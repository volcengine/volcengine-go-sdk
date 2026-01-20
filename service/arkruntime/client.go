package arkruntime

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/volcengine/volcengine-go-sdk/service/ark"
	"github.com/volcengine/volcengine-go-sdk/service/arkruntime/model"
	"github.com/volcengine/volcengine-go-sdk/service/arkruntime/pkg/encryption"
	"github.com/volcengine/volcengine-go-sdk/service/arkruntime/utils"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineerr"
)

type Client struct {
	config clientConfig

	requestBuilder utils.RequestBuilder

	arkClient               *ark.ARK
	resourceStsTokens       sync.Map
	rwLock                  *sync.RWMutex
	advisoryRefreshTimeout  int
	mandatoryRefreshTimeout int
	e2eeManager             sync.Map
	keyNonce                sync.Map

	batchHTTPClient      *http.Client
	modelBreakerProvider *utils.ModelBreakerProvider
}

func NewClientWithApiKey(apiKey string, setters ...ConfigOption) *Client {
	config := NewClientConfig(apiKey, "", "", setters...)
	return newClientWithConfig(config)
}

func NewClientWithAkSk(ak, sk string, setters ...ConfigOption) *Client {
	config := NewClientConfig("", ak, sk, setters...)
	return newClientWithConfig(config)
}

// NewClientWithConfig creates new API client for specified config.
func newClientWithConfig(config clientConfig) *Client {
	var arkClient *ark.ARK
	arkConfig := volcengine.NewConfig().
		WithCredentials(credentials.NewStaticCredentials(config.ak, config.sk, "")).
		WithRegion(config.region)

	sess, _ := session.NewSession(arkConfig)
	arkClient = ark.New(sess)

	return &Client{
		config:                  config,
		requestBuilder:          utils.NewRequestBuilder(),
		arkClient:               arkClient,
		resourceStsTokens:       sync.Map{},
		rwLock:                  &sync.RWMutex{},
		advisoryRefreshTimeout:  model.DefaultAdvisoryRefreshTimeout,
		mandatoryRefreshTimeout: model.DefaultMandatoryRefreshTimeout,
		e2eeManager:             sync.Map{},
		keyNonce:                sync.Map{},
		batchHTTPClient:         newBatchHTTPClient(config.batchMaxParallel),
		modelBreakerProvider:    utils.NewModelBreakerProvider(),
	}
}

func (c *Client) GetEndpointStsToken(ctx context.Context, endpointId string) (string, error) {
	return c.GetResourceStsToken(ctx, resourceTypeEndpoint, endpointId, "")
}

func (c *Client) GetResourceStsToken(ctx context.Context, resourceType string, resourceId string, projectName string) (string, error) {
	err := c.refresh(ctx, resourceType, resourceId, projectName)
	if err != nil {
		return "", err
	}

	token, ok := c.resourceStsTokens.Load(fmt.Sprintf(stsTokenKeyPattern, resourceType, resourceId))
	if ok {
		return token.(tokenInfo).token, nil
	}
	return "", nil
}

func (c *Client) refresh(ctx context.Context, resourceType string, resourceId string, projectName string) error {
	if !c.needRefresh(resourceType, resourceId, c.advisoryRefreshTimeout) {
		return nil
	}

	if c.rwLock.TryLock() {
		defer c.rwLock.Unlock()
		if !c.needRefresh(resourceType, resourceId, c.advisoryRefreshTimeout) {
			return nil
		}

		isMandatoryRefresh := c.needRefresh(resourceType, resourceId, c.mandatoryRefreshTimeout)
		return c.protectedRefresh(ctx, resourceType, resourceId, projectName, isMandatoryRefresh)
	} else if c.needRefresh(resourceType, resourceId, c.mandatoryRefreshTimeout) {
		c.rwLock.Lock()
		defer c.rwLock.Unlock()
		if !c.needRefresh(resourceType, resourceId, c.mandatoryRefreshTimeout) {
			return nil
		}
		return c.protectedRefresh(ctx, resourceType, resourceId, projectName, true)
	}
	return nil
}

func (c *Client) needRefresh(resourceType string, resourceId string, refreshIn int) bool {
	delta := c.advisoryRefreshTimeout
	if refreshIn > 0 {
		delta = refreshIn
	}

	token, ok := c.resourceStsTokens.Load(fmt.Sprintf(stsTokenKeyPattern, resourceType, resourceId))
	if ok {
		return token.(tokenInfo).expiredTime-time.Now().Unix() < int64(delta)
	}
	return true
}

func (c *Client) protectedRefresh(ctx context.Context, resourceType string, resourceId string, projectName string, isMandatory bool) error {
	input := ark.GetApiKeyInput{
		DurationSeconds: volcengine.Int32(model.DefaultStsTimeout),
		ResourceIds:     []*string{volcengine.String(resourceId)},
		ResourceType:    volcengine.String(resourceType),
	}

	if projectName != "" {
		input.ProjectName = volcengine.String(projectName)
	}

	resp, err := c.arkClient.GetApiKeyWithContext(ctx, &input)
	if err != nil {
		if isMandatory {
			return err
		} else {
			return nil
		}
	}
	c.resourceStsTokens.Store(fmt.Sprintf(stsTokenKeyPattern, resourceType, resourceId), tokenInfo{*resp.ApiKey, int64(*resp.ExpiredTime)})
	return nil
}

type requestOptions struct {
	body   interface{}
	header http.Header
	query  url.Values
}

type requestOption func(*requestOptions)

func withBody(body interface{}) requestOption {
	return func(args *requestOptions) {
		args.body = body
	}
}

func structToMap(obj interface{}) (map[string]interface{}, error) {
	data, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func withContentType(contentType string) requestOption {
	return func(args *requestOptions) {
		args.header.Set("Content-Type", contentType)
	}
}

func WithProjectName(project string) requestOption {
	return func(args *requestOptions) {
		args.header.Set("X-Project-Name", project)
	}
}

func WithCustomHeader(key, value string) requestOption {
	return func(args *requestOptions) {
		args.header.Set(key, value)
	}
}

func WithCustomHeaders(m map[string]string) requestOption {
	return func(args *requestOptions) {
		for k, v := range m {
			args.header.Set(k, v)
		}
	}
}

// WithQuery returns a RequestOption that sets the query value to the associated key. It overwrites
// any value if there was one already present.
func WithQuery(key, value string) requestOption {
	return func(args *requestOptions) {
		args.query.Set(key, value)
	}
}

func (c *Client) newRequest(ctx context.Context, method, url, resourceType, resourceId string, setters ...requestOption) (*http.Request, *model.RequestError) {
	// Default Options
	args := &requestOptions{
		body:   nil,
		header: make(http.Header),
	}

	requestID := utils.GenRequestId()
	args.header.Set(model.ClientRequestHeader, requestID)

	// parse resource type by resourceId
	// - endpoint: ep-*
	// - bot: bot-*
	// - presetendpoint: ep-m-* or modelID such as doubao-pro-32k-240525
	resourceType = c.getResourceTypeById(resourceId)

	for _, setter := range setters {
		setter(args)
	}

	errH := c.setCommonHeaders(ctx, args, resourceType, resourceId)
	if errH != nil {
		return nil, errH
	}

	// encrypt request body if is necessary
	if args.header.Get(model.ClientIsEncryptedHeader) == "true" {
		isStreaming := false
		if args.body != nil {
			if chatRequest, ok := args.body.(model.ChatRequest); ok {
				isStreaming = chatRequest.IsStream()
			}
		}
		if isStreaming {
			if err := c.encryptRequest(ctx, resourceId, args); err != nil {
				return nil, model.NewRequestError(http.StatusBadRequest, err, requestID)
			}
		}
	}

	// add query args
	if len(args.query) > 0 {
		url = url + "?" + args.query.Encode()
	}

	req, err := c.requestBuilder.Build(ctx, method, url, args.body, args.header)
	if err != nil {
		return nil, model.NewRequestError(http.StatusBadRequest, err, requestID)
	}

	return req, nil
}

func (c *Client) sendRequest(client *http.Client, req *http.Request, v model.Response) error {
	requestID := req.Header.Get(model.ClientRequestHeader)
	req.Header.Set("Accept", "application/json")

	// Check whether Content-Type is already set, Upload Files API requires
	// Content-Type == multipart/form-data
	contentType := req.Header.Get("Content-Type")
	if contentType == "" {
		req.Header.Set("Content-Type", "application/json")
	}

	res, err := client.Do(req)
	if err != nil {
		return model.NewRequestError(http.StatusInternalServerError, err, requestID)
	}

	defer res.Body.Close()

	if v != nil {
		v.SetHeader(res.Header)
	}

	if isFailureStatusCode(res) {
		return c.handleErrorResp(res)
	}

	err = decodeResponse(res.Body, v)
	if err != nil {
		err = &model.RequestError{
			HTTPStatusCode: res.StatusCode,
			Err:            err,
			RequestId:      requestID,
		}
	}
	keyNonce, ok := c.keyNonce.Load(requestID)
	if ok {
		err = encryption.DecryptChatResponse(keyNonce.([]byte), v)
		if err != nil {
			err = &model.RequestError{
				HTTPStatusCode: res.StatusCode,
				Err:            err,
				RequestId:      requestID,
			}
		}
	}
	return err
}

func (c *Client) Do(ctx context.Context, method, url, resourceType, resourceId string, v model.Response, setters ...requestOption) (err error) {
	err = utils.Retry(
		ctx,
		utils.RetryPolicy{
			MaxAttempts:    c.config.RetryTimes,
			InitialBackoff: model.ErrorRetryBaseDelay,
			MaxBackoff:     model.ErrorRetryMaxDelay,
		},
		func() bool { return true },
		func() error {
			req, innerErr := c.newRequest(ctx, method, url, resourceType, resourceId, setters...)
			if innerErr != nil {
				return innerErr
			}

			return c.sendRequest(c.config.HTTPClient, req, v)
		},
		nil,
		needRetryError,
	)
	return
}

func (c *Client) DoBatch(ctx context.Context, method, url, resourceType, resourceId string, v model.Response, setters ...requestOption) error {
	breaker := c.modelBreakerProvider.GetOrCreateBreaker(resourceId)

	for {
		breaker.Wait()

		select {
		case <-ctx.Done(): // whole context finish
			return ctx.Err()
		default:
		}

		err := func() error {
			req, er := c.newRequest(ctx, method, url, resourceType, resourceId, setters...)
			if er != nil {
				return er
			}

			return c.sendRequest(c.batchHTTPClient, req, v)
		}()

		// no error: just return on this try
		if err == nil {
			return nil
		}

		// no need to retry error
		if !needRetryError(err) {
			return err
		}

		retryAfter := c.getRetryAfter(v)
		if retryAfter > 0 {
			breaker.Reset(time.Duration(retryAfter) * time.Second)
		}

		time.Sleep(time.Duration(500+rand.Intn(1001)) * time.Millisecond)
	}
}

func sendChatCompletionRequestStream(client *Client, httpClient *http.Client, req *http.Request) (*utils.ChatCompletionStreamReader, error) {
	requestID := req.Header.Get(model.ClientRequestHeader)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "text/event-stream")
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Connection", "keep-alive")

	resp, err := httpClient.Do(req) //nolint:bodyclose // body is closed in stream.Close()
	if err != nil {
		return &utils.ChatCompletionStreamReader{}, model.NewRequestError(http.StatusInternalServerError, err, requestID)
	}
	if isFailureStatusCode(resp) {
		return &utils.ChatCompletionStreamReader{}, client.handleErrorResp(resp)
	}

	keyNonce, ok := client.keyNonce.Load(resp.Header.Get(model.ClientRequestHeader))
	if !ok {
		keyNonce = []byte{}
	}

	cleanup := func() {
		client.keyNonce.Delete(requestID)
	}

	return &utils.ChatCompletionStreamReader{
		EmptyMessagesLimit: client.config.EmptyMessagesLimit,
		Reader:             bufio.NewReader(resp.Body),
		Response:           resp,
		ErrAccumulator:     utils.NewErrorAccumulator(),
		Unmarshaler:        &utils.JSONUnmarshaler{},
		HttpHeader:         model.HttpHeader(resp.Header),
		KeyNonce:           keyNonce.([]byte),
		RequestID:          requestID,
		Cleanup:            cleanup,
	}, nil
}

func sendBotChatCompletionRequestStream(client *Client, httpClient *http.Client, req *http.Request) (*utils.BotChatCompletionStreamReader, error) {
	requestID := req.Header.Get(model.ClientRequestHeader)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "text/event-stream")
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Connection", "keep-alive")

	resp, err := httpClient.Do(req) //nolint:bodyclose // body is closed in stream.Close()
	if err != nil {
		return &utils.BotChatCompletionStreamReader{}, model.NewRequestError(http.StatusInternalServerError, err, requestID)
	}
	if isFailureStatusCode(resp) {
		return &utils.BotChatCompletionStreamReader{}, client.handleErrorResp(resp)
	}
	return &utils.BotChatCompletionStreamReader{
		ChatCompletionStreamReader: utils.ChatCompletionStreamReader{
			EmptyMessagesLimit: client.config.EmptyMessagesLimit,
			Reader:             bufio.NewReader(resp.Body),
			Response:           resp,
			ErrAccumulator:     utils.NewErrorAccumulator(),
			Unmarshaler:        &utils.JSONUnmarshaler{},
			HttpHeader:         model.HttpHeader(resp.Header),
		},
	}, nil
}

func sendCreateResponsesRequestStream(client *Client, httpClient *http.Client, req *http.Request) (*utils.ResponsesStreamReader, error) {
	requestID := req.Header.Get(model.ClientRequestHeader)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "text/event-stream")
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Connection", "keep-alive")

	resp, err := httpClient.Do(req) //nolint:bodyclose // body is closed in stream.Close()
	if err != nil {
		return &utils.ResponsesStreamReader{}, model.NewRequestError(http.StatusInternalServerError, err, requestID)
	}
	if isFailureStatusCode(resp) {
		return &utils.ResponsesStreamReader{}, client.handleErrorResp(resp)
	}
	return &utils.ResponsesStreamReader{
		ChatCompletionStreamReader: utils.ChatCompletionStreamReader{
			EmptyMessagesLimit: client.config.EmptyMessagesLimit,
			Response:           resp,
			ErrAccumulator:     utils.NewErrorAccumulator(),
			Unmarshaler:        &utils.JSONUnmarshaler{},
			HttpHeader:         model.HttpHeader(resp.Header),
		},
		Decoder: utils.NewEventStreamDecoder(resp.Body),
	}, nil
}

func sendImageGenerationStream(client *Client, httpClient *http.Client, req *http.Request) (*utils.ImageGenerationStreamReader, error) {
	requestID := req.Header.Get(model.ClientRequestHeader)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "text/event-stream")
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Connection", "keep-alive")

	resp, err := httpClient.Do(req) //nolint:bodyclose // body is closed in stream.Close()
	if err != nil {
		return &utils.ImageGenerationStreamReader{}, model.NewRequestError(http.StatusInternalServerError, err, requestID)
	}
	if isFailureStatusCode(resp) {
		return &utils.ImageGenerationStreamReader{}, client.handleErrorResp(resp)
	}
	return &utils.ImageGenerationStreamReader{
		ChatCompletionStreamReader: utils.ChatCompletionStreamReader{
			EmptyMessagesLimit: client.config.EmptyMessagesLimit,
			Reader:             bufio.NewReader(resp.Body),
			Response:           resp,
			ErrAccumulator:     utils.NewErrorAccumulator(),
			Unmarshaler:        &utils.JSONUnmarshaler{},
			HttpHeader:         model.HttpHeader(resp.Header),
		},
	}, nil
}

func (c *Client) BotChatCompletionRequestStreamDo(ctx context.Context, method, url, botId string, setters ...requestOption) (streamReader *utils.BotChatCompletionStreamReader, err error) {
	err = utils.Retry(
		ctx,
		utils.RetryPolicy{
			MaxAttempts:    c.config.RetryTimes,
			InitialBackoff: model.ErrorRetryBaseDelay,
			MaxBackoff:     model.ErrorRetryMaxDelay,
		},
		func() bool { return true },
		func() error {
			req, innerErr := c.newRequest(ctx, method, url, resourceTypeBot, botId, setters...)
			if innerErr != nil {
				return innerErr
			}

			streamReader, err = sendBotChatCompletionRequestStream(c, c.config.HTTPClient, req)
			return err
		},
		nil,
		needRetryError,
	)

	return
}

func (c *Client) ChatCompletionRequestStreamDo(ctx context.Context, method, url, resourceId string, setters ...requestOption) (streamReader *utils.ChatCompletionStreamReader, err error) {
	err = utils.Retry(
		ctx,
		utils.RetryPolicy{
			MaxAttempts:    c.config.RetryTimes,
			InitialBackoff: model.ErrorRetryBaseDelay,
			MaxBackoff:     model.ErrorRetryMaxDelay,
		},
		func() bool { return true },
		func() error {

			req, innerErr := c.newRequest(ctx, method, url, resourceTypeEndpoint, resourceId, setters...)
			if innerErr != nil {
				return innerErr
			}

			streamReader, err = sendChatCompletionRequestStream(c, c.config.HTTPClient, req)
			return err
		},
		nil,
		needRetryError,
	)

	return
}

// ResponsesRequestStreamDo executes a request.
func (c *Client) ResponsesRequestStreamDo(ctx context.Context, method, url, resourceType, resourceId string, setters ...requestOption) (resp *utils.ResponsesStreamReader, err error) {
	err = utils.Retry(
		ctx,
		utils.RetryPolicy{
			MaxAttempts:    c.config.RetryTimes,
			InitialBackoff: model.ErrorRetryBaseDelay,
			MaxBackoff:     model.ErrorRetryMaxDelay,
		},
		func() bool { return true },
		func() error {
			req, innerErr := c.newRequest(ctx, method, url, resourceType, resourceId, setters...)
			if innerErr != nil {
				return innerErr
			}
			resp, err = sendCreateResponsesRequestStream(c, c.config.HTTPClient, req)
			return err
		},
		nil,
		needRetryError,
	)
	return
}

func (c *Client) ImageGenerationStreamDo(ctx context.Context, method, url, resourceId string, setters ...requestOption) (streamReader *utils.ImageGenerationStreamReader, err error) {
	err = utils.Retry(
		ctx,
		utils.RetryPolicy{
			MaxAttempts:    c.config.RetryTimes,
			InitialBackoff: model.ErrorRetryBaseDelay,
			MaxBackoff:     model.ErrorRetryMaxDelay,
		},
		func() bool { return true },
		func() error {
			req, innerErr := c.newRequest(ctx, method, url, resourceTypeEndpoint, resourceId, setters...)
			if innerErr != nil {
				return innerErr
			}

			streamReader, err = sendImageGenerationStream(c, c.config.HTTPClient, req)

			return err
		},
		nil,
		needRetryError,
	)
	return
}

func (c *Client) setCommonHeaders(ctx context.Context, args *requestOptions, resourceType string, resourceId string) *model.RequestError {
	requestID := args.header.Get(model.ClientRequestHeader)
	if len(c.config.apiKey) > 0 {
		args.header.Set("Authorization", fmt.Sprintf("Bearer %s", c.config.apiKey))
	} else {
		if resourceTypeEndpoint == resourceType && !strings.HasPrefix(resourceId, "ep-") {
			return model.NewRequestError(http.StatusBadRequest, model.ErrBodyWithoutEndpoint, requestID)
		} else if resourceTypeBot == resourceType && !strings.HasPrefix(resourceId, "bot-") {
			return model.NewRequestError(http.StatusBadRequest, model.ErrBodyWithoutBot, requestID)
		}

		projectName := args.header.Get("X-Project-Name")
		if resourceTypePresetEndpoint == resourceType && projectName == "" {
			return model.NewRequestError(http.StatusBadRequest, model.ErrBodyWithoutProjectName, requestID)
		}

		token, err := c.GetResourceStsToken(ctx, resourceType, resourceId, projectName)
		if err != nil {
			if volcErr, ok := err.(volcengineerr.RequestFailure); ok {
				return model.NewRequestError(volcErr.StatusCode(), fmt.Errorf("failed to get resource sts token. err=%w", volcErr), volcErr.RequestID())
			}
			return model.NewRequestError(http.StatusInternalServerError, fmt.Errorf("failed to get resource sts token. err=%w", err), requestID)
		}
		args.header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	}
	return nil
}

func (c *Client) getResourceTypeById(resourceId string) string {
	switch {
	case strings.HasPrefix(resourceId, "ep-m-"):
		return resourceTypePresetEndpoint
	case strings.HasPrefix(resourceId, "ep-"):
		return resourceTypeEndpoint
	case strings.HasPrefix(resourceId, "bot-"):
		return resourceTypeBot
	default:
		return resourceTypePresetEndpoint
	}
}

func isFailureStatusCode(resp *http.Response) bool {
	return resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusBadRequest
}

func needRetryError(err error) bool {
	apiErr := &model.APIError{}
	reqErr := &model.RequestError{}
	if errors.As(err, &apiErr) {
		return apiErr.HTTPStatusCode >= http.StatusInternalServerError || apiErr.HTTPStatusCode == http.StatusTooManyRequests
	} else if errors.Is(err, io.EOF) {
		return true
	} else if errors.As(err, &reqErr) {
		return reqErr.HTTPStatusCode >= http.StatusInternalServerError
	}
	return false
}

func decodeResponse(body io.Reader, v interface{}) error {
	if v == nil {
		return nil
	}

	switch o := v.(type) {
	case *string:
		return decodeString(body, o)
	default:
		return json.NewDecoder(body).Decode(v)
	}
}

func decodeString(body io.Reader, output *string) error {
	b, err := io.ReadAll(body)
	if err != nil {
		return err
	}
	*output = string(b)
	return nil
}

func (c *Client) fullURL(suffix string) string {
	return fmt.Sprintf("%s%s", c.config.BaseURL, suffix)
}

func (c *Client) handleErrorResp(resp *http.Response) error {
	requestID := resp.Header.Get(model.ClientRequestHeader)
	var errRes model.ErrorResponse
	err := json.NewDecoder(resp.Body).Decode(&errRes)
	if err != nil || errRes.Error == nil {
		reqErr := model.NewRequestError(resp.StatusCode, err, requestID)
		if errRes.Error != nil {
			reqErr.Err = errRes.Error
		}
		return reqErr
	}

	errRes.Error.HTTPStatusCode = resp.StatusCode
	errRes.Error.RequestId = requestID
	return errRes.Error
}

func (c *Client) getRetryAfter(v model.Response) int64 {
	header := v.GetHeader()
	retryAfter := header[model.RetryAfterHeader]
	if len(retryAfter) == 0 || retryAfter[0] == "" {
		return 0
	}
	retryAfterInterval, err := strconv.ParseInt(retryAfter[0], 10, 64)
	if err != nil {
		return 0
	}
	return retryAfterInterval
}

func (c *Client) isAPIKeyAuthentication() bool {
	return c.config.apiKey != ""
}

func (c *Client) encryptRequest(ctx context.Context, resourceId string, args *requestOptions) error {
	e2eeClient, err := c.getE2eeClient(ctx, resourceId, args.header.Get("Authorization"))
	if err != nil {
		return err
	}
	keyNonce, sessionToken, err := e2eeClient.GenerateECIESKeyPair()
	if err != nil {
		return err
	}
	// add session token to header
	args.header.Set(model.ClientSessionTokenHeader, sessionToken)
	// store keyNonce to map
	c.rwLock.Lock()
	c.keyNonce.Store(args.header.Get(model.ClientRequestHeader), keyNonce)
	c.rwLock.Unlock()
	// encrypt request body
	err = encryption.EncryptChatRequest(ctx, keyNonce, args.body.(model.CreateChatCompletionRequest))
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) getE2eeClient(ctx context.Context, resourceId, auth string) (*E2eeClient, error) {
	cert, ok := c.e2eeManager.Load(resourceId)
	if ok {
		client := cert.(*E2eeClient)
		if client.isAICC == encryption.CheckIsModeAICC() {
			return client, nil
		}
	}
	// load by local file
	certPem, err := encryption.LoadLocalCertificate(resourceId)
	if err != nil || certPem == "" {
		// load from client request and save file to local
		certPem, err = c.loadServerCertificate(ctx, resourceId, auth)
		if err != nil {
			return nil, fmt.Errorf("loading Certificate failed: %w", err)
		}
		// save to local file
		err = encryption.SaveToLocalCertificate(resourceId, certPem)
		if err != nil {
			return nil, fmt.Errorf("saving Certificate failed: %w", err)
		}
	}
	e2eeClient, err := NewE2eeClient(certPem)
	if err != nil {
		return nil, fmt.Errorf("creating E2eeClient failed: %w", err)
	}
	// store to e2eeManager
	c.rwLock.Lock()
	defer c.rwLock.Unlock()
	c.e2eeManager.Store(resourceId, e2eeClient)
	return e2eeClient, nil
}

func (c *Client) loadServerCertificate(ctx context.Context, resourceId, auth string) (string, error) {
	url := c.fullURL("/e2e/get/certificate")
	requestBody := map[string]interface{}{"model": resourceId}
	// Check if AICC mode is enabled
	if encryption.CheckIsModeAICC() {
		requestBody["type"] = "AICCv0.1"
	}
	b, err := json.Marshal(requestBody)
	if err != nil {
		return "", fmt.Errorf("getting Certificate failed: %w", err)
	}
	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewReader(b))
	if err != nil {
		return "", fmt.Errorf("getting Certificate failed: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", auth)
	req.Header.Set("X-Session-Token", "/e2e/get/certificate")
	res, err := c.config.HTTPClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("getting Certificate failed: %w", err)
	}
	defer res.Body.Close()
	var cr model.CertificateResponse
	if err := json.NewDecoder(res.Body).Decode(&cr); err != nil {
		return "", fmt.Errorf("getting Certificate failed: %w", err)
	}
	if len(cr.Error) > 0 {
		return "", fmt.Errorf("getting Certificate failed: %v", cr.Error)
	}
	return cr.Certificate, nil
}

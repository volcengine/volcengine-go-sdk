package arkruntime

import (
	"bufio"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/volcengine/volcengine-go-sdk/service/ark"
	"github.com/volcengine/volcengine-go-sdk/service/arkruntime/model"
	"github.com/volcengine/volcengine-go-sdk/service/arkruntime/utils"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)

type Client struct {
	config ClientConfig

	requestBuilder utils.RequestBuilder

	arkClient               *ark.ARK
	resourceStsTokens       sync.Map
	rwLock                  *sync.RWMutex
	advisoryRefreshTimeout  int
	mandatoryRefreshTimeout int
}

func NewClientWithApiKey(apiKey string, setters ...configOption) *Client {
	config := NewClientConfig(apiKey, "", "", setters...)
	return newClientWithConfig(config)
}

func NewClientWithAkSk(ak, sk string, setters ...configOption) *Client {
	config := NewClientConfig("", ak, sk, setters...)
	return newClientWithConfig(config)
}

// NewClientWithConfig creates new API client for specified config.
func newClientWithConfig(config ClientConfig) *Client {
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
	}
}

func (c *Client) GetEndpointStsToken(ctx context.Context, endpointId string) (string, error) {
	return c.GetResourceStsToken(ctx, resourceTypeEndpoint, endpointId)
}

func (c *Client) GetResourceStsToken(ctx context.Context, resourceType string, resourceId string) (string, error) {
	err := c.refresh(ctx, resourceType, resourceId)
	if err != nil {
		return "", err
	}

	token, ok := c.resourceStsTokens.Load(fmt.Sprintf(stsTokenKeyPattern, resourceType, resourceId))
	if ok {
		return token.(tokenInfo).token, nil
	}
	return "", nil
}

func (c *Client) refresh(ctx context.Context, resourceType string, resourceId string) error {
	if !c.needRefresh(resourceType, resourceId, c.advisoryRefreshTimeout) {
		return nil
	}

	if c.rwLock.TryLock() {
		defer c.rwLock.Unlock()
		if !c.needRefresh(resourceType, resourceId, c.advisoryRefreshTimeout) {
			return nil
		}

		isMandatoryRefresh := c.needRefresh(resourceType, resourceId, c.mandatoryRefreshTimeout)
		return c.protectedRefresh(ctx, resourceType, resourceId, isMandatoryRefresh)
	} else if c.needRefresh(resourceType, resourceId, c.mandatoryRefreshTimeout) {
		c.rwLock.Lock()
		defer c.rwLock.Unlock()
		if !c.needRefresh(resourceType, resourceId, c.mandatoryRefreshTimeout) {
			return nil
		}
		return c.protectedRefresh(ctx, resourceType, resourceId, true)
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

func (c *Client) protectedRefresh(ctx context.Context, resourceType string, resourceId string, isMandatory bool) error {
	resp, err := c.arkClient.GetApiKeyWithContext(ctx, &ark.GetApiKeyInput{
		DurationSeconds: volcengine.Int32(model.DefaultStsTimeout),
		ResourceIds:     []*string{volcengine.String(resourceId)},
		ResourceType:    volcengine.String(resourceType),
	})
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
}

type requestOption func(*requestOptions)

func withBody(body interface{}) requestOption {
	return func(args *requestOptions) {
		args.body = body
	}
}

func withContentType(contentType string) requestOption {
	return func(args *requestOptions) {
		args.header.Set("Content-Type", contentType)
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

func (c *Client) newRequest(ctx context.Context, method, url, resourceType, resourceId string, setters ...requestOption) (*http.Request, error) {
	// Default Options
	args := &requestOptions{
		body:   nil,
		header: make(http.Header),
	}

	errH := c.setCommonHeaders(ctx, args, resourceType, resourceId)
	if errH != nil {
		return nil, errH
	}

	for _, setter := range setters {
		setter(args)
	}
	req, err := c.requestBuilder.Build(ctx, method, url, args.body, args.header)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func (c *Client) sendRequest(req *http.Request, v model.Response) error {
	requestId := req.Header.Get(model.ClientRequestHeader)
	req.Header.Set("Accept", "application/json")

	// Check whether Content-Type is already set, Upload Files API requires
	// Content-Type == multipart/form-data
	contentType := req.Header.Get("Content-Type")
	if contentType == "" {
		req.Header.Set("Content-Type", "application/json")
	}

	res, err := c.config.HTTPClient.Do(req)
	if err != nil {
		return &model.RequestError{
			HTTPStatusCode: 500,
			Err:            err,
			RequestId:      requestId,
		}
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
			RequestId:      requestId,
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

			return c.sendRequest(req, v)
		},
		nil,
		needRetryError,
	)
	return
}

func (c *Client) sendRequestRaw(req *http.Request) (response model.RawResponse, err error) {
	requestId := req.Header.Get(model.ClientRequestHeader)
	resp, err := c.config.HTTPClient.Do(req) //nolint:bodyclose // body should be closed by outer function
	if err != nil {
		err = &model.RequestError{
			HTTPStatusCode: 500,
			Err:            err,
			RequestId:      requestId,
		}
		return
	}

	if isFailureStatusCode(resp) {
		err = c.handleErrorResp(resp)
		return
	}

	response.SetHeader(resp.Header)
	response.ReadCloser = resp.Body
	return
}

func sendChatCompletionRequestStream(client *Client, req *http.Request) (*utils.ChatCompletionStreamReader, error) {
	requestId := req.Header.Get(model.ClientRequestHeader)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "text/event-stream")
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Connection", "keep-alive")

	resp, err := client.config.HTTPClient.Do(req) //nolint:bodyclose // body is closed in stream.Close()
	if err != nil {
		return &utils.ChatCompletionStreamReader{}, &model.RequestError{
			HTTPStatusCode: 500,
			Err:            err,
			RequestId:      requestId,
		}
	}
	if isFailureStatusCode(resp) {
		return &utils.ChatCompletionStreamReader{}, client.handleErrorResp(resp)
	}
	return &utils.ChatCompletionStreamReader{
		EmptyMessagesLimit: client.config.EmptyMessagesLimit,
		Reader:             bufio.NewReader(resp.Body),
		Response:           resp,
		ErrAccumulator:     utils.NewErrorAccumulator(),
		Unmarshaler:        &utils.JSONUnmarshaler{},
		HttpHeader:         model.HttpHeader(resp.Header),
	}, nil
}

func sendBotChatCompletionRequestStream(client *Client, req *http.Request) (*utils.BotChatCompletionStreamReader, error) {
	requestId := req.Header.Get(model.ClientRequestHeader)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "text/event-stream")
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Connection", "keep-alive")

	resp, err := client.config.HTTPClient.Do(req) //nolint:bodyclose // body is closed in stream.Close()
	if err != nil {
		return &utils.BotChatCompletionStreamReader{}, &model.RequestError{
			HTTPStatusCode: 500,
			Err:            err,
			RequestId:      requestId,
		}
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

			streamReader, err = sendBotChatCompletionRequestStream(c, req)
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

			streamReader, err = sendChatCompletionRequestStream(c, req)
			return err
		},
		nil,
		needRetryError,
	)

	return
}

func (c *Client) setCommonHeaders(ctx context.Context, args *requestOptions, resourceType string, resourceId string) error {
	args.header.Set(model.ClientRequestHeader, utils.GenRequestId())

	if len(c.config.apiKey) > 0 {
		args.header.Set("Authorization", fmt.Sprintf("Bearer %s", c.config.apiKey))
	} else {
		if resourceTypeEndpoint == resourceType && !strings.HasPrefix(resourceId, "ep-") {
			return model.ErrBodyWithoutEndpoint
		} else if resourceTypeBot == resourceType && !strings.HasPrefix(resourceId, "bot-") {
			return model.ErrBodyWithoutBot
		}
		token, err := c.GetResourceStsToken(ctx, resourceType, resourceId)
		if err != nil {
			return fmt.Errorf("failed to get resource sts token. err=%v", err)
		}
		args.header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	}
	return nil
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
	requestId := resp.Header.Get(model.ClientRequestHeader)
	var errRes model.ErrorResponse
	err := json.NewDecoder(resp.Body).Decode(&errRes)
	if err != nil || errRes.Error == nil {
		reqErr := &model.RequestError{
			HTTPStatusCode: resp.StatusCode,
			Err:            err,
			RequestId:      requestId,
		}
		if errRes.Error != nil {
			reqErr.Err = errRes.Error
		}
		return reqErr
	}

	errRes.Error.HTTPStatusCode = resp.StatusCode
	errRes.Error.RequestId = requestId
	return errRes.Error
}

package arkruntime

import (
	"context"
	"net/http"

	"github.com/volcengine/volcengine-go-sdk/service/arkruntime/model"
)

const (
	batchChatCompletionsSuffix      = "/batch/chat/completions"
	batchEmbeddingsSuffix           = "/batch/embeddings"
	batchMultiModalEmbeddingsSuffix = "/batch/embeddings/multimodal"
)

func newBatchHTTPClient(maxParallel int) *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			MaxConnsPerHost: maxParallel,
		},
	}
}

// CreateBatchChatCompletion — API call to Create a batch completion for the chat message.
func (c *Client) CreateBatchChatCompletion(
	ctx context.Context,
	request model.ChatRequest,
	setters ...requestOption,
) (response model.ChatCompletionResponse, err error) {
	if request.IsStream() {
		err = model.ErrChatCompletionStreamNotSupported
		return
	}
	requestOptions := append(setters, withBody(request))
	err = c.DoBatch(ctx, http.MethodPost, c.fullURL(batchChatCompletionsSuffix), resourceTypeEndpoint, request.GetModel(), &response, requestOptions...)
	if err != nil {
		return
	}
	return
}

// CreateBatchEmbeddings like CreateEmbeddings, but for batch processing.
func (c *Client) CreateBatchEmbeddings(
	ctx context.Context,
	conv model.EmbeddingRequestConverter,
	setters ...requestOption,
) (res model.EmbeddingResponse, err error) {
	baseReq := conv.Convert()

	requestOptions := append(setters, withBody(baseReq))
	if baseReq.EncodingFormat != model.EmbeddingEncodingFormatBase64 {
		err = c.DoBatch(ctx, http.MethodPost, c.fullURL(batchEmbeddingsSuffix), resourceTypeEndpoint, baseReq.Model, &res, requestOptions...)
		return
	}

	base64Response := &model.EmbeddingResponseBase64{}
	err = c.DoBatch(ctx, http.MethodPost, c.fullURL(batchEmbeddingsSuffix), resourceTypeEndpoint, baseReq.Model, base64Response, requestOptions...)
	if err != nil {
		return
	}

	res, err = base64Response.ToEmbeddingResponse()
	return
}

// CreateBatchMultiModalEmbeddings like CreateMultiModalEmbeddings, but for batch processing.
func (c *Client) CreateBatchMultiModalEmbeddings(
	ctx context.Context,
	request model.MultiModalEmbeddingRequest,
	setters ...requestOption,
) (res model.MultimodalEmbeddingResponse, err error) {

	requestOptions := append(setters, withBody(request))
	if request.EncodingFormat != nil && *request.EncodingFormat == model.EmbeddingEncodingFormatBase64 {
		base64Response := &model.MultiModalEmbeddingResponseBase64{}
		err = c.DoBatch(ctx, http.MethodPost, c.fullURL(batchMultiModalEmbeddingsSuffix), resourceTypeEndpoint, request.Model, base64Response, requestOptions...)
		if err != nil {
			return
		}
		res, err = base64Response.ToMultiModalEmbeddingResponse()
		return
	}

	err = c.DoBatch(ctx, http.MethodPost, c.fullURL(batchMultiModalEmbeddingsSuffix), resourceTypeEndpoint, request.Model, &res, requestOptions...)
	return
}

// Deprecated: use `arkruntime.WithBatchMaxParallel` option in `arkruntime.NewClientXXX` instead.
func (c *Client) StartBatchWorker(workerNum int) {
	if transport, ok := c.batchHTTPClient.Transport.(*http.Transport); ok {
		transport.MaxConnsPerHost = workerNum
	}
}

// Deprecated: no need to stop batch worker.
func (c *Client) StopBatchWorker() {
}

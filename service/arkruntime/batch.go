package arkruntime

import (
	"context"
	"net/http"

	"github.com/volcengine/volcengine-go-sdk/service/arkruntime/model"
)

const batchChatCompletionsSuffix = "/batch/chat/completions"

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

// Deprecated: use `arkruntime.WithBatchMaxParallel` option in `arkruntime.NewClientXXX` instead.
func (c *Client) StartBatchWorker(workerNum int) {
	if transport, ok := c.batchHTTPClient.Transport.(*http.Transport); ok {
		transport.MaxConnsPerHost = workerNum
	}
}

// Deprecated: no need to stop batch worker.
func (c *Client) StopBatchWorker() {
}

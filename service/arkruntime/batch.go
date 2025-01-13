package arkruntime

import (
	"context"
	"net/http"

	"github.com/volcengine/volcengine-go-sdk/service/arkruntime/model"
	"github.com/volcengine/volcengine-go-sdk/service/arkruntime/utils"
)

const batchChatCompletionsSuffix = "/batch/chat/completions"

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
	ctxWithTimeout, cancel := context.WithTimeout(ctx, c.config.HTTPClient.Timeout)
	defer cancel()
	requestOptions := append(setters, withBody(request))
	err = c.DoBatch(ctxWithTimeout, http.MethodPost, c.fullURL(batchChatCompletionsSuffix), resourceTypeEndpoint, request.GetModel(), &response, requestOptions...)
	if err != nil {
		return
	}
	return
}

func (c *Client) StartBatchWorker(workerNum int) {
	c.batchWorkerOnce.Do(func() {
		c.batchWorkerPool = utils.NewBatchWorkerPool(workerNum)
		go c.batchWorkerPool.Run()
	})
}

func (c *Client) StopBatchWorker() {
	if c.batchWorkerPool != nil {
		c.batchWorkerPool.Close()
	}
}

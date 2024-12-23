package arkruntime

import (
	"context"
	"net/http"

	"github.com/volcengine/volcengine-go-sdk/service/arkruntime/model"
	"github.com/volcengine/volcengine-go-sdk/service/arkruntime/utils"
)

const contextCreateSuffix = "/context/create"
const contextChatSuffix = "/context/chat/completions"

// CreateContext — API call to Create a new context cache instance.
func (c *Client) CreateContext(
	goCtx context.Context,
	request model.CreateContextRequest,
	setters ...requestOption,
) (response model.CreateContextResponse, err error) {

	requestOptions := append(setters, withBody(request))
	err = c.Do(goCtx, http.MethodPost, c.fullURL(contextCreateSuffix), resourceTypeEndpoint, request.Model, &response, requestOptions...)
	if err != nil {
		return
	}
	return
}

// CreateContextChatCompletion — API call to Create a completion for chat with context cache.
func (c *Client) CreateContextChatCompletion(
	ctx context.Context,
	request model.ContextChatCompletionRequest,
	setters ...requestOption,
) (response model.ChatCompletionResponse, err error) {

	if request.Stream {
		err = model.ErrChatCompletionStreamNotSupported
		return
	}

	requestOptions := append(setters, withBody(request))
	err = c.Do(ctx, http.MethodPost, c.fullURL(contextChatSuffix), resourceTypeEndpoint, request.Model, &response, requestOptions...)
	if err != nil {
		return
	}
	return
}

// CreateContextChatCompletionStream — API call to create a chat completion w/ streaming
// support. It sets whether to stream back partial progress. If set, tokens will be
// sent as data-only server-sent events as they become available, with the
// stream terminated by a data: [DONE] message.
func (c *Client) CreateContextChatCompletionStream(
	ctx context.Context,
	request model.ContextChatCompletionRequest,
	setters ...requestOption,
) (stream *utils.ChatCompletionStreamReader, err error) {

	request.Stream = true

	requestOptions := append(setters, withBody(request))

	resp, err := c.ChatCompletionRequestStreamDo(ctx, http.MethodPost, c.fullURL(contextChatSuffix), request.Model, requestOptions...)
	if err != nil {
		return
	}
	stream = resp
	return
}

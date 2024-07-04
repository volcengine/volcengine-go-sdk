package arkruntime

import (
	"context"
	"net/http"

	"github.com/volcengine/volcengine-go-sdk/service/arkruntime/model"
	"github.com/volcengine/volcengine-go-sdk/service/arkruntime/utils"
)

const chatCompletionsSuffix = "/chat/completions"

// CreateChatCompletion — API call to Create a completion for the chat message.
func (c *Client) CreateChatCompletion(
	ctx context.Context,
	request model.ChatCompletionRequest,
	setters ...requestOption,
) (response model.ChatCompletionResponse, err error) {
	if request.Stream {
		err = model.ErrChatCompletionStreamNotSupported
		return
	}

	requestOptions := append(setters, withBody(request))
	err = c.Do(ctx, http.MethodPost, c.fullURL(chatCompletionsSuffix), resourceTypeEndpoint, request.Model, &response, requestOptions...)
	if err != nil {
		return
	}
	return
}

// CreateChatCompletionStream — API call to create a chat completion w/ streaming
// support. It sets whether to stream back partial progress. If set, tokens will be
// sent as data-only server-sent events as they become available, with the
// stream terminated by a data: [DONE] message.
func (c *Client) CreateChatCompletionStream(
	ctx context.Context,
	request model.ChatCompletionRequest,
	setters ...requestOption,
) (stream *utils.ChatCompletionStreamReader, err error) {
	request.Stream = true

	requestOptions := append(setters, withBody(request))

	resp, err := c.ChatCompletionRequestStreamDo(ctx, http.MethodPost, c.fullURL(chatCompletionsSuffix), request.Model, requestOptions...)
	if err != nil {
		return
	}
	stream = resp
	return
}

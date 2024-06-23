package arkruntime

import (
	"context"
	"net/http"

	"github.com/volcengine/volcengine-go-sdk/service/arkruntime/model"
	"github.com/volcengine/volcengine-go-sdk/service/arkruntime/utils"
)

const botChatCompletionsSuffix = "/bots/chat/completions"

// CreateBotChatCompletion — API call to Create a completion for the bot chat.
func (c *Client) CreateBotChatCompletion(
	ctx context.Context,
	request model.BotChatCompletionRequest,
	setters ...requestOption,
) (response model.BotChatCompletionResponse, err error) {
	// set BotId to request.Model
	if request.BotId != "" {
		request.Model = request.BotId
	}

	if request.Stream {
		err = model.ErrChatCompletionStreamNotSupported
		return
	}

	requestOptions := append(setters, withBody(request))
	err = c.Do(ctx, http.MethodPost, c.fullURL(botChatCompletionsSuffix), resourceTypeBot, request.Model, &response, requestOptions...)
	if err != nil {
		return
	}
	return
}

// CreateBotChatCompletionStream — API call to create a chat completion w/ streaming
// support. It sets whether to stream back partial progress. If set, tokens will be
// sent as data-only server-sent events as they become available, with the
// stream terminated by a data: [DONE] message.
func (c *Client) CreateBotChatCompletionStream(
	ctx context.Context,
	request model.BotChatCompletionRequest,
	setters ...requestOption,
) (stream *utils.BotChatCompletionStreamReader, err error) {
	// set BotId to request.Model
	if request.BotId != "" {
		request.Model = request.BotId
	}

	request.Stream = true

	requestOptions := append(setters, withBody(request))

	resp, err := c.BotChatCompletionRequestStreamDo(ctx, http.MethodPost, c.fullURL(botChatCompletionsSuffix), request.Model, requestOptions...)
	if err != nil {
		return
	}
	stream = resp
	return
}

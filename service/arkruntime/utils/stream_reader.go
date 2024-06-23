package utils

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"net/http"

	"github.com/volcengine/volcengine-go-sdk/service/arkruntime/model"
)

var (
	headerData  = []byte("data: ")
	errorPrefix = []byte(`data: {"error":`)
)

type ChatCompletionStreamReader struct {
	EmptyMessagesLimit uint
	IsFinished         bool

	Reader         *bufio.Reader
	Response       *http.Response
	ErrAccumulator ErrorAccumulator
	Unmarshaler    Unmarshaler

	model.HttpHeader
}

type BotChatCompletionStreamReader struct {
	ChatCompletionStreamReader
}

func (stream *ChatCompletionStreamReader) Recv() (response model.ChatCompletionStreamResponse, err error) {
	if stream.IsFinished {
		err = io.EOF
		return
	}

	response, err = stream.processLines()
	return
}

func (stream *BotChatCompletionStreamReader) Recv() (response model.BotChatCompletionStreamResponse, err error) {
	if stream.IsFinished {
		err = io.EOF
		return
	}

	response, err = stream.processLines()
	return
}

func (stream *ChatCompletionStreamReader) processLines() (model.ChatCompletionStreamResponse, error) {
	var (
		emptyMessagesCount uint
		hasErrorPrefix     bool
	)

	for {
		rawLine, readErr := stream.Reader.ReadBytes('\n')
		if readErr != nil || hasErrorPrefix {
			respErr := stream.unmarshalError()
			if respErr != nil {
				return model.ChatCompletionStreamResponse{}, fmt.Errorf("error, %w", respErr.Error)
			}
			return model.ChatCompletionStreamResponse{}, readErr
		}

		noSpaceLine := bytes.TrimSpace(rawLine)
		if bytes.HasPrefix(noSpaceLine, errorPrefix) {
			hasErrorPrefix = true
		}
		if !bytes.HasPrefix(noSpaceLine, headerData) || hasErrorPrefix {
			if hasErrorPrefix {
				noSpaceLine = bytes.TrimPrefix(noSpaceLine, headerData)
			}
			writeErr := stream.ErrAccumulator.Write(noSpaceLine)
			if writeErr != nil {
				return model.ChatCompletionStreamResponse{}, writeErr
			}
			emptyMessagesCount++
			if emptyMessagesCount > stream.EmptyMessagesLimit {
				return model.ChatCompletionStreamResponse{}, model.ErrTooManyEmptyStreamMessages
			}

			continue
		}

		noPrefixLine := bytes.TrimPrefix(noSpaceLine, headerData)
		if string(noPrefixLine) == "[DONE]" {
			stream.IsFinished = true
			return model.ChatCompletionStreamResponse{}, io.EOF
		}

		var response model.ChatCompletionStreamResponse
		unmarshalErr := stream.Unmarshaler.Unmarshal(noPrefixLine, &response)
		if unmarshalErr != nil {
			return model.ChatCompletionStreamResponse{}, unmarshalErr
		}

		return response, nil
	}
}

func (stream *BotChatCompletionStreamReader) processLines() (model.BotChatCompletionStreamResponse, error) {
	var (
		emptyMessagesCount uint
		hasErrorPrefix     bool
	)

	for {
		rawLine, readErr := stream.Reader.ReadBytes('\n')
		if readErr != nil || hasErrorPrefix {
			respErr := stream.unmarshalError()
			if respErr != nil {
				return model.BotChatCompletionStreamResponse{}, fmt.Errorf("error, %w", respErr.Error)
			}
			return model.BotChatCompletionStreamResponse{}, readErr
		}

		noSpaceLine := bytes.TrimSpace(rawLine)
		if bytes.HasPrefix(noSpaceLine, errorPrefix) {
			hasErrorPrefix = true
		}
		if !bytes.HasPrefix(noSpaceLine, headerData) || hasErrorPrefix {
			if hasErrorPrefix {
				noSpaceLine = bytes.TrimPrefix(noSpaceLine, headerData)
			}
			writeErr := stream.ErrAccumulator.Write(noSpaceLine)
			if writeErr != nil {
				return model.BotChatCompletionStreamResponse{}, writeErr
			}
			emptyMessagesCount++
			if emptyMessagesCount > stream.EmptyMessagesLimit {
				return model.BotChatCompletionStreamResponse{}, model.ErrTooManyEmptyStreamMessages
			}

			continue
		}

		noPrefixLine := bytes.TrimPrefix(noSpaceLine, headerData)
		if string(noPrefixLine) == "[DONE]" {
			stream.IsFinished = true
			return model.BotChatCompletionStreamResponse{}, io.EOF
		}

		var response model.BotChatCompletionStreamResponse
		unmarshalErr := stream.Unmarshaler.Unmarshal(noPrefixLine, &response)
		if unmarshalErr != nil {
			return model.BotChatCompletionStreamResponse{}, unmarshalErr
		}

		return response, nil
	}
}

func (stream *ChatCompletionStreamReader) unmarshalError() (errResp *model.ErrorResponse) {
	errBytes := stream.ErrAccumulator.Bytes()
	if len(errBytes) == 0 {
		return
	}

	err := stream.Unmarshaler.Unmarshal(errBytes, &errResp)
	if err != nil {
		errResp = nil
	}

	if errResp != nil && errResp.Error != nil {
		if stream.Header() != nil {
			errResp.Error.RequestId = stream.Header().Get(model.ClientRequestHeader)
		}
	}

	return
}

func (stream *ChatCompletionStreamReader) Close() error {
	return stream.Response.Body.Close()
}

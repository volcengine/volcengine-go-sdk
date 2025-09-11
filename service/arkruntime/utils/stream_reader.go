package utils

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"net/http"

	"github.com/volcengine/volcengine-go-sdk/service/arkruntime/model"
	"github.com/volcengine/volcengine-go-sdk/service/arkruntime/model/responses"
)

var (
	headerData  = []byte("data:")
	errorPrefix = []byte(`{"error":`)
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

type ImageGenerationStreamReader struct {
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

type ResponsesStreamReader struct {
	ChatCompletionStreamReader
	Decoder *EventStreamDecoder
}

func (stream *ResponsesStreamReader) Recv() (response *responses.Event, err error) {
	if stream.IsFinished {
		err = io.EOF
		return
	}

	response, err = stream.processLines()
	return
}
func (stream *ImageGenerationStreamReader) Recv() (response model.ImagesStreamResponse, err error) {
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
				return model.ChatCompletionStreamResponse{}, respErr.Error
			}
			return model.ChatCompletionStreamResponse{}, readErr
		}

		// noSpaceLint is trimed with leading space
		noSpaceLine := bytes.TrimSpace(rawLine)
		// trimedLine is trimed with header and followed space (if exists)
		trimedLine := bytes.TrimSpace(bytes.TrimPrefix(noSpaceLine, headerData))
		if bytes.HasPrefix(trimedLine, errorPrefix) {
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

		if string(trimedLine) == "[DONE]" {
			stream.IsFinished = true
			return model.ChatCompletionStreamResponse{}, io.EOF
		}

		var response model.ChatCompletionStreamResponse
		unmarshalErr := stream.Unmarshaler.Unmarshal(trimedLine, &response)
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

		// noSpaceLint is trimed with leading space
		noSpaceLine := bytes.TrimSpace(rawLine)
		// trimedLine is trimed with header and followed space (if exists)
		trimedLine := bytes.TrimSpace(bytes.TrimPrefix(noSpaceLine, headerData))
		if bytes.HasPrefix(trimedLine, errorPrefix) {
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

		if string(trimedLine) == "[DONE]" {
			stream.IsFinished = true
			return model.BotChatCompletionStreamResponse{}, io.EOF
		}

		var response model.BotChatCompletionStreamResponse
		unmarshalErr := stream.Unmarshaler.Unmarshal(trimedLine, &response)
		if unmarshalErr != nil {
			return model.BotChatCompletionStreamResponse{}, unmarshalErr
		}

		return response, nil
	}
}

func (stream *ResponsesStreamReader) processLines() (*responses.Event, error) {
	var (
		emptyMessagesCount uint
	)

	for stream.Decoder.Next() {

		// trimedLine is trimed with header and followed space (if exists)
		if bytes.HasPrefix(stream.Decoder.Event().Data, errorPrefix) {
			writeErr := stream.ErrAccumulator.Write(stream.Decoder.Event().Data)
			if writeErr != nil {
				return nil, writeErr
			}
			emptyMessagesCount++
			if emptyMessagesCount > stream.EmptyMessagesLimit {
				return nil, model.ErrTooManyEmptyStreamMessages
			}
			if respErr := stream.unmarshalError(); respErr != nil {
				return nil, fmt.Errorf("error, %w", respErr.Error)
			}
			continue
		}

		if bytes.HasPrefix(stream.Decoder.Event().Data, []byte("[DONE]")) {
			// In this case we don't break because we still want to iterate through the full stream.
			stream.IsFinished = true
			return nil, io.EOF
		}

		response := &responses.Event{}
		unmarshalErr := stream.Unmarshaler.Unmarshal(stream.Decoder.Event().Data, response)
		if unmarshalErr != nil {
			return nil, unmarshalErr
		}

		return response, nil
	}
	return nil, stream.Decoder.Err()
}

func (stream *ImageGenerationStreamReader) processLines() (model.ImagesStreamResponse, error) {
	var (
		emptyMessagesCount uint
		hasErrorPrefix     bool
	)

	for {
		rawLine, readErr := stream.Reader.ReadBytes('\n')
		if readErr != nil || hasErrorPrefix {
			respErr := stream.unmarshalError()
			if respErr != nil {
				return model.ImagesStreamResponse{}, fmt.Errorf("error, %w", respErr.Error)
			}
			return model.ImagesStreamResponse{}, readErr
		}

		// noSpaceLint is trimed with leading space
		noSpaceLine := bytes.TrimSpace(rawLine)
		// trimedLine is trimed with header and followed space (if exists)
		trimedLine := bytes.TrimSpace(bytes.TrimPrefix(noSpaceLine, headerData))
		if bytes.HasPrefix(trimedLine, errorPrefix) {
			hasErrorPrefix = true
		}
		if !bytes.HasPrefix(noSpaceLine, headerData) || hasErrorPrefix {
			if hasErrorPrefix {
				noSpaceLine = bytes.TrimPrefix(noSpaceLine, headerData)
			}
			writeErr := stream.ErrAccumulator.Write(noSpaceLine)
			if writeErr != nil {
				return model.ImagesStreamResponse{}, writeErr
			}
			emptyMessagesCount++
			if emptyMessagesCount > stream.EmptyMessagesLimit {
				return model.ImagesStreamResponse{}, model.ErrTooManyEmptyStreamMessages
			}

			continue
		}

		if string(trimedLine) == "[DONE]" {
			stream.IsFinished = true
			return model.ImagesStreamResponse{}, io.EOF
		}

		var response model.ImagesStreamResponse
		unmarshalErr := stream.Unmarshaler.Unmarshal(trimedLine, &response)
		if unmarshalErr != nil {
			return model.ImagesStreamResponse{}, unmarshalErr

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

func (stream *BotChatCompletionStreamReader) Close() error {
	// fmt.Printf("%#v\n", stream)
	// fmt.Printf("%#v\n", stream.Response)
	// fmt.Printf("%#v\n", stream.Response.Body)
	return stream.ChatCompletionStreamReader.Response.Body.Close()
}

func (stream *ImageGenerationStreamReader) Close() error {
	return stream.Response.Body.Close()
}

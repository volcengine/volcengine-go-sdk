package model

import (
	"encoding/json"
	"errors"
	"fmt"
)

type APIError struct {
	Code           string  `json:"code,omitempty"`
	Message        string  `json:"message"`
	Param          *string `json:"param,omitempty"`
	Type           string  `json:"type"`
	HTTPStatusCode int     `json:"-"`
	RequestId      string  `json:"request_id"`
}

// RequestError provides information about generic request errors.
type RequestError struct {
	HTTPStatusCode int
	Err            error
	RequestId      string `json:"request_id"`
}

func NewRequestError(httpStatusCode int, rawErr error, requestID string) *RequestError {
	return &RequestError{
		HTTPStatusCode: httpStatusCode,
		Err:            rawErr,
		RequestId:      requestID,
	}
}

type ErrorResponse struct {
	Error *APIError `json:"error,omitempty"`
}

func (e *APIError) Error() string {
	s, _ := json.Marshal(e)
	return fmt.Sprintf("Error code: %d - %s", e.HTTPStatusCode, string(s))
}

func (e *RequestError) Error() string {
	return fmt.Sprintf("RequestError code: %d, err: %v, request_id: %s", e.HTTPStatusCode, e.Err, e.RequestId)
}

func (e *RequestError) Unwrap() error {
	return e.Err
}

var (
	ErrTooManyEmptyStreamMessages       = errors.New("stream has sent too many empty messages")
	ErrChatCompletionInvalidModel       = errors.New("this model is not supported with this method, please use CreateCompletion client method instead") //nolint:lll
	ErrChatCompletionStreamNotSupported = errors.New("streaming is not supported with this method, please use CreateChatCompletionStream")              //nolint:lll
	ErrContentFieldsMisused             = errors.New("can't use both Content and MultiContent properties simultaneously")
	ErrBodyWithoutEndpoint              = errors.New("can't fetch endpoint sts token without endpoint")
	ErrBodyWithoutBot                   = errors.New("can't fetch bot sts token without bot id")
	ErrAKSKNotSupported                 = errors.New("ak&sk authentication is currently not supported for this method, please use api key instead")
)

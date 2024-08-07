package model

import (
	"io"
	"net/http"
	"time"
)

const (
	ClientRequestHeader = "X-Client-Request-Id"

	DefaultMandatoryRefreshTimeout = 10 * 60          // 10 min
	DefaultAdvisoryRefreshTimeout  = 30 * 60          // 30 min
	DefaultStsTimeout              = 7 * 24 * 60 * 60 // 7 days

	InitialRetryDelay = 0.5
	MaxRetryDelay     = 8.0

	ErrorRetryBaseDelay = 500 * time.Millisecond
	ErrorRetryMaxDelay  = 8 * time.Second
)

type Usage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

type Response interface {
	SetHeader(http.Header)
}

type HttpHeader http.Header

func (h *HttpHeader) SetHeader(header http.Header) {
	*h = HttpHeader(header)
}

func (h *HttpHeader) Header() http.Header {
	return http.Header(*h)
}

type RawResponse struct {
	io.ReadCloser

	HttpHeader
}

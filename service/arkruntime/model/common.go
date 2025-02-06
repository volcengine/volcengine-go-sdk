package model

import (
	"io"
	"net/http"
	"time"
)

const (
	ClientRequestHeader = "X-Client-Request-Id"
	RetryAfterHeader    = "Retry-After"

	DefaultMandatoryRefreshTimeout = 10 * 60          // 10 min
	DefaultAdvisoryRefreshTimeout  = 30 * 60          // 30 min
	DefaultStsTimeout              = 7 * 24 * 60 * 60 // 7 days

	InitialRetryDelay = 0.5
	MaxRetryDelay     = 8.0

	ErrorRetryBaseDelay = 500 * time.Millisecond
	ErrorRetryMaxDelay  = 8 * time.Second
)

type PromptTokensDetail struct {
	CachedTokens int `json:"cached_tokens"`
}

type CompletionTokensDetails struct {
	ReasoningTokens int `json:"reasoning_tokens"`
}

type Usage struct {
	PromptTokens            int                     `json:"prompt_tokens"`
	CompletionTokens        int                     `json:"completion_tokens"`
	TotalTokens             int                     `json:"total_tokens"`
	PromptTokensDetails     PromptTokensDetail      `json:"prompt_tokens_details"`
	CompletionTokensDetails CompletionTokensDetails `json:"completion_tokens_details"`
}

type Response interface {
	SetHeader(http.Header)
	GetHeader() http.Header
}

type HttpHeader http.Header

func (h *HttpHeader) SetHeader(header http.Header) {
	*h = HttpHeader(header)
}

func (h *HttpHeader) GetHeader() http.Header {
	return http.Header(*h)
}

func (h *HttpHeader) Header() http.Header {
	return http.Header(*h)
}

type RawResponse struct {
	io.ReadCloser

	HttpHeader
}

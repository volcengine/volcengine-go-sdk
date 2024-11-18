package model

type TruncationStrategyType string

const (
	TruncationStrategyTypeLastHistoryTokens TruncationStrategyType = "last_history_tokens"
)

type TruncationStrategy struct {
	Type              TruncationStrategyType `json:"type"`
	LastHistoryTokens *int                   `json:"last_history_tokens,omitempty"`
}

type CreateContextRequest struct {
	Model              string                   `json:"model"`
	Messages           []*ChatCompletionMessage `json:"messages"`
	TTL                *int                     `json:"ttl,omitempty"`
	TruncationStrategy *TruncationStrategy      `json:"truncation_strategy,omitempty"`
}

type CreateContextResponse struct {
	ID                 string              `json:"id"`
	Model              string              `json:"model"`
	TTL                *int                `json:"ttl,omitempty"`
	TruncationStrategy *TruncationStrategy `json:"truncation_strategy,omitempty"`

	HttpHeader
}

type ContextChatCompletionRequest struct {
	ContextID        string                   `json:"context_id"`
	Model            string                   `json:"model"`
	Messages         []*ChatCompletionMessage `json:"messages"`
	MaxTokens        int                      `json:"max_tokens,omitempty"`
	Temperature      float32                  `json:"temperature,omitempty"`
	TopP             float32                  `json:"top_p,omitempty"`
	Stream           bool                     `json:"stream,omitempty"`
	Stop             []string                 `json:"stop,omitempty"`
	FrequencyPenalty float32                  `json:"frequency_penalty,omitempty"`
	LogitBias        map[string]int           `json:"logit_bias,omitempty"`
	LogProbs         bool                     `json:"logprobs,omitempty"`
	TopLogProbs      int                      `json:"top_logprobs,omitempty"`
	User             string                   `json:"user,omitempty"`
	FunctionCall     interface{}              `json:"function_call,omitempty"`
	Tools            []*Tool                  `json:"tools,omitempty"`
	ToolChoice       interface{}              `json:"tool_choice,omitempty"`
	StreamOptions    *StreamOptions           `json:"stream_options,omitempty"`
	Metadata         map[string]interface{}   `json:"metadata,omitempty"`
}

package model

import (
	"encoding/json"
)

const (
	ChatMessageRoleSystem    = "system"
	ChatMessageRoleUser      = "user"
	ChatMessageRoleAssistant = "assistant"
	ChatMessageRoleTool      = "tool"
)

type ImageURLDetail string

const (
	ImageURLDetailHigh ImageURLDetail = "high"
	ImageURLDetailLow  ImageURLDetail = "low"
	ImageURLDetailAuto ImageURLDetail = "auto"
)

type ChatMessageImageURL struct {
	URL    string         `json:"url,omitempty"`
	Detail ImageURLDetail `json:"detail,omitempty"`
}

type ChatMessageVideoURL struct {
	URL string   `json:"url"`
	FPS *float64 `json:"fps,omitempty"`
}

type ChatCompletionMessageContentPartType string

const (
	ChatCompletionMessageContentPartTypeText     ChatCompletionMessageContentPartType = "text"
	ChatCompletionMessageContentPartTypeImageURL ChatCompletionMessageContentPartType = "image_url"
	ChatCompletionMessageContentPartTypeVideoURL ChatCompletionMessageContentPartType = "video_url"
)

type ChatCompletionMessageContentPart struct {
	Type     ChatCompletionMessageContentPartType `json:"type,omitempty"`
	Text     string                               `json:"text,omitempty"`
	ImageURL *ChatMessageImageURL                 `json:"image_url,omitempty"`
	VideoURL *ChatMessageVideoURL                 `json:"video_url,omitempty"`
}

type ChatCompletionMessageContent struct {
	StringValue *string
	ListValue   []*ChatCompletionMessageContentPart
}

// MarshalJSON implements json.Marshaler.
func (j ChatCompletionMessageContent) MarshalJSON() ([]byte, error) {
	if j.StringValue != nil {
		return json.Marshal(j.StringValue)
	} else if j.ListValue != nil {
		return json.Marshal(j.ListValue)
	} else {
		return json.Marshal(nil)
	}
}

func (j *ChatCompletionMessageContent) UnmarshalJSON(b []byte) error {
	var raw string
	if err := json.Unmarshal(b, &raw); err == nil {
		*j = ChatCompletionMessageContent{}
		j.StringValue = &raw
		return nil
	}

	var raw2 []*ChatCompletionMessageContentPart
	if err := json.Unmarshal(b, &raw2); err == nil {
		*j = ChatCompletionMessageContent{}
		j.ListValue = raw2
		return nil
	}

	return nil
}

type ChatCompletionMessage struct {
	Role             string                        `json:"role"`
	Content          *ChatCompletionMessageContent `json:"content"`
	ReasoningContent *string                       `json:"reasoning_content,omitempty"`
	Name             *string                       `json:"name"`
	FunctionCall     *FunctionCall                 `json:"function_call,omitempty"`
	ToolCalls        []*ToolCall                   `json:"tool_calls,omitempty"`
	ToolCallID       string                        `json:"tool_call_id,omitempty"`
}

type ToolCall struct {
	ID       string       `json:"id"`
	Type     ToolType     `json:"type"`
	Function FunctionCall `json:"function"`
	Index    *int         `json:"index,omitempty"`
}

type FunctionCall struct {
	Name      string `json:"name,omitempty"`
	Arguments string `json:"arguments,omitempty"`
}

type ThinkingType string

const (
	ThinkingTypeEnabled  ThinkingType = "enabled"
	ThinkingTypeDisabled ThinkingType = "disabled"
	ThinkingTypeAuto     ThinkingType = "auto"
)

type Thinking struct {
	Type ThinkingType `json:"type"`
}

type ChatRequest interface {
	json.Marshaler
	WithStream(stream bool) ChatRequest
	IsStream() bool
	GetModel() string
}

// Deprecated: use `CreateChatCompletionRequest` instead.
// ChatCompletionRequest - When making a request using this struct, only non-zero fields take effect.
// This means that if your field value is 0, an empty string (""), false, or
// other zero values, it will not be sent to the server.
// The server will handle these fields according to their default values.
// If you need to specify a zero value, please use CreateChatCompletionRequest.
type ChatCompletionRequest struct {
	Model             string                   `json:"model"`
	Messages          []*ChatCompletionMessage `json:"messages"`
	MaxTokens         int                      `json:"max_tokens,omitempty"`
	Temperature       float32                  `json:"temperature,omitempty"`
	TopP              float32                  `json:"top_p,omitempty"`
	Stream            bool                     `json:"stream,omitempty"`
	Stop              []string                 `json:"stop,omitempty"`
	FrequencyPenalty  float32                  `json:"frequency_penalty,omitempty"`
	LogitBias         map[string]int           `json:"logit_bias,omitempty"`
	LogProbs          bool                     `json:"logprobs,omitempty"`
	TopLogProbs       int                      `json:"top_logprobs,omitempty"`
	User              string                   `json:"user,omitempty"`
	FunctionCall      interface{}              `json:"function_call,omitempty"`
	Tools             []*Tool                  `json:"tools,omitempty"`
	ToolChoice        interface{}              `json:"tool_choice,omitempty"`
	StreamOptions     *StreamOptions           `json:"stream_options,omitempty"`
	PresencePenalty   float32                  `json:"presence_penalty,omitempty"`
	RepetitionPenalty float32                  `json:"repetition_penalty,omitempty"`
	N                 int                      `json:"n,omitempty"`
	ResponseFormat    *ResponseFormat          `json:"response_format,omitempty"`
	ServiceTier       *string                  `json:"service_tier,omitempty"`
}

func (r ChatCompletionRequest) MarshalJSON() ([]byte, error) {
	type Alias ChatCompletionRequest
	return json.Marshal(&struct {
		*Alias
	}{
		Alias: (*Alias)(&r),
	})
}

func (r ChatCompletionRequest) WithStream(stream bool) ChatRequest {
	r.Stream = stream
	return r
}

func (r ChatCompletionRequest) IsStream() bool {
	return r.Stream
}

func (r ChatCompletionRequest) GetModel() string {
	return r.Model
}

// CreateChatCompletionRequest - When making a request using this struct, if your field value is 0,
// an empty string (""), false, or other zero values, it will be sent to
// the server. The server will handle these fields according to the specified values.
type CreateChatCompletionRequest struct {
	Model               string                   `json:"model"`
	Messages            []*ChatCompletionMessage `json:"messages"`
	MaxTokens           *int                     `json:"max_tokens,omitempty"`
	Temperature         *float32                 `json:"temperature,omitempty"`
	TopP                *float32                 `json:"top_p,omitempty"`
	Stream              *bool                    `json:"stream,omitempty"`
	Stop                []string                 `json:"stop,omitempty"`
	FrequencyPenalty    *float32                 `json:"frequency_penalty,omitempty"`
	LogitBias           map[string]int           `json:"logit_bias,omitempty"`
	LogProbs            *bool                    `json:"logprobs,omitempty"`
	TopLogProbs         *int                     `json:"top_logprobs,omitempty"`
	User                *string                  `json:"user,omitempty"`
	FunctionCall        interface{}              `json:"function_call,omitempty"`
	Tools               []*Tool                  `json:"tools,omitempty"`
	ToolChoice          interface{}              `json:"tool_choice,omitempty"`
	StreamOptions       *StreamOptions           `json:"stream_options,omitempty"`
	PresencePenalty     *float32                 `json:"presence_penalty,omitempty"`
	RepetitionPenalty   *float32                 `json:"repetition_penalty,omitempty"`
	N                   *int                     `json:"n,omitempty"`
	ResponseFormat      *ResponseFormat          `json:"response_format,omitempty"`
	ParallelToolCalls   *bool                    `json:"parallel_tool_calls,omitempty"`
	ServiceTier         *string                  `json:"service_tier,omitempty"`
	Thinking            *Thinking                `json:"thinking,omitempty"`
	MaxCompletionTokens *int                     `json:"max_completion_tokens,omitempty"`
}

func (r CreateChatCompletionRequest) MarshalJSON() ([]byte, error) {
	type Alias CreateChatCompletionRequest
	return json.Marshal(&struct {
		*Alias
	}{
		Alias: (*Alias)(&r),
	})
}

func (r CreateChatCompletionRequest) WithStream(stream bool) ChatRequest {
	r.Stream = &stream
	return r
}

func (r CreateChatCompletionRequest) IsStream() bool {
	if r.Stream == nil {
		return false
	}
	return *r.Stream
}

func (r CreateChatCompletionRequest) GetModel() string {
	return r.Model
}

type StreamOptions struct {
	// If set, an additional chunk will be streamed before the data: [DONE] message.
	// The usage field on this chunk shows the token usage statistics for the entire request,
	// and the choices field will always be an empty array.
	// All other chunks will also include a usage field, but with a null value.
	IncludeUsage bool `json:"include_usage,omitempty"`
	// if set, each data chunk will include a `usage` field
	// representing the current cumulative token usage for the entire request.
	ChunkIncludeUsage bool `json:"chunk_include_usage,omitempty"`
}

type ToolType string

const (
	ToolTypeFunction ToolType = "function"
)

type Tool struct {
	Type     ToolType            `json:"type"`
	Function *FunctionDefinition `json:"function,omitempty"`
}

const (
	ToolChoiceStringTypeAuto     = "auto"
	ToolChoiceStringTypeNone     = "none"
	ToolChoiceStringTypeRequired = "required"
)

type ToolChoice struct {
	Type     ToolType           `json:"type"`
	Function ToolChoiceFunction `json:"function,omitempty"`
}

type ToolChoiceFunction struct {
	Name string `json:"name"`
}

type FunctionDefinition struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	// Parameters is an object describing the function.
	// You can pass json.RawMessage to describe the schema,
	// or you can pass in a struct which serializes to the proper JSON schema.
	// The jsonschema package is provided for convenience, but you should
	// consider another specialized library if you require more complex schemas.
	Parameters interface{} `json:"parameters"`
}

// Deprecated: use FunctionDefinition instead.
type FunctionDefine = FunctionDefinition

type TopLogProbs struct {
	Token   string  `json:"token"`
	LogProb float64 `json:"logprob"`
	Bytes   []rune  `json:"bytes,omitempty"`
}

// LogProb represents the probability information for a token.
type LogProb struct {
	Token   string  `json:"token"`
	LogProb float64 `json:"logprob"`
	Bytes   []rune  `json:"bytes,omitempty"` // Omitting the field if it is null
	// TopLogProbs is a list of the most likely tokens and their log probability, at this token position.
	// In rare cases, there may be fewer than the number of requested top_logprobs returned.
	TopLogProbs []*TopLogProbs `json:"top_logprobs"`
}

// LogProbs is the top-level structure containing the log probability information.
type LogProbs struct {
	// Content is a list of message content tokens with log probability information.
	Content []*LogProb `json:"content"`
}

type ResponseFormatType string

type ResponseFormat struct {
	Type       ResponseFormatType                       `json:"type"`
	JSONSchema *ResponseFormatJSONSchemaJSONSchemaParam `json:"json_schema,omitempty"`
	// Deprecated: use `JSONSchema` instead.
	Schema interface{} `json:"schema,omitempty"`
}

type ResponseFormatJSONSchemaJSONSchemaParam struct {
	// The name of the response format. Must be a-z, A-Z, 0-9, or contain underscores
	// and dashes, with a maximum length of 64.
	Name string `json:"name"`
	// A description of what the response format is for, used by the model to determine
	// how to respond in the format.
	Description string `json:"description"`
	// The schema for the response format, described as a JSON Schema object.
	Schema interface{} `json:"schema"`
	// Whether to enable strict schema adherence when generating the output. If set to
	// true, the model will always follow the exact schema defined in the `schema`
	// field. Only a subset of JSON Schema is supported when `strict` is `true`.
	Strict bool `json:"strict"`
}

const (
	ResponseFormatJSONSchema ResponseFormatType = "json_schema"
	ResponseFormatJsonObject ResponseFormatType = "json_object"
	ResponseFormatText       ResponseFormatType = "text"
)

type FinishReason string

const (
	FinishReasonStop          FinishReason = "stop"
	FinishReasonLength        FinishReason = "length"
	FinishReasonFunctionCall  FinishReason = "function_call"
	FinishReasonToolCalls     FinishReason = "tool_calls"
	FinishReasonContentFilter FinishReason = "content_filter"
	FinishReasonNull          FinishReason = "null"
)

func (r FinishReason) MarshalJSON() ([]byte, error) {
	if r == FinishReasonNull || r == "" {
		return []byte("null"), nil
	}
	return []byte(`"` + string(r) + `"`), nil // best effort to not break future API changes
}

type ChatCompletionResponseChoicesElemModerationHitType string

const (
	ChatCompletionResponseChoicesElemModerationHitTypeViolence        ChatCompletionResponseChoicesElemModerationHitType = "violence"
	ChatCompletionResponseChoicesElemModerationHitTypeSevereViolation ChatCompletionResponseChoicesElemModerationHitType = "severe_violation"
)

type ChatCompletionChoice struct {
	Index   int                   `json:"index"`
	Message ChatCompletionMessage `json:"message"`
	// FinishReason
	// stop: API returned complete message,
	// or a message terminated by one of the stop sequences provided via the stop parameter
	// length: Incomplete model output due to max_tokens parameter or token limit
	// function_call: The model decided to call a function
	// content_filter: Omitted content due to a flag from our content filters
	// null: API response still in progress or incomplete
	FinishReason FinishReason `json:"finish_reason"`
	// ModerationHitType
	// The type of content moderation strategy hit.
	// Only after selecting a moderation strategy for the endpoint that supports returning moderation hit types,
	// API will return the corresponding values.
	ModerationHitType *ChatCompletionResponseChoicesElemModerationHitType `json:"moderation_hit_type,omitempty" yaml:"moderation_hit_type,omitempty" mapstructure:"moderation_hit_type,omitempty"`
	LogProbs          *LogProbs                                           `json:"logprobs,omitempty"`
}

// ChatCompletionResponse represents a response structure for chat completion API.
type ChatCompletionResponse struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Created int64  `json:"created"`
	Model   string `json:"model"`
	// mark the request is scale-tier or default, only exists for scale-tier
	ServiceTier string                  `json:"service_tier,omitempty"`
	Choices     []*ChatCompletionChoice `json:"choices"`
	Usage       Usage                   `json:"usage"`

	HttpHeader
}

type ChatCompletionStreamChoiceDelta struct {
	Content          string        `json:"content,omitempty"`
	Role             string        `json:"role,omitempty"`
	ReasoningContent *string       `json:"reasoning_content,omitempty"`
	FunctionCall     *FunctionCall `json:"function_call,omitempty"`
	ToolCalls        []*ToolCall   `json:"tool_calls,omitempty"`
}

type ChatCompletionStreamChoice struct {
	Index             int                                                 `json:"index"`
	Delta             ChatCompletionStreamChoiceDelta                     `json:"delta"`
	LogProbs          *LogProbs                                           `json:"logprobs,omitempty"`
	FinishReason      FinishReason                                        `json:"finish_reason"`
	ModerationHitType *ChatCompletionResponseChoicesElemModerationHitType `json:"moderation_hit_type,omitempty" yaml:"moderation_hit_type,omitempty" mapstructure:"moderation_hit_type,omitempty"`
}

type ChatCompletionStreamResponse struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Created int64  `json:"created"`
	Model   string `json:"model"`
	// mark the request is scale-tier or default, only exists for scale-tier
	ServiceTier string                        `json:"service_tier,omitempty"`
	Choices     []*ChatCompletionStreamChoice `json:"choices"`
	// An optional field that will only be present when you set stream_options: {"include_usage": true} in your request.
	// When present, it contains a null value except for the last chunk which contains the token usage statistics
	// for the entire request.
	Usage *Usage `json:"usage,omitempty"`
}

package model

type BotChatCompletionRequest struct {
	BotId             string                   `json:"bot_id,omitempty"`
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
	Metadata          map[string]interface{}   `json:"metadata,omitempty"`
}

type BotActionUsage struct {
	Name             string  `json:"name"`
	PromptTokens     string  `json:"prompt_tokens,omitempty"`
	CompletionTokens int     `json:"completion_tokens,omitempty"`
	TotalTokens      int     `json:"total_tokens,omitempty"`
	SearchCount      int     `json:"search_count,omitempty"`
	ActionName       *string `json:"action_name,omitempty"`
	Count            *int    `json:"count,omitempty"`
}

type BotModelUsage struct {
	Usage
	Name string `json:"name"`
}

type BotUsage struct {
	ModelUsage  []*BotModelUsage  `json:"model_usage,omitempty"`
	ActionUsage []*BotActionUsage `json:"action_usage,omitempty"`

	ActionDetails []*BotActionDetail `json:"action_details,omitempty"` // 本次请求插件调用详情
}

type BotActionDetail struct {
	Name        string           `json:"name"`                   // 插件分类名称，如 content_plugin（内容插件）等
	Count       int              `json:"count"`                  // 本次请求某插件分类中插件调用次数
	ToolDetails []*BotToolDetail `json:"tool_details,omitempty"` // 某插件分类中插件调用详细信息
}

type BotToolDetail struct {
	Name        string      `json:"name"`         // 具体调用的工具名称
	Input       interface{} `json:"input"`        // 插件输入参数，调用插件的数据结构
	Output      interface{} `json:"output"`       // 插件输出结果，调用插件返回的数据结构
	CreatedAt   int         `json:"created_at"`   // 插件调用开始时间
	CompletedAt int         `json:"completed_at"` // 插件调用结束时间
}

type BotCoverImage struct {
	Url    string `json:"url,omitempty"`
	Width  int    `json:"width,omitempty"`
	Height int    `json:"height,omitempty"`
}

type BotChatResultReference struct {
	Url                string                 `json:"url,omitempty"`
	LogoUrl            string                 `json:"logo_url,omitempty"`
	MobileUrl          string                 `json:"mobile_url,omitempty"`
	SiteName           string                 `json:"site_name,omitempty"`
	Title              string                 `json:"title,omitempty"`
	CoverImage         *BotCoverImage         `json:"cover_image,omitempty"`
	Summary            string                 `json:"summary,omitempty"`
	PublishTime        string                 `json:"publish_time,omitempty"`
	CollectionName     string                 `json:"collection_name,omitempty"`
	Project            string                 `json:"project,omitempty"`
	DocId              string                 `json:"doc_id,omitempty"`
	DocName            string                 `json:"doc_name,omitempty"`
	DocType            string                 `json:"doc_type,omitempty"`
	DocTitle           string                 `json:"doc_title,omitempty"`
	ChunkId            string                 `json:"chunk_id,omitempty"`
	ChunkTitle         string                 `json:"chunk_title,omitempty"`
	PageNums           string                 `json:"page_nums,omitempty"`
	OriginTextTokenLen int                    `json:"origin_text_token_len,omitempty"`
	FileName           string                 `json:"file_name,omitempty"`
	Extra              map[string]interface{} `json:"extra,omitempty"`
}

type BotChatCompletionResponse struct {
	ChatCompletionResponse
	Metadata   map[string]interface{}    `json:"metadata,omitempty"`
	BotUsage   *BotUsage                 `json:"bot_usage,omitempty"`
	References []*BotChatResultReference `json:"references,omitempty"`
}

type BotChatCompletionStreamResponse struct {
	ChatCompletionStreamResponse
	Metadata   map[string]interface{}    `json:"metadata,omitempty"`
	BotUsage   *BotUsage                 `json:"bot_usage,omitempty"`
	References []*BotChatResultReference `json:"references,omitempty"`
}

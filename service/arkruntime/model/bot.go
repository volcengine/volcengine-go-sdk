package model

type BotChatCompletionRequest struct {
	ChatCompletionRequest
	BotId    string                 `json:"bot_id,omitempty"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

type BotActionUsage struct {
	Name             string `json:"name"`
	PromptTokens     string `json:"prompt_tokens,omitempty"`
	CompletionTokens int    `json:"completion_tokens,omitempty"`
	TotalTokens      int    `json:"total_tokens,omitempty"`
	SearchCount      int    `json:"search_count,omitempty"`
}

type BotModelUsage struct {
	Usage
	Name string `json:"name"`
}

type BotUsage struct {
	ModelUsage  []*BotModelUsage  `json:"model_usage,omitempty"`
	ActionUsage []*BotActionUsage `json:"action_usage,omitempty"`
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

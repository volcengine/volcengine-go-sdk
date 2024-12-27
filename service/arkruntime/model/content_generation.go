package model

type ContentGenerationContentItemType string

const (
	ContentGenerationContentItemTypeText  ContentGenerationContentItemType = "text"
	ContentGenerationContentItemTypeImage ContentGenerationContentItemType = "image_url"
)

type CreateContentGenerationContentItem struct {
	Type     ContentGenerationContentItemType `json:"type"`
	Text     string                           `json:"text,omitempty"`
	ImageURL *ImageURL                        `json:"image_url,omitempty"`
}

type ImageURL struct {
	URL string `json:"url"`
}

type CreateContentGenerationTaskRequest struct {
	Model   string                                `json:"model"`
	Content []*CreateContentGenerationContentItem `json:"content"`
}

type CreateContentGenerationTaskResponse struct {
	ID string `json:"id"`

	HttpHeader
}

type GetContentGenerationTaskRequest struct {
	ID string `json:"id"`
}

type DeleteContentGenerationTaskRequest struct {
	ID string `json:"id"`
}

type GetContentGenerationTaskResponse struct {
	ID            string  `json:"id"`
	Model         string  `json:"model"`
	Status        string  `json:"status"`
	FailureReason string  `json:"failure_reason,omitempty"`
	Content       Content `json:"content"`
	Usage         Usage   `json:"usage"`
	CreatedAt     int64   `json:"created_at"`
	UpdatedAt     int64   `json:"updated_at"`

	HttpHeader
}

type Content struct {
	VideoURL string `json:"video_url"`
}

type ContentGenerationUsage struct {
	CompletionTokens int `json:"completion_tokens"`
}

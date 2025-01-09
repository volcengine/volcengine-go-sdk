package model

type ContentGenerationContentItemType string

const (
	ContentGenerationContentItemTypeText  ContentGenerationContentItemType = "text"
	ContentGenerationContentItemTypeImage ContentGenerationContentItemType = "image_url"
)

const (
	StatusSucceeded = "succeeded"
	StatusCancelled = "cancelled"
	StatusFailed    = "failed"
	StatusRunning   = "running"
	StatusQueued    = "queued"
)

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

type GetContentGenerationTaskResponse struct {
	ID            string  `json:"id"`
	Model         string  `json:"model"`
	Status        string  `json:"status"`
	FailureReason *string `json:"failure_reason,omitempty"`
	Content       Content `json:"content"`
	Usage         Usage   `json:"usage"`
	CreatedAt     int64   `json:"created_at"`
	UpdatedAt     int64   `json:"updated_at"`

	HttpHeader
}

type ListContentGenerationTasksRequest struct {
	PageNum  *int                              `json:"page_num,omitempty"`
	PageSize *int                              `json:"page_size,omitempty"`
	Filter   *ListContentGenerationTasksFilter `json:"filter,omitempty"`
}

type DeleteContentGenerationTaskRequest struct {
	ID string `json:"id"`
}

type ListContentGenerationTasksFilter struct {
	Status  *string   `json:"status,omitempty"`
	TaskIDs []*string `json:"task_ids,omitempty"`
	Model   *string   `json:"model,omitempty"`
}

type CreateContentGenerationContentItem struct {
	Type     ContentGenerationContentItemType `json:"type"`
	Text     *string                          `json:"text,omitempty"`
	ImageURL *ImageURL                        `json:"image_url,omitempty"`
}

type ImageURL struct {
	URL string `json:"url"`
}
type Content struct {
	VideoURL string `json:"video_url"`
}

type ContentGenerationUsage struct {
	CompletionTokens int `json:"completion_tokens"`
}

type ListContentGenerationTasksResponse struct {
	Total int64                           `json:"total"`
	Items []ListContentGenerationTaskItem `json:"items"`
	HttpHeader
}

type ListContentGenerationTaskItem struct {
	ID            string  `json:"id"`
	Model         string  `json:"model"`
	Status        string  `json:"status"`
	FailureReason *string `json:"failure_reason,omitempty"`
	Content       Content `json:"content"`
	Usage         Usage   `json:"usage"`
	CreatedAt     int64   `json:"created_at"`
	UpdatedAt     int64   `json:"updated_at"`
}

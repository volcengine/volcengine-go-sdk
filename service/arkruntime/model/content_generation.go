package model

import "encoding/json"

type ContentGenerationContentItemType string

const (
	ContentGenerationContentItemTypeText      ContentGenerationContentItemType = "text"
	ContentGenerationContentItemTypeImage     ContentGenerationContentItemType = "image_url"
	ContentGenerationContentItemTypeDraftTask ContentGenerationContentItemType = "draft_task"
)

const (
	StatusSucceeded = "succeeded"
	StatusCancelled = "cancelled"
	StatusFailed    = "failed"
	StatusRunning   = "running"
	StatusQueued    = "queued"
)

type CreateContentGenerationTaskRequest struct {
	Model                 string                                `json:"model"`
	Content               []*CreateContentGenerationContentItem `json:"content"`
	CallbackUrl           *string                               `json:"callback_url,omitempty"`
	ReturnLastFrame       *bool                                 `json:"return_last_frame,omitempty"`
	ServiceTier           *string                               `json:"service_tier,omitempty"`
	ExecutionExpiresAfter *int64                                `json:"execution_expires_after,omitempty"`
	GenerateAudio         *bool                                 `json:"generate_audio,omitempty"`
	Draft                 *bool                                 `json:"draft,omitempty"`
	CameraFixed           *bool                                 `json:"camera_fixed,omitempty"`
	Watermark             *bool                                 `json:"watermark,omitempty"`
	Seed                  *int64                                `json:"seed,omitempty"`
	Resolution            *string                               `json:"resolution,omitempty"`
	Ratio                 *string                               `json:"ratio,omitempty"`
	Duration              *int64                                `json:"duration,omitempty"`
	Frames                *int64                                `json:"frames,omitempty"`
	ExtraBody             `json:"-"`
}

func (r CreateContentGenerationTaskRequest) MarshalJSON() ([]byte, error) {
	type Alias CreateContentGenerationTaskRequest
	aliasValue := Alias(r)
	data, err := json.Marshal(&aliasValue)
	if err != nil {
		return nil, err
	}

	var jsonObj map[string]interface{}
	if err := json.Unmarshal(data, &jsonObj); err != nil {
		return nil, err
	}

	for k, v := range r.ExtraBody {
		jsonObj[k] = v
	}

	return json.Marshal(jsonObj)
}

type ExtraBody map[string]interface{}

type CreateContentGenerationTaskResponse struct {
	ID string `json:"id"`

	HttpHeader
}

type GetContentGenerationTaskRequest struct {
	ID string `json:"id"`
}

type GetContentGenerationTaskResponse struct {
	ID                    string                  `json:"id"`
	Model                 string                  `json:"model"`
	Status                string                  `json:"status"`
	Error                 *ContentGenerationError `json:"error,omitempty"`
	Content               Content                 `json:"content"`
	Usage                 Usage                   `json:"usage"`
	SubdivisionLevel      *string                 `json:"subdivisionlevel,omitempty"`
	FileFormat            *string                 `json:"fileformat,omitempty"`
	Frames                *int64                  `json:"frames"`
	FramesPerSecond       *int64                  `json:"framespersecond"`
	Resolution            *string                 `json:"resolution,omitempty"`
	Ratio                 *string                 `json:"ratio,omitempty"`
	Duration              *int64                  `json:"duration,omitempty"`
	CreatedAt             int64                   `json:"created_at"`
	UpdatedAt             int64                   `json:"updated_at"`
	Seed                  *int64                  `json:"seed,omitempty"`
	RevisedPrompt         *string                 `json:"revised_prompt,omitempty"`
	ServiceTier           *string                 `json:"service_tier,omitempty"`
	ExecutionExpiresAfter *int64                  `json:"execution_expires_after,omitempty"`
	GenerateAudio         *bool                   `json:"generate_audio,omitempty"`
	Draft                 *bool                   `json:"draft,omitempty"`
	DraftTaskID           *string                 `json:"draft_task_id,omitempty"`

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
	Status      *string   `json:"status,omitempty"`
	TaskIDs     []*string `json:"task_ids,omitempty"`
	Model       *string   `json:"model,omitempty"`
	ServiceTier *string   `json:"service_tier,omitempty"`
}

type CreateContentGenerationContentItem struct {
	Type      ContentGenerationContentItemType `json:"type"`
	Text      *string                          `json:"text,omitempty"`
	ImageURL  *ImageURL                        `json:"image_url,omitempty"`
	Role      *string                          `json:"role,omitempty"`
	DraftTask *DraftTask                       `json:"draft_task,omitempty"`
}

type DraftTask struct {
	ID string `json:"id"`
}

type ImageURL struct {
	URL string `json:"url"`
}
type Content struct {
	VideoURL     string `json:"video_url"`
	LastFrameURL string `json:"last_frame_url"`
	FileURL      string `json:"file_url"`
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
	ID                    string                  `json:"id"`
	Model                 string                  `json:"model"`
	Status                string                  `json:"status"`
	FailureReason         *ContentGenerationError `json:"failure_reason,omitempty"`
	Content               Content                 `json:"content"`
	Usage                 Usage                   `json:"usage"`
	SubdivisionLevel      *string                 `json:"subdivisionlevel,omitempty"`
	FileFormat            *string                 `json:"fileformat,omitempty"`
	Frames                *int64                  `json:"frames"`
	FramesPerSecond       *int64                  `json:"framespersecond"`
	CreatedAt             int64                   `json:"created_at"`
	UpdatedAt             int64                   `json:"updated_at"`
	Seed                  *int64                  `json:"seed,omitempty"`
	RevisedPrompt         *string                 `json:"revised_prompt,omitempty"`
	ServiceTier           *string                 `json:"service_tier,omitempty"`
	ExecutionExpiresAfter *int64                  `json:"execution_expires_after,omitempty"`
	GenerateAudio         *bool                   `json:"generate_audio,omitempty"`
	Draft                 *bool                   `json:"draft,omitempty"`
	DraftTaskID           *string                 `json:"draft_task_id,omitempty"`
}

type ContentGenerationError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

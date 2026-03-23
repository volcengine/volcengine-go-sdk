package model

import "fmt"

const (
	GenerateImagesResponseFormatBase64 = "b64_json"

	GenerateImagesResponseFormatURL = "url"

	GenerateImagesSizeAdaptive = "adaptive"

	OptimizePromptThinkingAuto = "auto"

	OptimizePromptThinkingEnabled = "enabled"

	OptimizePromptThinkingDisabled = "disabled"

	OptimizePromptModeStandard = "standard"

	OptimizePromptModeFast = "fast"

	SequentialImageGenerationAuto = "auto"

	SequentialImageGenerationDisabled = "disabled"

	ImageGenerationStreamEventPartialSucceeded = "image_generation.partial_succeeded"

	ImageGenerationStreamEventPartialFailed = "image_generation.partial_failed"

	ImageGenerationStreamEventCompleted = "image_generation.completed"

	ToolTypeWebSearch ContentGenerationToolType = "web_search"

	OutputFormatJPEG OutputFormat = "jpeg"

	OutputFormatPNG OutputFormat = "png"
)

type OutputFormat string

type ContentGenerationToolType string

type SequentialImageGeneration string

type OptimizePromptThinking string

type OptimizePromptMode string

type GenerateImagesRequest struct {
	Model                            string                            `json:"model"`
	Prompt                           string                            `json:"prompt"`
	Image                            interface{}                       `json:"image,omitempty"`
	ResponseFormat                   *string                           `json:"response_format,omitempty"`
	Seed                             *int64                            `json:"seed,omitempty"`
	GuidanceScale                    *float64                          `json:"guidance_scale,omitempty"`
	Size                             *string                           `json:"size,omitempty"`
	Watermark                        *bool                             `json:"watermark,omitempty"`
	OptimizePrompt                   *bool                             `json:"optimize_prompt,omitempty"`
	OptimizePromptOptions            *OptimizePromptOptions            `json:"optimize_prompt_options,omitempty"`
	SequentialImageGeneration        *SequentialImageGeneration        `json:"sequential_image_generation,omitempty"`
	SequentialImageGenerationOptions *SequentialImageGenerationOptions `json:"sequential_image_generation_options,omitempty"`
	Tools                            []*ContentGenerationTool          `json:"tools,omitempty"`
	OutputFormat                     *OutputFormat                     `json:"output_format,omitempty"`
}

func (req *GenerateImagesRequest) NormalizeImages() error {
	switch v := req.Image.(type) {
	case nil, string, *string, []string, *[]string:
		return nil
	case []interface{}:
		for i, image := range v {
			switch image.(type) {
			case string, *string:
				continue
			case Image:
				continue
			default:
				return fmt.Errorf("unsupported type for Image at index %d: %T", i, image)
			}
		}
	default:
		return fmt.Errorf("unsupported type for Image: %T", v)
	}
	return nil
}

type OptimizePromptOptions struct {
	Thinking *OptimizePromptThinking `json:"thinking,omitempty"`
	Mode     *OptimizePromptMode     `json:"mode,omitempty"`
}

type SequentialImageGenerationOptions struct {
	MaxImages *int `json:"max_images,omitempty"`
}

type Image struct {
	Url     *string `json:"url,omitempty"`
	B64Json *string `json:"b64_json,omitempty"`
	Size    string  `json:"size"`
}

type GenerateImagesUsage struct {
	GeneratedImages int64      `json:"generated_images"`
	OutputTokens    int64      `json:"output_tokens"`
	TotalTokens     int64      `json:"total_tokens"`
	ToolUsage       *ToolUsage `json:"tool_usage,omitempty"`
}

type GenerateImagesError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type ImagesResponse struct {
	Model   string                   `json:"model"`
	Created int64                    `json:"created"`
	Data    []*Image                 `json:"data"`
	Usage   *GenerateImagesUsage     `json:"usage,omitempty"`
	Error   *GenerateImagesError     `json:"error,omitempty"`
	Tools   []*ContentGenerationTool `json:"tools,omitempty"`

	HttpHeader
}

type ImagesStreamResponse struct {
	Type       string                   `json:"type"`
	Model      string                   `json:"model"`
	Created    int64                    `json:"created"`
	ImageIndex int64                    `json:"image_index"`
	Url        *string                  `json:"url,omitempty"`
	B64Json    *string                  `json:"b64_json,omitempty"`
	Size       string                   `json:"size"`
	Usage      *GenerateImagesUsage     `json:"usage,omitempty"`
	Error      *GenerateImagesError     `json:"error,omitempty"`
	Tools      []*ContentGenerationTool `json:"tools,omitempty"`

	HttpHeader
}

type ContentGenerationTool struct {
	Type ContentGenerationToolType `json:"type,required"`
}

type ToolUsage struct {
	WebSearch *int64 `json:"web_search,omitempty"`
}

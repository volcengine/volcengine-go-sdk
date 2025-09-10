package model

import "fmt"

const (
	GenerateImagesResponseFormatBase64 = "b64_json"

	GenerateImagesResponseFormatURL = "url"

	GenerateImagesSizeAdaptive = "adaptive"
)

type SequentialImageGeneration string

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
	SequentialImageGeneration        *SequentialImageGeneration        `json:"sequential_image_generation,omitempty"`
	SequentialImageGenerationOptions *SequentialImageGenerationOptions `json:"sequential_image_generation_options,omitempty"`
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

type SequentialImageGenerationOptions struct {
	MaxImages *int `json:"max_images,omitempty"`
}

type Image struct {
	Url     *string `json:"url,omitempty"`
	B64Json *string `json:"b64_json,omitempty"`
	Size    string  `json:"size"`
}

type GenerateImagesUsage struct {
	GeneratedImages int64 `json:"generated_images"`
	OutputTokens    int64 `json:"output_tokens"`
	TotalTokens     int64 `json:"total_tokens"`
}

type GenerateImagesError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type ImagesResponse struct {
	Model   string               `json:"model"`
	Created int64                `json:"created"`
	Data    []*Image             `json:"data"`
	Usage   *GenerateImagesUsage `json:"usage,omitempty"`
	Error   *GenerateImagesError `json:"error,omitempty"`

	HttpHeader
}

type ImagesStreamResponse struct {
	Type       string               `json:"type"`
	Model      string               `json:"model"`
	Created    int64                `json:"created"`
	ImageIndex int64                `json:"image_index"`
	Url        *string              `json:"url,omitempty"`
	B64Json    *string              `json:"b64_json,omitempty"`
	Size       string               `json:"size"`
	Usage      *GenerateImagesUsage `json:"usage,omitempty"`
	Error      *GenerateImagesError `json:"error,omitempty"`

	HttpHeader
}

package model

const (
	GenerateImagesResponseFormatBase64 = "b64_json"

	GenerateImagesResponseFormatURL = "url"

	GenerateImagesSizeAdaptive = "adaptive"
)

type GenerateImagesRequest struct {
	Model          string   `json:"model"`
	Prompt         string   `json:"prompt"`
	Image          *string  `json:"image,omitempty"`
	ResponseFormat *string  `json:"response_format,omitempty"`
	Seed           *int64   `json:"seed,omitempty"`
	GuidanceScale  *float64 `json:"guidance_scale,omitempty"`
	Size           *string  `json:"size,omitempty"`
	Watermark      *bool    `json:"watermark,omitempty"`
	OptimizePrompt *bool    `json:"optimize_prompt,omitempty"`
}

type GenerateImages4Request struct {
	Model                            string          `json:"model"`
	Prompt                           string          `json:"prompt"`
	Image                            any             `json:"image,omitempty"`
	ResponseFormat                   *string         `json:"response_format,omitempty"`
	GuidanceScale                    *float64        `json:"guidance_scale,omitempty"`
	Size                             *string         `json:"size,omitempty"`
	Watermark                        *bool           `json:"watermark,omitempty"`
	SequentialImageGeneration        *string         `json:"sequential_image_generation,omitempty"`
	SequentialImageGenerationOptions *GenerateOption `json:"sequential_image_generation_options,omitempty"`
}

type GenerateOption struct {
	MaxImages int `json:"max_images"`
}

type Image struct {
	Url     *string `json:"url,omitempty"`
	B64Json *string `json:"b64_json,omitempty"`
}

type GenerateImagesUsage struct {
	GeneratedImages int64 `json:"generated_images"`
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

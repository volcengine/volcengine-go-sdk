package model

const (
	GenerateImagesResponseFormatBase64 = "b64_json"

	GenerateImagesResponseFormatURL = "url"
)

type GenerateImagesRequest struct {
	Model          string   `json:"model"`
	Prompt         string   `json:"prompt"`
	ResponseFormat *string  `json:"response_format,omitempty"`
	Seed           *int64   `json:"seed,omitempty"`
	GuidanceScale  *float64 `json:"guidance_scale,omitempty"`
	Size           *string  `json:"size,omitempty"`
	Watermark      *bool    `json:"watermark,omitempty"`
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

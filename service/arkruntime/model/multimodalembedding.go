package model

type MultiModalEmbeddingInputType string

const (
	MultiModalEmbeddingInputTypeText     MultiModalEmbeddingInputType = "text"
	MultiModalEmbeddingInputTypeImageURL MultiModalEmbeddingInputType = "image_url"
)

type MultimodalEmbeddingImageURL struct {
	URL string `json:"url"`
}

type MultimodalEmbeddingInput struct {
	Type     MultiModalEmbeddingInputType `json:"type"`
	Text     *string                      `json:"text,omitempty"`
	ImageURL *MultimodalEmbeddingImageURL `json:"image_url,omitempty"`
}

// MultiModalEmbeddingRequest is the input to a create embeddings request.
type MultiModalEmbeddingRequest struct {
	Input []MultimodalEmbeddingInput `json:"input"`
	// ID of the model to use. You can use the List models API to see all of your available models,
	// or see our Model overview for descriptions of them.
	Model string `json:"model"`
	// EmbeddingEncodingFormat is the format of the embeddings data.
	// Currently, only "float" and "base64" are supported, however, "base64" is not officially documented.
	// If not specified will use "float".
	EncodingFormat *EmbeddingEncodingFormat `json:"encoding_format,omitempty"`
	// Dimensions Value range: 1024 or 2048.
	// Specifies the dimensionality of the output embedding vector.
	// This parameter is only supported in doubao-embedding-vision-250615 and later versions.
	Dimensions *int `json:"dimensions,omitempty"`
}

type MultimodalEmbedding struct {
	Embedding []float32 `json:"embedding"`
	Object    string    `json:"object"`
}

type MultimodalEmbeddingPromptTokensDetail struct {
	TextTokens  int `json:"text_tokens"`
	ImageTokens int `json:"image_tokens"`
}

type MultimodalEmbeddingUsage struct {
	PromptTokens        int                                   `json:"prompt_tokens"`
	TotalTokens         int                                   `json:"total_tokens"`
	PromptTokensDetails MultimodalEmbeddingPromptTokensDetail `json:"prompt_tokens_details"`
}

type MultimodalEmbeddingResponse struct {
	Id      string                   `json:"id"`
	Model   string                   `json:"model"`
	Created int64                    `json:"created"`
	Object  string                   `json:"object"`
	Data    MultimodalEmbedding      `json:"data"`
	Usage   MultimodalEmbeddingUsage `json:"usage"`

	HttpHeader
}

// MultiModalEmbeddingResponseBase64 is the response from a Create embeddings request with base64 encoding format.
type MultiModalEmbeddingResponseBase64 struct {
	Id      string                   `json:"id"`
	Model   string                   `json:"model"`
	Created int64                    `json:"created"`
	Object  string                   `json:"object"`
	Data    Base64Embedding          `json:"data"`
	Usage   MultimodalEmbeddingUsage `json:"usage"`

	HttpHeader
}

// ToMultiModalEmbeddingResponse converts an embeddingResponseBase64 to an MultimodalEmbeddingResponse.
func (r *MultiModalEmbeddingResponseBase64) ToMultiModalEmbeddingResponse() (MultimodalEmbeddingResponse, error) {
	base64Embedding := r.Data

	embedding, err := base64Embedding.Embedding.Decode()
	if err != nil {
		return MultimodalEmbeddingResponse{}, err
	}

	data := MultimodalEmbedding{
		Object:    base64Embedding.Object,
		Embedding: embedding,
	}

	return MultimodalEmbeddingResponse{
		Object: r.Object,
		Model:  r.Model,
		Data:   data,
		Usage:  r.Usage,
	}, nil
}

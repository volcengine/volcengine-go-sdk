package model

import (
	"bytes"
	"encoding/base64"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"math"

	"github.com/klauspost/compress/zstd"
)

type MultiModalEmbeddingInputType string

const (
	MultiModalEmbeddingInputTypeText     MultiModalEmbeddingInputType = "text"
	MultiModalEmbeddingInputTypeImageURL MultiModalEmbeddingInputType = "image_url"
	MultiModalEmbeddingInputTypeVideoURL MultiModalEmbeddingInputType = "video_url"
)

type SparseEmbeddingInputType string

const (
	SparseEmbeddingInputTypeEnabled  SparseEmbeddingInputType = "enabled"
	SparseEmbeddingInputTypeDisabled SparseEmbeddingInputType = "disabled"
)

type MultiEmbeddingType string

const (
	MultiEmbeddingTypeEnabled  MultiEmbeddingType = "enabled"
	MultiEmbeddingTypeDisabled MultiEmbeddingType = "disabled"
)

type MultiEmbeddingCompression string

const (
	// MultiEmbeddingCompressionZstd requests zstd-compressed, base64-encoded multi_embedding payloads.
	MultiEmbeddingCompressionZstd MultiEmbeddingCompression = "zstd"
	// MultiEmbeddingCompressionBlosc2 is reserved by the wire protocol but unsupported by this Go SDK.
	MultiEmbeddingCompressionBlosc2 MultiEmbeddingCompression = "blosc2"
)

const MultiEmbeddingDimension = 2048

type MultiEmbeddingConfig struct {
	Type        MultiEmbeddingType         `json:"type"`
	Compression *MultiEmbeddingCompression `json:"compression,omitempty"`
}

type MultimodalEmbeddingImageURL struct {
	URL string `json:"url"`
}

type MultimodalEmbeddingVideoURL struct {
	URL            string   `json:"url"`
	FPS            *float64 `json:"fps,omitempty"`
	MaxVideoTokens *int     `json:"max_video_tokens,omitempty"`
	MinFrameTokens *int     `json:"min_frame_tokens,omitempty"`
	MaxFrameTokens *int     `json:"max_frame_tokens,omitempty"`
	MinFrames      *int     `json:"min_frames,omitempty"`
}

type MultimodalEmbeddingInput struct {
	Type     MultiModalEmbeddingInputType `json:"type"`
	Text     *string                      `json:"text,omitempty"`
	ImageURL *MultimodalEmbeddingImageURL `json:"image_url,omitempty"`
	VideoURL *MultimodalEmbeddingVideoURL `json:"video_url,omitempty"`
}

type SparseEmbeddingInput struct {
	Type SparseEmbeddingInputType `json:"type"`
}

type SparseEmbedding struct {
	Index int     `json:"index"`
	Value float64 `json:"value"`
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

	// SparseEmbedding stands for whether to return sparse embedding.
	SparseEmbedding *SparseEmbeddingInput `json:"sparse_embedding,omitempty"`

	// MultiEmbedding controls whether to return token-level multi-vector embeddings.
	MultiEmbedding *MultiEmbeddingConfig `json:"multi_embedding,omitempty"`

	// Instructions stands for the system prompt for the model.
	Instructions *string `json:"instructions,omitempty"`
}

type MultimodalEmbedding struct {
	Embedding       []float32            `json:"embedding"`
	SparseEmbedding *[]SparseEmbedding   `json:"sparse_embedding,omitempty"`
	MultiEmbedding  *MultiEmbeddingValue `json:"multi_embedding,omitempty"`
	Object          string               `json:"object"`
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
		Object:          base64Embedding.Object,
		Embedding:       embedding,
		SparseEmbedding: base64Embedding.SparseEmbedding,
		MultiEmbedding:  base64Embedding.MultiEmbedding,
	}

	return MultimodalEmbeddingResponse{
		Id:         r.Id,
		Created:    r.Created,
		Object:     r.Object,
		Model:      r.Model,
		Data:       data,
		Usage:      r.Usage,
		HttpHeader: r.HttpHeader,
	}, nil
}

type MultiEmbeddingValue struct {
	Vectors    [][]float32
	Compressed string

	compressed bool
}

func (v *MultiEmbeddingValue) UnmarshalJSON(data []byte) error {
	data = bytes.TrimSpace(data)
	if len(data) == 0 {
		return fmt.Errorf("multi_embedding must be a JSON array or string")
	}
	if bytes.Equal(data, []byte("null")) {
		*v = MultiEmbeddingValue{}
		return nil
	}

	switch data[0] {
	case '"':
		var compressed string
		if err := json.Unmarshal(data, &compressed); err != nil {
			return err
		}
		*v = MultiEmbeddingValue{Compressed: compressed, compressed: true}
		return nil
	case '[':
		var vectors [][]float32
		if err := json.Unmarshal(data, &vectors); err != nil {
			return err
		}
		*v = MultiEmbeddingValue{Vectors: vectors}
		return nil
	default:
		return fmt.Errorf("multi_embedding must be a JSON array or string")
	}
}

func (v MultiEmbeddingValue) MarshalJSON() ([]byte, error) {
	if v.compressed || v.Compressed != "" {
		return json.Marshal(v.Compressed)
	}
	return json.Marshal(v.Vectors)
}

func (v *MultiEmbeddingValue) Decode(compression MultiEmbeddingCompression) ([][]float32, error) {
	if v == nil {
		return nil, fmt.Errorf("multi_embedding is nil")
	}
	if v.compressed || v.Compressed != "" {
		return DecodeMultiEmbedding(v.Compressed, compression)
	}
	if compression != "" {
		return nil, fmt.Errorf("multi_embedding is not compressed; got vectors for compression %q", compression)
	}
	return v.Vectors, nil
}

func DecodeMultiEmbedding(raw string, compression MultiEmbeddingCompression) ([][]float32, error) {
	switch compression {
	case MultiEmbeddingCompressionZstd:
		return decodeZstdMultiEmbedding(raw)
	case MultiEmbeddingCompressionBlosc2:
		return nil, ErrMultiEmbeddingBlosc2Unsupported
	case "":
		return nil, fmt.Errorf("multi_embedding.compression is required for compressed multi_embedding")
	default:
		return nil, fmt.Errorf("unsupported multi_embedding.compression %q, allowed values: %q", compression, MultiEmbeddingCompressionZstd)
	}
}

func (c *MultiEmbeddingConfig) Validate() error {
	if c == nil {
		return nil
	}

	switch c.Type {
	case MultiEmbeddingTypeEnabled:
		if c.Compression == nil {
			return nil
		}
		switch *c.Compression {
		case MultiEmbeddingCompressionZstd:
			return nil
		case MultiEmbeddingCompressionBlosc2:
			return ErrMultiEmbeddingBlosc2Unsupported
		default:
			return fmt.Errorf("invalid multi_embedding.compression %q, allowed values: %q", *c.Compression, MultiEmbeddingCompressionZstd)
		}
	case MultiEmbeddingTypeDisabled:
		if c.Compression != nil {
			return fmt.Errorf("multi_embedding.compression must be empty when multi_embedding.type is %q", c.Type)
		}
		return nil
	case "":
		return fmt.Errorf("multi_embedding.type is required")
	default:
		return fmt.Errorf("invalid multi_embedding.type %q, allowed values: %q, %q", c.Type, MultiEmbeddingTypeEnabled, MultiEmbeddingTypeDisabled)
	}
}

func (r MultiModalEmbeddingRequest) Validate() error {
	if err := r.MultiEmbedding.Validate(); err != nil {
		return err
	}
	for i := range r.Input {
		if err := r.Input[i].Validate(); err != nil {
			return fmt.Errorf("input[%d]: %w", i, err)
		}
	}
	return nil
}

func (i MultimodalEmbeddingInput) Validate() error {
	if i.VideoURL != nil {
		return i.VideoURL.Validate()
	}
	return nil
}

func (v *MultimodalEmbeddingVideoURL) Validate() error {
	if v == nil {
		return nil
	}
	if err := validateFloatRange("fps", v.FPS, 0.2, 5); err != nil {
		return err
	}
	if err := validateIntRange("max_video_tokens", v.MaxVideoTokens, 10240, 204800); err != nil {
		return err
	}
	if err := validateIntRange("min_frame_tokens", v.MinFrameTokens, 16, 128); err != nil {
		return err
	}
	if err := validateIntRange("max_frame_tokens", v.MaxFrameTokens, 128, 640); err != nil {
		return err
	}
	if err := validateIntRange("min_frames", v.MinFrames, 5, 16); err != nil {
		return err
	}
	if v.MinFrameTokens != nil && v.MaxFrameTokens != nil && *v.MaxFrameTokens < *v.MinFrameTokens {
		return fmt.Errorf("max_frame_tokens must be greater than or equal to min_frame_tokens")
	}
	return nil
}

func validateFloatRange(name string, value *float64, min, max float64) error {
	if value == nil {
		return nil
	}
	if *value < min || *value > max {
		return fmt.Errorf("%s must be in range [%g, %g]", name, min, max)
	}
	return nil
}

func validateIntRange(name string, value *int, min, max int) error {
	if value == nil {
		return nil
	}
	if *value < min || *value > max {
		return fmt.Errorf("%s must be in range [%d, %d]", name, min, max)
	}
	return nil
}

func decodeZstdMultiEmbedding(raw string) ([][]float32, error) {
	compressed, err := base64.StdEncoding.DecodeString(raw)
	if err != nil {
		return nil, fmt.Errorf("decode multi_embedding base64: %w", err)
	}

	decoder, err := zstd.NewReader(nil)
	if err != nil {
		return nil, fmt.Errorf("create zstd decoder: %w", err)
	}
	defer decoder.Close()

	decompressed, err := decoder.DecodeAll(compressed, nil)
	if err != nil {
		return nil, fmt.Errorf("decompress multi_embedding zstd: %w", err)
	}
	return decodeFP16Matrix(decompressed)
}

func decodeFP16Matrix(data []byte) ([][]float32, error) {
	if len(data)%2 != 0 {
		return nil, fmt.Errorf("decoded multi_embedding fp16 byte length must be divisible by 2, got %d", len(data))
	}

	elements := len(data) / 2
	if elements%MultiEmbeddingDimension != 0 {
		return nil, fmt.Errorf("decoded multi_embedding fp16 element count %d is not divisible by %d", elements, MultiEmbeddingDimension)
	}

	values := make([]float32, elements)
	for i := 0; i < elements; i++ {
		values[i] = float16ToFloat32(binary.LittleEndian.Uint16(data[i*2 : (i+1)*2]))
	}

	rows := elements / MultiEmbeddingDimension
	vectors := make([][]float32, rows)
	for i := 0; i < rows; i++ {
		start := i * MultiEmbeddingDimension
		vectors[i] = values[start : start+MultiEmbeddingDimension]
	}
	return vectors, nil
}

func float16ToFloat32(h uint16) float32 {
	sign := uint32(h&0x8000) << 16
	exponent := uint32(h>>10) & 0x1f
	mantissa := uint32(h & 0x03ff)

	switch exponent {
	case 0:
		if mantissa == 0 {
			return math.Float32frombits(sign)
		}
		exponent32 := int32(127 - 14)
		for mantissa&0x0400 == 0 {
			mantissa <<= 1
			exponent32--
		}
		mantissa &= 0x03ff
		return math.Float32frombits(sign | uint32(exponent32)<<23 | mantissa<<13)
	case 0x1f:
		return math.Float32frombits(sign | 0x7f800000 | mantissa<<13)
	default:
		return math.Float32frombits(sign | (exponent+112)<<23 | mantissa<<13)
	}
}

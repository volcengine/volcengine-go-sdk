package model

import (
	"encoding/base64"
	"encoding/binary"
	"encoding/json"
	"errors"
	"math"
	"strings"
	"testing"

	"github.com/klauspost/compress/zstd"
)

func TestMultiModalEmbeddingRequestMarshalAndValidate(t *testing.T) {
	text := "hello"
	defaultReq := MultiModalEmbeddingRequest{
		Model: "doubao-embedding-vision-251215",
		Input: []MultimodalEmbeddingInput{{
			Type: MultiModalEmbeddingInputTypeText,
			Text: &text,
		}},
	}
	payload, err := json.Marshal(defaultReq)
	if err != nil {
		t.Fatalf("json.Marshal() error = %v", err)
	}
	body := string(payload)
	for _, field := range []string{"multi_embedding", "dimensions", "fps", "max_video_tokens", "min_frame_tokens", "max_frame_tokens", "min_frames"} {
		if strings.Contains(body, field) {
			t.Fatalf("default request should omit %s: %s", field, body)
		}
	}

	compression := MultiEmbeddingCompressionZstd
	fps := 0.2
	maxVideoTokens := 10240
	minFrameTokens := 16
	maxFrameTokens := 128
	minFrames := 5
	videoReq := MultiModalEmbeddingRequest{
		Model: "doubao-embedding-vision-251215",
		Input: []MultimodalEmbeddingInput{{
			Type: MultiModalEmbeddingInputTypeVideoURL,
			VideoURL: &MultimodalEmbeddingVideoURL{
				URL:            "https://example.com/video.mp4",
				FPS:            &fps,
				MaxVideoTokens: &maxVideoTokens,
				MinFrameTokens: &minFrameTokens,
				MaxFrameTokens: &maxFrameTokens,
				MinFrames:      &minFrames,
			},
		}},
		MultiEmbedding: &MultiEmbeddingConfig{
			Type:        MultiEmbeddingTypeEnabled,
			Compression: &compression,
		},
	}
	if err := videoReq.Validate(); err != nil {
		t.Fatalf("Validate() error = %v", err)
	}
	payload, err = json.Marshal(videoReq)
	if err != nil {
		t.Fatalf("json.Marshal() error = %v", err)
	}
	body = string(payload)
	for _, want := range []string{
		`"multi_embedding":{"type":"enabled","compression":"zstd"}`,
		`"fps":0.2`,
		`"max_video_tokens":10240`,
		`"min_frame_tokens":16`,
		`"max_frame_tokens":128`,
		`"min_frames":5`,
	} {
		if !strings.Contains(body, want) {
			t.Fatalf("request JSON missing %s: %s", want, body)
		}
	}
}

func TestMultiModalEmbeddingRequestValidateRejectsInvalidValues(t *testing.T) {
	if err := (MultiModalEmbeddingRequest{MultiEmbedding: &MultiEmbeddingConfig{}}).Validate(); err == nil || !strings.Contains(err.Error(), "multi_embedding.type is required") {
		t.Fatalf("empty multi_embedding Validate() error = %v", err)
	}

	unknownCompression := MultiEmbeddingCompression("gzip")
	if err := (MultiModalEmbeddingRequest{MultiEmbedding: &MultiEmbeddingConfig{Type: MultiEmbeddingTypeEnabled, Compression: &unknownCompression}}).Validate(); err == nil || !strings.Contains(err.Error(), "invalid multi_embedding.compression") {
		t.Fatalf("unknown compression Validate() error = %v", err)
	}

	blosc2 := MultiEmbeddingCompressionBlosc2
	if err := (MultiModalEmbeddingRequest{MultiEmbedding: &MultiEmbeddingConfig{Type: MultiEmbeddingTypeEnabled, Compression: &blosc2}}).Validate(); !errors.Is(err, ErrMultiEmbeddingBlosc2Unsupported) {
		t.Fatalf("blosc2 Validate() error = %v", err)
	}

	zstdCompression := MultiEmbeddingCompressionZstd
	if err := (MultiModalEmbeddingRequest{MultiEmbedding: &MultiEmbeddingConfig{Type: MultiEmbeddingTypeDisabled, Compression: &zstdCompression}}).Validate(); err == nil || !strings.Contains(err.Error(), "must be empty") {
		t.Fatalf("disabled with compression Validate() error = %v", err)
	}

	fps := 0.1
	videoReq := MultiModalEmbeddingRequest{Input: []MultimodalEmbeddingInput{{
		Type: MultiModalEmbeddingInputTypeVideoURL,
		VideoURL: &MultimodalEmbeddingVideoURL{
			URL: "https://example.com/video.mp4",
			FPS: &fps,
		},
	}}}
	if err := videoReq.Validate(); err == nil || !strings.Contains(err.Error(), "fps must be in range") {
		t.Fatalf("bad fps Validate() error = %v", err)
	}
}

func TestMultiEmbeddingValueUnmarshalArrayAndString(t *testing.T) {
	var vectors MultiEmbeddingValue
	if err := json.Unmarshal([]byte(`[[1,2],[3,4]]`), &vectors); err != nil {
		t.Fatalf("json.Unmarshal(array) error = %v", err)
	}
	if len(vectors.Vectors) != 2 || vectors.Compressed != "" || vectors.Vectors[1][0] != 3 {
		t.Fatalf("unexpected array value: %+v", vectors)
	}

	var compressed MultiEmbeddingValue
	if err := json.Unmarshal([]byte(`"raw-zstd"`), &compressed); err != nil {
		t.Fatalf("json.Unmarshal(string) error = %v", err)
	}
	if compressed.Compressed != "raw-zstd" || len(compressed.Vectors) != 0 {
		t.Fatalf("unexpected compressed value: %+v", compressed)
	}
}

func TestMultiModalBase64ConversionPreservesMultiEmbedding(t *testing.T) {
	embeddingBytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(embeddingBytes, math.Float32bits(1.25))
	body := `{
		"id":"resp-id",
		"created":123,
		"object":"list",
		"model":"doubao-embedding-vision-251215",
		"data":{"object":"embedding","embedding":"` + base64.StdEncoding.EncodeToString(embeddingBytes) + `","multi_embedding":"raw-zstd"},
		"usage":{"prompt_tokens":1,"total_tokens":1,"prompt_tokens_details":{"text_tokens":1,"image_tokens":0}}
	}`

	var base64Resp MultiModalEmbeddingResponseBase64
	if err := json.Unmarshal([]byte(body), &base64Resp); err != nil {
		t.Fatalf("json.Unmarshal() error = %v", err)
	}
	resp, err := base64Resp.ToMultiModalEmbeddingResponse()
	if err != nil {
		t.Fatalf("ToMultiModalEmbeddingResponse() error = %v", err)
	}
	if resp.Id != "resp-id" || resp.Created != 123 {
		t.Fatalf("response metadata lost: %+v", resp)
	}
	if len(resp.Data.Embedding) != 1 || resp.Data.Embedding[0] != 1.25 {
		t.Fatalf("unexpected embedding: %+v", resp.Data.Embedding)
	}
	if resp.Data.MultiEmbedding == nil || resp.Data.MultiEmbedding.Compressed != "raw-zstd" {
		t.Fatalf("multi_embedding not preserved: %+v", resp.Data.MultiEmbedding)
	}
}

func TestDecodeMultiEmbeddingZstd(t *testing.T) {
	payload := make([]byte, MultiEmbeddingDimension*2)
	binary.LittleEndian.PutUint16(payload[0:2], 0x3c00) // 1
	binary.LittleEndian.PutUint16(payload[2:4], 0xc000) // -2
	binary.LittleEndian.PutUint16(payload[4:6], 0x3800) // 0.5

	matrix, err := DecodeMultiEmbedding(zstdBase64(t, payload), MultiEmbeddingCompressionZstd)
	if err != nil {
		t.Fatalf("DecodeMultiEmbedding() error = %v", err)
	}
	if len(matrix) != 1 || len(matrix[0]) != MultiEmbeddingDimension {
		t.Fatalf("unexpected shape: rows=%d dims=%d", len(matrix), len(matrix[0]))
	}
	if matrix[0][0] != 1 || matrix[0][1] != -2 || matrix[0][2] != 0.5 {
		t.Fatalf("unexpected decoded prefix: %v", matrix[0][:3])
	}
}

func TestDecodeMultiEmbeddingErrors(t *testing.T) {
	if _, err := DecodeMultiEmbedding("not-base64", MultiEmbeddingCompressionZstd); err == nil || !strings.Contains(err.Error(), "base64") {
		t.Fatalf("invalid base64 error = %v", err)
	}

	if _, err := DecodeMultiEmbedding(base64.StdEncoding.EncodeToString([]byte("not-zstd")), MultiEmbeddingCompressionZstd); err == nil || !strings.Contains(err.Error(), "zstd") {
		t.Fatalf("invalid zstd error = %v", err)
	}

	if _, err := DecodeMultiEmbedding(zstdBase64(t, []byte{0}), MultiEmbeddingCompressionZstd); err == nil || !strings.Contains(err.Error(), "divisible by 2") {
		t.Fatalf("odd fp16 bytes error = %v", err)
	}

	if _, err := DecodeMultiEmbedding(zstdBase64(t, []byte{0, 0}), MultiEmbeddingCompressionZstd); err == nil || !strings.Contains(err.Error(), "not divisible") {
		t.Fatalf("bad shape error = %v", err)
	}

	if _, err := DecodeMultiEmbedding("ignored", MultiEmbeddingCompressionBlosc2); !errors.Is(err, ErrMultiEmbeddingBlosc2Unsupported) {
		t.Fatalf("blosc2 error = %v", err)
	}
}

func zstdBase64(t *testing.T, payload []byte) string {
	t.Helper()
	encoder, err := zstd.NewWriter(nil)
	if err != nil {
		t.Fatalf("zstd.NewWriter() error = %v", err)
	}
	defer encoder.Close()
	compressed := encoder.EncodeAll(payload, nil)
	return base64.StdEncoding.EncodeToString(compressed)
}

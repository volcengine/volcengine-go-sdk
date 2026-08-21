package model

import (
	"encoding/json"
	"testing"
)

func TestGenerateImagesUsageInputImagesJSON(t *testing.T) {
	t.Run("omitted response field", func(t *testing.T) {
		var usage GenerateImagesUsage
		if err := json.Unmarshal([]byte(`{"generated_images":1}`), &usage); err != nil {
			t.Fatalf("json.Unmarshal() error = %v", err)
		}
		if usage.InputImages != nil {
			t.Fatalf("InputImages = %v, want nil", *usage.InputImages)
		}
	})

	t.Run("explicit zero response field", func(t *testing.T) {
		var usage GenerateImagesUsage
		if err := json.Unmarshal([]byte(`{"input_images":0}`), &usage); err != nil {
			t.Fatalf("json.Unmarshal() error = %v", err)
		}
		if usage.InputImages == nil {
			t.Fatal("InputImages = nil, want pointer to 0")
		}
		if got := *usage.InputImages; got != 0 {
			t.Fatalf("InputImages = %d, want 0", got)
		}
	})

	t.Run("nil field is omitted", func(t *testing.T) {
		payload, err := json.Marshal(GenerateImagesUsage{})
		if err != nil {
			t.Fatalf("json.Marshal() error = %v", err)
		}

		var body map[string]interface{}
		if err := json.Unmarshal(payload, &body); err != nil {
			t.Fatalf("json.Unmarshal() error = %v", err)
		}
		if _, ok := body["input_images"]; ok {
			t.Fatalf("input_images should be omitted: %s", payload)
		}
	})
}

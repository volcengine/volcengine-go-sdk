package model

import (
	"encoding/json"
	"testing"
)

func TestCreateContentGenerationTaskRequestOutputFormatJSON(t *testing.T) {
	t.Run("set", func(t *testing.T) {
		outputFormat := "mp4"
		request := CreateContentGenerationTaskRequest{
			OutputFormat: &outputFormat,
		}

		payload, err := json.Marshal(request)
		if err != nil {
			t.Fatalf("json.Marshal() error = %v", err)
		}

		var body map[string]interface{}
		if err := json.Unmarshal(payload, &body); err != nil {
			t.Fatalf("json.Unmarshal() error = %v", err)
		}
		if got := body["output_format"]; got != outputFormat {
			t.Fatalf("output_format = %v, want %q", got, outputFormat)
		}
	})

	t.Run("unset", func(t *testing.T) {
		payload, err := json.Marshal(CreateContentGenerationTaskRequest{})
		if err != nil {
			t.Fatalf("json.Marshal() error = %v", err)
		}

		var body map[string]interface{}
		if err := json.Unmarshal(payload, &body); err != nil {
			t.Fatalf("json.Unmarshal() error = %v", err)
		}
		if _, ok := body["output_format"]; ok {
			t.Fatalf("output_format should be omitted: %s", payload)
		}
	})
}

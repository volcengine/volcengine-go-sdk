package model

import (
	"encoding/json"
	"reflect"
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

func TestCreateContentGenerationTaskRequestOmniReferenceTaskTypeJSON(t *testing.T) {
	t.Run("set", func(t *testing.T) {
		omniReferenceTaskType := "first_last_frame"
		request := CreateContentGenerationTaskRequest{
			OmniReferenceTaskType: &omniReferenceTaskType,
		}

		payload, err := json.Marshal(request)
		if err != nil {
			t.Fatalf("json.Marshal() error = %v", err)
		}

		var body map[string]interface{}
		if err := json.Unmarshal(payload, &body); err != nil {
			t.Fatalf("json.Unmarshal() error = %v", err)
		}
		if got := body["omni_reference_task_type"]; got != omniReferenceTaskType {
			t.Fatalf("omni_reference_task_type = %v, want %q", got, omniReferenceTaskType)
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
		if _, ok := body["omni_reference_task_type"]; ok {
			t.Fatalf("omni_reference_task_type should be omitted: %s", payload)
		}
	})
}

func TestGetContentGenerationTaskResponseOutputFormatJSON(t *testing.T) {
	var response GetContentGenerationTaskResponse
	if err := json.Unmarshal([]byte(`{"output_format":"mp4"}`), &response); err != nil {
		t.Fatalf("json.Unmarshal() error = %v", err)
	}
	if response.OutputFormat == nil {
		t.Fatal("OutputFormat = nil, want pointer to \"mp4\"")
	}
	if got, want := *response.OutputFormat, "mp4"; got != want {
		t.Fatalf("OutputFormat = %q, want %q", got, want)
	}
}

func TestGetContentGenerationTaskResponseDurationJSON(t *testing.T) {
	tests := []struct {
		name    string
		payload string
		want    float64
	}{
		{name: "fractional", payload: `{"duration":5.5}`, want: 5.5},
		{name: "integer", payload: `{"duration":5}`, want: 5.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var response GetContentGenerationTaskResponse
			if err := json.Unmarshal([]byte(tt.payload), &response); err != nil {
				t.Fatalf("json.Unmarshal() error = %v", err)
			}
			if response.Duration == nil {
				t.Fatal("Duration = nil, want non-nil pointer")
			}

			value := reflect.ValueOf(*response.Duration)
			if got, want := value.Kind(), reflect.Float64; got != want {
				t.Fatalf("Duration kind = %v, want %v", got, want)
			}
			if got := value.Float(); got != tt.want {
				t.Fatalf("Duration = %v, want %v", got, tt.want)
			}
		})
	}
}

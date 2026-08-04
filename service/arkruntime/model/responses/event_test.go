package responses

import (
	"encoding/json"
	"testing"
)

func TestEventUnmarshalReasoningText(t *testing.T) {
	t.Run("delta", func(t *testing.T) {
		var event Event
		raw := []byte(`{"type":"response.reasoning_text.delta","content_index":0,"delta":"reasoning","item_id":"item_1","output_index":1,"sequence_number":4}`)

		if err := json.Unmarshal(raw, &event); err != nil {
			t.Fatalf("unmarshal reasoning text delta event: %v", err)
		}

		delta := event.GetReasoningRawTextDelta()
		if delta == nil {
			t.Fatal("reasoning text delta event was not decoded")
		}
		if delta.GetType() != EventType_response_reasoning_text_delta {
			t.Fatalf("unexpected event type: %v", delta.GetType())
		}
		if delta.GetDelta() != "reasoning" {
			t.Fatalf("unexpected delta: %q", delta.GetDelta())
		}
	})

	t.Run("done", func(t *testing.T) {
		var event Event
		raw := []byte(`{"type":"response.reasoning_text.done","content_index":0,"item_id":"item_1","output_index":1,"sequence_number":5,"text":"reasoning"}`)

		if err := json.Unmarshal(raw, &event); err != nil {
			t.Fatalf("unmarshal reasoning text done event: %v", err)
		}

		done := event.GetReasoningRawTextDone()
		if done == nil {
			t.Fatal("reasoning text done event was not decoded")
		}
		if done.GetType() != EventType_response_reasoning_text_done {
			t.Fatalf("unexpected event type: %v", done.GetType())
		}
		if done.GetText() != "reasoning" {
			t.Fatalf("unexpected text: %q", done.GetText())
		}
	})
}

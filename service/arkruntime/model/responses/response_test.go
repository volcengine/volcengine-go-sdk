package responses

import (
	"encoding/json"
	"testing"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func TestResponseObjectUnmarshalReasoningContent(t *testing.T) {
	var response ResponseObject
	raw := []byte(`{
		"id":"resp_1",
		"object":"response",
		"status":"completed",
		"output":[{
			"id":"rs_1",
			"type":"reasoning",
			"status":"completed",
			"content":[{"type":"reasoning_text","text":"reasoning"}]
		}]
	}`)

	if err := json.Unmarshal(raw, &response); err != nil {
		t.Fatalf("unmarshal response with reasoning content: %v", err)
	}
	if len(response.GetOutput()) != 1 {
		t.Fatalf("unexpected output item count: %d", len(response.GetOutput()))
	}

	reasoning := response.GetOutput()[0].GetReasoning()
	if reasoning == nil {
		t.Fatal("reasoning output item was not decoded")
	}
	if len(reasoning.GetContent()) != 1 {
		t.Fatalf("unexpected reasoning content count: %d", len(reasoning.GetContent()))
	}

	content := reasoning.GetContent()[0].GetText()
	if content == nil {
		t.Fatal("reasoning text content was not decoded")
	}
	if content.GetType() != ContentItemType_reasoning_text {
		t.Fatalf("unexpected reasoning content type: %v", content.GetType())
	}
	if content.GetText() != "reasoning" {
		t.Fatalf("unexpected reasoning content text: %q", content.GetText())
	}

	contentField := reasoning.ProtoReflect().Descriptor().Fields().ByNumber(protoreflect.FieldNumber(5))
	if contentField == nil {
		t.Fatal("ItemReasoning.content is missing from the protobuf descriptor")
	}
	if contentField.Name() != "content" || contentField.Message().Name() != "OutputContentItem" {
		t.Fatalf("unexpected ItemReasoning field 5 descriptor: %v", contentField)
	}

	wire, err := proto.Marshal(reasoning)
	if err != nil {
		t.Fatalf("marshal reasoning protobuf: %v", err)
	}
	var wireDecoded ItemReasoning
	if err := proto.Unmarshal(wire, &wireDecoded); err != nil {
		t.Fatalf("unmarshal reasoning protobuf: %v", err)
	}
	if len(wireDecoded.GetContent()) != 1 || wireDecoded.GetContent()[0].GetText().GetText() != "reasoning" {
		t.Fatalf("reasoning content was not preserved by protobuf round trip: %v", wireDecoded.GetContent())
	}

	protoJSON, err := protojson.Marshal(reasoning)
	if err != nil {
		t.Fatalf("marshal reasoning protojson: %v", err)
	}
	var protoJSONDecoded ItemReasoning
	if err := protojson.Unmarshal(protoJSON, &protoJSONDecoded); err != nil {
		t.Fatalf("unmarshal reasoning protojson: %v", err)
	}
	if len(protoJSONDecoded.GetContent()) != 1 || protoJSONDecoded.GetContent()[0].GetText().GetText() != "reasoning" {
		t.Fatalf("reasoning content was not preserved by protojson round trip: %v", protoJSONDecoded.GetContent())
	}
}

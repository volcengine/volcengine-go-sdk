package responses

import (
	"bytes"
	"encoding/json"
	"net/url"
	"reflect"
	"strings"

	"github.com/volcengine/volcengine-go-sdk/service/arkruntime/pkg/apiquery"
)

func unmarshal(message []byte, val interface{}) error {
	return json.Unmarshal(message, val)
}

// UnmarshalJSON ...
func (r *ResponsesRequest) UnmarshalJSON(bytes []byte) error {
	type ResponsesRequestAlias ResponsesRequest
	var alias ResponsesRequestAlias
	if err := unmarshal(bytes, &alias); err != nil {
		return err
	}
	*r = ResponsesRequest(alias) // nolint
	return nil
}

// UnmarshalJSON ...
func (r *ResponsesInput) UnmarshalJSON(bytes []byte) error {
	var err error
	oneof1 := ResponsesInput_StringValue{}
	if err = unmarshal(bytes, &oneof1.StringValue); err == nil {
		r.Union = &oneof1
		return nil
	}
	oneof2 := ResponsesInput_ListValue{
		ListValue: &InputItemList{},
	}
	if err = unmarshal(bytes, &oneof2.ListValue.ListValue); err == nil {
		r.Union = &oneof2
		return nil
	}
	return err
}

// MarshalJSON ...
func (r *ResponsesInput) MarshalJSON() ([]byte, error) {
	if v := r.GetListValue(); v != nil {
		return json.Marshal(v.ListValue)
	}
	v := r.GetStringValue()
	return json.Marshal(v)
}

// UnmarshalJSON ...
func (r *InputItem) UnmarshalJSON(bytes []byte) error {
	var err error
	var typeOnly struct {
		Type *ItemType_Enum `json:"type,omitempty"`
	}
	if err = json.Unmarshal(bytes, &typeOnly); err != nil {
		return err
	}
	unmarshalMessageFunc := func(bytes []byte, i *InputItem) (err error) { // nolint
		// Try unmarshal to easy message
		easyMessage := InputItem_EasyMessage{}
		if err = unmarshal(bytes, &easyMessage.EasyMessage); err == nil {
			r.Union = &easyMessage
			return nil
		}
		var roleOnly struct {
			Role MessageRole_Enum `json:"role"`
		}
		if err = json.Unmarshal(bytes, &roleOnly); err != nil {
			return err
		}
		if roleOnly.Role == MessageRole_assistant {
			outputMessage := InputItem_OutputMessage{}
			if err = unmarshal(bytes, &outputMessage.OutputMessage); err == nil {
				r.Union = &outputMessage
				return nil
			}
		} else {
			inputMessage := InputItem_InputMessage{}
			if err = unmarshal(bytes, &inputMessage.InputMessage); err == nil {
				r.Union = &inputMessage
				return nil
			}
		}
		return
	}
	switch {
	case typeOnly.Type == nil:
		// TODO: item reference
		// Try unmarshal to message
		err = unmarshalMessageFunc(bytes, r)
		if err == nil {
			return nil
		}
	case *typeOnly.Type == ItemType_message:
		// Try unmarshal to message
		err = unmarshalMessageFunc(bytes, r)
		if err == nil {
			return nil
		}
	case *typeOnly.Type == ItemType_function_call:
		fc := InputItem_FunctionToolCall{}
		if err = unmarshal(bytes, &fc.FunctionToolCall); err == nil {
			r.Union = &fc
			return nil
		}
	case *typeOnly.Type == ItemType_function_call_output:
		fco := InputItem_FunctionToolCallOutput{}
		if err = unmarshal(bytes, &fco.FunctionToolCallOutput); err == nil {
			r.Union = &fco
			return nil
		}
	case *typeOnly.Type == ItemType_reasoning:
		reasoning := InputItem_Reasoning{}
		if err = unmarshal(bytes, &reasoning.Reasoning); err == nil {
			r.Union = &reasoning
			return nil
		}
	case *typeOnly.Type == ItemType_image_process:
		imageProcess := InputItem_ImageProcess{}
		if err = unmarshal(bytes, &imageProcess.ImageProcess); err == nil {
			r.Union = &imageProcess
			return nil
		}
	}
	return err
}

// MarshalJSON ...
func (r *InputItem) MarshalJSON() ([]byte, error) {
	if v := r.GetEasyMessage(); v != nil {
		return json.Marshal(v)
	}
	if v := r.GetInputMessage(); v != nil {
		return json.Marshal(v)
	}
	if v := r.GetOutputMessage(); v != nil {
		return json.Marshal(v)
	}
	if v := r.GetFunctionToolCall(); v != nil {
		return json.Marshal(v)
	}
	if v := r.GetFunctionToolCallOutput(); v != nil {
		return json.Marshal(v)
	}
	if v := r.GetReasoning(); v != nil {
		return json.Marshal(v)
	}
	if v := r.GetImageProcess(); v != nil {
		return json.Marshal(v)
	}
	return json.Marshal(nil)
}

// MarshalJSON ...
func (x *ResponseImageProcessArgs) MarshalJSON() ([]byte, error) {
	if v := x.GetPointArgs(); v != nil {
		return json.Marshal(v)
	}
	if v := x.GetGroundingArgs(); v != nil {
		return json.Marshal(v)
	}
	if v := x.GetRotateArgs(); v != nil {
		return json.Marshal(v)
	}
	if v := x.GetZoomArgs(); v != nil {
		return json.Marshal(v)
	}
	return json.Marshal(nil)
}

// UnmarshalJSON ...
func (x *ResponseImageProcessArgs) UnmarshalJSON(bytes []byte) error {
	var err error
	oneof1 := ResponseImageProcessArgs_PointArgs{}
	if err = unmarshal(bytes, &oneof1.PointArgs); err == nil {
		x.Union = &oneof1
		return nil
	}
	oneof2 := ResponseImageProcessArgs_GroundingArgs{}
	if err = unmarshal(bytes, &oneof2.GroundingArgs); err == nil {
		x.Union = &oneof2
		return nil
	}
	oneof3 := ResponseImageProcessArgs_RotateArgs{}
	if err = unmarshal(bytes, &oneof3.RotateArgs); err == nil {
		x.Union = &oneof3
		return nil
	}
	oneof4 := ResponseImageProcessArgs_ZoomArgs{}
	if err = unmarshal(bytes, &oneof4.ZoomArgs); err == nil {
		x.Union = &oneof4
		return nil
	}
	return err
}

// UnmarshalJSON ...
func (r *MessageContent) UnmarshalJSON(bytes []byte) error {
	var err error
	// TODO: more
	oneof1 := MessageContent_StringValue{}
	if err = unmarshal(bytes, &oneof1.StringValue); err == nil {
		r.Union = &oneof1
		return nil
	}
	oneof2 := MessageContent_ListValue{
		ListValue: &ContentItemList{},
	}
	if err = unmarshal(bytes, &oneof2.ListValue.ListValue); err == nil {
		r.Union = &oneof2
		return nil
	}
	return err
}

// MarshalJSON ...
func (r *MessageContent) MarshalJSON() ([]byte, error) {
	if v := r.GetListValue(); v != nil {
		return json.Marshal(v.ListValue)
	}
	v := r.GetStringValue()
	return json.Marshal(v)
}

// UnmarshalJSON ...
func (r *ContentItem) UnmarshalJSON(bytes []byte) error {
	var err error
	var typeOnly struct {
		Type ContentItemType_Enum `json:"type"`
	}
	if err = unmarshal(bytes, &typeOnly); err != nil {
		return err
	}
	switch typeOnly.Type {
	case ContentItemType_input_text:
		oneof := ContentItem_Text{}
		if err = unmarshal(bytes, &oneof.Text); err == nil {
			r.Union = &oneof
			return nil
		}
	case ContentItemType_input_image:
		oneof := ContentItem_Image{}
		if err = unmarshal(bytes, &oneof.Image); err == nil {
			r.Union = &oneof
			return nil
		}
	case ContentItemType_input_video:
		oneof := ContentItem_Video{}
		if err = unmarshal(bytes, &oneof.Video); err == nil {
			r.Union = &oneof
			return nil
		}
	case ContentItemType_input_audio:
		oneof := ContentItem_Audio{}
		if err = unmarshal(bytes, &oneof.Audio); err == nil {
			r.Union = &oneof
			return nil
		}
	case ContentItemType_output_text:
		oneof := ContentItem_Text{}
		if err = unmarshal(bytes, &oneof.Text); err == nil {
			r.Union = &oneof
			return nil
		}
	case ContentItemType_summary_text:
		oneof := ContentItem_Text{}
		if err = unmarshal(bytes, &oneof.Text); err == nil {
			r.Union = &oneof
			return nil
		}
	default:
		err = &json.InvalidUnmarshalError{
			Type: reflect.TypeOf(r),
		}
	}
	return err
}

// MarshalJSON ...
func (r *ContentItem) MarshalJSON() ([]byte, error) {
	if v := r.GetImage(); v != nil {
		return json.Marshal(v)
	}
	if v := r.GetText(); v != nil {
		return json.Marshal(v)
	}
	if v := r.GetVideo(); v != nil {
		return json.Marshal(v)
	}
	if v := r.GetAudio(); v != nil {
		return json.Marshal(v)
	}
	return json.Marshal(nil)
}

// UnmarshalJSON ...
func (r *ResponsesToolChoice) UnmarshalJSON(bytes []byte) error {
	var err error
	mode := ResponsesToolChoice_Mode{}
	if err = unmarshal(bytes, &mode.Mode); err == nil {
		r.Union = &mode
		return nil
	}
	fc := ResponsesToolChoice_FunctionToolChoice{}
	if err = unmarshal(bytes, &fc.FunctionToolChoice); err == nil {
		r.Union = &fc
		return nil
	}
	return err
}

// MarshalJSON ...
func (r *ResponsesToolChoice) MarshalJSON() ([]byte, error) {
	if v := r.GetFunctionToolChoice(); v != nil {
		return json.Marshal(v)
	}
	return json.Marshal(r.GetMode())
}

// UnmarshalJSON ...
func (r *ResponsesTool) UnmarshalJSON(bytes []byte) error {
	var err error
	var typeOnly struct {
		Type ToolType_Enum `json:"type"`
	}
	if err = json.Unmarshal(bytes, &typeOnly); err != nil {
		return err
	}
	switch typeOnly.Type {
	case ToolType_function:
		fc := ResponsesTool_ToolFunction{}
		if err = unmarshal(bytes, &fc.ToolFunction); err == nil {
			r.Union = &fc
			return nil
		}
	case ToolType_web_search:
		ws := ResponsesTool_ToolWebSearch{}
		if err = unmarshal(bytes, &ws.ToolWebSearch); err == nil {
			r.Union = &ws
			return nil
		}
	case ToolType_image_process:
		ip := ResponsesTool_ToolImageProcess{}
		if err = unmarshal(bytes, &ip.ToolImageProcess); err == nil {
			r.Union = &ip
			return nil
		}
	default:
		err = &json.InvalidUnmarshalError{
			Type: reflect.TypeOf(r),
		}
	}
	return err
}

// MarshalJSON ...
func (r *ResponsesTool) MarshalJSON() ([]byte, error) {
	if v := r.GetToolFunction(); v != nil {
		return json.Marshal(v)
	}
	if v := r.GetToolWebSearch(); v != nil {
		return json.Marshal(v)
	}
	if v := r.GetToolImageProcess(); v != nil {
		return json.Marshal(v)
	}
	return json.Marshal(nil)
}

// UnmarshalJSON ...
func (r *OutputContentItem) UnmarshalJSON(bytes []byte) error {
	var err error
	text := OutputContentItem_Text{}
	if err = unmarshal(bytes, &text.Text); err == nil {
		r.Union = &text
		return nil
	}
	return err
}

// MarshalJSON ...
func (r *OutputContentItem) MarshalJSON() ([]byte, error) {
	if v := r.GetText(); v != nil {
		return json.Marshal(v)
	}
	return json.Marshal(nil)
}

// UnmarshalJSON ...
func (r *Bytes) UnmarshalJSON(bytes []byte) error {
	r.Value = bytes
	return nil
}

// MarshalJSON ...
func (r *Bytes) MarshalJSON() ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := json.Compact(buf, r.Value); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// UnmarshalJSON ...
func (r *OutputItem) UnmarshalJSON(bytes []byte) error {
	var err error
	var typeOnly struct {
		Type ItemType_Enum `json:"type"`
	}
	if err = json.Unmarshal(bytes, &typeOnly); err != nil {
		return err
	}
	switch typeOnly.Type {
	case ItemType_reasoning:
		oneof := OutputItem_Reasoning{}
		if err = unmarshal(bytes, &oneof.Reasoning); err == nil {
			r.Union = &oneof
			return nil
		}
	case ItemType_message:
		oneof := OutputItem_OutputMessage{}
		if err = unmarshal(bytes, &oneof.OutputMessage); err == nil {
			r.Union = &oneof
			return nil
		}
	case ItemType_function_call:
		oneof := OutputItem_FunctionToolCall{}
		if err = unmarshal(bytes, &oneof.FunctionToolCall); err == nil {
			r.Union = &oneof
			return nil
		}
	case ItemType_transcription:
		oneof := OutputItem_Transcription{}
		if err = unmarshal(bytes, &oneof.Transcription); err == nil {
			r.Union = &oneof
		}
	case ItemType_web_search_call:
		oneof := OutputItem_FunctionWebSearch{}
		if err = unmarshal(bytes, &oneof.FunctionWebSearch); err == nil {
			r.Union = &oneof
		}
	case ItemType_image_process:
		oneof := OutputItem_FunctionImageProcess{}
		if err = unmarshal(bytes, &oneof.FunctionImageProcess); err == nil {
			r.Union = &oneof
		}
	default:
		err = &json.InvalidUnmarshalError{
			Type: reflect.TypeOf(r),
		}
	}
	return err
}

// MarshalJSON ...
func (r *OutputItem) MarshalJSON() ([]byte, error) {
	if v := r.GetReasoning(); v != nil {
		return json.Marshal(v)
	}
	if v := r.GetOutputMessage(); v != nil {
		return json.Marshal(v)
	}
	if v := r.GetFunctionToolCall(); v != nil {
		return json.Marshal(v)
	}
	if v := r.GetTranscription(); v != nil {
		return json.Marshal(v)
	}
	if v := r.GetFunctionWebSearch(); v != nil {
		return json.Marshal(v)
	}
	if v := r.GetFunctionImageProcess(); v != nil {
		return json.Marshal(v)
	}
	return json.Marshal(nil)
}

// MarshalJSON ...
func (e *Event) MarshalJSON() ([]byte, error) {
	if v := e.GetResponse(); v != nil {
		return json.Marshal(v)
	}
	if v := e.GetItem(); v != nil {
		return json.Marshal(v)
	}
	if v := e.GetContentPart(); v != nil {
		return json.Marshal(v)
	}
	if v := e.GetReasoningPart(); v != nil {
		return json.Marshal(v)
	}
	if v := e.GetReasoningText(); v != nil {
		return json.Marshal(v)
	}
	if v := e.GetText(); v != nil {
		return json.Marshal(v)
	}
	if v := e.GetFunctionCallArguments(); v != nil {
		return json.Marshal(v)
	}
	if v := e.GetError(); v != nil {
		return json.Marshal(v)
	}
	if v := e.GetTranscriptionPart(); v != nil {
		return json.Marshal(v)
	}
	if v := e.GetTranscriptionText(); v != nil {
		return json.Marshal(v)
	}
	if v := e.GetResponseWebSearchCallInProgress(); v != nil {
		return json.Marshal(v)
	}
	if v := e.GetResponseWebSearchCallSearching(); v != nil {
		return json.Marshal(v)
	}
	if v := e.GetResponseWebSearchCallCompleted(); v != nil {
		return json.Marshal(v)
	}
	if v := e.GetResponseAnnotationAdded(); v != nil {
		return json.Marshal(v)
	}
	if v := e.GetResponseImageProcessCallInProgress(); v != nil {
		return json.Marshal(v)
	}
	if v := e.GetResponseImageProcessCallProcessing(); v != nil {
		return json.Marshal(v)
	}
	if v := e.GetResponseImageProcessCallCompleted(); v != nil {
		return json.Marshal(v)
	}
	return json.Marshal(nil)
}
func (e *Event) UnmarshalJSON(bytes []byte) error {
	var typeOnly struct {
		Type EventType_Enum `json:"type"`
	}
	if err := json.Unmarshal(bytes, &typeOnly); err != nil {
		return err
	}

	switch typeOnly.Type {
	case EventType_response_created:
		oneof := Event_Response{}
		if err := unmarshal(bytes, &oneof.Response); err != nil {
			return err
		}
		e.Event = &oneof
	case EventType_response_completed:
		oneof := Event_Response{}
		if err := unmarshal(bytes, &oneof.Response); err != nil {
			return err
		}
		e.Event = &oneof
	case EventType_response_failed:
		oneof := Event_Response{}
		if err := unmarshal(bytes, &oneof.Response); err != nil {
			return err
		}
		e.Event = &oneof
	case EventType_response_incomplete:
		oneof := Event_Response{}
		if err := unmarshal(bytes, &oneof.Response); err != nil {
			return err
		}
		e.Event = &oneof
	case EventType_response_in_progress:
		oneof := Event_Response{}
		if err := unmarshal(bytes, &oneof.Response); err != nil {
			return err
		}
		e.Event = &oneof
	case EventType_response_output_item_added:
		oneof := Event_Item{}
		if err := unmarshal(bytes, &oneof.Item); err != nil {
			return err
		}
		e.Event = &oneof
	case EventType_response_output_item_done:
		oneof := Event_Item{}
		if err := unmarshal(bytes, &oneof.Item); err != nil {
			return err
		}
		e.Event = &oneof
	case EventType_response_content_part_added:
		oneof := Event_ContentPart{}
		if err := unmarshal(bytes, &oneof.ContentPart); err != nil {
			return err
		}
		e.Event = &oneof
	case EventType_response_content_part_done:
		oneof := Event_ContentPart{}
		if err := unmarshal(bytes, &oneof.ContentPart); err != nil {
			return err
		}
		e.Event = &oneof
	case EventType_response_reasoning_summary_part_added:
		oneof := Event_ReasoningPart{}
		if err := unmarshal(bytes, &oneof.ReasoningPart); err != nil {
			return err
		}
		e.Event = &oneof
	case EventType_response_reasoning_summary_part_done:
		oneof := Event_ReasoningPart{}
		if err := unmarshal(bytes, &oneof.ReasoningPart); err != nil {
			return err
		}
		e.Event = &oneof
	case EventType_response_reasoning_summary_text_done:
		oneof := Event_ReasoningText{}
		if err := unmarshal(bytes, &oneof.ReasoningText); err != nil {
			return err
		}
		e.Event = &oneof
	case EventType_response_reasoning_summary_text_delta:
		oneof := Event_ReasoningText{}
		if err := unmarshal(bytes, &oneof.ReasoningText); err != nil {
			return err
		}
		e.Event = &oneof
	case EventType_response_output_text_delta:
		oneof := Event_Text{}
		if err := unmarshal(bytes, &oneof.Text); err != nil {
			return err
		}
		e.Event = &oneof
	case EventType_response_output_text_done:
		oneof := Event_Text{}
		if err := unmarshal(bytes, &oneof.Text); err != nil {
			return err
		}
		e.Event = &oneof
	case EventType_response_output_text_annotation_added:
		oneof := Event_ResponseAnnotationAdded{}
		if err := unmarshal(bytes, &oneof.ResponseAnnotationAdded); err != nil {
			return err
		}
		e.Event = &oneof
	case EventType_response_function_call_arguments_delta:
		oneof := Event_FunctionCallArguments{}
		if err := unmarshal(bytes, &oneof.FunctionCallArguments); err != nil {
			return err
		}
		e.Event = &oneof
	case EventType_response_function_call_arguments_done:
		oneof := Event_FunctionCallArguments{}
		if err := unmarshal(bytes, &oneof.FunctionCallArguments); err != nil {
			return err
		}
		e.Event = &oneof
	case EventType_response_web_search_call_completed:
		oneof := Event_ResponseWebSearchCallCompleted{}
		if err := unmarshal(bytes, &oneof.ResponseWebSearchCallCompleted); err != nil {
			return err
		}
		e.Event = &oneof
	case EventType_response_web_search_call_searching:
		oneof := Event_ResponseWebSearchCallSearching{}
		if err := unmarshal(bytes, &oneof.ResponseWebSearchCallSearching); err != nil {
			return err
		}
		e.Event = &oneof
	case EventType_response_web_search_call_in_progress:
		oneof := Event_ResponseWebSearchCallInProgress{}
		if err := unmarshal(bytes, &oneof.ResponseWebSearchCallInProgress); err != nil {
			return err
		}
		e.Event = &oneof
	case EventType_response_image_process_call_completed:
		oneof := Event_ResponseImageProcessCallCompleted{}
		if err := unmarshal(bytes, &oneof.ResponseImageProcessCallCompleted); err != nil {
			return err
		}
		e.Event = &oneof
	case EventType_response_image_process_call_in_progress:
		oneof := Event_ResponseImageProcessCallInProgress{}
		if err := unmarshal(bytes, &oneof.ResponseImageProcessCallInProgress); err != nil {
			return err
		}
		e.Event = &oneof
	case EventType_response_image_process_call_progressing:
		oneof := Event_ResponseImageProcessCallProcessing{}
		if err := unmarshal(bytes, &oneof.ResponseImageProcessCallProcessing); err != nil {
			return err
		}
		e.Event = &oneof
	default:
		return &json.InvalidUnmarshalError{Type: reflect.TypeOf(e)}
	}
	return nil
}

// URLQuery serializes [GetResponseRequest]'s query parameters as `url.Values`.
func (x *GetResponseRequest) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(x, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// URLQuery serializes [ListInputItemsRequest]'s query parameters as `url.Values`.
func (x *ListInputItemsRequest) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(x, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatBrackets,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// GetEventType ...
func (e *Event) GetEventType() string {
	if e == nil {
		return ""
	}
	eventType := EventType_unspecified
	if e.GetResponse() != nil {
		eventType = e.GetResponse().GetType()
	} else if e.GetText() != nil {
		eventType = e.GetText().GetType()
	} else if e.GetReasoningText() != nil {
		eventType = e.GetReasoningText().GetType()
	} else if e.GetItem() != nil {
		eventType = e.GetItem().GetType()
	} else if e.GetFunctionCallArguments() != nil {
		eventType = e.GetFunctionCallArguments().GetType()
	} else if e.GetError() != nil {
		eventType = e.GetError().GetType()
	} else if e.GetContentPart() != nil {
		eventType = e.GetContentPart().GetType()
	} else if e.GetReasoningPart() != nil {
		eventType = e.GetReasoningPart().GetType()
	} else if e.GetTranscriptionPart() != nil {
		eventType = e.GetTranscriptionPart().GetType()
	} else if e.GetTranscriptionText() != nil {
		eventType = e.GetTranscriptionText().GetType()
	} else if e.GetResponseAnnotationAdded() != nil {
		eventType = e.GetResponseAnnotationAdded().GetType()
	} else if e.GetResponseWebSearchCallInProgress() != nil {
		eventType = e.GetResponseWebSearchCallInProgress().GetType()
	} else if e.GetResponseWebSearchCallSearching() != nil {
		eventType = e.GetResponseWebSearchCallSearching().GetType()
	} else if e.GetResponseWebSearchCallCompleted() != nil {
		eventType = e.GetResponseWebSearchCallCompleted().GetType()
	} else if e.GetResponseImageProcessCallInProgress() != nil {
		eventType = e.GetResponseImageProcessCallInProgress().GetType()
	} else if e.GetResponseImageProcessCallProcessing() != nil {
		eventType = e.GetResponseImageProcessCallProcessing().GetType()
	} else if e.GetResponseImageProcessCallCompleted() != nil {
		eventType = e.GetResponseImageProcessCallCompleted().GetType()
	}
	marshaled, _ := json.Marshal(eventType)
	eventTypeString := string(marshaled)
	return strings.Trim(eventTypeString, "\"")
}

// IsDelta ...
func (e *Event) IsDelta() bool {
	if (e.GetText() != nil && e.GetText().GetType() == EventType_response_output_text_delta) ||
		(e.GetReasoningText() != nil && e.GetReasoningText().GetType() == EventType_response_output_text_delta) ||
		(e.GetFunctionCallArguments() != nil && e.GetFunctionCallArguments().GetType() == EventType_response_function_call_arguments_delta) ||
		(e.GetTranscriptionText() != nil && e.GetTranscriptionText().GetType() == EventType_response_transcription_text_delta) ||
		(e.GetResponseAnnotationAdded() != nil) ||
		(e.GetResponseWebSearchCallInProgress() != nil) ||
		(e.GetResponseWebSearchCallSearching() != nil) ||
		(e.GetResponseWebSearchCallCompleted() != nil) ||
		(e.GetResponseImageProcessCallInProgress() != nil) ||
		(e.GetResponseImageProcessCallCompleted() != nil) {
		return true
	}
	return false
}

// IsDeltaDone ...
func (e *Event) IsDeltaDone() bool {
	if (e.GetText() != nil &&
		e.GetText().GetType() == EventType_response_output_text_done) ||
		(e.GetReasoningText() != nil &&
			e.GetReasoningText().GetType() == EventType_response_reasoning_summary_text_done) ||
		(e.GetFunctionCallArguments() != nil &&
			e.GetFunctionCallArguments().GetType() == EventType_response_function_call_arguments_done) ||
		(e.GetTranscriptionText() != nil &&
			e.GetTranscriptionText().GetType() == EventType_response_transcription_text_done) {
		return true
	}
	return false
}

// IsResponseDone ...
func (e *Event) IsResponseDone() bool {
	if respEvent := e.GetResponse(); respEvent != nil {
		if respEvent.Type == EventType_response_completed ||
			respEvent.Type == EventType_response_failed ||
			respEvent.Type == EventType_response_incomplete {
			return true
		}
	}
	return false
}

// GetParamType ...
func (x *ResponseImageProcessArgs) GetParamType() ResponseImageProcessType_Enum {
	if x.GetPointArgs() != nil {
		return ResponseImageProcessType_point
	}
	if x.GetGroundingArgs() != nil {
		return ResponseImageProcessType_grounding
	}
	if x.GetRotateArgs() != nil {
		return ResponseImageProcessType_rotate
	}
	if x.GetZoomArgs() != nil {
		return ResponseImageProcessType_zoom
	}
	return ResponseImageProcessType_unspecified
}

// GetImageIndex ...
func (x *ResponseImageProcessArgs) GetImageIndex() int {
	if x.GetPointArgs() != nil {
		return int(x.GetPointArgs().GetImageIndex())
	}
	if x.GetGroundingArgs() != nil {
		return int(x.GetGroundingArgs().GetImageIndex())
	}
	if x.GetRotateArgs() != nil {
		return int(x.GetRotateArgs().GetImageIndex())
	}
	if x.GetZoomArgs() != nil {
		return int(x.GetZoomArgs().GetImageIndex())
	}
	return 0
}

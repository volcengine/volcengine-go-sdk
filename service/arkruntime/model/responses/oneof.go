package responses

import (
	"bytes"
	"encoding/json"
	"reflect"
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
	case *typeOnly.Type == ItemType_mcp_approval_request:
		mcpApprovalRequest := InputItem_McpApprovalRequest{}
		if err = unmarshal(bytes, &mcpApprovalRequest.McpApprovalRequest); err == nil {
			r.Union = &mcpApprovalRequest
			return nil
		}
	case *typeOnly.Type == ItemType_mcp_approval_response:
		mcpApprovalResponse := InputItem_McpApprovalResponse{}
		if err = unmarshal(bytes, &mcpApprovalResponse.McpApprovalResponse); err == nil {
			r.Union = &mcpApprovalResponse
			return nil
		}
	case *typeOnly.Type == ItemType_mcp_list_tools:
		mcpListTools := InputItem_McpListTools{}
		if err = unmarshal(bytes, &mcpListTools.McpListTools); err == nil {
			r.Union = &mcpListTools
			return nil
		}
	case *typeOnly.Type == ItemType_mcp_call:
		mcpCall := ItemFunctionMcpCall{}
		if err = unmarshal(bytes, &mcpCall); err == nil {
			r.Union = &InputItem_FunctionMcpCall{
				FunctionMcpCall: &mcpCall,
			}
			return nil
		}
	case *typeOnly.Type == ItemType_web_search_call:
		webSearchCall := ItemFunctionWebSearch{}
		if err = unmarshal(bytes, &webSearchCall); err == nil {
			r.Union = &InputItem_FunctionWebSearchCall{
				FunctionWebSearchCall: &webSearchCall,
			}
			return nil
		}
	case *typeOnly.Type == ItemType_knowledge_search_call:
		knowledgeSearchCall := ItemFunctionKnowledgeSearch{}
		if err = unmarshal(bytes, &knowledgeSearchCall); err == nil {
			r.Union = &InputItem_FunctionKnowledgeSearch{
				FunctionKnowledgeSearch: &knowledgeSearchCall,
			}
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
	if v := r.GetMcpApprovalRequest(); v != nil {
		return json.Marshal(v)
	}
	if v := r.GetMcpApprovalResponse(); v != nil {
		return json.Marshal(v)
	}
	if v := r.GetMcpListTools(); v != nil {
		return json.Marshal(v)
	}
	if v := r.GetFunctionMcpCall(); v != nil {
		return json.Marshal(v)
	}
	if v := r.GetFunctionWebSearchCall(); v != nil {
		return json.Marshal(v)
	}
	if v := r.GetFunctionKnowledgeSearch(); v != nil {
		return json.Marshal(v)
	}
	return json.Marshal(nil)
}

// MarshalJSON ...
func (i *ResponseImageProcessArgs) MarshalJSON() ([]byte, error) {
	if v := i.GetPointArgs(); v != nil {
		return json.Marshal(v)
	}
	if v := i.GetGroundingArgs(); v != nil {
		return json.Marshal(v)
	}
	if v := i.GetRotateArgs(); v != nil {
		return json.Marshal(v)
	}
	if v := i.GetZoomArgs(); v != nil {
		return json.Marshal(v)
	}
	return json.Marshal(nil)
}

// UnmarshalJSON ...
func (i *ResponseImageProcessArgs) UnmarshalJSON(bytes []byte) error {
	var err error
	oneof1 := ResponseImageProcessArgs_PointArgs{}
	if err = unmarshal(bytes, &oneof1.PointArgs); err == nil {
		i.Union = &oneof1
		return nil
	}
	oneof2 := ResponseImageProcessArgs_GroundingArgs{}
	if err = unmarshal(bytes, &oneof2.GroundingArgs); err == nil {
		i.Union = &oneof2
		return nil
	}
	oneof3 := ResponseImageProcessArgs_RotateArgs{}
	if err = unmarshal(bytes, &oneof3.RotateArgs); err == nil {
		i.Union = &oneof3
		return nil
	}
	oneof4 := ResponseImageProcessArgs_ZoomArgs{}
	if err = unmarshal(bytes, &oneof4.ZoomArgs); err == nil {
		i.Union = &oneof4
		return nil
	}
	return err
}

// MarshalJSON ...
func (r *McpAllowedTools) MarshalJSON() ([]byte, error) {
	if v := r.GetList(); v != nil {
		// Serialize as a string array, removing the "values" wrapper
		return json.Marshal(v.Values)
	}
	if v := r.GetFilter(); v != nil {
		// 序列化为对象，保持 "tool_names" 字段
		return json.Marshal(v)
	}
	return json.Marshal(nil)
}

// UnmarshalJSON ...
func (r *McpAllowedTools) UnmarshalJSON(bytes []byte) error {
	var err error

	// Try to parse as a string array
	var stringArray []string
	if err = json.Unmarshal(bytes, &stringArray); err == nil {
		list := McpAllowedTools_List{
			List: &McpAllowedToolsList{Values: stringArray},
		}
		r.Union = &list
		return nil
	}

	// Try to parse as filter
	filter := McpAllowedTools_Filter{
		Filter: &McpAllowedToolsFilter{},
	}
	if err = unmarshal(bytes, &filter.Filter); err == nil {
		r.Union = &filter
		return nil
	}

	return err
}

// MarshalJSON ...
func (r *McpRequireApproval) MarshalJSON() ([]byte, error) {
	if v := r.GetFilter(); v != nil {
		return json.Marshal(v)
	}
	return json.Marshal(r.GetMode())
}

// UnmarshalJSON ...
func (r *McpRequireApproval) UnmarshalJSON(bytes []byte) error {
	var err error
	mode := McpRequireApproval_Mode{}
	if err = unmarshal(bytes, &mode.Mode); err == nil {
		r.Union = &mode
		return nil
	}
	fc := McpRequireApproval_Filter{
		Filter: &McpToolApprovalFilter{},
	}
	if err = unmarshal(bytes, &fc.Filter); err == nil {
		r.Union = &fc
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
	if err = json.Unmarshal(bytes, &typeOnly); err != nil {
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
	case ContentItemType_input_file:
		oneof := ContentItem_File{}
		if err = unmarshal(bytes, &oneof.File); err == nil {
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
	if v := r.GetFile(); v != nil {
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

	var typeOnly struct {
		Type ToolType_Enum `json:"type"`
	}
	if err = json.Unmarshal(bytes, &typeOnly); err != nil {
		return err
	}
	switch typeOnly.Type {
	case ToolType_function:
		fc := ResponsesToolChoice_FunctionToolChoice{}
		if err = unmarshal(bytes, &fc.FunctionToolChoice); err == nil {
			r.Union = &fc
			return nil
		}
	case ToolType_web_search:
		ws := ResponsesToolChoice_WebSearchToolChoice{}
		if err = unmarshal(bytes, &ws.WebSearchToolChoice); err == nil {
			r.Union = &ws
			return nil
		}
	case ToolType_mcp:
		mcp := ResponsesToolChoice_McpToolChoice{}
		if err = unmarshal(bytes, &mcp.McpToolChoice); err == nil {
			r.Union = &mcp
			return nil
		}
	case ToolType_knowledge_search:
		ks := ResponsesToolChoice_KnowledgeSearchToolChoice{}
		if err = unmarshal(bytes, &ks.KnowledgeSearchToolChoice); err == nil {
			r.Union = &ks
			return nil
		}
	default:
		err = &json.InvalidUnmarshalError{
			Type: reflect.TypeOf(r),
		}
	}
	return err
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
	case ToolType_mcp:
		mcp := ResponsesTool_ToolMcp{}
		if err = unmarshal(bytes, &mcp.ToolMcp); err == nil {
			r.Union = &mcp
			return nil
		}
	case ToolType_knowledge_search:
		ks := ResponsesTool_ToolKnowledgeSearch{}
		if err = unmarshal(bytes, &ks.ToolKnowledgeSearch); err == nil {
			r.Union = &ks
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
	if v := r.GetToolMcp(); v != nil {
		return json.Marshal(v)
	}
	if v := r.GetToolKnowledgeSearch(); v != nil {
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
			return nil
		}
	case ItemType_web_search_call:
		oneof := OutputItem_FunctionWebSearch{}
		if err = unmarshal(bytes, &oneof.FunctionWebSearch); err == nil {
			r.Union = &oneof
			return nil
		}
	case ItemType_knowledge_search_call:
		oneof := OutputItem_FunctionKnowledgeSearch{}
		if err = unmarshal(bytes, &oneof.FunctionKnowledgeSearch); err == nil {
			r.Union = &oneof
			return nil
		}
	case ItemType_image_process:
		oneof := OutputItem_FunctionImageProcess{}
		if err = unmarshal(bytes, &oneof.FunctionImageProcess); err == nil {
			r.Union = &oneof
			return nil
		}
	case ItemType_mcp_approval_request:
		oneof := OutputItem_FunctionMcpApprovalRequest{}
		if err = unmarshal(bytes, &oneof.FunctionMcpApprovalRequest); err == nil {
			r.Union = &oneof
			return nil
		}
	case ItemType_mcp_call:
		oneof := OutputItem_FunctionMcpCall{}
		if err = unmarshal(bytes, &oneof.FunctionMcpCall); err == nil {
			r.Union = &oneof
			return nil
		}
	case ItemType_mcp_list_tools:
		oneof := OutputItem_FunctionMcpListTools{}
		if err = unmarshal(bytes, &oneof.FunctionMcpListTools); err == nil {
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
	if v := r.GetFunctionMcpApprovalRequest(); v != nil {
		return json.Marshal(v)
	}
	if v := r.GetFunctionMcpListTools(); v != nil {
		return json.Marshal(v)
	}
	if v := r.GetFunctionMcpCall(); v != nil {
		return json.Marshal(v)
	}
	if v := r.GetFunctionKnowledgeSearch(); v != nil {
		return json.Marshal(v)
	}
	return json.Marshal(nil)
}

// MarshalJSON ...
func (e *Event) MarshalJSON() ([]byte, error) {
	if v := e.GetResponse(); v != nil {
		return json.Marshal(v)
	}
	if v := e.GetResponseInProgress(); v != nil {
		return json.Marshal(v)
	}
	if v := e.GetResponseCompleted(); v != nil {
		return json.Marshal(v)
	}
	if v := e.GetResponseFailed(); v != nil {
		return json.Marshal(v)
	}
	if v := e.GetResponseIncomplete(); v != nil {
		return json.Marshal(v)
	}
	if v := e.GetItem(); v != nil {
		return json.Marshal(v)
	}
	if v := e.GetItemDone(); v != nil {
		return json.Marshal(v)
	}
	if v := e.GetContentPart(); v != nil {
		return json.Marshal(v)
	}
	if v := e.GetContentPartDone(); v != nil {
		return json.Marshal(v)
	}
	if v := e.GetReasoningPart(); v != nil {
		return json.Marshal(v)
	}
	if v := e.GetReasoningPartDone(); v != nil {
		return json.Marshal(v)
	}
	if v := e.GetReasoningText(); v != nil {
		return json.Marshal(v)
	}
	if v := e.GetReasoningTextDone(); v != nil {
		return json.Marshal(v)
	}
	if v := e.GetText(); v != nil {
		return json.Marshal(v)
	}
	if v := e.GetTextDone(); v != nil {
		return json.Marshal(v)
	}
	if v := e.GetFunctionCallArguments(); v != nil {
		return json.Marshal(v)
	}
	if v := e.GetFunctionCallArgumentsDone(); v != nil {
		return json.Marshal(v)
	}
	if v := e.GetError(); v != nil {
		return json.Marshal(v)
	}
	if v := e.GetTranscriptionPart(); v != nil {
		return json.Marshal(v)
	}
	if v := e.GetTranscriptionPartDone(); v != nil {
		return json.Marshal(v)
	}
	if v := e.GetTranscriptionText(); v != nil {
		return json.Marshal(v)
	}
	if v := e.GetTranscriptionTextDone(); v != nil {
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
	if v := e.GetResponseMcpListToolsInProgress(); v != nil {
		return json.Marshal(v)
	}
	if v := e.GetResponseMcpListToolsCompleted(); v != nil {
		return json.Marshal(v)
	}
	if v := e.GetResponseMcpCallInProgress(); v != nil {
		return json.Marshal(v)
	}
	if v := e.GetResponseMcpCallArgumentsDelta(); v != nil {
		return json.Marshal(v)
	}
	if v := e.GetResponseMcpCallArgumentsDone(); v != nil {
		return json.Marshal(v)
	}
	if v := e.GetResponseMcpCallCompleted(); v != nil {
		return json.Marshal(v)
	}
	if v := e.GetResponseMcpCallFailed(); v != nil {
		return json.Marshal(v)
	}
	if v := e.GetResponseMcpApprovalRequest(); v != nil {
		return json.Marshal(v)
	}
	if v := e.GetResponseKnowledgeSearchCallInProgress(); v != nil {
		return json.Marshal(v)
	}
	if v := e.GetResponseKnowledgeSearchCallSearching(); v != nil {
		return json.Marshal(v)
	}
	if v := e.GetResponseKnowledgeSearchCallCompleted(); v != nil {
		return json.Marshal(v)
	}
	if v := e.GetResponseKnowledgeSearchCallFailed(); v != nil {
		return json.Marshal(v)
	}
	return json.Marshal(nil)
}

// UnmarshalJSON ...
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
		oneof := Event_ResponseCompleted{}
		if err := unmarshal(bytes, &oneof.ResponseCompleted); err != nil {
			return err
		}
		e.Event = &oneof
	case EventType_response_failed:
		oneof := Event_ResponseFailed{}
		if err := unmarshal(bytes, &oneof.ResponseFailed); err != nil {
			return err
		}
		e.Event = &oneof
	case EventType_response_incomplete:
		oneof := Event_ResponseIncomplete{}
		if err := unmarshal(bytes, &oneof.ResponseIncomplete); err != nil {
			return err
		}
		e.Event = &oneof
	case EventType_response_in_progress:
		oneof := Event_ResponseInProgress{}
		if err := unmarshal(bytes, &oneof.ResponseInProgress); err != nil {
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
		oneof := Event_ItemDone{}
		if err := unmarshal(bytes, &oneof.ItemDone); err != nil {
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
		oneof := Event_ContentPartDone{}
		if err := unmarshal(bytes, &oneof.ContentPartDone); err != nil {
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
		oneof := Event_ReasoningPartDone{}
		if err := unmarshal(bytes, &oneof.ReasoningPartDone); err != nil {
			return err
		}
		e.Event = &oneof
	case EventType_response_reasoning_summary_text_done:
		oneof := Event_ReasoningTextDone{}
		if err := unmarshal(bytes, &oneof.ReasoningTextDone); err != nil {
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
		oneof := Event_TextDone{}
		if err := unmarshal(bytes, &oneof.TextDone); err != nil {
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
		oneof := Event_FunctionCallArgumentsDone{}
		if err := unmarshal(bytes, &oneof.FunctionCallArgumentsDone); err != nil {
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
	case EventType_response_mcp_list_tools_in_progress:
		oneof := Event_ResponseMcpListToolsInProgress{}
		if err := unmarshal(bytes, &oneof.ResponseMcpListToolsInProgress); err != nil {
			return err
		}
		e.Event = &oneof
	case EventType_response_mcp_list_tools_completed:
		oneof := Event_ResponseMcpListToolsCompleted{}
		if err := unmarshal(bytes, &oneof.ResponseMcpListToolsCompleted); err != nil {
			return err
		}
		e.Event = &oneof
	case EventType_response_mcp_call_in_progress:
		oneof := Event_ResponseMcpCallInProgress{}
		if err := unmarshal(bytes, &oneof.ResponseMcpCallInProgress); err != nil {
			return err
		}
		e.Event = &oneof
	case EventType_response_mcp_call_arguments_delta:
		oneof := Event_ResponseMcpCallArgumentsDelta{}
		if err := unmarshal(bytes, &oneof.ResponseMcpCallArgumentsDelta); err != nil {
			return err
		}
		e.Event = &oneof
	case EventType_response_mcp_call_arguments_done:
		oneof := Event_ResponseMcpCallArgumentsDone{}
		if err := unmarshal(bytes, &oneof.ResponseMcpCallArgumentsDone); err != nil {
			return err
		}
		e.Event = &oneof
	case EventType_response_mcp_call_completed:
		oneof := Event_ResponseMcpCallCompleted{}
		if err := unmarshal(bytes, &oneof.ResponseMcpCallCompleted); err != nil {
			return err
		}
		e.Event = &oneof
	case EventType_response_mcp_call_failed:
		oneof := Event_ResponseMcpCallFailed{}
		if err := unmarshal(bytes, &oneof.ResponseMcpCallFailed); err != nil {
			return err
		}
		e.Event = &oneof
	case EventType_response_knowledge_search_call_in_progress:
		oneof := Event_ResponseKnowledgeSearchCallInProgress{}
		if err := unmarshal(bytes, &oneof.ResponseKnowledgeSearchCallInProgress); err != nil {
			return err
		}
		e.Event = &oneof
	case EventType_response_knowledge_search_call_searching:
		oneof := Event_ResponseKnowledgeSearchCallSearching{}
		if err := unmarshal(bytes, &oneof.ResponseKnowledgeSearchCallSearching); err != nil {
			return err
		}
		e.Event = &oneof
	case EventType_response_knowledge_search_call_completed:
		oneof := Event_ResponseKnowledgeSearchCallCompleted{}
		if err := unmarshal(bytes, &oneof.ResponseKnowledgeSearchCallCompleted); err != nil {
			return err
		}
		e.Event = &oneof
	case EventType_response_knowledge_search_call_failed:
		oneof := Event_ResponseKnowledgeSearchCallFailed{}
		if err := unmarshal(bytes, &oneof.ResponseKnowledgeSearchCallFailed); err != nil {
			return err
		}
		e.Event = &oneof
	default:
		return &json.InvalidUnmarshalError{Type: reflect.TypeOf(e)}
	}
	return nil
}

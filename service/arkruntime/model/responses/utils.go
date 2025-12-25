package responses

import (
	"encoding/json"
	"net/url"
	"reflect"
	"strings"

	"github.com/volcengine/volcengine-go-sdk/service/arkruntime/pkg/apiquery"
)

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

// GetEventType ...
func (e *Event) GetEventType() string {
	if e == nil {
		return ""
	}
	// 1. Detect which oneof is set
	ev := e.GetEvent()
	if ev == nil {
		return ""
	}

	// 2. Find the underlying GetXxx method name
	oneofType := reflect.TypeOf(ev)
	if oneofType.Kind() == reflect.Ptr {
		oneofType = oneofType.Elem()
	}
	name := oneofType.Name()
	if !strings.HasPrefix(name, "Event_") {
		return ""
	}
	suffix := name[len("Event_"):]
	getterName := "Get" + suffix

	// 3. Call event.GetXxx() → underlying message (with .GetType())
	method := reflect.ValueOf(e).MethodByName(getterName)
	if !method.IsValid() {
		return ""
	}
	res := method.Call(nil)
	if len(res) == 0 || res[0].IsNil() {
		return ""
	}
	inner := res[0]

	// 4. Call inner.GetType()
	getType := inner.MethodByName("GetType")
	if !getType.IsValid() {
		return ""
	}

	typeRes := getType.Call(nil)
	if len(typeRes) == 0 {
		return ""
	}

	enumValue := typeRes[0].Interface()
	// 5. Marshal with Sonic
	raw, _ := json.Marshal(enumValue)
	return strings.Trim(string(raw), "\"")
}

// IsDelta ...
func (e *Event) IsDelta() bool {
	if (e.GetText() != nil) ||
		(e.GetReasoningText() != nil) ||
		(e.GetFunctionCallArguments() != nil) ||
		(e.GetTranscriptionText() != nil) ||
		(e.GetResponseAnnotationAdded() != nil) ||
		(e.GetResponseWebSearchCallInProgress() != nil) ||
		(e.GetResponseWebSearchCallSearching() != nil) ||
		(e.GetResponseWebSearchCallCompleted() != nil) ||
		(e.GetResponseImageProcessCallInProgress() != nil) ||
		(e.GetResponseImageProcessCallCompleted() != nil) ||
		(e.GetResponseMcpListToolsCompleted() != nil) ||
		(e.GetResponseMcpCallInProgress() != nil) ||
		(e.GetResponseMcpCallArgumentsDelta() != nil) ||
		(e.GetResponseMcpCallArgumentsDone() != nil) ||
		(e.GetResponseMcpCallCompleted() != nil) ||
		(e.GetResponseMcpCallFailed() != nil) ||
		(e.GetResponseKnowledgeSearchCallInProgress() != nil) ||
		(e.GetResponseKnowledgeSearchCallSearching() != nil) ||
		(e.GetResponseKnowledgeSearchCallCompleted() != nil) ||
		(e.GetResponseKnowledgeSearchCallFailed() != nil) {
		return true
	}
	return false
}

// IsDeltaDone ...
func (e *Event) IsDeltaDone() bool {
	if (e.GetTextDone() != nil) ||
		(e.GetReasoningTextDone() != nil) ||
		(e.GetFunctionCallArgumentsDone() != nil) ||
		(e.GetTranscriptionTextDone() != nil) {
		return true
	}
	return false
}

// IsResponseDone ...
func (e *Event) IsResponseDone() bool {
	if e.GetResponseCompleted() != nil ||
		e.GetResponseFailed() != nil ||
		e.GetResponseIncomplete() != nil {
		return true
	}
	return false
}

package responses

import (
	"encoding/json"
	"net/url"
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
	eventType := EventType_unspecified
	if e.GetResponse() != nil {
		eventType = e.GetResponse().GetType()
	} else if e.GetResponseInProgress() != nil {
		eventType = e.GetResponseInProgress().GetType()
	} else if e.GetResponseCompleted() != nil {
		eventType = e.GetResponseCompleted().GetType()
	} else if e.GetResponseFailed() != nil {
		eventType = e.GetResponseFailed().GetType()
	} else if e.GetResponseIncomplete() != nil {
		eventType = e.GetResponseIncomplete().GetType()
	} else if e.GetText() != nil {
		eventType = e.GetText().GetType()
	} else if e.GetTextDone() != nil {
		eventType = e.GetTextDone().GetType()
	} else if e.GetReasoningText() != nil {
		eventType = e.GetReasoningText().GetType()
	} else if e.GetReasoningTextDone() != nil {
		eventType = e.GetReasoningTextDone().GetType()
	} else if e.GetItem() != nil {
		eventType = e.GetItem().GetType()
	} else if e.GetItemDone() != nil {
		eventType = e.GetItemDone().GetType()
	} else if e.GetFunctionCallArguments() != nil {
		eventType = e.GetFunctionCallArguments().GetType()
	} else if e.GetFunctionCallArgumentsDone() != nil {
		eventType = e.GetFunctionCallArgumentsDone().GetType()
	} else if e.GetError() != nil {
		eventType = e.GetError().GetType()
	} else if e.GetContentPart() != nil {
		eventType = e.GetContentPart().GetType()
	} else if e.GetContentPartDone() != nil {
		eventType = e.GetContentPartDone().GetType()
	} else if e.GetReasoningPart() != nil {
		eventType = e.GetReasoningPart().GetType()
	} else if e.GetReasoningPartDone() != nil {
		eventType = e.GetReasoningPartDone().GetType()
	} else if e.GetTranscriptionPart() != nil {
		eventType = e.GetTranscriptionPart().GetType()
	} else if e.GetTranscriptionPartDone() != nil {
		eventType = e.GetTranscriptionPartDone().GetType()
	} else if e.GetTranscriptionText() != nil {
		eventType = e.GetTranscriptionText().GetType()
	} else if e.GetTranscriptionTextDone() != nil {
		eventType = e.GetTranscriptionTextDone().GetType()
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
	} else if e.GetResponseMcpListToolsInProgress() != nil {
		eventType = e.GetResponseMcpListToolsInProgress().GetType()
	} else if e.GetResponseMcpListToolsCompleted() != nil {
		eventType = e.GetResponseMcpListToolsCompleted().GetType()
	} else if e.GetResponseMcpCallInProgress() != nil {
		eventType = e.GetResponseMcpCallInProgress().GetType()
	} else if e.GetResponseMcpCallArgumentsDelta() != nil {
		eventType = e.GetResponseMcpCallArgumentsDelta().GetType()
	} else if e.GetResponseMcpCallArgumentsDone() != nil {
		eventType = e.GetResponseMcpCallArgumentsDone().GetType()
	} else if e.GetResponseMcpCallCompleted() != nil {
		eventType = e.GetResponseMcpCallCompleted().GetType()
	} else if e.GetResponseMcpCallFailed() != nil {
		eventType = e.GetResponseMcpCallFailed().GetType()
	} else if e.GetResponseKnowledgeSearchCallInProgress() != nil {
		eventType = e.GetResponseKnowledgeSearchCallInProgress().GetType()
	} else if e.GetResponseKnowledgeSearchCallSearching() != nil {
		eventType = e.GetResponseKnowledgeSearchCallSearching().GetType()
	} else if e.GetResponseKnowledgeSearchCallCompleted() != nil {
		eventType = e.GetResponseKnowledgeSearchCallCompleted().GetType()
	} else if e.GetResponseKnowledgeSearchCallFailed() != nil {
		eventType = e.GetResponseKnowledgeSearchCallFailed().GetType()
	}
	marshaled, _ := json.Marshal(eventType)
	eventTypeString := string(marshaled)
	return strings.Trim(eventTypeString, "\"")
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

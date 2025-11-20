package doubao_app

import (
	"encoding/json"
	"fmt"
)

func unmarshal(message []byte, val interface{}) error {
	return json.Unmarshal(message, val)
}

func (e *SseResponseChunkData) MarshalJSON() ([]byte, error) {
	if e.GetSseAck() != nil {
		return json.Marshal(e.GetSseAck())
	} else if e.GetStreamMsgNotify() != nil {
		return json.Marshal(e.GetStreamMsgNotify())
	} else if e.GetStreamChunk() != nil {
		return json.Marshal(e.GetStreamChunk())
	} else if e.GetSseReplyEnd() != nil {
		return json.Marshal(e.GetSseReplyEnd())
	}
	return nil, fmt.Errorf("invalid SseResponseChunkData")
}

func (e *SseResponseChunkData) UnmarshalJSON(data []byte) error {
	// Try StreamMsgNotify first since it's the most common case for streaming data
	oneof1 := SseResponseChunkData_StreamMsgNotify{}
	if err := unmarshal(data, &oneof1.StreamMsgNotify); err == nil && oneof1.StreamMsgNotify.Content != nil {
		e.Union = &oneof1
		return nil
	}

	// Try StreamChunk next
	oneof2 := SseResponseChunkData_StreamChunk{}
	if err := unmarshal(data, &oneof2.StreamChunk); err == nil && oneof2.StreamChunk.PatchOp != nil {
		e.Union = &oneof2
		return nil
	}

	// Try SseReplyEnd
	oneof3 := SseResponseChunkData_SseReplyEnd{}
	if err := unmarshal(data, &oneof3.SseReplyEnd); err == nil && oneof3.SseReplyEnd.EndType > 0 {
		e.Union = &oneof3
		return nil
	}

	// Try SseAck
	oneof4 := SseResponseChunkData_SseAck{}
	if err := unmarshal(data, &oneof4.SseAck); err == nil && oneof4.SseAck.TimeoutConf != nil {
		e.Union = &oneof4
		return nil
	}

	return fmt.Errorf("unmarshal SseResponseChunkData failed, raw data: %s", string(data))
}

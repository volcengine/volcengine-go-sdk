package wafruntime

import (
	"fmt"
	"github.com/volcengine/volcengine-go-sdk/service/waf"
)

const (
	LLM_STREAM_SEND_BASE_WINDOW int64 = 10
	LLM_STREAM_SEND_EXPONENT    int64 = 2
)

// LLMStreamSession 表示一个流会话的结构体，包含流缓冲区、流发送长度和消息 ID 等信息。
type LLMStreamSession struct {
	// 流缓冲区，用于存储流数据
	StreamBuf string
	// 流发送的长度
	StreamSendLen int
	// 消息的唯一标识符
	MsgID string
	// 存储默认响应体
	DefaultBody       *waf.CheckLLMResponseStreamOutput
	currentSendWindow int64
}

// GetStreamBuf 获取流缓冲区的值。
func (s *LLMStreamSession) GetStreamBuf() string {
	return s.StreamBuf
}

// SetStreamBuf 设置流缓冲区的值。
func (s *LLMStreamSession) SetStreamBuf(streamBuf string) {
	s.StreamBuf = streamBuf
}

// GetStreamSendLen 获取流发送长度的值。
func (s *LLMStreamSession) GetStreamSendLen() int {
	return s.StreamSendLen
}

// SetStreamSendLen 设置流发送长度的值。
func (s *LLMStreamSession) SetStreamSendLen(streamSendLen int) {
	s.StreamSendLen = streamSendLen
}

// GetMsgID 获取消息 ID 的值。
func (s *LLMStreamSession) GetMsgID() string {
	return s.MsgID
}

// SetMsgID 设置消息 ID 的值。
func (s *LLMStreamSession) SetMsgID(msgID string) {
	s.MsgID = msgID
}

// GetDefaultBody 获取默认响应体。
func (s *LLMStreamSession) GetDefaultBody() *waf.CheckLLMResponseStreamOutput {
	return s.DefaultBody
}

// SetDefaultBody 设置默认响应体。
func (s *LLMStreamSession) SetDefaultBody(defaultBody *waf.CheckLLMResponseStreamOutput) {
	s.DefaultBody = defaultBody
}

// AppendStreamBuf 向流缓冲区追加一个字符串，并更新流发送长度。
// 如果 defaultBody 不为空，也可以在这里进行相关操作
func (s *LLMStreamSession) AppendStreamBuf(str string) {
	if str != "" {
		s.StreamBuf += str
		s.StreamSendLen += len(str)
	}
}

// String 重写 String 方法，返回对象的字符串表示形式。
func (s *LLMStreamSession) String() string {
	return fmt.Sprintf("LLMStreamSession{StreamBuf='%s', StreamSendLen=%d, MsgID='%s', DefaultBody=%v}",
		s.StreamBuf, s.StreamSendLen, s.MsgID, s.DefaultBody)
}

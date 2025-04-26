package wafruntime

import (
	"github.com/volcengine/volcengine-go-sdk/service/waf"

	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
)

const opCheckLLMResponseStream = "CheckLLMResponseStream"

func (c *WAFRuntime) CheckLLMResponseStreamRequest(input *waf.CheckLLMResponseStreamInput, session *LLMStreamSession) (req *request.Request, output *waf.CheckLLMResponseStreamOutput) {
	op := &request.Operation{
		Name:       opCheckLLMResponseStream,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &waf.CheckLLMResponseStreamInput{}
	}
	session.StreamBuf += *input.Content
	session.StreamSendLen += len(*input.Content)
	if session.StreamSendLen < 10 && input.MsgID != nil && *input.UseStream != 2 {
		return nil, nil
	} else {
		session.StreamSendLen = 0
		*input.Content = session.StreamBuf
	}
	if session.MsgID != "" {
		input.MsgID = &session.MsgID
	}
	output = &waf.CheckLLMResponseStreamOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")
	//if c.MsgID == nil {
	//	c.MsgID = output.MsgID
	//}
	return
}

func (c *WAFRuntime) CheckLLMResponseStream(input *waf.CheckLLMResponseStreamInput, session *LLMStreamSession) (*waf.CheckLLMResponseStreamOutput, error) {
	req, out := c.CheckLLMResponseStreamRequest(input, session)
	if req == nil {
		return session.DefaultBody, nil
	}
	err := req.Send()
	if err != nil {
		return nil, err
	}
	if session.MsgID == "" {
		session.MsgID = *out.MsgID
	}
	session.DefaultBody = out
	return out, err
}

func (c *WAFRuntime) CheckLLMResponseStreamWithContext(ctx volcengine.Context, input *waf.CheckLLMResponseStreamInput, session *LLMStreamSession, opts ...request.Option) (*waf.CheckLLMResponseStreamOutput, error) {
	req, out := c.CheckLLMResponseStreamRequest(input, session)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, c.send(req)
}

package wafruntime

import (
	"github.com/volcengine/volcengine-go-sdk/service/waf"

	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
)

const opCheckLLMResponseStream = "CheckLLMResponseStream"

func (c *WAFRuntime) CheckLLMResponseStreamRequest(input *waf.CheckLLMResponseStreamInput) (req *request.Request, output *waf.CheckLLMResponseStreamOutput) {
	op := &request.Operation{
		Name:       opCheckLLMResponseStream,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &waf.CheckLLMResponseStreamInput{}
	}
	c.StreamBuf += *input.Content
	c.StreamSendLen += len(*input.Content)
	if c.StreamSendLen < 10 && input.MsgID != nil && *input.UseStream != 2 {
		return nil, nil
	} else {
		c.StreamSendLen = 0
		*input.Content = c.StreamBuf
	}
	if c.MsgID != nil {
		input.MsgID = c.MsgID
	}
	output = &waf.CheckLLMResponseStreamOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")
	//if c.MsgID == nil {
	//	c.MsgID = output.MsgID
	//}
	return
}

func (c *WAFRuntime) CheckLLMResponseStream(input *waf.CheckLLMResponseStreamInput) (*waf.CheckLLMResponseStreamOutput, error) {
	req, out := c.CheckLLMResponseStreamRequest(input)
	if req == nil {
		return c.defaultOut, nil
	}
	err := req.Send()
	if err != nil {
		return nil, err
	}
	if c.MsgID == nil {
		c.MsgID = out.MsgID
	}
	c.defaultOut = out
	return out, err
}

func (c *WAFRuntime) CheckLLMResponseStreamWithContext(ctx volcengine.Context, input *waf.CheckLLMResponseStreamInput, opts ...request.Option) (*waf.CheckLLMResponseStreamOutput, error) {
	req, out := c.CheckLLMResponseStreamRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, c.send(req)
}

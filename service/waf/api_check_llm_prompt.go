// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package waf

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opCheckLLMPromptCommon = "CheckLLMPrompt"

// CheckLLMPromptCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the CheckLLMPromptCommon operation. The "output" return
// value will be populated with the CheckLLMPromptCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CheckLLMPromptCommon Request to send the API call to the service.
// the "output" return value is not valid until after CheckLLMPromptCommon Send returns without error.
//
// See CheckLLMPromptCommon for more information on using the CheckLLMPromptCommon
// API call, and error handling.
//
//    // Example sending a request using the CheckLLMPromptCommonRequest method.
//    req, resp := client.CheckLLMPromptCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *WAF) CheckLLMPromptCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opCheckLLMPromptCommon,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &map[string]interface{}{}
	}

	output = &map[string]interface{}{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// CheckLLMPromptCommon API operation for WAF.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for WAF's
// API operation CheckLLMPromptCommon for usage and error information.
func (c *WAF) CheckLLMPromptCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.CheckLLMPromptCommonRequest(input)
	return out, req.Send()
}

// CheckLLMPromptCommonWithContext is the same as CheckLLMPromptCommon with the addition of
// the ability to pass a context and additional request options.
//
// See CheckLLMPromptCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *WAF) CheckLLMPromptCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.CheckLLMPromptCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opCheckLLMPrompt = "CheckLLMPrompt"

// CheckLLMPromptRequest generates a "volcengine/request.Request" representing the
// client's request for the CheckLLMPrompt operation. The "output" return
// value will be populated with the CheckLLMPromptCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CheckLLMPromptCommon Request to send the API call to the service.
// the "output" return value is not valid until after CheckLLMPromptCommon Send returns without error.
//
// See CheckLLMPrompt for more information on using the CheckLLMPrompt
// API call, and error handling.
//
//    // Example sending a request using the CheckLLMPromptRequest method.
//    req, resp := client.CheckLLMPromptRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *WAF) CheckLLMPromptRequest(input *CheckLLMPromptInput) (req *request.Request, output *CheckLLMPromptOutput) {
	op := &request.Operation{
		Name:       opCheckLLMPrompt,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CheckLLMPromptInput{}
	}

	output = &CheckLLMPromptOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// CheckLLMPrompt API operation for WAF.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for WAF's
// API operation CheckLLMPrompt for usage and error information.
func (c *WAF) CheckLLMPrompt(input *CheckLLMPromptInput) (*CheckLLMPromptOutput, error) {
	req, out := c.CheckLLMPromptRequest(input)
	return out, req.Send()
}

// CheckLLMPromptWithContext is the same as CheckLLMPrompt with the addition of
// the ability to pass a context and additional request options.
//
// See CheckLLMPrompt for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *WAF) CheckLLMPromptWithContext(ctx volcengine.Context, input *CheckLLMPromptInput, opts ...request.Option) (*CheckLLMPromptOutput, error) {
	req, out := c.CheckLLMPromptRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type CheckLLMPromptInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	// Content is a required field
	Content *string `type:"string" json:",omitempty" required:"true"`

	// ContentType is a required field
	ContentType *int32 `type:"int32" json:",omitempty" required:"true"`

	// Host is a required field
	Host *string `type:"string" json:",omitempty" required:"true"`

	MsgClass *int32 `type:"int32" json:",omitempty"`

	// Region is a required field
	Region *string `type:"string" json:",omitempty" required:"true"`
}

// String returns the string representation
func (s CheckLLMPromptInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CheckLLMPromptInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CheckLLMPromptInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "CheckLLMPromptInput"}
	if s.Content == nil {
		invalidParams.Add(request.NewErrParamRequired("Content"))
	}
	if s.ContentType == nil {
		invalidParams.Add(request.NewErrParamRequired("ContentType"))
	}
	if s.Host == nil {
		invalidParams.Add(request.NewErrParamRequired("Host"))
	}
	if s.Region == nil {
		invalidParams.Add(request.NewErrParamRequired("Region"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetContent sets the Content field's value.
func (s *CheckLLMPromptInput) SetContent(v string) *CheckLLMPromptInput {
	s.Content = &v
	return s
}

// SetContentType sets the ContentType field's value.
func (s *CheckLLMPromptInput) SetContentType(v int32) *CheckLLMPromptInput {
	s.ContentType = &v
	return s
}

// SetHost sets the Host field's value.
func (s *CheckLLMPromptInput) SetHost(v string) *CheckLLMPromptInput {
	s.Host = &v
	return s
}

// SetMsgClass sets the MsgClass field's value.
func (s *CheckLLMPromptInput) SetMsgClass(v int32) *CheckLLMPromptInput {
	s.MsgClass = &v
	return s
}

// SetRegion sets the Region field's value.
func (s *CheckLLMPromptInput) SetRegion(v string) *CheckLLMPromptInput {
	s.Region = &v
	return s
}

type CheckLLMPromptOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	Decision *DecisionForCheckLLMPromptOutput `type:"structure" json:",omitempty"`

	MsgID *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s CheckLLMPromptOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CheckLLMPromptOutput) GoString() string {
	return s.String()
}

// SetDecision sets the Decision field's value.
func (s *CheckLLMPromptOutput) SetDecision(v *DecisionForCheckLLMPromptOutput) *CheckLLMPromptOutput {
	s.Decision = v
	return s
}

// SetMsgID sets the MsgID field's value.
func (s *CheckLLMPromptOutput) SetMsgID(v string) *CheckLLMPromptOutput {
	s.MsgID = &v
	return s
}

type CustomMatchForCheckLLMPromptOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Word *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s CustomMatchForCheckLLMPromptOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CustomMatchForCheckLLMPromptOutput) GoString() string {
	return s.String()
}

// SetWord sets the Word field's value.
func (s *CustomMatchForCheckLLMPromptOutput) SetWord(v string) *CustomMatchForCheckLLMPromptOutput {
	s.Word = &v
	return s
}

type DecisionForCheckLLMPromptOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Action *int32 `type:"int32" json:",omitempty"`

	CustomMatches []*CustomMatchForCheckLLMPromptOutput `type:"list" json:",omitempty"`

	ErrCode *int32 `type:"int32" json:",omitempty"`

	ErrMsg *string `type:"string" json:",omitempty"`

	Labels []*string `type:"list" json:",omitempty"`

	Matches []*MatchForCheckLLMPromptOutput `type:"list" json:",omitempty"`
}

// String returns the string representation
func (s DecisionForCheckLLMPromptOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DecisionForCheckLLMPromptOutput) GoString() string {
	return s.String()
}

// SetAction sets the Action field's value.
func (s *DecisionForCheckLLMPromptOutput) SetAction(v int32) *DecisionForCheckLLMPromptOutput {
	s.Action = &v
	return s
}

// SetCustomMatches sets the CustomMatches field's value.
func (s *DecisionForCheckLLMPromptOutput) SetCustomMatches(v []*CustomMatchForCheckLLMPromptOutput) *DecisionForCheckLLMPromptOutput {
	s.CustomMatches = v
	return s
}

// SetErrCode sets the ErrCode field's value.
func (s *DecisionForCheckLLMPromptOutput) SetErrCode(v int32) *DecisionForCheckLLMPromptOutput {
	s.ErrCode = &v
	return s
}

// SetErrMsg sets the ErrMsg field's value.
func (s *DecisionForCheckLLMPromptOutput) SetErrMsg(v string) *DecisionForCheckLLMPromptOutput {
	s.ErrMsg = &v
	return s
}

// SetLabels sets the Labels field's value.
func (s *DecisionForCheckLLMPromptOutput) SetLabels(v []*string) *DecisionForCheckLLMPromptOutput {
	s.Labels = v
	return s
}

// SetMatches sets the Matches field's value.
func (s *DecisionForCheckLLMPromptOutput) SetMatches(v []*MatchForCheckLLMPromptOutput) *DecisionForCheckLLMPromptOutput {
	s.Matches = v
	return s
}

type MatchForCheckLLMPromptOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Label *string `type:"string" json:",omitempty"`

	Word *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s MatchForCheckLLMPromptOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s MatchForCheckLLMPromptOutput) GoString() string {
	return s.String()
}

// SetLabel sets the Label field's value.
func (s *MatchForCheckLLMPromptOutput) SetLabel(v string) *MatchForCheckLLMPromptOutput {
	s.Label = &v
	return s
}

// SetWord sets the Word field's value.
func (s *MatchForCheckLLMPromptOutput) SetWord(v string) *MatchForCheckLLMPromptOutput {
	s.Word = &v
	return s
}

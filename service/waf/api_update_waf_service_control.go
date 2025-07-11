// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package waf

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opUpdateWafServiceControlCommon = "UpdateWafServiceControl"

// UpdateWafServiceControlCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the UpdateWafServiceControlCommon operation. The "output" return
// value will be populated with the UpdateWafServiceControlCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned UpdateWafServiceControlCommon Request to send the API call to the service.
// the "output" return value is not valid until after UpdateWafServiceControlCommon Send returns without error.
//
// See UpdateWafServiceControlCommon for more information on using the UpdateWafServiceControlCommon
// API call, and error handling.
//
//    // Example sending a request using the UpdateWafServiceControlCommonRequest method.
//    req, resp := client.UpdateWafServiceControlCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *WAF) UpdateWafServiceControlCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opUpdateWafServiceControlCommon,
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

// UpdateWafServiceControlCommon API operation for WAF.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for WAF's
// API operation UpdateWafServiceControlCommon for usage and error information.
func (c *WAF) UpdateWafServiceControlCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.UpdateWafServiceControlCommonRequest(input)
	return out, req.Send()
}

// UpdateWafServiceControlCommonWithContext is the same as UpdateWafServiceControlCommon with the addition of
// the ability to pass a context and additional request options.
//
// See UpdateWafServiceControlCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *WAF) UpdateWafServiceControlCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.UpdateWafServiceControlCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opUpdateWafServiceControl = "UpdateWafServiceControl"

// UpdateWafServiceControlRequest generates a "volcengine/request.Request" representing the
// client's request for the UpdateWafServiceControl operation. The "output" return
// value will be populated with the UpdateWafServiceControlCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned UpdateWafServiceControlCommon Request to send the API call to the service.
// the "output" return value is not valid until after UpdateWafServiceControlCommon Send returns without error.
//
// See UpdateWafServiceControl for more information on using the UpdateWafServiceControl
// API call, and error handling.
//
//    // Example sending a request using the UpdateWafServiceControlRequest method.
//    req, resp := client.UpdateWafServiceControlRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *WAF) UpdateWafServiceControlRequest(input *UpdateWafServiceControlInput) (req *request.Request, output *UpdateWafServiceControlOutput) {
	op := &request.Operation{
		Name:       opUpdateWafServiceControl,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateWafServiceControlInput{}
	}

	output = &UpdateWafServiceControlOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// UpdateWafServiceControl API operation for WAF.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for WAF's
// API operation UpdateWafServiceControl for usage and error information.
func (c *WAF) UpdateWafServiceControl(input *UpdateWafServiceControlInput) (*UpdateWafServiceControlOutput, error) {
	req, out := c.UpdateWafServiceControlRequest(input)
	return out, req.Send()
}

// UpdateWafServiceControlWithContext is the same as UpdateWafServiceControl with the addition of
// the ability to pass a context and additional request options.
//
// See UpdateWafServiceControl for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *WAF) UpdateWafServiceControlWithContext(ctx volcengine.Context, input *UpdateWafServiceControlInput, opts ...request.Option) (*UpdateWafServiceControlOutput, error) {
	req, out := c.UpdateWafServiceControlRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type UpdateWafServiceControlInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	ApiEnable *int32 `type:"int32" json:",omitempty"`

	AutoCCEnable *int32 `type:"int32" json:",omitempty"`

	BlackIpEnable *int32 `type:"int32" json:",omitempty"`

	BlackLctEnable *int32 `type:"int32" json:",omitempty"`

	BotDytokenEnable *int32 `type:"int32" json:",omitempty"`

	BotFrequencyEnable *int32 `type:"int32" json:",omitempty"`

	BotRepeatEnable *int32 `type:"int32" json:",omitempty"`

	BotSequenceDefaultAction *int32 `type:"int32" json:",omitempty"`

	BotSequenceEnable *int32 `type:"int32" json:",omitempty"`

	CcEnable *int32 `type:"int32" json:",omitempty"`

	CustomBotEnable *int32 `type:"int32" json:",omitempty"`

	CustomRspEnable *int32 `type:"int32" json:",omitempty"`

	DlpEnable *int32 `type:"int32" json:",omitempty"`

	// Host is a required field
	Host *string `type:"string" json:",omitempty" required:"true"`

	ProjectName *string `type:"string" json:",omitempty"`

	SystemBotEnable *int32 `type:"int32" json:",omitempty"`

	TLSEnable *int32 `type:"int32" json:",omitempty"`

	TamperProofEnable *int32 `type:"int32" json:",omitempty"`

	WafEnable *int32 `type:"int32" json:",omitempty"`

	WafWhiteReqEnable *int32 `type:"int32" json:",omitempty"`

	WhiteEnable *int32 `type:"int32" json:",omitempty"`

	WhiteFieldEnable *int32 `type:"int32" json:",omitempty"`
}

// String returns the string representation
func (s UpdateWafServiceControlInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s UpdateWafServiceControlInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateWafServiceControlInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "UpdateWafServiceControlInput"}
	if s.Host == nil {
		invalidParams.Add(request.NewErrParamRequired("Host"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetApiEnable sets the ApiEnable field's value.
func (s *UpdateWafServiceControlInput) SetApiEnable(v int32) *UpdateWafServiceControlInput {
	s.ApiEnable = &v
	return s
}

// SetAutoCCEnable sets the AutoCCEnable field's value.
func (s *UpdateWafServiceControlInput) SetAutoCCEnable(v int32) *UpdateWafServiceControlInput {
	s.AutoCCEnable = &v
	return s
}

// SetBlackIpEnable sets the BlackIpEnable field's value.
func (s *UpdateWafServiceControlInput) SetBlackIpEnable(v int32) *UpdateWafServiceControlInput {
	s.BlackIpEnable = &v
	return s
}

// SetBlackLctEnable sets the BlackLctEnable field's value.
func (s *UpdateWafServiceControlInput) SetBlackLctEnable(v int32) *UpdateWafServiceControlInput {
	s.BlackLctEnable = &v
	return s
}

// SetBotDytokenEnable sets the BotDytokenEnable field's value.
func (s *UpdateWafServiceControlInput) SetBotDytokenEnable(v int32) *UpdateWafServiceControlInput {
	s.BotDytokenEnable = &v
	return s
}

// SetBotFrequencyEnable sets the BotFrequencyEnable field's value.
func (s *UpdateWafServiceControlInput) SetBotFrequencyEnable(v int32) *UpdateWafServiceControlInput {
	s.BotFrequencyEnable = &v
	return s
}

// SetBotRepeatEnable sets the BotRepeatEnable field's value.
func (s *UpdateWafServiceControlInput) SetBotRepeatEnable(v int32) *UpdateWafServiceControlInput {
	s.BotRepeatEnable = &v
	return s
}

// SetBotSequenceDefaultAction sets the BotSequenceDefaultAction field's value.
func (s *UpdateWafServiceControlInput) SetBotSequenceDefaultAction(v int32) *UpdateWafServiceControlInput {
	s.BotSequenceDefaultAction = &v
	return s
}

// SetBotSequenceEnable sets the BotSequenceEnable field's value.
func (s *UpdateWafServiceControlInput) SetBotSequenceEnable(v int32) *UpdateWafServiceControlInput {
	s.BotSequenceEnable = &v
	return s
}

// SetCcEnable sets the CcEnable field's value.
func (s *UpdateWafServiceControlInput) SetCcEnable(v int32) *UpdateWafServiceControlInput {
	s.CcEnable = &v
	return s
}

// SetCustomBotEnable sets the CustomBotEnable field's value.
func (s *UpdateWafServiceControlInput) SetCustomBotEnable(v int32) *UpdateWafServiceControlInput {
	s.CustomBotEnable = &v
	return s
}

// SetCustomRspEnable sets the CustomRspEnable field's value.
func (s *UpdateWafServiceControlInput) SetCustomRspEnable(v int32) *UpdateWafServiceControlInput {
	s.CustomRspEnable = &v
	return s
}

// SetDlpEnable sets the DlpEnable field's value.
func (s *UpdateWafServiceControlInput) SetDlpEnable(v int32) *UpdateWafServiceControlInput {
	s.DlpEnable = &v
	return s
}

// SetHost sets the Host field's value.
func (s *UpdateWafServiceControlInput) SetHost(v string) *UpdateWafServiceControlInput {
	s.Host = &v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *UpdateWafServiceControlInput) SetProjectName(v string) *UpdateWafServiceControlInput {
	s.ProjectName = &v
	return s
}

// SetSystemBotEnable sets the SystemBotEnable field's value.
func (s *UpdateWafServiceControlInput) SetSystemBotEnable(v int32) *UpdateWafServiceControlInput {
	s.SystemBotEnable = &v
	return s
}

// SetTLSEnable sets the TLSEnable field's value.
func (s *UpdateWafServiceControlInput) SetTLSEnable(v int32) *UpdateWafServiceControlInput {
	s.TLSEnable = &v
	return s
}

// SetTamperProofEnable sets the TamperProofEnable field's value.
func (s *UpdateWafServiceControlInput) SetTamperProofEnable(v int32) *UpdateWafServiceControlInput {
	s.TamperProofEnable = &v
	return s
}

// SetWafEnable sets the WafEnable field's value.
func (s *UpdateWafServiceControlInput) SetWafEnable(v int32) *UpdateWafServiceControlInput {
	s.WafEnable = &v
	return s
}

// SetWafWhiteReqEnable sets the WafWhiteReqEnable field's value.
func (s *UpdateWafServiceControlInput) SetWafWhiteReqEnable(v int32) *UpdateWafServiceControlInput {
	s.WafWhiteReqEnable = &v
	return s
}

// SetWhiteEnable sets the WhiteEnable field's value.
func (s *UpdateWafServiceControlInput) SetWhiteEnable(v int32) *UpdateWafServiceControlInput {
	s.WhiteEnable = &v
	return s
}

// SetWhiteFieldEnable sets the WhiteFieldEnable field's value.
func (s *UpdateWafServiceControlInput) SetWhiteFieldEnable(v int32) *UpdateWafServiceControlInput {
	s.WhiteFieldEnable = &v
	return s
}

type UpdateWafServiceControlOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s UpdateWafServiceControlOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s UpdateWafServiceControlOutput) GoString() string {
	return s.String()
}

// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package secagent

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opCheckIMConfigParamsCommon = "CheckIMConfigParams"

// CheckIMConfigParamsCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the CheckIMConfigParamsCommon operation. The "output" return
// value will be populated with the CheckIMConfigParamsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CheckIMConfigParamsCommon Request to send the API call to the service.
// the "output" return value is not valid until after CheckIMConfigParamsCommon Send returns without error.
//
// See CheckIMConfigParamsCommon for more information on using the CheckIMConfigParamsCommon
// API call, and error handling.
//
//    // Example sending a request using the CheckIMConfigParamsCommonRequest method.
//    req, resp := client.CheckIMConfigParamsCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *SECAGENT) CheckIMConfigParamsCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opCheckIMConfigParamsCommon,
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

// CheckIMConfigParamsCommon API operation for SEC_AGENT.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for SEC_AGENT's
// API operation CheckIMConfigParamsCommon for usage and error information.
func (c *SECAGENT) CheckIMConfigParamsCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.CheckIMConfigParamsCommonRequest(input)
	return out, req.Send()
}

// CheckIMConfigParamsCommonWithContext is the same as CheckIMConfigParamsCommon with the addition of
// the ability to pass a context and additional request options.
//
// See CheckIMConfigParamsCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *SECAGENT) CheckIMConfigParamsCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.CheckIMConfigParamsCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opCheckIMConfigParams = "CheckIMConfigParams"

// CheckIMConfigParamsRequest generates a "volcengine/request.Request" representing the
// client's request for the CheckIMConfigParams operation. The "output" return
// value will be populated with the CheckIMConfigParamsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CheckIMConfigParamsCommon Request to send the API call to the service.
// the "output" return value is not valid until after CheckIMConfigParamsCommon Send returns without error.
//
// See CheckIMConfigParams for more information on using the CheckIMConfigParams
// API call, and error handling.
//
//    // Example sending a request using the CheckIMConfigParamsRequest method.
//    req, resp := client.CheckIMConfigParamsRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *SECAGENT) CheckIMConfigParamsRequest(input *CheckIMConfigParamsInput) (req *request.Request, output *CheckIMConfigParamsOutput) {
	op := &request.Operation{
		Name:       opCheckIMConfigParams,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CheckIMConfigParamsInput{}
	}

	output = &CheckIMConfigParamsOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// CheckIMConfigParams API operation for SEC_AGENT.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for SEC_AGENT's
// API operation CheckIMConfigParams for usage and error information.
func (c *SECAGENT) CheckIMConfigParams(input *CheckIMConfigParamsInput) (*CheckIMConfigParamsOutput, error) {
	req, out := c.CheckIMConfigParamsRequest(input)
	return out, req.Send()
}

// CheckIMConfigParamsWithContext is the same as CheckIMConfigParams with the addition of
// the ability to pass a context and additional request options.
//
// See CheckIMConfigParams for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *SECAGENT) CheckIMConfigParamsWithContext(ctx volcengine.Context, input *CheckIMConfigParamsInput, opts ...request.Option) (*CheckIMConfigParamsOutput, error) {
	req, out := c.CheckIMConfigParamsRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type CheckIMConfigParamsInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	// NotifyPlatform is a required field
	NotifyPlatform *string `type:"string" json:",omitempty" required:"true"`

	// Signature is a required field
	Signature *string `type:"string" json:",omitempty" required:"true"`

	// Webhook is a required field
	Webhook *string `type:"string" json:",omitempty" required:"true"`
}

// String returns the string representation
func (s CheckIMConfigParamsInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CheckIMConfigParamsInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CheckIMConfigParamsInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "CheckIMConfigParamsInput"}
	if s.NotifyPlatform == nil {
		invalidParams.Add(request.NewErrParamRequired("NotifyPlatform"))
	}
	if s.Signature == nil {
		invalidParams.Add(request.NewErrParamRequired("Signature"))
	}
	if s.Webhook == nil {
		invalidParams.Add(request.NewErrParamRequired("Webhook"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetNotifyPlatform sets the NotifyPlatform field's value.
func (s *CheckIMConfigParamsInput) SetNotifyPlatform(v string) *CheckIMConfigParamsInput {
	s.NotifyPlatform = &v
	return s
}

// SetSignature sets the Signature field's value.
func (s *CheckIMConfigParamsInput) SetSignature(v string) *CheckIMConfigParamsInput {
	s.Signature = &v
	return s
}

// SetWebhook sets the Webhook field's value.
func (s *CheckIMConfigParamsInput) SetWebhook(v string) *CheckIMConfigParamsInput {
	s.Webhook = &v
	return s
}

type CheckIMConfigParamsOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	IsPass *bool `type:"boolean" json:",omitempty"`

	Message *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s CheckIMConfigParamsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CheckIMConfigParamsOutput) GoString() string {
	return s.String()
}

// SetIsPass sets the IsPass field's value.
func (s *CheckIMConfigParamsOutput) SetIsPass(v bool) *CheckIMConfigParamsOutput {
	s.IsPass = &v
	return s
}

// SetMessage sets the Message field's value.
func (s *CheckIMConfigParamsOutput) SetMessage(v string) *CheckIMConfigParamsOutput {
	s.Message = &v
	return s
}

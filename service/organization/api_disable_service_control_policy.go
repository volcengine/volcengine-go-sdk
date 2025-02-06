// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package organization

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDisableServiceControlPolicyCommon = "DisableServiceControlPolicy"

// DisableServiceControlPolicyCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DisableServiceControlPolicyCommon operation. The "output" return
// value will be populated with the DisableServiceControlPolicyCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DisableServiceControlPolicyCommon Request to send the API call to the service.
// the "output" return value is not valid until after DisableServiceControlPolicyCommon Send returns without error.
//
// See DisableServiceControlPolicyCommon for more information on using the DisableServiceControlPolicyCommon
// API call, and error handling.
//
//    // Example sending a request using the DisableServiceControlPolicyCommonRequest method.
//    req, resp := client.DisableServiceControlPolicyCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ORGANIZATION) DisableServiceControlPolicyCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDisableServiceControlPolicyCommon,
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

// DisableServiceControlPolicyCommon API operation for ORGANIZATION.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ORGANIZATION's
// API operation DisableServiceControlPolicyCommon for usage and error information.
func (c *ORGANIZATION) DisableServiceControlPolicyCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DisableServiceControlPolicyCommonRequest(input)
	return out, req.Send()
}

// DisableServiceControlPolicyCommonWithContext is the same as DisableServiceControlPolicyCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DisableServiceControlPolicyCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ORGANIZATION) DisableServiceControlPolicyCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DisableServiceControlPolicyCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDisableServiceControlPolicy = "DisableServiceControlPolicy"

// DisableServiceControlPolicyRequest generates a "volcengine/request.Request" representing the
// client's request for the DisableServiceControlPolicy operation. The "output" return
// value will be populated with the DisableServiceControlPolicyCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DisableServiceControlPolicyCommon Request to send the API call to the service.
// the "output" return value is not valid until after DisableServiceControlPolicyCommon Send returns without error.
//
// See DisableServiceControlPolicy for more information on using the DisableServiceControlPolicy
// API call, and error handling.
//
//    // Example sending a request using the DisableServiceControlPolicyRequest method.
//    req, resp := client.DisableServiceControlPolicyRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ORGANIZATION) DisableServiceControlPolicyRequest(input *DisableServiceControlPolicyInput) (req *request.Request, output *DisableServiceControlPolicyOutput) {
	op := &request.Operation{
		Name:       opDisableServiceControlPolicy,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DisableServiceControlPolicyInput{}
	}

	output = &DisableServiceControlPolicyOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// DisableServiceControlPolicy API operation for ORGANIZATION.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ORGANIZATION's
// API operation DisableServiceControlPolicy for usage and error information.
func (c *ORGANIZATION) DisableServiceControlPolicy(input *DisableServiceControlPolicyInput) (*DisableServiceControlPolicyOutput, error) {
	req, out := c.DisableServiceControlPolicyRequest(input)
	return out, req.Send()
}

// DisableServiceControlPolicyWithContext is the same as DisableServiceControlPolicy with the addition of
// the ability to pass a context and additional request options.
//
// See DisableServiceControlPolicy for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ORGANIZATION) DisableServiceControlPolicyWithContext(ctx volcengine.Context, input *DisableServiceControlPolicyInput, opts ...request.Option) (*DisableServiceControlPolicyOutput, error) {
	req, out := c.DisableServiceControlPolicyRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DisableServiceControlPolicyInput struct {
	_ struct{} `type:"structure" json:",omitempty"`
}

// String returns the string representation
func (s DisableServiceControlPolicyInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DisableServiceControlPolicyInput) GoString() string {
	return s.String()
}

type DisableServiceControlPolicyOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	Status *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s DisableServiceControlPolicyOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DisableServiceControlPolicyOutput) GoString() string {
	return s.String()
}

// SetStatus sets the Status field's value.
func (s *DisableServiceControlPolicyOutput) SetStatus(v string) *DisableServiceControlPolicyOutput {
	s.Status = &v
	return s
}

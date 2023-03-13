// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package alb

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDisableTLSAccessLogCommon = "DisableTLSAccessLog"

// DisableTLSAccessLogCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DisableTLSAccessLogCommon operation. The "output" return
// value will be populated with the DisableTLSAccessLogCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DisableTLSAccessLogCommon Request to send the API call to the service.
// the "output" return value is not valid until after DisableTLSAccessLogCommon Send returns without error.
//
// See DisableTLSAccessLogCommon for more information on using the DisableTLSAccessLogCommon
// API call, and error handling.
//
//    // Example sending a request using the DisableTLSAccessLogCommonRequest method.
//    req, resp := client.DisableTLSAccessLogCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ALB) DisableTLSAccessLogCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDisableTLSAccessLogCommon,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &map[string]interface{}{}
	}

	output = &map[string]interface{}{}
	req = c.newRequest(op, input, output)

	return
}

// DisableTLSAccessLogCommon API operation for ALB.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ALB's
// API operation DisableTLSAccessLogCommon for usage and error information.
func (c *ALB) DisableTLSAccessLogCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DisableTLSAccessLogCommonRequest(input)
	return out, req.Send()
}

// DisableTLSAccessLogCommonWithContext is the same as DisableTLSAccessLogCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DisableTLSAccessLogCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ALB) DisableTLSAccessLogCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DisableTLSAccessLogCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDisableTLSAccessLog = "DisableTLSAccessLog"

// DisableTLSAccessLogRequest generates a "volcengine/request.Request" representing the
// client's request for the DisableTLSAccessLog operation. The "output" return
// value will be populated with the DisableTLSAccessLogCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DisableTLSAccessLogCommon Request to send the API call to the service.
// the "output" return value is not valid until after DisableTLSAccessLogCommon Send returns without error.
//
// See DisableTLSAccessLog for more information on using the DisableTLSAccessLog
// API call, and error handling.
//
//    // Example sending a request using the DisableTLSAccessLogRequest method.
//    req, resp := client.DisableTLSAccessLogRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ALB) DisableTLSAccessLogRequest(input *DisableTLSAccessLogInput) (req *request.Request, output *DisableTLSAccessLogOutput) {
	op := &request.Operation{
		Name:       opDisableTLSAccessLog,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DisableTLSAccessLogInput{}
	}

	output = &DisableTLSAccessLogOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DisableTLSAccessLog API operation for ALB.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ALB's
// API operation DisableTLSAccessLog for usage and error information.
func (c *ALB) DisableTLSAccessLog(input *DisableTLSAccessLogInput) (*DisableTLSAccessLogOutput, error) {
	req, out := c.DisableTLSAccessLogRequest(input)
	return out, req.Send()
}

// DisableTLSAccessLogWithContext is the same as DisableTLSAccessLog with the addition of
// the ability to pass a context and additional request options.
//
// See DisableTLSAccessLog for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ALB) DisableTLSAccessLogWithContext(ctx volcengine.Context, input *DisableTLSAccessLogInput, opts ...request.Option) (*DisableTLSAccessLogOutput, error) {
	req, out := c.DisableTLSAccessLogRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DisableTLSAccessLogInput struct {
	_ struct{} `type:"structure"`

	// LoadBalancerId is a required field
	LoadBalancerId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DisableTLSAccessLogInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DisableTLSAccessLogInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DisableTLSAccessLogInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DisableTLSAccessLogInput"}
	if s.LoadBalancerId == nil {
		invalidParams.Add(request.NewErrParamRequired("LoadBalancerId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetLoadBalancerId sets the LoadBalancerId field's value.
func (s *DisableTLSAccessLogInput) SetLoadBalancerId(v string) *DisableTLSAccessLogInput {
	s.LoadBalancerId = &v
	return s
}

type DisableTLSAccessLogOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	RequestId *string `type:"string"`
}

// String returns the string representation
func (s DisableTLSAccessLogOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DisableTLSAccessLogOutput) GoString() string {
	return s.String()
}

// SetRequestId sets the RequestId field's value.
func (s *DisableTLSAccessLogOutput) SetRequestId(v string) *DisableTLSAccessLogOutput {
	s.RequestId = &v
	return s
}
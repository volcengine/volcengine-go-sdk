// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package autoscaling

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opEnableScalingPolicyCommon = "EnableScalingPolicy"

// EnableScalingPolicyCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the EnableScalingPolicyCommon operation. The "output" return
// value will be populated with the EnableScalingPolicyCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned EnableScalingPolicyCommon Request to send the API call to the service.
// the "output" return value is not valid until after EnableScalingPolicyCommon Send returns without error.
//
// See EnableScalingPolicyCommon for more information on using the EnableScalingPolicyCommon
// API call, and error handling.
//
//    // Example sending a request using the EnableScalingPolicyCommonRequest method.
//    req, resp := client.EnableScalingPolicyCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *AUTOSCALING) EnableScalingPolicyCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opEnableScalingPolicyCommon,
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

// EnableScalingPolicyCommon API operation for AUTO_SCALING.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for AUTO_SCALING's
// API operation EnableScalingPolicyCommon for usage and error information.
func (c *AUTOSCALING) EnableScalingPolicyCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.EnableScalingPolicyCommonRequest(input)
	return out, req.Send()
}

// EnableScalingPolicyCommonWithContext is the same as EnableScalingPolicyCommon with the addition of
// the ability to pass a context and additional request options.
//
// See EnableScalingPolicyCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *AUTOSCALING) EnableScalingPolicyCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.EnableScalingPolicyCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opEnableScalingPolicy = "EnableScalingPolicy"

// EnableScalingPolicyRequest generates a "volcengine/request.Request" representing the
// client's request for the EnableScalingPolicy operation. The "output" return
// value will be populated with the EnableScalingPolicyCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned EnableScalingPolicyCommon Request to send the API call to the service.
// the "output" return value is not valid until after EnableScalingPolicyCommon Send returns without error.
//
// See EnableScalingPolicy for more information on using the EnableScalingPolicy
// API call, and error handling.
//
//    // Example sending a request using the EnableScalingPolicyRequest method.
//    req, resp := client.EnableScalingPolicyRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *AUTOSCALING) EnableScalingPolicyRequest(input *EnableScalingPolicyInput) (req *request.Request, output *EnableScalingPolicyOutput) {
	op := &request.Operation{
		Name:       opEnableScalingPolicy,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &EnableScalingPolicyInput{}
	}

	output = &EnableScalingPolicyOutput{}
	req = c.newRequest(op, input, output)

	return
}

// EnableScalingPolicy API operation for AUTO_SCALING.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for AUTO_SCALING's
// API operation EnableScalingPolicy for usage and error information.
func (c *AUTOSCALING) EnableScalingPolicy(input *EnableScalingPolicyInput) (*EnableScalingPolicyOutput, error) {
	req, out := c.EnableScalingPolicyRequest(input)
	return out, req.Send()
}

// EnableScalingPolicyWithContext is the same as EnableScalingPolicy with the addition of
// the ability to pass a context and additional request options.
//
// See EnableScalingPolicy for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *AUTOSCALING) EnableScalingPolicyWithContext(ctx volcengine.Context, input *EnableScalingPolicyInput, opts ...request.Option) (*EnableScalingPolicyOutput, error) {
	req, out := c.EnableScalingPolicyRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type EnableScalingPolicyInput struct {
	_ struct{} `type:"structure"`

	ScalingPolicyId *string `type:"string"`
}

// String returns the string representation
func (s EnableScalingPolicyInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s EnableScalingPolicyInput) GoString() string {
	return s.String()
}

// SetScalingPolicyId sets the ScalingPolicyId field's value.
func (s *EnableScalingPolicyInput) SetScalingPolicyId(v string) *EnableScalingPolicyInput {
	s.ScalingPolicyId = &v
	return s
}

type EnableScalingPolicyOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	ScalingPolicyId *string `type:"string"`
}

// String returns the string representation
func (s EnableScalingPolicyOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s EnableScalingPolicyOutput) GoString() string {
	return s.String()
}

// SetScalingPolicyId sets the ScalingPolicyId field's value.
func (s *EnableScalingPolicyOutput) SetScalingPolicyId(v string) *EnableScalingPolicyOutput {
	s.ScalingPolicyId = &v
	return s
}

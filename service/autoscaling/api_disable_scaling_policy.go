// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package autoscaling

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDisableScalingPolicyCommon = "DisableScalingPolicy"

// DisableScalingPolicyCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DisableScalingPolicyCommon operation. The "output" return
// value will be populated with the DisableScalingPolicyCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DisableScalingPolicyCommon Request to send the API call to the service.
// the "output" return value is not valid until after DisableScalingPolicyCommon Send returns without error.
//
// See DisableScalingPolicyCommon for more information on using the DisableScalingPolicyCommon
// API call, and error handling.
//
//    // Example sending a request using the DisableScalingPolicyCommonRequest method.
//    req, resp := client.DisableScalingPolicyCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *AUTOSCALING) DisableScalingPolicyCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDisableScalingPolicyCommon,
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

// DisableScalingPolicyCommon API operation for AUTOSCALING.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for AUTOSCALING's
// API operation DisableScalingPolicyCommon for usage and error information.
func (c *AUTOSCALING) DisableScalingPolicyCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DisableScalingPolicyCommonRequest(input)
	return out, req.Send()
}

// DisableScalingPolicyCommonWithContext is the same as DisableScalingPolicyCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DisableScalingPolicyCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *AUTOSCALING) DisableScalingPolicyCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DisableScalingPolicyCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDisableScalingPolicy = "DisableScalingPolicy"

// DisableScalingPolicyRequest generates a "volcengine/request.Request" representing the
// client's request for the DisableScalingPolicy operation. The "output" return
// value will be populated with the DisableScalingPolicyCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DisableScalingPolicyCommon Request to send the API call to the service.
// the "output" return value is not valid until after DisableScalingPolicyCommon Send returns without error.
//
// See DisableScalingPolicy for more information on using the DisableScalingPolicy
// API call, and error handling.
//
//    // Example sending a request using the DisableScalingPolicyRequest method.
//    req, resp := client.DisableScalingPolicyRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *AUTOSCALING) DisableScalingPolicyRequest(input *DisableScalingPolicyInput) (req *request.Request, output *DisableScalingPolicyOutput) {
	op := &request.Operation{
		Name:       opDisableScalingPolicy,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DisableScalingPolicyInput{}
	}

	output = &DisableScalingPolicyOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DisableScalingPolicy API operation for AUTOSCALING.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for AUTOSCALING's
// API operation DisableScalingPolicy for usage and error information.
func (c *AUTOSCALING) DisableScalingPolicy(input *DisableScalingPolicyInput) (*DisableScalingPolicyOutput, error) {
	req, out := c.DisableScalingPolicyRequest(input)
	return out, req.Send()
}

// DisableScalingPolicyWithContext is the same as DisableScalingPolicy with the addition of
// the ability to pass a context and additional request options.
//
// See DisableScalingPolicy for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *AUTOSCALING) DisableScalingPolicyWithContext(ctx volcengine.Context, input *DisableScalingPolicyInput, opts ...request.Option) (*DisableScalingPolicyOutput, error) {
	req, out := c.DisableScalingPolicyRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DisableScalingPolicyInput struct {
	_ struct{} `type:"structure"`

	// ScalingPolicyId is a required field
	ScalingPolicyId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DisableScalingPolicyInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DisableScalingPolicyInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DisableScalingPolicyInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DisableScalingPolicyInput"}
	if s.ScalingPolicyId == nil {
		invalidParams.Add(request.NewErrParamRequired("ScalingPolicyId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetScalingPolicyId sets the ScalingPolicyId field's value.
func (s *DisableScalingPolicyInput) SetScalingPolicyId(v string) *DisableScalingPolicyInput {
	s.ScalingPolicyId = &v
	return s
}

type DisableScalingPolicyOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	ScalingPolicyId *string `type:"string"`
}

// String returns the string representation
func (s DisableScalingPolicyOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DisableScalingPolicyOutput) GoString() string {
	return s.String()
}

// SetScalingPolicyId sets the ScalingPolicyId field's value.
func (s *DisableScalingPolicyOutput) SetScalingPolicyId(v string) *DisableScalingPolicyOutput {
	s.ScalingPolicyId = &v
	return s
}

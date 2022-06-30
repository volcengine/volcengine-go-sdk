// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package autoscaling

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opAttachInstancesCommon = "AttachInstances"

// AttachInstancesCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the AttachInstancesCommon operation. The "output" return
// value will be populated with the AttachInstancesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned AttachInstancesCommon Request to send the API call to the service.
// the "output" return value is not valid until after AttachInstancesCommon Send returns without error.
//
// See AttachInstancesCommon for more information on using the AttachInstancesCommon
// API call, and error handling.
//
//    // Example sending a request using the AttachInstancesCommonRequest method.
//    req, resp := client.AttachInstancesCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *AUTOSCALING) AttachInstancesCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opAttachInstancesCommon,
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

// AttachInstancesCommon API operation for AUTO_SCALING.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for AUTO_SCALING's
// API operation AttachInstancesCommon for usage and error information.
func (c *AUTOSCALING) AttachInstancesCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.AttachInstancesCommonRequest(input)
	return out, req.Send()
}

// AttachInstancesCommonWithContext is the same as AttachInstancesCommon with the addition of
// the ability to pass a context and additional request options.
//
// See AttachInstancesCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *AUTOSCALING) AttachInstancesCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.AttachInstancesCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opAttachInstances = "AttachInstances"

// AttachInstancesRequest generates a "volcengine/request.Request" representing the
// client's request for the AttachInstances operation. The "output" return
// value will be populated with the AttachInstancesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned AttachInstancesCommon Request to send the API call to the service.
// the "output" return value is not valid until after AttachInstancesCommon Send returns without error.
//
// See AttachInstances for more information on using the AttachInstances
// API call, and error handling.
//
//    // Example sending a request using the AttachInstancesRequest method.
//    req, resp := client.AttachInstancesRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *AUTOSCALING) AttachInstancesRequest(input *AttachInstancesInput) (req *request.Request, output *AttachInstancesOutput) {
	op := &request.Operation{
		Name:       opAttachInstances,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &AttachInstancesInput{}
	}

	output = &AttachInstancesOutput{}
	req = c.newRequest(op, input, output)

	return
}

// AttachInstances API operation for AUTO_SCALING.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for AUTO_SCALING's
// API operation AttachInstances for usage and error information.
func (c *AUTOSCALING) AttachInstances(input *AttachInstancesInput) (*AttachInstancesOutput, error) {
	req, out := c.AttachInstancesRequest(input)
	return out, req.Send()
}

// AttachInstancesWithContext is the same as AttachInstances with the addition of
// the ability to pass a context and additional request options.
//
// See AttachInstances for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *AUTOSCALING) AttachInstancesWithContext(ctx volcengine.Context, input *AttachInstancesInput, opts ...request.Option) (*AttachInstancesOutput, error) {
	req, out := c.AttachInstancesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type AttachInstancesInput struct {
	_ struct{} `type:"structure"`

	Entrusted *bool `type:"boolean"`

	InstanceIds []*string `type:"list"`

	ScalingGroupId *string `type:"string"`
}

// String returns the string representation
func (s AttachInstancesInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s AttachInstancesInput) GoString() string {
	return s.String()
}

// SetEntrusted sets the Entrusted field's value.
func (s *AttachInstancesInput) SetEntrusted(v bool) *AttachInstancesInput {
	s.Entrusted = &v
	return s
}

// SetInstanceIds sets the InstanceIds field's value.
func (s *AttachInstancesInput) SetInstanceIds(v []*string) *AttachInstancesInput {
	s.InstanceIds = v
	return s
}

// SetScalingGroupId sets the ScalingGroupId field's value.
func (s *AttachInstancesInput) SetScalingGroupId(v string) *AttachInstancesInput {
	s.ScalingGroupId = &v
	return s
}

type AttachInstancesOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s AttachInstancesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s AttachInstancesOutput) GoString() string {
	return s.String()
}

// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package autoscaling

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDeleteScalingGroupCommon = "DeleteScalingGroup"

// DeleteScalingGroupCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DeleteScalingGroupCommon operation. The "output" return
// value will be populated with the DeleteScalingGroupCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeleteScalingGroupCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeleteScalingGroupCommon Send returns without error.
//
// See DeleteScalingGroupCommon for more information on using the DeleteScalingGroupCommon
// API call, and error handling.
//
//	// Example sending a request using the DeleteScalingGroupCommonRequest method.
//	req, resp := client.DeleteScalingGroupCommonRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *AUTOSCALING) DeleteScalingGroupCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDeleteScalingGroupCommon,
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

// DeleteScalingGroupCommon API operation for AUTO_SCALING.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for AUTO_SCALING's
// API operation DeleteScalingGroupCommon for usage and error information.
func (c *AUTOSCALING) DeleteScalingGroupCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DeleteScalingGroupCommonRequest(input)
	return out, req.Send()
}

// DeleteScalingGroupCommonWithContext is the same as DeleteScalingGroupCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteScalingGroupCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *AUTOSCALING) DeleteScalingGroupCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DeleteScalingGroupCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDeleteScalingGroup = "DeleteScalingGroup"

// DeleteScalingGroupRequest generates a "volcengine/request.Request" representing the
// client's request for the DeleteScalingGroup operation. The "output" return
// value will be populated with the DeleteScalingGroupCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeleteScalingGroupCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeleteScalingGroupCommon Send returns without error.
//
// See DeleteScalingGroup for more information on using the DeleteScalingGroup
// API call, and error handling.
//
//	// Example sending a request using the DeleteScalingGroupRequest method.
//	req, resp := client.DeleteScalingGroupRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *AUTOSCALING) DeleteScalingGroupRequest(input *DeleteScalingGroupInput) (req *request.Request, output *DeleteScalingGroupOutput) {
	op := &request.Operation{
		Name:       opDeleteScalingGroup,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteScalingGroupInput{}
	}

	output = &DeleteScalingGroupOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DeleteScalingGroup API operation for AUTO_SCALING.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for AUTO_SCALING's
// API operation DeleteScalingGroup for usage and error information.
func (c *AUTOSCALING) DeleteScalingGroup(input *DeleteScalingGroupInput) (*DeleteScalingGroupOutput, error) {
	req, out := c.DeleteScalingGroupRequest(input)
	return out, req.Send()
}

// DeleteScalingGroupWithContext is the same as DeleteScalingGroup with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteScalingGroup for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *AUTOSCALING) DeleteScalingGroupWithContext(ctx volcengine.Context, input *DeleteScalingGroupInput, opts ...request.Option) (*DeleteScalingGroupOutput, error) {
	req, out := c.DeleteScalingGroupRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DeleteScalingGroupInput struct {
	_ struct{} `type:"structure"`

	ScalingGroupId *string `type:"string"`
}

// String returns the string representation
func (s DeleteScalingGroupInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteScalingGroupInput) GoString() string {
	return s.String()
}

// SetScalingGroupId sets the ScalingGroupId field's value.
func (s *DeleteScalingGroupInput) SetScalingGroupId(v string) *DeleteScalingGroupInput {
	s.ScalingGroupId = &v
	return s
}

type DeleteScalingGroupOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	ScalingGroupId *string `type:"string"`
}

// String returns the string representation
func (s DeleteScalingGroupOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteScalingGroupOutput) GoString() string {
	return s.String()
}

// SetScalingGroupId sets the ScalingGroupId field's value.
func (s *DeleteScalingGroupOutput) SetScalingGroupId(v string) *DeleteScalingGroupOutput {
	s.ScalingGroupId = &v
	return s
}

// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package autoscaling

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opModifyScalingGroupCommon = "ModifyScalingGroup"

// ModifyScalingGroupCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ModifyScalingGroupCommon operation. The "output" return
// value will be populated with the ModifyScalingGroupCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyScalingGroupCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyScalingGroupCommon Send returns without error.
//
// See ModifyScalingGroupCommon for more information on using the ModifyScalingGroupCommon
// API call, and error handling.
//
//    // Example sending a request using the ModifyScalingGroupCommonRequest method.
//    req, resp := client.ModifyScalingGroupCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *AUTOSCALING) ModifyScalingGroupCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opModifyScalingGroupCommon,
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

// ModifyScalingGroupCommon API operation for AUTO_SCALING.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for AUTO_SCALING's
// API operation ModifyScalingGroupCommon for usage and error information.
func (c *AUTOSCALING) ModifyScalingGroupCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ModifyScalingGroupCommonRequest(input)
	return out, req.Send()
}

// ModifyScalingGroupCommonWithContext is the same as ModifyScalingGroupCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyScalingGroupCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *AUTOSCALING) ModifyScalingGroupCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ModifyScalingGroupCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opModifyScalingGroup = "ModifyScalingGroup"

// ModifyScalingGroupRequest generates a "volcengine/request.Request" representing the
// client's request for the ModifyScalingGroup operation. The "output" return
// value will be populated with the ModifyScalingGroupCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyScalingGroupCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyScalingGroupCommon Send returns without error.
//
// See ModifyScalingGroup for more information on using the ModifyScalingGroup
// API call, and error handling.
//
//    // Example sending a request using the ModifyScalingGroupRequest method.
//    req, resp := client.ModifyScalingGroupRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *AUTOSCALING) ModifyScalingGroupRequest(input *ModifyScalingGroupInput) (req *request.Request, output *ModifyScalingGroupOutput) {
	op := &request.Operation{
		Name:       opModifyScalingGroup,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ModifyScalingGroupInput{}
	}

	output = &ModifyScalingGroupOutput{}
	req = c.newRequest(op, input, output)

	return
}

// ModifyScalingGroup API operation for AUTO_SCALING.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for AUTO_SCALING's
// API operation ModifyScalingGroup for usage and error information.
func (c *AUTOSCALING) ModifyScalingGroup(input *ModifyScalingGroupInput) (*ModifyScalingGroupOutput, error) {
	req, out := c.ModifyScalingGroupRequest(input)
	return out, req.Send()
}

// ModifyScalingGroupWithContext is the same as ModifyScalingGroup with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyScalingGroup for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *AUTOSCALING) ModifyScalingGroupWithContext(ctx volcengine.Context, input *ModifyScalingGroupInput, opts ...request.Option) (*ModifyScalingGroupOutput, error) {
	req, out := c.ModifyScalingGroupRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ModifyScalingGroupInput struct {
	_ struct{} `type:"structure"`

	ActiveScalingConfigurationId *string `type:"string"`

	DesireInstanceNumber *int32 `type:"int32"`

	InstanceTerminatePolicy *string `type:"string"`

	MaxInstanceNumber *int32 `type:"int32"`

	MinInstanceNumber *int32 `type:"int32"`

	ScalingGroupId *string `type:"string"`

	ScalingGroupName *string `type:"string"`

	SubnetIds []*string `type:"list"`
}

// String returns the string representation
func (s ModifyScalingGroupInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyScalingGroupInput) GoString() string {
	return s.String()
}

// SetActiveScalingConfigurationId sets the ActiveScalingConfigurationId field's value.
func (s *ModifyScalingGroupInput) SetActiveScalingConfigurationId(v string) *ModifyScalingGroupInput {
	s.ActiveScalingConfigurationId = &v
	return s
}

// SetDesireInstanceNumber sets the DesireInstanceNumber field's value.
func (s *ModifyScalingGroupInput) SetDesireInstanceNumber(v int32) *ModifyScalingGroupInput {
	s.DesireInstanceNumber = &v
	return s
}

// SetInstanceTerminatePolicy sets the InstanceTerminatePolicy field's value.
func (s *ModifyScalingGroupInput) SetInstanceTerminatePolicy(v string) *ModifyScalingGroupInput {
	s.InstanceTerminatePolicy = &v
	return s
}

// SetMaxInstanceNumber sets the MaxInstanceNumber field's value.
func (s *ModifyScalingGroupInput) SetMaxInstanceNumber(v int32) *ModifyScalingGroupInput {
	s.MaxInstanceNumber = &v
	return s
}

// SetMinInstanceNumber sets the MinInstanceNumber field's value.
func (s *ModifyScalingGroupInput) SetMinInstanceNumber(v int32) *ModifyScalingGroupInput {
	s.MinInstanceNumber = &v
	return s
}

// SetScalingGroupId sets the ScalingGroupId field's value.
func (s *ModifyScalingGroupInput) SetScalingGroupId(v string) *ModifyScalingGroupInput {
	s.ScalingGroupId = &v
	return s
}

// SetScalingGroupName sets the ScalingGroupName field's value.
func (s *ModifyScalingGroupInput) SetScalingGroupName(v string) *ModifyScalingGroupInput {
	s.ScalingGroupName = &v
	return s
}

// SetSubnetIds sets the SubnetIds field's value.
func (s *ModifyScalingGroupInput) SetSubnetIds(v []*string) *ModifyScalingGroupInput {
	s.SubnetIds = v
	return s
}

type ModifyScalingGroupOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	ScalingGroupId *string `type:"string"`
}

// String returns the string representation
func (s ModifyScalingGroupOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyScalingGroupOutput) GoString() string {
	return s.String()
}

// SetScalingGroupId sets the ScalingGroupId field's value.
func (s *ModifyScalingGroupOutput) SetScalingGroupId(v string) *ModifyScalingGroupOutput {
	s.ScalingGroupId = &v
	return s
}

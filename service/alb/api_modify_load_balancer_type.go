// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package alb

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opModifyLoadBalancerTypeCommon = "ModifyLoadBalancerType"

// ModifyLoadBalancerTypeCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ModifyLoadBalancerTypeCommon operation. The "output" return
// value will be populated with the ModifyLoadBalancerTypeCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyLoadBalancerTypeCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyLoadBalancerTypeCommon Send returns without error.
//
// See ModifyLoadBalancerTypeCommon for more information on using the ModifyLoadBalancerTypeCommon
// API call, and error handling.
//
//    // Example sending a request using the ModifyLoadBalancerTypeCommonRequest method.
//    req, resp := client.ModifyLoadBalancerTypeCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ALB) ModifyLoadBalancerTypeCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opModifyLoadBalancerTypeCommon,
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

// ModifyLoadBalancerTypeCommon API operation for ALB.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ALB's
// API operation ModifyLoadBalancerTypeCommon for usage and error information.
func (c *ALB) ModifyLoadBalancerTypeCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ModifyLoadBalancerTypeCommonRequest(input)
	return out, req.Send()
}

// ModifyLoadBalancerTypeCommonWithContext is the same as ModifyLoadBalancerTypeCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyLoadBalancerTypeCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ALB) ModifyLoadBalancerTypeCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ModifyLoadBalancerTypeCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opModifyLoadBalancerType = "ModifyLoadBalancerType"

// ModifyLoadBalancerTypeRequest generates a "volcengine/request.Request" representing the
// client's request for the ModifyLoadBalancerType operation. The "output" return
// value will be populated with the ModifyLoadBalancerTypeCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyLoadBalancerTypeCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyLoadBalancerTypeCommon Send returns without error.
//
// See ModifyLoadBalancerType for more information on using the ModifyLoadBalancerType
// API call, and error handling.
//
//    // Example sending a request using the ModifyLoadBalancerTypeRequest method.
//    req, resp := client.ModifyLoadBalancerTypeRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ALB) ModifyLoadBalancerTypeRequest(input *ModifyLoadBalancerTypeInput) (req *request.Request, output *ModifyLoadBalancerTypeOutput) {
	op := &request.Operation{
		Name:       opModifyLoadBalancerType,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ModifyLoadBalancerTypeInput{}
	}

	output = &ModifyLoadBalancerTypeOutput{}
	req = c.newRequest(op, input, output)

	return
}

// ModifyLoadBalancerType API operation for ALB.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ALB's
// API operation ModifyLoadBalancerType for usage and error information.
func (c *ALB) ModifyLoadBalancerType(input *ModifyLoadBalancerTypeInput) (*ModifyLoadBalancerTypeOutput, error) {
	req, out := c.ModifyLoadBalancerTypeRequest(input)
	return out, req.Send()
}

// ModifyLoadBalancerTypeWithContext is the same as ModifyLoadBalancerType with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyLoadBalancerType for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ALB) ModifyLoadBalancerTypeWithContext(ctx volcengine.Context, input *ModifyLoadBalancerTypeInput, opts ...request.Option) (*ModifyLoadBalancerTypeOutput, error) {
	req, out := c.ModifyLoadBalancerTypeRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ModifyLoadBalancerTypeInput struct {
	_ struct{} `type:"structure"`

	LoadBalancerId *string `type:"string"`

	Type *string `type:"string"`

	ZoneMappings []*ZoneMappingForModifyLoadBalancerTypeInput `type:"list"`
}

// String returns the string representation
func (s ModifyLoadBalancerTypeInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyLoadBalancerTypeInput) GoString() string {
	return s.String()
}

// SetLoadBalancerId sets the LoadBalancerId field's value.
func (s *ModifyLoadBalancerTypeInput) SetLoadBalancerId(v string) *ModifyLoadBalancerTypeInput {
	s.LoadBalancerId = &v
	return s
}

// SetType sets the Type field's value.
func (s *ModifyLoadBalancerTypeInput) SetType(v string) *ModifyLoadBalancerTypeInput {
	s.Type = &v
	return s
}

// SetZoneMappings sets the ZoneMappings field's value.
func (s *ModifyLoadBalancerTypeInput) SetZoneMappings(v []*ZoneMappingForModifyLoadBalancerTypeInput) *ModifyLoadBalancerTypeInput {
	s.ZoneMappings = v
	return s
}

type ModifyLoadBalancerTypeOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	RequestId *string `type:"string"`
}

// String returns the string representation
func (s ModifyLoadBalancerTypeOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyLoadBalancerTypeOutput) GoString() string {
	return s.String()
}

// SetRequestId sets the RequestId field's value.
func (s *ModifyLoadBalancerTypeOutput) SetRequestId(v string) *ModifyLoadBalancerTypeOutput {
	s.RequestId = &v
	return s
}

type ZoneMappingForModifyLoadBalancerTypeInput struct {
	_ struct{} `type:"structure"`

	AllocationId *string `type:"string"`

	EipType *string `type:"string"`

	PopLocations *string `type:"string"`

	ZoneId *string `type:"string"`
}

// String returns the string representation
func (s ZoneMappingForModifyLoadBalancerTypeInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ZoneMappingForModifyLoadBalancerTypeInput) GoString() string {
	return s.String()
}

// SetAllocationId sets the AllocationId field's value.
func (s *ZoneMappingForModifyLoadBalancerTypeInput) SetAllocationId(v string) *ZoneMappingForModifyLoadBalancerTypeInput {
	s.AllocationId = &v
	return s
}

// SetEipType sets the EipType field's value.
func (s *ZoneMappingForModifyLoadBalancerTypeInput) SetEipType(v string) *ZoneMappingForModifyLoadBalancerTypeInput {
	s.EipType = &v
	return s
}

// SetPopLocations sets the PopLocations field's value.
func (s *ZoneMappingForModifyLoadBalancerTypeInput) SetPopLocations(v string) *ZoneMappingForModifyLoadBalancerTypeInput {
	s.PopLocations = &v
	return s
}

// SetZoneId sets the ZoneId field's value.
func (s *ZoneMappingForModifyLoadBalancerTypeInput) SetZoneId(v string) *ZoneMappingForModifyLoadBalancerTypeInput {
	s.ZoneId = &v
	return s
}

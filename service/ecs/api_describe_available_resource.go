// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package ecs

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribeAvailableResourceCommon = "DescribeAvailableResource"

// DescribeAvailableResourceCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeAvailableResourceCommon operation. The "output" return
// value will be populated with the DescribeAvailableResourceCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeAvailableResourceCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeAvailableResourceCommon Send returns without error.
//
// See DescribeAvailableResourceCommon for more information on using the DescribeAvailableResourceCommon
// API call, and error handling.
//
//    // Example sending a request using the DescribeAvailableResourceCommonRequest method.
//    req, resp := client.DescribeAvailableResourceCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ECS) DescribeAvailableResourceCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeAvailableResourceCommon,
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

// DescribeAvailableResourceCommon API operation for ECS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ECS's
// API operation DescribeAvailableResourceCommon for usage and error information.
func (c *ECS) DescribeAvailableResourceCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeAvailableResourceCommonRequest(input)
	return out, req.Send()
}

// DescribeAvailableResourceCommonWithContext is the same as DescribeAvailableResourceCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeAvailableResourceCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ECS) DescribeAvailableResourceCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeAvailableResourceCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeAvailableResource = "DescribeAvailableResource"

// DescribeAvailableResourceRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeAvailableResource operation. The "output" return
// value will be populated with the DescribeAvailableResourceCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeAvailableResourceCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeAvailableResourceCommon Send returns without error.
//
// See DescribeAvailableResource for more information on using the DescribeAvailableResource
// API call, and error handling.
//
//    // Example sending a request using the DescribeAvailableResourceRequest method.
//    req, resp := client.DescribeAvailableResourceRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ECS) DescribeAvailableResourceRequest(input *DescribeAvailableResourceInput) (req *request.Request, output *DescribeAvailableResourceOutput) {
	op := &request.Operation{
		Name:       opDescribeAvailableResource,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeAvailableResourceInput{}
	}

	output = &DescribeAvailableResourceOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DescribeAvailableResource API operation for ECS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ECS's
// API operation DescribeAvailableResource for usage and error information.
func (c *ECS) DescribeAvailableResource(input *DescribeAvailableResourceInput) (*DescribeAvailableResourceOutput, error) {
	req, out := c.DescribeAvailableResourceRequest(input)
	return out, req.Send()
}

// DescribeAvailableResourceWithContext is the same as DescribeAvailableResource with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeAvailableResource for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ECS) DescribeAvailableResourceWithContext(ctx volcengine.Context, input *DescribeAvailableResourceInput, opts ...request.Option) (*DescribeAvailableResourceOutput, error) {
	req, out := c.DescribeAvailableResourceRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type AvailableResourceForDescribeAvailableResourceOutput struct {
	_ struct{} `type:"structure"`

	SupportedResources []*SupportedResourceForDescribeAvailableResourceOutput `type:"list"`

	Type *string `type:"string"`
}

// String returns the string representation
func (s AvailableResourceForDescribeAvailableResourceOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s AvailableResourceForDescribeAvailableResourceOutput) GoString() string {
	return s.String()
}

// SetSupportedResources sets the SupportedResources field's value.
func (s *AvailableResourceForDescribeAvailableResourceOutput) SetSupportedResources(v []*SupportedResourceForDescribeAvailableResourceOutput) *AvailableResourceForDescribeAvailableResourceOutput {
	s.SupportedResources = v
	return s
}

// SetType sets the Type field's value.
func (s *AvailableResourceForDescribeAvailableResourceOutput) SetType(v string) *AvailableResourceForDescribeAvailableResourceOutput {
	s.Type = &v
	return s
}

type AvailableZoneForDescribeAvailableResourceOutput struct {
	_ struct{} `type:"structure"`

	AvailableResources []*AvailableResourceForDescribeAvailableResourceOutput `type:"list"`

	RegionId *string `type:"string"`

	Status *string `type:"string"`

	ZoneId *string `type:"string"`
}

// String returns the string representation
func (s AvailableZoneForDescribeAvailableResourceOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s AvailableZoneForDescribeAvailableResourceOutput) GoString() string {
	return s.String()
}

// SetAvailableResources sets the AvailableResources field's value.
func (s *AvailableZoneForDescribeAvailableResourceOutput) SetAvailableResources(v []*AvailableResourceForDescribeAvailableResourceOutput) *AvailableZoneForDescribeAvailableResourceOutput {
	s.AvailableResources = v
	return s
}

// SetRegionId sets the RegionId field's value.
func (s *AvailableZoneForDescribeAvailableResourceOutput) SetRegionId(v string) *AvailableZoneForDescribeAvailableResourceOutput {
	s.RegionId = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *AvailableZoneForDescribeAvailableResourceOutput) SetStatus(v string) *AvailableZoneForDescribeAvailableResourceOutput {
	s.Status = &v
	return s
}

// SetZoneId sets the ZoneId field's value.
func (s *AvailableZoneForDescribeAvailableResourceOutput) SetZoneId(v string) *AvailableZoneForDescribeAvailableResourceOutput {
	s.ZoneId = &v
	return s
}

type DescribeAvailableResourceInput struct {
	_ struct{} `type:"structure"`

	DestinationResource *string `type:"string"`

	InstanceType *string `type:"string"`

	InstanceTypeId *string `type:"string"`

	ZoneId *string `type:"string"`
}

// String returns the string representation
func (s DescribeAvailableResourceInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeAvailableResourceInput) GoString() string {
	return s.String()
}

// SetDestinationResource sets the DestinationResource field's value.
func (s *DescribeAvailableResourceInput) SetDestinationResource(v string) *DescribeAvailableResourceInput {
	s.DestinationResource = &v
	return s
}

// SetInstanceType sets the InstanceType field's value.
func (s *DescribeAvailableResourceInput) SetInstanceType(v string) *DescribeAvailableResourceInput {
	s.InstanceType = &v
	return s
}

// SetInstanceTypeId sets the InstanceTypeId field's value.
func (s *DescribeAvailableResourceInput) SetInstanceTypeId(v string) *DescribeAvailableResourceInput {
	s.InstanceTypeId = &v
	return s
}

// SetZoneId sets the ZoneId field's value.
func (s *DescribeAvailableResourceInput) SetZoneId(v string) *DescribeAvailableResourceInput {
	s.ZoneId = &v
	return s
}

type DescribeAvailableResourceOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	AvailableZones []*AvailableZoneForDescribeAvailableResourceOutput `type:"list"`
}

// String returns the string representation
func (s DescribeAvailableResourceOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeAvailableResourceOutput) GoString() string {
	return s.String()
}

// SetAvailableZones sets the AvailableZones field's value.
func (s *DescribeAvailableResourceOutput) SetAvailableZones(v []*AvailableZoneForDescribeAvailableResourceOutput) *DescribeAvailableResourceOutput {
	s.AvailableZones = v
	return s
}

type SupportedResourceForDescribeAvailableResourceOutput struct {
	_ struct{} `type:"structure"`

	Status *string `type:"string"`

	Value *string `type:"string"`
}

// String returns the string representation
func (s SupportedResourceForDescribeAvailableResourceOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s SupportedResourceForDescribeAvailableResourceOutput) GoString() string {
	return s.String()
}

// SetStatus sets the Status field's value.
func (s *SupportedResourceForDescribeAvailableResourceOutput) SetStatus(v string) *SupportedResourceForDescribeAvailableResourceOutput {
	s.Status = &v
	return s
}

// SetValue sets the Value field's value.
func (s *SupportedResourceForDescribeAvailableResourceOutput) SetValue(v string) *SupportedResourceForDescribeAvailableResourceOutput {
	s.Value = &v
	return s
}

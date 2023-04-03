// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package ecs

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opModifyInstanceMaintenanceAttributesCommon = "ModifyInstanceMaintenanceAttributes"

// ModifyInstanceMaintenanceAttributesCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ModifyInstanceMaintenanceAttributesCommon operation. The "output" return
// value will be populated with the ModifyInstanceMaintenanceAttributesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyInstanceMaintenanceAttributesCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyInstanceMaintenanceAttributesCommon Send returns without error.
//
// See ModifyInstanceMaintenanceAttributesCommon for more information on using the ModifyInstanceMaintenanceAttributesCommon
// API call, and error handling.
//
//    // Example sending a request using the ModifyInstanceMaintenanceAttributesCommonRequest method.
//    req, resp := client.ModifyInstanceMaintenanceAttributesCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ECS) ModifyInstanceMaintenanceAttributesCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opModifyInstanceMaintenanceAttributesCommon,
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

// ModifyInstanceMaintenanceAttributesCommon API operation for ECS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ECS's
// API operation ModifyInstanceMaintenanceAttributesCommon for usage and error information.
func (c *ECS) ModifyInstanceMaintenanceAttributesCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ModifyInstanceMaintenanceAttributesCommonRequest(input)
	return out, req.Send()
}

// ModifyInstanceMaintenanceAttributesCommonWithContext is the same as ModifyInstanceMaintenanceAttributesCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyInstanceMaintenanceAttributesCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ECS) ModifyInstanceMaintenanceAttributesCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ModifyInstanceMaintenanceAttributesCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opModifyInstanceMaintenanceAttributes = "ModifyInstanceMaintenanceAttributes"

// ModifyInstanceMaintenanceAttributesRequest generates a "volcengine/request.Request" representing the
// client's request for the ModifyInstanceMaintenanceAttributes operation. The "output" return
// value will be populated with the ModifyInstanceMaintenanceAttributesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyInstanceMaintenanceAttributesCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyInstanceMaintenanceAttributesCommon Send returns without error.
//
// See ModifyInstanceMaintenanceAttributes for more information on using the ModifyInstanceMaintenanceAttributes
// API call, and error handling.
//
//    // Example sending a request using the ModifyInstanceMaintenanceAttributesRequest method.
//    req, resp := client.ModifyInstanceMaintenanceAttributesRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ECS) ModifyInstanceMaintenanceAttributesRequest(input *ModifyInstanceMaintenanceAttributesInput) (req *request.Request, output *ModifyInstanceMaintenanceAttributesOutput) {
	op := &request.Operation{
		Name:       opModifyInstanceMaintenanceAttributes,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ModifyInstanceMaintenanceAttributesInput{}
	}

	output = &ModifyInstanceMaintenanceAttributesOutput{}
	req = c.newRequest(op, input, output)

	return
}

// ModifyInstanceMaintenanceAttributes API operation for ECS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ECS's
// API operation ModifyInstanceMaintenanceAttributes for usage and error information.
func (c *ECS) ModifyInstanceMaintenanceAttributes(input *ModifyInstanceMaintenanceAttributesInput) (*ModifyInstanceMaintenanceAttributesOutput, error) {
	req, out := c.ModifyInstanceMaintenanceAttributesRequest(input)
	return out, req.Send()
}

// ModifyInstanceMaintenanceAttributesWithContext is the same as ModifyInstanceMaintenanceAttributes with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyInstanceMaintenanceAttributes for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ECS) ModifyInstanceMaintenanceAttributesWithContext(ctx volcengine.Context, input *ModifyInstanceMaintenanceAttributesInput, opts ...request.Option) (*ModifyInstanceMaintenanceAttributesOutput, error) {
	req, out := c.ModifyInstanceMaintenanceAttributesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type MaintenanceAttributeAndActionForModifyInstanceMaintenanceAttributesInput struct {
	_ struct{} `type:"structure"`

	DefaultAction *string `type:"string"`

	MaintenanceType *string `type:"string"`
}

// String returns the string representation
func (s MaintenanceAttributeAndActionForModifyInstanceMaintenanceAttributesInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s MaintenanceAttributeAndActionForModifyInstanceMaintenanceAttributesInput) GoString() string {
	return s.String()
}

// SetDefaultAction sets the DefaultAction field's value.
func (s *MaintenanceAttributeAndActionForModifyInstanceMaintenanceAttributesInput) SetDefaultAction(v string) *MaintenanceAttributeAndActionForModifyInstanceMaintenanceAttributesInput {
	s.DefaultAction = &v
	return s
}

// SetMaintenanceType sets the MaintenanceType field's value.
func (s *MaintenanceAttributeAndActionForModifyInstanceMaintenanceAttributesInput) SetMaintenanceType(v string) *MaintenanceAttributeAndActionForModifyInstanceMaintenanceAttributesInput {
	s.MaintenanceType = &v
	return s
}

type ModifyInstanceMaintenanceAttributesInput struct {
	_ struct{} `type:"structure"`

	InstanceId *string `type:"string"`

	MaintenanceAttributeAndAction []*MaintenanceAttributeAndActionForModifyInstanceMaintenanceAttributesInput `type:"list"`
}

// String returns the string representation
func (s ModifyInstanceMaintenanceAttributesInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyInstanceMaintenanceAttributesInput) GoString() string {
	return s.String()
}

// SetInstanceId sets the InstanceId field's value.
func (s *ModifyInstanceMaintenanceAttributesInput) SetInstanceId(v string) *ModifyInstanceMaintenanceAttributesInput {
	s.InstanceId = &v
	return s
}

// SetMaintenanceAttributeAndAction sets the MaintenanceAttributeAndAction field's value.
func (s *ModifyInstanceMaintenanceAttributesInput) SetMaintenanceAttributeAndAction(v []*MaintenanceAttributeAndActionForModifyInstanceMaintenanceAttributesInput) *ModifyInstanceMaintenanceAttributesInput {
	s.MaintenanceAttributeAndAction = v
	return s
}

type ModifyInstanceMaintenanceAttributesOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	InstanceId *string `type:"string"`
}

// String returns the string representation
func (s ModifyInstanceMaintenanceAttributesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyInstanceMaintenanceAttributesOutput) GoString() string {
	return s.String()
}

// SetInstanceId sets the InstanceId field's value.
func (s *ModifyInstanceMaintenanceAttributesOutput) SetInstanceId(v string) *ModifyInstanceMaintenanceAttributesOutput {
	s.InstanceId = &v
	return s
}

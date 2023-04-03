// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package ecs

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribeInstanceAutoRenewAttributesCommon = "DescribeInstanceAutoRenewAttributes"

// DescribeInstanceAutoRenewAttributesCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeInstanceAutoRenewAttributesCommon operation. The "output" return
// value will be populated with the DescribeInstanceAutoRenewAttributesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeInstanceAutoRenewAttributesCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeInstanceAutoRenewAttributesCommon Send returns without error.
//
// See DescribeInstanceAutoRenewAttributesCommon for more information on using the DescribeInstanceAutoRenewAttributesCommon
// API call, and error handling.
//
//    // Example sending a request using the DescribeInstanceAutoRenewAttributesCommonRequest method.
//    req, resp := client.DescribeInstanceAutoRenewAttributesCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ECS) DescribeInstanceAutoRenewAttributesCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeInstanceAutoRenewAttributesCommon,
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

// DescribeInstanceAutoRenewAttributesCommon API operation for ECS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ECS's
// API operation DescribeInstanceAutoRenewAttributesCommon for usage and error information.
func (c *ECS) DescribeInstanceAutoRenewAttributesCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeInstanceAutoRenewAttributesCommonRequest(input)
	return out, req.Send()
}

// DescribeInstanceAutoRenewAttributesCommonWithContext is the same as DescribeInstanceAutoRenewAttributesCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeInstanceAutoRenewAttributesCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ECS) DescribeInstanceAutoRenewAttributesCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeInstanceAutoRenewAttributesCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeInstanceAutoRenewAttributes = "DescribeInstanceAutoRenewAttributes"

// DescribeInstanceAutoRenewAttributesRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeInstanceAutoRenewAttributes operation. The "output" return
// value will be populated with the DescribeInstanceAutoRenewAttributesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeInstanceAutoRenewAttributesCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeInstanceAutoRenewAttributesCommon Send returns without error.
//
// See DescribeInstanceAutoRenewAttributes for more information on using the DescribeInstanceAutoRenewAttributes
// API call, and error handling.
//
//    // Example sending a request using the DescribeInstanceAutoRenewAttributesRequest method.
//    req, resp := client.DescribeInstanceAutoRenewAttributesRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ECS) DescribeInstanceAutoRenewAttributesRequest(input *DescribeInstanceAutoRenewAttributesInput) (req *request.Request, output *DescribeInstanceAutoRenewAttributesOutput) {
	op := &request.Operation{
		Name:       opDescribeInstanceAutoRenewAttributes,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeInstanceAutoRenewAttributesInput{}
	}

	output = &DescribeInstanceAutoRenewAttributesOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DescribeInstanceAutoRenewAttributes API operation for ECS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ECS's
// API operation DescribeInstanceAutoRenewAttributes for usage and error information.
func (c *ECS) DescribeInstanceAutoRenewAttributes(input *DescribeInstanceAutoRenewAttributesInput) (*DescribeInstanceAutoRenewAttributesOutput, error) {
	req, out := c.DescribeInstanceAutoRenewAttributesRequest(input)
	return out, req.Send()
}

// DescribeInstanceAutoRenewAttributesWithContext is the same as DescribeInstanceAutoRenewAttributes with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeInstanceAutoRenewAttributes for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ECS) DescribeInstanceAutoRenewAttributesWithContext(ctx volcengine.Context, input *DescribeInstanceAutoRenewAttributesInput, opts ...request.Option) (*DescribeInstanceAutoRenewAttributesOutput, error) {
	req, out := c.DescribeInstanceAutoRenewAttributesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DescribeInstanceAutoRenewAttributesInput struct {
	_ struct{} `type:"structure"`

	InstanceIds []*string `type:"list"`

	MaxResults *int32 `type:"int32"`

	NextToken *string `type:"string"`
}

// String returns the string representation
func (s DescribeInstanceAutoRenewAttributesInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeInstanceAutoRenewAttributesInput) GoString() string {
	return s.String()
}

// SetInstanceIds sets the InstanceIds field's value.
func (s *DescribeInstanceAutoRenewAttributesInput) SetInstanceIds(v []*string) *DescribeInstanceAutoRenewAttributesInput {
	s.InstanceIds = v
	return s
}

// SetMaxResults sets the MaxResults field's value.
func (s *DescribeInstanceAutoRenewAttributesInput) SetMaxResults(v int32) *DescribeInstanceAutoRenewAttributesInput {
	s.MaxResults = &v
	return s
}

// SetNextToken sets the NextToken field's value.
func (s *DescribeInstanceAutoRenewAttributesInput) SetNextToken(v string) *DescribeInstanceAutoRenewAttributesInput {
	s.NextToken = &v
	return s
}

type DescribeInstanceAutoRenewAttributesOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	InstanceRenewAttributes []*InstanceRenewAttributeForDescribeInstanceAutoRenewAttributesOutput `type:"list"`

	NextToken *string `type:"string"`
}

// String returns the string representation
func (s DescribeInstanceAutoRenewAttributesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeInstanceAutoRenewAttributesOutput) GoString() string {
	return s.String()
}

// SetInstanceRenewAttributes sets the InstanceRenewAttributes field's value.
func (s *DescribeInstanceAutoRenewAttributesOutput) SetInstanceRenewAttributes(v []*InstanceRenewAttributeForDescribeInstanceAutoRenewAttributesOutput) *DescribeInstanceAutoRenewAttributesOutput {
	s.InstanceRenewAttributes = v
	return s
}

// SetNextToken sets the NextToken field's value.
func (s *DescribeInstanceAutoRenewAttributesOutput) SetNextToken(v string) *DescribeInstanceAutoRenewAttributesOutput {
	s.NextToken = &v
	return s
}

type InstanceRenewAttributeForDescribeInstanceAutoRenewAttributesOutput struct {
	_ struct{} `type:"structure"`

	InstanceId *string `type:"string"`

	RenewType *string `type:"string" enum:"RenewTypeForDescribeInstanceAutoRenewAttributesOutput"`
}

// String returns the string representation
func (s InstanceRenewAttributeForDescribeInstanceAutoRenewAttributesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s InstanceRenewAttributeForDescribeInstanceAutoRenewAttributesOutput) GoString() string {
	return s.String()
}

// SetInstanceId sets the InstanceId field's value.
func (s *InstanceRenewAttributeForDescribeInstanceAutoRenewAttributesOutput) SetInstanceId(v string) *InstanceRenewAttributeForDescribeInstanceAutoRenewAttributesOutput {
	s.InstanceId = &v
	return s
}

// SetRenewType sets the RenewType field's value.
func (s *InstanceRenewAttributeForDescribeInstanceAutoRenewAttributesOutput) SetRenewType(v string) *InstanceRenewAttributeForDescribeInstanceAutoRenewAttributesOutput {
	s.RenewType = &v
	return s
}

const (
	// RenewTypeForDescribeInstanceAutoRenewAttributesOutputAutoRenewDisabled is a RenewTypeForDescribeInstanceAutoRenewAttributesOutput enum value
	RenewTypeForDescribeInstanceAutoRenewAttributesOutputAutoRenewDisabled = "AutoRenewDisabled"

	// RenewTypeForDescribeInstanceAutoRenewAttributesOutputAutoRenewEnabled is a RenewTypeForDescribeInstanceAutoRenewAttributesOutput enum value
	RenewTypeForDescribeInstanceAutoRenewAttributesOutputAutoRenewEnabled = "AutoRenewEnabled"

	// RenewTypeForDescribeInstanceAutoRenewAttributesOutputNonRenewable is a RenewTypeForDescribeInstanceAutoRenewAttributesOutput enum value
	RenewTypeForDescribeInstanceAutoRenewAttributesOutputNonRenewable = "NonRenewable"
)
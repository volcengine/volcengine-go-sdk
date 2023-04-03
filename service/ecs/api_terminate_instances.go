// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package ecs

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opTerminateInstancesCommon = "TerminateInstances"

// TerminateInstancesCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the TerminateInstancesCommon operation. The "output" return
// value will be populated with the TerminateInstancesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned TerminateInstancesCommon Request to send the API call to the service.
// the "output" return value is not valid until after TerminateInstancesCommon Send returns without error.
//
// See TerminateInstancesCommon for more information on using the TerminateInstancesCommon
// API call, and error handling.
//
//    // Example sending a request using the TerminateInstancesCommonRequest method.
//    req, resp := client.TerminateInstancesCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ECS) TerminateInstancesCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opTerminateInstancesCommon,
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

// TerminateInstancesCommon API operation for ECS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ECS's
// API operation TerminateInstancesCommon for usage and error information.
func (c *ECS) TerminateInstancesCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.TerminateInstancesCommonRequest(input)
	return out, req.Send()
}

// TerminateInstancesCommonWithContext is the same as TerminateInstancesCommon with the addition of
// the ability to pass a context and additional request options.
//
// See TerminateInstancesCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ECS) TerminateInstancesCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.TerminateInstancesCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opTerminateInstances = "TerminateInstances"

// TerminateInstancesRequest generates a "volcengine/request.Request" representing the
// client's request for the TerminateInstances operation. The "output" return
// value will be populated with the TerminateInstancesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned TerminateInstancesCommon Request to send the API call to the service.
// the "output" return value is not valid until after TerminateInstancesCommon Send returns without error.
//
// See TerminateInstances for more information on using the TerminateInstances
// API call, and error handling.
//
//    // Example sending a request using the TerminateInstancesRequest method.
//    req, resp := client.TerminateInstancesRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ECS) TerminateInstancesRequest(input *TerminateInstancesInput) (req *request.Request, output *TerminateInstancesOutput) {
	op := &request.Operation{
		Name:       opTerminateInstances,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &TerminateInstancesInput{}
	}

	output = &TerminateInstancesOutput{}
	req = c.newRequest(op, input, output)

	return
}

// TerminateInstances API operation for ECS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ECS's
// API operation TerminateInstances for usage and error information.
func (c *ECS) TerminateInstances(input *TerminateInstancesInput) (*TerminateInstancesOutput, error) {
	req, out := c.TerminateInstancesRequest(input)
	return out, req.Send()
}

// TerminateInstancesWithContext is the same as TerminateInstances with the addition of
// the ability to pass a context and additional request options.
//
// See TerminateInstances for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ECS) TerminateInstancesWithContext(ctx volcengine.Context, input *TerminateInstancesInput, opts ...request.Option) (*TerminateInstancesOutput, error) {
	req, out := c.TerminateInstancesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ErrorForTerminateInstancesOutput struct {
	_ struct{} `type:"structure"`

	Code *string `type:"string"`

	Message *string `type:"string"`
}

// String returns the string representation
func (s ErrorForTerminateInstancesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ErrorForTerminateInstancesOutput) GoString() string {
	return s.String()
}

// SetCode sets the Code field's value.
func (s *ErrorForTerminateInstancesOutput) SetCode(v string) *ErrorForTerminateInstancesOutput {
	s.Code = &v
	return s
}

// SetMessage sets the Message field's value.
func (s *ErrorForTerminateInstancesOutput) SetMessage(v string) *ErrorForTerminateInstancesOutput {
	s.Message = &v
	return s
}

type OperationDetailForTerminateInstancesOutput struct {
	_ struct{} `type:"structure"`

	Error *ErrorForTerminateInstancesOutput `type:"structure"`

	InstanceId *string `type:"string"`
}

// String returns the string representation
func (s OperationDetailForTerminateInstancesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s OperationDetailForTerminateInstancesOutput) GoString() string {
	return s.String()
}

// SetError sets the Error field's value.
func (s *OperationDetailForTerminateInstancesOutput) SetError(v *ErrorForTerminateInstancesOutput) *OperationDetailForTerminateInstancesOutput {
	s.Error = v
	return s
}

// SetInstanceId sets the InstanceId field's value.
func (s *OperationDetailForTerminateInstancesOutput) SetInstanceId(v string) *OperationDetailForTerminateInstancesOutput {
	s.InstanceId = &v
	return s
}

type TerminateInstancesInput struct {
	_ struct{} `type:"structure"`

	ClientToken *string `type:"string"`

	InstanceIds []*string `type:"list"`
}

// String returns the string representation
func (s TerminateInstancesInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s TerminateInstancesInput) GoString() string {
	return s.String()
}

// SetClientToken sets the ClientToken field's value.
func (s *TerminateInstancesInput) SetClientToken(v string) *TerminateInstancesInput {
	s.ClientToken = &v
	return s
}

// SetInstanceIds sets the InstanceIds field's value.
func (s *TerminateInstancesInput) SetInstanceIds(v []*string) *TerminateInstancesInput {
	s.InstanceIds = v
	return s
}

type TerminateInstancesOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	OperationDetails []*OperationDetailForTerminateInstancesOutput `type:"list"`
}

// String returns the string representation
func (s TerminateInstancesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s TerminateInstancesOutput) GoString() string {
	return s.String()
}

// SetOperationDetails sets the OperationDetails field's value.
func (s *TerminateInstancesOutput) SetOperationDetails(v []*OperationDetailForTerminateInstancesOutput) *TerminateInstancesOutput {
	s.OperationDetails = v
	return s
}

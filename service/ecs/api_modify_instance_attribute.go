// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package ecs

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opModifyInstanceAttributeCommon = "ModifyInstanceAttribute"

// ModifyInstanceAttributeCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ModifyInstanceAttributeCommon operation. The "output" return
// value will be populated with the ModifyInstanceAttributeCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyInstanceAttributeCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyInstanceAttributeCommon Send returns without error.
//
// See ModifyInstanceAttributeCommon for more information on using the ModifyInstanceAttributeCommon
// API call, and error handling.
//
//    // Example sending a request using the ModifyInstanceAttributeCommonRequest method.
//    req, resp := client.ModifyInstanceAttributeCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ECS) ModifyInstanceAttributeCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opModifyInstanceAttributeCommon,
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

// ModifyInstanceAttributeCommon API operation for ECS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ECS's
// API operation ModifyInstanceAttributeCommon for usage and error information.
func (c *ECS) ModifyInstanceAttributeCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ModifyInstanceAttributeCommonRequest(input)
	return out, req.Send()
}

// ModifyInstanceAttributeCommonWithContext is the same as ModifyInstanceAttributeCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyInstanceAttributeCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ECS) ModifyInstanceAttributeCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ModifyInstanceAttributeCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opModifyInstanceAttribute = "ModifyInstanceAttribute"

// ModifyInstanceAttributeRequest generates a "volcengine/request.Request" representing the
// client's request for the ModifyInstanceAttribute operation. The "output" return
// value will be populated with the ModifyInstanceAttributeCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyInstanceAttributeCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyInstanceAttributeCommon Send returns without error.
//
// See ModifyInstanceAttribute for more information on using the ModifyInstanceAttribute
// API call, and error handling.
//
//    // Example sending a request using the ModifyInstanceAttributeRequest method.
//    req, resp := client.ModifyInstanceAttributeRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ECS) ModifyInstanceAttributeRequest(input *ModifyInstanceAttributeInput) (req *request.Request, output *ModifyInstanceAttributeOutput) {
	op := &request.Operation{
		Name:       opModifyInstanceAttribute,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ModifyInstanceAttributeInput{}
	}

	output = &ModifyInstanceAttributeOutput{}
	req = c.newRequest(op, input, output)

	return
}

// ModifyInstanceAttribute API operation for ECS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ECS's
// API operation ModifyInstanceAttribute for usage and error information.
func (c *ECS) ModifyInstanceAttribute(input *ModifyInstanceAttributeInput) (*ModifyInstanceAttributeOutput, error) {
	req, out := c.ModifyInstanceAttributeRequest(input)
	return out, req.Send()
}

// ModifyInstanceAttributeWithContext is the same as ModifyInstanceAttribute with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyInstanceAttribute for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ECS) ModifyInstanceAttributeWithContext(ctx volcengine.Context, input *ModifyInstanceAttributeInput, opts ...request.Option) (*ModifyInstanceAttributeOutput, error) {
	req, out := c.ModifyInstanceAttributeRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ModifyInstanceAttributeInput struct {
	_ struct{} `type:"structure"`

	ClientToken *string `type:"string"`

	Description *string `type:"string"`

	Hostname *string `type:"string"`

	// InstanceId is a required field
	InstanceId *string `type:"string" required:"true"`

	InstanceName *string `type:"string"`

	Password *string `type:"string"`

	UserData *string `type:"string"`
}

// String returns the string representation
func (s ModifyInstanceAttributeInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyInstanceAttributeInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ModifyInstanceAttributeInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ModifyInstanceAttributeInput"}
	if s.InstanceId == nil {
		invalidParams.Add(request.NewErrParamRequired("InstanceId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetClientToken sets the ClientToken field's value.
func (s *ModifyInstanceAttributeInput) SetClientToken(v string) *ModifyInstanceAttributeInput {
	s.ClientToken = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *ModifyInstanceAttributeInput) SetDescription(v string) *ModifyInstanceAttributeInput {
	s.Description = &v
	return s
}

// SetHostname sets the Hostname field's value.
func (s *ModifyInstanceAttributeInput) SetHostname(v string) *ModifyInstanceAttributeInput {
	s.Hostname = &v
	return s
}

// SetInstanceId sets the InstanceId field's value.
func (s *ModifyInstanceAttributeInput) SetInstanceId(v string) *ModifyInstanceAttributeInput {
	s.InstanceId = &v
	return s
}

// SetInstanceName sets the InstanceName field's value.
func (s *ModifyInstanceAttributeInput) SetInstanceName(v string) *ModifyInstanceAttributeInput {
	s.InstanceName = &v
	return s
}

// SetPassword sets the Password field's value.
func (s *ModifyInstanceAttributeInput) SetPassword(v string) *ModifyInstanceAttributeInput {
	s.Password = &v
	return s
}

// SetUserData sets the UserData field's value.
func (s *ModifyInstanceAttributeInput) SetUserData(v string) *ModifyInstanceAttributeInput {
	s.UserData = &v
	return s
}

type ModifyInstanceAttributeOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s ModifyInstanceAttributeOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyInstanceAttributeOutput) GoString() string {
	return s.String()
}

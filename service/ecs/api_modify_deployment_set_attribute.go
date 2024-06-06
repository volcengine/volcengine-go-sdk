// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package ecs

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opModifyDeploymentSetAttributeCommon = "ModifyDeploymentSetAttribute"

// ModifyDeploymentSetAttributeCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ModifyDeploymentSetAttributeCommon operation. The "output" return
// value will be populated with the ModifyDeploymentSetAttributeCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyDeploymentSetAttributeCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyDeploymentSetAttributeCommon Send returns without error.
//
// See ModifyDeploymentSetAttributeCommon for more information on using the ModifyDeploymentSetAttributeCommon
// API call, and error handling.
//
//    // Example sending a request using the ModifyDeploymentSetAttributeCommonRequest method.
//    req, resp := client.ModifyDeploymentSetAttributeCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ECS) ModifyDeploymentSetAttributeCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opModifyDeploymentSetAttributeCommon,
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

// ModifyDeploymentSetAttributeCommon API operation for ECS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ECS's
// API operation ModifyDeploymentSetAttributeCommon for usage and error information.
func (c *ECS) ModifyDeploymentSetAttributeCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ModifyDeploymentSetAttributeCommonRequest(input)
	return out, req.Send()
}

// ModifyDeploymentSetAttributeCommonWithContext is the same as ModifyDeploymentSetAttributeCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyDeploymentSetAttributeCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ECS) ModifyDeploymentSetAttributeCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ModifyDeploymentSetAttributeCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opModifyDeploymentSetAttribute = "ModifyDeploymentSetAttribute"

// ModifyDeploymentSetAttributeRequest generates a "volcengine/request.Request" representing the
// client's request for the ModifyDeploymentSetAttribute operation. The "output" return
// value will be populated with the ModifyDeploymentSetAttributeCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyDeploymentSetAttributeCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyDeploymentSetAttributeCommon Send returns without error.
//
// See ModifyDeploymentSetAttribute for more information on using the ModifyDeploymentSetAttribute
// API call, and error handling.
//
//    // Example sending a request using the ModifyDeploymentSetAttributeRequest method.
//    req, resp := client.ModifyDeploymentSetAttributeRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ECS) ModifyDeploymentSetAttributeRequest(input *ModifyDeploymentSetAttributeInput) (req *request.Request, output *ModifyDeploymentSetAttributeOutput) {
	op := &request.Operation{
		Name:       opModifyDeploymentSetAttribute,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ModifyDeploymentSetAttributeInput{}
	}

	output = &ModifyDeploymentSetAttributeOutput{}
	req = c.newRequest(op, input, output)

	return
}

// ModifyDeploymentSetAttribute API operation for ECS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ECS's
// API operation ModifyDeploymentSetAttribute for usage and error information.
func (c *ECS) ModifyDeploymentSetAttribute(input *ModifyDeploymentSetAttributeInput) (*ModifyDeploymentSetAttributeOutput, error) {
	req, out := c.ModifyDeploymentSetAttributeRequest(input)
	return out, req.Send()
}

// ModifyDeploymentSetAttributeWithContext is the same as ModifyDeploymentSetAttribute with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyDeploymentSetAttribute for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ECS) ModifyDeploymentSetAttributeWithContext(ctx volcengine.Context, input *ModifyDeploymentSetAttributeInput, opts ...request.Option) (*ModifyDeploymentSetAttributeOutput, error) {
	req, out := c.ModifyDeploymentSetAttributeRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ModifyDeploymentSetAttributeInput struct {
	_ struct{} `type:"structure"`

	ClientToken *string `type:"string"`

	// DeploymentSetId is a required field
	DeploymentSetId *string `type:"string" required:"true"`

	DeploymentSetName *string `type:"string"`

	Description *string `type:"string"`
}

// String returns the string representation
func (s ModifyDeploymentSetAttributeInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyDeploymentSetAttributeInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ModifyDeploymentSetAttributeInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ModifyDeploymentSetAttributeInput"}
	if s.DeploymentSetId == nil {
		invalidParams.Add(request.NewErrParamRequired("DeploymentSetId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetClientToken sets the ClientToken field's value.
func (s *ModifyDeploymentSetAttributeInput) SetClientToken(v string) *ModifyDeploymentSetAttributeInput {
	s.ClientToken = &v
	return s
}

// SetDeploymentSetId sets the DeploymentSetId field's value.
func (s *ModifyDeploymentSetAttributeInput) SetDeploymentSetId(v string) *ModifyDeploymentSetAttributeInput {
	s.DeploymentSetId = &v
	return s
}

// SetDeploymentSetName sets the DeploymentSetName field's value.
func (s *ModifyDeploymentSetAttributeInput) SetDeploymentSetName(v string) *ModifyDeploymentSetAttributeInput {
	s.DeploymentSetName = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *ModifyDeploymentSetAttributeInput) SetDescription(v string) *ModifyDeploymentSetAttributeInput {
	s.Description = &v
	return s
}

type ModifyDeploymentSetAttributeOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s ModifyDeploymentSetAttributeOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyDeploymentSetAttributeOutput) GoString() string {
	return s.String()
}

// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package iam

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDeleteServiceLinkedRoleCommon = "DeleteServiceLinkedRole"

// DeleteServiceLinkedRoleCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DeleteServiceLinkedRoleCommon operation. The "output" return
// value will be populated with the DeleteServiceLinkedRoleCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeleteServiceLinkedRoleCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeleteServiceLinkedRoleCommon Send returns without error.
//
// See DeleteServiceLinkedRoleCommon for more information on using the DeleteServiceLinkedRoleCommon
// API call, and error handling.
//
//    // Example sending a request using the DeleteServiceLinkedRoleCommonRequest method.
//    req, resp := client.DeleteServiceLinkedRoleCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *IAM) DeleteServiceLinkedRoleCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDeleteServiceLinkedRoleCommon,
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

// DeleteServiceLinkedRoleCommon API operation for IAM.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for IAM's
// API operation DeleteServiceLinkedRoleCommon for usage and error information.
func (c *IAM) DeleteServiceLinkedRoleCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DeleteServiceLinkedRoleCommonRequest(input)
	return out, req.Send()
}

// DeleteServiceLinkedRoleCommonWithContext is the same as DeleteServiceLinkedRoleCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteServiceLinkedRoleCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *IAM) DeleteServiceLinkedRoleCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DeleteServiceLinkedRoleCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDeleteServiceLinkedRole = "DeleteServiceLinkedRole"

// DeleteServiceLinkedRoleRequest generates a "volcengine/request.Request" representing the
// client's request for the DeleteServiceLinkedRole operation. The "output" return
// value will be populated with the DeleteServiceLinkedRoleCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeleteServiceLinkedRoleCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeleteServiceLinkedRoleCommon Send returns without error.
//
// See DeleteServiceLinkedRole for more information on using the DeleteServiceLinkedRole
// API call, and error handling.
//
//    // Example sending a request using the DeleteServiceLinkedRoleRequest method.
//    req, resp := client.DeleteServiceLinkedRoleRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *IAM) DeleteServiceLinkedRoleRequest(input *DeleteServiceLinkedRoleInput) (req *request.Request, output *DeleteServiceLinkedRoleOutput) {
	op := &request.Operation{
		Name:       opDeleteServiceLinkedRole,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteServiceLinkedRoleInput{}
	}

	output = &DeleteServiceLinkedRoleOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DeleteServiceLinkedRole API operation for IAM.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for IAM's
// API operation DeleteServiceLinkedRole for usage and error information.
func (c *IAM) DeleteServiceLinkedRole(input *DeleteServiceLinkedRoleInput) (*DeleteServiceLinkedRoleOutput, error) {
	req, out := c.DeleteServiceLinkedRoleRequest(input)
	return out, req.Send()
}

// DeleteServiceLinkedRoleWithContext is the same as DeleteServiceLinkedRole with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteServiceLinkedRole for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *IAM) DeleteServiceLinkedRoleWithContext(ctx volcengine.Context, input *DeleteServiceLinkedRoleInput, opts ...request.Option) (*DeleteServiceLinkedRoleOutput, error) {
	req, out := c.DeleteServiceLinkedRoleRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DeleteServiceLinkedRoleInput struct {
	_ struct{} `type:"structure"`

	// RoleName is a required field
	RoleName *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteServiceLinkedRoleInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteServiceLinkedRoleInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteServiceLinkedRoleInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DeleteServiceLinkedRoleInput"}
	if s.RoleName == nil {
		invalidParams.Add(request.NewErrParamRequired("RoleName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetRoleName sets the RoleName field's value.
func (s *DeleteServiceLinkedRoleInput) SetRoleName(v string) *DeleteServiceLinkedRoleInput {
	s.RoleName = &v
	return s
}

type DeleteServiceLinkedRoleOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s DeleteServiceLinkedRoleOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteServiceLinkedRoleOutput) GoString() string {
	return s.String()
}

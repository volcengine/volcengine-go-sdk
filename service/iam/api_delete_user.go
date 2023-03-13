// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package iam

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDeleteUserCommon = "DeleteUser"

// DeleteUserCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DeleteUserCommon operation. The "output" return
// value will be populated with the DeleteUserCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeleteUserCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeleteUserCommon Send returns without error.
//
// See DeleteUserCommon for more information on using the DeleteUserCommon
// API call, and error handling.
//
//    // Example sending a request using the DeleteUserCommonRequest method.
//    req, resp := client.DeleteUserCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *IAM) DeleteUserCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDeleteUserCommon,
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

// DeleteUserCommon API operation for IAM.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for IAM's
// API operation DeleteUserCommon for usage and error information.
func (c *IAM) DeleteUserCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DeleteUserCommonRequest(input)
	return out, req.Send()
}

// DeleteUserCommonWithContext is the same as DeleteUserCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteUserCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *IAM) DeleteUserCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DeleteUserCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDeleteUser = "DeleteUser"

// DeleteUserRequest generates a "volcengine/request.Request" representing the
// client's request for the DeleteUser operation. The "output" return
// value will be populated with the DeleteUserCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeleteUserCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeleteUserCommon Send returns without error.
//
// See DeleteUser for more information on using the DeleteUser
// API call, and error handling.
//
//    // Example sending a request using the DeleteUserRequest method.
//    req, resp := client.DeleteUserRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *IAM) DeleteUserRequest(input *DeleteUserInput) (req *request.Request, output *DeleteUserOutput) {
	op := &request.Operation{
		Name:       opDeleteUser,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteUserInput{}
	}

	output = &DeleteUserOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DeleteUser API operation for IAM.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for IAM's
// API operation DeleteUser for usage and error information.
func (c *IAM) DeleteUser(input *DeleteUserInput) (*DeleteUserOutput, error) {
	req, out := c.DeleteUserRequest(input)
	return out, req.Send()
}

// DeleteUserWithContext is the same as DeleteUser with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteUser for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *IAM) DeleteUserWithContext(ctx volcengine.Context, input *DeleteUserInput, opts ...request.Option) (*DeleteUserOutput, error) {
	req, out := c.DeleteUserRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DeleteUserInput struct {
	_ struct{} `type:"structure"`

	// UserName is a required field
	UserName *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteUserInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteUserInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteUserInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DeleteUserInput"}
	if s.UserName == nil {
		invalidParams.Add(request.NewErrParamRequired("UserName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetUserName sets the UserName field's value.
func (s *DeleteUserInput) SetUserName(v string) *DeleteUserInput {
	s.UserName = &v
	return s
}

type DeleteUserOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s DeleteUserOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteUserOutput) GoString() string {
	return s.String()
}
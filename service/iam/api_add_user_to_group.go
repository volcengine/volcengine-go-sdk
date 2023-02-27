// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package iam

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opAddUserToGroupCommon = "AddUserToGroup"

// AddUserToGroupCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the AddUserToGroupCommon operation. The "output" return
// value will be populated with the AddUserToGroupCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned AddUserToGroupCommon Request to send the API call to the service.
// the "output" return value is not valid until after AddUserToGroupCommon Send returns without error.
//
// See AddUserToGroupCommon for more information on using the AddUserToGroupCommon
// API call, and error handling.
//
//    // Example sending a request using the AddUserToGroupCommonRequest method.
//    req, resp := client.AddUserToGroupCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *IAM) AddUserToGroupCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opAddUserToGroupCommon,
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

// AddUserToGroupCommon API operation for IAM.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for IAM's
// API operation AddUserToGroupCommon for usage and error information.
func (c *IAM) AddUserToGroupCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.AddUserToGroupCommonRequest(input)
	return out, req.Send()
}

// AddUserToGroupCommonWithContext is the same as AddUserToGroupCommon with the addition of
// the ability to pass a context and additional request options.
//
// See AddUserToGroupCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *IAM) AddUserToGroupCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.AddUserToGroupCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opAddUserToGroup = "AddUserToGroup"

// AddUserToGroupRequest generates a "volcengine/request.Request" representing the
// client's request for the AddUserToGroup operation. The "output" return
// value will be populated with the AddUserToGroupCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned AddUserToGroupCommon Request to send the API call to the service.
// the "output" return value is not valid until after AddUserToGroupCommon Send returns without error.
//
// See AddUserToGroup for more information on using the AddUserToGroup
// API call, and error handling.
//
//    // Example sending a request using the AddUserToGroupRequest method.
//    req, resp := client.AddUserToGroupRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *IAM) AddUserToGroupRequest(input *AddUserToGroupInput) (req *request.Request, output *AddUserToGroupOutput) {
	op := &request.Operation{
		Name:       opAddUserToGroup,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &AddUserToGroupInput{}
	}

	output = &AddUserToGroupOutput{}
	req = c.newRequest(op, input, output)

	return
}

// AddUserToGroup API operation for IAM.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for IAM's
// API operation AddUserToGroup for usage and error information.
func (c *IAM) AddUserToGroup(input *AddUserToGroupInput) (*AddUserToGroupOutput, error) {
	req, out := c.AddUserToGroupRequest(input)
	return out, req.Send()
}

// AddUserToGroupWithContext is the same as AddUserToGroup with the addition of
// the ability to pass a context and additional request options.
//
// See AddUserToGroup for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *IAM) AddUserToGroupWithContext(ctx volcengine.Context, input *AddUserToGroupInput, opts ...request.Option) (*AddUserToGroupOutput, error) {
	req, out := c.AddUserToGroupRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type AddUserToGroupInput struct {
	_ struct{} `type:"structure"`

	// UserGroupName is a required field
	UserGroupName *string `type:"string" required:"true"`

	// UserName is a required field
	UserName *string `type:"string" required:"true"`
}

// String returns the string representation
func (s AddUserToGroupInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s AddUserToGroupInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AddUserToGroupInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "AddUserToGroupInput"}
	if s.UserGroupName == nil {
		invalidParams.Add(request.NewErrParamRequired("UserGroupName"))
	}
	if s.UserName == nil {
		invalidParams.Add(request.NewErrParamRequired("UserName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetUserGroupName sets the UserGroupName field's value.
func (s *AddUserToGroupInput) SetUserGroupName(v string) *AddUserToGroupInput {
	s.UserGroupName = &v
	return s
}

// SetUserName sets the UserName field's value.
func (s *AddUserToGroupInput) SetUserName(v string) *AddUserToGroupInput {
	s.UserName = &v
	return s
}

type AddUserToGroupOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s AddUserToGroupOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s AddUserToGroupOutput) GoString() string {
	return s.String()
}

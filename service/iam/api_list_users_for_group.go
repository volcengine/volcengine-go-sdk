// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package iam

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opListUsersForGroupCommon = "ListUsersForGroup"

// ListUsersForGroupCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ListUsersForGroupCommon operation. The "output" return
// value will be populated with the ListUsersForGroupCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListUsersForGroupCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListUsersForGroupCommon Send returns without error.
//
// See ListUsersForGroupCommon for more information on using the ListUsersForGroupCommon
// API call, and error handling.
//
//    // Example sending a request using the ListUsersForGroupCommonRequest method.
//    req, resp := client.ListUsersForGroupCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *IAM) ListUsersForGroupCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opListUsersForGroupCommon,
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

// ListUsersForGroupCommon API operation for IAM.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for IAM's
// API operation ListUsersForGroupCommon for usage and error information.
func (c *IAM) ListUsersForGroupCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ListUsersForGroupCommonRequest(input)
	return out, req.Send()
}

// ListUsersForGroupCommonWithContext is the same as ListUsersForGroupCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ListUsersForGroupCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *IAM) ListUsersForGroupCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ListUsersForGroupCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opListUsersForGroup = "ListUsersForGroup"

// ListUsersForGroupRequest generates a "volcengine/request.Request" representing the
// client's request for the ListUsersForGroup operation. The "output" return
// value will be populated with the ListUsersForGroupCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListUsersForGroupCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListUsersForGroupCommon Send returns without error.
//
// See ListUsersForGroup for more information on using the ListUsersForGroup
// API call, and error handling.
//
//    // Example sending a request using the ListUsersForGroupRequest method.
//    req, resp := client.ListUsersForGroupRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *IAM) ListUsersForGroupRequest(input *ListUsersForGroupInput) (req *request.Request, output *ListUsersForGroupOutput) {
	op := &request.Operation{
		Name:       opListUsersForGroup,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListUsersForGroupInput{}
	}

	output = &ListUsersForGroupOutput{}
	req = c.newRequest(op, input, output)

	return
}

// ListUsersForGroup API operation for IAM.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for IAM's
// API operation ListUsersForGroup for usage and error information.
func (c *IAM) ListUsersForGroup(input *ListUsersForGroupInput) (*ListUsersForGroupOutput, error) {
	req, out := c.ListUsersForGroupRequest(input)
	return out, req.Send()
}

// ListUsersForGroupWithContext is the same as ListUsersForGroup with the addition of
// the ability to pass a context and additional request options.
//
// See ListUsersForGroup for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *IAM) ListUsersForGroupWithContext(ctx volcengine.Context, input *ListUsersForGroupInput, opts ...request.Option) (*ListUsersForGroupOutput, error) {
	req, out := c.ListUsersForGroupRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ListUsersForGroupInput struct {
	_ struct{} `type:"structure"`

	// UserGroupName is a required field
	UserGroupName *string `type:"string" required:"true"`
}

// String returns the string representation
func (s ListUsersForGroupInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListUsersForGroupInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListUsersForGroupInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ListUsersForGroupInput"}
	if s.UserGroupName == nil {
		invalidParams.Add(request.NewErrParamRequired("UserGroupName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetUserGroupName sets the UserGroupName field's value.
func (s *ListUsersForGroupInput) SetUserGroupName(v string) *ListUsersForGroupInput {
	s.UserGroupName = &v
	return s
}

type ListUsersForGroupOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	ResponseMetadata *interface{} `type:"interface"`

	Result *interface{} `type:"interface"`
}

// String returns the string representation
func (s ListUsersForGroupOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListUsersForGroupOutput) GoString() string {
	return s.String()
}

// SetResponseMetadata sets the ResponseMetadata field's value.
func (s *ListUsersForGroupOutput) SetResponseMetadata(v interface{}) *ListUsersForGroupOutput {
	s.ResponseMetadata = &v
	return s
}

// SetResult sets the Result field's value.
func (s *ListUsersForGroupOutput) SetResult(v interface{}) *ListUsersForGroupOutput {
	s.Result = &v
	return s
}

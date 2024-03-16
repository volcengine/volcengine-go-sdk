// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package iam

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opListGroupsCommon = "ListGroups"

// ListGroupsCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ListGroupsCommon operation. The "output" return
// value will be populated with the ListGroupsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListGroupsCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListGroupsCommon Send returns without error.
//
// See ListGroupsCommon for more information on using the ListGroupsCommon
// API call, and error handling.
//
//    // Example sending a request using the ListGroupsCommonRequest method.
//    req, resp := client.ListGroupsCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *IAM) ListGroupsCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opListGroupsCommon,
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

// ListGroupsCommon API operation for IAM.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for IAM's
// API operation ListGroupsCommon for usage and error information.
func (c *IAM) ListGroupsCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ListGroupsCommonRequest(input)
	return out, req.Send()
}

// ListGroupsCommonWithContext is the same as ListGroupsCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ListGroupsCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *IAM) ListGroupsCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ListGroupsCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opListGroups = "ListGroups"

// ListGroupsRequest generates a "volcengine/request.Request" representing the
// client's request for the ListGroups operation. The "output" return
// value will be populated with the ListGroupsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListGroupsCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListGroupsCommon Send returns without error.
//
// See ListGroups for more information on using the ListGroups
// API call, and error handling.
//
//    // Example sending a request using the ListGroupsRequest method.
//    req, resp := client.ListGroupsRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *IAM) ListGroupsRequest(input *ListGroupsInput) (req *request.Request, output *ListGroupsOutput) {
	op := &request.Operation{
		Name:       opListGroups,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListGroupsInput{}
	}

	output = &ListGroupsOutput{}
	req = c.newRequest(op, input, output)

	return
}

// ListGroups API operation for IAM.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for IAM's
// API operation ListGroups for usage and error information.
func (c *IAM) ListGroups(input *ListGroupsInput) (*ListGroupsOutput, error) {
	req, out := c.ListGroupsRequest(input)
	return out, req.Send()
}

// ListGroupsWithContext is the same as ListGroups with the addition of
// the ability to pass a context and additional request options.
//
// See ListGroups for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *IAM) ListGroupsWithContext(ctx volcengine.Context, input *ListGroupsInput, opts ...request.Option) (*ListGroupsOutput, error) {
	req, out := c.ListGroupsRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ListGroupsInput struct {
	_ struct{} `type:"structure"`

	Limit *int32 `type:"int32"`

	Offset *int32 `type:"int32"`

	Query *string `type:"string"`
}

// String returns the string representation
func (s ListGroupsInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListGroupsInput) GoString() string {
	return s.String()
}

// SetLimit sets the Limit field's value.
func (s *ListGroupsInput) SetLimit(v int32) *ListGroupsInput {
	s.Limit = &v
	return s
}

// SetOffset sets the Offset field's value.
func (s *ListGroupsInput) SetOffset(v int32) *ListGroupsInput {
	s.Offset = &v
	return s
}

// SetQuery sets the Query field's value.
func (s *ListGroupsInput) SetQuery(v string) *ListGroupsInput {
	s.Query = &v
	return s
}

type ListGroupsOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	Limit *int32 `type:"int32"`

	Offset *int32 `type:"int32"`

	Total *int32 `type:"int32"`

	UserGroups []*UserGroupForListGroupsOutput `type:"list"`
}

// String returns the string representation
func (s ListGroupsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListGroupsOutput) GoString() string {
	return s.String()
}

// SetLimit sets the Limit field's value.
func (s *ListGroupsOutput) SetLimit(v int32) *ListGroupsOutput {
	s.Limit = &v
	return s
}

// SetOffset sets the Offset field's value.
func (s *ListGroupsOutput) SetOffset(v int32) *ListGroupsOutput {
	s.Offset = &v
	return s
}

// SetTotal sets the Total field's value.
func (s *ListGroupsOutput) SetTotal(v int32) *ListGroupsOutput {
	s.Total = &v
	return s
}

// SetUserGroups sets the UserGroups field's value.
func (s *ListGroupsOutput) SetUserGroups(v []*UserGroupForListGroupsOutput) *ListGroupsOutput {
	s.UserGroups = v
	return s
}

type UserGroupForListGroupsOutput struct {
	_ struct{} `type:"structure"`

	AccountID *int64 `type:"int64"`

	CreateDate *string `type:"string"`

	Description *string `type:"string"`

	DisplayName *string `type:"string"`

	UpdateDate *string `type:"string"`

	UserGroupName *string `type:"string"`
}

// String returns the string representation
func (s UserGroupForListGroupsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s UserGroupForListGroupsOutput) GoString() string {
	return s.String()
}

// SetAccountID sets the AccountID field's value.
func (s *UserGroupForListGroupsOutput) SetAccountID(v int64) *UserGroupForListGroupsOutput {
	s.AccountID = &v
	return s
}

// SetCreateDate sets the CreateDate field's value.
func (s *UserGroupForListGroupsOutput) SetCreateDate(v string) *UserGroupForListGroupsOutput {
	s.CreateDate = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *UserGroupForListGroupsOutput) SetDescription(v string) *UserGroupForListGroupsOutput {
	s.Description = &v
	return s
}

// SetDisplayName sets the DisplayName field's value.
func (s *UserGroupForListGroupsOutput) SetDisplayName(v string) *UserGroupForListGroupsOutput {
	s.DisplayName = &v
	return s
}

// SetUpdateDate sets the UpdateDate field's value.
func (s *UserGroupForListGroupsOutput) SetUpdateDate(v string) *UserGroupForListGroupsOutput {
	s.UpdateDate = &v
	return s
}

// SetUserGroupName sets the UserGroupName field's value.
func (s *UserGroupForListGroupsOutput) SetUserGroupName(v string) *UserGroupForListGroupsOutput {
	s.UserGroupName = &v
	return s
}

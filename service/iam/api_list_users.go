// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package iam

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opListUsersCommon = "ListUsers"

// ListUsersCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ListUsersCommon operation. The "output" return
// value will be populated with the ListUsersCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListUsersCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListUsersCommon Send returns without error.
//
// See ListUsersCommon for more information on using the ListUsersCommon
// API call, and error handling.
//
//    // Example sending a request using the ListUsersCommonRequest method.
//    req, resp := client.ListUsersCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *IAM) ListUsersCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opListUsersCommon,
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

// ListUsersCommon API operation for IAM.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for IAM's
// API operation ListUsersCommon for usage and error information.
func (c *IAM) ListUsersCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ListUsersCommonRequest(input)
	return out, req.Send()
}

// ListUsersCommonWithContext is the same as ListUsersCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ListUsersCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *IAM) ListUsersCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ListUsersCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opListUsers = "ListUsers"

// ListUsersRequest generates a "volcengine/request.Request" representing the
// client's request for the ListUsers operation. The "output" return
// value will be populated with the ListUsersCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListUsersCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListUsersCommon Send returns without error.
//
// See ListUsers for more information on using the ListUsers
// API call, and error handling.
//
//    // Example sending a request using the ListUsersRequest method.
//    req, resp := client.ListUsersRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *IAM) ListUsersRequest(input *ListUsersInput) (req *request.Request, output *ListUsersOutput) {
	op := &request.Operation{
		Name:       opListUsers,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListUsersInput{}
	}

	output = &ListUsersOutput{}
	req = c.newRequest(op, input, output)

	return
}

// ListUsers API operation for IAM.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for IAM's
// API operation ListUsers for usage and error information.
func (c *IAM) ListUsers(input *ListUsersInput) (*ListUsersOutput, error) {
	req, out := c.ListUsersRequest(input)
	return out, req.Send()
}

// ListUsersWithContext is the same as ListUsers with the addition of
// the ability to pass a context and additional request options.
//
// See ListUsers for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *IAM) ListUsersWithContext(ctx volcengine.Context, input *ListUsersInput, opts ...request.Option) (*ListUsersOutput, error) {
	req, out := c.ListUsersRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ListUsersInput struct {
	_ struct{} `type:"structure"`

	Limit *int64 `type:"integer"`

	Offset *string `type:"string"`

	Query *string `type:"string"`
}

// String returns the string representation
func (s ListUsersInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListUsersInput) GoString() string {
	return s.String()
}

// SetLimit sets the Limit field's value.
func (s *ListUsersInput) SetLimit(v int64) *ListUsersInput {
	s.Limit = &v
	return s
}

// SetOffset sets the Offset field's value.
func (s *ListUsersInput) SetOffset(v string) *ListUsersInput {
	s.Offset = &v
	return s
}

// SetQuery sets the Query field's value.
func (s *ListUsersInput) SetQuery(v string) *ListUsersInput {
	s.Query = &v
	return s
}

type ListUsersOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	Limit *int32 `type:"int32"`

	Offset *int32 `type:"int32"`

	Total *int32 `type:"int32"`

	UserMetadata []*UserMetadataForListUsersOutput `type:"list"`
}

// String returns the string representation
func (s ListUsersOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListUsersOutput) GoString() string {
	return s.String()
}

// SetLimit sets the Limit field's value.
func (s *ListUsersOutput) SetLimit(v int32) *ListUsersOutput {
	s.Limit = &v
	return s
}

// SetOffset sets the Offset field's value.
func (s *ListUsersOutput) SetOffset(v int32) *ListUsersOutput {
	s.Offset = &v
	return s
}

// SetTotal sets the Total field's value.
func (s *ListUsersOutput) SetTotal(v int32) *ListUsersOutput {
	s.Total = &v
	return s
}

// SetUserMetadata sets the UserMetadata field's value.
func (s *ListUsersOutput) SetUserMetadata(v []*UserMetadataForListUsersOutput) *ListUsersOutput {
	s.UserMetadata = v
	return s
}

type UserMetadataForListUsersOutput struct {
	_ struct{} `type:"structure"`

	AccountId *int64 `type:"int64"`

	CreateDate *string `type:"string"`

	Description *string `type:"string"`

	DisplayName *string `type:"string"`

	Email *string `type:"string"`

	EmailIsVerify *bool `type:"boolean"`

	Id *int32 `type:"int32"`

	MobilePhone *string `type:"string"`

	MobilePhoneIsVerify *bool `type:"boolean"`

	Trn *string `type:"string"`

	UpdateDate *string `type:"string"`

	UserName *string `type:"string"`
}

// String returns the string representation
func (s UserMetadataForListUsersOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s UserMetadataForListUsersOutput) GoString() string {
	return s.String()
}

// SetAccountId sets the AccountId field's value.
func (s *UserMetadataForListUsersOutput) SetAccountId(v int64) *UserMetadataForListUsersOutput {
	s.AccountId = &v
	return s
}

// SetCreateDate sets the CreateDate field's value.
func (s *UserMetadataForListUsersOutput) SetCreateDate(v string) *UserMetadataForListUsersOutput {
	s.CreateDate = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *UserMetadataForListUsersOutput) SetDescription(v string) *UserMetadataForListUsersOutput {
	s.Description = &v
	return s
}

// SetDisplayName sets the DisplayName field's value.
func (s *UserMetadataForListUsersOutput) SetDisplayName(v string) *UserMetadataForListUsersOutput {
	s.DisplayName = &v
	return s
}

// SetEmail sets the Email field's value.
func (s *UserMetadataForListUsersOutput) SetEmail(v string) *UserMetadataForListUsersOutput {
	s.Email = &v
	return s
}

// SetEmailIsVerify sets the EmailIsVerify field's value.
func (s *UserMetadataForListUsersOutput) SetEmailIsVerify(v bool) *UserMetadataForListUsersOutput {
	s.EmailIsVerify = &v
	return s
}

// SetId sets the Id field's value.
func (s *UserMetadataForListUsersOutput) SetId(v int32) *UserMetadataForListUsersOutput {
	s.Id = &v
	return s
}

// SetMobilePhone sets the MobilePhone field's value.
func (s *UserMetadataForListUsersOutput) SetMobilePhone(v string) *UserMetadataForListUsersOutput {
	s.MobilePhone = &v
	return s
}

// SetMobilePhoneIsVerify sets the MobilePhoneIsVerify field's value.
func (s *UserMetadataForListUsersOutput) SetMobilePhoneIsVerify(v bool) *UserMetadataForListUsersOutput {
	s.MobilePhoneIsVerify = &v
	return s
}

// SetTrn sets the Trn field's value.
func (s *UserMetadataForListUsersOutput) SetTrn(v string) *UserMetadataForListUsersOutput {
	s.Trn = &v
	return s
}

// SetUpdateDate sets the UpdateDate field's value.
func (s *UserMetadataForListUsersOutput) SetUpdateDate(v string) *UserMetadataForListUsersOutput {
	s.UpdateDate = &v
	return s
}

// SetUserName sets the UserName field's value.
func (s *UserMetadataForListUsersOutput) SetUserName(v string) *UserMetadataForListUsersOutput {
	s.UserName = &v
	return s
}

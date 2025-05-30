// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package iam

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opCreateRoleCommon = "CreateRole"

// CreateRoleCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateRoleCommon operation. The "output" return
// value will be populated with the CreateRoleCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateRoleCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateRoleCommon Send returns without error.
//
// See CreateRoleCommon for more information on using the CreateRoleCommon
// API call, and error handling.
//
//    // Example sending a request using the CreateRoleCommonRequest method.
//    req, resp := client.CreateRoleCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *IAM) CreateRoleCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opCreateRoleCommon,
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

// CreateRoleCommon API operation for IAM.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for IAM's
// API operation CreateRoleCommon for usage and error information.
func (c *IAM) CreateRoleCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.CreateRoleCommonRequest(input)
	return out, req.Send()
}

// CreateRoleCommonWithContext is the same as CreateRoleCommon with the addition of
// the ability to pass a context and additional request options.
//
// See CreateRoleCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *IAM) CreateRoleCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.CreateRoleCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opCreateRole = "CreateRole"

// CreateRoleRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateRole operation. The "output" return
// value will be populated with the CreateRoleCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateRoleCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateRoleCommon Send returns without error.
//
// See CreateRole for more information on using the CreateRole
// API call, and error handling.
//
//    // Example sending a request using the CreateRoleRequest method.
//    req, resp := client.CreateRoleRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *IAM) CreateRoleRequest(input *CreateRoleInput) (req *request.Request, output *CreateRoleOutput) {
	op := &request.Operation{
		Name:       opCreateRole,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateRoleInput{}
	}

	output = &CreateRoleOutput{}
	req = c.newRequest(op, input, output)

	return
}

// CreateRole API operation for IAM.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for IAM's
// API operation CreateRole for usage and error information.
func (c *IAM) CreateRole(input *CreateRoleInput) (*CreateRoleOutput, error) {
	req, out := c.CreateRoleRequest(input)
	return out, req.Send()
}

// CreateRoleWithContext is the same as CreateRole with the addition of
// the ability to pass a context and additional request options.
//
// See CreateRole for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *IAM) CreateRoleWithContext(ctx volcengine.Context, input *CreateRoleInput, opts ...request.Option) (*CreateRoleOutput, error) {
	req, out := c.CreateRoleRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type CreateRoleInput struct {
	_ struct{} `type:"structure"`

	Description *string `type:"string"`

	DisplayName *string `type:"string"`

	MaxSessionDuration *int32 `type:"int32"`

	// RoleName is a required field
	RoleName *string `type:"string" required:"true"`

	Tags []*TagForCreateRoleInput `type:"list"`

	TrustPolicyDocument *string `type:"string"`
}

// String returns the string representation
func (s CreateRoleInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateRoleInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateRoleInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "CreateRoleInput"}
	if s.RoleName == nil {
		invalidParams.Add(request.NewErrParamRequired("RoleName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetDescription sets the Description field's value.
func (s *CreateRoleInput) SetDescription(v string) *CreateRoleInput {
	s.Description = &v
	return s
}

// SetDisplayName sets the DisplayName field's value.
func (s *CreateRoleInput) SetDisplayName(v string) *CreateRoleInput {
	s.DisplayName = &v
	return s
}

// SetMaxSessionDuration sets the MaxSessionDuration field's value.
func (s *CreateRoleInput) SetMaxSessionDuration(v int32) *CreateRoleInput {
	s.MaxSessionDuration = &v
	return s
}

// SetRoleName sets the RoleName field's value.
func (s *CreateRoleInput) SetRoleName(v string) *CreateRoleInput {
	s.RoleName = &v
	return s
}

// SetTags sets the Tags field's value.
func (s *CreateRoleInput) SetTags(v []*TagForCreateRoleInput) *CreateRoleInput {
	s.Tags = v
	return s
}

// SetTrustPolicyDocument sets the TrustPolicyDocument field's value.
func (s *CreateRoleInput) SetTrustPolicyDocument(v string) *CreateRoleInput {
	s.TrustPolicyDocument = &v
	return s
}

type CreateRoleOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	Role *RoleForCreateRoleOutput `type:"structure"`
}

// String returns the string representation
func (s CreateRoleOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateRoleOutput) GoString() string {
	return s.String()
}

// SetRole sets the Role field's value.
func (s *CreateRoleOutput) SetRole(v *RoleForCreateRoleOutput) *CreateRoleOutput {
	s.Role = v
	return s
}

type RoleForCreateRoleOutput struct {
	_ struct{} `type:"structure"`

	CreateDate *string `type:"string"`

	Description *string `type:"string"`

	DisplayName *string `type:"string"`

	IsServiceLinkedRole *int32 `type:"int32"`

	MaxSessionDuration *int32 `type:"int32"`

	RoleId *int32 `type:"int32"`

	RoleName *string `type:"string"`

	Tags []*TagForCreateRoleOutput `type:"list"`

	Trn *string `type:"string"`

	TrustPolicyDocument *string `type:"string"`

	UpdateDate *string `type:"string"`
}

// String returns the string representation
func (s RoleForCreateRoleOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s RoleForCreateRoleOutput) GoString() string {
	return s.String()
}

// SetCreateDate sets the CreateDate field's value.
func (s *RoleForCreateRoleOutput) SetCreateDate(v string) *RoleForCreateRoleOutput {
	s.CreateDate = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *RoleForCreateRoleOutput) SetDescription(v string) *RoleForCreateRoleOutput {
	s.Description = &v
	return s
}

// SetDisplayName sets the DisplayName field's value.
func (s *RoleForCreateRoleOutput) SetDisplayName(v string) *RoleForCreateRoleOutput {
	s.DisplayName = &v
	return s
}

// SetIsServiceLinkedRole sets the IsServiceLinkedRole field's value.
func (s *RoleForCreateRoleOutput) SetIsServiceLinkedRole(v int32) *RoleForCreateRoleOutput {
	s.IsServiceLinkedRole = &v
	return s
}

// SetMaxSessionDuration sets the MaxSessionDuration field's value.
func (s *RoleForCreateRoleOutput) SetMaxSessionDuration(v int32) *RoleForCreateRoleOutput {
	s.MaxSessionDuration = &v
	return s
}

// SetRoleId sets the RoleId field's value.
func (s *RoleForCreateRoleOutput) SetRoleId(v int32) *RoleForCreateRoleOutput {
	s.RoleId = &v
	return s
}

// SetRoleName sets the RoleName field's value.
func (s *RoleForCreateRoleOutput) SetRoleName(v string) *RoleForCreateRoleOutput {
	s.RoleName = &v
	return s
}

// SetTags sets the Tags field's value.
func (s *RoleForCreateRoleOutput) SetTags(v []*TagForCreateRoleOutput) *RoleForCreateRoleOutput {
	s.Tags = v
	return s
}

// SetTrn sets the Trn field's value.
func (s *RoleForCreateRoleOutput) SetTrn(v string) *RoleForCreateRoleOutput {
	s.Trn = &v
	return s
}

// SetTrustPolicyDocument sets the TrustPolicyDocument field's value.
func (s *RoleForCreateRoleOutput) SetTrustPolicyDocument(v string) *RoleForCreateRoleOutput {
	s.TrustPolicyDocument = &v
	return s
}

// SetUpdateDate sets the UpdateDate field's value.
func (s *RoleForCreateRoleOutput) SetUpdateDate(v string) *RoleForCreateRoleOutput {
	s.UpdateDate = &v
	return s
}

type TagForCreateRoleInput struct {
	_ struct{} `type:"structure"`

	Key *string `type:"string"`

	Value *string `type:"string"`
}

// String returns the string representation
func (s TagForCreateRoleInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s TagForCreateRoleInput) GoString() string {
	return s.String()
}

// SetKey sets the Key field's value.
func (s *TagForCreateRoleInput) SetKey(v string) *TagForCreateRoleInput {
	s.Key = &v
	return s
}

// SetValue sets the Value field's value.
func (s *TagForCreateRoleInput) SetValue(v string) *TagForCreateRoleInput {
	s.Value = &v
	return s
}

type TagForCreateRoleOutput struct {
	_ struct{} `type:"structure"`

	Key *string `type:"string"`

	Value *string `type:"string"`
}

// String returns the string representation
func (s TagForCreateRoleOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s TagForCreateRoleOutput) GoString() string {
	return s.String()
}

// SetKey sets the Key field's value.
func (s *TagForCreateRoleOutput) SetKey(v string) *TagForCreateRoleOutput {
	s.Key = &v
	return s
}

// SetValue sets the Value field's value.
func (s *TagForCreateRoleOutput) SetValue(v string) *TagForCreateRoleOutput {
	s.Value = &v
	return s
}

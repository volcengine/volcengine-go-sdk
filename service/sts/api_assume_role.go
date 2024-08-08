// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package sts

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opAssumeRoleCommon = "AssumeRole"

// AssumeRoleCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the AssumeRoleCommon operation. The "output" return
// value will be populated with the AssumeRoleCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned AssumeRoleCommon Request to send the API call to the service.
// the "output" return value is not valid until after AssumeRoleCommon Send returns without error.
//
// See AssumeRoleCommon for more information on using the AssumeRoleCommon
// API call, and error handling.
//
//    // Example sending a request using the AssumeRoleCommonRequest method.
//    req, resp := client.AssumeRoleCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *STS) AssumeRoleCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opAssumeRoleCommon,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &map[string]interface{}{}
	}

	output = &map[string]interface{}{}
	req = c.newRequest(op, input, output)

	return
}

// AssumeRoleCommon API operation for STS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for STS's
// API operation AssumeRoleCommon for usage and error information.
func (c *STS) AssumeRoleCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.AssumeRoleCommonRequest(input)
	return out, req.Send()
}

// AssumeRoleCommonWithContext is the same as AssumeRoleCommon with the addition of
// the ability to pass a context and additional request options.
//
// See AssumeRoleCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *STS) AssumeRoleCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.AssumeRoleCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opAssumeRole = "AssumeRole"

// AssumeRoleRequest generates a "volcengine/request.Request" representing the
// client's request for the AssumeRole operation. The "output" return
// value will be populated with the AssumeRoleCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned AssumeRoleCommon Request to send the API call to the service.
// the "output" return value is not valid until after AssumeRoleCommon Send returns without error.
//
// See AssumeRole for more information on using the AssumeRole
// API call, and error handling.
//
//    // Example sending a request using the AssumeRoleRequest method.
//    req, resp := client.AssumeRoleRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *STS) AssumeRoleRequest(input *AssumeRoleInput) (req *request.Request, output *AssumeRoleOutput) {
	op := &request.Operation{
		Name:       opAssumeRole,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &AssumeRoleInput{}
	}

	output = &AssumeRoleOutput{}
	req = c.newRequest(op, input, output)

	return
}

// AssumeRole API operation for STS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for STS's
// API operation AssumeRole for usage and error information.
func (c *STS) AssumeRole(input *AssumeRoleInput) (*AssumeRoleOutput, error) {
	req, out := c.AssumeRoleRequest(input)
	return out, req.Send()
}

// AssumeRoleWithContext is the same as AssumeRole with the addition of
// the ability to pass a context and additional request options.
//
// See AssumeRole for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *STS) AssumeRoleWithContext(ctx volcengine.Context, input *AssumeRoleInput, opts ...request.Option) (*AssumeRoleOutput, error) {
	req, out := c.AssumeRoleRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type AssumeRoleInput struct {
	_ struct{} `type:"structure"`

	DurationSeconds *int32 `type:"int32"`

	Policy *string `type:"string"`

	// RoleSessionName is a required field
	RoleSessionName *string `type:"string" required:"true"`

	// RoleTrn is a required field
	RoleTrn *string `type:"string" required:"true"`

	Tags []*TagForAssumeRoleInput `type:"list"`
}

// String returns the string representation
func (s AssumeRoleInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s AssumeRoleInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AssumeRoleInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "AssumeRoleInput"}
	if s.RoleSessionName == nil {
		invalidParams.Add(request.NewErrParamRequired("RoleSessionName"))
	}
	if s.RoleTrn == nil {
		invalidParams.Add(request.NewErrParamRequired("RoleTrn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetDurationSeconds sets the DurationSeconds field's value.
func (s *AssumeRoleInput) SetDurationSeconds(v int32) *AssumeRoleInput {
	s.DurationSeconds = &v
	return s
}

// SetPolicy sets the Policy field's value.
func (s *AssumeRoleInput) SetPolicy(v string) *AssumeRoleInput {
	s.Policy = &v
	return s
}

// SetRoleSessionName sets the RoleSessionName field's value.
func (s *AssumeRoleInput) SetRoleSessionName(v string) *AssumeRoleInput {
	s.RoleSessionName = &v
	return s
}

// SetRoleTrn sets the RoleTrn field's value.
func (s *AssumeRoleInput) SetRoleTrn(v string) *AssumeRoleInput {
	s.RoleTrn = &v
	return s
}

// SetTags sets the Tags field's value.
func (s *AssumeRoleInput) SetTags(v []*TagForAssumeRoleInput) *AssumeRoleInput {
	s.Tags = v
	return s
}

type AssumeRoleOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	AssumedRoleUser *AssumedRoleUserForAssumeRoleOutput `type:"structure"`

	Credentials *CredentialsForAssumeRoleOutput `type:"structure"`
}

// String returns the string representation
func (s AssumeRoleOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s AssumeRoleOutput) GoString() string {
	return s.String()
}

// SetAssumedRoleUser sets the AssumedRoleUser field's value.
func (s *AssumeRoleOutput) SetAssumedRoleUser(v *AssumedRoleUserForAssumeRoleOutput) *AssumeRoleOutput {
	s.AssumedRoleUser = v
	return s
}

// SetCredentials sets the Credentials field's value.
func (s *AssumeRoleOutput) SetCredentials(v *CredentialsForAssumeRoleOutput) *AssumeRoleOutput {
	s.Credentials = v
	return s
}

type AssumedRoleUserForAssumeRoleOutput struct {
	_ struct{} `type:"structure"`

	AssumedRoleId *string `type:"string"`

	Trn *string `type:"string"`
}

// String returns the string representation
func (s AssumedRoleUserForAssumeRoleOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s AssumedRoleUserForAssumeRoleOutput) GoString() string {
	return s.String()
}

// SetAssumedRoleId sets the AssumedRoleId field's value.
func (s *AssumedRoleUserForAssumeRoleOutput) SetAssumedRoleId(v string) *AssumedRoleUserForAssumeRoleOutput {
	s.AssumedRoleId = &v
	return s
}

// SetTrn sets the Trn field's value.
func (s *AssumedRoleUserForAssumeRoleOutput) SetTrn(v string) *AssumedRoleUserForAssumeRoleOutput {
	s.Trn = &v
	return s
}

type CredentialsForAssumeRoleOutput struct {
	_ struct{} `type:"structure"`

	AccessKeyId *string `type:"string"`

	CurrentTime *string `type:"string"`

	ExpiredTime *string `type:"string"`

	SecretAccessKey *string `type:"string"`

	SessionToken *string `type:"string"`
}

// String returns the string representation
func (s CredentialsForAssumeRoleOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CredentialsForAssumeRoleOutput) GoString() string {
	return s.String()
}

// SetAccessKeyId sets the AccessKeyId field's value.
func (s *CredentialsForAssumeRoleOutput) SetAccessKeyId(v string) *CredentialsForAssumeRoleOutput {
	s.AccessKeyId = &v
	return s
}

// SetCurrentTime sets the CurrentTime field's value.
func (s *CredentialsForAssumeRoleOutput) SetCurrentTime(v string) *CredentialsForAssumeRoleOutput {
	s.CurrentTime = &v
	return s
}

// SetExpiredTime sets the ExpiredTime field's value.
func (s *CredentialsForAssumeRoleOutput) SetExpiredTime(v string) *CredentialsForAssumeRoleOutput {
	s.ExpiredTime = &v
	return s
}

// SetSecretAccessKey sets the SecretAccessKey field's value.
func (s *CredentialsForAssumeRoleOutput) SetSecretAccessKey(v string) *CredentialsForAssumeRoleOutput {
	s.SecretAccessKey = &v
	return s
}

// SetSessionToken sets the SessionToken field's value.
func (s *CredentialsForAssumeRoleOutput) SetSessionToken(v string) *CredentialsForAssumeRoleOutput {
	s.SessionToken = &v
	return s
}

type TagForAssumeRoleInput struct {
	_ struct{} `type:"structure"`

	Key *string `type:"string"`

	Value *string `type:"string"`
}

// String returns the string representation
func (s TagForAssumeRoleInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s TagForAssumeRoleInput) GoString() string {
	return s.String()
}

// SetKey sets the Key field's value.
func (s *TagForAssumeRoleInput) SetKey(v string) *TagForAssumeRoleInput {
	s.Key = &v
	return s
}

// SetValue sets the Value field's value.
func (s *TagForAssumeRoleInput) SetValue(v string) *TagForAssumeRoleInput {
	s.Value = &v
	return s
}
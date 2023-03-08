// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package iam

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opUpdateLoginProfileCommon = "UpdateLoginProfile"

// UpdateLoginProfileCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the UpdateLoginProfileCommon operation. The "output" return
// value will be populated with the UpdateLoginProfileCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned UpdateLoginProfileCommon Request to send the API call to the service.
// the "output" return value is not valid until after UpdateLoginProfileCommon Send returns without error.
//
// See UpdateLoginProfileCommon for more information on using the UpdateLoginProfileCommon
// API call, and error handling.
//
//    // Example sending a request using the UpdateLoginProfileCommonRequest method.
//    req, resp := client.UpdateLoginProfileCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *IAM) UpdateLoginProfileCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opUpdateLoginProfileCommon,
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

// UpdateLoginProfileCommon API operation for IAM.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for IAM's
// API operation UpdateLoginProfileCommon for usage and error information.
func (c *IAM) UpdateLoginProfileCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.UpdateLoginProfileCommonRequest(input)
	return out, req.Send()
}

// UpdateLoginProfileCommonWithContext is the same as UpdateLoginProfileCommon with the addition of
// the ability to pass a context and additional request options.
//
// See UpdateLoginProfileCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *IAM) UpdateLoginProfileCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.UpdateLoginProfileCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opUpdateLoginProfile = "UpdateLoginProfile"

// UpdateLoginProfileRequest generates a "volcengine/request.Request" representing the
// client's request for the UpdateLoginProfile operation. The "output" return
// value will be populated with the UpdateLoginProfileCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned UpdateLoginProfileCommon Request to send the API call to the service.
// the "output" return value is not valid until after UpdateLoginProfileCommon Send returns without error.
//
// See UpdateLoginProfile for more information on using the UpdateLoginProfile
// API call, and error handling.
//
//    // Example sending a request using the UpdateLoginProfileRequest method.
//    req, resp := client.UpdateLoginProfileRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *IAM) UpdateLoginProfileRequest(input *UpdateLoginProfileInput) (req *request.Request, output *UpdateLoginProfileOutput) {
	op := &request.Operation{
		Name:       opUpdateLoginProfile,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateLoginProfileInput{}
	}

	output = &UpdateLoginProfileOutput{}
	req = c.newRequest(op, input, output)

	return
}

// UpdateLoginProfile API operation for IAM.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for IAM's
// API operation UpdateLoginProfile for usage and error information.
func (c *IAM) UpdateLoginProfile(input *UpdateLoginProfileInput) (*UpdateLoginProfileOutput, error) {
	req, out := c.UpdateLoginProfileRequest(input)
	return out, req.Send()
}

// UpdateLoginProfileWithContext is the same as UpdateLoginProfile with the addition of
// the ability to pass a context and additional request options.
//
// See UpdateLoginProfile for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *IAM) UpdateLoginProfileWithContext(ctx volcengine.Context, input *UpdateLoginProfileInput, opts ...request.Option) (*UpdateLoginProfileOutput, error) {
	req, out := c.UpdateLoginProfileRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type LoginProfileForUpdateLoginProfileOutput struct {
	_ struct{} `type:"structure"`

	LastLoginDate *string `type:"string"`

	LastLoginIp *string `type:"string"`

	LoginAllowed *bool `type:"boolean"`

	PasswordResetRequired *bool `type:"boolean"`

	UserName *string `type:"string"`
}

// String returns the string representation
func (s LoginProfileForUpdateLoginProfileOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s LoginProfileForUpdateLoginProfileOutput) GoString() string {
	return s.String()
}

// SetLastLoginDate sets the LastLoginDate field's value.
func (s *LoginProfileForUpdateLoginProfileOutput) SetLastLoginDate(v string) *LoginProfileForUpdateLoginProfileOutput {
	s.LastLoginDate = &v
	return s
}

// SetLastLoginIp sets the LastLoginIp field's value.
func (s *LoginProfileForUpdateLoginProfileOutput) SetLastLoginIp(v string) *LoginProfileForUpdateLoginProfileOutput {
	s.LastLoginIp = &v
	return s
}

// SetLoginAllowed sets the LoginAllowed field's value.
func (s *LoginProfileForUpdateLoginProfileOutput) SetLoginAllowed(v bool) *LoginProfileForUpdateLoginProfileOutput {
	s.LoginAllowed = &v
	return s
}

// SetPasswordResetRequired sets the PasswordResetRequired field's value.
func (s *LoginProfileForUpdateLoginProfileOutput) SetPasswordResetRequired(v bool) *LoginProfileForUpdateLoginProfileOutput {
	s.PasswordResetRequired = &v
	return s
}

// SetUserName sets the UserName field's value.
func (s *LoginProfileForUpdateLoginProfileOutput) SetUserName(v string) *LoginProfileForUpdateLoginProfileOutput {
	s.UserName = &v
	return s
}

type UpdateLoginProfileInput struct {
	_ struct{} `type:"structure"`

	LoginAllowed *bool `type:"boolean"`

	Password *string `type:"string"`

	PasswordResetRequired *bool `type:"boolean"`

	SafeAuthExemptDuration *int64 `type:"integer"`

	SafeAuthFlag *bool `type:"boolean"`

	SafeAuthType *string `type:"string"`

	// UserName is a required field
	UserName *string `type:"string" required:"true"`
}

// String returns the string representation
func (s UpdateLoginProfileInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s UpdateLoginProfileInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateLoginProfileInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "UpdateLoginProfileInput"}
	if s.UserName == nil {
		invalidParams.Add(request.NewErrParamRequired("UserName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetLoginAllowed sets the LoginAllowed field's value.
func (s *UpdateLoginProfileInput) SetLoginAllowed(v bool) *UpdateLoginProfileInput {
	s.LoginAllowed = &v
	return s
}

// SetPassword sets the Password field's value.
func (s *UpdateLoginProfileInput) SetPassword(v string) *UpdateLoginProfileInput {
	s.Password = &v
	return s
}

// SetPasswordResetRequired sets the PasswordResetRequired field's value.
func (s *UpdateLoginProfileInput) SetPasswordResetRequired(v bool) *UpdateLoginProfileInput {
	s.PasswordResetRequired = &v
	return s
}

// SetSafeAuthExemptDuration sets the SafeAuthExemptDuration field's value.
func (s *UpdateLoginProfileInput) SetSafeAuthExemptDuration(v int64) *UpdateLoginProfileInput {
	s.SafeAuthExemptDuration = &v
	return s
}

// SetSafeAuthFlag sets the SafeAuthFlag field's value.
func (s *UpdateLoginProfileInput) SetSafeAuthFlag(v bool) *UpdateLoginProfileInput {
	s.SafeAuthFlag = &v
	return s
}

// SetSafeAuthType sets the SafeAuthType field's value.
func (s *UpdateLoginProfileInput) SetSafeAuthType(v string) *UpdateLoginProfileInput {
	s.SafeAuthType = &v
	return s
}

// SetUserName sets the UserName field's value.
func (s *UpdateLoginProfileInput) SetUserName(v string) *UpdateLoginProfileInput {
	s.UserName = &v
	return s
}

type UpdateLoginProfileOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	LoginProfile *LoginProfileForUpdateLoginProfileOutput `type:"structure"`
}

// String returns the string representation
func (s UpdateLoginProfileOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s UpdateLoginProfileOutput) GoString() string {
	return s.String()
}

// SetLoginProfile sets the LoginProfile field's value.
func (s *UpdateLoginProfileOutput) SetLoginProfile(v *LoginProfileForUpdateLoginProfileOutput) *UpdateLoginProfileOutput {
	s.LoginProfile = v
	return s
}

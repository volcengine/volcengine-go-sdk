// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package iam

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opAttachUserGroupPolicyCommon = "AttachUserGroupPolicy"

// AttachUserGroupPolicyCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the AttachUserGroupPolicyCommon operation. The "output" return
// value will be populated with the AttachUserGroupPolicyCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned AttachUserGroupPolicyCommon Request to send the API call to the service.
// the "output" return value is not valid until after AttachUserGroupPolicyCommon Send returns without error.
//
// See AttachUserGroupPolicyCommon for more information on using the AttachUserGroupPolicyCommon
// API call, and error handling.
//
//    // Example sending a request using the AttachUserGroupPolicyCommonRequest method.
//    req, resp := client.AttachUserGroupPolicyCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *IAM) AttachUserGroupPolicyCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opAttachUserGroupPolicyCommon,
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

// AttachUserGroupPolicyCommon API operation for IAM.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for IAM's
// API operation AttachUserGroupPolicyCommon for usage and error information.
func (c *IAM) AttachUserGroupPolicyCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.AttachUserGroupPolicyCommonRequest(input)
	return out, req.Send()
}

// AttachUserGroupPolicyCommonWithContext is the same as AttachUserGroupPolicyCommon with the addition of
// the ability to pass a context and additional request options.
//
// See AttachUserGroupPolicyCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *IAM) AttachUserGroupPolicyCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.AttachUserGroupPolicyCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opAttachUserGroupPolicy = "AttachUserGroupPolicy"

// AttachUserGroupPolicyRequest generates a "volcengine/request.Request" representing the
// client's request for the AttachUserGroupPolicy operation. The "output" return
// value will be populated with the AttachUserGroupPolicyCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned AttachUserGroupPolicyCommon Request to send the API call to the service.
// the "output" return value is not valid until after AttachUserGroupPolicyCommon Send returns without error.
//
// See AttachUserGroupPolicy for more information on using the AttachUserGroupPolicy
// API call, and error handling.
//
//    // Example sending a request using the AttachUserGroupPolicyRequest method.
//    req, resp := client.AttachUserGroupPolicyRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *IAM) AttachUserGroupPolicyRequest(input *AttachUserGroupPolicyInput) (req *request.Request, output *AttachUserGroupPolicyOutput) {
	op := &request.Operation{
		Name:       opAttachUserGroupPolicy,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &AttachUserGroupPolicyInput{}
	}

	output = &AttachUserGroupPolicyOutput{}
	req = c.newRequest(op, input, output)

	return
}

// AttachUserGroupPolicy API operation for IAM.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for IAM's
// API operation AttachUserGroupPolicy for usage and error information.
func (c *IAM) AttachUserGroupPolicy(input *AttachUserGroupPolicyInput) (*AttachUserGroupPolicyOutput, error) {
	req, out := c.AttachUserGroupPolicyRequest(input)
	return out, req.Send()
}

// AttachUserGroupPolicyWithContext is the same as AttachUserGroupPolicy with the addition of
// the ability to pass a context and additional request options.
//
// See AttachUserGroupPolicy for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *IAM) AttachUserGroupPolicyWithContext(ctx volcengine.Context, input *AttachUserGroupPolicyInput, opts ...request.Option) (*AttachUserGroupPolicyOutput, error) {
	req, out := c.AttachUserGroupPolicyRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type AttachUserGroupPolicyInput struct {
	_ struct{} `type:"structure"`

	// PolicyName is a required field
	PolicyName *string `min:"1" max:"64" type:"string" required:"true"`

	// PolicyType is a required field
	PolicyType *string `type:"string" required:"true" enum:"EnumOfPolicyTypeForAttachUserGroupPolicyInput"`

	// UserGroupName is a required field
	UserGroupName *string `min:"1" max:"64" type:"string" required:"true"`
}

// String returns the string representation
func (s AttachUserGroupPolicyInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s AttachUserGroupPolicyInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AttachUserGroupPolicyInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "AttachUserGroupPolicyInput"}
	if s.PolicyName == nil {
		invalidParams.Add(request.NewErrParamRequired("PolicyName"))
	}
	if s.PolicyName != nil && len(*s.PolicyName) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("PolicyName", 1))
	}
	if s.PolicyName != nil && len(*s.PolicyName) > 64 {
		invalidParams.Add(request.NewErrParamMaxLen("PolicyName", 64, *s.PolicyName))
	}
	if s.PolicyType == nil {
		invalidParams.Add(request.NewErrParamRequired("PolicyType"))
	}
	if s.UserGroupName == nil {
		invalidParams.Add(request.NewErrParamRequired("UserGroupName"))
	}
	if s.UserGroupName != nil && len(*s.UserGroupName) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("UserGroupName", 1))
	}
	if s.UserGroupName != nil && len(*s.UserGroupName) > 64 {
		invalidParams.Add(request.NewErrParamMaxLen("UserGroupName", 64, *s.UserGroupName))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetPolicyName sets the PolicyName field's value.
func (s *AttachUserGroupPolicyInput) SetPolicyName(v string) *AttachUserGroupPolicyInput {
	s.PolicyName = &v
	return s
}

// SetPolicyType sets the PolicyType field's value.
func (s *AttachUserGroupPolicyInput) SetPolicyType(v string) *AttachUserGroupPolicyInput {
	s.PolicyType = &v
	return s
}

// SetUserGroupName sets the UserGroupName field's value.
func (s *AttachUserGroupPolicyInput) SetUserGroupName(v string) *AttachUserGroupPolicyInput {
	s.UserGroupName = &v
	return s
}

type AttachUserGroupPolicyOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s AttachUserGroupPolicyOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s AttachUserGroupPolicyOutput) GoString() string {
	return s.String()
}

const (
	// EnumOfPolicyTypeForAttachUserGroupPolicyInputSystem is a EnumOfPolicyTypeForAttachUserGroupPolicyInput enum value
	EnumOfPolicyTypeForAttachUserGroupPolicyInputSystem = "System"

	// EnumOfPolicyTypeForAttachUserGroupPolicyInputCustom is a EnumOfPolicyTypeForAttachUserGroupPolicyInput enum value
	EnumOfPolicyTypeForAttachUserGroupPolicyInputCustom = "Custom"
)

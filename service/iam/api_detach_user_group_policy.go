// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package iam

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDetachUserGroupPolicyCommon = "DetachUserGroupPolicy"

// DetachUserGroupPolicyCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DetachUserGroupPolicyCommon operation. The "output" return
// value will be populated with the DetachUserGroupPolicyCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DetachUserGroupPolicyCommon Request to send the API call to the service.
// the "output" return value is not valid until after DetachUserGroupPolicyCommon Send returns without error.
//
// See DetachUserGroupPolicyCommon for more information on using the DetachUserGroupPolicyCommon
// API call, and error handling.
//
//    // Example sending a request using the DetachUserGroupPolicyCommonRequest method.
//    req, resp := client.DetachUserGroupPolicyCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *IAM) DetachUserGroupPolicyCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDetachUserGroupPolicyCommon,
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

// DetachUserGroupPolicyCommon API operation for IAM.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for IAM's
// API operation DetachUserGroupPolicyCommon for usage and error information.
func (c *IAM) DetachUserGroupPolicyCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DetachUserGroupPolicyCommonRequest(input)
	return out, req.Send()
}

// DetachUserGroupPolicyCommonWithContext is the same as DetachUserGroupPolicyCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DetachUserGroupPolicyCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *IAM) DetachUserGroupPolicyCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DetachUserGroupPolicyCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDetachUserGroupPolicy = "DetachUserGroupPolicy"

// DetachUserGroupPolicyRequest generates a "volcengine/request.Request" representing the
// client's request for the DetachUserGroupPolicy operation. The "output" return
// value will be populated with the DetachUserGroupPolicyCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DetachUserGroupPolicyCommon Request to send the API call to the service.
// the "output" return value is not valid until after DetachUserGroupPolicyCommon Send returns without error.
//
// See DetachUserGroupPolicy for more information on using the DetachUserGroupPolicy
// API call, and error handling.
//
//    // Example sending a request using the DetachUserGroupPolicyRequest method.
//    req, resp := client.DetachUserGroupPolicyRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *IAM) DetachUserGroupPolicyRequest(input *DetachUserGroupPolicyInput) (req *request.Request, output *DetachUserGroupPolicyOutput) {
	op := &request.Operation{
		Name:       opDetachUserGroupPolicy,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DetachUserGroupPolicyInput{}
	}

	output = &DetachUserGroupPolicyOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DetachUserGroupPolicy API operation for IAM.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for IAM's
// API operation DetachUserGroupPolicy for usage and error information.
func (c *IAM) DetachUserGroupPolicy(input *DetachUserGroupPolicyInput) (*DetachUserGroupPolicyOutput, error) {
	req, out := c.DetachUserGroupPolicyRequest(input)
	return out, req.Send()
}

// DetachUserGroupPolicyWithContext is the same as DetachUserGroupPolicy with the addition of
// the ability to pass a context and additional request options.
//
// See DetachUserGroupPolicy for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *IAM) DetachUserGroupPolicyWithContext(ctx volcengine.Context, input *DetachUserGroupPolicyInput, opts ...request.Option) (*DetachUserGroupPolicyOutput, error) {
	req, out := c.DetachUserGroupPolicyRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DetachUserGroupPolicyInput struct {
	_ struct{} `type:"structure"`

	// PolicyName is a required field
	PolicyName *string `min:"1" max:"64" type:"string" required:"true"`

	// PolicyType is a required field
	PolicyType *string `type:"string" required:"true" enum:"EnumOfPolicyTypeForDetachUserGroupPolicyInput"`

	// UserGroupName is a required field
	UserGroupName *string `min:"1" max:"64" type:"string" required:"true"`
}

// String returns the string representation
func (s DetachUserGroupPolicyInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DetachUserGroupPolicyInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DetachUserGroupPolicyInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DetachUserGroupPolicyInput"}
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
func (s *DetachUserGroupPolicyInput) SetPolicyName(v string) *DetachUserGroupPolicyInput {
	s.PolicyName = &v
	return s
}

// SetPolicyType sets the PolicyType field's value.
func (s *DetachUserGroupPolicyInput) SetPolicyType(v string) *DetachUserGroupPolicyInput {
	s.PolicyType = &v
	return s
}

// SetUserGroupName sets the UserGroupName field's value.
func (s *DetachUserGroupPolicyInput) SetUserGroupName(v string) *DetachUserGroupPolicyInput {
	s.UserGroupName = &v
	return s
}

type DetachUserGroupPolicyOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s DetachUserGroupPolicyOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DetachUserGroupPolicyOutput) GoString() string {
	return s.String()
}

const (
	// EnumOfPolicyTypeForDetachUserGroupPolicyInputSystem is a EnumOfPolicyTypeForDetachUserGroupPolicyInput enum value
	EnumOfPolicyTypeForDetachUserGroupPolicyInputSystem = "System"

	// EnumOfPolicyTypeForDetachUserGroupPolicyInputCustom is a EnumOfPolicyTypeForDetachUserGroupPolicyInput enum value
	EnumOfPolicyTypeForDetachUserGroupPolicyInputCustom = "Custom"
)

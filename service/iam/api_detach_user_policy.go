// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package iam

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDetachUserPolicyCommon = "DetachUserPolicy"

// DetachUserPolicyCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DetachUserPolicyCommon operation. The "output" return
// value will be populated with the DetachUserPolicyCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DetachUserPolicyCommon Request to send the API call to the service.
// the "output" return value is not valid until after DetachUserPolicyCommon Send returns without error.
//
// See DetachUserPolicyCommon for more information on using the DetachUserPolicyCommon
// API call, and error handling.
//
//    // Example sending a request using the DetachUserPolicyCommonRequest method.
//    req, resp := client.DetachUserPolicyCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *IAM) DetachUserPolicyCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDetachUserPolicyCommon,
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

// DetachUserPolicyCommon API operation for IAM.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for IAM's
// API operation DetachUserPolicyCommon for usage and error information.
func (c *IAM) DetachUserPolicyCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DetachUserPolicyCommonRequest(input)
	return out, req.Send()
}

// DetachUserPolicyCommonWithContext is the same as DetachUserPolicyCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DetachUserPolicyCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *IAM) DetachUserPolicyCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DetachUserPolicyCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDetachUserPolicy = "DetachUserPolicy"

// DetachUserPolicyRequest generates a "volcengine/request.Request" representing the
// client's request for the DetachUserPolicy operation. The "output" return
// value will be populated with the DetachUserPolicyCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DetachUserPolicyCommon Request to send the API call to the service.
// the "output" return value is not valid until after DetachUserPolicyCommon Send returns without error.
//
// See DetachUserPolicy for more information on using the DetachUserPolicy
// API call, and error handling.
//
//    // Example sending a request using the DetachUserPolicyRequest method.
//    req, resp := client.DetachUserPolicyRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *IAM) DetachUserPolicyRequest(input *DetachUserPolicyInput) (req *request.Request, output *DetachUserPolicyOutput) {
	op := &request.Operation{
		Name:       opDetachUserPolicy,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DetachUserPolicyInput{}
	}

	output = &DetachUserPolicyOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DetachUserPolicy API operation for IAM.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for IAM's
// API operation DetachUserPolicy for usage and error information.
func (c *IAM) DetachUserPolicy(input *DetachUserPolicyInput) (*DetachUserPolicyOutput, error) {
	req, out := c.DetachUserPolicyRequest(input)
	return out, req.Send()
}

// DetachUserPolicyWithContext is the same as DetachUserPolicy with the addition of
// the ability to pass a context and additional request options.
//
// See DetachUserPolicy for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *IAM) DetachUserPolicyWithContext(ctx volcengine.Context, input *DetachUserPolicyInput, opts ...request.Option) (*DetachUserPolicyOutput, error) {
	req, out := c.DetachUserPolicyRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DetachUserPolicyInput struct {
	_ struct{} `type:"structure"`

	// PolicyName is a required field
	PolicyName *string `type:"string" required:"true"`

	// PolicyType is a required field
	PolicyType *string `type:"string" required:"true"`

	// UserName is a required field
	UserName *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DetachUserPolicyInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DetachUserPolicyInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DetachUserPolicyInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DetachUserPolicyInput"}
	if s.PolicyName == nil {
		invalidParams.Add(request.NewErrParamRequired("PolicyName"))
	}
	if s.PolicyType == nil {
		invalidParams.Add(request.NewErrParamRequired("PolicyType"))
	}
	if s.UserName == nil {
		invalidParams.Add(request.NewErrParamRequired("UserName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetPolicyName sets the PolicyName field's value.
func (s *DetachUserPolicyInput) SetPolicyName(v string) *DetachUserPolicyInput {
	s.PolicyName = &v
	return s
}

// SetPolicyType sets the PolicyType field's value.
func (s *DetachUserPolicyInput) SetPolicyType(v string) *DetachUserPolicyInput {
	s.PolicyType = &v
	return s
}

// SetUserName sets the UserName field's value.
func (s *DetachUserPolicyInput) SetUserName(v string) *DetachUserPolicyInput {
	s.UserName = &v
	return s
}

type DetachUserPolicyOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s DetachUserPolicyOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DetachUserPolicyOutput) GoString() string {
	return s.String()
}

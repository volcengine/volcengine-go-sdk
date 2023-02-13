// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package iam

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opAttachUserPolicyCommon = "AttachUserPolicy"

// AttachUserPolicyCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the AttachUserPolicyCommon operation. The "output" return
// value will be populated with the AttachUserPolicyCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned AttachUserPolicyCommon Request to send the API call to the service.
// the "output" return value is not valid until after AttachUserPolicyCommon Send returns without error.
//
// See AttachUserPolicyCommon for more information on using the AttachUserPolicyCommon
// API call, and error handling.
//
//    // Example sending a request using the AttachUserPolicyCommonRequest method.
//    req, resp := client.AttachUserPolicyCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *IAM) AttachUserPolicyCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opAttachUserPolicyCommon,
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

// AttachUserPolicyCommon API operation for IAM.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for IAM's
// API operation AttachUserPolicyCommon for usage and error information.
func (c *IAM) AttachUserPolicyCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.AttachUserPolicyCommonRequest(input)
	return out, req.Send()
}

// AttachUserPolicyCommonWithContext is the same as AttachUserPolicyCommon with the addition of
// the ability to pass a context and additional request options.
//
// See AttachUserPolicyCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *IAM) AttachUserPolicyCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.AttachUserPolicyCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opAttachUserPolicy = "AttachUserPolicy"

// AttachUserPolicyRequest generates a "volcengine/request.Request" representing the
// client's request for the AttachUserPolicy operation. The "output" return
// value will be populated with the AttachUserPolicyCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned AttachUserPolicyCommon Request to send the API call to the service.
// the "output" return value is not valid until after AttachUserPolicyCommon Send returns without error.
//
// See AttachUserPolicy for more information on using the AttachUserPolicy
// API call, and error handling.
//
//    // Example sending a request using the AttachUserPolicyRequest method.
//    req, resp := client.AttachUserPolicyRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *IAM) AttachUserPolicyRequest(input *AttachUserPolicyInput) (req *request.Request, output *AttachUserPolicyOutput) {
	op := &request.Operation{
		Name:       opAttachUserPolicy,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &AttachUserPolicyInput{}
	}

	output = &AttachUserPolicyOutput{}
	req = c.newRequest(op, input, output)

	return
}

// AttachUserPolicy API operation for IAM.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for IAM's
// API operation AttachUserPolicy for usage and error information.
func (c *IAM) AttachUserPolicy(input *AttachUserPolicyInput) (*AttachUserPolicyOutput, error) {
	req, out := c.AttachUserPolicyRequest(input)
	return out, req.Send()
}

// AttachUserPolicyWithContext is the same as AttachUserPolicy with the addition of
// the ability to pass a context and additional request options.
//
// See AttachUserPolicy for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *IAM) AttachUserPolicyWithContext(ctx volcengine.Context, input *AttachUserPolicyInput, opts ...request.Option) (*AttachUserPolicyOutput, error) {
	req, out := c.AttachUserPolicyRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type AttachUserPolicyInput struct {
	_ struct{} `type:"structure"`

	// PolicyName is a required field
	PolicyName *string `type:"string" required:"true"`

	// PolicyType is a required field
	PolicyType *string `type:"string" required:"true"`

	// UserName is a required field
	UserName *string `type:"string" required:"true"`
}

// String returns the string representation
func (s AttachUserPolicyInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s AttachUserPolicyInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AttachUserPolicyInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "AttachUserPolicyInput"}
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
func (s *AttachUserPolicyInput) SetPolicyName(v string) *AttachUserPolicyInput {
	s.PolicyName = &v
	return s
}

// SetPolicyType sets the PolicyType field's value.
func (s *AttachUserPolicyInput) SetPolicyType(v string) *AttachUserPolicyInput {
	s.PolicyType = &v
	return s
}

// SetUserName sets the UserName field's value.
func (s *AttachUserPolicyInput) SetUserName(v string) *AttachUserPolicyInput {
	s.UserName = &v
	return s
}

type AttachUserPolicyOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	ResponseMetadata *interface{} `type:"interface"`
}

// String returns the string representation
func (s AttachUserPolicyOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s AttachUserPolicyOutput) GoString() string {
	return s.String()
}

// SetResponseMetadata sets the ResponseMetadata field's value.
func (s *AttachUserPolicyOutput) SetResponseMetadata(v interface{}) *AttachUserPolicyOutput {
	s.ResponseMetadata = &v
	return s
}

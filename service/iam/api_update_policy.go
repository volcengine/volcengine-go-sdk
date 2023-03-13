// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package iam

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opUpdatePolicyCommon = "UpdatePolicy"

// UpdatePolicyCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the UpdatePolicyCommon operation. The "output" return
// value will be populated with the UpdatePolicyCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned UpdatePolicyCommon Request to send the API call to the service.
// the "output" return value is not valid until after UpdatePolicyCommon Send returns without error.
//
// See UpdatePolicyCommon for more information on using the UpdatePolicyCommon
// API call, and error handling.
//
//    // Example sending a request using the UpdatePolicyCommonRequest method.
//    req, resp := client.UpdatePolicyCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *IAM) UpdatePolicyCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opUpdatePolicyCommon,
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

// UpdatePolicyCommon API operation for IAM.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for IAM's
// API operation UpdatePolicyCommon for usage and error information.
func (c *IAM) UpdatePolicyCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.UpdatePolicyCommonRequest(input)
	return out, req.Send()
}

// UpdatePolicyCommonWithContext is the same as UpdatePolicyCommon with the addition of
// the ability to pass a context and additional request options.
//
// See UpdatePolicyCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *IAM) UpdatePolicyCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.UpdatePolicyCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opUpdatePolicy = "UpdatePolicy"

// UpdatePolicyRequest generates a "volcengine/request.Request" representing the
// client's request for the UpdatePolicy operation. The "output" return
// value will be populated with the UpdatePolicyCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned UpdatePolicyCommon Request to send the API call to the service.
// the "output" return value is not valid until after UpdatePolicyCommon Send returns without error.
//
// See UpdatePolicy for more information on using the UpdatePolicy
// API call, and error handling.
//
//    // Example sending a request using the UpdatePolicyRequest method.
//    req, resp := client.UpdatePolicyRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *IAM) UpdatePolicyRequest(input *UpdatePolicyInput) (req *request.Request, output *UpdatePolicyOutput) {
	op := &request.Operation{
		Name:       opUpdatePolicy,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdatePolicyInput{}
	}

	output = &UpdatePolicyOutput{}
	req = c.newRequest(op, input, output)

	return
}

// UpdatePolicy API operation for IAM.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for IAM's
// API operation UpdatePolicy for usage and error information.
func (c *IAM) UpdatePolicy(input *UpdatePolicyInput) (*UpdatePolicyOutput, error) {
	req, out := c.UpdatePolicyRequest(input)
	return out, req.Send()
}

// UpdatePolicyWithContext is the same as UpdatePolicy with the addition of
// the ability to pass a context and additional request options.
//
// See UpdatePolicy for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *IAM) UpdatePolicyWithContext(ctx volcengine.Context, input *UpdatePolicyInput, opts ...request.Option) (*UpdatePolicyOutput, error) {
	req, out := c.UpdatePolicyRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type PolicyForUpdatePolicyOutput struct {
	_ struct{} `type:"structure"`

	CreateDate *string `type:"string"`

	Description *string `type:"string"`

	PolicyDocument *string `type:"string"`

	PolicyName *string `type:"string"`

	PolicyTrn *string `type:"string"`

	PolicyType *string `type:"string"`

	UpdateDate *string `type:"string"`
}

// String returns the string representation
func (s PolicyForUpdatePolicyOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s PolicyForUpdatePolicyOutput) GoString() string {
	return s.String()
}

// SetCreateDate sets the CreateDate field's value.
func (s *PolicyForUpdatePolicyOutput) SetCreateDate(v string) *PolicyForUpdatePolicyOutput {
	s.CreateDate = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *PolicyForUpdatePolicyOutput) SetDescription(v string) *PolicyForUpdatePolicyOutput {
	s.Description = &v
	return s
}

// SetPolicyDocument sets the PolicyDocument field's value.
func (s *PolicyForUpdatePolicyOutput) SetPolicyDocument(v string) *PolicyForUpdatePolicyOutput {
	s.PolicyDocument = &v
	return s
}

// SetPolicyName sets the PolicyName field's value.
func (s *PolicyForUpdatePolicyOutput) SetPolicyName(v string) *PolicyForUpdatePolicyOutput {
	s.PolicyName = &v
	return s
}

// SetPolicyTrn sets the PolicyTrn field's value.
func (s *PolicyForUpdatePolicyOutput) SetPolicyTrn(v string) *PolicyForUpdatePolicyOutput {
	s.PolicyTrn = &v
	return s
}

// SetPolicyType sets the PolicyType field's value.
func (s *PolicyForUpdatePolicyOutput) SetPolicyType(v string) *PolicyForUpdatePolicyOutput {
	s.PolicyType = &v
	return s
}

// SetUpdateDate sets the UpdateDate field's value.
func (s *PolicyForUpdatePolicyOutput) SetUpdateDate(v string) *PolicyForUpdatePolicyOutput {
	s.UpdateDate = &v
	return s
}

type UpdatePolicyInput struct {
	_ struct{} `type:"structure"`

	NewDescription *string `type:"string"`

	NewPolicyDocument *string `type:"string"`

	NewPolicyName *string `type:"string"`

	// PolicyName is a required field
	PolicyName *string `type:"string" required:"true"`
}

// String returns the string representation
func (s UpdatePolicyInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s UpdatePolicyInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdatePolicyInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "UpdatePolicyInput"}
	if s.PolicyName == nil {
		invalidParams.Add(request.NewErrParamRequired("PolicyName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetNewDescription sets the NewDescription field's value.
func (s *UpdatePolicyInput) SetNewDescription(v string) *UpdatePolicyInput {
	s.NewDescription = &v
	return s
}

// SetNewPolicyDocument sets the NewPolicyDocument field's value.
func (s *UpdatePolicyInput) SetNewPolicyDocument(v string) *UpdatePolicyInput {
	s.NewPolicyDocument = &v
	return s
}

// SetNewPolicyName sets the NewPolicyName field's value.
func (s *UpdatePolicyInput) SetNewPolicyName(v string) *UpdatePolicyInput {
	s.NewPolicyName = &v
	return s
}

// SetPolicyName sets the PolicyName field's value.
func (s *UpdatePolicyInput) SetPolicyName(v string) *UpdatePolicyInput {
	s.PolicyName = &v
	return s
}

type UpdatePolicyOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	Policy *PolicyForUpdatePolicyOutput `type:"structure"`
}

// String returns the string representation
func (s UpdatePolicyOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s UpdatePolicyOutput) GoString() string {
	return s.String()
}

// SetPolicy sets the Policy field's value.
func (s *UpdatePolicyOutput) SetPolicy(v *PolicyForUpdatePolicyOutput) *UpdatePolicyOutput {
	s.Policy = v
	return s
}
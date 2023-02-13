// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package iam

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opListPoliciesCommon = "ListPolicies"

// ListPoliciesCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ListPoliciesCommon operation. The "output" return
// value will be populated with the ListPoliciesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListPoliciesCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListPoliciesCommon Send returns without error.
//
// See ListPoliciesCommon for more information on using the ListPoliciesCommon
// API call, and error handling.
//
//    // Example sending a request using the ListPoliciesCommonRequest method.
//    req, resp := client.ListPoliciesCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *IAM) ListPoliciesCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opListPoliciesCommon,
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

// ListPoliciesCommon API operation for IAM.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for IAM's
// API operation ListPoliciesCommon for usage and error information.
func (c *IAM) ListPoliciesCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ListPoliciesCommonRequest(input)
	return out, req.Send()
}

// ListPoliciesCommonWithContext is the same as ListPoliciesCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ListPoliciesCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *IAM) ListPoliciesCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ListPoliciesCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opListPolicies = "ListPolicies"

// ListPoliciesRequest generates a "volcengine/request.Request" representing the
// client's request for the ListPolicies operation. The "output" return
// value will be populated with the ListPoliciesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListPoliciesCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListPoliciesCommon Send returns without error.
//
// See ListPolicies for more information on using the ListPolicies
// API call, and error handling.
//
//    // Example sending a request using the ListPoliciesRequest method.
//    req, resp := client.ListPoliciesRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *IAM) ListPoliciesRequest(input *ListPoliciesInput) (req *request.Request, output *ListPoliciesOutput) {
	op := &request.Operation{
		Name:       opListPolicies,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListPoliciesInput{}
	}

	output = &ListPoliciesOutput{}
	req = c.newRequest(op, input, output)

	return
}

// ListPolicies API operation for IAM.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for IAM's
// API operation ListPolicies for usage and error information.
func (c *IAM) ListPolicies(input *ListPoliciesInput) (*ListPoliciesOutput, error) {
	req, out := c.ListPoliciesRequest(input)
	return out, req.Send()
}

// ListPoliciesWithContext is the same as ListPolicies with the addition of
// the ability to pass a context and additional request options.
//
// See ListPolicies for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *IAM) ListPoliciesWithContext(ctx volcengine.Context, input *ListPoliciesInput, opts ...request.Option) (*ListPoliciesOutput, error) {
	req, out := c.ListPoliciesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ListPoliciesInput struct {
	_ struct{} `type:"structure"`

	Scope *string `type:"string"`
}

// String returns the string representation
func (s ListPoliciesInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListPoliciesInput) GoString() string {
	return s.String()
}

// SetScope sets the Scope field's value.
func (s *ListPoliciesInput) SetScope(v string) *ListPoliciesInput {
	s.Scope = &v
	return s
}

type ListPoliciesOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	ResponseMetadata *interface{} `type:"interface"`

	Result *interface{} `type:"interface"`
}

// String returns the string representation
func (s ListPoliciesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListPoliciesOutput) GoString() string {
	return s.String()
}

// SetResponseMetadata sets the ResponseMetadata field's value.
func (s *ListPoliciesOutput) SetResponseMetadata(v interface{}) *ListPoliciesOutput {
	s.ResponseMetadata = &v
	return s
}

// SetResult sets the Result field's value.
func (s *ListPoliciesOutput) SetResult(v interface{}) *ListPoliciesOutput {
	s.Result = &v
	return s
}

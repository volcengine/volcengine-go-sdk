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

type UpdateLoginProfileInput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s UpdateLoginProfileInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s UpdateLoginProfileInput) GoString() string {
	return s.String()
}

type UpdateLoginProfileOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	ResponseMetadata *interface{} `type:"interface"`
}

// String returns the string representation
func (s UpdateLoginProfileOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s UpdateLoginProfileOutput) GoString() string {
	return s.String()
}

// SetResponseMetadata sets the ResponseMetadata field's value.
func (s *UpdateLoginProfileOutput) SetResponseMetadata(v interface{}) *UpdateLoginProfileOutput {
	s.ResponseMetadata = &v
	return s
}

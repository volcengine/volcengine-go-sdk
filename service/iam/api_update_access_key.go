// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package iam

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opUpdateAccessKeyCommon = "UpdateAccessKey"

// UpdateAccessKeyCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the UpdateAccessKeyCommon operation. The "output" return
// value will be populated with the UpdateAccessKeyCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned UpdateAccessKeyCommon Request to send the API call to the service.
// the "output" return value is not valid until after UpdateAccessKeyCommon Send returns without error.
//
// See UpdateAccessKeyCommon for more information on using the UpdateAccessKeyCommon
// API call, and error handling.
//
//    // Example sending a request using the UpdateAccessKeyCommonRequest method.
//    req, resp := client.UpdateAccessKeyCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *IAM) UpdateAccessKeyCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opUpdateAccessKeyCommon,
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

// UpdateAccessKeyCommon API operation for IAM.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for IAM's
// API operation UpdateAccessKeyCommon for usage and error information.
func (c *IAM) UpdateAccessKeyCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.UpdateAccessKeyCommonRequest(input)
	return out, req.Send()
}

// UpdateAccessKeyCommonWithContext is the same as UpdateAccessKeyCommon with the addition of
// the ability to pass a context and additional request options.
//
// See UpdateAccessKeyCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *IAM) UpdateAccessKeyCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.UpdateAccessKeyCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opUpdateAccessKey = "UpdateAccessKey"

// UpdateAccessKeyRequest generates a "volcengine/request.Request" representing the
// client's request for the UpdateAccessKey operation. The "output" return
// value will be populated with the UpdateAccessKeyCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned UpdateAccessKeyCommon Request to send the API call to the service.
// the "output" return value is not valid until after UpdateAccessKeyCommon Send returns without error.
//
// See UpdateAccessKey for more information on using the UpdateAccessKey
// API call, and error handling.
//
//    // Example sending a request using the UpdateAccessKeyRequest method.
//    req, resp := client.UpdateAccessKeyRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *IAM) UpdateAccessKeyRequest(input *UpdateAccessKeyInput) (req *request.Request, output *UpdateAccessKeyOutput) {
	op := &request.Operation{
		Name:       opUpdateAccessKey,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateAccessKeyInput{}
	}

	output = &UpdateAccessKeyOutput{}
	req = c.newRequest(op, input, output)

	return
}

// UpdateAccessKey API operation for IAM.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for IAM's
// API operation UpdateAccessKey for usage and error information.
func (c *IAM) UpdateAccessKey(input *UpdateAccessKeyInput) (*UpdateAccessKeyOutput, error) {
	req, out := c.UpdateAccessKeyRequest(input)
	return out, req.Send()
}

// UpdateAccessKeyWithContext is the same as UpdateAccessKey with the addition of
// the ability to pass a context and additional request options.
//
// See UpdateAccessKey for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *IAM) UpdateAccessKeyWithContext(ctx volcengine.Context, input *UpdateAccessKeyInput, opts ...request.Option) (*UpdateAccessKeyOutput, error) {
	req, out := c.UpdateAccessKeyRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type UpdateAccessKeyInput struct {
	_ struct{} `type:"structure"`

	// AccessKeyId is a required field
	AccessKeyId *string `type:"string" required:"true"`

	// Status is a required field
	Status *string `type:"string" required:"true"`
}

// String returns the string representation
func (s UpdateAccessKeyInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s UpdateAccessKeyInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateAccessKeyInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "UpdateAccessKeyInput"}
	if s.AccessKeyId == nil {
		invalidParams.Add(request.NewErrParamRequired("AccessKeyId"))
	}
	if s.Status == nil {
		invalidParams.Add(request.NewErrParamRequired("Status"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetAccessKeyId sets the AccessKeyId field's value.
func (s *UpdateAccessKeyInput) SetAccessKeyId(v string) *UpdateAccessKeyInput {
	s.AccessKeyId = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *UpdateAccessKeyInput) SetStatus(v string) *UpdateAccessKeyInput {
	s.Status = &v
	return s
}

type UpdateAccessKeyOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	ResponseMetadata *interface{} `type:"interface"`
}

// String returns the string representation
func (s UpdateAccessKeyOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s UpdateAccessKeyOutput) GoString() string {
	return s.String()
}

// SetResponseMetadata sets the ResponseMetadata field's value.
func (s *UpdateAccessKeyOutput) SetResponseMetadata(v interface{}) *UpdateAccessKeyOutput {
	s.ResponseMetadata = &v
	return s
}

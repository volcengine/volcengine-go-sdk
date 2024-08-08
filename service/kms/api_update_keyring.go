// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package kms

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opUpdateKeyringCommon = "UpdateKeyring"

// UpdateKeyringCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the UpdateKeyringCommon operation. The "output" return
// value will be populated with the UpdateKeyringCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned UpdateKeyringCommon Request to send the API call to the service.
// the "output" return value is not valid until after UpdateKeyringCommon Send returns without error.
//
// See UpdateKeyringCommon for more information on using the UpdateKeyringCommon
// API call, and error handling.
//
//    // Example sending a request using the UpdateKeyringCommonRequest method.
//    req, resp := client.UpdateKeyringCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *KMS) UpdateKeyringCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opUpdateKeyringCommon,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &map[string]interface{}{}
	}

	output = &map[string]interface{}{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// UpdateKeyringCommon API operation for KMS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for KMS's
// API operation UpdateKeyringCommon for usage and error information.
func (c *KMS) UpdateKeyringCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.UpdateKeyringCommonRequest(input)
	return out, req.Send()
}

// UpdateKeyringCommonWithContext is the same as UpdateKeyringCommon with the addition of
// the ability to pass a context and additional request options.
//
// See UpdateKeyringCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *KMS) UpdateKeyringCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.UpdateKeyringCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opUpdateKeyring = "UpdateKeyring"

// UpdateKeyringRequest generates a "volcengine/request.Request" representing the
// client's request for the UpdateKeyring operation. The "output" return
// value will be populated with the UpdateKeyringCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned UpdateKeyringCommon Request to send the API call to the service.
// the "output" return value is not valid until after UpdateKeyringCommon Send returns without error.
//
// See UpdateKeyring for more information on using the UpdateKeyring
// API call, and error handling.
//
//    // Example sending a request using the UpdateKeyringRequest method.
//    req, resp := client.UpdateKeyringRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *KMS) UpdateKeyringRequest(input *UpdateKeyringInput) (req *request.Request, output *UpdateKeyringOutput) {
	op := &request.Operation{
		Name:       opUpdateKeyring,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateKeyringInput{}
	}

	output = &UpdateKeyringOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// UpdateKeyring API operation for KMS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for KMS's
// API operation UpdateKeyring for usage and error information.
func (c *KMS) UpdateKeyring(input *UpdateKeyringInput) (*UpdateKeyringOutput, error) {
	req, out := c.UpdateKeyringRequest(input)
	return out, req.Send()
}

// UpdateKeyringWithContext is the same as UpdateKeyring with the addition of
// the ability to pass a context and additional request options.
//
// See UpdateKeyring for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *KMS) UpdateKeyringWithContext(ctx volcengine.Context, input *UpdateKeyringInput, opts ...request.Option) (*UpdateKeyringOutput, error) {
	req, out := c.UpdateKeyringRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type UpdateKeyringInput struct {
	_ struct{} `type:"structure"`

	Description *string `max:"8192" type:"string"`

	// KeyringName is a required field
	KeyringName *string `min:"2" max:"31" type:"string" required:"true"`
}

// String returns the string representation
func (s UpdateKeyringInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s UpdateKeyringInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateKeyringInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "UpdateKeyringInput"}
	if s.Description != nil && len(*s.Description) > 8192 {
		invalidParams.Add(request.NewErrParamMaxLen("Description", 8192, *s.Description))
	}
	if s.KeyringName == nil {
		invalidParams.Add(request.NewErrParamRequired("KeyringName"))
	}
	if s.KeyringName != nil && len(*s.KeyringName) < 2 {
		invalidParams.Add(request.NewErrParamMinLen("KeyringName", 2))
	}
	if s.KeyringName != nil && len(*s.KeyringName) > 31 {
		invalidParams.Add(request.NewErrParamMaxLen("KeyringName", 31, *s.KeyringName))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetDescription sets the Description field's value.
func (s *UpdateKeyringInput) SetDescription(v string) *UpdateKeyringInput {
	s.Description = &v
	return s
}

// SetKeyringName sets the KeyringName field's value.
func (s *UpdateKeyringInput) SetKeyringName(v string) *UpdateKeyringInput {
	s.KeyringName = &v
	return s
}

type UpdateKeyringOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s UpdateKeyringOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s UpdateKeyringOutput) GoString() string {
	return s.String()
}

// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package kms

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opEnableKeyRotationCommon = "EnableKeyRotation"

// EnableKeyRotationCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the EnableKeyRotationCommon operation. The "output" return
// value will be populated with the EnableKeyRotationCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned EnableKeyRotationCommon Request to send the API call to the service.
// the "output" return value is not valid until after EnableKeyRotationCommon Send returns without error.
//
// See EnableKeyRotationCommon for more information on using the EnableKeyRotationCommon
// API call, and error handling.
//
//    // Example sending a request using the EnableKeyRotationCommonRequest method.
//    req, resp := client.EnableKeyRotationCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *KMS) EnableKeyRotationCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opEnableKeyRotationCommon,
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

// EnableKeyRotationCommon API operation for KMS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for KMS's
// API operation EnableKeyRotationCommon for usage and error information.
func (c *KMS) EnableKeyRotationCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.EnableKeyRotationCommonRequest(input)
	return out, req.Send()
}

// EnableKeyRotationCommonWithContext is the same as EnableKeyRotationCommon with the addition of
// the ability to pass a context and additional request options.
//
// See EnableKeyRotationCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *KMS) EnableKeyRotationCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.EnableKeyRotationCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opEnableKeyRotation = "EnableKeyRotation"

// EnableKeyRotationRequest generates a "volcengine/request.Request" representing the
// client's request for the EnableKeyRotation operation. The "output" return
// value will be populated with the EnableKeyRotationCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned EnableKeyRotationCommon Request to send the API call to the service.
// the "output" return value is not valid until after EnableKeyRotationCommon Send returns without error.
//
// See EnableKeyRotation for more information on using the EnableKeyRotation
// API call, and error handling.
//
//    // Example sending a request using the EnableKeyRotationRequest method.
//    req, resp := client.EnableKeyRotationRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *KMS) EnableKeyRotationRequest(input *EnableKeyRotationInput) (req *request.Request, output *EnableKeyRotationOutput) {
	op := &request.Operation{
		Name:       opEnableKeyRotation,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &EnableKeyRotationInput{}
	}

	output = &EnableKeyRotationOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// EnableKeyRotation API operation for KMS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for KMS's
// API operation EnableKeyRotation for usage and error information.
func (c *KMS) EnableKeyRotation(input *EnableKeyRotationInput) (*EnableKeyRotationOutput, error) {
	req, out := c.EnableKeyRotationRequest(input)
	return out, req.Send()
}

// EnableKeyRotationWithContext is the same as EnableKeyRotation with the addition of
// the ability to pass a context and additional request options.
//
// See EnableKeyRotation for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *KMS) EnableKeyRotationWithContext(ctx volcengine.Context, input *EnableKeyRotationInput, opts ...request.Option) (*EnableKeyRotationOutput, error) {
	req, out := c.EnableKeyRotationRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type EnableKeyRotationInput struct {
	_ struct{} `type:"structure"`

	// KeyName is a required field
	KeyName *string `min:"2" max:"31" type:"string" required:"true"`

	// KeyringName is a required field
	KeyringName *string `min:"2" max:"31" type:"string" required:"true"`
}

// String returns the string representation
func (s EnableKeyRotationInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s EnableKeyRotationInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *EnableKeyRotationInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "EnableKeyRotationInput"}
	if s.KeyName == nil {
		invalidParams.Add(request.NewErrParamRequired("KeyName"))
	}
	if s.KeyName != nil && len(*s.KeyName) < 2 {
		invalidParams.Add(request.NewErrParamMinLen("KeyName", 2))
	}
	if s.KeyName != nil && len(*s.KeyName) > 31 {
		invalidParams.Add(request.NewErrParamMaxLen("KeyName", 31, *s.KeyName))
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

// SetKeyName sets the KeyName field's value.
func (s *EnableKeyRotationInput) SetKeyName(v string) *EnableKeyRotationInput {
	s.KeyName = &v
	return s
}

// SetKeyringName sets the KeyringName field's value.
func (s *EnableKeyRotationInput) SetKeyringName(v string) *EnableKeyRotationInput {
	s.KeyringName = &v
	return s
}

type EnableKeyRotationOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s EnableKeyRotationOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s EnableKeyRotationOutput) GoString() string {
	return s.String()
}

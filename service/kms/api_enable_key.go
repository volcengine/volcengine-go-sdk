// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package kms

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opEnableKeyCommon = "EnableKey"

// EnableKeyCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the EnableKeyCommon operation. The "output" return
// value will be populated with the EnableKeyCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned EnableKeyCommon Request to send the API call to the service.
// the "output" return value is not valid until after EnableKeyCommon Send returns without error.
//
// See EnableKeyCommon for more information on using the EnableKeyCommon
// API call, and error handling.
//
//    // Example sending a request using the EnableKeyCommonRequest method.
//    req, resp := client.EnableKeyCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *KMS) EnableKeyCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opEnableKeyCommon,
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

// EnableKeyCommon API operation for KMS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for KMS's
// API operation EnableKeyCommon for usage and error information.
func (c *KMS) EnableKeyCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.EnableKeyCommonRequest(input)
	return out, req.Send()
}

// EnableKeyCommonWithContext is the same as EnableKeyCommon with the addition of
// the ability to pass a context and additional request options.
//
// See EnableKeyCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *KMS) EnableKeyCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.EnableKeyCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opEnableKey = "EnableKey"

// EnableKeyRequest generates a "volcengine/request.Request" representing the
// client's request for the EnableKey operation. The "output" return
// value will be populated with the EnableKeyCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned EnableKeyCommon Request to send the API call to the service.
// the "output" return value is not valid until after EnableKeyCommon Send returns without error.
//
// See EnableKey for more information on using the EnableKey
// API call, and error handling.
//
//    // Example sending a request using the EnableKeyRequest method.
//    req, resp := client.EnableKeyRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *KMS) EnableKeyRequest(input *EnableKeyInput) (req *request.Request, output *EnableKeyOutput) {
	op := &request.Operation{
		Name:       opEnableKey,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &EnableKeyInput{}
	}

	output = &EnableKeyOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// EnableKey API operation for KMS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for KMS's
// API operation EnableKey for usage and error information.
func (c *KMS) EnableKey(input *EnableKeyInput) (*EnableKeyOutput, error) {
	req, out := c.EnableKeyRequest(input)
	return out, req.Send()
}

// EnableKeyWithContext is the same as EnableKey with the addition of
// the ability to pass a context and additional request options.
//
// See EnableKey for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *KMS) EnableKeyWithContext(ctx volcengine.Context, input *EnableKeyInput, opts ...request.Option) (*EnableKeyOutput, error) {
	req, out := c.EnableKeyRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type EnableKeyInput struct {
	_ struct{} `type:"structure"`

	// KeyName is a required field
	KeyName *string `min:"2" max:"31" type:"string" required:"true"`

	// KeyringName is a required field
	KeyringName *string `min:"2" max:"31" type:"string" required:"true"`
}

// String returns the string representation
func (s EnableKeyInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s EnableKeyInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *EnableKeyInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "EnableKeyInput"}
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
func (s *EnableKeyInput) SetKeyName(v string) *EnableKeyInput {
	s.KeyName = &v
	return s
}

// SetKeyringName sets the KeyringName field's value.
func (s *EnableKeyInput) SetKeyringName(v string) *EnableKeyInput {
	s.KeyringName = &v
	return s
}

type EnableKeyOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s EnableKeyOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s EnableKeyOutput) GoString() string {
	return s.String()
}

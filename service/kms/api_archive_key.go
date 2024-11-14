// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package kms

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opArchiveKeyCommon = "ArchiveKey"

// ArchiveKeyCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ArchiveKeyCommon operation. The "output" return
// value will be populated with the ArchiveKeyCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ArchiveKeyCommon Request to send the API call to the service.
// the "output" return value is not valid until after ArchiveKeyCommon Send returns without error.
//
// See ArchiveKeyCommon for more information on using the ArchiveKeyCommon
// API call, and error handling.
//
//    // Example sending a request using the ArchiveKeyCommonRequest method.
//    req, resp := client.ArchiveKeyCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *KMS) ArchiveKeyCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opArchiveKeyCommon,
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

// ArchiveKeyCommon API operation for KMS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for KMS's
// API operation ArchiveKeyCommon for usage and error information.
func (c *KMS) ArchiveKeyCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ArchiveKeyCommonRequest(input)
	return out, req.Send()
}

// ArchiveKeyCommonWithContext is the same as ArchiveKeyCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ArchiveKeyCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *KMS) ArchiveKeyCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ArchiveKeyCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opArchiveKey = "ArchiveKey"

// ArchiveKeyRequest generates a "volcengine/request.Request" representing the
// client's request for the ArchiveKey operation. The "output" return
// value will be populated with the ArchiveKeyCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ArchiveKeyCommon Request to send the API call to the service.
// the "output" return value is not valid until after ArchiveKeyCommon Send returns without error.
//
// See ArchiveKey for more information on using the ArchiveKey
// API call, and error handling.
//
//    // Example sending a request using the ArchiveKeyRequest method.
//    req, resp := client.ArchiveKeyRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *KMS) ArchiveKeyRequest(input *ArchiveKeyInput) (req *request.Request, output *ArchiveKeyOutput) {
	op := &request.Operation{
		Name:       opArchiveKey,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ArchiveKeyInput{}
	}

	output = &ArchiveKeyOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// ArchiveKey API operation for KMS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for KMS's
// API operation ArchiveKey for usage and error information.
func (c *KMS) ArchiveKey(input *ArchiveKeyInput) (*ArchiveKeyOutput, error) {
	req, out := c.ArchiveKeyRequest(input)
	return out, req.Send()
}

// ArchiveKeyWithContext is the same as ArchiveKey with the addition of
// the ability to pass a context and additional request options.
//
// See ArchiveKey for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *KMS) ArchiveKeyWithContext(ctx volcengine.Context, input *ArchiveKeyInput, opts ...request.Option) (*ArchiveKeyOutput, error) {
	req, out := c.ArchiveKeyRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ArchiveKeyInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	// KeyName is a required field
	KeyName *string `min:"2" max:"31" type:"string" json:",omitempty" required:"true"`

	// KeyringName is a required field
	KeyringName *string `min:"2" max:"31" type:"string" json:",omitempty" required:"true"`
}

// String returns the string representation
func (s ArchiveKeyInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ArchiveKeyInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ArchiveKeyInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ArchiveKeyInput"}
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
func (s *ArchiveKeyInput) SetKeyName(v string) *ArchiveKeyInput {
	s.KeyName = &v
	return s
}

// SetKeyringName sets the KeyringName field's value.
func (s *ArchiveKeyInput) SetKeyringName(v string) *ArchiveKeyInput {
	s.KeyringName = &v
	return s
}

type ArchiveKeyOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s ArchiveKeyOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ArchiveKeyOutput) GoString() string {
	return s.String()
}

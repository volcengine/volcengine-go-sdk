// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package kms

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opCancelArchiveKeyCommon = "CancelArchiveKey"

// CancelArchiveKeyCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the CancelArchiveKeyCommon operation. The "output" return
// value will be populated with the CancelArchiveKeyCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CancelArchiveKeyCommon Request to send the API call to the service.
// the "output" return value is not valid until after CancelArchiveKeyCommon Send returns without error.
//
// See CancelArchiveKeyCommon for more information on using the CancelArchiveKeyCommon
// API call, and error handling.
//
//    // Example sending a request using the CancelArchiveKeyCommonRequest method.
//    req, resp := client.CancelArchiveKeyCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *KMS) CancelArchiveKeyCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opCancelArchiveKeyCommon,
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

// CancelArchiveKeyCommon API operation for KMS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for KMS's
// API operation CancelArchiveKeyCommon for usage and error information.
func (c *KMS) CancelArchiveKeyCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.CancelArchiveKeyCommonRequest(input)
	return out, req.Send()
}

// CancelArchiveKeyCommonWithContext is the same as CancelArchiveKeyCommon with the addition of
// the ability to pass a context and additional request options.
//
// See CancelArchiveKeyCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *KMS) CancelArchiveKeyCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.CancelArchiveKeyCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opCancelArchiveKey = "CancelArchiveKey"

// CancelArchiveKeyRequest generates a "volcengine/request.Request" representing the
// client's request for the CancelArchiveKey operation. The "output" return
// value will be populated with the CancelArchiveKeyCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CancelArchiveKeyCommon Request to send the API call to the service.
// the "output" return value is not valid until after CancelArchiveKeyCommon Send returns without error.
//
// See CancelArchiveKey for more information on using the CancelArchiveKey
// API call, and error handling.
//
//    // Example sending a request using the CancelArchiveKeyRequest method.
//    req, resp := client.CancelArchiveKeyRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *KMS) CancelArchiveKeyRequest(input *CancelArchiveKeyInput) (req *request.Request, output *CancelArchiveKeyOutput) {
	op := &request.Operation{
		Name:       opCancelArchiveKey,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CancelArchiveKeyInput{}
	}

	output = &CancelArchiveKeyOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// CancelArchiveKey API operation for KMS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for KMS's
// API operation CancelArchiveKey for usage and error information.
func (c *KMS) CancelArchiveKey(input *CancelArchiveKeyInput) (*CancelArchiveKeyOutput, error) {
	req, out := c.CancelArchiveKeyRequest(input)
	return out, req.Send()
}

// CancelArchiveKeyWithContext is the same as CancelArchiveKey with the addition of
// the ability to pass a context and additional request options.
//
// See CancelArchiveKey for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *KMS) CancelArchiveKeyWithContext(ctx volcengine.Context, input *CancelArchiveKeyInput, opts ...request.Option) (*CancelArchiveKeyOutput, error) {
	req, out := c.CancelArchiveKeyRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type CancelArchiveKeyInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	KeyID *string `type:"string" json:",omitempty"`

	KeyName *string `min:"2" max:"31" type:"string" json:",omitempty"`

	KeyringName *string `min:"2" max:"31" type:"string" json:",omitempty"`
}

// String returns the string representation
func (s CancelArchiveKeyInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CancelArchiveKeyInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CancelArchiveKeyInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "CancelArchiveKeyInput"}
	if s.KeyName != nil && len(*s.KeyName) < 2 {
		invalidParams.Add(request.NewErrParamMinLen("KeyName", 2))
	}
	if s.KeyName != nil && len(*s.KeyName) > 31 {
		invalidParams.Add(request.NewErrParamMaxLen("KeyName", 31, *s.KeyName))
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

// SetKeyID sets the KeyID field's value.
func (s *CancelArchiveKeyInput) SetKeyID(v string) *CancelArchiveKeyInput {
	s.KeyID = &v
	return s
}

// SetKeyName sets the KeyName field's value.
func (s *CancelArchiveKeyInput) SetKeyName(v string) *CancelArchiveKeyInput {
	s.KeyName = &v
	return s
}

// SetKeyringName sets the KeyringName field's value.
func (s *CancelArchiveKeyInput) SetKeyringName(v string) *CancelArchiveKeyInput {
	s.KeyringName = &v
	return s
}

type CancelArchiveKeyOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s CancelArchiveKeyOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CancelArchiveKeyOutput) GoString() string {
	return s.String()
}

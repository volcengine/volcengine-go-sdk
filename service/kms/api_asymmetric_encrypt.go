// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package kms

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opAsymmetricEncryptCommon = "AsymmetricEncrypt"

// AsymmetricEncryptCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the AsymmetricEncryptCommon operation. The "output" return
// value will be populated with the AsymmetricEncryptCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned AsymmetricEncryptCommon Request to send the API call to the service.
// the "output" return value is not valid until after AsymmetricEncryptCommon Send returns without error.
//
// See AsymmetricEncryptCommon for more information on using the AsymmetricEncryptCommon
// API call, and error handling.
//
//    // Example sending a request using the AsymmetricEncryptCommonRequest method.
//    req, resp := client.AsymmetricEncryptCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *KMS) AsymmetricEncryptCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opAsymmetricEncryptCommon,
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

// AsymmetricEncryptCommon API operation for KMS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for KMS's
// API operation AsymmetricEncryptCommon for usage and error information.
func (c *KMS) AsymmetricEncryptCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.AsymmetricEncryptCommonRequest(input)
	return out, req.Send()
}

// AsymmetricEncryptCommonWithContext is the same as AsymmetricEncryptCommon with the addition of
// the ability to pass a context and additional request options.
//
// See AsymmetricEncryptCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *KMS) AsymmetricEncryptCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.AsymmetricEncryptCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opAsymmetricEncrypt = "AsymmetricEncrypt"

// AsymmetricEncryptRequest generates a "volcengine/request.Request" representing the
// client's request for the AsymmetricEncrypt operation. The "output" return
// value will be populated with the AsymmetricEncryptCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned AsymmetricEncryptCommon Request to send the API call to the service.
// the "output" return value is not valid until after AsymmetricEncryptCommon Send returns without error.
//
// See AsymmetricEncrypt for more information on using the AsymmetricEncrypt
// API call, and error handling.
//
//    // Example sending a request using the AsymmetricEncryptRequest method.
//    req, resp := client.AsymmetricEncryptRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *KMS) AsymmetricEncryptRequest(input *AsymmetricEncryptInput) (req *request.Request, output *AsymmetricEncryptOutput) {
	op := &request.Operation{
		Name:       opAsymmetricEncrypt,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &AsymmetricEncryptInput{}
	}

	output = &AsymmetricEncryptOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// AsymmetricEncrypt API operation for KMS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for KMS's
// API operation AsymmetricEncrypt for usage and error information.
func (c *KMS) AsymmetricEncrypt(input *AsymmetricEncryptInput) (*AsymmetricEncryptOutput, error) {
	req, out := c.AsymmetricEncryptRequest(input)
	return out, req.Send()
}

// AsymmetricEncryptWithContext is the same as AsymmetricEncrypt with the addition of
// the ability to pass a context and additional request options.
//
// See AsymmetricEncrypt for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *KMS) AsymmetricEncryptWithContext(ctx volcengine.Context, input *AsymmetricEncryptInput, opts ...request.Option) (*AsymmetricEncryptOutput, error) {
	req, out := c.AsymmetricEncryptRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type AsymmetricEncryptInput struct {
	_ struct{} `type:"structure"`

	// Algorithm is a required field
	Algorithm *string `type:"string" required:"true"`

	KeyID *string `type:"string"`

	KeyName *string `min:"2" max:"31" type:"string"`

	KeyringName *string `min:"2" max:"31" type:"string"`

	// Plaintext is a required field
	Plaintext *string `min:"1" max:"4096" type:"string" required:"true"`
}

// String returns the string representation
func (s AsymmetricEncryptInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s AsymmetricEncryptInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AsymmetricEncryptInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "AsymmetricEncryptInput"}
	if s.Algorithm == nil {
		invalidParams.Add(request.NewErrParamRequired("Algorithm"))
	}
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
	if s.Plaintext == nil {
		invalidParams.Add(request.NewErrParamRequired("Plaintext"))
	}
	if s.Plaintext != nil && len(*s.Plaintext) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("Plaintext", 1))
	}
	if s.Plaintext != nil && len(*s.Plaintext) > 4096 {
		invalidParams.Add(request.NewErrParamMaxLen("Plaintext", 4096, *s.Plaintext))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetAlgorithm sets the Algorithm field's value.
func (s *AsymmetricEncryptInput) SetAlgorithm(v string) *AsymmetricEncryptInput {
	s.Algorithm = &v
	return s
}

// SetKeyID sets the KeyID field's value.
func (s *AsymmetricEncryptInput) SetKeyID(v string) *AsymmetricEncryptInput {
	s.KeyID = &v
	return s
}

// SetKeyName sets the KeyName field's value.
func (s *AsymmetricEncryptInput) SetKeyName(v string) *AsymmetricEncryptInput {
	s.KeyName = &v
	return s
}

// SetKeyringName sets the KeyringName field's value.
func (s *AsymmetricEncryptInput) SetKeyringName(v string) *AsymmetricEncryptInput {
	s.KeyringName = &v
	return s
}

// SetPlaintext sets the Plaintext field's value.
func (s *AsymmetricEncryptInput) SetPlaintext(v string) *AsymmetricEncryptInput {
	s.Plaintext = &v
	return s
}

type AsymmetricEncryptOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	CiphertextBlob *string `type:"string"`

	KeyID *string `type:"string"`
}

// String returns the string representation
func (s AsymmetricEncryptOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s AsymmetricEncryptOutput) GoString() string {
	return s.String()
}

// SetCiphertextBlob sets the CiphertextBlob field's value.
func (s *AsymmetricEncryptOutput) SetCiphertextBlob(v string) *AsymmetricEncryptOutput {
	s.CiphertextBlob = &v
	return s
}

// SetKeyID sets the KeyID field's value.
func (s *AsymmetricEncryptOutput) SetKeyID(v string) *AsymmetricEncryptOutput {
	s.KeyID = &v
	return s
}

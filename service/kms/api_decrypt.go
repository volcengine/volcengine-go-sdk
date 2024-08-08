// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package kms

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDecryptCommon = "Decrypt"

// DecryptCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DecryptCommon operation. The "output" return
// value will be populated with the DecryptCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DecryptCommon Request to send the API call to the service.
// the "output" return value is not valid until after DecryptCommon Send returns without error.
//
// See DecryptCommon for more information on using the DecryptCommon
// API call, and error handling.
//
//    // Example sending a request using the DecryptCommonRequest method.
//    req, resp := client.DecryptCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *KMS) DecryptCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDecryptCommon,
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

// DecryptCommon API operation for KMS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for KMS's
// API operation DecryptCommon for usage and error information.
func (c *KMS) DecryptCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DecryptCommonRequest(input)
	return out, req.Send()
}

// DecryptCommonWithContext is the same as DecryptCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DecryptCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *KMS) DecryptCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DecryptCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDecrypt = "Decrypt"

// DecryptRequest generates a "volcengine/request.Request" representing the
// client's request for the Decrypt operation. The "output" return
// value will be populated with the DecryptCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DecryptCommon Request to send the API call to the service.
// the "output" return value is not valid until after DecryptCommon Send returns without error.
//
// See Decrypt for more information on using the Decrypt
// API call, and error handling.
//
//    // Example sending a request using the DecryptRequest method.
//    req, resp := client.DecryptRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *KMS) DecryptRequest(input *DecryptInput) (req *request.Request, output *DecryptOutput) {
	op := &request.Operation{
		Name:       opDecrypt,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DecryptInput{}
	}

	output = &DecryptOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// Decrypt API operation for KMS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for KMS's
// API operation Decrypt for usage and error information.
func (c *KMS) Decrypt(input *DecryptInput) (*DecryptOutput, error) {
	req, out := c.DecryptRequest(input)
	return out, req.Send()
}

// DecryptWithContext is the same as Decrypt with the addition of
// the ability to pass a context and additional request options.
//
// See Decrypt for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *KMS) DecryptWithContext(ctx volcengine.Context, input *DecryptInput, opts ...request.Option) (*DecryptOutput, error) {
	req, out := c.DecryptRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DecryptInput struct {
	_ struct{} `type:"structure"`

	// CiphertextBlob is a required field
	CiphertextBlob *string `min:"19" type:"string" required:"true"`

	EncryptionContext *string `type:"string"`
}

// String returns the string representation
func (s DecryptInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DecryptInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DecryptInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DecryptInput"}
	if s.CiphertextBlob == nil {
		invalidParams.Add(request.NewErrParamRequired("CiphertextBlob"))
	}
	if s.CiphertextBlob != nil && len(*s.CiphertextBlob) < 19 {
		invalidParams.Add(request.NewErrParamMinLen("CiphertextBlob", 19))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetCiphertextBlob sets the CiphertextBlob field's value.
func (s *DecryptInput) SetCiphertextBlob(v string) *DecryptInput {
	s.CiphertextBlob = &v
	return s
}

// SetEncryptionContext sets the EncryptionContext field's value.
func (s *DecryptInput) SetEncryptionContext(v string) *DecryptInput {
	s.EncryptionContext = &v
	return s
}

type DecryptOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	Plaintext *string `type:"string"`
}

// String returns the string representation
func (s DecryptOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DecryptOutput) GoString() string {
	return s.String()
}

// SetPlaintext sets the Plaintext field's value.
func (s *DecryptOutput) SetPlaintext(v string) *DecryptOutput {
	s.Plaintext = &v
	return s
}

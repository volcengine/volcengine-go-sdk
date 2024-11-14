// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package kms

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opGetParametersForImportCommon = "GetParametersForImport"

// GetParametersForImportCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the GetParametersForImportCommon operation. The "output" return
// value will be populated with the GetParametersForImportCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned GetParametersForImportCommon Request to send the API call to the service.
// the "output" return value is not valid until after GetParametersForImportCommon Send returns without error.
//
// See GetParametersForImportCommon for more information on using the GetParametersForImportCommon
// API call, and error handling.
//
//    // Example sending a request using the GetParametersForImportCommonRequest method.
//    req, resp := client.GetParametersForImportCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *KMS) GetParametersForImportCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opGetParametersForImportCommon,
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

// GetParametersForImportCommon API operation for KMS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for KMS's
// API operation GetParametersForImportCommon for usage and error information.
func (c *KMS) GetParametersForImportCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.GetParametersForImportCommonRequest(input)
	return out, req.Send()
}

// GetParametersForImportCommonWithContext is the same as GetParametersForImportCommon with the addition of
// the ability to pass a context and additional request options.
//
// See GetParametersForImportCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *KMS) GetParametersForImportCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.GetParametersForImportCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opGetParametersForImport = "GetParametersForImport"

// GetParametersForImportRequest generates a "volcengine/request.Request" representing the
// client's request for the GetParametersForImport operation. The "output" return
// value will be populated with the GetParametersForImportCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned GetParametersForImportCommon Request to send the API call to the service.
// the "output" return value is not valid until after GetParametersForImportCommon Send returns without error.
//
// See GetParametersForImport for more information on using the GetParametersForImport
// API call, and error handling.
//
//    // Example sending a request using the GetParametersForImportRequest method.
//    req, resp := client.GetParametersForImportRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *KMS) GetParametersForImportRequest(input *GetParametersForImportInput) (req *request.Request, output *GetParametersForImportOutput) {
	op := &request.Operation{
		Name:       opGetParametersForImport,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetParametersForImportInput{}
	}

	output = &GetParametersForImportOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// GetParametersForImport API operation for KMS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for KMS's
// API operation GetParametersForImport for usage and error information.
func (c *KMS) GetParametersForImport(input *GetParametersForImportInput) (*GetParametersForImportOutput, error) {
	req, out := c.GetParametersForImportRequest(input)
	return out, req.Send()
}

// GetParametersForImportWithContext is the same as GetParametersForImport with the addition of
// the ability to pass a context and additional request options.
//
// See GetParametersForImport for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *KMS) GetParametersForImportWithContext(ctx volcengine.Context, input *GetParametersForImportInput, opts ...request.Option) (*GetParametersForImportOutput, error) {
	req, out := c.GetParametersForImportRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type GetParametersForImportInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	// KeyName is a required field
	KeyName *string `min:"2" max:"31" type:"string" json:",omitempty" required:"true"`

	// KeyringName is a required field
	KeyringName *string `min:"2" max:"31" type:"string" json:",omitempty" required:"true"`

	WrappingAlgorithm *string `type:"string" json:",omitempty"`

	WrappingKeySpec *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s GetParametersForImportInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s GetParametersForImportInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetParametersForImportInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "GetParametersForImportInput"}
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
func (s *GetParametersForImportInput) SetKeyName(v string) *GetParametersForImportInput {
	s.KeyName = &v
	return s
}

// SetKeyringName sets the KeyringName field's value.
func (s *GetParametersForImportInput) SetKeyringName(v string) *GetParametersForImportInput {
	s.KeyringName = &v
	return s
}

// SetWrappingAlgorithm sets the WrappingAlgorithm field's value.
func (s *GetParametersForImportInput) SetWrappingAlgorithm(v string) *GetParametersForImportInput {
	s.WrappingAlgorithm = &v
	return s
}

// SetWrappingKeySpec sets the WrappingKeySpec field's value.
func (s *GetParametersForImportInput) SetWrappingKeySpec(v string) *GetParametersForImportInput {
	s.WrappingKeySpec = &v
	return s
}

type GetParametersForImportOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	ImportToken *string `type:"string" json:",omitempty"`

	KeyID *string `type:"string" json:",omitempty"`

	KeyringID *string `type:"string" json:",omitempty"`

	PublicKey *string `type:"string" json:",omitempty"`

	TokenExpireTime *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s GetParametersForImportOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s GetParametersForImportOutput) GoString() string {
	return s.String()
}

// SetImportToken sets the ImportToken field's value.
func (s *GetParametersForImportOutput) SetImportToken(v string) *GetParametersForImportOutput {
	s.ImportToken = &v
	return s
}

// SetKeyID sets the KeyID field's value.
func (s *GetParametersForImportOutput) SetKeyID(v string) *GetParametersForImportOutput {
	s.KeyID = &v
	return s
}

// SetKeyringID sets the KeyringID field's value.
func (s *GetParametersForImportOutput) SetKeyringID(v string) *GetParametersForImportOutput {
	s.KeyringID = &v
	return s
}

// SetPublicKey sets the PublicKey field's value.
func (s *GetParametersForImportOutput) SetPublicKey(v string) *GetParametersForImportOutput {
	s.PublicKey = &v
	return s
}

// SetTokenExpireTime sets the TokenExpireTime field's value.
func (s *GetParametersForImportOutput) SetTokenExpireTime(v string) *GetParametersForImportOutput {
	s.TokenExpireTime = &v
	return s
}

// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package iam

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opAddSAMLProviderCertificateCommon = "AddSAMLProviderCertificate"

// AddSAMLProviderCertificateCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the AddSAMLProviderCertificateCommon operation. The "output" return
// value will be populated with the AddSAMLProviderCertificateCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned AddSAMLProviderCertificateCommon Request to send the API call to the service.
// the "output" return value is not valid until after AddSAMLProviderCertificateCommon Send returns without error.
//
// See AddSAMLProviderCertificateCommon for more information on using the AddSAMLProviderCertificateCommon
// API call, and error handling.
//
//	// Example sending a request using the AddSAMLProviderCertificateCommonRequest method.
//	req, resp := client.AddSAMLProviderCertificateCommonRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *IAM) AddSAMLProviderCertificateCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opAddSAMLProviderCertificateCommon,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &map[string]interface{}{}
	}

	output = &map[string]interface{}{}
	req = c.newRequest(op, input, output)

	return
}

// AddSAMLProviderCertificateCommon API operation for IAM.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for IAM's
// API operation AddSAMLProviderCertificateCommon for usage and error information.
func (c *IAM) AddSAMLProviderCertificateCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.AddSAMLProviderCertificateCommonRequest(input)
	return out, req.Send()
}

// AddSAMLProviderCertificateCommonWithContext is the same as AddSAMLProviderCertificateCommon with the addition of
// the ability to pass a context and additional request options.
//
// See AddSAMLProviderCertificateCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *IAM) AddSAMLProviderCertificateCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.AddSAMLProviderCertificateCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opAddSAMLProviderCertificate = "AddSAMLProviderCertificate"

// AddSAMLProviderCertificateRequest generates a "volcengine/request.Request" representing the
// client's request for the AddSAMLProviderCertificate operation. The "output" return
// value will be populated with the AddSAMLProviderCertificateCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned AddSAMLProviderCertificateCommon Request to send the API call to the service.
// the "output" return value is not valid until after AddSAMLProviderCertificateCommon Send returns without error.
//
// See AddSAMLProviderCertificate for more information on using the AddSAMLProviderCertificate
// API call, and error handling.
//
//	// Example sending a request using the AddSAMLProviderCertificateRequest method.
//	req, resp := client.AddSAMLProviderCertificateRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *IAM) AddSAMLProviderCertificateRequest(input *AddSAMLProviderCertificateInput) (req *request.Request, output *AddSAMLProviderCertificateOutput) {
	op := &request.Operation{
		Name:       opAddSAMLProviderCertificate,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &AddSAMLProviderCertificateInput{}
	}

	output = &AddSAMLProviderCertificateOutput{}
	req = c.newRequest(op, input, output)

	return
}

// AddSAMLProviderCertificate API operation for IAM.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for IAM's
// API operation AddSAMLProviderCertificate for usage and error information.
func (c *IAM) AddSAMLProviderCertificate(input *AddSAMLProviderCertificateInput) (*AddSAMLProviderCertificateOutput, error) {
	req, out := c.AddSAMLProviderCertificateRequest(input)
	return out, req.Send()
}

// AddSAMLProviderCertificateWithContext is the same as AddSAMLProviderCertificate with the addition of
// the ability to pass a context and additional request options.
//
// See AddSAMLProviderCertificate for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *IAM) AddSAMLProviderCertificateWithContext(ctx volcengine.Context, input *AddSAMLProviderCertificateInput, opts ...request.Option) (*AddSAMLProviderCertificateOutput, error) {
	req, out := c.AddSAMLProviderCertificateRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type AddSAMLProviderCertificateInput struct {
	_ struct{} `type:"structure"`

	// SAMLProviderName is a required field
	SAMLProviderName *string `type:"string" required:"true"`

	// X509Certificate is a required field
	X509Certificate *string `type:"string" required:"true"`
}

// String returns the string representation
func (s AddSAMLProviderCertificateInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s AddSAMLProviderCertificateInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AddSAMLProviderCertificateInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "AddSAMLProviderCertificateInput"}
	if s.SAMLProviderName == nil {
		invalidParams.Add(request.NewErrParamRequired("SAMLProviderName"))
	}
	if s.X509Certificate == nil {
		invalidParams.Add(request.NewErrParamRequired("X509Certificate"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetSAMLProviderName sets the SAMLProviderName field's value.
func (s *AddSAMLProviderCertificateInput) SetSAMLProviderName(v string) *AddSAMLProviderCertificateInput {
	s.SAMLProviderName = &v
	return s
}

// SetX509Certificate sets the X509Certificate field's value.
func (s *AddSAMLProviderCertificateInput) SetX509Certificate(v string) *AddSAMLProviderCertificateInput {
	s.X509Certificate = &v
	return s
}

type AddSAMLProviderCertificateOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	CertificateId *string `type:"string"`
}

// String returns the string representation
func (s AddSAMLProviderCertificateOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s AddSAMLProviderCertificateOutput) GoString() string {
	return s.String()
}

// SetCertificateId sets the CertificateId field's value.
func (s *AddSAMLProviderCertificateOutput) SetCertificateId(v string) *AddSAMLProviderCertificateOutput {
	s.CertificateId = &v
	return s
}

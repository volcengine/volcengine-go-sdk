// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package iam

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opListSAMLProviderCertificatesCommon = "ListSAMLProviderCertificates"

// ListSAMLProviderCertificatesCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ListSAMLProviderCertificatesCommon operation. The "output" return
// value will be populated with the ListSAMLProviderCertificatesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListSAMLProviderCertificatesCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListSAMLProviderCertificatesCommon Send returns without error.
//
// See ListSAMLProviderCertificatesCommon for more information on using the ListSAMLProviderCertificatesCommon
// API call, and error handling.
//
//    // Example sending a request using the ListSAMLProviderCertificatesCommonRequest method.
//    req, resp := client.ListSAMLProviderCertificatesCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *IAM) ListSAMLProviderCertificatesCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opListSAMLProviderCertificatesCommon,
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

// ListSAMLProviderCertificatesCommon API operation for IAM.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for IAM's
// API operation ListSAMLProviderCertificatesCommon for usage and error information.
func (c *IAM) ListSAMLProviderCertificatesCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ListSAMLProviderCertificatesCommonRequest(input)
	return out, req.Send()
}

// ListSAMLProviderCertificatesCommonWithContext is the same as ListSAMLProviderCertificatesCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ListSAMLProviderCertificatesCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *IAM) ListSAMLProviderCertificatesCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ListSAMLProviderCertificatesCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opListSAMLProviderCertificates = "ListSAMLProviderCertificates"

// ListSAMLProviderCertificatesRequest generates a "volcengine/request.Request" representing the
// client's request for the ListSAMLProviderCertificates operation. The "output" return
// value will be populated with the ListSAMLProviderCertificatesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListSAMLProviderCertificatesCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListSAMLProviderCertificatesCommon Send returns without error.
//
// See ListSAMLProviderCertificates for more information on using the ListSAMLProviderCertificates
// API call, and error handling.
//
//    // Example sending a request using the ListSAMLProviderCertificatesRequest method.
//    req, resp := client.ListSAMLProviderCertificatesRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *IAM) ListSAMLProviderCertificatesRequest(input *ListSAMLProviderCertificatesInput) (req *request.Request, output *ListSAMLProviderCertificatesOutput) {
	op := &request.Operation{
		Name:       opListSAMLProviderCertificates,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListSAMLProviderCertificatesInput{}
	}

	output = &ListSAMLProviderCertificatesOutput{}
	req = c.newRequest(op, input, output)

	return
}

// ListSAMLProviderCertificates API operation for IAM.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for IAM's
// API operation ListSAMLProviderCertificates for usage and error information.
func (c *IAM) ListSAMLProviderCertificates(input *ListSAMLProviderCertificatesInput) (*ListSAMLProviderCertificatesOutput, error) {
	req, out := c.ListSAMLProviderCertificatesRequest(input)
	return out, req.Send()
}

// ListSAMLProviderCertificatesWithContext is the same as ListSAMLProviderCertificates with the addition of
// the ability to pass a context and additional request options.
//
// See ListSAMLProviderCertificates for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *IAM) ListSAMLProviderCertificatesWithContext(ctx volcengine.Context, input *ListSAMLProviderCertificatesInput, opts ...request.Option) (*ListSAMLProviderCertificatesOutput, error) {
	req, out := c.ListSAMLProviderCertificatesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type CertificateForListSAMLProviderCertificatesOutput struct {
	_ struct{} `type:"structure"`

	CertificateId *string `type:"string"`

	CreateDate *string `type:"string"`

	Issuer *string `type:"string"`

	NotAfter *string `type:"string"`

	NotBefore *string `type:"string"`

	SerialNumber *string `type:"string"`

	SignatureAlgorithm *string `type:"string"`

	Subject *string `type:"string"`

	UpdateDate *string `type:"string"`

	Version *string `type:"string"`
}

// String returns the string representation
func (s CertificateForListSAMLProviderCertificatesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CertificateForListSAMLProviderCertificatesOutput) GoString() string {
	return s.String()
}

// SetCertificateId sets the CertificateId field's value.
func (s *CertificateForListSAMLProviderCertificatesOutput) SetCertificateId(v string) *CertificateForListSAMLProviderCertificatesOutput {
	s.CertificateId = &v
	return s
}

// SetCreateDate sets the CreateDate field's value.
func (s *CertificateForListSAMLProviderCertificatesOutput) SetCreateDate(v string) *CertificateForListSAMLProviderCertificatesOutput {
	s.CreateDate = &v
	return s
}

// SetIssuer sets the Issuer field's value.
func (s *CertificateForListSAMLProviderCertificatesOutput) SetIssuer(v string) *CertificateForListSAMLProviderCertificatesOutput {
	s.Issuer = &v
	return s
}

// SetNotAfter sets the NotAfter field's value.
func (s *CertificateForListSAMLProviderCertificatesOutput) SetNotAfter(v string) *CertificateForListSAMLProviderCertificatesOutput {
	s.NotAfter = &v
	return s
}

// SetNotBefore sets the NotBefore field's value.
func (s *CertificateForListSAMLProviderCertificatesOutput) SetNotBefore(v string) *CertificateForListSAMLProviderCertificatesOutput {
	s.NotBefore = &v
	return s
}

// SetSerialNumber sets the SerialNumber field's value.
func (s *CertificateForListSAMLProviderCertificatesOutput) SetSerialNumber(v string) *CertificateForListSAMLProviderCertificatesOutput {
	s.SerialNumber = &v
	return s
}

// SetSignatureAlgorithm sets the SignatureAlgorithm field's value.
func (s *CertificateForListSAMLProviderCertificatesOutput) SetSignatureAlgorithm(v string) *CertificateForListSAMLProviderCertificatesOutput {
	s.SignatureAlgorithm = &v
	return s
}

// SetSubject sets the Subject field's value.
func (s *CertificateForListSAMLProviderCertificatesOutput) SetSubject(v string) *CertificateForListSAMLProviderCertificatesOutput {
	s.Subject = &v
	return s
}

// SetUpdateDate sets the UpdateDate field's value.
func (s *CertificateForListSAMLProviderCertificatesOutput) SetUpdateDate(v string) *CertificateForListSAMLProviderCertificatesOutput {
	s.UpdateDate = &v
	return s
}

// SetVersion sets the Version field's value.
func (s *CertificateForListSAMLProviderCertificatesOutput) SetVersion(v string) *CertificateForListSAMLProviderCertificatesOutput {
	s.Version = &v
	return s
}

type ListSAMLProviderCertificatesInput struct {
	_ struct{} `type:"structure"`

	// SAMLProviderName is a required field
	SAMLProviderName *string `type:"string" required:"true"`
}

// String returns the string representation
func (s ListSAMLProviderCertificatesInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListSAMLProviderCertificatesInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListSAMLProviderCertificatesInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ListSAMLProviderCertificatesInput"}
	if s.SAMLProviderName == nil {
		invalidParams.Add(request.NewErrParamRequired("SAMLProviderName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetSAMLProviderName sets the SAMLProviderName field's value.
func (s *ListSAMLProviderCertificatesInput) SetSAMLProviderName(v string) *ListSAMLProviderCertificatesInput {
	s.SAMLProviderName = &v
	return s
}

type ListSAMLProviderCertificatesOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	Certificates []*CertificateForListSAMLProviderCertificatesOutput `type:"list"`

	Total *int32 `type:"int32"`
}

// String returns the string representation
func (s ListSAMLProviderCertificatesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListSAMLProviderCertificatesOutput) GoString() string {
	return s.String()
}

// SetCertificates sets the Certificates field's value.
func (s *ListSAMLProviderCertificatesOutput) SetCertificates(v []*CertificateForListSAMLProviderCertificatesOutput) *ListSAMLProviderCertificatesOutput {
	s.Certificates = v
	return s
}

// SetTotal sets the Total field's value.
func (s *ListSAMLProviderCertificatesOutput) SetTotal(v int32) *ListSAMLProviderCertificatesOutput {
	s.Total = &v
	return s
}

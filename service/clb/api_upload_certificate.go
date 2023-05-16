// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package clb

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opUploadCertificateCommon = "UploadCertificate"

// UploadCertificateCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the UploadCertificateCommon operation. The "output" return
// value will be populated with the UploadCertificateCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned UploadCertificateCommon Request to send the API call to the service.
// the "output" return value is not valid until after UploadCertificateCommon Send returns without error.
//
// See UploadCertificateCommon for more information on using the UploadCertificateCommon
// API call, and error handling.
//
//	// Example sending a request using the UploadCertificateCommonRequest method.
//	req, resp := client.UploadCertificateCommonRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *CLB) UploadCertificateCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opUploadCertificateCommon,
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

// UploadCertificateCommon API operation for CLB.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CLB's
// API operation UploadCertificateCommon for usage and error information.
func (c *CLB) UploadCertificateCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.UploadCertificateCommonRequest(input)
	return out, req.Send()
}

// UploadCertificateCommonWithContext is the same as UploadCertificateCommon with the addition of
// the ability to pass a context and additional request options.
//
// See UploadCertificateCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CLB) UploadCertificateCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.UploadCertificateCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opUploadCertificate = "UploadCertificate"

// UploadCertificateRequest generates a "volcengine/request.Request" representing the
// client's request for the UploadCertificate operation. The "output" return
// value will be populated with the UploadCertificateCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned UploadCertificateCommon Request to send the API call to the service.
// the "output" return value is not valid until after UploadCertificateCommon Send returns without error.
//
// See UploadCertificate for more information on using the UploadCertificate
// API call, and error handling.
//
//	// Example sending a request using the UploadCertificateRequest method.
//	req, resp := client.UploadCertificateRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *CLB) UploadCertificateRequest(input *UploadCertificateInput) (req *request.Request, output *UploadCertificateOutput) {
	op := &request.Operation{
		Name:       opUploadCertificate,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UploadCertificateInput{}
	}

	output = &UploadCertificateOutput{}
	req = c.newRequest(op, input, output)

	return
}

// UploadCertificate API operation for CLB.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CLB's
// API operation UploadCertificate for usage and error information.
func (c *CLB) UploadCertificate(input *UploadCertificateInput) (*UploadCertificateOutput, error) {
	req, out := c.UploadCertificateRequest(input)
	return out, req.Send()
}

// UploadCertificateWithContext is the same as UploadCertificate with the addition of
// the ability to pass a context and additional request options.
//
// See UploadCertificate for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CLB) UploadCertificateWithContext(ctx volcengine.Context, input *UploadCertificateInput, opts ...request.Option) (*UploadCertificateOutput, error) {
	req, out := c.UploadCertificateRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type UploadCertificateInput struct {
	_ struct{} `type:"structure"`

	CertificateName *string `type:"string"`

	Description *string `type:"string"`

	// PrivateKey is a required field
	PrivateKey *string `type:"string" required:"true"`

	ProjectName *string `type:"string"`

	// PublicKey is a required field
	PublicKey *string `type:"string" required:"true"`

	TagFilters *string `type:"string"`
}

// String returns the string representation
func (s UploadCertificateInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s UploadCertificateInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UploadCertificateInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "UploadCertificateInput"}
	if s.PrivateKey == nil {
		invalidParams.Add(request.NewErrParamRequired("PrivateKey"))
	}
	if s.PublicKey == nil {
		invalidParams.Add(request.NewErrParamRequired("PublicKey"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetCertificateName sets the CertificateName field's value.
func (s *UploadCertificateInput) SetCertificateName(v string) *UploadCertificateInput {
	s.CertificateName = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *UploadCertificateInput) SetDescription(v string) *UploadCertificateInput {
	s.Description = &v
	return s
}

// SetPrivateKey sets the PrivateKey field's value.
func (s *UploadCertificateInput) SetPrivateKey(v string) *UploadCertificateInput {
	s.PrivateKey = &v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *UploadCertificateInput) SetProjectName(v string) *UploadCertificateInput {
	s.ProjectName = &v
	return s
}

// SetPublicKey sets the PublicKey field's value.
func (s *UploadCertificateInput) SetPublicKey(v string) *UploadCertificateInput {
	s.PublicKey = &v
	return s
}

// SetTagFilters sets the TagFilters field's value.
func (s *UploadCertificateInput) SetTagFilters(v string) *UploadCertificateInput {
	s.TagFilters = &v
	return s
}

type UploadCertificateOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	CertificateId *string `type:"string"`

	RequestId *string `type:"string"`
}

// String returns the string representation
func (s UploadCertificateOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s UploadCertificateOutput) GoString() string {
	return s.String()
}

// SetCertificateId sets the CertificateId field's value.
func (s *UploadCertificateOutput) SetCertificateId(v string) *UploadCertificateOutput {
	s.CertificateId = &v
	return s
}

// SetRequestId sets the RequestId field's value.
func (s *UploadCertificateOutput) SetRequestId(v string) *UploadCertificateOutput {
	s.RequestId = &v
	return s
}

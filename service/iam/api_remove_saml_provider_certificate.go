// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package iam

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opRemoveSAMLProviderCertificateCommon = "RemoveSAMLProviderCertificate"

// RemoveSAMLProviderCertificateCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the RemoveSAMLProviderCertificateCommon operation. The "output" return
// value will be populated with the RemoveSAMLProviderCertificateCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned RemoveSAMLProviderCertificateCommon Request to send the API call to the service.
// the "output" return value is not valid until after RemoveSAMLProviderCertificateCommon Send returns without error.
//
// See RemoveSAMLProviderCertificateCommon for more information on using the RemoveSAMLProviderCertificateCommon
// API call, and error handling.
//
//    // Example sending a request using the RemoveSAMLProviderCertificateCommonRequest method.
//    req, resp := client.RemoveSAMLProviderCertificateCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *IAM) RemoveSAMLProviderCertificateCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opRemoveSAMLProviderCertificateCommon,
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

// RemoveSAMLProviderCertificateCommon API operation for IAM.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for IAM's
// API operation RemoveSAMLProviderCertificateCommon for usage and error information.
func (c *IAM) RemoveSAMLProviderCertificateCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.RemoveSAMLProviderCertificateCommonRequest(input)
	return out, req.Send()
}

// RemoveSAMLProviderCertificateCommonWithContext is the same as RemoveSAMLProviderCertificateCommon with the addition of
// the ability to pass a context and additional request options.
//
// See RemoveSAMLProviderCertificateCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *IAM) RemoveSAMLProviderCertificateCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.RemoveSAMLProviderCertificateCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opRemoveSAMLProviderCertificate = "RemoveSAMLProviderCertificate"

// RemoveSAMLProviderCertificateRequest generates a "volcengine/request.Request" representing the
// client's request for the RemoveSAMLProviderCertificate operation. The "output" return
// value will be populated with the RemoveSAMLProviderCertificateCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned RemoveSAMLProviderCertificateCommon Request to send the API call to the service.
// the "output" return value is not valid until after RemoveSAMLProviderCertificateCommon Send returns without error.
//
// See RemoveSAMLProviderCertificate for more information on using the RemoveSAMLProviderCertificate
// API call, and error handling.
//
//    // Example sending a request using the RemoveSAMLProviderCertificateRequest method.
//    req, resp := client.RemoveSAMLProviderCertificateRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *IAM) RemoveSAMLProviderCertificateRequest(input *RemoveSAMLProviderCertificateInput) (req *request.Request, output *RemoveSAMLProviderCertificateOutput) {
	op := &request.Operation{
		Name:       opRemoveSAMLProviderCertificate,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &RemoveSAMLProviderCertificateInput{}
	}

	output = &RemoveSAMLProviderCertificateOutput{}
	req = c.newRequest(op, input, output)

	return
}

// RemoveSAMLProviderCertificate API operation for IAM.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for IAM's
// API operation RemoveSAMLProviderCertificate for usage and error information.
func (c *IAM) RemoveSAMLProviderCertificate(input *RemoveSAMLProviderCertificateInput) (*RemoveSAMLProviderCertificateOutput, error) {
	req, out := c.RemoveSAMLProviderCertificateRequest(input)
	return out, req.Send()
}

// RemoveSAMLProviderCertificateWithContext is the same as RemoveSAMLProviderCertificate with the addition of
// the ability to pass a context and additional request options.
//
// See RemoveSAMLProviderCertificate for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *IAM) RemoveSAMLProviderCertificateWithContext(ctx volcengine.Context, input *RemoveSAMLProviderCertificateInput, opts ...request.Option) (*RemoveSAMLProviderCertificateOutput, error) {
	req, out := c.RemoveSAMLProviderCertificateRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type RemoveSAMLProviderCertificateInput struct {
	_ struct{} `type:"structure"`

	// CertificateId is a required field
	CertificateId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s RemoveSAMLProviderCertificateInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s RemoveSAMLProviderCertificateInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RemoveSAMLProviderCertificateInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "RemoveSAMLProviderCertificateInput"}
	if s.CertificateId == nil {
		invalidParams.Add(request.NewErrParamRequired("CertificateId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetCertificateId sets the CertificateId field's value.
func (s *RemoveSAMLProviderCertificateInput) SetCertificateId(v string) *RemoveSAMLProviderCertificateInput {
	s.CertificateId = &v
	return s
}

type RemoveSAMLProviderCertificateOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s RemoveSAMLProviderCertificateOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s RemoveSAMLProviderCertificateOutput) GoString() string {
	return s.String()
}

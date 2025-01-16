// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package alb

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDeleteCertificateCommon = "DeleteCertificate"

// DeleteCertificateCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DeleteCertificateCommon operation. The "output" return
// value will be populated with the DeleteCertificateCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeleteCertificateCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeleteCertificateCommon Send returns without error.
//
// See DeleteCertificateCommon for more information on using the DeleteCertificateCommon
// API call, and error handling.
//
//    // Example sending a request using the DeleteCertificateCommonRequest method.
//    req, resp := client.DeleteCertificateCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ALB) DeleteCertificateCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDeleteCertificateCommon,
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

// DeleteCertificateCommon API operation for ALB.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ALB's
// API operation DeleteCertificateCommon for usage and error information.
func (c *ALB) DeleteCertificateCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DeleteCertificateCommonRequest(input)
	return out, req.Send()
}

// DeleteCertificateCommonWithContext is the same as DeleteCertificateCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteCertificateCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ALB) DeleteCertificateCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DeleteCertificateCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDeleteCertificate = "DeleteCertificate"

// DeleteCertificateRequest generates a "volcengine/request.Request" representing the
// client's request for the DeleteCertificate operation. The "output" return
// value will be populated with the DeleteCertificateCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeleteCertificateCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeleteCertificateCommon Send returns without error.
//
// See DeleteCertificate for more information on using the DeleteCertificate
// API call, and error handling.
//
//    // Example sending a request using the DeleteCertificateRequest method.
//    req, resp := client.DeleteCertificateRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ALB) DeleteCertificateRequest(input *DeleteCertificateInput) (req *request.Request, output *DeleteCertificateOutput) {
	op := &request.Operation{
		Name:       opDeleteCertificate,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteCertificateInput{}
	}

	output = &DeleteCertificateOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DeleteCertificate API operation for ALB.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ALB's
// API operation DeleteCertificate for usage and error information.
func (c *ALB) DeleteCertificate(input *DeleteCertificateInput) (*DeleteCertificateOutput, error) {
	req, out := c.DeleteCertificateRequest(input)
	return out, req.Send()
}

// DeleteCertificateWithContext is the same as DeleteCertificate with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteCertificate for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ALB) DeleteCertificateWithContext(ctx volcengine.Context, input *DeleteCertificateInput, opts ...request.Option) (*DeleteCertificateOutput, error) {
	req, out := c.DeleteCertificateRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DeleteCertificateInput struct {
	_ struct{} `type:"structure"`

	CertificateId *string `type:"string"`
}

// String returns the string representation
func (s DeleteCertificateInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteCertificateInput) GoString() string {
	return s.String()
}

// SetCertificateId sets the CertificateId field's value.
func (s *DeleteCertificateInput) SetCertificateId(v string) *DeleteCertificateInput {
	s.CertificateId = &v
	return s
}

type DeleteCertificateOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	RequestId *string `type:"string"`
}

// String returns the string representation
func (s DeleteCertificateOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteCertificateOutput) GoString() string {
	return s.String()
}

// SetRequestId sets the RequestId field's value.
func (s *DeleteCertificateOutput) SetRequestId(v string) *DeleteCertificateOutput {
	s.RequestId = &v
	return s
}

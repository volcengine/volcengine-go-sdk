// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package certificateservice

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opCertificateDeleteInstanceCommon = "CertificateDeleteInstance"

// CertificateDeleteInstanceCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the CertificateDeleteInstanceCommon operation. The "output" return
// value will be populated with the CertificateDeleteInstanceCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CertificateDeleteInstanceCommon Request to send the API call to the service.
// the "output" return value is not valid until after CertificateDeleteInstanceCommon Send returns without error.
//
// See CertificateDeleteInstanceCommon for more information on using the CertificateDeleteInstanceCommon
// API call, and error handling.
//
//    // Example sending a request using the CertificateDeleteInstanceCommonRequest method.
//    req, resp := client.CertificateDeleteInstanceCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CERTIFICATESERVICE) CertificateDeleteInstanceCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opCertificateDeleteInstanceCommon,
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

// CertificateDeleteInstanceCommon API operation for CERTIFICATE_SERVICE.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CERTIFICATE_SERVICE's
// API operation CertificateDeleteInstanceCommon for usage and error information.
func (c *CERTIFICATESERVICE) CertificateDeleteInstanceCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.CertificateDeleteInstanceCommonRequest(input)
	return out, req.Send()
}

// CertificateDeleteInstanceCommonWithContext is the same as CertificateDeleteInstanceCommon with the addition of
// the ability to pass a context and additional request options.
//
// See CertificateDeleteInstanceCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CERTIFICATESERVICE) CertificateDeleteInstanceCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.CertificateDeleteInstanceCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opCertificateDeleteInstance = "CertificateDeleteInstance"

// CertificateDeleteInstanceRequest generates a "volcengine/request.Request" representing the
// client's request for the CertificateDeleteInstance operation. The "output" return
// value will be populated with the CertificateDeleteInstanceCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CertificateDeleteInstanceCommon Request to send the API call to the service.
// the "output" return value is not valid until after CertificateDeleteInstanceCommon Send returns without error.
//
// See CertificateDeleteInstance for more information on using the CertificateDeleteInstance
// API call, and error handling.
//
//    // Example sending a request using the CertificateDeleteInstanceRequest method.
//    req, resp := client.CertificateDeleteInstanceRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CERTIFICATESERVICE) CertificateDeleteInstanceRequest(input *CertificateDeleteInstanceInput) (req *request.Request, output *CertificateDeleteInstanceOutput) {
	op := &request.Operation{
		Name:       opCertificateDeleteInstance,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CertificateDeleteInstanceInput{}
	}

	output = &CertificateDeleteInstanceOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// CertificateDeleteInstance API operation for CERTIFICATE_SERVICE.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CERTIFICATE_SERVICE's
// API operation CertificateDeleteInstance for usage and error information.
func (c *CERTIFICATESERVICE) CertificateDeleteInstance(input *CertificateDeleteInstanceInput) (*CertificateDeleteInstanceOutput, error) {
	req, out := c.CertificateDeleteInstanceRequest(input)
	return out, req.Send()
}

// CertificateDeleteInstanceWithContext is the same as CertificateDeleteInstance with the addition of
// the ability to pass a context and additional request options.
//
// See CertificateDeleteInstance for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CERTIFICATESERVICE) CertificateDeleteInstanceWithContext(ctx volcengine.Context, input *CertificateDeleteInstanceInput, opts ...request.Option) (*CertificateDeleteInstanceOutput, error) {
	req, out := c.CertificateDeleteInstanceRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type CertificateDeleteInstanceInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	// InstanceId is a required field
	InstanceId *string `type:"string" json:",omitempty" required:"true"`
}

// String returns the string representation
func (s CertificateDeleteInstanceInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CertificateDeleteInstanceInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CertificateDeleteInstanceInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "CertificateDeleteInstanceInput"}
	if s.InstanceId == nil {
		invalidParams.Add(request.NewErrParamRequired("InstanceId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetInstanceId sets the InstanceId field's value.
func (s *CertificateDeleteInstanceInput) SetInstanceId(v string) *CertificateDeleteInstanceInput {
	s.InstanceId = &v
	return s
}

type CertificateDeleteInstanceOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s CertificateDeleteInstanceOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CertificateDeleteInstanceOutput) GoString() string {
	return s.String()
}

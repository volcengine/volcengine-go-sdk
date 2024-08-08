// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package waf

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opQueryCertificateIfReplaceCommon = "QueryCertificateIfReplace"

// QueryCertificateIfReplaceCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the QueryCertificateIfReplaceCommon operation. The "output" return
// value will be populated with the QueryCertificateIfReplaceCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned QueryCertificateIfReplaceCommon Request to send the API call to the service.
// the "output" return value is not valid until after QueryCertificateIfReplaceCommon Send returns without error.
//
// See QueryCertificateIfReplaceCommon for more information on using the QueryCertificateIfReplaceCommon
// API call, and error handling.
//
//    // Example sending a request using the QueryCertificateIfReplaceCommonRequest method.
//    req, resp := client.QueryCertificateIfReplaceCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *WAF) QueryCertificateIfReplaceCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opQueryCertificateIfReplaceCommon,
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

// QueryCertificateIfReplaceCommon API operation for WAF.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for WAF's
// API operation QueryCertificateIfReplaceCommon for usage and error information.
func (c *WAF) QueryCertificateIfReplaceCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.QueryCertificateIfReplaceCommonRequest(input)
	return out, req.Send()
}

// QueryCertificateIfReplaceCommonWithContext is the same as QueryCertificateIfReplaceCommon with the addition of
// the ability to pass a context and additional request options.
//
// See QueryCertificateIfReplaceCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *WAF) QueryCertificateIfReplaceCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.QueryCertificateIfReplaceCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opQueryCertificateIfReplace = "QueryCertificateIfReplace"

// QueryCertificateIfReplaceRequest generates a "volcengine/request.Request" representing the
// client's request for the QueryCertificateIfReplace operation. The "output" return
// value will be populated with the QueryCertificateIfReplaceCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned QueryCertificateIfReplaceCommon Request to send the API call to the service.
// the "output" return value is not valid until after QueryCertificateIfReplaceCommon Send returns without error.
//
// See QueryCertificateIfReplace for more information on using the QueryCertificateIfReplace
// API call, and error handling.
//
//    // Example sending a request using the QueryCertificateIfReplaceRequest method.
//    req, resp := client.QueryCertificateIfReplaceRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *WAF) QueryCertificateIfReplaceRequest(input *QueryCertificateIfReplaceInput) (req *request.Request, output *QueryCertificateIfReplaceOutput) {
	op := &request.Operation{
		Name:       opQueryCertificateIfReplace,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &QueryCertificateIfReplaceInput{}
	}

	output = &QueryCertificateIfReplaceOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// QueryCertificateIfReplace API operation for WAF.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for WAF's
// API operation QueryCertificateIfReplace for usage and error information.
func (c *WAF) QueryCertificateIfReplace(input *QueryCertificateIfReplaceInput) (*QueryCertificateIfReplaceOutput, error) {
	req, out := c.QueryCertificateIfReplaceRequest(input)
	return out, req.Send()
}

// QueryCertificateIfReplaceWithContext is the same as QueryCertificateIfReplace with the addition of
// the ability to pass a context and additional request options.
//
// See QueryCertificateIfReplace for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *WAF) QueryCertificateIfReplaceWithContext(ctx volcengine.Context, input *QueryCertificateIfReplaceInput, opts ...request.Option) (*QueryCertificateIfReplaceOutput, error) {
	req, out := c.QueryCertificateIfReplaceRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type QueryCertificateIfReplaceInput struct {
	_ struct{} `type:"structure"`

	// CertificateID is a required field
	CertificateID *int32 `type:"int32" required:"true"`

	// Domain is a required field
	Domain *string `type:"string" required:"true"`

	PublicRealServer *int32 `type:"int32"`
}

// String returns the string representation
func (s QueryCertificateIfReplaceInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s QueryCertificateIfReplaceInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *QueryCertificateIfReplaceInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "QueryCertificateIfReplaceInput"}
	if s.CertificateID == nil {
		invalidParams.Add(request.NewErrParamRequired("CertificateID"))
	}
	if s.Domain == nil {
		invalidParams.Add(request.NewErrParamRequired("Domain"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetCertificateID sets the CertificateID field's value.
func (s *QueryCertificateIfReplaceInput) SetCertificateID(v int32) *QueryCertificateIfReplaceInput {
	s.CertificateID = &v
	return s
}

// SetDomain sets the Domain field's value.
func (s *QueryCertificateIfReplaceInput) SetDomain(v string) *QueryCertificateIfReplaceInput {
	s.Domain = &v
	return s
}

// SetPublicRealServer sets the PublicRealServer field's value.
func (s *QueryCertificateIfReplaceInput) SetPublicRealServer(v int32) *QueryCertificateIfReplaceInput {
	s.PublicRealServer = &v
	return s
}

type QueryCertificateIfReplaceOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	IfReplace *bool `type:"boolean"`
}

// String returns the string representation
func (s QueryCertificateIfReplaceOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s QueryCertificateIfReplaceOutput) GoString() string {
	return s.String()
}

// SetIfReplace sets the IfReplace field's value.
func (s *QueryCertificateIfReplaceOutput) SetIfReplace(v bool) *QueryCertificateIfReplaceOutput {
	s.IfReplace = &v
	return s
}
// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package rdsmysql

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opCopyParameterTemplateCommon = "CopyParameterTemplate"

// CopyParameterTemplateCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the CopyParameterTemplateCommon operation. The "output" return
// value will be populated with the CopyParameterTemplateCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CopyParameterTemplateCommon Request to send the API call to the service.
// the "output" return value is not valid until after CopyParameterTemplateCommon Send returns without error.
//
// See CopyParameterTemplateCommon for more information on using the CopyParameterTemplateCommon
// API call, and error handling.
//
//    // Example sending a request using the CopyParameterTemplateCommonRequest method.
//    req, resp := client.CopyParameterTemplateCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *RDSMYSQL) CopyParameterTemplateCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opCopyParameterTemplateCommon,
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

// CopyParameterTemplateCommon API operation for RDS_MYSQL.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RDS_MYSQL's
// API operation CopyParameterTemplateCommon for usage and error information.
func (c *RDSMYSQL) CopyParameterTemplateCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.CopyParameterTemplateCommonRequest(input)
	return out, req.Send()
}

// CopyParameterTemplateCommonWithContext is the same as CopyParameterTemplateCommon with the addition of
// the ability to pass a context and additional request options.
//
// See CopyParameterTemplateCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RDSMYSQL) CopyParameterTemplateCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.CopyParameterTemplateCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opCopyParameterTemplate = "CopyParameterTemplate"

// CopyParameterTemplateRequest generates a "volcengine/request.Request" representing the
// client's request for the CopyParameterTemplate operation. The "output" return
// value will be populated with the CopyParameterTemplateCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CopyParameterTemplateCommon Request to send the API call to the service.
// the "output" return value is not valid until after CopyParameterTemplateCommon Send returns without error.
//
// See CopyParameterTemplate for more information on using the CopyParameterTemplate
// API call, and error handling.
//
//    // Example sending a request using the CopyParameterTemplateRequest method.
//    req, resp := client.CopyParameterTemplateRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *RDSMYSQL) CopyParameterTemplateRequest(input *CopyParameterTemplateInput) (req *request.Request, output *CopyParameterTemplateOutput) {
	op := &request.Operation{
		Name:       opCopyParameterTemplate,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CopyParameterTemplateInput{}
	}

	output = &CopyParameterTemplateOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// CopyParameterTemplate API operation for RDS_MYSQL.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RDS_MYSQL's
// API operation CopyParameterTemplate for usage and error information.
func (c *RDSMYSQL) CopyParameterTemplate(input *CopyParameterTemplateInput) (*CopyParameterTemplateOutput, error) {
	req, out := c.CopyParameterTemplateRequest(input)
	return out, req.Send()
}

// CopyParameterTemplateWithContext is the same as CopyParameterTemplate with the addition of
// the ability to pass a context and additional request options.
//
// See CopyParameterTemplate for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RDSMYSQL) CopyParameterTemplateWithContext(ctx volcengine.Context, input *CopyParameterTemplateInput, opts ...request.Option) (*CopyParameterTemplateOutput, error) {
	req, out := c.CopyParameterTemplateRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type CopyParameterTemplateInput struct {
	_ struct{} `type:"structure"`

	// SrcTemplateId is a required field
	SrcTemplateId *interface{} `type:"interface" required:"true"`

	TemplateDesc *string `max:"200" type:"string"`

	TemplateName *string `min:"2" max:"64" type:"string"`
}

// String returns the string representation
func (s CopyParameterTemplateInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CopyParameterTemplateInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CopyParameterTemplateInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "CopyParameterTemplateInput"}
	if s.SrcTemplateId == nil {
		invalidParams.Add(request.NewErrParamRequired("SrcTemplateId"))
	}
	if s.TemplateDesc != nil && len(*s.TemplateDesc) > 200 {
		invalidParams.Add(request.NewErrParamMaxLen("TemplateDesc", 200, *s.TemplateDesc))
	}
	if s.TemplateName != nil && len(*s.TemplateName) < 2 {
		invalidParams.Add(request.NewErrParamMinLen("TemplateName", 2))
	}
	if s.TemplateName != nil && len(*s.TemplateName) > 64 {
		invalidParams.Add(request.NewErrParamMaxLen("TemplateName", 64, *s.TemplateName))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetSrcTemplateId sets the SrcTemplateId field's value.
func (s *CopyParameterTemplateInput) SetSrcTemplateId(v interface{}) *CopyParameterTemplateInput {
	s.SrcTemplateId = &v
	return s
}

// SetTemplateDesc sets the TemplateDesc field's value.
func (s *CopyParameterTemplateInput) SetTemplateDesc(v string) *CopyParameterTemplateInput {
	s.TemplateDesc = &v
	return s
}

// SetTemplateName sets the TemplateName field's value.
func (s *CopyParameterTemplateInput) SetTemplateName(v string) *CopyParameterTemplateInput {
	s.TemplateName = &v
	return s
}

type CopyParameterTemplateOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s CopyParameterTemplateOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CopyParameterTemplateOutput) GoString() string {
	return s.String()
}

// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package rdsmysqlv2

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDeleteParameterTemplateCommon = "DeleteParameterTemplate"

// DeleteParameterTemplateCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DeleteParameterTemplateCommon operation. The "output" return
// value will be populated with the DeleteParameterTemplateCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeleteParameterTemplateCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeleteParameterTemplateCommon Send returns without error.
//
// See DeleteParameterTemplateCommon for more information on using the DeleteParameterTemplateCommon
// API call, and error handling.
//
//    // Example sending a request using the DeleteParameterTemplateCommonRequest method.
//    req, resp := client.DeleteParameterTemplateCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *RDSMYSQLV2) DeleteParameterTemplateCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDeleteParameterTemplateCommon,
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

// DeleteParameterTemplateCommon API operation for RDS_MYSQL_V2.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RDS_MYSQL_V2's
// API operation DeleteParameterTemplateCommon for usage and error information.
func (c *RDSMYSQLV2) DeleteParameterTemplateCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DeleteParameterTemplateCommonRequest(input)
	return out, req.Send()
}

// DeleteParameterTemplateCommonWithContext is the same as DeleteParameterTemplateCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteParameterTemplateCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RDSMYSQLV2) DeleteParameterTemplateCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DeleteParameterTemplateCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDeleteParameterTemplate = "DeleteParameterTemplate"

// DeleteParameterTemplateRequest generates a "volcengine/request.Request" representing the
// client's request for the DeleteParameterTemplate operation. The "output" return
// value will be populated with the DeleteParameterTemplateCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeleteParameterTemplateCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeleteParameterTemplateCommon Send returns without error.
//
// See DeleteParameterTemplate for more information on using the DeleteParameterTemplate
// API call, and error handling.
//
//    // Example sending a request using the DeleteParameterTemplateRequest method.
//    req, resp := client.DeleteParameterTemplateRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *RDSMYSQLV2) DeleteParameterTemplateRequest(input *DeleteParameterTemplateInput) (req *request.Request, output *DeleteParameterTemplateOutput) {
	op := &request.Operation{
		Name:       opDeleteParameterTemplate,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteParameterTemplateInput{}
	}

	output = &DeleteParameterTemplateOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// DeleteParameterTemplate API operation for RDS_MYSQL_V2.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RDS_MYSQL_V2's
// API operation DeleteParameterTemplate for usage and error information.
func (c *RDSMYSQLV2) DeleteParameterTemplate(input *DeleteParameterTemplateInput) (*DeleteParameterTemplateOutput, error) {
	req, out := c.DeleteParameterTemplateRequest(input)
	return out, req.Send()
}

// DeleteParameterTemplateWithContext is the same as DeleteParameterTemplate with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteParameterTemplate for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RDSMYSQLV2) DeleteParameterTemplateWithContext(ctx volcengine.Context, input *DeleteParameterTemplateInput, opts ...request.Option) (*DeleteParameterTemplateOutput, error) {
	req, out := c.DeleteParameterTemplateRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DeleteParameterTemplateInput struct {
	_ struct{} `type:"structure"`

	// TemplateId is a required field
	TemplateId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteParameterTemplateInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteParameterTemplateInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteParameterTemplateInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DeleteParameterTemplateInput"}
	if s.TemplateId == nil {
		invalidParams.Add(request.NewErrParamRequired("TemplateId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetTemplateId sets the TemplateId field's value.
func (s *DeleteParameterTemplateInput) SetTemplateId(v string) *DeleteParameterTemplateInput {
	s.TemplateId = &v
	return s
}

type DeleteParameterTemplateOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s DeleteParameterTemplateOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteParameterTemplateOutput) GoString() string {
	return s.String()
}

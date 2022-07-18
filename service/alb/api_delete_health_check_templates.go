// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package alb

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDeleteHealthCheckTemplatesCommon = "DeleteHealthCheckTemplates"

// DeleteHealthCheckTemplatesCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DeleteHealthCheckTemplatesCommon operation. The "output" return
// value will be populated with the DeleteHealthCheckTemplatesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeleteHealthCheckTemplatesCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeleteHealthCheckTemplatesCommon Send returns without error.
//
// See DeleteHealthCheckTemplatesCommon for more information on using the DeleteHealthCheckTemplatesCommon
// API call, and error handling.
//
//    // Example sending a request using the DeleteHealthCheckTemplatesCommonRequest method.
//    req, resp := client.DeleteHealthCheckTemplatesCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ALB) DeleteHealthCheckTemplatesCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDeleteHealthCheckTemplatesCommon,
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

// DeleteHealthCheckTemplatesCommon API operation for ALB.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ALB's
// API operation DeleteHealthCheckTemplatesCommon for usage and error information.
func (c *ALB) DeleteHealthCheckTemplatesCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DeleteHealthCheckTemplatesCommonRequest(input)
	return out, req.Send()
}

// DeleteHealthCheckTemplatesCommonWithContext is the same as DeleteHealthCheckTemplatesCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteHealthCheckTemplatesCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ALB) DeleteHealthCheckTemplatesCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DeleteHealthCheckTemplatesCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDeleteHealthCheckTemplates = "DeleteHealthCheckTemplates"

// DeleteHealthCheckTemplatesRequest generates a "volcengine/request.Request" representing the
// client's request for the DeleteHealthCheckTemplates operation. The "output" return
// value will be populated with the DeleteHealthCheckTemplatesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeleteHealthCheckTemplatesCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeleteHealthCheckTemplatesCommon Send returns without error.
//
// See DeleteHealthCheckTemplates for more information on using the DeleteHealthCheckTemplates
// API call, and error handling.
//
//    // Example sending a request using the DeleteHealthCheckTemplatesRequest method.
//    req, resp := client.DeleteHealthCheckTemplatesRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ALB) DeleteHealthCheckTemplatesRequest(input *DeleteHealthCheckTemplatesInput) (req *request.Request, output *DeleteHealthCheckTemplatesOutput) {
	op := &request.Operation{
		Name:       opDeleteHealthCheckTemplates,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteHealthCheckTemplatesInput{}
	}

	output = &DeleteHealthCheckTemplatesOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DeleteHealthCheckTemplates API operation for ALB.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ALB's
// API operation DeleteHealthCheckTemplates for usage and error information.
func (c *ALB) DeleteHealthCheckTemplates(input *DeleteHealthCheckTemplatesInput) (*DeleteHealthCheckTemplatesOutput, error) {
	req, out := c.DeleteHealthCheckTemplatesRequest(input)
	return out, req.Send()
}

// DeleteHealthCheckTemplatesWithContext is the same as DeleteHealthCheckTemplates with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteHealthCheckTemplates for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ALB) DeleteHealthCheckTemplatesWithContext(ctx volcengine.Context, input *DeleteHealthCheckTemplatesInput, opts ...request.Option) (*DeleteHealthCheckTemplatesOutput, error) {
	req, out := c.DeleteHealthCheckTemplatesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DeleteHealthCheckTemplatesInput struct {
	_ struct{} `type:"structure"`

	// HealthCheckTemplateIds is a required field
	HealthCheckTemplateIds []*string `type:"list" required:"true"`
}

// String returns the string representation
func (s DeleteHealthCheckTemplatesInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteHealthCheckTemplatesInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteHealthCheckTemplatesInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DeleteHealthCheckTemplatesInput"}
	if s.HealthCheckTemplateIds == nil {
		invalidParams.Add(request.NewErrParamRequired("HealthCheckTemplateIds"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetHealthCheckTemplateIds sets the HealthCheckTemplateIds field's value.
func (s *DeleteHealthCheckTemplatesInput) SetHealthCheckTemplateIds(v []*string) *DeleteHealthCheckTemplatesInput {
	s.HealthCheckTemplateIds = v
	return s
}

type DeleteHealthCheckTemplatesOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	RequestId *string `type:"string"`
}

// String returns the string representation
func (s DeleteHealthCheckTemplatesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteHealthCheckTemplatesOutput) GoString() string {
	return s.String()
}

// SetRequestId sets the RequestId field's value.
func (s *DeleteHealthCheckTemplatesOutput) SetRequestId(v string) *DeleteHealthCheckTemplatesOutput {
	s.RequestId = &v
	return s
}
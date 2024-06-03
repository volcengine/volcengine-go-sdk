// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package waf

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opQueryLLMGenerateCommon = "QueryLLMGenerate"

// QueryLLMGenerateCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the QueryLLMGenerateCommon operation. The "output" return
// value will be populated with the QueryLLMGenerateCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned QueryLLMGenerateCommon Request to send the API call to the service.
// the "output" return value is not valid until after QueryLLMGenerateCommon Send returns without error.
//
// See QueryLLMGenerateCommon for more information on using the QueryLLMGenerateCommon
// API call, and error handling.
//
//    // Example sending a request using the QueryLLMGenerateCommonRequest method.
//    req, resp := client.QueryLLMGenerateCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *WAF) QueryLLMGenerateCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opQueryLLMGenerateCommon,
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

// QueryLLMGenerateCommon API operation for WAF.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for WAF's
// API operation QueryLLMGenerateCommon for usage and error information.
func (c *WAF) QueryLLMGenerateCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.QueryLLMGenerateCommonRequest(input)
	return out, req.Send()
}

// QueryLLMGenerateCommonWithContext is the same as QueryLLMGenerateCommon with the addition of
// the ability to pass a context and additional request options.
//
// See QueryLLMGenerateCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *WAF) QueryLLMGenerateCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.QueryLLMGenerateCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opQueryLLMGenerate = "QueryLLMGenerate"

// QueryLLMGenerateRequest generates a "volcengine/request.Request" representing the
// client's request for the QueryLLMGenerate operation. The "output" return
// value will be populated with the QueryLLMGenerateCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned QueryLLMGenerateCommon Request to send the API call to the service.
// the "output" return value is not valid until after QueryLLMGenerateCommon Send returns without error.
//
// See QueryLLMGenerate for more information on using the QueryLLMGenerate
// API call, and error handling.
//
//    // Example sending a request using the QueryLLMGenerateRequest method.
//    req, resp := client.QueryLLMGenerateRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *WAF) QueryLLMGenerateRequest(input *QueryLLMGenerateInput) (req *request.Request, output *QueryLLMGenerateOutput) {
	op := &request.Operation{
		Name:       opQueryLLMGenerate,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &QueryLLMGenerateInput{}
	}

	output = &QueryLLMGenerateOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// QueryLLMGenerate API operation for WAF.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for WAF's
// API operation QueryLLMGenerate for usage and error information.
func (c *WAF) QueryLLMGenerate(input *QueryLLMGenerateInput) (*QueryLLMGenerateOutput, error) {
	req, out := c.QueryLLMGenerateRequest(input)
	return out, req.Send()
}

// QueryLLMGenerateWithContext is the same as QueryLLMGenerate with the addition of
// the ability to pass a context and additional request options.
//
// See QueryLLMGenerate for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *WAF) QueryLLMGenerateWithContext(ctx volcengine.Context, input *QueryLLMGenerateInput, opts ...request.Option) (*QueryLLMGenerateOutput, error) {
	req, out := c.QueryLLMGenerateRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type QueryLLMGenerateInput struct {
	_ struct{} `type:"structure"`

	// MsgID is a required field
	MsgID *string `type:"string" required:"true"`

	UseStream *bool `type:"boolean"`
}

// String returns the string representation
func (s QueryLLMGenerateInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s QueryLLMGenerateInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *QueryLLMGenerateInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "QueryLLMGenerateInput"}
	if s.MsgID == nil {
		invalidParams.Add(request.NewErrParamRequired("MsgID"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetMsgID sets the MsgID field's value.
func (s *QueryLLMGenerateInput) SetMsgID(v string) *QueryLLMGenerateInput {
	s.MsgID = &v
	return s
}

// SetUseStream sets the UseStream field's value.
func (s *QueryLLMGenerateInput) SetUseStream(v bool) *QueryLLMGenerateInput {
	s.UseStream = &v
	return s
}

type QueryLLMGenerateOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	Suggest *string `type:"string"`
}

// String returns the string representation
func (s QueryLLMGenerateOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s QueryLLMGenerateOutput) GoString() string {
	return s.String()
}

// SetSuggest sets the Suggest field's value.
func (s *QueryLLMGenerateOutput) SetSuggest(v string) *QueryLLMGenerateOutput {
	s.Suggest = &v
	return s
}

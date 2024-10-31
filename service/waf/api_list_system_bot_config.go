// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package waf

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opListSystemBotConfigCommon = "ListSystemBotConfig"

// ListSystemBotConfigCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ListSystemBotConfigCommon operation. The "output" return
// value will be populated with the ListSystemBotConfigCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListSystemBotConfigCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListSystemBotConfigCommon Send returns without error.
//
// See ListSystemBotConfigCommon for more information on using the ListSystemBotConfigCommon
// API call, and error handling.
//
//    // Example sending a request using the ListSystemBotConfigCommonRequest method.
//    req, resp := client.ListSystemBotConfigCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *WAF) ListSystemBotConfigCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opListSystemBotConfigCommon,
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

// ListSystemBotConfigCommon API operation for WAF.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for WAF's
// API operation ListSystemBotConfigCommon for usage and error information.
func (c *WAF) ListSystemBotConfigCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ListSystemBotConfigCommonRequest(input)
	return out, req.Send()
}

// ListSystemBotConfigCommonWithContext is the same as ListSystemBotConfigCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ListSystemBotConfigCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *WAF) ListSystemBotConfigCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ListSystemBotConfigCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opListSystemBotConfig = "ListSystemBotConfig"

// ListSystemBotConfigRequest generates a "volcengine/request.Request" representing the
// client's request for the ListSystemBotConfig operation. The "output" return
// value will be populated with the ListSystemBotConfigCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListSystemBotConfigCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListSystemBotConfigCommon Send returns without error.
//
// See ListSystemBotConfig for more information on using the ListSystemBotConfig
// API call, and error handling.
//
//    // Example sending a request using the ListSystemBotConfigRequest method.
//    req, resp := client.ListSystemBotConfigRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *WAF) ListSystemBotConfigRequest(input *ListSystemBotConfigInput) (req *request.Request, output *ListSystemBotConfigOutput) {
	op := &request.Operation{
		Name:       opListSystemBotConfig,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListSystemBotConfigInput{}
	}

	output = &ListSystemBotConfigOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// ListSystemBotConfig API operation for WAF.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for WAF's
// API operation ListSystemBotConfig for usage and error information.
func (c *WAF) ListSystemBotConfig(input *ListSystemBotConfigInput) (*ListSystemBotConfigOutput, error) {
	req, out := c.ListSystemBotConfigRequest(input)
	return out, req.Send()
}

// ListSystemBotConfigWithContext is the same as ListSystemBotConfig with the addition of
// the ability to pass a context and additional request options.
//
// See ListSystemBotConfig for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *WAF) ListSystemBotConfigWithContext(ctx volcengine.Context, input *ListSystemBotConfigInput, opts ...request.Option) (*ListSystemBotConfigOutput, error) {
	req, out := c.ListSystemBotConfigRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DataForListSystemBotConfigOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Action *string `type:"string" json:",omitempty"`

	BotType *string `type:"string" json:",omitempty"`

	Description *string `type:"string" json:",omitempty"`

	Enable *int32 `type:"int32" json:",omitempty"`

	RuleTag *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s DataForListSystemBotConfigOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DataForListSystemBotConfigOutput) GoString() string {
	return s.String()
}

// SetAction sets the Action field's value.
func (s *DataForListSystemBotConfigOutput) SetAction(v string) *DataForListSystemBotConfigOutput {
	s.Action = &v
	return s
}

// SetBotType sets the BotType field's value.
func (s *DataForListSystemBotConfigOutput) SetBotType(v string) *DataForListSystemBotConfigOutput {
	s.BotType = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *DataForListSystemBotConfigOutput) SetDescription(v string) *DataForListSystemBotConfigOutput {
	s.Description = &v
	return s
}

// SetEnable sets the Enable field's value.
func (s *DataForListSystemBotConfigOutput) SetEnable(v int32) *DataForListSystemBotConfigOutput {
	s.Enable = &v
	return s
}

// SetRuleTag sets the RuleTag field's value.
func (s *DataForListSystemBotConfigOutput) SetRuleTag(v string) *DataForListSystemBotConfigOutput {
	s.RuleTag = &v
	return s
}

type ListSystemBotConfigInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	// Host is a required field
	Host *string `type:"string" json:",omitempty" required:"true"`

	ProjectName *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s ListSystemBotConfigInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListSystemBotConfigInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListSystemBotConfigInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ListSystemBotConfigInput"}
	if s.Host == nil {
		invalidParams.Add(request.NewErrParamRequired("Host"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetHost sets the Host field's value.
func (s *ListSystemBotConfigInput) SetHost(v string) *ListSystemBotConfigInput {
	s.Host = &v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *ListSystemBotConfigInput) SetProjectName(v string) *ListSystemBotConfigInput {
	s.ProjectName = &v
	return s
}

type ListSystemBotConfigOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	Data []*DataForListSystemBotConfigOutput `type:"list" json:",omitempty"`
}

// String returns the string representation
func (s ListSystemBotConfigOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListSystemBotConfigOutput) GoString() string {
	return s.String()
}

// SetData sets the Data field's value.
func (s *ListSystemBotConfigOutput) SetData(v []*DataForListSystemBotConfigOutput) *ListSystemBotConfigOutput {
	s.Data = v
	return s
}

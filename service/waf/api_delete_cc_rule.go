// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package waf

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDeleteCCRuleCommon = "DeleteCCRule"

// DeleteCCRuleCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DeleteCCRuleCommon operation. The "output" return
// value will be populated with the DeleteCCRuleCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeleteCCRuleCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeleteCCRuleCommon Send returns without error.
//
// See DeleteCCRuleCommon for more information on using the DeleteCCRuleCommon
// API call, and error handling.
//
//    // Example sending a request using the DeleteCCRuleCommonRequest method.
//    req, resp := client.DeleteCCRuleCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *WAF) DeleteCCRuleCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDeleteCCRuleCommon,
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

// DeleteCCRuleCommon API operation for WAF.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for WAF's
// API operation DeleteCCRuleCommon for usage and error information.
func (c *WAF) DeleteCCRuleCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DeleteCCRuleCommonRequest(input)
	return out, req.Send()
}

// DeleteCCRuleCommonWithContext is the same as DeleteCCRuleCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteCCRuleCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *WAF) DeleteCCRuleCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DeleteCCRuleCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDeleteCCRule = "DeleteCCRule"

// DeleteCCRuleRequest generates a "volcengine/request.Request" representing the
// client's request for the DeleteCCRule operation. The "output" return
// value will be populated with the DeleteCCRuleCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeleteCCRuleCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeleteCCRuleCommon Send returns without error.
//
// See DeleteCCRule for more information on using the DeleteCCRule
// API call, and error handling.
//
//    // Example sending a request using the DeleteCCRuleRequest method.
//    req, resp := client.DeleteCCRuleRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *WAF) DeleteCCRuleRequest(input *DeleteCCRuleInput) (req *request.Request, output *DeleteCCRuleOutput) {
	op := &request.Operation{
		Name:       opDeleteCCRule,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteCCRuleInput{}
	}

	output = &DeleteCCRuleOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// DeleteCCRule API operation for WAF.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for WAF's
// API operation DeleteCCRule for usage and error information.
func (c *WAF) DeleteCCRule(input *DeleteCCRuleInput) (*DeleteCCRuleOutput, error) {
	req, out := c.DeleteCCRuleRequest(input)
	return out, req.Send()
}

// DeleteCCRuleWithContext is the same as DeleteCCRule with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteCCRule for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *WAF) DeleteCCRuleWithContext(ctx volcengine.Context, input *DeleteCCRuleInput, opts ...request.Option) (*DeleteCCRuleOutput, error) {
	req, out := c.DeleteCCRuleRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DeleteCCRuleInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	// Host is a required field
	Host *string `type:"string" json:",omitempty" required:"true"`

	// ID is a required field
	ID *int32 `type:"int32" json:",omitempty" required:"true"`

	ProjectName *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s DeleteCCRuleInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteCCRuleInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteCCRuleInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DeleteCCRuleInput"}
	if s.Host == nil {
		invalidParams.Add(request.NewErrParamRequired("Host"))
	}
	if s.ID == nil {
		invalidParams.Add(request.NewErrParamRequired("ID"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetHost sets the Host field's value.
func (s *DeleteCCRuleInput) SetHost(v string) *DeleteCCRuleInput {
	s.Host = &v
	return s
}

// SetID sets the ID field's value.
func (s *DeleteCCRuleInput) SetID(v int32) *DeleteCCRuleInput {
	s.ID = &v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *DeleteCCRuleInput) SetProjectName(v string) *DeleteCCRuleInput {
	s.ProjectName = &v
	return s
}

type DeleteCCRuleOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s DeleteCCRuleOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteCCRuleOutput) GoString() string {
	return s.String()
}

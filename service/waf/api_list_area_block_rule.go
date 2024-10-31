// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package waf

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opListAreaBlockRuleCommon = "ListAreaBlockRule"

// ListAreaBlockRuleCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ListAreaBlockRuleCommon operation. The "output" return
// value will be populated with the ListAreaBlockRuleCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListAreaBlockRuleCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListAreaBlockRuleCommon Send returns without error.
//
// See ListAreaBlockRuleCommon for more information on using the ListAreaBlockRuleCommon
// API call, and error handling.
//
//    // Example sending a request using the ListAreaBlockRuleCommonRequest method.
//    req, resp := client.ListAreaBlockRuleCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *WAF) ListAreaBlockRuleCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opListAreaBlockRuleCommon,
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

// ListAreaBlockRuleCommon API operation for WAF.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for WAF's
// API operation ListAreaBlockRuleCommon for usage and error information.
func (c *WAF) ListAreaBlockRuleCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ListAreaBlockRuleCommonRequest(input)
	return out, req.Send()
}

// ListAreaBlockRuleCommonWithContext is the same as ListAreaBlockRuleCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ListAreaBlockRuleCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *WAF) ListAreaBlockRuleCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ListAreaBlockRuleCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opListAreaBlockRule = "ListAreaBlockRule"

// ListAreaBlockRuleRequest generates a "volcengine/request.Request" representing the
// client's request for the ListAreaBlockRule operation. The "output" return
// value will be populated with the ListAreaBlockRuleCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListAreaBlockRuleCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListAreaBlockRuleCommon Send returns without error.
//
// See ListAreaBlockRule for more information on using the ListAreaBlockRule
// API call, and error handling.
//
//    // Example sending a request using the ListAreaBlockRuleRequest method.
//    req, resp := client.ListAreaBlockRuleRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *WAF) ListAreaBlockRuleRequest(input *ListAreaBlockRuleInput) (req *request.Request, output *ListAreaBlockRuleOutput) {
	op := &request.Operation{
		Name:       opListAreaBlockRule,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListAreaBlockRuleInput{}
	}

	output = &ListAreaBlockRuleOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// ListAreaBlockRule API operation for WAF.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for WAF's
// API operation ListAreaBlockRule for usage and error information.
func (c *WAF) ListAreaBlockRule(input *ListAreaBlockRuleInput) (*ListAreaBlockRuleOutput, error) {
	req, out := c.ListAreaBlockRuleRequest(input)
	return out, req.Send()
}

// ListAreaBlockRuleWithContext is the same as ListAreaBlockRule with the addition of
// the ability to pass a context and additional request options.
//
// See ListAreaBlockRule for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *WAF) ListAreaBlockRuleWithContext(ctx volcengine.Context, input *ListAreaBlockRuleInput, opts ...request.Option) (*ListAreaBlockRuleOutput, error) {
	req, out := c.ListAreaBlockRuleRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ListAreaBlockRuleInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	// Host is a required field
	Host *string `type:"string" json:",omitempty" required:"true"`

	ProjectName *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s ListAreaBlockRuleInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListAreaBlockRuleInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListAreaBlockRuleInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ListAreaBlockRuleInput"}
	if s.Host == nil {
		invalidParams.Add(request.NewErrParamRequired("Host"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetHost sets the Host field's value.
func (s *ListAreaBlockRuleInput) SetHost(v string) *ListAreaBlockRuleInput {
	s.Host = &v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *ListAreaBlockRuleInput) SetProjectName(v string) *ListAreaBlockRuleInput {
	s.ProjectName = &v
	return s
}

type ListAreaBlockRuleOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	Action *string `type:"string" json:",omitempty"`

	Country []*string `type:"list" json:",omitempty"`

	RuleTag *string `type:"string" json:",omitempty"`

	SubRegion []*string `type:"list" json:",omitempty"`
}

// String returns the string representation
func (s ListAreaBlockRuleOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListAreaBlockRuleOutput) GoString() string {
	return s.String()
}

// SetAction sets the Action field's value.
func (s *ListAreaBlockRuleOutput) SetAction(v string) *ListAreaBlockRuleOutput {
	s.Action = &v
	return s
}

// SetCountry sets the Country field's value.
func (s *ListAreaBlockRuleOutput) SetCountry(v []*string) *ListAreaBlockRuleOutput {
	s.Country = v
	return s
}

// SetRuleTag sets the RuleTag field's value.
func (s *ListAreaBlockRuleOutput) SetRuleTag(v string) *ListAreaBlockRuleOutput {
	s.RuleTag = &v
	return s
}

// SetSubRegion sets the SubRegion field's value.
func (s *ListAreaBlockRuleOutput) SetSubRegion(v []*string) *ListAreaBlockRuleOutput {
	s.SubRegion = v
	return s
}

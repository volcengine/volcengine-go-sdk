// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package waf

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opUpdateAreaBlockRuleCommon = "UpdateAreaBlockRule"

// UpdateAreaBlockRuleCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the UpdateAreaBlockRuleCommon operation. The "output" return
// value will be populated with the UpdateAreaBlockRuleCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned UpdateAreaBlockRuleCommon Request to send the API call to the service.
// the "output" return value is not valid until after UpdateAreaBlockRuleCommon Send returns without error.
//
// See UpdateAreaBlockRuleCommon for more information on using the UpdateAreaBlockRuleCommon
// API call, and error handling.
//
//    // Example sending a request using the UpdateAreaBlockRuleCommonRequest method.
//    req, resp := client.UpdateAreaBlockRuleCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *WAF) UpdateAreaBlockRuleCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opUpdateAreaBlockRuleCommon,
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

// UpdateAreaBlockRuleCommon API operation for WAF.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for WAF's
// API operation UpdateAreaBlockRuleCommon for usage and error information.
func (c *WAF) UpdateAreaBlockRuleCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.UpdateAreaBlockRuleCommonRequest(input)
	return out, req.Send()
}

// UpdateAreaBlockRuleCommonWithContext is the same as UpdateAreaBlockRuleCommon with the addition of
// the ability to pass a context and additional request options.
//
// See UpdateAreaBlockRuleCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *WAF) UpdateAreaBlockRuleCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.UpdateAreaBlockRuleCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opUpdateAreaBlockRule = "UpdateAreaBlockRule"

// UpdateAreaBlockRuleRequest generates a "volcengine/request.Request" representing the
// client's request for the UpdateAreaBlockRule operation. The "output" return
// value will be populated with the UpdateAreaBlockRuleCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned UpdateAreaBlockRuleCommon Request to send the API call to the service.
// the "output" return value is not valid until after UpdateAreaBlockRuleCommon Send returns without error.
//
// See UpdateAreaBlockRule for more information on using the UpdateAreaBlockRule
// API call, and error handling.
//
//    // Example sending a request using the UpdateAreaBlockRuleRequest method.
//    req, resp := client.UpdateAreaBlockRuleRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *WAF) UpdateAreaBlockRuleRequest(input *UpdateAreaBlockRuleInput) (req *request.Request, output *UpdateAreaBlockRuleOutput) {
	op := &request.Operation{
		Name:       opUpdateAreaBlockRule,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateAreaBlockRuleInput{}
	}

	output = &UpdateAreaBlockRuleOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// UpdateAreaBlockRule API operation for WAF.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for WAF's
// API operation UpdateAreaBlockRule for usage and error information.
func (c *WAF) UpdateAreaBlockRule(input *UpdateAreaBlockRuleInput) (*UpdateAreaBlockRuleOutput, error) {
	req, out := c.UpdateAreaBlockRuleRequest(input)
	return out, req.Send()
}

// UpdateAreaBlockRuleWithContext is the same as UpdateAreaBlockRule with the addition of
// the ability to pass a context and additional request options.
//
// See UpdateAreaBlockRule for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *WAF) UpdateAreaBlockRuleWithContext(ctx volcengine.Context, input *UpdateAreaBlockRuleInput, opts ...request.Option) (*UpdateAreaBlockRuleOutput, error) {
	req, out := c.UpdateAreaBlockRuleRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type UpdateAreaBlockRuleInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	// Action is a required field
	Action *string `type:"string" json:",omitempty" required:"true"`

	Country []*string `type:"list" json:",omitempty"`

	// Host is a required field
	Host *string `type:"string" json:",omitempty" required:"true"`

	ProjectName *string `type:"string" json:",omitempty"`

	SubRegion []*string `type:"list" json:",omitempty"`
}

// String returns the string representation
func (s UpdateAreaBlockRuleInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s UpdateAreaBlockRuleInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateAreaBlockRuleInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "UpdateAreaBlockRuleInput"}
	if s.Action == nil {
		invalidParams.Add(request.NewErrParamRequired("Action"))
	}
	if s.Host == nil {
		invalidParams.Add(request.NewErrParamRequired("Host"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetAction sets the Action field's value.
func (s *UpdateAreaBlockRuleInput) SetAction(v string) *UpdateAreaBlockRuleInput {
	s.Action = &v
	return s
}

// SetCountry sets the Country field's value.
func (s *UpdateAreaBlockRuleInput) SetCountry(v []*string) *UpdateAreaBlockRuleInput {
	s.Country = v
	return s
}

// SetHost sets the Host field's value.
func (s *UpdateAreaBlockRuleInput) SetHost(v string) *UpdateAreaBlockRuleInput {
	s.Host = &v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *UpdateAreaBlockRuleInput) SetProjectName(v string) *UpdateAreaBlockRuleInput {
	s.ProjectName = &v
	return s
}

// SetSubRegion sets the SubRegion field's value.
func (s *UpdateAreaBlockRuleInput) SetSubRegion(v []*string) *UpdateAreaBlockRuleInput {
	s.SubRegion = v
	return s
}

type UpdateAreaBlockRuleOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s UpdateAreaBlockRuleOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s UpdateAreaBlockRuleOutput) GoString() string {
	return s.String()
}

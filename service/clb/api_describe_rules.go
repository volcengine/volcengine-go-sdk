// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package clb

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribeRulesCommon = "DescribeRules"

// DescribeRulesCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeRulesCommon operation. The "output" return
// value will be populated with the DescribeRulesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeRulesCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeRulesCommon Send returns without error.
//
// See DescribeRulesCommon for more information on using the DescribeRulesCommon
// API call, and error handling.
//
//	// Example sending a request using the DescribeRulesCommonRequest method.
//	req, resp := client.DescribeRulesCommonRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *CLB) DescribeRulesCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeRulesCommon,
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

// DescribeRulesCommon API operation for CLB.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CLB's
// API operation DescribeRulesCommon for usage and error information.
func (c *CLB) DescribeRulesCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeRulesCommonRequest(input)
	return out, req.Send()
}

// DescribeRulesCommonWithContext is the same as DescribeRulesCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeRulesCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CLB) DescribeRulesCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeRulesCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeRules = "DescribeRules"

// DescribeRulesRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeRules operation. The "output" return
// value will be populated with the DescribeRulesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeRulesCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeRulesCommon Send returns without error.
//
// See DescribeRules for more information on using the DescribeRules
// API call, and error handling.
//
//	// Example sending a request using the DescribeRulesRequest method.
//	req, resp := client.DescribeRulesRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *CLB) DescribeRulesRequest(input *DescribeRulesInput) (req *request.Request, output *DescribeRulesOutput) {
	op := &request.Operation{
		Name:       opDescribeRules,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeRulesInput{}
	}

	output = &DescribeRulesOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DescribeRules API operation for CLB.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CLB's
// API operation DescribeRules for usage and error information.
func (c *CLB) DescribeRules(input *DescribeRulesInput) (*DescribeRulesOutput, error) {
	req, out := c.DescribeRulesRequest(input)
	return out, req.Send()
}

// DescribeRulesWithContext is the same as DescribeRules with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeRules for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CLB) DescribeRulesWithContext(ctx volcengine.Context, input *DescribeRulesInput, opts ...request.Option) (*DescribeRulesOutput, error) {
	req, out := c.DescribeRulesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DescribeRulesInput struct {
	_ struct{} `type:"structure"`

	// ListenerId is a required field
	ListenerId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeRulesInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeRulesInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeRulesInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DescribeRulesInput"}
	if s.ListenerId == nil {
		invalidParams.Add(request.NewErrParamRequired("ListenerId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetListenerId sets the ListenerId field's value.
func (s *DescribeRulesInput) SetListenerId(v string) *DescribeRulesInput {
	s.ListenerId = &v
	return s
}

type DescribeRulesOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	RequestId *string `type:"string"`

	Rules []*RuleForDescribeRulesOutput `type:"list"`
}

// String returns the string representation
func (s DescribeRulesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeRulesOutput) GoString() string {
	return s.String()
}

// SetRequestId sets the RequestId field's value.
func (s *DescribeRulesOutput) SetRequestId(v string) *DescribeRulesOutput {
	s.RequestId = &v
	return s
}

// SetRules sets the Rules field's value.
func (s *DescribeRulesOutput) SetRules(v []*RuleForDescribeRulesOutput) *DescribeRulesOutput {
	s.Rules = v
	return s
}

type RuleForDescribeRulesOutput struct {
	_ struct{} `type:"structure"`

	Description *string `type:"string"`

	Domain *string `type:"string"`

	RuleId *string `type:"string"`

	ServerGroupId *string `type:"string"`

	Url *string `type:"string"`
}

// String returns the string representation
func (s RuleForDescribeRulesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s RuleForDescribeRulesOutput) GoString() string {
	return s.String()
}

// SetDescription sets the Description field's value.
func (s *RuleForDescribeRulesOutput) SetDescription(v string) *RuleForDescribeRulesOutput {
	s.Description = &v
	return s
}

// SetDomain sets the Domain field's value.
func (s *RuleForDescribeRulesOutput) SetDomain(v string) *RuleForDescribeRulesOutput {
	s.Domain = &v
	return s
}

// SetRuleId sets the RuleId field's value.
func (s *RuleForDescribeRulesOutput) SetRuleId(v string) *RuleForDescribeRulesOutput {
	s.RuleId = &v
	return s
}

// SetServerGroupId sets the ServerGroupId field's value.
func (s *RuleForDescribeRulesOutput) SetServerGroupId(v string) *RuleForDescribeRulesOutput {
	s.ServerGroupId = &v
	return s
}

// SetUrl sets the Url field's value.
func (s *RuleForDescribeRulesOutput) SetUrl(v string) *RuleForDescribeRulesOutput {
	s.Url = &v
	return s
}

// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package alb

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
//    // Example sending a request using the DescribeRulesCommonRequest method.
//    req, resp := client.DescribeRulesCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ALB) DescribeRulesCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
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

// DescribeRulesCommon API operation for ALB.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ALB's
// API operation DescribeRulesCommon for usage and error information.
func (c *ALB) DescribeRulesCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
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
func (c *ALB) DescribeRulesCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
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
//    // Example sending a request using the DescribeRulesRequest method.
//    req, resp := client.DescribeRulesRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ALB) DescribeRulesRequest(input *DescribeRulesInput) (req *request.Request, output *DescribeRulesOutput) {
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

// DescribeRules API operation for ALB.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ALB's
// API operation DescribeRules for usage and error information.
func (c *ALB) DescribeRules(input *DescribeRulesInput) (*DescribeRulesOutput, error) {
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
func (c *ALB) DescribeRulesWithContext(ctx volcengine.Context, input *DescribeRulesInput, opts ...request.Option) (*DescribeRulesOutput, error) {
	req, out := c.DescribeRulesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DescribeRulesInput struct {
	_ struct{} `type:"structure"`

	ListenerId *string `type:"string"`
}

// String returns the string representation
func (s DescribeRulesInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeRulesInput) GoString() string {
	return s.String()
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

type ForwardGroupConfigForDescribeRulesOutput struct {
	_ struct{} `type:"structure"`

	ServerGroupTuples []*ServerGroupTupleForDescribeRulesOutput `type:"list"`

	StickySessionEnabled *string `type:"string"`

	StickySessionTimeout *int64 `type:"integer"`
}

// String returns the string representation
func (s ForwardGroupConfigForDescribeRulesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ForwardGroupConfigForDescribeRulesOutput) GoString() string {
	return s.String()
}

// SetServerGroupTuples sets the ServerGroupTuples field's value.
func (s *ForwardGroupConfigForDescribeRulesOutput) SetServerGroupTuples(v []*ServerGroupTupleForDescribeRulesOutput) *ForwardGroupConfigForDescribeRulesOutput {
	s.ServerGroupTuples = v
	return s
}

// SetStickySessionEnabled sets the StickySessionEnabled field's value.
func (s *ForwardGroupConfigForDescribeRulesOutput) SetStickySessionEnabled(v string) *ForwardGroupConfigForDescribeRulesOutput {
	s.StickySessionEnabled = &v
	return s
}

// SetStickySessionTimeout sets the StickySessionTimeout field's value.
func (s *ForwardGroupConfigForDescribeRulesOutput) SetStickySessionTimeout(v int64) *ForwardGroupConfigForDescribeRulesOutput {
	s.StickySessionTimeout = &v
	return s
}

type RedirectConfigForDescribeRulesOutput struct {
	_ struct{} `type:"structure"`

	RedirectDomain *string `type:"string"`

	RedirectHttpCode *string `type:"string"`

	RedirectPort *string `type:"string"`

	RedirectProtocol *string `type:"string"`

	RedirectUri *string `type:"string"`
}

// String returns the string representation
func (s RedirectConfigForDescribeRulesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s RedirectConfigForDescribeRulesOutput) GoString() string {
	return s.String()
}

// SetRedirectDomain sets the RedirectDomain field's value.
func (s *RedirectConfigForDescribeRulesOutput) SetRedirectDomain(v string) *RedirectConfigForDescribeRulesOutput {
	s.RedirectDomain = &v
	return s
}

// SetRedirectHttpCode sets the RedirectHttpCode field's value.
func (s *RedirectConfigForDescribeRulesOutput) SetRedirectHttpCode(v string) *RedirectConfigForDescribeRulesOutput {
	s.RedirectHttpCode = &v
	return s
}

// SetRedirectPort sets the RedirectPort field's value.
func (s *RedirectConfigForDescribeRulesOutput) SetRedirectPort(v string) *RedirectConfigForDescribeRulesOutput {
	s.RedirectPort = &v
	return s
}

// SetRedirectProtocol sets the RedirectProtocol field's value.
func (s *RedirectConfigForDescribeRulesOutput) SetRedirectProtocol(v string) *RedirectConfigForDescribeRulesOutput {
	s.RedirectProtocol = &v
	return s
}

// SetRedirectUri sets the RedirectUri field's value.
func (s *RedirectConfigForDescribeRulesOutput) SetRedirectUri(v string) *RedirectConfigForDescribeRulesOutput {
	s.RedirectUri = &v
	return s
}

type RewriteConfigForDescribeRulesOutput struct {
	_ struct{} `type:"structure"`

	RewritePath *string `type:"string"`
}

// String returns the string representation
func (s RewriteConfigForDescribeRulesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s RewriteConfigForDescribeRulesOutput) GoString() string {
	return s.String()
}

// SetRewritePath sets the RewritePath field's value.
func (s *RewriteConfigForDescribeRulesOutput) SetRewritePath(v string) *RewriteConfigForDescribeRulesOutput {
	s.RewritePath = &v
	return s
}

type RuleForDescribeRulesOutput struct {
	_ struct{} `type:"structure"`

	Description *string `type:"string"`

	Domain *string `type:"string"`

	ForwardGroupConfig *ForwardGroupConfigForDescribeRulesOutput `type:"structure"`

	RedirectConfig *RedirectConfigForDescribeRulesOutput `type:"structure"`

	RewriteConfig *RewriteConfigForDescribeRulesOutput `type:"structure"`

	RewriteEnabled *string `type:"string"`

	RuleAction *string `type:"string"`

	RuleId *string `type:"string"`

	ServerGroupId *string `type:"string"`

	TrafficLimitEnabled *string `type:"string"`

	TrafficLimitQPS *int64 `type:"integer"`

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

// SetForwardGroupConfig sets the ForwardGroupConfig field's value.
func (s *RuleForDescribeRulesOutput) SetForwardGroupConfig(v *ForwardGroupConfigForDescribeRulesOutput) *RuleForDescribeRulesOutput {
	s.ForwardGroupConfig = v
	return s
}

// SetRedirectConfig sets the RedirectConfig field's value.
func (s *RuleForDescribeRulesOutput) SetRedirectConfig(v *RedirectConfigForDescribeRulesOutput) *RuleForDescribeRulesOutput {
	s.RedirectConfig = v
	return s
}

// SetRewriteConfig sets the RewriteConfig field's value.
func (s *RuleForDescribeRulesOutput) SetRewriteConfig(v *RewriteConfigForDescribeRulesOutput) *RuleForDescribeRulesOutput {
	s.RewriteConfig = v
	return s
}

// SetRewriteEnabled sets the RewriteEnabled field's value.
func (s *RuleForDescribeRulesOutput) SetRewriteEnabled(v string) *RuleForDescribeRulesOutput {
	s.RewriteEnabled = &v
	return s
}

// SetRuleAction sets the RuleAction field's value.
func (s *RuleForDescribeRulesOutput) SetRuleAction(v string) *RuleForDescribeRulesOutput {
	s.RuleAction = &v
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

// SetTrafficLimitEnabled sets the TrafficLimitEnabled field's value.
func (s *RuleForDescribeRulesOutput) SetTrafficLimitEnabled(v string) *RuleForDescribeRulesOutput {
	s.TrafficLimitEnabled = &v
	return s
}

// SetTrafficLimitQPS sets the TrafficLimitQPS field's value.
func (s *RuleForDescribeRulesOutput) SetTrafficLimitQPS(v int64) *RuleForDescribeRulesOutput {
	s.TrafficLimitQPS = &v
	return s
}

// SetUrl sets the Url field's value.
func (s *RuleForDescribeRulesOutput) SetUrl(v string) *RuleForDescribeRulesOutput {
	s.Url = &v
	return s
}

type ServerGroupTupleForDescribeRulesOutput struct {
	_ struct{} `type:"structure"`

	ServerGroupId *string `type:"string"`

	Weight *int64 `type:"integer"`
}

// String returns the string representation
func (s ServerGroupTupleForDescribeRulesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ServerGroupTupleForDescribeRulesOutput) GoString() string {
	return s.String()
}

// SetServerGroupId sets the ServerGroupId field's value.
func (s *ServerGroupTupleForDescribeRulesOutput) SetServerGroupId(v string) *ServerGroupTupleForDescribeRulesOutput {
	s.ServerGroupId = &v
	return s
}

// SetWeight sets the Weight field's value.
func (s *ServerGroupTupleForDescribeRulesOutput) SetWeight(v int64) *ServerGroupTupleForDescribeRulesOutput {
	s.Weight = &v
	return s
}

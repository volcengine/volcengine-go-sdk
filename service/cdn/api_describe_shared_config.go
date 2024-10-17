// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package cdn

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribeSharedConfigCommon = "DescribeSharedConfig"

// DescribeSharedConfigCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeSharedConfigCommon operation. The "output" return
// value will be populated with the DescribeSharedConfigCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeSharedConfigCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeSharedConfigCommon Send returns without error.
//
// See DescribeSharedConfigCommon for more information on using the DescribeSharedConfigCommon
// API call, and error handling.
//
//    // Example sending a request using the DescribeSharedConfigCommonRequest method.
//    req, resp := client.DescribeSharedConfigCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CDN) DescribeSharedConfigCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeSharedConfigCommon,
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

// DescribeSharedConfigCommon API operation for CDN.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CDN's
// API operation DescribeSharedConfigCommon for usage and error information.
func (c *CDN) DescribeSharedConfigCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeSharedConfigCommonRequest(input)
	return out, req.Send()
}

// DescribeSharedConfigCommonWithContext is the same as DescribeSharedConfigCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeSharedConfigCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CDN) DescribeSharedConfigCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeSharedConfigCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeSharedConfig = "DescribeSharedConfig"

// DescribeSharedConfigRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeSharedConfig operation. The "output" return
// value will be populated with the DescribeSharedConfigCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeSharedConfigCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeSharedConfigCommon Send returns without error.
//
// See DescribeSharedConfig for more information on using the DescribeSharedConfig
// API call, and error handling.
//
//    // Example sending a request using the DescribeSharedConfigRequest method.
//    req, resp := client.DescribeSharedConfigRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CDN) DescribeSharedConfigRequest(input *DescribeSharedConfigInput) (req *request.Request, output *DescribeSharedConfigOutput) {
	op := &request.Operation{
		Name:       opDescribeSharedConfig,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeSharedConfigInput{}
	}

	output = &DescribeSharedConfigOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// DescribeSharedConfig API operation for CDN.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CDN's
// API operation DescribeSharedConfig for usage and error information.
func (c *CDN) DescribeSharedConfig(input *DescribeSharedConfigInput) (*DescribeSharedConfigOutput, error) {
	req, out := c.DescribeSharedConfigRequest(input)
	return out, req.Send()
}

// DescribeSharedConfigWithContext is the same as DescribeSharedConfig with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeSharedConfig for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CDN) DescribeSharedConfigWithContext(ctx volcengine.Context, input *DescribeSharedConfigInput, opts ...request.Option) (*DescribeSharedConfigOutput, error) {
	req, out := c.DescribeSharedConfigRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type AllowIpAccessRuleForDescribeSharedConfigOutput struct {
	_ struct{} `type:"structure"`

	Option *string `type:"string"`

	Rules []*string `type:"list"`
}

// String returns the string representation
func (s AllowIpAccessRuleForDescribeSharedConfigOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s AllowIpAccessRuleForDescribeSharedConfigOutput) GoString() string {
	return s.String()
}

// SetOption sets the Option field's value.
func (s *AllowIpAccessRuleForDescribeSharedConfigOutput) SetOption(v string) *AllowIpAccessRuleForDescribeSharedConfigOutput {
	s.Option = &v
	return s
}

// SetRules sets the Rules field's value.
func (s *AllowIpAccessRuleForDescribeSharedConfigOutput) SetRules(v []*string) *AllowIpAccessRuleForDescribeSharedConfigOutput {
	s.Rules = v
	return s
}

type AllowRefererAccessRuleForDescribeSharedConfigOutput struct {
	_ struct{} `type:"structure"`

	AllowEmpty *bool `type:"boolean"`

	CommonType *CommonTypeForDescribeSharedConfigOutput `type:"structure"`
}

// String returns the string representation
func (s AllowRefererAccessRuleForDescribeSharedConfigOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s AllowRefererAccessRuleForDescribeSharedConfigOutput) GoString() string {
	return s.String()
}

// SetAllowEmpty sets the AllowEmpty field's value.
func (s *AllowRefererAccessRuleForDescribeSharedConfigOutput) SetAllowEmpty(v bool) *AllowRefererAccessRuleForDescribeSharedConfigOutput {
	s.AllowEmpty = &v
	return s
}

// SetCommonType sets the CommonType field's value.
func (s *AllowRefererAccessRuleForDescribeSharedConfigOutput) SetCommonType(v *CommonTypeForDescribeSharedConfigOutput) *AllowRefererAccessRuleForDescribeSharedConfigOutput {
	s.CommonType = v
	return s
}

type CommonMatchListForDescribeSharedConfigOutput struct {
	_ struct{} `type:"structure"`

	CommonType *CommonTypeForDescribeSharedConfigOutput `type:"structure"`
}

// String returns the string representation
func (s CommonMatchListForDescribeSharedConfigOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CommonMatchListForDescribeSharedConfigOutput) GoString() string {
	return s.String()
}

// SetCommonType sets the CommonType field's value.
func (s *CommonMatchListForDescribeSharedConfigOutput) SetCommonType(v *CommonTypeForDescribeSharedConfigOutput) *CommonMatchListForDescribeSharedConfigOutput {
	s.CommonType = v
	return s
}

type CommonTypeForDescribeSharedConfigOutput struct {
	_ struct{} `type:"structure"`

	IgnoreCase *bool `type:"boolean"`

	Option *string `type:"string"`

	Rules []*string `type:"list"`
}

// String returns the string representation
func (s CommonTypeForDescribeSharedConfigOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CommonTypeForDescribeSharedConfigOutput) GoString() string {
	return s.String()
}

// SetIgnoreCase sets the IgnoreCase field's value.
func (s *CommonTypeForDescribeSharedConfigOutput) SetIgnoreCase(v bool) *CommonTypeForDescribeSharedConfigOutput {
	s.IgnoreCase = &v
	return s
}

// SetOption sets the Option field's value.
func (s *CommonTypeForDescribeSharedConfigOutput) SetOption(v string) *CommonTypeForDescribeSharedConfigOutput {
	s.Option = &v
	return s
}

// SetRules sets the Rules field's value.
func (s *CommonTypeForDescribeSharedConfigOutput) SetRules(v []*string) *CommonTypeForDescribeSharedConfigOutput {
	s.Rules = v
	return s
}

type DenyIpAccessRuleForDescribeSharedConfigOutput struct {
	_ struct{} `type:"structure"`

	Option *string `type:"string"`

	Rules []*string `type:"list"`
}

// String returns the string representation
func (s DenyIpAccessRuleForDescribeSharedConfigOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DenyIpAccessRuleForDescribeSharedConfigOutput) GoString() string {
	return s.String()
}

// SetOption sets the Option field's value.
func (s *DenyIpAccessRuleForDescribeSharedConfigOutput) SetOption(v string) *DenyIpAccessRuleForDescribeSharedConfigOutput {
	s.Option = &v
	return s
}

// SetRules sets the Rules field's value.
func (s *DenyIpAccessRuleForDescribeSharedConfigOutput) SetRules(v []*string) *DenyIpAccessRuleForDescribeSharedConfigOutput {
	s.Rules = v
	return s
}

type DenyRefererAccessRuleForDescribeSharedConfigOutput struct {
	_ struct{} `type:"structure"`

	AllowEmpty *bool `type:"boolean"`

	CommonType *CommonTypeForDescribeSharedConfigOutput `type:"structure"`
}

// String returns the string representation
func (s DenyRefererAccessRuleForDescribeSharedConfigOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DenyRefererAccessRuleForDescribeSharedConfigOutput) GoString() string {
	return s.String()
}

// SetAllowEmpty sets the AllowEmpty field's value.
func (s *DenyRefererAccessRuleForDescribeSharedConfigOutput) SetAllowEmpty(v bool) *DenyRefererAccessRuleForDescribeSharedConfigOutput {
	s.AllowEmpty = &v
	return s
}

// SetCommonType sets the CommonType field's value.
func (s *DenyRefererAccessRuleForDescribeSharedConfigOutput) SetCommonType(v *CommonTypeForDescribeSharedConfigOutput) *DenyRefererAccessRuleForDescribeSharedConfigOutput {
	s.CommonType = v
	return s
}

type DescribeSharedConfigInput struct {
	_ struct{} `type:"structure"`

	ConfigName *string `type:"string"`
}

// String returns the string representation
func (s DescribeSharedConfigInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeSharedConfigInput) GoString() string {
	return s.String()
}

// SetConfigName sets the ConfigName field's value.
func (s *DescribeSharedConfigInput) SetConfigName(v string) *DescribeSharedConfigInput {
	s.ConfigName = &v
	return s
}

type DescribeSharedConfigOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	AllowIpAccessRule *AllowIpAccessRuleForDescribeSharedConfigOutput `type:"structure"`

	AllowRefererAccessRule *AllowRefererAccessRuleForDescribeSharedConfigOutput `type:"structure"`

	CommonMatchList *CommonMatchListForDescribeSharedConfigOutput `type:"structure"`

	ConfigName *string `type:"string"`

	ConfigType *string `type:"string"`

	DenyIpAccessRule *DenyIpAccessRuleForDescribeSharedConfigOutput `type:"structure"`

	DenyRefererAccessRule *DenyRefererAccessRuleForDescribeSharedConfigOutput `type:"structure"`

	ErrorPageRule *ErrorPageRuleForDescribeSharedConfigOutput `type:"structure"`

	Project *string `type:"string"`
}

// String returns the string representation
func (s DescribeSharedConfigOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeSharedConfigOutput) GoString() string {
	return s.String()
}

// SetAllowIpAccessRule sets the AllowIpAccessRule field's value.
func (s *DescribeSharedConfigOutput) SetAllowIpAccessRule(v *AllowIpAccessRuleForDescribeSharedConfigOutput) *DescribeSharedConfigOutput {
	s.AllowIpAccessRule = v
	return s
}

// SetAllowRefererAccessRule sets the AllowRefererAccessRule field's value.
func (s *DescribeSharedConfigOutput) SetAllowRefererAccessRule(v *AllowRefererAccessRuleForDescribeSharedConfigOutput) *DescribeSharedConfigOutput {
	s.AllowRefererAccessRule = v
	return s
}

// SetCommonMatchList sets the CommonMatchList field's value.
func (s *DescribeSharedConfigOutput) SetCommonMatchList(v *CommonMatchListForDescribeSharedConfigOutput) *DescribeSharedConfigOutput {
	s.CommonMatchList = v
	return s
}

// SetConfigName sets the ConfigName field's value.
func (s *DescribeSharedConfigOutput) SetConfigName(v string) *DescribeSharedConfigOutput {
	s.ConfigName = &v
	return s
}

// SetConfigType sets the ConfigType field's value.
func (s *DescribeSharedConfigOutput) SetConfigType(v string) *DescribeSharedConfigOutput {
	s.ConfigType = &v
	return s
}

// SetDenyIpAccessRule sets the DenyIpAccessRule field's value.
func (s *DescribeSharedConfigOutput) SetDenyIpAccessRule(v *DenyIpAccessRuleForDescribeSharedConfigOutput) *DescribeSharedConfigOutput {
	s.DenyIpAccessRule = v
	return s
}

// SetDenyRefererAccessRule sets the DenyRefererAccessRule field's value.
func (s *DescribeSharedConfigOutput) SetDenyRefererAccessRule(v *DenyRefererAccessRuleForDescribeSharedConfigOutput) *DescribeSharedConfigOutput {
	s.DenyRefererAccessRule = v
	return s
}

// SetErrorPageRule sets the ErrorPageRule field's value.
func (s *DescribeSharedConfigOutput) SetErrorPageRule(v *ErrorPageRuleForDescribeSharedConfigOutput) *DescribeSharedConfigOutput {
	s.ErrorPageRule = v
	return s
}

// SetProject sets the Project field's value.
func (s *DescribeSharedConfigOutput) SetProject(v string) *DescribeSharedConfigOutput {
	s.Project = &v
	return s
}

type ErrorPageRuleForDescribeSharedConfigOutput struct {
	_ struct{} `type:"structure"`

	RuleContent *string `type:"string"`
}

// String returns the string representation
func (s ErrorPageRuleForDescribeSharedConfigOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ErrorPageRuleForDescribeSharedConfigOutput) GoString() string {
	return s.String()
}

// SetRuleContent sets the RuleContent field's value.
func (s *ErrorPageRuleForDescribeSharedConfigOutput) SetRuleContent(v string) *ErrorPageRuleForDescribeSharedConfigOutput {
	s.RuleContent = &v
	return s
}

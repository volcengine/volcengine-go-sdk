// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package fwcenter

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDeleteVpcFirewallAclRuleCommon = "DeleteVpcFirewallAclRule"

// DeleteVpcFirewallAclRuleCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DeleteVpcFirewallAclRuleCommon operation. The "output" return
// value will be populated with the DeleteVpcFirewallAclRuleCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeleteVpcFirewallAclRuleCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeleteVpcFirewallAclRuleCommon Send returns without error.
//
// See DeleteVpcFirewallAclRuleCommon for more information on using the DeleteVpcFirewallAclRuleCommon
// API call, and error handling.
//
//    // Example sending a request using the DeleteVpcFirewallAclRuleCommonRequest method.
//    req, resp := client.DeleteVpcFirewallAclRuleCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *FWCENTER) DeleteVpcFirewallAclRuleCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDeleteVpcFirewallAclRuleCommon,
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

// DeleteVpcFirewallAclRuleCommon API operation for FWCENTER.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for FWCENTER's
// API operation DeleteVpcFirewallAclRuleCommon for usage and error information.
func (c *FWCENTER) DeleteVpcFirewallAclRuleCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DeleteVpcFirewallAclRuleCommonRequest(input)
	return out, req.Send()
}

// DeleteVpcFirewallAclRuleCommonWithContext is the same as DeleteVpcFirewallAclRuleCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteVpcFirewallAclRuleCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *FWCENTER) DeleteVpcFirewallAclRuleCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DeleteVpcFirewallAclRuleCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDeleteVpcFirewallAclRule = "DeleteVpcFirewallAclRule"

// DeleteVpcFirewallAclRuleRequest generates a "volcengine/request.Request" representing the
// client's request for the DeleteVpcFirewallAclRule operation. The "output" return
// value will be populated with the DeleteVpcFirewallAclRuleCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeleteVpcFirewallAclRuleCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeleteVpcFirewallAclRuleCommon Send returns without error.
//
// See DeleteVpcFirewallAclRule for more information on using the DeleteVpcFirewallAclRule
// API call, and error handling.
//
//    // Example sending a request using the DeleteVpcFirewallAclRuleRequest method.
//    req, resp := client.DeleteVpcFirewallAclRuleRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *FWCENTER) DeleteVpcFirewallAclRuleRequest(input *DeleteVpcFirewallAclRuleInput) (req *request.Request, output *DeleteVpcFirewallAclRuleOutput) {
	op := &request.Operation{
		Name:       opDeleteVpcFirewallAclRule,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteVpcFirewallAclRuleInput{}
	}

	output = &DeleteVpcFirewallAclRuleOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// DeleteVpcFirewallAclRule API operation for FWCENTER.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for FWCENTER's
// API operation DeleteVpcFirewallAclRule for usage and error information.
func (c *FWCENTER) DeleteVpcFirewallAclRule(input *DeleteVpcFirewallAclRuleInput) (*DeleteVpcFirewallAclRuleOutput, error) {
	req, out := c.DeleteVpcFirewallAclRuleRequest(input)
	return out, req.Send()
}

// DeleteVpcFirewallAclRuleWithContext is the same as DeleteVpcFirewallAclRule with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteVpcFirewallAclRule for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *FWCENTER) DeleteVpcFirewallAclRuleWithContext(ctx volcengine.Context, input *DeleteVpcFirewallAclRuleInput, opts ...request.Option) (*DeleteVpcFirewallAclRuleOutput, error) {
	req, out := c.DeleteVpcFirewallAclRuleRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DeleteVpcFirewallAclRuleInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	RuleIds []*string `type:"list" json:",omitempty"`

	// VpcFirewallId is a required field
	VpcFirewallId *string `type:"string" json:",omitempty" required:"true"`
}

// String returns the string representation
func (s DeleteVpcFirewallAclRuleInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteVpcFirewallAclRuleInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteVpcFirewallAclRuleInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DeleteVpcFirewallAclRuleInput"}
	if s.VpcFirewallId == nil {
		invalidParams.Add(request.NewErrParamRequired("VpcFirewallId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetRuleIds sets the RuleIds field's value.
func (s *DeleteVpcFirewallAclRuleInput) SetRuleIds(v []*string) *DeleteVpcFirewallAclRuleInput {
	s.RuleIds = v
	return s
}

// SetVpcFirewallId sets the VpcFirewallId field's value.
func (s *DeleteVpcFirewallAclRuleInput) SetVpcFirewallId(v string) *DeleteVpcFirewallAclRuleInput {
	s.VpcFirewallId = &v
	return s
}

type DeleteVpcFirewallAclRuleOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	RuleIds []*string `type:"list" json:",omitempty"`
}

// String returns the string representation
func (s DeleteVpcFirewallAclRuleOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteVpcFirewallAclRuleOutput) GoString() string {
	return s.String()
}

// SetRuleIds sets the RuleIds field's value.
func (s *DeleteVpcFirewallAclRuleOutput) SetRuleIds(v []*string) *DeleteVpcFirewallAclRuleOutput {
	s.RuleIds = v
	return s
}

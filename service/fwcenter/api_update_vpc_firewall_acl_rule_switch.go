// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package fwcenter

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opUpdateVpcFirewallAclRuleSwitchCommon = "UpdateVpcFirewallAclRuleSwitch"

// UpdateVpcFirewallAclRuleSwitchCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the UpdateVpcFirewallAclRuleSwitchCommon operation. The "output" return
// value will be populated with the UpdateVpcFirewallAclRuleSwitchCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned UpdateVpcFirewallAclRuleSwitchCommon Request to send the API call to the service.
// the "output" return value is not valid until after UpdateVpcFirewallAclRuleSwitchCommon Send returns without error.
//
// See UpdateVpcFirewallAclRuleSwitchCommon for more information on using the UpdateVpcFirewallAclRuleSwitchCommon
// API call, and error handling.
//
//    // Example sending a request using the UpdateVpcFirewallAclRuleSwitchCommonRequest method.
//    req, resp := client.UpdateVpcFirewallAclRuleSwitchCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *FWCENTER) UpdateVpcFirewallAclRuleSwitchCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opUpdateVpcFirewallAclRuleSwitchCommon,
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

// UpdateVpcFirewallAclRuleSwitchCommon API operation for FWCENTER.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for FWCENTER's
// API operation UpdateVpcFirewallAclRuleSwitchCommon for usage and error information.
func (c *FWCENTER) UpdateVpcFirewallAclRuleSwitchCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.UpdateVpcFirewallAclRuleSwitchCommonRequest(input)
	return out, req.Send()
}

// UpdateVpcFirewallAclRuleSwitchCommonWithContext is the same as UpdateVpcFirewallAclRuleSwitchCommon with the addition of
// the ability to pass a context and additional request options.
//
// See UpdateVpcFirewallAclRuleSwitchCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *FWCENTER) UpdateVpcFirewallAclRuleSwitchCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.UpdateVpcFirewallAclRuleSwitchCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opUpdateVpcFirewallAclRuleSwitch = "UpdateVpcFirewallAclRuleSwitch"

// UpdateVpcFirewallAclRuleSwitchRequest generates a "volcengine/request.Request" representing the
// client's request for the UpdateVpcFirewallAclRuleSwitch operation. The "output" return
// value will be populated with the UpdateVpcFirewallAclRuleSwitchCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned UpdateVpcFirewallAclRuleSwitchCommon Request to send the API call to the service.
// the "output" return value is not valid until after UpdateVpcFirewallAclRuleSwitchCommon Send returns without error.
//
// See UpdateVpcFirewallAclRuleSwitch for more information on using the UpdateVpcFirewallAclRuleSwitch
// API call, and error handling.
//
//    // Example sending a request using the UpdateVpcFirewallAclRuleSwitchRequest method.
//    req, resp := client.UpdateVpcFirewallAclRuleSwitchRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *FWCENTER) UpdateVpcFirewallAclRuleSwitchRequest(input *UpdateVpcFirewallAclRuleSwitchInput) (req *request.Request, output *UpdateVpcFirewallAclRuleSwitchOutput) {
	op := &request.Operation{
		Name:       opUpdateVpcFirewallAclRuleSwitch,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateVpcFirewallAclRuleSwitchInput{}
	}

	output = &UpdateVpcFirewallAclRuleSwitchOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// UpdateVpcFirewallAclRuleSwitch API operation for FWCENTER.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for FWCENTER's
// API operation UpdateVpcFirewallAclRuleSwitch for usage and error information.
func (c *FWCENTER) UpdateVpcFirewallAclRuleSwitch(input *UpdateVpcFirewallAclRuleSwitchInput) (*UpdateVpcFirewallAclRuleSwitchOutput, error) {
	req, out := c.UpdateVpcFirewallAclRuleSwitchRequest(input)
	return out, req.Send()
}

// UpdateVpcFirewallAclRuleSwitchWithContext is the same as UpdateVpcFirewallAclRuleSwitch with the addition of
// the ability to pass a context and additional request options.
//
// See UpdateVpcFirewallAclRuleSwitch for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *FWCENTER) UpdateVpcFirewallAclRuleSwitchWithContext(ctx volcengine.Context, input *UpdateVpcFirewallAclRuleSwitchInput, opts ...request.Option) (*UpdateVpcFirewallAclRuleSwitchOutput, error) {
	req, out := c.UpdateVpcFirewallAclRuleSwitchRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type UpdateVpcFirewallAclRuleSwitchInput struct {
	_ struct{} `type:"structure"`

	RuleIds []*string `type:"list"`

	Status *bool `type:"boolean"`

	// VpcFirewallId is a required field
	VpcFirewallId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s UpdateVpcFirewallAclRuleSwitchInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s UpdateVpcFirewallAclRuleSwitchInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateVpcFirewallAclRuleSwitchInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "UpdateVpcFirewallAclRuleSwitchInput"}
	if s.VpcFirewallId == nil {
		invalidParams.Add(request.NewErrParamRequired("VpcFirewallId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetRuleIds sets the RuleIds field's value.
func (s *UpdateVpcFirewallAclRuleSwitchInput) SetRuleIds(v []*string) *UpdateVpcFirewallAclRuleSwitchInput {
	s.RuleIds = v
	return s
}

// SetStatus sets the Status field's value.
func (s *UpdateVpcFirewallAclRuleSwitchInput) SetStatus(v bool) *UpdateVpcFirewallAclRuleSwitchInput {
	s.Status = &v
	return s
}

// SetVpcFirewallId sets the VpcFirewallId field's value.
func (s *UpdateVpcFirewallAclRuleSwitchInput) SetVpcFirewallId(v string) *UpdateVpcFirewallAclRuleSwitchInput {
	s.VpcFirewallId = &v
	return s
}

type UpdateVpcFirewallAclRuleSwitchOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s UpdateVpcFirewallAclRuleSwitchOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s UpdateVpcFirewallAclRuleSwitchOutput) GoString() string {
	return s.String()
}

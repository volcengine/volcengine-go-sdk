// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package fwcenter

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opModifyVpcFirewallAclRuleCommon = "ModifyVpcFirewallAclRule"

// ModifyVpcFirewallAclRuleCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ModifyVpcFirewallAclRuleCommon operation. The "output" return
// value will be populated with the ModifyVpcFirewallAclRuleCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyVpcFirewallAclRuleCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyVpcFirewallAclRuleCommon Send returns without error.
//
// See ModifyVpcFirewallAclRuleCommon for more information on using the ModifyVpcFirewallAclRuleCommon
// API call, and error handling.
//
//    // Example sending a request using the ModifyVpcFirewallAclRuleCommonRequest method.
//    req, resp := client.ModifyVpcFirewallAclRuleCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *FWCENTER) ModifyVpcFirewallAclRuleCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opModifyVpcFirewallAclRuleCommon,
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

// ModifyVpcFirewallAclRuleCommon API operation for FWCENTER.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for FWCENTER's
// API operation ModifyVpcFirewallAclRuleCommon for usage and error information.
func (c *FWCENTER) ModifyVpcFirewallAclRuleCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ModifyVpcFirewallAclRuleCommonRequest(input)
	return out, req.Send()
}

// ModifyVpcFirewallAclRuleCommonWithContext is the same as ModifyVpcFirewallAclRuleCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyVpcFirewallAclRuleCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *FWCENTER) ModifyVpcFirewallAclRuleCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ModifyVpcFirewallAclRuleCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opModifyVpcFirewallAclRule = "ModifyVpcFirewallAclRule"

// ModifyVpcFirewallAclRuleRequest generates a "volcengine/request.Request" representing the
// client's request for the ModifyVpcFirewallAclRule operation. The "output" return
// value will be populated with the ModifyVpcFirewallAclRuleCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyVpcFirewallAclRuleCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyVpcFirewallAclRuleCommon Send returns without error.
//
// See ModifyVpcFirewallAclRule for more information on using the ModifyVpcFirewallAclRule
// API call, and error handling.
//
//    // Example sending a request using the ModifyVpcFirewallAclRuleRequest method.
//    req, resp := client.ModifyVpcFirewallAclRuleRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *FWCENTER) ModifyVpcFirewallAclRuleRequest(input *ModifyVpcFirewallAclRuleInput) (req *request.Request, output *ModifyVpcFirewallAclRuleOutput) {
	op := &request.Operation{
		Name:       opModifyVpcFirewallAclRule,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ModifyVpcFirewallAclRuleInput{}
	}

	output = &ModifyVpcFirewallAclRuleOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// ModifyVpcFirewallAclRule API operation for FWCENTER.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for FWCENTER's
// API operation ModifyVpcFirewallAclRule for usage and error information.
func (c *FWCENTER) ModifyVpcFirewallAclRule(input *ModifyVpcFirewallAclRuleInput) (*ModifyVpcFirewallAclRuleOutput, error) {
	req, out := c.ModifyVpcFirewallAclRuleRequest(input)
	return out, req.Send()
}

// ModifyVpcFirewallAclRuleWithContext is the same as ModifyVpcFirewallAclRule with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyVpcFirewallAclRule for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *FWCENTER) ModifyVpcFirewallAclRuleWithContext(ctx volcengine.Context, input *ModifyVpcFirewallAclRuleInput, opts ...request.Option) (*ModifyVpcFirewallAclRuleOutput, error) {
	req, out := c.ModifyVpcFirewallAclRuleRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ModifyVpcFirewallAclRuleInput struct {
	_ struct{} `type:"structure"`

	// Action is a required field
	Action *string `type:"string" required:"true"`

	Description *string `type:"string"`

	DestPort *string `type:"string"`

	DestPortType *string `type:"string"`

	// Destination is a required field
	Destination *string `type:"string" required:"true"`

	DestinationGroupType *string `type:"string"`

	// DestinationType is a required field
	DestinationType *string `type:"string" required:"true"`

	// Proto is a required field
	Proto *string `type:"string" required:"true"`

	// RuleId is a required field
	RuleId *string `type:"string" required:"true"`

	// Source is a required field
	Source *string `type:"string" required:"true"`

	SourceGroupType *string `type:"string"`

	// SourceType is a required field
	SourceType *string `type:"string" required:"true"`

	Status *bool `type:"boolean"`

	// VpcFirewallId is a required field
	VpcFirewallId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s ModifyVpcFirewallAclRuleInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyVpcFirewallAclRuleInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ModifyVpcFirewallAclRuleInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ModifyVpcFirewallAclRuleInput"}
	if s.Action == nil {
		invalidParams.Add(request.NewErrParamRequired("Action"))
	}
	if s.Destination == nil {
		invalidParams.Add(request.NewErrParamRequired("Destination"))
	}
	if s.DestinationType == nil {
		invalidParams.Add(request.NewErrParamRequired("DestinationType"))
	}
	if s.Proto == nil {
		invalidParams.Add(request.NewErrParamRequired("Proto"))
	}
	if s.RuleId == nil {
		invalidParams.Add(request.NewErrParamRequired("RuleId"))
	}
	if s.Source == nil {
		invalidParams.Add(request.NewErrParamRequired("Source"))
	}
	if s.SourceType == nil {
		invalidParams.Add(request.NewErrParamRequired("SourceType"))
	}
	if s.VpcFirewallId == nil {
		invalidParams.Add(request.NewErrParamRequired("VpcFirewallId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetAction sets the Action field's value.
func (s *ModifyVpcFirewallAclRuleInput) SetAction(v string) *ModifyVpcFirewallAclRuleInput {
	s.Action = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *ModifyVpcFirewallAclRuleInput) SetDescription(v string) *ModifyVpcFirewallAclRuleInput {
	s.Description = &v
	return s
}

// SetDestPort sets the DestPort field's value.
func (s *ModifyVpcFirewallAclRuleInput) SetDestPort(v string) *ModifyVpcFirewallAclRuleInput {
	s.DestPort = &v
	return s
}

// SetDestPortType sets the DestPortType field's value.
func (s *ModifyVpcFirewallAclRuleInput) SetDestPortType(v string) *ModifyVpcFirewallAclRuleInput {
	s.DestPortType = &v
	return s
}

// SetDestination sets the Destination field's value.
func (s *ModifyVpcFirewallAclRuleInput) SetDestination(v string) *ModifyVpcFirewallAclRuleInput {
	s.Destination = &v
	return s
}

// SetDestinationGroupType sets the DestinationGroupType field's value.
func (s *ModifyVpcFirewallAclRuleInput) SetDestinationGroupType(v string) *ModifyVpcFirewallAclRuleInput {
	s.DestinationGroupType = &v
	return s
}

// SetDestinationType sets the DestinationType field's value.
func (s *ModifyVpcFirewallAclRuleInput) SetDestinationType(v string) *ModifyVpcFirewallAclRuleInput {
	s.DestinationType = &v
	return s
}

// SetProto sets the Proto field's value.
func (s *ModifyVpcFirewallAclRuleInput) SetProto(v string) *ModifyVpcFirewallAclRuleInput {
	s.Proto = &v
	return s
}

// SetRuleId sets the RuleId field's value.
func (s *ModifyVpcFirewallAclRuleInput) SetRuleId(v string) *ModifyVpcFirewallAclRuleInput {
	s.RuleId = &v
	return s
}

// SetSource sets the Source field's value.
func (s *ModifyVpcFirewallAclRuleInput) SetSource(v string) *ModifyVpcFirewallAclRuleInput {
	s.Source = &v
	return s
}

// SetSourceGroupType sets the SourceGroupType field's value.
func (s *ModifyVpcFirewallAclRuleInput) SetSourceGroupType(v string) *ModifyVpcFirewallAclRuleInput {
	s.SourceGroupType = &v
	return s
}

// SetSourceType sets the SourceType field's value.
func (s *ModifyVpcFirewallAclRuleInput) SetSourceType(v string) *ModifyVpcFirewallAclRuleInput {
	s.SourceType = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *ModifyVpcFirewallAclRuleInput) SetStatus(v bool) *ModifyVpcFirewallAclRuleInput {
	s.Status = &v
	return s
}

// SetVpcFirewallId sets the VpcFirewallId field's value.
func (s *ModifyVpcFirewallAclRuleInput) SetVpcFirewallId(v string) *ModifyVpcFirewallAclRuleInput {
	s.VpcFirewallId = &v
	return s
}

type ModifyVpcFirewallAclRuleOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s ModifyVpcFirewallAclRuleOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyVpcFirewallAclRuleOutput) GoString() string {
	return s.String()
}

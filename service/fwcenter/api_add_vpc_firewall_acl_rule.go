// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package fwcenter

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opAddVpcFirewallAclRuleCommon = "AddVpcFirewallAclRule"

// AddVpcFirewallAclRuleCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the AddVpcFirewallAclRuleCommon operation. The "output" return
// value will be populated with the AddVpcFirewallAclRuleCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned AddVpcFirewallAclRuleCommon Request to send the API call to the service.
// the "output" return value is not valid until after AddVpcFirewallAclRuleCommon Send returns without error.
//
// See AddVpcFirewallAclRuleCommon for more information on using the AddVpcFirewallAclRuleCommon
// API call, and error handling.
//
//	// Example sending a request using the AddVpcFirewallAclRuleCommonRequest method.
//	req, resp := client.AddVpcFirewallAclRuleCommonRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *FWCENTER) AddVpcFirewallAclRuleCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opAddVpcFirewallAclRuleCommon,
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

// AddVpcFirewallAclRuleCommon API operation for FWCENTER.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for FWCENTER's
// API operation AddVpcFirewallAclRuleCommon for usage and error information.
func (c *FWCENTER) AddVpcFirewallAclRuleCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.AddVpcFirewallAclRuleCommonRequest(input)
	return out, req.Send()
}

// AddVpcFirewallAclRuleCommonWithContext is the same as AddVpcFirewallAclRuleCommon with the addition of
// the ability to pass a context and additional request options.
//
// See AddVpcFirewallAclRuleCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *FWCENTER) AddVpcFirewallAclRuleCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.AddVpcFirewallAclRuleCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opAddVpcFirewallAclRule = "AddVpcFirewallAclRule"

// AddVpcFirewallAclRuleRequest generates a "volcengine/request.Request" representing the
// client's request for the AddVpcFirewallAclRule operation. The "output" return
// value will be populated with the AddVpcFirewallAclRuleCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned AddVpcFirewallAclRuleCommon Request to send the API call to the service.
// the "output" return value is not valid until after AddVpcFirewallAclRuleCommon Send returns without error.
//
// See AddVpcFirewallAclRule for more information on using the AddVpcFirewallAclRule
// API call, and error handling.
//
//	// Example sending a request using the AddVpcFirewallAclRuleRequest method.
//	req, resp := client.AddVpcFirewallAclRuleRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *FWCENTER) AddVpcFirewallAclRuleRequest(input *AddVpcFirewallAclRuleInput) (req *request.Request, output *AddVpcFirewallAclRuleOutput) {
	op := &request.Operation{
		Name:       opAddVpcFirewallAclRule,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &AddVpcFirewallAclRuleInput{}
	}

	output = &AddVpcFirewallAclRuleOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// AddVpcFirewallAclRule API operation for FWCENTER.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for FWCENTER's
// API operation AddVpcFirewallAclRule for usage and error information.
func (c *FWCENTER) AddVpcFirewallAclRule(input *AddVpcFirewallAclRuleInput) (*AddVpcFirewallAclRuleOutput, error) {
	req, out := c.AddVpcFirewallAclRuleRequest(input)
	return out, req.Send()
}

// AddVpcFirewallAclRuleWithContext is the same as AddVpcFirewallAclRule with the addition of
// the ability to pass a context and additional request options.
//
// See AddVpcFirewallAclRule for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *FWCENTER) AddVpcFirewallAclRuleWithContext(ctx volcengine.Context, input *AddVpcFirewallAclRuleInput, opts ...request.Option) (*AddVpcFirewallAclRuleOutput, error) {
	req, out := c.AddVpcFirewallAclRuleRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type AddVpcFirewallAclRuleInput struct {
	_ struct{} `type:"structure"`

	// Action is a required field
	Action *string `type:"string" required:"true"`

	Description *string `type:"string"`

	DestPort *string `type:"string"`

	DestPortType *string `type:"string"`

	// Destination is a required field
	Destination *string `type:"string" required:"true"`

	// DestinationType is a required field
	DestinationType *string `type:"string" required:"true"`

	Prio *int32 `type:"int32"`

	// Proto is a required field
	Proto *string `type:"string" required:"true"`

	// Source is a required field
	Source *string `type:"string" required:"true"`

	// SourceType is a required field
	SourceType *string `type:"string" required:"true"`

	Status *bool `type:"boolean"`

	// VpcFirewallId is a required field
	VpcFirewallId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s AddVpcFirewallAclRuleInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s AddVpcFirewallAclRuleInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AddVpcFirewallAclRuleInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "AddVpcFirewallAclRuleInput"}
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
func (s *AddVpcFirewallAclRuleInput) SetAction(v string) *AddVpcFirewallAclRuleInput {
	s.Action = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *AddVpcFirewallAclRuleInput) SetDescription(v string) *AddVpcFirewallAclRuleInput {
	s.Description = &v
	return s
}

// SetDestPort sets the DestPort field's value.
func (s *AddVpcFirewallAclRuleInput) SetDestPort(v string) *AddVpcFirewallAclRuleInput {
	s.DestPort = &v
	return s
}

// SetDestPortType sets the DestPortType field's value.
func (s *AddVpcFirewallAclRuleInput) SetDestPortType(v string) *AddVpcFirewallAclRuleInput {
	s.DestPortType = &v
	return s
}

// SetDestination sets the Destination field's value.
func (s *AddVpcFirewallAclRuleInput) SetDestination(v string) *AddVpcFirewallAclRuleInput {
	s.Destination = &v
	return s
}

// SetDestinationType sets the DestinationType field's value.
func (s *AddVpcFirewallAclRuleInput) SetDestinationType(v string) *AddVpcFirewallAclRuleInput {
	s.DestinationType = &v
	return s
}

// SetPrio sets the Prio field's value.
func (s *AddVpcFirewallAclRuleInput) SetPrio(v int32) *AddVpcFirewallAclRuleInput {
	s.Prio = &v
	return s
}

// SetProto sets the Proto field's value.
func (s *AddVpcFirewallAclRuleInput) SetProto(v string) *AddVpcFirewallAclRuleInput {
	s.Proto = &v
	return s
}

// SetSource sets the Source field's value.
func (s *AddVpcFirewallAclRuleInput) SetSource(v string) *AddVpcFirewallAclRuleInput {
	s.Source = &v
	return s
}

// SetSourceType sets the SourceType field's value.
func (s *AddVpcFirewallAclRuleInput) SetSourceType(v string) *AddVpcFirewallAclRuleInput {
	s.SourceType = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *AddVpcFirewallAclRuleInput) SetStatus(v bool) *AddVpcFirewallAclRuleInput {
	s.Status = &v
	return s
}

// SetVpcFirewallId sets the VpcFirewallId field's value.
func (s *AddVpcFirewallAclRuleInput) SetVpcFirewallId(v string) *AddVpcFirewallAclRuleInput {
	s.VpcFirewallId = &v
	return s
}

type AddVpcFirewallAclRuleOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	RuleId *string `type:"string"`
}

// String returns the string representation
func (s AddVpcFirewallAclRuleOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s AddVpcFirewallAclRuleOutput) GoString() string {
	return s.String()
}

// SetRuleId sets the RuleId field's value.
func (s *AddVpcFirewallAclRuleOutput) SetRuleId(v string) *AddVpcFirewallAclRuleOutput {
	s.RuleId = &v
	return s
}
// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package fwcenter

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opAddNatFirewallControlPolicyCommon = "AddNatFirewallControlPolicy"

// AddNatFirewallControlPolicyCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the AddNatFirewallControlPolicyCommon operation. The "output" return
// value will be populated with the AddNatFirewallControlPolicyCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned AddNatFirewallControlPolicyCommon Request to send the API call to the service.
// the "output" return value is not valid until after AddNatFirewallControlPolicyCommon Send returns without error.
//
// See AddNatFirewallControlPolicyCommon for more information on using the AddNatFirewallControlPolicyCommon
// API call, and error handling.
//
//    // Example sending a request using the AddNatFirewallControlPolicyCommonRequest method.
//    req, resp := client.AddNatFirewallControlPolicyCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *FWCENTER) AddNatFirewallControlPolicyCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opAddNatFirewallControlPolicyCommon,
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

// AddNatFirewallControlPolicyCommon API operation for FWCENTER.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for FWCENTER's
// API operation AddNatFirewallControlPolicyCommon for usage and error information.
func (c *FWCENTER) AddNatFirewallControlPolicyCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.AddNatFirewallControlPolicyCommonRequest(input)
	return out, req.Send()
}

// AddNatFirewallControlPolicyCommonWithContext is the same as AddNatFirewallControlPolicyCommon with the addition of
// the ability to pass a context and additional request options.
//
// See AddNatFirewallControlPolicyCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *FWCENTER) AddNatFirewallControlPolicyCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.AddNatFirewallControlPolicyCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opAddNatFirewallControlPolicy = "AddNatFirewallControlPolicy"

// AddNatFirewallControlPolicyRequest generates a "volcengine/request.Request" representing the
// client's request for the AddNatFirewallControlPolicy operation. The "output" return
// value will be populated with the AddNatFirewallControlPolicyCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned AddNatFirewallControlPolicyCommon Request to send the API call to the service.
// the "output" return value is not valid until after AddNatFirewallControlPolicyCommon Send returns without error.
//
// See AddNatFirewallControlPolicy for more information on using the AddNatFirewallControlPolicy
// API call, and error handling.
//
//    // Example sending a request using the AddNatFirewallControlPolicyRequest method.
//    req, resp := client.AddNatFirewallControlPolicyRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *FWCENTER) AddNatFirewallControlPolicyRequest(input *AddNatFirewallControlPolicyInput) (req *request.Request, output *AddNatFirewallControlPolicyOutput) {
	op := &request.Operation{
		Name:       opAddNatFirewallControlPolicy,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &AddNatFirewallControlPolicyInput{}
	}

	output = &AddNatFirewallControlPolicyOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// AddNatFirewallControlPolicy API operation for FWCENTER.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for FWCENTER's
// API operation AddNatFirewallControlPolicy for usage and error information.
func (c *FWCENTER) AddNatFirewallControlPolicy(input *AddNatFirewallControlPolicyInput) (*AddNatFirewallControlPolicyOutput, error) {
	req, out := c.AddNatFirewallControlPolicyRequest(input)
	return out, req.Send()
}

// AddNatFirewallControlPolicyWithContext is the same as AddNatFirewallControlPolicy with the addition of
// the ability to pass a context and additional request options.
//
// See AddNatFirewallControlPolicy for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *FWCENTER) AddNatFirewallControlPolicyWithContext(ctx volcengine.Context, input *AddNatFirewallControlPolicyInput, opts ...request.Option) (*AddNatFirewallControlPolicyOutput, error) {
	req, out := c.AddNatFirewallControlPolicyRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type AddNatFirewallControlPolicyInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	// Action is a required field
	Action *string `type:"string" json:",omitempty" required:"true" enum:"EnumOfActionForAddNatFirewallControlPolicyInput"`

	Description *string `type:"string" json:",omitempty"`

	DestPort *string `type:"string" json:",omitempty"`

	DestPortType *string `type:"string" json:",omitempty" enum:"EnumOfDestPortTypeForAddNatFirewallControlPolicyInput"`

	// Destination is a required field
	Destination *string `type:"string" json:",omitempty" required:"true"`

	// DestinationType is a required field
	DestinationType *string `type:"string" json:",omitempty" required:"true" enum:"EnumOfDestinationTypeForAddNatFirewallControlPolicyInput"`

	// Direction is a required field
	Direction *string `type:"string" json:",omitempty" required:"true" enum:"EnumOfDirectionForAddNatFirewallControlPolicyInput"`

	EndTime *int32 `type:"int32" json:",omitempty"`

	// NatFirewallId is a required field
	NatFirewallId *string `type:"string" json:",omitempty" required:"true"`

	Prio *int32 `type:"int32" json:",omitempty"`

	// Proto is a required field
	Proto *string `type:"string" json:",omitempty" required:"true" enum:"EnumOfProtoForAddNatFirewallControlPolicyInput"`

	RepeatDays []*int32 `type:"list" json:",omitempty"`

	RepeatEndTime *string `type:"string" json:",omitempty"`

	RepeatStartTime *string `type:"string" json:",omitempty"`

	RepeatType *string `type:"string" json:",omitempty" enum:"EnumOfRepeatTypeForAddNatFirewallControlPolicyInput"`

	// Source is a required field
	Source *string `type:"string" json:",omitempty" required:"true"`

	// SourceType is a required field
	SourceType *string `type:"string" json:",omitempty" required:"true" enum:"EnumOfSourceTypeForAddNatFirewallControlPolicyInput"`

	StartTime *int32 `type:"int32" json:",omitempty"`

	Status *bool `type:"boolean" json:",omitempty"`
}

// String returns the string representation
func (s AddNatFirewallControlPolicyInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s AddNatFirewallControlPolicyInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AddNatFirewallControlPolicyInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "AddNatFirewallControlPolicyInput"}
	if s.Action == nil {
		invalidParams.Add(request.NewErrParamRequired("Action"))
	}
	if s.Destination == nil {
		invalidParams.Add(request.NewErrParamRequired("Destination"))
	}
	if s.DestinationType == nil {
		invalidParams.Add(request.NewErrParamRequired("DestinationType"))
	}
	if s.Direction == nil {
		invalidParams.Add(request.NewErrParamRequired("Direction"))
	}
	if s.NatFirewallId == nil {
		invalidParams.Add(request.NewErrParamRequired("NatFirewallId"))
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

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetAction sets the Action field's value.
func (s *AddNatFirewallControlPolicyInput) SetAction(v string) *AddNatFirewallControlPolicyInput {
	s.Action = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *AddNatFirewallControlPolicyInput) SetDescription(v string) *AddNatFirewallControlPolicyInput {
	s.Description = &v
	return s
}

// SetDestPort sets the DestPort field's value.
func (s *AddNatFirewallControlPolicyInput) SetDestPort(v string) *AddNatFirewallControlPolicyInput {
	s.DestPort = &v
	return s
}

// SetDestPortType sets the DestPortType field's value.
func (s *AddNatFirewallControlPolicyInput) SetDestPortType(v string) *AddNatFirewallControlPolicyInput {
	s.DestPortType = &v
	return s
}

// SetDestination sets the Destination field's value.
func (s *AddNatFirewallControlPolicyInput) SetDestination(v string) *AddNatFirewallControlPolicyInput {
	s.Destination = &v
	return s
}

// SetDestinationType sets the DestinationType field's value.
func (s *AddNatFirewallControlPolicyInput) SetDestinationType(v string) *AddNatFirewallControlPolicyInput {
	s.DestinationType = &v
	return s
}

// SetDirection sets the Direction field's value.
func (s *AddNatFirewallControlPolicyInput) SetDirection(v string) *AddNatFirewallControlPolicyInput {
	s.Direction = &v
	return s
}

// SetEndTime sets the EndTime field's value.
func (s *AddNatFirewallControlPolicyInput) SetEndTime(v int32) *AddNatFirewallControlPolicyInput {
	s.EndTime = &v
	return s
}

// SetNatFirewallId sets the NatFirewallId field's value.
func (s *AddNatFirewallControlPolicyInput) SetNatFirewallId(v string) *AddNatFirewallControlPolicyInput {
	s.NatFirewallId = &v
	return s
}

// SetPrio sets the Prio field's value.
func (s *AddNatFirewallControlPolicyInput) SetPrio(v int32) *AddNatFirewallControlPolicyInput {
	s.Prio = &v
	return s
}

// SetProto sets the Proto field's value.
func (s *AddNatFirewallControlPolicyInput) SetProto(v string) *AddNatFirewallControlPolicyInput {
	s.Proto = &v
	return s
}

// SetRepeatDays sets the RepeatDays field's value.
func (s *AddNatFirewallControlPolicyInput) SetRepeatDays(v []*int32) *AddNatFirewallControlPolicyInput {
	s.RepeatDays = v
	return s
}

// SetRepeatEndTime sets the RepeatEndTime field's value.
func (s *AddNatFirewallControlPolicyInput) SetRepeatEndTime(v string) *AddNatFirewallControlPolicyInput {
	s.RepeatEndTime = &v
	return s
}

// SetRepeatStartTime sets the RepeatStartTime field's value.
func (s *AddNatFirewallControlPolicyInput) SetRepeatStartTime(v string) *AddNatFirewallControlPolicyInput {
	s.RepeatStartTime = &v
	return s
}

// SetRepeatType sets the RepeatType field's value.
func (s *AddNatFirewallControlPolicyInput) SetRepeatType(v string) *AddNatFirewallControlPolicyInput {
	s.RepeatType = &v
	return s
}

// SetSource sets the Source field's value.
func (s *AddNatFirewallControlPolicyInput) SetSource(v string) *AddNatFirewallControlPolicyInput {
	s.Source = &v
	return s
}

// SetSourceType sets the SourceType field's value.
func (s *AddNatFirewallControlPolicyInput) SetSourceType(v string) *AddNatFirewallControlPolicyInput {
	s.SourceType = &v
	return s
}

// SetStartTime sets the StartTime field's value.
func (s *AddNatFirewallControlPolicyInput) SetStartTime(v int32) *AddNatFirewallControlPolicyInput {
	s.StartTime = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *AddNatFirewallControlPolicyInput) SetStatus(v bool) *AddNatFirewallControlPolicyInput {
	s.Status = &v
	return s
}

type AddNatFirewallControlPolicyOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	RuleId *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s AddNatFirewallControlPolicyOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s AddNatFirewallControlPolicyOutput) GoString() string {
	return s.String()
}

// SetRuleId sets the RuleId field's value.
func (s *AddNatFirewallControlPolicyOutput) SetRuleId(v string) *AddNatFirewallControlPolicyOutput {
	s.RuleId = &v
	return s
}

const (
	// EnumOfActionForAddNatFirewallControlPolicyInputAccept is a EnumOfActionForAddNatFirewallControlPolicyInput enum value
	EnumOfActionForAddNatFirewallControlPolicyInputAccept = "accept"

	// EnumOfActionForAddNatFirewallControlPolicyInputDeny is a EnumOfActionForAddNatFirewallControlPolicyInput enum value
	EnumOfActionForAddNatFirewallControlPolicyInputDeny = "deny"

	// EnumOfActionForAddNatFirewallControlPolicyInputMonitor is a EnumOfActionForAddNatFirewallControlPolicyInput enum value
	EnumOfActionForAddNatFirewallControlPolicyInputMonitor = "monitor"
)

const (
	// EnumOfDestPortTypeForAddNatFirewallControlPolicyInputPort is a EnumOfDestPortTypeForAddNatFirewallControlPolicyInput enum value
	EnumOfDestPortTypeForAddNatFirewallControlPolicyInputPort = "port"

	// EnumOfDestPortTypeForAddNatFirewallControlPolicyInputGroup is a EnumOfDestPortTypeForAddNatFirewallControlPolicyInput enum value
	EnumOfDestPortTypeForAddNatFirewallControlPolicyInputGroup = "group"
)

const (
	// EnumOfDestinationTypeForAddNatFirewallControlPolicyInputNet is a EnumOfDestinationTypeForAddNatFirewallControlPolicyInput enum value
	EnumOfDestinationTypeForAddNatFirewallControlPolicyInputNet = "net"

	// EnumOfDestinationTypeForAddNatFirewallControlPolicyInputLocation is a EnumOfDestinationTypeForAddNatFirewallControlPolicyInput enum value
	EnumOfDestinationTypeForAddNatFirewallControlPolicyInputLocation = "location"

	// EnumOfDestinationTypeForAddNatFirewallControlPolicyInputGroup is a EnumOfDestinationTypeForAddNatFirewallControlPolicyInput enum value
	EnumOfDestinationTypeForAddNatFirewallControlPolicyInputGroup = "group"

	// EnumOfDestinationTypeForAddNatFirewallControlPolicyInputDomain is a EnumOfDestinationTypeForAddNatFirewallControlPolicyInput enum value
	EnumOfDestinationTypeForAddNatFirewallControlPolicyInputDomain = "domain"
)

const (
	// EnumOfDirectionForAddNatFirewallControlPolicyInputIn is a EnumOfDirectionForAddNatFirewallControlPolicyInput enum value
	EnumOfDirectionForAddNatFirewallControlPolicyInputIn = "in"

	// EnumOfDirectionForAddNatFirewallControlPolicyInputOut is a EnumOfDirectionForAddNatFirewallControlPolicyInput enum value
	EnumOfDirectionForAddNatFirewallControlPolicyInputOut = "out"
)

const (
	// EnumOfProtoForAddNatFirewallControlPolicyInputIcmp is a EnumOfProtoForAddNatFirewallControlPolicyInput enum value
	EnumOfProtoForAddNatFirewallControlPolicyInputIcmp = "ICMP"

	// EnumOfProtoForAddNatFirewallControlPolicyInputTcp is a EnumOfProtoForAddNatFirewallControlPolicyInput enum value
	EnumOfProtoForAddNatFirewallControlPolicyInputTcp = "TCP"

	// EnumOfProtoForAddNatFirewallControlPolicyInputUdp is a EnumOfProtoForAddNatFirewallControlPolicyInput enum value
	EnumOfProtoForAddNatFirewallControlPolicyInputUdp = "UDP"

	// EnumOfProtoForAddNatFirewallControlPolicyInputAny is a EnumOfProtoForAddNatFirewallControlPolicyInput enum value
	EnumOfProtoForAddNatFirewallControlPolicyInputAny = "ANY"
)

const (
	// EnumOfRepeatTypeForAddNatFirewallControlPolicyInputPermanent is a EnumOfRepeatTypeForAddNatFirewallControlPolicyInput enum value
	EnumOfRepeatTypeForAddNatFirewallControlPolicyInputPermanent = "Permanent"

	// EnumOfRepeatTypeForAddNatFirewallControlPolicyInputOnce is a EnumOfRepeatTypeForAddNatFirewallControlPolicyInput enum value
	EnumOfRepeatTypeForAddNatFirewallControlPolicyInputOnce = "Once"

	// EnumOfRepeatTypeForAddNatFirewallControlPolicyInputDaily is a EnumOfRepeatTypeForAddNatFirewallControlPolicyInput enum value
	EnumOfRepeatTypeForAddNatFirewallControlPolicyInputDaily = "Daily"

	// EnumOfRepeatTypeForAddNatFirewallControlPolicyInputWeekly is a EnumOfRepeatTypeForAddNatFirewallControlPolicyInput enum value
	EnumOfRepeatTypeForAddNatFirewallControlPolicyInputWeekly = "Weekly"

	// EnumOfRepeatTypeForAddNatFirewallControlPolicyInputMonthly is a EnumOfRepeatTypeForAddNatFirewallControlPolicyInput enum value
	EnumOfRepeatTypeForAddNatFirewallControlPolicyInputMonthly = "Monthly"
)

const (
	// EnumOfSourceTypeForAddNatFirewallControlPolicyInputNet is a EnumOfSourceTypeForAddNatFirewallControlPolicyInput enum value
	EnumOfSourceTypeForAddNatFirewallControlPolicyInputNet = "net"

	// EnumOfSourceTypeForAddNatFirewallControlPolicyInputLocation is a EnumOfSourceTypeForAddNatFirewallControlPolicyInput enum value
	EnumOfSourceTypeForAddNatFirewallControlPolicyInputLocation = "location"

	// EnumOfSourceTypeForAddNatFirewallControlPolicyInputGroup is a EnumOfSourceTypeForAddNatFirewallControlPolicyInput enum value
	EnumOfSourceTypeForAddNatFirewallControlPolicyInputGroup = "group"
)

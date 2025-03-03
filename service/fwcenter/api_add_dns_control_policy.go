// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package fwcenter

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opAddDnsControlPolicyCommon = "AddDnsControlPolicy"

// AddDnsControlPolicyCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the AddDnsControlPolicyCommon operation. The "output" return
// value will be populated with the AddDnsControlPolicyCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned AddDnsControlPolicyCommon Request to send the API call to the service.
// the "output" return value is not valid until after AddDnsControlPolicyCommon Send returns without error.
//
// See AddDnsControlPolicyCommon for more information on using the AddDnsControlPolicyCommon
// API call, and error handling.
//
//    // Example sending a request using the AddDnsControlPolicyCommonRequest method.
//    req, resp := client.AddDnsControlPolicyCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *FWCENTER) AddDnsControlPolicyCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opAddDnsControlPolicyCommon,
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

// AddDnsControlPolicyCommon API operation for FWCENTER.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for FWCENTER's
// API operation AddDnsControlPolicyCommon for usage and error information.
func (c *FWCENTER) AddDnsControlPolicyCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.AddDnsControlPolicyCommonRequest(input)
	return out, req.Send()
}

// AddDnsControlPolicyCommonWithContext is the same as AddDnsControlPolicyCommon with the addition of
// the ability to pass a context and additional request options.
//
// See AddDnsControlPolicyCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *FWCENTER) AddDnsControlPolicyCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.AddDnsControlPolicyCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opAddDnsControlPolicy = "AddDnsControlPolicy"

// AddDnsControlPolicyRequest generates a "volcengine/request.Request" representing the
// client's request for the AddDnsControlPolicy operation. The "output" return
// value will be populated with the AddDnsControlPolicyCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned AddDnsControlPolicyCommon Request to send the API call to the service.
// the "output" return value is not valid until after AddDnsControlPolicyCommon Send returns without error.
//
// See AddDnsControlPolicy for more information on using the AddDnsControlPolicy
// API call, and error handling.
//
//    // Example sending a request using the AddDnsControlPolicyRequest method.
//    req, resp := client.AddDnsControlPolicyRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *FWCENTER) AddDnsControlPolicyRequest(input *AddDnsControlPolicyInput) (req *request.Request, output *AddDnsControlPolicyOutput) {
	op := &request.Operation{
		Name:       opAddDnsControlPolicy,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &AddDnsControlPolicyInput{}
	}

	output = &AddDnsControlPolicyOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// AddDnsControlPolicy API operation for FWCENTER.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for FWCENTER's
// API operation AddDnsControlPolicy for usage and error information.
func (c *FWCENTER) AddDnsControlPolicy(input *AddDnsControlPolicyInput) (*AddDnsControlPolicyOutput, error) {
	req, out := c.AddDnsControlPolicyRequest(input)
	return out, req.Send()
}

// AddDnsControlPolicyWithContext is the same as AddDnsControlPolicy with the addition of
// the ability to pass a context and additional request options.
//
// See AddDnsControlPolicy for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *FWCENTER) AddDnsControlPolicyWithContext(ctx volcengine.Context, input *AddDnsControlPolicyInput, opts ...request.Option) (*AddDnsControlPolicyOutput, error) {
	req, out := c.AddDnsControlPolicyRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type AddDnsControlPolicyInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Description *string `max:"100" type:"string" json:",omitempty"`

	// Destination is a required field
	Destination *string `type:"string" json:",omitempty" required:"true"`

	// DestinationType is a required field
	DestinationType *string `type:"string" json:",omitempty" required:"true" enum:"EnumOfDestinationTypeForAddDnsControlPolicyInput"`

	Source []*SourceForAddDnsControlPolicyInput `type:"list" json:",omitempty"`
}

// String returns the string representation
func (s AddDnsControlPolicyInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s AddDnsControlPolicyInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AddDnsControlPolicyInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "AddDnsControlPolicyInput"}
	if s.Description != nil && len(*s.Description) > 100 {
		invalidParams.Add(request.NewErrParamMaxLen("Description", 100, *s.Description))
	}
	if s.Destination == nil {
		invalidParams.Add(request.NewErrParamRequired("Destination"))
	}
	if s.DestinationType == nil {
		invalidParams.Add(request.NewErrParamRequired("DestinationType"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetDescription sets the Description field's value.
func (s *AddDnsControlPolicyInput) SetDescription(v string) *AddDnsControlPolicyInput {
	s.Description = &v
	return s
}

// SetDestination sets the Destination field's value.
func (s *AddDnsControlPolicyInput) SetDestination(v string) *AddDnsControlPolicyInput {
	s.Destination = &v
	return s
}

// SetDestinationType sets the DestinationType field's value.
func (s *AddDnsControlPolicyInput) SetDestinationType(v string) *AddDnsControlPolicyInput {
	s.DestinationType = &v
	return s
}

// SetSource sets the Source field's value.
func (s *AddDnsControlPolicyInput) SetSource(v []*SourceForAddDnsControlPolicyInput) *AddDnsControlPolicyInput {
	s.Source = v
	return s
}

type AddDnsControlPolicyOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	RuleId *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s AddDnsControlPolicyOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s AddDnsControlPolicyOutput) GoString() string {
	return s.String()
}

// SetRuleId sets the RuleId field's value.
func (s *AddDnsControlPolicyOutput) SetRuleId(v string) *AddDnsControlPolicyOutput {
	s.RuleId = &v
	return s
}

type SourceForAddDnsControlPolicyInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Region *string `type:"string" json:",omitempty"`

	VpcId *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s SourceForAddDnsControlPolicyInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s SourceForAddDnsControlPolicyInput) GoString() string {
	return s.String()
}

// SetRegion sets the Region field's value.
func (s *SourceForAddDnsControlPolicyInput) SetRegion(v string) *SourceForAddDnsControlPolicyInput {
	s.Region = &v
	return s
}

// SetVpcId sets the VpcId field's value.
func (s *SourceForAddDnsControlPolicyInput) SetVpcId(v string) *SourceForAddDnsControlPolicyInput {
	s.VpcId = &v
	return s
}

const (
	// EnumOfDestinationTypeForAddDnsControlPolicyInputGroup is a EnumOfDestinationTypeForAddDnsControlPolicyInput enum value
	EnumOfDestinationTypeForAddDnsControlPolicyInputGroup = "group"

	// EnumOfDestinationTypeForAddDnsControlPolicyInputDomain is a EnumOfDestinationTypeForAddDnsControlPolicyInput enum value
	EnumOfDestinationTypeForAddDnsControlPolicyInputDomain = "domain"
)

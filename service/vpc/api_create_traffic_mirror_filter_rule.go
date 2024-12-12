// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package vpc

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opCreateTrafficMirrorFilterRuleCommon = "CreateTrafficMirrorFilterRule"

// CreateTrafficMirrorFilterRuleCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateTrafficMirrorFilterRuleCommon operation. The "output" return
// value will be populated with the CreateTrafficMirrorFilterRuleCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateTrafficMirrorFilterRuleCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateTrafficMirrorFilterRuleCommon Send returns without error.
//
// See CreateTrafficMirrorFilterRuleCommon for more information on using the CreateTrafficMirrorFilterRuleCommon
// API call, and error handling.
//
//    // Example sending a request using the CreateTrafficMirrorFilterRuleCommonRequest method.
//    req, resp := client.CreateTrafficMirrorFilterRuleCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VPC) CreateTrafficMirrorFilterRuleCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opCreateTrafficMirrorFilterRuleCommon,
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

// CreateTrafficMirrorFilterRuleCommon API operation for VPC.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VPC's
// API operation CreateTrafficMirrorFilterRuleCommon for usage and error information.
func (c *VPC) CreateTrafficMirrorFilterRuleCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.CreateTrafficMirrorFilterRuleCommonRequest(input)
	return out, req.Send()
}

// CreateTrafficMirrorFilterRuleCommonWithContext is the same as CreateTrafficMirrorFilterRuleCommon with the addition of
// the ability to pass a context and additional request options.
//
// See CreateTrafficMirrorFilterRuleCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPC) CreateTrafficMirrorFilterRuleCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.CreateTrafficMirrorFilterRuleCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opCreateTrafficMirrorFilterRule = "CreateTrafficMirrorFilterRule"

// CreateTrafficMirrorFilterRuleRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateTrafficMirrorFilterRule operation. The "output" return
// value will be populated with the CreateTrafficMirrorFilterRuleCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateTrafficMirrorFilterRuleCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateTrafficMirrorFilterRuleCommon Send returns without error.
//
// See CreateTrafficMirrorFilterRule for more information on using the CreateTrafficMirrorFilterRule
// API call, and error handling.
//
//    // Example sending a request using the CreateTrafficMirrorFilterRuleRequest method.
//    req, resp := client.CreateTrafficMirrorFilterRuleRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VPC) CreateTrafficMirrorFilterRuleRequest(input *CreateTrafficMirrorFilterRuleInput) (req *request.Request, output *CreateTrafficMirrorFilterRuleOutput) {
	op := &request.Operation{
		Name:       opCreateTrafficMirrorFilterRule,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateTrafficMirrorFilterRuleInput{}
	}

	output = &CreateTrafficMirrorFilterRuleOutput{}
	req = c.newRequest(op, input, output)

	return
}

// CreateTrafficMirrorFilterRule API operation for VPC.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VPC's
// API operation CreateTrafficMirrorFilterRule for usage and error information.
func (c *VPC) CreateTrafficMirrorFilterRule(input *CreateTrafficMirrorFilterRuleInput) (*CreateTrafficMirrorFilterRuleOutput, error) {
	req, out := c.CreateTrafficMirrorFilterRuleRequest(input)
	return out, req.Send()
}

// CreateTrafficMirrorFilterRuleWithContext is the same as CreateTrafficMirrorFilterRule with the addition of
// the ability to pass a context and additional request options.
//
// See CreateTrafficMirrorFilterRule for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPC) CreateTrafficMirrorFilterRuleWithContext(ctx volcengine.Context, input *CreateTrafficMirrorFilterRuleInput, opts ...request.Option) (*CreateTrafficMirrorFilterRuleOutput, error) {
	req, out := c.CreateTrafficMirrorFilterRuleRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type CreateTrafficMirrorFilterRuleInput struct {
	_ struct{} `type:"structure"`

	ClientToken *string `type:"string"`

	Description *string `type:"string"`

	// DestinationCidrBlock is a required field
	DestinationCidrBlock *string `type:"string" required:"true"`

	DestinationPortRange *string `type:"string"`

	// Policy is a required field
	Policy *string `type:"string" required:"true" enum:"PolicyForCreateTrafficMirrorFilterRuleInput"`

	Priority *int64 `type:"integer"`

	// Protocol is a required field
	Protocol *string `type:"string" required:"true" enum:"ProtocolForCreateTrafficMirrorFilterRuleInput"`

	// SourceCidrBlock is a required field
	SourceCidrBlock *string `type:"string" required:"true"`

	SourcePortRange *string `type:"string"`

	// TrafficDirection is a required field
	TrafficDirection *string `type:"string" required:"true" enum:"TrafficDirectionForCreateTrafficMirrorFilterRuleInput"`

	// TrafficMirrorFilterId is a required field
	TrafficMirrorFilterId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s CreateTrafficMirrorFilterRuleInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateTrafficMirrorFilterRuleInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateTrafficMirrorFilterRuleInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "CreateTrafficMirrorFilterRuleInput"}
	if s.DestinationCidrBlock == nil {
		invalidParams.Add(request.NewErrParamRequired("DestinationCidrBlock"))
	}
	if s.Policy == nil {
		invalidParams.Add(request.NewErrParamRequired("Policy"))
	}
	if s.Protocol == nil {
		invalidParams.Add(request.NewErrParamRequired("Protocol"))
	}
	if s.SourceCidrBlock == nil {
		invalidParams.Add(request.NewErrParamRequired("SourceCidrBlock"))
	}
	if s.TrafficDirection == nil {
		invalidParams.Add(request.NewErrParamRequired("TrafficDirection"))
	}
	if s.TrafficMirrorFilterId == nil {
		invalidParams.Add(request.NewErrParamRequired("TrafficMirrorFilterId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetClientToken sets the ClientToken field's value.
func (s *CreateTrafficMirrorFilterRuleInput) SetClientToken(v string) *CreateTrafficMirrorFilterRuleInput {
	s.ClientToken = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *CreateTrafficMirrorFilterRuleInput) SetDescription(v string) *CreateTrafficMirrorFilterRuleInput {
	s.Description = &v
	return s
}

// SetDestinationCidrBlock sets the DestinationCidrBlock field's value.
func (s *CreateTrafficMirrorFilterRuleInput) SetDestinationCidrBlock(v string) *CreateTrafficMirrorFilterRuleInput {
	s.DestinationCidrBlock = &v
	return s
}

// SetDestinationPortRange sets the DestinationPortRange field's value.
func (s *CreateTrafficMirrorFilterRuleInput) SetDestinationPortRange(v string) *CreateTrafficMirrorFilterRuleInput {
	s.DestinationPortRange = &v
	return s
}

// SetPolicy sets the Policy field's value.
func (s *CreateTrafficMirrorFilterRuleInput) SetPolicy(v string) *CreateTrafficMirrorFilterRuleInput {
	s.Policy = &v
	return s
}

// SetPriority sets the Priority field's value.
func (s *CreateTrafficMirrorFilterRuleInput) SetPriority(v int64) *CreateTrafficMirrorFilterRuleInput {
	s.Priority = &v
	return s
}

// SetProtocol sets the Protocol field's value.
func (s *CreateTrafficMirrorFilterRuleInput) SetProtocol(v string) *CreateTrafficMirrorFilterRuleInput {
	s.Protocol = &v
	return s
}

// SetSourceCidrBlock sets the SourceCidrBlock field's value.
func (s *CreateTrafficMirrorFilterRuleInput) SetSourceCidrBlock(v string) *CreateTrafficMirrorFilterRuleInput {
	s.SourceCidrBlock = &v
	return s
}

// SetSourcePortRange sets the SourcePortRange field's value.
func (s *CreateTrafficMirrorFilterRuleInput) SetSourcePortRange(v string) *CreateTrafficMirrorFilterRuleInput {
	s.SourcePortRange = &v
	return s
}

// SetTrafficDirection sets the TrafficDirection field's value.
func (s *CreateTrafficMirrorFilterRuleInput) SetTrafficDirection(v string) *CreateTrafficMirrorFilterRuleInput {
	s.TrafficDirection = &v
	return s
}

// SetTrafficMirrorFilterId sets the TrafficMirrorFilterId field's value.
func (s *CreateTrafficMirrorFilterRuleInput) SetTrafficMirrorFilterId(v string) *CreateTrafficMirrorFilterRuleInput {
	s.TrafficMirrorFilterId = &v
	return s
}

type CreateTrafficMirrorFilterRuleOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	RequestId *string `type:"string"`

	TrafficMirrorFilterRuleId *string `type:"string"`
}

// String returns the string representation
func (s CreateTrafficMirrorFilterRuleOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateTrafficMirrorFilterRuleOutput) GoString() string {
	return s.String()
}

// SetRequestId sets the RequestId field's value.
func (s *CreateTrafficMirrorFilterRuleOutput) SetRequestId(v string) *CreateTrafficMirrorFilterRuleOutput {
	s.RequestId = &v
	return s
}

// SetTrafficMirrorFilterRuleId sets the TrafficMirrorFilterRuleId field's value.
func (s *CreateTrafficMirrorFilterRuleOutput) SetTrafficMirrorFilterRuleId(v string) *CreateTrafficMirrorFilterRuleOutput {
	s.TrafficMirrorFilterRuleId = &v
	return s
}

const (
	// PolicyForCreateTrafficMirrorFilterRuleInputAccept is a PolicyForCreateTrafficMirrorFilterRuleInput enum value
	PolicyForCreateTrafficMirrorFilterRuleInputAccept = "accept"

	// PolicyForCreateTrafficMirrorFilterRuleInputReject is a PolicyForCreateTrafficMirrorFilterRuleInput enum value
	PolicyForCreateTrafficMirrorFilterRuleInputReject = "reject"
)

const (
	// ProtocolForCreateTrafficMirrorFilterRuleInputTcp is a ProtocolForCreateTrafficMirrorFilterRuleInput enum value
	ProtocolForCreateTrafficMirrorFilterRuleInputTcp = "tcp"

	// ProtocolForCreateTrafficMirrorFilterRuleInputUdp is a ProtocolForCreateTrafficMirrorFilterRuleInput enum value
	ProtocolForCreateTrafficMirrorFilterRuleInputUdp = "udp"

	// ProtocolForCreateTrafficMirrorFilterRuleInputIcmp is a ProtocolForCreateTrafficMirrorFilterRuleInput enum value
	ProtocolForCreateTrafficMirrorFilterRuleInputIcmp = "icmp"

	// ProtocolForCreateTrafficMirrorFilterRuleInputAll is a ProtocolForCreateTrafficMirrorFilterRuleInput enum value
	ProtocolForCreateTrafficMirrorFilterRuleInputAll = "all"
)

const (
	// TrafficDirectionForCreateTrafficMirrorFilterRuleInputEgress is a TrafficDirectionForCreateTrafficMirrorFilterRuleInput enum value
	TrafficDirectionForCreateTrafficMirrorFilterRuleInputEgress = "egress"

	// TrafficDirectionForCreateTrafficMirrorFilterRuleInputIngress is a TrafficDirectionForCreateTrafficMirrorFilterRuleInput enum value
	TrafficDirectionForCreateTrafficMirrorFilterRuleInputIngress = "ingress"
)

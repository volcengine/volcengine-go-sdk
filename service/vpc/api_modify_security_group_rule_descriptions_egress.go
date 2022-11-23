// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package vpc

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opModifySecurityGroupRuleDescriptionsEgressCommon = "ModifySecurityGroupRuleDescriptionsEgress"

// ModifySecurityGroupRuleDescriptionsEgressCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ModifySecurityGroupRuleDescriptionsEgressCommon operation. The "output" return
// value will be populated with the ModifySecurityGroupRuleDescriptionsEgressCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifySecurityGroupRuleDescriptionsEgressCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifySecurityGroupRuleDescriptionsEgressCommon Send returns without error.
//
// See ModifySecurityGroupRuleDescriptionsEgressCommon for more information on using the ModifySecurityGroupRuleDescriptionsEgressCommon
// API call, and error handling.
//
//    // Example sending a request using the ModifySecurityGroupRuleDescriptionsEgressCommonRequest method.
//    req, resp := client.ModifySecurityGroupRuleDescriptionsEgressCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VPC) ModifySecurityGroupRuleDescriptionsEgressCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opModifySecurityGroupRuleDescriptionsEgressCommon,
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

// ModifySecurityGroupRuleDescriptionsEgressCommon API operation for VPC.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VPC's
// API operation ModifySecurityGroupRuleDescriptionsEgressCommon for usage and error information.
func (c *VPC) ModifySecurityGroupRuleDescriptionsEgressCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ModifySecurityGroupRuleDescriptionsEgressCommonRequest(input)
	return out, req.Send()
}

// ModifySecurityGroupRuleDescriptionsEgressCommonWithContext is the same as ModifySecurityGroupRuleDescriptionsEgressCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ModifySecurityGroupRuleDescriptionsEgressCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPC) ModifySecurityGroupRuleDescriptionsEgressCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ModifySecurityGroupRuleDescriptionsEgressCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opModifySecurityGroupRuleDescriptionsEgress = "ModifySecurityGroupRuleDescriptionsEgress"

// ModifySecurityGroupRuleDescriptionsEgressRequest generates a "volcengine/request.Request" representing the
// client's request for the ModifySecurityGroupRuleDescriptionsEgress operation. The "output" return
// value will be populated with the ModifySecurityGroupRuleDescriptionsEgressCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifySecurityGroupRuleDescriptionsEgressCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifySecurityGroupRuleDescriptionsEgressCommon Send returns without error.
//
// See ModifySecurityGroupRuleDescriptionsEgress for more information on using the ModifySecurityGroupRuleDescriptionsEgress
// API call, and error handling.
//
//    // Example sending a request using the ModifySecurityGroupRuleDescriptionsEgressRequest method.
//    req, resp := client.ModifySecurityGroupRuleDescriptionsEgressRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VPC) ModifySecurityGroupRuleDescriptionsEgressRequest(input *ModifySecurityGroupRuleDescriptionsEgressInput) (req *request.Request, output *ModifySecurityGroupRuleDescriptionsEgressOutput) {
	op := &request.Operation{
		Name:       opModifySecurityGroupRuleDescriptionsEgress,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ModifySecurityGroupRuleDescriptionsEgressInput{}
	}

	output = &ModifySecurityGroupRuleDescriptionsEgressOutput{}
	req = c.newRequest(op, input, output)

	return
}

// ModifySecurityGroupRuleDescriptionsEgress API operation for VPC.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VPC's
// API operation ModifySecurityGroupRuleDescriptionsEgress for usage and error information.
func (c *VPC) ModifySecurityGroupRuleDescriptionsEgress(input *ModifySecurityGroupRuleDescriptionsEgressInput) (*ModifySecurityGroupRuleDescriptionsEgressOutput, error) {
	req, out := c.ModifySecurityGroupRuleDescriptionsEgressRequest(input)
	return out, req.Send()
}

// ModifySecurityGroupRuleDescriptionsEgressWithContext is the same as ModifySecurityGroupRuleDescriptionsEgress with the addition of
// the ability to pass a context and additional request options.
//
// See ModifySecurityGroupRuleDescriptionsEgress for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPC) ModifySecurityGroupRuleDescriptionsEgressWithContext(ctx volcengine.Context, input *ModifySecurityGroupRuleDescriptionsEgressInput, opts ...request.Option) (*ModifySecurityGroupRuleDescriptionsEgressOutput, error) {
	req, out := c.ModifySecurityGroupRuleDescriptionsEgressRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ModifySecurityGroupRuleDescriptionsEgressInput struct {
	_ struct{} `type:"structure"`

	CidrIp *string `type:"string"`

	Description *string `type:"string"`

	Policy *string `type:"string"`

	// PortEnd is a required field
	PortEnd *int64 `type:"integer" required:"true"`

	// PortStart is a required field
	PortStart *int64 `type:"integer" required:"true"`

	Priority *int64 `type:"integer"`

	// Protocol is a required field
	Protocol *string `type:"string" required:"true"`

	// SecurityGroupId is a required field
	SecurityGroupId *string `type:"string" required:"true"`

	SourceGroupId *string `type:"string"`
}

// String returns the string representation
func (s ModifySecurityGroupRuleDescriptionsEgressInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifySecurityGroupRuleDescriptionsEgressInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ModifySecurityGroupRuleDescriptionsEgressInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ModifySecurityGroupRuleDescriptionsEgressInput"}
	if s.PortEnd == nil {
		invalidParams.Add(request.NewErrParamRequired("PortEnd"))
	}
	if s.PortStart == nil {
		invalidParams.Add(request.NewErrParamRequired("PortStart"))
	}
	if s.Protocol == nil {
		invalidParams.Add(request.NewErrParamRequired("Protocol"))
	}
	if s.SecurityGroupId == nil {
		invalidParams.Add(request.NewErrParamRequired("SecurityGroupId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetCidrIp sets the CidrIp field's value.
func (s *ModifySecurityGroupRuleDescriptionsEgressInput) SetCidrIp(v string) *ModifySecurityGroupRuleDescriptionsEgressInput {
	s.CidrIp = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *ModifySecurityGroupRuleDescriptionsEgressInput) SetDescription(v string) *ModifySecurityGroupRuleDescriptionsEgressInput {
	s.Description = &v
	return s
}

// SetPolicy sets the Policy field's value.
func (s *ModifySecurityGroupRuleDescriptionsEgressInput) SetPolicy(v string) *ModifySecurityGroupRuleDescriptionsEgressInput {
	s.Policy = &v
	return s
}

// SetPortEnd sets the PortEnd field's value.
func (s *ModifySecurityGroupRuleDescriptionsEgressInput) SetPortEnd(v int64) *ModifySecurityGroupRuleDescriptionsEgressInput {
	s.PortEnd = &v
	return s
}

// SetPortStart sets the PortStart field's value.
func (s *ModifySecurityGroupRuleDescriptionsEgressInput) SetPortStart(v int64) *ModifySecurityGroupRuleDescriptionsEgressInput {
	s.PortStart = &v
	return s
}

// SetPriority sets the Priority field's value.
func (s *ModifySecurityGroupRuleDescriptionsEgressInput) SetPriority(v int64) *ModifySecurityGroupRuleDescriptionsEgressInput {
	s.Priority = &v
	return s
}

// SetProtocol sets the Protocol field's value.
func (s *ModifySecurityGroupRuleDescriptionsEgressInput) SetProtocol(v string) *ModifySecurityGroupRuleDescriptionsEgressInput {
	s.Protocol = &v
	return s
}

// SetSecurityGroupId sets the SecurityGroupId field's value.
func (s *ModifySecurityGroupRuleDescriptionsEgressInput) SetSecurityGroupId(v string) *ModifySecurityGroupRuleDescriptionsEgressInput {
	s.SecurityGroupId = &v
	return s
}

// SetSourceGroupId sets the SourceGroupId field's value.
func (s *ModifySecurityGroupRuleDescriptionsEgressInput) SetSourceGroupId(v string) *ModifySecurityGroupRuleDescriptionsEgressInput {
	s.SourceGroupId = &v
	return s
}

type ModifySecurityGroupRuleDescriptionsEgressOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	RequestId *string `type:"string"`
}

// String returns the string representation
func (s ModifySecurityGroupRuleDescriptionsEgressOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifySecurityGroupRuleDescriptionsEgressOutput) GoString() string {
	return s.String()
}

// SetRequestId sets the RequestId field's value.
func (s *ModifySecurityGroupRuleDescriptionsEgressOutput) SetRequestId(v string) *ModifySecurityGroupRuleDescriptionsEgressOutput {
	s.RequestId = &v
	return s
}

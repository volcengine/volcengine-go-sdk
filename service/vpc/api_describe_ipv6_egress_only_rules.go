// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package vpc

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribeIpv6EgressOnlyRulesCommon = "DescribeIpv6EgressOnlyRules"

// DescribeIpv6EgressOnlyRulesCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeIpv6EgressOnlyRulesCommon operation. The "output" return
// value will be populated with the DescribeIpv6EgressOnlyRulesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeIpv6EgressOnlyRulesCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeIpv6EgressOnlyRulesCommon Send returns without error.
//
// See DescribeIpv6EgressOnlyRulesCommon for more information on using the DescribeIpv6EgressOnlyRulesCommon
// API call, and error handling.
//
//    // Example sending a request using the DescribeIpv6EgressOnlyRulesCommonRequest method.
//    req, resp := client.DescribeIpv6EgressOnlyRulesCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VPC) DescribeIpv6EgressOnlyRulesCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeIpv6EgressOnlyRulesCommon,
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

// DescribeIpv6EgressOnlyRulesCommon API operation for VPC.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VPC's
// API operation DescribeIpv6EgressOnlyRulesCommon for usage and error information.
func (c *VPC) DescribeIpv6EgressOnlyRulesCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeIpv6EgressOnlyRulesCommonRequest(input)
	return out, req.Send()
}

// DescribeIpv6EgressOnlyRulesCommonWithContext is the same as DescribeIpv6EgressOnlyRulesCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeIpv6EgressOnlyRulesCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPC) DescribeIpv6EgressOnlyRulesCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeIpv6EgressOnlyRulesCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeIpv6EgressOnlyRules = "DescribeIpv6EgressOnlyRules"

// DescribeIpv6EgressOnlyRulesRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeIpv6EgressOnlyRules operation. The "output" return
// value will be populated with the DescribeIpv6EgressOnlyRulesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeIpv6EgressOnlyRulesCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeIpv6EgressOnlyRulesCommon Send returns without error.
//
// See DescribeIpv6EgressOnlyRules for more information on using the DescribeIpv6EgressOnlyRules
// API call, and error handling.
//
//    // Example sending a request using the DescribeIpv6EgressOnlyRulesRequest method.
//    req, resp := client.DescribeIpv6EgressOnlyRulesRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VPC) DescribeIpv6EgressOnlyRulesRequest(input *DescribeIpv6EgressOnlyRulesInput) (req *request.Request, output *DescribeIpv6EgressOnlyRulesOutput) {
	op := &request.Operation{
		Name:       opDescribeIpv6EgressOnlyRules,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeIpv6EgressOnlyRulesInput{}
	}

	output = &DescribeIpv6EgressOnlyRulesOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DescribeIpv6EgressOnlyRules API operation for VPC.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VPC's
// API operation DescribeIpv6EgressOnlyRules for usage and error information.
func (c *VPC) DescribeIpv6EgressOnlyRules(input *DescribeIpv6EgressOnlyRulesInput) (*DescribeIpv6EgressOnlyRulesOutput, error) {
	req, out := c.DescribeIpv6EgressOnlyRulesRequest(input)
	return out, req.Send()
}

// DescribeIpv6EgressOnlyRulesWithContext is the same as DescribeIpv6EgressOnlyRules with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeIpv6EgressOnlyRules for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPC) DescribeIpv6EgressOnlyRulesWithContext(ctx volcengine.Context, input *DescribeIpv6EgressOnlyRulesInput, opts ...request.Option) (*DescribeIpv6EgressOnlyRulesOutput, error) {
	req, out := c.DescribeIpv6EgressOnlyRulesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DescribeIpv6EgressOnlyRulesInput struct {
	_ struct{} `type:"structure"`

	InstanceId *string `type:"string"`

	Ipv6Address *string `type:"string"`

	Ipv6EgressOnlyRuleIds *string `type:"string"`

	// Ipv6GatewayId is a required field
	Ipv6GatewayId *string `type:"string" required:"true"`

	MaxResults *int64 `type:"integer"`

	Name *string `type:"string"`

	NextToken *string `type:"string"`
}

// String returns the string representation
func (s DescribeIpv6EgressOnlyRulesInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeIpv6EgressOnlyRulesInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeIpv6EgressOnlyRulesInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DescribeIpv6EgressOnlyRulesInput"}
	if s.Ipv6GatewayId == nil {
		invalidParams.Add(request.NewErrParamRequired("Ipv6GatewayId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetInstanceId sets the InstanceId field's value.
func (s *DescribeIpv6EgressOnlyRulesInput) SetInstanceId(v string) *DescribeIpv6EgressOnlyRulesInput {
	s.InstanceId = &v
	return s
}

// SetIpv6Address sets the Ipv6Address field's value.
func (s *DescribeIpv6EgressOnlyRulesInput) SetIpv6Address(v string) *DescribeIpv6EgressOnlyRulesInput {
	s.Ipv6Address = &v
	return s
}

// SetIpv6EgressOnlyRuleIds sets the Ipv6EgressOnlyRuleIds field's value.
func (s *DescribeIpv6EgressOnlyRulesInput) SetIpv6EgressOnlyRuleIds(v string) *DescribeIpv6EgressOnlyRulesInput {
	s.Ipv6EgressOnlyRuleIds = &v
	return s
}

// SetIpv6GatewayId sets the Ipv6GatewayId field's value.
func (s *DescribeIpv6EgressOnlyRulesInput) SetIpv6GatewayId(v string) *DescribeIpv6EgressOnlyRulesInput {
	s.Ipv6GatewayId = &v
	return s
}

// SetMaxResults sets the MaxResults field's value.
func (s *DescribeIpv6EgressOnlyRulesInput) SetMaxResults(v int64) *DescribeIpv6EgressOnlyRulesInput {
	s.MaxResults = &v
	return s
}

// SetName sets the Name field's value.
func (s *DescribeIpv6EgressOnlyRulesInput) SetName(v string) *DescribeIpv6EgressOnlyRulesInput {
	s.Name = &v
	return s
}

// SetNextToken sets the NextToken field's value.
func (s *DescribeIpv6EgressOnlyRulesInput) SetNextToken(v string) *DescribeIpv6EgressOnlyRulesInput {
	s.NextToken = &v
	return s
}

type DescribeIpv6EgressOnlyRulesOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	Ipv6EgressRules []*Ipv6EgressRuleForDescribeIpv6EgressOnlyRulesOutput `type:"list"`

	NextToken *string `type:"string"`

	RequestId *string `type:"string"`
}

// String returns the string representation
func (s DescribeIpv6EgressOnlyRulesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeIpv6EgressOnlyRulesOutput) GoString() string {
	return s.String()
}

// SetIpv6EgressRules sets the Ipv6EgressRules field's value.
func (s *DescribeIpv6EgressOnlyRulesOutput) SetIpv6EgressRules(v []*Ipv6EgressRuleForDescribeIpv6EgressOnlyRulesOutput) *DescribeIpv6EgressOnlyRulesOutput {
	s.Ipv6EgressRules = v
	return s
}

// SetNextToken sets the NextToken field's value.
func (s *DescribeIpv6EgressOnlyRulesOutput) SetNextToken(v string) *DescribeIpv6EgressOnlyRulesOutput {
	s.NextToken = &v
	return s
}

// SetRequestId sets the RequestId field's value.
func (s *DescribeIpv6EgressOnlyRulesOutput) SetRequestId(v string) *DescribeIpv6EgressOnlyRulesOutput {
	s.RequestId = &v
	return s
}

type Ipv6EgressRuleForDescribeIpv6EgressOnlyRulesOutput struct {
	_ struct{} `type:"structure"`

	Description *string `type:"string"`

	InstanceId *string `type:"string"`

	Ipv6EgressOnlyRuleId *string `type:"string"`

	Name *string `type:"string"`

	Status *string `type:"string"`
}

// String returns the string representation
func (s Ipv6EgressRuleForDescribeIpv6EgressOnlyRulesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s Ipv6EgressRuleForDescribeIpv6EgressOnlyRulesOutput) GoString() string {
	return s.String()
}

// SetDescription sets the Description field's value.
func (s *Ipv6EgressRuleForDescribeIpv6EgressOnlyRulesOutput) SetDescription(v string) *Ipv6EgressRuleForDescribeIpv6EgressOnlyRulesOutput {
	s.Description = &v
	return s
}

// SetInstanceId sets the InstanceId field's value.
func (s *Ipv6EgressRuleForDescribeIpv6EgressOnlyRulesOutput) SetInstanceId(v string) *Ipv6EgressRuleForDescribeIpv6EgressOnlyRulesOutput {
	s.InstanceId = &v
	return s
}

// SetIpv6EgressOnlyRuleId sets the Ipv6EgressOnlyRuleId field's value.
func (s *Ipv6EgressRuleForDescribeIpv6EgressOnlyRulesOutput) SetIpv6EgressOnlyRuleId(v string) *Ipv6EgressRuleForDescribeIpv6EgressOnlyRulesOutput {
	s.Ipv6EgressOnlyRuleId = &v
	return s
}

// SetName sets the Name field's value.
func (s *Ipv6EgressRuleForDescribeIpv6EgressOnlyRulesOutput) SetName(v string) *Ipv6EgressRuleForDescribeIpv6EgressOnlyRulesOutput {
	s.Name = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *Ipv6EgressRuleForDescribeIpv6EgressOnlyRulesOutput) SetStatus(v string) *Ipv6EgressRuleForDescribeIpv6EgressOnlyRulesOutput {
	s.Status = &v
	return s
}
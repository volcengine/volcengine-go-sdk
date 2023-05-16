// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package vpc

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opRevokeSecurityGroupEgressCommon = "RevokeSecurityGroupEgress"

// RevokeSecurityGroupEgressCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the RevokeSecurityGroupEgressCommon operation. The "output" return
// value will be populated with the RevokeSecurityGroupEgressCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned RevokeSecurityGroupEgressCommon Request to send the API call to the service.
// the "output" return value is not valid until after RevokeSecurityGroupEgressCommon Send returns without error.
//
// See RevokeSecurityGroupEgressCommon for more information on using the RevokeSecurityGroupEgressCommon
// API call, and error handling.
//
//	// Example sending a request using the RevokeSecurityGroupEgressCommonRequest method.
//	req, resp := client.RevokeSecurityGroupEgressCommonRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *VPC) RevokeSecurityGroupEgressCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opRevokeSecurityGroupEgressCommon,
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

// RevokeSecurityGroupEgressCommon API operation for VPC.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VPC's
// API operation RevokeSecurityGroupEgressCommon for usage and error information.
func (c *VPC) RevokeSecurityGroupEgressCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.RevokeSecurityGroupEgressCommonRequest(input)
	return out, req.Send()
}

// RevokeSecurityGroupEgressCommonWithContext is the same as RevokeSecurityGroupEgressCommon with the addition of
// the ability to pass a context and additional request options.
//
// See RevokeSecurityGroupEgressCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPC) RevokeSecurityGroupEgressCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.RevokeSecurityGroupEgressCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opRevokeSecurityGroupEgress = "RevokeSecurityGroupEgress"

// RevokeSecurityGroupEgressRequest generates a "volcengine/request.Request" representing the
// client's request for the RevokeSecurityGroupEgress operation. The "output" return
// value will be populated with the RevokeSecurityGroupEgressCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned RevokeSecurityGroupEgressCommon Request to send the API call to the service.
// the "output" return value is not valid until after RevokeSecurityGroupEgressCommon Send returns without error.
//
// See RevokeSecurityGroupEgress for more information on using the RevokeSecurityGroupEgress
// API call, and error handling.
//
//	// Example sending a request using the RevokeSecurityGroupEgressRequest method.
//	req, resp := client.RevokeSecurityGroupEgressRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *VPC) RevokeSecurityGroupEgressRequest(input *RevokeSecurityGroupEgressInput) (req *request.Request, output *RevokeSecurityGroupEgressOutput) {
	op := &request.Operation{
		Name:       opRevokeSecurityGroupEgress,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &RevokeSecurityGroupEgressInput{}
	}

	output = &RevokeSecurityGroupEgressOutput{}
	req = c.newRequest(op, input, output)

	return
}

// RevokeSecurityGroupEgress API operation for VPC.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VPC's
// API operation RevokeSecurityGroupEgress for usage and error information.
func (c *VPC) RevokeSecurityGroupEgress(input *RevokeSecurityGroupEgressInput) (*RevokeSecurityGroupEgressOutput, error) {
	req, out := c.RevokeSecurityGroupEgressRequest(input)
	return out, req.Send()
}

// RevokeSecurityGroupEgressWithContext is the same as RevokeSecurityGroupEgress with the addition of
// the ability to pass a context and additional request options.
//
// See RevokeSecurityGroupEgress for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPC) RevokeSecurityGroupEgressWithContext(ctx volcengine.Context, input *RevokeSecurityGroupEgressInput, opts ...request.Option) (*RevokeSecurityGroupEgressOutput, error) {
	req, out := c.RevokeSecurityGroupEgressRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type RevokeSecurityGroupEgressInput struct {
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
func (s RevokeSecurityGroupEgressInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s RevokeSecurityGroupEgressInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RevokeSecurityGroupEgressInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "RevokeSecurityGroupEgressInput"}
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
func (s *RevokeSecurityGroupEgressInput) SetCidrIp(v string) *RevokeSecurityGroupEgressInput {
	s.CidrIp = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *RevokeSecurityGroupEgressInput) SetDescription(v string) *RevokeSecurityGroupEgressInput {
	s.Description = &v
	return s
}

// SetPolicy sets the Policy field's value.
func (s *RevokeSecurityGroupEgressInput) SetPolicy(v string) *RevokeSecurityGroupEgressInput {
	s.Policy = &v
	return s
}

// SetPortEnd sets the PortEnd field's value.
func (s *RevokeSecurityGroupEgressInput) SetPortEnd(v int64) *RevokeSecurityGroupEgressInput {
	s.PortEnd = &v
	return s
}

// SetPortStart sets the PortStart field's value.
func (s *RevokeSecurityGroupEgressInput) SetPortStart(v int64) *RevokeSecurityGroupEgressInput {
	s.PortStart = &v
	return s
}

// SetPriority sets the Priority field's value.
func (s *RevokeSecurityGroupEgressInput) SetPriority(v int64) *RevokeSecurityGroupEgressInput {
	s.Priority = &v
	return s
}

// SetProtocol sets the Protocol field's value.
func (s *RevokeSecurityGroupEgressInput) SetProtocol(v string) *RevokeSecurityGroupEgressInput {
	s.Protocol = &v
	return s
}

// SetSecurityGroupId sets the SecurityGroupId field's value.
func (s *RevokeSecurityGroupEgressInput) SetSecurityGroupId(v string) *RevokeSecurityGroupEgressInput {
	s.SecurityGroupId = &v
	return s
}

// SetSourceGroupId sets the SourceGroupId field's value.
func (s *RevokeSecurityGroupEgressInput) SetSourceGroupId(v string) *RevokeSecurityGroupEgressInput {
	s.SourceGroupId = &v
	return s
}

type RevokeSecurityGroupEgressOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	RequestId *string `type:"string"`
}

// String returns the string representation
func (s RevokeSecurityGroupEgressOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s RevokeSecurityGroupEgressOutput) GoString() string {
	return s.String()
}

// SetRequestId sets the RequestId field's value.
func (s *RevokeSecurityGroupEgressOutput) SetRequestId(v string) *RevokeSecurityGroupEgressOutput {
	s.RequestId = &v
	return s
}

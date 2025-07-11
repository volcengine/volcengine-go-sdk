// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package vpn

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opCreateVpnGatewayCommon = "CreateVpnGateway"

// CreateVpnGatewayCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateVpnGatewayCommon operation. The "output" return
// value will be populated with the CreateVpnGatewayCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateVpnGatewayCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateVpnGatewayCommon Send returns without error.
//
// See CreateVpnGatewayCommon for more information on using the CreateVpnGatewayCommon
// API call, and error handling.
//
//    // Example sending a request using the CreateVpnGatewayCommonRequest method.
//    req, resp := client.CreateVpnGatewayCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VPN) CreateVpnGatewayCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opCreateVpnGatewayCommon,
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

// CreateVpnGatewayCommon API operation for VPN.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VPN's
// API operation CreateVpnGatewayCommon for usage and error information.
func (c *VPN) CreateVpnGatewayCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.CreateVpnGatewayCommonRequest(input)
	return out, req.Send()
}

// CreateVpnGatewayCommonWithContext is the same as CreateVpnGatewayCommon with the addition of
// the ability to pass a context and additional request options.
//
// See CreateVpnGatewayCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPN) CreateVpnGatewayCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.CreateVpnGatewayCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opCreateVpnGateway = "CreateVpnGateway"

// CreateVpnGatewayRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateVpnGateway operation. The "output" return
// value will be populated with the CreateVpnGatewayCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateVpnGatewayCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateVpnGatewayCommon Send returns without error.
//
// See CreateVpnGateway for more information on using the CreateVpnGateway
// API call, and error handling.
//
//    // Example sending a request using the CreateVpnGatewayRequest method.
//    req, resp := client.CreateVpnGatewayRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VPN) CreateVpnGatewayRequest(input *CreateVpnGatewayInput) (req *request.Request, output *CreateVpnGatewayOutput) {
	op := &request.Operation{
		Name:       opCreateVpnGateway,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateVpnGatewayInput{}
	}

	output = &CreateVpnGatewayOutput{}
	req = c.newRequest(op, input, output)

	return
}

// CreateVpnGateway API operation for VPN.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VPN's
// API operation CreateVpnGateway for usage and error information.
func (c *VPN) CreateVpnGateway(input *CreateVpnGatewayInput) (*CreateVpnGatewayOutput, error) {
	req, out := c.CreateVpnGatewayRequest(input)
	return out, req.Send()
}

// CreateVpnGatewayWithContext is the same as CreateVpnGateway with the addition of
// the ability to pass a context and additional request options.
//
// See CreateVpnGateway for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPN) CreateVpnGatewayWithContext(ctx volcengine.Context, input *CreateVpnGatewayInput, opts ...request.Option) (*CreateVpnGatewayOutput, error) {
	req, out := c.CreateVpnGatewayRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type CreateVpnGatewayInput struct {
	_ struct{} `type:"structure"`

	Asn *int64 `type:"integer"`

	// Bandwidth is a required field
	Bandwidth *int64 `type:"integer" required:"true"`

	BillingType *int64 `type:"integer"`

	ClientToken *string `type:"string"`

	Description *string `type:"string"`

	DualTunnelEnabled *bool `type:"boolean"`

	IpsecEnabled *bool `type:"boolean"`

	Period *int64 `type:"integer"`

	PeriodUnit *string `type:"string"`

	ProjectName *string `type:"string"`

	SecondarySubnetId *string `type:"string"`

	SslEnabled *bool `type:"boolean"`

	SslMaxConnections *int64 `type:"integer"`

	// SubnetId is a required field
	SubnetId *string `type:"string" required:"true"`

	Tags []*TagForCreateVpnGatewayInput `type:"list"`

	// VpcId is a required field
	VpcId *string `type:"string" required:"true"`

	VpnGatewayName *string `type:"string"`
}

// String returns the string representation
func (s CreateVpnGatewayInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateVpnGatewayInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateVpnGatewayInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "CreateVpnGatewayInput"}
	if s.Bandwidth == nil {
		invalidParams.Add(request.NewErrParamRequired("Bandwidth"))
	}
	if s.SubnetId == nil {
		invalidParams.Add(request.NewErrParamRequired("SubnetId"))
	}
	if s.VpcId == nil {
		invalidParams.Add(request.NewErrParamRequired("VpcId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetAsn sets the Asn field's value.
func (s *CreateVpnGatewayInput) SetAsn(v int64) *CreateVpnGatewayInput {
	s.Asn = &v
	return s
}

// SetBandwidth sets the Bandwidth field's value.
func (s *CreateVpnGatewayInput) SetBandwidth(v int64) *CreateVpnGatewayInput {
	s.Bandwidth = &v
	return s
}

// SetBillingType sets the BillingType field's value.
func (s *CreateVpnGatewayInput) SetBillingType(v int64) *CreateVpnGatewayInput {
	s.BillingType = &v
	return s
}

// SetClientToken sets the ClientToken field's value.
func (s *CreateVpnGatewayInput) SetClientToken(v string) *CreateVpnGatewayInput {
	s.ClientToken = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *CreateVpnGatewayInput) SetDescription(v string) *CreateVpnGatewayInput {
	s.Description = &v
	return s
}

// SetDualTunnelEnabled sets the DualTunnelEnabled field's value.
func (s *CreateVpnGatewayInput) SetDualTunnelEnabled(v bool) *CreateVpnGatewayInput {
	s.DualTunnelEnabled = &v
	return s
}

// SetIpsecEnabled sets the IpsecEnabled field's value.
func (s *CreateVpnGatewayInput) SetIpsecEnabled(v bool) *CreateVpnGatewayInput {
	s.IpsecEnabled = &v
	return s
}

// SetPeriod sets the Period field's value.
func (s *CreateVpnGatewayInput) SetPeriod(v int64) *CreateVpnGatewayInput {
	s.Period = &v
	return s
}

// SetPeriodUnit sets the PeriodUnit field's value.
func (s *CreateVpnGatewayInput) SetPeriodUnit(v string) *CreateVpnGatewayInput {
	s.PeriodUnit = &v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *CreateVpnGatewayInput) SetProjectName(v string) *CreateVpnGatewayInput {
	s.ProjectName = &v
	return s
}

// SetSecondarySubnetId sets the SecondarySubnetId field's value.
func (s *CreateVpnGatewayInput) SetSecondarySubnetId(v string) *CreateVpnGatewayInput {
	s.SecondarySubnetId = &v
	return s
}

// SetSslEnabled sets the SslEnabled field's value.
func (s *CreateVpnGatewayInput) SetSslEnabled(v bool) *CreateVpnGatewayInput {
	s.SslEnabled = &v
	return s
}

// SetSslMaxConnections sets the SslMaxConnections field's value.
func (s *CreateVpnGatewayInput) SetSslMaxConnections(v int64) *CreateVpnGatewayInput {
	s.SslMaxConnections = &v
	return s
}

// SetSubnetId sets the SubnetId field's value.
func (s *CreateVpnGatewayInput) SetSubnetId(v string) *CreateVpnGatewayInput {
	s.SubnetId = &v
	return s
}

// SetTags sets the Tags field's value.
func (s *CreateVpnGatewayInput) SetTags(v []*TagForCreateVpnGatewayInput) *CreateVpnGatewayInput {
	s.Tags = v
	return s
}

// SetVpcId sets the VpcId field's value.
func (s *CreateVpnGatewayInput) SetVpcId(v string) *CreateVpnGatewayInput {
	s.VpcId = &v
	return s
}

// SetVpnGatewayName sets the VpnGatewayName field's value.
func (s *CreateVpnGatewayInput) SetVpnGatewayName(v string) *CreateVpnGatewayInput {
	s.VpnGatewayName = &v
	return s
}

type CreateVpnGatewayOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	OrderId *string `type:"string"`

	RequestId *string `type:"string"`

	VpnGatewayId *string `type:"string"`
}

// String returns the string representation
func (s CreateVpnGatewayOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateVpnGatewayOutput) GoString() string {
	return s.String()
}

// SetOrderId sets the OrderId field's value.
func (s *CreateVpnGatewayOutput) SetOrderId(v string) *CreateVpnGatewayOutput {
	s.OrderId = &v
	return s
}

// SetRequestId sets the RequestId field's value.
func (s *CreateVpnGatewayOutput) SetRequestId(v string) *CreateVpnGatewayOutput {
	s.RequestId = &v
	return s
}

// SetVpnGatewayId sets the VpnGatewayId field's value.
func (s *CreateVpnGatewayOutput) SetVpnGatewayId(v string) *CreateVpnGatewayOutput {
	s.VpnGatewayId = &v
	return s
}

type TagForCreateVpnGatewayInput struct {
	_ struct{} `type:"structure"`

	Key *string `type:"string"`

	Value *string `type:"string"`
}

// String returns the string representation
func (s TagForCreateVpnGatewayInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s TagForCreateVpnGatewayInput) GoString() string {
	return s.String()
}

// SetKey sets the Key field's value.
func (s *TagForCreateVpnGatewayInput) SetKey(v string) *TagForCreateVpnGatewayInput {
	s.Key = &v
	return s
}

// SetValue sets the Value field's value.
func (s *TagForCreateVpnGatewayInput) SetValue(v string) *TagForCreateVpnGatewayInput {
	s.Value = &v
	return s
}

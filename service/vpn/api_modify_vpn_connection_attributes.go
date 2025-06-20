// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package vpn

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opModifyVpnConnectionAttributesCommon = "ModifyVpnConnectionAttributes"

// ModifyVpnConnectionAttributesCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ModifyVpnConnectionAttributesCommon operation. The "output" return
// value will be populated with the ModifyVpnConnectionAttributesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyVpnConnectionAttributesCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyVpnConnectionAttributesCommon Send returns without error.
//
// See ModifyVpnConnectionAttributesCommon for more information on using the ModifyVpnConnectionAttributesCommon
// API call, and error handling.
//
//    // Example sending a request using the ModifyVpnConnectionAttributesCommonRequest method.
//    req, resp := client.ModifyVpnConnectionAttributesCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VPN) ModifyVpnConnectionAttributesCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opModifyVpnConnectionAttributesCommon,
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

// ModifyVpnConnectionAttributesCommon API operation for VPN.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VPN's
// API operation ModifyVpnConnectionAttributesCommon for usage and error information.
func (c *VPN) ModifyVpnConnectionAttributesCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ModifyVpnConnectionAttributesCommonRequest(input)
	return out, req.Send()
}

// ModifyVpnConnectionAttributesCommonWithContext is the same as ModifyVpnConnectionAttributesCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyVpnConnectionAttributesCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPN) ModifyVpnConnectionAttributesCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ModifyVpnConnectionAttributesCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opModifyVpnConnectionAttributes = "ModifyVpnConnectionAttributes"

// ModifyVpnConnectionAttributesRequest generates a "volcengine/request.Request" representing the
// client's request for the ModifyVpnConnectionAttributes operation. The "output" return
// value will be populated with the ModifyVpnConnectionAttributesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyVpnConnectionAttributesCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyVpnConnectionAttributesCommon Send returns without error.
//
// See ModifyVpnConnectionAttributes for more information on using the ModifyVpnConnectionAttributes
// API call, and error handling.
//
//    // Example sending a request using the ModifyVpnConnectionAttributesRequest method.
//    req, resp := client.ModifyVpnConnectionAttributesRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VPN) ModifyVpnConnectionAttributesRequest(input *ModifyVpnConnectionAttributesInput) (req *request.Request, output *ModifyVpnConnectionAttributesOutput) {
	op := &request.Operation{
		Name:       opModifyVpnConnectionAttributes,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ModifyVpnConnectionAttributesInput{}
	}

	output = &ModifyVpnConnectionAttributesOutput{}
	req = c.newRequest(op, input, output)

	return
}

// ModifyVpnConnectionAttributes API operation for VPN.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VPN's
// API operation ModifyVpnConnectionAttributes for usage and error information.
func (c *VPN) ModifyVpnConnectionAttributes(input *ModifyVpnConnectionAttributesInput) (*ModifyVpnConnectionAttributesOutput, error) {
	req, out := c.ModifyVpnConnectionAttributesRequest(input)
	return out, req.Send()
}

// ModifyVpnConnectionAttributesWithContext is the same as ModifyVpnConnectionAttributes with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyVpnConnectionAttributes for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPN) ModifyVpnConnectionAttributesWithContext(ctx volcengine.Context, input *ModifyVpnConnectionAttributesInput, opts ...request.Option) (*ModifyVpnConnectionAttributesOutput, error) {
	req, out := c.ModifyVpnConnectionAttributesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type BgpConfigForModifyVpnConnectionAttributesInput struct {
	_ struct{} `type:"structure"`

	EnableBgp *bool `type:"boolean"`

	LocalBgpIp *string `type:"string"`

	TunnelCidr *string `type:"string"`
}

// String returns the string representation
func (s BgpConfigForModifyVpnConnectionAttributesInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s BgpConfigForModifyVpnConnectionAttributesInput) GoString() string {
	return s.String()
}

// SetEnableBgp sets the EnableBgp field's value.
func (s *BgpConfigForModifyVpnConnectionAttributesInput) SetEnableBgp(v bool) *BgpConfigForModifyVpnConnectionAttributesInput {
	s.EnableBgp = &v
	return s
}

// SetLocalBgpIp sets the LocalBgpIp field's value.
func (s *BgpConfigForModifyVpnConnectionAttributesInput) SetLocalBgpIp(v string) *BgpConfigForModifyVpnConnectionAttributesInput {
	s.LocalBgpIp = &v
	return s
}

// SetTunnelCidr sets the TunnelCidr field's value.
func (s *BgpConfigForModifyVpnConnectionAttributesInput) SetTunnelCidr(v string) *BgpConfigForModifyVpnConnectionAttributesInput {
	s.TunnelCidr = &v
	return s
}

type IkeConfigForModifyVpnConnectionAttributesInput struct {
	_ struct{} `type:"structure"`

	AuthAlg *string `type:"string"`

	DhGroup *string `type:"string"`

	EncAlg *string `type:"string"`

	Lifetime *int64 `type:"integer"`

	LocalId *string `type:"string"`

	Mode *string `type:"string"`

	Psk *string `type:"string"`

	RemoteId *string `type:"string"`

	Version *string `type:"string"`
}

// String returns the string representation
func (s IkeConfigForModifyVpnConnectionAttributesInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s IkeConfigForModifyVpnConnectionAttributesInput) GoString() string {
	return s.String()
}

// SetAuthAlg sets the AuthAlg field's value.
func (s *IkeConfigForModifyVpnConnectionAttributesInput) SetAuthAlg(v string) *IkeConfigForModifyVpnConnectionAttributesInput {
	s.AuthAlg = &v
	return s
}

// SetDhGroup sets the DhGroup field's value.
func (s *IkeConfigForModifyVpnConnectionAttributesInput) SetDhGroup(v string) *IkeConfigForModifyVpnConnectionAttributesInput {
	s.DhGroup = &v
	return s
}

// SetEncAlg sets the EncAlg field's value.
func (s *IkeConfigForModifyVpnConnectionAttributesInput) SetEncAlg(v string) *IkeConfigForModifyVpnConnectionAttributesInput {
	s.EncAlg = &v
	return s
}

// SetLifetime sets the Lifetime field's value.
func (s *IkeConfigForModifyVpnConnectionAttributesInput) SetLifetime(v int64) *IkeConfigForModifyVpnConnectionAttributesInput {
	s.Lifetime = &v
	return s
}

// SetLocalId sets the LocalId field's value.
func (s *IkeConfigForModifyVpnConnectionAttributesInput) SetLocalId(v string) *IkeConfigForModifyVpnConnectionAttributesInput {
	s.LocalId = &v
	return s
}

// SetMode sets the Mode field's value.
func (s *IkeConfigForModifyVpnConnectionAttributesInput) SetMode(v string) *IkeConfigForModifyVpnConnectionAttributesInput {
	s.Mode = &v
	return s
}

// SetPsk sets the Psk field's value.
func (s *IkeConfigForModifyVpnConnectionAttributesInput) SetPsk(v string) *IkeConfigForModifyVpnConnectionAttributesInput {
	s.Psk = &v
	return s
}

// SetRemoteId sets the RemoteId field's value.
func (s *IkeConfigForModifyVpnConnectionAttributesInput) SetRemoteId(v string) *IkeConfigForModifyVpnConnectionAttributesInput {
	s.RemoteId = &v
	return s
}

// SetVersion sets the Version field's value.
func (s *IkeConfigForModifyVpnConnectionAttributesInput) SetVersion(v string) *IkeConfigForModifyVpnConnectionAttributesInput {
	s.Version = &v
	return s
}

type IpsecConfigForModifyVpnConnectionAttributesInput struct {
	_ struct{} `type:"structure"`

	AuthAlg *string `type:"string"`

	DhGroup *string `type:"string"`

	EncAlg *string `type:"string"`

	Lifetime *int64 `type:"integer"`
}

// String returns the string representation
func (s IpsecConfigForModifyVpnConnectionAttributesInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s IpsecConfigForModifyVpnConnectionAttributesInput) GoString() string {
	return s.String()
}

// SetAuthAlg sets the AuthAlg field's value.
func (s *IpsecConfigForModifyVpnConnectionAttributesInput) SetAuthAlg(v string) *IpsecConfigForModifyVpnConnectionAttributesInput {
	s.AuthAlg = &v
	return s
}

// SetDhGroup sets the DhGroup field's value.
func (s *IpsecConfigForModifyVpnConnectionAttributesInput) SetDhGroup(v string) *IpsecConfigForModifyVpnConnectionAttributesInput {
	s.DhGroup = &v
	return s
}

// SetEncAlg sets the EncAlg field's value.
func (s *IpsecConfigForModifyVpnConnectionAttributesInput) SetEncAlg(v string) *IpsecConfigForModifyVpnConnectionAttributesInput {
	s.EncAlg = &v
	return s
}

// SetLifetime sets the Lifetime field's value.
func (s *IpsecConfigForModifyVpnConnectionAttributesInput) SetLifetime(v int64) *IpsecConfigForModifyVpnConnectionAttributesInput {
	s.Lifetime = &v
	return s
}

type ModifyVpnConnectionAttributesInput struct {
	_ struct{} `type:"structure"`

	BgpConfig *BgpConfigForModifyVpnConnectionAttributesInput `type:"structure"`

	CustomerGatewayId *string `type:"string"`

	Description *string `type:"string"`

	DpdAction *string `type:"string"`

	IkeConfig *IkeConfigForModifyVpnConnectionAttributesInput `type:"structure"`

	IpsecConfig *IpsecConfigForModifyVpnConnectionAttributesInput `type:"structure"`

	LocalSubnet []*string `type:"list"`

	LogEnabled *bool `type:"boolean"`

	NatTraversal *bool `type:"boolean"`

	NegotiateInstantly *bool `type:"boolean"`

	RemoteSubnet []*string `type:"list"`

	TunnelOptions []*TunnelOptionForModifyVpnConnectionAttributesInput `type:"list"`

	// VpnConnectionId is a required field
	VpnConnectionId *string `type:"string" required:"true"`

	VpnConnectionName *string `type:"string"`
}

// String returns the string representation
func (s ModifyVpnConnectionAttributesInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyVpnConnectionAttributesInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ModifyVpnConnectionAttributesInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ModifyVpnConnectionAttributesInput"}
	if s.VpnConnectionId == nil {
		invalidParams.Add(request.NewErrParamRequired("VpnConnectionId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetBgpConfig sets the BgpConfig field's value.
func (s *ModifyVpnConnectionAttributesInput) SetBgpConfig(v *BgpConfigForModifyVpnConnectionAttributesInput) *ModifyVpnConnectionAttributesInput {
	s.BgpConfig = v
	return s
}

// SetCustomerGatewayId sets the CustomerGatewayId field's value.
func (s *ModifyVpnConnectionAttributesInput) SetCustomerGatewayId(v string) *ModifyVpnConnectionAttributesInput {
	s.CustomerGatewayId = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *ModifyVpnConnectionAttributesInput) SetDescription(v string) *ModifyVpnConnectionAttributesInput {
	s.Description = &v
	return s
}

// SetDpdAction sets the DpdAction field's value.
func (s *ModifyVpnConnectionAttributesInput) SetDpdAction(v string) *ModifyVpnConnectionAttributesInput {
	s.DpdAction = &v
	return s
}

// SetIkeConfig sets the IkeConfig field's value.
func (s *ModifyVpnConnectionAttributesInput) SetIkeConfig(v *IkeConfigForModifyVpnConnectionAttributesInput) *ModifyVpnConnectionAttributesInput {
	s.IkeConfig = v
	return s
}

// SetIpsecConfig sets the IpsecConfig field's value.
func (s *ModifyVpnConnectionAttributesInput) SetIpsecConfig(v *IpsecConfigForModifyVpnConnectionAttributesInput) *ModifyVpnConnectionAttributesInput {
	s.IpsecConfig = v
	return s
}

// SetLocalSubnet sets the LocalSubnet field's value.
func (s *ModifyVpnConnectionAttributesInput) SetLocalSubnet(v []*string) *ModifyVpnConnectionAttributesInput {
	s.LocalSubnet = v
	return s
}

// SetLogEnabled sets the LogEnabled field's value.
func (s *ModifyVpnConnectionAttributesInput) SetLogEnabled(v bool) *ModifyVpnConnectionAttributesInput {
	s.LogEnabled = &v
	return s
}

// SetNatTraversal sets the NatTraversal field's value.
func (s *ModifyVpnConnectionAttributesInput) SetNatTraversal(v bool) *ModifyVpnConnectionAttributesInput {
	s.NatTraversal = &v
	return s
}

// SetNegotiateInstantly sets the NegotiateInstantly field's value.
func (s *ModifyVpnConnectionAttributesInput) SetNegotiateInstantly(v bool) *ModifyVpnConnectionAttributesInput {
	s.NegotiateInstantly = &v
	return s
}

// SetRemoteSubnet sets the RemoteSubnet field's value.
func (s *ModifyVpnConnectionAttributesInput) SetRemoteSubnet(v []*string) *ModifyVpnConnectionAttributesInput {
	s.RemoteSubnet = v
	return s
}

// SetTunnelOptions sets the TunnelOptions field's value.
func (s *ModifyVpnConnectionAttributesInput) SetTunnelOptions(v []*TunnelOptionForModifyVpnConnectionAttributesInput) *ModifyVpnConnectionAttributesInput {
	s.TunnelOptions = v
	return s
}

// SetVpnConnectionId sets the VpnConnectionId field's value.
func (s *ModifyVpnConnectionAttributesInput) SetVpnConnectionId(v string) *ModifyVpnConnectionAttributesInput {
	s.VpnConnectionId = &v
	return s
}

// SetVpnConnectionName sets the VpnConnectionName field's value.
func (s *ModifyVpnConnectionAttributesInput) SetVpnConnectionName(v string) *ModifyVpnConnectionAttributesInput {
	s.VpnConnectionName = &v
	return s
}

type ModifyVpnConnectionAttributesOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	OrderId *string `type:"string"`

	RequestId *string `type:"string"`
}

// String returns the string representation
func (s ModifyVpnConnectionAttributesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyVpnConnectionAttributesOutput) GoString() string {
	return s.String()
}

// SetOrderId sets the OrderId field's value.
func (s *ModifyVpnConnectionAttributesOutput) SetOrderId(v string) *ModifyVpnConnectionAttributesOutput {
	s.OrderId = &v
	return s
}

// SetRequestId sets the RequestId field's value.
func (s *ModifyVpnConnectionAttributesOutput) SetRequestId(v string) *ModifyVpnConnectionAttributesOutput {
	s.RequestId = &v
	return s
}

type TunnelOptionForModifyVpnConnectionAttributesInput struct {
	_ struct{} `type:"structure"`

	CustomerGatewayId *string `type:"string"`

	DpdAction *string `type:"string"`

	IkeConfig *IkeConfigForModifyVpnConnectionAttributesInput `type:"structure"`

	IpsecConfig *IpsecConfigForModifyVpnConnectionAttributesInput `type:"structure"`

	NatTraversal *bool `type:"boolean"`

	TunnelId *string `type:"string"`
}

// String returns the string representation
func (s TunnelOptionForModifyVpnConnectionAttributesInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s TunnelOptionForModifyVpnConnectionAttributesInput) GoString() string {
	return s.String()
}

// SetCustomerGatewayId sets the CustomerGatewayId field's value.
func (s *TunnelOptionForModifyVpnConnectionAttributesInput) SetCustomerGatewayId(v string) *TunnelOptionForModifyVpnConnectionAttributesInput {
	s.CustomerGatewayId = &v
	return s
}

// SetDpdAction sets the DpdAction field's value.
func (s *TunnelOptionForModifyVpnConnectionAttributesInput) SetDpdAction(v string) *TunnelOptionForModifyVpnConnectionAttributesInput {
	s.DpdAction = &v
	return s
}

// SetIkeConfig sets the IkeConfig field's value.
func (s *TunnelOptionForModifyVpnConnectionAttributesInput) SetIkeConfig(v *IkeConfigForModifyVpnConnectionAttributesInput) *TunnelOptionForModifyVpnConnectionAttributesInput {
	s.IkeConfig = v
	return s
}

// SetIpsecConfig sets the IpsecConfig field's value.
func (s *TunnelOptionForModifyVpnConnectionAttributesInput) SetIpsecConfig(v *IpsecConfigForModifyVpnConnectionAttributesInput) *TunnelOptionForModifyVpnConnectionAttributesInput {
	s.IpsecConfig = v
	return s
}

// SetNatTraversal sets the NatTraversal field's value.
func (s *TunnelOptionForModifyVpnConnectionAttributesInput) SetNatTraversal(v bool) *TunnelOptionForModifyVpnConnectionAttributesInput {
	s.NatTraversal = &v
	return s
}

// SetTunnelId sets the TunnelId field's value.
func (s *TunnelOptionForModifyVpnConnectionAttributesInput) SetTunnelId(v string) *TunnelOptionForModifyVpnConnectionAttributesInput {
	s.TunnelId = &v
	return s
}

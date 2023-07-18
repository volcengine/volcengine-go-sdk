// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package vpn

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opSetVpnGatewayRenewalCommon = "SetVpnGatewayRenewal"

// SetVpnGatewayRenewalCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the SetVpnGatewayRenewalCommon operation. The "output" return
// value will be populated with the SetVpnGatewayRenewalCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned SetVpnGatewayRenewalCommon Request to send the API call to the service.
// the "output" return value is not valid until after SetVpnGatewayRenewalCommon Send returns without error.
//
// See SetVpnGatewayRenewalCommon for more information on using the SetVpnGatewayRenewalCommon
// API call, and error handling.
//
//	// Example sending a request using the SetVpnGatewayRenewalCommonRequest method.
//	req, resp := client.SetVpnGatewayRenewalCommonRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *VPN) SetVpnGatewayRenewalCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opSetVpnGatewayRenewalCommon,
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

// SetVpnGatewayRenewalCommon API operation for VPN.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VPN's
// API operation SetVpnGatewayRenewalCommon for usage and error information.
func (c *VPN) SetVpnGatewayRenewalCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.SetVpnGatewayRenewalCommonRequest(input)
	return out, req.Send()
}

// SetVpnGatewayRenewalCommonWithContext is the same as SetVpnGatewayRenewalCommon with the addition of
// the ability to pass a context and additional request options.
//
// See SetVpnGatewayRenewalCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPN) SetVpnGatewayRenewalCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.SetVpnGatewayRenewalCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opSetVpnGatewayRenewal = "SetVpnGatewayRenewal"

// SetVpnGatewayRenewalRequest generates a "volcengine/request.Request" representing the
// client's request for the SetVpnGatewayRenewal operation. The "output" return
// value will be populated with the SetVpnGatewayRenewalCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned SetVpnGatewayRenewalCommon Request to send the API call to the service.
// the "output" return value is not valid until after SetVpnGatewayRenewalCommon Send returns without error.
//
// See SetVpnGatewayRenewal for more information on using the SetVpnGatewayRenewal
// API call, and error handling.
//
//	// Example sending a request using the SetVpnGatewayRenewalRequest method.
//	req, resp := client.SetVpnGatewayRenewalRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *VPN) SetVpnGatewayRenewalRequest(input *SetVpnGatewayRenewalInput) (req *request.Request, output *SetVpnGatewayRenewalOutput) {
	op := &request.Operation{
		Name:       opSetVpnGatewayRenewal,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &SetVpnGatewayRenewalInput{}
	}

	output = &SetVpnGatewayRenewalOutput{}
	req = c.newRequest(op, input, output)

	return
}

// SetVpnGatewayRenewal API operation for VPN.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VPN's
// API operation SetVpnGatewayRenewal for usage and error information.
func (c *VPN) SetVpnGatewayRenewal(input *SetVpnGatewayRenewalInput) (*SetVpnGatewayRenewalOutput, error) {
	req, out := c.SetVpnGatewayRenewalRequest(input)
	return out, req.Send()
}

// SetVpnGatewayRenewalWithContext is the same as SetVpnGatewayRenewal with the addition of
// the ability to pass a context and additional request options.
//
// See SetVpnGatewayRenewal for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPN) SetVpnGatewayRenewalWithContext(ctx volcengine.Context, input *SetVpnGatewayRenewalInput, opts ...request.Option) (*SetVpnGatewayRenewalOutput, error) {
	req, out := c.SetVpnGatewayRenewalRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type SetVpnGatewayRenewalInput struct {
	_ struct{} `type:"structure"`

	RemainRenewTimes *int64 `type:"integer"`

	RenewPeriod *int64 `type:"integer"`

	// RenewType is a required field
	RenewType *int64 `min:"1" max:"3" type:"integer" required:"true"`

	// VpnGatewayId is a required field
	VpnGatewayId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s SetVpnGatewayRenewalInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s SetVpnGatewayRenewalInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *SetVpnGatewayRenewalInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "SetVpnGatewayRenewalInput"}
	if s.RenewType == nil {
		invalidParams.Add(request.NewErrParamRequired("RenewType"))
	}
	if s.RenewType != nil && *s.RenewType < 1 {
		invalidParams.Add(request.NewErrParamMinValue("RenewType", 1))
	}
	if s.RenewType != nil && *s.RenewType > 3 {
		invalidParams.Add(request.NewErrParamMaxValue("RenewType", 3))
	}
	if s.VpnGatewayId == nil {
		invalidParams.Add(request.NewErrParamRequired("VpnGatewayId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetRemainRenewTimes sets the RemainRenewTimes field's value.
func (s *SetVpnGatewayRenewalInput) SetRemainRenewTimes(v int64) *SetVpnGatewayRenewalInput {
	s.RemainRenewTimes = &v
	return s
}

// SetRenewPeriod sets the RenewPeriod field's value.
func (s *SetVpnGatewayRenewalInput) SetRenewPeriod(v int64) *SetVpnGatewayRenewalInput {
	s.RenewPeriod = &v
	return s
}

// SetRenewType sets the RenewType field's value.
func (s *SetVpnGatewayRenewalInput) SetRenewType(v int64) *SetVpnGatewayRenewalInput {
	s.RenewType = &v
	return s
}

// SetVpnGatewayId sets the VpnGatewayId field's value.
func (s *SetVpnGatewayRenewalInput) SetVpnGatewayId(v string) *SetVpnGatewayRenewalInput {
	s.VpnGatewayId = &v
	return s
}

type SetVpnGatewayRenewalOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	RequestId *string `type:"string"`
}

// String returns the string representation
func (s SetVpnGatewayRenewalOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s SetVpnGatewayRenewalOutput) GoString() string {
	return s.String()
}

// SetRequestId sets the RequestId field's value.
func (s *SetVpnGatewayRenewalOutput) SetRequestId(v string) *SetVpnGatewayRenewalOutput {
	s.RequestId = &v
	return s
}

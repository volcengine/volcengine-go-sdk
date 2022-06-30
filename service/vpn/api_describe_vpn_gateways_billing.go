// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package vpn

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribeVpnGatewaysBillingCommon = "DescribeVpnGatewaysBilling"

// DescribeVpnGatewaysBillingCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeVpnGatewaysBillingCommon operation. The "output" return
// value will be populated with the DescribeVpnGatewaysBillingCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeVpnGatewaysBillingCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeVpnGatewaysBillingCommon Send returns without error.
//
// See DescribeVpnGatewaysBillingCommon for more information on using the DescribeVpnGatewaysBillingCommon
// API call, and error handling.
//
//    // Example sending a request using the DescribeVpnGatewaysBillingCommonRequest method.
//    req, resp := client.DescribeVpnGatewaysBillingCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VPN) DescribeVpnGatewaysBillingCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeVpnGatewaysBillingCommon,
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

// DescribeVpnGatewaysBillingCommon API operation for VPN.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VPN's
// API operation DescribeVpnGatewaysBillingCommon for usage and error information.
func (c *VPN) DescribeVpnGatewaysBillingCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeVpnGatewaysBillingCommonRequest(input)
	return out, req.Send()
}

// DescribeVpnGatewaysBillingCommonWithContext is the same as DescribeVpnGatewaysBillingCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeVpnGatewaysBillingCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPN) DescribeVpnGatewaysBillingCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeVpnGatewaysBillingCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeVpnGatewaysBilling = "DescribeVpnGatewaysBilling"

// DescribeVpnGatewaysBillingRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeVpnGatewaysBilling operation. The "output" return
// value will be populated with the DescribeVpnGatewaysBillingCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeVpnGatewaysBillingCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeVpnGatewaysBillingCommon Send returns without error.
//
// See DescribeVpnGatewaysBilling for more information on using the DescribeVpnGatewaysBilling
// API call, and error handling.
//
//    // Example sending a request using the DescribeVpnGatewaysBillingRequest method.
//    req, resp := client.DescribeVpnGatewaysBillingRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VPN) DescribeVpnGatewaysBillingRequest(input *DescribeVpnGatewaysBillingInput) (req *request.Request, output *DescribeVpnGatewaysBillingOutput) {
	op := &request.Operation{
		Name:       opDescribeVpnGatewaysBilling,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeVpnGatewaysBillingInput{}
	}

	output = &DescribeVpnGatewaysBillingOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DescribeVpnGatewaysBilling API operation for VPN.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VPN's
// API operation DescribeVpnGatewaysBilling for usage and error information.
func (c *VPN) DescribeVpnGatewaysBilling(input *DescribeVpnGatewaysBillingInput) (*DescribeVpnGatewaysBillingOutput, error) {
	req, out := c.DescribeVpnGatewaysBillingRequest(input)
	return out, req.Send()
}

// DescribeVpnGatewaysBillingWithContext is the same as DescribeVpnGatewaysBilling with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeVpnGatewaysBilling for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPN) DescribeVpnGatewaysBillingWithContext(ctx volcengine.Context, input *DescribeVpnGatewaysBillingInput, opts ...request.Option) (*DescribeVpnGatewaysBillingOutput, error) {
	req, out := c.DescribeVpnGatewaysBillingRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DescribeVpnGatewaysBillingInput struct {
	_ struct{} `type:"structure"`

	PageNumber *int64 `type:"integer"`

	PageSize *int64 `max:"100" type:"integer"`

	// VpnGatewayIds is a required field
	VpnGatewayIds []*string `type:"list" required:"true"`
}

// String returns the string representation
func (s DescribeVpnGatewaysBillingInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeVpnGatewaysBillingInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeVpnGatewaysBillingInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DescribeVpnGatewaysBillingInput"}
	if s.PageSize != nil && *s.PageSize > 100 {
		invalidParams.Add(request.NewErrParamMaxValue("PageSize", 100))
	}
	if s.VpnGatewayIds == nil {
		invalidParams.Add(request.NewErrParamRequired("VpnGatewayIds"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetPageNumber sets the PageNumber field's value.
func (s *DescribeVpnGatewaysBillingInput) SetPageNumber(v int64) *DescribeVpnGatewaysBillingInput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *DescribeVpnGatewaysBillingInput) SetPageSize(v int64) *DescribeVpnGatewaysBillingInput {
	s.PageSize = &v
	return s
}

// SetVpnGatewayIds sets the VpnGatewayIds field's value.
func (s *DescribeVpnGatewaysBillingInput) SetVpnGatewayIds(v []*string) *DescribeVpnGatewaysBillingInput {
	s.VpnGatewayIds = v
	return s
}

type DescribeVpnGatewaysBillingOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	PageNumber *int64 `type:"integer"`

	PageSize *int64 `type:"integer"`

	RequestId *string `type:"string"`

	TotalCount *int64 `type:"integer"`

	VpnGateways []*VpnGatewayForDescribeVpnGatewaysBillingOutput `type:"list"`
}

// String returns the string representation
func (s DescribeVpnGatewaysBillingOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeVpnGatewaysBillingOutput) GoString() string {
	return s.String()
}

// SetPageNumber sets the PageNumber field's value.
func (s *DescribeVpnGatewaysBillingOutput) SetPageNumber(v int64) *DescribeVpnGatewaysBillingOutput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *DescribeVpnGatewaysBillingOutput) SetPageSize(v int64) *DescribeVpnGatewaysBillingOutput {
	s.PageSize = &v
	return s
}

// SetRequestId sets the RequestId field's value.
func (s *DescribeVpnGatewaysBillingOutput) SetRequestId(v string) *DescribeVpnGatewaysBillingOutput {
	s.RequestId = &v
	return s
}

// SetTotalCount sets the TotalCount field's value.
func (s *DescribeVpnGatewaysBillingOutput) SetTotalCount(v int64) *DescribeVpnGatewaysBillingOutput {
	s.TotalCount = &v
	return s
}

// SetVpnGateways sets the VpnGateways field's value.
func (s *DescribeVpnGatewaysBillingOutput) SetVpnGateways(v []*VpnGatewayForDescribeVpnGatewaysBillingOutput) *DescribeVpnGatewaysBillingOutput {
	s.VpnGateways = v
	return s
}

type VpnGatewayForDescribeVpnGatewaysBillingOutput struct {
	_ struct{} `type:"structure"`

	BillingStatus *int64 `type:"integer"`

	BillingType *int64 `type:"integer"`

	ExpiredTime *string `type:"string"`

	ReclaimTime *string `type:"string"`

	RemainRenewTimes *int64 `type:"integer"`

	RenewType *int64 `type:"integer"`

	VpnGatewayId *string `type:"string"`
}

// String returns the string representation
func (s VpnGatewayForDescribeVpnGatewaysBillingOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s VpnGatewayForDescribeVpnGatewaysBillingOutput) GoString() string {
	return s.String()
}

// SetBillingStatus sets the BillingStatus field's value.
func (s *VpnGatewayForDescribeVpnGatewaysBillingOutput) SetBillingStatus(v int64) *VpnGatewayForDescribeVpnGatewaysBillingOutput {
	s.BillingStatus = &v
	return s
}

// SetBillingType sets the BillingType field's value.
func (s *VpnGatewayForDescribeVpnGatewaysBillingOutput) SetBillingType(v int64) *VpnGatewayForDescribeVpnGatewaysBillingOutput {
	s.BillingType = &v
	return s
}

// SetExpiredTime sets the ExpiredTime field's value.
func (s *VpnGatewayForDescribeVpnGatewaysBillingOutput) SetExpiredTime(v string) *VpnGatewayForDescribeVpnGatewaysBillingOutput {
	s.ExpiredTime = &v
	return s
}

// SetReclaimTime sets the ReclaimTime field's value.
func (s *VpnGatewayForDescribeVpnGatewaysBillingOutput) SetReclaimTime(v string) *VpnGatewayForDescribeVpnGatewaysBillingOutput {
	s.ReclaimTime = &v
	return s
}

// SetRemainRenewTimes sets the RemainRenewTimes field's value.
func (s *VpnGatewayForDescribeVpnGatewaysBillingOutput) SetRemainRenewTimes(v int64) *VpnGatewayForDescribeVpnGatewaysBillingOutput {
	s.RemainRenewTimes = &v
	return s
}

// SetRenewType sets the RenewType field's value.
func (s *VpnGatewayForDescribeVpnGatewaysBillingOutput) SetRenewType(v int64) *VpnGatewayForDescribeVpnGatewaysBillingOutput {
	s.RenewType = &v
	return s
}

// SetVpnGatewayId sets the VpnGatewayId field's value.
func (s *VpnGatewayForDescribeVpnGatewaysBillingOutput) SetVpnGatewayId(v string) *VpnGatewayForDescribeVpnGatewaysBillingOutput {
	s.VpnGatewayId = &v
	return s
}

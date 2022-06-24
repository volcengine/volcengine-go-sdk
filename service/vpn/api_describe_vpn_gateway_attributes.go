// Code generated by volcstack with private/model/cli/gen-api/main.go. DO NOT EDIT.

package vpn

import (
	"github.com/volcengine/volcstack-go-sdk/volcstack"
	"github.com/volcengine/volcstack-go-sdk/volcstack/request"
	"github.com/volcengine/volcstack-go-sdk/volcstack/response"
	"github.com/volcengine/volcstack-go-sdk/volcstack/volcstackutil"
)

const opDescribeVpnGatewayAttributesCommon = "DescribeVpnGatewayAttributes"

// DescribeVpnGatewayAttributesCommonRequest generates a "volcstack/request.Request" representing the
// client's request for the DescribeVpnGatewayAttributesCommon operation. The "output" return
// value will be populated with the DescribeVpnGatewayAttributesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeVpnGatewayAttributesCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeVpnGatewayAttributesCommon Send returns without error.
//
// See DescribeVpnGatewayAttributesCommon for more information on using the DescribeVpnGatewayAttributesCommon
// API call, and error handling.
//
//    // Example sending a request using the DescribeVpnGatewayAttributesCommonRequest method.
//    req, resp := client.DescribeVpnGatewayAttributesCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VPN) DescribeVpnGatewayAttributesCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeVpnGatewayAttributesCommon,
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

// DescribeVpnGatewayAttributesCommon API operation for VPN.
//
// Returns volcstackerr.Error for service API and SDK errors. Use runtime type assertions
// with volcstackerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCSTACK API reference guide for VPN's
// API operation DescribeVpnGatewayAttributesCommon for usage and error information.
func (c *VPN) DescribeVpnGatewayAttributesCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeVpnGatewayAttributesCommonRequest(input)
	return out, req.Send()
}

// DescribeVpnGatewayAttributesCommonWithContext is the same as DescribeVpnGatewayAttributesCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeVpnGatewayAttributesCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPN) DescribeVpnGatewayAttributesCommonWithContext(ctx volcstack.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeVpnGatewayAttributesCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeVpnGatewayAttributes = "DescribeVpnGatewayAttributes"

// DescribeVpnGatewayAttributesRequest generates a "volcstack/request.Request" representing the
// client's request for the DescribeVpnGatewayAttributes operation. The "output" return
// value will be populated with the DescribeVpnGatewayAttributesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeVpnGatewayAttributesCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeVpnGatewayAttributesCommon Send returns without error.
//
// See DescribeVpnGatewayAttributes for more information on using the DescribeVpnGatewayAttributes
// API call, and error handling.
//
//    // Example sending a request using the DescribeVpnGatewayAttributesRequest method.
//    req, resp := client.DescribeVpnGatewayAttributesRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VPN) DescribeVpnGatewayAttributesRequest(input *DescribeVpnGatewayAttributesInput) (req *request.Request, output *DescribeVpnGatewayAttributesOutput) {
	op := &request.Operation{
		Name:       opDescribeVpnGatewayAttributes,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeVpnGatewayAttributesInput{}
	}

	output = &DescribeVpnGatewayAttributesOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DescribeVpnGatewayAttributes API operation for VPN.
//
// Returns volcstackerr.Error for service API and SDK errors. Use runtime type assertions
// with volcstackerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCSTACK API reference guide for VPN's
// API operation DescribeVpnGatewayAttributes for usage and error information.
func (c *VPN) DescribeVpnGatewayAttributes(input *DescribeVpnGatewayAttributesInput) (*DescribeVpnGatewayAttributesOutput, error) {
	req, out := c.DescribeVpnGatewayAttributesRequest(input)
	return out, req.Send()
}

// DescribeVpnGatewayAttributesWithContext is the same as DescribeVpnGatewayAttributes with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeVpnGatewayAttributes for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPN) DescribeVpnGatewayAttributesWithContext(ctx volcstack.Context, input *DescribeVpnGatewayAttributesInput, opts ...request.Option) (*DescribeVpnGatewayAttributesOutput, error) {
	req, out := c.DescribeVpnGatewayAttributesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DescribeVpnGatewayAttributesInput struct {
	_ struct{} `type:"structure"`

	// VpnGatewayId is a required field
	VpnGatewayId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeVpnGatewayAttributesInput) String() string {
	return volcstackutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeVpnGatewayAttributesInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeVpnGatewayAttributesInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DescribeVpnGatewayAttributesInput"}
	if s.VpnGatewayId == nil {
		invalidParams.Add(request.NewErrParamRequired("VpnGatewayId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetVpnGatewayId sets the VpnGatewayId field's value.
func (s *DescribeVpnGatewayAttributesInput) SetVpnGatewayId(v string) *DescribeVpnGatewayAttributesInput {
	s.VpnGatewayId = &v
	return s
}

type DescribeVpnGatewayAttributesOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	AccountId *string `type:"string"`

	Bandwidth *int64 `type:"integer"`

	BillingType *int64 `type:"integer"`

	BusinessStatus *string `type:"string"`

	ConnectionCount *int64 `type:"integer"`

	CreationTime *string `type:"string"`

	DeletedTime *string `type:"string"`

	Description *string `type:"string"`

	ExpiredTime *string `type:"string"`

	IpAddress *string `type:"string"`

	LockReason *string `type:"string"`

	RequestId *string `type:"string"`

	RouteCount *int64 `type:"integer"`

	Status *string `type:"string"`

	SubnetId *string `type:"string"`

	UpdateTime *string `type:"string"`

	VpcId *string `type:"string"`

	VpnGatewayId *string `type:"string"`

	VpnGatewayName *string `type:"string"`
}

// String returns the string representation
func (s DescribeVpnGatewayAttributesOutput) String() string {
	return volcstackutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeVpnGatewayAttributesOutput) GoString() string {
	return s.String()
}

// SetAccountId sets the AccountId field's value.
func (s *DescribeVpnGatewayAttributesOutput) SetAccountId(v string) *DescribeVpnGatewayAttributesOutput {
	s.AccountId = &v
	return s
}

// SetBandwidth sets the Bandwidth field's value.
func (s *DescribeVpnGatewayAttributesOutput) SetBandwidth(v int64) *DescribeVpnGatewayAttributesOutput {
	s.Bandwidth = &v
	return s
}

// SetBillingType sets the BillingType field's value.
func (s *DescribeVpnGatewayAttributesOutput) SetBillingType(v int64) *DescribeVpnGatewayAttributesOutput {
	s.BillingType = &v
	return s
}

// SetBusinessStatus sets the BusinessStatus field's value.
func (s *DescribeVpnGatewayAttributesOutput) SetBusinessStatus(v string) *DescribeVpnGatewayAttributesOutput {
	s.BusinessStatus = &v
	return s
}

// SetConnectionCount sets the ConnectionCount field's value.
func (s *DescribeVpnGatewayAttributesOutput) SetConnectionCount(v int64) *DescribeVpnGatewayAttributesOutput {
	s.ConnectionCount = &v
	return s
}

// SetCreationTime sets the CreationTime field's value.
func (s *DescribeVpnGatewayAttributesOutput) SetCreationTime(v string) *DescribeVpnGatewayAttributesOutput {
	s.CreationTime = &v
	return s
}

// SetDeletedTime sets the DeletedTime field's value.
func (s *DescribeVpnGatewayAttributesOutput) SetDeletedTime(v string) *DescribeVpnGatewayAttributesOutput {
	s.DeletedTime = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *DescribeVpnGatewayAttributesOutput) SetDescription(v string) *DescribeVpnGatewayAttributesOutput {
	s.Description = &v
	return s
}

// SetExpiredTime sets the ExpiredTime field's value.
func (s *DescribeVpnGatewayAttributesOutput) SetExpiredTime(v string) *DescribeVpnGatewayAttributesOutput {
	s.ExpiredTime = &v
	return s
}

// SetIpAddress sets the IpAddress field's value.
func (s *DescribeVpnGatewayAttributesOutput) SetIpAddress(v string) *DescribeVpnGatewayAttributesOutput {
	s.IpAddress = &v
	return s
}

// SetLockReason sets the LockReason field's value.
func (s *DescribeVpnGatewayAttributesOutput) SetLockReason(v string) *DescribeVpnGatewayAttributesOutput {
	s.LockReason = &v
	return s
}

// SetRequestId sets the RequestId field's value.
func (s *DescribeVpnGatewayAttributesOutput) SetRequestId(v string) *DescribeVpnGatewayAttributesOutput {
	s.RequestId = &v
	return s
}

// SetRouteCount sets the RouteCount field's value.
func (s *DescribeVpnGatewayAttributesOutput) SetRouteCount(v int64) *DescribeVpnGatewayAttributesOutput {
	s.RouteCount = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *DescribeVpnGatewayAttributesOutput) SetStatus(v string) *DescribeVpnGatewayAttributesOutput {
	s.Status = &v
	return s
}

// SetSubnetId sets the SubnetId field's value.
func (s *DescribeVpnGatewayAttributesOutput) SetSubnetId(v string) *DescribeVpnGatewayAttributesOutput {
	s.SubnetId = &v
	return s
}

// SetUpdateTime sets the UpdateTime field's value.
func (s *DescribeVpnGatewayAttributesOutput) SetUpdateTime(v string) *DescribeVpnGatewayAttributesOutput {
	s.UpdateTime = &v
	return s
}

// SetVpcId sets the VpcId field's value.
func (s *DescribeVpnGatewayAttributesOutput) SetVpcId(v string) *DescribeVpnGatewayAttributesOutput {
	s.VpcId = &v
	return s
}

// SetVpnGatewayId sets the VpnGatewayId field's value.
func (s *DescribeVpnGatewayAttributesOutput) SetVpnGatewayId(v string) *DescribeVpnGatewayAttributesOutput {
	s.VpnGatewayId = &v
	return s
}

// SetVpnGatewayName sets the VpnGatewayName field's value.
func (s *DescribeVpnGatewayAttributesOutput) SetVpnGatewayName(v string) *DescribeVpnGatewayAttributesOutput {
	s.VpnGatewayName = &v
	return s
}
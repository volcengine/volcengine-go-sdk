// Code generated by volcstack with private/model/cli/gen-api/main.go. DO NOT EDIT.

package directconnect

import (
	"github.com/volcengine/volcstack-go-sdk/volcstack"
	"github.com/volcengine/volcstack-go-sdk/volcstack/request"
	"github.com/volcengine/volcstack-go-sdk/volcstack/response"
	"github.com/volcengine/volcstack-go-sdk/volcstack/volcstackutil"
)

const opDescribeDirectConnectVirtualInterfaceAttributesCommon = "DescribeDirectConnectVirtualInterfaceAttributes"

// DescribeDirectConnectVirtualInterfaceAttributesCommonRequest generates a "volcstack/request.Request" representing the
// client's request for the DescribeDirectConnectVirtualInterfaceAttributesCommon operation. The "output" return
// value will be populated with the DescribeDirectConnectVirtualInterfaceAttributesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeDirectConnectVirtualInterfaceAttributesCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeDirectConnectVirtualInterfaceAttributesCommon Send returns without error.
//
// See DescribeDirectConnectVirtualInterfaceAttributesCommon for more information on using the DescribeDirectConnectVirtualInterfaceAttributesCommon
// API call, and error handling.
//
//    // Example sending a request using the DescribeDirectConnectVirtualInterfaceAttributesCommonRequest method.
//    req, resp := client.DescribeDirectConnectVirtualInterfaceAttributesCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *DIRECTCONNECT) DescribeDirectConnectVirtualInterfaceAttributesCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeDirectConnectVirtualInterfaceAttributesCommon,
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

// DescribeDirectConnectVirtualInterfaceAttributesCommon API operation for DIRECTCONNECT.
//
// Returns volcstackerr.Error for service API and SDK errors. Use runtime type assertions
// with volcstackerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCSTACK API reference guide for DIRECTCONNECT's
// API operation DescribeDirectConnectVirtualInterfaceAttributesCommon for usage and error information.
func (c *DIRECTCONNECT) DescribeDirectConnectVirtualInterfaceAttributesCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeDirectConnectVirtualInterfaceAttributesCommonRequest(input)
	return out, req.Send()
}

// DescribeDirectConnectVirtualInterfaceAttributesCommonWithContext is the same as DescribeDirectConnectVirtualInterfaceAttributesCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeDirectConnectVirtualInterfaceAttributesCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *DIRECTCONNECT) DescribeDirectConnectVirtualInterfaceAttributesCommonWithContext(ctx volcstack.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeDirectConnectVirtualInterfaceAttributesCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeDirectConnectVirtualInterfaceAttributes = "DescribeDirectConnectVirtualInterfaceAttributes"

// DescribeDirectConnectVirtualInterfaceAttributesRequest generates a "volcstack/request.Request" representing the
// client's request for the DescribeDirectConnectVirtualInterfaceAttributes operation. The "output" return
// value will be populated with the DescribeDirectConnectVirtualInterfaceAttributesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeDirectConnectVirtualInterfaceAttributesCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeDirectConnectVirtualInterfaceAttributesCommon Send returns without error.
//
// See DescribeDirectConnectVirtualInterfaceAttributes for more information on using the DescribeDirectConnectVirtualInterfaceAttributes
// API call, and error handling.
//
//    // Example sending a request using the DescribeDirectConnectVirtualInterfaceAttributesRequest method.
//    req, resp := client.DescribeDirectConnectVirtualInterfaceAttributesRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *DIRECTCONNECT) DescribeDirectConnectVirtualInterfaceAttributesRequest(input *DescribeDirectConnectVirtualInterfaceAttributesInput) (req *request.Request, output *DescribeDirectConnectVirtualInterfaceAttributesOutput) {
	op := &request.Operation{
		Name:       opDescribeDirectConnectVirtualInterfaceAttributes,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeDirectConnectVirtualInterfaceAttributesInput{}
	}

	output = &DescribeDirectConnectVirtualInterfaceAttributesOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DescribeDirectConnectVirtualInterfaceAttributes API operation for DIRECTCONNECT.
//
// Returns volcstackerr.Error for service API and SDK errors. Use runtime type assertions
// with volcstackerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCSTACK API reference guide for DIRECTCONNECT's
// API operation DescribeDirectConnectVirtualInterfaceAttributes for usage and error information.
func (c *DIRECTCONNECT) DescribeDirectConnectVirtualInterfaceAttributes(input *DescribeDirectConnectVirtualInterfaceAttributesInput) (*DescribeDirectConnectVirtualInterfaceAttributesOutput, error) {
	req, out := c.DescribeDirectConnectVirtualInterfaceAttributesRequest(input)
	return out, req.Send()
}

// DescribeDirectConnectVirtualInterfaceAttributesWithContext is the same as DescribeDirectConnectVirtualInterfaceAttributes with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeDirectConnectVirtualInterfaceAttributes for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *DIRECTCONNECT) DescribeDirectConnectVirtualInterfaceAttributesWithContext(ctx volcstack.Context, input *DescribeDirectConnectVirtualInterfaceAttributesInput, opts ...request.Option) (*DescribeDirectConnectVirtualInterfaceAttributesOutput, error) {
	req, out := c.DescribeDirectConnectVirtualInterfaceAttributesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DescribeDirectConnectVirtualInterfaceAttributesInput struct {
	_ struct{} `type:"structure"`

	// VirtualInterfaceId is a required field
	VirtualInterfaceId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeDirectConnectVirtualInterfaceAttributesInput) String() string {
	return volcstackutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeDirectConnectVirtualInterfaceAttributesInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeDirectConnectVirtualInterfaceAttributesInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DescribeDirectConnectVirtualInterfaceAttributesInput"}
	if s.VirtualInterfaceId == nil {
		invalidParams.Add(request.NewErrParamRequired("VirtualInterfaceId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetVirtualInterfaceId sets the VirtualInterfaceId field's value.
func (s *DescribeDirectConnectVirtualInterfaceAttributesInput) SetVirtualInterfaceId(v string) *DescribeDirectConnectVirtualInterfaceAttributesInput {
	s.VirtualInterfaceId = &v
	return s
}

type DescribeDirectConnectVirtualInterfaceAttributesOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	AccountId *string `type:"string"`

	Bandwidth *int64 `type:"integer"`

	BfdDetectInterval *int64 `type:"integer"`

	BfdDetectMultiplier *int64 `type:"integer"`

	CreationTime *string `type:"string"`

	Description *string `type:"string"`

	DirectConnectConnectionId *string `type:"string"`

	DirectConnectGatewayId *string `type:"string"`

	EnableBfd *bool `type:"boolean"`

	LocalIp *string `type:"string"`

	PeerIp *string `type:"string"`

	RequestId *string `type:"string"`

	RouteType *string `type:"string"`

	Status *string `type:"string"`

	UpdateTime *string `type:"string"`

	VirtualInterfaceId *string `type:"string"`

	VirtualInterfaceName *string `type:"string"`

	VlanId *int64 `type:"integer"`
}

// String returns the string representation
func (s DescribeDirectConnectVirtualInterfaceAttributesOutput) String() string {
	return volcstackutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeDirectConnectVirtualInterfaceAttributesOutput) GoString() string {
	return s.String()
}

// SetAccountId sets the AccountId field's value.
func (s *DescribeDirectConnectVirtualInterfaceAttributesOutput) SetAccountId(v string) *DescribeDirectConnectVirtualInterfaceAttributesOutput {
	s.AccountId = &v
	return s
}

// SetBandwidth sets the Bandwidth field's value.
func (s *DescribeDirectConnectVirtualInterfaceAttributesOutput) SetBandwidth(v int64) *DescribeDirectConnectVirtualInterfaceAttributesOutput {
	s.Bandwidth = &v
	return s
}

// SetBfdDetectInterval sets the BfdDetectInterval field's value.
func (s *DescribeDirectConnectVirtualInterfaceAttributesOutput) SetBfdDetectInterval(v int64) *DescribeDirectConnectVirtualInterfaceAttributesOutput {
	s.BfdDetectInterval = &v
	return s
}

// SetBfdDetectMultiplier sets the BfdDetectMultiplier field's value.
func (s *DescribeDirectConnectVirtualInterfaceAttributesOutput) SetBfdDetectMultiplier(v int64) *DescribeDirectConnectVirtualInterfaceAttributesOutput {
	s.BfdDetectMultiplier = &v
	return s
}

// SetCreationTime sets the CreationTime field's value.
func (s *DescribeDirectConnectVirtualInterfaceAttributesOutput) SetCreationTime(v string) *DescribeDirectConnectVirtualInterfaceAttributesOutput {
	s.CreationTime = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *DescribeDirectConnectVirtualInterfaceAttributesOutput) SetDescription(v string) *DescribeDirectConnectVirtualInterfaceAttributesOutput {
	s.Description = &v
	return s
}

// SetDirectConnectConnectionId sets the DirectConnectConnectionId field's value.
func (s *DescribeDirectConnectVirtualInterfaceAttributesOutput) SetDirectConnectConnectionId(v string) *DescribeDirectConnectVirtualInterfaceAttributesOutput {
	s.DirectConnectConnectionId = &v
	return s
}

// SetDirectConnectGatewayId sets the DirectConnectGatewayId field's value.
func (s *DescribeDirectConnectVirtualInterfaceAttributesOutput) SetDirectConnectGatewayId(v string) *DescribeDirectConnectVirtualInterfaceAttributesOutput {
	s.DirectConnectGatewayId = &v
	return s
}

// SetEnableBfd sets the EnableBfd field's value.
func (s *DescribeDirectConnectVirtualInterfaceAttributesOutput) SetEnableBfd(v bool) *DescribeDirectConnectVirtualInterfaceAttributesOutput {
	s.EnableBfd = &v
	return s
}

// SetLocalIp sets the LocalIp field's value.
func (s *DescribeDirectConnectVirtualInterfaceAttributesOutput) SetLocalIp(v string) *DescribeDirectConnectVirtualInterfaceAttributesOutput {
	s.LocalIp = &v
	return s
}

// SetPeerIp sets the PeerIp field's value.
func (s *DescribeDirectConnectVirtualInterfaceAttributesOutput) SetPeerIp(v string) *DescribeDirectConnectVirtualInterfaceAttributesOutput {
	s.PeerIp = &v
	return s
}

// SetRequestId sets the RequestId field's value.
func (s *DescribeDirectConnectVirtualInterfaceAttributesOutput) SetRequestId(v string) *DescribeDirectConnectVirtualInterfaceAttributesOutput {
	s.RequestId = &v
	return s
}

// SetRouteType sets the RouteType field's value.
func (s *DescribeDirectConnectVirtualInterfaceAttributesOutput) SetRouteType(v string) *DescribeDirectConnectVirtualInterfaceAttributesOutput {
	s.RouteType = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *DescribeDirectConnectVirtualInterfaceAttributesOutput) SetStatus(v string) *DescribeDirectConnectVirtualInterfaceAttributesOutput {
	s.Status = &v
	return s
}

// SetUpdateTime sets the UpdateTime field's value.
func (s *DescribeDirectConnectVirtualInterfaceAttributesOutput) SetUpdateTime(v string) *DescribeDirectConnectVirtualInterfaceAttributesOutput {
	s.UpdateTime = &v
	return s
}

// SetVirtualInterfaceId sets the VirtualInterfaceId field's value.
func (s *DescribeDirectConnectVirtualInterfaceAttributesOutput) SetVirtualInterfaceId(v string) *DescribeDirectConnectVirtualInterfaceAttributesOutput {
	s.VirtualInterfaceId = &v
	return s
}

// SetVirtualInterfaceName sets the VirtualInterfaceName field's value.
func (s *DescribeDirectConnectVirtualInterfaceAttributesOutput) SetVirtualInterfaceName(v string) *DescribeDirectConnectVirtualInterfaceAttributesOutput {
	s.VirtualInterfaceName = &v
	return s
}

// SetVlanId sets the VlanId field's value.
func (s *DescribeDirectConnectVirtualInterfaceAttributesOutput) SetVlanId(v int64) *DescribeDirectConnectVirtualInterfaceAttributesOutput {
	s.VlanId = &v
	return s
}
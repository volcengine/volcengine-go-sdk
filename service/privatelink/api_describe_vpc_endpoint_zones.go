// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package privatelink

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribeVpcEndpointZonesCommon = "DescribeVpcEndpointZones"

// DescribeVpcEndpointZonesCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeVpcEndpointZonesCommon operation. The "output" return
// value will be populated with the DescribeVpcEndpointZonesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeVpcEndpointZonesCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeVpcEndpointZonesCommon Send returns without error.
//
// See DescribeVpcEndpointZonesCommon for more information on using the DescribeVpcEndpointZonesCommon
// API call, and error handling.
//
//    // Example sending a request using the DescribeVpcEndpointZonesCommonRequest method.
//    req, resp := client.DescribeVpcEndpointZonesCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *PRIVATELINK) DescribeVpcEndpointZonesCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeVpcEndpointZonesCommon,
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

// DescribeVpcEndpointZonesCommon API operation for PRIVATELINK.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for PRIVATELINK's
// API operation DescribeVpcEndpointZonesCommon for usage and error information.
func (c *PRIVATELINK) DescribeVpcEndpointZonesCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeVpcEndpointZonesCommonRequest(input)
	return out, req.Send()
}

// DescribeVpcEndpointZonesCommonWithContext is the same as DescribeVpcEndpointZonesCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeVpcEndpointZonesCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *PRIVATELINK) DescribeVpcEndpointZonesCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeVpcEndpointZonesCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeVpcEndpointZones = "DescribeVpcEndpointZones"

// DescribeVpcEndpointZonesRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeVpcEndpointZones operation. The "output" return
// value will be populated with the DescribeVpcEndpointZonesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeVpcEndpointZonesCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeVpcEndpointZonesCommon Send returns without error.
//
// See DescribeVpcEndpointZones for more information on using the DescribeVpcEndpointZones
// API call, and error handling.
//
//    // Example sending a request using the DescribeVpcEndpointZonesRequest method.
//    req, resp := client.DescribeVpcEndpointZonesRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *PRIVATELINK) DescribeVpcEndpointZonesRequest(input *DescribeVpcEndpointZonesInput) (req *request.Request, output *DescribeVpcEndpointZonesOutput) {
	op := &request.Operation{
		Name:       opDescribeVpcEndpointZones,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeVpcEndpointZonesInput{}
	}

	output = &DescribeVpcEndpointZonesOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DescribeVpcEndpointZones API operation for PRIVATELINK.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for PRIVATELINK's
// API operation DescribeVpcEndpointZones for usage and error information.
func (c *PRIVATELINK) DescribeVpcEndpointZones(input *DescribeVpcEndpointZonesInput) (*DescribeVpcEndpointZonesOutput, error) {
	req, out := c.DescribeVpcEndpointZonesRequest(input)
	return out, req.Send()
}

// DescribeVpcEndpointZonesWithContext is the same as DescribeVpcEndpointZones with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeVpcEndpointZones for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *PRIVATELINK) DescribeVpcEndpointZonesWithContext(ctx volcengine.Context, input *DescribeVpcEndpointZonesInput, opts ...request.Option) (*DescribeVpcEndpointZonesOutput, error) {
	req, out := c.DescribeVpcEndpointZonesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DescribeVpcEndpointZonesInput struct {
	_ struct{} `type:"structure"`

	// EndpointId is a required field
	EndpointId *string `type:"string" required:"true"`

	PageNumber *int32 `type:"int32"`

	PageSize *int32 `type:"int32"`
}

// String returns the string representation
func (s DescribeVpcEndpointZonesInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeVpcEndpointZonesInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeVpcEndpointZonesInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DescribeVpcEndpointZonesInput"}
	if s.EndpointId == nil {
		invalidParams.Add(request.NewErrParamRequired("EndpointId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetEndpointId sets the EndpointId field's value.
func (s *DescribeVpcEndpointZonesInput) SetEndpointId(v string) *DescribeVpcEndpointZonesInput {
	s.EndpointId = &v
	return s
}

// SetPageNumber sets the PageNumber field's value.
func (s *DescribeVpcEndpointZonesInput) SetPageNumber(v int32) *DescribeVpcEndpointZonesInput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *DescribeVpcEndpointZonesInput) SetPageSize(v int32) *DescribeVpcEndpointZonesInput {
	s.PageSize = &v
	return s
}

type DescribeVpcEndpointZonesOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	EndpointZones []*EndpointZoneForDescribeVpcEndpointZonesOutput `type:"list"`

	PageNumber *int32 `type:"int32"`

	PageSize *int32 `type:"int32"`

	RequestId *string `type:"string"`

	TotalCount *int32 `type:"int32"`
}

// String returns the string representation
func (s DescribeVpcEndpointZonesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeVpcEndpointZonesOutput) GoString() string {
	return s.String()
}

// SetEndpointZones sets the EndpointZones field's value.
func (s *DescribeVpcEndpointZonesOutput) SetEndpointZones(v []*EndpointZoneForDescribeVpcEndpointZonesOutput) *DescribeVpcEndpointZonesOutput {
	s.EndpointZones = v
	return s
}

// SetPageNumber sets the PageNumber field's value.
func (s *DescribeVpcEndpointZonesOutput) SetPageNumber(v int32) *DescribeVpcEndpointZonesOutput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *DescribeVpcEndpointZonesOutput) SetPageSize(v int32) *DescribeVpcEndpointZonesOutput {
	s.PageSize = &v
	return s
}

// SetRequestId sets the RequestId field's value.
func (s *DescribeVpcEndpointZonesOutput) SetRequestId(v string) *DescribeVpcEndpointZonesOutput {
	s.RequestId = &v
	return s
}

// SetTotalCount sets the TotalCount field's value.
func (s *DescribeVpcEndpointZonesOutput) SetTotalCount(v int32) *DescribeVpcEndpointZonesOutput {
	s.TotalCount = &v
	return s
}

type EndpointZoneForDescribeVpcEndpointZonesOutput struct {
	_ struct{} `type:"structure"`

	Ipv6ServiceStatus *string `type:"string"`

	NetworkInterfaceIP *string `type:"string"`

	NetworkInterfaceId *string `type:"string"`

	NetworkInterfaceIpv6 *string `type:"string"`

	ServiceStatus *string `type:"string"`

	SubnetId *string `type:"string"`

	ZoneDomain *string `type:"string"`

	ZoneId *string `type:"string"`

	ZoneStatus *string `type:"string"`
}

// String returns the string representation
func (s EndpointZoneForDescribeVpcEndpointZonesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s EndpointZoneForDescribeVpcEndpointZonesOutput) GoString() string {
	return s.String()
}

// SetIpv6ServiceStatus sets the Ipv6ServiceStatus field's value.
func (s *EndpointZoneForDescribeVpcEndpointZonesOutput) SetIpv6ServiceStatus(v string) *EndpointZoneForDescribeVpcEndpointZonesOutput {
	s.Ipv6ServiceStatus = &v
	return s
}

// SetNetworkInterfaceIP sets the NetworkInterfaceIP field's value.
func (s *EndpointZoneForDescribeVpcEndpointZonesOutput) SetNetworkInterfaceIP(v string) *EndpointZoneForDescribeVpcEndpointZonesOutput {
	s.NetworkInterfaceIP = &v
	return s
}

// SetNetworkInterfaceId sets the NetworkInterfaceId field's value.
func (s *EndpointZoneForDescribeVpcEndpointZonesOutput) SetNetworkInterfaceId(v string) *EndpointZoneForDescribeVpcEndpointZonesOutput {
	s.NetworkInterfaceId = &v
	return s
}

// SetNetworkInterfaceIpv6 sets the NetworkInterfaceIpv6 field's value.
func (s *EndpointZoneForDescribeVpcEndpointZonesOutput) SetNetworkInterfaceIpv6(v string) *EndpointZoneForDescribeVpcEndpointZonesOutput {
	s.NetworkInterfaceIpv6 = &v
	return s
}

// SetServiceStatus sets the ServiceStatus field's value.
func (s *EndpointZoneForDescribeVpcEndpointZonesOutput) SetServiceStatus(v string) *EndpointZoneForDescribeVpcEndpointZonesOutput {
	s.ServiceStatus = &v
	return s
}

// SetSubnetId sets the SubnetId field's value.
func (s *EndpointZoneForDescribeVpcEndpointZonesOutput) SetSubnetId(v string) *EndpointZoneForDescribeVpcEndpointZonesOutput {
	s.SubnetId = &v
	return s
}

// SetZoneDomain sets the ZoneDomain field's value.
func (s *EndpointZoneForDescribeVpcEndpointZonesOutput) SetZoneDomain(v string) *EndpointZoneForDescribeVpcEndpointZonesOutput {
	s.ZoneDomain = &v
	return s
}

// SetZoneId sets the ZoneId field's value.
func (s *EndpointZoneForDescribeVpcEndpointZonesOutput) SetZoneId(v string) *EndpointZoneForDescribeVpcEndpointZonesOutput {
	s.ZoneId = &v
	return s
}

// SetZoneStatus sets the ZoneStatus field's value.
func (s *EndpointZoneForDescribeVpcEndpointZonesOutput) SetZoneStatus(v string) *EndpointZoneForDescribeVpcEndpointZonesOutput {
	s.ZoneStatus = &v
	return s
}

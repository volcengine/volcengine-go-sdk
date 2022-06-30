// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package vpc

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribeSubnetsCommon = "DescribeSubnets"

// DescribeSubnetsCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeSubnetsCommon operation. The "output" return
// value will be populated with the DescribeSubnetsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeSubnetsCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeSubnetsCommon Send returns without error.
//
// See DescribeSubnetsCommon for more information on using the DescribeSubnetsCommon
// API call, and error handling.
//
//    // Example sending a request using the DescribeSubnetsCommonRequest method.
//    req, resp := client.DescribeSubnetsCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VPC) DescribeSubnetsCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeSubnetsCommon,
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

// DescribeSubnetsCommon API operation for VPC.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VPC's
// API operation DescribeSubnetsCommon for usage and error information.
func (c *VPC) DescribeSubnetsCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeSubnetsCommonRequest(input)
	return out, req.Send()
}

// DescribeSubnetsCommonWithContext is the same as DescribeSubnetsCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeSubnetsCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPC) DescribeSubnetsCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeSubnetsCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeSubnets = "DescribeSubnets"

// DescribeSubnetsRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeSubnets operation. The "output" return
// value will be populated with the DescribeSubnetsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeSubnetsCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeSubnetsCommon Send returns without error.
//
// See DescribeSubnets for more information on using the DescribeSubnets
// API call, and error handling.
//
//    // Example sending a request using the DescribeSubnetsRequest method.
//    req, resp := client.DescribeSubnetsRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VPC) DescribeSubnetsRequest(input *DescribeSubnetsInput) (req *request.Request, output *DescribeSubnetsOutput) {
	op := &request.Operation{
		Name:       opDescribeSubnets,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeSubnetsInput{}
	}

	output = &DescribeSubnetsOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DescribeSubnets API operation for VPC.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VPC's
// API operation DescribeSubnets for usage and error information.
func (c *VPC) DescribeSubnets(input *DescribeSubnetsInput) (*DescribeSubnetsOutput, error) {
	req, out := c.DescribeSubnetsRequest(input)
	return out, req.Send()
}

// DescribeSubnetsWithContext is the same as DescribeSubnets with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeSubnets for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPC) DescribeSubnetsWithContext(ctx volcengine.Context, input *DescribeSubnetsInput, opts ...request.Option) (*DescribeSubnetsOutput, error) {
	req, out := c.DescribeSubnetsRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DescribeSubnetsInput struct {
	_ struct{} `type:"structure"`

	PageNumber *int64 `type:"integer"`

	PageSize *int64 `max:"100" type:"integer"`

	RouteTableId *string `type:"string"`

	SubnetIds []*string `type:"list"`

	SubnetName *string `type:"string"`

	// VpcId is a required field
	VpcId *string `type:"string" required:"true"`

	ZoneId *string `type:"string"`
}

// String returns the string representation
func (s DescribeSubnetsInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeSubnetsInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeSubnetsInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DescribeSubnetsInput"}
	if s.PageSize != nil && *s.PageSize > 100 {
		invalidParams.Add(request.NewErrParamMaxValue("PageSize", 100))
	}
	if s.VpcId == nil {
		invalidParams.Add(request.NewErrParamRequired("VpcId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetPageNumber sets the PageNumber field's value.
func (s *DescribeSubnetsInput) SetPageNumber(v int64) *DescribeSubnetsInput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *DescribeSubnetsInput) SetPageSize(v int64) *DescribeSubnetsInput {
	s.PageSize = &v
	return s
}

// SetRouteTableId sets the RouteTableId field's value.
func (s *DescribeSubnetsInput) SetRouteTableId(v string) *DescribeSubnetsInput {
	s.RouteTableId = &v
	return s
}

// SetSubnetIds sets the SubnetIds field's value.
func (s *DescribeSubnetsInput) SetSubnetIds(v []*string) *DescribeSubnetsInput {
	s.SubnetIds = v
	return s
}

// SetSubnetName sets the SubnetName field's value.
func (s *DescribeSubnetsInput) SetSubnetName(v string) *DescribeSubnetsInput {
	s.SubnetName = &v
	return s
}

// SetVpcId sets the VpcId field's value.
func (s *DescribeSubnetsInput) SetVpcId(v string) *DescribeSubnetsInput {
	s.VpcId = &v
	return s
}

// SetZoneId sets the ZoneId field's value.
func (s *DescribeSubnetsInput) SetZoneId(v string) *DescribeSubnetsInput {
	s.ZoneId = &v
	return s
}

type DescribeSubnetsOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	PageNumber *int64 `type:"integer"`

	PageSize *int64 `type:"integer"`

	RequestId *string `type:"string"`

	Subnets []*SubnetForDescribeSubnetsOutput `type:"list"`

	TotalCount *int64 `type:"integer"`
}

// String returns the string representation
func (s DescribeSubnetsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeSubnetsOutput) GoString() string {
	return s.String()
}

// SetPageNumber sets the PageNumber field's value.
func (s *DescribeSubnetsOutput) SetPageNumber(v int64) *DescribeSubnetsOutput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *DescribeSubnetsOutput) SetPageSize(v int64) *DescribeSubnetsOutput {
	s.PageSize = &v
	return s
}

// SetRequestId sets the RequestId field's value.
func (s *DescribeSubnetsOutput) SetRequestId(v string) *DescribeSubnetsOutput {
	s.RequestId = &v
	return s
}

// SetSubnets sets the Subnets field's value.
func (s *DescribeSubnetsOutput) SetSubnets(v []*SubnetForDescribeSubnetsOutput) *DescribeSubnetsOutput {
	s.Subnets = v
	return s
}

// SetTotalCount sets the TotalCount field's value.
func (s *DescribeSubnetsOutput) SetTotalCount(v int64) *DescribeSubnetsOutput {
	s.TotalCount = &v
	return s
}

type RouteTableForDescribeSubnetsOutput struct {
	_ struct{} `type:"structure"`

	RouteTableId *string `type:"string"`

	RouteTableType *string `type:"string"`
}

// String returns the string representation
func (s RouteTableForDescribeSubnetsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s RouteTableForDescribeSubnetsOutput) GoString() string {
	return s.String()
}

// SetRouteTableId sets the RouteTableId field's value.
func (s *RouteTableForDescribeSubnetsOutput) SetRouteTableId(v string) *RouteTableForDescribeSubnetsOutput {
	s.RouteTableId = &v
	return s
}

// SetRouteTableType sets the RouteTableType field's value.
func (s *RouteTableForDescribeSubnetsOutput) SetRouteTableType(v string) *RouteTableForDescribeSubnetsOutput {
	s.RouteTableType = &v
	return s
}

type SubnetForDescribeSubnetsOutput struct {
	_ struct{} `type:"structure"`

	AccountId *string `type:"string"`

	AvailableIpAddressCount *int64 `type:"integer"`

	CidrBlock *string `type:"string"`

	CreationTime *string `type:"string"`

	Description *string `type:"string"`

	NetworkAclId *string `type:"string"`

	RouteTable *RouteTableForDescribeSubnetsOutput `type:"structure"`

	Status *string `type:"string"`

	SubnetId *string `type:"string"`

	SubnetName *string `type:"string"`

	TotalIpv4Count *int64 `type:"integer"`

	UpdateTime *string `type:"string"`

	VpcId *string `type:"string"`

	ZoneId *string `type:"string"`
}

// String returns the string representation
func (s SubnetForDescribeSubnetsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s SubnetForDescribeSubnetsOutput) GoString() string {
	return s.String()
}

// SetAccountId sets the AccountId field's value.
func (s *SubnetForDescribeSubnetsOutput) SetAccountId(v string) *SubnetForDescribeSubnetsOutput {
	s.AccountId = &v
	return s
}

// SetAvailableIpAddressCount sets the AvailableIpAddressCount field's value.
func (s *SubnetForDescribeSubnetsOutput) SetAvailableIpAddressCount(v int64) *SubnetForDescribeSubnetsOutput {
	s.AvailableIpAddressCount = &v
	return s
}

// SetCidrBlock sets the CidrBlock field's value.
func (s *SubnetForDescribeSubnetsOutput) SetCidrBlock(v string) *SubnetForDescribeSubnetsOutput {
	s.CidrBlock = &v
	return s
}

// SetCreationTime sets the CreationTime field's value.
func (s *SubnetForDescribeSubnetsOutput) SetCreationTime(v string) *SubnetForDescribeSubnetsOutput {
	s.CreationTime = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *SubnetForDescribeSubnetsOutput) SetDescription(v string) *SubnetForDescribeSubnetsOutput {
	s.Description = &v
	return s
}

// SetNetworkAclId sets the NetworkAclId field's value.
func (s *SubnetForDescribeSubnetsOutput) SetNetworkAclId(v string) *SubnetForDescribeSubnetsOutput {
	s.NetworkAclId = &v
	return s
}

// SetRouteTable sets the RouteTable field's value.
func (s *SubnetForDescribeSubnetsOutput) SetRouteTable(v *RouteTableForDescribeSubnetsOutput) *SubnetForDescribeSubnetsOutput {
	s.RouteTable = v
	return s
}

// SetStatus sets the Status field's value.
func (s *SubnetForDescribeSubnetsOutput) SetStatus(v string) *SubnetForDescribeSubnetsOutput {
	s.Status = &v
	return s
}

// SetSubnetId sets the SubnetId field's value.
func (s *SubnetForDescribeSubnetsOutput) SetSubnetId(v string) *SubnetForDescribeSubnetsOutput {
	s.SubnetId = &v
	return s
}

// SetSubnetName sets the SubnetName field's value.
func (s *SubnetForDescribeSubnetsOutput) SetSubnetName(v string) *SubnetForDescribeSubnetsOutput {
	s.SubnetName = &v
	return s
}

// SetTotalIpv4Count sets the TotalIpv4Count field's value.
func (s *SubnetForDescribeSubnetsOutput) SetTotalIpv4Count(v int64) *SubnetForDescribeSubnetsOutput {
	s.TotalIpv4Count = &v
	return s
}

// SetUpdateTime sets the UpdateTime field's value.
func (s *SubnetForDescribeSubnetsOutput) SetUpdateTime(v string) *SubnetForDescribeSubnetsOutput {
	s.UpdateTime = &v
	return s
}

// SetVpcId sets the VpcId field's value.
func (s *SubnetForDescribeSubnetsOutput) SetVpcId(v string) *SubnetForDescribeSubnetsOutput {
	s.VpcId = &v
	return s
}

// SetZoneId sets the ZoneId field's value.
func (s *SubnetForDescribeSubnetsOutput) SetZoneId(v string) *SubnetForDescribeSubnetsOutput {
	s.ZoneId = &v
	return s
}

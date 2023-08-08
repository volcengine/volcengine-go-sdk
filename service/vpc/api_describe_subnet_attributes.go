// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package vpc

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribeSubnetAttributesCommon = "DescribeSubnetAttributes"

// DescribeSubnetAttributesCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeSubnetAttributesCommon operation. The "output" return
// value will be populated with the DescribeSubnetAttributesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeSubnetAttributesCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeSubnetAttributesCommon Send returns without error.
//
// See DescribeSubnetAttributesCommon for more information on using the DescribeSubnetAttributesCommon
// API call, and error handling.
//
//	// Example sending a request using the DescribeSubnetAttributesCommonRequest method.
//	req, resp := client.DescribeSubnetAttributesCommonRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *VPC) DescribeSubnetAttributesCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeSubnetAttributesCommon,
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

// DescribeSubnetAttributesCommon API operation for VPC.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VPC's
// API operation DescribeSubnetAttributesCommon for usage and error information.
func (c *VPC) DescribeSubnetAttributesCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeSubnetAttributesCommonRequest(input)
	return out, req.Send()
}

// DescribeSubnetAttributesCommonWithContext is the same as DescribeSubnetAttributesCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeSubnetAttributesCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPC) DescribeSubnetAttributesCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeSubnetAttributesCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeSubnetAttributes = "DescribeSubnetAttributes"

// DescribeSubnetAttributesRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeSubnetAttributes operation. The "output" return
// value will be populated with the DescribeSubnetAttributesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeSubnetAttributesCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeSubnetAttributesCommon Send returns without error.
//
// See DescribeSubnetAttributes for more information on using the DescribeSubnetAttributes
// API call, and error handling.
//
//	// Example sending a request using the DescribeSubnetAttributesRequest method.
//	req, resp := client.DescribeSubnetAttributesRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *VPC) DescribeSubnetAttributesRequest(input *DescribeSubnetAttributesInput) (req *request.Request, output *DescribeSubnetAttributesOutput) {
	op := &request.Operation{
		Name:       opDescribeSubnetAttributes,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeSubnetAttributesInput{}
	}

	output = &DescribeSubnetAttributesOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DescribeSubnetAttributes API operation for VPC.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VPC's
// API operation DescribeSubnetAttributes for usage and error information.
func (c *VPC) DescribeSubnetAttributes(input *DescribeSubnetAttributesInput) (*DescribeSubnetAttributesOutput, error) {
	req, out := c.DescribeSubnetAttributesRequest(input)
	return out, req.Send()
}

// DescribeSubnetAttributesWithContext is the same as DescribeSubnetAttributes with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeSubnetAttributes for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPC) DescribeSubnetAttributesWithContext(ctx volcengine.Context, input *DescribeSubnetAttributesInput, opts ...request.Option) (*DescribeSubnetAttributesOutput, error) {
	req, out := c.DescribeSubnetAttributesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DescribeSubnetAttributesInput struct {
	_ struct{} `type:"structure"`

	// SubnetId is a required field
	SubnetId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeSubnetAttributesInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeSubnetAttributesInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeSubnetAttributesInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DescribeSubnetAttributesInput"}
	if s.SubnetId == nil {
		invalidParams.Add(request.NewErrParamRequired("SubnetId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetSubnetId sets the SubnetId field's value.
func (s *DescribeSubnetAttributesInput) SetSubnetId(v string) *DescribeSubnetAttributesInput {
	s.SubnetId = &v
	return s
}

type DescribeSubnetAttributesOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	AccountId *string `type:"string"`

	AvailableIpAddressCount *int64 `type:"integer"`

	CidrBlock *string `type:"string"`

	CreationTime *string `type:"string"`

	Description *string `type:"string"`

	Ipv6CidrBlock *string `type:"string"`

	IsDefault *bool `type:"boolean"`

	NetworkAclId *string `type:"string"`

	ProjectName *string `type:"string"`

	RequestId *string `type:"string"`

	RouteTable *RouteTableForDescribeSubnetAttributesOutput `type:"structure"`

	Status *string `type:"string"`

	SubnetId *string `type:"string"`

	SubnetName *string `type:"string"`

	TotalIpv4Count *int64 `type:"integer"`

	UpdateTime *string `type:"string"`

	VpcId *string `type:"string"`

	ZoneId *string `type:"string"`
}

// String returns the string representation
func (s DescribeSubnetAttributesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeSubnetAttributesOutput) GoString() string {
	return s.String()
}

// SetAccountId sets the AccountId field's value.
func (s *DescribeSubnetAttributesOutput) SetAccountId(v string) *DescribeSubnetAttributesOutput {
	s.AccountId = &v
	return s
}

// SetAvailableIpAddressCount sets the AvailableIpAddressCount field's value.
func (s *DescribeSubnetAttributesOutput) SetAvailableIpAddressCount(v int64) *DescribeSubnetAttributesOutput {
	s.AvailableIpAddressCount = &v
	return s
}

// SetCidrBlock sets the CidrBlock field's value.
func (s *DescribeSubnetAttributesOutput) SetCidrBlock(v string) *DescribeSubnetAttributesOutput {
	s.CidrBlock = &v
	return s
}

// SetCreationTime sets the CreationTime field's value.
func (s *DescribeSubnetAttributesOutput) SetCreationTime(v string) *DescribeSubnetAttributesOutput {
	s.CreationTime = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *DescribeSubnetAttributesOutput) SetDescription(v string) *DescribeSubnetAttributesOutput {
	s.Description = &v
	return s
}

// SetIpv6CidrBlock sets the Ipv6CidrBlock field's value.
func (s *DescribeSubnetAttributesOutput) SetIpv6CidrBlock(v string) *DescribeSubnetAttributesOutput {
	s.Ipv6CidrBlock = &v
	return s
}

// SetIsDefault sets the IsDefault field's value.
func (s *DescribeSubnetAttributesOutput) SetIsDefault(v bool) *DescribeSubnetAttributesOutput {
	s.IsDefault = &v
	return s
}

// SetNetworkAclId sets the NetworkAclId field's value.
func (s *DescribeSubnetAttributesOutput) SetNetworkAclId(v string) *DescribeSubnetAttributesOutput {
	s.NetworkAclId = &v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *DescribeSubnetAttributesOutput) SetProjectName(v string) *DescribeSubnetAttributesOutput {
	s.ProjectName = &v
	return s
}

// SetRequestId sets the RequestId field's value.
func (s *DescribeSubnetAttributesOutput) SetRequestId(v string) *DescribeSubnetAttributesOutput {
	s.RequestId = &v
	return s
}

// SetRouteTable sets the RouteTable field's value.
func (s *DescribeSubnetAttributesOutput) SetRouteTable(v *RouteTableForDescribeSubnetAttributesOutput) *DescribeSubnetAttributesOutput {
	s.RouteTable = v
	return s
}

// SetStatus sets the Status field's value.
func (s *DescribeSubnetAttributesOutput) SetStatus(v string) *DescribeSubnetAttributesOutput {
	s.Status = &v
	return s
}

// SetSubnetId sets the SubnetId field's value.
func (s *DescribeSubnetAttributesOutput) SetSubnetId(v string) *DescribeSubnetAttributesOutput {
	s.SubnetId = &v
	return s
}

// SetSubnetName sets the SubnetName field's value.
func (s *DescribeSubnetAttributesOutput) SetSubnetName(v string) *DescribeSubnetAttributesOutput {
	s.SubnetName = &v
	return s
}

// SetTotalIpv4Count sets the TotalIpv4Count field's value.
func (s *DescribeSubnetAttributesOutput) SetTotalIpv4Count(v int64) *DescribeSubnetAttributesOutput {
	s.TotalIpv4Count = &v
	return s
}

// SetUpdateTime sets the UpdateTime field's value.
func (s *DescribeSubnetAttributesOutput) SetUpdateTime(v string) *DescribeSubnetAttributesOutput {
	s.UpdateTime = &v
	return s
}

// SetVpcId sets the VpcId field's value.
func (s *DescribeSubnetAttributesOutput) SetVpcId(v string) *DescribeSubnetAttributesOutput {
	s.VpcId = &v
	return s
}

// SetZoneId sets the ZoneId field's value.
func (s *DescribeSubnetAttributesOutput) SetZoneId(v string) *DescribeSubnetAttributesOutput {
	s.ZoneId = &v
	return s
}

type RouteTableForDescribeSubnetAttributesOutput struct {
	_ struct{} `type:"structure"`

	RouteTableId *string `type:"string"`

	RouteTableType *string `type:"string"`
}

// String returns the string representation
func (s RouteTableForDescribeSubnetAttributesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s RouteTableForDescribeSubnetAttributesOutput) GoString() string {
	return s.String()
}

// SetRouteTableId sets the RouteTableId field's value.
func (s *RouteTableForDescribeSubnetAttributesOutput) SetRouteTableId(v string) *RouteTableForDescribeSubnetAttributesOutput {
	s.RouteTableId = &v
	return s
}

// SetRouteTableType sets the RouteTableType field's value.
func (s *RouteTableForDescribeSubnetAttributesOutput) SetRouteTableType(v string) *RouteTableForDescribeSubnetAttributesOutput {
	s.RouteTableType = &v
	return s
}

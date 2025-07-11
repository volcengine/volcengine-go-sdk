// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package clb

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribeNetworkLoadBalancersCommon = "DescribeNetworkLoadBalancers"

// DescribeNetworkLoadBalancersCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeNetworkLoadBalancersCommon operation. The "output" return
// value will be populated with the DescribeNetworkLoadBalancersCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeNetworkLoadBalancersCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeNetworkLoadBalancersCommon Send returns without error.
//
// See DescribeNetworkLoadBalancersCommon for more information on using the DescribeNetworkLoadBalancersCommon
// API call, and error handling.
//
//    // Example sending a request using the DescribeNetworkLoadBalancersCommonRequest method.
//    req, resp := client.DescribeNetworkLoadBalancersCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CLB) DescribeNetworkLoadBalancersCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeNetworkLoadBalancersCommon,
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

// DescribeNetworkLoadBalancersCommon API operation for CLB.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CLB's
// API operation DescribeNetworkLoadBalancersCommon for usage and error information.
func (c *CLB) DescribeNetworkLoadBalancersCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeNetworkLoadBalancersCommonRequest(input)
	return out, req.Send()
}

// DescribeNetworkLoadBalancersCommonWithContext is the same as DescribeNetworkLoadBalancersCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeNetworkLoadBalancersCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CLB) DescribeNetworkLoadBalancersCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeNetworkLoadBalancersCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeNetworkLoadBalancers = "DescribeNetworkLoadBalancers"

// DescribeNetworkLoadBalancersRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeNetworkLoadBalancers operation. The "output" return
// value will be populated with the DescribeNetworkLoadBalancersCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeNetworkLoadBalancersCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeNetworkLoadBalancersCommon Send returns without error.
//
// See DescribeNetworkLoadBalancers for more information on using the DescribeNetworkLoadBalancers
// API call, and error handling.
//
//    // Example sending a request using the DescribeNetworkLoadBalancersRequest method.
//    req, resp := client.DescribeNetworkLoadBalancersRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CLB) DescribeNetworkLoadBalancersRequest(input *DescribeNetworkLoadBalancersInput) (req *request.Request, output *DescribeNetworkLoadBalancersOutput) {
	op := &request.Operation{
		Name:       opDescribeNetworkLoadBalancers,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeNetworkLoadBalancersInput{}
	}

	output = &DescribeNetworkLoadBalancersOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DescribeNetworkLoadBalancers API operation for CLB.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CLB's
// API operation DescribeNetworkLoadBalancers for usage and error information.
func (c *CLB) DescribeNetworkLoadBalancers(input *DescribeNetworkLoadBalancersInput) (*DescribeNetworkLoadBalancersOutput, error) {
	req, out := c.DescribeNetworkLoadBalancersRequest(input)
	return out, req.Send()
}

// DescribeNetworkLoadBalancersWithContext is the same as DescribeNetworkLoadBalancers with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeNetworkLoadBalancers for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CLB) DescribeNetworkLoadBalancersWithContext(ctx volcengine.Context, input *DescribeNetworkLoadBalancersInput, opts ...request.Option) (*DescribeNetworkLoadBalancersOutput, error) {
	req, out := c.DescribeNetworkLoadBalancersRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DescribeNetworkLoadBalancersInput struct {
	_ struct{} `type:"structure"`

	IpAddressVersion *string `type:"string"`

	LoadBalancerIds []*string `type:"list"`

	LoadBalancerName *string `type:"string"`

	MaxResults *int64 `type:"integer"`

	NextToken *string `type:"string"`

	ProjectName *string `type:"string"`

	Status *string `type:"string"`

	TagFilters []*TagFilterForDescribeNetworkLoadBalancersInput `type:"list"`

	VpcId *string `type:"string"`

	ZoneId *string `type:"string"`
}

// String returns the string representation
func (s DescribeNetworkLoadBalancersInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeNetworkLoadBalancersInput) GoString() string {
	return s.String()
}

// SetIpAddressVersion sets the IpAddressVersion field's value.
func (s *DescribeNetworkLoadBalancersInput) SetIpAddressVersion(v string) *DescribeNetworkLoadBalancersInput {
	s.IpAddressVersion = &v
	return s
}

// SetLoadBalancerIds sets the LoadBalancerIds field's value.
func (s *DescribeNetworkLoadBalancersInput) SetLoadBalancerIds(v []*string) *DescribeNetworkLoadBalancersInput {
	s.LoadBalancerIds = v
	return s
}

// SetLoadBalancerName sets the LoadBalancerName field's value.
func (s *DescribeNetworkLoadBalancersInput) SetLoadBalancerName(v string) *DescribeNetworkLoadBalancersInput {
	s.LoadBalancerName = &v
	return s
}

// SetMaxResults sets the MaxResults field's value.
func (s *DescribeNetworkLoadBalancersInput) SetMaxResults(v int64) *DescribeNetworkLoadBalancersInput {
	s.MaxResults = &v
	return s
}

// SetNextToken sets the NextToken field's value.
func (s *DescribeNetworkLoadBalancersInput) SetNextToken(v string) *DescribeNetworkLoadBalancersInput {
	s.NextToken = &v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *DescribeNetworkLoadBalancersInput) SetProjectName(v string) *DescribeNetworkLoadBalancersInput {
	s.ProjectName = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *DescribeNetworkLoadBalancersInput) SetStatus(v string) *DescribeNetworkLoadBalancersInput {
	s.Status = &v
	return s
}

// SetTagFilters sets the TagFilters field's value.
func (s *DescribeNetworkLoadBalancersInput) SetTagFilters(v []*TagFilterForDescribeNetworkLoadBalancersInput) *DescribeNetworkLoadBalancersInput {
	s.TagFilters = v
	return s
}

// SetVpcId sets the VpcId field's value.
func (s *DescribeNetworkLoadBalancersInput) SetVpcId(v string) *DescribeNetworkLoadBalancersInput {
	s.VpcId = &v
	return s
}

// SetZoneId sets the ZoneId field's value.
func (s *DescribeNetworkLoadBalancersInput) SetZoneId(v string) *DescribeNetworkLoadBalancersInput {
	s.ZoneId = &v
	return s
}

type DescribeNetworkLoadBalancersOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	LoadBalancers []*LoadBalancerForDescribeNetworkLoadBalancersOutput `type:"list"`

	NextToken *string `type:"string"`

	RequestId *string `type:"string"`
}

// String returns the string representation
func (s DescribeNetworkLoadBalancersOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeNetworkLoadBalancersOutput) GoString() string {
	return s.String()
}

// SetLoadBalancers sets the LoadBalancers field's value.
func (s *DescribeNetworkLoadBalancersOutput) SetLoadBalancers(v []*LoadBalancerForDescribeNetworkLoadBalancersOutput) *DescribeNetworkLoadBalancersOutput {
	s.LoadBalancers = v
	return s
}

// SetNextToken sets the NextToken field's value.
func (s *DescribeNetworkLoadBalancersOutput) SetNextToken(v string) *DescribeNetworkLoadBalancersOutput {
	s.NextToken = &v
	return s
}

// SetRequestId sets the RequestId field's value.
func (s *DescribeNetworkLoadBalancersOutput) SetRequestId(v string) *DescribeNetworkLoadBalancersOutput {
	s.RequestId = &v
	return s
}

type LoadBalancerForDescribeNetworkLoadBalancersOutput struct {
	_ struct{} `type:"structure"`

	AccountId *string `type:"string"`

	BillingStatus *string `type:"string"`

	BillingType *int64 `type:"integer"`

	CreateTime *string `type:"string"`

	CrossZoneEnabled *bool `type:"boolean"`

	DNSName *string `type:"string"`

	Description *string `type:"string"`

	ExpectedOverdueTime *string `type:"string"`

	IpAddressVersion *string `type:"string"`

	Ipv4BandwidthPackageId *string `type:"string"`

	Ipv4NetworkType *string `type:"string"`

	Ipv6BandwidthPackageId *string `type:"string"`

	Ipv6NetworkType *string `type:"string"`

	LoadBalancerId *string `type:"string"`

	LoadBalancerName *string `type:"string"`

	ManagedSecurityGroupId *string `type:"string"`

	ModificationProtectionStatus *string `type:"string"`

	OverdueTime *string `type:"string"`

	ProjectName *string `type:"string"`

	ReclaimedTime *string `type:"string"`

	SecurityGroupIds []*string `type:"list"`

	Status *string `type:"string"`

	Tags []*TagForDescribeNetworkLoadBalancersOutput `type:"list"`

	UpdateTime *string `type:"string"`

	VpcId *string `type:"string"`

	ZoneMappings []*ZoneMappingForDescribeNetworkLoadBalancersOutput `type:"list"`
}

// String returns the string representation
func (s LoadBalancerForDescribeNetworkLoadBalancersOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s LoadBalancerForDescribeNetworkLoadBalancersOutput) GoString() string {
	return s.String()
}

// SetAccountId sets the AccountId field's value.
func (s *LoadBalancerForDescribeNetworkLoadBalancersOutput) SetAccountId(v string) *LoadBalancerForDescribeNetworkLoadBalancersOutput {
	s.AccountId = &v
	return s
}

// SetBillingStatus sets the BillingStatus field's value.
func (s *LoadBalancerForDescribeNetworkLoadBalancersOutput) SetBillingStatus(v string) *LoadBalancerForDescribeNetworkLoadBalancersOutput {
	s.BillingStatus = &v
	return s
}

// SetBillingType sets the BillingType field's value.
func (s *LoadBalancerForDescribeNetworkLoadBalancersOutput) SetBillingType(v int64) *LoadBalancerForDescribeNetworkLoadBalancersOutput {
	s.BillingType = &v
	return s
}

// SetCreateTime sets the CreateTime field's value.
func (s *LoadBalancerForDescribeNetworkLoadBalancersOutput) SetCreateTime(v string) *LoadBalancerForDescribeNetworkLoadBalancersOutput {
	s.CreateTime = &v
	return s
}

// SetCrossZoneEnabled sets the CrossZoneEnabled field's value.
func (s *LoadBalancerForDescribeNetworkLoadBalancersOutput) SetCrossZoneEnabled(v bool) *LoadBalancerForDescribeNetworkLoadBalancersOutput {
	s.CrossZoneEnabled = &v
	return s
}

// SetDNSName sets the DNSName field's value.
func (s *LoadBalancerForDescribeNetworkLoadBalancersOutput) SetDNSName(v string) *LoadBalancerForDescribeNetworkLoadBalancersOutput {
	s.DNSName = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *LoadBalancerForDescribeNetworkLoadBalancersOutput) SetDescription(v string) *LoadBalancerForDescribeNetworkLoadBalancersOutput {
	s.Description = &v
	return s
}

// SetExpectedOverdueTime sets the ExpectedOverdueTime field's value.
func (s *LoadBalancerForDescribeNetworkLoadBalancersOutput) SetExpectedOverdueTime(v string) *LoadBalancerForDescribeNetworkLoadBalancersOutput {
	s.ExpectedOverdueTime = &v
	return s
}

// SetIpAddressVersion sets the IpAddressVersion field's value.
func (s *LoadBalancerForDescribeNetworkLoadBalancersOutput) SetIpAddressVersion(v string) *LoadBalancerForDescribeNetworkLoadBalancersOutput {
	s.IpAddressVersion = &v
	return s
}

// SetIpv4BandwidthPackageId sets the Ipv4BandwidthPackageId field's value.
func (s *LoadBalancerForDescribeNetworkLoadBalancersOutput) SetIpv4BandwidthPackageId(v string) *LoadBalancerForDescribeNetworkLoadBalancersOutput {
	s.Ipv4BandwidthPackageId = &v
	return s
}

// SetIpv4NetworkType sets the Ipv4NetworkType field's value.
func (s *LoadBalancerForDescribeNetworkLoadBalancersOutput) SetIpv4NetworkType(v string) *LoadBalancerForDescribeNetworkLoadBalancersOutput {
	s.Ipv4NetworkType = &v
	return s
}

// SetIpv6BandwidthPackageId sets the Ipv6BandwidthPackageId field's value.
func (s *LoadBalancerForDescribeNetworkLoadBalancersOutput) SetIpv6BandwidthPackageId(v string) *LoadBalancerForDescribeNetworkLoadBalancersOutput {
	s.Ipv6BandwidthPackageId = &v
	return s
}

// SetIpv6NetworkType sets the Ipv6NetworkType field's value.
func (s *LoadBalancerForDescribeNetworkLoadBalancersOutput) SetIpv6NetworkType(v string) *LoadBalancerForDescribeNetworkLoadBalancersOutput {
	s.Ipv6NetworkType = &v
	return s
}

// SetLoadBalancerId sets the LoadBalancerId field's value.
func (s *LoadBalancerForDescribeNetworkLoadBalancersOutput) SetLoadBalancerId(v string) *LoadBalancerForDescribeNetworkLoadBalancersOutput {
	s.LoadBalancerId = &v
	return s
}

// SetLoadBalancerName sets the LoadBalancerName field's value.
func (s *LoadBalancerForDescribeNetworkLoadBalancersOutput) SetLoadBalancerName(v string) *LoadBalancerForDescribeNetworkLoadBalancersOutput {
	s.LoadBalancerName = &v
	return s
}

// SetManagedSecurityGroupId sets the ManagedSecurityGroupId field's value.
func (s *LoadBalancerForDescribeNetworkLoadBalancersOutput) SetManagedSecurityGroupId(v string) *LoadBalancerForDescribeNetworkLoadBalancersOutput {
	s.ManagedSecurityGroupId = &v
	return s
}

// SetModificationProtectionStatus sets the ModificationProtectionStatus field's value.
func (s *LoadBalancerForDescribeNetworkLoadBalancersOutput) SetModificationProtectionStatus(v string) *LoadBalancerForDescribeNetworkLoadBalancersOutput {
	s.ModificationProtectionStatus = &v
	return s
}

// SetOverdueTime sets the OverdueTime field's value.
func (s *LoadBalancerForDescribeNetworkLoadBalancersOutput) SetOverdueTime(v string) *LoadBalancerForDescribeNetworkLoadBalancersOutput {
	s.OverdueTime = &v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *LoadBalancerForDescribeNetworkLoadBalancersOutput) SetProjectName(v string) *LoadBalancerForDescribeNetworkLoadBalancersOutput {
	s.ProjectName = &v
	return s
}

// SetReclaimedTime sets the ReclaimedTime field's value.
func (s *LoadBalancerForDescribeNetworkLoadBalancersOutput) SetReclaimedTime(v string) *LoadBalancerForDescribeNetworkLoadBalancersOutput {
	s.ReclaimedTime = &v
	return s
}

// SetSecurityGroupIds sets the SecurityGroupIds field's value.
func (s *LoadBalancerForDescribeNetworkLoadBalancersOutput) SetSecurityGroupIds(v []*string) *LoadBalancerForDescribeNetworkLoadBalancersOutput {
	s.SecurityGroupIds = v
	return s
}

// SetStatus sets the Status field's value.
func (s *LoadBalancerForDescribeNetworkLoadBalancersOutput) SetStatus(v string) *LoadBalancerForDescribeNetworkLoadBalancersOutput {
	s.Status = &v
	return s
}

// SetTags sets the Tags field's value.
func (s *LoadBalancerForDescribeNetworkLoadBalancersOutput) SetTags(v []*TagForDescribeNetworkLoadBalancersOutput) *LoadBalancerForDescribeNetworkLoadBalancersOutput {
	s.Tags = v
	return s
}

// SetUpdateTime sets the UpdateTime field's value.
func (s *LoadBalancerForDescribeNetworkLoadBalancersOutput) SetUpdateTime(v string) *LoadBalancerForDescribeNetworkLoadBalancersOutput {
	s.UpdateTime = &v
	return s
}

// SetVpcId sets the VpcId field's value.
func (s *LoadBalancerForDescribeNetworkLoadBalancersOutput) SetVpcId(v string) *LoadBalancerForDescribeNetworkLoadBalancersOutput {
	s.VpcId = &v
	return s
}

// SetZoneMappings sets the ZoneMappings field's value.
func (s *LoadBalancerForDescribeNetworkLoadBalancersOutput) SetZoneMappings(v []*ZoneMappingForDescribeNetworkLoadBalancersOutput) *LoadBalancerForDescribeNetworkLoadBalancersOutput {
	s.ZoneMappings = v
	return s
}

type TagFilterForDescribeNetworkLoadBalancersInput struct {
	_ struct{} `type:"structure"`

	Key *string `type:"string"`

	Values []*string `type:"list"`
}

// String returns the string representation
func (s TagFilterForDescribeNetworkLoadBalancersInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s TagFilterForDescribeNetworkLoadBalancersInput) GoString() string {
	return s.String()
}

// SetKey sets the Key field's value.
func (s *TagFilterForDescribeNetworkLoadBalancersInput) SetKey(v string) *TagFilterForDescribeNetworkLoadBalancersInput {
	s.Key = &v
	return s
}

// SetValues sets the Values field's value.
func (s *TagFilterForDescribeNetworkLoadBalancersInput) SetValues(v []*string) *TagFilterForDescribeNetworkLoadBalancersInput {
	s.Values = v
	return s
}

type TagForDescribeNetworkLoadBalancersOutput struct {
	_ struct{} `type:"structure"`

	Key *string `type:"string"`

	Value *string `type:"string"`
}

// String returns the string representation
func (s TagForDescribeNetworkLoadBalancersOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s TagForDescribeNetworkLoadBalancersOutput) GoString() string {
	return s.String()
}

// SetKey sets the Key field's value.
func (s *TagForDescribeNetworkLoadBalancersOutput) SetKey(v string) *TagForDescribeNetworkLoadBalancersOutput {
	s.Key = &v
	return s
}

// SetValue sets the Value field's value.
func (s *TagForDescribeNetworkLoadBalancersOutput) SetValue(v string) *TagForDescribeNetworkLoadBalancersOutput {
	s.Value = &v
	return s
}

type ZoneMappingForDescribeNetworkLoadBalancersOutput struct {
	_ struct{} `type:"structure"`

	EniId *string `type:"string"`

	Ipv4Address *string `type:"string"`

	Ipv4EipAddress *string `type:"string"`

	Ipv4EipId *string `type:"string"`

	Ipv4HcStatus *string `type:"string"`

	Ipv6Address *string `type:"string"`

	Ipv6EipId *string `type:"string"`

	Ipv6HcStatus *string `type:"string"`

	SubnetId *string `type:"string"`

	ZoneId *string `type:"string"`
}

// String returns the string representation
func (s ZoneMappingForDescribeNetworkLoadBalancersOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ZoneMappingForDescribeNetworkLoadBalancersOutput) GoString() string {
	return s.String()
}

// SetEniId sets the EniId field's value.
func (s *ZoneMappingForDescribeNetworkLoadBalancersOutput) SetEniId(v string) *ZoneMappingForDescribeNetworkLoadBalancersOutput {
	s.EniId = &v
	return s
}

// SetIpv4Address sets the Ipv4Address field's value.
func (s *ZoneMappingForDescribeNetworkLoadBalancersOutput) SetIpv4Address(v string) *ZoneMappingForDescribeNetworkLoadBalancersOutput {
	s.Ipv4Address = &v
	return s
}

// SetIpv4EipAddress sets the Ipv4EipAddress field's value.
func (s *ZoneMappingForDescribeNetworkLoadBalancersOutput) SetIpv4EipAddress(v string) *ZoneMappingForDescribeNetworkLoadBalancersOutput {
	s.Ipv4EipAddress = &v
	return s
}

// SetIpv4EipId sets the Ipv4EipId field's value.
func (s *ZoneMappingForDescribeNetworkLoadBalancersOutput) SetIpv4EipId(v string) *ZoneMappingForDescribeNetworkLoadBalancersOutput {
	s.Ipv4EipId = &v
	return s
}

// SetIpv4HcStatus sets the Ipv4HcStatus field's value.
func (s *ZoneMappingForDescribeNetworkLoadBalancersOutput) SetIpv4HcStatus(v string) *ZoneMappingForDescribeNetworkLoadBalancersOutput {
	s.Ipv4HcStatus = &v
	return s
}

// SetIpv6Address sets the Ipv6Address field's value.
func (s *ZoneMappingForDescribeNetworkLoadBalancersOutput) SetIpv6Address(v string) *ZoneMappingForDescribeNetworkLoadBalancersOutput {
	s.Ipv6Address = &v
	return s
}

// SetIpv6EipId sets the Ipv6EipId field's value.
func (s *ZoneMappingForDescribeNetworkLoadBalancersOutput) SetIpv6EipId(v string) *ZoneMappingForDescribeNetworkLoadBalancersOutput {
	s.Ipv6EipId = &v
	return s
}

// SetIpv6HcStatus sets the Ipv6HcStatus field's value.
func (s *ZoneMappingForDescribeNetworkLoadBalancersOutput) SetIpv6HcStatus(v string) *ZoneMappingForDescribeNetworkLoadBalancersOutput {
	s.Ipv6HcStatus = &v
	return s
}

// SetSubnetId sets the SubnetId field's value.
func (s *ZoneMappingForDescribeNetworkLoadBalancersOutput) SetSubnetId(v string) *ZoneMappingForDescribeNetworkLoadBalancersOutput {
	s.SubnetId = &v
	return s
}

// SetZoneId sets the ZoneId field's value.
func (s *ZoneMappingForDescribeNetworkLoadBalancersOutput) SetZoneId(v string) *ZoneMappingForDescribeNetworkLoadBalancersOutput {
	s.ZoneId = &v
	return s
}

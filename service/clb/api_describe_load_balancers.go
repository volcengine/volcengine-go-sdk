// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package clb

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribeLoadBalancersCommon = "DescribeLoadBalancers"

// DescribeLoadBalancersCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeLoadBalancersCommon operation. The "output" return
// value will be populated with the DescribeLoadBalancersCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeLoadBalancersCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeLoadBalancersCommon Send returns without error.
//
// See DescribeLoadBalancersCommon for more information on using the DescribeLoadBalancersCommon
// API call, and error handling.
//
//	// Example sending a request using the DescribeLoadBalancersCommonRequest method.
//	req, resp := client.DescribeLoadBalancersCommonRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *CLB) DescribeLoadBalancersCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeLoadBalancersCommon,
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

// DescribeLoadBalancersCommon API operation for CLB.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CLB's
// API operation DescribeLoadBalancersCommon for usage and error information.
func (c *CLB) DescribeLoadBalancersCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeLoadBalancersCommonRequest(input)
	return out, req.Send()
}

// DescribeLoadBalancersCommonWithContext is the same as DescribeLoadBalancersCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeLoadBalancersCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CLB) DescribeLoadBalancersCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeLoadBalancersCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeLoadBalancers = "DescribeLoadBalancers"

// DescribeLoadBalancersRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeLoadBalancers operation. The "output" return
// value will be populated with the DescribeLoadBalancersCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeLoadBalancersCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeLoadBalancersCommon Send returns without error.
//
// See DescribeLoadBalancers for more information on using the DescribeLoadBalancers
// API call, and error handling.
//
//	// Example sending a request using the DescribeLoadBalancersRequest method.
//	req, resp := client.DescribeLoadBalancersRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *CLB) DescribeLoadBalancersRequest(input *DescribeLoadBalancersInput) (req *request.Request, output *DescribeLoadBalancersOutput) {
	op := &request.Operation{
		Name:       opDescribeLoadBalancers,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeLoadBalancersInput{}
	}

	output = &DescribeLoadBalancersOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DescribeLoadBalancers API operation for CLB.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CLB's
// API operation DescribeLoadBalancers for usage and error information.
func (c *CLB) DescribeLoadBalancers(input *DescribeLoadBalancersInput) (*DescribeLoadBalancersOutput, error) {
	req, out := c.DescribeLoadBalancersRequest(input)
	return out, req.Send()
}

// DescribeLoadBalancersWithContext is the same as DescribeLoadBalancers with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeLoadBalancers for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CLB) DescribeLoadBalancersWithContext(ctx volcengine.Context, input *DescribeLoadBalancersInput, opts ...request.Option) (*DescribeLoadBalancersOutput, error) {
	req, out := c.DescribeLoadBalancersRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DescribeLoadBalancersInput struct {
	_ struct{} `type:"structure"`

	AddressIpVersion *string `type:"string"`

	EipAddress *string `type:"string"`

	EniAddress *string `type:"string"`

	LoadBalancerIds []*string `type:"list"`

	LoadBalancerName *string `type:"string"`

	PageNumber *int64 `type:"integer"`

	PageSize *int64 `type:"integer"`

	ProjectName *string `type:"string"`

	TagFilters []*TagFilterForDescribeLoadBalancersInput `type:"list"`

	VpcId *string `type:"string"`
}

// String returns the string representation
func (s DescribeLoadBalancersInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeLoadBalancersInput) GoString() string {
	return s.String()
}

// SetAddressIpVersion sets the AddressIpVersion field's value.
func (s *DescribeLoadBalancersInput) SetAddressIpVersion(v string) *DescribeLoadBalancersInput {
	s.AddressIpVersion = &v
	return s
}

// SetEipAddress sets the EipAddress field's value.
func (s *DescribeLoadBalancersInput) SetEipAddress(v string) *DescribeLoadBalancersInput {
	s.EipAddress = &v
	return s
}

// SetEniAddress sets the EniAddress field's value.
func (s *DescribeLoadBalancersInput) SetEniAddress(v string) *DescribeLoadBalancersInput {
	s.EniAddress = &v
	return s
}

// SetLoadBalancerIds sets the LoadBalancerIds field's value.
func (s *DescribeLoadBalancersInput) SetLoadBalancerIds(v []*string) *DescribeLoadBalancersInput {
	s.LoadBalancerIds = v
	return s
}

// SetLoadBalancerName sets the LoadBalancerName field's value.
func (s *DescribeLoadBalancersInput) SetLoadBalancerName(v string) *DescribeLoadBalancersInput {
	s.LoadBalancerName = &v
	return s
}

// SetPageNumber sets the PageNumber field's value.
func (s *DescribeLoadBalancersInput) SetPageNumber(v int64) *DescribeLoadBalancersInput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *DescribeLoadBalancersInput) SetPageSize(v int64) *DescribeLoadBalancersInput {
	s.PageSize = &v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *DescribeLoadBalancersInput) SetProjectName(v string) *DescribeLoadBalancersInput {
	s.ProjectName = &v
	return s
}

// SetTagFilters sets the TagFilters field's value.
func (s *DescribeLoadBalancersInput) SetTagFilters(v []*TagFilterForDescribeLoadBalancersInput) *DescribeLoadBalancersInput {
	s.TagFilters = v
	return s
}

// SetVpcId sets the VpcId field's value.
func (s *DescribeLoadBalancersInput) SetVpcId(v string) *DescribeLoadBalancersInput {
	s.VpcId = &v
	return s
}

type DescribeLoadBalancersOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	LoadBalancers []*LoadBalancerForDescribeLoadBalancersOutput `type:"list"`

	PageNumber *int64 `type:"integer"`

	PageSize *int64 `type:"integer"`

	RequestId *string `type:"string"`

	TotalCount *int64 `type:"integer"`
}

// String returns the string representation
func (s DescribeLoadBalancersOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeLoadBalancersOutput) GoString() string {
	return s.String()
}

// SetLoadBalancers sets the LoadBalancers field's value.
func (s *DescribeLoadBalancersOutput) SetLoadBalancers(v []*LoadBalancerForDescribeLoadBalancersOutput) *DescribeLoadBalancersOutput {
	s.LoadBalancers = v
	return s
}

// SetPageNumber sets the PageNumber field's value.
func (s *DescribeLoadBalancersOutput) SetPageNumber(v int64) *DescribeLoadBalancersOutput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *DescribeLoadBalancersOutput) SetPageSize(v int64) *DescribeLoadBalancersOutput {
	s.PageSize = &v
	return s
}

// SetRequestId sets the RequestId field's value.
func (s *DescribeLoadBalancersOutput) SetRequestId(v string) *DescribeLoadBalancersOutput {
	s.RequestId = &v
	return s
}

// SetTotalCount sets the TotalCount field's value.
func (s *DescribeLoadBalancersOutput) SetTotalCount(v int64) *DescribeLoadBalancersOutput {
	s.TotalCount = &v
	return s
}

type LoadBalancerForDescribeLoadBalancersOutput struct {
	_ struct{} `type:"structure"`

	AddressIpVersion *string `type:"string"`

	BusinessStatus *string `type:"string"`

	CreateTime *string `type:"string"`

	DeletedTime *string `type:"string"`

	Description *string `type:"string"`

	EipAddress *string `type:"string"`

	EipID *string `type:"string"`

	EniAddress *string `type:"string"`

	EniID *string `type:"string"`

	EniIpv6Address *string `type:"string"`

	ExpiredTime *string `type:"string"`

	Ipv6EipId *string `type:"string"`

	LoadBalancerBillingType *int64 `type:"integer"`

	LoadBalancerId *string `type:"string"`

	LoadBalancerName *string `type:"string"`

	LoadBalancerSpec *string `type:"string"`

	LockReason *string `type:"string"`

	MasterZoneId *string `type:"string"`

	ModificationProtectionReason *string `type:"string"`

	ModificationProtectionStatus *string `type:"string"`

	OverdueTime *string `type:"string"`

	ProjectName *string `type:"string"`

	ServiceManaged *bool `type:"boolean"`

	SlaveZoneId *string `type:"string"`

	Status *string `type:"string"`

	SubnetId *string `type:"string"`

	Tags []*TagForDescribeLoadBalancersOutput `type:"list"`

	Type *string `type:"string"`

	UpdateTime *string `type:"string"`

	VpcId *string `type:"string"`
}

// String returns the string representation
func (s LoadBalancerForDescribeLoadBalancersOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s LoadBalancerForDescribeLoadBalancersOutput) GoString() string {
	return s.String()
}

// SetAddressIpVersion sets the AddressIpVersion field's value.
func (s *LoadBalancerForDescribeLoadBalancersOutput) SetAddressIpVersion(v string) *LoadBalancerForDescribeLoadBalancersOutput {
	s.AddressIpVersion = &v
	return s
}

// SetBusinessStatus sets the BusinessStatus field's value.
func (s *LoadBalancerForDescribeLoadBalancersOutput) SetBusinessStatus(v string) *LoadBalancerForDescribeLoadBalancersOutput {
	s.BusinessStatus = &v
	return s
}

// SetCreateTime sets the CreateTime field's value.
func (s *LoadBalancerForDescribeLoadBalancersOutput) SetCreateTime(v string) *LoadBalancerForDescribeLoadBalancersOutput {
	s.CreateTime = &v
	return s
}

// SetDeletedTime sets the DeletedTime field's value.
func (s *LoadBalancerForDescribeLoadBalancersOutput) SetDeletedTime(v string) *LoadBalancerForDescribeLoadBalancersOutput {
	s.DeletedTime = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *LoadBalancerForDescribeLoadBalancersOutput) SetDescription(v string) *LoadBalancerForDescribeLoadBalancersOutput {
	s.Description = &v
	return s
}

// SetEipAddress sets the EipAddress field's value.
func (s *LoadBalancerForDescribeLoadBalancersOutput) SetEipAddress(v string) *LoadBalancerForDescribeLoadBalancersOutput {
	s.EipAddress = &v
	return s
}

// SetEipID sets the EipID field's value.
func (s *LoadBalancerForDescribeLoadBalancersOutput) SetEipID(v string) *LoadBalancerForDescribeLoadBalancersOutput {
	s.EipID = &v
	return s
}

// SetEniAddress sets the EniAddress field's value.
func (s *LoadBalancerForDescribeLoadBalancersOutput) SetEniAddress(v string) *LoadBalancerForDescribeLoadBalancersOutput {
	s.EniAddress = &v
	return s
}

// SetEniID sets the EniID field's value.
func (s *LoadBalancerForDescribeLoadBalancersOutput) SetEniID(v string) *LoadBalancerForDescribeLoadBalancersOutput {
	s.EniID = &v
	return s
}

// SetEniIpv6Address sets the EniIpv6Address field's value.
func (s *LoadBalancerForDescribeLoadBalancersOutput) SetEniIpv6Address(v string) *LoadBalancerForDescribeLoadBalancersOutput {
	s.EniIpv6Address = &v
	return s
}

// SetExpiredTime sets the ExpiredTime field's value.
func (s *LoadBalancerForDescribeLoadBalancersOutput) SetExpiredTime(v string) *LoadBalancerForDescribeLoadBalancersOutput {
	s.ExpiredTime = &v
	return s
}

// SetIpv6EipId sets the Ipv6EipId field's value.
func (s *LoadBalancerForDescribeLoadBalancersOutput) SetIpv6EipId(v string) *LoadBalancerForDescribeLoadBalancersOutput {
	s.Ipv6EipId = &v
	return s
}

// SetLoadBalancerBillingType sets the LoadBalancerBillingType field's value.
func (s *LoadBalancerForDescribeLoadBalancersOutput) SetLoadBalancerBillingType(v int64) *LoadBalancerForDescribeLoadBalancersOutput {
	s.LoadBalancerBillingType = &v
	return s
}

// SetLoadBalancerId sets the LoadBalancerId field's value.
func (s *LoadBalancerForDescribeLoadBalancersOutput) SetLoadBalancerId(v string) *LoadBalancerForDescribeLoadBalancersOutput {
	s.LoadBalancerId = &v
	return s
}

// SetLoadBalancerName sets the LoadBalancerName field's value.
func (s *LoadBalancerForDescribeLoadBalancersOutput) SetLoadBalancerName(v string) *LoadBalancerForDescribeLoadBalancersOutput {
	s.LoadBalancerName = &v
	return s
}

// SetLoadBalancerSpec sets the LoadBalancerSpec field's value.
func (s *LoadBalancerForDescribeLoadBalancersOutput) SetLoadBalancerSpec(v string) *LoadBalancerForDescribeLoadBalancersOutput {
	s.LoadBalancerSpec = &v
	return s
}

// SetLockReason sets the LockReason field's value.
func (s *LoadBalancerForDescribeLoadBalancersOutput) SetLockReason(v string) *LoadBalancerForDescribeLoadBalancersOutput {
	s.LockReason = &v
	return s
}

// SetMasterZoneId sets the MasterZoneId field's value.
func (s *LoadBalancerForDescribeLoadBalancersOutput) SetMasterZoneId(v string) *LoadBalancerForDescribeLoadBalancersOutput {
	s.MasterZoneId = &v
	return s
}

// SetModificationProtectionReason sets the ModificationProtectionReason field's value.
func (s *LoadBalancerForDescribeLoadBalancersOutput) SetModificationProtectionReason(v string) *LoadBalancerForDescribeLoadBalancersOutput {
	s.ModificationProtectionReason = &v
	return s
}

// SetModificationProtectionStatus sets the ModificationProtectionStatus field's value.
func (s *LoadBalancerForDescribeLoadBalancersOutput) SetModificationProtectionStatus(v string) *LoadBalancerForDescribeLoadBalancersOutput {
	s.ModificationProtectionStatus = &v
	return s
}

// SetOverdueTime sets the OverdueTime field's value.
func (s *LoadBalancerForDescribeLoadBalancersOutput) SetOverdueTime(v string) *LoadBalancerForDescribeLoadBalancersOutput {
	s.OverdueTime = &v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *LoadBalancerForDescribeLoadBalancersOutput) SetProjectName(v string) *LoadBalancerForDescribeLoadBalancersOutput {
	s.ProjectName = &v
	return s
}

// SetServiceManaged sets the ServiceManaged field's value.
func (s *LoadBalancerForDescribeLoadBalancersOutput) SetServiceManaged(v bool) *LoadBalancerForDescribeLoadBalancersOutput {
	s.ServiceManaged = &v
	return s
}

// SetSlaveZoneId sets the SlaveZoneId field's value.
func (s *LoadBalancerForDescribeLoadBalancersOutput) SetSlaveZoneId(v string) *LoadBalancerForDescribeLoadBalancersOutput {
	s.SlaveZoneId = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *LoadBalancerForDescribeLoadBalancersOutput) SetStatus(v string) *LoadBalancerForDescribeLoadBalancersOutput {
	s.Status = &v
	return s
}

// SetSubnetId sets the SubnetId field's value.
func (s *LoadBalancerForDescribeLoadBalancersOutput) SetSubnetId(v string) *LoadBalancerForDescribeLoadBalancersOutput {
	s.SubnetId = &v
	return s
}

// SetTags sets the Tags field's value.
func (s *LoadBalancerForDescribeLoadBalancersOutput) SetTags(v []*TagForDescribeLoadBalancersOutput) *LoadBalancerForDescribeLoadBalancersOutput {
	s.Tags = v
	return s
}

// SetType sets the Type field's value.
func (s *LoadBalancerForDescribeLoadBalancersOutput) SetType(v string) *LoadBalancerForDescribeLoadBalancersOutput {
	s.Type = &v
	return s
}

// SetUpdateTime sets the UpdateTime field's value.
func (s *LoadBalancerForDescribeLoadBalancersOutput) SetUpdateTime(v string) *LoadBalancerForDescribeLoadBalancersOutput {
	s.UpdateTime = &v
	return s
}

// SetVpcId sets the VpcId field's value.
func (s *LoadBalancerForDescribeLoadBalancersOutput) SetVpcId(v string) *LoadBalancerForDescribeLoadBalancersOutput {
	s.VpcId = &v
	return s
}

type TagFilterForDescribeLoadBalancersInput struct {
	_ struct{} `type:"structure"`

	Key *string `type:"string"`

	Values []*string `type:"list"`
}

// String returns the string representation
func (s TagFilterForDescribeLoadBalancersInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s TagFilterForDescribeLoadBalancersInput) GoString() string {
	return s.String()
}

// SetKey sets the Key field's value.
func (s *TagFilterForDescribeLoadBalancersInput) SetKey(v string) *TagFilterForDescribeLoadBalancersInput {
	s.Key = &v
	return s
}

// SetValues sets the Values field's value.
func (s *TagFilterForDescribeLoadBalancersInput) SetValues(v []*string) *TagFilterForDescribeLoadBalancersInput {
	s.Values = v
	return s
}

type TagForDescribeLoadBalancersOutput struct {
	_ struct{} `type:"structure"`

	Key *string `type:"string"`

	Value *string `type:"string"`
}

// String returns the string representation
func (s TagForDescribeLoadBalancersOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s TagForDescribeLoadBalancersOutput) GoString() string {
	return s.String()
}

// SetKey sets the Key field's value.
func (s *TagForDescribeLoadBalancersOutput) SetKey(v string) *TagForDescribeLoadBalancersOutput {
	s.Key = &v
	return s
}

// SetValue sets the Value field's value.
func (s *TagForDescribeLoadBalancersOutput) SetValue(v string) *TagForDescribeLoadBalancersOutput {
	s.Value = &v
	return s
}

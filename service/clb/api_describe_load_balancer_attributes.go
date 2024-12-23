// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package clb

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribeLoadBalancerAttributesCommon = "DescribeLoadBalancerAttributes"

// DescribeLoadBalancerAttributesCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeLoadBalancerAttributesCommon operation. The "output" return
// value will be populated with the DescribeLoadBalancerAttributesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeLoadBalancerAttributesCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeLoadBalancerAttributesCommon Send returns without error.
//
// See DescribeLoadBalancerAttributesCommon for more information on using the DescribeLoadBalancerAttributesCommon
// API call, and error handling.
//
//    // Example sending a request using the DescribeLoadBalancerAttributesCommonRequest method.
//    req, resp := client.DescribeLoadBalancerAttributesCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CLB) DescribeLoadBalancerAttributesCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeLoadBalancerAttributesCommon,
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

// DescribeLoadBalancerAttributesCommon API operation for CLB.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CLB's
// API operation DescribeLoadBalancerAttributesCommon for usage and error information.
func (c *CLB) DescribeLoadBalancerAttributesCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeLoadBalancerAttributesCommonRequest(input)
	return out, req.Send()
}

// DescribeLoadBalancerAttributesCommonWithContext is the same as DescribeLoadBalancerAttributesCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeLoadBalancerAttributesCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CLB) DescribeLoadBalancerAttributesCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeLoadBalancerAttributesCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeLoadBalancerAttributes = "DescribeLoadBalancerAttributes"

// DescribeLoadBalancerAttributesRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeLoadBalancerAttributes operation. The "output" return
// value will be populated with the DescribeLoadBalancerAttributesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeLoadBalancerAttributesCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeLoadBalancerAttributesCommon Send returns without error.
//
// See DescribeLoadBalancerAttributes for more information on using the DescribeLoadBalancerAttributes
// API call, and error handling.
//
//    // Example sending a request using the DescribeLoadBalancerAttributesRequest method.
//    req, resp := client.DescribeLoadBalancerAttributesRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CLB) DescribeLoadBalancerAttributesRequest(input *DescribeLoadBalancerAttributesInput) (req *request.Request, output *DescribeLoadBalancerAttributesOutput) {
	op := &request.Operation{
		Name:       opDescribeLoadBalancerAttributes,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeLoadBalancerAttributesInput{}
	}

	output = &DescribeLoadBalancerAttributesOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DescribeLoadBalancerAttributes API operation for CLB.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CLB's
// API operation DescribeLoadBalancerAttributes for usage and error information.
func (c *CLB) DescribeLoadBalancerAttributes(input *DescribeLoadBalancerAttributesInput) (*DescribeLoadBalancerAttributesOutput, error) {
	req, out := c.DescribeLoadBalancerAttributesRequest(input)
	return out, req.Send()
}

// DescribeLoadBalancerAttributesWithContext is the same as DescribeLoadBalancerAttributes with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeLoadBalancerAttributes for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CLB) DescribeLoadBalancerAttributesWithContext(ctx volcengine.Context, input *DescribeLoadBalancerAttributesInput, opts ...request.Option) (*DescribeLoadBalancerAttributesOutput, error) {
	req, out := c.DescribeLoadBalancerAttributesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type AccessLogForDescribeLoadBalancerAttributesOutput struct {
	_ struct{} `type:"structure"`

	BucketName *string `type:"string"`

	Enabled *bool `type:"boolean"`
}

// String returns the string representation
func (s AccessLogForDescribeLoadBalancerAttributesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s AccessLogForDescribeLoadBalancerAttributesOutput) GoString() string {
	return s.String()
}

// SetBucketName sets the BucketName field's value.
func (s *AccessLogForDescribeLoadBalancerAttributesOutput) SetBucketName(v string) *AccessLogForDescribeLoadBalancerAttributesOutput {
	s.BucketName = &v
	return s
}

// SetEnabled sets the Enabled field's value.
func (s *AccessLogForDescribeLoadBalancerAttributesOutput) SetEnabled(v bool) *AccessLogForDescribeLoadBalancerAttributesOutput {
	s.Enabled = &v
	return s
}

type DescribeLoadBalancerAttributesInput struct {
	_ struct{} `type:"structure"`

	// LoadBalancerId is a required field
	LoadBalancerId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeLoadBalancerAttributesInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeLoadBalancerAttributesInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeLoadBalancerAttributesInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DescribeLoadBalancerAttributesInput"}
	if s.LoadBalancerId == nil {
		invalidParams.Add(request.NewErrParamRequired("LoadBalancerId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetLoadBalancerId sets the LoadBalancerId field's value.
func (s *DescribeLoadBalancerAttributesInput) SetLoadBalancerId(v string) *DescribeLoadBalancerAttributesInput {
	s.LoadBalancerId = &v
	return s
}

type DescribeLoadBalancerAttributesOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	AccessLog *AccessLogForDescribeLoadBalancerAttributesOutput `type:"structure"`

	AccountId *string `type:"string"`

	AddressIpVersion *string `type:"string"`

	AllowedPorts []*string `type:"list"`

	BusinessStatus *string `type:"string"`

	BypassSecurityGroupEnabled *string `type:"string"`

	CreateTime *string `type:"string"`

	DeletedTime *string `type:"string"`

	Description *string `type:"string"`

	Eip *EipForDescribeLoadBalancerAttributesOutput `type:"structure"`

	EipAddress *string `type:"string"`

	EipID *string `type:"string"`

	Enabled *bool `type:"boolean"`

	EniAddress *string `type:"string"`

	EniAddressNum *int64 `type:"integer"`

	EniAddresses []*EniAddressForDescribeLoadBalancerAttributesOutput `type:"list"`

	EniID *string `type:"string"`

	EniIpv6Address *string `type:"string"`

	ExclusiveClusterId *string `type:"string"`

	ExpiredTime *string `type:"string"`

	Ipv6AddressBandwidth *Ipv6AddressBandwidthForDescribeLoadBalancerAttributesOutput `type:"structure"`

	Ipv6EipId *string `type:"string"`

	Listeners []*ListenerForDescribeLoadBalancerAttributesOutput `type:"list"`

	LoadBalancerBillingType *int64 `type:"integer"`

	LoadBalancerId *string `type:"string"`

	LoadBalancerName *string `type:"string"`

	LoadBalancerSpec *string `type:"string"`

	LockReason *string `type:"string"`

	LogTopicId *string `type:"string"`

	MasterZoneId *string `type:"string"`

	ModificationProtectionReason *string `type:"string"`

	ModificationProtectionStatus *string `type:"string"`

	NewArch *bool `type:"boolean"`

	OverdueTime *string `type:"string"`

	ProjectName *string `type:"string"`

	RequestId *string `type:"string"`

	ServerGroups []*ServerGroupForDescribeLoadBalancerAttributesOutput `type:"list"`

	ServiceManaged *bool `type:"boolean"`

	SlaveZoneId *string `type:"string"`

	Status *string `type:"string"`

	SubnetId *string `type:"string"`

	Tags []*TagForDescribeLoadBalancerAttributesOutput `type:"list"`

	TimestampRemoveEnabled *string `type:"string"`

	Type *string `type:"string"`

	UpdateTime *string `type:"string"`

	VpcId *string `type:"string"`
}

// String returns the string representation
func (s DescribeLoadBalancerAttributesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeLoadBalancerAttributesOutput) GoString() string {
	return s.String()
}

// SetAccessLog sets the AccessLog field's value.
func (s *DescribeLoadBalancerAttributesOutput) SetAccessLog(v *AccessLogForDescribeLoadBalancerAttributesOutput) *DescribeLoadBalancerAttributesOutput {
	s.AccessLog = v
	return s
}

// SetAccountId sets the AccountId field's value.
func (s *DescribeLoadBalancerAttributesOutput) SetAccountId(v string) *DescribeLoadBalancerAttributesOutput {
	s.AccountId = &v
	return s
}

// SetAddressIpVersion sets the AddressIpVersion field's value.
func (s *DescribeLoadBalancerAttributesOutput) SetAddressIpVersion(v string) *DescribeLoadBalancerAttributesOutput {
	s.AddressIpVersion = &v
	return s
}

// SetAllowedPorts sets the AllowedPorts field's value.
func (s *DescribeLoadBalancerAttributesOutput) SetAllowedPorts(v []*string) *DescribeLoadBalancerAttributesOutput {
	s.AllowedPorts = v
	return s
}

// SetBusinessStatus sets the BusinessStatus field's value.
func (s *DescribeLoadBalancerAttributesOutput) SetBusinessStatus(v string) *DescribeLoadBalancerAttributesOutput {
	s.BusinessStatus = &v
	return s
}

// SetBypassSecurityGroupEnabled sets the BypassSecurityGroupEnabled field's value.
func (s *DescribeLoadBalancerAttributesOutput) SetBypassSecurityGroupEnabled(v string) *DescribeLoadBalancerAttributesOutput {
	s.BypassSecurityGroupEnabled = &v
	return s
}

// SetCreateTime sets the CreateTime field's value.
func (s *DescribeLoadBalancerAttributesOutput) SetCreateTime(v string) *DescribeLoadBalancerAttributesOutput {
	s.CreateTime = &v
	return s
}

// SetDeletedTime sets the DeletedTime field's value.
func (s *DescribeLoadBalancerAttributesOutput) SetDeletedTime(v string) *DescribeLoadBalancerAttributesOutput {
	s.DeletedTime = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *DescribeLoadBalancerAttributesOutput) SetDescription(v string) *DescribeLoadBalancerAttributesOutput {
	s.Description = &v
	return s
}

// SetEip sets the Eip field's value.
func (s *DescribeLoadBalancerAttributesOutput) SetEip(v *EipForDescribeLoadBalancerAttributesOutput) *DescribeLoadBalancerAttributesOutput {
	s.Eip = v
	return s
}

// SetEipAddress sets the EipAddress field's value.
func (s *DescribeLoadBalancerAttributesOutput) SetEipAddress(v string) *DescribeLoadBalancerAttributesOutput {
	s.EipAddress = &v
	return s
}

// SetEipID sets the EipID field's value.
func (s *DescribeLoadBalancerAttributesOutput) SetEipID(v string) *DescribeLoadBalancerAttributesOutput {
	s.EipID = &v
	return s
}

// SetEnabled sets the Enabled field's value.
func (s *DescribeLoadBalancerAttributesOutput) SetEnabled(v bool) *DescribeLoadBalancerAttributesOutput {
	s.Enabled = &v
	return s
}

// SetEniAddress sets the EniAddress field's value.
func (s *DescribeLoadBalancerAttributesOutput) SetEniAddress(v string) *DescribeLoadBalancerAttributesOutput {
	s.EniAddress = &v
	return s
}

// SetEniAddressNum sets the EniAddressNum field's value.
func (s *DescribeLoadBalancerAttributesOutput) SetEniAddressNum(v int64) *DescribeLoadBalancerAttributesOutput {
	s.EniAddressNum = &v
	return s
}

// SetEniAddresses sets the EniAddresses field's value.
func (s *DescribeLoadBalancerAttributesOutput) SetEniAddresses(v []*EniAddressForDescribeLoadBalancerAttributesOutput) *DescribeLoadBalancerAttributesOutput {
	s.EniAddresses = v
	return s
}

// SetEniID sets the EniID field's value.
func (s *DescribeLoadBalancerAttributesOutput) SetEniID(v string) *DescribeLoadBalancerAttributesOutput {
	s.EniID = &v
	return s
}

// SetEniIpv6Address sets the EniIpv6Address field's value.
func (s *DescribeLoadBalancerAttributesOutput) SetEniIpv6Address(v string) *DescribeLoadBalancerAttributesOutput {
	s.EniIpv6Address = &v
	return s
}

// SetExclusiveClusterId sets the ExclusiveClusterId field's value.
func (s *DescribeLoadBalancerAttributesOutput) SetExclusiveClusterId(v string) *DescribeLoadBalancerAttributesOutput {
	s.ExclusiveClusterId = &v
	return s
}

// SetExpiredTime sets the ExpiredTime field's value.
func (s *DescribeLoadBalancerAttributesOutput) SetExpiredTime(v string) *DescribeLoadBalancerAttributesOutput {
	s.ExpiredTime = &v
	return s
}

// SetIpv6AddressBandwidth sets the Ipv6AddressBandwidth field's value.
func (s *DescribeLoadBalancerAttributesOutput) SetIpv6AddressBandwidth(v *Ipv6AddressBandwidthForDescribeLoadBalancerAttributesOutput) *DescribeLoadBalancerAttributesOutput {
	s.Ipv6AddressBandwidth = v
	return s
}

// SetIpv6EipId sets the Ipv6EipId field's value.
func (s *DescribeLoadBalancerAttributesOutput) SetIpv6EipId(v string) *DescribeLoadBalancerAttributesOutput {
	s.Ipv6EipId = &v
	return s
}

// SetListeners sets the Listeners field's value.
func (s *DescribeLoadBalancerAttributesOutput) SetListeners(v []*ListenerForDescribeLoadBalancerAttributesOutput) *DescribeLoadBalancerAttributesOutput {
	s.Listeners = v
	return s
}

// SetLoadBalancerBillingType sets the LoadBalancerBillingType field's value.
func (s *DescribeLoadBalancerAttributesOutput) SetLoadBalancerBillingType(v int64) *DescribeLoadBalancerAttributesOutput {
	s.LoadBalancerBillingType = &v
	return s
}

// SetLoadBalancerId sets the LoadBalancerId field's value.
func (s *DescribeLoadBalancerAttributesOutput) SetLoadBalancerId(v string) *DescribeLoadBalancerAttributesOutput {
	s.LoadBalancerId = &v
	return s
}

// SetLoadBalancerName sets the LoadBalancerName field's value.
func (s *DescribeLoadBalancerAttributesOutput) SetLoadBalancerName(v string) *DescribeLoadBalancerAttributesOutput {
	s.LoadBalancerName = &v
	return s
}

// SetLoadBalancerSpec sets the LoadBalancerSpec field's value.
func (s *DescribeLoadBalancerAttributesOutput) SetLoadBalancerSpec(v string) *DescribeLoadBalancerAttributesOutput {
	s.LoadBalancerSpec = &v
	return s
}

// SetLockReason sets the LockReason field's value.
func (s *DescribeLoadBalancerAttributesOutput) SetLockReason(v string) *DescribeLoadBalancerAttributesOutput {
	s.LockReason = &v
	return s
}

// SetLogTopicId sets the LogTopicId field's value.
func (s *DescribeLoadBalancerAttributesOutput) SetLogTopicId(v string) *DescribeLoadBalancerAttributesOutput {
	s.LogTopicId = &v
	return s
}

// SetMasterZoneId sets the MasterZoneId field's value.
func (s *DescribeLoadBalancerAttributesOutput) SetMasterZoneId(v string) *DescribeLoadBalancerAttributesOutput {
	s.MasterZoneId = &v
	return s
}

// SetModificationProtectionReason sets the ModificationProtectionReason field's value.
func (s *DescribeLoadBalancerAttributesOutput) SetModificationProtectionReason(v string) *DescribeLoadBalancerAttributesOutput {
	s.ModificationProtectionReason = &v
	return s
}

// SetModificationProtectionStatus sets the ModificationProtectionStatus field's value.
func (s *DescribeLoadBalancerAttributesOutput) SetModificationProtectionStatus(v string) *DescribeLoadBalancerAttributesOutput {
	s.ModificationProtectionStatus = &v
	return s
}

// SetNewArch sets the NewArch field's value.
func (s *DescribeLoadBalancerAttributesOutput) SetNewArch(v bool) *DescribeLoadBalancerAttributesOutput {
	s.NewArch = &v
	return s
}

// SetOverdueTime sets the OverdueTime field's value.
func (s *DescribeLoadBalancerAttributesOutput) SetOverdueTime(v string) *DescribeLoadBalancerAttributesOutput {
	s.OverdueTime = &v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *DescribeLoadBalancerAttributesOutput) SetProjectName(v string) *DescribeLoadBalancerAttributesOutput {
	s.ProjectName = &v
	return s
}

// SetRequestId sets the RequestId field's value.
func (s *DescribeLoadBalancerAttributesOutput) SetRequestId(v string) *DescribeLoadBalancerAttributesOutput {
	s.RequestId = &v
	return s
}

// SetServerGroups sets the ServerGroups field's value.
func (s *DescribeLoadBalancerAttributesOutput) SetServerGroups(v []*ServerGroupForDescribeLoadBalancerAttributesOutput) *DescribeLoadBalancerAttributesOutput {
	s.ServerGroups = v
	return s
}

// SetServiceManaged sets the ServiceManaged field's value.
func (s *DescribeLoadBalancerAttributesOutput) SetServiceManaged(v bool) *DescribeLoadBalancerAttributesOutput {
	s.ServiceManaged = &v
	return s
}

// SetSlaveZoneId sets the SlaveZoneId field's value.
func (s *DescribeLoadBalancerAttributesOutput) SetSlaveZoneId(v string) *DescribeLoadBalancerAttributesOutput {
	s.SlaveZoneId = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *DescribeLoadBalancerAttributesOutput) SetStatus(v string) *DescribeLoadBalancerAttributesOutput {
	s.Status = &v
	return s
}

// SetSubnetId sets the SubnetId field's value.
func (s *DescribeLoadBalancerAttributesOutput) SetSubnetId(v string) *DescribeLoadBalancerAttributesOutput {
	s.SubnetId = &v
	return s
}

// SetTags sets the Tags field's value.
func (s *DescribeLoadBalancerAttributesOutput) SetTags(v []*TagForDescribeLoadBalancerAttributesOutput) *DescribeLoadBalancerAttributesOutput {
	s.Tags = v
	return s
}

// SetTimestampRemoveEnabled sets the TimestampRemoveEnabled field's value.
func (s *DescribeLoadBalancerAttributesOutput) SetTimestampRemoveEnabled(v string) *DescribeLoadBalancerAttributesOutput {
	s.TimestampRemoveEnabled = &v
	return s
}

// SetType sets the Type field's value.
func (s *DescribeLoadBalancerAttributesOutput) SetType(v string) *DescribeLoadBalancerAttributesOutput {
	s.Type = &v
	return s
}

// SetUpdateTime sets the UpdateTime field's value.
func (s *DescribeLoadBalancerAttributesOutput) SetUpdateTime(v string) *DescribeLoadBalancerAttributesOutput {
	s.UpdateTime = &v
	return s
}

// SetVpcId sets the VpcId field's value.
func (s *DescribeLoadBalancerAttributesOutput) SetVpcId(v string) *DescribeLoadBalancerAttributesOutput {
	s.VpcId = &v
	return s
}

type EipForDescribeLoadBalancerAttributesOutput struct {
	_ struct{} `type:"structure"`

	Bandwidth *int64 `type:"integer"`

	BandwidthPackageId *string `type:"string"`

	EipAddress *string `type:"string"`

	EipBillingType *int64 `type:"integer"`

	ISP *string `type:"string"`

	SecurityProtectionTypes []*string `type:"list"`
}

// String returns the string representation
func (s EipForDescribeLoadBalancerAttributesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s EipForDescribeLoadBalancerAttributesOutput) GoString() string {
	return s.String()
}

// SetBandwidth sets the Bandwidth field's value.
func (s *EipForDescribeLoadBalancerAttributesOutput) SetBandwidth(v int64) *EipForDescribeLoadBalancerAttributesOutput {
	s.Bandwidth = &v
	return s
}

// SetBandwidthPackageId sets the BandwidthPackageId field's value.
func (s *EipForDescribeLoadBalancerAttributesOutput) SetBandwidthPackageId(v string) *EipForDescribeLoadBalancerAttributesOutput {
	s.BandwidthPackageId = &v
	return s
}

// SetEipAddress sets the EipAddress field's value.
func (s *EipForDescribeLoadBalancerAttributesOutput) SetEipAddress(v string) *EipForDescribeLoadBalancerAttributesOutput {
	s.EipAddress = &v
	return s
}

// SetEipBillingType sets the EipBillingType field's value.
func (s *EipForDescribeLoadBalancerAttributesOutput) SetEipBillingType(v int64) *EipForDescribeLoadBalancerAttributesOutput {
	s.EipBillingType = &v
	return s
}

// SetISP sets the ISP field's value.
func (s *EipForDescribeLoadBalancerAttributesOutput) SetISP(v string) *EipForDescribeLoadBalancerAttributesOutput {
	s.ISP = &v
	return s
}

// SetSecurityProtectionTypes sets the SecurityProtectionTypes field's value.
func (s *EipForDescribeLoadBalancerAttributesOutput) SetSecurityProtectionTypes(v []*string) *EipForDescribeLoadBalancerAttributesOutput {
	s.SecurityProtectionTypes = v
	return s
}

type EniAddressForDescribeLoadBalancerAttributesOutput struct {
	_ struct{} `type:"structure"`

	EipAddress *string `type:"string"`

	EipId *string `type:"string"`

	EniAddress *string `type:"string"`
}

// String returns the string representation
func (s EniAddressForDescribeLoadBalancerAttributesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s EniAddressForDescribeLoadBalancerAttributesOutput) GoString() string {
	return s.String()
}

// SetEipAddress sets the EipAddress field's value.
func (s *EniAddressForDescribeLoadBalancerAttributesOutput) SetEipAddress(v string) *EniAddressForDescribeLoadBalancerAttributesOutput {
	s.EipAddress = &v
	return s
}

// SetEipId sets the EipId field's value.
func (s *EniAddressForDescribeLoadBalancerAttributesOutput) SetEipId(v string) *EniAddressForDescribeLoadBalancerAttributesOutput {
	s.EipId = &v
	return s
}

// SetEniAddress sets the EniAddress field's value.
func (s *EniAddressForDescribeLoadBalancerAttributesOutput) SetEniAddress(v string) *EniAddressForDescribeLoadBalancerAttributesOutput {
	s.EniAddress = &v
	return s
}

type Ipv6AddressBandwidthForDescribeLoadBalancerAttributesOutput struct {
	_ struct{} `type:"structure"`

	Bandwidth *int64 `type:"integer"`

	BandwidthPackageId *string `type:"string"`

	BillingType *int64 `type:"integer"`

	ISP *string `type:"string"`

	NetworkType *string `type:"string"`
}

// String returns the string representation
func (s Ipv6AddressBandwidthForDescribeLoadBalancerAttributesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s Ipv6AddressBandwidthForDescribeLoadBalancerAttributesOutput) GoString() string {
	return s.String()
}

// SetBandwidth sets the Bandwidth field's value.
func (s *Ipv6AddressBandwidthForDescribeLoadBalancerAttributesOutput) SetBandwidth(v int64) *Ipv6AddressBandwidthForDescribeLoadBalancerAttributesOutput {
	s.Bandwidth = &v
	return s
}

// SetBandwidthPackageId sets the BandwidthPackageId field's value.
func (s *Ipv6AddressBandwidthForDescribeLoadBalancerAttributesOutput) SetBandwidthPackageId(v string) *Ipv6AddressBandwidthForDescribeLoadBalancerAttributesOutput {
	s.BandwidthPackageId = &v
	return s
}

// SetBillingType sets the BillingType field's value.
func (s *Ipv6AddressBandwidthForDescribeLoadBalancerAttributesOutput) SetBillingType(v int64) *Ipv6AddressBandwidthForDescribeLoadBalancerAttributesOutput {
	s.BillingType = &v
	return s
}

// SetISP sets the ISP field's value.
func (s *Ipv6AddressBandwidthForDescribeLoadBalancerAttributesOutput) SetISP(v string) *Ipv6AddressBandwidthForDescribeLoadBalancerAttributesOutput {
	s.ISP = &v
	return s
}

// SetNetworkType sets the NetworkType field's value.
func (s *Ipv6AddressBandwidthForDescribeLoadBalancerAttributesOutput) SetNetworkType(v string) *Ipv6AddressBandwidthForDescribeLoadBalancerAttributesOutput {
	s.NetworkType = &v
	return s
}

type ListenerForDescribeLoadBalancerAttributesOutput struct {
	_ struct{} `type:"structure"`

	ListenerId *string `type:"string"`

	ListenerName *string `type:"string"`
}

// String returns the string representation
func (s ListenerForDescribeLoadBalancerAttributesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListenerForDescribeLoadBalancerAttributesOutput) GoString() string {
	return s.String()
}

// SetListenerId sets the ListenerId field's value.
func (s *ListenerForDescribeLoadBalancerAttributesOutput) SetListenerId(v string) *ListenerForDescribeLoadBalancerAttributesOutput {
	s.ListenerId = &v
	return s
}

// SetListenerName sets the ListenerName field's value.
func (s *ListenerForDescribeLoadBalancerAttributesOutput) SetListenerName(v string) *ListenerForDescribeLoadBalancerAttributesOutput {
	s.ListenerName = &v
	return s
}

type ServerGroupForDescribeLoadBalancerAttributesOutput struct {
	_ struct{} `type:"structure"`

	ServerGroupId *string `type:"string"`

	ServerGroupName *string `type:"string"`
}

// String returns the string representation
func (s ServerGroupForDescribeLoadBalancerAttributesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ServerGroupForDescribeLoadBalancerAttributesOutput) GoString() string {
	return s.String()
}

// SetServerGroupId sets the ServerGroupId field's value.
func (s *ServerGroupForDescribeLoadBalancerAttributesOutput) SetServerGroupId(v string) *ServerGroupForDescribeLoadBalancerAttributesOutput {
	s.ServerGroupId = &v
	return s
}

// SetServerGroupName sets the ServerGroupName field's value.
func (s *ServerGroupForDescribeLoadBalancerAttributesOutput) SetServerGroupName(v string) *ServerGroupForDescribeLoadBalancerAttributesOutput {
	s.ServerGroupName = &v
	return s
}

type TagForDescribeLoadBalancerAttributesOutput struct {
	_ struct{} `type:"structure"`

	Key *string `type:"string"`

	Value *string `type:"string"`
}

// String returns the string representation
func (s TagForDescribeLoadBalancerAttributesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s TagForDescribeLoadBalancerAttributesOutput) GoString() string {
	return s.String()
}

// SetKey sets the Key field's value.
func (s *TagForDescribeLoadBalancerAttributesOutput) SetKey(v string) *TagForDescribeLoadBalancerAttributesOutput {
	s.Key = &v
	return s
}

// SetValue sets the Value field's value.
func (s *TagForDescribeLoadBalancerAttributesOutput) SetValue(v string) *TagForDescribeLoadBalancerAttributesOutput {
	s.Value = &v
	return s
}

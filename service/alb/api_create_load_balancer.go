// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package alb

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opCreateLoadBalancerCommon = "CreateLoadBalancer"

// CreateLoadBalancerCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateLoadBalancerCommon operation. The "output" return
// value will be populated with the CreateLoadBalancerCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateLoadBalancerCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateLoadBalancerCommon Send returns without error.
//
// See CreateLoadBalancerCommon for more information on using the CreateLoadBalancerCommon
// API call, and error handling.
//
//    // Example sending a request using the CreateLoadBalancerCommonRequest method.
//    req, resp := client.CreateLoadBalancerCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ALB) CreateLoadBalancerCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opCreateLoadBalancerCommon,
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

// CreateLoadBalancerCommon API operation for ALB.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ALB's
// API operation CreateLoadBalancerCommon for usage and error information.
func (c *ALB) CreateLoadBalancerCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.CreateLoadBalancerCommonRequest(input)
	return out, req.Send()
}

// CreateLoadBalancerCommonWithContext is the same as CreateLoadBalancerCommon with the addition of
// the ability to pass a context and additional request options.
//
// See CreateLoadBalancerCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ALB) CreateLoadBalancerCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.CreateLoadBalancerCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opCreateLoadBalancer = "CreateLoadBalancer"

// CreateLoadBalancerRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateLoadBalancer operation. The "output" return
// value will be populated with the CreateLoadBalancerCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateLoadBalancerCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateLoadBalancerCommon Send returns without error.
//
// See CreateLoadBalancer for more information on using the CreateLoadBalancer
// API call, and error handling.
//
//    // Example sending a request using the CreateLoadBalancerRequest method.
//    req, resp := client.CreateLoadBalancerRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ALB) CreateLoadBalancerRequest(input *CreateLoadBalancerInput) (req *request.Request, output *CreateLoadBalancerOutput) {
	op := &request.Operation{
		Name:       opCreateLoadBalancer,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateLoadBalancerInput{}
	}

	output = &CreateLoadBalancerOutput{}
	req = c.newRequest(op, input, output)

	return
}

// CreateLoadBalancer API operation for ALB.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ALB's
// API operation CreateLoadBalancer for usage and error information.
func (c *ALB) CreateLoadBalancer(input *CreateLoadBalancerInput) (*CreateLoadBalancerOutput, error) {
	req, out := c.CreateLoadBalancerRequest(input)
	return out, req.Send()
}

// CreateLoadBalancerWithContext is the same as CreateLoadBalancer with the addition of
// the ability to pass a context and additional request options.
//
// See CreateLoadBalancer for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ALB) CreateLoadBalancerWithContext(ctx volcengine.Context, input *CreateLoadBalancerInput, opts ...request.Option) (*CreateLoadBalancerOutput, error) {
	req, out := c.CreateLoadBalancerRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type CreateLoadBalancerInput struct {
	_ struct{} `type:"structure"`

	AddressIpVersion *string `type:"string"`

	BandwidthPackageId *string `type:"string"`

	DeleteProtection *string `type:"string"`

	Description *string `type:"string"`

	EipBillingConfig *EipBillingConfigForCreateLoadBalancerInput `type:"structure"`

	Ipv6EipBillingConfig *Ipv6EipBillingConfigForCreateLoadBalancerInput `type:"structure"`

	LoadBalancerBillingType *int64 `type:"integer"`

	LoadBalancerName *string `type:"string"`

	ModificationProtectionReason *string `type:"string"`

	ModificationProtectionStatus *string `type:"string"`

	ProjectName *string `type:"string"`

	RegionId *string `type:"string"`

	SubnetId *string `type:"string"`

	Tags []*TagForCreateLoadBalancerInput `type:"list"`

	Type *string `type:"string"`

	VpcId *string `type:"string"`

	ZoneMappings []*ZoneMappingForCreateLoadBalancerInput `type:"list"`
}

// String returns the string representation
func (s CreateLoadBalancerInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateLoadBalancerInput) GoString() string {
	return s.String()
}

// SetAddressIpVersion sets the AddressIpVersion field's value.
func (s *CreateLoadBalancerInput) SetAddressIpVersion(v string) *CreateLoadBalancerInput {
	s.AddressIpVersion = &v
	return s
}

// SetBandwidthPackageId sets the BandwidthPackageId field's value.
func (s *CreateLoadBalancerInput) SetBandwidthPackageId(v string) *CreateLoadBalancerInput {
	s.BandwidthPackageId = &v
	return s
}

// SetDeleteProtection sets the DeleteProtection field's value.
func (s *CreateLoadBalancerInput) SetDeleteProtection(v string) *CreateLoadBalancerInput {
	s.DeleteProtection = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *CreateLoadBalancerInput) SetDescription(v string) *CreateLoadBalancerInput {
	s.Description = &v
	return s
}

// SetEipBillingConfig sets the EipBillingConfig field's value.
func (s *CreateLoadBalancerInput) SetEipBillingConfig(v *EipBillingConfigForCreateLoadBalancerInput) *CreateLoadBalancerInput {
	s.EipBillingConfig = v
	return s
}

// SetIpv6EipBillingConfig sets the Ipv6EipBillingConfig field's value.
func (s *CreateLoadBalancerInput) SetIpv6EipBillingConfig(v *Ipv6EipBillingConfigForCreateLoadBalancerInput) *CreateLoadBalancerInput {
	s.Ipv6EipBillingConfig = v
	return s
}

// SetLoadBalancerBillingType sets the LoadBalancerBillingType field's value.
func (s *CreateLoadBalancerInput) SetLoadBalancerBillingType(v int64) *CreateLoadBalancerInput {
	s.LoadBalancerBillingType = &v
	return s
}

// SetLoadBalancerName sets the LoadBalancerName field's value.
func (s *CreateLoadBalancerInput) SetLoadBalancerName(v string) *CreateLoadBalancerInput {
	s.LoadBalancerName = &v
	return s
}

// SetModificationProtectionReason sets the ModificationProtectionReason field's value.
func (s *CreateLoadBalancerInput) SetModificationProtectionReason(v string) *CreateLoadBalancerInput {
	s.ModificationProtectionReason = &v
	return s
}

// SetModificationProtectionStatus sets the ModificationProtectionStatus field's value.
func (s *CreateLoadBalancerInput) SetModificationProtectionStatus(v string) *CreateLoadBalancerInput {
	s.ModificationProtectionStatus = &v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *CreateLoadBalancerInput) SetProjectName(v string) *CreateLoadBalancerInput {
	s.ProjectName = &v
	return s
}

// SetRegionId sets the RegionId field's value.
func (s *CreateLoadBalancerInput) SetRegionId(v string) *CreateLoadBalancerInput {
	s.RegionId = &v
	return s
}

// SetSubnetId sets the SubnetId field's value.
func (s *CreateLoadBalancerInput) SetSubnetId(v string) *CreateLoadBalancerInput {
	s.SubnetId = &v
	return s
}

// SetTags sets the Tags field's value.
func (s *CreateLoadBalancerInput) SetTags(v []*TagForCreateLoadBalancerInput) *CreateLoadBalancerInput {
	s.Tags = v
	return s
}

// SetType sets the Type field's value.
func (s *CreateLoadBalancerInput) SetType(v string) *CreateLoadBalancerInput {
	s.Type = &v
	return s
}

// SetVpcId sets the VpcId field's value.
func (s *CreateLoadBalancerInput) SetVpcId(v string) *CreateLoadBalancerInput {
	s.VpcId = &v
	return s
}

// SetZoneMappings sets the ZoneMappings field's value.
func (s *CreateLoadBalancerInput) SetZoneMappings(v []*ZoneMappingForCreateLoadBalancerInput) *CreateLoadBalancerInput {
	s.ZoneMappings = v
	return s
}

type CreateLoadBalancerOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	LoadBalancerId *string `type:"string"`

	RequestId *string `type:"string"`
}

// String returns the string representation
func (s CreateLoadBalancerOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateLoadBalancerOutput) GoString() string {
	return s.String()
}

// SetLoadBalancerId sets the LoadBalancerId field's value.
func (s *CreateLoadBalancerOutput) SetLoadBalancerId(v string) *CreateLoadBalancerOutput {
	s.LoadBalancerId = &v
	return s
}

// SetRequestId sets the RequestId field's value.
func (s *CreateLoadBalancerOutput) SetRequestId(v string) *CreateLoadBalancerOutput {
	s.RequestId = &v
	return s
}

type EipBillingConfigForCreateLoadBalancerInput struct {
	_ struct{} `type:"structure"`

	Bandwidth *int64 `type:"integer"`

	EipBillingType *int64 `type:"integer"`

	ISP *string `type:"string"`

	SecurityProtectionInstanceId *string `type:"string"`

	SecurityProtectionTypes *string `type:"string"`
}

// String returns the string representation
func (s EipBillingConfigForCreateLoadBalancerInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s EipBillingConfigForCreateLoadBalancerInput) GoString() string {
	return s.String()
}

// SetBandwidth sets the Bandwidth field's value.
func (s *EipBillingConfigForCreateLoadBalancerInput) SetBandwidth(v int64) *EipBillingConfigForCreateLoadBalancerInput {
	s.Bandwidth = &v
	return s
}

// SetEipBillingType sets the EipBillingType field's value.
func (s *EipBillingConfigForCreateLoadBalancerInput) SetEipBillingType(v int64) *EipBillingConfigForCreateLoadBalancerInput {
	s.EipBillingType = &v
	return s
}

// SetISP sets the ISP field's value.
func (s *EipBillingConfigForCreateLoadBalancerInput) SetISP(v string) *EipBillingConfigForCreateLoadBalancerInput {
	s.ISP = &v
	return s
}

// SetSecurityProtectionInstanceId sets the SecurityProtectionInstanceId field's value.
func (s *EipBillingConfigForCreateLoadBalancerInput) SetSecurityProtectionInstanceId(v string) *EipBillingConfigForCreateLoadBalancerInput {
	s.SecurityProtectionInstanceId = &v
	return s
}

// SetSecurityProtectionTypes sets the SecurityProtectionTypes field's value.
func (s *EipBillingConfigForCreateLoadBalancerInput) SetSecurityProtectionTypes(v string) *EipBillingConfigForCreateLoadBalancerInput {
	s.SecurityProtectionTypes = &v
	return s
}

type Ipv6EipBillingConfigForCreateLoadBalancerInput struct {
	_ struct{} `type:"structure"`

	Bandwidth *int64 `type:"integer"`

	BillingType *int64 `type:"integer"`

	ISP *string `type:"string"`
}

// String returns the string representation
func (s Ipv6EipBillingConfigForCreateLoadBalancerInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s Ipv6EipBillingConfigForCreateLoadBalancerInput) GoString() string {
	return s.String()
}

// SetBandwidth sets the Bandwidth field's value.
func (s *Ipv6EipBillingConfigForCreateLoadBalancerInput) SetBandwidth(v int64) *Ipv6EipBillingConfigForCreateLoadBalancerInput {
	s.Bandwidth = &v
	return s
}

// SetBillingType sets the BillingType field's value.
func (s *Ipv6EipBillingConfigForCreateLoadBalancerInput) SetBillingType(v int64) *Ipv6EipBillingConfigForCreateLoadBalancerInput {
	s.BillingType = &v
	return s
}

// SetISP sets the ISP field's value.
func (s *Ipv6EipBillingConfigForCreateLoadBalancerInput) SetISP(v string) *Ipv6EipBillingConfigForCreateLoadBalancerInput {
	s.ISP = &v
	return s
}

type TagForCreateLoadBalancerInput struct {
	_ struct{} `type:"structure"`

	Key *string `type:"string"`

	Value *string `type:"string"`
}

// String returns the string representation
func (s TagForCreateLoadBalancerInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s TagForCreateLoadBalancerInput) GoString() string {
	return s.String()
}

// SetKey sets the Key field's value.
func (s *TagForCreateLoadBalancerInput) SetKey(v string) *TagForCreateLoadBalancerInput {
	s.Key = &v
	return s
}

// SetValue sets the Value field's value.
func (s *TagForCreateLoadBalancerInput) SetValue(v string) *TagForCreateLoadBalancerInput {
	s.Value = &v
	return s
}

type ZoneMappingForCreateLoadBalancerInput struct {
	_ struct{} `type:"structure"`

	SubnetId *string `type:"string"`

	ZoneId *string `type:"string"`
}

// String returns the string representation
func (s ZoneMappingForCreateLoadBalancerInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ZoneMappingForCreateLoadBalancerInput) GoString() string {
	return s.String()
}

// SetSubnetId sets the SubnetId field's value.
func (s *ZoneMappingForCreateLoadBalancerInput) SetSubnetId(v string) *ZoneMappingForCreateLoadBalancerInput {
	s.SubnetId = &v
	return s
}

// SetZoneId sets the ZoneId field's value.
func (s *ZoneMappingForCreateLoadBalancerInput) SetZoneId(v string) *ZoneMappingForCreateLoadBalancerInput {
	s.ZoneId = &v
	return s
}

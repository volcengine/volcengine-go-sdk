// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package ecs

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribeLaunchTemplateVersionsCommon = "DescribeLaunchTemplateVersions"

// DescribeLaunchTemplateVersionsCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeLaunchTemplateVersionsCommon operation. The "output" return
// value will be populated with the DescribeLaunchTemplateVersionsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeLaunchTemplateVersionsCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeLaunchTemplateVersionsCommon Send returns without error.
//
// See DescribeLaunchTemplateVersionsCommon for more information on using the DescribeLaunchTemplateVersionsCommon
// API call, and error handling.
//
//    // Example sending a request using the DescribeLaunchTemplateVersionsCommonRequest method.
//    req, resp := client.DescribeLaunchTemplateVersionsCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ECS) DescribeLaunchTemplateVersionsCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeLaunchTemplateVersionsCommon,
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

// DescribeLaunchTemplateVersionsCommon API operation for ECS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ECS's
// API operation DescribeLaunchTemplateVersionsCommon for usage and error information.
func (c *ECS) DescribeLaunchTemplateVersionsCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeLaunchTemplateVersionsCommonRequest(input)
	return out, req.Send()
}

// DescribeLaunchTemplateVersionsCommonWithContext is the same as DescribeLaunchTemplateVersionsCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeLaunchTemplateVersionsCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ECS) DescribeLaunchTemplateVersionsCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeLaunchTemplateVersionsCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeLaunchTemplateVersions = "DescribeLaunchTemplateVersions"

// DescribeLaunchTemplateVersionsRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeLaunchTemplateVersions operation. The "output" return
// value will be populated with the DescribeLaunchTemplateVersionsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeLaunchTemplateVersionsCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeLaunchTemplateVersionsCommon Send returns without error.
//
// See DescribeLaunchTemplateVersions for more information on using the DescribeLaunchTemplateVersions
// API call, and error handling.
//
//    // Example sending a request using the DescribeLaunchTemplateVersionsRequest method.
//    req, resp := client.DescribeLaunchTemplateVersionsRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ECS) DescribeLaunchTemplateVersionsRequest(input *DescribeLaunchTemplateVersionsInput) (req *request.Request, output *DescribeLaunchTemplateVersionsOutput) {
	op := &request.Operation{
		Name:       opDescribeLaunchTemplateVersions,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeLaunchTemplateVersionsInput{}
	}

	output = &DescribeLaunchTemplateVersionsOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DescribeLaunchTemplateVersions API operation for ECS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ECS's
// API operation DescribeLaunchTemplateVersions for usage and error information.
func (c *ECS) DescribeLaunchTemplateVersions(input *DescribeLaunchTemplateVersionsInput) (*DescribeLaunchTemplateVersionsOutput, error) {
	req, out := c.DescribeLaunchTemplateVersionsRequest(input)
	return out, req.Send()
}

// DescribeLaunchTemplateVersionsWithContext is the same as DescribeLaunchTemplateVersions with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeLaunchTemplateVersions for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ECS) DescribeLaunchTemplateVersionsWithContext(ctx volcengine.Context, input *DescribeLaunchTemplateVersionsInput, opts ...request.Option) (*DescribeLaunchTemplateVersionsOutput, error) {
	req, out := c.DescribeLaunchTemplateVersionsRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DescribeLaunchTemplateVersionsInput struct {
	_ struct{} `type:"structure"`

	DefaultVersion *bool `type:"boolean"`

	LaunchTemplateId *string `type:"string"`

	LaunchTemplateName *string `type:"string"`

	LaunchTemplateVersions []*int64 `type:"list"`

	PageNumber *int32 `type:"int32"`

	PageSize *int32 `type:"int32"`
}

// String returns the string representation
func (s DescribeLaunchTemplateVersionsInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeLaunchTemplateVersionsInput) GoString() string {
	return s.String()
}

// SetDefaultVersion sets the DefaultVersion field's value.
func (s *DescribeLaunchTemplateVersionsInput) SetDefaultVersion(v bool) *DescribeLaunchTemplateVersionsInput {
	s.DefaultVersion = &v
	return s
}

// SetLaunchTemplateId sets the LaunchTemplateId field's value.
func (s *DescribeLaunchTemplateVersionsInput) SetLaunchTemplateId(v string) *DescribeLaunchTemplateVersionsInput {
	s.LaunchTemplateId = &v
	return s
}

// SetLaunchTemplateName sets the LaunchTemplateName field's value.
func (s *DescribeLaunchTemplateVersionsInput) SetLaunchTemplateName(v string) *DescribeLaunchTemplateVersionsInput {
	s.LaunchTemplateName = &v
	return s
}

// SetLaunchTemplateVersions sets the LaunchTemplateVersions field's value.
func (s *DescribeLaunchTemplateVersionsInput) SetLaunchTemplateVersions(v []*int64) *DescribeLaunchTemplateVersionsInput {
	s.LaunchTemplateVersions = v
	return s
}

// SetPageNumber sets the PageNumber field's value.
func (s *DescribeLaunchTemplateVersionsInput) SetPageNumber(v int32) *DescribeLaunchTemplateVersionsInput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *DescribeLaunchTemplateVersionsInput) SetPageSize(v int32) *DescribeLaunchTemplateVersionsInput {
	s.PageSize = &v
	return s
}

type DescribeLaunchTemplateVersionsOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	LaunchTemplateVersions []*LaunchTemplateVersionForDescribeLaunchTemplateVersionsOutput `type:"list"`

	PageNumber *int32 `type:"int32"`

	PageSize *int32 `type:"int32"`

	TotalCount *int32 `type:"int32"`
}

// String returns the string representation
func (s DescribeLaunchTemplateVersionsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeLaunchTemplateVersionsOutput) GoString() string {
	return s.String()
}

// SetLaunchTemplateVersions sets the LaunchTemplateVersions field's value.
func (s *DescribeLaunchTemplateVersionsOutput) SetLaunchTemplateVersions(v []*LaunchTemplateVersionForDescribeLaunchTemplateVersionsOutput) *DescribeLaunchTemplateVersionsOutput {
	s.LaunchTemplateVersions = v
	return s
}

// SetPageNumber sets the PageNumber field's value.
func (s *DescribeLaunchTemplateVersionsOutput) SetPageNumber(v int32) *DescribeLaunchTemplateVersionsOutput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *DescribeLaunchTemplateVersionsOutput) SetPageSize(v int32) *DescribeLaunchTemplateVersionsOutput {
	s.PageSize = &v
	return s
}

// SetTotalCount sets the TotalCount field's value.
func (s *DescribeLaunchTemplateVersionsOutput) SetTotalCount(v int32) *DescribeLaunchTemplateVersionsOutput {
	s.TotalCount = &v
	return s
}

type EipForDescribeLaunchTemplateVersionsOutput struct {
	_ struct{} `type:"structure"`

	Bandwidth *int32 `type:"int32"`

	BillingType *int64 `type:"int64"`

	ISP *string `type:"string"`
}

// String returns the string representation
func (s EipForDescribeLaunchTemplateVersionsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s EipForDescribeLaunchTemplateVersionsOutput) GoString() string {
	return s.String()
}

// SetBandwidth sets the Bandwidth field's value.
func (s *EipForDescribeLaunchTemplateVersionsOutput) SetBandwidth(v int32) *EipForDescribeLaunchTemplateVersionsOutput {
	s.Bandwidth = &v
	return s
}

// SetBillingType sets the BillingType field's value.
func (s *EipForDescribeLaunchTemplateVersionsOutput) SetBillingType(v int64) *EipForDescribeLaunchTemplateVersionsOutput {
	s.BillingType = &v
	return s
}

// SetISP sets the ISP field's value.
func (s *EipForDescribeLaunchTemplateVersionsOutput) SetISP(v string) *EipForDescribeLaunchTemplateVersionsOutput {
	s.ISP = &v
	return s
}

type LaunchTemplateVersionDataForDescribeLaunchTemplateVersionsOutput struct {
	_ struct{} `type:"structure"`

	Description *string `type:"string"`

	Eip *EipForDescribeLaunchTemplateVersionsOutput `type:"structure"`

	HostName *string `type:"string"`

	HpcClusterId *string `type:"string"`

	ImageId *string `type:"string"`

	InstanceChargeType *string `type:"string"`

	InstanceName *string `type:"string"`

	InstanceTypeId *string `type:"string"`

	KeyPairName *string `type:"string"`

	NetworkInterfaces []*NetworkInterfaceForDescribeLaunchTemplateVersionsOutput `type:"list"`

	ProjectName *string `type:"string"`

	SecurityEnhancementStrategy *string `type:"string"`

	SpotStrategy *string `type:"string"`

	SuffixIndex *int32 `type:"int32"`

	Tags []*TagForDescribeLaunchTemplateVersionsOutput `type:"list"`

	UniqueSuffix *bool `type:"boolean"`

	UserData *string `type:"string"`

	Volumes []*VolumeForDescribeLaunchTemplateVersionsOutput `type:"list"`

	VpcId *string `type:"string"`

	ZoneId *string `type:"string"`
}

// String returns the string representation
func (s LaunchTemplateVersionDataForDescribeLaunchTemplateVersionsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s LaunchTemplateVersionDataForDescribeLaunchTemplateVersionsOutput) GoString() string {
	return s.String()
}

// SetDescription sets the Description field's value.
func (s *LaunchTemplateVersionDataForDescribeLaunchTemplateVersionsOutput) SetDescription(v string) *LaunchTemplateVersionDataForDescribeLaunchTemplateVersionsOutput {
	s.Description = &v
	return s
}

// SetEip sets the Eip field's value.
func (s *LaunchTemplateVersionDataForDescribeLaunchTemplateVersionsOutput) SetEip(v *EipForDescribeLaunchTemplateVersionsOutput) *LaunchTemplateVersionDataForDescribeLaunchTemplateVersionsOutput {
	s.Eip = v
	return s
}

// SetHostName sets the HostName field's value.
func (s *LaunchTemplateVersionDataForDescribeLaunchTemplateVersionsOutput) SetHostName(v string) *LaunchTemplateVersionDataForDescribeLaunchTemplateVersionsOutput {
	s.HostName = &v
	return s
}

// SetHpcClusterId sets the HpcClusterId field's value.
func (s *LaunchTemplateVersionDataForDescribeLaunchTemplateVersionsOutput) SetHpcClusterId(v string) *LaunchTemplateVersionDataForDescribeLaunchTemplateVersionsOutput {
	s.HpcClusterId = &v
	return s
}

// SetImageId sets the ImageId field's value.
func (s *LaunchTemplateVersionDataForDescribeLaunchTemplateVersionsOutput) SetImageId(v string) *LaunchTemplateVersionDataForDescribeLaunchTemplateVersionsOutput {
	s.ImageId = &v
	return s
}

// SetInstanceChargeType sets the InstanceChargeType field's value.
func (s *LaunchTemplateVersionDataForDescribeLaunchTemplateVersionsOutput) SetInstanceChargeType(v string) *LaunchTemplateVersionDataForDescribeLaunchTemplateVersionsOutput {
	s.InstanceChargeType = &v
	return s
}

// SetInstanceName sets the InstanceName field's value.
func (s *LaunchTemplateVersionDataForDescribeLaunchTemplateVersionsOutput) SetInstanceName(v string) *LaunchTemplateVersionDataForDescribeLaunchTemplateVersionsOutput {
	s.InstanceName = &v
	return s
}

// SetInstanceTypeId sets the InstanceTypeId field's value.
func (s *LaunchTemplateVersionDataForDescribeLaunchTemplateVersionsOutput) SetInstanceTypeId(v string) *LaunchTemplateVersionDataForDescribeLaunchTemplateVersionsOutput {
	s.InstanceTypeId = &v
	return s
}

// SetKeyPairName sets the KeyPairName field's value.
func (s *LaunchTemplateVersionDataForDescribeLaunchTemplateVersionsOutput) SetKeyPairName(v string) *LaunchTemplateVersionDataForDescribeLaunchTemplateVersionsOutput {
	s.KeyPairName = &v
	return s
}

// SetNetworkInterfaces sets the NetworkInterfaces field's value.
func (s *LaunchTemplateVersionDataForDescribeLaunchTemplateVersionsOutput) SetNetworkInterfaces(v []*NetworkInterfaceForDescribeLaunchTemplateVersionsOutput) *LaunchTemplateVersionDataForDescribeLaunchTemplateVersionsOutput {
	s.NetworkInterfaces = v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *LaunchTemplateVersionDataForDescribeLaunchTemplateVersionsOutput) SetProjectName(v string) *LaunchTemplateVersionDataForDescribeLaunchTemplateVersionsOutput {
	s.ProjectName = &v
	return s
}

// SetSecurityEnhancementStrategy sets the SecurityEnhancementStrategy field's value.
func (s *LaunchTemplateVersionDataForDescribeLaunchTemplateVersionsOutput) SetSecurityEnhancementStrategy(v string) *LaunchTemplateVersionDataForDescribeLaunchTemplateVersionsOutput {
	s.SecurityEnhancementStrategy = &v
	return s
}

// SetSpotStrategy sets the SpotStrategy field's value.
func (s *LaunchTemplateVersionDataForDescribeLaunchTemplateVersionsOutput) SetSpotStrategy(v string) *LaunchTemplateVersionDataForDescribeLaunchTemplateVersionsOutput {
	s.SpotStrategy = &v
	return s
}

// SetSuffixIndex sets the SuffixIndex field's value.
func (s *LaunchTemplateVersionDataForDescribeLaunchTemplateVersionsOutput) SetSuffixIndex(v int32) *LaunchTemplateVersionDataForDescribeLaunchTemplateVersionsOutput {
	s.SuffixIndex = &v
	return s
}

// SetTags sets the Tags field's value.
func (s *LaunchTemplateVersionDataForDescribeLaunchTemplateVersionsOutput) SetTags(v []*TagForDescribeLaunchTemplateVersionsOutput) *LaunchTemplateVersionDataForDescribeLaunchTemplateVersionsOutput {
	s.Tags = v
	return s
}

// SetUniqueSuffix sets the UniqueSuffix field's value.
func (s *LaunchTemplateVersionDataForDescribeLaunchTemplateVersionsOutput) SetUniqueSuffix(v bool) *LaunchTemplateVersionDataForDescribeLaunchTemplateVersionsOutput {
	s.UniqueSuffix = &v
	return s
}

// SetUserData sets the UserData field's value.
func (s *LaunchTemplateVersionDataForDescribeLaunchTemplateVersionsOutput) SetUserData(v string) *LaunchTemplateVersionDataForDescribeLaunchTemplateVersionsOutput {
	s.UserData = &v
	return s
}

// SetVolumes sets the Volumes field's value.
func (s *LaunchTemplateVersionDataForDescribeLaunchTemplateVersionsOutput) SetVolumes(v []*VolumeForDescribeLaunchTemplateVersionsOutput) *LaunchTemplateVersionDataForDescribeLaunchTemplateVersionsOutput {
	s.Volumes = v
	return s
}

// SetVpcId sets the VpcId field's value.
func (s *LaunchTemplateVersionDataForDescribeLaunchTemplateVersionsOutput) SetVpcId(v string) *LaunchTemplateVersionDataForDescribeLaunchTemplateVersionsOutput {
	s.VpcId = &v
	return s
}

// SetZoneId sets the ZoneId field's value.
func (s *LaunchTemplateVersionDataForDescribeLaunchTemplateVersionsOutput) SetZoneId(v string) *LaunchTemplateVersionDataForDescribeLaunchTemplateVersionsOutput {
	s.ZoneId = &v
	return s
}

type LaunchTemplateVersionForDescribeLaunchTemplateVersionsOutput struct {
	_ struct{} `type:"structure"`

	CreatedAt *string `type:"string"`

	LaunchTemplateId *string `type:"string"`

	LaunchTemplateVersionData *LaunchTemplateVersionDataForDescribeLaunchTemplateVersionsOutput `type:"structure"`

	UpdatedAt *string `type:"string"`

	VersionDescription *string `type:"string"`

	VersionNumber *int32 `type:"int32"`
}

// String returns the string representation
func (s LaunchTemplateVersionForDescribeLaunchTemplateVersionsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s LaunchTemplateVersionForDescribeLaunchTemplateVersionsOutput) GoString() string {
	return s.String()
}

// SetCreatedAt sets the CreatedAt field's value.
func (s *LaunchTemplateVersionForDescribeLaunchTemplateVersionsOutput) SetCreatedAt(v string) *LaunchTemplateVersionForDescribeLaunchTemplateVersionsOutput {
	s.CreatedAt = &v
	return s
}

// SetLaunchTemplateId sets the LaunchTemplateId field's value.
func (s *LaunchTemplateVersionForDescribeLaunchTemplateVersionsOutput) SetLaunchTemplateId(v string) *LaunchTemplateVersionForDescribeLaunchTemplateVersionsOutput {
	s.LaunchTemplateId = &v
	return s
}

// SetLaunchTemplateVersionData sets the LaunchTemplateVersionData field's value.
func (s *LaunchTemplateVersionForDescribeLaunchTemplateVersionsOutput) SetLaunchTemplateVersionData(v *LaunchTemplateVersionDataForDescribeLaunchTemplateVersionsOutput) *LaunchTemplateVersionForDescribeLaunchTemplateVersionsOutput {
	s.LaunchTemplateVersionData = v
	return s
}

// SetUpdatedAt sets the UpdatedAt field's value.
func (s *LaunchTemplateVersionForDescribeLaunchTemplateVersionsOutput) SetUpdatedAt(v string) *LaunchTemplateVersionForDescribeLaunchTemplateVersionsOutput {
	s.UpdatedAt = &v
	return s
}

// SetVersionDescription sets the VersionDescription field's value.
func (s *LaunchTemplateVersionForDescribeLaunchTemplateVersionsOutput) SetVersionDescription(v string) *LaunchTemplateVersionForDescribeLaunchTemplateVersionsOutput {
	s.VersionDescription = &v
	return s
}

// SetVersionNumber sets the VersionNumber field's value.
func (s *LaunchTemplateVersionForDescribeLaunchTemplateVersionsOutput) SetVersionNumber(v int32) *LaunchTemplateVersionForDescribeLaunchTemplateVersionsOutput {
	s.VersionNumber = &v
	return s
}

type NetworkInterfaceForDescribeLaunchTemplateVersionsOutput struct {
	_ struct{} `type:"structure"`

	SecurityGroupIds []*string `type:"list"`

	SubnetId *string `type:"string"`
}

// String returns the string representation
func (s NetworkInterfaceForDescribeLaunchTemplateVersionsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s NetworkInterfaceForDescribeLaunchTemplateVersionsOutput) GoString() string {
	return s.String()
}

// SetSecurityGroupIds sets the SecurityGroupIds field's value.
func (s *NetworkInterfaceForDescribeLaunchTemplateVersionsOutput) SetSecurityGroupIds(v []*string) *NetworkInterfaceForDescribeLaunchTemplateVersionsOutput {
	s.SecurityGroupIds = v
	return s
}

// SetSubnetId sets the SubnetId field's value.
func (s *NetworkInterfaceForDescribeLaunchTemplateVersionsOutput) SetSubnetId(v string) *NetworkInterfaceForDescribeLaunchTemplateVersionsOutput {
	s.SubnetId = &v
	return s
}

type TagForDescribeLaunchTemplateVersionsOutput struct {
	_ struct{} `type:"structure"`

	Key *string `type:"string"`

	Value *string `type:"string"`
}

// String returns the string representation
func (s TagForDescribeLaunchTemplateVersionsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s TagForDescribeLaunchTemplateVersionsOutput) GoString() string {
	return s.String()
}

// SetKey sets the Key field's value.
func (s *TagForDescribeLaunchTemplateVersionsOutput) SetKey(v string) *TagForDescribeLaunchTemplateVersionsOutput {
	s.Key = &v
	return s
}

// SetValue sets the Value field's value.
func (s *TagForDescribeLaunchTemplateVersionsOutput) SetValue(v string) *TagForDescribeLaunchTemplateVersionsOutput {
	s.Value = &v
	return s
}

type VolumeForDescribeLaunchTemplateVersionsOutput struct {
	_ struct{} `type:"structure"`

	DeleteWithInstance *bool `type:"boolean"`

	Size *int32 `type:"int32"`

	VolumeType *string `type:"string"`
}

// String returns the string representation
func (s VolumeForDescribeLaunchTemplateVersionsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s VolumeForDescribeLaunchTemplateVersionsOutput) GoString() string {
	return s.String()
}

// SetDeleteWithInstance sets the DeleteWithInstance field's value.
func (s *VolumeForDescribeLaunchTemplateVersionsOutput) SetDeleteWithInstance(v bool) *VolumeForDescribeLaunchTemplateVersionsOutput {
	s.DeleteWithInstance = &v
	return s
}

// SetSize sets the Size field's value.
func (s *VolumeForDescribeLaunchTemplateVersionsOutput) SetSize(v int32) *VolumeForDescribeLaunchTemplateVersionsOutput {
	s.Size = &v
	return s
}

// SetVolumeType sets the VolumeType field's value.
func (s *VolumeForDescribeLaunchTemplateVersionsOutput) SetVolumeType(v string) *VolumeForDescribeLaunchTemplateVersionsOutput {
	s.VolumeType = &v
	return s
}
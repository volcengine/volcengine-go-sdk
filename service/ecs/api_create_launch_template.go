// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package ecs

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opCreateLaunchTemplateCommon = "CreateLaunchTemplate"

// CreateLaunchTemplateCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateLaunchTemplateCommon operation. The "output" return
// value will be populated with the CreateLaunchTemplateCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateLaunchTemplateCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateLaunchTemplateCommon Send returns without error.
//
// See CreateLaunchTemplateCommon for more information on using the CreateLaunchTemplateCommon
// API call, and error handling.
//
//    // Example sending a request using the CreateLaunchTemplateCommonRequest method.
//    req, resp := client.CreateLaunchTemplateCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ECS) CreateLaunchTemplateCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opCreateLaunchTemplateCommon,
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

// CreateLaunchTemplateCommon API operation for ECS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ECS's
// API operation CreateLaunchTemplateCommon for usage and error information.
func (c *ECS) CreateLaunchTemplateCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.CreateLaunchTemplateCommonRequest(input)
	return out, req.Send()
}

// CreateLaunchTemplateCommonWithContext is the same as CreateLaunchTemplateCommon with the addition of
// the ability to pass a context and additional request options.
//
// See CreateLaunchTemplateCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ECS) CreateLaunchTemplateCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.CreateLaunchTemplateCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opCreateLaunchTemplate = "CreateLaunchTemplate"

// CreateLaunchTemplateRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateLaunchTemplate operation. The "output" return
// value will be populated with the CreateLaunchTemplateCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateLaunchTemplateCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateLaunchTemplateCommon Send returns without error.
//
// See CreateLaunchTemplate for more information on using the CreateLaunchTemplate
// API call, and error handling.
//
//    // Example sending a request using the CreateLaunchTemplateRequest method.
//    req, resp := client.CreateLaunchTemplateRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ECS) CreateLaunchTemplateRequest(input *CreateLaunchTemplateInput) (req *request.Request, output *CreateLaunchTemplateOutput) {
	op := &request.Operation{
		Name:       opCreateLaunchTemplate,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateLaunchTemplateInput{}
	}

	output = &CreateLaunchTemplateOutput{}
	req = c.newRequest(op, input, output)

	return
}

// CreateLaunchTemplate API operation for ECS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ECS's
// API operation CreateLaunchTemplate for usage and error information.
func (c *ECS) CreateLaunchTemplate(input *CreateLaunchTemplateInput) (*CreateLaunchTemplateOutput, error) {
	req, out := c.CreateLaunchTemplateRequest(input)
	return out, req.Send()
}

// CreateLaunchTemplateWithContext is the same as CreateLaunchTemplate with the addition of
// the ability to pass a context and additional request options.
//
// See CreateLaunchTemplate for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ECS) CreateLaunchTemplateWithContext(ctx volcengine.Context, input *CreateLaunchTemplateInput, opts ...request.Option) (*CreateLaunchTemplateOutput, error) {
	req, out := c.CreateLaunchTemplateRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type CreateLaunchTemplateInput struct {
	_ struct{} `type:"structure"`

	Description *string `type:"string"`

	Eip *EipForCreateLaunchTemplateInput `type:"structure"`

	HostName *string `type:"string"`

	HpcClusterId *string `type:"string"`

	ImageId *string `type:"string"`

	ImageName *string `type:"string"`

	InstanceChargeType *string `type:"string"`

	InstanceName *string `type:"string"`

	InstanceTypeId *string `type:"string"`

	KeyPairName *string `type:"string"`

	// LaunchTemplateName is a required field
	LaunchTemplateName *string `type:"string" required:"true"`

	NetworkInterfaces []*NetworkInterfaceForCreateLaunchTemplateInput `type:"list"`

	ProjectName *string `type:"string"`

	SecurityEnhancementStrategy *string `type:"string"`

	SpotStrategy *string `type:"string"`

	SuffixIndex *int32 `type:"int32"`

	Tags []*TagForCreateLaunchTemplateInput `type:"list"`

	UniqueSuffix *bool `type:"boolean"`

	UserData *string `type:"string"`

	VersionDescription *string `type:"string"`

	Volumes []*VolumeForCreateLaunchTemplateInput `type:"list"`

	VpcId *string `type:"string"`

	// ZoneId is a required field
	ZoneId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s CreateLaunchTemplateInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateLaunchTemplateInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateLaunchTemplateInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "CreateLaunchTemplateInput"}
	if s.LaunchTemplateName == nil {
		invalidParams.Add(request.NewErrParamRequired("LaunchTemplateName"))
	}
	if s.ZoneId == nil {
		invalidParams.Add(request.NewErrParamRequired("ZoneId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetDescription sets the Description field's value.
func (s *CreateLaunchTemplateInput) SetDescription(v string) *CreateLaunchTemplateInput {
	s.Description = &v
	return s
}

// SetEip sets the Eip field's value.
func (s *CreateLaunchTemplateInput) SetEip(v *EipForCreateLaunchTemplateInput) *CreateLaunchTemplateInput {
	s.Eip = v
	return s
}

// SetHostName sets the HostName field's value.
func (s *CreateLaunchTemplateInput) SetHostName(v string) *CreateLaunchTemplateInput {
	s.HostName = &v
	return s
}

// SetHpcClusterId sets the HpcClusterId field's value.
func (s *CreateLaunchTemplateInput) SetHpcClusterId(v string) *CreateLaunchTemplateInput {
	s.HpcClusterId = &v
	return s
}

// SetImageId sets the ImageId field's value.
func (s *CreateLaunchTemplateInput) SetImageId(v string) *CreateLaunchTemplateInput {
	s.ImageId = &v
	return s
}

// SetImageName sets the ImageName field's value.
func (s *CreateLaunchTemplateInput) SetImageName(v string) *CreateLaunchTemplateInput {
	s.ImageName = &v
	return s
}

// SetInstanceChargeType sets the InstanceChargeType field's value.
func (s *CreateLaunchTemplateInput) SetInstanceChargeType(v string) *CreateLaunchTemplateInput {
	s.InstanceChargeType = &v
	return s
}

// SetInstanceName sets the InstanceName field's value.
func (s *CreateLaunchTemplateInput) SetInstanceName(v string) *CreateLaunchTemplateInput {
	s.InstanceName = &v
	return s
}

// SetInstanceTypeId sets the InstanceTypeId field's value.
func (s *CreateLaunchTemplateInput) SetInstanceTypeId(v string) *CreateLaunchTemplateInput {
	s.InstanceTypeId = &v
	return s
}

// SetKeyPairName sets the KeyPairName field's value.
func (s *CreateLaunchTemplateInput) SetKeyPairName(v string) *CreateLaunchTemplateInput {
	s.KeyPairName = &v
	return s
}

// SetLaunchTemplateName sets the LaunchTemplateName field's value.
func (s *CreateLaunchTemplateInput) SetLaunchTemplateName(v string) *CreateLaunchTemplateInput {
	s.LaunchTemplateName = &v
	return s
}

// SetNetworkInterfaces sets the NetworkInterfaces field's value.
func (s *CreateLaunchTemplateInput) SetNetworkInterfaces(v []*NetworkInterfaceForCreateLaunchTemplateInput) *CreateLaunchTemplateInput {
	s.NetworkInterfaces = v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *CreateLaunchTemplateInput) SetProjectName(v string) *CreateLaunchTemplateInput {
	s.ProjectName = &v
	return s
}

// SetSecurityEnhancementStrategy sets the SecurityEnhancementStrategy field's value.
func (s *CreateLaunchTemplateInput) SetSecurityEnhancementStrategy(v string) *CreateLaunchTemplateInput {
	s.SecurityEnhancementStrategy = &v
	return s
}

// SetSpotStrategy sets the SpotStrategy field's value.
func (s *CreateLaunchTemplateInput) SetSpotStrategy(v string) *CreateLaunchTemplateInput {
	s.SpotStrategy = &v
	return s
}

// SetSuffixIndex sets the SuffixIndex field's value.
func (s *CreateLaunchTemplateInput) SetSuffixIndex(v int32) *CreateLaunchTemplateInput {
	s.SuffixIndex = &v
	return s
}

// SetTags sets the Tags field's value.
func (s *CreateLaunchTemplateInput) SetTags(v []*TagForCreateLaunchTemplateInput) *CreateLaunchTemplateInput {
	s.Tags = v
	return s
}

// SetUniqueSuffix sets the UniqueSuffix field's value.
func (s *CreateLaunchTemplateInput) SetUniqueSuffix(v bool) *CreateLaunchTemplateInput {
	s.UniqueSuffix = &v
	return s
}

// SetUserData sets the UserData field's value.
func (s *CreateLaunchTemplateInput) SetUserData(v string) *CreateLaunchTemplateInput {
	s.UserData = &v
	return s
}

// SetVersionDescription sets the VersionDescription field's value.
func (s *CreateLaunchTemplateInput) SetVersionDescription(v string) *CreateLaunchTemplateInput {
	s.VersionDescription = &v
	return s
}

// SetVolumes sets the Volumes field's value.
func (s *CreateLaunchTemplateInput) SetVolumes(v []*VolumeForCreateLaunchTemplateInput) *CreateLaunchTemplateInput {
	s.Volumes = v
	return s
}

// SetVpcId sets the VpcId field's value.
func (s *CreateLaunchTemplateInput) SetVpcId(v string) *CreateLaunchTemplateInput {
	s.VpcId = &v
	return s
}

// SetZoneId sets the ZoneId field's value.
func (s *CreateLaunchTemplateInput) SetZoneId(v string) *CreateLaunchTemplateInput {
	s.ZoneId = &v
	return s
}

type CreateLaunchTemplateOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	LaunchTemplateId *string `type:"string"`
}

// String returns the string representation
func (s CreateLaunchTemplateOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateLaunchTemplateOutput) GoString() string {
	return s.String()
}

// SetLaunchTemplateId sets the LaunchTemplateId field's value.
func (s *CreateLaunchTemplateOutput) SetLaunchTemplateId(v string) *CreateLaunchTemplateOutput {
	s.LaunchTemplateId = &v
	return s
}

type EipForCreateLaunchTemplateInput struct {
	_ struct{} `type:"structure"`

	Bandwidth *int32 `type:"int32"`

	BillingType *int64 `type:"int64"`

	ISP *string `type:"string"`
}

// String returns the string representation
func (s EipForCreateLaunchTemplateInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s EipForCreateLaunchTemplateInput) GoString() string {
	return s.String()
}

// SetBandwidth sets the Bandwidth field's value.
func (s *EipForCreateLaunchTemplateInput) SetBandwidth(v int32) *EipForCreateLaunchTemplateInput {
	s.Bandwidth = &v
	return s
}

// SetBillingType sets the BillingType field's value.
func (s *EipForCreateLaunchTemplateInput) SetBillingType(v int64) *EipForCreateLaunchTemplateInput {
	s.BillingType = &v
	return s
}

// SetISP sets the ISP field's value.
func (s *EipForCreateLaunchTemplateInput) SetISP(v string) *EipForCreateLaunchTemplateInput {
	s.ISP = &v
	return s
}

type NetworkInterfaceForCreateLaunchTemplateInput struct {
	_ struct{} `type:"structure"`

	SecurityGroupIds []*string `type:"list"`

	SubnetId *string `type:"string"`
}

// String returns the string representation
func (s NetworkInterfaceForCreateLaunchTemplateInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s NetworkInterfaceForCreateLaunchTemplateInput) GoString() string {
	return s.String()
}

// SetSecurityGroupIds sets the SecurityGroupIds field's value.
func (s *NetworkInterfaceForCreateLaunchTemplateInput) SetSecurityGroupIds(v []*string) *NetworkInterfaceForCreateLaunchTemplateInput {
	s.SecurityGroupIds = v
	return s
}

// SetSubnetId sets the SubnetId field's value.
func (s *NetworkInterfaceForCreateLaunchTemplateInput) SetSubnetId(v string) *NetworkInterfaceForCreateLaunchTemplateInput {
	s.SubnetId = &v
	return s
}

type TagForCreateLaunchTemplateInput struct {
	_ struct{} `type:"structure"`

	Key *string `type:"string"`

	Value *string `type:"string"`
}

// String returns the string representation
func (s TagForCreateLaunchTemplateInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s TagForCreateLaunchTemplateInput) GoString() string {
	return s.String()
}

// SetKey sets the Key field's value.
func (s *TagForCreateLaunchTemplateInput) SetKey(v string) *TagForCreateLaunchTemplateInput {
	s.Key = &v
	return s
}

// SetValue sets the Value field's value.
func (s *TagForCreateLaunchTemplateInput) SetValue(v string) *TagForCreateLaunchTemplateInput {
	s.Value = &v
	return s
}

type VolumeForCreateLaunchTemplateInput struct {
	_ struct{} `type:"structure"`

	DeleteWithInstance *bool `type:"boolean"`

	Size *int32 `type:"int32"`

	VolumeType *string `type:"string"`
}

// String returns the string representation
func (s VolumeForCreateLaunchTemplateInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s VolumeForCreateLaunchTemplateInput) GoString() string {
	return s.String()
}

// SetDeleteWithInstance sets the DeleteWithInstance field's value.
func (s *VolumeForCreateLaunchTemplateInput) SetDeleteWithInstance(v bool) *VolumeForCreateLaunchTemplateInput {
	s.DeleteWithInstance = &v
	return s
}

// SetSize sets the Size field's value.
func (s *VolumeForCreateLaunchTemplateInput) SetSize(v int32) *VolumeForCreateLaunchTemplateInput {
	s.Size = &v
	return s
}

// SetVolumeType sets the VolumeType field's value.
func (s *VolumeForCreateLaunchTemplateInput) SetVolumeType(v string) *VolumeForCreateLaunchTemplateInput {
	s.VolumeType = &v
	return s
}
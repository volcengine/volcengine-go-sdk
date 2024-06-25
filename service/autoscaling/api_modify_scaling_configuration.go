// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package autoscaling

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opModifyScalingConfigurationCommon = "ModifyScalingConfiguration"

// ModifyScalingConfigurationCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ModifyScalingConfigurationCommon operation. The "output" return
// value will be populated with the ModifyScalingConfigurationCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyScalingConfigurationCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyScalingConfigurationCommon Send returns without error.
//
// See ModifyScalingConfigurationCommon for more information on using the ModifyScalingConfigurationCommon
// API call, and error handling.
//
//    // Example sending a request using the ModifyScalingConfigurationCommonRequest method.
//    req, resp := client.ModifyScalingConfigurationCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *AUTOSCALING) ModifyScalingConfigurationCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opModifyScalingConfigurationCommon,
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

// ModifyScalingConfigurationCommon API operation for AUTO_SCALING.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for AUTO_SCALING's
// API operation ModifyScalingConfigurationCommon for usage and error information.
func (c *AUTOSCALING) ModifyScalingConfigurationCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ModifyScalingConfigurationCommonRequest(input)
	return out, req.Send()
}

// ModifyScalingConfigurationCommonWithContext is the same as ModifyScalingConfigurationCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyScalingConfigurationCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *AUTOSCALING) ModifyScalingConfigurationCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ModifyScalingConfigurationCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opModifyScalingConfiguration = "ModifyScalingConfiguration"

// ModifyScalingConfigurationRequest generates a "volcengine/request.Request" representing the
// client's request for the ModifyScalingConfiguration operation. The "output" return
// value will be populated with the ModifyScalingConfigurationCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyScalingConfigurationCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyScalingConfigurationCommon Send returns without error.
//
// See ModifyScalingConfiguration for more information on using the ModifyScalingConfiguration
// API call, and error handling.
//
//    // Example sending a request using the ModifyScalingConfigurationRequest method.
//    req, resp := client.ModifyScalingConfigurationRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *AUTOSCALING) ModifyScalingConfigurationRequest(input *ModifyScalingConfigurationInput) (req *request.Request, output *ModifyScalingConfigurationOutput) {
	op := &request.Operation{
		Name:       opModifyScalingConfiguration,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ModifyScalingConfigurationInput{}
	}

	output = &ModifyScalingConfigurationOutput{}
	req = c.newRequest(op, input, output)

	return
}

// ModifyScalingConfiguration API operation for AUTO_SCALING.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for AUTO_SCALING's
// API operation ModifyScalingConfiguration for usage and error information.
func (c *AUTOSCALING) ModifyScalingConfiguration(input *ModifyScalingConfigurationInput) (*ModifyScalingConfigurationOutput, error) {
	req, out := c.ModifyScalingConfigurationRequest(input)
	return out, req.Send()
}

// ModifyScalingConfigurationWithContext is the same as ModifyScalingConfiguration with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyScalingConfiguration for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *AUTOSCALING) ModifyScalingConfigurationWithContext(ctx volcengine.Context, input *ModifyScalingConfigurationInput, opts ...request.Option) (*ModifyScalingConfigurationOutput, error) {
	req, out := c.ModifyScalingConfigurationRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type EipForModifyScalingConfigurationInput struct {
	_ struct{} `type:"structure"`

	Bandwidth *int32 `type:"int32"`

	BandwidthPackageId *string `type:"string"`

	BillingType *string `type:"string"`

	ISP *string `type:"string"`

	SecurityProtectionInstanceId *int32 `type:"int32"`

	SecurityProtectionTypes []*string `type:"list"`
}

// String returns the string representation
func (s EipForModifyScalingConfigurationInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s EipForModifyScalingConfigurationInput) GoString() string {
	return s.String()
}

// SetBandwidth sets the Bandwidth field's value.
func (s *EipForModifyScalingConfigurationInput) SetBandwidth(v int32) *EipForModifyScalingConfigurationInput {
	s.Bandwidth = &v
	return s
}

// SetBandwidthPackageId sets the BandwidthPackageId field's value.
func (s *EipForModifyScalingConfigurationInput) SetBandwidthPackageId(v string) *EipForModifyScalingConfigurationInput {
	s.BandwidthPackageId = &v
	return s
}

// SetBillingType sets the BillingType field's value.
func (s *EipForModifyScalingConfigurationInput) SetBillingType(v string) *EipForModifyScalingConfigurationInput {
	s.BillingType = &v
	return s
}

// SetISP sets the ISP field's value.
func (s *EipForModifyScalingConfigurationInput) SetISP(v string) *EipForModifyScalingConfigurationInput {
	s.ISP = &v
	return s
}

// SetSecurityProtectionInstanceId sets the SecurityProtectionInstanceId field's value.
func (s *EipForModifyScalingConfigurationInput) SetSecurityProtectionInstanceId(v int32) *EipForModifyScalingConfigurationInput {
	s.SecurityProtectionInstanceId = &v
	return s
}

// SetSecurityProtectionTypes sets the SecurityProtectionTypes field's value.
func (s *EipForModifyScalingConfigurationInput) SetSecurityProtectionTypes(v []*string) *EipForModifyScalingConfigurationInput {
	s.SecurityProtectionTypes = v
	return s
}

type InstanceTypeOverrideForModifyScalingConfigurationInput struct {
	_ struct{} `type:"structure"`

	InstanceType *string `type:"string"`

	PriceLimit *float64 `type:"float"`
}

// String returns the string representation
func (s InstanceTypeOverrideForModifyScalingConfigurationInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s InstanceTypeOverrideForModifyScalingConfigurationInput) GoString() string {
	return s.String()
}

// SetInstanceType sets the InstanceType field's value.
func (s *InstanceTypeOverrideForModifyScalingConfigurationInput) SetInstanceType(v string) *InstanceTypeOverrideForModifyScalingConfigurationInput {
	s.InstanceType = &v
	return s
}

// SetPriceLimit sets the PriceLimit field's value.
func (s *InstanceTypeOverrideForModifyScalingConfigurationInput) SetPriceLimit(v float64) *InstanceTypeOverrideForModifyScalingConfigurationInput {
	s.PriceLimit = &v
	return s
}

type ModifyScalingConfigurationInput struct {
	_ struct{} `type:"structure"`

	Eip *EipForModifyScalingConfigurationInput `type:"structure"`

	HostName *string `type:"string"`

	HpcClusterId *string `type:"string"`

	ImageId *string `type:"string"`

	InstanceDescription *string `type:"string"`

	InstanceName *string `type:"string"`

	InstanceTypeOverrides []*InstanceTypeOverrideForModifyScalingConfigurationInput `type:"list"`

	InstanceTypes []*string `type:"list"`

	Ipv6AddressCount *int32 `type:"int32"`

	KeyPairName *string `type:"string"`

	Password *string `type:"string"`

	ProjectName *string `type:"string"`

	// ScalingConfigurationId is a required field
	ScalingConfigurationId *string `type:"string" required:"true"`

	ScalingConfigurationName *string `min:"1" max:"128" type:"string"`

	SecurityEnhancementStrategy *string `type:"string"`

	SecurityGroupIds []*string `type:"list"`

	SpotStrategy *string `type:"string"`

	Tags *string `type:"string"`

	UserData *string `type:"string"`

	Volumes []*VolumeForModifyScalingConfigurationInput `type:"list"`

	ZoneId *string `type:"string"`
}

// String returns the string representation
func (s ModifyScalingConfigurationInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyScalingConfigurationInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ModifyScalingConfigurationInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ModifyScalingConfigurationInput"}
	if s.ScalingConfigurationId == nil {
		invalidParams.Add(request.NewErrParamRequired("ScalingConfigurationId"))
	}
	if s.ScalingConfigurationName != nil && len(*s.ScalingConfigurationName) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("ScalingConfigurationName", 1))
	}
	if s.ScalingConfigurationName != nil && len(*s.ScalingConfigurationName) > 128 {
		invalidParams.Add(request.NewErrParamMaxLen("ScalingConfigurationName", 128, *s.ScalingConfigurationName))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetEip sets the Eip field's value.
func (s *ModifyScalingConfigurationInput) SetEip(v *EipForModifyScalingConfigurationInput) *ModifyScalingConfigurationInput {
	s.Eip = v
	return s
}

// SetHostName sets the HostName field's value.
func (s *ModifyScalingConfigurationInput) SetHostName(v string) *ModifyScalingConfigurationInput {
	s.HostName = &v
	return s
}

// SetHpcClusterId sets the HpcClusterId field's value.
func (s *ModifyScalingConfigurationInput) SetHpcClusterId(v string) *ModifyScalingConfigurationInput {
	s.HpcClusterId = &v
	return s
}

// SetImageId sets the ImageId field's value.
func (s *ModifyScalingConfigurationInput) SetImageId(v string) *ModifyScalingConfigurationInput {
	s.ImageId = &v
	return s
}

// SetInstanceDescription sets the InstanceDescription field's value.
func (s *ModifyScalingConfigurationInput) SetInstanceDescription(v string) *ModifyScalingConfigurationInput {
	s.InstanceDescription = &v
	return s
}

// SetInstanceName sets the InstanceName field's value.
func (s *ModifyScalingConfigurationInput) SetInstanceName(v string) *ModifyScalingConfigurationInput {
	s.InstanceName = &v
	return s
}

// SetInstanceTypeOverrides sets the InstanceTypeOverrides field's value.
func (s *ModifyScalingConfigurationInput) SetInstanceTypeOverrides(v []*InstanceTypeOverrideForModifyScalingConfigurationInput) *ModifyScalingConfigurationInput {
	s.InstanceTypeOverrides = v
	return s
}

// SetInstanceTypes sets the InstanceTypes field's value.
func (s *ModifyScalingConfigurationInput) SetInstanceTypes(v []*string) *ModifyScalingConfigurationInput {
	s.InstanceTypes = v
	return s
}

// SetIpv6AddressCount sets the Ipv6AddressCount field's value.
func (s *ModifyScalingConfigurationInput) SetIpv6AddressCount(v int32) *ModifyScalingConfigurationInput {
	s.Ipv6AddressCount = &v
	return s
}

// SetKeyPairName sets the KeyPairName field's value.
func (s *ModifyScalingConfigurationInput) SetKeyPairName(v string) *ModifyScalingConfigurationInput {
	s.KeyPairName = &v
	return s
}

// SetPassword sets the Password field's value.
func (s *ModifyScalingConfigurationInput) SetPassword(v string) *ModifyScalingConfigurationInput {
	s.Password = &v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *ModifyScalingConfigurationInput) SetProjectName(v string) *ModifyScalingConfigurationInput {
	s.ProjectName = &v
	return s
}

// SetScalingConfigurationId sets the ScalingConfigurationId field's value.
func (s *ModifyScalingConfigurationInput) SetScalingConfigurationId(v string) *ModifyScalingConfigurationInput {
	s.ScalingConfigurationId = &v
	return s
}

// SetScalingConfigurationName sets the ScalingConfigurationName field's value.
func (s *ModifyScalingConfigurationInput) SetScalingConfigurationName(v string) *ModifyScalingConfigurationInput {
	s.ScalingConfigurationName = &v
	return s
}

// SetSecurityEnhancementStrategy sets the SecurityEnhancementStrategy field's value.
func (s *ModifyScalingConfigurationInput) SetSecurityEnhancementStrategy(v string) *ModifyScalingConfigurationInput {
	s.SecurityEnhancementStrategy = &v
	return s
}

// SetSecurityGroupIds sets the SecurityGroupIds field's value.
func (s *ModifyScalingConfigurationInput) SetSecurityGroupIds(v []*string) *ModifyScalingConfigurationInput {
	s.SecurityGroupIds = v
	return s
}

// SetSpotStrategy sets the SpotStrategy field's value.
func (s *ModifyScalingConfigurationInput) SetSpotStrategy(v string) *ModifyScalingConfigurationInput {
	s.SpotStrategy = &v
	return s
}

// SetTags sets the Tags field's value.
func (s *ModifyScalingConfigurationInput) SetTags(v string) *ModifyScalingConfigurationInput {
	s.Tags = &v
	return s
}

// SetUserData sets the UserData field's value.
func (s *ModifyScalingConfigurationInput) SetUserData(v string) *ModifyScalingConfigurationInput {
	s.UserData = &v
	return s
}

// SetVolumes sets the Volumes field's value.
func (s *ModifyScalingConfigurationInput) SetVolumes(v []*VolumeForModifyScalingConfigurationInput) *ModifyScalingConfigurationInput {
	s.Volumes = v
	return s
}

// SetZoneId sets the ZoneId field's value.
func (s *ModifyScalingConfigurationInput) SetZoneId(v string) *ModifyScalingConfigurationInput {
	s.ZoneId = &v
	return s
}

type ModifyScalingConfigurationOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	ScalingConfigurationId *string `type:"string"`
}

// String returns the string representation
func (s ModifyScalingConfigurationOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyScalingConfigurationOutput) GoString() string {
	return s.String()
}

// SetScalingConfigurationId sets the ScalingConfigurationId field's value.
func (s *ModifyScalingConfigurationOutput) SetScalingConfigurationId(v string) *ModifyScalingConfigurationOutput {
	s.ScalingConfigurationId = &v
	return s
}

type VolumeForModifyScalingConfigurationInput struct {
	_ struct{} `type:"structure"`

	DeleteWithInstance *bool `type:"boolean"`

	Size *int32 `type:"int32"`

	VolumeType *string `type:"string"`
}

// String returns the string representation
func (s VolumeForModifyScalingConfigurationInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s VolumeForModifyScalingConfigurationInput) GoString() string {
	return s.String()
}

// SetDeleteWithInstance sets the DeleteWithInstance field's value.
func (s *VolumeForModifyScalingConfigurationInput) SetDeleteWithInstance(v bool) *VolumeForModifyScalingConfigurationInput {
	s.DeleteWithInstance = &v
	return s
}

// SetSize sets the Size field's value.
func (s *VolumeForModifyScalingConfigurationInput) SetSize(v int32) *VolumeForModifyScalingConfigurationInput {
	s.Size = &v
	return s
}

// SetVolumeType sets the VolumeType field's value.
func (s *VolumeForModifyScalingConfigurationInput) SetVolumeType(v string) *VolumeForModifyScalingConfigurationInput {
	s.VolumeType = &v
	return s
}

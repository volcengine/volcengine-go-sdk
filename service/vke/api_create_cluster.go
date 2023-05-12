// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package vke

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opCreateClusterCommon = "CreateCluster"

// CreateClusterCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateClusterCommon operation. The "output" return
// value will be populated with the CreateClusterCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateClusterCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateClusterCommon Send returns without error.
//
// See CreateClusterCommon for more information on using the CreateClusterCommon
// API call, and error handling.
//
//	// Example sending a request using the CreateClusterCommonRequest method.
//	req, resp := client.CreateClusterCommonRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *VKE) CreateClusterCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opCreateClusterCommon,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &map[string]interface{}{}
	}

	output = &map[string]interface{}{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// CreateClusterCommon API operation for VKE.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VKE's
// API operation CreateClusterCommon for usage and error information.
func (c *VKE) CreateClusterCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.CreateClusterCommonRequest(input)
	return out, req.Send()
}

// CreateClusterCommonWithContext is the same as CreateClusterCommon with the addition of
// the ability to pass a context and additional request options.
//
// See CreateClusterCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VKE) CreateClusterCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.CreateClusterCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opCreateCluster = "CreateCluster"

// CreateClusterRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateCluster operation. The "output" return
// value will be populated with the CreateClusterCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateClusterCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateClusterCommon Send returns without error.
//
// See CreateCluster for more information on using the CreateCluster
// API call, and error handling.
//
//	// Example sending a request using the CreateClusterRequest method.
//	req, resp := client.CreateClusterRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *VKE) CreateClusterRequest(input *CreateClusterInput) (req *request.Request, output *CreateClusterOutput) {
	op := &request.Operation{
		Name:       opCreateCluster,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateClusterInput{}
	}

	output = &CreateClusterOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// CreateCluster API operation for VKE.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VKE's
// API operation CreateCluster for usage and error information.
func (c *VKE) CreateCluster(input *CreateClusterInput) (*CreateClusterOutput, error) {
	req, out := c.CreateClusterRequest(input)
	return out, req.Send()
}

// CreateClusterWithContext is the same as CreateCluster with the addition of
// the ability to pass a context and additional request options.
//
// See CreateCluster for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VKE) CreateClusterWithContext(ctx volcengine.Context, input *CreateClusterInput, opts ...request.Option) (*CreateClusterOutput, error) {
	req, out := c.CreateClusterRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ApiServerPublicAccessConfigForCreateClusterInput struct {
	_ struct{} `type:"structure"`

	PublicAccessNetworkConfig *PublicAccessNetworkConfigForCreateClusterInput `type:"structure"`
}

// String returns the string representation
func (s ApiServerPublicAccessConfigForCreateClusterInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ApiServerPublicAccessConfigForCreateClusterInput) GoString() string {
	return s.String()
}

// SetPublicAccessNetworkConfig sets the PublicAccessNetworkConfig field's value.
func (s *ApiServerPublicAccessConfigForCreateClusterInput) SetPublicAccessNetworkConfig(v *PublicAccessNetworkConfigForCreateClusterInput) *ApiServerPublicAccessConfigForCreateClusterInput {
	s.PublicAccessNetworkConfig = v
	return s
}

type ClusterConfigForCreateClusterInput struct {
	_ struct{} `type:"structure"`

	ApiServerPublicAccessConfig *ApiServerPublicAccessConfigForCreateClusterInput `type:"structure"`

	ApiServerPublicAccessEnabled *bool `type:"boolean"`

	ResourcePublicAccessDefaultEnabled *bool `type:"boolean"`

	SubnetIds []*string `type:"list"`
}

// String returns the string representation
func (s ClusterConfigForCreateClusterInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ClusterConfigForCreateClusterInput) GoString() string {
	return s.String()
}

// SetApiServerPublicAccessConfig sets the ApiServerPublicAccessConfig field's value.
func (s *ClusterConfigForCreateClusterInput) SetApiServerPublicAccessConfig(v *ApiServerPublicAccessConfigForCreateClusterInput) *ClusterConfigForCreateClusterInput {
	s.ApiServerPublicAccessConfig = v
	return s
}

// SetApiServerPublicAccessEnabled sets the ApiServerPublicAccessEnabled field's value.
func (s *ClusterConfigForCreateClusterInput) SetApiServerPublicAccessEnabled(v bool) *ClusterConfigForCreateClusterInput {
	s.ApiServerPublicAccessEnabled = &v
	return s
}

// SetResourcePublicAccessDefaultEnabled sets the ResourcePublicAccessDefaultEnabled field's value.
func (s *ClusterConfigForCreateClusterInput) SetResourcePublicAccessDefaultEnabled(v bool) *ClusterConfigForCreateClusterInput {
	s.ResourcePublicAccessDefaultEnabled = &v
	return s
}

// SetSubnetIds sets the SubnetIds field's value.
func (s *ClusterConfigForCreateClusterInput) SetSubnetIds(v []*string) *ClusterConfigForCreateClusterInput {
	s.SubnetIds = v
	return s
}

type CreateClusterInput struct {
	_ struct{} `type:"structure"`

	ClientToken *string `type:"string"`

	ClusterConfig *ClusterConfigForCreateClusterInput `type:"structure"`

	DeleteProtectionEnabled *bool `type:"boolean"`

	Description *string `type:"string"`

	KubernetesVersion *string `type:"string"`

	LoggingConfig *LoggingConfigForCreateClusterInput `type:"structure"`

	Name *string `type:"string"`

	PodsConfig *PodsConfigForCreateClusterInput `type:"structure"`

	ServicesConfig *ServicesConfigForCreateClusterInput `type:"structure"`

	Tags []*TagForCreateClusterInput `type:"list"`
}

// String returns the string representation
func (s CreateClusterInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateClusterInput) GoString() string {
	return s.String()
}

// SetClientToken sets the ClientToken field's value.
func (s *CreateClusterInput) SetClientToken(v string) *CreateClusterInput {
	s.ClientToken = &v
	return s
}

// SetClusterConfig sets the ClusterConfig field's value.
func (s *CreateClusterInput) SetClusterConfig(v *ClusterConfigForCreateClusterInput) *CreateClusterInput {
	s.ClusterConfig = v
	return s
}

// SetDeleteProtectionEnabled sets the DeleteProtectionEnabled field's value.
func (s *CreateClusterInput) SetDeleteProtectionEnabled(v bool) *CreateClusterInput {
	s.DeleteProtectionEnabled = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *CreateClusterInput) SetDescription(v string) *CreateClusterInput {
	s.Description = &v
	return s
}

// SetKubernetesVersion sets the KubernetesVersion field's value.
func (s *CreateClusterInput) SetKubernetesVersion(v string) *CreateClusterInput {
	s.KubernetesVersion = &v
	return s
}

// SetLoggingConfig sets the LoggingConfig field's value.
func (s *CreateClusterInput) SetLoggingConfig(v *LoggingConfigForCreateClusterInput) *CreateClusterInput {
	s.LoggingConfig = v
	return s
}

// SetName sets the Name field's value.
func (s *CreateClusterInput) SetName(v string) *CreateClusterInput {
	s.Name = &v
	return s
}

// SetPodsConfig sets the PodsConfig field's value.
func (s *CreateClusterInput) SetPodsConfig(v *PodsConfigForCreateClusterInput) *CreateClusterInput {
	s.PodsConfig = v
	return s
}

// SetServicesConfig sets the ServicesConfig field's value.
func (s *CreateClusterInput) SetServicesConfig(v *ServicesConfigForCreateClusterInput) *CreateClusterInput {
	s.ServicesConfig = v
	return s
}

// SetTags sets the Tags field's value.
func (s *CreateClusterInput) SetTags(v []*TagForCreateClusterInput) *CreateClusterInput {
	s.Tags = v
	return s
}

type CreateClusterOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	Id *string `type:"string"`
}

// String returns the string representation
func (s CreateClusterOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateClusterOutput) GoString() string {
	return s.String()
}

// SetId sets the Id field's value.
func (s *CreateClusterOutput) SetId(v string) *CreateClusterOutput {
	s.Id = &v
	return s
}

type FlannelConfigForCreateClusterInput struct {
	_ struct{} `type:"structure"`

	MaxPodsPerNode *int32 `type:"int32"`

	PodCidrs []*string `type:"list"`
}

// String returns the string representation
func (s FlannelConfigForCreateClusterInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s FlannelConfigForCreateClusterInput) GoString() string {
	return s.String()
}

// SetMaxPodsPerNode sets the MaxPodsPerNode field's value.
func (s *FlannelConfigForCreateClusterInput) SetMaxPodsPerNode(v int32) *FlannelConfigForCreateClusterInput {
	s.MaxPodsPerNode = &v
	return s
}

// SetPodCidrs sets the PodCidrs field's value.
func (s *FlannelConfigForCreateClusterInput) SetPodCidrs(v []*string) *FlannelConfigForCreateClusterInput {
	s.PodCidrs = v
	return s
}

type LogSetupForCreateClusterInput struct {
	_ struct{} `type:"structure"`

	Enabled *bool `type:"boolean"`

	LogTtl *int32 `type:"int32"`

	LogType *string `type:"string" enum:"EnumOfLogTypeForCreateClusterInput"`
}

// String returns the string representation
func (s LogSetupForCreateClusterInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s LogSetupForCreateClusterInput) GoString() string {
	return s.String()
}

// SetEnabled sets the Enabled field's value.
func (s *LogSetupForCreateClusterInput) SetEnabled(v bool) *LogSetupForCreateClusterInput {
	s.Enabled = &v
	return s
}

// SetLogTtl sets the LogTtl field's value.
func (s *LogSetupForCreateClusterInput) SetLogTtl(v int32) *LogSetupForCreateClusterInput {
	s.LogTtl = &v
	return s
}

// SetLogType sets the LogType field's value.
func (s *LogSetupForCreateClusterInput) SetLogType(v string) *LogSetupForCreateClusterInput {
	s.LogType = &v
	return s
}

type LoggingConfigForCreateClusterInput struct {
	_ struct{} `type:"structure"`

	LogProjectId *string `type:"string"`

	LogSetups []*LogSetupForCreateClusterInput `type:"list"`
}

// String returns the string representation
func (s LoggingConfigForCreateClusterInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s LoggingConfigForCreateClusterInput) GoString() string {
	return s.String()
}

// SetLogProjectId sets the LogProjectId field's value.
func (s *LoggingConfigForCreateClusterInput) SetLogProjectId(v string) *LoggingConfigForCreateClusterInput {
	s.LogProjectId = &v
	return s
}

// SetLogSetups sets the LogSetups field's value.
func (s *LoggingConfigForCreateClusterInput) SetLogSetups(v []*LogSetupForCreateClusterInput) *LoggingConfigForCreateClusterInput {
	s.LogSetups = v
	return s
}

type PodsConfigForCreateClusterInput struct {
	_ struct{} `type:"structure"`

	FlannelConfig *FlannelConfigForCreateClusterInput `type:"structure"`

	PodNetworkMode *string `type:"string" enum:"EnumOfPodNetworkModeForCreateClusterInput"`

	VpcCniConfig *VpcCniConfigForCreateClusterInput `type:"structure"`
}

// String returns the string representation
func (s PodsConfigForCreateClusterInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s PodsConfigForCreateClusterInput) GoString() string {
	return s.String()
}

// SetFlannelConfig sets the FlannelConfig field's value.
func (s *PodsConfigForCreateClusterInput) SetFlannelConfig(v *FlannelConfigForCreateClusterInput) *PodsConfigForCreateClusterInput {
	s.FlannelConfig = v
	return s
}

// SetPodNetworkMode sets the PodNetworkMode field's value.
func (s *PodsConfigForCreateClusterInput) SetPodNetworkMode(v string) *PodsConfigForCreateClusterInput {
	s.PodNetworkMode = &v
	return s
}

// SetVpcCniConfig sets the VpcCniConfig field's value.
func (s *PodsConfigForCreateClusterInput) SetVpcCniConfig(v *VpcCniConfigForCreateClusterInput) *PodsConfigForCreateClusterInput {
	s.VpcCniConfig = v
	return s
}

type PublicAccessNetworkConfigForCreateClusterInput struct {
	_ struct{} `type:"structure"`

	Bandwidth *int32 `type:"int32"`

	BillingType *int32 `type:"int32"`
}

// String returns the string representation
func (s PublicAccessNetworkConfigForCreateClusterInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s PublicAccessNetworkConfigForCreateClusterInput) GoString() string {
	return s.String()
}

// SetBandwidth sets the Bandwidth field's value.
func (s *PublicAccessNetworkConfigForCreateClusterInput) SetBandwidth(v int32) *PublicAccessNetworkConfigForCreateClusterInput {
	s.Bandwidth = &v
	return s
}

// SetBillingType sets the BillingType field's value.
func (s *PublicAccessNetworkConfigForCreateClusterInput) SetBillingType(v int32) *PublicAccessNetworkConfigForCreateClusterInput {
	s.BillingType = &v
	return s
}

type ServicesConfigForCreateClusterInput struct {
	_ struct{} `type:"structure"`

	ServiceCidrsv4 []*string `type:"list"`
}

// String returns the string representation
func (s ServicesConfigForCreateClusterInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ServicesConfigForCreateClusterInput) GoString() string {
	return s.String()
}

// SetServiceCidrsv4 sets the ServiceCidrsv4 field's value.
func (s *ServicesConfigForCreateClusterInput) SetServiceCidrsv4(v []*string) *ServicesConfigForCreateClusterInput {
	s.ServiceCidrsv4 = v
	return s
}

type TagForCreateClusterInput struct {
	_ struct{} `type:"structure"`

	Key *string `type:"string"`

	Value *string `type:"string"`
}

// String returns the string representation
func (s TagForCreateClusterInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s TagForCreateClusterInput) GoString() string {
	return s.String()
}

// SetKey sets the Key field's value.
func (s *TagForCreateClusterInput) SetKey(v string) *TagForCreateClusterInput {
	s.Key = &v
	return s
}

// SetValue sets the Value field's value.
func (s *TagForCreateClusterInput) SetValue(v string) *TagForCreateClusterInput {
	s.Value = &v
	return s
}

type VpcCniConfigForCreateClusterInput struct {
	_ struct{} `type:"structure"`

	SubnetIds []*string `type:"list"`
}

// String returns the string representation
func (s VpcCniConfigForCreateClusterInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s VpcCniConfigForCreateClusterInput) GoString() string {
	return s.String()
}

// SetSubnetIds sets the SubnetIds field's value.
func (s *VpcCniConfigForCreateClusterInput) SetSubnetIds(v []*string) *VpcCniConfigForCreateClusterInput {
	s.SubnetIds = v
	return s
}

const (
	// EnumOfLogTypeForCreateClusterInputAudit is a EnumOfLogTypeForCreateClusterInput enum value
	EnumOfLogTypeForCreateClusterInputAudit = "Audit"
)

const (
	// EnumOfPodNetworkModeForCreateClusterInputCalicoBgp is a EnumOfPodNetworkModeForCreateClusterInput enum value
	EnumOfPodNetworkModeForCreateClusterInputCalicoBgp = "CalicoBgp"

	// EnumOfPodNetworkModeForCreateClusterInputCalicoVxlan is a EnumOfPodNetworkModeForCreateClusterInput enum value
	EnumOfPodNetworkModeForCreateClusterInputCalicoVxlan = "CalicoVxlan"

	// EnumOfPodNetworkModeForCreateClusterInputCarma is a EnumOfPodNetworkModeForCreateClusterInput enum value
	EnumOfPodNetworkModeForCreateClusterInputCarma = "Carma"

	// EnumOfPodNetworkModeForCreateClusterInputCilium is a EnumOfPodNetworkModeForCreateClusterInput enum value
	EnumOfPodNetworkModeForCreateClusterInputCilium = "Cilium"

	// EnumOfPodNetworkModeForCreateClusterInputDefault is a EnumOfPodNetworkModeForCreateClusterInput enum value
	EnumOfPodNetworkModeForCreateClusterInputDefault = "Default"

	// EnumOfPodNetworkModeForCreateClusterInputFlannel is a EnumOfPodNetworkModeForCreateClusterInput enum value
	EnumOfPodNetworkModeForCreateClusterInputFlannel = "Flannel"

	// EnumOfPodNetworkModeForCreateClusterInputKubeOvn is a EnumOfPodNetworkModeForCreateClusterInput enum value
	EnumOfPodNetworkModeForCreateClusterInputKubeOvn = "KubeOvn"

	// EnumOfPodNetworkModeForCreateClusterInputVpcCniDedicated is a EnumOfPodNetworkModeForCreateClusterInput enum value
	EnumOfPodNetworkModeForCreateClusterInputVpcCniDedicated = "VpcCniDedicated"

	// EnumOfPodNetworkModeForCreateClusterInputVpcCniShared is a EnumOfPodNetworkModeForCreateClusterInput enum value
	EnumOfPodNetworkModeForCreateClusterInputVpcCniShared = "VpcCniShared"
)

// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package vke

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opListNodePoolsCommon = "ListNodePools"

// ListNodePoolsCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ListNodePoolsCommon operation. The "output" return
// value will be populated with the ListNodePoolsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListNodePoolsCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListNodePoolsCommon Send returns without error.
//
// See ListNodePoolsCommon for more information on using the ListNodePoolsCommon
// API call, and error handling.
//
//    // Example sending a request using the ListNodePoolsCommonRequest method.
//    req, resp := client.ListNodePoolsCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VKE) ListNodePoolsCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opListNodePoolsCommon,
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

// ListNodePoolsCommon API operation for VKE.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VKE's
// API operation ListNodePoolsCommon for usage and error information.
func (c *VKE) ListNodePoolsCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ListNodePoolsCommonRequest(input)
	return out, req.Send()
}

// ListNodePoolsCommonWithContext is the same as ListNodePoolsCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ListNodePoolsCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VKE) ListNodePoolsCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ListNodePoolsCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opListNodePools = "ListNodePools"

// ListNodePoolsRequest generates a "volcengine/request.Request" representing the
// client's request for the ListNodePools operation. The "output" return
// value will be populated with the ListNodePoolsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListNodePoolsCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListNodePoolsCommon Send returns without error.
//
// See ListNodePools for more information on using the ListNodePools
// API call, and error handling.
//
//    // Example sending a request using the ListNodePoolsRequest method.
//    req, resp := client.ListNodePoolsRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VKE) ListNodePoolsRequest(input *ListNodePoolsInput) (req *request.Request, output *ListNodePoolsOutput) {
	op := &request.Operation{
		Name:       opListNodePools,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListNodePoolsInput{}
	}

	output = &ListNodePoolsOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// ListNodePools API operation for VKE.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VKE's
// API operation ListNodePools for usage and error information.
func (c *VKE) ListNodePools(input *ListNodePoolsInput) (*ListNodePoolsOutput, error) {
	req, out := c.ListNodePoolsRequest(input)
	return out, req.Send()
}

// ListNodePoolsWithContext is the same as ListNodePools with the addition of
// the ability to pass a context and additional request options.
//
// See ListNodePools for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VKE) ListNodePoolsWithContext(ctx volcengine.Context, input *ListNodePoolsInput, opts ...request.Option) (*ListNodePoolsOutput, error) {
	req, out := c.ListNodePoolsRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type AutoScalingForListNodePoolsOutput struct {
	_ struct{} `type:"structure"`

	DesiredReplicas *int32 `type:"int32"`

	Enabled *bool `type:"boolean"`

	MaxReplicas *int32 `type:"int32"`

	MinReplicas *int32 `type:"int32"`

	Priority *int32 `type:"int32"`

	SubnetPolicy *string `type:"string" enum:"EnumOfSubnetPolicyForListNodePoolsOutput"`
}

// String returns the string representation
func (s AutoScalingForListNodePoolsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s AutoScalingForListNodePoolsOutput) GoString() string {
	return s.String()
}

// SetDesiredReplicas sets the DesiredReplicas field's value.
func (s *AutoScalingForListNodePoolsOutput) SetDesiredReplicas(v int32) *AutoScalingForListNodePoolsOutput {
	s.DesiredReplicas = &v
	return s
}

// SetEnabled sets the Enabled field's value.
func (s *AutoScalingForListNodePoolsOutput) SetEnabled(v bool) *AutoScalingForListNodePoolsOutput {
	s.Enabled = &v
	return s
}

// SetMaxReplicas sets the MaxReplicas field's value.
func (s *AutoScalingForListNodePoolsOutput) SetMaxReplicas(v int32) *AutoScalingForListNodePoolsOutput {
	s.MaxReplicas = &v
	return s
}

// SetMinReplicas sets the MinReplicas field's value.
func (s *AutoScalingForListNodePoolsOutput) SetMinReplicas(v int32) *AutoScalingForListNodePoolsOutput {
	s.MinReplicas = &v
	return s
}

// SetPriority sets the Priority field's value.
func (s *AutoScalingForListNodePoolsOutput) SetPriority(v int32) *AutoScalingForListNodePoolsOutput {
	s.Priority = &v
	return s
}

// SetSubnetPolicy sets the SubnetPolicy field's value.
func (s *AutoScalingForListNodePoolsOutput) SetSubnetPolicy(v string) *AutoScalingForListNodePoolsOutput {
	s.SubnetPolicy = &v
	return s
}

type ConditionForListNodePoolsOutput struct {
	_ struct{} `type:"structure"`

	Type *string `type:"string"`
}

// String returns the string representation
func (s ConditionForListNodePoolsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ConditionForListNodePoolsOutput) GoString() string {
	return s.String()
}

// SetType sets the Type field's value.
func (s *ConditionForListNodePoolsOutput) SetType(v string) *ConditionForListNodePoolsOutput {
	s.Type = &v
	return s
}

type ConvertTagForListNodePoolsOutput struct {
	_ struct{} `type:"structure"`

	Key *string `type:"string"`

	Type *string `type:"string"`

	Value *string `type:"string"`
}

// String returns the string representation
func (s ConvertTagForListNodePoolsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ConvertTagForListNodePoolsOutput) GoString() string {
	return s.String()
}

// SetKey sets the Key field's value.
func (s *ConvertTagForListNodePoolsOutput) SetKey(v string) *ConvertTagForListNodePoolsOutput {
	s.Key = &v
	return s
}

// SetType sets the Type field's value.
func (s *ConvertTagForListNodePoolsOutput) SetType(v string) *ConvertTagForListNodePoolsOutput {
	s.Type = &v
	return s
}

// SetValue sets the Value field's value.
func (s *ConvertTagForListNodePoolsOutput) SetValue(v string) *ConvertTagForListNodePoolsOutput {
	s.Value = &v
	return s
}

type DataVolumeForListNodePoolsOutput struct {
	_ struct{} `type:"structure"`

	MountPoint *string `type:"string"`

	Size *int32 `type:"int32"`

	Type *string `type:"string" enum:"EnumOfTypeForListNodePoolsOutput"`
}

// String returns the string representation
func (s DataVolumeForListNodePoolsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DataVolumeForListNodePoolsOutput) GoString() string {
	return s.String()
}

// SetMountPoint sets the MountPoint field's value.
func (s *DataVolumeForListNodePoolsOutput) SetMountPoint(v string) *DataVolumeForListNodePoolsOutput {
	s.MountPoint = &v
	return s
}

// SetSize sets the Size field's value.
func (s *DataVolumeForListNodePoolsOutput) SetSize(v int32) *DataVolumeForListNodePoolsOutput {
	s.Size = &v
	return s
}

// SetType sets the Type field's value.
func (s *DataVolumeForListNodePoolsOutput) SetType(v string) *DataVolumeForListNodePoolsOutput {
	s.Type = &v
	return s
}

type FilterForListNodePoolsInput struct {
	_ struct{} `type:"structure"`

	AutoScalingEnabled *bool `type:"boolean" json:"AutoScaling.Enabled"`

	ClusterIds []*string `type:"list"`

	CreateClientToken *string `type:"string"`

	Ids []*string `type:"list"`

	Name *string `type:"string"`

	Statuses []*StatusForListNodePoolsInput `type:"list"`

	UpdateClientToken *string `type:"string"`
}

// String returns the string representation
func (s FilterForListNodePoolsInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s FilterForListNodePoolsInput) GoString() string {
	return s.String()
}

// SetAutoScalingEnabled sets the AutoScalingEnabled field's value.
func (s *FilterForListNodePoolsInput) SetAutoScalingEnabled(v bool) *FilterForListNodePoolsInput {
	s.AutoScalingEnabled = &v
	return s
}

// SetClusterIds sets the ClusterIds field's value.
func (s *FilterForListNodePoolsInput) SetClusterIds(v []*string) *FilterForListNodePoolsInput {
	s.ClusterIds = v
	return s
}

// SetCreateClientToken sets the CreateClientToken field's value.
func (s *FilterForListNodePoolsInput) SetCreateClientToken(v string) *FilterForListNodePoolsInput {
	s.CreateClientToken = &v
	return s
}

// SetIds sets the Ids field's value.
func (s *FilterForListNodePoolsInput) SetIds(v []*string) *FilterForListNodePoolsInput {
	s.Ids = v
	return s
}

// SetName sets the Name field's value.
func (s *FilterForListNodePoolsInput) SetName(v string) *FilterForListNodePoolsInput {
	s.Name = &v
	return s
}

// SetStatuses sets the Statuses field's value.
func (s *FilterForListNodePoolsInput) SetStatuses(v []*StatusForListNodePoolsInput) *FilterForListNodePoolsInput {
	s.Statuses = v
	return s
}

// SetUpdateClientToken sets the UpdateClientToken field's value.
func (s *FilterForListNodePoolsInput) SetUpdateClientToken(v string) *FilterForListNodePoolsInput {
	s.UpdateClientToken = &v
	return s
}

type ItemForListNodePoolsOutput struct {
	_ struct{} `type:"structure"`

	AutoScaling *AutoScalingForListNodePoolsOutput `type:"structure"`

	ClusterId *string `type:"string"`

	CreateClientToken *string `type:"string"`

	CreateTime *string `type:"string"`

	Id *string `type:"string"`

	KubernetesConfig *KubernetesConfigForListNodePoolsOutput `type:"structure"`

	Name *string `type:"string"`

	NodeConfig *NodeConfigForListNodePoolsOutput `type:"structure"`

	NodeStatistics *NodeStatisticsForListNodePoolsOutput `type:"structure"`

	Status *StatusForListNodePoolsOutput `type:"structure"`

	Tags []*ConvertTagForListNodePoolsOutput `type:"list"`

	UpdateClientToken *string `type:"string"`

	UpdateTime *string `type:"string"`
}

// String returns the string representation
func (s ItemForListNodePoolsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ItemForListNodePoolsOutput) GoString() string {
	return s.String()
}

// SetAutoScaling sets the AutoScaling field's value.
func (s *ItemForListNodePoolsOutput) SetAutoScaling(v *AutoScalingForListNodePoolsOutput) *ItemForListNodePoolsOutput {
	s.AutoScaling = v
	return s
}

// SetClusterId sets the ClusterId field's value.
func (s *ItemForListNodePoolsOutput) SetClusterId(v string) *ItemForListNodePoolsOutput {
	s.ClusterId = &v
	return s
}

// SetCreateClientToken sets the CreateClientToken field's value.
func (s *ItemForListNodePoolsOutput) SetCreateClientToken(v string) *ItemForListNodePoolsOutput {
	s.CreateClientToken = &v
	return s
}

// SetCreateTime sets the CreateTime field's value.
func (s *ItemForListNodePoolsOutput) SetCreateTime(v string) *ItemForListNodePoolsOutput {
	s.CreateTime = &v
	return s
}

// SetId sets the Id field's value.
func (s *ItemForListNodePoolsOutput) SetId(v string) *ItemForListNodePoolsOutput {
	s.Id = &v
	return s
}

// SetKubernetesConfig sets the KubernetesConfig field's value.
func (s *ItemForListNodePoolsOutput) SetKubernetesConfig(v *KubernetesConfigForListNodePoolsOutput) *ItemForListNodePoolsOutput {
	s.KubernetesConfig = v
	return s
}

// SetName sets the Name field's value.
func (s *ItemForListNodePoolsOutput) SetName(v string) *ItemForListNodePoolsOutput {
	s.Name = &v
	return s
}

// SetNodeConfig sets the NodeConfig field's value.
func (s *ItemForListNodePoolsOutput) SetNodeConfig(v *NodeConfigForListNodePoolsOutput) *ItemForListNodePoolsOutput {
	s.NodeConfig = v
	return s
}

// SetNodeStatistics sets the NodeStatistics field's value.
func (s *ItemForListNodePoolsOutput) SetNodeStatistics(v *NodeStatisticsForListNodePoolsOutput) *ItemForListNodePoolsOutput {
	s.NodeStatistics = v
	return s
}

// SetStatus sets the Status field's value.
func (s *ItemForListNodePoolsOutput) SetStatus(v *StatusForListNodePoolsOutput) *ItemForListNodePoolsOutput {
	s.Status = v
	return s
}

// SetTags sets the Tags field's value.
func (s *ItemForListNodePoolsOutput) SetTags(v []*ConvertTagForListNodePoolsOutput) *ItemForListNodePoolsOutput {
	s.Tags = v
	return s
}

// SetUpdateClientToken sets the UpdateClientToken field's value.
func (s *ItemForListNodePoolsOutput) SetUpdateClientToken(v string) *ItemForListNodePoolsOutput {
	s.UpdateClientToken = &v
	return s
}

// SetUpdateTime sets the UpdateTime field's value.
func (s *ItemForListNodePoolsOutput) SetUpdateTime(v string) *ItemForListNodePoolsOutput {
	s.UpdateTime = &v
	return s
}

type KubernetesConfigForListNodePoolsOutput struct {
	_ struct{} `type:"structure"`

	Cordon *bool `type:"boolean"`

	Labels []*LabelForListNodePoolsOutput `type:"list"`

	Taints []*TaintForListNodePoolsOutput `type:"list"`
}

// String returns the string representation
func (s KubernetesConfigForListNodePoolsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s KubernetesConfigForListNodePoolsOutput) GoString() string {
	return s.String()
}

// SetCordon sets the Cordon field's value.
func (s *KubernetesConfigForListNodePoolsOutput) SetCordon(v bool) *KubernetesConfigForListNodePoolsOutput {
	s.Cordon = &v
	return s
}

// SetLabels sets the Labels field's value.
func (s *KubernetesConfigForListNodePoolsOutput) SetLabels(v []*LabelForListNodePoolsOutput) *KubernetesConfigForListNodePoolsOutput {
	s.Labels = v
	return s
}

// SetTaints sets the Taints field's value.
func (s *KubernetesConfigForListNodePoolsOutput) SetTaints(v []*TaintForListNodePoolsOutput) *KubernetesConfigForListNodePoolsOutput {
	s.Taints = v
	return s
}

type LabelForListNodePoolsOutput struct {
	_ struct{} `type:"structure"`

	Key *string `type:"string"`

	Value *string `type:"string"`
}

// String returns the string representation
func (s LabelForListNodePoolsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s LabelForListNodePoolsOutput) GoString() string {
	return s.String()
}

// SetKey sets the Key field's value.
func (s *LabelForListNodePoolsOutput) SetKey(v string) *LabelForListNodePoolsOutput {
	s.Key = &v
	return s
}

// SetValue sets the Value field's value.
func (s *LabelForListNodePoolsOutput) SetValue(v string) *LabelForListNodePoolsOutput {
	s.Value = &v
	return s
}

type ListNodePoolsInput struct {
	_ struct{} `type:"structure"`

	Filter *FilterForListNodePoolsInput `type:"structure"`

	PageNumber *int32 `type:"int32"`

	PageSize *int32 `type:"int32"`

	Tags []*TagForListNodePoolsInput `type:"list"`
}

// String returns the string representation
func (s ListNodePoolsInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListNodePoolsInput) GoString() string {
	return s.String()
}

// SetFilter sets the Filter field's value.
func (s *ListNodePoolsInput) SetFilter(v *FilterForListNodePoolsInput) *ListNodePoolsInput {
	s.Filter = v
	return s
}

// SetPageNumber sets the PageNumber field's value.
func (s *ListNodePoolsInput) SetPageNumber(v int32) *ListNodePoolsInput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *ListNodePoolsInput) SetPageSize(v int32) *ListNodePoolsInput {
	s.PageSize = &v
	return s
}

// SetTags sets the Tags field's value.
func (s *ListNodePoolsInput) SetTags(v []*TagForListNodePoolsInput) *ListNodePoolsInput {
	s.Tags = v
	return s
}

type ListNodePoolsOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	Items []*ItemForListNodePoolsOutput `type:"list"`

	PageNumber *int32 `type:"int32"`

	PageSize *int32 `type:"int32"`

	TotalCount *int32 `type:"int32"`
}

// String returns the string representation
func (s ListNodePoolsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListNodePoolsOutput) GoString() string {
	return s.String()
}

// SetItems sets the Items field's value.
func (s *ListNodePoolsOutput) SetItems(v []*ItemForListNodePoolsOutput) *ListNodePoolsOutput {
	s.Items = v
	return s
}

// SetPageNumber sets the PageNumber field's value.
func (s *ListNodePoolsOutput) SetPageNumber(v int32) *ListNodePoolsOutput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *ListNodePoolsOutput) SetPageSize(v int32) *ListNodePoolsOutput {
	s.PageSize = &v
	return s
}

// SetTotalCount sets the TotalCount field's value.
func (s *ListNodePoolsOutput) SetTotalCount(v int32) *ListNodePoolsOutput {
	s.TotalCount = &v
	return s
}

type LoginForListNodePoolsOutput struct {
	_ struct{} `type:"structure"`

	SshKeyPairName *string `type:"string"`

	Type *string `type:"string"`
}

// String returns the string representation
func (s LoginForListNodePoolsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s LoginForListNodePoolsOutput) GoString() string {
	return s.String()
}

// SetSshKeyPairName sets the SshKeyPairName field's value.
func (s *LoginForListNodePoolsOutput) SetSshKeyPairName(v string) *LoginForListNodePoolsOutput {
	s.SshKeyPairName = &v
	return s
}

// SetType sets the Type field's value.
func (s *LoginForListNodePoolsOutput) SetType(v string) *LoginForListNodePoolsOutput {
	s.Type = &v
	return s
}

type NodeConfigForListNodePoolsOutput struct {
	_ struct{} `type:"structure"`

	AdditionalContainerStorageEnabled *bool `type:"boolean"`

	AutoRenew *bool `type:"boolean"`

	AutoRenewPeriod *int32 `type:"int32"`

	DataVolumes []*DataVolumeForListNodePoolsOutput `type:"list"`

	HpcClusterIds []*string `type:"list"`

	ImageId *string `type:"string"`

	InitializeScript *string `type:"string"`

	InstanceChargeType *string `type:"string"`

	InstanceTypeIds []*string `type:"list"`

	NamePrefix *string `type:"string"`

	Period *int32 `type:"int32"`

	Security *SecurityForListNodePoolsOutput `type:"structure"`

	SubnetIds []*string `type:"list"`

	SystemVolume *SystemVolumeForListNodePoolsOutput `type:"structure"`

	Tags []*TagForListNodePoolsOutput `type:"list"`
}

// String returns the string representation
func (s NodeConfigForListNodePoolsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s NodeConfigForListNodePoolsOutput) GoString() string {
	return s.String()
}

// SetAdditionalContainerStorageEnabled sets the AdditionalContainerStorageEnabled field's value.
func (s *NodeConfigForListNodePoolsOutput) SetAdditionalContainerStorageEnabled(v bool) *NodeConfigForListNodePoolsOutput {
	s.AdditionalContainerStorageEnabled = &v
	return s
}

// SetAutoRenew sets the AutoRenew field's value.
func (s *NodeConfigForListNodePoolsOutput) SetAutoRenew(v bool) *NodeConfigForListNodePoolsOutput {
	s.AutoRenew = &v
	return s
}

// SetAutoRenewPeriod sets the AutoRenewPeriod field's value.
func (s *NodeConfigForListNodePoolsOutput) SetAutoRenewPeriod(v int32) *NodeConfigForListNodePoolsOutput {
	s.AutoRenewPeriod = &v
	return s
}

// SetDataVolumes sets the DataVolumes field's value.
func (s *NodeConfigForListNodePoolsOutput) SetDataVolumes(v []*DataVolumeForListNodePoolsOutput) *NodeConfigForListNodePoolsOutput {
	s.DataVolumes = v
	return s
}

// SetHpcClusterIds sets the HpcClusterIds field's value.
func (s *NodeConfigForListNodePoolsOutput) SetHpcClusterIds(v []*string) *NodeConfigForListNodePoolsOutput {
	s.HpcClusterIds = v
	return s
}

// SetImageId sets the ImageId field's value.
func (s *NodeConfigForListNodePoolsOutput) SetImageId(v string) *NodeConfigForListNodePoolsOutput {
	s.ImageId = &v
	return s
}

// SetInitializeScript sets the InitializeScript field's value.
func (s *NodeConfigForListNodePoolsOutput) SetInitializeScript(v string) *NodeConfigForListNodePoolsOutput {
	s.InitializeScript = &v
	return s
}

// SetInstanceChargeType sets the InstanceChargeType field's value.
func (s *NodeConfigForListNodePoolsOutput) SetInstanceChargeType(v string) *NodeConfigForListNodePoolsOutput {
	s.InstanceChargeType = &v
	return s
}

// SetInstanceTypeIds sets the InstanceTypeIds field's value.
func (s *NodeConfigForListNodePoolsOutput) SetInstanceTypeIds(v []*string) *NodeConfigForListNodePoolsOutput {
	s.InstanceTypeIds = v
	return s
}

// SetNamePrefix sets the NamePrefix field's value.
func (s *NodeConfigForListNodePoolsOutput) SetNamePrefix(v string) *NodeConfigForListNodePoolsOutput {
	s.NamePrefix = &v
	return s
}

// SetPeriod sets the Period field's value.
func (s *NodeConfigForListNodePoolsOutput) SetPeriod(v int32) *NodeConfigForListNodePoolsOutput {
	s.Period = &v
	return s
}

// SetSecurity sets the Security field's value.
func (s *NodeConfigForListNodePoolsOutput) SetSecurity(v *SecurityForListNodePoolsOutput) *NodeConfigForListNodePoolsOutput {
	s.Security = v
	return s
}

// SetSubnetIds sets the SubnetIds field's value.
func (s *NodeConfigForListNodePoolsOutput) SetSubnetIds(v []*string) *NodeConfigForListNodePoolsOutput {
	s.SubnetIds = v
	return s
}

// SetSystemVolume sets the SystemVolume field's value.
func (s *NodeConfigForListNodePoolsOutput) SetSystemVolume(v *SystemVolumeForListNodePoolsOutput) *NodeConfigForListNodePoolsOutput {
	s.SystemVolume = v
	return s
}

// SetTags sets the Tags field's value.
func (s *NodeConfigForListNodePoolsOutput) SetTags(v []*TagForListNodePoolsOutput) *NodeConfigForListNodePoolsOutput {
	s.Tags = v
	return s
}

type NodeStatisticsForListNodePoolsOutput struct {
	_ struct{} `type:"structure"`

	CreatingCount *int32 `type:"int32"`

	DeletingCount *int32 `type:"int32"`

	FailedCount *int32 `type:"int32"`

	RunningCount *int32 `type:"int32"`

	TotalCount *int32 `type:"int32"`

	UpdatingCount *int32 `type:"int32"`
}

// String returns the string representation
func (s NodeStatisticsForListNodePoolsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s NodeStatisticsForListNodePoolsOutput) GoString() string {
	return s.String()
}

// SetCreatingCount sets the CreatingCount field's value.
func (s *NodeStatisticsForListNodePoolsOutput) SetCreatingCount(v int32) *NodeStatisticsForListNodePoolsOutput {
	s.CreatingCount = &v
	return s
}

// SetDeletingCount sets the DeletingCount field's value.
func (s *NodeStatisticsForListNodePoolsOutput) SetDeletingCount(v int32) *NodeStatisticsForListNodePoolsOutput {
	s.DeletingCount = &v
	return s
}

// SetFailedCount sets the FailedCount field's value.
func (s *NodeStatisticsForListNodePoolsOutput) SetFailedCount(v int32) *NodeStatisticsForListNodePoolsOutput {
	s.FailedCount = &v
	return s
}

// SetRunningCount sets the RunningCount field's value.
func (s *NodeStatisticsForListNodePoolsOutput) SetRunningCount(v int32) *NodeStatisticsForListNodePoolsOutput {
	s.RunningCount = &v
	return s
}

// SetTotalCount sets the TotalCount field's value.
func (s *NodeStatisticsForListNodePoolsOutput) SetTotalCount(v int32) *NodeStatisticsForListNodePoolsOutput {
	s.TotalCount = &v
	return s
}

// SetUpdatingCount sets the UpdatingCount field's value.
func (s *NodeStatisticsForListNodePoolsOutput) SetUpdatingCount(v int32) *NodeStatisticsForListNodePoolsOutput {
	s.UpdatingCount = &v
	return s
}

type SecurityForListNodePoolsOutput struct {
	_ struct{} `type:"structure"`

	Login *LoginForListNodePoolsOutput `type:"structure"`

	SecurityGroupIds []*string `type:"list"`

	SecurityStrategies []*string `type:"list"`

	SecurityStrategyEnabled *bool `type:"boolean"`
}

// String returns the string representation
func (s SecurityForListNodePoolsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s SecurityForListNodePoolsOutput) GoString() string {
	return s.String()
}

// SetLogin sets the Login field's value.
func (s *SecurityForListNodePoolsOutput) SetLogin(v *LoginForListNodePoolsOutput) *SecurityForListNodePoolsOutput {
	s.Login = v
	return s
}

// SetSecurityGroupIds sets the SecurityGroupIds field's value.
func (s *SecurityForListNodePoolsOutput) SetSecurityGroupIds(v []*string) *SecurityForListNodePoolsOutput {
	s.SecurityGroupIds = v
	return s
}

// SetSecurityStrategies sets the SecurityStrategies field's value.
func (s *SecurityForListNodePoolsOutput) SetSecurityStrategies(v []*string) *SecurityForListNodePoolsOutput {
	s.SecurityStrategies = v
	return s
}

// SetSecurityStrategyEnabled sets the SecurityStrategyEnabled field's value.
func (s *SecurityForListNodePoolsOutput) SetSecurityStrategyEnabled(v bool) *SecurityForListNodePoolsOutput {
	s.SecurityStrategyEnabled = &v
	return s
}

type StatusForListNodePoolsInput struct {
	_ struct{} `type:"structure"`

	ConditionsType *string `type:"string" json:"Conditions.Type" enum:"EnumOfConditionsTypeForListNodePoolsInput"`

	Phase *string `type:"string" enum:"EnumOfPhaseForListNodePoolsInput"`
}

// String returns the string representation
func (s StatusForListNodePoolsInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s StatusForListNodePoolsInput) GoString() string {
	return s.String()
}

// SetConditionsType sets the ConditionsType field's value.
func (s *StatusForListNodePoolsInput) SetConditionsType(v string) *StatusForListNodePoolsInput {
	s.ConditionsType = &v
	return s
}

// SetPhase sets the Phase field's value.
func (s *StatusForListNodePoolsInput) SetPhase(v string) *StatusForListNodePoolsInput {
	s.Phase = &v
	return s
}

type StatusForListNodePoolsOutput struct {
	_ struct{} `type:"structure"`

	Conditions []*ConditionForListNodePoolsOutput `type:"list"`

	Phase *string `type:"string"`
}

// String returns the string representation
func (s StatusForListNodePoolsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s StatusForListNodePoolsOutput) GoString() string {
	return s.String()
}

// SetConditions sets the Conditions field's value.
func (s *StatusForListNodePoolsOutput) SetConditions(v []*ConditionForListNodePoolsOutput) *StatusForListNodePoolsOutput {
	s.Conditions = v
	return s
}

// SetPhase sets the Phase field's value.
func (s *StatusForListNodePoolsOutput) SetPhase(v string) *StatusForListNodePoolsOutput {
	s.Phase = &v
	return s
}

type SystemVolumeForListNodePoolsOutput struct {
	_ struct{} `type:"structure"`

	Size *int32 `type:"int32"`

	Type *string `type:"string" enum:"EnumOfTypeForListNodePoolsOutput"`
}

// String returns the string representation
func (s SystemVolumeForListNodePoolsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s SystemVolumeForListNodePoolsOutput) GoString() string {
	return s.String()
}

// SetSize sets the Size field's value.
func (s *SystemVolumeForListNodePoolsOutput) SetSize(v int32) *SystemVolumeForListNodePoolsOutput {
	s.Size = &v
	return s
}

// SetType sets the Type field's value.
func (s *SystemVolumeForListNodePoolsOutput) SetType(v string) *SystemVolumeForListNodePoolsOutput {
	s.Type = &v
	return s
}

type TagForListNodePoolsInput struct {
	_ struct{} `type:"structure"`

	Key *string `type:"string"`

	Value *string `type:"string"`
}

// String returns the string representation
func (s TagForListNodePoolsInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s TagForListNodePoolsInput) GoString() string {
	return s.String()
}

// SetKey sets the Key field's value.
func (s *TagForListNodePoolsInput) SetKey(v string) *TagForListNodePoolsInput {
	s.Key = &v
	return s
}

// SetValue sets the Value field's value.
func (s *TagForListNodePoolsInput) SetValue(v string) *TagForListNodePoolsInput {
	s.Value = &v
	return s
}

type TagForListNodePoolsOutput struct {
	_ struct{} `type:"structure"`

	Key *string `type:"string"`

	Value *string `type:"string"`
}

// String returns the string representation
func (s TagForListNodePoolsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s TagForListNodePoolsOutput) GoString() string {
	return s.String()
}

// SetKey sets the Key field's value.
func (s *TagForListNodePoolsOutput) SetKey(v string) *TagForListNodePoolsOutput {
	s.Key = &v
	return s
}

// SetValue sets the Value field's value.
func (s *TagForListNodePoolsOutput) SetValue(v string) *TagForListNodePoolsOutput {
	s.Value = &v
	return s
}

type TaintForListNodePoolsOutput struct {
	_ struct{} `type:"structure"`

	Effect *string `type:"string" enum:"EnumOfEffectForListNodePoolsOutput"`

	Key *string `type:"string"`

	Value *string `type:"string"`
}

// String returns the string representation
func (s TaintForListNodePoolsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s TaintForListNodePoolsOutput) GoString() string {
	return s.String()
}

// SetEffect sets the Effect field's value.
func (s *TaintForListNodePoolsOutput) SetEffect(v string) *TaintForListNodePoolsOutput {
	s.Effect = &v
	return s
}

// SetKey sets the Key field's value.
func (s *TaintForListNodePoolsOutput) SetKey(v string) *TaintForListNodePoolsOutput {
	s.Key = &v
	return s
}

// SetValue sets the Value field's value.
func (s *TaintForListNodePoolsOutput) SetValue(v string) *TaintForListNodePoolsOutput {
	s.Value = &v
	return s
}

const (
	// EnumOfConditionsTypeForListNodePoolsInputBalance is a EnumOfConditionsTypeForListNodePoolsInput enum value
	EnumOfConditionsTypeForListNodePoolsInputBalance = "Balance"

	// EnumOfConditionsTypeForListNodePoolsInputClusterNotRunning is a EnumOfConditionsTypeForListNodePoolsInput enum value
	EnumOfConditionsTypeForListNodePoolsInputClusterNotRunning = "ClusterNotRunning"

	// EnumOfConditionsTypeForListNodePoolsInputClusterVersionUpgrading is a EnumOfConditionsTypeForListNodePoolsInput enum value
	EnumOfConditionsTypeForListNodePoolsInputClusterVersionUpgrading = "ClusterVersionUpgrading"

	// EnumOfConditionsTypeForListNodePoolsInputLimitedByQuota is a EnumOfConditionsTypeForListNodePoolsInput enum value
	EnumOfConditionsTypeForListNodePoolsInputLimitedByQuota = "LimitedByQuota"

	// EnumOfConditionsTypeForListNodePoolsInputOk is a EnumOfConditionsTypeForListNodePoolsInput enum value
	EnumOfConditionsTypeForListNodePoolsInputOk = "Ok"

	// EnumOfConditionsTypeForListNodePoolsInputProgressing is a EnumOfConditionsTypeForListNodePoolsInput enum value
	EnumOfConditionsTypeForListNodePoolsInputProgressing = "Progressing"

	// EnumOfConditionsTypeForListNodePoolsInputResourceCleanupFailed is a EnumOfConditionsTypeForListNodePoolsInput enum value
	EnumOfConditionsTypeForListNodePoolsInputResourceCleanupFailed = "ResourceCleanupFailed"

	// EnumOfConditionsTypeForListNodePoolsInputStockOut is a EnumOfConditionsTypeForListNodePoolsInput enum value
	EnumOfConditionsTypeForListNodePoolsInputStockOut = "StockOut"

	// EnumOfConditionsTypeForListNodePoolsInputUnknown is a EnumOfConditionsTypeForListNodePoolsInput enum value
	EnumOfConditionsTypeForListNodePoolsInputUnknown = "Unknown"

	// EnumOfConditionsTypeForListNodePoolsInputVersionPartlyUpgraded is a EnumOfConditionsTypeForListNodePoolsInput enum value
	EnumOfConditionsTypeForListNodePoolsInputVersionPartlyUpgraded = "VersionPartlyUpgraded"
)

const (
	// EnumOfEffectForListNodePoolsOutputNoSchedule is a EnumOfEffectForListNodePoolsOutput enum value
	EnumOfEffectForListNodePoolsOutputNoSchedule = "NoSchedule"

	// EnumOfEffectForListNodePoolsOutputNoExecute is a EnumOfEffectForListNodePoolsOutput enum value
	EnumOfEffectForListNodePoolsOutputNoExecute = "NoExecute"

	// EnumOfEffectForListNodePoolsOutputPreferNoSchedule is a EnumOfEffectForListNodePoolsOutput enum value
	EnumOfEffectForListNodePoolsOutputPreferNoSchedule = "PreferNoSchedule"
)

const (
	// EnumOfPhaseForListNodePoolsInputCreating is a EnumOfPhaseForListNodePoolsInput enum value
	EnumOfPhaseForListNodePoolsInputCreating = "Creating"

	// EnumOfPhaseForListNodePoolsInputDeleting is a EnumOfPhaseForListNodePoolsInput enum value
	EnumOfPhaseForListNodePoolsInputDeleting = "Deleting"

	// EnumOfPhaseForListNodePoolsInputFailed is a EnumOfPhaseForListNodePoolsInput enum value
	EnumOfPhaseForListNodePoolsInputFailed = "Failed"

	// EnumOfPhaseForListNodePoolsInputRunning is a EnumOfPhaseForListNodePoolsInput enum value
	EnumOfPhaseForListNodePoolsInputRunning = "Running"

	// EnumOfPhaseForListNodePoolsInputScaling is a EnumOfPhaseForListNodePoolsInput enum value
	EnumOfPhaseForListNodePoolsInputScaling = "Scaling"

	// EnumOfPhaseForListNodePoolsInputUpdating is a EnumOfPhaseForListNodePoolsInput enum value
	EnumOfPhaseForListNodePoolsInputUpdating = "Updating"
)

const (
	// EnumOfSecurityStrategyListForListNodePoolsOutputHids is a EnumOfSecurityStrategyListForListNodePoolsOutput enum value
	EnumOfSecurityStrategyListForListNodePoolsOutputHids = "Hids"
)

const (
	// EnumOfSubnetPolicyForListNodePoolsOutputZoneBalance is a EnumOfSubnetPolicyForListNodePoolsOutput enum value
	EnumOfSubnetPolicyForListNodePoolsOutputZoneBalance = "ZoneBalance"

	// EnumOfSubnetPolicyForListNodePoolsOutputPriority is a EnumOfSubnetPolicyForListNodePoolsOutput enum value
	EnumOfSubnetPolicyForListNodePoolsOutputPriority = "Priority"
)

const (
	// EnumOfTypeForListNodePoolsOutputEssdPl0 is a EnumOfTypeForListNodePoolsOutput enum value
	EnumOfTypeForListNodePoolsOutputEssdPl0 = "ESSD_PL0"

	// EnumOfTypeForListNodePoolsOutputEssdFlexPl is a EnumOfTypeForListNodePoolsOutput enum value
	EnumOfTypeForListNodePoolsOutputEssdFlexPl = "ESSD_FlexPL"

	// EnumOfTypeForListNodePoolsOutputEssdPl1 is a EnumOfTypeForListNodePoolsOutput enum value
	EnumOfTypeForListNodePoolsOutputEssdPl1 = "ESSD_PL1"

	// EnumOfTypeForListNodePoolsOutputEssd is a EnumOfTypeForListNodePoolsOutput enum value
	EnumOfTypeForListNodePoolsOutputEssd = "ESSD"

	// EnumOfTypeForListNodePoolsOutputPtssd is a EnumOfTypeForListNodePoolsOutput enum value
	EnumOfTypeForListNodePoolsOutputPtssd = "PTSSD"
)

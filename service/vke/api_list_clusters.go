// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package vke

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opListClustersCommon = "ListClusters"

// ListClustersCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ListClustersCommon operation. The "output" return
// value will be populated with the ListClustersCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListClustersCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListClustersCommon Send returns without error.
//
// See ListClustersCommon for more information on using the ListClustersCommon
// API call, and error handling.
//
//    // Example sending a request using the ListClustersCommonRequest method.
//    req, resp := client.ListClustersCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VKE) ListClustersCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opListClustersCommon,
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

// ListClustersCommon API operation for VKE.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VKE's
// API operation ListClustersCommon for usage and error information.
func (c *VKE) ListClustersCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ListClustersCommonRequest(input)
	return out, req.Send()
}

// ListClustersCommonWithContext is the same as ListClustersCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ListClustersCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VKE) ListClustersCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ListClustersCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opListClusters = "ListClusters"

// ListClustersRequest generates a "volcengine/request.Request" representing the
// client's request for the ListClusters operation. The "output" return
// value will be populated with the ListClustersCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListClustersCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListClustersCommon Send returns without error.
//
// See ListClusters for more information on using the ListClusters
// API call, and error handling.
//
//    // Example sending a request using the ListClustersRequest method.
//    req, resp := client.ListClustersRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VKE) ListClustersRequest(input *ListClustersInput) (req *request.Request, output *ListClustersOutput) {
	op := &request.Operation{
		Name:       opListClusters,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListClustersInput{}
	}

	output = &ListClustersOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// ListClusters API operation for VKE.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VKE's
// API operation ListClusters for usage and error information.
func (c *VKE) ListClusters(input *ListClustersInput) (*ListClustersOutput, error) {
	req, out := c.ListClustersRequest(input)
	return out, req.Send()
}

// ListClustersWithContext is the same as ListClusters with the addition of
// the ability to pass a context and additional request options.
//
// See ListClusters for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VKE) ListClustersWithContext(ctx volcengine.Context, input *ListClustersInput, opts ...request.Option) (*ListClustersOutput, error) {
	req, out := c.ListClustersRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ApiServerEndpointsForListClustersOutput struct {
	_ struct{} `type:"structure"`

	PrivateIp *PrivateIpForListClustersOutput `type:"structure"`

	PublicIp *PublicIpForListClustersOutput `type:"structure"`
}

// String returns the string representation
func (s ApiServerEndpointsForListClustersOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ApiServerEndpointsForListClustersOutput) GoString() string {
	return s.String()
}

// SetPrivateIp sets the PrivateIp field's value.
func (s *ApiServerEndpointsForListClustersOutput) SetPrivateIp(v *PrivateIpForListClustersOutput) *ApiServerEndpointsForListClustersOutput {
	s.PrivateIp = v
	return s
}

// SetPublicIp sets the PublicIp field's value.
func (s *ApiServerEndpointsForListClustersOutput) SetPublicIp(v *PublicIpForListClustersOutput) *ApiServerEndpointsForListClustersOutput {
	s.PublicIp = v
	return s
}

type ApiServerPublicAccessConfigForListClustersOutput struct {
	_ struct{} `type:"structure"`

	AccessSourceIpsv4 []*string `type:"list"`

	PublicAccessNetworkConfig *PublicAccessNetworkConfigForListClustersOutput `type:"structure"`
}

// String returns the string representation
func (s ApiServerPublicAccessConfigForListClustersOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ApiServerPublicAccessConfigForListClustersOutput) GoString() string {
	return s.String()
}

// SetAccessSourceIpsv4 sets the AccessSourceIpsv4 field's value.
func (s *ApiServerPublicAccessConfigForListClustersOutput) SetAccessSourceIpsv4(v []*string) *ApiServerPublicAccessConfigForListClustersOutput {
	s.AccessSourceIpsv4 = v
	return s
}

// SetPublicAccessNetworkConfig sets the PublicAccessNetworkConfig field's value.
func (s *ApiServerPublicAccessConfigForListClustersOutput) SetPublicAccessNetworkConfig(v *PublicAccessNetworkConfigForListClustersOutput) *ApiServerPublicAccessConfigForListClustersOutput {
	s.PublicAccessNetworkConfig = v
	return s
}

type ClusterConfigForListClustersOutput struct {
	_ struct{} `type:"structure"`

	ApiServerEndpoints *ApiServerEndpointsForListClustersOutput `type:"structure"`

	ApiServerPublicAccessConfig *ApiServerPublicAccessConfigForListClustersOutput `type:"structure"`

	ApiServerPublicAccessEnabled *bool `type:"boolean"`

	ResourcePublicAccessDefaultEnabled *bool `type:"boolean"`

	SecurityGroupIds []*string `type:"list"`

	SubnetIds []*string `type:"list"`

	VpcId *string `type:"string"`
}

// String returns the string representation
func (s ClusterConfigForListClustersOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ClusterConfigForListClustersOutput) GoString() string {
	return s.String()
}

// SetApiServerEndpoints sets the ApiServerEndpoints field's value.
func (s *ClusterConfigForListClustersOutput) SetApiServerEndpoints(v *ApiServerEndpointsForListClustersOutput) *ClusterConfigForListClustersOutput {
	s.ApiServerEndpoints = v
	return s
}

// SetApiServerPublicAccessConfig sets the ApiServerPublicAccessConfig field's value.
func (s *ClusterConfigForListClustersOutput) SetApiServerPublicAccessConfig(v *ApiServerPublicAccessConfigForListClustersOutput) *ClusterConfigForListClustersOutput {
	s.ApiServerPublicAccessConfig = v
	return s
}

// SetApiServerPublicAccessEnabled sets the ApiServerPublicAccessEnabled field's value.
func (s *ClusterConfigForListClustersOutput) SetApiServerPublicAccessEnabled(v bool) *ClusterConfigForListClustersOutput {
	s.ApiServerPublicAccessEnabled = &v
	return s
}

// SetResourcePublicAccessDefaultEnabled sets the ResourcePublicAccessDefaultEnabled field's value.
func (s *ClusterConfigForListClustersOutput) SetResourcePublicAccessDefaultEnabled(v bool) *ClusterConfigForListClustersOutput {
	s.ResourcePublicAccessDefaultEnabled = &v
	return s
}

// SetSecurityGroupIds sets the SecurityGroupIds field's value.
func (s *ClusterConfigForListClustersOutput) SetSecurityGroupIds(v []*string) *ClusterConfigForListClustersOutput {
	s.SecurityGroupIds = v
	return s
}

// SetSubnetIds sets the SubnetIds field's value.
func (s *ClusterConfigForListClustersOutput) SetSubnetIds(v []*string) *ClusterConfigForListClustersOutput {
	s.SubnetIds = v
	return s
}

// SetVpcId sets the VpcId field's value.
func (s *ClusterConfigForListClustersOutput) SetVpcId(v string) *ClusterConfigForListClustersOutput {
	s.VpcId = &v
	return s
}

type ConditionForListClustersOutput struct {
	_ struct{} `type:"structure"`

	Type *string `type:"string"`
}

// String returns the string representation
func (s ConditionForListClustersOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ConditionForListClustersOutput) GoString() string {
	return s.String()
}

// SetType sets the Type field's value.
func (s *ConditionForListClustersOutput) SetType(v string) *ConditionForListClustersOutput {
	s.Type = &v
	return s
}

type FilterForListClustersInput struct {
	_ struct{} `type:"structure"`

	CreateClientToken *string `type:"string"`

	DeleteProtectionEnabled *bool `type:"boolean"`

	Ids []*string `type:"list"`

	Name *string `type:"string"`

	PodsConfigPodNetworkMode *string `type:"string" json:"PodsConfig.PodNetworkMode" enum:"EnumOfPodsConfigPodNetworkModeForListClustersInput"`

	Statuses []*StatusForListClustersInput `type:"list"`

	UpdateClientToken *string `type:"string"`
}

// String returns the string representation
func (s FilterForListClustersInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s FilterForListClustersInput) GoString() string {
	return s.String()
}

// SetCreateClientToken sets the CreateClientToken field's value.
func (s *FilterForListClustersInput) SetCreateClientToken(v string) *FilterForListClustersInput {
	s.CreateClientToken = &v
	return s
}

// SetDeleteProtectionEnabled sets the DeleteProtectionEnabled field's value.
func (s *FilterForListClustersInput) SetDeleteProtectionEnabled(v bool) *FilterForListClustersInput {
	s.DeleteProtectionEnabled = &v
	return s
}

// SetIds sets the Ids field's value.
func (s *FilterForListClustersInput) SetIds(v []*string) *FilterForListClustersInput {
	s.Ids = v
	return s
}

// SetName sets the Name field's value.
func (s *FilterForListClustersInput) SetName(v string) *FilterForListClustersInput {
	s.Name = &v
	return s
}

// SetPodsConfigPodNetworkMode sets the PodsConfigPodNetworkMode field's value.
func (s *FilterForListClustersInput) SetPodsConfigPodNetworkMode(v string) *FilterForListClustersInput {
	s.PodsConfigPodNetworkMode = &v
	return s
}

// SetStatuses sets the Statuses field's value.
func (s *FilterForListClustersInput) SetStatuses(v []*StatusForListClustersInput) *FilterForListClustersInput {
	s.Statuses = v
	return s
}

// SetUpdateClientToken sets the UpdateClientToken field's value.
func (s *FilterForListClustersInput) SetUpdateClientToken(v string) *FilterForListClustersInput {
	s.UpdateClientToken = &v
	return s
}

type FlannelConfigForListClustersOutput struct {
	_ struct{} `type:"structure"`

	MaxPodsPerNode *int32 `type:"int32"`

	PodCidrs []*string `type:"list"`
}

// String returns the string representation
func (s FlannelConfigForListClustersOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s FlannelConfigForListClustersOutput) GoString() string {
	return s.String()
}

// SetMaxPodsPerNode sets the MaxPodsPerNode field's value.
func (s *FlannelConfigForListClustersOutput) SetMaxPodsPerNode(v int32) *FlannelConfigForListClustersOutput {
	s.MaxPodsPerNode = &v
	return s
}

// SetPodCidrs sets the PodCidrs field's value.
func (s *FlannelConfigForListClustersOutput) SetPodCidrs(v []*string) *FlannelConfigForListClustersOutput {
	s.PodCidrs = v
	return s
}

type ItemForListClustersOutput struct {
	_ struct{} `type:"structure"`

	ChargeType *string `type:"string"`

	ClusterConfig *ClusterConfigForListClustersOutput `type:"structure"`

	CreateClientToken *string `type:"string"`

	CreateTime *string `type:"string"`

	DeleteProtectionEnabled *bool `type:"boolean"`

	Description *string `type:"string"`

	Id *string `type:"string"`

	KubernetesVersion *string `type:"string"`

	Name *string `type:"string"`

	NodeStatistics *NodeStatisticsForListClustersOutput `type:"structure"`

	PodsConfig *PodsConfigForListClustersOutput `type:"structure"`

	ServicesConfig *ServicesConfigForListClustersOutput `type:"structure"`

	Status *StatusForListClustersOutput `type:"structure"`

	UpdateClientToken *string `type:"string"`

	UpdateTime *string `type:"string"`
}

// String returns the string representation
func (s ItemForListClustersOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ItemForListClustersOutput) GoString() string {
	return s.String()
}

// SetChargeType sets the ChargeType field's value.
func (s *ItemForListClustersOutput) SetChargeType(v string) *ItemForListClustersOutput {
	s.ChargeType = &v
	return s
}

// SetClusterConfig sets the ClusterConfig field's value.
func (s *ItemForListClustersOutput) SetClusterConfig(v *ClusterConfigForListClustersOutput) *ItemForListClustersOutput {
	s.ClusterConfig = v
	return s
}

// SetCreateClientToken sets the CreateClientToken field's value.
func (s *ItemForListClustersOutput) SetCreateClientToken(v string) *ItemForListClustersOutput {
	s.CreateClientToken = &v
	return s
}

// SetCreateTime sets the CreateTime field's value.
func (s *ItemForListClustersOutput) SetCreateTime(v string) *ItemForListClustersOutput {
	s.CreateTime = &v
	return s
}

// SetDeleteProtectionEnabled sets the DeleteProtectionEnabled field's value.
func (s *ItemForListClustersOutput) SetDeleteProtectionEnabled(v bool) *ItemForListClustersOutput {
	s.DeleteProtectionEnabled = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *ItemForListClustersOutput) SetDescription(v string) *ItemForListClustersOutput {
	s.Description = &v
	return s
}

// SetId sets the Id field's value.
func (s *ItemForListClustersOutput) SetId(v string) *ItemForListClustersOutput {
	s.Id = &v
	return s
}

// SetKubernetesVersion sets the KubernetesVersion field's value.
func (s *ItemForListClustersOutput) SetKubernetesVersion(v string) *ItemForListClustersOutput {
	s.KubernetesVersion = &v
	return s
}

// SetName sets the Name field's value.
func (s *ItemForListClustersOutput) SetName(v string) *ItemForListClustersOutput {
	s.Name = &v
	return s
}

// SetNodeStatistics sets the NodeStatistics field's value.
func (s *ItemForListClustersOutput) SetNodeStatistics(v *NodeStatisticsForListClustersOutput) *ItemForListClustersOutput {
	s.NodeStatistics = v
	return s
}

// SetPodsConfig sets the PodsConfig field's value.
func (s *ItemForListClustersOutput) SetPodsConfig(v *PodsConfigForListClustersOutput) *ItemForListClustersOutput {
	s.PodsConfig = v
	return s
}

// SetServicesConfig sets the ServicesConfig field's value.
func (s *ItemForListClustersOutput) SetServicesConfig(v *ServicesConfigForListClustersOutput) *ItemForListClustersOutput {
	s.ServicesConfig = v
	return s
}

// SetStatus sets the Status field's value.
func (s *ItemForListClustersOutput) SetStatus(v *StatusForListClustersOutput) *ItemForListClustersOutput {
	s.Status = v
	return s
}

// SetUpdateClientToken sets the UpdateClientToken field's value.
func (s *ItemForListClustersOutput) SetUpdateClientToken(v string) *ItemForListClustersOutput {
	s.UpdateClientToken = &v
	return s
}

// SetUpdateTime sets the UpdateTime field's value.
func (s *ItemForListClustersOutput) SetUpdateTime(v string) *ItemForListClustersOutput {
	s.UpdateTime = &v
	return s
}

type ListClustersInput struct {
	_ struct{} `type:"structure"`

	Filter *FilterForListClustersInput `type:"structure"`

	PageNumber *int32 `type:"int32"`

	PageSize *int32 `type:"int32"`
}

// String returns the string representation
func (s ListClustersInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListClustersInput) GoString() string {
	return s.String()
}

// SetFilter sets the Filter field's value.
func (s *ListClustersInput) SetFilter(v *FilterForListClustersInput) *ListClustersInput {
	s.Filter = v
	return s
}

// SetPageNumber sets the PageNumber field's value.
func (s *ListClustersInput) SetPageNumber(v int32) *ListClustersInput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *ListClustersInput) SetPageSize(v int32) *ListClustersInput {
	s.PageSize = &v
	return s
}

type ListClustersOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	Items []*ItemForListClustersOutput `type:"list"`

	PageNumber *int32 `type:"int32"`

	PageSize *int32 `type:"int32"`

	TotalCount *int32 `type:"int32"`
}

// String returns the string representation
func (s ListClustersOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListClustersOutput) GoString() string {
	return s.String()
}

// SetItems sets the Items field's value.
func (s *ListClustersOutput) SetItems(v []*ItemForListClustersOutput) *ListClustersOutput {
	s.Items = v
	return s
}

// SetPageNumber sets the PageNumber field's value.
func (s *ListClustersOutput) SetPageNumber(v int32) *ListClustersOutput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *ListClustersOutput) SetPageSize(v int32) *ListClustersOutput {
	s.PageSize = &v
	return s
}

// SetTotalCount sets the TotalCount field's value.
func (s *ListClustersOutput) SetTotalCount(v int32) *ListClustersOutput {
	s.TotalCount = &v
	return s
}

type NodeStatisticsForListClustersOutput struct {
	_ struct{} `type:"structure"`

	CreatingCount *int32 `type:"int32"`

	DeletingCount *int32 `type:"int32"`

	FailedCount *int32 `type:"int32"`

	RunningCount *int32 `type:"int32"`

	TotalCount *int32 `type:"int32"`

	UpdatingCount *int32 `type:"int32"`
}

// String returns the string representation
func (s NodeStatisticsForListClustersOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s NodeStatisticsForListClustersOutput) GoString() string {
	return s.String()
}

// SetCreatingCount sets the CreatingCount field's value.
func (s *NodeStatisticsForListClustersOutput) SetCreatingCount(v int32) *NodeStatisticsForListClustersOutput {
	s.CreatingCount = &v
	return s
}

// SetDeletingCount sets the DeletingCount field's value.
func (s *NodeStatisticsForListClustersOutput) SetDeletingCount(v int32) *NodeStatisticsForListClustersOutput {
	s.DeletingCount = &v
	return s
}

// SetFailedCount sets the FailedCount field's value.
func (s *NodeStatisticsForListClustersOutput) SetFailedCount(v int32) *NodeStatisticsForListClustersOutput {
	s.FailedCount = &v
	return s
}

// SetRunningCount sets the RunningCount field's value.
func (s *NodeStatisticsForListClustersOutput) SetRunningCount(v int32) *NodeStatisticsForListClustersOutput {
	s.RunningCount = &v
	return s
}

// SetTotalCount sets the TotalCount field's value.
func (s *NodeStatisticsForListClustersOutput) SetTotalCount(v int32) *NodeStatisticsForListClustersOutput {
	s.TotalCount = &v
	return s
}

// SetUpdatingCount sets the UpdatingCount field's value.
func (s *NodeStatisticsForListClustersOutput) SetUpdatingCount(v int32) *NodeStatisticsForListClustersOutput {
	s.UpdatingCount = &v
	return s
}

type PodsConfigForListClustersOutput struct {
	_ struct{} `type:"structure"`

	FlannelConfig *FlannelConfigForListClustersOutput `type:"structure"`

	PodNetworkMode *string `type:"string"`

	VpcCniConfig *VpcCniConfigForListClustersOutput `type:"structure"`
}

// String returns the string representation
func (s PodsConfigForListClustersOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s PodsConfigForListClustersOutput) GoString() string {
	return s.String()
}

// SetFlannelConfig sets the FlannelConfig field's value.
func (s *PodsConfigForListClustersOutput) SetFlannelConfig(v *FlannelConfigForListClustersOutput) *PodsConfigForListClustersOutput {
	s.FlannelConfig = v
	return s
}

// SetPodNetworkMode sets the PodNetworkMode field's value.
func (s *PodsConfigForListClustersOutput) SetPodNetworkMode(v string) *PodsConfigForListClustersOutput {
	s.PodNetworkMode = &v
	return s
}

// SetVpcCniConfig sets the VpcCniConfig field's value.
func (s *PodsConfigForListClustersOutput) SetVpcCniConfig(v *VpcCniConfigForListClustersOutput) *PodsConfigForListClustersOutput {
	s.VpcCniConfig = v
	return s
}

type PrivateIpForListClustersOutput struct {
	_ struct{} `type:"structure"`

	Ipv4 *string `type:"string"`
}

// String returns the string representation
func (s PrivateIpForListClustersOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s PrivateIpForListClustersOutput) GoString() string {
	return s.String()
}

// SetIpv4 sets the Ipv4 field's value.
func (s *PrivateIpForListClustersOutput) SetIpv4(v string) *PrivateIpForListClustersOutput {
	s.Ipv4 = &v
	return s
}

type PublicAccessNetworkConfigForListClustersOutput struct {
	_ struct{} `type:"structure"`

	Bandwidth *int32 `type:"int32"`

	BillingType *int32 `type:"int32"`

	Isp *string `type:"string"`
}

// String returns the string representation
func (s PublicAccessNetworkConfigForListClustersOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s PublicAccessNetworkConfigForListClustersOutput) GoString() string {
	return s.String()
}

// SetBandwidth sets the Bandwidth field's value.
func (s *PublicAccessNetworkConfigForListClustersOutput) SetBandwidth(v int32) *PublicAccessNetworkConfigForListClustersOutput {
	s.Bandwidth = &v
	return s
}

// SetBillingType sets the BillingType field's value.
func (s *PublicAccessNetworkConfigForListClustersOutput) SetBillingType(v int32) *PublicAccessNetworkConfigForListClustersOutput {
	s.BillingType = &v
	return s
}

// SetIsp sets the Isp field's value.
func (s *PublicAccessNetworkConfigForListClustersOutput) SetIsp(v string) *PublicAccessNetworkConfigForListClustersOutput {
	s.Isp = &v
	return s
}

type PublicIpForListClustersOutput struct {
	_ struct{} `type:"structure"`

	Ipv4 *string `type:"string"`
}

// String returns the string representation
func (s PublicIpForListClustersOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s PublicIpForListClustersOutput) GoString() string {
	return s.String()
}

// SetIpv4 sets the Ipv4 field's value.
func (s *PublicIpForListClustersOutput) SetIpv4(v string) *PublicIpForListClustersOutput {
	s.Ipv4 = &v
	return s
}

type ServicesConfigForListClustersOutput struct {
	_ struct{} `type:"structure"`

	ServiceCidrsv4 []*string `type:"list"`
}

// String returns the string representation
func (s ServicesConfigForListClustersOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ServicesConfigForListClustersOutput) GoString() string {
	return s.String()
}

// SetServiceCidrsv4 sets the ServiceCidrsv4 field's value.
func (s *ServicesConfigForListClustersOutput) SetServiceCidrsv4(v []*string) *ServicesConfigForListClustersOutput {
	s.ServiceCidrsv4 = v
	return s
}

type StatusForListClustersInput struct {
	_ struct{} `type:"structure"`

	ConditionsType *string `type:"string" json:"Conditions.Type" enum:"EnumOfConditionsTypeForListClustersInput"`

	Phase *string `type:"string" enum:"EnumOfPhaseForListClustersInput"`
}

// String returns the string representation
func (s StatusForListClustersInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s StatusForListClustersInput) GoString() string {
	return s.String()
}

// SetConditionsType sets the ConditionsType field's value.
func (s *StatusForListClustersInput) SetConditionsType(v string) *StatusForListClustersInput {
	s.ConditionsType = &v
	return s
}

// SetPhase sets the Phase field's value.
func (s *StatusForListClustersInput) SetPhase(v string) *StatusForListClustersInput {
	s.Phase = &v
	return s
}

type StatusForListClustersOutput struct {
	_ struct{} `type:"structure"`

	Conditions []*ConditionForListClustersOutput `type:"list"`

	Phase *string `type:"string"`
}

// String returns the string representation
func (s StatusForListClustersOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s StatusForListClustersOutput) GoString() string {
	return s.String()
}

// SetConditions sets the Conditions field's value.
func (s *StatusForListClustersOutput) SetConditions(v []*ConditionForListClustersOutput) *StatusForListClustersOutput {
	s.Conditions = v
	return s
}

// SetPhase sets the Phase field's value.
func (s *StatusForListClustersOutput) SetPhase(v string) *StatusForListClustersOutput {
	s.Phase = &v
	return s
}

type VpcCniConfigForListClustersOutput struct {
	_ struct{} `type:"structure"`

	SubnetIds []*string `type:"list"`
}

// String returns the string representation
func (s VpcCniConfigForListClustersOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s VpcCniConfigForListClustersOutput) GoString() string {
	return s.String()
}

// SetSubnetIds sets the SubnetIds field's value.
func (s *VpcCniConfigForListClustersOutput) SetSubnetIds(v []*string) *VpcCniConfigForListClustersOutput {
	s.SubnetIds = v
	return s
}

const (
	// EnumOfConditionsTypeForListClustersInputBalance is a EnumOfConditionsTypeForListClustersInput enum value
	EnumOfConditionsTypeForListClustersInputBalance = "Balance"

	// EnumOfConditionsTypeForListClustersInputClusterVersionUpgrading is a EnumOfConditionsTypeForListClustersInput enum value
	EnumOfConditionsTypeForListClustersInputClusterVersionUpgrading = "ClusterVersionUpgrading"

	// EnumOfConditionsTypeForListClustersInputCreateError is a EnumOfConditionsTypeForListClustersInput enum value
	EnumOfConditionsTypeForListClustersInputCreateError = "CreateError"

	// EnumOfConditionsTypeForListClustersInputDegraded is a EnumOfConditionsTypeForListClustersInput enum value
	EnumOfConditionsTypeForListClustersInputDegraded = "Degraded"

	// EnumOfConditionsTypeForListClustersInputDisconnected is a EnumOfConditionsTypeForListClustersInput enum value
	EnumOfConditionsTypeForListClustersInputDisconnected = "Disconnected"

	// EnumOfConditionsTypeForListClustersInputInvalidCertificate is a EnumOfConditionsTypeForListClustersInput enum value
	EnumOfConditionsTypeForListClustersInputInvalidCertificate = "InvalidCertificate"

	// EnumOfConditionsTypeForListClustersInputLimitedByQuota is a EnumOfConditionsTypeForListClustersInput enum value
	EnumOfConditionsTypeForListClustersInputLimitedByQuota = "LimitedByQuota"

	// EnumOfConditionsTypeForListClustersInputOk is a EnumOfConditionsTypeForListClustersInput enum value
	EnumOfConditionsTypeForListClustersInputOk = "Ok"

	// EnumOfConditionsTypeForListClustersInputProgressing is a EnumOfConditionsTypeForListClustersInput enum value
	EnumOfConditionsTypeForListClustersInputProgressing = "Progressing"

	// EnumOfConditionsTypeForListClustersInputResourceCleanupFailed is a EnumOfConditionsTypeForListClustersInput enum value
	EnumOfConditionsTypeForListClustersInputResourceCleanupFailed = "ResourceCleanupFailed"

	// EnumOfConditionsTypeForListClustersInputSecurity is a EnumOfConditionsTypeForListClustersInput enum value
	EnumOfConditionsTypeForListClustersInputSecurity = "Security"

	// EnumOfConditionsTypeForListClustersInputSetByProvider is a EnumOfConditionsTypeForListClustersInput enum value
	EnumOfConditionsTypeForListClustersInputSetByProvider = "SetByProvider"

	// EnumOfConditionsTypeForListClustersInputSetByUser is a EnumOfConditionsTypeForListClustersInput enum value
	EnumOfConditionsTypeForListClustersInputSetByUser = "SetByUser"

	// EnumOfConditionsTypeForListClustersInputStockOut is a EnumOfConditionsTypeForListClustersInput enum value
	EnumOfConditionsTypeForListClustersInputStockOut = "StockOut"

	// EnumOfConditionsTypeForListClustersInputUnknown is a EnumOfConditionsTypeForListClustersInput enum value
	EnumOfConditionsTypeForListClustersInputUnknown = "Unknown"

	// EnumOfConditionsTypeForListClustersInputWaiting is a EnumOfConditionsTypeForListClustersInput enum value
	EnumOfConditionsTypeForListClustersInputWaiting = "Waiting"
)

const (
	// EnumOfPhaseForListClustersInputCreating is a EnumOfPhaseForListClustersInput enum value
	EnumOfPhaseForListClustersInputCreating = "Creating"

	// EnumOfPhaseForListClustersInputDeleting is a EnumOfPhaseForListClustersInput enum value
	EnumOfPhaseForListClustersInputDeleting = "Deleting"

	// EnumOfPhaseForListClustersInputFailed is a EnumOfPhaseForListClustersInput enum value
	EnumOfPhaseForListClustersInputFailed = "Failed"

	// EnumOfPhaseForListClustersInputRunning is a EnumOfPhaseForListClustersInput enum value
	EnumOfPhaseForListClustersInputRunning = "Running"

	// EnumOfPhaseForListClustersInputStarting is a EnumOfPhaseForListClustersInput enum value
	EnumOfPhaseForListClustersInputStarting = "Starting"

	// EnumOfPhaseForListClustersInputStopped is a EnumOfPhaseForListClustersInput enum value
	EnumOfPhaseForListClustersInputStopped = "Stopped"

	// EnumOfPhaseForListClustersInputUpdating is a EnumOfPhaseForListClustersInput enum value
	EnumOfPhaseForListClustersInputUpdating = "Updating"
)

const (
	// EnumOfPodsConfigPodNetworkModeForListClustersInputCalicoBgp is a EnumOfPodsConfigPodNetworkModeForListClustersInput enum value
	EnumOfPodsConfigPodNetworkModeForListClustersInputCalicoBgp = "CalicoBgp"

	// EnumOfPodsConfigPodNetworkModeForListClustersInputCalicoVxlan is a EnumOfPodsConfigPodNetworkModeForListClustersInput enum value
	EnumOfPodsConfigPodNetworkModeForListClustersInputCalicoVxlan = "CalicoVxlan"

	// EnumOfPodsConfigPodNetworkModeForListClustersInputCarma is a EnumOfPodsConfigPodNetworkModeForListClustersInput enum value
	EnumOfPodsConfigPodNetworkModeForListClustersInputCarma = "Carma"

	// EnumOfPodsConfigPodNetworkModeForListClustersInputCilium is a EnumOfPodsConfigPodNetworkModeForListClustersInput enum value
	EnumOfPodsConfigPodNetworkModeForListClustersInputCilium = "Cilium"

	// EnumOfPodsConfigPodNetworkModeForListClustersInputDefault is a EnumOfPodsConfigPodNetworkModeForListClustersInput enum value
	EnumOfPodsConfigPodNetworkModeForListClustersInputDefault = "Default"

	// EnumOfPodsConfigPodNetworkModeForListClustersInputFlannel is a EnumOfPodsConfigPodNetworkModeForListClustersInput enum value
	EnumOfPodsConfigPodNetworkModeForListClustersInputFlannel = "Flannel"

	// EnumOfPodsConfigPodNetworkModeForListClustersInputKubeOvn is a EnumOfPodsConfigPodNetworkModeForListClustersInput enum value
	EnumOfPodsConfigPodNetworkModeForListClustersInputKubeOvn = "KubeOvn"

	// EnumOfPodsConfigPodNetworkModeForListClustersInputVpcCniDedicated is a EnumOfPodsConfigPodNetworkModeForListClustersInput enum value
	EnumOfPodsConfigPodNetworkModeForListClustersInputVpcCniDedicated = "VpcCniDedicated"

	// EnumOfPodsConfigPodNetworkModeForListClustersInputVpcCniShared is a EnumOfPodsConfigPodNetworkModeForListClustersInput enum value
	EnumOfPodsConfigPodNetworkModeForListClustersInputVpcCniShared = "VpcCniShared"
)

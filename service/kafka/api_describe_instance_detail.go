// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package kafka

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribeInstanceDetailCommon = "DescribeInstanceDetail"

// DescribeInstanceDetailCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeInstanceDetailCommon operation. The "output" return
// value will be populated with the DescribeInstanceDetailCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeInstanceDetailCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeInstanceDetailCommon Send returns without error.
//
// See DescribeInstanceDetailCommon for more information on using the DescribeInstanceDetailCommon
// API call, and error handling.
//
//	// Example sending a request using the DescribeInstanceDetailCommonRequest method.
//	req, resp := client.DescribeInstanceDetailCommonRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *KAFKA) DescribeInstanceDetailCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeInstanceDetailCommon,
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

// DescribeInstanceDetailCommon API operation for KAFKA.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for KAFKA's
// API operation DescribeInstanceDetailCommon for usage and error information.
func (c *KAFKA) DescribeInstanceDetailCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeInstanceDetailCommonRequest(input)
	return out, req.Send()
}

// DescribeInstanceDetailCommonWithContext is the same as DescribeInstanceDetailCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeInstanceDetailCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *KAFKA) DescribeInstanceDetailCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeInstanceDetailCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeInstanceDetail = "DescribeInstanceDetail"

// DescribeInstanceDetailRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeInstanceDetail operation. The "output" return
// value will be populated with the DescribeInstanceDetailCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeInstanceDetailCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeInstanceDetailCommon Send returns without error.
//
// See DescribeInstanceDetail for more information on using the DescribeInstanceDetail
// API call, and error handling.
//
//	// Example sending a request using the DescribeInstanceDetailRequest method.
//	req, resp := client.DescribeInstanceDetailRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *KAFKA) DescribeInstanceDetailRequest(input *DescribeInstanceDetailInput) (req *request.Request, output *DescribeInstanceDetailOutput) {
	op := &request.Operation{
		Name:       opDescribeInstanceDetail,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeInstanceDetailInput{}
	}

	output = &DescribeInstanceDetailOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// DescribeInstanceDetail API operation for KAFKA.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for KAFKA's
// API operation DescribeInstanceDetail for usage and error information.
func (c *KAFKA) DescribeInstanceDetail(input *DescribeInstanceDetailInput) (*DescribeInstanceDetailOutput, error) {
	req, out := c.DescribeInstanceDetailRequest(input)
	return out, req.Send()
}

// DescribeInstanceDetailWithContext is the same as DescribeInstanceDetail with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeInstanceDetail for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *KAFKA) DescribeInstanceDetailWithContext(ctx volcengine.Context, input *DescribeInstanceDetailInput, opts ...request.Option) (*DescribeInstanceDetailOutput, error) {
	req, out := c.DescribeInstanceDetailRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type BasicInstanceInfoForDescribeInstanceDetailOutput struct {
	_ struct{} `type:"structure"`

	AccountId *string `type:"string"`

	ChargeDetail *ChargeDetailForDescribeInstanceDetailOutput `type:"structure"`

	ComputeSpec *string `type:"string"`

	CreateTime *string `type:"string"`

	EipId *string `type:"string"`

	InstanceDescription *string `type:"string"`

	InstanceId *string `type:"string"`

	InstanceName *string `type:"string"`

	InstanceStatus *string `type:"string"`

	PrivateDomainOnPublic *bool `type:"boolean"`

	ProjectName *string `type:"string"`

	RegionId *string `type:"string"`

	StorageSpace *int32 `type:"int32"`

	StorageType *string `type:"string"`

	SubnetId *string `type:"string"`

	Tags map[string]*string `type:"map"`

	UsableGroupNumber *int32 `type:"int32"`

	UsablePartitionNumber *int32 `type:"int32"`

	UsedGroupNumber *int32 `type:"int32"`

	UsedPartitionNumber *int32 `type:"int32"`

	UsedStorageSpace *int32 `type:"int32"`

	UsedTopicNumber *int32 `type:"int32"`

	Version *string `type:"string"`

	VpcId *string `type:"string"`

	ZoneId *string `type:"string"`
}

// String returns the string representation
func (s BasicInstanceInfoForDescribeInstanceDetailOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s BasicInstanceInfoForDescribeInstanceDetailOutput) GoString() string {
	return s.String()
}

// SetAccountId sets the AccountId field's value.
func (s *BasicInstanceInfoForDescribeInstanceDetailOutput) SetAccountId(v string) *BasicInstanceInfoForDescribeInstanceDetailOutput {
	s.AccountId = &v
	return s
}

// SetChargeDetail sets the ChargeDetail field's value.
func (s *BasicInstanceInfoForDescribeInstanceDetailOutput) SetChargeDetail(v *ChargeDetailForDescribeInstanceDetailOutput) *BasicInstanceInfoForDescribeInstanceDetailOutput {
	s.ChargeDetail = v
	return s
}

// SetComputeSpec sets the ComputeSpec field's value.
func (s *BasicInstanceInfoForDescribeInstanceDetailOutput) SetComputeSpec(v string) *BasicInstanceInfoForDescribeInstanceDetailOutput {
	s.ComputeSpec = &v
	return s
}

// SetCreateTime sets the CreateTime field's value.
func (s *BasicInstanceInfoForDescribeInstanceDetailOutput) SetCreateTime(v string) *BasicInstanceInfoForDescribeInstanceDetailOutput {
	s.CreateTime = &v
	return s
}

// SetEipId sets the EipId field's value.
func (s *BasicInstanceInfoForDescribeInstanceDetailOutput) SetEipId(v string) *BasicInstanceInfoForDescribeInstanceDetailOutput {
	s.EipId = &v
	return s
}

// SetInstanceDescription sets the InstanceDescription field's value.
func (s *BasicInstanceInfoForDescribeInstanceDetailOutput) SetInstanceDescription(v string) *BasicInstanceInfoForDescribeInstanceDetailOutput {
	s.InstanceDescription = &v
	return s
}

// SetInstanceId sets the InstanceId field's value.
func (s *BasicInstanceInfoForDescribeInstanceDetailOutput) SetInstanceId(v string) *BasicInstanceInfoForDescribeInstanceDetailOutput {
	s.InstanceId = &v
	return s
}

// SetInstanceName sets the InstanceName field's value.
func (s *BasicInstanceInfoForDescribeInstanceDetailOutput) SetInstanceName(v string) *BasicInstanceInfoForDescribeInstanceDetailOutput {
	s.InstanceName = &v
	return s
}

// SetInstanceStatus sets the InstanceStatus field's value.
func (s *BasicInstanceInfoForDescribeInstanceDetailOutput) SetInstanceStatus(v string) *BasicInstanceInfoForDescribeInstanceDetailOutput {
	s.InstanceStatus = &v
	return s
}

// SetPrivateDomainOnPublic sets the PrivateDomainOnPublic field's value.
func (s *BasicInstanceInfoForDescribeInstanceDetailOutput) SetPrivateDomainOnPublic(v bool) *BasicInstanceInfoForDescribeInstanceDetailOutput {
	s.PrivateDomainOnPublic = &v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *BasicInstanceInfoForDescribeInstanceDetailOutput) SetProjectName(v string) *BasicInstanceInfoForDescribeInstanceDetailOutput {
	s.ProjectName = &v
	return s
}

// SetRegionId sets the RegionId field's value.
func (s *BasicInstanceInfoForDescribeInstanceDetailOutput) SetRegionId(v string) *BasicInstanceInfoForDescribeInstanceDetailOutput {
	s.RegionId = &v
	return s
}

// SetStorageSpace sets the StorageSpace field's value.
func (s *BasicInstanceInfoForDescribeInstanceDetailOutput) SetStorageSpace(v int32) *BasicInstanceInfoForDescribeInstanceDetailOutput {
	s.StorageSpace = &v
	return s
}

// SetStorageType sets the StorageType field's value.
func (s *BasicInstanceInfoForDescribeInstanceDetailOutput) SetStorageType(v string) *BasicInstanceInfoForDescribeInstanceDetailOutput {
	s.StorageType = &v
	return s
}

// SetSubnetId sets the SubnetId field's value.
func (s *BasicInstanceInfoForDescribeInstanceDetailOutput) SetSubnetId(v string) *BasicInstanceInfoForDescribeInstanceDetailOutput {
	s.SubnetId = &v
	return s
}

// SetTags sets the Tags field's value.
func (s *BasicInstanceInfoForDescribeInstanceDetailOutput) SetTags(v map[string]*string) *BasicInstanceInfoForDescribeInstanceDetailOutput {
	s.Tags = v
	return s
}

// SetUsableGroupNumber sets the UsableGroupNumber field's value.
func (s *BasicInstanceInfoForDescribeInstanceDetailOutput) SetUsableGroupNumber(v int32) *BasicInstanceInfoForDescribeInstanceDetailOutput {
	s.UsableGroupNumber = &v
	return s
}

// SetUsablePartitionNumber sets the UsablePartitionNumber field's value.
func (s *BasicInstanceInfoForDescribeInstanceDetailOutput) SetUsablePartitionNumber(v int32) *BasicInstanceInfoForDescribeInstanceDetailOutput {
	s.UsablePartitionNumber = &v
	return s
}

// SetUsedGroupNumber sets the UsedGroupNumber field's value.
func (s *BasicInstanceInfoForDescribeInstanceDetailOutput) SetUsedGroupNumber(v int32) *BasicInstanceInfoForDescribeInstanceDetailOutput {
	s.UsedGroupNumber = &v
	return s
}

// SetUsedPartitionNumber sets the UsedPartitionNumber field's value.
func (s *BasicInstanceInfoForDescribeInstanceDetailOutput) SetUsedPartitionNumber(v int32) *BasicInstanceInfoForDescribeInstanceDetailOutput {
	s.UsedPartitionNumber = &v
	return s
}

// SetUsedStorageSpace sets the UsedStorageSpace field's value.
func (s *BasicInstanceInfoForDescribeInstanceDetailOutput) SetUsedStorageSpace(v int32) *BasicInstanceInfoForDescribeInstanceDetailOutput {
	s.UsedStorageSpace = &v
	return s
}

// SetUsedTopicNumber sets the UsedTopicNumber field's value.
func (s *BasicInstanceInfoForDescribeInstanceDetailOutput) SetUsedTopicNumber(v int32) *BasicInstanceInfoForDescribeInstanceDetailOutput {
	s.UsedTopicNumber = &v
	return s
}

// SetVersion sets the Version field's value.
func (s *BasicInstanceInfoForDescribeInstanceDetailOutput) SetVersion(v string) *BasicInstanceInfoForDescribeInstanceDetailOutput {
	s.Version = &v
	return s
}

// SetVpcId sets the VpcId field's value.
func (s *BasicInstanceInfoForDescribeInstanceDetailOutput) SetVpcId(v string) *BasicInstanceInfoForDescribeInstanceDetailOutput {
	s.VpcId = &v
	return s
}

// SetZoneId sets the ZoneId field's value.
func (s *BasicInstanceInfoForDescribeInstanceDetailOutput) SetZoneId(v string) *BasicInstanceInfoForDescribeInstanceDetailOutput {
	s.ZoneId = &v
	return s
}

type ChargeDetailForDescribeInstanceDetailOutput struct {
	_ struct{} `type:"structure"`

	AutoRenew *bool `type:"boolean"`

	ChargeExpireTime *string `type:"string"`

	ChargeStartTime *string `type:"string"`

	ChargeStatus *string `type:"string"`

	ChargeType *string `type:"string"`

	OverdueReclaimTime *string `type:"string"`

	OverdueTime *string `type:"string"`

	PeriodUnit *string `type:"string"`
}

// String returns the string representation
func (s ChargeDetailForDescribeInstanceDetailOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ChargeDetailForDescribeInstanceDetailOutput) GoString() string {
	return s.String()
}

// SetAutoRenew sets the AutoRenew field's value.
func (s *ChargeDetailForDescribeInstanceDetailOutput) SetAutoRenew(v bool) *ChargeDetailForDescribeInstanceDetailOutput {
	s.AutoRenew = &v
	return s
}

// SetChargeExpireTime sets the ChargeExpireTime field's value.
func (s *ChargeDetailForDescribeInstanceDetailOutput) SetChargeExpireTime(v string) *ChargeDetailForDescribeInstanceDetailOutput {
	s.ChargeExpireTime = &v
	return s
}

// SetChargeStartTime sets the ChargeStartTime field's value.
func (s *ChargeDetailForDescribeInstanceDetailOutput) SetChargeStartTime(v string) *ChargeDetailForDescribeInstanceDetailOutput {
	s.ChargeStartTime = &v
	return s
}

// SetChargeStatus sets the ChargeStatus field's value.
func (s *ChargeDetailForDescribeInstanceDetailOutput) SetChargeStatus(v string) *ChargeDetailForDescribeInstanceDetailOutput {
	s.ChargeStatus = &v
	return s
}

// SetChargeType sets the ChargeType field's value.
func (s *ChargeDetailForDescribeInstanceDetailOutput) SetChargeType(v string) *ChargeDetailForDescribeInstanceDetailOutput {
	s.ChargeType = &v
	return s
}

// SetOverdueReclaimTime sets the OverdueReclaimTime field's value.
func (s *ChargeDetailForDescribeInstanceDetailOutput) SetOverdueReclaimTime(v string) *ChargeDetailForDescribeInstanceDetailOutput {
	s.OverdueReclaimTime = &v
	return s
}

// SetOverdueTime sets the OverdueTime field's value.
func (s *ChargeDetailForDescribeInstanceDetailOutput) SetOverdueTime(v string) *ChargeDetailForDescribeInstanceDetailOutput {
	s.OverdueTime = &v
	return s
}

// SetPeriodUnit sets the PeriodUnit field's value.
func (s *ChargeDetailForDescribeInstanceDetailOutput) SetPeriodUnit(v string) *ChargeDetailForDescribeInstanceDetailOutput {
	s.PeriodUnit = &v
	return s
}

type ConnectionInfoForDescribeInstanceDetailOutput struct {
	_ struct{} `type:"structure"`

	EndpointType *string `type:"string"`

	InternalEndpoint *string `type:"string"`

	NetworkType *string `type:"string"`

	PublicEndpoint *string `type:"string"`
}

// String returns the string representation
func (s ConnectionInfoForDescribeInstanceDetailOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ConnectionInfoForDescribeInstanceDetailOutput) GoString() string {
	return s.String()
}

// SetEndpointType sets the EndpointType field's value.
func (s *ConnectionInfoForDescribeInstanceDetailOutput) SetEndpointType(v string) *ConnectionInfoForDescribeInstanceDetailOutput {
	s.EndpointType = &v
	return s
}

// SetInternalEndpoint sets the InternalEndpoint field's value.
func (s *ConnectionInfoForDescribeInstanceDetailOutput) SetInternalEndpoint(v string) *ConnectionInfoForDescribeInstanceDetailOutput {
	s.InternalEndpoint = &v
	return s
}

// SetNetworkType sets the NetworkType field's value.
func (s *ConnectionInfoForDescribeInstanceDetailOutput) SetNetworkType(v string) *ConnectionInfoForDescribeInstanceDetailOutput {
	s.NetworkType = &v
	return s
}

// SetPublicEndpoint sets the PublicEndpoint field's value.
func (s *ConnectionInfoForDescribeInstanceDetailOutput) SetPublicEndpoint(v string) *ConnectionInfoForDescribeInstanceDetailOutput {
	s.PublicEndpoint = &v
	return s
}

type DescribeInstanceDetailInput struct {
	_ struct{} `type:"structure"`

	InstanceId *string `type:"string"`
}

// String returns the string representation
func (s DescribeInstanceDetailInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeInstanceDetailInput) GoString() string {
	return s.String()
}

// SetInstanceId sets the InstanceId field's value.
func (s *DescribeInstanceDetailInput) SetInstanceId(v string) *DescribeInstanceDetailInput {
	s.InstanceId = &v
	return s
}

type DescribeInstanceDetailOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	BasicInstanceInfo *BasicInstanceInfoForDescribeInstanceDetailOutput `type:"structure"`

	ChargeDetail *ChargeDetailForDescribeInstanceDetailOutput `type:"structure"`

	ConnectionInfo []*ConnectionInfoForDescribeInstanceDetailOutput `type:"list"`

	Parameters *string `type:"string"`
}

// String returns the string representation
func (s DescribeInstanceDetailOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeInstanceDetailOutput) GoString() string {
	return s.String()
}

// SetBasicInstanceInfo sets the BasicInstanceInfo field's value.
func (s *DescribeInstanceDetailOutput) SetBasicInstanceInfo(v *BasicInstanceInfoForDescribeInstanceDetailOutput) *DescribeInstanceDetailOutput {
	s.BasicInstanceInfo = v
	return s
}

// SetChargeDetail sets the ChargeDetail field's value.
func (s *DescribeInstanceDetailOutput) SetChargeDetail(v *ChargeDetailForDescribeInstanceDetailOutput) *DescribeInstanceDetailOutput {
	s.ChargeDetail = v
	return s
}

// SetConnectionInfo sets the ConnectionInfo field's value.
func (s *DescribeInstanceDetailOutput) SetConnectionInfo(v []*ConnectionInfoForDescribeInstanceDetailOutput) *DescribeInstanceDetailOutput {
	s.ConnectionInfo = v
	return s
}

// SetParameters sets the Parameters field's value.
func (s *DescribeInstanceDetailOutput) SetParameters(v string) *DescribeInstanceDetailOutput {
	s.Parameters = &v
	return s
}

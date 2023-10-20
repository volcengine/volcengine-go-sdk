// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package rdsmysqlv2

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribeDBInstancesCommon = "DescribeDBInstances"

// DescribeDBInstancesCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeDBInstancesCommon operation. The "output" return
// value will be populated with the DescribeDBInstancesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeDBInstancesCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeDBInstancesCommon Send returns without error.
//
// See DescribeDBInstancesCommon for more information on using the DescribeDBInstancesCommon
// API call, and error handling.
//
//	// Example sending a request using the DescribeDBInstancesCommonRequest method.
//	req, resp := client.DescribeDBInstancesCommonRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *RDSMYSQLV2) DescribeDBInstancesCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeDBInstancesCommon,
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

// DescribeDBInstancesCommon API operation for RDS_MYSQL_V2.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RDS_MYSQL_V2's
// API operation DescribeDBInstancesCommon for usage and error information.
func (c *RDSMYSQLV2) DescribeDBInstancesCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeDBInstancesCommonRequest(input)
	return out, req.Send()
}

// DescribeDBInstancesCommonWithContext is the same as DescribeDBInstancesCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeDBInstancesCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RDSMYSQLV2) DescribeDBInstancesCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeDBInstancesCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeDBInstances = "DescribeDBInstances"

// DescribeDBInstancesRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeDBInstances operation. The "output" return
// value will be populated with the DescribeDBInstancesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeDBInstancesCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeDBInstancesCommon Send returns without error.
//
// See DescribeDBInstances for more information on using the DescribeDBInstances
// API call, and error handling.
//
//	// Example sending a request using the DescribeDBInstancesRequest method.
//	req, resp := client.DescribeDBInstancesRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *RDSMYSQLV2) DescribeDBInstancesRequest(input *DescribeDBInstancesInput) (req *request.Request, output *DescribeDBInstancesOutput) {
	op := &request.Operation{
		Name:       opDescribeDBInstances,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeDBInstancesInput{}
	}

	output = &DescribeDBInstancesOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// DescribeDBInstances API operation for RDS_MYSQL_V2.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RDS_MYSQL_V2's
// API operation DescribeDBInstances for usage and error information.
func (c *RDSMYSQLV2) DescribeDBInstances(input *DescribeDBInstancesInput) (*DescribeDBInstancesOutput, error) {
	req, out := c.DescribeDBInstancesRequest(input)
	return out, req.Send()
}

// DescribeDBInstancesWithContext is the same as DescribeDBInstances with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeDBInstances for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RDSMYSQLV2) DescribeDBInstancesWithContext(ctx volcengine.Context, input *DescribeDBInstancesInput, opts ...request.Option) (*DescribeDBInstancesOutput, error) {
	req, out := c.DescribeDBInstancesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type AddressObjectForDescribeDBInstancesOutput struct {
	_ struct{} `type:"structure"`

	DNSVisibility *bool `type:"boolean"`

	Domain *string `type:"string"`

	EipId *string `type:"string"`

	IPAddress *string `type:"string"`

	NetworkType *string `type:"string"`

	Port *string `type:"string"`

	SubnetId *string `type:"string"`
}

// String returns the string representation
func (s AddressObjectForDescribeDBInstancesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s AddressObjectForDescribeDBInstancesOutput) GoString() string {
	return s.String()
}

// SetDNSVisibility sets the DNSVisibility field's value.
func (s *AddressObjectForDescribeDBInstancesOutput) SetDNSVisibility(v bool) *AddressObjectForDescribeDBInstancesOutput {
	s.DNSVisibility = &v
	return s
}

// SetDomain sets the Domain field's value.
func (s *AddressObjectForDescribeDBInstancesOutput) SetDomain(v string) *AddressObjectForDescribeDBInstancesOutput {
	s.Domain = &v
	return s
}

// SetEipId sets the EipId field's value.
func (s *AddressObjectForDescribeDBInstancesOutput) SetEipId(v string) *AddressObjectForDescribeDBInstancesOutput {
	s.EipId = &v
	return s
}

// SetIPAddress sets the IPAddress field's value.
func (s *AddressObjectForDescribeDBInstancesOutput) SetIPAddress(v string) *AddressObjectForDescribeDBInstancesOutput {
	s.IPAddress = &v
	return s
}

// SetNetworkType sets the NetworkType field's value.
func (s *AddressObjectForDescribeDBInstancesOutput) SetNetworkType(v string) *AddressObjectForDescribeDBInstancesOutput {
	s.NetworkType = &v
	return s
}

// SetPort sets the Port field's value.
func (s *AddressObjectForDescribeDBInstancesOutput) SetPort(v string) *AddressObjectForDescribeDBInstancesOutput {
	s.Port = &v
	return s
}

// SetSubnetId sets the SubnetId field's value.
func (s *AddressObjectForDescribeDBInstancesOutput) SetSubnetId(v string) *AddressObjectForDescribeDBInstancesOutput {
	s.SubnetId = &v
	return s
}

type ChargeDetailForDescribeDBInstancesOutput struct {
	_ struct{} `type:"structure"`

	AutoRenew *bool `type:"boolean"`

	ChargeEndTime *string `type:"string"`

	ChargeStartTime *string `type:"string"`

	ChargeStatus *string `type:"string"`

	ChargeType *string `type:"string"`

	Number *int32 `type:"int32"`

	OverdueReclaimTime *string `type:"string"`

	OverdueTime *string `type:"string"`

	Period *int32 `type:"int32"`

	PeriodUnit *string `type:"string"`
}

// String returns the string representation
func (s ChargeDetailForDescribeDBInstancesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ChargeDetailForDescribeDBInstancesOutput) GoString() string {
	return s.String()
}

// SetAutoRenew sets the AutoRenew field's value.
func (s *ChargeDetailForDescribeDBInstancesOutput) SetAutoRenew(v bool) *ChargeDetailForDescribeDBInstancesOutput {
	s.AutoRenew = &v
	return s
}

// SetChargeEndTime sets the ChargeEndTime field's value.
func (s *ChargeDetailForDescribeDBInstancesOutput) SetChargeEndTime(v string) *ChargeDetailForDescribeDBInstancesOutput {
	s.ChargeEndTime = &v
	return s
}

// SetChargeStartTime sets the ChargeStartTime field's value.
func (s *ChargeDetailForDescribeDBInstancesOutput) SetChargeStartTime(v string) *ChargeDetailForDescribeDBInstancesOutput {
	s.ChargeStartTime = &v
	return s
}

// SetChargeStatus sets the ChargeStatus field's value.
func (s *ChargeDetailForDescribeDBInstancesOutput) SetChargeStatus(v string) *ChargeDetailForDescribeDBInstancesOutput {
	s.ChargeStatus = &v
	return s
}

// SetChargeType sets the ChargeType field's value.
func (s *ChargeDetailForDescribeDBInstancesOutput) SetChargeType(v string) *ChargeDetailForDescribeDBInstancesOutput {
	s.ChargeType = &v
	return s
}

// SetNumber sets the Number field's value.
func (s *ChargeDetailForDescribeDBInstancesOutput) SetNumber(v int32) *ChargeDetailForDescribeDBInstancesOutput {
	s.Number = &v
	return s
}

// SetOverdueReclaimTime sets the OverdueReclaimTime field's value.
func (s *ChargeDetailForDescribeDBInstancesOutput) SetOverdueReclaimTime(v string) *ChargeDetailForDescribeDBInstancesOutput {
	s.OverdueReclaimTime = &v
	return s
}

// SetOverdueTime sets the OverdueTime field's value.
func (s *ChargeDetailForDescribeDBInstancesOutput) SetOverdueTime(v string) *ChargeDetailForDescribeDBInstancesOutput {
	s.OverdueTime = &v
	return s
}

// SetPeriod sets the Period field's value.
func (s *ChargeDetailForDescribeDBInstancesOutput) SetPeriod(v int32) *ChargeDetailForDescribeDBInstancesOutput {
	s.Period = &v
	return s
}

// SetPeriodUnit sets the PeriodUnit field's value.
func (s *ChargeDetailForDescribeDBInstancesOutput) SetPeriodUnit(v string) *ChargeDetailForDescribeDBInstancesOutput {
	s.PeriodUnit = &v
	return s
}

type DescribeDBInstancesInput struct {
	_ struct{} `type:"structure"`

	ChargeType *string `type:"string" enum:"EnumOfChargeTypeForDescribeDBInstancesInput"`

	CreateTimeEnd *string `type:"string"`

	CreateTimeStart *string `type:"string"`

	DBEngineVersion *string `type:"string" enum:"EnumOfDBEngineVersionForDescribeDBInstancesInput"`

	InstanceId *string `type:"string"`

	InstanceName *string `type:"string"`

	InstanceStatus *string `type:"string" enum:"EnumOfInstanceStatusForDescribeDBInstancesInput"`

	InstanceType *string `type:"string" enum:"EnumOfInstanceTypeForDescribeDBInstancesInput"`

	PageNumber *int32 `type:"int32"`

	PageSize *int32 `type:"int32"`

	ZoneId *string `type:"string"`
}

// String returns the string representation
func (s DescribeDBInstancesInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeDBInstancesInput) GoString() string {
	return s.String()
}

// SetChargeType sets the ChargeType field's value.
func (s *DescribeDBInstancesInput) SetChargeType(v string) *DescribeDBInstancesInput {
	s.ChargeType = &v
	return s
}

// SetCreateTimeEnd sets the CreateTimeEnd field's value.
func (s *DescribeDBInstancesInput) SetCreateTimeEnd(v string) *DescribeDBInstancesInput {
	s.CreateTimeEnd = &v
	return s
}

// SetCreateTimeStart sets the CreateTimeStart field's value.
func (s *DescribeDBInstancesInput) SetCreateTimeStart(v string) *DescribeDBInstancesInput {
	s.CreateTimeStart = &v
	return s
}

// SetDBEngineVersion sets the DBEngineVersion field's value.
func (s *DescribeDBInstancesInput) SetDBEngineVersion(v string) *DescribeDBInstancesInput {
	s.DBEngineVersion = &v
	return s
}

// SetInstanceId sets the InstanceId field's value.
func (s *DescribeDBInstancesInput) SetInstanceId(v string) *DescribeDBInstancesInput {
	s.InstanceId = &v
	return s
}

// SetInstanceName sets the InstanceName field's value.
func (s *DescribeDBInstancesInput) SetInstanceName(v string) *DescribeDBInstancesInput {
	s.InstanceName = &v
	return s
}

// SetInstanceStatus sets the InstanceStatus field's value.
func (s *DescribeDBInstancesInput) SetInstanceStatus(v string) *DescribeDBInstancesInput {
	s.InstanceStatus = &v
	return s
}

// SetInstanceType sets the InstanceType field's value.
func (s *DescribeDBInstancesInput) SetInstanceType(v string) *DescribeDBInstancesInput {
	s.InstanceType = &v
	return s
}

// SetPageNumber sets the PageNumber field's value.
func (s *DescribeDBInstancesInput) SetPageNumber(v int32) *DescribeDBInstancesInput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *DescribeDBInstancesInput) SetPageSize(v int32) *DescribeDBInstancesInput {
	s.PageSize = &v
	return s
}

// SetZoneId sets the ZoneId field's value.
func (s *DescribeDBInstancesInput) SetZoneId(v string) *DescribeDBInstancesInput {
	s.ZoneId = &v
	return s
}

type DescribeDBInstancesOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	Instances []*InstanceForDescribeDBInstancesOutput `type:"list"`

	Total *int32 `type:"int32"`
}

// String returns the string representation
func (s DescribeDBInstancesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeDBInstancesOutput) GoString() string {
	return s.String()
}

// SetInstances sets the Instances field's value.
func (s *DescribeDBInstancesOutput) SetInstances(v []*InstanceForDescribeDBInstancesOutput) *DescribeDBInstancesOutput {
	s.Instances = v
	return s
}

// SetTotal sets the Total field's value.
func (s *DescribeDBInstancesOutput) SetTotal(v int32) *DescribeDBInstancesOutput {
	s.Total = &v
	return s
}

type InstanceForDescribeDBInstancesOutput struct {
	_ struct{} `type:"structure"`

	AccountId *string `type:"string"`

	AddressObject []*AddressObjectForDescribeDBInstancesOutput `type:"list"`

	AllowListVersion *string `type:"string"`

	ChargeDetail *ChargeDetailForDescribeDBInstancesOutput `type:"structure"`

	CreateTime *string `type:"string"`

	DBEngineVersion *string `type:"string"`

	InstanceId *string `type:"string"`

	InstanceName *string `type:"string"`

	InstanceStatus *string `type:"string"`

	InstanceType *string `type:"string"`

	LowerCaseTableNames *string `type:"string"`

	NodeDetailInfo []*NodeDetailInfoForDescribeDBInstancesOutput `type:"list"`

	NodeNumber *int32 `type:"int32"`

	NodeSpec *string `type:"string"`

	Port *string `type:"string"`

	ProjectName *string `type:"string"`

	RegionId *string `type:"string"`

	ServerCollation *string `type:"string"`

	ShardNumber *int32 `type:"int32"`

	StorageSpace *int32 `type:"int32"`

	StorageType *string `type:"string"`

	StorageUse *int32 `type:"int32"`

	SubnetId *string `type:"string"`

	TimeZone *string `type:"string"`

	VpcId *string `type:"string"`

	ZoneId *string `type:"string"`
}

// String returns the string representation
func (s InstanceForDescribeDBInstancesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s InstanceForDescribeDBInstancesOutput) GoString() string {
	return s.String()
}

// SetAccountId sets the AccountId field's value.
func (s *InstanceForDescribeDBInstancesOutput) SetAccountId(v string) *InstanceForDescribeDBInstancesOutput {
	s.AccountId = &v
	return s
}

// SetAddressObject sets the AddressObject field's value.
func (s *InstanceForDescribeDBInstancesOutput) SetAddressObject(v []*AddressObjectForDescribeDBInstancesOutput) *InstanceForDescribeDBInstancesOutput {
	s.AddressObject = v
	return s
}

// SetAllowListVersion sets the AllowListVersion field's value.
func (s *InstanceForDescribeDBInstancesOutput) SetAllowListVersion(v string) *InstanceForDescribeDBInstancesOutput {
	s.AllowListVersion = &v
	return s
}

// SetChargeDetail sets the ChargeDetail field's value.
func (s *InstanceForDescribeDBInstancesOutput) SetChargeDetail(v *ChargeDetailForDescribeDBInstancesOutput) *InstanceForDescribeDBInstancesOutput {
	s.ChargeDetail = v
	return s
}

// SetCreateTime sets the CreateTime field's value.
func (s *InstanceForDescribeDBInstancesOutput) SetCreateTime(v string) *InstanceForDescribeDBInstancesOutput {
	s.CreateTime = &v
	return s
}

// SetDBEngineVersion sets the DBEngineVersion field's value.
func (s *InstanceForDescribeDBInstancesOutput) SetDBEngineVersion(v string) *InstanceForDescribeDBInstancesOutput {
	s.DBEngineVersion = &v
	return s
}

// SetInstanceId sets the InstanceId field's value.
func (s *InstanceForDescribeDBInstancesOutput) SetInstanceId(v string) *InstanceForDescribeDBInstancesOutput {
	s.InstanceId = &v
	return s
}

// SetInstanceName sets the InstanceName field's value.
func (s *InstanceForDescribeDBInstancesOutput) SetInstanceName(v string) *InstanceForDescribeDBInstancesOutput {
	s.InstanceName = &v
	return s
}

// SetInstanceStatus sets the InstanceStatus field's value.
func (s *InstanceForDescribeDBInstancesOutput) SetInstanceStatus(v string) *InstanceForDescribeDBInstancesOutput {
	s.InstanceStatus = &v
	return s
}

// SetInstanceType sets the InstanceType field's value.
func (s *InstanceForDescribeDBInstancesOutput) SetInstanceType(v string) *InstanceForDescribeDBInstancesOutput {
	s.InstanceType = &v
	return s
}

// SetLowerCaseTableNames sets the LowerCaseTableNames field's value.
func (s *InstanceForDescribeDBInstancesOutput) SetLowerCaseTableNames(v string) *InstanceForDescribeDBInstancesOutput {
	s.LowerCaseTableNames = &v
	return s
}

// SetNodeDetailInfo sets the NodeDetailInfo field's value.
func (s *InstanceForDescribeDBInstancesOutput) SetNodeDetailInfo(v []*NodeDetailInfoForDescribeDBInstancesOutput) *InstanceForDescribeDBInstancesOutput {
	s.NodeDetailInfo = v
	return s
}

// SetNodeNumber sets the NodeNumber field's value.
func (s *InstanceForDescribeDBInstancesOutput) SetNodeNumber(v int32) *InstanceForDescribeDBInstancesOutput {
	s.NodeNumber = &v
	return s
}

// SetNodeSpec sets the NodeSpec field's value.
func (s *InstanceForDescribeDBInstancesOutput) SetNodeSpec(v string) *InstanceForDescribeDBInstancesOutput {
	s.NodeSpec = &v
	return s
}

// SetPort sets the Port field's value.
func (s *InstanceForDescribeDBInstancesOutput) SetPort(v string) *InstanceForDescribeDBInstancesOutput {
	s.Port = &v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *InstanceForDescribeDBInstancesOutput) SetProjectName(v string) *InstanceForDescribeDBInstancesOutput {
	s.ProjectName = &v
	return s
}

// SetRegionId sets the RegionId field's value.
func (s *InstanceForDescribeDBInstancesOutput) SetRegionId(v string) *InstanceForDescribeDBInstancesOutput {
	s.RegionId = &v
	return s
}

// SetServerCollation sets the ServerCollation field's value.
func (s *InstanceForDescribeDBInstancesOutput) SetServerCollation(v string) *InstanceForDescribeDBInstancesOutput {
	s.ServerCollation = &v
	return s
}

// SetShardNumber sets the ShardNumber field's value.
func (s *InstanceForDescribeDBInstancesOutput) SetShardNumber(v int32) *InstanceForDescribeDBInstancesOutput {
	s.ShardNumber = &v
	return s
}

// SetStorageSpace sets the StorageSpace field's value.
func (s *InstanceForDescribeDBInstancesOutput) SetStorageSpace(v int32) *InstanceForDescribeDBInstancesOutput {
	s.StorageSpace = &v
	return s
}

// SetStorageType sets the StorageType field's value.
func (s *InstanceForDescribeDBInstancesOutput) SetStorageType(v string) *InstanceForDescribeDBInstancesOutput {
	s.StorageType = &v
	return s
}

// SetStorageUse sets the StorageUse field's value.
func (s *InstanceForDescribeDBInstancesOutput) SetStorageUse(v int32) *InstanceForDescribeDBInstancesOutput {
	s.StorageUse = &v
	return s
}

// SetSubnetId sets the SubnetId field's value.
func (s *InstanceForDescribeDBInstancesOutput) SetSubnetId(v string) *InstanceForDescribeDBInstancesOutput {
	s.SubnetId = &v
	return s
}

// SetTimeZone sets the TimeZone field's value.
func (s *InstanceForDescribeDBInstancesOutput) SetTimeZone(v string) *InstanceForDescribeDBInstancesOutput {
	s.TimeZone = &v
	return s
}

// SetVpcId sets the VpcId field's value.
func (s *InstanceForDescribeDBInstancesOutput) SetVpcId(v string) *InstanceForDescribeDBInstancesOutput {
	s.VpcId = &v
	return s
}

// SetZoneId sets the ZoneId field's value.
func (s *InstanceForDescribeDBInstancesOutput) SetZoneId(v string) *InstanceForDescribeDBInstancesOutput {
	s.ZoneId = &v
	return s
}

type NodeDetailInfoForDescribeDBInstancesOutput struct {
	_ struct{} `type:"structure"`

	CreateTime *string `type:"string"`

	InstanceId *string `type:"string"`

	Memory *int32 `type:"int32"`

	NodeId *string `type:"string"`

	NodeSpec *string `type:"string"`

	NodeStatus *string `type:"string"`

	NodeType *string `type:"string"`

	RegionId *string `type:"string"`

	ShardId *string `type:"string"`

	SyncDelay *int64 `type:"int64"`

	UpdateTime *string `type:"string"`

	VCPU *int32 `type:"int32"`

	ZoneId *string `type:"string"`
}

// String returns the string representation
func (s NodeDetailInfoForDescribeDBInstancesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s NodeDetailInfoForDescribeDBInstancesOutput) GoString() string {
	return s.String()
}

// SetCreateTime sets the CreateTime field's value.
func (s *NodeDetailInfoForDescribeDBInstancesOutput) SetCreateTime(v string) *NodeDetailInfoForDescribeDBInstancesOutput {
	s.CreateTime = &v
	return s
}

// SetInstanceId sets the InstanceId field's value.
func (s *NodeDetailInfoForDescribeDBInstancesOutput) SetInstanceId(v string) *NodeDetailInfoForDescribeDBInstancesOutput {
	s.InstanceId = &v
	return s
}

// SetMemory sets the Memory field's value.
func (s *NodeDetailInfoForDescribeDBInstancesOutput) SetMemory(v int32) *NodeDetailInfoForDescribeDBInstancesOutput {
	s.Memory = &v
	return s
}

// SetNodeId sets the NodeId field's value.
func (s *NodeDetailInfoForDescribeDBInstancesOutput) SetNodeId(v string) *NodeDetailInfoForDescribeDBInstancesOutput {
	s.NodeId = &v
	return s
}

// SetNodeSpec sets the NodeSpec field's value.
func (s *NodeDetailInfoForDescribeDBInstancesOutput) SetNodeSpec(v string) *NodeDetailInfoForDescribeDBInstancesOutput {
	s.NodeSpec = &v
	return s
}

// SetNodeStatus sets the NodeStatus field's value.
func (s *NodeDetailInfoForDescribeDBInstancesOutput) SetNodeStatus(v string) *NodeDetailInfoForDescribeDBInstancesOutput {
	s.NodeStatus = &v
	return s
}

// SetNodeType sets the NodeType field's value.
func (s *NodeDetailInfoForDescribeDBInstancesOutput) SetNodeType(v string) *NodeDetailInfoForDescribeDBInstancesOutput {
	s.NodeType = &v
	return s
}

// SetRegionId sets the RegionId field's value.
func (s *NodeDetailInfoForDescribeDBInstancesOutput) SetRegionId(v string) *NodeDetailInfoForDescribeDBInstancesOutput {
	s.RegionId = &v
	return s
}

// SetShardId sets the ShardId field's value.
func (s *NodeDetailInfoForDescribeDBInstancesOutput) SetShardId(v string) *NodeDetailInfoForDescribeDBInstancesOutput {
	s.ShardId = &v
	return s
}

// SetSyncDelay sets the SyncDelay field's value.
func (s *NodeDetailInfoForDescribeDBInstancesOutput) SetSyncDelay(v int64) *NodeDetailInfoForDescribeDBInstancesOutput {
	s.SyncDelay = &v
	return s
}

// SetUpdateTime sets the UpdateTime field's value.
func (s *NodeDetailInfoForDescribeDBInstancesOutput) SetUpdateTime(v string) *NodeDetailInfoForDescribeDBInstancesOutput {
	s.UpdateTime = &v
	return s
}

// SetVCPU sets the VCPU field's value.
func (s *NodeDetailInfoForDescribeDBInstancesOutput) SetVCPU(v int32) *NodeDetailInfoForDescribeDBInstancesOutput {
	s.VCPU = &v
	return s
}

// SetZoneId sets the ZoneId field's value.
func (s *NodeDetailInfoForDescribeDBInstancesOutput) SetZoneId(v string) *NodeDetailInfoForDescribeDBInstancesOutput {
	s.ZoneId = &v
	return s
}

const (
	// EnumOfChargeTypeForDescribeDBInstancesInputNotEnabled is a EnumOfChargeTypeForDescribeDBInstancesInput enum value
	EnumOfChargeTypeForDescribeDBInstancesInputNotEnabled = "NotEnabled"

	// EnumOfChargeTypeForDescribeDBInstancesInputPostPaid is a EnumOfChargeTypeForDescribeDBInstancesInput enum value
	EnumOfChargeTypeForDescribeDBInstancesInputPostPaid = "PostPaid"

	// EnumOfChargeTypeForDescribeDBInstancesInputPrePaid is a EnumOfChargeTypeForDescribeDBInstancesInput enum value
	EnumOfChargeTypeForDescribeDBInstancesInputPrePaid = "PrePaid"
)

const (
	// EnumOfDBEngineVersionForDescribeDBInstancesInputMySql56 is a EnumOfDBEngineVersionForDescribeDBInstancesInput enum value
	EnumOfDBEngineVersionForDescribeDBInstancesInputMySql56 = "MySQL_5_6"

	// EnumOfDBEngineVersionForDescribeDBInstancesInputMySql57 is a EnumOfDBEngineVersionForDescribeDBInstancesInput enum value
	EnumOfDBEngineVersionForDescribeDBInstancesInputMySql57 = "MySQL_5_7"

	// EnumOfDBEngineVersionForDescribeDBInstancesInputMySql80 is a EnumOfDBEngineVersionForDescribeDBInstancesInput enum value
	EnumOfDBEngineVersionForDescribeDBInstancesInputMySql80 = "MySQL_8_0"

	// EnumOfDBEngineVersionForDescribeDBInstancesInputSqlserver2019Ent is a EnumOfDBEngineVersionForDescribeDBInstancesInput enum value
	EnumOfDBEngineVersionForDescribeDBInstancesInputSqlserver2019Ent = "SQLServer_2019_Ent"

	// EnumOfDBEngineVersionForDescribeDBInstancesInputSqlserver2019Std is a EnumOfDBEngineVersionForDescribeDBInstancesInput enum value
	EnumOfDBEngineVersionForDescribeDBInstancesInputSqlserver2019Std = "SQLServer_2019_Std"

	// EnumOfDBEngineVersionForDescribeDBInstancesInputSqlserver2019Web is a EnumOfDBEngineVersionForDescribeDBInstancesInput enum value
	EnumOfDBEngineVersionForDescribeDBInstancesInputSqlserver2019Web = "SQLServer_2019_Web"
)

const (
	// EnumOfInstanceStatusForDescribeDBInstancesInputAllowListMaintaining is a EnumOfInstanceStatusForDescribeDBInstancesInput enum value
	EnumOfInstanceStatusForDescribeDBInstancesInputAllowListMaintaining = "AllowListMaintaining"

	// EnumOfInstanceStatusForDescribeDBInstancesInputClosed is a EnumOfInstanceStatusForDescribeDBInstancesInput enum value
	EnumOfInstanceStatusForDescribeDBInstancesInputClosed = "Closed"

	// EnumOfInstanceStatusForDescribeDBInstancesInputClosing is a EnumOfInstanceStatusForDescribeDBInstancesInput enum value
	EnumOfInstanceStatusForDescribeDBInstancesInputClosing = "Closing"

	// EnumOfInstanceStatusForDescribeDBInstancesInputCreateFailed is a EnumOfInstanceStatusForDescribeDBInstancesInput enum value
	EnumOfInstanceStatusForDescribeDBInstancesInputCreateFailed = "CreateFailed"

	// EnumOfInstanceStatusForDescribeDBInstancesInputCreating is a EnumOfInstanceStatusForDescribeDBInstancesInput enum value
	EnumOfInstanceStatusForDescribeDBInstancesInputCreating = "Creating"

	// EnumOfInstanceStatusForDescribeDBInstancesInputDeleting is a EnumOfInstanceStatusForDescribeDBInstancesInput enum value
	EnumOfInstanceStatusForDescribeDBInstancesInputDeleting = "Deleting"

	// EnumOfInstanceStatusForDescribeDBInstancesInputDestroyed is a EnumOfInstanceStatusForDescribeDBInstancesInput enum value
	EnumOfInstanceStatusForDescribeDBInstancesInputDestroyed = "Destroyed"

	// EnumOfInstanceStatusForDescribeDBInstancesInputDestroying is a EnumOfInstanceStatusForDescribeDBInstancesInput enum value
	EnumOfInstanceStatusForDescribeDBInstancesInputDestroying = "Destroying"

	// EnumOfInstanceStatusForDescribeDBInstancesInputError is a EnumOfInstanceStatusForDescribeDBInstancesInput enum value
	EnumOfInstanceStatusForDescribeDBInstancesInputError = "Error"

	// EnumOfInstanceStatusForDescribeDBInstancesInputImporting is a EnumOfInstanceStatusForDescribeDBInstancesInput enum value
	EnumOfInstanceStatusForDescribeDBInstancesInputImporting = "Importing"

	// EnumOfInstanceStatusForDescribeDBInstancesInputMaintaining is a EnumOfInstanceStatusForDescribeDBInstancesInput enum value
	EnumOfInstanceStatusForDescribeDBInstancesInputMaintaining = "Maintaining"

	// EnumOfInstanceStatusForDescribeDBInstancesInputMasterChanging is a EnumOfInstanceStatusForDescribeDBInstancesInput enum value
	EnumOfInstanceStatusForDescribeDBInstancesInputMasterChanging = "MasterChanging"

	// EnumOfInstanceStatusForDescribeDBInstancesInputMigrating is a EnumOfInstanceStatusForDescribeDBInstancesInput enum value
	EnumOfInstanceStatusForDescribeDBInstancesInputMigrating = "Migrating"

	// EnumOfInstanceStatusForDescribeDBInstancesInputProxyCreating is a EnumOfInstanceStatusForDescribeDBInstancesInput enum value
	EnumOfInstanceStatusForDescribeDBInstancesInputProxyCreating = "ProxyCreating"

	// EnumOfInstanceStatusForDescribeDBInstancesInputProxyDeleting is a EnumOfInstanceStatusForDescribeDBInstancesInput enum value
	EnumOfInstanceStatusForDescribeDBInstancesInputProxyDeleting = "ProxyDeleting"

	// EnumOfInstanceStatusForDescribeDBInstancesInputReclaiming is a EnumOfInstanceStatusForDescribeDBInstancesInput enum value
	EnumOfInstanceStatusForDescribeDBInstancesInputReclaiming = "Reclaiming"

	// EnumOfInstanceStatusForDescribeDBInstancesInputRecycled is a EnumOfInstanceStatusForDescribeDBInstancesInput enum value
	EnumOfInstanceStatusForDescribeDBInstancesInputRecycled = "Recycled"

	// EnumOfInstanceStatusForDescribeDBInstancesInputReleased is a EnumOfInstanceStatusForDescribeDBInstancesInput enum value
	EnumOfInstanceStatusForDescribeDBInstancesInputReleased = "Released"

	// EnumOfInstanceStatusForDescribeDBInstancesInputRestarting is a EnumOfInstanceStatusForDescribeDBInstancesInput enum value
	EnumOfInstanceStatusForDescribeDBInstancesInputRestarting = "Restarting"

	// EnumOfInstanceStatusForDescribeDBInstancesInputRestoring is a EnumOfInstanceStatusForDescribeDBInstancesInput enum value
	EnumOfInstanceStatusForDescribeDBInstancesInputRestoring = "Restoring"

	// EnumOfInstanceStatusForDescribeDBInstancesInputResuming is a EnumOfInstanceStatusForDescribeDBInstancesInput enum value
	EnumOfInstanceStatusForDescribeDBInstancesInputResuming = "Resuming"

	// EnumOfInstanceStatusForDescribeDBInstancesInputRunning is a EnumOfInstanceStatusForDescribeDBInstancesInput enum value
	EnumOfInstanceStatusForDescribeDBInstancesInputRunning = "Running"

	// EnumOfInstanceStatusForDescribeDBInstancesInputSslupdating is a EnumOfInstanceStatusForDescribeDBInstancesInput enum value
	EnumOfInstanceStatusForDescribeDBInstancesInputSslupdating = "SSLUpdating"

	// EnumOfInstanceStatusForDescribeDBInstancesInputTdeupdating is a EnumOfInstanceStatusForDescribeDBInstancesInput enum value
	EnumOfInstanceStatusForDescribeDBInstancesInputTdeupdating = "TDEUpdating"

	// EnumOfInstanceStatusForDescribeDBInstancesInputUnknown is a EnumOfInstanceStatusForDescribeDBInstancesInput enum value
	EnumOfInstanceStatusForDescribeDBInstancesInputUnknown = "Unknown"

	// EnumOfInstanceStatusForDescribeDBInstancesInputUpdating is a EnumOfInstanceStatusForDescribeDBInstancesInput enum value
	EnumOfInstanceStatusForDescribeDBInstancesInputUpdating = "Updating"

	// EnumOfInstanceStatusForDescribeDBInstancesInputUpgrading is a EnumOfInstanceStatusForDescribeDBInstancesInput enum value
	EnumOfInstanceStatusForDescribeDBInstancesInputUpgrading = "Upgrading"

	// EnumOfInstanceStatusForDescribeDBInstancesInputWaitingPaid is a EnumOfInstanceStatusForDescribeDBInstancesInput enum value
	EnumOfInstanceStatusForDescribeDBInstancesInputWaitingPaid = "WaitingPaid"
)

const (
	// EnumOfInstanceTypeForDescribeDBInstancesInputBasic is a EnumOfInstanceTypeForDescribeDBInstancesInput enum value
	EnumOfInstanceTypeForDescribeDBInstancesInputBasic = "Basic"

	// EnumOfInstanceTypeForDescribeDBInstancesInputCluster is a EnumOfInstanceTypeForDescribeDBInstancesInput enum value
	EnumOfInstanceTypeForDescribeDBInstancesInputCluster = "Cluster"

	// EnumOfInstanceTypeForDescribeDBInstancesInputDoubleNode is a EnumOfInstanceTypeForDescribeDBInstancesInput enum value
	EnumOfInstanceTypeForDescribeDBInstancesInputDoubleNode = "DoubleNode"

	// EnumOfInstanceTypeForDescribeDBInstancesInputHa is a EnumOfInstanceTypeForDescribeDBInstancesInput enum value
	EnumOfInstanceTypeForDescribeDBInstancesInputHa = "HA"

	// EnumOfInstanceTypeForDescribeDBInstancesInputMultiNode is a EnumOfInstanceTypeForDescribeDBInstancesInput enum value
	EnumOfInstanceTypeForDescribeDBInstancesInputMultiNode = "MultiNode"
)

// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package rdsmssql

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
//    // Example sending a request using the DescribeDBInstancesCommonRequest method.
//    req, resp := client.DescribeDBInstancesCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *RDSMSSQL) DescribeDBInstancesCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
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

// DescribeDBInstancesCommon API operation for RDS_MSSQL.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RDS_MSSQL's
// API operation DescribeDBInstancesCommon for usage and error information.
func (c *RDSMSSQL) DescribeDBInstancesCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
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
func (c *RDSMSSQL) DescribeDBInstancesCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
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
//    // Example sending a request using the DescribeDBInstancesRequest method.
//    req, resp := client.DescribeDBInstancesRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *RDSMSSQL) DescribeDBInstancesRequest(input *DescribeDBInstancesInput) (req *request.Request, output *DescribeDBInstancesOutput) {
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

// DescribeDBInstances API operation for RDS_MSSQL.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RDS_MSSQL's
// API operation DescribeDBInstances for usage and error information.
func (c *RDSMSSQL) DescribeDBInstances(input *DescribeDBInstancesInput) (*DescribeDBInstancesOutput, error) {
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
func (c *RDSMSSQL) DescribeDBInstancesWithContext(ctx volcengine.Context, input *DescribeDBInstancesInput, opts ...request.Option) (*DescribeDBInstancesOutput, error) {
	req, out := c.DescribeDBInstancesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ChargeDetailForDescribeDBInstancesOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	AutoRenew *bool `type:"boolean" json:",omitempty"`

	ChargeEndTime *string `type:"string" json:",omitempty"`

	ChargeStartTime *string `type:"string" json:",omitempty"`

	ChargeStatus *string `type:"string" json:",omitempty"`

	ChargeType *string `type:"string" json:",omitempty"`

	OverdueReclaimTime *string `type:"string" json:",omitempty"`

	OverdueTime *string `type:"string" json:",omitempty"`

	Period *int32 `type:"int32" json:",omitempty"`

	PeriodUnit *string `type:"string" json:",omitempty"`
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
	_ struct{} `type:"structure" json:",omitempty"`

	ChargeType *string `type:"string" json:",omitempty"`

	CreateTimeEnd *string `type:"string" json:",omitempty"`

	CreateTimeStart *string `type:"string" json:",omitempty"`

	DBEngineVersion *string `type:"string" json:",omitempty"`

	InstanceCategory *string `type:"string" json:",omitempty"`

	InstanceId *string `type:"string" json:",omitempty"`

	InstanceName *string `type:"string" json:",omitempty"`

	InstanceStatus *string `type:"string" json:",omitempty"`

	InstanceType *string `type:"string" json:",omitempty"`

	PageNumber *int32 `type:"int32" json:",omitempty"`

	PageSize *int32 `type:"int32" json:",omitempty"`

	PrimaryInstanceId *string `type:"string" json:",omitempty"`

	ProjectName *string `type:"string" json:",omitempty"`

	ServerCollation *string `type:"string" json:",omitempty"`

	TagFilters []*TagFilterForDescribeDBInstancesInput `type:"list" json:",omitempty"`

	ZoneId *string `type:"string" json:",omitempty"`
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

// SetInstanceCategory sets the InstanceCategory field's value.
func (s *DescribeDBInstancesInput) SetInstanceCategory(v string) *DescribeDBInstancesInput {
	s.InstanceCategory = &v
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

// SetPrimaryInstanceId sets the PrimaryInstanceId field's value.
func (s *DescribeDBInstancesInput) SetPrimaryInstanceId(v string) *DescribeDBInstancesInput {
	s.PrimaryInstanceId = &v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *DescribeDBInstancesInput) SetProjectName(v string) *DescribeDBInstancesInput {
	s.ProjectName = &v
	return s
}

// SetServerCollation sets the ServerCollation field's value.
func (s *DescribeDBInstancesInput) SetServerCollation(v string) *DescribeDBInstancesInput {
	s.ServerCollation = &v
	return s
}

// SetTagFilters sets the TagFilters field's value.
func (s *DescribeDBInstancesInput) SetTagFilters(v []*TagFilterForDescribeDBInstancesInput) *DescribeDBInstancesInput {
	s.TagFilters = v
	return s
}

// SetZoneId sets the ZoneId field's value.
func (s *DescribeDBInstancesInput) SetZoneId(v string) *DescribeDBInstancesInput {
	s.ZoneId = &v
	return s
}

type DescribeDBInstancesOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	InstancesInfo []*InstancesInfoForDescribeDBInstancesOutput `type:"list" json:",omitempty"`

	Total *int32 `type:"int32" json:",omitempty"`
}

// String returns the string representation
func (s DescribeDBInstancesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeDBInstancesOutput) GoString() string {
	return s.String()
}

// SetInstancesInfo sets the InstancesInfo field's value.
func (s *DescribeDBInstancesOutput) SetInstancesInfo(v []*InstancesInfoForDescribeDBInstancesOutput) *DescribeDBInstancesOutput {
	s.InstancesInfo = v
	return s
}

// SetTotal sets the Total field's value.
func (s *DescribeDBInstancesOutput) SetTotal(v int32) *DescribeDBInstancesOutput {
	s.Total = &v
	return s
}

type InstancesInfoForDescribeDBInstancesOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	ChargeDetail *ChargeDetailForDescribeDBInstancesOutput `type:"structure" json:",omitempty"`

	CreateTime *string `type:"string" json:",omitempty"`

	DBEngineVersion *string `type:"string" json:",omitempty"`

	InstanceCategory *string `type:"string" json:",omitempty"`

	InstanceId *string `type:"string" json:",omitempty"`

	InstanceName *string `type:"string" json:",omitempty"`

	InstanceStatus *string `type:"string" json:",omitempty"`

	InstanceType *string `type:"string" json:",omitempty"`

	MaintenanceTime *string `type:"string" json:",omitempty"`

	NodeDetailInfo []*NodeDetailInfoForDescribeDBInstancesOutput `type:"list" json:",omitempty"`

	NodeSpec *string `type:"string" json:",omitempty"`

	Port *string `type:"string" json:",omitempty"`

	PrimaryInstanceId *string `type:"string" json:",omitempty"`

	ProjectName *string `type:"string" json:",omitempty"`

	ReadOnlyNumber *int32 `type:"int32" json:",omitempty"`

	RegionId *string `type:"string" json:",omitempty"`

	ServerCollation *string `type:"string" json:",omitempty"`

	StorageSpace *int32 `type:"int32" json:",omitempty"`

	StorageType *string `type:"string" json:",omitempty"`

	SubnetId *string `type:"string" json:",omitempty"`

	Tags []*TagForDescribeDBInstancesOutput `type:"list" json:",omitempty"`

	TimeZone *string `type:"string" json:",omitempty"`

	VpcId *string `type:"string" json:",omitempty"`

	ZoneId *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s InstancesInfoForDescribeDBInstancesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s InstancesInfoForDescribeDBInstancesOutput) GoString() string {
	return s.String()
}

// SetChargeDetail sets the ChargeDetail field's value.
func (s *InstancesInfoForDescribeDBInstancesOutput) SetChargeDetail(v *ChargeDetailForDescribeDBInstancesOutput) *InstancesInfoForDescribeDBInstancesOutput {
	s.ChargeDetail = v
	return s
}

// SetCreateTime sets the CreateTime field's value.
func (s *InstancesInfoForDescribeDBInstancesOutput) SetCreateTime(v string) *InstancesInfoForDescribeDBInstancesOutput {
	s.CreateTime = &v
	return s
}

// SetDBEngineVersion sets the DBEngineVersion field's value.
func (s *InstancesInfoForDescribeDBInstancesOutput) SetDBEngineVersion(v string) *InstancesInfoForDescribeDBInstancesOutput {
	s.DBEngineVersion = &v
	return s
}

// SetInstanceCategory sets the InstanceCategory field's value.
func (s *InstancesInfoForDescribeDBInstancesOutput) SetInstanceCategory(v string) *InstancesInfoForDescribeDBInstancesOutput {
	s.InstanceCategory = &v
	return s
}

// SetInstanceId sets the InstanceId field's value.
func (s *InstancesInfoForDescribeDBInstancesOutput) SetInstanceId(v string) *InstancesInfoForDescribeDBInstancesOutput {
	s.InstanceId = &v
	return s
}

// SetInstanceName sets the InstanceName field's value.
func (s *InstancesInfoForDescribeDBInstancesOutput) SetInstanceName(v string) *InstancesInfoForDescribeDBInstancesOutput {
	s.InstanceName = &v
	return s
}

// SetInstanceStatus sets the InstanceStatus field's value.
func (s *InstancesInfoForDescribeDBInstancesOutput) SetInstanceStatus(v string) *InstancesInfoForDescribeDBInstancesOutput {
	s.InstanceStatus = &v
	return s
}

// SetInstanceType sets the InstanceType field's value.
func (s *InstancesInfoForDescribeDBInstancesOutput) SetInstanceType(v string) *InstancesInfoForDescribeDBInstancesOutput {
	s.InstanceType = &v
	return s
}

// SetMaintenanceTime sets the MaintenanceTime field's value.
func (s *InstancesInfoForDescribeDBInstancesOutput) SetMaintenanceTime(v string) *InstancesInfoForDescribeDBInstancesOutput {
	s.MaintenanceTime = &v
	return s
}

// SetNodeDetailInfo sets the NodeDetailInfo field's value.
func (s *InstancesInfoForDescribeDBInstancesOutput) SetNodeDetailInfo(v []*NodeDetailInfoForDescribeDBInstancesOutput) *InstancesInfoForDescribeDBInstancesOutput {
	s.NodeDetailInfo = v
	return s
}

// SetNodeSpec sets the NodeSpec field's value.
func (s *InstancesInfoForDescribeDBInstancesOutput) SetNodeSpec(v string) *InstancesInfoForDescribeDBInstancesOutput {
	s.NodeSpec = &v
	return s
}

// SetPort sets the Port field's value.
func (s *InstancesInfoForDescribeDBInstancesOutput) SetPort(v string) *InstancesInfoForDescribeDBInstancesOutput {
	s.Port = &v
	return s
}

// SetPrimaryInstanceId sets the PrimaryInstanceId field's value.
func (s *InstancesInfoForDescribeDBInstancesOutput) SetPrimaryInstanceId(v string) *InstancesInfoForDescribeDBInstancesOutput {
	s.PrimaryInstanceId = &v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *InstancesInfoForDescribeDBInstancesOutput) SetProjectName(v string) *InstancesInfoForDescribeDBInstancesOutput {
	s.ProjectName = &v
	return s
}

// SetReadOnlyNumber sets the ReadOnlyNumber field's value.
func (s *InstancesInfoForDescribeDBInstancesOutput) SetReadOnlyNumber(v int32) *InstancesInfoForDescribeDBInstancesOutput {
	s.ReadOnlyNumber = &v
	return s
}

// SetRegionId sets the RegionId field's value.
func (s *InstancesInfoForDescribeDBInstancesOutput) SetRegionId(v string) *InstancesInfoForDescribeDBInstancesOutput {
	s.RegionId = &v
	return s
}

// SetServerCollation sets the ServerCollation field's value.
func (s *InstancesInfoForDescribeDBInstancesOutput) SetServerCollation(v string) *InstancesInfoForDescribeDBInstancesOutput {
	s.ServerCollation = &v
	return s
}

// SetStorageSpace sets the StorageSpace field's value.
func (s *InstancesInfoForDescribeDBInstancesOutput) SetStorageSpace(v int32) *InstancesInfoForDescribeDBInstancesOutput {
	s.StorageSpace = &v
	return s
}

// SetStorageType sets the StorageType field's value.
func (s *InstancesInfoForDescribeDBInstancesOutput) SetStorageType(v string) *InstancesInfoForDescribeDBInstancesOutput {
	s.StorageType = &v
	return s
}

// SetSubnetId sets the SubnetId field's value.
func (s *InstancesInfoForDescribeDBInstancesOutput) SetSubnetId(v string) *InstancesInfoForDescribeDBInstancesOutput {
	s.SubnetId = &v
	return s
}

// SetTags sets the Tags field's value.
func (s *InstancesInfoForDescribeDBInstancesOutput) SetTags(v []*TagForDescribeDBInstancesOutput) *InstancesInfoForDescribeDBInstancesOutput {
	s.Tags = v
	return s
}

// SetTimeZone sets the TimeZone field's value.
func (s *InstancesInfoForDescribeDBInstancesOutput) SetTimeZone(v string) *InstancesInfoForDescribeDBInstancesOutput {
	s.TimeZone = &v
	return s
}

// SetVpcId sets the VpcId field's value.
func (s *InstancesInfoForDescribeDBInstancesOutput) SetVpcId(v string) *InstancesInfoForDescribeDBInstancesOutput {
	s.VpcId = &v
	return s
}

// SetZoneId sets the ZoneId field's value.
func (s *InstancesInfoForDescribeDBInstancesOutput) SetZoneId(v string) *InstancesInfoForDescribeDBInstancesOutput {
	s.ZoneId = &v
	return s
}

type NodeDetailInfoForDescribeDBInstancesOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	CreateTime *string `type:"string" json:",omitempty"`

	InstanceId *string `type:"string" json:",omitempty"`

	Memory *int32 `type:"int32" json:",omitempty"`

	NodeIP *string `type:"string" json:",omitempty"`

	NodeId *string `type:"string" json:",omitempty"`

	NodeSpec *string `type:"string" json:",omitempty"`

	NodeStatus *string `type:"string" json:",omitempty"`

	NodeType *string `type:"string" json:",omitempty"`

	RegionId *string `type:"string" json:",omitempty"`

	UpdateTime *string `type:"string" json:",omitempty"`

	VCPU *int32 `type:"int32" json:",omitempty"`

	ZoneId *string `type:"string" json:",omitempty"`
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

// SetNodeIP sets the NodeIP field's value.
func (s *NodeDetailInfoForDescribeDBInstancesOutput) SetNodeIP(v string) *NodeDetailInfoForDescribeDBInstancesOutput {
	s.NodeIP = &v
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

type TagFilterForDescribeDBInstancesInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Key *string `type:"string" json:",omitempty"`

	Value *string `type:"string" json:",omitempty"`

	Values []*string `type:"list" json:",omitempty"`
}

// String returns the string representation
func (s TagFilterForDescribeDBInstancesInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s TagFilterForDescribeDBInstancesInput) GoString() string {
	return s.String()
}

// SetKey sets the Key field's value.
func (s *TagFilterForDescribeDBInstancesInput) SetKey(v string) *TagFilterForDescribeDBInstancesInput {
	s.Key = &v
	return s
}

// SetValue sets the Value field's value.
func (s *TagFilterForDescribeDBInstancesInput) SetValue(v string) *TagFilterForDescribeDBInstancesInput {
	s.Value = &v
	return s
}

// SetValues sets the Values field's value.
func (s *TagFilterForDescribeDBInstancesInput) SetValues(v []*string) *TagFilterForDescribeDBInstancesInput {
	s.Values = v
	return s
}

type TagForDescribeDBInstancesOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Key *string `type:"string" json:",omitempty"`

	Value *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s TagForDescribeDBInstancesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s TagForDescribeDBInstancesOutput) GoString() string {
	return s.String()
}

// SetKey sets the Key field's value.
func (s *TagForDescribeDBInstancesOutput) SetKey(v string) *TagForDescribeDBInstancesOutput {
	s.Key = &v
	return s
}

// SetValue sets the Value field's value.
func (s *TagForDescribeDBInstancesOutput) SetValue(v string) *TagForDescribeDBInstancesOutput {
	s.Value = &v
	return s
}

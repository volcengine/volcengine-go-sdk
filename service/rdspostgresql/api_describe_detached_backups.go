// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package rdspostgresql

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribeDetachedBackupsCommon = "DescribeDetachedBackups"

// DescribeDetachedBackupsCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeDetachedBackupsCommon operation. The "output" return
// value will be populated with the DescribeDetachedBackupsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeDetachedBackupsCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeDetachedBackupsCommon Send returns without error.
//
// See DescribeDetachedBackupsCommon for more information on using the DescribeDetachedBackupsCommon
// API call, and error handling.
//
//    // Example sending a request using the DescribeDetachedBackupsCommonRequest method.
//    req, resp := client.DescribeDetachedBackupsCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *RDSPOSTGRESQL) DescribeDetachedBackupsCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeDetachedBackupsCommon,
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

// DescribeDetachedBackupsCommon API operation for RDS_POSTGRESQL.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RDS_POSTGRESQL's
// API operation DescribeDetachedBackupsCommon for usage and error information.
func (c *RDSPOSTGRESQL) DescribeDetachedBackupsCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeDetachedBackupsCommonRequest(input)
	return out, req.Send()
}

// DescribeDetachedBackupsCommonWithContext is the same as DescribeDetachedBackupsCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeDetachedBackupsCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RDSPOSTGRESQL) DescribeDetachedBackupsCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeDetachedBackupsCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeDetachedBackups = "DescribeDetachedBackups"

// DescribeDetachedBackupsRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeDetachedBackups operation. The "output" return
// value will be populated with the DescribeDetachedBackupsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeDetachedBackupsCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeDetachedBackupsCommon Send returns without error.
//
// See DescribeDetachedBackups for more information on using the DescribeDetachedBackups
// API call, and error handling.
//
//    // Example sending a request using the DescribeDetachedBackupsRequest method.
//    req, resp := client.DescribeDetachedBackupsRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *RDSPOSTGRESQL) DescribeDetachedBackupsRequest(input *DescribeDetachedBackupsInput) (req *request.Request, output *DescribeDetachedBackupsOutput) {
	op := &request.Operation{
		Name:       opDescribeDetachedBackups,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeDetachedBackupsInput{}
	}

	output = &DescribeDetachedBackupsOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// DescribeDetachedBackups API operation for RDS_POSTGRESQL.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RDS_POSTGRESQL's
// API operation DescribeDetachedBackups for usage and error information.
func (c *RDSPOSTGRESQL) DescribeDetachedBackups(input *DescribeDetachedBackupsInput) (*DescribeDetachedBackupsOutput, error) {
	req, out := c.DescribeDetachedBackupsRequest(input)
	return out, req.Send()
}

// DescribeDetachedBackupsWithContext is the same as DescribeDetachedBackups with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeDetachedBackups for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RDSPOSTGRESQL) DescribeDetachedBackupsWithContext(ctx volcengine.Context, input *DescribeDetachedBackupsInput, opts ...request.Option) (*DescribeDetachedBackupsOutput, error) {
	req, out := c.DescribeDetachedBackupsRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type BackupForDescribeDetachedBackupsOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	BackupEndTime *string `type:"string" json:",omitempty"`

	BackupFileName *string `type:"string" json:",omitempty"`

	BackupFileSize *int64 `type:"int64" json:",omitempty"`

	BackupId *string `type:"string" json:",omitempty"`

	BackupProgress *int32 `type:"int32" json:",omitempty"`

	BackupStartTime *string `type:"string" json:",omitempty"`

	BackupStatus *string `type:"string" json:",omitempty"`

	BackupType *string `type:"string" json:",omitempty"`

	CreateType *string `type:"string" json:",omitempty"`

	InstanceInfo *InstanceInfoForDescribeDetachedBackupsOutput `type:"structure" json:",omitempty"`
}

// String returns the string representation
func (s BackupForDescribeDetachedBackupsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s BackupForDescribeDetachedBackupsOutput) GoString() string {
	return s.String()
}

// SetBackupEndTime sets the BackupEndTime field's value.
func (s *BackupForDescribeDetachedBackupsOutput) SetBackupEndTime(v string) *BackupForDescribeDetachedBackupsOutput {
	s.BackupEndTime = &v
	return s
}

// SetBackupFileName sets the BackupFileName field's value.
func (s *BackupForDescribeDetachedBackupsOutput) SetBackupFileName(v string) *BackupForDescribeDetachedBackupsOutput {
	s.BackupFileName = &v
	return s
}

// SetBackupFileSize sets the BackupFileSize field's value.
func (s *BackupForDescribeDetachedBackupsOutput) SetBackupFileSize(v int64) *BackupForDescribeDetachedBackupsOutput {
	s.BackupFileSize = &v
	return s
}

// SetBackupId sets the BackupId field's value.
func (s *BackupForDescribeDetachedBackupsOutput) SetBackupId(v string) *BackupForDescribeDetachedBackupsOutput {
	s.BackupId = &v
	return s
}

// SetBackupProgress sets the BackupProgress field's value.
func (s *BackupForDescribeDetachedBackupsOutput) SetBackupProgress(v int32) *BackupForDescribeDetachedBackupsOutput {
	s.BackupProgress = &v
	return s
}

// SetBackupStartTime sets the BackupStartTime field's value.
func (s *BackupForDescribeDetachedBackupsOutput) SetBackupStartTime(v string) *BackupForDescribeDetachedBackupsOutput {
	s.BackupStartTime = &v
	return s
}

// SetBackupStatus sets the BackupStatus field's value.
func (s *BackupForDescribeDetachedBackupsOutput) SetBackupStatus(v string) *BackupForDescribeDetachedBackupsOutput {
	s.BackupStatus = &v
	return s
}

// SetBackupType sets the BackupType field's value.
func (s *BackupForDescribeDetachedBackupsOutput) SetBackupType(v string) *BackupForDescribeDetachedBackupsOutput {
	s.BackupType = &v
	return s
}

// SetCreateType sets the CreateType field's value.
func (s *BackupForDescribeDetachedBackupsOutput) SetCreateType(v string) *BackupForDescribeDetachedBackupsOutput {
	s.CreateType = &v
	return s
}

// SetInstanceInfo sets the InstanceInfo field's value.
func (s *BackupForDescribeDetachedBackupsOutput) SetInstanceInfo(v *InstanceInfoForDescribeDetachedBackupsOutput) *BackupForDescribeDetachedBackupsOutput {
	s.InstanceInfo = v
	return s
}

type ChargeDetailForDescribeDetachedBackupsOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	AutoRenew *bool `type:"boolean" json:",omitempty"`

	ChargeEndTime *string `type:"string" json:",omitempty"`

	ChargeStartTime *string `type:"string" json:",omitempty"`

	ChargeStatus *string `type:"string" json:",omitempty"`

	ChargeType *string `type:"string" json:",omitempty"`

	Number *int32 `type:"int32" json:",omitempty"`

	OverdueReclaimTime *string `type:"string" json:",omitempty"`

	OverdueTime *string `type:"string" json:",omitempty"`

	Period *int32 `type:"int32" json:",omitempty"`

	PeriodUnit *string `type:"string" json:",omitempty"`

	TempModifyEndTime *string `type:"string" json:",omitempty"`

	TempModifyStartTime *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s ChargeDetailForDescribeDetachedBackupsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ChargeDetailForDescribeDetachedBackupsOutput) GoString() string {
	return s.String()
}

// SetAutoRenew sets the AutoRenew field's value.
func (s *ChargeDetailForDescribeDetachedBackupsOutput) SetAutoRenew(v bool) *ChargeDetailForDescribeDetachedBackupsOutput {
	s.AutoRenew = &v
	return s
}

// SetChargeEndTime sets the ChargeEndTime field's value.
func (s *ChargeDetailForDescribeDetachedBackupsOutput) SetChargeEndTime(v string) *ChargeDetailForDescribeDetachedBackupsOutput {
	s.ChargeEndTime = &v
	return s
}

// SetChargeStartTime sets the ChargeStartTime field's value.
func (s *ChargeDetailForDescribeDetachedBackupsOutput) SetChargeStartTime(v string) *ChargeDetailForDescribeDetachedBackupsOutput {
	s.ChargeStartTime = &v
	return s
}

// SetChargeStatus sets the ChargeStatus field's value.
func (s *ChargeDetailForDescribeDetachedBackupsOutput) SetChargeStatus(v string) *ChargeDetailForDescribeDetachedBackupsOutput {
	s.ChargeStatus = &v
	return s
}

// SetChargeType sets the ChargeType field's value.
func (s *ChargeDetailForDescribeDetachedBackupsOutput) SetChargeType(v string) *ChargeDetailForDescribeDetachedBackupsOutput {
	s.ChargeType = &v
	return s
}

// SetNumber sets the Number field's value.
func (s *ChargeDetailForDescribeDetachedBackupsOutput) SetNumber(v int32) *ChargeDetailForDescribeDetachedBackupsOutput {
	s.Number = &v
	return s
}

// SetOverdueReclaimTime sets the OverdueReclaimTime field's value.
func (s *ChargeDetailForDescribeDetachedBackupsOutput) SetOverdueReclaimTime(v string) *ChargeDetailForDescribeDetachedBackupsOutput {
	s.OverdueReclaimTime = &v
	return s
}

// SetOverdueTime sets the OverdueTime field's value.
func (s *ChargeDetailForDescribeDetachedBackupsOutput) SetOverdueTime(v string) *ChargeDetailForDescribeDetachedBackupsOutput {
	s.OverdueTime = &v
	return s
}

// SetPeriod sets the Period field's value.
func (s *ChargeDetailForDescribeDetachedBackupsOutput) SetPeriod(v int32) *ChargeDetailForDescribeDetachedBackupsOutput {
	s.Period = &v
	return s
}

// SetPeriodUnit sets the PeriodUnit field's value.
func (s *ChargeDetailForDescribeDetachedBackupsOutput) SetPeriodUnit(v string) *ChargeDetailForDescribeDetachedBackupsOutput {
	s.PeriodUnit = &v
	return s
}

// SetTempModifyEndTime sets the TempModifyEndTime field's value.
func (s *ChargeDetailForDescribeDetachedBackupsOutput) SetTempModifyEndTime(v string) *ChargeDetailForDescribeDetachedBackupsOutput {
	s.TempModifyEndTime = &v
	return s
}

// SetTempModifyStartTime sets the TempModifyStartTime field's value.
func (s *ChargeDetailForDescribeDetachedBackupsOutput) SetTempModifyStartTime(v string) *ChargeDetailForDescribeDetachedBackupsOutput {
	s.TempModifyStartTime = &v
	return s
}

type DescribeDetachedBackupsInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	BackupEndTime *string `type:"string" json:",omitempty"`

	BackupId *string `type:"string" json:",omitempty"`

	BackupStartTime *string `type:"string" json:",omitempty"`

	BackupStatus *string `type:"string" json:",omitempty"`

	BackupType *string `type:"string" json:",omitempty"`

	InstanceId *string `type:"string" json:",omitempty"`

	InstanceName *string `type:"string" json:",omitempty"`

	PageNumber *int32 `type:"int32" json:",omitempty"`

	PageSize *int32 `type:"int32" json:",omitempty"`

	ProjectName *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s DescribeDetachedBackupsInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeDetachedBackupsInput) GoString() string {
	return s.String()
}

// SetBackupEndTime sets the BackupEndTime field's value.
func (s *DescribeDetachedBackupsInput) SetBackupEndTime(v string) *DescribeDetachedBackupsInput {
	s.BackupEndTime = &v
	return s
}

// SetBackupId sets the BackupId field's value.
func (s *DescribeDetachedBackupsInput) SetBackupId(v string) *DescribeDetachedBackupsInput {
	s.BackupId = &v
	return s
}

// SetBackupStartTime sets the BackupStartTime field's value.
func (s *DescribeDetachedBackupsInput) SetBackupStartTime(v string) *DescribeDetachedBackupsInput {
	s.BackupStartTime = &v
	return s
}

// SetBackupStatus sets the BackupStatus field's value.
func (s *DescribeDetachedBackupsInput) SetBackupStatus(v string) *DescribeDetachedBackupsInput {
	s.BackupStatus = &v
	return s
}

// SetBackupType sets the BackupType field's value.
func (s *DescribeDetachedBackupsInput) SetBackupType(v string) *DescribeDetachedBackupsInput {
	s.BackupType = &v
	return s
}

// SetInstanceId sets the InstanceId field's value.
func (s *DescribeDetachedBackupsInput) SetInstanceId(v string) *DescribeDetachedBackupsInput {
	s.InstanceId = &v
	return s
}

// SetInstanceName sets the InstanceName field's value.
func (s *DescribeDetachedBackupsInput) SetInstanceName(v string) *DescribeDetachedBackupsInput {
	s.InstanceName = &v
	return s
}

// SetPageNumber sets the PageNumber field's value.
func (s *DescribeDetachedBackupsInput) SetPageNumber(v int32) *DescribeDetachedBackupsInput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *DescribeDetachedBackupsInput) SetPageSize(v int32) *DescribeDetachedBackupsInput {
	s.PageSize = &v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *DescribeDetachedBackupsInput) SetProjectName(v string) *DescribeDetachedBackupsInput {
	s.ProjectName = &v
	return s
}

type DescribeDetachedBackupsOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	Backups []*BackupForDescribeDetachedBackupsOutput `type:"list" json:",omitempty"`

	Total *int32 `type:"int32" json:",omitempty"`
}

// String returns the string representation
func (s DescribeDetachedBackupsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeDetachedBackupsOutput) GoString() string {
	return s.String()
}

// SetBackups sets the Backups field's value.
func (s *DescribeDetachedBackupsOutput) SetBackups(v []*BackupForDescribeDetachedBackupsOutput) *DescribeDetachedBackupsOutput {
	s.Backups = v
	return s
}

// SetTotal sets the Total field's value.
func (s *DescribeDetachedBackupsOutput) SetTotal(v int32) *DescribeDetachedBackupsOutput {
	s.Total = &v
	return s
}

type InstanceInfoForDescribeDetachedBackupsOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	ChargeDetail *ChargeDetailForDescribeDetachedBackupsOutput `type:"structure" json:",omitempty"`

	DBEngineVersion *string `type:"string" json:",omitempty"`

	InstanceId *string `type:"string" json:",omitempty"`

	InstanceName *string `type:"string" json:",omitempty"`

	InstanceStatus *string `type:"string" json:",omitempty"`

	InstanceType *string `type:"string" json:",omitempty"`

	Memory *int32 `type:"int32" json:",omitempty"`

	Nodes []*NodeForDescribeDetachedBackupsOutput `type:"list" json:",omitempty"`

	ProjectName *string `type:"string" json:",omitempty"`

	StorageSpace *int32 `type:"int32" json:",omitempty"`

	StorageType *string `type:"string" json:",omitempty"`

	SubnetId *string `type:"string" json:",omitempty"`

	VCPU *int32 `type:"int32" json:",omitempty"`

	VpcId *string `type:"string" json:",omitempty"`

	ZoneId *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s InstanceInfoForDescribeDetachedBackupsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s InstanceInfoForDescribeDetachedBackupsOutput) GoString() string {
	return s.String()
}

// SetChargeDetail sets the ChargeDetail field's value.
func (s *InstanceInfoForDescribeDetachedBackupsOutput) SetChargeDetail(v *ChargeDetailForDescribeDetachedBackupsOutput) *InstanceInfoForDescribeDetachedBackupsOutput {
	s.ChargeDetail = v
	return s
}

// SetDBEngineVersion sets the DBEngineVersion field's value.
func (s *InstanceInfoForDescribeDetachedBackupsOutput) SetDBEngineVersion(v string) *InstanceInfoForDescribeDetachedBackupsOutput {
	s.DBEngineVersion = &v
	return s
}

// SetInstanceId sets the InstanceId field's value.
func (s *InstanceInfoForDescribeDetachedBackupsOutput) SetInstanceId(v string) *InstanceInfoForDescribeDetachedBackupsOutput {
	s.InstanceId = &v
	return s
}

// SetInstanceName sets the InstanceName field's value.
func (s *InstanceInfoForDescribeDetachedBackupsOutput) SetInstanceName(v string) *InstanceInfoForDescribeDetachedBackupsOutput {
	s.InstanceName = &v
	return s
}

// SetInstanceStatus sets the InstanceStatus field's value.
func (s *InstanceInfoForDescribeDetachedBackupsOutput) SetInstanceStatus(v string) *InstanceInfoForDescribeDetachedBackupsOutput {
	s.InstanceStatus = &v
	return s
}

// SetInstanceType sets the InstanceType field's value.
func (s *InstanceInfoForDescribeDetachedBackupsOutput) SetInstanceType(v string) *InstanceInfoForDescribeDetachedBackupsOutput {
	s.InstanceType = &v
	return s
}

// SetMemory sets the Memory field's value.
func (s *InstanceInfoForDescribeDetachedBackupsOutput) SetMemory(v int32) *InstanceInfoForDescribeDetachedBackupsOutput {
	s.Memory = &v
	return s
}

// SetNodes sets the Nodes field's value.
func (s *InstanceInfoForDescribeDetachedBackupsOutput) SetNodes(v []*NodeForDescribeDetachedBackupsOutput) *InstanceInfoForDescribeDetachedBackupsOutput {
	s.Nodes = v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *InstanceInfoForDescribeDetachedBackupsOutput) SetProjectName(v string) *InstanceInfoForDescribeDetachedBackupsOutput {
	s.ProjectName = &v
	return s
}

// SetStorageSpace sets the StorageSpace field's value.
func (s *InstanceInfoForDescribeDetachedBackupsOutput) SetStorageSpace(v int32) *InstanceInfoForDescribeDetachedBackupsOutput {
	s.StorageSpace = &v
	return s
}

// SetStorageType sets the StorageType field's value.
func (s *InstanceInfoForDescribeDetachedBackupsOutput) SetStorageType(v string) *InstanceInfoForDescribeDetachedBackupsOutput {
	s.StorageType = &v
	return s
}

// SetSubnetId sets the SubnetId field's value.
func (s *InstanceInfoForDescribeDetachedBackupsOutput) SetSubnetId(v string) *InstanceInfoForDescribeDetachedBackupsOutput {
	s.SubnetId = &v
	return s
}

// SetVCPU sets the VCPU field's value.
func (s *InstanceInfoForDescribeDetachedBackupsOutput) SetVCPU(v int32) *InstanceInfoForDescribeDetachedBackupsOutput {
	s.VCPU = &v
	return s
}

// SetVpcId sets the VpcId field's value.
func (s *InstanceInfoForDescribeDetachedBackupsOutput) SetVpcId(v string) *InstanceInfoForDescribeDetachedBackupsOutput {
	s.VpcId = &v
	return s
}

// SetZoneId sets the ZoneId field's value.
func (s *InstanceInfoForDescribeDetachedBackupsOutput) SetZoneId(v string) *InstanceInfoForDescribeDetachedBackupsOutput {
	s.ZoneId = &v
	return s
}

type NodeForDescribeDetachedBackupsOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	CreateTime *string `type:"string" json:",omitempty"`

	InstanceId *string `type:"string" json:",omitempty"`

	Memory *int32 `type:"int32" json:",omitempty"`

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
func (s NodeForDescribeDetachedBackupsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s NodeForDescribeDetachedBackupsOutput) GoString() string {
	return s.String()
}

// SetCreateTime sets the CreateTime field's value.
func (s *NodeForDescribeDetachedBackupsOutput) SetCreateTime(v string) *NodeForDescribeDetachedBackupsOutput {
	s.CreateTime = &v
	return s
}

// SetInstanceId sets the InstanceId field's value.
func (s *NodeForDescribeDetachedBackupsOutput) SetInstanceId(v string) *NodeForDescribeDetachedBackupsOutput {
	s.InstanceId = &v
	return s
}

// SetMemory sets the Memory field's value.
func (s *NodeForDescribeDetachedBackupsOutput) SetMemory(v int32) *NodeForDescribeDetachedBackupsOutput {
	s.Memory = &v
	return s
}

// SetNodeId sets the NodeId field's value.
func (s *NodeForDescribeDetachedBackupsOutput) SetNodeId(v string) *NodeForDescribeDetachedBackupsOutput {
	s.NodeId = &v
	return s
}

// SetNodeSpec sets the NodeSpec field's value.
func (s *NodeForDescribeDetachedBackupsOutput) SetNodeSpec(v string) *NodeForDescribeDetachedBackupsOutput {
	s.NodeSpec = &v
	return s
}

// SetNodeStatus sets the NodeStatus field's value.
func (s *NodeForDescribeDetachedBackupsOutput) SetNodeStatus(v string) *NodeForDescribeDetachedBackupsOutput {
	s.NodeStatus = &v
	return s
}

// SetNodeType sets the NodeType field's value.
func (s *NodeForDescribeDetachedBackupsOutput) SetNodeType(v string) *NodeForDescribeDetachedBackupsOutput {
	s.NodeType = &v
	return s
}

// SetRegionId sets the RegionId field's value.
func (s *NodeForDescribeDetachedBackupsOutput) SetRegionId(v string) *NodeForDescribeDetachedBackupsOutput {
	s.RegionId = &v
	return s
}

// SetUpdateTime sets the UpdateTime field's value.
func (s *NodeForDescribeDetachedBackupsOutput) SetUpdateTime(v string) *NodeForDescribeDetachedBackupsOutput {
	s.UpdateTime = &v
	return s
}

// SetVCPU sets the VCPU field's value.
func (s *NodeForDescribeDetachedBackupsOutput) SetVCPU(v int32) *NodeForDescribeDetachedBackupsOutput {
	s.VCPU = &v
	return s
}

// SetZoneId sets the ZoneId field's value.
func (s *NodeForDescribeDetachedBackupsOutput) SetZoneId(v string) *NodeForDescribeDetachedBackupsOutput {
	s.ZoneId = &v
	return s
}

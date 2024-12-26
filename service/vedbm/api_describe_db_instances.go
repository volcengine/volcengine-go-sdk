// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package vedbm

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
func (c *VEDBM) DescribeDBInstancesCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
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

// DescribeDBInstancesCommon API operation for VEDBM.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VEDBM's
// API operation DescribeDBInstancesCommon for usage and error information.
func (c *VEDBM) DescribeDBInstancesCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
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
func (c *VEDBM) DescribeDBInstancesCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
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
func (c *VEDBM) DescribeDBInstancesRequest(input *DescribeDBInstancesInput) (req *request.Request, output *DescribeDBInstancesOutput) {
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

// DescribeDBInstances API operation for VEDBM.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VEDBM's
// API operation DescribeDBInstances for usage and error information.
func (c *VEDBM) DescribeDBInstances(input *DescribeDBInstancesInput) (*DescribeDBInstancesOutput, error) {
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
func (c *VEDBM) DescribeDBInstancesWithContext(ctx volcengine.Context, input *DescribeDBInstancesInput, opts ...request.Option) (*DescribeDBInstancesOutput, error) {
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

	ChargeStatus *string `type:"string" json:",omitempty" enum:"EnumOfChargeStatusForDescribeDBInstancesOutput"`

	ChargeType *string `type:"string" json:",omitempty" enum:"EnumOfChargeTypeForDescribeDBInstancesOutput"`

	OverdueReclaimTime *string `type:"string" json:",omitempty"`

	OverdueTime *string `type:"string" json:",omitempty"`
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

type DescribeDBInstancesInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	ChargeType *string `type:"string" json:",omitempty" enum:"EnumOfChargeTypeForDescribeDBInstancesInput"`

	CreateTimeEnd *string `type:"string" json:",omitempty"`

	CreateTimeStart *string `type:"string" json:",omitempty"`

	DBEngineVersion *string `type:"string" json:",omitempty" enum:"EnumOfDBEngineVersionForDescribeDBInstancesInput"`

	InstanceId *string `type:"string" json:",omitempty"`

	InstanceName *string `type:"string" json:",omitempty"`

	InstanceStatus *string `type:"string" json:",omitempty" enum:"EnumOfInstanceStatusForDescribeDBInstancesInput"`

	NodeSpec *string `type:"string" json:",omitempty" enum:"EnumOfNodeSpecForDescribeDBInstancesInput"`

	PageNumber *int32 `type:"int32" json:",omitempty"`

	PageSize *int32 `type:"int32" json:",omitempty"`

	ProjectName *string `type:"string" json:",omitempty"`

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

// SetNodeSpec sets the NodeSpec field's value.
func (s *DescribeDBInstancesInput) SetNodeSpec(v string) *DescribeDBInstancesInput {
	s.NodeSpec = &v
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

// SetProjectName sets the ProjectName field's value.
func (s *DescribeDBInstancesInput) SetProjectName(v string) *DescribeDBInstancesInput {
	s.ProjectName = &v
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

	Instances []*InstanceForDescribeDBInstancesOutput `type:"list" json:",omitempty"`

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
	_ struct{} `type:"structure" json:",omitempty"`

	ChargeDetail *ChargeDetailForDescribeDBInstancesOutput `type:"structure" json:",omitempty"`

	CreateTime *string `type:"string" json:",omitempty"`

	DBEngineVersion *string `type:"string" json:",omitempty" enum:"EnumOfDBEngineVersionForDescribeDBInstancesOutput"`

	DBRevisionVersion *string `type:"string" json:",omitempty"`

	DeletionProtection *string `type:"string" json:",omitempty" enum:"EnumOfDeletionProtectionForDescribeDBInstancesOutput"`

	InstanceId *string `type:"string" json:",omitempty"`

	InstanceName *string `type:"string" json:",omitempty"`

	InstanceStatus *string `type:"string" json:",omitempty" enum:"EnumOfInstanceStatusForDescribeDBInstancesOutput"`

	LowerCaseTableNames *string `type:"string" json:",omitempty"`

	Nodes []*NodeForDescribeDBInstancesOutput `type:"list" json:",omitempty"`

	PrePaidStorageInGB *int32 `type:"int32" json:",omitempty"`

	ProjectName *string `type:"string" json:",omitempty"`

	RegionId *string `type:"string" json:",omitempty"`

	SpecFamily *string `type:"string" json:",omitempty"`

	StorageChargeType *string `type:"string" json:",omitempty" enum:"EnumOfStorageChargeTypeForDescribeDBInstancesOutput"`

	StorageUsedGiB *float64 `type:"double" json:",omitempty"`

	SubnetId *string `type:"string" json:",omitempty"`

	Tags []*TagForDescribeDBInstancesOutput `type:"list" json:",omitempty"`

	TimeZone *string `type:"string" json:",omitempty"`

	VpcId *string `type:"string" json:",omitempty"`

	ZoneIds *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s InstanceForDescribeDBInstancesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s InstanceForDescribeDBInstancesOutput) GoString() string {
	return s.String()
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

// SetDBRevisionVersion sets the DBRevisionVersion field's value.
func (s *InstanceForDescribeDBInstancesOutput) SetDBRevisionVersion(v string) *InstanceForDescribeDBInstancesOutput {
	s.DBRevisionVersion = &v
	return s
}

// SetDeletionProtection sets the DeletionProtection field's value.
func (s *InstanceForDescribeDBInstancesOutput) SetDeletionProtection(v string) *InstanceForDescribeDBInstancesOutput {
	s.DeletionProtection = &v
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

// SetLowerCaseTableNames sets the LowerCaseTableNames field's value.
func (s *InstanceForDescribeDBInstancesOutput) SetLowerCaseTableNames(v string) *InstanceForDescribeDBInstancesOutput {
	s.LowerCaseTableNames = &v
	return s
}

// SetNodes sets the Nodes field's value.
func (s *InstanceForDescribeDBInstancesOutput) SetNodes(v []*NodeForDescribeDBInstancesOutput) *InstanceForDescribeDBInstancesOutput {
	s.Nodes = v
	return s
}

// SetPrePaidStorageInGB sets the PrePaidStorageInGB field's value.
func (s *InstanceForDescribeDBInstancesOutput) SetPrePaidStorageInGB(v int32) *InstanceForDescribeDBInstancesOutput {
	s.PrePaidStorageInGB = &v
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

// SetSpecFamily sets the SpecFamily field's value.
func (s *InstanceForDescribeDBInstancesOutput) SetSpecFamily(v string) *InstanceForDescribeDBInstancesOutput {
	s.SpecFamily = &v
	return s
}

// SetStorageChargeType sets the StorageChargeType field's value.
func (s *InstanceForDescribeDBInstancesOutput) SetStorageChargeType(v string) *InstanceForDescribeDBInstancesOutput {
	s.StorageChargeType = &v
	return s
}

// SetStorageUsedGiB sets the StorageUsedGiB field's value.
func (s *InstanceForDescribeDBInstancesOutput) SetStorageUsedGiB(v float64) *InstanceForDescribeDBInstancesOutput {
	s.StorageUsedGiB = &v
	return s
}

// SetSubnetId sets the SubnetId field's value.
func (s *InstanceForDescribeDBInstancesOutput) SetSubnetId(v string) *InstanceForDescribeDBInstancesOutput {
	s.SubnetId = &v
	return s
}

// SetTags sets the Tags field's value.
func (s *InstanceForDescribeDBInstancesOutput) SetTags(v []*TagForDescribeDBInstancesOutput) *InstanceForDescribeDBInstancesOutput {
	s.Tags = v
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

// SetZoneIds sets the ZoneIds field's value.
func (s *InstanceForDescribeDBInstancesOutput) SetZoneIds(v string) *InstanceForDescribeDBInstancesOutput {
	s.ZoneIds = &v
	return s
}

type NodeForDescribeDBInstancesOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Memory *int32 `type:"int32" json:",omitempty"`

	NodeId *string `type:"string" json:",omitempty"`

	NodeSpec *string `type:"string" json:",omitempty" enum:"EnumOfNodeSpecForDescribeDBInstancesOutput"`

	NodeType *string `type:"string" json:",omitempty" enum:"EnumOfNodeTypeForDescribeDBInstancesOutput"`

	VCPU *int32 `type:"int32" json:"vCPU,omitempty"`

	ZoneId *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s NodeForDescribeDBInstancesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s NodeForDescribeDBInstancesOutput) GoString() string {
	return s.String()
}

// SetMemory sets the Memory field's value.
func (s *NodeForDescribeDBInstancesOutput) SetMemory(v int32) *NodeForDescribeDBInstancesOutput {
	s.Memory = &v
	return s
}

// SetNodeId sets the NodeId field's value.
func (s *NodeForDescribeDBInstancesOutput) SetNodeId(v string) *NodeForDescribeDBInstancesOutput {
	s.NodeId = &v
	return s
}

// SetNodeSpec sets the NodeSpec field's value.
func (s *NodeForDescribeDBInstancesOutput) SetNodeSpec(v string) *NodeForDescribeDBInstancesOutput {
	s.NodeSpec = &v
	return s
}

// SetNodeType sets the NodeType field's value.
func (s *NodeForDescribeDBInstancesOutput) SetNodeType(v string) *NodeForDescribeDBInstancesOutput {
	s.NodeType = &v
	return s
}

// SetVCPU sets the VCPU field's value.
func (s *NodeForDescribeDBInstancesOutput) SetVCPU(v int32) *NodeForDescribeDBInstancesOutput {
	s.VCPU = &v
	return s
}

// SetZoneId sets the ZoneId field's value.
func (s *NodeForDescribeDBInstancesOutput) SetZoneId(v string) *NodeForDescribeDBInstancesOutput {
	s.ZoneId = &v
	return s
}

type TagFilterForDescribeDBInstancesInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Key *string `type:"string" json:",omitempty"`

	Value *string `type:"string" json:",omitempty"`
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

const (
	// EnumOfChargeStatusForDescribeDBInstancesOutputNormal is a EnumOfChargeStatusForDescribeDBInstancesOutput enum value
	EnumOfChargeStatusForDescribeDBInstancesOutputNormal = "Normal"

	// EnumOfChargeStatusForDescribeDBInstancesOutputOverdue is a EnumOfChargeStatusForDescribeDBInstancesOutput enum value
	EnumOfChargeStatusForDescribeDBInstancesOutputOverdue = "Overdue"

	// EnumOfChargeStatusForDescribeDBInstancesOutputShutdown is a EnumOfChargeStatusForDescribeDBInstancesOutput enum value
	EnumOfChargeStatusForDescribeDBInstancesOutputShutdown = "Shutdown"
)

const (
	// EnumOfChargeTypeForDescribeDBInstancesInputPostPaid is a EnumOfChargeTypeForDescribeDBInstancesInput enum value
	EnumOfChargeTypeForDescribeDBInstancesInputPostPaid = "PostPaid"

	// EnumOfChargeTypeForDescribeDBInstancesInputPrePaid is a EnumOfChargeTypeForDescribeDBInstancesInput enum value
	EnumOfChargeTypeForDescribeDBInstancesInputPrePaid = "PrePaid"
)

const (
	// EnumOfChargeTypeForDescribeDBInstancesOutputPostPaid is a EnumOfChargeTypeForDescribeDBInstancesOutput enum value
	EnumOfChargeTypeForDescribeDBInstancesOutputPostPaid = "PostPaid"

	// EnumOfChargeTypeForDescribeDBInstancesOutputPrePaid is a EnumOfChargeTypeForDescribeDBInstancesOutput enum value
	EnumOfChargeTypeForDescribeDBInstancesOutputPrePaid = "PrePaid"
)

const (
	// EnumOfDBEngineVersionForDescribeDBInstancesInputMySql80 is a EnumOfDBEngineVersionForDescribeDBInstancesInput enum value
	EnumOfDBEngineVersionForDescribeDBInstancesInputMySql80 = "MySQL_8_0"
)

const (
	// EnumOfDBEngineVersionForDescribeDBInstancesOutputMySql80 is a EnumOfDBEngineVersionForDescribeDBInstancesOutput enum value
	EnumOfDBEngineVersionForDescribeDBInstancesOutputMySql80 = "MySQL_8_0"

	// EnumOfDBEngineVersionForDescribeDBInstancesOutputMySql57 is a EnumOfDBEngineVersionForDescribeDBInstancesOutput enum value
	EnumOfDBEngineVersionForDescribeDBInstancesOutputMySql57 = "MySQL_5_7"
)

const (
	// EnumOfDeletionProtectionForDescribeDBInstancesOutputDisabled is a EnumOfDeletionProtectionForDescribeDBInstancesOutput enum value
	EnumOfDeletionProtectionForDescribeDBInstancesOutputDisabled = "disabled"

	// EnumOfDeletionProtectionForDescribeDBInstancesOutputEnabled is a EnumOfDeletionProtectionForDescribeDBInstancesOutput enum value
	EnumOfDeletionProtectionForDescribeDBInstancesOutputEnabled = "enabled"
)

const (
	// EnumOfInstanceStatusForDescribeDBInstancesInputWaitingPaid is a EnumOfInstanceStatusForDescribeDBInstancesInput enum value
	EnumOfInstanceStatusForDescribeDBInstancesInputWaitingPaid = "WaitingPaid"

	// EnumOfInstanceStatusForDescribeDBInstancesInputRunning is a EnumOfInstanceStatusForDescribeDBInstancesInput enum value
	EnumOfInstanceStatusForDescribeDBInstancesInputRunning = "Running"

	// EnumOfInstanceStatusForDescribeDBInstancesInputCreating is a EnumOfInstanceStatusForDescribeDBInstancesInput enum value
	EnumOfInstanceStatusForDescribeDBInstancesInputCreating = "Creating"

	// EnumOfInstanceStatusForDescribeDBInstancesInputScaling is a EnumOfInstanceStatusForDescribeDBInstancesInput enum value
	EnumOfInstanceStatusForDescribeDBInstancesInputScaling = "Scaling"

	// EnumOfInstanceStatusForDescribeDBInstancesInputRestarting is a EnumOfInstanceStatusForDescribeDBInstancesInput enum value
	EnumOfInstanceStatusForDescribeDBInstancesInputRestarting = "Restarting"

	// EnumOfInstanceStatusForDescribeDBInstancesInputRestoring is a EnumOfInstanceStatusForDescribeDBInstancesInput enum value
	EnumOfInstanceStatusForDescribeDBInstancesInputRestoring = "Restoring"

	// EnumOfInstanceStatusForDescribeDBInstancesInputUpgrading is a EnumOfInstanceStatusForDescribeDBInstancesInput enum value
	EnumOfInstanceStatusForDescribeDBInstancesInputUpgrading = "Upgrading"

	// EnumOfInstanceStatusForDescribeDBInstancesInputPrimaryChanging is a EnumOfInstanceStatusForDescribeDBInstancesInput enum value
	EnumOfInstanceStatusForDescribeDBInstancesInputPrimaryChanging = "PrimaryChanging"

	// EnumOfInstanceStatusForDescribeDBInstancesInputUnavailable is a EnumOfInstanceStatusForDescribeDBInstancesInput enum value
	EnumOfInstanceStatusForDescribeDBInstancesInputUnavailable = "Unavailable"

	// EnumOfInstanceStatusForDescribeDBInstancesInputDeleting is a EnumOfInstanceStatusForDescribeDBInstancesInput enum value
	EnumOfInstanceStatusForDescribeDBInstancesInputDeleting = "Deleting"

	// EnumOfInstanceStatusForDescribeDBInstancesInputDeleted is a EnumOfInstanceStatusForDescribeDBInstancesInput enum value
	EnumOfInstanceStatusForDescribeDBInstancesInputDeleted = "Deleted"

	// EnumOfInstanceStatusForDescribeDBInstancesInputCreateFailed is a EnumOfInstanceStatusForDescribeDBInstancesInput enum value
	EnumOfInstanceStatusForDescribeDBInstancesInputCreateFailed = "CreateFailed"

	// EnumOfInstanceStatusForDescribeDBInstancesInputClosing is a EnumOfInstanceStatusForDescribeDBInstancesInput enum value
	EnumOfInstanceStatusForDescribeDBInstancesInputClosing = "Closing"

	// EnumOfInstanceStatusForDescribeDBInstancesInputExpired is a EnumOfInstanceStatusForDescribeDBInstancesInput enum value
	EnumOfInstanceStatusForDescribeDBInstancesInputExpired = "Expired"

	// EnumOfInstanceStatusForDescribeDBInstancesInputOwing is a EnumOfInstanceStatusForDescribeDBInstancesInput enum value
	EnumOfInstanceStatusForDescribeDBInstancesInputOwing = "Owing"

	// EnumOfInstanceStatusForDescribeDBInstancesInputResuming is a EnumOfInstanceStatusForDescribeDBInstancesInput enum value
	EnumOfInstanceStatusForDescribeDBInstancesInputResuming = "Resuming"

	// EnumOfInstanceStatusForDescribeDBInstancesInputAllowListMaintaining is a EnumOfInstanceStatusForDescribeDBInstancesInput enum value
	EnumOfInstanceStatusForDescribeDBInstancesInputAllowListMaintaining = "AllowListMaintaining"

	// EnumOfInstanceStatusForDescribeDBInstancesInputError is a EnumOfInstanceStatusForDescribeDBInstancesInput enum value
	EnumOfInstanceStatusForDescribeDBInstancesInputError = "Error"
)

const (
	// EnumOfInstanceStatusForDescribeDBInstancesOutputWaitingPaid is a EnumOfInstanceStatusForDescribeDBInstancesOutput enum value
	EnumOfInstanceStatusForDescribeDBInstancesOutputWaitingPaid = "WaitingPaid"

	// EnumOfInstanceStatusForDescribeDBInstancesOutputRunning is a EnumOfInstanceStatusForDescribeDBInstancesOutput enum value
	EnumOfInstanceStatusForDescribeDBInstancesOutputRunning = "Running"

	// EnumOfInstanceStatusForDescribeDBInstancesOutputCreating is a EnumOfInstanceStatusForDescribeDBInstancesOutput enum value
	EnumOfInstanceStatusForDescribeDBInstancesOutputCreating = "Creating"

	// EnumOfInstanceStatusForDescribeDBInstancesOutputScaling is a EnumOfInstanceStatusForDescribeDBInstancesOutput enum value
	EnumOfInstanceStatusForDescribeDBInstancesOutputScaling = "Scaling"

	// EnumOfInstanceStatusForDescribeDBInstancesOutputRestarting is a EnumOfInstanceStatusForDescribeDBInstancesOutput enum value
	EnumOfInstanceStatusForDescribeDBInstancesOutputRestarting = "Restarting"

	// EnumOfInstanceStatusForDescribeDBInstancesOutputRestoring is a EnumOfInstanceStatusForDescribeDBInstancesOutput enum value
	EnumOfInstanceStatusForDescribeDBInstancesOutputRestoring = "Restoring"

	// EnumOfInstanceStatusForDescribeDBInstancesOutputUpgrading is a EnumOfInstanceStatusForDescribeDBInstancesOutput enum value
	EnumOfInstanceStatusForDescribeDBInstancesOutputUpgrading = "Upgrading"

	// EnumOfInstanceStatusForDescribeDBInstancesOutputPrimaryChanging is a EnumOfInstanceStatusForDescribeDBInstancesOutput enum value
	EnumOfInstanceStatusForDescribeDBInstancesOutputPrimaryChanging = "PrimaryChanging"

	// EnumOfInstanceStatusForDescribeDBInstancesOutputUnavailable is a EnumOfInstanceStatusForDescribeDBInstancesOutput enum value
	EnumOfInstanceStatusForDescribeDBInstancesOutputUnavailable = "Unavailable"

	// EnumOfInstanceStatusForDescribeDBInstancesOutputDeleting is a EnumOfInstanceStatusForDescribeDBInstancesOutput enum value
	EnumOfInstanceStatusForDescribeDBInstancesOutputDeleting = "Deleting"

	// EnumOfInstanceStatusForDescribeDBInstancesOutputDeleted is a EnumOfInstanceStatusForDescribeDBInstancesOutput enum value
	EnumOfInstanceStatusForDescribeDBInstancesOutputDeleted = "Deleted"

	// EnumOfInstanceStatusForDescribeDBInstancesOutputCreateFailed is a EnumOfInstanceStatusForDescribeDBInstancesOutput enum value
	EnumOfInstanceStatusForDescribeDBInstancesOutputCreateFailed = "CreateFailed"

	// EnumOfInstanceStatusForDescribeDBInstancesOutputClosing is a EnumOfInstanceStatusForDescribeDBInstancesOutput enum value
	EnumOfInstanceStatusForDescribeDBInstancesOutputClosing = "Closing"

	// EnumOfInstanceStatusForDescribeDBInstancesOutputExpired is a EnumOfInstanceStatusForDescribeDBInstancesOutput enum value
	EnumOfInstanceStatusForDescribeDBInstancesOutputExpired = "Expired"

	// EnumOfInstanceStatusForDescribeDBInstancesOutputOwing is a EnumOfInstanceStatusForDescribeDBInstancesOutput enum value
	EnumOfInstanceStatusForDescribeDBInstancesOutputOwing = "Owing"

	// EnumOfInstanceStatusForDescribeDBInstancesOutputResuming is a EnumOfInstanceStatusForDescribeDBInstancesOutput enum value
	EnumOfInstanceStatusForDescribeDBInstancesOutputResuming = "Resuming"

	// EnumOfInstanceStatusForDescribeDBInstancesOutputAllowListMaintaining is a EnumOfInstanceStatusForDescribeDBInstancesOutput enum value
	EnumOfInstanceStatusForDescribeDBInstancesOutputAllowListMaintaining = "AllowListMaintaining"
)

const (
	// EnumOfNodeSpecForDescribeDBInstancesInputVedbMysqlX4Large is a EnumOfNodeSpecForDescribeDBInstancesInput enum value
	EnumOfNodeSpecForDescribeDBInstancesInputVedbMysqlX4Large = "vedb.mysql.x4.large"

	// EnumOfNodeSpecForDescribeDBInstancesInputVedbMysqlX8Large is a EnumOfNodeSpecForDescribeDBInstancesInput enum value
	EnumOfNodeSpecForDescribeDBInstancesInputVedbMysqlX8Large = "vedb.mysql.x8.large"

	// EnumOfNodeSpecForDescribeDBInstancesInputVedbMysqlX4Xlarge is a EnumOfNodeSpecForDescribeDBInstancesInput enum value
	EnumOfNodeSpecForDescribeDBInstancesInputVedbMysqlX4Xlarge = "vedb.mysql.x4.xlarge"

	// EnumOfNodeSpecForDescribeDBInstancesInputVedbMysqlX8Xlarge is a EnumOfNodeSpecForDescribeDBInstancesInput enum value
	EnumOfNodeSpecForDescribeDBInstancesInputVedbMysqlX8Xlarge = "vedb.mysql.x8.xlarge"

	// EnumOfNodeSpecForDescribeDBInstancesInputVedbMysqlX42xlarge is a EnumOfNodeSpecForDescribeDBInstancesInput enum value
	EnumOfNodeSpecForDescribeDBInstancesInputVedbMysqlX42xlarge = "vedb.mysql.x4.2xlarge"

	// EnumOfNodeSpecForDescribeDBInstancesInputVedbMysqlX82xlarge is a EnumOfNodeSpecForDescribeDBInstancesInput enum value
	EnumOfNodeSpecForDescribeDBInstancesInputVedbMysqlX82xlarge = "vedb.mysql.x8.2xlarge"

	// EnumOfNodeSpecForDescribeDBInstancesInputVedbMysqlX44xlarge is a EnumOfNodeSpecForDescribeDBInstancesInput enum value
	EnumOfNodeSpecForDescribeDBInstancesInputVedbMysqlX44xlarge = "vedb.mysql.x4.4xlarge"

	// EnumOfNodeSpecForDescribeDBInstancesInputVedbMysqlX84xlarge is a EnumOfNodeSpecForDescribeDBInstancesInput enum value
	EnumOfNodeSpecForDescribeDBInstancesInputVedbMysqlX84xlarge = "vedb.mysql.x8.4xlarge"

	// EnumOfNodeSpecForDescribeDBInstancesInputVedbMysqlX86xlarge is a EnumOfNodeSpecForDescribeDBInstancesInput enum value
	EnumOfNodeSpecForDescribeDBInstancesInputVedbMysqlX86xlarge = "vedb.mysql.x8.6xlarge"

	// EnumOfNodeSpecForDescribeDBInstancesInputVedbMysqlX48xlarge is a EnumOfNodeSpecForDescribeDBInstancesInput enum value
	EnumOfNodeSpecForDescribeDBInstancesInputVedbMysqlX48xlarge = "vedb.mysql.x4.8xlarge"

	// EnumOfNodeSpecForDescribeDBInstancesInputVedbMysqlX88xlarge is a EnumOfNodeSpecForDescribeDBInstancesInput enum value
	EnumOfNodeSpecForDescribeDBInstancesInputVedbMysqlX88xlarge = "vedb.mysql.x8.8xlarge"

	// EnumOfNodeSpecForDescribeDBInstancesInputVedbMysqlG4Large is a EnumOfNodeSpecForDescribeDBInstancesInput enum value
	EnumOfNodeSpecForDescribeDBInstancesInputVedbMysqlG4Large = "vedb.mysql.g4.large"

	// EnumOfNodeSpecForDescribeDBInstancesInputVedbMysqlG4Xlarge is a EnumOfNodeSpecForDescribeDBInstancesInput enum value
	EnumOfNodeSpecForDescribeDBInstancesInputVedbMysqlG4Xlarge = "vedb.mysql.g4.xlarge"

	// EnumOfNodeSpecForDescribeDBInstancesInputVedbMysqlG42xlarge is a EnumOfNodeSpecForDescribeDBInstancesInput enum value
	EnumOfNodeSpecForDescribeDBInstancesInputVedbMysqlG42xlarge = "vedb.mysql.g4.2xlarge"

	// EnumOfNodeSpecForDescribeDBInstancesInputVedbMysqlG82xlarge is a EnumOfNodeSpecForDescribeDBInstancesInput enum value
	EnumOfNodeSpecForDescribeDBInstancesInputVedbMysqlG82xlarge = "vedb.mysql.g8.2xlarge"

	// EnumOfNodeSpecForDescribeDBInstancesInputVedbMysqlG44xlarge is a EnumOfNodeSpecForDescribeDBInstancesInput enum value
	EnumOfNodeSpecForDescribeDBInstancesInputVedbMysqlG44xlarge = "vedb.mysql.g4.4xlarge"
)

const (
	// EnumOfNodeSpecForDescribeDBInstancesOutputVedbMysqlX4Large is a EnumOfNodeSpecForDescribeDBInstancesOutput enum value
	EnumOfNodeSpecForDescribeDBInstancesOutputVedbMysqlX4Large = "vedb.mysql.x4.large"

	// EnumOfNodeSpecForDescribeDBInstancesOutputVedbMysqlX8Large is a EnumOfNodeSpecForDescribeDBInstancesOutput enum value
	EnumOfNodeSpecForDescribeDBInstancesOutputVedbMysqlX8Large = "vedb.mysql.x8.large"

	// EnumOfNodeSpecForDescribeDBInstancesOutputVedbMysqlX4Xlarge is a EnumOfNodeSpecForDescribeDBInstancesOutput enum value
	EnumOfNodeSpecForDescribeDBInstancesOutputVedbMysqlX4Xlarge = "vedb.mysql.x4.xlarge"

	// EnumOfNodeSpecForDescribeDBInstancesOutputVedbMysqlX8Xlarge is a EnumOfNodeSpecForDescribeDBInstancesOutput enum value
	EnumOfNodeSpecForDescribeDBInstancesOutputVedbMysqlX8Xlarge = "vedb.mysql.x8.xlarge"

	// EnumOfNodeSpecForDescribeDBInstancesOutputVedbMysqlX42xlarge is a EnumOfNodeSpecForDescribeDBInstancesOutput enum value
	EnumOfNodeSpecForDescribeDBInstancesOutputVedbMysqlX42xlarge = "vedb.mysql.x4.2xlarge"

	// EnumOfNodeSpecForDescribeDBInstancesOutputVedbMysqlX82xlarge is a EnumOfNodeSpecForDescribeDBInstancesOutput enum value
	EnumOfNodeSpecForDescribeDBInstancesOutputVedbMysqlX82xlarge = "vedb.mysql.x8.2xlarge"

	// EnumOfNodeSpecForDescribeDBInstancesOutputVedbMysqlX44xlarge is a EnumOfNodeSpecForDescribeDBInstancesOutput enum value
	EnumOfNodeSpecForDescribeDBInstancesOutputVedbMysqlX44xlarge = "vedb.mysql.x4.4xlarge"

	// EnumOfNodeSpecForDescribeDBInstancesOutputVedbMysqlX84xlarge is a EnumOfNodeSpecForDescribeDBInstancesOutput enum value
	EnumOfNodeSpecForDescribeDBInstancesOutputVedbMysqlX84xlarge = "vedb.mysql.x8.4xlarge"

	// EnumOfNodeSpecForDescribeDBInstancesOutputVedbMysqlX86xlarge is a EnumOfNodeSpecForDescribeDBInstancesOutput enum value
	EnumOfNodeSpecForDescribeDBInstancesOutputVedbMysqlX86xlarge = "vedb.mysql.x8.6xlarge"

	// EnumOfNodeSpecForDescribeDBInstancesOutputVedbMysqlX48xlarge is a EnumOfNodeSpecForDescribeDBInstancesOutput enum value
	EnumOfNodeSpecForDescribeDBInstancesOutputVedbMysqlX48xlarge = "vedb.mysql.x4.8xlarge"

	// EnumOfNodeSpecForDescribeDBInstancesOutputVedbMysqlX88xlarge is a EnumOfNodeSpecForDescribeDBInstancesOutput enum value
	EnumOfNodeSpecForDescribeDBInstancesOutputVedbMysqlX88xlarge = "vedb.mysql.x8.8xlarge"

	// EnumOfNodeSpecForDescribeDBInstancesOutputVedbMysqlG4Large is a EnumOfNodeSpecForDescribeDBInstancesOutput enum value
	EnumOfNodeSpecForDescribeDBInstancesOutputVedbMysqlG4Large = "vedb.mysql.g4.large"

	// EnumOfNodeSpecForDescribeDBInstancesOutputVedbMysqlG4Xlarge is a EnumOfNodeSpecForDescribeDBInstancesOutput enum value
	EnumOfNodeSpecForDescribeDBInstancesOutputVedbMysqlG4Xlarge = "vedb.mysql.g4.xlarge"

	// EnumOfNodeSpecForDescribeDBInstancesOutputVedbMysqlG42xlarge is a EnumOfNodeSpecForDescribeDBInstancesOutput enum value
	EnumOfNodeSpecForDescribeDBInstancesOutputVedbMysqlG42xlarge = "vedb.mysql.g4.2xlarge"

	// EnumOfNodeSpecForDescribeDBInstancesOutputVedbMysqlG82xlarge is a EnumOfNodeSpecForDescribeDBInstancesOutput enum value
	EnumOfNodeSpecForDescribeDBInstancesOutputVedbMysqlG82xlarge = "vedb.mysql.g8.2xlarge"

	// EnumOfNodeSpecForDescribeDBInstancesOutputVedbMysqlG44xlarge is a EnumOfNodeSpecForDescribeDBInstancesOutput enum value
	EnumOfNodeSpecForDescribeDBInstancesOutputVedbMysqlG44xlarge = "vedb.mysql.g4.4xlarge"
)

const (
	// EnumOfNodeTypeForDescribeDBInstancesOutputPrimary is a EnumOfNodeTypeForDescribeDBInstancesOutput enum value
	EnumOfNodeTypeForDescribeDBInstancesOutputPrimary = "Primary"

	// EnumOfNodeTypeForDescribeDBInstancesOutputReadOnly is a EnumOfNodeTypeForDescribeDBInstancesOutput enum value
	EnumOfNodeTypeForDescribeDBInstancesOutputReadOnly = "ReadOnly"
)

const (
	// EnumOfStorageChargeTypeForDescribeDBInstancesOutputPostPaid is a EnumOfStorageChargeTypeForDescribeDBInstancesOutput enum value
	EnumOfStorageChargeTypeForDescribeDBInstancesOutputPostPaid = "PostPaid"

	// EnumOfStorageChargeTypeForDescribeDBInstancesOutputPrePaid is a EnumOfStorageChargeTypeForDescribeDBInstancesOutput enum value
	EnumOfStorageChargeTypeForDescribeDBInstancesOutputPrePaid = "PrePaid"
)

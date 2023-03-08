// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package rdsmysqlv2

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribeDBInstanceDetailCommon = "DescribeDBInstanceDetail"

// DescribeDBInstanceDetailCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeDBInstanceDetailCommon operation. The "output" return
// value will be populated with the DescribeDBInstanceDetailCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeDBInstanceDetailCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeDBInstanceDetailCommon Send returns without error.
//
// See DescribeDBInstanceDetailCommon for more information on using the DescribeDBInstanceDetailCommon
// API call, and error handling.
//
//    // Example sending a request using the DescribeDBInstanceDetailCommonRequest method.
//    req, resp := client.DescribeDBInstanceDetailCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *RDSMYSQLV2) DescribeDBInstanceDetailCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeDBInstanceDetailCommon,
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

// DescribeDBInstanceDetailCommon API operation for RDS_MYSQL_V2.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RDS_MYSQL_V2's
// API operation DescribeDBInstanceDetailCommon for usage and error information.
func (c *RDSMYSQLV2) DescribeDBInstanceDetailCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeDBInstanceDetailCommonRequest(input)
	return out, req.Send()
}

// DescribeDBInstanceDetailCommonWithContext is the same as DescribeDBInstanceDetailCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeDBInstanceDetailCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RDSMYSQLV2) DescribeDBInstanceDetailCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeDBInstanceDetailCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeDBInstanceDetail = "DescribeDBInstanceDetail"

// DescribeDBInstanceDetailRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeDBInstanceDetail operation. The "output" return
// value will be populated with the DescribeDBInstanceDetailCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeDBInstanceDetailCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeDBInstanceDetailCommon Send returns without error.
//
// See DescribeDBInstanceDetail for more information on using the DescribeDBInstanceDetail
// API call, and error handling.
//
//    // Example sending a request using the DescribeDBInstanceDetailRequest method.
//    req, resp := client.DescribeDBInstanceDetailRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *RDSMYSQLV2) DescribeDBInstanceDetailRequest(input *DescribeDBInstanceDetailInput) (req *request.Request, output *DescribeDBInstanceDetailOutput) {
	op := &request.Operation{
		Name:       opDescribeDBInstanceDetail,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeDBInstanceDetailInput{}
	}

	output = &DescribeDBInstanceDetailOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// DescribeDBInstanceDetail API operation for RDS_MYSQL_V2.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RDS_MYSQL_V2's
// API operation DescribeDBInstanceDetail for usage and error information.
func (c *RDSMYSQLV2) DescribeDBInstanceDetail(input *DescribeDBInstanceDetailInput) (*DescribeDBInstanceDetailOutput, error) {
	req, out := c.DescribeDBInstanceDetailRequest(input)
	return out, req.Send()
}

// DescribeDBInstanceDetailWithContext is the same as DescribeDBInstanceDetail with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeDBInstanceDetail for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RDSMYSQLV2) DescribeDBInstanceDetailWithContext(ctx volcengine.Context, input *DescribeDBInstanceDetailInput, opts ...request.Option) (*DescribeDBInstanceDetailOutput, error) {
	req, out := c.DescribeDBInstanceDetailRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type AddressForDescribeDBInstanceDetailOutput struct {
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
func (s AddressForDescribeDBInstanceDetailOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s AddressForDescribeDBInstanceDetailOutput) GoString() string {
	return s.String()
}

// SetDNSVisibility sets the DNSVisibility field's value.
func (s *AddressForDescribeDBInstanceDetailOutput) SetDNSVisibility(v bool) *AddressForDescribeDBInstanceDetailOutput {
	s.DNSVisibility = &v
	return s
}

// SetDomain sets the Domain field's value.
func (s *AddressForDescribeDBInstanceDetailOutput) SetDomain(v string) *AddressForDescribeDBInstanceDetailOutput {
	s.Domain = &v
	return s
}

// SetEipId sets the EipId field's value.
func (s *AddressForDescribeDBInstanceDetailOutput) SetEipId(v string) *AddressForDescribeDBInstanceDetailOutput {
	s.EipId = &v
	return s
}

// SetIPAddress sets the IPAddress field's value.
func (s *AddressForDescribeDBInstanceDetailOutput) SetIPAddress(v string) *AddressForDescribeDBInstanceDetailOutput {
	s.IPAddress = &v
	return s
}

// SetNetworkType sets the NetworkType field's value.
func (s *AddressForDescribeDBInstanceDetailOutput) SetNetworkType(v string) *AddressForDescribeDBInstanceDetailOutput {
	s.NetworkType = &v
	return s
}

// SetPort sets the Port field's value.
func (s *AddressForDescribeDBInstanceDetailOutput) SetPort(v string) *AddressForDescribeDBInstanceDetailOutput {
	s.Port = &v
	return s
}

// SetSubnetId sets the SubnetId field's value.
func (s *AddressForDescribeDBInstanceDetailOutput) SetSubnetId(v string) *AddressForDescribeDBInstanceDetailOutput {
	s.SubnetId = &v
	return s
}

type BasicInfoForDescribeDBInstanceDetailOutput struct {
	_ struct{} `type:"structure"`

	BackupUse *float64 `type:"double"`

	CreateTime *string `type:"string"`

	DBEngineVersion *string `type:"string"`

	DataSyncMode *string `type:"string"`

	InstanceId *string `type:"string"`

	InstanceName *string `type:"string"`

	InstanceStatus *string `type:"string"`

	LowerCaseTableNames *string `type:"string"`

	MaintenanceWindow *MaintenanceWindowForDescribeDBInstanceDetailOutput `type:"structure"`

	Memory *int32 `type:"int32"`

	NodeNumber *string `type:"string"`

	NodeSpec *string `type:"string"`

	ProjectName *string `type:"string"`

	RegionId *string `type:"string"`

	StorageSpace *int64 `type:"int64"`

	StorageType *string `type:"string"`

	StorageUse *float64 `type:"double"`

	SubnetId *string `type:"string"`

	Tags []*TagForDescribeDBInstanceDetailOutput `type:"list"`

	TimeZone *string `type:"string"`

	UpdateTime *string `type:"string"`

	VCPU *int32 `type:"int32"`

	VpcId *string `type:"string"`

	ZoneId *string `type:"string"`
}

// String returns the string representation
func (s BasicInfoForDescribeDBInstanceDetailOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s BasicInfoForDescribeDBInstanceDetailOutput) GoString() string {
	return s.String()
}

// SetBackupUse sets the BackupUse field's value.
func (s *BasicInfoForDescribeDBInstanceDetailOutput) SetBackupUse(v float64) *BasicInfoForDescribeDBInstanceDetailOutput {
	s.BackupUse = &v
	return s
}

// SetCreateTime sets the CreateTime field's value.
func (s *BasicInfoForDescribeDBInstanceDetailOutput) SetCreateTime(v string) *BasicInfoForDescribeDBInstanceDetailOutput {
	s.CreateTime = &v
	return s
}

// SetDBEngineVersion sets the DBEngineVersion field's value.
func (s *BasicInfoForDescribeDBInstanceDetailOutput) SetDBEngineVersion(v string) *BasicInfoForDescribeDBInstanceDetailOutput {
	s.DBEngineVersion = &v
	return s
}

// SetDataSyncMode sets the DataSyncMode field's value.
func (s *BasicInfoForDescribeDBInstanceDetailOutput) SetDataSyncMode(v string) *BasicInfoForDescribeDBInstanceDetailOutput {
	s.DataSyncMode = &v
	return s
}

// SetInstanceId sets the InstanceId field's value.
func (s *BasicInfoForDescribeDBInstanceDetailOutput) SetInstanceId(v string) *BasicInfoForDescribeDBInstanceDetailOutput {
	s.InstanceId = &v
	return s
}

// SetInstanceName sets the InstanceName field's value.
func (s *BasicInfoForDescribeDBInstanceDetailOutput) SetInstanceName(v string) *BasicInfoForDescribeDBInstanceDetailOutput {
	s.InstanceName = &v
	return s
}

// SetInstanceStatus sets the InstanceStatus field's value.
func (s *BasicInfoForDescribeDBInstanceDetailOutput) SetInstanceStatus(v string) *BasicInfoForDescribeDBInstanceDetailOutput {
	s.InstanceStatus = &v
	return s
}

// SetLowerCaseTableNames sets the LowerCaseTableNames field's value.
func (s *BasicInfoForDescribeDBInstanceDetailOutput) SetLowerCaseTableNames(v string) *BasicInfoForDescribeDBInstanceDetailOutput {
	s.LowerCaseTableNames = &v
	return s
}

// SetMaintenanceWindow sets the MaintenanceWindow field's value.
func (s *BasicInfoForDescribeDBInstanceDetailOutput) SetMaintenanceWindow(v *MaintenanceWindowForDescribeDBInstanceDetailOutput) *BasicInfoForDescribeDBInstanceDetailOutput {
	s.MaintenanceWindow = v
	return s
}

// SetMemory sets the Memory field's value.
func (s *BasicInfoForDescribeDBInstanceDetailOutput) SetMemory(v int32) *BasicInfoForDescribeDBInstanceDetailOutput {
	s.Memory = &v
	return s
}

// SetNodeNumber sets the NodeNumber field's value.
func (s *BasicInfoForDescribeDBInstanceDetailOutput) SetNodeNumber(v string) *BasicInfoForDescribeDBInstanceDetailOutput {
	s.NodeNumber = &v
	return s
}

// SetNodeSpec sets the NodeSpec field's value.
func (s *BasicInfoForDescribeDBInstanceDetailOutput) SetNodeSpec(v string) *BasicInfoForDescribeDBInstanceDetailOutput {
	s.NodeSpec = &v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *BasicInfoForDescribeDBInstanceDetailOutput) SetProjectName(v string) *BasicInfoForDescribeDBInstanceDetailOutput {
	s.ProjectName = &v
	return s
}

// SetRegionId sets the RegionId field's value.
func (s *BasicInfoForDescribeDBInstanceDetailOutput) SetRegionId(v string) *BasicInfoForDescribeDBInstanceDetailOutput {
	s.RegionId = &v
	return s
}

// SetStorageSpace sets the StorageSpace field's value.
func (s *BasicInfoForDescribeDBInstanceDetailOutput) SetStorageSpace(v int64) *BasicInfoForDescribeDBInstanceDetailOutput {
	s.StorageSpace = &v
	return s
}

// SetStorageType sets the StorageType field's value.
func (s *BasicInfoForDescribeDBInstanceDetailOutput) SetStorageType(v string) *BasicInfoForDescribeDBInstanceDetailOutput {
	s.StorageType = &v
	return s
}

// SetStorageUse sets the StorageUse field's value.
func (s *BasicInfoForDescribeDBInstanceDetailOutput) SetStorageUse(v float64) *BasicInfoForDescribeDBInstanceDetailOutput {
	s.StorageUse = &v
	return s
}

// SetSubnetId sets the SubnetId field's value.
func (s *BasicInfoForDescribeDBInstanceDetailOutput) SetSubnetId(v string) *BasicInfoForDescribeDBInstanceDetailOutput {
	s.SubnetId = &v
	return s
}

// SetTags sets the Tags field's value.
func (s *BasicInfoForDescribeDBInstanceDetailOutput) SetTags(v []*TagForDescribeDBInstanceDetailOutput) *BasicInfoForDescribeDBInstanceDetailOutput {
	s.Tags = v
	return s
}

// SetTimeZone sets the TimeZone field's value.
func (s *BasicInfoForDescribeDBInstanceDetailOutput) SetTimeZone(v string) *BasicInfoForDescribeDBInstanceDetailOutput {
	s.TimeZone = &v
	return s
}

// SetUpdateTime sets the UpdateTime field's value.
func (s *BasicInfoForDescribeDBInstanceDetailOutput) SetUpdateTime(v string) *BasicInfoForDescribeDBInstanceDetailOutput {
	s.UpdateTime = &v
	return s
}

// SetVCPU sets the VCPU field's value.
func (s *BasicInfoForDescribeDBInstanceDetailOutput) SetVCPU(v int32) *BasicInfoForDescribeDBInstanceDetailOutput {
	s.VCPU = &v
	return s
}

// SetVpcId sets the VpcId field's value.
func (s *BasicInfoForDescribeDBInstanceDetailOutput) SetVpcId(v string) *BasicInfoForDescribeDBInstanceDetailOutput {
	s.VpcId = &v
	return s
}

// SetZoneId sets the ZoneId field's value.
func (s *BasicInfoForDescribeDBInstanceDetailOutput) SetZoneId(v string) *BasicInfoForDescribeDBInstanceDetailOutput {
	s.ZoneId = &v
	return s
}

type ChargeDetailForDescribeDBInstanceDetailOutput struct {
	_ struct{} `type:"structure"`

	AutoRenew *bool `type:"boolean"`

	ChargeEndTime *string `type:"string"`

	ChargeStartTime *string `type:"string"`

	ChargeStatus *string `type:"string"`

	ChargeType *string `type:"string"`

	OverdueReclaimTime *string `type:"string"`

	OverdueTime *string `type:"string"`

	Period *int32 `type:"int32"`

	PeriodUnit *string `type:"string"`
}

// String returns the string representation
func (s ChargeDetailForDescribeDBInstanceDetailOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ChargeDetailForDescribeDBInstanceDetailOutput) GoString() string {
	return s.String()
}

// SetAutoRenew sets the AutoRenew field's value.
func (s *ChargeDetailForDescribeDBInstanceDetailOutput) SetAutoRenew(v bool) *ChargeDetailForDescribeDBInstanceDetailOutput {
	s.AutoRenew = &v
	return s
}

// SetChargeEndTime sets the ChargeEndTime field's value.
func (s *ChargeDetailForDescribeDBInstanceDetailOutput) SetChargeEndTime(v string) *ChargeDetailForDescribeDBInstanceDetailOutput {
	s.ChargeEndTime = &v
	return s
}

// SetChargeStartTime sets the ChargeStartTime field's value.
func (s *ChargeDetailForDescribeDBInstanceDetailOutput) SetChargeStartTime(v string) *ChargeDetailForDescribeDBInstanceDetailOutput {
	s.ChargeStartTime = &v
	return s
}

// SetChargeStatus sets the ChargeStatus field's value.
func (s *ChargeDetailForDescribeDBInstanceDetailOutput) SetChargeStatus(v string) *ChargeDetailForDescribeDBInstanceDetailOutput {
	s.ChargeStatus = &v
	return s
}

// SetChargeType sets the ChargeType field's value.
func (s *ChargeDetailForDescribeDBInstanceDetailOutput) SetChargeType(v string) *ChargeDetailForDescribeDBInstanceDetailOutput {
	s.ChargeType = &v
	return s
}

// SetOverdueReclaimTime sets the OverdueReclaimTime field's value.
func (s *ChargeDetailForDescribeDBInstanceDetailOutput) SetOverdueReclaimTime(v string) *ChargeDetailForDescribeDBInstanceDetailOutput {
	s.OverdueReclaimTime = &v
	return s
}

// SetOverdueTime sets the OverdueTime field's value.
func (s *ChargeDetailForDescribeDBInstanceDetailOutput) SetOverdueTime(v string) *ChargeDetailForDescribeDBInstanceDetailOutput {
	s.OverdueTime = &v
	return s
}

// SetPeriod sets the Period field's value.
func (s *ChargeDetailForDescribeDBInstanceDetailOutput) SetPeriod(v int32) *ChargeDetailForDescribeDBInstanceDetailOutput {
	s.Period = &v
	return s
}

// SetPeriodUnit sets the PeriodUnit field's value.
func (s *ChargeDetailForDescribeDBInstanceDetailOutput) SetPeriodUnit(v string) *ChargeDetailForDescribeDBInstanceDetailOutput {
	s.PeriodUnit = &v
	return s
}

type DescribeDBInstanceDetailInput struct {
	_ struct{} `type:"structure"`

	// InstanceId is a required field
	InstanceId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeDBInstanceDetailInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeDBInstanceDetailInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeDBInstanceDetailInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DescribeDBInstanceDetailInput"}
	if s.InstanceId == nil {
		invalidParams.Add(request.NewErrParamRequired("InstanceId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetInstanceId sets the InstanceId field's value.
func (s *DescribeDBInstanceDetailInput) SetInstanceId(v string) *DescribeDBInstanceDetailInput {
	s.InstanceId = &v
	return s
}

type DescribeDBInstanceDetailOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	BasicInfo *BasicInfoForDescribeDBInstanceDetailOutput `type:"structure"`

	ChargeDetail *ChargeDetailForDescribeDBInstanceDetailOutput `type:"structure"`

	Endpoints []*EndpointForDescribeDBInstanceDetailOutput `type:"list"`

	Nodes []*NodeForDescribeDBInstanceDetailOutput `type:"list"`
}

// String returns the string representation
func (s DescribeDBInstanceDetailOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeDBInstanceDetailOutput) GoString() string {
	return s.String()
}

// SetBasicInfo sets the BasicInfo field's value.
func (s *DescribeDBInstanceDetailOutput) SetBasicInfo(v *BasicInfoForDescribeDBInstanceDetailOutput) *DescribeDBInstanceDetailOutput {
	s.BasicInfo = v
	return s
}

// SetChargeDetail sets the ChargeDetail field's value.
func (s *DescribeDBInstanceDetailOutput) SetChargeDetail(v *ChargeDetailForDescribeDBInstanceDetailOutput) *DescribeDBInstanceDetailOutput {
	s.ChargeDetail = v
	return s
}

// SetEndpoints sets the Endpoints field's value.
func (s *DescribeDBInstanceDetailOutput) SetEndpoints(v []*EndpointForDescribeDBInstanceDetailOutput) *DescribeDBInstanceDetailOutput {
	s.Endpoints = v
	return s
}

// SetNodes sets the Nodes field's value.
func (s *DescribeDBInstanceDetailOutput) SetNodes(v []*NodeForDescribeDBInstanceDetailOutput) *DescribeDBInstanceDetailOutput {
	s.Nodes = v
	return s
}

type EndpointForDescribeDBInstanceDetailOutput struct {
	_ struct{} `type:"structure"`

	Addresses []*AddressForDescribeDBInstanceDetailOutput `type:"list"`

	AutoAddNewNodes *string `type:"string"`

	Description *string `type:"string"`

	EnableReadOnly *string `type:"string"`

	EnableReadWriteSplitting *string `type:"string"`

	EndpointId *string `type:"string"`

	EndpointName *string `type:"string"`

	EndpointType *string `type:"string"`

	ReadOnlyNodeWeight []*ReadOnlyNodeWeightForDescribeDBInstanceDetailOutput `type:"list"`

	ReadWriteMode *string `type:"string"`
}

// String returns the string representation
func (s EndpointForDescribeDBInstanceDetailOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s EndpointForDescribeDBInstanceDetailOutput) GoString() string {
	return s.String()
}

// SetAddresses sets the Addresses field's value.
func (s *EndpointForDescribeDBInstanceDetailOutput) SetAddresses(v []*AddressForDescribeDBInstanceDetailOutput) *EndpointForDescribeDBInstanceDetailOutput {
	s.Addresses = v
	return s
}

// SetAutoAddNewNodes sets the AutoAddNewNodes field's value.
func (s *EndpointForDescribeDBInstanceDetailOutput) SetAutoAddNewNodes(v string) *EndpointForDescribeDBInstanceDetailOutput {
	s.AutoAddNewNodes = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *EndpointForDescribeDBInstanceDetailOutput) SetDescription(v string) *EndpointForDescribeDBInstanceDetailOutput {
	s.Description = &v
	return s
}

// SetEnableReadOnly sets the EnableReadOnly field's value.
func (s *EndpointForDescribeDBInstanceDetailOutput) SetEnableReadOnly(v string) *EndpointForDescribeDBInstanceDetailOutput {
	s.EnableReadOnly = &v
	return s
}

// SetEnableReadWriteSplitting sets the EnableReadWriteSplitting field's value.
func (s *EndpointForDescribeDBInstanceDetailOutput) SetEnableReadWriteSplitting(v string) *EndpointForDescribeDBInstanceDetailOutput {
	s.EnableReadWriteSplitting = &v
	return s
}

// SetEndpointId sets the EndpointId field's value.
func (s *EndpointForDescribeDBInstanceDetailOutput) SetEndpointId(v string) *EndpointForDescribeDBInstanceDetailOutput {
	s.EndpointId = &v
	return s
}

// SetEndpointName sets the EndpointName field's value.
func (s *EndpointForDescribeDBInstanceDetailOutput) SetEndpointName(v string) *EndpointForDescribeDBInstanceDetailOutput {
	s.EndpointName = &v
	return s
}

// SetEndpointType sets the EndpointType field's value.
func (s *EndpointForDescribeDBInstanceDetailOutput) SetEndpointType(v string) *EndpointForDescribeDBInstanceDetailOutput {
	s.EndpointType = &v
	return s
}

// SetReadOnlyNodeWeight sets the ReadOnlyNodeWeight field's value.
func (s *EndpointForDescribeDBInstanceDetailOutput) SetReadOnlyNodeWeight(v []*ReadOnlyNodeWeightForDescribeDBInstanceDetailOutput) *EndpointForDescribeDBInstanceDetailOutput {
	s.ReadOnlyNodeWeight = v
	return s
}

// SetReadWriteMode sets the ReadWriteMode field's value.
func (s *EndpointForDescribeDBInstanceDetailOutput) SetReadWriteMode(v string) *EndpointForDescribeDBInstanceDetailOutput {
	s.ReadWriteMode = &v
	return s
}

type MaintenanceWindowForDescribeDBInstanceDetailOutput struct {
	_ struct{} `type:"structure"`

	DayKind *string `type:"string"`

	DayOfMonth []*int64 `type:"list"`

	DayOfWeek []*string `type:"list"`

	MaintenanceTime *string `type:"string"`
}

// String returns the string representation
func (s MaintenanceWindowForDescribeDBInstanceDetailOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s MaintenanceWindowForDescribeDBInstanceDetailOutput) GoString() string {
	return s.String()
}

// SetDayKind sets the DayKind field's value.
func (s *MaintenanceWindowForDescribeDBInstanceDetailOutput) SetDayKind(v string) *MaintenanceWindowForDescribeDBInstanceDetailOutput {
	s.DayKind = &v
	return s
}

// SetDayOfMonth sets the DayOfMonth field's value.
func (s *MaintenanceWindowForDescribeDBInstanceDetailOutput) SetDayOfMonth(v []*int64) *MaintenanceWindowForDescribeDBInstanceDetailOutput {
	s.DayOfMonth = v
	return s
}

// SetDayOfWeek sets the DayOfWeek field's value.
func (s *MaintenanceWindowForDescribeDBInstanceDetailOutput) SetDayOfWeek(v []*string) *MaintenanceWindowForDescribeDBInstanceDetailOutput {
	s.DayOfWeek = v
	return s
}

// SetMaintenanceTime sets the MaintenanceTime field's value.
func (s *MaintenanceWindowForDescribeDBInstanceDetailOutput) SetMaintenanceTime(v string) *MaintenanceWindowForDescribeDBInstanceDetailOutput {
	s.MaintenanceTime = &v
	return s
}

type NodeForDescribeDBInstanceDetailOutput struct {
	_ struct{} `type:"structure"`

	CreateTime *string `type:"string"`

	InstanceId *string `type:"string"`

	Memory *int32 `type:"int32"`

	NodeId *string `type:"string"`

	NodeSpec *string `type:"string"`

	NodeStatus *string `type:"string"`

	NodeType *string `type:"string"`

	RegionId *string `type:"string"`

	UpdateTime *string `type:"string"`

	VCPU *int32 `type:"int32"`

	ZoneId *string `type:"string"`
}

// String returns the string representation
func (s NodeForDescribeDBInstanceDetailOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s NodeForDescribeDBInstanceDetailOutput) GoString() string {
	return s.String()
}

// SetCreateTime sets the CreateTime field's value.
func (s *NodeForDescribeDBInstanceDetailOutput) SetCreateTime(v string) *NodeForDescribeDBInstanceDetailOutput {
	s.CreateTime = &v
	return s
}

// SetInstanceId sets the InstanceId field's value.
func (s *NodeForDescribeDBInstanceDetailOutput) SetInstanceId(v string) *NodeForDescribeDBInstanceDetailOutput {
	s.InstanceId = &v
	return s
}

// SetMemory sets the Memory field's value.
func (s *NodeForDescribeDBInstanceDetailOutput) SetMemory(v int32) *NodeForDescribeDBInstanceDetailOutput {
	s.Memory = &v
	return s
}

// SetNodeId sets the NodeId field's value.
func (s *NodeForDescribeDBInstanceDetailOutput) SetNodeId(v string) *NodeForDescribeDBInstanceDetailOutput {
	s.NodeId = &v
	return s
}

// SetNodeSpec sets the NodeSpec field's value.
func (s *NodeForDescribeDBInstanceDetailOutput) SetNodeSpec(v string) *NodeForDescribeDBInstanceDetailOutput {
	s.NodeSpec = &v
	return s
}

// SetNodeStatus sets the NodeStatus field's value.
func (s *NodeForDescribeDBInstanceDetailOutput) SetNodeStatus(v string) *NodeForDescribeDBInstanceDetailOutput {
	s.NodeStatus = &v
	return s
}

// SetNodeType sets the NodeType field's value.
func (s *NodeForDescribeDBInstanceDetailOutput) SetNodeType(v string) *NodeForDescribeDBInstanceDetailOutput {
	s.NodeType = &v
	return s
}

// SetRegionId sets the RegionId field's value.
func (s *NodeForDescribeDBInstanceDetailOutput) SetRegionId(v string) *NodeForDescribeDBInstanceDetailOutput {
	s.RegionId = &v
	return s
}

// SetUpdateTime sets the UpdateTime field's value.
func (s *NodeForDescribeDBInstanceDetailOutput) SetUpdateTime(v string) *NodeForDescribeDBInstanceDetailOutput {
	s.UpdateTime = &v
	return s
}

// SetVCPU sets the VCPU field's value.
func (s *NodeForDescribeDBInstanceDetailOutput) SetVCPU(v int32) *NodeForDescribeDBInstanceDetailOutput {
	s.VCPU = &v
	return s
}

// SetZoneId sets the ZoneId field's value.
func (s *NodeForDescribeDBInstanceDetailOutput) SetZoneId(v string) *NodeForDescribeDBInstanceDetailOutput {
	s.ZoneId = &v
	return s
}

type ReadOnlyNodeWeightForDescribeDBInstanceDetailOutput struct {
	_ struct{} `type:"structure"`

	NodeId *string `type:"string"`

	NodeType *string `type:"string"`

	Weight *int32 `type:"int32"`
}

// String returns the string representation
func (s ReadOnlyNodeWeightForDescribeDBInstanceDetailOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ReadOnlyNodeWeightForDescribeDBInstanceDetailOutput) GoString() string {
	return s.String()
}

// SetNodeId sets the NodeId field's value.
func (s *ReadOnlyNodeWeightForDescribeDBInstanceDetailOutput) SetNodeId(v string) *ReadOnlyNodeWeightForDescribeDBInstanceDetailOutput {
	s.NodeId = &v
	return s
}

// SetNodeType sets the NodeType field's value.
func (s *ReadOnlyNodeWeightForDescribeDBInstanceDetailOutput) SetNodeType(v string) *ReadOnlyNodeWeightForDescribeDBInstanceDetailOutput {
	s.NodeType = &v
	return s
}

// SetWeight sets the Weight field's value.
func (s *ReadOnlyNodeWeightForDescribeDBInstanceDetailOutput) SetWeight(v int32) *ReadOnlyNodeWeightForDescribeDBInstanceDetailOutput {
	s.Weight = &v
	return s
}

type TagForDescribeDBInstanceDetailOutput struct {
	_ struct{} `type:"structure"`

	Key *string `type:"string"`

	Value *string `type:"string"`
}

// String returns the string representation
func (s TagForDescribeDBInstanceDetailOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s TagForDescribeDBInstanceDetailOutput) GoString() string {
	return s.String()
}

// SetKey sets the Key field's value.
func (s *TagForDescribeDBInstanceDetailOutput) SetKey(v string) *TagForDescribeDBInstanceDetailOutput {
	s.Key = &v
	return s
}

// SetValue sets the Value field's value.
func (s *TagForDescribeDBInstanceDetailOutput) SetValue(v string) *TagForDescribeDBInstanceDetailOutput {
	s.Value = &v
	return s
}

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
//    // Example sending a request using the DescribeDBInstancesCommonRequest method.
//    req, resp := client.DescribeDBInstancesCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
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
//    // Example sending a request using the DescribeDBInstancesRequest method.
//    req, resp := client.DescribeDBInstancesRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
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
	_ struct{} `type:"structure" json:",omitempty"`

	DNSVisibility *bool `type:"boolean" json:",omitempty"`

	Domain *string `type:"string" json:",omitempty"`

	EipId *string `type:"string" json:",omitempty"`

	IPAddress *string `type:"string" json:",omitempty"`

	InternetProtocol *string `type:"string" json:",omitempty"`

	NetworkType *string `type:"string" json:",omitempty"`

	Port *string `type:"string" json:",omitempty"`

	SubnetId *string `type:"string" json:",omitempty"`
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

// SetInternetProtocol sets the InternetProtocol field's value.
func (s *AddressObjectForDescribeDBInstancesOutput) SetInternetProtocol(v string) *AddressObjectForDescribeDBInstancesOutput {
	s.InternetProtocol = &v
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

	TempModifyEndTime *string `type:"string" json:",omitempty"`

	TempModifyStartTime *string `type:"string" json:",omitempty"`
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

// SetTempModifyEndTime sets the TempModifyEndTime field's value.
func (s *ChargeDetailForDescribeDBInstancesOutput) SetTempModifyEndTime(v string) *ChargeDetailForDescribeDBInstancesOutput {
	s.TempModifyEndTime = &v
	return s
}

// SetTempModifyStartTime sets the TempModifyStartTime field's value.
func (s *ChargeDetailForDescribeDBInstancesOutput) SetTempModifyStartTime(v string) *ChargeDetailForDescribeDBInstancesOutput {
	s.TempModifyStartTime = &v
	return s
}

type DescribeDBInstancesInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	ChargeType *string `type:"string" json:",omitempty"`

	CreateTimeEnd *string `type:"string" json:",omitempty"`

	CreateTimeStart *string `type:"string" json:",omitempty"`

	DBEngineVersion *string `type:"string" json:",omitempty"`

	InstanceId *string `type:"string" json:",omitempty"`

	InstanceName *string `type:"string" json:",omitempty"`

	InstanceStatus *string `type:"string" json:",omitempty"`

	InstanceType *string `type:"string" json:",omitempty"`

	KernelVersion []*string `type:"list" json:",omitempty"`

	NodeSpec *string `type:"string" json:",omitempty"`

	PageNumber *int32 `type:"int32" json:",omitempty"`

	PageSize *int32 `type:"int32" json:",omitempty"`

	PrivateNetworkIpAddress *string `type:"string" json:",omitempty"`

	PrivateNetworkVpcId *string `type:"string" json:",omitempty"`

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

// SetInstanceType sets the InstanceType field's value.
func (s *DescribeDBInstancesInput) SetInstanceType(v string) *DescribeDBInstancesInput {
	s.InstanceType = &v
	return s
}

// SetKernelVersion sets the KernelVersion field's value.
func (s *DescribeDBInstancesInput) SetKernelVersion(v []*string) *DescribeDBInstancesInput {
	s.KernelVersion = v
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

// SetPrivateNetworkIpAddress sets the PrivateNetworkIpAddress field's value.
func (s *DescribeDBInstancesInput) SetPrivateNetworkIpAddress(v string) *DescribeDBInstancesInput {
	s.PrivateNetworkIpAddress = &v
	return s
}

// SetPrivateNetworkVpcId sets the PrivateNetworkVpcId field's value.
func (s *DescribeDBInstancesInput) SetPrivateNetworkVpcId(v string) *DescribeDBInstancesInput {
	s.PrivateNetworkVpcId = &v
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

	AddressObject []*AddressObjectForDescribeDBInstancesOutput `type:"list" json:",omitempty"`

	AllowListVersion *string `type:"string" json:",omitempty"`

	ChargeDetail *ChargeDetailForDescribeDBInstancesOutput `type:"structure" json:",omitempty"`

	CreateTime *string `type:"string" json:",omitempty"`

	CurrentKernelVersion *string `type:"string" json:",omitempty"`

	DBEngineVersion *string `type:"string" json:",omitempty"`

	InstanceId *string `type:"string" json:",omitempty"`

	InstanceName *string `type:"string" json:",omitempty"`

	InstanceStatus *string `type:"string" json:",omitempty"`

	InstanceType *string `type:"string" json:",omitempty"`

	LowerCaseTableNames *string `type:"string" json:",omitempty"`

	MaintenanceWindow *MaintenanceWindowForDescribeDBInstancesOutput `type:"structure" json:",omitempty"`

	NodeCPUUsedPercentage *float64 `type:"double" json:",omitempty"`

	NodeMemoryUsedPercentage *float64 `type:"double" json:",omitempty"`

	NodeNumber *int32 `type:"int32" json:",omitempty"`

	NodeSpaceUsedPercentage *float64 `type:"double" json:",omitempty"`

	NodeSpec *string `type:"string" json:",omitempty"`

	ProjectName *string `type:"string" json:",omitempty"`

	RegionId *string `type:"string" json:",omitempty"`

	StorageSpace *int32 `type:"int32" json:",omitempty"`

	StorageType *string `type:"string" json:",omitempty"`

	SubnetId *string `type:"string" json:",omitempty"`

	Tags []*TagForDescribeDBInstancesOutput `type:"list" json:",omitempty"`

	TimeZone *string `type:"string" json:",omitempty"`

	VpcId *string `type:"string" json:",omitempty"`

	ZoneId *string `type:"string" json:",omitempty"`

	ZoneIds []*string `type:"list" json:",omitempty"`
}

// String returns the string representation
func (s InstanceForDescribeDBInstancesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s InstanceForDescribeDBInstancesOutput) GoString() string {
	return s.String()
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

// SetCurrentKernelVersion sets the CurrentKernelVersion field's value.
func (s *InstanceForDescribeDBInstancesOutput) SetCurrentKernelVersion(v string) *InstanceForDescribeDBInstancesOutput {
	s.CurrentKernelVersion = &v
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

// SetMaintenanceWindow sets the MaintenanceWindow field's value.
func (s *InstanceForDescribeDBInstancesOutput) SetMaintenanceWindow(v *MaintenanceWindowForDescribeDBInstancesOutput) *InstanceForDescribeDBInstancesOutput {
	s.MaintenanceWindow = v
	return s
}

// SetNodeCPUUsedPercentage sets the NodeCPUUsedPercentage field's value.
func (s *InstanceForDescribeDBInstancesOutput) SetNodeCPUUsedPercentage(v float64) *InstanceForDescribeDBInstancesOutput {
	s.NodeCPUUsedPercentage = &v
	return s
}

// SetNodeMemoryUsedPercentage sets the NodeMemoryUsedPercentage field's value.
func (s *InstanceForDescribeDBInstancesOutput) SetNodeMemoryUsedPercentage(v float64) *InstanceForDescribeDBInstancesOutput {
	s.NodeMemoryUsedPercentage = &v
	return s
}

// SetNodeNumber sets the NodeNumber field's value.
func (s *InstanceForDescribeDBInstancesOutput) SetNodeNumber(v int32) *InstanceForDescribeDBInstancesOutput {
	s.NodeNumber = &v
	return s
}

// SetNodeSpaceUsedPercentage sets the NodeSpaceUsedPercentage field's value.
func (s *InstanceForDescribeDBInstancesOutput) SetNodeSpaceUsedPercentage(v float64) *InstanceForDescribeDBInstancesOutput {
	s.NodeSpaceUsedPercentage = &v
	return s
}

// SetNodeSpec sets the NodeSpec field's value.
func (s *InstanceForDescribeDBInstancesOutput) SetNodeSpec(v string) *InstanceForDescribeDBInstancesOutput {
	s.NodeSpec = &v
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

// SetZoneId sets the ZoneId field's value.
func (s *InstanceForDescribeDBInstancesOutput) SetZoneId(v string) *InstanceForDescribeDBInstancesOutput {
	s.ZoneId = &v
	return s
}

// SetZoneIds sets the ZoneIds field's value.
func (s *InstanceForDescribeDBInstancesOutput) SetZoneIds(v []*string) *InstanceForDescribeDBInstancesOutput {
	s.ZoneIds = v
	return s
}

type MaintenanceWindowForDescribeDBInstancesOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	DayKind *string `type:"string" json:",omitempty"`

	DayOfWeek []*string `type:"list" json:",omitempty"`

	MaintenanceTime *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s MaintenanceWindowForDescribeDBInstancesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s MaintenanceWindowForDescribeDBInstancesOutput) GoString() string {
	return s.String()
}

// SetDayKind sets the DayKind field's value.
func (s *MaintenanceWindowForDescribeDBInstancesOutput) SetDayKind(v string) *MaintenanceWindowForDescribeDBInstancesOutput {
	s.DayKind = &v
	return s
}

// SetDayOfWeek sets the DayOfWeek field's value.
func (s *MaintenanceWindowForDescribeDBInstancesOutput) SetDayOfWeek(v []*string) *MaintenanceWindowForDescribeDBInstancesOutput {
	s.DayOfWeek = v
	return s
}

// SetMaintenanceTime sets the MaintenanceTime field's value.
func (s *MaintenanceWindowForDescribeDBInstancesOutput) SetMaintenanceTime(v string) *MaintenanceWindowForDescribeDBInstancesOutput {
	s.MaintenanceTime = &v
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

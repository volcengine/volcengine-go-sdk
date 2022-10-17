// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package rdsmysql

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opListDBInstancesCommon = "ListDBInstances"

// ListDBInstancesCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ListDBInstancesCommon operation. The "output" return
// value will be populated with the ListDBInstancesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListDBInstancesCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListDBInstancesCommon Send returns without error.
//
// See ListDBInstancesCommon for more information on using the ListDBInstancesCommon
// API call, and error handling.
//
//    // Example sending a request using the ListDBInstancesCommonRequest method.
//    req, resp := client.ListDBInstancesCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *RDSMYSQL) ListDBInstancesCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opListDBInstancesCommon,
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

// ListDBInstancesCommon API operation for RDS_MYSQL.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RDS_MYSQL's
// API operation ListDBInstancesCommon for usage and error information.
func (c *RDSMYSQL) ListDBInstancesCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ListDBInstancesCommonRequest(input)
	return out, req.Send()
}

// ListDBInstancesCommonWithContext is the same as ListDBInstancesCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ListDBInstancesCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RDSMYSQL) ListDBInstancesCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ListDBInstancesCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opListDBInstances = "ListDBInstances"

// ListDBInstancesRequest generates a "volcengine/request.Request" representing the
// client's request for the ListDBInstances operation. The "output" return
// value will be populated with the ListDBInstancesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListDBInstancesCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListDBInstancesCommon Send returns without error.
//
// See ListDBInstances for more information on using the ListDBInstances
// API call, and error handling.
//
//    // Example sending a request using the ListDBInstancesRequest method.
//    req, resp := client.ListDBInstancesRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *RDSMYSQL) ListDBInstancesRequest(input *ListDBInstancesInput) (req *request.Request, output *ListDBInstancesOutput) {
	op := &request.Operation{
		Name:       opListDBInstances,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListDBInstancesInput{}
	}

	output = &ListDBInstancesOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// ListDBInstances API operation for RDS_MYSQL.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RDS_MYSQL's
// API operation ListDBInstances for usage and error information.
func (c *RDSMYSQL) ListDBInstances(input *ListDBInstancesInput) (*ListDBInstancesOutput, error) {
	req, out := c.ListDBInstancesRequest(input)
	return out, req.Send()
}

// ListDBInstancesWithContext is the same as ListDBInstances with the addition of
// the ability to pass a context and additional request options.
//
// See ListDBInstances for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RDSMYSQL) ListDBInstancesWithContext(ctx volcengine.Context, input *ListDBInstancesInput, opts ...request.Option) (*ListDBInstancesOutput, error) {
	req, out := c.ListDBInstancesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DataForListDBInstancesOutput struct {
	_ struct{} `type:"structure"`

	ChargeStatus *string `type:"string" enum:"EnumOfChargeStatusForListDBInstancesOutput"`

	ChargeType *string `type:"string" enum:"EnumOfChargeTypeForListDBInstancesOutput"`

	CreateTime *string `type:"string"`

	DBEngine *string `type:"string" enum:"EnumOfDBEngineForListDBInstancesOutput"`

	DBEngineVersion *string `type:"string" enum:"EnumOfDBEngineVersionForListDBInstancesOutput"`

	InstanceId *string `type:"string"`

	InstanceName *string `type:"string"`

	InstanceSpec *InstanceSpecForListDBInstancesOutput `type:"structure"`

	InstanceStatus *string `type:"string" enum:"EnumOfInstanceStatusForListDBInstancesOutput"`

	InstanceType *string `type:"string" enum:"EnumOfInstanceTypeForListDBInstancesOutput"`

	ReadOnlyInstanceIds []*string `type:"list"`

	Region *string `type:"string"`

	StorageSpaceGB *int32 `type:"int32"`

	UpdateTime *string `type:"string"`

	VpcID *string `type:"string"`

	Zone *string `type:"string"`
}

// String returns the string representation
func (s DataForListDBInstancesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DataForListDBInstancesOutput) GoString() string {
	return s.String()
}

// SetChargeStatus sets the ChargeStatus field's value.
func (s *DataForListDBInstancesOutput) SetChargeStatus(v string) *DataForListDBInstancesOutput {
	s.ChargeStatus = &v
	return s
}

// SetChargeType sets the ChargeType field's value.
func (s *DataForListDBInstancesOutput) SetChargeType(v string) *DataForListDBInstancesOutput {
	s.ChargeType = &v
	return s
}

// SetCreateTime sets the CreateTime field's value.
func (s *DataForListDBInstancesOutput) SetCreateTime(v string) *DataForListDBInstancesOutput {
	s.CreateTime = &v
	return s
}

// SetDBEngine sets the DBEngine field's value.
func (s *DataForListDBInstancesOutput) SetDBEngine(v string) *DataForListDBInstancesOutput {
	s.DBEngine = &v
	return s
}

// SetDBEngineVersion sets the DBEngineVersion field's value.
func (s *DataForListDBInstancesOutput) SetDBEngineVersion(v string) *DataForListDBInstancesOutput {
	s.DBEngineVersion = &v
	return s
}

// SetInstanceId sets the InstanceId field's value.
func (s *DataForListDBInstancesOutput) SetInstanceId(v string) *DataForListDBInstancesOutput {
	s.InstanceId = &v
	return s
}

// SetInstanceName sets the InstanceName field's value.
func (s *DataForListDBInstancesOutput) SetInstanceName(v string) *DataForListDBInstancesOutput {
	s.InstanceName = &v
	return s
}

// SetInstanceSpec sets the InstanceSpec field's value.
func (s *DataForListDBInstancesOutput) SetInstanceSpec(v *InstanceSpecForListDBInstancesOutput) *DataForListDBInstancesOutput {
	s.InstanceSpec = v
	return s
}

// SetInstanceStatus sets the InstanceStatus field's value.
func (s *DataForListDBInstancesOutput) SetInstanceStatus(v string) *DataForListDBInstancesOutput {
	s.InstanceStatus = &v
	return s
}

// SetInstanceType sets the InstanceType field's value.
func (s *DataForListDBInstancesOutput) SetInstanceType(v string) *DataForListDBInstancesOutput {
	s.InstanceType = &v
	return s
}

// SetReadOnlyInstanceIds sets the ReadOnlyInstanceIds field's value.
func (s *DataForListDBInstancesOutput) SetReadOnlyInstanceIds(v []*string) *DataForListDBInstancesOutput {
	s.ReadOnlyInstanceIds = v
	return s
}

// SetRegion sets the Region field's value.
func (s *DataForListDBInstancesOutput) SetRegion(v string) *DataForListDBInstancesOutput {
	s.Region = &v
	return s
}

// SetStorageSpaceGB sets the StorageSpaceGB field's value.
func (s *DataForListDBInstancesOutput) SetStorageSpaceGB(v int32) *DataForListDBInstancesOutput {
	s.StorageSpaceGB = &v
	return s
}

// SetUpdateTime sets the UpdateTime field's value.
func (s *DataForListDBInstancesOutput) SetUpdateTime(v string) *DataForListDBInstancesOutput {
	s.UpdateTime = &v
	return s
}

// SetVpcID sets the VpcID field's value.
func (s *DataForListDBInstancesOutput) SetVpcID(v string) *DataForListDBInstancesOutput {
	s.VpcID = &v
	return s
}

// SetZone sets the Zone field's value.
func (s *DataForListDBInstancesOutput) SetZone(v string) *DataForListDBInstancesOutput {
	s.Zone = &v
	return s
}

type InstanceSpecForListDBInstancesOutput struct {
	_ struct{} `type:"structure"`

	CpuNum *float64 `type:"double"`

	MemInGb *float64 `type:"double"`

	SpecName *string `type:"string"`
}

// String returns the string representation
func (s InstanceSpecForListDBInstancesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s InstanceSpecForListDBInstancesOutput) GoString() string {
	return s.String()
}

// SetCpuNum sets the CpuNum field's value.
func (s *InstanceSpecForListDBInstancesOutput) SetCpuNum(v float64) *InstanceSpecForListDBInstancesOutput {
	s.CpuNum = &v
	return s
}

// SetMemInGb sets the MemInGb field's value.
func (s *InstanceSpecForListDBInstancesOutput) SetMemInGb(v float64) *InstanceSpecForListDBInstancesOutput {
	s.MemInGb = &v
	return s
}

// SetSpecName sets the SpecName field's value.
func (s *InstanceSpecForListDBInstancesOutput) SetSpecName(v string) *InstanceSpecForListDBInstancesOutput {
	s.SpecName = &v
	return s
}

type ListDBInstancesInput struct {
	_ struct{} `type:"structure"`

	CreateEndTime *string `type:"string"`

	CreateStartTime *string `type:"string"`

	InstanceId *string `type:"string"`

	InstanceStatus *string `type:"string" enum:"EnumOfInstanceStatusForListDBInstancesInput"`

	InstanceType *string `type:"string" enum:"EnumOfInstanceTypeForListDBInstancesInput"`

	Limit *int32 `type:"int32"`

	Offset *int32 `type:"int32"`

	Region *string `type:"string"`

	Zone *string `type:"string"`
}

// String returns the string representation
func (s ListDBInstancesInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListDBInstancesInput) GoString() string {
	return s.String()
}

// SetCreateEndTime sets the CreateEndTime field's value.
func (s *ListDBInstancesInput) SetCreateEndTime(v string) *ListDBInstancesInput {
	s.CreateEndTime = &v
	return s
}

// SetCreateStartTime sets the CreateStartTime field's value.
func (s *ListDBInstancesInput) SetCreateStartTime(v string) *ListDBInstancesInput {
	s.CreateStartTime = &v
	return s
}

// SetInstanceId sets the InstanceId field's value.
func (s *ListDBInstancesInput) SetInstanceId(v string) *ListDBInstancesInput {
	s.InstanceId = &v
	return s
}

// SetInstanceStatus sets the InstanceStatus field's value.
func (s *ListDBInstancesInput) SetInstanceStatus(v string) *ListDBInstancesInput {
	s.InstanceStatus = &v
	return s
}

// SetInstanceType sets the InstanceType field's value.
func (s *ListDBInstancesInput) SetInstanceType(v string) *ListDBInstancesInput {
	s.InstanceType = &v
	return s
}

// SetLimit sets the Limit field's value.
func (s *ListDBInstancesInput) SetLimit(v int32) *ListDBInstancesInput {
	s.Limit = &v
	return s
}

// SetOffset sets the Offset field's value.
func (s *ListDBInstancesInput) SetOffset(v int32) *ListDBInstancesInput {
	s.Offset = &v
	return s
}

// SetRegion sets the Region field's value.
func (s *ListDBInstancesInput) SetRegion(v string) *ListDBInstancesInput {
	s.Region = &v
	return s
}

// SetZone sets the Zone field's value.
func (s *ListDBInstancesInput) SetZone(v string) *ListDBInstancesInput {
	s.Zone = &v
	return s
}

type ListDBInstancesOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	Datas []*DataForListDBInstancesOutput `type:"list"`

	Total *int32 `type:"int32"`
}

// String returns the string representation
func (s ListDBInstancesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListDBInstancesOutput) GoString() string {
	return s.String()
}

// SetDatas sets the Datas field's value.
func (s *ListDBInstancesOutput) SetDatas(v []*DataForListDBInstancesOutput) *ListDBInstancesOutput {
	s.Datas = v
	return s
}

// SetTotal sets the Total field's value.
func (s *ListDBInstancesOutput) SetTotal(v int32) *ListDBInstancesOutput {
	s.Total = &v
	return s
}

const (
	// EnumOfChargeStatusForListDBInstancesOutputNormal is a EnumOfChargeStatusForListDBInstancesOutput enum value
	EnumOfChargeStatusForListDBInstancesOutputNormal = "Normal"

	// EnumOfChargeStatusForListDBInstancesOutputOverdue is a EnumOfChargeStatusForListDBInstancesOutput enum value
	EnumOfChargeStatusForListDBInstancesOutputOverdue = "Overdue"

	// EnumOfChargeStatusForListDBInstancesOutputUnpaid is a EnumOfChargeStatusForListDBInstancesOutput enum value
	EnumOfChargeStatusForListDBInstancesOutputUnpaid = "Unpaid"
)

const (
	// EnumOfChargeTypeForListDBInstancesOutputNotEnabled is a EnumOfChargeTypeForListDBInstancesOutput enum value
	EnumOfChargeTypeForListDBInstancesOutputNotEnabled = "NotEnabled"

	// EnumOfChargeTypeForListDBInstancesOutputPostPaid is a EnumOfChargeTypeForListDBInstancesOutput enum value
	EnumOfChargeTypeForListDBInstancesOutputPostPaid = "PostPaid"

	// EnumOfChargeTypeForListDBInstancesOutputPrepaid is a EnumOfChargeTypeForListDBInstancesOutput enum value
	EnumOfChargeTypeForListDBInstancesOutputPrepaid = "Prepaid"
)

const (
	// EnumOfDBEngineForListDBInstancesOutputMySql is a EnumOfDBEngineForListDBInstancesOutput enum value
	EnumOfDBEngineForListDBInstancesOutputMySql = "MySQL"
)

const (
	// EnumOfDBEngineVersionForListDBInstancesOutputMySql80 is a EnumOfDBEngineVersionForListDBInstancesOutput enum value
	EnumOfDBEngineVersionForListDBInstancesOutputMySql80 = "MySQL_8_0"

	// EnumOfDBEngineVersionForListDBInstancesOutputMySqlCommunity57 is a EnumOfDBEngineVersionForListDBInstancesOutput enum value
	EnumOfDBEngineVersionForListDBInstancesOutputMySqlCommunity57 = "MySQL_Community_5_7"
)

const (
	// EnumOfInstanceStatusForListDBInstancesInputAllowListMaintaining is a EnumOfInstanceStatusForListDBInstancesInput enum value
	EnumOfInstanceStatusForListDBInstancesInputAllowListMaintaining = "AllowListMaintaining"

	// EnumOfInstanceStatusForListDBInstancesInputClosed is a EnumOfInstanceStatusForListDBInstancesInput enum value
	EnumOfInstanceStatusForListDBInstancesInputClosed = "Closed"

	// EnumOfInstanceStatusForListDBInstancesInputClosing is a EnumOfInstanceStatusForListDBInstancesInput enum value
	EnumOfInstanceStatusForListDBInstancesInputClosing = "Closing"

	// EnumOfInstanceStatusForListDBInstancesInputCreateFailed is a EnumOfInstanceStatusForListDBInstancesInput enum value
	EnumOfInstanceStatusForListDBInstancesInputCreateFailed = "CreateFailed"

	// EnumOfInstanceStatusForListDBInstancesInputCreating is a EnumOfInstanceStatusForListDBInstancesInput enum value
	EnumOfInstanceStatusForListDBInstancesInputCreating = "Creating"

	// EnumOfInstanceStatusForListDBInstancesInputDeleting is a EnumOfInstanceStatusForListDBInstancesInput enum value
	EnumOfInstanceStatusForListDBInstancesInputDeleting = "Deleting"

	// EnumOfInstanceStatusForListDBInstancesInputDestroyed is a EnumOfInstanceStatusForListDBInstancesInput enum value
	EnumOfInstanceStatusForListDBInstancesInputDestroyed = "Destroyed"

	// EnumOfInstanceStatusForListDBInstancesInputDestroying is a EnumOfInstanceStatusForListDBInstancesInput enum value
	EnumOfInstanceStatusForListDBInstancesInputDestroying = "Destroying"

	// EnumOfInstanceStatusForListDBInstancesInputError is a EnumOfInstanceStatusForListDBInstancesInput enum value
	EnumOfInstanceStatusForListDBInstancesInputError = "Error"

	// EnumOfInstanceStatusForListDBInstancesInputImporting is a EnumOfInstanceStatusForListDBInstancesInput enum value
	EnumOfInstanceStatusForListDBInstancesInputImporting = "Importing"

	// EnumOfInstanceStatusForListDBInstancesInputMaintaining is a EnumOfInstanceStatusForListDBInstancesInput enum value
	EnumOfInstanceStatusForListDBInstancesInputMaintaining = "Maintaining"

	// EnumOfInstanceStatusForListDBInstancesInputMasterChanging is a EnumOfInstanceStatusForListDBInstancesInput enum value
	EnumOfInstanceStatusForListDBInstancesInputMasterChanging = "MasterChanging"

	// EnumOfInstanceStatusForListDBInstancesInputMigrating is a EnumOfInstanceStatusForListDBInstancesInput enum value
	EnumOfInstanceStatusForListDBInstancesInputMigrating = "Migrating"

	// EnumOfInstanceStatusForListDBInstancesInputReclaiming is a EnumOfInstanceStatusForListDBInstancesInput enum value
	EnumOfInstanceStatusForListDBInstancesInputReclaiming = "Reclaiming"

	// EnumOfInstanceStatusForListDBInstancesInputRecycled is a EnumOfInstanceStatusForListDBInstancesInput enum value
	EnumOfInstanceStatusForListDBInstancesInputRecycled = "Recycled"

	// EnumOfInstanceStatusForListDBInstancesInputReleased is a EnumOfInstanceStatusForListDBInstancesInput enum value
	EnumOfInstanceStatusForListDBInstancesInputReleased = "Released"

	// EnumOfInstanceStatusForListDBInstancesInputRestarting is a EnumOfInstanceStatusForListDBInstancesInput enum value
	EnumOfInstanceStatusForListDBInstancesInputRestarting = "Restarting"

	// EnumOfInstanceStatusForListDBInstancesInputRestoring is a EnumOfInstanceStatusForListDBInstancesInput enum value
	EnumOfInstanceStatusForListDBInstancesInputRestoring = "Restoring"

	// EnumOfInstanceStatusForListDBInstancesInputResuming is a EnumOfInstanceStatusForListDBInstancesInput enum value
	EnumOfInstanceStatusForListDBInstancesInputResuming = "Resuming"

	// EnumOfInstanceStatusForListDBInstancesInputRunning is a EnumOfInstanceStatusForListDBInstancesInput enum value
	EnumOfInstanceStatusForListDBInstancesInputRunning = "Running"

	// EnumOfInstanceStatusForListDBInstancesInputSslupdating is a EnumOfInstanceStatusForListDBInstancesInput enum value
	EnumOfInstanceStatusForListDBInstancesInputSslupdating = "SSLUpdating"

	// EnumOfInstanceStatusForListDBInstancesInputTdeupdating is a EnumOfInstanceStatusForListDBInstancesInput enum value
	EnumOfInstanceStatusForListDBInstancesInputTdeupdating = "TDEUpdating"

	// EnumOfInstanceStatusForListDBInstancesInputUnknown is a EnumOfInstanceStatusForListDBInstancesInput enum value
	EnumOfInstanceStatusForListDBInstancesInputUnknown = "Unknown"

	// EnumOfInstanceStatusForListDBInstancesInputUpdating is a EnumOfInstanceStatusForListDBInstancesInput enum value
	EnumOfInstanceStatusForListDBInstancesInputUpdating = "Updating"

	// EnumOfInstanceStatusForListDBInstancesInputUpgrading is a EnumOfInstanceStatusForListDBInstancesInput enum value
	EnumOfInstanceStatusForListDBInstancesInputUpgrading = "Upgrading"

	// EnumOfInstanceStatusForListDBInstancesInputWaitingPaid is a EnumOfInstanceStatusForListDBInstancesInput enum value
	EnumOfInstanceStatusForListDBInstancesInputWaitingPaid = "WaitingPaid"
)

const (
	// EnumOfInstanceStatusForListDBInstancesOutputAllowListMaintaining is a EnumOfInstanceStatusForListDBInstancesOutput enum value
	EnumOfInstanceStatusForListDBInstancesOutputAllowListMaintaining = "AllowListMaintaining"

	// EnumOfInstanceStatusForListDBInstancesOutputClosed is a EnumOfInstanceStatusForListDBInstancesOutput enum value
	EnumOfInstanceStatusForListDBInstancesOutputClosed = "Closed"

	// EnumOfInstanceStatusForListDBInstancesOutputClosing is a EnumOfInstanceStatusForListDBInstancesOutput enum value
	EnumOfInstanceStatusForListDBInstancesOutputClosing = "Closing"

	// EnumOfInstanceStatusForListDBInstancesOutputCreateFailed is a EnumOfInstanceStatusForListDBInstancesOutput enum value
	EnumOfInstanceStatusForListDBInstancesOutputCreateFailed = "CreateFailed"

	// EnumOfInstanceStatusForListDBInstancesOutputCreating is a EnumOfInstanceStatusForListDBInstancesOutput enum value
	EnumOfInstanceStatusForListDBInstancesOutputCreating = "Creating"

	// EnumOfInstanceStatusForListDBInstancesOutputDeleting is a EnumOfInstanceStatusForListDBInstancesOutput enum value
	EnumOfInstanceStatusForListDBInstancesOutputDeleting = "Deleting"

	// EnumOfInstanceStatusForListDBInstancesOutputDestroyed is a EnumOfInstanceStatusForListDBInstancesOutput enum value
	EnumOfInstanceStatusForListDBInstancesOutputDestroyed = "Destroyed"

	// EnumOfInstanceStatusForListDBInstancesOutputDestroying is a EnumOfInstanceStatusForListDBInstancesOutput enum value
	EnumOfInstanceStatusForListDBInstancesOutputDestroying = "Destroying"

	// EnumOfInstanceStatusForListDBInstancesOutputError is a EnumOfInstanceStatusForListDBInstancesOutput enum value
	EnumOfInstanceStatusForListDBInstancesOutputError = "Error"

	// EnumOfInstanceStatusForListDBInstancesOutputImporting is a EnumOfInstanceStatusForListDBInstancesOutput enum value
	EnumOfInstanceStatusForListDBInstancesOutputImporting = "Importing"

	// EnumOfInstanceStatusForListDBInstancesOutputMaintaining is a EnumOfInstanceStatusForListDBInstancesOutput enum value
	EnumOfInstanceStatusForListDBInstancesOutputMaintaining = "Maintaining"

	// EnumOfInstanceStatusForListDBInstancesOutputMasterChanging is a EnumOfInstanceStatusForListDBInstancesOutput enum value
	EnumOfInstanceStatusForListDBInstancesOutputMasterChanging = "MasterChanging"

	// EnumOfInstanceStatusForListDBInstancesOutputMigrating is a EnumOfInstanceStatusForListDBInstancesOutput enum value
	EnumOfInstanceStatusForListDBInstancesOutputMigrating = "Migrating"

	// EnumOfInstanceStatusForListDBInstancesOutputReclaiming is a EnumOfInstanceStatusForListDBInstancesOutput enum value
	EnumOfInstanceStatusForListDBInstancesOutputReclaiming = "Reclaiming"

	// EnumOfInstanceStatusForListDBInstancesOutputRecycled is a EnumOfInstanceStatusForListDBInstancesOutput enum value
	EnumOfInstanceStatusForListDBInstancesOutputRecycled = "Recycled"

	// EnumOfInstanceStatusForListDBInstancesOutputReleased is a EnumOfInstanceStatusForListDBInstancesOutput enum value
	EnumOfInstanceStatusForListDBInstancesOutputReleased = "Released"

	// EnumOfInstanceStatusForListDBInstancesOutputRestarting is a EnumOfInstanceStatusForListDBInstancesOutput enum value
	EnumOfInstanceStatusForListDBInstancesOutputRestarting = "Restarting"

	// EnumOfInstanceStatusForListDBInstancesOutputRestoring is a EnumOfInstanceStatusForListDBInstancesOutput enum value
	EnumOfInstanceStatusForListDBInstancesOutputRestoring = "Restoring"

	// EnumOfInstanceStatusForListDBInstancesOutputResuming is a EnumOfInstanceStatusForListDBInstancesOutput enum value
	EnumOfInstanceStatusForListDBInstancesOutputResuming = "Resuming"

	// EnumOfInstanceStatusForListDBInstancesOutputRunning is a EnumOfInstanceStatusForListDBInstancesOutput enum value
	EnumOfInstanceStatusForListDBInstancesOutputRunning = "Running"

	// EnumOfInstanceStatusForListDBInstancesOutputSslupdating is a EnumOfInstanceStatusForListDBInstancesOutput enum value
	EnumOfInstanceStatusForListDBInstancesOutputSslupdating = "SSLUpdating"

	// EnumOfInstanceStatusForListDBInstancesOutputTdeupdating is a EnumOfInstanceStatusForListDBInstancesOutput enum value
	EnumOfInstanceStatusForListDBInstancesOutputTdeupdating = "TDEUpdating"

	// EnumOfInstanceStatusForListDBInstancesOutputUnknown is a EnumOfInstanceStatusForListDBInstancesOutput enum value
	EnumOfInstanceStatusForListDBInstancesOutputUnknown = "Unknown"

	// EnumOfInstanceStatusForListDBInstancesOutputUpdating is a EnumOfInstanceStatusForListDBInstancesOutput enum value
	EnumOfInstanceStatusForListDBInstancesOutputUpdating = "Updating"

	// EnumOfInstanceStatusForListDBInstancesOutputUpgrading is a EnumOfInstanceStatusForListDBInstancesOutput enum value
	EnumOfInstanceStatusForListDBInstancesOutputUpgrading = "Upgrading"

	// EnumOfInstanceStatusForListDBInstancesOutputWaitingPaid is a EnumOfInstanceStatusForListDBInstancesOutput enum value
	EnumOfInstanceStatusForListDBInstancesOutputWaitingPaid = "WaitingPaid"
)

const (
	// EnumOfInstanceTypeForListDBInstancesInputBasic is a EnumOfInstanceTypeForListDBInstancesInput enum value
	EnumOfInstanceTypeForListDBInstancesInputBasic = "Basic"

	// EnumOfInstanceTypeForListDBInstancesInputCluster is a EnumOfInstanceTypeForListDBInstancesInput enum value
	EnumOfInstanceTypeForListDBInstancesInputCluster = "Cluster"

	// EnumOfInstanceTypeForListDBInstancesInputFinance is a EnumOfInstanceTypeForListDBInstancesInput enum value
	EnumOfInstanceTypeForListDBInstancesInputFinance = "Finance"

	// EnumOfInstanceTypeForListDBInstancesInputHa is a EnumOfInstanceTypeForListDBInstancesInput enum value
	EnumOfInstanceTypeForListDBInstancesInputHa = "HA"
)

const (
	// EnumOfInstanceTypeForListDBInstancesOutputBasic is a EnumOfInstanceTypeForListDBInstancesOutput enum value
	EnumOfInstanceTypeForListDBInstancesOutputBasic = "Basic"

	// EnumOfInstanceTypeForListDBInstancesOutputCluster is a EnumOfInstanceTypeForListDBInstancesOutput enum value
	EnumOfInstanceTypeForListDBInstancesOutputCluster = "Cluster"

	// EnumOfInstanceTypeForListDBInstancesOutputFinance is a EnumOfInstanceTypeForListDBInstancesOutput enum value
	EnumOfInstanceTypeForListDBInstancesOutputFinance = "Finance"

	// EnumOfInstanceTypeForListDBInstancesOutputHa is a EnumOfInstanceTypeForListDBInstancesOutput enum value
	EnumOfInstanceTypeForListDBInstancesOutputHa = "HA"
)

// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package mongodb

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
func (c *MONGODB) DescribeDBInstancesCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
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

// DescribeDBInstancesCommon API operation for MONGODB.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for MONGODB's
// API operation DescribeDBInstancesCommon for usage and error information.
func (c *MONGODB) DescribeDBInstancesCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
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
func (c *MONGODB) DescribeDBInstancesCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
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
func (c *MONGODB) DescribeDBInstancesRequest(input *DescribeDBInstancesInput) (req *request.Request, output *DescribeDBInstancesOutput) {
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

// DescribeDBInstances API operation for MONGODB.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for MONGODB's
// API operation DescribeDBInstances for usage and error information.
func (c *MONGODB) DescribeDBInstances(input *DescribeDBInstancesInput) (*DescribeDBInstancesOutput, error) {
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
func (c *MONGODB) DescribeDBInstancesWithContext(ctx volcengine.Context, input *DescribeDBInstancesInput, opts ...request.Option) (*DescribeDBInstancesOutput, error) {
	req, out := c.DescribeDBInstancesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DBInstanceForDescribeDBInstancesOutput struct {
	_ struct{} `type:"structure"`

	AutoRenew *bool `type:"boolean"`

	ChargeStatus *string `type:"string" enum:"EnumOfChargeStatusForDescribeDBInstancesOutput"`

	ChargeType *string `type:"string" enum:"EnumOfChargeTypeForDescribeDBInstancesOutput"`

	ClosedTime *string `type:"string"`

	ConfigServersId *string `type:"string"`

	CreateTime *string `type:"string"`

	DBEngine *string `type:"string" enum:"EnumOfDBEngineForDescribeDBInstancesOutput"`

	DBEngineVersion *string `type:"string" enum:"EnumOfDBEngineVersionForDescribeDBInstancesOutput"`

	DBEngineVersionStr *string `type:"string"`

	ExpiredTime *string `type:"string"`

	InstanceId *string `type:"string"`

	InstanceName *string `type:"string"`

	InstanceStatus *string `type:"string" enum:"EnumOfInstanceStatusForDescribeDBInstancesOutput"`

	InstanceType *string `type:"string" enum:"EnumOfInstanceTypeForDescribeDBInstancesOutput"`

	MongosId *string `type:"string"`

	ProjectName *string `type:"string"`

	ReclaimTime *string `type:"string"`

	SubnetId *string `type:"string"`

	UpdateTime *string `type:"string"`

	VpcId *string `type:"string"`

	ZoneId *string `type:"string"`
}

// String returns the string representation
func (s DBInstanceForDescribeDBInstancesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DBInstanceForDescribeDBInstancesOutput) GoString() string {
	return s.String()
}

// SetAutoRenew sets the AutoRenew field's value.
func (s *DBInstanceForDescribeDBInstancesOutput) SetAutoRenew(v bool) *DBInstanceForDescribeDBInstancesOutput {
	s.AutoRenew = &v
	return s
}

// SetChargeStatus sets the ChargeStatus field's value.
func (s *DBInstanceForDescribeDBInstancesOutput) SetChargeStatus(v string) *DBInstanceForDescribeDBInstancesOutput {
	s.ChargeStatus = &v
	return s
}

// SetChargeType sets the ChargeType field's value.
func (s *DBInstanceForDescribeDBInstancesOutput) SetChargeType(v string) *DBInstanceForDescribeDBInstancesOutput {
	s.ChargeType = &v
	return s
}

// SetClosedTime sets the ClosedTime field's value.
func (s *DBInstanceForDescribeDBInstancesOutput) SetClosedTime(v string) *DBInstanceForDescribeDBInstancesOutput {
	s.ClosedTime = &v
	return s
}

// SetConfigServersId sets the ConfigServersId field's value.
func (s *DBInstanceForDescribeDBInstancesOutput) SetConfigServersId(v string) *DBInstanceForDescribeDBInstancesOutput {
	s.ConfigServersId = &v
	return s
}

// SetCreateTime sets the CreateTime field's value.
func (s *DBInstanceForDescribeDBInstancesOutput) SetCreateTime(v string) *DBInstanceForDescribeDBInstancesOutput {
	s.CreateTime = &v
	return s
}

// SetDBEngine sets the DBEngine field's value.
func (s *DBInstanceForDescribeDBInstancesOutput) SetDBEngine(v string) *DBInstanceForDescribeDBInstancesOutput {
	s.DBEngine = &v
	return s
}

// SetDBEngineVersion sets the DBEngineVersion field's value.
func (s *DBInstanceForDescribeDBInstancesOutput) SetDBEngineVersion(v string) *DBInstanceForDescribeDBInstancesOutput {
	s.DBEngineVersion = &v
	return s
}

// SetDBEngineVersionStr sets the DBEngineVersionStr field's value.
func (s *DBInstanceForDescribeDBInstancesOutput) SetDBEngineVersionStr(v string) *DBInstanceForDescribeDBInstancesOutput {
	s.DBEngineVersionStr = &v
	return s
}

// SetExpiredTime sets the ExpiredTime field's value.
func (s *DBInstanceForDescribeDBInstancesOutput) SetExpiredTime(v string) *DBInstanceForDescribeDBInstancesOutput {
	s.ExpiredTime = &v
	return s
}

// SetInstanceId sets the InstanceId field's value.
func (s *DBInstanceForDescribeDBInstancesOutput) SetInstanceId(v string) *DBInstanceForDescribeDBInstancesOutput {
	s.InstanceId = &v
	return s
}

// SetInstanceName sets the InstanceName field's value.
func (s *DBInstanceForDescribeDBInstancesOutput) SetInstanceName(v string) *DBInstanceForDescribeDBInstancesOutput {
	s.InstanceName = &v
	return s
}

// SetInstanceStatus sets the InstanceStatus field's value.
func (s *DBInstanceForDescribeDBInstancesOutput) SetInstanceStatus(v string) *DBInstanceForDescribeDBInstancesOutput {
	s.InstanceStatus = &v
	return s
}

// SetInstanceType sets the InstanceType field's value.
func (s *DBInstanceForDescribeDBInstancesOutput) SetInstanceType(v string) *DBInstanceForDescribeDBInstancesOutput {
	s.InstanceType = &v
	return s
}

// SetMongosId sets the MongosId field's value.
func (s *DBInstanceForDescribeDBInstancesOutput) SetMongosId(v string) *DBInstanceForDescribeDBInstancesOutput {
	s.MongosId = &v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *DBInstanceForDescribeDBInstancesOutput) SetProjectName(v string) *DBInstanceForDescribeDBInstancesOutput {
	s.ProjectName = &v
	return s
}

// SetReclaimTime sets the ReclaimTime field's value.
func (s *DBInstanceForDescribeDBInstancesOutput) SetReclaimTime(v string) *DBInstanceForDescribeDBInstancesOutput {
	s.ReclaimTime = &v
	return s
}

// SetSubnetId sets the SubnetId field's value.
func (s *DBInstanceForDescribeDBInstancesOutput) SetSubnetId(v string) *DBInstanceForDescribeDBInstancesOutput {
	s.SubnetId = &v
	return s
}

// SetUpdateTime sets the UpdateTime field's value.
func (s *DBInstanceForDescribeDBInstancesOutput) SetUpdateTime(v string) *DBInstanceForDescribeDBInstancesOutput {
	s.UpdateTime = &v
	return s
}

// SetVpcId sets the VpcId field's value.
func (s *DBInstanceForDescribeDBInstancesOutput) SetVpcId(v string) *DBInstanceForDescribeDBInstancesOutput {
	s.VpcId = &v
	return s
}

// SetZoneId sets the ZoneId field's value.
func (s *DBInstanceForDescribeDBInstancesOutput) SetZoneId(v string) *DBInstanceForDescribeDBInstancesOutput {
	s.ZoneId = &v
	return s
}

type DescribeDBInstancesInput struct {
	_ struct{} `type:"structure"`

	CreateEndTime *string `type:"string"`

	CreateStartTime *string `type:"string"`

	DBEngine *string `type:"string" enum:"EnumOfDBEngineForDescribeDBInstancesInput"`

	DBEngineVersion *string `type:"string" enum:"EnumOfDBEngineVersionForDescribeDBInstancesInput"`

	InstanceId *string `type:"string"`

	InstanceName *string `max:"64" type:"string"`

	InstanceStatus *string `type:"string" enum:"EnumOfInstanceStatusForDescribeDBInstancesInput"`

	InstanceType *string `type:"string" enum:"EnumOfInstanceTypeForDescribeDBInstancesInput"`

	PageNumber *int32 `type:"int32"`

	PageSize *int32 `type:"int32"`

	UpdateEndTime *string `type:"string"`

	UpdateStartTime *string `type:"string"`

	VpcId *string `type:"string"`

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

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeDBInstancesInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DescribeDBInstancesInput"}
	if s.InstanceName != nil && len(*s.InstanceName) > 64 {
		invalidParams.Add(request.NewErrParamMaxLen("InstanceName", 64, *s.InstanceName))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetCreateEndTime sets the CreateEndTime field's value.
func (s *DescribeDBInstancesInput) SetCreateEndTime(v string) *DescribeDBInstancesInput {
	s.CreateEndTime = &v
	return s
}

// SetCreateStartTime sets the CreateStartTime field's value.
func (s *DescribeDBInstancesInput) SetCreateStartTime(v string) *DescribeDBInstancesInput {
	s.CreateStartTime = &v
	return s
}

// SetDBEngine sets the DBEngine field's value.
func (s *DescribeDBInstancesInput) SetDBEngine(v string) *DescribeDBInstancesInput {
	s.DBEngine = &v
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

// SetUpdateEndTime sets the UpdateEndTime field's value.
func (s *DescribeDBInstancesInput) SetUpdateEndTime(v string) *DescribeDBInstancesInput {
	s.UpdateEndTime = &v
	return s
}

// SetUpdateStartTime sets the UpdateStartTime field's value.
func (s *DescribeDBInstancesInput) SetUpdateStartTime(v string) *DescribeDBInstancesInput {
	s.UpdateStartTime = &v
	return s
}

// SetVpcId sets the VpcId field's value.
func (s *DescribeDBInstancesInput) SetVpcId(v string) *DescribeDBInstancesInput {
	s.VpcId = &v
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

	DBInstances []*DBInstanceForDescribeDBInstancesOutput `type:"list"`

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

// SetDBInstances sets the DBInstances field's value.
func (s *DescribeDBInstancesOutput) SetDBInstances(v []*DBInstanceForDescribeDBInstancesOutput) *DescribeDBInstancesOutput {
	s.DBInstances = v
	return s
}

// SetTotal sets the Total field's value.
func (s *DescribeDBInstancesOutput) SetTotal(v int32) *DescribeDBInstancesOutput {
	s.Total = &v
	return s
}

const (
	// EnumOfChargeStatusForDescribeDBInstancesOutputNormal is a EnumOfChargeStatusForDescribeDBInstancesOutput enum value
	EnumOfChargeStatusForDescribeDBInstancesOutputNormal = "Normal"

	// EnumOfChargeStatusForDescribeDBInstancesOutputOverdue is a EnumOfChargeStatusForDescribeDBInstancesOutput enum value
	EnumOfChargeStatusForDescribeDBInstancesOutputOverdue = "Overdue"

	// EnumOfChargeStatusForDescribeDBInstancesOutputOwing is a EnumOfChargeStatusForDescribeDBInstancesOutput enum value
	EnumOfChargeStatusForDescribeDBInstancesOutputOwing = "Owing"

	// EnumOfChargeStatusForDescribeDBInstancesOutputRenewing is a EnumOfChargeStatusForDescribeDBInstancesOutput enum value
	EnumOfChargeStatusForDescribeDBInstancesOutputRenewing = "Renewing"

	// EnumOfChargeStatusForDescribeDBInstancesOutputUnDeploy is a EnumOfChargeStatusForDescribeDBInstancesOutput enum value
	EnumOfChargeStatusForDescribeDBInstancesOutputUnDeploy = "UnDeploy"

	// EnumOfChargeStatusForDescribeDBInstancesOutputUnsubscribing is a EnumOfChargeStatusForDescribeDBInstancesOutput enum value
	EnumOfChargeStatusForDescribeDBInstancesOutputUnsubscribing = "Unsubscribing"

	// EnumOfChargeStatusForDescribeDBInstancesOutputWaitingPaid is a EnumOfChargeStatusForDescribeDBInstancesOutput enum value
	EnumOfChargeStatusForDescribeDBInstancesOutputWaitingPaid = "WaitingPaid"
)

const (
	// EnumOfChargeTypeForDescribeDBInstancesOutputNotEnabled is a EnumOfChargeTypeForDescribeDBInstancesOutput enum value
	EnumOfChargeTypeForDescribeDBInstancesOutputNotEnabled = "NotEnabled"

	// EnumOfChargeTypeForDescribeDBInstancesOutputPostPaid is a EnumOfChargeTypeForDescribeDBInstancesOutput enum value
	EnumOfChargeTypeForDescribeDBInstancesOutputPostPaid = "PostPaid"

	// EnumOfChargeTypeForDescribeDBInstancesOutputPrepaid is a EnumOfChargeTypeForDescribeDBInstancesOutput enum value
	EnumOfChargeTypeForDescribeDBInstancesOutputPrepaid = "Prepaid"
)

const (
	// EnumOfDBEngineForDescribeDBInstancesInputMongoDb is a EnumOfDBEngineForDescribeDBInstancesInput enum value
	EnumOfDBEngineForDescribeDBInstancesInputMongoDb = "MongoDB"
)

const (
	// EnumOfDBEngineForDescribeDBInstancesOutputMongoDb is a EnumOfDBEngineForDescribeDBInstancesOutput enum value
	EnumOfDBEngineForDescribeDBInstancesOutputMongoDb = "MongoDB"
)

const (
	// EnumOfDBEngineVersionForDescribeDBInstancesInputMongoDb40 is a EnumOfDBEngineVersionForDescribeDBInstancesInput enum value
	EnumOfDBEngineVersionForDescribeDBInstancesInputMongoDb40 = "MongoDB_4_0"

	// EnumOfDBEngineVersionForDescribeDBInstancesInputMongoDb50 is a EnumOfDBEngineVersionForDescribeDBInstancesInput enum value
	EnumOfDBEngineVersionForDescribeDBInstancesInputMongoDb50 = "MongoDB_5_0"
)

const (
	// EnumOfDBEngineVersionForDescribeDBInstancesOutputMongoDb40 is a EnumOfDBEngineVersionForDescribeDBInstancesOutput enum value
	EnumOfDBEngineVersionForDescribeDBInstancesOutputMongoDb40 = "MongoDB_4_0"

	// EnumOfDBEngineVersionForDescribeDBInstancesOutputMongoDb50 is a EnumOfDBEngineVersionForDescribeDBInstancesOutput enum value
	EnumOfDBEngineVersionForDescribeDBInstancesOutputMongoDb50 = "MongoDB_5_0"
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

	// EnumOfInstanceStatusForDescribeDBInstancesInputDeleted is a EnumOfInstanceStatusForDescribeDBInstancesInput enum value
	EnumOfInstanceStatusForDescribeDBInstancesInputDeleted = "Deleted"

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

	// EnumOfInstanceStatusForDescribeDBInstancesInputMigrating is a EnumOfInstanceStatusForDescribeDBInstancesInput enum value
	EnumOfInstanceStatusForDescribeDBInstancesInputMigrating = "Migrating"

	// EnumOfInstanceStatusForDescribeDBInstancesInputNetCreating is a EnumOfInstanceStatusForDescribeDBInstancesInput enum value
	EnumOfInstanceStatusForDescribeDBInstancesInputNetCreating = "NetCreating"

	// EnumOfInstanceStatusForDescribeDBInstancesInputNetReleasing is a EnumOfInstanceStatusForDescribeDBInstancesInput enum value
	EnumOfInstanceStatusForDescribeDBInstancesInputNetReleasing = "NetReleasing"

	// EnumOfInstanceStatusForDescribeDBInstancesInputNetworkMaintaining is a EnumOfInstanceStatusForDescribeDBInstancesInput enum value
	EnumOfInstanceStatusForDescribeDBInstancesInputNetworkMaintaining = "NetworkMaintaining"

	// EnumOfInstanceStatusForDescribeDBInstancesInputRebuilding is a EnumOfInstanceStatusForDescribeDBInstancesInput enum value
	EnumOfInstanceStatusForDescribeDBInstancesInputRebuilding = "Rebuilding"

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

	// EnumOfInstanceStatusForDescribeDBInstancesInputScaling is a EnumOfInstanceStatusForDescribeDBInstancesInput enum value
	EnumOfInstanceStatusForDescribeDBInstancesInputScaling = "Scaling"

	// EnumOfInstanceStatusForDescribeDBInstancesInputTdemaintaining is a EnumOfInstanceStatusForDescribeDBInstancesInput enum value
	EnumOfInstanceStatusForDescribeDBInstancesInputTdemaintaining = "TDEMaintaining"

	// EnumOfInstanceStatusForDescribeDBInstancesInputTaskFailedAvailable is a EnumOfInstanceStatusForDescribeDBInstancesInput enum value
	EnumOfInstanceStatusForDescribeDBInstancesInputTaskFailedAvailable = "TaskFailed_Available"

	// EnumOfInstanceStatusForDescribeDBInstancesInputUnavailable is a EnumOfInstanceStatusForDescribeDBInstancesInput enum value
	EnumOfInstanceStatusForDescribeDBInstancesInputUnavailable = "Unavailable"

	// EnumOfInstanceStatusForDescribeDBInstancesInputUpdating is a EnumOfInstanceStatusForDescribeDBInstancesInput enum value
	EnumOfInstanceStatusForDescribeDBInstancesInputUpdating = "Updating"

	// EnumOfInstanceStatusForDescribeDBInstancesInputUpgrading is a EnumOfInstanceStatusForDescribeDBInstancesInput enum value
	EnumOfInstanceStatusForDescribeDBInstancesInputUpgrading = "Upgrading"

	// EnumOfInstanceStatusForDescribeDBInstancesInputWaitingPaid is a EnumOfInstanceStatusForDescribeDBInstancesInput enum value
	EnumOfInstanceStatusForDescribeDBInstancesInputWaitingPaid = "WaitingPaid"
)

const (
	// EnumOfInstanceStatusForDescribeDBInstancesOutputAllowListMaintaining is a EnumOfInstanceStatusForDescribeDBInstancesOutput enum value
	EnumOfInstanceStatusForDescribeDBInstancesOutputAllowListMaintaining = "AllowListMaintaining"

	// EnumOfInstanceStatusForDescribeDBInstancesOutputClosed is a EnumOfInstanceStatusForDescribeDBInstancesOutput enum value
	EnumOfInstanceStatusForDescribeDBInstancesOutputClosed = "Closed"

	// EnumOfInstanceStatusForDescribeDBInstancesOutputClosing is a EnumOfInstanceStatusForDescribeDBInstancesOutput enum value
	EnumOfInstanceStatusForDescribeDBInstancesOutputClosing = "Closing"

	// EnumOfInstanceStatusForDescribeDBInstancesOutputCreateFailed is a EnumOfInstanceStatusForDescribeDBInstancesOutput enum value
	EnumOfInstanceStatusForDescribeDBInstancesOutputCreateFailed = "CreateFailed"

	// EnumOfInstanceStatusForDescribeDBInstancesOutputCreating is a EnumOfInstanceStatusForDescribeDBInstancesOutput enum value
	EnumOfInstanceStatusForDescribeDBInstancesOutputCreating = "Creating"

	// EnumOfInstanceStatusForDescribeDBInstancesOutputDeleted is a EnumOfInstanceStatusForDescribeDBInstancesOutput enum value
	EnumOfInstanceStatusForDescribeDBInstancesOutputDeleted = "Deleted"

	// EnumOfInstanceStatusForDescribeDBInstancesOutputDeleting is a EnumOfInstanceStatusForDescribeDBInstancesOutput enum value
	EnumOfInstanceStatusForDescribeDBInstancesOutputDeleting = "Deleting"

	// EnumOfInstanceStatusForDescribeDBInstancesOutputDestroyed is a EnumOfInstanceStatusForDescribeDBInstancesOutput enum value
	EnumOfInstanceStatusForDescribeDBInstancesOutputDestroyed = "Destroyed"

	// EnumOfInstanceStatusForDescribeDBInstancesOutputDestroying is a EnumOfInstanceStatusForDescribeDBInstancesOutput enum value
	EnumOfInstanceStatusForDescribeDBInstancesOutputDestroying = "Destroying"

	// EnumOfInstanceStatusForDescribeDBInstancesOutputError is a EnumOfInstanceStatusForDescribeDBInstancesOutput enum value
	EnumOfInstanceStatusForDescribeDBInstancesOutputError = "Error"

	// EnumOfInstanceStatusForDescribeDBInstancesOutputImporting is a EnumOfInstanceStatusForDescribeDBInstancesOutput enum value
	EnumOfInstanceStatusForDescribeDBInstancesOutputImporting = "Importing"

	// EnumOfInstanceStatusForDescribeDBInstancesOutputMigrating is a EnumOfInstanceStatusForDescribeDBInstancesOutput enum value
	EnumOfInstanceStatusForDescribeDBInstancesOutputMigrating = "Migrating"

	// EnumOfInstanceStatusForDescribeDBInstancesOutputNetCreating is a EnumOfInstanceStatusForDescribeDBInstancesOutput enum value
	EnumOfInstanceStatusForDescribeDBInstancesOutputNetCreating = "NetCreating"

	// EnumOfInstanceStatusForDescribeDBInstancesOutputNetReleasing is a EnumOfInstanceStatusForDescribeDBInstancesOutput enum value
	EnumOfInstanceStatusForDescribeDBInstancesOutputNetReleasing = "NetReleasing"

	// EnumOfInstanceStatusForDescribeDBInstancesOutputNetworkMaintaining is a EnumOfInstanceStatusForDescribeDBInstancesOutput enum value
	EnumOfInstanceStatusForDescribeDBInstancesOutputNetworkMaintaining = "NetworkMaintaining"

	// EnumOfInstanceStatusForDescribeDBInstancesOutputRebuilding is a EnumOfInstanceStatusForDescribeDBInstancesOutput enum value
	EnumOfInstanceStatusForDescribeDBInstancesOutputRebuilding = "Rebuilding"

	// EnumOfInstanceStatusForDescribeDBInstancesOutputReclaiming is a EnumOfInstanceStatusForDescribeDBInstancesOutput enum value
	EnumOfInstanceStatusForDescribeDBInstancesOutputReclaiming = "Reclaiming"

	// EnumOfInstanceStatusForDescribeDBInstancesOutputRecycled is a EnumOfInstanceStatusForDescribeDBInstancesOutput enum value
	EnumOfInstanceStatusForDescribeDBInstancesOutputRecycled = "Recycled"

	// EnumOfInstanceStatusForDescribeDBInstancesOutputReleased is a EnumOfInstanceStatusForDescribeDBInstancesOutput enum value
	EnumOfInstanceStatusForDescribeDBInstancesOutputReleased = "Released"

	// EnumOfInstanceStatusForDescribeDBInstancesOutputRestarting is a EnumOfInstanceStatusForDescribeDBInstancesOutput enum value
	EnumOfInstanceStatusForDescribeDBInstancesOutputRestarting = "Restarting"

	// EnumOfInstanceStatusForDescribeDBInstancesOutputRestoring is a EnumOfInstanceStatusForDescribeDBInstancesOutput enum value
	EnumOfInstanceStatusForDescribeDBInstancesOutputRestoring = "Restoring"

	// EnumOfInstanceStatusForDescribeDBInstancesOutputResuming is a EnumOfInstanceStatusForDescribeDBInstancesOutput enum value
	EnumOfInstanceStatusForDescribeDBInstancesOutputResuming = "Resuming"

	// EnumOfInstanceStatusForDescribeDBInstancesOutputRunning is a EnumOfInstanceStatusForDescribeDBInstancesOutput enum value
	EnumOfInstanceStatusForDescribeDBInstancesOutputRunning = "Running"

	// EnumOfInstanceStatusForDescribeDBInstancesOutputSslupdating is a EnumOfInstanceStatusForDescribeDBInstancesOutput enum value
	EnumOfInstanceStatusForDescribeDBInstancesOutputSslupdating = "SSLUpdating"

	// EnumOfInstanceStatusForDescribeDBInstancesOutputScaling is a EnumOfInstanceStatusForDescribeDBInstancesOutput enum value
	EnumOfInstanceStatusForDescribeDBInstancesOutputScaling = "Scaling"

	// EnumOfInstanceStatusForDescribeDBInstancesOutputTdemaintaining is a EnumOfInstanceStatusForDescribeDBInstancesOutput enum value
	EnumOfInstanceStatusForDescribeDBInstancesOutputTdemaintaining = "TDEMaintaining"

	// EnumOfInstanceStatusForDescribeDBInstancesOutputTaskFailedAvailable is a EnumOfInstanceStatusForDescribeDBInstancesOutput enum value
	EnumOfInstanceStatusForDescribeDBInstancesOutputTaskFailedAvailable = "TaskFailed_Available"

	// EnumOfInstanceStatusForDescribeDBInstancesOutputUnavailable is a EnumOfInstanceStatusForDescribeDBInstancesOutput enum value
	EnumOfInstanceStatusForDescribeDBInstancesOutputUnavailable = "Unavailable"

	// EnumOfInstanceStatusForDescribeDBInstancesOutputUpdating is a EnumOfInstanceStatusForDescribeDBInstancesOutput enum value
	EnumOfInstanceStatusForDescribeDBInstancesOutputUpdating = "Updating"

	// EnumOfInstanceStatusForDescribeDBInstancesOutputUpgrading is a EnumOfInstanceStatusForDescribeDBInstancesOutput enum value
	EnumOfInstanceStatusForDescribeDBInstancesOutputUpgrading = "Upgrading"

	// EnumOfInstanceStatusForDescribeDBInstancesOutputWaitingPaid is a EnumOfInstanceStatusForDescribeDBInstancesOutput enum value
	EnumOfInstanceStatusForDescribeDBInstancesOutputWaitingPaid = "WaitingPaid"
)

const (
	// EnumOfInstanceTypeForDescribeDBInstancesInputReplicaSet is a EnumOfInstanceTypeForDescribeDBInstancesInput enum value
	EnumOfInstanceTypeForDescribeDBInstancesInputReplicaSet = "ReplicaSet"

	// EnumOfInstanceTypeForDescribeDBInstancesInputShardedCluster is a EnumOfInstanceTypeForDescribeDBInstancesInput enum value
	EnumOfInstanceTypeForDescribeDBInstancesInputShardedCluster = "ShardedCluster"
)

const (
	// EnumOfInstanceTypeForDescribeDBInstancesOutputReplicaSet is a EnumOfInstanceTypeForDescribeDBInstancesOutput enum value
	EnumOfInstanceTypeForDescribeDBInstancesOutputReplicaSet = "ReplicaSet"

	// EnumOfInstanceTypeForDescribeDBInstancesOutputShardedCluster is a EnumOfInstanceTypeForDescribeDBInstancesOutput enum value
	EnumOfInstanceTypeForDescribeDBInstancesOutputShardedCluster = "ShardedCluster"
)

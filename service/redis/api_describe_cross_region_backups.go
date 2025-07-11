// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package redis

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribeCrossRegionBackupsCommon = "DescribeCrossRegionBackups"

// DescribeCrossRegionBackupsCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeCrossRegionBackupsCommon operation. The "output" return
// value will be populated with the DescribeCrossRegionBackupsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeCrossRegionBackupsCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeCrossRegionBackupsCommon Send returns without error.
//
// See DescribeCrossRegionBackupsCommon for more information on using the DescribeCrossRegionBackupsCommon
// API call, and error handling.
//
//    // Example sending a request using the DescribeCrossRegionBackupsCommonRequest method.
//    req, resp := client.DescribeCrossRegionBackupsCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *REDIS) DescribeCrossRegionBackupsCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeCrossRegionBackupsCommon,
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

// DescribeCrossRegionBackupsCommon API operation for REDIS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for REDIS's
// API operation DescribeCrossRegionBackupsCommon for usage and error information.
func (c *REDIS) DescribeCrossRegionBackupsCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeCrossRegionBackupsCommonRequest(input)
	return out, req.Send()
}

// DescribeCrossRegionBackupsCommonWithContext is the same as DescribeCrossRegionBackupsCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeCrossRegionBackupsCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *REDIS) DescribeCrossRegionBackupsCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeCrossRegionBackupsCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeCrossRegionBackups = "DescribeCrossRegionBackups"

// DescribeCrossRegionBackupsRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeCrossRegionBackups operation. The "output" return
// value will be populated with the DescribeCrossRegionBackupsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeCrossRegionBackupsCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeCrossRegionBackupsCommon Send returns without error.
//
// See DescribeCrossRegionBackups for more information on using the DescribeCrossRegionBackups
// API call, and error handling.
//
//    // Example sending a request using the DescribeCrossRegionBackupsRequest method.
//    req, resp := client.DescribeCrossRegionBackupsRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *REDIS) DescribeCrossRegionBackupsRequest(input *DescribeCrossRegionBackupsInput) (req *request.Request, output *DescribeCrossRegionBackupsOutput) {
	op := &request.Operation{
		Name:       opDescribeCrossRegionBackups,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeCrossRegionBackupsInput{}
	}

	output = &DescribeCrossRegionBackupsOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// DescribeCrossRegionBackups API operation for REDIS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for REDIS's
// API operation DescribeCrossRegionBackups for usage and error information.
func (c *REDIS) DescribeCrossRegionBackups(input *DescribeCrossRegionBackupsInput) (*DescribeCrossRegionBackupsOutput, error) {
	req, out := c.DescribeCrossRegionBackupsRequest(input)
	return out, req.Send()
}

// DescribeCrossRegionBackupsWithContext is the same as DescribeCrossRegionBackups with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeCrossRegionBackups for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *REDIS) DescribeCrossRegionBackupsWithContext(ctx volcengine.Context, input *DescribeCrossRegionBackupsInput, opts ...request.Option) (*DescribeCrossRegionBackupsOutput, error) {
	req, out := c.DescribeCrossRegionBackupsRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type BackupForDescribeCrossRegionBackupsOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	BackupPointId *string `type:"string" json:",omitempty"`

	BackupPointName *string `type:"string" json:",omitempty"`

	BackupStrategy *string `type:"string" json:",omitempty" enum:"EnumOfBackupStrategyForDescribeCrossRegionBackupsOutput"`

	BackupType *string `type:"string" json:",omitempty" enum:"EnumOfBackupTypeForDescribeCrossRegionBackupsOutput"`

	EndTime *string `type:"string" json:",omitempty"`

	InstanceId *string `type:"string" json:",omitempty"`

	InstanceInfo *InstanceInfoForDescribeCrossRegionBackupsOutput `type:"structure" json:",omitempty"`

	ProjectName *string `type:"string" json:",omitempty"`

	Size *int64 `type:"int64" json:",omitempty"`

	SourceRegion *string `type:"string" json:",omitempty"`

	StartTime *string `type:"string" json:",omitempty"`

	Status *string `type:"string" json:",omitempty" enum:"EnumOfStatusForDescribeCrossRegionBackupsOutput"`

	TTL *int32 `type:"int32" json:",omitempty"`

	TargetRegion *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s BackupForDescribeCrossRegionBackupsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s BackupForDescribeCrossRegionBackupsOutput) GoString() string {
	return s.String()
}

// SetBackupPointId sets the BackupPointId field's value.
func (s *BackupForDescribeCrossRegionBackupsOutput) SetBackupPointId(v string) *BackupForDescribeCrossRegionBackupsOutput {
	s.BackupPointId = &v
	return s
}

// SetBackupPointName sets the BackupPointName field's value.
func (s *BackupForDescribeCrossRegionBackupsOutput) SetBackupPointName(v string) *BackupForDescribeCrossRegionBackupsOutput {
	s.BackupPointName = &v
	return s
}

// SetBackupStrategy sets the BackupStrategy field's value.
func (s *BackupForDescribeCrossRegionBackupsOutput) SetBackupStrategy(v string) *BackupForDescribeCrossRegionBackupsOutput {
	s.BackupStrategy = &v
	return s
}

// SetBackupType sets the BackupType field's value.
func (s *BackupForDescribeCrossRegionBackupsOutput) SetBackupType(v string) *BackupForDescribeCrossRegionBackupsOutput {
	s.BackupType = &v
	return s
}

// SetEndTime sets the EndTime field's value.
func (s *BackupForDescribeCrossRegionBackupsOutput) SetEndTime(v string) *BackupForDescribeCrossRegionBackupsOutput {
	s.EndTime = &v
	return s
}

// SetInstanceId sets the InstanceId field's value.
func (s *BackupForDescribeCrossRegionBackupsOutput) SetInstanceId(v string) *BackupForDescribeCrossRegionBackupsOutput {
	s.InstanceId = &v
	return s
}

// SetInstanceInfo sets the InstanceInfo field's value.
func (s *BackupForDescribeCrossRegionBackupsOutput) SetInstanceInfo(v *InstanceInfoForDescribeCrossRegionBackupsOutput) *BackupForDescribeCrossRegionBackupsOutput {
	s.InstanceInfo = v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *BackupForDescribeCrossRegionBackupsOutput) SetProjectName(v string) *BackupForDescribeCrossRegionBackupsOutput {
	s.ProjectName = &v
	return s
}

// SetSize sets the Size field's value.
func (s *BackupForDescribeCrossRegionBackupsOutput) SetSize(v int64) *BackupForDescribeCrossRegionBackupsOutput {
	s.Size = &v
	return s
}

// SetSourceRegion sets the SourceRegion field's value.
func (s *BackupForDescribeCrossRegionBackupsOutput) SetSourceRegion(v string) *BackupForDescribeCrossRegionBackupsOutput {
	s.SourceRegion = &v
	return s
}

// SetStartTime sets the StartTime field's value.
func (s *BackupForDescribeCrossRegionBackupsOutput) SetStartTime(v string) *BackupForDescribeCrossRegionBackupsOutput {
	s.StartTime = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *BackupForDescribeCrossRegionBackupsOutput) SetStatus(v string) *BackupForDescribeCrossRegionBackupsOutput {
	s.Status = &v
	return s
}

// SetTTL sets the TTL field's value.
func (s *BackupForDescribeCrossRegionBackupsOutput) SetTTL(v int32) *BackupForDescribeCrossRegionBackupsOutput {
	s.TTL = &v
	return s
}

// SetTargetRegion sets the TargetRegion field's value.
func (s *BackupForDescribeCrossRegionBackupsOutput) SetTargetRegion(v string) *BackupForDescribeCrossRegionBackupsOutput {
	s.TargetRegion = &v
	return s
}

type DescribeCrossRegionBackupsInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	BackupPointId *string `type:"string" json:",omitempty"`

	BackupPointName *string `type:"string" json:",omitempty"`

	BackupStrategyList []*string `type:"list" json:",omitempty"`

	EndTime *string `type:"string" json:",omitempty"`

	InstanceId *string `type:"string" json:",omitempty"`

	PageNumber *int32 `type:"int32" json:",omitempty"`

	PageSize *int32 `type:"int32" json:",omitempty"`

	ProjectName *string `type:"string" json:",omitempty"`

	Scope *string `type:"string" json:",omitempty" enum:"EnumOfScopeForDescribeCrossRegionBackupsInput"`

	StartTime *string `type:"string" json:",omitempty"`

	Status *string `type:"string" json:",omitempty" enum:"EnumOfStatusForDescribeCrossRegionBackupsInput"`
}

// String returns the string representation
func (s DescribeCrossRegionBackupsInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeCrossRegionBackupsInput) GoString() string {
	return s.String()
}

// SetBackupPointId sets the BackupPointId field's value.
func (s *DescribeCrossRegionBackupsInput) SetBackupPointId(v string) *DescribeCrossRegionBackupsInput {
	s.BackupPointId = &v
	return s
}

// SetBackupPointName sets the BackupPointName field's value.
func (s *DescribeCrossRegionBackupsInput) SetBackupPointName(v string) *DescribeCrossRegionBackupsInput {
	s.BackupPointName = &v
	return s
}

// SetBackupStrategyList sets the BackupStrategyList field's value.
func (s *DescribeCrossRegionBackupsInput) SetBackupStrategyList(v []*string) *DescribeCrossRegionBackupsInput {
	s.BackupStrategyList = v
	return s
}

// SetEndTime sets the EndTime field's value.
func (s *DescribeCrossRegionBackupsInput) SetEndTime(v string) *DescribeCrossRegionBackupsInput {
	s.EndTime = &v
	return s
}

// SetInstanceId sets the InstanceId field's value.
func (s *DescribeCrossRegionBackupsInput) SetInstanceId(v string) *DescribeCrossRegionBackupsInput {
	s.InstanceId = &v
	return s
}

// SetPageNumber sets the PageNumber field's value.
func (s *DescribeCrossRegionBackupsInput) SetPageNumber(v int32) *DescribeCrossRegionBackupsInput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *DescribeCrossRegionBackupsInput) SetPageSize(v int32) *DescribeCrossRegionBackupsInput {
	s.PageSize = &v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *DescribeCrossRegionBackupsInput) SetProjectName(v string) *DescribeCrossRegionBackupsInput {
	s.ProjectName = &v
	return s
}

// SetScope sets the Scope field's value.
func (s *DescribeCrossRegionBackupsInput) SetScope(v string) *DescribeCrossRegionBackupsInput {
	s.Scope = &v
	return s
}

// SetStartTime sets the StartTime field's value.
func (s *DescribeCrossRegionBackupsInput) SetStartTime(v string) *DescribeCrossRegionBackupsInput {
	s.StartTime = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *DescribeCrossRegionBackupsInput) SetStatus(v string) *DescribeCrossRegionBackupsInput {
	s.Status = &v
	return s
}

type DescribeCrossRegionBackupsOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	Backups []*BackupForDescribeCrossRegionBackupsOutput `type:"list" json:",omitempty"`

	Total *int32 `type:"int32" json:",omitempty"`
}

// String returns the string representation
func (s DescribeCrossRegionBackupsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeCrossRegionBackupsOutput) GoString() string {
	return s.String()
}

// SetBackups sets the Backups field's value.
func (s *DescribeCrossRegionBackupsOutput) SetBackups(v []*BackupForDescribeCrossRegionBackupsOutput) *DescribeCrossRegionBackupsOutput {
	s.Backups = v
	return s
}

// SetTotal sets the Total field's value.
func (s *DescribeCrossRegionBackupsOutput) SetTotal(v int32) *DescribeCrossRegionBackupsOutput {
	s.Total = &v
	return s
}

type InstanceInfoForDescribeCrossRegionBackupsOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	AccountId *int64 `type:"int64" json:",omitempty"`

	ArchType *string `type:"string" json:",omitempty"`

	ChargeType *string `type:"string" json:",omitempty"`

	DeletionProtection *string `type:"string" json:",omitempty"`

	EngineVersion *string `type:"string" json:",omitempty"`

	ExpiredTime *string `type:"string" json:",omitempty"`

	InstanceId *string `type:"string" json:",omitempty"`

	InstanceName *string `type:"string" json:",omitempty"`

	MaintenanceTime *string `type:"string" json:",omitempty"`

	NetworkType *string `type:"string" json:",omitempty"`

	RegionId *string `type:"string" json:",omitempty"`

	Replicas *int32 `type:"int32" json:",omitempty"`

	ShardCapacity *float64 `type:"double" json:",omitempty"`

	ShardNumber *int32 `type:"int32" json:",omitempty"`

	TotalCapacity *int64 `type:"int64" json:",omitempty"`

	VpcId *string `type:"string" json:",omitempty"`

	ZoneIds []*string `type:"list" json:",omitempty"`
}

// String returns the string representation
func (s InstanceInfoForDescribeCrossRegionBackupsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s InstanceInfoForDescribeCrossRegionBackupsOutput) GoString() string {
	return s.String()
}

// SetAccountId sets the AccountId field's value.
func (s *InstanceInfoForDescribeCrossRegionBackupsOutput) SetAccountId(v int64) *InstanceInfoForDescribeCrossRegionBackupsOutput {
	s.AccountId = &v
	return s
}

// SetArchType sets the ArchType field's value.
func (s *InstanceInfoForDescribeCrossRegionBackupsOutput) SetArchType(v string) *InstanceInfoForDescribeCrossRegionBackupsOutput {
	s.ArchType = &v
	return s
}

// SetChargeType sets the ChargeType field's value.
func (s *InstanceInfoForDescribeCrossRegionBackupsOutput) SetChargeType(v string) *InstanceInfoForDescribeCrossRegionBackupsOutput {
	s.ChargeType = &v
	return s
}

// SetDeletionProtection sets the DeletionProtection field's value.
func (s *InstanceInfoForDescribeCrossRegionBackupsOutput) SetDeletionProtection(v string) *InstanceInfoForDescribeCrossRegionBackupsOutput {
	s.DeletionProtection = &v
	return s
}

// SetEngineVersion sets the EngineVersion field's value.
func (s *InstanceInfoForDescribeCrossRegionBackupsOutput) SetEngineVersion(v string) *InstanceInfoForDescribeCrossRegionBackupsOutput {
	s.EngineVersion = &v
	return s
}

// SetExpiredTime sets the ExpiredTime field's value.
func (s *InstanceInfoForDescribeCrossRegionBackupsOutput) SetExpiredTime(v string) *InstanceInfoForDescribeCrossRegionBackupsOutput {
	s.ExpiredTime = &v
	return s
}

// SetInstanceId sets the InstanceId field's value.
func (s *InstanceInfoForDescribeCrossRegionBackupsOutput) SetInstanceId(v string) *InstanceInfoForDescribeCrossRegionBackupsOutput {
	s.InstanceId = &v
	return s
}

// SetInstanceName sets the InstanceName field's value.
func (s *InstanceInfoForDescribeCrossRegionBackupsOutput) SetInstanceName(v string) *InstanceInfoForDescribeCrossRegionBackupsOutput {
	s.InstanceName = &v
	return s
}

// SetMaintenanceTime sets the MaintenanceTime field's value.
func (s *InstanceInfoForDescribeCrossRegionBackupsOutput) SetMaintenanceTime(v string) *InstanceInfoForDescribeCrossRegionBackupsOutput {
	s.MaintenanceTime = &v
	return s
}

// SetNetworkType sets the NetworkType field's value.
func (s *InstanceInfoForDescribeCrossRegionBackupsOutput) SetNetworkType(v string) *InstanceInfoForDescribeCrossRegionBackupsOutput {
	s.NetworkType = &v
	return s
}

// SetRegionId sets the RegionId field's value.
func (s *InstanceInfoForDescribeCrossRegionBackupsOutput) SetRegionId(v string) *InstanceInfoForDescribeCrossRegionBackupsOutput {
	s.RegionId = &v
	return s
}

// SetReplicas sets the Replicas field's value.
func (s *InstanceInfoForDescribeCrossRegionBackupsOutput) SetReplicas(v int32) *InstanceInfoForDescribeCrossRegionBackupsOutput {
	s.Replicas = &v
	return s
}

// SetShardCapacity sets the ShardCapacity field's value.
func (s *InstanceInfoForDescribeCrossRegionBackupsOutput) SetShardCapacity(v float64) *InstanceInfoForDescribeCrossRegionBackupsOutput {
	s.ShardCapacity = &v
	return s
}

// SetShardNumber sets the ShardNumber field's value.
func (s *InstanceInfoForDescribeCrossRegionBackupsOutput) SetShardNumber(v int32) *InstanceInfoForDescribeCrossRegionBackupsOutput {
	s.ShardNumber = &v
	return s
}

// SetTotalCapacity sets the TotalCapacity field's value.
func (s *InstanceInfoForDescribeCrossRegionBackupsOutput) SetTotalCapacity(v int64) *InstanceInfoForDescribeCrossRegionBackupsOutput {
	s.TotalCapacity = &v
	return s
}

// SetVpcId sets the VpcId field's value.
func (s *InstanceInfoForDescribeCrossRegionBackupsOutput) SetVpcId(v string) *InstanceInfoForDescribeCrossRegionBackupsOutput {
	s.VpcId = &v
	return s
}

// SetZoneIds sets the ZoneIds field's value.
func (s *InstanceInfoForDescribeCrossRegionBackupsOutput) SetZoneIds(v []*string) *InstanceInfoForDescribeCrossRegionBackupsOutput {
	s.ZoneIds = v
	return s
}

const (
	// EnumOfBackupStrategyForDescribeCrossRegionBackupsOutputManualBackup is a EnumOfBackupStrategyForDescribeCrossRegionBackupsOutput enum value
	EnumOfBackupStrategyForDescribeCrossRegionBackupsOutputManualBackup = "ManualBackup"

	// EnumOfBackupStrategyForDescribeCrossRegionBackupsOutputAutomatedBackup is a EnumOfBackupStrategyForDescribeCrossRegionBackupsOutput enum value
	EnumOfBackupStrategyForDescribeCrossRegionBackupsOutputAutomatedBackup = "AutomatedBackup"

	// EnumOfBackupStrategyForDescribeCrossRegionBackupsOutputDataFlashBack is a EnumOfBackupStrategyForDescribeCrossRegionBackupsOutput enum value
	EnumOfBackupStrategyForDescribeCrossRegionBackupsOutputDataFlashBack = "DataFlashBack"

	// EnumOfBackupStrategyForDescribeCrossRegionBackupsOutputAllStrategy is a EnumOfBackupStrategyForDescribeCrossRegionBackupsOutput enum value
	EnumOfBackupStrategyForDescribeCrossRegionBackupsOutputAllStrategy = "AllStrategy"
)

const (
	// EnumOfBackupStrategyListListForDescribeCrossRegionBackupsInputManualBackup is a EnumOfBackupStrategyListListForDescribeCrossRegionBackupsInput enum value
	EnumOfBackupStrategyListListForDescribeCrossRegionBackupsInputManualBackup = "ManualBackup"

	// EnumOfBackupStrategyListListForDescribeCrossRegionBackupsInputAutomatedBackup is a EnumOfBackupStrategyListListForDescribeCrossRegionBackupsInput enum value
	EnumOfBackupStrategyListListForDescribeCrossRegionBackupsInputAutomatedBackup = "AutomatedBackup"

	// EnumOfBackupStrategyListListForDescribeCrossRegionBackupsInputDataFlashBack is a EnumOfBackupStrategyListListForDescribeCrossRegionBackupsInput enum value
	EnumOfBackupStrategyListListForDescribeCrossRegionBackupsInputDataFlashBack = "DataFlashBack"

	// EnumOfBackupStrategyListListForDescribeCrossRegionBackupsInputAllStrategy is a EnumOfBackupStrategyListListForDescribeCrossRegionBackupsInput enum value
	EnumOfBackupStrategyListListForDescribeCrossRegionBackupsInputAllStrategy = "AllStrategy"
)

const (
	// EnumOfBackupTypeForDescribeCrossRegionBackupsOutputInvalid is a EnumOfBackupTypeForDescribeCrossRegionBackupsOutput enum value
	EnumOfBackupTypeForDescribeCrossRegionBackupsOutputInvalid = "Invalid"

	// EnumOfBackupTypeForDescribeCrossRegionBackupsOutputFull is a EnumOfBackupTypeForDescribeCrossRegionBackupsOutput enum value
	EnumOfBackupTypeForDescribeCrossRegionBackupsOutputFull = "Full"

	// EnumOfBackupTypeForDescribeCrossRegionBackupsOutputInc is a EnumOfBackupTypeForDescribeCrossRegionBackupsOutput enum value
	EnumOfBackupTypeForDescribeCrossRegionBackupsOutputInc = "Inc"

	// EnumOfBackupTypeForDescribeCrossRegionBackupsOutputAll is a EnumOfBackupTypeForDescribeCrossRegionBackupsOutput enum value
	EnumOfBackupTypeForDescribeCrossRegionBackupsOutputAll = "All"
)

const (
	// EnumOfScopeForDescribeCrossRegionBackupsInputOneInstance is a EnumOfScopeForDescribeCrossRegionBackupsInput enum value
	EnumOfScopeForDescribeCrossRegionBackupsInputOneInstance = "OneInstance"

	// EnumOfScopeForDescribeCrossRegionBackupsInputDestroyedInstances is a EnumOfScopeForDescribeCrossRegionBackupsInput enum value
	EnumOfScopeForDescribeCrossRegionBackupsInputDestroyedInstances = "DestroyedInstances"

	// EnumOfScopeForDescribeCrossRegionBackupsInputAccountInstances is a EnumOfScopeForDescribeCrossRegionBackupsInput enum value
	EnumOfScopeForDescribeCrossRegionBackupsInputAccountInstances = "AccountInstances"
)

const (
	// EnumOfStatusForDescribeCrossRegionBackupsInputCreating is a EnumOfStatusForDescribeCrossRegionBackupsInput enum value
	EnumOfStatusForDescribeCrossRegionBackupsInputCreating = "Creating"

	// EnumOfStatusForDescribeCrossRegionBackupsInputAvailable is a EnumOfStatusForDescribeCrossRegionBackupsInput enum value
	EnumOfStatusForDescribeCrossRegionBackupsInputAvailable = "Available"

	// EnumOfStatusForDescribeCrossRegionBackupsInputUnavailable is a EnumOfStatusForDescribeCrossRegionBackupsInput enum value
	EnumOfStatusForDescribeCrossRegionBackupsInputUnavailable = "Unavailable"

	// EnumOfStatusForDescribeCrossRegionBackupsInputDeleting is a EnumOfStatusForDescribeCrossRegionBackupsInput enum value
	EnumOfStatusForDescribeCrossRegionBackupsInputDeleting = "Deleting"
)

const (
	// EnumOfStatusForDescribeCrossRegionBackupsOutputCreating is a EnumOfStatusForDescribeCrossRegionBackupsOutput enum value
	EnumOfStatusForDescribeCrossRegionBackupsOutputCreating = "Creating"

	// EnumOfStatusForDescribeCrossRegionBackupsOutputAvailable is a EnumOfStatusForDescribeCrossRegionBackupsOutput enum value
	EnumOfStatusForDescribeCrossRegionBackupsOutputAvailable = "Available"

	// EnumOfStatusForDescribeCrossRegionBackupsOutputUnavailable is a EnumOfStatusForDescribeCrossRegionBackupsOutput enum value
	EnumOfStatusForDescribeCrossRegionBackupsOutputUnavailable = "Unavailable"

	// EnumOfStatusForDescribeCrossRegionBackupsOutputDeleting is a EnumOfStatusForDescribeCrossRegionBackupsOutput enum value
	EnumOfStatusForDescribeCrossRegionBackupsOutputDeleting = "Deleting"
)

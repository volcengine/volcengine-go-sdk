// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package rdsmysqlv2

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opModifyBackupPolicyCommon = "ModifyBackupPolicy"

// ModifyBackupPolicyCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ModifyBackupPolicyCommon operation. The "output" return
// value will be populated with the ModifyBackupPolicyCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyBackupPolicyCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyBackupPolicyCommon Send returns without error.
//
// See ModifyBackupPolicyCommon for more information on using the ModifyBackupPolicyCommon
// API call, and error handling.
//
//    // Example sending a request using the ModifyBackupPolicyCommonRequest method.
//    req, resp := client.ModifyBackupPolicyCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *RDSMYSQLV2) ModifyBackupPolicyCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opModifyBackupPolicyCommon,
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

// ModifyBackupPolicyCommon API operation for RDS_MYSQL_V2.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RDS_MYSQL_V2's
// API operation ModifyBackupPolicyCommon for usage and error information.
func (c *RDSMYSQLV2) ModifyBackupPolicyCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ModifyBackupPolicyCommonRequest(input)
	return out, req.Send()
}

// ModifyBackupPolicyCommonWithContext is the same as ModifyBackupPolicyCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyBackupPolicyCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RDSMYSQLV2) ModifyBackupPolicyCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ModifyBackupPolicyCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opModifyBackupPolicy = "ModifyBackupPolicy"

// ModifyBackupPolicyRequest generates a "volcengine/request.Request" representing the
// client's request for the ModifyBackupPolicy operation. The "output" return
// value will be populated with the ModifyBackupPolicyCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyBackupPolicyCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyBackupPolicyCommon Send returns without error.
//
// See ModifyBackupPolicy for more information on using the ModifyBackupPolicy
// API call, and error handling.
//
//    // Example sending a request using the ModifyBackupPolicyRequest method.
//    req, resp := client.ModifyBackupPolicyRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *RDSMYSQLV2) ModifyBackupPolicyRequest(input *ModifyBackupPolicyInput) (req *request.Request, output *ModifyBackupPolicyOutput) {
	op := &request.Operation{
		Name:       opModifyBackupPolicy,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ModifyBackupPolicyInput{}
	}

	output = &ModifyBackupPolicyOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// ModifyBackupPolicy API operation for RDS_MYSQL_V2.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RDS_MYSQL_V2's
// API operation ModifyBackupPolicy for usage and error information.
func (c *RDSMYSQLV2) ModifyBackupPolicy(input *ModifyBackupPolicyInput) (*ModifyBackupPolicyOutput, error) {
	req, out := c.ModifyBackupPolicyRequest(input)
	return out, req.Send()
}

// ModifyBackupPolicyWithContext is the same as ModifyBackupPolicy with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyBackupPolicy for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RDSMYSQLV2) ModifyBackupPolicyWithContext(ctx volcengine.Context, input *ModifyBackupPolicyInput, opts ...request.Option) (*ModifyBackupPolicyOutput, error) {
	req, out := c.ModifyBackupPolicyRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ModifyBackupPolicyInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	BinlogBackupAllRetention *bool `type:"boolean" json:",omitempty"`

	BinlogBackupEnabled *bool `type:"boolean" json:",omitempty"`

	BinlogBackupEncryptionEnabled *bool `type:"boolean" json:",omitempty"`

	BinlogFileCountsEnable *bool `type:"boolean" json:",omitempty"`

	BinlogLimitCount *int32 `type:"int32" json:",omitempty"`

	BinlogLocalRetentionHour *int32 `type:"int32" json:",omitempty"`

	BinlogSpaceLimitEnable *bool `type:"boolean" json:",omitempty"`

	BinlogStoragePercentage *int32 `type:"int32" json:",omitempty"`

	DataBackupAllRetention *bool `type:"boolean" json:",omitempty"`

	DataBackupEncryptionEnabled *bool `type:"boolean" json:",omitempty"`

	DataBackupRetentionDay *int32 `type:"int32" json:",omitempty"`

	DataFullBackupPeriods []*string `type:"list" json:",omitempty"`

	DataFullBackupStartUTCHour *int32 `type:"int32" json:",omitempty"`

	DataFullBackupTime *string `type:"string" json:",omitempty"`

	DataIncrBackupPeriods []*string `type:"list" json:",omitempty"`

	DataKeepDaysAfterReleased *int32 `type:"int32" json:",omitempty"`

	DataKeepPolicyAfterReleased *string `type:"string" json:",omitempty"`

	HourlyIncrBackupEnable *bool `type:"boolean" json:",omitempty"`

	IncrBackupHourPeriod *int32 `type:"int32" json:",omitempty"`

	// InstanceId is a required field
	InstanceId *string `type:"string" json:",omitempty" required:"true"`

	LockDDLTime *int32 `type:"int32" json:",omitempty"`

	LogBackupRetentionDay *int32 `type:"int32" json:",omitempty"`

	RetentionPolicySynced *bool `type:"boolean" json:",omitempty"`
}

// String returns the string representation
func (s ModifyBackupPolicyInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyBackupPolicyInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ModifyBackupPolicyInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ModifyBackupPolicyInput"}
	if s.InstanceId == nil {
		invalidParams.Add(request.NewErrParamRequired("InstanceId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetBinlogBackupAllRetention sets the BinlogBackupAllRetention field's value.
func (s *ModifyBackupPolicyInput) SetBinlogBackupAllRetention(v bool) *ModifyBackupPolicyInput {
	s.BinlogBackupAllRetention = &v
	return s
}

// SetBinlogBackupEnabled sets the BinlogBackupEnabled field's value.
func (s *ModifyBackupPolicyInput) SetBinlogBackupEnabled(v bool) *ModifyBackupPolicyInput {
	s.BinlogBackupEnabled = &v
	return s
}

// SetBinlogBackupEncryptionEnabled sets the BinlogBackupEncryptionEnabled field's value.
func (s *ModifyBackupPolicyInput) SetBinlogBackupEncryptionEnabled(v bool) *ModifyBackupPolicyInput {
	s.BinlogBackupEncryptionEnabled = &v
	return s
}

// SetBinlogFileCountsEnable sets the BinlogFileCountsEnable field's value.
func (s *ModifyBackupPolicyInput) SetBinlogFileCountsEnable(v bool) *ModifyBackupPolicyInput {
	s.BinlogFileCountsEnable = &v
	return s
}

// SetBinlogLimitCount sets the BinlogLimitCount field's value.
func (s *ModifyBackupPolicyInput) SetBinlogLimitCount(v int32) *ModifyBackupPolicyInput {
	s.BinlogLimitCount = &v
	return s
}

// SetBinlogLocalRetentionHour sets the BinlogLocalRetentionHour field's value.
func (s *ModifyBackupPolicyInput) SetBinlogLocalRetentionHour(v int32) *ModifyBackupPolicyInput {
	s.BinlogLocalRetentionHour = &v
	return s
}

// SetBinlogSpaceLimitEnable sets the BinlogSpaceLimitEnable field's value.
func (s *ModifyBackupPolicyInput) SetBinlogSpaceLimitEnable(v bool) *ModifyBackupPolicyInput {
	s.BinlogSpaceLimitEnable = &v
	return s
}

// SetBinlogStoragePercentage sets the BinlogStoragePercentage field's value.
func (s *ModifyBackupPolicyInput) SetBinlogStoragePercentage(v int32) *ModifyBackupPolicyInput {
	s.BinlogStoragePercentage = &v
	return s
}

// SetDataBackupAllRetention sets the DataBackupAllRetention field's value.
func (s *ModifyBackupPolicyInput) SetDataBackupAllRetention(v bool) *ModifyBackupPolicyInput {
	s.DataBackupAllRetention = &v
	return s
}

// SetDataBackupEncryptionEnabled sets the DataBackupEncryptionEnabled field's value.
func (s *ModifyBackupPolicyInput) SetDataBackupEncryptionEnabled(v bool) *ModifyBackupPolicyInput {
	s.DataBackupEncryptionEnabled = &v
	return s
}

// SetDataBackupRetentionDay sets the DataBackupRetentionDay field's value.
func (s *ModifyBackupPolicyInput) SetDataBackupRetentionDay(v int32) *ModifyBackupPolicyInput {
	s.DataBackupRetentionDay = &v
	return s
}

// SetDataFullBackupPeriods sets the DataFullBackupPeriods field's value.
func (s *ModifyBackupPolicyInput) SetDataFullBackupPeriods(v []*string) *ModifyBackupPolicyInput {
	s.DataFullBackupPeriods = v
	return s
}

// SetDataFullBackupStartUTCHour sets the DataFullBackupStartUTCHour field's value.
func (s *ModifyBackupPolicyInput) SetDataFullBackupStartUTCHour(v int32) *ModifyBackupPolicyInput {
	s.DataFullBackupStartUTCHour = &v
	return s
}

// SetDataFullBackupTime sets the DataFullBackupTime field's value.
func (s *ModifyBackupPolicyInput) SetDataFullBackupTime(v string) *ModifyBackupPolicyInput {
	s.DataFullBackupTime = &v
	return s
}

// SetDataIncrBackupPeriods sets the DataIncrBackupPeriods field's value.
func (s *ModifyBackupPolicyInput) SetDataIncrBackupPeriods(v []*string) *ModifyBackupPolicyInput {
	s.DataIncrBackupPeriods = v
	return s
}

// SetDataKeepDaysAfterReleased sets the DataKeepDaysAfterReleased field's value.
func (s *ModifyBackupPolicyInput) SetDataKeepDaysAfterReleased(v int32) *ModifyBackupPolicyInput {
	s.DataKeepDaysAfterReleased = &v
	return s
}

// SetDataKeepPolicyAfterReleased sets the DataKeepPolicyAfterReleased field's value.
func (s *ModifyBackupPolicyInput) SetDataKeepPolicyAfterReleased(v string) *ModifyBackupPolicyInput {
	s.DataKeepPolicyAfterReleased = &v
	return s
}

// SetHourlyIncrBackupEnable sets the HourlyIncrBackupEnable field's value.
func (s *ModifyBackupPolicyInput) SetHourlyIncrBackupEnable(v bool) *ModifyBackupPolicyInput {
	s.HourlyIncrBackupEnable = &v
	return s
}

// SetIncrBackupHourPeriod sets the IncrBackupHourPeriod field's value.
func (s *ModifyBackupPolicyInput) SetIncrBackupHourPeriod(v int32) *ModifyBackupPolicyInput {
	s.IncrBackupHourPeriod = &v
	return s
}

// SetInstanceId sets the InstanceId field's value.
func (s *ModifyBackupPolicyInput) SetInstanceId(v string) *ModifyBackupPolicyInput {
	s.InstanceId = &v
	return s
}

// SetLockDDLTime sets the LockDDLTime field's value.
func (s *ModifyBackupPolicyInput) SetLockDDLTime(v int32) *ModifyBackupPolicyInput {
	s.LockDDLTime = &v
	return s
}

// SetLogBackupRetentionDay sets the LogBackupRetentionDay field's value.
func (s *ModifyBackupPolicyInput) SetLogBackupRetentionDay(v int32) *ModifyBackupPolicyInput {
	s.LogBackupRetentionDay = &v
	return s
}

// SetRetentionPolicySynced sets the RetentionPolicySynced field's value.
func (s *ModifyBackupPolicyInput) SetRetentionPolicySynced(v bool) *ModifyBackupPolicyInput {
	s.RetentionPolicySynced = &v
	return s
}

type ModifyBackupPolicyOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	BinlogBackupAllRetention *bool `type:"boolean" json:",omitempty"`

	BinlogBackupEnabled *bool `type:"boolean" json:",omitempty"`

	BinlogBackupEncryptionEnabled *bool `type:"boolean" json:",omitempty"`

	BinlogFileCountsEnable *bool `type:"boolean" json:",omitempty"`

	BinlogLimitCount *int32 `type:"int32" json:",omitempty"`

	BinlogLocalRetentionHour *int32 `type:"int32" json:",omitempty"`

	BinlogSpaceLimitEnable *bool `type:"boolean" json:",omitempty"`

	BinlogStoragePercentage *int32 `type:"int32" json:",omitempty"`

	DataBackupAllRetention *bool `type:"boolean" json:",omitempty"`

	DataBackupEncryptionEnabled *bool `type:"boolean" json:",omitempty"`

	DataBackupRetentionDay *int32 `type:"int32" json:",omitempty"`

	DataFullBackupPeriods []*string `type:"list" json:",omitempty"`

	DataFullBackupStartUTCHour *int32 `type:"int32" json:",omitempty"`

	DataFullBackupTime *string `type:"string" json:",omitempty"`

	DataIncrBackupPeriods []*string `type:"list" json:",omitempty"`

	DataKeepDaysAfterReleased *int32 `type:"int32" json:",omitempty"`

	DataKeepPolicyAfterReleased *string `type:"string" json:",omitempty"`

	HourlyIncrBackupEnable *bool `type:"boolean" json:",omitempty"`

	IncrBackupHourPeriod *int32 `type:"int32" json:",omitempty"`

	InstanceId *string `type:"string" json:",omitempty"`

	LockDDLTime *int32 `type:"int32" json:",omitempty"`

	LogBackupRetentionDay *int32 `type:"int32" json:",omitempty"`

	RetentionPolicySynced *bool `type:"boolean" json:",omitempty"`
}

// String returns the string representation
func (s ModifyBackupPolicyOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyBackupPolicyOutput) GoString() string {
	return s.String()
}

// SetBinlogBackupAllRetention sets the BinlogBackupAllRetention field's value.
func (s *ModifyBackupPolicyOutput) SetBinlogBackupAllRetention(v bool) *ModifyBackupPolicyOutput {
	s.BinlogBackupAllRetention = &v
	return s
}

// SetBinlogBackupEnabled sets the BinlogBackupEnabled field's value.
func (s *ModifyBackupPolicyOutput) SetBinlogBackupEnabled(v bool) *ModifyBackupPolicyOutput {
	s.BinlogBackupEnabled = &v
	return s
}

// SetBinlogBackupEncryptionEnabled sets the BinlogBackupEncryptionEnabled field's value.
func (s *ModifyBackupPolicyOutput) SetBinlogBackupEncryptionEnabled(v bool) *ModifyBackupPolicyOutput {
	s.BinlogBackupEncryptionEnabled = &v
	return s
}

// SetBinlogFileCountsEnable sets the BinlogFileCountsEnable field's value.
func (s *ModifyBackupPolicyOutput) SetBinlogFileCountsEnable(v bool) *ModifyBackupPolicyOutput {
	s.BinlogFileCountsEnable = &v
	return s
}

// SetBinlogLimitCount sets the BinlogLimitCount field's value.
func (s *ModifyBackupPolicyOutput) SetBinlogLimitCount(v int32) *ModifyBackupPolicyOutput {
	s.BinlogLimitCount = &v
	return s
}

// SetBinlogLocalRetentionHour sets the BinlogLocalRetentionHour field's value.
func (s *ModifyBackupPolicyOutput) SetBinlogLocalRetentionHour(v int32) *ModifyBackupPolicyOutput {
	s.BinlogLocalRetentionHour = &v
	return s
}

// SetBinlogSpaceLimitEnable sets the BinlogSpaceLimitEnable field's value.
func (s *ModifyBackupPolicyOutput) SetBinlogSpaceLimitEnable(v bool) *ModifyBackupPolicyOutput {
	s.BinlogSpaceLimitEnable = &v
	return s
}

// SetBinlogStoragePercentage sets the BinlogStoragePercentage field's value.
func (s *ModifyBackupPolicyOutput) SetBinlogStoragePercentage(v int32) *ModifyBackupPolicyOutput {
	s.BinlogStoragePercentage = &v
	return s
}

// SetDataBackupAllRetention sets the DataBackupAllRetention field's value.
func (s *ModifyBackupPolicyOutput) SetDataBackupAllRetention(v bool) *ModifyBackupPolicyOutput {
	s.DataBackupAllRetention = &v
	return s
}

// SetDataBackupEncryptionEnabled sets the DataBackupEncryptionEnabled field's value.
func (s *ModifyBackupPolicyOutput) SetDataBackupEncryptionEnabled(v bool) *ModifyBackupPolicyOutput {
	s.DataBackupEncryptionEnabled = &v
	return s
}

// SetDataBackupRetentionDay sets the DataBackupRetentionDay field's value.
func (s *ModifyBackupPolicyOutput) SetDataBackupRetentionDay(v int32) *ModifyBackupPolicyOutput {
	s.DataBackupRetentionDay = &v
	return s
}

// SetDataFullBackupPeriods sets the DataFullBackupPeriods field's value.
func (s *ModifyBackupPolicyOutput) SetDataFullBackupPeriods(v []*string) *ModifyBackupPolicyOutput {
	s.DataFullBackupPeriods = v
	return s
}

// SetDataFullBackupStartUTCHour sets the DataFullBackupStartUTCHour field's value.
func (s *ModifyBackupPolicyOutput) SetDataFullBackupStartUTCHour(v int32) *ModifyBackupPolicyOutput {
	s.DataFullBackupStartUTCHour = &v
	return s
}

// SetDataFullBackupTime sets the DataFullBackupTime field's value.
func (s *ModifyBackupPolicyOutput) SetDataFullBackupTime(v string) *ModifyBackupPolicyOutput {
	s.DataFullBackupTime = &v
	return s
}

// SetDataIncrBackupPeriods sets the DataIncrBackupPeriods field's value.
func (s *ModifyBackupPolicyOutput) SetDataIncrBackupPeriods(v []*string) *ModifyBackupPolicyOutput {
	s.DataIncrBackupPeriods = v
	return s
}

// SetDataKeepDaysAfterReleased sets the DataKeepDaysAfterReleased field's value.
func (s *ModifyBackupPolicyOutput) SetDataKeepDaysAfterReleased(v int32) *ModifyBackupPolicyOutput {
	s.DataKeepDaysAfterReleased = &v
	return s
}

// SetDataKeepPolicyAfterReleased sets the DataKeepPolicyAfterReleased field's value.
func (s *ModifyBackupPolicyOutput) SetDataKeepPolicyAfterReleased(v string) *ModifyBackupPolicyOutput {
	s.DataKeepPolicyAfterReleased = &v
	return s
}

// SetHourlyIncrBackupEnable sets the HourlyIncrBackupEnable field's value.
func (s *ModifyBackupPolicyOutput) SetHourlyIncrBackupEnable(v bool) *ModifyBackupPolicyOutput {
	s.HourlyIncrBackupEnable = &v
	return s
}

// SetIncrBackupHourPeriod sets the IncrBackupHourPeriod field's value.
func (s *ModifyBackupPolicyOutput) SetIncrBackupHourPeriod(v int32) *ModifyBackupPolicyOutput {
	s.IncrBackupHourPeriod = &v
	return s
}

// SetInstanceId sets the InstanceId field's value.
func (s *ModifyBackupPolicyOutput) SetInstanceId(v string) *ModifyBackupPolicyOutput {
	s.InstanceId = &v
	return s
}

// SetLockDDLTime sets the LockDDLTime field's value.
func (s *ModifyBackupPolicyOutput) SetLockDDLTime(v int32) *ModifyBackupPolicyOutput {
	s.LockDDLTime = &v
	return s
}

// SetLogBackupRetentionDay sets the LogBackupRetentionDay field's value.
func (s *ModifyBackupPolicyOutput) SetLogBackupRetentionDay(v int32) *ModifyBackupPolicyOutput {
	s.LogBackupRetentionDay = &v
	return s
}

// SetRetentionPolicySynced sets the RetentionPolicySynced field's value.
func (s *ModifyBackupPolicyOutput) SetRetentionPolicySynced(v bool) *ModifyBackupPolicyOutput {
	s.RetentionPolicySynced = &v
	return s
}

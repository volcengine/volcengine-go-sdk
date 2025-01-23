// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package rdsmysqlv2

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribeBackupsCommon = "DescribeBackups"

// DescribeBackupsCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeBackupsCommon operation. The "output" return
// value will be populated with the DescribeBackupsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeBackupsCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeBackupsCommon Send returns without error.
//
// See DescribeBackupsCommon for more information on using the DescribeBackupsCommon
// API call, and error handling.
//
//    // Example sending a request using the DescribeBackupsCommonRequest method.
//    req, resp := client.DescribeBackupsCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *RDSMYSQLV2) DescribeBackupsCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeBackupsCommon,
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

// DescribeBackupsCommon API operation for RDS_MYSQL_V2.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RDS_MYSQL_V2's
// API operation DescribeBackupsCommon for usage and error information.
func (c *RDSMYSQLV2) DescribeBackupsCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeBackupsCommonRequest(input)
	return out, req.Send()
}

// DescribeBackupsCommonWithContext is the same as DescribeBackupsCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeBackupsCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RDSMYSQLV2) DescribeBackupsCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeBackupsCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeBackups = "DescribeBackups"

// DescribeBackupsRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeBackups operation. The "output" return
// value will be populated with the DescribeBackupsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeBackupsCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeBackupsCommon Send returns without error.
//
// See DescribeBackups for more information on using the DescribeBackups
// API call, and error handling.
//
//    // Example sending a request using the DescribeBackupsRequest method.
//    req, resp := client.DescribeBackupsRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *RDSMYSQLV2) DescribeBackupsRequest(input *DescribeBackupsInput) (req *request.Request, output *DescribeBackupsOutput) {
	op := &request.Operation{
		Name:       opDescribeBackups,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeBackupsInput{}
	}

	output = &DescribeBackupsOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// DescribeBackups API operation for RDS_MYSQL_V2.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RDS_MYSQL_V2's
// API operation DescribeBackups for usage and error information.
func (c *RDSMYSQLV2) DescribeBackups(input *DescribeBackupsInput) (*DescribeBackupsOutput, error) {
	req, out := c.DescribeBackupsRequest(input)
	return out, req.Send()
}

// DescribeBackupsWithContext is the same as DescribeBackups with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeBackups for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RDSMYSQLV2) DescribeBackupsWithContext(ctx volcengine.Context, input *DescribeBackupsInput, opts ...request.Option) (*DescribeBackupsOutput, error) {
	req, out := c.DescribeBackupsRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type BackupForDescribeBackupsOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	BackupEndTime *string `type:"string" json:",omitempty"`

	BackupFileName *string `type:"string" json:",omitempty"`

	BackupFileSize *int64 `type:"int64" json:",omitempty"`

	BackupId *string `type:"string" json:",omitempty"`

	BackupMethod *string `type:"string" json:",omitempty"`

	BackupRegion *string `type:"string" json:",omitempty"`

	BackupStartTime *string `type:"string" json:",omitempty"`

	BackupStatus *string `type:"string" json:",omitempty"`

	BackupType *string `type:"string" json:",omitempty"`

	ConsistentTime *string `type:"string" json:",omitempty"`

	CreateType *string `type:"string" json:",omitempty"`

	DBTableInfos []*DBTableInfoForDescribeBackupsOutput `type:"list" json:",omitempty"`

	DownloadStatus *string `type:"string" json:",omitempty"`

	ErrorMessage *string `type:"string" json:",omitempty"`

	ExpiredTime *string `type:"string" json:",omitempty"`

	IsEncrypted *bool `type:"boolean" json:",omitempty"`

	IsExpired *bool `type:"boolean" json:",omitempty"`
}

// String returns the string representation
func (s BackupForDescribeBackupsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s BackupForDescribeBackupsOutput) GoString() string {
	return s.String()
}

// SetBackupEndTime sets the BackupEndTime field's value.
func (s *BackupForDescribeBackupsOutput) SetBackupEndTime(v string) *BackupForDescribeBackupsOutput {
	s.BackupEndTime = &v
	return s
}

// SetBackupFileName sets the BackupFileName field's value.
func (s *BackupForDescribeBackupsOutput) SetBackupFileName(v string) *BackupForDescribeBackupsOutput {
	s.BackupFileName = &v
	return s
}

// SetBackupFileSize sets the BackupFileSize field's value.
func (s *BackupForDescribeBackupsOutput) SetBackupFileSize(v int64) *BackupForDescribeBackupsOutput {
	s.BackupFileSize = &v
	return s
}

// SetBackupId sets the BackupId field's value.
func (s *BackupForDescribeBackupsOutput) SetBackupId(v string) *BackupForDescribeBackupsOutput {
	s.BackupId = &v
	return s
}

// SetBackupMethod sets the BackupMethod field's value.
func (s *BackupForDescribeBackupsOutput) SetBackupMethod(v string) *BackupForDescribeBackupsOutput {
	s.BackupMethod = &v
	return s
}

// SetBackupRegion sets the BackupRegion field's value.
func (s *BackupForDescribeBackupsOutput) SetBackupRegion(v string) *BackupForDescribeBackupsOutput {
	s.BackupRegion = &v
	return s
}

// SetBackupStartTime sets the BackupStartTime field's value.
func (s *BackupForDescribeBackupsOutput) SetBackupStartTime(v string) *BackupForDescribeBackupsOutput {
	s.BackupStartTime = &v
	return s
}

// SetBackupStatus sets the BackupStatus field's value.
func (s *BackupForDescribeBackupsOutput) SetBackupStatus(v string) *BackupForDescribeBackupsOutput {
	s.BackupStatus = &v
	return s
}

// SetBackupType sets the BackupType field's value.
func (s *BackupForDescribeBackupsOutput) SetBackupType(v string) *BackupForDescribeBackupsOutput {
	s.BackupType = &v
	return s
}

// SetConsistentTime sets the ConsistentTime field's value.
func (s *BackupForDescribeBackupsOutput) SetConsistentTime(v string) *BackupForDescribeBackupsOutput {
	s.ConsistentTime = &v
	return s
}

// SetCreateType sets the CreateType field's value.
func (s *BackupForDescribeBackupsOutput) SetCreateType(v string) *BackupForDescribeBackupsOutput {
	s.CreateType = &v
	return s
}

// SetDBTableInfos sets the DBTableInfos field's value.
func (s *BackupForDescribeBackupsOutput) SetDBTableInfos(v []*DBTableInfoForDescribeBackupsOutput) *BackupForDescribeBackupsOutput {
	s.DBTableInfos = v
	return s
}

// SetDownloadStatus sets the DownloadStatus field's value.
func (s *BackupForDescribeBackupsOutput) SetDownloadStatus(v string) *BackupForDescribeBackupsOutput {
	s.DownloadStatus = &v
	return s
}

// SetErrorMessage sets the ErrorMessage field's value.
func (s *BackupForDescribeBackupsOutput) SetErrorMessage(v string) *BackupForDescribeBackupsOutput {
	s.ErrorMessage = &v
	return s
}

// SetExpiredTime sets the ExpiredTime field's value.
func (s *BackupForDescribeBackupsOutput) SetExpiredTime(v string) *BackupForDescribeBackupsOutput {
	s.ExpiredTime = &v
	return s
}

// SetIsEncrypted sets the IsEncrypted field's value.
func (s *BackupForDescribeBackupsOutput) SetIsEncrypted(v bool) *BackupForDescribeBackupsOutput {
	s.IsEncrypted = &v
	return s
}

// SetIsExpired sets the IsExpired field's value.
func (s *BackupForDescribeBackupsOutput) SetIsExpired(v bool) *BackupForDescribeBackupsOutput {
	s.IsExpired = &v
	return s
}

type DBTableInfoForDescribeBackupsOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Database *string `type:"string" json:",omitempty"`

	Tables []*string `type:"list" json:",omitempty"`
}

// String returns the string representation
func (s DBTableInfoForDescribeBackupsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DBTableInfoForDescribeBackupsOutput) GoString() string {
	return s.String()
}

// SetDatabase sets the Database field's value.
func (s *DBTableInfoForDescribeBackupsOutput) SetDatabase(v string) *DBTableInfoForDescribeBackupsOutput {
	s.Database = &v
	return s
}

// SetTables sets the Tables field's value.
func (s *DBTableInfoForDescribeBackupsOutput) SetTables(v []*string) *DBTableInfoForDescribeBackupsOutput {
	s.Tables = v
	return s
}

type DescribeBackupsInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	BackupEndTime *string `type:"string" json:",omitempty"`

	BackupId *string `type:"string" json:",omitempty"`

	BackupMethod *string `type:"string" json:",omitempty"`

	BackupStartTime *string `type:"string" json:",omitempty"`

	BackupStatus *string `type:"string" json:",omitempty"`

	BackupType *string `type:"string" json:",omitempty"`

	CreateType *string `type:"string" json:",omitempty"`

	// InstanceId is a required field
	InstanceId *string `type:"string" json:",omitempty" required:"true"`

	PageNumber *int32 `type:"int32" json:",omitempty"`

	PageSize *int32 `type:"int32" json:",omitempty"`
}

// String returns the string representation
func (s DescribeBackupsInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeBackupsInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeBackupsInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DescribeBackupsInput"}
	if s.InstanceId == nil {
		invalidParams.Add(request.NewErrParamRequired("InstanceId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetBackupEndTime sets the BackupEndTime field's value.
func (s *DescribeBackupsInput) SetBackupEndTime(v string) *DescribeBackupsInput {
	s.BackupEndTime = &v
	return s
}

// SetBackupId sets the BackupId field's value.
func (s *DescribeBackupsInput) SetBackupId(v string) *DescribeBackupsInput {
	s.BackupId = &v
	return s
}

// SetBackupMethod sets the BackupMethod field's value.
func (s *DescribeBackupsInput) SetBackupMethod(v string) *DescribeBackupsInput {
	s.BackupMethod = &v
	return s
}

// SetBackupStartTime sets the BackupStartTime field's value.
func (s *DescribeBackupsInput) SetBackupStartTime(v string) *DescribeBackupsInput {
	s.BackupStartTime = &v
	return s
}

// SetBackupStatus sets the BackupStatus field's value.
func (s *DescribeBackupsInput) SetBackupStatus(v string) *DescribeBackupsInput {
	s.BackupStatus = &v
	return s
}

// SetBackupType sets the BackupType field's value.
func (s *DescribeBackupsInput) SetBackupType(v string) *DescribeBackupsInput {
	s.BackupType = &v
	return s
}

// SetCreateType sets the CreateType field's value.
func (s *DescribeBackupsInput) SetCreateType(v string) *DescribeBackupsInput {
	s.CreateType = &v
	return s
}

// SetInstanceId sets the InstanceId field's value.
func (s *DescribeBackupsInput) SetInstanceId(v string) *DescribeBackupsInput {
	s.InstanceId = &v
	return s
}

// SetPageNumber sets the PageNumber field's value.
func (s *DescribeBackupsInput) SetPageNumber(v int32) *DescribeBackupsInput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *DescribeBackupsInput) SetPageSize(v int32) *DescribeBackupsInput {
	s.PageSize = &v
	return s
}

type DescribeBackupsOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	Backups []*BackupForDescribeBackupsOutput `type:"list" json:",omitempty"`

	Total *int32 `type:"int32" json:",omitempty"`
}

// String returns the string representation
func (s DescribeBackupsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeBackupsOutput) GoString() string {
	return s.String()
}

// SetBackups sets the Backups field's value.
func (s *DescribeBackupsOutput) SetBackups(v []*BackupForDescribeBackupsOutput) *DescribeBackupsOutput {
	s.Backups = v
	return s
}

// SetTotal sets the Total field's value.
func (s *DescribeBackupsOutput) SetTotal(v int32) *DescribeBackupsOutput {
	s.Total = &v
	return s
}

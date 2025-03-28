// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package rdsmssql

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribeBackupDetailCommon = "DescribeBackupDetail"

// DescribeBackupDetailCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeBackupDetailCommon operation. The "output" return
// value will be populated with the DescribeBackupDetailCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeBackupDetailCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeBackupDetailCommon Send returns without error.
//
// See DescribeBackupDetailCommon for more information on using the DescribeBackupDetailCommon
// API call, and error handling.
//
//    // Example sending a request using the DescribeBackupDetailCommonRequest method.
//    req, resp := client.DescribeBackupDetailCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *RDSMSSQL) DescribeBackupDetailCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeBackupDetailCommon,
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

// DescribeBackupDetailCommon API operation for RDS_MSSQL.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RDS_MSSQL's
// API operation DescribeBackupDetailCommon for usage and error information.
func (c *RDSMSSQL) DescribeBackupDetailCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeBackupDetailCommonRequest(input)
	return out, req.Send()
}

// DescribeBackupDetailCommonWithContext is the same as DescribeBackupDetailCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeBackupDetailCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RDSMSSQL) DescribeBackupDetailCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeBackupDetailCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeBackupDetail = "DescribeBackupDetail"

// DescribeBackupDetailRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeBackupDetail operation. The "output" return
// value will be populated with the DescribeBackupDetailCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeBackupDetailCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeBackupDetailCommon Send returns without error.
//
// See DescribeBackupDetail for more information on using the DescribeBackupDetail
// API call, and error handling.
//
//    // Example sending a request using the DescribeBackupDetailRequest method.
//    req, resp := client.DescribeBackupDetailRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *RDSMSSQL) DescribeBackupDetailRequest(input *DescribeBackupDetailInput) (req *request.Request, output *DescribeBackupDetailOutput) {
	op := &request.Operation{
		Name:       opDescribeBackupDetail,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeBackupDetailInput{}
	}

	output = &DescribeBackupDetailOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// DescribeBackupDetail API operation for RDS_MSSQL.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RDS_MSSQL's
// API operation DescribeBackupDetail for usage and error information.
func (c *RDSMSSQL) DescribeBackupDetail(input *DescribeBackupDetailInput) (*DescribeBackupDetailOutput, error) {
	req, out := c.DescribeBackupDetailRequest(input)
	return out, req.Send()
}

// DescribeBackupDetailWithContext is the same as DescribeBackupDetail with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeBackupDetail for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RDSMSSQL) DescribeBackupDetailWithContext(ctx volcengine.Context, input *DescribeBackupDetailInput, opts ...request.Option) (*DescribeBackupDetailOutput, error) {
	req, out := c.DescribeBackupDetailRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type BackupDatabaseDetailForDescribeBackupDetailOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	BackupDownloadLinkEIP *string `type:"string" json:",omitempty"`

	BackupDownloadLinkInner *string `type:"string" json:",omitempty"`

	BackupEndTime *string `type:"string" json:",omitempty"`

	BackupFileName *string `type:"string" json:",omitempty"`

	BackupFileSize *int64 `type:"int64" json:",omitempty"`

	BackupStartTime *string `type:"string" json:",omitempty"`

	BackupType *string `type:"string" json:",omitempty"`

	DatabaseName *string `type:"string" json:",omitempty"`

	DownloadProgress *int32 `type:"int32" json:",omitempty"`

	DownloadStatus *string `type:"string" json:",omitempty"`

	LinkExpiredTime *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s BackupDatabaseDetailForDescribeBackupDetailOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s BackupDatabaseDetailForDescribeBackupDetailOutput) GoString() string {
	return s.String()
}

// SetBackupDownloadLinkEIP sets the BackupDownloadLinkEIP field's value.
func (s *BackupDatabaseDetailForDescribeBackupDetailOutput) SetBackupDownloadLinkEIP(v string) *BackupDatabaseDetailForDescribeBackupDetailOutput {
	s.BackupDownloadLinkEIP = &v
	return s
}

// SetBackupDownloadLinkInner sets the BackupDownloadLinkInner field's value.
func (s *BackupDatabaseDetailForDescribeBackupDetailOutput) SetBackupDownloadLinkInner(v string) *BackupDatabaseDetailForDescribeBackupDetailOutput {
	s.BackupDownloadLinkInner = &v
	return s
}

// SetBackupEndTime sets the BackupEndTime field's value.
func (s *BackupDatabaseDetailForDescribeBackupDetailOutput) SetBackupEndTime(v string) *BackupDatabaseDetailForDescribeBackupDetailOutput {
	s.BackupEndTime = &v
	return s
}

// SetBackupFileName sets the BackupFileName field's value.
func (s *BackupDatabaseDetailForDescribeBackupDetailOutput) SetBackupFileName(v string) *BackupDatabaseDetailForDescribeBackupDetailOutput {
	s.BackupFileName = &v
	return s
}

// SetBackupFileSize sets the BackupFileSize field's value.
func (s *BackupDatabaseDetailForDescribeBackupDetailOutput) SetBackupFileSize(v int64) *BackupDatabaseDetailForDescribeBackupDetailOutput {
	s.BackupFileSize = &v
	return s
}

// SetBackupStartTime sets the BackupStartTime field's value.
func (s *BackupDatabaseDetailForDescribeBackupDetailOutput) SetBackupStartTime(v string) *BackupDatabaseDetailForDescribeBackupDetailOutput {
	s.BackupStartTime = &v
	return s
}

// SetBackupType sets the BackupType field's value.
func (s *BackupDatabaseDetailForDescribeBackupDetailOutput) SetBackupType(v string) *BackupDatabaseDetailForDescribeBackupDetailOutput {
	s.BackupType = &v
	return s
}

// SetDatabaseName sets the DatabaseName field's value.
func (s *BackupDatabaseDetailForDescribeBackupDetailOutput) SetDatabaseName(v string) *BackupDatabaseDetailForDescribeBackupDetailOutput {
	s.DatabaseName = &v
	return s
}

// SetDownloadProgress sets the DownloadProgress field's value.
func (s *BackupDatabaseDetailForDescribeBackupDetailOutput) SetDownloadProgress(v int32) *BackupDatabaseDetailForDescribeBackupDetailOutput {
	s.DownloadProgress = &v
	return s
}

// SetDownloadStatus sets the DownloadStatus field's value.
func (s *BackupDatabaseDetailForDescribeBackupDetailOutput) SetDownloadStatus(v string) *BackupDatabaseDetailForDescribeBackupDetailOutput {
	s.DownloadStatus = &v
	return s
}

// SetLinkExpiredTime sets the LinkExpiredTime field's value.
func (s *BackupDatabaseDetailForDescribeBackupDetailOutput) SetLinkExpiredTime(v string) *BackupDatabaseDetailForDescribeBackupDetailOutput {
	s.LinkExpiredTime = &v
	return s
}

type BackupsInfoForDescribeBackupDetailOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	BackupDatabaseDetail []*BackupDatabaseDetailForDescribeBackupDetailOutput `type:"list" json:",omitempty"`

	BackupEndTime *string `type:"string" json:",omitempty"`

	BackupFileSize *int64 `type:"int64" json:",omitempty"`

	BackupId *string `type:"string" json:",omitempty"`

	BackupMethod *string `type:"string" json:",omitempty"`

	BackupStartTime *string `type:"string" json:",omitempty"`

	BackupStatus *string `type:"string" json:",omitempty"`

	BackupType *string `type:"string" json:",omitempty"`

	CreateType *string `type:"string" json:",omitempty"`

	DownloadProgress *int32 `type:"int32" json:",omitempty"`

	DownloadStatus *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s BackupsInfoForDescribeBackupDetailOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s BackupsInfoForDescribeBackupDetailOutput) GoString() string {
	return s.String()
}

// SetBackupDatabaseDetail sets the BackupDatabaseDetail field's value.
func (s *BackupsInfoForDescribeBackupDetailOutput) SetBackupDatabaseDetail(v []*BackupDatabaseDetailForDescribeBackupDetailOutput) *BackupsInfoForDescribeBackupDetailOutput {
	s.BackupDatabaseDetail = v
	return s
}

// SetBackupEndTime sets the BackupEndTime field's value.
func (s *BackupsInfoForDescribeBackupDetailOutput) SetBackupEndTime(v string) *BackupsInfoForDescribeBackupDetailOutput {
	s.BackupEndTime = &v
	return s
}

// SetBackupFileSize sets the BackupFileSize field's value.
func (s *BackupsInfoForDescribeBackupDetailOutput) SetBackupFileSize(v int64) *BackupsInfoForDescribeBackupDetailOutput {
	s.BackupFileSize = &v
	return s
}

// SetBackupId sets the BackupId field's value.
func (s *BackupsInfoForDescribeBackupDetailOutput) SetBackupId(v string) *BackupsInfoForDescribeBackupDetailOutput {
	s.BackupId = &v
	return s
}

// SetBackupMethod sets the BackupMethod field's value.
func (s *BackupsInfoForDescribeBackupDetailOutput) SetBackupMethod(v string) *BackupsInfoForDescribeBackupDetailOutput {
	s.BackupMethod = &v
	return s
}

// SetBackupStartTime sets the BackupStartTime field's value.
func (s *BackupsInfoForDescribeBackupDetailOutput) SetBackupStartTime(v string) *BackupsInfoForDescribeBackupDetailOutput {
	s.BackupStartTime = &v
	return s
}

// SetBackupStatus sets the BackupStatus field's value.
func (s *BackupsInfoForDescribeBackupDetailOutput) SetBackupStatus(v string) *BackupsInfoForDescribeBackupDetailOutput {
	s.BackupStatus = &v
	return s
}

// SetBackupType sets the BackupType field's value.
func (s *BackupsInfoForDescribeBackupDetailOutput) SetBackupType(v string) *BackupsInfoForDescribeBackupDetailOutput {
	s.BackupType = &v
	return s
}

// SetCreateType sets the CreateType field's value.
func (s *BackupsInfoForDescribeBackupDetailOutput) SetCreateType(v string) *BackupsInfoForDescribeBackupDetailOutput {
	s.CreateType = &v
	return s
}

// SetDownloadProgress sets the DownloadProgress field's value.
func (s *BackupsInfoForDescribeBackupDetailOutput) SetDownloadProgress(v int32) *BackupsInfoForDescribeBackupDetailOutput {
	s.DownloadProgress = &v
	return s
}

// SetDownloadStatus sets the DownloadStatus field's value.
func (s *BackupsInfoForDescribeBackupDetailOutput) SetDownloadStatus(v string) *BackupsInfoForDescribeBackupDetailOutput {
	s.DownloadStatus = &v
	return s
}

type DescribeBackupDetailInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	// BackupId is a required field
	BackupId *string `type:"string" json:",omitempty" required:"true"`

	// InstanceId is a required field
	InstanceId *string `type:"string" json:",omitempty" required:"true"`
}

// String returns the string representation
func (s DescribeBackupDetailInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeBackupDetailInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeBackupDetailInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DescribeBackupDetailInput"}
	if s.BackupId == nil {
		invalidParams.Add(request.NewErrParamRequired("BackupId"))
	}
	if s.InstanceId == nil {
		invalidParams.Add(request.NewErrParamRequired("InstanceId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetBackupId sets the BackupId field's value.
func (s *DescribeBackupDetailInput) SetBackupId(v string) *DescribeBackupDetailInput {
	s.BackupId = &v
	return s
}

// SetInstanceId sets the InstanceId field's value.
func (s *DescribeBackupDetailInput) SetInstanceId(v string) *DescribeBackupDetailInput {
	s.InstanceId = &v
	return s
}

type DescribeBackupDetailOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	BackupsInfo *BackupsInfoForDescribeBackupDetailOutput `type:"structure" json:",omitempty"`
}

// String returns the string representation
func (s DescribeBackupDetailOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeBackupDetailOutput) GoString() string {
	return s.String()
}

// SetBackupsInfo sets the BackupsInfo field's value.
func (s *DescribeBackupDetailOutput) SetBackupsInfo(v *BackupsInfoForDescribeBackupDetailOutput) *DescribeBackupDetailOutput {
	s.BackupsInfo = v
	return s
}

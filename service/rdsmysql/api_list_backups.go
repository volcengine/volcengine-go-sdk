// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package rdsmysql

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opListBackupsCommon = "ListBackups"

// ListBackupsCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ListBackupsCommon operation. The "output" return
// value will be populated with the ListBackupsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListBackupsCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListBackupsCommon Send returns without error.
//
// See ListBackupsCommon for more information on using the ListBackupsCommon
// API call, and error handling.
//
//    // Example sending a request using the ListBackupsCommonRequest method.
//    req, resp := client.ListBackupsCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *RDSMYSQL) ListBackupsCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opListBackupsCommon,
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

// ListBackupsCommon API operation for RDS_MYSQL.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RDS_MYSQL's
// API operation ListBackupsCommon for usage and error information.
func (c *RDSMYSQL) ListBackupsCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ListBackupsCommonRequest(input)
	return out, req.Send()
}

// ListBackupsCommonWithContext is the same as ListBackupsCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ListBackupsCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RDSMYSQL) ListBackupsCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ListBackupsCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opListBackups = "ListBackups"

// ListBackupsRequest generates a "volcengine/request.Request" representing the
// client's request for the ListBackups operation. The "output" return
// value will be populated with the ListBackupsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListBackupsCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListBackupsCommon Send returns without error.
//
// See ListBackups for more information on using the ListBackups
// API call, and error handling.
//
//    // Example sending a request using the ListBackupsRequest method.
//    req, resp := client.ListBackupsRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *RDSMYSQL) ListBackupsRequest(input *ListBackupsInput) (req *request.Request, output *ListBackupsOutput) {
	op := &request.Operation{
		Name:       opListBackups,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListBackupsInput{}
	}

	output = &ListBackupsOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// ListBackups API operation for RDS_MYSQL.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RDS_MYSQL's
// API operation ListBackups for usage and error information.
func (c *RDSMYSQL) ListBackups(input *ListBackupsInput) (*ListBackupsOutput, error) {
	req, out := c.ListBackupsRequest(input)
	return out, req.Send()
}

// ListBackupsWithContext is the same as ListBackups with the addition of
// the ability to pass a context and additional request options.
//
// See ListBackups for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RDSMYSQL) ListBackupsWithContext(ctx volcengine.Context, input *ListBackupsInput, opts ...request.Option) (*ListBackupsOutput, error) {
	req, out := c.ListBackupsRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DataForListBackupsOutput struct {
	_ struct{} `type:"structure"`

	BackupEndTime *string `type:"string"`

	BackupFileName *string `type:"string"`

	BackupFileSize *int64 `type:"int64"`

	BackupId *string `type:"string"`

	BackupMode *string `type:"string"`

	BackupStartTime *string `type:"string"`

	BackupStatus *string `type:"string"`

	BackupStrategy *string `type:"string"`

	BackupType *string `type:"string"`

	CreateType *string `type:"string"`
}

// String returns the string representation
func (s DataForListBackupsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DataForListBackupsOutput) GoString() string {
	return s.String()
}

// SetBackupEndTime sets the BackupEndTime field's value.
func (s *DataForListBackupsOutput) SetBackupEndTime(v string) *DataForListBackupsOutput {
	s.BackupEndTime = &v
	return s
}

// SetBackupFileName sets the BackupFileName field's value.
func (s *DataForListBackupsOutput) SetBackupFileName(v string) *DataForListBackupsOutput {
	s.BackupFileName = &v
	return s
}

// SetBackupFileSize sets the BackupFileSize field's value.
func (s *DataForListBackupsOutput) SetBackupFileSize(v int64) *DataForListBackupsOutput {
	s.BackupFileSize = &v
	return s
}

// SetBackupId sets the BackupId field's value.
func (s *DataForListBackupsOutput) SetBackupId(v string) *DataForListBackupsOutput {
	s.BackupId = &v
	return s
}

// SetBackupMode sets the BackupMode field's value.
func (s *DataForListBackupsOutput) SetBackupMode(v string) *DataForListBackupsOutput {
	s.BackupMode = &v
	return s
}

// SetBackupStartTime sets the BackupStartTime field's value.
func (s *DataForListBackupsOutput) SetBackupStartTime(v string) *DataForListBackupsOutput {
	s.BackupStartTime = &v
	return s
}

// SetBackupStatus sets the BackupStatus field's value.
func (s *DataForListBackupsOutput) SetBackupStatus(v string) *DataForListBackupsOutput {
	s.BackupStatus = &v
	return s
}

// SetBackupStrategy sets the BackupStrategy field's value.
func (s *DataForListBackupsOutput) SetBackupStrategy(v string) *DataForListBackupsOutput {
	s.BackupStrategy = &v
	return s
}

// SetBackupType sets the BackupType field's value.
func (s *DataForListBackupsOutput) SetBackupType(v string) *DataForListBackupsOutput {
	s.BackupType = &v
	return s
}

// SetCreateType sets the CreateType field's value.
func (s *DataForListBackupsOutput) SetCreateType(v string) *DataForListBackupsOutput {
	s.CreateType = &v
	return s
}

type ListBackupsInput struct {
	_ struct{} `type:"structure"`

	BackupDataType *string `type:"string" enum:"EnumOfBackupDataTypeForListBackupsInput"`

	BackupStatus *string `type:"string" enum:"EnumOfBackupStatusForListBackupsInput"`

	EndTime *string `type:"string"`

	// InstanceId is a required field
	InstanceId *string `type:"string" required:"true"`

	Limit *int32 `type:"int32"`

	Offset *int32 `type:"int32"`

	StartTime *string `type:"string"`
}

// String returns the string representation
func (s ListBackupsInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListBackupsInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListBackupsInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ListBackupsInput"}
	if s.InstanceId == nil {
		invalidParams.Add(request.NewErrParamRequired("InstanceId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetBackupDataType sets the BackupDataType field's value.
func (s *ListBackupsInput) SetBackupDataType(v string) *ListBackupsInput {
	s.BackupDataType = &v
	return s
}

// SetBackupStatus sets the BackupStatus field's value.
func (s *ListBackupsInput) SetBackupStatus(v string) *ListBackupsInput {
	s.BackupStatus = &v
	return s
}

// SetEndTime sets the EndTime field's value.
func (s *ListBackupsInput) SetEndTime(v string) *ListBackupsInput {
	s.EndTime = &v
	return s
}

// SetInstanceId sets the InstanceId field's value.
func (s *ListBackupsInput) SetInstanceId(v string) *ListBackupsInput {
	s.InstanceId = &v
	return s
}

// SetLimit sets the Limit field's value.
func (s *ListBackupsInput) SetLimit(v int32) *ListBackupsInput {
	s.Limit = &v
	return s
}

// SetOffset sets the Offset field's value.
func (s *ListBackupsInput) SetOffset(v int32) *ListBackupsInput {
	s.Offset = &v
	return s
}

// SetStartTime sets the StartTime field's value.
func (s *ListBackupsInput) SetStartTime(v string) *ListBackupsInput {
	s.StartTime = &v
	return s
}

type ListBackupsOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	Datas []*DataForListBackupsOutput `type:"list"`

	Total *int32 `type:"int32"`
}

// String returns the string representation
func (s ListBackupsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListBackupsOutput) GoString() string {
	return s.String()
}

// SetDatas sets the Datas field's value.
func (s *ListBackupsOutput) SetDatas(v []*DataForListBackupsOutput) *ListBackupsOutput {
	s.Datas = v
	return s
}

// SetTotal sets the Total field's value.
func (s *ListBackupsOutput) SetTotal(v int32) *ListBackupsOutput {
	s.Total = &v
	return s
}

const (
	// EnumOfBackupDataTypeForListBackupsInputData is a EnumOfBackupDataTypeForListBackupsInput enum value
	EnumOfBackupDataTypeForListBackupsInputData = "Data"

	// EnumOfBackupDataTypeForListBackupsInputLog is a EnumOfBackupDataTypeForListBackupsInput enum value
	EnumOfBackupDataTypeForListBackupsInputLog = "Log"
)

const (
	// EnumOfBackupStatusForListBackupsInputFailed is a EnumOfBackupStatusForListBackupsInput enum value
	EnumOfBackupStatusForListBackupsInputFailed = "Failed"

	// EnumOfBackupStatusForListBackupsInputRunning is a EnumOfBackupStatusForListBackupsInput enum value
	EnumOfBackupStatusForListBackupsInputRunning = "Running"

	// EnumOfBackupStatusForListBackupsInputSuccess is a EnumOfBackupStatusForListBackupsInput enum value
	EnumOfBackupStatusForListBackupsInputSuccess = "Success"
)

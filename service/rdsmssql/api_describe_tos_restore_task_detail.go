// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package rdsmssql

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribeTosRestoreTaskDetailCommon = "DescribeTosRestoreTaskDetail"

// DescribeTosRestoreTaskDetailCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeTosRestoreTaskDetailCommon operation. The "output" return
// value will be populated with the DescribeTosRestoreTaskDetailCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeTosRestoreTaskDetailCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeTosRestoreTaskDetailCommon Send returns without error.
//
// See DescribeTosRestoreTaskDetailCommon for more information on using the DescribeTosRestoreTaskDetailCommon
// API call, and error handling.
//
//    // Example sending a request using the DescribeTosRestoreTaskDetailCommonRequest method.
//    req, resp := client.DescribeTosRestoreTaskDetailCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *RDSMSSQL) DescribeTosRestoreTaskDetailCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeTosRestoreTaskDetailCommon,
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

// DescribeTosRestoreTaskDetailCommon API operation for RDS_MSSQL.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RDS_MSSQL's
// API operation DescribeTosRestoreTaskDetailCommon for usage and error information.
func (c *RDSMSSQL) DescribeTosRestoreTaskDetailCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeTosRestoreTaskDetailCommonRequest(input)
	return out, req.Send()
}

// DescribeTosRestoreTaskDetailCommonWithContext is the same as DescribeTosRestoreTaskDetailCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeTosRestoreTaskDetailCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RDSMSSQL) DescribeTosRestoreTaskDetailCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeTosRestoreTaskDetailCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeTosRestoreTaskDetail = "DescribeTosRestoreTaskDetail"

// DescribeTosRestoreTaskDetailRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeTosRestoreTaskDetail operation. The "output" return
// value will be populated with the DescribeTosRestoreTaskDetailCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeTosRestoreTaskDetailCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeTosRestoreTaskDetailCommon Send returns without error.
//
// See DescribeTosRestoreTaskDetail for more information on using the DescribeTosRestoreTaskDetail
// API call, and error handling.
//
//    // Example sending a request using the DescribeTosRestoreTaskDetailRequest method.
//    req, resp := client.DescribeTosRestoreTaskDetailRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *RDSMSSQL) DescribeTosRestoreTaskDetailRequest(input *DescribeTosRestoreTaskDetailInput) (req *request.Request, output *DescribeTosRestoreTaskDetailOutput) {
	op := &request.Operation{
		Name:       opDescribeTosRestoreTaskDetail,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeTosRestoreTaskDetailInput{}
	}

	output = &DescribeTosRestoreTaskDetailOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// DescribeTosRestoreTaskDetail API operation for RDS_MSSQL.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RDS_MSSQL's
// API operation DescribeTosRestoreTaskDetail for usage and error information.
func (c *RDSMSSQL) DescribeTosRestoreTaskDetail(input *DescribeTosRestoreTaskDetailInput) (*DescribeTosRestoreTaskDetailOutput, error) {
	req, out := c.DescribeTosRestoreTaskDetailRequest(input)
	return out, req.Send()
}

// DescribeTosRestoreTaskDetailWithContext is the same as DescribeTosRestoreTaskDetail with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeTosRestoreTaskDetail for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RDSMSSQL) DescribeTosRestoreTaskDetailWithContext(ctx volcengine.Context, input *DescribeTosRestoreTaskDetailInput, opts ...request.Option) (*DescribeTosRestoreTaskDetailOutput, error) {
	req, out := c.DescribeTosRestoreTaskDetailRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DescribeTosRestoreTaskDetailInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	PageNumber *int32 `type:"int32" json:",omitempty"`

	PageSize *int32 `type:"int32" json:",omitempty"`

	ProjectName *string `type:"string" json:",omitempty"`

	// RestoreTaskId is a required field
	RestoreTaskId *string `type:"string" json:",omitempty" required:"true"`
}

// String returns the string representation
func (s DescribeTosRestoreTaskDetailInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeTosRestoreTaskDetailInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeTosRestoreTaskDetailInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DescribeTosRestoreTaskDetailInput"}
	if s.RestoreTaskId == nil {
		invalidParams.Add(request.NewErrParamRequired("RestoreTaskId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetPageNumber sets the PageNumber field's value.
func (s *DescribeTosRestoreTaskDetailInput) SetPageNumber(v int32) *DescribeTosRestoreTaskDetailInput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *DescribeTosRestoreTaskDetailInput) SetPageSize(v int32) *DescribeTosRestoreTaskDetailInput {
	s.PageSize = &v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *DescribeTosRestoreTaskDetailInput) SetProjectName(v string) *DescribeTosRestoreTaskDetailInput {
	s.ProjectName = &v
	return s
}

// SetRestoreTaskId sets the RestoreTaskId field's value.
func (s *DescribeTosRestoreTaskDetailInput) SetRestoreTaskId(v string) *DescribeTosRestoreTaskDetailInput {
	s.RestoreTaskId = &v
	return s
}

type DescribeTosRestoreTaskDetailOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	RestoreTaskDetails []*RestoreTaskDetailForDescribeTosRestoreTaskDetailOutput `type:"list" json:",omitempty"`

	Total *int32 `type:"int32" json:",omitempty"`
}

// String returns the string representation
func (s DescribeTosRestoreTaskDetailOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeTosRestoreTaskDetailOutput) GoString() string {
	return s.String()
}

// SetRestoreTaskDetails sets the RestoreTaskDetails field's value.
func (s *DescribeTosRestoreTaskDetailOutput) SetRestoreTaskDetails(v []*RestoreTaskDetailForDescribeTosRestoreTaskDetailOutput) *DescribeTosRestoreTaskDetailOutput {
	s.RestoreTaskDetails = v
	return s
}

// SetTotal sets the Total field's value.
func (s *DescribeTosRestoreTaskDetailOutput) SetTotal(v int32) *DescribeTosRestoreTaskDetailOutput {
	s.Total = &v
	return s
}

type RestoreTaskDetailForDescribeTosRestoreTaskDetailOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	BackupType *string `type:"string" json:",omitempty"`

	DBName *string `type:"string" json:",omitempty"`

	NewDBName *string `type:"string" json:",omitempty"`

	RestoreDesc *string `type:"string" json:",omitempty"`

	RestoreEndTime *string `type:"string" json:",omitempty"`

	RestoreFileName *string `type:"string" json:",omitempty"`

	RestoreFileSize *string `type:"string" json:",omitempty"`

	RestoreStartTime *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s RestoreTaskDetailForDescribeTosRestoreTaskDetailOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s RestoreTaskDetailForDescribeTosRestoreTaskDetailOutput) GoString() string {
	return s.String()
}

// SetBackupType sets the BackupType field's value.
func (s *RestoreTaskDetailForDescribeTosRestoreTaskDetailOutput) SetBackupType(v string) *RestoreTaskDetailForDescribeTosRestoreTaskDetailOutput {
	s.BackupType = &v
	return s
}

// SetDBName sets the DBName field's value.
func (s *RestoreTaskDetailForDescribeTosRestoreTaskDetailOutput) SetDBName(v string) *RestoreTaskDetailForDescribeTosRestoreTaskDetailOutput {
	s.DBName = &v
	return s
}

// SetNewDBName sets the NewDBName field's value.
func (s *RestoreTaskDetailForDescribeTosRestoreTaskDetailOutput) SetNewDBName(v string) *RestoreTaskDetailForDescribeTosRestoreTaskDetailOutput {
	s.NewDBName = &v
	return s
}

// SetRestoreDesc sets the RestoreDesc field's value.
func (s *RestoreTaskDetailForDescribeTosRestoreTaskDetailOutput) SetRestoreDesc(v string) *RestoreTaskDetailForDescribeTosRestoreTaskDetailOutput {
	s.RestoreDesc = &v
	return s
}

// SetRestoreEndTime sets the RestoreEndTime field's value.
func (s *RestoreTaskDetailForDescribeTosRestoreTaskDetailOutput) SetRestoreEndTime(v string) *RestoreTaskDetailForDescribeTosRestoreTaskDetailOutput {
	s.RestoreEndTime = &v
	return s
}

// SetRestoreFileName sets the RestoreFileName field's value.
func (s *RestoreTaskDetailForDescribeTosRestoreTaskDetailOutput) SetRestoreFileName(v string) *RestoreTaskDetailForDescribeTosRestoreTaskDetailOutput {
	s.RestoreFileName = &v
	return s
}

// SetRestoreFileSize sets the RestoreFileSize field's value.
func (s *RestoreTaskDetailForDescribeTosRestoreTaskDetailOutput) SetRestoreFileSize(v string) *RestoreTaskDetailForDescribeTosRestoreTaskDetailOutput {
	s.RestoreFileSize = &v
	return s
}

// SetRestoreStartTime sets the RestoreStartTime field's value.
func (s *RestoreTaskDetailForDescribeTosRestoreTaskDetailOutput) SetRestoreStartTime(v string) *RestoreTaskDetailForDescribeTosRestoreTaskDetailOutput {
	s.RestoreStartTime = &v
	return s
}

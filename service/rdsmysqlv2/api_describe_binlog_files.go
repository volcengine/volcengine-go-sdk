// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package rdsmysqlv2

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribeBinlogFilesCommon = "DescribeBinlogFiles"

// DescribeBinlogFilesCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeBinlogFilesCommon operation. The "output" return
// value will be populated with the DescribeBinlogFilesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeBinlogFilesCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeBinlogFilesCommon Send returns without error.
//
// See DescribeBinlogFilesCommon for more information on using the DescribeBinlogFilesCommon
// API call, and error handling.
//
//    // Example sending a request using the DescribeBinlogFilesCommonRequest method.
//    req, resp := client.DescribeBinlogFilesCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *RDSMYSQLV2) DescribeBinlogFilesCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeBinlogFilesCommon,
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

// DescribeBinlogFilesCommon API operation for RDS_MYSQL_V2.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RDS_MYSQL_V2's
// API operation DescribeBinlogFilesCommon for usage and error information.
func (c *RDSMYSQLV2) DescribeBinlogFilesCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeBinlogFilesCommonRequest(input)
	return out, req.Send()
}

// DescribeBinlogFilesCommonWithContext is the same as DescribeBinlogFilesCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeBinlogFilesCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RDSMYSQLV2) DescribeBinlogFilesCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeBinlogFilesCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeBinlogFiles = "DescribeBinlogFiles"

// DescribeBinlogFilesRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeBinlogFiles operation. The "output" return
// value will be populated with the DescribeBinlogFilesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeBinlogFilesCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeBinlogFilesCommon Send returns without error.
//
// See DescribeBinlogFiles for more information on using the DescribeBinlogFiles
// API call, and error handling.
//
//    // Example sending a request using the DescribeBinlogFilesRequest method.
//    req, resp := client.DescribeBinlogFilesRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *RDSMYSQLV2) DescribeBinlogFilesRequest(input *DescribeBinlogFilesInput) (req *request.Request, output *DescribeBinlogFilesOutput) {
	op := &request.Operation{
		Name:       opDescribeBinlogFiles,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeBinlogFilesInput{}
	}

	output = &DescribeBinlogFilesOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// DescribeBinlogFiles API operation for RDS_MYSQL_V2.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RDS_MYSQL_V2's
// API operation DescribeBinlogFiles for usage and error information.
func (c *RDSMYSQLV2) DescribeBinlogFiles(input *DescribeBinlogFilesInput) (*DescribeBinlogFilesOutput, error) {
	req, out := c.DescribeBinlogFilesRequest(input)
	return out, req.Send()
}

// DescribeBinlogFilesWithContext is the same as DescribeBinlogFiles with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeBinlogFiles for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RDSMYSQLV2) DescribeBinlogFilesWithContext(ctx volcengine.Context, input *DescribeBinlogFilesInput, opts ...request.Option) (*DescribeBinlogFilesOutput, error) {
	req, out := c.DescribeBinlogFilesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type BinlogFileForDescribeBinlogFilesOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	BackupRegion *string `type:"string" json:",omitempty"`

	BackupStatus *string `type:"string" json:",omitempty"`

	DownloadStatus *string `type:"string" json:",omitempty"`

	FileName *string `type:"string" json:",omitempty"`

	FileSize *int64 `type:"int64" json:",omitempty"`

	IsEncrypted *bool `type:"boolean" json:",omitempty"`

	NodeId *string `type:"string" json:",omitempty"`

	UpdateTime *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s BinlogFileForDescribeBinlogFilesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s BinlogFileForDescribeBinlogFilesOutput) GoString() string {
	return s.String()
}

// SetBackupRegion sets the BackupRegion field's value.
func (s *BinlogFileForDescribeBinlogFilesOutput) SetBackupRegion(v string) *BinlogFileForDescribeBinlogFilesOutput {
	s.BackupRegion = &v
	return s
}

// SetBackupStatus sets the BackupStatus field's value.
func (s *BinlogFileForDescribeBinlogFilesOutput) SetBackupStatus(v string) *BinlogFileForDescribeBinlogFilesOutput {
	s.BackupStatus = &v
	return s
}

// SetDownloadStatus sets the DownloadStatus field's value.
func (s *BinlogFileForDescribeBinlogFilesOutput) SetDownloadStatus(v string) *BinlogFileForDescribeBinlogFilesOutput {
	s.DownloadStatus = &v
	return s
}

// SetFileName sets the FileName field's value.
func (s *BinlogFileForDescribeBinlogFilesOutput) SetFileName(v string) *BinlogFileForDescribeBinlogFilesOutput {
	s.FileName = &v
	return s
}

// SetFileSize sets the FileSize field's value.
func (s *BinlogFileForDescribeBinlogFilesOutput) SetFileSize(v int64) *BinlogFileForDescribeBinlogFilesOutput {
	s.FileSize = &v
	return s
}

// SetIsEncrypted sets the IsEncrypted field's value.
func (s *BinlogFileForDescribeBinlogFilesOutput) SetIsEncrypted(v bool) *BinlogFileForDescribeBinlogFilesOutput {
	s.IsEncrypted = &v
	return s
}

// SetNodeId sets the NodeId field's value.
func (s *BinlogFileForDescribeBinlogFilesOutput) SetNodeId(v string) *BinlogFileForDescribeBinlogFilesOutput {
	s.NodeId = &v
	return s
}

// SetUpdateTime sets the UpdateTime field's value.
func (s *BinlogFileForDescribeBinlogFilesOutput) SetUpdateTime(v string) *BinlogFileForDescribeBinlogFilesOutput {
	s.UpdateTime = &v
	return s
}

type DescribeBinlogFilesInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	BackupFileNumber *int32 `type:"int32" json:",omitempty"`

	Context *string `type:"string" json:",omitempty"`

	// EndTime is a required field
	EndTime *string `type:"string" json:",omitempty" required:"true"`

	// InstanceId is a required field
	InstanceId *string `type:"string" json:",omitempty" required:"true"`

	// StartTime is a required field
	StartTime *string `type:"string" json:",omitempty" required:"true"`
}

// String returns the string representation
func (s DescribeBinlogFilesInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeBinlogFilesInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeBinlogFilesInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DescribeBinlogFilesInput"}
	if s.EndTime == nil {
		invalidParams.Add(request.NewErrParamRequired("EndTime"))
	}
	if s.InstanceId == nil {
		invalidParams.Add(request.NewErrParamRequired("InstanceId"))
	}
	if s.StartTime == nil {
		invalidParams.Add(request.NewErrParamRequired("StartTime"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetBackupFileNumber sets the BackupFileNumber field's value.
func (s *DescribeBinlogFilesInput) SetBackupFileNumber(v int32) *DescribeBinlogFilesInput {
	s.BackupFileNumber = &v
	return s
}

// SetContext sets the Context field's value.
func (s *DescribeBinlogFilesInput) SetContext(v string) *DescribeBinlogFilesInput {
	s.Context = &v
	return s
}

// SetEndTime sets the EndTime field's value.
func (s *DescribeBinlogFilesInput) SetEndTime(v string) *DescribeBinlogFilesInput {
	s.EndTime = &v
	return s
}

// SetInstanceId sets the InstanceId field's value.
func (s *DescribeBinlogFilesInput) SetInstanceId(v string) *DescribeBinlogFilesInput {
	s.InstanceId = &v
	return s
}

// SetStartTime sets the StartTime field's value.
func (s *DescribeBinlogFilesInput) SetStartTime(v string) *DescribeBinlogFilesInput {
	s.StartTime = &v
	return s
}

type DescribeBinlogFilesOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	BinlogFiles []*BinlogFileForDescribeBinlogFilesOutput `type:"list" json:",omitempty"`

	Context *string `type:"string" json:",omitempty"`

	Total *int32 `type:"int32" json:",omitempty"`
}

// String returns the string representation
func (s DescribeBinlogFilesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeBinlogFilesOutput) GoString() string {
	return s.String()
}

// SetBinlogFiles sets the BinlogFiles field's value.
func (s *DescribeBinlogFilesOutput) SetBinlogFiles(v []*BinlogFileForDescribeBinlogFilesOutput) *DescribeBinlogFilesOutput {
	s.BinlogFiles = v
	return s
}

// SetContext sets the Context field's value.
func (s *DescribeBinlogFilesOutput) SetContext(v string) *DescribeBinlogFilesOutput {
	s.Context = &v
	return s
}

// SetTotal sets the Total field's value.
func (s *DescribeBinlogFilesOutput) SetTotal(v int32) *DescribeBinlogFilesOutput {
	s.Total = &v
	return s
}

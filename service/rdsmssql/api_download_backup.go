// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package rdsmssql

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDownloadBackupCommon = "DownloadBackup"

// DownloadBackupCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DownloadBackupCommon operation. The "output" return
// value will be populated with the DownloadBackupCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DownloadBackupCommon Request to send the API call to the service.
// the "output" return value is not valid until after DownloadBackupCommon Send returns without error.
//
// See DownloadBackupCommon for more information on using the DownloadBackupCommon
// API call, and error handling.
//
//    // Example sending a request using the DownloadBackupCommonRequest method.
//    req, resp := client.DownloadBackupCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *RDSMSSQL) DownloadBackupCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDownloadBackupCommon,
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

// DownloadBackupCommon API operation for RDS_MSSQL.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RDS_MSSQL's
// API operation DownloadBackupCommon for usage and error information.
func (c *RDSMSSQL) DownloadBackupCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DownloadBackupCommonRequest(input)
	return out, req.Send()
}

// DownloadBackupCommonWithContext is the same as DownloadBackupCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DownloadBackupCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RDSMSSQL) DownloadBackupCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DownloadBackupCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDownloadBackup = "DownloadBackup"

// DownloadBackupRequest generates a "volcengine/request.Request" representing the
// client's request for the DownloadBackup operation. The "output" return
// value will be populated with the DownloadBackupCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DownloadBackupCommon Request to send the API call to the service.
// the "output" return value is not valid until after DownloadBackupCommon Send returns without error.
//
// See DownloadBackup for more information on using the DownloadBackup
// API call, and error handling.
//
//    // Example sending a request using the DownloadBackupRequest method.
//    req, resp := client.DownloadBackupRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *RDSMSSQL) DownloadBackupRequest(input *DownloadBackupInput) (req *request.Request, output *DownloadBackupOutput) {
	op := &request.Operation{
		Name:       opDownloadBackup,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DownloadBackupInput{}
	}

	output = &DownloadBackupOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// DownloadBackup API operation for RDS_MSSQL.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RDS_MSSQL's
// API operation DownloadBackup for usage and error information.
func (c *RDSMSSQL) DownloadBackup(input *DownloadBackupInput) (*DownloadBackupOutput, error) {
	req, out := c.DownloadBackupRequest(input)
	return out, req.Send()
}

// DownloadBackupWithContext is the same as DownloadBackup with the addition of
// the ability to pass a context and additional request options.
//
// See DownloadBackup for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RDSMSSQL) DownloadBackupWithContext(ctx volcengine.Context, input *DownloadBackupInput, opts ...request.Option) (*DownloadBackupOutput, error) {
	req, out := c.DownloadBackupRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DownloadBackupInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	// BackupId is a required field
	BackupId *string `type:"string" json:",omitempty" required:"true"`

	// InstanceId is a required field
	InstanceId *string `type:"string" json:",omitempty" required:"true"`
}

// String returns the string representation
func (s DownloadBackupInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DownloadBackupInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DownloadBackupInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DownloadBackupInput"}
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
func (s *DownloadBackupInput) SetBackupId(v string) *DownloadBackupInput {
	s.BackupId = &v
	return s
}

// SetInstanceId sets the InstanceId field's value.
func (s *DownloadBackupInput) SetInstanceId(v string) *DownloadBackupInput {
	s.InstanceId = &v
	return s
}

type DownloadBackupOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	BackupId *string `type:"string" json:",omitempty"`

	InstanceId *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s DownloadBackupOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DownloadBackupOutput) GoString() string {
	return s.String()
}

// SetBackupId sets the BackupId field's value.
func (s *DownloadBackupOutput) SetBackupId(v string) *DownloadBackupOutput {
	s.BackupId = &v
	return s
}

// SetInstanceId sets the InstanceId field's value.
func (s *DownloadBackupOutput) SetInstanceId(v string) *DownloadBackupOutput {
	s.InstanceId = &v
	return s
}

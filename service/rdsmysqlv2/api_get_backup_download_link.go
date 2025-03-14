// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package rdsmysqlv2

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opGetBackupDownloadLinkCommon = "GetBackupDownloadLink"

// GetBackupDownloadLinkCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the GetBackupDownloadLinkCommon operation. The "output" return
// value will be populated with the GetBackupDownloadLinkCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned GetBackupDownloadLinkCommon Request to send the API call to the service.
// the "output" return value is not valid until after GetBackupDownloadLinkCommon Send returns without error.
//
// See GetBackupDownloadLinkCommon for more information on using the GetBackupDownloadLinkCommon
// API call, and error handling.
//
//    // Example sending a request using the GetBackupDownloadLinkCommonRequest method.
//    req, resp := client.GetBackupDownloadLinkCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *RDSMYSQLV2) GetBackupDownloadLinkCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opGetBackupDownloadLinkCommon,
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

// GetBackupDownloadLinkCommon API operation for RDS_MYSQL_V2.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RDS_MYSQL_V2's
// API operation GetBackupDownloadLinkCommon for usage and error information.
func (c *RDSMYSQLV2) GetBackupDownloadLinkCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.GetBackupDownloadLinkCommonRequest(input)
	return out, req.Send()
}

// GetBackupDownloadLinkCommonWithContext is the same as GetBackupDownloadLinkCommon with the addition of
// the ability to pass a context and additional request options.
//
// See GetBackupDownloadLinkCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RDSMYSQLV2) GetBackupDownloadLinkCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.GetBackupDownloadLinkCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opGetBackupDownloadLink = "GetBackupDownloadLink"

// GetBackupDownloadLinkRequest generates a "volcengine/request.Request" representing the
// client's request for the GetBackupDownloadLink operation. The "output" return
// value will be populated with the GetBackupDownloadLinkCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned GetBackupDownloadLinkCommon Request to send the API call to the service.
// the "output" return value is not valid until after GetBackupDownloadLinkCommon Send returns without error.
//
// See GetBackupDownloadLink for more information on using the GetBackupDownloadLink
// API call, and error handling.
//
//    // Example sending a request using the GetBackupDownloadLinkRequest method.
//    req, resp := client.GetBackupDownloadLinkRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *RDSMYSQLV2) GetBackupDownloadLinkRequest(input *GetBackupDownloadLinkInput) (req *request.Request, output *GetBackupDownloadLinkOutput) {
	op := &request.Operation{
		Name:       opGetBackupDownloadLink,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetBackupDownloadLinkInput{}
	}

	output = &GetBackupDownloadLinkOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// GetBackupDownloadLink API operation for RDS_MYSQL_V2.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RDS_MYSQL_V2's
// API operation GetBackupDownloadLink for usage and error information.
func (c *RDSMYSQLV2) GetBackupDownloadLink(input *GetBackupDownloadLinkInput) (*GetBackupDownloadLinkOutput, error) {
	req, out := c.GetBackupDownloadLinkRequest(input)
	return out, req.Send()
}

// GetBackupDownloadLinkWithContext is the same as GetBackupDownloadLink with the addition of
// the ability to pass a context and additional request options.
//
// See GetBackupDownloadLink for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RDSMYSQLV2) GetBackupDownloadLinkWithContext(ctx volcengine.Context, input *GetBackupDownloadLinkInput, opts ...request.Option) (*GetBackupDownloadLinkOutput, error) {
	req, out := c.GetBackupDownloadLinkRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type GetBackupDownloadLinkInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	// BackupId is a required field
	BackupId *string `type:"string" json:",omitempty" required:"true"`

	// InstanceId is a required field
	InstanceId *string `type:"string" json:",omitempty" required:"true"`
}

// String returns the string representation
func (s GetBackupDownloadLinkInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s GetBackupDownloadLinkInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetBackupDownloadLinkInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "GetBackupDownloadLinkInput"}
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
func (s *GetBackupDownloadLinkInput) SetBackupId(v string) *GetBackupDownloadLinkInput {
	s.BackupId = &v
	return s
}

// SetInstanceId sets the InstanceId field's value.
func (s *GetBackupDownloadLinkInput) SetInstanceId(v string) *GetBackupDownloadLinkInput {
	s.InstanceId = &v
	return s
}

type GetBackupDownloadLinkOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	BackupDownloadLink *string `type:"string" json:",omitempty"`

	BackupFileName *string `type:"string" json:",omitempty"`

	BackupFileSize *int64 `type:"int64" json:",omitempty"`

	BackupId *string `type:"string" json:",omitempty"`

	BackupType *string `type:"string" json:",omitempty"`

	DownloadProgress *int64 `type:"int64" json:",omitempty"`

	InnerBackupDownloadLink *string `type:"string" json:",omitempty"`

	InstanceId *string `type:"string" json:",omitempty"`

	LinkExpiredTime *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s GetBackupDownloadLinkOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s GetBackupDownloadLinkOutput) GoString() string {
	return s.String()
}

// SetBackupDownloadLink sets the BackupDownloadLink field's value.
func (s *GetBackupDownloadLinkOutput) SetBackupDownloadLink(v string) *GetBackupDownloadLinkOutput {
	s.BackupDownloadLink = &v
	return s
}

// SetBackupFileName sets the BackupFileName field's value.
func (s *GetBackupDownloadLinkOutput) SetBackupFileName(v string) *GetBackupDownloadLinkOutput {
	s.BackupFileName = &v
	return s
}

// SetBackupFileSize sets the BackupFileSize field's value.
func (s *GetBackupDownloadLinkOutput) SetBackupFileSize(v int64) *GetBackupDownloadLinkOutput {
	s.BackupFileSize = &v
	return s
}

// SetBackupId sets the BackupId field's value.
func (s *GetBackupDownloadLinkOutput) SetBackupId(v string) *GetBackupDownloadLinkOutput {
	s.BackupId = &v
	return s
}

// SetBackupType sets the BackupType field's value.
func (s *GetBackupDownloadLinkOutput) SetBackupType(v string) *GetBackupDownloadLinkOutput {
	s.BackupType = &v
	return s
}

// SetDownloadProgress sets the DownloadProgress field's value.
func (s *GetBackupDownloadLinkOutput) SetDownloadProgress(v int64) *GetBackupDownloadLinkOutput {
	s.DownloadProgress = &v
	return s
}

// SetInnerBackupDownloadLink sets the InnerBackupDownloadLink field's value.
func (s *GetBackupDownloadLinkOutput) SetInnerBackupDownloadLink(v string) *GetBackupDownloadLinkOutput {
	s.InnerBackupDownloadLink = &v
	return s
}

// SetInstanceId sets the InstanceId field's value.
func (s *GetBackupDownloadLinkOutput) SetInstanceId(v string) *GetBackupDownloadLinkOutput {
	s.InstanceId = &v
	return s
}

// SetLinkExpiredTime sets the LinkExpiredTime field's value.
func (s *GetBackupDownloadLinkOutput) SetLinkExpiredTime(v string) *GetBackupDownloadLinkOutput {
	s.LinkExpiredTime = &v
	return s
}

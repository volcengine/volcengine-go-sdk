// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package dns

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opListUserZoneBackupsCommon = "ListUserZoneBackups"

// ListUserZoneBackupsCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ListUserZoneBackupsCommon operation. The "output" return
// value will be populated with the ListUserZoneBackupsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListUserZoneBackupsCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListUserZoneBackupsCommon Send returns without error.
//
// See ListUserZoneBackupsCommon for more information on using the ListUserZoneBackupsCommon
// API call, and error handling.
//
//    // Example sending a request using the ListUserZoneBackupsCommonRequest method.
//    req, resp := client.ListUserZoneBackupsCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *DNS) ListUserZoneBackupsCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opListUserZoneBackupsCommon,
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

// ListUserZoneBackupsCommon API operation for DNS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for DNS's
// API operation ListUserZoneBackupsCommon for usage and error information.
func (c *DNS) ListUserZoneBackupsCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ListUserZoneBackupsCommonRequest(input)
	return out, req.Send()
}

// ListUserZoneBackupsCommonWithContext is the same as ListUserZoneBackupsCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ListUserZoneBackupsCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *DNS) ListUserZoneBackupsCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ListUserZoneBackupsCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opListUserZoneBackups = "ListUserZoneBackups"

// ListUserZoneBackupsRequest generates a "volcengine/request.Request" representing the
// client's request for the ListUserZoneBackups operation. The "output" return
// value will be populated with the ListUserZoneBackupsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListUserZoneBackupsCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListUserZoneBackupsCommon Send returns without error.
//
// See ListUserZoneBackups for more information on using the ListUserZoneBackups
// API call, and error handling.
//
//    // Example sending a request using the ListUserZoneBackupsRequest method.
//    req, resp := client.ListUserZoneBackupsRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *DNS) ListUserZoneBackupsRequest(input *ListUserZoneBackupsInput) (req *request.Request, output *ListUserZoneBackupsOutput) {
	op := &request.Operation{
		Name:       opListUserZoneBackups,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListUserZoneBackupsInput{}
	}

	output = &ListUserZoneBackupsOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// ListUserZoneBackups API operation for DNS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for DNS's
// API operation ListUserZoneBackups for usage and error information.
func (c *DNS) ListUserZoneBackups(input *ListUserZoneBackupsInput) (*ListUserZoneBackupsOutput, error) {
	req, out := c.ListUserZoneBackupsRequest(input)
	return out, req.Send()
}

// ListUserZoneBackupsWithContext is the same as ListUserZoneBackups with the addition of
// the ability to pass a context and additional request options.
//
// See ListUserZoneBackups for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *DNS) ListUserZoneBackupsWithContext(ctx volcengine.Context, input *ListUserZoneBackupsInput, opts ...request.Option) (*ListUserZoneBackupsOutput, error) {
	req, out := c.ListUserZoneBackupsRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type BackupInfoForListUserZoneBackupsOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	BackupID *string `type:"string" json:",omitempty"`

	BackupTime *string `type:"string" json:",omitempty"`

	RecordCount *int32 `type:"int32" json:",omitempty"`
}

// String returns the string representation
func (s BackupInfoForListUserZoneBackupsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s BackupInfoForListUserZoneBackupsOutput) GoString() string {
	return s.String()
}

// SetBackupID sets the BackupID field's value.
func (s *BackupInfoForListUserZoneBackupsOutput) SetBackupID(v string) *BackupInfoForListUserZoneBackupsOutput {
	s.BackupID = &v
	return s
}

// SetBackupTime sets the BackupTime field's value.
func (s *BackupInfoForListUserZoneBackupsOutput) SetBackupTime(v string) *BackupInfoForListUserZoneBackupsOutput {
	s.BackupTime = &v
	return s
}

// SetRecordCount sets the RecordCount field's value.
func (s *BackupInfoForListUserZoneBackupsOutput) SetRecordCount(v int32) *BackupInfoForListUserZoneBackupsOutput {
	s.RecordCount = &v
	return s
}

type ListUserZoneBackupsInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	// ZID is a required field
	ZID *int64 `type:"int64" json:",omitempty" required:"true"`
}

// String returns the string representation
func (s ListUserZoneBackupsInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListUserZoneBackupsInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListUserZoneBackupsInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ListUserZoneBackupsInput"}
	if s.ZID == nil {
		invalidParams.Add(request.NewErrParamRequired("ZID"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetZID sets the ZID field's value.
func (s *ListUserZoneBackupsInput) SetZID(v int64) *ListUserZoneBackupsInput {
	s.ZID = &v
	return s
}

type ListUserZoneBackupsOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	BackupInfos []*BackupInfoForListUserZoneBackupsOutput `type:"list" json:",omitempty"`
}

// String returns the string representation
func (s ListUserZoneBackupsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListUserZoneBackupsOutput) GoString() string {
	return s.String()
}

// SetBackupInfos sets the BackupInfos field's value.
func (s *ListUserZoneBackupsOutput) SetBackupInfos(v []*BackupInfoForListUserZoneBackupsOutput) *ListUserZoneBackupsOutput {
	s.BackupInfos = v
	return s
}
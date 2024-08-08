// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package rdsmysqlv2

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribeBackupEncryptionStatusCommon = "DescribeBackupEncryptionStatus"

// DescribeBackupEncryptionStatusCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeBackupEncryptionStatusCommon operation. The "output" return
// value will be populated with the DescribeBackupEncryptionStatusCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeBackupEncryptionStatusCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeBackupEncryptionStatusCommon Send returns without error.
//
// See DescribeBackupEncryptionStatusCommon for more information on using the DescribeBackupEncryptionStatusCommon
// API call, and error handling.
//
//    // Example sending a request using the DescribeBackupEncryptionStatusCommonRequest method.
//    req, resp := client.DescribeBackupEncryptionStatusCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *RDSMYSQLV2) DescribeBackupEncryptionStatusCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeBackupEncryptionStatusCommon,
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

// DescribeBackupEncryptionStatusCommon API operation for RDS_MYSQL_V2.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RDS_MYSQL_V2's
// API operation DescribeBackupEncryptionStatusCommon for usage and error information.
func (c *RDSMYSQLV2) DescribeBackupEncryptionStatusCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeBackupEncryptionStatusCommonRequest(input)
	return out, req.Send()
}

// DescribeBackupEncryptionStatusCommonWithContext is the same as DescribeBackupEncryptionStatusCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeBackupEncryptionStatusCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RDSMYSQLV2) DescribeBackupEncryptionStatusCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeBackupEncryptionStatusCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeBackupEncryptionStatus = "DescribeBackupEncryptionStatus"

// DescribeBackupEncryptionStatusRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeBackupEncryptionStatus operation. The "output" return
// value will be populated with the DescribeBackupEncryptionStatusCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeBackupEncryptionStatusCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeBackupEncryptionStatusCommon Send returns without error.
//
// See DescribeBackupEncryptionStatus for more information on using the DescribeBackupEncryptionStatus
// API call, and error handling.
//
//    // Example sending a request using the DescribeBackupEncryptionStatusRequest method.
//    req, resp := client.DescribeBackupEncryptionStatusRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *RDSMYSQLV2) DescribeBackupEncryptionStatusRequest(input *DescribeBackupEncryptionStatusInput) (req *request.Request, output *DescribeBackupEncryptionStatusOutput) {
	op := &request.Operation{
		Name:       opDescribeBackupEncryptionStatus,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeBackupEncryptionStatusInput{}
	}

	output = &DescribeBackupEncryptionStatusOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// DescribeBackupEncryptionStatus API operation for RDS_MYSQL_V2.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RDS_MYSQL_V2's
// API operation DescribeBackupEncryptionStatus for usage and error information.
func (c *RDSMYSQLV2) DescribeBackupEncryptionStatus(input *DescribeBackupEncryptionStatusInput) (*DescribeBackupEncryptionStatusOutput, error) {
	req, out := c.DescribeBackupEncryptionStatusRequest(input)
	return out, req.Send()
}

// DescribeBackupEncryptionStatusWithContext is the same as DescribeBackupEncryptionStatus with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeBackupEncryptionStatus for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RDSMYSQLV2) DescribeBackupEncryptionStatusWithContext(ctx volcengine.Context, input *DescribeBackupEncryptionStatusInput, opts ...request.Option) (*DescribeBackupEncryptionStatusOutput, error) {
	req, out := c.DescribeBackupEncryptionStatusRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DescribeBackupEncryptionStatusInput struct {
	_ struct{} `type:"structure"`

	// InstanceId is a required field
	InstanceId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeBackupEncryptionStatusInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeBackupEncryptionStatusInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeBackupEncryptionStatusInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DescribeBackupEncryptionStatusInput"}
	if s.InstanceId == nil {
		invalidParams.Add(request.NewErrParamRequired("InstanceId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetInstanceId sets the InstanceId field's value.
func (s *DescribeBackupEncryptionStatusInput) SetInstanceId(v string) *DescribeBackupEncryptionStatusInput {
	s.InstanceId = &v
	return s
}

type DescribeBackupEncryptionStatusOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	DataEncryptionStatus *string `type:"string"`

	LogEncryptionStatus *string `type:"string"`
}

// String returns the string representation
func (s DescribeBackupEncryptionStatusOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeBackupEncryptionStatusOutput) GoString() string {
	return s.String()
}

// SetDataEncryptionStatus sets the DataEncryptionStatus field's value.
func (s *DescribeBackupEncryptionStatusOutput) SetDataEncryptionStatus(v string) *DescribeBackupEncryptionStatusOutput {
	s.DataEncryptionStatus = &v
	return s
}

// SetLogEncryptionStatus sets the LogEncryptionStatus field's value.
func (s *DescribeBackupEncryptionStatusOutput) SetLogEncryptionStatus(v string) *DescribeBackupEncryptionStatusOutput {
	s.LogEncryptionStatus = &v
	return s
}

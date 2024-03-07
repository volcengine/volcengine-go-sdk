// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package rdsmssql

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opRestoreToExistedInstanceCommon = "RestoreToExistedInstance"

// RestoreToExistedInstanceCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the RestoreToExistedInstanceCommon operation. The "output" return
// value will be populated with the RestoreToExistedInstanceCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned RestoreToExistedInstanceCommon Request to send the API call to the service.
// the "output" return value is not valid until after RestoreToExistedInstanceCommon Send returns without error.
//
// See RestoreToExistedInstanceCommon for more information on using the RestoreToExistedInstanceCommon
// API call, and error handling.
//
//	// Example sending a request using the RestoreToExistedInstanceCommonRequest method.
//	req, resp := client.RestoreToExistedInstanceCommonRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *RDSMSSQL) RestoreToExistedInstanceCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opRestoreToExistedInstanceCommon,
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

// RestoreToExistedInstanceCommon API operation for RDS_MSSQL.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RDS_MSSQL's
// API operation RestoreToExistedInstanceCommon for usage and error information.
func (c *RDSMSSQL) RestoreToExistedInstanceCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.RestoreToExistedInstanceCommonRequest(input)
	return out, req.Send()
}

// RestoreToExistedInstanceCommonWithContext is the same as RestoreToExistedInstanceCommon with the addition of
// the ability to pass a context and additional request options.
//
// See RestoreToExistedInstanceCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RDSMSSQL) RestoreToExistedInstanceCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.RestoreToExistedInstanceCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opRestoreToExistedInstance = "RestoreToExistedInstance"

// RestoreToExistedInstanceRequest generates a "volcengine/request.Request" representing the
// client's request for the RestoreToExistedInstance operation. The "output" return
// value will be populated with the RestoreToExistedInstanceCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned RestoreToExistedInstanceCommon Request to send the API call to the service.
// the "output" return value is not valid until after RestoreToExistedInstanceCommon Send returns without error.
//
// See RestoreToExistedInstance for more information on using the RestoreToExistedInstance
// API call, and error handling.
//
//	// Example sending a request using the RestoreToExistedInstanceRequest method.
//	req, resp := client.RestoreToExistedInstanceRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *RDSMSSQL) RestoreToExistedInstanceRequest(input *RestoreToExistedInstanceInput) (req *request.Request, output *RestoreToExistedInstanceOutput) {
	op := &request.Operation{
		Name:       opRestoreToExistedInstance,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &RestoreToExistedInstanceInput{}
	}

	output = &RestoreToExistedInstanceOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// RestoreToExistedInstance API operation for RDS_MSSQL.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RDS_MSSQL's
// API operation RestoreToExistedInstance for usage and error information.
func (c *RDSMSSQL) RestoreToExistedInstance(input *RestoreToExistedInstanceInput) (*RestoreToExistedInstanceOutput, error) {
	req, out := c.RestoreToExistedInstanceRequest(input)
	return out, req.Send()
}

// RestoreToExistedInstanceWithContext is the same as RestoreToExistedInstance with the addition of
// the ability to pass a context and additional request options.
//
// See RestoreToExistedInstance for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RDSMSSQL) RestoreToExistedInstanceWithContext(ctx volcengine.Context, input *RestoreToExistedInstanceInput, opts ...request.Option) (*RestoreToExistedInstanceOutput, error) {
	req, out := c.RestoreToExistedInstanceRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DatabaseForRestoreToExistedInstanceInput struct {
	_ struct{} `type:"structure"`

	DBName *string `type:"string"`

	NewDBName *string `type:"string"`
}

// String returns the string representation
func (s DatabaseForRestoreToExistedInstanceInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DatabaseForRestoreToExistedInstanceInput) GoString() string {
	return s.String()
}

// SetDBName sets the DBName field's value.
func (s *DatabaseForRestoreToExistedInstanceInput) SetDBName(v string) *DatabaseForRestoreToExistedInstanceInput {
	s.DBName = &v
	return s
}

// SetNewDBName sets the NewDBName field's value.
func (s *DatabaseForRestoreToExistedInstanceInput) SetNewDBName(v string) *DatabaseForRestoreToExistedInstanceInput {
	s.NewDBName = &v
	return s
}

type RestoreToExistedInstanceInput struct {
	_ struct{} `type:"structure"`

	BackupId *string `type:"string"`

	Databases []*DatabaseForRestoreToExistedInstanceInput `type:"list"`

	RestoreTime *string `type:"string"`

	// SourceDBInstanceId is a required field
	SourceDBInstanceId *string `type:"string" required:"true"`

	// TargetDBInstanceId is a required field
	TargetDBInstanceId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s RestoreToExistedInstanceInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s RestoreToExistedInstanceInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RestoreToExistedInstanceInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "RestoreToExistedInstanceInput"}
	if s.SourceDBInstanceId == nil {
		invalidParams.Add(request.NewErrParamRequired("SourceDBInstanceId"))
	}
	if s.TargetDBInstanceId == nil {
		invalidParams.Add(request.NewErrParamRequired("TargetDBInstanceId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetBackupId sets the BackupId field's value.
func (s *RestoreToExistedInstanceInput) SetBackupId(v string) *RestoreToExistedInstanceInput {
	s.BackupId = &v
	return s
}

// SetDatabases sets the Databases field's value.
func (s *RestoreToExistedInstanceInput) SetDatabases(v []*DatabaseForRestoreToExistedInstanceInput) *RestoreToExistedInstanceInput {
	s.Databases = v
	return s
}

// SetRestoreTime sets the RestoreTime field's value.
func (s *RestoreToExistedInstanceInput) SetRestoreTime(v string) *RestoreToExistedInstanceInput {
	s.RestoreTime = &v
	return s
}

// SetSourceDBInstanceId sets the SourceDBInstanceId field's value.
func (s *RestoreToExistedInstanceInput) SetSourceDBInstanceId(v string) *RestoreToExistedInstanceInput {
	s.SourceDBInstanceId = &v
	return s
}

// SetTargetDBInstanceId sets the TargetDBInstanceId field's value.
func (s *RestoreToExistedInstanceInput) SetTargetDBInstanceId(v string) *RestoreToExistedInstanceInput {
	s.TargetDBInstanceId = &v
	return s
}

type RestoreToExistedInstanceOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	InstanceId *string `type:"string"`
}

// String returns the string representation
func (s RestoreToExistedInstanceOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s RestoreToExistedInstanceOutput) GoString() string {
	return s.String()
}

// SetInstanceId sets the InstanceId field's value.
func (s *RestoreToExistedInstanceOutput) SetInstanceId(v string) *RestoreToExistedInstanceOutput {
	s.InstanceId = &v
	return s
}

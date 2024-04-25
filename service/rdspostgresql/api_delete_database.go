// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package rdspostgresql

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDeleteDatabaseCommon = "DeleteDatabase"

// DeleteDatabaseCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DeleteDatabaseCommon operation. The "output" return
// value will be populated with the DeleteDatabaseCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeleteDatabaseCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeleteDatabaseCommon Send returns without error.
//
// See DeleteDatabaseCommon for more information on using the DeleteDatabaseCommon
// API call, and error handling.
//
//    // Example sending a request using the DeleteDatabaseCommonRequest method.
//    req, resp := client.DeleteDatabaseCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *RDSPOSTGRESQL) DeleteDatabaseCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDeleteDatabaseCommon,
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

// DeleteDatabaseCommon API operation for RDS_POSTGRESQL.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RDS_POSTGRESQL's
// API operation DeleteDatabaseCommon for usage and error information.
func (c *RDSPOSTGRESQL) DeleteDatabaseCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DeleteDatabaseCommonRequest(input)
	return out, req.Send()
}

// DeleteDatabaseCommonWithContext is the same as DeleteDatabaseCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteDatabaseCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RDSPOSTGRESQL) DeleteDatabaseCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DeleteDatabaseCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDeleteDatabase = "DeleteDatabase"

// DeleteDatabaseRequest generates a "volcengine/request.Request" representing the
// client's request for the DeleteDatabase operation. The "output" return
// value will be populated with the DeleteDatabaseCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeleteDatabaseCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeleteDatabaseCommon Send returns without error.
//
// See DeleteDatabase for more information on using the DeleteDatabase
// API call, and error handling.
//
//    // Example sending a request using the DeleteDatabaseRequest method.
//    req, resp := client.DeleteDatabaseRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *RDSPOSTGRESQL) DeleteDatabaseRequest(input *DeleteDatabaseInput) (req *request.Request, output *DeleteDatabaseOutput) {
	op := &request.Operation{
		Name:       opDeleteDatabase,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteDatabaseInput{}
	}

	output = &DeleteDatabaseOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// DeleteDatabase API operation for RDS_POSTGRESQL.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RDS_POSTGRESQL's
// API operation DeleteDatabase for usage and error information.
func (c *RDSPOSTGRESQL) DeleteDatabase(input *DeleteDatabaseInput) (*DeleteDatabaseOutput, error) {
	req, out := c.DeleteDatabaseRequest(input)
	return out, req.Send()
}

// DeleteDatabaseWithContext is the same as DeleteDatabase with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteDatabase for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RDSPOSTGRESQL) DeleteDatabaseWithContext(ctx volcengine.Context, input *DeleteDatabaseInput, opts ...request.Option) (*DeleteDatabaseOutput, error) {
	req, out := c.DeleteDatabaseRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DeleteDatabaseInput struct {
	_ struct{} `type:"structure"`

	// DBName is a required field
	DBName *string `type:"string" required:"true"`

	// InstanceId is a required field
	InstanceId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteDatabaseInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteDatabaseInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteDatabaseInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DeleteDatabaseInput"}
	if s.DBName == nil {
		invalidParams.Add(request.NewErrParamRequired("DBName"))
	}
	if s.InstanceId == nil {
		invalidParams.Add(request.NewErrParamRequired("InstanceId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetDBName sets the DBName field's value.
func (s *DeleteDatabaseInput) SetDBName(v string) *DeleteDatabaseInput {
	s.DBName = &v
	return s
}

// SetInstanceId sets the InstanceId field's value.
func (s *DeleteDatabaseInput) SetInstanceId(v string) *DeleteDatabaseInput {
	s.InstanceId = &v
	return s
}

type DeleteDatabaseOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s DeleteDatabaseOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteDatabaseOutput) GoString() string {
	return s.String()
}

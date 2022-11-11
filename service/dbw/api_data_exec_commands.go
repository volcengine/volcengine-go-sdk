// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package dbw

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDataExecCommandsCommon = "DataExecCommands"

// DataExecCommandsCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DataExecCommandsCommon operation. The "output" return
// value will be populated with the DataExecCommandsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DataExecCommandsCommon Request to send the API call to the service.
// the "output" return value is not valid until after DataExecCommandsCommon Send returns without error.
//
// See DataExecCommandsCommon for more information on using the DataExecCommandsCommon
// API call, and error handling.
//
//    // Example sending a request using the DataExecCommandsCommonRequest method.
//    req, resp := client.DataExecCommandsCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *DBW) DataExecCommandsCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDataExecCommandsCommon,
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

// DataExecCommandsCommon API operation for DBW.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for DBW's
// API operation DataExecCommandsCommon for usage and error information.
func (c *DBW) DataExecCommandsCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DataExecCommandsCommonRequest(input)
	return out, req.Send()
}

// DataExecCommandsCommonWithContext is the same as DataExecCommandsCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DataExecCommandsCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *DBW) DataExecCommandsCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DataExecCommandsCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDataExecCommands = "DataExecCommands"

// DataExecCommandsRequest generates a "volcengine/request.Request" representing the
// client's request for the DataExecCommands operation. The "output" return
// value will be populated with the DataExecCommandsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DataExecCommandsCommon Request to send the API call to the service.
// the "output" return value is not valid until after DataExecCommandsCommon Send returns without error.
//
// See DataExecCommands for more information on using the DataExecCommands
// API call, and error handling.
//
//    // Example sending a request using the DataExecCommandsRequest method.
//    req, resp := client.DataExecCommandsRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *DBW) DataExecCommandsRequest(input *DataExecCommandsInput) (req *request.Request, output *DataExecCommandsOutput) {
	op := &request.Operation{
		Name:       opDataExecCommands,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DataExecCommandsInput{}
	}

	output = &DataExecCommandsOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// DataExecCommands API operation for DBW.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for DBW's
// API operation DataExecCommands for usage and error information.
func (c *DBW) DataExecCommands(input *DataExecCommandsInput) (*DataExecCommandsOutput, error) {
	req, out := c.DataExecCommandsRequest(input)
	return out, req.Send()
}

// DataExecCommandsWithContext is the same as DataExecCommands with the addition of
// the ability to pass a context and additional request options.
//
// See DataExecCommands for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *DBW) DataExecCommandsWithContext(ctx volcengine.Context, input *DataExecCommandsInput, opts ...request.Option) (*DataExecCommandsOutput, error) {
	req, out := c.DataExecCommandsRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DataExecCommandsInput struct {
	_ struct{} `type:"structure"`

	Commands *string `type:"string"`

	DatabaseName *string `type:"string"`

	SessionId *string `type:"string"`

	TimeOutSeconds *int64 `type:"int64"`
}

// String returns the string representation
func (s DataExecCommandsInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DataExecCommandsInput) GoString() string {
	return s.String()
}

// SetCommands sets the Commands field's value.
func (s *DataExecCommandsInput) SetCommands(v string) *DataExecCommandsInput {
	s.Commands = &v
	return s
}

// SetDatabaseName sets the DatabaseName field's value.
func (s *DataExecCommandsInput) SetDatabaseName(v string) *DataExecCommandsInput {
	s.DatabaseName = &v
	return s
}

// SetSessionId sets the SessionId field's value.
func (s *DataExecCommandsInput) SetSessionId(v string) *DataExecCommandsInput {
	s.SessionId = &v
	return s
}

// SetTimeOutSeconds sets the TimeOutSeconds field's value.
func (s *DataExecCommandsInput) SetTimeOutSeconds(v int64) *DataExecCommandsInput {
	s.TimeOutSeconds = &v
	return s
}

type DataExecCommandsOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	Results []*ResultForDataExecCommandsOutput `type:"list"`
}

// String returns the string representation
func (s DataExecCommandsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DataExecCommandsOutput) GoString() string {
	return s.String()
}

// SetResults sets the Results field's value.
func (s *DataExecCommandsOutput) SetResults(v []*ResultForDataExecCommandsOutput) *DataExecCommandsOutput {
	s.Results = v
	return s
}

type ResultForDataExecCommandsOutput struct {
	_ struct{} `type:"structure"`

	ColumnNames []*string `type:"list"`

	CommandStr *string `type:"string"`

	ReasonDetail *string `type:"string"`

	RowCount *int32 `type:"int32"`

	Rows []*RowForDataExecCommandsOutput `type:"list"`

	RunTime *int64 `type:"int64"`

	State *string `type:"string" enum:"EnumOfStateForDataExecCommandsOutput"`
}

// String returns the string representation
func (s ResultForDataExecCommandsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ResultForDataExecCommandsOutput) GoString() string {
	return s.String()
}

// SetColumnNames sets the ColumnNames field's value.
func (s *ResultForDataExecCommandsOutput) SetColumnNames(v []*string) *ResultForDataExecCommandsOutput {
	s.ColumnNames = v
	return s
}

// SetCommandStr sets the CommandStr field's value.
func (s *ResultForDataExecCommandsOutput) SetCommandStr(v string) *ResultForDataExecCommandsOutput {
	s.CommandStr = &v
	return s
}

// SetReasonDetail sets the ReasonDetail field's value.
func (s *ResultForDataExecCommandsOutput) SetReasonDetail(v string) *ResultForDataExecCommandsOutput {
	s.ReasonDetail = &v
	return s
}

// SetRowCount sets the RowCount field's value.
func (s *ResultForDataExecCommandsOutput) SetRowCount(v int32) *ResultForDataExecCommandsOutput {
	s.RowCount = &v
	return s
}

// SetRows sets the Rows field's value.
func (s *ResultForDataExecCommandsOutput) SetRows(v []*RowForDataExecCommandsOutput) *ResultForDataExecCommandsOutput {
	s.Rows = v
	return s
}

// SetRunTime sets the RunTime field's value.
func (s *ResultForDataExecCommandsOutput) SetRunTime(v int64) *ResultForDataExecCommandsOutput {
	s.RunTime = &v
	return s
}

// SetState sets the State field's value.
func (s *ResultForDataExecCommandsOutput) SetState(v string) *ResultForDataExecCommandsOutput {
	s.State = &v
	return s
}

type RowForDataExecCommandsOutput struct {
	_ struct{} `type:"structure"`

	Cells []*string `type:"list"`
}

// String returns the string representation
func (s RowForDataExecCommandsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s RowForDataExecCommandsOutput) GoString() string {
	return s.String()
}

// SetCells sets the Cells field's value.
func (s *RowForDataExecCommandsOutput) SetCells(v []*string) *RowForDataExecCommandsOutput {
	s.Cells = v
	return s
}

const (
	// EnumOfStateForDataExecCommandsOutputCancel is a EnumOfStateForDataExecCommandsOutput enum value
	EnumOfStateForDataExecCommandsOutputCancel = "Cancel"

	// EnumOfStateForDataExecCommandsOutputFailed is a EnumOfStateForDataExecCommandsOutput enum value
	EnumOfStateForDataExecCommandsOutputFailed = "Failed"

	// EnumOfStateForDataExecCommandsOutputSuccess is a EnumOfStateForDataExecCommandsOutput enum value
	EnumOfStateForDataExecCommandsOutputSuccess = "Success"
)

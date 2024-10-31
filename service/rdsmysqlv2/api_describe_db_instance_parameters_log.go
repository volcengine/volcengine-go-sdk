// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package rdsmysqlv2

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribeDBInstanceParametersLogCommon = "DescribeDBInstanceParametersLog"

// DescribeDBInstanceParametersLogCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeDBInstanceParametersLogCommon operation. The "output" return
// value will be populated with the DescribeDBInstanceParametersLogCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeDBInstanceParametersLogCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeDBInstanceParametersLogCommon Send returns without error.
//
// See DescribeDBInstanceParametersLogCommon for more information on using the DescribeDBInstanceParametersLogCommon
// API call, and error handling.
//
//    // Example sending a request using the DescribeDBInstanceParametersLogCommonRequest method.
//    req, resp := client.DescribeDBInstanceParametersLogCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *RDSMYSQLV2) DescribeDBInstanceParametersLogCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeDBInstanceParametersLogCommon,
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

// DescribeDBInstanceParametersLogCommon API operation for RDS_MYSQL_V2.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RDS_MYSQL_V2's
// API operation DescribeDBInstanceParametersLogCommon for usage and error information.
func (c *RDSMYSQLV2) DescribeDBInstanceParametersLogCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeDBInstanceParametersLogCommonRequest(input)
	return out, req.Send()
}

// DescribeDBInstanceParametersLogCommonWithContext is the same as DescribeDBInstanceParametersLogCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeDBInstanceParametersLogCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RDSMYSQLV2) DescribeDBInstanceParametersLogCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeDBInstanceParametersLogCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeDBInstanceParametersLog = "DescribeDBInstanceParametersLog"

// DescribeDBInstanceParametersLogRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeDBInstanceParametersLog operation. The "output" return
// value will be populated with the DescribeDBInstanceParametersLogCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeDBInstanceParametersLogCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeDBInstanceParametersLogCommon Send returns without error.
//
// See DescribeDBInstanceParametersLog for more information on using the DescribeDBInstanceParametersLog
// API call, and error handling.
//
//    // Example sending a request using the DescribeDBInstanceParametersLogRequest method.
//    req, resp := client.DescribeDBInstanceParametersLogRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *RDSMYSQLV2) DescribeDBInstanceParametersLogRequest(input *DescribeDBInstanceParametersLogInput) (req *request.Request, output *DescribeDBInstanceParametersLogOutput) {
	op := &request.Operation{
		Name:       opDescribeDBInstanceParametersLog,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeDBInstanceParametersLogInput{}
	}

	output = &DescribeDBInstanceParametersLogOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// DescribeDBInstanceParametersLog API operation for RDS_MYSQL_V2.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RDS_MYSQL_V2's
// API operation DescribeDBInstanceParametersLog for usage and error information.
func (c *RDSMYSQLV2) DescribeDBInstanceParametersLog(input *DescribeDBInstanceParametersLogInput) (*DescribeDBInstanceParametersLogOutput, error) {
	req, out := c.DescribeDBInstanceParametersLogRequest(input)
	return out, req.Send()
}

// DescribeDBInstanceParametersLogWithContext is the same as DescribeDBInstanceParametersLog with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeDBInstanceParametersLog for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RDSMYSQLV2) DescribeDBInstanceParametersLogWithContext(ctx volcengine.Context, input *DescribeDBInstanceParametersLogInput, opts ...request.Option) (*DescribeDBInstanceParametersLogOutput, error) {
	req, out := c.DescribeDBInstanceParametersLogRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DescribeDBInstanceParametersLogInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	// EndTime is a required field
	EndTime *string `type:"string" json:",omitempty" required:"true"`

	// InstanceId is a required field
	InstanceId *string `type:"string" json:",omitempty" required:"true"`

	// StartTime is a required field
	StartTime *string `type:"string" json:",omitempty" required:"true"`
}

// String returns the string representation
func (s DescribeDBInstanceParametersLogInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeDBInstanceParametersLogInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeDBInstanceParametersLogInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DescribeDBInstanceParametersLogInput"}
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

// SetEndTime sets the EndTime field's value.
func (s *DescribeDBInstanceParametersLogInput) SetEndTime(v string) *DescribeDBInstanceParametersLogInput {
	s.EndTime = &v
	return s
}

// SetInstanceId sets the InstanceId field's value.
func (s *DescribeDBInstanceParametersLogInput) SetInstanceId(v string) *DescribeDBInstanceParametersLogInput {
	s.InstanceId = &v
	return s
}

// SetStartTime sets the StartTime field's value.
func (s *DescribeDBInstanceParametersLogInput) SetStartTime(v string) *DescribeDBInstanceParametersLogInput {
	s.StartTime = &v
	return s
}

type DescribeDBInstanceParametersLogOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	ParameterChangeLogs []*ParameterChangeLogForDescribeDBInstanceParametersLogOutput `type:"list" json:",omitempty"`

	Total *int32 `type:"int32" json:",omitempty"`
}

// String returns the string representation
func (s DescribeDBInstanceParametersLogOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeDBInstanceParametersLogOutput) GoString() string {
	return s.String()
}

// SetParameterChangeLogs sets the ParameterChangeLogs field's value.
func (s *DescribeDBInstanceParametersLogOutput) SetParameterChangeLogs(v []*ParameterChangeLogForDescribeDBInstanceParametersLogOutput) *DescribeDBInstanceParametersLogOutput {
	s.ParameterChangeLogs = v
	return s
}

// SetTotal sets the Total field's value.
func (s *DescribeDBInstanceParametersLogOutput) SetTotal(v int32) *DescribeDBInstanceParametersLogOutput {
	s.Total = &v
	return s
}

type ParameterChangeLogForDescribeDBInstanceParametersLogOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	CustomNodeIds []*string `type:"list" json:",omitempty"`

	ModifyTime *string `type:"string" json:",omitempty"`

	NewParameterValue *string `type:"string" json:",omitempty"`

	OldParameterValue *string `type:"string" json:",omitempty"`

	ParamApplyScope *string `type:"string" json:",omitempty" enum:"EnumOfParamApplyScopeForDescribeDBInstanceParametersLogOutput"`

	ParameterName *string `type:"string" json:",omitempty"`

	Status *string `type:"string" json:",omitempty" enum:"EnumOfStatusForDescribeDBInstanceParametersLogOutput"`
}

// String returns the string representation
func (s ParameterChangeLogForDescribeDBInstanceParametersLogOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ParameterChangeLogForDescribeDBInstanceParametersLogOutput) GoString() string {
	return s.String()
}

// SetCustomNodeIds sets the CustomNodeIds field's value.
func (s *ParameterChangeLogForDescribeDBInstanceParametersLogOutput) SetCustomNodeIds(v []*string) *ParameterChangeLogForDescribeDBInstanceParametersLogOutput {
	s.CustomNodeIds = v
	return s
}

// SetModifyTime sets the ModifyTime field's value.
func (s *ParameterChangeLogForDescribeDBInstanceParametersLogOutput) SetModifyTime(v string) *ParameterChangeLogForDescribeDBInstanceParametersLogOutput {
	s.ModifyTime = &v
	return s
}

// SetNewParameterValue sets the NewParameterValue field's value.
func (s *ParameterChangeLogForDescribeDBInstanceParametersLogOutput) SetNewParameterValue(v string) *ParameterChangeLogForDescribeDBInstanceParametersLogOutput {
	s.NewParameterValue = &v
	return s
}

// SetOldParameterValue sets the OldParameterValue field's value.
func (s *ParameterChangeLogForDescribeDBInstanceParametersLogOutput) SetOldParameterValue(v string) *ParameterChangeLogForDescribeDBInstanceParametersLogOutput {
	s.OldParameterValue = &v
	return s
}

// SetParamApplyScope sets the ParamApplyScope field's value.
func (s *ParameterChangeLogForDescribeDBInstanceParametersLogOutput) SetParamApplyScope(v string) *ParameterChangeLogForDescribeDBInstanceParametersLogOutput {
	s.ParamApplyScope = &v
	return s
}

// SetParameterName sets the ParameterName field's value.
func (s *ParameterChangeLogForDescribeDBInstanceParametersLogOutput) SetParameterName(v string) *ParameterChangeLogForDescribeDBInstanceParametersLogOutput {
	s.ParameterName = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *ParameterChangeLogForDescribeDBInstanceParametersLogOutput) SetStatus(v string) *ParameterChangeLogForDescribeDBInstanceParametersLogOutput {
	s.Status = &v
	return s
}

const (
	// EnumOfParamApplyScopeForDescribeDBInstanceParametersLogOutputAllNode is a EnumOfParamApplyScopeForDescribeDBInstanceParametersLogOutput enum value
	EnumOfParamApplyScopeForDescribeDBInstanceParametersLogOutputAllNode = "AllNode"

	// EnumOfParamApplyScopeForDescribeDBInstanceParametersLogOutputOnlyMasterSlaveNode is a EnumOfParamApplyScopeForDescribeDBInstanceParametersLogOutput enum value
	EnumOfParamApplyScopeForDescribeDBInstanceParametersLogOutputOnlyMasterSlaveNode = "OnlyMasterSlaveNode"

	// EnumOfParamApplyScopeForDescribeDBInstanceParametersLogOutputOnlyReadOnlyNode is a EnumOfParamApplyScopeForDescribeDBInstanceParametersLogOutput enum value
	EnumOfParamApplyScopeForDescribeDBInstanceParametersLogOutputOnlyReadOnlyNode = "OnlyReadOnlyNode"

	// EnumOfParamApplyScopeForDescribeDBInstanceParametersLogOutputCustomNode is a EnumOfParamApplyScopeForDescribeDBInstanceParametersLogOutput enum value
	EnumOfParamApplyScopeForDescribeDBInstanceParametersLogOutputCustomNode = "CustomNode"
)

const (
	// EnumOfStatusForDescribeDBInstanceParametersLogOutputApplied is a EnumOfStatusForDescribeDBInstanceParametersLogOutput enum value
	EnumOfStatusForDescribeDBInstanceParametersLogOutputApplied = "Applied"

	// EnumOfStatusForDescribeDBInstanceParametersLogOutputInvalid is a EnumOfStatusForDescribeDBInstanceParametersLogOutput enum value
	EnumOfStatusForDescribeDBInstanceParametersLogOutputInvalid = "Invalid"

	// EnumOfStatusForDescribeDBInstanceParametersLogOutputSyncing is a EnumOfStatusForDescribeDBInstanceParametersLogOutput enum value
	EnumOfStatusForDescribeDBInstanceParametersLogOutputSyncing = "Syncing"
)

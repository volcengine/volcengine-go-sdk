// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package mcs

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opGetAlarmSyncTaskCommon = "GetAlarmSyncTask"

// GetAlarmSyncTaskCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the GetAlarmSyncTaskCommon operation. The "output" return
// value will be populated with the GetAlarmSyncTaskCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned GetAlarmSyncTaskCommon Request to send the API call to the service.
// the "output" return value is not valid until after GetAlarmSyncTaskCommon Send returns without error.
//
// See GetAlarmSyncTaskCommon for more information on using the GetAlarmSyncTaskCommon
// API call, and error handling.
//
//    // Example sending a request using the GetAlarmSyncTaskCommonRequest method.
//    req, resp := client.GetAlarmSyncTaskCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *MCS) GetAlarmSyncTaskCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opGetAlarmSyncTaskCommon,
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

// GetAlarmSyncTaskCommon API operation for MCS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for MCS's
// API operation GetAlarmSyncTaskCommon for usage and error information.
func (c *MCS) GetAlarmSyncTaskCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.GetAlarmSyncTaskCommonRequest(input)
	return out, req.Send()
}

// GetAlarmSyncTaskCommonWithContext is the same as GetAlarmSyncTaskCommon with the addition of
// the ability to pass a context and additional request options.
//
// See GetAlarmSyncTaskCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *MCS) GetAlarmSyncTaskCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.GetAlarmSyncTaskCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opGetAlarmSyncTask = "GetAlarmSyncTask"

// GetAlarmSyncTaskRequest generates a "volcengine/request.Request" representing the
// client's request for the GetAlarmSyncTask operation. The "output" return
// value will be populated with the GetAlarmSyncTaskCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned GetAlarmSyncTaskCommon Request to send the API call to the service.
// the "output" return value is not valid until after GetAlarmSyncTaskCommon Send returns without error.
//
// See GetAlarmSyncTask for more information on using the GetAlarmSyncTask
// API call, and error handling.
//
//    // Example sending a request using the GetAlarmSyncTaskRequest method.
//    req, resp := client.GetAlarmSyncTaskRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *MCS) GetAlarmSyncTaskRequest(input *GetAlarmSyncTaskInput) (req *request.Request, output *GetAlarmSyncTaskOutput) {
	op := &request.Operation{
		Name:       opGetAlarmSyncTask,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetAlarmSyncTaskInput{}
	}

	output = &GetAlarmSyncTaskOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// GetAlarmSyncTask API operation for MCS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for MCS's
// API operation GetAlarmSyncTask for usage and error information.
func (c *MCS) GetAlarmSyncTask(input *GetAlarmSyncTaskInput) (*GetAlarmSyncTaskOutput, error) {
	req, out := c.GetAlarmSyncTaskRequest(input)
	return out, req.Send()
}

// GetAlarmSyncTaskWithContext is the same as GetAlarmSyncTask with the addition of
// the ability to pass a context and additional request options.
//
// See GetAlarmSyncTask for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *MCS) GetAlarmSyncTaskWithContext(ctx volcengine.Context, input *GetAlarmSyncTaskInput, opts ...request.Option) (*GetAlarmSyncTaskOutput, error) {
	req, out := c.GetAlarmSyncTaskRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type BaseInfoForGetAlarmSyncTaskOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	EndTimeMilli *int64 `type:"int64" json:",omitempty"`

	ErrCode *string `type:"string" json:",omitempty"`

	ErrMsg *string `type:"string" json:",omitempty"`

	StartTimeMilli *int64 `type:"int64" json:",omitempty"`

	TaskID *string `type:"string" json:",omitempty"`

	TaskStatus *string `type:"string" json:",omitempty" enum:"EnumOfTaskStatusForGetAlarmSyncTaskOutput"`

	TriggerType *string `type:"string" json:",omitempty" enum:"EnumOfTriggerTypeForGetAlarmSyncTaskOutput"`
}

// String returns the string representation
func (s BaseInfoForGetAlarmSyncTaskOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s BaseInfoForGetAlarmSyncTaskOutput) GoString() string {
	return s.String()
}

// SetEndTimeMilli sets the EndTimeMilli field's value.
func (s *BaseInfoForGetAlarmSyncTaskOutput) SetEndTimeMilli(v int64) *BaseInfoForGetAlarmSyncTaskOutput {
	s.EndTimeMilli = &v
	return s
}

// SetErrCode sets the ErrCode field's value.
func (s *BaseInfoForGetAlarmSyncTaskOutput) SetErrCode(v string) *BaseInfoForGetAlarmSyncTaskOutput {
	s.ErrCode = &v
	return s
}

// SetErrMsg sets the ErrMsg field's value.
func (s *BaseInfoForGetAlarmSyncTaskOutput) SetErrMsg(v string) *BaseInfoForGetAlarmSyncTaskOutput {
	s.ErrMsg = &v
	return s
}

// SetStartTimeMilli sets the StartTimeMilli field's value.
func (s *BaseInfoForGetAlarmSyncTaskOutput) SetStartTimeMilli(v int64) *BaseInfoForGetAlarmSyncTaskOutput {
	s.StartTimeMilli = &v
	return s
}

// SetTaskID sets the TaskID field's value.
func (s *BaseInfoForGetAlarmSyncTaskOutput) SetTaskID(v string) *BaseInfoForGetAlarmSyncTaskOutput {
	s.TaskID = &v
	return s
}

// SetTaskStatus sets the TaskStatus field's value.
func (s *BaseInfoForGetAlarmSyncTaskOutput) SetTaskStatus(v string) *BaseInfoForGetAlarmSyncTaskOutput {
	s.TaskStatus = &v
	return s
}

// SetTriggerType sets the TriggerType field's value.
func (s *BaseInfoForGetAlarmSyncTaskOutput) SetTriggerType(v string) *BaseInfoForGetAlarmSyncTaskOutput {
	s.TriggerType = &v
	return s
}

type GetAlarmSyncTaskInput struct {
	_ struct{} `type:"structure" json:",omitempty"`
}

// String returns the string representation
func (s GetAlarmSyncTaskInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s GetAlarmSyncTaskInput) GoString() string {
	return s.String()
}

type GetAlarmSyncTaskOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	BaseInfo *BaseInfoForGetAlarmSyncTaskOutput `type:"structure" json:",omitempty"`
}

// String returns the string representation
func (s GetAlarmSyncTaskOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s GetAlarmSyncTaskOutput) GoString() string {
	return s.String()
}

// SetBaseInfo sets the BaseInfo field's value.
func (s *GetAlarmSyncTaskOutput) SetBaseInfo(v *BaseInfoForGetAlarmSyncTaskOutput) *GetAlarmSyncTaskOutput {
	s.BaseInfo = v
	return s
}

const (
	// EnumOfTaskStatusForGetAlarmSyncTaskOutputWaiting is a EnumOfTaskStatusForGetAlarmSyncTaskOutput enum value
	EnumOfTaskStatusForGetAlarmSyncTaskOutputWaiting = "waiting"

	// EnumOfTaskStatusForGetAlarmSyncTaskOutputProcessing is a EnumOfTaskStatusForGetAlarmSyncTaskOutput enum value
	EnumOfTaskStatusForGetAlarmSyncTaskOutputProcessing = "processing"

	// EnumOfTaskStatusForGetAlarmSyncTaskOutputSucceeded is a EnumOfTaskStatusForGetAlarmSyncTaskOutput enum value
	EnumOfTaskStatusForGetAlarmSyncTaskOutputSucceeded = "succeeded"

	// EnumOfTaskStatusForGetAlarmSyncTaskOutputCanceled is a EnumOfTaskStatusForGetAlarmSyncTaskOutput enum value
	EnumOfTaskStatusForGetAlarmSyncTaskOutputCanceled = "canceled"

	// EnumOfTaskStatusForGetAlarmSyncTaskOutputFailed is a EnumOfTaskStatusForGetAlarmSyncTaskOutput enum value
	EnumOfTaskStatusForGetAlarmSyncTaskOutputFailed = "failed"

	// EnumOfTaskStatusForGetAlarmSyncTaskOutputTimeouted is a EnumOfTaskStatusForGetAlarmSyncTaskOutput enum value
	EnumOfTaskStatusForGetAlarmSyncTaskOutputTimeouted = "timeouted"
)

const (
	// EnumOfTriggerTypeForGetAlarmSyncTaskOutputByUser is a EnumOfTriggerTypeForGetAlarmSyncTaskOutput enum value
	EnumOfTriggerTypeForGetAlarmSyncTaskOutputByUser = "by_user"

	// EnumOfTriggerTypeForGetAlarmSyncTaskOutputByTimer is a EnumOfTriggerTypeForGetAlarmSyncTaskOutput enum value
	EnumOfTriggerTypeForGetAlarmSyncTaskOutputByTimer = "by_timer"
)

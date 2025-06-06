// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package seccenter20240508

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opGetRaspAlarmStatisticsCommon = "GetRaspAlarmStatistics"

// GetRaspAlarmStatisticsCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the GetRaspAlarmStatisticsCommon operation. The "output" return
// value will be populated with the GetRaspAlarmStatisticsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned GetRaspAlarmStatisticsCommon Request to send the API call to the service.
// the "output" return value is not valid until after GetRaspAlarmStatisticsCommon Send returns without error.
//
// See GetRaspAlarmStatisticsCommon for more information on using the GetRaspAlarmStatisticsCommon
// API call, and error handling.
//
//    // Example sending a request using the GetRaspAlarmStatisticsCommonRequest method.
//    req, resp := client.GetRaspAlarmStatisticsCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *SECCENTER20240508) GetRaspAlarmStatisticsCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opGetRaspAlarmStatisticsCommon,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &map[string]interface{}{}
	}

	output = &map[string]interface{}{}
	req = c.newRequest(op, input, output)

	return
}

// GetRaspAlarmStatisticsCommon API operation for SECCENTER20240508.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for SECCENTER20240508's
// API operation GetRaspAlarmStatisticsCommon for usage and error information.
func (c *SECCENTER20240508) GetRaspAlarmStatisticsCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.GetRaspAlarmStatisticsCommonRequest(input)
	return out, req.Send()
}

// GetRaspAlarmStatisticsCommonWithContext is the same as GetRaspAlarmStatisticsCommon with the addition of
// the ability to pass a context and additional request options.
//
// See GetRaspAlarmStatisticsCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *SECCENTER20240508) GetRaspAlarmStatisticsCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.GetRaspAlarmStatisticsCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opGetRaspAlarmStatistics = "GetRaspAlarmStatistics"

// GetRaspAlarmStatisticsRequest generates a "volcengine/request.Request" representing the
// client's request for the GetRaspAlarmStatistics operation. The "output" return
// value will be populated with the GetRaspAlarmStatisticsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned GetRaspAlarmStatisticsCommon Request to send the API call to the service.
// the "output" return value is not valid until after GetRaspAlarmStatisticsCommon Send returns without error.
//
// See GetRaspAlarmStatistics for more information on using the GetRaspAlarmStatistics
// API call, and error handling.
//
//    // Example sending a request using the GetRaspAlarmStatisticsRequest method.
//    req, resp := client.GetRaspAlarmStatisticsRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *SECCENTER20240508) GetRaspAlarmStatisticsRequest(input *GetRaspAlarmStatisticsInput) (req *request.Request, output *GetRaspAlarmStatisticsOutput) {
	op := &request.Operation{
		Name:       opGetRaspAlarmStatistics,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetRaspAlarmStatisticsInput{}
	}

	output = &GetRaspAlarmStatisticsOutput{}
	req = c.newRequest(op, input, output)

	return
}

// GetRaspAlarmStatistics API operation for SECCENTER20240508.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for SECCENTER20240508's
// API operation GetRaspAlarmStatistics for usage and error information.
func (c *SECCENTER20240508) GetRaspAlarmStatistics(input *GetRaspAlarmStatisticsInput) (*GetRaspAlarmStatisticsOutput, error) {
	req, out := c.GetRaspAlarmStatisticsRequest(input)
	return out, req.Send()
}

// GetRaspAlarmStatisticsWithContext is the same as GetRaspAlarmStatistics with the addition of
// the ability to pass a context and additional request options.
//
// See GetRaspAlarmStatistics for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *SECCENTER20240508) GetRaspAlarmStatisticsWithContext(ctx volcengine.Context, input *GetRaspAlarmStatisticsInput, opts ...request.Option) (*GetRaspAlarmStatisticsOutput, error) {
	req, out := c.GetRaspAlarmStatisticsRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type GetRaspAlarmStatisticsInput struct {
	_ struct{} `type:"structure"`

	AgentID *string `type:"string"`

	ClusterID *string `type:"string"`

	EndTime *int32 `type:"int32"`

	MlpInstanceID *string `type:"string"`

	StartTime *int32 `type:"int32"`
}

// String returns the string representation
func (s GetRaspAlarmStatisticsInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s GetRaspAlarmStatisticsInput) GoString() string {
	return s.String()
}

// SetAgentID sets the AgentID field's value.
func (s *GetRaspAlarmStatisticsInput) SetAgentID(v string) *GetRaspAlarmStatisticsInput {
	s.AgentID = &v
	return s
}

// SetClusterID sets the ClusterID field's value.
func (s *GetRaspAlarmStatisticsInput) SetClusterID(v string) *GetRaspAlarmStatisticsInput {
	s.ClusterID = &v
	return s
}

// SetEndTime sets the EndTime field's value.
func (s *GetRaspAlarmStatisticsInput) SetEndTime(v int32) *GetRaspAlarmStatisticsInput {
	s.EndTime = &v
	return s
}

// SetMlpInstanceID sets the MlpInstanceID field's value.
func (s *GetRaspAlarmStatisticsInput) SetMlpInstanceID(v string) *GetRaspAlarmStatisticsInput {
	s.MlpInstanceID = &v
	return s
}

// SetStartTime sets the StartTime field's value.
func (s *GetRaspAlarmStatisticsInput) SetStartTime(v int32) *GetRaspAlarmStatisticsInput {
	s.StartTime = &v
	return s
}

type GetRaspAlarmStatisticsOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	AlarmCriticalCount *int32 `type:"int32"`

	AlarmHighCount *int32 `type:"int32"`

	AlarmLowCount *int32 `type:"int32"`

	AlarmMediumCount *int32 `type:"int32"`

	AlarmProcessedCount *int32 `type:"int32"`

	AlarmTotal *int32 `type:"int32"`

	AlarmWhiteCount *int32 `type:"int32"`

	IsolateFileCount *int32 `type:"int32"`

	IsolateTodayCount *int32 `type:"int32"`

	PrecisionDefenseCount *int32 `type:"int32"`

	PrecisionDefenseTodayCount *int32 `type:"int32"`

	UnhandledTodayCount *int32 `type:"int32"`
}

// String returns the string representation
func (s GetRaspAlarmStatisticsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s GetRaspAlarmStatisticsOutput) GoString() string {
	return s.String()
}

// SetAlarmCriticalCount sets the AlarmCriticalCount field's value.
func (s *GetRaspAlarmStatisticsOutput) SetAlarmCriticalCount(v int32) *GetRaspAlarmStatisticsOutput {
	s.AlarmCriticalCount = &v
	return s
}

// SetAlarmHighCount sets the AlarmHighCount field's value.
func (s *GetRaspAlarmStatisticsOutput) SetAlarmHighCount(v int32) *GetRaspAlarmStatisticsOutput {
	s.AlarmHighCount = &v
	return s
}

// SetAlarmLowCount sets the AlarmLowCount field's value.
func (s *GetRaspAlarmStatisticsOutput) SetAlarmLowCount(v int32) *GetRaspAlarmStatisticsOutput {
	s.AlarmLowCount = &v
	return s
}

// SetAlarmMediumCount sets the AlarmMediumCount field's value.
func (s *GetRaspAlarmStatisticsOutput) SetAlarmMediumCount(v int32) *GetRaspAlarmStatisticsOutput {
	s.AlarmMediumCount = &v
	return s
}

// SetAlarmProcessedCount sets the AlarmProcessedCount field's value.
func (s *GetRaspAlarmStatisticsOutput) SetAlarmProcessedCount(v int32) *GetRaspAlarmStatisticsOutput {
	s.AlarmProcessedCount = &v
	return s
}

// SetAlarmTotal sets the AlarmTotal field's value.
func (s *GetRaspAlarmStatisticsOutput) SetAlarmTotal(v int32) *GetRaspAlarmStatisticsOutput {
	s.AlarmTotal = &v
	return s
}

// SetAlarmWhiteCount sets the AlarmWhiteCount field's value.
func (s *GetRaspAlarmStatisticsOutput) SetAlarmWhiteCount(v int32) *GetRaspAlarmStatisticsOutput {
	s.AlarmWhiteCount = &v
	return s
}

// SetIsolateFileCount sets the IsolateFileCount field's value.
func (s *GetRaspAlarmStatisticsOutput) SetIsolateFileCount(v int32) *GetRaspAlarmStatisticsOutput {
	s.IsolateFileCount = &v
	return s
}

// SetIsolateTodayCount sets the IsolateTodayCount field's value.
func (s *GetRaspAlarmStatisticsOutput) SetIsolateTodayCount(v int32) *GetRaspAlarmStatisticsOutput {
	s.IsolateTodayCount = &v
	return s
}

// SetPrecisionDefenseCount sets the PrecisionDefenseCount field's value.
func (s *GetRaspAlarmStatisticsOutput) SetPrecisionDefenseCount(v int32) *GetRaspAlarmStatisticsOutput {
	s.PrecisionDefenseCount = &v
	return s
}

// SetPrecisionDefenseTodayCount sets the PrecisionDefenseTodayCount field's value.
func (s *GetRaspAlarmStatisticsOutput) SetPrecisionDefenseTodayCount(v int32) *GetRaspAlarmStatisticsOutput {
	s.PrecisionDefenseTodayCount = &v
	return s
}

// SetUnhandledTodayCount sets the UnhandledTodayCount field's value.
func (s *GetRaspAlarmStatisticsOutput) SetUnhandledTodayCount(v int32) *GetRaspAlarmStatisticsOutput {
	s.UnhandledTodayCount = &v
	return s
}

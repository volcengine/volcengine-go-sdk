// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package advdefence

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescWebAtkStatisticsCommon = "DescWebAtkStatistics"

// DescWebAtkStatisticsCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescWebAtkStatisticsCommon operation. The "output" return
// value will be populated with the DescWebAtkStatisticsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescWebAtkStatisticsCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescWebAtkStatisticsCommon Send returns without error.
//
// See DescWebAtkStatisticsCommon for more information on using the DescWebAtkStatisticsCommon
// API call, and error handling.
//
//    // Example sending a request using the DescWebAtkStatisticsCommonRequest method.
//    req, resp := client.DescWebAtkStatisticsCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ADVDEFENCE) DescWebAtkStatisticsCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescWebAtkStatisticsCommon,
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

// DescWebAtkStatisticsCommon API operation for ADVDEFENCE.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ADVDEFENCE's
// API operation DescWebAtkStatisticsCommon for usage and error information.
func (c *ADVDEFENCE) DescWebAtkStatisticsCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescWebAtkStatisticsCommonRequest(input)
	return out, req.Send()
}

// DescWebAtkStatisticsCommonWithContext is the same as DescWebAtkStatisticsCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescWebAtkStatisticsCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ADVDEFENCE) DescWebAtkStatisticsCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescWebAtkStatisticsCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescWebAtkStatistics = "DescWebAtkStatistics"

// DescWebAtkStatisticsRequest generates a "volcengine/request.Request" representing the
// client's request for the DescWebAtkStatistics operation. The "output" return
// value will be populated with the DescWebAtkStatisticsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescWebAtkStatisticsCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescWebAtkStatisticsCommon Send returns without error.
//
// See DescWebAtkStatistics for more information on using the DescWebAtkStatistics
// API call, and error handling.
//
//    // Example sending a request using the DescWebAtkStatisticsRequest method.
//    req, resp := client.DescWebAtkStatisticsRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ADVDEFENCE) DescWebAtkStatisticsRequest(input *DescWebAtkStatisticsInput) (req *request.Request, output *DescWebAtkStatisticsOutput) {
	op := &request.Operation{
		Name:       opDescWebAtkStatistics,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescWebAtkStatisticsInput{}
	}

	output = &DescWebAtkStatisticsOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// DescWebAtkStatistics API operation for ADVDEFENCE.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ADVDEFENCE's
// API operation DescWebAtkStatistics for usage and error information.
func (c *ADVDEFENCE) DescWebAtkStatistics(input *DescWebAtkStatisticsInput) (*DescWebAtkStatisticsOutput, error) {
	req, out := c.DescWebAtkStatisticsRequest(input)
	return out, req.Send()
}

// DescWebAtkStatisticsWithContext is the same as DescWebAtkStatistics with the addition of
// the ability to pass a context and additional request options.
//
// See DescWebAtkStatistics for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ADVDEFENCE) DescWebAtkStatisticsWithContext(ctx volcengine.Context, input *DescWebAtkStatisticsInput, opts ...request.Option) (*DescWebAtkStatisticsOutput, error) {
	req, out := c.DescWebAtkStatisticsRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type AttackFlowForDescWebAtkStatisticsOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Count *int32 `type:"int32" json:",omitempty"`

	TimeStamp *int32 `type:"int32" json:",omitempty"`
}

// String returns the string representation
func (s AttackFlowForDescWebAtkStatisticsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s AttackFlowForDescWebAtkStatisticsOutput) GoString() string {
	return s.String()
}

// SetCount sets the Count field's value.
func (s *AttackFlowForDescWebAtkStatisticsOutput) SetCount(v int32) *AttackFlowForDescWebAtkStatisticsOutput {
	s.Count = &v
	return s
}

// SetTimeStamp sets the TimeStamp field's value.
func (s *AttackFlowForDescWebAtkStatisticsOutput) SetTimeStamp(v int32) *AttackFlowForDescWebAtkStatisticsOutput {
	s.TimeStamp = &v
	return s
}

type BackSrcFlowForDescWebAtkStatisticsOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Count *int32 `type:"int32" json:",omitempty"`

	TimeStamp *int32 `type:"int32" json:",omitempty"`
}

// String returns the string representation
func (s BackSrcFlowForDescWebAtkStatisticsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s BackSrcFlowForDescWebAtkStatisticsOutput) GoString() string {
	return s.String()
}

// SetCount sets the Count field's value.
func (s *BackSrcFlowForDescWebAtkStatisticsOutput) SetCount(v int32) *BackSrcFlowForDescWebAtkStatisticsOutput {
	s.Count = &v
	return s
}

// SetTimeStamp sets the TimeStamp field's value.
func (s *BackSrcFlowForDescWebAtkStatisticsOutput) SetTimeStamp(v int32) *BackSrcFlowForDescWebAtkStatisticsOutput {
	s.TimeStamp = &v
	return s
}

type DescWebAtkStatisticsInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	BeginTime *int32 `type:"int32" json:",omitempty"`

	EndTime *int32 `type:"int32" json:",omitempty"`

	Hosts []*string `type:"list" json:",omitempty"`
}

// String returns the string representation
func (s DescWebAtkStatisticsInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescWebAtkStatisticsInput) GoString() string {
	return s.String()
}

// SetBeginTime sets the BeginTime field's value.
func (s *DescWebAtkStatisticsInput) SetBeginTime(v int32) *DescWebAtkStatisticsInput {
	s.BeginTime = &v
	return s
}

// SetEndTime sets the EndTime field's value.
func (s *DescWebAtkStatisticsInput) SetEndTime(v int32) *DescWebAtkStatisticsInput {
	s.EndTime = &v
	return s
}

// SetHosts sets the Hosts field's value.
func (s *DescWebAtkStatisticsInput) SetHosts(v []*string) *DescWebAtkStatisticsInput {
	s.Hosts = v
	return s
}

type DescWebAtkStatisticsOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	AttackFlow []*AttackFlowForDescWebAtkStatisticsOutput `type:"list" json:",omitempty"`

	BackSrcFlow []*BackSrcFlowForDescWebAtkStatisticsOutput `type:"list" json:",omitempty"`

	InQueryFlow []*InQueryFlowForDescWebAtkStatisticsOutput `type:"list" json:",omitempty"`
}

// String returns the string representation
func (s DescWebAtkStatisticsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescWebAtkStatisticsOutput) GoString() string {
	return s.String()
}

// SetAttackFlow sets the AttackFlow field's value.
func (s *DescWebAtkStatisticsOutput) SetAttackFlow(v []*AttackFlowForDescWebAtkStatisticsOutput) *DescWebAtkStatisticsOutput {
	s.AttackFlow = v
	return s
}

// SetBackSrcFlow sets the BackSrcFlow field's value.
func (s *DescWebAtkStatisticsOutput) SetBackSrcFlow(v []*BackSrcFlowForDescWebAtkStatisticsOutput) *DescWebAtkStatisticsOutput {
	s.BackSrcFlow = v
	return s
}

// SetInQueryFlow sets the InQueryFlow field's value.
func (s *DescWebAtkStatisticsOutput) SetInQueryFlow(v []*InQueryFlowForDescWebAtkStatisticsOutput) *DescWebAtkStatisticsOutput {
	s.InQueryFlow = v
	return s
}

type InQueryFlowForDescWebAtkStatisticsOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Count *int32 `type:"int32" json:",omitempty"`

	TimeStamp *int32 `type:"int32" json:",omitempty"`
}

// String returns the string representation
func (s InQueryFlowForDescWebAtkStatisticsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s InQueryFlowForDescWebAtkStatisticsOutput) GoString() string {
	return s.String()
}

// SetCount sets the Count field's value.
func (s *InQueryFlowForDescWebAtkStatisticsOutput) SetCount(v int32) *InQueryFlowForDescWebAtkStatisticsOutput {
	s.Count = &v
	return s
}

// SetTimeStamp sets the TimeStamp field's value.
func (s *InQueryFlowForDescWebAtkStatisticsOutput) SetTimeStamp(v int32) *InQueryFlowForDescWebAtkStatisticsOutput {
	s.TimeStamp = &v
	return s
}
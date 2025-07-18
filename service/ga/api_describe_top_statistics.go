// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package ga

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribeTopStatisticsCommon = "DescribeTopStatistics"

// DescribeTopStatisticsCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeTopStatisticsCommon operation. The "output" return
// value will be populated with the DescribeTopStatisticsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeTopStatisticsCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeTopStatisticsCommon Send returns without error.
//
// See DescribeTopStatisticsCommon for more information on using the DescribeTopStatisticsCommon
// API call, and error handling.
//
//    // Example sending a request using the DescribeTopStatisticsCommonRequest method.
//    req, resp := client.DescribeTopStatisticsCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *GA) DescribeTopStatisticsCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeTopStatisticsCommon,
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

// DescribeTopStatisticsCommon API operation for GA.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for GA's
// API operation DescribeTopStatisticsCommon for usage and error information.
func (c *GA) DescribeTopStatisticsCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeTopStatisticsCommonRequest(input)
	return out, req.Send()
}

// DescribeTopStatisticsCommonWithContext is the same as DescribeTopStatisticsCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeTopStatisticsCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *GA) DescribeTopStatisticsCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeTopStatisticsCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeTopStatistics = "DescribeTopStatistics"

// DescribeTopStatisticsRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeTopStatistics operation. The "output" return
// value will be populated with the DescribeTopStatisticsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeTopStatisticsCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeTopStatisticsCommon Send returns without error.
//
// See DescribeTopStatistics for more information on using the DescribeTopStatistics
// API call, and error handling.
//
//    // Example sending a request using the DescribeTopStatisticsRequest method.
//    req, resp := client.DescribeTopStatisticsRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *GA) DescribeTopStatisticsRequest(input *DescribeTopStatisticsInput) (req *request.Request, output *DescribeTopStatisticsOutput) {
	op := &request.Operation{
		Name:       opDescribeTopStatistics,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeTopStatisticsInput{}
	}

	output = &DescribeTopStatisticsOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// DescribeTopStatistics API operation for GA.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for GA's
// API operation DescribeTopStatistics for usage and error information.
func (c *GA) DescribeTopStatistics(input *DescribeTopStatisticsInput) (*DescribeTopStatisticsOutput, error) {
	req, out := c.DescribeTopStatisticsRequest(input)
	return out, req.Send()
}

// DescribeTopStatisticsWithContext is the same as DescribeTopStatistics with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeTopStatistics for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *GA) DescribeTopStatisticsWithContext(ctx volcengine.Context, input *DescribeTopStatisticsInput, opts ...request.Option) (*DescribeTopStatisticsOutput, error) {
	req, out := c.DescribeTopStatisticsRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DescribeTopStatisticsInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	// EndTime is a required field
	EndTime *string `type:"string" json:",omitempty" required:"true"`

	InputId []*string `type:"list" json:"inputId,omitempty"`

	// InputIdType is a required field
	InputIdType *string `type:"string" json:",omitempty" required:"true"`

	SortMetric *string `type:"string" json:",omitempty"`

	// StartTime is a required field
	StartTime *string `type:"string" json:",omitempty" required:"true"`
}

// String returns the string representation
func (s DescribeTopStatisticsInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeTopStatisticsInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeTopStatisticsInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DescribeTopStatisticsInput"}
	if s.EndTime == nil {
		invalidParams.Add(request.NewErrParamRequired("EndTime"))
	}
	if s.InputIdType == nil {
		invalidParams.Add(request.NewErrParamRequired("InputIdType"))
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
func (s *DescribeTopStatisticsInput) SetEndTime(v string) *DescribeTopStatisticsInput {
	s.EndTime = &v
	return s
}

// SetInputId sets the InputId field's value.
func (s *DescribeTopStatisticsInput) SetInputId(v []*string) *DescribeTopStatisticsInput {
	s.InputId = v
	return s
}

// SetInputIdType sets the InputIdType field's value.
func (s *DescribeTopStatisticsInput) SetInputIdType(v string) *DescribeTopStatisticsInput {
	s.InputIdType = &v
	return s
}

// SetSortMetric sets the SortMetric field's value.
func (s *DescribeTopStatisticsInput) SetSortMetric(v string) *DescribeTopStatisticsInput {
	s.SortMetric = &v
	return s
}

// SetStartTime sets the StartTime field's value.
func (s *DescribeTopStatisticsInput) SetStartTime(v string) *DescribeTopStatisticsInput {
	s.StartTime = &v
	return s
}

type DescribeTopStatisticsOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	SortMetric *string `type:"string" json:",omitempty"`

	TopStatistics []*TopStatisticForDescribeTopStatisticsOutput `type:"list" json:",omitempty"`
}

// String returns the string representation
func (s DescribeTopStatisticsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeTopStatisticsOutput) GoString() string {
	return s.String()
}

// SetSortMetric sets the SortMetric field's value.
func (s *DescribeTopStatisticsOutput) SetSortMetric(v string) *DescribeTopStatisticsOutput {
	s.SortMetric = &v
	return s
}

// SetTopStatistics sets the TopStatistics field's value.
func (s *DescribeTopStatisticsOutput) SetTopStatistics(v []*TopStatisticForDescribeTopStatisticsOutput) *DescribeTopStatisticsOutput {
	s.TopStatistics = v
	return s
}

type TopStatisticForDescribeTopStatisticsOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	ID *string `type:"string" json:",omitempty"`

	MaxBandwidth *float64 `type:"float" json:",omitempty"`

	MaxBandwidth95 *float64 `type:"float" json:",omitempty"`

	MaxConnectionNum *float64 `type:"float" json:",omitempty"`

	Name *string `type:"string" json:",omitempty"`

	Rank *int32 `type:"int32" json:",omitempty"`

	TotalTraffic *float64 `type:"float" json:",omitempty"`
}

// String returns the string representation
func (s TopStatisticForDescribeTopStatisticsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s TopStatisticForDescribeTopStatisticsOutput) GoString() string {
	return s.String()
}

// SetID sets the ID field's value.
func (s *TopStatisticForDescribeTopStatisticsOutput) SetID(v string) *TopStatisticForDescribeTopStatisticsOutput {
	s.ID = &v
	return s
}

// SetMaxBandwidth sets the MaxBandwidth field's value.
func (s *TopStatisticForDescribeTopStatisticsOutput) SetMaxBandwidth(v float64) *TopStatisticForDescribeTopStatisticsOutput {
	s.MaxBandwidth = &v
	return s
}

// SetMaxBandwidth95 sets the MaxBandwidth95 field's value.
func (s *TopStatisticForDescribeTopStatisticsOutput) SetMaxBandwidth95(v float64) *TopStatisticForDescribeTopStatisticsOutput {
	s.MaxBandwidth95 = &v
	return s
}

// SetMaxConnectionNum sets the MaxConnectionNum field's value.
func (s *TopStatisticForDescribeTopStatisticsOutput) SetMaxConnectionNum(v float64) *TopStatisticForDescribeTopStatisticsOutput {
	s.MaxConnectionNum = &v
	return s
}

// SetName sets the Name field's value.
func (s *TopStatisticForDescribeTopStatisticsOutput) SetName(v string) *TopStatisticForDescribeTopStatisticsOutput {
	s.Name = &v
	return s
}

// SetRank sets the Rank field's value.
func (s *TopStatisticForDescribeTopStatisticsOutput) SetRank(v int32) *TopStatisticForDescribeTopStatisticsOutput {
	s.Rank = &v
	return s
}

// SetTotalTraffic sets the TotalTraffic field's value.
func (s *TopStatisticForDescribeTopStatisticsOutput) SetTotalTraffic(v float64) *TopStatisticForDescribeTopStatisticsOutput {
	s.TotalTraffic = &v
	return s
}

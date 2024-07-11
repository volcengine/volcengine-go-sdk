// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package cdn

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribeDistrictRankingCommon = "DescribeDistrictRanking"

// DescribeDistrictRankingCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeDistrictRankingCommon operation. The "output" return
// value will be populated with the DescribeDistrictRankingCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeDistrictRankingCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeDistrictRankingCommon Send returns without error.
//
// See DescribeDistrictRankingCommon for more information on using the DescribeDistrictRankingCommon
// API call, and error handling.
//
//    // Example sending a request using the DescribeDistrictRankingCommonRequest method.
//    req, resp := client.DescribeDistrictRankingCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CDN) DescribeDistrictRankingCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeDistrictRankingCommon,
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

// DescribeDistrictRankingCommon API operation for CDN.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CDN's
// API operation DescribeDistrictRankingCommon for usage and error information.
func (c *CDN) DescribeDistrictRankingCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeDistrictRankingCommonRequest(input)
	return out, req.Send()
}

// DescribeDistrictRankingCommonWithContext is the same as DescribeDistrictRankingCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeDistrictRankingCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CDN) DescribeDistrictRankingCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeDistrictRankingCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeDistrictRanking = "DescribeDistrictRanking"

// DescribeDistrictRankingRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeDistrictRanking operation. The "output" return
// value will be populated with the DescribeDistrictRankingCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeDistrictRankingCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeDistrictRankingCommon Send returns without error.
//
// See DescribeDistrictRanking for more information on using the DescribeDistrictRanking
// API call, and error handling.
//
//    // Example sending a request using the DescribeDistrictRankingRequest method.
//    req, resp := client.DescribeDistrictRankingRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CDN) DescribeDistrictRankingRequest(input *DescribeDistrictRankingInput) (req *request.Request, output *DescribeDistrictRankingOutput) {
	op := &request.Operation{
		Name:       opDescribeDistrictRanking,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeDistrictRankingInput{}
	}

	output = &DescribeDistrictRankingOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// DescribeDistrictRanking API operation for CDN.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CDN's
// API operation DescribeDistrictRanking for usage and error information.
func (c *CDN) DescribeDistrictRanking(input *DescribeDistrictRankingInput) (*DescribeDistrictRankingOutput, error) {
	req, out := c.DescribeDistrictRankingRequest(input)
	return out, req.Send()
}

// DescribeDistrictRankingWithContext is the same as DescribeDistrictRanking with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeDistrictRanking for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CDN) DescribeDistrictRankingWithContext(ctx volcengine.Context, input *DescribeDistrictRankingInput, opts ...request.Option) (*DescribeDistrictRankingOutput, error) {
	req, out := c.DescribeDistrictRankingRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DescribeDistrictRankingInput struct {
	_ struct{} `type:"structure"`

	Domain *string `type:"string"`

	// EndTime is a required field
	EndTime *int64 `type:"int64" required:"true"`

	Interval *string `type:"string"`

	// Item is a required field
	Item *string `type:"string" required:"true"`

	// Metric is a required field
	Metric *string `type:"string" required:"true"`

	Project *string `type:"string"`

	// StartTime is a required field
	StartTime *int64 `type:"int64" required:"true"`
}

// String returns the string representation
func (s DescribeDistrictRankingInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeDistrictRankingInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeDistrictRankingInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DescribeDistrictRankingInput"}
	if s.EndTime == nil {
		invalidParams.Add(request.NewErrParamRequired("EndTime"))
	}
	if s.Item == nil {
		invalidParams.Add(request.NewErrParamRequired("Item"))
	}
	if s.Metric == nil {
		invalidParams.Add(request.NewErrParamRequired("Metric"))
	}
	if s.StartTime == nil {
		invalidParams.Add(request.NewErrParamRequired("StartTime"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetDomain sets the Domain field's value.
func (s *DescribeDistrictRankingInput) SetDomain(v string) *DescribeDistrictRankingInput {
	s.Domain = &v
	return s
}

// SetEndTime sets the EndTime field's value.
func (s *DescribeDistrictRankingInput) SetEndTime(v int64) *DescribeDistrictRankingInput {
	s.EndTime = &v
	return s
}

// SetInterval sets the Interval field's value.
func (s *DescribeDistrictRankingInput) SetInterval(v string) *DescribeDistrictRankingInput {
	s.Interval = &v
	return s
}

// SetItem sets the Item field's value.
func (s *DescribeDistrictRankingInput) SetItem(v string) *DescribeDistrictRankingInput {
	s.Item = &v
	return s
}

// SetMetric sets the Metric field's value.
func (s *DescribeDistrictRankingInput) SetMetric(v string) *DescribeDistrictRankingInput {
	s.Metric = &v
	return s
}

// SetProject sets the Project field's value.
func (s *DescribeDistrictRankingInput) SetProject(v string) *DescribeDistrictRankingInput {
	s.Project = &v
	return s
}

// SetStartTime sets the StartTime field's value.
func (s *DescribeDistrictRankingInput) SetStartTime(v int64) *DescribeDistrictRankingInput {
	s.StartTime = &v
	return s
}

type DescribeDistrictRankingOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	Item *string `type:"string"`

	TopDataDetails []*TopDataDetailForDescribeDistrictRankingOutput `type:"list"`
}

// String returns the string representation
func (s DescribeDistrictRankingOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeDistrictRankingOutput) GoString() string {
	return s.String()
}

// SetItem sets the Item field's value.
func (s *DescribeDistrictRankingOutput) SetItem(v string) *DescribeDistrictRankingOutput {
	s.Item = &v
	return s
}

// SetTopDataDetails sets the TopDataDetails field's value.
func (s *DescribeDistrictRankingOutput) SetTopDataDetails(v []*TopDataDetailForDescribeDistrictRankingOutput) *DescribeDistrictRankingOutput {
	s.TopDataDetails = v
	return s
}

type TopDataDetailForDescribeDistrictRankingOutput struct {
	_ struct{} `type:"structure"`

	Metric *string `type:"string"`

	ValueDetails []*ValueDetailForDescribeDistrictRankingOutput `type:"list"`
}

// String returns the string representation
func (s TopDataDetailForDescribeDistrictRankingOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s TopDataDetailForDescribeDistrictRankingOutput) GoString() string {
	return s.String()
}

// SetMetric sets the Metric field's value.
func (s *TopDataDetailForDescribeDistrictRankingOutput) SetMetric(v string) *TopDataDetailForDescribeDistrictRankingOutput {
	s.Metric = &v
	return s
}

// SetValueDetails sets the ValueDetails field's value.
func (s *TopDataDetailForDescribeDistrictRankingOutput) SetValueDetails(v []*ValueDetailForDescribeDistrictRankingOutput) *TopDataDetailForDescribeDistrictRankingOutput {
	s.ValueDetails = v
	return s
}

type ValueDetailForDescribeDistrictRankingOutput struct {
	_ struct{} `type:"structure"`

	ItemKey *string `type:"string"`

	Ratio *float64 `type:"double"`

	Timestamp *int64 `type:"int64"`

	Value *float64 `type:"double"`
}

// String returns the string representation
func (s ValueDetailForDescribeDistrictRankingOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ValueDetailForDescribeDistrictRankingOutput) GoString() string {
	return s.String()
}

// SetItemKey sets the ItemKey field's value.
func (s *ValueDetailForDescribeDistrictRankingOutput) SetItemKey(v string) *ValueDetailForDescribeDistrictRankingOutput {
	s.ItemKey = &v
	return s
}

// SetRatio sets the Ratio field's value.
func (s *ValueDetailForDescribeDistrictRankingOutput) SetRatio(v float64) *ValueDetailForDescribeDistrictRankingOutput {
	s.Ratio = &v
	return s
}

// SetTimestamp sets the Timestamp field's value.
func (s *ValueDetailForDescribeDistrictRankingOutput) SetTimestamp(v int64) *ValueDetailForDescribeDistrictRankingOutput {
	s.Timestamp = &v
	return s
}

// SetValue sets the Value field's value.
func (s *ValueDetailForDescribeDistrictRankingOutput) SetValue(v float64) *ValueDetailForDescribeDistrictRankingOutput {
	s.Value = &v
	return s
}

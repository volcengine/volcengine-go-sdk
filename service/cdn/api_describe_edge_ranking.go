// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package cdn

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribeEdgeRankingCommon = "DescribeEdgeRanking"

// DescribeEdgeRankingCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeEdgeRankingCommon operation. The "output" return
// value will be populated with the DescribeEdgeRankingCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeEdgeRankingCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeEdgeRankingCommon Send returns without error.
//
// See DescribeEdgeRankingCommon for more information on using the DescribeEdgeRankingCommon
// API call, and error handling.
//
//    // Example sending a request using the DescribeEdgeRankingCommonRequest method.
//    req, resp := client.DescribeEdgeRankingCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CDN) DescribeEdgeRankingCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeEdgeRankingCommon,
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

// DescribeEdgeRankingCommon API operation for CDN.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CDN's
// API operation DescribeEdgeRankingCommon for usage and error information.
func (c *CDN) DescribeEdgeRankingCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeEdgeRankingCommonRequest(input)
	return out, req.Send()
}

// DescribeEdgeRankingCommonWithContext is the same as DescribeEdgeRankingCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeEdgeRankingCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CDN) DescribeEdgeRankingCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeEdgeRankingCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeEdgeRanking = "DescribeEdgeRanking"

// DescribeEdgeRankingRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeEdgeRanking operation. The "output" return
// value will be populated with the DescribeEdgeRankingCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeEdgeRankingCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeEdgeRankingCommon Send returns without error.
//
// See DescribeEdgeRanking for more information on using the DescribeEdgeRanking
// API call, and error handling.
//
//    // Example sending a request using the DescribeEdgeRankingRequest method.
//    req, resp := client.DescribeEdgeRankingRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CDN) DescribeEdgeRankingRequest(input *DescribeEdgeRankingInput) (req *request.Request, output *DescribeEdgeRankingOutput) {
	op := &request.Operation{
		Name:       opDescribeEdgeRanking,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeEdgeRankingInput{}
	}

	output = &DescribeEdgeRankingOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// DescribeEdgeRanking API operation for CDN.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CDN's
// API operation DescribeEdgeRanking for usage and error information.
func (c *CDN) DescribeEdgeRanking(input *DescribeEdgeRankingInput) (*DescribeEdgeRankingOutput, error) {
	req, out := c.DescribeEdgeRankingRequest(input)
	return out, req.Send()
}

// DescribeEdgeRankingWithContext is the same as DescribeEdgeRanking with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeEdgeRanking for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CDN) DescribeEdgeRankingWithContext(ctx volcengine.Context, input *DescribeEdgeRankingInput, opts ...request.Option) (*DescribeEdgeRankingOutput, error) {
	req, out := c.DescribeEdgeRankingRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DescribeEdgeRankingInput struct {
	_ struct{} `type:"structure"`

	BillingRegion *string `type:"string"`

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
func (s DescribeEdgeRankingInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeEdgeRankingInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeEdgeRankingInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DescribeEdgeRankingInput"}
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

// SetBillingRegion sets the BillingRegion field's value.
func (s *DescribeEdgeRankingInput) SetBillingRegion(v string) *DescribeEdgeRankingInput {
	s.BillingRegion = &v
	return s
}

// SetDomain sets the Domain field's value.
func (s *DescribeEdgeRankingInput) SetDomain(v string) *DescribeEdgeRankingInput {
	s.Domain = &v
	return s
}

// SetEndTime sets the EndTime field's value.
func (s *DescribeEdgeRankingInput) SetEndTime(v int64) *DescribeEdgeRankingInput {
	s.EndTime = &v
	return s
}

// SetInterval sets the Interval field's value.
func (s *DescribeEdgeRankingInput) SetInterval(v string) *DescribeEdgeRankingInput {
	s.Interval = &v
	return s
}

// SetItem sets the Item field's value.
func (s *DescribeEdgeRankingInput) SetItem(v string) *DescribeEdgeRankingInput {
	s.Item = &v
	return s
}

// SetMetric sets the Metric field's value.
func (s *DescribeEdgeRankingInput) SetMetric(v string) *DescribeEdgeRankingInput {
	s.Metric = &v
	return s
}

// SetProject sets the Project field's value.
func (s *DescribeEdgeRankingInput) SetProject(v string) *DescribeEdgeRankingInput {
	s.Project = &v
	return s
}

// SetStartTime sets the StartTime field's value.
func (s *DescribeEdgeRankingInput) SetStartTime(v int64) *DescribeEdgeRankingInput {
	s.StartTime = &v
	return s
}

type DescribeEdgeRankingOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	Item *string `type:"string"`

	TopDataDetails []*TopDataDetailForDescribeEdgeRankingOutput `type:"list"`
}

// String returns the string representation
func (s DescribeEdgeRankingOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeEdgeRankingOutput) GoString() string {
	return s.String()
}

// SetItem sets the Item field's value.
func (s *DescribeEdgeRankingOutput) SetItem(v string) *DescribeEdgeRankingOutput {
	s.Item = &v
	return s
}

// SetTopDataDetails sets the TopDataDetails field's value.
func (s *DescribeEdgeRankingOutput) SetTopDataDetails(v []*TopDataDetailForDescribeEdgeRankingOutput) *DescribeEdgeRankingOutput {
	s.TopDataDetails = v
	return s
}

type TopDataDetailForDescribeEdgeRankingOutput struct {
	_ struct{} `type:"structure"`

	Metric *string `type:"string"`

	ValueDetails []*ValueDetailForDescribeEdgeRankingOutput `type:"list"`
}

// String returns the string representation
func (s TopDataDetailForDescribeEdgeRankingOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s TopDataDetailForDescribeEdgeRankingOutput) GoString() string {
	return s.String()
}

// SetMetric sets the Metric field's value.
func (s *TopDataDetailForDescribeEdgeRankingOutput) SetMetric(v string) *TopDataDetailForDescribeEdgeRankingOutput {
	s.Metric = &v
	return s
}

// SetValueDetails sets the ValueDetails field's value.
func (s *TopDataDetailForDescribeEdgeRankingOutput) SetValueDetails(v []*ValueDetailForDescribeEdgeRankingOutput) *TopDataDetailForDescribeEdgeRankingOutput {
	s.ValueDetails = v
	return s
}

type ValueDetailForDescribeEdgeRankingOutput struct {
	_ struct{} `type:"structure"`

	ItemKey *string `type:"string"`

	Ratio *float64 `type:"double"`

	Timestamp *int64 `type:"int64"`

	Value *float64 `type:"double"`
}

// String returns the string representation
func (s ValueDetailForDescribeEdgeRankingOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ValueDetailForDescribeEdgeRankingOutput) GoString() string {
	return s.String()
}

// SetItemKey sets the ItemKey field's value.
func (s *ValueDetailForDescribeEdgeRankingOutput) SetItemKey(v string) *ValueDetailForDescribeEdgeRankingOutput {
	s.ItemKey = &v
	return s
}

// SetRatio sets the Ratio field's value.
func (s *ValueDetailForDescribeEdgeRankingOutput) SetRatio(v float64) *ValueDetailForDescribeEdgeRankingOutput {
	s.Ratio = &v
	return s
}

// SetTimestamp sets the Timestamp field's value.
func (s *ValueDetailForDescribeEdgeRankingOutput) SetTimestamp(v int64) *ValueDetailForDescribeEdgeRankingOutput {
	s.Timestamp = &v
	return s
}

// SetValue sets the Value field's value.
func (s *ValueDetailForDescribeEdgeRankingOutput) SetValue(v float64) *ValueDetailForDescribeEdgeRankingOutput {
	s.Value = &v
	return s
}

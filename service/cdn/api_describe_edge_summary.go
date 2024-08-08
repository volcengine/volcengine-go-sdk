// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package cdn

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribeEdgeSummaryCommon = "DescribeEdgeSummary"

// DescribeEdgeSummaryCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeEdgeSummaryCommon operation. The "output" return
// value will be populated with the DescribeEdgeSummaryCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeEdgeSummaryCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeEdgeSummaryCommon Send returns without error.
//
// See DescribeEdgeSummaryCommon for more information on using the DescribeEdgeSummaryCommon
// API call, and error handling.
//
//    // Example sending a request using the DescribeEdgeSummaryCommonRequest method.
//    req, resp := client.DescribeEdgeSummaryCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CDN) DescribeEdgeSummaryCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeEdgeSummaryCommon,
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

// DescribeEdgeSummaryCommon API operation for CDN.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CDN's
// API operation DescribeEdgeSummaryCommon for usage and error information.
func (c *CDN) DescribeEdgeSummaryCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeEdgeSummaryCommonRequest(input)
	return out, req.Send()
}

// DescribeEdgeSummaryCommonWithContext is the same as DescribeEdgeSummaryCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeEdgeSummaryCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CDN) DescribeEdgeSummaryCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeEdgeSummaryCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeEdgeSummary = "DescribeEdgeSummary"

// DescribeEdgeSummaryRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeEdgeSummary operation. The "output" return
// value will be populated with the DescribeEdgeSummaryCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeEdgeSummaryCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeEdgeSummaryCommon Send returns without error.
//
// See DescribeEdgeSummary for more information on using the DescribeEdgeSummary
// API call, and error handling.
//
//    // Example sending a request using the DescribeEdgeSummaryRequest method.
//    req, resp := client.DescribeEdgeSummaryRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CDN) DescribeEdgeSummaryRequest(input *DescribeEdgeSummaryInput) (req *request.Request, output *DescribeEdgeSummaryOutput) {
	op := &request.Operation{
		Name:       opDescribeEdgeSummary,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeEdgeSummaryInput{}
	}

	output = &DescribeEdgeSummaryOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// DescribeEdgeSummary API operation for CDN.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CDN's
// API operation DescribeEdgeSummary for usage and error information.
func (c *CDN) DescribeEdgeSummary(input *DescribeEdgeSummaryInput) (*DescribeEdgeSummaryOutput, error) {
	req, out := c.DescribeEdgeSummaryRequest(input)
	return out, req.Send()
}

// DescribeEdgeSummaryWithContext is the same as DescribeEdgeSummary with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeEdgeSummary for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CDN) DescribeEdgeSummaryWithContext(ctx volcengine.Context, input *DescribeEdgeSummaryInput, opts ...request.Option) (*DescribeEdgeSummaryOutput, error) {
	req, out := c.DescribeEdgeSummaryRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DescribeEdgeSummaryInput struct {
	_ struct{} `type:"structure"`

	BillingRegion *string `type:"string"`

	Domain *string `type:"string"`

	// EndTime is a required field
	EndTime *int64 `type:"int64" required:"true"`

	Interval *string `type:"string"`

	// Metric is a required field
	Metric *string `type:"string" required:"true"`

	Project *string `type:"string"`

	// StartTime is a required field
	StartTime *int64 `type:"int64" required:"true"`
}

// String returns the string representation
func (s DescribeEdgeSummaryInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeEdgeSummaryInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeEdgeSummaryInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DescribeEdgeSummaryInput"}
	if s.EndTime == nil {
		invalidParams.Add(request.NewErrParamRequired("EndTime"))
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
func (s *DescribeEdgeSummaryInput) SetBillingRegion(v string) *DescribeEdgeSummaryInput {
	s.BillingRegion = &v
	return s
}

// SetDomain sets the Domain field's value.
func (s *DescribeEdgeSummaryInput) SetDomain(v string) *DescribeEdgeSummaryInput {
	s.Domain = &v
	return s
}

// SetEndTime sets the EndTime field's value.
func (s *DescribeEdgeSummaryInput) SetEndTime(v int64) *DescribeEdgeSummaryInput {
	s.EndTime = &v
	return s
}

// SetInterval sets the Interval field's value.
func (s *DescribeEdgeSummaryInput) SetInterval(v string) *DescribeEdgeSummaryInput {
	s.Interval = &v
	return s
}

// SetMetric sets the Metric field's value.
func (s *DescribeEdgeSummaryInput) SetMetric(v string) *DescribeEdgeSummaryInput {
	s.Metric = &v
	return s
}

// SetProject sets the Project field's value.
func (s *DescribeEdgeSummaryInput) SetProject(v string) *DescribeEdgeSummaryInput {
	s.Project = &v
	return s
}

// SetStartTime sets the StartTime field's value.
func (s *DescribeEdgeSummaryInput) SetStartTime(v int64) *DescribeEdgeSummaryInput {
	s.StartTime = &v
	return s
}

type DescribeEdgeSummaryOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	MetricDataList []*MetricDataListForDescribeEdgeSummaryOutput `type:"list"`
}

// String returns the string representation
func (s DescribeEdgeSummaryOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeEdgeSummaryOutput) GoString() string {
	return s.String()
}

// SetMetricDataList sets the MetricDataList field's value.
func (s *DescribeEdgeSummaryOutput) SetMetricDataList(v []*MetricDataListForDescribeEdgeSummaryOutput) *DescribeEdgeSummaryOutput {
	s.MetricDataList = v
	return s
}

type MetricDataListForDescribeEdgeSummaryOutput struct {
	_ struct{} `type:"structure"`

	Metric *string `type:"string"`

	Value *float64 `type:"double"`
}

// String returns the string representation
func (s MetricDataListForDescribeEdgeSummaryOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s MetricDataListForDescribeEdgeSummaryOutput) GoString() string {
	return s.String()
}

// SetMetric sets the Metric field's value.
func (s *MetricDataListForDescribeEdgeSummaryOutput) SetMetric(v string) *MetricDataListForDescribeEdgeSummaryOutput {
	s.Metric = &v
	return s
}

// SetValue sets the Value field's value.
func (s *MetricDataListForDescribeEdgeSummaryOutput) SetValue(v float64) *MetricDataListForDescribeEdgeSummaryOutput {
	s.Value = &v
	return s
}
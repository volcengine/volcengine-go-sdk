// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package cdn

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribeEdgeStatusCodeRankingCommon = "DescribeEdgeStatusCodeRanking"

// DescribeEdgeStatusCodeRankingCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeEdgeStatusCodeRankingCommon operation. The "output" return
// value will be populated with the DescribeEdgeStatusCodeRankingCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeEdgeStatusCodeRankingCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeEdgeStatusCodeRankingCommon Send returns without error.
//
// See DescribeEdgeStatusCodeRankingCommon for more information on using the DescribeEdgeStatusCodeRankingCommon
// API call, and error handling.
//
//    // Example sending a request using the DescribeEdgeStatusCodeRankingCommonRequest method.
//    req, resp := client.DescribeEdgeStatusCodeRankingCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CDN) DescribeEdgeStatusCodeRankingCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeEdgeStatusCodeRankingCommon,
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

// DescribeEdgeStatusCodeRankingCommon API operation for CDN.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CDN's
// API operation DescribeEdgeStatusCodeRankingCommon for usage and error information.
func (c *CDN) DescribeEdgeStatusCodeRankingCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeEdgeStatusCodeRankingCommonRequest(input)
	return out, req.Send()
}

// DescribeEdgeStatusCodeRankingCommonWithContext is the same as DescribeEdgeStatusCodeRankingCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeEdgeStatusCodeRankingCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CDN) DescribeEdgeStatusCodeRankingCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeEdgeStatusCodeRankingCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeEdgeStatusCodeRanking = "DescribeEdgeStatusCodeRanking"

// DescribeEdgeStatusCodeRankingRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeEdgeStatusCodeRanking operation. The "output" return
// value will be populated with the DescribeEdgeStatusCodeRankingCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeEdgeStatusCodeRankingCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeEdgeStatusCodeRankingCommon Send returns without error.
//
// See DescribeEdgeStatusCodeRanking for more information on using the DescribeEdgeStatusCodeRanking
// API call, and error handling.
//
//    // Example sending a request using the DescribeEdgeStatusCodeRankingRequest method.
//    req, resp := client.DescribeEdgeStatusCodeRankingRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CDN) DescribeEdgeStatusCodeRankingRequest(input *DescribeEdgeStatusCodeRankingInput) (req *request.Request, output *DescribeEdgeStatusCodeRankingOutput) {
	op := &request.Operation{
		Name:       opDescribeEdgeStatusCodeRanking,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeEdgeStatusCodeRankingInput{}
	}

	output = &DescribeEdgeStatusCodeRankingOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// DescribeEdgeStatusCodeRanking API operation for CDN.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CDN's
// API operation DescribeEdgeStatusCodeRanking for usage and error information.
func (c *CDN) DescribeEdgeStatusCodeRanking(input *DescribeEdgeStatusCodeRankingInput) (*DescribeEdgeStatusCodeRankingOutput, error) {
	req, out := c.DescribeEdgeStatusCodeRankingRequest(input)
	return out, req.Send()
}

// DescribeEdgeStatusCodeRankingWithContext is the same as DescribeEdgeStatusCodeRanking with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeEdgeStatusCodeRanking for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CDN) DescribeEdgeStatusCodeRankingWithContext(ctx volcengine.Context, input *DescribeEdgeStatusCodeRankingInput, opts ...request.Option) (*DescribeEdgeStatusCodeRankingOutput, error) {
	req, out := c.DescribeEdgeStatusCodeRankingRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DescribeEdgeStatusCodeRankingInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Domain *string `type:"string" json:",omitempty"`

	// EndTime is a required field
	EndTime *int64 `type:"int64" json:",omitempty" required:"true"`

	Interval *string `type:"string" json:",omitempty"`

	// Item is a required field
	Item *string `type:"string" json:",omitempty" required:"true"`

	// Metric is a required field
	Metric *string `type:"string" json:",omitempty" required:"true"`

	Project *string `type:"string" json:",omitempty"`

	// StartTime is a required field
	StartTime *int64 `type:"int64" json:",omitempty" required:"true"`
}

// String returns the string representation
func (s DescribeEdgeStatusCodeRankingInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeEdgeStatusCodeRankingInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeEdgeStatusCodeRankingInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DescribeEdgeStatusCodeRankingInput"}
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
func (s *DescribeEdgeStatusCodeRankingInput) SetDomain(v string) *DescribeEdgeStatusCodeRankingInput {
	s.Domain = &v
	return s
}

// SetEndTime sets the EndTime field's value.
func (s *DescribeEdgeStatusCodeRankingInput) SetEndTime(v int64) *DescribeEdgeStatusCodeRankingInput {
	s.EndTime = &v
	return s
}

// SetInterval sets the Interval field's value.
func (s *DescribeEdgeStatusCodeRankingInput) SetInterval(v string) *DescribeEdgeStatusCodeRankingInput {
	s.Interval = &v
	return s
}

// SetItem sets the Item field's value.
func (s *DescribeEdgeStatusCodeRankingInput) SetItem(v string) *DescribeEdgeStatusCodeRankingInput {
	s.Item = &v
	return s
}

// SetMetric sets the Metric field's value.
func (s *DescribeEdgeStatusCodeRankingInput) SetMetric(v string) *DescribeEdgeStatusCodeRankingInput {
	s.Metric = &v
	return s
}

// SetProject sets the Project field's value.
func (s *DescribeEdgeStatusCodeRankingInput) SetProject(v string) *DescribeEdgeStatusCodeRankingInput {
	s.Project = &v
	return s
}

// SetStartTime sets the StartTime field's value.
func (s *DescribeEdgeStatusCodeRankingInput) SetStartTime(v int64) *DescribeEdgeStatusCodeRankingInput {
	s.StartTime = &v
	return s
}

type DescribeEdgeStatusCodeRankingOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	Item *string `type:"string" json:",omitempty"`

	Metric *string `type:"string" json:",omitempty"`

	TopDataDetails []*TopDataDetailForDescribeEdgeStatusCodeRankingOutput `type:"list" json:",omitempty"`
}

// String returns the string representation
func (s DescribeEdgeStatusCodeRankingOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeEdgeStatusCodeRankingOutput) GoString() string {
	return s.String()
}

// SetItem sets the Item field's value.
func (s *DescribeEdgeStatusCodeRankingOutput) SetItem(v string) *DescribeEdgeStatusCodeRankingOutput {
	s.Item = &v
	return s
}

// SetMetric sets the Metric field's value.
func (s *DescribeEdgeStatusCodeRankingOutput) SetMetric(v string) *DescribeEdgeStatusCodeRankingOutput {
	s.Metric = &v
	return s
}

// SetTopDataDetails sets the TopDataDetails field's value.
func (s *DescribeEdgeStatusCodeRankingOutput) SetTopDataDetails(v []*TopDataDetailForDescribeEdgeStatusCodeRankingOutput) *DescribeEdgeStatusCodeRankingOutput {
	s.TopDataDetails = v
	return s
}

type TopDataDetailForDescribeEdgeStatusCodeRankingOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	ItemKey *string `type:"string" json:",omitempty"`

	Status2xx *float64 `type:"double" json:",omitempty"`

	Status2xxRatio *float64 `type:"double" json:",omitempty"`

	Status3xx *float64 `type:"double" json:",omitempty"`

	Status3xxRatio *float64 `type:"double" json:",omitempty"`

	Status4xx *float64 `type:"double" json:",omitempty"`

	Status4xxRatio *float64 `type:"double" json:",omitempty"`

	Status5xx *float64 `type:"double" json:",omitempty"`

	Status5xxRatio *float64 `type:"double" json:",omitempty"`
}

// String returns the string representation
func (s TopDataDetailForDescribeEdgeStatusCodeRankingOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s TopDataDetailForDescribeEdgeStatusCodeRankingOutput) GoString() string {
	return s.String()
}

// SetItemKey sets the ItemKey field's value.
func (s *TopDataDetailForDescribeEdgeStatusCodeRankingOutput) SetItemKey(v string) *TopDataDetailForDescribeEdgeStatusCodeRankingOutput {
	s.ItemKey = &v
	return s
}

// SetStatus2xx sets the Status2xx field's value.
func (s *TopDataDetailForDescribeEdgeStatusCodeRankingOutput) SetStatus2xx(v float64) *TopDataDetailForDescribeEdgeStatusCodeRankingOutput {
	s.Status2xx = &v
	return s
}

// SetStatus2xxRatio sets the Status2xxRatio field's value.
func (s *TopDataDetailForDescribeEdgeStatusCodeRankingOutput) SetStatus2xxRatio(v float64) *TopDataDetailForDescribeEdgeStatusCodeRankingOutput {
	s.Status2xxRatio = &v
	return s
}

// SetStatus3xx sets the Status3xx field's value.
func (s *TopDataDetailForDescribeEdgeStatusCodeRankingOutput) SetStatus3xx(v float64) *TopDataDetailForDescribeEdgeStatusCodeRankingOutput {
	s.Status3xx = &v
	return s
}

// SetStatus3xxRatio sets the Status3xxRatio field's value.
func (s *TopDataDetailForDescribeEdgeStatusCodeRankingOutput) SetStatus3xxRatio(v float64) *TopDataDetailForDescribeEdgeStatusCodeRankingOutput {
	s.Status3xxRatio = &v
	return s
}

// SetStatus4xx sets the Status4xx field's value.
func (s *TopDataDetailForDescribeEdgeStatusCodeRankingOutput) SetStatus4xx(v float64) *TopDataDetailForDescribeEdgeStatusCodeRankingOutput {
	s.Status4xx = &v
	return s
}

// SetStatus4xxRatio sets the Status4xxRatio field's value.
func (s *TopDataDetailForDescribeEdgeStatusCodeRankingOutput) SetStatus4xxRatio(v float64) *TopDataDetailForDescribeEdgeStatusCodeRankingOutput {
	s.Status4xxRatio = &v
	return s
}

// SetStatus5xx sets the Status5xx field's value.
func (s *TopDataDetailForDescribeEdgeStatusCodeRankingOutput) SetStatus5xx(v float64) *TopDataDetailForDescribeEdgeStatusCodeRankingOutput {
	s.Status5xx = &v
	return s
}

// SetStatus5xxRatio sets the Status5xxRatio field's value.
func (s *TopDataDetailForDescribeEdgeStatusCodeRankingOutput) SetStatus5xxRatio(v float64) *TopDataDetailForDescribeEdgeStatusCodeRankingOutput {
	s.Status5xxRatio = &v
	return s
}

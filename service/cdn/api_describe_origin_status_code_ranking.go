// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package cdn

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribeOriginStatusCodeRankingCommon = "DescribeOriginStatusCodeRanking"

// DescribeOriginStatusCodeRankingCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeOriginStatusCodeRankingCommon operation. The "output" return
// value will be populated with the DescribeOriginStatusCodeRankingCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeOriginStatusCodeRankingCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeOriginStatusCodeRankingCommon Send returns without error.
//
// See DescribeOriginStatusCodeRankingCommon for more information on using the DescribeOriginStatusCodeRankingCommon
// API call, and error handling.
//
//    // Example sending a request using the DescribeOriginStatusCodeRankingCommonRequest method.
//    req, resp := client.DescribeOriginStatusCodeRankingCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CDN) DescribeOriginStatusCodeRankingCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeOriginStatusCodeRankingCommon,
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

// DescribeOriginStatusCodeRankingCommon API operation for CDN.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CDN's
// API operation DescribeOriginStatusCodeRankingCommon for usage and error information.
func (c *CDN) DescribeOriginStatusCodeRankingCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeOriginStatusCodeRankingCommonRequest(input)
	return out, req.Send()
}

// DescribeOriginStatusCodeRankingCommonWithContext is the same as DescribeOriginStatusCodeRankingCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeOriginStatusCodeRankingCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CDN) DescribeOriginStatusCodeRankingCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeOriginStatusCodeRankingCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeOriginStatusCodeRanking = "DescribeOriginStatusCodeRanking"

// DescribeOriginStatusCodeRankingRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeOriginStatusCodeRanking operation. The "output" return
// value will be populated with the DescribeOriginStatusCodeRankingCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeOriginStatusCodeRankingCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeOriginStatusCodeRankingCommon Send returns without error.
//
// See DescribeOriginStatusCodeRanking for more information on using the DescribeOriginStatusCodeRanking
// API call, and error handling.
//
//    // Example sending a request using the DescribeOriginStatusCodeRankingRequest method.
//    req, resp := client.DescribeOriginStatusCodeRankingRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CDN) DescribeOriginStatusCodeRankingRequest(input *DescribeOriginStatusCodeRankingInput) (req *request.Request, output *DescribeOriginStatusCodeRankingOutput) {
	op := &request.Operation{
		Name:       opDescribeOriginStatusCodeRanking,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeOriginStatusCodeRankingInput{}
	}

	output = &DescribeOriginStatusCodeRankingOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// DescribeOriginStatusCodeRanking API operation for CDN.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CDN's
// API operation DescribeOriginStatusCodeRanking for usage and error information.
func (c *CDN) DescribeOriginStatusCodeRanking(input *DescribeOriginStatusCodeRankingInput) (*DescribeOriginStatusCodeRankingOutput, error) {
	req, out := c.DescribeOriginStatusCodeRankingRequest(input)
	return out, req.Send()
}

// DescribeOriginStatusCodeRankingWithContext is the same as DescribeOriginStatusCodeRanking with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeOriginStatusCodeRanking for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CDN) DescribeOriginStatusCodeRankingWithContext(ctx volcengine.Context, input *DescribeOriginStatusCodeRankingInput, opts ...request.Option) (*DescribeOriginStatusCodeRankingOutput, error) {
	req, out := c.DescribeOriginStatusCodeRankingRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DescribeOriginStatusCodeRankingInput struct {
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
func (s DescribeOriginStatusCodeRankingInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeOriginStatusCodeRankingInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeOriginStatusCodeRankingInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DescribeOriginStatusCodeRankingInput"}
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
func (s *DescribeOriginStatusCodeRankingInput) SetDomain(v string) *DescribeOriginStatusCodeRankingInput {
	s.Domain = &v
	return s
}

// SetEndTime sets the EndTime field's value.
func (s *DescribeOriginStatusCodeRankingInput) SetEndTime(v int64) *DescribeOriginStatusCodeRankingInput {
	s.EndTime = &v
	return s
}

// SetInterval sets the Interval field's value.
func (s *DescribeOriginStatusCodeRankingInput) SetInterval(v string) *DescribeOriginStatusCodeRankingInput {
	s.Interval = &v
	return s
}

// SetItem sets the Item field's value.
func (s *DescribeOriginStatusCodeRankingInput) SetItem(v string) *DescribeOriginStatusCodeRankingInput {
	s.Item = &v
	return s
}

// SetMetric sets the Metric field's value.
func (s *DescribeOriginStatusCodeRankingInput) SetMetric(v string) *DescribeOriginStatusCodeRankingInput {
	s.Metric = &v
	return s
}

// SetProject sets the Project field's value.
func (s *DescribeOriginStatusCodeRankingInput) SetProject(v string) *DescribeOriginStatusCodeRankingInput {
	s.Project = &v
	return s
}

// SetStartTime sets the StartTime field's value.
func (s *DescribeOriginStatusCodeRankingInput) SetStartTime(v int64) *DescribeOriginStatusCodeRankingInput {
	s.StartTime = &v
	return s
}

type DescribeOriginStatusCodeRankingOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	Item *string `type:"string"`

	Metric *string `type:"string"`

	TopDataDetails []*TopDataDetailForDescribeOriginStatusCodeRankingOutput `type:"list"`
}

// String returns the string representation
func (s DescribeOriginStatusCodeRankingOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeOriginStatusCodeRankingOutput) GoString() string {
	return s.String()
}

// SetItem sets the Item field's value.
func (s *DescribeOriginStatusCodeRankingOutput) SetItem(v string) *DescribeOriginStatusCodeRankingOutput {
	s.Item = &v
	return s
}

// SetMetric sets the Metric field's value.
func (s *DescribeOriginStatusCodeRankingOutput) SetMetric(v string) *DescribeOriginStatusCodeRankingOutput {
	s.Metric = &v
	return s
}

// SetTopDataDetails sets the TopDataDetails field's value.
func (s *DescribeOriginStatusCodeRankingOutput) SetTopDataDetails(v []*TopDataDetailForDescribeOriginStatusCodeRankingOutput) *DescribeOriginStatusCodeRankingOutput {
	s.TopDataDetails = v
	return s
}

type TopDataDetailForDescribeOriginStatusCodeRankingOutput struct {
	_ struct{} `type:"structure"`

	ItemKey *string `type:"string"`

	Status2xx *float64 `type:"double"`

	Status2xxRatio *float64 `type:"double"`

	Status3xx *float64 `type:"double"`

	Status3xxRatio *float64 `type:"double"`

	Status4xx *float64 `type:"double"`

	Status4xxRatio *float64 `type:"double"`

	Status5xx *float64 `type:"double"`

	Status5xxRatio *float64 `type:"double"`
}

// String returns the string representation
func (s TopDataDetailForDescribeOriginStatusCodeRankingOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s TopDataDetailForDescribeOriginStatusCodeRankingOutput) GoString() string {
	return s.String()
}

// SetItemKey sets the ItemKey field's value.
func (s *TopDataDetailForDescribeOriginStatusCodeRankingOutput) SetItemKey(v string) *TopDataDetailForDescribeOriginStatusCodeRankingOutput {
	s.ItemKey = &v
	return s
}

// SetStatus2xx sets the Status2xx field's value.
func (s *TopDataDetailForDescribeOriginStatusCodeRankingOutput) SetStatus2xx(v float64) *TopDataDetailForDescribeOriginStatusCodeRankingOutput {
	s.Status2xx = &v
	return s
}

// SetStatus2xxRatio sets the Status2xxRatio field's value.
func (s *TopDataDetailForDescribeOriginStatusCodeRankingOutput) SetStatus2xxRatio(v float64) *TopDataDetailForDescribeOriginStatusCodeRankingOutput {
	s.Status2xxRatio = &v
	return s
}

// SetStatus3xx sets the Status3xx field's value.
func (s *TopDataDetailForDescribeOriginStatusCodeRankingOutput) SetStatus3xx(v float64) *TopDataDetailForDescribeOriginStatusCodeRankingOutput {
	s.Status3xx = &v
	return s
}

// SetStatus3xxRatio sets the Status3xxRatio field's value.
func (s *TopDataDetailForDescribeOriginStatusCodeRankingOutput) SetStatus3xxRatio(v float64) *TopDataDetailForDescribeOriginStatusCodeRankingOutput {
	s.Status3xxRatio = &v
	return s
}

// SetStatus4xx sets the Status4xx field's value.
func (s *TopDataDetailForDescribeOriginStatusCodeRankingOutput) SetStatus4xx(v float64) *TopDataDetailForDescribeOriginStatusCodeRankingOutput {
	s.Status4xx = &v
	return s
}

// SetStatus4xxRatio sets the Status4xxRatio field's value.
func (s *TopDataDetailForDescribeOriginStatusCodeRankingOutput) SetStatus4xxRatio(v float64) *TopDataDetailForDescribeOriginStatusCodeRankingOutput {
	s.Status4xxRatio = &v
	return s
}

// SetStatus5xx sets the Status5xx field's value.
func (s *TopDataDetailForDescribeOriginStatusCodeRankingOutput) SetStatus5xx(v float64) *TopDataDetailForDescribeOriginStatusCodeRankingOutput {
	s.Status5xx = &v
	return s
}

// SetStatus5xxRatio sets the Status5xxRatio field's value.
func (s *TopDataDetailForDescribeOriginStatusCodeRankingOutput) SetStatus5xxRatio(v float64) *TopDataDetailForDescribeOriginStatusCodeRankingOutput {
	s.Status5xxRatio = &v
	return s
}
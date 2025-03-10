// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package cdn

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribeStatisticalRankingCommon = "DescribeStatisticalRanking"

// DescribeStatisticalRankingCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeStatisticalRankingCommon operation. The "output" return
// value will be populated with the DescribeStatisticalRankingCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeStatisticalRankingCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeStatisticalRankingCommon Send returns without error.
//
// See DescribeStatisticalRankingCommon for more information on using the DescribeStatisticalRankingCommon
// API call, and error handling.
//
//    // Example sending a request using the DescribeStatisticalRankingCommonRequest method.
//    req, resp := client.DescribeStatisticalRankingCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CDN) DescribeStatisticalRankingCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeStatisticalRankingCommon,
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

// DescribeStatisticalRankingCommon API operation for CDN.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CDN's
// API operation DescribeStatisticalRankingCommon for usage and error information.
func (c *CDN) DescribeStatisticalRankingCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeStatisticalRankingCommonRequest(input)
	return out, req.Send()
}

// DescribeStatisticalRankingCommonWithContext is the same as DescribeStatisticalRankingCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeStatisticalRankingCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CDN) DescribeStatisticalRankingCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeStatisticalRankingCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeStatisticalRanking = "DescribeStatisticalRanking"

// DescribeStatisticalRankingRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeStatisticalRanking operation. The "output" return
// value will be populated with the DescribeStatisticalRankingCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeStatisticalRankingCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeStatisticalRankingCommon Send returns without error.
//
// See DescribeStatisticalRanking for more information on using the DescribeStatisticalRanking
// API call, and error handling.
//
//    // Example sending a request using the DescribeStatisticalRankingRequest method.
//    req, resp := client.DescribeStatisticalRankingRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CDN) DescribeStatisticalRankingRequest(input *DescribeStatisticalRankingInput) (req *request.Request, output *DescribeStatisticalRankingOutput) {
	op := &request.Operation{
		Name:       opDescribeStatisticalRanking,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeStatisticalRankingInput{}
	}

	output = &DescribeStatisticalRankingOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// DescribeStatisticalRanking API operation for CDN.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CDN's
// API operation DescribeStatisticalRanking for usage and error information.
func (c *CDN) DescribeStatisticalRanking(input *DescribeStatisticalRankingInput) (*DescribeStatisticalRankingOutput, error) {
	req, out := c.DescribeStatisticalRankingRequest(input)
	return out, req.Send()
}

// DescribeStatisticalRankingWithContext is the same as DescribeStatisticalRanking with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeStatisticalRanking for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CDN) DescribeStatisticalRankingWithContext(ctx volcengine.Context, input *DescribeStatisticalRankingInput, opts ...request.Option) (*DescribeStatisticalRankingOutput, error) {
	req, out := c.DescribeStatisticalRankingRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DescribeStatisticalRankingInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Area *string `type:"string" json:",omitempty"`

	// Domain is a required field
	Domain *string `type:"string" json:",omitempty" required:"true"`

	// EndTime is a required field
	EndTime *int64 `type:"int64" json:",omitempty" required:"true"`

	// Item is a required field
	Item *string `type:"string" json:",omitempty" required:"true"`

	// Metric is a required field
	Metric *string `type:"string" json:",omitempty" required:"true"`

	// StartTime is a required field
	StartTime *int64 `type:"int64" json:",omitempty" required:"true"`

	UaType *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s DescribeStatisticalRankingInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeStatisticalRankingInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeStatisticalRankingInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DescribeStatisticalRankingInput"}
	if s.Domain == nil {
		invalidParams.Add(request.NewErrParamRequired("Domain"))
	}
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

// SetArea sets the Area field's value.
func (s *DescribeStatisticalRankingInput) SetArea(v string) *DescribeStatisticalRankingInput {
	s.Area = &v
	return s
}

// SetDomain sets the Domain field's value.
func (s *DescribeStatisticalRankingInput) SetDomain(v string) *DescribeStatisticalRankingInput {
	s.Domain = &v
	return s
}

// SetEndTime sets the EndTime field's value.
func (s *DescribeStatisticalRankingInput) SetEndTime(v int64) *DescribeStatisticalRankingInput {
	s.EndTime = &v
	return s
}

// SetItem sets the Item field's value.
func (s *DescribeStatisticalRankingInput) SetItem(v string) *DescribeStatisticalRankingInput {
	s.Item = &v
	return s
}

// SetMetric sets the Metric field's value.
func (s *DescribeStatisticalRankingInput) SetMetric(v string) *DescribeStatisticalRankingInput {
	s.Metric = &v
	return s
}

// SetStartTime sets the StartTime field's value.
func (s *DescribeStatisticalRankingInput) SetStartTime(v int64) *DescribeStatisticalRankingInput {
	s.StartTime = &v
	return s
}

// SetUaType sets the UaType field's value.
func (s *DescribeStatisticalRankingInput) SetUaType(v string) *DescribeStatisticalRankingInput {
	s.UaType = &v
	return s
}

type DescribeStatisticalRankingOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	Item *string `type:"string" json:",omitempty"`

	Metric *string `type:"string" json:",omitempty"`

	RankingDataList []*RankingDataListForDescribeStatisticalRankingOutput `type:"list" json:",omitempty"`

	UaType *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s DescribeStatisticalRankingOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeStatisticalRankingOutput) GoString() string {
	return s.String()
}

// SetItem sets the Item field's value.
func (s *DescribeStatisticalRankingOutput) SetItem(v string) *DescribeStatisticalRankingOutput {
	s.Item = &v
	return s
}

// SetMetric sets the Metric field's value.
func (s *DescribeStatisticalRankingOutput) SetMetric(v string) *DescribeStatisticalRankingOutput {
	s.Metric = &v
	return s
}

// SetRankingDataList sets the RankingDataList field's value.
func (s *DescribeStatisticalRankingOutput) SetRankingDataList(v []*RankingDataListForDescribeStatisticalRankingOutput) *DescribeStatisticalRankingOutput {
	s.RankingDataList = v
	return s
}

// SetUaType sets the UaType field's value.
func (s *DescribeStatisticalRankingOutput) SetUaType(v string) *DescribeStatisticalRankingOutput {
	s.UaType = &v
	return s
}

type RankingDataListForDescribeStatisticalRankingOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	ItemKey *string `type:"string" json:",omitempty"`

	ItemKeyCN *string `type:"string" json:",omitempty"`

	Value *float64 `type:"double" json:",omitempty"`
}

// String returns the string representation
func (s RankingDataListForDescribeStatisticalRankingOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s RankingDataListForDescribeStatisticalRankingOutput) GoString() string {
	return s.String()
}

// SetItemKey sets the ItemKey field's value.
func (s *RankingDataListForDescribeStatisticalRankingOutput) SetItemKey(v string) *RankingDataListForDescribeStatisticalRankingOutput {
	s.ItemKey = &v
	return s
}

// SetItemKeyCN sets the ItemKeyCN field's value.
func (s *RankingDataListForDescribeStatisticalRankingOutput) SetItemKeyCN(v string) *RankingDataListForDescribeStatisticalRankingOutput {
	s.ItemKeyCN = &v
	return s
}

// SetValue sets the Value field's value.
func (s *RankingDataListForDescribeStatisticalRankingOutput) SetValue(v float64) *RankingDataListForDescribeStatisticalRankingOutput {
	s.Value = &v
	return s
}

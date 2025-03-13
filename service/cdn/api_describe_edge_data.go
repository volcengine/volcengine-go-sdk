// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package cdn

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribeEdgeDataCommon = "DescribeEdgeData"

// DescribeEdgeDataCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeEdgeDataCommon operation. The "output" return
// value will be populated with the DescribeEdgeDataCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeEdgeDataCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeEdgeDataCommon Send returns without error.
//
// See DescribeEdgeDataCommon for more information on using the DescribeEdgeDataCommon
// API call, and error handling.
//
//    // Example sending a request using the DescribeEdgeDataCommonRequest method.
//    req, resp := client.DescribeEdgeDataCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CDN) DescribeEdgeDataCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeEdgeDataCommon,
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

// DescribeEdgeDataCommon API operation for CDN.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CDN's
// API operation DescribeEdgeDataCommon for usage and error information.
func (c *CDN) DescribeEdgeDataCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeEdgeDataCommonRequest(input)
	return out, req.Send()
}

// DescribeEdgeDataCommonWithContext is the same as DescribeEdgeDataCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeEdgeDataCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CDN) DescribeEdgeDataCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeEdgeDataCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeEdgeData = "DescribeEdgeData"

// DescribeEdgeDataRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeEdgeData operation. The "output" return
// value will be populated with the DescribeEdgeDataCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeEdgeDataCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeEdgeDataCommon Send returns without error.
//
// See DescribeEdgeData for more information on using the DescribeEdgeData
// API call, and error handling.
//
//    // Example sending a request using the DescribeEdgeDataRequest method.
//    req, resp := client.DescribeEdgeDataRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CDN) DescribeEdgeDataRequest(input *DescribeEdgeDataInput) (req *request.Request, output *DescribeEdgeDataOutput) {
	op := &request.Operation{
		Name:       opDescribeEdgeData,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeEdgeDataInput{}
	}

	output = &DescribeEdgeDataOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// DescribeEdgeData API operation for CDN.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CDN's
// API operation DescribeEdgeData for usage and error information.
func (c *CDN) DescribeEdgeData(input *DescribeEdgeDataInput) (*DescribeEdgeDataOutput, error) {
	req, out := c.DescribeEdgeDataRequest(input)
	return out, req.Send()
}

// DescribeEdgeDataWithContext is the same as DescribeEdgeData with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeEdgeData for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CDN) DescribeEdgeDataWithContext(ctx volcengine.Context, input *DescribeEdgeDataInput, opts ...request.Option) (*DescribeEdgeDataOutput, error) {
	req, out := c.DescribeEdgeDataRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DescribeEdgeDataInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	BillingRegion *string `type:"string" json:",omitempty"`

	Domain *string `type:"string" json:",omitempty"`

	// EndTime is a required field
	EndTime *int64 `type:"int64" json:",omitempty" required:"true"`

	Interval *string `type:"string" json:",omitempty"`

	// Metric is a required field
	Metric *string `type:"string" json:",omitempty" required:"true"`

	Project *string `type:"string" json:",omitempty"`

	// StartTime is a required field
	StartTime *int64 `type:"int64" json:",omitempty" required:"true"`
}

// String returns the string representation
func (s DescribeEdgeDataInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeEdgeDataInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeEdgeDataInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DescribeEdgeDataInput"}
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
func (s *DescribeEdgeDataInput) SetBillingRegion(v string) *DescribeEdgeDataInput {
	s.BillingRegion = &v
	return s
}

// SetDomain sets the Domain field's value.
func (s *DescribeEdgeDataInput) SetDomain(v string) *DescribeEdgeDataInput {
	s.Domain = &v
	return s
}

// SetEndTime sets the EndTime field's value.
func (s *DescribeEdgeDataInput) SetEndTime(v int64) *DescribeEdgeDataInput {
	s.EndTime = &v
	return s
}

// SetInterval sets the Interval field's value.
func (s *DescribeEdgeDataInput) SetInterval(v string) *DescribeEdgeDataInput {
	s.Interval = &v
	return s
}

// SetMetric sets the Metric field's value.
func (s *DescribeEdgeDataInput) SetMetric(v string) *DescribeEdgeDataInput {
	s.Metric = &v
	return s
}

// SetProject sets the Project field's value.
func (s *DescribeEdgeDataInput) SetProject(v string) *DescribeEdgeDataInput {
	s.Project = &v
	return s
}

// SetStartTime sets the StartTime field's value.
func (s *DescribeEdgeDataInput) SetStartTime(v int64) *DescribeEdgeDataInput {
	s.StartTime = &v
	return s
}

type DescribeEdgeDataOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	MetricDataList []*MetricDataListForDescribeEdgeDataOutput `type:"list" json:",omitempty"`
}

// String returns the string representation
func (s DescribeEdgeDataOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeEdgeDataOutput) GoString() string {
	return s.String()
}

// SetMetricDataList sets the MetricDataList field's value.
func (s *DescribeEdgeDataOutput) SetMetricDataList(v []*MetricDataListForDescribeEdgeDataOutput) *DescribeEdgeDataOutput {
	s.MetricDataList = v
	return s
}

type MetricDataListForDescribeEdgeDataOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metric *string `type:"string" json:",omitempty"`

	Values []*ValueForDescribeEdgeDataOutput `type:"list" json:",omitempty"`
}

// String returns the string representation
func (s MetricDataListForDescribeEdgeDataOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s MetricDataListForDescribeEdgeDataOutput) GoString() string {
	return s.String()
}

// SetMetric sets the Metric field's value.
func (s *MetricDataListForDescribeEdgeDataOutput) SetMetric(v string) *MetricDataListForDescribeEdgeDataOutput {
	s.Metric = &v
	return s
}

// SetValues sets the Values field's value.
func (s *MetricDataListForDescribeEdgeDataOutput) SetValues(v []*ValueForDescribeEdgeDataOutput) *MetricDataListForDescribeEdgeDataOutput {
	s.Values = v
	return s
}

type ValueForDescribeEdgeDataOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Item *string `type:"string" json:",omitempty"`

	TimeStamp *int64 `type:"int64" json:",omitempty"`

	Value *float64 `type:"double" json:",omitempty"`
}

// String returns the string representation
func (s ValueForDescribeEdgeDataOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ValueForDescribeEdgeDataOutput) GoString() string {
	return s.String()
}

// SetItem sets the Item field's value.
func (s *ValueForDescribeEdgeDataOutput) SetItem(v string) *ValueForDescribeEdgeDataOutput {
	s.Item = &v
	return s
}

// SetTimeStamp sets the TimeStamp field's value.
func (s *ValueForDescribeEdgeDataOutput) SetTimeStamp(v int64) *ValueForDescribeEdgeDataOutput {
	s.TimeStamp = &v
	return s
}

// SetValue sets the Value field's value.
func (s *ValueForDescribeEdgeDataOutput) SetValue(v float64) *ValueForDescribeEdgeDataOutput {
	s.Value = &v
	return s
}

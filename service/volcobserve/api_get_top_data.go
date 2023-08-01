// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package volcobserve

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opGetTopDataCommon = "GetTopData"

// GetTopDataCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the GetTopDataCommon operation. The "output" return
// value will be populated with the GetTopDataCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned GetTopDataCommon Request to send the API call to the service.
// the "output" return value is not valid until after GetTopDataCommon Send returns without error.
//
// See GetTopDataCommon for more information on using the GetTopDataCommon
// API call, and error handling.
//
//	// Example sending a request using the GetTopDataCommonRequest method.
//	req, resp := client.GetTopDataCommonRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *VOLCOBSERVE) GetTopDataCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opGetTopDataCommon,
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

// GetTopDataCommon API operation for VOLC_OBSERVE.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VOLC_OBSERVE's
// API operation GetTopDataCommon for usage and error information.
func (c *VOLCOBSERVE) GetTopDataCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.GetTopDataCommonRequest(input)
	return out, req.Send()
}

// GetTopDataCommonWithContext is the same as GetTopDataCommon with the addition of
// the ability to pass a context and additional request options.
//
// See GetTopDataCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VOLCOBSERVE) GetTopDataCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.GetTopDataCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opGetTopData = "GetTopData"

// GetTopDataRequest generates a "volcengine/request.Request" representing the
// client's request for the GetTopData operation. The "output" return
// value will be populated with the GetTopDataCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned GetTopDataCommon Request to send the API call to the service.
// the "output" return value is not valid until after GetTopDataCommon Send returns without error.
//
// See GetTopData for more information on using the GetTopData
// API call, and error handling.
//
//	// Example sending a request using the GetTopDataRequest method.
//	req, resp := client.GetTopDataRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *VOLCOBSERVE) GetTopDataRequest(input *GetTopDataInput) (req *request.Request, output *GetTopDataOutput) {
	op := &request.Operation{
		Name:       opGetTopData,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetTopDataInput{}
	}

	output = &GetTopDataOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// GetTopData API operation for VOLC_OBSERVE.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VOLC_OBSERVE's
// API operation GetTopData for usage and error information.
func (c *VOLCOBSERVE) GetTopData(input *GetTopDataInput) (*GetTopDataOutput, error) {
	req, out := c.GetTopDataRequest(input)
	return out, req.Send()
}

// GetTopDataWithContext is the same as GetTopData with the addition of
// the ability to pass a context and additional request options.
//
// See GetTopData for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VOLCOBSERVE) GetTopDataWithContext(ctx volcengine.Context, input *GetTopDataInput, opts ...request.Option) (*GetTopDataOutput, error) {
	req, out := c.GetTopDataRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DataForGetTopDataOutput struct {
	_ struct{} `type:"structure"`

	Asc *bool `type:"boolean"`

	OrderByMetricName *string `type:"string"`

	TopDataResults []*TopDataResultForGetTopDataOutput `type:"list"`
}

// String returns the string representation
func (s DataForGetTopDataOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DataForGetTopDataOutput) GoString() string {
	return s.String()
}

// SetAsc sets the Asc field's value.
func (s *DataForGetTopDataOutput) SetAsc(v bool) *DataForGetTopDataOutput {
	s.Asc = &v
	return s
}

// SetOrderByMetricName sets the OrderByMetricName field's value.
func (s *DataForGetTopDataOutput) SetOrderByMetricName(v string) *DataForGetTopDataOutput {
	s.OrderByMetricName = &v
	return s
}

// SetTopDataResults sets the TopDataResults field's value.
func (s *DataForGetTopDataOutput) SetTopDataResults(v []*TopDataResultForGetTopDataOutput) *DataForGetTopDataOutput {
	s.TopDataResults = v
	return s
}

type DimensionForGetTopDataInput struct {
	_ struct{} `type:"structure"`

	Name *string `type:"string"`

	Value *string `type:"string"`
}

// String returns the string representation
func (s DimensionForGetTopDataInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DimensionForGetTopDataInput) GoString() string {
	return s.String()
}

// SetName sets the Name field's value.
func (s *DimensionForGetTopDataInput) SetName(v string) *DimensionForGetTopDataInput {
	s.Name = &v
	return s
}

// SetValue sets the Value field's value.
func (s *DimensionForGetTopDataInput) SetValue(v string) *DimensionForGetTopDataInput {
	s.Value = &v
	return s
}

type GetTopDataInput struct {
	_ struct{} `type:"structure"`

	Asc *bool `type:"boolean"`

	EndTime *int64 `type:"integer"`

	GroupBy []*string `type:"list"`

	Instances []*InstanceForGetTopDataInput `type:"list"`

	MetricNames []*string `type:"list"`

	Namespace *string `type:"string"`

	OrderByMetricName *string `type:"string"`

	StartTime *int64 `type:"integer"`

	SubNamespace *string `type:"string"`

	TopN *int64 `type:"integer"`
}

// String returns the string representation
func (s GetTopDataInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s GetTopDataInput) GoString() string {
	return s.String()
}

// SetAsc sets the Asc field's value.
func (s *GetTopDataInput) SetAsc(v bool) *GetTopDataInput {
	s.Asc = &v
	return s
}

// SetEndTime sets the EndTime field's value.
func (s *GetTopDataInput) SetEndTime(v int64) *GetTopDataInput {
	s.EndTime = &v
	return s
}

// SetGroupBy sets the GroupBy field's value.
func (s *GetTopDataInput) SetGroupBy(v []*string) *GetTopDataInput {
	s.GroupBy = v
	return s
}

// SetInstances sets the Instances field's value.
func (s *GetTopDataInput) SetInstances(v []*InstanceForGetTopDataInput) *GetTopDataInput {
	s.Instances = v
	return s
}

// SetMetricNames sets the MetricNames field's value.
func (s *GetTopDataInput) SetMetricNames(v []*string) *GetTopDataInput {
	s.MetricNames = v
	return s
}

// SetNamespace sets the Namespace field's value.
func (s *GetTopDataInput) SetNamespace(v string) *GetTopDataInput {
	s.Namespace = &v
	return s
}

// SetOrderByMetricName sets the OrderByMetricName field's value.
func (s *GetTopDataInput) SetOrderByMetricName(v string) *GetTopDataInput {
	s.OrderByMetricName = &v
	return s
}

// SetStartTime sets the StartTime field's value.
func (s *GetTopDataInput) SetStartTime(v int64) *GetTopDataInput {
	s.StartTime = &v
	return s
}

// SetSubNamespace sets the SubNamespace field's value.
func (s *GetTopDataInput) SetSubNamespace(v string) *GetTopDataInput {
	s.SubNamespace = &v
	return s
}

// SetTopN sets the TopN field's value.
func (s *GetTopDataInput) SetTopN(v int64) *GetTopDataInput {
	s.TopN = &v
	return s
}

type GetTopDataOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	Data *DataForGetTopDataOutput `type:"structure"`
}

// String returns the string representation
func (s GetTopDataOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s GetTopDataOutput) GoString() string {
	return s.String()
}

// SetData sets the Data field's value.
func (s *GetTopDataOutput) SetData(v *DataForGetTopDataOutput) *GetTopDataOutput {
	s.Data = v
	return s
}

type InstanceForGetTopDataInput struct {
	_ struct{} `type:"structure"`

	Dimensions []*DimensionForGetTopDataInput `type:"list"`
}

// String returns the string representation
func (s InstanceForGetTopDataInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s InstanceForGetTopDataInput) GoString() string {
	return s.String()
}

// SetDimensions sets the Dimensions field's value.
func (s *InstanceForGetTopDataInput) SetDimensions(v []*DimensionForGetTopDataInput) *InstanceForGetTopDataInput {
	s.Dimensions = v
	return s
}

type TopDataResultForGetTopDataOutput struct {
	_ struct{} `type:"structure"`

	GroupKeys map[string]*interface{} `type:"map"`

	MetricData map[string]*interface{} `type:"map"`
}

// String returns the string representation
func (s TopDataResultForGetTopDataOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s TopDataResultForGetTopDataOutput) GoString() string {
	return s.String()
}

// SetGroupKeys sets the GroupKeys field's value.
func (s *TopDataResultForGetTopDataOutput) SetGroupKeys(v map[string]*interface{}) *TopDataResultForGetTopDataOutput {
	s.GroupKeys = v
	return s
}

// SetMetricData sets the MetricData field's value.
func (s *TopDataResultForGetTopDataOutput) SetMetricData(v map[string]*interface{}) *TopDataResultForGetTopDataOutput {
	s.MetricData = v
	return s
}
// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package cdn

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribeUserDataCommon = "DescribeUserData"

// DescribeUserDataCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeUserDataCommon operation. The "output" return
// value will be populated with the DescribeUserDataCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeUserDataCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeUserDataCommon Send returns without error.
//
// See DescribeUserDataCommon for more information on using the DescribeUserDataCommon
// API call, and error handling.
//
//    // Example sending a request using the DescribeUserDataCommonRequest method.
//    req, resp := client.DescribeUserDataCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CDN) DescribeUserDataCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeUserDataCommon,
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

// DescribeUserDataCommon API operation for CDN.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CDN's
// API operation DescribeUserDataCommon for usage and error information.
func (c *CDN) DescribeUserDataCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeUserDataCommonRequest(input)
	return out, req.Send()
}

// DescribeUserDataCommonWithContext is the same as DescribeUserDataCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeUserDataCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CDN) DescribeUserDataCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeUserDataCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeUserData = "DescribeUserData"

// DescribeUserDataRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeUserData operation. The "output" return
// value will be populated with the DescribeUserDataCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeUserDataCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeUserDataCommon Send returns without error.
//
// See DescribeUserData for more information on using the DescribeUserData
// API call, and error handling.
//
//    // Example sending a request using the DescribeUserDataRequest method.
//    req, resp := client.DescribeUserDataRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CDN) DescribeUserDataRequest(input *DescribeUserDataInput) (req *request.Request, output *DescribeUserDataOutput) {
	op := &request.Operation{
		Name:       opDescribeUserData,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeUserDataInput{}
	}

	output = &DescribeUserDataOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// DescribeUserData API operation for CDN.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CDN's
// API operation DescribeUserData for usage and error information.
func (c *CDN) DescribeUserData(input *DescribeUserDataInput) (*DescribeUserDataOutput, error) {
	req, out := c.DescribeUserDataRequest(input)
	return out, req.Send()
}

// DescribeUserDataWithContext is the same as DescribeUserData with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeUserData for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CDN) DescribeUserDataWithContext(ctx volcengine.Context, input *DescribeUserDataInput, opts ...request.Option) (*DescribeUserDataOutput, error) {
	req, out := c.DescribeUserDataRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DescribeUserDataInput struct {
	_ struct{} `type:"structure"`

	// Domain is a required field
	Domain *string `type:"string" required:"true"`

	// EndTime is a required field
	EndTime *int64 `type:"int64" required:"true"`

	// Interval is a required field
	Interval *string `type:"string" required:"true"`

	IpVersion *string `type:"string"`

	Location *string `type:"string"`

	Province *string `type:"string"`

	// StartTime is a required field
	StartTime *int64 `type:"int64" required:"true"`
}

// String returns the string representation
func (s DescribeUserDataInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeUserDataInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeUserDataInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DescribeUserDataInput"}
	if s.Domain == nil {
		invalidParams.Add(request.NewErrParamRequired("Domain"))
	}
	if s.EndTime == nil {
		invalidParams.Add(request.NewErrParamRequired("EndTime"))
	}
	if s.Interval == nil {
		invalidParams.Add(request.NewErrParamRequired("Interval"))
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
func (s *DescribeUserDataInput) SetDomain(v string) *DescribeUserDataInput {
	s.Domain = &v
	return s
}

// SetEndTime sets the EndTime field's value.
func (s *DescribeUserDataInput) SetEndTime(v int64) *DescribeUserDataInput {
	s.EndTime = &v
	return s
}

// SetInterval sets the Interval field's value.
func (s *DescribeUserDataInput) SetInterval(v string) *DescribeUserDataInput {
	s.Interval = &v
	return s
}

// SetIpVersion sets the IpVersion field's value.
func (s *DescribeUserDataInput) SetIpVersion(v string) *DescribeUserDataInput {
	s.IpVersion = &v
	return s
}

// SetLocation sets the Location field's value.
func (s *DescribeUserDataInput) SetLocation(v string) *DescribeUserDataInput {
	s.Location = &v
	return s
}

// SetProvince sets the Province field's value.
func (s *DescribeUserDataInput) SetProvince(v string) *DescribeUserDataInput {
	s.Province = &v
	return s
}

// SetStartTime sets the StartTime field's value.
func (s *DescribeUserDataInput) SetStartTime(v int64) *DescribeUserDataInput {
	s.StartTime = &v
	return s
}

type DescribeUserDataOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	MetricDataList []*MetricDataListForDescribeUserDataOutput `type:"list"`
}

// String returns the string representation
func (s DescribeUserDataOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeUserDataOutput) GoString() string {
	return s.String()
}

// SetMetricDataList sets the MetricDataList field's value.
func (s *DescribeUserDataOutput) SetMetricDataList(v []*MetricDataListForDescribeUserDataOutput) *DescribeUserDataOutput {
	s.MetricDataList = v
	return s
}

type MetricDataListForDescribeUserDataOutput struct {
	_ struct{} `type:"structure"`

	TimeStamp *int64 `type:"int64"`

	Value *float64 `type:"double"`
}

// String returns the string representation
func (s MetricDataListForDescribeUserDataOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s MetricDataListForDescribeUserDataOutput) GoString() string {
	return s.String()
}

// SetTimeStamp sets the TimeStamp field's value.
func (s *MetricDataListForDescribeUserDataOutput) SetTimeStamp(v int64) *MetricDataListForDescribeUserDataOutput {
	s.TimeStamp = &v
	return s
}

// SetValue sets the Value field's value.
func (s *MetricDataListForDescribeUserDataOutput) SetValue(v float64) *MetricDataListForDescribeUserDataOutput {
	s.Value = &v
	return s
}
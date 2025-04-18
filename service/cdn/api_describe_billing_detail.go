// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package cdn

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribeBillingDetailCommon = "DescribeBillingDetail"

// DescribeBillingDetailCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeBillingDetailCommon operation. The "output" return
// value will be populated with the DescribeBillingDetailCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeBillingDetailCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeBillingDetailCommon Send returns without error.
//
// See DescribeBillingDetailCommon for more information on using the DescribeBillingDetailCommon
// API call, and error handling.
//
//    // Example sending a request using the DescribeBillingDetailCommonRequest method.
//    req, resp := client.DescribeBillingDetailCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CDN) DescribeBillingDetailCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeBillingDetailCommon,
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

// DescribeBillingDetailCommon API operation for CDN.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CDN's
// API operation DescribeBillingDetailCommon for usage and error information.
func (c *CDN) DescribeBillingDetailCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeBillingDetailCommonRequest(input)
	return out, req.Send()
}

// DescribeBillingDetailCommonWithContext is the same as DescribeBillingDetailCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeBillingDetailCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CDN) DescribeBillingDetailCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeBillingDetailCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeBillingDetail = "DescribeBillingDetail"

// DescribeBillingDetailRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeBillingDetail operation. The "output" return
// value will be populated with the DescribeBillingDetailCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeBillingDetailCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeBillingDetailCommon Send returns without error.
//
// See DescribeBillingDetail for more information on using the DescribeBillingDetail
// API call, and error handling.
//
//    // Example sending a request using the DescribeBillingDetailRequest method.
//    req, resp := client.DescribeBillingDetailRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CDN) DescribeBillingDetailRequest(input *DescribeBillingDetailInput) (req *request.Request, output *DescribeBillingDetailOutput) {
	op := &request.Operation{
		Name:       opDescribeBillingDetail,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeBillingDetailInput{}
	}

	output = &DescribeBillingDetailOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// DescribeBillingDetail API operation for CDN.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CDN's
// API operation DescribeBillingDetail for usage and error information.
func (c *CDN) DescribeBillingDetail(input *DescribeBillingDetailInput) (*DescribeBillingDetailOutput, error) {
	req, out := c.DescribeBillingDetailRequest(input)
	return out, req.Send()
}

// DescribeBillingDetailWithContext is the same as DescribeBillingDetail with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeBillingDetail for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CDN) DescribeBillingDetailWithContext(ctx volcengine.Context, input *DescribeBillingDetailInput, opts ...request.Option) (*DescribeBillingDetailOutput, error) {
	req, out := c.DescribeBillingDetailRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DescribeBillingDetailInput struct {
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

	TimeZone *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s DescribeBillingDetailInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeBillingDetailInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeBillingDetailInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DescribeBillingDetailInput"}
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
func (s *DescribeBillingDetailInput) SetBillingRegion(v string) *DescribeBillingDetailInput {
	s.BillingRegion = &v
	return s
}

// SetDomain sets the Domain field's value.
func (s *DescribeBillingDetailInput) SetDomain(v string) *DescribeBillingDetailInput {
	s.Domain = &v
	return s
}

// SetEndTime sets the EndTime field's value.
func (s *DescribeBillingDetailInput) SetEndTime(v int64) *DescribeBillingDetailInput {
	s.EndTime = &v
	return s
}

// SetInterval sets the Interval field's value.
func (s *DescribeBillingDetailInput) SetInterval(v string) *DescribeBillingDetailInput {
	s.Interval = &v
	return s
}

// SetMetric sets the Metric field's value.
func (s *DescribeBillingDetailInput) SetMetric(v string) *DescribeBillingDetailInput {
	s.Metric = &v
	return s
}

// SetProject sets the Project field's value.
func (s *DescribeBillingDetailInput) SetProject(v string) *DescribeBillingDetailInput {
	s.Project = &v
	return s
}

// SetStartTime sets the StartTime field's value.
func (s *DescribeBillingDetailInput) SetStartTime(v int64) *DescribeBillingDetailInput {
	s.StartTime = &v
	return s
}

// SetTimeZone sets the TimeZone field's value.
func (s *DescribeBillingDetailInput) SetTimeZone(v string) *DescribeBillingDetailInput {
	s.TimeZone = &v
	return s
}

type DescribeBillingDetailOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	TimestampPoint []*TimestampPointForDescribeBillingDetailOutput `type:"list" json:",omitempty"`
}

// String returns the string representation
func (s DescribeBillingDetailOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeBillingDetailOutput) GoString() string {
	return s.String()
}

// SetTimestampPoint sets the TimestampPoint field's value.
func (s *DescribeBillingDetailOutput) SetTimestampPoint(v []*TimestampPointForDescribeBillingDetailOutput) *DescribeBillingDetailOutput {
	s.TimestampPoint = v
	return s
}

type TimestampPointForDescribeBillingDetailOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Item *string `type:"string" json:",omitempty"`

	TimeStamp *int64 `type:"int64" json:",omitempty"`

	Value *float64 `type:"double" json:",omitempty"`
}

// String returns the string representation
func (s TimestampPointForDescribeBillingDetailOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s TimestampPointForDescribeBillingDetailOutput) GoString() string {
	return s.String()
}

// SetItem sets the Item field's value.
func (s *TimestampPointForDescribeBillingDetailOutput) SetItem(v string) *TimestampPointForDescribeBillingDetailOutput {
	s.Item = &v
	return s
}

// SetTimeStamp sets the TimeStamp field's value.
func (s *TimestampPointForDescribeBillingDetailOutput) SetTimeStamp(v int64) *TimestampPointForDescribeBillingDetailOutput {
	s.TimeStamp = &v
	return s
}

// SetValue sets the Value field's value.
func (s *TimestampPointForDescribeBillingDetailOutput) SetValue(v float64) *TimestampPointForDescribeBillingDetailOutput {
	s.Value = &v
	return s
}

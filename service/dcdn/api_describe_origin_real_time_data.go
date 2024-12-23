// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package dcdn

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribeOriginRealTimeDataCommon = "DescribeOriginRealTimeData"

// DescribeOriginRealTimeDataCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeOriginRealTimeDataCommon operation. The "output" return
// value will be populated with the DescribeOriginRealTimeDataCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeOriginRealTimeDataCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeOriginRealTimeDataCommon Send returns without error.
//
// See DescribeOriginRealTimeDataCommon for more information on using the DescribeOriginRealTimeDataCommon
// API call, and error handling.
//
//    // Example sending a request using the DescribeOriginRealTimeDataCommonRequest method.
//    req, resp := client.DescribeOriginRealTimeDataCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *DCDN) DescribeOriginRealTimeDataCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeOriginRealTimeDataCommon,
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

// DescribeOriginRealTimeDataCommon API operation for DCDN.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for DCDN's
// API operation DescribeOriginRealTimeDataCommon for usage and error information.
func (c *DCDN) DescribeOriginRealTimeDataCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeOriginRealTimeDataCommonRequest(input)
	return out, req.Send()
}

// DescribeOriginRealTimeDataCommonWithContext is the same as DescribeOriginRealTimeDataCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeOriginRealTimeDataCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *DCDN) DescribeOriginRealTimeDataCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeOriginRealTimeDataCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeOriginRealTimeData = "DescribeOriginRealTimeData"

// DescribeOriginRealTimeDataRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeOriginRealTimeData operation. The "output" return
// value will be populated with the DescribeOriginRealTimeDataCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeOriginRealTimeDataCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeOriginRealTimeDataCommon Send returns without error.
//
// See DescribeOriginRealTimeData for more information on using the DescribeOriginRealTimeData
// API call, and error handling.
//
//    // Example sending a request using the DescribeOriginRealTimeDataRequest method.
//    req, resp := client.DescribeOriginRealTimeDataRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *DCDN) DescribeOriginRealTimeDataRequest(input *DescribeOriginRealTimeDataInput) (req *request.Request, output *DescribeOriginRealTimeDataOutput) {
	op := &request.Operation{
		Name:       opDescribeOriginRealTimeData,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeOriginRealTimeDataInput{}
	}

	output = &DescribeOriginRealTimeDataOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// DescribeOriginRealTimeData API operation for DCDN.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for DCDN's
// API operation DescribeOriginRealTimeData for usage and error information.
func (c *DCDN) DescribeOriginRealTimeData(input *DescribeOriginRealTimeDataInput) (*DescribeOriginRealTimeDataOutput, error) {
	req, out := c.DescribeOriginRealTimeDataRequest(input)
	return out, req.Send()
}

// DescribeOriginRealTimeDataWithContext is the same as DescribeOriginRealTimeData with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeOriginRealTimeData for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *DCDN) DescribeOriginRealTimeDataWithContext(ctx volcengine.Context, input *DescribeOriginRealTimeDataInput, opts ...request.Option) (*DescribeOriginRealTimeDataOutput, error) {
	req, out := c.DescribeOriginRealTimeDataRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DescribeOriginRealTimeDataInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Domains []*string `type:"list" json:",omitempty"`

	// EndTime is a required field
	EndTime *string `type:"string" json:",omitempty" required:"true"`

	IPVersion *string `type:"string" json:",omitempty"`

	Metrics []*string `type:"list" json:",omitempty"`

	ProjectName []*string `type:"list" json:",omitempty"`

	Protocol []*string `type:"list" json:",omitempty"`

	// StartTime is a required field
	StartTime *string `type:"string" json:",omitempty" required:"true"`

	Type *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s DescribeOriginRealTimeDataInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeOriginRealTimeDataInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeOriginRealTimeDataInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DescribeOriginRealTimeDataInput"}
	if s.EndTime == nil {
		invalidParams.Add(request.NewErrParamRequired("EndTime"))
	}
	if s.StartTime == nil {
		invalidParams.Add(request.NewErrParamRequired("StartTime"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetDomains sets the Domains field's value.
func (s *DescribeOriginRealTimeDataInput) SetDomains(v []*string) *DescribeOriginRealTimeDataInput {
	s.Domains = v
	return s
}

// SetEndTime sets the EndTime field's value.
func (s *DescribeOriginRealTimeDataInput) SetEndTime(v string) *DescribeOriginRealTimeDataInput {
	s.EndTime = &v
	return s
}

// SetIPVersion sets the IPVersion field's value.
func (s *DescribeOriginRealTimeDataInput) SetIPVersion(v string) *DescribeOriginRealTimeDataInput {
	s.IPVersion = &v
	return s
}

// SetMetrics sets the Metrics field's value.
func (s *DescribeOriginRealTimeDataInput) SetMetrics(v []*string) *DescribeOriginRealTimeDataInput {
	s.Metrics = v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *DescribeOriginRealTimeDataInput) SetProjectName(v []*string) *DescribeOriginRealTimeDataInput {
	s.ProjectName = v
	return s
}

// SetProtocol sets the Protocol field's value.
func (s *DescribeOriginRealTimeDataInput) SetProtocol(v []*string) *DescribeOriginRealTimeDataInput {
	s.Protocol = v
	return s
}

// SetStartTime sets the StartTime field's value.
func (s *DescribeOriginRealTimeDataInput) SetStartTime(v string) *DescribeOriginRealTimeDataInput {
	s.StartTime = &v
	return s
}

// SetType sets the Type field's value.
func (s *DescribeOriginRealTimeDataInput) SetType(v string) *DescribeOriginRealTimeDataInput {
	s.Type = &v
	return s
}

type DescribeOriginRealTimeDataOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	DomainCount *int32 `type:"int32" json:",omitempty"`

	EndTime *string `type:"string" json:",omitempty"`

	Metrics []*string `type:"list" json:",omitempty"`

	Results []*ResultForDescribeOriginRealTimeDataOutput `type:"list" json:",omitempty"`

	StartTime *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s DescribeOriginRealTimeDataOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeOriginRealTimeDataOutput) GoString() string {
	return s.String()
}

// SetDomainCount sets the DomainCount field's value.
func (s *DescribeOriginRealTimeDataOutput) SetDomainCount(v int32) *DescribeOriginRealTimeDataOutput {
	s.DomainCount = &v
	return s
}

// SetEndTime sets the EndTime field's value.
func (s *DescribeOriginRealTimeDataOutput) SetEndTime(v string) *DescribeOriginRealTimeDataOutput {
	s.EndTime = &v
	return s
}

// SetMetrics sets the Metrics field's value.
func (s *DescribeOriginRealTimeDataOutput) SetMetrics(v []*string) *DescribeOriginRealTimeDataOutput {
	s.Metrics = v
	return s
}

// SetResults sets the Results field's value.
func (s *DescribeOriginRealTimeDataOutput) SetResults(v []*ResultForDescribeOriginRealTimeDataOutput) *DescribeOriginRealTimeDataOutput {
	s.Results = v
	return s
}

// SetStartTime sets the StartTime field's value.
func (s *DescribeOriginRealTimeDataOutput) SetStartTime(v string) *DescribeOriginRealTimeDataOutput {
	s.StartTime = &v
	return s
}

type DetailInfoForDescribeOriginRealTimeDataOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Name *string `type:"string" json:",omitempty"`

	Value *float64 `type:"float" json:",omitempty"`
}

// String returns the string representation
func (s DetailInfoForDescribeOriginRealTimeDataOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DetailInfoForDescribeOriginRealTimeDataOutput) GoString() string {
	return s.String()
}

// SetName sets the Name field's value.
func (s *DetailInfoForDescribeOriginRealTimeDataOutput) SetName(v string) *DetailInfoForDescribeOriginRealTimeDataOutput {
	s.Name = &v
	return s
}

// SetValue sets the Value field's value.
func (s *DetailInfoForDescribeOriginRealTimeDataOutput) SetValue(v float64) *DetailInfoForDescribeOriginRealTimeDataOutput {
	s.Value = &v
	return s
}

type RealTimeResultForDescribeOriginRealTimeDataOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	DetailInfo []*DetailInfoForDescribeOriginRealTimeDataOutput `type:"list" json:",omitempty"`

	DomainName *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s RealTimeResultForDescribeOriginRealTimeDataOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s RealTimeResultForDescribeOriginRealTimeDataOutput) GoString() string {
	return s.String()
}

// SetDetailInfo sets the DetailInfo field's value.
func (s *RealTimeResultForDescribeOriginRealTimeDataOutput) SetDetailInfo(v []*DetailInfoForDescribeOriginRealTimeDataOutput) *RealTimeResultForDescribeOriginRealTimeDataOutput {
	s.DetailInfo = v
	return s
}

// SetDomainName sets the DomainName field's value.
func (s *RealTimeResultForDescribeOriginRealTimeDataOutput) SetDomainName(v string) *RealTimeResultForDescribeOriginRealTimeDataOutput {
	s.DomainName = &v
	return s
}

type ResultForDescribeOriginRealTimeDataOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	RealTimeResults []*RealTimeResultForDescribeOriginRealTimeDataOutput `type:"list" json:",omitempty"`

	TimeStamp *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s ResultForDescribeOriginRealTimeDataOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ResultForDescribeOriginRealTimeDataOutput) GoString() string {
	return s.String()
}

// SetRealTimeResults sets the RealTimeResults field's value.
func (s *ResultForDescribeOriginRealTimeDataOutput) SetRealTimeResults(v []*RealTimeResultForDescribeOriginRealTimeDataOutput) *ResultForDescribeOriginRealTimeDataOutput {
	s.RealTimeResults = v
	return s
}

// SetTimeStamp sets the TimeStamp field's value.
func (s *ResultForDescribeOriginRealTimeDataOutput) SetTimeStamp(v string) *ResultForDescribeOriginRealTimeDataOutput {
	s.TimeStamp = &v
	return s
}

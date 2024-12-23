// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package dcdn

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribeDomainRegionDataCommon = "DescribeDomainRegionData"

// DescribeDomainRegionDataCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeDomainRegionDataCommon operation. The "output" return
// value will be populated with the DescribeDomainRegionDataCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeDomainRegionDataCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeDomainRegionDataCommon Send returns without error.
//
// See DescribeDomainRegionDataCommon for more information on using the DescribeDomainRegionDataCommon
// API call, and error handling.
//
//    // Example sending a request using the DescribeDomainRegionDataCommonRequest method.
//    req, resp := client.DescribeDomainRegionDataCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *DCDN) DescribeDomainRegionDataCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeDomainRegionDataCommon,
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

// DescribeDomainRegionDataCommon API operation for DCDN.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for DCDN's
// API operation DescribeDomainRegionDataCommon for usage and error information.
func (c *DCDN) DescribeDomainRegionDataCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeDomainRegionDataCommonRequest(input)
	return out, req.Send()
}

// DescribeDomainRegionDataCommonWithContext is the same as DescribeDomainRegionDataCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeDomainRegionDataCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *DCDN) DescribeDomainRegionDataCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeDomainRegionDataCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeDomainRegionData = "DescribeDomainRegionData"

// DescribeDomainRegionDataRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeDomainRegionData operation. The "output" return
// value will be populated with the DescribeDomainRegionDataCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeDomainRegionDataCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeDomainRegionDataCommon Send returns without error.
//
// See DescribeDomainRegionData for more information on using the DescribeDomainRegionData
// API call, and error handling.
//
//    // Example sending a request using the DescribeDomainRegionDataRequest method.
//    req, resp := client.DescribeDomainRegionDataRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *DCDN) DescribeDomainRegionDataRequest(input *DescribeDomainRegionDataInput) (req *request.Request, output *DescribeDomainRegionDataOutput) {
	op := &request.Operation{
		Name:       opDescribeDomainRegionData,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeDomainRegionDataInput{}
	}

	output = &DescribeDomainRegionDataOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// DescribeDomainRegionData API operation for DCDN.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for DCDN's
// API operation DescribeDomainRegionData for usage and error information.
func (c *DCDN) DescribeDomainRegionData(input *DescribeDomainRegionDataInput) (*DescribeDomainRegionDataOutput, error) {
	req, out := c.DescribeDomainRegionDataRequest(input)
	return out, req.Send()
}

// DescribeDomainRegionDataWithContext is the same as DescribeDomainRegionData with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeDomainRegionData for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *DCDN) DescribeDomainRegionDataWithContext(ctx volcengine.Context, input *DescribeDomainRegionDataInput, opts ...request.Option) (*DescribeDomainRegionDataOutput, error) {
	req, out := c.DescribeDomainRegionDataRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DescribeDomainRegionDataInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	// Domain is a required field
	Domain *string `type:"string" json:",omitempty" required:"true"`

	// EndTime is a required field
	EndTime *string `type:"string" json:",omitempty" required:"true"`

	ProjectName []*string `type:"list" json:",omitempty"`

	// StartTime is a required field
	StartTime *string `type:"string" json:",omitempty" required:"true"`
}

// String returns the string representation
func (s DescribeDomainRegionDataInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeDomainRegionDataInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeDomainRegionDataInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DescribeDomainRegionDataInput"}
	if s.Domain == nil {
		invalidParams.Add(request.NewErrParamRequired("Domain"))
	}
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

// SetDomain sets the Domain field's value.
func (s *DescribeDomainRegionDataInput) SetDomain(v string) *DescribeDomainRegionDataInput {
	s.Domain = &v
	return s
}

// SetEndTime sets the EndTime field's value.
func (s *DescribeDomainRegionDataInput) SetEndTime(v string) *DescribeDomainRegionDataInput {
	s.EndTime = &v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *DescribeDomainRegionDataInput) SetProjectName(v []*string) *DescribeDomainRegionDataInput {
	s.ProjectName = v
	return s
}

// SetStartTime sets the StartTime field's value.
func (s *DescribeDomainRegionDataInput) SetStartTime(v string) *DescribeDomainRegionDataInput {
	s.StartTime = &v
	return s
}

type DescribeDomainRegionDataOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	DomainName *string `type:"string" json:",omitempty"`

	EndTime *string `type:"string" json:",omitempty"`

	StartTime *string `type:"string" json:",omitempty"`

	Values []*ValueForDescribeDomainRegionDataOutput `type:"list" json:",omitempty"`
}

// String returns the string representation
func (s DescribeDomainRegionDataOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeDomainRegionDataOutput) GoString() string {
	return s.String()
}

// SetDomainName sets the DomainName field's value.
func (s *DescribeDomainRegionDataOutput) SetDomainName(v string) *DescribeDomainRegionDataOutput {
	s.DomainName = &v
	return s
}

// SetEndTime sets the EndTime field's value.
func (s *DescribeDomainRegionDataOutput) SetEndTime(v string) *DescribeDomainRegionDataOutput {
	s.EndTime = &v
	return s
}

// SetStartTime sets the StartTime field's value.
func (s *DescribeDomainRegionDataOutput) SetStartTime(v string) *DescribeDomainRegionDataOutput {
	s.StartTime = &v
	return s
}

// SetValues sets the Values field's value.
func (s *DescribeDomainRegionDataOutput) SetValues(v []*ValueForDescribeDomainRegionDataOutput) *DescribeDomainRegionDataOutput {
	s.Values = v
	return s
}

type ValueForDescribeDomainRegionDataOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	AvgResponseTime *float64 `type:"float" json:",omitempty"`

	Bandwidth *float64 `type:"float" json:",omitempty"`

	Region *string `type:"string" json:",omitempty"`

	RegionEname *string `type:"string" json:",omitempty"`

	Request *float64 `type:"float" json:",omitempty"`

	RequestProportion *float64 `type:"float" json:",omitempty"`

	Traffic *float64 `type:"float" json:",omitempty"`

	TrafficProportion *float64 `type:"float" json:",omitempty"`
}

// String returns the string representation
func (s ValueForDescribeDomainRegionDataOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ValueForDescribeDomainRegionDataOutput) GoString() string {
	return s.String()
}

// SetAvgResponseTime sets the AvgResponseTime field's value.
func (s *ValueForDescribeDomainRegionDataOutput) SetAvgResponseTime(v float64) *ValueForDescribeDomainRegionDataOutput {
	s.AvgResponseTime = &v
	return s
}

// SetBandwidth sets the Bandwidth field's value.
func (s *ValueForDescribeDomainRegionDataOutput) SetBandwidth(v float64) *ValueForDescribeDomainRegionDataOutput {
	s.Bandwidth = &v
	return s
}

// SetRegion sets the Region field's value.
func (s *ValueForDescribeDomainRegionDataOutput) SetRegion(v string) *ValueForDescribeDomainRegionDataOutput {
	s.Region = &v
	return s
}

// SetRegionEname sets the RegionEname field's value.
func (s *ValueForDescribeDomainRegionDataOutput) SetRegionEname(v string) *ValueForDescribeDomainRegionDataOutput {
	s.RegionEname = &v
	return s
}

// SetRequest sets the Request field's value.
func (s *ValueForDescribeDomainRegionDataOutput) SetRequest(v float64) *ValueForDescribeDomainRegionDataOutput {
	s.Request = &v
	return s
}

// SetRequestProportion sets the RequestProportion field's value.
func (s *ValueForDescribeDomainRegionDataOutput) SetRequestProportion(v float64) *ValueForDescribeDomainRegionDataOutput {
	s.RequestProportion = &v
	return s
}

// SetTraffic sets the Traffic field's value.
func (s *ValueForDescribeDomainRegionDataOutput) SetTraffic(v float64) *ValueForDescribeDomainRegionDataOutput {
	s.Traffic = &v
	return s
}

// SetTrafficProportion sets the TrafficProportion field's value.
func (s *ValueForDescribeDomainRegionDataOutput) SetTrafficProportion(v float64) *ValueForDescribeDomainRegionDataOutput {
	s.TrafficProportion = &v
	return s
}

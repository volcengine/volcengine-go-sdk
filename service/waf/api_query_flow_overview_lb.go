// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package waf

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opQueryFlowOverviewLbCommon = "QueryFlowOverviewLb"

// QueryFlowOverviewLbCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the QueryFlowOverviewLbCommon operation. The "output" return
// value will be populated with the QueryFlowOverviewLbCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned QueryFlowOverviewLbCommon Request to send the API call to the service.
// the "output" return value is not valid until after QueryFlowOverviewLbCommon Send returns without error.
//
// See QueryFlowOverviewLbCommon for more information on using the QueryFlowOverviewLbCommon
// API call, and error handling.
//
//    // Example sending a request using the QueryFlowOverviewLbCommonRequest method.
//    req, resp := client.QueryFlowOverviewLbCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *WAF) QueryFlowOverviewLbCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opQueryFlowOverviewLbCommon,
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

// QueryFlowOverviewLbCommon API operation for WAF.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for WAF's
// API operation QueryFlowOverviewLbCommon for usage and error information.
func (c *WAF) QueryFlowOverviewLbCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.QueryFlowOverviewLbCommonRequest(input)
	return out, req.Send()
}

// QueryFlowOverviewLbCommonWithContext is the same as QueryFlowOverviewLbCommon with the addition of
// the ability to pass a context and additional request options.
//
// See QueryFlowOverviewLbCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *WAF) QueryFlowOverviewLbCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.QueryFlowOverviewLbCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opQueryFlowOverviewLb = "QueryFlowOverviewLb"

// QueryFlowOverviewLbRequest generates a "volcengine/request.Request" representing the
// client's request for the QueryFlowOverviewLb operation. The "output" return
// value will be populated with the QueryFlowOverviewLbCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned QueryFlowOverviewLbCommon Request to send the API call to the service.
// the "output" return value is not valid until after QueryFlowOverviewLbCommon Send returns without error.
//
// See QueryFlowOverviewLb for more information on using the QueryFlowOverviewLb
// API call, and error handling.
//
//    // Example sending a request using the QueryFlowOverviewLbRequest method.
//    req, resp := client.QueryFlowOverviewLbRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *WAF) QueryFlowOverviewLbRequest(input *QueryFlowOverviewLbInput) (req *request.Request, output *QueryFlowOverviewLbOutput) {
	op := &request.Operation{
		Name:       opQueryFlowOverviewLb,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &QueryFlowOverviewLbInput{}
	}

	output = &QueryFlowOverviewLbOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// QueryFlowOverviewLb API operation for WAF.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for WAF's
// API operation QueryFlowOverviewLb for usage and error information.
func (c *WAF) QueryFlowOverviewLb(input *QueryFlowOverviewLbInput) (*QueryFlowOverviewLbOutput, error) {
	req, out := c.QueryFlowOverviewLbRequest(input)
	return out, req.Send()
}

// QueryFlowOverviewLbWithContext is the same as QueryFlowOverviewLb with the addition of
// the ability to pass a context and additional request options.
//
// See QueryFlowOverviewLb for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *WAF) QueryFlowOverviewLbWithContext(ctx volcengine.Context, input *QueryFlowOverviewLbInput, opts ...request.Option) (*QueryFlowOverviewLbOutput, error) {
	req, out := c.QueryFlowOverviewLbRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type QueryFlowOverviewLbInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	// EndTime is a required field
	EndTime *int32 `type:"int32" json:",omitempty" required:"true"`

	Host *string `type:"string" json:",omitempty"`

	// StartTime is a required field
	StartTime *int32 `type:"int32" json:",omitempty" required:"true"`
}

// String returns the string representation
func (s QueryFlowOverviewLbInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s QueryFlowOverviewLbInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *QueryFlowOverviewLbInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "QueryFlowOverviewLbInput"}
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

// SetEndTime sets the EndTime field's value.
func (s *QueryFlowOverviewLbInput) SetEndTime(v int32) *QueryFlowOverviewLbInput {
	s.EndTime = &v
	return s
}

// SetHost sets the Host field's value.
func (s *QueryFlowOverviewLbInput) SetHost(v string) *QueryFlowOverviewLbInput {
	s.Host = &v
	return s
}

// SetStartTime sets the StartTime field's value.
func (s *QueryFlowOverviewLbInput) SetStartTime(v int32) *QueryFlowOverviewLbInput {
	s.StartTime = &v
	return s
}

type QueryFlowOverviewLbOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	BackSourceCountPeak *float64 `type:"float" json:",omitempty"`

	BackSourceQPSPeak *float64 `type:"float" json:",omitempty"`

	ReqBandwidthPeak *float64 `type:"float" json:",omitempty"`

	ReqCountPeak *float64 `type:"float" json:",omitempty"`

	ReqQPSPeak *float64 `type:"float" json:",omitempty"`
}

// String returns the string representation
func (s QueryFlowOverviewLbOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s QueryFlowOverviewLbOutput) GoString() string {
	return s.String()
}

// SetBackSourceCountPeak sets the BackSourceCountPeak field's value.
func (s *QueryFlowOverviewLbOutput) SetBackSourceCountPeak(v float64) *QueryFlowOverviewLbOutput {
	s.BackSourceCountPeak = &v
	return s
}

// SetBackSourceQPSPeak sets the BackSourceQPSPeak field's value.
func (s *QueryFlowOverviewLbOutput) SetBackSourceQPSPeak(v float64) *QueryFlowOverviewLbOutput {
	s.BackSourceQPSPeak = &v
	return s
}

// SetReqBandwidthPeak sets the ReqBandwidthPeak field's value.
func (s *QueryFlowOverviewLbOutput) SetReqBandwidthPeak(v float64) *QueryFlowOverviewLbOutput {
	s.ReqBandwidthPeak = &v
	return s
}

// SetReqCountPeak sets the ReqCountPeak field's value.
func (s *QueryFlowOverviewLbOutput) SetReqCountPeak(v float64) *QueryFlowOverviewLbOutput {
	s.ReqCountPeak = &v
	return s
}

// SetReqQPSPeak sets the ReqQPSPeak field's value.
func (s *QueryFlowOverviewLbOutput) SetReqQPSPeak(v float64) *QueryFlowOverviewLbOutput {
	s.ReqQPSPeak = &v
	return s
}

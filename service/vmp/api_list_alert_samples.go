// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package vmp

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opListAlertSamplesCommon = "ListAlertSamples"

// ListAlertSamplesCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ListAlertSamplesCommon operation. The "output" return
// value will be populated with the ListAlertSamplesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListAlertSamplesCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListAlertSamplesCommon Send returns without error.
//
// See ListAlertSamplesCommon for more information on using the ListAlertSamplesCommon
// API call, and error handling.
//
//    // Example sending a request using the ListAlertSamplesCommonRequest method.
//    req, resp := client.ListAlertSamplesCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VMP) ListAlertSamplesCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opListAlertSamplesCommon,
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

// ListAlertSamplesCommon API operation for VMP.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VMP's
// API operation ListAlertSamplesCommon for usage and error information.
func (c *VMP) ListAlertSamplesCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ListAlertSamplesCommonRequest(input)
	return out, req.Send()
}

// ListAlertSamplesCommonWithContext is the same as ListAlertSamplesCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ListAlertSamplesCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VMP) ListAlertSamplesCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ListAlertSamplesCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opListAlertSamples = "ListAlertSamples"

// ListAlertSamplesRequest generates a "volcengine/request.Request" representing the
// client's request for the ListAlertSamples operation. The "output" return
// value will be populated with the ListAlertSamplesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListAlertSamplesCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListAlertSamplesCommon Send returns without error.
//
// See ListAlertSamples for more information on using the ListAlertSamples
// API call, and error handling.
//
//    // Example sending a request using the ListAlertSamplesRequest method.
//    req, resp := client.ListAlertSamplesRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VMP) ListAlertSamplesRequest(input *ListAlertSamplesInput) (req *request.Request, output *ListAlertSamplesOutput) {
	op := &request.Operation{
		Name:       opListAlertSamples,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListAlertSamplesInput{}
	}

	output = &ListAlertSamplesOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// ListAlertSamples API operation for VMP.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VMP's
// API operation ListAlertSamples for usage and error information.
func (c *VMP) ListAlertSamples(input *ListAlertSamplesInput) (*ListAlertSamplesOutput, error) {
	req, out := c.ListAlertSamplesRequest(input)
	return out, req.Send()
}

// ListAlertSamplesWithContext is the same as ListAlertSamples with the addition of
// the ability to pass a context and additional request options.
//
// See ListAlertSamples for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VMP) ListAlertSamplesWithContext(ctx volcengine.Context, input *ListAlertSamplesInput, opts ...request.Option) (*ListAlertSamplesOutput, error) {
	req, out := c.ListAlertSamplesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type FilterForListAlertSamplesInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	AlertId *string `type:"string" json:",omitempty"`

	SampleSince *int64 `type:"int64" json:",omitempty"`

	SampleUntil *int64 `type:"int64" json:",omitempty"`
}

// String returns the string representation
func (s FilterForListAlertSamplesInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s FilterForListAlertSamplesInput) GoString() string {
	return s.String()
}

// SetAlertId sets the AlertId field's value.
func (s *FilterForListAlertSamplesInput) SetAlertId(v string) *FilterForListAlertSamplesInput {
	s.AlertId = &v
	return s
}

// SetSampleSince sets the SampleSince field's value.
func (s *FilterForListAlertSamplesInput) SetSampleSince(v int64) *FilterForListAlertSamplesInput {
	s.SampleSince = &v
	return s
}

// SetSampleUntil sets the SampleUntil field's value.
func (s *FilterForListAlertSamplesInput) SetSampleUntil(v int64) *FilterForListAlertSamplesInput {
	s.SampleUntil = &v
	return s
}

type ItemForListAlertSamplesOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	AlertId *string `type:"string" json:",omitempty"`

	Level *string `type:"string" json:",omitempty"`

	Phase *string `type:"string" json:",omitempty"`

	Timestamp *int64 `type:"int64" json:",omitempty"`

	Value *float64 `type:"double" json:",omitempty"`
}

// String returns the string representation
func (s ItemForListAlertSamplesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ItemForListAlertSamplesOutput) GoString() string {
	return s.String()
}

// SetAlertId sets the AlertId field's value.
func (s *ItemForListAlertSamplesOutput) SetAlertId(v string) *ItemForListAlertSamplesOutput {
	s.AlertId = &v
	return s
}

// SetLevel sets the Level field's value.
func (s *ItemForListAlertSamplesOutput) SetLevel(v string) *ItemForListAlertSamplesOutput {
	s.Level = &v
	return s
}

// SetPhase sets the Phase field's value.
func (s *ItemForListAlertSamplesOutput) SetPhase(v string) *ItemForListAlertSamplesOutput {
	s.Phase = &v
	return s
}

// SetTimestamp sets the Timestamp field's value.
func (s *ItemForListAlertSamplesOutput) SetTimestamp(v int64) *ItemForListAlertSamplesOutput {
	s.Timestamp = &v
	return s
}

// SetValue sets the Value field's value.
func (s *ItemForListAlertSamplesOutput) SetValue(v float64) *ItemForListAlertSamplesOutput {
	s.Value = &v
	return s
}

type ListAlertSamplesInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Filter *FilterForListAlertSamplesInput `type:"structure" json:",omitempty"`

	Limit *int64 `type:"int64" json:",omitempty"`
}

// String returns the string representation
func (s ListAlertSamplesInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListAlertSamplesInput) GoString() string {
	return s.String()
}

// SetFilter sets the Filter field's value.
func (s *ListAlertSamplesInput) SetFilter(v *FilterForListAlertSamplesInput) *ListAlertSamplesInput {
	s.Filter = v
	return s
}

// SetLimit sets the Limit field's value.
func (s *ListAlertSamplesInput) SetLimit(v int64) *ListAlertSamplesInput {
	s.Limit = &v
	return s
}

type ListAlertSamplesOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	Items []*ItemForListAlertSamplesOutput `type:"list" json:",omitempty"`

	Total *int64 `type:"int64" json:",omitempty"`
}

// String returns the string representation
func (s ListAlertSamplesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListAlertSamplesOutput) GoString() string {
	return s.String()
}

// SetItems sets the Items field's value.
func (s *ListAlertSamplesOutput) SetItems(v []*ItemForListAlertSamplesOutput) *ListAlertSamplesOutput {
	s.Items = v
	return s
}

// SetTotal sets the Total field's value.
func (s *ListAlertSamplesOutput) SetTotal(v int64) *ListAlertSamplesOutput {
	s.Total = &v
	return s
}

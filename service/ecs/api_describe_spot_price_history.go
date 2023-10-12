// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package ecs

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribeSpotPriceHistoryCommon = "DescribeSpotPriceHistory"

// DescribeSpotPriceHistoryCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeSpotPriceHistoryCommon operation. The "output" return
// value will be populated with the DescribeSpotPriceHistoryCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeSpotPriceHistoryCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeSpotPriceHistoryCommon Send returns without error.
//
// See DescribeSpotPriceHistoryCommon for more information on using the DescribeSpotPriceHistoryCommon
// API call, and error handling.
//
//	// Example sending a request using the DescribeSpotPriceHistoryCommonRequest method.
//	req, resp := client.DescribeSpotPriceHistoryCommonRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *ECS) DescribeSpotPriceHistoryCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeSpotPriceHistoryCommon,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &map[string]interface{}{}
	}

	output = &map[string]interface{}{}
	req = c.newRequest(op, input, output)

	return
}

// DescribeSpotPriceHistoryCommon API operation for ECS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ECS's
// API operation DescribeSpotPriceHistoryCommon for usage and error information.
func (c *ECS) DescribeSpotPriceHistoryCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeSpotPriceHistoryCommonRequest(input)
	return out, req.Send()
}

// DescribeSpotPriceHistoryCommonWithContext is the same as DescribeSpotPriceHistoryCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeSpotPriceHistoryCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ECS) DescribeSpotPriceHistoryCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeSpotPriceHistoryCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeSpotPriceHistory = "DescribeSpotPriceHistory"

// DescribeSpotPriceHistoryRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeSpotPriceHistory operation. The "output" return
// value will be populated with the DescribeSpotPriceHistoryCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeSpotPriceHistoryCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeSpotPriceHistoryCommon Send returns without error.
//
// See DescribeSpotPriceHistory for more information on using the DescribeSpotPriceHistory
// API call, and error handling.
//
//	// Example sending a request using the DescribeSpotPriceHistoryRequest method.
//	req, resp := client.DescribeSpotPriceHistoryRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *ECS) DescribeSpotPriceHistoryRequest(input *DescribeSpotPriceHistoryInput) (req *request.Request, output *DescribeSpotPriceHistoryOutput) {
	op := &request.Operation{
		Name:       opDescribeSpotPriceHistory,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeSpotPriceHistoryInput{}
	}

	output = &DescribeSpotPriceHistoryOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DescribeSpotPriceHistory API operation for ECS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ECS's
// API operation DescribeSpotPriceHistory for usage and error information.
func (c *ECS) DescribeSpotPriceHistory(input *DescribeSpotPriceHistoryInput) (*DescribeSpotPriceHistoryOutput, error) {
	req, out := c.DescribeSpotPriceHistoryRequest(input)
	return out, req.Send()
}

// DescribeSpotPriceHistoryWithContext is the same as DescribeSpotPriceHistory with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeSpotPriceHistory for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ECS) DescribeSpotPriceHistoryWithContext(ctx volcengine.Context, input *DescribeSpotPriceHistoryInput, opts ...request.Option) (*DescribeSpotPriceHistoryOutput, error) {
	req, out := c.DescribeSpotPriceHistoryRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DescribeSpotPriceHistoryInput struct {
	_ struct{} `type:"structure"`

	InstanceTypeId *string `type:"string"`

	MaxResults *int32 `type:"int32"`

	NextToken *string `type:"string"`

	TimestampEnd *string `type:"string"`

	TimestampStart *string `type:"string"`

	ZoneId *string `type:"string"`
}

// String returns the string representation
func (s DescribeSpotPriceHistoryInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeSpotPriceHistoryInput) GoString() string {
	return s.String()
}

// SetInstanceTypeId sets the InstanceTypeId field's value.
func (s *DescribeSpotPriceHistoryInput) SetInstanceTypeId(v string) *DescribeSpotPriceHistoryInput {
	s.InstanceTypeId = &v
	return s
}

// SetMaxResults sets the MaxResults field's value.
func (s *DescribeSpotPriceHistoryInput) SetMaxResults(v int32) *DescribeSpotPriceHistoryInput {
	s.MaxResults = &v
	return s
}

// SetNextToken sets the NextToken field's value.
func (s *DescribeSpotPriceHistoryInput) SetNextToken(v string) *DescribeSpotPriceHistoryInput {
	s.NextToken = &v
	return s
}

// SetTimestampEnd sets the TimestampEnd field's value.
func (s *DescribeSpotPriceHistoryInput) SetTimestampEnd(v string) *DescribeSpotPriceHistoryInput {
	s.TimestampEnd = &v
	return s
}

// SetTimestampStart sets the TimestampStart field's value.
func (s *DescribeSpotPriceHistoryInput) SetTimestampStart(v string) *DescribeSpotPriceHistoryInput {
	s.TimestampStart = &v
	return s
}

// SetZoneId sets the ZoneId field's value.
func (s *DescribeSpotPriceHistoryInput) SetZoneId(v string) *DescribeSpotPriceHistoryInput {
	s.ZoneId = &v
	return s
}

type DescribeSpotPriceHistoryOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	NextToken *string `type:"string"`

	SpotPrices []*SpotPriceForDescribeSpotPriceHistoryOutput `type:"list"`
}

// String returns the string representation
func (s DescribeSpotPriceHistoryOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeSpotPriceHistoryOutput) GoString() string {
	return s.String()
}

// SetNextToken sets the NextToken field's value.
func (s *DescribeSpotPriceHistoryOutput) SetNextToken(v string) *DescribeSpotPriceHistoryOutput {
	s.NextToken = &v
	return s
}

// SetSpotPrices sets the SpotPrices field's value.
func (s *DescribeSpotPriceHistoryOutput) SetSpotPrices(v []*SpotPriceForDescribeSpotPriceHistoryOutput) *DescribeSpotPriceHistoryOutput {
	s.SpotPrices = v
	return s
}

type SpotPriceForDescribeSpotPriceHistoryOutput struct {
	_ struct{} `type:"structure"`

	InstanceTypeId *string `type:"string"`

	SpotPrice *float64 `type:"float"`

	Timestamp *string `type:"string"`

	ZoneId *string `type:"string"`
}

// String returns the string representation
func (s SpotPriceForDescribeSpotPriceHistoryOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s SpotPriceForDescribeSpotPriceHistoryOutput) GoString() string {
	return s.String()
}

// SetInstanceTypeId sets the InstanceTypeId field's value.
func (s *SpotPriceForDescribeSpotPriceHistoryOutput) SetInstanceTypeId(v string) *SpotPriceForDescribeSpotPriceHistoryOutput {
	s.InstanceTypeId = &v
	return s
}

// SetSpotPrice sets the SpotPrice field's value.
func (s *SpotPriceForDescribeSpotPriceHistoryOutput) SetSpotPrice(v float64) *SpotPriceForDescribeSpotPriceHistoryOutput {
	s.SpotPrice = &v
	return s
}

// SetTimestamp sets the Timestamp field's value.
func (s *SpotPriceForDescribeSpotPriceHistoryOutput) SetTimestamp(v string) *SpotPriceForDescribeSpotPriceHistoryOutput {
	s.Timestamp = &v
	return s
}

// SetZoneId sets the ZoneId field's value.
func (s *SpotPriceForDescribeSpotPriceHistoryOutput) SetZoneId(v string) *SpotPriceForDescribeSpotPriceHistoryOutput {
	s.ZoneId = &v
	return s
}
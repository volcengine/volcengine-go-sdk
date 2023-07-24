// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package kafka

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribeConsumedPartitionsCommon = "DescribeConsumedPartitions"

// DescribeConsumedPartitionsCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeConsumedPartitionsCommon operation. The "output" return
// value will be populated with the DescribeConsumedPartitionsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeConsumedPartitionsCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeConsumedPartitionsCommon Send returns without error.
//
// See DescribeConsumedPartitionsCommon for more information on using the DescribeConsumedPartitionsCommon
// API call, and error handling.
//
//	// Example sending a request using the DescribeConsumedPartitionsCommonRequest method.
//	req, resp := client.DescribeConsumedPartitionsCommonRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *KAFKA) DescribeConsumedPartitionsCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeConsumedPartitionsCommon,
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

// DescribeConsumedPartitionsCommon API operation for KAFKA.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for KAFKA's
// API operation DescribeConsumedPartitionsCommon for usage and error information.
func (c *KAFKA) DescribeConsumedPartitionsCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeConsumedPartitionsCommonRequest(input)
	return out, req.Send()
}

// DescribeConsumedPartitionsCommonWithContext is the same as DescribeConsumedPartitionsCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeConsumedPartitionsCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *KAFKA) DescribeConsumedPartitionsCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeConsumedPartitionsCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeConsumedPartitions = "DescribeConsumedPartitions"

// DescribeConsumedPartitionsRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeConsumedPartitions operation. The "output" return
// value will be populated with the DescribeConsumedPartitionsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeConsumedPartitionsCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeConsumedPartitionsCommon Send returns without error.
//
// See DescribeConsumedPartitions for more information on using the DescribeConsumedPartitions
// API call, and error handling.
//
//	// Example sending a request using the DescribeConsumedPartitionsRequest method.
//	req, resp := client.DescribeConsumedPartitionsRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *KAFKA) DescribeConsumedPartitionsRequest(input *DescribeConsumedPartitionsInput) (req *request.Request, output *DescribeConsumedPartitionsOutput) {
	op := &request.Operation{
		Name:       opDescribeConsumedPartitions,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeConsumedPartitionsInput{}
	}

	output = &DescribeConsumedPartitionsOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// DescribeConsumedPartitions API operation for KAFKA.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for KAFKA's
// API operation DescribeConsumedPartitions for usage and error information.
func (c *KAFKA) DescribeConsumedPartitions(input *DescribeConsumedPartitionsInput) (*DescribeConsumedPartitionsOutput, error) {
	req, out := c.DescribeConsumedPartitionsRequest(input)
	return out, req.Send()
}

// DescribeConsumedPartitionsWithContext is the same as DescribeConsumedPartitions with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeConsumedPartitions for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *KAFKA) DescribeConsumedPartitionsWithContext(ctx volcengine.Context, input *DescribeConsumedPartitionsInput, opts ...request.Option) (*DescribeConsumedPartitionsOutput, error) {
	req, out := c.DescribeConsumedPartitionsRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ConsumedPartitionsInfoForDescribeConsumedPartitionsOutput struct {
	_ struct{} `type:"structure"`

	Accumulation *int64 `type:"int64"`

	ConsumedClient *string `type:"string"`

	ConsumedOffset *int64 `type:"int64"`

	EndOffset *int64 `type:"int64"`

	PartitionId *int32 `type:"int32"`

	StartOffset *int64 `type:"int64"`
}

// String returns the string representation
func (s ConsumedPartitionsInfoForDescribeConsumedPartitionsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ConsumedPartitionsInfoForDescribeConsumedPartitionsOutput) GoString() string {
	return s.String()
}

// SetAccumulation sets the Accumulation field's value.
func (s *ConsumedPartitionsInfoForDescribeConsumedPartitionsOutput) SetAccumulation(v int64) *ConsumedPartitionsInfoForDescribeConsumedPartitionsOutput {
	s.Accumulation = &v
	return s
}

// SetConsumedClient sets the ConsumedClient field's value.
func (s *ConsumedPartitionsInfoForDescribeConsumedPartitionsOutput) SetConsumedClient(v string) *ConsumedPartitionsInfoForDescribeConsumedPartitionsOutput {
	s.ConsumedClient = &v
	return s
}

// SetConsumedOffset sets the ConsumedOffset field's value.
func (s *ConsumedPartitionsInfoForDescribeConsumedPartitionsOutput) SetConsumedOffset(v int64) *ConsumedPartitionsInfoForDescribeConsumedPartitionsOutput {
	s.ConsumedOffset = &v
	return s
}

// SetEndOffset sets the EndOffset field's value.
func (s *ConsumedPartitionsInfoForDescribeConsumedPartitionsOutput) SetEndOffset(v int64) *ConsumedPartitionsInfoForDescribeConsumedPartitionsOutput {
	s.EndOffset = &v
	return s
}

// SetPartitionId sets the PartitionId field's value.
func (s *ConsumedPartitionsInfoForDescribeConsumedPartitionsOutput) SetPartitionId(v int32) *ConsumedPartitionsInfoForDescribeConsumedPartitionsOutput {
	s.PartitionId = &v
	return s
}

// SetStartOffset sets the StartOffset field's value.
func (s *ConsumedPartitionsInfoForDescribeConsumedPartitionsOutput) SetStartOffset(v int64) *ConsumedPartitionsInfoForDescribeConsumedPartitionsOutput {
	s.StartOffset = &v
	return s
}

type DescribeConsumedPartitionsInput struct {
	_ struct{} `type:"structure"`

	GroupId *string `type:"string"`

	InstanceId *string `type:"string"`

	PageNumber *int32 `type:"int32"`

	PageSize *int32 `type:"int32"`

	TopicName *string `type:"string"`
}

// String returns the string representation
func (s DescribeConsumedPartitionsInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeConsumedPartitionsInput) GoString() string {
	return s.String()
}

// SetGroupId sets the GroupId field's value.
func (s *DescribeConsumedPartitionsInput) SetGroupId(v string) *DescribeConsumedPartitionsInput {
	s.GroupId = &v
	return s
}

// SetInstanceId sets the InstanceId field's value.
func (s *DescribeConsumedPartitionsInput) SetInstanceId(v string) *DescribeConsumedPartitionsInput {
	s.InstanceId = &v
	return s
}

// SetPageNumber sets the PageNumber field's value.
func (s *DescribeConsumedPartitionsInput) SetPageNumber(v int32) *DescribeConsumedPartitionsInput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *DescribeConsumedPartitionsInput) SetPageSize(v int32) *DescribeConsumedPartitionsInput {
	s.PageSize = &v
	return s
}

// SetTopicName sets the TopicName field's value.
func (s *DescribeConsumedPartitionsInput) SetTopicName(v string) *DescribeConsumedPartitionsInput {
	s.TopicName = &v
	return s
}

type DescribeConsumedPartitionsOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	ConsumedPartitionsInfo []*ConsumedPartitionsInfoForDescribeConsumedPartitionsOutput `type:"list"`

	Total *int32 `type:"int32"`
}

// String returns the string representation
func (s DescribeConsumedPartitionsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeConsumedPartitionsOutput) GoString() string {
	return s.String()
}

// SetConsumedPartitionsInfo sets the ConsumedPartitionsInfo field's value.
func (s *DescribeConsumedPartitionsOutput) SetConsumedPartitionsInfo(v []*ConsumedPartitionsInfoForDescribeConsumedPartitionsOutput) *DescribeConsumedPartitionsOutput {
	s.ConsumedPartitionsInfo = v
	return s
}

// SetTotal sets the Total field's value.
func (s *DescribeConsumedPartitionsOutput) SetTotal(v int32) *DescribeConsumedPartitionsOutput {
	s.Total = &v
	return s
}

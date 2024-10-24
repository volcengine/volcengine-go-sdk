// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package rocketmq

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribeConsumedTopicDetailCommon = "DescribeConsumedTopicDetail"

// DescribeConsumedTopicDetailCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeConsumedTopicDetailCommon operation. The "output" return
// value will be populated with the DescribeConsumedTopicDetailCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeConsumedTopicDetailCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeConsumedTopicDetailCommon Send returns without error.
//
// See DescribeConsumedTopicDetailCommon for more information on using the DescribeConsumedTopicDetailCommon
// API call, and error handling.
//
//    // Example sending a request using the DescribeConsumedTopicDetailCommonRequest method.
//    req, resp := client.DescribeConsumedTopicDetailCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ROCKETMQ) DescribeConsumedTopicDetailCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeConsumedTopicDetailCommon,
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

// DescribeConsumedTopicDetailCommon API operation for ROCKETMQ.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ROCKETMQ's
// API operation DescribeConsumedTopicDetailCommon for usage and error information.
func (c *ROCKETMQ) DescribeConsumedTopicDetailCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeConsumedTopicDetailCommonRequest(input)
	return out, req.Send()
}

// DescribeConsumedTopicDetailCommonWithContext is the same as DescribeConsumedTopicDetailCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeConsumedTopicDetailCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ROCKETMQ) DescribeConsumedTopicDetailCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeConsumedTopicDetailCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeConsumedTopicDetail = "DescribeConsumedTopicDetail"

// DescribeConsumedTopicDetailRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeConsumedTopicDetail operation. The "output" return
// value will be populated with the DescribeConsumedTopicDetailCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeConsumedTopicDetailCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeConsumedTopicDetailCommon Send returns without error.
//
// See DescribeConsumedTopicDetail for more information on using the DescribeConsumedTopicDetail
// API call, and error handling.
//
//    // Example sending a request using the DescribeConsumedTopicDetailRequest method.
//    req, resp := client.DescribeConsumedTopicDetailRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ROCKETMQ) DescribeConsumedTopicDetailRequest(input *DescribeConsumedTopicDetailInput) (req *request.Request, output *DescribeConsumedTopicDetailOutput) {
	op := &request.Operation{
		Name:       opDescribeConsumedTopicDetail,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeConsumedTopicDetailInput{}
	}

	output = &DescribeConsumedTopicDetailOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// DescribeConsumedTopicDetail API operation for ROCKETMQ.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ROCKETMQ's
// API operation DescribeConsumedTopicDetail for usage and error information.
func (c *ROCKETMQ) DescribeConsumedTopicDetail(input *DescribeConsumedTopicDetailInput) (*DescribeConsumedTopicDetailOutput, error) {
	req, out := c.DescribeConsumedTopicDetailRequest(input)
	return out, req.Send()
}

// DescribeConsumedTopicDetailWithContext is the same as DescribeConsumedTopicDetail with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeConsumedTopicDetail for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ROCKETMQ) DescribeConsumedTopicDetailWithContext(ctx volcengine.Context, input *DescribeConsumedTopicDetailInput, opts ...request.Option) (*DescribeConsumedTopicDetailOutput, error) {
	req, out := c.DescribeConsumedTopicDetailRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ConsumedQueueInfoForDescribeConsumedTopicDetailOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	ConsumeOffset *int32 `type:"int32" json:",omitempty"`

	DiffCount *int32 `type:"int32" json:",omitempty"`

	EndOffset *int32 `type:"int32" json:",omitempty"`

	LastTimeStamp *int64 `type:"int64" json:",omitempty"`

	QueueId *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s ConsumedQueueInfoForDescribeConsumedTopicDetailOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ConsumedQueueInfoForDescribeConsumedTopicDetailOutput) GoString() string {
	return s.String()
}

// SetConsumeOffset sets the ConsumeOffset field's value.
func (s *ConsumedQueueInfoForDescribeConsumedTopicDetailOutput) SetConsumeOffset(v int32) *ConsumedQueueInfoForDescribeConsumedTopicDetailOutput {
	s.ConsumeOffset = &v
	return s
}

// SetDiffCount sets the DiffCount field's value.
func (s *ConsumedQueueInfoForDescribeConsumedTopicDetailOutput) SetDiffCount(v int32) *ConsumedQueueInfoForDescribeConsumedTopicDetailOutput {
	s.DiffCount = &v
	return s
}

// SetEndOffset sets the EndOffset field's value.
func (s *ConsumedQueueInfoForDescribeConsumedTopicDetailOutput) SetEndOffset(v int32) *ConsumedQueueInfoForDescribeConsumedTopicDetailOutput {
	s.EndOffset = &v
	return s
}

// SetLastTimeStamp sets the LastTimeStamp field's value.
func (s *ConsumedQueueInfoForDescribeConsumedTopicDetailOutput) SetLastTimeStamp(v int64) *ConsumedQueueInfoForDescribeConsumedTopicDetailOutput {
	s.LastTimeStamp = &v
	return s
}

// SetQueueId sets the QueueId field's value.
func (s *ConsumedQueueInfoForDescribeConsumedTopicDetailOutput) SetQueueId(v string) *ConsumedQueueInfoForDescribeConsumedTopicDetailOutput {
	s.QueueId = &v
	return s
}

type DescribeConsumedTopicDetailInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	// GroupId is a required field
	GroupId *string `type:"string" json:",omitempty" required:"true"`

	// InstanceId is a required field
	InstanceId *string `type:"string" json:",omitempty" required:"true"`

	// PageNumber is a required field
	PageNumber *int32 `type:"int32" json:",omitempty" required:"true"`

	// PageSize is a required field
	PageSize *int32 `type:"int32" json:",omitempty" required:"true"`

	// TopicName is a required field
	TopicName *string `type:"string" json:",omitempty" required:"true"`
}

// String returns the string representation
func (s DescribeConsumedTopicDetailInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeConsumedTopicDetailInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeConsumedTopicDetailInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DescribeConsumedTopicDetailInput"}
	if s.GroupId == nil {
		invalidParams.Add(request.NewErrParamRequired("GroupId"))
	}
	if s.InstanceId == nil {
		invalidParams.Add(request.NewErrParamRequired("InstanceId"))
	}
	if s.PageNumber == nil {
		invalidParams.Add(request.NewErrParamRequired("PageNumber"))
	}
	if s.PageSize == nil {
		invalidParams.Add(request.NewErrParamRequired("PageSize"))
	}
	if s.TopicName == nil {
		invalidParams.Add(request.NewErrParamRequired("TopicName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetGroupId sets the GroupId field's value.
func (s *DescribeConsumedTopicDetailInput) SetGroupId(v string) *DescribeConsumedTopicDetailInput {
	s.GroupId = &v
	return s
}

// SetInstanceId sets the InstanceId field's value.
func (s *DescribeConsumedTopicDetailInput) SetInstanceId(v string) *DescribeConsumedTopicDetailInput {
	s.InstanceId = &v
	return s
}

// SetPageNumber sets the PageNumber field's value.
func (s *DescribeConsumedTopicDetailInput) SetPageNumber(v int32) *DescribeConsumedTopicDetailInput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *DescribeConsumedTopicDetailInput) SetPageSize(v int32) *DescribeConsumedTopicDetailInput {
	s.PageSize = &v
	return s
}

// SetTopicName sets the TopicName field's value.
func (s *DescribeConsumedTopicDetailInput) SetTopicName(v string) *DescribeConsumedTopicDetailInput {
	s.TopicName = &v
	return s
}

type DescribeConsumedTopicDetailOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	ConsumedQueueInfo []*ConsumedQueueInfoForDescribeConsumedTopicDetailOutput `type:"list" json:",omitempty"`

	Total *int32 `type:"int32" json:",omitempty"`
}

// String returns the string representation
func (s DescribeConsumedTopicDetailOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeConsumedTopicDetailOutput) GoString() string {
	return s.String()
}

// SetConsumedQueueInfo sets the ConsumedQueueInfo field's value.
func (s *DescribeConsumedTopicDetailOutput) SetConsumedQueueInfo(v []*ConsumedQueueInfoForDescribeConsumedTopicDetailOutput) *DescribeConsumedTopicDetailOutput {
	s.ConsumedQueueInfo = v
	return s
}

// SetTotal sets the Total field's value.
func (s *DescribeConsumedTopicDetailOutput) SetTotal(v int32) *DescribeConsumedTopicDetailOutput {
	s.Total = &v
	return s
}
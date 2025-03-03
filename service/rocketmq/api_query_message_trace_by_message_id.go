// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package rocketmq

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opQueryMessageTraceByMessageIdCommon = "QueryMessageTraceByMessageId"

// QueryMessageTraceByMessageIdCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the QueryMessageTraceByMessageIdCommon operation. The "output" return
// value will be populated with the QueryMessageTraceByMessageIdCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned QueryMessageTraceByMessageIdCommon Request to send the API call to the service.
// the "output" return value is not valid until after QueryMessageTraceByMessageIdCommon Send returns without error.
//
// See QueryMessageTraceByMessageIdCommon for more information on using the QueryMessageTraceByMessageIdCommon
// API call, and error handling.
//
//    // Example sending a request using the QueryMessageTraceByMessageIdCommonRequest method.
//    req, resp := client.QueryMessageTraceByMessageIdCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ROCKETMQ) QueryMessageTraceByMessageIdCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opQueryMessageTraceByMessageIdCommon,
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

// QueryMessageTraceByMessageIdCommon API operation for ROCKETMQ.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ROCKETMQ's
// API operation QueryMessageTraceByMessageIdCommon for usage and error information.
func (c *ROCKETMQ) QueryMessageTraceByMessageIdCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.QueryMessageTraceByMessageIdCommonRequest(input)
	return out, req.Send()
}

// QueryMessageTraceByMessageIdCommonWithContext is the same as QueryMessageTraceByMessageIdCommon with the addition of
// the ability to pass a context and additional request options.
//
// See QueryMessageTraceByMessageIdCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ROCKETMQ) QueryMessageTraceByMessageIdCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.QueryMessageTraceByMessageIdCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opQueryMessageTraceByMessageId = "QueryMessageTraceByMessageId"

// QueryMessageTraceByMessageIdRequest generates a "volcengine/request.Request" representing the
// client's request for the QueryMessageTraceByMessageId operation. The "output" return
// value will be populated with the QueryMessageTraceByMessageIdCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned QueryMessageTraceByMessageIdCommon Request to send the API call to the service.
// the "output" return value is not valid until after QueryMessageTraceByMessageIdCommon Send returns without error.
//
// See QueryMessageTraceByMessageId for more information on using the QueryMessageTraceByMessageId
// API call, and error handling.
//
//    // Example sending a request using the QueryMessageTraceByMessageIdRequest method.
//    req, resp := client.QueryMessageTraceByMessageIdRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ROCKETMQ) QueryMessageTraceByMessageIdRequest(input *QueryMessageTraceByMessageIdInput) (req *request.Request, output *QueryMessageTraceByMessageIdOutput) {
	op := &request.Operation{
		Name:       opQueryMessageTraceByMessageId,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &QueryMessageTraceByMessageIdInput{}
	}

	output = &QueryMessageTraceByMessageIdOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// QueryMessageTraceByMessageId API operation for ROCKETMQ.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ROCKETMQ's
// API operation QueryMessageTraceByMessageId for usage and error information.
func (c *ROCKETMQ) QueryMessageTraceByMessageId(input *QueryMessageTraceByMessageIdInput) (*QueryMessageTraceByMessageIdOutput, error) {
	req, out := c.QueryMessageTraceByMessageIdRequest(input)
	return out, req.Send()
}

// QueryMessageTraceByMessageIdWithContext is the same as QueryMessageTraceByMessageId with the addition of
// the ability to pass a context and additional request options.
//
// See QueryMessageTraceByMessageId for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ROCKETMQ) QueryMessageTraceByMessageIdWithContext(ctx volcengine.Context, input *QueryMessageTraceByMessageIdInput, opts ...request.Option) (*QueryMessageTraceByMessageIdOutput, error) {
	req, out := c.QueryMessageTraceByMessageIdRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ConsumerTraceInfoForQueryMessageTraceByMessageIdOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	ConsumeStatus *string `type:"string" json:",omitempty"`

	ConsumeSuccess *bool `type:"boolean" json:",omitempty"`

	ConsumerHost *string `type:"string" json:",omitempty"`

	EndProcessTimestamp *int64 `type:"int64" json:",omitempty"`

	GroupName *string `type:"string" json:",omitempty"`

	ProcessCostTimeMs *int32 `type:"int32" json:",omitempty"`

	ReconsumeTimes *int32 `type:"int32" json:",omitempty"`

	StartProcessTimestamp *int64 `type:"int64" json:",omitempty"`
}

// String returns the string representation
func (s ConsumerTraceInfoForQueryMessageTraceByMessageIdOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ConsumerTraceInfoForQueryMessageTraceByMessageIdOutput) GoString() string {
	return s.String()
}

// SetConsumeStatus sets the ConsumeStatus field's value.
func (s *ConsumerTraceInfoForQueryMessageTraceByMessageIdOutput) SetConsumeStatus(v string) *ConsumerTraceInfoForQueryMessageTraceByMessageIdOutput {
	s.ConsumeStatus = &v
	return s
}

// SetConsumeSuccess sets the ConsumeSuccess field's value.
func (s *ConsumerTraceInfoForQueryMessageTraceByMessageIdOutput) SetConsumeSuccess(v bool) *ConsumerTraceInfoForQueryMessageTraceByMessageIdOutput {
	s.ConsumeSuccess = &v
	return s
}

// SetConsumerHost sets the ConsumerHost field's value.
func (s *ConsumerTraceInfoForQueryMessageTraceByMessageIdOutput) SetConsumerHost(v string) *ConsumerTraceInfoForQueryMessageTraceByMessageIdOutput {
	s.ConsumerHost = &v
	return s
}

// SetEndProcessTimestamp sets the EndProcessTimestamp field's value.
func (s *ConsumerTraceInfoForQueryMessageTraceByMessageIdOutput) SetEndProcessTimestamp(v int64) *ConsumerTraceInfoForQueryMessageTraceByMessageIdOutput {
	s.EndProcessTimestamp = &v
	return s
}

// SetGroupName sets the GroupName field's value.
func (s *ConsumerTraceInfoForQueryMessageTraceByMessageIdOutput) SetGroupName(v string) *ConsumerTraceInfoForQueryMessageTraceByMessageIdOutput {
	s.GroupName = &v
	return s
}

// SetProcessCostTimeMs sets the ProcessCostTimeMs field's value.
func (s *ConsumerTraceInfoForQueryMessageTraceByMessageIdOutput) SetProcessCostTimeMs(v int32) *ConsumerTraceInfoForQueryMessageTraceByMessageIdOutput {
	s.ProcessCostTimeMs = &v
	return s
}

// SetReconsumeTimes sets the ReconsumeTimes field's value.
func (s *ConsumerTraceInfoForQueryMessageTraceByMessageIdOutput) SetReconsumeTimes(v int32) *ConsumerTraceInfoForQueryMessageTraceByMessageIdOutput {
	s.ReconsumeTimes = &v
	return s
}

// SetStartProcessTimestamp sets the StartProcessTimestamp field's value.
func (s *ConsumerTraceInfoForQueryMessageTraceByMessageIdOutput) SetStartProcessTimestamp(v int64) *ConsumerTraceInfoForQueryMessageTraceByMessageIdOutput {
	s.StartProcessTimestamp = &v
	return s
}

type ProducerTraceInfoForQueryMessageTraceByMessageIdOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	MessageBornHost *string `type:"string" json:",omitempty"`

	MessageBornTimestamp *int64 `type:"int64" json:",omitempty"`

	SendCostTimeMs *int32 `type:"int32" json:",omitempty"`

	SendStatus *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s ProducerTraceInfoForQueryMessageTraceByMessageIdOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ProducerTraceInfoForQueryMessageTraceByMessageIdOutput) GoString() string {
	return s.String()
}

// SetMessageBornHost sets the MessageBornHost field's value.
func (s *ProducerTraceInfoForQueryMessageTraceByMessageIdOutput) SetMessageBornHost(v string) *ProducerTraceInfoForQueryMessageTraceByMessageIdOutput {
	s.MessageBornHost = &v
	return s
}

// SetMessageBornTimestamp sets the MessageBornTimestamp field's value.
func (s *ProducerTraceInfoForQueryMessageTraceByMessageIdOutput) SetMessageBornTimestamp(v int64) *ProducerTraceInfoForQueryMessageTraceByMessageIdOutput {
	s.MessageBornTimestamp = &v
	return s
}

// SetSendCostTimeMs sets the SendCostTimeMs field's value.
func (s *ProducerTraceInfoForQueryMessageTraceByMessageIdOutput) SetSendCostTimeMs(v int32) *ProducerTraceInfoForQueryMessageTraceByMessageIdOutput {
	s.SendCostTimeMs = &v
	return s
}

// SetSendStatus sets the SendStatus field's value.
func (s *ProducerTraceInfoForQueryMessageTraceByMessageIdOutput) SetSendStatus(v string) *ProducerTraceInfoForQueryMessageTraceByMessageIdOutput {
	s.SendStatus = &v
	return s
}

type QueryMessageTraceByMessageIdInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	// InstanceId is a required field
	InstanceId *string `type:"string" json:",omitempty" required:"true"`

	MessageBornTimestamp *int64 `type:"int64" json:",omitempty"`

	// MessageId is a required field
	MessageId *string `type:"string" json:",omitempty" required:"true"`

	// NeedProduceTraceInfo is a required field
	NeedProduceTraceInfo *bool `type:"boolean" json:",omitempty" required:"true"`

	// QueryEndTimestamp is a required field
	QueryEndTimestamp *int64 `type:"int64" json:",omitempty" required:"true"`

	// QueryStartTimestamp is a required field
	QueryStartTimestamp *int64 `type:"int64" json:",omitempty" required:"true"`

	// Topic is a required field
	Topic *string `type:"string" json:",omitempty" required:"true"`
}

// String returns the string representation
func (s QueryMessageTraceByMessageIdInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s QueryMessageTraceByMessageIdInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *QueryMessageTraceByMessageIdInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "QueryMessageTraceByMessageIdInput"}
	if s.InstanceId == nil {
		invalidParams.Add(request.NewErrParamRequired("InstanceId"))
	}
	if s.MessageId == nil {
		invalidParams.Add(request.NewErrParamRequired("MessageId"))
	}
	if s.NeedProduceTraceInfo == nil {
		invalidParams.Add(request.NewErrParamRequired("NeedProduceTraceInfo"))
	}
	if s.QueryEndTimestamp == nil {
		invalidParams.Add(request.NewErrParamRequired("QueryEndTimestamp"))
	}
	if s.QueryStartTimestamp == nil {
		invalidParams.Add(request.NewErrParamRequired("QueryStartTimestamp"))
	}
	if s.Topic == nil {
		invalidParams.Add(request.NewErrParamRequired("Topic"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetInstanceId sets the InstanceId field's value.
func (s *QueryMessageTraceByMessageIdInput) SetInstanceId(v string) *QueryMessageTraceByMessageIdInput {
	s.InstanceId = &v
	return s
}

// SetMessageBornTimestamp sets the MessageBornTimestamp field's value.
func (s *QueryMessageTraceByMessageIdInput) SetMessageBornTimestamp(v int64) *QueryMessageTraceByMessageIdInput {
	s.MessageBornTimestamp = &v
	return s
}

// SetMessageId sets the MessageId field's value.
func (s *QueryMessageTraceByMessageIdInput) SetMessageId(v string) *QueryMessageTraceByMessageIdInput {
	s.MessageId = &v
	return s
}

// SetNeedProduceTraceInfo sets the NeedProduceTraceInfo field's value.
func (s *QueryMessageTraceByMessageIdInput) SetNeedProduceTraceInfo(v bool) *QueryMessageTraceByMessageIdInput {
	s.NeedProduceTraceInfo = &v
	return s
}

// SetQueryEndTimestamp sets the QueryEndTimestamp field's value.
func (s *QueryMessageTraceByMessageIdInput) SetQueryEndTimestamp(v int64) *QueryMessageTraceByMessageIdInput {
	s.QueryEndTimestamp = &v
	return s
}

// SetQueryStartTimestamp sets the QueryStartTimestamp field's value.
func (s *QueryMessageTraceByMessageIdInput) SetQueryStartTimestamp(v int64) *QueryMessageTraceByMessageIdInput {
	s.QueryStartTimestamp = &v
	return s
}

// SetTopic sets the Topic field's value.
func (s *QueryMessageTraceByMessageIdInput) SetTopic(v string) *QueryMessageTraceByMessageIdInput {
	s.Topic = &v
	return s
}

type QueryMessageTraceByMessageIdOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	ConsumerTraceInfos []*ConsumerTraceInfoForQueryMessageTraceByMessageIdOutput `type:"list" json:",omitempty"`

	ProducerTraceInfo []*ProducerTraceInfoForQueryMessageTraceByMessageIdOutput `type:"list" json:",omitempty"`
}

// String returns the string representation
func (s QueryMessageTraceByMessageIdOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s QueryMessageTraceByMessageIdOutput) GoString() string {
	return s.String()
}

// SetConsumerTraceInfos sets the ConsumerTraceInfos field's value.
func (s *QueryMessageTraceByMessageIdOutput) SetConsumerTraceInfos(v []*ConsumerTraceInfoForQueryMessageTraceByMessageIdOutput) *QueryMessageTraceByMessageIdOutput {
	s.ConsumerTraceInfos = v
	return s
}

// SetProducerTraceInfo sets the ProducerTraceInfo field's value.
func (s *QueryMessageTraceByMessageIdOutput) SetProducerTraceInfo(v []*ProducerTraceInfoForQueryMessageTraceByMessageIdOutput) *QueryMessageTraceByMessageIdOutput {
	s.ProducerTraceInfo = v
	return s
}

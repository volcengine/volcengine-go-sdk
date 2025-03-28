// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package rocketmq

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opQueryMessageByMsgKeyCommon = "QueryMessageByMsgKey"

// QueryMessageByMsgKeyCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the QueryMessageByMsgKeyCommon operation. The "output" return
// value will be populated with the QueryMessageByMsgKeyCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned QueryMessageByMsgKeyCommon Request to send the API call to the service.
// the "output" return value is not valid until after QueryMessageByMsgKeyCommon Send returns without error.
//
// See QueryMessageByMsgKeyCommon for more information on using the QueryMessageByMsgKeyCommon
// API call, and error handling.
//
//    // Example sending a request using the QueryMessageByMsgKeyCommonRequest method.
//    req, resp := client.QueryMessageByMsgKeyCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ROCKETMQ) QueryMessageByMsgKeyCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opQueryMessageByMsgKeyCommon,
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

// QueryMessageByMsgKeyCommon API operation for ROCKETMQ.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ROCKETMQ's
// API operation QueryMessageByMsgKeyCommon for usage and error information.
func (c *ROCKETMQ) QueryMessageByMsgKeyCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.QueryMessageByMsgKeyCommonRequest(input)
	return out, req.Send()
}

// QueryMessageByMsgKeyCommonWithContext is the same as QueryMessageByMsgKeyCommon with the addition of
// the ability to pass a context and additional request options.
//
// See QueryMessageByMsgKeyCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ROCKETMQ) QueryMessageByMsgKeyCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.QueryMessageByMsgKeyCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opQueryMessageByMsgKey = "QueryMessageByMsgKey"

// QueryMessageByMsgKeyRequest generates a "volcengine/request.Request" representing the
// client's request for the QueryMessageByMsgKey operation. The "output" return
// value will be populated with the QueryMessageByMsgKeyCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned QueryMessageByMsgKeyCommon Request to send the API call to the service.
// the "output" return value is not valid until after QueryMessageByMsgKeyCommon Send returns without error.
//
// See QueryMessageByMsgKey for more information on using the QueryMessageByMsgKey
// API call, and error handling.
//
//    // Example sending a request using the QueryMessageByMsgKeyRequest method.
//    req, resp := client.QueryMessageByMsgKeyRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ROCKETMQ) QueryMessageByMsgKeyRequest(input *QueryMessageByMsgKeyInput) (req *request.Request, output *QueryMessageByMsgKeyOutput) {
	op := &request.Operation{
		Name:       opQueryMessageByMsgKey,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &QueryMessageByMsgKeyInput{}
	}

	output = &QueryMessageByMsgKeyOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// QueryMessageByMsgKey API operation for ROCKETMQ.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ROCKETMQ's
// API operation QueryMessageByMsgKey for usage and error information.
func (c *ROCKETMQ) QueryMessageByMsgKey(input *QueryMessageByMsgKeyInput) (*QueryMessageByMsgKeyOutput, error) {
	req, out := c.QueryMessageByMsgKeyRequest(input)
	return out, req.Send()
}

// QueryMessageByMsgKeyWithContext is the same as QueryMessageByMsgKey with the addition of
// the ability to pass a context and additional request options.
//
// See QueryMessageByMsgKey for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ROCKETMQ) QueryMessageByMsgKeyWithContext(ctx volcengine.Context, input *QueryMessageByMsgKeyInput, opts ...request.Option) (*QueryMessageByMsgKeyOutput, error) {
	req, out := c.QueryMessageByMsgKeyRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type MessageListForQueryMessageByMsgKeyOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Body *string `type:"string" json:",omitempty"`

	CreateTimestamp *int64 `type:"int64" json:",omitempty"`

	IsExist *bool `type:"boolean" json:",omitempty"`

	MessageId *string `type:"string" json:",omitempty"`

	MessageKey *string `type:"string" json:",omitempty"`

	MessageSize *int32 `type:"int32" json:",omitempty"`

	ProducerHost *string `type:"string" json:",omitempty"`

	StoreTimestamp *int64 `type:"int64" json:",omitempty"`

	Tag *string `type:"string" json:",omitempty"`

	TopicName *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s MessageListForQueryMessageByMsgKeyOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s MessageListForQueryMessageByMsgKeyOutput) GoString() string {
	return s.String()
}

// SetBody sets the Body field's value.
func (s *MessageListForQueryMessageByMsgKeyOutput) SetBody(v string) *MessageListForQueryMessageByMsgKeyOutput {
	s.Body = &v
	return s
}

// SetCreateTimestamp sets the CreateTimestamp field's value.
func (s *MessageListForQueryMessageByMsgKeyOutput) SetCreateTimestamp(v int64) *MessageListForQueryMessageByMsgKeyOutput {
	s.CreateTimestamp = &v
	return s
}

// SetIsExist sets the IsExist field's value.
func (s *MessageListForQueryMessageByMsgKeyOutput) SetIsExist(v bool) *MessageListForQueryMessageByMsgKeyOutput {
	s.IsExist = &v
	return s
}

// SetMessageId sets the MessageId field's value.
func (s *MessageListForQueryMessageByMsgKeyOutput) SetMessageId(v string) *MessageListForQueryMessageByMsgKeyOutput {
	s.MessageId = &v
	return s
}

// SetMessageKey sets the MessageKey field's value.
func (s *MessageListForQueryMessageByMsgKeyOutput) SetMessageKey(v string) *MessageListForQueryMessageByMsgKeyOutput {
	s.MessageKey = &v
	return s
}

// SetMessageSize sets the MessageSize field's value.
func (s *MessageListForQueryMessageByMsgKeyOutput) SetMessageSize(v int32) *MessageListForQueryMessageByMsgKeyOutput {
	s.MessageSize = &v
	return s
}

// SetProducerHost sets the ProducerHost field's value.
func (s *MessageListForQueryMessageByMsgKeyOutput) SetProducerHost(v string) *MessageListForQueryMessageByMsgKeyOutput {
	s.ProducerHost = &v
	return s
}

// SetStoreTimestamp sets the StoreTimestamp field's value.
func (s *MessageListForQueryMessageByMsgKeyOutput) SetStoreTimestamp(v int64) *MessageListForQueryMessageByMsgKeyOutput {
	s.StoreTimestamp = &v
	return s
}

// SetTag sets the Tag field's value.
func (s *MessageListForQueryMessageByMsgKeyOutput) SetTag(v string) *MessageListForQueryMessageByMsgKeyOutput {
	s.Tag = &v
	return s
}

// SetTopicName sets the TopicName field's value.
func (s *MessageListForQueryMessageByMsgKeyOutput) SetTopicName(v string) *MessageListForQueryMessageByMsgKeyOutput {
	s.TopicName = &v
	return s
}

type QueryMessageByMsgKeyInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	// InstanceId is a required field
	InstanceId *string `type:"string" json:",omitempty" required:"true"`

	// MessageKey is a required field
	MessageKey *string `type:"string" json:",omitempty" required:"true"`

	// TopicName is a required field
	TopicName *string `type:"string" json:",omitempty" required:"true"`
}

// String returns the string representation
func (s QueryMessageByMsgKeyInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s QueryMessageByMsgKeyInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *QueryMessageByMsgKeyInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "QueryMessageByMsgKeyInput"}
	if s.InstanceId == nil {
		invalidParams.Add(request.NewErrParamRequired("InstanceId"))
	}
	if s.MessageKey == nil {
		invalidParams.Add(request.NewErrParamRequired("MessageKey"))
	}
	if s.TopicName == nil {
		invalidParams.Add(request.NewErrParamRequired("TopicName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetInstanceId sets the InstanceId field's value.
func (s *QueryMessageByMsgKeyInput) SetInstanceId(v string) *QueryMessageByMsgKeyInput {
	s.InstanceId = &v
	return s
}

// SetMessageKey sets the MessageKey field's value.
func (s *QueryMessageByMsgKeyInput) SetMessageKey(v string) *QueryMessageByMsgKeyInput {
	s.MessageKey = &v
	return s
}

// SetTopicName sets the TopicName field's value.
func (s *QueryMessageByMsgKeyInput) SetTopicName(v string) *QueryMessageByMsgKeyInput {
	s.TopicName = &v
	return s
}

type QueryMessageByMsgKeyOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	HasMoreMessage *bool `type:"boolean" json:",omitempty"`

	MessageList []*MessageListForQueryMessageByMsgKeyOutput `type:"list" json:",omitempty"`
}

// String returns the string representation
func (s QueryMessageByMsgKeyOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s QueryMessageByMsgKeyOutput) GoString() string {
	return s.String()
}

// SetHasMoreMessage sets the HasMoreMessage field's value.
func (s *QueryMessageByMsgKeyOutput) SetHasMoreMessage(v bool) *QueryMessageByMsgKeyOutput {
	s.HasMoreMessage = &v
	return s
}

// SetMessageList sets the MessageList field's value.
func (s *QueryMessageByMsgKeyOutput) SetMessageList(v []*MessageListForQueryMessageByMsgKeyOutput) *QueryMessageByMsgKeyOutput {
	s.MessageList = v
	return s
}

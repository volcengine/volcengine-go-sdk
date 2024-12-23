// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package rocketmq

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opQueryDLQMessageByIdCommon = "QueryDLQMessageById"

// QueryDLQMessageByIdCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the QueryDLQMessageByIdCommon operation. The "output" return
// value will be populated with the QueryDLQMessageByIdCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned QueryDLQMessageByIdCommon Request to send the API call to the service.
// the "output" return value is not valid until after QueryDLQMessageByIdCommon Send returns without error.
//
// See QueryDLQMessageByIdCommon for more information on using the QueryDLQMessageByIdCommon
// API call, and error handling.
//
//    // Example sending a request using the QueryDLQMessageByIdCommonRequest method.
//    req, resp := client.QueryDLQMessageByIdCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ROCKETMQ) QueryDLQMessageByIdCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opQueryDLQMessageByIdCommon,
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

// QueryDLQMessageByIdCommon API operation for ROCKETMQ.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ROCKETMQ's
// API operation QueryDLQMessageByIdCommon for usage and error information.
func (c *ROCKETMQ) QueryDLQMessageByIdCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.QueryDLQMessageByIdCommonRequest(input)
	return out, req.Send()
}

// QueryDLQMessageByIdCommonWithContext is the same as QueryDLQMessageByIdCommon with the addition of
// the ability to pass a context and additional request options.
//
// See QueryDLQMessageByIdCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ROCKETMQ) QueryDLQMessageByIdCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.QueryDLQMessageByIdCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opQueryDLQMessageById = "QueryDLQMessageById"

// QueryDLQMessageByIdRequest generates a "volcengine/request.Request" representing the
// client's request for the QueryDLQMessageById operation. The "output" return
// value will be populated with the QueryDLQMessageByIdCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned QueryDLQMessageByIdCommon Request to send the API call to the service.
// the "output" return value is not valid until after QueryDLQMessageByIdCommon Send returns without error.
//
// See QueryDLQMessageById for more information on using the QueryDLQMessageById
// API call, and error handling.
//
//    // Example sending a request using the QueryDLQMessageByIdRequest method.
//    req, resp := client.QueryDLQMessageByIdRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ROCKETMQ) QueryDLQMessageByIdRequest(input *QueryDLQMessageByIdInput) (req *request.Request, output *QueryDLQMessageByIdOutput) {
	op := &request.Operation{
		Name:       opQueryDLQMessageById,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &QueryDLQMessageByIdInput{}
	}

	output = &QueryDLQMessageByIdOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// QueryDLQMessageById API operation for ROCKETMQ.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ROCKETMQ's
// API operation QueryDLQMessageById for usage and error information.
func (c *ROCKETMQ) QueryDLQMessageById(input *QueryDLQMessageByIdInput) (*QueryDLQMessageByIdOutput, error) {
	req, out := c.QueryDLQMessageByIdRequest(input)
	return out, req.Send()
}

// QueryDLQMessageByIdWithContext is the same as QueryDLQMessageById with the addition of
// the ability to pass a context and additional request options.
//
// See QueryDLQMessageById for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ROCKETMQ) QueryDLQMessageByIdWithContext(ctx volcengine.Context, input *QueryDLQMessageByIdInput, opts ...request.Option) (*QueryDLQMessageByIdOutput, error) {
	req, out := c.QueryDLQMessageByIdRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DLQMessageInfoForQueryDLQMessageByIdOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	CreateTimestamp *int64 `type:"int64" json:",omitempty"`

	IsExist *bool `type:"boolean" json:",omitempty"`

	MessageId *string `type:"string" json:",omitempty"`

	MessageKey *string `type:"string" json:",omitempty"`

	MessageSize *int32 `type:"int32" json:",omitempty"`

	ProducerHost *string `type:"string" json:",omitempty"`

	ReconsumeTimes *int32 `type:"int32" json:",omitempty"`

	StoreTimestamp *int64 `type:"int64" json:",omitempty"`

	Tag *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s DLQMessageInfoForQueryDLQMessageByIdOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DLQMessageInfoForQueryDLQMessageByIdOutput) GoString() string {
	return s.String()
}

// SetCreateTimestamp sets the CreateTimestamp field's value.
func (s *DLQMessageInfoForQueryDLQMessageByIdOutput) SetCreateTimestamp(v int64) *DLQMessageInfoForQueryDLQMessageByIdOutput {
	s.CreateTimestamp = &v
	return s
}

// SetIsExist sets the IsExist field's value.
func (s *DLQMessageInfoForQueryDLQMessageByIdOutput) SetIsExist(v bool) *DLQMessageInfoForQueryDLQMessageByIdOutput {
	s.IsExist = &v
	return s
}

// SetMessageId sets the MessageId field's value.
func (s *DLQMessageInfoForQueryDLQMessageByIdOutput) SetMessageId(v string) *DLQMessageInfoForQueryDLQMessageByIdOutput {
	s.MessageId = &v
	return s
}

// SetMessageKey sets the MessageKey field's value.
func (s *DLQMessageInfoForQueryDLQMessageByIdOutput) SetMessageKey(v string) *DLQMessageInfoForQueryDLQMessageByIdOutput {
	s.MessageKey = &v
	return s
}

// SetMessageSize sets the MessageSize field's value.
func (s *DLQMessageInfoForQueryDLQMessageByIdOutput) SetMessageSize(v int32) *DLQMessageInfoForQueryDLQMessageByIdOutput {
	s.MessageSize = &v
	return s
}

// SetProducerHost sets the ProducerHost field's value.
func (s *DLQMessageInfoForQueryDLQMessageByIdOutput) SetProducerHost(v string) *DLQMessageInfoForQueryDLQMessageByIdOutput {
	s.ProducerHost = &v
	return s
}

// SetReconsumeTimes sets the ReconsumeTimes field's value.
func (s *DLQMessageInfoForQueryDLQMessageByIdOutput) SetReconsumeTimes(v int32) *DLQMessageInfoForQueryDLQMessageByIdOutput {
	s.ReconsumeTimes = &v
	return s
}

// SetStoreTimestamp sets the StoreTimestamp field's value.
func (s *DLQMessageInfoForQueryDLQMessageByIdOutput) SetStoreTimestamp(v int64) *DLQMessageInfoForQueryDLQMessageByIdOutput {
	s.StoreTimestamp = &v
	return s
}

// SetTag sets the Tag field's value.
func (s *DLQMessageInfoForQueryDLQMessageByIdOutput) SetTag(v string) *DLQMessageInfoForQueryDLQMessageByIdOutput {
	s.Tag = &v
	return s
}

type QueryDLQMessageByIdInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	// GroupId is a required field
	GroupId *string `type:"string" json:",omitempty" required:"true"`

	// InstanceId is a required field
	InstanceId *string `type:"string" json:",omitempty" required:"true"`

	// MessageId is a required field
	MessageId *string `type:"string" json:",omitempty" required:"true"`
}

// String returns the string representation
func (s QueryDLQMessageByIdInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s QueryDLQMessageByIdInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *QueryDLQMessageByIdInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "QueryDLQMessageByIdInput"}
	if s.GroupId == nil {
		invalidParams.Add(request.NewErrParamRequired("GroupId"))
	}
	if s.InstanceId == nil {
		invalidParams.Add(request.NewErrParamRequired("InstanceId"))
	}
	if s.MessageId == nil {
		invalidParams.Add(request.NewErrParamRequired("MessageId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetGroupId sets the GroupId field's value.
func (s *QueryDLQMessageByIdInput) SetGroupId(v string) *QueryDLQMessageByIdInput {
	s.GroupId = &v
	return s
}

// SetInstanceId sets the InstanceId field's value.
func (s *QueryDLQMessageByIdInput) SetInstanceId(v string) *QueryDLQMessageByIdInput {
	s.InstanceId = &v
	return s
}

// SetMessageId sets the MessageId field's value.
func (s *QueryDLQMessageByIdInput) SetMessageId(v string) *QueryDLQMessageByIdInput {
	s.MessageId = &v
	return s
}

type QueryDLQMessageByIdOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	DLQMessageInfo *DLQMessageInfoForQueryDLQMessageByIdOutput `type:"structure" json:",omitempty"`
}

// String returns the string representation
func (s QueryDLQMessageByIdOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s QueryDLQMessageByIdOutput) GoString() string {
	return s.String()
}

// SetDLQMessageInfo sets the DLQMessageInfo field's value.
func (s *QueryDLQMessageByIdOutput) SetDLQMessageInfo(v *DLQMessageInfoForQueryDLQMessageByIdOutput) *QueryDLQMessageByIdOutput {
	s.DLQMessageInfo = v
	return s
}

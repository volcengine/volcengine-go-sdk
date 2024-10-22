// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package rocketmq

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDeleteTopicCommon = "DeleteTopic"

// DeleteTopicCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DeleteTopicCommon operation. The "output" return
// value will be populated with the DeleteTopicCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeleteTopicCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeleteTopicCommon Send returns without error.
//
// See DeleteTopicCommon for more information on using the DeleteTopicCommon
// API call, and error handling.
//
//    // Example sending a request using the DeleteTopicCommonRequest method.
//    req, resp := client.DeleteTopicCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ROCKETMQ) DeleteTopicCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDeleteTopicCommon,
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

// DeleteTopicCommon API operation for ROCKETMQ.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ROCKETMQ's
// API operation DeleteTopicCommon for usage and error information.
func (c *ROCKETMQ) DeleteTopicCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DeleteTopicCommonRequest(input)
	return out, req.Send()
}

// DeleteTopicCommonWithContext is the same as DeleteTopicCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteTopicCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ROCKETMQ) DeleteTopicCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DeleteTopicCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDeleteTopic = "DeleteTopic"

// DeleteTopicRequest generates a "volcengine/request.Request" representing the
// client's request for the DeleteTopic operation. The "output" return
// value will be populated with the DeleteTopicCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeleteTopicCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeleteTopicCommon Send returns without error.
//
// See DeleteTopic for more information on using the DeleteTopic
// API call, and error handling.
//
//    // Example sending a request using the DeleteTopicRequest method.
//    req, resp := client.DeleteTopicRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ROCKETMQ) DeleteTopicRequest(input *DeleteTopicInput) (req *request.Request, output *DeleteTopicOutput) {
	op := &request.Operation{
		Name:       opDeleteTopic,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteTopicInput{}
	}

	output = &DeleteTopicOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// DeleteTopic API operation for ROCKETMQ.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ROCKETMQ's
// API operation DeleteTopic for usage and error information.
func (c *ROCKETMQ) DeleteTopic(input *DeleteTopicInput) (*DeleteTopicOutput, error) {
	req, out := c.DeleteTopicRequest(input)
	return out, req.Send()
}

// DeleteTopicWithContext is the same as DeleteTopic with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteTopic for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ROCKETMQ) DeleteTopicWithContext(ctx volcengine.Context, input *DeleteTopicInput, opts ...request.Option) (*DeleteTopicOutput, error) {
	req, out := c.DeleteTopicRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DeleteTopicInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	// InstanceId is a required field
	InstanceId *string `type:"string" json:",omitempty" required:"true"`

	// TopicName is a required field
	TopicName *string `type:"string" json:",omitempty" required:"true"`
}

// String returns the string representation
func (s DeleteTopicInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteTopicInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteTopicInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DeleteTopicInput"}
	if s.InstanceId == nil {
		invalidParams.Add(request.NewErrParamRequired("InstanceId"))
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
func (s *DeleteTopicInput) SetInstanceId(v string) *DeleteTopicInput {
	s.InstanceId = &v
	return s
}

// SetTopicName sets the TopicName field's value.
func (s *DeleteTopicInput) SetTopicName(v string) *DeleteTopicInput {
	s.TopicName = &v
	return s
}

type DeleteTopicOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s DeleteTopicOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteTopicOutput) GoString() string {
	return s.String()
}

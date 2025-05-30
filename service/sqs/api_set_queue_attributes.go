// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package sqs

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opSetQueueAttributesCommon = "SetQueueAttributes"

// SetQueueAttributesCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the SetQueueAttributesCommon operation. The "output" return
// value will be populated with the SetQueueAttributesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned SetQueueAttributesCommon Request to send the API call to the service.
// the "output" return value is not valid until after SetQueueAttributesCommon Send returns without error.
//
// See SetQueueAttributesCommon for more information on using the SetQueueAttributesCommon
// API call, and error handling.
//
//    // Example sending a request using the SetQueueAttributesCommonRequest method.
//    req, resp := client.SetQueueAttributesCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *SQS) SetQueueAttributesCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opSetQueueAttributesCommon,
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

// SetQueueAttributesCommon API operation for SQS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for SQS's
// API operation SetQueueAttributesCommon for usage and error information.
func (c *SQS) SetQueueAttributesCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.SetQueueAttributesCommonRequest(input)
	return out, req.Send()
}

// SetQueueAttributesCommonWithContext is the same as SetQueueAttributesCommon with the addition of
// the ability to pass a context and additional request options.
//
// See SetQueueAttributesCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *SQS) SetQueueAttributesCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.SetQueueAttributesCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opSetQueueAttributes = "SetQueueAttributes"

// SetQueueAttributesRequest generates a "volcengine/request.Request" representing the
// client's request for the SetQueueAttributes operation. The "output" return
// value will be populated with the SetQueueAttributesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned SetQueueAttributesCommon Request to send the API call to the service.
// the "output" return value is not valid until after SetQueueAttributesCommon Send returns without error.
//
// See SetQueueAttributes for more information on using the SetQueueAttributes
// API call, and error handling.
//
//    // Example sending a request using the SetQueueAttributesRequest method.
//    req, resp := client.SetQueueAttributesRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *SQS) SetQueueAttributesRequest(input *SetQueueAttributesInput) (req *request.Request, output *SetQueueAttributesOutput) {
	op := &request.Operation{
		Name:       opSetQueueAttributes,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &SetQueueAttributesInput{}
	}

	output = &SetQueueAttributesOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// SetQueueAttributes API operation for SQS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for SQS's
// API operation SetQueueAttributes for usage and error information.
func (c *SQS) SetQueueAttributes(input *SetQueueAttributesInput) (*SetQueueAttributesOutput, error) {
	req, out := c.SetQueueAttributesRequest(input)
	return out, req.Send()
}

// SetQueueAttributesWithContext is the same as SetQueueAttributes with the addition of
// the ability to pass a context and additional request options.
//
// See SetQueueAttributes for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *SQS) SetQueueAttributesWithContext(ctx volcengine.Context, input *SetQueueAttributesInput, opts ...request.Option) (*SetQueueAttributesOutput, error) {
	req, out := c.SetQueueAttributesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type SetQueueAttributesInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	MessageRetentionPeriod *int32 `min:"60" max:"604800" type:"int32" json:",omitempty"`

	// QueueTrn is a required field
	QueueTrn *string `type:"string" json:",omitempty" required:"true"`

	VisibilityTimeout *int32 `max:"600" type:"int32" json:",omitempty"`
}

// String returns the string representation
func (s SetQueueAttributesInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s SetQueueAttributesInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *SetQueueAttributesInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "SetQueueAttributesInput"}
	if s.MessageRetentionPeriod != nil && *s.MessageRetentionPeriod < 60 {
		invalidParams.Add(request.NewErrParamMinValue("MessageRetentionPeriod", 60))
	}
	if s.MessageRetentionPeriod != nil && *s.MessageRetentionPeriod > 604800 {
		invalidParams.Add(request.NewErrParamMaxValue("MessageRetentionPeriod", 604800))
	}
	if s.QueueTrn == nil {
		invalidParams.Add(request.NewErrParamRequired("QueueTrn"))
	}
	if s.VisibilityTimeout != nil && *s.VisibilityTimeout > 600 {
		invalidParams.Add(request.NewErrParamMaxValue("VisibilityTimeout", 600))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetMessageRetentionPeriod sets the MessageRetentionPeriod field's value.
func (s *SetQueueAttributesInput) SetMessageRetentionPeriod(v int32) *SetQueueAttributesInput {
	s.MessageRetentionPeriod = &v
	return s
}

// SetQueueTrn sets the QueueTrn field's value.
func (s *SetQueueAttributesInput) SetQueueTrn(v string) *SetQueueAttributesInput {
	s.QueueTrn = &v
	return s
}

// SetVisibilityTimeout sets the VisibilityTimeout field's value.
func (s *SetQueueAttributesInput) SetVisibilityTimeout(v int32) *SetQueueAttributesInput {
	s.VisibilityTimeout = &v
	return s
}

type SetQueueAttributesOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s SetQueueAttributesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s SetQueueAttributesOutput) GoString() string {
	return s.String()
}

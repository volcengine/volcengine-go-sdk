// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package sqs

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opGetQueueAttributesCommon = "GetQueueAttributes"

// GetQueueAttributesCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the GetQueueAttributesCommon operation. The "output" return
// value will be populated with the GetQueueAttributesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned GetQueueAttributesCommon Request to send the API call to the service.
// the "output" return value is not valid until after GetQueueAttributesCommon Send returns without error.
//
// See GetQueueAttributesCommon for more information on using the GetQueueAttributesCommon
// API call, and error handling.
//
//    // Example sending a request using the GetQueueAttributesCommonRequest method.
//    req, resp := client.GetQueueAttributesCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *SQS) GetQueueAttributesCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opGetQueueAttributesCommon,
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

// GetQueueAttributesCommon API operation for SQS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for SQS's
// API operation GetQueueAttributesCommon for usage and error information.
func (c *SQS) GetQueueAttributesCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.GetQueueAttributesCommonRequest(input)
	return out, req.Send()
}

// GetQueueAttributesCommonWithContext is the same as GetQueueAttributesCommon with the addition of
// the ability to pass a context and additional request options.
//
// See GetQueueAttributesCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *SQS) GetQueueAttributesCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.GetQueueAttributesCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opGetQueueAttributes = "GetQueueAttributes"

// GetQueueAttributesRequest generates a "volcengine/request.Request" representing the
// client's request for the GetQueueAttributes operation. The "output" return
// value will be populated with the GetQueueAttributesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned GetQueueAttributesCommon Request to send the API call to the service.
// the "output" return value is not valid until after GetQueueAttributesCommon Send returns without error.
//
// See GetQueueAttributes for more information on using the GetQueueAttributes
// API call, and error handling.
//
//    // Example sending a request using the GetQueueAttributesRequest method.
//    req, resp := client.GetQueueAttributesRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *SQS) GetQueueAttributesRequest(input *GetQueueAttributesInput) (req *request.Request, output *GetQueueAttributesOutput) {
	op := &request.Operation{
		Name:       opGetQueueAttributes,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetQueueAttributesInput{}
	}

	output = &GetQueueAttributesOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// GetQueueAttributes API operation for SQS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for SQS's
// API operation GetQueueAttributes for usage and error information.
func (c *SQS) GetQueueAttributes(input *GetQueueAttributesInput) (*GetQueueAttributesOutput, error) {
	req, out := c.GetQueueAttributesRequest(input)
	return out, req.Send()
}

// GetQueueAttributesWithContext is the same as GetQueueAttributes with the addition of
// the ability to pass a context and additional request options.
//
// See GetQueueAttributes for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *SQS) GetQueueAttributesWithContext(ctx volcengine.Context, input *GetQueueAttributesInput, opts ...request.Option) (*GetQueueAttributesOutput, error) {
	req, out := c.GetQueueAttributesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type GetQueueAttributesInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	// QueueTrn is a required field
	QueueTrn *string `type:"string" json:",omitempty" required:"true"`
}

// String returns the string representation
func (s GetQueueAttributesInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s GetQueueAttributesInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetQueueAttributesInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "GetQueueAttributesInput"}
	if s.QueueTrn == nil {
		invalidParams.Add(request.NewErrParamRequired("QueueTrn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetQueueTrn sets the QueueTrn field's value.
func (s *GetQueueAttributesInput) SetQueueTrn(v string) *GetQueueAttributesInput {
	s.QueueTrn = &v
	return s
}

type GetQueueAttributesOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	CreatedAt *string `type:"string" json:",omitempty"`

	MessageRetentionPeriod *int32 `type:"int32" json:",omitempty"`

	QueueDisplayName *string `type:"string" json:",omitempty"`

	QueueTrn *string `type:"string" json:",omitempty"`

	VisibilityTimeout *int32 `type:"int32" json:",omitempty"`
}

// String returns the string representation
func (s GetQueueAttributesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s GetQueueAttributesOutput) GoString() string {
	return s.String()
}

// SetCreatedAt sets the CreatedAt field's value.
func (s *GetQueueAttributesOutput) SetCreatedAt(v string) *GetQueueAttributesOutput {
	s.CreatedAt = &v
	return s
}

// SetMessageRetentionPeriod sets the MessageRetentionPeriod field's value.
func (s *GetQueueAttributesOutput) SetMessageRetentionPeriod(v int32) *GetQueueAttributesOutput {
	s.MessageRetentionPeriod = &v
	return s
}

// SetQueueDisplayName sets the QueueDisplayName field's value.
func (s *GetQueueAttributesOutput) SetQueueDisplayName(v string) *GetQueueAttributesOutput {
	s.QueueDisplayName = &v
	return s
}

// SetQueueTrn sets the QueueTrn field's value.
func (s *GetQueueAttributesOutput) SetQueueTrn(v string) *GetQueueAttributesOutput {
	s.QueueTrn = &v
	return s
}

// SetVisibilityTimeout sets the VisibilityTimeout field's value.
func (s *GetQueueAttributesOutput) SetVisibilityTimeout(v int32) *GetQueueAttributesOutput {
	s.VisibilityTimeout = &v
	return s
}

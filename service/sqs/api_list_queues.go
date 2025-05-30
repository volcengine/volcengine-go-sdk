// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package sqs

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opListQueuesCommon = "ListQueues"

// ListQueuesCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ListQueuesCommon operation. The "output" return
// value will be populated with the ListQueuesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListQueuesCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListQueuesCommon Send returns without error.
//
// See ListQueuesCommon for more information on using the ListQueuesCommon
// API call, and error handling.
//
//    // Example sending a request using the ListQueuesCommonRequest method.
//    req, resp := client.ListQueuesCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *SQS) ListQueuesCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opListQueuesCommon,
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

// ListQueuesCommon API operation for SQS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for SQS's
// API operation ListQueuesCommon for usage and error information.
func (c *SQS) ListQueuesCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ListQueuesCommonRequest(input)
	return out, req.Send()
}

// ListQueuesCommonWithContext is the same as ListQueuesCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ListQueuesCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *SQS) ListQueuesCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ListQueuesCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opListQueues = "ListQueues"

// ListQueuesRequest generates a "volcengine/request.Request" representing the
// client's request for the ListQueues operation. The "output" return
// value will be populated with the ListQueuesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListQueuesCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListQueuesCommon Send returns without error.
//
// See ListQueues for more information on using the ListQueues
// API call, and error handling.
//
//    // Example sending a request using the ListQueuesRequest method.
//    req, resp := client.ListQueuesRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *SQS) ListQueuesRequest(input *ListQueuesInput) (req *request.Request, output *ListQueuesOutput) {
	op := &request.Operation{
		Name:       opListQueues,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListQueuesInput{}
	}

	output = &ListQueuesOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// ListQueues API operation for SQS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for SQS's
// API operation ListQueues for usage and error information.
func (c *SQS) ListQueues(input *ListQueuesInput) (*ListQueuesOutput, error) {
	req, out := c.ListQueuesRequest(input)
	return out, req.Send()
}

// ListQueuesWithContext is the same as ListQueues with the addition of
// the ability to pass a context and additional request options.
//
// See ListQueues for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *SQS) ListQueuesWithContext(ctx volcengine.Context, input *ListQueuesInput, opts ...request.Option) (*ListQueuesOutput, error) {
	req, out := c.ListQueuesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ListQueuesInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	MaxResults *int32 `min:"1" max:"100" type:"int32" json:",omitempty"`

	NextToken *string `type:"string" json:",omitempty"`

	QueueDisplayNamePrefix *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s ListQueuesInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListQueuesInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListQueuesInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ListQueuesInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(request.NewErrParamMinValue("MaxResults", 1))
	}
	if s.MaxResults != nil && *s.MaxResults > 100 {
		invalidParams.Add(request.NewErrParamMaxValue("MaxResults", 100))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetMaxResults sets the MaxResults field's value.
func (s *ListQueuesInput) SetMaxResults(v int32) *ListQueuesInput {
	s.MaxResults = &v
	return s
}

// SetNextToken sets the NextToken field's value.
func (s *ListQueuesInput) SetNextToken(v string) *ListQueuesInput {
	s.NextToken = &v
	return s
}

// SetQueueDisplayNamePrefix sets the QueueDisplayNamePrefix field's value.
func (s *ListQueuesInput) SetQueueDisplayNamePrefix(v string) *ListQueuesInput {
	s.QueueDisplayNamePrefix = &v
	return s
}

type ListQueuesOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	NextToken *string `type:"string" json:",omitempty"`

	Queues []*QueueForListQueuesOutput `type:"list" json:",omitempty"`

	TotalCount *int32 `type:"int32" json:",omitempty"`
}

// String returns the string representation
func (s ListQueuesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListQueuesOutput) GoString() string {
	return s.String()
}

// SetNextToken sets the NextToken field's value.
func (s *ListQueuesOutput) SetNextToken(v string) *ListQueuesOutput {
	s.NextToken = &v
	return s
}

// SetQueues sets the Queues field's value.
func (s *ListQueuesOutput) SetQueues(v []*QueueForListQueuesOutput) *ListQueuesOutput {
	s.Queues = v
	return s
}

// SetTotalCount sets the TotalCount field's value.
func (s *ListQueuesOutput) SetTotalCount(v int32) *ListQueuesOutput {
	s.TotalCount = &v
	return s
}

type QueueForListQueuesOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	CreatedAt *string `type:"string" json:",omitempty"`

	QueueDisplayName *string `type:"string" json:",omitempty"`

	QueueTrn *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s QueueForListQueuesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s QueueForListQueuesOutput) GoString() string {
	return s.String()
}

// SetCreatedAt sets the CreatedAt field's value.
func (s *QueueForListQueuesOutput) SetCreatedAt(v string) *QueueForListQueuesOutput {
	s.CreatedAt = &v
	return s
}

// SetQueueDisplayName sets the QueueDisplayName field's value.
func (s *QueueForListQueuesOutput) SetQueueDisplayName(v string) *QueueForListQueuesOutput {
	s.QueueDisplayName = &v
	return s
}

// SetQueueTrn sets the QueueTrn field's value.
func (s *QueueForListQueuesOutput) SetQueueTrn(v string) *QueueForListQueuesOutput {
	s.QueueTrn = &v
	return s
}

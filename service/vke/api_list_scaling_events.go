// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package vke

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opListScalingEventsCommon = "ListScalingEvents"

// ListScalingEventsCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ListScalingEventsCommon operation. The "output" return
// value will be populated with the ListScalingEventsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListScalingEventsCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListScalingEventsCommon Send returns without error.
//
// See ListScalingEventsCommon for more information on using the ListScalingEventsCommon
// API call, and error handling.
//
//    // Example sending a request using the ListScalingEventsCommonRequest method.
//    req, resp := client.ListScalingEventsCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VKE) ListScalingEventsCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opListScalingEventsCommon,
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

// ListScalingEventsCommon API operation for VKE.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VKE's
// API operation ListScalingEventsCommon for usage and error information.
func (c *VKE) ListScalingEventsCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ListScalingEventsCommonRequest(input)
	return out, req.Send()
}

// ListScalingEventsCommonWithContext is the same as ListScalingEventsCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ListScalingEventsCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VKE) ListScalingEventsCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ListScalingEventsCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opListScalingEvents = "ListScalingEvents"

// ListScalingEventsRequest generates a "volcengine/request.Request" representing the
// client's request for the ListScalingEvents operation. The "output" return
// value will be populated with the ListScalingEventsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListScalingEventsCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListScalingEventsCommon Send returns without error.
//
// See ListScalingEvents for more information on using the ListScalingEvents
// API call, and error handling.
//
//    // Example sending a request using the ListScalingEventsRequest method.
//    req, resp := client.ListScalingEventsRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VKE) ListScalingEventsRequest(input *ListScalingEventsInput) (req *request.Request, output *ListScalingEventsOutput) {
	op := &request.Operation{
		Name:       opListScalingEvents,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListScalingEventsInput{}
	}

	output = &ListScalingEventsOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// ListScalingEvents API operation for VKE.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VKE's
// API operation ListScalingEvents for usage and error information.
func (c *VKE) ListScalingEvents(input *ListScalingEventsInput) (*ListScalingEventsOutput, error) {
	req, out := c.ListScalingEventsRequest(input)
	return out, req.Send()
}

// ListScalingEventsWithContext is the same as ListScalingEvents with the addition of
// the ability to pass a context and additional request options.
//
// See ListScalingEvents for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VKE) ListScalingEventsWithContext(ctx volcengine.Context, input *ListScalingEventsInput, opts ...request.Option) (*ListScalingEventsOutput, error) {
	req, out := c.ListScalingEventsRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type FilterForListScalingEventsInput struct {
	_ struct{} `type:"structure"`

	EndTime *string `type:"string"`

	Ids []*string `type:"list"`

	NodeIds []*string `type:"list"`

	NodePoolIds []*string `type:"list"`

	StartTime *string `type:"string"`

	Statuses []*string `type:"list"`

	Type *string `type:"string"`
}

// String returns the string representation
func (s FilterForListScalingEventsInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s FilterForListScalingEventsInput) GoString() string {
	return s.String()
}

// SetEndTime sets the EndTime field's value.
func (s *FilterForListScalingEventsInput) SetEndTime(v string) *FilterForListScalingEventsInput {
	s.EndTime = &v
	return s
}

// SetIds sets the Ids field's value.
func (s *FilterForListScalingEventsInput) SetIds(v []*string) *FilterForListScalingEventsInput {
	s.Ids = v
	return s
}

// SetNodeIds sets the NodeIds field's value.
func (s *FilterForListScalingEventsInput) SetNodeIds(v []*string) *FilterForListScalingEventsInput {
	s.NodeIds = v
	return s
}

// SetNodePoolIds sets the NodePoolIds field's value.
func (s *FilterForListScalingEventsInput) SetNodePoolIds(v []*string) *FilterForListScalingEventsInput {
	s.NodePoolIds = v
	return s
}

// SetStartTime sets the StartTime field's value.
func (s *FilterForListScalingEventsInput) SetStartTime(v string) *FilterForListScalingEventsInput {
	s.StartTime = &v
	return s
}

// SetStatuses sets the Statuses field's value.
func (s *FilterForListScalingEventsInput) SetStatuses(v []*string) *FilterForListScalingEventsInput {
	s.Statuses = v
	return s
}

// SetType sets the Type field's value.
func (s *FilterForListScalingEventsInput) SetType(v string) *FilterForListScalingEventsInput {
	s.Type = &v
	return s
}

type ItemForListScalingEventsOutput struct {
	_ struct{} `type:"structure"`

	EndTime *string `type:"string"`

	Id *string `type:"string"`

	Message *string `type:"string"`

	NodeIds []*string `type:"list"`

	NodePoolId *string `type:"string"`

	Replicas *int32 `type:"int32"`

	StartTime *string `type:"string"`

	Status *string `type:"string"`

	Type *string `type:"string"`
}

// String returns the string representation
func (s ItemForListScalingEventsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ItemForListScalingEventsOutput) GoString() string {
	return s.String()
}

// SetEndTime sets the EndTime field's value.
func (s *ItemForListScalingEventsOutput) SetEndTime(v string) *ItemForListScalingEventsOutput {
	s.EndTime = &v
	return s
}

// SetId sets the Id field's value.
func (s *ItemForListScalingEventsOutput) SetId(v string) *ItemForListScalingEventsOutput {
	s.Id = &v
	return s
}

// SetMessage sets the Message field's value.
func (s *ItemForListScalingEventsOutput) SetMessage(v string) *ItemForListScalingEventsOutput {
	s.Message = &v
	return s
}

// SetNodeIds sets the NodeIds field's value.
func (s *ItemForListScalingEventsOutput) SetNodeIds(v []*string) *ItemForListScalingEventsOutput {
	s.NodeIds = v
	return s
}

// SetNodePoolId sets the NodePoolId field's value.
func (s *ItemForListScalingEventsOutput) SetNodePoolId(v string) *ItemForListScalingEventsOutput {
	s.NodePoolId = &v
	return s
}

// SetReplicas sets the Replicas field's value.
func (s *ItemForListScalingEventsOutput) SetReplicas(v int32) *ItemForListScalingEventsOutput {
	s.Replicas = &v
	return s
}

// SetStartTime sets the StartTime field's value.
func (s *ItemForListScalingEventsOutput) SetStartTime(v string) *ItemForListScalingEventsOutput {
	s.StartTime = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *ItemForListScalingEventsOutput) SetStatus(v string) *ItemForListScalingEventsOutput {
	s.Status = &v
	return s
}

// SetType sets the Type field's value.
func (s *ItemForListScalingEventsOutput) SetType(v string) *ItemForListScalingEventsOutput {
	s.Type = &v
	return s
}

type ListScalingEventsInput struct {
	_ struct{} `type:"structure"`

	// ClusterId is a required field
	ClusterId *string `type:"string" required:"true"`

	Filter *FilterForListScalingEventsInput `type:"structure"`

	PageNumber *int32 `type:"int32"`

	PageSize *int32 `type:"int32"`
}

// String returns the string representation
func (s ListScalingEventsInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListScalingEventsInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListScalingEventsInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ListScalingEventsInput"}
	if s.ClusterId == nil {
		invalidParams.Add(request.NewErrParamRequired("ClusterId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetClusterId sets the ClusterId field's value.
func (s *ListScalingEventsInput) SetClusterId(v string) *ListScalingEventsInput {
	s.ClusterId = &v
	return s
}

// SetFilter sets the Filter field's value.
func (s *ListScalingEventsInput) SetFilter(v *FilterForListScalingEventsInput) *ListScalingEventsInput {
	s.Filter = v
	return s
}

// SetPageNumber sets the PageNumber field's value.
func (s *ListScalingEventsInput) SetPageNumber(v int32) *ListScalingEventsInput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *ListScalingEventsInput) SetPageSize(v int32) *ListScalingEventsInput {
	s.PageSize = &v
	return s
}

type ListScalingEventsOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	Items []*ItemForListScalingEventsOutput `type:"list"`

	PageNumber *int32 `type:"int32"`

	PageSize *int32 `type:"int32"`

	TotalCount *int32 `type:"int32"`
}

// String returns the string representation
func (s ListScalingEventsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListScalingEventsOutput) GoString() string {
	return s.String()
}

// SetItems sets the Items field's value.
func (s *ListScalingEventsOutput) SetItems(v []*ItemForListScalingEventsOutput) *ListScalingEventsOutput {
	s.Items = v
	return s
}

// SetPageNumber sets the PageNumber field's value.
func (s *ListScalingEventsOutput) SetPageNumber(v int32) *ListScalingEventsOutput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *ListScalingEventsOutput) SetPageSize(v int32) *ListScalingEventsOutput {
	s.PageSize = &v
	return s
}

// SetTotalCount sets the TotalCount field's value.
func (s *ListScalingEventsOutput) SetTotalCount(v int32) *ListScalingEventsOutput {
	s.TotalCount = &v
	return s
}
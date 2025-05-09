// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package mongodb

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribePlannedEventsCommon = "DescribePlannedEvents"

// DescribePlannedEventsCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribePlannedEventsCommon operation. The "output" return
// value will be populated with the DescribePlannedEventsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribePlannedEventsCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribePlannedEventsCommon Send returns without error.
//
// See DescribePlannedEventsCommon for more information on using the DescribePlannedEventsCommon
// API call, and error handling.
//
//    // Example sending a request using the DescribePlannedEventsCommonRequest method.
//    req, resp := client.DescribePlannedEventsCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *MONGODB) DescribePlannedEventsCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribePlannedEventsCommon,
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

// DescribePlannedEventsCommon API operation for MONGODB.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for MONGODB's
// API operation DescribePlannedEventsCommon for usage and error information.
func (c *MONGODB) DescribePlannedEventsCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribePlannedEventsCommonRequest(input)
	return out, req.Send()
}

// DescribePlannedEventsCommonWithContext is the same as DescribePlannedEventsCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribePlannedEventsCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *MONGODB) DescribePlannedEventsCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribePlannedEventsCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribePlannedEvents = "DescribePlannedEvents"

// DescribePlannedEventsRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribePlannedEvents operation. The "output" return
// value will be populated with the DescribePlannedEventsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribePlannedEventsCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribePlannedEventsCommon Send returns without error.
//
// See DescribePlannedEvents for more information on using the DescribePlannedEvents
// API call, and error handling.
//
//    // Example sending a request using the DescribePlannedEventsRequest method.
//    req, resp := client.DescribePlannedEventsRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *MONGODB) DescribePlannedEventsRequest(input *DescribePlannedEventsInput) (req *request.Request, output *DescribePlannedEventsOutput) {
	op := &request.Operation{
		Name:       opDescribePlannedEvents,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribePlannedEventsInput{}
	}

	output = &DescribePlannedEventsOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// DescribePlannedEvents API operation for MONGODB.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for MONGODB's
// API operation DescribePlannedEvents for usage and error information.
func (c *MONGODB) DescribePlannedEvents(input *DescribePlannedEventsInput) (*DescribePlannedEventsOutput, error) {
	req, out := c.DescribePlannedEventsRequest(input)
	return out, req.Send()
}

// DescribePlannedEventsWithContext is the same as DescribePlannedEvents with the addition of
// the ability to pass a context and additional request options.
//
// See DescribePlannedEvents for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *MONGODB) DescribePlannedEventsWithContext(ctx volcengine.Context, input *DescribePlannedEventsInput, opts ...request.Option) (*DescribePlannedEventsOutput, error) {
	req, out := c.DescribePlannedEventsRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DescribePlannedEventsInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	EndTime *string `type:"string" json:",omitempty"`

	EventAction *string `type:"string" json:",omitempty"`

	EventId *string `type:"string" json:",omitempty"`

	EventStatus []*string `type:"list" json:",omitempty"`

	EventType []*string `type:"list" json:",omitempty"`

	InstanceId *string `type:"string" json:",omitempty"`

	InstanceName *string `type:"string" json:",omitempty"`

	PageNumber *int32 `type:"int32" json:",omitempty"`

	PageSize *int32 `type:"int32" json:",omitempty"`

	ProjectName *string `type:"string" json:",omitempty"`

	StartTime *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s DescribePlannedEventsInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribePlannedEventsInput) GoString() string {
	return s.String()
}

// SetEndTime sets the EndTime field's value.
func (s *DescribePlannedEventsInput) SetEndTime(v string) *DescribePlannedEventsInput {
	s.EndTime = &v
	return s
}

// SetEventAction sets the EventAction field's value.
func (s *DescribePlannedEventsInput) SetEventAction(v string) *DescribePlannedEventsInput {
	s.EventAction = &v
	return s
}

// SetEventId sets the EventId field's value.
func (s *DescribePlannedEventsInput) SetEventId(v string) *DescribePlannedEventsInput {
	s.EventId = &v
	return s
}

// SetEventStatus sets the EventStatus field's value.
func (s *DescribePlannedEventsInput) SetEventStatus(v []*string) *DescribePlannedEventsInput {
	s.EventStatus = v
	return s
}

// SetEventType sets the EventType field's value.
func (s *DescribePlannedEventsInput) SetEventType(v []*string) *DescribePlannedEventsInput {
	s.EventType = v
	return s
}

// SetInstanceId sets the InstanceId field's value.
func (s *DescribePlannedEventsInput) SetInstanceId(v string) *DescribePlannedEventsInput {
	s.InstanceId = &v
	return s
}

// SetInstanceName sets the InstanceName field's value.
func (s *DescribePlannedEventsInput) SetInstanceName(v string) *DescribePlannedEventsInput {
	s.InstanceName = &v
	return s
}

// SetPageNumber sets the PageNumber field's value.
func (s *DescribePlannedEventsInput) SetPageNumber(v int32) *DescribePlannedEventsInput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *DescribePlannedEventsInput) SetPageSize(v int32) *DescribePlannedEventsInput {
	s.PageSize = &v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *DescribePlannedEventsInput) SetProjectName(v string) *DescribePlannedEventsInput {
	s.ProjectName = &v
	return s
}

// SetStartTime sets the StartTime field's value.
func (s *DescribePlannedEventsInput) SetStartTime(v string) *DescribePlannedEventsInput {
	s.StartTime = &v
	return s
}

type DescribePlannedEventsOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	Events []*EventForDescribePlannedEventsOutput `type:"list" json:",omitempty"`

	Total *int32 `type:"int32" json:",omitempty"`
}

// String returns the string representation
func (s DescribePlannedEventsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribePlannedEventsOutput) GoString() string {
	return s.String()
}

// SetEvents sets the Events field's value.
func (s *DescribePlannedEventsOutput) SetEvents(v []*EventForDescribePlannedEventsOutput) *DescribePlannedEventsOutput {
	s.Events = v
	return s
}

// SetTotal sets the Total field's value.
func (s *DescribePlannedEventsOutput) SetTotal(v int32) *DescribePlannedEventsOutput {
	s.Total = &v
	return s
}

type EventForDescribePlannedEventsOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	CostTimeMs *int32 `type:"int32" json:",omitempty"`

	EventAction *string `type:"string" json:",omitempty"`

	EventId *string `type:"string" json:",omitempty"`

	EventStatus *string `type:"string" json:",omitempty"`

	EventType *string `type:"string" json:",omitempty"`

	FinishTime *string `type:"string" json:",omitempty"`

	InstanceId *string `type:"string" json:",omitempty"`

	InstanceName *string `type:"string" json:",omitempty"`

	MaxDelayTime *string `type:"string" json:",omitempty"`

	PlannedBeginTime *string `type:"string" json:",omitempty"`

	PlannedEndTime *string `type:"string" json:",omitempty"`

	ResourceName *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s EventForDescribePlannedEventsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s EventForDescribePlannedEventsOutput) GoString() string {
	return s.String()
}

// SetCostTimeMs sets the CostTimeMs field's value.
func (s *EventForDescribePlannedEventsOutput) SetCostTimeMs(v int32) *EventForDescribePlannedEventsOutput {
	s.CostTimeMs = &v
	return s
}

// SetEventAction sets the EventAction field's value.
func (s *EventForDescribePlannedEventsOutput) SetEventAction(v string) *EventForDescribePlannedEventsOutput {
	s.EventAction = &v
	return s
}

// SetEventId sets the EventId field's value.
func (s *EventForDescribePlannedEventsOutput) SetEventId(v string) *EventForDescribePlannedEventsOutput {
	s.EventId = &v
	return s
}

// SetEventStatus sets the EventStatus field's value.
func (s *EventForDescribePlannedEventsOutput) SetEventStatus(v string) *EventForDescribePlannedEventsOutput {
	s.EventStatus = &v
	return s
}

// SetEventType sets the EventType field's value.
func (s *EventForDescribePlannedEventsOutput) SetEventType(v string) *EventForDescribePlannedEventsOutput {
	s.EventType = &v
	return s
}

// SetFinishTime sets the FinishTime field's value.
func (s *EventForDescribePlannedEventsOutput) SetFinishTime(v string) *EventForDescribePlannedEventsOutput {
	s.FinishTime = &v
	return s
}

// SetInstanceId sets the InstanceId field's value.
func (s *EventForDescribePlannedEventsOutput) SetInstanceId(v string) *EventForDescribePlannedEventsOutput {
	s.InstanceId = &v
	return s
}

// SetInstanceName sets the InstanceName field's value.
func (s *EventForDescribePlannedEventsOutput) SetInstanceName(v string) *EventForDescribePlannedEventsOutput {
	s.InstanceName = &v
	return s
}

// SetMaxDelayTime sets the MaxDelayTime field's value.
func (s *EventForDescribePlannedEventsOutput) SetMaxDelayTime(v string) *EventForDescribePlannedEventsOutput {
	s.MaxDelayTime = &v
	return s
}

// SetPlannedBeginTime sets the PlannedBeginTime field's value.
func (s *EventForDescribePlannedEventsOutput) SetPlannedBeginTime(v string) *EventForDescribePlannedEventsOutput {
	s.PlannedBeginTime = &v
	return s
}

// SetPlannedEndTime sets the PlannedEndTime field's value.
func (s *EventForDescribePlannedEventsOutput) SetPlannedEndTime(v string) *EventForDescribePlannedEventsOutput {
	s.PlannedEndTime = &v
	return s
}

// SetResourceName sets the ResourceName field's value.
func (s *EventForDescribePlannedEventsOutput) SetResourceName(v string) *EventForDescribePlannedEventsOutput {
	s.ResourceName = &v
	return s
}

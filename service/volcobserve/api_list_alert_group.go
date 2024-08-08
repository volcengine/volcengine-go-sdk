// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package volcobserve

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opListAlertGroupCommon = "ListAlertGroup"

// ListAlertGroupCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ListAlertGroupCommon operation. The "output" return
// value will be populated with the ListAlertGroupCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListAlertGroupCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListAlertGroupCommon Send returns without error.
//
// See ListAlertGroupCommon for more information on using the ListAlertGroupCommon
// API call, and error handling.
//
//    // Example sending a request using the ListAlertGroupCommonRequest method.
//    req, resp := client.ListAlertGroupCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VOLCOBSERVE) ListAlertGroupCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opListAlertGroupCommon,
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

// ListAlertGroupCommon API operation for VOLC_OBSERVE.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VOLC_OBSERVE's
// API operation ListAlertGroupCommon for usage and error information.
func (c *VOLCOBSERVE) ListAlertGroupCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ListAlertGroupCommonRequest(input)
	return out, req.Send()
}

// ListAlertGroupCommonWithContext is the same as ListAlertGroupCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ListAlertGroupCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VOLCOBSERVE) ListAlertGroupCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ListAlertGroupCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opListAlertGroup = "ListAlertGroup"

// ListAlertGroupRequest generates a "volcengine/request.Request" representing the
// client's request for the ListAlertGroup operation. The "output" return
// value will be populated with the ListAlertGroupCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListAlertGroupCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListAlertGroupCommon Send returns without error.
//
// See ListAlertGroup for more information on using the ListAlertGroup
// API call, and error handling.
//
//    // Example sending a request using the ListAlertGroupRequest method.
//    req, resp := client.ListAlertGroupRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VOLCOBSERVE) ListAlertGroupRequest(input *ListAlertGroupInput) (req *request.Request, output *ListAlertGroupOutput) {
	op := &request.Operation{
		Name:       opListAlertGroup,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListAlertGroupInput{}
	}

	output = &ListAlertGroupOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// ListAlertGroup API operation for VOLC_OBSERVE.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VOLC_OBSERVE's
// API operation ListAlertGroup for usage and error information.
func (c *VOLCOBSERVE) ListAlertGroup(input *ListAlertGroupInput) (*ListAlertGroupOutput, error) {
	req, out := c.ListAlertGroupRequest(input)
	return out, req.Send()
}

// ListAlertGroupWithContext is the same as ListAlertGroup with the addition of
// the ability to pass a context and additional request options.
//
// See ListAlertGroup for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VOLCOBSERVE) ListAlertGroupWithContext(ctx volcengine.Context, input *ListAlertGroupInput, opts ...request.Option) (*ListAlertGroupOutput, error) {
	req, out := c.ListAlertGroupRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DataForListAlertGroupOutput struct {
	_ struct{} `type:"structure"`

	AlertCount *int64 `type:"integer"`

	AlertState *string `type:"string"`

	Dimension *string `type:"string"`

	Duration *string `type:"string"`

	EndAt *string `type:"string"`

	Id *string `type:"string"`

	Level *string `type:"string"`

	Namespace *string `type:"string"`

	Region *string `type:"string"`

	ResourceId *string `type:"string"`

	ResourceName *string `type:"string"`

	ResourceStatus *string `type:"string"`

	ResourceStatusCN *string `type:"string"`

	ResourceTags []*ResourceTagForListAlertGroupOutput `type:"list"`

	ResourceType *string `type:"string"`

	RuleId *string `type:"string"`

	RuleName *string `type:"string"`

	RuleStatus *string `type:"string"`

	RuleTriggerCondition *string `type:"string"`

	SendAlertCount *int64 `type:"integer"`

	StartAt *string `type:"string"`

	SubNamespace *string `type:"string"`
}

// String returns the string representation
func (s DataForListAlertGroupOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DataForListAlertGroupOutput) GoString() string {
	return s.String()
}

// SetAlertCount sets the AlertCount field's value.
func (s *DataForListAlertGroupOutput) SetAlertCount(v int64) *DataForListAlertGroupOutput {
	s.AlertCount = &v
	return s
}

// SetAlertState sets the AlertState field's value.
func (s *DataForListAlertGroupOutput) SetAlertState(v string) *DataForListAlertGroupOutput {
	s.AlertState = &v
	return s
}

// SetDimension sets the Dimension field's value.
func (s *DataForListAlertGroupOutput) SetDimension(v string) *DataForListAlertGroupOutput {
	s.Dimension = &v
	return s
}

// SetDuration sets the Duration field's value.
func (s *DataForListAlertGroupOutput) SetDuration(v string) *DataForListAlertGroupOutput {
	s.Duration = &v
	return s
}

// SetEndAt sets the EndAt field's value.
func (s *DataForListAlertGroupOutput) SetEndAt(v string) *DataForListAlertGroupOutput {
	s.EndAt = &v
	return s
}

// SetId sets the Id field's value.
func (s *DataForListAlertGroupOutput) SetId(v string) *DataForListAlertGroupOutput {
	s.Id = &v
	return s
}

// SetLevel sets the Level field's value.
func (s *DataForListAlertGroupOutput) SetLevel(v string) *DataForListAlertGroupOutput {
	s.Level = &v
	return s
}

// SetNamespace sets the Namespace field's value.
func (s *DataForListAlertGroupOutput) SetNamespace(v string) *DataForListAlertGroupOutput {
	s.Namespace = &v
	return s
}

// SetRegion sets the Region field's value.
func (s *DataForListAlertGroupOutput) SetRegion(v string) *DataForListAlertGroupOutput {
	s.Region = &v
	return s
}

// SetResourceId sets the ResourceId field's value.
func (s *DataForListAlertGroupOutput) SetResourceId(v string) *DataForListAlertGroupOutput {
	s.ResourceId = &v
	return s
}

// SetResourceName sets the ResourceName field's value.
func (s *DataForListAlertGroupOutput) SetResourceName(v string) *DataForListAlertGroupOutput {
	s.ResourceName = &v
	return s
}

// SetResourceStatus sets the ResourceStatus field's value.
func (s *DataForListAlertGroupOutput) SetResourceStatus(v string) *DataForListAlertGroupOutput {
	s.ResourceStatus = &v
	return s
}

// SetResourceStatusCN sets the ResourceStatusCN field's value.
func (s *DataForListAlertGroupOutput) SetResourceStatusCN(v string) *DataForListAlertGroupOutput {
	s.ResourceStatusCN = &v
	return s
}

// SetResourceTags sets the ResourceTags field's value.
func (s *DataForListAlertGroupOutput) SetResourceTags(v []*ResourceTagForListAlertGroupOutput) *DataForListAlertGroupOutput {
	s.ResourceTags = v
	return s
}

// SetResourceType sets the ResourceType field's value.
func (s *DataForListAlertGroupOutput) SetResourceType(v string) *DataForListAlertGroupOutput {
	s.ResourceType = &v
	return s
}

// SetRuleId sets the RuleId field's value.
func (s *DataForListAlertGroupOutput) SetRuleId(v string) *DataForListAlertGroupOutput {
	s.RuleId = &v
	return s
}

// SetRuleName sets the RuleName field's value.
func (s *DataForListAlertGroupOutput) SetRuleName(v string) *DataForListAlertGroupOutput {
	s.RuleName = &v
	return s
}

// SetRuleStatus sets the RuleStatus field's value.
func (s *DataForListAlertGroupOutput) SetRuleStatus(v string) *DataForListAlertGroupOutput {
	s.RuleStatus = &v
	return s
}

// SetRuleTriggerCondition sets the RuleTriggerCondition field's value.
func (s *DataForListAlertGroupOutput) SetRuleTriggerCondition(v string) *DataForListAlertGroupOutput {
	s.RuleTriggerCondition = &v
	return s
}

// SetSendAlertCount sets the SendAlertCount field's value.
func (s *DataForListAlertGroupOutput) SetSendAlertCount(v int64) *DataForListAlertGroupOutput {
	s.SendAlertCount = &v
	return s
}

// SetStartAt sets the StartAt field's value.
func (s *DataForListAlertGroupOutput) SetStartAt(v string) *DataForListAlertGroupOutput {
	s.StartAt = &v
	return s
}

// SetSubNamespace sets the SubNamespace field's value.
func (s *DataForListAlertGroupOutput) SetSubNamespace(v string) *DataForListAlertGroupOutput {
	s.SubNamespace = &v
	return s
}

type ListAlertGroupInput struct {
	_ struct{} `type:"structure"`

	AlertStates []*string `type:"list"`

	EndAt *int64 `type:"integer"`

	Levels []*string `type:"list"`

	Namespaces []*string `type:"list"`

	PageNumber *int64 `type:"integer"`

	PageSize *int64 `type:"integer"`

	ResourceId *string `type:"string"`

	RuleIds []*string `type:"list"`

	RuleName *string `type:"string"`

	StartAt *int64 `type:"integer"`
}

// String returns the string representation
func (s ListAlertGroupInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListAlertGroupInput) GoString() string {
	return s.String()
}

// SetAlertStates sets the AlertStates field's value.
func (s *ListAlertGroupInput) SetAlertStates(v []*string) *ListAlertGroupInput {
	s.AlertStates = v
	return s
}

// SetEndAt sets the EndAt field's value.
func (s *ListAlertGroupInput) SetEndAt(v int64) *ListAlertGroupInput {
	s.EndAt = &v
	return s
}

// SetLevels sets the Levels field's value.
func (s *ListAlertGroupInput) SetLevels(v []*string) *ListAlertGroupInput {
	s.Levels = v
	return s
}

// SetNamespaces sets the Namespaces field's value.
func (s *ListAlertGroupInput) SetNamespaces(v []*string) *ListAlertGroupInput {
	s.Namespaces = v
	return s
}

// SetPageNumber sets the PageNumber field's value.
func (s *ListAlertGroupInput) SetPageNumber(v int64) *ListAlertGroupInput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *ListAlertGroupInput) SetPageSize(v int64) *ListAlertGroupInput {
	s.PageSize = &v
	return s
}

// SetResourceId sets the ResourceId field's value.
func (s *ListAlertGroupInput) SetResourceId(v string) *ListAlertGroupInput {
	s.ResourceId = &v
	return s
}

// SetRuleIds sets the RuleIds field's value.
func (s *ListAlertGroupInput) SetRuleIds(v []*string) *ListAlertGroupInput {
	s.RuleIds = v
	return s
}

// SetRuleName sets the RuleName field's value.
func (s *ListAlertGroupInput) SetRuleName(v string) *ListAlertGroupInput {
	s.RuleName = &v
	return s
}

// SetStartAt sets the StartAt field's value.
func (s *ListAlertGroupInput) SetStartAt(v int64) *ListAlertGroupInput {
	s.StartAt = &v
	return s
}

type ListAlertGroupOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	Asc *bool `type:"boolean"`

	Data []*DataForListAlertGroupOutput `type:"list"`

	OrderBy *string `type:"string"`

	PageNumber *int64 `type:"integer"`

	PageSize *int64 `type:"integer"`

	TotalCount *int64 `type:"integer"`
}

// String returns the string representation
func (s ListAlertGroupOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListAlertGroupOutput) GoString() string {
	return s.String()
}

// SetAsc sets the Asc field's value.
func (s *ListAlertGroupOutput) SetAsc(v bool) *ListAlertGroupOutput {
	s.Asc = &v
	return s
}

// SetData sets the Data field's value.
func (s *ListAlertGroupOutput) SetData(v []*DataForListAlertGroupOutput) *ListAlertGroupOutput {
	s.Data = v
	return s
}

// SetOrderBy sets the OrderBy field's value.
func (s *ListAlertGroupOutput) SetOrderBy(v string) *ListAlertGroupOutput {
	s.OrderBy = &v
	return s
}

// SetPageNumber sets the PageNumber field's value.
func (s *ListAlertGroupOutput) SetPageNumber(v int64) *ListAlertGroupOutput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *ListAlertGroupOutput) SetPageSize(v int64) *ListAlertGroupOutput {
	s.PageSize = &v
	return s
}

// SetTotalCount sets the TotalCount field's value.
func (s *ListAlertGroupOutput) SetTotalCount(v int64) *ListAlertGroupOutput {
	s.TotalCount = &v
	return s
}

type ResourceTagForListAlertGroupOutput struct {
	_ struct{} `type:"structure"`

	Name *string `type:"string"`

	Value *string `type:"string"`
}

// String returns the string representation
func (s ResourceTagForListAlertGroupOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ResourceTagForListAlertGroupOutput) GoString() string {
	return s.String()
}

// SetName sets the Name field's value.
func (s *ResourceTagForListAlertGroupOutput) SetName(v string) *ResourceTagForListAlertGroupOutput {
	s.Name = &v
	return s
}

// SetValue sets the Value field's value.
func (s *ResourceTagForListAlertGroupOutput) SetValue(v string) *ResourceTagForListAlertGroupOutput {
	s.Value = &v
	return s
}
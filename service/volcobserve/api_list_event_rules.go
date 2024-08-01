// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package volcobserve

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opListEventRulesCommon = "ListEventRules"

// ListEventRulesCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ListEventRulesCommon operation. The "output" return
// value will be populated with the ListEventRulesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListEventRulesCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListEventRulesCommon Send returns without error.
//
// See ListEventRulesCommon for more information on using the ListEventRulesCommon
// API call, and error handling.
//
//    // Example sending a request using the ListEventRulesCommonRequest method.
//    req, resp := client.ListEventRulesCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VOLCOBSERVE) ListEventRulesCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opListEventRulesCommon,
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

// ListEventRulesCommon API operation for VOLC_OBSERVE.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VOLC_OBSERVE's
// API operation ListEventRulesCommon for usage and error information.
func (c *VOLCOBSERVE) ListEventRulesCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ListEventRulesCommonRequest(input)
	return out, req.Send()
}

// ListEventRulesCommonWithContext is the same as ListEventRulesCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ListEventRulesCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VOLCOBSERVE) ListEventRulesCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ListEventRulesCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opListEventRules = "ListEventRules"

// ListEventRulesRequest generates a "volcengine/request.Request" representing the
// client's request for the ListEventRules operation. The "output" return
// value will be populated with the ListEventRulesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListEventRulesCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListEventRulesCommon Send returns without error.
//
// See ListEventRules for more information on using the ListEventRules
// API call, and error handling.
//
//    // Example sending a request using the ListEventRulesRequest method.
//    req, resp := client.ListEventRulesRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VOLCOBSERVE) ListEventRulesRequest(input *ListEventRulesInput) (req *request.Request, output *ListEventRulesOutput) {
	op := &request.Operation{
		Name:       opListEventRules,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListEventRulesInput{}
	}

	output = &ListEventRulesOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// ListEventRules API operation for VOLC_OBSERVE.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VOLC_OBSERVE's
// API operation ListEventRules for usage and error information.
func (c *VOLCOBSERVE) ListEventRules(input *ListEventRulesInput) (*ListEventRulesOutput, error) {
	req, out := c.ListEventRulesRequest(input)
	return out, req.Send()
}

// ListEventRulesWithContext is the same as ListEventRules with the addition of
// the ability to pass a context and additional request options.
//
// See ListEventRules for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VOLCOBSERVE) ListEventRulesWithContext(ctx volcengine.Context, input *ListEventRulesInput, opts ...request.Option) (*ListEventRulesOutput, error) {
	req, out := c.ListEventRulesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DataForListEventRulesOutput struct {
	_ struct{} `type:"structure"`

	AccountId *string `type:"string"`

	ContactGroupIds []*string `type:"list"`

	ContactMethods []*string `type:"list"`

	CreatedAt *int64 `type:"integer"`

	Description *string `type:"string"`

	EffectEndAt *string `type:"string"`

	EffectStartAt *string `type:"string"`

	EnableState *string `type:"string"`

	Endpoint *string `type:"string"`

	EventBusName *string `type:"string"`

	EventType []*string `type:"list"`

	FilterPattern map[string]*interface{} `type:"map"`

	Level *string `type:"string"`

	MessageQueue *MessageQueueForListEventRulesOutput `type:"structure"`

	Region *string `type:"string"`

	RuleId *string `type:"string"`

	RuleName *string `type:"string"`

	Source *string `type:"string"`

	TLSTarget []*TLSTargetForListEventRulesOutput `type:"list"`

	UpdatedAt *int64 `type:"integer"`

	WebhookIds []*string `type:"list"`
}

// String returns the string representation
func (s DataForListEventRulesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DataForListEventRulesOutput) GoString() string {
	return s.String()
}

// SetAccountId sets the AccountId field's value.
func (s *DataForListEventRulesOutput) SetAccountId(v string) *DataForListEventRulesOutput {
	s.AccountId = &v
	return s
}

// SetContactGroupIds sets the ContactGroupIds field's value.
func (s *DataForListEventRulesOutput) SetContactGroupIds(v []*string) *DataForListEventRulesOutput {
	s.ContactGroupIds = v
	return s
}

// SetContactMethods sets the ContactMethods field's value.
func (s *DataForListEventRulesOutput) SetContactMethods(v []*string) *DataForListEventRulesOutput {
	s.ContactMethods = v
	return s
}

// SetCreatedAt sets the CreatedAt field's value.
func (s *DataForListEventRulesOutput) SetCreatedAt(v int64) *DataForListEventRulesOutput {
	s.CreatedAt = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *DataForListEventRulesOutput) SetDescription(v string) *DataForListEventRulesOutput {
	s.Description = &v
	return s
}

// SetEffectEndAt sets the EffectEndAt field's value.
func (s *DataForListEventRulesOutput) SetEffectEndAt(v string) *DataForListEventRulesOutput {
	s.EffectEndAt = &v
	return s
}

// SetEffectStartAt sets the EffectStartAt field's value.
func (s *DataForListEventRulesOutput) SetEffectStartAt(v string) *DataForListEventRulesOutput {
	s.EffectStartAt = &v
	return s
}

// SetEnableState sets the EnableState field's value.
func (s *DataForListEventRulesOutput) SetEnableState(v string) *DataForListEventRulesOutput {
	s.EnableState = &v
	return s
}

// SetEndpoint sets the Endpoint field's value.
func (s *DataForListEventRulesOutput) SetEndpoint(v string) *DataForListEventRulesOutput {
	s.Endpoint = &v
	return s
}

// SetEventBusName sets the EventBusName field's value.
func (s *DataForListEventRulesOutput) SetEventBusName(v string) *DataForListEventRulesOutput {
	s.EventBusName = &v
	return s
}

// SetEventType sets the EventType field's value.
func (s *DataForListEventRulesOutput) SetEventType(v []*string) *DataForListEventRulesOutput {
	s.EventType = v
	return s
}

// SetFilterPattern sets the FilterPattern field's value.
func (s *DataForListEventRulesOutput) SetFilterPattern(v map[string]*interface{}) *DataForListEventRulesOutput {
	s.FilterPattern = v
	return s
}

// SetLevel sets the Level field's value.
func (s *DataForListEventRulesOutput) SetLevel(v string) *DataForListEventRulesOutput {
	s.Level = &v
	return s
}

// SetMessageQueue sets the MessageQueue field's value.
func (s *DataForListEventRulesOutput) SetMessageQueue(v *MessageQueueForListEventRulesOutput) *DataForListEventRulesOutput {
	s.MessageQueue = v
	return s
}

// SetRegion sets the Region field's value.
func (s *DataForListEventRulesOutput) SetRegion(v string) *DataForListEventRulesOutput {
	s.Region = &v
	return s
}

// SetRuleId sets the RuleId field's value.
func (s *DataForListEventRulesOutput) SetRuleId(v string) *DataForListEventRulesOutput {
	s.RuleId = &v
	return s
}

// SetRuleName sets the RuleName field's value.
func (s *DataForListEventRulesOutput) SetRuleName(v string) *DataForListEventRulesOutput {
	s.RuleName = &v
	return s
}

// SetSource sets the Source field's value.
func (s *DataForListEventRulesOutput) SetSource(v string) *DataForListEventRulesOutput {
	s.Source = &v
	return s
}

// SetTLSTarget sets the TLSTarget field's value.
func (s *DataForListEventRulesOutput) SetTLSTarget(v []*TLSTargetForListEventRulesOutput) *DataForListEventRulesOutput {
	s.TLSTarget = v
	return s
}

// SetUpdatedAt sets the UpdatedAt field's value.
func (s *DataForListEventRulesOutput) SetUpdatedAt(v int64) *DataForListEventRulesOutput {
	s.UpdatedAt = &v
	return s
}

// SetWebhookIds sets the WebhookIds field's value.
func (s *DataForListEventRulesOutput) SetWebhookIds(v []*string) *DataForListEventRulesOutput {
	s.WebhookIds = v
	return s
}

type ListEventRulesInput struct {
	_ struct{} `type:"structure"`

	Asc *bool `type:"boolean"`

	OrderBy *string `type:"string"`

	PageNumber *int64 `type:"integer"`

	PageSize *int64 `type:"integer"`

	RuleName *string `type:"string"`

	Source []*string `type:"list"`

	Status []*string `type:"list"`
}

// String returns the string representation
func (s ListEventRulesInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListEventRulesInput) GoString() string {
	return s.String()
}

// SetAsc sets the Asc field's value.
func (s *ListEventRulesInput) SetAsc(v bool) *ListEventRulesInput {
	s.Asc = &v
	return s
}

// SetOrderBy sets the OrderBy field's value.
func (s *ListEventRulesInput) SetOrderBy(v string) *ListEventRulesInput {
	s.OrderBy = &v
	return s
}

// SetPageNumber sets the PageNumber field's value.
func (s *ListEventRulesInput) SetPageNumber(v int64) *ListEventRulesInput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *ListEventRulesInput) SetPageSize(v int64) *ListEventRulesInput {
	s.PageSize = &v
	return s
}

// SetRuleName sets the RuleName field's value.
func (s *ListEventRulesInput) SetRuleName(v string) *ListEventRulesInput {
	s.RuleName = &v
	return s
}

// SetSource sets the Source field's value.
func (s *ListEventRulesInput) SetSource(v []*string) *ListEventRulesInput {
	s.Source = v
	return s
}

// SetStatus sets the Status field's value.
func (s *ListEventRulesInput) SetStatus(v []*string) *ListEventRulesInput {
	s.Status = v
	return s
}

type ListEventRulesOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	Asc *bool `type:"boolean"`

	Data []*DataForListEventRulesOutput `type:"list"`

	OrderBy *string `type:"string"`

	PageNumber *int64 `type:"integer"`

	PageSize *int64 `type:"integer"`

	TotalCount *int64 `type:"integer"`
}

// String returns the string representation
func (s ListEventRulesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListEventRulesOutput) GoString() string {
	return s.String()
}

// SetAsc sets the Asc field's value.
func (s *ListEventRulesOutput) SetAsc(v bool) *ListEventRulesOutput {
	s.Asc = &v
	return s
}

// SetData sets the Data field's value.
func (s *ListEventRulesOutput) SetData(v []*DataForListEventRulesOutput) *ListEventRulesOutput {
	s.Data = v
	return s
}

// SetOrderBy sets the OrderBy field's value.
func (s *ListEventRulesOutput) SetOrderBy(v string) *ListEventRulesOutput {
	s.OrderBy = &v
	return s
}

// SetPageNumber sets the PageNumber field's value.
func (s *ListEventRulesOutput) SetPageNumber(v int64) *ListEventRulesOutput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *ListEventRulesOutput) SetPageSize(v int64) *ListEventRulesOutput {
	s.PageSize = &v
	return s
}

// SetTotalCount sets the TotalCount field's value.
func (s *ListEventRulesOutput) SetTotalCount(v int64) *ListEventRulesOutput {
	s.TotalCount = &v
	return s
}

type MessageQueueForListEventRulesOutput struct {
	_ struct{} `type:"structure"`

	InstanceId *string `type:"string"`

	Region *string `type:"string"`

	Topic *string `type:"string"`

	Type *string `type:"string"`

	VpcId *string `type:"string"`
}

// String returns the string representation
func (s MessageQueueForListEventRulesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s MessageQueueForListEventRulesOutput) GoString() string {
	return s.String()
}

// SetInstanceId sets the InstanceId field's value.
func (s *MessageQueueForListEventRulesOutput) SetInstanceId(v string) *MessageQueueForListEventRulesOutput {
	s.InstanceId = &v
	return s
}

// SetRegion sets the Region field's value.
func (s *MessageQueueForListEventRulesOutput) SetRegion(v string) *MessageQueueForListEventRulesOutput {
	s.Region = &v
	return s
}

// SetTopic sets the Topic field's value.
func (s *MessageQueueForListEventRulesOutput) SetTopic(v string) *MessageQueueForListEventRulesOutput {
	s.Topic = &v
	return s
}

// SetType sets the Type field's value.
func (s *MessageQueueForListEventRulesOutput) SetType(v string) *MessageQueueForListEventRulesOutput {
	s.Type = &v
	return s
}

// SetVpcId sets the VpcId field's value.
func (s *MessageQueueForListEventRulesOutput) SetVpcId(v string) *MessageQueueForListEventRulesOutput {
	s.VpcId = &v
	return s
}

type TLSTargetForListEventRulesOutput struct {
	_ struct{} `type:"structure"`

	ProjectId *string `type:"string"`

	ProjectName *string `type:"string"`

	RegionNameCN *string `type:"string"`

	RegionNameEN *string `type:"string"`

	TopicId *string `type:"string"`
}

// String returns the string representation
func (s TLSTargetForListEventRulesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s TLSTargetForListEventRulesOutput) GoString() string {
	return s.String()
}

// SetProjectId sets the ProjectId field's value.
func (s *TLSTargetForListEventRulesOutput) SetProjectId(v string) *TLSTargetForListEventRulesOutput {
	s.ProjectId = &v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *TLSTargetForListEventRulesOutput) SetProjectName(v string) *TLSTargetForListEventRulesOutput {
	s.ProjectName = &v
	return s
}

// SetRegionNameCN sets the RegionNameCN field's value.
func (s *TLSTargetForListEventRulesOutput) SetRegionNameCN(v string) *TLSTargetForListEventRulesOutput {
	s.RegionNameCN = &v
	return s
}

// SetRegionNameEN sets the RegionNameEN field's value.
func (s *TLSTargetForListEventRulesOutput) SetRegionNameEN(v string) *TLSTargetForListEventRulesOutput {
	s.RegionNameEN = &v
	return s
}

// SetTopicId sets the TopicId field's value.
func (s *TLSTargetForListEventRulesOutput) SetTopicId(v string) *TLSTargetForListEventRulesOutput {
	s.TopicId = &v
	return s
}

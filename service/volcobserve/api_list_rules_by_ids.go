// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package volcobserve

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opListRulesByIdsCommon = "ListRulesByIds"

// ListRulesByIdsCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ListRulesByIdsCommon operation. The "output" return
// value will be populated with the ListRulesByIdsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListRulesByIdsCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListRulesByIdsCommon Send returns without error.
//
// See ListRulesByIdsCommon for more information on using the ListRulesByIdsCommon
// API call, and error handling.
//
//    // Example sending a request using the ListRulesByIdsCommonRequest method.
//    req, resp := client.ListRulesByIdsCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VOLCOBSERVE) ListRulesByIdsCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opListRulesByIdsCommon,
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

// ListRulesByIdsCommon API operation for VOLC_OBSERVE.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VOLC_OBSERVE's
// API operation ListRulesByIdsCommon for usage and error information.
func (c *VOLCOBSERVE) ListRulesByIdsCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ListRulesByIdsCommonRequest(input)
	return out, req.Send()
}

// ListRulesByIdsCommonWithContext is the same as ListRulesByIdsCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ListRulesByIdsCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VOLCOBSERVE) ListRulesByIdsCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ListRulesByIdsCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opListRulesByIds = "ListRulesByIds"

// ListRulesByIdsRequest generates a "volcengine/request.Request" representing the
// client's request for the ListRulesByIds operation. The "output" return
// value will be populated with the ListRulesByIdsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListRulesByIdsCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListRulesByIdsCommon Send returns without error.
//
// See ListRulesByIds for more information on using the ListRulesByIds
// API call, and error handling.
//
//    // Example sending a request using the ListRulesByIdsRequest method.
//    req, resp := client.ListRulesByIdsRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VOLCOBSERVE) ListRulesByIdsRequest(input *ListRulesByIdsInput) (req *request.Request, output *ListRulesByIdsOutput) {
	op := &request.Operation{
		Name:       opListRulesByIds,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListRulesByIdsInput{}
	}

	output = &ListRulesByIdsOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// ListRulesByIds API operation for VOLC_OBSERVE.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VOLC_OBSERVE's
// API operation ListRulesByIds for usage and error information.
func (c *VOLCOBSERVE) ListRulesByIds(input *ListRulesByIdsInput) (*ListRulesByIdsOutput, error) {
	req, out := c.ListRulesByIdsRequest(input)
	return out, req.Send()
}

// ListRulesByIdsWithContext is the same as ListRulesByIds with the addition of
// the ability to pass a context and additional request options.
//
// See ListRulesByIds for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VOLCOBSERVE) ListRulesByIdsWithContext(ctx volcengine.Context, input *ListRulesByIdsInput, opts ...request.Option) (*ListRulesByIdsOutput, error) {
	req, out := c.ListRulesByIdsRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ConditionForListRulesByIdsOutput struct {
	_ struct{} `type:"structure"`

	ComparisonOperator *string `type:"string"`

	MetricName *string `type:"string"`

	MetricUnit *string `type:"string"`

	Statistics *string `type:"string"`

	Threshold *string `type:"string"`
}

// String returns the string representation
func (s ConditionForListRulesByIdsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ConditionForListRulesByIdsOutput) GoString() string {
	return s.String()
}

// SetComparisonOperator sets the ComparisonOperator field's value.
func (s *ConditionForListRulesByIdsOutput) SetComparisonOperator(v string) *ConditionForListRulesByIdsOutput {
	s.ComparisonOperator = &v
	return s
}

// SetMetricName sets the MetricName field's value.
func (s *ConditionForListRulesByIdsOutput) SetMetricName(v string) *ConditionForListRulesByIdsOutput {
	s.MetricName = &v
	return s
}

// SetMetricUnit sets the MetricUnit field's value.
func (s *ConditionForListRulesByIdsOutput) SetMetricUnit(v string) *ConditionForListRulesByIdsOutput {
	s.MetricUnit = &v
	return s
}

// SetStatistics sets the Statistics field's value.
func (s *ConditionForListRulesByIdsOutput) SetStatistics(v string) *ConditionForListRulesByIdsOutput {
	s.Statistics = &v
	return s
}

// SetThreshold sets the Threshold field's value.
func (s *ConditionForListRulesByIdsOutput) SetThreshold(v string) *ConditionForListRulesByIdsOutput {
	s.Threshold = &v
	return s
}

type DataForListRulesByIdsOutput struct {
	_ struct{} `type:"structure"`

	AlertMethods []*string `type:"list"`

	AlertState *string `type:"string"`

	ConditionOperator *string `type:"string"`

	Conditions []*ConditionForListRulesByIdsOutput `type:"list"`

	ContactGroupIds []*string `type:"list"`

	CreatedAt *string `type:"string"`

	Description *string `type:"string"`

	EffectEndAt *string `type:"string"`

	EffectStartAt *string `type:"string"`

	EnableState *string `type:"string"`

	EvaluationCount *int64 `type:"integer"`

	Id *string `type:"string"`

	Level *string `type:"string"`

	MultipleConditions *bool `type:"boolean"`

	Namespace *string `type:"string"`

	OriginalDimensions map[string][]*string `type:"map"`

	RecoveryNotify *RecoveryNotifyForListRulesByIdsOutput `type:"structure"`

	Regions []*string `type:"list"`

	RuleName *string `type:"string"`

	SilenceTime *int64 `type:"integer"`

	SubNamespace *string `type:"string"`

	UpdatedAt *string `type:"string"`

	WebHook *string `type:"string"`

	WebhookIds []*string `type:"list"`
}

// String returns the string representation
func (s DataForListRulesByIdsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DataForListRulesByIdsOutput) GoString() string {
	return s.String()
}

// SetAlertMethods sets the AlertMethods field's value.
func (s *DataForListRulesByIdsOutput) SetAlertMethods(v []*string) *DataForListRulesByIdsOutput {
	s.AlertMethods = v
	return s
}

// SetAlertState sets the AlertState field's value.
func (s *DataForListRulesByIdsOutput) SetAlertState(v string) *DataForListRulesByIdsOutput {
	s.AlertState = &v
	return s
}

// SetConditionOperator sets the ConditionOperator field's value.
func (s *DataForListRulesByIdsOutput) SetConditionOperator(v string) *DataForListRulesByIdsOutput {
	s.ConditionOperator = &v
	return s
}

// SetConditions sets the Conditions field's value.
func (s *DataForListRulesByIdsOutput) SetConditions(v []*ConditionForListRulesByIdsOutput) *DataForListRulesByIdsOutput {
	s.Conditions = v
	return s
}

// SetContactGroupIds sets the ContactGroupIds field's value.
func (s *DataForListRulesByIdsOutput) SetContactGroupIds(v []*string) *DataForListRulesByIdsOutput {
	s.ContactGroupIds = v
	return s
}

// SetCreatedAt sets the CreatedAt field's value.
func (s *DataForListRulesByIdsOutput) SetCreatedAt(v string) *DataForListRulesByIdsOutput {
	s.CreatedAt = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *DataForListRulesByIdsOutput) SetDescription(v string) *DataForListRulesByIdsOutput {
	s.Description = &v
	return s
}

// SetEffectEndAt sets the EffectEndAt field's value.
func (s *DataForListRulesByIdsOutput) SetEffectEndAt(v string) *DataForListRulesByIdsOutput {
	s.EffectEndAt = &v
	return s
}

// SetEffectStartAt sets the EffectStartAt field's value.
func (s *DataForListRulesByIdsOutput) SetEffectStartAt(v string) *DataForListRulesByIdsOutput {
	s.EffectStartAt = &v
	return s
}

// SetEnableState sets the EnableState field's value.
func (s *DataForListRulesByIdsOutput) SetEnableState(v string) *DataForListRulesByIdsOutput {
	s.EnableState = &v
	return s
}

// SetEvaluationCount sets the EvaluationCount field's value.
func (s *DataForListRulesByIdsOutput) SetEvaluationCount(v int64) *DataForListRulesByIdsOutput {
	s.EvaluationCount = &v
	return s
}

// SetId sets the Id field's value.
func (s *DataForListRulesByIdsOutput) SetId(v string) *DataForListRulesByIdsOutput {
	s.Id = &v
	return s
}

// SetLevel sets the Level field's value.
func (s *DataForListRulesByIdsOutput) SetLevel(v string) *DataForListRulesByIdsOutput {
	s.Level = &v
	return s
}

// SetMultipleConditions sets the MultipleConditions field's value.
func (s *DataForListRulesByIdsOutput) SetMultipleConditions(v bool) *DataForListRulesByIdsOutput {
	s.MultipleConditions = &v
	return s
}

// SetNamespace sets the Namespace field's value.
func (s *DataForListRulesByIdsOutput) SetNamespace(v string) *DataForListRulesByIdsOutput {
	s.Namespace = &v
	return s
}

// SetOriginalDimensions sets the OriginalDimensions field's value.
func (s *DataForListRulesByIdsOutput) SetOriginalDimensions(v map[string][]*string) *DataForListRulesByIdsOutput {
	s.OriginalDimensions = v
	return s
}

// SetRecoveryNotify sets the RecoveryNotify field's value.
func (s *DataForListRulesByIdsOutput) SetRecoveryNotify(v *RecoveryNotifyForListRulesByIdsOutput) *DataForListRulesByIdsOutput {
	s.RecoveryNotify = v
	return s
}

// SetRegions sets the Regions field's value.
func (s *DataForListRulesByIdsOutput) SetRegions(v []*string) *DataForListRulesByIdsOutput {
	s.Regions = v
	return s
}

// SetRuleName sets the RuleName field's value.
func (s *DataForListRulesByIdsOutput) SetRuleName(v string) *DataForListRulesByIdsOutput {
	s.RuleName = &v
	return s
}

// SetSilenceTime sets the SilenceTime field's value.
func (s *DataForListRulesByIdsOutput) SetSilenceTime(v int64) *DataForListRulesByIdsOutput {
	s.SilenceTime = &v
	return s
}

// SetSubNamespace sets the SubNamespace field's value.
func (s *DataForListRulesByIdsOutput) SetSubNamespace(v string) *DataForListRulesByIdsOutput {
	s.SubNamespace = &v
	return s
}

// SetUpdatedAt sets the UpdatedAt field's value.
func (s *DataForListRulesByIdsOutput) SetUpdatedAt(v string) *DataForListRulesByIdsOutput {
	s.UpdatedAt = &v
	return s
}

// SetWebHook sets the WebHook field's value.
func (s *DataForListRulesByIdsOutput) SetWebHook(v string) *DataForListRulesByIdsOutput {
	s.WebHook = &v
	return s
}

// SetWebhookIds sets the WebhookIds field's value.
func (s *DataForListRulesByIdsOutput) SetWebhookIds(v []*string) *DataForListRulesByIdsOutput {
	s.WebhookIds = v
	return s
}

type ListRulesByIdsInput struct {
	_ struct{} `type:"structure"`

	Ids []*string `type:"list"`
}

// String returns the string representation
func (s ListRulesByIdsInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListRulesByIdsInput) GoString() string {
	return s.String()
}

// SetIds sets the Ids field's value.
func (s *ListRulesByIdsInput) SetIds(v []*string) *ListRulesByIdsInput {
	s.Ids = v
	return s
}

type ListRulesByIdsOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	Data []*DataForListRulesByIdsOutput `type:"list"`
}

// String returns the string representation
func (s ListRulesByIdsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListRulesByIdsOutput) GoString() string {
	return s.String()
}

// SetData sets the Data field's value.
func (s *ListRulesByIdsOutput) SetData(v []*DataForListRulesByIdsOutput) *ListRulesByIdsOutput {
	s.Data = v
	return s
}

type RecoveryNotifyForListRulesByIdsOutput struct {
	_ struct{} `type:"structure"`

	Enable *bool `type:"boolean"`
}

// String returns the string representation
func (s RecoveryNotifyForListRulesByIdsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s RecoveryNotifyForListRulesByIdsOutput) GoString() string {
	return s.String()
}

// SetEnable sets the Enable field's value.
func (s *RecoveryNotifyForListRulesByIdsOutput) SetEnable(v bool) *RecoveryNotifyForListRulesByIdsOutput {
	s.Enable = &v
	return s
}
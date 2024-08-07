// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package volcobserve

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opListRulesCommon = "ListRules"

// ListRulesCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ListRulesCommon operation. The "output" return
// value will be populated with the ListRulesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListRulesCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListRulesCommon Send returns without error.
//
// See ListRulesCommon for more information on using the ListRulesCommon
// API call, and error handling.
//
//    // Example sending a request using the ListRulesCommonRequest method.
//    req, resp := client.ListRulesCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VOLCOBSERVE) ListRulesCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opListRulesCommon,
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

// ListRulesCommon API operation for VOLC_OBSERVE.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VOLC_OBSERVE's
// API operation ListRulesCommon for usage and error information.
func (c *VOLCOBSERVE) ListRulesCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ListRulesCommonRequest(input)
	return out, req.Send()
}

// ListRulesCommonWithContext is the same as ListRulesCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ListRulesCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VOLCOBSERVE) ListRulesCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ListRulesCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opListRules = "ListRules"

// ListRulesRequest generates a "volcengine/request.Request" representing the
// client's request for the ListRules operation. The "output" return
// value will be populated with the ListRulesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListRulesCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListRulesCommon Send returns without error.
//
// See ListRules for more information on using the ListRules
// API call, and error handling.
//
//    // Example sending a request using the ListRulesRequest method.
//    req, resp := client.ListRulesRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VOLCOBSERVE) ListRulesRequest(input *ListRulesInput) (req *request.Request, output *ListRulesOutput) {
	op := &request.Operation{
		Name:       opListRules,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListRulesInput{}
	}

	output = &ListRulesOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// ListRules API operation for VOLC_OBSERVE.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VOLC_OBSERVE's
// API operation ListRules for usage and error information.
func (c *VOLCOBSERVE) ListRules(input *ListRulesInput) (*ListRulesOutput, error) {
	req, out := c.ListRulesRequest(input)
	return out, req.Send()
}

// ListRulesWithContext is the same as ListRules with the addition of
// the ability to pass a context and additional request options.
//
// See ListRules for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VOLCOBSERVE) ListRulesWithContext(ctx volcengine.Context, input *ListRulesInput, opts ...request.Option) (*ListRulesOutput, error) {
	req, out := c.ListRulesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ConditionForListRulesOutput struct {
	_ struct{} `type:"structure"`

	ComparisonOperator *string `type:"string"`

	MetricName *string `type:"string"`

	MetricUnit *string `type:"string"`

	Statistics *string `type:"string"`

	Threshold *string `type:"string"`
}

// String returns the string representation
func (s ConditionForListRulesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ConditionForListRulesOutput) GoString() string {
	return s.String()
}

// SetComparisonOperator sets the ComparisonOperator field's value.
func (s *ConditionForListRulesOutput) SetComparisonOperator(v string) *ConditionForListRulesOutput {
	s.ComparisonOperator = &v
	return s
}

// SetMetricName sets the MetricName field's value.
func (s *ConditionForListRulesOutput) SetMetricName(v string) *ConditionForListRulesOutput {
	s.MetricName = &v
	return s
}

// SetMetricUnit sets the MetricUnit field's value.
func (s *ConditionForListRulesOutput) SetMetricUnit(v string) *ConditionForListRulesOutput {
	s.MetricUnit = &v
	return s
}

// SetStatistics sets the Statistics field's value.
func (s *ConditionForListRulesOutput) SetStatistics(v string) *ConditionForListRulesOutput {
	s.Statistics = &v
	return s
}

// SetThreshold sets the Threshold field's value.
func (s *ConditionForListRulesOutput) SetThreshold(v string) *ConditionForListRulesOutput {
	s.Threshold = &v
	return s
}

type DataForListRulesOutput struct {
	_ struct{} `type:"structure"`

	AlertMethods []*string `type:"list"`

	AlertState *string `type:"string"`

	ConditionOperator *string `type:"string"`

	Conditions []*ConditionForListRulesOutput `type:"list"`

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

	RecoveryNotify *RecoveryNotifyForListRulesOutput `type:"structure"`

	Regions []*string `type:"list"`

	RuleName *string `type:"string"`

	SilenceTime *int64 `type:"integer"`

	SubNamespace *string `type:"string"`

	UpdatedAt *string `type:"string"`

	WebHook *string `type:"string"`

	WebhookIds []*string `type:"list"`
}

// String returns the string representation
func (s DataForListRulesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DataForListRulesOutput) GoString() string {
	return s.String()
}

// SetAlertMethods sets the AlertMethods field's value.
func (s *DataForListRulesOutput) SetAlertMethods(v []*string) *DataForListRulesOutput {
	s.AlertMethods = v
	return s
}

// SetAlertState sets the AlertState field's value.
func (s *DataForListRulesOutput) SetAlertState(v string) *DataForListRulesOutput {
	s.AlertState = &v
	return s
}

// SetConditionOperator sets the ConditionOperator field's value.
func (s *DataForListRulesOutput) SetConditionOperator(v string) *DataForListRulesOutput {
	s.ConditionOperator = &v
	return s
}

// SetConditions sets the Conditions field's value.
func (s *DataForListRulesOutput) SetConditions(v []*ConditionForListRulesOutput) *DataForListRulesOutput {
	s.Conditions = v
	return s
}

// SetContactGroupIds sets the ContactGroupIds field's value.
func (s *DataForListRulesOutput) SetContactGroupIds(v []*string) *DataForListRulesOutput {
	s.ContactGroupIds = v
	return s
}

// SetCreatedAt sets the CreatedAt field's value.
func (s *DataForListRulesOutput) SetCreatedAt(v string) *DataForListRulesOutput {
	s.CreatedAt = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *DataForListRulesOutput) SetDescription(v string) *DataForListRulesOutput {
	s.Description = &v
	return s
}

// SetEffectEndAt sets the EffectEndAt field's value.
func (s *DataForListRulesOutput) SetEffectEndAt(v string) *DataForListRulesOutput {
	s.EffectEndAt = &v
	return s
}

// SetEffectStartAt sets the EffectStartAt field's value.
func (s *DataForListRulesOutput) SetEffectStartAt(v string) *DataForListRulesOutput {
	s.EffectStartAt = &v
	return s
}

// SetEnableState sets the EnableState field's value.
func (s *DataForListRulesOutput) SetEnableState(v string) *DataForListRulesOutput {
	s.EnableState = &v
	return s
}

// SetEvaluationCount sets the EvaluationCount field's value.
func (s *DataForListRulesOutput) SetEvaluationCount(v int64) *DataForListRulesOutput {
	s.EvaluationCount = &v
	return s
}

// SetId sets the Id field's value.
func (s *DataForListRulesOutput) SetId(v string) *DataForListRulesOutput {
	s.Id = &v
	return s
}

// SetLevel sets the Level field's value.
func (s *DataForListRulesOutput) SetLevel(v string) *DataForListRulesOutput {
	s.Level = &v
	return s
}

// SetMultipleConditions sets the MultipleConditions field's value.
func (s *DataForListRulesOutput) SetMultipleConditions(v bool) *DataForListRulesOutput {
	s.MultipleConditions = &v
	return s
}

// SetNamespace sets the Namespace field's value.
func (s *DataForListRulesOutput) SetNamespace(v string) *DataForListRulesOutput {
	s.Namespace = &v
	return s
}

// SetOriginalDimensions sets the OriginalDimensions field's value.
func (s *DataForListRulesOutput) SetOriginalDimensions(v map[string][]*string) *DataForListRulesOutput {
	s.OriginalDimensions = v
	return s
}

// SetRecoveryNotify sets the RecoveryNotify field's value.
func (s *DataForListRulesOutput) SetRecoveryNotify(v *RecoveryNotifyForListRulesOutput) *DataForListRulesOutput {
	s.RecoveryNotify = v
	return s
}

// SetRegions sets the Regions field's value.
func (s *DataForListRulesOutput) SetRegions(v []*string) *DataForListRulesOutput {
	s.Regions = v
	return s
}

// SetRuleName sets the RuleName field's value.
func (s *DataForListRulesOutput) SetRuleName(v string) *DataForListRulesOutput {
	s.RuleName = &v
	return s
}

// SetSilenceTime sets the SilenceTime field's value.
func (s *DataForListRulesOutput) SetSilenceTime(v int64) *DataForListRulesOutput {
	s.SilenceTime = &v
	return s
}

// SetSubNamespace sets the SubNamespace field's value.
func (s *DataForListRulesOutput) SetSubNamespace(v string) *DataForListRulesOutput {
	s.SubNamespace = &v
	return s
}

// SetUpdatedAt sets the UpdatedAt field's value.
func (s *DataForListRulesOutput) SetUpdatedAt(v string) *DataForListRulesOutput {
	s.UpdatedAt = &v
	return s
}

// SetWebHook sets the WebHook field's value.
func (s *DataForListRulesOutput) SetWebHook(v string) *DataForListRulesOutput {
	s.WebHook = &v
	return s
}

// SetWebhookIds sets the WebhookIds field's value.
func (s *DataForListRulesOutput) SetWebhookIds(v []*string) *DataForListRulesOutput {
	s.WebhookIds = v
	return s
}

type ListRulesInput struct {
	_ struct{} `type:"structure"`

	AlertState []*string `type:"list"`

	EnableState []*string `type:"list"`

	Level []*string `type:"list"`

	Namespace []*string `type:"list"`

	PageNumber *int64 `type:"integer"`

	PageSize *int64 `type:"integer"`

	RuleName *string `type:"string"`
}

// String returns the string representation
func (s ListRulesInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListRulesInput) GoString() string {
	return s.String()
}

// SetAlertState sets the AlertState field's value.
func (s *ListRulesInput) SetAlertState(v []*string) *ListRulesInput {
	s.AlertState = v
	return s
}

// SetEnableState sets the EnableState field's value.
func (s *ListRulesInput) SetEnableState(v []*string) *ListRulesInput {
	s.EnableState = v
	return s
}

// SetLevel sets the Level field's value.
func (s *ListRulesInput) SetLevel(v []*string) *ListRulesInput {
	s.Level = v
	return s
}

// SetNamespace sets the Namespace field's value.
func (s *ListRulesInput) SetNamespace(v []*string) *ListRulesInput {
	s.Namespace = v
	return s
}

// SetPageNumber sets the PageNumber field's value.
func (s *ListRulesInput) SetPageNumber(v int64) *ListRulesInput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *ListRulesInput) SetPageSize(v int64) *ListRulesInput {
	s.PageSize = &v
	return s
}

// SetRuleName sets the RuleName field's value.
func (s *ListRulesInput) SetRuleName(v string) *ListRulesInput {
	s.RuleName = &v
	return s
}

type ListRulesOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	Data []*DataForListRulesOutput `type:"list"`

	PageNumber *int64 `type:"integer"`

	PageSize *int64 `type:"integer"`

	TotalCount *int64 `type:"integer"`
}

// String returns the string representation
func (s ListRulesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListRulesOutput) GoString() string {
	return s.String()
}

// SetData sets the Data field's value.
func (s *ListRulesOutput) SetData(v []*DataForListRulesOutput) *ListRulesOutput {
	s.Data = v
	return s
}

// SetPageNumber sets the PageNumber field's value.
func (s *ListRulesOutput) SetPageNumber(v int64) *ListRulesOutput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *ListRulesOutput) SetPageSize(v int64) *ListRulesOutput {
	s.PageSize = &v
	return s
}

// SetTotalCount sets the TotalCount field's value.
func (s *ListRulesOutput) SetTotalCount(v int64) *ListRulesOutput {
	s.TotalCount = &v
	return s
}

type RecoveryNotifyForListRulesOutput struct {
	_ struct{} `type:"structure"`

	Enable *bool `type:"boolean"`
}

// String returns the string representation
func (s RecoveryNotifyForListRulesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s RecoveryNotifyForListRulesOutput) GoString() string {
	return s.String()
}

// SetEnable sets the Enable field's value.
func (s *RecoveryNotifyForListRulesOutput) SetEnable(v bool) *RecoveryNotifyForListRulesOutput {
	s.Enable = &v
	return s
}

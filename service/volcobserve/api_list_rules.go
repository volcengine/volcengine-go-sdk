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
	_ struct{} `type:"structure" json:",omitempty"`

	ComparisonOperator *string `type:"string" json:",omitempty"`

	MetricName *string `type:"string" json:",omitempty"`

	MetricUnit *string `type:"string" json:",omitempty"`

	Statistics *string `type:"string" json:",omitempty"`

	Threshold *string `type:"string" json:",omitempty"`
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
	_ struct{} `type:"structure" json:",omitempty"`

	AlertMethods []*string `type:"list" json:",omitempty"`

	AlertState *string `type:"string" json:",omitempty"`

	ConditionOperator *string `type:"string" json:",omitempty"`

	Conditions []*ConditionForListRulesOutput `type:"list" json:",omitempty"`

	ContactGroupIds []*string `type:"list" json:",omitempty"`

	CreatedAt *string `type:"string" json:",omitempty"`

	Description *string `type:"string" json:",omitempty"`

	DimensionConditions *DimensionConditionsForListRulesOutput `type:"structure" json:",omitempty"`

	EffectEndAt *string `type:"string" json:",omitempty"`

	EffectStartAt *string `type:"string" json:",omitempty"`

	EnableState *string `type:"string" json:",omitempty"`

	EvaluationCount *int64 `type:"integer" json:",omitempty"`

	Id *string `type:"string" json:",omitempty"`

	Level *string `type:"string" json:",omitempty"`

	MultipleConditions *bool `type:"boolean" json:",omitempty"`

	Namespace *string `type:"string" json:",omitempty"`

	OriginalDimensions map[string][]*string `type:"map" json:",omitempty"`

	RecoveryNotify *RecoveryNotifyForListRulesOutput `type:"structure" json:",omitempty"`

	Regions []*string `type:"list" json:",omitempty"`

	RuleName *string `type:"string" json:",omitempty"`

	RuleType *string `type:"string" json:",omitempty"`

	SilenceTime *int64 `type:"integer" json:",omitempty"`

	SubNamespace *string `type:"string" json:",omitempty"`

	UpdatedAt *string `type:"string" json:",omitempty"`

	WebHook *string `type:"string" json:",omitempty"`

	WebhookIds []*string `type:"list" json:",omitempty"`
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

// SetDimensionConditions sets the DimensionConditions field's value.
func (s *DataForListRulesOutput) SetDimensionConditions(v *DimensionConditionsForListRulesOutput) *DataForListRulesOutput {
	s.DimensionConditions = v
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

// SetRuleType sets the RuleType field's value.
func (s *DataForListRulesOutput) SetRuleType(v string) *DataForListRulesOutput {
	s.RuleType = &v
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

type DimensionConditionsForListRulesOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	MetaCondition *MetaConditionForListRulesOutput `type:"structure" json:",omitempty"`

	ProjectCondition *ProjectConditionForListRulesOutput `type:"structure" json:",omitempty"`

	TagCondition *TagConditionForListRulesOutput `type:"structure" json:",omitempty"`

	Type *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s DimensionConditionsForListRulesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DimensionConditionsForListRulesOutput) GoString() string {
	return s.String()
}

// SetMetaCondition sets the MetaCondition field's value.
func (s *DimensionConditionsForListRulesOutput) SetMetaCondition(v *MetaConditionForListRulesOutput) *DimensionConditionsForListRulesOutput {
	s.MetaCondition = v
	return s
}

// SetProjectCondition sets the ProjectCondition field's value.
func (s *DimensionConditionsForListRulesOutput) SetProjectCondition(v *ProjectConditionForListRulesOutput) *DimensionConditionsForListRulesOutput {
	s.ProjectCondition = v
	return s
}

// SetTagCondition sets the TagCondition field's value.
func (s *DimensionConditionsForListRulesOutput) SetTagCondition(v *TagConditionForListRulesOutput) *DimensionConditionsForListRulesOutput {
	s.TagCondition = v
	return s
}

// SetType sets the Type field's value.
func (s *DimensionConditionsForListRulesOutput) SetType(v string) *DimensionConditionsForListRulesOutput {
	s.Type = &v
	return s
}

type ListRulesInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	AlertState []*string `type:"list" json:",omitempty"`

	EnableState []*string `type:"list" json:",omitempty"`

	Level []*string `type:"list" json:",omitempty"`

	Namespace []*string `type:"list" json:",omitempty"`

	PageNumber *int64 `type:"integer" json:",omitempty"`

	PageSize *int64 `type:"integer" json:",omitempty"`

	RuleName *string `type:"string" json:",omitempty"`
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
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	Data []*DataForListRulesOutput `type:"list" json:",omitempty"`

	PageNumber *int64 `type:"integer" json:",omitempty"`

	PageSize *int64 `type:"integer" json:",omitempty"`

	TotalCount *int64 `type:"integer" json:",omitempty"`
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

type MetaConditionForListRulesOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	AllDimensions *bool `type:"boolean" json:",omitempty"`

	Condition *string `type:"string" json:",omitempty"`

	Metas []*MetaForListRulesOutput `type:"list" json:",omitempty"`
}

// String returns the string representation
func (s MetaConditionForListRulesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s MetaConditionForListRulesOutput) GoString() string {
	return s.String()
}

// SetAllDimensions sets the AllDimensions field's value.
func (s *MetaConditionForListRulesOutput) SetAllDimensions(v bool) *MetaConditionForListRulesOutput {
	s.AllDimensions = &v
	return s
}

// SetCondition sets the Condition field's value.
func (s *MetaConditionForListRulesOutput) SetCondition(v string) *MetaConditionForListRulesOutput {
	s.Condition = &v
	return s
}

// SetMetas sets the Metas field's value.
func (s *MetaConditionForListRulesOutput) SetMetas(v []*MetaForListRulesOutput) *MetaConditionForListRulesOutput {
	s.Metas = v
	return s
}

type MetaForListRulesOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Comparator *string `type:"string" json:",omitempty"`

	Key *string `type:"string" json:",omitempty"`

	Values []*string `type:"list" json:",omitempty"`
}

// String returns the string representation
func (s MetaForListRulesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s MetaForListRulesOutput) GoString() string {
	return s.String()
}

// SetComparator sets the Comparator field's value.
func (s *MetaForListRulesOutput) SetComparator(v string) *MetaForListRulesOutput {
	s.Comparator = &v
	return s
}

// SetKey sets the Key field's value.
func (s *MetaForListRulesOutput) SetKey(v string) *MetaForListRulesOutput {
	s.Key = &v
	return s
}

// SetValues sets the Values field's value.
func (s *MetaForListRulesOutput) SetValues(v []*string) *MetaForListRulesOutput {
	s.Values = v
	return s
}

type ProjectConditionForListRulesOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Projects []*string `type:"list" json:",omitempty"`
}

// String returns the string representation
func (s ProjectConditionForListRulesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ProjectConditionForListRulesOutput) GoString() string {
	return s.String()
}

// SetProjects sets the Projects field's value.
func (s *ProjectConditionForListRulesOutput) SetProjects(v []*string) *ProjectConditionForListRulesOutput {
	s.Projects = v
	return s
}

type RecoveryNotifyForListRulesOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Enable *bool `type:"boolean" json:",omitempty"`
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

type TagConditionForListRulesOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Condition *string `type:"string" json:",omitempty"`

	Tags []*TagForListRulesOutput `type:"list" json:",omitempty"`
}

// String returns the string representation
func (s TagConditionForListRulesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s TagConditionForListRulesOutput) GoString() string {
	return s.String()
}

// SetCondition sets the Condition field's value.
func (s *TagConditionForListRulesOutput) SetCondition(v string) *TagConditionForListRulesOutput {
	s.Condition = &v
	return s
}

// SetTags sets the Tags field's value.
func (s *TagConditionForListRulesOutput) SetTags(v []*TagForListRulesOutput) *TagConditionForListRulesOutput {
	s.Tags = v
	return s
}

type TagForListRulesOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Comparator *string `type:"string" json:",omitempty"`

	Key *string `type:"string" json:",omitempty"`

	Values []*string `type:"list" json:",omitempty"`
}

// String returns the string representation
func (s TagForListRulesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s TagForListRulesOutput) GoString() string {
	return s.String()
}

// SetComparator sets the Comparator field's value.
func (s *TagForListRulesOutput) SetComparator(v string) *TagForListRulesOutput {
	s.Comparator = &v
	return s
}

// SetKey sets the Key field's value.
func (s *TagForListRulesOutput) SetKey(v string) *TagForListRulesOutput {
	s.Key = &v
	return s
}

// SetValues sets the Values field's value.
func (s *TagForListRulesOutput) SetValues(v []*string) *TagForListRulesOutput {
	s.Values = v
	return s
}

// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package waf

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opListBotAnalyseProtectRuleCommon = "ListBotAnalyseProtectRule"

// ListBotAnalyseProtectRuleCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ListBotAnalyseProtectRuleCommon operation. The "output" return
// value will be populated with the ListBotAnalyseProtectRuleCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListBotAnalyseProtectRuleCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListBotAnalyseProtectRuleCommon Send returns without error.
//
// See ListBotAnalyseProtectRuleCommon for more information on using the ListBotAnalyseProtectRuleCommon
// API call, and error handling.
//
//    // Example sending a request using the ListBotAnalyseProtectRuleCommonRequest method.
//    req, resp := client.ListBotAnalyseProtectRuleCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *WAF) ListBotAnalyseProtectRuleCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opListBotAnalyseProtectRuleCommon,
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

// ListBotAnalyseProtectRuleCommon API operation for WAF.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for WAF's
// API operation ListBotAnalyseProtectRuleCommon for usage and error information.
func (c *WAF) ListBotAnalyseProtectRuleCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ListBotAnalyseProtectRuleCommonRequest(input)
	return out, req.Send()
}

// ListBotAnalyseProtectRuleCommonWithContext is the same as ListBotAnalyseProtectRuleCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ListBotAnalyseProtectRuleCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *WAF) ListBotAnalyseProtectRuleCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ListBotAnalyseProtectRuleCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opListBotAnalyseProtectRule = "ListBotAnalyseProtectRule"

// ListBotAnalyseProtectRuleRequest generates a "volcengine/request.Request" representing the
// client's request for the ListBotAnalyseProtectRule operation. The "output" return
// value will be populated with the ListBotAnalyseProtectRuleCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListBotAnalyseProtectRuleCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListBotAnalyseProtectRuleCommon Send returns without error.
//
// See ListBotAnalyseProtectRule for more information on using the ListBotAnalyseProtectRule
// API call, and error handling.
//
//    // Example sending a request using the ListBotAnalyseProtectRuleRequest method.
//    req, resp := client.ListBotAnalyseProtectRuleRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *WAF) ListBotAnalyseProtectRuleRequest(input *ListBotAnalyseProtectRuleInput) (req *request.Request, output *ListBotAnalyseProtectRuleOutput) {
	op := &request.Operation{
		Name:       opListBotAnalyseProtectRule,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListBotAnalyseProtectRuleInput{}
	}

	output = &ListBotAnalyseProtectRuleOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// ListBotAnalyseProtectRule API operation for WAF.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for WAF's
// API operation ListBotAnalyseProtectRule for usage and error information.
func (c *WAF) ListBotAnalyseProtectRule(input *ListBotAnalyseProtectRuleInput) (*ListBotAnalyseProtectRuleOutput, error) {
	req, out := c.ListBotAnalyseProtectRuleRequest(input)
	return out, req.Send()
}

// ListBotAnalyseProtectRuleWithContext is the same as ListBotAnalyseProtectRule with the addition of
// the ability to pass a context and additional request options.
//
// See ListBotAnalyseProtectRule for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *WAF) ListBotAnalyseProtectRuleWithContext(ctx volcengine.Context, input *ListBotAnalyseProtectRuleInput, opts ...request.Option) (*ListBotAnalyseProtectRuleOutput, error) {
	req, out := c.ListBotAnalyseProtectRuleRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type AccurateRuleForListBotAnalyseProtectRuleOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	HttpObj *string `type:"string" json:",omitempty"`

	ObjType *int32 `type:"int32" json:",omitempty"`

	Opretar *int32 `type:"int32" json:",omitempty"`

	Property *int32 `type:"int32" json:",omitempty"`

	ValueString *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s AccurateRuleForListBotAnalyseProtectRuleOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s AccurateRuleForListBotAnalyseProtectRuleOutput) GoString() string {
	return s.String()
}

// SetHttpObj sets the HttpObj field's value.
func (s *AccurateRuleForListBotAnalyseProtectRuleOutput) SetHttpObj(v string) *AccurateRuleForListBotAnalyseProtectRuleOutput {
	s.HttpObj = &v
	return s
}

// SetObjType sets the ObjType field's value.
func (s *AccurateRuleForListBotAnalyseProtectRuleOutput) SetObjType(v int32) *AccurateRuleForListBotAnalyseProtectRuleOutput {
	s.ObjType = &v
	return s
}

// SetOpretar sets the Opretar field's value.
func (s *AccurateRuleForListBotAnalyseProtectRuleOutput) SetOpretar(v int32) *AccurateRuleForListBotAnalyseProtectRuleOutput {
	s.Opretar = &v
	return s
}

// SetProperty sets the Property field's value.
func (s *AccurateRuleForListBotAnalyseProtectRuleOutput) SetProperty(v int32) *AccurateRuleForListBotAnalyseProtectRuleOutput {
	s.Property = &v
	return s
}

// SetValueString sets the ValueString field's value.
func (s *AccurateRuleForListBotAnalyseProtectRuleOutput) SetValueString(v string) *AccurateRuleForListBotAnalyseProtectRuleOutput {
	s.ValueString = &v
	return s
}

type DataForListBotAnalyseProtectRuleOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	EnableCount *int32 `type:"int32" json:",omitempty"`

	Path *string `type:"string" json:",omitempty"`

	RuleGroup []*RuleGroupForListBotAnalyseProtectRuleOutput `type:"list" json:",omitempty"`

	TotalCount *int32 `type:"int32" json:",omitempty"`
}

// String returns the string representation
func (s DataForListBotAnalyseProtectRuleOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DataForListBotAnalyseProtectRuleOutput) GoString() string {
	return s.String()
}

// SetEnableCount sets the EnableCount field's value.
func (s *DataForListBotAnalyseProtectRuleOutput) SetEnableCount(v int32) *DataForListBotAnalyseProtectRuleOutput {
	s.EnableCount = &v
	return s
}

// SetPath sets the Path field's value.
func (s *DataForListBotAnalyseProtectRuleOutput) SetPath(v string) *DataForListBotAnalyseProtectRuleOutput {
	s.Path = &v
	return s
}

// SetRuleGroup sets the RuleGroup field's value.
func (s *DataForListBotAnalyseProtectRuleOutput) SetRuleGroup(v []*RuleGroupForListBotAnalyseProtectRuleOutput) *DataForListBotAnalyseProtectRuleOutput {
	s.RuleGroup = v
	return s
}

// SetTotalCount sets the TotalCount field's value.
func (s *DataForListBotAnalyseProtectRuleOutput) SetTotalCount(v int32) *DataForListBotAnalyseProtectRuleOutput {
	s.TotalCount = &v
	return s
}

type GroupForListBotAnalyseProtectRuleOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	AccurateGroupPriority *int32 `type:"int32" json:",omitempty"`

	AccurateRules []*AccurateRuleForListBotAnalyseProtectRuleOutput `type:"list" json:",omitempty"`

	Id *int32 `type:"int32" json:",omitempty"`

	Logic *int32 `type:"int32" json:",omitempty"`
}

// String returns the string representation
func (s GroupForListBotAnalyseProtectRuleOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s GroupForListBotAnalyseProtectRuleOutput) GoString() string {
	return s.String()
}

// SetAccurateGroupPriority sets the AccurateGroupPriority field's value.
func (s *GroupForListBotAnalyseProtectRuleOutput) SetAccurateGroupPriority(v int32) *GroupForListBotAnalyseProtectRuleOutput {
	s.AccurateGroupPriority = &v
	return s
}

// SetAccurateRules sets the AccurateRules field's value.
func (s *GroupForListBotAnalyseProtectRuleOutput) SetAccurateRules(v []*AccurateRuleForListBotAnalyseProtectRuleOutput) *GroupForListBotAnalyseProtectRuleOutput {
	s.AccurateRules = v
	return s
}

// SetId sets the Id field's value.
func (s *GroupForListBotAnalyseProtectRuleOutput) SetId(v int32) *GroupForListBotAnalyseProtectRuleOutput {
	s.Id = &v
	return s
}

// SetLogic sets the Logic field's value.
func (s *GroupForListBotAnalyseProtectRuleOutput) SetLogic(v int32) *GroupForListBotAnalyseProtectRuleOutput {
	s.Logic = &v
	return s
}

type ListBotAnalyseProtectRuleInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	// BotSpace is a required field
	BotSpace *string `type:"string" json:",omitempty" required:"true"`

	// Host is a required field
	Host *string `type:"string" json:",omitempty" required:"true"`

	Name *string `type:"string" json:",omitempty"`

	Page *int32 `type:"int32" json:",omitempty"`

	PageSize *int32 `type:"int32" json:",omitempty"`

	Path *string `type:"string" json:",omitempty"`

	ProjectName *string `type:"string" json:",omitempty"`

	// Region is a required field
	Region *string `type:"string" json:",omitempty" required:"true"`

	RuleTag *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s ListBotAnalyseProtectRuleInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListBotAnalyseProtectRuleInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListBotAnalyseProtectRuleInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ListBotAnalyseProtectRuleInput"}
	if s.BotSpace == nil {
		invalidParams.Add(request.NewErrParamRequired("BotSpace"))
	}
	if s.Host == nil {
		invalidParams.Add(request.NewErrParamRequired("Host"))
	}
	if s.Region == nil {
		invalidParams.Add(request.NewErrParamRequired("Region"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetBotSpace sets the BotSpace field's value.
func (s *ListBotAnalyseProtectRuleInput) SetBotSpace(v string) *ListBotAnalyseProtectRuleInput {
	s.BotSpace = &v
	return s
}

// SetHost sets the Host field's value.
func (s *ListBotAnalyseProtectRuleInput) SetHost(v string) *ListBotAnalyseProtectRuleInput {
	s.Host = &v
	return s
}

// SetName sets the Name field's value.
func (s *ListBotAnalyseProtectRuleInput) SetName(v string) *ListBotAnalyseProtectRuleInput {
	s.Name = &v
	return s
}

// SetPage sets the Page field's value.
func (s *ListBotAnalyseProtectRuleInput) SetPage(v int32) *ListBotAnalyseProtectRuleInput {
	s.Page = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *ListBotAnalyseProtectRuleInput) SetPageSize(v int32) *ListBotAnalyseProtectRuleInput {
	s.PageSize = &v
	return s
}

// SetPath sets the Path field's value.
func (s *ListBotAnalyseProtectRuleInput) SetPath(v string) *ListBotAnalyseProtectRuleInput {
	s.Path = &v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *ListBotAnalyseProtectRuleInput) SetProjectName(v string) *ListBotAnalyseProtectRuleInput {
	s.ProjectName = &v
	return s
}

// SetRegion sets the Region field's value.
func (s *ListBotAnalyseProtectRuleInput) SetRegion(v string) *ListBotAnalyseProtectRuleInput {
	s.Region = &v
	return s
}

// SetRuleTag sets the RuleTag field's value.
func (s *ListBotAnalyseProtectRuleInput) SetRuleTag(v string) *ListBotAnalyseProtectRuleInput {
	s.RuleTag = &v
	return s
}

type ListBotAnalyseProtectRuleOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	Count *int32 `type:"int32" json:",omitempty"`

	CurrentPage *int32 `type:"int32" json:",omitempty"`

	Data []*DataForListBotAnalyseProtectRuleOutput `type:"list" json:",omitempty"`

	PageNumber *int32 `type:"int32" json:",omitempty"`

	PageSize *int32 `type:"int32" json:",omitempty"`

	TotalCount *int32 `type:"int32" json:",omitempty"`
}

// String returns the string representation
func (s ListBotAnalyseProtectRuleOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListBotAnalyseProtectRuleOutput) GoString() string {
	return s.String()
}

// SetCount sets the Count field's value.
func (s *ListBotAnalyseProtectRuleOutput) SetCount(v int32) *ListBotAnalyseProtectRuleOutput {
	s.Count = &v
	return s
}

// SetCurrentPage sets the CurrentPage field's value.
func (s *ListBotAnalyseProtectRuleOutput) SetCurrentPage(v int32) *ListBotAnalyseProtectRuleOutput {
	s.CurrentPage = &v
	return s
}

// SetData sets the Data field's value.
func (s *ListBotAnalyseProtectRuleOutput) SetData(v []*DataForListBotAnalyseProtectRuleOutput) *ListBotAnalyseProtectRuleOutput {
	s.Data = v
	return s
}

// SetPageNumber sets the PageNumber field's value.
func (s *ListBotAnalyseProtectRuleOutput) SetPageNumber(v int32) *ListBotAnalyseProtectRuleOutput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *ListBotAnalyseProtectRuleOutput) SetPageSize(v int32) *ListBotAnalyseProtectRuleOutput {
	s.PageSize = &v
	return s
}

// SetTotalCount sets the TotalCount field's value.
func (s *ListBotAnalyseProtectRuleOutput) SetTotalCount(v int32) *ListBotAnalyseProtectRuleOutput {
	s.TotalCount = &v
	return s
}

type RuleForListBotAnalyseProtectRuleOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	AccurateGroup *GroupForListBotAnalyseProtectRuleOutput `type:"structure" json:",omitempty"`

	AccurateGroupPriority *int32 `type:"int32" json:",omitempty"`

	ActionAfterVerification *int32 `type:"int32" json:",omitempty"`

	ActionType *int32 `type:"int32" json:",omitempty"`

	EffectTime *int32 `type:"int32" json:",omitempty"`

	Enable *int32 `type:"int32" json:",omitempty"`

	ExemptionTime *int32 `type:"int32" json:",omitempty"`

	Field *string `type:"string" json:",omitempty"`

	Host *string `type:"string" json:",omitempty"`

	Id *int32 `type:"int32" json:",omitempty"`

	Name *string `type:"string" json:",omitempty"`

	PassRatio *float64 `type:"float" json:",omitempty"`

	Path *string `type:"string" json:",omitempty"`

	PathThreshold *int32 `type:"int32" json:",omitempty"`

	RulePriority *int32 `type:"int32" json:",omitempty"`

	RuleTag *string `type:"string" json:",omitempty"`

	SingleProportion *float64 `type:"float" json:",omitempty"`

	SingleThreshold *int32 `type:"int32" json:",omitempty"`

	StatisticalDuration *int32 `type:"int32" json:",omitempty"`

	StatisticalType *int32 `type:"int32" json:",omitempty"`

	UpdateTime *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s RuleForListBotAnalyseProtectRuleOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s RuleForListBotAnalyseProtectRuleOutput) GoString() string {
	return s.String()
}

// SetAccurateGroup sets the AccurateGroup field's value.
func (s *RuleForListBotAnalyseProtectRuleOutput) SetAccurateGroup(v *GroupForListBotAnalyseProtectRuleOutput) *RuleForListBotAnalyseProtectRuleOutput {
	s.AccurateGroup = v
	return s
}

// SetAccurateGroupPriority sets the AccurateGroupPriority field's value.
func (s *RuleForListBotAnalyseProtectRuleOutput) SetAccurateGroupPriority(v int32) *RuleForListBotAnalyseProtectRuleOutput {
	s.AccurateGroupPriority = &v
	return s
}

// SetActionAfterVerification sets the ActionAfterVerification field's value.
func (s *RuleForListBotAnalyseProtectRuleOutput) SetActionAfterVerification(v int32) *RuleForListBotAnalyseProtectRuleOutput {
	s.ActionAfterVerification = &v
	return s
}

// SetActionType sets the ActionType field's value.
func (s *RuleForListBotAnalyseProtectRuleOutput) SetActionType(v int32) *RuleForListBotAnalyseProtectRuleOutput {
	s.ActionType = &v
	return s
}

// SetEffectTime sets the EffectTime field's value.
func (s *RuleForListBotAnalyseProtectRuleOutput) SetEffectTime(v int32) *RuleForListBotAnalyseProtectRuleOutput {
	s.EffectTime = &v
	return s
}

// SetEnable sets the Enable field's value.
func (s *RuleForListBotAnalyseProtectRuleOutput) SetEnable(v int32) *RuleForListBotAnalyseProtectRuleOutput {
	s.Enable = &v
	return s
}

// SetExemptionTime sets the ExemptionTime field's value.
func (s *RuleForListBotAnalyseProtectRuleOutput) SetExemptionTime(v int32) *RuleForListBotAnalyseProtectRuleOutput {
	s.ExemptionTime = &v
	return s
}

// SetField sets the Field field's value.
func (s *RuleForListBotAnalyseProtectRuleOutput) SetField(v string) *RuleForListBotAnalyseProtectRuleOutput {
	s.Field = &v
	return s
}

// SetHost sets the Host field's value.
func (s *RuleForListBotAnalyseProtectRuleOutput) SetHost(v string) *RuleForListBotAnalyseProtectRuleOutput {
	s.Host = &v
	return s
}

// SetId sets the Id field's value.
func (s *RuleForListBotAnalyseProtectRuleOutput) SetId(v int32) *RuleForListBotAnalyseProtectRuleOutput {
	s.Id = &v
	return s
}

// SetName sets the Name field's value.
func (s *RuleForListBotAnalyseProtectRuleOutput) SetName(v string) *RuleForListBotAnalyseProtectRuleOutput {
	s.Name = &v
	return s
}

// SetPassRatio sets the PassRatio field's value.
func (s *RuleForListBotAnalyseProtectRuleOutput) SetPassRatio(v float64) *RuleForListBotAnalyseProtectRuleOutput {
	s.PassRatio = &v
	return s
}

// SetPath sets the Path field's value.
func (s *RuleForListBotAnalyseProtectRuleOutput) SetPath(v string) *RuleForListBotAnalyseProtectRuleOutput {
	s.Path = &v
	return s
}

// SetPathThreshold sets the PathThreshold field's value.
func (s *RuleForListBotAnalyseProtectRuleOutput) SetPathThreshold(v int32) *RuleForListBotAnalyseProtectRuleOutput {
	s.PathThreshold = &v
	return s
}

// SetRulePriority sets the RulePriority field's value.
func (s *RuleForListBotAnalyseProtectRuleOutput) SetRulePriority(v int32) *RuleForListBotAnalyseProtectRuleOutput {
	s.RulePriority = &v
	return s
}

// SetRuleTag sets the RuleTag field's value.
func (s *RuleForListBotAnalyseProtectRuleOutput) SetRuleTag(v string) *RuleForListBotAnalyseProtectRuleOutput {
	s.RuleTag = &v
	return s
}

// SetSingleProportion sets the SingleProportion field's value.
func (s *RuleForListBotAnalyseProtectRuleOutput) SetSingleProportion(v float64) *RuleForListBotAnalyseProtectRuleOutput {
	s.SingleProportion = &v
	return s
}

// SetSingleThreshold sets the SingleThreshold field's value.
func (s *RuleForListBotAnalyseProtectRuleOutput) SetSingleThreshold(v int32) *RuleForListBotAnalyseProtectRuleOutput {
	s.SingleThreshold = &v
	return s
}

// SetStatisticalDuration sets the StatisticalDuration field's value.
func (s *RuleForListBotAnalyseProtectRuleOutput) SetStatisticalDuration(v int32) *RuleForListBotAnalyseProtectRuleOutput {
	s.StatisticalDuration = &v
	return s
}

// SetStatisticalType sets the StatisticalType field's value.
func (s *RuleForListBotAnalyseProtectRuleOutput) SetStatisticalType(v int32) *RuleForListBotAnalyseProtectRuleOutput {
	s.StatisticalType = &v
	return s
}

// SetUpdateTime sets the UpdateTime field's value.
func (s *RuleForListBotAnalyseProtectRuleOutput) SetUpdateTime(v string) *RuleForListBotAnalyseProtectRuleOutput {
	s.UpdateTime = &v
	return s
}

type RuleGroupForListBotAnalyseProtectRuleOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Group *GroupForListBotAnalyseProtectRuleOutput `type:"structure" json:",omitempty"`

	Rules []*RuleForListBotAnalyseProtectRuleOutput `type:"list" json:",omitempty"`
}

// String returns the string representation
func (s RuleGroupForListBotAnalyseProtectRuleOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s RuleGroupForListBotAnalyseProtectRuleOutput) GoString() string {
	return s.String()
}

// SetGroup sets the Group field's value.
func (s *RuleGroupForListBotAnalyseProtectRuleOutput) SetGroup(v *GroupForListBotAnalyseProtectRuleOutput) *RuleGroupForListBotAnalyseProtectRuleOutput {
	s.Group = v
	return s
}

// SetRules sets the Rules field's value.
func (s *RuleGroupForListBotAnalyseProtectRuleOutput) SetRules(v []*RuleForListBotAnalyseProtectRuleOutput) *RuleGroupForListBotAnalyseProtectRuleOutput {
	s.Rules = v
	return s
}

// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package waf

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opListCCRuleCommon = "ListCCRule"

// ListCCRuleCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ListCCRuleCommon operation. The "output" return
// value will be populated with the ListCCRuleCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListCCRuleCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListCCRuleCommon Send returns without error.
//
// See ListCCRuleCommon for more information on using the ListCCRuleCommon
// API call, and error handling.
//
//    // Example sending a request using the ListCCRuleCommonRequest method.
//    req, resp := client.ListCCRuleCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *WAF) ListCCRuleCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opListCCRuleCommon,
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

// ListCCRuleCommon API operation for WAF.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for WAF's
// API operation ListCCRuleCommon for usage and error information.
func (c *WAF) ListCCRuleCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ListCCRuleCommonRequest(input)
	return out, req.Send()
}

// ListCCRuleCommonWithContext is the same as ListCCRuleCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ListCCRuleCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *WAF) ListCCRuleCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ListCCRuleCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opListCCRule = "ListCCRule"

// ListCCRuleRequest generates a "volcengine/request.Request" representing the
// client's request for the ListCCRule operation. The "output" return
// value will be populated with the ListCCRuleCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListCCRuleCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListCCRuleCommon Send returns without error.
//
// See ListCCRule for more information on using the ListCCRule
// API call, and error handling.
//
//    // Example sending a request using the ListCCRuleRequest method.
//    req, resp := client.ListCCRuleRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *WAF) ListCCRuleRequest(input *ListCCRuleInput) (req *request.Request, output *ListCCRuleOutput) {
	op := &request.Operation{
		Name:       opListCCRule,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListCCRuleInput{}
	}

	output = &ListCCRuleOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// ListCCRule API operation for WAF.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for WAF's
// API operation ListCCRule for usage and error information.
func (c *WAF) ListCCRule(input *ListCCRuleInput) (*ListCCRuleOutput, error) {
	req, out := c.ListCCRuleRequest(input)
	return out, req.Send()
}

// ListCCRuleWithContext is the same as ListCCRule with the addition of
// the ability to pass a context and additional request options.
//
// See ListCCRule for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *WAF) ListCCRuleWithContext(ctx volcengine.Context, input *ListCCRuleInput, opts ...request.Option) (*ListCCRuleOutput, error) {
	req, out := c.ListCCRuleRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type AccurateRuleForListCCRuleOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	HttpObj *string `type:"string" json:",omitempty"`

	ObjType *int32 `type:"int32" json:",omitempty"`

	Opretar *int32 `type:"int32" json:",omitempty"`

	Property *int32 `type:"int32" json:",omitempty"`

	ValueString *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s AccurateRuleForListCCRuleOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s AccurateRuleForListCCRuleOutput) GoString() string {
	return s.String()
}

// SetHttpObj sets the HttpObj field's value.
func (s *AccurateRuleForListCCRuleOutput) SetHttpObj(v string) *AccurateRuleForListCCRuleOutput {
	s.HttpObj = &v
	return s
}

// SetObjType sets the ObjType field's value.
func (s *AccurateRuleForListCCRuleOutput) SetObjType(v int32) *AccurateRuleForListCCRuleOutput {
	s.ObjType = &v
	return s
}

// SetOpretar sets the Opretar field's value.
func (s *AccurateRuleForListCCRuleOutput) SetOpretar(v int32) *AccurateRuleForListCCRuleOutput {
	s.Opretar = &v
	return s
}

// SetProperty sets the Property field's value.
func (s *AccurateRuleForListCCRuleOutput) SetProperty(v int32) *AccurateRuleForListCCRuleOutput {
	s.Property = &v
	return s
}

// SetValueString sets the ValueString field's value.
func (s *AccurateRuleForListCCRuleOutput) SetValueString(v string) *AccurateRuleForListCCRuleOutput {
	s.ValueString = &v
	return s
}

type CronConfForListCCRuleOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Crontab *string `type:"string" json:",omitempty"`

	PathThreshold *int32 `type:"int32" json:",omitempty"`

	SingleThreshold *int32 `type:"int32" json:",omitempty"`
}

// String returns the string representation
func (s CronConfForListCCRuleOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CronConfForListCCRuleOutput) GoString() string {
	return s.String()
}

// SetCrontab sets the Crontab field's value.
func (s *CronConfForListCCRuleOutput) SetCrontab(v string) *CronConfForListCCRuleOutput {
	s.Crontab = &v
	return s
}

// SetPathThreshold sets the PathThreshold field's value.
func (s *CronConfForListCCRuleOutput) SetPathThreshold(v int32) *CronConfForListCCRuleOutput {
	s.PathThreshold = &v
	return s
}

// SetSingleThreshold sets the SingleThreshold field's value.
func (s *CronConfForListCCRuleOutput) SetSingleThreshold(v int32) *CronConfForListCCRuleOutput {
	s.SingleThreshold = &v
	return s
}

type GroupForListCCRuleOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	AccurateRules []*AccurateRuleForListCCRuleOutput `type:"list" json:",omitempty"`

	Logic *int32 `type:"int32" json:",omitempty"`
}

// String returns the string representation
func (s GroupForListCCRuleOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s GroupForListCCRuleOutput) GoString() string {
	return s.String()
}

// SetAccurateRules sets the AccurateRules field's value.
func (s *GroupForListCCRuleOutput) SetAccurateRules(v []*AccurateRuleForListCCRuleOutput) *GroupForListCCRuleOutput {
	s.AccurateRules = v
	return s
}

// SetLogic sets the Logic field's value.
func (s *GroupForListCCRuleOutput) SetLogic(v int32) *GroupForListCCRuleOutput {
	s.Logic = &v
	return s
}

type ListCCRuleInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	CCType []*int32 `type:"list" json:",omitempty"`

	// Host is a required field
	Host *string `type:"string" json:",omitempty" required:"true"`

	Page *int32 `type:"int32" json:",omitempty"`

	PageSize *int32 `type:"int32" json:",omitempty"`

	PathOrderBy *string `type:"string" json:",omitempty"`

	ProjectName *string `type:"string" json:",omitempty"`

	RuleName *string `type:"string" json:",omitempty"`

	RuleTag *string `type:"string" json:",omitempty"`

	Url *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s ListCCRuleInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListCCRuleInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListCCRuleInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ListCCRuleInput"}
	if s.Host == nil {
		invalidParams.Add(request.NewErrParamRequired("Host"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetCCType sets the CCType field's value.
func (s *ListCCRuleInput) SetCCType(v []*int32) *ListCCRuleInput {
	s.CCType = v
	return s
}

// SetHost sets the Host field's value.
func (s *ListCCRuleInput) SetHost(v string) *ListCCRuleInput {
	s.Host = &v
	return s
}

// SetPage sets the Page field's value.
func (s *ListCCRuleInput) SetPage(v int32) *ListCCRuleInput {
	s.Page = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *ListCCRuleInput) SetPageSize(v int32) *ListCCRuleInput {
	s.PageSize = &v
	return s
}

// SetPathOrderBy sets the PathOrderBy field's value.
func (s *ListCCRuleInput) SetPathOrderBy(v string) *ListCCRuleInput {
	s.PathOrderBy = &v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *ListCCRuleInput) SetProjectName(v string) *ListCCRuleInput {
	s.ProjectName = &v
	return s
}

// SetRuleName sets the RuleName field's value.
func (s *ListCCRuleInput) SetRuleName(v string) *ListCCRuleInput {
	s.RuleName = &v
	return s
}

// SetRuleTag sets the RuleTag field's value.
func (s *ListCCRuleInput) SetRuleTag(v string) *ListCCRuleInput {
	s.RuleTag = &v
	return s
}

// SetUrl sets the Url field's value.
func (s *ListCCRuleInput) SetUrl(v string) *ListCCRuleInput {
	s.Url = &v
	return s
}

type ListCCRuleOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	Count *int32 `type:"int32" json:",omitempty"`

	CurrentPage *int32 `type:"int32" json:",omitempty"`

	EnableCount *int32 `type:"int32" json:",omitempty"`

	InsertTime *string `type:"string" json:",omitempty"`

	PageSize *int32 `type:"int32" json:",omitempty"`

	RuleGroup []*RuleGroupForListCCRuleOutput `type:"list" json:",omitempty"`

	TotalCount *int32 `type:"int32" json:",omitempty"`

	Url *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s ListCCRuleOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListCCRuleOutput) GoString() string {
	return s.String()
}

// SetCount sets the Count field's value.
func (s *ListCCRuleOutput) SetCount(v int32) *ListCCRuleOutput {
	s.Count = &v
	return s
}

// SetCurrentPage sets the CurrentPage field's value.
func (s *ListCCRuleOutput) SetCurrentPage(v int32) *ListCCRuleOutput {
	s.CurrentPage = &v
	return s
}

// SetEnableCount sets the EnableCount field's value.
func (s *ListCCRuleOutput) SetEnableCount(v int32) *ListCCRuleOutput {
	s.EnableCount = &v
	return s
}

// SetInsertTime sets the InsertTime field's value.
func (s *ListCCRuleOutput) SetInsertTime(v string) *ListCCRuleOutput {
	s.InsertTime = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *ListCCRuleOutput) SetPageSize(v int32) *ListCCRuleOutput {
	s.PageSize = &v
	return s
}

// SetRuleGroup sets the RuleGroup field's value.
func (s *ListCCRuleOutput) SetRuleGroup(v []*RuleGroupForListCCRuleOutput) *ListCCRuleOutput {
	s.RuleGroup = v
	return s
}

// SetTotalCount sets the TotalCount field's value.
func (s *ListCCRuleOutput) SetTotalCount(v int32) *ListCCRuleOutput {
	s.TotalCount = &v
	return s
}

// SetUrl sets the Url field's value.
func (s *ListCCRuleOutput) SetUrl(v string) *ListCCRuleOutput {
	s.Url = &v
	return s
}

type RuleForListCCRuleOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	AccurateGroup *GroupForListCCRuleOutput `type:"structure" json:",omitempty"`

	AccurateGroupPriority *int32 `type:"int32" json:",omitempty"`

	CCType *int32 `type:"int32" json:",omitempty"`

	CountTime *int32 `type:"int32" json:",omitempty"`

	CronConfs []*CronConfForListCCRuleOutput `type:"list" json:",omitempty"`

	CronEnable *int32 `type:"int32" json:",omitempty"`

	EffectTime *int32 `type:"int32" json:",omitempty"`

	Enable *int32 `type:"int32" json:",omitempty"`

	ExemptionTime *int32 `type:"int32" json:",omitempty"`

	Field *string `type:"string" json:",omitempty"`

	Host *string `type:"string" json:",omitempty"`

	Id *int32 `type:"int32" json:",omitempty"`

	Name *string `type:"string" json:",omitempty"`

	PathThreshold *int32 `type:"int32" json:",omitempty"`

	RulePriority *int32 `type:"int32" json:",omitempty"`

	RuleTag *string `type:"string" json:",omitempty"`

	SingleThreshold *int32 `type:"int32" json:",omitempty"`

	UpdateTime *string `type:"string" json:",omitempty"`

	Url *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s RuleForListCCRuleOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s RuleForListCCRuleOutput) GoString() string {
	return s.String()
}

// SetAccurateGroup sets the AccurateGroup field's value.
func (s *RuleForListCCRuleOutput) SetAccurateGroup(v *GroupForListCCRuleOutput) *RuleForListCCRuleOutput {
	s.AccurateGroup = v
	return s
}

// SetAccurateGroupPriority sets the AccurateGroupPriority field's value.
func (s *RuleForListCCRuleOutput) SetAccurateGroupPriority(v int32) *RuleForListCCRuleOutput {
	s.AccurateGroupPriority = &v
	return s
}

// SetCCType sets the CCType field's value.
func (s *RuleForListCCRuleOutput) SetCCType(v int32) *RuleForListCCRuleOutput {
	s.CCType = &v
	return s
}

// SetCountTime sets the CountTime field's value.
func (s *RuleForListCCRuleOutput) SetCountTime(v int32) *RuleForListCCRuleOutput {
	s.CountTime = &v
	return s
}

// SetCronConfs sets the CronConfs field's value.
func (s *RuleForListCCRuleOutput) SetCronConfs(v []*CronConfForListCCRuleOutput) *RuleForListCCRuleOutput {
	s.CronConfs = v
	return s
}

// SetCronEnable sets the CronEnable field's value.
func (s *RuleForListCCRuleOutput) SetCronEnable(v int32) *RuleForListCCRuleOutput {
	s.CronEnable = &v
	return s
}

// SetEffectTime sets the EffectTime field's value.
func (s *RuleForListCCRuleOutput) SetEffectTime(v int32) *RuleForListCCRuleOutput {
	s.EffectTime = &v
	return s
}

// SetEnable sets the Enable field's value.
func (s *RuleForListCCRuleOutput) SetEnable(v int32) *RuleForListCCRuleOutput {
	s.Enable = &v
	return s
}

// SetExemptionTime sets the ExemptionTime field's value.
func (s *RuleForListCCRuleOutput) SetExemptionTime(v int32) *RuleForListCCRuleOutput {
	s.ExemptionTime = &v
	return s
}

// SetField sets the Field field's value.
func (s *RuleForListCCRuleOutput) SetField(v string) *RuleForListCCRuleOutput {
	s.Field = &v
	return s
}

// SetHost sets the Host field's value.
func (s *RuleForListCCRuleOutput) SetHost(v string) *RuleForListCCRuleOutput {
	s.Host = &v
	return s
}

// SetId sets the Id field's value.
func (s *RuleForListCCRuleOutput) SetId(v int32) *RuleForListCCRuleOutput {
	s.Id = &v
	return s
}

// SetName sets the Name field's value.
func (s *RuleForListCCRuleOutput) SetName(v string) *RuleForListCCRuleOutput {
	s.Name = &v
	return s
}

// SetPathThreshold sets the PathThreshold field's value.
func (s *RuleForListCCRuleOutput) SetPathThreshold(v int32) *RuleForListCCRuleOutput {
	s.PathThreshold = &v
	return s
}

// SetRulePriority sets the RulePriority field's value.
func (s *RuleForListCCRuleOutput) SetRulePriority(v int32) *RuleForListCCRuleOutput {
	s.RulePriority = &v
	return s
}

// SetRuleTag sets the RuleTag field's value.
func (s *RuleForListCCRuleOutput) SetRuleTag(v string) *RuleForListCCRuleOutput {
	s.RuleTag = &v
	return s
}

// SetSingleThreshold sets the SingleThreshold field's value.
func (s *RuleForListCCRuleOutput) SetSingleThreshold(v int32) *RuleForListCCRuleOutput {
	s.SingleThreshold = &v
	return s
}

// SetUpdateTime sets the UpdateTime field's value.
func (s *RuleForListCCRuleOutput) SetUpdateTime(v string) *RuleForListCCRuleOutput {
	s.UpdateTime = &v
	return s
}

// SetUrl sets the Url field's value.
func (s *RuleForListCCRuleOutput) SetUrl(v string) *RuleForListCCRuleOutput {
	s.Url = &v
	return s
}

type RuleGroupForListCCRuleOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Group *GroupForListCCRuleOutput `type:"structure" json:",omitempty"`

	Rules []*RuleForListCCRuleOutput `type:"list" json:",omitempty"`
}

// String returns the string representation
func (s RuleGroupForListCCRuleOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s RuleGroupForListCCRuleOutput) GoString() string {
	return s.String()
}

// SetGroup sets the Group field's value.
func (s *RuleGroupForListCCRuleOutput) SetGroup(v *GroupForListCCRuleOutput) *RuleGroupForListCCRuleOutput {
	s.Group = v
	return s
}

// SetRules sets the Rules field's value.
func (s *RuleGroupForListCCRuleOutput) SetRules(v []*RuleForListCCRuleOutput) *RuleGroupForListCCRuleOutput {
	s.Rules = v
	return s
}

// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package waf

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opListAclRuleCommon = "ListAclRule"

// ListAclRuleCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ListAclRuleCommon operation. The "output" return
// value will be populated with the ListAclRuleCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListAclRuleCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListAclRuleCommon Send returns without error.
//
// See ListAclRuleCommon for more information on using the ListAclRuleCommon
// API call, and error handling.
//
//    // Example sending a request using the ListAclRuleCommonRequest method.
//    req, resp := client.ListAclRuleCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *WAF) ListAclRuleCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opListAclRuleCommon,
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

// ListAclRuleCommon API operation for WAF.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for WAF's
// API operation ListAclRuleCommon for usage and error information.
func (c *WAF) ListAclRuleCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ListAclRuleCommonRequest(input)
	return out, req.Send()
}

// ListAclRuleCommonWithContext is the same as ListAclRuleCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ListAclRuleCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *WAF) ListAclRuleCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ListAclRuleCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opListAclRule = "ListAclRule"

// ListAclRuleRequest generates a "volcengine/request.Request" representing the
// client's request for the ListAclRule operation. The "output" return
// value will be populated with the ListAclRuleCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListAclRuleCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListAclRuleCommon Send returns without error.
//
// See ListAclRule for more information on using the ListAclRule
// API call, and error handling.
//
//    // Example sending a request using the ListAclRuleRequest method.
//    req, resp := client.ListAclRuleRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *WAF) ListAclRuleRequest(input *ListAclRuleInput) (req *request.Request, output *ListAclRuleOutput) {
	op := &request.Operation{
		Name:       opListAclRule,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListAclRuleInput{}
	}

	output = &ListAclRuleOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// ListAclRule API operation for WAF.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for WAF's
// API operation ListAclRule for usage and error information.
func (c *WAF) ListAclRule(input *ListAclRuleInput) (*ListAclRuleOutput, error) {
	req, out := c.ListAclRuleRequest(input)
	return out, req.Send()
}

// ListAclRuleWithContext is the same as ListAclRule with the addition of
// the ability to pass a context and additional request options.
//
// See ListAclRule for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *WAF) ListAclRuleWithContext(ctx volcengine.Context, input *ListAclRuleInput, opts ...request.Option) (*ListAclRuleOutput, error) {
	req, out := c.ListAclRuleRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type AccurateGroupForListAclRuleOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	AccurateRules []*AccurateRuleForListAclRuleOutput `type:"list" json:",omitempty"`

	Logic *int32 `type:"int32" json:",omitempty"`
}

// String returns the string representation
func (s AccurateGroupForListAclRuleOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s AccurateGroupForListAclRuleOutput) GoString() string {
	return s.String()
}

// SetAccurateRules sets the AccurateRules field's value.
func (s *AccurateGroupForListAclRuleOutput) SetAccurateRules(v []*AccurateRuleForListAclRuleOutput) *AccurateGroupForListAclRuleOutput {
	s.AccurateRules = v
	return s
}

// SetLogic sets the Logic field's value.
func (s *AccurateGroupForListAclRuleOutput) SetLogic(v int32) *AccurateGroupForListAclRuleOutput {
	s.Logic = &v
	return s
}

type AccurateRuleForListAclRuleOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	HttpObj *string `type:"string" json:",omitempty"`

	ObjType *int32 `type:"int32" json:",omitempty"`

	Opretar *int32 `type:"int32" json:",omitempty"`

	Property *int32 `type:"int32" json:",omitempty"`

	ValueString *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s AccurateRuleForListAclRuleOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s AccurateRuleForListAclRuleOutput) GoString() string {
	return s.String()
}

// SetHttpObj sets the HttpObj field's value.
func (s *AccurateRuleForListAclRuleOutput) SetHttpObj(v string) *AccurateRuleForListAclRuleOutput {
	s.HttpObj = &v
	return s
}

// SetObjType sets the ObjType field's value.
func (s *AccurateRuleForListAclRuleOutput) SetObjType(v int32) *AccurateRuleForListAclRuleOutput {
	s.ObjType = &v
	return s
}

// SetOpretar sets the Opretar field's value.
func (s *AccurateRuleForListAclRuleOutput) SetOpretar(v int32) *AccurateRuleForListAclRuleOutput {
	s.Opretar = &v
	return s
}

// SetProperty sets the Property field's value.
func (s *AccurateRuleForListAclRuleOutput) SetProperty(v int32) *AccurateRuleForListAclRuleOutput {
	s.Property = &v
	return s
}

// SetValueString sets the ValueString field's value.
func (s *AccurateRuleForListAclRuleOutput) SetValueString(v string) *AccurateRuleForListAclRuleOutput {
	s.ValueString = &v
	return s
}

type HostGroupForListAclRuleOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	HostGroupId *int32 `type:"int32" json:",omitempty"`

	Name *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s HostGroupForListAclRuleOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s HostGroupForListAclRuleOutput) GoString() string {
	return s.String()
}

// SetHostGroupId sets the HostGroupId field's value.
func (s *HostGroupForListAclRuleOutput) SetHostGroupId(v int32) *HostGroupForListAclRuleOutput {
	s.HostGroupId = &v
	return s
}

// SetName sets the Name field's value.
func (s *HostGroupForListAclRuleOutput) SetName(v string) *HostGroupForListAclRuleOutput {
	s.Name = &v
	return s
}

type ListAclRuleInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	// AclType is a required field
	AclType *string `type:"string" json:",omitempty" required:"true"`

	Action []*string `type:"list" json:",omitempty"`

	DefenceHost []*string `type:"list" json:",omitempty"`

	Enable []*int32 `type:"list" json:",omitempty"`

	Page *int32 `type:"int32" json:",omitempty"`

	PageSize *int32 `type:"int32" json:",omitempty"`

	RuleName *string `type:"string" json:",omitempty"`

	RuleTag *string `type:"string" json:",omitempty"`

	TimeOrderBy *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s ListAclRuleInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListAclRuleInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListAclRuleInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ListAclRuleInput"}
	if s.AclType == nil {
		invalidParams.Add(request.NewErrParamRequired("AclType"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetAclType sets the AclType field's value.
func (s *ListAclRuleInput) SetAclType(v string) *ListAclRuleInput {
	s.AclType = &v
	return s
}

// SetAction sets the Action field's value.
func (s *ListAclRuleInput) SetAction(v []*string) *ListAclRuleInput {
	s.Action = v
	return s
}

// SetDefenceHost sets the DefenceHost field's value.
func (s *ListAclRuleInput) SetDefenceHost(v []*string) *ListAclRuleInput {
	s.DefenceHost = v
	return s
}

// SetEnable sets the Enable field's value.
func (s *ListAclRuleInput) SetEnable(v []*int32) *ListAclRuleInput {
	s.Enable = v
	return s
}

// SetPage sets the Page field's value.
func (s *ListAclRuleInput) SetPage(v int32) *ListAclRuleInput {
	s.Page = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *ListAclRuleInput) SetPageSize(v int32) *ListAclRuleInput {
	s.PageSize = &v
	return s
}

// SetRuleName sets the RuleName field's value.
func (s *ListAclRuleInput) SetRuleName(v string) *ListAclRuleInput {
	s.RuleName = &v
	return s
}

// SetRuleTag sets the RuleTag field's value.
func (s *ListAclRuleInput) SetRuleTag(v string) *ListAclRuleInput {
	s.RuleTag = &v
	return s
}

// SetTimeOrderBy sets the TimeOrderBy field's value.
func (s *ListAclRuleInput) SetTimeOrderBy(v string) *ListAclRuleInput {
	s.TimeOrderBy = &v
	return s
}

type ListAclRuleOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	Count *int32 `type:"int32" json:",omitempty"`

	CurrentPage *int32 `type:"int32" json:",omitempty"`

	PageSize *int32 `type:"int32" json:",omitempty"`

	Rules []*RuleForListAclRuleOutput `type:"list" json:",omitempty"`

	TotalCount *int32 `type:"int32" json:",omitempty"`
}

// String returns the string representation
func (s ListAclRuleOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListAclRuleOutput) GoString() string {
	return s.String()
}

// SetCount sets the Count field's value.
func (s *ListAclRuleOutput) SetCount(v int32) *ListAclRuleOutput {
	s.Count = &v
	return s
}

// SetCurrentPage sets the CurrentPage field's value.
func (s *ListAclRuleOutput) SetCurrentPage(v int32) *ListAclRuleOutput {
	s.CurrentPage = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *ListAclRuleOutput) SetPageSize(v int32) *ListAclRuleOutput {
	s.PageSize = &v
	return s
}

// SetRules sets the Rules field's value.
func (s *ListAclRuleOutput) SetRules(v []*RuleForListAclRuleOutput) *ListAclRuleOutput {
	s.Rules = v
	return s
}

// SetTotalCount sets the TotalCount field's value.
func (s *ListAclRuleOutput) SetTotalCount(v int32) *ListAclRuleOutput {
	s.TotalCount = &v
	return s
}

type RuleForListAclRuleOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	AccurateGroup *AccurateGroupForListAclRuleOutput `type:"structure" json:",omitempty"`

	Action *string `type:"string" json:",omitempty"`

	Advanced *int32 `type:"int32" json:",omitempty"`

	Description *string `type:"string" json:",omitempty"`

	Enable *int32 `type:"int32" json:",omitempty"`

	HostAddType *int32 `type:"int32" json:",omitempty"`

	HostGroupId []*int32 `type:"list" json:",omitempty"`

	HostGroups []*HostGroupForListAclRuleOutput `type:"list" json:",omitempty"`

	HostList []*string `type:"list" json:",omitempty"`

	ID *int32 `type:"int32" json:",omitempty"`

	IpGroups *int32 `type:"int32" json:",omitempty"`

	IpList []*string `type:"list" json:",omitempty"`

	IpLocationCountry []*string `type:"list" json:",omitempty"`

	Name []*string `type:"list" json:",omitempty"`

	PrefixSwitch []*string `type:"list" json:",omitempty"`

	RuleTag *string `type:"string" json:",omitempty"`

	UpdateTime *string `type:"string" json:",omitempty"`

	Url *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s RuleForListAclRuleOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s RuleForListAclRuleOutput) GoString() string {
	return s.String()
}

// SetAccurateGroup sets the AccurateGroup field's value.
func (s *RuleForListAclRuleOutput) SetAccurateGroup(v *AccurateGroupForListAclRuleOutput) *RuleForListAclRuleOutput {
	s.AccurateGroup = v
	return s
}

// SetAction sets the Action field's value.
func (s *RuleForListAclRuleOutput) SetAction(v string) *RuleForListAclRuleOutput {
	s.Action = &v
	return s
}

// SetAdvanced sets the Advanced field's value.
func (s *RuleForListAclRuleOutput) SetAdvanced(v int32) *RuleForListAclRuleOutput {
	s.Advanced = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *RuleForListAclRuleOutput) SetDescription(v string) *RuleForListAclRuleOutput {
	s.Description = &v
	return s
}

// SetEnable sets the Enable field's value.
func (s *RuleForListAclRuleOutput) SetEnable(v int32) *RuleForListAclRuleOutput {
	s.Enable = &v
	return s
}

// SetHostAddType sets the HostAddType field's value.
func (s *RuleForListAclRuleOutput) SetHostAddType(v int32) *RuleForListAclRuleOutput {
	s.HostAddType = &v
	return s
}

// SetHostGroupId sets the HostGroupId field's value.
func (s *RuleForListAclRuleOutput) SetHostGroupId(v []*int32) *RuleForListAclRuleOutput {
	s.HostGroupId = v
	return s
}

// SetHostGroups sets the HostGroups field's value.
func (s *RuleForListAclRuleOutput) SetHostGroups(v []*HostGroupForListAclRuleOutput) *RuleForListAclRuleOutput {
	s.HostGroups = v
	return s
}

// SetHostList sets the HostList field's value.
func (s *RuleForListAclRuleOutput) SetHostList(v []*string) *RuleForListAclRuleOutput {
	s.HostList = v
	return s
}

// SetID sets the ID field's value.
func (s *RuleForListAclRuleOutput) SetID(v int32) *RuleForListAclRuleOutput {
	s.ID = &v
	return s
}

// SetIpGroups sets the IpGroups field's value.
func (s *RuleForListAclRuleOutput) SetIpGroups(v int32) *RuleForListAclRuleOutput {
	s.IpGroups = &v
	return s
}

// SetIpList sets the IpList field's value.
func (s *RuleForListAclRuleOutput) SetIpList(v []*string) *RuleForListAclRuleOutput {
	s.IpList = v
	return s
}

// SetIpLocationCountry sets the IpLocationCountry field's value.
func (s *RuleForListAclRuleOutput) SetIpLocationCountry(v []*string) *RuleForListAclRuleOutput {
	s.IpLocationCountry = v
	return s
}

// SetName sets the Name field's value.
func (s *RuleForListAclRuleOutput) SetName(v []*string) *RuleForListAclRuleOutput {
	s.Name = v
	return s
}

// SetPrefixSwitch sets the PrefixSwitch field's value.
func (s *RuleForListAclRuleOutput) SetPrefixSwitch(v []*string) *RuleForListAclRuleOutput {
	s.PrefixSwitch = v
	return s
}

// SetRuleTag sets the RuleTag field's value.
func (s *RuleForListAclRuleOutput) SetRuleTag(v string) *RuleForListAclRuleOutput {
	s.RuleTag = &v
	return s
}

// SetUpdateTime sets the UpdateTime field's value.
func (s *RuleForListAclRuleOutput) SetUpdateTime(v string) *RuleForListAclRuleOutput {
	s.UpdateTime = &v
	return s
}

// SetUrl sets the Url field's value.
func (s *RuleForListAclRuleOutput) SetUrl(v string) *RuleForListAclRuleOutput {
	s.Url = &v
	return s
}

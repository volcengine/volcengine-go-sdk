// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package waf

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opListHostGroupCommon = "ListHostGroup"

// ListHostGroupCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ListHostGroupCommon operation. The "output" return
// value will be populated with the ListHostGroupCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListHostGroupCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListHostGroupCommon Send returns without error.
//
// See ListHostGroupCommon for more information on using the ListHostGroupCommon
// API call, and error handling.
//
//    // Example sending a request using the ListHostGroupCommonRequest method.
//    req, resp := client.ListHostGroupCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *WAF) ListHostGroupCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opListHostGroupCommon,
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

// ListHostGroupCommon API operation for WAF.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for WAF's
// API operation ListHostGroupCommon for usage and error information.
func (c *WAF) ListHostGroupCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ListHostGroupCommonRequest(input)
	return out, req.Send()
}

// ListHostGroupCommonWithContext is the same as ListHostGroupCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ListHostGroupCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *WAF) ListHostGroupCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ListHostGroupCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opListHostGroup = "ListHostGroup"

// ListHostGroupRequest generates a "volcengine/request.Request" representing the
// client's request for the ListHostGroup operation. The "output" return
// value will be populated with the ListHostGroupCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListHostGroupCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListHostGroupCommon Send returns without error.
//
// See ListHostGroup for more information on using the ListHostGroup
// API call, and error handling.
//
//    // Example sending a request using the ListHostGroupRequest method.
//    req, resp := client.ListHostGroupRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *WAF) ListHostGroupRequest(input *ListHostGroupInput) (req *request.Request, output *ListHostGroupOutput) {
	op := &request.Operation{
		Name:       opListHostGroup,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListHostGroupInput{}
	}

	output = &ListHostGroupOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// ListHostGroup API operation for WAF.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for WAF's
// API operation ListHostGroup for usage and error information.
func (c *WAF) ListHostGroup(input *ListHostGroupInput) (*ListHostGroupOutput, error) {
	req, out := c.ListHostGroupRequest(input)
	return out, req.Send()
}

// ListHostGroupWithContext is the same as ListHostGroup with the addition of
// the ability to pass a context and additional request options.
//
// See ListHostGroup for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *WAF) ListHostGroupWithContext(ctx volcengine.Context, input *ListHostGroupInput, opts ...request.Option) (*ListHostGroupOutput, error) {
	req, out := c.ListHostGroupRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type HostGroupListForListHostGroupOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Description *string `type:"string" json:",omitempty"`

	HostCount *int32 `type:"int32" json:",omitempty"`

	HostGroupId *int32 `type:"int32" json:",omitempty"`

	HostList []*string `type:"list" json:",omitempty"`

	Name *string `type:"string" json:",omitempty"`

	RelatedRules []*RelatedRuleForListHostGroupOutput `type:"list" json:",omitempty"`

	UpdateTime *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s HostGroupListForListHostGroupOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s HostGroupListForListHostGroupOutput) GoString() string {
	return s.String()
}

// SetDescription sets the Description field's value.
func (s *HostGroupListForListHostGroupOutput) SetDescription(v string) *HostGroupListForListHostGroupOutput {
	s.Description = &v
	return s
}

// SetHostCount sets the HostCount field's value.
func (s *HostGroupListForListHostGroupOutput) SetHostCount(v int32) *HostGroupListForListHostGroupOutput {
	s.HostCount = &v
	return s
}

// SetHostGroupId sets the HostGroupId field's value.
func (s *HostGroupListForListHostGroupOutput) SetHostGroupId(v int32) *HostGroupListForListHostGroupOutput {
	s.HostGroupId = &v
	return s
}

// SetHostList sets the HostList field's value.
func (s *HostGroupListForListHostGroupOutput) SetHostList(v []*string) *HostGroupListForListHostGroupOutput {
	s.HostList = v
	return s
}

// SetName sets the Name field's value.
func (s *HostGroupListForListHostGroupOutput) SetName(v string) *HostGroupListForListHostGroupOutput {
	s.Name = &v
	return s
}

// SetRelatedRules sets the RelatedRules field's value.
func (s *HostGroupListForListHostGroupOutput) SetRelatedRules(v []*RelatedRuleForListHostGroupOutput) *HostGroupListForListHostGroupOutput {
	s.RelatedRules = v
	return s
}

// SetUpdateTime sets the UpdateTime field's value.
func (s *HostGroupListForListHostGroupOutput) SetUpdateTime(v string) *HostGroupListForListHostGroupOutput {
	s.UpdateTime = &v
	return s
}

type ListHostGroupInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	HostFix *string `type:"string" json:",omitempty"`

	HostGroupID *int32 `type:"int32" json:",omitempty"`

	ListAll *bool `type:"boolean" json:",omitempty"`

	NameFix *string `type:"string" json:",omitempty"`

	Page *int32 `type:"int32" json:",omitempty"`

	PageSize *int32 `type:"int32" json:",omitempty"`

	RuleTag *string `type:"string" json:",omitempty"`

	TimeOrderBy *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s ListHostGroupInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListHostGroupInput) GoString() string {
	return s.String()
}

// SetHostFix sets the HostFix field's value.
func (s *ListHostGroupInput) SetHostFix(v string) *ListHostGroupInput {
	s.HostFix = &v
	return s
}

// SetHostGroupID sets the HostGroupID field's value.
func (s *ListHostGroupInput) SetHostGroupID(v int32) *ListHostGroupInput {
	s.HostGroupID = &v
	return s
}

// SetListAll sets the ListAll field's value.
func (s *ListHostGroupInput) SetListAll(v bool) *ListHostGroupInput {
	s.ListAll = &v
	return s
}

// SetNameFix sets the NameFix field's value.
func (s *ListHostGroupInput) SetNameFix(v string) *ListHostGroupInput {
	s.NameFix = &v
	return s
}

// SetPage sets the Page field's value.
func (s *ListHostGroupInput) SetPage(v int32) *ListHostGroupInput {
	s.Page = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *ListHostGroupInput) SetPageSize(v int32) *ListHostGroupInput {
	s.PageSize = &v
	return s
}

// SetRuleTag sets the RuleTag field's value.
func (s *ListHostGroupInput) SetRuleTag(v string) *ListHostGroupInput {
	s.RuleTag = &v
	return s
}

// SetTimeOrderBy sets the TimeOrderBy field's value.
func (s *ListHostGroupInput) SetTimeOrderBy(v string) *ListHostGroupInput {
	s.TimeOrderBy = &v
	return s
}

type ListHostGroupOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	HostGroupList []*HostGroupListForListHostGroupOutput `type:"list" json:",omitempty"`
}

// String returns the string representation
func (s ListHostGroupOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListHostGroupOutput) GoString() string {
	return s.String()
}

// SetHostGroupList sets the HostGroupList field's value.
func (s *ListHostGroupOutput) SetHostGroupList(v []*HostGroupListForListHostGroupOutput) *ListHostGroupOutput {
	s.HostGroupList = v
	return s
}

type RelatedRuleForListHostGroupOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Host *string `type:"string" json:",omitempty"`

	RuleName *string `type:"string" json:",omitempty"`

	RuleTag *string `type:"string" json:",omitempty"`

	RuleType *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s RelatedRuleForListHostGroupOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s RelatedRuleForListHostGroupOutput) GoString() string {
	return s.String()
}

// SetHost sets the Host field's value.
func (s *RelatedRuleForListHostGroupOutput) SetHost(v string) *RelatedRuleForListHostGroupOutput {
	s.Host = &v
	return s
}

// SetRuleName sets the RuleName field's value.
func (s *RelatedRuleForListHostGroupOutput) SetRuleName(v string) *RelatedRuleForListHostGroupOutput {
	s.RuleName = &v
	return s
}

// SetRuleTag sets the RuleTag field's value.
func (s *RelatedRuleForListHostGroupOutput) SetRuleTag(v string) *RelatedRuleForListHostGroupOutput {
	s.RuleTag = &v
	return s
}

// SetRuleType sets the RuleType field's value.
func (s *RelatedRuleForListHostGroupOutput) SetRuleType(v string) *RelatedRuleForListHostGroupOutput {
	s.RuleType = &v
	return s
}
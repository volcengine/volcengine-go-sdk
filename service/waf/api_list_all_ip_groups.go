// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package waf

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opListAllIpGroupsCommon = "ListAllIpGroups"

// ListAllIpGroupsCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ListAllIpGroupsCommon operation. The "output" return
// value will be populated with the ListAllIpGroupsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListAllIpGroupsCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListAllIpGroupsCommon Send returns without error.
//
// See ListAllIpGroupsCommon for more information on using the ListAllIpGroupsCommon
// API call, and error handling.
//
//    // Example sending a request using the ListAllIpGroupsCommonRequest method.
//    req, resp := client.ListAllIpGroupsCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *WAF) ListAllIpGroupsCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opListAllIpGroupsCommon,
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

// ListAllIpGroupsCommon API operation for WAF.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for WAF's
// API operation ListAllIpGroupsCommon for usage and error information.
func (c *WAF) ListAllIpGroupsCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ListAllIpGroupsCommonRequest(input)
	return out, req.Send()
}

// ListAllIpGroupsCommonWithContext is the same as ListAllIpGroupsCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ListAllIpGroupsCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *WAF) ListAllIpGroupsCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ListAllIpGroupsCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opListAllIpGroups = "ListAllIpGroups"

// ListAllIpGroupsRequest generates a "volcengine/request.Request" representing the
// client's request for the ListAllIpGroups operation. The "output" return
// value will be populated with the ListAllIpGroupsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListAllIpGroupsCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListAllIpGroupsCommon Send returns without error.
//
// See ListAllIpGroups for more information on using the ListAllIpGroups
// API call, and error handling.
//
//    // Example sending a request using the ListAllIpGroupsRequest method.
//    req, resp := client.ListAllIpGroupsRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *WAF) ListAllIpGroupsRequest(input *ListAllIpGroupsInput) (req *request.Request, output *ListAllIpGroupsOutput) {
	op := &request.Operation{
		Name:       opListAllIpGroups,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListAllIpGroupsInput{}
	}

	output = &ListAllIpGroupsOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// ListAllIpGroups API operation for WAF.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for WAF's
// API operation ListAllIpGroups for usage and error information.
func (c *WAF) ListAllIpGroups(input *ListAllIpGroupsInput) (*ListAllIpGroupsOutput, error) {
	req, out := c.ListAllIpGroupsRequest(input)
	return out, req.Send()
}

// ListAllIpGroupsWithContext is the same as ListAllIpGroups with the addition of
// the ability to pass a context and additional request options.
//
// See ListAllIpGroups for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *WAF) ListAllIpGroupsWithContext(ctx volcengine.Context, input *ListAllIpGroupsInput, opts ...request.Option) (*ListAllIpGroupsOutput, error) {
	req, out := c.ListAllIpGroupsRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type IpGroupListForListAllIpGroupsOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	IpCount *int32 `type:"int32" json:",omitempty"`

	IpGroupId *int32 `type:"int32" json:",omitempty"`

	Name *string `type:"string" json:",omitempty"`

	RelatedRules []*RelatedRuleForListAllIpGroupsOutput `type:"list" json:",omitempty"`

	UpdateTime *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s IpGroupListForListAllIpGroupsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s IpGroupListForListAllIpGroupsOutput) GoString() string {
	return s.String()
}

// SetIpCount sets the IpCount field's value.
func (s *IpGroupListForListAllIpGroupsOutput) SetIpCount(v int32) *IpGroupListForListAllIpGroupsOutput {
	s.IpCount = &v
	return s
}

// SetIpGroupId sets the IpGroupId field's value.
func (s *IpGroupListForListAllIpGroupsOutput) SetIpGroupId(v int32) *IpGroupListForListAllIpGroupsOutput {
	s.IpGroupId = &v
	return s
}

// SetName sets the Name field's value.
func (s *IpGroupListForListAllIpGroupsOutput) SetName(v string) *IpGroupListForListAllIpGroupsOutput {
	s.Name = &v
	return s
}

// SetRelatedRules sets the RelatedRules field's value.
func (s *IpGroupListForListAllIpGroupsOutput) SetRelatedRules(v []*RelatedRuleForListAllIpGroupsOutput) *IpGroupListForListAllIpGroupsOutput {
	s.RelatedRules = v
	return s
}

// SetUpdateTime sets the UpdateTime field's value.
func (s *IpGroupListForListAllIpGroupsOutput) SetUpdateTime(v string) *IpGroupListForListAllIpGroupsOutput {
	s.UpdateTime = &v
	return s
}

type ListAllIpGroupsInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Ip *string `type:"string" json:",omitempty"`

	Page *string `type:"string" json:",omitempty"`

	PageSize *int32 `type:"int32" json:",omitempty"`

	ProjectName *string `type:"string" json:",omitempty"`

	RuleTag *string `type:"string" json:",omitempty"`

	// TimeOrderBy is a required field
	TimeOrderBy *string `type:"string" json:",omitempty" required:"true" enum:"EnumOfTimeOrderByForListAllIpGroupsInput"`
}

// String returns the string representation
func (s ListAllIpGroupsInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListAllIpGroupsInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListAllIpGroupsInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ListAllIpGroupsInput"}
	if s.TimeOrderBy == nil {
		invalidParams.Add(request.NewErrParamRequired("TimeOrderBy"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetIp sets the Ip field's value.
func (s *ListAllIpGroupsInput) SetIp(v string) *ListAllIpGroupsInput {
	s.Ip = &v
	return s
}

// SetPage sets the Page field's value.
func (s *ListAllIpGroupsInput) SetPage(v string) *ListAllIpGroupsInput {
	s.Page = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *ListAllIpGroupsInput) SetPageSize(v int32) *ListAllIpGroupsInput {
	s.PageSize = &v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *ListAllIpGroupsInput) SetProjectName(v string) *ListAllIpGroupsInput {
	s.ProjectName = &v
	return s
}

// SetRuleTag sets the RuleTag field's value.
func (s *ListAllIpGroupsInput) SetRuleTag(v string) *ListAllIpGroupsInput {
	s.RuleTag = &v
	return s
}

// SetTimeOrderBy sets the TimeOrderBy field's value.
func (s *ListAllIpGroupsInput) SetTimeOrderBy(v string) *ListAllIpGroupsInput {
	s.TimeOrderBy = &v
	return s
}

type ListAllIpGroupsOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	Count *int32 `type:"int32" json:",omitempty"`

	IpGroupCount *int32 `type:"int32" json:",omitempty"`

	IpGroupList []*IpGroupListForListAllIpGroupsOutput `type:"list" json:",omitempty"`

	IpGroupQuota *int32 `type:"int32" json:",omitempty"`

	IpLimitQuota *int32 `type:"int32" json:",omitempty"`

	PageNumber *int32 `type:"int32" json:",omitempty"`

	PageSize *int32 `type:"int32" json:",omitempty"`

	TotalCount *int32 `type:"int32" json:",omitempty"`
}

// String returns the string representation
func (s ListAllIpGroupsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListAllIpGroupsOutput) GoString() string {
	return s.String()
}

// SetCount sets the Count field's value.
func (s *ListAllIpGroupsOutput) SetCount(v int32) *ListAllIpGroupsOutput {
	s.Count = &v
	return s
}

// SetIpGroupCount sets the IpGroupCount field's value.
func (s *ListAllIpGroupsOutput) SetIpGroupCount(v int32) *ListAllIpGroupsOutput {
	s.IpGroupCount = &v
	return s
}

// SetIpGroupList sets the IpGroupList field's value.
func (s *ListAllIpGroupsOutput) SetIpGroupList(v []*IpGroupListForListAllIpGroupsOutput) *ListAllIpGroupsOutput {
	s.IpGroupList = v
	return s
}

// SetIpGroupQuota sets the IpGroupQuota field's value.
func (s *ListAllIpGroupsOutput) SetIpGroupQuota(v int32) *ListAllIpGroupsOutput {
	s.IpGroupQuota = &v
	return s
}

// SetIpLimitQuota sets the IpLimitQuota field's value.
func (s *ListAllIpGroupsOutput) SetIpLimitQuota(v int32) *ListAllIpGroupsOutput {
	s.IpLimitQuota = &v
	return s
}

// SetPageNumber sets the PageNumber field's value.
func (s *ListAllIpGroupsOutput) SetPageNumber(v int32) *ListAllIpGroupsOutput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *ListAllIpGroupsOutput) SetPageSize(v int32) *ListAllIpGroupsOutput {
	s.PageSize = &v
	return s
}

// SetTotalCount sets the TotalCount field's value.
func (s *ListAllIpGroupsOutput) SetTotalCount(v int32) *ListAllIpGroupsOutput {
	s.TotalCount = &v
	return s
}

type RelatedRuleForListAllIpGroupsOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Host *string `type:"string" json:",omitempty"`

	RuleName *string `type:"string" json:",omitempty"`

	RuleTag *string `type:"string" json:",omitempty"`

	RuleType *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s RelatedRuleForListAllIpGroupsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s RelatedRuleForListAllIpGroupsOutput) GoString() string {
	return s.String()
}

// SetHost sets the Host field's value.
func (s *RelatedRuleForListAllIpGroupsOutput) SetHost(v string) *RelatedRuleForListAllIpGroupsOutput {
	s.Host = &v
	return s
}

// SetRuleName sets the RuleName field's value.
func (s *RelatedRuleForListAllIpGroupsOutput) SetRuleName(v string) *RelatedRuleForListAllIpGroupsOutput {
	s.RuleName = &v
	return s
}

// SetRuleTag sets the RuleTag field's value.
func (s *RelatedRuleForListAllIpGroupsOutput) SetRuleTag(v string) *RelatedRuleForListAllIpGroupsOutput {
	s.RuleTag = &v
	return s
}

// SetRuleType sets the RuleType field's value.
func (s *RelatedRuleForListAllIpGroupsOutput) SetRuleType(v string) *RelatedRuleForListAllIpGroupsOutput {
	s.RuleType = &v
	return s
}

const (
	// EnumOfTimeOrderByForListAllIpGroupsInputAsc is a EnumOfTimeOrderByForListAllIpGroupsInput enum value
	EnumOfTimeOrderByForListAllIpGroupsInputAsc = "ASC"

	// EnumOfTimeOrderByForListAllIpGroupsInputDesc is a EnumOfTimeOrderByForListAllIpGroupsInput enum value
	EnumOfTimeOrderByForListAllIpGroupsInputDesc = "DESC"
)

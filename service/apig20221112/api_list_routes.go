// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package apig20221112

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opListRoutesCommon = "ListRoutes"

// ListRoutesCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ListRoutesCommon operation. The "output" return
// value will be populated with the ListRoutesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListRoutesCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListRoutesCommon Send returns without error.
//
// See ListRoutesCommon for more information on using the ListRoutesCommon
// API call, and error handling.
//
//    // Example sending a request using the ListRoutesCommonRequest method.
//    req, resp := client.ListRoutesCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *APIG20221112) ListRoutesCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opListRoutesCommon,
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

// ListRoutesCommon API operation for APIG20221112.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for APIG20221112's
// API operation ListRoutesCommon for usage and error information.
func (c *APIG20221112) ListRoutesCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ListRoutesCommonRequest(input)
	return out, req.Send()
}

// ListRoutesCommonWithContext is the same as ListRoutesCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ListRoutesCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *APIG20221112) ListRoutesCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ListRoutesCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opListRoutes = "ListRoutes"

// ListRoutesRequest generates a "volcengine/request.Request" representing the
// client's request for the ListRoutes operation. The "output" return
// value will be populated with the ListRoutesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListRoutesCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListRoutesCommon Send returns without error.
//
// See ListRoutes for more information on using the ListRoutes
// API call, and error handling.
//
//    // Example sending a request using the ListRoutesRequest method.
//    req, resp := client.ListRoutesRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *APIG20221112) ListRoutesRequest(input *ListRoutesInput) (req *request.Request, output *ListRoutesOutput) {
	op := &request.Operation{
		Name:       opListRoutes,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListRoutesInput{}
	}

	output = &ListRoutesOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// ListRoutes API operation for APIG20221112.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for APIG20221112's
// API operation ListRoutes for usage and error information.
func (c *APIG20221112) ListRoutes(input *ListRoutesInput) (*ListRoutesOutput, error) {
	req, out := c.ListRoutesRequest(input)
	return out, req.Send()
}

// ListRoutesWithContext is the same as ListRoutes with the addition of
// the ability to pass a context and additional request options.
//
// See ListRoutes for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *APIG20221112) ListRoutesWithContext(ctx volcengine.Context, input *ListRoutesInput, opts ...request.Option) (*ListRoutesOutput, error) {
	req, out := c.ListRoutesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type AIProviderSettingsForListRoutesOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Model *string `type:"string" json:",omitempty"`

	TargetPath *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s AIProviderSettingsForListRoutesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s AIProviderSettingsForListRoutesOutput) GoString() string {
	return s.String()
}

// SetModel sets the Model field's value.
func (s *AIProviderSettingsForListRoutesOutput) SetModel(v string) *AIProviderSettingsForListRoutesOutput {
	s.Model = &v
	return s
}

// SetTargetPath sets the TargetPath field's value.
func (s *AIProviderSettingsForListRoutesOutput) SetTargetPath(v string) *AIProviderSettingsForListRoutesOutput {
	s.TargetPath = &v
	return s
}

type AdvancedSettingForListRoutesOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	CorsPolicySetting *CorsPolicySettingForListRoutesOutput `type:"structure" json:",omitempty"`

	HeaderOperations []*HeaderOperationForListRoutesOutput `type:"list" json:",omitempty"`

	MirrorPolicies []*MirrorPolicyForListRoutesOutput `type:"list" json:",omitempty"`

	RetryPolicySetting *RetryPolicySettingForListRoutesOutput `type:"structure" json:",omitempty"`

	TimeoutSetting *TimeoutSettingForListRoutesOutput `type:"structure" json:",omitempty"`

	URLRewriteSetting *URLRewriteSettingForListRoutesOutput `type:"structure" json:",omitempty"`
}

// String returns the string representation
func (s AdvancedSettingForListRoutesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s AdvancedSettingForListRoutesOutput) GoString() string {
	return s.String()
}

// SetCorsPolicySetting sets the CorsPolicySetting field's value.
func (s *AdvancedSettingForListRoutesOutput) SetCorsPolicySetting(v *CorsPolicySettingForListRoutesOutput) *AdvancedSettingForListRoutesOutput {
	s.CorsPolicySetting = v
	return s
}

// SetHeaderOperations sets the HeaderOperations field's value.
func (s *AdvancedSettingForListRoutesOutput) SetHeaderOperations(v []*HeaderOperationForListRoutesOutput) *AdvancedSettingForListRoutesOutput {
	s.HeaderOperations = v
	return s
}

// SetMirrorPolicies sets the MirrorPolicies field's value.
func (s *AdvancedSettingForListRoutesOutput) SetMirrorPolicies(v []*MirrorPolicyForListRoutesOutput) *AdvancedSettingForListRoutesOutput {
	s.MirrorPolicies = v
	return s
}

// SetRetryPolicySetting sets the RetryPolicySetting field's value.
func (s *AdvancedSettingForListRoutesOutput) SetRetryPolicySetting(v *RetryPolicySettingForListRoutesOutput) *AdvancedSettingForListRoutesOutput {
	s.RetryPolicySetting = v
	return s
}

// SetTimeoutSetting sets the TimeoutSetting field's value.
func (s *AdvancedSettingForListRoutesOutput) SetTimeoutSetting(v *TimeoutSettingForListRoutesOutput) *AdvancedSettingForListRoutesOutput {
	s.TimeoutSetting = v
	return s
}

// SetURLRewriteSetting sets the URLRewriteSetting field's value.
func (s *AdvancedSettingForListRoutesOutput) SetURLRewriteSetting(v *URLRewriteSettingForListRoutesOutput) *AdvancedSettingForListRoutesOutput {
	s.URLRewriteSetting = v
	return s
}

type CorsPolicySettingForListRoutesOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Enable *bool `type:"boolean" json:",omitempty"`
}

// String returns the string representation
func (s CorsPolicySettingForListRoutesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CorsPolicySettingForListRoutesOutput) GoString() string {
	return s.String()
}

// SetEnable sets the Enable field's value.
func (s *CorsPolicySettingForListRoutesOutput) SetEnable(v bool) *CorsPolicySettingForListRoutesOutput {
	s.Enable = &v
	return s
}

type CustomDomainForListRoutesOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Domain *string `type:"string" json:",omitempty"`

	Id *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s CustomDomainForListRoutesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CustomDomainForListRoutesOutput) GoString() string {
	return s.String()
}

// SetDomain sets the Domain field's value.
func (s *CustomDomainForListRoutesOutput) SetDomain(v string) *CustomDomainForListRoutesOutput {
	s.Domain = &v
	return s
}

// SetId sets the Id field's value.
func (s *CustomDomainForListRoutesOutput) SetId(v string) *CustomDomainForListRoutesOutput {
	s.Id = &v
	return s
}

type DomainForListRoutesOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Domain *string `type:"string" json:",omitempty"`

	Type *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s DomainForListRoutesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DomainForListRoutesOutput) GoString() string {
	return s.String()
}

// SetDomain sets the Domain field's value.
func (s *DomainForListRoutesOutput) SetDomain(v string) *DomainForListRoutesOutput {
	s.Domain = &v
	return s
}

// SetType sets the Type field's value.
func (s *DomainForListRoutesOutput) SetType(v string) *DomainForListRoutesOutput {
	s.Type = &v
	return s
}

type FilterForListRoutesInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Name *string `type:"string" json:",omitempty"`

	Path *string `type:"string" json:",omitempty"`

	UpstreamIds []*string `type:"list" json:",omitempty"`
}

// String returns the string representation
func (s FilterForListRoutesInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s FilterForListRoutesInput) GoString() string {
	return s.String()
}

// SetName sets the Name field's value.
func (s *FilterForListRoutesInput) SetName(v string) *FilterForListRoutesInput {
	s.Name = &v
	return s
}

// SetPath sets the Path field's value.
func (s *FilterForListRoutesInput) SetPath(v string) *FilterForListRoutesInput {
	s.Path = &v
	return s
}

// SetUpstreamIds sets the UpstreamIds field's value.
func (s *FilterForListRoutesInput) SetUpstreamIds(v []*string) *FilterForListRoutesInput {
	s.UpstreamIds = v
	return s
}

type HeaderForListRoutesOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Key *string `type:"string" json:",omitempty"`

	Value *ValueForListRoutesOutput `type:"structure" json:",omitempty"`
}

// String returns the string representation
func (s HeaderForListRoutesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s HeaderForListRoutesOutput) GoString() string {
	return s.String()
}

// SetKey sets the Key field's value.
func (s *HeaderForListRoutesOutput) SetKey(v string) *HeaderForListRoutesOutput {
	s.Key = &v
	return s
}

// SetValue sets the Value field's value.
func (s *HeaderForListRoutesOutput) SetValue(v *ValueForListRoutesOutput) *HeaderForListRoutesOutput {
	s.Value = v
	return s
}

type HeaderOperationForListRoutesOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	DirectionType *string `type:"string" json:",omitempty"`

	Key *string `type:"string" json:",omitempty"`

	Operation *string `type:"string" json:",omitempty"`

	Value *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s HeaderOperationForListRoutesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s HeaderOperationForListRoutesOutput) GoString() string {
	return s.String()
}

// SetDirectionType sets the DirectionType field's value.
func (s *HeaderOperationForListRoutesOutput) SetDirectionType(v string) *HeaderOperationForListRoutesOutput {
	s.DirectionType = &v
	return s
}

// SetKey sets the Key field's value.
func (s *HeaderOperationForListRoutesOutput) SetKey(v string) *HeaderOperationForListRoutesOutput {
	s.Key = &v
	return s
}

// SetOperation sets the Operation field's value.
func (s *HeaderOperationForListRoutesOutput) SetOperation(v string) *HeaderOperationForListRoutesOutput {
	s.Operation = &v
	return s
}

// SetValue sets the Value field's value.
func (s *HeaderOperationForListRoutesOutput) SetValue(v string) *HeaderOperationForListRoutesOutput {
	s.Value = &v
	return s
}

type ItemForListRoutesOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	AdvancedSetting *AdvancedSettingForListRoutesOutput `type:"structure" json:",omitempty"`

	CreateTime *string `type:"string" json:",omitempty"`

	CustomDomains []*CustomDomainForListRoutesOutput `type:"list" json:",omitempty"`

	Domains []*DomainForListRoutesOutput `type:"list" json:",omitempty"`

	Enable *bool `type:"boolean" json:",omitempty"`

	Id *string `type:"string" json:",omitempty"`

	MatchRule *MatchRuleForListRoutesOutput `type:"structure" json:",omitempty"`

	Name *string `type:"string" json:",omitempty"`

	Priority *int64 `type:"int64" json:",omitempty"`

	Reason *string `type:"string" json:",omitempty"`

	ResourceType *string `type:"string" json:",omitempty"`

	ServiceId *string `type:"string" json:",omitempty"`

	ServiceName *string `type:"string" json:",omitempty"`

	Status *string `type:"string" json:",omitempty"`

	Tags []*TagForListRoutesOutput `type:"list" json:",omitempty"`

	UpdateTime *string `type:"string" json:",omitempty"`

	UpstreamList []*UpstreamListForListRoutesOutput `type:"list" json:",omitempty"`
}

// String returns the string representation
func (s ItemForListRoutesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ItemForListRoutesOutput) GoString() string {
	return s.String()
}

// SetAdvancedSetting sets the AdvancedSetting field's value.
func (s *ItemForListRoutesOutput) SetAdvancedSetting(v *AdvancedSettingForListRoutesOutput) *ItemForListRoutesOutput {
	s.AdvancedSetting = v
	return s
}

// SetCreateTime sets the CreateTime field's value.
func (s *ItemForListRoutesOutput) SetCreateTime(v string) *ItemForListRoutesOutput {
	s.CreateTime = &v
	return s
}

// SetCustomDomains sets the CustomDomains field's value.
func (s *ItemForListRoutesOutput) SetCustomDomains(v []*CustomDomainForListRoutesOutput) *ItemForListRoutesOutput {
	s.CustomDomains = v
	return s
}

// SetDomains sets the Domains field's value.
func (s *ItemForListRoutesOutput) SetDomains(v []*DomainForListRoutesOutput) *ItemForListRoutesOutput {
	s.Domains = v
	return s
}

// SetEnable sets the Enable field's value.
func (s *ItemForListRoutesOutput) SetEnable(v bool) *ItemForListRoutesOutput {
	s.Enable = &v
	return s
}

// SetId sets the Id field's value.
func (s *ItemForListRoutesOutput) SetId(v string) *ItemForListRoutesOutput {
	s.Id = &v
	return s
}

// SetMatchRule sets the MatchRule field's value.
func (s *ItemForListRoutesOutput) SetMatchRule(v *MatchRuleForListRoutesOutput) *ItemForListRoutesOutput {
	s.MatchRule = v
	return s
}

// SetName sets the Name field's value.
func (s *ItemForListRoutesOutput) SetName(v string) *ItemForListRoutesOutput {
	s.Name = &v
	return s
}

// SetPriority sets the Priority field's value.
func (s *ItemForListRoutesOutput) SetPriority(v int64) *ItemForListRoutesOutput {
	s.Priority = &v
	return s
}

// SetReason sets the Reason field's value.
func (s *ItemForListRoutesOutput) SetReason(v string) *ItemForListRoutesOutput {
	s.Reason = &v
	return s
}

// SetResourceType sets the ResourceType field's value.
func (s *ItemForListRoutesOutput) SetResourceType(v string) *ItemForListRoutesOutput {
	s.ResourceType = &v
	return s
}

// SetServiceId sets the ServiceId field's value.
func (s *ItemForListRoutesOutput) SetServiceId(v string) *ItemForListRoutesOutput {
	s.ServiceId = &v
	return s
}

// SetServiceName sets the ServiceName field's value.
func (s *ItemForListRoutesOutput) SetServiceName(v string) *ItemForListRoutesOutput {
	s.ServiceName = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *ItemForListRoutesOutput) SetStatus(v string) *ItemForListRoutesOutput {
	s.Status = &v
	return s
}

// SetTags sets the Tags field's value.
func (s *ItemForListRoutesOutput) SetTags(v []*TagForListRoutesOutput) *ItemForListRoutesOutput {
	s.Tags = v
	return s
}

// SetUpdateTime sets the UpdateTime field's value.
func (s *ItemForListRoutesOutput) SetUpdateTime(v string) *ItemForListRoutesOutput {
	s.UpdateTime = &v
	return s
}

// SetUpstreamList sets the UpstreamList field's value.
func (s *ItemForListRoutesOutput) SetUpstreamList(v []*UpstreamListForListRoutesOutput) *ItemForListRoutesOutput {
	s.UpstreamList = v
	return s
}

type ListRoutesInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Filter *FilterForListRoutesInput `type:"structure" json:",omitempty"`

	GatewayId *string `type:"string" json:",omitempty"`

	PageNumber *int64 `type:"int64" json:",omitempty"`

	PageSize *int64 `type:"int64" json:",omitempty"`

	ResourceType *string `type:"string" json:",omitempty"`

	ServiceId *string `type:"string" json:",omitempty"`

	UpstreamId *string `type:"string" json:",omitempty"`

	UpstreamVersion *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s ListRoutesInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListRoutesInput) GoString() string {
	return s.String()
}

// SetFilter sets the Filter field's value.
func (s *ListRoutesInput) SetFilter(v *FilterForListRoutesInput) *ListRoutesInput {
	s.Filter = v
	return s
}

// SetGatewayId sets the GatewayId field's value.
func (s *ListRoutesInput) SetGatewayId(v string) *ListRoutesInput {
	s.GatewayId = &v
	return s
}

// SetPageNumber sets the PageNumber field's value.
func (s *ListRoutesInput) SetPageNumber(v int64) *ListRoutesInput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *ListRoutesInput) SetPageSize(v int64) *ListRoutesInput {
	s.PageSize = &v
	return s
}

// SetResourceType sets the ResourceType field's value.
func (s *ListRoutesInput) SetResourceType(v string) *ListRoutesInput {
	s.ResourceType = &v
	return s
}

// SetServiceId sets the ServiceId field's value.
func (s *ListRoutesInput) SetServiceId(v string) *ListRoutesInput {
	s.ServiceId = &v
	return s
}

// SetUpstreamId sets the UpstreamId field's value.
func (s *ListRoutesInput) SetUpstreamId(v string) *ListRoutesInput {
	s.UpstreamId = &v
	return s
}

// SetUpstreamVersion sets the UpstreamVersion field's value.
func (s *ListRoutesInput) SetUpstreamVersion(v string) *ListRoutesInput {
	s.UpstreamVersion = &v
	return s
}

type ListRoutesOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	Items []*ItemForListRoutesOutput `type:"list" json:",omitempty"`

	Total *int64 `type:"int64" json:",omitempty"`
}

// String returns the string representation
func (s ListRoutesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListRoutesOutput) GoString() string {
	return s.String()
}

// SetItems sets the Items field's value.
func (s *ListRoutesOutput) SetItems(v []*ItemForListRoutesOutput) *ListRoutesOutput {
	s.Items = v
	return s
}

// SetTotal sets the Total field's value.
func (s *ListRoutesOutput) SetTotal(v int64) *ListRoutesOutput {
	s.Total = &v
	return s
}

type MatchRuleForListRoutesOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Header []*HeaderForListRoutesOutput `type:"list" json:",omitempty"`

	Method []*string `type:"list" json:",omitempty"`

	Path *PathForListRoutesOutput `type:"structure" json:",omitempty"`

	QueryString []*QueryStringForListRoutesOutput `type:"list" json:",omitempty"`
}

// String returns the string representation
func (s MatchRuleForListRoutesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s MatchRuleForListRoutesOutput) GoString() string {
	return s.String()
}

// SetHeader sets the Header field's value.
func (s *MatchRuleForListRoutesOutput) SetHeader(v []*HeaderForListRoutesOutput) *MatchRuleForListRoutesOutput {
	s.Header = v
	return s
}

// SetMethod sets the Method field's value.
func (s *MatchRuleForListRoutesOutput) SetMethod(v []*string) *MatchRuleForListRoutesOutput {
	s.Method = v
	return s
}

// SetPath sets the Path field's value.
func (s *MatchRuleForListRoutesOutput) SetPath(v *PathForListRoutesOutput) *MatchRuleForListRoutesOutput {
	s.Path = v
	return s
}

// SetQueryString sets the QueryString field's value.
func (s *MatchRuleForListRoutesOutput) SetQueryString(v []*QueryStringForListRoutesOutput) *MatchRuleForListRoutesOutput {
	s.QueryString = v
	return s
}

type MirrorPolicyForListRoutesOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Percent *PercentForListRoutesOutput `type:"structure" json:",omitempty"`

	Upstream *UpstreamForListRoutesOutput `type:"structure" json:",omitempty"`
}

// String returns the string representation
func (s MirrorPolicyForListRoutesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s MirrorPolicyForListRoutesOutput) GoString() string {
	return s.String()
}

// SetPercent sets the Percent field's value.
func (s *MirrorPolicyForListRoutesOutput) SetPercent(v *PercentForListRoutesOutput) *MirrorPolicyForListRoutesOutput {
	s.Percent = v
	return s
}

// SetUpstream sets the Upstream field's value.
func (s *MirrorPolicyForListRoutesOutput) SetUpstream(v *UpstreamForListRoutesOutput) *MirrorPolicyForListRoutesOutput {
	s.Upstream = v
	return s
}

type PathForListRoutesOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	MatchContent *string `type:"string" json:",omitempty"`

	MatchType *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s PathForListRoutesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s PathForListRoutesOutput) GoString() string {
	return s.String()
}

// SetMatchContent sets the MatchContent field's value.
func (s *PathForListRoutesOutput) SetMatchContent(v string) *PathForListRoutesOutput {
	s.MatchContent = &v
	return s
}

// SetMatchType sets the MatchType field's value.
func (s *PathForListRoutesOutput) SetMatchType(v string) *PathForListRoutesOutput {
	s.MatchType = &v
	return s
}

type PercentForListRoutesOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Value *float64 `type:"double" json:",omitempty"`
}

// String returns the string representation
func (s PercentForListRoutesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s PercentForListRoutesOutput) GoString() string {
	return s.String()
}

// SetValue sets the Value field's value.
func (s *PercentForListRoutesOutput) SetValue(v float64) *PercentForListRoutesOutput {
	s.Value = &v
	return s
}

type QueryStringForListRoutesOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Key *string `type:"string" json:",omitempty"`

	Value *PathForListRoutesOutput `type:"structure" json:",omitempty"`
}

// String returns the string representation
func (s QueryStringForListRoutesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s QueryStringForListRoutesOutput) GoString() string {
	return s.String()
}

// SetKey sets the Key field's value.
func (s *QueryStringForListRoutesOutput) SetKey(v string) *QueryStringForListRoutesOutput {
	s.Key = &v
	return s
}

// SetValue sets the Value field's value.
func (s *QueryStringForListRoutesOutput) SetValue(v *PathForListRoutesOutput) *QueryStringForListRoutesOutput {
	s.Value = v
	return s
}

type RetryPolicySettingForListRoutesOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Attempts *int64 `type:"int64" json:",omitempty"`

	Enable *bool `type:"boolean" json:",omitempty"`

	HttpCodes []*string `type:"list" json:",omitempty"`

	PerTryTimeout *int64 `type:"int64" json:",omitempty"`

	RetryOn []*string `type:"list" json:",omitempty"`
}

// String returns the string representation
func (s RetryPolicySettingForListRoutesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s RetryPolicySettingForListRoutesOutput) GoString() string {
	return s.String()
}

// SetAttempts sets the Attempts field's value.
func (s *RetryPolicySettingForListRoutesOutput) SetAttempts(v int64) *RetryPolicySettingForListRoutesOutput {
	s.Attempts = &v
	return s
}

// SetEnable sets the Enable field's value.
func (s *RetryPolicySettingForListRoutesOutput) SetEnable(v bool) *RetryPolicySettingForListRoutesOutput {
	s.Enable = &v
	return s
}

// SetHttpCodes sets the HttpCodes field's value.
func (s *RetryPolicySettingForListRoutesOutput) SetHttpCodes(v []*string) *RetryPolicySettingForListRoutesOutput {
	s.HttpCodes = v
	return s
}

// SetPerTryTimeout sets the PerTryTimeout field's value.
func (s *RetryPolicySettingForListRoutesOutput) SetPerTryTimeout(v int64) *RetryPolicySettingForListRoutesOutput {
	s.PerTryTimeout = &v
	return s
}

// SetRetryOn sets the RetryOn field's value.
func (s *RetryPolicySettingForListRoutesOutput) SetRetryOn(v []*string) *RetryPolicySettingForListRoutesOutput {
	s.RetryOn = v
	return s
}

type TagForListRoutesOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Key *string `type:"string" json:",omitempty"`

	Value *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s TagForListRoutesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s TagForListRoutesOutput) GoString() string {
	return s.String()
}

// SetKey sets the Key field's value.
func (s *TagForListRoutesOutput) SetKey(v string) *TagForListRoutesOutput {
	s.Key = &v
	return s
}

// SetValue sets the Value field's value.
func (s *TagForListRoutesOutput) SetValue(v string) *TagForListRoutesOutput {
	s.Value = &v
	return s
}

type TimeoutSettingForListRoutesOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Enable *bool `type:"boolean" json:",omitempty"`

	Timeout *int64 `type:"int64" json:",omitempty"`
}

// String returns the string representation
func (s TimeoutSettingForListRoutesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s TimeoutSettingForListRoutesOutput) GoString() string {
	return s.String()
}

// SetEnable sets the Enable field's value.
func (s *TimeoutSettingForListRoutesOutput) SetEnable(v bool) *TimeoutSettingForListRoutesOutput {
	s.Enable = &v
	return s
}

// SetTimeout sets the Timeout field's value.
func (s *TimeoutSettingForListRoutesOutput) SetTimeout(v int64) *TimeoutSettingForListRoutesOutput {
	s.Timeout = &v
	return s
}

type URLRewriteSettingForListRoutesOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Enable *bool `type:"boolean" json:",omitempty"`

	URLRewrite *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s URLRewriteSettingForListRoutesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s URLRewriteSettingForListRoutesOutput) GoString() string {
	return s.String()
}

// SetEnable sets the Enable field's value.
func (s *URLRewriteSettingForListRoutesOutput) SetEnable(v bool) *URLRewriteSettingForListRoutesOutput {
	s.Enable = &v
	return s
}

// SetURLRewrite sets the URLRewrite field's value.
func (s *URLRewriteSettingForListRoutesOutput) SetURLRewrite(v string) *URLRewriteSettingForListRoutesOutput {
	s.URLRewrite = &v
	return s
}

type UpstreamForListRoutesOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Type *string `type:"string" json:",omitempty"`

	UpstreamId *string `type:"string" json:",omitempty"`

	Version *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s UpstreamForListRoutesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s UpstreamForListRoutesOutput) GoString() string {
	return s.String()
}

// SetType sets the Type field's value.
func (s *UpstreamForListRoutesOutput) SetType(v string) *UpstreamForListRoutesOutput {
	s.Type = &v
	return s
}

// SetUpstreamId sets the UpstreamId field's value.
func (s *UpstreamForListRoutesOutput) SetUpstreamId(v string) *UpstreamForListRoutesOutput {
	s.UpstreamId = &v
	return s
}

// SetVersion sets the Version field's value.
func (s *UpstreamForListRoutesOutput) SetVersion(v string) *UpstreamForListRoutesOutput {
	s.Version = &v
	return s
}

type UpstreamListForListRoutesOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	AIProviderSettings *AIProviderSettingsForListRoutesOutput `type:"structure" json:",omitempty"`

	UpstreamId *string `type:"string" json:",omitempty"`

	Version *string `type:"string" json:",omitempty"`

	Weight *int32 `type:"int32" json:",omitempty"`
}

// String returns the string representation
func (s UpstreamListForListRoutesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s UpstreamListForListRoutesOutput) GoString() string {
	return s.String()
}

// SetAIProviderSettings sets the AIProviderSettings field's value.
func (s *UpstreamListForListRoutesOutput) SetAIProviderSettings(v *AIProviderSettingsForListRoutesOutput) *UpstreamListForListRoutesOutput {
	s.AIProviderSettings = v
	return s
}

// SetUpstreamId sets the UpstreamId field's value.
func (s *UpstreamListForListRoutesOutput) SetUpstreamId(v string) *UpstreamListForListRoutesOutput {
	s.UpstreamId = &v
	return s
}

// SetVersion sets the Version field's value.
func (s *UpstreamListForListRoutesOutput) SetVersion(v string) *UpstreamListForListRoutesOutput {
	s.Version = &v
	return s
}

// SetWeight sets the Weight field's value.
func (s *UpstreamListForListRoutesOutput) SetWeight(v int32) *UpstreamListForListRoutesOutput {
	s.Weight = &v
	return s
}

type ValueForListRoutesOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	MatchContent *string `type:"string" json:",omitempty"`

	MatchType *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s ValueForListRoutesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ValueForListRoutesOutput) GoString() string {
	return s.String()
}

// SetMatchContent sets the MatchContent field's value.
func (s *ValueForListRoutesOutput) SetMatchContent(v string) *ValueForListRoutesOutput {
	s.MatchContent = &v
	return s
}

// SetMatchType sets the MatchType field's value.
func (s *ValueForListRoutesOutput) SetMatchType(v string) *ValueForListRoutesOutput {
	s.MatchType = &v
	return s
}

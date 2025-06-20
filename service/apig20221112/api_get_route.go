// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package apig20221112

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opGetRouteCommon = "GetRoute"

// GetRouteCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the GetRouteCommon operation. The "output" return
// value will be populated with the GetRouteCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned GetRouteCommon Request to send the API call to the service.
// the "output" return value is not valid until after GetRouteCommon Send returns without error.
//
// See GetRouteCommon for more information on using the GetRouteCommon
// API call, and error handling.
//
//    // Example sending a request using the GetRouteCommonRequest method.
//    req, resp := client.GetRouteCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *APIG20221112) GetRouteCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opGetRouteCommon,
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

// GetRouteCommon API operation for APIG20221112.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for APIG20221112's
// API operation GetRouteCommon for usage and error information.
func (c *APIG20221112) GetRouteCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.GetRouteCommonRequest(input)
	return out, req.Send()
}

// GetRouteCommonWithContext is the same as GetRouteCommon with the addition of
// the ability to pass a context and additional request options.
//
// See GetRouteCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *APIG20221112) GetRouteCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.GetRouteCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opGetRoute = "GetRoute"

// GetRouteRequest generates a "volcengine/request.Request" representing the
// client's request for the GetRoute operation. The "output" return
// value will be populated with the GetRouteCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned GetRouteCommon Request to send the API call to the service.
// the "output" return value is not valid until after GetRouteCommon Send returns without error.
//
// See GetRoute for more information on using the GetRoute
// API call, and error handling.
//
//    // Example sending a request using the GetRouteRequest method.
//    req, resp := client.GetRouteRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *APIG20221112) GetRouteRequest(input *GetRouteInput) (req *request.Request, output *GetRouteOutput) {
	op := &request.Operation{
		Name:       opGetRoute,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetRouteInput{}
	}

	output = &GetRouteOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// GetRoute API operation for APIG20221112.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for APIG20221112's
// API operation GetRoute for usage and error information.
func (c *APIG20221112) GetRoute(input *GetRouteInput) (*GetRouteOutput, error) {
	req, out := c.GetRouteRequest(input)
	return out, req.Send()
}

// GetRouteWithContext is the same as GetRoute with the addition of
// the ability to pass a context and additional request options.
//
// See GetRoute for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *APIG20221112) GetRouteWithContext(ctx volcengine.Context, input *GetRouteInput, opts ...request.Option) (*GetRouteOutput, error) {
	req, out := c.GetRouteRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type AIProviderSettingsForGetRouteOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Model *string `type:"string" json:",omitempty"`

	TargetPath *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s AIProviderSettingsForGetRouteOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s AIProviderSettingsForGetRouteOutput) GoString() string {
	return s.String()
}

// SetModel sets the Model field's value.
func (s *AIProviderSettingsForGetRouteOutput) SetModel(v string) *AIProviderSettingsForGetRouteOutput {
	s.Model = &v
	return s
}

// SetTargetPath sets the TargetPath field's value.
func (s *AIProviderSettingsForGetRouteOutput) SetTargetPath(v string) *AIProviderSettingsForGetRouteOutput {
	s.TargetPath = &v
	return s
}

type AdvancedSettingForGetRouteOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	CorsPolicySetting *CorsPolicySettingForGetRouteOutput `type:"structure" json:",omitempty"`

	HeaderOperations []*HeaderOperationForGetRouteOutput `type:"list" json:",omitempty"`

	MirrorPolicies []*MirrorPolicyForGetRouteOutput `type:"list" json:",omitempty"`

	RetryPolicySetting *RetryPolicySettingForGetRouteOutput `type:"structure" json:",omitempty"`

	TimeoutSetting *TimeoutSettingForGetRouteOutput `type:"structure" json:",omitempty"`

	URLRewriteSetting *URLRewriteSettingForGetRouteOutput `type:"structure" json:",omitempty"`
}

// String returns the string representation
func (s AdvancedSettingForGetRouteOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s AdvancedSettingForGetRouteOutput) GoString() string {
	return s.String()
}

// SetCorsPolicySetting sets the CorsPolicySetting field's value.
func (s *AdvancedSettingForGetRouteOutput) SetCorsPolicySetting(v *CorsPolicySettingForGetRouteOutput) *AdvancedSettingForGetRouteOutput {
	s.CorsPolicySetting = v
	return s
}

// SetHeaderOperations sets the HeaderOperations field's value.
func (s *AdvancedSettingForGetRouteOutput) SetHeaderOperations(v []*HeaderOperationForGetRouteOutput) *AdvancedSettingForGetRouteOutput {
	s.HeaderOperations = v
	return s
}

// SetMirrorPolicies sets the MirrorPolicies field's value.
func (s *AdvancedSettingForGetRouteOutput) SetMirrorPolicies(v []*MirrorPolicyForGetRouteOutput) *AdvancedSettingForGetRouteOutput {
	s.MirrorPolicies = v
	return s
}

// SetRetryPolicySetting sets the RetryPolicySetting field's value.
func (s *AdvancedSettingForGetRouteOutput) SetRetryPolicySetting(v *RetryPolicySettingForGetRouteOutput) *AdvancedSettingForGetRouteOutput {
	s.RetryPolicySetting = v
	return s
}

// SetTimeoutSetting sets the TimeoutSetting field's value.
func (s *AdvancedSettingForGetRouteOutput) SetTimeoutSetting(v *TimeoutSettingForGetRouteOutput) *AdvancedSettingForGetRouteOutput {
	s.TimeoutSetting = v
	return s
}

// SetURLRewriteSetting sets the URLRewriteSetting field's value.
func (s *AdvancedSettingForGetRouteOutput) SetURLRewriteSetting(v *URLRewriteSettingForGetRouteOutput) *AdvancedSettingForGetRouteOutput {
	s.URLRewriteSetting = v
	return s
}

type CorsPolicySettingForGetRouteOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Enable *bool `type:"boolean" json:",omitempty"`
}

// String returns the string representation
func (s CorsPolicySettingForGetRouteOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CorsPolicySettingForGetRouteOutput) GoString() string {
	return s.String()
}

// SetEnable sets the Enable field's value.
func (s *CorsPolicySettingForGetRouteOutput) SetEnable(v bool) *CorsPolicySettingForGetRouteOutput {
	s.Enable = &v
	return s
}

type CustomDomainForGetRouteOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Domain *string `type:"string" json:",omitempty"`

	Id *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s CustomDomainForGetRouteOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CustomDomainForGetRouteOutput) GoString() string {
	return s.String()
}

// SetDomain sets the Domain field's value.
func (s *CustomDomainForGetRouteOutput) SetDomain(v string) *CustomDomainForGetRouteOutput {
	s.Domain = &v
	return s
}

// SetId sets the Id field's value.
func (s *CustomDomainForGetRouteOutput) SetId(v string) *CustomDomainForGetRouteOutput {
	s.Id = &v
	return s
}

type DomainForGetRouteOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Domain *string `type:"string" json:",omitempty"`

	Type *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s DomainForGetRouteOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DomainForGetRouteOutput) GoString() string {
	return s.String()
}

// SetDomain sets the Domain field's value.
func (s *DomainForGetRouteOutput) SetDomain(v string) *DomainForGetRouteOutput {
	s.Domain = &v
	return s
}

// SetType sets the Type field's value.
func (s *DomainForGetRouteOutput) SetType(v string) *DomainForGetRouteOutput {
	s.Type = &v
	return s
}

type GetRouteInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	// Id is a required field
	Id *string `type:"string" json:",omitempty" required:"true"`
}

// String returns the string representation
func (s GetRouteInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s GetRouteInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetRouteInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "GetRouteInput"}
	if s.Id == nil {
		invalidParams.Add(request.NewErrParamRequired("Id"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetId sets the Id field's value.
func (s *GetRouteInput) SetId(v string) *GetRouteInput {
	s.Id = &v
	return s
}

type GetRouteOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	Route *RouteForGetRouteOutput `type:"structure" json:",omitempty"`
}

// String returns the string representation
func (s GetRouteOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s GetRouteOutput) GoString() string {
	return s.String()
}

// SetRoute sets the Route field's value.
func (s *GetRouteOutput) SetRoute(v *RouteForGetRouteOutput) *GetRouteOutput {
	s.Route = v
	return s
}

type HeaderForGetRouteOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Key *string `type:"string" json:",omitempty"`

	Value *ValueForGetRouteOutput `type:"structure" json:",omitempty"`
}

// String returns the string representation
func (s HeaderForGetRouteOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s HeaderForGetRouteOutput) GoString() string {
	return s.String()
}

// SetKey sets the Key field's value.
func (s *HeaderForGetRouteOutput) SetKey(v string) *HeaderForGetRouteOutput {
	s.Key = &v
	return s
}

// SetValue sets the Value field's value.
func (s *HeaderForGetRouteOutput) SetValue(v *ValueForGetRouteOutput) *HeaderForGetRouteOutput {
	s.Value = v
	return s
}

type HeaderOperationForGetRouteOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	DirectionType *string `type:"string" json:",omitempty"`

	Key *string `type:"string" json:",omitempty"`

	Operation *string `type:"string" json:",omitempty"`

	Value *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s HeaderOperationForGetRouteOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s HeaderOperationForGetRouteOutput) GoString() string {
	return s.String()
}

// SetDirectionType sets the DirectionType field's value.
func (s *HeaderOperationForGetRouteOutput) SetDirectionType(v string) *HeaderOperationForGetRouteOutput {
	s.DirectionType = &v
	return s
}

// SetKey sets the Key field's value.
func (s *HeaderOperationForGetRouteOutput) SetKey(v string) *HeaderOperationForGetRouteOutput {
	s.Key = &v
	return s
}

// SetOperation sets the Operation field's value.
func (s *HeaderOperationForGetRouteOutput) SetOperation(v string) *HeaderOperationForGetRouteOutput {
	s.Operation = &v
	return s
}

// SetValue sets the Value field's value.
func (s *HeaderOperationForGetRouteOutput) SetValue(v string) *HeaderOperationForGetRouteOutput {
	s.Value = &v
	return s
}

type MatchRuleForGetRouteOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Header []*HeaderForGetRouteOutput `type:"list" json:",omitempty"`

	Method []*string `type:"list" json:",omitempty"`

	Path *PathForGetRouteOutput `type:"structure" json:",omitempty"`

	QueryString []*QueryStringForGetRouteOutput `type:"list" json:",omitempty"`
}

// String returns the string representation
func (s MatchRuleForGetRouteOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s MatchRuleForGetRouteOutput) GoString() string {
	return s.String()
}

// SetHeader sets the Header field's value.
func (s *MatchRuleForGetRouteOutput) SetHeader(v []*HeaderForGetRouteOutput) *MatchRuleForGetRouteOutput {
	s.Header = v
	return s
}

// SetMethod sets the Method field's value.
func (s *MatchRuleForGetRouteOutput) SetMethod(v []*string) *MatchRuleForGetRouteOutput {
	s.Method = v
	return s
}

// SetPath sets the Path field's value.
func (s *MatchRuleForGetRouteOutput) SetPath(v *PathForGetRouteOutput) *MatchRuleForGetRouteOutput {
	s.Path = v
	return s
}

// SetQueryString sets the QueryString field's value.
func (s *MatchRuleForGetRouteOutput) SetQueryString(v []*QueryStringForGetRouteOutput) *MatchRuleForGetRouteOutput {
	s.QueryString = v
	return s
}

type MirrorPolicyForGetRouteOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Percent *PercentForGetRouteOutput `type:"structure" json:",omitempty"`

	Upstream *UpstreamForGetRouteOutput `type:"structure" json:",omitempty"`
}

// String returns the string representation
func (s MirrorPolicyForGetRouteOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s MirrorPolicyForGetRouteOutput) GoString() string {
	return s.String()
}

// SetPercent sets the Percent field's value.
func (s *MirrorPolicyForGetRouteOutput) SetPercent(v *PercentForGetRouteOutput) *MirrorPolicyForGetRouteOutput {
	s.Percent = v
	return s
}

// SetUpstream sets the Upstream field's value.
func (s *MirrorPolicyForGetRouteOutput) SetUpstream(v *UpstreamForGetRouteOutput) *MirrorPolicyForGetRouteOutput {
	s.Upstream = v
	return s
}

type PathForGetRouteOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	MatchContent *string `type:"string" json:",omitempty"`

	MatchType *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s PathForGetRouteOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s PathForGetRouteOutput) GoString() string {
	return s.String()
}

// SetMatchContent sets the MatchContent field's value.
func (s *PathForGetRouteOutput) SetMatchContent(v string) *PathForGetRouteOutput {
	s.MatchContent = &v
	return s
}

// SetMatchType sets the MatchType field's value.
func (s *PathForGetRouteOutput) SetMatchType(v string) *PathForGetRouteOutput {
	s.MatchType = &v
	return s
}

type PercentForGetRouteOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Value *float64 `type:"double" json:",omitempty"`
}

// String returns the string representation
func (s PercentForGetRouteOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s PercentForGetRouteOutput) GoString() string {
	return s.String()
}

// SetValue sets the Value field's value.
func (s *PercentForGetRouteOutput) SetValue(v float64) *PercentForGetRouteOutput {
	s.Value = &v
	return s
}

type QueryStringForGetRouteOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Key *string `type:"string" json:",omitempty"`

	Value *PathForGetRouteOutput `type:"structure" json:",omitempty"`
}

// String returns the string representation
func (s QueryStringForGetRouteOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s QueryStringForGetRouteOutput) GoString() string {
	return s.String()
}

// SetKey sets the Key field's value.
func (s *QueryStringForGetRouteOutput) SetKey(v string) *QueryStringForGetRouteOutput {
	s.Key = &v
	return s
}

// SetValue sets the Value field's value.
func (s *QueryStringForGetRouteOutput) SetValue(v *PathForGetRouteOutput) *QueryStringForGetRouteOutput {
	s.Value = v
	return s
}

type RetryPolicySettingForGetRouteOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Attempts *int64 `type:"int64" json:",omitempty"`

	Enable *bool `type:"boolean" json:",omitempty"`

	HttpCodes []*string `type:"list" json:",omitempty"`

	PerTryTimeout *int64 `type:"int64" json:",omitempty"`

	RetryOn []*string `type:"list" json:",omitempty"`
}

// String returns the string representation
func (s RetryPolicySettingForGetRouteOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s RetryPolicySettingForGetRouteOutput) GoString() string {
	return s.String()
}

// SetAttempts sets the Attempts field's value.
func (s *RetryPolicySettingForGetRouteOutput) SetAttempts(v int64) *RetryPolicySettingForGetRouteOutput {
	s.Attempts = &v
	return s
}

// SetEnable sets the Enable field's value.
func (s *RetryPolicySettingForGetRouteOutput) SetEnable(v bool) *RetryPolicySettingForGetRouteOutput {
	s.Enable = &v
	return s
}

// SetHttpCodes sets the HttpCodes field's value.
func (s *RetryPolicySettingForGetRouteOutput) SetHttpCodes(v []*string) *RetryPolicySettingForGetRouteOutput {
	s.HttpCodes = v
	return s
}

// SetPerTryTimeout sets the PerTryTimeout field's value.
func (s *RetryPolicySettingForGetRouteOutput) SetPerTryTimeout(v int64) *RetryPolicySettingForGetRouteOutput {
	s.PerTryTimeout = &v
	return s
}

// SetRetryOn sets the RetryOn field's value.
func (s *RetryPolicySettingForGetRouteOutput) SetRetryOn(v []*string) *RetryPolicySettingForGetRouteOutput {
	s.RetryOn = v
	return s
}

type RouteForGetRouteOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	AdvancedSetting *AdvancedSettingForGetRouteOutput `type:"structure" json:",omitempty"`

	CreateTime *string `type:"string" json:",omitempty"`

	CustomDomains []*CustomDomainForGetRouteOutput `type:"list" json:",omitempty"`

	Domains []*DomainForGetRouteOutput `type:"list" json:",omitempty"`

	Enable *bool `type:"boolean" json:",omitempty"`

	Id *string `type:"string" json:",omitempty"`

	MatchRule *MatchRuleForGetRouteOutput `type:"structure" json:",omitempty"`

	Name *string `type:"string" json:",omitempty"`

	Priority *int64 `type:"int64" json:",omitempty"`

	Reason *string `type:"string" json:",omitempty"`

	ResourceType *string `type:"string" json:",omitempty"`

	ServiceId *string `type:"string" json:",omitempty"`

	ServiceName *string `type:"string" json:",omitempty"`

	Status *string `type:"string" json:",omitempty"`

	Tags []*TagForGetRouteOutput `type:"list" json:",omitempty"`

	UpdateTime *string `type:"string" json:",omitempty"`

	UpstreamList []*UpstreamListForGetRouteOutput `type:"list" json:",omitempty"`
}

// String returns the string representation
func (s RouteForGetRouteOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s RouteForGetRouteOutput) GoString() string {
	return s.String()
}

// SetAdvancedSetting sets the AdvancedSetting field's value.
func (s *RouteForGetRouteOutput) SetAdvancedSetting(v *AdvancedSettingForGetRouteOutput) *RouteForGetRouteOutput {
	s.AdvancedSetting = v
	return s
}

// SetCreateTime sets the CreateTime field's value.
func (s *RouteForGetRouteOutput) SetCreateTime(v string) *RouteForGetRouteOutput {
	s.CreateTime = &v
	return s
}

// SetCustomDomains sets the CustomDomains field's value.
func (s *RouteForGetRouteOutput) SetCustomDomains(v []*CustomDomainForGetRouteOutput) *RouteForGetRouteOutput {
	s.CustomDomains = v
	return s
}

// SetDomains sets the Domains field's value.
func (s *RouteForGetRouteOutput) SetDomains(v []*DomainForGetRouteOutput) *RouteForGetRouteOutput {
	s.Domains = v
	return s
}

// SetEnable sets the Enable field's value.
func (s *RouteForGetRouteOutput) SetEnable(v bool) *RouteForGetRouteOutput {
	s.Enable = &v
	return s
}

// SetId sets the Id field's value.
func (s *RouteForGetRouteOutput) SetId(v string) *RouteForGetRouteOutput {
	s.Id = &v
	return s
}

// SetMatchRule sets the MatchRule field's value.
func (s *RouteForGetRouteOutput) SetMatchRule(v *MatchRuleForGetRouteOutput) *RouteForGetRouteOutput {
	s.MatchRule = v
	return s
}

// SetName sets the Name field's value.
func (s *RouteForGetRouteOutput) SetName(v string) *RouteForGetRouteOutput {
	s.Name = &v
	return s
}

// SetPriority sets the Priority field's value.
func (s *RouteForGetRouteOutput) SetPriority(v int64) *RouteForGetRouteOutput {
	s.Priority = &v
	return s
}

// SetReason sets the Reason field's value.
func (s *RouteForGetRouteOutput) SetReason(v string) *RouteForGetRouteOutput {
	s.Reason = &v
	return s
}

// SetResourceType sets the ResourceType field's value.
func (s *RouteForGetRouteOutput) SetResourceType(v string) *RouteForGetRouteOutput {
	s.ResourceType = &v
	return s
}

// SetServiceId sets the ServiceId field's value.
func (s *RouteForGetRouteOutput) SetServiceId(v string) *RouteForGetRouteOutput {
	s.ServiceId = &v
	return s
}

// SetServiceName sets the ServiceName field's value.
func (s *RouteForGetRouteOutput) SetServiceName(v string) *RouteForGetRouteOutput {
	s.ServiceName = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *RouteForGetRouteOutput) SetStatus(v string) *RouteForGetRouteOutput {
	s.Status = &v
	return s
}

// SetTags sets the Tags field's value.
func (s *RouteForGetRouteOutput) SetTags(v []*TagForGetRouteOutput) *RouteForGetRouteOutput {
	s.Tags = v
	return s
}

// SetUpdateTime sets the UpdateTime field's value.
func (s *RouteForGetRouteOutput) SetUpdateTime(v string) *RouteForGetRouteOutput {
	s.UpdateTime = &v
	return s
}

// SetUpstreamList sets the UpstreamList field's value.
func (s *RouteForGetRouteOutput) SetUpstreamList(v []*UpstreamListForGetRouteOutput) *RouteForGetRouteOutput {
	s.UpstreamList = v
	return s
}

type TagForGetRouteOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Key *string `type:"string" json:",omitempty"`

	Value *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s TagForGetRouteOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s TagForGetRouteOutput) GoString() string {
	return s.String()
}

// SetKey sets the Key field's value.
func (s *TagForGetRouteOutput) SetKey(v string) *TagForGetRouteOutput {
	s.Key = &v
	return s
}

// SetValue sets the Value field's value.
func (s *TagForGetRouteOutput) SetValue(v string) *TagForGetRouteOutput {
	s.Value = &v
	return s
}

type TimeoutSettingForGetRouteOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Enable *bool `type:"boolean" json:",omitempty"`

	Timeout *int64 `type:"int64" json:",omitempty"`
}

// String returns the string representation
func (s TimeoutSettingForGetRouteOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s TimeoutSettingForGetRouteOutput) GoString() string {
	return s.String()
}

// SetEnable sets the Enable field's value.
func (s *TimeoutSettingForGetRouteOutput) SetEnable(v bool) *TimeoutSettingForGetRouteOutput {
	s.Enable = &v
	return s
}

// SetTimeout sets the Timeout field's value.
func (s *TimeoutSettingForGetRouteOutput) SetTimeout(v int64) *TimeoutSettingForGetRouteOutput {
	s.Timeout = &v
	return s
}

type URLRewriteSettingForGetRouteOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Enable *bool `type:"boolean" json:",omitempty"`

	URLRewrite *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s URLRewriteSettingForGetRouteOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s URLRewriteSettingForGetRouteOutput) GoString() string {
	return s.String()
}

// SetEnable sets the Enable field's value.
func (s *URLRewriteSettingForGetRouteOutput) SetEnable(v bool) *URLRewriteSettingForGetRouteOutput {
	s.Enable = &v
	return s
}

// SetURLRewrite sets the URLRewrite field's value.
func (s *URLRewriteSettingForGetRouteOutput) SetURLRewrite(v string) *URLRewriteSettingForGetRouteOutput {
	s.URLRewrite = &v
	return s
}

type UpstreamForGetRouteOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Type *string `type:"string" json:",omitempty"`

	UpstreamId *string `type:"string" json:",omitempty"`

	Version *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s UpstreamForGetRouteOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s UpstreamForGetRouteOutput) GoString() string {
	return s.String()
}

// SetType sets the Type field's value.
func (s *UpstreamForGetRouteOutput) SetType(v string) *UpstreamForGetRouteOutput {
	s.Type = &v
	return s
}

// SetUpstreamId sets the UpstreamId field's value.
func (s *UpstreamForGetRouteOutput) SetUpstreamId(v string) *UpstreamForGetRouteOutput {
	s.UpstreamId = &v
	return s
}

// SetVersion sets the Version field's value.
func (s *UpstreamForGetRouteOutput) SetVersion(v string) *UpstreamForGetRouteOutput {
	s.Version = &v
	return s
}

type UpstreamListForGetRouteOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	AIProviderSettings *AIProviderSettingsForGetRouteOutput `type:"structure" json:",omitempty"`

	UpstreamId *string `type:"string" json:",omitempty"`

	Version *string `type:"string" json:",omitempty"`

	Weight *int32 `type:"int32" json:",omitempty"`
}

// String returns the string representation
func (s UpstreamListForGetRouteOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s UpstreamListForGetRouteOutput) GoString() string {
	return s.String()
}

// SetAIProviderSettings sets the AIProviderSettings field's value.
func (s *UpstreamListForGetRouteOutput) SetAIProviderSettings(v *AIProviderSettingsForGetRouteOutput) *UpstreamListForGetRouteOutput {
	s.AIProviderSettings = v
	return s
}

// SetUpstreamId sets the UpstreamId field's value.
func (s *UpstreamListForGetRouteOutput) SetUpstreamId(v string) *UpstreamListForGetRouteOutput {
	s.UpstreamId = &v
	return s
}

// SetVersion sets the Version field's value.
func (s *UpstreamListForGetRouteOutput) SetVersion(v string) *UpstreamListForGetRouteOutput {
	s.Version = &v
	return s
}

// SetWeight sets the Weight field's value.
func (s *UpstreamListForGetRouteOutput) SetWeight(v int32) *UpstreamListForGetRouteOutput {
	s.Weight = &v
	return s
}

type ValueForGetRouteOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	MatchContent *string `type:"string" json:",omitempty"`

	MatchType *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s ValueForGetRouteOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ValueForGetRouteOutput) GoString() string {
	return s.String()
}

// SetMatchContent sets the MatchContent field's value.
func (s *ValueForGetRouteOutput) SetMatchContent(v string) *ValueForGetRouteOutput {
	s.MatchContent = &v
	return s
}

// SetMatchType sets the MatchType field's value.
func (s *ValueForGetRouteOutput) SetMatchType(v string) *ValueForGetRouteOutput {
	s.MatchType = &v
	return s
}

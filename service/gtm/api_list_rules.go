// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package gtm

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
func (c *GTM) ListRulesCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opListRulesCommon,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &map[string]interface{}{}
	}

	output = &map[string]interface{}{}
	req = c.newRequest(op, input, output)

	return
}

// ListRulesCommon API operation for GTM.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for GTM's
// API operation ListRulesCommon for usage and error information.
func (c *GTM) ListRulesCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
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
func (c *GTM) ListRulesCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
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
func (c *GTM) ListRulesRequest(input *ListRulesInput) (req *request.Request, output *ListRulesOutput) {
	op := &request.Operation{
		Name:       opListRules,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListRulesInput{}
	}

	output = &ListRulesOutput{}
	req = c.newRequest(op, input, output)

	return
}

// ListRules API operation for GTM.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for GTM's
// API operation ListRules for usage and error information.
func (c *GTM) ListRules(input *ListRulesInput) (*ListRulesOutput, error) {
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
func (c *GTM) ListRulesWithContext(ctx volcengine.Context, input *ListRulesInput, opts ...request.Option) (*ListRulesOutput, error) {
	req, out := c.ListRulesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type AddressForListRulesOutput struct {
	_ struct{} `type:"structure"`

	Active *bool `type:"boolean"`

	Capacity *int64 `type:"int64"`

	Geo *string `type:"string"`

	Latency *int32 `type:"int32"`

	Mode *string `type:"string"`

	RectifiedGeos []*string `type:"list"`

	Value *string `type:"string"`

	Weight *int32 `type:"int32"`
}

// String returns the string representation
func (s AddressForListRulesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s AddressForListRulesOutput) GoString() string {
	return s.String()
}

// SetActive sets the Active field's value.
func (s *AddressForListRulesOutput) SetActive(v bool) *AddressForListRulesOutput {
	s.Active = &v
	return s
}

// SetCapacity sets the Capacity field's value.
func (s *AddressForListRulesOutput) SetCapacity(v int64) *AddressForListRulesOutput {
	s.Capacity = &v
	return s
}

// SetGeo sets the Geo field's value.
func (s *AddressForListRulesOutput) SetGeo(v string) *AddressForListRulesOutput {
	s.Geo = &v
	return s
}

// SetLatency sets the Latency field's value.
func (s *AddressForListRulesOutput) SetLatency(v int32) *AddressForListRulesOutput {
	s.Latency = &v
	return s
}

// SetMode sets the Mode field's value.
func (s *AddressForListRulesOutput) SetMode(v string) *AddressForListRulesOutput {
	s.Mode = &v
	return s
}

// SetRectifiedGeos sets the RectifiedGeos field's value.
func (s *AddressForListRulesOutput) SetRectifiedGeos(v []*string) *AddressForListRulesOutput {
	s.RectifiedGeos = v
	return s
}

// SetValue sets the Value field's value.
func (s *AddressForListRulesOutput) SetValue(v string) *AddressForListRulesOutput {
	s.Value = &v
	return s
}

// SetWeight sets the Weight field's value.
func (s *AddressForListRulesOutput) SetWeight(v int32) *AddressForListRulesOutput {
	s.Weight = &v
	return s
}

type DataForListRulesOutput struct {
	_ struct{} `type:"structure"`

	CreatedTime *string `type:"string"`

	Disable *bool `type:"boolean"`

	EffectivePoolSetIndex *int32 `type:"int32"`

	Line *string `type:"string"`

	Name *string `type:"string"`

	PoolSetMode *string `type:"string" enum:"EnumOfPoolSetModeForListRulesOutput"`

	PoolSets []*PoolSetForListRulesOutput `type:"list"`

	Probe *ProbeForListRulesOutput `type:"structure"`

	ProbeMode *string `type:"string" enum:"EnumOfProbeModeForListRulesOutput"`

	Remark *string `type:"string"`

	RuleId *string `type:"string"`

	UpdatedTime *string `type:"string"`

	Weight *int32 `type:"int32"`
}

// String returns the string representation
func (s DataForListRulesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DataForListRulesOutput) GoString() string {
	return s.String()
}

// SetCreatedTime sets the CreatedTime field's value.
func (s *DataForListRulesOutput) SetCreatedTime(v string) *DataForListRulesOutput {
	s.CreatedTime = &v
	return s
}

// SetDisable sets the Disable field's value.
func (s *DataForListRulesOutput) SetDisable(v bool) *DataForListRulesOutput {
	s.Disable = &v
	return s
}

// SetEffectivePoolSetIndex sets the EffectivePoolSetIndex field's value.
func (s *DataForListRulesOutput) SetEffectivePoolSetIndex(v int32) *DataForListRulesOutput {
	s.EffectivePoolSetIndex = &v
	return s
}

// SetLine sets the Line field's value.
func (s *DataForListRulesOutput) SetLine(v string) *DataForListRulesOutput {
	s.Line = &v
	return s
}

// SetName sets the Name field's value.
func (s *DataForListRulesOutput) SetName(v string) *DataForListRulesOutput {
	s.Name = &v
	return s
}

// SetPoolSetMode sets the PoolSetMode field's value.
func (s *DataForListRulesOutput) SetPoolSetMode(v string) *DataForListRulesOutput {
	s.PoolSetMode = &v
	return s
}

// SetPoolSets sets the PoolSets field's value.
func (s *DataForListRulesOutput) SetPoolSets(v []*PoolSetForListRulesOutput) *DataForListRulesOutput {
	s.PoolSets = v
	return s
}

// SetProbe sets the Probe field's value.
func (s *DataForListRulesOutput) SetProbe(v *ProbeForListRulesOutput) *DataForListRulesOutput {
	s.Probe = v
	return s
}

// SetProbeMode sets the ProbeMode field's value.
func (s *DataForListRulesOutput) SetProbeMode(v string) *DataForListRulesOutput {
	s.ProbeMode = &v
	return s
}

// SetRemark sets the Remark field's value.
func (s *DataForListRulesOutput) SetRemark(v string) *DataForListRulesOutput {
	s.Remark = &v
	return s
}

// SetRuleId sets the RuleId field's value.
func (s *DataForListRulesOutput) SetRuleId(v string) *DataForListRulesOutput {
	s.RuleId = &v
	return s
}

// SetUpdatedTime sets the UpdatedTime field's value.
func (s *DataForListRulesOutput) SetUpdatedTime(v string) *DataForListRulesOutput {
	s.UpdatedTime = &v
	return s
}

// SetWeight sets the Weight field's value.
func (s *DataForListRulesOutput) SetWeight(v int32) *DataForListRulesOutput {
	s.Weight = &v
	return s
}

type HttpHeaderForListRulesOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s HttpHeaderForListRulesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s HttpHeaderForListRulesOutput) GoString() string {
	return s.String()
}

type HttpUsabilityCodeForListRulesOutput struct {
	_ struct{} `type:"structure"`

	Codes []*int32 `type:"list"`

	Operator *string `type:"string"`
}

// String returns the string representation
func (s HttpUsabilityCodeForListRulesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s HttpUsabilityCodeForListRulesOutput) GoString() string {
	return s.String()
}

// SetCodes sets the Codes field's value.
func (s *HttpUsabilityCodeForListRulesOutput) SetCodes(v []*int32) *HttpUsabilityCodeForListRulesOutput {
	s.Codes = v
	return s
}

// SetOperator sets the Operator field's value.
func (s *HttpUsabilityCodeForListRulesOutput) SetOperator(v string) *HttpUsabilityCodeForListRulesOutput {
	s.Operator = &v
	return s
}

type ListRulesInput struct {
	_ struct{} `type:"structure"`

	Address *string `type:"string"`

	// GtmId is a required field
	GtmId *string `type:"string" required:"true"`

	Line *string `type:"string"`

	Name *string `type:"string"`

	PageNumber *int32 `type:"int32"`

	PageSize *int32 `type:"int32"`

	// PolicyType is a required field
	PolicyType *string `type:"string" required:"true"`

	SearchMode *string `type:"string"`

	SortBy *string `type:"string"`

	SortOrder *string `type:"string"`
}

// String returns the string representation
func (s ListRulesInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListRulesInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListRulesInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ListRulesInput"}
	if s.GtmId == nil {
		invalidParams.Add(request.NewErrParamRequired("GtmId"))
	}
	if s.PolicyType == nil {
		invalidParams.Add(request.NewErrParamRequired("PolicyType"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetAddress sets the Address field's value.
func (s *ListRulesInput) SetAddress(v string) *ListRulesInput {
	s.Address = &v
	return s
}

// SetGtmId sets the GtmId field's value.
func (s *ListRulesInput) SetGtmId(v string) *ListRulesInput {
	s.GtmId = &v
	return s
}

// SetLine sets the Line field's value.
func (s *ListRulesInput) SetLine(v string) *ListRulesInput {
	s.Line = &v
	return s
}

// SetName sets the Name field's value.
func (s *ListRulesInput) SetName(v string) *ListRulesInput {
	s.Name = &v
	return s
}

// SetPageNumber sets the PageNumber field's value.
func (s *ListRulesInput) SetPageNumber(v int32) *ListRulesInput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *ListRulesInput) SetPageSize(v int32) *ListRulesInput {
	s.PageSize = &v
	return s
}

// SetPolicyType sets the PolicyType field's value.
func (s *ListRulesInput) SetPolicyType(v string) *ListRulesInput {
	s.PolicyType = &v
	return s
}

// SetSearchMode sets the SearchMode field's value.
func (s *ListRulesInput) SetSearchMode(v string) *ListRulesInput {
	s.SearchMode = &v
	return s
}

// SetSortBy sets the SortBy field's value.
func (s *ListRulesInput) SetSortBy(v string) *ListRulesInput {
	s.SortBy = &v
	return s
}

// SetSortOrder sets the SortOrder field's value.
func (s *ListRulesInput) SetSortOrder(v string) *ListRulesInput {
	s.SortOrder = &v
	return s
}

type ListRulesOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	Data []*DataForListRulesOutput `type:"list"`

	PageNumber *int32 `type:"int32"`

	PageSize *int32 `type:"int32"`

	TotalCount *int32 `type:"int32"`
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
func (s *ListRulesOutput) SetPageNumber(v int32) *ListRulesOutput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *ListRulesOutput) SetPageSize(v int32) *ListRulesOutput {
	s.PageSize = &v
	return s
}

// SetTotalCount sets the TotalCount field's value.
func (s *ListRulesOutput) SetTotalCount(v int32) *ListRulesOutput {
	s.TotalCount = &v
	return s
}

type PoolForListRulesOutput struct {
	_ struct{} `type:"structure"`

	Address []*AddressForListRulesOutput `type:"list"`

	InactiveAddrCount *int32 `type:"int32"`

	Name *string `type:"string"`

	PoolId *string `type:"string"`

	Weight *int32 `type:"int32"`
}

// String returns the string representation
func (s PoolForListRulesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s PoolForListRulesOutput) GoString() string {
	return s.String()
}

// SetAddress sets the Address field's value.
func (s *PoolForListRulesOutput) SetAddress(v []*AddressForListRulesOutput) *PoolForListRulesOutput {
	s.Address = v
	return s
}

// SetInactiveAddrCount sets the InactiveAddrCount field's value.
func (s *PoolForListRulesOutput) SetInactiveAddrCount(v int32) *PoolForListRulesOutput {
	s.InactiveAddrCount = &v
	return s
}

// SetName sets the Name field's value.
func (s *PoolForListRulesOutput) SetName(v string) *PoolForListRulesOutput {
	s.Name = &v
	return s
}

// SetPoolId sets the PoolId field's value.
func (s *PoolForListRulesOutput) SetPoolId(v string) *PoolForListRulesOutput {
	s.PoolId = &v
	return s
}

// SetWeight sets the Weight field's value.
func (s *PoolForListRulesOutput) SetWeight(v int32) *PoolForListRulesOutput {
	s.Weight = &v
	return s
}

type PoolSetForListRulesOutput struct {
	_ struct{} `type:"structure"`

	Active *bool `type:"boolean"`

	ActiveAddrCount *int32 `type:"int32"`

	Name *string `type:"string"`

	PoolSetId *string `type:"string"`

	Pools []*PoolForListRulesOutput `type:"list"`
}

// String returns the string representation
func (s PoolSetForListRulesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s PoolSetForListRulesOutput) GoString() string {
	return s.String()
}

// SetActive sets the Active field's value.
func (s *PoolSetForListRulesOutput) SetActive(v bool) *PoolSetForListRulesOutput {
	s.Active = &v
	return s
}

// SetActiveAddrCount sets the ActiveAddrCount field's value.
func (s *PoolSetForListRulesOutput) SetActiveAddrCount(v int32) *PoolSetForListRulesOutput {
	s.ActiveAddrCount = &v
	return s
}

// SetName sets the Name field's value.
func (s *PoolSetForListRulesOutput) SetName(v string) *PoolSetForListRulesOutput {
	s.Name = &v
	return s
}

// SetPoolSetId sets the PoolSetId field's value.
func (s *PoolSetForListRulesOutput) SetPoolSetId(v string) *PoolSetForListRulesOutput {
	s.PoolSetId = &v
	return s
}

// SetPools sets the Pools field's value.
func (s *PoolSetForListRulesOutput) SetPools(v []*PoolForListRulesOutput) *PoolSetForListRulesOutput {
	s.Pools = v
	return s
}

type ProbeForListRulesOutput struct {
	_ struct{} `type:"structure"`

	AdvisedNodeCount *int32 `type:"int32"`

	Disable *bool `type:"boolean"`

	FailedCount *int32 `type:"int32"`

	Host *string `type:"string"`

	HttpHeader *HttpHeaderForListRulesOutput `type:"structure"`

	HttpMethod *string `type:"string"`

	HttpUsabilityCodes []*HttpUsabilityCodeForListRulesOutput `type:"list"`

	Interval *int32 `type:"int32"`

	IsManualNodes *bool `type:"boolean"`

	Nodes []*string `type:"list"`

	PingCount *int32 `type:"int32"`

	PingLossPercent *int32 `type:"int32"`

	Port *int32 `type:"int32"`

	Protocol *string `type:"string"`

	Timeout *int32 `type:"int32"`

	Url *string `type:"string"`
}

// String returns the string representation
func (s ProbeForListRulesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ProbeForListRulesOutput) GoString() string {
	return s.String()
}

// SetAdvisedNodeCount sets the AdvisedNodeCount field's value.
func (s *ProbeForListRulesOutput) SetAdvisedNodeCount(v int32) *ProbeForListRulesOutput {
	s.AdvisedNodeCount = &v
	return s
}

// SetDisable sets the Disable field's value.
func (s *ProbeForListRulesOutput) SetDisable(v bool) *ProbeForListRulesOutput {
	s.Disable = &v
	return s
}

// SetFailedCount sets the FailedCount field's value.
func (s *ProbeForListRulesOutput) SetFailedCount(v int32) *ProbeForListRulesOutput {
	s.FailedCount = &v
	return s
}

// SetHost sets the Host field's value.
func (s *ProbeForListRulesOutput) SetHost(v string) *ProbeForListRulesOutput {
	s.Host = &v
	return s
}

// SetHttpHeader sets the HttpHeader field's value.
func (s *ProbeForListRulesOutput) SetHttpHeader(v *HttpHeaderForListRulesOutput) *ProbeForListRulesOutput {
	s.HttpHeader = v
	return s
}

// SetHttpMethod sets the HttpMethod field's value.
func (s *ProbeForListRulesOutput) SetHttpMethod(v string) *ProbeForListRulesOutput {
	s.HttpMethod = &v
	return s
}

// SetHttpUsabilityCodes sets the HttpUsabilityCodes field's value.
func (s *ProbeForListRulesOutput) SetHttpUsabilityCodes(v []*HttpUsabilityCodeForListRulesOutput) *ProbeForListRulesOutput {
	s.HttpUsabilityCodes = v
	return s
}

// SetInterval sets the Interval field's value.
func (s *ProbeForListRulesOutput) SetInterval(v int32) *ProbeForListRulesOutput {
	s.Interval = &v
	return s
}

// SetIsManualNodes sets the IsManualNodes field's value.
func (s *ProbeForListRulesOutput) SetIsManualNodes(v bool) *ProbeForListRulesOutput {
	s.IsManualNodes = &v
	return s
}

// SetNodes sets the Nodes field's value.
func (s *ProbeForListRulesOutput) SetNodes(v []*string) *ProbeForListRulesOutput {
	s.Nodes = v
	return s
}

// SetPingCount sets the PingCount field's value.
func (s *ProbeForListRulesOutput) SetPingCount(v int32) *ProbeForListRulesOutput {
	s.PingCount = &v
	return s
}

// SetPingLossPercent sets the PingLossPercent field's value.
func (s *ProbeForListRulesOutput) SetPingLossPercent(v int32) *ProbeForListRulesOutput {
	s.PingLossPercent = &v
	return s
}

// SetPort sets the Port field's value.
func (s *ProbeForListRulesOutput) SetPort(v int32) *ProbeForListRulesOutput {
	s.Port = &v
	return s
}

// SetProtocol sets the Protocol field's value.
func (s *ProbeForListRulesOutput) SetProtocol(v string) *ProbeForListRulesOutput {
	s.Protocol = &v
	return s
}

// SetTimeout sets the Timeout field's value.
func (s *ProbeForListRulesOutput) SetTimeout(v int32) *ProbeForListRulesOutput {
	s.Timeout = &v
	return s
}

// SetUrl sets the Url field's value.
func (s *ProbeForListRulesOutput) SetUrl(v string) *ProbeForListRulesOutput {
	s.Url = &v
	return s
}

const (
	// EnumOfPoolSetModeForListRulesOutputAuto is a EnumOfPoolSetModeForListRulesOutput enum value
	EnumOfPoolSetModeForListRulesOutputAuto = "auto"

	// EnumOfPoolSetModeForListRulesOutputManual is a EnumOfPoolSetModeForListRulesOutput enum value
	EnumOfPoolSetModeForListRulesOutputManual = "manual"
)

const (
	// EnumOfProbeModeForListRulesOutputDefault is a EnumOfProbeModeForListRulesOutput enum value
	EnumOfProbeModeForListRulesOutputDefault = "default"

	// EnumOfProbeModeForListRulesOutputCustomizedBaseConfig is a EnumOfProbeModeForListRulesOutput enum value
	EnumOfProbeModeForListRulesOutputCustomizedBaseConfig = "customized_base_config"

	// EnumOfProbeModeForListRulesOutputCustomizedProbeNodes is a EnumOfProbeModeForListRulesOutput enum value
	EnumOfProbeModeForListRulesOutputCustomizedProbeNodes = "customized_probe_nodes"

	// EnumOfProbeModeForListRulesOutputCustomizedAll is a EnumOfProbeModeForListRulesOutput enum value
	EnumOfProbeModeForListRulesOutputCustomizedAll = "customized_all"
)

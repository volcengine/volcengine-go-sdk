// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package mcs

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opGetRiskStatCommon = "GetRiskStat"

// GetRiskStatCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the GetRiskStatCommon operation. The "output" return
// value will be populated with the GetRiskStatCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned GetRiskStatCommon Request to send the API call to the service.
// the "output" return value is not valid until after GetRiskStatCommon Send returns without error.
//
// See GetRiskStatCommon for more information on using the GetRiskStatCommon
// API call, and error handling.
//
//    // Example sending a request using the GetRiskStatCommonRequest method.
//    req, resp := client.GetRiskStatCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *MCS) GetRiskStatCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opGetRiskStatCommon,
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

// GetRiskStatCommon API operation for MCS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for MCS's
// API operation GetRiskStatCommon for usage and error information.
func (c *MCS) GetRiskStatCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.GetRiskStatCommonRequest(input)
	return out, req.Send()
}

// GetRiskStatCommonWithContext is the same as GetRiskStatCommon with the addition of
// the ability to pass a context and additional request options.
//
// See GetRiskStatCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *MCS) GetRiskStatCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.GetRiskStatCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opGetRiskStat = "GetRiskStat"

// GetRiskStatRequest generates a "volcengine/request.Request" representing the
// client's request for the GetRiskStat operation. The "output" return
// value will be populated with the GetRiskStatCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned GetRiskStatCommon Request to send the API call to the service.
// the "output" return value is not valid until after GetRiskStatCommon Send returns without error.
//
// See GetRiskStat for more information on using the GetRiskStat
// API call, and error handling.
//
//    // Example sending a request using the GetRiskStatRequest method.
//    req, resp := client.GetRiskStatRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *MCS) GetRiskStatRequest(input *GetRiskStatInput) (req *request.Request, output *GetRiskStatOutput) {
	op := &request.Operation{
		Name:       opGetRiskStat,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetRiskStatInput{}
	}

	output = &GetRiskStatOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// GetRiskStat API operation for MCS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for MCS's
// API operation GetRiskStat for usage and error information.
func (c *MCS) GetRiskStat(input *GetRiskStatInput) (*GetRiskStatOutput, error) {
	req, out := c.GetRiskStatRequest(input)
	return out, req.Send()
}

// GetRiskStatWithContext is the same as GetRiskStat with the addition of
// the ability to pass a context and additional request options.
//
// See GetRiskStat for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *MCS) GetRiskStatWithContext(ctx volcengine.Context, input *GetRiskStatInput, opts ...request.Option) (*GetRiskStatOutput, error) {
	req, out := c.GetRiskStatRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type CloudAccountForGetRiskStatInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	CloudAccountName *string `type:"string" json:",omitempty"`

	CloudAccountUID *string `type:"string" json:",omitempty"`

	CloudVendor *string `type:"string" json:",omitempty" enum:"EnumOfCloudVendorForGetRiskStatInput"`
}

// String returns the string representation
func (s CloudAccountForGetRiskStatInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CloudAccountForGetRiskStatInput) GoString() string {
	return s.String()
}

// SetCloudAccountName sets the CloudAccountName field's value.
func (s *CloudAccountForGetRiskStatInput) SetCloudAccountName(v string) *CloudAccountForGetRiskStatInput {
	s.CloudAccountName = &v
	return s
}

// SetCloudAccountUID sets the CloudAccountUID field's value.
func (s *CloudAccountForGetRiskStatInput) SetCloudAccountUID(v string) *CloudAccountForGetRiskStatInput {
	s.CloudAccountUID = &v
	return s
}

// SetCloudVendor sets the CloudVendor field's value.
func (s *CloudAccountForGetRiskStatInput) SetCloudVendor(v string) *CloudAccountForGetRiskStatInput {
	s.CloudVendor = &v
	return s
}

type GetRiskStatInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	CloudAccounts []*CloudAccountForGetRiskStatInput `type:"list" json:",omitempty"`

	CloudVendors []*string `type:"list" json:",omitempty"`

	DisplayMode *string `type:"string" json:",omitempty" enum:"EnumOfDisplayModeForGetRiskStatInput"`

	RiskType *string `type:"string" json:",omitempty" enum:"EnumOfRiskTypeForGetRiskStatInput"`

	TimeRange *TimeRangeForGetRiskStatInput `type:"structure" json:",omitempty"`
}

// String returns the string representation
func (s GetRiskStatInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s GetRiskStatInput) GoString() string {
	return s.String()
}

// SetCloudAccounts sets the CloudAccounts field's value.
func (s *GetRiskStatInput) SetCloudAccounts(v []*CloudAccountForGetRiskStatInput) *GetRiskStatInput {
	s.CloudAccounts = v
	return s
}

// SetCloudVendors sets the CloudVendors field's value.
func (s *GetRiskStatInput) SetCloudVendors(v []*string) *GetRiskStatInput {
	s.CloudVendors = v
	return s
}

// SetDisplayMode sets the DisplayMode field's value.
func (s *GetRiskStatInput) SetDisplayMode(v string) *GetRiskStatInput {
	s.DisplayMode = &v
	return s
}

// SetRiskType sets the RiskType field's value.
func (s *GetRiskStatInput) SetRiskType(v string) *GetRiskStatInput {
	s.RiskType = &v
	return s
}

// SetTimeRange sets the TimeRange field's value.
func (s *GetRiskStatInput) SetTimeRange(v *TimeRangeForGetRiskStatInput) *GetRiskStatInput {
	s.TimeRange = v
	return s
}

type GetRiskStatOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	NotFixed *NotFixedForGetRiskStatOutput `type:"structure" json:",omitempty"`

	StatByStatus *StatByStatusForGetRiskStatOutput `type:"structure" json:",omitempty"`

	Trend *TrendForGetRiskStatOutput `type:"structure" json:",omitempty"`
}

// String returns the string representation
func (s GetRiskStatOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s GetRiskStatOutput) GoString() string {
	return s.String()
}

// SetNotFixed sets the NotFixed field's value.
func (s *GetRiskStatOutput) SetNotFixed(v *NotFixedForGetRiskStatOutput) *GetRiskStatOutput {
	s.NotFixed = v
	return s
}

// SetStatByStatus sets the StatByStatus field's value.
func (s *GetRiskStatOutput) SetStatByStatus(v *StatByStatusForGetRiskStatOutput) *GetRiskStatOutput {
	s.StatByStatus = v
	return s
}

// SetTrend sets the Trend field's value.
func (s *GetRiskStatOutput) SetTrend(v *TrendForGetRiskStatOutput) *GetRiskStatOutput {
	s.Trend = v
	return s
}

type NotFixedForGetRiskStatOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Critical *int64 `type:"int64" json:",omitempty"`

	High *int64 `type:"int64" json:",omitempty"`

	Low *int64 `type:"int64" json:",omitempty"`

	Medium *int64 `type:"int64" json:",omitempty"`

	Security *int64 `type:"int64" json:",omitempty"`

	Total *int64 `type:"int64" json:",omitempty"`
}

// String returns the string representation
func (s NotFixedForGetRiskStatOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s NotFixedForGetRiskStatOutput) GoString() string {
	return s.String()
}

// SetCritical sets the Critical field's value.
func (s *NotFixedForGetRiskStatOutput) SetCritical(v int64) *NotFixedForGetRiskStatOutput {
	s.Critical = &v
	return s
}

// SetHigh sets the High field's value.
func (s *NotFixedForGetRiskStatOutput) SetHigh(v int64) *NotFixedForGetRiskStatOutput {
	s.High = &v
	return s
}

// SetLow sets the Low field's value.
func (s *NotFixedForGetRiskStatOutput) SetLow(v int64) *NotFixedForGetRiskStatOutput {
	s.Low = &v
	return s
}

// SetMedium sets the Medium field's value.
func (s *NotFixedForGetRiskStatOutput) SetMedium(v int64) *NotFixedForGetRiskStatOutput {
	s.Medium = &v
	return s
}

// SetSecurity sets the Security field's value.
func (s *NotFixedForGetRiskStatOutput) SetSecurity(v int64) *NotFixedForGetRiskStatOutput {
	s.Security = &v
	return s
}

// SetTotal sets the Total field's value.
func (s *NotFixedForGetRiskStatOutput) SetTotal(v int64) *NotFixedForGetRiskStatOutput {
	s.Total = &v
	return s
}

type StatByStatusForGetRiskStatOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Handled *int64 `type:"int64" json:",omitempty"`

	Ignored *int64 `type:"int64" json:",omitempty"`

	Total *int64 `type:"int64" json:",omitempty"`

	Unhandled *int64 `type:"int64" json:",omitempty"`

	Whitened *int64 `type:"int64" json:",omitempty"`
}

// String returns the string representation
func (s StatByStatusForGetRiskStatOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s StatByStatusForGetRiskStatOutput) GoString() string {
	return s.String()
}

// SetHandled sets the Handled field's value.
func (s *StatByStatusForGetRiskStatOutput) SetHandled(v int64) *StatByStatusForGetRiskStatOutput {
	s.Handled = &v
	return s
}

// SetIgnored sets the Ignored field's value.
func (s *StatByStatusForGetRiskStatOutput) SetIgnored(v int64) *StatByStatusForGetRiskStatOutput {
	s.Ignored = &v
	return s
}

// SetTotal sets the Total field's value.
func (s *StatByStatusForGetRiskStatOutput) SetTotal(v int64) *StatByStatusForGetRiskStatOutput {
	s.Total = &v
	return s
}

// SetUnhandled sets the Unhandled field's value.
func (s *StatByStatusForGetRiskStatOutput) SetUnhandled(v int64) *StatByStatusForGetRiskStatOutput {
	s.Unhandled = &v
	return s
}

// SetWhitened sets the Whitened field's value.
func (s *StatByStatusForGetRiskStatOutput) SetWhitened(v int64) *StatByStatusForGetRiskStatOutput {
	s.Whitened = &v
	return s
}

type TimeRangeForGetRiskStatInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	EndTimeMilli *int64 `type:"int64" json:",omitempty"`

	StartTimeMilli *int64 `type:"int64" json:",omitempty"`
}

// String returns the string representation
func (s TimeRangeForGetRiskStatInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s TimeRangeForGetRiskStatInput) GoString() string {
	return s.String()
}

// SetEndTimeMilli sets the EndTimeMilli field's value.
func (s *TimeRangeForGetRiskStatInput) SetEndTimeMilli(v int64) *TimeRangeForGetRiskStatInput {
	s.EndTimeMilli = &v
	return s
}

// SetStartTimeMilli sets the StartTimeMilli field's value.
func (s *TimeRangeForGetRiskStatInput) SetStartTimeMilli(v int64) *TimeRangeForGetRiskStatInput {
	s.StartTimeMilli = &v
	return s
}

type TotalForGetRiskStatOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Critical *int64 `type:"int64" json:",omitempty"`

	High *int64 `type:"int64" json:",omitempty"`

	Low *int64 `type:"int64" json:",omitempty"`

	Medium *int64 `type:"int64" json:",omitempty"`

	Security *int64 `type:"int64" json:",omitempty"`

	Total *int64 `type:"int64" json:",omitempty"`
}

// String returns the string representation
func (s TotalForGetRiskStatOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s TotalForGetRiskStatOutput) GoString() string {
	return s.String()
}

// SetCritical sets the Critical field's value.
func (s *TotalForGetRiskStatOutput) SetCritical(v int64) *TotalForGetRiskStatOutput {
	s.Critical = &v
	return s
}

// SetHigh sets the High field's value.
func (s *TotalForGetRiskStatOutput) SetHigh(v int64) *TotalForGetRiskStatOutput {
	s.High = &v
	return s
}

// SetLow sets the Low field's value.
func (s *TotalForGetRiskStatOutput) SetLow(v int64) *TotalForGetRiskStatOutput {
	s.Low = &v
	return s
}

// SetMedium sets the Medium field's value.
func (s *TotalForGetRiskStatOutput) SetMedium(v int64) *TotalForGetRiskStatOutput {
	s.Medium = &v
	return s
}

// SetSecurity sets the Security field's value.
func (s *TotalForGetRiskStatOutput) SetSecurity(v int64) *TotalForGetRiskStatOutput {
	s.Security = &v
	return s
}

// SetTotal sets the Total field's value.
func (s *TotalForGetRiskStatOutput) SetTotal(v int64) *TotalForGetRiskStatOutput {
	s.Total = &v
	return s
}

type TrendForGetRiskStatOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Total *TotalForGetRiskStatOutput `type:"structure" json:",omitempty"`

	TrendInfos []*TrendInfoForGetRiskStatOutput `type:"list" json:",omitempty"`
}

// String returns the string representation
func (s TrendForGetRiskStatOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s TrendForGetRiskStatOutput) GoString() string {
	return s.String()
}

// SetTotal sets the Total field's value.
func (s *TrendForGetRiskStatOutput) SetTotal(v *TotalForGetRiskStatOutput) *TrendForGetRiskStatOutput {
	s.Total = v
	return s
}

// SetTrendInfos sets the TrendInfos field's value.
func (s *TrendForGetRiskStatOutput) SetTrendInfos(v []*TrendInfoForGetRiskStatOutput) *TrendForGetRiskStatOutput {
	s.TrendInfos = v
	return s
}

type TrendInfoForGetRiskStatOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Date *string `type:"string" json:",omitempty"`

	Hour *string `type:"string" json:",omitempty"`

	NewAdded *TotalForGetRiskStatOutput `type:"structure" json:",omitempty"`

	Stat *TotalForGetRiskStatOutput `type:"structure" json:",omitempty"`
}

// String returns the string representation
func (s TrendInfoForGetRiskStatOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s TrendInfoForGetRiskStatOutput) GoString() string {
	return s.String()
}

// SetDate sets the Date field's value.
func (s *TrendInfoForGetRiskStatOutput) SetDate(v string) *TrendInfoForGetRiskStatOutput {
	s.Date = &v
	return s
}

// SetHour sets the Hour field's value.
func (s *TrendInfoForGetRiskStatOutput) SetHour(v string) *TrendInfoForGetRiskStatOutput {
	s.Hour = &v
	return s
}

// SetNewAdded sets the NewAdded field's value.
func (s *TrendInfoForGetRiskStatOutput) SetNewAdded(v *TotalForGetRiskStatOutput) *TrendInfoForGetRiskStatOutput {
	s.NewAdded = v
	return s
}

// SetStat sets the Stat field's value.
func (s *TrendInfoForGetRiskStatOutput) SetStat(v *TotalForGetRiskStatOutput) *TrendInfoForGetRiskStatOutput {
	s.Stat = v
	return s
}

const (
	// EnumOfCloudVendorForGetRiskStatInputVolcengine is a EnumOfCloudVendorForGetRiskStatInput enum value
	EnumOfCloudVendorForGetRiskStatInputVolcengine = "volcengine"

	// EnumOfCloudVendorForGetRiskStatInputAliyun is a EnumOfCloudVendorForGetRiskStatInput enum value
	EnumOfCloudVendorForGetRiskStatInputAliyun = "aliyun"

	// EnumOfCloudVendorForGetRiskStatInputHuaweicloud is a EnumOfCloudVendorForGetRiskStatInput enum value
	EnumOfCloudVendorForGetRiskStatInputHuaweicloud = "huaweicloud"

	// EnumOfCloudVendorForGetRiskStatInputTencent is a EnumOfCloudVendorForGetRiskStatInput enum value
	EnumOfCloudVendorForGetRiskStatInputTencent = "tencent"
)

const (
	// EnumOfCloudVendorListForGetRiskStatInputVolcengine is a EnumOfCloudVendorListForGetRiskStatInput enum value
	EnumOfCloudVendorListForGetRiskStatInputVolcengine = "volcengine"

	// EnumOfCloudVendorListForGetRiskStatInputAliyun is a EnumOfCloudVendorListForGetRiskStatInput enum value
	EnumOfCloudVendorListForGetRiskStatInputAliyun = "aliyun"

	// EnumOfCloudVendorListForGetRiskStatInputHuaweicloud is a EnumOfCloudVendorListForGetRiskStatInput enum value
	EnumOfCloudVendorListForGetRiskStatInputHuaweicloud = "huaweicloud"

	// EnumOfCloudVendorListForGetRiskStatInputTencent is a EnumOfCloudVendorListForGetRiskStatInput enum value
	EnumOfCloudVendorListForGetRiskStatInputTencent = "tencent"
)

const (
	// EnumOfDisplayModeForGetRiskStatInputDate is a EnumOfDisplayModeForGetRiskStatInput enum value
	EnumOfDisplayModeForGetRiskStatInputDate = "date"

	// EnumOfDisplayModeForGetRiskStatInputDateHour is a EnumOfDisplayModeForGetRiskStatInput enum value
	EnumOfDisplayModeForGetRiskStatInputDateHour = "date-hour"
)

const (
	// EnumOfRiskTypeForGetRiskStatInputRiskModelVulnerability is a EnumOfRiskTypeForGetRiskStatInput enum value
	EnumOfRiskTypeForGetRiskStatInputRiskModelVulnerability = "risk_model_vulnerability"

	// EnumOfRiskTypeForGetRiskStatInputRiskModelMcStrategy is a EnumOfRiskTypeForGetRiskStatInput enum value
	EnumOfRiskTypeForGetRiskStatInputRiskModelMcStrategy = "risk_model_mc_strategy"

	// EnumOfRiskTypeForGetRiskStatInputRiskModelAttackChainAnalysis is a EnumOfRiskTypeForGetRiskStatInput enum value
	EnumOfRiskTypeForGetRiskStatInputRiskModelAttackChainAnalysis = "risk_model_attack_chain_analysis"
)

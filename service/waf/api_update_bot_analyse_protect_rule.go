// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package waf

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opUpdateBotAnalyseProtectRuleCommon = "UpdateBotAnalyseProtectRule"

// UpdateBotAnalyseProtectRuleCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the UpdateBotAnalyseProtectRuleCommon operation. The "output" return
// value will be populated with the UpdateBotAnalyseProtectRuleCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned UpdateBotAnalyseProtectRuleCommon Request to send the API call to the service.
// the "output" return value is not valid until after UpdateBotAnalyseProtectRuleCommon Send returns without error.
//
// See UpdateBotAnalyseProtectRuleCommon for more information on using the UpdateBotAnalyseProtectRuleCommon
// API call, and error handling.
//
//    // Example sending a request using the UpdateBotAnalyseProtectRuleCommonRequest method.
//    req, resp := client.UpdateBotAnalyseProtectRuleCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *WAF) UpdateBotAnalyseProtectRuleCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opUpdateBotAnalyseProtectRuleCommon,
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

// UpdateBotAnalyseProtectRuleCommon API operation for WAF.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for WAF's
// API operation UpdateBotAnalyseProtectRuleCommon for usage and error information.
func (c *WAF) UpdateBotAnalyseProtectRuleCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.UpdateBotAnalyseProtectRuleCommonRequest(input)
	return out, req.Send()
}

// UpdateBotAnalyseProtectRuleCommonWithContext is the same as UpdateBotAnalyseProtectRuleCommon with the addition of
// the ability to pass a context and additional request options.
//
// See UpdateBotAnalyseProtectRuleCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *WAF) UpdateBotAnalyseProtectRuleCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.UpdateBotAnalyseProtectRuleCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opUpdateBotAnalyseProtectRule = "UpdateBotAnalyseProtectRule"

// UpdateBotAnalyseProtectRuleRequest generates a "volcengine/request.Request" representing the
// client's request for the UpdateBotAnalyseProtectRule operation. The "output" return
// value will be populated with the UpdateBotAnalyseProtectRuleCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned UpdateBotAnalyseProtectRuleCommon Request to send the API call to the service.
// the "output" return value is not valid until after UpdateBotAnalyseProtectRuleCommon Send returns without error.
//
// See UpdateBotAnalyseProtectRule for more information on using the UpdateBotAnalyseProtectRule
// API call, and error handling.
//
//    // Example sending a request using the UpdateBotAnalyseProtectRuleRequest method.
//    req, resp := client.UpdateBotAnalyseProtectRuleRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *WAF) UpdateBotAnalyseProtectRuleRequest(input *UpdateBotAnalyseProtectRuleInput) (req *request.Request, output *UpdateBotAnalyseProtectRuleOutput) {
	op := &request.Operation{
		Name:       opUpdateBotAnalyseProtectRule,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateBotAnalyseProtectRuleInput{}
	}

	output = &UpdateBotAnalyseProtectRuleOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// UpdateBotAnalyseProtectRule API operation for WAF.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for WAF's
// API operation UpdateBotAnalyseProtectRule for usage and error information.
func (c *WAF) UpdateBotAnalyseProtectRule(input *UpdateBotAnalyseProtectRuleInput) (*UpdateBotAnalyseProtectRuleOutput, error) {
	req, out := c.UpdateBotAnalyseProtectRuleRequest(input)
	return out, req.Send()
}

// UpdateBotAnalyseProtectRuleWithContext is the same as UpdateBotAnalyseProtectRule with the addition of
// the ability to pass a context and additional request options.
//
// See UpdateBotAnalyseProtectRule for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *WAF) UpdateBotAnalyseProtectRuleWithContext(ctx volcengine.Context, input *UpdateBotAnalyseProtectRuleInput, opts ...request.Option) (*UpdateBotAnalyseProtectRuleOutput, error) {
	req, out := c.UpdateBotAnalyseProtectRuleRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type AccurateGroupForUpdateBotAnalyseProtectRuleInput struct {
	_ struct{} `type:"structure"`

	AccurateRules []*AccurateRuleForUpdateBotAnalyseProtectRuleInput `type:"list"`

	Logic *int32 `type:"int32"`
}

// String returns the string representation
func (s AccurateGroupForUpdateBotAnalyseProtectRuleInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s AccurateGroupForUpdateBotAnalyseProtectRuleInput) GoString() string {
	return s.String()
}

// SetAccurateRules sets the AccurateRules field's value.
func (s *AccurateGroupForUpdateBotAnalyseProtectRuleInput) SetAccurateRules(v []*AccurateRuleForUpdateBotAnalyseProtectRuleInput) *AccurateGroupForUpdateBotAnalyseProtectRuleInput {
	s.AccurateRules = v
	return s
}

// SetLogic sets the Logic field's value.
func (s *AccurateGroupForUpdateBotAnalyseProtectRuleInput) SetLogic(v int32) *AccurateGroupForUpdateBotAnalyseProtectRuleInput {
	s.Logic = &v
	return s
}

type AccurateRuleForUpdateBotAnalyseProtectRuleInput struct {
	_ struct{} `type:"structure"`

	HttpObj *string `type:"string"`

	ObjType *int32 `type:"int32"`

	Opretar *int32 `type:"int32"`

	Property *int32 `type:"int32"`

	ValueString *string `type:"string"`
}

// String returns the string representation
func (s AccurateRuleForUpdateBotAnalyseProtectRuleInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s AccurateRuleForUpdateBotAnalyseProtectRuleInput) GoString() string {
	return s.String()
}

// SetHttpObj sets the HttpObj field's value.
func (s *AccurateRuleForUpdateBotAnalyseProtectRuleInput) SetHttpObj(v string) *AccurateRuleForUpdateBotAnalyseProtectRuleInput {
	s.HttpObj = &v
	return s
}

// SetObjType sets the ObjType field's value.
func (s *AccurateRuleForUpdateBotAnalyseProtectRuleInput) SetObjType(v int32) *AccurateRuleForUpdateBotAnalyseProtectRuleInput {
	s.ObjType = &v
	return s
}

// SetOpretar sets the Opretar field's value.
func (s *AccurateRuleForUpdateBotAnalyseProtectRuleInput) SetOpretar(v int32) *AccurateRuleForUpdateBotAnalyseProtectRuleInput {
	s.Opretar = &v
	return s
}

// SetProperty sets the Property field's value.
func (s *AccurateRuleForUpdateBotAnalyseProtectRuleInput) SetProperty(v int32) *AccurateRuleForUpdateBotAnalyseProtectRuleInput {
	s.Property = &v
	return s
}

// SetValueString sets the ValueString field's value.
func (s *AccurateRuleForUpdateBotAnalyseProtectRuleInput) SetValueString(v string) *AccurateRuleForUpdateBotAnalyseProtectRuleInput {
	s.ValueString = &v
	return s
}

type UpdateBotAnalyseProtectRuleInput struct {
	_ struct{} `type:"structure"`

	AccurateGroup []*AccurateGroupForUpdateBotAnalyseProtectRuleInput `type:"list"`

	ActionAfterVerification *int32 `type:"int32"`

	// ActionType is a required field
	ActionType *int32 `type:"int32" required:"true"`

	// EffectTime is a required field
	EffectTime *int32 `type:"int32" required:"true"`

	// Enable is a required field
	Enable *int32 `type:"int32" required:"true"`

	ExemptionTime *int32 `type:"int32"`

	// Field is a required field
	Field *string `type:"string" required:"true"`

	// Host is a required field
	Host *string `type:"string" required:"true"`

	// Id is a required field
	Id *int32 `type:"int32" required:"true"`

	// Name is a required field
	Name *string `type:"string" required:"true"`

	// Path is a required field
	Path *string `type:"string" required:"true"`

	PathThreshold *int32 `type:"int32"`

	// RulePriority is a required field
	RulePriority *int32 `type:"int32" required:"true"`

	SingleProportion *string `type:"string"`

	// SingleThreshold is a required field
	SingleThreshold *int32 `type:"int32" required:"true"`

	// StatisticalDuration is a required field
	StatisticalDuration *int32 `type:"int32" required:"true"`

	// StatisticalType is a required field
	StatisticalType *int32 `type:"int32" required:"true"`
}

// String returns the string representation
func (s UpdateBotAnalyseProtectRuleInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s UpdateBotAnalyseProtectRuleInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateBotAnalyseProtectRuleInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "UpdateBotAnalyseProtectRuleInput"}
	if s.ActionType == nil {
		invalidParams.Add(request.NewErrParamRequired("ActionType"))
	}
	if s.EffectTime == nil {
		invalidParams.Add(request.NewErrParamRequired("EffectTime"))
	}
	if s.Enable == nil {
		invalidParams.Add(request.NewErrParamRequired("Enable"))
	}
	if s.Field == nil {
		invalidParams.Add(request.NewErrParamRequired("Field"))
	}
	if s.Host == nil {
		invalidParams.Add(request.NewErrParamRequired("Host"))
	}
	if s.Id == nil {
		invalidParams.Add(request.NewErrParamRequired("Id"))
	}
	if s.Name == nil {
		invalidParams.Add(request.NewErrParamRequired("Name"))
	}
	if s.Path == nil {
		invalidParams.Add(request.NewErrParamRequired("Path"))
	}
	if s.RulePriority == nil {
		invalidParams.Add(request.NewErrParamRequired("RulePriority"))
	}
	if s.SingleThreshold == nil {
		invalidParams.Add(request.NewErrParamRequired("SingleThreshold"))
	}
	if s.StatisticalDuration == nil {
		invalidParams.Add(request.NewErrParamRequired("StatisticalDuration"))
	}
	if s.StatisticalType == nil {
		invalidParams.Add(request.NewErrParamRequired("StatisticalType"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetAccurateGroup sets the AccurateGroup field's value.
func (s *UpdateBotAnalyseProtectRuleInput) SetAccurateGroup(v []*AccurateGroupForUpdateBotAnalyseProtectRuleInput) *UpdateBotAnalyseProtectRuleInput {
	s.AccurateGroup = v
	return s
}

// SetActionAfterVerification sets the ActionAfterVerification field's value.
func (s *UpdateBotAnalyseProtectRuleInput) SetActionAfterVerification(v int32) *UpdateBotAnalyseProtectRuleInput {
	s.ActionAfterVerification = &v
	return s
}

// SetActionType sets the ActionType field's value.
func (s *UpdateBotAnalyseProtectRuleInput) SetActionType(v int32) *UpdateBotAnalyseProtectRuleInput {
	s.ActionType = &v
	return s
}

// SetEffectTime sets the EffectTime field's value.
func (s *UpdateBotAnalyseProtectRuleInput) SetEffectTime(v int32) *UpdateBotAnalyseProtectRuleInput {
	s.EffectTime = &v
	return s
}

// SetEnable sets the Enable field's value.
func (s *UpdateBotAnalyseProtectRuleInput) SetEnable(v int32) *UpdateBotAnalyseProtectRuleInput {
	s.Enable = &v
	return s
}

// SetExemptionTime sets the ExemptionTime field's value.
func (s *UpdateBotAnalyseProtectRuleInput) SetExemptionTime(v int32) *UpdateBotAnalyseProtectRuleInput {
	s.ExemptionTime = &v
	return s
}

// SetField sets the Field field's value.
func (s *UpdateBotAnalyseProtectRuleInput) SetField(v string) *UpdateBotAnalyseProtectRuleInput {
	s.Field = &v
	return s
}

// SetHost sets the Host field's value.
func (s *UpdateBotAnalyseProtectRuleInput) SetHost(v string) *UpdateBotAnalyseProtectRuleInput {
	s.Host = &v
	return s
}

// SetId sets the Id field's value.
func (s *UpdateBotAnalyseProtectRuleInput) SetId(v int32) *UpdateBotAnalyseProtectRuleInput {
	s.Id = &v
	return s
}

// SetName sets the Name field's value.
func (s *UpdateBotAnalyseProtectRuleInput) SetName(v string) *UpdateBotAnalyseProtectRuleInput {
	s.Name = &v
	return s
}

// SetPath sets the Path field's value.
func (s *UpdateBotAnalyseProtectRuleInput) SetPath(v string) *UpdateBotAnalyseProtectRuleInput {
	s.Path = &v
	return s
}

// SetPathThreshold sets the PathThreshold field's value.
func (s *UpdateBotAnalyseProtectRuleInput) SetPathThreshold(v int32) *UpdateBotAnalyseProtectRuleInput {
	s.PathThreshold = &v
	return s
}

// SetRulePriority sets the RulePriority field's value.
func (s *UpdateBotAnalyseProtectRuleInput) SetRulePriority(v int32) *UpdateBotAnalyseProtectRuleInput {
	s.RulePriority = &v
	return s
}

// SetSingleProportion sets the SingleProportion field's value.
func (s *UpdateBotAnalyseProtectRuleInput) SetSingleProportion(v string) *UpdateBotAnalyseProtectRuleInput {
	s.SingleProportion = &v
	return s
}

// SetSingleThreshold sets the SingleThreshold field's value.
func (s *UpdateBotAnalyseProtectRuleInput) SetSingleThreshold(v int32) *UpdateBotAnalyseProtectRuleInput {
	s.SingleThreshold = &v
	return s
}

// SetStatisticalDuration sets the StatisticalDuration field's value.
func (s *UpdateBotAnalyseProtectRuleInput) SetStatisticalDuration(v int32) *UpdateBotAnalyseProtectRuleInput {
	s.StatisticalDuration = &v
	return s
}

// SetStatisticalType sets the StatisticalType field's value.
func (s *UpdateBotAnalyseProtectRuleInput) SetStatisticalType(v int32) *UpdateBotAnalyseProtectRuleInput {
	s.StatisticalType = &v
	return s
}

type UpdateBotAnalyseProtectRuleOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s UpdateBotAnalyseProtectRuleOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s UpdateBotAnalyseProtectRuleOutput) GoString() string {
	return s.String()
}
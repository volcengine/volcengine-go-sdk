// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package autoscaling

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opModifyScalingPolicyCommon = "ModifyScalingPolicy"

// ModifyScalingPolicyCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ModifyScalingPolicyCommon operation. The "output" return
// value will be populated with the ModifyScalingPolicyCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyScalingPolicyCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyScalingPolicyCommon Send returns without error.
//
// See ModifyScalingPolicyCommon for more information on using the ModifyScalingPolicyCommon
// API call, and error handling.
//
//    // Example sending a request using the ModifyScalingPolicyCommonRequest method.
//    req, resp := client.ModifyScalingPolicyCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *AUTOSCALING) ModifyScalingPolicyCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opModifyScalingPolicyCommon,
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

// ModifyScalingPolicyCommon API operation for AUTO_SCALING.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for AUTO_SCALING's
// API operation ModifyScalingPolicyCommon for usage and error information.
func (c *AUTOSCALING) ModifyScalingPolicyCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ModifyScalingPolicyCommonRequest(input)
	return out, req.Send()
}

// ModifyScalingPolicyCommonWithContext is the same as ModifyScalingPolicyCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyScalingPolicyCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *AUTOSCALING) ModifyScalingPolicyCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ModifyScalingPolicyCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opModifyScalingPolicy = "ModifyScalingPolicy"

// ModifyScalingPolicyRequest generates a "volcengine/request.Request" representing the
// client's request for the ModifyScalingPolicy operation. The "output" return
// value will be populated with the ModifyScalingPolicyCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyScalingPolicyCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyScalingPolicyCommon Send returns without error.
//
// See ModifyScalingPolicy for more information on using the ModifyScalingPolicy
// API call, and error handling.
//
//    // Example sending a request using the ModifyScalingPolicyRequest method.
//    req, resp := client.ModifyScalingPolicyRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *AUTOSCALING) ModifyScalingPolicyRequest(input *ModifyScalingPolicyInput) (req *request.Request, output *ModifyScalingPolicyOutput) {
	op := &request.Operation{
		Name:       opModifyScalingPolicy,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ModifyScalingPolicyInput{}
	}

	output = &ModifyScalingPolicyOutput{}
	req = c.newRequest(op, input, output)

	return
}

// ModifyScalingPolicy API operation for AUTO_SCALING.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for AUTO_SCALING's
// API operation ModifyScalingPolicy for usage and error information.
func (c *AUTOSCALING) ModifyScalingPolicy(input *ModifyScalingPolicyInput) (*ModifyScalingPolicyOutput, error) {
	req, out := c.ModifyScalingPolicyRequest(input)
	return out, req.Send()
}

// ModifyScalingPolicyWithContext is the same as ModifyScalingPolicy with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyScalingPolicy for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *AUTOSCALING) ModifyScalingPolicyWithContext(ctx volcengine.Context, input *ModifyScalingPolicyInput, opts ...request.Option) (*ModifyScalingPolicyOutput, error) {
	req, out := c.ModifyScalingPolicyRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type AlarmPolicyConditionForModifyScalingPolicyInput struct {
	_ struct{} `type:"structure"`

	ComparisonOperator *string `type:"string"`

	MetricName *string `type:"string"`

	MetricUnit *string `type:"string"`

	Threshold *string `type:"string"`
}

// String returns the string representation
func (s AlarmPolicyConditionForModifyScalingPolicyInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s AlarmPolicyConditionForModifyScalingPolicyInput) GoString() string {
	return s.String()
}

// SetComparisonOperator sets the ComparisonOperator field's value.
func (s *AlarmPolicyConditionForModifyScalingPolicyInput) SetComparisonOperator(v string) *AlarmPolicyConditionForModifyScalingPolicyInput {
	s.ComparisonOperator = &v
	return s
}

// SetMetricName sets the MetricName field's value.
func (s *AlarmPolicyConditionForModifyScalingPolicyInput) SetMetricName(v string) *AlarmPolicyConditionForModifyScalingPolicyInput {
	s.MetricName = &v
	return s
}

// SetMetricUnit sets the MetricUnit field's value.
func (s *AlarmPolicyConditionForModifyScalingPolicyInput) SetMetricUnit(v string) *AlarmPolicyConditionForModifyScalingPolicyInput {
	s.MetricUnit = &v
	return s
}

// SetThreshold sets the Threshold field's value.
func (s *AlarmPolicyConditionForModifyScalingPolicyInput) SetThreshold(v string) *AlarmPolicyConditionForModifyScalingPolicyInput {
	s.Threshold = &v
	return s
}

type AlarmPolicyForModifyScalingPolicyInput struct {
	_ struct{} `type:"structure"`

	Condition *AlarmPolicyConditionForModifyScalingPolicyInput `type:"structure"`

	EvaluationCount *int32 `min:"1" max:"180" type:"int32"`

	RuleType *string `type:"string"`
}

// String returns the string representation
func (s AlarmPolicyForModifyScalingPolicyInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s AlarmPolicyForModifyScalingPolicyInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AlarmPolicyForModifyScalingPolicyInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "AlarmPolicyForModifyScalingPolicyInput"}
	if s.EvaluationCount != nil && *s.EvaluationCount < 1 {
		invalidParams.Add(request.NewErrParamMinValue("EvaluationCount", 1))
	}
	if s.EvaluationCount != nil && *s.EvaluationCount > 180 {
		invalidParams.Add(request.NewErrParamMaxValue("EvaluationCount", 180))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetCondition sets the Condition field's value.
func (s *AlarmPolicyForModifyScalingPolicyInput) SetCondition(v *AlarmPolicyConditionForModifyScalingPolicyInput) *AlarmPolicyForModifyScalingPolicyInput {
	s.Condition = v
	return s
}

// SetEvaluationCount sets the EvaluationCount field's value.
func (s *AlarmPolicyForModifyScalingPolicyInput) SetEvaluationCount(v int32) *AlarmPolicyForModifyScalingPolicyInput {
	s.EvaluationCount = &v
	return s
}

// SetRuleType sets the RuleType field's value.
func (s *AlarmPolicyForModifyScalingPolicyInput) SetRuleType(v string) *AlarmPolicyForModifyScalingPolicyInput {
	s.RuleType = &v
	return s
}

type ModifyScalingPolicyInput struct {
	_ struct{} `type:"structure"`

	AdjustmentType *string `type:"string"`

	AdjustmentValue *int32 `type:"int32"`

	AlarmPolicy *AlarmPolicyForModifyScalingPolicyInput `type:"structure"`

	Cooldown *int32 `type:"int32"`

	// ScalingPolicyId is a required field
	ScalingPolicyId *string `type:"string" required:"true"`

	ScalingPolicyName *string `type:"string"`

	ScheduledPolicy *ScheduledPolicyForModifyScalingPolicyInput `type:"structure"`
}

// String returns the string representation
func (s ModifyScalingPolicyInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyScalingPolicyInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ModifyScalingPolicyInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ModifyScalingPolicyInput"}
	if s.ScalingPolicyId == nil {
		invalidParams.Add(request.NewErrParamRequired("ScalingPolicyId"))
	}
	if s.AlarmPolicy != nil {
		if err := s.AlarmPolicy.Validate(); err != nil {
			invalidParams.AddNested("AlarmPolicy", err.(request.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetAdjustmentType sets the AdjustmentType field's value.
func (s *ModifyScalingPolicyInput) SetAdjustmentType(v string) *ModifyScalingPolicyInput {
	s.AdjustmentType = &v
	return s
}

// SetAdjustmentValue sets the AdjustmentValue field's value.
func (s *ModifyScalingPolicyInput) SetAdjustmentValue(v int32) *ModifyScalingPolicyInput {
	s.AdjustmentValue = &v
	return s
}

// SetAlarmPolicy sets the AlarmPolicy field's value.
func (s *ModifyScalingPolicyInput) SetAlarmPolicy(v *AlarmPolicyForModifyScalingPolicyInput) *ModifyScalingPolicyInput {
	s.AlarmPolicy = v
	return s
}

// SetCooldown sets the Cooldown field's value.
func (s *ModifyScalingPolicyInput) SetCooldown(v int32) *ModifyScalingPolicyInput {
	s.Cooldown = &v
	return s
}

// SetScalingPolicyId sets the ScalingPolicyId field's value.
func (s *ModifyScalingPolicyInput) SetScalingPolicyId(v string) *ModifyScalingPolicyInput {
	s.ScalingPolicyId = &v
	return s
}

// SetScalingPolicyName sets the ScalingPolicyName field's value.
func (s *ModifyScalingPolicyInput) SetScalingPolicyName(v string) *ModifyScalingPolicyInput {
	s.ScalingPolicyName = &v
	return s
}

// SetScheduledPolicy sets the ScheduledPolicy field's value.
func (s *ModifyScalingPolicyInput) SetScheduledPolicy(v *ScheduledPolicyForModifyScalingPolicyInput) *ModifyScalingPolicyInput {
	s.ScheduledPolicy = v
	return s
}

type ModifyScalingPolicyOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	ScalingPolicyId *string `type:"string"`
}

// String returns the string representation
func (s ModifyScalingPolicyOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyScalingPolicyOutput) GoString() string {
	return s.String()
}

// SetScalingPolicyId sets the ScalingPolicyId field's value.
func (s *ModifyScalingPolicyOutput) SetScalingPolicyId(v string) *ModifyScalingPolicyOutput {
	s.ScalingPolicyId = &v
	return s
}

type ScheduledPolicyForModifyScalingPolicyInput struct {
	_ struct{} `type:"structure"`

	LaunchTime *string `type:"string"`

	RecurrenceEndTime *string `type:"string"`

	RecurrenceType *string `type:"string"`

	RecurrenceValue *string `type:"string"`
}

// String returns the string representation
func (s ScheduledPolicyForModifyScalingPolicyInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ScheduledPolicyForModifyScalingPolicyInput) GoString() string {
	return s.String()
}

// SetLaunchTime sets the LaunchTime field's value.
func (s *ScheduledPolicyForModifyScalingPolicyInput) SetLaunchTime(v string) *ScheduledPolicyForModifyScalingPolicyInput {
	s.LaunchTime = &v
	return s
}

// SetRecurrenceEndTime sets the RecurrenceEndTime field's value.
func (s *ScheduledPolicyForModifyScalingPolicyInput) SetRecurrenceEndTime(v string) *ScheduledPolicyForModifyScalingPolicyInput {
	s.RecurrenceEndTime = &v
	return s
}

// SetRecurrenceType sets the RecurrenceType field's value.
func (s *ScheduledPolicyForModifyScalingPolicyInput) SetRecurrenceType(v string) *ScheduledPolicyForModifyScalingPolicyInput {
	s.RecurrenceType = &v
	return s
}

// SetRecurrenceValue sets the RecurrenceValue field's value.
func (s *ScheduledPolicyForModifyScalingPolicyInput) SetRecurrenceValue(v string) *ScheduledPolicyForModifyScalingPolicyInput {
	s.RecurrenceValue = &v
	return s
}

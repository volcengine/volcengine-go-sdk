// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package mcdn

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribeAlertStrategyCommon = "DescribeAlertStrategy"

// DescribeAlertStrategyCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeAlertStrategyCommon operation. The "output" return
// value will be populated with the DescribeAlertStrategyCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeAlertStrategyCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeAlertStrategyCommon Send returns without error.
//
// See DescribeAlertStrategyCommon for more information on using the DescribeAlertStrategyCommon
// API call, and error handling.
//
//    // Example sending a request using the DescribeAlertStrategyCommonRequest method.
//    req, resp := client.DescribeAlertStrategyCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *MCDN) DescribeAlertStrategyCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeAlertStrategyCommon,
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

// DescribeAlertStrategyCommon API operation for MCDN.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for MCDN's
// API operation DescribeAlertStrategyCommon for usage and error information.
func (c *MCDN) DescribeAlertStrategyCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeAlertStrategyCommonRequest(input)
	return out, req.Send()
}

// DescribeAlertStrategyCommonWithContext is the same as DescribeAlertStrategyCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeAlertStrategyCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *MCDN) DescribeAlertStrategyCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeAlertStrategyCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeAlertStrategy = "DescribeAlertStrategy"

// DescribeAlertStrategyRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeAlertStrategy operation. The "output" return
// value will be populated with the DescribeAlertStrategyCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeAlertStrategyCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeAlertStrategyCommon Send returns without error.
//
// See DescribeAlertStrategy for more information on using the DescribeAlertStrategy
// API call, and error handling.
//
//    // Example sending a request using the DescribeAlertStrategyRequest method.
//    req, resp := client.DescribeAlertStrategyRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *MCDN) DescribeAlertStrategyRequest(input *DescribeAlertStrategyInput) (req *request.Request, output *DescribeAlertStrategyOutput) {
	op := &request.Operation{
		Name:       opDescribeAlertStrategy,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeAlertStrategyInput{}
	}

	output = &DescribeAlertStrategyOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// DescribeAlertStrategy API operation for MCDN.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for MCDN's
// API operation DescribeAlertStrategy for usage and error information.
func (c *MCDN) DescribeAlertStrategy(input *DescribeAlertStrategyInput) (*DescribeAlertStrategyOutput, error) {
	req, out := c.DescribeAlertStrategyRequest(input)
	return out, req.Send()
}

// DescribeAlertStrategyWithContext is the same as DescribeAlertStrategy with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeAlertStrategy for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *MCDN) DescribeAlertStrategyWithContext(ctx volcengine.Context, input *DescribeAlertStrategyInput, opts ...request.Option) (*DescribeAlertStrategyOutput, error) {
	req, out := c.DescribeAlertStrategyRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type AlertRuleForDescribeAlertStrategyOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Conditions []*ConditionForDescribeAlertStrategyOutput `type:"list" json:",omitempty"`

	EnableEndTime *string `type:"string" json:",omitempty"`

	EnableStartTime *string `type:"string" json:",omitempty"`

	Frequency *int64 `type:"int64" json:",omitempty"`

	Level *string `type:"string" json:",omitempty"`

	Logic *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s AlertRuleForDescribeAlertStrategyOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s AlertRuleForDescribeAlertStrategyOutput) GoString() string {
	return s.String()
}

// SetConditions sets the Conditions field's value.
func (s *AlertRuleForDescribeAlertStrategyOutput) SetConditions(v []*ConditionForDescribeAlertStrategyOutput) *AlertRuleForDescribeAlertStrategyOutput {
	s.Conditions = v
	return s
}

// SetEnableEndTime sets the EnableEndTime field's value.
func (s *AlertRuleForDescribeAlertStrategyOutput) SetEnableEndTime(v string) *AlertRuleForDescribeAlertStrategyOutput {
	s.EnableEndTime = &v
	return s
}

// SetEnableStartTime sets the EnableStartTime field's value.
func (s *AlertRuleForDescribeAlertStrategyOutput) SetEnableStartTime(v string) *AlertRuleForDescribeAlertStrategyOutput {
	s.EnableStartTime = &v
	return s
}

// SetFrequency sets the Frequency field's value.
func (s *AlertRuleForDescribeAlertStrategyOutput) SetFrequency(v int64) *AlertRuleForDescribeAlertStrategyOutput {
	s.Frequency = &v
	return s
}

// SetLevel sets the Level field's value.
func (s *AlertRuleForDescribeAlertStrategyOutput) SetLevel(v string) *AlertRuleForDescribeAlertStrategyOutput {
	s.Level = &v
	return s
}

// SetLogic sets the Logic field's value.
func (s *AlertRuleForDescribeAlertStrategyOutput) SetLogic(v string) *AlertRuleForDescribeAlertStrategyOutput {
	s.Logic = &v
	return s
}

type AlertStrategyForDescribeAlertStrategyOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	AlertRule *AlertRuleForDescribeAlertStrategyOutput `type:"structure" json:",omitempty"`

	Domains []*DomainForDescribeAlertStrategyOutput `type:"list" json:",omitempty"`

	Id *string `type:"string" json:",omitempty"`

	Name *string `type:"string" json:",omitempty"`

	ProbeTasks []*ProbeTaskForDescribeAlertStrategyOutput `type:"list" json:",omitempty"`

	ResourceScope *string `type:"string" json:",omitempty"`

	ResourceTypes []*string `type:"list" json:",omitempty"`

	Status *string `type:"string" json:",omitempty"`

	SubscribeRule *SubscribeRuleForDescribeAlertStrategyOutput `type:"structure" json:",omitempty"`

	TriggerType *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s AlertStrategyForDescribeAlertStrategyOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s AlertStrategyForDescribeAlertStrategyOutput) GoString() string {
	return s.String()
}

// SetAlertRule sets the AlertRule field's value.
func (s *AlertStrategyForDescribeAlertStrategyOutput) SetAlertRule(v *AlertRuleForDescribeAlertStrategyOutput) *AlertStrategyForDescribeAlertStrategyOutput {
	s.AlertRule = v
	return s
}

// SetDomains sets the Domains field's value.
func (s *AlertStrategyForDescribeAlertStrategyOutput) SetDomains(v []*DomainForDescribeAlertStrategyOutput) *AlertStrategyForDescribeAlertStrategyOutput {
	s.Domains = v
	return s
}

// SetId sets the Id field's value.
func (s *AlertStrategyForDescribeAlertStrategyOutput) SetId(v string) *AlertStrategyForDescribeAlertStrategyOutput {
	s.Id = &v
	return s
}

// SetName sets the Name field's value.
func (s *AlertStrategyForDescribeAlertStrategyOutput) SetName(v string) *AlertStrategyForDescribeAlertStrategyOutput {
	s.Name = &v
	return s
}

// SetProbeTasks sets the ProbeTasks field's value.
func (s *AlertStrategyForDescribeAlertStrategyOutput) SetProbeTasks(v []*ProbeTaskForDescribeAlertStrategyOutput) *AlertStrategyForDescribeAlertStrategyOutput {
	s.ProbeTasks = v
	return s
}

// SetResourceScope sets the ResourceScope field's value.
func (s *AlertStrategyForDescribeAlertStrategyOutput) SetResourceScope(v string) *AlertStrategyForDescribeAlertStrategyOutput {
	s.ResourceScope = &v
	return s
}

// SetResourceTypes sets the ResourceTypes field's value.
func (s *AlertStrategyForDescribeAlertStrategyOutput) SetResourceTypes(v []*string) *AlertStrategyForDescribeAlertStrategyOutput {
	s.ResourceTypes = v
	return s
}

// SetStatus sets the Status field's value.
func (s *AlertStrategyForDescribeAlertStrategyOutput) SetStatus(v string) *AlertStrategyForDescribeAlertStrategyOutput {
	s.Status = &v
	return s
}

// SetSubscribeRule sets the SubscribeRule field's value.
func (s *AlertStrategyForDescribeAlertStrategyOutput) SetSubscribeRule(v *SubscribeRuleForDescribeAlertStrategyOutput) *AlertStrategyForDescribeAlertStrategyOutput {
	s.SubscribeRule = v
	return s
}

// SetTriggerType sets the TriggerType field's value.
func (s *AlertStrategyForDescribeAlertStrategyOutput) SetTriggerType(v string) *AlertStrategyForDescribeAlertStrategyOutput {
	s.TriggerType = &v
	return s
}

type ConditionForDescribeAlertStrategyOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Format *string `type:"string" json:",omitempty"`

	MetricId *string `type:"string" json:",omitempty"`

	Operator *string `type:"string" json:",omitempty"`

	Period *int64 `type:"int64" json:",omitempty"`

	Sensitivity *string `type:"string" json:",omitempty"`

	ThresholdType *string `type:"string" json:",omitempty"`

	Value *float64 `type:"double" json:",omitempty"`
}

// String returns the string representation
func (s ConditionForDescribeAlertStrategyOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ConditionForDescribeAlertStrategyOutput) GoString() string {
	return s.String()
}

// SetFormat sets the Format field's value.
func (s *ConditionForDescribeAlertStrategyOutput) SetFormat(v string) *ConditionForDescribeAlertStrategyOutput {
	s.Format = &v
	return s
}

// SetMetricId sets the MetricId field's value.
func (s *ConditionForDescribeAlertStrategyOutput) SetMetricId(v string) *ConditionForDescribeAlertStrategyOutput {
	s.MetricId = &v
	return s
}

// SetOperator sets the Operator field's value.
func (s *ConditionForDescribeAlertStrategyOutput) SetOperator(v string) *ConditionForDescribeAlertStrategyOutput {
	s.Operator = &v
	return s
}

// SetPeriod sets the Period field's value.
func (s *ConditionForDescribeAlertStrategyOutput) SetPeriod(v int64) *ConditionForDescribeAlertStrategyOutput {
	s.Period = &v
	return s
}

// SetSensitivity sets the Sensitivity field's value.
func (s *ConditionForDescribeAlertStrategyOutput) SetSensitivity(v string) *ConditionForDescribeAlertStrategyOutput {
	s.Sensitivity = &v
	return s
}

// SetThresholdType sets the ThresholdType field's value.
func (s *ConditionForDescribeAlertStrategyOutput) SetThresholdType(v string) *ConditionForDescribeAlertStrategyOutput {
	s.ThresholdType = &v
	return s
}

// SetValue sets the Value field's value.
func (s *ConditionForDescribeAlertStrategyOutput) SetValue(v float64) *ConditionForDescribeAlertStrategyOutput {
	s.Value = &v
	return s
}

type ContactGroupForDescribeAlertStrategyOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Id *string `type:"string" json:",omitempty"`

	Name *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s ContactGroupForDescribeAlertStrategyOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ContactGroupForDescribeAlertStrategyOutput) GoString() string {
	return s.String()
}

// SetId sets the Id field's value.
func (s *ContactGroupForDescribeAlertStrategyOutput) SetId(v string) *ContactGroupForDescribeAlertStrategyOutput {
	s.Id = &v
	return s
}

// SetName sets the Name field's value.
func (s *ContactGroupForDescribeAlertStrategyOutput) SetName(v string) *ContactGroupForDescribeAlertStrategyOutput {
	s.Name = &v
	return s
}

type ContactRobotForDescribeAlertStrategyOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Id *string `type:"string" json:",omitempty"`

	Name *string `type:"string" json:",omitempty"`

	RobotType *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s ContactRobotForDescribeAlertStrategyOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ContactRobotForDescribeAlertStrategyOutput) GoString() string {
	return s.String()
}

// SetId sets the Id field's value.
func (s *ContactRobotForDescribeAlertStrategyOutput) SetId(v string) *ContactRobotForDescribeAlertStrategyOutput {
	s.Id = &v
	return s
}

// SetName sets the Name field's value.
func (s *ContactRobotForDescribeAlertStrategyOutput) SetName(v string) *ContactRobotForDescribeAlertStrategyOutput {
	s.Name = &v
	return s
}

// SetRobotType sets the RobotType field's value.
func (s *ContactRobotForDescribeAlertStrategyOutput) SetRobotType(v string) *ContactRobotForDescribeAlertStrategyOutput {
	s.RobotType = &v
	return s
}

type DescribeAlertStrategyInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	// Id is a required field
	Id *string `type:"string" json:",omitempty" required:"true"`
}

// String returns the string representation
func (s DescribeAlertStrategyInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeAlertStrategyInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeAlertStrategyInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DescribeAlertStrategyInput"}
	if s.Id == nil {
		invalidParams.Add(request.NewErrParamRequired("Id"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetId sets the Id field's value.
func (s *DescribeAlertStrategyInput) SetId(v string) *DescribeAlertStrategyInput {
	s.Id = &v
	return s
}

type DescribeAlertStrategyOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	AlertStrategy *AlertStrategyForDescribeAlertStrategyOutput `type:"structure" json:",omitempty"`
}

// String returns the string representation
func (s DescribeAlertStrategyOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeAlertStrategyOutput) GoString() string {
	return s.String()
}

// SetAlertStrategy sets the AlertStrategy field's value.
func (s *DescribeAlertStrategyOutput) SetAlertStrategy(v *AlertStrategyForDescribeAlertStrategyOutput) *DescribeAlertStrategyOutput {
	s.AlertStrategy = v
	return s
}

type DomainForDescribeAlertStrategyOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Domain *string `type:"string" json:",omitempty"`

	Vendor *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s DomainForDescribeAlertStrategyOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DomainForDescribeAlertStrategyOutput) GoString() string {
	return s.String()
}

// SetDomain sets the Domain field's value.
func (s *DomainForDescribeAlertStrategyOutput) SetDomain(v string) *DomainForDescribeAlertStrategyOutput {
	s.Domain = &v
	return s
}

// SetVendor sets the Vendor field's value.
func (s *DomainForDescribeAlertStrategyOutput) SetVendor(v string) *DomainForDescribeAlertStrategyOutput {
	s.Vendor = &v
	return s
}

type NotifyConfigForDescribeAlertStrategyOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	IgnoreDisabledDomains *bool `type:"boolean" json:",omitempty"`

	Level *string `type:"string" json:",omitempty"`

	SendType []*string `type:"list" json:",omitempty"`
}

// String returns the string representation
func (s NotifyConfigForDescribeAlertStrategyOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s NotifyConfigForDescribeAlertStrategyOutput) GoString() string {
	return s.String()
}

// SetIgnoreDisabledDomains sets the IgnoreDisabledDomains field's value.
func (s *NotifyConfigForDescribeAlertStrategyOutput) SetIgnoreDisabledDomains(v bool) *NotifyConfigForDescribeAlertStrategyOutput {
	s.IgnoreDisabledDomains = &v
	return s
}

// SetLevel sets the Level field's value.
func (s *NotifyConfigForDescribeAlertStrategyOutput) SetLevel(v string) *NotifyConfigForDescribeAlertStrategyOutput {
	s.Level = &v
	return s
}

// SetSendType sets the SendType field's value.
func (s *NotifyConfigForDescribeAlertStrategyOutput) SetSendType(v []*string) *NotifyConfigForDescribeAlertStrategyOutput {
	s.SendType = v
	return s
}

type ProbeTaskForDescribeAlertStrategyOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Name *string `type:"string" json:",omitempty"`

	TaskId *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s ProbeTaskForDescribeAlertStrategyOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ProbeTaskForDescribeAlertStrategyOutput) GoString() string {
	return s.String()
}

// SetName sets the Name field's value.
func (s *ProbeTaskForDescribeAlertStrategyOutput) SetName(v string) *ProbeTaskForDescribeAlertStrategyOutput {
	s.Name = &v
	return s
}

// SetTaskId sets the TaskId field's value.
func (s *ProbeTaskForDescribeAlertStrategyOutput) SetTaskId(v string) *ProbeTaskForDescribeAlertStrategyOutput {
	s.TaskId = &v
	return s
}

type SubscribeRuleForDescribeAlertStrategyOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	ContactGroups []*ContactGroupForDescribeAlertStrategyOutput `type:"list" json:",omitempty"`

	ContactRobots []*ContactRobotForDescribeAlertStrategyOutput `type:"list" json:",omitempty"`

	NotifyConfig *NotifyConfigForDescribeAlertStrategyOutput `type:"structure" json:",omitempty"`
}

// String returns the string representation
func (s SubscribeRuleForDescribeAlertStrategyOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s SubscribeRuleForDescribeAlertStrategyOutput) GoString() string {
	return s.String()
}

// SetContactGroups sets the ContactGroups field's value.
func (s *SubscribeRuleForDescribeAlertStrategyOutput) SetContactGroups(v []*ContactGroupForDescribeAlertStrategyOutput) *SubscribeRuleForDescribeAlertStrategyOutput {
	s.ContactGroups = v
	return s
}

// SetContactRobots sets the ContactRobots field's value.
func (s *SubscribeRuleForDescribeAlertStrategyOutput) SetContactRobots(v []*ContactRobotForDescribeAlertStrategyOutput) *SubscribeRuleForDescribeAlertStrategyOutput {
	s.ContactRobots = v
	return s
}

// SetNotifyConfig sets the NotifyConfig field's value.
func (s *SubscribeRuleForDescribeAlertStrategyOutput) SetNotifyConfig(v *NotifyConfigForDescribeAlertStrategyOutput) *SubscribeRuleForDescribeAlertStrategyOutput {
	s.NotifyConfig = v
	return s
}

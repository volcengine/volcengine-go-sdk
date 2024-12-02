// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package volcobserve

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opUpdateEventRuleCommon = "UpdateEventRule"

// UpdateEventRuleCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the UpdateEventRuleCommon operation. The "output" return
// value will be populated with the UpdateEventRuleCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned UpdateEventRuleCommon Request to send the API call to the service.
// the "output" return value is not valid until after UpdateEventRuleCommon Send returns without error.
//
// See UpdateEventRuleCommon for more information on using the UpdateEventRuleCommon
// API call, and error handling.
//
//    // Example sending a request using the UpdateEventRuleCommonRequest method.
//    req, resp := client.UpdateEventRuleCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VOLCOBSERVE) UpdateEventRuleCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opUpdateEventRuleCommon,
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

// UpdateEventRuleCommon API operation for VOLC_OBSERVE.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VOLC_OBSERVE's
// API operation UpdateEventRuleCommon for usage and error information.
func (c *VOLCOBSERVE) UpdateEventRuleCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.UpdateEventRuleCommonRequest(input)
	return out, req.Send()
}

// UpdateEventRuleCommonWithContext is the same as UpdateEventRuleCommon with the addition of
// the ability to pass a context and additional request options.
//
// See UpdateEventRuleCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VOLCOBSERVE) UpdateEventRuleCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.UpdateEventRuleCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opUpdateEventRule = "UpdateEventRule"

// UpdateEventRuleRequest generates a "volcengine/request.Request" representing the
// client's request for the UpdateEventRule operation. The "output" return
// value will be populated with the UpdateEventRuleCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned UpdateEventRuleCommon Request to send the API call to the service.
// the "output" return value is not valid until after UpdateEventRuleCommon Send returns without error.
//
// See UpdateEventRule for more information on using the UpdateEventRule
// API call, and error handling.
//
//    // Example sending a request using the UpdateEventRuleRequest method.
//    req, resp := client.UpdateEventRuleRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VOLCOBSERVE) UpdateEventRuleRequest(input *UpdateEventRuleInput) (req *request.Request, output *UpdateEventRuleOutput) {
	op := &request.Operation{
		Name:       opUpdateEventRule,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateEventRuleInput{}
	}

	output = &UpdateEventRuleOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// UpdateEventRule API operation for VOLC_OBSERVE.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VOLC_OBSERVE's
// API operation UpdateEventRule for usage and error information.
func (c *VOLCOBSERVE) UpdateEventRule(input *UpdateEventRuleInput) (*UpdateEventRuleOutput, error) {
	req, out := c.UpdateEventRuleRequest(input)
	return out, req.Send()
}

// UpdateEventRuleWithContext is the same as UpdateEventRule with the addition of
// the ability to pass a context and additional request options.
//
// See UpdateEventRule for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VOLCOBSERVE) UpdateEventRuleWithContext(ctx volcengine.Context, input *UpdateEventRuleInput, opts ...request.Option) (*UpdateEventRuleOutput, error) {
	req, out := c.UpdateEventRuleRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DataForUpdateEventRuleOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	RuleId *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s DataForUpdateEventRuleOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DataForUpdateEventRuleOutput) GoString() string {
	return s.String()
}

// SetRuleId sets the RuleId field's value.
func (s *DataForUpdateEventRuleOutput) SetRuleId(v string) *DataForUpdateEventRuleOutput {
	s.RuleId = &v
	return s
}

type EffectiveTimeForUpdateEventRuleInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	EndTime *string `type:"string" json:",omitempty"`

	StartTime *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s EffectiveTimeForUpdateEventRuleInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s EffectiveTimeForUpdateEventRuleInput) GoString() string {
	return s.String()
}

// SetEndTime sets the EndTime field's value.
func (s *EffectiveTimeForUpdateEventRuleInput) SetEndTime(v string) *EffectiveTimeForUpdateEventRuleInput {
	s.EndTime = &v
	return s
}

// SetStartTime sets the StartTime field's value.
func (s *EffectiveTimeForUpdateEventRuleInput) SetStartTime(v string) *EffectiveTimeForUpdateEventRuleInput {
	s.StartTime = &v
	return s
}

type MessageQueueForUpdateEventRuleInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	AuthEncrypt []*int64 `type:"list" json:",omitempty"`

	Endpoints *string `type:"string" json:",omitempty"`

	InstanceId *string `type:"string" json:",omitempty"`

	Password *string `type:"string" json:",omitempty"`

	Region *string `type:"string" json:",omitempty"`

	Topic *string `type:"string" json:",omitempty"`

	Type *string `type:"string" json:",omitempty"`

	Username *string `type:"string" json:",omitempty"`

	VpcId *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s MessageQueueForUpdateEventRuleInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s MessageQueueForUpdateEventRuleInput) GoString() string {
	return s.String()
}

// SetAuthEncrypt sets the AuthEncrypt field's value.
func (s *MessageQueueForUpdateEventRuleInput) SetAuthEncrypt(v []*int64) *MessageQueueForUpdateEventRuleInput {
	s.AuthEncrypt = v
	return s
}

// SetEndpoints sets the Endpoints field's value.
func (s *MessageQueueForUpdateEventRuleInput) SetEndpoints(v string) *MessageQueueForUpdateEventRuleInput {
	s.Endpoints = &v
	return s
}

// SetInstanceId sets the InstanceId field's value.
func (s *MessageQueueForUpdateEventRuleInput) SetInstanceId(v string) *MessageQueueForUpdateEventRuleInput {
	s.InstanceId = &v
	return s
}

// SetPassword sets the Password field's value.
func (s *MessageQueueForUpdateEventRuleInput) SetPassword(v string) *MessageQueueForUpdateEventRuleInput {
	s.Password = &v
	return s
}

// SetRegion sets the Region field's value.
func (s *MessageQueueForUpdateEventRuleInput) SetRegion(v string) *MessageQueueForUpdateEventRuleInput {
	s.Region = &v
	return s
}

// SetTopic sets the Topic field's value.
func (s *MessageQueueForUpdateEventRuleInput) SetTopic(v string) *MessageQueueForUpdateEventRuleInput {
	s.Topic = &v
	return s
}

// SetType sets the Type field's value.
func (s *MessageQueueForUpdateEventRuleInput) SetType(v string) *MessageQueueForUpdateEventRuleInput {
	s.Type = &v
	return s
}

// SetUsername sets the Username field's value.
func (s *MessageQueueForUpdateEventRuleInput) SetUsername(v string) *MessageQueueForUpdateEventRuleInput {
	s.Username = &v
	return s
}

// SetVpcId sets the VpcId field's value.
func (s *MessageQueueForUpdateEventRuleInput) SetVpcId(v string) *MessageQueueForUpdateEventRuleInput {
	s.VpcId = &v
	return s
}

type TLSTargetForUpdateEventRuleInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	ProjectId *string `type:"string" json:",omitempty"`

	ProjectName *string `type:"string" json:",omitempty"`

	RegionNameCN *string `type:"string" json:",omitempty"`

	RegionNameEN *string `type:"string" json:",omitempty"`

	TopicId *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s TLSTargetForUpdateEventRuleInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s TLSTargetForUpdateEventRuleInput) GoString() string {
	return s.String()
}

// SetProjectId sets the ProjectId field's value.
func (s *TLSTargetForUpdateEventRuleInput) SetProjectId(v string) *TLSTargetForUpdateEventRuleInput {
	s.ProjectId = &v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *TLSTargetForUpdateEventRuleInput) SetProjectName(v string) *TLSTargetForUpdateEventRuleInput {
	s.ProjectName = &v
	return s
}

// SetRegionNameCN sets the RegionNameCN field's value.
func (s *TLSTargetForUpdateEventRuleInput) SetRegionNameCN(v string) *TLSTargetForUpdateEventRuleInput {
	s.RegionNameCN = &v
	return s
}

// SetRegionNameEN sets the RegionNameEN field's value.
func (s *TLSTargetForUpdateEventRuleInput) SetRegionNameEN(v string) *TLSTargetForUpdateEventRuleInput {
	s.RegionNameEN = &v
	return s
}

// SetTopicId sets the TopicId field's value.
func (s *TLSTargetForUpdateEventRuleInput) SetTopicId(v string) *TLSTargetForUpdateEventRuleInput {
	s.TopicId = &v
	return s
}

type UpdateEventRuleInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	ContactGroupIds []*string `type:"list" json:",omitempty"`

	ContactMethods []*string `type:"list" json:",omitempty"`

	Description *string `type:"string" json:",omitempty"`

	EffectiveTime *EffectiveTimeForUpdateEventRuleInput `type:"structure" json:",omitempty"`

	Endpoint *string `type:"string" json:",omitempty"`

	// EventBusName is a required field
	EventBusName *string `min:"2" max:"127" type:"string" json:",omitempty" required:"true" enum:"EnumOfEventBusNameForUpdateEventRuleInput"`

	// EventSource is a required field
	EventSource *string `type:"string" json:",omitempty" required:"true"`

	EventType []*string `type:"list" json:",omitempty"`

	FilterPattern map[string]*interface{} `type:"map" json:",omitempty"`

	// Level is a required field
	Level *string `type:"string" json:",omitempty" required:"true"`

	MessageQueue []*MessageQueueForUpdateEventRuleInput `type:"list" json:",omitempty"`

	// RuleId is a required field
	RuleId *string `type:"string" json:",omitempty" required:"true"`

	// RuleName is a required field
	RuleName *string `min:"2" max:"127" type:"string" json:",omitempty" required:"true"`

	Status *string `type:"string" json:",omitempty" enum:"EnumOfStatusForUpdateEventRuleInput"`

	TLSTarget []*TLSTargetForUpdateEventRuleInput `type:"list" json:",omitempty"`

	WebhookIds []*string `type:"list" json:",omitempty"`
}

// String returns the string representation
func (s UpdateEventRuleInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s UpdateEventRuleInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateEventRuleInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "UpdateEventRuleInput"}
	if s.EventBusName == nil {
		invalidParams.Add(request.NewErrParamRequired("EventBusName"))
	}
	if s.EventBusName != nil && len(*s.EventBusName) < 2 {
		invalidParams.Add(request.NewErrParamMinLen("EventBusName", 2))
	}
	if s.EventBusName != nil && len(*s.EventBusName) > 127 {
		invalidParams.Add(request.NewErrParamMaxLen("EventBusName", 127, *s.EventBusName))
	}
	if s.EventSource == nil {
		invalidParams.Add(request.NewErrParamRequired("EventSource"))
	}
	if s.Level == nil {
		invalidParams.Add(request.NewErrParamRequired("Level"))
	}
	if s.RuleId == nil {
		invalidParams.Add(request.NewErrParamRequired("RuleId"))
	}
	if s.RuleName == nil {
		invalidParams.Add(request.NewErrParamRequired("RuleName"))
	}
	if s.RuleName != nil && len(*s.RuleName) < 2 {
		invalidParams.Add(request.NewErrParamMinLen("RuleName", 2))
	}
	if s.RuleName != nil && len(*s.RuleName) > 127 {
		invalidParams.Add(request.NewErrParamMaxLen("RuleName", 127, *s.RuleName))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetContactGroupIds sets the ContactGroupIds field's value.
func (s *UpdateEventRuleInput) SetContactGroupIds(v []*string) *UpdateEventRuleInput {
	s.ContactGroupIds = v
	return s
}

// SetContactMethods sets the ContactMethods field's value.
func (s *UpdateEventRuleInput) SetContactMethods(v []*string) *UpdateEventRuleInput {
	s.ContactMethods = v
	return s
}

// SetDescription sets the Description field's value.
func (s *UpdateEventRuleInput) SetDescription(v string) *UpdateEventRuleInput {
	s.Description = &v
	return s
}

// SetEffectiveTime sets the EffectiveTime field's value.
func (s *UpdateEventRuleInput) SetEffectiveTime(v *EffectiveTimeForUpdateEventRuleInput) *UpdateEventRuleInput {
	s.EffectiveTime = v
	return s
}

// SetEndpoint sets the Endpoint field's value.
func (s *UpdateEventRuleInput) SetEndpoint(v string) *UpdateEventRuleInput {
	s.Endpoint = &v
	return s
}

// SetEventBusName sets the EventBusName field's value.
func (s *UpdateEventRuleInput) SetEventBusName(v string) *UpdateEventRuleInput {
	s.EventBusName = &v
	return s
}

// SetEventSource sets the EventSource field's value.
func (s *UpdateEventRuleInput) SetEventSource(v string) *UpdateEventRuleInput {
	s.EventSource = &v
	return s
}

// SetEventType sets the EventType field's value.
func (s *UpdateEventRuleInput) SetEventType(v []*string) *UpdateEventRuleInput {
	s.EventType = v
	return s
}

// SetFilterPattern sets the FilterPattern field's value.
func (s *UpdateEventRuleInput) SetFilterPattern(v map[string]*interface{}) *UpdateEventRuleInput {
	s.FilterPattern = v
	return s
}

// SetLevel sets the Level field's value.
func (s *UpdateEventRuleInput) SetLevel(v string) *UpdateEventRuleInput {
	s.Level = &v
	return s
}

// SetMessageQueue sets the MessageQueue field's value.
func (s *UpdateEventRuleInput) SetMessageQueue(v []*MessageQueueForUpdateEventRuleInput) *UpdateEventRuleInput {
	s.MessageQueue = v
	return s
}

// SetRuleId sets the RuleId field's value.
func (s *UpdateEventRuleInput) SetRuleId(v string) *UpdateEventRuleInput {
	s.RuleId = &v
	return s
}

// SetRuleName sets the RuleName field's value.
func (s *UpdateEventRuleInput) SetRuleName(v string) *UpdateEventRuleInput {
	s.RuleName = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *UpdateEventRuleInput) SetStatus(v string) *UpdateEventRuleInput {
	s.Status = &v
	return s
}

// SetTLSTarget sets the TLSTarget field's value.
func (s *UpdateEventRuleInput) SetTLSTarget(v []*TLSTargetForUpdateEventRuleInput) *UpdateEventRuleInput {
	s.TLSTarget = v
	return s
}

// SetWebhookIds sets the WebhookIds field's value.
func (s *UpdateEventRuleInput) SetWebhookIds(v []*string) *UpdateEventRuleInput {
	s.WebhookIds = v
	return s
}

type UpdateEventRuleOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	Data *DataForUpdateEventRuleOutput `type:"structure" json:",omitempty"`
}

// String returns the string representation
func (s UpdateEventRuleOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s UpdateEventRuleOutput) GoString() string {
	return s.String()
}

// SetData sets the Data field's value.
func (s *UpdateEventRuleOutput) SetData(v *DataForUpdateEventRuleOutput) *UpdateEventRuleOutput {
	s.Data = v
	return s
}

const (
	// EnumOfEventBusNameForUpdateEventRuleInputDefault is a EnumOfEventBusNameForUpdateEventRuleInput enum value
	EnumOfEventBusNameForUpdateEventRuleInputDefault = "default"
)

const (
	// EnumOfStatusForUpdateEventRuleInputEnable is a EnumOfStatusForUpdateEventRuleInput enum value
	EnumOfStatusForUpdateEventRuleInputEnable = "enable"

	// EnumOfStatusForUpdateEventRuleInputDisable is a EnumOfStatusForUpdateEventRuleInput enum value
	EnumOfStatusForUpdateEventRuleInputDisable = "disable"
)

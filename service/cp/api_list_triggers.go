// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package cp

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opListTriggersCommon = "ListTriggers"

// ListTriggersCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ListTriggersCommon operation. The "output" return
// value will be populated with the ListTriggersCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListTriggersCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListTriggersCommon Send returns without error.
//
// See ListTriggersCommon for more information on using the ListTriggersCommon
// API call, and error handling.
//
//    // Example sending a request using the ListTriggersCommonRequest method.
//    req, resp := client.ListTriggersCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CP) ListTriggersCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opListTriggersCommon,
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

// ListTriggersCommon API operation for CP.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CP's
// API operation ListTriggersCommon for usage and error information.
func (c *CP) ListTriggersCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ListTriggersCommonRequest(input)
	return out, req.Send()
}

// ListTriggersCommonWithContext is the same as ListTriggersCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ListTriggersCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CP) ListTriggersCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ListTriggersCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opListTriggers = "ListTriggers"

// ListTriggersRequest generates a "volcengine/request.Request" representing the
// client's request for the ListTriggers operation. The "output" return
// value will be populated with the ListTriggersCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListTriggersCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListTriggersCommon Send returns without error.
//
// See ListTriggers for more information on using the ListTriggers
// API call, and error handling.
//
//    // Example sending a request using the ListTriggersRequest method.
//    req, resp := client.ListTriggersRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CP) ListTriggersRequest(input *ListTriggersInput) (req *request.Request, output *ListTriggersOutput) {
	op := &request.Operation{
		Name:       opListTriggers,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListTriggersInput{}
	}

	output = &ListTriggersOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// ListTriggers API operation for CP.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CP's
// API operation ListTriggers for usage and error information.
func (c *CP) ListTriggers(input *ListTriggersInput) (*ListTriggersOutput, error) {
	req, out := c.ListTriggersRequest(input)
	return out, req.Send()
}

// ListTriggersWithContext is the same as ListTriggers with the addition of
// the ability to pass a context and additional request options.
//
// See ListTriggers for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CP) ListTriggersWithContext(ctx volcengine.Context, input *ListTriggersInput, opts ...request.Option) (*ListTriggersOutput, error) {
	req, out := c.ListTriggersRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type CRForListTriggersOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Filter *FilterForListTriggersOutput `type:"structure" json:",omitempty"`
}

// String returns the string representation
func (s CRForListTriggersOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CRForListTriggersOutput) GoString() string {
	return s.String()
}

// SetFilter sets the Filter field's value.
func (s *CRForListTriggersOutput) SetFilter(v *FilterForListTriggersOutput) *CRForListTriggersOutput {
	s.Filter = v
	return s
}

type ConfigForListTriggersOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	References []*string `type:"list" json:",omitempty"`
}

// String returns the string representation
func (s ConfigForListTriggersOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ConfigForListTriggersOutput) GoString() string {
	return s.String()
}

// SetReferences sets the References field's value.
func (s *ConfigForListTriggersOutput) SetReferences(v []*string) *ConfigForListTriggersOutput {
	s.References = v
	return s
}

type ConfigurationForListTriggersOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Schedule *ScheduleForListTriggersOutput `type:"structure" json:",omitempty"`

	Webhook *WebhookForListTriggersOutput `type:"structure" json:",omitempty"`
}

// String returns the string representation
func (s ConfigurationForListTriggersOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ConfigurationForListTriggersOutput) GoString() string {
	return s.String()
}

// SetSchedule sets the Schedule field's value.
func (s *ConfigurationForListTriggersOutput) SetSchedule(v *ScheduleForListTriggersOutput) *ConfigurationForListTriggersOutput {
	s.Schedule = v
	return s
}

// SetWebhook sets the Webhook field's value.
func (s *ConfigurationForListTriggersOutput) SetWebhook(v *WebhookForListTriggersOutput) *ConfigurationForListTriggersOutput {
	s.Webhook = v
	return s
}

type ConvertFilterForListTriggersOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Config *ConfigForListTriggersOutput `type:"structure" json:",omitempty"`

	EventType *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s ConvertFilterForListTriggersOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ConvertFilterForListTriggersOutput) GoString() string {
	return s.String()
}

// SetConfig sets the Config field's value.
func (s *ConvertFilterForListTriggersOutput) SetConfig(v *ConfigForListTriggersOutput) *ConvertFilterForListTriggersOutput {
	s.Config = v
	return s
}

// SetEventType sets the EventType field's value.
func (s *ConvertFilterForListTriggersOutput) SetEventType(v string) *ConvertFilterForListTriggersOutput {
	s.EventType = &v
	return s
}

type CreatorForListTriggersOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	AccountId *int64 `type:"int64" json:",omitempty"`

	UserId *int64 `type:"int64" json:",omitempty"`
}

// String returns the string representation
func (s CreatorForListTriggersOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreatorForListTriggersOutput) GoString() string {
	return s.String()
}

// SetAccountId sets the AccountId field's value.
func (s *CreatorForListTriggersOutput) SetAccountId(v int64) *CreatorForListTriggersOutput {
	s.AccountId = &v
	return s
}

// SetUserId sets the UserId field's value.
func (s *CreatorForListTriggersOutput) SetUserId(v int64) *CreatorForListTriggersOutput {
	s.UserId = &v
	return s
}

type FilterForListTriggersInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Ids []*string `type:"list" json:",omitempty"`

	Name *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s FilterForListTriggersInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s FilterForListTriggersInput) GoString() string {
	return s.String()
}

// SetIds sets the Ids field's value.
func (s *FilterForListTriggersInput) SetIds(v []*string) *FilterForListTriggersInput {
	s.Ids = v
	return s
}

// SetName sets the Name field's value.
func (s *FilterForListTriggersInput) SetName(v string) *FilterForListTriggersInput {
	s.Name = &v
	return s
}

type FilterForListTriggersOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	CREventType *string `type:"string" json:",omitempty"`

	Condition *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s FilterForListTriggersOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s FilterForListTriggersOutput) GoString() string {
	return s.String()
}

// SetCREventType sets the CREventType field's value.
func (s *FilterForListTriggersOutput) SetCREventType(v string) *FilterForListTriggersOutput {
	s.CREventType = &v
	return s
}

// SetCondition sets the Condition field's value.
func (s *FilterForListTriggersOutput) SetCondition(v string) *FilterForListTriggersOutput {
	s.Condition = &v
	return s
}

type GitForListTriggersOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Filters []*ConvertFilterForListTriggersOutput `type:"list" json:",omitempty"`

	ResourceId *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s GitForListTriggersOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s GitForListTriggersOutput) GoString() string {
	return s.String()
}

// SetFilters sets the Filters field's value.
func (s *GitForListTriggersOutput) SetFilters(v []*ConvertFilterForListTriggersOutput) *GitForListTriggersOutput {
	s.Filters = v
	return s
}

// SetResourceId sets the ResourceId field's value.
func (s *GitForListTriggersOutput) SetResourceId(v string) *GitForListTriggersOutput {
	s.ResourceId = &v
	return s
}

type ItemForListTriggersOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Configuration *ConfigurationForListTriggersOutput `type:"structure" json:",omitempty"`

	CreateTime *string `type:"string" json:",omitempty"`

	Creator *CreatorForListTriggersOutput `type:"structure" json:",omitempty"`

	Description *string `type:"string" json:",omitempty"`

	Id *string `type:"string" json:",omitempty"`

	Name *string `type:"string" json:",omitempty"`

	Parameters []*ParameterForListTriggersOutput `type:"list" json:",omitempty"`

	PipelineId *string `type:"string" json:",omitempty"`

	Resources []*ResourceForListTriggersOutput `type:"list" json:",omitempty"`

	Type *string `type:"string" json:",omitempty" enum:"EnumOfTypeForListTriggersOutput"`

	UpdateTime *string `type:"string" json:",omitempty"`

	WorkspaceId *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s ItemForListTriggersOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ItemForListTriggersOutput) GoString() string {
	return s.String()
}

// SetConfiguration sets the Configuration field's value.
func (s *ItemForListTriggersOutput) SetConfiguration(v *ConfigurationForListTriggersOutput) *ItemForListTriggersOutput {
	s.Configuration = v
	return s
}

// SetCreateTime sets the CreateTime field's value.
func (s *ItemForListTriggersOutput) SetCreateTime(v string) *ItemForListTriggersOutput {
	s.CreateTime = &v
	return s
}

// SetCreator sets the Creator field's value.
func (s *ItemForListTriggersOutput) SetCreator(v *CreatorForListTriggersOutput) *ItemForListTriggersOutput {
	s.Creator = v
	return s
}

// SetDescription sets the Description field's value.
func (s *ItemForListTriggersOutput) SetDescription(v string) *ItemForListTriggersOutput {
	s.Description = &v
	return s
}

// SetId sets the Id field's value.
func (s *ItemForListTriggersOutput) SetId(v string) *ItemForListTriggersOutput {
	s.Id = &v
	return s
}

// SetName sets the Name field's value.
func (s *ItemForListTriggersOutput) SetName(v string) *ItemForListTriggersOutput {
	s.Name = &v
	return s
}

// SetParameters sets the Parameters field's value.
func (s *ItemForListTriggersOutput) SetParameters(v []*ParameterForListTriggersOutput) *ItemForListTriggersOutput {
	s.Parameters = v
	return s
}

// SetPipelineId sets the PipelineId field's value.
func (s *ItemForListTriggersOutput) SetPipelineId(v string) *ItemForListTriggersOutput {
	s.PipelineId = &v
	return s
}

// SetResources sets the Resources field's value.
func (s *ItemForListTriggersOutput) SetResources(v []*ResourceForListTriggersOutput) *ItemForListTriggersOutput {
	s.Resources = v
	return s
}

// SetType sets the Type field's value.
func (s *ItemForListTriggersOutput) SetType(v string) *ItemForListTriggersOutput {
	s.Type = &v
	return s
}

// SetUpdateTime sets the UpdateTime field's value.
func (s *ItemForListTriggersOutput) SetUpdateTime(v string) *ItemForListTriggersOutput {
	s.UpdateTime = &v
	return s
}

// SetWorkspaceId sets the WorkspaceId field's value.
func (s *ItemForListTriggersOutput) SetWorkspaceId(v string) *ItemForListTriggersOutput {
	s.WorkspaceId = &v
	return s
}

type ListTriggersInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Filter *FilterForListTriggersInput `type:"structure" json:",omitempty"`

	PageNumber *int64 `type:"int64" json:",omitempty"`

	PageSize *int64 `type:"int64" json:",omitempty"`

	// PipelineId is a required field
	PipelineId *string `type:"string" json:",omitempty" required:"true"`

	// WorkspaceId is a required field
	WorkspaceId *string `type:"string" json:",omitempty" required:"true"`
}

// String returns the string representation
func (s ListTriggersInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListTriggersInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListTriggersInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ListTriggersInput"}
	if s.PipelineId == nil {
		invalidParams.Add(request.NewErrParamRequired("PipelineId"))
	}
	if s.WorkspaceId == nil {
		invalidParams.Add(request.NewErrParamRequired("WorkspaceId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetFilter sets the Filter field's value.
func (s *ListTriggersInput) SetFilter(v *FilterForListTriggersInput) *ListTriggersInput {
	s.Filter = v
	return s
}

// SetPageNumber sets the PageNumber field's value.
func (s *ListTriggersInput) SetPageNumber(v int64) *ListTriggersInput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *ListTriggersInput) SetPageSize(v int64) *ListTriggersInput {
	s.PageSize = &v
	return s
}

// SetPipelineId sets the PipelineId field's value.
func (s *ListTriggersInput) SetPipelineId(v string) *ListTriggersInput {
	s.PipelineId = &v
	return s
}

// SetWorkspaceId sets the WorkspaceId field's value.
func (s *ListTriggersInput) SetWorkspaceId(v string) *ListTriggersInput {
	s.WorkspaceId = &v
	return s
}

type ListTriggersOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	Items []*ItemForListTriggersOutput `type:"list" json:",omitempty"`

	PageNumber *int64 `type:"int64" json:",omitempty"`

	PageSize *int64 `type:"int64" json:",omitempty"`

	TotalCount *int64 `type:"int64" json:",omitempty"`
}

// String returns the string representation
func (s ListTriggersOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListTriggersOutput) GoString() string {
	return s.String()
}

// SetItems sets the Items field's value.
func (s *ListTriggersOutput) SetItems(v []*ItemForListTriggersOutput) *ListTriggersOutput {
	s.Items = v
	return s
}

// SetPageNumber sets the PageNumber field's value.
func (s *ListTriggersOutput) SetPageNumber(v int64) *ListTriggersOutput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *ListTriggersOutput) SetPageSize(v int64) *ListTriggersOutput {
	s.PageSize = &v
	return s
}

// SetTotalCount sets the TotalCount field's value.
func (s *ListTriggersOutput) SetTotalCount(v int64) *ListTriggersOutput {
	s.TotalCount = &v
	return s
}

type ParameterForListTriggersOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Key *string `type:"string" json:",omitempty"`

	Value *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s ParameterForListTriggersOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ParameterForListTriggersOutput) GoString() string {
	return s.String()
}

// SetKey sets the Key field's value.
func (s *ParameterForListTriggersOutput) SetKey(v string) *ParameterForListTriggersOutput {
	s.Key = &v
	return s
}

// SetValue sets the Value field's value.
func (s *ParameterForListTriggersOutput) SetValue(v string) *ParameterForListTriggersOutput {
	s.Value = &v
	return s
}

type ResourceForListTriggersOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Reference *string `type:"string" json:",omitempty"`

	ResourceId *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s ResourceForListTriggersOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ResourceForListTriggersOutput) GoString() string {
	return s.String()
}

// SetReference sets the Reference field's value.
func (s *ResourceForListTriggersOutput) SetReference(v string) *ResourceForListTriggersOutput {
	s.Reference = &v
	return s
}

// SetResourceId sets the ResourceId field's value.
func (s *ResourceForListTriggersOutput) SetResourceId(v string) *ResourceForListTriggersOutput {
	s.ResourceId = &v
	return s
}

type ScheduleForListTriggersOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	ScheduleConfig *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s ScheduleForListTriggersOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ScheduleForListTriggersOutput) GoString() string {
	return s.String()
}

// SetScheduleConfig sets the ScheduleConfig field's value.
func (s *ScheduleForListTriggersOutput) SetScheduleConfig(v string) *ScheduleForListTriggersOutput {
	s.ScheduleConfig = &v
	return s
}

type WebhookForListTriggersOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	CR *CRForListTriggersOutput `type:"structure" json:",omitempty"`

	Git *GitForListTriggersOutput `type:"structure" json:",omitempty"`

	URL *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s WebhookForListTriggersOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s WebhookForListTriggersOutput) GoString() string {
	return s.String()
}

// SetCR sets the CR field's value.
func (s *WebhookForListTriggersOutput) SetCR(v *CRForListTriggersOutput) *WebhookForListTriggersOutput {
	s.CR = v
	return s
}

// SetGit sets the Git field's value.
func (s *WebhookForListTriggersOutput) SetGit(v *GitForListTriggersOutput) *WebhookForListTriggersOutput {
	s.Git = v
	return s
}

// SetURL sets the URL field's value.
func (s *WebhookForListTriggersOutput) SetURL(v string) *WebhookForListTriggersOutput {
	s.URL = &v
	return s
}

const (
	// EnumOfTypeForListTriggersOutputWebhook is a EnumOfTypeForListTriggersOutput enum value
	EnumOfTypeForListTriggersOutputWebhook = "Webhook"

	// EnumOfTypeForListTriggersOutputGitWebhook is a EnumOfTypeForListTriggersOutput enum value
	EnumOfTypeForListTriggersOutputGitWebhook = "GitWebhook"

	// EnumOfTypeForListTriggersOutputPerforceWebhook is a EnumOfTypeForListTriggersOutput enum value
	EnumOfTypeForListTriggersOutputPerforceWebhook = "PerforceWebhook"

	// EnumOfTypeForListTriggersOutputCrwebhook is a EnumOfTypeForListTriggersOutput enum value
	EnumOfTypeForListTriggersOutputCrwebhook = "CRWebhook"

	// EnumOfTypeForListTriggersOutputSchedule is a EnumOfTypeForListTriggersOutput enum value
	EnumOfTypeForListTriggersOutputSchedule = "Schedule"

	// EnumOfTypeForListTriggersOutputManual is a EnumOfTypeForListTriggersOutput enum value
	EnumOfTypeForListTriggersOutputManual = "Manual"
)

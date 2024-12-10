// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package cp

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opListPipelineRunsCommon = "ListPipelineRuns"

// ListPipelineRunsCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ListPipelineRunsCommon operation. The "output" return
// value will be populated with the ListPipelineRunsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListPipelineRunsCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListPipelineRunsCommon Send returns without error.
//
// See ListPipelineRunsCommon for more information on using the ListPipelineRunsCommon
// API call, and error handling.
//
//    // Example sending a request using the ListPipelineRunsCommonRequest method.
//    req, resp := client.ListPipelineRunsCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CP) ListPipelineRunsCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opListPipelineRunsCommon,
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

// ListPipelineRunsCommon API operation for CP.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CP's
// API operation ListPipelineRunsCommon for usage and error information.
func (c *CP) ListPipelineRunsCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ListPipelineRunsCommonRequest(input)
	return out, req.Send()
}

// ListPipelineRunsCommonWithContext is the same as ListPipelineRunsCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ListPipelineRunsCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CP) ListPipelineRunsCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ListPipelineRunsCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opListPipelineRuns = "ListPipelineRuns"

// ListPipelineRunsRequest generates a "volcengine/request.Request" representing the
// client's request for the ListPipelineRuns operation. The "output" return
// value will be populated with the ListPipelineRunsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListPipelineRunsCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListPipelineRunsCommon Send returns without error.
//
// See ListPipelineRuns for more information on using the ListPipelineRuns
// API call, and error handling.
//
//    // Example sending a request using the ListPipelineRunsRequest method.
//    req, resp := client.ListPipelineRunsRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CP) ListPipelineRunsRequest(input *ListPipelineRunsInput) (req *request.Request, output *ListPipelineRunsOutput) {
	op := &request.Operation{
		Name:       opListPipelineRuns,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListPipelineRunsInput{}
	}

	output = &ListPipelineRunsOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// ListPipelineRuns API operation for CP.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CP's
// API operation ListPipelineRuns for usage and error information.
func (c *CP) ListPipelineRuns(input *ListPipelineRunsInput) (*ListPipelineRunsOutput, error) {
	req, out := c.ListPipelineRunsRequest(input)
	return out, req.Send()
}

// ListPipelineRunsWithContext is the same as ListPipelineRuns with the addition of
// the ability to pass a context and additional request options.
//
// See ListPipelineRuns for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CP) ListPipelineRunsWithContext(ctx volcengine.Context, input *ListPipelineRunsInput, opts ...request.Option) (*ListPipelineRunsOutput, error) {
	req, out := c.ListPipelineRunsRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ConfigurationForListPipelineRunsOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	GitResource *GitResourceForListPipelineRunsOutput `type:"structure" json:",omitempty"`
}

// String returns the string representation
func (s ConfigurationForListPipelineRunsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ConfigurationForListPipelineRunsOutput) GoString() string {
	return s.String()
}

// SetGitResource sets the GitResource field's value.
func (s *ConfigurationForListPipelineRunsOutput) SetGitResource(v *GitResourceForListPipelineRunsOutput) *ConfigurationForListPipelineRunsOutput {
	s.GitResource = v
	return s
}

type FilterForListPipelineRunsInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Ids []*string `type:"list" json:",omitempty"`

	Statuses []*string `type:"list" json:",omitempty"`
}

// String returns the string representation
func (s FilterForListPipelineRunsInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s FilterForListPipelineRunsInput) GoString() string {
	return s.String()
}

// SetIds sets the Ids field's value.
func (s *FilterForListPipelineRunsInput) SetIds(v []*string) *FilterForListPipelineRunsInput {
	s.Ids = v
	return s
}

// SetStatuses sets the Statuses field's value.
func (s *FilterForListPipelineRunsInput) SetStatuses(v []*string) *FilterForListPipelineRunsInput {
	s.Statuses = v
	return s
}

type GitResourceForListPipelineRunsOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	CloneDepth *int32 `type:"int32" json:",omitempty"`

	DefaultBranch *string `type:"string" json:",omitempty"`

	Reference *string `type:"string" json:",omitempty"`

	Type *string `type:"string" json:",omitempty"`

	URL *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s GitResourceForListPipelineRunsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s GitResourceForListPipelineRunsOutput) GoString() string {
	return s.String()
}

// SetCloneDepth sets the CloneDepth field's value.
func (s *GitResourceForListPipelineRunsOutput) SetCloneDepth(v int32) *GitResourceForListPipelineRunsOutput {
	s.CloneDepth = &v
	return s
}

// SetDefaultBranch sets the DefaultBranch field's value.
func (s *GitResourceForListPipelineRunsOutput) SetDefaultBranch(v string) *GitResourceForListPipelineRunsOutput {
	s.DefaultBranch = &v
	return s
}

// SetReference sets the Reference field's value.
func (s *GitResourceForListPipelineRunsOutput) SetReference(v string) *GitResourceForListPipelineRunsOutput {
	s.Reference = &v
	return s
}

// SetType sets the Type field's value.
func (s *GitResourceForListPipelineRunsOutput) SetType(v string) *GitResourceForListPipelineRunsOutput {
	s.Type = &v
	return s
}

// SetURL sets the URL field's value.
func (s *GitResourceForListPipelineRunsOutput) SetURL(v string) *GitResourceForListPipelineRunsOutput {
	s.URL = &v
	return s
}

type ItemForListPipelineRunsOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	CreateTime *string `type:"string" json:",omitempty"`

	Description *string `type:"string" json:",omitempty"`

	FinishTime *string `type:"string" json:",omitempty"`

	Id *string `type:"string" json:",omitempty"`

	Index *int64 `type:"int64" json:",omitempty"`

	Parameters []*ParameterForListPipelineRunsOutput `type:"list" json:",omitempty"`

	PipelineId *string `type:"string" json:",omitempty"`

	Resources []*ResourceForListPipelineRunsOutput `type:"list" json:",omitempty"`

	Spec *string `type:"string" json:",omitempty"`

	Stages []*StageForListPipelineRunsOutput `type:"list" json:",omitempty"`

	StartTime *string `type:"string" json:",omitempty"`

	Status *string `type:"string" json:",omitempty"`

	SystemParameters []*SystemParameterForListPipelineRunsOutput `type:"list" json:",omitempty"`

	Trigger *TriggerForListPipelineRunsOutput `type:"structure" json:",omitempty"`

	UpdateTime *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s ItemForListPipelineRunsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ItemForListPipelineRunsOutput) GoString() string {
	return s.String()
}

// SetCreateTime sets the CreateTime field's value.
func (s *ItemForListPipelineRunsOutput) SetCreateTime(v string) *ItemForListPipelineRunsOutput {
	s.CreateTime = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *ItemForListPipelineRunsOutput) SetDescription(v string) *ItemForListPipelineRunsOutput {
	s.Description = &v
	return s
}

// SetFinishTime sets the FinishTime field's value.
func (s *ItemForListPipelineRunsOutput) SetFinishTime(v string) *ItemForListPipelineRunsOutput {
	s.FinishTime = &v
	return s
}

// SetId sets the Id field's value.
func (s *ItemForListPipelineRunsOutput) SetId(v string) *ItemForListPipelineRunsOutput {
	s.Id = &v
	return s
}

// SetIndex sets the Index field's value.
func (s *ItemForListPipelineRunsOutput) SetIndex(v int64) *ItemForListPipelineRunsOutput {
	s.Index = &v
	return s
}

// SetParameters sets the Parameters field's value.
func (s *ItemForListPipelineRunsOutput) SetParameters(v []*ParameterForListPipelineRunsOutput) *ItemForListPipelineRunsOutput {
	s.Parameters = v
	return s
}

// SetPipelineId sets the PipelineId field's value.
func (s *ItemForListPipelineRunsOutput) SetPipelineId(v string) *ItemForListPipelineRunsOutput {
	s.PipelineId = &v
	return s
}

// SetResources sets the Resources field's value.
func (s *ItemForListPipelineRunsOutput) SetResources(v []*ResourceForListPipelineRunsOutput) *ItemForListPipelineRunsOutput {
	s.Resources = v
	return s
}

// SetSpec sets the Spec field's value.
func (s *ItemForListPipelineRunsOutput) SetSpec(v string) *ItemForListPipelineRunsOutput {
	s.Spec = &v
	return s
}

// SetStages sets the Stages field's value.
func (s *ItemForListPipelineRunsOutput) SetStages(v []*StageForListPipelineRunsOutput) *ItemForListPipelineRunsOutput {
	s.Stages = v
	return s
}

// SetStartTime sets the StartTime field's value.
func (s *ItemForListPipelineRunsOutput) SetStartTime(v string) *ItemForListPipelineRunsOutput {
	s.StartTime = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *ItemForListPipelineRunsOutput) SetStatus(v string) *ItemForListPipelineRunsOutput {
	s.Status = &v
	return s
}

// SetSystemParameters sets the SystemParameters field's value.
func (s *ItemForListPipelineRunsOutput) SetSystemParameters(v []*SystemParameterForListPipelineRunsOutput) *ItemForListPipelineRunsOutput {
	s.SystemParameters = v
	return s
}

// SetTrigger sets the Trigger field's value.
func (s *ItemForListPipelineRunsOutput) SetTrigger(v *TriggerForListPipelineRunsOutput) *ItemForListPipelineRunsOutput {
	s.Trigger = v
	return s
}

// SetUpdateTime sets the UpdateTime field's value.
func (s *ItemForListPipelineRunsOutput) SetUpdateTime(v string) *ItemForListPipelineRunsOutput {
	s.UpdateTime = &v
	return s
}

type ListPipelineRunsInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Filter *FilterForListPipelineRunsInput `type:"structure" json:",omitempty"`

	MaxResults *int64 `type:"int64" json:",omitempty"`

	NextToken *string `type:"string" json:",omitempty"`

	// PipelineId is a required field
	PipelineId *string `type:"string" json:",omitempty" required:"true"`

	// WorkspaceId is a required field
	WorkspaceId *string `type:"string" json:",omitempty" required:"true"`
}

// String returns the string representation
func (s ListPipelineRunsInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListPipelineRunsInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListPipelineRunsInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ListPipelineRunsInput"}
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
func (s *ListPipelineRunsInput) SetFilter(v *FilterForListPipelineRunsInput) *ListPipelineRunsInput {
	s.Filter = v
	return s
}

// SetMaxResults sets the MaxResults field's value.
func (s *ListPipelineRunsInput) SetMaxResults(v int64) *ListPipelineRunsInput {
	s.MaxResults = &v
	return s
}

// SetNextToken sets the NextToken field's value.
func (s *ListPipelineRunsInput) SetNextToken(v string) *ListPipelineRunsInput {
	s.NextToken = &v
	return s
}

// SetPipelineId sets the PipelineId field's value.
func (s *ListPipelineRunsInput) SetPipelineId(v string) *ListPipelineRunsInput {
	s.PipelineId = &v
	return s
}

// SetWorkspaceId sets the WorkspaceId field's value.
func (s *ListPipelineRunsInput) SetWorkspaceId(v string) *ListPipelineRunsInput {
	s.WorkspaceId = &v
	return s
}

type ListPipelineRunsOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	Items []*ItemForListPipelineRunsOutput `type:"list" json:",omitempty"`

	NextToken *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s ListPipelineRunsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListPipelineRunsOutput) GoString() string {
	return s.String()
}

// SetItems sets the Items field's value.
func (s *ListPipelineRunsOutput) SetItems(v []*ItemForListPipelineRunsOutput) *ListPipelineRunsOutput {
	s.Items = v
	return s
}

// SetNextToken sets the NextToken field's value.
func (s *ListPipelineRunsOutput) SetNextToken(v string) *ListPipelineRunsOutput {
	s.NextToken = &v
	return s
}

type ParameterForListPipelineRunsOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Description *string `type:"string" json:",omitempty"`

	Dynamic *bool `type:"boolean" json:",omitempty"`

	Key *string `type:"string" json:",omitempty"`

	OptionValues []*string `type:"list" json:",omitempty"`

	Secret *bool `type:"boolean" json:",omitempty"`

	Value *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s ParameterForListPipelineRunsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ParameterForListPipelineRunsOutput) GoString() string {
	return s.String()
}

// SetDescription sets the Description field's value.
func (s *ParameterForListPipelineRunsOutput) SetDescription(v string) *ParameterForListPipelineRunsOutput {
	s.Description = &v
	return s
}

// SetDynamic sets the Dynamic field's value.
func (s *ParameterForListPipelineRunsOutput) SetDynamic(v bool) *ParameterForListPipelineRunsOutput {
	s.Dynamic = &v
	return s
}

// SetKey sets the Key field's value.
func (s *ParameterForListPipelineRunsOutput) SetKey(v string) *ParameterForListPipelineRunsOutput {
	s.Key = &v
	return s
}

// SetOptionValues sets the OptionValues field's value.
func (s *ParameterForListPipelineRunsOutput) SetOptionValues(v []*string) *ParameterForListPipelineRunsOutput {
	s.OptionValues = v
	return s
}

// SetSecret sets the Secret field's value.
func (s *ParameterForListPipelineRunsOutput) SetSecret(v bool) *ParameterForListPipelineRunsOutput {
	s.Secret = &v
	return s
}

// SetValue sets the Value field's value.
func (s *ParameterForListPipelineRunsOutput) SetValue(v string) *ParameterForListPipelineRunsOutput {
	s.Value = &v
	return s
}

type ResourceForListPipelineRunsOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Configuration *ConfigurationForListPipelineRunsOutput `type:"structure" json:",omitempty"`

	Id *string `type:"string" json:",omitempty"`

	Name *string `type:"string" json:",omitempty"`

	Type *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s ResourceForListPipelineRunsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ResourceForListPipelineRunsOutput) GoString() string {
	return s.String()
}

// SetConfiguration sets the Configuration field's value.
func (s *ResourceForListPipelineRunsOutput) SetConfiguration(v *ConfigurationForListPipelineRunsOutput) *ResourceForListPipelineRunsOutput {
	s.Configuration = v
	return s
}

// SetId sets the Id field's value.
func (s *ResourceForListPipelineRunsOutput) SetId(v string) *ResourceForListPipelineRunsOutput {
	s.Id = &v
	return s
}

// SetName sets the Name field's value.
func (s *ResourceForListPipelineRunsOutput) SetName(v string) *ResourceForListPipelineRunsOutput {
	s.Name = &v
	return s
}

// SetType sets the Type field's value.
func (s *ResourceForListPipelineRunsOutput) SetType(v string) *ResourceForListPipelineRunsOutput {
	s.Type = &v
	return s
}

type StageForListPipelineRunsOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Id *string `type:"string" json:",omitempty"`

	Name *string `type:"string" json:",omitempty"`

	Status *string `type:"string" json:",omitempty"`

	Tasks []*TaskForListPipelineRunsOutput `type:"list" json:",omitempty"`
}

// String returns the string representation
func (s StageForListPipelineRunsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s StageForListPipelineRunsOutput) GoString() string {
	return s.String()
}

// SetId sets the Id field's value.
func (s *StageForListPipelineRunsOutput) SetId(v string) *StageForListPipelineRunsOutput {
	s.Id = &v
	return s
}

// SetName sets the Name field's value.
func (s *StageForListPipelineRunsOutput) SetName(v string) *StageForListPipelineRunsOutput {
	s.Name = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *StageForListPipelineRunsOutput) SetStatus(v string) *StageForListPipelineRunsOutput {
	s.Status = &v
	return s
}

// SetTasks sets the Tasks field's value.
func (s *StageForListPipelineRunsOutput) SetTasks(v []*TaskForListPipelineRunsOutput) *StageForListPipelineRunsOutput {
	s.Tasks = v
	return s
}

type SystemParameterForListPipelineRunsOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Description *string `type:"string" json:",omitempty"`

	Dynamic *bool `type:"boolean" json:",omitempty"`

	Key *string `type:"string" json:",omitempty"`

	OptionValues []*string `type:"list" json:",omitempty"`

	Secret *bool `type:"boolean" json:",omitempty"`

	Value *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s SystemParameterForListPipelineRunsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s SystemParameterForListPipelineRunsOutput) GoString() string {
	return s.String()
}

// SetDescription sets the Description field's value.
func (s *SystemParameterForListPipelineRunsOutput) SetDescription(v string) *SystemParameterForListPipelineRunsOutput {
	s.Description = &v
	return s
}

// SetDynamic sets the Dynamic field's value.
func (s *SystemParameterForListPipelineRunsOutput) SetDynamic(v bool) *SystemParameterForListPipelineRunsOutput {
	s.Dynamic = &v
	return s
}

// SetKey sets the Key field's value.
func (s *SystemParameterForListPipelineRunsOutput) SetKey(v string) *SystemParameterForListPipelineRunsOutput {
	s.Key = &v
	return s
}

// SetOptionValues sets the OptionValues field's value.
func (s *SystemParameterForListPipelineRunsOutput) SetOptionValues(v []*string) *SystemParameterForListPipelineRunsOutput {
	s.OptionValues = v
	return s
}

// SetSecret sets the Secret field's value.
func (s *SystemParameterForListPipelineRunsOutput) SetSecret(v bool) *SystemParameterForListPipelineRunsOutput {
	s.Secret = &v
	return s
}

// SetValue sets the Value field's value.
func (s *SystemParameterForListPipelineRunsOutput) SetValue(v string) *SystemParameterForListPipelineRunsOutput {
	s.Value = &v
	return s
}

type TaskForListPipelineRunsOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	DependsOn []*string `type:"list" json:",omitempty"`

	Id *string `type:"string" json:",omitempty"`

	Name *string `type:"string" json:",omitempty"`

	Status *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s TaskForListPipelineRunsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s TaskForListPipelineRunsOutput) GoString() string {
	return s.String()
}

// SetDependsOn sets the DependsOn field's value.
func (s *TaskForListPipelineRunsOutput) SetDependsOn(v []*string) *TaskForListPipelineRunsOutput {
	s.DependsOn = v
	return s
}

// SetId sets the Id field's value.
func (s *TaskForListPipelineRunsOutput) SetId(v string) *TaskForListPipelineRunsOutput {
	s.Id = &v
	return s
}

// SetName sets the Name field's value.
func (s *TaskForListPipelineRunsOutput) SetName(v string) *TaskForListPipelineRunsOutput {
	s.Name = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *TaskForListPipelineRunsOutput) SetStatus(v string) *TaskForListPipelineRunsOutput {
	s.Status = &v
	return s
}

type TriggerForListPipelineRunsOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	TriggerId *string `type:"string" json:",omitempty"`

	Type *string `type:"string" json:",omitempty"`

	User *UserForListPipelineRunsOutput `type:"structure" json:",omitempty"`
}

// String returns the string representation
func (s TriggerForListPipelineRunsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s TriggerForListPipelineRunsOutput) GoString() string {
	return s.String()
}

// SetTriggerId sets the TriggerId field's value.
func (s *TriggerForListPipelineRunsOutput) SetTriggerId(v string) *TriggerForListPipelineRunsOutput {
	s.TriggerId = &v
	return s
}

// SetType sets the Type field's value.
func (s *TriggerForListPipelineRunsOutput) SetType(v string) *TriggerForListPipelineRunsOutput {
	s.Type = &v
	return s
}

// SetUser sets the User field's value.
func (s *TriggerForListPipelineRunsOutput) SetUser(v *UserForListPipelineRunsOutput) *TriggerForListPipelineRunsOutput {
	s.User = v
	return s
}

type UserForListPipelineRunsOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	AccountId *int64 `type:"int64" json:",omitempty"`

	UserId *int64 `type:"int64" json:",omitempty"`
}

// String returns the string representation
func (s UserForListPipelineRunsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s UserForListPipelineRunsOutput) GoString() string {
	return s.String()
}

// SetAccountId sets the AccountId field's value.
func (s *UserForListPipelineRunsOutput) SetAccountId(v int64) *UserForListPipelineRunsOutput {
	s.AccountId = &v
	return s
}

// SetUserId sets the UserId field's value.
func (s *UserForListPipelineRunsOutput) SetUserId(v int64) *UserForListPipelineRunsOutput {
	s.UserId = &v
	return s
}

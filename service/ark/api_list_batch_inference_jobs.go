// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package ark

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opListBatchInferenceJobsCommon = "ListBatchInferenceJobs"

// ListBatchInferenceJobsCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ListBatchInferenceJobsCommon operation. The "output" return
// value will be populated with the ListBatchInferenceJobsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListBatchInferenceJobsCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListBatchInferenceJobsCommon Send returns without error.
//
// See ListBatchInferenceJobsCommon for more information on using the ListBatchInferenceJobsCommon
// API call, and error handling.
//
//    // Example sending a request using the ListBatchInferenceJobsCommonRequest method.
//    req, resp := client.ListBatchInferenceJobsCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ARK) ListBatchInferenceJobsCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opListBatchInferenceJobsCommon,
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

// ListBatchInferenceJobsCommon API operation for ARK.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ARK's
// API operation ListBatchInferenceJobsCommon for usage and error information.
func (c *ARK) ListBatchInferenceJobsCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ListBatchInferenceJobsCommonRequest(input)
	return out, req.Send()
}

// ListBatchInferenceJobsCommonWithContext is the same as ListBatchInferenceJobsCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ListBatchInferenceJobsCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ARK) ListBatchInferenceJobsCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ListBatchInferenceJobsCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opListBatchInferenceJobs = "ListBatchInferenceJobs"

// ListBatchInferenceJobsRequest generates a "volcengine/request.Request" representing the
// client's request for the ListBatchInferenceJobs operation. The "output" return
// value will be populated with the ListBatchInferenceJobsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListBatchInferenceJobsCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListBatchInferenceJobsCommon Send returns without error.
//
// See ListBatchInferenceJobs for more information on using the ListBatchInferenceJobs
// API call, and error handling.
//
//    // Example sending a request using the ListBatchInferenceJobsRequest method.
//    req, resp := client.ListBatchInferenceJobsRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ARK) ListBatchInferenceJobsRequest(input *ListBatchInferenceJobsInput) (req *request.Request, output *ListBatchInferenceJobsOutput) {
	op := &request.Operation{
		Name:       opListBatchInferenceJobs,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListBatchInferenceJobsInput{}
	}

	output = &ListBatchInferenceJobsOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// ListBatchInferenceJobs API operation for ARK.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ARK's
// API operation ListBatchInferenceJobs for usage and error information.
func (c *ARK) ListBatchInferenceJobs(input *ListBatchInferenceJobsInput) (*ListBatchInferenceJobsOutput, error) {
	req, out := c.ListBatchInferenceJobsRequest(input)
	return out, req.Send()
}

// ListBatchInferenceJobsWithContext is the same as ListBatchInferenceJobs with the addition of
// the ability to pass a context and additional request options.
//
// See ListBatchInferenceJobs for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ARK) ListBatchInferenceJobsWithContext(ctx volcengine.Context, input *ListBatchInferenceJobsInput, opts ...request.Option) (*ListBatchInferenceJobsOutput, error) {
	req, out := c.ListBatchInferenceJobsRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type FilterForListBatchInferenceJobsInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	CustomModelIds []*string `type:"list" json:",omitempty"`

	FoundationModels []*FoundationModelForListBatchInferenceJobsInput `type:"list" json:",omitempty"`

	Ids []*string `type:"list" json:",omitempty"`

	Name *string `type:"string" json:",omitempty"`

	Phases []*string `type:"list" json:",omitempty"`
}

// String returns the string representation
func (s FilterForListBatchInferenceJobsInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s FilterForListBatchInferenceJobsInput) GoString() string {
	return s.String()
}

// SetCustomModelIds sets the CustomModelIds field's value.
func (s *FilterForListBatchInferenceJobsInput) SetCustomModelIds(v []*string) *FilterForListBatchInferenceJobsInput {
	s.CustomModelIds = v
	return s
}

// SetFoundationModels sets the FoundationModels field's value.
func (s *FilterForListBatchInferenceJobsInput) SetFoundationModels(v []*FoundationModelForListBatchInferenceJobsInput) *FilterForListBatchInferenceJobsInput {
	s.FoundationModels = v
	return s
}

// SetIds sets the Ids field's value.
func (s *FilterForListBatchInferenceJobsInput) SetIds(v []*string) *FilterForListBatchInferenceJobsInput {
	s.Ids = v
	return s
}

// SetName sets the Name field's value.
func (s *FilterForListBatchInferenceJobsInput) SetName(v string) *FilterForListBatchInferenceJobsInput {
	s.Name = &v
	return s
}

// SetPhases sets the Phases field's value.
func (s *FilterForListBatchInferenceJobsInput) SetPhases(v []*string) *FilterForListBatchInferenceJobsInput {
	s.Phases = v
	return s
}

type FoundationModelForListBatchInferenceJobsInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	ModelVersions []*string `type:"list" json:",omitempty"`

	Name *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s FoundationModelForListBatchInferenceJobsInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s FoundationModelForListBatchInferenceJobsInput) GoString() string {
	return s.String()
}

// SetModelVersions sets the ModelVersions field's value.
func (s *FoundationModelForListBatchInferenceJobsInput) SetModelVersions(v []*string) *FoundationModelForListBatchInferenceJobsInput {
	s.ModelVersions = v
	return s
}

// SetName sets the Name field's value.
func (s *FoundationModelForListBatchInferenceJobsInput) SetName(v string) *FoundationModelForListBatchInferenceJobsInput {
	s.Name = &v
	return s
}

type FoundationModelForListBatchInferenceJobsOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	ModelVersion *string `type:"string" json:",omitempty"`

	Name *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s FoundationModelForListBatchInferenceJobsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s FoundationModelForListBatchInferenceJobsOutput) GoString() string {
	return s.String()
}

// SetModelVersion sets the ModelVersion field's value.
func (s *FoundationModelForListBatchInferenceJobsOutput) SetModelVersion(v string) *FoundationModelForListBatchInferenceJobsOutput {
	s.ModelVersion = &v
	return s
}

// SetName sets the Name field's value.
func (s *FoundationModelForListBatchInferenceJobsOutput) SetName(v string) *FoundationModelForListBatchInferenceJobsOutput {
	s.Name = &v
	return s
}

type InputFileTosLocationForListBatchInferenceJobsOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	BucketName *string `type:"string" json:",omitempty"`

	ObjectKey *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s InputFileTosLocationForListBatchInferenceJobsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s InputFileTosLocationForListBatchInferenceJobsOutput) GoString() string {
	return s.String()
}

// SetBucketName sets the BucketName field's value.
func (s *InputFileTosLocationForListBatchInferenceJobsOutput) SetBucketName(v string) *InputFileTosLocationForListBatchInferenceJobsOutput {
	s.BucketName = &v
	return s
}

// SetObjectKey sets the ObjectKey field's value.
func (s *InputFileTosLocationForListBatchInferenceJobsOutput) SetObjectKey(v string) *InputFileTosLocationForListBatchInferenceJobsOutput {
	s.ObjectKey = &v
	return s
}

type ItemForListBatchInferenceJobsOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	CompletionWindow *string `type:"string" json:",omitempty"`

	CreateTime *string `type:"string" json:",omitempty"`

	Description *string `type:"string" json:",omitempty"`

	ExpireTime *string `type:"string" json:",omitempty"`

	Id *string `type:"string" json:",omitempty"`

	InputFileTosLocation *InputFileTosLocationForListBatchInferenceJobsOutput `type:"structure" json:",omitempty"`

	ModelReference *ModelReferenceForListBatchInferenceJobsOutput `type:"structure" json:",omitempty"`

	Name *string `type:"string" json:",omitempty"`

	OutputDirTosLocation *OutputDirTosLocationForListBatchInferenceJobsOutput `type:"structure" json:",omitempty"`

	ProjectName *string `type:"string" json:",omitempty"`

	RequestCounts *RequestCountsForListBatchInferenceJobsOutput `type:"structure" json:",omitempty"`

	Status *StatusForListBatchInferenceJobsOutput `type:"structure" json:",omitempty"`

	Tags []*TagForListBatchInferenceJobsOutput `type:"list" json:",omitempty"`

	UpdateTime *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s ItemForListBatchInferenceJobsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ItemForListBatchInferenceJobsOutput) GoString() string {
	return s.String()
}

// SetCompletionWindow sets the CompletionWindow field's value.
func (s *ItemForListBatchInferenceJobsOutput) SetCompletionWindow(v string) *ItemForListBatchInferenceJobsOutput {
	s.CompletionWindow = &v
	return s
}

// SetCreateTime sets the CreateTime field's value.
func (s *ItemForListBatchInferenceJobsOutput) SetCreateTime(v string) *ItemForListBatchInferenceJobsOutput {
	s.CreateTime = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *ItemForListBatchInferenceJobsOutput) SetDescription(v string) *ItemForListBatchInferenceJobsOutput {
	s.Description = &v
	return s
}

// SetExpireTime sets the ExpireTime field's value.
func (s *ItemForListBatchInferenceJobsOutput) SetExpireTime(v string) *ItemForListBatchInferenceJobsOutput {
	s.ExpireTime = &v
	return s
}

// SetId sets the Id field's value.
func (s *ItemForListBatchInferenceJobsOutput) SetId(v string) *ItemForListBatchInferenceJobsOutput {
	s.Id = &v
	return s
}

// SetInputFileTosLocation sets the InputFileTosLocation field's value.
func (s *ItemForListBatchInferenceJobsOutput) SetInputFileTosLocation(v *InputFileTosLocationForListBatchInferenceJobsOutput) *ItemForListBatchInferenceJobsOutput {
	s.InputFileTosLocation = v
	return s
}

// SetModelReference sets the ModelReference field's value.
func (s *ItemForListBatchInferenceJobsOutput) SetModelReference(v *ModelReferenceForListBatchInferenceJobsOutput) *ItemForListBatchInferenceJobsOutput {
	s.ModelReference = v
	return s
}

// SetName sets the Name field's value.
func (s *ItemForListBatchInferenceJobsOutput) SetName(v string) *ItemForListBatchInferenceJobsOutput {
	s.Name = &v
	return s
}

// SetOutputDirTosLocation sets the OutputDirTosLocation field's value.
func (s *ItemForListBatchInferenceJobsOutput) SetOutputDirTosLocation(v *OutputDirTosLocationForListBatchInferenceJobsOutput) *ItemForListBatchInferenceJobsOutput {
	s.OutputDirTosLocation = v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *ItemForListBatchInferenceJobsOutput) SetProjectName(v string) *ItemForListBatchInferenceJobsOutput {
	s.ProjectName = &v
	return s
}

// SetRequestCounts sets the RequestCounts field's value.
func (s *ItemForListBatchInferenceJobsOutput) SetRequestCounts(v *RequestCountsForListBatchInferenceJobsOutput) *ItemForListBatchInferenceJobsOutput {
	s.RequestCounts = v
	return s
}

// SetStatus sets the Status field's value.
func (s *ItemForListBatchInferenceJobsOutput) SetStatus(v *StatusForListBatchInferenceJobsOutput) *ItemForListBatchInferenceJobsOutput {
	s.Status = v
	return s
}

// SetTags sets the Tags field's value.
func (s *ItemForListBatchInferenceJobsOutput) SetTags(v []*TagForListBatchInferenceJobsOutput) *ItemForListBatchInferenceJobsOutput {
	s.Tags = v
	return s
}

// SetUpdateTime sets the UpdateTime field's value.
func (s *ItemForListBatchInferenceJobsOutput) SetUpdateTime(v string) *ItemForListBatchInferenceJobsOutput {
	s.UpdateTime = &v
	return s
}

type ListBatchInferenceJobsInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Filter *FilterForListBatchInferenceJobsInput `type:"structure" json:",omitempty"`

	PageNumber *int32 `type:"int32" json:",omitempty"`

	PageSize *int32 `type:"int32" json:",omitempty"`

	ProjectName *string `type:"string" json:",omitempty"`

	SortBy *string `type:"string" json:",omitempty"`

	SortOrder *string `type:"string" json:",omitempty"`

	TagFilters []*TagFilterForListBatchInferenceJobsInput `type:"list" json:",omitempty"`
}

// String returns the string representation
func (s ListBatchInferenceJobsInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListBatchInferenceJobsInput) GoString() string {
	return s.String()
}

// SetFilter sets the Filter field's value.
func (s *ListBatchInferenceJobsInput) SetFilter(v *FilterForListBatchInferenceJobsInput) *ListBatchInferenceJobsInput {
	s.Filter = v
	return s
}

// SetPageNumber sets the PageNumber field's value.
func (s *ListBatchInferenceJobsInput) SetPageNumber(v int32) *ListBatchInferenceJobsInput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *ListBatchInferenceJobsInput) SetPageSize(v int32) *ListBatchInferenceJobsInput {
	s.PageSize = &v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *ListBatchInferenceJobsInput) SetProjectName(v string) *ListBatchInferenceJobsInput {
	s.ProjectName = &v
	return s
}

// SetSortBy sets the SortBy field's value.
func (s *ListBatchInferenceJobsInput) SetSortBy(v string) *ListBatchInferenceJobsInput {
	s.SortBy = &v
	return s
}

// SetSortOrder sets the SortOrder field's value.
func (s *ListBatchInferenceJobsInput) SetSortOrder(v string) *ListBatchInferenceJobsInput {
	s.SortOrder = &v
	return s
}

// SetTagFilters sets the TagFilters field's value.
func (s *ListBatchInferenceJobsInput) SetTagFilters(v []*TagFilterForListBatchInferenceJobsInput) *ListBatchInferenceJobsInput {
	s.TagFilters = v
	return s
}

type ListBatchInferenceJobsOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	Items []*ItemForListBatchInferenceJobsOutput `type:"list" json:",omitempty"`

	PageNumber *int32 `type:"int32" json:",omitempty"`

	PageSize *int32 `type:"int32" json:",omitempty"`

	TotalCount *int32 `type:"int32" json:",omitempty"`
}

// String returns the string representation
func (s ListBatchInferenceJobsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListBatchInferenceJobsOutput) GoString() string {
	return s.String()
}

// SetItems sets the Items field's value.
func (s *ListBatchInferenceJobsOutput) SetItems(v []*ItemForListBatchInferenceJobsOutput) *ListBatchInferenceJobsOutput {
	s.Items = v
	return s
}

// SetPageNumber sets the PageNumber field's value.
func (s *ListBatchInferenceJobsOutput) SetPageNumber(v int32) *ListBatchInferenceJobsOutput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *ListBatchInferenceJobsOutput) SetPageSize(v int32) *ListBatchInferenceJobsOutput {
	s.PageSize = &v
	return s
}

// SetTotalCount sets the TotalCount field's value.
func (s *ListBatchInferenceJobsOutput) SetTotalCount(v int32) *ListBatchInferenceJobsOutput {
	s.TotalCount = &v
	return s
}

type ModelReferenceForListBatchInferenceJobsOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	CustomModelId *string `type:"string" json:",omitempty"`

	FoundationModel *FoundationModelForListBatchInferenceJobsOutput `type:"structure" json:",omitempty"`
}

// String returns the string representation
func (s ModelReferenceForListBatchInferenceJobsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ModelReferenceForListBatchInferenceJobsOutput) GoString() string {
	return s.String()
}

// SetCustomModelId sets the CustomModelId field's value.
func (s *ModelReferenceForListBatchInferenceJobsOutput) SetCustomModelId(v string) *ModelReferenceForListBatchInferenceJobsOutput {
	s.CustomModelId = &v
	return s
}

// SetFoundationModel sets the FoundationModel field's value.
func (s *ModelReferenceForListBatchInferenceJobsOutput) SetFoundationModel(v *FoundationModelForListBatchInferenceJobsOutput) *ModelReferenceForListBatchInferenceJobsOutput {
	s.FoundationModel = v
	return s
}

type OutputDirTosLocationForListBatchInferenceJobsOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	BucketName *string `type:"string" json:",omitempty"`

	ObjectKey *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s OutputDirTosLocationForListBatchInferenceJobsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s OutputDirTosLocationForListBatchInferenceJobsOutput) GoString() string {
	return s.String()
}

// SetBucketName sets the BucketName field's value.
func (s *OutputDirTosLocationForListBatchInferenceJobsOutput) SetBucketName(v string) *OutputDirTosLocationForListBatchInferenceJobsOutput {
	s.BucketName = &v
	return s
}

// SetObjectKey sets the ObjectKey field's value.
func (s *OutputDirTosLocationForListBatchInferenceJobsOutput) SetObjectKey(v string) *OutputDirTosLocationForListBatchInferenceJobsOutput {
	s.ObjectKey = &v
	return s
}

type RequestCountsForListBatchInferenceJobsOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Completed *int64 `type:"int64" json:",omitempty"`

	Failed *int64 `type:"int64" json:",omitempty"`

	Total *int64 `type:"int64" json:",omitempty"`
}

// String returns the string representation
func (s RequestCountsForListBatchInferenceJobsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s RequestCountsForListBatchInferenceJobsOutput) GoString() string {
	return s.String()
}

// SetCompleted sets the Completed field's value.
func (s *RequestCountsForListBatchInferenceJobsOutput) SetCompleted(v int64) *RequestCountsForListBatchInferenceJobsOutput {
	s.Completed = &v
	return s
}

// SetFailed sets the Failed field's value.
func (s *RequestCountsForListBatchInferenceJobsOutput) SetFailed(v int64) *RequestCountsForListBatchInferenceJobsOutput {
	s.Failed = &v
	return s
}

// SetTotal sets the Total field's value.
func (s *RequestCountsForListBatchInferenceJobsOutput) SetTotal(v int64) *RequestCountsForListBatchInferenceJobsOutput {
	s.Total = &v
	return s
}

type StatusForListBatchInferenceJobsOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Message *string `type:"string" json:",omitempty"`

	Phase *string `type:"string" json:",omitempty"`

	PhaseTime *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s StatusForListBatchInferenceJobsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s StatusForListBatchInferenceJobsOutput) GoString() string {
	return s.String()
}

// SetMessage sets the Message field's value.
func (s *StatusForListBatchInferenceJobsOutput) SetMessage(v string) *StatusForListBatchInferenceJobsOutput {
	s.Message = &v
	return s
}

// SetPhase sets the Phase field's value.
func (s *StatusForListBatchInferenceJobsOutput) SetPhase(v string) *StatusForListBatchInferenceJobsOutput {
	s.Phase = &v
	return s
}

// SetPhaseTime sets the PhaseTime field's value.
func (s *StatusForListBatchInferenceJobsOutput) SetPhaseTime(v string) *StatusForListBatchInferenceJobsOutput {
	s.PhaseTime = &v
	return s
}

type TagFilterForListBatchInferenceJobsInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Key *string `type:"string" json:",omitempty"`

	Values []*string `type:"list" json:",omitempty"`
}

// String returns the string representation
func (s TagFilterForListBatchInferenceJobsInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s TagFilterForListBatchInferenceJobsInput) GoString() string {
	return s.String()
}

// SetKey sets the Key field's value.
func (s *TagFilterForListBatchInferenceJobsInput) SetKey(v string) *TagFilterForListBatchInferenceJobsInput {
	s.Key = &v
	return s
}

// SetValues sets the Values field's value.
func (s *TagFilterForListBatchInferenceJobsInput) SetValues(v []*string) *TagFilterForListBatchInferenceJobsInput {
	s.Values = v
	return s
}

type TagForListBatchInferenceJobsOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Key *string `type:"string" json:",omitempty"`

	Value *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s TagForListBatchInferenceJobsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s TagForListBatchInferenceJobsOutput) GoString() string {
	return s.String()
}

// SetKey sets the Key field's value.
func (s *TagForListBatchInferenceJobsOutput) SetKey(v string) *TagForListBatchInferenceJobsOutput {
	s.Key = &v
	return s
}

// SetValue sets the Value field's value.
func (s *TagForListBatchInferenceJobsOutput) SetValue(v string) *TagForListBatchInferenceJobsOutput {
	s.Value = &v
	return s
}

const (
	// EnumOfPhaseListForListBatchInferenceJobsInputQueued is a EnumOfPhaseListForListBatchInferenceJobsInput enum value
	EnumOfPhaseListForListBatchInferenceJobsInputQueued = "Queued"

	// EnumOfPhaseListForListBatchInferenceJobsInputRunning is a EnumOfPhaseListForListBatchInferenceJobsInput enum value
	EnumOfPhaseListForListBatchInferenceJobsInputRunning = "Running"

	// EnumOfPhaseListForListBatchInferenceJobsInputCompleted is a EnumOfPhaseListForListBatchInferenceJobsInput enum value
	EnumOfPhaseListForListBatchInferenceJobsInputCompleted = "Completed"

	// EnumOfPhaseListForListBatchInferenceJobsInputTerminating is a EnumOfPhaseListForListBatchInferenceJobsInput enum value
	EnumOfPhaseListForListBatchInferenceJobsInputTerminating = "Terminating"

	// EnumOfPhaseListForListBatchInferenceJobsInputTerminated is a EnumOfPhaseListForListBatchInferenceJobsInput enum value
	EnumOfPhaseListForListBatchInferenceJobsInputTerminated = "Terminated"

	// EnumOfPhaseListForListBatchInferenceJobsInputFailed is a EnumOfPhaseListForListBatchInferenceJobsInput enum value
	EnumOfPhaseListForListBatchInferenceJobsInputFailed = "Failed"
)

// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package vke

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opListAddonsCommon = "ListAddons"

// ListAddonsCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ListAddonsCommon operation. The "output" return
// value will be populated with the ListAddonsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListAddonsCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListAddonsCommon Send returns without error.
//
// See ListAddonsCommon for more information on using the ListAddonsCommon
// API call, and error handling.
//
//	// Example sending a request using the ListAddonsCommonRequest method.
//	req, resp := client.ListAddonsCommonRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *VKE) ListAddonsCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opListAddonsCommon,
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

// ListAddonsCommon API operation for VKE.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VKE's
// API operation ListAddonsCommon for usage and error information.
func (c *VKE) ListAddonsCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ListAddonsCommonRequest(input)
	return out, req.Send()
}

// ListAddonsCommonWithContext is the same as ListAddonsCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ListAddonsCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VKE) ListAddonsCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ListAddonsCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opListAddons = "ListAddons"

// ListAddonsRequest generates a "volcengine/request.Request" representing the
// client's request for the ListAddons operation. The "output" return
// value will be populated with the ListAddonsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListAddonsCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListAddonsCommon Send returns without error.
//
// See ListAddons for more information on using the ListAddons
// API call, and error handling.
//
//	// Example sending a request using the ListAddonsRequest method.
//	req, resp := client.ListAddonsRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *VKE) ListAddonsRequest(input *ListAddonsInput) (req *request.Request, output *ListAddonsOutput) {
	op := &request.Operation{
		Name:       opListAddons,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListAddonsInput{}
	}

	output = &ListAddonsOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// ListAddons API operation for VKE.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VKE's
// API operation ListAddons for usage and error information.
func (c *VKE) ListAddons(input *ListAddonsInput) (*ListAddonsOutput, error) {
	req, out := c.ListAddonsRequest(input)
	return out, req.Send()
}

// ListAddonsWithContext is the same as ListAddons with the addition of
// the ability to pass a context and additional request options.
//
// See ListAddons for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VKE) ListAddonsWithContext(ctx volcengine.Context, input *ListAddonsInput, opts ...request.Option) (*ListAddonsOutput, error) {
	req, out := c.ListAddonsRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ConditionForListAddonsOutput struct {
	_ struct{} `type:"structure"`

	Type *string `type:"string"`
}

// String returns the string representation
func (s ConditionForListAddonsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ConditionForListAddonsOutput) GoString() string {
	return s.String()
}

// SetType sets the Type field's value.
func (s *ConditionForListAddonsOutput) SetType(v string) *ConditionForListAddonsOutput {
	s.Type = &v
	return s
}

type FilterForListAddonsInput struct {
	_ struct{} `type:"structure"`

	ClusterIds []*string `type:"list"`

	CreateClientToken *string `type:"string"`

	DeployModes []*string `type:"list"`

	DeployNodeTypes []*string `type:"list"`

	Names []*string `type:"list"`

	Statuses []*StatusForListAddonsInput `type:"list"`

	UpdateClientToken *string `type:"string"`
}

// String returns the string representation
func (s FilterForListAddonsInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s FilterForListAddonsInput) GoString() string {
	return s.String()
}

// SetClusterIds sets the ClusterIds field's value.
func (s *FilterForListAddonsInput) SetClusterIds(v []*string) *FilterForListAddonsInput {
	s.ClusterIds = v
	return s
}

// SetCreateClientToken sets the CreateClientToken field's value.
func (s *FilterForListAddonsInput) SetCreateClientToken(v string) *FilterForListAddonsInput {
	s.CreateClientToken = &v
	return s
}

// SetDeployModes sets the DeployModes field's value.
func (s *FilterForListAddonsInput) SetDeployModes(v []*string) *FilterForListAddonsInput {
	s.DeployModes = v
	return s
}

// SetDeployNodeTypes sets the DeployNodeTypes field's value.
func (s *FilterForListAddonsInput) SetDeployNodeTypes(v []*string) *FilterForListAddonsInput {
	s.DeployNodeTypes = v
	return s
}

// SetNames sets the Names field's value.
func (s *FilterForListAddonsInput) SetNames(v []*string) *FilterForListAddonsInput {
	s.Names = v
	return s
}

// SetStatuses sets the Statuses field's value.
func (s *FilterForListAddonsInput) SetStatuses(v []*StatusForListAddonsInput) *FilterForListAddonsInput {
	s.Statuses = v
	return s
}

// SetUpdateClientToken sets the UpdateClientToken field's value.
func (s *FilterForListAddonsInput) SetUpdateClientToken(v string) *FilterForListAddonsInput {
	s.UpdateClientToken = &v
	return s
}

type ItemForListAddonsOutput struct {
	_ struct{} `type:"structure"`

	ClusterId *string `type:"string"`

	Config *string `type:"string"`

	CreateClientToken *string `type:"string"`

	CreateTime *string `type:"string"`

	DeployMode *string `type:"string"`

	DeployNodeType *string `type:"string"`

	Name *string `type:"string"`

	Status *StatusForListAddonsOutput `type:"structure"`

	UpdateClientToken *string `type:"string"`

	UpdateTime *string `type:"string"`

	Version *string `type:"string"`
}

// String returns the string representation
func (s ItemForListAddonsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ItemForListAddonsOutput) GoString() string {
	return s.String()
}

// SetClusterId sets the ClusterId field's value.
func (s *ItemForListAddonsOutput) SetClusterId(v string) *ItemForListAddonsOutput {
	s.ClusterId = &v
	return s
}

// SetConfig sets the Config field's value.
func (s *ItemForListAddonsOutput) SetConfig(v string) *ItemForListAddonsOutput {
	s.Config = &v
	return s
}

// SetCreateClientToken sets the CreateClientToken field's value.
func (s *ItemForListAddonsOutput) SetCreateClientToken(v string) *ItemForListAddonsOutput {
	s.CreateClientToken = &v
	return s
}

// SetCreateTime sets the CreateTime field's value.
func (s *ItemForListAddonsOutput) SetCreateTime(v string) *ItemForListAddonsOutput {
	s.CreateTime = &v
	return s
}

// SetDeployMode sets the DeployMode field's value.
func (s *ItemForListAddonsOutput) SetDeployMode(v string) *ItemForListAddonsOutput {
	s.DeployMode = &v
	return s
}

// SetDeployNodeType sets the DeployNodeType field's value.
func (s *ItemForListAddonsOutput) SetDeployNodeType(v string) *ItemForListAddonsOutput {
	s.DeployNodeType = &v
	return s
}

// SetName sets the Name field's value.
func (s *ItemForListAddonsOutput) SetName(v string) *ItemForListAddonsOutput {
	s.Name = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *ItemForListAddonsOutput) SetStatus(v *StatusForListAddonsOutput) *ItemForListAddonsOutput {
	s.Status = v
	return s
}

// SetUpdateClientToken sets the UpdateClientToken field's value.
func (s *ItemForListAddonsOutput) SetUpdateClientToken(v string) *ItemForListAddonsOutput {
	s.UpdateClientToken = &v
	return s
}

// SetUpdateTime sets the UpdateTime field's value.
func (s *ItemForListAddonsOutput) SetUpdateTime(v string) *ItemForListAddonsOutput {
	s.UpdateTime = &v
	return s
}

// SetVersion sets the Version field's value.
func (s *ItemForListAddonsOutput) SetVersion(v string) *ItemForListAddonsOutput {
	s.Version = &v
	return s
}

type ListAddonsInput struct {
	_ struct{} `type:"structure"`

	Filter *FilterForListAddonsInput `type:"structure"`

	PageNumber *int32 `type:"int32"`

	PageSize *int32 `type:"int32"`
}

// String returns the string representation
func (s ListAddonsInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListAddonsInput) GoString() string {
	return s.String()
}

// SetFilter sets the Filter field's value.
func (s *ListAddonsInput) SetFilter(v *FilterForListAddonsInput) *ListAddonsInput {
	s.Filter = v
	return s
}

// SetPageNumber sets the PageNumber field's value.
func (s *ListAddonsInput) SetPageNumber(v int32) *ListAddonsInput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *ListAddonsInput) SetPageSize(v int32) *ListAddonsInput {
	s.PageSize = &v
	return s
}

type ListAddonsOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	Items []*ItemForListAddonsOutput `type:"list"`

	PageNumber *int32 `type:"int32"`

	PageSize *int32 `type:"int32"`

	TotalCount *int32 `type:"int32"`
}

// String returns the string representation
func (s ListAddonsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListAddonsOutput) GoString() string {
	return s.String()
}

// SetItems sets the Items field's value.
func (s *ListAddonsOutput) SetItems(v []*ItemForListAddonsOutput) *ListAddonsOutput {
	s.Items = v
	return s
}

// SetPageNumber sets the PageNumber field's value.
func (s *ListAddonsOutput) SetPageNumber(v int32) *ListAddonsOutput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *ListAddonsOutput) SetPageSize(v int32) *ListAddonsOutput {
	s.PageSize = &v
	return s
}

// SetTotalCount sets the TotalCount field's value.
func (s *ListAddonsOutput) SetTotalCount(v int32) *ListAddonsOutput {
	s.TotalCount = &v
	return s
}

type StatusForListAddonsInput struct {
	_ struct{} `type:"structure"`

	ConditionsType *string `type:"string" json:"Conditions.Type" enum:"EnumOfConditionsTypeForListAddonsInput"`

	Phase *string `type:"string" json:"Phase" enum:"EnumOfPhaseForListAddonsInput"`
}

// String returns the string representation
func (s StatusForListAddonsInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s StatusForListAddonsInput) GoString() string {
	return s.String()
}

// SetConditionsType sets the ConditionsType field's value.
func (s *StatusForListAddonsInput) SetConditionsType(v string) *StatusForListAddonsInput {
	s.ConditionsType = &v
	return s
}

// SetPhase sets the Phase field's value.
func (s *StatusForListAddonsInput) SetPhase(v string) *StatusForListAddonsInput {
	s.Phase = &v
	return s
}

type StatusForListAddonsOutput struct {
	_ struct{} `type:"structure"`

	Conditions []*ConditionForListAddonsOutput `type:"list"`

	Phase *string `type:"string"`
}

// String returns the string representation
func (s StatusForListAddonsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s StatusForListAddonsOutput) GoString() string {
	return s.String()
}

// SetConditions sets the Conditions field's value.
func (s *StatusForListAddonsOutput) SetConditions(v []*ConditionForListAddonsOutput) *StatusForListAddonsOutput {
	s.Conditions = v
	return s
}

// SetPhase sets the Phase field's value.
func (s *StatusForListAddonsOutput) SetPhase(v string) *StatusForListAddonsOutput {
	s.Phase = &v
	return s
}

const (
	// EnumOfConditionsTypeForListAddonsInputClusterNotRunning is a EnumOfConditionsTypeForListAddonsInput enum value
	EnumOfConditionsTypeForListAddonsInputClusterNotRunning = "ClusterNotRunning"

	// EnumOfConditionsTypeForListAddonsInputClusterVersionUpgrading is a EnumOfConditionsTypeForListAddonsInput enum value
	EnumOfConditionsTypeForListAddonsInputClusterVersionUpgrading = "ClusterVersionUpgrading"

	// EnumOfConditionsTypeForListAddonsInputCrashLoopBackOff is a EnumOfConditionsTypeForListAddonsInput enum value
	EnumOfConditionsTypeForListAddonsInputCrashLoopBackOff = "CrashLoopBackOff"

	// EnumOfConditionsTypeForListAddonsInputDegraded is a EnumOfConditionsTypeForListAddonsInput enum value
	EnumOfConditionsTypeForListAddonsInputDegraded = "Degraded"

	// EnumOfConditionsTypeForListAddonsInputImagePullBackOff is a EnumOfConditionsTypeForListAddonsInput enum value
	EnumOfConditionsTypeForListAddonsInputImagePullBackOff = "ImagePullBackOff"

	// EnumOfConditionsTypeForListAddonsInputNameConflict is a EnumOfConditionsTypeForListAddonsInput enum value
	EnumOfConditionsTypeForListAddonsInputNameConflict = "NameConflict"

	// EnumOfConditionsTypeForListAddonsInputOk is a EnumOfConditionsTypeForListAddonsInput enum value
	EnumOfConditionsTypeForListAddonsInputOk = "Ok"

	// EnumOfConditionsTypeForListAddonsInputProgressing is a EnumOfConditionsTypeForListAddonsInput enum value
	EnumOfConditionsTypeForListAddonsInputProgressing = "Progressing"

	// EnumOfConditionsTypeForListAddonsInputResourceCleanupFailed is a EnumOfConditionsTypeForListAddonsInput enum value
	EnumOfConditionsTypeForListAddonsInputResourceCleanupFailed = "ResourceCleanupFailed"

	// EnumOfConditionsTypeForListAddonsInputSchedulingFailed is a EnumOfConditionsTypeForListAddonsInput enum value
	EnumOfConditionsTypeForListAddonsInputSchedulingFailed = "SchedulingFailed"

	// EnumOfConditionsTypeForListAddonsInputUnknown is a EnumOfConditionsTypeForListAddonsInput enum value
	EnumOfConditionsTypeForListAddonsInputUnknown = "Unknown"
)

const (
	// EnumOfDeployModeListForListAddonsInputManaged is a EnumOfDeployModeListForListAddonsInput enum value
	EnumOfDeployModeListForListAddonsInputManaged = "Managed"

	// EnumOfDeployModeListForListAddonsInputUnmanaged is a EnumOfDeployModeListForListAddonsInput enum value
	EnumOfDeployModeListForListAddonsInputUnmanaged = "Unmanaged"
)

const (
	// EnumOfDeployNodeTypeListForListAddonsInputEdgeNode is a EnumOfDeployNodeTypeListForListAddonsInput enum value
	EnumOfDeployNodeTypeListForListAddonsInputEdgeNode = "EdgeNode"

	// EnumOfDeployNodeTypeListForListAddonsInputNode is a EnumOfDeployNodeTypeListForListAddonsInput enum value
	EnumOfDeployNodeTypeListForListAddonsInputNode = "Node"

	// EnumOfDeployNodeTypeListForListAddonsInputVirtualNode is a EnumOfDeployNodeTypeListForListAddonsInput enum value
	EnumOfDeployNodeTypeListForListAddonsInputVirtualNode = "VirtualNode"
)

const (
	// EnumOfPhaseForListAddonsInputCreating is a EnumOfPhaseForListAddonsInput enum value
	EnumOfPhaseForListAddonsInputCreating = "Creating"

	// EnumOfPhaseForListAddonsInputDeleting is a EnumOfPhaseForListAddonsInput enum value
	EnumOfPhaseForListAddonsInputDeleting = "Deleting"

	// EnumOfPhaseForListAddonsInputFailed is a EnumOfPhaseForListAddonsInput enum value
	EnumOfPhaseForListAddonsInputFailed = "Failed"

	// EnumOfPhaseForListAddonsInputNone is a EnumOfPhaseForListAddonsInput enum value
	EnumOfPhaseForListAddonsInputNone = "None"

	// EnumOfPhaseForListAddonsInputRunning is a EnumOfPhaseForListAddonsInput enum value
	EnumOfPhaseForListAddonsInputRunning = "Running"

	// EnumOfPhaseForListAddonsInputUpdating is a EnumOfPhaseForListAddonsInput enum value
	EnumOfPhaseForListAddonsInputUpdating = "Updating"
)

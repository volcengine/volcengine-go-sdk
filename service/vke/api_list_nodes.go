// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package vke

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opListNodesCommon = "ListNodes"

// ListNodesCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ListNodesCommon operation. The "output" return
// value will be populated with the ListNodesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListNodesCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListNodesCommon Send returns without error.
//
// See ListNodesCommon for more information on using the ListNodesCommon
// API call, and error handling.
//
//    // Example sending a request using the ListNodesCommonRequest method.
//    req, resp := client.ListNodesCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VKE) ListNodesCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opListNodesCommon,
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

// ListNodesCommon API operation for VKE.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VKE's
// API operation ListNodesCommon for usage and error information.
func (c *VKE) ListNodesCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ListNodesCommonRequest(input)
	return out, req.Send()
}

// ListNodesCommonWithContext is the same as ListNodesCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ListNodesCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VKE) ListNodesCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ListNodesCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opListNodes = "ListNodes"

// ListNodesRequest generates a "volcengine/request.Request" representing the
// client's request for the ListNodes operation. The "output" return
// value will be populated with the ListNodesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListNodesCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListNodesCommon Send returns without error.
//
// See ListNodes for more information on using the ListNodes
// API call, and error handling.
//
//    // Example sending a request using the ListNodesRequest method.
//    req, resp := client.ListNodesRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VKE) ListNodesRequest(input *ListNodesInput) (req *request.Request, output *ListNodesOutput) {
	op := &request.Operation{
		Name:       opListNodes,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListNodesInput{}
	}

	output = &ListNodesOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// ListNodes API operation for VKE.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VKE's
// API operation ListNodes for usage and error information.
func (c *VKE) ListNodes(input *ListNodesInput) (*ListNodesOutput, error) {
	req, out := c.ListNodesRequest(input)
	return out, req.Send()
}

// ListNodesWithContext is the same as ListNodes with the addition of
// the ability to pass a context and additional request options.
//
// See ListNodes for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VKE) ListNodesWithContext(ctx volcengine.Context, input *ListNodesInput, opts ...request.Option) (*ListNodesOutput, error) {
	req, out := c.ListNodesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ConditionForListNodesOutput struct {
	_ struct{} `type:"structure"`

	Type *string `type:"string"`
}

// String returns the string representation
func (s ConditionForListNodesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ConditionForListNodesOutput) GoString() string {
	return s.String()
}

// SetType sets the Type field's value.
func (s *ConditionForListNodesOutput) SetType(v string) *ConditionForListNodesOutput {
	s.Type = &v
	return s
}

type FilterForListNodesInput struct {
	_ struct{} `type:"structure"`

	ClusterIds []*string `type:"list"`

	CreateClientToken *string `type:"string"`

	Ids []*string `type:"list"`

	Name *string `type:"string"`

	NodePoolIds []*string `type:"list"`

	Statuses []*StatusForListNodesInput `type:"list"`

	ZoneIds []*string `type:"list"`
}

// String returns the string representation
func (s FilterForListNodesInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s FilterForListNodesInput) GoString() string {
	return s.String()
}

// SetClusterIds sets the ClusterIds field's value.
func (s *FilterForListNodesInput) SetClusterIds(v []*string) *FilterForListNodesInput {
	s.ClusterIds = v
	return s
}

// SetCreateClientToken sets the CreateClientToken field's value.
func (s *FilterForListNodesInput) SetCreateClientToken(v string) *FilterForListNodesInput {
	s.CreateClientToken = &v
	return s
}

// SetIds sets the Ids field's value.
func (s *FilterForListNodesInput) SetIds(v []*string) *FilterForListNodesInput {
	s.Ids = v
	return s
}

// SetName sets the Name field's value.
func (s *FilterForListNodesInput) SetName(v string) *FilterForListNodesInput {
	s.Name = &v
	return s
}

// SetNodePoolIds sets the NodePoolIds field's value.
func (s *FilterForListNodesInput) SetNodePoolIds(v []*string) *FilterForListNodesInput {
	s.NodePoolIds = v
	return s
}

// SetStatuses sets the Statuses field's value.
func (s *FilterForListNodesInput) SetStatuses(v []*StatusForListNodesInput) *FilterForListNodesInput {
	s.Statuses = v
	return s
}

// SetZoneIds sets the ZoneIds field's value.
func (s *FilterForListNodesInput) SetZoneIds(v []*string) *FilterForListNodesInput {
	s.ZoneIds = v
	return s
}

type ItemForListNodesOutput struct {
	_ struct{} `type:"structure"`

	AdditionalContainerStorageEnabled *bool `type:"boolean"`

	ClusterId *string `type:"string"`

	ContainerStoragePath *string `type:"string"`

	CreateClientToken *string `type:"string"`

	CreateTime *string `type:"string"`

	Id *string `type:"string"`

	InstanceId *string `type:"string"`

	IsVirtual *bool `type:"boolean"`

	Name *string `type:"string"`

	NodePoolId *string `type:"string"`

	Roles []*string `type:"list"`

	Status *StatusForListNodesOutput `type:"structure"`

	UpdateTime *string `type:"string"`

	ZoneId *string `type:"string"`
}

// String returns the string representation
func (s ItemForListNodesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ItemForListNodesOutput) GoString() string {
	return s.String()
}

// SetAdditionalContainerStorageEnabled sets the AdditionalContainerStorageEnabled field's value.
func (s *ItemForListNodesOutput) SetAdditionalContainerStorageEnabled(v bool) *ItemForListNodesOutput {
	s.AdditionalContainerStorageEnabled = &v
	return s
}

// SetClusterId sets the ClusterId field's value.
func (s *ItemForListNodesOutput) SetClusterId(v string) *ItemForListNodesOutput {
	s.ClusterId = &v
	return s
}

// SetContainerStoragePath sets the ContainerStoragePath field's value.
func (s *ItemForListNodesOutput) SetContainerStoragePath(v string) *ItemForListNodesOutput {
	s.ContainerStoragePath = &v
	return s
}

// SetCreateClientToken sets the CreateClientToken field's value.
func (s *ItemForListNodesOutput) SetCreateClientToken(v string) *ItemForListNodesOutput {
	s.CreateClientToken = &v
	return s
}

// SetCreateTime sets the CreateTime field's value.
func (s *ItemForListNodesOutput) SetCreateTime(v string) *ItemForListNodesOutput {
	s.CreateTime = &v
	return s
}

// SetId sets the Id field's value.
func (s *ItemForListNodesOutput) SetId(v string) *ItemForListNodesOutput {
	s.Id = &v
	return s
}

// SetInstanceId sets the InstanceId field's value.
func (s *ItemForListNodesOutput) SetInstanceId(v string) *ItemForListNodesOutput {
	s.InstanceId = &v
	return s
}

// SetIsVirtual sets the IsVirtual field's value.
func (s *ItemForListNodesOutput) SetIsVirtual(v bool) *ItemForListNodesOutput {
	s.IsVirtual = &v
	return s
}

// SetName sets the Name field's value.
func (s *ItemForListNodesOutput) SetName(v string) *ItemForListNodesOutput {
	s.Name = &v
	return s
}

// SetNodePoolId sets the NodePoolId field's value.
func (s *ItemForListNodesOutput) SetNodePoolId(v string) *ItemForListNodesOutput {
	s.NodePoolId = &v
	return s
}

// SetRoles sets the Roles field's value.
func (s *ItemForListNodesOutput) SetRoles(v []*string) *ItemForListNodesOutput {
	s.Roles = v
	return s
}

// SetStatus sets the Status field's value.
func (s *ItemForListNodesOutput) SetStatus(v *StatusForListNodesOutput) *ItemForListNodesOutput {
	s.Status = v
	return s
}

// SetUpdateTime sets the UpdateTime field's value.
func (s *ItemForListNodesOutput) SetUpdateTime(v string) *ItemForListNodesOutput {
	s.UpdateTime = &v
	return s
}

// SetZoneId sets the ZoneId field's value.
func (s *ItemForListNodesOutput) SetZoneId(v string) *ItemForListNodesOutput {
	s.ZoneId = &v
	return s
}

type ListNodesInput struct {
	_ struct{} `type:"structure"`

	Filter *FilterForListNodesInput `type:"structure"`

	PageNumber *int32 `type:"int32"`

	PageSize *int32 `type:"int32"`
}

// String returns the string representation
func (s ListNodesInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListNodesInput) GoString() string {
	return s.String()
}

// SetFilter sets the Filter field's value.
func (s *ListNodesInput) SetFilter(v *FilterForListNodesInput) *ListNodesInput {
	s.Filter = v
	return s
}

// SetPageNumber sets the PageNumber field's value.
func (s *ListNodesInput) SetPageNumber(v int32) *ListNodesInput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *ListNodesInput) SetPageSize(v int32) *ListNodesInput {
	s.PageSize = &v
	return s
}

type ListNodesOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	Items []*ItemForListNodesOutput `type:"list"`

	PageNumber *int32 `type:"int32"`

	PageSize *int32 `type:"int32"`

	TotalCount *int32 `type:"int32"`
}

// String returns the string representation
func (s ListNodesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListNodesOutput) GoString() string {
	return s.String()
}

// SetItems sets the Items field's value.
func (s *ListNodesOutput) SetItems(v []*ItemForListNodesOutput) *ListNodesOutput {
	s.Items = v
	return s
}

// SetPageNumber sets the PageNumber field's value.
func (s *ListNodesOutput) SetPageNumber(v int32) *ListNodesOutput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *ListNodesOutput) SetPageSize(v int32) *ListNodesOutput {
	s.PageSize = &v
	return s
}

// SetTotalCount sets the TotalCount field's value.
func (s *ListNodesOutput) SetTotalCount(v int32) *ListNodesOutput {
	s.TotalCount = &v
	return s
}

type StatusForListNodesInput struct {
	_ struct{} `type:"structure"`

	ConditionsType *string `type:"string" json:"Conditions.Type" enum:"EnumOfConditionsTypeForListNodesInput"`

	Phase *string `type:"string" enum:"EnumOfPhaseForListNodesInput"`
}

// String returns the string representation
func (s StatusForListNodesInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s StatusForListNodesInput) GoString() string {
	return s.String()
}

// SetConditionsType sets the ConditionsType field's value.
func (s *StatusForListNodesInput) SetConditionsType(v string) *StatusForListNodesInput {
	s.ConditionsType = &v
	return s
}

// SetPhase sets the Phase field's value.
func (s *StatusForListNodesInput) SetPhase(v string) *StatusForListNodesInput {
	s.Phase = &v
	return s
}

type StatusForListNodesOutput struct {
	_ struct{} `type:"structure"`

	Conditions []*ConditionForListNodesOutput `type:"list"`

	Phase *string `type:"string"`
}

// String returns the string representation
func (s StatusForListNodesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s StatusForListNodesOutput) GoString() string {
	return s.String()
}

// SetConditions sets the Conditions field's value.
func (s *StatusForListNodesOutput) SetConditions(v []*ConditionForListNodesOutput) *StatusForListNodesOutput {
	s.Conditions = v
	return s
}

// SetPhase sets the Phase field's value.
func (s *StatusForListNodesOutput) SetPhase(v string) *StatusForListNodesOutput {
	s.Phase = &v
	return s
}

const (
	// EnumOfConditionsTypeForListNodesInputBalance is a EnumOfConditionsTypeForListNodesInput enum value
	EnumOfConditionsTypeForListNodesInputBalance = "Balance"

	// EnumOfConditionsTypeForListNodesInputInitializeFailed is a EnumOfConditionsTypeForListNodesInput enum value
	EnumOfConditionsTypeForListNodesInputInitializeFailed = "InitializeFailed"

	// EnumOfConditionsTypeForListNodesInputNotReady is a EnumOfConditionsTypeForListNodesInput enum value
	EnumOfConditionsTypeForListNodesInputNotReady = "NotReady"

	// EnumOfConditionsTypeForListNodesInputOk is a EnumOfConditionsTypeForListNodesInput enum value
	EnumOfConditionsTypeForListNodesInputOk = "Ok"

	// EnumOfConditionsTypeForListNodesInputProgressing is a EnumOfConditionsTypeForListNodesInput enum value
	EnumOfConditionsTypeForListNodesInputProgressing = "Progressing"

	// EnumOfConditionsTypeForListNodesInputResourceCleanupFailed is a EnumOfConditionsTypeForListNodesInput enum value
	EnumOfConditionsTypeForListNodesInputResourceCleanupFailed = "ResourceCleanupFailed"

	// EnumOfConditionsTypeForListNodesInputSecurity is a EnumOfConditionsTypeForListNodesInput enum value
	EnumOfConditionsTypeForListNodesInputSecurity = "Security"

	// EnumOfConditionsTypeForListNodesInputUnknown is a EnumOfConditionsTypeForListNodesInput enum value
	EnumOfConditionsTypeForListNodesInputUnknown = "Unknown"

	// EnumOfConditionsTypeForListNodesInputUnschedulable is a EnumOfConditionsTypeForListNodesInput enum value
	EnumOfConditionsTypeForListNodesInputUnschedulable = "Unschedulable"
)

const (
	// EnumOfPhaseForListNodesInputCreating is a EnumOfPhaseForListNodesInput enum value
	EnumOfPhaseForListNodesInputCreating = "Creating"

	// EnumOfPhaseForListNodesInputDeleting is a EnumOfPhaseForListNodesInput enum value
	EnumOfPhaseForListNodesInputDeleting = "Deleting"

	// EnumOfPhaseForListNodesInputFailed is a EnumOfPhaseForListNodesInput enum value
	EnumOfPhaseForListNodesInputFailed = "Failed"

	// EnumOfPhaseForListNodesInputRunning is a EnumOfPhaseForListNodesInput enum value
	EnumOfPhaseForListNodesInputRunning = "Running"

	// EnumOfPhaseForListNodesInputStarting is a EnumOfPhaseForListNodesInput enum value
	EnumOfPhaseForListNodesInputStarting = "Starting"

	// EnumOfPhaseForListNodesInputStopped is a EnumOfPhaseForListNodesInput enum value
	EnumOfPhaseForListNodesInputStopped = "Stopped"

	// EnumOfPhaseForListNodesInputStopping is a EnumOfPhaseForListNodesInput enum value
	EnumOfPhaseForListNodesInputStopping = "Stopping"

	// EnumOfPhaseForListNodesInputUpdating is a EnumOfPhaseForListNodesInput enum value
	EnumOfPhaseForListNodesInputUpdating = "Updating"
)

const (
	// EnumOfRoleListForListNodesOutputEtcd is a EnumOfRoleListForListNodesOutput enum value
	EnumOfRoleListForListNodesOutputEtcd = "Etcd"

	// EnumOfRoleListForListNodesOutputMaster is a EnumOfRoleListForListNodesOutput enum value
	EnumOfRoleListForListNodesOutputMaster = "Master"

	// EnumOfRoleListForListNodesOutputWorker is a EnumOfRoleListForListNodesOutput enum value
	EnumOfRoleListForListNodesOutputWorker = "Worker"
)

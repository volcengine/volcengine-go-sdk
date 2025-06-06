// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package vmp

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opListWorkspaceStatusCommon = "ListWorkspaceStatus"

// ListWorkspaceStatusCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ListWorkspaceStatusCommon operation. The "output" return
// value will be populated with the ListWorkspaceStatusCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListWorkspaceStatusCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListWorkspaceStatusCommon Send returns without error.
//
// See ListWorkspaceStatusCommon for more information on using the ListWorkspaceStatusCommon
// API call, and error handling.
//
//    // Example sending a request using the ListWorkspaceStatusCommonRequest method.
//    req, resp := client.ListWorkspaceStatusCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VMP) ListWorkspaceStatusCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opListWorkspaceStatusCommon,
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

// ListWorkspaceStatusCommon API operation for VMP.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VMP's
// API operation ListWorkspaceStatusCommon for usage and error information.
func (c *VMP) ListWorkspaceStatusCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ListWorkspaceStatusCommonRequest(input)
	return out, req.Send()
}

// ListWorkspaceStatusCommonWithContext is the same as ListWorkspaceStatusCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ListWorkspaceStatusCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VMP) ListWorkspaceStatusCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ListWorkspaceStatusCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opListWorkspaceStatus = "ListWorkspaceStatus"

// ListWorkspaceStatusRequest generates a "volcengine/request.Request" representing the
// client's request for the ListWorkspaceStatus operation. The "output" return
// value will be populated with the ListWorkspaceStatusCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListWorkspaceStatusCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListWorkspaceStatusCommon Send returns without error.
//
// See ListWorkspaceStatus for more information on using the ListWorkspaceStatus
// API call, and error handling.
//
//    // Example sending a request using the ListWorkspaceStatusRequest method.
//    req, resp := client.ListWorkspaceStatusRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VMP) ListWorkspaceStatusRequest(input *ListWorkspaceStatusInput) (req *request.Request, output *ListWorkspaceStatusOutput) {
	op := &request.Operation{
		Name:       opListWorkspaceStatus,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListWorkspaceStatusInput{}
	}

	output = &ListWorkspaceStatusOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// ListWorkspaceStatus API operation for VMP.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VMP's
// API operation ListWorkspaceStatus for usage and error information.
func (c *VMP) ListWorkspaceStatus(input *ListWorkspaceStatusInput) (*ListWorkspaceStatusOutput, error) {
	req, out := c.ListWorkspaceStatusRequest(input)
	return out, req.Send()
}

// ListWorkspaceStatusWithContext is the same as ListWorkspaceStatus with the addition of
// the ability to pass a context and additional request options.
//
// See ListWorkspaceStatus for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VMP) ListWorkspaceStatusWithContext(ctx volcengine.Context, input *ListWorkspaceStatusInput, opts ...request.Option) (*ListWorkspaceStatusOutput, error) {
	req, out := c.ListWorkspaceStatusRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type FiltersForListWorkspaceStatusInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Ids []*string `type:"list" json:",omitempty"`

	InstanceTypeIds []*string `type:"list" json:",omitempty"`

	Name *string `type:"string" json:",omitempty"`

	Statuses []*string `type:"list" json:",omitempty"`
}

// String returns the string representation
func (s FiltersForListWorkspaceStatusInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s FiltersForListWorkspaceStatusInput) GoString() string {
	return s.String()
}

// SetIds sets the Ids field's value.
func (s *FiltersForListWorkspaceStatusInput) SetIds(v []*string) *FiltersForListWorkspaceStatusInput {
	s.Ids = v
	return s
}

// SetInstanceTypeIds sets the InstanceTypeIds field's value.
func (s *FiltersForListWorkspaceStatusInput) SetInstanceTypeIds(v []*string) *FiltersForListWorkspaceStatusInput {
	s.InstanceTypeIds = v
	return s
}

// SetName sets the Name field's value.
func (s *FiltersForListWorkspaceStatusInput) SetName(v string) *FiltersForListWorkspaceStatusInput {
	s.Name = &v
	return s
}

// SetStatuses sets the Statuses field's value.
func (s *FiltersForListWorkspaceStatusInput) SetStatuses(v []*string) *FiltersForListWorkspaceStatusInput {
	s.Statuses = v
	return s
}

type ItemForListWorkspaceStatusOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Id *string `type:"string" json:",omitempty"`

	InstanceTypeId *string `type:"string" json:",omitempty"`

	Name *string `type:"string" json:",omitempty"`

	Status *string `type:"string" json:",omitempty"`

	Usage *UsageForListWorkspaceStatusOutput `type:"structure" json:",omitempty"`
}

// String returns the string representation
func (s ItemForListWorkspaceStatusOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ItemForListWorkspaceStatusOutput) GoString() string {
	return s.String()
}

// SetId sets the Id field's value.
func (s *ItemForListWorkspaceStatusOutput) SetId(v string) *ItemForListWorkspaceStatusOutput {
	s.Id = &v
	return s
}

// SetInstanceTypeId sets the InstanceTypeId field's value.
func (s *ItemForListWorkspaceStatusOutput) SetInstanceTypeId(v string) *ItemForListWorkspaceStatusOutput {
	s.InstanceTypeId = &v
	return s
}

// SetName sets the Name field's value.
func (s *ItemForListWorkspaceStatusOutput) SetName(v string) *ItemForListWorkspaceStatusOutput {
	s.Name = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *ItemForListWorkspaceStatusOutput) SetStatus(v string) *ItemForListWorkspaceStatusOutput {
	s.Status = &v
	return s
}

// SetUsage sets the Usage field's value.
func (s *ItemForListWorkspaceStatusOutput) SetUsage(v *UsageForListWorkspaceStatusOutput) *ItemForListWorkspaceStatusOutput {
	s.Usage = v
	return s
}

type ListWorkspaceStatusInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Filters *FiltersForListWorkspaceStatusInput `type:"structure" json:",omitempty"`

	PageNumber *int64 `type:"int64" json:",omitempty"`

	PageSize *int64 `type:"int64" json:",omitempty"`

	ProjectName *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s ListWorkspaceStatusInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListWorkspaceStatusInput) GoString() string {
	return s.String()
}

// SetFilters sets the Filters field's value.
func (s *ListWorkspaceStatusInput) SetFilters(v *FiltersForListWorkspaceStatusInput) *ListWorkspaceStatusInput {
	s.Filters = v
	return s
}

// SetPageNumber sets the PageNumber field's value.
func (s *ListWorkspaceStatusInput) SetPageNumber(v int64) *ListWorkspaceStatusInput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *ListWorkspaceStatusInput) SetPageSize(v int64) *ListWorkspaceStatusInput {
	s.PageSize = &v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *ListWorkspaceStatusInput) SetProjectName(v string) *ListWorkspaceStatusInput {
	s.ProjectName = &v
	return s
}

type ListWorkspaceStatusOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	Items []*ItemForListWorkspaceStatusOutput `type:"list" json:",omitempty"`

	Total *int64 `type:"int64" json:",omitempty"`
}

// String returns the string representation
func (s ListWorkspaceStatusOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListWorkspaceStatusOutput) GoString() string {
	return s.String()
}

// SetItems sets the Items field's value.
func (s *ListWorkspaceStatusOutput) SetItems(v []*ItemForListWorkspaceStatusOutput) *ListWorkspaceStatusOutput {
	s.Items = v
	return s
}

// SetTotal sets the Total field's value.
func (s *ListWorkspaceStatusOutput) SetTotal(v int64) *ListWorkspaceStatusOutput {
	s.Total = &v
	return s
}

type UsageForListWorkspaceStatusOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	ActiveSeries *int64 `type:"int64" json:",omitempty"`

	IngestedSamplesPerSecond *float64 `type:"double" json:",omitempty"`
}

// String returns the string representation
func (s UsageForListWorkspaceStatusOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s UsageForListWorkspaceStatusOutput) GoString() string {
	return s.String()
}

// SetActiveSeries sets the ActiveSeries field's value.
func (s *UsageForListWorkspaceStatusOutput) SetActiveSeries(v int64) *UsageForListWorkspaceStatusOutput {
	s.ActiveSeries = &v
	return s
}

// SetIngestedSamplesPerSecond sets the IngestedSamplesPerSecond field's value.
func (s *UsageForListWorkspaceStatusOutput) SetIngestedSamplesPerSecond(v float64) *UsageForListWorkspaceStatusOutput {
	s.IngestedSamplesPerSecond = &v
	return s
}

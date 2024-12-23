// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package cp

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opListWorkspacesCommon = "ListWorkspaces"

// ListWorkspacesCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ListWorkspacesCommon operation. The "output" return
// value will be populated with the ListWorkspacesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListWorkspacesCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListWorkspacesCommon Send returns without error.
//
// See ListWorkspacesCommon for more information on using the ListWorkspacesCommon
// API call, and error handling.
//
//    // Example sending a request using the ListWorkspacesCommonRequest method.
//    req, resp := client.ListWorkspacesCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CP) ListWorkspacesCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opListWorkspacesCommon,
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

// ListWorkspacesCommon API operation for CP.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CP's
// API operation ListWorkspacesCommon for usage and error information.
func (c *CP) ListWorkspacesCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ListWorkspacesCommonRequest(input)
	return out, req.Send()
}

// ListWorkspacesCommonWithContext is the same as ListWorkspacesCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ListWorkspacesCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CP) ListWorkspacesCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ListWorkspacesCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opListWorkspaces = "ListWorkspaces"

// ListWorkspacesRequest generates a "volcengine/request.Request" representing the
// client's request for the ListWorkspaces operation. The "output" return
// value will be populated with the ListWorkspacesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListWorkspacesCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListWorkspacesCommon Send returns without error.
//
// See ListWorkspaces for more information on using the ListWorkspaces
// API call, and error handling.
//
//    // Example sending a request using the ListWorkspacesRequest method.
//    req, resp := client.ListWorkspacesRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CP) ListWorkspacesRequest(input *ListWorkspacesInput) (req *request.Request, output *ListWorkspacesOutput) {
	op := &request.Operation{
		Name:       opListWorkspaces,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListWorkspacesInput{}
	}

	output = &ListWorkspacesOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// ListWorkspaces API operation for CP.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CP's
// API operation ListWorkspaces for usage and error information.
func (c *CP) ListWorkspaces(input *ListWorkspacesInput) (*ListWorkspacesOutput, error) {
	req, out := c.ListWorkspacesRequest(input)
	return out, req.Send()
}

// ListWorkspacesWithContext is the same as ListWorkspaces with the addition of
// the ability to pass a context and additional request options.
//
// See ListWorkspaces for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CP) ListWorkspacesWithContext(ctx volcengine.Context, input *ListWorkspacesInput, opts ...request.Option) (*ListWorkspacesOutput, error) {
	req, out := c.ListWorkspacesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type CreatorForListWorkspacesOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	AccountId *int64 `type:"int64" json:",omitempty"`

	UserId *int64 `type:"int64" json:",omitempty"`
}

// String returns the string representation
func (s CreatorForListWorkspacesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreatorForListWorkspacesOutput) GoString() string {
	return s.String()
}

// SetAccountId sets the AccountId field's value.
func (s *CreatorForListWorkspacesOutput) SetAccountId(v int64) *CreatorForListWorkspacesOutput {
	s.AccountId = &v
	return s
}

// SetUserId sets the UserId field's value.
func (s *CreatorForListWorkspacesOutput) SetUserId(v int64) *CreatorForListWorkspacesOutput {
	s.UserId = &v
	return s
}

type FilterForListWorkspacesInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Ids []*string `type:"list" json:",omitempty"`

	Name *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s FilterForListWorkspacesInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s FilterForListWorkspacesInput) GoString() string {
	return s.String()
}

// SetIds sets the Ids field's value.
func (s *FilterForListWorkspacesInput) SetIds(v []*string) *FilterForListWorkspacesInput {
	s.Ids = v
	return s
}

// SetName sets the Name field's value.
func (s *FilterForListWorkspacesInput) SetName(v string) *FilterForListWorkspacesInput {
	s.Name = &v
	return s
}

type ItemForListWorkspacesOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	CreateTime *string `type:"string" json:",omitempty"`

	Creator *CreatorForListWorkspacesOutput `type:"structure" json:",omitempty"`

	Description *string `type:"string" json:",omitempty"`

	Id *string `type:"string" json:",omitempty"`

	Name *string `type:"string" json:",omitempty"`

	UpdateTime *string `type:"string" json:",omitempty"`

	Visibility *string `type:"string" json:",omitempty"`

	VisibleUsers []*VisibleUserForListWorkspacesOutput `type:"list" json:",omitempty"`
}

// String returns the string representation
func (s ItemForListWorkspacesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ItemForListWorkspacesOutput) GoString() string {
	return s.String()
}

// SetCreateTime sets the CreateTime field's value.
func (s *ItemForListWorkspacesOutput) SetCreateTime(v string) *ItemForListWorkspacesOutput {
	s.CreateTime = &v
	return s
}

// SetCreator sets the Creator field's value.
func (s *ItemForListWorkspacesOutput) SetCreator(v *CreatorForListWorkspacesOutput) *ItemForListWorkspacesOutput {
	s.Creator = v
	return s
}

// SetDescription sets the Description field's value.
func (s *ItemForListWorkspacesOutput) SetDescription(v string) *ItemForListWorkspacesOutput {
	s.Description = &v
	return s
}

// SetId sets the Id field's value.
func (s *ItemForListWorkspacesOutput) SetId(v string) *ItemForListWorkspacesOutput {
	s.Id = &v
	return s
}

// SetName sets the Name field's value.
func (s *ItemForListWorkspacesOutput) SetName(v string) *ItemForListWorkspacesOutput {
	s.Name = &v
	return s
}

// SetUpdateTime sets the UpdateTime field's value.
func (s *ItemForListWorkspacesOutput) SetUpdateTime(v string) *ItemForListWorkspacesOutput {
	s.UpdateTime = &v
	return s
}

// SetVisibility sets the Visibility field's value.
func (s *ItemForListWorkspacesOutput) SetVisibility(v string) *ItemForListWorkspacesOutput {
	s.Visibility = &v
	return s
}

// SetVisibleUsers sets the VisibleUsers field's value.
func (s *ItemForListWorkspacesOutput) SetVisibleUsers(v []*VisibleUserForListWorkspacesOutput) *ItemForListWorkspacesOutput {
	s.VisibleUsers = v
	return s
}

type ListWorkspacesInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Filter *FilterForListWorkspacesInput `type:"structure" json:",omitempty"`

	PageNumber *int64 `type:"int64" json:",omitempty"`

	PageSize *int64 `type:"int64" json:",omitempty"`
}

// String returns the string representation
func (s ListWorkspacesInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListWorkspacesInput) GoString() string {
	return s.String()
}

// SetFilter sets the Filter field's value.
func (s *ListWorkspacesInput) SetFilter(v *FilterForListWorkspacesInput) *ListWorkspacesInput {
	s.Filter = v
	return s
}

// SetPageNumber sets the PageNumber field's value.
func (s *ListWorkspacesInput) SetPageNumber(v int64) *ListWorkspacesInput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *ListWorkspacesInput) SetPageSize(v int64) *ListWorkspacesInput {
	s.PageSize = &v
	return s
}

type ListWorkspacesOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	Items []*ItemForListWorkspacesOutput `type:"list" json:",omitempty"`

	PageNumber *int64 `type:"int64" json:",omitempty"`

	PageSize *int64 `type:"int64" json:",omitempty"`

	TotalCount *int64 `type:"int64" json:",omitempty"`
}

// String returns the string representation
func (s ListWorkspacesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListWorkspacesOutput) GoString() string {
	return s.String()
}

// SetItems sets the Items field's value.
func (s *ListWorkspacesOutput) SetItems(v []*ItemForListWorkspacesOutput) *ListWorkspacesOutput {
	s.Items = v
	return s
}

// SetPageNumber sets the PageNumber field's value.
func (s *ListWorkspacesOutput) SetPageNumber(v int64) *ListWorkspacesOutput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *ListWorkspacesOutput) SetPageSize(v int64) *ListWorkspacesOutput {
	s.PageSize = &v
	return s
}

// SetTotalCount sets the TotalCount field's value.
func (s *ListWorkspacesOutput) SetTotalCount(v int64) *ListWorkspacesOutput {
	s.TotalCount = &v
	return s
}

type VisibleUserForListWorkspacesOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	AccountId *int64 `type:"int64" json:",omitempty"`

	UserId *int64 `type:"int64" json:",omitempty"`
}

// String returns the string representation
func (s VisibleUserForListWorkspacesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s VisibleUserForListWorkspacesOutput) GoString() string {
	return s.String()
}

// SetAccountId sets the AccountId field's value.
func (s *VisibleUserForListWorkspacesOutput) SetAccountId(v int64) *VisibleUserForListWorkspacesOutput {
	s.AccountId = &v
	return s
}

// SetUserId sets the UserId field's value.
func (s *VisibleUserForListWorkspacesOutput) SetUserId(v int64) *VisibleUserForListWorkspacesOutput {
	s.UserId = &v
	return s
}

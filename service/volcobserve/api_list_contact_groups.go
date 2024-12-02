// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package volcobserve

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opListContactGroupsCommon = "ListContactGroups"

// ListContactGroupsCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ListContactGroupsCommon operation. The "output" return
// value will be populated with the ListContactGroupsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListContactGroupsCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListContactGroupsCommon Send returns without error.
//
// See ListContactGroupsCommon for more information on using the ListContactGroupsCommon
// API call, and error handling.
//
//    // Example sending a request using the ListContactGroupsCommonRequest method.
//    req, resp := client.ListContactGroupsCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VOLCOBSERVE) ListContactGroupsCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opListContactGroupsCommon,
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

// ListContactGroupsCommon API operation for VOLC_OBSERVE.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VOLC_OBSERVE's
// API operation ListContactGroupsCommon for usage and error information.
func (c *VOLCOBSERVE) ListContactGroupsCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ListContactGroupsCommonRequest(input)
	return out, req.Send()
}

// ListContactGroupsCommonWithContext is the same as ListContactGroupsCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ListContactGroupsCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VOLCOBSERVE) ListContactGroupsCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ListContactGroupsCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opListContactGroups = "ListContactGroups"

// ListContactGroupsRequest generates a "volcengine/request.Request" representing the
// client's request for the ListContactGroups operation. The "output" return
// value will be populated with the ListContactGroupsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListContactGroupsCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListContactGroupsCommon Send returns without error.
//
// See ListContactGroups for more information on using the ListContactGroups
// API call, and error handling.
//
//    // Example sending a request using the ListContactGroupsRequest method.
//    req, resp := client.ListContactGroupsRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VOLCOBSERVE) ListContactGroupsRequest(input *ListContactGroupsInput) (req *request.Request, output *ListContactGroupsOutput) {
	op := &request.Operation{
		Name:       opListContactGroups,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListContactGroupsInput{}
	}

	output = &ListContactGroupsOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// ListContactGroups API operation for VOLC_OBSERVE.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VOLC_OBSERVE's
// API operation ListContactGroups for usage and error information.
func (c *VOLCOBSERVE) ListContactGroups(input *ListContactGroupsInput) (*ListContactGroupsOutput, error) {
	req, out := c.ListContactGroupsRequest(input)
	return out, req.Send()
}

// ListContactGroupsWithContext is the same as ListContactGroups with the addition of
// the ability to pass a context and additional request options.
//
// See ListContactGroups for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VOLCOBSERVE) ListContactGroupsWithContext(ctx volcengine.Context, input *ListContactGroupsInput, opts ...request.Option) (*ListContactGroupsOutput, error) {
	req, out := c.ListContactGroupsRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DataForListContactGroupsOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	AccountId *string `type:"string" json:",omitempty"`

	CreatedAt *string `type:"string" json:",omitempty"`

	Description *string `type:"string" json:",omitempty"`

	Id *string `type:"string" json:",omitempty"`

	Name *string `type:"string" json:",omitempty"`

	UpdatedAt *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s DataForListContactGroupsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DataForListContactGroupsOutput) GoString() string {
	return s.String()
}

// SetAccountId sets the AccountId field's value.
func (s *DataForListContactGroupsOutput) SetAccountId(v string) *DataForListContactGroupsOutput {
	s.AccountId = &v
	return s
}

// SetCreatedAt sets the CreatedAt field's value.
func (s *DataForListContactGroupsOutput) SetCreatedAt(v string) *DataForListContactGroupsOutput {
	s.CreatedAt = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *DataForListContactGroupsOutput) SetDescription(v string) *DataForListContactGroupsOutput {
	s.Description = &v
	return s
}

// SetId sets the Id field's value.
func (s *DataForListContactGroupsOutput) SetId(v string) *DataForListContactGroupsOutput {
	s.Id = &v
	return s
}

// SetName sets the Name field's value.
func (s *DataForListContactGroupsOutput) SetName(v string) *DataForListContactGroupsOutput {
	s.Name = &v
	return s
}

// SetUpdatedAt sets the UpdatedAt field's value.
func (s *DataForListContactGroupsOutput) SetUpdatedAt(v string) *DataForListContactGroupsOutput {
	s.UpdatedAt = &v
	return s
}

type ListContactGroupsInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Name *string `type:"string" json:",omitempty"`

	PageNumber *int64 `type:"integer" json:",omitempty"`

	PageSize *int64 `type:"integer" json:",omitempty"`
}

// String returns the string representation
func (s ListContactGroupsInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListContactGroupsInput) GoString() string {
	return s.String()
}

// SetName sets the Name field's value.
func (s *ListContactGroupsInput) SetName(v string) *ListContactGroupsInput {
	s.Name = &v
	return s
}

// SetPageNumber sets the PageNumber field's value.
func (s *ListContactGroupsInput) SetPageNumber(v int64) *ListContactGroupsInput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *ListContactGroupsInput) SetPageSize(v int64) *ListContactGroupsInput {
	s.PageSize = &v
	return s
}

type ListContactGroupsOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	Data []*DataForListContactGroupsOutput `type:"list" json:",omitempty"`

	PageNumber *int64 `type:"integer" json:",omitempty"`

	PageSize *int64 `type:"integer" json:",omitempty"`

	TotalCount *int64 `type:"integer" json:",omitempty"`
}

// String returns the string representation
func (s ListContactGroupsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListContactGroupsOutput) GoString() string {
	return s.String()
}

// SetData sets the Data field's value.
func (s *ListContactGroupsOutput) SetData(v []*DataForListContactGroupsOutput) *ListContactGroupsOutput {
	s.Data = v
	return s
}

// SetPageNumber sets the PageNumber field's value.
func (s *ListContactGroupsOutput) SetPageNumber(v int64) *ListContactGroupsOutput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *ListContactGroupsOutput) SetPageSize(v int64) *ListContactGroupsOutput {
	s.PageSize = &v
	return s
}

// SetTotalCount sets the TotalCount field's value.
func (s *ListContactGroupsOutput) SetTotalCount(v int64) *ListContactGroupsOutput {
	s.TotalCount = &v
	return s
}

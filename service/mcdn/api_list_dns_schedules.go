// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package mcdn

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opListDnsSchedulesCommon = "ListDnsSchedules"

// ListDnsSchedulesCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ListDnsSchedulesCommon operation. The "output" return
// value will be populated with the ListDnsSchedulesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListDnsSchedulesCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListDnsSchedulesCommon Send returns without error.
//
// See ListDnsSchedulesCommon for more information on using the ListDnsSchedulesCommon
// API call, and error handling.
//
//    // Example sending a request using the ListDnsSchedulesCommonRequest method.
//    req, resp := client.ListDnsSchedulesCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *MCDN) ListDnsSchedulesCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opListDnsSchedulesCommon,
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

// ListDnsSchedulesCommon API operation for MCDN.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for MCDN's
// API operation ListDnsSchedulesCommon for usage and error information.
func (c *MCDN) ListDnsSchedulesCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ListDnsSchedulesCommonRequest(input)
	return out, req.Send()
}

// ListDnsSchedulesCommonWithContext is the same as ListDnsSchedulesCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ListDnsSchedulesCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *MCDN) ListDnsSchedulesCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ListDnsSchedulesCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opListDnsSchedules = "ListDnsSchedules"

// ListDnsSchedulesRequest generates a "volcengine/request.Request" representing the
// client's request for the ListDnsSchedules operation. The "output" return
// value will be populated with the ListDnsSchedulesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListDnsSchedulesCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListDnsSchedulesCommon Send returns without error.
//
// See ListDnsSchedules for more information on using the ListDnsSchedules
// API call, and error handling.
//
//    // Example sending a request using the ListDnsSchedulesRequest method.
//    req, resp := client.ListDnsSchedulesRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *MCDN) ListDnsSchedulesRequest(input *ListDnsSchedulesInput) (req *request.Request, output *ListDnsSchedulesOutput) {
	op := &request.Operation{
		Name:       opListDnsSchedules,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListDnsSchedulesInput{}
	}

	output = &ListDnsSchedulesOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// ListDnsSchedules API operation for MCDN.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for MCDN's
// API operation ListDnsSchedules for usage and error information.
func (c *MCDN) ListDnsSchedules(input *ListDnsSchedulesInput) (*ListDnsSchedulesOutput, error) {
	req, out := c.ListDnsSchedulesRequest(input)
	return out, req.Send()
}

// ListDnsSchedulesWithContext is the same as ListDnsSchedules with the addition of
// the ability to pass a context and additional request options.
//
// See ListDnsSchedules for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *MCDN) ListDnsSchedulesWithContext(ctx volcengine.Context, input *ListDnsSchedulesInput, opts ...request.Option) (*ListDnsSchedulesOutput, error) {
	req, out := c.ListDnsSchedulesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DnsScheduleForListDnsSchedulesOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	CloudAccountIds []*string `type:"list" json:",omitempty"`

	CreatedAt *int64 `type:"int64" json:",omitempty"`

	DomainName *string `type:"string" json:",omitempty"`

	Id *string `type:"string" json:",omitempty"`

	Region *string `type:"string" json:",omitempty"`

	ScheduleCname *string `type:"string" json:",omitempty"`

	ScheduleStatus *string `type:"string" json:",omitempty"`

	ScheduleStrategies []*string `type:"list" json:",omitempty"`

	UpdatedAt *int64 `type:"int64" json:",omitempty"`

	Vendors []*string `type:"list" json:",omitempty"`
}

// String returns the string representation
func (s DnsScheduleForListDnsSchedulesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DnsScheduleForListDnsSchedulesOutput) GoString() string {
	return s.String()
}

// SetCloudAccountIds sets the CloudAccountIds field's value.
func (s *DnsScheduleForListDnsSchedulesOutput) SetCloudAccountIds(v []*string) *DnsScheduleForListDnsSchedulesOutput {
	s.CloudAccountIds = v
	return s
}

// SetCreatedAt sets the CreatedAt field's value.
func (s *DnsScheduleForListDnsSchedulesOutput) SetCreatedAt(v int64) *DnsScheduleForListDnsSchedulesOutput {
	s.CreatedAt = &v
	return s
}

// SetDomainName sets the DomainName field's value.
func (s *DnsScheduleForListDnsSchedulesOutput) SetDomainName(v string) *DnsScheduleForListDnsSchedulesOutput {
	s.DomainName = &v
	return s
}

// SetId sets the Id field's value.
func (s *DnsScheduleForListDnsSchedulesOutput) SetId(v string) *DnsScheduleForListDnsSchedulesOutput {
	s.Id = &v
	return s
}

// SetRegion sets the Region field's value.
func (s *DnsScheduleForListDnsSchedulesOutput) SetRegion(v string) *DnsScheduleForListDnsSchedulesOutput {
	s.Region = &v
	return s
}

// SetScheduleCname sets the ScheduleCname field's value.
func (s *DnsScheduleForListDnsSchedulesOutput) SetScheduleCname(v string) *DnsScheduleForListDnsSchedulesOutput {
	s.ScheduleCname = &v
	return s
}

// SetScheduleStatus sets the ScheduleStatus field's value.
func (s *DnsScheduleForListDnsSchedulesOutput) SetScheduleStatus(v string) *DnsScheduleForListDnsSchedulesOutput {
	s.ScheduleStatus = &v
	return s
}

// SetScheduleStrategies sets the ScheduleStrategies field's value.
func (s *DnsScheduleForListDnsSchedulesOutput) SetScheduleStrategies(v []*string) *DnsScheduleForListDnsSchedulesOutput {
	s.ScheduleStrategies = v
	return s
}

// SetUpdatedAt sets the UpdatedAt field's value.
func (s *DnsScheduleForListDnsSchedulesOutput) SetUpdatedAt(v int64) *DnsScheduleForListDnsSchedulesOutput {
	s.UpdatedAt = &v
	return s
}

// SetVendors sets the Vendors field's value.
func (s *DnsScheduleForListDnsSchedulesOutput) SetVendors(v []*string) *DnsScheduleForListDnsSchedulesOutput {
	s.Vendors = v
	return s
}

type ListDnsSchedulesInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	CloudAccountIds []*string `type:"list" json:",omitempty"`

	DomainName *string `type:"string" json:",omitempty"`

	ExactDomainName *string `type:"string" json:",omitempty"`

	Pagination *PaginationForListDnsSchedulesInput `type:"structure" json:",omitempty"`

	Region *string `type:"string" json:",omitempty"`

	ScheduleStatus *string `type:"string" json:",omitempty"`

	ScheduleStrategy *string `type:"string" json:",omitempty"`

	SortBy *string `type:"string" json:",omitempty"`

	SortOrder *string `type:"string" json:",omitempty"`

	Vendors []*string `type:"list" json:",omitempty"`
}

// String returns the string representation
func (s ListDnsSchedulesInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListDnsSchedulesInput) GoString() string {
	return s.String()
}

// SetCloudAccountIds sets the CloudAccountIds field's value.
func (s *ListDnsSchedulesInput) SetCloudAccountIds(v []*string) *ListDnsSchedulesInput {
	s.CloudAccountIds = v
	return s
}

// SetDomainName sets the DomainName field's value.
func (s *ListDnsSchedulesInput) SetDomainName(v string) *ListDnsSchedulesInput {
	s.DomainName = &v
	return s
}

// SetExactDomainName sets the ExactDomainName field's value.
func (s *ListDnsSchedulesInput) SetExactDomainName(v string) *ListDnsSchedulesInput {
	s.ExactDomainName = &v
	return s
}

// SetPagination sets the Pagination field's value.
func (s *ListDnsSchedulesInput) SetPagination(v *PaginationForListDnsSchedulesInput) *ListDnsSchedulesInput {
	s.Pagination = v
	return s
}

// SetRegion sets the Region field's value.
func (s *ListDnsSchedulesInput) SetRegion(v string) *ListDnsSchedulesInput {
	s.Region = &v
	return s
}

// SetScheduleStatus sets the ScheduleStatus field's value.
func (s *ListDnsSchedulesInput) SetScheduleStatus(v string) *ListDnsSchedulesInput {
	s.ScheduleStatus = &v
	return s
}

// SetScheduleStrategy sets the ScheduleStrategy field's value.
func (s *ListDnsSchedulesInput) SetScheduleStrategy(v string) *ListDnsSchedulesInput {
	s.ScheduleStrategy = &v
	return s
}

// SetSortBy sets the SortBy field's value.
func (s *ListDnsSchedulesInput) SetSortBy(v string) *ListDnsSchedulesInput {
	s.SortBy = &v
	return s
}

// SetSortOrder sets the SortOrder field's value.
func (s *ListDnsSchedulesInput) SetSortOrder(v string) *ListDnsSchedulesInput {
	s.SortOrder = &v
	return s
}

// SetVendors sets the Vendors field's value.
func (s *ListDnsSchedulesInput) SetVendors(v []*string) *ListDnsSchedulesInput {
	s.Vendors = v
	return s
}

type ListDnsSchedulesOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	DnsSchedules []*DnsScheduleForListDnsSchedulesOutput `type:"list" json:",omitempty"`

	Pagination *PaginationForListDnsSchedulesOutput `type:"structure" json:",omitempty"`
}

// String returns the string representation
func (s ListDnsSchedulesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListDnsSchedulesOutput) GoString() string {
	return s.String()
}

// SetDnsSchedules sets the DnsSchedules field's value.
func (s *ListDnsSchedulesOutput) SetDnsSchedules(v []*DnsScheduleForListDnsSchedulesOutput) *ListDnsSchedulesOutput {
	s.DnsSchedules = v
	return s
}

// SetPagination sets the Pagination field's value.
func (s *ListDnsSchedulesOutput) SetPagination(v *PaginationForListDnsSchedulesOutput) *ListDnsSchedulesOutput {
	s.Pagination = v
	return s
}

type PaginationForListDnsSchedulesInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	PageNum *int64 `type:"int64" json:",omitempty"`

	PageSize *int64 `type:"int64" json:",omitempty"`
}

// String returns the string representation
func (s PaginationForListDnsSchedulesInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s PaginationForListDnsSchedulesInput) GoString() string {
	return s.String()
}

// SetPageNum sets the PageNum field's value.
func (s *PaginationForListDnsSchedulesInput) SetPageNum(v int64) *PaginationForListDnsSchedulesInput {
	s.PageNum = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *PaginationForListDnsSchedulesInput) SetPageSize(v int64) *PaginationForListDnsSchedulesInput {
	s.PageSize = &v
	return s
}

type PaginationForListDnsSchedulesOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	PageNum *int64 `type:"int64" json:",omitempty"`

	PageSize *int64 `type:"int64" json:",omitempty"`

	Total *int64 `type:"int64" json:",omitempty"`
}

// String returns the string representation
func (s PaginationForListDnsSchedulesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s PaginationForListDnsSchedulesOutput) GoString() string {
	return s.String()
}

// SetPageNum sets the PageNum field's value.
func (s *PaginationForListDnsSchedulesOutput) SetPageNum(v int64) *PaginationForListDnsSchedulesOutput {
	s.PageNum = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *PaginationForListDnsSchedulesOutput) SetPageSize(v int64) *PaginationForListDnsSchedulesOutput {
	s.PageSize = &v
	return s
}

// SetTotal sets the Total field's value.
func (s *PaginationForListDnsSchedulesOutput) SetTotal(v int64) *PaginationForListDnsSchedulesOutput {
	s.Total = &v
	return s
}
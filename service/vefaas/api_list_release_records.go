// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package vefaas

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opListReleaseRecordsCommon = "ListReleaseRecords"

// ListReleaseRecordsCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ListReleaseRecordsCommon operation. The "output" return
// value will be populated with the ListReleaseRecordsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListReleaseRecordsCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListReleaseRecordsCommon Send returns without error.
//
// See ListReleaseRecordsCommon for more information on using the ListReleaseRecordsCommon
// API call, and error handling.
//
//    // Example sending a request using the ListReleaseRecordsCommonRequest method.
//    req, resp := client.ListReleaseRecordsCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VEFAAS) ListReleaseRecordsCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opListReleaseRecordsCommon,
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

// ListReleaseRecordsCommon API operation for VEFAAS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VEFAAS's
// API operation ListReleaseRecordsCommon for usage and error information.
func (c *VEFAAS) ListReleaseRecordsCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ListReleaseRecordsCommonRequest(input)
	return out, req.Send()
}

// ListReleaseRecordsCommonWithContext is the same as ListReleaseRecordsCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ListReleaseRecordsCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VEFAAS) ListReleaseRecordsCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ListReleaseRecordsCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opListReleaseRecords = "ListReleaseRecords"

// ListReleaseRecordsRequest generates a "volcengine/request.Request" representing the
// client's request for the ListReleaseRecords operation. The "output" return
// value will be populated with the ListReleaseRecordsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListReleaseRecordsCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListReleaseRecordsCommon Send returns without error.
//
// See ListReleaseRecords for more information on using the ListReleaseRecords
// API call, and error handling.
//
//    // Example sending a request using the ListReleaseRecordsRequest method.
//    req, resp := client.ListReleaseRecordsRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VEFAAS) ListReleaseRecordsRequest(input *ListReleaseRecordsInput) (req *request.Request, output *ListReleaseRecordsOutput) {
	op := &request.Operation{
		Name:       opListReleaseRecords,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListReleaseRecordsInput{}
	}

	output = &ListReleaseRecordsOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// ListReleaseRecords API operation for VEFAAS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VEFAAS's
// API operation ListReleaseRecords for usage and error information.
func (c *VEFAAS) ListReleaseRecords(input *ListReleaseRecordsInput) (*ListReleaseRecordsOutput, error) {
	req, out := c.ListReleaseRecordsRequest(input)
	return out, req.Send()
}

// ListReleaseRecordsWithContext is the same as ListReleaseRecords with the addition of
// the ability to pass a context and additional request options.
//
// See ListReleaseRecords for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VEFAAS) ListReleaseRecordsWithContext(ctx volcengine.Context, input *ListReleaseRecordsInput, opts ...request.Option) (*ListReleaseRecordsOutput, error) {
	req, out := c.ListReleaseRecordsRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type FilterForListReleaseRecordsInput struct {
	_ struct{} `type:"structure"`

	Name *string `type:"string"`

	Values []*string `type:"list"`
}

// String returns the string representation
func (s FilterForListReleaseRecordsInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s FilterForListReleaseRecordsInput) GoString() string {
	return s.String()
}

// SetName sets the Name field's value.
func (s *FilterForListReleaseRecordsInput) SetName(v string) *FilterForListReleaseRecordsInput {
	s.Name = &v
	return s
}

// SetValues sets the Values field's value.
func (s *FilterForListReleaseRecordsInput) SetValues(v []*string) *FilterForListReleaseRecordsInput {
	s.Values = v
	return s
}

type ItemForListReleaseRecordsOutput struct {
	_ struct{} `type:"structure"`

	CreationTime *string `type:"string"`

	Description *string `type:"string"`

	FinishTime *string `type:"string"`

	FunctionId *string `type:"string"`

	Id *string `type:"string"`

	LastUpdateTime *string `type:"string"`

	SourceRevisionNumber *int32 `type:"int32"`

	Status *string `type:"string"`

	TargetRevisionNumber *int32 `type:"int32"`
}

// String returns the string representation
func (s ItemForListReleaseRecordsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ItemForListReleaseRecordsOutput) GoString() string {
	return s.String()
}

// SetCreationTime sets the CreationTime field's value.
func (s *ItemForListReleaseRecordsOutput) SetCreationTime(v string) *ItemForListReleaseRecordsOutput {
	s.CreationTime = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *ItemForListReleaseRecordsOutput) SetDescription(v string) *ItemForListReleaseRecordsOutput {
	s.Description = &v
	return s
}

// SetFinishTime sets the FinishTime field's value.
func (s *ItemForListReleaseRecordsOutput) SetFinishTime(v string) *ItemForListReleaseRecordsOutput {
	s.FinishTime = &v
	return s
}

// SetFunctionId sets the FunctionId field's value.
func (s *ItemForListReleaseRecordsOutput) SetFunctionId(v string) *ItemForListReleaseRecordsOutput {
	s.FunctionId = &v
	return s
}

// SetId sets the Id field's value.
func (s *ItemForListReleaseRecordsOutput) SetId(v string) *ItemForListReleaseRecordsOutput {
	s.Id = &v
	return s
}

// SetLastUpdateTime sets the LastUpdateTime field's value.
func (s *ItemForListReleaseRecordsOutput) SetLastUpdateTime(v string) *ItemForListReleaseRecordsOutput {
	s.LastUpdateTime = &v
	return s
}

// SetSourceRevisionNumber sets the SourceRevisionNumber field's value.
func (s *ItemForListReleaseRecordsOutput) SetSourceRevisionNumber(v int32) *ItemForListReleaseRecordsOutput {
	s.SourceRevisionNumber = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *ItemForListReleaseRecordsOutput) SetStatus(v string) *ItemForListReleaseRecordsOutput {
	s.Status = &v
	return s
}

// SetTargetRevisionNumber sets the TargetRevisionNumber field's value.
func (s *ItemForListReleaseRecordsOutput) SetTargetRevisionNumber(v int32) *ItemForListReleaseRecordsOutput {
	s.TargetRevisionNumber = &v
	return s
}

type ListReleaseRecordsInput struct {
	_ struct{} `type:"structure"`

	Filters []*FilterForListReleaseRecordsInput `type:"list"`

	// FunctionId is a required field
	FunctionId *string `type:"string" required:"true"`

	OrderBy *OrderByForListReleaseRecordsInput `type:"structure"`

	PageNumber *int32 `type:"int32"`

	PageSize *int32 `type:"int32"`
}

// String returns the string representation
func (s ListReleaseRecordsInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListReleaseRecordsInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListReleaseRecordsInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ListReleaseRecordsInput"}
	if s.FunctionId == nil {
		invalidParams.Add(request.NewErrParamRequired("FunctionId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetFilters sets the Filters field's value.
func (s *ListReleaseRecordsInput) SetFilters(v []*FilterForListReleaseRecordsInput) *ListReleaseRecordsInput {
	s.Filters = v
	return s
}

// SetFunctionId sets the FunctionId field's value.
func (s *ListReleaseRecordsInput) SetFunctionId(v string) *ListReleaseRecordsInput {
	s.FunctionId = &v
	return s
}

// SetOrderBy sets the OrderBy field's value.
func (s *ListReleaseRecordsInput) SetOrderBy(v *OrderByForListReleaseRecordsInput) *ListReleaseRecordsInput {
	s.OrderBy = v
	return s
}

// SetPageNumber sets the PageNumber field's value.
func (s *ListReleaseRecordsInput) SetPageNumber(v int32) *ListReleaseRecordsInput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *ListReleaseRecordsInput) SetPageSize(v int32) *ListReleaseRecordsInput {
	s.PageSize = &v
	return s
}

type ListReleaseRecordsOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	Items []*ItemForListReleaseRecordsOutput `type:"list"`

	Total *int32 `type:"int32"`
}

// String returns the string representation
func (s ListReleaseRecordsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListReleaseRecordsOutput) GoString() string {
	return s.String()
}

// SetItems sets the Items field's value.
func (s *ListReleaseRecordsOutput) SetItems(v []*ItemForListReleaseRecordsOutput) *ListReleaseRecordsOutput {
	s.Items = v
	return s
}

// SetTotal sets the Total field's value.
func (s *ListReleaseRecordsOutput) SetTotal(v int32) *ListReleaseRecordsOutput {
	s.Total = &v
	return s
}

type OrderByForListReleaseRecordsInput struct {
	_ struct{} `type:"structure"`

	Ascend *bool `type:"boolean"`

	Key *string `type:"string"`
}

// String returns the string representation
func (s OrderByForListReleaseRecordsInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s OrderByForListReleaseRecordsInput) GoString() string {
	return s.String()
}

// SetAscend sets the Ascend field's value.
func (s *OrderByForListReleaseRecordsInput) SetAscend(v bool) *OrderByForListReleaseRecordsInput {
	s.Ascend = &v
	return s
}

// SetKey sets the Key field's value.
func (s *OrderByForListReleaseRecordsInput) SetKey(v string) *OrderByForListReleaseRecordsInput {
	s.Key = &v
	return s
}
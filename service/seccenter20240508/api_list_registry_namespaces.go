// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package seccenter20240508

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opListRegistryNamespacesCommon = "ListRegistryNamespaces"

// ListRegistryNamespacesCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ListRegistryNamespacesCommon operation. The "output" return
// value will be populated with the ListRegistryNamespacesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListRegistryNamespacesCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListRegistryNamespacesCommon Send returns without error.
//
// See ListRegistryNamespacesCommon for more information on using the ListRegistryNamespacesCommon
// API call, and error handling.
//
//    // Example sending a request using the ListRegistryNamespacesCommonRequest method.
//    req, resp := client.ListRegistryNamespacesCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *SECCENTER20240508) ListRegistryNamespacesCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opListRegistryNamespacesCommon,
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

// ListRegistryNamespacesCommon API operation for SECCENTER20240508.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for SECCENTER20240508's
// API operation ListRegistryNamespacesCommon for usage and error information.
func (c *SECCENTER20240508) ListRegistryNamespacesCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ListRegistryNamespacesCommonRequest(input)
	return out, req.Send()
}

// ListRegistryNamespacesCommonWithContext is the same as ListRegistryNamespacesCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ListRegistryNamespacesCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *SECCENTER20240508) ListRegistryNamespacesCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ListRegistryNamespacesCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opListRegistryNamespaces = "ListRegistryNamespaces"

// ListRegistryNamespacesRequest generates a "volcengine/request.Request" representing the
// client's request for the ListRegistryNamespaces operation. The "output" return
// value will be populated with the ListRegistryNamespacesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListRegistryNamespacesCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListRegistryNamespacesCommon Send returns without error.
//
// See ListRegistryNamespaces for more information on using the ListRegistryNamespaces
// API call, and error handling.
//
//    // Example sending a request using the ListRegistryNamespacesRequest method.
//    req, resp := client.ListRegistryNamespacesRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *SECCENTER20240508) ListRegistryNamespacesRequest(input *ListRegistryNamespacesInput) (req *request.Request, output *ListRegistryNamespacesOutput) {
	op := &request.Operation{
		Name:       opListRegistryNamespaces,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListRegistryNamespacesInput{}
	}

	output = &ListRegistryNamespacesOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// ListRegistryNamespaces API operation for SECCENTER20240508.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for SECCENTER20240508's
// API operation ListRegistryNamespaces for usage and error information.
func (c *SECCENTER20240508) ListRegistryNamespaces(input *ListRegistryNamespacesInput) (*ListRegistryNamespacesOutput, error) {
	req, out := c.ListRegistryNamespacesRequest(input)
	return out, req.Send()
}

// ListRegistryNamespacesWithContext is the same as ListRegistryNamespaces with the addition of
// the ability to pass a context and additional request options.
//
// See ListRegistryNamespaces for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *SECCENTER20240508) ListRegistryNamespacesWithContext(ctx volcengine.Context, input *ListRegistryNamespacesInput, opts ...request.Option) (*ListRegistryNamespacesOutput, error) {
	req, out := c.ListRegistryNamespacesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DataForListRegistryNamespacesOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	ID *string `type:"string" json:",omitempty"`

	Namespace *string `type:"string" json:",omitempty"`

	NamespaceCreateTime *int32 `type:"int32" json:",omitempty"`

	Region *string `type:"string" json:",omitempty"`

	RegistryName *string `type:"string" json:",omitempty"`

	RegistryType *string `type:"string" json:",omitempty"`

	TotalImageCount *int32 `type:"int32" json:",omitempty"`
}

// String returns the string representation
func (s DataForListRegistryNamespacesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DataForListRegistryNamespacesOutput) GoString() string {
	return s.String()
}

// SetID sets the ID field's value.
func (s *DataForListRegistryNamespacesOutput) SetID(v string) *DataForListRegistryNamespacesOutput {
	s.ID = &v
	return s
}

// SetNamespace sets the Namespace field's value.
func (s *DataForListRegistryNamespacesOutput) SetNamespace(v string) *DataForListRegistryNamespacesOutput {
	s.Namespace = &v
	return s
}

// SetNamespaceCreateTime sets the NamespaceCreateTime field's value.
func (s *DataForListRegistryNamespacesOutput) SetNamespaceCreateTime(v int32) *DataForListRegistryNamespacesOutput {
	s.NamespaceCreateTime = &v
	return s
}

// SetRegion sets the Region field's value.
func (s *DataForListRegistryNamespacesOutput) SetRegion(v string) *DataForListRegistryNamespacesOutput {
	s.Region = &v
	return s
}

// SetRegistryName sets the RegistryName field's value.
func (s *DataForListRegistryNamespacesOutput) SetRegistryName(v string) *DataForListRegistryNamespacesOutput {
	s.RegistryName = &v
	return s
}

// SetRegistryType sets the RegistryType field's value.
func (s *DataForListRegistryNamespacesOutput) SetRegistryType(v string) *DataForListRegistryNamespacesOutput {
	s.RegistryType = &v
	return s
}

// SetTotalImageCount sets the TotalImageCount field's value.
func (s *DataForListRegistryNamespacesOutput) SetTotalImageCount(v int32) *DataForListRegistryNamespacesOutput {
	s.TotalImageCount = &v
	return s
}

type FilterForListRegistryNamespacesInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Namespace *string `type:"string" json:",omitempty"`

	NamespaceIDs []*string `type:"list" json:",omitempty"`

	RegistryName *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s FilterForListRegistryNamespacesInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s FilterForListRegistryNamespacesInput) GoString() string {
	return s.String()
}

// SetNamespace sets the Namespace field's value.
func (s *FilterForListRegistryNamespacesInput) SetNamespace(v string) *FilterForListRegistryNamespacesInput {
	s.Namespace = &v
	return s
}

// SetNamespaceIDs sets the NamespaceIDs field's value.
func (s *FilterForListRegistryNamespacesInput) SetNamespaceIDs(v []*string) *FilterForListRegistryNamespacesInput {
	s.NamespaceIDs = v
	return s
}

// SetRegistryName sets the RegistryName field's value.
func (s *FilterForListRegistryNamespacesInput) SetRegistryName(v string) *FilterForListRegistryNamespacesInput {
	s.RegistryName = &v
	return s
}

type ListRegistryNamespacesInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Filter *FilterForListRegistryNamespacesInput `type:"structure" json:",omitempty"`

	// PageNumber is a required field
	PageNumber *int32 `type:"int32" json:",omitempty" required:"true"`

	// PageSize is a required field
	PageSize *int32 `type:"int32" json:",omitempty" required:"true"`

	SortBy *string `type:"string" json:",omitempty"`

	SortOrder *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s ListRegistryNamespacesInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListRegistryNamespacesInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListRegistryNamespacesInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ListRegistryNamespacesInput"}
	if s.PageNumber == nil {
		invalidParams.Add(request.NewErrParamRequired("PageNumber"))
	}
	if s.PageSize == nil {
		invalidParams.Add(request.NewErrParamRequired("PageSize"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetFilter sets the Filter field's value.
func (s *ListRegistryNamespacesInput) SetFilter(v *FilterForListRegistryNamespacesInput) *ListRegistryNamespacesInput {
	s.Filter = v
	return s
}

// SetPageNumber sets the PageNumber field's value.
func (s *ListRegistryNamespacesInput) SetPageNumber(v int32) *ListRegistryNamespacesInput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *ListRegistryNamespacesInput) SetPageSize(v int32) *ListRegistryNamespacesInput {
	s.PageSize = &v
	return s
}

// SetSortBy sets the SortBy field's value.
func (s *ListRegistryNamespacesInput) SetSortBy(v string) *ListRegistryNamespacesInput {
	s.SortBy = &v
	return s
}

// SetSortOrder sets the SortOrder field's value.
func (s *ListRegistryNamespacesInput) SetSortOrder(v string) *ListRegistryNamespacesInput {
	s.SortOrder = &v
	return s
}

type ListRegistryNamespacesOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	Data []*DataForListRegistryNamespacesOutput `type:"list" json:",omitempty"`

	PageNumber *int32 `type:"int32" json:",omitempty"`

	PageSize *int32 `type:"int32" json:",omitempty"`

	TotalCount *int32 `type:"int32" json:",omitempty"`
}

// String returns the string representation
func (s ListRegistryNamespacesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListRegistryNamespacesOutput) GoString() string {
	return s.String()
}

// SetData sets the Data field's value.
func (s *ListRegistryNamespacesOutput) SetData(v []*DataForListRegistryNamespacesOutput) *ListRegistryNamespacesOutput {
	s.Data = v
	return s
}

// SetPageNumber sets the PageNumber field's value.
func (s *ListRegistryNamespacesOutput) SetPageNumber(v int32) *ListRegistryNamespacesOutput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *ListRegistryNamespacesOutput) SetPageSize(v int32) *ListRegistryNamespacesOutput {
	s.PageSize = &v
	return s
}

// SetTotalCount sets the TotalCount field's value.
func (s *ListRegistryNamespacesOutput) SetTotalCount(v int32) *ListRegistryNamespacesOutput {
	s.TotalCount = &v
	return s
}

// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package cr

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opListDomainsCommon = "ListDomains"

// ListDomainsCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ListDomainsCommon operation. The "output" return
// value will be populated with the ListDomainsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListDomainsCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListDomainsCommon Send returns without error.
//
// See ListDomainsCommon for more information on using the ListDomainsCommon
// API call, and error handling.
//
//    // Example sending a request using the ListDomainsCommonRequest method.
//    req, resp := client.ListDomainsCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CR) ListDomainsCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opListDomainsCommon,
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

// ListDomainsCommon API operation for CR.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CR's
// API operation ListDomainsCommon for usage and error information.
func (c *CR) ListDomainsCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ListDomainsCommonRequest(input)
	return out, req.Send()
}

// ListDomainsCommonWithContext is the same as ListDomainsCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ListDomainsCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CR) ListDomainsCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ListDomainsCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opListDomains = "ListDomains"

// ListDomainsRequest generates a "volcengine/request.Request" representing the
// client's request for the ListDomains operation. The "output" return
// value will be populated with the ListDomainsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListDomainsCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListDomainsCommon Send returns without error.
//
// See ListDomains for more information on using the ListDomains
// API call, and error handling.
//
//    // Example sending a request using the ListDomainsRequest method.
//    req, resp := client.ListDomainsRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CR) ListDomainsRequest(input *ListDomainsInput) (req *request.Request, output *ListDomainsOutput) {
	op := &request.Operation{
		Name:       opListDomains,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListDomainsInput{}
	}

	output = &ListDomainsOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// ListDomains API operation for CR.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CR's
// API operation ListDomains for usage and error information.
func (c *CR) ListDomains(input *ListDomainsInput) (*ListDomainsOutput, error) {
	req, out := c.ListDomainsRequest(input)
	return out, req.Send()
}

// ListDomainsWithContext is the same as ListDomains with the addition of
// the ability to pass a context and additional request options.
//
// See ListDomains for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CR) ListDomainsWithContext(ctx volcengine.Context, input *ListDomainsInput, opts ...request.Option) (*ListDomainsOutput, error) {
	req, out := c.ListDomainsRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ItemForListDomainsOutput struct {
	_ struct{} `type:"structure"`

	CreateTime *string `type:"string"`

	Default *bool `type:"boolean"`

	Domain *string `type:"string"`

	Type *string `type:"string"`
}

// String returns the string representation
func (s ItemForListDomainsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ItemForListDomainsOutput) GoString() string {
	return s.String()
}

// SetCreateTime sets the CreateTime field's value.
func (s *ItemForListDomainsOutput) SetCreateTime(v string) *ItemForListDomainsOutput {
	s.CreateTime = &v
	return s
}

// SetDefault sets the Default field's value.
func (s *ItemForListDomainsOutput) SetDefault(v bool) *ItemForListDomainsOutput {
	s.Default = &v
	return s
}

// SetDomain sets the Domain field's value.
func (s *ItemForListDomainsOutput) SetDomain(v string) *ItemForListDomainsOutput {
	s.Domain = &v
	return s
}

// SetType sets the Type field's value.
func (s *ItemForListDomainsOutput) SetType(v string) *ItemForListDomainsOutput {
	s.Type = &v
	return s
}

type ListDomainsInput struct {
	_ struct{} `type:"structure"`

	PageNumber *int64 `type:"int64"`

	PageSize *int64 `type:"int64"`

	// Registry is a required field
	Registry *string `type:"string" required:"true"`
}

// String returns the string representation
func (s ListDomainsInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListDomainsInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListDomainsInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ListDomainsInput"}
	if s.Registry == nil {
		invalidParams.Add(request.NewErrParamRequired("Registry"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetPageNumber sets the PageNumber field's value.
func (s *ListDomainsInput) SetPageNumber(v int64) *ListDomainsInput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *ListDomainsInput) SetPageSize(v int64) *ListDomainsInput {
	s.PageSize = &v
	return s
}

// SetRegistry sets the Registry field's value.
func (s *ListDomainsInput) SetRegistry(v string) *ListDomainsInput {
	s.Registry = &v
	return s
}

type ListDomainsOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	Items []*ItemForListDomainsOutput `type:"list"`

	PageNumber *int64 `type:"int64"`

	PageSize *int64 `type:"int64"`

	Registry *string `type:"string"`

	TotalCount *int64 `type:"int64"`
}

// String returns the string representation
func (s ListDomainsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListDomainsOutput) GoString() string {
	return s.String()
}

// SetItems sets the Items field's value.
func (s *ListDomainsOutput) SetItems(v []*ItemForListDomainsOutput) *ListDomainsOutput {
	s.Items = v
	return s
}

// SetPageNumber sets the PageNumber field's value.
func (s *ListDomainsOutput) SetPageNumber(v int64) *ListDomainsOutput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *ListDomainsOutput) SetPageSize(v int64) *ListDomainsOutput {
	s.PageSize = &v
	return s
}

// SetRegistry sets the Registry field's value.
func (s *ListDomainsOutput) SetRegistry(v string) *ListDomainsOutput {
	s.Registry = &v
	return s
}

// SetTotalCount sets the TotalCount field's value.
func (s *ListDomainsOutput) SetTotalCount(v int64) *ListDomainsOutput {
	s.TotalCount = &v
	return s
}

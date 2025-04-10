// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package edx

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opListCrossAccountVIFAuthorityCommon = "ListCrossAccountVIFAuthority"

// ListCrossAccountVIFAuthorityCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ListCrossAccountVIFAuthorityCommon operation. The "output" return
// value will be populated with the ListCrossAccountVIFAuthorityCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListCrossAccountVIFAuthorityCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListCrossAccountVIFAuthorityCommon Send returns without error.
//
// See ListCrossAccountVIFAuthorityCommon for more information on using the ListCrossAccountVIFAuthorityCommon
// API call, and error handling.
//
//    // Example sending a request using the ListCrossAccountVIFAuthorityCommonRequest method.
//    req, resp := client.ListCrossAccountVIFAuthorityCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *EDX) ListCrossAccountVIFAuthorityCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opListCrossAccountVIFAuthorityCommon,
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

// ListCrossAccountVIFAuthorityCommon API operation for EDX.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for EDX's
// API operation ListCrossAccountVIFAuthorityCommon for usage and error information.
func (c *EDX) ListCrossAccountVIFAuthorityCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ListCrossAccountVIFAuthorityCommonRequest(input)
	return out, req.Send()
}

// ListCrossAccountVIFAuthorityCommonWithContext is the same as ListCrossAccountVIFAuthorityCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ListCrossAccountVIFAuthorityCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *EDX) ListCrossAccountVIFAuthorityCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ListCrossAccountVIFAuthorityCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opListCrossAccountVIFAuthority = "ListCrossAccountVIFAuthority"

// ListCrossAccountVIFAuthorityRequest generates a "volcengine/request.Request" representing the
// client's request for the ListCrossAccountVIFAuthority operation. The "output" return
// value will be populated with the ListCrossAccountVIFAuthorityCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListCrossAccountVIFAuthorityCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListCrossAccountVIFAuthorityCommon Send returns without error.
//
// See ListCrossAccountVIFAuthority for more information on using the ListCrossAccountVIFAuthority
// API call, and error handling.
//
//    // Example sending a request using the ListCrossAccountVIFAuthorityRequest method.
//    req, resp := client.ListCrossAccountVIFAuthorityRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *EDX) ListCrossAccountVIFAuthorityRequest(input *ListCrossAccountVIFAuthorityInput) (req *request.Request, output *ListCrossAccountVIFAuthorityOutput) {
	op := &request.Operation{
		Name:       opListCrossAccountVIFAuthority,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListCrossAccountVIFAuthorityInput{}
	}

	output = &ListCrossAccountVIFAuthorityOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// ListCrossAccountVIFAuthority API operation for EDX.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for EDX's
// API operation ListCrossAccountVIFAuthority for usage and error information.
func (c *EDX) ListCrossAccountVIFAuthority(input *ListCrossAccountVIFAuthorityInput) (*ListCrossAccountVIFAuthorityOutput, error) {
	req, out := c.ListCrossAccountVIFAuthorityRequest(input)
	return out, req.Send()
}

// ListCrossAccountVIFAuthorityWithContext is the same as ListCrossAccountVIFAuthority with the addition of
// the ability to pass a context and additional request options.
//
// See ListCrossAccountVIFAuthority for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *EDX) ListCrossAccountVIFAuthorityWithContext(ctx volcengine.Context, input *ListCrossAccountVIFAuthorityInput, opts ...request.Option) (*ListCrossAccountVIFAuthorityOutput, error) {
	req, out := c.ListCrossAccountVIFAuthorityRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type AuthorityListForListCrossAccountVIFAuthorityOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	CreateTime *string `type:"string" json:",omitempty"`

	State *string `type:"string" json:",omitempty"`

	VGWAccountId *int32 `type:"int32" json:",omitempty"`

	VIFAccountId *int32 `type:"int32" json:",omitempty"`

	VIFInstanceId *string `type:"string" json:",omitempty"`

	VIFInstanceName *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s AuthorityListForListCrossAccountVIFAuthorityOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s AuthorityListForListCrossAccountVIFAuthorityOutput) GoString() string {
	return s.String()
}

// SetCreateTime sets the CreateTime field's value.
func (s *AuthorityListForListCrossAccountVIFAuthorityOutput) SetCreateTime(v string) *AuthorityListForListCrossAccountVIFAuthorityOutput {
	s.CreateTime = &v
	return s
}

// SetState sets the State field's value.
func (s *AuthorityListForListCrossAccountVIFAuthorityOutput) SetState(v string) *AuthorityListForListCrossAccountVIFAuthorityOutput {
	s.State = &v
	return s
}

// SetVGWAccountId sets the VGWAccountId field's value.
func (s *AuthorityListForListCrossAccountVIFAuthorityOutput) SetVGWAccountId(v int32) *AuthorityListForListCrossAccountVIFAuthorityOutput {
	s.VGWAccountId = &v
	return s
}

// SetVIFAccountId sets the VIFAccountId field's value.
func (s *AuthorityListForListCrossAccountVIFAuthorityOutput) SetVIFAccountId(v int32) *AuthorityListForListCrossAccountVIFAuthorityOutput {
	s.VIFAccountId = &v
	return s
}

// SetVIFInstanceId sets the VIFInstanceId field's value.
func (s *AuthorityListForListCrossAccountVIFAuthorityOutput) SetVIFInstanceId(v string) *AuthorityListForListCrossAccountVIFAuthorityOutput {
	s.VIFInstanceId = &v
	return s
}

// SetVIFInstanceName sets the VIFInstanceName field's value.
func (s *AuthorityListForListCrossAccountVIFAuthorityOutput) SetVIFInstanceName(v string) *AuthorityListForListCrossAccountVIFAuthorityOutput {
	s.VIFInstanceName = &v
	return s
}

type ListCrossAccountVIFAuthorityInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	PageNumber *int32 `type:"int32" json:",omitempty"`

	PageSize *int32 `type:"int32" json:",omitempty"`

	State *string `type:"string" json:",omitempty"`

	VGWAccountID *int32 `type:"int32" json:",omitempty"`

	VIFInstanceID *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s ListCrossAccountVIFAuthorityInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListCrossAccountVIFAuthorityInput) GoString() string {
	return s.String()
}

// SetPageNumber sets the PageNumber field's value.
func (s *ListCrossAccountVIFAuthorityInput) SetPageNumber(v int32) *ListCrossAccountVIFAuthorityInput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *ListCrossAccountVIFAuthorityInput) SetPageSize(v int32) *ListCrossAccountVIFAuthorityInput {
	s.PageSize = &v
	return s
}

// SetState sets the State field's value.
func (s *ListCrossAccountVIFAuthorityInput) SetState(v string) *ListCrossAccountVIFAuthorityInput {
	s.State = &v
	return s
}

// SetVGWAccountID sets the VGWAccountID field's value.
func (s *ListCrossAccountVIFAuthorityInput) SetVGWAccountID(v int32) *ListCrossAccountVIFAuthorityInput {
	s.VGWAccountID = &v
	return s
}

// SetVIFInstanceID sets the VIFInstanceID field's value.
func (s *ListCrossAccountVIFAuthorityInput) SetVIFInstanceID(v string) *ListCrossAccountVIFAuthorityInput {
	s.VIFInstanceID = &v
	return s
}

type ListCrossAccountVIFAuthorityOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	AuthorityList []*AuthorityListForListCrossAccountVIFAuthorityOutput `type:"list" json:",omitempty"`

	AuthorizedCount *int32 `type:"int32" json:",omitempty"`

	AuthorizingCount *int32 `type:"int32" json:",omitempty"`

	CanceledCount *int32 `type:"int32" json:",omitempty"`

	PageNumber *int32 `type:"int32" json:",omitempty"`

	PageSize *int32 `type:"int32" json:",omitempty"`

	RejectCount *int32 `type:"int32" json:",omitempty"`

	TotalCount *int32 `type:"int32" json:",omitempty"`
}

// String returns the string representation
func (s ListCrossAccountVIFAuthorityOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListCrossAccountVIFAuthorityOutput) GoString() string {
	return s.String()
}

// SetAuthorityList sets the AuthorityList field's value.
func (s *ListCrossAccountVIFAuthorityOutput) SetAuthorityList(v []*AuthorityListForListCrossAccountVIFAuthorityOutput) *ListCrossAccountVIFAuthorityOutput {
	s.AuthorityList = v
	return s
}

// SetAuthorizedCount sets the AuthorizedCount field's value.
func (s *ListCrossAccountVIFAuthorityOutput) SetAuthorizedCount(v int32) *ListCrossAccountVIFAuthorityOutput {
	s.AuthorizedCount = &v
	return s
}

// SetAuthorizingCount sets the AuthorizingCount field's value.
func (s *ListCrossAccountVIFAuthorityOutput) SetAuthorizingCount(v int32) *ListCrossAccountVIFAuthorityOutput {
	s.AuthorizingCount = &v
	return s
}

// SetCanceledCount sets the CanceledCount field's value.
func (s *ListCrossAccountVIFAuthorityOutput) SetCanceledCount(v int32) *ListCrossAccountVIFAuthorityOutput {
	s.CanceledCount = &v
	return s
}

// SetPageNumber sets the PageNumber field's value.
func (s *ListCrossAccountVIFAuthorityOutput) SetPageNumber(v int32) *ListCrossAccountVIFAuthorityOutput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *ListCrossAccountVIFAuthorityOutput) SetPageSize(v int32) *ListCrossAccountVIFAuthorityOutput {
	s.PageSize = &v
	return s
}

// SetRejectCount sets the RejectCount field's value.
func (s *ListCrossAccountVIFAuthorityOutput) SetRejectCount(v int32) *ListCrossAccountVIFAuthorityOutput {
	s.RejectCount = &v
	return s
}

// SetTotalCount sets the TotalCount field's value.
func (s *ListCrossAccountVIFAuthorityOutput) SetTotalCount(v int32) *ListCrossAccountVIFAuthorityOutput {
	s.TotalCount = &v
	return s
}

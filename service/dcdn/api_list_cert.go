// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package dcdn

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opListCertCommon = "ListCert"

// ListCertCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ListCertCommon operation. The "output" return
// value will be populated with the ListCertCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListCertCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListCertCommon Send returns without error.
//
// See ListCertCommon for more information on using the ListCertCommon
// API call, and error handling.
//
//    // Example sending a request using the ListCertCommonRequest method.
//    req, resp := client.ListCertCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *DCDN) ListCertCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opListCertCommon,
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

// ListCertCommon API operation for DCDN.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for DCDN's
// API operation ListCertCommon for usage and error information.
func (c *DCDN) ListCertCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ListCertCommonRequest(input)
	return out, req.Send()
}

// ListCertCommonWithContext is the same as ListCertCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ListCertCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *DCDN) ListCertCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ListCertCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opListCert = "ListCert"

// ListCertRequest generates a "volcengine/request.Request" representing the
// client's request for the ListCert operation. The "output" return
// value will be populated with the ListCertCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListCertCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListCertCommon Send returns without error.
//
// See ListCert for more information on using the ListCert
// API call, and error handling.
//
//    // Example sending a request using the ListCertRequest method.
//    req, resp := client.ListCertRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *DCDN) ListCertRequest(input *ListCertInput) (req *request.Request, output *ListCertOutput) {
	op := &request.Operation{
		Name:       opListCert,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListCertInput{}
	}

	output = &ListCertOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// ListCert API operation for DCDN.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for DCDN's
// API operation ListCert for usage and error information.
func (c *DCDN) ListCert(input *ListCertInput) (*ListCertOutput, error) {
	req, out := c.ListCertRequest(input)
	return out, req.Send()
}

// ListCertWithContext is the same as ListCert with the addition of
// the ability to pass a context and additional request options.
//
// See ListCert for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *DCDN) ListCertWithContext(ctx volcengine.Context, input *ListCertInput, opts ...request.Option) (*ListCertOutput, error) {
	req, out := c.ListCertRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type CertListForListCertOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	CertId *string `type:"string" json:",omitempty"`

	CertName *string `type:"string" json:",omitempty"`

	CertSource *string `type:"string" json:",omitempty"`

	CertStatus *string `type:"string" json:",omitempty"`

	Domain []*string `type:"list" json:",omitempty"`

	Expire *string `type:"string" json:",omitempty"`

	KeyType *string `type:"string" json:",omitempty"`

	MatchDomain []*string `type:"list" json:",omitempty"`
}

// String returns the string representation
func (s CertListForListCertOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CertListForListCertOutput) GoString() string {
	return s.String()
}

// SetCertId sets the CertId field's value.
func (s *CertListForListCertOutput) SetCertId(v string) *CertListForListCertOutput {
	s.CertId = &v
	return s
}

// SetCertName sets the CertName field's value.
func (s *CertListForListCertOutput) SetCertName(v string) *CertListForListCertOutput {
	s.CertName = &v
	return s
}

// SetCertSource sets the CertSource field's value.
func (s *CertListForListCertOutput) SetCertSource(v string) *CertListForListCertOutput {
	s.CertSource = &v
	return s
}

// SetCertStatus sets the CertStatus field's value.
func (s *CertListForListCertOutput) SetCertStatus(v string) *CertListForListCertOutput {
	s.CertStatus = &v
	return s
}

// SetDomain sets the Domain field's value.
func (s *CertListForListCertOutput) SetDomain(v []*string) *CertListForListCertOutput {
	s.Domain = v
	return s
}

// SetExpire sets the Expire field's value.
func (s *CertListForListCertOutput) SetExpire(v string) *CertListForListCertOutput {
	s.Expire = &v
	return s
}

// SetKeyType sets the KeyType field's value.
func (s *CertListForListCertOutput) SetKeyType(v string) *CertListForListCertOutput {
	s.KeyType = &v
	return s
}

// SetMatchDomain sets the MatchDomain field's value.
func (s *CertListForListCertOutput) SetMatchDomain(v []*string) *CertListForListCertOutput {
	s.MatchDomain = v
	return s
}

type ListCertInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	BindDomain *string `type:"string" json:",omitempty"`

	BoundStatus *string `type:"string" json:",omitempty"`

	CertId *string `type:"string" json:",omitempty"`

	CertName *string `type:"string" json:",omitempty"`

	CertSource *string `type:"string" json:",omitempty"`

	CertStatus *string `type:"string" json:",omitempty"`

	ContainDomain *bool `type:"boolean" json:",omitempty"`

	ExpireSortOrder *string `type:"string" json:",omitempty"`

	MatchDomain *string `type:"string" json:",omitempty"`

	PageNumber *int32 `type:"int32" json:",omitempty"`

	PageSize *int32 `type:"int32" json:",omitempty"`

	ProjectName []*string `type:"list" json:",omitempty"`
}

// String returns the string representation
func (s ListCertInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListCertInput) GoString() string {
	return s.String()
}

// SetBindDomain sets the BindDomain field's value.
func (s *ListCertInput) SetBindDomain(v string) *ListCertInput {
	s.BindDomain = &v
	return s
}

// SetBoundStatus sets the BoundStatus field's value.
func (s *ListCertInput) SetBoundStatus(v string) *ListCertInput {
	s.BoundStatus = &v
	return s
}

// SetCertId sets the CertId field's value.
func (s *ListCertInput) SetCertId(v string) *ListCertInput {
	s.CertId = &v
	return s
}

// SetCertName sets the CertName field's value.
func (s *ListCertInput) SetCertName(v string) *ListCertInput {
	s.CertName = &v
	return s
}

// SetCertSource sets the CertSource field's value.
func (s *ListCertInput) SetCertSource(v string) *ListCertInput {
	s.CertSource = &v
	return s
}

// SetCertStatus sets the CertStatus field's value.
func (s *ListCertInput) SetCertStatus(v string) *ListCertInput {
	s.CertStatus = &v
	return s
}

// SetContainDomain sets the ContainDomain field's value.
func (s *ListCertInput) SetContainDomain(v bool) *ListCertInput {
	s.ContainDomain = &v
	return s
}

// SetExpireSortOrder sets the ExpireSortOrder field's value.
func (s *ListCertInput) SetExpireSortOrder(v string) *ListCertInput {
	s.ExpireSortOrder = &v
	return s
}

// SetMatchDomain sets the MatchDomain field's value.
func (s *ListCertInput) SetMatchDomain(v string) *ListCertInput {
	s.MatchDomain = &v
	return s
}

// SetPageNumber sets the PageNumber field's value.
func (s *ListCertInput) SetPageNumber(v int32) *ListCertInput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *ListCertInput) SetPageSize(v int32) *ListCertInput {
	s.PageSize = &v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *ListCertInput) SetProjectName(v []*string) *ListCertInput {
	s.ProjectName = v
	return s
}

type ListCertOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	CertList []*CertListForListCertOutput `type:"list" json:",omitempty"`

	NeedAuth *bool `type:"boolean" json:",omitempty"`

	PageNumber *int32 `type:"int32" json:",omitempty"`

	PageSize *int32 `type:"int32" json:",omitempty"`

	Total *int32 `type:"int32" json:",omitempty"`
}

// String returns the string representation
func (s ListCertOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListCertOutput) GoString() string {
	return s.String()
}

// SetCertList sets the CertList field's value.
func (s *ListCertOutput) SetCertList(v []*CertListForListCertOutput) *ListCertOutput {
	s.CertList = v
	return s
}

// SetNeedAuth sets the NeedAuth field's value.
func (s *ListCertOutput) SetNeedAuth(v bool) *ListCertOutput {
	s.NeedAuth = &v
	return s
}

// SetPageNumber sets the PageNumber field's value.
func (s *ListCertOutput) SetPageNumber(v int32) *ListCertOutput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *ListCertOutput) SetPageSize(v int32) *ListCertOutput {
	s.PageSize = &v
	return s
}

// SetTotal sets the Total field's value.
func (s *ListCertOutput) SetTotal(v int32) *ListCertOutput {
	s.Total = &v
	return s
}

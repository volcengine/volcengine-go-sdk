// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package dcdn

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opListCertBindCommon = "ListCertBind"

// ListCertBindCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ListCertBindCommon operation. The "output" return
// value will be populated with the ListCertBindCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListCertBindCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListCertBindCommon Send returns without error.
//
// See ListCertBindCommon for more information on using the ListCertBindCommon
// API call, and error handling.
//
//    // Example sending a request using the ListCertBindCommonRequest method.
//    req, resp := client.ListCertBindCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *DCDN) ListCertBindCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opListCertBindCommon,
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

// ListCertBindCommon API operation for DCDN.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for DCDN's
// API operation ListCertBindCommon for usage and error information.
func (c *DCDN) ListCertBindCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ListCertBindCommonRequest(input)
	return out, req.Send()
}

// ListCertBindCommonWithContext is the same as ListCertBindCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ListCertBindCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *DCDN) ListCertBindCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ListCertBindCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opListCertBind = "ListCertBind"

// ListCertBindRequest generates a "volcengine/request.Request" representing the
// client's request for the ListCertBind operation. The "output" return
// value will be populated with the ListCertBindCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListCertBindCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListCertBindCommon Send returns without error.
//
// See ListCertBind for more information on using the ListCertBind
// API call, and error handling.
//
//    // Example sending a request using the ListCertBindRequest method.
//    req, resp := client.ListCertBindRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *DCDN) ListCertBindRequest(input *ListCertBindInput) (req *request.Request, output *ListCertBindOutput) {
	op := &request.Operation{
		Name:       opListCertBind,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListCertBindInput{}
	}

	output = &ListCertBindOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// ListCertBind API operation for DCDN.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for DCDN's
// API operation ListCertBind for usage and error information.
func (c *DCDN) ListCertBind(input *ListCertBindInput) (*ListCertBindOutput, error) {
	req, out := c.ListCertBindRequest(input)
	return out, req.Send()
}

// ListCertBindWithContext is the same as ListCertBind with the addition of
// the ability to pass a context and additional request options.
//
// See ListCertBind for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *DCDN) ListCertBindWithContext(ctx volcengine.Context, input *ListCertBindInput, opts ...request.Option) (*ListCertBindOutput, error) {
	req, out := c.ListCertBindRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type BindListForListCertBindOutput struct {
	_ struct{} `type:"structure"`

	CertId *string `type:"string"`

	CertName *string `type:"string"`

	CertSource *string `type:"string"`

	DeployStatus *string `type:"string"`

	DomainId *string `type:"string"`

	DomainName *string `type:"string"`

	Expire *string `type:"string"`
}

// String returns the string representation
func (s BindListForListCertBindOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s BindListForListCertBindOutput) GoString() string {
	return s.String()
}

// SetCertId sets the CertId field's value.
func (s *BindListForListCertBindOutput) SetCertId(v string) *BindListForListCertBindOutput {
	s.CertId = &v
	return s
}

// SetCertName sets the CertName field's value.
func (s *BindListForListCertBindOutput) SetCertName(v string) *BindListForListCertBindOutput {
	s.CertName = &v
	return s
}

// SetCertSource sets the CertSource field's value.
func (s *BindListForListCertBindOutput) SetCertSource(v string) *BindListForListCertBindOutput {
	s.CertSource = &v
	return s
}

// SetDeployStatus sets the DeployStatus field's value.
func (s *BindListForListCertBindOutput) SetDeployStatus(v string) *BindListForListCertBindOutput {
	s.DeployStatus = &v
	return s
}

// SetDomainId sets the DomainId field's value.
func (s *BindListForListCertBindOutput) SetDomainId(v string) *BindListForListCertBindOutput {
	s.DomainId = &v
	return s
}

// SetDomainName sets the DomainName field's value.
func (s *BindListForListCertBindOutput) SetDomainName(v string) *BindListForListCertBindOutput {
	s.DomainName = &v
	return s
}

// SetExpire sets the Expire field's value.
func (s *BindListForListCertBindOutput) SetExpire(v string) *BindListForListCertBindOutput {
	s.Expire = &v
	return s
}

type ListCertBindInput struct {
	_ struct{} `type:"structure"`

	ProjectName []*string `type:"list"`

	SearchKey *string `type:"string"`
}

// String returns the string representation
func (s ListCertBindInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListCertBindInput) GoString() string {
	return s.String()
}

// SetProjectName sets the ProjectName field's value.
func (s *ListCertBindInput) SetProjectName(v []*string) *ListCertBindInput {
	s.ProjectName = v
	return s
}

// SetSearchKey sets the SearchKey field's value.
func (s *ListCertBindInput) SetSearchKey(v string) *ListCertBindInput {
	s.SearchKey = &v
	return s
}

type ListCertBindOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	BindList []*BindListForListCertBindOutput `type:"list"`
}

// String returns the string representation
func (s ListCertBindOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListCertBindOutput) GoString() string {
	return s.String()
}

// SetBindList sets the BindList field's value.
func (s *ListCertBindOutput) SetBindList(v []*BindListForListCertBindOutput) *ListCertBindOutput {
	s.BindList = v
	return s
}

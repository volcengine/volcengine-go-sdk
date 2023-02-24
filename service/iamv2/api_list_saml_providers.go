// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package iamv2

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opListSAMLProvidersCommon = "ListSAMLProviders"

// ListSAMLProvidersCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ListSAMLProvidersCommon operation. The "output" return
// value will be populated with the ListSAMLProvidersCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListSAMLProvidersCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListSAMLProvidersCommon Send returns without error.
//
// See ListSAMLProvidersCommon for more information on using the ListSAMLProvidersCommon
// API call, and error handling.
//
//    // Example sending a request using the ListSAMLProvidersCommonRequest method.
//    req, resp := client.ListSAMLProvidersCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *IAMV2) ListSAMLProvidersCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opListSAMLProvidersCommon,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &map[string]interface{}{}
	}

	output = &map[string]interface{}{}
	req = c.newRequest(op, input, output)

	return
}

// ListSAMLProvidersCommon API operation for IAMV2.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for IAMV2's
// API operation ListSAMLProvidersCommon for usage and error information.
func (c *IAMV2) ListSAMLProvidersCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ListSAMLProvidersCommonRequest(input)
	return out, req.Send()
}

// ListSAMLProvidersCommonWithContext is the same as ListSAMLProvidersCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ListSAMLProvidersCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *IAMV2) ListSAMLProvidersCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ListSAMLProvidersCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opListSAMLProviders = "ListSAMLProviders"

// ListSAMLProvidersRequest generates a "volcengine/request.Request" representing the
// client's request for the ListSAMLProviders operation. The "output" return
// value will be populated with the ListSAMLProvidersCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListSAMLProvidersCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListSAMLProvidersCommon Send returns without error.
//
// See ListSAMLProviders for more information on using the ListSAMLProviders
// API call, and error handling.
//
//    // Example sending a request using the ListSAMLProvidersRequest method.
//    req, resp := client.ListSAMLProvidersRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *IAMV2) ListSAMLProvidersRequest(input *ListSAMLProvidersInput) (req *request.Request, output *ListSAMLProvidersOutput) {
	op := &request.Operation{
		Name:       opListSAMLProviders,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListSAMLProvidersInput{}
	}

	output = &ListSAMLProvidersOutput{}
	req = c.newRequest(op, input, output)

	return
}

// ListSAMLProviders API operation for IAMV2.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for IAMV2's
// API operation ListSAMLProviders for usage and error information.
func (c *IAMV2) ListSAMLProviders(input *ListSAMLProvidersInput) (*ListSAMLProvidersOutput, error) {
	req, out := c.ListSAMLProvidersRequest(input)
	return out, req.Send()
}

// ListSAMLProvidersWithContext is the same as ListSAMLProviders with the addition of
// the ability to pass a context and additional request options.
//
// See ListSAMLProviders for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *IAMV2) ListSAMLProvidersWithContext(ctx volcengine.Context, input *ListSAMLProvidersInput, opts ...request.Option) (*ListSAMLProvidersOutput, error) {
	req, out := c.ListSAMLProvidersRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ListSAMLProvidersInput struct {
	_ struct{} `type:"structure"`

	Limit *int64 `type:"integer"`

	Offset *int64 `type:"integer"`
}

// String returns the string representation
func (s ListSAMLProvidersInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListSAMLProvidersInput) GoString() string {
	return s.String()
}

// SetLimit sets the Limit field's value.
func (s *ListSAMLProvidersInput) SetLimit(v int64) *ListSAMLProvidersInput {
	s.Limit = &v
	return s
}

// SetOffset sets the Offset field's value.
func (s *ListSAMLProvidersInput) SetOffset(v int64) *ListSAMLProvidersInput {
	s.Offset = &v
	return s
}

type ListSAMLProvidersOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	Limit *int32 `type:"int32"`

	Offset *int32 `type:"int32"`

	SAMLProviders []*SAMLProviderForListSAMLProvidersOutput `type:"list"`

	Total *int32 `type:"int32"`
}

// String returns the string representation
func (s ListSAMLProvidersOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListSAMLProvidersOutput) GoString() string {
	return s.String()
}

// SetLimit sets the Limit field's value.
func (s *ListSAMLProvidersOutput) SetLimit(v int32) *ListSAMLProvidersOutput {
	s.Limit = &v
	return s
}

// SetOffset sets the Offset field's value.
func (s *ListSAMLProvidersOutput) SetOffset(v int32) *ListSAMLProvidersOutput {
	s.Offset = &v
	return s
}

// SetSAMLProviders sets the SAMLProviders field's value.
func (s *ListSAMLProvidersOutput) SetSAMLProviders(v []*SAMLProviderForListSAMLProvidersOutput) *ListSAMLProvidersOutput {
	s.SAMLProviders = v
	return s
}

// SetTotal sets the Total field's value.
func (s *ListSAMLProvidersOutput) SetTotal(v int32) *ListSAMLProvidersOutput {
	s.Total = &v
	return s
}

type SAMLProviderForListSAMLProvidersOutput struct {
	_ struct{} `type:"structure"`

	CreateDate *string `type:"string"`

	Description *string `type:"string"`

	SAMLProviderName *string `type:"string"`

	SSOType *int32 `type:"int32"`

	Status *int32 `type:"int32"`

	Trn *string `type:"string"`

	UpdateDate *string `type:"string"`
}

// String returns the string representation
func (s SAMLProviderForListSAMLProvidersOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s SAMLProviderForListSAMLProvidersOutput) GoString() string {
	return s.String()
}

// SetCreateDate sets the CreateDate field's value.
func (s *SAMLProviderForListSAMLProvidersOutput) SetCreateDate(v string) *SAMLProviderForListSAMLProvidersOutput {
	s.CreateDate = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *SAMLProviderForListSAMLProvidersOutput) SetDescription(v string) *SAMLProviderForListSAMLProvidersOutput {
	s.Description = &v
	return s
}

// SetSAMLProviderName sets the SAMLProviderName field's value.
func (s *SAMLProviderForListSAMLProvidersOutput) SetSAMLProviderName(v string) *SAMLProviderForListSAMLProvidersOutput {
	s.SAMLProviderName = &v
	return s
}

// SetSSOType sets the SSOType field's value.
func (s *SAMLProviderForListSAMLProvidersOutput) SetSSOType(v int32) *SAMLProviderForListSAMLProvidersOutput {
	s.SSOType = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *SAMLProviderForListSAMLProvidersOutput) SetStatus(v int32) *SAMLProviderForListSAMLProvidersOutput {
	s.Status = &v
	return s
}

// SetTrn sets the Trn field's value.
func (s *SAMLProviderForListSAMLProvidersOutput) SetTrn(v string) *SAMLProviderForListSAMLProvidersOutput {
	s.Trn = &v
	return s
}

// SetUpdateDate sets the UpdateDate field's value.
func (s *SAMLProviderForListSAMLProvidersOutput) SetUpdateDate(v string) *SAMLProviderForListSAMLProvidersOutput {
	s.UpdateDate = &v
	return s
}

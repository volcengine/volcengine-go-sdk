// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package vke

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opListKubeconfigsCommon = "ListKubeconfigs"

// ListKubeconfigsCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ListKubeconfigsCommon operation. The "output" return
// value will be populated with the ListKubeconfigsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListKubeconfigsCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListKubeconfigsCommon Send returns without error.
//
// See ListKubeconfigsCommon for more information on using the ListKubeconfigsCommon
// API call, and error handling.
//
//    // Example sending a request using the ListKubeconfigsCommonRequest method.
//    req, resp := client.ListKubeconfigsCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VKE) ListKubeconfigsCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opListKubeconfigsCommon,
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

// ListKubeconfigsCommon API operation for VKE.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VKE's
// API operation ListKubeconfigsCommon for usage and error information.
func (c *VKE) ListKubeconfigsCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ListKubeconfigsCommonRequest(input)
	return out, req.Send()
}

// ListKubeconfigsCommonWithContext is the same as ListKubeconfigsCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ListKubeconfigsCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VKE) ListKubeconfigsCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ListKubeconfigsCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opListKubeconfigs = "ListKubeconfigs"

// ListKubeconfigsRequest generates a "volcengine/request.Request" representing the
// client's request for the ListKubeconfigs operation. The "output" return
// value will be populated with the ListKubeconfigsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListKubeconfigsCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListKubeconfigsCommon Send returns without error.
//
// See ListKubeconfigs for more information on using the ListKubeconfigs
// API call, and error handling.
//
//    // Example sending a request using the ListKubeconfigsRequest method.
//    req, resp := client.ListKubeconfigsRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VKE) ListKubeconfigsRequest(input *ListKubeconfigsInput) (req *request.Request, output *ListKubeconfigsOutput) {
	op := &request.Operation{
		Name:       opListKubeconfigs,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListKubeconfigsInput{}
	}

	output = &ListKubeconfigsOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// ListKubeconfigs API operation for VKE.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VKE's
// API operation ListKubeconfigs for usage and error information.
func (c *VKE) ListKubeconfigs(input *ListKubeconfigsInput) (*ListKubeconfigsOutput, error) {
	req, out := c.ListKubeconfigsRequest(input)
	return out, req.Send()
}

// ListKubeconfigsWithContext is the same as ListKubeconfigs with the addition of
// the ability to pass a context and additional request options.
//
// See ListKubeconfigs for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VKE) ListKubeconfigsWithContext(ctx volcengine.Context, input *ListKubeconfigsInput, opts ...request.Option) (*ListKubeconfigsOutput, error) {
	req, out := c.ListKubeconfigsRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type FilterForListKubeconfigsInput struct {
	_ struct{} `type:"structure"`

	ClusterIds []*string `type:"list"`

	Ids []*string `type:"list"`

	RoleIds []*int64 `type:"list"`

	Types []*string `type:"list"`

	UserIds []*int64 `type:"list"`
}

// String returns the string representation
func (s FilterForListKubeconfigsInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s FilterForListKubeconfigsInput) GoString() string {
	return s.String()
}

// SetClusterIds sets the ClusterIds field's value.
func (s *FilterForListKubeconfigsInput) SetClusterIds(v []*string) *FilterForListKubeconfigsInput {
	s.ClusterIds = v
	return s
}

// SetIds sets the Ids field's value.
func (s *FilterForListKubeconfigsInput) SetIds(v []*string) *FilterForListKubeconfigsInput {
	s.Ids = v
	return s
}

// SetRoleIds sets the RoleIds field's value.
func (s *FilterForListKubeconfigsInput) SetRoleIds(v []*int64) *FilterForListKubeconfigsInput {
	s.RoleIds = v
	return s
}

// SetTypes sets the Types field's value.
func (s *FilterForListKubeconfigsInput) SetTypes(v []*string) *FilterForListKubeconfigsInput {
	s.Types = v
	return s
}

// SetUserIds sets the UserIds field's value.
func (s *FilterForListKubeconfigsInput) SetUserIds(v []*int64) *FilterForListKubeconfigsInput {
	s.UserIds = v
	return s
}

type ItemForListKubeconfigsOutput struct {
	_ struct{} `type:"structure"`

	ClusterId *string `type:"string"`

	CreateTime *string `type:"string"`

	ExpireTime *string `type:"string"`

	Id *string `type:"string"`

	Kubeconfig *string `type:"string"`

	RoleId *int64 `type:"int64"`

	Type *string `type:"string"`

	UserId *int64 `type:"int64"`
}

// String returns the string representation
func (s ItemForListKubeconfigsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ItemForListKubeconfigsOutput) GoString() string {
	return s.String()
}

// SetClusterId sets the ClusterId field's value.
func (s *ItemForListKubeconfigsOutput) SetClusterId(v string) *ItemForListKubeconfigsOutput {
	s.ClusterId = &v
	return s
}

// SetCreateTime sets the CreateTime field's value.
func (s *ItemForListKubeconfigsOutput) SetCreateTime(v string) *ItemForListKubeconfigsOutput {
	s.CreateTime = &v
	return s
}

// SetExpireTime sets the ExpireTime field's value.
func (s *ItemForListKubeconfigsOutput) SetExpireTime(v string) *ItemForListKubeconfigsOutput {
	s.ExpireTime = &v
	return s
}

// SetId sets the Id field's value.
func (s *ItemForListKubeconfigsOutput) SetId(v string) *ItemForListKubeconfigsOutput {
	s.Id = &v
	return s
}

// SetKubeconfig sets the Kubeconfig field's value.
func (s *ItemForListKubeconfigsOutput) SetKubeconfig(v string) *ItemForListKubeconfigsOutput {
	s.Kubeconfig = &v
	return s
}

// SetRoleId sets the RoleId field's value.
func (s *ItemForListKubeconfigsOutput) SetRoleId(v int64) *ItemForListKubeconfigsOutput {
	s.RoleId = &v
	return s
}

// SetType sets the Type field's value.
func (s *ItemForListKubeconfigsOutput) SetType(v string) *ItemForListKubeconfigsOutput {
	s.Type = &v
	return s
}

// SetUserId sets the UserId field's value.
func (s *ItemForListKubeconfigsOutput) SetUserId(v int64) *ItemForListKubeconfigsOutput {
	s.UserId = &v
	return s
}

type ListKubeconfigsInput struct {
	_ struct{} `type:"structure"`

	Filter *FilterForListKubeconfigsInput `type:"structure"`

	PageNumber *int32 `type:"int32"`

	PageSize *int32 `type:"int32"`
}

// String returns the string representation
func (s ListKubeconfigsInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListKubeconfigsInput) GoString() string {
	return s.String()
}

// SetFilter sets the Filter field's value.
func (s *ListKubeconfigsInput) SetFilter(v *FilterForListKubeconfigsInput) *ListKubeconfigsInput {
	s.Filter = v
	return s
}

// SetPageNumber sets the PageNumber field's value.
func (s *ListKubeconfigsInput) SetPageNumber(v int32) *ListKubeconfigsInput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *ListKubeconfigsInput) SetPageSize(v int32) *ListKubeconfigsInput {
	s.PageSize = &v
	return s
}

type ListKubeconfigsOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	Items []*ItemForListKubeconfigsOutput `type:"list"`

	PageNumber *int32 `type:"int32"`

	PageSize *int32 `type:"int32"`

	TotalCount *int32 `type:"int32"`
}

// String returns the string representation
func (s ListKubeconfigsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListKubeconfigsOutput) GoString() string {
	return s.String()
}

// SetItems sets the Items field's value.
func (s *ListKubeconfigsOutput) SetItems(v []*ItemForListKubeconfigsOutput) *ListKubeconfigsOutput {
	s.Items = v
	return s
}

// SetPageNumber sets the PageNumber field's value.
func (s *ListKubeconfigsOutput) SetPageNumber(v int32) *ListKubeconfigsOutput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *ListKubeconfigsOutput) SetPageSize(v int32) *ListKubeconfigsOutput {
	s.PageSize = &v
	return s
}

// SetTotalCount sets the TotalCount field's value.
func (s *ListKubeconfigsOutput) SetTotalCount(v int32) *ListKubeconfigsOutput {
	s.TotalCount = &v
	return s
}

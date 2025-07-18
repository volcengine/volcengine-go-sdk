// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package vke

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opListPermissionsCommon = "ListPermissions"

// ListPermissionsCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ListPermissionsCommon operation. The "output" return
// value will be populated with the ListPermissionsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListPermissionsCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListPermissionsCommon Send returns without error.
//
// See ListPermissionsCommon for more information on using the ListPermissionsCommon
// API call, and error handling.
//
//    // Example sending a request using the ListPermissionsCommonRequest method.
//    req, resp := client.ListPermissionsCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VKE) ListPermissionsCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opListPermissionsCommon,
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

// ListPermissionsCommon API operation for VKE.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VKE's
// API operation ListPermissionsCommon for usage and error information.
func (c *VKE) ListPermissionsCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ListPermissionsCommonRequest(input)
	return out, req.Send()
}

// ListPermissionsCommonWithContext is the same as ListPermissionsCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ListPermissionsCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VKE) ListPermissionsCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ListPermissionsCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opListPermissions = "ListPermissions"

// ListPermissionsRequest generates a "volcengine/request.Request" representing the
// client's request for the ListPermissions operation. The "output" return
// value will be populated with the ListPermissionsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListPermissionsCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListPermissionsCommon Send returns without error.
//
// See ListPermissions for more information on using the ListPermissions
// API call, and error handling.
//
//    // Example sending a request using the ListPermissionsRequest method.
//    req, resp := client.ListPermissionsRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VKE) ListPermissionsRequest(input *ListPermissionsInput) (req *request.Request, output *ListPermissionsOutput) {
	op := &request.Operation{
		Name:       opListPermissions,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListPermissionsInput{}
	}

	output = &ListPermissionsOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// ListPermissions API operation for VKE.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VKE's
// API operation ListPermissions for usage and error information.
func (c *VKE) ListPermissions(input *ListPermissionsInput) (*ListPermissionsOutput, error) {
	req, out := c.ListPermissionsRequest(input)
	return out, req.Send()
}

// ListPermissionsWithContext is the same as ListPermissions with the addition of
// the ability to pass a context and additional request options.
//
// See ListPermissions for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VKE) ListPermissionsWithContext(ctx volcengine.Context, input *ListPermissionsInput, opts ...request.Option) (*ListPermissionsOutput, error) {
	req, out := c.ListPermissionsRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type FilterForListPermissionsInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	ClusterIds []*string `type:"list" json:",omitempty"`

	GranteeIds []*int32 `type:"list" json:",omitempty"`

	GranteeType *string `type:"string" json:",omitempty" enum:"EnumOfGranteeTypeForListPermissionsInput"`

	Ids []*string `type:"list" json:",omitempty"`

	Namespaces []*string `type:"list" json:",omitempty"`

	RoleNames []*string `type:"list" json:",omitempty"`

	Status *string `type:"string" json:",omitempty" enum:"EnumOfStatusForListPermissionsInput"`
}

// String returns the string representation
func (s FilterForListPermissionsInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s FilterForListPermissionsInput) GoString() string {
	return s.String()
}

// SetClusterIds sets the ClusterIds field's value.
func (s *FilterForListPermissionsInput) SetClusterIds(v []*string) *FilterForListPermissionsInput {
	s.ClusterIds = v
	return s
}

// SetGranteeIds sets the GranteeIds field's value.
func (s *FilterForListPermissionsInput) SetGranteeIds(v []*int32) *FilterForListPermissionsInput {
	s.GranteeIds = v
	return s
}

// SetGranteeType sets the GranteeType field's value.
func (s *FilterForListPermissionsInput) SetGranteeType(v string) *FilterForListPermissionsInput {
	s.GranteeType = &v
	return s
}

// SetIds sets the Ids field's value.
func (s *FilterForListPermissionsInput) SetIds(v []*string) *FilterForListPermissionsInput {
	s.Ids = v
	return s
}

// SetNamespaces sets the Namespaces field's value.
func (s *FilterForListPermissionsInput) SetNamespaces(v []*string) *FilterForListPermissionsInput {
	s.Namespaces = v
	return s
}

// SetRoleNames sets the RoleNames field's value.
func (s *FilterForListPermissionsInput) SetRoleNames(v []*string) *FilterForListPermissionsInput {
	s.RoleNames = v
	return s
}

// SetStatus sets the Status field's value.
func (s *FilterForListPermissionsInput) SetStatus(v string) *FilterForListPermissionsInput {
	s.Status = &v
	return s
}

type ItemForListPermissionsOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	AuthorizedAt *string `type:"string" json:",omitempty"`

	AuthorizerId *int64 `type:"int64" json:",omitempty"`

	AuthorizerName *string `type:"string" json:",omitempty"`

	AuthorizerType *string `type:"string" json:",omitempty" enum:"EnumOfAuthorizerTypeForListPermissionsOutput"`

	ClusterId *string `type:"string" json:",omitempty"`

	GrantedAt *string `type:"string" json:",omitempty"`

	GranteeId *int64 `type:"int64" json:",omitempty"`

	GranteeType *string `type:"string" json:",omitempty" enum:"EnumOfGranteeTypeForListPermissionsOutput"`

	Id *string `type:"string" json:",omitempty"`

	IsCustomRole *bool `type:"boolean" json:",omitempty"`

	KubeRoleBindingName *string `type:"string" json:",omitempty"`

	Message *string `type:"string" json:",omitempty"`

	Namespace *string `type:"string" json:",omitempty"`

	ProjectSelector *string `type:"string" json:",omitempty"`

	RevokedAt *string `type:"string" json:",omitempty"`

	RoleName *string `type:"string" json:",omitempty"`

	Status *string `type:"string" json:",omitempty" enum:"EnumOfStatusForListPermissionsOutput"`
}

// String returns the string representation
func (s ItemForListPermissionsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ItemForListPermissionsOutput) GoString() string {
	return s.String()
}

// SetAuthorizedAt sets the AuthorizedAt field's value.
func (s *ItemForListPermissionsOutput) SetAuthorizedAt(v string) *ItemForListPermissionsOutput {
	s.AuthorizedAt = &v
	return s
}

// SetAuthorizerId sets the AuthorizerId field's value.
func (s *ItemForListPermissionsOutput) SetAuthorizerId(v int64) *ItemForListPermissionsOutput {
	s.AuthorizerId = &v
	return s
}

// SetAuthorizerName sets the AuthorizerName field's value.
func (s *ItemForListPermissionsOutput) SetAuthorizerName(v string) *ItemForListPermissionsOutput {
	s.AuthorizerName = &v
	return s
}

// SetAuthorizerType sets the AuthorizerType field's value.
func (s *ItemForListPermissionsOutput) SetAuthorizerType(v string) *ItemForListPermissionsOutput {
	s.AuthorizerType = &v
	return s
}

// SetClusterId sets the ClusterId field's value.
func (s *ItemForListPermissionsOutput) SetClusterId(v string) *ItemForListPermissionsOutput {
	s.ClusterId = &v
	return s
}

// SetGrantedAt sets the GrantedAt field's value.
func (s *ItemForListPermissionsOutput) SetGrantedAt(v string) *ItemForListPermissionsOutput {
	s.GrantedAt = &v
	return s
}

// SetGranteeId sets the GranteeId field's value.
func (s *ItemForListPermissionsOutput) SetGranteeId(v int64) *ItemForListPermissionsOutput {
	s.GranteeId = &v
	return s
}

// SetGranteeType sets the GranteeType field's value.
func (s *ItemForListPermissionsOutput) SetGranteeType(v string) *ItemForListPermissionsOutput {
	s.GranteeType = &v
	return s
}

// SetId sets the Id field's value.
func (s *ItemForListPermissionsOutput) SetId(v string) *ItemForListPermissionsOutput {
	s.Id = &v
	return s
}

// SetIsCustomRole sets the IsCustomRole field's value.
func (s *ItemForListPermissionsOutput) SetIsCustomRole(v bool) *ItemForListPermissionsOutput {
	s.IsCustomRole = &v
	return s
}

// SetKubeRoleBindingName sets the KubeRoleBindingName field's value.
func (s *ItemForListPermissionsOutput) SetKubeRoleBindingName(v string) *ItemForListPermissionsOutput {
	s.KubeRoleBindingName = &v
	return s
}

// SetMessage sets the Message field's value.
func (s *ItemForListPermissionsOutput) SetMessage(v string) *ItemForListPermissionsOutput {
	s.Message = &v
	return s
}

// SetNamespace sets the Namespace field's value.
func (s *ItemForListPermissionsOutput) SetNamespace(v string) *ItemForListPermissionsOutput {
	s.Namespace = &v
	return s
}

// SetProjectSelector sets the ProjectSelector field's value.
func (s *ItemForListPermissionsOutput) SetProjectSelector(v string) *ItemForListPermissionsOutput {
	s.ProjectSelector = &v
	return s
}

// SetRevokedAt sets the RevokedAt field's value.
func (s *ItemForListPermissionsOutput) SetRevokedAt(v string) *ItemForListPermissionsOutput {
	s.RevokedAt = &v
	return s
}

// SetRoleName sets the RoleName field's value.
func (s *ItemForListPermissionsOutput) SetRoleName(v string) *ItemForListPermissionsOutput {
	s.RoleName = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *ItemForListPermissionsOutput) SetStatus(v string) *ItemForListPermissionsOutput {
	s.Status = &v
	return s
}

type ListPermissionsInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Filter *FilterForListPermissionsInput `type:"structure" json:",omitempty"`

	PageNumber *int32 `type:"int32" json:",omitempty"`

	PageSize *int32 `type:"int32" json:",omitempty"`
}

// String returns the string representation
func (s ListPermissionsInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListPermissionsInput) GoString() string {
	return s.String()
}

// SetFilter sets the Filter field's value.
func (s *ListPermissionsInput) SetFilter(v *FilterForListPermissionsInput) *ListPermissionsInput {
	s.Filter = v
	return s
}

// SetPageNumber sets the PageNumber field's value.
func (s *ListPermissionsInput) SetPageNumber(v int32) *ListPermissionsInput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *ListPermissionsInput) SetPageSize(v int32) *ListPermissionsInput {
	s.PageSize = &v
	return s
}

type ListPermissionsOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	Items []*ItemForListPermissionsOutput `type:"list" json:",omitempty"`

	PageNumber *int32 `type:"int32" json:",omitempty"`

	PageSize *int32 `type:"int32" json:",omitempty"`

	TotalCount *int32 `type:"int32" json:",omitempty"`
}

// String returns the string representation
func (s ListPermissionsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListPermissionsOutput) GoString() string {
	return s.String()
}

// SetItems sets the Items field's value.
func (s *ListPermissionsOutput) SetItems(v []*ItemForListPermissionsOutput) *ListPermissionsOutput {
	s.Items = v
	return s
}

// SetPageNumber sets the PageNumber field's value.
func (s *ListPermissionsOutput) SetPageNumber(v int32) *ListPermissionsOutput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *ListPermissionsOutput) SetPageSize(v int32) *ListPermissionsOutput {
	s.PageSize = &v
	return s
}

// SetTotalCount sets the TotalCount field's value.
func (s *ListPermissionsOutput) SetTotalCount(v int32) *ListPermissionsOutput {
	s.TotalCount = &v
	return s
}

const (
	// EnumOfAuthorizerTypeForListPermissionsOutputUser is a EnumOfAuthorizerTypeForListPermissionsOutput enum value
	EnumOfAuthorizerTypeForListPermissionsOutputUser = "User"

	// EnumOfAuthorizerTypeForListPermissionsOutputRole is a EnumOfAuthorizerTypeForListPermissionsOutput enum value
	EnumOfAuthorizerTypeForListPermissionsOutputRole = "Role"

	// EnumOfAuthorizerTypeForListPermissionsOutputAccount is a EnumOfAuthorizerTypeForListPermissionsOutput enum value
	EnumOfAuthorizerTypeForListPermissionsOutputAccount = "Account"
)

const (
	// EnumOfGranteeTypeForListPermissionsInputUser is a EnumOfGranteeTypeForListPermissionsInput enum value
	EnumOfGranteeTypeForListPermissionsInputUser = "User"

	// EnumOfGranteeTypeForListPermissionsInputRole is a EnumOfGranteeTypeForListPermissionsInput enum value
	EnumOfGranteeTypeForListPermissionsInputRole = "Role"

	// EnumOfGranteeTypeForListPermissionsInputAccount is a EnumOfGranteeTypeForListPermissionsInput enum value
	EnumOfGranteeTypeForListPermissionsInputAccount = "Account"
)

const (
	// EnumOfGranteeTypeForListPermissionsOutputUser is a EnumOfGranteeTypeForListPermissionsOutput enum value
	EnumOfGranteeTypeForListPermissionsOutputUser = "User"

	// EnumOfGranteeTypeForListPermissionsOutputRole is a EnumOfGranteeTypeForListPermissionsOutput enum value
	EnumOfGranteeTypeForListPermissionsOutputRole = "Role"

	// EnumOfGranteeTypeForListPermissionsOutputAccount is a EnumOfGranteeTypeForListPermissionsOutput enum value
	EnumOfGranteeTypeForListPermissionsOutputAccount = "Account"
)

const (
	// EnumOfStatusForListPermissionsInputPending is a EnumOfStatusForListPermissionsInput enum value
	EnumOfStatusForListPermissionsInputPending = "Pending"

	// EnumOfStatusForListPermissionsInputPartialSuccess is a EnumOfStatusForListPermissionsInput enum value
	EnumOfStatusForListPermissionsInputPartialSuccess = "PartialSuccess"

	// EnumOfStatusForListPermissionsInputSuccess is a EnumOfStatusForListPermissionsInput enum value
	EnumOfStatusForListPermissionsInputSuccess = "Success"

	// EnumOfStatusForListPermissionsInputFailed is a EnumOfStatusForListPermissionsInput enum value
	EnumOfStatusForListPermissionsInputFailed = "Failed"
)

const (
	// EnumOfStatusForListPermissionsOutputPending is a EnumOfStatusForListPermissionsOutput enum value
	EnumOfStatusForListPermissionsOutputPending = "Pending"

	// EnumOfStatusForListPermissionsOutputPartialSuccess is a EnumOfStatusForListPermissionsOutput enum value
	EnumOfStatusForListPermissionsOutputPartialSuccess = "PartialSuccess"

	// EnumOfStatusForListPermissionsOutputSuccess is a EnumOfStatusForListPermissionsOutput enum value
	EnumOfStatusForListPermissionsOutputSuccess = "Success"

	// EnumOfStatusForListPermissionsOutputFailed is a EnumOfStatusForListPermissionsOutput enum value
	EnumOfStatusForListPermissionsOutputFailed = "Failed"
)

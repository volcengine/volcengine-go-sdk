// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package organization

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opListServiceControlPoliciesCommon = "ListServiceControlPolicies"

// ListServiceControlPoliciesCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ListServiceControlPoliciesCommon operation. The "output" return
// value will be populated with the ListServiceControlPoliciesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListServiceControlPoliciesCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListServiceControlPoliciesCommon Send returns without error.
//
// See ListServiceControlPoliciesCommon for more information on using the ListServiceControlPoliciesCommon
// API call, and error handling.
//
//    // Example sending a request using the ListServiceControlPoliciesCommonRequest method.
//    req, resp := client.ListServiceControlPoliciesCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ORGANIZATION) ListServiceControlPoliciesCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opListServiceControlPoliciesCommon,
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

// ListServiceControlPoliciesCommon API operation for ORGANIZATION.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ORGANIZATION's
// API operation ListServiceControlPoliciesCommon for usage and error information.
func (c *ORGANIZATION) ListServiceControlPoliciesCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ListServiceControlPoliciesCommonRequest(input)
	return out, req.Send()
}

// ListServiceControlPoliciesCommonWithContext is the same as ListServiceControlPoliciesCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ListServiceControlPoliciesCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ORGANIZATION) ListServiceControlPoliciesCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ListServiceControlPoliciesCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opListServiceControlPolicies = "ListServiceControlPolicies"

// ListServiceControlPoliciesRequest generates a "volcengine/request.Request" representing the
// client's request for the ListServiceControlPolicies operation. The "output" return
// value will be populated with the ListServiceControlPoliciesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListServiceControlPoliciesCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListServiceControlPoliciesCommon Send returns without error.
//
// See ListServiceControlPolicies for more information on using the ListServiceControlPolicies
// API call, and error handling.
//
//    // Example sending a request using the ListServiceControlPoliciesRequest method.
//    req, resp := client.ListServiceControlPoliciesRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ORGANIZATION) ListServiceControlPoliciesRequest(input *ListServiceControlPoliciesInput) (req *request.Request, output *ListServiceControlPoliciesOutput) {
	op := &request.Operation{
		Name:       opListServiceControlPolicies,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListServiceControlPoliciesInput{}
	}

	output = &ListServiceControlPoliciesOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// ListServiceControlPolicies API operation for ORGANIZATION.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ORGANIZATION's
// API operation ListServiceControlPolicies for usage and error information.
func (c *ORGANIZATION) ListServiceControlPolicies(input *ListServiceControlPoliciesInput) (*ListServiceControlPoliciesOutput, error) {
	req, out := c.ListServiceControlPoliciesRequest(input)
	return out, req.Send()
}

// ListServiceControlPoliciesWithContext is the same as ListServiceControlPolicies with the addition of
// the ability to pass a context and additional request options.
//
// See ListServiceControlPolicies for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ORGANIZATION) ListServiceControlPoliciesWithContext(ctx volcengine.Context, input *ListServiceControlPoliciesInput, opts ...request.Option) (*ListServiceControlPoliciesOutput, error) {
	req, out := c.ListServiceControlPoliciesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ListServiceControlPoliciesInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	PageNumber *int32 `type:"int32" json:",omitempty"`

	PageSize *int32 `type:"int32" json:",omitempty"`

	PolicyType *string `type:"string" json:",omitempty"`

	Query *string `type:"string" json:",omitempty"`

	SortBy *string `type:"string" json:",omitempty"`

	SortOrder *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s ListServiceControlPoliciesInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListServiceControlPoliciesInput) GoString() string {
	return s.String()
}

// SetPageNumber sets the PageNumber field's value.
func (s *ListServiceControlPoliciesInput) SetPageNumber(v int32) *ListServiceControlPoliciesInput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *ListServiceControlPoliciesInput) SetPageSize(v int32) *ListServiceControlPoliciesInput {
	s.PageSize = &v
	return s
}

// SetPolicyType sets the PolicyType field's value.
func (s *ListServiceControlPoliciesInput) SetPolicyType(v string) *ListServiceControlPoliciesInput {
	s.PolicyType = &v
	return s
}

// SetQuery sets the Query field's value.
func (s *ListServiceControlPoliciesInput) SetQuery(v string) *ListServiceControlPoliciesInput {
	s.Query = &v
	return s
}

// SetSortBy sets the SortBy field's value.
func (s *ListServiceControlPoliciesInput) SetSortBy(v string) *ListServiceControlPoliciesInput {
	s.SortBy = &v
	return s
}

// SetSortOrder sets the SortOrder field's value.
func (s *ListServiceControlPoliciesInput) SetSortOrder(v string) *ListServiceControlPoliciesInput {
	s.SortOrder = &v
	return s
}

type ListServiceControlPoliciesOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	PageNumber *int32 `type:"int32" json:",omitempty"`

	PageSize *int32 `type:"int32" json:",omitempty"`

	ServiceControlPolicies []*ServiceControlPolicyForListServiceControlPoliciesOutput `type:"list" json:",omitempty"`

	TotalCount *int32 `type:"int32" json:",omitempty"`
}

// String returns the string representation
func (s ListServiceControlPoliciesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListServiceControlPoliciesOutput) GoString() string {
	return s.String()
}

// SetPageNumber sets the PageNumber field's value.
func (s *ListServiceControlPoliciesOutput) SetPageNumber(v int32) *ListServiceControlPoliciesOutput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *ListServiceControlPoliciesOutput) SetPageSize(v int32) *ListServiceControlPoliciesOutput {
	s.PageSize = &v
	return s
}

// SetServiceControlPolicies sets the ServiceControlPolicies field's value.
func (s *ListServiceControlPoliciesOutput) SetServiceControlPolicies(v []*ServiceControlPolicyForListServiceControlPoliciesOutput) *ListServiceControlPoliciesOutput {
	s.ServiceControlPolicies = v
	return s
}

// SetTotalCount sets the TotalCount field's value.
func (s *ListServiceControlPoliciesOutput) SetTotalCount(v int32) *ListServiceControlPoliciesOutput {
	s.TotalCount = &v
	return s
}

type ServiceControlPolicyForListServiceControlPoliciesOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	CreateDate *string `type:"string" json:",omitempty"`

	Description *string `type:"string" json:",omitempty"`

	PolicyID *string `type:"string" json:",omitempty"`

	PolicyName *string `type:"string" json:",omitempty"`

	PolicyType *string `type:"string" json:",omitempty"`

	UpdateDate *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s ServiceControlPolicyForListServiceControlPoliciesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ServiceControlPolicyForListServiceControlPoliciesOutput) GoString() string {
	return s.String()
}

// SetCreateDate sets the CreateDate field's value.
func (s *ServiceControlPolicyForListServiceControlPoliciesOutput) SetCreateDate(v string) *ServiceControlPolicyForListServiceControlPoliciesOutput {
	s.CreateDate = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *ServiceControlPolicyForListServiceControlPoliciesOutput) SetDescription(v string) *ServiceControlPolicyForListServiceControlPoliciesOutput {
	s.Description = &v
	return s
}

// SetPolicyID sets the PolicyID field's value.
func (s *ServiceControlPolicyForListServiceControlPoliciesOutput) SetPolicyID(v string) *ServiceControlPolicyForListServiceControlPoliciesOutput {
	s.PolicyID = &v
	return s
}

// SetPolicyName sets the PolicyName field's value.
func (s *ServiceControlPolicyForListServiceControlPoliciesOutput) SetPolicyName(v string) *ServiceControlPolicyForListServiceControlPoliciesOutput {
	s.PolicyName = &v
	return s
}

// SetPolicyType sets the PolicyType field's value.
func (s *ServiceControlPolicyForListServiceControlPoliciesOutput) SetPolicyType(v string) *ServiceControlPolicyForListServiceControlPoliciesOutput {
	s.PolicyType = &v
	return s
}

// SetUpdateDate sets the UpdateDate field's value.
func (s *ServiceControlPolicyForListServiceControlPoliciesOutput) SetUpdateDate(v string) *ServiceControlPolicyForListServiceControlPoliciesOutput {
	s.UpdateDate = &v
	return s
}

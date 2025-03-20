// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudidentity

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opListPermissionPoliciesInPermissionSetCommon = "ListPermissionPoliciesInPermissionSet"

// ListPermissionPoliciesInPermissionSetCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ListPermissionPoliciesInPermissionSetCommon operation. The "output" return
// value will be populated with the ListPermissionPoliciesInPermissionSetCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListPermissionPoliciesInPermissionSetCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListPermissionPoliciesInPermissionSetCommon Send returns without error.
//
// See ListPermissionPoliciesInPermissionSetCommon for more information on using the ListPermissionPoliciesInPermissionSetCommon
// API call, and error handling.
//
//    // Example sending a request using the ListPermissionPoliciesInPermissionSetCommonRequest method.
//    req, resp := client.ListPermissionPoliciesInPermissionSetCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CLOUDIDENTITY) ListPermissionPoliciesInPermissionSetCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opListPermissionPoliciesInPermissionSetCommon,
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

// ListPermissionPoliciesInPermissionSetCommon API operation for CLOUDIDENTITY.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CLOUDIDENTITY's
// API operation ListPermissionPoliciesInPermissionSetCommon for usage and error information.
func (c *CLOUDIDENTITY) ListPermissionPoliciesInPermissionSetCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ListPermissionPoliciesInPermissionSetCommonRequest(input)
	return out, req.Send()
}

// ListPermissionPoliciesInPermissionSetCommonWithContext is the same as ListPermissionPoliciesInPermissionSetCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ListPermissionPoliciesInPermissionSetCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CLOUDIDENTITY) ListPermissionPoliciesInPermissionSetCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ListPermissionPoliciesInPermissionSetCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opListPermissionPoliciesInPermissionSet = "ListPermissionPoliciesInPermissionSet"

// ListPermissionPoliciesInPermissionSetRequest generates a "volcengine/request.Request" representing the
// client's request for the ListPermissionPoliciesInPermissionSet operation. The "output" return
// value will be populated with the ListPermissionPoliciesInPermissionSetCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListPermissionPoliciesInPermissionSetCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListPermissionPoliciesInPermissionSetCommon Send returns without error.
//
// See ListPermissionPoliciesInPermissionSet for more information on using the ListPermissionPoliciesInPermissionSet
// API call, and error handling.
//
//    // Example sending a request using the ListPermissionPoliciesInPermissionSetRequest method.
//    req, resp := client.ListPermissionPoliciesInPermissionSetRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CLOUDIDENTITY) ListPermissionPoliciesInPermissionSetRequest(input *ListPermissionPoliciesInPermissionSetInput) (req *request.Request, output *ListPermissionPoliciesInPermissionSetOutput) {
	op := &request.Operation{
		Name:       opListPermissionPoliciesInPermissionSet,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListPermissionPoliciesInPermissionSetInput{}
	}

	output = &ListPermissionPoliciesInPermissionSetOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// ListPermissionPoliciesInPermissionSet API operation for CLOUDIDENTITY.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CLOUDIDENTITY's
// API operation ListPermissionPoliciesInPermissionSet for usage and error information.
func (c *CLOUDIDENTITY) ListPermissionPoliciesInPermissionSet(input *ListPermissionPoliciesInPermissionSetInput) (*ListPermissionPoliciesInPermissionSetOutput, error) {
	req, out := c.ListPermissionPoliciesInPermissionSetRequest(input)
	return out, req.Send()
}

// ListPermissionPoliciesInPermissionSetWithContext is the same as ListPermissionPoliciesInPermissionSet with the addition of
// the ability to pass a context and additional request options.
//
// See ListPermissionPoliciesInPermissionSet for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CLOUDIDENTITY) ListPermissionPoliciesInPermissionSetWithContext(ctx volcengine.Context, input *ListPermissionPoliciesInPermissionSetInput, opts ...request.Option) (*ListPermissionPoliciesInPermissionSetOutput, error) {
	req, out := c.ListPermissionPoliciesInPermissionSetRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ListPermissionPoliciesInPermissionSetInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	PageNumber *int32 `type:"int32" json:",omitempty"`

	PageSize *int32 `type:"int32" json:",omitempty"`

	PermissionPolicyType *string `type:"string" json:",omitempty" enum:"EnumOfPermissionPolicyTypeForListPermissionPoliciesInPermissionSetInput"`

	// PermissionSetId is a required field
	PermissionSetId *string `type:"string" json:",omitempty" required:"true"`
}

// String returns the string representation
func (s ListPermissionPoliciesInPermissionSetInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListPermissionPoliciesInPermissionSetInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListPermissionPoliciesInPermissionSetInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ListPermissionPoliciesInPermissionSetInput"}
	if s.PermissionSetId == nil {
		invalidParams.Add(request.NewErrParamRequired("PermissionSetId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetPageNumber sets the PageNumber field's value.
func (s *ListPermissionPoliciesInPermissionSetInput) SetPageNumber(v int32) *ListPermissionPoliciesInPermissionSetInput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *ListPermissionPoliciesInPermissionSetInput) SetPageSize(v int32) *ListPermissionPoliciesInPermissionSetInput {
	s.PageSize = &v
	return s
}

// SetPermissionPolicyType sets the PermissionPolicyType field's value.
func (s *ListPermissionPoliciesInPermissionSetInput) SetPermissionPolicyType(v string) *ListPermissionPoliciesInPermissionSetInput {
	s.PermissionPolicyType = &v
	return s
}

// SetPermissionSetId sets the PermissionSetId field's value.
func (s *ListPermissionPoliciesInPermissionSetInput) SetPermissionSetId(v string) *ListPermissionPoliciesInPermissionSetInput {
	s.PermissionSetId = &v
	return s
}

type ListPermissionPoliciesInPermissionSetOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	PageNumber *int32 `type:"int32" json:",omitempty"`

	PageSize *int32 `type:"int32" json:",omitempty"`

	PermissionPolicies []*PermissionPolicyForListPermissionPoliciesInPermissionSetOutput `type:"list" json:",omitempty"`

	Total *int64 `type:"int64" json:",omitempty"`
}

// String returns the string representation
func (s ListPermissionPoliciesInPermissionSetOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListPermissionPoliciesInPermissionSetOutput) GoString() string {
	return s.String()
}

// SetPageNumber sets the PageNumber field's value.
func (s *ListPermissionPoliciesInPermissionSetOutput) SetPageNumber(v int32) *ListPermissionPoliciesInPermissionSetOutput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *ListPermissionPoliciesInPermissionSetOutput) SetPageSize(v int32) *ListPermissionPoliciesInPermissionSetOutput {
	s.PageSize = &v
	return s
}

// SetPermissionPolicies sets the PermissionPolicies field's value.
func (s *ListPermissionPoliciesInPermissionSetOutput) SetPermissionPolicies(v []*PermissionPolicyForListPermissionPoliciesInPermissionSetOutput) *ListPermissionPoliciesInPermissionSetOutput {
	s.PermissionPolicies = v
	return s
}

// SetTotal sets the Total field's value.
func (s *ListPermissionPoliciesInPermissionSetOutput) SetTotal(v int64) *ListPermissionPoliciesInPermissionSetOutput {
	s.Total = &v
	return s
}

type PermissionPolicyForListPermissionPoliciesInPermissionSetOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	CreatedTime *string `type:"string" json:",omitempty"`

	PermissionPolicyDocument *string `type:"string" json:",omitempty"`

	PermissionPolicyName *string `type:"string" json:",omitempty"`

	PermissionPolicyType *string `type:"string" json:",omitempty" enum:"EnumOfPermissionPolicyTypeForListPermissionPoliciesInPermissionSetOutput"`
}

// String returns the string representation
func (s PermissionPolicyForListPermissionPoliciesInPermissionSetOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s PermissionPolicyForListPermissionPoliciesInPermissionSetOutput) GoString() string {
	return s.String()
}

// SetCreatedTime sets the CreatedTime field's value.
func (s *PermissionPolicyForListPermissionPoliciesInPermissionSetOutput) SetCreatedTime(v string) *PermissionPolicyForListPermissionPoliciesInPermissionSetOutput {
	s.CreatedTime = &v
	return s
}

// SetPermissionPolicyDocument sets the PermissionPolicyDocument field's value.
func (s *PermissionPolicyForListPermissionPoliciesInPermissionSetOutput) SetPermissionPolicyDocument(v string) *PermissionPolicyForListPermissionPoliciesInPermissionSetOutput {
	s.PermissionPolicyDocument = &v
	return s
}

// SetPermissionPolicyName sets the PermissionPolicyName field's value.
func (s *PermissionPolicyForListPermissionPoliciesInPermissionSetOutput) SetPermissionPolicyName(v string) *PermissionPolicyForListPermissionPoliciesInPermissionSetOutput {
	s.PermissionPolicyName = &v
	return s
}

// SetPermissionPolicyType sets the PermissionPolicyType field's value.
func (s *PermissionPolicyForListPermissionPoliciesInPermissionSetOutput) SetPermissionPolicyType(v string) *PermissionPolicyForListPermissionPoliciesInPermissionSetOutput {
	s.PermissionPolicyType = &v
	return s
}

const (
	// EnumOfPermissionPolicyTypeForListPermissionPoliciesInPermissionSetInputSystem is a EnumOfPermissionPolicyTypeForListPermissionPoliciesInPermissionSetInput enum value
	EnumOfPermissionPolicyTypeForListPermissionPoliciesInPermissionSetInputSystem = "System"

	// EnumOfPermissionPolicyTypeForListPermissionPoliciesInPermissionSetInputInline is a EnumOfPermissionPolicyTypeForListPermissionPoliciesInPermissionSetInput enum value
	EnumOfPermissionPolicyTypeForListPermissionPoliciesInPermissionSetInputInline = "Inline"
)

const (
	// EnumOfPermissionPolicyTypeForListPermissionPoliciesInPermissionSetOutputSystem is a EnumOfPermissionPolicyTypeForListPermissionPoliciesInPermissionSetOutput enum value
	EnumOfPermissionPolicyTypeForListPermissionPoliciesInPermissionSetOutputSystem = "System"

	// EnumOfPermissionPolicyTypeForListPermissionPoliciesInPermissionSetOutputInline is a EnumOfPermissionPolicyTypeForListPermissionPoliciesInPermissionSetOutput enum value
	EnumOfPermissionPolicyTypeForListPermissionPoliciesInPermissionSetOutputInline = "Inline"
)

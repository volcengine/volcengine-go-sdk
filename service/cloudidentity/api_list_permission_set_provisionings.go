// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudidentity

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opListPermissionSetProvisioningsCommon = "ListPermissionSetProvisionings"

// ListPermissionSetProvisioningsCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ListPermissionSetProvisioningsCommon operation. The "output" return
// value will be populated with the ListPermissionSetProvisioningsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListPermissionSetProvisioningsCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListPermissionSetProvisioningsCommon Send returns without error.
//
// See ListPermissionSetProvisioningsCommon for more information on using the ListPermissionSetProvisioningsCommon
// API call, and error handling.
//
//    // Example sending a request using the ListPermissionSetProvisioningsCommonRequest method.
//    req, resp := client.ListPermissionSetProvisioningsCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CLOUDIDENTITY) ListPermissionSetProvisioningsCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opListPermissionSetProvisioningsCommon,
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

// ListPermissionSetProvisioningsCommon API operation for CLOUDIDENTITY.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CLOUDIDENTITY's
// API operation ListPermissionSetProvisioningsCommon for usage and error information.
func (c *CLOUDIDENTITY) ListPermissionSetProvisioningsCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ListPermissionSetProvisioningsCommonRequest(input)
	return out, req.Send()
}

// ListPermissionSetProvisioningsCommonWithContext is the same as ListPermissionSetProvisioningsCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ListPermissionSetProvisioningsCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CLOUDIDENTITY) ListPermissionSetProvisioningsCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ListPermissionSetProvisioningsCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opListPermissionSetProvisionings = "ListPermissionSetProvisionings"

// ListPermissionSetProvisioningsRequest generates a "volcengine/request.Request" representing the
// client's request for the ListPermissionSetProvisionings operation. The "output" return
// value will be populated with the ListPermissionSetProvisioningsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListPermissionSetProvisioningsCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListPermissionSetProvisioningsCommon Send returns without error.
//
// See ListPermissionSetProvisionings for more information on using the ListPermissionSetProvisionings
// API call, and error handling.
//
//    // Example sending a request using the ListPermissionSetProvisioningsRequest method.
//    req, resp := client.ListPermissionSetProvisioningsRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CLOUDIDENTITY) ListPermissionSetProvisioningsRequest(input *ListPermissionSetProvisioningsInput) (req *request.Request, output *ListPermissionSetProvisioningsOutput) {
	op := &request.Operation{
		Name:       opListPermissionSetProvisionings,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListPermissionSetProvisioningsInput{}
	}

	output = &ListPermissionSetProvisioningsOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// ListPermissionSetProvisionings API operation for CLOUDIDENTITY.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CLOUDIDENTITY's
// API operation ListPermissionSetProvisionings for usage and error information.
func (c *CLOUDIDENTITY) ListPermissionSetProvisionings(input *ListPermissionSetProvisioningsInput) (*ListPermissionSetProvisioningsOutput, error) {
	req, out := c.ListPermissionSetProvisioningsRequest(input)
	return out, req.Send()
}

// ListPermissionSetProvisioningsWithContext is the same as ListPermissionSetProvisionings with the addition of
// the ability to pass a context and additional request options.
//
// See ListPermissionSetProvisionings for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CLOUDIDENTITY) ListPermissionSetProvisioningsWithContext(ctx volcengine.Context, input *ListPermissionSetProvisioningsInput, opts ...request.Option) (*ListPermissionSetProvisioningsOutput, error) {
	req, out := c.ListPermissionSetProvisioningsRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ListPermissionSetProvisioningsInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	PageNumber *int32 `type:"int32" json:",omitempty"`

	PageSize *int32 `type:"int32" json:",omitempty"`

	PermissionSetId *string `type:"string" json:",omitempty"`

	ProvisioningStatus *string `type:"string" json:",omitempty" enum:"EnumOfProvisioningStatusForListPermissionSetProvisioningsInput"`

	TargetId *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s ListPermissionSetProvisioningsInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListPermissionSetProvisioningsInput) GoString() string {
	return s.String()
}

// SetPageNumber sets the PageNumber field's value.
func (s *ListPermissionSetProvisioningsInput) SetPageNumber(v int32) *ListPermissionSetProvisioningsInput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *ListPermissionSetProvisioningsInput) SetPageSize(v int32) *ListPermissionSetProvisioningsInput {
	s.PageSize = &v
	return s
}

// SetPermissionSetId sets the PermissionSetId field's value.
func (s *ListPermissionSetProvisioningsInput) SetPermissionSetId(v string) *ListPermissionSetProvisioningsInput {
	s.PermissionSetId = &v
	return s
}

// SetProvisioningStatus sets the ProvisioningStatus field's value.
func (s *ListPermissionSetProvisioningsInput) SetProvisioningStatus(v string) *ListPermissionSetProvisioningsInput {
	s.ProvisioningStatus = &v
	return s
}

// SetTargetId sets the TargetId field's value.
func (s *ListPermissionSetProvisioningsInput) SetTargetId(v string) *ListPermissionSetProvisioningsInput {
	s.TargetId = &v
	return s
}

type ListPermissionSetProvisioningsOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	PageNumber *int32 `type:"int32" json:",omitempty"`

	PageSize *int32 `type:"int32" json:",omitempty"`

	PermissionSetProvisionings []*PermissionSetProvisioningForListPermissionSetProvisioningsOutput `type:"list" json:",omitempty"`

	Total *int64 `type:"int64" json:",omitempty"`
}

// String returns the string representation
func (s ListPermissionSetProvisioningsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListPermissionSetProvisioningsOutput) GoString() string {
	return s.String()
}

// SetPageNumber sets the PageNumber field's value.
func (s *ListPermissionSetProvisioningsOutput) SetPageNumber(v int32) *ListPermissionSetProvisioningsOutput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *ListPermissionSetProvisioningsOutput) SetPageSize(v int32) *ListPermissionSetProvisioningsOutput {
	s.PageSize = &v
	return s
}

// SetPermissionSetProvisionings sets the PermissionSetProvisionings field's value.
func (s *ListPermissionSetProvisioningsOutput) SetPermissionSetProvisionings(v []*PermissionSetProvisioningForListPermissionSetProvisioningsOutput) *ListPermissionSetProvisioningsOutput {
	s.PermissionSetProvisionings = v
	return s
}

// SetTotal sets the Total field's value.
func (s *ListPermissionSetProvisioningsOutput) SetTotal(v int64) *ListPermissionSetProvisioningsOutput {
	s.Total = &v
	return s
}

type PermissionSetProvisioningForListPermissionSetProvisioningsOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	CreatedTime *string `type:"string" json:",omitempty"`

	PermissionSetId *string `type:"string" json:",omitempty"`

	PermissionSetName *string `type:"string" json:",omitempty"`

	ProvisioningStatus *string `type:"string" json:",omitempty"`

	TargetId *string `type:"string" json:",omitempty"`

	TargetName *string `type:"string" json:",omitempty"`

	UpdatedTime *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s PermissionSetProvisioningForListPermissionSetProvisioningsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s PermissionSetProvisioningForListPermissionSetProvisioningsOutput) GoString() string {
	return s.String()
}

// SetCreatedTime sets the CreatedTime field's value.
func (s *PermissionSetProvisioningForListPermissionSetProvisioningsOutput) SetCreatedTime(v string) *PermissionSetProvisioningForListPermissionSetProvisioningsOutput {
	s.CreatedTime = &v
	return s
}

// SetPermissionSetId sets the PermissionSetId field's value.
func (s *PermissionSetProvisioningForListPermissionSetProvisioningsOutput) SetPermissionSetId(v string) *PermissionSetProvisioningForListPermissionSetProvisioningsOutput {
	s.PermissionSetId = &v
	return s
}

// SetPermissionSetName sets the PermissionSetName field's value.
func (s *PermissionSetProvisioningForListPermissionSetProvisioningsOutput) SetPermissionSetName(v string) *PermissionSetProvisioningForListPermissionSetProvisioningsOutput {
	s.PermissionSetName = &v
	return s
}

// SetProvisioningStatus sets the ProvisioningStatus field's value.
func (s *PermissionSetProvisioningForListPermissionSetProvisioningsOutput) SetProvisioningStatus(v string) *PermissionSetProvisioningForListPermissionSetProvisioningsOutput {
	s.ProvisioningStatus = &v
	return s
}

// SetTargetId sets the TargetId field's value.
func (s *PermissionSetProvisioningForListPermissionSetProvisioningsOutput) SetTargetId(v string) *PermissionSetProvisioningForListPermissionSetProvisioningsOutput {
	s.TargetId = &v
	return s
}

// SetTargetName sets the TargetName field's value.
func (s *PermissionSetProvisioningForListPermissionSetProvisioningsOutput) SetTargetName(v string) *PermissionSetProvisioningForListPermissionSetProvisioningsOutput {
	s.TargetName = &v
	return s
}

// SetUpdatedTime sets the UpdatedTime field's value.
func (s *PermissionSetProvisioningForListPermissionSetProvisioningsOutput) SetUpdatedTime(v string) *PermissionSetProvisioningForListPermissionSetProvisioningsOutput {
	s.UpdatedTime = &v
	return s
}

const (
	// EnumOfProvisioningStatusForListPermissionSetProvisioningsInputProvisioned is a EnumOfProvisioningStatusForListPermissionSetProvisioningsInput enum value
	EnumOfProvisioningStatusForListPermissionSetProvisioningsInputProvisioned = "Provisioned"

	// EnumOfProvisioningStatusForListPermissionSetProvisioningsInputReprovisionRequired is a EnumOfProvisioningStatusForListPermissionSetProvisioningsInput enum value
	EnumOfProvisioningStatusForListPermissionSetProvisioningsInputReprovisionRequired = "ReprovisionRequired"

	// EnumOfProvisioningStatusForListPermissionSetProvisioningsInputProvisionFailed is a EnumOfProvisioningStatusForListPermissionSetProvisioningsInput enum value
	EnumOfProvisioningStatusForListPermissionSetProvisioningsInputProvisionFailed = "ProvisionFailed"

	// EnumOfProvisioningStatusForListPermissionSetProvisioningsInputRunning is a EnumOfProvisioningStatusForListPermissionSetProvisioningsInput enum value
	EnumOfProvisioningStatusForListPermissionSetProvisioningsInputRunning = "Running"
)

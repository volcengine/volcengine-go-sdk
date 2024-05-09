// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package organization

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opListOrganizationalUnitsCommon = "ListOrganizationalUnits"

// ListOrganizationalUnitsCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ListOrganizationalUnitsCommon operation. The "output" return
// value will be populated with the ListOrganizationalUnitsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListOrganizationalUnitsCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListOrganizationalUnitsCommon Send returns without error.
//
// See ListOrganizationalUnitsCommon for more information on using the ListOrganizationalUnitsCommon
// API call, and error handling.
//
//    // Example sending a request using the ListOrganizationalUnitsCommonRequest method.
//    req, resp := client.ListOrganizationalUnitsCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ORGANIZATION) ListOrganizationalUnitsCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opListOrganizationalUnitsCommon,
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

// ListOrganizationalUnitsCommon API operation for ORGANIZATION.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ORGANIZATION's
// API operation ListOrganizationalUnitsCommon for usage and error information.
func (c *ORGANIZATION) ListOrganizationalUnitsCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ListOrganizationalUnitsCommonRequest(input)
	return out, req.Send()
}

// ListOrganizationalUnitsCommonWithContext is the same as ListOrganizationalUnitsCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ListOrganizationalUnitsCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ORGANIZATION) ListOrganizationalUnitsCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ListOrganizationalUnitsCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opListOrganizationalUnits = "ListOrganizationalUnits"

// ListOrganizationalUnitsRequest generates a "volcengine/request.Request" representing the
// client's request for the ListOrganizationalUnits operation. The "output" return
// value will be populated with the ListOrganizationalUnitsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListOrganizationalUnitsCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListOrganizationalUnitsCommon Send returns without error.
//
// See ListOrganizationalUnits for more information on using the ListOrganizationalUnits
// API call, and error handling.
//
//    // Example sending a request using the ListOrganizationalUnitsRequest method.
//    req, resp := client.ListOrganizationalUnitsRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ORGANIZATION) ListOrganizationalUnitsRequest(input *ListOrganizationalUnitsInput) (req *request.Request, output *ListOrganizationalUnitsOutput) {
	op := &request.Operation{
		Name:       opListOrganizationalUnits,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListOrganizationalUnitsInput{}
	}

	output = &ListOrganizationalUnitsOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// ListOrganizationalUnits API operation for ORGANIZATION.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ORGANIZATION's
// API operation ListOrganizationalUnits for usage and error information.
func (c *ORGANIZATION) ListOrganizationalUnits(input *ListOrganizationalUnitsInput) (*ListOrganizationalUnitsOutput, error) {
	req, out := c.ListOrganizationalUnitsRequest(input)
	return out, req.Send()
}

// ListOrganizationalUnitsWithContext is the same as ListOrganizationalUnits with the addition of
// the ability to pass a context and additional request options.
//
// See ListOrganizationalUnits for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ORGANIZATION) ListOrganizationalUnitsWithContext(ctx volcengine.Context, input *ListOrganizationalUnitsInput, opts ...request.Option) (*ListOrganizationalUnitsOutput, error) {
	req, out := c.ListOrganizationalUnitsRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ListOrganizationalUnitsInput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s ListOrganizationalUnitsInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListOrganizationalUnitsInput) GoString() string {
	return s.String()
}

type ListOrganizationalUnitsOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	SubUnitList []*SubUnitListForListOrganizationalUnitsOutput `type:"list"`
}

// String returns the string representation
func (s ListOrganizationalUnitsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListOrganizationalUnitsOutput) GoString() string {
	return s.String()
}

// SetSubUnitList sets the SubUnitList field's value.
func (s *ListOrganizationalUnitsOutput) SetSubUnitList(v []*SubUnitListForListOrganizationalUnitsOutput) *ListOrganizationalUnitsOutput {
	s.SubUnitList = v
	return s
}

type SubUnitListForListOrganizationalUnitsOutput struct {
	_ struct{} `type:"structure"`

	CreatedTime *string `type:"string"`

	DeleteUk *string `type:"string"`

	DeletedTime *string `type:"string"`

	Depth *int32 `type:"int32"`

	Description *string `type:"string"`

	ID *string `type:"string"`

	Name *string `type:"string"`

	OrgID *string `type:"string"`

	OrgType *int32 `type:"int32"`

	Owner *string `type:"string"`

	ParentID *string `type:"string"`

	UpdatedTime *string `type:"string"`
}

// String returns the string representation
func (s SubUnitListForListOrganizationalUnitsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s SubUnitListForListOrganizationalUnitsOutput) GoString() string {
	return s.String()
}

// SetCreatedTime sets the CreatedTime field's value.
func (s *SubUnitListForListOrganizationalUnitsOutput) SetCreatedTime(v string) *SubUnitListForListOrganizationalUnitsOutput {
	s.CreatedTime = &v
	return s
}

// SetDeleteUk sets the DeleteUk field's value.
func (s *SubUnitListForListOrganizationalUnitsOutput) SetDeleteUk(v string) *SubUnitListForListOrganizationalUnitsOutput {
	s.DeleteUk = &v
	return s
}

// SetDeletedTime sets the DeletedTime field's value.
func (s *SubUnitListForListOrganizationalUnitsOutput) SetDeletedTime(v string) *SubUnitListForListOrganizationalUnitsOutput {
	s.DeletedTime = &v
	return s
}

// SetDepth sets the Depth field's value.
func (s *SubUnitListForListOrganizationalUnitsOutput) SetDepth(v int32) *SubUnitListForListOrganizationalUnitsOutput {
	s.Depth = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *SubUnitListForListOrganizationalUnitsOutput) SetDescription(v string) *SubUnitListForListOrganizationalUnitsOutput {
	s.Description = &v
	return s
}

// SetID sets the ID field's value.
func (s *SubUnitListForListOrganizationalUnitsOutput) SetID(v string) *SubUnitListForListOrganizationalUnitsOutput {
	s.ID = &v
	return s
}

// SetName sets the Name field's value.
func (s *SubUnitListForListOrganizationalUnitsOutput) SetName(v string) *SubUnitListForListOrganizationalUnitsOutput {
	s.Name = &v
	return s
}

// SetOrgID sets the OrgID field's value.
func (s *SubUnitListForListOrganizationalUnitsOutput) SetOrgID(v string) *SubUnitListForListOrganizationalUnitsOutput {
	s.OrgID = &v
	return s
}

// SetOrgType sets the OrgType field's value.
func (s *SubUnitListForListOrganizationalUnitsOutput) SetOrgType(v int32) *SubUnitListForListOrganizationalUnitsOutput {
	s.OrgType = &v
	return s
}

// SetOwner sets the Owner field's value.
func (s *SubUnitListForListOrganizationalUnitsOutput) SetOwner(v string) *SubUnitListForListOrganizationalUnitsOutput {
	s.Owner = &v
	return s
}

// SetParentID sets the ParentID field's value.
func (s *SubUnitListForListOrganizationalUnitsOutput) SetParentID(v string) *SubUnitListForListOrganizationalUnitsOutput {
	s.ParentID = &v
	return s
}

// SetUpdatedTime sets the UpdatedTime field's value.
func (s *SubUnitListForListOrganizationalUnitsOutput) SetUpdatedTime(v string) *SubUnitListForListOrganizationalUnitsOutput {
	s.UpdatedTime = &v
	return s
}

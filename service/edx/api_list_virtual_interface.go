// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package edx

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opListVirtualInterfaceCommon = "ListVirtualInterface"

// ListVirtualInterfaceCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ListVirtualInterfaceCommon operation. The "output" return
// value will be populated with the ListVirtualInterfaceCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListVirtualInterfaceCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListVirtualInterfaceCommon Send returns without error.
//
// See ListVirtualInterfaceCommon for more information on using the ListVirtualInterfaceCommon
// API call, and error handling.
//
//    // Example sending a request using the ListVirtualInterfaceCommonRequest method.
//    req, resp := client.ListVirtualInterfaceCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *EDX) ListVirtualInterfaceCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opListVirtualInterfaceCommon,
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

// ListVirtualInterfaceCommon API operation for EDX.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for EDX's
// API operation ListVirtualInterfaceCommon for usage and error information.
func (c *EDX) ListVirtualInterfaceCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ListVirtualInterfaceCommonRequest(input)
	return out, req.Send()
}

// ListVirtualInterfaceCommonWithContext is the same as ListVirtualInterfaceCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ListVirtualInterfaceCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *EDX) ListVirtualInterfaceCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ListVirtualInterfaceCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opListVirtualInterface = "ListVirtualInterface"

// ListVirtualInterfaceRequest generates a "volcengine/request.Request" representing the
// client's request for the ListVirtualInterface operation. The "output" return
// value will be populated with the ListVirtualInterfaceCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListVirtualInterfaceCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListVirtualInterfaceCommon Send returns without error.
//
// See ListVirtualInterface for more information on using the ListVirtualInterface
// API call, and error handling.
//
//    // Example sending a request using the ListVirtualInterfaceRequest method.
//    req, resp := client.ListVirtualInterfaceRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *EDX) ListVirtualInterfaceRequest(input *ListVirtualInterfaceInput) (req *request.Request, output *ListVirtualInterfaceOutput) {
	op := &request.Operation{
		Name:       opListVirtualInterface,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListVirtualInterfaceInput{}
	}

	output = &ListVirtualInterfaceOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// ListVirtualInterface API operation for EDX.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for EDX's
// API operation ListVirtualInterface for usage and error information.
func (c *EDX) ListVirtualInterface(input *ListVirtualInterfaceInput) (*ListVirtualInterfaceOutput, error) {
	req, out := c.ListVirtualInterfaceRequest(input)
	return out, req.Send()
}

// ListVirtualInterfaceWithContext is the same as ListVirtualInterface with the addition of
// the ability to pass a context and additional request options.
//
// See ListVirtualInterface for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *EDX) ListVirtualInterfaceWithContext(ctx volcengine.Context, input *ListVirtualInterfaceInput, opts ...request.Option) (*ListVirtualInterfaceOutput, error) {
	req, out := c.ListVirtualInterfaceRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ListVirtualInterfaceInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	DXPInstanceID *string `type:"string" json:",omitempty"`

	PageNumber *int32 `type:"int32" json:",omitempty"`

	PageSize *int32 `type:"int32" json:",omitempty"`

	SortBy *string `type:"string" json:",omitempty"`

	SortOrder *string `type:"string" json:",omitempty"`

	State *string `type:"string" json:",omitempty"`

	VIFInstanceID *string `type:"string" json:",omitempty"`

	VIFVGWInstanceID *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s ListVirtualInterfaceInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListVirtualInterfaceInput) GoString() string {
	return s.String()
}

// SetDXPInstanceID sets the DXPInstanceID field's value.
func (s *ListVirtualInterfaceInput) SetDXPInstanceID(v string) *ListVirtualInterfaceInput {
	s.DXPInstanceID = &v
	return s
}

// SetPageNumber sets the PageNumber field's value.
func (s *ListVirtualInterfaceInput) SetPageNumber(v int32) *ListVirtualInterfaceInput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *ListVirtualInterfaceInput) SetPageSize(v int32) *ListVirtualInterfaceInput {
	s.PageSize = &v
	return s
}

// SetSortBy sets the SortBy field's value.
func (s *ListVirtualInterfaceInput) SetSortBy(v string) *ListVirtualInterfaceInput {
	s.SortBy = &v
	return s
}

// SetSortOrder sets the SortOrder field's value.
func (s *ListVirtualInterfaceInput) SetSortOrder(v string) *ListVirtualInterfaceInput {
	s.SortOrder = &v
	return s
}

// SetState sets the State field's value.
func (s *ListVirtualInterfaceInput) SetState(v string) *ListVirtualInterfaceInput {
	s.State = &v
	return s
}

// SetVIFInstanceID sets the VIFInstanceID field's value.
func (s *ListVirtualInterfaceInput) SetVIFInstanceID(v string) *ListVirtualInterfaceInput {
	s.VIFInstanceID = &v
	return s
}

// SetVIFVGWInstanceID sets the VIFVGWInstanceID field's value.
func (s *ListVirtualInterfaceInput) SetVIFVGWInstanceID(v string) *ListVirtualInterfaceInput {
	s.VIFVGWInstanceID = &v
	return s
}

type ListVirtualInterfaceOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	PageNumber *int32 `type:"int32" json:",omitempty"`

	PageSize *int32 `type:"int32" json:",omitempty"`

	TotalCount *int32 `type:"int32" json:",omitempty"`

	VIFList []*VIFListForListVirtualInterfaceOutput `type:"list" json:",omitempty"`
}

// String returns the string representation
func (s ListVirtualInterfaceOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListVirtualInterfaceOutput) GoString() string {
	return s.String()
}

// SetPageNumber sets the PageNumber field's value.
func (s *ListVirtualInterfaceOutput) SetPageNumber(v int32) *ListVirtualInterfaceOutput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *ListVirtualInterfaceOutput) SetPageSize(v int32) *ListVirtualInterfaceOutput {
	s.PageSize = &v
	return s
}

// SetTotalCount sets the TotalCount field's value.
func (s *ListVirtualInterfaceOutput) SetTotalCount(v int32) *ListVirtualInterfaceOutput {
	s.TotalCount = &v
	return s
}

// SetVIFList sets the VIFList field's value.
func (s *ListVirtualInterfaceOutput) SetVIFList(v []*VIFListForListVirtualInterfaceOutput) *ListVirtualInterfaceOutput {
	s.VIFList = v
	return s
}

type VIFListForListVirtualInterfaceOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	AccessPoint *string `type:"string" json:",omitempty"`

	AccountId *int32 `type:"int32" json:",omitempty"`

	CreateTime *string `type:"string" json:",omitempty"`

	CrossAccount *bool `type:"boolean" json:",omitempty"`

	DXPInstanceId *string `type:"string" json:",omitempty"`

	DXPInstanceName *string `type:"string" json:",omitempty"`

	InstanceId *string `type:"string" json:",omitempty"`

	LocalIP *string `type:"string" json:",omitempty"`

	MaxBandwidth *int32 `type:"int32" json:",omitempty"`

	Name *string `type:"string" json:",omitempty"`

	PeerIP *string `type:"string" json:",omitempty"`

	State *string `type:"string" json:",omitempty"`

	VIFVGWInstanceId *string `type:"string" json:",omitempty"`

	VlanId *int32 `type:"int32" json:",omitempty"`
}

// String returns the string representation
func (s VIFListForListVirtualInterfaceOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s VIFListForListVirtualInterfaceOutput) GoString() string {
	return s.String()
}

// SetAccessPoint sets the AccessPoint field's value.
func (s *VIFListForListVirtualInterfaceOutput) SetAccessPoint(v string) *VIFListForListVirtualInterfaceOutput {
	s.AccessPoint = &v
	return s
}

// SetAccountId sets the AccountId field's value.
func (s *VIFListForListVirtualInterfaceOutput) SetAccountId(v int32) *VIFListForListVirtualInterfaceOutput {
	s.AccountId = &v
	return s
}

// SetCreateTime sets the CreateTime field's value.
func (s *VIFListForListVirtualInterfaceOutput) SetCreateTime(v string) *VIFListForListVirtualInterfaceOutput {
	s.CreateTime = &v
	return s
}

// SetCrossAccount sets the CrossAccount field's value.
func (s *VIFListForListVirtualInterfaceOutput) SetCrossAccount(v bool) *VIFListForListVirtualInterfaceOutput {
	s.CrossAccount = &v
	return s
}

// SetDXPInstanceId sets the DXPInstanceId field's value.
func (s *VIFListForListVirtualInterfaceOutput) SetDXPInstanceId(v string) *VIFListForListVirtualInterfaceOutput {
	s.DXPInstanceId = &v
	return s
}

// SetDXPInstanceName sets the DXPInstanceName field's value.
func (s *VIFListForListVirtualInterfaceOutput) SetDXPInstanceName(v string) *VIFListForListVirtualInterfaceOutput {
	s.DXPInstanceName = &v
	return s
}

// SetInstanceId sets the InstanceId field's value.
func (s *VIFListForListVirtualInterfaceOutput) SetInstanceId(v string) *VIFListForListVirtualInterfaceOutput {
	s.InstanceId = &v
	return s
}

// SetLocalIP sets the LocalIP field's value.
func (s *VIFListForListVirtualInterfaceOutput) SetLocalIP(v string) *VIFListForListVirtualInterfaceOutput {
	s.LocalIP = &v
	return s
}

// SetMaxBandwidth sets the MaxBandwidth field's value.
func (s *VIFListForListVirtualInterfaceOutput) SetMaxBandwidth(v int32) *VIFListForListVirtualInterfaceOutput {
	s.MaxBandwidth = &v
	return s
}

// SetName sets the Name field's value.
func (s *VIFListForListVirtualInterfaceOutput) SetName(v string) *VIFListForListVirtualInterfaceOutput {
	s.Name = &v
	return s
}

// SetPeerIP sets the PeerIP field's value.
func (s *VIFListForListVirtualInterfaceOutput) SetPeerIP(v string) *VIFListForListVirtualInterfaceOutput {
	s.PeerIP = &v
	return s
}

// SetState sets the State field's value.
func (s *VIFListForListVirtualInterfaceOutput) SetState(v string) *VIFListForListVirtualInterfaceOutput {
	s.State = &v
	return s
}

// SetVIFVGWInstanceId sets the VIFVGWInstanceId field's value.
func (s *VIFListForListVirtualInterfaceOutput) SetVIFVGWInstanceId(v string) *VIFListForListVirtualInterfaceOutput {
	s.VIFVGWInstanceId = &v
	return s
}

// SetVlanId sets the VlanId field's value.
func (s *VIFListForListVirtualInterfaceOutput) SetVlanId(v int32) *VIFListForListVirtualInterfaceOutput {
	s.VlanId = &v
	return s
}

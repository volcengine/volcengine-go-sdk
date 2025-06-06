// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package edx

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opListEDXBandwidthPkgCommon = "ListEDXBandwidthPkg"

// ListEDXBandwidthPkgCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ListEDXBandwidthPkgCommon operation. The "output" return
// value will be populated with the ListEDXBandwidthPkgCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListEDXBandwidthPkgCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListEDXBandwidthPkgCommon Send returns without error.
//
// See ListEDXBandwidthPkgCommon for more information on using the ListEDXBandwidthPkgCommon
// API call, and error handling.
//
//    // Example sending a request using the ListEDXBandwidthPkgCommonRequest method.
//    req, resp := client.ListEDXBandwidthPkgCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *EDX) ListEDXBandwidthPkgCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opListEDXBandwidthPkgCommon,
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

// ListEDXBandwidthPkgCommon API operation for EDX.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for EDX's
// API operation ListEDXBandwidthPkgCommon for usage and error information.
func (c *EDX) ListEDXBandwidthPkgCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ListEDXBandwidthPkgCommonRequest(input)
	return out, req.Send()
}

// ListEDXBandwidthPkgCommonWithContext is the same as ListEDXBandwidthPkgCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ListEDXBandwidthPkgCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *EDX) ListEDXBandwidthPkgCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ListEDXBandwidthPkgCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opListEDXBandwidthPkg = "ListEDXBandwidthPkg"

// ListEDXBandwidthPkgRequest generates a "volcengine/request.Request" representing the
// client's request for the ListEDXBandwidthPkg operation. The "output" return
// value will be populated with the ListEDXBandwidthPkgCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListEDXBandwidthPkgCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListEDXBandwidthPkgCommon Send returns without error.
//
// See ListEDXBandwidthPkg for more information on using the ListEDXBandwidthPkg
// API call, and error handling.
//
//    // Example sending a request using the ListEDXBandwidthPkgRequest method.
//    req, resp := client.ListEDXBandwidthPkgRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *EDX) ListEDXBandwidthPkgRequest(input *ListEDXBandwidthPkgInput) (req *request.Request, output *ListEDXBandwidthPkgOutput) {
	op := &request.Operation{
		Name:       opListEDXBandwidthPkg,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListEDXBandwidthPkgInput{}
	}

	output = &ListEDXBandwidthPkgOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// ListEDXBandwidthPkg API operation for EDX.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for EDX's
// API operation ListEDXBandwidthPkg for usage and error information.
func (c *EDX) ListEDXBandwidthPkg(input *ListEDXBandwidthPkgInput) (*ListEDXBandwidthPkgOutput, error) {
	req, out := c.ListEDXBandwidthPkgRequest(input)
	return out, req.Send()
}

// ListEDXBandwidthPkgWithContext is the same as ListEDXBandwidthPkg with the addition of
// the ability to pass a context and additional request options.
//
// See ListEDXBandwidthPkg for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *EDX) ListEDXBandwidthPkgWithContext(ctx volcengine.Context, input *ListEDXBandwidthPkgInput, opts ...request.Option) (*ListEDXBandwidthPkgOutput, error) {
	req, out := c.ListEDXBandwidthPkgRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type BandwidthPkgListForListEDXBandwidthPkgOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	AllocatedBandwidth *int32 `type:"int32" json:",omitempty"`

	BandwidthPkgId *string `type:"string" json:",omitempty"`

	BillingMode *string `type:"string" json:",omitempty"`

	ConnectArea *string `type:"string" json:",omitempty"`

	CreateTime *string `type:"string" json:",omitempty"`

	EDXInstanceId *string `type:"string" json:",omitempty"`

	Name *string `type:"string" json:",omitempty"`

	PaidMode *string `type:"string" json:",omitempty"`

	State *string `type:"string" json:",omitempty"`

	TotalBandwidth *int32 `type:"int32" json:",omitempty"`
}

// String returns the string representation
func (s BandwidthPkgListForListEDXBandwidthPkgOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s BandwidthPkgListForListEDXBandwidthPkgOutput) GoString() string {
	return s.String()
}

// SetAllocatedBandwidth sets the AllocatedBandwidth field's value.
func (s *BandwidthPkgListForListEDXBandwidthPkgOutput) SetAllocatedBandwidth(v int32) *BandwidthPkgListForListEDXBandwidthPkgOutput {
	s.AllocatedBandwidth = &v
	return s
}

// SetBandwidthPkgId sets the BandwidthPkgId field's value.
func (s *BandwidthPkgListForListEDXBandwidthPkgOutput) SetBandwidthPkgId(v string) *BandwidthPkgListForListEDXBandwidthPkgOutput {
	s.BandwidthPkgId = &v
	return s
}

// SetBillingMode sets the BillingMode field's value.
func (s *BandwidthPkgListForListEDXBandwidthPkgOutput) SetBillingMode(v string) *BandwidthPkgListForListEDXBandwidthPkgOutput {
	s.BillingMode = &v
	return s
}

// SetConnectArea sets the ConnectArea field's value.
func (s *BandwidthPkgListForListEDXBandwidthPkgOutput) SetConnectArea(v string) *BandwidthPkgListForListEDXBandwidthPkgOutput {
	s.ConnectArea = &v
	return s
}

// SetCreateTime sets the CreateTime field's value.
func (s *BandwidthPkgListForListEDXBandwidthPkgOutput) SetCreateTime(v string) *BandwidthPkgListForListEDXBandwidthPkgOutput {
	s.CreateTime = &v
	return s
}

// SetEDXInstanceId sets the EDXInstanceId field's value.
func (s *BandwidthPkgListForListEDXBandwidthPkgOutput) SetEDXInstanceId(v string) *BandwidthPkgListForListEDXBandwidthPkgOutput {
	s.EDXInstanceId = &v
	return s
}

// SetName sets the Name field's value.
func (s *BandwidthPkgListForListEDXBandwidthPkgOutput) SetName(v string) *BandwidthPkgListForListEDXBandwidthPkgOutput {
	s.Name = &v
	return s
}

// SetPaidMode sets the PaidMode field's value.
func (s *BandwidthPkgListForListEDXBandwidthPkgOutput) SetPaidMode(v string) *BandwidthPkgListForListEDXBandwidthPkgOutput {
	s.PaidMode = &v
	return s
}

// SetState sets the State field's value.
func (s *BandwidthPkgListForListEDXBandwidthPkgOutput) SetState(v string) *BandwidthPkgListForListEDXBandwidthPkgOutput {
	s.State = &v
	return s
}

// SetTotalBandwidth sets the TotalBandwidth field's value.
func (s *BandwidthPkgListForListEDXBandwidthPkgOutput) SetTotalBandwidth(v int32) *BandwidthPkgListForListEDXBandwidthPkgOutput {
	s.TotalBandwidth = &v
	return s
}

type ListEDXBandwidthPkgInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	BandwidthPkgId *string `type:"string" json:",omitempty"`

	// EDXInstanceId is a required field
	EDXInstanceId *string `type:"string" json:",omitempty" required:"true"`

	Name *string `type:"string" json:",omitempty"`

	PageNumber *string `type:"string" json:",omitempty"`

	PageSize *int32 `type:"int32" json:",omitempty"`

	State *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s ListEDXBandwidthPkgInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListEDXBandwidthPkgInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListEDXBandwidthPkgInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ListEDXBandwidthPkgInput"}
	if s.EDXInstanceId == nil {
		invalidParams.Add(request.NewErrParamRequired("EDXInstanceId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetBandwidthPkgId sets the BandwidthPkgId field's value.
func (s *ListEDXBandwidthPkgInput) SetBandwidthPkgId(v string) *ListEDXBandwidthPkgInput {
	s.BandwidthPkgId = &v
	return s
}

// SetEDXInstanceId sets the EDXInstanceId field's value.
func (s *ListEDXBandwidthPkgInput) SetEDXInstanceId(v string) *ListEDXBandwidthPkgInput {
	s.EDXInstanceId = &v
	return s
}

// SetName sets the Name field's value.
func (s *ListEDXBandwidthPkgInput) SetName(v string) *ListEDXBandwidthPkgInput {
	s.Name = &v
	return s
}

// SetPageNumber sets the PageNumber field's value.
func (s *ListEDXBandwidthPkgInput) SetPageNumber(v string) *ListEDXBandwidthPkgInput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *ListEDXBandwidthPkgInput) SetPageSize(v int32) *ListEDXBandwidthPkgInput {
	s.PageSize = &v
	return s
}

// SetState sets the State field's value.
func (s *ListEDXBandwidthPkgInput) SetState(v string) *ListEDXBandwidthPkgInput {
	s.State = &v
	return s
}

type ListEDXBandwidthPkgOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	BandwidthPkgList []*BandwidthPkgListForListEDXBandwidthPkgOutput `type:"list" json:",omitempty"`

	PageNumber *int32 `type:"int32" json:",omitempty"`

	PageSize *int32 `type:"int32" json:",omitempty"`

	TotalCount *int32 `type:"int32" json:",omitempty"`
}

// String returns the string representation
func (s ListEDXBandwidthPkgOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListEDXBandwidthPkgOutput) GoString() string {
	return s.String()
}

// SetBandwidthPkgList sets the BandwidthPkgList field's value.
func (s *ListEDXBandwidthPkgOutput) SetBandwidthPkgList(v []*BandwidthPkgListForListEDXBandwidthPkgOutput) *ListEDXBandwidthPkgOutput {
	s.BandwidthPkgList = v
	return s
}

// SetPageNumber sets the PageNumber field's value.
func (s *ListEDXBandwidthPkgOutput) SetPageNumber(v int32) *ListEDXBandwidthPkgOutput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *ListEDXBandwidthPkgOutput) SetPageSize(v int32) *ListEDXBandwidthPkgOutput {
	s.PageSize = &v
	return s
}

// SetTotalCount sets the TotalCount field's value.
func (s *ListEDXBandwidthPkgOutput) SetTotalCount(v int32) *ListEDXBandwidthPkgOutput {
	s.TotalCount = &v
	return s
}

// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package mcdn

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribeContentQuotaCommon = "DescribeContentQuota"

// DescribeContentQuotaCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeContentQuotaCommon operation. The "output" return
// value will be populated with the DescribeContentQuotaCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeContentQuotaCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeContentQuotaCommon Send returns without error.
//
// See DescribeContentQuotaCommon for more information on using the DescribeContentQuotaCommon
// API call, and error handling.
//
//    // Example sending a request using the DescribeContentQuotaCommonRequest method.
//    req, resp := client.DescribeContentQuotaCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *MCDN) DescribeContentQuotaCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeContentQuotaCommon,
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

// DescribeContentQuotaCommon API operation for MCDN.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for MCDN's
// API operation DescribeContentQuotaCommon for usage and error information.
func (c *MCDN) DescribeContentQuotaCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeContentQuotaCommonRequest(input)
	return out, req.Send()
}

// DescribeContentQuotaCommonWithContext is the same as DescribeContentQuotaCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeContentQuotaCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *MCDN) DescribeContentQuotaCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeContentQuotaCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeContentQuota = "DescribeContentQuota"

// DescribeContentQuotaRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeContentQuota operation. The "output" return
// value will be populated with the DescribeContentQuotaCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeContentQuotaCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeContentQuotaCommon Send returns without error.
//
// See DescribeContentQuota for more information on using the DescribeContentQuota
// API call, and error handling.
//
//    // Example sending a request using the DescribeContentQuotaRequest method.
//    req, resp := client.DescribeContentQuotaRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *MCDN) DescribeContentQuotaRequest(input *DescribeContentQuotaInput) (req *request.Request, output *DescribeContentQuotaOutput) {
	op := &request.Operation{
		Name:       opDescribeContentQuota,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeContentQuotaInput{}
	}

	output = &DescribeContentQuotaOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// DescribeContentQuota API operation for MCDN.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for MCDN's
// API operation DescribeContentQuota for usage and error information.
func (c *MCDN) DescribeContentQuota(input *DescribeContentQuotaInput) (*DescribeContentQuotaOutput, error) {
	req, out := c.DescribeContentQuotaRequest(input)
	return out, req.Send()
}

// DescribeContentQuotaWithContext is the same as DescribeContentQuota with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeContentQuota for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *MCDN) DescribeContentQuotaWithContext(ctx volcengine.Context, input *DescribeContentQuotaInput, opts ...request.Option) (*DescribeContentQuotaOutput, error) {
	req, out := c.DescribeContentQuotaRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DescribeContentQuotaInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	// CloudAccountId is a required field
	CloudAccountId *string `type:"string" json:",omitempty" required:"true"`
}

// String returns the string representation
func (s DescribeContentQuotaInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeContentQuotaInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeContentQuotaInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DescribeContentQuotaInput"}
	if s.CloudAccountId == nil {
		invalidParams.Add(request.NewErrParamRequired("CloudAccountId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetCloudAccountId sets the CloudAccountId field's value.
func (s *DescribeContentQuotaInput) SetCloudAccountId(v string) *DescribeContentQuotaInput {
	s.CloudAccountId = &v
	return s
}

type DescribeContentQuotaOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	Quotas []*QuotaForDescribeContentQuotaOutput `type:"list" json:",omitempty"`

	VendorsMetaData []*VendorsMetaDataForDescribeContentQuotaOutput `type:"list" json:",omitempty"`
}

// String returns the string representation
func (s DescribeContentQuotaOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeContentQuotaOutput) GoString() string {
	return s.String()
}

// SetQuotas sets the Quotas field's value.
func (s *DescribeContentQuotaOutput) SetQuotas(v []*QuotaForDescribeContentQuotaOutput) *DescribeContentQuotaOutput {
	s.Quotas = v
	return s
}

// SetVendorsMetaData sets the VendorsMetaData field's value.
func (s *DescribeContentQuotaOutput) SetVendorsMetaData(v []*VendorsMetaDataForDescribeContentQuotaOutput) *DescribeContentQuotaOutput {
	s.VendorsMetaData = v
	return s
}

type ErrorForDescribeContentQuotaOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Code *string `type:"string" json:",omitempty"`

	Detail *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s ErrorForDescribeContentQuotaOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ErrorForDescribeContentQuotaOutput) GoString() string {
	return s.String()
}

// SetCode sets the Code field's value.
func (s *ErrorForDescribeContentQuotaOutput) SetCode(v string) *ErrorForDescribeContentQuotaOutput {
	s.Code = &v
	return s
}

// SetDetail sets the Detail field's value.
func (s *ErrorForDescribeContentQuotaOutput) SetDetail(v string) *ErrorForDescribeContentQuotaOutput {
	s.Detail = &v
	return s
}

type QuotaForDescribeContentQuotaOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Area *string `type:"string" json:",omitempty"`

	CloudAccountId *string `type:"string" json:",omitempty"`

	CloudAccountName *string `type:"string" json:",omitempty"`

	PreloadQuota *int64 `type:"int64" json:",omitempty"`

	PreloadRemain *int64 `type:"int64" json:",omitempty"`

	ProductType *string `type:"string" json:",omitempty"`

	RefreshDirQuota *int64 `type:"int64" json:",omitempty"`

	RefreshDirRemain *int64 `type:"int64" json:",omitempty"`

	RefreshUrlQuota *int64 `type:"int64" json:",omitempty"`

	RefreshUrlRemain *int64 `type:"int64" json:",omitempty"`

	SubProduct *string `type:"string" json:",omitempty"`

	Vendor *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s QuotaForDescribeContentQuotaOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s QuotaForDescribeContentQuotaOutput) GoString() string {
	return s.String()
}

// SetArea sets the Area field's value.
func (s *QuotaForDescribeContentQuotaOutput) SetArea(v string) *QuotaForDescribeContentQuotaOutput {
	s.Area = &v
	return s
}

// SetCloudAccountId sets the CloudAccountId field's value.
func (s *QuotaForDescribeContentQuotaOutput) SetCloudAccountId(v string) *QuotaForDescribeContentQuotaOutput {
	s.CloudAccountId = &v
	return s
}

// SetCloudAccountName sets the CloudAccountName field's value.
func (s *QuotaForDescribeContentQuotaOutput) SetCloudAccountName(v string) *QuotaForDescribeContentQuotaOutput {
	s.CloudAccountName = &v
	return s
}

// SetPreloadQuota sets the PreloadQuota field's value.
func (s *QuotaForDescribeContentQuotaOutput) SetPreloadQuota(v int64) *QuotaForDescribeContentQuotaOutput {
	s.PreloadQuota = &v
	return s
}

// SetPreloadRemain sets the PreloadRemain field's value.
func (s *QuotaForDescribeContentQuotaOutput) SetPreloadRemain(v int64) *QuotaForDescribeContentQuotaOutput {
	s.PreloadRemain = &v
	return s
}

// SetProductType sets the ProductType field's value.
func (s *QuotaForDescribeContentQuotaOutput) SetProductType(v string) *QuotaForDescribeContentQuotaOutput {
	s.ProductType = &v
	return s
}

// SetRefreshDirQuota sets the RefreshDirQuota field's value.
func (s *QuotaForDescribeContentQuotaOutput) SetRefreshDirQuota(v int64) *QuotaForDescribeContentQuotaOutput {
	s.RefreshDirQuota = &v
	return s
}

// SetRefreshDirRemain sets the RefreshDirRemain field's value.
func (s *QuotaForDescribeContentQuotaOutput) SetRefreshDirRemain(v int64) *QuotaForDescribeContentQuotaOutput {
	s.RefreshDirRemain = &v
	return s
}

// SetRefreshUrlQuota sets the RefreshUrlQuota field's value.
func (s *QuotaForDescribeContentQuotaOutput) SetRefreshUrlQuota(v int64) *QuotaForDescribeContentQuotaOutput {
	s.RefreshUrlQuota = &v
	return s
}

// SetRefreshUrlRemain sets the RefreshUrlRemain field's value.
func (s *QuotaForDescribeContentQuotaOutput) SetRefreshUrlRemain(v int64) *QuotaForDescribeContentQuotaOutput {
	s.RefreshUrlRemain = &v
	return s
}

// SetSubProduct sets the SubProduct field's value.
func (s *QuotaForDescribeContentQuotaOutput) SetSubProduct(v string) *QuotaForDescribeContentQuotaOutput {
	s.SubProduct = &v
	return s
}

// SetVendor sets the Vendor field's value.
func (s *QuotaForDescribeContentQuotaOutput) SetVendor(v string) *QuotaForDescribeContentQuotaOutput {
	s.Vendor = &v
	return s
}

type VendorsMetaDataForDescribeContentQuotaOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	CloudAccountId *string `type:"string" json:",omitempty"`

	Cost *float64 `type:"double" json:",omitempty"`

	Error *ErrorForDescribeContentQuotaOutput `type:"structure" json:",omitempty"`

	ProductType *string `type:"string" json:",omitempty"`

	RequestId *string `type:"string" json:",omitempty"`

	Vendor *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s VendorsMetaDataForDescribeContentQuotaOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s VendorsMetaDataForDescribeContentQuotaOutput) GoString() string {
	return s.String()
}

// SetCloudAccountId sets the CloudAccountId field's value.
func (s *VendorsMetaDataForDescribeContentQuotaOutput) SetCloudAccountId(v string) *VendorsMetaDataForDescribeContentQuotaOutput {
	s.CloudAccountId = &v
	return s
}

// SetCost sets the Cost field's value.
func (s *VendorsMetaDataForDescribeContentQuotaOutput) SetCost(v float64) *VendorsMetaDataForDescribeContentQuotaOutput {
	s.Cost = &v
	return s
}

// SetError sets the Error field's value.
func (s *VendorsMetaDataForDescribeContentQuotaOutput) SetError(v *ErrorForDescribeContentQuotaOutput) *VendorsMetaDataForDescribeContentQuotaOutput {
	s.Error = v
	return s
}

// SetProductType sets the ProductType field's value.
func (s *VendorsMetaDataForDescribeContentQuotaOutput) SetProductType(v string) *VendorsMetaDataForDescribeContentQuotaOutput {
	s.ProductType = &v
	return s
}

// SetRequestId sets the RequestId field's value.
func (s *VendorsMetaDataForDescribeContentQuotaOutput) SetRequestId(v string) *VendorsMetaDataForDescribeContentQuotaOutput {
	s.RequestId = &v
	return s
}

// SetVendor sets the Vendor field's value.
func (s *VendorsMetaDataForDescribeContentQuotaOutput) SetVendor(v string) *VendorsMetaDataForDescribeContentQuotaOutput {
	s.Vendor = &v
	return s
}

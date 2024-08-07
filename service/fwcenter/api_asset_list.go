// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package fwcenter

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opAssetListCommon = "AssetList"

// AssetListCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the AssetListCommon operation. The "output" return
// value will be populated with the AssetListCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned AssetListCommon Request to send the API call to the service.
// the "output" return value is not valid until after AssetListCommon Send returns without error.
//
// See AssetListCommon for more information on using the AssetListCommon
// API call, and error handling.
//
//    // Example sending a request using the AssetListCommonRequest method.
//    req, resp := client.AssetListCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *FWCENTER) AssetListCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opAssetListCommon,
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

// AssetListCommon API operation for FWCENTER.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for FWCENTER's
// API operation AssetListCommon for usage and error information.
func (c *FWCENTER) AssetListCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.AssetListCommonRequest(input)
	return out, req.Send()
}

// AssetListCommonWithContext is the same as AssetListCommon with the addition of
// the ability to pass a context and additional request options.
//
// See AssetListCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *FWCENTER) AssetListCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.AssetListCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opAssetList = "AssetList"

// AssetListRequest generates a "volcengine/request.Request" representing the
// client's request for the AssetList operation. The "output" return
// value will be populated with the AssetListCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned AssetListCommon Request to send the API call to the service.
// the "output" return value is not valid until after AssetListCommon Send returns without error.
//
// See AssetList for more information on using the AssetList
// API call, and error handling.
//
//    // Example sending a request using the AssetListRequest method.
//    req, resp := client.AssetListRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *FWCENTER) AssetListRequest(input *AssetListInput) (req *request.Request, output *AssetListOutput) {
	op := &request.Operation{
		Name:       opAssetList,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &AssetListInput{}
	}

	output = &AssetListOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// AssetList API operation for FWCENTER.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for FWCENTER's
// API operation AssetList for usage and error information.
func (c *FWCENTER) AssetList(input *AssetListInput) (*AssetListOutput, error) {
	req, out := c.AssetListRequest(input)
	return out, req.Send()
}

// AssetListWithContext is the same as AssetList with the addition of
// the ability to pass a context and additional request options.
//
// See AssetList for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *FWCENTER) AssetListWithContext(ctx volcengine.Context, input *AssetListInput, opts ...request.Option) (*AssetListOutput, error) {
	req, out := c.AssetListRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type AssetListInput struct {
	_ struct{} `type:"structure"`

	Asset *string `type:"string" json:"asset"`

	Asset_type []*string `type:"list" json:"asset_type"`

	Current_page *int32 `type:"int32" json:"current_page"`

	Page_size *int32 `type:"int32" json:"page_size"`

	// Stat is a required field
	Stat *int32 `type:"int32" json:"stat" required:"true"`
}

// String returns the string representation
func (s AssetListInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s AssetListInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AssetListInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "AssetListInput"}
	if s.Stat == nil {
		invalidParams.Add(request.NewErrParamRequired("Stat"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetAsset sets the Asset field's value.
func (s *AssetListInput) SetAsset(v string) *AssetListInput {
	s.Asset = &v
	return s
}

// SetAsset_type sets the Asset_type field's value.
func (s *AssetListInput) SetAsset_type(v []*string) *AssetListInput {
	s.Asset_type = v
	return s
}

// SetCurrent_page sets the Current_page field's value.
func (s *AssetListInput) SetCurrent_page(v int32) *AssetListInput {
	s.Current_page = &v
	return s
}

// SetPage_size sets the Page_size field's value.
func (s *AssetListInput) SetPage_size(v int32) *AssetListInput {
	s.Page_size = &v
	return s
}

// SetStat sets the Stat field's value.
func (s *AssetListInput) SetStat(v int32) *AssetListInput {
	s.Stat = &v
	return s
}

type AssetListOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	Count *int32 `type:"int32"`

	Data []*DataForAssetListOutput `type:"list"`

	PageNumber *int32 `type:"int32"`

	PageSize *int32 `type:"int32"`

	TotalCount *int32 `type:"int32"`
}

// String returns the string representation
func (s AssetListOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s AssetListOutput) GoString() string {
	return s.String()
}

// SetCount sets the Count field's value.
func (s *AssetListOutput) SetCount(v int32) *AssetListOutput {
	s.Count = &v
	return s
}

// SetData sets the Data field's value.
func (s *AssetListOutput) SetData(v []*DataForAssetListOutput) *AssetListOutput {
	s.Data = v
	return s
}

// SetPageNumber sets the PageNumber field's value.
func (s *AssetListOutput) SetPageNumber(v int32) *AssetListOutput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *AssetListOutput) SetPageSize(v int32) *AssetListOutput {
	s.PageSize = &v
	return s
}

// SetTotalCount sets the TotalCount field's value.
func (s *AssetListOutput) SetTotalCount(v int32) *AssetListOutput {
	s.TotalCount = &v
	return s
}

type DataForAssetListOutput struct {
	_ struct{} `type:"structure"`

	Account_id *string `type:"string" json:"account_id"`

	Cluster *int32 `type:"int32" json:"cluster"`

	Enable *bool `type:"boolean" json:"enable"`

	Id *int32 `type:"int32" json:"id"`

	Ip *string `type:"string" json:"ip"`

	Name *string `type:"string" json:"name"`

	Region *string `type:"string" json:"region"`

	Type *string `type:"string" json:"type"`
}

// String returns the string representation
func (s DataForAssetListOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DataForAssetListOutput) GoString() string {
	return s.String()
}

// SetAccount_id sets the Account_id field's value.
func (s *DataForAssetListOutput) SetAccount_id(v string) *DataForAssetListOutput {
	s.Account_id = &v
	return s
}

// SetCluster sets the Cluster field's value.
func (s *DataForAssetListOutput) SetCluster(v int32) *DataForAssetListOutput {
	s.Cluster = &v
	return s
}

// SetEnable sets the Enable field's value.
func (s *DataForAssetListOutput) SetEnable(v bool) *DataForAssetListOutput {
	s.Enable = &v
	return s
}

// SetId sets the Id field's value.
func (s *DataForAssetListOutput) SetId(v int32) *DataForAssetListOutput {
	s.Id = &v
	return s
}

// SetIp sets the Ip field's value.
func (s *DataForAssetListOutput) SetIp(v string) *DataForAssetListOutput {
	s.Ip = &v
	return s
}

// SetName sets the Name field's value.
func (s *DataForAssetListOutput) SetName(v string) *DataForAssetListOutput {
	s.Name = &v
	return s
}

// SetRegion sets the Region field's value.
func (s *DataForAssetListOutput) SetRegion(v string) *DataForAssetListOutput {
	s.Region = &v
	return s
}

// SetType sets the Type field's value.
func (s *DataForAssetListOutput) SetType(v string) *DataForAssetListOutput {
	s.Type = &v
	return s
}

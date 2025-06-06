// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package seccenter20240508

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opListCleanHistoryCommon = "ListCleanHistory"

// ListCleanHistoryCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ListCleanHistoryCommon operation. The "output" return
// value will be populated with the ListCleanHistoryCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListCleanHistoryCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListCleanHistoryCommon Send returns without error.
//
// See ListCleanHistoryCommon for more information on using the ListCleanHistoryCommon
// API call, and error handling.
//
//    // Example sending a request using the ListCleanHistoryCommonRequest method.
//    req, resp := client.ListCleanHistoryCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *SECCENTER20240508) ListCleanHistoryCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opListCleanHistoryCommon,
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

// ListCleanHistoryCommon API operation for SECCENTER20240508.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for SECCENTER20240508's
// API operation ListCleanHistoryCommon for usage and error information.
func (c *SECCENTER20240508) ListCleanHistoryCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ListCleanHistoryCommonRequest(input)
	return out, req.Send()
}

// ListCleanHistoryCommonWithContext is the same as ListCleanHistoryCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ListCleanHistoryCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *SECCENTER20240508) ListCleanHistoryCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ListCleanHistoryCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opListCleanHistory = "ListCleanHistory"

// ListCleanHistoryRequest generates a "volcengine/request.Request" representing the
// client's request for the ListCleanHistory operation. The "output" return
// value will be populated with the ListCleanHistoryCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListCleanHistoryCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListCleanHistoryCommon Send returns without error.
//
// See ListCleanHistory for more information on using the ListCleanHistory
// API call, and error handling.
//
//    // Example sending a request using the ListCleanHistoryRequest method.
//    req, resp := client.ListCleanHistoryRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *SECCENTER20240508) ListCleanHistoryRequest(input *ListCleanHistoryInput) (req *request.Request, output *ListCleanHistoryOutput) {
	op := &request.Operation{
		Name:       opListCleanHistory,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListCleanHistoryInput{}
	}

	output = &ListCleanHistoryOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// ListCleanHistory API operation for SECCENTER20240508.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for SECCENTER20240508's
// API operation ListCleanHistory for usage and error information.
func (c *SECCENTER20240508) ListCleanHistory(input *ListCleanHistoryInput) (*ListCleanHistoryOutput, error) {
	req, out := c.ListCleanHistoryRequest(input)
	return out, req.Send()
}

// ListCleanHistoryWithContext is the same as ListCleanHistory with the addition of
// the ability to pass a context and additional request options.
//
// See ListCleanHistory for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *SECCENTER20240508) ListCleanHistoryWithContext(ctx volcengine.Context, input *ListCleanHistoryInput, opts ...request.Option) (*ListCleanHistoryOutput, error) {
	req, out := c.ListCleanHistoryRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DataForListCleanHistoryOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	CleanTime *int32 `type:"int32" json:",omitempty"`

	CloudProvider *string `type:"string" json:",omitempty"`

	GroupName *string `type:"string" json:",omitempty"`

	HostId *string `type:"string" json:",omitempty"`

	HostName *string `type:"string" json:",omitempty"`

	OfflineTime *int32 `type:"int32" json:",omitempty"`

	Platform *string `type:"string" json:",omitempty"`

	PrivateIP *string `type:"string" json:",omitempty"`

	PublicIP *string `type:"string" json:",omitempty"`

	Region *string `type:"string" json:",omitempty"`

	Tags []*string `type:"list" json:",omitempty"`

	VpcId *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s DataForListCleanHistoryOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DataForListCleanHistoryOutput) GoString() string {
	return s.String()
}

// SetCleanTime sets the CleanTime field's value.
func (s *DataForListCleanHistoryOutput) SetCleanTime(v int32) *DataForListCleanHistoryOutput {
	s.CleanTime = &v
	return s
}

// SetCloudProvider sets the CloudProvider field's value.
func (s *DataForListCleanHistoryOutput) SetCloudProvider(v string) *DataForListCleanHistoryOutput {
	s.CloudProvider = &v
	return s
}

// SetGroupName sets the GroupName field's value.
func (s *DataForListCleanHistoryOutput) SetGroupName(v string) *DataForListCleanHistoryOutput {
	s.GroupName = &v
	return s
}

// SetHostId sets the HostId field's value.
func (s *DataForListCleanHistoryOutput) SetHostId(v string) *DataForListCleanHistoryOutput {
	s.HostId = &v
	return s
}

// SetHostName sets the HostName field's value.
func (s *DataForListCleanHistoryOutput) SetHostName(v string) *DataForListCleanHistoryOutput {
	s.HostName = &v
	return s
}

// SetOfflineTime sets the OfflineTime field's value.
func (s *DataForListCleanHistoryOutput) SetOfflineTime(v int32) *DataForListCleanHistoryOutput {
	s.OfflineTime = &v
	return s
}

// SetPlatform sets the Platform field's value.
func (s *DataForListCleanHistoryOutput) SetPlatform(v string) *DataForListCleanHistoryOutput {
	s.Platform = &v
	return s
}

// SetPrivateIP sets the PrivateIP field's value.
func (s *DataForListCleanHistoryOutput) SetPrivateIP(v string) *DataForListCleanHistoryOutput {
	s.PrivateIP = &v
	return s
}

// SetPublicIP sets the PublicIP field's value.
func (s *DataForListCleanHistoryOutput) SetPublicIP(v string) *DataForListCleanHistoryOutput {
	s.PublicIP = &v
	return s
}

// SetRegion sets the Region field's value.
func (s *DataForListCleanHistoryOutput) SetRegion(v string) *DataForListCleanHistoryOutput {
	s.Region = &v
	return s
}

// SetTags sets the Tags field's value.
func (s *DataForListCleanHistoryOutput) SetTags(v []*string) *DataForListCleanHistoryOutput {
	s.Tags = v
	return s
}

// SetVpcId sets the VpcId field's value.
func (s *DataForListCleanHistoryOutput) SetVpcId(v string) *DataForListCleanHistoryOutput {
	s.VpcId = &v
	return s
}

type ListCleanHistoryInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	AgentId *string `type:"string" json:",omitempty"`

	CloudProviders []*string `type:"list" json:",omitempty"`

	Hostname *string `type:"string" json:",omitempty"`

	Ip *string `type:"string" json:",omitempty"`

	LeafGroupIDs []*string `type:"list" json:",omitempty"`

	// PageNumber is a required field
	PageNumber *int32 `type:"int32" json:",omitempty" required:"true"`

	// PageSize is a required field
	PageSize *int32 `type:"int32" json:",omitempty" required:"true"`

	Platforms []*string `type:"list" json:",omitempty"`

	Regions []*string `type:"list" json:",omitempty"`

	SortBy *string `type:"string" json:",omitempty"`

	SortOrder *string `type:"string" json:",omitempty"`

	Tags []*string `type:"list" json:",omitempty"`
}

// String returns the string representation
func (s ListCleanHistoryInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListCleanHistoryInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListCleanHistoryInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ListCleanHistoryInput"}
	if s.PageNumber == nil {
		invalidParams.Add(request.NewErrParamRequired("PageNumber"))
	}
	if s.PageSize == nil {
		invalidParams.Add(request.NewErrParamRequired("PageSize"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetAgentId sets the AgentId field's value.
func (s *ListCleanHistoryInput) SetAgentId(v string) *ListCleanHistoryInput {
	s.AgentId = &v
	return s
}

// SetCloudProviders sets the CloudProviders field's value.
func (s *ListCleanHistoryInput) SetCloudProviders(v []*string) *ListCleanHistoryInput {
	s.CloudProviders = v
	return s
}

// SetHostname sets the Hostname field's value.
func (s *ListCleanHistoryInput) SetHostname(v string) *ListCleanHistoryInput {
	s.Hostname = &v
	return s
}

// SetIp sets the Ip field's value.
func (s *ListCleanHistoryInput) SetIp(v string) *ListCleanHistoryInput {
	s.Ip = &v
	return s
}

// SetLeafGroupIDs sets the LeafGroupIDs field's value.
func (s *ListCleanHistoryInput) SetLeafGroupIDs(v []*string) *ListCleanHistoryInput {
	s.LeafGroupIDs = v
	return s
}

// SetPageNumber sets the PageNumber field's value.
func (s *ListCleanHistoryInput) SetPageNumber(v int32) *ListCleanHistoryInput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *ListCleanHistoryInput) SetPageSize(v int32) *ListCleanHistoryInput {
	s.PageSize = &v
	return s
}

// SetPlatforms sets the Platforms field's value.
func (s *ListCleanHistoryInput) SetPlatforms(v []*string) *ListCleanHistoryInput {
	s.Platforms = v
	return s
}

// SetRegions sets the Regions field's value.
func (s *ListCleanHistoryInput) SetRegions(v []*string) *ListCleanHistoryInput {
	s.Regions = v
	return s
}

// SetSortBy sets the SortBy field's value.
func (s *ListCleanHistoryInput) SetSortBy(v string) *ListCleanHistoryInput {
	s.SortBy = &v
	return s
}

// SetSortOrder sets the SortOrder field's value.
func (s *ListCleanHistoryInput) SetSortOrder(v string) *ListCleanHistoryInput {
	s.SortOrder = &v
	return s
}

// SetTags sets the Tags field's value.
func (s *ListCleanHistoryInput) SetTags(v []*string) *ListCleanHistoryInput {
	s.Tags = v
	return s
}

type ListCleanHistoryOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	Data []*DataForListCleanHistoryOutput `type:"list" json:",omitempty"`

	PageNumber *int32 `type:"int32" json:",omitempty"`

	PageSize *int32 `type:"int32" json:",omitempty"`

	TotalCount *int32 `type:"int32" json:",omitempty"`
}

// String returns the string representation
func (s ListCleanHistoryOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListCleanHistoryOutput) GoString() string {
	return s.String()
}

// SetData sets the Data field's value.
func (s *ListCleanHistoryOutput) SetData(v []*DataForListCleanHistoryOutput) *ListCleanHistoryOutput {
	s.Data = v
	return s
}

// SetPageNumber sets the PageNumber field's value.
func (s *ListCleanHistoryOutput) SetPageNumber(v int32) *ListCleanHistoryOutput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *ListCleanHistoryOutput) SetPageSize(v int32) *ListCleanHistoryOutput {
	s.PageSize = &v
	return s
}

// SetTotalCount sets the TotalCount field's value.
func (s *ListCleanHistoryOutput) SetTotalCount(v int32) *ListCleanHistoryOutput {
	s.TotalCount = &v
	return s
}

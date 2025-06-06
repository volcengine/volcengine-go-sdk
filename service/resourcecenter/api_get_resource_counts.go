// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package resourcecenter

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opGetResourceCountsCommon = "GetResourceCounts"

// GetResourceCountsCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the GetResourceCountsCommon operation. The "output" return
// value will be populated with the GetResourceCountsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned GetResourceCountsCommon Request to send the API call to the service.
// the "output" return value is not valid until after GetResourceCountsCommon Send returns without error.
//
// See GetResourceCountsCommon for more information on using the GetResourceCountsCommon
// API call, and error handling.
//
//    // Example sending a request using the GetResourceCountsCommonRequest method.
//    req, resp := client.GetResourceCountsCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *RESOURCECENTER) GetResourceCountsCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opGetResourceCountsCommon,
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

// GetResourceCountsCommon API operation for RESOURCECENTER.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RESOURCECENTER's
// API operation GetResourceCountsCommon for usage and error information.
func (c *RESOURCECENTER) GetResourceCountsCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.GetResourceCountsCommonRequest(input)
	return out, req.Send()
}

// GetResourceCountsCommonWithContext is the same as GetResourceCountsCommon with the addition of
// the ability to pass a context and additional request options.
//
// See GetResourceCountsCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RESOURCECENTER) GetResourceCountsCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.GetResourceCountsCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opGetResourceCounts = "GetResourceCounts"

// GetResourceCountsRequest generates a "volcengine/request.Request" representing the
// client's request for the GetResourceCounts operation. The "output" return
// value will be populated with the GetResourceCountsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned GetResourceCountsCommon Request to send the API call to the service.
// the "output" return value is not valid until after GetResourceCountsCommon Send returns without error.
//
// See GetResourceCounts for more information on using the GetResourceCounts
// API call, and error handling.
//
//    // Example sending a request using the GetResourceCountsRequest method.
//    req, resp := client.GetResourceCountsRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *RESOURCECENTER) GetResourceCountsRequest(input *GetResourceCountsInput) (req *request.Request, output *GetResourceCountsOutput) {
	op := &request.Operation{
		Name:       opGetResourceCounts,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetResourceCountsInput{}
	}

	output = &GetResourceCountsOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// GetResourceCounts API operation for RESOURCECENTER.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RESOURCECENTER's
// API operation GetResourceCounts for usage and error information.
func (c *RESOURCECENTER) GetResourceCounts(input *GetResourceCountsInput) (*GetResourceCountsOutput, error) {
	req, out := c.GetResourceCountsRequest(input)
	return out, req.Send()
}

// GetResourceCountsWithContext is the same as GetResourceCounts with the addition of
// the ability to pass a context and additional request options.
//
// See GetResourceCounts for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RESOURCECENTER) GetResourceCountsWithContext(ctx volcengine.Context, input *GetResourceCountsInput, opts ...request.Option) (*GetResourceCountsOutput, error) {
	req, out := c.GetResourceCountsRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type FilterForGetResourceCountsInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Key *string `type:"string" json:",omitempty"`

	MatchType *string `type:"string" json:",omitempty"`

	Values []*string `type:"list" json:",omitempty"`
}

// String returns the string representation
func (s FilterForGetResourceCountsInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s FilterForGetResourceCountsInput) GoString() string {
	return s.String()
}

// SetKey sets the Key field's value.
func (s *FilterForGetResourceCountsInput) SetKey(v string) *FilterForGetResourceCountsInput {
	s.Key = &v
	return s
}

// SetMatchType sets the MatchType field's value.
func (s *FilterForGetResourceCountsInput) SetMatchType(v string) *FilterForGetResourceCountsInput {
	s.MatchType = &v
	return s
}

// SetValues sets the Values field's value.
func (s *FilterForGetResourceCountsInput) SetValues(v []*string) *FilterForGetResourceCountsInput {
	s.Values = v
	return s
}

type FilterForGetResourceCountsOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Key *string `type:"string" json:",omitempty"`

	MatchType *string `type:"string" json:",omitempty"`

	Values []*string `type:"list" json:",omitempty"`
}

// String returns the string representation
func (s FilterForGetResourceCountsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s FilterForGetResourceCountsOutput) GoString() string {
	return s.String()
}

// SetKey sets the Key field's value.
func (s *FilterForGetResourceCountsOutput) SetKey(v string) *FilterForGetResourceCountsOutput {
	s.Key = &v
	return s
}

// SetMatchType sets the MatchType field's value.
func (s *FilterForGetResourceCountsOutput) SetMatchType(v string) *FilterForGetResourceCountsOutput {
	s.MatchType = &v
	return s
}

// SetValues sets the Values field's value.
func (s *FilterForGetResourceCountsOutput) SetValues(v []*string) *FilterForGetResourceCountsOutput {
	s.Values = v
	return s
}

type GetResourceCountsInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Filter []*FilterForGetResourceCountsInput `type:"list" json:",omitempty"`

	// GroupByKey is a required field
	GroupByKey *string `type:"string" json:",omitempty" required:"true" enum:"EnumOfGroupByKeyForGetResourceCountsInput"`
}

// String returns the string representation
func (s GetResourceCountsInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s GetResourceCountsInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetResourceCountsInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "GetResourceCountsInput"}
	if s.GroupByKey == nil {
		invalidParams.Add(request.NewErrParamRequired("GroupByKey"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetFilter sets the Filter field's value.
func (s *GetResourceCountsInput) SetFilter(v []*FilterForGetResourceCountsInput) *GetResourceCountsInput {
	s.Filter = v
	return s
}

// SetGroupByKey sets the GroupByKey field's value.
func (s *GetResourceCountsInput) SetGroupByKey(v string) *GetResourceCountsInput {
	s.GroupByKey = &v
	return s
}

type GetResourceCountsOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	Filter []*FilterForGetResourceCountsOutput `type:"list" json:",omitempty"`

	GroupByKey *string `type:"string" json:",omitempty"`

	ResourceCounts []*ResourceCountForGetResourceCountsOutput `type:"list" json:",omitempty"`
}

// String returns the string representation
func (s GetResourceCountsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s GetResourceCountsOutput) GoString() string {
	return s.String()
}

// SetFilter sets the Filter field's value.
func (s *GetResourceCountsOutput) SetFilter(v []*FilterForGetResourceCountsOutput) *GetResourceCountsOutput {
	s.Filter = v
	return s
}

// SetGroupByKey sets the GroupByKey field's value.
func (s *GetResourceCountsOutput) SetGroupByKey(v string) *GetResourceCountsOutput {
	s.GroupByKey = &v
	return s
}

// SetResourceCounts sets the ResourceCounts field's value.
func (s *GetResourceCountsOutput) SetResourceCounts(v []*ResourceCountForGetResourceCountsOutput) *GetResourceCountsOutput {
	s.ResourceCounts = v
	return s
}

type ResourceCountForGetResourceCountsOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Count *int32 `type:"int32" json:",omitempty"`

	GroupName *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s ResourceCountForGetResourceCountsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ResourceCountForGetResourceCountsOutput) GoString() string {
	return s.String()
}

// SetCount sets the Count field's value.
func (s *ResourceCountForGetResourceCountsOutput) SetCount(v int32) *ResourceCountForGetResourceCountsOutput {
	s.Count = &v
	return s
}

// SetGroupName sets the GroupName field's value.
func (s *ResourceCountForGetResourceCountsOutput) SetGroupName(v string) *ResourceCountForGetResourceCountsOutput {
	s.GroupName = &v
	return s
}

const (
	// EnumOfGroupByKeyForGetResourceCountsInputResourceType is a EnumOfGroupByKeyForGetResourceCountsInput enum value
	EnumOfGroupByKeyForGetResourceCountsInputResourceType = "ResourceType"

	// EnumOfGroupByKeyForGetResourceCountsInputRegion is a EnumOfGroupByKeyForGetResourceCountsInput enum value
	EnumOfGroupByKeyForGetResourceCountsInputRegion = "Region"
)

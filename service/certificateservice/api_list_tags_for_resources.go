// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package certificateservice

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opListTagsForResourcesCommon = "ListTagsForResources"

// ListTagsForResourcesCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ListTagsForResourcesCommon operation. The "output" return
// value will be populated with the ListTagsForResourcesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListTagsForResourcesCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListTagsForResourcesCommon Send returns without error.
//
// See ListTagsForResourcesCommon for more information on using the ListTagsForResourcesCommon
// API call, and error handling.
//
//    // Example sending a request using the ListTagsForResourcesCommonRequest method.
//    req, resp := client.ListTagsForResourcesCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CERTIFICATESERVICE) ListTagsForResourcesCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opListTagsForResourcesCommon,
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

// ListTagsForResourcesCommon API operation for CERTIFICATE_SERVICE.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CERTIFICATE_SERVICE's
// API operation ListTagsForResourcesCommon for usage and error information.
func (c *CERTIFICATESERVICE) ListTagsForResourcesCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ListTagsForResourcesCommonRequest(input)
	return out, req.Send()
}

// ListTagsForResourcesCommonWithContext is the same as ListTagsForResourcesCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ListTagsForResourcesCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CERTIFICATESERVICE) ListTagsForResourcesCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ListTagsForResourcesCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opListTagsForResources = "ListTagsForResources"

// ListTagsForResourcesRequest generates a "volcengine/request.Request" representing the
// client's request for the ListTagsForResources operation. The "output" return
// value will be populated with the ListTagsForResourcesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListTagsForResourcesCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListTagsForResourcesCommon Send returns without error.
//
// See ListTagsForResources for more information on using the ListTagsForResources
// API call, and error handling.
//
//    // Example sending a request using the ListTagsForResourcesRequest method.
//    req, resp := client.ListTagsForResourcesRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CERTIFICATESERVICE) ListTagsForResourcesRequest(input *ListTagsForResourcesInput) (req *request.Request, output *ListTagsForResourcesOutput) {
	op := &request.Operation{
		Name:       opListTagsForResources,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListTagsForResourcesInput{}
	}

	output = &ListTagsForResourcesOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// ListTagsForResources API operation for CERTIFICATE_SERVICE.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CERTIFICATE_SERVICE's
// API operation ListTagsForResources for usage and error information.
func (c *CERTIFICATESERVICE) ListTagsForResources(input *ListTagsForResourcesInput) (*ListTagsForResourcesOutput, error) {
	req, out := c.ListTagsForResourcesRequest(input)
	return out, req.Send()
}

// ListTagsForResourcesWithContext is the same as ListTagsForResources with the addition of
// the ability to pass a context and additional request options.
//
// See ListTagsForResources for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CERTIFICATESERVICE) ListTagsForResourcesWithContext(ctx volcengine.Context, input *ListTagsForResourcesInput, opts ...request.Option) (*ListTagsForResourcesOutput, error) {
	req, out := c.ListTagsForResourcesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ListTagsForResourcesInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	PageNumber *int32 `type:"int32" json:",omitempty"`

	PageSize *int32 `type:"int32" json:",omitempty"`

	ResourceIds []*string `type:"list" json:",omitempty"`

	// ResourceType is a required field
	ResourceType *string `type:"string" json:",omitempty" required:"true"`

	TagFilters []*TagFilterForListTagsForResourcesInput `type:"list" json:",omitempty"`
}

// String returns the string representation
func (s ListTagsForResourcesInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListTagsForResourcesInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListTagsForResourcesInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ListTagsForResourcesInput"}
	if s.ResourceType == nil {
		invalidParams.Add(request.NewErrParamRequired("ResourceType"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetPageNumber sets the PageNumber field's value.
func (s *ListTagsForResourcesInput) SetPageNumber(v int32) *ListTagsForResourcesInput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *ListTagsForResourcesInput) SetPageSize(v int32) *ListTagsForResourcesInput {
	s.PageSize = &v
	return s
}

// SetResourceIds sets the ResourceIds field's value.
func (s *ListTagsForResourcesInput) SetResourceIds(v []*string) *ListTagsForResourcesInput {
	s.ResourceIds = v
	return s
}

// SetResourceType sets the ResourceType field's value.
func (s *ListTagsForResourcesInput) SetResourceType(v string) *ListTagsForResourcesInput {
	s.ResourceType = &v
	return s
}

// SetTagFilters sets the TagFilters field's value.
func (s *ListTagsForResourcesInput) SetTagFilters(v []*TagFilterForListTagsForResourcesInput) *ListTagsForResourcesInput {
	s.TagFilters = v
	return s
}

type ListTagsForResourcesOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	PageNumber *int32 `type:"int32" json:",omitempty"`

	PageSize *int32 `type:"int32" json:",omitempty"`

	ResourceTags []*ResourceTagForListTagsForResourcesOutput `type:"list" json:",omitempty"`

	TotalCount *int32 `type:"int32" json:",omitempty"`
}

// String returns the string representation
func (s ListTagsForResourcesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListTagsForResourcesOutput) GoString() string {
	return s.String()
}

// SetPageNumber sets the PageNumber field's value.
func (s *ListTagsForResourcesOutput) SetPageNumber(v int32) *ListTagsForResourcesOutput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *ListTagsForResourcesOutput) SetPageSize(v int32) *ListTagsForResourcesOutput {
	s.PageSize = &v
	return s
}

// SetResourceTags sets the ResourceTags field's value.
func (s *ListTagsForResourcesOutput) SetResourceTags(v []*ResourceTagForListTagsForResourcesOutput) *ListTagsForResourcesOutput {
	s.ResourceTags = v
	return s
}

// SetTotalCount sets the TotalCount field's value.
func (s *ListTagsForResourcesOutput) SetTotalCount(v int32) *ListTagsForResourcesOutput {
	s.TotalCount = &v
	return s
}

type ResourceTagForListTagsForResourcesOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	ResourceId *string `type:"string" json:",omitempty"`

	ResourceType *string `type:"string" json:",omitempty"`

	TagKey *string `type:"string" json:",omitempty"`

	TagValue *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s ResourceTagForListTagsForResourcesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ResourceTagForListTagsForResourcesOutput) GoString() string {
	return s.String()
}

// SetResourceId sets the ResourceId field's value.
func (s *ResourceTagForListTagsForResourcesOutput) SetResourceId(v string) *ResourceTagForListTagsForResourcesOutput {
	s.ResourceId = &v
	return s
}

// SetResourceType sets the ResourceType field's value.
func (s *ResourceTagForListTagsForResourcesOutput) SetResourceType(v string) *ResourceTagForListTagsForResourcesOutput {
	s.ResourceType = &v
	return s
}

// SetTagKey sets the TagKey field's value.
func (s *ResourceTagForListTagsForResourcesOutput) SetTagKey(v string) *ResourceTagForListTagsForResourcesOutput {
	s.TagKey = &v
	return s
}

// SetTagValue sets the TagValue field's value.
func (s *ResourceTagForListTagsForResourcesOutput) SetTagValue(v string) *ResourceTagForListTagsForResourcesOutput {
	s.TagValue = &v
	return s
}

type TagFilterForListTagsForResourcesInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Key *string `type:"string" json:",omitempty"`

	Values []*string `type:"list" json:",omitempty"`
}

// String returns the string representation
func (s TagFilterForListTagsForResourcesInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s TagFilterForListTagsForResourcesInput) GoString() string {
	return s.String()
}

// SetKey sets the Key field's value.
func (s *TagFilterForListTagsForResourcesInput) SetKey(v string) *TagFilterForListTagsForResourcesInput {
	s.Key = &v
	return s
}

// SetValues sets the Values field's value.
func (s *TagFilterForListTagsForResourcesInput) SetValues(v []*string) *TagFilterForListTagsForResourcesInput {
	s.Values = v
	return s
}

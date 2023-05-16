// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package vpc

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
//	// Example sending a request using the ListTagsForResourcesCommonRequest method.
//	req, resp := client.ListTagsForResourcesCommonRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *VPC) ListTagsForResourcesCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opListTagsForResourcesCommon,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &map[string]interface{}{}
	}

	output = &map[string]interface{}{}
	req = c.newRequest(op, input, output)

	return
}

// ListTagsForResourcesCommon API operation for VPC.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VPC's
// API operation ListTagsForResourcesCommon for usage and error information.
func (c *VPC) ListTagsForResourcesCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
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
func (c *VPC) ListTagsForResourcesCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
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
//	// Example sending a request using the ListTagsForResourcesRequest method.
//	req, resp := client.ListTagsForResourcesRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *VPC) ListTagsForResourcesRequest(input *ListTagsForResourcesInput) (req *request.Request, output *ListTagsForResourcesOutput) {
	op := &request.Operation{
		Name:       opListTagsForResources,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListTagsForResourcesInput{}
	}

	output = &ListTagsForResourcesOutput{}
	req = c.newRequest(op, input, output)

	return
}

// ListTagsForResources API operation for VPC.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VPC's
// API operation ListTagsForResources for usage and error information.
func (c *VPC) ListTagsForResources(input *ListTagsForResourcesInput) (*ListTagsForResourcesOutput, error) {
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
func (c *VPC) ListTagsForResourcesWithContext(ctx volcengine.Context, input *ListTagsForResourcesInput, opts ...request.Option) (*ListTagsForResourcesOutput, error) {
	req, out := c.ListTagsForResourcesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ListTagsForResourcesInput struct {
	_ struct{} `type:"structure"`

	MaxResults *int64 `max:"100" type:"integer"`

	NextToken *string `type:"string"`

	ResourceIds []*string `type:"list"`

	// ResourceType is a required field
	ResourceType *string `type:"string" required:"true" enum:"ResourceTypeForListTagsForResourcesInput"`

	TagFilters []*TagFilterForListTagsForResourcesInput `type:"list"`

	TagType *string `type:"string"`
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
	if s.MaxResults != nil && *s.MaxResults > 100 {
		invalidParams.Add(request.NewErrParamMaxValue("MaxResults", 100))
	}
	if s.ResourceType == nil {
		invalidParams.Add(request.NewErrParamRequired("ResourceType"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetMaxResults sets the MaxResults field's value.
func (s *ListTagsForResourcesInput) SetMaxResults(v int64) *ListTagsForResourcesInput {
	s.MaxResults = &v
	return s
}

// SetNextToken sets the NextToken field's value.
func (s *ListTagsForResourcesInput) SetNextToken(v string) *ListTagsForResourcesInput {
	s.NextToken = &v
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

// SetTagType sets the TagType field's value.
func (s *ListTagsForResourcesInput) SetTagType(v string) *ListTagsForResourcesInput {
	s.TagType = &v
	return s
}

type ListTagsForResourcesOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	NextToken *string `type:"string"`

	RequestId *string `type:"string"`

	ResourceTags []*ResourceTagForListTagsForResourcesOutput `type:"list"`
}

// String returns the string representation
func (s ListTagsForResourcesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListTagsForResourcesOutput) GoString() string {
	return s.String()
}

// SetNextToken sets the NextToken field's value.
func (s *ListTagsForResourcesOutput) SetNextToken(v string) *ListTagsForResourcesOutput {
	s.NextToken = &v
	return s
}

// SetRequestId sets the RequestId field's value.
func (s *ListTagsForResourcesOutput) SetRequestId(v string) *ListTagsForResourcesOutput {
	s.RequestId = &v
	return s
}

// SetResourceTags sets the ResourceTags field's value.
func (s *ListTagsForResourcesOutput) SetResourceTags(v []*ResourceTagForListTagsForResourcesOutput) *ListTagsForResourcesOutput {
	s.ResourceTags = v
	return s
}

type ResourceTagForListTagsForResourcesOutput struct {
	_ struct{} `type:"structure"`

	ResourceId *string `type:"string"`

	ResourceType *string `type:"string"`

	TagKey *string `type:"string"`

	TagValue *string `type:"string"`
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
	_ struct{} `type:"structure"`

	Key *string `type:"string"`

	Values []*string `type:"list"`
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

const (
	// ResourceTypeForListTagsForResourcesInputVpc is a ResourceTypeForListTagsForResourcesInput enum value
	ResourceTypeForListTagsForResourcesInputVpc = "vpc"

	// ResourceTypeForListTagsForResourcesInputEni is a ResourceTypeForListTagsForResourcesInput enum value
	ResourceTypeForListTagsForResourcesInputEni = "eni"

	// ResourceTypeForListTagsForResourcesInputSecuritygroup is a ResourceTypeForListTagsForResourcesInput enum value
	ResourceTypeForListTagsForResourcesInputSecuritygroup = "securitygroup"

	// ResourceTypeForListTagsForResourcesInputEip is a ResourceTypeForListTagsForResourcesInput enum value
	ResourceTypeForListTagsForResourcesInputEip = "eip"

	// ResourceTypeForListTagsForResourcesInputBandwidthpackage is a ResourceTypeForListTagsForResourcesInput enum value
	ResourceTypeForListTagsForResourcesInputBandwidthpackage = "bandwidthpackage"

	// ResourceTypeForListTagsForResourcesInputVpngateway is a ResourceTypeForListTagsForResourcesInput enum value
	ResourceTypeForListTagsForResourcesInputVpngateway = "vpngateway"

	// ResourceTypeForListTagsForResourcesInputNgw is a ResourceTypeForListTagsForResourcesInput enum value
	ResourceTypeForListTagsForResourcesInputNgw = "ngw"

	// ResourceTypeForListTagsForResourcesInputDirectconnectconnection is a ResourceTypeForListTagsForResourcesInput enum value
	ResourceTypeForListTagsForResourcesInputDirectconnectconnection = "directconnectconnection"

	// ResourceTypeForListTagsForResourcesInputDirectconnectgateway is a ResourceTypeForListTagsForResourcesInput enum value
	ResourceTypeForListTagsForResourcesInputDirectconnectgateway = "directconnectgateway"

	// ResourceTypeForListTagsForResourcesInputDirectconnectvirtualinterface is a ResourceTypeForListTagsForResourcesInput enum value
	ResourceTypeForListTagsForResourcesInputDirectconnectvirtualinterface = "directconnectvirtualinterface"
)

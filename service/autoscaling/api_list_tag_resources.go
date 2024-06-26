// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package autoscaling

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opListTagResourcesCommon = "ListTagResources"

// ListTagResourcesCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ListTagResourcesCommon operation. The "output" return
// value will be populated with the ListTagResourcesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListTagResourcesCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListTagResourcesCommon Send returns without error.
//
// See ListTagResourcesCommon for more information on using the ListTagResourcesCommon
// API call, and error handling.
//
//    // Example sending a request using the ListTagResourcesCommonRequest method.
//    req, resp := client.ListTagResourcesCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *AUTOSCALING) ListTagResourcesCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opListTagResourcesCommon,
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

// ListTagResourcesCommon API operation for AUTO_SCALING.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for AUTO_SCALING's
// API operation ListTagResourcesCommon for usage and error information.
func (c *AUTOSCALING) ListTagResourcesCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ListTagResourcesCommonRequest(input)
	return out, req.Send()
}

// ListTagResourcesCommonWithContext is the same as ListTagResourcesCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ListTagResourcesCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *AUTOSCALING) ListTagResourcesCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ListTagResourcesCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opListTagResources = "ListTagResources"

// ListTagResourcesRequest generates a "volcengine/request.Request" representing the
// client's request for the ListTagResources operation. The "output" return
// value will be populated with the ListTagResourcesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListTagResourcesCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListTagResourcesCommon Send returns without error.
//
// See ListTagResources for more information on using the ListTagResources
// API call, and error handling.
//
//    // Example sending a request using the ListTagResourcesRequest method.
//    req, resp := client.ListTagResourcesRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *AUTOSCALING) ListTagResourcesRequest(input *ListTagResourcesInput) (req *request.Request, output *ListTagResourcesOutput) {
	op := &request.Operation{
		Name:       opListTagResources,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListTagResourcesInput{}
	}

	output = &ListTagResourcesOutput{}
	req = c.newRequest(op, input, output)

	return
}

// ListTagResources API operation for AUTO_SCALING.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for AUTO_SCALING's
// API operation ListTagResources for usage and error information.
func (c *AUTOSCALING) ListTagResources(input *ListTagResourcesInput) (*ListTagResourcesOutput, error) {
	req, out := c.ListTagResourcesRequest(input)
	return out, req.Send()
}

// ListTagResourcesWithContext is the same as ListTagResources with the addition of
// the ability to pass a context and additional request options.
//
// See ListTagResources for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *AUTOSCALING) ListTagResourcesWithContext(ctx volcengine.Context, input *ListTagResourcesInput, opts ...request.Option) (*ListTagResourcesOutput, error) {
	req, out := c.ListTagResourcesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ListTagResourcesInput struct {
	_ struct{} `type:"structure"`

	MaxResults *int32 `min:"1" max:"100" type:"int32"`

	NextToken *string `type:"string"`

	ResourceIds []*string `type:"list"`

	// ResourceType is a required field
	ResourceType *string `type:"string" required:"true"`

	TagFilters []*TagFilterForListTagResourcesInput `type:"list"`
}

// String returns the string representation
func (s ListTagResourcesInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListTagResourcesInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListTagResourcesInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ListTagResourcesInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(request.NewErrParamMinValue("MaxResults", 1))
	}
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
func (s *ListTagResourcesInput) SetMaxResults(v int32) *ListTagResourcesInput {
	s.MaxResults = &v
	return s
}

// SetNextToken sets the NextToken field's value.
func (s *ListTagResourcesInput) SetNextToken(v string) *ListTagResourcesInput {
	s.NextToken = &v
	return s
}

// SetResourceIds sets the ResourceIds field's value.
func (s *ListTagResourcesInput) SetResourceIds(v []*string) *ListTagResourcesInput {
	s.ResourceIds = v
	return s
}

// SetResourceType sets the ResourceType field's value.
func (s *ListTagResourcesInput) SetResourceType(v string) *ListTagResourcesInput {
	s.ResourceType = &v
	return s
}

// SetTagFilters sets the TagFilters field's value.
func (s *ListTagResourcesInput) SetTagFilters(v []*TagFilterForListTagResourcesInput) *ListTagResourcesInput {
	s.TagFilters = v
	return s
}

type ListTagResourcesOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	NextToken *string `type:"string"`

	TagResources []*TagResourceForListTagResourcesOutput `type:"list"`
}

// String returns the string representation
func (s ListTagResourcesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListTagResourcesOutput) GoString() string {
	return s.String()
}

// SetNextToken sets the NextToken field's value.
func (s *ListTagResourcesOutput) SetNextToken(v string) *ListTagResourcesOutput {
	s.NextToken = &v
	return s
}

// SetTagResources sets the TagResources field's value.
func (s *ListTagResourcesOutput) SetTagResources(v []*TagResourceForListTagResourcesOutput) *ListTagResourcesOutput {
	s.TagResources = v
	return s
}

type TagFilterForListTagResourcesInput struct {
	_ struct{} `type:"structure"`

	Key *string `type:"string"`

	Value *string `type:"string"`
}

// String returns the string representation
func (s TagFilterForListTagResourcesInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s TagFilterForListTagResourcesInput) GoString() string {
	return s.String()
}

// SetKey sets the Key field's value.
func (s *TagFilterForListTagResourcesInput) SetKey(v string) *TagFilterForListTagResourcesInput {
	s.Key = &v
	return s
}

// SetValue sets the Value field's value.
func (s *TagFilterForListTagResourcesInput) SetValue(v string) *TagFilterForListTagResourcesInput {
	s.Value = &v
	return s
}

type TagResourceForListTagResourcesOutput struct {
	_ struct{} `type:"structure"`

	ResourceId *string `type:"string"`

	ResourceType *string `type:"string"`

	TagKey *string `type:"string"`

	TagValue *string `type:"string"`
}

// String returns the string representation
func (s TagResourceForListTagResourcesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s TagResourceForListTagResourcesOutput) GoString() string {
	return s.String()
}

// SetResourceId sets the ResourceId field's value.
func (s *TagResourceForListTagResourcesOutput) SetResourceId(v string) *TagResourceForListTagResourcesOutput {
	s.ResourceId = &v
	return s
}

// SetResourceType sets the ResourceType field's value.
func (s *TagResourceForListTagResourcesOutput) SetResourceType(v string) *TagResourceForListTagResourcesOutput {
	s.ResourceType = &v
	return s
}

// SetTagKey sets the TagKey field's value.
func (s *TagResourceForListTagResourcesOutput) SetTagKey(v string) *TagResourceForListTagResourcesOutput {
	s.TagKey = &v
	return s
}

// SetTagValue sets the TagValue field's value.
func (s *TagResourceForListTagResourcesOutput) SetTagValue(v string) *TagResourceForListTagResourcesOutput {
	s.TagValue = &v
	return s
}

// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package resourcecenter

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opListResourceTypesCommon = "ListResourceTypes"

// ListResourceTypesCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ListResourceTypesCommon operation. The "output" return
// value will be populated with the ListResourceTypesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListResourceTypesCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListResourceTypesCommon Send returns without error.
//
// See ListResourceTypesCommon for more information on using the ListResourceTypesCommon
// API call, and error handling.
//
//    // Example sending a request using the ListResourceTypesCommonRequest method.
//    req, resp := client.ListResourceTypesCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *RESOURCECENTER) ListResourceTypesCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opListResourceTypesCommon,
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

// ListResourceTypesCommon API operation for RESOURCECENTER.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RESOURCECENTER's
// API operation ListResourceTypesCommon for usage and error information.
func (c *RESOURCECENTER) ListResourceTypesCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ListResourceTypesCommonRequest(input)
	return out, req.Send()
}

// ListResourceTypesCommonWithContext is the same as ListResourceTypesCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ListResourceTypesCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RESOURCECENTER) ListResourceTypesCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ListResourceTypesCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opListResourceTypes = "ListResourceTypes"

// ListResourceTypesRequest generates a "volcengine/request.Request" representing the
// client's request for the ListResourceTypes operation. The "output" return
// value will be populated with the ListResourceTypesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListResourceTypesCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListResourceTypesCommon Send returns without error.
//
// See ListResourceTypes for more information on using the ListResourceTypes
// API call, and error handling.
//
//    // Example sending a request using the ListResourceTypesRequest method.
//    req, resp := client.ListResourceTypesRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *RESOURCECENTER) ListResourceTypesRequest(input *ListResourceTypesInput) (req *request.Request, output *ListResourceTypesOutput) {
	op := &request.Operation{
		Name:       opListResourceTypes,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListResourceTypesInput{}
	}

	output = &ListResourceTypesOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// ListResourceTypes API operation for RESOURCECENTER.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RESOURCECENTER's
// API operation ListResourceTypes for usage and error information.
func (c *RESOURCECENTER) ListResourceTypes(input *ListResourceTypesInput) (*ListResourceTypesOutput, error) {
	req, out := c.ListResourceTypesRequest(input)
	return out, req.Send()
}

// ListResourceTypesWithContext is the same as ListResourceTypes with the addition of
// the ability to pass a context and additional request options.
//
// See ListResourceTypes for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RESOURCECENTER) ListResourceTypesWithContext(ctx volcengine.Context, input *ListResourceTypesInput, opts ...request.Option) (*ListResourceTypesOutput, error) {
	req, out := c.ListResourceTypesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ListResourceTypesInput struct {
	_ struct{} `type:"structure" json:",omitempty"`
}

// String returns the string representation
func (s ListResourceTypesInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListResourceTypesInput) GoString() string {
	return s.String()
}

type ListResourceTypesOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	ResourceTypes []*ResourceTypeForListResourceTypesOutput `type:"list" json:",omitempty"`
}

// String returns the string representation
func (s ListResourceTypesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListResourceTypesOutput) GoString() string {
	return s.String()
}

// SetResourceTypes sets the ResourceTypes field's value.
func (s *ListResourceTypesOutput) SetResourceTypes(v []*ResourceTypeForListResourceTypesOutput) *ListResourceTypesOutput {
	s.ResourceTypes = v
	return s
}

type ResourceTypeForListResourceTypesOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	ResourceType *string `type:"string" json:",omitempty"`

	ResourceTypeName *string `type:"string" json:",omitempty"`

	Service *string `type:"string" json:",omitempty"`

	ServiceName *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s ResourceTypeForListResourceTypesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ResourceTypeForListResourceTypesOutput) GoString() string {
	return s.String()
}

// SetResourceType sets the ResourceType field's value.
func (s *ResourceTypeForListResourceTypesOutput) SetResourceType(v string) *ResourceTypeForListResourceTypesOutput {
	s.ResourceType = &v
	return s
}

// SetResourceTypeName sets the ResourceTypeName field's value.
func (s *ResourceTypeForListResourceTypesOutput) SetResourceTypeName(v string) *ResourceTypeForListResourceTypesOutput {
	s.ResourceTypeName = &v
	return s
}

// SetService sets the Service field's value.
func (s *ResourceTypeForListResourceTypesOutput) SetService(v string) *ResourceTypeForListResourceTypesOutput {
	s.Service = &v
	return s
}

// SetServiceName sets the ServiceName field's value.
func (s *ResourceTypeForListResourceTypesOutput) SetServiceName(v string) *ResourceTypeForListResourceTypesOutput {
	s.ServiceName = &v
	return s
}

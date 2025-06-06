// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package cdn

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opListResourceTagsCommon = "ListResourceTags"

// ListResourceTagsCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ListResourceTagsCommon operation. The "output" return
// value will be populated with the ListResourceTagsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListResourceTagsCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListResourceTagsCommon Send returns without error.
//
// See ListResourceTagsCommon for more information on using the ListResourceTagsCommon
// API call, and error handling.
//
//    // Example sending a request using the ListResourceTagsCommonRequest method.
//    req, resp := client.ListResourceTagsCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CDN) ListResourceTagsCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opListResourceTagsCommon,
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

// ListResourceTagsCommon API operation for CDN.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CDN's
// API operation ListResourceTagsCommon for usage and error information.
func (c *CDN) ListResourceTagsCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ListResourceTagsCommonRequest(input)
	return out, req.Send()
}

// ListResourceTagsCommonWithContext is the same as ListResourceTagsCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ListResourceTagsCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CDN) ListResourceTagsCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ListResourceTagsCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opListResourceTags = "ListResourceTags"

// ListResourceTagsRequest generates a "volcengine/request.Request" representing the
// client's request for the ListResourceTags operation. The "output" return
// value will be populated with the ListResourceTagsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListResourceTagsCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListResourceTagsCommon Send returns without error.
//
// See ListResourceTags for more information on using the ListResourceTags
// API call, and error handling.
//
//    // Example sending a request using the ListResourceTagsRequest method.
//    req, resp := client.ListResourceTagsRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CDN) ListResourceTagsRequest(input *ListResourceTagsInput) (req *request.Request, output *ListResourceTagsOutput) {
	op := &request.Operation{
		Name:       opListResourceTags,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListResourceTagsInput{}
	}

	output = &ListResourceTagsOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// ListResourceTags API operation for CDN.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CDN's
// API operation ListResourceTags for usage and error information.
func (c *CDN) ListResourceTags(input *ListResourceTagsInput) (*ListResourceTagsOutput, error) {
	req, out := c.ListResourceTagsRequest(input)
	return out, req.Send()
}

// ListResourceTagsWithContext is the same as ListResourceTags with the addition of
// the ability to pass a context and additional request options.
//
// See ListResourceTags for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CDN) ListResourceTagsWithContext(ctx volcengine.Context, input *ListResourceTagsInput, opts ...request.Option) (*ListResourceTagsOutput, error) {
	req, out := c.ListResourceTagsRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ListResourceTagsInput struct {
	_ struct{} `type:"structure" json:",omitempty"`
}

// String returns the string representation
func (s ListResourceTagsInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListResourceTagsInput) GoString() string {
	return s.String()
}

type ListResourceTagsOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	ResourceTags []*ResourceTagForListResourceTagsOutput `type:"list" json:",omitempty"`
}

// String returns the string representation
func (s ListResourceTagsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListResourceTagsOutput) GoString() string {
	return s.String()
}

// SetResourceTags sets the ResourceTags field's value.
func (s *ListResourceTagsOutput) SetResourceTags(v []*ResourceTagForListResourceTagsOutput) *ListResourceTagsOutput {
	s.ResourceTags = v
	return s
}

type ResourceTagForListResourceTagsOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Key *string `type:"string" json:",omitempty"`

	Value *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s ResourceTagForListResourceTagsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ResourceTagForListResourceTagsOutput) GoString() string {
	return s.String()
}

// SetKey sets the Key field's value.
func (s *ResourceTagForListResourceTagsOutput) SetKey(v string) *ResourceTagForListResourceTagsOutput {
	s.Key = &v
	return s
}

// SetValue sets the Value field's value.
func (s *ResourceTagForListResourceTagsOutput) SetValue(v string) *ResourceTagForListResourceTagsOutput {
	s.Value = &v
	return s
}

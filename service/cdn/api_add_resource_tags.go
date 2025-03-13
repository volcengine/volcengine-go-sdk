// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package cdn

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opAddResourceTagsCommon = "AddResourceTags"

// AddResourceTagsCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the AddResourceTagsCommon operation. The "output" return
// value will be populated with the AddResourceTagsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned AddResourceTagsCommon Request to send the API call to the service.
// the "output" return value is not valid until after AddResourceTagsCommon Send returns without error.
//
// See AddResourceTagsCommon for more information on using the AddResourceTagsCommon
// API call, and error handling.
//
//    // Example sending a request using the AddResourceTagsCommonRequest method.
//    req, resp := client.AddResourceTagsCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CDN) AddResourceTagsCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opAddResourceTagsCommon,
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

// AddResourceTagsCommon API operation for CDN.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CDN's
// API operation AddResourceTagsCommon for usage and error information.
func (c *CDN) AddResourceTagsCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.AddResourceTagsCommonRequest(input)
	return out, req.Send()
}

// AddResourceTagsCommonWithContext is the same as AddResourceTagsCommon with the addition of
// the ability to pass a context and additional request options.
//
// See AddResourceTagsCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CDN) AddResourceTagsCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.AddResourceTagsCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opAddResourceTags = "AddResourceTags"

// AddResourceTagsRequest generates a "volcengine/request.Request" representing the
// client's request for the AddResourceTags operation. The "output" return
// value will be populated with the AddResourceTagsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned AddResourceTagsCommon Request to send the API call to the service.
// the "output" return value is not valid until after AddResourceTagsCommon Send returns without error.
//
// See AddResourceTags for more information on using the AddResourceTags
// API call, and error handling.
//
//    // Example sending a request using the AddResourceTagsRequest method.
//    req, resp := client.AddResourceTagsRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CDN) AddResourceTagsRequest(input *AddResourceTagsInput) (req *request.Request, output *AddResourceTagsOutput) {
	op := &request.Operation{
		Name:       opAddResourceTags,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &AddResourceTagsInput{}
	}

	output = &AddResourceTagsOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// AddResourceTags API operation for CDN.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CDN's
// API operation AddResourceTags for usage and error information.
func (c *CDN) AddResourceTags(input *AddResourceTagsInput) (*AddResourceTagsOutput, error) {
	req, out := c.AddResourceTagsRequest(input)
	return out, req.Send()
}

// AddResourceTagsWithContext is the same as AddResourceTags with the addition of
// the ability to pass a context and additional request options.
//
// See AddResourceTags for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CDN) AddResourceTagsWithContext(ctx volcengine.Context, input *AddResourceTagsInput, opts ...request.Option) (*AddResourceTagsOutput, error) {
	req, out := c.AddResourceTagsRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type AddResourceTagsInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	ResourceTags []*ResourceTagForAddResourceTagsInput `type:"list" json:",omitempty"`

	Resources []*string `type:"list" json:",omitempty"`
}

// String returns the string representation
func (s AddResourceTagsInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s AddResourceTagsInput) GoString() string {
	return s.String()
}

// SetResourceTags sets the ResourceTags field's value.
func (s *AddResourceTagsInput) SetResourceTags(v []*ResourceTagForAddResourceTagsInput) *AddResourceTagsInput {
	s.ResourceTags = v
	return s
}

// SetResources sets the Resources field's value.
func (s *AddResourceTagsInput) SetResources(v []*string) *AddResourceTagsInput {
	s.Resources = v
	return s
}

type AddResourceTagsOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s AddResourceTagsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s AddResourceTagsOutput) GoString() string {
	return s.String()
}

type ResourceTagForAddResourceTagsInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Key *string `type:"string" json:",omitempty"`

	Value *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s ResourceTagForAddResourceTagsInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ResourceTagForAddResourceTagsInput) GoString() string {
	return s.String()
}

// SetKey sets the Key field's value.
func (s *ResourceTagForAddResourceTagsInput) SetKey(v string) *ResourceTagForAddResourceTagsInput {
	s.Key = &v
	return s
}

// SetValue sets the Value field's value.
func (s *ResourceTagForAddResourceTagsInput) SetValue(v string) *ResourceTagForAddResourceTagsInput {
	s.Value = &v
	return s
}

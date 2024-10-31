// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package redis

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opAddTagsToResourceCommon = "AddTagsToResource"

// AddTagsToResourceCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the AddTagsToResourceCommon operation. The "output" return
// value will be populated with the AddTagsToResourceCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned AddTagsToResourceCommon Request to send the API call to the service.
// the "output" return value is not valid until after AddTagsToResourceCommon Send returns without error.
//
// See AddTagsToResourceCommon for more information on using the AddTagsToResourceCommon
// API call, and error handling.
//
//    // Example sending a request using the AddTagsToResourceCommonRequest method.
//    req, resp := client.AddTagsToResourceCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *REDIS) AddTagsToResourceCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opAddTagsToResourceCommon,
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

// AddTagsToResourceCommon API operation for REDIS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for REDIS's
// API operation AddTagsToResourceCommon for usage and error information.
func (c *REDIS) AddTagsToResourceCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.AddTagsToResourceCommonRequest(input)
	return out, req.Send()
}

// AddTagsToResourceCommonWithContext is the same as AddTagsToResourceCommon with the addition of
// the ability to pass a context and additional request options.
//
// See AddTagsToResourceCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *REDIS) AddTagsToResourceCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.AddTagsToResourceCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opAddTagsToResource = "AddTagsToResource"

// AddTagsToResourceRequest generates a "volcengine/request.Request" representing the
// client's request for the AddTagsToResource operation. The "output" return
// value will be populated with the AddTagsToResourceCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned AddTagsToResourceCommon Request to send the API call to the service.
// the "output" return value is not valid until after AddTagsToResourceCommon Send returns without error.
//
// See AddTagsToResource for more information on using the AddTagsToResource
// API call, and error handling.
//
//    // Example sending a request using the AddTagsToResourceRequest method.
//    req, resp := client.AddTagsToResourceRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *REDIS) AddTagsToResourceRequest(input *AddTagsToResourceInput) (req *request.Request, output *AddTagsToResourceOutput) {
	op := &request.Operation{
		Name:       opAddTagsToResource,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &AddTagsToResourceInput{}
	}

	output = &AddTagsToResourceOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// AddTagsToResource API operation for REDIS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for REDIS's
// API operation AddTagsToResource for usage and error information.
func (c *REDIS) AddTagsToResource(input *AddTagsToResourceInput) (*AddTagsToResourceOutput, error) {
	req, out := c.AddTagsToResourceRequest(input)
	return out, req.Send()
}

// AddTagsToResourceWithContext is the same as AddTagsToResource with the addition of
// the ability to pass a context and additional request options.
//
// See AddTagsToResource for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *REDIS) AddTagsToResourceWithContext(ctx volcengine.Context, input *AddTagsToResourceInput, opts ...request.Option) (*AddTagsToResourceOutput, error) {
	req, out := c.AddTagsToResourceRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type AddTagsToResourceInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	ClientToken *string `type:"string" json:",omitempty"`

	InstanceIds []*string `type:"list" json:",omitempty"`

	Tags []*TagForAddTagsToResourceInput `type:"list" json:",omitempty"`
}

// String returns the string representation
func (s AddTagsToResourceInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s AddTagsToResourceInput) GoString() string {
	return s.String()
}

// SetClientToken sets the ClientToken field's value.
func (s *AddTagsToResourceInput) SetClientToken(v string) *AddTagsToResourceInput {
	s.ClientToken = &v
	return s
}

// SetInstanceIds sets the InstanceIds field's value.
func (s *AddTagsToResourceInput) SetInstanceIds(v []*string) *AddTagsToResourceInput {
	s.InstanceIds = v
	return s
}

// SetTags sets the Tags field's value.
func (s *AddTagsToResourceInput) SetTags(v []*TagForAddTagsToResourceInput) *AddTagsToResourceInput {
	s.Tags = v
	return s
}

type AddTagsToResourceOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s AddTagsToResourceOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s AddTagsToResourceOutput) GoString() string {
	return s.String()
}

type TagForAddTagsToResourceInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Key *string `type:"string" json:",omitempty"`

	Value *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s TagForAddTagsToResourceInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s TagForAddTagsToResourceInput) GoString() string {
	return s.String()
}

// SetKey sets the Key field's value.
func (s *TagForAddTagsToResourceInput) SetKey(v string) *TagForAddTagsToResourceInput {
	s.Key = &v
	return s
}

// SetValue sets the Value field's value.
func (s *TagForAddTagsToResourceInput) SetValue(v string) *TagForAddTagsToResourceInput {
	s.Value = &v
	return s
}

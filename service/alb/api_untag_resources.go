// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package alb

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opUntagResourcesCommon = "UntagResources"

// UntagResourcesCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the UntagResourcesCommon operation. The "output" return
// value will be populated with the UntagResourcesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned UntagResourcesCommon Request to send the API call to the service.
// the "output" return value is not valid until after UntagResourcesCommon Send returns without error.
//
// See UntagResourcesCommon for more information on using the UntagResourcesCommon
// API call, and error handling.
//
//    // Example sending a request using the UntagResourcesCommonRequest method.
//    req, resp := client.UntagResourcesCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ALB) UntagResourcesCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opUntagResourcesCommon,
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

// UntagResourcesCommon API operation for ALB.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ALB's
// API operation UntagResourcesCommon for usage and error information.
func (c *ALB) UntagResourcesCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.UntagResourcesCommonRequest(input)
	return out, req.Send()
}

// UntagResourcesCommonWithContext is the same as UntagResourcesCommon with the addition of
// the ability to pass a context and additional request options.
//
// See UntagResourcesCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ALB) UntagResourcesCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.UntagResourcesCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opUntagResources = "UntagResources"

// UntagResourcesRequest generates a "volcengine/request.Request" representing the
// client's request for the UntagResources operation. The "output" return
// value will be populated with the UntagResourcesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned UntagResourcesCommon Request to send the API call to the service.
// the "output" return value is not valid until after UntagResourcesCommon Send returns without error.
//
// See UntagResources for more information on using the UntagResources
// API call, and error handling.
//
//    // Example sending a request using the UntagResourcesRequest method.
//    req, resp := client.UntagResourcesRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ALB) UntagResourcesRequest(input *UntagResourcesInput) (req *request.Request, output *UntagResourcesOutput) {
	op := &request.Operation{
		Name:       opUntagResources,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UntagResourcesInput{}
	}

	output = &UntagResourcesOutput{}
	req = c.newRequest(op, input, output)

	return
}

// UntagResources API operation for ALB.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ALB's
// API operation UntagResources for usage and error information.
func (c *ALB) UntagResources(input *UntagResourcesInput) (*UntagResourcesOutput, error) {
	req, out := c.UntagResourcesRequest(input)
	return out, req.Send()
}

// UntagResourcesWithContext is the same as UntagResources with the addition of
// the ability to pass a context and additional request options.
//
// See UntagResources for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ALB) UntagResourcesWithContext(ctx volcengine.Context, input *UntagResourcesInput, opts ...request.Option) (*UntagResourcesOutput, error) {
	req, out := c.UntagResourcesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type UntagResourcesInput struct {
	_ struct{} `type:"structure"`

	ResourceIds []*string `type:"list"`

	ResourceType *string `type:"string" enum:"ResourceTypeForUntagResourcesInput"`

	TagKeys []*string `type:"list"`
}

// String returns the string representation
func (s UntagResourcesInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s UntagResourcesInput) GoString() string {
	return s.String()
}

// SetResourceIds sets the ResourceIds field's value.
func (s *UntagResourcesInput) SetResourceIds(v []*string) *UntagResourcesInput {
	s.ResourceIds = v
	return s
}

// SetResourceType sets the ResourceType field's value.
func (s *UntagResourcesInput) SetResourceType(v string) *UntagResourcesInput {
	s.ResourceType = &v
	return s
}

// SetTagKeys sets the TagKeys field's value.
func (s *UntagResourcesInput) SetTagKeys(v []*string) *UntagResourcesInput {
	s.TagKeys = v
	return s
}

type UntagResourcesOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	RequestId *string `type:"string"`
}

// String returns the string representation
func (s UntagResourcesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s UntagResourcesOutput) GoString() string {
	return s.String()
}

// SetRequestId sets the RequestId field's value.
func (s *UntagResourcesOutput) SetRequestId(v string) *UntagResourcesOutput {
	s.RequestId = &v
	return s
}

const (
	// ResourceTypeForUntagResourcesInputLoadbalancer is a ResourceTypeForUntagResourcesInput enum value
	ResourceTypeForUntagResourcesInputLoadbalancer = "loadbalancer"

	// ResourceTypeForUntagResourcesInputListener is a ResourceTypeForUntagResourcesInput enum value
	ResourceTypeForUntagResourcesInputListener = "listener"

	// ResourceTypeForUntagResourcesInputServergroup is a ResourceTypeForUntagResourcesInput enum value
	ResourceTypeForUntagResourcesInputServergroup = "servergroup"
)

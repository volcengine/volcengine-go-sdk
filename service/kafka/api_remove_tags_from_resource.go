// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package kafka

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opRemoveTagsFromResourceCommon = "RemoveTagsFromResource"

// RemoveTagsFromResourceCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the RemoveTagsFromResourceCommon operation. The "output" return
// value will be populated with the RemoveTagsFromResourceCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned RemoveTagsFromResourceCommon Request to send the API call to the service.
// the "output" return value is not valid until after RemoveTagsFromResourceCommon Send returns without error.
//
// See RemoveTagsFromResourceCommon for more information on using the RemoveTagsFromResourceCommon
// API call, and error handling.
//
//    // Example sending a request using the RemoveTagsFromResourceCommonRequest method.
//    req, resp := client.RemoveTagsFromResourceCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *KAFKA) RemoveTagsFromResourceCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opRemoveTagsFromResourceCommon,
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

// RemoveTagsFromResourceCommon API operation for KAFKA.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for KAFKA's
// API operation RemoveTagsFromResourceCommon for usage and error information.
func (c *KAFKA) RemoveTagsFromResourceCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.RemoveTagsFromResourceCommonRequest(input)
	return out, req.Send()
}

// RemoveTagsFromResourceCommonWithContext is the same as RemoveTagsFromResourceCommon with the addition of
// the ability to pass a context and additional request options.
//
// See RemoveTagsFromResourceCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *KAFKA) RemoveTagsFromResourceCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.RemoveTagsFromResourceCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opRemoveTagsFromResource = "RemoveTagsFromResource"

// RemoveTagsFromResourceRequest generates a "volcengine/request.Request" representing the
// client's request for the RemoveTagsFromResource operation. The "output" return
// value will be populated with the RemoveTagsFromResourceCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned RemoveTagsFromResourceCommon Request to send the API call to the service.
// the "output" return value is not valid until after RemoveTagsFromResourceCommon Send returns without error.
//
// See RemoveTagsFromResource for more information on using the RemoveTagsFromResource
// API call, and error handling.
//
//    // Example sending a request using the RemoveTagsFromResourceRequest method.
//    req, resp := client.RemoveTagsFromResourceRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *KAFKA) RemoveTagsFromResourceRequest(input *RemoveTagsFromResourceInput) (req *request.Request, output *RemoveTagsFromResourceOutput) {
	op := &request.Operation{
		Name:       opRemoveTagsFromResource,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &RemoveTagsFromResourceInput{}
	}

	output = &RemoveTagsFromResourceOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// RemoveTagsFromResource API operation for KAFKA.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for KAFKA's
// API operation RemoveTagsFromResource for usage and error information.
func (c *KAFKA) RemoveTagsFromResource(input *RemoveTagsFromResourceInput) (*RemoveTagsFromResourceOutput, error) {
	req, out := c.RemoveTagsFromResourceRequest(input)
	return out, req.Send()
}

// RemoveTagsFromResourceWithContext is the same as RemoveTagsFromResource with the addition of
// the ability to pass a context and additional request options.
//
// See RemoveTagsFromResource for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *KAFKA) RemoveTagsFromResourceWithContext(ctx volcengine.Context, input *RemoveTagsFromResourceInput, opts ...request.Option) (*RemoveTagsFromResourceOutput, error) {
	req, out := c.RemoveTagsFromResourceRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type RemoveTagsFromResourceInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	All *bool `type:"boolean" json:",omitempty"`

	InstanceIds []*string `type:"list" json:",omitempty"`

	TagKeys []*string `type:"list" json:",omitempty"`
}

// String returns the string representation
func (s RemoveTagsFromResourceInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s RemoveTagsFromResourceInput) GoString() string {
	return s.String()
}

// SetAll sets the All field's value.
func (s *RemoveTagsFromResourceInput) SetAll(v bool) *RemoveTagsFromResourceInput {
	s.All = &v
	return s
}

// SetInstanceIds sets the InstanceIds field's value.
func (s *RemoveTagsFromResourceInput) SetInstanceIds(v []*string) *RemoveTagsFromResourceInput {
	s.InstanceIds = v
	return s
}

// SetTagKeys sets the TagKeys field's value.
func (s *RemoveTagsFromResourceInput) SetTagKeys(v []*string) *RemoveTagsFromResourceInput {
	s.TagKeys = v
	return s
}

type RemoveTagsFromResourceOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s RemoveTagsFromResourceOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s RemoveTagsFromResourceOutput) GoString() string {
	return s.String()
}

// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package vke

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDeleteKubeconfigsCommon = "DeleteKubeconfigs"

// DeleteKubeconfigsCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DeleteKubeconfigsCommon operation. The "output" return
// value will be populated with the DeleteKubeconfigsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeleteKubeconfigsCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeleteKubeconfigsCommon Send returns without error.
//
// See DeleteKubeconfigsCommon for more information on using the DeleteKubeconfigsCommon
// API call, and error handling.
//
//    // Example sending a request using the DeleteKubeconfigsCommonRequest method.
//    req, resp := client.DeleteKubeconfigsCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VKE) DeleteKubeconfigsCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDeleteKubeconfigsCommon,
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

// DeleteKubeconfigsCommon API operation for VKE.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VKE's
// API operation DeleteKubeconfigsCommon for usage and error information.
func (c *VKE) DeleteKubeconfigsCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DeleteKubeconfigsCommonRequest(input)
	return out, req.Send()
}

// DeleteKubeconfigsCommonWithContext is the same as DeleteKubeconfigsCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteKubeconfigsCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VKE) DeleteKubeconfigsCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DeleteKubeconfigsCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDeleteKubeconfigs = "DeleteKubeconfigs"

// DeleteKubeconfigsRequest generates a "volcengine/request.Request" representing the
// client's request for the DeleteKubeconfigs operation. The "output" return
// value will be populated with the DeleteKubeconfigsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeleteKubeconfigsCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeleteKubeconfigsCommon Send returns without error.
//
// See DeleteKubeconfigs for more information on using the DeleteKubeconfigs
// API call, and error handling.
//
//    // Example sending a request using the DeleteKubeconfigsRequest method.
//    req, resp := client.DeleteKubeconfigsRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VKE) DeleteKubeconfigsRequest(input *DeleteKubeconfigsInput) (req *request.Request, output *DeleteKubeconfigsOutput) {
	op := &request.Operation{
		Name:       opDeleteKubeconfigs,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteKubeconfigsInput{}
	}

	output = &DeleteKubeconfigsOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// DeleteKubeconfigs API operation for VKE.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VKE's
// API operation DeleteKubeconfigs for usage and error information.
func (c *VKE) DeleteKubeconfigs(input *DeleteKubeconfigsInput) (*DeleteKubeconfigsOutput, error) {
	req, out := c.DeleteKubeconfigsRequest(input)
	return out, req.Send()
}

// DeleteKubeconfigsWithContext is the same as DeleteKubeconfigs with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteKubeconfigs for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VKE) DeleteKubeconfigsWithContext(ctx volcengine.Context, input *DeleteKubeconfigsInput, opts ...request.Option) (*DeleteKubeconfigsOutput, error) {
	req, out := c.DeleteKubeconfigsRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DeleteKubeconfigsInput struct {
	_ struct{} `type:"structure"`

	ClusterId *string `type:"string"`

	Ids []*string `type:"list"`
}

// String returns the string representation
func (s DeleteKubeconfigsInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteKubeconfigsInput) GoString() string {
	return s.String()
}

// SetClusterId sets the ClusterId field's value.
func (s *DeleteKubeconfigsInput) SetClusterId(v string) *DeleteKubeconfigsInput {
	s.ClusterId = &v
	return s
}

// SetIds sets the Ids field's value.
func (s *DeleteKubeconfigsInput) SetIds(v []*string) *DeleteKubeconfigsInput {
	s.Ids = v
	return s
}

type DeleteKubeconfigsOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s DeleteKubeconfigsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteKubeconfigsOutput) GoString() string {
	return s.String()
}

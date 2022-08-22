// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package cr

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opStartRegistryCommon = "StartRegistry"

// StartRegistryCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the StartRegistryCommon operation. The "output" return
// value will be populated with the StartRegistryCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned StartRegistryCommon Request to send the API call to the service.
// the "output" return value is not valid until after StartRegistryCommon Send returns without error.
//
// See StartRegistryCommon for more information on using the StartRegistryCommon
// API call, and error handling.
//
//    // Example sending a request using the StartRegistryCommonRequest method.
//    req, resp := client.StartRegistryCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CR) StartRegistryCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opStartRegistryCommon,
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

// StartRegistryCommon API operation for CR.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CR's
// API operation StartRegistryCommon for usage and error information.
func (c *CR) StartRegistryCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.StartRegistryCommonRequest(input)
	return out, req.Send()
}

// StartRegistryCommonWithContext is the same as StartRegistryCommon with the addition of
// the ability to pass a context and additional request options.
//
// See StartRegistryCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CR) StartRegistryCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.StartRegistryCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opStartRegistry = "StartRegistry"

// StartRegistryRequest generates a "volcengine/request.Request" representing the
// client's request for the StartRegistry operation. The "output" return
// value will be populated with the StartRegistryCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned StartRegistryCommon Request to send the API call to the service.
// the "output" return value is not valid until after StartRegistryCommon Send returns without error.
//
// See StartRegistry for more information on using the StartRegistry
// API call, and error handling.
//
//    // Example sending a request using the StartRegistryRequest method.
//    req, resp := client.StartRegistryRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CR) StartRegistryRequest(input *StartRegistryInput) (req *request.Request, output *StartRegistryOutput) {
	op := &request.Operation{
		Name:       opStartRegistry,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &StartRegistryInput{}
	}

	output = &StartRegistryOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// StartRegistry API operation for CR.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CR's
// API operation StartRegistry for usage and error information.
func (c *CR) StartRegistry(input *StartRegistryInput) (*StartRegistryOutput, error) {
	req, out := c.StartRegistryRequest(input)
	return out, req.Send()
}

// StartRegistryWithContext is the same as StartRegistry with the addition of
// the ability to pass a context and additional request options.
//
// See StartRegistry for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CR) StartRegistryWithContext(ctx volcengine.Context, input *StartRegistryInput, opts ...request.Option) (*StartRegistryOutput, error) {
	req, out := c.StartRegistryRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type StartRegistryInput struct {
	_ struct{} `type:"structure"`

	Name *string `type:"string"`
}

// String returns the string representation
func (s StartRegistryInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s StartRegistryInput) GoString() string {
	return s.String()
}

// SetName sets the Name field's value.
func (s *StartRegistryInput) SetName(v string) *StartRegistryInput {
	s.Name = &v
	return s
}

type StartRegistryOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s StartRegistryOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s StartRegistryOutput) GoString() string {
	return s.String()
}
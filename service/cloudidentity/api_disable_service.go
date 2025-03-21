// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudidentity

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDisableServiceCommon = "DisableService"

// DisableServiceCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DisableServiceCommon operation. The "output" return
// value will be populated with the DisableServiceCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DisableServiceCommon Request to send the API call to the service.
// the "output" return value is not valid until after DisableServiceCommon Send returns without error.
//
// See DisableServiceCommon for more information on using the DisableServiceCommon
// API call, and error handling.
//
//    // Example sending a request using the DisableServiceCommonRequest method.
//    req, resp := client.DisableServiceCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CLOUDIDENTITY) DisableServiceCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDisableServiceCommon,
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

// DisableServiceCommon API operation for CLOUDIDENTITY.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CLOUDIDENTITY's
// API operation DisableServiceCommon for usage and error information.
func (c *CLOUDIDENTITY) DisableServiceCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DisableServiceCommonRequest(input)
	return out, req.Send()
}

// DisableServiceCommonWithContext is the same as DisableServiceCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DisableServiceCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CLOUDIDENTITY) DisableServiceCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DisableServiceCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDisableService = "DisableService"

// DisableServiceRequest generates a "volcengine/request.Request" representing the
// client's request for the DisableService operation. The "output" return
// value will be populated with the DisableServiceCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DisableServiceCommon Request to send the API call to the service.
// the "output" return value is not valid until after DisableServiceCommon Send returns without error.
//
// See DisableService for more information on using the DisableService
// API call, and error handling.
//
//    // Example sending a request using the DisableServiceRequest method.
//    req, resp := client.DisableServiceRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CLOUDIDENTITY) DisableServiceRequest(input *DisableServiceInput) (req *request.Request, output *DisableServiceOutput) {
	op := &request.Operation{
		Name:       opDisableService,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DisableServiceInput{}
	}

	output = &DisableServiceOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// DisableService API operation for CLOUDIDENTITY.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CLOUDIDENTITY's
// API operation DisableService for usage and error information.
func (c *CLOUDIDENTITY) DisableService(input *DisableServiceInput) (*DisableServiceOutput, error) {
	req, out := c.DisableServiceRequest(input)
	return out, req.Send()
}

// DisableServiceWithContext is the same as DisableService with the addition of
// the ability to pass a context and additional request options.
//
// See DisableService for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CLOUDIDENTITY) DisableServiceWithContext(ctx volcengine.Context, input *DisableServiceInput, opts ...request.Option) (*DisableServiceOutput, error) {
	req, out := c.DisableServiceRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DisableServiceInput struct {
	_ struct{} `type:"structure" json:",omitempty"`
}

// String returns the string representation
func (s DisableServiceInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DisableServiceInput) GoString() string {
	return s.String()
}

type DisableServiceOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s DisableServiceOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DisableServiceOutput) GoString() string {
	return s.String()
}

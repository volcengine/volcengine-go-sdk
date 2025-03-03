// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package vke

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opGetGlobalDefaultDeleteOptionCommon = "GetGlobalDefaultDeleteOption"

// GetGlobalDefaultDeleteOptionCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the GetGlobalDefaultDeleteOptionCommon operation. The "output" return
// value will be populated with the GetGlobalDefaultDeleteOptionCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned GetGlobalDefaultDeleteOptionCommon Request to send the API call to the service.
// the "output" return value is not valid until after GetGlobalDefaultDeleteOptionCommon Send returns without error.
//
// See GetGlobalDefaultDeleteOptionCommon for more information on using the GetGlobalDefaultDeleteOptionCommon
// API call, and error handling.
//
//    // Example sending a request using the GetGlobalDefaultDeleteOptionCommonRequest method.
//    req, resp := client.GetGlobalDefaultDeleteOptionCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VKE) GetGlobalDefaultDeleteOptionCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opGetGlobalDefaultDeleteOptionCommon,
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

// GetGlobalDefaultDeleteOptionCommon API operation for VKE.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VKE's
// API operation GetGlobalDefaultDeleteOptionCommon for usage and error information.
func (c *VKE) GetGlobalDefaultDeleteOptionCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.GetGlobalDefaultDeleteOptionCommonRequest(input)
	return out, req.Send()
}

// GetGlobalDefaultDeleteOptionCommonWithContext is the same as GetGlobalDefaultDeleteOptionCommon with the addition of
// the ability to pass a context and additional request options.
//
// See GetGlobalDefaultDeleteOptionCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VKE) GetGlobalDefaultDeleteOptionCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.GetGlobalDefaultDeleteOptionCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opGetGlobalDefaultDeleteOption = "GetGlobalDefaultDeleteOption"

// GetGlobalDefaultDeleteOptionRequest generates a "volcengine/request.Request" representing the
// client's request for the GetGlobalDefaultDeleteOption operation. The "output" return
// value will be populated with the GetGlobalDefaultDeleteOptionCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned GetGlobalDefaultDeleteOptionCommon Request to send the API call to the service.
// the "output" return value is not valid until after GetGlobalDefaultDeleteOptionCommon Send returns without error.
//
// See GetGlobalDefaultDeleteOption for more information on using the GetGlobalDefaultDeleteOption
// API call, and error handling.
//
//    // Example sending a request using the GetGlobalDefaultDeleteOptionRequest method.
//    req, resp := client.GetGlobalDefaultDeleteOptionRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VKE) GetGlobalDefaultDeleteOptionRequest(input *GetGlobalDefaultDeleteOptionInput) (req *request.Request, output *GetGlobalDefaultDeleteOptionOutput) {
	op := &request.Operation{
		Name:       opGetGlobalDefaultDeleteOption,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetGlobalDefaultDeleteOptionInput{}
	}

	output = &GetGlobalDefaultDeleteOptionOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// GetGlobalDefaultDeleteOption API operation for VKE.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VKE's
// API operation GetGlobalDefaultDeleteOption for usage and error information.
func (c *VKE) GetGlobalDefaultDeleteOption(input *GetGlobalDefaultDeleteOptionInput) (*GetGlobalDefaultDeleteOptionOutput, error) {
	req, out := c.GetGlobalDefaultDeleteOptionRequest(input)
	return out, req.Send()
}

// GetGlobalDefaultDeleteOptionWithContext is the same as GetGlobalDefaultDeleteOption with the addition of
// the ability to pass a context and additional request options.
//
// See GetGlobalDefaultDeleteOption for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VKE) GetGlobalDefaultDeleteOptionWithContext(ctx volcengine.Context, input *GetGlobalDefaultDeleteOptionInput, opts ...request.Option) (*GetGlobalDefaultDeleteOptionOutput, error) {
	req, out := c.GetGlobalDefaultDeleteOptionRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type GetGlobalDefaultDeleteOptionInput struct {
	_ struct{} `type:"structure" json:",omitempty"`
}

// String returns the string representation
func (s GetGlobalDefaultDeleteOptionInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s GetGlobalDefaultDeleteOptionInput) GoString() string {
	return s.String()
}

type GetGlobalDefaultDeleteOptionOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	DefaultDeleteAllResources *bool `type:"boolean" json:",omitempty"`

	LastChangeTime *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s GetGlobalDefaultDeleteOptionOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s GetGlobalDefaultDeleteOptionOutput) GoString() string {
	return s.String()
}

// SetDefaultDeleteAllResources sets the DefaultDeleteAllResources field's value.
func (s *GetGlobalDefaultDeleteOptionOutput) SetDefaultDeleteAllResources(v bool) *GetGlobalDefaultDeleteOptionOutput {
	s.DefaultDeleteAllResources = &v
	return s
}

// SetLastChangeTime sets the LastChangeTime field's value.
func (s *GetGlobalDefaultDeleteOptionOutput) SetLastChangeTime(v string) *GetGlobalDefaultDeleteOptionOutput {
	s.LastChangeTime = &v
	return s
}

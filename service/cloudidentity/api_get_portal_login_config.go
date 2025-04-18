// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudidentity

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opGetPortalLoginConfigCommon = "GetPortalLoginConfig"

// GetPortalLoginConfigCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the GetPortalLoginConfigCommon operation. The "output" return
// value will be populated with the GetPortalLoginConfigCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned GetPortalLoginConfigCommon Request to send the API call to the service.
// the "output" return value is not valid until after GetPortalLoginConfigCommon Send returns without error.
//
// See GetPortalLoginConfigCommon for more information on using the GetPortalLoginConfigCommon
// API call, and error handling.
//
//    // Example sending a request using the GetPortalLoginConfigCommonRequest method.
//    req, resp := client.GetPortalLoginConfigCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CLOUDIDENTITY) GetPortalLoginConfigCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opGetPortalLoginConfigCommon,
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

// GetPortalLoginConfigCommon API operation for CLOUDIDENTITY.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CLOUDIDENTITY's
// API operation GetPortalLoginConfigCommon for usage and error information.
func (c *CLOUDIDENTITY) GetPortalLoginConfigCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.GetPortalLoginConfigCommonRequest(input)
	return out, req.Send()
}

// GetPortalLoginConfigCommonWithContext is the same as GetPortalLoginConfigCommon with the addition of
// the ability to pass a context and additional request options.
//
// See GetPortalLoginConfigCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CLOUDIDENTITY) GetPortalLoginConfigCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.GetPortalLoginConfigCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opGetPortalLoginConfig = "GetPortalLoginConfig"

// GetPortalLoginConfigRequest generates a "volcengine/request.Request" representing the
// client's request for the GetPortalLoginConfig operation. The "output" return
// value will be populated with the GetPortalLoginConfigCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned GetPortalLoginConfigCommon Request to send the API call to the service.
// the "output" return value is not valid until after GetPortalLoginConfigCommon Send returns without error.
//
// See GetPortalLoginConfig for more information on using the GetPortalLoginConfig
// API call, and error handling.
//
//    // Example sending a request using the GetPortalLoginConfigRequest method.
//    req, resp := client.GetPortalLoginConfigRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CLOUDIDENTITY) GetPortalLoginConfigRequest(input *GetPortalLoginConfigInput) (req *request.Request, output *GetPortalLoginConfigOutput) {
	op := &request.Operation{
		Name:       opGetPortalLoginConfig,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetPortalLoginConfigInput{}
	}

	output = &GetPortalLoginConfigOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// GetPortalLoginConfig API operation for CLOUDIDENTITY.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CLOUDIDENTITY's
// API operation GetPortalLoginConfig for usage and error information.
func (c *CLOUDIDENTITY) GetPortalLoginConfig(input *GetPortalLoginConfigInput) (*GetPortalLoginConfigOutput, error) {
	req, out := c.GetPortalLoginConfigRequest(input)
	return out, req.Send()
}

// GetPortalLoginConfigWithContext is the same as GetPortalLoginConfig with the addition of
// the ability to pass a context and additional request options.
//
// See GetPortalLoginConfig for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CLOUDIDENTITY) GetPortalLoginConfigWithContext(ctx volcengine.Context, input *GetPortalLoginConfigInput, opts ...request.Option) (*GetPortalLoginConfigOutput, error) {
	req, out := c.GetPortalLoginConfigRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type GetPortalLoginConfigInput struct {
	_ struct{} `type:"structure" json:",omitempty"`
}

// String returns the string representation
func (s GetPortalLoginConfigInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s GetPortalLoginConfigInput) GoString() string {
	return s.String()
}

type GetPortalLoginConfigOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	LoginType *string `type:"string" json:",omitempty"`

	PortalURL *string `type:"string" json:",omitempty"`

	Subdomain *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s GetPortalLoginConfigOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s GetPortalLoginConfigOutput) GoString() string {
	return s.String()
}

// SetLoginType sets the LoginType field's value.
func (s *GetPortalLoginConfigOutput) SetLoginType(v string) *GetPortalLoginConfigOutput {
	s.LoginType = &v
	return s
}

// SetPortalURL sets the PortalURL field's value.
func (s *GetPortalLoginConfigOutput) SetPortalURL(v string) *GetPortalLoginConfigOutput {
	s.PortalURL = &v
	return s
}

// SetSubdomain sets the Subdomain field's value.
func (s *GetPortalLoginConfigOutput) SetSubdomain(v string) *GetPortalLoginConfigOutput {
	s.Subdomain = &v
	return s
}

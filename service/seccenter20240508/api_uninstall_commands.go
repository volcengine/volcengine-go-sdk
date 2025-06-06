// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package seccenter20240508

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opUninstallCommandsCommon = "UninstallCommands"

// UninstallCommandsCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the UninstallCommandsCommon operation. The "output" return
// value will be populated with the UninstallCommandsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned UninstallCommandsCommon Request to send the API call to the service.
// the "output" return value is not valid until after UninstallCommandsCommon Send returns without error.
//
// See UninstallCommandsCommon for more information on using the UninstallCommandsCommon
// API call, and error handling.
//
//    // Example sending a request using the UninstallCommandsCommonRequest method.
//    req, resp := client.UninstallCommandsCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *SECCENTER20240508) UninstallCommandsCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opUninstallCommandsCommon,
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

// UninstallCommandsCommon API operation for SECCENTER20240508.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for SECCENTER20240508's
// API operation UninstallCommandsCommon for usage and error information.
func (c *SECCENTER20240508) UninstallCommandsCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.UninstallCommandsCommonRequest(input)
	return out, req.Send()
}

// UninstallCommandsCommonWithContext is the same as UninstallCommandsCommon with the addition of
// the ability to pass a context and additional request options.
//
// See UninstallCommandsCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *SECCENTER20240508) UninstallCommandsCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.UninstallCommandsCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opUninstallCommands = "UninstallCommands"

// UninstallCommandsRequest generates a "volcengine/request.Request" representing the
// client's request for the UninstallCommands operation. The "output" return
// value will be populated with the UninstallCommandsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned UninstallCommandsCommon Request to send the API call to the service.
// the "output" return value is not valid until after UninstallCommandsCommon Send returns without error.
//
// See UninstallCommands for more information on using the UninstallCommands
// API call, and error handling.
//
//    // Example sending a request using the UninstallCommandsRequest method.
//    req, resp := client.UninstallCommandsRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *SECCENTER20240508) UninstallCommandsRequest(input *UninstallCommandsInput) (req *request.Request, output *UninstallCommandsOutput) {
	op := &request.Operation{
		Name:       opUninstallCommands,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UninstallCommandsInput{}
	}

	output = &UninstallCommandsOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// UninstallCommands API operation for SECCENTER20240508.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for SECCENTER20240508's
// API operation UninstallCommands for usage and error information.
func (c *SECCENTER20240508) UninstallCommands(input *UninstallCommandsInput) (*UninstallCommandsOutput, error) {
	req, out := c.UninstallCommandsRequest(input)
	return out, req.Send()
}

// UninstallCommandsWithContext is the same as UninstallCommands with the addition of
// the ability to pass a context and additional request options.
//
// See UninstallCommands for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *SECCENTER20240508) UninstallCommandsWithContext(ctx volcengine.Context, input *UninstallCommandsInput, opts ...request.Option) (*UninstallCommandsOutput, error) {
	req, out := c.UninstallCommandsRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type UninstallCommandsInput struct {
	_ struct{} `type:"structure" json:",omitempty"`
}

// String returns the string representation
func (s UninstallCommandsInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s UninstallCommandsInput) GoString() string {
	return s.String()
}

type UninstallCommandsOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	NotVolcLinux *string `type:"string" json:",omitempty"`

	VolcLinux *string `type:"string" json:",omitempty"`

	Windows *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s UninstallCommandsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s UninstallCommandsOutput) GoString() string {
	return s.String()
}

// SetNotVolcLinux sets the NotVolcLinux field's value.
func (s *UninstallCommandsOutput) SetNotVolcLinux(v string) *UninstallCommandsOutput {
	s.NotVolcLinux = &v
	return s
}

// SetVolcLinux sets the VolcLinux field's value.
func (s *UninstallCommandsOutput) SetVolcLinux(v string) *UninstallCommandsOutput {
	s.VolcLinux = &v
	return s
}

// SetWindows sets the Windows field's value.
func (s *UninstallCommandsOutput) SetWindows(v string) *UninstallCommandsOutput {
	s.Windows = &v
	return s
}

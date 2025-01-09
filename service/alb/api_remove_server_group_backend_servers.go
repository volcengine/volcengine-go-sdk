// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package alb

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opRemoveServerGroupBackendServersCommon = "RemoveServerGroupBackendServers"

// RemoveServerGroupBackendServersCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the RemoveServerGroupBackendServersCommon operation. The "output" return
// value will be populated with the RemoveServerGroupBackendServersCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned RemoveServerGroupBackendServersCommon Request to send the API call to the service.
// the "output" return value is not valid until after RemoveServerGroupBackendServersCommon Send returns without error.
//
// See RemoveServerGroupBackendServersCommon for more information on using the RemoveServerGroupBackendServersCommon
// API call, and error handling.
//
//    // Example sending a request using the RemoveServerGroupBackendServersCommonRequest method.
//    req, resp := client.RemoveServerGroupBackendServersCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ALB) RemoveServerGroupBackendServersCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opRemoveServerGroupBackendServersCommon,
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

// RemoveServerGroupBackendServersCommon API operation for ALB.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ALB's
// API operation RemoveServerGroupBackendServersCommon for usage and error information.
func (c *ALB) RemoveServerGroupBackendServersCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.RemoveServerGroupBackendServersCommonRequest(input)
	return out, req.Send()
}

// RemoveServerGroupBackendServersCommonWithContext is the same as RemoveServerGroupBackendServersCommon with the addition of
// the ability to pass a context and additional request options.
//
// See RemoveServerGroupBackendServersCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ALB) RemoveServerGroupBackendServersCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.RemoveServerGroupBackendServersCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opRemoveServerGroupBackendServers = "RemoveServerGroupBackendServers"

// RemoveServerGroupBackendServersRequest generates a "volcengine/request.Request" representing the
// client's request for the RemoveServerGroupBackendServers operation. The "output" return
// value will be populated with the RemoveServerGroupBackendServersCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned RemoveServerGroupBackendServersCommon Request to send the API call to the service.
// the "output" return value is not valid until after RemoveServerGroupBackendServersCommon Send returns without error.
//
// See RemoveServerGroupBackendServers for more information on using the RemoveServerGroupBackendServers
// API call, and error handling.
//
//    // Example sending a request using the RemoveServerGroupBackendServersRequest method.
//    req, resp := client.RemoveServerGroupBackendServersRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ALB) RemoveServerGroupBackendServersRequest(input *RemoveServerGroupBackendServersInput) (req *request.Request, output *RemoveServerGroupBackendServersOutput) {
	op := &request.Operation{
		Name:       opRemoveServerGroupBackendServers,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &RemoveServerGroupBackendServersInput{}
	}

	output = &RemoveServerGroupBackendServersOutput{}
	req = c.newRequest(op, input, output)

	return
}

// RemoveServerGroupBackendServers API operation for ALB.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ALB's
// API operation RemoveServerGroupBackendServers for usage and error information.
func (c *ALB) RemoveServerGroupBackendServers(input *RemoveServerGroupBackendServersInput) (*RemoveServerGroupBackendServersOutput, error) {
	req, out := c.RemoveServerGroupBackendServersRequest(input)
	return out, req.Send()
}

// RemoveServerGroupBackendServersWithContext is the same as RemoveServerGroupBackendServers with the addition of
// the ability to pass a context and additional request options.
//
// See RemoveServerGroupBackendServers for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ALB) RemoveServerGroupBackendServersWithContext(ctx volcengine.Context, input *RemoveServerGroupBackendServersInput, opts ...request.Option) (*RemoveServerGroupBackendServersOutput, error) {
	req, out := c.RemoveServerGroupBackendServersRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type RemoveServerGroupBackendServersInput struct {
	_ struct{} `type:"structure"`

	ServerGroupId *string `type:"string"`

	ServerIds []*string `type:"list"`
}

// String returns the string representation
func (s RemoveServerGroupBackendServersInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s RemoveServerGroupBackendServersInput) GoString() string {
	return s.String()
}

// SetServerGroupId sets the ServerGroupId field's value.
func (s *RemoveServerGroupBackendServersInput) SetServerGroupId(v string) *RemoveServerGroupBackendServersInput {
	s.ServerGroupId = &v
	return s
}

// SetServerIds sets the ServerIds field's value.
func (s *RemoveServerGroupBackendServersInput) SetServerIds(v []*string) *RemoveServerGroupBackendServersInput {
	s.ServerIds = v
	return s
}

type RemoveServerGroupBackendServersOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	RequestId *string `type:"string"`
}

// String returns the string representation
func (s RemoveServerGroupBackendServersOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s RemoveServerGroupBackendServersOutput) GoString() string {
	return s.String()
}

// SetRequestId sets the RequestId field's value.
func (s *RemoveServerGroupBackendServersOutput) SetRequestId(v string) *RemoveServerGroupBackendServersOutput {
	s.RequestId = &v
	return s
}

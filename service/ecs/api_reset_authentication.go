// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package ecs

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opResetAuthenticationCommon = "ResetAuthentication"

// ResetAuthenticationCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ResetAuthenticationCommon operation. The "output" return
// value will be populated with the ResetAuthenticationCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ResetAuthenticationCommon Request to send the API call to the service.
// the "output" return value is not valid until after ResetAuthenticationCommon Send returns without error.
//
// See ResetAuthenticationCommon for more information on using the ResetAuthenticationCommon
// API call, and error handling.
//
//    // Example sending a request using the ResetAuthenticationCommonRequest method.
//    req, resp := client.ResetAuthenticationCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ECS) ResetAuthenticationCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opResetAuthenticationCommon,
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

// ResetAuthenticationCommon API operation for ECS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ECS's
// API operation ResetAuthenticationCommon for usage and error information.
func (c *ECS) ResetAuthenticationCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ResetAuthenticationCommonRequest(input)
	return out, req.Send()
}

// ResetAuthenticationCommonWithContext is the same as ResetAuthenticationCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ResetAuthenticationCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ECS) ResetAuthenticationCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ResetAuthenticationCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opResetAuthentication = "ResetAuthentication"

// ResetAuthenticationRequest generates a "volcengine/request.Request" representing the
// client's request for the ResetAuthentication operation. The "output" return
// value will be populated with the ResetAuthenticationCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ResetAuthenticationCommon Request to send the API call to the service.
// the "output" return value is not valid until after ResetAuthenticationCommon Send returns without error.
//
// See ResetAuthentication for more information on using the ResetAuthentication
// API call, and error handling.
//
//    // Example sending a request using the ResetAuthenticationRequest method.
//    req, resp := client.ResetAuthenticationRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ECS) ResetAuthenticationRequest(input *ResetAuthenticationInput) (req *request.Request, output *ResetAuthenticationOutput) {
	op := &request.Operation{
		Name:       opResetAuthentication,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ResetAuthenticationInput{}
	}

	output = &ResetAuthenticationOutput{}
	req = c.newRequest(op, input, output)

	return
}

// ResetAuthentication API operation for ECS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ECS's
// API operation ResetAuthentication for usage and error information.
func (c *ECS) ResetAuthentication(input *ResetAuthenticationInput) (*ResetAuthenticationOutput, error) {
	req, out := c.ResetAuthenticationRequest(input)
	return out, req.Send()
}

// ResetAuthenticationWithContext is the same as ResetAuthentication with the addition of
// the ability to pass a context and additional request options.
//
// See ResetAuthentication for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ECS) ResetAuthenticationWithContext(ctx volcengine.Context, input *ResetAuthenticationInput, opts ...request.Option) (*ResetAuthenticationOutput, error) {
	req, out := c.ResetAuthenticationRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ResetAuthenticationInput struct {
	_ struct{} `type:"structure"`

	InstanceId *string `type:"string"`

	Password *string `type:"string"`

	Pubkey *string `type:"string"`
}

// String returns the string representation
func (s ResetAuthenticationInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ResetAuthenticationInput) GoString() string {
	return s.String()
}

// SetInstanceId sets the InstanceId field's value.
func (s *ResetAuthenticationInput) SetInstanceId(v string) *ResetAuthenticationInput {
	s.InstanceId = &v
	return s
}

// SetPassword sets the Password field's value.
func (s *ResetAuthenticationInput) SetPassword(v string) *ResetAuthenticationInput {
	s.Password = &v
	return s
}

// SetPubkey sets the Pubkey field's value.
func (s *ResetAuthenticationInput) SetPubkey(v string) *ResetAuthenticationInput {
	s.Pubkey = &v
	return s
}

type ResetAuthenticationOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s ResetAuthenticationOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ResetAuthenticationOutput) GoString() string {
	return s.String()
}
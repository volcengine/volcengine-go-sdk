// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package seccenter20240508

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opInstallAgentClientCommon = "InstallAgentClient"

// InstallAgentClientCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the InstallAgentClientCommon operation. The "output" return
// value will be populated with the InstallAgentClientCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned InstallAgentClientCommon Request to send the API call to the service.
// the "output" return value is not valid until after InstallAgentClientCommon Send returns without error.
//
// See InstallAgentClientCommon for more information on using the InstallAgentClientCommon
// API call, and error handling.
//
//    // Example sending a request using the InstallAgentClientCommonRequest method.
//    req, resp := client.InstallAgentClientCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *SECCENTER20240508) InstallAgentClientCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opInstallAgentClientCommon,
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

// InstallAgentClientCommon API operation for SECCENTER20240508.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for SECCENTER20240508's
// API operation InstallAgentClientCommon for usage and error information.
func (c *SECCENTER20240508) InstallAgentClientCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.InstallAgentClientCommonRequest(input)
	return out, req.Send()
}

// InstallAgentClientCommonWithContext is the same as InstallAgentClientCommon with the addition of
// the ability to pass a context and additional request options.
//
// See InstallAgentClientCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *SECCENTER20240508) InstallAgentClientCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.InstallAgentClientCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opInstallAgentClient = "InstallAgentClient"

// InstallAgentClientRequest generates a "volcengine/request.Request" representing the
// client's request for the InstallAgentClient operation. The "output" return
// value will be populated with the InstallAgentClientCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned InstallAgentClientCommon Request to send the API call to the service.
// the "output" return value is not valid until after InstallAgentClientCommon Send returns without error.
//
// See InstallAgentClient for more information on using the InstallAgentClient
// API call, and error handling.
//
//    // Example sending a request using the InstallAgentClientRequest method.
//    req, resp := client.InstallAgentClientRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *SECCENTER20240508) InstallAgentClientRequest(input *InstallAgentClientInput) (req *request.Request, output *InstallAgentClientOutput) {
	op := &request.Operation{
		Name:       opInstallAgentClient,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &InstallAgentClientInput{}
	}

	output = &InstallAgentClientOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// InstallAgentClient API operation for SECCENTER20240508.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for SECCENTER20240508's
// API operation InstallAgentClient for usage and error information.
func (c *SECCENTER20240508) InstallAgentClient(input *InstallAgentClientInput) (*InstallAgentClientOutput, error) {
	req, out := c.InstallAgentClientRequest(input)
	return out, req.Send()
}

// InstallAgentClientWithContext is the same as InstallAgentClient with the addition of
// the ability to pass a context and additional request options.
//
// See InstallAgentClient for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *SECCENTER20240508) InstallAgentClientWithContext(ctx volcengine.Context, input *InstallAgentClientInput, opts ...request.Option) (*InstallAgentClientOutput, error) {
	req, out := c.InstallAgentClientRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DetailForInstallAgentClientOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	AgentID *string `type:"string" json:",omitempty"`

	Code *int32 `type:"int32" json:",omitempty"`

	Msg *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s DetailForInstallAgentClientOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DetailForInstallAgentClientOutput) GoString() string {
	return s.String()
}

// SetAgentID sets the AgentID field's value.
func (s *DetailForInstallAgentClientOutput) SetAgentID(v string) *DetailForInstallAgentClientOutput {
	s.AgentID = &v
	return s
}

// SetCode sets the Code field's value.
func (s *DetailForInstallAgentClientOutput) SetCode(v int32) *DetailForInstallAgentClientOutput {
	s.Code = &v
	return s
}

// SetMsg sets the Msg field's value.
func (s *DetailForInstallAgentClientOutput) SetMsg(v string) *DetailForInstallAgentClientOutput {
	s.Msg = &v
	return s
}

type InstallAgentClientInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	AgentIDs []*string `type:"list" json:",omitempty"`
}

// String returns the string representation
func (s InstallAgentClientInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s InstallAgentClientInput) GoString() string {
	return s.String()
}

// SetAgentIDs sets the AgentIDs field's value.
func (s *InstallAgentClientInput) SetAgentIDs(v []*string) *InstallAgentClientInput {
	s.AgentIDs = v
	return s
}

type InstallAgentClientOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	Detail []*DetailForInstallAgentClientOutput `type:"list" json:",omitempty"`

	FailureCount *int32 `type:"int32" json:",omitempty"`

	SuccessCount *int32 `type:"int32" json:",omitempty"`
}

// String returns the string representation
func (s InstallAgentClientOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s InstallAgentClientOutput) GoString() string {
	return s.String()
}

// SetDetail sets the Detail field's value.
func (s *InstallAgentClientOutput) SetDetail(v []*DetailForInstallAgentClientOutput) *InstallAgentClientOutput {
	s.Detail = v
	return s
}

// SetFailureCount sets the FailureCount field's value.
func (s *InstallAgentClientOutput) SetFailureCount(v int32) *InstallAgentClientOutput {
	s.FailureCount = &v
	return s
}

// SetSuccessCount sets the SuccessCount field's value.
func (s *InstallAgentClientOutput) SetSuccessCount(v int32) *InstallAgentClientOutput {
	s.SuccessCount = &v
	return s
}

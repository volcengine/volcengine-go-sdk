// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package redis

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opSwitchOverCommon = "SwitchOver"

// SwitchOverCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the SwitchOverCommon operation. The "output" return
// value will be populated with the SwitchOverCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned SwitchOverCommon Request to send the API call to the service.
// the "output" return value is not valid until after SwitchOverCommon Send returns without error.
//
// See SwitchOverCommon for more information on using the SwitchOverCommon
// API call, and error handling.
//
//    // Example sending a request using the SwitchOverCommonRequest method.
//    req, resp := client.SwitchOverCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *REDIS) SwitchOverCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opSwitchOverCommon,
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

// SwitchOverCommon API operation for REDIS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for REDIS's
// API operation SwitchOverCommon for usage and error information.
func (c *REDIS) SwitchOverCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.SwitchOverCommonRequest(input)
	return out, req.Send()
}

// SwitchOverCommonWithContext is the same as SwitchOverCommon with the addition of
// the ability to pass a context and additional request options.
//
// See SwitchOverCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *REDIS) SwitchOverCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.SwitchOverCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opSwitchOver = "SwitchOver"

// SwitchOverRequest generates a "volcengine/request.Request" representing the
// client's request for the SwitchOver operation. The "output" return
// value will be populated with the SwitchOverCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned SwitchOverCommon Request to send the API call to the service.
// the "output" return value is not valid until after SwitchOverCommon Send returns without error.
//
// See SwitchOver for more information on using the SwitchOver
// API call, and error handling.
//
//    // Example sending a request using the SwitchOverRequest method.
//    req, resp := client.SwitchOverRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *REDIS) SwitchOverRequest(input *SwitchOverInput) (req *request.Request, output *SwitchOverOutput) {
	op := &request.Operation{
		Name:       opSwitchOver,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &SwitchOverInput{}
	}

	output = &SwitchOverOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// SwitchOver API operation for REDIS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for REDIS's
// API operation SwitchOver for usage and error information.
func (c *REDIS) SwitchOver(input *SwitchOverInput) (*SwitchOverOutput, error) {
	req, out := c.SwitchOverRequest(input)
	return out, req.Send()
}

// SwitchOverWithContext is the same as SwitchOver with the addition of
// the ability to pass a context and additional request options.
//
// See SwitchOver for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *REDIS) SwitchOverWithContext(ctx volcengine.Context, input *SwitchOverInput, opts ...request.Option) (*SwitchOverOutput, error) {
	req, out := c.SwitchOverRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type SwitchOverInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	ClientToken *string `type:"string" json:",omitempty"`

	// InstanceId is a required field
	InstanceId *string `type:"string" json:",omitempty" required:"true"`

	// TargetPrimaryNodeId is a required field
	TargetPrimaryNodeId *string `type:"string" json:",omitempty" required:"true"`
}

// String returns the string representation
func (s SwitchOverInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s SwitchOverInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *SwitchOverInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "SwitchOverInput"}
	if s.InstanceId == nil {
		invalidParams.Add(request.NewErrParamRequired("InstanceId"))
	}
	if s.TargetPrimaryNodeId == nil {
		invalidParams.Add(request.NewErrParamRequired("TargetPrimaryNodeId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetClientToken sets the ClientToken field's value.
func (s *SwitchOverInput) SetClientToken(v string) *SwitchOverInput {
	s.ClientToken = &v
	return s
}

// SetInstanceId sets the InstanceId field's value.
func (s *SwitchOverInput) SetInstanceId(v string) *SwitchOverInput {
	s.InstanceId = &v
	return s
}

// SetTargetPrimaryNodeId sets the TargetPrimaryNodeId field's value.
func (s *SwitchOverInput) SetTargetPrimaryNodeId(v string) *SwitchOverInput {
	s.TargetPrimaryNodeId = &v
	return s
}

type SwitchOverOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s SwitchOverOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s SwitchOverOutput) GoString() string {
	return s.String()
}

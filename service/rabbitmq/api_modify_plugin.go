// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package rabbitmq

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opModifyPluginCommon = "ModifyPlugin"

// ModifyPluginCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ModifyPluginCommon operation. The "output" return
// value will be populated with the ModifyPluginCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyPluginCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyPluginCommon Send returns without error.
//
// See ModifyPluginCommon for more information on using the ModifyPluginCommon
// API call, and error handling.
//
//    // Example sending a request using the ModifyPluginCommonRequest method.
//    req, resp := client.ModifyPluginCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *RABBITMQ) ModifyPluginCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opModifyPluginCommon,
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

// ModifyPluginCommon API operation for RABBITMQ.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RABBITMQ's
// API operation ModifyPluginCommon for usage and error information.
func (c *RABBITMQ) ModifyPluginCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ModifyPluginCommonRequest(input)
	return out, req.Send()
}

// ModifyPluginCommonWithContext is the same as ModifyPluginCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyPluginCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RABBITMQ) ModifyPluginCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ModifyPluginCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opModifyPlugin = "ModifyPlugin"

// ModifyPluginRequest generates a "volcengine/request.Request" representing the
// client's request for the ModifyPlugin operation. The "output" return
// value will be populated with the ModifyPluginCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyPluginCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyPluginCommon Send returns without error.
//
// See ModifyPlugin for more information on using the ModifyPlugin
// API call, and error handling.
//
//    // Example sending a request using the ModifyPluginRequest method.
//    req, resp := client.ModifyPluginRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *RABBITMQ) ModifyPluginRequest(input *ModifyPluginInput) (req *request.Request, output *ModifyPluginOutput) {
	op := &request.Operation{
		Name:       opModifyPlugin,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ModifyPluginInput{}
	}

	output = &ModifyPluginOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// ModifyPlugin API operation for RABBITMQ.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RABBITMQ's
// API operation ModifyPlugin for usage and error information.
func (c *RABBITMQ) ModifyPlugin(input *ModifyPluginInput) (*ModifyPluginOutput, error) {
	req, out := c.ModifyPluginRequest(input)
	return out, req.Send()
}

// ModifyPluginWithContext is the same as ModifyPlugin with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyPlugin for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RABBITMQ) ModifyPluginWithContext(ctx volcengine.Context, input *ModifyPluginInput, opts ...request.Option) (*ModifyPluginOutput, error) {
	req, out := c.ModifyPluginRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ModifyPluginInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	// InstanceId is a required field
	InstanceId *string `type:"string" json:",omitempty" required:"true"`

	Plugins []*PluginForModifyPluginInput `type:"list" json:",omitempty"`
}

// String returns the string representation
func (s ModifyPluginInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyPluginInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ModifyPluginInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ModifyPluginInput"}
	if s.InstanceId == nil {
		invalidParams.Add(request.NewErrParamRequired("InstanceId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetInstanceId sets the InstanceId field's value.
func (s *ModifyPluginInput) SetInstanceId(v string) *ModifyPluginInput {
	s.InstanceId = &v
	return s
}

// SetPlugins sets the Plugins field's value.
func (s *ModifyPluginInput) SetPlugins(v []*PluginForModifyPluginInput) *ModifyPluginInput {
	s.Plugins = v
	return s
}

type ModifyPluginOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s ModifyPluginOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyPluginOutput) GoString() string {
	return s.String()
}

type PluginForModifyPluginInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Enabled *bool `type:"boolean" json:",omitempty"`

	PluginName *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s PluginForModifyPluginInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s PluginForModifyPluginInput) GoString() string {
	return s.String()
}

// SetEnabled sets the Enabled field's value.
func (s *PluginForModifyPluginInput) SetEnabled(v bool) *PluginForModifyPluginInput {
	s.Enabled = &v
	return s
}

// SetPluginName sets the PluginName field's value.
func (s *PluginForModifyPluginInput) SetPluginName(v string) *PluginForModifyPluginInput {
	s.PluginName = &v
	return s
}
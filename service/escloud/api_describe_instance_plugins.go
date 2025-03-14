// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package escloud

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribeInstancePluginsCommon = "DescribeInstancePlugins"

// DescribeInstancePluginsCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeInstancePluginsCommon operation. The "output" return
// value will be populated with the DescribeInstancePluginsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeInstancePluginsCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeInstancePluginsCommon Send returns without error.
//
// See DescribeInstancePluginsCommon for more information on using the DescribeInstancePluginsCommon
// API call, and error handling.
//
//    // Example sending a request using the DescribeInstancePluginsCommonRequest method.
//    req, resp := client.DescribeInstancePluginsCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ESCLOUD) DescribeInstancePluginsCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeInstancePluginsCommon,
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

// DescribeInstancePluginsCommon API operation for ESCLOUD.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ESCLOUD's
// API operation DescribeInstancePluginsCommon for usage and error information.
func (c *ESCLOUD) DescribeInstancePluginsCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeInstancePluginsCommonRequest(input)
	return out, req.Send()
}

// DescribeInstancePluginsCommonWithContext is the same as DescribeInstancePluginsCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeInstancePluginsCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ESCLOUD) DescribeInstancePluginsCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeInstancePluginsCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeInstancePlugins = "DescribeInstancePlugins"

// DescribeInstancePluginsRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeInstancePlugins operation. The "output" return
// value will be populated with the DescribeInstancePluginsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeInstancePluginsCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeInstancePluginsCommon Send returns without error.
//
// See DescribeInstancePlugins for more information on using the DescribeInstancePlugins
// API call, and error handling.
//
//    // Example sending a request using the DescribeInstancePluginsRequest method.
//    req, resp := client.DescribeInstancePluginsRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ESCLOUD) DescribeInstancePluginsRequest(input *DescribeInstancePluginsInput) (req *request.Request, output *DescribeInstancePluginsOutput) {
	op := &request.Operation{
		Name:       opDescribeInstancePlugins,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeInstancePluginsInput{}
	}

	output = &DescribeInstancePluginsOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// DescribeInstancePlugins API operation for ESCLOUD.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ESCLOUD's
// API operation DescribeInstancePlugins for usage and error information.
func (c *ESCLOUD) DescribeInstancePlugins(input *DescribeInstancePluginsInput) (*DescribeInstancePluginsOutput, error) {
	req, out := c.DescribeInstancePluginsRequest(input)
	return out, req.Send()
}

// DescribeInstancePluginsWithContext is the same as DescribeInstancePlugins with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeInstancePlugins for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ESCLOUD) DescribeInstancePluginsWithContext(ctx volcengine.Context, input *DescribeInstancePluginsInput, opts ...request.Option) (*DescribeInstancePluginsOutput, error) {
	req, out := c.DescribeInstancePluginsRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DescribeInstancePluginsInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	// InstanceId is a required field
	InstanceId *string `type:"string" json:",omitempty" required:"true"`
}

// String returns the string representation
func (s DescribeInstancePluginsInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeInstancePluginsInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeInstancePluginsInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DescribeInstancePluginsInput"}
	if s.InstanceId == nil {
		invalidParams.Add(request.NewErrParamRequired("InstanceId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetInstanceId sets the InstanceId field's value.
func (s *DescribeInstancePluginsInput) SetInstanceId(v string) *DescribeInstancePluginsInput {
	s.InstanceId = &v
	return s
}

type DescribeInstancePluginsOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	InstancePlugins []*InstancePluginForDescribeInstancePluginsOutput `type:"list" json:",omitempty"`
}

// String returns the string representation
func (s DescribeInstancePluginsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeInstancePluginsOutput) GoString() string {
	return s.String()
}

// SetInstancePlugins sets the InstancePlugins field's value.
func (s *DescribeInstancePluginsOutput) SetInstancePlugins(v []*InstancePluginForDescribeInstancePluginsOutput) *DescribeInstancePluginsOutput {
	s.InstancePlugins = v
	return s
}

type InstancePluginForDescribeInstancePluginsOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Description *string `type:"string" json:",omitempty"`

	PluginName *string `type:"string" json:",omitempty"`

	Status *string `type:"string" json:",omitempty"`

	Version *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s InstancePluginForDescribeInstancePluginsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s InstancePluginForDescribeInstancePluginsOutput) GoString() string {
	return s.String()
}

// SetDescription sets the Description field's value.
func (s *InstancePluginForDescribeInstancePluginsOutput) SetDescription(v string) *InstancePluginForDescribeInstancePluginsOutput {
	s.Description = &v
	return s
}

// SetPluginName sets the PluginName field's value.
func (s *InstancePluginForDescribeInstancePluginsOutput) SetPluginName(v string) *InstancePluginForDescribeInstancePluginsOutput {
	s.PluginName = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *InstancePluginForDescribeInstancePluginsOutput) SetStatus(v string) *InstancePluginForDescribeInstancePluginsOutput {
	s.Status = &v
	return s
}

// SetVersion sets the Version field's value.
func (s *InstancePluginForDescribeInstancePluginsOutput) SetVersion(v string) *InstancePluginForDescribeInstancePluginsOutput {
	s.Version = &v
	return s
}

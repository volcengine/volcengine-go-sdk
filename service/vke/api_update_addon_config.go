// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package vke

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opUpdateAddonConfigCommon = "UpdateAddonConfig"

// UpdateAddonConfigCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the UpdateAddonConfigCommon operation. The "output" return
// value will be populated with the UpdateAddonConfigCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned UpdateAddonConfigCommon Request to send the API call to the service.
// the "output" return value is not valid until after UpdateAddonConfigCommon Send returns without error.
//
// See UpdateAddonConfigCommon for more information on using the UpdateAddonConfigCommon
// API call, and error handling.
//
//    // Example sending a request using the UpdateAddonConfigCommonRequest method.
//    req, resp := client.UpdateAddonConfigCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VKE) UpdateAddonConfigCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opUpdateAddonConfigCommon,
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

// UpdateAddonConfigCommon API operation for VKE.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VKE's
// API operation UpdateAddonConfigCommon for usage and error information.
func (c *VKE) UpdateAddonConfigCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.UpdateAddonConfigCommonRequest(input)
	return out, req.Send()
}

// UpdateAddonConfigCommonWithContext is the same as UpdateAddonConfigCommon with the addition of
// the ability to pass a context and additional request options.
//
// See UpdateAddonConfigCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VKE) UpdateAddonConfigCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.UpdateAddonConfigCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opUpdateAddonConfig = "UpdateAddonConfig"

// UpdateAddonConfigRequest generates a "volcengine/request.Request" representing the
// client's request for the UpdateAddonConfig operation. The "output" return
// value will be populated with the UpdateAddonConfigCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned UpdateAddonConfigCommon Request to send the API call to the service.
// the "output" return value is not valid until after UpdateAddonConfigCommon Send returns without error.
//
// See UpdateAddonConfig for more information on using the UpdateAddonConfig
// API call, and error handling.
//
//    // Example sending a request using the UpdateAddonConfigRequest method.
//    req, resp := client.UpdateAddonConfigRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VKE) UpdateAddonConfigRequest(input *UpdateAddonConfigInput) (req *request.Request, output *UpdateAddonConfigOutput) {
	op := &request.Operation{
		Name:       opUpdateAddonConfig,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateAddonConfigInput{}
	}

	output = &UpdateAddonConfigOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// UpdateAddonConfig API operation for VKE.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VKE's
// API operation UpdateAddonConfig for usage and error information.
func (c *VKE) UpdateAddonConfig(input *UpdateAddonConfigInput) (*UpdateAddonConfigOutput, error) {
	req, out := c.UpdateAddonConfigRequest(input)
	return out, req.Send()
}

// UpdateAddonConfigWithContext is the same as UpdateAddonConfig with the addition of
// the ability to pass a context and additional request options.
//
// See UpdateAddonConfig for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VKE) UpdateAddonConfigWithContext(ctx volcengine.Context, input *UpdateAddonConfigInput, opts ...request.Option) (*UpdateAddonConfigOutput, error) {
	req, out := c.UpdateAddonConfigRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type UpdateAddonConfigInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	ClientToken *string `type:"string" json:",omitempty"`

	// ClusterId is a required field
	ClusterId *string `type:"string" json:",omitempty" required:"true"`

	Config *string `type:"string" json:",omitempty"`

	// Name is a required field
	Name *string `type:"string" json:",omitempty" required:"true"`
}

// String returns the string representation
func (s UpdateAddonConfigInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s UpdateAddonConfigInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateAddonConfigInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "UpdateAddonConfigInput"}
	if s.ClusterId == nil {
		invalidParams.Add(request.NewErrParamRequired("ClusterId"))
	}
	if s.Name == nil {
		invalidParams.Add(request.NewErrParamRequired("Name"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetClientToken sets the ClientToken field's value.
func (s *UpdateAddonConfigInput) SetClientToken(v string) *UpdateAddonConfigInput {
	s.ClientToken = &v
	return s
}

// SetClusterId sets the ClusterId field's value.
func (s *UpdateAddonConfigInput) SetClusterId(v string) *UpdateAddonConfigInput {
	s.ClusterId = &v
	return s
}

// SetConfig sets the Config field's value.
func (s *UpdateAddonConfigInput) SetConfig(v string) *UpdateAddonConfigInput {
	s.Config = &v
	return s
}

// SetName sets the Name field's value.
func (s *UpdateAddonConfigInput) SetName(v string) *UpdateAddonConfigInput {
	s.Name = &v
	return s
}

type UpdateAddonConfigOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s UpdateAddonConfigOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s UpdateAddonConfigOutput) GoString() string {
	return s.String()
}

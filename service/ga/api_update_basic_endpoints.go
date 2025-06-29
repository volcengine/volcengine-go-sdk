// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package ga

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opUpdateBasicEndpointsCommon = "UpdateBasicEndpoints"

// UpdateBasicEndpointsCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the UpdateBasicEndpointsCommon operation. The "output" return
// value will be populated with the UpdateBasicEndpointsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned UpdateBasicEndpointsCommon Request to send the API call to the service.
// the "output" return value is not valid until after UpdateBasicEndpointsCommon Send returns without error.
//
// See UpdateBasicEndpointsCommon for more information on using the UpdateBasicEndpointsCommon
// API call, and error handling.
//
//    // Example sending a request using the UpdateBasicEndpointsCommonRequest method.
//    req, resp := client.UpdateBasicEndpointsCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *GA) UpdateBasicEndpointsCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opUpdateBasicEndpointsCommon,
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

// UpdateBasicEndpointsCommon API operation for GA.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for GA's
// API operation UpdateBasicEndpointsCommon for usage and error information.
func (c *GA) UpdateBasicEndpointsCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.UpdateBasicEndpointsCommonRequest(input)
	return out, req.Send()
}

// UpdateBasicEndpointsCommonWithContext is the same as UpdateBasicEndpointsCommon with the addition of
// the ability to pass a context and additional request options.
//
// See UpdateBasicEndpointsCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *GA) UpdateBasicEndpointsCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.UpdateBasicEndpointsCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opUpdateBasicEndpoints = "UpdateBasicEndpoints"

// UpdateBasicEndpointsRequest generates a "volcengine/request.Request" representing the
// client's request for the UpdateBasicEndpoints operation. The "output" return
// value will be populated with the UpdateBasicEndpointsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned UpdateBasicEndpointsCommon Request to send the API call to the service.
// the "output" return value is not valid until after UpdateBasicEndpointsCommon Send returns without error.
//
// See UpdateBasicEndpoints for more information on using the UpdateBasicEndpoints
// API call, and error handling.
//
//    // Example sending a request using the UpdateBasicEndpointsRequest method.
//    req, resp := client.UpdateBasicEndpointsRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *GA) UpdateBasicEndpointsRequest(input *UpdateBasicEndpointsInput) (req *request.Request, output *UpdateBasicEndpointsOutput) {
	op := &request.Operation{
		Name:       opUpdateBasicEndpoints,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateBasicEndpointsInput{}
	}

	output = &UpdateBasicEndpointsOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// UpdateBasicEndpoints API operation for GA.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for GA's
// API operation UpdateBasicEndpoints for usage and error information.
func (c *GA) UpdateBasicEndpoints(input *UpdateBasicEndpointsInput) (*UpdateBasicEndpointsOutput, error) {
	req, out := c.UpdateBasicEndpointsRequest(input)
	return out, req.Send()
}

// UpdateBasicEndpointsWithContext is the same as UpdateBasicEndpoints with the addition of
// the ability to pass a context and additional request options.
//
// See UpdateBasicEndpoints for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *GA) UpdateBasicEndpointsWithContext(ctx volcengine.Context, input *UpdateBasicEndpointsInput, opts ...request.Option) (*UpdateBasicEndpointsOutput, error) {
	req, out := c.UpdateBasicEndpointsRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type EndpointForUpdateBasicEndpointsInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	EndpointAddress *string `type:"string" json:",omitempty"`

	EndpointId *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s EndpointForUpdateBasicEndpointsInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s EndpointForUpdateBasicEndpointsInput) GoString() string {
	return s.String()
}

// SetEndpointAddress sets the EndpointAddress field's value.
func (s *EndpointForUpdateBasicEndpointsInput) SetEndpointAddress(v string) *EndpointForUpdateBasicEndpointsInput {
	s.EndpointAddress = &v
	return s
}

// SetEndpointId sets the EndpointId field's value.
func (s *EndpointForUpdateBasicEndpointsInput) SetEndpointId(v string) *EndpointForUpdateBasicEndpointsInput {
	s.EndpointId = &v
	return s
}

type UpdateBasicEndpointsInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	// AcceleratorId is a required field
	AcceleratorId *string `type:"string" json:",omitempty" required:"true"`

	// EndpointGroupId is a required field
	EndpointGroupId *string `type:"string" json:",omitempty" required:"true"`

	Endpoints []*EndpointForUpdateBasicEndpointsInput `type:"list" json:",omitempty"`
}

// String returns the string representation
func (s UpdateBasicEndpointsInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s UpdateBasicEndpointsInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateBasicEndpointsInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "UpdateBasicEndpointsInput"}
	if s.AcceleratorId == nil {
		invalidParams.Add(request.NewErrParamRequired("AcceleratorId"))
	}
	if s.EndpointGroupId == nil {
		invalidParams.Add(request.NewErrParamRequired("EndpointGroupId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetAcceleratorId sets the AcceleratorId field's value.
func (s *UpdateBasicEndpointsInput) SetAcceleratorId(v string) *UpdateBasicEndpointsInput {
	s.AcceleratorId = &v
	return s
}

// SetEndpointGroupId sets the EndpointGroupId field's value.
func (s *UpdateBasicEndpointsInput) SetEndpointGroupId(v string) *UpdateBasicEndpointsInput {
	s.EndpointGroupId = &v
	return s
}

// SetEndpoints sets the Endpoints field's value.
func (s *UpdateBasicEndpointsInput) SetEndpoints(v []*EndpointForUpdateBasicEndpointsInput) *UpdateBasicEndpointsInput {
	s.Endpoints = v
	return s
}

type UpdateBasicEndpointsOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s UpdateBasicEndpointsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s UpdateBasicEndpointsOutput) GoString() string {
	return s.String()
}

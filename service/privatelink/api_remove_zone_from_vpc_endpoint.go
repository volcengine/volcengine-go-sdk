// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package privatelink

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opRemoveZoneFromVpcEndpointCommon = "RemoveZoneFromVpcEndpoint"

// RemoveZoneFromVpcEndpointCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the RemoveZoneFromVpcEndpointCommon operation. The "output" return
// value will be populated with the RemoveZoneFromVpcEndpointCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned RemoveZoneFromVpcEndpointCommon Request to send the API call to the service.
// the "output" return value is not valid until after RemoveZoneFromVpcEndpointCommon Send returns without error.
//
// See RemoveZoneFromVpcEndpointCommon for more information on using the RemoveZoneFromVpcEndpointCommon
// API call, and error handling.
//
//    // Example sending a request using the RemoveZoneFromVpcEndpointCommonRequest method.
//    req, resp := client.RemoveZoneFromVpcEndpointCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *PRIVATELINK) RemoveZoneFromVpcEndpointCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opRemoveZoneFromVpcEndpointCommon,
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

// RemoveZoneFromVpcEndpointCommon API operation for PRIVATELINK.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for PRIVATELINK's
// API operation RemoveZoneFromVpcEndpointCommon for usage and error information.
func (c *PRIVATELINK) RemoveZoneFromVpcEndpointCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.RemoveZoneFromVpcEndpointCommonRequest(input)
	return out, req.Send()
}

// RemoveZoneFromVpcEndpointCommonWithContext is the same as RemoveZoneFromVpcEndpointCommon with the addition of
// the ability to pass a context and additional request options.
//
// See RemoveZoneFromVpcEndpointCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *PRIVATELINK) RemoveZoneFromVpcEndpointCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.RemoveZoneFromVpcEndpointCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opRemoveZoneFromVpcEndpoint = "RemoveZoneFromVpcEndpoint"

// RemoveZoneFromVpcEndpointRequest generates a "volcengine/request.Request" representing the
// client's request for the RemoveZoneFromVpcEndpoint operation. The "output" return
// value will be populated with the RemoveZoneFromVpcEndpointCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned RemoveZoneFromVpcEndpointCommon Request to send the API call to the service.
// the "output" return value is not valid until after RemoveZoneFromVpcEndpointCommon Send returns without error.
//
// See RemoveZoneFromVpcEndpoint for more information on using the RemoveZoneFromVpcEndpoint
// API call, and error handling.
//
//    // Example sending a request using the RemoveZoneFromVpcEndpointRequest method.
//    req, resp := client.RemoveZoneFromVpcEndpointRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *PRIVATELINK) RemoveZoneFromVpcEndpointRequest(input *RemoveZoneFromVpcEndpointInput) (req *request.Request, output *RemoveZoneFromVpcEndpointOutput) {
	op := &request.Operation{
		Name:       opRemoveZoneFromVpcEndpoint,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &RemoveZoneFromVpcEndpointInput{}
	}

	output = &RemoveZoneFromVpcEndpointOutput{}
	req = c.newRequest(op, input, output)

	return
}

// RemoveZoneFromVpcEndpoint API operation for PRIVATELINK.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for PRIVATELINK's
// API operation RemoveZoneFromVpcEndpoint for usage and error information.
func (c *PRIVATELINK) RemoveZoneFromVpcEndpoint(input *RemoveZoneFromVpcEndpointInput) (*RemoveZoneFromVpcEndpointOutput, error) {
	req, out := c.RemoveZoneFromVpcEndpointRequest(input)
	return out, req.Send()
}

// RemoveZoneFromVpcEndpointWithContext is the same as RemoveZoneFromVpcEndpoint with the addition of
// the ability to pass a context and additional request options.
//
// See RemoveZoneFromVpcEndpoint for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *PRIVATELINK) RemoveZoneFromVpcEndpointWithContext(ctx volcengine.Context, input *RemoveZoneFromVpcEndpointInput, opts ...request.Option) (*RemoveZoneFromVpcEndpointOutput, error) {
	req, out := c.RemoveZoneFromVpcEndpointRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type RemoveZoneFromVpcEndpointInput struct {
	_ struct{} `type:"structure"`

	// EndpointId is a required field
	EndpointId *string `type:"string" required:"true"`

	// ZoneId is a required field
	ZoneId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s RemoveZoneFromVpcEndpointInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s RemoveZoneFromVpcEndpointInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RemoveZoneFromVpcEndpointInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "RemoveZoneFromVpcEndpointInput"}
	if s.EndpointId == nil {
		invalidParams.Add(request.NewErrParamRequired("EndpointId"))
	}
	if s.ZoneId == nil {
		invalidParams.Add(request.NewErrParamRequired("ZoneId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetEndpointId sets the EndpointId field's value.
func (s *RemoveZoneFromVpcEndpointInput) SetEndpointId(v string) *RemoveZoneFromVpcEndpointInput {
	s.EndpointId = &v
	return s
}

// SetZoneId sets the ZoneId field's value.
func (s *RemoveZoneFromVpcEndpointInput) SetZoneId(v string) *RemoveZoneFromVpcEndpointInput {
	s.ZoneId = &v
	return s
}

type RemoveZoneFromVpcEndpointOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	RequestId *string `type:"string"`
}

// String returns the string representation
func (s RemoveZoneFromVpcEndpointOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s RemoveZoneFromVpcEndpointOutput) GoString() string {
	return s.String()
}

// SetRequestId sets the RequestId field's value.
func (s *RemoveZoneFromVpcEndpointOutput) SetRequestId(v string) *RemoveZoneFromVpcEndpointOutput {
	s.RequestId = &v
	return s
}

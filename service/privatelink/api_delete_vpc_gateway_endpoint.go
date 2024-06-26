// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package privatelink

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDeleteVpcGatewayEndpointCommon = "DeleteVpcGatewayEndpoint"

// DeleteVpcGatewayEndpointCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DeleteVpcGatewayEndpointCommon operation. The "output" return
// value will be populated with the DeleteVpcGatewayEndpointCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeleteVpcGatewayEndpointCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeleteVpcGatewayEndpointCommon Send returns without error.
//
// See DeleteVpcGatewayEndpointCommon for more information on using the DeleteVpcGatewayEndpointCommon
// API call, and error handling.
//
//    // Example sending a request using the DeleteVpcGatewayEndpointCommonRequest method.
//    req, resp := client.DeleteVpcGatewayEndpointCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *PRIVATELINK) DeleteVpcGatewayEndpointCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDeleteVpcGatewayEndpointCommon,
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

// DeleteVpcGatewayEndpointCommon API operation for PRIVATELINK.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for PRIVATELINK's
// API operation DeleteVpcGatewayEndpointCommon for usage and error information.
func (c *PRIVATELINK) DeleteVpcGatewayEndpointCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DeleteVpcGatewayEndpointCommonRequest(input)
	return out, req.Send()
}

// DeleteVpcGatewayEndpointCommonWithContext is the same as DeleteVpcGatewayEndpointCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteVpcGatewayEndpointCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *PRIVATELINK) DeleteVpcGatewayEndpointCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DeleteVpcGatewayEndpointCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDeleteVpcGatewayEndpoint = "DeleteVpcGatewayEndpoint"

// DeleteVpcGatewayEndpointRequest generates a "volcengine/request.Request" representing the
// client's request for the DeleteVpcGatewayEndpoint operation. The "output" return
// value will be populated with the DeleteVpcGatewayEndpointCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeleteVpcGatewayEndpointCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeleteVpcGatewayEndpointCommon Send returns without error.
//
// See DeleteVpcGatewayEndpoint for more information on using the DeleteVpcGatewayEndpoint
// API call, and error handling.
//
//    // Example sending a request using the DeleteVpcGatewayEndpointRequest method.
//    req, resp := client.DeleteVpcGatewayEndpointRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *PRIVATELINK) DeleteVpcGatewayEndpointRequest(input *DeleteVpcGatewayEndpointInput) (req *request.Request, output *DeleteVpcGatewayEndpointOutput) {
	op := &request.Operation{
		Name:       opDeleteVpcGatewayEndpoint,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteVpcGatewayEndpointInput{}
	}

	output = &DeleteVpcGatewayEndpointOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DeleteVpcGatewayEndpoint API operation for PRIVATELINK.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for PRIVATELINK's
// API operation DeleteVpcGatewayEndpoint for usage and error information.
func (c *PRIVATELINK) DeleteVpcGatewayEndpoint(input *DeleteVpcGatewayEndpointInput) (*DeleteVpcGatewayEndpointOutput, error) {
	req, out := c.DeleteVpcGatewayEndpointRequest(input)
	return out, req.Send()
}

// DeleteVpcGatewayEndpointWithContext is the same as DeleteVpcGatewayEndpoint with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteVpcGatewayEndpoint for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *PRIVATELINK) DeleteVpcGatewayEndpointWithContext(ctx volcengine.Context, input *DeleteVpcGatewayEndpointInput, opts ...request.Option) (*DeleteVpcGatewayEndpointOutput, error) {
	req, out := c.DeleteVpcGatewayEndpointRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DeleteVpcGatewayEndpointInput struct {
	_ struct{} `type:"structure"`

	// EndpointId is a required field
	EndpointId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteVpcGatewayEndpointInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteVpcGatewayEndpointInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteVpcGatewayEndpointInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DeleteVpcGatewayEndpointInput"}
	if s.EndpointId == nil {
		invalidParams.Add(request.NewErrParamRequired("EndpointId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetEndpointId sets the EndpointId field's value.
func (s *DeleteVpcGatewayEndpointInput) SetEndpointId(v string) *DeleteVpcGatewayEndpointInput {
	s.EndpointId = &v
	return s
}

type DeleteVpcGatewayEndpointOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	RequestId *string `type:"string"`
}

// String returns the string representation
func (s DeleteVpcGatewayEndpointOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteVpcGatewayEndpointOutput) GoString() string {
	return s.String()
}

// SetRequestId sets the RequestId field's value.
func (s *DeleteVpcGatewayEndpointOutput) SetRequestId(v string) *DeleteVpcGatewayEndpointOutput {
	s.RequestId = &v
	return s
}

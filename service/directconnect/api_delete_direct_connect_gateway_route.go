// Code generated by volcstack with private/model/cli/gen-api/main.go. DO NOT EDIT.

package directconnect

import (
	"github.com/volcengine/volcstack-go-sdk/volcstack"
	"github.com/volcengine/volcstack-go-sdk/volcstack/request"
	"github.com/volcengine/volcstack-go-sdk/volcstack/response"
	"github.com/volcengine/volcstack-go-sdk/volcstack/volcstackutil"
)

const opDeleteDirectConnectGatewayRouteCommon = "DeleteDirectConnectGatewayRoute"

// DeleteDirectConnectGatewayRouteCommonRequest generates a "volcstack/request.Request" representing the
// client's request for the DeleteDirectConnectGatewayRouteCommon operation. The "output" return
// value will be populated with the DeleteDirectConnectGatewayRouteCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeleteDirectConnectGatewayRouteCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeleteDirectConnectGatewayRouteCommon Send returns without error.
//
// See DeleteDirectConnectGatewayRouteCommon for more information on using the DeleteDirectConnectGatewayRouteCommon
// API call, and error handling.
//
//    // Example sending a request using the DeleteDirectConnectGatewayRouteCommonRequest method.
//    req, resp := client.DeleteDirectConnectGatewayRouteCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *DIRECTCONNECT) DeleteDirectConnectGatewayRouteCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDeleteDirectConnectGatewayRouteCommon,
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

// DeleteDirectConnectGatewayRouteCommon API operation for DIRECTCONNECT.
//
// Returns volcstackerr.Error for service API and SDK errors. Use runtime type assertions
// with volcstackerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCSTACK API reference guide for DIRECTCONNECT's
// API operation DeleteDirectConnectGatewayRouteCommon for usage and error information.
func (c *DIRECTCONNECT) DeleteDirectConnectGatewayRouteCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DeleteDirectConnectGatewayRouteCommonRequest(input)
	return out, req.Send()
}

// DeleteDirectConnectGatewayRouteCommonWithContext is the same as DeleteDirectConnectGatewayRouteCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteDirectConnectGatewayRouteCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *DIRECTCONNECT) DeleteDirectConnectGatewayRouteCommonWithContext(ctx volcstack.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DeleteDirectConnectGatewayRouteCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDeleteDirectConnectGatewayRoute = "DeleteDirectConnectGatewayRoute"

// DeleteDirectConnectGatewayRouteRequest generates a "volcstack/request.Request" representing the
// client's request for the DeleteDirectConnectGatewayRoute operation. The "output" return
// value will be populated with the DeleteDirectConnectGatewayRouteCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeleteDirectConnectGatewayRouteCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeleteDirectConnectGatewayRouteCommon Send returns without error.
//
// See DeleteDirectConnectGatewayRoute for more information on using the DeleteDirectConnectGatewayRoute
// API call, and error handling.
//
//    // Example sending a request using the DeleteDirectConnectGatewayRouteRequest method.
//    req, resp := client.DeleteDirectConnectGatewayRouteRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *DIRECTCONNECT) DeleteDirectConnectGatewayRouteRequest(input *DeleteDirectConnectGatewayRouteInput) (req *request.Request, output *DeleteDirectConnectGatewayRouteOutput) {
	op := &request.Operation{
		Name:       opDeleteDirectConnectGatewayRoute,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteDirectConnectGatewayRouteInput{}
	}

	output = &DeleteDirectConnectGatewayRouteOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DeleteDirectConnectGatewayRoute API operation for DIRECTCONNECT.
//
// Returns volcstackerr.Error for service API and SDK errors. Use runtime type assertions
// with volcstackerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCSTACK API reference guide for DIRECTCONNECT's
// API operation DeleteDirectConnectGatewayRoute for usage and error information.
func (c *DIRECTCONNECT) DeleteDirectConnectGatewayRoute(input *DeleteDirectConnectGatewayRouteInput) (*DeleteDirectConnectGatewayRouteOutput, error) {
	req, out := c.DeleteDirectConnectGatewayRouteRequest(input)
	return out, req.Send()
}

// DeleteDirectConnectGatewayRouteWithContext is the same as DeleteDirectConnectGatewayRoute with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteDirectConnectGatewayRoute for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *DIRECTCONNECT) DeleteDirectConnectGatewayRouteWithContext(ctx volcstack.Context, input *DeleteDirectConnectGatewayRouteInput, opts ...request.Option) (*DeleteDirectConnectGatewayRouteOutput, error) {
	req, out := c.DeleteDirectConnectGatewayRouteRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DeleteDirectConnectGatewayRouteInput struct {
	_ struct{} `type:"structure"`

	// DirectConnectGatewayRouteId is a required field
	DirectConnectGatewayRouteId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteDirectConnectGatewayRouteInput) String() string {
	return volcstackutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteDirectConnectGatewayRouteInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteDirectConnectGatewayRouteInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DeleteDirectConnectGatewayRouteInput"}
	if s.DirectConnectGatewayRouteId == nil {
		invalidParams.Add(request.NewErrParamRequired("DirectConnectGatewayRouteId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetDirectConnectGatewayRouteId sets the DirectConnectGatewayRouteId field's value.
func (s *DeleteDirectConnectGatewayRouteInput) SetDirectConnectGatewayRouteId(v string) *DeleteDirectConnectGatewayRouteInput {
	s.DirectConnectGatewayRouteId = &v
	return s
}

type DeleteDirectConnectGatewayRouteOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	RequestId *string `type:"string"`
}

// String returns the string representation
func (s DeleteDirectConnectGatewayRouteOutput) String() string {
	return volcstackutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteDirectConnectGatewayRouteOutput) GoString() string {
	return s.String()
}

// SetRequestId sets the RequestId field's value.
func (s *DeleteDirectConnectGatewayRouteOutput) SetRequestId(v string) *DeleteDirectConnectGatewayRouteOutput {
	s.RequestId = &v
	return s
}
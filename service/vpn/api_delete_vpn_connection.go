// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package vpn

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDeleteVpnConnectionCommon = "DeleteVpnConnection"

// DeleteVpnConnectionCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DeleteVpnConnectionCommon operation. The "output" return
// value will be populated with the DeleteVpnConnectionCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeleteVpnConnectionCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeleteVpnConnectionCommon Send returns without error.
//
// See DeleteVpnConnectionCommon for more information on using the DeleteVpnConnectionCommon
// API call, and error handling.
//
//    // Example sending a request using the DeleteVpnConnectionCommonRequest method.
//    req, resp := client.DeleteVpnConnectionCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VPN) DeleteVpnConnectionCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDeleteVpnConnectionCommon,
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

// DeleteVpnConnectionCommon API operation for VPN.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VPN's
// API operation DeleteVpnConnectionCommon for usage and error information.
func (c *VPN) DeleteVpnConnectionCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DeleteVpnConnectionCommonRequest(input)
	return out, req.Send()
}

// DeleteVpnConnectionCommonWithContext is the same as DeleteVpnConnectionCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteVpnConnectionCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPN) DeleteVpnConnectionCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DeleteVpnConnectionCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDeleteVpnConnection = "DeleteVpnConnection"

// DeleteVpnConnectionRequest generates a "volcengine/request.Request" representing the
// client's request for the DeleteVpnConnection operation. The "output" return
// value will be populated with the DeleteVpnConnectionCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeleteVpnConnectionCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeleteVpnConnectionCommon Send returns without error.
//
// See DeleteVpnConnection for more information on using the DeleteVpnConnection
// API call, and error handling.
//
//    // Example sending a request using the DeleteVpnConnectionRequest method.
//    req, resp := client.DeleteVpnConnectionRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VPN) DeleteVpnConnectionRequest(input *DeleteVpnConnectionInput) (req *request.Request, output *DeleteVpnConnectionOutput) {
	op := &request.Operation{
		Name:       opDeleteVpnConnection,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteVpnConnectionInput{}
	}

	output = &DeleteVpnConnectionOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DeleteVpnConnection API operation for VPN.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VPN's
// API operation DeleteVpnConnection for usage and error information.
func (c *VPN) DeleteVpnConnection(input *DeleteVpnConnectionInput) (*DeleteVpnConnectionOutput, error) {
	req, out := c.DeleteVpnConnectionRequest(input)
	return out, req.Send()
}

// DeleteVpnConnectionWithContext is the same as DeleteVpnConnection with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteVpnConnection for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPN) DeleteVpnConnectionWithContext(ctx volcengine.Context, input *DeleteVpnConnectionInput, opts ...request.Option) (*DeleteVpnConnectionOutput, error) {
	req, out := c.DeleteVpnConnectionRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DeleteVpnConnectionInput struct {
	_ struct{} `type:"structure"`

	// VpnConnectionId is a required field
	VpnConnectionId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteVpnConnectionInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteVpnConnectionInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteVpnConnectionInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DeleteVpnConnectionInput"}
	if s.VpnConnectionId == nil {
		invalidParams.Add(request.NewErrParamRequired("VpnConnectionId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetVpnConnectionId sets the VpnConnectionId field's value.
func (s *DeleteVpnConnectionInput) SetVpnConnectionId(v string) *DeleteVpnConnectionInput {
	s.VpnConnectionId = &v
	return s
}

type DeleteVpnConnectionOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	OrderId *string `type:"string"`

	RequestId *string `type:"string"`
}

// String returns the string representation
func (s DeleteVpnConnectionOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteVpnConnectionOutput) GoString() string {
	return s.String()
}

// SetOrderId sets the OrderId field's value.
func (s *DeleteVpnConnectionOutput) SetOrderId(v string) *DeleteVpnConnectionOutput {
	s.OrderId = &v
	return s
}

// SetRequestId sets the RequestId field's value.
func (s *DeleteVpnConnectionOutput) SetRequestId(v string) *DeleteVpnConnectionOutput {
	s.RequestId = &v
	return s
}

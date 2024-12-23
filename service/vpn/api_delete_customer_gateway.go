// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package vpn

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDeleteCustomerGatewayCommon = "DeleteCustomerGateway"

// DeleteCustomerGatewayCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DeleteCustomerGatewayCommon operation. The "output" return
// value will be populated with the DeleteCustomerGatewayCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeleteCustomerGatewayCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeleteCustomerGatewayCommon Send returns without error.
//
// See DeleteCustomerGatewayCommon for more information on using the DeleteCustomerGatewayCommon
// API call, and error handling.
//
//    // Example sending a request using the DeleteCustomerGatewayCommonRequest method.
//    req, resp := client.DeleteCustomerGatewayCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VPN) DeleteCustomerGatewayCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDeleteCustomerGatewayCommon,
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

// DeleteCustomerGatewayCommon API operation for VPN.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VPN's
// API operation DeleteCustomerGatewayCommon for usage and error information.
func (c *VPN) DeleteCustomerGatewayCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DeleteCustomerGatewayCommonRequest(input)
	return out, req.Send()
}

// DeleteCustomerGatewayCommonWithContext is the same as DeleteCustomerGatewayCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteCustomerGatewayCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPN) DeleteCustomerGatewayCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DeleteCustomerGatewayCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDeleteCustomerGateway = "DeleteCustomerGateway"

// DeleteCustomerGatewayRequest generates a "volcengine/request.Request" representing the
// client's request for the DeleteCustomerGateway operation. The "output" return
// value will be populated with the DeleteCustomerGatewayCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeleteCustomerGatewayCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeleteCustomerGatewayCommon Send returns without error.
//
// See DeleteCustomerGateway for more information on using the DeleteCustomerGateway
// API call, and error handling.
//
//    // Example sending a request using the DeleteCustomerGatewayRequest method.
//    req, resp := client.DeleteCustomerGatewayRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VPN) DeleteCustomerGatewayRequest(input *DeleteCustomerGatewayInput) (req *request.Request, output *DeleteCustomerGatewayOutput) {
	op := &request.Operation{
		Name:       opDeleteCustomerGateway,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteCustomerGatewayInput{}
	}

	output = &DeleteCustomerGatewayOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DeleteCustomerGateway API operation for VPN.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VPN's
// API operation DeleteCustomerGateway for usage and error information.
func (c *VPN) DeleteCustomerGateway(input *DeleteCustomerGatewayInput) (*DeleteCustomerGatewayOutput, error) {
	req, out := c.DeleteCustomerGatewayRequest(input)
	return out, req.Send()
}

// DeleteCustomerGatewayWithContext is the same as DeleteCustomerGateway with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteCustomerGateway for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPN) DeleteCustomerGatewayWithContext(ctx volcengine.Context, input *DeleteCustomerGatewayInput, opts ...request.Option) (*DeleteCustomerGatewayOutput, error) {
	req, out := c.DeleteCustomerGatewayRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DeleteCustomerGatewayInput struct {
	_ struct{} `type:"structure"`

	// CustomerGatewayId is a required field
	CustomerGatewayId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteCustomerGatewayInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteCustomerGatewayInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteCustomerGatewayInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DeleteCustomerGatewayInput"}
	if s.CustomerGatewayId == nil {
		invalidParams.Add(request.NewErrParamRequired("CustomerGatewayId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetCustomerGatewayId sets the CustomerGatewayId field's value.
func (s *DeleteCustomerGatewayInput) SetCustomerGatewayId(v string) *DeleteCustomerGatewayInput {
	s.CustomerGatewayId = &v
	return s
}

type DeleteCustomerGatewayOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	RequestId *string `type:"string"`
}

// String returns the string representation
func (s DeleteCustomerGatewayOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteCustomerGatewayOutput) GoString() string {
	return s.String()
}

// SetRequestId sets the RequestId field's value.
func (s *DeleteCustomerGatewayOutput) SetRequestId(v string) *DeleteCustomerGatewayOutput {
	s.RequestId = &v
	return s
}

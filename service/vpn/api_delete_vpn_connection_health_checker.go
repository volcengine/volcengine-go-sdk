// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package vpn

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDeleteVpnConnectionHealthCheckerCommon = "DeleteVpnConnectionHealthChecker"

// DeleteVpnConnectionHealthCheckerCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DeleteVpnConnectionHealthCheckerCommon operation. The "output" return
// value will be populated with the DeleteVpnConnectionHealthCheckerCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeleteVpnConnectionHealthCheckerCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeleteVpnConnectionHealthCheckerCommon Send returns without error.
//
// See DeleteVpnConnectionHealthCheckerCommon for more information on using the DeleteVpnConnectionHealthCheckerCommon
// API call, and error handling.
//
//    // Example sending a request using the DeleteVpnConnectionHealthCheckerCommonRequest method.
//    req, resp := client.DeleteVpnConnectionHealthCheckerCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VPN) DeleteVpnConnectionHealthCheckerCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDeleteVpnConnectionHealthCheckerCommon,
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

// DeleteVpnConnectionHealthCheckerCommon API operation for VPN.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VPN's
// API operation DeleteVpnConnectionHealthCheckerCommon for usage and error information.
func (c *VPN) DeleteVpnConnectionHealthCheckerCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DeleteVpnConnectionHealthCheckerCommonRequest(input)
	return out, req.Send()
}

// DeleteVpnConnectionHealthCheckerCommonWithContext is the same as DeleteVpnConnectionHealthCheckerCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteVpnConnectionHealthCheckerCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPN) DeleteVpnConnectionHealthCheckerCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DeleteVpnConnectionHealthCheckerCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDeleteVpnConnectionHealthChecker = "DeleteVpnConnectionHealthChecker"

// DeleteVpnConnectionHealthCheckerRequest generates a "volcengine/request.Request" representing the
// client's request for the DeleteVpnConnectionHealthChecker operation. The "output" return
// value will be populated with the DeleteVpnConnectionHealthCheckerCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeleteVpnConnectionHealthCheckerCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeleteVpnConnectionHealthCheckerCommon Send returns without error.
//
// See DeleteVpnConnectionHealthChecker for more information on using the DeleteVpnConnectionHealthChecker
// API call, and error handling.
//
//    // Example sending a request using the DeleteVpnConnectionHealthCheckerRequest method.
//    req, resp := client.DeleteVpnConnectionHealthCheckerRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VPN) DeleteVpnConnectionHealthCheckerRequest(input *DeleteVpnConnectionHealthCheckerInput) (req *request.Request, output *DeleteVpnConnectionHealthCheckerOutput) {
	op := &request.Operation{
		Name:       opDeleteVpnConnectionHealthChecker,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteVpnConnectionHealthCheckerInput{}
	}

	output = &DeleteVpnConnectionHealthCheckerOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DeleteVpnConnectionHealthChecker API operation for VPN.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VPN's
// API operation DeleteVpnConnectionHealthChecker for usage and error information.
func (c *VPN) DeleteVpnConnectionHealthChecker(input *DeleteVpnConnectionHealthCheckerInput) (*DeleteVpnConnectionHealthCheckerOutput, error) {
	req, out := c.DeleteVpnConnectionHealthCheckerRequest(input)
	return out, req.Send()
}

// DeleteVpnConnectionHealthCheckerWithContext is the same as DeleteVpnConnectionHealthChecker with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteVpnConnectionHealthChecker for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPN) DeleteVpnConnectionHealthCheckerWithContext(ctx volcengine.Context, input *DeleteVpnConnectionHealthCheckerInput, opts ...request.Option) (*DeleteVpnConnectionHealthCheckerOutput, error) {
	req, out := c.DeleteVpnConnectionHealthCheckerRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DeleteVpnConnectionHealthCheckerInput struct {
	_ struct{} `type:"structure"`

	// CheckerId is a required field
	CheckerId *string `type:"string" required:"true"`

	ClientToken *string `type:"string"`

	// VpnConnectionId is a required field
	VpnConnectionId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteVpnConnectionHealthCheckerInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteVpnConnectionHealthCheckerInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteVpnConnectionHealthCheckerInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DeleteVpnConnectionHealthCheckerInput"}
	if s.CheckerId == nil {
		invalidParams.Add(request.NewErrParamRequired("CheckerId"))
	}
	if s.VpnConnectionId == nil {
		invalidParams.Add(request.NewErrParamRequired("VpnConnectionId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetCheckerId sets the CheckerId field's value.
func (s *DeleteVpnConnectionHealthCheckerInput) SetCheckerId(v string) *DeleteVpnConnectionHealthCheckerInput {
	s.CheckerId = &v
	return s
}

// SetClientToken sets the ClientToken field's value.
func (s *DeleteVpnConnectionHealthCheckerInput) SetClientToken(v string) *DeleteVpnConnectionHealthCheckerInput {
	s.ClientToken = &v
	return s
}

// SetVpnConnectionId sets the VpnConnectionId field's value.
func (s *DeleteVpnConnectionHealthCheckerInput) SetVpnConnectionId(v string) *DeleteVpnConnectionHealthCheckerInput {
	s.VpnConnectionId = &v
	return s
}

type DeleteVpnConnectionHealthCheckerOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	RequestId *string `type:"string"`
}

// String returns the string representation
func (s DeleteVpnConnectionHealthCheckerOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteVpnConnectionHealthCheckerOutput) GoString() string {
	return s.String()
}

// SetRequestId sets the RequestId field's value.
func (s *DeleteVpnConnectionHealthCheckerOutput) SetRequestId(v string) *DeleteVpnConnectionHealthCheckerOutput {
	s.RequestId = &v
	return s
}
// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package vpn

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opCreateVpnGatewayRouteCommon = "CreateVpnGatewayRoute"

// CreateVpnGatewayRouteCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateVpnGatewayRouteCommon operation. The "output" return
// value will be populated with the CreateVpnGatewayRouteCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateVpnGatewayRouteCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateVpnGatewayRouteCommon Send returns without error.
//
// See CreateVpnGatewayRouteCommon for more information on using the CreateVpnGatewayRouteCommon
// API call, and error handling.
//
//	// Example sending a request using the CreateVpnGatewayRouteCommonRequest method.
//	req, resp := client.CreateVpnGatewayRouteCommonRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *VPN) CreateVpnGatewayRouteCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opCreateVpnGatewayRouteCommon,
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

// CreateVpnGatewayRouteCommon API operation for VPN.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VPN's
// API operation CreateVpnGatewayRouteCommon for usage and error information.
func (c *VPN) CreateVpnGatewayRouteCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.CreateVpnGatewayRouteCommonRequest(input)
	return out, req.Send()
}

// CreateVpnGatewayRouteCommonWithContext is the same as CreateVpnGatewayRouteCommon with the addition of
// the ability to pass a context and additional request options.
//
// See CreateVpnGatewayRouteCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPN) CreateVpnGatewayRouteCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.CreateVpnGatewayRouteCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opCreateVpnGatewayRoute = "CreateVpnGatewayRoute"

// CreateVpnGatewayRouteRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateVpnGatewayRoute operation. The "output" return
// value will be populated with the CreateVpnGatewayRouteCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateVpnGatewayRouteCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateVpnGatewayRouteCommon Send returns without error.
//
// See CreateVpnGatewayRoute for more information on using the CreateVpnGatewayRoute
// API call, and error handling.
//
//	// Example sending a request using the CreateVpnGatewayRouteRequest method.
//	req, resp := client.CreateVpnGatewayRouteRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *VPN) CreateVpnGatewayRouteRequest(input *CreateVpnGatewayRouteInput) (req *request.Request, output *CreateVpnGatewayRouteOutput) {
	op := &request.Operation{
		Name:       opCreateVpnGatewayRoute,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateVpnGatewayRouteInput{}
	}

	output = &CreateVpnGatewayRouteOutput{}
	req = c.newRequest(op, input, output)

	return
}

// CreateVpnGatewayRoute API operation for VPN.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VPN's
// API operation CreateVpnGatewayRoute for usage and error information.
func (c *VPN) CreateVpnGatewayRoute(input *CreateVpnGatewayRouteInput) (*CreateVpnGatewayRouteOutput, error) {
	req, out := c.CreateVpnGatewayRouteRequest(input)
	return out, req.Send()
}

// CreateVpnGatewayRouteWithContext is the same as CreateVpnGatewayRoute with the addition of
// the ability to pass a context and additional request options.
//
// See CreateVpnGatewayRoute for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPN) CreateVpnGatewayRouteWithContext(ctx volcengine.Context, input *CreateVpnGatewayRouteInput, opts ...request.Option) (*CreateVpnGatewayRouteOutput, error) {
	req, out := c.CreateVpnGatewayRouteRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type CreateVpnGatewayRouteInput struct {
	_ struct{} `type:"structure"`

	// DestinationCidrBlock is a required field
	DestinationCidrBlock *string `type:"string" required:"true"`

	// NextHopId is a required field
	NextHopId *string `type:"string" required:"true"`

	// VpnGatewayId is a required field
	VpnGatewayId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s CreateVpnGatewayRouteInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateVpnGatewayRouteInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateVpnGatewayRouteInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "CreateVpnGatewayRouteInput"}
	if s.DestinationCidrBlock == nil {
		invalidParams.Add(request.NewErrParamRequired("DestinationCidrBlock"))
	}
	if s.NextHopId == nil {
		invalidParams.Add(request.NewErrParamRequired("NextHopId"))
	}
	if s.VpnGatewayId == nil {
		invalidParams.Add(request.NewErrParamRequired("VpnGatewayId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetDestinationCidrBlock sets the DestinationCidrBlock field's value.
func (s *CreateVpnGatewayRouteInput) SetDestinationCidrBlock(v string) *CreateVpnGatewayRouteInput {
	s.DestinationCidrBlock = &v
	return s
}

// SetNextHopId sets the NextHopId field's value.
func (s *CreateVpnGatewayRouteInput) SetNextHopId(v string) *CreateVpnGatewayRouteInput {
	s.NextHopId = &v
	return s
}

// SetVpnGatewayId sets the VpnGatewayId field's value.
func (s *CreateVpnGatewayRouteInput) SetVpnGatewayId(v string) *CreateVpnGatewayRouteInput {
	s.VpnGatewayId = &v
	return s
}

type CreateVpnGatewayRouteOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	RequestId *string `type:"string"`

	VpnGatewayRouteId *string `type:"string"`
}

// String returns the string representation
func (s CreateVpnGatewayRouteOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateVpnGatewayRouteOutput) GoString() string {
	return s.String()
}

// SetRequestId sets the RequestId field's value.
func (s *CreateVpnGatewayRouteOutput) SetRequestId(v string) *CreateVpnGatewayRouteOutput {
	s.RequestId = &v
	return s
}

// SetVpnGatewayRouteId sets the VpnGatewayRouteId field's value.
func (s *CreateVpnGatewayRouteOutput) SetVpnGatewayRouteId(v string) *CreateVpnGatewayRouteOutput {
	s.VpnGatewayRouteId = &v
	return s
}

// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package vpn

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opCreateCustomerGatewayCommon = "CreateCustomerGateway"

// CreateCustomerGatewayCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateCustomerGatewayCommon operation. The "output" return
// value will be populated with the CreateCustomerGatewayCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateCustomerGatewayCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateCustomerGatewayCommon Send returns without error.
//
// See CreateCustomerGatewayCommon for more information on using the CreateCustomerGatewayCommon
// API call, and error handling.
//
//    // Example sending a request using the CreateCustomerGatewayCommonRequest method.
//    req, resp := client.CreateCustomerGatewayCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VPN) CreateCustomerGatewayCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opCreateCustomerGatewayCommon,
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

// CreateCustomerGatewayCommon API operation for VPN.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VPN's
// API operation CreateCustomerGatewayCommon for usage and error information.
func (c *VPN) CreateCustomerGatewayCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.CreateCustomerGatewayCommonRequest(input)
	return out, req.Send()
}

// CreateCustomerGatewayCommonWithContext is the same as CreateCustomerGatewayCommon with the addition of
// the ability to pass a context and additional request options.
//
// See CreateCustomerGatewayCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPN) CreateCustomerGatewayCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.CreateCustomerGatewayCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opCreateCustomerGateway = "CreateCustomerGateway"

// CreateCustomerGatewayRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateCustomerGateway operation. The "output" return
// value will be populated with the CreateCustomerGatewayCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateCustomerGatewayCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateCustomerGatewayCommon Send returns without error.
//
// See CreateCustomerGateway for more information on using the CreateCustomerGateway
// API call, and error handling.
//
//    // Example sending a request using the CreateCustomerGatewayRequest method.
//    req, resp := client.CreateCustomerGatewayRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VPN) CreateCustomerGatewayRequest(input *CreateCustomerGatewayInput) (req *request.Request, output *CreateCustomerGatewayOutput) {
	op := &request.Operation{
		Name:       opCreateCustomerGateway,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateCustomerGatewayInput{}
	}

	output = &CreateCustomerGatewayOutput{}
	req = c.newRequest(op, input, output)

	return
}

// CreateCustomerGateway API operation for VPN.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VPN's
// API operation CreateCustomerGateway for usage and error information.
func (c *VPN) CreateCustomerGateway(input *CreateCustomerGatewayInput) (*CreateCustomerGatewayOutput, error) {
	req, out := c.CreateCustomerGatewayRequest(input)
	return out, req.Send()
}

// CreateCustomerGatewayWithContext is the same as CreateCustomerGateway with the addition of
// the ability to pass a context and additional request options.
//
// See CreateCustomerGateway for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPN) CreateCustomerGatewayWithContext(ctx volcengine.Context, input *CreateCustomerGatewayInput, opts ...request.Option) (*CreateCustomerGatewayOutput, error) {
	req, out := c.CreateCustomerGatewayRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type CreateCustomerGatewayInput struct {
	_ struct{} `type:"structure"`

	ClientToken *string `type:"string"`

	CustomerGatewayName *string `min:"1" max:"128" type:"string"`

	Description *string `min:"1" max:"255" type:"string"`

	// IpAddress is a required field
	IpAddress *string `type:"string" required:"true"`
}

// String returns the string representation
func (s CreateCustomerGatewayInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateCustomerGatewayInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateCustomerGatewayInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "CreateCustomerGatewayInput"}
	if s.CustomerGatewayName != nil && len(*s.CustomerGatewayName) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("CustomerGatewayName", 1))
	}
	if s.CustomerGatewayName != nil && len(*s.CustomerGatewayName) > 128 {
		invalidParams.Add(request.NewErrParamMaxLen("CustomerGatewayName", 128, *s.CustomerGatewayName))
	}
	if s.Description != nil && len(*s.Description) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("Description", 1))
	}
	if s.Description != nil && len(*s.Description) > 255 {
		invalidParams.Add(request.NewErrParamMaxLen("Description", 255, *s.Description))
	}
	if s.IpAddress == nil {
		invalidParams.Add(request.NewErrParamRequired("IpAddress"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetClientToken sets the ClientToken field's value.
func (s *CreateCustomerGatewayInput) SetClientToken(v string) *CreateCustomerGatewayInput {
	s.ClientToken = &v
	return s
}

// SetCustomerGatewayName sets the CustomerGatewayName field's value.
func (s *CreateCustomerGatewayInput) SetCustomerGatewayName(v string) *CreateCustomerGatewayInput {
	s.CustomerGatewayName = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *CreateCustomerGatewayInput) SetDescription(v string) *CreateCustomerGatewayInput {
	s.Description = &v
	return s
}

// SetIpAddress sets the IpAddress field's value.
func (s *CreateCustomerGatewayInput) SetIpAddress(v string) *CreateCustomerGatewayInput {
	s.IpAddress = &v
	return s
}

type CreateCustomerGatewayOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	CustomerGatewayId *string `type:"string"`

	RequestId *string `type:"string"`
}

// String returns the string representation
func (s CreateCustomerGatewayOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateCustomerGatewayOutput) GoString() string {
	return s.String()
}

// SetCustomerGatewayId sets the CustomerGatewayId field's value.
func (s *CreateCustomerGatewayOutput) SetCustomerGatewayId(v string) *CreateCustomerGatewayOutput {
	s.CustomerGatewayId = &v
	return s
}

// SetRequestId sets the RequestId field's value.
func (s *CreateCustomerGatewayOutput) SetRequestId(v string) *CreateCustomerGatewayOutput {
	s.RequestId = &v
	return s
}
// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package vpc

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opModifyIpv6GatewayAttributeCommon = "ModifyIpv6GatewayAttribute"

// ModifyIpv6GatewayAttributeCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ModifyIpv6GatewayAttributeCommon operation. The "output" return
// value will be populated with the ModifyIpv6GatewayAttributeCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyIpv6GatewayAttributeCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyIpv6GatewayAttributeCommon Send returns without error.
//
// See ModifyIpv6GatewayAttributeCommon for more information on using the ModifyIpv6GatewayAttributeCommon
// API call, and error handling.
//
//    // Example sending a request using the ModifyIpv6GatewayAttributeCommonRequest method.
//    req, resp := client.ModifyIpv6GatewayAttributeCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VPC) ModifyIpv6GatewayAttributeCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opModifyIpv6GatewayAttributeCommon,
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

// ModifyIpv6GatewayAttributeCommon API operation for VPC.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VPC's
// API operation ModifyIpv6GatewayAttributeCommon for usage and error information.
func (c *VPC) ModifyIpv6GatewayAttributeCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ModifyIpv6GatewayAttributeCommonRequest(input)
	return out, req.Send()
}

// ModifyIpv6GatewayAttributeCommonWithContext is the same as ModifyIpv6GatewayAttributeCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyIpv6GatewayAttributeCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPC) ModifyIpv6GatewayAttributeCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ModifyIpv6GatewayAttributeCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opModifyIpv6GatewayAttribute = "ModifyIpv6GatewayAttribute"

// ModifyIpv6GatewayAttributeRequest generates a "volcengine/request.Request" representing the
// client's request for the ModifyIpv6GatewayAttribute operation. The "output" return
// value will be populated with the ModifyIpv6GatewayAttributeCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyIpv6GatewayAttributeCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyIpv6GatewayAttributeCommon Send returns without error.
//
// See ModifyIpv6GatewayAttribute for more information on using the ModifyIpv6GatewayAttribute
// API call, and error handling.
//
//    // Example sending a request using the ModifyIpv6GatewayAttributeRequest method.
//    req, resp := client.ModifyIpv6GatewayAttributeRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VPC) ModifyIpv6GatewayAttributeRequest(input *ModifyIpv6GatewayAttributeInput) (req *request.Request, output *ModifyIpv6GatewayAttributeOutput) {
	op := &request.Operation{
		Name:       opModifyIpv6GatewayAttribute,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ModifyIpv6GatewayAttributeInput{}
	}

	output = &ModifyIpv6GatewayAttributeOutput{}
	req = c.newRequest(op, input, output)

	return
}

// ModifyIpv6GatewayAttribute API operation for VPC.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VPC's
// API operation ModifyIpv6GatewayAttribute for usage and error information.
func (c *VPC) ModifyIpv6GatewayAttribute(input *ModifyIpv6GatewayAttributeInput) (*ModifyIpv6GatewayAttributeOutput, error) {
	req, out := c.ModifyIpv6GatewayAttributeRequest(input)
	return out, req.Send()
}

// ModifyIpv6GatewayAttributeWithContext is the same as ModifyIpv6GatewayAttribute with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyIpv6GatewayAttribute for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPC) ModifyIpv6GatewayAttributeWithContext(ctx volcengine.Context, input *ModifyIpv6GatewayAttributeInput, opts ...request.Option) (*ModifyIpv6GatewayAttributeOutput, error) {
	req, out := c.ModifyIpv6GatewayAttributeRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ModifyIpv6GatewayAttributeInput struct {
	_ struct{} `type:"structure"`

	Description *string `min:"1" max:"255" type:"string"`

	// Ipv6GatewayId is a required field
	Ipv6GatewayId *string `type:"string" required:"true"`

	Name *string `min:"1" max:"128" type:"string"`
}

// String returns the string representation
func (s ModifyIpv6GatewayAttributeInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyIpv6GatewayAttributeInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ModifyIpv6GatewayAttributeInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ModifyIpv6GatewayAttributeInput"}
	if s.Description != nil && len(*s.Description) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("Description", 1))
	}
	if s.Description != nil && len(*s.Description) > 255 {
		invalidParams.Add(request.NewErrParamMaxLen("Description", 255, *s.Description))
	}
	if s.Ipv6GatewayId == nil {
		invalidParams.Add(request.NewErrParamRequired("Ipv6GatewayId"))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("Name", 1))
	}
	if s.Name != nil && len(*s.Name) > 128 {
		invalidParams.Add(request.NewErrParamMaxLen("Name", 128, *s.Name))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetDescription sets the Description field's value.
func (s *ModifyIpv6GatewayAttributeInput) SetDescription(v string) *ModifyIpv6GatewayAttributeInput {
	s.Description = &v
	return s
}

// SetIpv6GatewayId sets the Ipv6GatewayId field's value.
func (s *ModifyIpv6GatewayAttributeInput) SetIpv6GatewayId(v string) *ModifyIpv6GatewayAttributeInput {
	s.Ipv6GatewayId = &v
	return s
}

// SetName sets the Name field's value.
func (s *ModifyIpv6GatewayAttributeInput) SetName(v string) *ModifyIpv6GatewayAttributeInput {
	s.Name = &v
	return s
}

type ModifyIpv6GatewayAttributeOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	RequestId *string `type:"string"`
}

// String returns the string representation
func (s ModifyIpv6GatewayAttributeOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyIpv6GatewayAttributeOutput) GoString() string {
	return s.String()
}

// SetRequestId sets the RequestId field's value.
func (s *ModifyIpv6GatewayAttributeOutput) SetRequestId(v string) *ModifyIpv6GatewayAttributeOutput {
	s.RequestId = &v
	return s
}

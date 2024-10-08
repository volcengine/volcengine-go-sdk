// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package privatelink

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opModifyVpcEndpointServiceAttributesCommon = "ModifyVpcEndpointServiceAttributes"

// ModifyVpcEndpointServiceAttributesCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ModifyVpcEndpointServiceAttributesCommon operation. The "output" return
// value will be populated with the ModifyVpcEndpointServiceAttributesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyVpcEndpointServiceAttributesCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyVpcEndpointServiceAttributesCommon Send returns without error.
//
// See ModifyVpcEndpointServiceAttributesCommon for more information on using the ModifyVpcEndpointServiceAttributesCommon
// API call, and error handling.
//
//    // Example sending a request using the ModifyVpcEndpointServiceAttributesCommonRequest method.
//    req, resp := client.ModifyVpcEndpointServiceAttributesCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *PRIVATELINK) ModifyVpcEndpointServiceAttributesCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opModifyVpcEndpointServiceAttributesCommon,
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

// ModifyVpcEndpointServiceAttributesCommon API operation for PRIVATELINK.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for PRIVATELINK's
// API operation ModifyVpcEndpointServiceAttributesCommon for usage and error information.
func (c *PRIVATELINK) ModifyVpcEndpointServiceAttributesCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ModifyVpcEndpointServiceAttributesCommonRequest(input)
	return out, req.Send()
}

// ModifyVpcEndpointServiceAttributesCommonWithContext is the same as ModifyVpcEndpointServiceAttributesCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyVpcEndpointServiceAttributesCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *PRIVATELINK) ModifyVpcEndpointServiceAttributesCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ModifyVpcEndpointServiceAttributesCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opModifyVpcEndpointServiceAttributes = "ModifyVpcEndpointServiceAttributes"

// ModifyVpcEndpointServiceAttributesRequest generates a "volcengine/request.Request" representing the
// client's request for the ModifyVpcEndpointServiceAttributes operation. The "output" return
// value will be populated with the ModifyVpcEndpointServiceAttributesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyVpcEndpointServiceAttributesCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyVpcEndpointServiceAttributesCommon Send returns without error.
//
// See ModifyVpcEndpointServiceAttributes for more information on using the ModifyVpcEndpointServiceAttributes
// API call, and error handling.
//
//    // Example sending a request using the ModifyVpcEndpointServiceAttributesRequest method.
//    req, resp := client.ModifyVpcEndpointServiceAttributesRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *PRIVATELINK) ModifyVpcEndpointServiceAttributesRequest(input *ModifyVpcEndpointServiceAttributesInput) (req *request.Request, output *ModifyVpcEndpointServiceAttributesOutput) {
	op := &request.Operation{
		Name:       opModifyVpcEndpointServiceAttributes,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ModifyVpcEndpointServiceAttributesInput{}
	}

	output = &ModifyVpcEndpointServiceAttributesOutput{}
	req = c.newRequest(op, input, output)

	return
}

// ModifyVpcEndpointServiceAttributes API operation for PRIVATELINK.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for PRIVATELINK's
// API operation ModifyVpcEndpointServiceAttributes for usage and error information.
func (c *PRIVATELINK) ModifyVpcEndpointServiceAttributes(input *ModifyVpcEndpointServiceAttributesInput) (*ModifyVpcEndpointServiceAttributesOutput, error) {
	req, out := c.ModifyVpcEndpointServiceAttributesRequest(input)
	return out, req.Send()
}

// ModifyVpcEndpointServiceAttributesWithContext is the same as ModifyVpcEndpointServiceAttributes with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyVpcEndpointServiceAttributes for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *PRIVATELINK) ModifyVpcEndpointServiceAttributesWithContext(ctx volcengine.Context, input *ModifyVpcEndpointServiceAttributesInput, opts ...request.Option) (*ModifyVpcEndpointServiceAttributesOutput, error) {
	req, out := c.ModifyVpcEndpointServiceAttributesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ModifyVpcEndpointServiceAttributesInput struct {
	_ struct{} `type:"structure"`

	AutoAcceptEnabled *bool `type:"boolean"`

	Description *string `type:"string"`

	IpAddressVersions []*string `type:"list"`

	PrivateDNSEnabled *string `type:"string"`

	PrivateDNSName *string `type:"string"`

	PrivateDNSType *string `type:"string"`

	// ServiceId is a required field
	ServiceId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s ModifyVpcEndpointServiceAttributesInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyVpcEndpointServiceAttributesInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ModifyVpcEndpointServiceAttributesInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ModifyVpcEndpointServiceAttributesInput"}
	if s.ServiceId == nil {
		invalidParams.Add(request.NewErrParamRequired("ServiceId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetAutoAcceptEnabled sets the AutoAcceptEnabled field's value.
func (s *ModifyVpcEndpointServiceAttributesInput) SetAutoAcceptEnabled(v bool) *ModifyVpcEndpointServiceAttributesInput {
	s.AutoAcceptEnabled = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *ModifyVpcEndpointServiceAttributesInput) SetDescription(v string) *ModifyVpcEndpointServiceAttributesInput {
	s.Description = &v
	return s
}

// SetIpAddressVersions sets the IpAddressVersions field's value.
func (s *ModifyVpcEndpointServiceAttributesInput) SetIpAddressVersions(v []*string) *ModifyVpcEndpointServiceAttributesInput {
	s.IpAddressVersions = v
	return s
}

// SetPrivateDNSEnabled sets the PrivateDNSEnabled field's value.
func (s *ModifyVpcEndpointServiceAttributesInput) SetPrivateDNSEnabled(v string) *ModifyVpcEndpointServiceAttributesInput {
	s.PrivateDNSEnabled = &v
	return s
}

// SetPrivateDNSName sets the PrivateDNSName field's value.
func (s *ModifyVpcEndpointServiceAttributesInput) SetPrivateDNSName(v string) *ModifyVpcEndpointServiceAttributesInput {
	s.PrivateDNSName = &v
	return s
}

// SetPrivateDNSType sets the PrivateDNSType field's value.
func (s *ModifyVpcEndpointServiceAttributesInput) SetPrivateDNSType(v string) *ModifyVpcEndpointServiceAttributesInput {
	s.PrivateDNSType = &v
	return s
}

// SetServiceId sets the ServiceId field's value.
func (s *ModifyVpcEndpointServiceAttributesInput) SetServiceId(v string) *ModifyVpcEndpointServiceAttributesInput {
	s.ServiceId = &v
	return s
}

type ModifyVpcEndpointServiceAttributesOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	RequestId *string `type:"string"`
}

// String returns the string representation
func (s ModifyVpcEndpointServiceAttributesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyVpcEndpointServiceAttributesOutput) GoString() string {
	return s.String()
}

// SetRequestId sets the RequestId field's value.
func (s *ModifyVpcEndpointServiceAttributesOutput) SetRequestId(v string) *ModifyVpcEndpointServiceAttributesOutput {
	s.RequestId = &v
	return s
}

// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package privatelink

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opModifyVpcEndpointAttributesCommon = "ModifyVpcEndpointAttributes"

// ModifyVpcEndpointAttributesCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ModifyVpcEndpointAttributesCommon operation. The "output" return
// value will be populated with the ModifyVpcEndpointAttributesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyVpcEndpointAttributesCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyVpcEndpointAttributesCommon Send returns without error.
//
// See ModifyVpcEndpointAttributesCommon for more information on using the ModifyVpcEndpointAttributesCommon
// API call, and error handling.
//
//    // Example sending a request using the ModifyVpcEndpointAttributesCommonRequest method.
//    req, resp := client.ModifyVpcEndpointAttributesCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *PRIVATELINK) ModifyVpcEndpointAttributesCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opModifyVpcEndpointAttributesCommon,
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

// ModifyVpcEndpointAttributesCommon API operation for PRIVATELINK.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for PRIVATELINK's
// API operation ModifyVpcEndpointAttributesCommon for usage and error information.
func (c *PRIVATELINK) ModifyVpcEndpointAttributesCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ModifyVpcEndpointAttributesCommonRequest(input)
	return out, req.Send()
}

// ModifyVpcEndpointAttributesCommonWithContext is the same as ModifyVpcEndpointAttributesCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyVpcEndpointAttributesCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *PRIVATELINK) ModifyVpcEndpointAttributesCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ModifyVpcEndpointAttributesCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opModifyVpcEndpointAttributes = "ModifyVpcEndpointAttributes"

// ModifyVpcEndpointAttributesRequest generates a "volcengine/request.Request" representing the
// client's request for the ModifyVpcEndpointAttributes operation. The "output" return
// value will be populated with the ModifyVpcEndpointAttributesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyVpcEndpointAttributesCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyVpcEndpointAttributesCommon Send returns without error.
//
// See ModifyVpcEndpointAttributes for more information on using the ModifyVpcEndpointAttributes
// API call, and error handling.
//
//    // Example sending a request using the ModifyVpcEndpointAttributesRequest method.
//    req, resp := client.ModifyVpcEndpointAttributesRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *PRIVATELINK) ModifyVpcEndpointAttributesRequest(input *ModifyVpcEndpointAttributesInput) (req *request.Request, output *ModifyVpcEndpointAttributesOutput) {
	op := &request.Operation{
		Name:       opModifyVpcEndpointAttributes,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ModifyVpcEndpointAttributesInput{}
	}

	output = &ModifyVpcEndpointAttributesOutput{}
	req = c.newRequest(op, input, output)

	return
}

// ModifyVpcEndpointAttributes API operation for PRIVATELINK.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for PRIVATELINK's
// API operation ModifyVpcEndpointAttributes for usage and error information.
func (c *PRIVATELINK) ModifyVpcEndpointAttributes(input *ModifyVpcEndpointAttributesInput) (*ModifyVpcEndpointAttributesOutput, error) {
	req, out := c.ModifyVpcEndpointAttributesRequest(input)
	return out, req.Send()
}

// ModifyVpcEndpointAttributesWithContext is the same as ModifyVpcEndpointAttributes with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyVpcEndpointAttributes for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *PRIVATELINK) ModifyVpcEndpointAttributesWithContext(ctx volcengine.Context, input *ModifyVpcEndpointAttributesInput, opts ...request.Option) (*ModifyVpcEndpointAttributesOutput, error) {
	req, out := c.ModifyVpcEndpointAttributesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ModifyVpcEndpointAttributesInput struct {
	_ struct{} `type:"structure"`

	Description *string `type:"string"`

	// EndpointId is a required field
	EndpointId *string `type:"string" required:"true"`

	EndpointName *string `type:"string"`

	PrivateDNSEnabled *bool `type:"boolean"`
}

// String returns the string representation
func (s ModifyVpcEndpointAttributesInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyVpcEndpointAttributesInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ModifyVpcEndpointAttributesInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ModifyVpcEndpointAttributesInput"}
	if s.EndpointId == nil {
		invalidParams.Add(request.NewErrParamRequired("EndpointId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetDescription sets the Description field's value.
func (s *ModifyVpcEndpointAttributesInput) SetDescription(v string) *ModifyVpcEndpointAttributesInput {
	s.Description = &v
	return s
}

// SetEndpointId sets the EndpointId field's value.
func (s *ModifyVpcEndpointAttributesInput) SetEndpointId(v string) *ModifyVpcEndpointAttributesInput {
	s.EndpointId = &v
	return s
}

// SetEndpointName sets the EndpointName field's value.
func (s *ModifyVpcEndpointAttributesInput) SetEndpointName(v string) *ModifyVpcEndpointAttributesInput {
	s.EndpointName = &v
	return s
}

// SetPrivateDNSEnabled sets the PrivateDNSEnabled field's value.
func (s *ModifyVpcEndpointAttributesInput) SetPrivateDNSEnabled(v bool) *ModifyVpcEndpointAttributesInput {
	s.PrivateDNSEnabled = &v
	return s
}

type ModifyVpcEndpointAttributesOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	RequestId *string `type:"string"`
}

// String returns the string representation
func (s ModifyVpcEndpointAttributesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyVpcEndpointAttributesOutput) GoString() string {
	return s.String()
}

// SetRequestId sets the RequestId field's value.
func (s *ModifyVpcEndpointAttributesOutput) SetRequestId(v string) *ModifyVpcEndpointAttributesOutput {
	s.RequestId = &v
	return s
}
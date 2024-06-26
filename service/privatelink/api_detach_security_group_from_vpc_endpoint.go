// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package privatelink

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDetachSecurityGroupFromVpcEndpointCommon = "DetachSecurityGroupFromVpcEndpoint"

// DetachSecurityGroupFromVpcEndpointCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DetachSecurityGroupFromVpcEndpointCommon operation. The "output" return
// value will be populated with the DetachSecurityGroupFromVpcEndpointCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DetachSecurityGroupFromVpcEndpointCommon Request to send the API call to the service.
// the "output" return value is not valid until after DetachSecurityGroupFromVpcEndpointCommon Send returns without error.
//
// See DetachSecurityGroupFromVpcEndpointCommon for more information on using the DetachSecurityGroupFromVpcEndpointCommon
// API call, and error handling.
//
//    // Example sending a request using the DetachSecurityGroupFromVpcEndpointCommonRequest method.
//    req, resp := client.DetachSecurityGroupFromVpcEndpointCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *PRIVATELINK) DetachSecurityGroupFromVpcEndpointCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDetachSecurityGroupFromVpcEndpointCommon,
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

// DetachSecurityGroupFromVpcEndpointCommon API operation for PRIVATELINK.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for PRIVATELINK's
// API operation DetachSecurityGroupFromVpcEndpointCommon for usage and error information.
func (c *PRIVATELINK) DetachSecurityGroupFromVpcEndpointCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DetachSecurityGroupFromVpcEndpointCommonRequest(input)
	return out, req.Send()
}

// DetachSecurityGroupFromVpcEndpointCommonWithContext is the same as DetachSecurityGroupFromVpcEndpointCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DetachSecurityGroupFromVpcEndpointCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *PRIVATELINK) DetachSecurityGroupFromVpcEndpointCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DetachSecurityGroupFromVpcEndpointCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDetachSecurityGroupFromVpcEndpoint = "DetachSecurityGroupFromVpcEndpoint"

// DetachSecurityGroupFromVpcEndpointRequest generates a "volcengine/request.Request" representing the
// client's request for the DetachSecurityGroupFromVpcEndpoint operation. The "output" return
// value will be populated with the DetachSecurityGroupFromVpcEndpointCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DetachSecurityGroupFromVpcEndpointCommon Request to send the API call to the service.
// the "output" return value is not valid until after DetachSecurityGroupFromVpcEndpointCommon Send returns without error.
//
// See DetachSecurityGroupFromVpcEndpoint for more information on using the DetachSecurityGroupFromVpcEndpoint
// API call, and error handling.
//
//    // Example sending a request using the DetachSecurityGroupFromVpcEndpointRequest method.
//    req, resp := client.DetachSecurityGroupFromVpcEndpointRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *PRIVATELINK) DetachSecurityGroupFromVpcEndpointRequest(input *DetachSecurityGroupFromVpcEndpointInput) (req *request.Request, output *DetachSecurityGroupFromVpcEndpointOutput) {
	op := &request.Operation{
		Name:       opDetachSecurityGroupFromVpcEndpoint,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DetachSecurityGroupFromVpcEndpointInput{}
	}

	output = &DetachSecurityGroupFromVpcEndpointOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DetachSecurityGroupFromVpcEndpoint API operation for PRIVATELINK.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for PRIVATELINK's
// API operation DetachSecurityGroupFromVpcEndpoint for usage and error information.
func (c *PRIVATELINK) DetachSecurityGroupFromVpcEndpoint(input *DetachSecurityGroupFromVpcEndpointInput) (*DetachSecurityGroupFromVpcEndpointOutput, error) {
	req, out := c.DetachSecurityGroupFromVpcEndpointRequest(input)
	return out, req.Send()
}

// DetachSecurityGroupFromVpcEndpointWithContext is the same as DetachSecurityGroupFromVpcEndpoint with the addition of
// the ability to pass a context and additional request options.
//
// See DetachSecurityGroupFromVpcEndpoint for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *PRIVATELINK) DetachSecurityGroupFromVpcEndpointWithContext(ctx volcengine.Context, input *DetachSecurityGroupFromVpcEndpointInput, opts ...request.Option) (*DetachSecurityGroupFromVpcEndpointOutput, error) {
	req, out := c.DetachSecurityGroupFromVpcEndpointRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DetachSecurityGroupFromVpcEndpointInput struct {
	_ struct{} `type:"structure"`

	// EndpointId is a required field
	EndpointId *string `type:"string" required:"true"`

	// SecurityGroupId is a required field
	SecurityGroupId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DetachSecurityGroupFromVpcEndpointInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DetachSecurityGroupFromVpcEndpointInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DetachSecurityGroupFromVpcEndpointInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DetachSecurityGroupFromVpcEndpointInput"}
	if s.EndpointId == nil {
		invalidParams.Add(request.NewErrParamRequired("EndpointId"))
	}
	if s.SecurityGroupId == nil {
		invalidParams.Add(request.NewErrParamRequired("SecurityGroupId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetEndpointId sets the EndpointId field's value.
func (s *DetachSecurityGroupFromVpcEndpointInput) SetEndpointId(v string) *DetachSecurityGroupFromVpcEndpointInput {
	s.EndpointId = &v
	return s
}

// SetSecurityGroupId sets the SecurityGroupId field's value.
func (s *DetachSecurityGroupFromVpcEndpointInput) SetSecurityGroupId(v string) *DetachSecurityGroupFromVpcEndpointInput {
	s.SecurityGroupId = &v
	return s
}

type DetachSecurityGroupFromVpcEndpointOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	RequestId *string `type:"string"`
}

// String returns the string representation
func (s DetachSecurityGroupFromVpcEndpointOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DetachSecurityGroupFromVpcEndpointOutput) GoString() string {
	return s.String()
}

// SetRequestId sets the RequestId field's value.
func (s *DetachSecurityGroupFromVpcEndpointOutput) SetRequestId(v string) *DetachSecurityGroupFromVpcEndpointOutput {
	s.RequestId = &v
	return s
}

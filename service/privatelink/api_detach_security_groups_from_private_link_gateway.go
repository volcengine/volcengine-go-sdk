// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package privatelink

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDetachSecurityGroupsFromPrivateLinkGatewayCommon = "DetachSecurityGroupsFromPrivateLinkGateway"

// DetachSecurityGroupsFromPrivateLinkGatewayCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DetachSecurityGroupsFromPrivateLinkGatewayCommon operation. The "output" return
// value will be populated with the DetachSecurityGroupsFromPrivateLinkGatewayCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DetachSecurityGroupsFromPrivateLinkGatewayCommon Request to send the API call to the service.
// the "output" return value is not valid until after DetachSecurityGroupsFromPrivateLinkGatewayCommon Send returns without error.
//
// See DetachSecurityGroupsFromPrivateLinkGatewayCommon for more information on using the DetachSecurityGroupsFromPrivateLinkGatewayCommon
// API call, and error handling.
//
//    // Example sending a request using the DetachSecurityGroupsFromPrivateLinkGatewayCommonRequest method.
//    req, resp := client.DetachSecurityGroupsFromPrivateLinkGatewayCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *PRIVATELINK) DetachSecurityGroupsFromPrivateLinkGatewayCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDetachSecurityGroupsFromPrivateLinkGatewayCommon,
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

// DetachSecurityGroupsFromPrivateLinkGatewayCommon API operation for PRIVATELINK.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for PRIVATELINK's
// API operation DetachSecurityGroupsFromPrivateLinkGatewayCommon for usage and error information.
func (c *PRIVATELINK) DetachSecurityGroupsFromPrivateLinkGatewayCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DetachSecurityGroupsFromPrivateLinkGatewayCommonRequest(input)
	return out, req.Send()
}

// DetachSecurityGroupsFromPrivateLinkGatewayCommonWithContext is the same as DetachSecurityGroupsFromPrivateLinkGatewayCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DetachSecurityGroupsFromPrivateLinkGatewayCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *PRIVATELINK) DetachSecurityGroupsFromPrivateLinkGatewayCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DetachSecurityGroupsFromPrivateLinkGatewayCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDetachSecurityGroupsFromPrivateLinkGateway = "DetachSecurityGroupsFromPrivateLinkGateway"

// DetachSecurityGroupsFromPrivateLinkGatewayRequest generates a "volcengine/request.Request" representing the
// client's request for the DetachSecurityGroupsFromPrivateLinkGateway operation. The "output" return
// value will be populated with the DetachSecurityGroupsFromPrivateLinkGatewayCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DetachSecurityGroupsFromPrivateLinkGatewayCommon Request to send the API call to the service.
// the "output" return value is not valid until after DetachSecurityGroupsFromPrivateLinkGatewayCommon Send returns without error.
//
// See DetachSecurityGroupsFromPrivateLinkGateway for more information on using the DetachSecurityGroupsFromPrivateLinkGateway
// API call, and error handling.
//
//    // Example sending a request using the DetachSecurityGroupsFromPrivateLinkGatewayRequest method.
//    req, resp := client.DetachSecurityGroupsFromPrivateLinkGatewayRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *PRIVATELINK) DetachSecurityGroupsFromPrivateLinkGatewayRequest(input *DetachSecurityGroupsFromPrivateLinkGatewayInput) (req *request.Request, output *DetachSecurityGroupsFromPrivateLinkGatewayOutput) {
	op := &request.Operation{
		Name:       opDetachSecurityGroupsFromPrivateLinkGateway,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DetachSecurityGroupsFromPrivateLinkGatewayInput{}
	}

	output = &DetachSecurityGroupsFromPrivateLinkGatewayOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DetachSecurityGroupsFromPrivateLinkGateway API operation for PRIVATELINK.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for PRIVATELINK's
// API operation DetachSecurityGroupsFromPrivateLinkGateway for usage and error information.
func (c *PRIVATELINK) DetachSecurityGroupsFromPrivateLinkGateway(input *DetachSecurityGroupsFromPrivateLinkGatewayInput) (*DetachSecurityGroupsFromPrivateLinkGatewayOutput, error) {
	req, out := c.DetachSecurityGroupsFromPrivateLinkGatewayRequest(input)
	return out, req.Send()
}

// DetachSecurityGroupsFromPrivateLinkGatewayWithContext is the same as DetachSecurityGroupsFromPrivateLinkGateway with the addition of
// the ability to pass a context and additional request options.
//
// See DetachSecurityGroupsFromPrivateLinkGateway for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *PRIVATELINK) DetachSecurityGroupsFromPrivateLinkGatewayWithContext(ctx volcengine.Context, input *DetachSecurityGroupsFromPrivateLinkGatewayInput, opts ...request.Option) (*DetachSecurityGroupsFromPrivateLinkGatewayOutput, error) {
	req, out := c.DetachSecurityGroupsFromPrivateLinkGatewayRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DetachSecurityGroupsFromPrivateLinkGatewayInput struct {
	_ struct{} `type:"structure"`

	// PrivateLinkGatewayId is a required field
	PrivateLinkGatewayId *string `type:"string" required:"true"`

	// SecurityGroupIds is a required field
	SecurityGroupIds []*string `type:"list" required:"true"`
}

// String returns the string representation
func (s DetachSecurityGroupsFromPrivateLinkGatewayInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DetachSecurityGroupsFromPrivateLinkGatewayInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DetachSecurityGroupsFromPrivateLinkGatewayInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DetachSecurityGroupsFromPrivateLinkGatewayInput"}
	if s.PrivateLinkGatewayId == nil {
		invalidParams.Add(request.NewErrParamRequired("PrivateLinkGatewayId"))
	}
	if s.SecurityGroupIds == nil {
		invalidParams.Add(request.NewErrParamRequired("SecurityGroupIds"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetPrivateLinkGatewayId sets the PrivateLinkGatewayId field's value.
func (s *DetachSecurityGroupsFromPrivateLinkGatewayInput) SetPrivateLinkGatewayId(v string) *DetachSecurityGroupsFromPrivateLinkGatewayInput {
	s.PrivateLinkGatewayId = &v
	return s
}

// SetSecurityGroupIds sets the SecurityGroupIds field's value.
func (s *DetachSecurityGroupsFromPrivateLinkGatewayInput) SetSecurityGroupIds(v []*string) *DetachSecurityGroupsFromPrivateLinkGatewayInput {
	s.SecurityGroupIds = v
	return s
}

type DetachSecurityGroupsFromPrivateLinkGatewayOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	RequestId *string `type:"string"`
}

// String returns the string representation
func (s DetachSecurityGroupsFromPrivateLinkGatewayOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DetachSecurityGroupsFromPrivateLinkGatewayOutput) GoString() string {
	return s.String()
}

// SetRequestId sets the RequestId field's value.
func (s *DetachSecurityGroupsFromPrivateLinkGatewayOutput) SetRequestId(v string) *DetachSecurityGroupsFromPrivateLinkGatewayOutput {
	s.RequestId = &v
	return s
}

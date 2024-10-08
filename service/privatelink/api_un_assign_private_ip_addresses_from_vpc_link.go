// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package privatelink

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opUnAssignPrivateIpAddressesFromVpcLinkCommon = "UnAssignPrivateIpAddressesFromVpcLink"

// UnAssignPrivateIpAddressesFromVpcLinkCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the UnAssignPrivateIpAddressesFromVpcLinkCommon operation. The "output" return
// value will be populated with the UnAssignPrivateIpAddressesFromVpcLinkCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned UnAssignPrivateIpAddressesFromVpcLinkCommon Request to send the API call to the service.
// the "output" return value is not valid until after UnAssignPrivateIpAddressesFromVpcLinkCommon Send returns without error.
//
// See UnAssignPrivateIpAddressesFromVpcLinkCommon for more information on using the UnAssignPrivateIpAddressesFromVpcLinkCommon
// API call, and error handling.
//
//    // Example sending a request using the UnAssignPrivateIpAddressesFromVpcLinkCommonRequest method.
//    req, resp := client.UnAssignPrivateIpAddressesFromVpcLinkCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *PRIVATELINK) UnAssignPrivateIpAddressesFromVpcLinkCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opUnAssignPrivateIpAddressesFromVpcLinkCommon,
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

// UnAssignPrivateIpAddressesFromVpcLinkCommon API operation for PRIVATELINK.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for PRIVATELINK's
// API operation UnAssignPrivateIpAddressesFromVpcLinkCommon for usage and error information.
func (c *PRIVATELINK) UnAssignPrivateIpAddressesFromVpcLinkCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.UnAssignPrivateIpAddressesFromVpcLinkCommonRequest(input)
	return out, req.Send()
}

// UnAssignPrivateIpAddressesFromVpcLinkCommonWithContext is the same as UnAssignPrivateIpAddressesFromVpcLinkCommon with the addition of
// the ability to pass a context and additional request options.
//
// See UnAssignPrivateIpAddressesFromVpcLinkCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *PRIVATELINK) UnAssignPrivateIpAddressesFromVpcLinkCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.UnAssignPrivateIpAddressesFromVpcLinkCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opUnAssignPrivateIpAddressesFromVpcLink = "UnAssignPrivateIpAddressesFromVpcLink"

// UnAssignPrivateIpAddressesFromVpcLinkRequest generates a "volcengine/request.Request" representing the
// client's request for the UnAssignPrivateIpAddressesFromVpcLink operation. The "output" return
// value will be populated with the UnAssignPrivateIpAddressesFromVpcLinkCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned UnAssignPrivateIpAddressesFromVpcLinkCommon Request to send the API call to the service.
// the "output" return value is not valid until after UnAssignPrivateIpAddressesFromVpcLinkCommon Send returns without error.
//
// See UnAssignPrivateIpAddressesFromVpcLink for more information on using the UnAssignPrivateIpAddressesFromVpcLink
// API call, and error handling.
//
//    // Example sending a request using the UnAssignPrivateIpAddressesFromVpcLinkRequest method.
//    req, resp := client.UnAssignPrivateIpAddressesFromVpcLinkRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *PRIVATELINK) UnAssignPrivateIpAddressesFromVpcLinkRequest(input *UnAssignPrivateIpAddressesFromVpcLinkInput) (req *request.Request, output *UnAssignPrivateIpAddressesFromVpcLinkOutput) {
	op := &request.Operation{
		Name:       opUnAssignPrivateIpAddressesFromVpcLink,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UnAssignPrivateIpAddressesFromVpcLinkInput{}
	}

	output = &UnAssignPrivateIpAddressesFromVpcLinkOutput{}
	req = c.newRequest(op, input, output)

	return
}

// UnAssignPrivateIpAddressesFromVpcLink API operation for PRIVATELINK.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for PRIVATELINK's
// API operation UnAssignPrivateIpAddressesFromVpcLink for usage and error information.
func (c *PRIVATELINK) UnAssignPrivateIpAddressesFromVpcLink(input *UnAssignPrivateIpAddressesFromVpcLinkInput) (*UnAssignPrivateIpAddressesFromVpcLinkOutput, error) {
	req, out := c.UnAssignPrivateIpAddressesFromVpcLinkRequest(input)
	return out, req.Send()
}

// UnAssignPrivateIpAddressesFromVpcLinkWithContext is the same as UnAssignPrivateIpAddressesFromVpcLink with the addition of
// the ability to pass a context and additional request options.
//
// See UnAssignPrivateIpAddressesFromVpcLink for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *PRIVATELINK) UnAssignPrivateIpAddressesFromVpcLinkWithContext(ctx volcengine.Context, input *UnAssignPrivateIpAddressesFromVpcLinkInput, opts ...request.Option) (*UnAssignPrivateIpAddressesFromVpcLinkOutput, error) {
	req, out := c.UnAssignPrivateIpAddressesFromVpcLinkRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type UnAssignPrivateIpAddressesFromVpcLinkInput struct {
	_ struct{} `type:"structure"`

	// PrivateIpAddresses is a required field
	PrivateIpAddresses []*string `type:"list" required:"true"`

	// VpcLinkId is a required field
	VpcLinkId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s UnAssignPrivateIpAddressesFromVpcLinkInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s UnAssignPrivateIpAddressesFromVpcLinkInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UnAssignPrivateIpAddressesFromVpcLinkInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "UnAssignPrivateIpAddressesFromVpcLinkInput"}
	if s.PrivateIpAddresses == nil {
		invalidParams.Add(request.NewErrParamRequired("PrivateIpAddresses"))
	}
	if s.VpcLinkId == nil {
		invalidParams.Add(request.NewErrParamRequired("VpcLinkId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetPrivateIpAddresses sets the PrivateIpAddresses field's value.
func (s *UnAssignPrivateIpAddressesFromVpcLinkInput) SetPrivateIpAddresses(v []*string) *UnAssignPrivateIpAddressesFromVpcLinkInput {
	s.PrivateIpAddresses = v
	return s
}

// SetVpcLinkId sets the VpcLinkId field's value.
func (s *UnAssignPrivateIpAddressesFromVpcLinkInput) SetVpcLinkId(v string) *UnAssignPrivateIpAddressesFromVpcLinkInput {
	s.VpcLinkId = &v
	return s
}

type UnAssignPrivateIpAddressesFromVpcLinkOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	RequestId *string `type:"string"`
}

// String returns the string representation
func (s UnAssignPrivateIpAddressesFromVpcLinkOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s UnAssignPrivateIpAddressesFromVpcLinkOutput) GoString() string {
	return s.String()
}

// SetRequestId sets the RequestId field's value.
func (s *UnAssignPrivateIpAddressesFromVpcLinkOutput) SetRequestId(v string) *UnAssignPrivateIpAddressesFromVpcLinkOutput {
	s.RequestId = &v
	return s
}

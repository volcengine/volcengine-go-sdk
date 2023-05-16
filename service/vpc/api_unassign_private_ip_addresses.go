// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package vpc

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opUnassignPrivateIpAddressesCommon = "UnassignPrivateIpAddresses"

// UnassignPrivateIpAddressesCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the UnassignPrivateIpAddressesCommon operation. The "output" return
// value will be populated with the UnassignPrivateIpAddressesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned UnassignPrivateIpAddressesCommon Request to send the API call to the service.
// the "output" return value is not valid until after UnassignPrivateIpAddressesCommon Send returns without error.
//
// See UnassignPrivateIpAddressesCommon for more information on using the UnassignPrivateIpAddressesCommon
// API call, and error handling.
//
//	// Example sending a request using the UnassignPrivateIpAddressesCommonRequest method.
//	req, resp := client.UnassignPrivateIpAddressesCommonRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *VPC) UnassignPrivateIpAddressesCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opUnassignPrivateIpAddressesCommon,
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

// UnassignPrivateIpAddressesCommon API operation for VPC.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VPC's
// API operation UnassignPrivateIpAddressesCommon for usage and error information.
func (c *VPC) UnassignPrivateIpAddressesCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.UnassignPrivateIpAddressesCommonRequest(input)
	return out, req.Send()
}

// UnassignPrivateIpAddressesCommonWithContext is the same as UnassignPrivateIpAddressesCommon with the addition of
// the ability to pass a context and additional request options.
//
// See UnassignPrivateIpAddressesCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPC) UnassignPrivateIpAddressesCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.UnassignPrivateIpAddressesCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opUnassignPrivateIpAddresses = "UnassignPrivateIpAddresses"

// UnassignPrivateIpAddressesRequest generates a "volcengine/request.Request" representing the
// client's request for the UnassignPrivateIpAddresses operation. The "output" return
// value will be populated with the UnassignPrivateIpAddressesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned UnassignPrivateIpAddressesCommon Request to send the API call to the service.
// the "output" return value is not valid until after UnassignPrivateIpAddressesCommon Send returns without error.
//
// See UnassignPrivateIpAddresses for more information on using the UnassignPrivateIpAddresses
// API call, and error handling.
//
//	// Example sending a request using the UnassignPrivateIpAddressesRequest method.
//	req, resp := client.UnassignPrivateIpAddressesRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *VPC) UnassignPrivateIpAddressesRequest(input *UnassignPrivateIpAddressesInput) (req *request.Request, output *UnassignPrivateIpAddressesOutput) {
	op := &request.Operation{
		Name:       opUnassignPrivateIpAddresses,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UnassignPrivateIpAddressesInput{}
	}

	output = &UnassignPrivateIpAddressesOutput{}
	req = c.newRequest(op, input, output)

	return
}

// UnassignPrivateIpAddresses API operation for VPC.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VPC's
// API operation UnassignPrivateIpAddresses for usage and error information.
func (c *VPC) UnassignPrivateIpAddresses(input *UnassignPrivateIpAddressesInput) (*UnassignPrivateIpAddressesOutput, error) {
	req, out := c.UnassignPrivateIpAddressesRequest(input)
	return out, req.Send()
}

// UnassignPrivateIpAddressesWithContext is the same as UnassignPrivateIpAddresses with the addition of
// the ability to pass a context and additional request options.
//
// See UnassignPrivateIpAddresses for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPC) UnassignPrivateIpAddressesWithContext(ctx volcengine.Context, input *UnassignPrivateIpAddressesInput, opts ...request.Option) (*UnassignPrivateIpAddressesOutput, error) {
	req, out := c.UnassignPrivateIpAddressesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type UnassignPrivateIpAddressesInput struct {
	_ struct{} `type:"structure"`

	// NetworkInterfaceId is a required field
	NetworkInterfaceId *string `type:"string" required:"true"`

	// PrivateIpAddress is a required field
	PrivateIpAddress []*string `type:"list" required:"true"`
}

// String returns the string representation
func (s UnassignPrivateIpAddressesInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s UnassignPrivateIpAddressesInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UnassignPrivateIpAddressesInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "UnassignPrivateIpAddressesInput"}
	if s.NetworkInterfaceId == nil {
		invalidParams.Add(request.NewErrParamRequired("NetworkInterfaceId"))
	}
	if s.PrivateIpAddress == nil {
		invalidParams.Add(request.NewErrParamRequired("PrivateIpAddress"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetNetworkInterfaceId sets the NetworkInterfaceId field's value.
func (s *UnassignPrivateIpAddressesInput) SetNetworkInterfaceId(v string) *UnassignPrivateIpAddressesInput {
	s.NetworkInterfaceId = &v
	return s
}

// SetPrivateIpAddress sets the PrivateIpAddress field's value.
func (s *UnassignPrivateIpAddressesInput) SetPrivateIpAddress(v []*string) *UnassignPrivateIpAddressesInput {
	s.PrivateIpAddress = v
	return s
}

type UnassignPrivateIpAddressesOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	RequestId *string `type:"string"`
}

// String returns the string representation
func (s UnassignPrivateIpAddressesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s UnassignPrivateIpAddressesOutput) GoString() string {
	return s.String()
}

// SetRequestId sets the RequestId field's value.
func (s *UnassignPrivateIpAddressesOutput) SetRequestId(v string) *UnassignPrivateIpAddressesOutput {
	s.RequestId = &v
	return s
}

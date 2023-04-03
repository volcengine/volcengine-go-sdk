// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package vpc

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opReleaseIpv6AddressBandwidthCommon = "ReleaseIpv6AddressBandwidth"

// ReleaseIpv6AddressBandwidthCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ReleaseIpv6AddressBandwidthCommon operation. The "output" return
// value will be populated with the ReleaseIpv6AddressBandwidthCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ReleaseIpv6AddressBandwidthCommon Request to send the API call to the service.
// the "output" return value is not valid until after ReleaseIpv6AddressBandwidthCommon Send returns without error.
//
// See ReleaseIpv6AddressBandwidthCommon for more information on using the ReleaseIpv6AddressBandwidthCommon
// API call, and error handling.
//
//    // Example sending a request using the ReleaseIpv6AddressBandwidthCommonRequest method.
//    req, resp := client.ReleaseIpv6AddressBandwidthCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VPC) ReleaseIpv6AddressBandwidthCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opReleaseIpv6AddressBandwidthCommon,
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

// ReleaseIpv6AddressBandwidthCommon API operation for VPC.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VPC's
// API operation ReleaseIpv6AddressBandwidthCommon for usage and error information.
func (c *VPC) ReleaseIpv6AddressBandwidthCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ReleaseIpv6AddressBandwidthCommonRequest(input)
	return out, req.Send()
}

// ReleaseIpv6AddressBandwidthCommonWithContext is the same as ReleaseIpv6AddressBandwidthCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ReleaseIpv6AddressBandwidthCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPC) ReleaseIpv6AddressBandwidthCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ReleaseIpv6AddressBandwidthCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opReleaseIpv6AddressBandwidth = "ReleaseIpv6AddressBandwidth"

// ReleaseIpv6AddressBandwidthRequest generates a "volcengine/request.Request" representing the
// client's request for the ReleaseIpv6AddressBandwidth operation. The "output" return
// value will be populated with the ReleaseIpv6AddressBandwidthCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ReleaseIpv6AddressBandwidthCommon Request to send the API call to the service.
// the "output" return value is not valid until after ReleaseIpv6AddressBandwidthCommon Send returns without error.
//
// See ReleaseIpv6AddressBandwidth for more information on using the ReleaseIpv6AddressBandwidth
// API call, and error handling.
//
//    // Example sending a request using the ReleaseIpv6AddressBandwidthRequest method.
//    req, resp := client.ReleaseIpv6AddressBandwidthRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VPC) ReleaseIpv6AddressBandwidthRequest(input *ReleaseIpv6AddressBandwidthInput) (req *request.Request, output *ReleaseIpv6AddressBandwidthOutput) {
	op := &request.Operation{
		Name:       opReleaseIpv6AddressBandwidth,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ReleaseIpv6AddressBandwidthInput{}
	}

	output = &ReleaseIpv6AddressBandwidthOutput{}
	req = c.newRequest(op, input, output)

	return
}

// ReleaseIpv6AddressBandwidth API operation for VPC.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VPC's
// API operation ReleaseIpv6AddressBandwidth for usage and error information.
func (c *VPC) ReleaseIpv6AddressBandwidth(input *ReleaseIpv6AddressBandwidthInput) (*ReleaseIpv6AddressBandwidthOutput, error) {
	req, out := c.ReleaseIpv6AddressBandwidthRequest(input)
	return out, req.Send()
}

// ReleaseIpv6AddressBandwidthWithContext is the same as ReleaseIpv6AddressBandwidth with the addition of
// the ability to pass a context and additional request options.
//
// See ReleaseIpv6AddressBandwidth for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPC) ReleaseIpv6AddressBandwidthWithContext(ctx volcengine.Context, input *ReleaseIpv6AddressBandwidthInput, opts ...request.Option) (*ReleaseIpv6AddressBandwidthOutput, error) {
	req, out := c.ReleaseIpv6AddressBandwidthRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ReleaseIpv6AddressBandwidthInput struct {
	_ struct{} `type:"structure"`

	// AllocationId is a required field
	AllocationId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s ReleaseIpv6AddressBandwidthInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ReleaseIpv6AddressBandwidthInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ReleaseIpv6AddressBandwidthInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ReleaseIpv6AddressBandwidthInput"}
	if s.AllocationId == nil {
		invalidParams.Add(request.NewErrParamRequired("AllocationId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetAllocationId sets the AllocationId field's value.
func (s *ReleaseIpv6AddressBandwidthInput) SetAllocationId(v string) *ReleaseIpv6AddressBandwidthInput {
	s.AllocationId = &v
	return s
}

type ReleaseIpv6AddressBandwidthOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	PreOrderNumber *string `type:"string"`

	RequestId *string `type:"string"`
}

// String returns the string representation
func (s ReleaseIpv6AddressBandwidthOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ReleaseIpv6AddressBandwidthOutput) GoString() string {
	return s.String()
}

// SetPreOrderNumber sets the PreOrderNumber field's value.
func (s *ReleaseIpv6AddressBandwidthOutput) SetPreOrderNumber(v string) *ReleaseIpv6AddressBandwidthOutput {
	s.PreOrderNumber = &v
	return s
}

// SetRequestId sets the RequestId field's value.
func (s *ReleaseIpv6AddressBandwidthOutput) SetRequestId(v string) *ReleaseIpv6AddressBandwidthOutput {
	s.RequestId = &v
	return s
}

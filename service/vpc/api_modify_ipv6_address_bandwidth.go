// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package vpc

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opModifyIpv6AddressBandwidthCommon = "ModifyIpv6AddressBandwidth"

// ModifyIpv6AddressBandwidthCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ModifyIpv6AddressBandwidthCommon operation. The "output" return
// value will be populated with the ModifyIpv6AddressBandwidthCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyIpv6AddressBandwidthCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyIpv6AddressBandwidthCommon Send returns without error.
//
// See ModifyIpv6AddressBandwidthCommon for more information on using the ModifyIpv6AddressBandwidthCommon
// API call, and error handling.
//
//    // Example sending a request using the ModifyIpv6AddressBandwidthCommonRequest method.
//    req, resp := client.ModifyIpv6AddressBandwidthCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VPC) ModifyIpv6AddressBandwidthCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opModifyIpv6AddressBandwidthCommon,
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

// ModifyIpv6AddressBandwidthCommon API operation for VPC.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VPC's
// API operation ModifyIpv6AddressBandwidthCommon for usage and error information.
func (c *VPC) ModifyIpv6AddressBandwidthCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ModifyIpv6AddressBandwidthCommonRequest(input)
	return out, req.Send()
}

// ModifyIpv6AddressBandwidthCommonWithContext is the same as ModifyIpv6AddressBandwidthCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyIpv6AddressBandwidthCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPC) ModifyIpv6AddressBandwidthCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ModifyIpv6AddressBandwidthCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opModifyIpv6AddressBandwidth = "ModifyIpv6AddressBandwidth"

// ModifyIpv6AddressBandwidthRequest generates a "volcengine/request.Request" representing the
// client's request for the ModifyIpv6AddressBandwidth operation. The "output" return
// value will be populated with the ModifyIpv6AddressBandwidthCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyIpv6AddressBandwidthCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyIpv6AddressBandwidthCommon Send returns without error.
//
// See ModifyIpv6AddressBandwidth for more information on using the ModifyIpv6AddressBandwidth
// API call, and error handling.
//
//    // Example sending a request using the ModifyIpv6AddressBandwidthRequest method.
//    req, resp := client.ModifyIpv6AddressBandwidthRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VPC) ModifyIpv6AddressBandwidthRequest(input *ModifyIpv6AddressBandwidthInput) (req *request.Request, output *ModifyIpv6AddressBandwidthOutput) {
	op := &request.Operation{
		Name:       opModifyIpv6AddressBandwidth,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ModifyIpv6AddressBandwidthInput{}
	}

	output = &ModifyIpv6AddressBandwidthOutput{}
	req = c.newRequest(op, input, output)

	return
}

// ModifyIpv6AddressBandwidth API operation for VPC.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VPC's
// API operation ModifyIpv6AddressBandwidth for usage and error information.
func (c *VPC) ModifyIpv6AddressBandwidth(input *ModifyIpv6AddressBandwidthInput) (*ModifyIpv6AddressBandwidthOutput, error) {
	req, out := c.ModifyIpv6AddressBandwidthRequest(input)
	return out, req.Send()
}

// ModifyIpv6AddressBandwidthWithContext is the same as ModifyIpv6AddressBandwidth with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyIpv6AddressBandwidth for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPC) ModifyIpv6AddressBandwidthWithContext(ctx volcengine.Context, input *ModifyIpv6AddressBandwidthInput, opts ...request.Option) (*ModifyIpv6AddressBandwidthOutput, error) {
	req, out := c.ModifyIpv6AddressBandwidthRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ModifyIpv6AddressBandwidthInput struct {
	_ struct{} `type:"structure"`

	// AllocationId is a required field
	AllocationId *string `type:"string" required:"true"`

	Bandwidth *int64 `type:"integer"`
}

// String returns the string representation
func (s ModifyIpv6AddressBandwidthInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyIpv6AddressBandwidthInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ModifyIpv6AddressBandwidthInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ModifyIpv6AddressBandwidthInput"}
	if s.AllocationId == nil {
		invalidParams.Add(request.NewErrParamRequired("AllocationId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetAllocationId sets the AllocationId field's value.
func (s *ModifyIpv6AddressBandwidthInput) SetAllocationId(v string) *ModifyIpv6AddressBandwidthInput {
	s.AllocationId = &v
	return s
}

// SetBandwidth sets the Bandwidth field's value.
func (s *ModifyIpv6AddressBandwidthInput) SetBandwidth(v int64) *ModifyIpv6AddressBandwidthInput {
	s.Bandwidth = &v
	return s
}

type ModifyIpv6AddressBandwidthOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	RequestId *string `type:"string"`
}

// String returns the string representation
func (s ModifyIpv6AddressBandwidthOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyIpv6AddressBandwidthOutput) GoString() string {
	return s.String()
}

// SetRequestId sets the RequestId field's value.
func (s *ModifyIpv6AddressBandwidthOutput) SetRequestId(v string) *ModifyIpv6AddressBandwidthOutput {
	s.RequestId = &v
	return s
}
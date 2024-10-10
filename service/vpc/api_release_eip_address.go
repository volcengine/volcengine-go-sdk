// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package vpc

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opReleaseEipAddressCommon = "ReleaseEipAddress"

// ReleaseEipAddressCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ReleaseEipAddressCommon operation. The "output" return
// value will be populated with the ReleaseEipAddressCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ReleaseEipAddressCommon Request to send the API call to the service.
// the "output" return value is not valid until after ReleaseEipAddressCommon Send returns without error.
//
// See ReleaseEipAddressCommon for more information on using the ReleaseEipAddressCommon
// API call, and error handling.
//
//    // Example sending a request using the ReleaseEipAddressCommonRequest method.
//    req, resp := client.ReleaseEipAddressCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VPC) ReleaseEipAddressCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opReleaseEipAddressCommon,
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

// ReleaseEipAddressCommon API operation for VPC.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VPC's
// API operation ReleaseEipAddressCommon for usage and error information.
func (c *VPC) ReleaseEipAddressCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ReleaseEipAddressCommonRequest(input)
	return out, req.Send()
}

// ReleaseEipAddressCommonWithContext is the same as ReleaseEipAddressCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ReleaseEipAddressCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPC) ReleaseEipAddressCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ReleaseEipAddressCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opReleaseEipAddress = "ReleaseEipAddress"

// ReleaseEipAddressRequest generates a "volcengine/request.Request" representing the
// client's request for the ReleaseEipAddress operation. The "output" return
// value will be populated with the ReleaseEipAddressCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ReleaseEipAddressCommon Request to send the API call to the service.
// the "output" return value is not valid until after ReleaseEipAddressCommon Send returns without error.
//
// See ReleaseEipAddress for more information on using the ReleaseEipAddress
// API call, and error handling.
//
//    // Example sending a request using the ReleaseEipAddressRequest method.
//    req, resp := client.ReleaseEipAddressRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VPC) ReleaseEipAddressRequest(input *ReleaseEipAddressInput) (req *request.Request, output *ReleaseEipAddressOutput) {
	op := &request.Operation{
		Name:       opReleaseEipAddress,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ReleaseEipAddressInput{}
	}

	output = &ReleaseEipAddressOutput{}
	req = c.newRequest(op, input, output)

	return
}

// ReleaseEipAddress API operation for VPC.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VPC's
// API operation ReleaseEipAddress for usage and error information.
func (c *VPC) ReleaseEipAddress(input *ReleaseEipAddressInput) (*ReleaseEipAddressOutput, error) {
	req, out := c.ReleaseEipAddressRequest(input)
	return out, req.Send()
}

// ReleaseEipAddressWithContext is the same as ReleaseEipAddress with the addition of
// the ability to pass a context and additional request options.
//
// See ReleaseEipAddress for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPC) ReleaseEipAddressWithContext(ctx volcengine.Context, input *ReleaseEipAddressInput, opts ...request.Option) (*ReleaseEipAddressOutput, error) {
	req, out := c.ReleaseEipAddressRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ReleaseEipAddressInput struct {
	_ struct{} `type:"structure"`

	// AllocationId is a required field
	AllocationId *string `type:"string" required:"true"`

	ClientToken *string `type:"string"`
}

// String returns the string representation
func (s ReleaseEipAddressInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ReleaseEipAddressInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ReleaseEipAddressInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ReleaseEipAddressInput"}
	if s.AllocationId == nil {
		invalidParams.Add(request.NewErrParamRequired("AllocationId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetAllocationId sets the AllocationId field's value.
func (s *ReleaseEipAddressInput) SetAllocationId(v string) *ReleaseEipAddressInput {
	s.AllocationId = &v
	return s
}

// SetClientToken sets the ClientToken field's value.
func (s *ReleaseEipAddressInput) SetClientToken(v string) *ReleaseEipAddressInput {
	s.ClientToken = &v
	return s
}

type ReleaseEipAddressOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	RequestId *string `type:"string"`
}

// String returns the string representation
func (s ReleaseEipAddressOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ReleaseEipAddressOutput) GoString() string {
	return s.String()
}

// SetRequestId sets the RequestId field's value.
func (s *ReleaseEipAddressOutput) SetRequestId(v string) *ReleaseEipAddressOutput {
	s.RequestId = &v
	return s
}

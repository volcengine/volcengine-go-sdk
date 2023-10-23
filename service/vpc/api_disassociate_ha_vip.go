// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package vpc

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDisassociateHaVipCommon = "DisassociateHaVip"

// DisassociateHaVipCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DisassociateHaVipCommon operation. The "output" return
// value will be populated with the DisassociateHaVipCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DisassociateHaVipCommon Request to send the API call to the service.
// the "output" return value is not valid until after DisassociateHaVipCommon Send returns without error.
//
// See DisassociateHaVipCommon for more information on using the DisassociateHaVipCommon
// API call, and error handling.
//
//    // Example sending a request using the DisassociateHaVipCommonRequest method.
//    req, resp := client.DisassociateHaVipCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VPC) DisassociateHaVipCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDisassociateHaVipCommon,
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

// DisassociateHaVipCommon API operation for VPC.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VPC's
// API operation DisassociateHaVipCommon for usage and error information.
func (c *VPC) DisassociateHaVipCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DisassociateHaVipCommonRequest(input)
	return out, req.Send()
}

// DisassociateHaVipCommonWithContext is the same as DisassociateHaVipCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DisassociateHaVipCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPC) DisassociateHaVipCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DisassociateHaVipCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDisassociateHaVip = "DisassociateHaVip"

// DisassociateHaVipRequest generates a "volcengine/request.Request" representing the
// client's request for the DisassociateHaVip operation. The "output" return
// value will be populated with the DisassociateHaVipCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DisassociateHaVipCommon Request to send the API call to the service.
// the "output" return value is not valid until after DisassociateHaVipCommon Send returns without error.
//
// See DisassociateHaVip for more information on using the DisassociateHaVip
// API call, and error handling.
//
//    // Example sending a request using the DisassociateHaVipRequest method.
//    req, resp := client.DisassociateHaVipRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VPC) DisassociateHaVipRequest(input *DisassociateHaVipInput) (req *request.Request, output *DisassociateHaVipOutput) {
	op := &request.Operation{
		Name:       opDisassociateHaVip,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DisassociateHaVipInput{}
	}

	output = &DisassociateHaVipOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DisassociateHaVip API operation for VPC.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VPC's
// API operation DisassociateHaVip for usage and error information.
func (c *VPC) DisassociateHaVip(input *DisassociateHaVipInput) (*DisassociateHaVipOutput, error) {
	req, out := c.DisassociateHaVipRequest(input)
	return out, req.Send()
}

// DisassociateHaVipWithContext is the same as DisassociateHaVip with the addition of
// the ability to pass a context and additional request options.
//
// See DisassociateHaVip for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPC) DisassociateHaVipWithContext(ctx volcengine.Context, input *DisassociateHaVipInput, opts ...request.Option) (*DisassociateHaVipOutput, error) {
	req, out := c.DisassociateHaVipRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DisassociateHaVipInput struct {
	_ struct{} `type:"structure"`

	ClientToken *string `type:"string"`

	// HaVipId is a required field
	HaVipId *string `type:"string" required:"true"`

	// InstanceId is a required field
	InstanceId *string `type:"string" required:"true"`

	InstanceType *string `type:"string"`
}

// String returns the string representation
func (s DisassociateHaVipInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DisassociateHaVipInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DisassociateHaVipInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DisassociateHaVipInput"}
	if s.HaVipId == nil {
		invalidParams.Add(request.NewErrParamRequired("HaVipId"))
	}
	if s.InstanceId == nil {
		invalidParams.Add(request.NewErrParamRequired("InstanceId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetClientToken sets the ClientToken field's value.
func (s *DisassociateHaVipInput) SetClientToken(v string) *DisassociateHaVipInput {
	s.ClientToken = &v
	return s
}

// SetHaVipId sets the HaVipId field's value.
func (s *DisassociateHaVipInput) SetHaVipId(v string) *DisassociateHaVipInput {
	s.HaVipId = &v
	return s
}

// SetInstanceId sets the InstanceId field's value.
func (s *DisassociateHaVipInput) SetInstanceId(v string) *DisassociateHaVipInput {
	s.InstanceId = &v
	return s
}

// SetInstanceType sets the InstanceType field's value.
func (s *DisassociateHaVipInput) SetInstanceType(v string) *DisassociateHaVipInput {
	s.InstanceType = &v
	return s
}

type DisassociateHaVipOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	RequestId *string `type:"string"`
}

// String returns the string representation
func (s DisassociateHaVipOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DisassociateHaVipOutput) GoString() string {
	return s.String()
}

// SetRequestId sets the RequestId field's value.
func (s *DisassociateHaVipOutput) SetRequestId(v string) *DisassociateHaVipOutput {
	s.RequestId = &v
	return s
}

// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package vpc

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDisassociateVpcUserCidrBlockCommon = "DisassociateVpcUserCidrBlock"

// DisassociateVpcUserCidrBlockCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DisassociateVpcUserCidrBlockCommon operation. The "output" return
// value will be populated with the DisassociateVpcUserCidrBlockCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DisassociateVpcUserCidrBlockCommon Request to send the API call to the service.
// the "output" return value is not valid until after DisassociateVpcUserCidrBlockCommon Send returns without error.
//
// See DisassociateVpcUserCidrBlockCommon for more information on using the DisassociateVpcUserCidrBlockCommon
// API call, and error handling.
//
//    // Example sending a request using the DisassociateVpcUserCidrBlockCommonRequest method.
//    req, resp := client.DisassociateVpcUserCidrBlockCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VPC) DisassociateVpcUserCidrBlockCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDisassociateVpcUserCidrBlockCommon,
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

// DisassociateVpcUserCidrBlockCommon API operation for VPC.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VPC's
// API operation DisassociateVpcUserCidrBlockCommon for usage and error information.
func (c *VPC) DisassociateVpcUserCidrBlockCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DisassociateVpcUserCidrBlockCommonRequest(input)
	return out, req.Send()
}

// DisassociateVpcUserCidrBlockCommonWithContext is the same as DisassociateVpcUserCidrBlockCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DisassociateVpcUserCidrBlockCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPC) DisassociateVpcUserCidrBlockCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DisassociateVpcUserCidrBlockCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDisassociateVpcUserCidrBlock = "DisassociateVpcUserCidrBlock"

// DisassociateVpcUserCidrBlockRequest generates a "volcengine/request.Request" representing the
// client's request for the DisassociateVpcUserCidrBlock operation. The "output" return
// value will be populated with the DisassociateVpcUserCidrBlockCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DisassociateVpcUserCidrBlockCommon Request to send the API call to the service.
// the "output" return value is not valid until after DisassociateVpcUserCidrBlockCommon Send returns without error.
//
// See DisassociateVpcUserCidrBlock for more information on using the DisassociateVpcUserCidrBlock
// API call, and error handling.
//
//    // Example sending a request using the DisassociateVpcUserCidrBlockRequest method.
//    req, resp := client.DisassociateVpcUserCidrBlockRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VPC) DisassociateVpcUserCidrBlockRequest(input *DisassociateVpcUserCidrBlockInput) (req *request.Request, output *DisassociateVpcUserCidrBlockOutput) {
	op := &request.Operation{
		Name:       opDisassociateVpcUserCidrBlock,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DisassociateVpcUserCidrBlockInput{}
	}

	output = &DisassociateVpcUserCidrBlockOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DisassociateVpcUserCidrBlock API operation for VPC.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VPC's
// API operation DisassociateVpcUserCidrBlock for usage and error information.
func (c *VPC) DisassociateVpcUserCidrBlock(input *DisassociateVpcUserCidrBlockInput) (*DisassociateVpcUserCidrBlockOutput, error) {
	req, out := c.DisassociateVpcUserCidrBlockRequest(input)
	return out, req.Send()
}

// DisassociateVpcUserCidrBlockWithContext is the same as DisassociateVpcUserCidrBlock with the addition of
// the ability to pass a context and additional request options.
//
// See DisassociateVpcUserCidrBlock for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPC) DisassociateVpcUserCidrBlockWithContext(ctx volcengine.Context, input *DisassociateVpcUserCidrBlockInput, opts ...request.Option) (*DisassociateVpcUserCidrBlockOutput, error) {
	req, out := c.DisassociateVpcUserCidrBlockRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DisassociateVpcUserCidrBlockInput struct {
	_ struct{} `type:"structure"`

	// UserCidrBlock is a required field
	UserCidrBlock *string `type:"string" required:"true"`

	// VpcId is a required field
	VpcId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DisassociateVpcUserCidrBlockInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DisassociateVpcUserCidrBlockInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DisassociateVpcUserCidrBlockInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DisassociateVpcUserCidrBlockInput"}
	if s.UserCidrBlock == nil {
		invalidParams.Add(request.NewErrParamRequired("UserCidrBlock"))
	}
	if s.VpcId == nil {
		invalidParams.Add(request.NewErrParamRequired("VpcId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetUserCidrBlock sets the UserCidrBlock field's value.
func (s *DisassociateVpcUserCidrBlockInput) SetUserCidrBlock(v string) *DisassociateVpcUserCidrBlockInput {
	s.UserCidrBlock = &v
	return s
}

// SetVpcId sets the VpcId field's value.
func (s *DisassociateVpcUserCidrBlockInput) SetVpcId(v string) *DisassociateVpcUserCidrBlockInput {
	s.VpcId = &v
	return s
}

type DisassociateVpcUserCidrBlockOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	RequestId *string `type:"string"`
}

// String returns the string representation
func (s DisassociateVpcUserCidrBlockOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DisassociateVpcUserCidrBlockOutput) GoString() string {
	return s.String()
}

// SetRequestId sets the RequestId field's value.
func (s *DisassociateVpcUserCidrBlockOutput) SetRequestId(v string) *DisassociateVpcUserCidrBlockOutput {
	s.RequestId = &v
	return s
}
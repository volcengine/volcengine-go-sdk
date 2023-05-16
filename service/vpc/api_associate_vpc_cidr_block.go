// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package vpc

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opAssociateVpcCidrBlockCommon = "AssociateVpcCidrBlock"

// AssociateVpcCidrBlockCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the AssociateVpcCidrBlockCommon operation. The "output" return
// value will be populated with the AssociateVpcCidrBlockCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned AssociateVpcCidrBlockCommon Request to send the API call to the service.
// the "output" return value is not valid until after AssociateVpcCidrBlockCommon Send returns without error.
//
// See AssociateVpcCidrBlockCommon for more information on using the AssociateVpcCidrBlockCommon
// API call, and error handling.
//
//	// Example sending a request using the AssociateVpcCidrBlockCommonRequest method.
//	req, resp := client.AssociateVpcCidrBlockCommonRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *VPC) AssociateVpcCidrBlockCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opAssociateVpcCidrBlockCommon,
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

// AssociateVpcCidrBlockCommon API operation for VPC.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VPC's
// API operation AssociateVpcCidrBlockCommon for usage and error information.
func (c *VPC) AssociateVpcCidrBlockCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.AssociateVpcCidrBlockCommonRequest(input)
	return out, req.Send()
}

// AssociateVpcCidrBlockCommonWithContext is the same as AssociateVpcCidrBlockCommon with the addition of
// the ability to pass a context and additional request options.
//
// See AssociateVpcCidrBlockCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPC) AssociateVpcCidrBlockCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.AssociateVpcCidrBlockCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opAssociateVpcCidrBlock = "AssociateVpcCidrBlock"

// AssociateVpcCidrBlockRequest generates a "volcengine/request.Request" representing the
// client's request for the AssociateVpcCidrBlock operation. The "output" return
// value will be populated with the AssociateVpcCidrBlockCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned AssociateVpcCidrBlockCommon Request to send the API call to the service.
// the "output" return value is not valid until after AssociateVpcCidrBlockCommon Send returns without error.
//
// See AssociateVpcCidrBlock for more information on using the AssociateVpcCidrBlock
// API call, and error handling.
//
//	// Example sending a request using the AssociateVpcCidrBlockRequest method.
//	req, resp := client.AssociateVpcCidrBlockRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *VPC) AssociateVpcCidrBlockRequest(input *AssociateVpcCidrBlockInput) (req *request.Request, output *AssociateVpcCidrBlockOutput) {
	op := &request.Operation{
		Name:       opAssociateVpcCidrBlock,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &AssociateVpcCidrBlockInput{}
	}

	output = &AssociateVpcCidrBlockOutput{}
	req = c.newRequest(op, input, output)

	return
}

// AssociateVpcCidrBlock API operation for VPC.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VPC's
// API operation AssociateVpcCidrBlock for usage and error information.
func (c *VPC) AssociateVpcCidrBlock(input *AssociateVpcCidrBlockInput) (*AssociateVpcCidrBlockOutput, error) {
	req, out := c.AssociateVpcCidrBlockRequest(input)
	return out, req.Send()
}

// AssociateVpcCidrBlockWithContext is the same as AssociateVpcCidrBlock with the addition of
// the ability to pass a context and additional request options.
//
// See AssociateVpcCidrBlock for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPC) AssociateVpcCidrBlockWithContext(ctx volcengine.Context, input *AssociateVpcCidrBlockInput, opts ...request.Option) (*AssociateVpcCidrBlockOutput, error) {
	req, out := c.AssociateVpcCidrBlockRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type AssociateVpcCidrBlockInput struct {
	_ struct{} `type:"structure"`

	SecondaryCidrBlock *string `type:"string"`

	// VpcId is a required field
	VpcId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s AssociateVpcCidrBlockInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s AssociateVpcCidrBlockInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AssociateVpcCidrBlockInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "AssociateVpcCidrBlockInput"}
	if s.VpcId == nil {
		invalidParams.Add(request.NewErrParamRequired("VpcId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetSecondaryCidrBlock sets the SecondaryCidrBlock field's value.
func (s *AssociateVpcCidrBlockInput) SetSecondaryCidrBlock(v string) *AssociateVpcCidrBlockInput {
	s.SecondaryCidrBlock = &v
	return s
}

// SetVpcId sets the VpcId field's value.
func (s *AssociateVpcCidrBlockInput) SetVpcId(v string) *AssociateVpcCidrBlockInput {
	s.VpcId = &v
	return s
}

type AssociateVpcCidrBlockOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	RequestId *string `type:"string"`

	SecondaryCidrBlocks []*string `type:"list"`

	VpcId *string `type:"string"`
}

// String returns the string representation
func (s AssociateVpcCidrBlockOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s AssociateVpcCidrBlockOutput) GoString() string {
	return s.String()
}

// SetRequestId sets the RequestId field's value.
func (s *AssociateVpcCidrBlockOutput) SetRequestId(v string) *AssociateVpcCidrBlockOutput {
	s.RequestId = &v
	return s
}

// SetSecondaryCidrBlocks sets the SecondaryCidrBlocks field's value.
func (s *AssociateVpcCidrBlockOutput) SetSecondaryCidrBlocks(v []*string) *AssociateVpcCidrBlockOutput {
	s.SecondaryCidrBlocks = v
	return s
}

// SetVpcId sets the VpcId field's value.
func (s *AssociateVpcCidrBlockOutput) SetVpcId(v string) *AssociateVpcCidrBlockOutput {
	s.VpcId = &v
	return s
}

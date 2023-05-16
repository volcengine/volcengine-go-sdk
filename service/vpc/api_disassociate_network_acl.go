// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package vpc

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDisassociateNetworkAclCommon = "DisassociateNetworkAcl"

// DisassociateNetworkAclCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DisassociateNetworkAclCommon operation. The "output" return
// value will be populated with the DisassociateNetworkAclCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DisassociateNetworkAclCommon Request to send the API call to the service.
// the "output" return value is not valid until after DisassociateNetworkAclCommon Send returns without error.
//
// See DisassociateNetworkAclCommon for more information on using the DisassociateNetworkAclCommon
// API call, and error handling.
//
//	// Example sending a request using the DisassociateNetworkAclCommonRequest method.
//	req, resp := client.DisassociateNetworkAclCommonRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *VPC) DisassociateNetworkAclCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDisassociateNetworkAclCommon,
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

// DisassociateNetworkAclCommon API operation for VPC.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VPC's
// API operation DisassociateNetworkAclCommon for usage and error information.
func (c *VPC) DisassociateNetworkAclCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DisassociateNetworkAclCommonRequest(input)
	return out, req.Send()
}

// DisassociateNetworkAclCommonWithContext is the same as DisassociateNetworkAclCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DisassociateNetworkAclCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPC) DisassociateNetworkAclCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DisassociateNetworkAclCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDisassociateNetworkAcl = "DisassociateNetworkAcl"

// DisassociateNetworkAclRequest generates a "volcengine/request.Request" representing the
// client's request for the DisassociateNetworkAcl operation. The "output" return
// value will be populated with the DisassociateNetworkAclCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DisassociateNetworkAclCommon Request to send the API call to the service.
// the "output" return value is not valid until after DisassociateNetworkAclCommon Send returns without error.
//
// See DisassociateNetworkAcl for more information on using the DisassociateNetworkAcl
// API call, and error handling.
//
//	// Example sending a request using the DisassociateNetworkAclRequest method.
//	req, resp := client.DisassociateNetworkAclRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *VPC) DisassociateNetworkAclRequest(input *DisassociateNetworkAclInput) (req *request.Request, output *DisassociateNetworkAclOutput) {
	op := &request.Operation{
		Name:       opDisassociateNetworkAcl,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DisassociateNetworkAclInput{}
	}

	output = &DisassociateNetworkAclOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DisassociateNetworkAcl API operation for VPC.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VPC's
// API operation DisassociateNetworkAcl for usage and error information.
func (c *VPC) DisassociateNetworkAcl(input *DisassociateNetworkAclInput) (*DisassociateNetworkAclOutput, error) {
	req, out := c.DisassociateNetworkAclRequest(input)
	return out, req.Send()
}

// DisassociateNetworkAclWithContext is the same as DisassociateNetworkAcl with the addition of
// the ability to pass a context and additional request options.
//
// See DisassociateNetworkAcl for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPC) DisassociateNetworkAclWithContext(ctx volcengine.Context, input *DisassociateNetworkAclInput, opts ...request.Option) (*DisassociateNetworkAclOutput, error) {
	req, out := c.DisassociateNetworkAclRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DisassociateNetworkAclInput struct {
	_ struct{} `type:"structure"`

	// NetworkAclId is a required field
	NetworkAclId *string `type:"string" required:"true"`

	Resource []*ResourceForDisassociateNetworkAclInput `type:"list"`
}

// String returns the string representation
func (s DisassociateNetworkAclInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DisassociateNetworkAclInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DisassociateNetworkAclInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DisassociateNetworkAclInput"}
	if s.NetworkAclId == nil {
		invalidParams.Add(request.NewErrParamRequired("NetworkAclId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetNetworkAclId sets the NetworkAclId field's value.
func (s *DisassociateNetworkAclInput) SetNetworkAclId(v string) *DisassociateNetworkAclInput {
	s.NetworkAclId = &v
	return s
}

// SetResource sets the Resource field's value.
func (s *DisassociateNetworkAclInput) SetResource(v []*ResourceForDisassociateNetworkAclInput) *DisassociateNetworkAclInput {
	s.Resource = v
	return s
}

type DisassociateNetworkAclOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	RequestId *string `type:"string"`
}

// String returns the string representation
func (s DisassociateNetworkAclOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DisassociateNetworkAclOutput) GoString() string {
	return s.String()
}

// SetRequestId sets the RequestId field's value.
func (s *DisassociateNetworkAclOutput) SetRequestId(v string) *DisassociateNetworkAclOutput {
	s.RequestId = &v
	return s
}

type ResourceForDisassociateNetworkAclInput struct {
	_ struct{} `type:"structure"`

	ResourceId *string `type:"string"`
}

// String returns the string representation
func (s ResourceForDisassociateNetworkAclInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ResourceForDisassociateNetworkAclInput) GoString() string {
	return s.String()
}

// SetResourceId sets the ResourceId field's value.
func (s *ResourceForDisassociateNetworkAclInput) SetResourceId(v string) *ResourceForDisassociateNetworkAclInput {
	s.ResourceId = &v
	return s
}

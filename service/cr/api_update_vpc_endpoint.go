// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package cr

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opUpdateVpcEndpointCommon = "UpdateVpcEndpoint"

// UpdateVpcEndpointCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the UpdateVpcEndpointCommon operation. The "output" return
// value will be populated with the UpdateVpcEndpointCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned UpdateVpcEndpointCommon Request to send the API call to the service.
// the "output" return value is not valid until after UpdateVpcEndpointCommon Send returns without error.
//
// See UpdateVpcEndpointCommon for more information on using the UpdateVpcEndpointCommon
// API call, and error handling.
//
//    // Example sending a request using the UpdateVpcEndpointCommonRequest method.
//    req, resp := client.UpdateVpcEndpointCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CR) UpdateVpcEndpointCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opUpdateVpcEndpointCommon,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &map[string]interface{}{}
	}

	output = &map[string]interface{}{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// UpdateVpcEndpointCommon API operation for CR.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CR's
// API operation UpdateVpcEndpointCommon for usage and error information.
func (c *CR) UpdateVpcEndpointCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.UpdateVpcEndpointCommonRequest(input)
	return out, req.Send()
}

// UpdateVpcEndpointCommonWithContext is the same as UpdateVpcEndpointCommon with the addition of
// the ability to pass a context and additional request options.
//
// See UpdateVpcEndpointCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CR) UpdateVpcEndpointCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.UpdateVpcEndpointCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opUpdateVpcEndpoint = "UpdateVpcEndpoint"

// UpdateVpcEndpointRequest generates a "volcengine/request.Request" representing the
// client's request for the UpdateVpcEndpoint operation. The "output" return
// value will be populated with the UpdateVpcEndpointCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned UpdateVpcEndpointCommon Request to send the API call to the service.
// the "output" return value is not valid until after UpdateVpcEndpointCommon Send returns without error.
//
// See UpdateVpcEndpoint for more information on using the UpdateVpcEndpoint
// API call, and error handling.
//
//    // Example sending a request using the UpdateVpcEndpointRequest method.
//    req, resp := client.UpdateVpcEndpointRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CR) UpdateVpcEndpointRequest(input *UpdateVpcEndpointInput) (req *request.Request, output *UpdateVpcEndpointOutput) {
	op := &request.Operation{
		Name:       opUpdateVpcEndpoint,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateVpcEndpointInput{}
	}

	output = &UpdateVpcEndpointOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// UpdateVpcEndpoint API operation for CR.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CR's
// API operation UpdateVpcEndpoint for usage and error information.
func (c *CR) UpdateVpcEndpoint(input *UpdateVpcEndpointInput) (*UpdateVpcEndpointOutput, error) {
	req, out := c.UpdateVpcEndpointRequest(input)
	return out, req.Send()
}

// UpdateVpcEndpointWithContext is the same as UpdateVpcEndpoint with the addition of
// the ability to pass a context and additional request options.
//
// See UpdateVpcEndpoint for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CR) UpdateVpcEndpointWithContext(ctx volcengine.Context, input *UpdateVpcEndpointInput, opts ...request.Option) (*UpdateVpcEndpointOutput, error) {
	req, out := c.UpdateVpcEndpointRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type UpdateVpcEndpointInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	// Registry is a required field
	Registry *string `min:"3" max:"30" type:"string" json:",omitempty" required:"true"`

	Vpcs []*VpcForUpdateVpcEndpointInput `type:"list" json:",omitempty"`
}

// String returns the string representation
func (s UpdateVpcEndpointInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s UpdateVpcEndpointInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateVpcEndpointInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "UpdateVpcEndpointInput"}
	if s.Registry == nil {
		invalidParams.Add(request.NewErrParamRequired("Registry"))
	}
	if s.Registry != nil && len(*s.Registry) < 3 {
		invalidParams.Add(request.NewErrParamMinLen("Registry", 3))
	}
	if s.Registry != nil && len(*s.Registry) > 30 {
		invalidParams.Add(request.NewErrParamMaxLen("Registry", 30, *s.Registry))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetRegistry sets the Registry field's value.
func (s *UpdateVpcEndpointInput) SetRegistry(v string) *UpdateVpcEndpointInput {
	s.Registry = &v
	return s
}

// SetVpcs sets the Vpcs field's value.
func (s *UpdateVpcEndpointInput) SetVpcs(v []*VpcForUpdateVpcEndpointInput) *UpdateVpcEndpointInput {
	s.Vpcs = v
	return s
}

type UpdateVpcEndpointOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s UpdateVpcEndpointOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s UpdateVpcEndpointOutput) GoString() string {
	return s.String()
}

type VpcForUpdateVpcEndpointInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	AccountId *int64 `type:"int64" json:",omitempty"`

	SubnetId *string `type:"string" json:",omitempty"`

	VpcId *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s VpcForUpdateVpcEndpointInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s VpcForUpdateVpcEndpointInput) GoString() string {
	return s.String()
}

// SetAccountId sets the AccountId field's value.
func (s *VpcForUpdateVpcEndpointInput) SetAccountId(v int64) *VpcForUpdateVpcEndpointInput {
	s.AccountId = &v
	return s
}

// SetSubnetId sets the SubnetId field's value.
func (s *VpcForUpdateVpcEndpointInput) SetSubnetId(v string) *VpcForUpdateVpcEndpointInput {
	s.SubnetId = &v
	return s
}

// SetVpcId sets the VpcId field's value.
func (s *VpcForUpdateVpcEndpointInput) SetVpcId(v string) *VpcForUpdateVpcEndpointInput {
	s.VpcId = &v
	return s
}

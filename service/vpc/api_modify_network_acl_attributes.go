// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package vpc

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opModifyNetworkAclAttributesCommon = "ModifyNetworkAclAttributes"

// ModifyNetworkAclAttributesCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ModifyNetworkAclAttributesCommon operation. The "output" return
// value will be populated with the ModifyNetworkAclAttributesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyNetworkAclAttributesCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyNetworkAclAttributesCommon Send returns without error.
//
// See ModifyNetworkAclAttributesCommon for more information on using the ModifyNetworkAclAttributesCommon
// API call, and error handling.
//
//    // Example sending a request using the ModifyNetworkAclAttributesCommonRequest method.
//    req, resp := client.ModifyNetworkAclAttributesCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VPC) ModifyNetworkAclAttributesCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opModifyNetworkAclAttributesCommon,
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

// ModifyNetworkAclAttributesCommon API operation for VPC.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VPC's
// API operation ModifyNetworkAclAttributesCommon for usage and error information.
func (c *VPC) ModifyNetworkAclAttributesCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ModifyNetworkAclAttributesCommonRequest(input)
	return out, req.Send()
}

// ModifyNetworkAclAttributesCommonWithContext is the same as ModifyNetworkAclAttributesCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyNetworkAclAttributesCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPC) ModifyNetworkAclAttributesCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ModifyNetworkAclAttributesCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opModifyNetworkAclAttributes = "ModifyNetworkAclAttributes"

// ModifyNetworkAclAttributesRequest generates a "volcengine/request.Request" representing the
// client's request for the ModifyNetworkAclAttributes operation. The "output" return
// value will be populated with the ModifyNetworkAclAttributesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyNetworkAclAttributesCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyNetworkAclAttributesCommon Send returns without error.
//
// See ModifyNetworkAclAttributes for more information on using the ModifyNetworkAclAttributes
// API call, and error handling.
//
//    // Example sending a request using the ModifyNetworkAclAttributesRequest method.
//    req, resp := client.ModifyNetworkAclAttributesRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VPC) ModifyNetworkAclAttributesRequest(input *ModifyNetworkAclAttributesInput) (req *request.Request, output *ModifyNetworkAclAttributesOutput) {
	op := &request.Operation{
		Name:       opModifyNetworkAclAttributes,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ModifyNetworkAclAttributesInput{}
	}

	output = &ModifyNetworkAclAttributesOutput{}
	req = c.newRequest(op, input, output)

	return
}

// ModifyNetworkAclAttributes API operation for VPC.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VPC's
// API operation ModifyNetworkAclAttributes for usage and error information.
func (c *VPC) ModifyNetworkAclAttributes(input *ModifyNetworkAclAttributesInput) (*ModifyNetworkAclAttributesOutput, error) {
	req, out := c.ModifyNetworkAclAttributesRequest(input)
	return out, req.Send()
}

// ModifyNetworkAclAttributesWithContext is the same as ModifyNetworkAclAttributes with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyNetworkAclAttributes for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPC) ModifyNetworkAclAttributesWithContext(ctx volcengine.Context, input *ModifyNetworkAclAttributesInput, opts ...request.Option) (*ModifyNetworkAclAttributesOutput, error) {
	req, out := c.ModifyNetworkAclAttributesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ModifyNetworkAclAttributesInput struct {
	_ struct{} `type:"structure"`

	Description *string `type:"string"`

	// NetworkAclId is a required field
	NetworkAclId *string `type:"string" required:"true"`

	NetworkAclName *string `type:"string"`
}

// String returns the string representation
func (s ModifyNetworkAclAttributesInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyNetworkAclAttributesInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ModifyNetworkAclAttributesInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ModifyNetworkAclAttributesInput"}
	if s.NetworkAclId == nil {
		invalidParams.Add(request.NewErrParamRequired("NetworkAclId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetDescription sets the Description field's value.
func (s *ModifyNetworkAclAttributesInput) SetDescription(v string) *ModifyNetworkAclAttributesInput {
	s.Description = &v
	return s
}

// SetNetworkAclId sets the NetworkAclId field's value.
func (s *ModifyNetworkAclAttributesInput) SetNetworkAclId(v string) *ModifyNetworkAclAttributesInput {
	s.NetworkAclId = &v
	return s
}

// SetNetworkAclName sets the NetworkAclName field's value.
func (s *ModifyNetworkAclAttributesInput) SetNetworkAclName(v string) *ModifyNetworkAclAttributesInput {
	s.NetworkAclName = &v
	return s
}

type ModifyNetworkAclAttributesOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	RequestId *string `type:"string"`
}

// String returns the string representation
func (s ModifyNetworkAclAttributesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyNetworkAclAttributesOutput) GoString() string {
	return s.String()
}

// SetRequestId sets the RequestId field's value.
func (s *ModifyNetworkAclAttributesOutput) SetRequestId(v string) *ModifyNetworkAclAttributesOutput {
	s.RequestId = &v
	return s
}

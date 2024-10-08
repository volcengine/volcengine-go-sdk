// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package privatelink

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opModifyVpcLinkAttributesCommon = "ModifyVpcLinkAttributes"

// ModifyVpcLinkAttributesCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ModifyVpcLinkAttributesCommon operation. The "output" return
// value will be populated with the ModifyVpcLinkAttributesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyVpcLinkAttributesCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyVpcLinkAttributesCommon Send returns without error.
//
// See ModifyVpcLinkAttributesCommon for more information on using the ModifyVpcLinkAttributesCommon
// API call, and error handling.
//
//    // Example sending a request using the ModifyVpcLinkAttributesCommonRequest method.
//    req, resp := client.ModifyVpcLinkAttributesCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *PRIVATELINK) ModifyVpcLinkAttributesCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opModifyVpcLinkAttributesCommon,
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

// ModifyVpcLinkAttributesCommon API operation for PRIVATELINK.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for PRIVATELINK's
// API operation ModifyVpcLinkAttributesCommon for usage and error information.
func (c *PRIVATELINK) ModifyVpcLinkAttributesCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ModifyVpcLinkAttributesCommonRequest(input)
	return out, req.Send()
}

// ModifyVpcLinkAttributesCommonWithContext is the same as ModifyVpcLinkAttributesCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyVpcLinkAttributesCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *PRIVATELINK) ModifyVpcLinkAttributesCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ModifyVpcLinkAttributesCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opModifyVpcLinkAttributes = "ModifyVpcLinkAttributes"

// ModifyVpcLinkAttributesRequest generates a "volcengine/request.Request" representing the
// client's request for the ModifyVpcLinkAttributes operation. The "output" return
// value will be populated with the ModifyVpcLinkAttributesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyVpcLinkAttributesCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyVpcLinkAttributesCommon Send returns without error.
//
// See ModifyVpcLinkAttributes for more information on using the ModifyVpcLinkAttributes
// API call, and error handling.
//
//    // Example sending a request using the ModifyVpcLinkAttributesRequest method.
//    req, resp := client.ModifyVpcLinkAttributesRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *PRIVATELINK) ModifyVpcLinkAttributesRequest(input *ModifyVpcLinkAttributesInput) (req *request.Request, output *ModifyVpcLinkAttributesOutput) {
	op := &request.Operation{
		Name:       opModifyVpcLinkAttributes,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ModifyVpcLinkAttributesInput{}
	}

	output = &ModifyVpcLinkAttributesOutput{}
	req = c.newRequest(op, input, output)

	return
}

// ModifyVpcLinkAttributes API operation for PRIVATELINK.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for PRIVATELINK's
// API operation ModifyVpcLinkAttributes for usage and error information.
func (c *PRIVATELINK) ModifyVpcLinkAttributes(input *ModifyVpcLinkAttributesInput) (*ModifyVpcLinkAttributesOutput, error) {
	req, out := c.ModifyVpcLinkAttributesRequest(input)
	return out, req.Send()
}

// ModifyVpcLinkAttributesWithContext is the same as ModifyVpcLinkAttributes with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyVpcLinkAttributes for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *PRIVATELINK) ModifyVpcLinkAttributesWithContext(ctx volcengine.Context, input *ModifyVpcLinkAttributesInput, opts ...request.Option) (*ModifyVpcLinkAttributesOutput, error) {
	req, out := c.ModifyVpcLinkAttributesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ModifyVpcLinkAttributesInput struct {
	_ struct{} `type:"structure"`

	Description *string `type:"string"`

	// VpcLinkId is a required field
	VpcLinkId *string `type:"string" required:"true"`

	VpcLinkName *string `type:"string"`
}

// String returns the string representation
func (s ModifyVpcLinkAttributesInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyVpcLinkAttributesInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ModifyVpcLinkAttributesInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ModifyVpcLinkAttributesInput"}
	if s.VpcLinkId == nil {
		invalidParams.Add(request.NewErrParamRequired("VpcLinkId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetDescription sets the Description field's value.
func (s *ModifyVpcLinkAttributesInput) SetDescription(v string) *ModifyVpcLinkAttributesInput {
	s.Description = &v
	return s
}

// SetVpcLinkId sets the VpcLinkId field's value.
func (s *ModifyVpcLinkAttributesInput) SetVpcLinkId(v string) *ModifyVpcLinkAttributesInput {
	s.VpcLinkId = &v
	return s
}

// SetVpcLinkName sets the VpcLinkName field's value.
func (s *ModifyVpcLinkAttributesInput) SetVpcLinkName(v string) *ModifyVpcLinkAttributesInput {
	s.VpcLinkName = &v
	return s
}

type ModifyVpcLinkAttributesOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	RequestId *string `type:"string"`
}

// String returns the string representation
func (s ModifyVpcLinkAttributesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyVpcLinkAttributesOutput) GoString() string {
	return s.String()
}

// SetRequestId sets the RequestId field's value.
func (s *ModifyVpcLinkAttributesOutput) SetRequestId(v string) *ModifyVpcLinkAttributesOutput {
	s.RequestId = &v
	return s
}

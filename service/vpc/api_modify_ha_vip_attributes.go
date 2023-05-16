// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package vpc

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opModifyHaVipAttributesCommon = "ModifyHaVipAttributes"

// ModifyHaVipAttributesCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ModifyHaVipAttributesCommon operation. The "output" return
// value will be populated with the ModifyHaVipAttributesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyHaVipAttributesCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyHaVipAttributesCommon Send returns without error.
//
// See ModifyHaVipAttributesCommon for more information on using the ModifyHaVipAttributesCommon
// API call, and error handling.
//
//	// Example sending a request using the ModifyHaVipAttributesCommonRequest method.
//	req, resp := client.ModifyHaVipAttributesCommonRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *VPC) ModifyHaVipAttributesCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opModifyHaVipAttributesCommon,
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

// ModifyHaVipAttributesCommon API operation for VPC.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VPC's
// API operation ModifyHaVipAttributesCommon for usage and error information.
func (c *VPC) ModifyHaVipAttributesCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ModifyHaVipAttributesCommonRequest(input)
	return out, req.Send()
}

// ModifyHaVipAttributesCommonWithContext is the same as ModifyHaVipAttributesCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyHaVipAttributesCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPC) ModifyHaVipAttributesCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ModifyHaVipAttributesCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opModifyHaVipAttributes = "ModifyHaVipAttributes"

// ModifyHaVipAttributesRequest generates a "volcengine/request.Request" representing the
// client's request for the ModifyHaVipAttributes operation. The "output" return
// value will be populated with the ModifyHaVipAttributesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyHaVipAttributesCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyHaVipAttributesCommon Send returns without error.
//
// See ModifyHaVipAttributes for more information on using the ModifyHaVipAttributes
// API call, and error handling.
//
//	// Example sending a request using the ModifyHaVipAttributesRequest method.
//	req, resp := client.ModifyHaVipAttributesRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *VPC) ModifyHaVipAttributesRequest(input *ModifyHaVipAttributesInput) (req *request.Request, output *ModifyHaVipAttributesOutput) {
	op := &request.Operation{
		Name:       opModifyHaVipAttributes,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ModifyHaVipAttributesInput{}
	}

	output = &ModifyHaVipAttributesOutput{}
	req = c.newRequest(op, input, output)

	return
}

// ModifyHaVipAttributes API operation for VPC.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VPC's
// API operation ModifyHaVipAttributes for usage and error information.
func (c *VPC) ModifyHaVipAttributes(input *ModifyHaVipAttributesInput) (*ModifyHaVipAttributesOutput, error) {
	req, out := c.ModifyHaVipAttributesRequest(input)
	return out, req.Send()
}

// ModifyHaVipAttributesWithContext is the same as ModifyHaVipAttributes with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyHaVipAttributes for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPC) ModifyHaVipAttributesWithContext(ctx volcengine.Context, input *ModifyHaVipAttributesInput, opts ...request.Option) (*ModifyHaVipAttributesOutput, error) {
	req, out := c.ModifyHaVipAttributesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ModifyHaVipAttributesInput struct {
	_ struct{} `type:"structure"`

	Description *string `min:"1" max:"255" type:"string"`

	// HaVipId is a required field
	HaVipId *string `type:"string" required:"true"`

	HaVipName *string `min:"1" max:"128" type:"string"`
}

// String returns the string representation
func (s ModifyHaVipAttributesInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyHaVipAttributesInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ModifyHaVipAttributesInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ModifyHaVipAttributesInput"}
	if s.Description != nil && len(*s.Description) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("Description", 1))
	}
	if s.Description != nil && len(*s.Description) > 255 {
		invalidParams.Add(request.NewErrParamMaxLen("Description", 255, *s.Description))
	}
	if s.HaVipId == nil {
		invalidParams.Add(request.NewErrParamRequired("HaVipId"))
	}
	if s.HaVipName != nil && len(*s.HaVipName) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("HaVipName", 1))
	}
	if s.HaVipName != nil && len(*s.HaVipName) > 128 {
		invalidParams.Add(request.NewErrParamMaxLen("HaVipName", 128, *s.HaVipName))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetDescription sets the Description field's value.
func (s *ModifyHaVipAttributesInput) SetDescription(v string) *ModifyHaVipAttributesInput {
	s.Description = &v
	return s
}

// SetHaVipId sets the HaVipId field's value.
func (s *ModifyHaVipAttributesInput) SetHaVipId(v string) *ModifyHaVipAttributesInput {
	s.HaVipId = &v
	return s
}

// SetHaVipName sets the HaVipName field's value.
func (s *ModifyHaVipAttributesInput) SetHaVipName(v string) *ModifyHaVipAttributesInput {
	s.HaVipName = &v
	return s
}

type ModifyHaVipAttributesOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	RequestId *string `type:"string"`
}

// String returns the string representation
func (s ModifyHaVipAttributesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyHaVipAttributesOutput) GoString() string {
	return s.String()
}

// SetRequestId sets the RequestId field's value.
func (s *ModifyHaVipAttributesOutput) SetRequestId(v string) *ModifyHaVipAttributesOutput {
	s.RequestId = &v
	return s
}

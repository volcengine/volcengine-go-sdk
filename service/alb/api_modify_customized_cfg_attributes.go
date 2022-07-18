// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package alb

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opModifyCustomizedCfgAttributesCommon = "ModifyCustomizedCfgAttributes"

// ModifyCustomizedCfgAttributesCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ModifyCustomizedCfgAttributesCommon operation. The "output" return
// value will be populated with the ModifyCustomizedCfgAttributesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyCustomizedCfgAttributesCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyCustomizedCfgAttributesCommon Send returns without error.
//
// See ModifyCustomizedCfgAttributesCommon for more information on using the ModifyCustomizedCfgAttributesCommon
// API call, and error handling.
//
//    // Example sending a request using the ModifyCustomizedCfgAttributesCommonRequest method.
//    req, resp := client.ModifyCustomizedCfgAttributesCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ALB) ModifyCustomizedCfgAttributesCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opModifyCustomizedCfgAttributesCommon,
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

// ModifyCustomizedCfgAttributesCommon API operation for ALB.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ALB's
// API operation ModifyCustomizedCfgAttributesCommon for usage and error information.
func (c *ALB) ModifyCustomizedCfgAttributesCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ModifyCustomizedCfgAttributesCommonRequest(input)
	return out, req.Send()
}

// ModifyCustomizedCfgAttributesCommonWithContext is the same as ModifyCustomizedCfgAttributesCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyCustomizedCfgAttributesCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ALB) ModifyCustomizedCfgAttributesCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ModifyCustomizedCfgAttributesCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opModifyCustomizedCfgAttributes = "ModifyCustomizedCfgAttributes"

// ModifyCustomizedCfgAttributesRequest generates a "volcengine/request.Request" representing the
// client's request for the ModifyCustomizedCfgAttributes operation. The "output" return
// value will be populated with the ModifyCustomizedCfgAttributesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyCustomizedCfgAttributesCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyCustomizedCfgAttributesCommon Send returns without error.
//
// See ModifyCustomizedCfgAttributes for more information on using the ModifyCustomizedCfgAttributes
// API call, and error handling.
//
//    // Example sending a request using the ModifyCustomizedCfgAttributesRequest method.
//    req, resp := client.ModifyCustomizedCfgAttributesRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ALB) ModifyCustomizedCfgAttributesRequest(input *ModifyCustomizedCfgAttributesInput) (req *request.Request, output *ModifyCustomizedCfgAttributesOutput) {
	op := &request.Operation{
		Name:       opModifyCustomizedCfgAttributes,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ModifyCustomizedCfgAttributesInput{}
	}

	output = &ModifyCustomizedCfgAttributesOutput{}
	req = c.newRequest(op, input, output)

	return
}

// ModifyCustomizedCfgAttributes API operation for ALB.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ALB's
// API operation ModifyCustomizedCfgAttributes for usage and error information.
func (c *ALB) ModifyCustomizedCfgAttributes(input *ModifyCustomizedCfgAttributesInput) (*ModifyCustomizedCfgAttributesOutput, error) {
	req, out := c.ModifyCustomizedCfgAttributesRequest(input)
	return out, req.Send()
}

// ModifyCustomizedCfgAttributesWithContext is the same as ModifyCustomizedCfgAttributes with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyCustomizedCfgAttributes for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ALB) ModifyCustomizedCfgAttributesWithContext(ctx volcengine.Context, input *ModifyCustomizedCfgAttributesInput, opts ...request.Option) (*ModifyCustomizedCfgAttributesOutput, error) {
	req, out := c.ModifyCustomizedCfgAttributesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ModifyCustomizedCfgAttributesInput struct {
	_ struct{} `type:"structure"`

	// CustomizedCfgContent is a required field
	CustomizedCfgContent *string `min:"1" max:"4096" type:"string" required:"true"`

	// CustomizedCfgId is a required field
	CustomizedCfgId *string `type:"string" required:"true"`

	CustomizedCfgName *string `min:"1" max:"128" type:"string"`

	Description *string `min:"1" max:"255" type:"string"`
}

// String returns the string representation
func (s ModifyCustomizedCfgAttributesInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyCustomizedCfgAttributesInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ModifyCustomizedCfgAttributesInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ModifyCustomizedCfgAttributesInput"}
	if s.CustomizedCfgContent == nil {
		invalidParams.Add(request.NewErrParamRequired("CustomizedCfgContent"))
	}
	if s.CustomizedCfgContent != nil && len(*s.CustomizedCfgContent) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("CustomizedCfgContent", 1))
	}
	if s.CustomizedCfgContent != nil && len(*s.CustomizedCfgContent) > 4096 {
		invalidParams.Add(request.NewErrParamMaxLen("CustomizedCfgContent", 4096, *s.CustomizedCfgContent))
	}
	if s.CustomizedCfgId == nil {
		invalidParams.Add(request.NewErrParamRequired("CustomizedCfgId"))
	}
	if s.CustomizedCfgName != nil && len(*s.CustomizedCfgName) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("CustomizedCfgName", 1))
	}
	if s.CustomizedCfgName != nil && len(*s.CustomizedCfgName) > 128 {
		invalidParams.Add(request.NewErrParamMaxLen("CustomizedCfgName", 128, *s.CustomizedCfgName))
	}
	if s.Description != nil && len(*s.Description) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("Description", 1))
	}
	if s.Description != nil && len(*s.Description) > 255 {
		invalidParams.Add(request.NewErrParamMaxLen("Description", 255, *s.Description))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetCustomizedCfgContent sets the CustomizedCfgContent field's value.
func (s *ModifyCustomizedCfgAttributesInput) SetCustomizedCfgContent(v string) *ModifyCustomizedCfgAttributesInput {
	s.CustomizedCfgContent = &v
	return s
}

// SetCustomizedCfgId sets the CustomizedCfgId field's value.
func (s *ModifyCustomizedCfgAttributesInput) SetCustomizedCfgId(v string) *ModifyCustomizedCfgAttributesInput {
	s.CustomizedCfgId = &v
	return s
}

// SetCustomizedCfgName sets the CustomizedCfgName field's value.
func (s *ModifyCustomizedCfgAttributesInput) SetCustomizedCfgName(v string) *ModifyCustomizedCfgAttributesInput {
	s.CustomizedCfgName = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *ModifyCustomizedCfgAttributesInput) SetDescription(v string) *ModifyCustomizedCfgAttributesInput {
	s.Description = &v
	return s
}

type ModifyCustomizedCfgAttributesOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	RequestId *string `type:"string"`
}

// String returns the string representation
func (s ModifyCustomizedCfgAttributesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyCustomizedCfgAttributesOutput) GoString() string {
	return s.String()
}

// SetRequestId sets the RequestId field's value.
func (s *ModifyCustomizedCfgAttributesOutput) SetRequestId(v string) *ModifyCustomizedCfgAttributesOutput {
	s.RequestId = &v
	return s
}

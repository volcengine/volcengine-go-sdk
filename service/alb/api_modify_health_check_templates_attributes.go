// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package alb

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opModifyHealthCheckTemplatesAttributesCommon = "ModifyHealthCheckTemplatesAttributes"

// ModifyHealthCheckTemplatesAttributesCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ModifyHealthCheckTemplatesAttributesCommon operation. The "output" return
// value will be populated with the ModifyHealthCheckTemplatesAttributesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyHealthCheckTemplatesAttributesCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyHealthCheckTemplatesAttributesCommon Send returns without error.
//
// See ModifyHealthCheckTemplatesAttributesCommon for more information on using the ModifyHealthCheckTemplatesAttributesCommon
// API call, and error handling.
//
//    // Example sending a request using the ModifyHealthCheckTemplatesAttributesCommonRequest method.
//    req, resp := client.ModifyHealthCheckTemplatesAttributesCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ALB) ModifyHealthCheckTemplatesAttributesCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opModifyHealthCheckTemplatesAttributesCommon,
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

// ModifyHealthCheckTemplatesAttributesCommon API operation for ALB.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ALB's
// API operation ModifyHealthCheckTemplatesAttributesCommon for usage and error information.
func (c *ALB) ModifyHealthCheckTemplatesAttributesCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ModifyHealthCheckTemplatesAttributesCommonRequest(input)
	return out, req.Send()
}

// ModifyHealthCheckTemplatesAttributesCommonWithContext is the same as ModifyHealthCheckTemplatesAttributesCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyHealthCheckTemplatesAttributesCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ALB) ModifyHealthCheckTemplatesAttributesCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ModifyHealthCheckTemplatesAttributesCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opModifyHealthCheckTemplatesAttributes = "ModifyHealthCheckTemplatesAttributes"

// ModifyHealthCheckTemplatesAttributesRequest generates a "volcengine/request.Request" representing the
// client's request for the ModifyHealthCheckTemplatesAttributes operation. The "output" return
// value will be populated with the ModifyHealthCheckTemplatesAttributesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyHealthCheckTemplatesAttributesCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyHealthCheckTemplatesAttributesCommon Send returns without error.
//
// See ModifyHealthCheckTemplatesAttributes for more information on using the ModifyHealthCheckTemplatesAttributes
// API call, and error handling.
//
//    // Example sending a request using the ModifyHealthCheckTemplatesAttributesRequest method.
//    req, resp := client.ModifyHealthCheckTemplatesAttributesRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ALB) ModifyHealthCheckTemplatesAttributesRequest(input *ModifyHealthCheckTemplatesAttributesInput) (req *request.Request, output *ModifyHealthCheckTemplatesAttributesOutput) {
	op := &request.Operation{
		Name:       opModifyHealthCheckTemplatesAttributes,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ModifyHealthCheckTemplatesAttributesInput{}
	}

	output = &ModifyHealthCheckTemplatesAttributesOutput{}
	req = c.newRequest(op, input, output)

	return
}

// ModifyHealthCheckTemplatesAttributes API operation for ALB.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ALB's
// API operation ModifyHealthCheckTemplatesAttributes for usage and error information.
func (c *ALB) ModifyHealthCheckTemplatesAttributes(input *ModifyHealthCheckTemplatesAttributesInput) (*ModifyHealthCheckTemplatesAttributesOutput, error) {
	req, out := c.ModifyHealthCheckTemplatesAttributesRequest(input)
	return out, req.Send()
}

// ModifyHealthCheckTemplatesAttributesWithContext is the same as ModifyHealthCheckTemplatesAttributes with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyHealthCheckTemplatesAttributes for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ALB) ModifyHealthCheckTemplatesAttributesWithContext(ctx volcengine.Context, input *ModifyHealthCheckTemplatesAttributesInput, opts ...request.Option) (*ModifyHealthCheckTemplatesAttributesOutput, error) {
	req, out := c.ModifyHealthCheckTemplatesAttributesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type HealthCheckTemplateForModifyHealthCheckTemplatesAttributesInput struct {
	_ struct{} `type:"structure"`

	Description *string `type:"string"`

	HealthCheckDomain *string `type:"string"`

	HealthCheckHttpCode *string `type:"string"`

	HealthCheckInterval *int64 `type:"integer"`

	HealthCheckMethod *string `type:"string"`

	HealthCheckProtocol *string `type:"string"`

	// HealthCheckTemplateId is a required field
	HealthCheckTemplateId *string `type:"string" required:"true"`

	// HealthCheckTemplateName is a required field
	HealthCheckTemplateName *string `type:"string" required:"true"`

	HealthCheckTimeout *int64 `type:"integer"`

	HealthCheckURI *string `type:"string"`

	HealthyThreshold *int64 `type:"integer"`

	UnhealthyThreshold *int64 `type:"integer"`
}

// String returns the string representation
func (s HealthCheckTemplateForModifyHealthCheckTemplatesAttributesInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s HealthCheckTemplateForModifyHealthCheckTemplatesAttributesInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *HealthCheckTemplateForModifyHealthCheckTemplatesAttributesInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "HealthCheckTemplateForModifyHealthCheckTemplatesAttributesInput"}
	if s.HealthCheckTemplateId == nil {
		invalidParams.Add(request.NewErrParamRequired("HealthCheckTemplateId"))
	}
	if s.HealthCheckTemplateName == nil {
		invalidParams.Add(request.NewErrParamRequired("HealthCheckTemplateName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetDescription sets the Description field's value.
func (s *HealthCheckTemplateForModifyHealthCheckTemplatesAttributesInput) SetDescription(v string) *HealthCheckTemplateForModifyHealthCheckTemplatesAttributesInput {
	s.Description = &v
	return s
}

// SetHealthCheckDomain sets the HealthCheckDomain field's value.
func (s *HealthCheckTemplateForModifyHealthCheckTemplatesAttributesInput) SetHealthCheckDomain(v string) *HealthCheckTemplateForModifyHealthCheckTemplatesAttributesInput {
	s.HealthCheckDomain = &v
	return s
}

// SetHealthCheckHttpCode sets the HealthCheckHttpCode field's value.
func (s *HealthCheckTemplateForModifyHealthCheckTemplatesAttributesInput) SetHealthCheckHttpCode(v string) *HealthCheckTemplateForModifyHealthCheckTemplatesAttributesInput {
	s.HealthCheckHttpCode = &v
	return s
}

// SetHealthCheckInterval sets the HealthCheckInterval field's value.
func (s *HealthCheckTemplateForModifyHealthCheckTemplatesAttributesInput) SetHealthCheckInterval(v int64) *HealthCheckTemplateForModifyHealthCheckTemplatesAttributesInput {
	s.HealthCheckInterval = &v
	return s
}

// SetHealthCheckMethod sets the HealthCheckMethod field's value.
func (s *HealthCheckTemplateForModifyHealthCheckTemplatesAttributesInput) SetHealthCheckMethod(v string) *HealthCheckTemplateForModifyHealthCheckTemplatesAttributesInput {
	s.HealthCheckMethod = &v
	return s
}

// SetHealthCheckProtocol sets the HealthCheckProtocol field's value.
func (s *HealthCheckTemplateForModifyHealthCheckTemplatesAttributesInput) SetHealthCheckProtocol(v string) *HealthCheckTemplateForModifyHealthCheckTemplatesAttributesInput {
	s.HealthCheckProtocol = &v
	return s
}

// SetHealthCheckTemplateId sets the HealthCheckTemplateId field's value.
func (s *HealthCheckTemplateForModifyHealthCheckTemplatesAttributesInput) SetHealthCheckTemplateId(v string) *HealthCheckTemplateForModifyHealthCheckTemplatesAttributesInput {
	s.HealthCheckTemplateId = &v
	return s
}

// SetHealthCheckTemplateName sets the HealthCheckTemplateName field's value.
func (s *HealthCheckTemplateForModifyHealthCheckTemplatesAttributesInput) SetHealthCheckTemplateName(v string) *HealthCheckTemplateForModifyHealthCheckTemplatesAttributesInput {
	s.HealthCheckTemplateName = &v
	return s
}

// SetHealthCheckTimeout sets the HealthCheckTimeout field's value.
func (s *HealthCheckTemplateForModifyHealthCheckTemplatesAttributesInput) SetHealthCheckTimeout(v int64) *HealthCheckTemplateForModifyHealthCheckTemplatesAttributesInput {
	s.HealthCheckTimeout = &v
	return s
}

// SetHealthCheckURI sets the HealthCheckURI field's value.
func (s *HealthCheckTemplateForModifyHealthCheckTemplatesAttributesInput) SetHealthCheckURI(v string) *HealthCheckTemplateForModifyHealthCheckTemplatesAttributesInput {
	s.HealthCheckURI = &v
	return s
}

// SetHealthyThreshold sets the HealthyThreshold field's value.
func (s *HealthCheckTemplateForModifyHealthCheckTemplatesAttributesInput) SetHealthyThreshold(v int64) *HealthCheckTemplateForModifyHealthCheckTemplatesAttributesInput {
	s.HealthyThreshold = &v
	return s
}

// SetUnhealthyThreshold sets the UnhealthyThreshold field's value.
func (s *HealthCheckTemplateForModifyHealthCheckTemplatesAttributesInput) SetUnhealthyThreshold(v int64) *HealthCheckTemplateForModifyHealthCheckTemplatesAttributesInput {
	s.UnhealthyThreshold = &v
	return s
}

type ModifyHealthCheckTemplatesAttributesInput struct {
	_ struct{} `type:"structure"`

	// HealthCheckTemplates is a required field
	HealthCheckTemplates []*HealthCheckTemplateForModifyHealthCheckTemplatesAttributesInput `type:"list" required:"true"`
}

// String returns the string representation
func (s ModifyHealthCheckTemplatesAttributesInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyHealthCheckTemplatesAttributesInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ModifyHealthCheckTemplatesAttributesInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ModifyHealthCheckTemplatesAttributesInput"}
	if s.HealthCheckTemplates == nil {
		invalidParams.Add(request.NewErrParamRequired("HealthCheckTemplates"))
	}
	if s.HealthCheckTemplates != nil {
		for i, v := range s.HealthCheckTemplates {
			if v == nil {
				continue
			}
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "HealthCheckTemplates", i), err.(request.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetHealthCheckTemplates sets the HealthCheckTemplates field's value.
func (s *ModifyHealthCheckTemplatesAttributesInput) SetHealthCheckTemplates(v []*HealthCheckTemplateForModifyHealthCheckTemplatesAttributesInput) *ModifyHealthCheckTemplatesAttributesInput {
	s.HealthCheckTemplates = v
	return s
}

type ModifyHealthCheckTemplatesAttributesOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	RequestId *string `type:"string"`
}

// String returns the string representation
func (s ModifyHealthCheckTemplatesAttributesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyHealthCheckTemplatesAttributesOutput) GoString() string {
	return s.String()
}

// SetRequestId sets the RequestId field's value.
func (s *ModifyHealthCheckTemplatesAttributesOutput) SetRequestId(v string) *ModifyHealthCheckTemplatesAttributesOutput {
	s.RequestId = &v
	return s
}
// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package rdsmysqlv2

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opModifyParameterTemplateCommon = "ModifyParameterTemplate"

// ModifyParameterTemplateCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ModifyParameterTemplateCommon operation. The "output" return
// value will be populated with the ModifyParameterTemplateCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyParameterTemplateCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyParameterTemplateCommon Send returns without error.
//
// See ModifyParameterTemplateCommon for more information on using the ModifyParameterTemplateCommon
// API call, and error handling.
//
//	// Example sending a request using the ModifyParameterTemplateCommonRequest method.
//	req, resp := client.ModifyParameterTemplateCommonRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *RDSMYSQLV2) ModifyParameterTemplateCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opModifyParameterTemplateCommon,
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

// ModifyParameterTemplateCommon API operation for RDS_MYSQL_V2.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RDS_MYSQL_V2's
// API operation ModifyParameterTemplateCommon for usage and error information.
func (c *RDSMYSQLV2) ModifyParameterTemplateCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ModifyParameterTemplateCommonRequest(input)
	return out, req.Send()
}

// ModifyParameterTemplateCommonWithContext is the same as ModifyParameterTemplateCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyParameterTemplateCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RDSMYSQLV2) ModifyParameterTemplateCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ModifyParameterTemplateCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opModifyParameterTemplate = "ModifyParameterTemplate"

// ModifyParameterTemplateRequest generates a "volcengine/request.Request" representing the
// client's request for the ModifyParameterTemplate operation. The "output" return
// value will be populated with the ModifyParameterTemplateCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyParameterTemplateCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyParameterTemplateCommon Send returns without error.
//
// See ModifyParameterTemplate for more information on using the ModifyParameterTemplate
// API call, and error handling.
//
//	// Example sending a request using the ModifyParameterTemplateRequest method.
//	req, resp := client.ModifyParameterTemplateRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *RDSMYSQLV2) ModifyParameterTemplateRequest(input *ModifyParameterTemplateInput) (req *request.Request, output *ModifyParameterTemplateOutput) {
	op := &request.Operation{
		Name:       opModifyParameterTemplate,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ModifyParameterTemplateInput{}
	}

	output = &ModifyParameterTemplateOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// ModifyParameterTemplate API operation for RDS_MYSQL_V2.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RDS_MYSQL_V2's
// API operation ModifyParameterTemplate for usage and error information.
func (c *RDSMYSQLV2) ModifyParameterTemplate(input *ModifyParameterTemplateInput) (*ModifyParameterTemplateOutput, error) {
	req, out := c.ModifyParameterTemplateRequest(input)
	return out, req.Send()
}

// ModifyParameterTemplateWithContext is the same as ModifyParameterTemplate with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyParameterTemplate for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RDSMYSQLV2) ModifyParameterTemplateWithContext(ctx volcengine.Context, input *ModifyParameterTemplateInput, opts ...request.Option) (*ModifyParameterTemplateOutput, error) {
	req, out := c.ModifyParameterTemplateRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type CustomParamForModifyParameterTemplateInput struct {
	_ struct{} `type:"structure"`

	DefaultValue *string `type:"string"`

	Description *string `type:"string"`

	ExpectValue *string `type:"string"`

	Name *string `type:"string"`

	Restart *bool `type:"boolean"`

	RunningValue *string `type:"string"`

	ValueRange *string `type:"string"`
}

// String returns the string representation
func (s CustomParamForModifyParameterTemplateInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CustomParamForModifyParameterTemplateInput) GoString() string {
	return s.String()
}

// SetDefaultValue sets the DefaultValue field's value.
func (s *CustomParamForModifyParameterTemplateInput) SetDefaultValue(v string) *CustomParamForModifyParameterTemplateInput {
	s.DefaultValue = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *CustomParamForModifyParameterTemplateInput) SetDescription(v string) *CustomParamForModifyParameterTemplateInput {
	s.Description = &v
	return s
}

// SetExpectValue sets the ExpectValue field's value.
func (s *CustomParamForModifyParameterTemplateInput) SetExpectValue(v string) *CustomParamForModifyParameterTemplateInput {
	s.ExpectValue = &v
	return s
}

// SetName sets the Name field's value.
func (s *CustomParamForModifyParameterTemplateInput) SetName(v string) *CustomParamForModifyParameterTemplateInput {
	s.Name = &v
	return s
}

// SetRestart sets the Restart field's value.
func (s *CustomParamForModifyParameterTemplateInput) SetRestart(v bool) *CustomParamForModifyParameterTemplateInput {
	s.Restart = &v
	return s
}

// SetRunningValue sets the RunningValue field's value.
func (s *CustomParamForModifyParameterTemplateInput) SetRunningValue(v string) *CustomParamForModifyParameterTemplateInput {
	s.RunningValue = &v
	return s
}

// SetValueRange sets the ValueRange field's value.
func (s *CustomParamForModifyParameterTemplateInput) SetValueRange(v string) *CustomParamForModifyParameterTemplateInput {
	s.ValueRange = &v
	return s
}

type ModifyParameterTemplateInput struct {
	_ struct{} `type:"structure"`

	CustomParams []*CustomParamForModifyParameterTemplateInput `type:"list"`

	TemplateDesc *string `max:"200" type:"string"`

	// TemplateId is a required field
	TemplateId *string `type:"string" required:"true"`

	TemplateName *string `min:"2" max:"64" type:"string"`

	TemplateParams []*TemplateParamForModifyParameterTemplateInput `type:"list"`
}

// String returns the string representation
func (s ModifyParameterTemplateInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyParameterTemplateInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ModifyParameterTemplateInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ModifyParameterTemplateInput"}
	if s.TemplateDesc != nil && len(*s.TemplateDesc) > 200 {
		invalidParams.Add(request.NewErrParamMaxLen("TemplateDesc", 200, *s.TemplateDesc))
	}
	if s.TemplateId == nil {
		invalidParams.Add(request.NewErrParamRequired("TemplateId"))
	}
	if s.TemplateName != nil && len(*s.TemplateName) < 2 {
		invalidParams.Add(request.NewErrParamMinLen("TemplateName", 2))
	}
	if s.TemplateName != nil && len(*s.TemplateName) > 64 {
		invalidParams.Add(request.NewErrParamMaxLen("TemplateName", 64, *s.TemplateName))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetCustomParams sets the CustomParams field's value.
func (s *ModifyParameterTemplateInput) SetCustomParams(v []*CustomParamForModifyParameterTemplateInput) *ModifyParameterTemplateInput {
	s.CustomParams = v
	return s
}

// SetTemplateDesc sets the TemplateDesc field's value.
func (s *ModifyParameterTemplateInput) SetTemplateDesc(v string) *ModifyParameterTemplateInput {
	s.TemplateDesc = &v
	return s
}

// SetTemplateId sets the TemplateId field's value.
func (s *ModifyParameterTemplateInput) SetTemplateId(v string) *ModifyParameterTemplateInput {
	s.TemplateId = &v
	return s
}

// SetTemplateName sets the TemplateName field's value.
func (s *ModifyParameterTemplateInput) SetTemplateName(v string) *ModifyParameterTemplateInput {
	s.TemplateName = &v
	return s
}

// SetTemplateParams sets the TemplateParams field's value.
func (s *ModifyParameterTemplateInput) SetTemplateParams(v []*TemplateParamForModifyParameterTemplateInput) *ModifyParameterTemplateInput {
	s.TemplateParams = v
	return s
}

type ModifyParameterTemplateOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s ModifyParameterTemplateOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyParameterTemplateOutput) GoString() string {
	return s.String()
}

type TemplateParamForModifyParameterTemplateInput struct {
	_ struct{} `type:"structure"`

	DefaultValue *string `type:"string"`

	Description *string `type:"string"`

	ExpectValue *string `type:"string"`

	Name *string `type:"string"`

	Restart *bool `type:"boolean"`

	RunningValue *string `type:"string"`

	ValueRange *string `type:"string"`
}

// String returns the string representation
func (s TemplateParamForModifyParameterTemplateInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s TemplateParamForModifyParameterTemplateInput) GoString() string {
	return s.String()
}

// SetDefaultValue sets the DefaultValue field's value.
func (s *TemplateParamForModifyParameterTemplateInput) SetDefaultValue(v string) *TemplateParamForModifyParameterTemplateInput {
	s.DefaultValue = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *TemplateParamForModifyParameterTemplateInput) SetDescription(v string) *TemplateParamForModifyParameterTemplateInput {
	s.Description = &v
	return s
}

// SetExpectValue sets the ExpectValue field's value.
func (s *TemplateParamForModifyParameterTemplateInput) SetExpectValue(v string) *TemplateParamForModifyParameterTemplateInput {
	s.ExpectValue = &v
	return s
}

// SetName sets the Name field's value.
func (s *TemplateParamForModifyParameterTemplateInput) SetName(v string) *TemplateParamForModifyParameterTemplateInput {
	s.Name = &v
	return s
}

// SetRestart sets the Restart field's value.
func (s *TemplateParamForModifyParameterTemplateInput) SetRestart(v bool) *TemplateParamForModifyParameterTemplateInput {
	s.Restart = &v
	return s
}

// SetRunningValue sets the RunningValue field's value.
func (s *TemplateParamForModifyParameterTemplateInput) SetRunningValue(v string) *TemplateParamForModifyParameterTemplateInput {
	s.RunningValue = &v
	return s
}

// SetValueRange sets the ValueRange field's value.
func (s *TemplateParamForModifyParameterTemplateInput) SetValueRange(v string) *TemplateParamForModifyParameterTemplateInput {
	s.ValueRange = &v
	return s
}
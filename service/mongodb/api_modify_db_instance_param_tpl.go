// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package mongodb

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opModifyDBInstanceParamTplCommon = "ModifyDBInstanceParamTpl"

// ModifyDBInstanceParamTplCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ModifyDBInstanceParamTplCommon operation. The "output" return
// value will be populated with the ModifyDBInstanceParamTplCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyDBInstanceParamTplCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyDBInstanceParamTplCommon Send returns without error.
//
// See ModifyDBInstanceParamTplCommon for more information on using the ModifyDBInstanceParamTplCommon
// API call, and error handling.
//
//    // Example sending a request using the ModifyDBInstanceParamTplCommonRequest method.
//    req, resp := client.ModifyDBInstanceParamTplCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *MONGODB) ModifyDBInstanceParamTplCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opModifyDBInstanceParamTplCommon,
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

// ModifyDBInstanceParamTplCommon API operation for MONGODB.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for MONGODB's
// API operation ModifyDBInstanceParamTplCommon for usage and error information.
func (c *MONGODB) ModifyDBInstanceParamTplCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ModifyDBInstanceParamTplCommonRequest(input)
	return out, req.Send()
}

// ModifyDBInstanceParamTplCommonWithContext is the same as ModifyDBInstanceParamTplCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyDBInstanceParamTplCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *MONGODB) ModifyDBInstanceParamTplCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ModifyDBInstanceParamTplCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opModifyDBInstanceParamTpl = "ModifyDBInstanceParamTpl"

// ModifyDBInstanceParamTplRequest generates a "volcengine/request.Request" representing the
// client's request for the ModifyDBInstanceParamTpl operation. The "output" return
// value will be populated with the ModifyDBInstanceParamTplCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyDBInstanceParamTplCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyDBInstanceParamTplCommon Send returns without error.
//
// See ModifyDBInstanceParamTpl for more information on using the ModifyDBInstanceParamTpl
// API call, and error handling.
//
//    // Example sending a request using the ModifyDBInstanceParamTplRequest method.
//    req, resp := client.ModifyDBInstanceParamTplRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *MONGODB) ModifyDBInstanceParamTplRequest(input *ModifyDBInstanceParamTplInput) (req *request.Request, output *ModifyDBInstanceParamTplOutput) {
	op := &request.Operation{
		Name:       opModifyDBInstanceParamTpl,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ModifyDBInstanceParamTplInput{}
	}

	output = &ModifyDBInstanceParamTplOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// ModifyDBInstanceParamTpl API operation for MONGODB.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for MONGODB's
// API operation ModifyDBInstanceParamTpl for usage and error information.
func (c *MONGODB) ModifyDBInstanceParamTpl(input *ModifyDBInstanceParamTplInput) (*ModifyDBInstanceParamTplOutput, error) {
	req, out := c.ModifyDBInstanceParamTplRequest(input)
	return out, req.Send()
}

// ModifyDBInstanceParamTplWithContext is the same as ModifyDBInstanceParamTpl with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyDBInstanceParamTpl for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *MONGODB) ModifyDBInstanceParamTplWithContext(ctx volcengine.Context, input *ModifyDBInstanceParamTplInput, opts ...request.Option) (*ModifyDBInstanceParamTplOutput, error) {
	req, out := c.ModifyDBInstanceParamTplRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ModifyDBInstanceParamTplInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Description *string `type:"string" json:",omitempty"`

	Name *string `type:"string" json:",omitempty"`

	Params []*ParamForModifyDBInstanceParamTplInput `type:"list" json:",omitempty"`

	ProjectName *string `type:"string" json:",omitempty"`

	// TemplateId is a required field
	TemplateId *string `type:"string" json:",omitempty" required:"true"`
}

// String returns the string representation
func (s ModifyDBInstanceParamTplInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyDBInstanceParamTplInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ModifyDBInstanceParamTplInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ModifyDBInstanceParamTplInput"}
	if s.TemplateId == nil {
		invalidParams.Add(request.NewErrParamRequired("TemplateId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetDescription sets the Description field's value.
func (s *ModifyDBInstanceParamTplInput) SetDescription(v string) *ModifyDBInstanceParamTplInput {
	s.Description = &v
	return s
}

// SetName sets the Name field's value.
func (s *ModifyDBInstanceParamTplInput) SetName(v string) *ModifyDBInstanceParamTplInput {
	s.Name = &v
	return s
}

// SetParams sets the Params field's value.
func (s *ModifyDBInstanceParamTplInput) SetParams(v []*ParamForModifyDBInstanceParamTplInput) *ModifyDBInstanceParamTplInput {
	s.Params = v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *ModifyDBInstanceParamTplInput) SetProjectName(v string) *ModifyDBInstanceParamTplInput {
	s.ProjectName = &v
	return s
}

// SetTemplateId sets the TemplateId field's value.
func (s *ModifyDBInstanceParamTplInput) SetTemplateId(v string) *ModifyDBInstanceParamTplInput {
	s.TemplateId = &v
	return s
}

type ModifyDBInstanceParamTplOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s ModifyDBInstanceParamTplOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyDBInstanceParamTplOutput) GoString() string {
	return s.String()
}

type ParamForModifyDBInstanceParamTplInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	ParamName *string `type:"string" json:",omitempty"`

	ParamValue *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s ParamForModifyDBInstanceParamTplInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ParamForModifyDBInstanceParamTplInput) GoString() string {
	return s.String()
}

// SetParamName sets the ParamName field's value.
func (s *ParamForModifyDBInstanceParamTplInput) SetParamName(v string) *ParamForModifyDBInstanceParamTplInput {
	s.ParamName = &v
	return s
}

// SetParamValue sets the ParamValue field's value.
func (s *ParamForModifyDBInstanceParamTplInput) SetParamValue(v string) *ParamForModifyDBInstanceParamTplInput {
	s.ParamValue = &v
	return s
}

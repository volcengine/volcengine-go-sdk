// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package waf

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opCreateVulWhiteFieldCommon = "CreateVulWhiteField"

// CreateVulWhiteFieldCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateVulWhiteFieldCommon operation. The "output" return
// value will be populated with the CreateVulWhiteFieldCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateVulWhiteFieldCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateVulWhiteFieldCommon Send returns without error.
//
// See CreateVulWhiteFieldCommon for more information on using the CreateVulWhiteFieldCommon
// API call, and error handling.
//
//    // Example sending a request using the CreateVulWhiteFieldCommonRequest method.
//    req, resp := client.CreateVulWhiteFieldCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *WAF) CreateVulWhiteFieldCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opCreateVulWhiteFieldCommon,
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

// CreateVulWhiteFieldCommon API operation for WAF.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for WAF's
// API operation CreateVulWhiteFieldCommon for usage and error information.
func (c *WAF) CreateVulWhiteFieldCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.CreateVulWhiteFieldCommonRequest(input)
	return out, req.Send()
}

// CreateVulWhiteFieldCommonWithContext is the same as CreateVulWhiteFieldCommon with the addition of
// the ability to pass a context and additional request options.
//
// See CreateVulWhiteFieldCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *WAF) CreateVulWhiteFieldCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.CreateVulWhiteFieldCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opCreateVulWhiteField = "CreateVulWhiteField"

// CreateVulWhiteFieldRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateVulWhiteField operation. The "output" return
// value will be populated with the CreateVulWhiteFieldCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateVulWhiteFieldCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateVulWhiteFieldCommon Send returns without error.
//
// See CreateVulWhiteField for more information on using the CreateVulWhiteField
// API call, and error handling.
//
//    // Example sending a request using the CreateVulWhiteFieldRequest method.
//    req, resp := client.CreateVulWhiteFieldRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *WAF) CreateVulWhiteFieldRequest(input *CreateVulWhiteFieldInput) (req *request.Request, output *CreateVulWhiteFieldOutput) {
	op := &request.Operation{
		Name:       opCreateVulWhiteField,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateVulWhiteFieldInput{}
	}

	output = &CreateVulWhiteFieldOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// CreateVulWhiteField API operation for WAF.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for WAF's
// API operation CreateVulWhiteField for usage and error information.
func (c *WAF) CreateVulWhiteField(input *CreateVulWhiteFieldInput) (*CreateVulWhiteFieldOutput, error) {
	req, out := c.CreateVulWhiteFieldRequest(input)
	return out, req.Send()
}

// CreateVulWhiteFieldWithContext is the same as CreateVulWhiteField with the addition of
// the ability to pass a context and additional request options.
//
// See CreateVulWhiteField for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *WAF) CreateVulWhiteFieldWithContext(ctx volcengine.Context, input *CreateVulWhiteFieldInput, opts ...request.Option) (*CreateVulWhiteFieldOutput, error) {
	req, out := c.CreateVulWhiteFieldRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type CreateVulWhiteFieldInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Enable *int32 `type:"int32" json:",omitempty"`

	FieldArea *string `type:"string" json:",omitempty" enum:"EnumOfFieldAreaForCreateVulWhiteFieldInput"`

	FieldList *string `type:"string" json:",omitempty"`

	Host *string `type:"string" json:",omitempty"`

	Name *string `type:"string" json:",omitempty"`

	ProjectName *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s CreateVulWhiteFieldInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateVulWhiteFieldInput) GoString() string {
	return s.String()
}

// SetEnable sets the Enable field's value.
func (s *CreateVulWhiteFieldInput) SetEnable(v int32) *CreateVulWhiteFieldInput {
	s.Enable = &v
	return s
}

// SetFieldArea sets the FieldArea field's value.
func (s *CreateVulWhiteFieldInput) SetFieldArea(v string) *CreateVulWhiteFieldInput {
	s.FieldArea = &v
	return s
}

// SetFieldList sets the FieldList field's value.
func (s *CreateVulWhiteFieldInput) SetFieldList(v string) *CreateVulWhiteFieldInput {
	s.FieldList = &v
	return s
}

// SetHost sets the Host field's value.
func (s *CreateVulWhiteFieldInput) SetHost(v string) *CreateVulWhiteFieldInput {
	s.Host = &v
	return s
}

// SetName sets the Name field's value.
func (s *CreateVulWhiteFieldInput) SetName(v string) *CreateVulWhiteFieldInput {
	s.Name = &v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *CreateVulWhiteFieldInput) SetProjectName(v string) *CreateVulWhiteFieldInput {
	s.ProjectName = &v
	return s
}

type CreateVulWhiteFieldOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	Id *int32 `type:"int32" json:",omitempty"`
}

// String returns the string representation
func (s CreateVulWhiteFieldOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateVulWhiteFieldOutput) GoString() string {
	return s.String()
}

// SetId sets the Id field's value.
func (s *CreateVulWhiteFieldOutput) SetId(v int32) *CreateVulWhiteFieldOutput {
	s.Id = &v
	return s
}

const (
	// EnumOfFieldAreaForCreateVulWhiteFieldInputArgs is a EnumOfFieldAreaForCreateVulWhiteFieldInput enum value
	EnumOfFieldAreaForCreateVulWhiteFieldInputArgs = "args"

	// EnumOfFieldAreaForCreateVulWhiteFieldInputUrl is a EnumOfFieldAreaForCreateVulWhiteFieldInput enum value
	EnumOfFieldAreaForCreateVulWhiteFieldInputUrl = "url"

	// EnumOfFieldAreaForCreateVulWhiteFieldInputCookies is a EnumOfFieldAreaForCreateVulWhiteFieldInput enum value
	EnumOfFieldAreaForCreateVulWhiteFieldInputCookies = "cookies"

	// EnumOfFieldAreaForCreateVulWhiteFieldInputHeaders is a EnumOfFieldAreaForCreateVulWhiteFieldInput enum value
	EnumOfFieldAreaForCreateVulWhiteFieldInputHeaders = "headers"

	// EnumOfFieldAreaForCreateVulWhiteFieldInputBodydetail is a EnumOfFieldAreaForCreateVulWhiteFieldInput enum value
	EnumOfFieldAreaForCreateVulWhiteFieldInputBodydetail = "bodydetail"
)
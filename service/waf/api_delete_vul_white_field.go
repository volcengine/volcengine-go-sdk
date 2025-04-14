// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package waf

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDeleteVulWhiteFieldCommon = "DeleteVulWhiteField"

// DeleteVulWhiteFieldCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DeleteVulWhiteFieldCommon operation. The "output" return
// value will be populated with the DeleteVulWhiteFieldCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeleteVulWhiteFieldCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeleteVulWhiteFieldCommon Send returns without error.
//
// See DeleteVulWhiteFieldCommon for more information on using the DeleteVulWhiteFieldCommon
// API call, and error handling.
//
//    // Example sending a request using the DeleteVulWhiteFieldCommonRequest method.
//    req, resp := client.DeleteVulWhiteFieldCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *WAF) DeleteVulWhiteFieldCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDeleteVulWhiteFieldCommon,
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

// DeleteVulWhiteFieldCommon API operation for WAF.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for WAF's
// API operation DeleteVulWhiteFieldCommon for usage and error information.
func (c *WAF) DeleteVulWhiteFieldCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DeleteVulWhiteFieldCommonRequest(input)
	return out, req.Send()
}

// DeleteVulWhiteFieldCommonWithContext is the same as DeleteVulWhiteFieldCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteVulWhiteFieldCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *WAF) DeleteVulWhiteFieldCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DeleteVulWhiteFieldCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDeleteVulWhiteField = "DeleteVulWhiteField"

// DeleteVulWhiteFieldRequest generates a "volcengine/request.Request" representing the
// client's request for the DeleteVulWhiteField operation. The "output" return
// value will be populated with the DeleteVulWhiteFieldCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeleteVulWhiteFieldCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeleteVulWhiteFieldCommon Send returns without error.
//
// See DeleteVulWhiteField for more information on using the DeleteVulWhiteField
// API call, and error handling.
//
//    // Example sending a request using the DeleteVulWhiteFieldRequest method.
//    req, resp := client.DeleteVulWhiteFieldRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *WAF) DeleteVulWhiteFieldRequest(input *DeleteVulWhiteFieldInput) (req *request.Request, output *DeleteVulWhiteFieldOutput) {
	op := &request.Operation{
		Name:       opDeleteVulWhiteField,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteVulWhiteFieldInput{}
	}

	output = &DeleteVulWhiteFieldOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// DeleteVulWhiteField API operation for WAF.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for WAF's
// API operation DeleteVulWhiteField for usage and error information.
func (c *WAF) DeleteVulWhiteField(input *DeleteVulWhiteFieldInput) (*DeleteVulWhiteFieldOutput, error) {
	req, out := c.DeleteVulWhiteFieldRequest(input)
	return out, req.Send()
}

// DeleteVulWhiteFieldWithContext is the same as DeleteVulWhiteField with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteVulWhiteField for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *WAF) DeleteVulWhiteFieldWithContext(ctx volcengine.Context, input *DeleteVulWhiteFieldInput, opts ...request.Option) (*DeleteVulWhiteFieldOutput, error) {
	req, out := c.DeleteVulWhiteFieldRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DeleteVulWhiteFieldInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	// ID is a required field
	ID *int32 `type:"int32" json:",omitempty" required:"true"`

	ProjectName *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s DeleteVulWhiteFieldInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteVulWhiteFieldInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteVulWhiteFieldInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DeleteVulWhiteFieldInput"}
	if s.ID == nil {
		invalidParams.Add(request.NewErrParamRequired("ID"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetID sets the ID field's value.
func (s *DeleteVulWhiteFieldInput) SetID(v int32) *DeleteVulWhiteFieldInput {
	s.ID = &v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *DeleteVulWhiteFieldInput) SetProjectName(v string) *DeleteVulWhiteFieldInput {
	s.ProjectName = &v
	return s
}

type DeleteVulWhiteFieldOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s DeleteVulWhiteFieldOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteVulWhiteFieldOutput) GoString() string {
	return s.String()
}

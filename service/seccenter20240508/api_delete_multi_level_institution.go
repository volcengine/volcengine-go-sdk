// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package seccenter20240508

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDeleteMultiLevelInstitutionCommon = "DeleteMultiLevelInstitution"

// DeleteMultiLevelInstitutionCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DeleteMultiLevelInstitutionCommon operation. The "output" return
// value will be populated with the DeleteMultiLevelInstitutionCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeleteMultiLevelInstitutionCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeleteMultiLevelInstitutionCommon Send returns without error.
//
// See DeleteMultiLevelInstitutionCommon for more information on using the DeleteMultiLevelInstitutionCommon
// API call, and error handling.
//
//    // Example sending a request using the DeleteMultiLevelInstitutionCommonRequest method.
//    req, resp := client.DeleteMultiLevelInstitutionCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *SECCENTER20240508) DeleteMultiLevelInstitutionCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDeleteMultiLevelInstitutionCommon,
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

// DeleteMultiLevelInstitutionCommon API operation for SECCENTER20240508.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for SECCENTER20240508's
// API operation DeleteMultiLevelInstitutionCommon for usage and error information.
func (c *SECCENTER20240508) DeleteMultiLevelInstitutionCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DeleteMultiLevelInstitutionCommonRequest(input)
	return out, req.Send()
}

// DeleteMultiLevelInstitutionCommonWithContext is the same as DeleteMultiLevelInstitutionCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteMultiLevelInstitutionCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *SECCENTER20240508) DeleteMultiLevelInstitutionCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DeleteMultiLevelInstitutionCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDeleteMultiLevelInstitution = "DeleteMultiLevelInstitution"

// DeleteMultiLevelInstitutionRequest generates a "volcengine/request.Request" representing the
// client's request for the DeleteMultiLevelInstitution operation. The "output" return
// value will be populated with the DeleteMultiLevelInstitutionCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeleteMultiLevelInstitutionCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeleteMultiLevelInstitutionCommon Send returns without error.
//
// See DeleteMultiLevelInstitution for more information on using the DeleteMultiLevelInstitution
// API call, and error handling.
//
//    // Example sending a request using the DeleteMultiLevelInstitutionRequest method.
//    req, resp := client.DeleteMultiLevelInstitutionRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *SECCENTER20240508) DeleteMultiLevelInstitutionRequest(input *DeleteMultiLevelInstitutionInput) (req *request.Request, output *DeleteMultiLevelInstitutionOutput) {
	op := &request.Operation{
		Name:       opDeleteMultiLevelInstitution,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteMultiLevelInstitutionInput{}
	}

	output = &DeleteMultiLevelInstitutionOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// DeleteMultiLevelInstitution API operation for SECCENTER20240508.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for SECCENTER20240508's
// API operation DeleteMultiLevelInstitution for usage and error information.
func (c *SECCENTER20240508) DeleteMultiLevelInstitution(input *DeleteMultiLevelInstitutionInput) (*DeleteMultiLevelInstitutionOutput, error) {
	req, out := c.DeleteMultiLevelInstitutionRequest(input)
	return out, req.Send()
}

// DeleteMultiLevelInstitutionWithContext is the same as DeleteMultiLevelInstitution with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteMultiLevelInstitution for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *SECCENTER20240508) DeleteMultiLevelInstitutionWithContext(ctx volcengine.Context, input *DeleteMultiLevelInstitutionInput, opts ...request.Option) (*DeleteMultiLevelInstitutionOutput, error) {
	req, out := c.DeleteMultiLevelInstitutionRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DeleteMultiLevelInstitutionInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	ID *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s DeleteMultiLevelInstitutionInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteMultiLevelInstitutionInput) GoString() string {
	return s.String()
}

// SetID sets the ID field's value.
func (s *DeleteMultiLevelInstitutionInput) SetID(v string) *DeleteMultiLevelInstitutionInput {
	s.ID = &v
	return s
}

type DeleteMultiLevelInstitutionOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s DeleteMultiLevelInstitutionOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteMultiLevelInstitutionOutput) GoString() string {
	return s.String()
}

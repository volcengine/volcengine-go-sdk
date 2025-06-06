// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package seccenter20240508

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opUpdateRegularCleanCommon = "UpdateRegularClean"

// UpdateRegularCleanCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the UpdateRegularCleanCommon operation. The "output" return
// value will be populated with the UpdateRegularCleanCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned UpdateRegularCleanCommon Request to send the API call to the service.
// the "output" return value is not valid until after UpdateRegularCleanCommon Send returns without error.
//
// See UpdateRegularCleanCommon for more information on using the UpdateRegularCleanCommon
// API call, and error handling.
//
//    // Example sending a request using the UpdateRegularCleanCommonRequest method.
//    req, resp := client.UpdateRegularCleanCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *SECCENTER20240508) UpdateRegularCleanCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opUpdateRegularCleanCommon,
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

// UpdateRegularCleanCommon API operation for SECCENTER20240508.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for SECCENTER20240508's
// API operation UpdateRegularCleanCommon for usage and error information.
func (c *SECCENTER20240508) UpdateRegularCleanCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.UpdateRegularCleanCommonRequest(input)
	return out, req.Send()
}

// UpdateRegularCleanCommonWithContext is the same as UpdateRegularCleanCommon with the addition of
// the ability to pass a context and additional request options.
//
// See UpdateRegularCleanCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *SECCENTER20240508) UpdateRegularCleanCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.UpdateRegularCleanCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opUpdateRegularClean = "UpdateRegularClean"

// UpdateRegularCleanRequest generates a "volcengine/request.Request" representing the
// client's request for the UpdateRegularClean operation. The "output" return
// value will be populated with the UpdateRegularCleanCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned UpdateRegularCleanCommon Request to send the API call to the service.
// the "output" return value is not valid until after UpdateRegularCleanCommon Send returns without error.
//
// See UpdateRegularClean for more information on using the UpdateRegularClean
// API call, and error handling.
//
//    // Example sending a request using the UpdateRegularCleanRequest method.
//    req, resp := client.UpdateRegularCleanRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *SECCENTER20240508) UpdateRegularCleanRequest(input *UpdateRegularCleanInput) (req *request.Request, output *UpdateRegularCleanOutput) {
	op := &request.Operation{
		Name:       opUpdateRegularClean,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateRegularCleanInput{}
	}

	output = &UpdateRegularCleanOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// UpdateRegularClean API operation for SECCENTER20240508.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for SECCENTER20240508's
// API operation UpdateRegularClean for usage and error information.
func (c *SECCENTER20240508) UpdateRegularClean(input *UpdateRegularCleanInput) (*UpdateRegularCleanOutput, error) {
	req, out := c.UpdateRegularCleanRequest(input)
	return out, req.Send()
}

// UpdateRegularCleanWithContext is the same as UpdateRegularClean with the addition of
// the ability to pass a context and additional request options.
//
// See UpdateRegularClean for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *SECCENTER20240508) UpdateRegularCleanWithContext(ctx volcengine.Context, input *UpdateRegularCleanInput, opts ...request.Option) (*UpdateRegularCleanOutput, error) {
	req, out := c.UpdateRegularCleanRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type UpdateRegularCleanInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	CleanDays *int32 `type:"int32" json:",omitempty"`

	Switch *bool `type:"boolean" json:",omitempty"`
}

// String returns the string representation
func (s UpdateRegularCleanInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s UpdateRegularCleanInput) GoString() string {
	return s.String()
}

// SetCleanDays sets the CleanDays field's value.
func (s *UpdateRegularCleanInput) SetCleanDays(v int32) *UpdateRegularCleanInput {
	s.CleanDays = &v
	return s
}

// SetSwitch sets the Switch field's value.
func (s *UpdateRegularCleanInput) SetSwitch(v bool) *UpdateRegularCleanInput {
	s.Switch = &v
	return s
}

type UpdateRegularCleanOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	Data *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s UpdateRegularCleanOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s UpdateRegularCleanOutput) GoString() string {
	return s.String()
}

// SetData sets the Data field's value.
func (s *UpdateRegularCleanOutput) SetData(v string) *UpdateRegularCleanOutput {
	s.Data = &v
	return s
}

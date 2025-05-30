// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package seccenter20240508

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opRecoverFileCommon = "RecoverFile"

// RecoverFileCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the RecoverFileCommon operation. The "output" return
// value will be populated with the RecoverFileCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned RecoverFileCommon Request to send the API call to the service.
// the "output" return value is not valid until after RecoverFileCommon Send returns without error.
//
// See RecoverFileCommon for more information on using the RecoverFileCommon
// API call, and error handling.
//
//    // Example sending a request using the RecoverFileCommonRequest method.
//    req, resp := client.RecoverFileCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *SECCENTER20240508) RecoverFileCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opRecoverFileCommon,
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

// RecoverFileCommon API operation for SECCENTER20240508.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for SECCENTER20240508's
// API operation RecoverFileCommon for usage and error information.
func (c *SECCENTER20240508) RecoverFileCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.RecoverFileCommonRequest(input)
	return out, req.Send()
}

// RecoverFileCommonWithContext is the same as RecoverFileCommon with the addition of
// the ability to pass a context and additional request options.
//
// See RecoverFileCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *SECCENTER20240508) RecoverFileCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.RecoverFileCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opRecoverFile = "RecoverFile"

// RecoverFileRequest generates a "volcengine/request.Request" representing the
// client's request for the RecoverFile operation. The "output" return
// value will be populated with the RecoverFileCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned RecoverFileCommon Request to send the API call to the service.
// the "output" return value is not valid until after RecoverFileCommon Send returns without error.
//
// See RecoverFile for more information on using the RecoverFile
// API call, and error handling.
//
//    // Example sending a request using the RecoverFileRequest method.
//    req, resp := client.RecoverFileRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *SECCENTER20240508) RecoverFileRequest(input *RecoverFileInput) (req *request.Request, output *RecoverFileOutput) {
	op := &request.Operation{
		Name:       opRecoverFile,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &RecoverFileInput{}
	}

	output = &RecoverFileOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// RecoverFile API operation for SECCENTER20240508.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for SECCENTER20240508's
// API operation RecoverFile for usage and error information.
func (c *SECCENTER20240508) RecoverFile(input *RecoverFileInput) (*RecoverFileOutput, error) {
	req, out := c.RecoverFileRequest(input)
	return out, req.Send()
}

// RecoverFileWithContext is the same as RecoverFile with the addition of
// the ability to pass a context and additional request options.
//
// See RecoverFile for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *SECCENTER20240508) RecoverFileWithContext(ctx volcengine.Context, input *RecoverFileInput, opts ...request.Option) (*RecoverFileOutput, error) {
	req, out := c.RecoverFileRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type RecoverFileInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	FileboxID *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s RecoverFileInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s RecoverFileInput) GoString() string {
	return s.String()
}

// SetFileboxID sets the FileboxID field's value.
func (s *RecoverFileInput) SetFileboxID(v string) *RecoverFileInput {
	s.FileboxID = &v
	return s
}

type RecoverFileOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	Data *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s RecoverFileOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s RecoverFileOutput) GoString() string {
	return s.String()
}

// SetData sets the Data field's value.
func (s *RecoverFileOutput) SetData(v string) *RecoverFileOutput {
	s.Data = &v
	return s
}

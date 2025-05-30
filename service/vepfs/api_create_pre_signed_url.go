// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package vepfs

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opCreatePreSignedUrlCommon = "CreatePreSignedUrl"

// CreatePreSignedUrlCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the CreatePreSignedUrlCommon operation. The "output" return
// value will be populated with the CreatePreSignedUrlCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreatePreSignedUrlCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreatePreSignedUrlCommon Send returns without error.
//
// See CreatePreSignedUrlCommon for more information on using the CreatePreSignedUrlCommon
// API call, and error handling.
//
//    // Example sending a request using the CreatePreSignedUrlCommonRequest method.
//    req, resp := client.CreatePreSignedUrlCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VEPFS) CreatePreSignedUrlCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opCreatePreSignedUrlCommon,
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

// CreatePreSignedUrlCommon API operation for VEPFS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VEPFS's
// API operation CreatePreSignedUrlCommon for usage and error information.
func (c *VEPFS) CreatePreSignedUrlCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.CreatePreSignedUrlCommonRequest(input)
	return out, req.Send()
}

// CreatePreSignedUrlCommonWithContext is the same as CreatePreSignedUrlCommon with the addition of
// the ability to pass a context and additional request options.
//
// See CreatePreSignedUrlCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VEPFS) CreatePreSignedUrlCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.CreatePreSignedUrlCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opCreatePreSignedUrl = "CreatePreSignedUrl"

// CreatePreSignedUrlRequest generates a "volcengine/request.Request" representing the
// client's request for the CreatePreSignedUrl operation. The "output" return
// value will be populated with the CreatePreSignedUrlCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreatePreSignedUrlCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreatePreSignedUrlCommon Send returns without error.
//
// See CreatePreSignedUrl for more information on using the CreatePreSignedUrl
// API call, and error handling.
//
//    // Example sending a request using the CreatePreSignedUrlRequest method.
//    req, resp := client.CreatePreSignedUrlRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VEPFS) CreatePreSignedUrlRequest(input *CreatePreSignedUrlInput) (req *request.Request, output *CreatePreSignedUrlOutput) {
	op := &request.Operation{
		Name:       opCreatePreSignedUrl,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreatePreSignedUrlInput{}
	}

	output = &CreatePreSignedUrlOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// CreatePreSignedUrl API operation for VEPFS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VEPFS's
// API operation CreatePreSignedUrl for usage and error information.
func (c *VEPFS) CreatePreSignedUrl(input *CreatePreSignedUrlInput) (*CreatePreSignedUrlOutput, error) {
	req, out := c.CreatePreSignedUrlRequest(input)
	return out, req.Send()
}

// CreatePreSignedUrlWithContext is the same as CreatePreSignedUrl with the addition of
// the ability to pass a context and additional request options.
//
// See CreatePreSignedUrl for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VEPFS) CreatePreSignedUrlWithContext(ctx volcengine.Context, input *CreatePreSignedUrlInput, opts ...request.Option) (*CreatePreSignedUrlOutput, error) {
	req, out := c.CreatePreSignedUrlRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type CreatePreSignedUrlInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Private *bool `type:"boolean" json:",omitempty"`
}

// String returns the string representation
func (s CreatePreSignedUrlInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreatePreSignedUrlInput) GoString() string {
	return s.String()
}

// SetPrivate sets the Private field's value.
func (s *CreatePreSignedUrlInput) SetPrivate(v bool) *CreatePreSignedUrlInput {
	s.Private = &v
	return s
}

type CreatePreSignedUrlOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	Expires *int32 `type:"int32" json:",omitempty"`

	FileKey *string `type:"string" json:",omitempty"`

	FileUrl *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s CreatePreSignedUrlOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreatePreSignedUrlOutput) GoString() string {
	return s.String()
}

// SetExpires sets the Expires field's value.
func (s *CreatePreSignedUrlOutput) SetExpires(v int32) *CreatePreSignedUrlOutput {
	s.Expires = &v
	return s
}

// SetFileKey sets the FileKey field's value.
func (s *CreatePreSignedUrlOutput) SetFileKey(v string) *CreatePreSignedUrlOutput {
	s.FileKey = &v
	return s
}

// SetFileUrl sets the FileUrl field's value.
func (s *CreatePreSignedUrlOutput) SetFileUrl(v string) *CreatePreSignedUrlOutput {
	s.FileUrl = &v
	return s
}

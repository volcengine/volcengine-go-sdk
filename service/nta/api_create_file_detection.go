// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package nta

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opCreateFileDetectionCommon = "CreateFileDetection"

// CreateFileDetectionCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateFileDetectionCommon operation. The "output" return
// value will be populated with the CreateFileDetectionCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateFileDetectionCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateFileDetectionCommon Send returns without error.
//
// See CreateFileDetectionCommon for more information on using the CreateFileDetectionCommon
// API call, and error handling.
//
//    // Example sending a request using the CreateFileDetectionCommonRequest method.
//    req, resp := client.CreateFileDetectionCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *NTA) CreateFileDetectionCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opCreateFileDetectionCommon,
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

// CreateFileDetectionCommon API operation for NTA.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for NTA's
// API operation CreateFileDetectionCommon for usage and error information.
func (c *NTA) CreateFileDetectionCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.CreateFileDetectionCommonRequest(input)
	return out, req.Send()
}

// CreateFileDetectionCommonWithContext is the same as CreateFileDetectionCommon with the addition of
// the ability to pass a context and additional request options.
//
// See CreateFileDetectionCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *NTA) CreateFileDetectionCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.CreateFileDetectionCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opCreateFileDetection = "CreateFileDetection"

// CreateFileDetectionRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateFileDetection operation. The "output" return
// value will be populated with the CreateFileDetectionCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateFileDetectionCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateFileDetectionCommon Send returns without error.
//
// See CreateFileDetection for more information on using the CreateFileDetection
// API call, and error handling.
//
//    // Example sending a request using the CreateFileDetectionRequest method.
//    req, resp := client.CreateFileDetectionRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *NTA) CreateFileDetectionRequest(input *CreateFileDetectionInput) (req *request.Request, output *CreateFileDetectionOutput) {
	op := &request.Operation{
		Name:       opCreateFileDetection,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateFileDetectionInput{}
	}

	output = &CreateFileDetectionOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// CreateFileDetection API operation for NTA.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for NTA's
// API operation CreateFileDetection for usage and error information.
func (c *NTA) CreateFileDetection(input *CreateFileDetectionInput) (*CreateFileDetectionOutput, error) {
	req, out := c.CreateFileDetectionRequest(input)
	return out, req.Send()
}

// CreateFileDetectionWithContext is the same as CreateFileDetection with the addition of
// the ability to pass a context and additional request options.
//
// See CreateFileDetection for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *NTA) CreateFileDetectionWithContext(ctx volcengine.Context, input *CreateFileDetectionInput, opts ...request.Option) (*CreateFileDetectionOutput, error) {
	req, out := c.CreateFileDetectionRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type CreateFileDetectionInput struct {
	_ struct{} `type:"structure"`

	// SignURL is a required field
	SignURL *string `min:"1" max:"2048" type:"string" required:"true"`
}

// String returns the string representation
func (s CreateFileDetectionInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateFileDetectionInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateFileDetectionInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "CreateFileDetectionInput"}
	if s.SignURL == nil {
		invalidParams.Add(request.NewErrParamRequired("SignURL"))
	}
	if s.SignURL != nil && len(*s.SignURL) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("SignURL", 1))
	}
	if s.SignURL != nil && len(*s.SignURL) > 2048 {
		invalidParams.Add(request.NewErrParamMaxLen("SignURL", 2048, *s.SignURL))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetSignURL sets the SignURL field's value.
func (s *CreateFileDetectionInput) SetSignURL(v string) *CreateFileDetectionInput {
	s.SignURL = &v
	return s
}

type CreateFileDetectionOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	QueryKey *string `type:"string"`
}

// String returns the string representation
func (s CreateFileDetectionOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateFileDetectionOutput) GoString() string {
	return s.String()
}

// SetQueryKey sets the QueryKey field's value.
func (s *CreateFileDetectionOutput) SetQueryKey(v string) *CreateFileDetectionOutput {
	s.QueryKey = &v
	return s
}

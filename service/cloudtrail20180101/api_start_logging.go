// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudtrail20180101

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opStartLoggingCommon = "StartLogging"

// StartLoggingCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the StartLoggingCommon operation. The "output" return
// value will be populated with the StartLoggingCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned StartLoggingCommon Request to send the API call to the service.
// the "output" return value is not valid until after StartLoggingCommon Send returns without error.
//
// See StartLoggingCommon for more information on using the StartLoggingCommon
// API call, and error handling.
//
//    // Example sending a request using the StartLoggingCommonRequest method.
//    req, resp := client.StartLoggingCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CLOUDTRAIL20180101) StartLoggingCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opStartLoggingCommon,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &map[string]interface{}{}
	}

	output = &map[string]interface{}{}
	req = c.newRequest(op, input, output)

	return
}

// StartLoggingCommon API operation for CLOUD_TRAIL20180101.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CLOUD_TRAIL20180101's
// API operation StartLoggingCommon for usage and error information.
func (c *CLOUDTRAIL20180101) StartLoggingCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.StartLoggingCommonRequest(input)
	return out, req.Send()
}

// StartLoggingCommonWithContext is the same as StartLoggingCommon with the addition of
// the ability to pass a context and additional request options.
//
// See StartLoggingCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CLOUDTRAIL20180101) StartLoggingCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.StartLoggingCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opStartLogging = "StartLogging"

// StartLoggingRequest generates a "volcengine/request.Request" representing the
// client's request for the StartLogging operation. The "output" return
// value will be populated with the StartLoggingCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned StartLoggingCommon Request to send the API call to the service.
// the "output" return value is not valid until after StartLoggingCommon Send returns without error.
//
// See StartLogging for more information on using the StartLogging
// API call, and error handling.
//
//    // Example sending a request using the StartLoggingRequest method.
//    req, resp := client.StartLoggingRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CLOUDTRAIL20180101) StartLoggingRequest(input *StartLoggingInput) (req *request.Request, output *StartLoggingOutput) {
	op := &request.Operation{
		Name:       opStartLogging,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &StartLoggingInput{}
	}

	output = &StartLoggingOutput{}
	req = c.newRequest(op, input, output)

	return
}

// StartLogging API operation for CLOUD_TRAIL20180101.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CLOUD_TRAIL20180101's
// API operation StartLogging for usage and error information.
func (c *CLOUDTRAIL20180101) StartLogging(input *StartLoggingInput) (*StartLoggingOutput, error) {
	req, out := c.StartLoggingRequest(input)
	return out, req.Send()
}

// StartLoggingWithContext is the same as StartLogging with the addition of
// the ability to pass a context and additional request options.
//
// See StartLogging for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CLOUDTRAIL20180101) StartLoggingWithContext(ctx volcengine.Context, input *StartLoggingInput, opts ...request.Option) (*StartLoggingOutput, error) {
	req, out := c.StartLoggingRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type StartLoggingInput struct {
	_ struct{} `type:"structure"`

	TrailName *string `type:"string"`
}

// String returns the string representation
func (s StartLoggingInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s StartLoggingInput) GoString() string {
	return s.String()
}

// SetTrailName sets the TrailName field's value.
func (s *StartLoggingInput) SetTrailName(v string) *StartLoggingInput {
	s.TrailName = &v
	return s
}

type StartLoggingOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s StartLoggingOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s StartLoggingOutput) GoString() string {
	return s.String()
}
// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package spark

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opStartApplicationCommon = "startApplication"

// StartApplicationCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the StartApplicationCommon operation. The "output" return
// value will be populated with the StartApplicationCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned StartApplicationCommon Request to send the API call to the service.
// the "output" return value is not valid until after StartApplicationCommon Send returns without error.
//
// See StartApplicationCommon for more information on using the StartApplicationCommon
// API call, and error handling.
//
//    // Example sending a request using the StartApplicationCommonRequest method.
//    req, resp := client.StartApplicationCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *SPARK) StartApplicationCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opStartApplicationCommon,
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

// StartApplicationCommon API operation for SPARK.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for SPARK's
// API operation StartApplicationCommon for usage and error information.
func (c *SPARK) StartApplicationCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.StartApplicationCommonRequest(input)
	return out, req.Send()
}

// StartApplicationCommonWithContext is the same as StartApplicationCommon with the addition of
// the ability to pass a context and additional request options.
//
// See StartApplicationCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *SPARK) StartApplicationCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.StartApplicationCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opStartApplication = "startApplication"

// StartApplicationRequest generates a "volcengine/request.Request" representing the
// client's request for the StartApplication operation. The "output" return
// value will be populated with the StartApplicationCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned StartApplicationCommon Request to send the API call to the service.
// the "output" return value is not valid until after StartApplicationCommon Send returns without error.
//
// See StartApplication for more information on using the StartApplication
// API call, and error handling.
//
//    // Example sending a request using the StartApplicationRequest method.
//    req, resp := client.StartApplicationRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *SPARK) StartApplicationRequest(input *StartApplicationInput) (req *request.Request, output *StartApplicationOutput) {
	op := &request.Operation{
		Name:       opStartApplication,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &StartApplicationInput{}
	}

	output = &StartApplicationOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// StartApplication API operation for SPARK.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for SPARK's
// API operation StartApplication for usage and error information.
func (c *SPARK) StartApplication(input *StartApplicationInput) (*StartApplicationOutput, error) {
	req, out := c.StartApplicationRequest(input)
	return out, req.Send()
}

// StartApplicationWithContext is the same as StartApplication with the addition of
// the ability to pass a context and additional request options.
//
// See StartApplication for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *SPARK) StartApplicationWithContext(ctx volcengine.Context, input *StartApplicationInput, opts ...request.Option) (*StartApplicationOutput, error) {
	req, out := c.StartApplicationRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type StartApplicationInput struct {
	_ struct{} `type:"structure"`

	ApplicationTrn *string `type:"string"`

	Args *string `type:"string"`

	Conf map[string]*string `type:"map"`
}

// String returns the string representation
func (s StartApplicationInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s StartApplicationInput) GoString() string {
	return s.String()
}

// SetApplicationTrn sets the ApplicationTrn field's value.
func (s *StartApplicationInput) SetApplicationTrn(v string) *StartApplicationInput {
	s.ApplicationTrn = &v
	return s
}

// SetArgs sets the Args field's value.
func (s *StartApplicationInput) SetArgs(v string) *StartApplicationInput {
	s.Args = &v
	return s
}

// SetConf sets the Conf field's value.
func (s *StartApplicationInput) SetConf(v map[string]*string) *StartApplicationInput {
	s.Conf = v
	return s
}

type StartApplicationOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	Id *string `type:"string"`

	InstanceId *string `type:"string"`

	Status *string `type:"string"`

	Success *bool `type:"boolean"`
}

// String returns the string representation
func (s StartApplicationOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s StartApplicationOutput) GoString() string {
	return s.String()
}

// SetId sets the Id field's value.
func (s *StartApplicationOutput) SetId(v string) *StartApplicationOutput {
	s.Id = &v
	return s
}

// SetInstanceId sets the InstanceId field's value.
func (s *StartApplicationOutput) SetInstanceId(v string) *StartApplicationOutput {
	s.InstanceId = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *StartApplicationOutput) SetStatus(v string) *StartApplicationOutput {
	s.Status = &v
	return s
}

// SetSuccess sets the Success field's value.
func (s *StartApplicationOutput) SetSuccess(v bool) *StartApplicationOutput {
	s.Success = &v
	return s
}

// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package spark

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDeleteApplicationCommon = "deleteApplication"

// DeleteApplicationCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DeleteApplicationCommon operation. The "output" return
// value will be populated with the DeleteApplicationCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeleteApplicationCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeleteApplicationCommon Send returns without error.
//
// See DeleteApplicationCommon for more information on using the DeleteApplicationCommon
// API call, and error handling.
//
//    // Example sending a request using the DeleteApplicationCommonRequest method.
//    req, resp := client.DeleteApplicationCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *SPARK) DeleteApplicationCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDeleteApplicationCommon,
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

// DeleteApplicationCommon API operation for SPARK.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for SPARK's
// API operation DeleteApplicationCommon for usage and error information.
func (c *SPARK) DeleteApplicationCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DeleteApplicationCommonRequest(input)
	return out, req.Send()
}

// DeleteApplicationCommonWithContext is the same as DeleteApplicationCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteApplicationCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *SPARK) DeleteApplicationCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DeleteApplicationCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDeleteApplication = "deleteApplication"

// DeleteApplicationRequest generates a "volcengine/request.Request" representing the
// client's request for the DeleteApplication operation. The "output" return
// value will be populated with the DeleteApplicationCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeleteApplicationCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeleteApplicationCommon Send returns without error.
//
// See DeleteApplication for more information on using the DeleteApplication
// API call, and error handling.
//
//    // Example sending a request using the DeleteApplicationRequest method.
//    req, resp := client.DeleteApplicationRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *SPARK) DeleteApplicationRequest(input *DeleteApplicationInput) (req *request.Request, output *DeleteApplicationOutput) {
	op := &request.Operation{
		Name:       opDeleteApplication,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteApplicationInput{}
	}

	output = &DeleteApplicationOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// DeleteApplication API operation for SPARK.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for SPARK's
// API operation DeleteApplication for usage and error information.
func (c *SPARK) DeleteApplication(input *DeleteApplicationInput) (*DeleteApplicationOutput, error) {
	req, out := c.DeleteApplicationRequest(input)
	return out, req.Send()
}

// DeleteApplicationWithContext is the same as DeleteApplication with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteApplication for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *SPARK) DeleteApplicationWithContext(ctx volcengine.Context, input *DeleteApplicationInput, opts ...request.Option) (*DeleteApplicationOutput, error) {
	req, out := c.DeleteApplicationRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DeleteApplicationInput struct {
	_ struct{} `type:"structure"`

	ApplicationTrn *string `type:"string"`
}

// String returns the string representation
func (s DeleteApplicationInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteApplicationInput) GoString() string {
	return s.String()
}

// SetApplicationTrn sets the ApplicationTrn field's value.
func (s *DeleteApplicationInput) SetApplicationTrn(v string) *DeleteApplicationInput {
	s.ApplicationTrn = &v
	return s
}

type DeleteApplicationOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	Message *string `type:"string"`

	Status *string `type:"string"`
}

// String returns the string representation
func (s DeleteApplicationOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteApplicationOutput) GoString() string {
	return s.String()
}

// SetMessage sets the Message field's value.
func (s *DeleteApplicationOutput) SetMessage(v string) *DeleteApplicationOutput {
	s.Message = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *DeleteApplicationOutput) SetStatus(v string) *DeleteApplicationOutput {
	s.Status = &v
	return s
}

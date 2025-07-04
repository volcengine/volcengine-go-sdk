// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package clb

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDeleteNLBListenerCommon = "DeleteNLBListener"

// DeleteNLBListenerCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DeleteNLBListenerCommon operation. The "output" return
// value will be populated with the DeleteNLBListenerCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeleteNLBListenerCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeleteNLBListenerCommon Send returns without error.
//
// See DeleteNLBListenerCommon for more information on using the DeleteNLBListenerCommon
// API call, and error handling.
//
//    // Example sending a request using the DeleteNLBListenerCommonRequest method.
//    req, resp := client.DeleteNLBListenerCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CLB) DeleteNLBListenerCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDeleteNLBListenerCommon,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &map[string]interface{}{}
	}

	output = &map[string]interface{}{}
	req = c.newRequest(op, input, output)

	return
}

// DeleteNLBListenerCommon API operation for CLB.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CLB's
// API operation DeleteNLBListenerCommon for usage and error information.
func (c *CLB) DeleteNLBListenerCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DeleteNLBListenerCommonRequest(input)
	return out, req.Send()
}

// DeleteNLBListenerCommonWithContext is the same as DeleteNLBListenerCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteNLBListenerCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CLB) DeleteNLBListenerCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DeleteNLBListenerCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDeleteNLBListener = "DeleteNLBListener"

// DeleteNLBListenerRequest generates a "volcengine/request.Request" representing the
// client's request for the DeleteNLBListener operation. The "output" return
// value will be populated with the DeleteNLBListenerCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeleteNLBListenerCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeleteNLBListenerCommon Send returns without error.
//
// See DeleteNLBListener for more information on using the DeleteNLBListener
// API call, and error handling.
//
//    // Example sending a request using the DeleteNLBListenerRequest method.
//    req, resp := client.DeleteNLBListenerRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CLB) DeleteNLBListenerRequest(input *DeleteNLBListenerInput) (req *request.Request, output *DeleteNLBListenerOutput) {
	op := &request.Operation{
		Name:       opDeleteNLBListener,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteNLBListenerInput{}
	}

	output = &DeleteNLBListenerOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DeleteNLBListener API operation for CLB.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CLB's
// API operation DeleteNLBListener for usage and error information.
func (c *CLB) DeleteNLBListener(input *DeleteNLBListenerInput) (*DeleteNLBListenerOutput, error) {
	req, out := c.DeleteNLBListenerRequest(input)
	return out, req.Send()
}

// DeleteNLBListenerWithContext is the same as DeleteNLBListener with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteNLBListener for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CLB) DeleteNLBListenerWithContext(ctx volcengine.Context, input *DeleteNLBListenerInput, opts ...request.Option) (*DeleteNLBListenerOutput, error) {
	req, out := c.DeleteNLBListenerRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DeleteNLBListenerInput struct {
	_ struct{} `type:"structure"`

	// ListenerId is a required field
	ListenerId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteNLBListenerInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteNLBListenerInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteNLBListenerInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DeleteNLBListenerInput"}
	if s.ListenerId == nil {
		invalidParams.Add(request.NewErrParamRequired("ListenerId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetListenerId sets the ListenerId field's value.
func (s *DeleteNLBListenerInput) SetListenerId(v string) *DeleteNLBListenerInput {
	s.ListenerId = &v
	return s
}

type DeleteNLBListenerOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	RequestId *string `type:"string"`
}

// String returns the string representation
func (s DeleteNLBListenerOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteNLBListenerOutput) GoString() string {
	return s.String()
}

// SetRequestId sets the RequestId field's value.
func (s *DeleteNLBListenerOutput) SetRequestId(v string) *DeleteNLBListenerOutput {
	s.RequestId = &v
	return s
}

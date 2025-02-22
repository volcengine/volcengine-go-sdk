// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package cp

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDeleteComponentStepCommon = "DeleteComponentStep"

// DeleteComponentStepCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DeleteComponentStepCommon operation. The "output" return
// value will be populated with the DeleteComponentStepCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeleteComponentStepCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeleteComponentStepCommon Send returns without error.
//
// See DeleteComponentStepCommon for more information on using the DeleteComponentStepCommon
// API call, and error handling.
//
//    // Example sending a request using the DeleteComponentStepCommonRequest method.
//    req, resp := client.DeleteComponentStepCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CP) DeleteComponentStepCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDeleteComponentStepCommon,
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

// DeleteComponentStepCommon API operation for CP.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CP's
// API operation DeleteComponentStepCommon for usage and error information.
func (c *CP) DeleteComponentStepCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DeleteComponentStepCommonRequest(input)
	return out, req.Send()
}

// DeleteComponentStepCommonWithContext is the same as DeleteComponentStepCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteComponentStepCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CP) DeleteComponentStepCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DeleteComponentStepCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDeleteComponentStep = "DeleteComponentStep"

// DeleteComponentStepRequest generates a "volcengine/request.Request" representing the
// client's request for the DeleteComponentStep operation. The "output" return
// value will be populated with the DeleteComponentStepCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeleteComponentStepCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeleteComponentStepCommon Send returns without error.
//
// See DeleteComponentStep for more information on using the DeleteComponentStep
// API call, and error handling.
//
//    // Example sending a request using the DeleteComponentStepRequest method.
//    req, resp := client.DeleteComponentStepRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CP) DeleteComponentStepRequest(input *DeleteComponentStepInput) (req *request.Request, output *DeleteComponentStepOutput) {
	op := &request.Operation{
		Name:       opDeleteComponentStep,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteComponentStepInput{}
	}

	output = &DeleteComponentStepOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// DeleteComponentStep API operation for CP.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CP's
// API operation DeleteComponentStep for usage and error information.
func (c *CP) DeleteComponentStep(input *DeleteComponentStepInput) (*DeleteComponentStepOutput, error) {
	req, out := c.DeleteComponentStepRequest(input)
	return out, req.Send()
}

// DeleteComponentStepWithContext is the same as DeleteComponentStep with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteComponentStep for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CP) DeleteComponentStepWithContext(ctx volcengine.Context, input *DeleteComponentStepInput, opts ...request.Option) (*DeleteComponentStepOutput, error) {
	req, out := c.DeleteComponentStepRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DeleteComponentStepInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	// Identifier is a required field
	Identifier *string `type:"string" json:",omitempty" required:"true"`
}

// String returns the string representation
func (s DeleteComponentStepInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteComponentStepInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteComponentStepInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DeleteComponentStepInput"}
	if s.Identifier == nil {
		invalidParams.Add(request.NewErrParamRequired("Identifier"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetIdentifier sets the Identifier field's value.
func (s *DeleteComponentStepInput) SetIdentifier(v string) *DeleteComponentStepInput {
	s.Identifier = &v
	return s
}

type DeleteComponentStepOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s DeleteComponentStepOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteComponentStepOutput) GoString() string {
	return s.String()
}

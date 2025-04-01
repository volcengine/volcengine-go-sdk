// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package edx

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDeleteDXPInstanceCommon = "DeleteDXPInstance"

// DeleteDXPInstanceCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DeleteDXPInstanceCommon operation. The "output" return
// value will be populated with the DeleteDXPInstanceCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeleteDXPInstanceCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeleteDXPInstanceCommon Send returns without error.
//
// See DeleteDXPInstanceCommon for more information on using the DeleteDXPInstanceCommon
// API call, and error handling.
//
//    // Example sending a request using the DeleteDXPInstanceCommonRequest method.
//    req, resp := client.DeleteDXPInstanceCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *EDX) DeleteDXPInstanceCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDeleteDXPInstanceCommon,
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

// DeleteDXPInstanceCommon API operation for EDX.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for EDX's
// API operation DeleteDXPInstanceCommon for usage and error information.
func (c *EDX) DeleteDXPInstanceCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DeleteDXPInstanceCommonRequest(input)
	return out, req.Send()
}

// DeleteDXPInstanceCommonWithContext is the same as DeleteDXPInstanceCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteDXPInstanceCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *EDX) DeleteDXPInstanceCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DeleteDXPInstanceCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDeleteDXPInstance = "DeleteDXPInstance"

// DeleteDXPInstanceRequest generates a "volcengine/request.Request" representing the
// client's request for the DeleteDXPInstance operation. The "output" return
// value will be populated with the DeleteDXPInstanceCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeleteDXPInstanceCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeleteDXPInstanceCommon Send returns without error.
//
// See DeleteDXPInstance for more information on using the DeleteDXPInstance
// API call, and error handling.
//
//    // Example sending a request using the DeleteDXPInstanceRequest method.
//    req, resp := client.DeleteDXPInstanceRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *EDX) DeleteDXPInstanceRequest(input *DeleteDXPInstanceInput) (req *request.Request, output *DeleteDXPInstanceOutput) {
	op := &request.Operation{
		Name:       opDeleteDXPInstance,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteDXPInstanceInput{}
	}

	output = &DeleteDXPInstanceOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// DeleteDXPInstance API operation for EDX.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for EDX's
// API operation DeleteDXPInstance for usage and error information.
func (c *EDX) DeleteDXPInstance(input *DeleteDXPInstanceInput) (*DeleteDXPInstanceOutput, error) {
	req, out := c.DeleteDXPInstanceRequest(input)
	return out, req.Send()
}

// DeleteDXPInstanceWithContext is the same as DeleteDXPInstance with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteDXPInstance for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *EDX) DeleteDXPInstanceWithContext(ctx volcengine.Context, input *DeleteDXPInstanceInput, opts ...request.Option) (*DeleteDXPInstanceOutput, error) {
	req, out := c.DeleteDXPInstanceRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DeleteDXPInstanceInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	// Area is a required field
	Area *string `type:"string" json:",omitempty" required:"true"`

	// InstanceId is a required field
	InstanceId *string `type:"string" json:",omitempty" required:"true"`
}

// String returns the string representation
func (s DeleteDXPInstanceInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteDXPInstanceInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteDXPInstanceInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DeleteDXPInstanceInput"}
	if s.Area == nil {
		invalidParams.Add(request.NewErrParamRequired("Area"))
	}
	if s.InstanceId == nil {
		invalidParams.Add(request.NewErrParamRequired("InstanceId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetArea sets the Area field's value.
func (s *DeleteDXPInstanceInput) SetArea(v string) *DeleteDXPInstanceInput {
	s.Area = &v
	return s
}

// SetInstanceId sets the InstanceId field's value.
func (s *DeleteDXPInstanceInput) SetInstanceId(v string) *DeleteDXPInstanceInput {
	s.InstanceId = &v
	return s
}

type DeleteDXPInstanceOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s DeleteDXPInstanceOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteDXPInstanceOutput) GoString() string {
	return s.String()
}

// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package mlplatform

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opStartDevInstanceCommon = "StartDevInstance"

// StartDevInstanceCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the StartDevInstanceCommon operation. The "output" return
// value will be populated with the StartDevInstanceCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned StartDevInstanceCommon Request to send the API call to the service.
// the "output" return value is not valid until after StartDevInstanceCommon Send returns without error.
//
// See StartDevInstanceCommon for more information on using the StartDevInstanceCommon
// API call, and error handling.
//
//	// Example sending a request using the StartDevInstanceCommonRequest method.
//	req, resp := client.StartDevInstanceCommonRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *MLPLATFORM) StartDevInstanceCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opStartDevInstanceCommon,
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

// StartDevInstanceCommon API operation for ML_PLATFORM.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ML_PLATFORM's
// API operation StartDevInstanceCommon for usage and error information.
func (c *MLPLATFORM) StartDevInstanceCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.StartDevInstanceCommonRequest(input)
	return out, req.Send()
}

// StartDevInstanceCommonWithContext is the same as StartDevInstanceCommon with the addition of
// the ability to pass a context and additional request options.
//
// See StartDevInstanceCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *MLPLATFORM) StartDevInstanceCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.StartDevInstanceCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opStartDevInstance = "StartDevInstance"

// StartDevInstanceRequest generates a "volcengine/request.Request" representing the
// client's request for the StartDevInstance operation. The "output" return
// value will be populated with the StartDevInstanceCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned StartDevInstanceCommon Request to send the API call to the service.
// the "output" return value is not valid until after StartDevInstanceCommon Send returns without error.
//
// See StartDevInstance for more information on using the StartDevInstance
// API call, and error handling.
//
//	// Example sending a request using the StartDevInstanceRequest method.
//	req, resp := client.StartDevInstanceRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *MLPLATFORM) StartDevInstanceRequest(input *StartDevInstanceInput) (req *request.Request, output *StartDevInstanceOutput) {
	op := &request.Operation{
		Name:       opStartDevInstance,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &StartDevInstanceInput{}
	}

	output = &StartDevInstanceOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// StartDevInstance API operation for ML_PLATFORM.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ML_PLATFORM's
// API operation StartDevInstance for usage and error information.
func (c *MLPLATFORM) StartDevInstance(input *StartDevInstanceInput) (*StartDevInstanceOutput, error) {
	req, out := c.StartDevInstanceRequest(input)
	return out, req.Send()
}

// StartDevInstanceWithContext is the same as StartDevInstance with the addition of
// the ability to pass a context and additional request options.
//
// See StartDevInstance for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *MLPLATFORM) StartDevInstanceWithContext(ctx volcengine.Context, input *StartDevInstanceInput, opts ...request.Option) (*StartDevInstanceOutput, error) {
	req, out := c.StartDevInstanceRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type StartDevInstanceInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Id *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s StartDevInstanceInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s StartDevInstanceInput) GoString() string {
	return s.String()
}

// SetId sets the Id field's value.
func (s *StartDevInstanceInput) SetId(v string) *StartDevInstanceInput {
	s.Id = &v
	return s
}

type StartDevInstanceOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	Id *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s StartDevInstanceOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s StartDevInstanceOutput) GoString() string {
	return s.String()
}

// SetId sets the Id field's value.
func (s *StartDevInstanceOutput) SetId(v string) *StartDevInstanceOutput {
	s.Id = &v
	return s
}
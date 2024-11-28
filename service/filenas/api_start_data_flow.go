// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package filenas

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opStartDataFlowCommon = "StartDataFlow"

// StartDataFlowCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the StartDataFlowCommon operation. The "output" return
// value will be populated with the StartDataFlowCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned StartDataFlowCommon Request to send the API call to the service.
// the "output" return value is not valid until after StartDataFlowCommon Send returns without error.
//
// See StartDataFlowCommon for more information on using the StartDataFlowCommon
// API call, and error handling.
//
//    // Example sending a request using the StartDataFlowCommonRequest method.
//    req, resp := client.StartDataFlowCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *FILENAS) StartDataFlowCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opStartDataFlowCommon,
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

// StartDataFlowCommon API operation for FILENAS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for FILENAS's
// API operation StartDataFlowCommon for usage and error information.
func (c *FILENAS) StartDataFlowCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.StartDataFlowCommonRequest(input)
	return out, req.Send()
}

// StartDataFlowCommonWithContext is the same as StartDataFlowCommon with the addition of
// the ability to pass a context and additional request options.
//
// See StartDataFlowCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *FILENAS) StartDataFlowCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.StartDataFlowCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opStartDataFlow = "StartDataFlow"

// StartDataFlowRequest generates a "volcengine/request.Request" representing the
// client's request for the StartDataFlow operation. The "output" return
// value will be populated with the StartDataFlowCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned StartDataFlowCommon Request to send the API call to the service.
// the "output" return value is not valid until after StartDataFlowCommon Send returns without error.
//
// See StartDataFlow for more information on using the StartDataFlow
// API call, and error handling.
//
//    // Example sending a request using the StartDataFlowRequest method.
//    req, resp := client.StartDataFlowRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *FILENAS) StartDataFlowRequest(input *StartDataFlowInput) (req *request.Request, output *StartDataFlowOutput) {
	op := &request.Operation{
		Name:       opStartDataFlow,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &StartDataFlowInput{}
	}

	output = &StartDataFlowOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// StartDataFlow API operation for FILENAS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for FILENAS's
// API operation StartDataFlow for usage and error information.
func (c *FILENAS) StartDataFlow(input *StartDataFlowInput) (*StartDataFlowOutput, error) {
	req, out := c.StartDataFlowRequest(input)
	return out, req.Send()
}

// StartDataFlowWithContext is the same as StartDataFlow with the addition of
// the ability to pass a context and additional request options.
//
// See StartDataFlow for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *FILENAS) StartDataFlowWithContext(ctx volcengine.Context, input *StartDataFlowInput, opts ...request.Option) (*StartDataFlowOutput, error) {
	req, out := c.StartDataFlowRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type StartDataFlowInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	// Id is a required field
	Id *string `type:"string" json:",omitempty" required:"true"`

	// Types is a required field
	Types *string `type:"string" json:",omitempty" required:"true"`
}

// String returns the string representation
func (s StartDataFlowInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s StartDataFlowInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StartDataFlowInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "StartDataFlowInput"}
	if s.Id == nil {
		invalidParams.Add(request.NewErrParamRequired("Id"))
	}
	if s.Types == nil {
		invalidParams.Add(request.NewErrParamRequired("Types"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetId sets the Id field's value.
func (s *StartDataFlowInput) SetId(v string) *StartDataFlowInput {
	s.Id = &v
	return s
}

// SetTypes sets the Types field's value.
func (s *StartDataFlowInput) SetTypes(v string) *StartDataFlowInput {
	s.Types = &v
	return s
}

type StartDataFlowOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s StartDataFlowOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s StartDataFlowOutput) GoString() string {
	return s.String()
}

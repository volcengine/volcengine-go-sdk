// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package filenas

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opStopDataFlowCommon = "StopDataFlow"

// StopDataFlowCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the StopDataFlowCommon operation. The "output" return
// value will be populated with the StopDataFlowCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned StopDataFlowCommon Request to send the API call to the service.
// the "output" return value is not valid until after StopDataFlowCommon Send returns without error.
//
// See StopDataFlowCommon for more information on using the StopDataFlowCommon
// API call, and error handling.
//
//    // Example sending a request using the StopDataFlowCommonRequest method.
//    req, resp := client.StopDataFlowCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *FILENAS) StopDataFlowCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opStopDataFlowCommon,
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

// StopDataFlowCommon API operation for FILENAS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for FILENAS's
// API operation StopDataFlowCommon for usage and error information.
func (c *FILENAS) StopDataFlowCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.StopDataFlowCommonRequest(input)
	return out, req.Send()
}

// StopDataFlowCommonWithContext is the same as StopDataFlowCommon with the addition of
// the ability to pass a context and additional request options.
//
// See StopDataFlowCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *FILENAS) StopDataFlowCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.StopDataFlowCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opStopDataFlow = "StopDataFlow"

// StopDataFlowRequest generates a "volcengine/request.Request" representing the
// client's request for the StopDataFlow operation. The "output" return
// value will be populated with the StopDataFlowCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned StopDataFlowCommon Request to send the API call to the service.
// the "output" return value is not valid until after StopDataFlowCommon Send returns without error.
//
// See StopDataFlow for more information on using the StopDataFlow
// API call, and error handling.
//
//    // Example sending a request using the StopDataFlowRequest method.
//    req, resp := client.StopDataFlowRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *FILENAS) StopDataFlowRequest(input *StopDataFlowInput) (req *request.Request, output *StopDataFlowOutput) {
	op := &request.Operation{
		Name:       opStopDataFlow,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &StopDataFlowInput{}
	}

	output = &StopDataFlowOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// StopDataFlow API operation for FILENAS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for FILENAS's
// API operation StopDataFlow for usage and error information.
func (c *FILENAS) StopDataFlow(input *StopDataFlowInput) (*StopDataFlowOutput, error) {
	req, out := c.StopDataFlowRequest(input)
	return out, req.Send()
}

// StopDataFlowWithContext is the same as StopDataFlow with the addition of
// the ability to pass a context and additional request options.
//
// See StopDataFlow for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *FILENAS) StopDataFlowWithContext(ctx volcengine.Context, input *StopDataFlowInput, opts ...request.Option) (*StopDataFlowOutput, error) {
	req, out := c.StopDataFlowRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type StopDataFlowInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	// Id is a required field
	Id *string `type:"string" json:",omitempty" required:"true"`

	// Types is a required field
	Types *string `type:"string" json:",omitempty" required:"true"`
}

// String returns the string representation
func (s StopDataFlowInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s StopDataFlowInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StopDataFlowInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "StopDataFlowInput"}
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
func (s *StopDataFlowInput) SetId(v string) *StopDataFlowInput {
	s.Id = &v
	return s
}

// SetTypes sets the Types field's value.
func (s *StopDataFlowInput) SetTypes(v string) *StopDataFlowInput {
	s.Types = &v
	return s
}

type StopDataFlowOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s StopDataFlowOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s StopDataFlowOutput) GoString() string {
	return s.String()
}

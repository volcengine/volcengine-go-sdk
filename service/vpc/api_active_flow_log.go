// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package vpc

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opActiveFlowLogCommon = "ActiveFlowLog"

// ActiveFlowLogCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ActiveFlowLogCommon operation. The "output" return
// value will be populated with the ActiveFlowLogCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ActiveFlowLogCommon Request to send the API call to the service.
// the "output" return value is not valid until after ActiveFlowLogCommon Send returns without error.
//
// See ActiveFlowLogCommon for more information on using the ActiveFlowLogCommon
// API call, and error handling.
//
//    // Example sending a request using the ActiveFlowLogCommonRequest method.
//    req, resp := client.ActiveFlowLogCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VPC) ActiveFlowLogCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opActiveFlowLogCommon,
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

// ActiveFlowLogCommon API operation for VPC.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VPC's
// API operation ActiveFlowLogCommon for usage and error information.
func (c *VPC) ActiveFlowLogCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ActiveFlowLogCommonRequest(input)
	return out, req.Send()
}

// ActiveFlowLogCommonWithContext is the same as ActiveFlowLogCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ActiveFlowLogCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPC) ActiveFlowLogCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ActiveFlowLogCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opActiveFlowLog = "ActiveFlowLog"

// ActiveFlowLogRequest generates a "volcengine/request.Request" representing the
// client's request for the ActiveFlowLog operation. The "output" return
// value will be populated with the ActiveFlowLogCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ActiveFlowLogCommon Request to send the API call to the service.
// the "output" return value is not valid until after ActiveFlowLogCommon Send returns without error.
//
// See ActiveFlowLog for more information on using the ActiveFlowLog
// API call, and error handling.
//
//    // Example sending a request using the ActiveFlowLogRequest method.
//    req, resp := client.ActiveFlowLogRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VPC) ActiveFlowLogRequest(input *ActiveFlowLogInput) (req *request.Request, output *ActiveFlowLogOutput) {
	op := &request.Operation{
		Name:       opActiveFlowLog,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ActiveFlowLogInput{}
	}

	output = &ActiveFlowLogOutput{}
	req = c.newRequest(op, input, output)

	return
}

// ActiveFlowLog API operation for VPC.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VPC's
// API operation ActiveFlowLog for usage and error information.
func (c *VPC) ActiveFlowLog(input *ActiveFlowLogInput) (*ActiveFlowLogOutput, error) {
	req, out := c.ActiveFlowLogRequest(input)
	return out, req.Send()
}

// ActiveFlowLogWithContext is the same as ActiveFlowLog with the addition of
// the ability to pass a context and additional request options.
//
// See ActiveFlowLog for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPC) ActiveFlowLogWithContext(ctx volcengine.Context, input *ActiveFlowLogInput, opts ...request.Option) (*ActiveFlowLogOutput, error) {
	req, out := c.ActiveFlowLogRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ActiveFlowLogInput struct {
	_ struct{} `type:"structure"`

	// ClientToken is a required field
	ClientToken *string `type:"string" required:"true"`

	// FlowLogId is a required field
	FlowLogId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s ActiveFlowLogInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ActiveFlowLogInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ActiveFlowLogInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ActiveFlowLogInput"}
	if s.ClientToken == nil {
		invalidParams.Add(request.NewErrParamRequired("ClientToken"))
	}
	if s.FlowLogId == nil {
		invalidParams.Add(request.NewErrParamRequired("FlowLogId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetClientToken sets the ClientToken field's value.
func (s *ActiveFlowLogInput) SetClientToken(v string) *ActiveFlowLogInput {
	s.ClientToken = &v
	return s
}

// SetFlowLogId sets the FlowLogId field's value.
func (s *ActiveFlowLogInput) SetFlowLogId(v string) *ActiveFlowLogInput {
	s.FlowLogId = &v
	return s
}

type ActiveFlowLogOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	RequestId *string `type:"string"`
}

// String returns the string representation
func (s ActiveFlowLogOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ActiveFlowLogOutput) GoString() string {
	return s.String()
}

// SetRequestId sets the RequestId field's value.
func (s *ActiveFlowLogOutput) SetRequestId(v string) *ActiveFlowLogOutput {
	s.RequestId = &v
	return s
}
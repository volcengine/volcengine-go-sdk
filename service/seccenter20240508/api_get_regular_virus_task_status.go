// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package seccenter20240508

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opGetRegularVirusTaskStatusCommon = "GetRegularVirusTaskStatus"

// GetRegularVirusTaskStatusCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the GetRegularVirusTaskStatusCommon operation. The "output" return
// value will be populated with the GetRegularVirusTaskStatusCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned GetRegularVirusTaskStatusCommon Request to send the API call to the service.
// the "output" return value is not valid until after GetRegularVirusTaskStatusCommon Send returns without error.
//
// See GetRegularVirusTaskStatusCommon for more information on using the GetRegularVirusTaskStatusCommon
// API call, and error handling.
//
//    // Example sending a request using the GetRegularVirusTaskStatusCommonRequest method.
//    req, resp := client.GetRegularVirusTaskStatusCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *SECCENTER20240508) GetRegularVirusTaskStatusCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opGetRegularVirusTaskStatusCommon,
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

// GetRegularVirusTaskStatusCommon API operation for SECCENTER20240508.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for SECCENTER20240508's
// API operation GetRegularVirusTaskStatusCommon for usage and error information.
func (c *SECCENTER20240508) GetRegularVirusTaskStatusCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.GetRegularVirusTaskStatusCommonRequest(input)
	return out, req.Send()
}

// GetRegularVirusTaskStatusCommonWithContext is the same as GetRegularVirusTaskStatusCommon with the addition of
// the ability to pass a context and additional request options.
//
// See GetRegularVirusTaskStatusCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *SECCENTER20240508) GetRegularVirusTaskStatusCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.GetRegularVirusTaskStatusCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opGetRegularVirusTaskStatus = "GetRegularVirusTaskStatus"

// GetRegularVirusTaskStatusRequest generates a "volcengine/request.Request" representing the
// client's request for the GetRegularVirusTaskStatus operation. The "output" return
// value will be populated with the GetRegularVirusTaskStatusCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned GetRegularVirusTaskStatusCommon Request to send the API call to the service.
// the "output" return value is not valid until after GetRegularVirusTaskStatusCommon Send returns without error.
//
// See GetRegularVirusTaskStatus for more information on using the GetRegularVirusTaskStatus
// API call, and error handling.
//
//    // Example sending a request using the GetRegularVirusTaskStatusRequest method.
//    req, resp := client.GetRegularVirusTaskStatusRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *SECCENTER20240508) GetRegularVirusTaskStatusRequest(input *GetRegularVirusTaskStatusInput) (req *request.Request, output *GetRegularVirusTaskStatusOutput) {
	op := &request.Operation{
		Name:       opGetRegularVirusTaskStatus,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetRegularVirusTaskStatusInput{}
	}

	output = &GetRegularVirusTaskStatusOutput{}
	req = c.newRequest(op, input, output)

	return
}

// GetRegularVirusTaskStatus API operation for SECCENTER20240508.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for SECCENTER20240508's
// API operation GetRegularVirusTaskStatus for usage and error information.
func (c *SECCENTER20240508) GetRegularVirusTaskStatus(input *GetRegularVirusTaskStatusInput) (*GetRegularVirusTaskStatusOutput, error) {
	req, out := c.GetRegularVirusTaskStatusRequest(input)
	return out, req.Send()
}

// GetRegularVirusTaskStatusWithContext is the same as GetRegularVirusTaskStatus with the addition of
// the ability to pass a context and additional request options.
//
// See GetRegularVirusTaskStatus for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *SECCENTER20240508) GetRegularVirusTaskStatusWithContext(ctx volcengine.Context, input *GetRegularVirusTaskStatusInput, opts ...request.Option) (*GetRegularVirusTaskStatusOutput, error) {
	req, out := c.GetRegularVirusTaskStatusRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type GetRegularVirusTaskStatusInput struct {
	_ struct{} `type:"structure"`

	TopGroupID *string `type:"string"`
}

// String returns the string representation
func (s GetRegularVirusTaskStatusInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s GetRegularVirusTaskStatusInput) GoString() string {
	return s.String()
}

// SetTopGroupID sets the TopGroupID field's value.
func (s *GetRegularVirusTaskStatusInput) SetTopGroupID(v string) *GetRegularVirusTaskStatusInput {
	s.TopGroupID = &v
	return s
}

type GetRegularVirusTaskStatusOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	Status *string `type:"string"`
}

// String returns the string representation
func (s GetRegularVirusTaskStatusOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s GetRegularVirusTaskStatusOutput) GoString() string {
	return s.String()
}

// SetStatus sets the Status field's value.
func (s *GetRegularVirusTaskStatusOutput) SetStatus(v string) *GetRegularVirusTaskStatusOutput {
	s.Status = &v
	return s
}

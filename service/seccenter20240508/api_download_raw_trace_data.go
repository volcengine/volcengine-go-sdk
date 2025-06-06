// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package seccenter20240508

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDownloadRawTraceDataCommon = "DownloadRawTraceData"

// DownloadRawTraceDataCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DownloadRawTraceDataCommon operation. The "output" return
// value will be populated with the DownloadRawTraceDataCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DownloadRawTraceDataCommon Request to send the API call to the service.
// the "output" return value is not valid until after DownloadRawTraceDataCommon Send returns without error.
//
// See DownloadRawTraceDataCommon for more information on using the DownloadRawTraceDataCommon
// API call, and error handling.
//
//    // Example sending a request using the DownloadRawTraceDataCommonRequest method.
//    req, resp := client.DownloadRawTraceDataCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *SECCENTER20240508) DownloadRawTraceDataCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDownloadRawTraceDataCommon,
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

// DownloadRawTraceDataCommon API operation for SECCENTER20240508.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for SECCENTER20240508's
// API operation DownloadRawTraceDataCommon for usage and error information.
func (c *SECCENTER20240508) DownloadRawTraceDataCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DownloadRawTraceDataCommonRequest(input)
	return out, req.Send()
}

// DownloadRawTraceDataCommonWithContext is the same as DownloadRawTraceDataCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DownloadRawTraceDataCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *SECCENTER20240508) DownloadRawTraceDataCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DownloadRawTraceDataCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDownloadRawTraceData = "DownloadRawTraceData"

// DownloadRawTraceDataRequest generates a "volcengine/request.Request" representing the
// client's request for the DownloadRawTraceData operation. The "output" return
// value will be populated with the DownloadRawTraceDataCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DownloadRawTraceDataCommon Request to send the API call to the service.
// the "output" return value is not valid until after DownloadRawTraceDataCommon Send returns without error.
//
// See DownloadRawTraceData for more information on using the DownloadRawTraceData
// API call, and error handling.
//
//    // Example sending a request using the DownloadRawTraceDataRequest method.
//    req, resp := client.DownloadRawTraceDataRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *SECCENTER20240508) DownloadRawTraceDataRequest(input *DownloadRawTraceDataInput) (req *request.Request, output *DownloadRawTraceDataOutput) {
	op := &request.Operation{
		Name:       opDownloadRawTraceData,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DownloadRawTraceDataInput{}
	}

	output = &DownloadRawTraceDataOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DownloadRawTraceData API operation for SECCENTER20240508.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for SECCENTER20240508's
// API operation DownloadRawTraceData for usage and error information.
func (c *SECCENTER20240508) DownloadRawTraceData(input *DownloadRawTraceDataInput) (*DownloadRawTraceDataOutput, error) {
	req, out := c.DownloadRawTraceDataRequest(input)
	return out, req.Send()
}

// DownloadRawTraceDataWithContext is the same as DownloadRawTraceData with the addition of
// the ability to pass a context and additional request options.
//
// See DownloadRawTraceData for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *SECCENTER20240508) DownloadRawTraceDataWithContext(ctx volcengine.Context, input *DownloadRawTraceDataInput, opts ...request.Option) (*DownloadRawTraceDataOutput, error) {
	req, out := c.DownloadRawTraceDataRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DownloadRawTraceDataInput struct {
	_ struct{} `type:"structure"`

	TraceID *string `type:"string"`
}

// String returns the string representation
func (s DownloadRawTraceDataInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DownloadRawTraceDataInput) GoString() string {
	return s.String()
}

// SetTraceID sets the TraceID field's value.
func (s *DownloadRawTraceDataInput) SetTraceID(v string) *DownloadRawTraceDataInput {
	s.TraceID = &v
	return s
}

type DownloadRawTraceDataOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	FileName *string `type:"string"`
}

// String returns the string representation
func (s DownloadRawTraceDataOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DownloadRawTraceDataOutput) GoString() string {
	return s.String()
}

// SetFileName sets the FileName field's value.
func (s *DownloadRawTraceDataOutput) SetFileName(v string) *DownloadRawTraceDataOutput {
	s.FileName = &v
	return s
}

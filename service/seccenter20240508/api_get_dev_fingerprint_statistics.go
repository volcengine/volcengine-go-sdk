// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package seccenter20240508

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opGetDevFingerprintStatisticsCommon = "GetDevFingerprintStatistics"

// GetDevFingerprintStatisticsCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the GetDevFingerprintStatisticsCommon operation. The "output" return
// value will be populated with the GetDevFingerprintStatisticsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned GetDevFingerprintStatisticsCommon Request to send the API call to the service.
// the "output" return value is not valid until after GetDevFingerprintStatisticsCommon Send returns without error.
//
// See GetDevFingerprintStatisticsCommon for more information on using the GetDevFingerprintStatisticsCommon
// API call, and error handling.
//
//    // Example sending a request using the GetDevFingerprintStatisticsCommonRequest method.
//    req, resp := client.GetDevFingerprintStatisticsCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *SECCENTER20240508) GetDevFingerprintStatisticsCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opGetDevFingerprintStatisticsCommon,
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

// GetDevFingerprintStatisticsCommon API operation for SECCENTER20240508.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for SECCENTER20240508's
// API operation GetDevFingerprintStatisticsCommon for usage and error information.
func (c *SECCENTER20240508) GetDevFingerprintStatisticsCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.GetDevFingerprintStatisticsCommonRequest(input)
	return out, req.Send()
}

// GetDevFingerprintStatisticsCommonWithContext is the same as GetDevFingerprintStatisticsCommon with the addition of
// the ability to pass a context and additional request options.
//
// See GetDevFingerprintStatisticsCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *SECCENTER20240508) GetDevFingerprintStatisticsCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.GetDevFingerprintStatisticsCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opGetDevFingerprintStatistics = "GetDevFingerprintStatistics"

// GetDevFingerprintStatisticsRequest generates a "volcengine/request.Request" representing the
// client's request for the GetDevFingerprintStatistics operation. The "output" return
// value will be populated with the GetDevFingerprintStatisticsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned GetDevFingerprintStatisticsCommon Request to send the API call to the service.
// the "output" return value is not valid until after GetDevFingerprintStatisticsCommon Send returns without error.
//
// See GetDevFingerprintStatistics for more information on using the GetDevFingerprintStatistics
// API call, and error handling.
//
//    // Example sending a request using the GetDevFingerprintStatisticsRequest method.
//    req, resp := client.GetDevFingerprintStatisticsRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *SECCENTER20240508) GetDevFingerprintStatisticsRequest(input *GetDevFingerprintStatisticsInput) (req *request.Request, output *GetDevFingerprintStatisticsOutput) {
	op := &request.Operation{
		Name:       opGetDevFingerprintStatistics,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetDevFingerprintStatisticsInput{}
	}

	output = &GetDevFingerprintStatisticsOutput{}
	req = c.newRequest(op, input, output)

	return
}

// GetDevFingerprintStatistics API operation for SECCENTER20240508.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for SECCENTER20240508's
// API operation GetDevFingerprintStatistics for usage and error information.
func (c *SECCENTER20240508) GetDevFingerprintStatistics(input *GetDevFingerprintStatisticsInput) (*GetDevFingerprintStatisticsOutput, error) {
	req, out := c.GetDevFingerprintStatisticsRequest(input)
	return out, req.Send()
}

// GetDevFingerprintStatisticsWithContext is the same as GetDevFingerprintStatistics with the addition of
// the ability to pass a context and additional request options.
//
// See GetDevFingerprintStatistics for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *SECCENTER20240508) GetDevFingerprintStatisticsWithContext(ctx volcengine.Context, input *GetDevFingerprintStatisticsInput, opts ...request.Option) (*GetDevFingerprintStatisticsOutput, error) {
	req, out := c.GetDevFingerprintStatisticsRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type GetDevFingerprintStatisticsInput struct {
	_ struct{} `type:"structure"`

	AssetID *string `type:"string"`
}

// String returns the string representation
func (s GetDevFingerprintStatisticsInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s GetDevFingerprintStatisticsInput) GoString() string {
	return s.String()
}

// SetAssetID sets the AssetID field's value.
func (s *GetDevFingerprintStatisticsInput) SetAssetID(v string) *GetDevFingerprintStatisticsInput {
	s.AssetID = &v
	return s
}

type GetDevFingerprintStatisticsOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	Port *int32 `type:"int32"`

	Process *int32 `type:"int32"`

	Software *int32 `type:"int32"`
}

// String returns the string representation
func (s GetDevFingerprintStatisticsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s GetDevFingerprintStatisticsOutput) GoString() string {
	return s.String()
}

// SetPort sets the Port field's value.
func (s *GetDevFingerprintStatisticsOutput) SetPort(v int32) *GetDevFingerprintStatisticsOutput {
	s.Port = &v
	return s
}

// SetProcess sets the Process field's value.
func (s *GetDevFingerprintStatisticsOutput) SetProcess(v int32) *GetDevFingerprintStatisticsOutput {
	s.Process = &v
	return s
}

// SetSoftware sets the Software field's value.
func (s *GetDevFingerprintStatisticsOutput) SetSoftware(v int32) *GetDevFingerprintStatisticsOutput {
	s.Software = &v
	return s
}

// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package nta

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opGetFileDetectionCommon = "GetFileDetection"

// GetFileDetectionCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the GetFileDetectionCommon operation. The "output" return
// value will be populated with the GetFileDetectionCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned GetFileDetectionCommon Request to send the API call to the service.
// the "output" return value is not valid until after GetFileDetectionCommon Send returns without error.
//
// See GetFileDetectionCommon for more information on using the GetFileDetectionCommon
// API call, and error handling.
//
//    // Example sending a request using the GetFileDetectionCommonRequest method.
//    req, resp := client.GetFileDetectionCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *NTA) GetFileDetectionCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opGetFileDetectionCommon,
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

// GetFileDetectionCommon API operation for NTA.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for NTA's
// API operation GetFileDetectionCommon for usage and error information.
func (c *NTA) GetFileDetectionCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.GetFileDetectionCommonRequest(input)
	return out, req.Send()
}

// GetFileDetectionCommonWithContext is the same as GetFileDetectionCommon with the addition of
// the ability to pass a context and additional request options.
//
// See GetFileDetectionCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *NTA) GetFileDetectionCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.GetFileDetectionCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opGetFileDetection = "GetFileDetection"

// GetFileDetectionRequest generates a "volcengine/request.Request" representing the
// client's request for the GetFileDetection operation. The "output" return
// value will be populated with the GetFileDetectionCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned GetFileDetectionCommon Request to send the API call to the service.
// the "output" return value is not valid until after GetFileDetectionCommon Send returns without error.
//
// See GetFileDetection for more information on using the GetFileDetection
// API call, and error handling.
//
//    // Example sending a request using the GetFileDetectionRequest method.
//    req, resp := client.GetFileDetectionRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *NTA) GetFileDetectionRequest(input *GetFileDetectionInput) (req *request.Request, output *GetFileDetectionOutput) {
	op := &request.Operation{
		Name:       opGetFileDetection,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetFileDetectionInput{}
	}

	output = &GetFileDetectionOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// GetFileDetection API operation for NTA.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for NTA's
// API operation GetFileDetection for usage and error information.
func (c *NTA) GetFileDetection(input *GetFileDetectionInput) (*GetFileDetectionOutput, error) {
	req, out := c.GetFileDetectionRequest(input)
	return out, req.Send()
}

// GetFileDetectionWithContext is the same as GetFileDetection with the addition of
// the ability to pass a context and additional request options.
//
// See GetFileDetection for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *NTA) GetFileDetectionWithContext(ctx volcengine.Context, input *GetFileDetectionInput, opts ...request.Option) (*GetFileDetectionOutput, error) {
	req, out := c.GetFileDetectionRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type GetFileDetectionInput struct {
	_ struct{} `type:"structure"`

	// QueryKey is a required field
	QueryKey *string `max:"64" type:"string" required:"true"`
}

// String returns the string representation
func (s GetFileDetectionInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s GetFileDetectionInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetFileDetectionInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "GetFileDetectionInput"}
	if s.QueryKey == nil {
		invalidParams.Add(request.NewErrParamRequired("QueryKey"))
	}
	if s.QueryKey != nil && len(*s.QueryKey) > 64 {
		invalidParams.Add(request.NewErrParamMaxLen("QueryKey", 64, *s.QueryKey))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetQueryKey sets the QueryKey field's value.
func (s *GetFileDetectionInput) SetQueryKey(v string) *GetFileDetectionInput {
	s.QueryKey = &v
	return s
}

type GetFileDetectionOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	FileMd5 *string `type:"string"`

	FileSecType *int32 `type:"int32"`

	FileSize *int64 `type:"int64"`

	Finish *int32 `type:"int32"`

	VirusDesc *string `type:"string"`
}

// String returns the string representation
func (s GetFileDetectionOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s GetFileDetectionOutput) GoString() string {
	return s.String()
}

// SetFileMd5 sets the FileMd5 field's value.
func (s *GetFileDetectionOutput) SetFileMd5(v string) *GetFileDetectionOutput {
	s.FileMd5 = &v
	return s
}

// SetFileSecType sets the FileSecType field's value.
func (s *GetFileDetectionOutput) SetFileSecType(v int32) *GetFileDetectionOutput {
	s.FileSecType = &v
	return s
}

// SetFileSize sets the FileSize field's value.
func (s *GetFileDetectionOutput) SetFileSize(v int64) *GetFileDetectionOutput {
	s.FileSize = &v
	return s
}

// SetFinish sets the Finish field's value.
func (s *GetFileDetectionOutput) SetFinish(v int32) *GetFileDetectionOutput {
	s.Finish = &v
	return s
}

// SetVirusDesc sets the VirusDesc field's value.
func (s *GetFileDetectionOutput) SetVirusDesc(v string) *GetFileDetectionOutput {
	s.VirusDesc = &v
	return s
}

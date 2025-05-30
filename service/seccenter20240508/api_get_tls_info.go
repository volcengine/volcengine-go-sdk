// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package seccenter20240508

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opGetTLSInfoCommon = "GetTLSInfo"

// GetTLSInfoCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the GetTLSInfoCommon operation. The "output" return
// value will be populated with the GetTLSInfoCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned GetTLSInfoCommon Request to send the API call to the service.
// the "output" return value is not valid until after GetTLSInfoCommon Send returns without error.
//
// See GetTLSInfoCommon for more information on using the GetTLSInfoCommon
// API call, and error handling.
//
//    // Example sending a request using the GetTLSInfoCommonRequest method.
//    req, resp := client.GetTLSInfoCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *SECCENTER20240508) GetTLSInfoCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opGetTLSInfoCommon,
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

// GetTLSInfoCommon API operation for SECCENTER20240508.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for SECCENTER20240508's
// API operation GetTLSInfoCommon for usage and error information.
func (c *SECCENTER20240508) GetTLSInfoCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.GetTLSInfoCommonRequest(input)
	return out, req.Send()
}

// GetTLSInfoCommonWithContext is the same as GetTLSInfoCommon with the addition of
// the ability to pass a context and additional request options.
//
// See GetTLSInfoCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *SECCENTER20240508) GetTLSInfoCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.GetTLSInfoCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opGetTLSInfo = "GetTLSInfo"

// GetTLSInfoRequest generates a "volcengine/request.Request" representing the
// client's request for the GetTLSInfo operation. The "output" return
// value will be populated with the GetTLSInfoCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned GetTLSInfoCommon Request to send the API call to the service.
// the "output" return value is not valid until after GetTLSInfoCommon Send returns without error.
//
// See GetTLSInfo for more information on using the GetTLSInfo
// API call, and error handling.
//
//    // Example sending a request using the GetTLSInfoRequest method.
//    req, resp := client.GetTLSInfoRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *SECCENTER20240508) GetTLSInfoRequest(input *GetTLSInfoInput) (req *request.Request, output *GetTLSInfoOutput) {
	op := &request.Operation{
		Name:       opGetTLSInfo,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetTLSInfoInput{}
	}

	output = &GetTLSInfoOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// GetTLSInfo API operation for SECCENTER20240508.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for SECCENTER20240508's
// API operation GetTLSInfo for usage and error information.
func (c *SECCENTER20240508) GetTLSInfo(input *GetTLSInfoInput) (*GetTLSInfoOutput, error) {
	req, out := c.GetTLSInfoRequest(input)
	return out, req.Send()
}

// GetTLSInfoWithContext is the same as GetTLSInfo with the addition of
// the ability to pass a context and additional request options.
//
// See GetTLSInfo for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *SECCENTER20240508) GetTLSInfoWithContext(ctx volcengine.Context, input *GetTLSInfoInput, opts ...request.Option) (*GetTLSInfoOutput, error) {
	req, out := c.GetTLSInfoRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type GetTLSInfoInput struct {
	_ struct{} `type:"structure" json:",omitempty"`
}

// String returns the string representation
func (s GetTLSInfoInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s GetTLSInfoInput) GoString() string {
	return s.String()
}

type GetTLSInfoOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	Authorized *bool `type:"boolean" json:",omitempty"`

	Open *bool `type:"boolean" json:",omitempty"`

	ProjectId *string `type:"string" json:",omitempty"`

	QuotaTotal *int32 `type:"int32" json:",omitempty"`

	QuotaUsed *int32 `type:"int32" json:",omitempty"`

	StorageDays *int32 `type:"int32" json:",omitempty"`

	Threshold *int32 `type:"int32" json:",omitempty"`

	TopicId *string `type:"string" json:",omitempty"`

	VulnTopicId *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s GetTLSInfoOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s GetTLSInfoOutput) GoString() string {
	return s.String()
}

// SetAuthorized sets the Authorized field's value.
func (s *GetTLSInfoOutput) SetAuthorized(v bool) *GetTLSInfoOutput {
	s.Authorized = &v
	return s
}

// SetOpen sets the Open field's value.
func (s *GetTLSInfoOutput) SetOpen(v bool) *GetTLSInfoOutput {
	s.Open = &v
	return s
}

// SetProjectId sets the ProjectId field's value.
func (s *GetTLSInfoOutput) SetProjectId(v string) *GetTLSInfoOutput {
	s.ProjectId = &v
	return s
}

// SetQuotaTotal sets the QuotaTotal field's value.
func (s *GetTLSInfoOutput) SetQuotaTotal(v int32) *GetTLSInfoOutput {
	s.QuotaTotal = &v
	return s
}

// SetQuotaUsed sets the QuotaUsed field's value.
func (s *GetTLSInfoOutput) SetQuotaUsed(v int32) *GetTLSInfoOutput {
	s.QuotaUsed = &v
	return s
}

// SetStorageDays sets the StorageDays field's value.
func (s *GetTLSInfoOutput) SetStorageDays(v int32) *GetTLSInfoOutput {
	s.StorageDays = &v
	return s
}

// SetThreshold sets the Threshold field's value.
func (s *GetTLSInfoOutput) SetThreshold(v int32) *GetTLSInfoOutput {
	s.Threshold = &v
	return s
}

// SetTopicId sets the TopicId field's value.
func (s *GetTLSInfoOutput) SetTopicId(v string) *GetTLSInfoOutput {
	s.TopicId = &v
	return s
}

// SetVulnTopicId sets the VulnTopicId field's value.
func (s *GetTLSInfoOutput) SetVulnTopicId(v string) *GetTLSInfoOutput {
	s.VulnTopicId = &v
	return s
}

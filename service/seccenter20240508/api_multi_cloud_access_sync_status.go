// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package seccenter20240508

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opMultiCloudAccessSyncStatusCommon = "MultiCloudAccessSyncStatus"

// MultiCloudAccessSyncStatusCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the MultiCloudAccessSyncStatusCommon operation. The "output" return
// value will be populated with the MultiCloudAccessSyncStatusCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned MultiCloudAccessSyncStatusCommon Request to send the API call to the service.
// the "output" return value is not valid until after MultiCloudAccessSyncStatusCommon Send returns without error.
//
// See MultiCloudAccessSyncStatusCommon for more information on using the MultiCloudAccessSyncStatusCommon
// API call, and error handling.
//
//    // Example sending a request using the MultiCloudAccessSyncStatusCommonRequest method.
//    req, resp := client.MultiCloudAccessSyncStatusCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *SECCENTER20240508) MultiCloudAccessSyncStatusCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opMultiCloudAccessSyncStatusCommon,
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

// MultiCloudAccessSyncStatusCommon API operation for SECCENTER20240508.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for SECCENTER20240508's
// API operation MultiCloudAccessSyncStatusCommon for usage and error information.
func (c *SECCENTER20240508) MultiCloudAccessSyncStatusCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.MultiCloudAccessSyncStatusCommonRequest(input)
	return out, req.Send()
}

// MultiCloudAccessSyncStatusCommonWithContext is the same as MultiCloudAccessSyncStatusCommon with the addition of
// the ability to pass a context and additional request options.
//
// See MultiCloudAccessSyncStatusCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *SECCENTER20240508) MultiCloudAccessSyncStatusCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.MultiCloudAccessSyncStatusCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opMultiCloudAccessSyncStatus = "MultiCloudAccessSyncStatus"

// MultiCloudAccessSyncStatusRequest generates a "volcengine/request.Request" representing the
// client's request for the MultiCloudAccessSyncStatus operation. The "output" return
// value will be populated with the MultiCloudAccessSyncStatusCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned MultiCloudAccessSyncStatusCommon Request to send the API call to the service.
// the "output" return value is not valid until after MultiCloudAccessSyncStatusCommon Send returns without error.
//
// See MultiCloudAccessSyncStatus for more information on using the MultiCloudAccessSyncStatus
// API call, and error handling.
//
//    // Example sending a request using the MultiCloudAccessSyncStatusRequest method.
//    req, resp := client.MultiCloudAccessSyncStatusRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *SECCENTER20240508) MultiCloudAccessSyncStatusRequest(input *MultiCloudAccessSyncStatusInput) (req *request.Request, output *MultiCloudAccessSyncStatusOutput) {
	op := &request.Operation{
		Name:       opMultiCloudAccessSyncStatus,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &MultiCloudAccessSyncStatusInput{}
	}

	output = &MultiCloudAccessSyncStatusOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// MultiCloudAccessSyncStatus API operation for SECCENTER20240508.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for SECCENTER20240508's
// API operation MultiCloudAccessSyncStatus for usage and error information.
func (c *SECCENTER20240508) MultiCloudAccessSyncStatus(input *MultiCloudAccessSyncStatusInput) (*MultiCloudAccessSyncStatusOutput, error) {
	req, out := c.MultiCloudAccessSyncStatusRequest(input)
	return out, req.Send()
}

// MultiCloudAccessSyncStatusWithContext is the same as MultiCloudAccessSyncStatus with the addition of
// the ability to pass a context and additional request options.
//
// See MultiCloudAccessSyncStatus for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *SECCENTER20240508) MultiCloudAccessSyncStatusWithContext(ctx volcengine.Context, input *MultiCloudAccessSyncStatusInput, opts ...request.Option) (*MultiCloudAccessSyncStatusOutput, error) {
	req, out := c.MultiCloudAccessSyncStatusRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type MultiCloudAccessSyncStatusInput struct {
	_ struct{} `type:"structure" json:",omitempty"`
}

// String returns the string representation
func (s MultiCloudAccessSyncStatusInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s MultiCloudAccessSyncStatusInput) GoString() string {
	return s.String()
}

type MultiCloudAccessSyncStatusOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	Status *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s MultiCloudAccessSyncStatusOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s MultiCloudAccessSyncStatusOutput) GoString() string {
	return s.String()
}

// SetStatus sets the Status field's value.
func (s *MultiCloudAccessSyncStatusOutput) SetStatus(v string) *MultiCloudAccessSyncStatusOutput {
	s.Status = &v
	return s
}

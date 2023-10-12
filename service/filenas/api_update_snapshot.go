// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package filenas

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opUpdateSnapshotCommon = "UpdateSnapshot"

// UpdateSnapshotCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the UpdateSnapshotCommon operation. The "output" return
// value will be populated with the UpdateSnapshotCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned UpdateSnapshotCommon Request to send the API call to the service.
// the "output" return value is not valid until after UpdateSnapshotCommon Send returns without error.
//
// See UpdateSnapshotCommon for more information on using the UpdateSnapshotCommon
// API call, and error handling.
//
//    // Example sending a request using the UpdateSnapshotCommonRequest method.
//    req, resp := client.UpdateSnapshotCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *FILENAS) UpdateSnapshotCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opUpdateSnapshotCommon,
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

// UpdateSnapshotCommon API operation for FILENAS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for FILENAS's
// API operation UpdateSnapshotCommon for usage and error information.
func (c *FILENAS) UpdateSnapshotCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.UpdateSnapshotCommonRequest(input)
	return out, req.Send()
}

// UpdateSnapshotCommonWithContext is the same as UpdateSnapshotCommon with the addition of
// the ability to pass a context and additional request options.
//
// See UpdateSnapshotCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *FILENAS) UpdateSnapshotCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.UpdateSnapshotCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opUpdateSnapshot = "UpdateSnapshot"

// UpdateSnapshotRequest generates a "volcengine/request.Request" representing the
// client's request for the UpdateSnapshot operation. The "output" return
// value will be populated with the UpdateSnapshotCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned UpdateSnapshotCommon Request to send the API call to the service.
// the "output" return value is not valid until after UpdateSnapshotCommon Send returns without error.
//
// See UpdateSnapshot for more information on using the UpdateSnapshot
// API call, and error handling.
//
//    // Example sending a request using the UpdateSnapshotRequest method.
//    req, resp := client.UpdateSnapshotRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *FILENAS) UpdateSnapshotRequest(input *UpdateSnapshotInput) (req *request.Request, output *UpdateSnapshotOutput) {
	op := &request.Operation{
		Name:       opUpdateSnapshot,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateSnapshotInput{}
	}

	output = &UpdateSnapshotOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// UpdateSnapshot API operation for FILENAS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for FILENAS's
// API operation UpdateSnapshot for usage and error information.
func (c *FILENAS) UpdateSnapshot(input *UpdateSnapshotInput) (*UpdateSnapshotOutput, error) {
	req, out := c.UpdateSnapshotRequest(input)
	return out, req.Send()
}

// UpdateSnapshotWithContext is the same as UpdateSnapshot with the addition of
// the ability to pass a context and additional request options.
//
// See UpdateSnapshot for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *FILENAS) UpdateSnapshotWithContext(ctx volcengine.Context, input *UpdateSnapshotInput, opts ...request.Option) (*UpdateSnapshotOutput, error) {
	req, out := c.UpdateSnapshotRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type UpdateSnapshotInput struct {
	_ struct{} `type:"structure"`

	Description *string `type:"string"`

	SnapshotId *string `type:"string"`

	SnapshotName *string `type:"string"`
}

// String returns the string representation
func (s UpdateSnapshotInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s UpdateSnapshotInput) GoString() string {
	return s.String()
}

// SetDescription sets the Description field's value.
func (s *UpdateSnapshotInput) SetDescription(v string) *UpdateSnapshotInput {
	s.Description = &v
	return s
}

// SetSnapshotId sets the SnapshotId field's value.
func (s *UpdateSnapshotInput) SetSnapshotId(v string) *UpdateSnapshotInput {
	s.SnapshotId = &v
	return s
}

// SetSnapshotName sets the SnapshotName field's value.
func (s *UpdateSnapshotInput) SetSnapshotName(v string) *UpdateSnapshotInput {
	s.SnapshotName = &v
	return s
}

type UpdateSnapshotOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s UpdateSnapshotOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s UpdateSnapshotOutput) GoString() string {
	return s.String()
}

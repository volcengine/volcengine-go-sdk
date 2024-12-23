// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package mcdn

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opSubmitRefreshTaskCommon = "SubmitRefreshTask"

// SubmitRefreshTaskCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the SubmitRefreshTaskCommon operation. The "output" return
// value will be populated with the SubmitRefreshTaskCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned SubmitRefreshTaskCommon Request to send the API call to the service.
// the "output" return value is not valid until after SubmitRefreshTaskCommon Send returns without error.
//
// See SubmitRefreshTaskCommon for more information on using the SubmitRefreshTaskCommon
// API call, and error handling.
//
//    // Example sending a request using the SubmitRefreshTaskCommonRequest method.
//    req, resp := client.SubmitRefreshTaskCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *MCDN) SubmitRefreshTaskCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opSubmitRefreshTaskCommon,
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

// SubmitRefreshTaskCommon API operation for MCDN.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for MCDN's
// API operation SubmitRefreshTaskCommon for usage and error information.
func (c *MCDN) SubmitRefreshTaskCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.SubmitRefreshTaskCommonRequest(input)
	return out, req.Send()
}

// SubmitRefreshTaskCommonWithContext is the same as SubmitRefreshTaskCommon with the addition of
// the ability to pass a context and additional request options.
//
// See SubmitRefreshTaskCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *MCDN) SubmitRefreshTaskCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.SubmitRefreshTaskCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opSubmitRefreshTask = "SubmitRefreshTask"

// SubmitRefreshTaskRequest generates a "volcengine/request.Request" representing the
// client's request for the SubmitRefreshTask operation. The "output" return
// value will be populated with the SubmitRefreshTaskCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned SubmitRefreshTaskCommon Request to send the API call to the service.
// the "output" return value is not valid until after SubmitRefreshTaskCommon Send returns without error.
//
// See SubmitRefreshTask for more information on using the SubmitRefreshTask
// API call, and error handling.
//
//    // Example sending a request using the SubmitRefreshTaskRequest method.
//    req, resp := client.SubmitRefreshTaskRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *MCDN) SubmitRefreshTaskRequest(input *SubmitRefreshTaskInput) (req *request.Request, output *SubmitRefreshTaskOutput) {
	op := &request.Operation{
		Name:       opSubmitRefreshTask,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &SubmitRefreshTaskInput{}
	}

	output = &SubmitRefreshTaskOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// SubmitRefreshTask API operation for MCDN.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for MCDN's
// API operation SubmitRefreshTask for usage and error information.
func (c *MCDN) SubmitRefreshTask(input *SubmitRefreshTaskInput) (*SubmitRefreshTaskOutput, error) {
	req, out := c.SubmitRefreshTaskRequest(input)
	return out, req.Send()
}

// SubmitRefreshTaskWithContext is the same as SubmitRefreshTask with the addition of
// the ability to pass a context and additional request options.
//
// See SubmitRefreshTask for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *MCDN) SubmitRefreshTaskWithContext(ctx volcengine.Context, input *SubmitRefreshTaskInput, opts ...request.Option) (*SubmitRefreshTaskOutput, error) {
	req, out := c.SubmitRefreshTaskRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type SubmitRefreshTaskInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	// Type is a required field
	Type *string `type:"string" json:",omitempty" required:"true"`

	// Urls is a required field
	Urls *string `type:"string" json:",omitempty" required:"true"`

	Vendor *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s SubmitRefreshTaskInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s SubmitRefreshTaskInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *SubmitRefreshTaskInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "SubmitRefreshTaskInput"}
	if s.Type == nil {
		invalidParams.Add(request.NewErrParamRequired("Type"))
	}
	if s.Urls == nil {
		invalidParams.Add(request.NewErrParamRequired("Urls"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetType sets the Type field's value.
func (s *SubmitRefreshTaskInput) SetType(v string) *SubmitRefreshTaskInput {
	s.Type = &v
	return s
}

// SetUrls sets the Urls field's value.
func (s *SubmitRefreshTaskInput) SetUrls(v string) *SubmitRefreshTaskInput {
	s.Urls = &v
	return s
}

// SetVendor sets the Vendor field's value.
func (s *SubmitRefreshTaskInput) SetVendor(v string) *SubmitRefreshTaskInput {
	s.Vendor = &v
	return s
}

type SubmitRefreshTaskOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	TaskId *string `type:"string" json:",omitempty"`

	TaskIds []*string `type:"list" json:",omitempty"`
}

// String returns the string representation
func (s SubmitRefreshTaskOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s SubmitRefreshTaskOutput) GoString() string {
	return s.String()
}

// SetTaskId sets the TaskId field's value.
func (s *SubmitRefreshTaskOutput) SetTaskId(v string) *SubmitRefreshTaskOutput {
	s.TaskId = &v
	return s
}

// SetTaskIds sets the TaskIds field's value.
func (s *SubmitRefreshTaskOutput) SetTaskIds(v []*string) *SubmitRefreshTaskOutput {
	s.TaskIds = v
	return s
}

// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package cdn

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opSubmitPreloadTaskCommon = "SubmitPreloadTask"

// SubmitPreloadTaskCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the SubmitPreloadTaskCommon operation. The "output" return
// value will be populated with the SubmitPreloadTaskCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned SubmitPreloadTaskCommon Request to send the API call to the service.
// the "output" return value is not valid until after SubmitPreloadTaskCommon Send returns without error.
//
// See SubmitPreloadTaskCommon for more information on using the SubmitPreloadTaskCommon
// API call, and error handling.
//
//    // Example sending a request using the SubmitPreloadTaskCommonRequest method.
//    req, resp := client.SubmitPreloadTaskCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CDN) SubmitPreloadTaskCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opSubmitPreloadTaskCommon,
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

// SubmitPreloadTaskCommon API operation for CDN.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CDN's
// API operation SubmitPreloadTaskCommon for usage and error information.
func (c *CDN) SubmitPreloadTaskCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.SubmitPreloadTaskCommonRequest(input)
	return out, req.Send()
}

// SubmitPreloadTaskCommonWithContext is the same as SubmitPreloadTaskCommon with the addition of
// the ability to pass a context and additional request options.
//
// See SubmitPreloadTaskCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CDN) SubmitPreloadTaskCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.SubmitPreloadTaskCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opSubmitPreloadTask = "SubmitPreloadTask"

// SubmitPreloadTaskRequest generates a "volcengine/request.Request" representing the
// client's request for the SubmitPreloadTask operation. The "output" return
// value will be populated with the SubmitPreloadTaskCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned SubmitPreloadTaskCommon Request to send the API call to the service.
// the "output" return value is not valid until after SubmitPreloadTaskCommon Send returns without error.
//
// See SubmitPreloadTask for more information on using the SubmitPreloadTask
// API call, and error handling.
//
//    // Example sending a request using the SubmitPreloadTaskRequest method.
//    req, resp := client.SubmitPreloadTaskRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CDN) SubmitPreloadTaskRequest(input *SubmitPreloadTaskInput) (req *request.Request, output *SubmitPreloadTaskOutput) {
	op := &request.Operation{
		Name:       opSubmitPreloadTask,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &SubmitPreloadTaskInput{}
	}

	output = &SubmitPreloadTaskOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// SubmitPreloadTask API operation for CDN.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CDN's
// API operation SubmitPreloadTask for usage and error information.
func (c *CDN) SubmitPreloadTask(input *SubmitPreloadTaskInput) (*SubmitPreloadTaskOutput, error) {
	req, out := c.SubmitPreloadTaskRequest(input)
	return out, req.Send()
}

// SubmitPreloadTaskWithContext is the same as SubmitPreloadTask with the addition of
// the ability to pass a context and additional request options.
//
// See SubmitPreloadTask for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CDN) SubmitPreloadTaskWithContext(ctx volcengine.Context, input *SubmitPreloadTaskInput, opts ...request.Option) (*SubmitPreloadTaskOutput, error) {
	req, out := c.SubmitPreloadTaskRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type RequestHeaderInstanceForSubmitPreloadTaskInput struct {
	_ struct{} `type:"structure"`

	Key *string `type:"string"`

	Value *string `type:"string"`
}

// String returns the string representation
func (s RequestHeaderInstanceForSubmitPreloadTaskInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s RequestHeaderInstanceForSubmitPreloadTaskInput) GoString() string {
	return s.String()
}

// SetKey sets the Key field's value.
func (s *RequestHeaderInstanceForSubmitPreloadTaskInput) SetKey(v string) *RequestHeaderInstanceForSubmitPreloadTaskInput {
	s.Key = &v
	return s
}

// SetValue sets the Value field's value.
func (s *RequestHeaderInstanceForSubmitPreloadTaskInput) SetValue(v string) *RequestHeaderInstanceForSubmitPreloadTaskInput {
	s.Value = &v
	return s
}

type SubmitPreloadTaskInput struct {
	_ struct{} `type:"structure"`

	Area *string `type:"string"`

	ConcurrentLimit *int64 `type:"int64"`

	Layer *string `type:"string"`

	RequestHeaderInstances []*RequestHeaderInstanceForSubmitPreloadTaskInput `type:"list"`

	// Urls is a required field
	Urls *string `type:"string" required:"true"`
}

// String returns the string representation
func (s SubmitPreloadTaskInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s SubmitPreloadTaskInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *SubmitPreloadTaskInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "SubmitPreloadTaskInput"}
	if s.Urls == nil {
		invalidParams.Add(request.NewErrParamRequired("Urls"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetArea sets the Area field's value.
func (s *SubmitPreloadTaskInput) SetArea(v string) *SubmitPreloadTaskInput {
	s.Area = &v
	return s
}

// SetConcurrentLimit sets the ConcurrentLimit field's value.
func (s *SubmitPreloadTaskInput) SetConcurrentLimit(v int64) *SubmitPreloadTaskInput {
	s.ConcurrentLimit = &v
	return s
}

// SetLayer sets the Layer field's value.
func (s *SubmitPreloadTaskInput) SetLayer(v string) *SubmitPreloadTaskInput {
	s.Layer = &v
	return s
}

// SetRequestHeaderInstances sets the RequestHeaderInstances field's value.
func (s *SubmitPreloadTaskInput) SetRequestHeaderInstances(v []*RequestHeaderInstanceForSubmitPreloadTaskInput) *SubmitPreloadTaskInput {
	s.RequestHeaderInstances = v
	return s
}

// SetUrls sets the Urls field's value.
func (s *SubmitPreloadTaskInput) SetUrls(v string) *SubmitPreloadTaskInput {
	s.Urls = &v
	return s
}

type SubmitPreloadTaskOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	CommitNum *int64 `type:"int64"`

	TaskID *string `type:"string"`
}

// String returns the string representation
func (s SubmitPreloadTaskOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s SubmitPreloadTaskOutput) GoString() string {
	return s.String()
}

// SetCommitNum sets the CommitNum field's value.
func (s *SubmitPreloadTaskOutput) SetCommitNum(v int64) *SubmitPreloadTaskOutput {
	s.CommitNum = &v
	return s
}

// SetTaskID sets the TaskID field's value.
func (s *SubmitPreloadTaskOutput) SetTaskID(v string) *SubmitPreloadTaskOutput {
	s.TaskID = &v
	return s
}
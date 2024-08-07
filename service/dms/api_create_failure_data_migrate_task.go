// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package dms

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opCreateFailureDataMigrateTaskCommon = "CreateFailureDataMigrateTask"

// CreateFailureDataMigrateTaskCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateFailureDataMigrateTaskCommon operation. The "output" return
// value will be populated with the CreateFailureDataMigrateTaskCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateFailureDataMigrateTaskCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateFailureDataMigrateTaskCommon Send returns without error.
//
// See CreateFailureDataMigrateTaskCommon for more information on using the CreateFailureDataMigrateTaskCommon
// API call, and error handling.
//
//    // Example sending a request using the CreateFailureDataMigrateTaskCommonRequest method.
//    req, resp := client.CreateFailureDataMigrateTaskCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *DMS) CreateFailureDataMigrateTaskCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opCreateFailureDataMigrateTaskCommon,
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

// CreateFailureDataMigrateTaskCommon API operation for DMS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for DMS's
// API operation CreateFailureDataMigrateTaskCommon for usage and error information.
func (c *DMS) CreateFailureDataMigrateTaskCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.CreateFailureDataMigrateTaskCommonRequest(input)
	return out, req.Send()
}

// CreateFailureDataMigrateTaskCommonWithContext is the same as CreateFailureDataMigrateTaskCommon with the addition of
// the ability to pass a context and additional request options.
//
// See CreateFailureDataMigrateTaskCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *DMS) CreateFailureDataMigrateTaskCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.CreateFailureDataMigrateTaskCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opCreateFailureDataMigrateTask = "CreateFailureDataMigrateTask"

// CreateFailureDataMigrateTaskRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateFailureDataMigrateTask operation. The "output" return
// value will be populated with the CreateFailureDataMigrateTaskCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateFailureDataMigrateTaskCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateFailureDataMigrateTaskCommon Send returns without error.
//
// See CreateFailureDataMigrateTask for more information on using the CreateFailureDataMigrateTask
// API call, and error handling.
//
//    // Example sending a request using the CreateFailureDataMigrateTaskRequest method.
//    req, resp := client.CreateFailureDataMigrateTaskRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *DMS) CreateFailureDataMigrateTaskRequest(input *CreateFailureDataMigrateTaskInput) (req *request.Request, output *CreateFailureDataMigrateTaskOutput) {
	op := &request.Operation{
		Name:       opCreateFailureDataMigrateTask,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateFailureDataMigrateTaskInput{}
	}

	output = &CreateFailureDataMigrateTaskOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// CreateFailureDataMigrateTask API operation for DMS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for DMS's
// API operation CreateFailureDataMigrateTask for usage and error information.
func (c *DMS) CreateFailureDataMigrateTask(input *CreateFailureDataMigrateTaskInput) (*CreateFailureDataMigrateTaskOutput, error) {
	req, out := c.CreateFailureDataMigrateTaskRequest(input)
	return out, req.Send()
}

// CreateFailureDataMigrateTaskWithContext is the same as CreateFailureDataMigrateTask with the addition of
// the ability to pass a context and additional request options.
//
// See CreateFailureDataMigrateTask for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *DMS) CreateFailureDataMigrateTaskWithContext(ctx volcengine.Context, input *CreateFailureDataMigrateTaskInput, opts ...request.Option) (*CreateFailureDataMigrateTaskOutput, error) {
	req, out := c.CreateFailureDataMigrateTaskRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type CreateFailureDataMigrateTaskInput struct {
	_ struct{} `type:"structure"`

	// OriginTaskID is a required field
	OriginTaskID *int64 `type:"int64" required:"true"`

	// TaskName is a required field
	TaskName *string `min:"3" max:"32" type:"string" required:"true"`
}

// String returns the string representation
func (s CreateFailureDataMigrateTaskInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateFailureDataMigrateTaskInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateFailureDataMigrateTaskInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "CreateFailureDataMigrateTaskInput"}
	if s.OriginTaskID == nil {
		invalidParams.Add(request.NewErrParamRequired("OriginTaskID"))
	}
	if s.TaskName == nil {
		invalidParams.Add(request.NewErrParamRequired("TaskName"))
	}
	if s.TaskName != nil && len(*s.TaskName) < 3 {
		invalidParams.Add(request.NewErrParamMinLen("TaskName", 3))
	}
	if s.TaskName != nil && len(*s.TaskName) > 32 {
		invalidParams.Add(request.NewErrParamMaxLen("TaskName", 32, *s.TaskName))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetOriginTaskID sets the OriginTaskID field's value.
func (s *CreateFailureDataMigrateTaskInput) SetOriginTaskID(v int64) *CreateFailureDataMigrateTaskInput {
	s.OriginTaskID = &v
	return s
}

// SetTaskName sets the TaskName field's value.
func (s *CreateFailureDataMigrateTaskInput) SetTaskName(v string) *CreateFailureDataMigrateTaskInput {
	s.TaskName = &v
	return s
}

type CreateFailureDataMigrateTaskOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	TaskID *int64 `type:"int64"`
}

// String returns the string representation
func (s CreateFailureDataMigrateTaskOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateFailureDataMigrateTaskOutput) GoString() string {
	return s.String()
}

// SetTaskID sets the TaskID field's value.
func (s *CreateFailureDataMigrateTaskOutput) SetTaskID(v int64) *CreateFailureDataMigrateTaskOutput {
	s.TaskID = &v
	return s
}

// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package rdsmysqlv2

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opModifyTaskCommon = "ModifyTask"

// ModifyTaskCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ModifyTaskCommon operation. The "output" return
// value will be populated with the ModifyTaskCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyTaskCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyTaskCommon Send returns without error.
//
// See ModifyTaskCommon for more information on using the ModifyTaskCommon
// API call, and error handling.
//
//    // Example sending a request using the ModifyTaskCommonRequest method.
//    req, resp := client.ModifyTaskCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *RDSMYSQLV2) ModifyTaskCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opModifyTaskCommon,
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

// ModifyTaskCommon API operation for RDS_MYSQL_V2.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RDS_MYSQL_V2's
// API operation ModifyTaskCommon for usage and error information.
func (c *RDSMYSQLV2) ModifyTaskCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ModifyTaskCommonRequest(input)
	return out, req.Send()
}

// ModifyTaskCommonWithContext is the same as ModifyTaskCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyTaskCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RDSMYSQLV2) ModifyTaskCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ModifyTaskCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opModifyTask = "ModifyTask"

// ModifyTaskRequest generates a "volcengine/request.Request" representing the
// client's request for the ModifyTask operation. The "output" return
// value will be populated with the ModifyTaskCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyTaskCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyTaskCommon Send returns without error.
//
// See ModifyTask for more information on using the ModifyTask
// API call, and error handling.
//
//    // Example sending a request using the ModifyTaskRequest method.
//    req, resp := client.ModifyTaskRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *RDSMYSQLV2) ModifyTaskRequest(input *ModifyTaskInput) (req *request.Request, output *ModifyTaskOutput) {
	op := &request.Operation{
		Name:       opModifyTask,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ModifyTaskInput{}
	}

	output = &ModifyTaskOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// ModifyTask API operation for RDS_MYSQL_V2.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RDS_MYSQL_V2's
// API operation ModifyTask for usage and error information.
func (c *RDSMYSQLV2) ModifyTask(input *ModifyTaskInput) (*ModifyTaskOutput, error) {
	req, out := c.ModifyTaskRequest(input)
	return out, req.Send()
}

// ModifyTaskWithContext is the same as ModifyTask with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyTask for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RDSMYSQLV2) ModifyTaskWithContext(ctx volcengine.Context, input *ModifyTaskInput, opts ...request.Option) (*ModifyTaskOutput, error) {
	req, out := c.ModifyTaskRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ModifyTaskInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	// InstanceId is a required field
	InstanceId *string `type:"string" json:",omitempty" required:"true"`

	ModifyBeginTime *string `type:"string" json:",omitempty"`

	ModifyEndTime *string `type:"string" json:",omitempty"`

	ModifyType *string `type:"string" json:",omitempty"`

	ProjectName *string `type:"string" json:",omitempty"`

	SwitchTime *string `type:"string" json:",omitempty"`

	// TaskEventOperation is a required field
	TaskEventOperation *string `type:"string" json:",omitempty" required:"true"`

	// TaskId is a required field
	TaskId *string `type:"string" json:",omitempty" required:"true"`
}

// String returns the string representation
func (s ModifyTaskInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyTaskInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ModifyTaskInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ModifyTaskInput"}
	if s.InstanceId == nil {
		invalidParams.Add(request.NewErrParamRequired("InstanceId"))
	}
	if s.TaskEventOperation == nil {
		invalidParams.Add(request.NewErrParamRequired("TaskEventOperation"))
	}
	if s.TaskId == nil {
		invalidParams.Add(request.NewErrParamRequired("TaskId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetInstanceId sets the InstanceId field's value.
func (s *ModifyTaskInput) SetInstanceId(v string) *ModifyTaskInput {
	s.InstanceId = &v
	return s
}

// SetModifyBeginTime sets the ModifyBeginTime field's value.
func (s *ModifyTaskInput) SetModifyBeginTime(v string) *ModifyTaskInput {
	s.ModifyBeginTime = &v
	return s
}

// SetModifyEndTime sets the ModifyEndTime field's value.
func (s *ModifyTaskInput) SetModifyEndTime(v string) *ModifyTaskInput {
	s.ModifyEndTime = &v
	return s
}

// SetModifyType sets the ModifyType field's value.
func (s *ModifyTaskInput) SetModifyType(v string) *ModifyTaskInput {
	s.ModifyType = &v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *ModifyTaskInput) SetProjectName(v string) *ModifyTaskInput {
	s.ProjectName = &v
	return s
}

// SetSwitchTime sets the SwitchTime field's value.
func (s *ModifyTaskInput) SetSwitchTime(v string) *ModifyTaskInput {
	s.SwitchTime = &v
	return s
}

// SetTaskEventOperation sets the TaskEventOperation field's value.
func (s *ModifyTaskInput) SetTaskEventOperation(v string) *ModifyTaskInput {
	s.TaskEventOperation = &v
	return s
}

// SetTaskId sets the TaskId field's value.
func (s *ModifyTaskInput) SetTaskId(v string) *ModifyTaskInput {
	s.TaskId = &v
	return s
}

type ModifyTaskOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s ModifyTaskOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyTaskOutput) GoString() string {
	return s.String()
}

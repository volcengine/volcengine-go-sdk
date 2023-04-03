// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package ecs

import (
	"encoding/json"

	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opUpdateLongTaskCommon = "UpdateLongTask"

// UpdateLongTaskCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the UpdateLongTaskCommon operation. The "output" return
// value will be populated with the UpdateLongTaskCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned UpdateLongTaskCommon Request to send the API call to the service.
// the "output" return value is not valid until after UpdateLongTaskCommon Send returns without error.
//
// See UpdateLongTaskCommon for more information on using the UpdateLongTaskCommon
// API call, and error handling.
//
//    // Example sending a request using the UpdateLongTaskCommonRequest method.
//    req, resp := client.UpdateLongTaskCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ECS) UpdateLongTaskCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opUpdateLongTaskCommon,
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

// UpdateLongTaskCommon API operation for ECS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ECS's
// API operation UpdateLongTaskCommon for usage and error information.
func (c *ECS) UpdateLongTaskCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.UpdateLongTaskCommonRequest(input)
	return out, req.Send()
}

// UpdateLongTaskCommonWithContext is the same as UpdateLongTaskCommon with the addition of
// the ability to pass a context and additional request options.
//
// See UpdateLongTaskCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ECS) UpdateLongTaskCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.UpdateLongTaskCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opUpdateLongTask = "UpdateLongTask"

// UpdateLongTaskRequest generates a "volcengine/request.Request" representing the
// client's request for the UpdateLongTask operation. The "output" return
// value will be populated with the UpdateLongTaskCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned UpdateLongTaskCommon Request to send the API call to the service.
// the "output" return value is not valid until after UpdateLongTaskCommon Send returns without error.
//
// See UpdateLongTask for more information on using the UpdateLongTask
// API call, and error handling.
//
//    // Example sending a request using the UpdateLongTaskRequest method.
//    req, resp := client.UpdateLongTaskRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ECS) UpdateLongTaskRequest(input *UpdateLongTaskInput) (req *request.Request, output *UpdateLongTaskOutput) {
	op := &request.Operation{
		Name:       opUpdateLongTask,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateLongTaskInput{}
	}

	output = &UpdateLongTaskOutput{}
	req = c.newRequest(op, input, output)

	return
}

// UpdateLongTask API operation for ECS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ECS's
// API operation UpdateLongTask for usage and error information.
func (c *ECS) UpdateLongTask(input *UpdateLongTaskInput) (*UpdateLongTaskOutput, error) {
	req, out := c.UpdateLongTaskRequest(input)
	return out, req.Send()
}

// UpdateLongTaskWithContext is the same as UpdateLongTask with the addition of
// the ability to pass a context and additional request options.
//
// See UpdateLongTask for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ECS) UpdateLongTaskWithContext(ctx volcengine.Context, input *UpdateLongTaskInput, opts ...request.Option) (*UpdateLongTaskOutput, error) {
	req, out := c.UpdateLongTaskRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type AutoGrowthForUpdateLongTaskInput struct {
	_ struct{} `type:"structure"`

	Expect *json.Number `type:"json_number"`
}

// String returns the string representation
func (s AutoGrowthForUpdateLongTaskInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s AutoGrowthForUpdateLongTaskInput) GoString() string {
	return s.String()
}

// SetExpect sets the Expect field's value.
func (s *AutoGrowthForUpdateLongTaskInput) SetExpect(v json.Number) *AutoGrowthForUpdateLongTaskInput {
	s.Expect = &v
	return s
}

type UpdateLongTaskInput struct {
	_ struct{} `type:"structure"`

	AutoGrowth *AutoGrowthForUpdateLongTaskInput `type:"structure"`

	EndAt *string `type:"string"`

	Progress *int64 `type:"int64"`

	Status *string `type:"string" enum:"StatusForUpdateLongTaskInput"`

	TaskId *string `type:"string"`

	UpdatedAt *string `type:"string"`
}

// String returns the string representation
func (s UpdateLongTaskInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s UpdateLongTaskInput) GoString() string {
	return s.String()
}

// SetAutoGrowth sets the AutoGrowth field's value.
func (s *UpdateLongTaskInput) SetAutoGrowth(v *AutoGrowthForUpdateLongTaskInput) *UpdateLongTaskInput {
	s.AutoGrowth = v
	return s
}

// SetEndAt sets the EndAt field's value.
func (s *UpdateLongTaskInput) SetEndAt(v string) *UpdateLongTaskInput {
	s.EndAt = &v
	return s
}

// SetProgress sets the Progress field's value.
func (s *UpdateLongTaskInput) SetProgress(v int64) *UpdateLongTaskInput {
	s.Progress = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *UpdateLongTaskInput) SetStatus(v string) *UpdateLongTaskInput {
	s.Status = &v
	return s
}

// SetTaskId sets the TaskId field's value.
func (s *UpdateLongTaskInput) SetTaskId(v string) *UpdateLongTaskInput {
	s.TaskId = &v
	return s
}

// SetUpdatedAt sets the UpdatedAt field's value.
func (s *UpdateLongTaskInput) SetUpdatedAt(v string) *UpdateLongTaskInput {
	s.UpdatedAt = &v
	return s
}

type UpdateLongTaskOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s UpdateLongTaskOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s UpdateLongTaskOutput) GoString() string {
	return s.String()
}

const (
	// StatusForUpdateLongTaskInputUnknownStatus is a StatusForUpdateLongTaskInput enum value
	StatusForUpdateLongTaskInputUnknownStatus = "UnknownStatus"

	// StatusForUpdateLongTaskInputPending is a StatusForUpdateLongTaskInput enum value
	StatusForUpdateLongTaskInputPending = "Pending"

	// StatusForUpdateLongTaskInputRunning is a StatusForUpdateLongTaskInput enum value
	StatusForUpdateLongTaskInputRunning = "Running"

	// StatusForUpdateLongTaskInputSucceeded is a StatusForUpdateLongTaskInput enum value
	StatusForUpdateLongTaskInputSucceeded = "Succeeded"

	// StatusForUpdateLongTaskInputFailed is a StatusForUpdateLongTaskInput enum value
	StatusForUpdateLongTaskInputFailed = "Failed"
)
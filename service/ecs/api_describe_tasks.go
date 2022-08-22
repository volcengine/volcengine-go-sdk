// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package ecs

import (
	"encoding/json"

	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribeTasksCommon = "DescribeTasks"

// DescribeTasksCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeTasksCommon operation. The "output" return
// value will be populated with the DescribeTasksCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeTasksCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeTasksCommon Send returns without error.
//
// See DescribeTasksCommon for more information on using the DescribeTasksCommon
// API call, and error handling.
//
//    // Example sending a request using the DescribeTasksCommonRequest method.
//    req, resp := client.DescribeTasksCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ECS) DescribeTasksCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeTasksCommon,
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

// DescribeTasksCommon API operation for ECS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ECS's
// API operation DescribeTasksCommon for usage and error information.
func (c *ECS) DescribeTasksCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeTasksCommonRequest(input)
	return out, req.Send()
}

// DescribeTasksCommonWithContext is the same as DescribeTasksCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeTasksCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ECS) DescribeTasksCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeTasksCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeTasks = "DescribeTasks"

// DescribeTasksRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeTasks operation. The "output" return
// value will be populated with the DescribeTasksCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeTasksCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeTasksCommon Send returns without error.
//
// See DescribeTasks for more information on using the DescribeTasks
// API call, and error handling.
//
//    // Example sending a request using the DescribeTasksRequest method.
//    req, resp := client.DescribeTasksRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ECS) DescribeTasksRequest(input *DescribeTasksInput) (req *request.Request, output *DescribeTasksOutput) {
	op := &request.Operation{
		Name:       opDescribeTasks,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeTasksInput{}
	}

	output = &DescribeTasksOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DescribeTasks API operation for ECS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ECS's
// API operation DescribeTasks for usage and error information.
func (c *ECS) DescribeTasks(input *DescribeTasksInput) (*DescribeTasksOutput, error) {
	req, out := c.DescribeTasksRequest(input)
	return out, req.Send()
}

// DescribeTasksWithContext is the same as DescribeTasks with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeTasks for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ECS) DescribeTasksWithContext(ctx volcengine.Context, input *DescribeTasksInput, opts ...request.Option) (*DescribeTasksOutput, error) {
	req, out := c.DescribeTasksRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DescribeTasksInput struct {
	_ struct{} `type:"structure"`

	MaxResults *json.Number `type:"json_number"`

	NextToken *string `type:"string"`

	ResourceId *string `type:"string"`

	TaskIds []*string `type:"list"`
}

// String returns the string representation
func (s DescribeTasksInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeTasksInput) GoString() string {
	return s.String()
}

// SetMaxResults sets the MaxResults field's value.
func (s *DescribeTasksInput) SetMaxResults(v json.Number) *DescribeTasksInput {
	s.MaxResults = &v
	return s
}

// SetNextToken sets the NextToken field's value.
func (s *DescribeTasksInput) SetNextToken(v string) *DescribeTasksInput {
	s.NextToken = &v
	return s
}

// SetResourceId sets the ResourceId field's value.
func (s *DescribeTasksInput) SetResourceId(v string) *DescribeTasksInput {
	s.ResourceId = &v
	return s
}

// SetTaskIds sets the TaskIds field's value.
func (s *DescribeTasksInput) SetTaskIds(v []*string) *DescribeTasksInput {
	s.TaskIds = v
	return s
}

type DescribeTasksOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	NextToken *string `type:"string"`

	Tasks []*TaskForDescribeTasksOutput `type:"list"`
}

// String returns the string representation
func (s DescribeTasksOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeTasksOutput) GoString() string {
	return s.String()
}

// SetNextToken sets the NextToken field's value.
func (s *DescribeTasksOutput) SetNextToken(v string) *DescribeTasksOutput {
	s.NextToken = &v
	return s
}

// SetTasks sets the Tasks field's value.
func (s *DescribeTasksOutput) SetTasks(v []*TaskForDescribeTasksOutput) *DescribeTasksOutput {
	s.Tasks = v
	return s
}

type TaskForDescribeTasksOutput struct {
	_ struct{} `type:"structure"`

	CreatedAt *string `type:"string"`

	EndAt *string `type:"string"`

	Id *string `type:"string"`

	Process *int64 `type:"int64"`

	ResourceId *string `type:"string"`

	Status *string `type:"string" enum:"StatusForDescribeTasksOutput"`

	Type *string `type:"string" enum:"TypeForDescribeTasksOutput"`

	UpdatedAt *string `type:"string"`
}

// String returns the string representation
func (s TaskForDescribeTasksOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s TaskForDescribeTasksOutput) GoString() string {
	return s.String()
}

// SetCreatedAt sets the CreatedAt field's value.
func (s *TaskForDescribeTasksOutput) SetCreatedAt(v string) *TaskForDescribeTasksOutput {
	s.CreatedAt = &v
	return s
}

// SetEndAt sets the EndAt field's value.
func (s *TaskForDescribeTasksOutput) SetEndAt(v string) *TaskForDescribeTasksOutput {
	s.EndAt = &v
	return s
}

// SetId sets the Id field's value.
func (s *TaskForDescribeTasksOutput) SetId(v string) *TaskForDescribeTasksOutput {
	s.Id = &v
	return s
}

// SetProcess sets the Process field's value.
func (s *TaskForDescribeTasksOutput) SetProcess(v int64) *TaskForDescribeTasksOutput {
	s.Process = &v
	return s
}

// SetResourceId sets the ResourceId field's value.
func (s *TaskForDescribeTasksOutput) SetResourceId(v string) *TaskForDescribeTasksOutput {
	s.ResourceId = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *TaskForDescribeTasksOutput) SetStatus(v string) *TaskForDescribeTasksOutput {
	s.Status = &v
	return s
}

// SetType sets the Type field's value.
func (s *TaskForDescribeTasksOutput) SetType(v string) *TaskForDescribeTasksOutput {
	s.Type = &v
	return s
}

// SetUpdatedAt sets the UpdatedAt field's value.
func (s *TaskForDescribeTasksOutput) SetUpdatedAt(v string) *TaskForDescribeTasksOutput {
	s.UpdatedAt = &v
	return s
}

const (
	// StatusForDescribeTasksOutputUnknownStatus is a StatusForDescribeTasksOutput enum value
	StatusForDescribeTasksOutputUnknownStatus = "UnknownStatus"

	// StatusForDescribeTasksOutputPending is a StatusForDescribeTasksOutput enum value
	StatusForDescribeTasksOutputPending = "Pending"

	// StatusForDescribeTasksOutputRunning is a StatusForDescribeTasksOutput enum value
	StatusForDescribeTasksOutputRunning = "Running"

	// StatusForDescribeTasksOutputSucceeded is a StatusForDescribeTasksOutput enum value
	StatusForDescribeTasksOutputSucceeded = "Succeeded"

	// StatusForDescribeTasksOutputFailed is a StatusForDescribeTasksOutput enum value
	StatusForDescribeTasksOutputFailed = "Failed"
)

const (
	// TypeForDescribeTasksOutputUnknownType is a TypeForDescribeTasksOutput enum value
	TypeForDescribeTasksOutputUnknownType = "UnknownType"

	// TypeForDescribeTasksOutputExportImage is a TypeForDescribeTasksOutput enum value
	TypeForDescribeTasksOutputExportImage = "ExportImage"

	// TypeForDescribeTasksOutputCopyImage is a TypeForDescribeTasksOutput enum value
	TypeForDescribeTasksOutputCopyImage = "CopyImage"
)

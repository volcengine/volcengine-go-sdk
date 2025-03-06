// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package ecs

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opInvokeCommandCommon = "InvokeCommand"

// InvokeCommandCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the InvokeCommandCommon operation. The "output" return
// value will be populated with the InvokeCommandCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned InvokeCommandCommon Request to send the API call to the service.
// the "output" return value is not valid until after InvokeCommandCommon Send returns without error.
//
// See InvokeCommandCommon for more information on using the InvokeCommandCommon
// API call, and error handling.
//
//    // Example sending a request using the InvokeCommandCommonRequest method.
//    req, resp := client.InvokeCommandCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ECS) InvokeCommandCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opInvokeCommandCommon,
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

// InvokeCommandCommon API operation for ECS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ECS's
// API operation InvokeCommandCommon for usage and error information.
func (c *ECS) InvokeCommandCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.InvokeCommandCommonRequest(input)
	return out, req.Send()
}

// InvokeCommandCommonWithContext is the same as InvokeCommandCommon with the addition of
// the ability to pass a context and additional request options.
//
// See InvokeCommandCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ECS) InvokeCommandCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.InvokeCommandCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opInvokeCommand = "InvokeCommand"

// InvokeCommandRequest generates a "volcengine/request.Request" representing the
// client's request for the InvokeCommand operation. The "output" return
// value will be populated with the InvokeCommandCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned InvokeCommandCommon Request to send the API call to the service.
// the "output" return value is not valid until after InvokeCommandCommon Send returns without error.
//
// See InvokeCommand for more information on using the InvokeCommand
// API call, and error handling.
//
//    // Example sending a request using the InvokeCommandRequest method.
//    req, resp := client.InvokeCommandRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ECS) InvokeCommandRequest(input *InvokeCommandInput) (req *request.Request, output *InvokeCommandOutput) {
	op := &request.Operation{
		Name:       opInvokeCommand,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &InvokeCommandInput{}
	}

	output = &InvokeCommandOutput{}
	req = c.newRequest(op, input, output)

	return
}

// InvokeCommand API operation for ECS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ECS's
// API operation InvokeCommand for usage and error information.
func (c *ECS) InvokeCommand(input *InvokeCommandInput) (*InvokeCommandOutput, error) {
	req, out := c.InvokeCommandRequest(input)
	return out, req.Send()
}

// InvokeCommandWithContext is the same as InvokeCommand with the addition of
// the ability to pass a context and additional request options.
//
// See InvokeCommand for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ECS) InvokeCommandWithContext(ctx volcengine.Context, input *InvokeCommandInput, opts ...request.Option) (*InvokeCommandOutput, error) {
	req, out := c.InvokeCommandRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type InvokeCommandInput struct {
	_ struct{} `type:"structure"`

	// CommandId is a required field
	CommandId *string `type:"string" required:"true"`

	Frequency *string `type:"string"`

	// InstanceIds is a required field
	InstanceIds []*string `type:"list" required:"true"`

	InvocationDescription *string `type:"string"`

	// InvocationName is a required field
	InvocationName *string `type:"string" required:"true"`

	LaunchTime *string `type:"string"`

	Parameters *string `type:"string"`

	ProjectName *string `type:"string"`

	RecurrenceEndTime *string `type:"string"`

	RepeatMode *string `type:"string"`

	Tags []*TagForInvokeCommandInput `type:"list"`

	Timeout *int32 `type:"int32"`

	Username *string `type:"string"`

	WindowsPassword *string `type:"string"`

	WorkingDir *string `type:"string"`
}

// String returns the string representation
func (s InvokeCommandInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s InvokeCommandInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *InvokeCommandInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "InvokeCommandInput"}
	if s.CommandId == nil {
		invalidParams.Add(request.NewErrParamRequired("CommandId"))
	}
	if s.InstanceIds == nil {
		invalidParams.Add(request.NewErrParamRequired("InstanceIds"))
	}
	if s.InvocationName == nil {
		invalidParams.Add(request.NewErrParamRequired("InvocationName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetCommandId sets the CommandId field's value.
func (s *InvokeCommandInput) SetCommandId(v string) *InvokeCommandInput {
	s.CommandId = &v
	return s
}

// SetFrequency sets the Frequency field's value.
func (s *InvokeCommandInput) SetFrequency(v string) *InvokeCommandInput {
	s.Frequency = &v
	return s
}

// SetInstanceIds sets the InstanceIds field's value.
func (s *InvokeCommandInput) SetInstanceIds(v []*string) *InvokeCommandInput {
	s.InstanceIds = v
	return s
}

// SetInvocationDescription sets the InvocationDescription field's value.
func (s *InvokeCommandInput) SetInvocationDescription(v string) *InvokeCommandInput {
	s.InvocationDescription = &v
	return s
}

// SetInvocationName sets the InvocationName field's value.
func (s *InvokeCommandInput) SetInvocationName(v string) *InvokeCommandInput {
	s.InvocationName = &v
	return s
}

// SetLaunchTime sets the LaunchTime field's value.
func (s *InvokeCommandInput) SetLaunchTime(v string) *InvokeCommandInput {
	s.LaunchTime = &v
	return s
}

// SetParameters sets the Parameters field's value.
func (s *InvokeCommandInput) SetParameters(v string) *InvokeCommandInput {
	s.Parameters = &v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *InvokeCommandInput) SetProjectName(v string) *InvokeCommandInput {
	s.ProjectName = &v
	return s
}

// SetRecurrenceEndTime sets the RecurrenceEndTime field's value.
func (s *InvokeCommandInput) SetRecurrenceEndTime(v string) *InvokeCommandInput {
	s.RecurrenceEndTime = &v
	return s
}

// SetRepeatMode sets the RepeatMode field's value.
func (s *InvokeCommandInput) SetRepeatMode(v string) *InvokeCommandInput {
	s.RepeatMode = &v
	return s
}

// SetTags sets the Tags field's value.
func (s *InvokeCommandInput) SetTags(v []*TagForInvokeCommandInput) *InvokeCommandInput {
	s.Tags = v
	return s
}

// SetTimeout sets the Timeout field's value.
func (s *InvokeCommandInput) SetTimeout(v int32) *InvokeCommandInput {
	s.Timeout = &v
	return s
}

// SetUsername sets the Username field's value.
func (s *InvokeCommandInput) SetUsername(v string) *InvokeCommandInput {
	s.Username = &v
	return s
}

// SetWindowsPassword sets the WindowsPassword field's value.
func (s *InvokeCommandInput) SetWindowsPassword(v string) *InvokeCommandInput {
	s.WindowsPassword = &v
	return s
}

// SetWorkingDir sets the WorkingDir field's value.
func (s *InvokeCommandInput) SetWorkingDir(v string) *InvokeCommandInput {
	s.WorkingDir = &v
	return s
}

type InvokeCommandOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	InvocationId *string `type:"string"`
}

// String returns the string representation
func (s InvokeCommandOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s InvokeCommandOutput) GoString() string {
	return s.String()
}

// SetInvocationId sets the InvocationId field's value.
func (s *InvokeCommandOutput) SetInvocationId(v string) *InvokeCommandOutput {
	s.InvocationId = &v
	return s
}

type TagForInvokeCommandInput struct {
	_ struct{} `type:"structure"`

	Key *string `type:"string"`

	Value *string `type:"string"`
}

// String returns the string representation
func (s TagForInvokeCommandInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s TagForInvokeCommandInput) GoString() string {
	return s.String()
}

// SetKey sets the Key field's value.
func (s *TagForInvokeCommandInput) SetKey(v string) *TagForInvokeCommandInput {
	s.Key = &v
	return s
}

// SetValue sets the Value field's value.
func (s *TagForInvokeCommandInput) SetValue(v string) *TagForInvokeCommandInput {
	s.Value = &v
	return s
}

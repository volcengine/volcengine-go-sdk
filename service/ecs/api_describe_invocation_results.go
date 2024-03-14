// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package ecs

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribeInvocationResultsCommon = "DescribeInvocationResults"

// DescribeInvocationResultsCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeInvocationResultsCommon operation. The "output" return
// value will be populated with the DescribeInvocationResultsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeInvocationResultsCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeInvocationResultsCommon Send returns without error.
//
// See DescribeInvocationResultsCommon for more information on using the DescribeInvocationResultsCommon
// API call, and error handling.
//
//    // Example sending a request using the DescribeInvocationResultsCommonRequest method.
//    req, resp := client.DescribeInvocationResultsCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ECS) DescribeInvocationResultsCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeInvocationResultsCommon,
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

// DescribeInvocationResultsCommon API operation for ECS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ECS's
// API operation DescribeInvocationResultsCommon for usage and error information.
func (c *ECS) DescribeInvocationResultsCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeInvocationResultsCommonRequest(input)
	return out, req.Send()
}

// DescribeInvocationResultsCommonWithContext is the same as DescribeInvocationResultsCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeInvocationResultsCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ECS) DescribeInvocationResultsCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeInvocationResultsCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeInvocationResults = "DescribeInvocationResults"

// DescribeInvocationResultsRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeInvocationResults operation. The "output" return
// value will be populated with the DescribeInvocationResultsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeInvocationResultsCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeInvocationResultsCommon Send returns without error.
//
// See DescribeInvocationResults for more information on using the DescribeInvocationResults
// API call, and error handling.
//
//    // Example sending a request using the DescribeInvocationResultsRequest method.
//    req, resp := client.DescribeInvocationResultsRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ECS) DescribeInvocationResultsRequest(input *DescribeInvocationResultsInput) (req *request.Request, output *DescribeInvocationResultsOutput) {
	op := &request.Operation{
		Name:       opDescribeInvocationResults,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeInvocationResultsInput{}
	}

	output = &DescribeInvocationResultsOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DescribeInvocationResults API operation for ECS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ECS's
// API operation DescribeInvocationResults for usage and error information.
func (c *ECS) DescribeInvocationResults(input *DescribeInvocationResultsInput) (*DescribeInvocationResultsOutput, error) {
	req, out := c.DescribeInvocationResultsRequest(input)
	return out, req.Send()
}

// DescribeInvocationResultsWithContext is the same as DescribeInvocationResults with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeInvocationResults for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ECS) DescribeInvocationResultsWithContext(ctx volcengine.Context, input *DescribeInvocationResultsInput, opts ...request.Option) (*DescribeInvocationResultsOutput, error) {
	req, out := c.DescribeInvocationResultsRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DescribeInvocationResultsInput struct {
	_ struct{} `type:"structure"`

	CommandId *string `type:"string"`

	InstanceId *string `type:"string"`

	// InvocationId is a required field
	InvocationId *string `type:"string" required:"true"`

	InvocationResultStatus *string `type:"string"`

	PageNumber *int32 `type:"int32"`

	PageSize *int32 `type:"int32"`
}

// String returns the string representation
func (s DescribeInvocationResultsInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeInvocationResultsInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeInvocationResultsInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DescribeInvocationResultsInput"}
	if s.InvocationId == nil {
		invalidParams.Add(request.NewErrParamRequired("InvocationId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetCommandId sets the CommandId field's value.
func (s *DescribeInvocationResultsInput) SetCommandId(v string) *DescribeInvocationResultsInput {
	s.CommandId = &v
	return s
}

// SetInstanceId sets the InstanceId field's value.
func (s *DescribeInvocationResultsInput) SetInstanceId(v string) *DescribeInvocationResultsInput {
	s.InstanceId = &v
	return s
}

// SetInvocationId sets the InvocationId field's value.
func (s *DescribeInvocationResultsInput) SetInvocationId(v string) *DescribeInvocationResultsInput {
	s.InvocationId = &v
	return s
}

// SetInvocationResultStatus sets the InvocationResultStatus field's value.
func (s *DescribeInvocationResultsInput) SetInvocationResultStatus(v string) *DescribeInvocationResultsInput {
	s.InvocationResultStatus = &v
	return s
}

// SetPageNumber sets the PageNumber field's value.
func (s *DescribeInvocationResultsInput) SetPageNumber(v int32) *DescribeInvocationResultsInput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *DescribeInvocationResultsInput) SetPageSize(v int32) *DescribeInvocationResultsInput {
	s.PageSize = &v
	return s
}

type DescribeInvocationResultsOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	InvocationResults []*InvocationResultForDescribeInvocationResultsOutput `type:"list"`

	PageNumber *int32 `type:"int32"`

	PageSize *int32 `type:"int32"`

	TotalCount *int32 `type:"int32"`
}

// String returns the string representation
func (s DescribeInvocationResultsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeInvocationResultsOutput) GoString() string {
	return s.String()
}

// SetInvocationResults sets the InvocationResults field's value.
func (s *DescribeInvocationResultsOutput) SetInvocationResults(v []*InvocationResultForDescribeInvocationResultsOutput) *DescribeInvocationResultsOutput {
	s.InvocationResults = v
	return s
}

// SetPageNumber sets the PageNumber field's value.
func (s *DescribeInvocationResultsOutput) SetPageNumber(v int32) *DescribeInvocationResultsOutput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *DescribeInvocationResultsOutput) SetPageSize(v int32) *DescribeInvocationResultsOutput {
	s.PageSize = &v
	return s
}

// SetTotalCount sets the TotalCount field's value.
func (s *DescribeInvocationResultsOutput) SetTotalCount(v int32) *DescribeInvocationResultsOutput {
	s.TotalCount = &v
	return s
}

type InvocationResultForDescribeInvocationResultsOutput struct {
	_ struct{} `type:"structure"`

	CommandId *string `type:"string"`

	EndTime *string `type:"string"`

	ErrorCode *string `type:"string"`

	ErrorMessage *string `type:"string"`

	ExitCode *int32 `type:"int32"`

	InstanceId *string `type:"string"`

	InvocationId *string `type:"string"`

	InvocationResultId *string `type:"string"`

	InvocationResultStatus *string `type:"string"`

	Output *string `type:"string"`

	StartTime *string `type:"string"`

	Username *string `type:"string"`
}

// String returns the string representation
func (s InvocationResultForDescribeInvocationResultsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s InvocationResultForDescribeInvocationResultsOutput) GoString() string {
	return s.String()
}

// SetCommandId sets the CommandId field's value.
func (s *InvocationResultForDescribeInvocationResultsOutput) SetCommandId(v string) *InvocationResultForDescribeInvocationResultsOutput {
	s.CommandId = &v
	return s
}

// SetEndTime sets the EndTime field's value.
func (s *InvocationResultForDescribeInvocationResultsOutput) SetEndTime(v string) *InvocationResultForDescribeInvocationResultsOutput {
	s.EndTime = &v
	return s
}

// SetErrorCode sets the ErrorCode field's value.
func (s *InvocationResultForDescribeInvocationResultsOutput) SetErrorCode(v string) *InvocationResultForDescribeInvocationResultsOutput {
	s.ErrorCode = &v
	return s
}

// SetErrorMessage sets the ErrorMessage field's value.
func (s *InvocationResultForDescribeInvocationResultsOutput) SetErrorMessage(v string) *InvocationResultForDescribeInvocationResultsOutput {
	s.ErrorMessage = &v
	return s
}

// SetExitCode sets the ExitCode field's value.
func (s *InvocationResultForDescribeInvocationResultsOutput) SetExitCode(v int32) *InvocationResultForDescribeInvocationResultsOutput {
	s.ExitCode = &v
	return s
}

// SetInstanceId sets the InstanceId field's value.
func (s *InvocationResultForDescribeInvocationResultsOutput) SetInstanceId(v string) *InvocationResultForDescribeInvocationResultsOutput {
	s.InstanceId = &v
	return s
}

// SetInvocationId sets the InvocationId field's value.
func (s *InvocationResultForDescribeInvocationResultsOutput) SetInvocationId(v string) *InvocationResultForDescribeInvocationResultsOutput {
	s.InvocationId = &v
	return s
}

// SetInvocationResultId sets the InvocationResultId field's value.
func (s *InvocationResultForDescribeInvocationResultsOutput) SetInvocationResultId(v string) *InvocationResultForDescribeInvocationResultsOutput {
	s.InvocationResultId = &v
	return s
}

// SetInvocationResultStatus sets the InvocationResultStatus field's value.
func (s *InvocationResultForDescribeInvocationResultsOutput) SetInvocationResultStatus(v string) *InvocationResultForDescribeInvocationResultsOutput {
	s.InvocationResultStatus = &v
	return s
}

// SetOutput sets the Output field's value.
func (s *InvocationResultForDescribeInvocationResultsOutput) SetOutput(v string) *InvocationResultForDescribeInvocationResultsOutput {
	s.Output = &v
	return s
}

// SetStartTime sets the StartTime field's value.
func (s *InvocationResultForDescribeInvocationResultsOutput) SetStartTime(v string) *InvocationResultForDescribeInvocationResultsOutput {
	s.StartTime = &v
	return s
}

// SetUsername sets the Username field's value.
func (s *InvocationResultForDescribeInvocationResultsOutput) SetUsername(v string) *InvocationResultForDescribeInvocationResultsOutput {
	s.Username = &v
	return s
}

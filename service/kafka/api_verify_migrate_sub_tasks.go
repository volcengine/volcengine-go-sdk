// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package kafka

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opVerifyMigrateSubTasksCommon = "VerifyMigrateSubTasks"

// VerifyMigrateSubTasksCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the VerifyMigrateSubTasksCommon operation. The "output" return
// value will be populated with the VerifyMigrateSubTasksCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned VerifyMigrateSubTasksCommon Request to send the API call to the service.
// the "output" return value is not valid until after VerifyMigrateSubTasksCommon Send returns without error.
//
// See VerifyMigrateSubTasksCommon for more information on using the VerifyMigrateSubTasksCommon
// API call, and error handling.
//
//    // Example sending a request using the VerifyMigrateSubTasksCommonRequest method.
//    req, resp := client.VerifyMigrateSubTasksCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *KAFKA) VerifyMigrateSubTasksCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opVerifyMigrateSubTasksCommon,
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

// VerifyMigrateSubTasksCommon API operation for KAFKA.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for KAFKA's
// API operation VerifyMigrateSubTasksCommon for usage and error information.
func (c *KAFKA) VerifyMigrateSubTasksCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.VerifyMigrateSubTasksCommonRequest(input)
	return out, req.Send()
}

// VerifyMigrateSubTasksCommonWithContext is the same as VerifyMigrateSubTasksCommon with the addition of
// the ability to pass a context and additional request options.
//
// See VerifyMigrateSubTasksCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *KAFKA) VerifyMigrateSubTasksCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.VerifyMigrateSubTasksCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opVerifyMigrateSubTasks = "VerifyMigrateSubTasks"

// VerifyMigrateSubTasksRequest generates a "volcengine/request.Request" representing the
// client's request for the VerifyMigrateSubTasks operation. The "output" return
// value will be populated with the VerifyMigrateSubTasksCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned VerifyMigrateSubTasksCommon Request to send the API call to the service.
// the "output" return value is not valid until after VerifyMigrateSubTasksCommon Send returns without error.
//
// See VerifyMigrateSubTasks for more information on using the VerifyMigrateSubTasks
// API call, and error handling.
//
//    // Example sending a request using the VerifyMigrateSubTasksRequest method.
//    req, resp := client.VerifyMigrateSubTasksRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *KAFKA) VerifyMigrateSubTasksRequest(input *VerifyMigrateSubTasksInput) (req *request.Request, output *VerifyMigrateSubTasksOutput) {
	op := &request.Operation{
		Name:       opVerifyMigrateSubTasks,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &VerifyMigrateSubTasksInput{}
	}

	output = &VerifyMigrateSubTasksOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// VerifyMigrateSubTasks API operation for KAFKA.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for KAFKA's
// API operation VerifyMigrateSubTasks for usage and error information.
func (c *KAFKA) VerifyMigrateSubTasks(input *VerifyMigrateSubTasksInput) (*VerifyMigrateSubTasksOutput, error) {
	req, out := c.VerifyMigrateSubTasksRequest(input)
	return out, req.Send()
}

// VerifyMigrateSubTasksWithContext is the same as VerifyMigrateSubTasks with the addition of
// the ability to pass a context and additional request options.
//
// See VerifyMigrateSubTasks for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *KAFKA) VerifyMigrateSubTasksWithContext(ctx volcengine.Context, input *VerifyMigrateSubTasksInput, opts ...request.Option) (*VerifyMigrateSubTasksOutput, error) {
	req, out := c.VerifyMigrateSubTasksRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ItemForVerifyMigrateSubTasksInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Config *string `type:"string" json:",omitempty"`

	Name *string `type:"string" json:",omitempty"`

	Type *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s ItemForVerifyMigrateSubTasksInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ItemForVerifyMigrateSubTasksInput) GoString() string {
	return s.String()
}

// SetConfig sets the Config field's value.
func (s *ItemForVerifyMigrateSubTasksInput) SetConfig(v string) *ItemForVerifyMigrateSubTasksInput {
	s.Config = &v
	return s
}

// SetName sets the Name field's value.
func (s *ItemForVerifyMigrateSubTasksInput) SetName(v string) *ItemForVerifyMigrateSubTasksInput {
	s.Name = &v
	return s
}

// SetType sets the Type field's value.
func (s *ItemForVerifyMigrateSubTasksInput) SetType(v string) *ItemForVerifyMigrateSubTasksInput {
	s.Type = &v
	return s
}

type SubTaskResultForVerifyMigrateSubTasksOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Code *string `type:"string" json:",omitempty"`

	Message *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s SubTaskResultForVerifyMigrateSubTasksOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s SubTaskResultForVerifyMigrateSubTasksOutput) GoString() string {
	return s.String()
}

// SetCode sets the Code field's value.
func (s *SubTaskResultForVerifyMigrateSubTasksOutput) SetCode(v string) *SubTaskResultForVerifyMigrateSubTasksOutput {
	s.Code = &v
	return s
}

// SetMessage sets the Message field's value.
func (s *SubTaskResultForVerifyMigrateSubTasksOutput) SetMessage(v string) *SubTaskResultForVerifyMigrateSubTasksOutput {
	s.Message = &v
	return s
}

type SummaryForVerifyMigrateSubTasksOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Code *string `type:"string" json:",omitempty"`

	Message *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s SummaryForVerifyMigrateSubTasksOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s SummaryForVerifyMigrateSubTasksOutput) GoString() string {
	return s.String()
}

// SetCode sets the Code field's value.
func (s *SummaryForVerifyMigrateSubTasksOutput) SetCode(v string) *SummaryForVerifyMigrateSubTasksOutput {
	s.Code = &v
	return s
}

// SetMessage sets the Message field's value.
func (s *SummaryForVerifyMigrateSubTasksOutput) SetMessage(v string) *SummaryForVerifyMigrateSubTasksOutput {
	s.Message = &v
	return s
}

type VerifyMigrateSubTasksInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Items []*ItemForVerifyMigrateSubTasksInput `type:"list" json:",omitempty"`

	// TaskId is a required field
	TaskId *string `type:"string" json:",omitempty" required:"true"`
}

// String returns the string representation
func (s VerifyMigrateSubTasksInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s VerifyMigrateSubTasksInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *VerifyMigrateSubTasksInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "VerifyMigrateSubTasksInput"}
	if s.TaskId == nil {
		invalidParams.Add(request.NewErrParamRequired("TaskId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetItems sets the Items field's value.
func (s *VerifyMigrateSubTasksInput) SetItems(v []*ItemForVerifyMigrateSubTasksInput) *VerifyMigrateSubTasksInput {
	s.Items = v
	return s
}

// SetTaskId sets the TaskId field's value.
func (s *VerifyMigrateSubTasksInput) SetTaskId(v string) *VerifyMigrateSubTasksInput {
	s.TaskId = &v
	return s
}

type VerifyMigrateSubTasksOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	SubTaskResults []*SubTaskResultForVerifyMigrateSubTasksOutput `type:"list" json:",omitempty"`

	Summary *SummaryForVerifyMigrateSubTasksOutput `type:"structure" json:",omitempty"`
}

// String returns the string representation
func (s VerifyMigrateSubTasksOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s VerifyMigrateSubTasksOutput) GoString() string {
	return s.String()
}

// SetSubTaskResults sets the SubTaskResults field's value.
func (s *VerifyMigrateSubTasksOutput) SetSubTaskResults(v []*SubTaskResultForVerifyMigrateSubTasksOutput) *VerifyMigrateSubTasksOutput {
	s.SubTaskResults = v
	return s
}

// SetSummary sets the Summary field's value.
func (s *VerifyMigrateSubTasksOutput) SetSummary(v *SummaryForVerifyMigrateSubTasksOutput) *VerifyMigrateSubTasksOutput {
	s.Summary = v
	return s
}

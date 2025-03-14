// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package cdn

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribeContentTasksCommon = "DescribeContentTasks"

// DescribeContentTasksCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeContentTasksCommon operation. The "output" return
// value will be populated with the DescribeContentTasksCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeContentTasksCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeContentTasksCommon Send returns without error.
//
// See DescribeContentTasksCommon for more information on using the DescribeContentTasksCommon
// API call, and error handling.
//
//    // Example sending a request using the DescribeContentTasksCommonRequest method.
//    req, resp := client.DescribeContentTasksCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CDN) DescribeContentTasksCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeContentTasksCommon,
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

// DescribeContentTasksCommon API operation for CDN.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CDN's
// API operation DescribeContentTasksCommon for usage and error information.
func (c *CDN) DescribeContentTasksCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeContentTasksCommonRequest(input)
	return out, req.Send()
}

// DescribeContentTasksCommonWithContext is the same as DescribeContentTasksCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeContentTasksCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CDN) DescribeContentTasksCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeContentTasksCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeContentTasks = "DescribeContentTasks"

// DescribeContentTasksRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeContentTasks operation. The "output" return
// value will be populated with the DescribeContentTasksCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeContentTasksCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeContentTasksCommon Send returns without error.
//
// See DescribeContentTasks for more information on using the DescribeContentTasks
// API call, and error handling.
//
//    // Example sending a request using the DescribeContentTasksRequest method.
//    req, resp := client.DescribeContentTasksRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CDN) DescribeContentTasksRequest(input *DescribeContentTasksInput) (req *request.Request, output *DescribeContentTasksOutput) {
	op := &request.Operation{
		Name:       opDescribeContentTasks,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeContentTasksInput{}
	}

	output = &DescribeContentTasksOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// DescribeContentTasks API operation for CDN.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CDN's
// API operation DescribeContentTasks for usage and error information.
func (c *CDN) DescribeContentTasks(input *DescribeContentTasksInput) (*DescribeContentTasksOutput, error) {
	req, out := c.DescribeContentTasksRequest(input)
	return out, req.Send()
}

// DescribeContentTasksWithContext is the same as DescribeContentTasks with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeContentTasks for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CDN) DescribeContentTasksWithContext(ctx volcengine.Context, input *DescribeContentTasksInput, opts ...request.Option) (*DescribeContentTasksOutput, error) {
	req, out := c.DescribeContentTasksRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DataForDescribeContentTasksOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	CreateTime *int64 `type:"int64" json:",omitempty"`

	Delete *bool `type:"boolean" json:",omitempty"`

	Layer *string `type:"string" json:",omitempty"`

	Process *string `type:"string" json:",omitempty"`

	RefreshPrefix *bool `type:"boolean" json:",omitempty"`

	Remark *string `type:"string" json:",omitempty"`

	Status *string `type:"string" json:",omitempty"`

	TaskID *string `type:"string" json:",omitempty"`

	TaskType *string `type:"string" json:",omitempty"`

	Url *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s DataForDescribeContentTasksOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DataForDescribeContentTasksOutput) GoString() string {
	return s.String()
}

// SetCreateTime sets the CreateTime field's value.
func (s *DataForDescribeContentTasksOutput) SetCreateTime(v int64) *DataForDescribeContentTasksOutput {
	s.CreateTime = &v
	return s
}

// SetDelete sets the Delete field's value.
func (s *DataForDescribeContentTasksOutput) SetDelete(v bool) *DataForDescribeContentTasksOutput {
	s.Delete = &v
	return s
}

// SetLayer sets the Layer field's value.
func (s *DataForDescribeContentTasksOutput) SetLayer(v string) *DataForDescribeContentTasksOutput {
	s.Layer = &v
	return s
}

// SetProcess sets the Process field's value.
func (s *DataForDescribeContentTasksOutput) SetProcess(v string) *DataForDescribeContentTasksOutput {
	s.Process = &v
	return s
}

// SetRefreshPrefix sets the RefreshPrefix field's value.
func (s *DataForDescribeContentTasksOutput) SetRefreshPrefix(v bool) *DataForDescribeContentTasksOutput {
	s.RefreshPrefix = &v
	return s
}

// SetRemark sets the Remark field's value.
func (s *DataForDescribeContentTasksOutput) SetRemark(v string) *DataForDescribeContentTasksOutput {
	s.Remark = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *DataForDescribeContentTasksOutput) SetStatus(v string) *DataForDescribeContentTasksOutput {
	s.Status = &v
	return s
}

// SetTaskID sets the TaskID field's value.
func (s *DataForDescribeContentTasksOutput) SetTaskID(v string) *DataForDescribeContentTasksOutput {
	s.TaskID = &v
	return s
}

// SetTaskType sets the TaskType field's value.
func (s *DataForDescribeContentTasksOutput) SetTaskType(v string) *DataForDescribeContentTasksOutput {
	s.TaskType = &v
	return s
}

// SetUrl sets the Url field's value.
func (s *DataForDescribeContentTasksOutput) SetUrl(v string) *DataForDescribeContentTasksOutput {
	s.Url = &v
	return s
}

type DescribeContentTasksInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	DomainName *string `type:"string" json:",omitempty"`

	EndTime *int64 `type:"int64" json:",omitempty"`

	PageNum *int32 `type:"int32" json:",omitempty"`

	PageSize *int32 `type:"int32" json:",omitempty"`

	Remark *string `type:"string" json:",omitempty"`

	StartTime *int32 `type:"int32" json:",omitempty"`

	Status *string `type:"string" json:",omitempty"`

	TaskID *string `type:"string" json:",omitempty"`

	// TaskType is a required field
	TaskType *string `type:"string" json:",omitempty" required:"true"`

	Url *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s DescribeContentTasksInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeContentTasksInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeContentTasksInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DescribeContentTasksInput"}
	if s.TaskType == nil {
		invalidParams.Add(request.NewErrParamRequired("TaskType"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetDomainName sets the DomainName field's value.
func (s *DescribeContentTasksInput) SetDomainName(v string) *DescribeContentTasksInput {
	s.DomainName = &v
	return s
}

// SetEndTime sets the EndTime field's value.
func (s *DescribeContentTasksInput) SetEndTime(v int64) *DescribeContentTasksInput {
	s.EndTime = &v
	return s
}

// SetPageNum sets the PageNum field's value.
func (s *DescribeContentTasksInput) SetPageNum(v int32) *DescribeContentTasksInput {
	s.PageNum = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *DescribeContentTasksInput) SetPageSize(v int32) *DescribeContentTasksInput {
	s.PageSize = &v
	return s
}

// SetRemark sets the Remark field's value.
func (s *DescribeContentTasksInput) SetRemark(v string) *DescribeContentTasksInput {
	s.Remark = &v
	return s
}

// SetStartTime sets the StartTime field's value.
func (s *DescribeContentTasksInput) SetStartTime(v int32) *DescribeContentTasksInput {
	s.StartTime = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *DescribeContentTasksInput) SetStatus(v string) *DescribeContentTasksInput {
	s.Status = &v
	return s
}

// SetTaskID sets the TaskID field's value.
func (s *DescribeContentTasksInput) SetTaskID(v string) *DescribeContentTasksInput {
	s.TaskID = &v
	return s
}

// SetTaskType sets the TaskType field's value.
func (s *DescribeContentTasksInput) SetTaskType(v string) *DescribeContentTasksInput {
	s.TaskType = &v
	return s
}

// SetUrl sets the Url field's value.
func (s *DescribeContentTasksInput) SetUrl(v string) *DescribeContentTasksInput {
	s.Url = &v
	return s
}

type DescribeContentTasksOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	Data []*DataForDescribeContentTasksOutput `type:"list" json:",omitempty"`

	PageNum *int32 `type:"int32" json:",omitempty"`

	PageSize *int32 `type:"int32" json:",omitempty"`

	Total *int32 `type:"int32" json:",omitempty"`
}

// String returns the string representation
func (s DescribeContentTasksOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeContentTasksOutput) GoString() string {
	return s.String()
}

// SetData sets the Data field's value.
func (s *DescribeContentTasksOutput) SetData(v []*DataForDescribeContentTasksOutput) *DescribeContentTasksOutput {
	s.Data = v
	return s
}

// SetPageNum sets the PageNum field's value.
func (s *DescribeContentTasksOutput) SetPageNum(v int32) *DescribeContentTasksOutput {
	s.PageNum = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *DescribeContentTasksOutput) SetPageSize(v int32) *DescribeContentTasksOutput {
	s.PageSize = &v
	return s
}

// SetTotal sets the Total field's value.
func (s *DescribeContentTasksOutput) SetTotal(v int32) *DescribeContentTasksOutput {
	s.Total = &v
	return s
}

// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package mongodb

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribeNormalLogsCommon = "DescribeNormalLogs"

// DescribeNormalLogsCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeNormalLogsCommon operation. The "output" return
// value will be populated with the DescribeNormalLogsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeNormalLogsCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeNormalLogsCommon Send returns without error.
//
// See DescribeNormalLogsCommon for more information on using the DescribeNormalLogsCommon
// API call, and error handling.
//
//    // Example sending a request using the DescribeNormalLogsCommonRequest method.
//    req, resp := client.DescribeNormalLogsCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *MONGODB) DescribeNormalLogsCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeNormalLogsCommon,
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

// DescribeNormalLogsCommon API operation for MONGODB.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for MONGODB's
// API operation DescribeNormalLogsCommon for usage and error information.
func (c *MONGODB) DescribeNormalLogsCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeNormalLogsCommonRequest(input)
	return out, req.Send()
}

// DescribeNormalLogsCommonWithContext is the same as DescribeNormalLogsCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeNormalLogsCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *MONGODB) DescribeNormalLogsCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeNormalLogsCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeNormalLogs = "DescribeNormalLogs"

// DescribeNormalLogsRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeNormalLogs operation. The "output" return
// value will be populated with the DescribeNormalLogsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeNormalLogsCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeNormalLogsCommon Send returns without error.
//
// See DescribeNormalLogs for more information on using the DescribeNormalLogs
// API call, and error handling.
//
//    // Example sending a request using the DescribeNormalLogsRequest method.
//    req, resp := client.DescribeNormalLogsRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *MONGODB) DescribeNormalLogsRequest(input *DescribeNormalLogsInput) (req *request.Request, output *DescribeNormalLogsOutput) {
	op := &request.Operation{
		Name:       opDescribeNormalLogs,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeNormalLogsInput{}
	}

	output = &DescribeNormalLogsOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// DescribeNormalLogs API operation for MONGODB.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for MONGODB's
// API operation DescribeNormalLogs for usage and error information.
func (c *MONGODB) DescribeNormalLogs(input *DescribeNormalLogsInput) (*DescribeNormalLogsOutput, error) {
	req, out := c.DescribeNormalLogsRequest(input)
	return out, req.Send()
}

// DescribeNormalLogsWithContext is the same as DescribeNormalLogs with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeNormalLogs for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *MONGODB) DescribeNormalLogsWithContext(ctx volcengine.Context, input *DescribeNormalLogsInput, opts ...request.Option) (*DescribeNormalLogsOutput, error) {
	req, out := c.DescribeNormalLogsRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DataForDescribeNormalLogsOutput struct {
	_ struct{} `type:"structure"`

	Connection *string `type:"string"`

	CreateTime *string `type:"string"`

	LogType *string `type:"string"`

	Message *string `type:"string"`
}

// String returns the string representation
func (s DataForDescribeNormalLogsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DataForDescribeNormalLogsOutput) GoString() string {
	return s.String()
}

// SetConnection sets the Connection field's value.
func (s *DataForDescribeNormalLogsOutput) SetConnection(v string) *DataForDescribeNormalLogsOutput {
	s.Connection = &v
	return s
}

// SetCreateTime sets the CreateTime field's value.
func (s *DataForDescribeNormalLogsOutput) SetCreateTime(v string) *DataForDescribeNormalLogsOutput {
	s.CreateTime = &v
	return s
}

// SetLogType sets the LogType field's value.
func (s *DataForDescribeNormalLogsOutput) SetLogType(v string) *DataForDescribeNormalLogsOutput {
	s.LogType = &v
	return s
}

// SetMessage sets the Message field's value.
func (s *DataForDescribeNormalLogsOutput) SetMessage(v string) *DataForDescribeNormalLogsOutput {
	s.Message = &v
	return s
}

type DescribeNormalLogsInput struct {
	_ struct{} `type:"structure"`

	Context *string `type:"string"`

	// EndTime is a required field
	EndTime *int64 `type:"int64" required:"true"`

	// InstanceId is a required field
	InstanceId *string `type:"string" required:"true"`

	// Limit is a required field
	Limit *int64 `type:"int64" required:"true"`

	LogLevel *string `type:"string" enum:"EnumOfLogLevelForDescribeNormalLogsInput"`

	// PodName is a required field
	PodName *string `type:"string" required:"true"`

	Sort *string `type:"string" enum:"EnumOfSortForDescribeNormalLogsInput"`

	// StartTime is a required field
	StartTime *int64 `type:"int64" required:"true"`
}

// String returns the string representation
func (s DescribeNormalLogsInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeNormalLogsInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeNormalLogsInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DescribeNormalLogsInput"}
	if s.EndTime == nil {
		invalidParams.Add(request.NewErrParamRequired("EndTime"))
	}
	if s.InstanceId == nil {
		invalidParams.Add(request.NewErrParamRequired("InstanceId"))
	}
	if s.Limit == nil {
		invalidParams.Add(request.NewErrParamRequired("Limit"))
	}
	if s.PodName == nil {
		invalidParams.Add(request.NewErrParamRequired("PodName"))
	}
	if s.StartTime == nil {
		invalidParams.Add(request.NewErrParamRequired("StartTime"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetContext sets the Context field's value.
func (s *DescribeNormalLogsInput) SetContext(v string) *DescribeNormalLogsInput {
	s.Context = &v
	return s
}

// SetEndTime sets the EndTime field's value.
func (s *DescribeNormalLogsInput) SetEndTime(v int64) *DescribeNormalLogsInput {
	s.EndTime = &v
	return s
}

// SetInstanceId sets the InstanceId field's value.
func (s *DescribeNormalLogsInput) SetInstanceId(v string) *DescribeNormalLogsInput {
	s.InstanceId = &v
	return s
}

// SetLimit sets the Limit field's value.
func (s *DescribeNormalLogsInput) SetLimit(v int64) *DescribeNormalLogsInput {
	s.Limit = &v
	return s
}

// SetLogLevel sets the LogLevel field's value.
func (s *DescribeNormalLogsInput) SetLogLevel(v string) *DescribeNormalLogsInput {
	s.LogLevel = &v
	return s
}

// SetPodName sets the PodName field's value.
func (s *DescribeNormalLogsInput) SetPodName(v string) *DescribeNormalLogsInput {
	s.PodName = &v
	return s
}

// SetSort sets the Sort field's value.
func (s *DescribeNormalLogsInput) SetSort(v string) *DescribeNormalLogsInput {
	s.Sort = &v
	return s
}

// SetStartTime sets the StartTime field's value.
func (s *DescribeNormalLogsInput) SetStartTime(v int64) *DescribeNormalLogsInput {
	s.StartTime = &v
	return s
}

type DescribeNormalLogsOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	Context *string `type:"string"`

	Datas []*DataForDescribeNormalLogsOutput `type:"list"`

	ListOver *bool `type:"boolean"`

	Total *int32 `type:"int32"`
}

// String returns the string representation
func (s DescribeNormalLogsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeNormalLogsOutput) GoString() string {
	return s.String()
}

// SetContext sets the Context field's value.
func (s *DescribeNormalLogsOutput) SetContext(v string) *DescribeNormalLogsOutput {
	s.Context = &v
	return s
}

// SetDatas sets the Datas field's value.
func (s *DescribeNormalLogsOutput) SetDatas(v []*DataForDescribeNormalLogsOutput) *DescribeNormalLogsOutput {
	s.Datas = v
	return s
}

// SetListOver sets the ListOver field's value.
func (s *DescribeNormalLogsOutput) SetListOver(v bool) *DescribeNormalLogsOutput {
	s.ListOver = &v
	return s
}

// SetTotal sets the Total field's value.
func (s *DescribeNormalLogsOutput) SetTotal(v int32) *DescribeNormalLogsOutput {
	s.Total = &v
	return s
}

const (
	// EnumOfLogLevelForDescribeNormalLogsInputErrorLog is a EnumOfLogLevelForDescribeNormalLogsInput enum value
	EnumOfLogLevelForDescribeNormalLogsInputErrorLog = "ErrorLog"

	// EnumOfLogLevelForDescribeNormalLogsInputRunningLog is a EnumOfLogLevelForDescribeNormalLogsInput enum value
	EnumOfLogLevelForDescribeNormalLogsInputRunningLog = "RunningLog"
)

const (
	// EnumOfSortForDescribeNormalLogsInputAsc is a EnumOfSortForDescribeNormalLogsInput enum value
	EnumOfSortForDescribeNormalLogsInputAsc = "ASC"

	// EnumOfSortForDescribeNormalLogsInputDesc is a EnumOfSortForDescribeNormalLogsInput enum value
	EnumOfSortForDescribeNormalLogsInputDesc = "DESC"
)

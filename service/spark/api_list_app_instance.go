// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package spark

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opListAppInstanceCommon = "listAppInstance"

// ListAppInstanceCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ListAppInstanceCommon operation. The "output" return
// value will be populated with the ListAppInstanceCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListAppInstanceCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListAppInstanceCommon Send returns without error.
//
// See ListAppInstanceCommon for more information on using the ListAppInstanceCommon
// API call, and error handling.
//
//    // Example sending a request using the ListAppInstanceCommonRequest method.
//    req, resp := client.ListAppInstanceCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *SPARK) ListAppInstanceCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opListAppInstanceCommon,
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

// ListAppInstanceCommon API operation for SPARK.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for SPARK's
// API operation ListAppInstanceCommon for usage and error information.
func (c *SPARK) ListAppInstanceCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ListAppInstanceCommonRequest(input)
	return out, req.Send()
}

// ListAppInstanceCommonWithContext is the same as ListAppInstanceCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ListAppInstanceCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *SPARK) ListAppInstanceCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ListAppInstanceCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opListAppInstance = "listAppInstance"

// ListAppInstanceRequest generates a "volcengine/request.Request" representing the
// client's request for the ListAppInstance operation. The "output" return
// value will be populated with the ListAppInstanceCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListAppInstanceCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListAppInstanceCommon Send returns without error.
//
// See ListAppInstance for more information on using the ListAppInstance
// API call, and error handling.
//
//    // Example sending a request using the ListAppInstanceRequest method.
//    req, resp := client.ListAppInstanceRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *SPARK) ListAppInstanceRequest(input *ListAppInstanceInput) (req *request.Request, output *ListAppInstanceOutput) {
	op := &request.Operation{
		Name:       opListAppInstance,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListAppInstanceInput{}
	}

	output = &ListAppInstanceOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// ListAppInstance API operation for SPARK.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for SPARK's
// API operation ListAppInstance for usage and error information.
func (c *SPARK) ListAppInstance(input *ListAppInstanceInput) (*ListAppInstanceOutput, error) {
	req, out := c.ListAppInstanceRequest(input)
	return out, req.Send()
}

// ListAppInstanceWithContext is the same as ListAppInstance with the addition of
// the ability to pass a context and additional request options.
//
// See ListAppInstance for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *SPARK) ListAppInstanceWithContext(ctx volcengine.Context, input *ListAppInstanceInput, opts ...request.Option) (*ListAppInstanceOutput, error) {
	req, out := c.ListAppInstanceRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ListAppInstanceInput struct {
	_ struct{} `type:"structure"`

	ApplicationName *string `type:"string"`

	ApplicationTrn *string `type:"string"`

	EngineVersion *string `type:"string" enum:"EnumOfEngineVersionForlistAppInstanceInput"`

	InstanceId *int64 `type:"int64"`

	PageNum *int32 `type:"int32"`

	PageSize *int32 `type:"int32"`

	ProjectId *string `type:"string"`

	ResourcePoolTrn *string `type:"string"`

	SortField *string `type:"string"`

	SortOrder *string `type:"string"`

	State *string `type:"string" enum:"EnumOfStateForlistAppInstanceInput"`
}

// String returns the string representation
func (s ListAppInstanceInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListAppInstanceInput) GoString() string {
	return s.String()
}

// SetApplicationName sets the ApplicationName field's value.
func (s *ListAppInstanceInput) SetApplicationName(v string) *ListAppInstanceInput {
	s.ApplicationName = &v
	return s
}

// SetApplicationTrn sets the ApplicationTrn field's value.
func (s *ListAppInstanceInput) SetApplicationTrn(v string) *ListAppInstanceInput {
	s.ApplicationTrn = &v
	return s
}

// SetEngineVersion sets the EngineVersion field's value.
func (s *ListAppInstanceInput) SetEngineVersion(v string) *ListAppInstanceInput {
	s.EngineVersion = &v
	return s
}

// SetInstanceId sets the InstanceId field's value.
func (s *ListAppInstanceInput) SetInstanceId(v int64) *ListAppInstanceInput {
	s.InstanceId = &v
	return s
}

// SetPageNum sets the PageNum field's value.
func (s *ListAppInstanceInput) SetPageNum(v int32) *ListAppInstanceInput {
	s.PageNum = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *ListAppInstanceInput) SetPageSize(v int32) *ListAppInstanceInput {
	s.PageSize = &v
	return s
}

// SetProjectId sets the ProjectId field's value.
func (s *ListAppInstanceInput) SetProjectId(v string) *ListAppInstanceInput {
	s.ProjectId = &v
	return s
}

// SetResourcePoolTrn sets the ResourcePoolTrn field's value.
func (s *ListAppInstanceInput) SetResourcePoolTrn(v string) *ListAppInstanceInput {
	s.ResourcePoolTrn = &v
	return s
}

// SetSortField sets the SortField field's value.
func (s *ListAppInstanceInput) SetSortField(v string) *ListAppInstanceInput {
	s.SortField = &v
	return s
}

// SetSortOrder sets the SortOrder field's value.
func (s *ListAppInstanceInput) SetSortOrder(v string) *ListAppInstanceInput {
	s.SortOrder = &v
	return s
}

// SetState sets the State field's value.
func (s *ListAppInstanceInput) SetState(v string) *ListAppInstanceInput {
	s.State = &v
	return s
}

type ListAppInstanceOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	Current *int64 `type:"int64"`

	Records []*interface{} `type:"list"`

	Size *int64 `type:"int64"`

	Total *int64 `type:"int64"`
}

// String returns the string representation
func (s ListAppInstanceOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListAppInstanceOutput) GoString() string {
	return s.String()
}

// SetCurrent sets the Current field's value.
func (s *ListAppInstanceOutput) SetCurrent(v int64) *ListAppInstanceOutput {
	s.Current = &v
	return s
}

// SetRecords sets the Records field's value.
func (s *ListAppInstanceOutput) SetRecords(v []*interface{}) *ListAppInstanceOutput {
	s.Records = v
	return s
}

// SetSize sets the Size field's value.
func (s *ListAppInstanceOutput) SetSize(v int64) *ListAppInstanceOutput {
	s.Size = &v
	return s
}

// SetTotal sets the Total field's value.
func (s *ListAppInstanceOutput) SetTotal(v int64) *ListAppInstanceOutput {
	s.Total = &v
	return s
}

const (
	// EnumOfEngineVersionForlistAppInstanceInputFlinkVersion111 is a EnumOfEngineVersionForlistAppInstanceInput enum value
	EnumOfEngineVersionForlistAppInstanceInputFlinkVersion111 = "FLINK_VERSION_1_11"

	// EnumOfEngineVersionForlistAppInstanceInputFlinkVersion112 is a EnumOfEngineVersionForlistAppInstanceInput enum value
	EnumOfEngineVersionForlistAppInstanceInputFlinkVersion112 = "FLINK_VERSION_1_12"

	// EnumOfEngineVersionForlistAppInstanceInputFlinkVersion116 is a EnumOfEngineVersionForlistAppInstanceInput enum value
	EnumOfEngineVersionForlistAppInstanceInputFlinkVersion116 = "FLINK_VERSION_1_16"

	// EnumOfEngineVersionForlistAppInstanceInputSparkVersion301Os is a EnumOfEngineVersionForlistAppInstanceInput enum value
	EnumOfEngineVersionForlistAppInstanceInputSparkVersion301Os = "SPARK_VERSION_3_0_1_OS"

	// EnumOfEngineVersionForlistAppInstanceInputSparkVersion322 is a EnumOfEngineVersionForlistAppInstanceInput enum value
	EnumOfEngineVersionForlistAppInstanceInputSparkVersion322 = "SPARK_VERSION_3_2_2"

	// EnumOfEngineVersionForlistAppInstanceInputRayVersion113 is a EnumOfEngineVersionForlistAppInstanceInput enum value
	EnumOfEngineVersionForlistAppInstanceInputRayVersion113 = "RAY_VERSION_1_13"

	// EnumOfEngineVersionForlistAppInstanceInputPrimusVersion113 is a EnumOfEngineVersionForlistAppInstanceInput enum value
	EnumOfEngineVersionForlistAppInstanceInputPrimusVersion113 = "PRIMUS_VERSION_1_13"

	// EnumOfEngineVersionForlistAppInstanceInputUnknown is a EnumOfEngineVersionForlistAppInstanceInput enum value
	EnumOfEngineVersionForlistAppInstanceInputUnknown = "UNKNOWN"
)

const (
	// EnumOfStateForlistAppInstanceInputDeployed is a EnumOfStateForlistAppInstanceInput enum value
	EnumOfStateForlistAppInstanceInputDeployed = "DEPLOYED"

	// EnumOfStateForlistAppInstanceInputCreated is a EnumOfStateForlistAppInstanceInput enum value
	EnumOfStateForlistAppInstanceInputCreated = "CREATED"

	// EnumOfStateForlistAppInstanceInputStarting is a EnumOfStateForlistAppInstanceInput enum value
	EnumOfStateForlistAppInstanceInputStarting = "STARTING"

	// EnumOfStateForlistAppInstanceInputRunning is a EnumOfStateForlistAppInstanceInput enum value
	EnumOfStateForlistAppInstanceInputRunning = "RUNNING"

	// EnumOfStateForlistAppInstanceInputFailed is a EnumOfStateForlistAppInstanceInput enum value
	EnumOfStateForlistAppInstanceInputFailed = "FAILED"

	// EnumOfStateForlistAppInstanceInputCancelling is a EnumOfStateForlistAppInstanceInput enum value
	EnumOfStateForlistAppInstanceInputCancelling = "CANCELLING"

	// EnumOfStateForlistAppInstanceInputSucceeded is a EnumOfStateForlistAppInstanceInput enum value
	EnumOfStateForlistAppInstanceInputSucceeded = "SUCCEEDED"

	// EnumOfStateForlistAppInstanceInputStopped is a EnumOfStateForlistAppInstanceInput enum value
	EnumOfStateForlistAppInstanceInputStopped = "STOPPED"

	// EnumOfStateForlistAppInstanceInputUnknown is a EnumOfStateForlistAppInstanceInput enum value
	EnumOfStateForlistAppInstanceInputUnknown = "UNKNOWN"
)

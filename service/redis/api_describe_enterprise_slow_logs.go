// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package redis

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribeEnterpriseSlowLogsCommon = "DescribeEnterpriseSlowLogs"

// DescribeEnterpriseSlowLogsCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeEnterpriseSlowLogsCommon operation. The "output" return
// value will be populated with the DescribeEnterpriseSlowLogsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeEnterpriseSlowLogsCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeEnterpriseSlowLogsCommon Send returns without error.
//
// See DescribeEnterpriseSlowLogsCommon for more information on using the DescribeEnterpriseSlowLogsCommon
// API call, and error handling.
//
//    // Example sending a request using the DescribeEnterpriseSlowLogsCommonRequest method.
//    req, resp := client.DescribeEnterpriseSlowLogsCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *REDIS) DescribeEnterpriseSlowLogsCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeEnterpriseSlowLogsCommon,
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

// DescribeEnterpriseSlowLogsCommon API operation for REDIS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for REDIS's
// API operation DescribeEnterpriseSlowLogsCommon for usage and error information.
func (c *REDIS) DescribeEnterpriseSlowLogsCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeEnterpriseSlowLogsCommonRequest(input)
	return out, req.Send()
}

// DescribeEnterpriseSlowLogsCommonWithContext is the same as DescribeEnterpriseSlowLogsCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeEnterpriseSlowLogsCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *REDIS) DescribeEnterpriseSlowLogsCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeEnterpriseSlowLogsCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeEnterpriseSlowLogs = "DescribeEnterpriseSlowLogs"

// DescribeEnterpriseSlowLogsRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeEnterpriseSlowLogs operation. The "output" return
// value will be populated with the DescribeEnterpriseSlowLogsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeEnterpriseSlowLogsCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeEnterpriseSlowLogsCommon Send returns without error.
//
// See DescribeEnterpriseSlowLogs for more information on using the DescribeEnterpriseSlowLogs
// API call, and error handling.
//
//    // Example sending a request using the DescribeEnterpriseSlowLogsRequest method.
//    req, resp := client.DescribeEnterpriseSlowLogsRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *REDIS) DescribeEnterpriseSlowLogsRequest(input *DescribeEnterpriseSlowLogsInput) (req *request.Request, output *DescribeEnterpriseSlowLogsOutput) {
	op := &request.Operation{
		Name:       opDescribeEnterpriseSlowLogs,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeEnterpriseSlowLogsInput{}
	}

	output = &DescribeEnterpriseSlowLogsOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// DescribeEnterpriseSlowLogs API operation for REDIS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for REDIS's
// API operation DescribeEnterpriseSlowLogs for usage and error information.
func (c *REDIS) DescribeEnterpriseSlowLogs(input *DescribeEnterpriseSlowLogsInput) (*DescribeEnterpriseSlowLogsOutput, error) {
	req, out := c.DescribeEnterpriseSlowLogsRequest(input)
	return out, req.Send()
}

// DescribeEnterpriseSlowLogsWithContext is the same as DescribeEnterpriseSlowLogs with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeEnterpriseSlowLogs for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *REDIS) DescribeEnterpriseSlowLogsWithContext(ctx volcengine.Context, input *DescribeEnterpriseSlowLogsInput, opts ...request.Option) (*DescribeEnterpriseSlowLogsOutput, error) {
	req, out := c.DescribeEnterpriseSlowLogsRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DescribeEnterpriseSlowLogsInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	// InstanceId is a required field
	InstanceId *string `type:"string" json:",omitempty" required:"true"`
}

// String returns the string representation
func (s DescribeEnterpriseSlowLogsInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeEnterpriseSlowLogsInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeEnterpriseSlowLogsInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DescribeEnterpriseSlowLogsInput"}
	if s.InstanceId == nil {
		invalidParams.Add(request.NewErrParamRequired("InstanceId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetInstanceId sets the InstanceId field's value.
func (s *DescribeEnterpriseSlowLogsInput) SetInstanceId(v string) *DescribeEnterpriseSlowLogsInput {
	s.InstanceId = &v
	return s
}

type DescribeEnterpriseSlowLogsOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	SlowQuery []*SlowQueryForDescribeEnterpriseSlowLogsOutput `type:"list" json:",omitempty"`
}

// String returns the string representation
func (s DescribeEnterpriseSlowLogsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeEnterpriseSlowLogsOutput) GoString() string {
	return s.String()
}

// SetSlowQuery sets the SlowQuery field's value.
func (s *DescribeEnterpriseSlowLogsOutput) SetSlowQuery(v []*SlowQueryForDescribeEnterpriseSlowLogsOutput) *DescribeEnterpriseSlowLogsOutput {
	s.SlowQuery = v
	return s
}

type SlowQueryForDescribeEnterpriseSlowLogsOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	DBName *string `type:"string" json:",omitempty"`

	ExecutionStartTime *string `type:"string" json:",omitempty"`

	HostAddress *string `type:"string" json:",omitempty"`

	InstanceId *string `type:"string" json:",omitempty"`

	NodeId *string `type:"string" json:",omitempty"`

	QueryText *string `type:"string" json:",omitempty"`

	QueryTimes *int64 `type:"int64" json:",omitempty"`

	UserName *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s SlowQueryForDescribeEnterpriseSlowLogsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s SlowQueryForDescribeEnterpriseSlowLogsOutput) GoString() string {
	return s.String()
}

// SetDBName sets the DBName field's value.
func (s *SlowQueryForDescribeEnterpriseSlowLogsOutput) SetDBName(v string) *SlowQueryForDescribeEnterpriseSlowLogsOutput {
	s.DBName = &v
	return s
}

// SetExecutionStartTime sets the ExecutionStartTime field's value.
func (s *SlowQueryForDescribeEnterpriseSlowLogsOutput) SetExecutionStartTime(v string) *SlowQueryForDescribeEnterpriseSlowLogsOutput {
	s.ExecutionStartTime = &v
	return s
}

// SetHostAddress sets the HostAddress field's value.
func (s *SlowQueryForDescribeEnterpriseSlowLogsOutput) SetHostAddress(v string) *SlowQueryForDescribeEnterpriseSlowLogsOutput {
	s.HostAddress = &v
	return s
}

// SetInstanceId sets the InstanceId field's value.
func (s *SlowQueryForDescribeEnterpriseSlowLogsOutput) SetInstanceId(v string) *SlowQueryForDescribeEnterpriseSlowLogsOutput {
	s.InstanceId = &v
	return s
}

// SetNodeId sets the NodeId field's value.
func (s *SlowQueryForDescribeEnterpriseSlowLogsOutput) SetNodeId(v string) *SlowQueryForDescribeEnterpriseSlowLogsOutput {
	s.NodeId = &v
	return s
}

// SetQueryText sets the QueryText field's value.
func (s *SlowQueryForDescribeEnterpriseSlowLogsOutput) SetQueryText(v string) *SlowQueryForDescribeEnterpriseSlowLogsOutput {
	s.QueryText = &v
	return s
}

// SetQueryTimes sets the QueryTimes field's value.
func (s *SlowQueryForDescribeEnterpriseSlowLogsOutput) SetQueryTimes(v int64) *SlowQueryForDescribeEnterpriseSlowLogsOutput {
	s.QueryTimes = &v
	return s
}

// SetUserName sets the UserName field's value.
func (s *SlowQueryForDescribeEnterpriseSlowLogsOutput) SetUserName(v string) *SlowQueryForDescribeEnterpriseSlowLogsOutput {
	s.UserName = &v
	return s
}

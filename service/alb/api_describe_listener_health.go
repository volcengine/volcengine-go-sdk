// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package alb

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribeListenerHealthCommon = "DescribeListenerHealth"

// DescribeListenerHealthCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeListenerHealthCommon operation. The "output" return
// value will be populated with the DescribeListenerHealthCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeListenerHealthCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeListenerHealthCommon Send returns without error.
//
// See DescribeListenerHealthCommon for more information on using the DescribeListenerHealthCommon
// API call, and error handling.
//
//    // Example sending a request using the DescribeListenerHealthCommonRequest method.
//    req, resp := client.DescribeListenerHealthCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ALB) DescribeListenerHealthCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeListenerHealthCommon,
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

// DescribeListenerHealthCommon API operation for ALB.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ALB's
// API operation DescribeListenerHealthCommon for usage and error information.
func (c *ALB) DescribeListenerHealthCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeListenerHealthCommonRequest(input)
	return out, req.Send()
}

// DescribeListenerHealthCommonWithContext is the same as DescribeListenerHealthCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeListenerHealthCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ALB) DescribeListenerHealthCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeListenerHealthCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeListenerHealth = "DescribeListenerHealth"

// DescribeListenerHealthRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeListenerHealth operation. The "output" return
// value will be populated with the DescribeListenerHealthCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeListenerHealthCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeListenerHealthCommon Send returns without error.
//
// See DescribeListenerHealth for more information on using the DescribeListenerHealth
// API call, and error handling.
//
//    // Example sending a request using the DescribeListenerHealthRequest method.
//    req, resp := client.DescribeListenerHealthRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ALB) DescribeListenerHealthRequest(input *DescribeListenerHealthInput) (req *request.Request, output *DescribeListenerHealthOutput) {
	op := &request.Operation{
		Name:       opDescribeListenerHealth,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeListenerHealthInput{}
	}

	output = &DescribeListenerHealthOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DescribeListenerHealth API operation for ALB.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ALB's
// API operation DescribeListenerHealth for usage and error information.
func (c *ALB) DescribeListenerHealth(input *DescribeListenerHealthInput) (*DescribeListenerHealthOutput, error) {
	req, out := c.DescribeListenerHealthRequest(input)
	return out, req.Send()
}

// DescribeListenerHealthWithContext is the same as DescribeListenerHealth with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeListenerHealth for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ALB) DescribeListenerHealthWithContext(ctx volcengine.Context, input *DescribeListenerHealthInput, opts ...request.Option) (*DescribeListenerHealthOutput, error) {
	req, out := c.DescribeListenerHealthRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DescribeListenerHealthInput struct {
	_ struct{} `type:"structure"`

	ListenerIds []*string `type:"list"`

	OnlyUnHealthy *string `type:"string"`

	ProjectName *string `type:"string"`
}

// String returns the string representation
func (s DescribeListenerHealthInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeListenerHealthInput) GoString() string {
	return s.String()
}

// SetListenerIds sets the ListenerIds field's value.
func (s *DescribeListenerHealthInput) SetListenerIds(v []*string) *DescribeListenerHealthInput {
	s.ListenerIds = v
	return s
}

// SetOnlyUnHealthy sets the OnlyUnHealthy field's value.
func (s *DescribeListenerHealthInput) SetOnlyUnHealthy(v string) *DescribeListenerHealthInput {
	s.OnlyUnHealthy = &v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *DescribeListenerHealthInput) SetProjectName(v string) *DescribeListenerHealthInput {
	s.ProjectName = &v
	return s
}

type DescribeListenerHealthOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	Listeners []*ListenerForDescribeListenerHealthOutput `type:"list"`

	RequestId *string `type:"string"`

	TotalCount *int64 `type:"integer"`
}

// String returns the string representation
func (s DescribeListenerHealthOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeListenerHealthOutput) GoString() string {
	return s.String()
}

// SetListeners sets the Listeners field's value.
func (s *DescribeListenerHealthOutput) SetListeners(v []*ListenerForDescribeListenerHealthOutput) *DescribeListenerHealthOutput {
	s.Listeners = v
	return s
}

// SetRequestId sets the RequestId field's value.
func (s *DescribeListenerHealthOutput) SetRequestId(v string) *DescribeListenerHealthOutput {
	s.RequestId = &v
	return s
}

// SetTotalCount sets the TotalCount field's value.
func (s *DescribeListenerHealthOutput) SetTotalCount(v int64) *DescribeListenerHealthOutput {
	s.TotalCount = &v
	return s
}

type ListenerForDescribeListenerHealthOutput struct {
	_ struct{} `type:"structure"`

	ListenerId *string `type:"string"`

	Results []*ResultForDescribeListenerHealthOutput `type:"list"`

	Status *string `type:"string"`

	TotalBackendServerCount *int64 `type:"integer"`

	UnHealthyCount *int64 `type:"integer"`
}

// String returns the string representation
func (s ListenerForDescribeListenerHealthOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListenerForDescribeListenerHealthOutput) GoString() string {
	return s.String()
}

// SetListenerId sets the ListenerId field's value.
func (s *ListenerForDescribeListenerHealthOutput) SetListenerId(v string) *ListenerForDescribeListenerHealthOutput {
	s.ListenerId = &v
	return s
}

// SetResults sets the Results field's value.
func (s *ListenerForDescribeListenerHealthOutput) SetResults(v []*ResultForDescribeListenerHealthOutput) *ListenerForDescribeListenerHealthOutput {
	s.Results = v
	return s
}

// SetStatus sets the Status field's value.
func (s *ListenerForDescribeListenerHealthOutput) SetStatus(v string) *ListenerForDescribeListenerHealthOutput {
	s.Status = &v
	return s
}

// SetTotalBackendServerCount sets the TotalBackendServerCount field's value.
func (s *ListenerForDescribeListenerHealthOutput) SetTotalBackendServerCount(v int64) *ListenerForDescribeListenerHealthOutput {
	s.TotalBackendServerCount = &v
	return s
}

// SetUnHealthyCount sets the UnHealthyCount field's value.
func (s *ListenerForDescribeListenerHealthOutput) SetUnHealthyCount(v int64) *ListenerForDescribeListenerHealthOutput {
	s.UnHealthyCount = &v
	return s
}

type ResultForDescribeListenerHealthOutput struct {
	_ struct{} `type:"structure"`

	InstanceId *string `type:"string"`

	Ip *string `type:"string"`

	Port *int64 `type:"integer"`

	RuleNumber *int64 `type:"integer"`

	ServerGroupId *string `type:"string"`

	ServerGroupName *string `type:"string"`

	ServerId *string `type:"string"`

	Status *string `type:"string"`

	Type *string `type:"string"`
}

// String returns the string representation
func (s ResultForDescribeListenerHealthOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ResultForDescribeListenerHealthOutput) GoString() string {
	return s.String()
}

// SetInstanceId sets the InstanceId field's value.
func (s *ResultForDescribeListenerHealthOutput) SetInstanceId(v string) *ResultForDescribeListenerHealthOutput {
	s.InstanceId = &v
	return s
}

// SetIp sets the Ip field's value.
func (s *ResultForDescribeListenerHealthOutput) SetIp(v string) *ResultForDescribeListenerHealthOutput {
	s.Ip = &v
	return s
}

// SetPort sets the Port field's value.
func (s *ResultForDescribeListenerHealthOutput) SetPort(v int64) *ResultForDescribeListenerHealthOutput {
	s.Port = &v
	return s
}

// SetRuleNumber sets the RuleNumber field's value.
func (s *ResultForDescribeListenerHealthOutput) SetRuleNumber(v int64) *ResultForDescribeListenerHealthOutput {
	s.RuleNumber = &v
	return s
}

// SetServerGroupId sets the ServerGroupId field's value.
func (s *ResultForDescribeListenerHealthOutput) SetServerGroupId(v string) *ResultForDescribeListenerHealthOutput {
	s.ServerGroupId = &v
	return s
}

// SetServerGroupName sets the ServerGroupName field's value.
func (s *ResultForDescribeListenerHealthOutput) SetServerGroupName(v string) *ResultForDescribeListenerHealthOutput {
	s.ServerGroupName = &v
	return s
}

// SetServerId sets the ServerId field's value.
func (s *ResultForDescribeListenerHealthOutput) SetServerId(v string) *ResultForDescribeListenerHealthOutput {
	s.ServerId = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *ResultForDescribeListenerHealthOutput) SetStatus(v string) *ResultForDescribeListenerHealthOutput {
	s.Status = &v
	return s
}

// SetType sets the Type field's value.
func (s *ResultForDescribeListenerHealthOutput) SetType(v string) *ResultForDescribeListenerHealthOutput {
	s.Type = &v
	return s
}

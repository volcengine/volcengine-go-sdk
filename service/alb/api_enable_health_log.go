// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package alb

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opEnableHealthLogCommon = "EnableHealthLog"

// EnableHealthLogCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the EnableHealthLogCommon operation. The "output" return
// value will be populated with the EnableHealthLogCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned EnableHealthLogCommon Request to send the API call to the service.
// the "output" return value is not valid until after EnableHealthLogCommon Send returns without error.
//
// See EnableHealthLogCommon for more information on using the EnableHealthLogCommon
// API call, and error handling.
//
//    // Example sending a request using the EnableHealthLogCommonRequest method.
//    req, resp := client.EnableHealthLogCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ALB) EnableHealthLogCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opEnableHealthLogCommon,
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

// EnableHealthLogCommon API operation for ALB.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ALB's
// API operation EnableHealthLogCommon for usage and error information.
func (c *ALB) EnableHealthLogCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.EnableHealthLogCommonRequest(input)
	return out, req.Send()
}

// EnableHealthLogCommonWithContext is the same as EnableHealthLogCommon with the addition of
// the ability to pass a context and additional request options.
//
// See EnableHealthLogCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ALB) EnableHealthLogCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.EnableHealthLogCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opEnableHealthLog = "EnableHealthLog"

// EnableHealthLogRequest generates a "volcengine/request.Request" representing the
// client's request for the EnableHealthLog operation. The "output" return
// value will be populated with the EnableHealthLogCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned EnableHealthLogCommon Request to send the API call to the service.
// the "output" return value is not valid until after EnableHealthLogCommon Send returns without error.
//
// See EnableHealthLog for more information on using the EnableHealthLog
// API call, and error handling.
//
//    // Example sending a request using the EnableHealthLogRequest method.
//    req, resp := client.EnableHealthLogRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ALB) EnableHealthLogRequest(input *EnableHealthLogInput) (req *request.Request, output *EnableHealthLogOutput) {
	op := &request.Operation{
		Name:       opEnableHealthLog,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &EnableHealthLogInput{}
	}

	output = &EnableHealthLogOutput{}
	req = c.newRequest(op, input, output)

	return
}

// EnableHealthLog API operation for ALB.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ALB's
// API operation EnableHealthLog for usage and error information.
func (c *ALB) EnableHealthLog(input *EnableHealthLogInput) (*EnableHealthLogOutput, error) {
	req, out := c.EnableHealthLogRequest(input)
	return out, req.Send()
}

// EnableHealthLogWithContext is the same as EnableHealthLog with the addition of
// the ability to pass a context and additional request options.
//
// See EnableHealthLog for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ALB) EnableHealthLogWithContext(ctx volcengine.Context, input *EnableHealthLogInput, opts ...request.Option) (*EnableHealthLogOutput, error) {
	req, out := c.EnableHealthLogRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type EnableHealthLogInput struct {
	_ struct{} `type:"structure"`

	// LoadBalancerId is a required field
	LoadBalancerId *string `type:"string" required:"true"`

	// ProjectId is a required field
	ProjectId *string `type:"string" required:"true"`

	// TopicId is a required field
	TopicId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s EnableHealthLogInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s EnableHealthLogInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *EnableHealthLogInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "EnableHealthLogInput"}
	if s.LoadBalancerId == nil {
		invalidParams.Add(request.NewErrParamRequired("LoadBalancerId"))
	}
	if s.ProjectId == nil {
		invalidParams.Add(request.NewErrParamRequired("ProjectId"))
	}
	if s.TopicId == nil {
		invalidParams.Add(request.NewErrParamRequired("TopicId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetLoadBalancerId sets the LoadBalancerId field's value.
func (s *EnableHealthLogInput) SetLoadBalancerId(v string) *EnableHealthLogInput {
	s.LoadBalancerId = &v
	return s
}

// SetProjectId sets the ProjectId field's value.
func (s *EnableHealthLogInput) SetProjectId(v string) *EnableHealthLogInput {
	s.ProjectId = &v
	return s
}

// SetTopicId sets the TopicId field's value.
func (s *EnableHealthLogInput) SetTopicId(v string) *EnableHealthLogInput {
	s.TopicId = &v
	return s
}

type EnableHealthLogOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	RequestId *string `type:"string"`
}

// String returns the string representation
func (s EnableHealthLogOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s EnableHealthLogOutput) GoString() string {
	return s.String()
}

// SetRequestId sets the RequestId field's value.
func (s *EnableHealthLogOutput) SetRequestId(v string) *EnableHealthLogOutput {
	s.RequestId = &v
	return s
}

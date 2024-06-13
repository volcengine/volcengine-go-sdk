// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package autoscaling

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribeLifecycleHooksCommon = "DescribeLifecycleHooks"

// DescribeLifecycleHooksCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeLifecycleHooksCommon operation. The "output" return
// value will be populated with the DescribeLifecycleHooksCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeLifecycleHooksCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeLifecycleHooksCommon Send returns without error.
//
// See DescribeLifecycleHooksCommon for more information on using the DescribeLifecycleHooksCommon
// API call, and error handling.
//
//    // Example sending a request using the DescribeLifecycleHooksCommonRequest method.
//    req, resp := client.DescribeLifecycleHooksCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *AUTOSCALING) DescribeLifecycleHooksCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeLifecycleHooksCommon,
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

// DescribeLifecycleHooksCommon API operation for AUTO_SCALING.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for AUTO_SCALING's
// API operation DescribeLifecycleHooksCommon for usage and error information.
func (c *AUTOSCALING) DescribeLifecycleHooksCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeLifecycleHooksCommonRequest(input)
	return out, req.Send()
}

// DescribeLifecycleHooksCommonWithContext is the same as DescribeLifecycleHooksCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeLifecycleHooksCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *AUTOSCALING) DescribeLifecycleHooksCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeLifecycleHooksCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeLifecycleHooks = "DescribeLifecycleHooks"

// DescribeLifecycleHooksRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeLifecycleHooks operation. The "output" return
// value will be populated with the DescribeLifecycleHooksCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeLifecycleHooksCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeLifecycleHooksCommon Send returns without error.
//
// See DescribeLifecycleHooks for more information on using the DescribeLifecycleHooks
// API call, and error handling.
//
//    // Example sending a request using the DescribeLifecycleHooksRequest method.
//    req, resp := client.DescribeLifecycleHooksRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *AUTOSCALING) DescribeLifecycleHooksRequest(input *DescribeLifecycleHooksInput) (req *request.Request, output *DescribeLifecycleHooksOutput) {
	op := &request.Operation{
		Name:       opDescribeLifecycleHooks,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeLifecycleHooksInput{}
	}

	output = &DescribeLifecycleHooksOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DescribeLifecycleHooks API operation for AUTO_SCALING.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for AUTO_SCALING's
// API operation DescribeLifecycleHooks for usage and error information.
func (c *AUTOSCALING) DescribeLifecycleHooks(input *DescribeLifecycleHooksInput) (*DescribeLifecycleHooksOutput, error) {
	req, out := c.DescribeLifecycleHooksRequest(input)
	return out, req.Send()
}

// DescribeLifecycleHooksWithContext is the same as DescribeLifecycleHooks with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeLifecycleHooks for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *AUTOSCALING) DescribeLifecycleHooksWithContext(ctx volcengine.Context, input *DescribeLifecycleHooksInput, opts ...request.Option) (*DescribeLifecycleHooksOutput, error) {
	req, out := c.DescribeLifecycleHooksRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DescribeLifecycleHooksInput struct {
	_ struct{} `type:"structure"`

	LifecycleHookIds []*string `type:"list"`

	LifecycleHookName *string `type:"string"`

	PageNumber *int32 `type:"int32"`

	PageSize *int32 `type:"int32"`

	// ScalingGroupId is a required field
	ScalingGroupId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeLifecycleHooksInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeLifecycleHooksInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeLifecycleHooksInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DescribeLifecycleHooksInput"}
	if s.ScalingGroupId == nil {
		invalidParams.Add(request.NewErrParamRequired("ScalingGroupId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetLifecycleHookIds sets the LifecycleHookIds field's value.
func (s *DescribeLifecycleHooksInput) SetLifecycleHookIds(v []*string) *DescribeLifecycleHooksInput {
	s.LifecycleHookIds = v
	return s
}

// SetLifecycleHookName sets the LifecycleHookName field's value.
func (s *DescribeLifecycleHooksInput) SetLifecycleHookName(v string) *DescribeLifecycleHooksInput {
	s.LifecycleHookName = &v
	return s
}

// SetPageNumber sets the PageNumber field's value.
func (s *DescribeLifecycleHooksInput) SetPageNumber(v int32) *DescribeLifecycleHooksInput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *DescribeLifecycleHooksInput) SetPageSize(v int32) *DescribeLifecycleHooksInput {
	s.PageSize = &v
	return s
}

// SetScalingGroupId sets the ScalingGroupId field's value.
func (s *DescribeLifecycleHooksInput) SetScalingGroupId(v string) *DescribeLifecycleHooksInput {
	s.ScalingGroupId = &v
	return s
}

type DescribeLifecycleHooksOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	LifecycleHooks []*LifecycleHookForDescribeLifecycleHooksOutput `type:"list"`

	PageNumber *int32 `type:"int32"`

	PageSize *int32 `type:"int32"`

	TotalCount *int32 `type:"int32"`
}

// String returns the string representation
func (s DescribeLifecycleHooksOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeLifecycleHooksOutput) GoString() string {
	return s.String()
}

// SetLifecycleHooks sets the LifecycleHooks field's value.
func (s *DescribeLifecycleHooksOutput) SetLifecycleHooks(v []*LifecycleHookForDescribeLifecycleHooksOutput) *DescribeLifecycleHooksOutput {
	s.LifecycleHooks = v
	return s
}

// SetPageNumber sets the PageNumber field's value.
func (s *DescribeLifecycleHooksOutput) SetPageNumber(v int32) *DescribeLifecycleHooksOutput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *DescribeLifecycleHooksOutput) SetPageSize(v int32) *DescribeLifecycleHooksOutput {
	s.PageSize = &v
	return s
}

// SetTotalCount sets the TotalCount field's value.
func (s *DescribeLifecycleHooksOutput) SetTotalCount(v int32) *DescribeLifecycleHooksOutput {
	s.TotalCount = &v
	return s
}

type LifecycleCommandForDescribeLifecycleHooksOutput struct {
	_ struct{} `type:"structure"`

	CommandId *string `type:"string"`

	Parameters *string `type:"string"`
}

// String returns the string representation
func (s LifecycleCommandForDescribeLifecycleHooksOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s LifecycleCommandForDescribeLifecycleHooksOutput) GoString() string {
	return s.String()
}

// SetCommandId sets the CommandId field's value.
func (s *LifecycleCommandForDescribeLifecycleHooksOutput) SetCommandId(v string) *LifecycleCommandForDescribeLifecycleHooksOutput {
	s.CommandId = &v
	return s
}

// SetParameters sets the Parameters field's value.
func (s *LifecycleCommandForDescribeLifecycleHooksOutput) SetParameters(v string) *LifecycleCommandForDescribeLifecycleHooksOutput {
	s.Parameters = &v
	return s
}

type LifecycleHookForDescribeLifecycleHooksOutput struct {
	_ struct{} `type:"structure"`

	LifecycleCommand *LifecycleCommandForDescribeLifecycleHooksOutput `type:"structure"`

	LifecycleHookId *string `type:"string"`

	LifecycleHookName *string `type:"string"`

	LifecycleHookPolicy *string `type:"string"`

	LifecycleHookTimeout *int32 `type:"int32"`

	LifecycleHookType *string `type:"string"`

	ScalingGroupId *string `type:"string"`
}

// String returns the string representation
func (s LifecycleHookForDescribeLifecycleHooksOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s LifecycleHookForDescribeLifecycleHooksOutput) GoString() string {
	return s.String()
}

// SetLifecycleCommand sets the LifecycleCommand field's value.
func (s *LifecycleHookForDescribeLifecycleHooksOutput) SetLifecycleCommand(v *LifecycleCommandForDescribeLifecycleHooksOutput) *LifecycleHookForDescribeLifecycleHooksOutput {
	s.LifecycleCommand = v
	return s
}

// SetLifecycleHookId sets the LifecycleHookId field's value.
func (s *LifecycleHookForDescribeLifecycleHooksOutput) SetLifecycleHookId(v string) *LifecycleHookForDescribeLifecycleHooksOutput {
	s.LifecycleHookId = &v
	return s
}

// SetLifecycleHookName sets the LifecycleHookName field's value.
func (s *LifecycleHookForDescribeLifecycleHooksOutput) SetLifecycleHookName(v string) *LifecycleHookForDescribeLifecycleHooksOutput {
	s.LifecycleHookName = &v
	return s
}

// SetLifecycleHookPolicy sets the LifecycleHookPolicy field's value.
func (s *LifecycleHookForDescribeLifecycleHooksOutput) SetLifecycleHookPolicy(v string) *LifecycleHookForDescribeLifecycleHooksOutput {
	s.LifecycleHookPolicy = &v
	return s
}

// SetLifecycleHookTimeout sets the LifecycleHookTimeout field's value.
func (s *LifecycleHookForDescribeLifecycleHooksOutput) SetLifecycleHookTimeout(v int32) *LifecycleHookForDescribeLifecycleHooksOutput {
	s.LifecycleHookTimeout = &v
	return s
}

// SetLifecycleHookType sets the LifecycleHookType field's value.
func (s *LifecycleHookForDescribeLifecycleHooksOutput) SetLifecycleHookType(v string) *LifecycleHookForDescribeLifecycleHooksOutput {
	s.LifecycleHookType = &v
	return s
}

// SetScalingGroupId sets the ScalingGroupId field's value.
func (s *LifecycleHookForDescribeLifecycleHooksOutput) SetScalingGroupId(v string) *LifecycleHookForDescribeLifecycleHooksOutput {
	s.ScalingGroupId = &v
	return s
}

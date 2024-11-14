// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package escloud

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opRestartInstanceCommon = "RestartInstance"

// RestartInstanceCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the RestartInstanceCommon operation. The "output" return
// value will be populated with the RestartInstanceCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned RestartInstanceCommon Request to send the API call to the service.
// the "output" return value is not valid until after RestartInstanceCommon Send returns without error.
//
// See RestartInstanceCommon for more information on using the RestartInstanceCommon
// API call, and error handling.
//
//    // Example sending a request using the RestartInstanceCommonRequest method.
//    req, resp := client.RestartInstanceCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ESCLOUD) RestartInstanceCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opRestartInstanceCommon,
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

// RestartInstanceCommon API operation for ESCLOUD.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ESCLOUD's
// API operation RestartInstanceCommon for usage and error information.
func (c *ESCLOUD) RestartInstanceCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.RestartInstanceCommonRequest(input)
	return out, req.Send()
}

// RestartInstanceCommonWithContext is the same as RestartInstanceCommon with the addition of
// the ability to pass a context and additional request options.
//
// See RestartInstanceCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ESCLOUD) RestartInstanceCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.RestartInstanceCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opRestartInstance = "RestartInstance"

// RestartInstanceRequest generates a "volcengine/request.Request" representing the
// client's request for the RestartInstance operation. The "output" return
// value will be populated with the RestartInstanceCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned RestartInstanceCommon Request to send the API call to the service.
// the "output" return value is not valid until after RestartInstanceCommon Send returns without error.
//
// See RestartInstance for more information on using the RestartInstance
// API call, and error handling.
//
//    // Example sending a request using the RestartInstanceRequest method.
//    req, resp := client.RestartInstanceRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ESCLOUD) RestartInstanceRequest(input *RestartInstanceInput) (req *request.Request, output *RestartInstanceOutput) {
	op := &request.Operation{
		Name:       opRestartInstance,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &RestartInstanceInput{}
	}

	output = &RestartInstanceOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// RestartInstance API operation for ESCLOUD.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ESCLOUD's
// API operation RestartInstance for usage and error information.
func (c *ESCLOUD) RestartInstance(input *RestartInstanceInput) (*RestartInstanceOutput, error) {
	req, out := c.RestartInstanceRequest(input)
	return out, req.Send()
}

// RestartInstanceWithContext is the same as RestartInstance with the addition of
// the ability to pass a context and additional request options.
//
// See RestartInstance for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ESCLOUD) RestartInstanceWithContext(ctx volcengine.Context, input *RestartInstanceInput, opts ...request.Option) (*RestartInstanceOutput, error) {
	req, out := c.RestartInstanceRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type RestartInstanceInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Force *bool `type:"boolean" json:",omitempty"`

	// InstanceId is a required field
	InstanceId *string `type:"string" json:",omitempty" required:"true"`
}

// String returns the string representation
func (s RestartInstanceInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s RestartInstanceInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RestartInstanceInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "RestartInstanceInput"}
	if s.InstanceId == nil {
		invalidParams.Add(request.NewErrParamRequired("InstanceId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetForce sets the Force field's value.
func (s *RestartInstanceInput) SetForce(v bool) *RestartInstanceInput {
	s.Force = &v
	return s
}

// SetInstanceId sets the InstanceId field's value.
func (s *RestartInstanceInput) SetInstanceId(v string) *RestartInstanceInput {
	s.InstanceId = &v
	return s
}

type RestartInstanceOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	InstanceId *string `type:"string" json:",omitempty"`

	TaskId *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s RestartInstanceOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s RestartInstanceOutput) GoString() string {
	return s.String()
}

// SetInstanceId sets the InstanceId field's value.
func (s *RestartInstanceOutput) SetInstanceId(v string) *RestartInstanceOutput {
	s.InstanceId = &v
	return s
}

// SetTaskId sets the TaskId field's value.
func (s *RestartInstanceOutput) SetTaskId(v string) *RestartInstanceOutput {
	s.TaskId = &v
	return s
}

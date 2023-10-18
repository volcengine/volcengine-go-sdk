// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package vedbm

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opRestartDBInstanceCommon = "RestartDBInstance"

// RestartDBInstanceCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the RestartDBInstanceCommon operation. The "output" return
// value will be populated with the RestartDBInstanceCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned RestartDBInstanceCommon Request to send the API call to the service.
// the "output" return value is not valid until after RestartDBInstanceCommon Send returns without error.
//
// See RestartDBInstanceCommon for more information on using the RestartDBInstanceCommon
// API call, and error handling.
//
//    // Example sending a request using the RestartDBInstanceCommonRequest method.
//    req, resp := client.RestartDBInstanceCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VEDBM) RestartDBInstanceCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opRestartDBInstanceCommon,
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

// RestartDBInstanceCommon API operation for VEDBM.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VEDBM's
// API operation RestartDBInstanceCommon for usage and error information.
func (c *VEDBM) RestartDBInstanceCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.RestartDBInstanceCommonRequest(input)
	return out, req.Send()
}

// RestartDBInstanceCommonWithContext is the same as RestartDBInstanceCommon with the addition of
// the ability to pass a context and additional request options.
//
// See RestartDBInstanceCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VEDBM) RestartDBInstanceCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.RestartDBInstanceCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opRestartDBInstance = "RestartDBInstance"

// RestartDBInstanceRequest generates a "volcengine/request.Request" representing the
// client's request for the RestartDBInstance operation. The "output" return
// value will be populated with the RestartDBInstanceCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned RestartDBInstanceCommon Request to send the API call to the service.
// the "output" return value is not valid until after RestartDBInstanceCommon Send returns without error.
//
// See RestartDBInstance for more information on using the RestartDBInstance
// API call, and error handling.
//
//    // Example sending a request using the RestartDBInstanceRequest method.
//    req, resp := client.RestartDBInstanceRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VEDBM) RestartDBInstanceRequest(input *RestartDBInstanceInput) (req *request.Request, output *RestartDBInstanceOutput) {
	op := &request.Operation{
		Name:       opRestartDBInstance,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &RestartDBInstanceInput{}
	}

	output = &RestartDBInstanceOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// RestartDBInstance API operation for VEDBM.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VEDBM's
// API operation RestartDBInstance for usage and error information.
func (c *VEDBM) RestartDBInstance(input *RestartDBInstanceInput) (*RestartDBInstanceOutput, error) {
	req, out := c.RestartDBInstanceRequest(input)
	return out, req.Send()
}

// RestartDBInstanceWithContext is the same as RestartDBInstance with the addition of
// the ability to pass a context and additional request options.
//
// See RestartDBInstance for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VEDBM) RestartDBInstanceWithContext(ctx volcengine.Context, input *RestartDBInstanceInput, opts ...request.Option) (*RestartDBInstanceOutput, error) {
	req, out := c.RestartDBInstanceRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type RestartDBInstanceInput struct {
	_ struct{} `type:"structure"`

	InstanceId *string `type:"string"`
}

// String returns the string representation
func (s RestartDBInstanceInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s RestartDBInstanceInput) GoString() string {
	return s.String()
}

// SetInstanceId sets the InstanceId field's value.
func (s *RestartDBInstanceInput) SetInstanceId(v string) *RestartDBInstanceInput {
	s.InstanceId = &v
	return s
}

type RestartDBInstanceOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s RestartDBInstanceOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s RestartDBInstanceOutput) GoString() string {
	return s.String()
}

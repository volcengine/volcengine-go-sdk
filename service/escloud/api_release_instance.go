// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package escloud

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opReleaseInstanceCommon = "ReleaseInstance"

// ReleaseInstanceCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ReleaseInstanceCommon operation. The "output" return
// value will be populated with the ReleaseInstanceCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ReleaseInstanceCommon Request to send the API call to the service.
// the "output" return value is not valid until after ReleaseInstanceCommon Send returns without error.
//
// See ReleaseInstanceCommon for more information on using the ReleaseInstanceCommon
// API call, and error handling.
//
//    // Example sending a request using the ReleaseInstanceCommonRequest method.
//    req, resp := client.ReleaseInstanceCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ESCLOUD) ReleaseInstanceCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opReleaseInstanceCommon,
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

// ReleaseInstanceCommon API operation for ESCLOUD.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ESCLOUD's
// API operation ReleaseInstanceCommon for usage and error information.
func (c *ESCLOUD) ReleaseInstanceCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ReleaseInstanceCommonRequest(input)
	return out, req.Send()
}

// ReleaseInstanceCommonWithContext is the same as ReleaseInstanceCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ReleaseInstanceCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ESCLOUD) ReleaseInstanceCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ReleaseInstanceCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opReleaseInstance = "ReleaseInstance"

// ReleaseInstanceRequest generates a "volcengine/request.Request" representing the
// client's request for the ReleaseInstance operation. The "output" return
// value will be populated with the ReleaseInstanceCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ReleaseInstanceCommon Request to send the API call to the service.
// the "output" return value is not valid until after ReleaseInstanceCommon Send returns without error.
//
// See ReleaseInstance for more information on using the ReleaseInstance
// API call, and error handling.
//
//    // Example sending a request using the ReleaseInstanceRequest method.
//    req, resp := client.ReleaseInstanceRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ESCLOUD) ReleaseInstanceRequest(input *ReleaseInstanceInput) (req *request.Request, output *ReleaseInstanceOutput) {
	op := &request.Operation{
		Name:       opReleaseInstance,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ReleaseInstanceInput{}
	}

	output = &ReleaseInstanceOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// ReleaseInstance API operation for ESCLOUD.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ESCLOUD's
// API operation ReleaseInstance for usage and error information.
func (c *ESCLOUD) ReleaseInstance(input *ReleaseInstanceInput) (*ReleaseInstanceOutput, error) {
	req, out := c.ReleaseInstanceRequest(input)
	return out, req.Send()
}

// ReleaseInstanceWithContext is the same as ReleaseInstance with the addition of
// the ability to pass a context and additional request options.
//
// See ReleaseInstance for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ESCLOUD) ReleaseInstanceWithContext(ctx volcengine.Context, input *ReleaseInstanceInput, opts ...request.Option) (*ReleaseInstanceOutput, error) {
	req, out := c.ReleaseInstanceRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ReleaseInstanceInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	// InstanceId is a required field
	InstanceId *string `type:"string" json:",omitempty" required:"true"`
}

// String returns the string representation
func (s ReleaseInstanceInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ReleaseInstanceInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ReleaseInstanceInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ReleaseInstanceInput"}
	if s.InstanceId == nil {
		invalidParams.Add(request.NewErrParamRequired("InstanceId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetInstanceId sets the InstanceId field's value.
func (s *ReleaseInstanceInput) SetInstanceId(v string) *ReleaseInstanceInput {
	s.InstanceId = &v
	return s
}

type ReleaseInstanceOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	Message *string `type:"string" json:",omitempty"`

	Status *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s ReleaseInstanceOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ReleaseInstanceOutput) GoString() string {
	return s.String()
}

// SetMessage sets the Message field's value.
func (s *ReleaseInstanceOutput) SetMessage(v string) *ReleaseInstanceOutput {
	s.Message = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *ReleaseInstanceOutput) SetStatus(v string) *ReleaseInstanceOutput {
	s.Status = &v
	return s
}
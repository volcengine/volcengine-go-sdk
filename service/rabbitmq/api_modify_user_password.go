// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package rabbitmq

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opModifyUserPasswordCommon = "ModifyUserPassword"

// ModifyUserPasswordCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ModifyUserPasswordCommon operation. The "output" return
// value will be populated with the ModifyUserPasswordCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyUserPasswordCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyUserPasswordCommon Send returns without error.
//
// See ModifyUserPasswordCommon for more information on using the ModifyUserPasswordCommon
// API call, and error handling.
//
//    // Example sending a request using the ModifyUserPasswordCommonRequest method.
//    req, resp := client.ModifyUserPasswordCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *RABBITMQ) ModifyUserPasswordCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opModifyUserPasswordCommon,
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

// ModifyUserPasswordCommon API operation for RABBITMQ.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RABBITMQ's
// API operation ModifyUserPasswordCommon for usage and error information.
func (c *RABBITMQ) ModifyUserPasswordCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ModifyUserPasswordCommonRequest(input)
	return out, req.Send()
}

// ModifyUserPasswordCommonWithContext is the same as ModifyUserPasswordCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyUserPasswordCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RABBITMQ) ModifyUserPasswordCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ModifyUserPasswordCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opModifyUserPassword = "ModifyUserPassword"

// ModifyUserPasswordRequest generates a "volcengine/request.Request" representing the
// client's request for the ModifyUserPassword operation. The "output" return
// value will be populated with the ModifyUserPasswordCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyUserPasswordCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyUserPasswordCommon Send returns without error.
//
// See ModifyUserPassword for more information on using the ModifyUserPassword
// API call, and error handling.
//
//    // Example sending a request using the ModifyUserPasswordRequest method.
//    req, resp := client.ModifyUserPasswordRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *RABBITMQ) ModifyUserPasswordRequest(input *ModifyUserPasswordInput) (req *request.Request, output *ModifyUserPasswordOutput) {
	op := &request.Operation{
		Name:       opModifyUserPassword,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ModifyUserPasswordInput{}
	}

	output = &ModifyUserPasswordOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// ModifyUserPassword API operation for RABBITMQ.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RABBITMQ's
// API operation ModifyUserPassword for usage and error information.
func (c *RABBITMQ) ModifyUserPassword(input *ModifyUserPasswordInput) (*ModifyUserPasswordOutput, error) {
	req, out := c.ModifyUserPasswordRequest(input)
	return out, req.Send()
}

// ModifyUserPasswordWithContext is the same as ModifyUserPassword with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyUserPassword for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RABBITMQ) ModifyUserPasswordWithContext(ctx volcengine.Context, input *ModifyUserPasswordInput, opts ...request.Option) (*ModifyUserPasswordOutput, error) {
	req, out := c.ModifyUserPasswordRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ModifyUserPasswordInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	// InstanceId is a required field
	InstanceId *string `type:"string" json:",omitempty" required:"true"`

	// Password is a required field
	Password *string `type:"string" json:",omitempty" required:"true"`

	// UserName is a required field
	UserName *string `type:"string" json:",omitempty" required:"true"`
}

// String returns the string representation
func (s ModifyUserPasswordInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyUserPasswordInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ModifyUserPasswordInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ModifyUserPasswordInput"}
	if s.InstanceId == nil {
		invalidParams.Add(request.NewErrParamRequired("InstanceId"))
	}
	if s.Password == nil {
		invalidParams.Add(request.NewErrParamRequired("Password"))
	}
	if s.UserName == nil {
		invalidParams.Add(request.NewErrParamRequired("UserName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetInstanceId sets the InstanceId field's value.
func (s *ModifyUserPasswordInput) SetInstanceId(v string) *ModifyUserPasswordInput {
	s.InstanceId = &v
	return s
}

// SetPassword sets the Password field's value.
func (s *ModifyUserPasswordInput) SetPassword(v string) *ModifyUserPasswordInput {
	s.Password = &v
	return s
}

// SetUserName sets the UserName field's value.
func (s *ModifyUserPasswordInput) SetUserName(v string) *ModifyUserPasswordInput {
	s.UserName = &v
	return s
}

type ModifyUserPasswordOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s ModifyUserPasswordOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyUserPasswordOutput) GoString() string {
	return s.String()
}
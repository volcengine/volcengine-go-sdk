// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package redis

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opModifyDBInstanceVpcAuthModeCommon = "ModifyDBInstanceVpcAuthMode"

// ModifyDBInstanceVpcAuthModeCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ModifyDBInstanceVpcAuthModeCommon operation. The "output" return
// value will be populated with the ModifyDBInstanceVpcAuthModeCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyDBInstanceVpcAuthModeCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyDBInstanceVpcAuthModeCommon Send returns without error.
//
// See ModifyDBInstanceVpcAuthModeCommon for more information on using the ModifyDBInstanceVpcAuthModeCommon
// API call, and error handling.
//
//    // Example sending a request using the ModifyDBInstanceVpcAuthModeCommonRequest method.
//    req, resp := client.ModifyDBInstanceVpcAuthModeCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *REDIS) ModifyDBInstanceVpcAuthModeCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opModifyDBInstanceVpcAuthModeCommon,
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

// ModifyDBInstanceVpcAuthModeCommon API operation for REDIS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for REDIS's
// API operation ModifyDBInstanceVpcAuthModeCommon for usage and error information.
func (c *REDIS) ModifyDBInstanceVpcAuthModeCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ModifyDBInstanceVpcAuthModeCommonRequest(input)
	return out, req.Send()
}

// ModifyDBInstanceVpcAuthModeCommonWithContext is the same as ModifyDBInstanceVpcAuthModeCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyDBInstanceVpcAuthModeCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *REDIS) ModifyDBInstanceVpcAuthModeCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ModifyDBInstanceVpcAuthModeCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opModifyDBInstanceVpcAuthMode = "ModifyDBInstanceVpcAuthMode"

// ModifyDBInstanceVpcAuthModeRequest generates a "volcengine/request.Request" representing the
// client's request for the ModifyDBInstanceVpcAuthMode operation. The "output" return
// value will be populated with the ModifyDBInstanceVpcAuthModeCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyDBInstanceVpcAuthModeCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyDBInstanceVpcAuthModeCommon Send returns without error.
//
// See ModifyDBInstanceVpcAuthMode for more information on using the ModifyDBInstanceVpcAuthMode
// API call, and error handling.
//
//    // Example sending a request using the ModifyDBInstanceVpcAuthModeRequest method.
//    req, resp := client.ModifyDBInstanceVpcAuthModeRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *REDIS) ModifyDBInstanceVpcAuthModeRequest(input *ModifyDBInstanceVpcAuthModeInput) (req *request.Request, output *ModifyDBInstanceVpcAuthModeOutput) {
	op := &request.Operation{
		Name:       opModifyDBInstanceVpcAuthMode,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ModifyDBInstanceVpcAuthModeInput{}
	}

	output = &ModifyDBInstanceVpcAuthModeOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// ModifyDBInstanceVpcAuthMode API operation for REDIS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for REDIS's
// API operation ModifyDBInstanceVpcAuthMode for usage and error information.
func (c *REDIS) ModifyDBInstanceVpcAuthMode(input *ModifyDBInstanceVpcAuthModeInput) (*ModifyDBInstanceVpcAuthModeOutput, error) {
	req, out := c.ModifyDBInstanceVpcAuthModeRequest(input)
	return out, req.Send()
}

// ModifyDBInstanceVpcAuthModeWithContext is the same as ModifyDBInstanceVpcAuthMode with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyDBInstanceVpcAuthMode for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *REDIS) ModifyDBInstanceVpcAuthModeWithContext(ctx volcengine.Context, input *ModifyDBInstanceVpcAuthModeInput, opts ...request.Option) (*ModifyDBInstanceVpcAuthModeOutput, error) {
	req, out := c.ModifyDBInstanceVpcAuthModeRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ModifyDBInstanceVpcAuthModeInput struct {
	_ struct{} `type:"structure"`

	ClientToken *string `type:"string"`

	// InstanceId is a required field
	InstanceId *string `type:"string" required:"true"`

	// VpcAuthMode is a required field
	VpcAuthMode *string `type:"string" required:"true"`
}

// String returns the string representation
func (s ModifyDBInstanceVpcAuthModeInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyDBInstanceVpcAuthModeInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ModifyDBInstanceVpcAuthModeInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ModifyDBInstanceVpcAuthModeInput"}
	if s.InstanceId == nil {
		invalidParams.Add(request.NewErrParamRequired("InstanceId"))
	}
	if s.VpcAuthMode == nil {
		invalidParams.Add(request.NewErrParamRequired("VpcAuthMode"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetClientToken sets the ClientToken field's value.
func (s *ModifyDBInstanceVpcAuthModeInput) SetClientToken(v string) *ModifyDBInstanceVpcAuthModeInput {
	s.ClientToken = &v
	return s
}

// SetInstanceId sets the InstanceId field's value.
func (s *ModifyDBInstanceVpcAuthModeInput) SetInstanceId(v string) *ModifyDBInstanceVpcAuthModeInput {
	s.InstanceId = &v
	return s
}

// SetVpcAuthMode sets the VpcAuthMode field's value.
func (s *ModifyDBInstanceVpcAuthModeInput) SetVpcAuthMode(v string) *ModifyDBInstanceVpcAuthModeInput {
	s.VpcAuthMode = &v
	return s
}

type ModifyDBInstanceVpcAuthModeOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s ModifyDBInstanceVpcAuthModeOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyDBInstanceVpcAuthModeOutput) GoString() string {
	return s.String()
}

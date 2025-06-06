// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package edx

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opModifyDXPInstanceCommon = "ModifyDXPInstance"

// ModifyDXPInstanceCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ModifyDXPInstanceCommon operation. The "output" return
// value will be populated with the ModifyDXPInstanceCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyDXPInstanceCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyDXPInstanceCommon Send returns without error.
//
// See ModifyDXPInstanceCommon for more information on using the ModifyDXPInstanceCommon
// API call, and error handling.
//
//    // Example sending a request using the ModifyDXPInstanceCommonRequest method.
//    req, resp := client.ModifyDXPInstanceCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *EDX) ModifyDXPInstanceCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opModifyDXPInstanceCommon,
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

// ModifyDXPInstanceCommon API operation for EDX.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for EDX's
// API operation ModifyDXPInstanceCommon for usage and error information.
func (c *EDX) ModifyDXPInstanceCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ModifyDXPInstanceCommonRequest(input)
	return out, req.Send()
}

// ModifyDXPInstanceCommonWithContext is the same as ModifyDXPInstanceCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyDXPInstanceCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *EDX) ModifyDXPInstanceCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ModifyDXPInstanceCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opModifyDXPInstance = "ModifyDXPInstance"

// ModifyDXPInstanceRequest generates a "volcengine/request.Request" representing the
// client's request for the ModifyDXPInstance operation. The "output" return
// value will be populated with the ModifyDXPInstanceCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyDXPInstanceCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyDXPInstanceCommon Send returns without error.
//
// See ModifyDXPInstance for more information on using the ModifyDXPInstance
// API call, and error handling.
//
//    // Example sending a request using the ModifyDXPInstanceRequest method.
//    req, resp := client.ModifyDXPInstanceRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *EDX) ModifyDXPInstanceRequest(input *ModifyDXPInstanceInput) (req *request.Request, output *ModifyDXPInstanceOutput) {
	op := &request.Operation{
		Name:       opModifyDXPInstance,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ModifyDXPInstanceInput{}
	}

	output = &ModifyDXPInstanceOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// ModifyDXPInstance API operation for EDX.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for EDX's
// API operation ModifyDXPInstance for usage and error information.
func (c *EDX) ModifyDXPInstance(input *ModifyDXPInstanceInput) (*ModifyDXPInstanceOutput, error) {
	req, out := c.ModifyDXPInstanceRequest(input)
	return out, req.Send()
}

// ModifyDXPInstanceWithContext is the same as ModifyDXPInstance with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyDXPInstance for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *EDX) ModifyDXPInstanceWithContext(ctx volcengine.Context, input *ModifyDXPInstanceInput, opts ...request.Option) (*ModifyDXPInstanceOutput, error) {
	req, out := c.ModifyDXPInstanceRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ModifyDXPInstanceInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Area *string `type:"string" json:",omitempty"`

	AutoRenew *string `type:"string" json:",omitempty"`

	// InstanceId is a required field
	InstanceId *string `type:"string" json:",omitempty" required:"true"`

	InstanceName *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s ModifyDXPInstanceInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyDXPInstanceInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ModifyDXPInstanceInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ModifyDXPInstanceInput"}
	if s.InstanceId == nil {
		invalidParams.Add(request.NewErrParamRequired("InstanceId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetArea sets the Area field's value.
func (s *ModifyDXPInstanceInput) SetArea(v string) *ModifyDXPInstanceInput {
	s.Area = &v
	return s
}

// SetAutoRenew sets the AutoRenew field's value.
func (s *ModifyDXPInstanceInput) SetAutoRenew(v string) *ModifyDXPInstanceInput {
	s.AutoRenew = &v
	return s
}

// SetInstanceId sets the InstanceId field's value.
func (s *ModifyDXPInstanceInput) SetInstanceId(v string) *ModifyDXPInstanceInput {
	s.InstanceId = &v
	return s
}

// SetInstanceName sets the InstanceName field's value.
func (s *ModifyDXPInstanceInput) SetInstanceName(v string) *ModifyDXPInstanceInput {
	s.InstanceName = &v
	return s
}

type ModifyDXPInstanceOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s ModifyDXPInstanceOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyDXPInstanceOutput) GoString() string {
	return s.String()
}

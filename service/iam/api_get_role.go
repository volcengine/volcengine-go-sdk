// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package iam

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opGetRoleCommon = "GetRole"

// GetRoleCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the GetRoleCommon operation. The "output" return
// value will be populated with the GetRoleCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned GetRoleCommon Request to send the API call to the service.
// the "output" return value is not valid until after GetRoleCommon Send returns without error.
//
// See GetRoleCommon for more information on using the GetRoleCommon
// API call, and error handling.
//
//    // Example sending a request using the GetRoleCommonRequest method.
//    req, resp := client.GetRoleCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *IAM) GetRoleCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opGetRoleCommon,
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

// GetRoleCommon API operation for IAM.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for IAM's
// API operation GetRoleCommon for usage and error information.
func (c *IAM) GetRoleCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.GetRoleCommonRequest(input)
	return out, req.Send()
}

// GetRoleCommonWithContext is the same as GetRoleCommon with the addition of
// the ability to pass a context and additional request options.
//
// See GetRoleCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *IAM) GetRoleCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.GetRoleCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opGetRole = "GetRole"

// GetRoleRequest generates a "volcengine/request.Request" representing the
// client's request for the GetRole operation. The "output" return
// value will be populated with the GetRoleCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned GetRoleCommon Request to send the API call to the service.
// the "output" return value is not valid until after GetRoleCommon Send returns without error.
//
// See GetRole for more information on using the GetRole
// API call, and error handling.
//
//    // Example sending a request using the GetRoleRequest method.
//    req, resp := client.GetRoleRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *IAM) GetRoleRequest(input *GetRoleInput) (req *request.Request, output *GetRoleOutput) {
	op := &request.Operation{
		Name:       opGetRole,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetRoleInput{}
	}

	output = &GetRoleOutput{}
	req = c.newRequest(op, input, output)

	return
}

// GetRole API operation for IAM.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for IAM's
// API operation GetRole for usage and error information.
func (c *IAM) GetRole(input *GetRoleInput) (*GetRoleOutput, error) {
	req, out := c.GetRoleRequest(input)
	return out, req.Send()
}

// GetRoleWithContext is the same as GetRole with the addition of
// the ability to pass a context and additional request options.
//
// See GetRole for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *IAM) GetRoleWithContext(ctx volcengine.Context, input *GetRoleInput, opts ...request.Option) (*GetRoleOutput, error) {
	req, out := c.GetRoleRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type GetRoleInput struct {
	_ struct{} `type:"structure"`

	// RoleName is a required field
	RoleName *string `type:"string" required:"true"`
}

// String returns the string representation
func (s GetRoleInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s GetRoleInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetRoleInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "GetRoleInput"}
	if s.RoleName == nil {
		invalidParams.Add(request.NewErrParamRequired("RoleName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetRoleName sets the RoleName field's value.
func (s *GetRoleInput) SetRoleName(v string) *GetRoleInput {
	s.RoleName = &v
	return s
}

type GetRoleOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	ResponseMetadata *interface{} `type:"interface"`

	Result *interface{} `type:"interface"`
}

// String returns the string representation
func (s GetRoleOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s GetRoleOutput) GoString() string {
	return s.String()
}

// SetResponseMetadata sets the ResponseMetadata field's value.
func (s *GetRoleOutput) SetResponseMetadata(v interface{}) *GetRoleOutput {
	s.ResponseMetadata = &v
	return s
}

// SetResult sets the Result field's value.
func (s *GetRoleOutput) SetResult(v interface{}) *GetRoleOutput {
	s.Result = &v
	return s
}

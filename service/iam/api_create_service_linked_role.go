// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package iam

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opCreateServiceLinkedRoleCommon = "CreateServiceLinkedRole"

// CreateServiceLinkedRoleCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateServiceLinkedRoleCommon operation. The "output" return
// value will be populated with the CreateServiceLinkedRoleCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateServiceLinkedRoleCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateServiceLinkedRoleCommon Send returns without error.
//
// See CreateServiceLinkedRoleCommon for more information on using the CreateServiceLinkedRoleCommon
// API call, and error handling.
//
//    // Example sending a request using the CreateServiceLinkedRoleCommonRequest method.
//    req, resp := client.CreateServiceLinkedRoleCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *IAM) CreateServiceLinkedRoleCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opCreateServiceLinkedRoleCommon,
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

// CreateServiceLinkedRoleCommon API operation for IAM.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for IAM's
// API operation CreateServiceLinkedRoleCommon for usage and error information.
func (c *IAM) CreateServiceLinkedRoleCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.CreateServiceLinkedRoleCommonRequest(input)
	return out, req.Send()
}

// CreateServiceLinkedRoleCommonWithContext is the same as CreateServiceLinkedRoleCommon with the addition of
// the ability to pass a context and additional request options.
//
// See CreateServiceLinkedRoleCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *IAM) CreateServiceLinkedRoleCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.CreateServiceLinkedRoleCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opCreateServiceLinkedRole = "CreateServiceLinkedRole"

// CreateServiceLinkedRoleRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateServiceLinkedRole operation. The "output" return
// value will be populated with the CreateServiceLinkedRoleCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateServiceLinkedRoleCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateServiceLinkedRoleCommon Send returns without error.
//
// See CreateServiceLinkedRole for more information on using the CreateServiceLinkedRole
// API call, and error handling.
//
//    // Example sending a request using the CreateServiceLinkedRoleRequest method.
//    req, resp := client.CreateServiceLinkedRoleRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *IAM) CreateServiceLinkedRoleRequest(input *CreateServiceLinkedRoleInput) (req *request.Request, output *CreateServiceLinkedRoleOutput) {
	op := &request.Operation{
		Name:       opCreateServiceLinkedRole,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateServiceLinkedRoleInput{}
	}

	output = &CreateServiceLinkedRoleOutput{}
	req = c.newRequest(op, input, output)

	return
}

// CreateServiceLinkedRole API operation for IAM.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for IAM's
// API operation CreateServiceLinkedRole for usage and error information.
func (c *IAM) CreateServiceLinkedRole(input *CreateServiceLinkedRoleInput) (*CreateServiceLinkedRoleOutput, error) {
	req, out := c.CreateServiceLinkedRoleRequest(input)
	return out, req.Send()
}

// CreateServiceLinkedRoleWithContext is the same as CreateServiceLinkedRole with the addition of
// the ability to pass a context and additional request options.
//
// See CreateServiceLinkedRole for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *IAM) CreateServiceLinkedRoleWithContext(ctx volcengine.Context, input *CreateServiceLinkedRoleInput, opts ...request.Option) (*CreateServiceLinkedRoleOutput, error) {
	req, out := c.CreateServiceLinkedRoleRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type CreateServiceLinkedRoleInput struct {
	_ struct{} `type:"structure"`

	// ServiceName is a required field
	ServiceName *string `type:"string" required:"true"`
}

// String returns the string representation
func (s CreateServiceLinkedRoleInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateServiceLinkedRoleInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateServiceLinkedRoleInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "CreateServiceLinkedRoleInput"}
	if s.ServiceName == nil {
		invalidParams.Add(request.NewErrParamRequired("ServiceName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetServiceName sets the ServiceName field's value.
func (s *CreateServiceLinkedRoleInput) SetServiceName(v string) *CreateServiceLinkedRoleInput {
	s.ServiceName = &v
	return s
}

type CreateServiceLinkedRoleOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	ResponseMetadata *interface{} `type:"interface"`
}

// String returns the string representation
func (s CreateServiceLinkedRoleOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateServiceLinkedRoleOutput) GoString() string {
	return s.String()
}

// SetResponseMetadata sets the ResponseMetadata field's value.
func (s *CreateServiceLinkedRoleOutput) SetResponseMetadata(v interface{}) *CreateServiceLinkedRoleOutput {
	s.ResponseMetadata = &v
	return s
}

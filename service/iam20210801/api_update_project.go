// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package iam20210801

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opUpdateProjectCommon = "UpdateProject"

// UpdateProjectCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the UpdateProjectCommon operation. The "output" return
// value will be populated with the UpdateProjectCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned UpdateProjectCommon Request to send the API call to the service.
// the "output" return value is not valid until after UpdateProjectCommon Send returns without error.
//
// See UpdateProjectCommon for more information on using the UpdateProjectCommon
// API call, and error handling.
//
//    // Example sending a request using the UpdateProjectCommonRequest method.
//    req, resp := client.UpdateProjectCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *IAM20210801) UpdateProjectCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opUpdateProjectCommon,
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

// UpdateProjectCommon API operation for IAM20210801.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for IAM20210801's
// API operation UpdateProjectCommon for usage and error information.
func (c *IAM20210801) UpdateProjectCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.UpdateProjectCommonRequest(input)
	return out, req.Send()
}

// UpdateProjectCommonWithContext is the same as UpdateProjectCommon with the addition of
// the ability to pass a context and additional request options.
//
// See UpdateProjectCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *IAM20210801) UpdateProjectCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.UpdateProjectCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opUpdateProject = "UpdateProject"

// UpdateProjectRequest generates a "volcengine/request.Request" representing the
// client's request for the UpdateProject operation. The "output" return
// value will be populated with the UpdateProjectCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned UpdateProjectCommon Request to send the API call to the service.
// the "output" return value is not valid until after UpdateProjectCommon Send returns without error.
//
// See UpdateProject for more information on using the UpdateProject
// API call, and error handling.
//
//    // Example sending a request using the UpdateProjectRequest method.
//    req, resp := client.UpdateProjectRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *IAM20210801) UpdateProjectRequest(input *UpdateProjectInput) (req *request.Request, output *UpdateProjectOutput) {
	op := &request.Operation{
		Name:       opUpdateProject,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateProjectInput{}
	}

	output = &UpdateProjectOutput{}
	req = c.newRequest(op, input, output)

	return
}

// UpdateProject API operation for IAM20210801.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for IAM20210801's
// API operation UpdateProject for usage and error information.
func (c *IAM20210801) UpdateProject(input *UpdateProjectInput) (*UpdateProjectOutput, error) {
	req, out := c.UpdateProjectRequest(input)
	return out, req.Send()
}

// UpdateProjectWithContext is the same as UpdateProject with the addition of
// the ability to pass a context and additional request options.
//
// See UpdateProject for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *IAM20210801) UpdateProjectWithContext(ctx volcengine.Context, input *UpdateProjectInput, opts ...request.Option) (*UpdateProjectOutput, error) {
	req, out := c.UpdateProjectRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type UpdateProjectInput struct {
	_ struct{} `type:"structure"`

	Description *string `type:"string"`

	DisplayName *string `type:"string"`

	// ProjectName is a required field
	ProjectName *string `type:"string" required:"true"`
}

// String returns the string representation
func (s UpdateProjectInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s UpdateProjectInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateProjectInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "UpdateProjectInput"}
	if s.ProjectName == nil {
		invalidParams.Add(request.NewErrParamRequired("ProjectName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetDescription sets the Description field's value.
func (s *UpdateProjectInput) SetDescription(v string) *UpdateProjectInput {
	s.Description = &v
	return s
}

// SetDisplayName sets the DisplayName field's value.
func (s *UpdateProjectInput) SetDisplayName(v string) *UpdateProjectInput {
	s.DisplayName = &v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *UpdateProjectInput) SetProjectName(v string) *UpdateProjectInput {
	s.ProjectName = &v
	return s
}

type UpdateProjectOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s UpdateProjectOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s UpdateProjectOutput) GoString() string {
	return s.String()
}
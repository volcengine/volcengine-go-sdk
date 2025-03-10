// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package iam20210801

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opGetProjectCommon = "GetProject"

// GetProjectCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the GetProjectCommon operation. The "output" return
// value will be populated with the GetProjectCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned GetProjectCommon Request to send the API call to the service.
// the "output" return value is not valid until after GetProjectCommon Send returns without error.
//
// See GetProjectCommon for more information on using the GetProjectCommon
// API call, and error handling.
//
//    // Example sending a request using the GetProjectCommonRequest method.
//    req, resp := client.GetProjectCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *IAM20210801) GetProjectCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opGetProjectCommon,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &map[string]interface{}{}
	}

	output = &map[string]interface{}{}
	req = c.newRequest(op, input, output)

	return
}

// GetProjectCommon API operation for IAM20210801.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for IAM20210801's
// API operation GetProjectCommon for usage and error information.
func (c *IAM20210801) GetProjectCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.GetProjectCommonRequest(input)
	return out, req.Send()
}

// GetProjectCommonWithContext is the same as GetProjectCommon with the addition of
// the ability to pass a context and additional request options.
//
// See GetProjectCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *IAM20210801) GetProjectCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.GetProjectCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opGetProject = "GetProject"

// GetProjectRequest generates a "volcengine/request.Request" representing the
// client's request for the GetProject operation. The "output" return
// value will be populated with the GetProjectCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned GetProjectCommon Request to send the API call to the service.
// the "output" return value is not valid until after GetProjectCommon Send returns without error.
//
// See GetProject for more information on using the GetProject
// API call, and error handling.
//
//    // Example sending a request using the GetProjectRequest method.
//    req, resp := client.GetProjectRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *IAM20210801) GetProjectRequest(input *GetProjectInput) (req *request.Request, output *GetProjectOutput) {
	op := &request.Operation{
		Name:       opGetProject,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetProjectInput{}
	}

	output = &GetProjectOutput{}
	req = c.newRequest(op, input, output)

	return
}

// GetProject API operation for IAM20210801.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for IAM20210801's
// API operation GetProject for usage and error information.
func (c *IAM20210801) GetProject(input *GetProjectInput) (*GetProjectOutput, error) {
	req, out := c.GetProjectRequest(input)
	return out, req.Send()
}

// GetProjectWithContext is the same as GetProject with the addition of
// the ability to pass a context and additional request options.
//
// See GetProject for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *IAM20210801) GetProjectWithContext(ctx volcengine.Context, input *GetProjectInput, opts ...request.Option) (*GetProjectOutput, error) {
	req, out := c.GetProjectRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type GetProjectInput struct {
	_ struct{} `type:"structure"`

	// ProjectName is a required field
	ProjectName *string `type:"string" required:"true"`
}

// String returns the string representation
func (s GetProjectInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s GetProjectInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetProjectInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "GetProjectInput"}
	if s.ProjectName == nil {
		invalidParams.Add(request.NewErrParamRequired("ProjectName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetProjectName sets the ProjectName field's value.
func (s *GetProjectInput) SetProjectName(v string) *GetProjectInput {
	s.ProjectName = &v
	return s
}

type GetProjectOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	AccountID *int32 `type:"int32"`

	CreateDate *string `type:"string"`

	Description *string `type:"string"`

	DisplayName *string `type:"string"`

	ParentProjectName *string `type:"string"`

	Path *string `type:"string"`

	ProjectName *string `type:"string"`

	Status *string `type:"string"`

	UpdateDate *string `type:"string"`
}

// String returns the string representation
func (s GetProjectOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s GetProjectOutput) GoString() string {
	return s.String()
}

// SetAccountID sets the AccountID field's value.
func (s *GetProjectOutput) SetAccountID(v int32) *GetProjectOutput {
	s.AccountID = &v
	return s
}

// SetCreateDate sets the CreateDate field's value.
func (s *GetProjectOutput) SetCreateDate(v string) *GetProjectOutput {
	s.CreateDate = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *GetProjectOutput) SetDescription(v string) *GetProjectOutput {
	s.Description = &v
	return s
}

// SetDisplayName sets the DisplayName field's value.
func (s *GetProjectOutput) SetDisplayName(v string) *GetProjectOutput {
	s.DisplayName = &v
	return s
}

// SetParentProjectName sets the ParentProjectName field's value.
func (s *GetProjectOutput) SetParentProjectName(v string) *GetProjectOutput {
	s.ParentProjectName = &v
	return s
}

// SetPath sets the Path field's value.
func (s *GetProjectOutput) SetPath(v string) *GetProjectOutput {
	s.Path = &v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *GetProjectOutput) SetProjectName(v string) *GetProjectOutput {
	s.ProjectName = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *GetProjectOutput) SetStatus(v string) *GetProjectOutput {
	s.Status = &v
	return s
}

// SetUpdateDate sets the UpdateDate field's value.
func (s *GetProjectOutput) SetUpdateDate(v string) *GetProjectOutput {
	s.UpdateDate = &v
	return s
}

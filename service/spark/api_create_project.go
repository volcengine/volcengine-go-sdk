// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package spark

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opCreateProjectCommon = "createProject"

// CreateProjectCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateProjectCommon operation. The "output" return
// value will be populated with the CreateProjectCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateProjectCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateProjectCommon Send returns without error.
//
// See CreateProjectCommon for more information on using the CreateProjectCommon
// API call, and error handling.
//
//    // Example sending a request using the CreateProjectCommonRequest method.
//    req, resp := client.CreateProjectCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *SPARK) CreateProjectCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opCreateProjectCommon,
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

// CreateProjectCommon API operation for SPARK.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for SPARK's
// API operation CreateProjectCommon for usage and error information.
func (c *SPARK) CreateProjectCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.CreateProjectCommonRequest(input)
	return out, req.Send()
}

// CreateProjectCommonWithContext is the same as CreateProjectCommon with the addition of
// the ability to pass a context and additional request options.
//
// See CreateProjectCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *SPARK) CreateProjectCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.CreateProjectCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opCreateProject = "createProject"

// CreateProjectRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateProject operation. The "output" return
// value will be populated with the CreateProjectCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateProjectCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateProjectCommon Send returns without error.
//
// See CreateProject for more information on using the CreateProject
// API call, and error handling.
//
//    // Example sending a request using the CreateProjectRequest method.
//    req, resp := client.CreateProjectRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *SPARK) CreateProjectRequest(input *CreateProjectInput) (req *request.Request, output *CreateProjectOutput) {
	op := &request.Operation{
		Name:       opCreateProject,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateProjectInput{}
	}

	output = &CreateProjectOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// CreateProject API operation for SPARK.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for SPARK's
// API operation CreateProject for usage and error information.
func (c *SPARK) CreateProject(input *CreateProjectInput) (*CreateProjectOutput, error) {
	req, out := c.CreateProjectRequest(input)
	return out, req.Send()
}

// CreateProjectWithContext is the same as CreateProject with the addition of
// the ability to pass a context and additional request options.
//
// See CreateProject for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *SPARK) CreateProjectWithContext(ctx volcengine.Context, input *CreateProjectInput, opts ...request.Option) (*CreateProjectOutput, error) {
	req, out := c.CreateProjectRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type CreateProjectInput struct {
	_ struct{} `type:"structure"`

	AuthorityType *string `type:"string"`

	Description *string `type:"string"`

	OwnerAccountId *string `type:"string"`

	OwnerId *string `type:"string"`

	ProjectId *string `type:"string"`

	ProjectName *string `type:"string"`

	ProjectType *string `type:"string"`
}

// String returns the string representation
func (s CreateProjectInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateProjectInput) GoString() string {
	return s.String()
}

// SetAuthorityType sets the AuthorityType field's value.
func (s *CreateProjectInput) SetAuthorityType(v string) *CreateProjectInput {
	s.AuthorityType = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *CreateProjectInput) SetDescription(v string) *CreateProjectInput {
	s.Description = &v
	return s
}

// SetOwnerAccountId sets the OwnerAccountId field's value.
func (s *CreateProjectInput) SetOwnerAccountId(v string) *CreateProjectInput {
	s.OwnerAccountId = &v
	return s
}

// SetOwnerId sets the OwnerId field's value.
func (s *CreateProjectInput) SetOwnerId(v string) *CreateProjectInput {
	s.OwnerId = &v
	return s
}

// SetProjectId sets the ProjectId field's value.
func (s *CreateProjectInput) SetProjectId(v string) *CreateProjectInput {
	s.ProjectId = &v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *CreateProjectInput) SetProjectName(v string) *CreateProjectInput {
	s.ProjectName = &v
	return s
}

// SetProjectType sets the ProjectType field's value.
func (s *CreateProjectInput) SetProjectType(v string) *CreateProjectInput {
	s.ProjectType = &v
	return s
}

type CreateProjectOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	AuthorityType *string `type:"string"`

	CreateDate *string `type:"string"`

	DatabaseName *string `type:"string"`

	Description *string `type:"string"`

	OwnerId *string `type:"string"`

	OwnerName *string `type:"string"`

	ProjectId *string `type:"string"`

	ProjectName *string `type:"string"`

	Region *string `type:"string"`
}

// String returns the string representation
func (s CreateProjectOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateProjectOutput) GoString() string {
	return s.String()
}

// SetAuthorityType sets the AuthorityType field's value.
func (s *CreateProjectOutput) SetAuthorityType(v string) *CreateProjectOutput {
	s.AuthorityType = &v
	return s
}

// SetCreateDate sets the CreateDate field's value.
func (s *CreateProjectOutput) SetCreateDate(v string) *CreateProjectOutput {
	s.CreateDate = &v
	return s
}

// SetDatabaseName sets the DatabaseName field's value.
func (s *CreateProjectOutput) SetDatabaseName(v string) *CreateProjectOutput {
	s.DatabaseName = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *CreateProjectOutput) SetDescription(v string) *CreateProjectOutput {
	s.Description = &v
	return s
}

// SetOwnerId sets the OwnerId field's value.
func (s *CreateProjectOutput) SetOwnerId(v string) *CreateProjectOutput {
	s.OwnerId = &v
	return s
}

// SetOwnerName sets the OwnerName field's value.
func (s *CreateProjectOutput) SetOwnerName(v string) *CreateProjectOutput {
	s.OwnerName = &v
	return s
}

// SetProjectId sets the ProjectId field's value.
func (s *CreateProjectOutput) SetProjectId(v string) *CreateProjectOutput {
	s.ProjectId = &v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *CreateProjectOutput) SetProjectName(v string) *CreateProjectOutput {
	s.ProjectName = &v
	return s
}

// SetRegion sets the Region field's value.
func (s *CreateProjectOutput) SetRegion(v string) *CreateProjectOutput {
	s.Region = &v
	return s
}

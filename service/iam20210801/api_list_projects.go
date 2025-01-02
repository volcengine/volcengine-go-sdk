// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package iam20210801

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opListProjectsCommon = "ListProjects"

// ListProjectsCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ListProjectsCommon operation. The "output" return
// value will be populated with the ListProjectsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListProjectsCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListProjectsCommon Send returns without error.
//
// See ListProjectsCommon for more information on using the ListProjectsCommon
// API call, and error handling.
//
//    // Example sending a request using the ListProjectsCommonRequest method.
//    req, resp := client.ListProjectsCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *IAM20210801) ListProjectsCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opListProjectsCommon,
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

// ListProjectsCommon API operation for IAM20210801.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for IAM20210801's
// API operation ListProjectsCommon for usage and error information.
func (c *IAM20210801) ListProjectsCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ListProjectsCommonRequest(input)
	return out, req.Send()
}

// ListProjectsCommonWithContext is the same as ListProjectsCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ListProjectsCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *IAM20210801) ListProjectsCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ListProjectsCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opListProjects = "ListProjects"

// ListProjectsRequest generates a "volcengine/request.Request" representing the
// client's request for the ListProjects operation. The "output" return
// value will be populated with the ListProjectsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListProjectsCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListProjectsCommon Send returns without error.
//
// See ListProjects for more information on using the ListProjects
// API call, and error handling.
//
//    // Example sending a request using the ListProjectsRequest method.
//    req, resp := client.ListProjectsRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *IAM20210801) ListProjectsRequest(input *ListProjectsInput) (req *request.Request, output *ListProjectsOutput) {
	op := &request.Operation{
		Name:       opListProjects,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListProjectsInput{}
	}

	output = &ListProjectsOutput{}
	req = c.newRequest(op, input, output)

	return
}

// ListProjects API operation for IAM20210801.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for IAM20210801's
// API operation ListProjects for usage and error information.
func (c *IAM20210801) ListProjects(input *ListProjectsInput) (*ListProjectsOutput, error) {
	req, out := c.ListProjectsRequest(input)
	return out, req.Send()
}

// ListProjectsWithContext is the same as ListProjects with the addition of
// the ability to pass a context and additional request options.
//
// See ListProjects for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *IAM20210801) ListProjectsWithContext(ctx volcengine.Context, input *ListProjectsInput, opts ...request.Option) (*ListProjectsOutput, error) {
	req, out := c.ListProjectsRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ListProjectsInput struct {
	_ struct{} `type:"structure"`

	Limit *int32 `type:"int32"`

	Offset *int32 `type:"int32"`

	Query *string `type:"string"`

	WithParentProject *int32 `type:"int32"`
}

// String returns the string representation
func (s ListProjectsInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListProjectsInput) GoString() string {
	return s.String()
}

// SetLimit sets the Limit field's value.
func (s *ListProjectsInput) SetLimit(v int32) *ListProjectsInput {
	s.Limit = &v
	return s
}

// SetOffset sets the Offset field's value.
func (s *ListProjectsInput) SetOffset(v int32) *ListProjectsInput {
	s.Offset = &v
	return s
}

// SetQuery sets the Query field's value.
func (s *ListProjectsInput) SetQuery(v string) *ListProjectsInput {
	s.Query = &v
	return s
}

// SetWithParentProject sets the WithParentProject field's value.
func (s *ListProjectsInput) SetWithParentProject(v int32) *ListProjectsInput {
	s.WithParentProject = &v
	return s
}

type ListProjectsOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	Limit *int32 `type:"int32"`

	Offset *int32 `type:"int32"`

	Projects []*ProjectForListProjectsOutput `type:"list"`

	Total *int32 `type:"int32"`
}

// String returns the string representation
func (s ListProjectsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListProjectsOutput) GoString() string {
	return s.String()
}

// SetLimit sets the Limit field's value.
func (s *ListProjectsOutput) SetLimit(v int32) *ListProjectsOutput {
	s.Limit = &v
	return s
}

// SetOffset sets the Offset field's value.
func (s *ListProjectsOutput) SetOffset(v int32) *ListProjectsOutput {
	s.Offset = &v
	return s
}

// SetProjects sets the Projects field's value.
func (s *ListProjectsOutput) SetProjects(v []*ProjectForListProjectsOutput) *ListProjectsOutput {
	s.Projects = v
	return s
}

// SetTotal sets the Total field's value.
func (s *ListProjectsOutput) SetTotal(v int32) *ListProjectsOutput {
	s.Total = &v
	return s
}

type ProjectForListProjectsOutput struct {
	_ struct{} `type:"structure"`

	AccountID *int64 `type:"int64"`

	CreateDate *string `type:"string"`

	Description *string `type:"string"`

	DisplayName *string `type:"string"`

	HasPermission *bool `type:"boolean"`

	ParentProjectName *string `type:"string"`

	Path *string `type:"string"`

	ProjectName *string `type:"string"`

	Status *string `type:"string"`

	UpdateDate *string `type:"string"`
}

// String returns the string representation
func (s ProjectForListProjectsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ProjectForListProjectsOutput) GoString() string {
	return s.String()
}

// SetAccountID sets the AccountID field's value.
func (s *ProjectForListProjectsOutput) SetAccountID(v int64) *ProjectForListProjectsOutput {
	s.AccountID = &v
	return s
}

// SetCreateDate sets the CreateDate field's value.
func (s *ProjectForListProjectsOutput) SetCreateDate(v string) *ProjectForListProjectsOutput {
	s.CreateDate = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *ProjectForListProjectsOutput) SetDescription(v string) *ProjectForListProjectsOutput {
	s.Description = &v
	return s
}

// SetDisplayName sets the DisplayName field's value.
func (s *ProjectForListProjectsOutput) SetDisplayName(v string) *ProjectForListProjectsOutput {
	s.DisplayName = &v
	return s
}

// SetHasPermission sets the HasPermission field's value.
func (s *ProjectForListProjectsOutput) SetHasPermission(v bool) *ProjectForListProjectsOutput {
	s.HasPermission = &v
	return s
}

// SetParentProjectName sets the ParentProjectName field's value.
func (s *ProjectForListProjectsOutput) SetParentProjectName(v string) *ProjectForListProjectsOutput {
	s.ParentProjectName = &v
	return s
}

// SetPath sets the Path field's value.
func (s *ProjectForListProjectsOutput) SetPath(v string) *ProjectForListProjectsOutput {
	s.Path = &v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *ProjectForListProjectsOutput) SetProjectName(v string) *ProjectForListProjectsOutput {
	s.ProjectName = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *ProjectForListProjectsOutput) SetStatus(v string) *ProjectForListProjectsOutput {
	s.Status = &v
	return s
}

// SetUpdateDate sets the UpdateDate field's value.
func (s *ProjectForListProjectsOutput) SetUpdateDate(v string) *ProjectForListProjectsOutput {
	s.UpdateDate = &v
	return s
}

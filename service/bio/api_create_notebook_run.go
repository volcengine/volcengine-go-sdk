// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package bio

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opCreateNotebookRunCommon = "CreateNotebookRun"

// CreateNotebookRunCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateNotebookRunCommon operation. The "output" return
// value will be populated with the CreateNotebookRunCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateNotebookRunCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateNotebookRunCommon Send returns without error.
//
// See CreateNotebookRunCommon for more information on using the CreateNotebookRunCommon
// API call, and error handling.
//
//    // Example sending a request using the CreateNotebookRunCommonRequest method.
//    req, resp := client.CreateNotebookRunCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *BIO) CreateNotebookRunCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opCreateNotebookRunCommon,
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

// CreateNotebookRunCommon API operation for BIO.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for BIO's
// API operation CreateNotebookRunCommon for usage and error information.
func (c *BIO) CreateNotebookRunCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.CreateNotebookRunCommonRequest(input)
	return out, req.Send()
}

// CreateNotebookRunCommonWithContext is the same as CreateNotebookRunCommon with the addition of
// the ability to pass a context and additional request options.
//
// See CreateNotebookRunCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *BIO) CreateNotebookRunCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.CreateNotebookRunCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opCreateNotebookRun = "CreateNotebookRun"

// CreateNotebookRunRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateNotebookRun operation. The "output" return
// value will be populated with the CreateNotebookRunCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateNotebookRunCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateNotebookRunCommon Send returns without error.
//
// See CreateNotebookRun for more information on using the CreateNotebookRun
// API call, and error handling.
//
//    // Example sending a request using the CreateNotebookRunRequest method.
//    req, resp := client.CreateNotebookRunRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *BIO) CreateNotebookRunRequest(input *CreateNotebookRunInput) (req *request.Request, output *CreateNotebookRunOutput) {
	op := &request.Operation{
		Name:       opCreateNotebookRun,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateNotebookRunInput{}
	}

	output = &CreateNotebookRunOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// CreateNotebookRun API operation for BIO.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for BIO's
// API operation CreateNotebookRun for usage and error information.
func (c *BIO) CreateNotebookRun(input *CreateNotebookRunInput) (*CreateNotebookRunOutput, error) {
	req, out := c.CreateNotebookRunRequest(input)
	return out, req.Send()
}

// CreateNotebookRunWithContext is the same as CreateNotebookRun with the addition of
// the ability to pass a context and additional request options.
//
// See CreateNotebookRun for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *BIO) CreateNotebookRunWithContext(ctx volcengine.Context, input *CreateNotebookRunInput, opts ...request.Option) (*CreateNotebookRunOutput, error) {
	req, out := c.CreateNotebookRunRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type CreateNotebookRunInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	// Name is a required field
	Name *string `type:"string" json:",omitempty" required:"true"`

	// WorkspaceID is a required field
	WorkspaceID *string `type:"string" json:",omitempty" required:"true"`
}

// String returns the string representation
func (s CreateNotebookRunInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateNotebookRunInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateNotebookRunInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "CreateNotebookRunInput"}
	if s.Name == nil {
		invalidParams.Add(request.NewErrParamRequired("Name"))
	}
	if s.WorkspaceID == nil {
		invalidParams.Add(request.NewErrParamRequired("WorkspaceID"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetName sets the Name field's value.
func (s *CreateNotebookRunInput) SetName(v string) *CreateNotebookRunInput {
	s.Name = &v
	return s
}

// SetWorkspaceID sets the WorkspaceID field's value.
func (s *CreateNotebookRunInput) SetWorkspaceID(v string) *CreateNotebookRunInput {
	s.WorkspaceID = &v
	return s
}

type CreateNotebookRunOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	JobID *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s CreateNotebookRunOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateNotebookRunOutput) GoString() string {
	return s.String()
}

// SetJobID sets the JobID field's value.
func (s *CreateNotebookRunOutput) SetJobID(v string) *CreateNotebookRunOutput {
	s.JobID = &v
	return s
}

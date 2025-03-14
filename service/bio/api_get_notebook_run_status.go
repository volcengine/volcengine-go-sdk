// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package bio

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opGetNotebookRunStatusCommon = "GetNotebookRunStatus"

// GetNotebookRunStatusCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the GetNotebookRunStatusCommon operation. The "output" return
// value will be populated with the GetNotebookRunStatusCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned GetNotebookRunStatusCommon Request to send the API call to the service.
// the "output" return value is not valid until after GetNotebookRunStatusCommon Send returns without error.
//
// See GetNotebookRunStatusCommon for more information on using the GetNotebookRunStatusCommon
// API call, and error handling.
//
//    // Example sending a request using the GetNotebookRunStatusCommonRequest method.
//    req, resp := client.GetNotebookRunStatusCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *BIO) GetNotebookRunStatusCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opGetNotebookRunStatusCommon,
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

// GetNotebookRunStatusCommon API operation for BIO.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for BIO's
// API operation GetNotebookRunStatusCommon for usage and error information.
func (c *BIO) GetNotebookRunStatusCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.GetNotebookRunStatusCommonRequest(input)
	return out, req.Send()
}

// GetNotebookRunStatusCommonWithContext is the same as GetNotebookRunStatusCommon with the addition of
// the ability to pass a context and additional request options.
//
// See GetNotebookRunStatusCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *BIO) GetNotebookRunStatusCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.GetNotebookRunStatusCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opGetNotebookRunStatus = "GetNotebookRunStatus"

// GetNotebookRunStatusRequest generates a "volcengine/request.Request" representing the
// client's request for the GetNotebookRunStatus operation. The "output" return
// value will be populated with the GetNotebookRunStatusCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned GetNotebookRunStatusCommon Request to send the API call to the service.
// the "output" return value is not valid until after GetNotebookRunStatusCommon Send returns without error.
//
// See GetNotebookRunStatus for more information on using the GetNotebookRunStatus
// API call, and error handling.
//
//    // Example sending a request using the GetNotebookRunStatusRequest method.
//    req, resp := client.GetNotebookRunStatusRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *BIO) GetNotebookRunStatusRequest(input *GetNotebookRunStatusInput) (req *request.Request, output *GetNotebookRunStatusOutput) {
	op := &request.Operation{
		Name:       opGetNotebookRunStatus,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetNotebookRunStatusInput{}
	}

	output = &GetNotebookRunStatusOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// GetNotebookRunStatus API operation for BIO.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for BIO's
// API operation GetNotebookRunStatus for usage and error information.
func (c *BIO) GetNotebookRunStatus(input *GetNotebookRunStatusInput) (*GetNotebookRunStatusOutput, error) {
	req, out := c.GetNotebookRunStatusRequest(input)
	return out, req.Send()
}

// GetNotebookRunStatusWithContext is the same as GetNotebookRunStatus with the addition of
// the ability to pass a context and additional request options.
//
// See GetNotebookRunStatus for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *BIO) GetNotebookRunStatusWithContext(ctx volcengine.Context, input *GetNotebookRunStatusInput, opts ...request.Option) (*GetNotebookRunStatusOutput, error) {
	req, out := c.GetNotebookRunStatusRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type GetNotebookRunStatusInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	// JobID is a required field
	JobID *string `type:"string" json:",omitempty" required:"true"`

	// WorkspaceID is a required field
	WorkspaceID *string `type:"string" json:",omitempty" required:"true"`
}

// String returns the string representation
func (s GetNotebookRunStatusInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s GetNotebookRunStatusInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetNotebookRunStatusInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "GetNotebookRunStatusInput"}
	if s.JobID == nil {
		invalidParams.Add(request.NewErrParamRequired("JobID"))
	}
	if s.WorkspaceID == nil {
		invalidParams.Add(request.NewErrParamRequired("WorkspaceID"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetJobID sets the JobID field's value.
func (s *GetNotebookRunStatusInput) SetJobID(v string) *GetNotebookRunStatusInput {
	s.JobID = &v
	return s
}

// SetWorkspaceID sets the WorkspaceID field's value.
func (s *GetNotebookRunStatusInput) SetWorkspaceID(v string) *GetNotebookRunStatusInput {
	s.WorkspaceID = &v
	return s
}

type GetNotebookRunStatusOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	ExitCode *int32 `type:"int32" json:",omitempty"`

	Log *string `type:"string" json:",omitempty"`

	Status *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s GetNotebookRunStatusOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s GetNotebookRunStatusOutput) GoString() string {
	return s.String()
}

// SetExitCode sets the ExitCode field's value.
func (s *GetNotebookRunStatusOutput) SetExitCode(v int32) *GetNotebookRunStatusOutput {
	s.ExitCode = &v
	return s
}

// SetLog sets the Log field's value.
func (s *GetNotebookRunStatusOutput) SetLog(v string) *GetNotebookRunStatusOutput {
	s.Log = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *GetNotebookRunStatusOutput) SetStatus(v string) *GetNotebookRunStatusOutput {
	s.Status = &v
	return s
}

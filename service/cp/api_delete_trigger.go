// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package cp

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDeleteTriggerCommon = "DeleteTrigger"

// DeleteTriggerCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DeleteTriggerCommon operation. The "output" return
// value will be populated with the DeleteTriggerCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeleteTriggerCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeleteTriggerCommon Send returns without error.
//
// See DeleteTriggerCommon for more information on using the DeleteTriggerCommon
// API call, and error handling.
//
//    // Example sending a request using the DeleteTriggerCommonRequest method.
//    req, resp := client.DeleteTriggerCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CP) DeleteTriggerCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDeleteTriggerCommon,
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

// DeleteTriggerCommon API operation for CP.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CP's
// API operation DeleteTriggerCommon for usage and error information.
func (c *CP) DeleteTriggerCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DeleteTriggerCommonRequest(input)
	return out, req.Send()
}

// DeleteTriggerCommonWithContext is the same as DeleteTriggerCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteTriggerCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CP) DeleteTriggerCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DeleteTriggerCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDeleteTrigger = "DeleteTrigger"

// DeleteTriggerRequest generates a "volcengine/request.Request" representing the
// client's request for the DeleteTrigger operation. The "output" return
// value will be populated with the DeleteTriggerCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeleteTriggerCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeleteTriggerCommon Send returns without error.
//
// See DeleteTrigger for more information on using the DeleteTrigger
// API call, and error handling.
//
//    // Example sending a request using the DeleteTriggerRequest method.
//    req, resp := client.DeleteTriggerRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CP) DeleteTriggerRequest(input *DeleteTriggerInput) (req *request.Request, output *DeleteTriggerOutput) {
	op := &request.Operation{
		Name:       opDeleteTrigger,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteTriggerInput{}
	}

	output = &DeleteTriggerOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// DeleteTrigger API operation for CP.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CP's
// API operation DeleteTrigger for usage and error information.
func (c *CP) DeleteTrigger(input *DeleteTriggerInput) (*DeleteTriggerOutput, error) {
	req, out := c.DeleteTriggerRequest(input)
	return out, req.Send()
}

// DeleteTriggerWithContext is the same as DeleteTrigger with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteTrigger for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CP) DeleteTriggerWithContext(ctx volcengine.Context, input *DeleteTriggerInput, opts ...request.Option) (*DeleteTriggerOutput, error) {
	req, out := c.DeleteTriggerRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DeleteTriggerInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	// Id is a required field
	Id *string `type:"string" json:",omitempty" required:"true"`

	// PipelineId is a required field
	PipelineId *string `type:"string" json:",omitempty" required:"true"`

	// WorkspaceId is a required field
	WorkspaceId *string `type:"string" json:",omitempty" required:"true"`
}

// String returns the string representation
func (s DeleteTriggerInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteTriggerInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteTriggerInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DeleteTriggerInput"}
	if s.Id == nil {
		invalidParams.Add(request.NewErrParamRequired("Id"))
	}
	if s.PipelineId == nil {
		invalidParams.Add(request.NewErrParamRequired("PipelineId"))
	}
	if s.WorkspaceId == nil {
		invalidParams.Add(request.NewErrParamRequired("WorkspaceId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetId sets the Id field's value.
func (s *DeleteTriggerInput) SetId(v string) *DeleteTriggerInput {
	s.Id = &v
	return s
}

// SetPipelineId sets the PipelineId field's value.
func (s *DeleteTriggerInput) SetPipelineId(v string) *DeleteTriggerInput {
	s.PipelineId = &v
	return s
}

// SetWorkspaceId sets the WorkspaceId field's value.
func (s *DeleteTriggerInput) SetWorkspaceId(v string) *DeleteTriggerInput {
	s.WorkspaceId = &v
	return s
}

type DeleteTriggerOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s DeleteTriggerOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteTriggerOutput) GoString() string {
	return s.String()
}
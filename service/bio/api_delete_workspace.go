// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package bio

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDeleteWorkspaceCommon = "DeleteWorkspace"

// DeleteWorkspaceCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DeleteWorkspaceCommon operation. The "output" return
// value will be populated with the DeleteWorkspaceCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeleteWorkspaceCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeleteWorkspaceCommon Send returns without error.
//
// See DeleteWorkspaceCommon for more information on using the DeleteWorkspaceCommon
// API call, and error handling.
//
//    // Example sending a request using the DeleteWorkspaceCommonRequest method.
//    req, resp := client.DeleteWorkspaceCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *BIO) DeleteWorkspaceCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDeleteWorkspaceCommon,
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

// DeleteWorkspaceCommon API operation for BIO.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for BIO's
// API operation DeleteWorkspaceCommon for usage and error information.
func (c *BIO) DeleteWorkspaceCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DeleteWorkspaceCommonRequest(input)
	return out, req.Send()
}

// DeleteWorkspaceCommonWithContext is the same as DeleteWorkspaceCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteWorkspaceCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *BIO) DeleteWorkspaceCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DeleteWorkspaceCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDeleteWorkspace = "DeleteWorkspace"

// DeleteWorkspaceRequest generates a "volcengine/request.Request" representing the
// client's request for the DeleteWorkspace operation. The "output" return
// value will be populated with the DeleteWorkspaceCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeleteWorkspaceCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeleteWorkspaceCommon Send returns without error.
//
// See DeleteWorkspace for more information on using the DeleteWorkspace
// API call, and error handling.
//
//    // Example sending a request using the DeleteWorkspaceRequest method.
//    req, resp := client.DeleteWorkspaceRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *BIO) DeleteWorkspaceRequest(input *DeleteWorkspaceInput) (req *request.Request, output *DeleteWorkspaceOutput) {
	op := &request.Operation{
		Name:       opDeleteWorkspace,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteWorkspaceInput{}
	}

	output = &DeleteWorkspaceOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// DeleteWorkspace API operation for BIO.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for BIO's
// API operation DeleteWorkspace for usage and error information.
func (c *BIO) DeleteWorkspace(input *DeleteWorkspaceInput) (*DeleteWorkspaceOutput, error) {
	req, out := c.DeleteWorkspaceRequest(input)
	return out, req.Send()
}

// DeleteWorkspaceWithContext is the same as DeleteWorkspace with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteWorkspace for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *BIO) DeleteWorkspaceWithContext(ctx volcengine.Context, input *DeleteWorkspaceInput, opts ...request.Option) (*DeleteWorkspaceOutput, error) {
	req, out := c.DeleteWorkspaceRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DeleteWorkspaceInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	// ID is a required field
	ID *string `type:"string" json:",omitempty" required:"true"`
}

// String returns the string representation
func (s DeleteWorkspaceInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteWorkspaceInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteWorkspaceInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DeleteWorkspaceInput"}
	if s.ID == nil {
		invalidParams.Add(request.NewErrParamRequired("ID"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetID sets the ID field's value.
func (s *DeleteWorkspaceInput) SetID(v string) *DeleteWorkspaceInput {
	s.ID = &v
	return s
}

type DeleteWorkspaceOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s DeleteWorkspaceOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteWorkspaceOutput) GoString() string {
	return s.String()
}

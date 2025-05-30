// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package filenas

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDeleteFileSystemCommon = "DeleteFileSystem"

// DeleteFileSystemCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DeleteFileSystemCommon operation. The "output" return
// value will be populated with the DeleteFileSystemCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeleteFileSystemCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeleteFileSystemCommon Send returns without error.
//
// See DeleteFileSystemCommon for more information on using the DeleteFileSystemCommon
// API call, and error handling.
//
//    // Example sending a request using the DeleteFileSystemCommonRequest method.
//    req, resp := client.DeleteFileSystemCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *FILENAS) DeleteFileSystemCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDeleteFileSystemCommon,
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

// DeleteFileSystemCommon API operation for FILENAS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for FILENAS's
// API operation DeleteFileSystemCommon for usage and error information.
func (c *FILENAS) DeleteFileSystemCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DeleteFileSystemCommonRequest(input)
	return out, req.Send()
}

// DeleteFileSystemCommonWithContext is the same as DeleteFileSystemCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteFileSystemCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *FILENAS) DeleteFileSystemCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DeleteFileSystemCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDeleteFileSystem = "DeleteFileSystem"

// DeleteFileSystemRequest generates a "volcengine/request.Request" representing the
// client's request for the DeleteFileSystem operation. The "output" return
// value will be populated with the DeleteFileSystemCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeleteFileSystemCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeleteFileSystemCommon Send returns without error.
//
// See DeleteFileSystem for more information on using the DeleteFileSystem
// API call, and error handling.
//
//    // Example sending a request using the DeleteFileSystemRequest method.
//    req, resp := client.DeleteFileSystemRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *FILENAS) DeleteFileSystemRequest(input *DeleteFileSystemInput) (req *request.Request, output *DeleteFileSystemOutput) {
	op := &request.Operation{
		Name:       opDeleteFileSystem,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteFileSystemInput{}
	}

	output = &DeleteFileSystemOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// DeleteFileSystem API operation for FILENAS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for FILENAS's
// API operation DeleteFileSystem for usage and error information.
func (c *FILENAS) DeleteFileSystem(input *DeleteFileSystemInput) (*DeleteFileSystemOutput, error) {
	req, out := c.DeleteFileSystemRequest(input)
	return out, req.Send()
}

// DeleteFileSystemWithContext is the same as DeleteFileSystem with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteFileSystem for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *FILENAS) DeleteFileSystemWithContext(ctx volcengine.Context, input *DeleteFileSystemInput, opts ...request.Option) (*DeleteFileSystemOutput, error) {
	req, out := c.DeleteFileSystemRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DeleteFileSystemInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	// FileSystemId is a required field
	FileSystemId *string `type:"string" json:",omitempty" required:"true"`
}

// String returns the string representation
func (s DeleteFileSystemInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteFileSystemInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteFileSystemInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DeleteFileSystemInput"}
	if s.FileSystemId == nil {
		invalidParams.Add(request.NewErrParamRequired("FileSystemId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetFileSystemId sets the FileSystemId field's value.
func (s *DeleteFileSystemInput) SetFileSystemId(v string) *DeleteFileSystemInput {
	s.FileSystemId = &v
	return s
}

type DeleteFileSystemOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	OrderNo *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s DeleteFileSystemOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteFileSystemOutput) GoString() string {
	return s.String()
}

// SetOrderNo sets the OrderNo field's value.
func (s *DeleteFileSystemOutput) SetOrderNo(v string) *DeleteFileSystemOutput {
	s.OrderNo = &v
	return s
}

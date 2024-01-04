// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package cen

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDeleteCenCommon = "DeleteCen"

// DeleteCenCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DeleteCenCommon operation. The "output" return
// value will be populated with the DeleteCenCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeleteCenCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeleteCenCommon Send returns without error.
//
// See DeleteCenCommon for more information on using the DeleteCenCommon
// API call, and error handling.
//
//    // Example sending a request using the DeleteCenCommonRequest method.
//    req, resp := client.DeleteCenCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CEN) DeleteCenCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDeleteCenCommon,
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

// DeleteCenCommon API operation for CEN.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CEN's
// API operation DeleteCenCommon for usage and error information.
func (c *CEN) DeleteCenCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DeleteCenCommonRequest(input)
	return out, req.Send()
}

// DeleteCenCommonWithContext is the same as DeleteCenCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteCenCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CEN) DeleteCenCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DeleteCenCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDeleteCen = "DeleteCen"

// DeleteCenRequest generates a "volcengine/request.Request" representing the
// client's request for the DeleteCen operation. The "output" return
// value will be populated with the DeleteCenCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeleteCenCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeleteCenCommon Send returns without error.
//
// See DeleteCen for more information on using the DeleteCen
// API call, and error handling.
//
//    // Example sending a request using the DeleteCenRequest method.
//    req, resp := client.DeleteCenRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CEN) DeleteCenRequest(input *DeleteCenInput) (req *request.Request, output *DeleteCenOutput) {
	op := &request.Operation{
		Name:       opDeleteCen,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteCenInput{}
	}

	output = &DeleteCenOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DeleteCen API operation for CEN.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CEN's
// API operation DeleteCen for usage and error information.
func (c *CEN) DeleteCen(input *DeleteCenInput) (*DeleteCenOutput, error) {
	req, out := c.DeleteCenRequest(input)
	return out, req.Send()
}

// DeleteCenWithContext is the same as DeleteCen with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteCen for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CEN) DeleteCenWithContext(ctx volcengine.Context, input *DeleteCenInput, opts ...request.Option) (*DeleteCenOutput, error) {
	req, out := c.DeleteCenRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DeleteCenInput struct {
	_ struct{} `type:"structure"`

	// CenId is a required field
	CenId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteCenInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteCenInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteCenInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DeleteCenInput"}
	if s.CenId == nil {
		invalidParams.Add(request.NewErrParamRequired("CenId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetCenId sets the CenId field's value.
func (s *DeleteCenInput) SetCenId(v string) *DeleteCenInput {
	s.CenId = &v
	return s
}

type DeleteCenOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	RequestId *string `type:"string"`
}

// String returns the string representation
func (s DeleteCenOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteCenOutput) GoString() string {
	return s.String()
}

// SetRequestId sets the RequestId field's value.
func (s *DeleteCenOutput) SetRequestId(v string) *DeleteCenOutput {
	s.RequestId = &v
	return s
}

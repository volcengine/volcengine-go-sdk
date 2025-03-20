// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package rdspostgresql

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDeleteAllowListCommon = "DeleteAllowList"

// DeleteAllowListCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DeleteAllowListCommon operation. The "output" return
// value will be populated with the DeleteAllowListCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeleteAllowListCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeleteAllowListCommon Send returns without error.
//
// See DeleteAllowListCommon for more information on using the DeleteAllowListCommon
// API call, and error handling.
//
//    // Example sending a request using the DeleteAllowListCommonRequest method.
//    req, resp := client.DeleteAllowListCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *RDSPOSTGRESQL) DeleteAllowListCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDeleteAllowListCommon,
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

// DeleteAllowListCommon API operation for RDS_POSTGRESQL.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RDS_POSTGRESQL's
// API operation DeleteAllowListCommon for usage and error information.
func (c *RDSPOSTGRESQL) DeleteAllowListCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DeleteAllowListCommonRequest(input)
	return out, req.Send()
}

// DeleteAllowListCommonWithContext is the same as DeleteAllowListCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteAllowListCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RDSPOSTGRESQL) DeleteAllowListCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DeleteAllowListCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDeleteAllowList = "DeleteAllowList"

// DeleteAllowListRequest generates a "volcengine/request.Request" representing the
// client's request for the DeleteAllowList operation. The "output" return
// value will be populated with the DeleteAllowListCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeleteAllowListCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeleteAllowListCommon Send returns without error.
//
// See DeleteAllowList for more information on using the DeleteAllowList
// API call, and error handling.
//
//    // Example sending a request using the DeleteAllowListRequest method.
//    req, resp := client.DeleteAllowListRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *RDSPOSTGRESQL) DeleteAllowListRequest(input *DeleteAllowListInput) (req *request.Request, output *DeleteAllowListOutput) {
	op := &request.Operation{
		Name:       opDeleteAllowList,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteAllowListInput{}
	}

	output = &DeleteAllowListOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// DeleteAllowList API operation for RDS_POSTGRESQL.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RDS_POSTGRESQL's
// API operation DeleteAllowList for usage and error information.
func (c *RDSPOSTGRESQL) DeleteAllowList(input *DeleteAllowListInput) (*DeleteAllowListOutput, error) {
	req, out := c.DeleteAllowListRequest(input)
	return out, req.Send()
}

// DeleteAllowListWithContext is the same as DeleteAllowList with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteAllowList for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RDSPOSTGRESQL) DeleteAllowListWithContext(ctx volcengine.Context, input *DeleteAllowListInput, opts ...request.Option) (*DeleteAllowListOutput, error) {
	req, out := c.DeleteAllowListRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DeleteAllowListInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	// AllowListId is a required field
	AllowListId *string `type:"string" json:",omitempty" required:"true"`
}

// String returns the string representation
func (s DeleteAllowListInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteAllowListInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteAllowListInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DeleteAllowListInput"}
	if s.AllowListId == nil {
		invalidParams.Add(request.NewErrParamRequired("AllowListId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetAllowListId sets the AllowListId field's value.
func (s *DeleteAllowListInput) SetAllowListId(v string) *DeleteAllowListInput {
	s.AllowListId = &v
	return s
}

type DeleteAllowListOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s DeleteAllowListOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteAllowListOutput) GoString() string {
	return s.String()
}

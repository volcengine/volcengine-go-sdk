// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package rdsmysql

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opListVpcsCommon = "ListVpcs"

// ListVpcsCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ListVpcsCommon operation. The "output" return
// value will be populated with the ListVpcsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListVpcsCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListVpcsCommon Send returns without error.
//
// See ListVpcsCommon for more information on using the ListVpcsCommon
// API call, and error handling.
//
//    // Example sending a request using the ListVpcsCommonRequest method.
//    req, resp := client.ListVpcsCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *RDSMYSQL) ListVpcsCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opListVpcsCommon,
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

// ListVpcsCommon API operation for RDS_MYSQL.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RDS_MYSQL's
// API operation ListVpcsCommon for usage and error information.
func (c *RDSMYSQL) ListVpcsCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ListVpcsCommonRequest(input)
	return out, req.Send()
}

// ListVpcsCommonWithContext is the same as ListVpcsCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ListVpcsCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RDSMYSQL) ListVpcsCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ListVpcsCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opListVpcs = "ListVpcs"

// ListVpcsRequest generates a "volcengine/request.Request" representing the
// client's request for the ListVpcs operation. The "output" return
// value will be populated with the ListVpcsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListVpcsCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListVpcsCommon Send returns without error.
//
// See ListVpcs for more information on using the ListVpcs
// API call, and error handling.
//
//    // Example sending a request using the ListVpcsRequest method.
//    req, resp := client.ListVpcsRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *RDSMYSQL) ListVpcsRequest(input *ListVpcsInput) (req *request.Request, output *ListVpcsOutput) {
	op := &request.Operation{
		Name:       opListVpcs,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListVpcsInput{}
	}

	output = &ListVpcsOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// ListVpcs API operation for RDS_MYSQL.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RDS_MYSQL's
// API operation ListVpcs for usage and error information.
func (c *RDSMYSQL) ListVpcs(input *ListVpcsInput) (*ListVpcsOutput, error) {
	req, out := c.ListVpcsRequest(input)
	return out, req.Send()
}

// ListVpcsWithContext is the same as ListVpcs with the addition of
// the ability to pass a context and additional request options.
//
// See ListVpcs for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RDSMYSQL) ListVpcsWithContext(ctx volcengine.Context, input *ListVpcsInput, opts ...request.Option) (*ListVpcsOutput, error) {
	req, out := c.ListVpcsRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ListVpcsInput struct {
	_ struct{} `type:"structure"`

	Region *string `type:"string"`
}

// String returns the string representation
func (s ListVpcsInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListVpcsInput) GoString() string {
	return s.String()
}

// SetRegion sets the Region field's value.
func (s *ListVpcsInput) SetRegion(v string) *ListVpcsInput {
	s.Region = &v
	return s
}

type ListVpcsOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s ListVpcsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListVpcsOutput) GoString() string {
	return s.String()
}
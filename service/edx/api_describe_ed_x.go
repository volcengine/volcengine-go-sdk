// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package edx

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribeEDXCommon = "DescribeEDX"

// DescribeEDXCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeEDXCommon operation. The "output" return
// value will be populated with the DescribeEDXCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeEDXCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeEDXCommon Send returns without error.
//
// See DescribeEDXCommon for more information on using the DescribeEDXCommon
// API call, and error handling.
//
//    // Example sending a request using the DescribeEDXCommonRequest method.
//    req, resp := client.DescribeEDXCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *EDX) DescribeEDXCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeEDXCommon,
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

// DescribeEDXCommon API operation for EDX.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for EDX's
// API operation DescribeEDXCommon for usage and error information.
func (c *EDX) DescribeEDXCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeEDXCommonRequest(input)
	return out, req.Send()
}

// DescribeEDXCommonWithContext is the same as DescribeEDXCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeEDXCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *EDX) DescribeEDXCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeEDXCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeEDX = "DescribeEDX"

// DescribeEDXRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeEDX operation. The "output" return
// value will be populated with the DescribeEDXCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeEDXCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeEDXCommon Send returns without error.
//
// See DescribeEDX for more information on using the DescribeEDX
// API call, and error handling.
//
//    // Example sending a request using the DescribeEDXRequest method.
//    req, resp := client.DescribeEDXRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *EDX) DescribeEDXRequest(input *DescribeEDXInput) (req *request.Request, output *DescribeEDXOutput) {
	op := &request.Operation{
		Name:       opDescribeEDX,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeEDXInput{}
	}

	output = &DescribeEDXOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// DescribeEDX API operation for EDX.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for EDX's
// API operation DescribeEDX for usage and error information.
func (c *EDX) DescribeEDX(input *DescribeEDXInput) (*DescribeEDXOutput, error) {
	req, out := c.DescribeEDXRequest(input)
	return out, req.Send()
}

// DescribeEDXWithContext is the same as DescribeEDX with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeEDX for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *EDX) DescribeEDXWithContext(ctx volcengine.Context, input *DescribeEDXInput, opts ...request.Option) (*DescribeEDXOutput, error) {
	req, out := c.DescribeEDXRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DescribeEDXInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	// InstanceId is a required field
	InstanceId *string `type:"string" json:",omitempty" required:"true"`
}

// String returns the string representation
func (s DescribeEDXInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeEDXInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeEDXInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DescribeEDXInput"}
	if s.InstanceId == nil {
		invalidParams.Add(request.NewErrParamRequired("InstanceId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetInstanceId sets the InstanceId field's value.
func (s *DescribeEDXInput) SetInstanceId(v string) *DescribeEDXInput {
	s.InstanceId = &v
	return s
}

type DescribeEDXOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	EDX *ForDescribeEDXOutput `type:"structure" json:",omitempty"`
}

// String returns the string representation
func (s DescribeEDXOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeEDXOutput) GoString() string {
	return s.String()
}

// SetEDX sets the EDX field's value.
func (s *DescribeEDXOutput) SetEDX(v *ForDescribeEDXOutput) *DescribeEDXOutput {
	s.EDX = v
	return s
}

// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package edx

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opListEDXAvailableVIFVGWCommon = "ListEDXAvailableVIFVGW"

// ListEDXAvailableVIFVGWCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ListEDXAvailableVIFVGWCommon operation. The "output" return
// value will be populated with the ListEDXAvailableVIFVGWCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListEDXAvailableVIFVGWCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListEDXAvailableVIFVGWCommon Send returns without error.
//
// See ListEDXAvailableVIFVGWCommon for more information on using the ListEDXAvailableVIFVGWCommon
// API call, and error handling.
//
//    // Example sending a request using the ListEDXAvailableVIFVGWCommonRequest method.
//    req, resp := client.ListEDXAvailableVIFVGWCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *EDX) ListEDXAvailableVIFVGWCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opListEDXAvailableVIFVGWCommon,
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

// ListEDXAvailableVIFVGWCommon API operation for EDX.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for EDX's
// API operation ListEDXAvailableVIFVGWCommon for usage and error information.
func (c *EDX) ListEDXAvailableVIFVGWCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ListEDXAvailableVIFVGWCommonRequest(input)
	return out, req.Send()
}

// ListEDXAvailableVIFVGWCommonWithContext is the same as ListEDXAvailableVIFVGWCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ListEDXAvailableVIFVGWCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *EDX) ListEDXAvailableVIFVGWCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ListEDXAvailableVIFVGWCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opListEDXAvailableVIFVGW = "ListEDXAvailableVIFVGW"

// ListEDXAvailableVIFVGWRequest generates a "volcengine/request.Request" representing the
// client's request for the ListEDXAvailableVIFVGW operation. The "output" return
// value will be populated with the ListEDXAvailableVIFVGWCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListEDXAvailableVIFVGWCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListEDXAvailableVIFVGWCommon Send returns without error.
//
// See ListEDXAvailableVIFVGW for more information on using the ListEDXAvailableVIFVGW
// API call, and error handling.
//
//    // Example sending a request using the ListEDXAvailableVIFVGWRequest method.
//    req, resp := client.ListEDXAvailableVIFVGWRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *EDX) ListEDXAvailableVIFVGWRequest(input *ListEDXAvailableVIFVGWInput) (req *request.Request, output *ListEDXAvailableVIFVGWOutput) {
	op := &request.Operation{
		Name:       opListEDXAvailableVIFVGW,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListEDXAvailableVIFVGWInput{}
	}

	output = &ListEDXAvailableVIFVGWOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// ListEDXAvailableVIFVGW API operation for EDX.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for EDX's
// API operation ListEDXAvailableVIFVGW for usage and error information.
func (c *EDX) ListEDXAvailableVIFVGW(input *ListEDXAvailableVIFVGWInput) (*ListEDXAvailableVIFVGWOutput, error) {
	req, out := c.ListEDXAvailableVIFVGWRequest(input)
	return out, req.Send()
}

// ListEDXAvailableVIFVGWWithContext is the same as ListEDXAvailableVIFVGW with the addition of
// the ability to pass a context and additional request options.
//
// See ListEDXAvailableVIFVGW for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *EDX) ListEDXAvailableVIFVGWWithContext(ctx volcengine.Context, input *ListEDXAvailableVIFVGWInput, opts ...request.Option) (*ListEDXAvailableVIFVGWOutput, error) {
	req, out := c.ListEDXAvailableVIFVGWRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type AvailableVIFVGWInstanceListForListEDXAvailableVIFVGWOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	InstanceId *string `type:"string" json:",omitempty"`

	Name *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s AvailableVIFVGWInstanceListForListEDXAvailableVIFVGWOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s AvailableVIFVGWInstanceListForListEDXAvailableVIFVGWOutput) GoString() string {
	return s.String()
}

// SetInstanceId sets the InstanceId field's value.
func (s *AvailableVIFVGWInstanceListForListEDXAvailableVIFVGWOutput) SetInstanceId(v string) *AvailableVIFVGWInstanceListForListEDXAvailableVIFVGWOutput {
	s.InstanceId = &v
	return s
}

// SetName sets the Name field's value.
func (s *AvailableVIFVGWInstanceListForListEDXAvailableVIFVGWOutput) SetName(v string) *AvailableVIFVGWInstanceListForListEDXAvailableVIFVGWOutput {
	s.Name = &v
	return s
}

type ListEDXAvailableVIFVGWInput struct {
	_ struct{} `type:"structure" json:",omitempty"`
}

// String returns the string representation
func (s ListEDXAvailableVIFVGWInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListEDXAvailableVIFVGWInput) GoString() string {
	return s.String()
}

type ListEDXAvailableVIFVGWOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	AvailableVIFVGWInstanceList []*AvailableVIFVGWInstanceListForListEDXAvailableVIFVGWOutput `type:"list" json:",omitempty"`
}

// String returns the string representation
func (s ListEDXAvailableVIFVGWOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListEDXAvailableVIFVGWOutput) GoString() string {
	return s.String()
}

// SetAvailableVIFVGWInstanceList sets the AvailableVIFVGWInstanceList field's value.
func (s *ListEDXAvailableVIFVGWOutput) SetAvailableVIFVGWInstanceList(v []*AvailableVIFVGWInstanceListForListEDXAvailableVIFVGWOutput) *ListEDXAvailableVIFVGWOutput {
	s.AvailableVIFVGWInstanceList = v
	return s
}

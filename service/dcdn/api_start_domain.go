// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package dcdn

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opStartDomainCommon = "StartDomain"

// StartDomainCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the StartDomainCommon operation. The "output" return
// value will be populated with the StartDomainCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned StartDomainCommon Request to send the API call to the service.
// the "output" return value is not valid until after StartDomainCommon Send returns without error.
//
// See StartDomainCommon for more information on using the StartDomainCommon
// API call, and error handling.
//
//    // Example sending a request using the StartDomainCommonRequest method.
//    req, resp := client.StartDomainCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *DCDN) StartDomainCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opStartDomainCommon,
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

// StartDomainCommon API operation for DCDN.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for DCDN's
// API operation StartDomainCommon for usage and error information.
func (c *DCDN) StartDomainCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.StartDomainCommonRequest(input)
	return out, req.Send()
}

// StartDomainCommonWithContext is the same as StartDomainCommon with the addition of
// the ability to pass a context and additional request options.
//
// See StartDomainCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *DCDN) StartDomainCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.StartDomainCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opStartDomain = "StartDomain"

// StartDomainRequest generates a "volcengine/request.Request" representing the
// client's request for the StartDomain operation. The "output" return
// value will be populated with the StartDomainCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned StartDomainCommon Request to send the API call to the service.
// the "output" return value is not valid until after StartDomainCommon Send returns without error.
//
// See StartDomain for more information on using the StartDomain
// API call, and error handling.
//
//    // Example sending a request using the StartDomainRequest method.
//    req, resp := client.StartDomainRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *DCDN) StartDomainRequest(input *StartDomainInput) (req *request.Request, output *StartDomainOutput) {
	op := &request.Operation{
		Name:       opStartDomain,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &StartDomainInput{}
	}

	output = &StartDomainOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// StartDomain API operation for DCDN.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for DCDN's
// API operation StartDomain for usage and error information.
func (c *DCDN) StartDomain(input *StartDomainInput) (*StartDomainOutput, error) {
	req, out := c.StartDomainRequest(input)
	return out, req.Send()
}

// StartDomainWithContext is the same as StartDomain with the addition of
// the ability to pass a context and additional request options.
//
// See StartDomain for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *DCDN) StartDomainWithContext(ctx volcengine.Context, input *StartDomainInput, opts ...request.Option) (*StartDomainOutput, error) {
	req, out := c.StartDomainRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type StartDomainInput struct {
	_ struct{} `type:"structure"`

	Domains []*string `type:"list"`
}

// String returns the string representation
func (s StartDomainInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s StartDomainInput) GoString() string {
	return s.String()
}

// SetDomains sets the Domains field's value.
func (s *StartDomainInput) SetDomains(v []*string) *StartDomainInput {
	s.Domains = v
	return s
}

type StartDomainOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s StartDomainOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s StartDomainOutput) GoString() string {
	return s.String()
}
// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package transitrouter

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opRejectSharedTransitRouterCommon = "RejectSharedTransitRouter"

// RejectSharedTransitRouterCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the RejectSharedTransitRouterCommon operation. The "output" return
// value will be populated with the RejectSharedTransitRouterCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned RejectSharedTransitRouterCommon Request to send the API call to the service.
// the "output" return value is not valid until after RejectSharedTransitRouterCommon Send returns without error.
//
// See RejectSharedTransitRouterCommon for more information on using the RejectSharedTransitRouterCommon
// API call, and error handling.
//
//    // Example sending a request using the RejectSharedTransitRouterCommonRequest method.
//    req, resp := client.RejectSharedTransitRouterCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *TRANSITROUTER) RejectSharedTransitRouterCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opRejectSharedTransitRouterCommon,
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

// RejectSharedTransitRouterCommon API operation for TRANSITROUTER.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for TRANSITROUTER's
// API operation RejectSharedTransitRouterCommon for usage and error information.
func (c *TRANSITROUTER) RejectSharedTransitRouterCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.RejectSharedTransitRouterCommonRequest(input)
	return out, req.Send()
}

// RejectSharedTransitRouterCommonWithContext is the same as RejectSharedTransitRouterCommon with the addition of
// the ability to pass a context and additional request options.
//
// See RejectSharedTransitRouterCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *TRANSITROUTER) RejectSharedTransitRouterCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.RejectSharedTransitRouterCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opRejectSharedTransitRouter = "RejectSharedTransitRouter"

// RejectSharedTransitRouterRequest generates a "volcengine/request.Request" representing the
// client's request for the RejectSharedTransitRouter operation. The "output" return
// value will be populated with the RejectSharedTransitRouterCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned RejectSharedTransitRouterCommon Request to send the API call to the service.
// the "output" return value is not valid until after RejectSharedTransitRouterCommon Send returns without error.
//
// See RejectSharedTransitRouter for more information on using the RejectSharedTransitRouter
// API call, and error handling.
//
//    // Example sending a request using the RejectSharedTransitRouterRequest method.
//    req, resp := client.RejectSharedTransitRouterRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *TRANSITROUTER) RejectSharedTransitRouterRequest(input *RejectSharedTransitRouterInput) (req *request.Request, output *RejectSharedTransitRouterOutput) {
	op := &request.Operation{
		Name:       opRejectSharedTransitRouter,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &RejectSharedTransitRouterInput{}
	}

	output = &RejectSharedTransitRouterOutput{}
	req = c.newRequest(op, input, output)

	return
}

// RejectSharedTransitRouter API operation for TRANSITROUTER.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for TRANSITROUTER's
// API operation RejectSharedTransitRouter for usage and error information.
func (c *TRANSITROUTER) RejectSharedTransitRouter(input *RejectSharedTransitRouterInput) (*RejectSharedTransitRouterOutput, error) {
	req, out := c.RejectSharedTransitRouterRequest(input)
	return out, req.Send()
}

// RejectSharedTransitRouterWithContext is the same as RejectSharedTransitRouter with the addition of
// the ability to pass a context and additional request options.
//
// See RejectSharedTransitRouter for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *TRANSITROUTER) RejectSharedTransitRouterWithContext(ctx volcengine.Context, input *RejectSharedTransitRouterInput, opts ...request.Option) (*RejectSharedTransitRouterOutput, error) {
	req, out := c.RejectSharedTransitRouterRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type RejectSharedTransitRouterInput struct {
	_ struct{} `type:"structure"`

	// TransitRouterId is a required field
	TransitRouterId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s RejectSharedTransitRouterInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s RejectSharedTransitRouterInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RejectSharedTransitRouterInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "RejectSharedTransitRouterInput"}
	if s.TransitRouterId == nil {
		invalidParams.Add(request.NewErrParamRequired("TransitRouterId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetTransitRouterId sets the TransitRouterId field's value.
func (s *RejectSharedTransitRouterInput) SetTransitRouterId(v string) *RejectSharedTransitRouterInput {
	s.TransitRouterId = &v
	return s
}

type RejectSharedTransitRouterOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s RejectSharedTransitRouterOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s RejectSharedTransitRouterOutput) GoString() string {
	return s.String()
}

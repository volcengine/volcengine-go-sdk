// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package transitrouter

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDeleteTransitRouterForwardPolicyEntryCommon = "DeleteTransitRouterForwardPolicyEntry"

// DeleteTransitRouterForwardPolicyEntryCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DeleteTransitRouterForwardPolicyEntryCommon operation. The "output" return
// value will be populated with the DeleteTransitRouterForwardPolicyEntryCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeleteTransitRouterForwardPolicyEntryCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeleteTransitRouterForwardPolicyEntryCommon Send returns without error.
//
// See DeleteTransitRouterForwardPolicyEntryCommon for more information on using the DeleteTransitRouterForwardPolicyEntryCommon
// API call, and error handling.
//
//    // Example sending a request using the DeleteTransitRouterForwardPolicyEntryCommonRequest method.
//    req, resp := client.DeleteTransitRouterForwardPolicyEntryCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *TRANSITROUTER) DeleteTransitRouterForwardPolicyEntryCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDeleteTransitRouterForwardPolicyEntryCommon,
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

// DeleteTransitRouterForwardPolicyEntryCommon API operation for TRANSITROUTER.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for TRANSITROUTER's
// API operation DeleteTransitRouterForwardPolicyEntryCommon for usage and error information.
func (c *TRANSITROUTER) DeleteTransitRouterForwardPolicyEntryCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DeleteTransitRouterForwardPolicyEntryCommonRequest(input)
	return out, req.Send()
}

// DeleteTransitRouterForwardPolicyEntryCommonWithContext is the same as DeleteTransitRouterForwardPolicyEntryCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteTransitRouterForwardPolicyEntryCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *TRANSITROUTER) DeleteTransitRouterForwardPolicyEntryCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DeleteTransitRouterForwardPolicyEntryCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDeleteTransitRouterForwardPolicyEntry = "DeleteTransitRouterForwardPolicyEntry"

// DeleteTransitRouterForwardPolicyEntryRequest generates a "volcengine/request.Request" representing the
// client's request for the DeleteTransitRouterForwardPolicyEntry operation. The "output" return
// value will be populated with the DeleteTransitRouterForwardPolicyEntryCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeleteTransitRouterForwardPolicyEntryCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeleteTransitRouterForwardPolicyEntryCommon Send returns without error.
//
// See DeleteTransitRouterForwardPolicyEntry for more information on using the DeleteTransitRouterForwardPolicyEntry
// API call, and error handling.
//
//    // Example sending a request using the DeleteTransitRouterForwardPolicyEntryRequest method.
//    req, resp := client.DeleteTransitRouterForwardPolicyEntryRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *TRANSITROUTER) DeleteTransitRouterForwardPolicyEntryRequest(input *DeleteTransitRouterForwardPolicyEntryInput) (req *request.Request, output *DeleteTransitRouterForwardPolicyEntryOutput) {
	op := &request.Operation{
		Name:       opDeleteTransitRouterForwardPolicyEntry,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteTransitRouterForwardPolicyEntryInput{}
	}

	output = &DeleteTransitRouterForwardPolicyEntryOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DeleteTransitRouterForwardPolicyEntry API operation for TRANSITROUTER.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for TRANSITROUTER's
// API operation DeleteTransitRouterForwardPolicyEntry for usage and error information.
func (c *TRANSITROUTER) DeleteTransitRouterForwardPolicyEntry(input *DeleteTransitRouterForwardPolicyEntryInput) (*DeleteTransitRouterForwardPolicyEntryOutput, error) {
	req, out := c.DeleteTransitRouterForwardPolicyEntryRequest(input)
	return out, req.Send()
}

// DeleteTransitRouterForwardPolicyEntryWithContext is the same as DeleteTransitRouterForwardPolicyEntry with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteTransitRouterForwardPolicyEntry for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *TRANSITROUTER) DeleteTransitRouterForwardPolicyEntryWithContext(ctx volcengine.Context, input *DeleteTransitRouterForwardPolicyEntryInput, opts ...request.Option) (*DeleteTransitRouterForwardPolicyEntryOutput, error) {
	req, out := c.DeleteTransitRouterForwardPolicyEntryRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DeleteTransitRouterForwardPolicyEntryInput struct {
	_ struct{} `type:"structure"`

	// TransitRouterForwardPolicyEntryId is a required field
	TransitRouterForwardPolicyEntryId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteTransitRouterForwardPolicyEntryInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteTransitRouterForwardPolicyEntryInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteTransitRouterForwardPolicyEntryInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DeleteTransitRouterForwardPolicyEntryInput"}
	if s.TransitRouterForwardPolicyEntryId == nil {
		invalidParams.Add(request.NewErrParamRequired("TransitRouterForwardPolicyEntryId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetTransitRouterForwardPolicyEntryId sets the TransitRouterForwardPolicyEntryId field's value.
func (s *DeleteTransitRouterForwardPolicyEntryInput) SetTransitRouterForwardPolicyEntryId(v string) *DeleteTransitRouterForwardPolicyEntryInput {
	s.TransitRouterForwardPolicyEntryId = &v
	return s
}

type DeleteTransitRouterForwardPolicyEntryOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s DeleteTransitRouterForwardPolicyEntryOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteTransitRouterForwardPolicyEntryOutput) GoString() string {
	return s.String()
}

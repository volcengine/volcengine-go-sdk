// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package transitrouter

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDeleteTransitRouterRouteEntryCommon = "DeleteTransitRouterRouteEntry"

// DeleteTransitRouterRouteEntryCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DeleteTransitRouterRouteEntryCommon operation. The "output" return
// value will be populated with the DeleteTransitRouterRouteEntryCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeleteTransitRouterRouteEntryCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeleteTransitRouterRouteEntryCommon Send returns without error.
//
// See DeleteTransitRouterRouteEntryCommon for more information on using the DeleteTransitRouterRouteEntryCommon
// API call, and error handling.
//
//    // Example sending a request using the DeleteTransitRouterRouteEntryCommonRequest method.
//    req, resp := client.DeleteTransitRouterRouteEntryCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *TRANSITROUTER) DeleteTransitRouterRouteEntryCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDeleteTransitRouterRouteEntryCommon,
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

// DeleteTransitRouterRouteEntryCommon API operation for TRANSITROUTER.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for TRANSITROUTER's
// API operation DeleteTransitRouterRouteEntryCommon for usage and error information.
func (c *TRANSITROUTER) DeleteTransitRouterRouteEntryCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DeleteTransitRouterRouteEntryCommonRequest(input)
	return out, req.Send()
}

// DeleteTransitRouterRouteEntryCommonWithContext is the same as DeleteTransitRouterRouteEntryCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteTransitRouterRouteEntryCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *TRANSITROUTER) DeleteTransitRouterRouteEntryCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DeleteTransitRouterRouteEntryCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDeleteTransitRouterRouteEntry = "DeleteTransitRouterRouteEntry"

// DeleteTransitRouterRouteEntryRequest generates a "volcengine/request.Request" representing the
// client's request for the DeleteTransitRouterRouteEntry operation. The "output" return
// value will be populated with the DeleteTransitRouterRouteEntryCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeleteTransitRouterRouteEntryCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeleteTransitRouterRouteEntryCommon Send returns without error.
//
// See DeleteTransitRouterRouteEntry for more information on using the DeleteTransitRouterRouteEntry
// API call, and error handling.
//
//    // Example sending a request using the DeleteTransitRouterRouteEntryRequest method.
//    req, resp := client.DeleteTransitRouterRouteEntryRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *TRANSITROUTER) DeleteTransitRouterRouteEntryRequest(input *DeleteTransitRouterRouteEntryInput) (req *request.Request, output *DeleteTransitRouterRouteEntryOutput) {
	op := &request.Operation{
		Name:       opDeleteTransitRouterRouteEntry,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteTransitRouterRouteEntryInput{}
	}

	output = &DeleteTransitRouterRouteEntryOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DeleteTransitRouterRouteEntry API operation for TRANSITROUTER.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for TRANSITROUTER's
// API operation DeleteTransitRouterRouteEntry for usage and error information.
func (c *TRANSITROUTER) DeleteTransitRouterRouteEntry(input *DeleteTransitRouterRouteEntryInput) (*DeleteTransitRouterRouteEntryOutput, error) {
	req, out := c.DeleteTransitRouterRouteEntryRequest(input)
	return out, req.Send()
}

// DeleteTransitRouterRouteEntryWithContext is the same as DeleteTransitRouterRouteEntry with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteTransitRouterRouteEntry for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *TRANSITROUTER) DeleteTransitRouterRouteEntryWithContext(ctx volcengine.Context, input *DeleteTransitRouterRouteEntryInput, opts ...request.Option) (*DeleteTransitRouterRouteEntryOutput, error) {
	req, out := c.DeleteTransitRouterRouteEntryRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DeleteTransitRouterRouteEntryInput struct {
	_ struct{} `type:"structure"`

	// TransitRouterRouteEntryId is a required field
	TransitRouterRouteEntryId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteTransitRouterRouteEntryInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteTransitRouterRouteEntryInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteTransitRouterRouteEntryInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DeleteTransitRouterRouteEntryInput"}
	if s.TransitRouterRouteEntryId == nil {
		invalidParams.Add(request.NewErrParamRequired("TransitRouterRouteEntryId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetTransitRouterRouteEntryId sets the TransitRouterRouteEntryId field's value.
func (s *DeleteTransitRouterRouteEntryInput) SetTransitRouterRouteEntryId(v string) *DeleteTransitRouterRouteEntryInput {
	s.TransitRouterRouteEntryId = &v
	return s
}

type DeleteTransitRouterRouteEntryOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s DeleteTransitRouterRouteEntryOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteTransitRouterRouteEntryOutput) GoString() string {
	return s.String()
}
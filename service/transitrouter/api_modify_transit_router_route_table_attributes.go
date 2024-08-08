// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package transitrouter

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opModifyTransitRouterRouteTableAttributesCommon = "ModifyTransitRouterRouteTableAttributes"

// ModifyTransitRouterRouteTableAttributesCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ModifyTransitRouterRouteTableAttributesCommon operation. The "output" return
// value will be populated with the ModifyTransitRouterRouteTableAttributesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyTransitRouterRouteTableAttributesCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyTransitRouterRouteTableAttributesCommon Send returns without error.
//
// See ModifyTransitRouterRouteTableAttributesCommon for more information on using the ModifyTransitRouterRouteTableAttributesCommon
// API call, and error handling.
//
//    // Example sending a request using the ModifyTransitRouterRouteTableAttributesCommonRequest method.
//    req, resp := client.ModifyTransitRouterRouteTableAttributesCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *TRANSITROUTER) ModifyTransitRouterRouteTableAttributesCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opModifyTransitRouterRouteTableAttributesCommon,
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

// ModifyTransitRouterRouteTableAttributesCommon API operation for TRANSITROUTER.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for TRANSITROUTER's
// API operation ModifyTransitRouterRouteTableAttributesCommon for usage and error information.
func (c *TRANSITROUTER) ModifyTransitRouterRouteTableAttributesCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ModifyTransitRouterRouteTableAttributesCommonRequest(input)
	return out, req.Send()
}

// ModifyTransitRouterRouteTableAttributesCommonWithContext is the same as ModifyTransitRouterRouteTableAttributesCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyTransitRouterRouteTableAttributesCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *TRANSITROUTER) ModifyTransitRouterRouteTableAttributesCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ModifyTransitRouterRouteTableAttributesCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opModifyTransitRouterRouteTableAttributes = "ModifyTransitRouterRouteTableAttributes"

// ModifyTransitRouterRouteTableAttributesRequest generates a "volcengine/request.Request" representing the
// client's request for the ModifyTransitRouterRouteTableAttributes operation. The "output" return
// value will be populated with the ModifyTransitRouterRouteTableAttributesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyTransitRouterRouteTableAttributesCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyTransitRouterRouteTableAttributesCommon Send returns without error.
//
// See ModifyTransitRouterRouteTableAttributes for more information on using the ModifyTransitRouterRouteTableAttributes
// API call, and error handling.
//
//    // Example sending a request using the ModifyTransitRouterRouteTableAttributesRequest method.
//    req, resp := client.ModifyTransitRouterRouteTableAttributesRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *TRANSITROUTER) ModifyTransitRouterRouteTableAttributesRequest(input *ModifyTransitRouterRouteTableAttributesInput) (req *request.Request, output *ModifyTransitRouterRouteTableAttributesOutput) {
	op := &request.Operation{
		Name:       opModifyTransitRouterRouteTableAttributes,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ModifyTransitRouterRouteTableAttributesInput{}
	}

	output = &ModifyTransitRouterRouteTableAttributesOutput{}
	req = c.newRequest(op, input, output)

	return
}

// ModifyTransitRouterRouteTableAttributes API operation for TRANSITROUTER.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for TRANSITROUTER's
// API operation ModifyTransitRouterRouteTableAttributes for usage and error information.
func (c *TRANSITROUTER) ModifyTransitRouterRouteTableAttributes(input *ModifyTransitRouterRouteTableAttributesInput) (*ModifyTransitRouterRouteTableAttributesOutput, error) {
	req, out := c.ModifyTransitRouterRouteTableAttributesRequest(input)
	return out, req.Send()
}

// ModifyTransitRouterRouteTableAttributesWithContext is the same as ModifyTransitRouterRouteTableAttributes with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyTransitRouterRouteTableAttributes for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *TRANSITROUTER) ModifyTransitRouterRouteTableAttributesWithContext(ctx volcengine.Context, input *ModifyTransitRouterRouteTableAttributesInput, opts ...request.Option) (*ModifyTransitRouterRouteTableAttributesOutput, error) {
	req, out := c.ModifyTransitRouterRouteTableAttributesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ModifyTransitRouterRouteTableAttributesInput struct {
	_ struct{} `type:"structure"`

	Description *string `type:"string"`

	// TransitRouterRouteTableId is a required field
	TransitRouterRouteTableId *string `type:"string" required:"true"`

	TransitRouterRouteTableName *string `type:"string"`
}

// String returns the string representation
func (s ModifyTransitRouterRouteTableAttributesInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyTransitRouterRouteTableAttributesInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ModifyTransitRouterRouteTableAttributesInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ModifyTransitRouterRouteTableAttributesInput"}
	if s.TransitRouterRouteTableId == nil {
		invalidParams.Add(request.NewErrParamRequired("TransitRouterRouteTableId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetDescription sets the Description field's value.
func (s *ModifyTransitRouterRouteTableAttributesInput) SetDescription(v string) *ModifyTransitRouterRouteTableAttributesInput {
	s.Description = &v
	return s
}

// SetTransitRouterRouteTableId sets the TransitRouterRouteTableId field's value.
func (s *ModifyTransitRouterRouteTableAttributesInput) SetTransitRouterRouteTableId(v string) *ModifyTransitRouterRouteTableAttributesInput {
	s.TransitRouterRouteTableId = &v
	return s
}

// SetTransitRouterRouteTableName sets the TransitRouterRouteTableName field's value.
func (s *ModifyTransitRouterRouteTableAttributesInput) SetTransitRouterRouteTableName(v string) *ModifyTransitRouterRouteTableAttributesInput {
	s.TransitRouterRouteTableName = &v
	return s
}

type ModifyTransitRouterRouteTableAttributesOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s ModifyTransitRouterRouteTableAttributesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyTransitRouterRouteTableAttributesOutput) GoString() string {
	return s.String()
}
// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package transitrouter

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opModifyTransitRouterForwardPolicyTableAttributesCommon = "ModifyTransitRouterForwardPolicyTableAttributes"

// ModifyTransitRouterForwardPolicyTableAttributesCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ModifyTransitRouterForwardPolicyTableAttributesCommon operation. The "output" return
// value will be populated with the ModifyTransitRouterForwardPolicyTableAttributesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyTransitRouterForwardPolicyTableAttributesCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyTransitRouterForwardPolicyTableAttributesCommon Send returns without error.
//
// See ModifyTransitRouterForwardPolicyTableAttributesCommon for more information on using the ModifyTransitRouterForwardPolicyTableAttributesCommon
// API call, and error handling.
//
//    // Example sending a request using the ModifyTransitRouterForwardPolicyTableAttributesCommonRequest method.
//    req, resp := client.ModifyTransitRouterForwardPolicyTableAttributesCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *TRANSITROUTER) ModifyTransitRouterForwardPolicyTableAttributesCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opModifyTransitRouterForwardPolicyTableAttributesCommon,
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

// ModifyTransitRouterForwardPolicyTableAttributesCommon API operation for TRANSITROUTER.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for TRANSITROUTER's
// API operation ModifyTransitRouterForwardPolicyTableAttributesCommon for usage and error information.
func (c *TRANSITROUTER) ModifyTransitRouterForwardPolicyTableAttributesCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ModifyTransitRouterForwardPolicyTableAttributesCommonRequest(input)
	return out, req.Send()
}

// ModifyTransitRouterForwardPolicyTableAttributesCommonWithContext is the same as ModifyTransitRouterForwardPolicyTableAttributesCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyTransitRouterForwardPolicyTableAttributesCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *TRANSITROUTER) ModifyTransitRouterForwardPolicyTableAttributesCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ModifyTransitRouterForwardPolicyTableAttributesCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opModifyTransitRouterForwardPolicyTableAttributes = "ModifyTransitRouterForwardPolicyTableAttributes"

// ModifyTransitRouterForwardPolicyTableAttributesRequest generates a "volcengine/request.Request" representing the
// client's request for the ModifyTransitRouterForwardPolicyTableAttributes operation. The "output" return
// value will be populated with the ModifyTransitRouterForwardPolicyTableAttributesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyTransitRouterForwardPolicyTableAttributesCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyTransitRouterForwardPolicyTableAttributesCommon Send returns without error.
//
// See ModifyTransitRouterForwardPolicyTableAttributes for more information on using the ModifyTransitRouterForwardPolicyTableAttributes
// API call, and error handling.
//
//    // Example sending a request using the ModifyTransitRouterForwardPolicyTableAttributesRequest method.
//    req, resp := client.ModifyTransitRouterForwardPolicyTableAttributesRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *TRANSITROUTER) ModifyTransitRouterForwardPolicyTableAttributesRequest(input *ModifyTransitRouterForwardPolicyTableAttributesInput) (req *request.Request, output *ModifyTransitRouterForwardPolicyTableAttributesOutput) {
	op := &request.Operation{
		Name:       opModifyTransitRouterForwardPolicyTableAttributes,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ModifyTransitRouterForwardPolicyTableAttributesInput{}
	}

	output = &ModifyTransitRouterForwardPolicyTableAttributesOutput{}
	req = c.newRequest(op, input, output)

	return
}

// ModifyTransitRouterForwardPolicyTableAttributes API operation for TRANSITROUTER.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for TRANSITROUTER's
// API operation ModifyTransitRouterForwardPolicyTableAttributes for usage and error information.
func (c *TRANSITROUTER) ModifyTransitRouterForwardPolicyTableAttributes(input *ModifyTransitRouterForwardPolicyTableAttributesInput) (*ModifyTransitRouterForwardPolicyTableAttributesOutput, error) {
	req, out := c.ModifyTransitRouterForwardPolicyTableAttributesRequest(input)
	return out, req.Send()
}

// ModifyTransitRouterForwardPolicyTableAttributesWithContext is the same as ModifyTransitRouterForwardPolicyTableAttributes with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyTransitRouterForwardPolicyTableAttributes for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *TRANSITROUTER) ModifyTransitRouterForwardPolicyTableAttributesWithContext(ctx volcengine.Context, input *ModifyTransitRouterForwardPolicyTableAttributesInput, opts ...request.Option) (*ModifyTransitRouterForwardPolicyTableAttributesOutput, error) {
	req, out := c.ModifyTransitRouterForwardPolicyTableAttributesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ModifyTransitRouterForwardPolicyTableAttributesInput struct {
	_ struct{} `type:"structure"`

	Description *string `type:"string"`

	// TransitRouterForwardPolicyTableId is a required field
	TransitRouterForwardPolicyTableId *string `type:"string" required:"true"`

	TransitRouterForwardPolicyTableName *string `type:"string"`
}

// String returns the string representation
func (s ModifyTransitRouterForwardPolicyTableAttributesInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyTransitRouterForwardPolicyTableAttributesInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ModifyTransitRouterForwardPolicyTableAttributesInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ModifyTransitRouterForwardPolicyTableAttributesInput"}
	if s.TransitRouterForwardPolicyTableId == nil {
		invalidParams.Add(request.NewErrParamRequired("TransitRouterForwardPolicyTableId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetDescription sets the Description field's value.
func (s *ModifyTransitRouterForwardPolicyTableAttributesInput) SetDescription(v string) *ModifyTransitRouterForwardPolicyTableAttributesInput {
	s.Description = &v
	return s
}

// SetTransitRouterForwardPolicyTableId sets the TransitRouterForwardPolicyTableId field's value.
func (s *ModifyTransitRouterForwardPolicyTableAttributesInput) SetTransitRouterForwardPolicyTableId(v string) *ModifyTransitRouterForwardPolicyTableAttributesInput {
	s.TransitRouterForwardPolicyTableId = &v
	return s
}

// SetTransitRouterForwardPolicyTableName sets the TransitRouterForwardPolicyTableName field's value.
func (s *ModifyTransitRouterForwardPolicyTableAttributesInput) SetTransitRouterForwardPolicyTableName(v string) *ModifyTransitRouterForwardPolicyTableAttributesInput {
	s.TransitRouterForwardPolicyTableName = &v
	return s
}

type ModifyTransitRouterForwardPolicyTableAttributesOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s ModifyTransitRouterForwardPolicyTableAttributesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyTransitRouterForwardPolicyTableAttributesOutput) GoString() string {
	return s.String()
}

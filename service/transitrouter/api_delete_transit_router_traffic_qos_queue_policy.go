// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package transitrouter

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDeleteTransitRouterTrafficQosQueuePolicyCommon = "DeleteTransitRouterTrafficQosQueuePolicy"

// DeleteTransitRouterTrafficQosQueuePolicyCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DeleteTransitRouterTrafficQosQueuePolicyCommon operation. The "output" return
// value will be populated with the DeleteTransitRouterTrafficQosQueuePolicyCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeleteTransitRouterTrafficQosQueuePolicyCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeleteTransitRouterTrafficQosQueuePolicyCommon Send returns without error.
//
// See DeleteTransitRouterTrafficQosQueuePolicyCommon for more information on using the DeleteTransitRouterTrafficQosQueuePolicyCommon
// API call, and error handling.
//
//    // Example sending a request using the DeleteTransitRouterTrafficQosQueuePolicyCommonRequest method.
//    req, resp := client.DeleteTransitRouterTrafficQosQueuePolicyCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *TRANSITROUTER) DeleteTransitRouterTrafficQosQueuePolicyCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDeleteTransitRouterTrafficQosQueuePolicyCommon,
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

// DeleteTransitRouterTrafficQosQueuePolicyCommon API operation for TRANSITROUTER.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for TRANSITROUTER's
// API operation DeleteTransitRouterTrafficQosQueuePolicyCommon for usage and error information.
func (c *TRANSITROUTER) DeleteTransitRouterTrafficQosQueuePolicyCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DeleteTransitRouterTrafficQosQueuePolicyCommonRequest(input)
	return out, req.Send()
}

// DeleteTransitRouterTrafficQosQueuePolicyCommonWithContext is the same as DeleteTransitRouterTrafficQosQueuePolicyCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteTransitRouterTrafficQosQueuePolicyCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *TRANSITROUTER) DeleteTransitRouterTrafficQosQueuePolicyCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DeleteTransitRouterTrafficQosQueuePolicyCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDeleteTransitRouterTrafficQosQueuePolicy = "DeleteTransitRouterTrafficQosQueuePolicy"

// DeleteTransitRouterTrafficQosQueuePolicyRequest generates a "volcengine/request.Request" representing the
// client's request for the DeleteTransitRouterTrafficQosQueuePolicy operation. The "output" return
// value will be populated with the DeleteTransitRouterTrafficQosQueuePolicyCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeleteTransitRouterTrafficQosQueuePolicyCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeleteTransitRouterTrafficQosQueuePolicyCommon Send returns without error.
//
// See DeleteTransitRouterTrafficQosQueuePolicy for more information on using the DeleteTransitRouterTrafficQosQueuePolicy
// API call, and error handling.
//
//    // Example sending a request using the DeleteTransitRouterTrafficQosQueuePolicyRequest method.
//    req, resp := client.DeleteTransitRouterTrafficQosQueuePolicyRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *TRANSITROUTER) DeleteTransitRouterTrafficQosQueuePolicyRequest(input *DeleteTransitRouterTrafficQosQueuePolicyInput) (req *request.Request, output *DeleteTransitRouterTrafficQosQueuePolicyOutput) {
	op := &request.Operation{
		Name:       opDeleteTransitRouterTrafficQosQueuePolicy,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteTransitRouterTrafficQosQueuePolicyInput{}
	}

	output = &DeleteTransitRouterTrafficQosQueuePolicyOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DeleteTransitRouterTrafficQosQueuePolicy API operation for TRANSITROUTER.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for TRANSITROUTER's
// API operation DeleteTransitRouterTrafficQosQueuePolicy for usage and error information.
func (c *TRANSITROUTER) DeleteTransitRouterTrafficQosQueuePolicy(input *DeleteTransitRouterTrafficQosQueuePolicyInput) (*DeleteTransitRouterTrafficQosQueuePolicyOutput, error) {
	req, out := c.DeleteTransitRouterTrafficQosQueuePolicyRequest(input)
	return out, req.Send()
}

// DeleteTransitRouterTrafficQosQueuePolicyWithContext is the same as DeleteTransitRouterTrafficQosQueuePolicy with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteTransitRouterTrafficQosQueuePolicy for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *TRANSITROUTER) DeleteTransitRouterTrafficQosQueuePolicyWithContext(ctx volcengine.Context, input *DeleteTransitRouterTrafficQosQueuePolicyInput, opts ...request.Option) (*DeleteTransitRouterTrafficQosQueuePolicyOutput, error) {
	req, out := c.DeleteTransitRouterTrafficQosQueuePolicyRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DeleteTransitRouterTrafficQosQueuePolicyInput struct {
	_ struct{} `type:"structure"`

	// TransitRouterTrafficQosQueuePolicyId is a required field
	TransitRouterTrafficQosQueuePolicyId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteTransitRouterTrafficQosQueuePolicyInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteTransitRouterTrafficQosQueuePolicyInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteTransitRouterTrafficQosQueuePolicyInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DeleteTransitRouterTrafficQosQueuePolicyInput"}
	if s.TransitRouterTrafficQosQueuePolicyId == nil {
		invalidParams.Add(request.NewErrParamRequired("TransitRouterTrafficQosQueuePolicyId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetTransitRouterTrafficQosQueuePolicyId sets the TransitRouterTrafficQosQueuePolicyId field's value.
func (s *DeleteTransitRouterTrafficQosQueuePolicyInput) SetTransitRouterTrafficQosQueuePolicyId(v string) *DeleteTransitRouterTrafficQosQueuePolicyInput {
	s.TransitRouterTrafficQosQueuePolicyId = &v
	return s
}

type DeleteTransitRouterTrafficQosQueuePolicyOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s DeleteTransitRouterTrafficQosQueuePolicyOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteTransitRouterTrafficQosQueuePolicyOutput) GoString() string {
	return s.String()
}
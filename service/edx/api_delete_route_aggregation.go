// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package edx

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDeleteRouteAggregationCommon = "DeleteRouteAggregation"

// DeleteRouteAggregationCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DeleteRouteAggregationCommon operation. The "output" return
// value will be populated with the DeleteRouteAggregationCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeleteRouteAggregationCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeleteRouteAggregationCommon Send returns without error.
//
// See DeleteRouteAggregationCommon for more information on using the DeleteRouteAggregationCommon
// API call, and error handling.
//
//    // Example sending a request using the DeleteRouteAggregationCommonRequest method.
//    req, resp := client.DeleteRouteAggregationCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *EDX) DeleteRouteAggregationCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDeleteRouteAggregationCommon,
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

// DeleteRouteAggregationCommon API operation for EDX.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for EDX's
// API operation DeleteRouteAggregationCommon for usage and error information.
func (c *EDX) DeleteRouteAggregationCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DeleteRouteAggregationCommonRequest(input)
	return out, req.Send()
}

// DeleteRouteAggregationCommonWithContext is the same as DeleteRouteAggregationCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteRouteAggregationCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *EDX) DeleteRouteAggregationCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DeleteRouteAggregationCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDeleteRouteAggregation = "DeleteRouteAggregation"

// DeleteRouteAggregationRequest generates a "volcengine/request.Request" representing the
// client's request for the DeleteRouteAggregation operation. The "output" return
// value will be populated with the DeleteRouteAggregationCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeleteRouteAggregationCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeleteRouteAggregationCommon Send returns without error.
//
// See DeleteRouteAggregation for more information on using the DeleteRouteAggregation
// API call, and error handling.
//
//    // Example sending a request using the DeleteRouteAggregationRequest method.
//    req, resp := client.DeleteRouteAggregationRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *EDX) DeleteRouteAggregationRequest(input *DeleteRouteAggregationInput) (req *request.Request, output *DeleteRouteAggregationOutput) {
	op := &request.Operation{
		Name:       opDeleteRouteAggregation,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteRouteAggregationInput{}
	}

	output = &DeleteRouteAggregationOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// DeleteRouteAggregation API operation for EDX.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for EDX's
// API operation DeleteRouteAggregation for usage and error information.
func (c *EDX) DeleteRouteAggregation(input *DeleteRouteAggregationInput) (*DeleteRouteAggregationOutput, error) {
	req, out := c.DeleteRouteAggregationRequest(input)
	return out, req.Send()
}

// DeleteRouteAggregationWithContext is the same as DeleteRouteAggregation with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteRouteAggregation for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *EDX) DeleteRouteAggregationWithContext(ctx volcengine.Context, input *DeleteRouteAggregationInput, opts ...request.Option) (*DeleteRouteAggregationOutput, error) {
	req, out := c.DeleteRouteAggregationRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DeleteRouteAggregationInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	// AggregationCidr is a required field
	AggregationCidr *string `type:"string" json:",omitempty" required:"true"`

	// Direction is a required field
	Direction *string `type:"string" json:",omitempty" required:"true"`

	// VIFInstanceId is a required field
	VIFInstanceId *string `type:"string" json:",omitempty" required:"true"`
}

// String returns the string representation
func (s DeleteRouteAggregationInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteRouteAggregationInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteRouteAggregationInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DeleteRouteAggregationInput"}
	if s.AggregationCidr == nil {
		invalidParams.Add(request.NewErrParamRequired("AggregationCidr"))
	}
	if s.Direction == nil {
		invalidParams.Add(request.NewErrParamRequired("Direction"))
	}
	if s.VIFInstanceId == nil {
		invalidParams.Add(request.NewErrParamRequired("VIFInstanceId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetAggregationCidr sets the AggregationCidr field's value.
func (s *DeleteRouteAggregationInput) SetAggregationCidr(v string) *DeleteRouteAggregationInput {
	s.AggregationCidr = &v
	return s
}

// SetDirection sets the Direction field's value.
func (s *DeleteRouteAggregationInput) SetDirection(v string) *DeleteRouteAggregationInput {
	s.Direction = &v
	return s
}

// SetVIFInstanceId sets the VIFInstanceId field's value.
func (s *DeleteRouteAggregationInput) SetVIFInstanceId(v string) *DeleteRouteAggregationInput {
	s.VIFInstanceId = &v
	return s
}

type DeleteRouteAggregationOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s DeleteRouteAggregationOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteRouteAggregationOutput) GoString() string {
	return s.String()
}

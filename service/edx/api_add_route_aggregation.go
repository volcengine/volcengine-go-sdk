// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package edx

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opAddRouteAggregationCommon = "AddRouteAggregation"

// AddRouteAggregationCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the AddRouteAggregationCommon operation. The "output" return
// value will be populated with the AddRouteAggregationCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned AddRouteAggregationCommon Request to send the API call to the service.
// the "output" return value is not valid until after AddRouteAggregationCommon Send returns without error.
//
// See AddRouteAggregationCommon for more information on using the AddRouteAggregationCommon
// API call, and error handling.
//
//    // Example sending a request using the AddRouteAggregationCommonRequest method.
//    req, resp := client.AddRouteAggregationCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *EDX) AddRouteAggregationCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opAddRouteAggregationCommon,
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

// AddRouteAggregationCommon API operation for EDX.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for EDX's
// API operation AddRouteAggregationCommon for usage and error information.
func (c *EDX) AddRouteAggregationCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.AddRouteAggregationCommonRequest(input)
	return out, req.Send()
}

// AddRouteAggregationCommonWithContext is the same as AddRouteAggregationCommon with the addition of
// the ability to pass a context and additional request options.
//
// See AddRouteAggregationCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *EDX) AddRouteAggregationCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.AddRouteAggregationCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opAddRouteAggregation = "AddRouteAggregation"

// AddRouteAggregationRequest generates a "volcengine/request.Request" representing the
// client's request for the AddRouteAggregation operation. The "output" return
// value will be populated with the AddRouteAggregationCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned AddRouteAggregationCommon Request to send the API call to the service.
// the "output" return value is not valid until after AddRouteAggregationCommon Send returns without error.
//
// See AddRouteAggregation for more information on using the AddRouteAggregation
// API call, and error handling.
//
//    // Example sending a request using the AddRouteAggregationRequest method.
//    req, resp := client.AddRouteAggregationRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *EDX) AddRouteAggregationRequest(input *AddRouteAggregationInput) (req *request.Request, output *AddRouteAggregationOutput) {
	op := &request.Operation{
		Name:       opAddRouteAggregation,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &AddRouteAggregationInput{}
	}

	output = &AddRouteAggregationOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// AddRouteAggregation API operation for EDX.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for EDX's
// API operation AddRouteAggregation for usage and error information.
func (c *EDX) AddRouteAggregation(input *AddRouteAggregationInput) (*AddRouteAggregationOutput, error) {
	req, out := c.AddRouteAggregationRequest(input)
	return out, req.Send()
}

// AddRouteAggregationWithContext is the same as AddRouteAggregation with the addition of
// the ability to pass a context and additional request options.
//
// See AddRouteAggregation for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *EDX) AddRouteAggregationWithContext(ctx volcengine.Context, input *AddRouteAggregationInput, opts ...request.Option) (*AddRouteAggregationOutput, error) {
	req, out := c.AddRouteAggregationRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type AddRouteAggregationInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	// AggregationCidr is a required field
	AggregationCidr *string `type:"string" json:",omitempty" required:"true"`

	Description *string `type:"string" json:",omitempty"`

	DetailSuppressed *bool `type:"boolean" json:",omitempty"`

	// Direction is a required field
	Direction *string `type:"string" json:",omitempty" required:"true"`

	// VIFInstanceId is a required field
	VIFInstanceId *string `type:"string" json:",omitempty" required:"true"`
}

// String returns the string representation
func (s AddRouteAggregationInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s AddRouteAggregationInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AddRouteAggregationInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "AddRouteAggregationInput"}
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
func (s *AddRouteAggregationInput) SetAggregationCidr(v string) *AddRouteAggregationInput {
	s.AggregationCidr = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *AddRouteAggregationInput) SetDescription(v string) *AddRouteAggregationInput {
	s.Description = &v
	return s
}

// SetDetailSuppressed sets the DetailSuppressed field's value.
func (s *AddRouteAggregationInput) SetDetailSuppressed(v bool) *AddRouteAggregationInput {
	s.DetailSuppressed = &v
	return s
}

// SetDirection sets the Direction field's value.
func (s *AddRouteAggregationInput) SetDirection(v string) *AddRouteAggregationInput {
	s.Direction = &v
	return s
}

// SetVIFInstanceId sets the VIFInstanceId field's value.
func (s *AddRouteAggregationInput) SetVIFInstanceId(v string) *AddRouteAggregationInput {
	s.VIFInstanceId = &v
	return s
}

type AddRouteAggregationOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s AddRouteAggregationOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s AddRouteAggregationOutput) GoString() string {
	return s.String()
}

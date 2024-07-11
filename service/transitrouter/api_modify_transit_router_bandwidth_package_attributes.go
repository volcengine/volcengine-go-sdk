// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package transitrouter

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opModifyTransitRouterBandwidthPackageAttributesCommon = "ModifyTransitRouterBandwidthPackageAttributes"

// ModifyTransitRouterBandwidthPackageAttributesCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ModifyTransitRouterBandwidthPackageAttributesCommon operation. The "output" return
// value will be populated with the ModifyTransitRouterBandwidthPackageAttributesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyTransitRouterBandwidthPackageAttributesCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyTransitRouterBandwidthPackageAttributesCommon Send returns without error.
//
// See ModifyTransitRouterBandwidthPackageAttributesCommon for more information on using the ModifyTransitRouterBandwidthPackageAttributesCommon
// API call, and error handling.
//
//    // Example sending a request using the ModifyTransitRouterBandwidthPackageAttributesCommonRequest method.
//    req, resp := client.ModifyTransitRouterBandwidthPackageAttributesCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *TRANSITROUTER) ModifyTransitRouterBandwidthPackageAttributesCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opModifyTransitRouterBandwidthPackageAttributesCommon,
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

// ModifyTransitRouterBandwidthPackageAttributesCommon API operation for TRANSITROUTER.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for TRANSITROUTER's
// API operation ModifyTransitRouterBandwidthPackageAttributesCommon for usage and error information.
func (c *TRANSITROUTER) ModifyTransitRouterBandwidthPackageAttributesCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ModifyTransitRouterBandwidthPackageAttributesCommonRequest(input)
	return out, req.Send()
}

// ModifyTransitRouterBandwidthPackageAttributesCommonWithContext is the same as ModifyTransitRouterBandwidthPackageAttributesCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyTransitRouterBandwidthPackageAttributesCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *TRANSITROUTER) ModifyTransitRouterBandwidthPackageAttributesCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ModifyTransitRouterBandwidthPackageAttributesCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opModifyTransitRouterBandwidthPackageAttributes = "ModifyTransitRouterBandwidthPackageAttributes"

// ModifyTransitRouterBandwidthPackageAttributesRequest generates a "volcengine/request.Request" representing the
// client's request for the ModifyTransitRouterBandwidthPackageAttributes operation. The "output" return
// value will be populated with the ModifyTransitRouterBandwidthPackageAttributesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyTransitRouterBandwidthPackageAttributesCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyTransitRouterBandwidthPackageAttributesCommon Send returns without error.
//
// See ModifyTransitRouterBandwidthPackageAttributes for more information on using the ModifyTransitRouterBandwidthPackageAttributes
// API call, and error handling.
//
//    // Example sending a request using the ModifyTransitRouterBandwidthPackageAttributesRequest method.
//    req, resp := client.ModifyTransitRouterBandwidthPackageAttributesRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *TRANSITROUTER) ModifyTransitRouterBandwidthPackageAttributesRequest(input *ModifyTransitRouterBandwidthPackageAttributesInput) (req *request.Request, output *ModifyTransitRouterBandwidthPackageAttributesOutput) {
	op := &request.Operation{
		Name:       opModifyTransitRouterBandwidthPackageAttributes,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ModifyTransitRouterBandwidthPackageAttributesInput{}
	}

	output = &ModifyTransitRouterBandwidthPackageAttributesOutput{}
	req = c.newRequest(op, input, output)

	return
}

// ModifyTransitRouterBandwidthPackageAttributes API operation for TRANSITROUTER.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for TRANSITROUTER's
// API operation ModifyTransitRouterBandwidthPackageAttributes for usage and error information.
func (c *TRANSITROUTER) ModifyTransitRouterBandwidthPackageAttributes(input *ModifyTransitRouterBandwidthPackageAttributesInput) (*ModifyTransitRouterBandwidthPackageAttributesOutput, error) {
	req, out := c.ModifyTransitRouterBandwidthPackageAttributesRequest(input)
	return out, req.Send()
}

// ModifyTransitRouterBandwidthPackageAttributesWithContext is the same as ModifyTransitRouterBandwidthPackageAttributes with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyTransitRouterBandwidthPackageAttributes for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *TRANSITROUTER) ModifyTransitRouterBandwidthPackageAttributesWithContext(ctx volcengine.Context, input *ModifyTransitRouterBandwidthPackageAttributesInput, opts ...request.Option) (*ModifyTransitRouterBandwidthPackageAttributesOutput, error) {
	req, out := c.ModifyTransitRouterBandwidthPackageAttributesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ModifyTransitRouterBandwidthPackageAttributesInput struct {
	_ struct{} `type:"structure"`

	Bandwidth *int32 `type:"int32"`

	Description *string `type:"string"`

	// TransitRouterBandwidthPackageId is a required field
	TransitRouterBandwidthPackageId *string `type:"string" required:"true"`

	TransitRouterBandwidthPackageName *string `type:"string"`
}

// String returns the string representation
func (s ModifyTransitRouterBandwidthPackageAttributesInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyTransitRouterBandwidthPackageAttributesInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ModifyTransitRouterBandwidthPackageAttributesInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ModifyTransitRouterBandwidthPackageAttributesInput"}
	if s.TransitRouterBandwidthPackageId == nil {
		invalidParams.Add(request.NewErrParamRequired("TransitRouterBandwidthPackageId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetBandwidth sets the Bandwidth field's value.
func (s *ModifyTransitRouterBandwidthPackageAttributesInput) SetBandwidth(v int32) *ModifyTransitRouterBandwidthPackageAttributesInput {
	s.Bandwidth = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *ModifyTransitRouterBandwidthPackageAttributesInput) SetDescription(v string) *ModifyTransitRouterBandwidthPackageAttributesInput {
	s.Description = &v
	return s
}

// SetTransitRouterBandwidthPackageId sets the TransitRouterBandwidthPackageId field's value.
func (s *ModifyTransitRouterBandwidthPackageAttributesInput) SetTransitRouterBandwidthPackageId(v string) *ModifyTransitRouterBandwidthPackageAttributesInput {
	s.TransitRouterBandwidthPackageId = &v
	return s
}

// SetTransitRouterBandwidthPackageName sets the TransitRouterBandwidthPackageName field's value.
func (s *ModifyTransitRouterBandwidthPackageAttributesInput) SetTransitRouterBandwidthPackageName(v string) *ModifyTransitRouterBandwidthPackageAttributesInput {
	s.TransitRouterBandwidthPackageName = &v
	return s
}

type ModifyTransitRouterBandwidthPackageAttributesOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s ModifyTransitRouterBandwidthPackageAttributesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyTransitRouterBandwidthPackageAttributesOutput) GoString() string {
	return s.String()
}

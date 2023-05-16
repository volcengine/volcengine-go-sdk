// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package vpc

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opModifyRouteEntryCommon = "ModifyRouteEntry"

// ModifyRouteEntryCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ModifyRouteEntryCommon operation. The "output" return
// value will be populated with the ModifyRouteEntryCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyRouteEntryCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyRouteEntryCommon Send returns without error.
//
// See ModifyRouteEntryCommon for more information on using the ModifyRouteEntryCommon
// API call, and error handling.
//
//	// Example sending a request using the ModifyRouteEntryCommonRequest method.
//	req, resp := client.ModifyRouteEntryCommonRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *VPC) ModifyRouteEntryCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opModifyRouteEntryCommon,
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

// ModifyRouteEntryCommon API operation for VPC.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VPC's
// API operation ModifyRouteEntryCommon for usage and error information.
func (c *VPC) ModifyRouteEntryCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ModifyRouteEntryCommonRequest(input)
	return out, req.Send()
}

// ModifyRouteEntryCommonWithContext is the same as ModifyRouteEntryCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyRouteEntryCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPC) ModifyRouteEntryCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ModifyRouteEntryCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opModifyRouteEntry = "ModifyRouteEntry"

// ModifyRouteEntryRequest generates a "volcengine/request.Request" representing the
// client's request for the ModifyRouteEntry operation. The "output" return
// value will be populated with the ModifyRouteEntryCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyRouteEntryCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyRouteEntryCommon Send returns without error.
//
// See ModifyRouteEntry for more information on using the ModifyRouteEntry
// API call, and error handling.
//
//	// Example sending a request using the ModifyRouteEntryRequest method.
//	req, resp := client.ModifyRouteEntryRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *VPC) ModifyRouteEntryRequest(input *ModifyRouteEntryInput) (req *request.Request, output *ModifyRouteEntryOutput) {
	op := &request.Operation{
		Name:       opModifyRouteEntry,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ModifyRouteEntryInput{}
	}

	output = &ModifyRouteEntryOutput{}
	req = c.newRequest(op, input, output)

	return
}

// ModifyRouteEntry API operation for VPC.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VPC's
// API operation ModifyRouteEntry for usage and error information.
func (c *VPC) ModifyRouteEntry(input *ModifyRouteEntryInput) (*ModifyRouteEntryOutput, error) {
	req, out := c.ModifyRouteEntryRequest(input)
	return out, req.Send()
}

// ModifyRouteEntryWithContext is the same as ModifyRouteEntry with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyRouteEntry for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPC) ModifyRouteEntryWithContext(ctx volcengine.Context, input *ModifyRouteEntryInput, opts ...request.Option) (*ModifyRouteEntryOutput, error) {
	req, out := c.ModifyRouteEntryRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ModifyRouteEntryInput struct {
	_ struct{} `type:"structure"`

	Description *string `min:"1" max:"255" type:"string"`

	// RouteEntryId is a required field
	RouteEntryId *string `type:"string" required:"true"`

	RouteEntryName *string `min:"1" max:"128" type:"string"`
}

// String returns the string representation
func (s ModifyRouteEntryInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyRouteEntryInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ModifyRouteEntryInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ModifyRouteEntryInput"}
	if s.Description != nil && len(*s.Description) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("Description", 1))
	}
	if s.Description != nil && len(*s.Description) > 255 {
		invalidParams.Add(request.NewErrParamMaxLen("Description", 255, *s.Description))
	}
	if s.RouteEntryId == nil {
		invalidParams.Add(request.NewErrParamRequired("RouteEntryId"))
	}
	if s.RouteEntryName != nil && len(*s.RouteEntryName) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("RouteEntryName", 1))
	}
	if s.RouteEntryName != nil && len(*s.RouteEntryName) > 128 {
		invalidParams.Add(request.NewErrParamMaxLen("RouteEntryName", 128, *s.RouteEntryName))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetDescription sets the Description field's value.
func (s *ModifyRouteEntryInput) SetDescription(v string) *ModifyRouteEntryInput {
	s.Description = &v
	return s
}

// SetRouteEntryId sets the RouteEntryId field's value.
func (s *ModifyRouteEntryInput) SetRouteEntryId(v string) *ModifyRouteEntryInput {
	s.RouteEntryId = &v
	return s
}

// SetRouteEntryName sets the RouteEntryName field's value.
func (s *ModifyRouteEntryInput) SetRouteEntryName(v string) *ModifyRouteEntryInput {
	s.RouteEntryName = &v
	return s
}

type ModifyRouteEntryOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	RequestId *string `type:"string"`
}

// String returns the string representation
func (s ModifyRouteEntryOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyRouteEntryOutput) GoString() string {
	return s.String()
}

// SetRequestId sets the RequestId field's value.
func (s *ModifyRouteEntryOutput) SetRequestId(v string) *ModifyRouteEntryOutput {
	s.RequestId = &v
	return s
}

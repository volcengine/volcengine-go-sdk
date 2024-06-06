// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package vpc

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opModifyEipAddressAttributesCommon = "ModifyEipAddressAttributes"

// ModifyEipAddressAttributesCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ModifyEipAddressAttributesCommon operation. The "output" return
// value will be populated with the ModifyEipAddressAttributesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyEipAddressAttributesCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyEipAddressAttributesCommon Send returns without error.
//
// See ModifyEipAddressAttributesCommon for more information on using the ModifyEipAddressAttributesCommon
// API call, and error handling.
//
//    // Example sending a request using the ModifyEipAddressAttributesCommonRequest method.
//    req, resp := client.ModifyEipAddressAttributesCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VPC) ModifyEipAddressAttributesCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opModifyEipAddressAttributesCommon,
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

// ModifyEipAddressAttributesCommon API operation for VPC.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VPC's
// API operation ModifyEipAddressAttributesCommon for usage and error information.
func (c *VPC) ModifyEipAddressAttributesCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ModifyEipAddressAttributesCommonRequest(input)
	return out, req.Send()
}

// ModifyEipAddressAttributesCommonWithContext is the same as ModifyEipAddressAttributesCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyEipAddressAttributesCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPC) ModifyEipAddressAttributesCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ModifyEipAddressAttributesCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opModifyEipAddressAttributes = "ModifyEipAddressAttributes"

// ModifyEipAddressAttributesRequest generates a "volcengine/request.Request" representing the
// client's request for the ModifyEipAddressAttributes operation. The "output" return
// value will be populated with the ModifyEipAddressAttributesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyEipAddressAttributesCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyEipAddressAttributesCommon Send returns without error.
//
// See ModifyEipAddressAttributes for more information on using the ModifyEipAddressAttributes
// API call, and error handling.
//
//    // Example sending a request using the ModifyEipAddressAttributesRequest method.
//    req, resp := client.ModifyEipAddressAttributesRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VPC) ModifyEipAddressAttributesRequest(input *ModifyEipAddressAttributesInput) (req *request.Request, output *ModifyEipAddressAttributesOutput) {
	op := &request.Operation{
		Name:       opModifyEipAddressAttributes,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ModifyEipAddressAttributesInput{}
	}

	output = &ModifyEipAddressAttributesOutput{}
	req = c.newRequest(op, input, output)

	return
}

// ModifyEipAddressAttributes API operation for VPC.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VPC's
// API operation ModifyEipAddressAttributes for usage and error information.
func (c *VPC) ModifyEipAddressAttributes(input *ModifyEipAddressAttributesInput) (*ModifyEipAddressAttributesOutput, error) {
	req, out := c.ModifyEipAddressAttributesRequest(input)
	return out, req.Send()
}

// ModifyEipAddressAttributesWithContext is the same as ModifyEipAddressAttributes with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyEipAddressAttributes for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPC) ModifyEipAddressAttributesWithContext(ctx volcengine.Context, input *ModifyEipAddressAttributesInput, opts ...request.Option) (*ModifyEipAddressAttributesOutput, error) {
	req, out := c.ModifyEipAddressAttributesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ModifyEipAddressAttributesInput struct {
	_ struct{} `type:"structure"`

	// AllocationId is a required field
	AllocationId *string `type:"string" required:"true"`

	Bandwidth *int64 `min:"1" max:"1000" type:"integer"`

	Description *string `min:"1" max:"255" type:"string"`

	Name *string `min:"1" max:"128" type:"string"`

	ReleaseWithInstance *bool `type:"boolean"`
}

// String returns the string representation
func (s ModifyEipAddressAttributesInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyEipAddressAttributesInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ModifyEipAddressAttributesInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ModifyEipAddressAttributesInput"}
	if s.AllocationId == nil {
		invalidParams.Add(request.NewErrParamRequired("AllocationId"))
	}
	if s.Bandwidth != nil && *s.Bandwidth < 1 {
		invalidParams.Add(request.NewErrParamMinValue("Bandwidth", 1))
	}
	if s.Bandwidth != nil && *s.Bandwidth > 1000 {
		invalidParams.Add(request.NewErrParamMaxValue("Bandwidth", 1000))
	}
	if s.Description != nil && len(*s.Description) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("Description", 1))
	}
	if s.Description != nil && len(*s.Description) > 255 {
		invalidParams.Add(request.NewErrParamMaxLen("Description", 255, *s.Description))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("Name", 1))
	}
	if s.Name != nil && len(*s.Name) > 128 {
		invalidParams.Add(request.NewErrParamMaxLen("Name", 128, *s.Name))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetAllocationId sets the AllocationId field's value.
func (s *ModifyEipAddressAttributesInput) SetAllocationId(v string) *ModifyEipAddressAttributesInput {
	s.AllocationId = &v
	return s
}

// SetBandwidth sets the Bandwidth field's value.
func (s *ModifyEipAddressAttributesInput) SetBandwidth(v int64) *ModifyEipAddressAttributesInput {
	s.Bandwidth = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *ModifyEipAddressAttributesInput) SetDescription(v string) *ModifyEipAddressAttributesInput {
	s.Description = &v
	return s
}

// SetName sets the Name field's value.
func (s *ModifyEipAddressAttributesInput) SetName(v string) *ModifyEipAddressAttributesInput {
	s.Name = &v
	return s
}

// SetReleaseWithInstance sets the ReleaseWithInstance field's value.
func (s *ModifyEipAddressAttributesInput) SetReleaseWithInstance(v bool) *ModifyEipAddressAttributesInput {
	s.ReleaseWithInstance = &v
	return s
}

type ModifyEipAddressAttributesOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	RequestId *string `type:"string"`
}

// String returns the string representation
func (s ModifyEipAddressAttributesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyEipAddressAttributesOutput) GoString() string {
	return s.String()
}

// SetRequestId sets the RequestId field's value.
func (s *ModifyEipAddressAttributesOutput) SetRequestId(v string) *ModifyEipAddressAttributesOutput {
	s.RequestId = &v
	return s
}

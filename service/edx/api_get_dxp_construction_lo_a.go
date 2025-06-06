// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package edx

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opGetDXPConstructionLOACommon = "GetDXPConstructionLOA"

// GetDXPConstructionLOACommonRequest generates a "volcengine/request.Request" representing the
// client's request for the GetDXPConstructionLOACommon operation. The "output" return
// value will be populated with the GetDXPConstructionLOACommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned GetDXPConstructionLOACommon Request to send the API call to the service.
// the "output" return value is not valid until after GetDXPConstructionLOACommon Send returns without error.
//
// See GetDXPConstructionLOACommon for more information on using the GetDXPConstructionLOACommon
// API call, and error handling.
//
//    // Example sending a request using the GetDXPConstructionLOACommonRequest method.
//    req, resp := client.GetDXPConstructionLOACommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *EDX) GetDXPConstructionLOACommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opGetDXPConstructionLOACommon,
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

// GetDXPConstructionLOACommon API operation for EDX.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for EDX's
// API operation GetDXPConstructionLOACommon for usage and error information.
func (c *EDX) GetDXPConstructionLOACommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.GetDXPConstructionLOACommonRequest(input)
	return out, req.Send()
}

// GetDXPConstructionLOACommonWithContext is the same as GetDXPConstructionLOACommon with the addition of
// the ability to pass a context and additional request options.
//
// See GetDXPConstructionLOACommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *EDX) GetDXPConstructionLOACommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.GetDXPConstructionLOACommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opGetDXPConstructionLOA = "GetDXPConstructionLOA"

// GetDXPConstructionLOARequest generates a "volcengine/request.Request" representing the
// client's request for the GetDXPConstructionLOA operation. The "output" return
// value will be populated with the GetDXPConstructionLOACommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned GetDXPConstructionLOACommon Request to send the API call to the service.
// the "output" return value is not valid until after GetDXPConstructionLOACommon Send returns without error.
//
// See GetDXPConstructionLOA for more information on using the GetDXPConstructionLOA
// API call, and error handling.
//
//    // Example sending a request using the GetDXPConstructionLOARequest method.
//    req, resp := client.GetDXPConstructionLOARequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *EDX) GetDXPConstructionLOARequest(input *GetDXPConstructionLOAInput) (req *request.Request, output *GetDXPConstructionLOAOutput) {
	op := &request.Operation{
		Name:       opGetDXPConstructionLOA,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetDXPConstructionLOAInput{}
	}

	output = &GetDXPConstructionLOAOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// GetDXPConstructionLOA API operation for EDX.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for EDX's
// API operation GetDXPConstructionLOA for usage and error information.
func (c *EDX) GetDXPConstructionLOA(input *GetDXPConstructionLOAInput) (*GetDXPConstructionLOAOutput, error) {
	req, out := c.GetDXPConstructionLOARequest(input)
	return out, req.Send()
}

// GetDXPConstructionLOAWithContext is the same as GetDXPConstructionLOA with the addition of
// the ability to pass a context and additional request options.
//
// See GetDXPConstructionLOA for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *EDX) GetDXPConstructionLOAWithContext(ctx volcengine.Context, input *GetDXPConstructionLOAInput, opts ...request.Option) (*GetDXPConstructionLOAOutput, error) {
	req, out := c.GetDXPConstructionLOARequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type GetDXPConstructionLOAInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	// InstanceId is a required field
	InstanceId *string `type:"string" json:",omitempty" required:"true"`
}

// String returns the string representation
func (s GetDXPConstructionLOAInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s GetDXPConstructionLOAInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetDXPConstructionLOAInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "GetDXPConstructionLOAInput"}
	if s.InstanceId == nil {
		invalidParams.Add(request.NewErrParamRequired("InstanceId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetInstanceId sets the InstanceId field's value.
func (s *GetDXPConstructionLOAInput) SetInstanceId(v string) *GetDXPConstructionLOAInput {
	s.InstanceId = &v
	return s
}

type GetDXPConstructionLOAOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	LOADATA *string `type:"string" json:",omitempty"`

	LOAName *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s GetDXPConstructionLOAOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s GetDXPConstructionLOAOutput) GoString() string {
	return s.String()
}

// SetLOADATA sets the LOADATA field's value.
func (s *GetDXPConstructionLOAOutput) SetLOADATA(v string) *GetDXPConstructionLOAOutput {
	s.LOADATA = &v
	return s
}

// SetLOAName sets the LOAName field's value.
func (s *GetDXPConstructionLOAOutput) SetLOAName(v string) *GetDXPConstructionLOAOutput {
	s.LOAName = &v
	return s
}

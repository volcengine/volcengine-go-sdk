// Code generated by volcstack with private/model/cli/gen-api/main.go. DO NOT EDIT.

package directconnect

import (
	"github.com/volcengine/volcstack-go-sdk/volcstack"
	"github.com/volcengine/volcstack-go-sdk/volcstack/request"
	"github.com/volcengine/volcstack-go-sdk/volcstack/response"
	"github.com/volcengine/volcstack-go-sdk/volcstack/volcstackutil"
)

const opModifyDirectConnectConnectionAttributesCommon = "ModifyDirectConnectConnectionAttributes"

// ModifyDirectConnectConnectionAttributesCommonRequest generates a "volcstack/request.Request" representing the
// client's request for the ModifyDirectConnectConnectionAttributesCommon operation. The "output" return
// value will be populated with the ModifyDirectConnectConnectionAttributesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyDirectConnectConnectionAttributesCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyDirectConnectConnectionAttributesCommon Send returns without error.
//
// See ModifyDirectConnectConnectionAttributesCommon for more information on using the ModifyDirectConnectConnectionAttributesCommon
// API call, and error handling.
//
//    // Example sending a request using the ModifyDirectConnectConnectionAttributesCommonRequest method.
//    req, resp := client.ModifyDirectConnectConnectionAttributesCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *DIRECTCONNECT) ModifyDirectConnectConnectionAttributesCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opModifyDirectConnectConnectionAttributesCommon,
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

// ModifyDirectConnectConnectionAttributesCommon API operation for DIRECTCONNECT.
//
// Returns volcstackerr.Error for service API and SDK errors. Use runtime type assertions
// with volcstackerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCSTACK API reference guide for DIRECTCONNECT's
// API operation ModifyDirectConnectConnectionAttributesCommon for usage and error information.
func (c *DIRECTCONNECT) ModifyDirectConnectConnectionAttributesCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ModifyDirectConnectConnectionAttributesCommonRequest(input)
	return out, req.Send()
}

// ModifyDirectConnectConnectionAttributesCommonWithContext is the same as ModifyDirectConnectConnectionAttributesCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyDirectConnectConnectionAttributesCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *DIRECTCONNECT) ModifyDirectConnectConnectionAttributesCommonWithContext(ctx volcstack.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ModifyDirectConnectConnectionAttributesCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opModifyDirectConnectConnectionAttributes = "ModifyDirectConnectConnectionAttributes"

// ModifyDirectConnectConnectionAttributesRequest generates a "volcstack/request.Request" representing the
// client's request for the ModifyDirectConnectConnectionAttributes operation. The "output" return
// value will be populated with the ModifyDirectConnectConnectionAttributesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyDirectConnectConnectionAttributesCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyDirectConnectConnectionAttributesCommon Send returns without error.
//
// See ModifyDirectConnectConnectionAttributes for more information on using the ModifyDirectConnectConnectionAttributes
// API call, and error handling.
//
//    // Example sending a request using the ModifyDirectConnectConnectionAttributesRequest method.
//    req, resp := client.ModifyDirectConnectConnectionAttributesRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *DIRECTCONNECT) ModifyDirectConnectConnectionAttributesRequest(input *ModifyDirectConnectConnectionAttributesInput) (req *request.Request, output *ModifyDirectConnectConnectionAttributesOutput) {
	op := &request.Operation{
		Name:       opModifyDirectConnectConnectionAttributes,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ModifyDirectConnectConnectionAttributesInput{}
	}

	output = &ModifyDirectConnectConnectionAttributesOutput{}
	req = c.newRequest(op, input, output)

	return
}

// ModifyDirectConnectConnectionAttributes API operation for DIRECTCONNECT.
//
// Returns volcstackerr.Error for service API and SDK errors. Use runtime type assertions
// with volcstackerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCSTACK API reference guide for DIRECTCONNECT's
// API operation ModifyDirectConnectConnectionAttributes for usage and error information.
func (c *DIRECTCONNECT) ModifyDirectConnectConnectionAttributes(input *ModifyDirectConnectConnectionAttributesInput) (*ModifyDirectConnectConnectionAttributesOutput, error) {
	req, out := c.ModifyDirectConnectConnectionAttributesRequest(input)
	return out, req.Send()
}

// ModifyDirectConnectConnectionAttributesWithContext is the same as ModifyDirectConnectConnectionAttributes with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyDirectConnectConnectionAttributes for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *DIRECTCONNECT) ModifyDirectConnectConnectionAttributesWithContext(ctx volcstack.Context, input *ModifyDirectConnectConnectionAttributesInput, opts ...request.Option) (*ModifyDirectConnectConnectionAttributesOutput, error) {
	req, out := c.ModifyDirectConnectConnectionAttributesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ModifyDirectConnectConnectionAttributesInput struct {
	_ struct{} `type:"structure"`

	Description *string `type:"string"`

	// DirectConnectConnectionId is a required field
	DirectConnectConnectionId *string `type:"string" required:"true"`

	DirectConnectConnectionName *string `type:"string"`
}

// String returns the string representation
func (s ModifyDirectConnectConnectionAttributesInput) String() string {
	return volcstackutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyDirectConnectConnectionAttributesInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ModifyDirectConnectConnectionAttributesInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ModifyDirectConnectConnectionAttributesInput"}
	if s.DirectConnectConnectionId == nil {
		invalidParams.Add(request.NewErrParamRequired("DirectConnectConnectionId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetDescription sets the Description field's value.
func (s *ModifyDirectConnectConnectionAttributesInput) SetDescription(v string) *ModifyDirectConnectConnectionAttributesInput {
	s.Description = &v
	return s
}

// SetDirectConnectConnectionId sets the DirectConnectConnectionId field's value.
func (s *ModifyDirectConnectConnectionAttributesInput) SetDirectConnectConnectionId(v string) *ModifyDirectConnectConnectionAttributesInput {
	s.DirectConnectConnectionId = &v
	return s
}

// SetDirectConnectConnectionName sets the DirectConnectConnectionName field's value.
func (s *ModifyDirectConnectConnectionAttributesInput) SetDirectConnectConnectionName(v string) *ModifyDirectConnectConnectionAttributesInput {
	s.DirectConnectConnectionName = &v
	return s
}

type ModifyDirectConnectConnectionAttributesOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	RequestId *string `type:"string"`
}

// String returns the string representation
func (s ModifyDirectConnectConnectionAttributesOutput) String() string {
	return volcstackutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyDirectConnectConnectionAttributesOutput) GoString() string {
	return s.String()
}

// SetRequestId sets the RequestId field's value.
func (s *ModifyDirectConnectConnectionAttributesOutput) SetRequestId(v string) *ModifyDirectConnectConnectionAttributesOutput {
	s.RequestId = &v
	return s
}
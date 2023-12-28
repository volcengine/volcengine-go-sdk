// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package directconnect

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opModifyBgpPeerAttributesCommon = "ModifyBgpPeerAttributes"

// ModifyBgpPeerAttributesCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ModifyBgpPeerAttributesCommon operation. The "output" return
// value will be populated with the ModifyBgpPeerAttributesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyBgpPeerAttributesCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyBgpPeerAttributesCommon Send returns without error.
//
// See ModifyBgpPeerAttributesCommon for more information on using the ModifyBgpPeerAttributesCommon
// API call, and error handling.
//
//    // Example sending a request using the ModifyBgpPeerAttributesCommonRequest method.
//    req, resp := client.ModifyBgpPeerAttributesCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *DIRECTCONNECT) ModifyBgpPeerAttributesCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opModifyBgpPeerAttributesCommon,
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

// ModifyBgpPeerAttributesCommon API operation for DIRECTCONNECT.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for DIRECTCONNECT's
// API operation ModifyBgpPeerAttributesCommon for usage and error information.
func (c *DIRECTCONNECT) ModifyBgpPeerAttributesCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ModifyBgpPeerAttributesCommonRequest(input)
	return out, req.Send()
}

// ModifyBgpPeerAttributesCommonWithContext is the same as ModifyBgpPeerAttributesCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyBgpPeerAttributesCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *DIRECTCONNECT) ModifyBgpPeerAttributesCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ModifyBgpPeerAttributesCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opModifyBgpPeerAttributes = "ModifyBgpPeerAttributes"

// ModifyBgpPeerAttributesRequest generates a "volcengine/request.Request" representing the
// client's request for the ModifyBgpPeerAttributes operation. The "output" return
// value will be populated with the ModifyBgpPeerAttributesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyBgpPeerAttributesCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyBgpPeerAttributesCommon Send returns without error.
//
// See ModifyBgpPeerAttributes for more information on using the ModifyBgpPeerAttributes
// API call, and error handling.
//
//    // Example sending a request using the ModifyBgpPeerAttributesRequest method.
//    req, resp := client.ModifyBgpPeerAttributesRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *DIRECTCONNECT) ModifyBgpPeerAttributesRequest(input *ModifyBgpPeerAttributesInput) (req *request.Request, output *ModifyBgpPeerAttributesOutput) {
	op := &request.Operation{
		Name:       opModifyBgpPeerAttributes,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ModifyBgpPeerAttributesInput{}
	}

	output = &ModifyBgpPeerAttributesOutput{}
	req = c.newRequest(op, input, output)

	return
}

// ModifyBgpPeerAttributes API operation for DIRECTCONNECT.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for DIRECTCONNECT's
// API operation ModifyBgpPeerAttributes for usage and error information.
func (c *DIRECTCONNECT) ModifyBgpPeerAttributes(input *ModifyBgpPeerAttributesInput) (*ModifyBgpPeerAttributesOutput, error) {
	req, out := c.ModifyBgpPeerAttributesRequest(input)
	return out, req.Send()
}

// ModifyBgpPeerAttributesWithContext is the same as ModifyBgpPeerAttributes with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyBgpPeerAttributes for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *DIRECTCONNECT) ModifyBgpPeerAttributesWithContext(ctx volcengine.Context, input *ModifyBgpPeerAttributesInput, opts ...request.Option) (*ModifyBgpPeerAttributesOutput, error) {
	req, out := c.ModifyBgpPeerAttributesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ModifyBgpPeerAttributesInput struct {
	_ struct{} `type:"structure"`

	// BgpPeerId is a required field
	BgpPeerId *string `type:"string" required:"true"`

	BgpPeerName *string `type:"string"`

	Description *string `type:"string"`
}

// String returns the string representation
func (s ModifyBgpPeerAttributesInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyBgpPeerAttributesInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ModifyBgpPeerAttributesInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ModifyBgpPeerAttributesInput"}
	if s.BgpPeerId == nil {
		invalidParams.Add(request.NewErrParamRequired("BgpPeerId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetBgpPeerId sets the BgpPeerId field's value.
func (s *ModifyBgpPeerAttributesInput) SetBgpPeerId(v string) *ModifyBgpPeerAttributesInput {
	s.BgpPeerId = &v
	return s
}

// SetBgpPeerName sets the BgpPeerName field's value.
func (s *ModifyBgpPeerAttributesInput) SetBgpPeerName(v string) *ModifyBgpPeerAttributesInput {
	s.BgpPeerName = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *ModifyBgpPeerAttributesInput) SetDescription(v string) *ModifyBgpPeerAttributesInput {
	s.Description = &v
	return s
}

type ModifyBgpPeerAttributesOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	RequestId *string `type:"string"`
}

// String returns the string representation
func (s ModifyBgpPeerAttributesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyBgpPeerAttributesOutput) GoString() string {
	return s.String()
}

// SetRequestId sets the RequestId field's value.
func (s *ModifyBgpPeerAttributesOutput) SetRequestId(v string) *ModifyBgpPeerAttributesOutput {
	s.RequestId = &v
	return s
}

// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package edx

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDeleteEDXPeerLinkCommon = "DeleteEDXPeerLink"

// DeleteEDXPeerLinkCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DeleteEDXPeerLinkCommon operation. The "output" return
// value will be populated with the DeleteEDXPeerLinkCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeleteEDXPeerLinkCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeleteEDXPeerLinkCommon Send returns without error.
//
// See DeleteEDXPeerLinkCommon for more information on using the DeleteEDXPeerLinkCommon
// API call, and error handling.
//
//    // Example sending a request using the DeleteEDXPeerLinkCommonRequest method.
//    req, resp := client.DeleteEDXPeerLinkCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *EDX) DeleteEDXPeerLinkCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDeleteEDXPeerLinkCommon,
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

// DeleteEDXPeerLinkCommon API operation for EDX.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for EDX's
// API operation DeleteEDXPeerLinkCommon for usage and error information.
func (c *EDX) DeleteEDXPeerLinkCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DeleteEDXPeerLinkCommonRequest(input)
	return out, req.Send()
}

// DeleteEDXPeerLinkCommonWithContext is the same as DeleteEDXPeerLinkCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteEDXPeerLinkCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *EDX) DeleteEDXPeerLinkCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DeleteEDXPeerLinkCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDeleteEDXPeerLink = "DeleteEDXPeerLink"

// DeleteEDXPeerLinkRequest generates a "volcengine/request.Request" representing the
// client's request for the DeleteEDXPeerLink operation. The "output" return
// value will be populated with the DeleteEDXPeerLinkCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeleteEDXPeerLinkCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeleteEDXPeerLinkCommon Send returns without error.
//
// See DeleteEDXPeerLink for more information on using the DeleteEDXPeerLink
// API call, and error handling.
//
//    // Example sending a request using the DeleteEDXPeerLinkRequest method.
//    req, resp := client.DeleteEDXPeerLinkRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *EDX) DeleteEDXPeerLinkRequest(input *DeleteEDXPeerLinkInput) (req *request.Request, output *DeleteEDXPeerLinkOutput) {
	op := &request.Operation{
		Name:       opDeleteEDXPeerLink,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteEDXPeerLinkInput{}
	}

	output = &DeleteEDXPeerLinkOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// DeleteEDXPeerLink API operation for EDX.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for EDX's
// API operation DeleteEDXPeerLink for usage and error information.
func (c *EDX) DeleteEDXPeerLink(input *DeleteEDXPeerLinkInput) (*DeleteEDXPeerLinkOutput, error) {
	req, out := c.DeleteEDXPeerLinkRequest(input)
	return out, req.Send()
}

// DeleteEDXPeerLinkWithContext is the same as DeleteEDXPeerLink with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteEDXPeerLink for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *EDX) DeleteEDXPeerLinkWithContext(ctx volcengine.Context, input *DeleteEDXPeerLinkInput, opts ...request.Option) (*DeleteEDXPeerLinkOutput, error) {
	req, out := c.DeleteEDXPeerLinkRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DeleteEDXPeerLinkInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	// PeerLinkId is a required field
	PeerLinkId *string `type:"string" json:",omitempty" required:"true"`
}

// String returns the string representation
func (s DeleteEDXPeerLinkInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteEDXPeerLinkInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteEDXPeerLinkInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DeleteEDXPeerLinkInput"}
	if s.PeerLinkId == nil {
		invalidParams.Add(request.NewErrParamRequired("PeerLinkId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetPeerLinkId sets the PeerLinkId field's value.
func (s *DeleteEDXPeerLinkInput) SetPeerLinkId(v string) *DeleteEDXPeerLinkInput {
	s.PeerLinkId = &v
	return s
}

type DeleteEDXPeerLinkOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s DeleteEDXPeerLinkOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteEDXPeerLinkOutput) GoString() string {
	return s.String()
}

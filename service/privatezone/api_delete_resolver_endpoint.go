// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package privatezone

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDeleteResolverEndpointCommon = "DeleteResolverEndpoint"

// DeleteResolverEndpointCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DeleteResolverEndpointCommon operation. The "output" return
// value will be populated with the DeleteResolverEndpointCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeleteResolverEndpointCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeleteResolverEndpointCommon Send returns without error.
//
// See DeleteResolverEndpointCommon for more information on using the DeleteResolverEndpointCommon
// API call, and error handling.
//
//    // Example sending a request using the DeleteResolverEndpointCommonRequest method.
//    req, resp := client.DeleteResolverEndpointCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *PRIVATEZONE) DeleteResolverEndpointCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDeleteResolverEndpointCommon,
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

// DeleteResolverEndpointCommon API operation for PRIVATE_ZONE.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for PRIVATE_ZONE's
// API operation DeleteResolverEndpointCommon for usage and error information.
func (c *PRIVATEZONE) DeleteResolverEndpointCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DeleteResolverEndpointCommonRequest(input)
	return out, req.Send()
}

// DeleteResolverEndpointCommonWithContext is the same as DeleteResolverEndpointCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteResolverEndpointCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *PRIVATEZONE) DeleteResolverEndpointCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DeleteResolverEndpointCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDeleteResolverEndpoint = "DeleteResolverEndpoint"

// DeleteResolverEndpointRequest generates a "volcengine/request.Request" representing the
// client's request for the DeleteResolverEndpoint operation. The "output" return
// value will be populated with the DeleteResolverEndpointCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeleteResolverEndpointCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeleteResolverEndpointCommon Send returns without error.
//
// See DeleteResolverEndpoint for more information on using the DeleteResolverEndpoint
// API call, and error handling.
//
//    // Example sending a request using the DeleteResolverEndpointRequest method.
//    req, resp := client.DeleteResolverEndpointRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *PRIVATEZONE) DeleteResolverEndpointRequest(input *DeleteResolverEndpointInput) (req *request.Request, output *DeleteResolverEndpointOutput) {
	op := &request.Operation{
		Name:       opDeleteResolverEndpoint,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteResolverEndpointInput{}
	}

	output = &DeleteResolverEndpointOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// DeleteResolverEndpoint API operation for PRIVATE_ZONE.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for PRIVATE_ZONE's
// API operation DeleteResolverEndpoint for usage and error information.
func (c *PRIVATEZONE) DeleteResolverEndpoint(input *DeleteResolverEndpointInput) (*DeleteResolverEndpointOutput, error) {
	req, out := c.DeleteResolverEndpointRequest(input)
	return out, req.Send()
}

// DeleteResolverEndpointWithContext is the same as DeleteResolverEndpoint with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteResolverEndpoint for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *PRIVATEZONE) DeleteResolverEndpointWithContext(ctx volcengine.Context, input *DeleteResolverEndpointInput, opts ...request.Option) (*DeleteResolverEndpointOutput, error) {
	req, out := c.DeleteResolverEndpointRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DeleteResolverEndpointInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	// EndpointID is a required field
	EndpointID *int64 `type:"int64" json:",omitempty" required:"true"`

	EndpointTrn *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s DeleteResolverEndpointInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteResolverEndpointInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteResolverEndpointInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DeleteResolverEndpointInput"}
	if s.EndpointID == nil {
		invalidParams.Add(request.NewErrParamRequired("EndpointID"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetEndpointID sets the EndpointID field's value.
func (s *DeleteResolverEndpointInput) SetEndpointID(v int64) *DeleteResolverEndpointInput {
	s.EndpointID = &v
	return s
}

// SetEndpointTrn sets the EndpointTrn field's value.
func (s *DeleteResolverEndpointInput) SetEndpointTrn(v string) *DeleteResolverEndpointInput {
	s.EndpointTrn = &v
	return s
}

type DeleteResolverEndpointOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s DeleteResolverEndpointOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteResolverEndpointOutput) GoString() string {
	return s.String()
}

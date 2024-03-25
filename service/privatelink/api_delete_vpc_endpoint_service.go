// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package privatelink

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDeleteVpcEndpointServiceCommon = "DeleteVpcEndpointService"

// DeleteVpcEndpointServiceCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DeleteVpcEndpointServiceCommon operation. The "output" return
// value will be populated with the DeleteVpcEndpointServiceCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeleteVpcEndpointServiceCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeleteVpcEndpointServiceCommon Send returns without error.
//
// See DeleteVpcEndpointServiceCommon for more information on using the DeleteVpcEndpointServiceCommon
// API call, and error handling.
//
//    // Example sending a request using the DeleteVpcEndpointServiceCommonRequest method.
//    req, resp := client.DeleteVpcEndpointServiceCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *PRIVATELINK) DeleteVpcEndpointServiceCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDeleteVpcEndpointServiceCommon,
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

// DeleteVpcEndpointServiceCommon API operation for PRIVATELINK.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for PRIVATELINK's
// API operation DeleteVpcEndpointServiceCommon for usage and error information.
func (c *PRIVATELINK) DeleteVpcEndpointServiceCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DeleteVpcEndpointServiceCommonRequest(input)
	return out, req.Send()
}

// DeleteVpcEndpointServiceCommonWithContext is the same as DeleteVpcEndpointServiceCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteVpcEndpointServiceCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *PRIVATELINK) DeleteVpcEndpointServiceCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DeleteVpcEndpointServiceCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDeleteVpcEndpointService = "DeleteVpcEndpointService"

// DeleteVpcEndpointServiceRequest generates a "volcengine/request.Request" representing the
// client's request for the DeleteVpcEndpointService operation. The "output" return
// value will be populated with the DeleteVpcEndpointServiceCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeleteVpcEndpointServiceCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeleteVpcEndpointServiceCommon Send returns without error.
//
// See DeleteVpcEndpointService for more information on using the DeleteVpcEndpointService
// API call, and error handling.
//
//    // Example sending a request using the DeleteVpcEndpointServiceRequest method.
//    req, resp := client.DeleteVpcEndpointServiceRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *PRIVATELINK) DeleteVpcEndpointServiceRequest(input *DeleteVpcEndpointServiceInput) (req *request.Request, output *DeleteVpcEndpointServiceOutput) {
	op := &request.Operation{
		Name:       opDeleteVpcEndpointService,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteVpcEndpointServiceInput{}
	}

	output = &DeleteVpcEndpointServiceOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DeleteVpcEndpointService API operation for PRIVATELINK.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for PRIVATELINK's
// API operation DeleteVpcEndpointService for usage and error information.
func (c *PRIVATELINK) DeleteVpcEndpointService(input *DeleteVpcEndpointServiceInput) (*DeleteVpcEndpointServiceOutput, error) {
	req, out := c.DeleteVpcEndpointServiceRequest(input)
	return out, req.Send()
}

// DeleteVpcEndpointServiceWithContext is the same as DeleteVpcEndpointService with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteVpcEndpointService for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *PRIVATELINK) DeleteVpcEndpointServiceWithContext(ctx volcengine.Context, input *DeleteVpcEndpointServiceInput, opts ...request.Option) (*DeleteVpcEndpointServiceOutput, error) {
	req, out := c.DeleteVpcEndpointServiceRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DeleteVpcEndpointServiceInput struct {
	_ struct{} `type:"structure"`

	// ServiceId is a required field
	ServiceId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteVpcEndpointServiceInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteVpcEndpointServiceInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteVpcEndpointServiceInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DeleteVpcEndpointServiceInput"}
	if s.ServiceId == nil {
		invalidParams.Add(request.NewErrParamRequired("ServiceId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetServiceId sets the ServiceId field's value.
func (s *DeleteVpcEndpointServiceInput) SetServiceId(v string) *DeleteVpcEndpointServiceInput {
	s.ServiceId = &v
	return s
}

type DeleteVpcEndpointServiceOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	RequestId *string `type:"string"`
}

// String returns the string representation
func (s DeleteVpcEndpointServiceOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteVpcEndpointServiceOutput) GoString() string {
	return s.String()
}

// SetRequestId sets the RequestId field's value.
func (s *DeleteVpcEndpointServiceOutput) SetRequestId(v string) *DeleteVpcEndpointServiceOutput {
	s.RequestId = &v
	return s
}

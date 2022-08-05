// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package redis

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDeleteDBEndpointPublicAddressCommon = "DeleteDBEndpointPublicAddress"

// DeleteDBEndpointPublicAddressCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DeleteDBEndpointPublicAddressCommon operation. The "output" return
// value will be populated with the DeleteDBEndpointPublicAddressCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeleteDBEndpointPublicAddressCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeleteDBEndpointPublicAddressCommon Send returns without error.
//
// See DeleteDBEndpointPublicAddressCommon for more information on using the DeleteDBEndpointPublicAddressCommon
// API call, and error handling.
//
//    // Example sending a request using the DeleteDBEndpointPublicAddressCommonRequest method.
//    req, resp := client.DeleteDBEndpointPublicAddressCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *REDIS) DeleteDBEndpointPublicAddressCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDeleteDBEndpointPublicAddressCommon,
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

// DeleteDBEndpointPublicAddressCommon API operation for REDIS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for REDIS's
// API operation DeleteDBEndpointPublicAddressCommon for usage and error information.
func (c *REDIS) DeleteDBEndpointPublicAddressCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DeleteDBEndpointPublicAddressCommonRequest(input)
	return out, req.Send()
}

// DeleteDBEndpointPublicAddressCommonWithContext is the same as DeleteDBEndpointPublicAddressCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteDBEndpointPublicAddressCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *REDIS) DeleteDBEndpointPublicAddressCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DeleteDBEndpointPublicAddressCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDeleteDBEndpointPublicAddress = "DeleteDBEndpointPublicAddress"

// DeleteDBEndpointPublicAddressRequest generates a "volcengine/request.Request" representing the
// client's request for the DeleteDBEndpointPublicAddress operation. The "output" return
// value will be populated with the DeleteDBEndpointPublicAddressCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeleteDBEndpointPublicAddressCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeleteDBEndpointPublicAddressCommon Send returns without error.
//
// See DeleteDBEndpointPublicAddress for more information on using the DeleteDBEndpointPublicAddress
// API call, and error handling.
//
//    // Example sending a request using the DeleteDBEndpointPublicAddressRequest method.
//    req, resp := client.DeleteDBEndpointPublicAddressRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *REDIS) DeleteDBEndpointPublicAddressRequest(input *DeleteDBEndpointPublicAddressInput) (req *request.Request, output *DeleteDBEndpointPublicAddressOutput) {
	op := &request.Operation{
		Name:       opDeleteDBEndpointPublicAddress,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteDBEndpointPublicAddressInput{}
	}

	output = &DeleteDBEndpointPublicAddressOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// DeleteDBEndpointPublicAddress API operation for REDIS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for REDIS's
// API operation DeleteDBEndpointPublicAddress for usage and error information.
func (c *REDIS) DeleteDBEndpointPublicAddress(input *DeleteDBEndpointPublicAddressInput) (*DeleteDBEndpointPublicAddressOutput, error) {
	req, out := c.DeleteDBEndpointPublicAddressRequest(input)
	return out, req.Send()
}

// DeleteDBEndpointPublicAddressWithContext is the same as DeleteDBEndpointPublicAddress with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteDBEndpointPublicAddress for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *REDIS) DeleteDBEndpointPublicAddressWithContext(ctx volcengine.Context, input *DeleteDBEndpointPublicAddressInput, opts ...request.Option) (*DeleteDBEndpointPublicAddressOutput, error) {
	req, out := c.DeleteDBEndpointPublicAddressRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DeleteDBEndpointPublicAddressInput struct {
	_ struct{} `type:"structure"`

	// EipId is a required field
	EipId *string `type:"string" required:"true"`

	// InstanceId is a required field
	InstanceId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteDBEndpointPublicAddressInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteDBEndpointPublicAddressInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteDBEndpointPublicAddressInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DeleteDBEndpointPublicAddressInput"}
	if s.EipId == nil {
		invalidParams.Add(request.NewErrParamRequired("EipId"))
	}
	if s.InstanceId == nil {
		invalidParams.Add(request.NewErrParamRequired("InstanceId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetEipId sets the EipId field's value.
func (s *DeleteDBEndpointPublicAddressInput) SetEipId(v string) *DeleteDBEndpointPublicAddressInput {
	s.EipId = &v
	return s
}

// SetInstanceId sets the InstanceId field's value.
func (s *DeleteDBEndpointPublicAddressInput) SetInstanceId(v string) *DeleteDBEndpointPublicAddressInput {
	s.InstanceId = &v
	return s
}

type DeleteDBEndpointPublicAddressOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s DeleteDBEndpointPublicAddressOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteDBEndpointPublicAddressOutput) GoString() string {
	return s.String()
}
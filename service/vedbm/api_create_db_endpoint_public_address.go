// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package vedbm

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opCreateDBEndpointPublicAddressCommon = "CreateDBEndpointPublicAddress"

// CreateDBEndpointPublicAddressCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateDBEndpointPublicAddressCommon operation. The "output" return
// value will be populated with the CreateDBEndpointPublicAddressCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateDBEndpointPublicAddressCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateDBEndpointPublicAddressCommon Send returns without error.
//
// See CreateDBEndpointPublicAddressCommon for more information on using the CreateDBEndpointPublicAddressCommon
// API call, and error handling.
//
//    // Example sending a request using the CreateDBEndpointPublicAddressCommonRequest method.
//    req, resp := client.CreateDBEndpointPublicAddressCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VEDBM) CreateDBEndpointPublicAddressCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opCreateDBEndpointPublicAddressCommon,
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

// CreateDBEndpointPublicAddressCommon API operation for VEDBM.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VEDBM's
// API operation CreateDBEndpointPublicAddressCommon for usage and error information.
func (c *VEDBM) CreateDBEndpointPublicAddressCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.CreateDBEndpointPublicAddressCommonRequest(input)
	return out, req.Send()
}

// CreateDBEndpointPublicAddressCommonWithContext is the same as CreateDBEndpointPublicAddressCommon with the addition of
// the ability to pass a context and additional request options.
//
// See CreateDBEndpointPublicAddressCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VEDBM) CreateDBEndpointPublicAddressCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.CreateDBEndpointPublicAddressCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opCreateDBEndpointPublicAddress = "CreateDBEndpointPublicAddress"

// CreateDBEndpointPublicAddressRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateDBEndpointPublicAddress operation. The "output" return
// value will be populated with the CreateDBEndpointPublicAddressCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateDBEndpointPublicAddressCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateDBEndpointPublicAddressCommon Send returns without error.
//
// See CreateDBEndpointPublicAddress for more information on using the CreateDBEndpointPublicAddress
// API call, and error handling.
//
//    // Example sending a request using the CreateDBEndpointPublicAddressRequest method.
//    req, resp := client.CreateDBEndpointPublicAddressRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VEDBM) CreateDBEndpointPublicAddressRequest(input *CreateDBEndpointPublicAddressInput) (req *request.Request, output *CreateDBEndpointPublicAddressOutput) {
	op := &request.Operation{
		Name:       opCreateDBEndpointPublicAddress,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateDBEndpointPublicAddressInput{}
	}

	output = &CreateDBEndpointPublicAddressOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// CreateDBEndpointPublicAddress API operation for VEDBM.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VEDBM's
// API operation CreateDBEndpointPublicAddress for usage and error information.
func (c *VEDBM) CreateDBEndpointPublicAddress(input *CreateDBEndpointPublicAddressInput) (*CreateDBEndpointPublicAddressOutput, error) {
	req, out := c.CreateDBEndpointPublicAddressRequest(input)
	return out, req.Send()
}

// CreateDBEndpointPublicAddressWithContext is the same as CreateDBEndpointPublicAddress with the addition of
// the ability to pass a context and additional request options.
//
// See CreateDBEndpointPublicAddress for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VEDBM) CreateDBEndpointPublicAddressWithContext(ctx volcengine.Context, input *CreateDBEndpointPublicAddressInput, opts ...request.Option) (*CreateDBEndpointPublicAddressOutput, error) {
	req, out := c.CreateDBEndpointPublicAddressRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type CreateDBEndpointPublicAddressInput struct {
	_ struct{} `type:"structure"`

	EipId *string `type:"string"`

	EndpointId *string `type:"string"`

	InstanceId *string `type:"string"`
}

// String returns the string representation
func (s CreateDBEndpointPublicAddressInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateDBEndpointPublicAddressInput) GoString() string {
	return s.String()
}

// SetEipId sets the EipId field's value.
func (s *CreateDBEndpointPublicAddressInput) SetEipId(v string) *CreateDBEndpointPublicAddressInput {
	s.EipId = &v
	return s
}

// SetEndpointId sets the EndpointId field's value.
func (s *CreateDBEndpointPublicAddressInput) SetEndpointId(v string) *CreateDBEndpointPublicAddressInput {
	s.EndpointId = &v
	return s
}

// SetInstanceId sets the InstanceId field's value.
func (s *CreateDBEndpointPublicAddressInput) SetInstanceId(v string) *CreateDBEndpointPublicAddressInput {
	s.InstanceId = &v
	return s
}

type CreateDBEndpointPublicAddressOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s CreateDBEndpointPublicAddressOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateDBEndpointPublicAddressOutput) GoString() string {
	return s.String()
}

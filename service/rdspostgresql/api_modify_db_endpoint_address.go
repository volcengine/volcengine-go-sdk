// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package rdspostgresql

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opModifyDBEndpointAddressCommon = "ModifyDBEndpointAddress"

// ModifyDBEndpointAddressCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ModifyDBEndpointAddressCommon operation. The "output" return
// value will be populated with the ModifyDBEndpointAddressCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyDBEndpointAddressCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyDBEndpointAddressCommon Send returns without error.
//
// See ModifyDBEndpointAddressCommon for more information on using the ModifyDBEndpointAddressCommon
// API call, and error handling.
//
//    // Example sending a request using the ModifyDBEndpointAddressCommonRequest method.
//    req, resp := client.ModifyDBEndpointAddressCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *RDSPOSTGRESQL) ModifyDBEndpointAddressCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opModifyDBEndpointAddressCommon,
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

// ModifyDBEndpointAddressCommon API operation for RDS_POSTGRESQL.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RDS_POSTGRESQL's
// API operation ModifyDBEndpointAddressCommon for usage and error information.
func (c *RDSPOSTGRESQL) ModifyDBEndpointAddressCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ModifyDBEndpointAddressCommonRequest(input)
	return out, req.Send()
}

// ModifyDBEndpointAddressCommonWithContext is the same as ModifyDBEndpointAddressCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyDBEndpointAddressCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RDSPOSTGRESQL) ModifyDBEndpointAddressCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ModifyDBEndpointAddressCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opModifyDBEndpointAddress = "ModifyDBEndpointAddress"

// ModifyDBEndpointAddressRequest generates a "volcengine/request.Request" representing the
// client's request for the ModifyDBEndpointAddress operation. The "output" return
// value will be populated with the ModifyDBEndpointAddressCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyDBEndpointAddressCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyDBEndpointAddressCommon Send returns without error.
//
// See ModifyDBEndpointAddress for more information on using the ModifyDBEndpointAddress
// API call, and error handling.
//
//    // Example sending a request using the ModifyDBEndpointAddressRequest method.
//    req, resp := client.ModifyDBEndpointAddressRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *RDSPOSTGRESQL) ModifyDBEndpointAddressRequest(input *ModifyDBEndpointAddressInput) (req *request.Request, output *ModifyDBEndpointAddressOutput) {
	op := &request.Operation{
		Name:       opModifyDBEndpointAddress,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ModifyDBEndpointAddressInput{}
	}

	output = &ModifyDBEndpointAddressOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// ModifyDBEndpointAddress API operation for RDS_POSTGRESQL.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RDS_POSTGRESQL's
// API operation ModifyDBEndpointAddress for usage and error information.
func (c *RDSPOSTGRESQL) ModifyDBEndpointAddress(input *ModifyDBEndpointAddressInput) (*ModifyDBEndpointAddressOutput, error) {
	req, out := c.ModifyDBEndpointAddressRequest(input)
	return out, req.Send()
}

// ModifyDBEndpointAddressWithContext is the same as ModifyDBEndpointAddress with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyDBEndpointAddress for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RDSPOSTGRESQL) ModifyDBEndpointAddressWithContext(ctx volcengine.Context, input *ModifyDBEndpointAddressInput, opts ...request.Option) (*ModifyDBEndpointAddressOutput, error) {
	req, out := c.ModifyDBEndpointAddressRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ModifyDBEndpointAddressInput struct {
	_ struct{} `type:"structure"`

	EndpointId *string `type:"string"`

	// InstanceId is a required field
	InstanceId *string `type:"string" required:"true"`

	// NetworkType is a required field
	NetworkType *string `type:"string" required:"true" enum:"EnumOfNetworkTypeForModifyDBEndpointAddressInput"`

	Port *string `type:"string"`
}

// String returns the string representation
func (s ModifyDBEndpointAddressInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyDBEndpointAddressInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ModifyDBEndpointAddressInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ModifyDBEndpointAddressInput"}
	if s.InstanceId == nil {
		invalidParams.Add(request.NewErrParamRequired("InstanceId"))
	}
	if s.NetworkType == nil {
		invalidParams.Add(request.NewErrParamRequired("NetworkType"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetEndpointId sets the EndpointId field's value.
func (s *ModifyDBEndpointAddressInput) SetEndpointId(v string) *ModifyDBEndpointAddressInput {
	s.EndpointId = &v
	return s
}

// SetInstanceId sets the InstanceId field's value.
func (s *ModifyDBEndpointAddressInput) SetInstanceId(v string) *ModifyDBEndpointAddressInput {
	s.InstanceId = &v
	return s
}

// SetNetworkType sets the NetworkType field's value.
func (s *ModifyDBEndpointAddressInput) SetNetworkType(v string) *ModifyDBEndpointAddressInput {
	s.NetworkType = &v
	return s
}

// SetPort sets the Port field's value.
func (s *ModifyDBEndpointAddressInput) SetPort(v string) *ModifyDBEndpointAddressInput {
	s.Port = &v
	return s
}

type ModifyDBEndpointAddressOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s ModifyDBEndpointAddressOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyDBEndpointAddressOutput) GoString() string {
	return s.String()
}

const (
	// EnumOfNetworkTypeForModifyDBEndpointAddressInputPrivate is a EnumOfNetworkTypeForModifyDBEndpointAddressInput enum value
	EnumOfNetworkTypeForModifyDBEndpointAddressInputPrivate = "Private"
)

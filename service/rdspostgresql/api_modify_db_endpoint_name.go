// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package rdspostgresql

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opModifyDBEndpointNameCommon = "ModifyDBEndpointName"

// ModifyDBEndpointNameCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ModifyDBEndpointNameCommon operation. The "output" return
// value will be populated with the ModifyDBEndpointNameCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyDBEndpointNameCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyDBEndpointNameCommon Send returns without error.
//
// See ModifyDBEndpointNameCommon for more information on using the ModifyDBEndpointNameCommon
// API call, and error handling.
//
//    // Example sending a request using the ModifyDBEndpointNameCommonRequest method.
//    req, resp := client.ModifyDBEndpointNameCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *RDSPOSTGRESQL) ModifyDBEndpointNameCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opModifyDBEndpointNameCommon,
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

// ModifyDBEndpointNameCommon API operation for RDS_POSTGRESQL.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RDS_POSTGRESQL's
// API operation ModifyDBEndpointNameCommon for usage and error information.
func (c *RDSPOSTGRESQL) ModifyDBEndpointNameCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ModifyDBEndpointNameCommonRequest(input)
	return out, req.Send()
}

// ModifyDBEndpointNameCommonWithContext is the same as ModifyDBEndpointNameCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyDBEndpointNameCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RDSPOSTGRESQL) ModifyDBEndpointNameCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ModifyDBEndpointNameCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opModifyDBEndpointName = "ModifyDBEndpointName"

// ModifyDBEndpointNameRequest generates a "volcengine/request.Request" representing the
// client's request for the ModifyDBEndpointName operation. The "output" return
// value will be populated with the ModifyDBEndpointNameCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyDBEndpointNameCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyDBEndpointNameCommon Send returns without error.
//
// See ModifyDBEndpointName for more information on using the ModifyDBEndpointName
// API call, and error handling.
//
//    // Example sending a request using the ModifyDBEndpointNameRequest method.
//    req, resp := client.ModifyDBEndpointNameRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *RDSPOSTGRESQL) ModifyDBEndpointNameRequest(input *ModifyDBEndpointNameInput) (req *request.Request, output *ModifyDBEndpointNameOutput) {
	op := &request.Operation{
		Name:       opModifyDBEndpointName,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ModifyDBEndpointNameInput{}
	}

	output = &ModifyDBEndpointNameOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// ModifyDBEndpointName API operation for RDS_POSTGRESQL.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RDS_POSTGRESQL's
// API operation ModifyDBEndpointName for usage and error information.
func (c *RDSPOSTGRESQL) ModifyDBEndpointName(input *ModifyDBEndpointNameInput) (*ModifyDBEndpointNameOutput, error) {
	req, out := c.ModifyDBEndpointNameRequest(input)
	return out, req.Send()
}

// ModifyDBEndpointNameWithContext is the same as ModifyDBEndpointName with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyDBEndpointName for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RDSPOSTGRESQL) ModifyDBEndpointNameWithContext(ctx volcengine.Context, input *ModifyDBEndpointNameInput, opts ...request.Option) (*ModifyDBEndpointNameOutput, error) {
	req, out := c.ModifyDBEndpointNameRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ModifyDBEndpointNameInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	// EndpointId is a required field
	EndpointId *string `type:"string" json:",omitempty" required:"true"`

	// EndpointName is a required field
	EndpointName *string `type:"string" json:",omitempty" required:"true"`

	// InstanceId is a required field
	InstanceId *string `type:"string" json:",omitempty" required:"true"`
}

// String returns the string representation
func (s ModifyDBEndpointNameInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyDBEndpointNameInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ModifyDBEndpointNameInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ModifyDBEndpointNameInput"}
	if s.EndpointId == nil {
		invalidParams.Add(request.NewErrParamRequired("EndpointId"))
	}
	if s.EndpointName == nil {
		invalidParams.Add(request.NewErrParamRequired("EndpointName"))
	}
	if s.InstanceId == nil {
		invalidParams.Add(request.NewErrParamRequired("InstanceId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetEndpointId sets the EndpointId field's value.
func (s *ModifyDBEndpointNameInput) SetEndpointId(v string) *ModifyDBEndpointNameInput {
	s.EndpointId = &v
	return s
}

// SetEndpointName sets the EndpointName field's value.
func (s *ModifyDBEndpointNameInput) SetEndpointName(v string) *ModifyDBEndpointNameInput {
	s.EndpointName = &v
	return s
}

// SetInstanceId sets the InstanceId field's value.
func (s *ModifyDBEndpointNameInput) SetInstanceId(v string) *ModifyDBEndpointNameInput {
	s.InstanceId = &v
	return s
}

type ModifyDBEndpointNameOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s ModifyDBEndpointNameOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyDBEndpointNameOutput) GoString() string {
	return s.String()
}

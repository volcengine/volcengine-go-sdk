// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package rdspostgresql

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opCreateDBEndpointCommon = "CreateDBEndpoint"

// CreateDBEndpointCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateDBEndpointCommon operation. The "output" return
// value will be populated with the CreateDBEndpointCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateDBEndpointCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateDBEndpointCommon Send returns without error.
//
// See CreateDBEndpointCommon for more information on using the CreateDBEndpointCommon
// API call, and error handling.
//
//    // Example sending a request using the CreateDBEndpointCommonRequest method.
//    req, resp := client.CreateDBEndpointCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *RDSPOSTGRESQL) CreateDBEndpointCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opCreateDBEndpointCommon,
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

// CreateDBEndpointCommon API operation for RDS_POSTGRESQL.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RDS_POSTGRESQL's
// API operation CreateDBEndpointCommon for usage and error information.
func (c *RDSPOSTGRESQL) CreateDBEndpointCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.CreateDBEndpointCommonRequest(input)
	return out, req.Send()
}

// CreateDBEndpointCommonWithContext is the same as CreateDBEndpointCommon with the addition of
// the ability to pass a context and additional request options.
//
// See CreateDBEndpointCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RDSPOSTGRESQL) CreateDBEndpointCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.CreateDBEndpointCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opCreateDBEndpoint = "CreateDBEndpoint"

// CreateDBEndpointRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateDBEndpoint operation. The "output" return
// value will be populated with the CreateDBEndpointCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateDBEndpointCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateDBEndpointCommon Send returns without error.
//
// See CreateDBEndpoint for more information on using the CreateDBEndpoint
// API call, and error handling.
//
//    // Example sending a request using the CreateDBEndpointRequest method.
//    req, resp := client.CreateDBEndpointRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *RDSPOSTGRESQL) CreateDBEndpointRequest(input *CreateDBEndpointInput) (req *request.Request, output *CreateDBEndpointOutput) {
	op := &request.Operation{
		Name:       opCreateDBEndpoint,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateDBEndpointInput{}
	}

	output = &CreateDBEndpointOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// CreateDBEndpoint API operation for RDS_POSTGRESQL.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RDS_POSTGRESQL's
// API operation CreateDBEndpoint for usage and error information.
func (c *RDSPOSTGRESQL) CreateDBEndpoint(input *CreateDBEndpointInput) (*CreateDBEndpointOutput, error) {
	req, out := c.CreateDBEndpointRequest(input)
	return out, req.Send()
}

// CreateDBEndpointWithContext is the same as CreateDBEndpoint with the addition of
// the ability to pass a context and additional request options.
//
// See CreateDBEndpoint for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RDSPOSTGRESQL) CreateDBEndpointWithContext(ctx volcengine.Context, input *CreateDBEndpointInput, opts ...request.Option) (*CreateDBEndpointOutput, error) {
	req, out := c.CreateDBEndpointRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type CreateDBEndpointInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	EndpointName *string `type:"string" json:",omitempty"`

	// EndpointType is a required field
	EndpointType *string `type:"string" json:",omitempty" required:"true"`

	// InstanceId is a required field
	InstanceId *string `type:"string" json:",omitempty" required:"true"`

	Nodes *string `type:"string" json:",omitempty"`

	ReadWriteMode *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s CreateDBEndpointInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateDBEndpointInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateDBEndpointInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "CreateDBEndpointInput"}
	if s.EndpointType == nil {
		invalidParams.Add(request.NewErrParamRequired("EndpointType"))
	}
	if s.InstanceId == nil {
		invalidParams.Add(request.NewErrParamRequired("InstanceId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetEndpointName sets the EndpointName field's value.
func (s *CreateDBEndpointInput) SetEndpointName(v string) *CreateDBEndpointInput {
	s.EndpointName = &v
	return s
}

// SetEndpointType sets the EndpointType field's value.
func (s *CreateDBEndpointInput) SetEndpointType(v string) *CreateDBEndpointInput {
	s.EndpointType = &v
	return s
}

// SetInstanceId sets the InstanceId field's value.
func (s *CreateDBEndpointInput) SetInstanceId(v string) *CreateDBEndpointInput {
	s.InstanceId = &v
	return s
}

// SetNodes sets the Nodes field's value.
func (s *CreateDBEndpointInput) SetNodes(v string) *CreateDBEndpointInput {
	s.Nodes = &v
	return s
}

// SetReadWriteMode sets the ReadWriteMode field's value.
func (s *CreateDBEndpointInput) SetReadWriteMode(v string) *CreateDBEndpointInput {
	s.ReadWriteMode = &v
	return s
}

type CreateDBEndpointOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s CreateDBEndpointOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateDBEndpointOutput) GoString() string {
	return s.String()
}

// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package vedbm

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
func (c *VEDBM) CreateDBEndpointCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
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

// CreateDBEndpointCommon API operation for VEDBM.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VEDBM's
// API operation CreateDBEndpointCommon for usage and error information.
func (c *VEDBM) CreateDBEndpointCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
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
func (c *VEDBM) CreateDBEndpointCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
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
func (c *VEDBM) CreateDBEndpointRequest(input *CreateDBEndpointInput) (req *request.Request, output *CreateDBEndpointOutput) {
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

// CreateDBEndpoint API operation for VEDBM.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VEDBM's
// API operation CreateDBEndpoint for usage and error information.
func (c *VEDBM) CreateDBEndpoint(input *CreateDBEndpointInput) (*CreateDBEndpointOutput, error) {
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
func (c *VEDBM) CreateDBEndpointWithContext(ctx volcengine.Context, input *CreateDBEndpointInput, opts ...request.Option) (*CreateDBEndpointOutput, error) {
	req, out := c.CreateDBEndpointRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type CreateDBEndpointInput struct {
	_ struct{} `type:"structure"`

	AutoAddNewNodes *bool `type:"boolean"`

	ConsistLevel *string `type:"string"`

	ConsistTimeout *int32 `type:"int32"`

	ConsistTimeoutAction *string `type:"string"`

	Description *string `type:"string"`

	DistributedTransaction *bool `type:"boolean"`

	EndpointName *string `type:"string"`

	EndpointType *string `type:"string"`

	InstanceId *string `type:"string"`

	MasterAcceptReadRequests *bool `type:"boolean"`

	Nodes *string `type:"string"`

	ReadWriteMode *string `type:"string"`
}

// String returns the string representation
func (s CreateDBEndpointInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateDBEndpointInput) GoString() string {
	return s.String()
}

// SetAutoAddNewNodes sets the AutoAddNewNodes field's value.
func (s *CreateDBEndpointInput) SetAutoAddNewNodes(v bool) *CreateDBEndpointInput {
	s.AutoAddNewNodes = &v
	return s
}

// SetConsistLevel sets the ConsistLevel field's value.
func (s *CreateDBEndpointInput) SetConsistLevel(v string) *CreateDBEndpointInput {
	s.ConsistLevel = &v
	return s
}

// SetConsistTimeout sets the ConsistTimeout field's value.
func (s *CreateDBEndpointInput) SetConsistTimeout(v int32) *CreateDBEndpointInput {
	s.ConsistTimeout = &v
	return s
}

// SetConsistTimeoutAction sets the ConsistTimeoutAction field's value.
func (s *CreateDBEndpointInput) SetConsistTimeoutAction(v string) *CreateDBEndpointInput {
	s.ConsistTimeoutAction = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *CreateDBEndpointInput) SetDescription(v string) *CreateDBEndpointInput {
	s.Description = &v
	return s
}

// SetDistributedTransaction sets the DistributedTransaction field's value.
func (s *CreateDBEndpointInput) SetDistributedTransaction(v bool) *CreateDBEndpointInput {
	s.DistributedTransaction = &v
	return s
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

// SetMasterAcceptReadRequests sets the MasterAcceptReadRequests field's value.
func (s *CreateDBEndpointInput) SetMasterAcceptReadRequests(v bool) *CreateDBEndpointInput {
	s.MasterAcceptReadRequests = &v
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
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	EndpointId *string `type:"string"`
}

// String returns the string representation
func (s CreateDBEndpointOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateDBEndpointOutput) GoString() string {
	return s.String()
}

// SetEndpointId sets the EndpointId field's value.
func (s *CreateDBEndpointOutput) SetEndpointId(v string) *CreateDBEndpointOutput {
	s.EndpointId = &v
	return s
}

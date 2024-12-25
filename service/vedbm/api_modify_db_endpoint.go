// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package vedbm

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opModifyDBEndpointCommon = "ModifyDBEndpoint"

// ModifyDBEndpointCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ModifyDBEndpointCommon operation. The "output" return
// value will be populated with the ModifyDBEndpointCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyDBEndpointCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyDBEndpointCommon Send returns without error.
//
// See ModifyDBEndpointCommon for more information on using the ModifyDBEndpointCommon
// API call, and error handling.
//
//    // Example sending a request using the ModifyDBEndpointCommonRequest method.
//    req, resp := client.ModifyDBEndpointCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VEDBM) ModifyDBEndpointCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opModifyDBEndpointCommon,
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

// ModifyDBEndpointCommon API operation for VEDBM.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VEDBM's
// API operation ModifyDBEndpointCommon for usage and error information.
func (c *VEDBM) ModifyDBEndpointCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ModifyDBEndpointCommonRequest(input)
	return out, req.Send()
}

// ModifyDBEndpointCommonWithContext is the same as ModifyDBEndpointCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyDBEndpointCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VEDBM) ModifyDBEndpointCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ModifyDBEndpointCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opModifyDBEndpoint = "ModifyDBEndpoint"

// ModifyDBEndpointRequest generates a "volcengine/request.Request" representing the
// client's request for the ModifyDBEndpoint operation. The "output" return
// value will be populated with the ModifyDBEndpointCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyDBEndpointCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyDBEndpointCommon Send returns without error.
//
// See ModifyDBEndpoint for more information on using the ModifyDBEndpoint
// API call, and error handling.
//
//    // Example sending a request using the ModifyDBEndpointRequest method.
//    req, resp := client.ModifyDBEndpointRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VEDBM) ModifyDBEndpointRequest(input *ModifyDBEndpointInput) (req *request.Request, output *ModifyDBEndpointOutput) {
	op := &request.Operation{
		Name:       opModifyDBEndpoint,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ModifyDBEndpointInput{}
	}

	output = &ModifyDBEndpointOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// ModifyDBEndpoint API operation for VEDBM.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VEDBM's
// API operation ModifyDBEndpoint for usage and error information.
func (c *VEDBM) ModifyDBEndpoint(input *ModifyDBEndpointInput) (*ModifyDBEndpointOutput, error) {
	req, out := c.ModifyDBEndpointRequest(input)
	return out, req.Send()
}

// ModifyDBEndpointWithContext is the same as ModifyDBEndpoint with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyDBEndpoint for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VEDBM) ModifyDBEndpointWithContext(ctx volcengine.Context, input *ModifyDBEndpointInput, opts ...request.Option) (*ModifyDBEndpointOutput, error) {
	req, out := c.ModifyDBEndpointRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ModifyDBEndpointInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	AutoAddNewNodes *bool `type:"boolean" json:",omitempty"`

	ConsistLevel *string `type:"string" json:",omitempty" enum:"EnumOfConsistLevelForModifyDBEndpointInput"`

	ConsistTimeout *int32 `type:"int32" json:",omitempty"`

	ConsistTimeoutAction *string `type:"string" json:",omitempty" enum:"EnumOfConsistTimeoutActionForModifyDBEndpointInput"`

	Description *string `type:"string" json:",omitempty"`

	DistributedTransaction *bool `type:"boolean" json:",omitempty"`

	// EndpointId is a required field
	EndpointId *string `type:"string" json:",omitempty" required:"true"`

	EndpointName *string `type:"string" json:",omitempty"`

	// InstanceId is a required field
	InstanceId *string `type:"string" json:",omitempty" required:"true"`

	MasterAcceptReadRequests *bool `type:"boolean" json:",omitempty"`

	Nodes *string `type:"string" json:",omitempty"`

	ReadWriteMode *string `type:"string" json:",omitempty" enum:"EnumOfReadWriteModeForModifyDBEndpointInput"`
}

// String returns the string representation
func (s ModifyDBEndpointInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyDBEndpointInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ModifyDBEndpointInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ModifyDBEndpointInput"}
	if s.EndpointId == nil {
		invalidParams.Add(request.NewErrParamRequired("EndpointId"))
	}
	if s.InstanceId == nil {
		invalidParams.Add(request.NewErrParamRequired("InstanceId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetAutoAddNewNodes sets the AutoAddNewNodes field's value.
func (s *ModifyDBEndpointInput) SetAutoAddNewNodes(v bool) *ModifyDBEndpointInput {
	s.AutoAddNewNodes = &v
	return s
}

// SetConsistLevel sets the ConsistLevel field's value.
func (s *ModifyDBEndpointInput) SetConsistLevel(v string) *ModifyDBEndpointInput {
	s.ConsistLevel = &v
	return s
}

// SetConsistTimeout sets the ConsistTimeout field's value.
func (s *ModifyDBEndpointInput) SetConsistTimeout(v int32) *ModifyDBEndpointInput {
	s.ConsistTimeout = &v
	return s
}

// SetConsistTimeoutAction sets the ConsistTimeoutAction field's value.
func (s *ModifyDBEndpointInput) SetConsistTimeoutAction(v string) *ModifyDBEndpointInput {
	s.ConsistTimeoutAction = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *ModifyDBEndpointInput) SetDescription(v string) *ModifyDBEndpointInput {
	s.Description = &v
	return s
}

// SetDistributedTransaction sets the DistributedTransaction field's value.
func (s *ModifyDBEndpointInput) SetDistributedTransaction(v bool) *ModifyDBEndpointInput {
	s.DistributedTransaction = &v
	return s
}

// SetEndpointId sets the EndpointId field's value.
func (s *ModifyDBEndpointInput) SetEndpointId(v string) *ModifyDBEndpointInput {
	s.EndpointId = &v
	return s
}

// SetEndpointName sets the EndpointName field's value.
func (s *ModifyDBEndpointInput) SetEndpointName(v string) *ModifyDBEndpointInput {
	s.EndpointName = &v
	return s
}

// SetInstanceId sets the InstanceId field's value.
func (s *ModifyDBEndpointInput) SetInstanceId(v string) *ModifyDBEndpointInput {
	s.InstanceId = &v
	return s
}

// SetMasterAcceptReadRequests sets the MasterAcceptReadRequests field's value.
func (s *ModifyDBEndpointInput) SetMasterAcceptReadRequests(v bool) *ModifyDBEndpointInput {
	s.MasterAcceptReadRequests = &v
	return s
}

// SetNodes sets the Nodes field's value.
func (s *ModifyDBEndpointInput) SetNodes(v string) *ModifyDBEndpointInput {
	s.Nodes = &v
	return s
}

// SetReadWriteMode sets the ReadWriteMode field's value.
func (s *ModifyDBEndpointInput) SetReadWriteMode(v string) *ModifyDBEndpointInput {
	s.ReadWriteMode = &v
	return s
}

type ModifyDBEndpointOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s ModifyDBEndpointOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyDBEndpointOutput) GoString() string {
	return s.String()
}

const (
	// EnumOfConsistLevelForModifyDBEndpointInputEventual is a EnumOfConsistLevelForModifyDBEndpointInput enum value
	EnumOfConsistLevelForModifyDBEndpointInputEventual = "Eventual"

	// EnumOfConsistLevelForModifyDBEndpointInputSession is a EnumOfConsistLevelForModifyDBEndpointInput enum value
	EnumOfConsistLevelForModifyDBEndpointInputSession = "Session"

	// EnumOfConsistLevelForModifyDBEndpointInputGlobal is a EnumOfConsistLevelForModifyDBEndpointInput enum value
	EnumOfConsistLevelForModifyDBEndpointInputGlobal = "Global"
)

const (
	// EnumOfConsistTimeoutActionForModifyDBEndpointInputReturnError is a EnumOfConsistTimeoutActionForModifyDBEndpointInput enum value
	EnumOfConsistTimeoutActionForModifyDBEndpointInputReturnError = "ReturnError"

	// EnumOfConsistTimeoutActionForModifyDBEndpointInputReadMaster is a EnumOfConsistTimeoutActionForModifyDBEndpointInput enum value
	EnumOfConsistTimeoutActionForModifyDBEndpointInputReadMaster = "ReadMaster"
)

const (
	// EnumOfReadWriteModeForModifyDBEndpointInputReadOnly is a EnumOfReadWriteModeForModifyDBEndpointInput enum value
	EnumOfReadWriteModeForModifyDBEndpointInputReadOnly = "ReadOnly"

	// EnumOfReadWriteModeForModifyDBEndpointInputReadWrite is a EnumOfReadWriteModeForModifyDBEndpointInput enum value
	EnumOfReadWriteModeForModifyDBEndpointInputReadWrite = "ReadWrite"
)

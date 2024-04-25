// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package rdspostgresql

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opModifyDBEndpointReadWriteFlagCommon = "ModifyDBEndpointReadWriteFlag"

// ModifyDBEndpointReadWriteFlagCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ModifyDBEndpointReadWriteFlagCommon operation. The "output" return
// value will be populated with the ModifyDBEndpointReadWriteFlagCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyDBEndpointReadWriteFlagCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyDBEndpointReadWriteFlagCommon Send returns without error.
//
// See ModifyDBEndpointReadWriteFlagCommon for more information on using the ModifyDBEndpointReadWriteFlagCommon
// API call, and error handling.
//
//    // Example sending a request using the ModifyDBEndpointReadWriteFlagCommonRequest method.
//    req, resp := client.ModifyDBEndpointReadWriteFlagCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *RDSPOSTGRESQL) ModifyDBEndpointReadWriteFlagCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opModifyDBEndpointReadWriteFlagCommon,
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

// ModifyDBEndpointReadWriteFlagCommon API operation for RDS_POSTGRESQL.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RDS_POSTGRESQL's
// API operation ModifyDBEndpointReadWriteFlagCommon for usage and error information.
func (c *RDSPOSTGRESQL) ModifyDBEndpointReadWriteFlagCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ModifyDBEndpointReadWriteFlagCommonRequest(input)
	return out, req.Send()
}

// ModifyDBEndpointReadWriteFlagCommonWithContext is the same as ModifyDBEndpointReadWriteFlagCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyDBEndpointReadWriteFlagCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RDSPOSTGRESQL) ModifyDBEndpointReadWriteFlagCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ModifyDBEndpointReadWriteFlagCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opModifyDBEndpointReadWriteFlag = "ModifyDBEndpointReadWriteFlag"

// ModifyDBEndpointReadWriteFlagRequest generates a "volcengine/request.Request" representing the
// client's request for the ModifyDBEndpointReadWriteFlag operation. The "output" return
// value will be populated with the ModifyDBEndpointReadWriteFlagCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyDBEndpointReadWriteFlagCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyDBEndpointReadWriteFlagCommon Send returns without error.
//
// See ModifyDBEndpointReadWriteFlag for more information on using the ModifyDBEndpointReadWriteFlag
// API call, and error handling.
//
//    // Example sending a request using the ModifyDBEndpointReadWriteFlagRequest method.
//    req, resp := client.ModifyDBEndpointReadWriteFlagRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *RDSPOSTGRESQL) ModifyDBEndpointReadWriteFlagRequest(input *ModifyDBEndpointReadWriteFlagInput) (req *request.Request, output *ModifyDBEndpointReadWriteFlagOutput) {
	op := &request.Operation{
		Name:       opModifyDBEndpointReadWriteFlag,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ModifyDBEndpointReadWriteFlagInput{}
	}

	output = &ModifyDBEndpointReadWriteFlagOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// ModifyDBEndpointReadWriteFlag API operation for RDS_POSTGRESQL.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RDS_POSTGRESQL's
// API operation ModifyDBEndpointReadWriteFlag for usage and error information.
func (c *RDSPOSTGRESQL) ModifyDBEndpointReadWriteFlag(input *ModifyDBEndpointReadWriteFlagInput) (*ModifyDBEndpointReadWriteFlagOutput, error) {
	req, out := c.ModifyDBEndpointReadWriteFlagRequest(input)
	return out, req.Send()
}

// ModifyDBEndpointReadWriteFlagWithContext is the same as ModifyDBEndpointReadWriteFlag with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyDBEndpointReadWriteFlag for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RDSPOSTGRESQL) ModifyDBEndpointReadWriteFlagWithContext(ctx volcengine.Context, input *ModifyDBEndpointReadWriteFlagInput, opts ...request.Option) (*ModifyDBEndpointReadWriteFlagOutput, error) {
	req, out := c.ModifyDBEndpointReadWriteFlagRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ModifyDBEndpointReadWriteFlagInput struct {
	_ struct{} `type:"structure"`

	// EndpointId is a required field
	EndpointId *string `type:"string" required:"true"`

	// InstanceId is a required field
	InstanceId *string `type:"string" required:"true"`

	ReadOnlyNodeDistributionType *string `type:"string" enum:"EnumOfReadOnlyNodeDistributionTypeForModifyDBEndpointReadWriteFlagInput"`

	ReadOnlyNodeMaxDelayTime *int32 `type:"int32"`

	ReadOnlyNodeWeight []*ReadOnlyNodeWeightForModifyDBEndpointReadWriteFlagInput `type:"list"`

	// ReadWriteSpliting is a required field
	ReadWriteSpliting *bool `type:"boolean" required:"true"`
}

// String returns the string representation
func (s ModifyDBEndpointReadWriteFlagInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyDBEndpointReadWriteFlagInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ModifyDBEndpointReadWriteFlagInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ModifyDBEndpointReadWriteFlagInput"}
	if s.EndpointId == nil {
		invalidParams.Add(request.NewErrParamRequired("EndpointId"))
	}
	if s.InstanceId == nil {
		invalidParams.Add(request.NewErrParamRequired("InstanceId"))
	}
	if s.ReadWriteSpliting == nil {
		invalidParams.Add(request.NewErrParamRequired("ReadWriteSpliting"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetEndpointId sets the EndpointId field's value.
func (s *ModifyDBEndpointReadWriteFlagInput) SetEndpointId(v string) *ModifyDBEndpointReadWriteFlagInput {
	s.EndpointId = &v
	return s
}

// SetInstanceId sets the InstanceId field's value.
func (s *ModifyDBEndpointReadWriteFlagInput) SetInstanceId(v string) *ModifyDBEndpointReadWriteFlagInput {
	s.InstanceId = &v
	return s
}

// SetReadOnlyNodeDistributionType sets the ReadOnlyNodeDistributionType field's value.
func (s *ModifyDBEndpointReadWriteFlagInput) SetReadOnlyNodeDistributionType(v string) *ModifyDBEndpointReadWriteFlagInput {
	s.ReadOnlyNodeDistributionType = &v
	return s
}

// SetReadOnlyNodeMaxDelayTime sets the ReadOnlyNodeMaxDelayTime field's value.
func (s *ModifyDBEndpointReadWriteFlagInput) SetReadOnlyNodeMaxDelayTime(v int32) *ModifyDBEndpointReadWriteFlagInput {
	s.ReadOnlyNodeMaxDelayTime = &v
	return s
}

// SetReadOnlyNodeWeight sets the ReadOnlyNodeWeight field's value.
func (s *ModifyDBEndpointReadWriteFlagInput) SetReadOnlyNodeWeight(v []*ReadOnlyNodeWeightForModifyDBEndpointReadWriteFlagInput) *ModifyDBEndpointReadWriteFlagInput {
	s.ReadOnlyNodeWeight = v
	return s
}

// SetReadWriteSpliting sets the ReadWriteSpliting field's value.
func (s *ModifyDBEndpointReadWriteFlagInput) SetReadWriteSpliting(v bool) *ModifyDBEndpointReadWriteFlagInput {
	s.ReadWriteSpliting = &v
	return s
}

type ModifyDBEndpointReadWriteFlagOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s ModifyDBEndpointReadWriteFlagOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyDBEndpointReadWriteFlagOutput) GoString() string {
	return s.String()
}

type ReadOnlyNodeWeightForModifyDBEndpointReadWriteFlagInput struct {
	_ struct{} `type:"structure"`

	NodeId *string `type:"string"`

	NodeType *string `type:"string" enum:"EnumOfNodeTypeForModifyDBEndpointReadWriteFlagInput"`

	Weight *int32 `type:"int32"`
}

// String returns the string representation
func (s ReadOnlyNodeWeightForModifyDBEndpointReadWriteFlagInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ReadOnlyNodeWeightForModifyDBEndpointReadWriteFlagInput) GoString() string {
	return s.String()
}

// SetNodeId sets the NodeId field's value.
func (s *ReadOnlyNodeWeightForModifyDBEndpointReadWriteFlagInput) SetNodeId(v string) *ReadOnlyNodeWeightForModifyDBEndpointReadWriteFlagInput {
	s.NodeId = &v
	return s
}

// SetNodeType sets the NodeType field's value.
func (s *ReadOnlyNodeWeightForModifyDBEndpointReadWriteFlagInput) SetNodeType(v string) *ReadOnlyNodeWeightForModifyDBEndpointReadWriteFlagInput {
	s.NodeType = &v
	return s
}

// SetWeight sets the Weight field's value.
func (s *ReadOnlyNodeWeightForModifyDBEndpointReadWriteFlagInput) SetWeight(v int32) *ReadOnlyNodeWeightForModifyDBEndpointReadWriteFlagInput {
	s.Weight = &v
	return s
}

const (
	// EnumOfNodeTypeForModifyDBEndpointReadWriteFlagInputPrimary is a EnumOfNodeTypeForModifyDBEndpointReadWriteFlagInput enum value
	EnumOfNodeTypeForModifyDBEndpointReadWriteFlagInputPrimary = "Primary"

	// EnumOfNodeTypeForModifyDBEndpointReadWriteFlagInputReadOnly is a EnumOfNodeTypeForModifyDBEndpointReadWriteFlagInput enum value
	EnumOfNodeTypeForModifyDBEndpointReadWriteFlagInputReadOnly = "ReadOnly"
)

const (
	// EnumOfReadOnlyNodeDistributionTypeForModifyDBEndpointReadWriteFlagInputDefault is a EnumOfReadOnlyNodeDistributionTypeForModifyDBEndpointReadWriteFlagInput enum value
	EnumOfReadOnlyNodeDistributionTypeForModifyDBEndpointReadWriteFlagInputDefault = "Default"

	// EnumOfReadOnlyNodeDistributionTypeForModifyDBEndpointReadWriteFlagInputCustom is a EnumOfReadOnlyNodeDistributionTypeForModifyDBEndpointReadWriteFlagInput enum value
	EnumOfReadOnlyNodeDistributionTypeForModifyDBEndpointReadWriteFlagInputCustom = "Custom"
)

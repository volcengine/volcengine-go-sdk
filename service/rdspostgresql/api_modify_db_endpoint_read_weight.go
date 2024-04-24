// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package rdspostgresql

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opModifyDBEndpointReadWeightCommon = "ModifyDBEndpointReadWeight"

// ModifyDBEndpointReadWeightCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ModifyDBEndpointReadWeightCommon operation. The "output" return
// value will be populated with the ModifyDBEndpointReadWeightCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyDBEndpointReadWeightCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyDBEndpointReadWeightCommon Send returns without error.
//
// See ModifyDBEndpointReadWeightCommon for more information on using the ModifyDBEndpointReadWeightCommon
// API call, and error handling.
//
//    // Example sending a request using the ModifyDBEndpointReadWeightCommonRequest method.
//    req, resp := client.ModifyDBEndpointReadWeightCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *RDSPOSTGRESQL) ModifyDBEndpointReadWeightCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opModifyDBEndpointReadWeightCommon,
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

// ModifyDBEndpointReadWeightCommon API operation for RDS_POSTGRESQL.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RDS_POSTGRESQL's
// API operation ModifyDBEndpointReadWeightCommon for usage and error information.
func (c *RDSPOSTGRESQL) ModifyDBEndpointReadWeightCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ModifyDBEndpointReadWeightCommonRequest(input)
	return out, req.Send()
}

// ModifyDBEndpointReadWeightCommonWithContext is the same as ModifyDBEndpointReadWeightCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyDBEndpointReadWeightCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RDSPOSTGRESQL) ModifyDBEndpointReadWeightCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ModifyDBEndpointReadWeightCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opModifyDBEndpointReadWeight = "ModifyDBEndpointReadWeight"

// ModifyDBEndpointReadWeightRequest generates a "volcengine/request.Request" representing the
// client's request for the ModifyDBEndpointReadWeight operation. The "output" return
// value will be populated with the ModifyDBEndpointReadWeightCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyDBEndpointReadWeightCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyDBEndpointReadWeightCommon Send returns without error.
//
// See ModifyDBEndpointReadWeight for more information on using the ModifyDBEndpointReadWeight
// API call, and error handling.
//
//    // Example sending a request using the ModifyDBEndpointReadWeightRequest method.
//    req, resp := client.ModifyDBEndpointReadWeightRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *RDSPOSTGRESQL) ModifyDBEndpointReadWeightRequest(input *ModifyDBEndpointReadWeightInput) (req *request.Request, output *ModifyDBEndpointReadWeightOutput) {
	op := &request.Operation{
		Name:       opModifyDBEndpointReadWeight,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ModifyDBEndpointReadWeightInput{}
	}

	output = &ModifyDBEndpointReadWeightOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// ModifyDBEndpointReadWeight API operation for RDS_POSTGRESQL.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RDS_POSTGRESQL's
// API operation ModifyDBEndpointReadWeight for usage and error information.
func (c *RDSPOSTGRESQL) ModifyDBEndpointReadWeight(input *ModifyDBEndpointReadWeightInput) (*ModifyDBEndpointReadWeightOutput, error) {
	req, out := c.ModifyDBEndpointReadWeightRequest(input)
	return out, req.Send()
}

// ModifyDBEndpointReadWeightWithContext is the same as ModifyDBEndpointReadWeight with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyDBEndpointReadWeight for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RDSPOSTGRESQL) ModifyDBEndpointReadWeightWithContext(ctx volcengine.Context, input *ModifyDBEndpointReadWeightInput, opts ...request.Option) (*ModifyDBEndpointReadWeightOutput, error) {
	req, out := c.ModifyDBEndpointReadWeightRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ModifyDBEndpointReadWeightInput struct {
	_ struct{} `type:"structure"`

	// EndpointId is a required field
	EndpointId *string `type:"string" required:"true"`

	// InstanceId is a required field
	InstanceId *string `type:"string" required:"true"`

	// ReadOnlyNodeDistributionType is a required field
	ReadOnlyNodeDistributionType *string `type:"string" required:"true" enum:"EnumOfReadOnlyNodeDistributionTypeForModifyDBEndpointReadWeightInput"`

	ReadOnlyNodeWeight []*ReadOnlyNodeWeightForModifyDBEndpointReadWeightInput `type:"list"`
}

// String returns the string representation
func (s ModifyDBEndpointReadWeightInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyDBEndpointReadWeightInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ModifyDBEndpointReadWeightInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ModifyDBEndpointReadWeightInput"}
	if s.EndpointId == nil {
		invalidParams.Add(request.NewErrParamRequired("EndpointId"))
	}
	if s.InstanceId == nil {
		invalidParams.Add(request.NewErrParamRequired("InstanceId"))
	}
	if s.ReadOnlyNodeDistributionType == nil {
		invalidParams.Add(request.NewErrParamRequired("ReadOnlyNodeDistributionType"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetEndpointId sets the EndpointId field's value.
func (s *ModifyDBEndpointReadWeightInput) SetEndpointId(v string) *ModifyDBEndpointReadWeightInput {
	s.EndpointId = &v
	return s
}

// SetInstanceId sets the InstanceId field's value.
func (s *ModifyDBEndpointReadWeightInput) SetInstanceId(v string) *ModifyDBEndpointReadWeightInput {
	s.InstanceId = &v
	return s
}

// SetReadOnlyNodeDistributionType sets the ReadOnlyNodeDistributionType field's value.
func (s *ModifyDBEndpointReadWeightInput) SetReadOnlyNodeDistributionType(v string) *ModifyDBEndpointReadWeightInput {
	s.ReadOnlyNodeDistributionType = &v
	return s
}

// SetReadOnlyNodeWeight sets the ReadOnlyNodeWeight field's value.
func (s *ModifyDBEndpointReadWeightInput) SetReadOnlyNodeWeight(v []*ReadOnlyNodeWeightForModifyDBEndpointReadWeightInput) *ModifyDBEndpointReadWeightInput {
	s.ReadOnlyNodeWeight = v
	return s
}

type ModifyDBEndpointReadWeightOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s ModifyDBEndpointReadWeightOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyDBEndpointReadWeightOutput) GoString() string {
	return s.String()
}

type ReadOnlyNodeWeightForModifyDBEndpointReadWeightInput struct {
	_ struct{} `type:"structure"`

	NodeId *string `type:"string"`

	NodeType *string `type:"string" enum:"EnumOfNodeTypeForModifyDBEndpointReadWeightInput"`

	Weight *int32 `type:"int32"`
}

// String returns the string representation
func (s ReadOnlyNodeWeightForModifyDBEndpointReadWeightInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ReadOnlyNodeWeightForModifyDBEndpointReadWeightInput) GoString() string {
	return s.String()
}

// SetNodeId sets the NodeId field's value.
func (s *ReadOnlyNodeWeightForModifyDBEndpointReadWeightInput) SetNodeId(v string) *ReadOnlyNodeWeightForModifyDBEndpointReadWeightInput {
	s.NodeId = &v
	return s
}

// SetNodeType sets the NodeType field's value.
func (s *ReadOnlyNodeWeightForModifyDBEndpointReadWeightInput) SetNodeType(v string) *ReadOnlyNodeWeightForModifyDBEndpointReadWeightInput {
	s.NodeType = &v
	return s
}

// SetWeight sets the Weight field's value.
func (s *ReadOnlyNodeWeightForModifyDBEndpointReadWeightInput) SetWeight(v int32) *ReadOnlyNodeWeightForModifyDBEndpointReadWeightInput {
	s.Weight = &v
	return s
}

const (
	// EnumOfNodeTypeForModifyDBEndpointReadWeightInputPrimary is a EnumOfNodeTypeForModifyDBEndpointReadWeightInput enum value
	EnumOfNodeTypeForModifyDBEndpointReadWeightInputPrimary = "Primary"

	// EnumOfNodeTypeForModifyDBEndpointReadWeightInputReadOnly is a EnumOfNodeTypeForModifyDBEndpointReadWeightInput enum value
	EnumOfNodeTypeForModifyDBEndpointReadWeightInputReadOnly = "ReadOnly"
)

const (
	// EnumOfReadOnlyNodeDistributionTypeForModifyDBEndpointReadWeightInputDefault is a EnumOfReadOnlyNodeDistributionTypeForModifyDBEndpointReadWeightInput enum value
	EnumOfReadOnlyNodeDistributionTypeForModifyDBEndpointReadWeightInputDefault = "Default"

	// EnumOfReadOnlyNodeDistributionTypeForModifyDBEndpointReadWeightInputCustom is a EnumOfReadOnlyNodeDistributionTypeForModifyDBEndpointReadWeightInput enum value
	EnumOfReadOnlyNodeDistributionTypeForModifyDBEndpointReadWeightInputCustom = "Custom"
)

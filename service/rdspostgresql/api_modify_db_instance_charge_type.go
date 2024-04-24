// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package rdspostgresql

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opModifyDBInstanceChargeTypeCommon = "ModifyDBInstanceChargeType"

// ModifyDBInstanceChargeTypeCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ModifyDBInstanceChargeTypeCommon operation. The "output" return
// value will be populated with the ModifyDBInstanceChargeTypeCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyDBInstanceChargeTypeCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyDBInstanceChargeTypeCommon Send returns without error.
//
// See ModifyDBInstanceChargeTypeCommon for more information on using the ModifyDBInstanceChargeTypeCommon
// API call, and error handling.
//
//    // Example sending a request using the ModifyDBInstanceChargeTypeCommonRequest method.
//    req, resp := client.ModifyDBInstanceChargeTypeCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *RDSPOSTGRESQL) ModifyDBInstanceChargeTypeCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opModifyDBInstanceChargeTypeCommon,
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

// ModifyDBInstanceChargeTypeCommon API operation for RDS_POSTGRESQL.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RDS_POSTGRESQL's
// API operation ModifyDBInstanceChargeTypeCommon for usage and error information.
func (c *RDSPOSTGRESQL) ModifyDBInstanceChargeTypeCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ModifyDBInstanceChargeTypeCommonRequest(input)
	return out, req.Send()
}

// ModifyDBInstanceChargeTypeCommonWithContext is the same as ModifyDBInstanceChargeTypeCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyDBInstanceChargeTypeCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RDSPOSTGRESQL) ModifyDBInstanceChargeTypeCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ModifyDBInstanceChargeTypeCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opModifyDBInstanceChargeType = "ModifyDBInstanceChargeType"

// ModifyDBInstanceChargeTypeRequest generates a "volcengine/request.Request" representing the
// client's request for the ModifyDBInstanceChargeType operation. The "output" return
// value will be populated with the ModifyDBInstanceChargeTypeCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyDBInstanceChargeTypeCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyDBInstanceChargeTypeCommon Send returns without error.
//
// See ModifyDBInstanceChargeType for more information on using the ModifyDBInstanceChargeType
// API call, and error handling.
//
//    // Example sending a request using the ModifyDBInstanceChargeTypeRequest method.
//    req, resp := client.ModifyDBInstanceChargeTypeRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *RDSPOSTGRESQL) ModifyDBInstanceChargeTypeRequest(input *ModifyDBInstanceChargeTypeInput) (req *request.Request, output *ModifyDBInstanceChargeTypeOutput) {
	op := &request.Operation{
		Name:       opModifyDBInstanceChargeType,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ModifyDBInstanceChargeTypeInput{}
	}

	output = &ModifyDBInstanceChargeTypeOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// ModifyDBInstanceChargeType API operation for RDS_POSTGRESQL.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RDS_POSTGRESQL's
// API operation ModifyDBInstanceChargeType for usage and error information.
func (c *RDSPOSTGRESQL) ModifyDBInstanceChargeType(input *ModifyDBInstanceChargeTypeInput) (*ModifyDBInstanceChargeTypeOutput, error) {
	req, out := c.ModifyDBInstanceChargeTypeRequest(input)
	return out, req.Send()
}

// ModifyDBInstanceChargeTypeWithContext is the same as ModifyDBInstanceChargeType with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyDBInstanceChargeType for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RDSPOSTGRESQL) ModifyDBInstanceChargeTypeWithContext(ctx volcengine.Context, input *ModifyDBInstanceChargeTypeInput, opts ...request.Option) (*ModifyDBInstanceChargeTypeOutput, error) {
	req, out := c.ModifyDBInstanceChargeTypeRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ModifyDBInstanceChargeTypeInput struct {
	_ struct{} `type:"structure"`

	AutoRenew *bool `type:"boolean"`

	// ChargeType is a required field
	ChargeType *string `type:"string" required:"true" enum:"EnumOfChargeTypeForModifyDBInstanceChargeTypeInput"`

	// InstanceId is a required field
	InstanceId *string `type:"string" required:"true"`

	Period *int32 `type:"int32"`

	PeriodUnit *string `type:"string" enum:"EnumOfPeriodUnitForModifyDBInstanceChargeTypeInput"`
}

// String returns the string representation
func (s ModifyDBInstanceChargeTypeInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyDBInstanceChargeTypeInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ModifyDBInstanceChargeTypeInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ModifyDBInstanceChargeTypeInput"}
	if s.ChargeType == nil {
		invalidParams.Add(request.NewErrParamRequired("ChargeType"))
	}
	if s.InstanceId == nil {
		invalidParams.Add(request.NewErrParamRequired("InstanceId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetAutoRenew sets the AutoRenew field's value.
func (s *ModifyDBInstanceChargeTypeInput) SetAutoRenew(v bool) *ModifyDBInstanceChargeTypeInput {
	s.AutoRenew = &v
	return s
}

// SetChargeType sets the ChargeType field's value.
func (s *ModifyDBInstanceChargeTypeInput) SetChargeType(v string) *ModifyDBInstanceChargeTypeInput {
	s.ChargeType = &v
	return s
}

// SetInstanceId sets the InstanceId field's value.
func (s *ModifyDBInstanceChargeTypeInput) SetInstanceId(v string) *ModifyDBInstanceChargeTypeInput {
	s.InstanceId = &v
	return s
}

// SetPeriod sets the Period field's value.
func (s *ModifyDBInstanceChargeTypeInput) SetPeriod(v int32) *ModifyDBInstanceChargeTypeInput {
	s.Period = &v
	return s
}

// SetPeriodUnit sets the PeriodUnit field's value.
func (s *ModifyDBInstanceChargeTypeInput) SetPeriodUnit(v string) *ModifyDBInstanceChargeTypeInput {
	s.PeriodUnit = &v
	return s
}

type ModifyDBInstanceChargeTypeOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	InstanceId *string `type:"string"`

	OrderId *string `type:"string"`
}

// String returns the string representation
func (s ModifyDBInstanceChargeTypeOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyDBInstanceChargeTypeOutput) GoString() string {
	return s.String()
}

// SetInstanceId sets the InstanceId field's value.
func (s *ModifyDBInstanceChargeTypeOutput) SetInstanceId(v string) *ModifyDBInstanceChargeTypeOutput {
	s.InstanceId = &v
	return s
}

// SetOrderId sets the OrderId field's value.
func (s *ModifyDBInstanceChargeTypeOutput) SetOrderId(v string) *ModifyDBInstanceChargeTypeOutput {
	s.OrderId = &v
	return s
}

const (
	// EnumOfChargeTypeForModifyDBInstanceChargeTypeInputPostPaid is a EnumOfChargeTypeForModifyDBInstanceChargeTypeInput enum value
	EnumOfChargeTypeForModifyDBInstanceChargeTypeInputPostPaid = "PostPaid"

	// EnumOfChargeTypeForModifyDBInstanceChargeTypeInputPrePaid is a EnumOfChargeTypeForModifyDBInstanceChargeTypeInput enum value
	EnumOfChargeTypeForModifyDBInstanceChargeTypeInputPrePaid = "PrePaid"
)

const (
	// EnumOfPeriodUnitForModifyDBInstanceChargeTypeInputMonth is a EnumOfPeriodUnitForModifyDBInstanceChargeTypeInput enum value
	EnumOfPeriodUnitForModifyDBInstanceChargeTypeInputMonth = "Month"

	// EnumOfPeriodUnitForModifyDBInstanceChargeTypeInputYear is a EnumOfPeriodUnitForModifyDBInstanceChargeTypeInput enum value
	EnumOfPeriodUnitForModifyDBInstanceChargeTypeInputYear = "Year"
)

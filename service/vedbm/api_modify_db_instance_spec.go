// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package vedbm

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opModifyDBInstanceSpecCommon = "ModifyDBInstanceSpec"

// ModifyDBInstanceSpecCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ModifyDBInstanceSpecCommon operation. The "output" return
// value will be populated with the ModifyDBInstanceSpecCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyDBInstanceSpecCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyDBInstanceSpecCommon Send returns without error.
//
// See ModifyDBInstanceSpecCommon for more information on using the ModifyDBInstanceSpecCommon
// API call, and error handling.
//
//    // Example sending a request using the ModifyDBInstanceSpecCommonRequest method.
//    req, resp := client.ModifyDBInstanceSpecCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VEDBM) ModifyDBInstanceSpecCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opModifyDBInstanceSpecCommon,
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

// ModifyDBInstanceSpecCommon API operation for VEDBM.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VEDBM's
// API operation ModifyDBInstanceSpecCommon for usage and error information.
func (c *VEDBM) ModifyDBInstanceSpecCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ModifyDBInstanceSpecCommonRequest(input)
	return out, req.Send()
}

// ModifyDBInstanceSpecCommonWithContext is the same as ModifyDBInstanceSpecCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyDBInstanceSpecCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VEDBM) ModifyDBInstanceSpecCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ModifyDBInstanceSpecCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opModifyDBInstanceSpec = "ModifyDBInstanceSpec"

// ModifyDBInstanceSpecRequest generates a "volcengine/request.Request" representing the
// client's request for the ModifyDBInstanceSpec operation. The "output" return
// value will be populated with the ModifyDBInstanceSpecCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyDBInstanceSpecCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyDBInstanceSpecCommon Send returns without error.
//
// See ModifyDBInstanceSpec for more information on using the ModifyDBInstanceSpec
// API call, and error handling.
//
//    // Example sending a request using the ModifyDBInstanceSpecRequest method.
//    req, resp := client.ModifyDBInstanceSpecRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VEDBM) ModifyDBInstanceSpecRequest(input *ModifyDBInstanceSpecInput) (req *request.Request, output *ModifyDBInstanceSpecOutput) {
	op := &request.Operation{
		Name:       opModifyDBInstanceSpec,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ModifyDBInstanceSpecInput{}
	}

	output = &ModifyDBInstanceSpecOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// ModifyDBInstanceSpec API operation for VEDBM.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VEDBM's
// API operation ModifyDBInstanceSpec for usage and error information.
func (c *VEDBM) ModifyDBInstanceSpec(input *ModifyDBInstanceSpecInput) (*ModifyDBInstanceSpecOutput, error) {
	req, out := c.ModifyDBInstanceSpecRequest(input)
	return out, req.Send()
}

// ModifyDBInstanceSpecWithContext is the same as ModifyDBInstanceSpec with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyDBInstanceSpec for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VEDBM) ModifyDBInstanceSpecWithContext(ctx volcengine.Context, input *ModifyDBInstanceSpecInput, opts ...request.Option) (*ModifyDBInstanceSpecOutput, error) {
	req, out := c.ModifyDBInstanceSpecRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ModifyDBInstanceSpecInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	// InstanceId is a required field
	InstanceId *string `type:"string" json:",omitempty" required:"true"`

	// NodeNumber is a required field
	NodeNumber *int32 `type:"int32" json:",omitempty" required:"true"`

	// NodeSpec is a required field
	NodeSpec *string `type:"string" json:",omitempty" required:"true" enum:"EnumOfNodeSpecForModifyDBInstanceSpecInput"`

	PrePaidStorageInGB *int32 `type:"int32" json:",omitempty"`
}

// String returns the string representation
func (s ModifyDBInstanceSpecInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyDBInstanceSpecInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ModifyDBInstanceSpecInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ModifyDBInstanceSpecInput"}
	if s.InstanceId == nil {
		invalidParams.Add(request.NewErrParamRequired("InstanceId"))
	}
	if s.NodeNumber == nil {
		invalidParams.Add(request.NewErrParamRequired("NodeNumber"))
	}
	if s.NodeSpec == nil {
		invalidParams.Add(request.NewErrParamRequired("NodeSpec"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetInstanceId sets the InstanceId field's value.
func (s *ModifyDBInstanceSpecInput) SetInstanceId(v string) *ModifyDBInstanceSpecInput {
	s.InstanceId = &v
	return s
}

// SetNodeNumber sets the NodeNumber field's value.
func (s *ModifyDBInstanceSpecInput) SetNodeNumber(v int32) *ModifyDBInstanceSpecInput {
	s.NodeNumber = &v
	return s
}

// SetNodeSpec sets the NodeSpec field's value.
func (s *ModifyDBInstanceSpecInput) SetNodeSpec(v string) *ModifyDBInstanceSpecInput {
	s.NodeSpec = &v
	return s
}

// SetPrePaidStorageInGB sets the PrePaidStorageInGB field's value.
func (s *ModifyDBInstanceSpecInput) SetPrePaidStorageInGB(v int32) *ModifyDBInstanceSpecInput {
	s.PrePaidStorageInGB = &v
	return s
}

type ModifyDBInstanceSpecOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	InstanceId *string `type:"string" json:",omitempty"`

	OrderId *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s ModifyDBInstanceSpecOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyDBInstanceSpecOutput) GoString() string {
	return s.String()
}

// SetInstanceId sets the InstanceId field's value.
func (s *ModifyDBInstanceSpecOutput) SetInstanceId(v string) *ModifyDBInstanceSpecOutput {
	s.InstanceId = &v
	return s
}

// SetOrderId sets the OrderId field's value.
func (s *ModifyDBInstanceSpecOutput) SetOrderId(v string) *ModifyDBInstanceSpecOutput {
	s.OrderId = &v
	return s
}

const (
	// EnumOfNodeSpecForModifyDBInstanceSpecInputVedbMysqlG42xlarge is a EnumOfNodeSpecForModifyDBInstanceSpecInput enum value
	EnumOfNodeSpecForModifyDBInstanceSpecInputVedbMysqlG42xlarge = "vedb.mysql.g4.2xlarge"

	// EnumOfNodeSpecForModifyDBInstanceSpecInputVedbMysqlG44xlarge is a EnumOfNodeSpecForModifyDBInstanceSpecInput enum value
	EnumOfNodeSpecForModifyDBInstanceSpecInputVedbMysqlG44xlarge = "vedb.mysql.g4.4xlarge"

	// EnumOfNodeSpecForModifyDBInstanceSpecInputVedbMysqlG4Large is a EnumOfNodeSpecForModifyDBInstanceSpecInput enum value
	EnumOfNodeSpecForModifyDBInstanceSpecInputVedbMysqlG4Large = "vedb.mysql.g4.large"

	// EnumOfNodeSpecForModifyDBInstanceSpecInputVedbMysqlG4Xlarge is a EnumOfNodeSpecForModifyDBInstanceSpecInput enum value
	EnumOfNodeSpecForModifyDBInstanceSpecInputVedbMysqlG4Xlarge = "vedb.mysql.g4.xlarge"

	// EnumOfNodeSpecForModifyDBInstanceSpecInputVedbMysqlG82xlarge is a EnumOfNodeSpecForModifyDBInstanceSpecInput enum value
	EnumOfNodeSpecForModifyDBInstanceSpecInputVedbMysqlG82xlarge = "vedb.mysql.g8.2xlarge"

	// EnumOfNodeSpecForModifyDBInstanceSpecInputVedbMysqlX42xlarge is a EnumOfNodeSpecForModifyDBInstanceSpecInput enum value
	EnumOfNodeSpecForModifyDBInstanceSpecInputVedbMysqlX42xlarge = "vedb.mysql.x4.2xlarge"

	// EnumOfNodeSpecForModifyDBInstanceSpecInputVedbMysqlX44xlarge is a EnumOfNodeSpecForModifyDBInstanceSpecInput enum value
	EnumOfNodeSpecForModifyDBInstanceSpecInputVedbMysqlX44xlarge = "vedb.mysql.x4.4xlarge"

	// EnumOfNodeSpecForModifyDBInstanceSpecInputVedbMysqlX48xlarge is a EnumOfNodeSpecForModifyDBInstanceSpecInput enum value
	EnumOfNodeSpecForModifyDBInstanceSpecInputVedbMysqlX48xlarge = "vedb.mysql.x4.8xlarge"

	// EnumOfNodeSpecForModifyDBInstanceSpecInputVedbMysqlX4Large is a EnumOfNodeSpecForModifyDBInstanceSpecInput enum value
	EnumOfNodeSpecForModifyDBInstanceSpecInputVedbMysqlX4Large = "vedb.mysql.x4.large"

	// EnumOfNodeSpecForModifyDBInstanceSpecInputVedbMysqlX4Xlarge is a EnumOfNodeSpecForModifyDBInstanceSpecInput enum value
	EnumOfNodeSpecForModifyDBInstanceSpecInputVedbMysqlX4Xlarge = "vedb.mysql.x4.xlarge"

	// EnumOfNodeSpecForModifyDBInstanceSpecInputVedbMysqlX82xlarge is a EnumOfNodeSpecForModifyDBInstanceSpecInput enum value
	EnumOfNodeSpecForModifyDBInstanceSpecInputVedbMysqlX82xlarge = "vedb.mysql.x8.2xlarge"

	// EnumOfNodeSpecForModifyDBInstanceSpecInputVedbMysqlX84xlarge is a EnumOfNodeSpecForModifyDBInstanceSpecInput enum value
	EnumOfNodeSpecForModifyDBInstanceSpecInputVedbMysqlX84xlarge = "vedb.mysql.x8.4xlarge"

	// EnumOfNodeSpecForModifyDBInstanceSpecInputVedbMysqlX86xlarge is a EnumOfNodeSpecForModifyDBInstanceSpecInput enum value
	EnumOfNodeSpecForModifyDBInstanceSpecInputVedbMysqlX86xlarge = "vedb.mysql.x8.6xlarge"

	// EnumOfNodeSpecForModifyDBInstanceSpecInputVedbMysqlX88xlarge is a EnumOfNodeSpecForModifyDBInstanceSpecInput enum value
	EnumOfNodeSpecForModifyDBInstanceSpecInputVedbMysqlX88xlarge = "vedb.mysql.x8.8xlarge"

	// EnumOfNodeSpecForModifyDBInstanceSpecInputVedbMysqlX8Large is a EnumOfNodeSpecForModifyDBInstanceSpecInput enum value
	EnumOfNodeSpecForModifyDBInstanceSpecInputVedbMysqlX8Large = "vedb.mysql.x8.large"

	// EnumOfNodeSpecForModifyDBInstanceSpecInputVedbMysqlX8Xlarge is a EnumOfNodeSpecForModifyDBInstanceSpecInput enum value
	EnumOfNodeSpecForModifyDBInstanceSpecInputVedbMysqlX8Xlarge = "vedb.mysql.x8.xlarge"
)

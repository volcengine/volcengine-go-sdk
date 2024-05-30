// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package rdspostgresql

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opModifyDBInstanceAvailabilityZoneCommon = "ModifyDBInstanceAvailabilityZone"

// ModifyDBInstanceAvailabilityZoneCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ModifyDBInstanceAvailabilityZoneCommon operation. The "output" return
// value will be populated with the ModifyDBInstanceAvailabilityZoneCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyDBInstanceAvailabilityZoneCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyDBInstanceAvailabilityZoneCommon Send returns without error.
//
// See ModifyDBInstanceAvailabilityZoneCommon for more information on using the ModifyDBInstanceAvailabilityZoneCommon
// API call, and error handling.
//
//    // Example sending a request using the ModifyDBInstanceAvailabilityZoneCommonRequest method.
//    req, resp := client.ModifyDBInstanceAvailabilityZoneCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *RDSPOSTGRESQL) ModifyDBInstanceAvailabilityZoneCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opModifyDBInstanceAvailabilityZoneCommon,
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

// ModifyDBInstanceAvailabilityZoneCommon API operation for RDS_POSTGRESQL.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RDS_POSTGRESQL's
// API operation ModifyDBInstanceAvailabilityZoneCommon for usage and error information.
func (c *RDSPOSTGRESQL) ModifyDBInstanceAvailabilityZoneCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ModifyDBInstanceAvailabilityZoneCommonRequest(input)
	return out, req.Send()
}

// ModifyDBInstanceAvailabilityZoneCommonWithContext is the same as ModifyDBInstanceAvailabilityZoneCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyDBInstanceAvailabilityZoneCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RDSPOSTGRESQL) ModifyDBInstanceAvailabilityZoneCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ModifyDBInstanceAvailabilityZoneCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opModifyDBInstanceAvailabilityZone = "ModifyDBInstanceAvailabilityZone"

// ModifyDBInstanceAvailabilityZoneRequest generates a "volcengine/request.Request" representing the
// client's request for the ModifyDBInstanceAvailabilityZone operation. The "output" return
// value will be populated with the ModifyDBInstanceAvailabilityZoneCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyDBInstanceAvailabilityZoneCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyDBInstanceAvailabilityZoneCommon Send returns without error.
//
// See ModifyDBInstanceAvailabilityZone for more information on using the ModifyDBInstanceAvailabilityZone
// API call, and error handling.
//
//    // Example sending a request using the ModifyDBInstanceAvailabilityZoneRequest method.
//    req, resp := client.ModifyDBInstanceAvailabilityZoneRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *RDSPOSTGRESQL) ModifyDBInstanceAvailabilityZoneRequest(input *ModifyDBInstanceAvailabilityZoneInput) (req *request.Request, output *ModifyDBInstanceAvailabilityZoneOutput) {
	op := &request.Operation{
		Name:       opModifyDBInstanceAvailabilityZone,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ModifyDBInstanceAvailabilityZoneInput{}
	}

	output = &ModifyDBInstanceAvailabilityZoneOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// ModifyDBInstanceAvailabilityZone API operation for RDS_POSTGRESQL.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RDS_POSTGRESQL's
// API operation ModifyDBInstanceAvailabilityZone for usage and error information.
func (c *RDSPOSTGRESQL) ModifyDBInstanceAvailabilityZone(input *ModifyDBInstanceAvailabilityZoneInput) (*ModifyDBInstanceAvailabilityZoneOutput, error) {
	req, out := c.ModifyDBInstanceAvailabilityZoneRequest(input)
	return out, req.Send()
}

// ModifyDBInstanceAvailabilityZoneWithContext is the same as ModifyDBInstanceAvailabilityZone with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyDBInstanceAvailabilityZone for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RDSPOSTGRESQL) ModifyDBInstanceAvailabilityZoneWithContext(ctx volcengine.Context, input *ModifyDBInstanceAvailabilityZoneInput, opts ...request.Option) (*ModifyDBInstanceAvailabilityZoneOutput, error) {
	req, out := c.ModifyDBInstanceAvailabilityZoneRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ModifyDBInstanceAvailabilityZoneInput struct {
	_ struct{} `type:"structure"`

	// InstanceId is a required field
	InstanceId *string `type:"string" required:"true"`

	NodeInfo []*NodeInfoForModifyDBInstanceAvailabilityZoneInput `type:"list"`
}

// String returns the string representation
func (s ModifyDBInstanceAvailabilityZoneInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyDBInstanceAvailabilityZoneInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ModifyDBInstanceAvailabilityZoneInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ModifyDBInstanceAvailabilityZoneInput"}
	if s.InstanceId == nil {
		invalidParams.Add(request.NewErrParamRequired("InstanceId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetInstanceId sets the InstanceId field's value.
func (s *ModifyDBInstanceAvailabilityZoneInput) SetInstanceId(v string) *ModifyDBInstanceAvailabilityZoneInput {
	s.InstanceId = &v
	return s
}

// SetNodeInfo sets the NodeInfo field's value.
func (s *ModifyDBInstanceAvailabilityZoneInput) SetNodeInfo(v []*NodeInfoForModifyDBInstanceAvailabilityZoneInput) *ModifyDBInstanceAvailabilityZoneInput {
	s.NodeInfo = v
	return s
}

type ModifyDBInstanceAvailabilityZoneOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	InstanceId *string `type:"string"`

	OrderId *string `type:"string"`
}

// String returns the string representation
func (s ModifyDBInstanceAvailabilityZoneOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyDBInstanceAvailabilityZoneOutput) GoString() string {
	return s.String()
}

// SetInstanceId sets the InstanceId field's value.
func (s *ModifyDBInstanceAvailabilityZoneOutput) SetInstanceId(v string) *ModifyDBInstanceAvailabilityZoneOutput {
	s.InstanceId = &v
	return s
}

// SetOrderId sets the OrderId field's value.
func (s *ModifyDBInstanceAvailabilityZoneOutput) SetOrderId(v string) *ModifyDBInstanceAvailabilityZoneOutput {
	s.OrderId = &v
	return s
}

type NodeInfoForModifyDBInstanceAvailabilityZoneInput struct {
	_ struct{} `type:"structure"`

	NodeId *string `type:"string"`

	ZoneId *string `type:"string"`
}

// String returns the string representation
func (s NodeInfoForModifyDBInstanceAvailabilityZoneInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s NodeInfoForModifyDBInstanceAvailabilityZoneInput) GoString() string {
	return s.String()
}

// SetNodeId sets the NodeId field's value.
func (s *NodeInfoForModifyDBInstanceAvailabilityZoneInput) SetNodeId(v string) *NodeInfoForModifyDBInstanceAvailabilityZoneInput {
	s.NodeId = &v
	return s
}

// SetZoneId sets the ZoneId field's value.
func (s *NodeInfoForModifyDBInstanceAvailabilityZoneInput) SetZoneId(v string) *NodeInfoForModifyDBInstanceAvailabilityZoneInput {
	s.ZoneId = &v
	return s
}
// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package rdsmysqlv2

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
func (c *RDSMYSQLV2) ModifyDBInstanceSpecCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
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

// ModifyDBInstanceSpecCommon API operation for RDS_MYSQL_V2.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RDS_MYSQL_V2's
// API operation ModifyDBInstanceSpecCommon for usage and error information.
func (c *RDSMYSQLV2) ModifyDBInstanceSpecCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
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
func (c *RDSMYSQLV2) ModifyDBInstanceSpecCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
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
func (c *RDSMYSQLV2) ModifyDBInstanceSpecRequest(input *ModifyDBInstanceSpecInput) (req *request.Request, output *ModifyDBInstanceSpecOutput) {
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

// ModifyDBInstanceSpec API operation for RDS_MYSQL_V2.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RDS_MYSQL_V2's
// API operation ModifyDBInstanceSpec for usage and error information.
func (c *RDSMYSQLV2) ModifyDBInstanceSpec(input *ModifyDBInstanceSpecInput) (*ModifyDBInstanceSpecOutput, error) {
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
func (c *RDSMYSQLV2) ModifyDBInstanceSpecWithContext(ctx volcengine.Context, input *ModifyDBInstanceSpecInput, opts ...request.Option) (*ModifyDBInstanceSpecOutput, error) {
	req, out := c.ModifyDBInstanceSpecRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ModifyDBInstanceSpecInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	// InstanceId is a required field
	InstanceId *string `type:"string" json:",omitempty" required:"true"`

	ModifyType *string `type:"string" json:",omitempty" enum:"EnumOfModifyTypeForModifyDBInstanceSpecInput"`

	NodeInfo []*NodeInfoForModifyDBInstanceSpecInput `type:"list" json:",omitempty"`

	StorageSpace *int32 `type:"int32" json:",omitempty"`

	// StorageType is a required field
	StorageType *string `type:"string" json:",omitempty" required:"true" enum:"EnumOfStorageTypeForModifyDBInstanceSpecInput"`

	SwitchType *string `type:"string" json:",omitempty" enum:"EnumOfSwitchTypeForModifyDBInstanceSpecInput"`
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
	if s.StorageType == nil {
		invalidParams.Add(request.NewErrParamRequired("StorageType"))
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

// SetModifyType sets the ModifyType field's value.
func (s *ModifyDBInstanceSpecInput) SetModifyType(v string) *ModifyDBInstanceSpecInput {
	s.ModifyType = &v
	return s
}

// SetNodeInfo sets the NodeInfo field's value.
func (s *ModifyDBInstanceSpecInput) SetNodeInfo(v []*NodeInfoForModifyDBInstanceSpecInput) *ModifyDBInstanceSpecInput {
	s.NodeInfo = v
	return s
}

// SetStorageSpace sets the StorageSpace field's value.
func (s *ModifyDBInstanceSpecInput) SetStorageSpace(v int32) *ModifyDBInstanceSpecInput {
	s.StorageSpace = &v
	return s
}

// SetStorageType sets the StorageType field's value.
func (s *ModifyDBInstanceSpecInput) SetStorageType(v string) *ModifyDBInstanceSpecInput {
	s.StorageType = &v
	return s
}

// SetSwitchType sets the SwitchType field's value.
func (s *ModifyDBInstanceSpecInput) SetSwitchType(v string) *ModifyDBInstanceSpecInput {
	s.SwitchType = &v
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

type NodeInfoForModifyDBInstanceSpecInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	NodeId *string `type:"string" json:",omitempty"`

	NodeOperateType *string `type:"string" json:",omitempty"`

	NodeSpec *string `type:"string" json:",omitempty"`

	NodeType *string `type:"string" json:",omitempty"`

	ZoneId *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s NodeInfoForModifyDBInstanceSpecInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s NodeInfoForModifyDBInstanceSpecInput) GoString() string {
	return s.String()
}

// SetNodeId sets the NodeId field's value.
func (s *NodeInfoForModifyDBInstanceSpecInput) SetNodeId(v string) *NodeInfoForModifyDBInstanceSpecInput {
	s.NodeId = &v
	return s
}

// SetNodeOperateType sets the NodeOperateType field's value.
func (s *NodeInfoForModifyDBInstanceSpecInput) SetNodeOperateType(v string) *NodeInfoForModifyDBInstanceSpecInput {
	s.NodeOperateType = &v
	return s
}

// SetNodeSpec sets the NodeSpec field's value.
func (s *NodeInfoForModifyDBInstanceSpecInput) SetNodeSpec(v string) *NodeInfoForModifyDBInstanceSpecInput {
	s.NodeSpec = &v
	return s
}

// SetNodeType sets the NodeType field's value.
func (s *NodeInfoForModifyDBInstanceSpecInput) SetNodeType(v string) *NodeInfoForModifyDBInstanceSpecInput {
	s.NodeType = &v
	return s
}

// SetZoneId sets the ZoneId field's value.
func (s *NodeInfoForModifyDBInstanceSpecInput) SetZoneId(v string) *NodeInfoForModifyDBInstanceSpecInput {
	s.ZoneId = &v
	return s
}

const (
	// EnumOfModifyTypeForModifyDBInstanceSpecInputUsually is a EnumOfModifyTypeForModifyDBInstanceSpecInput enum value
	EnumOfModifyTypeForModifyDBInstanceSpecInputUsually = "Usually"

	// EnumOfModifyTypeForModifyDBInstanceSpecInputTemporary is a EnumOfModifyTypeForModifyDBInstanceSpecInput enum value
	EnumOfModifyTypeForModifyDBInstanceSpecInputTemporary = "Temporary"
)

const (
	// EnumOfStorageTypeForModifyDBInstanceSpecInputLocalSsd is a EnumOfStorageTypeForModifyDBInstanceSpecInput enum value
	EnumOfStorageTypeForModifyDBInstanceSpecInputLocalSsd = "LocalSSD"

	// EnumOfStorageTypeForModifyDBInstanceSpecInputCloudStorage is a EnumOfStorageTypeForModifyDBInstanceSpecInput enum value
	EnumOfStorageTypeForModifyDBInstanceSpecInputCloudStorage = "CloudStorage"

	// EnumOfStorageTypeForModifyDBInstanceSpecInputEssdpl1 is a EnumOfStorageTypeForModifyDBInstanceSpecInput enum value
	EnumOfStorageTypeForModifyDBInstanceSpecInputEssdpl1 = "ESSDPL1"

	// EnumOfStorageTypeForModifyDBInstanceSpecInputEssdpl2 is a EnumOfStorageTypeForModifyDBInstanceSpecInput enum value
	EnumOfStorageTypeForModifyDBInstanceSpecInputEssdpl2 = "ESSDPL2"
)

const (
	// EnumOfSwitchTypeForModifyDBInstanceSpecInputImmediate is a EnumOfSwitchTypeForModifyDBInstanceSpecInput enum value
	EnumOfSwitchTypeForModifyDBInstanceSpecInputImmediate = "Immediate"

	// EnumOfSwitchTypeForModifyDBInstanceSpecInputMaintainTime is a EnumOfSwitchTypeForModifyDBInstanceSpecInput enum value
	EnumOfSwitchTypeForModifyDBInstanceSpecInputMaintainTime = "MaintainTime"

	// EnumOfSwitchTypeForModifyDBInstanceSpecInputSpecifiedTime is a EnumOfSwitchTypeForModifyDBInstanceSpecInput enum value
	EnumOfSwitchTypeForModifyDBInstanceSpecInputSpecifiedTime = "SpecifiedTime"
)

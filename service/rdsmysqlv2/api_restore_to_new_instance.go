// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package rdsmysqlv2

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opRestoreToNewInstanceCommon = "RestoreToNewInstance"

// RestoreToNewInstanceCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the RestoreToNewInstanceCommon operation. The "output" return
// value will be populated with the RestoreToNewInstanceCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned RestoreToNewInstanceCommon Request to send the API call to the service.
// the "output" return value is not valid until after RestoreToNewInstanceCommon Send returns without error.
//
// See RestoreToNewInstanceCommon for more information on using the RestoreToNewInstanceCommon
// API call, and error handling.
//
//    // Example sending a request using the RestoreToNewInstanceCommonRequest method.
//    req, resp := client.RestoreToNewInstanceCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *RDSMYSQLV2) RestoreToNewInstanceCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opRestoreToNewInstanceCommon,
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

// RestoreToNewInstanceCommon API operation for RDS_MYSQL_V2.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RDS_MYSQL_V2's
// API operation RestoreToNewInstanceCommon for usage and error information.
func (c *RDSMYSQLV2) RestoreToNewInstanceCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.RestoreToNewInstanceCommonRequest(input)
	return out, req.Send()
}

// RestoreToNewInstanceCommonWithContext is the same as RestoreToNewInstanceCommon with the addition of
// the ability to pass a context and additional request options.
//
// See RestoreToNewInstanceCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RDSMYSQLV2) RestoreToNewInstanceCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.RestoreToNewInstanceCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opRestoreToNewInstance = "RestoreToNewInstance"

// RestoreToNewInstanceRequest generates a "volcengine/request.Request" representing the
// client's request for the RestoreToNewInstance operation. The "output" return
// value will be populated with the RestoreToNewInstanceCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned RestoreToNewInstanceCommon Request to send the API call to the service.
// the "output" return value is not valid until after RestoreToNewInstanceCommon Send returns without error.
//
// See RestoreToNewInstance for more information on using the RestoreToNewInstance
// API call, and error handling.
//
//    // Example sending a request using the RestoreToNewInstanceRequest method.
//    req, resp := client.RestoreToNewInstanceRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *RDSMYSQLV2) RestoreToNewInstanceRequest(input *RestoreToNewInstanceInput) (req *request.Request, output *RestoreToNewInstanceOutput) {
	op := &request.Operation{
		Name:       opRestoreToNewInstance,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &RestoreToNewInstanceInput{}
	}

	output = &RestoreToNewInstanceOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// RestoreToNewInstance API operation for RDS_MYSQL_V2.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RDS_MYSQL_V2's
// API operation RestoreToNewInstance for usage and error information.
func (c *RDSMYSQLV2) RestoreToNewInstance(input *RestoreToNewInstanceInput) (*RestoreToNewInstanceOutput, error) {
	req, out := c.RestoreToNewInstanceRequest(input)
	return out, req.Send()
}

// RestoreToNewInstanceWithContext is the same as RestoreToNewInstance with the addition of
// the ability to pass a context and additional request options.
//
// See RestoreToNewInstance for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RDSMYSQLV2) RestoreToNewInstanceWithContext(ctx volcengine.Context, input *RestoreToNewInstanceInput, opts ...request.Option) (*RestoreToNewInstanceOutput, error) {
	req, out := c.RestoreToNewInstanceRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ChargeInfoForRestoreToNewInstanceInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	AutoRenew *bool `type:"boolean" json:",omitempty"`

	ChargeType *string `type:"string" json:",omitempty"`

	Number *int32 `type:"int32" json:",omitempty"`

	Period *int32 `type:"int32" json:",omitempty"`

	PeriodUnit *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s ChargeInfoForRestoreToNewInstanceInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ChargeInfoForRestoreToNewInstanceInput) GoString() string {
	return s.String()
}

// SetAutoRenew sets the AutoRenew field's value.
func (s *ChargeInfoForRestoreToNewInstanceInput) SetAutoRenew(v bool) *ChargeInfoForRestoreToNewInstanceInput {
	s.AutoRenew = &v
	return s
}

// SetChargeType sets the ChargeType field's value.
func (s *ChargeInfoForRestoreToNewInstanceInput) SetChargeType(v string) *ChargeInfoForRestoreToNewInstanceInput {
	s.ChargeType = &v
	return s
}

// SetNumber sets the Number field's value.
func (s *ChargeInfoForRestoreToNewInstanceInput) SetNumber(v int32) *ChargeInfoForRestoreToNewInstanceInput {
	s.Number = &v
	return s
}

// SetPeriod sets the Period field's value.
func (s *ChargeInfoForRestoreToNewInstanceInput) SetPeriod(v int32) *ChargeInfoForRestoreToNewInstanceInput {
	s.Period = &v
	return s
}

// SetPeriodUnit sets the PeriodUnit field's value.
func (s *ChargeInfoForRestoreToNewInstanceInput) SetPeriodUnit(v string) *ChargeInfoForRestoreToNewInstanceInput {
	s.PeriodUnit = &v
	return s
}

type InstanceTagForRestoreToNewInstanceInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Key *string `type:"string" json:",omitempty"`

	Value *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s InstanceTagForRestoreToNewInstanceInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s InstanceTagForRestoreToNewInstanceInput) GoString() string {
	return s.String()
}

// SetKey sets the Key field's value.
func (s *InstanceTagForRestoreToNewInstanceInput) SetKey(v string) *InstanceTagForRestoreToNewInstanceInput {
	s.Key = &v
	return s
}

// SetValue sets the Value field's value.
func (s *InstanceTagForRestoreToNewInstanceInput) SetValue(v string) *InstanceTagForRestoreToNewInstanceInput {
	s.Value = &v
	return s
}

type NodeInfoForRestoreToNewInstanceInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	NodeId *string `type:"string" json:",omitempty"`

	NodeOperateType *string `type:"string" json:",omitempty"`

	NodeSpec *string `type:"string" json:",omitempty"`

	NodeType *string `type:"string" json:",omitempty"`

	ZoneId *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s NodeInfoForRestoreToNewInstanceInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s NodeInfoForRestoreToNewInstanceInput) GoString() string {
	return s.String()
}

// SetNodeId sets the NodeId field's value.
func (s *NodeInfoForRestoreToNewInstanceInput) SetNodeId(v string) *NodeInfoForRestoreToNewInstanceInput {
	s.NodeId = &v
	return s
}

// SetNodeOperateType sets the NodeOperateType field's value.
func (s *NodeInfoForRestoreToNewInstanceInput) SetNodeOperateType(v string) *NodeInfoForRestoreToNewInstanceInput {
	s.NodeOperateType = &v
	return s
}

// SetNodeSpec sets the NodeSpec field's value.
func (s *NodeInfoForRestoreToNewInstanceInput) SetNodeSpec(v string) *NodeInfoForRestoreToNewInstanceInput {
	s.NodeSpec = &v
	return s
}

// SetNodeType sets the NodeType field's value.
func (s *NodeInfoForRestoreToNewInstanceInput) SetNodeType(v string) *NodeInfoForRestoreToNewInstanceInput {
	s.NodeType = &v
	return s
}

// SetZoneId sets the ZoneId field's value.
func (s *NodeInfoForRestoreToNewInstanceInput) SetZoneId(v string) *NodeInfoForRestoreToNewInstanceInput {
	s.ZoneId = &v
	return s
}

type RestoreToNewInstanceInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	AllowListIds []*string `type:"list" json:",omitempty"`

	BackupId *string `type:"string" json:",omitempty"`

	ChargeInfo *ChargeInfoForRestoreToNewInstanceInput `type:"structure" json:",omitempty"`

	DBParamGroupId *string `type:"string" json:",omitempty"`

	InstanceName *string `type:"string" json:",omitempty"`

	InstanceTags []*InstanceTagForRestoreToNewInstanceInput `type:"list" json:",omitempty"`

	NodeInfo []*NodeInfoForRestoreToNewInstanceInput `type:"list" json:",omitempty"`

	Port *int32 `type:"int32" json:",omitempty"`

	ProjectName *string `type:"string" json:",omitempty"`

	RestoreTime *string `type:"string" json:",omitempty"`

	// SrcInstanceId is a required field
	SrcInstanceId *string `type:"string" json:",omitempty" required:"true"`

	StorageSpace *int32 `type:"int32" json:",omitempty"`

	// StorageType is a required field
	StorageType *string `type:"string" json:",omitempty" required:"true"`

	// SubnetId is a required field
	SubnetId *string `type:"string" json:",omitempty" required:"true"`

	// VpcId is a required field
	VpcId *string `type:"string" json:",omitempty" required:"true"`
}

// String returns the string representation
func (s RestoreToNewInstanceInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s RestoreToNewInstanceInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RestoreToNewInstanceInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "RestoreToNewInstanceInput"}
	if s.SrcInstanceId == nil {
		invalidParams.Add(request.NewErrParamRequired("SrcInstanceId"))
	}
	if s.StorageType == nil {
		invalidParams.Add(request.NewErrParamRequired("StorageType"))
	}
	if s.SubnetId == nil {
		invalidParams.Add(request.NewErrParamRequired("SubnetId"))
	}
	if s.VpcId == nil {
		invalidParams.Add(request.NewErrParamRequired("VpcId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetAllowListIds sets the AllowListIds field's value.
func (s *RestoreToNewInstanceInput) SetAllowListIds(v []*string) *RestoreToNewInstanceInput {
	s.AllowListIds = v
	return s
}

// SetBackupId sets the BackupId field's value.
func (s *RestoreToNewInstanceInput) SetBackupId(v string) *RestoreToNewInstanceInput {
	s.BackupId = &v
	return s
}

// SetChargeInfo sets the ChargeInfo field's value.
func (s *RestoreToNewInstanceInput) SetChargeInfo(v *ChargeInfoForRestoreToNewInstanceInput) *RestoreToNewInstanceInput {
	s.ChargeInfo = v
	return s
}

// SetDBParamGroupId sets the DBParamGroupId field's value.
func (s *RestoreToNewInstanceInput) SetDBParamGroupId(v string) *RestoreToNewInstanceInput {
	s.DBParamGroupId = &v
	return s
}

// SetInstanceName sets the InstanceName field's value.
func (s *RestoreToNewInstanceInput) SetInstanceName(v string) *RestoreToNewInstanceInput {
	s.InstanceName = &v
	return s
}

// SetInstanceTags sets the InstanceTags field's value.
func (s *RestoreToNewInstanceInput) SetInstanceTags(v []*InstanceTagForRestoreToNewInstanceInput) *RestoreToNewInstanceInput {
	s.InstanceTags = v
	return s
}

// SetNodeInfo sets the NodeInfo field's value.
func (s *RestoreToNewInstanceInput) SetNodeInfo(v []*NodeInfoForRestoreToNewInstanceInput) *RestoreToNewInstanceInput {
	s.NodeInfo = v
	return s
}

// SetPort sets the Port field's value.
func (s *RestoreToNewInstanceInput) SetPort(v int32) *RestoreToNewInstanceInput {
	s.Port = &v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *RestoreToNewInstanceInput) SetProjectName(v string) *RestoreToNewInstanceInput {
	s.ProjectName = &v
	return s
}

// SetRestoreTime sets the RestoreTime field's value.
func (s *RestoreToNewInstanceInput) SetRestoreTime(v string) *RestoreToNewInstanceInput {
	s.RestoreTime = &v
	return s
}

// SetSrcInstanceId sets the SrcInstanceId field's value.
func (s *RestoreToNewInstanceInput) SetSrcInstanceId(v string) *RestoreToNewInstanceInput {
	s.SrcInstanceId = &v
	return s
}

// SetStorageSpace sets the StorageSpace field's value.
func (s *RestoreToNewInstanceInput) SetStorageSpace(v int32) *RestoreToNewInstanceInput {
	s.StorageSpace = &v
	return s
}

// SetStorageType sets the StorageType field's value.
func (s *RestoreToNewInstanceInput) SetStorageType(v string) *RestoreToNewInstanceInput {
	s.StorageType = &v
	return s
}

// SetSubnetId sets the SubnetId field's value.
func (s *RestoreToNewInstanceInput) SetSubnetId(v string) *RestoreToNewInstanceInput {
	s.SubnetId = &v
	return s
}

// SetVpcId sets the VpcId field's value.
func (s *RestoreToNewInstanceInput) SetVpcId(v string) *RestoreToNewInstanceInput {
	s.VpcId = &v
	return s
}

type RestoreToNewInstanceOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	InstanceId *string `type:"string" json:",omitempty"`

	OrderId *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s RestoreToNewInstanceOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s RestoreToNewInstanceOutput) GoString() string {
	return s.String()
}

// SetInstanceId sets the InstanceId field's value.
func (s *RestoreToNewInstanceOutput) SetInstanceId(v string) *RestoreToNewInstanceOutput {
	s.InstanceId = &v
	return s
}

// SetOrderId sets the OrderId field's value.
func (s *RestoreToNewInstanceOutput) SetOrderId(v string) *RestoreToNewInstanceOutput {
	s.OrderId = &v
	return s
}

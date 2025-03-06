// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package rdsmysqlv2

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opRestoreToCrossRegionInstanceCommon = "RestoreToCrossRegionInstance"

// RestoreToCrossRegionInstanceCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the RestoreToCrossRegionInstanceCommon operation. The "output" return
// value will be populated with the RestoreToCrossRegionInstanceCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned RestoreToCrossRegionInstanceCommon Request to send the API call to the service.
// the "output" return value is not valid until after RestoreToCrossRegionInstanceCommon Send returns without error.
//
// See RestoreToCrossRegionInstanceCommon for more information on using the RestoreToCrossRegionInstanceCommon
// API call, and error handling.
//
//    // Example sending a request using the RestoreToCrossRegionInstanceCommonRequest method.
//    req, resp := client.RestoreToCrossRegionInstanceCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *RDSMYSQLV2) RestoreToCrossRegionInstanceCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opRestoreToCrossRegionInstanceCommon,
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

// RestoreToCrossRegionInstanceCommon API operation for RDS_MYSQL_V2.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RDS_MYSQL_V2's
// API operation RestoreToCrossRegionInstanceCommon for usage and error information.
func (c *RDSMYSQLV2) RestoreToCrossRegionInstanceCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.RestoreToCrossRegionInstanceCommonRequest(input)
	return out, req.Send()
}

// RestoreToCrossRegionInstanceCommonWithContext is the same as RestoreToCrossRegionInstanceCommon with the addition of
// the ability to pass a context and additional request options.
//
// See RestoreToCrossRegionInstanceCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RDSMYSQLV2) RestoreToCrossRegionInstanceCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.RestoreToCrossRegionInstanceCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opRestoreToCrossRegionInstance = "RestoreToCrossRegionInstance"

// RestoreToCrossRegionInstanceRequest generates a "volcengine/request.Request" representing the
// client's request for the RestoreToCrossRegionInstance operation. The "output" return
// value will be populated with the RestoreToCrossRegionInstanceCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned RestoreToCrossRegionInstanceCommon Request to send the API call to the service.
// the "output" return value is not valid until after RestoreToCrossRegionInstanceCommon Send returns without error.
//
// See RestoreToCrossRegionInstance for more information on using the RestoreToCrossRegionInstance
// API call, and error handling.
//
//    // Example sending a request using the RestoreToCrossRegionInstanceRequest method.
//    req, resp := client.RestoreToCrossRegionInstanceRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *RDSMYSQLV2) RestoreToCrossRegionInstanceRequest(input *RestoreToCrossRegionInstanceInput) (req *request.Request, output *RestoreToCrossRegionInstanceOutput) {
	op := &request.Operation{
		Name:       opRestoreToCrossRegionInstance,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &RestoreToCrossRegionInstanceInput{}
	}

	output = &RestoreToCrossRegionInstanceOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// RestoreToCrossRegionInstance API operation for RDS_MYSQL_V2.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RDS_MYSQL_V2's
// API operation RestoreToCrossRegionInstance for usage and error information.
func (c *RDSMYSQLV2) RestoreToCrossRegionInstance(input *RestoreToCrossRegionInstanceInput) (*RestoreToCrossRegionInstanceOutput, error) {
	req, out := c.RestoreToCrossRegionInstanceRequest(input)
	return out, req.Send()
}

// RestoreToCrossRegionInstanceWithContext is the same as RestoreToCrossRegionInstance with the addition of
// the ability to pass a context and additional request options.
//
// See RestoreToCrossRegionInstance for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RDSMYSQLV2) RestoreToCrossRegionInstanceWithContext(ctx volcengine.Context, input *RestoreToCrossRegionInstanceInput, opts ...request.Option) (*RestoreToCrossRegionInstanceOutput, error) {
	req, out := c.RestoreToCrossRegionInstanceRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ChargeInfoForRestoreToCrossRegionInstanceInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	AutoRenew *bool `type:"boolean" json:",omitempty"`

	ChargeType *string `type:"string" json:",omitempty"`

	Number *int32 `type:"int32" json:",omitempty"`

	Period *int32 `type:"int32" json:",omitempty"`

	PeriodUnit *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s ChargeInfoForRestoreToCrossRegionInstanceInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ChargeInfoForRestoreToCrossRegionInstanceInput) GoString() string {
	return s.String()
}

// SetAutoRenew sets the AutoRenew field's value.
func (s *ChargeInfoForRestoreToCrossRegionInstanceInput) SetAutoRenew(v bool) *ChargeInfoForRestoreToCrossRegionInstanceInput {
	s.AutoRenew = &v
	return s
}

// SetChargeType sets the ChargeType field's value.
func (s *ChargeInfoForRestoreToCrossRegionInstanceInput) SetChargeType(v string) *ChargeInfoForRestoreToCrossRegionInstanceInput {
	s.ChargeType = &v
	return s
}

// SetNumber sets the Number field's value.
func (s *ChargeInfoForRestoreToCrossRegionInstanceInput) SetNumber(v int32) *ChargeInfoForRestoreToCrossRegionInstanceInput {
	s.Number = &v
	return s
}

// SetPeriod sets the Period field's value.
func (s *ChargeInfoForRestoreToCrossRegionInstanceInput) SetPeriod(v int32) *ChargeInfoForRestoreToCrossRegionInstanceInput {
	s.Period = &v
	return s
}

// SetPeriodUnit sets the PeriodUnit field's value.
func (s *ChargeInfoForRestoreToCrossRegionInstanceInput) SetPeriodUnit(v string) *ChargeInfoForRestoreToCrossRegionInstanceInput {
	s.PeriodUnit = &v
	return s
}

type InstanceTagForRestoreToCrossRegionInstanceInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Key *string `type:"string" json:",omitempty"`

	Value *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s InstanceTagForRestoreToCrossRegionInstanceInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s InstanceTagForRestoreToCrossRegionInstanceInput) GoString() string {
	return s.String()
}

// SetKey sets the Key field's value.
func (s *InstanceTagForRestoreToCrossRegionInstanceInput) SetKey(v string) *InstanceTagForRestoreToCrossRegionInstanceInput {
	s.Key = &v
	return s
}

// SetValue sets the Value field's value.
func (s *InstanceTagForRestoreToCrossRegionInstanceInput) SetValue(v string) *InstanceTagForRestoreToCrossRegionInstanceInput {
	s.Value = &v
	return s
}

type NodeInfoForRestoreToCrossRegionInstanceInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	NodeId *string `type:"string" json:",omitempty"`

	NodeOperateType *string `type:"string" json:",omitempty"`

	NodeSpec *string `type:"string" json:",omitempty"`

	NodeType *string `type:"string" json:",omitempty"`

	ZoneId *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s NodeInfoForRestoreToCrossRegionInstanceInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s NodeInfoForRestoreToCrossRegionInstanceInput) GoString() string {
	return s.String()
}

// SetNodeId sets the NodeId field's value.
func (s *NodeInfoForRestoreToCrossRegionInstanceInput) SetNodeId(v string) *NodeInfoForRestoreToCrossRegionInstanceInput {
	s.NodeId = &v
	return s
}

// SetNodeOperateType sets the NodeOperateType field's value.
func (s *NodeInfoForRestoreToCrossRegionInstanceInput) SetNodeOperateType(v string) *NodeInfoForRestoreToCrossRegionInstanceInput {
	s.NodeOperateType = &v
	return s
}

// SetNodeSpec sets the NodeSpec field's value.
func (s *NodeInfoForRestoreToCrossRegionInstanceInput) SetNodeSpec(v string) *NodeInfoForRestoreToCrossRegionInstanceInput {
	s.NodeSpec = &v
	return s
}

// SetNodeType sets the NodeType field's value.
func (s *NodeInfoForRestoreToCrossRegionInstanceInput) SetNodeType(v string) *NodeInfoForRestoreToCrossRegionInstanceInput {
	s.NodeType = &v
	return s
}

// SetZoneId sets the ZoneId field's value.
func (s *NodeInfoForRestoreToCrossRegionInstanceInput) SetZoneId(v string) *NodeInfoForRestoreToCrossRegionInstanceInput {
	s.ZoneId = &v
	return s
}

type RestoreToCrossRegionInstanceInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	AllowListIds []*string `type:"list" json:",omitempty"`

	BackupId *string `type:"string" json:",omitempty"`

	ChargeInfo *ChargeInfoForRestoreToCrossRegionInstanceInput `type:"structure" json:",omitempty"`

	DBParamGroupId *string `type:"string" json:",omitempty"`

	// DstRegionId is a required field
	DstRegionId *string `type:"string" json:",omitempty" required:"true"`

	InstanceName *string `type:"string" json:",omitempty"`

	InstanceTags []*InstanceTagForRestoreToCrossRegionInstanceInput `type:"list" json:",omitempty"`

	NodeInfo []*NodeInfoForRestoreToCrossRegionInstanceInput `type:"list" json:",omitempty"`

	Port *int32 `type:"int32" json:",omitempty"`

	ProjectName *string `type:"string" json:",omitempty"`

	RestoreTime *string `type:"string" json:",omitempty"`

	// SrcRegionId is a required field
	SrcRegionId *string `type:"string" json:",omitempty" required:"true"`

	// SrcRegionInstanceId is a required field
	SrcRegionInstanceId *string `type:"string" json:",omitempty" required:"true"`

	StorageSpace *int32 `type:"int32" json:",omitempty"`

	// StorageType is a required field
	StorageType *string `type:"string" json:",omitempty" required:"true" enum:"EnumOfStorageTypeForRestoreToCrossRegionInstanceInput"`

	// SubnetId is a required field
	SubnetId *string `type:"string" json:",omitempty" required:"true"`

	// VpcId is a required field
	VpcId *string `type:"string" json:",omitempty" required:"true"`
}

// String returns the string representation
func (s RestoreToCrossRegionInstanceInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s RestoreToCrossRegionInstanceInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RestoreToCrossRegionInstanceInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "RestoreToCrossRegionInstanceInput"}
	if s.DstRegionId == nil {
		invalidParams.Add(request.NewErrParamRequired("DstRegionId"))
	}
	if s.SrcRegionId == nil {
		invalidParams.Add(request.NewErrParamRequired("SrcRegionId"))
	}
	if s.SrcRegionInstanceId == nil {
		invalidParams.Add(request.NewErrParamRequired("SrcRegionInstanceId"))
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
func (s *RestoreToCrossRegionInstanceInput) SetAllowListIds(v []*string) *RestoreToCrossRegionInstanceInput {
	s.AllowListIds = v
	return s
}

// SetBackupId sets the BackupId field's value.
func (s *RestoreToCrossRegionInstanceInput) SetBackupId(v string) *RestoreToCrossRegionInstanceInput {
	s.BackupId = &v
	return s
}

// SetChargeInfo sets the ChargeInfo field's value.
func (s *RestoreToCrossRegionInstanceInput) SetChargeInfo(v *ChargeInfoForRestoreToCrossRegionInstanceInput) *RestoreToCrossRegionInstanceInput {
	s.ChargeInfo = v
	return s
}

// SetDBParamGroupId sets the DBParamGroupId field's value.
func (s *RestoreToCrossRegionInstanceInput) SetDBParamGroupId(v string) *RestoreToCrossRegionInstanceInput {
	s.DBParamGroupId = &v
	return s
}

// SetDstRegionId sets the DstRegionId field's value.
func (s *RestoreToCrossRegionInstanceInput) SetDstRegionId(v string) *RestoreToCrossRegionInstanceInput {
	s.DstRegionId = &v
	return s
}

// SetInstanceName sets the InstanceName field's value.
func (s *RestoreToCrossRegionInstanceInput) SetInstanceName(v string) *RestoreToCrossRegionInstanceInput {
	s.InstanceName = &v
	return s
}

// SetInstanceTags sets the InstanceTags field's value.
func (s *RestoreToCrossRegionInstanceInput) SetInstanceTags(v []*InstanceTagForRestoreToCrossRegionInstanceInput) *RestoreToCrossRegionInstanceInput {
	s.InstanceTags = v
	return s
}

// SetNodeInfo sets the NodeInfo field's value.
func (s *RestoreToCrossRegionInstanceInput) SetNodeInfo(v []*NodeInfoForRestoreToCrossRegionInstanceInput) *RestoreToCrossRegionInstanceInput {
	s.NodeInfo = v
	return s
}

// SetPort sets the Port field's value.
func (s *RestoreToCrossRegionInstanceInput) SetPort(v int32) *RestoreToCrossRegionInstanceInput {
	s.Port = &v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *RestoreToCrossRegionInstanceInput) SetProjectName(v string) *RestoreToCrossRegionInstanceInput {
	s.ProjectName = &v
	return s
}

// SetRestoreTime sets the RestoreTime field's value.
func (s *RestoreToCrossRegionInstanceInput) SetRestoreTime(v string) *RestoreToCrossRegionInstanceInput {
	s.RestoreTime = &v
	return s
}

// SetSrcRegionId sets the SrcRegionId field's value.
func (s *RestoreToCrossRegionInstanceInput) SetSrcRegionId(v string) *RestoreToCrossRegionInstanceInput {
	s.SrcRegionId = &v
	return s
}

// SetSrcRegionInstanceId sets the SrcRegionInstanceId field's value.
func (s *RestoreToCrossRegionInstanceInput) SetSrcRegionInstanceId(v string) *RestoreToCrossRegionInstanceInput {
	s.SrcRegionInstanceId = &v
	return s
}

// SetStorageSpace sets the StorageSpace field's value.
func (s *RestoreToCrossRegionInstanceInput) SetStorageSpace(v int32) *RestoreToCrossRegionInstanceInput {
	s.StorageSpace = &v
	return s
}

// SetStorageType sets the StorageType field's value.
func (s *RestoreToCrossRegionInstanceInput) SetStorageType(v string) *RestoreToCrossRegionInstanceInput {
	s.StorageType = &v
	return s
}

// SetSubnetId sets the SubnetId field's value.
func (s *RestoreToCrossRegionInstanceInput) SetSubnetId(v string) *RestoreToCrossRegionInstanceInput {
	s.SubnetId = &v
	return s
}

// SetVpcId sets the VpcId field's value.
func (s *RestoreToCrossRegionInstanceInput) SetVpcId(v string) *RestoreToCrossRegionInstanceInput {
	s.VpcId = &v
	return s
}

type RestoreToCrossRegionInstanceOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	InstanceId *string `type:"string" json:",omitempty"`

	OrderId *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s RestoreToCrossRegionInstanceOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s RestoreToCrossRegionInstanceOutput) GoString() string {
	return s.String()
}

// SetInstanceId sets the InstanceId field's value.
func (s *RestoreToCrossRegionInstanceOutput) SetInstanceId(v string) *RestoreToCrossRegionInstanceOutput {
	s.InstanceId = &v
	return s
}

// SetOrderId sets the OrderId field's value.
func (s *RestoreToCrossRegionInstanceOutput) SetOrderId(v string) *RestoreToCrossRegionInstanceOutput {
	s.OrderId = &v
	return s
}

const (
	// EnumOfStorageTypeForRestoreToCrossRegionInstanceInputLocalSsd is a EnumOfStorageTypeForRestoreToCrossRegionInstanceInput enum value
	EnumOfStorageTypeForRestoreToCrossRegionInstanceInputLocalSsd = "LocalSSD"

	// EnumOfStorageTypeForRestoreToCrossRegionInstanceInputCloudStorage is a EnumOfStorageTypeForRestoreToCrossRegionInstanceInput enum value
	EnumOfStorageTypeForRestoreToCrossRegionInstanceInputCloudStorage = "CloudStorage"

	// EnumOfStorageTypeForRestoreToCrossRegionInstanceInputEssdpl1 is a EnumOfStorageTypeForRestoreToCrossRegionInstanceInput enum value
	EnumOfStorageTypeForRestoreToCrossRegionInstanceInputEssdpl1 = "ESSDPL1"

	// EnumOfStorageTypeForRestoreToCrossRegionInstanceInputEssdpl2 is a EnumOfStorageTypeForRestoreToCrossRegionInstanceInput enum value
	EnumOfStorageTypeForRestoreToCrossRegionInstanceInputEssdpl2 = "ESSDPL2"
)

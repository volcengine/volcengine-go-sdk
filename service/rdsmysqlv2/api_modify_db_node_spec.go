// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package rdsmysqlv2

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opModifyDBNodeSpecCommon = "ModifyDBNodeSpec"

// ModifyDBNodeSpecCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ModifyDBNodeSpecCommon operation. The "output" return
// value will be populated with the ModifyDBNodeSpecCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyDBNodeSpecCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyDBNodeSpecCommon Send returns without error.
//
// See ModifyDBNodeSpecCommon for more information on using the ModifyDBNodeSpecCommon
// API call, and error handling.
//
//    // Example sending a request using the ModifyDBNodeSpecCommonRequest method.
//    req, resp := client.ModifyDBNodeSpecCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *RDSMYSQLV2) ModifyDBNodeSpecCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opModifyDBNodeSpecCommon,
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

// ModifyDBNodeSpecCommon API operation for RDS_MYSQL_V2.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RDS_MYSQL_V2's
// API operation ModifyDBNodeSpecCommon for usage and error information.
func (c *RDSMYSQLV2) ModifyDBNodeSpecCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ModifyDBNodeSpecCommonRequest(input)
	return out, req.Send()
}

// ModifyDBNodeSpecCommonWithContext is the same as ModifyDBNodeSpecCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyDBNodeSpecCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RDSMYSQLV2) ModifyDBNodeSpecCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ModifyDBNodeSpecCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opModifyDBNodeSpec = "ModifyDBNodeSpec"

// ModifyDBNodeSpecRequest generates a "volcengine/request.Request" representing the
// client's request for the ModifyDBNodeSpec operation. The "output" return
// value will be populated with the ModifyDBNodeSpecCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyDBNodeSpecCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyDBNodeSpecCommon Send returns without error.
//
// See ModifyDBNodeSpec for more information on using the ModifyDBNodeSpec
// API call, and error handling.
//
//    // Example sending a request using the ModifyDBNodeSpecRequest method.
//    req, resp := client.ModifyDBNodeSpecRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *RDSMYSQLV2) ModifyDBNodeSpecRequest(input *ModifyDBNodeSpecInput) (req *request.Request, output *ModifyDBNodeSpecOutput) {
	op := &request.Operation{
		Name:       opModifyDBNodeSpec,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ModifyDBNodeSpecInput{}
	}

	output = &ModifyDBNodeSpecOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// ModifyDBNodeSpec API operation for RDS_MYSQL_V2.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RDS_MYSQL_V2's
// API operation ModifyDBNodeSpec for usage and error information.
func (c *RDSMYSQLV2) ModifyDBNodeSpec(input *ModifyDBNodeSpecInput) (*ModifyDBNodeSpecOutput, error) {
	req, out := c.ModifyDBNodeSpecRequest(input)
	return out, req.Send()
}

// ModifyDBNodeSpecWithContext is the same as ModifyDBNodeSpec with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyDBNodeSpec for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RDSMYSQLV2) ModifyDBNodeSpecWithContext(ctx volcengine.Context, input *ModifyDBNodeSpecInput, opts ...request.Option) (*ModifyDBNodeSpecOutput, error) {
	req, out := c.ModifyDBNodeSpecRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type EstimationResultForModifyDBNodeSpecOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Effects []*string `type:"list" json:",omitempty"`

	Plans []*string `type:"list" json:",omitempty"`
}

// String returns the string representation
func (s EstimationResultForModifyDBNodeSpecOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s EstimationResultForModifyDBNodeSpecOutput) GoString() string {
	return s.String()
}

// SetEffects sets the Effects field's value.
func (s *EstimationResultForModifyDBNodeSpecOutput) SetEffects(v []*string) *EstimationResultForModifyDBNodeSpecOutput {
	s.Effects = v
	return s
}

// SetPlans sets the Plans field's value.
func (s *EstimationResultForModifyDBNodeSpecOutput) SetPlans(v []*string) *EstimationResultForModifyDBNodeSpecOutput {
	s.Plans = v
	return s
}

type ModifyDBNodeSpecInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	EstimateOnly *bool `type:"boolean" json:",omitempty"`

	// InstanceId is a required field
	InstanceId *string `type:"string" json:",omitempty" required:"true"`

	NodeInfo []*NodeInfoForModifyDBNodeSpecInput `type:"list" json:",omitempty"`

	RequestSource *string `type:"string" json:",omitempty"`

	SpecifiedSwitchEndTime *string `type:"string" json:",omitempty"`

	SpecifiedSwitchStartTime *string `type:"string" json:",omitempty"`

	StorageSpace *int32 `type:"int32" json:",omitempty"`

	// StorageType is a required field
	StorageType *string `type:"string" json:",omitempty" required:"true"`

	SwitchType *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s ModifyDBNodeSpecInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyDBNodeSpecInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ModifyDBNodeSpecInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ModifyDBNodeSpecInput"}
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

// SetEstimateOnly sets the EstimateOnly field's value.
func (s *ModifyDBNodeSpecInput) SetEstimateOnly(v bool) *ModifyDBNodeSpecInput {
	s.EstimateOnly = &v
	return s
}

// SetInstanceId sets the InstanceId field's value.
func (s *ModifyDBNodeSpecInput) SetInstanceId(v string) *ModifyDBNodeSpecInput {
	s.InstanceId = &v
	return s
}

// SetNodeInfo sets the NodeInfo field's value.
func (s *ModifyDBNodeSpecInput) SetNodeInfo(v []*NodeInfoForModifyDBNodeSpecInput) *ModifyDBNodeSpecInput {
	s.NodeInfo = v
	return s
}

// SetRequestSource sets the RequestSource field's value.
func (s *ModifyDBNodeSpecInput) SetRequestSource(v string) *ModifyDBNodeSpecInput {
	s.RequestSource = &v
	return s
}

// SetSpecifiedSwitchEndTime sets the SpecifiedSwitchEndTime field's value.
func (s *ModifyDBNodeSpecInput) SetSpecifiedSwitchEndTime(v string) *ModifyDBNodeSpecInput {
	s.SpecifiedSwitchEndTime = &v
	return s
}

// SetSpecifiedSwitchStartTime sets the SpecifiedSwitchStartTime field's value.
func (s *ModifyDBNodeSpecInput) SetSpecifiedSwitchStartTime(v string) *ModifyDBNodeSpecInput {
	s.SpecifiedSwitchStartTime = &v
	return s
}

// SetStorageSpace sets the StorageSpace field's value.
func (s *ModifyDBNodeSpecInput) SetStorageSpace(v int32) *ModifyDBNodeSpecInput {
	s.StorageSpace = &v
	return s
}

// SetStorageType sets the StorageType field's value.
func (s *ModifyDBNodeSpecInput) SetStorageType(v string) *ModifyDBNodeSpecInput {
	s.StorageType = &v
	return s
}

// SetSwitchType sets the SwitchType field's value.
func (s *ModifyDBNodeSpecInput) SetSwitchType(v string) *ModifyDBNodeSpecInput {
	s.SwitchType = &v
	return s
}

type ModifyDBNodeSpecOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	EstimationResult *EstimationResultForModifyDBNodeSpecOutput `type:"structure" json:",omitempty"`

	InstanceId *string `type:"string" json:",omitempty"`

	OrderId *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s ModifyDBNodeSpecOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyDBNodeSpecOutput) GoString() string {
	return s.String()
}

// SetEstimationResult sets the EstimationResult field's value.
func (s *ModifyDBNodeSpecOutput) SetEstimationResult(v *EstimationResultForModifyDBNodeSpecOutput) *ModifyDBNodeSpecOutput {
	s.EstimationResult = v
	return s
}

// SetInstanceId sets the InstanceId field's value.
func (s *ModifyDBNodeSpecOutput) SetInstanceId(v string) *ModifyDBNodeSpecOutput {
	s.InstanceId = &v
	return s
}

// SetOrderId sets the OrderId field's value.
func (s *ModifyDBNodeSpecOutput) SetOrderId(v string) *ModifyDBNodeSpecOutput {
	s.OrderId = &v
	return s
}

type NodeInfoForModifyDBNodeSpecInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	NodeId *string `type:"string" json:",omitempty"`

	NodeOperateType *string `type:"string" json:",omitempty"`

	NodeSpec *string `type:"string" json:",omitempty"`

	NodeType *string `type:"string" json:",omitempty"`

	ZoneId *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s NodeInfoForModifyDBNodeSpecInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s NodeInfoForModifyDBNodeSpecInput) GoString() string {
	return s.String()
}

// SetNodeId sets the NodeId field's value.
func (s *NodeInfoForModifyDBNodeSpecInput) SetNodeId(v string) *NodeInfoForModifyDBNodeSpecInput {
	s.NodeId = &v
	return s
}

// SetNodeOperateType sets the NodeOperateType field's value.
func (s *NodeInfoForModifyDBNodeSpecInput) SetNodeOperateType(v string) *NodeInfoForModifyDBNodeSpecInput {
	s.NodeOperateType = &v
	return s
}

// SetNodeSpec sets the NodeSpec field's value.
func (s *NodeInfoForModifyDBNodeSpecInput) SetNodeSpec(v string) *NodeInfoForModifyDBNodeSpecInput {
	s.NodeSpec = &v
	return s
}

// SetNodeType sets the NodeType field's value.
func (s *NodeInfoForModifyDBNodeSpecInput) SetNodeType(v string) *NodeInfoForModifyDBNodeSpecInput {
	s.NodeType = &v
	return s
}

// SetZoneId sets the ZoneId field's value.
func (s *NodeInfoForModifyDBNodeSpecInput) SetZoneId(v string) *NodeInfoForModifyDBNodeSpecInput {
	s.ZoneId = &v
	return s
}

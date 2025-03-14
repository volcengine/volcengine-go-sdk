// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package rdsmysqlv2

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opCreateDBNodesCommon = "CreateDBNodes"

// CreateDBNodesCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateDBNodesCommon operation. The "output" return
// value will be populated with the CreateDBNodesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateDBNodesCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateDBNodesCommon Send returns without error.
//
// See CreateDBNodesCommon for more information on using the CreateDBNodesCommon
// API call, and error handling.
//
//    // Example sending a request using the CreateDBNodesCommonRequest method.
//    req, resp := client.CreateDBNodesCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *RDSMYSQLV2) CreateDBNodesCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opCreateDBNodesCommon,
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

// CreateDBNodesCommon API operation for RDS_MYSQL_V2.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RDS_MYSQL_V2's
// API operation CreateDBNodesCommon for usage and error information.
func (c *RDSMYSQLV2) CreateDBNodesCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.CreateDBNodesCommonRequest(input)
	return out, req.Send()
}

// CreateDBNodesCommonWithContext is the same as CreateDBNodesCommon with the addition of
// the ability to pass a context and additional request options.
//
// See CreateDBNodesCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RDSMYSQLV2) CreateDBNodesCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.CreateDBNodesCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opCreateDBNodes = "CreateDBNodes"

// CreateDBNodesRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateDBNodes operation. The "output" return
// value will be populated with the CreateDBNodesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateDBNodesCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateDBNodesCommon Send returns without error.
//
// See CreateDBNodes for more information on using the CreateDBNodes
// API call, and error handling.
//
//    // Example sending a request using the CreateDBNodesRequest method.
//    req, resp := client.CreateDBNodesRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *RDSMYSQLV2) CreateDBNodesRequest(input *CreateDBNodesInput) (req *request.Request, output *CreateDBNodesOutput) {
	op := &request.Operation{
		Name:       opCreateDBNodes,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateDBNodesInput{}
	}

	output = &CreateDBNodesOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// CreateDBNodes API operation for RDS_MYSQL_V2.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RDS_MYSQL_V2's
// API operation CreateDBNodes for usage and error information.
func (c *RDSMYSQLV2) CreateDBNodes(input *CreateDBNodesInput) (*CreateDBNodesOutput, error) {
	req, out := c.CreateDBNodesRequest(input)
	return out, req.Send()
}

// CreateDBNodesWithContext is the same as CreateDBNodes with the addition of
// the ability to pass a context and additional request options.
//
// See CreateDBNodes for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RDSMYSQLV2) CreateDBNodesWithContext(ctx volcengine.Context, input *CreateDBNodesInput, opts ...request.Option) (*CreateDBNodesOutput, error) {
	req, out := c.CreateDBNodesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type CreateDBNodesInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	EstimateOnly *bool `type:"boolean" json:",omitempty"`

	// InstanceId is a required field
	InstanceId *string `type:"string" json:",omitempty" required:"true"`

	NodeInfo []*NodeInfoForCreateDBNodesInput `type:"list" json:",omitempty"`

	RequestSource *string `type:"string" json:",omitempty"`

	SpecifiedSwitchEndTime *string `type:"string" json:",omitempty"`

	SpecifiedSwitchStartTime *string `type:"string" json:",omitempty"`

	SwitchType *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s CreateDBNodesInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateDBNodesInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateDBNodesInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "CreateDBNodesInput"}
	if s.InstanceId == nil {
		invalidParams.Add(request.NewErrParamRequired("InstanceId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetEstimateOnly sets the EstimateOnly field's value.
func (s *CreateDBNodesInput) SetEstimateOnly(v bool) *CreateDBNodesInput {
	s.EstimateOnly = &v
	return s
}

// SetInstanceId sets the InstanceId field's value.
func (s *CreateDBNodesInput) SetInstanceId(v string) *CreateDBNodesInput {
	s.InstanceId = &v
	return s
}

// SetNodeInfo sets the NodeInfo field's value.
func (s *CreateDBNodesInput) SetNodeInfo(v []*NodeInfoForCreateDBNodesInput) *CreateDBNodesInput {
	s.NodeInfo = v
	return s
}

// SetRequestSource sets the RequestSource field's value.
func (s *CreateDBNodesInput) SetRequestSource(v string) *CreateDBNodesInput {
	s.RequestSource = &v
	return s
}

// SetSpecifiedSwitchEndTime sets the SpecifiedSwitchEndTime field's value.
func (s *CreateDBNodesInput) SetSpecifiedSwitchEndTime(v string) *CreateDBNodesInput {
	s.SpecifiedSwitchEndTime = &v
	return s
}

// SetSpecifiedSwitchStartTime sets the SpecifiedSwitchStartTime field's value.
func (s *CreateDBNodesInput) SetSpecifiedSwitchStartTime(v string) *CreateDBNodesInput {
	s.SpecifiedSwitchStartTime = &v
	return s
}

// SetSwitchType sets the SwitchType field's value.
func (s *CreateDBNodesInput) SetSwitchType(v string) *CreateDBNodesInput {
	s.SwitchType = &v
	return s
}

type CreateDBNodesOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	EstimationResult *EstimationResultForCreateDBNodesOutput `type:"structure" json:",omitempty"`

	InstanceId *string `type:"string" json:",omitempty"`

	OrderId *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s CreateDBNodesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateDBNodesOutput) GoString() string {
	return s.String()
}

// SetEstimationResult sets the EstimationResult field's value.
func (s *CreateDBNodesOutput) SetEstimationResult(v *EstimationResultForCreateDBNodesOutput) *CreateDBNodesOutput {
	s.EstimationResult = v
	return s
}

// SetInstanceId sets the InstanceId field's value.
func (s *CreateDBNodesOutput) SetInstanceId(v string) *CreateDBNodesOutput {
	s.InstanceId = &v
	return s
}

// SetOrderId sets the OrderId field's value.
func (s *CreateDBNodesOutput) SetOrderId(v string) *CreateDBNodesOutput {
	s.OrderId = &v
	return s
}

type EstimationResultForCreateDBNodesOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Effects []*string `type:"list" json:",omitempty"`

	Plans []*string `type:"list" json:",omitempty"`
}

// String returns the string representation
func (s EstimationResultForCreateDBNodesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s EstimationResultForCreateDBNodesOutput) GoString() string {
	return s.String()
}

// SetEffects sets the Effects field's value.
func (s *EstimationResultForCreateDBNodesOutput) SetEffects(v []*string) *EstimationResultForCreateDBNodesOutput {
	s.Effects = v
	return s
}

// SetPlans sets the Plans field's value.
func (s *EstimationResultForCreateDBNodesOutput) SetPlans(v []*string) *EstimationResultForCreateDBNodesOutput {
	s.Plans = v
	return s
}

type NodeInfoForCreateDBNodesInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	NodeId *string `type:"string" json:",omitempty"`

	NodeOperateType *string `type:"string" json:",omitempty"`

	NodeSpec *string `type:"string" json:",omitempty"`

	NodeType *string `type:"string" json:",omitempty"`

	ZoneId *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s NodeInfoForCreateDBNodesInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s NodeInfoForCreateDBNodesInput) GoString() string {
	return s.String()
}

// SetNodeId sets the NodeId field's value.
func (s *NodeInfoForCreateDBNodesInput) SetNodeId(v string) *NodeInfoForCreateDBNodesInput {
	s.NodeId = &v
	return s
}

// SetNodeOperateType sets the NodeOperateType field's value.
func (s *NodeInfoForCreateDBNodesInput) SetNodeOperateType(v string) *NodeInfoForCreateDBNodesInput {
	s.NodeOperateType = &v
	return s
}

// SetNodeSpec sets the NodeSpec field's value.
func (s *NodeInfoForCreateDBNodesInput) SetNodeSpec(v string) *NodeInfoForCreateDBNodesInput {
	s.NodeSpec = &v
	return s
}

// SetNodeType sets the NodeType field's value.
func (s *NodeInfoForCreateDBNodesInput) SetNodeType(v string) *NodeInfoForCreateDBNodesInput {
	s.NodeType = &v
	return s
}

// SetZoneId sets the ZoneId field's value.
func (s *NodeInfoForCreateDBNodesInput) SetZoneId(v string) *NodeInfoForCreateDBNodesInput {
	s.ZoneId = &v
	return s
}

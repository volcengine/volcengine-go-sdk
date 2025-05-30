// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package rdspostgresql

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribeDBInstancePriceDifferenceCommon = "DescribeDBInstancePriceDifference"

// DescribeDBInstancePriceDifferenceCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeDBInstancePriceDifferenceCommon operation. The "output" return
// value will be populated with the DescribeDBInstancePriceDifferenceCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeDBInstancePriceDifferenceCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeDBInstancePriceDifferenceCommon Send returns without error.
//
// See DescribeDBInstancePriceDifferenceCommon for more information on using the DescribeDBInstancePriceDifferenceCommon
// API call, and error handling.
//
//    // Example sending a request using the DescribeDBInstancePriceDifferenceCommonRequest method.
//    req, resp := client.DescribeDBInstancePriceDifferenceCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *RDSPOSTGRESQL) DescribeDBInstancePriceDifferenceCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeDBInstancePriceDifferenceCommon,
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

// DescribeDBInstancePriceDifferenceCommon API operation for RDS_POSTGRESQL.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RDS_POSTGRESQL's
// API operation DescribeDBInstancePriceDifferenceCommon for usage and error information.
func (c *RDSPOSTGRESQL) DescribeDBInstancePriceDifferenceCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeDBInstancePriceDifferenceCommonRequest(input)
	return out, req.Send()
}

// DescribeDBInstancePriceDifferenceCommonWithContext is the same as DescribeDBInstancePriceDifferenceCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeDBInstancePriceDifferenceCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RDSPOSTGRESQL) DescribeDBInstancePriceDifferenceCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeDBInstancePriceDifferenceCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeDBInstancePriceDifference = "DescribeDBInstancePriceDifference"

// DescribeDBInstancePriceDifferenceRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeDBInstancePriceDifference operation. The "output" return
// value will be populated with the DescribeDBInstancePriceDifferenceCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeDBInstancePriceDifferenceCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeDBInstancePriceDifferenceCommon Send returns without error.
//
// See DescribeDBInstancePriceDifference for more information on using the DescribeDBInstancePriceDifference
// API call, and error handling.
//
//    // Example sending a request using the DescribeDBInstancePriceDifferenceRequest method.
//    req, resp := client.DescribeDBInstancePriceDifferenceRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *RDSPOSTGRESQL) DescribeDBInstancePriceDifferenceRequest(input *DescribeDBInstancePriceDifferenceInput) (req *request.Request, output *DescribeDBInstancePriceDifferenceOutput) {
	op := &request.Operation{
		Name:       opDescribeDBInstancePriceDifference,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeDBInstancePriceDifferenceInput{}
	}

	output = &DescribeDBInstancePriceDifferenceOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// DescribeDBInstancePriceDifference API operation for RDS_POSTGRESQL.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RDS_POSTGRESQL's
// API operation DescribeDBInstancePriceDifference for usage and error information.
func (c *RDSPOSTGRESQL) DescribeDBInstancePriceDifference(input *DescribeDBInstancePriceDifferenceInput) (*DescribeDBInstancePriceDifferenceOutput, error) {
	req, out := c.DescribeDBInstancePriceDifferenceRequest(input)
	return out, req.Send()
}

// DescribeDBInstancePriceDifferenceWithContext is the same as DescribeDBInstancePriceDifference with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeDBInstancePriceDifference for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RDSPOSTGRESQL) DescribeDBInstancePriceDifferenceWithContext(ctx volcengine.Context, input *DescribeDBInstancePriceDifferenceInput, opts ...request.Option) (*DescribeDBInstancePriceDifferenceOutput, error) {
	req, out := c.DescribeDBInstancePriceDifferenceRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ChargeInfoForDescribeDBInstancePriceDifferenceInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	AutoRenew *bool `type:"boolean" json:",omitempty"`

	ChargeType *string `type:"string" json:",omitempty"`

	Number *int32 `type:"int32" json:",omitempty"`

	Period *int32 `type:"int32" json:",omitempty"`

	PeriodUnit *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s ChargeInfoForDescribeDBInstancePriceDifferenceInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ChargeInfoForDescribeDBInstancePriceDifferenceInput) GoString() string {
	return s.String()
}

// SetAutoRenew sets the AutoRenew field's value.
func (s *ChargeInfoForDescribeDBInstancePriceDifferenceInput) SetAutoRenew(v bool) *ChargeInfoForDescribeDBInstancePriceDifferenceInput {
	s.AutoRenew = &v
	return s
}

// SetChargeType sets the ChargeType field's value.
func (s *ChargeInfoForDescribeDBInstancePriceDifferenceInput) SetChargeType(v string) *ChargeInfoForDescribeDBInstancePriceDifferenceInput {
	s.ChargeType = &v
	return s
}

// SetNumber sets the Number field's value.
func (s *ChargeInfoForDescribeDBInstancePriceDifferenceInput) SetNumber(v int32) *ChargeInfoForDescribeDBInstancePriceDifferenceInput {
	s.Number = &v
	return s
}

// SetPeriod sets the Period field's value.
func (s *ChargeInfoForDescribeDBInstancePriceDifferenceInput) SetPeriod(v int32) *ChargeInfoForDescribeDBInstancePriceDifferenceInput {
	s.Period = &v
	return s
}

// SetPeriodUnit sets the PeriodUnit field's value.
func (s *ChargeInfoForDescribeDBInstancePriceDifferenceInput) SetPeriodUnit(v string) *ChargeInfoForDescribeDBInstancePriceDifferenceInput {
	s.PeriodUnit = &v
	return s
}

type ChargeItemPriceForDescribeDBInstancePriceDifferenceOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	ChargeItemKey *string `type:"string" json:",omitempty"`

	ChargeItemType *string `type:"string" json:",omitempty"`

	ChargeItemValue *int32 `type:"int32" json:",omitempty"`

	DiscountPrice *float64 `type:"double" json:",omitempty"`

	NodeNumPerInstance *int32 `type:"int32" json:",omitempty"`

	OriginalPrice *float64 `type:"double" json:",omitempty"`

	PayablePrice *float64 `type:"double" json:",omitempty"`
}

// String returns the string representation
func (s ChargeItemPriceForDescribeDBInstancePriceDifferenceOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ChargeItemPriceForDescribeDBInstancePriceDifferenceOutput) GoString() string {
	return s.String()
}

// SetChargeItemKey sets the ChargeItemKey field's value.
func (s *ChargeItemPriceForDescribeDBInstancePriceDifferenceOutput) SetChargeItemKey(v string) *ChargeItemPriceForDescribeDBInstancePriceDifferenceOutput {
	s.ChargeItemKey = &v
	return s
}

// SetChargeItemType sets the ChargeItemType field's value.
func (s *ChargeItemPriceForDescribeDBInstancePriceDifferenceOutput) SetChargeItemType(v string) *ChargeItemPriceForDescribeDBInstancePriceDifferenceOutput {
	s.ChargeItemType = &v
	return s
}

// SetChargeItemValue sets the ChargeItemValue field's value.
func (s *ChargeItemPriceForDescribeDBInstancePriceDifferenceOutput) SetChargeItemValue(v int32) *ChargeItemPriceForDescribeDBInstancePriceDifferenceOutput {
	s.ChargeItemValue = &v
	return s
}

// SetDiscountPrice sets the DiscountPrice field's value.
func (s *ChargeItemPriceForDescribeDBInstancePriceDifferenceOutput) SetDiscountPrice(v float64) *ChargeItemPriceForDescribeDBInstancePriceDifferenceOutput {
	s.DiscountPrice = &v
	return s
}

// SetNodeNumPerInstance sets the NodeNumPerInstance field's value.
func (s *ChargeItemPriceForDescribeDBInstancePriceDifferenceOutput) SetNodeNumPerInstance(v int32) *ChargeItemPriceForDescribeDBInstancePriceDifferenceOutput {
	s.NodeNumPerInstance = &v
	return s
}

// SetOriginalPrice sets the OriginalPrice field's value.
func (s *ChargeItemPriceForDescribeDBInstancePriceDifferenceOutput) SetOriginalPrice(v float64) *ChargeItemPriceForDescribeDBInstancePriceDifferenceOutput {
	s.OriginalPrice = &v
	return s
}

// SetPayablePrice sets the PayablePrice field's value.
func (s *ChargeItemPriceForDescribeDBInstancePriceDifferenceOutput) SetPayablePrice(v float64) *ChargeItemPriceForDescribeDBInstancePriceDifferenceOutput {
	s.PayablePrice = &v
	return s
}

type DescribeDBInstancePriceDifferenceInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	ChargeInfo *ChargeInfoForDescribeDBInstancePriceDifferenceInput `type:"structure" json:",omitempty"`

	// InstanceId is a required field
	InstanceId *string `type:"string" json:",omitempty" required:"true"`

	ModifyType *string `type:"string" json:",omitempty"`

	NodeInfo []*NodeInfoForDescribeDBInstancePriceDifferenceInput `type:"list" json:",omitempty"`

	RollbackTime *string `type:"string" json:",omitempty"`

	// StorageSpace is a required field
	StorageSpace *int32 `type:"int32" json:",omitempty" required:"true"`

	// StorageType is a required field
	StorageType *string `type:"string" json:",omitempty" required:"true"`
}

// String returns the string representation
func (s DescribeDBInstancePriceDifferenceInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeDBInstancePriceDifferenceInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeDBInstancePriceDifferenceInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DescribeDBInstancePriceDifferenceInput"}
	if s.InstanceId == nil {
		invalidParams.Add(request.NewErrParamRequired("InstanceId"))
	}
	if s.StorageSpace == nil {
		invalidParams.Add(request.NewErrParamRequired("StorageSpace"))
	}
	if s.StorageType == nil {
		invalidParams.Add(request.NewErrParamRequired("StorageType"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetChargeInfo sets the ChargeInfo field's value.
func (s *DescribeDBInstancePriceDifferenceInput) SetChargeInfo(v *ChargeInfoForDescribeDBInstancePriceDifferenceInput) *DescribeDBInstancePriceDifferenceInput {
	s.ChargeInfo = v
	return s
}

// SetInstanceId sets the InstanceId field's value.
func (s *DescribeDBInstancePriceDifferenceInput) SetInstanceId(v string) *DescribeDBInstancePriceDifferenceInput {
	s.InstanceId = &v
	return s
}

// SetModifyType sets the ModifyType field's value.
func (s *DescribeDBInstancePriceDifferenceInput) SetModifyType(v string) *DescribeDBInstancePriceDifferenceInput {
	s.ModifyType = &v
	return s
}

// SetNodeInfo sets the NodeInfo field's value.
func (s *DescribeDBInstancePriceDifferenceInput) SetNodeInfo(v []*NodeInfoForDescribeDBInstancePriceDifferenceInput) *DescribeDBInstancePriceDifferenceInput {
	s.NodeInfo = v
	return s
}

// SetRollbackTime sets the RollbackTime field's value.
func (s *DescribeDBInstancePriceDifferenceInput) SetRollbackTime(v string) *DescribeDBInstancePriceDifferenceInput {
	s.RollbackTime = &v
	return s
}

// SetStorageSpace sets the StorageSpace field's value.
func (s *DescribeDBInstancePriceDifferenceInput) SetStorageSpace(v int32) *DescribeDBInstancePriceDifferenceInput {
	s.StorageSpace = &v
	return s
}

// SetStorageType sets the StorageType field's value.
func (s *DescribeDBInstancePriceDifferenceInput) SetStorageType(v string) *DescribeDBInstancePriceDifferenceInput {
	s.StorageType = &v
	return s
}

type DescribeDBInstancePriceDifferenceOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	ChargeItemPrices []*ChargeItemPriceForDescribeDBInstancePriceDifferenceOutput `type:"list" json:",omitempty"`

	Currency *string `type:"string" json:",omitempty"`

	DiscountPrice *float64 `type:"double" json:",omitempty"`

	OriginalPrice *float64 `type:"double" json:",omitempty"`

	PayablePrice *float64 `type:"double" json:",omitempty"`
}

// String returns the string representation
func (s DescribeDBInstancePriceDifferenceOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeDBInstancePriceDifferenceOutput) GoString() string {
	return s.String()
}

// SetChargeItemPrices sets the ChargeItemPrices field's value.
func (s *DescribeDBInstancePriceDifferenceOutput) SetChargeItemPrices(v []*ChargeItemPriceForDescribeDBInstancePriceDifferenceOutput) *DescribeDBInstancePriceDifferenceOutput {
	s.ChargeItemPrices = v
	return s
}

// SetCurrency sets the Currency field's value.
func (s *DescribeDBInstancePriceDifferenceOutput) SetCurrency(v string) *DescribeDBInstancePriceDifferenceOutput {
	s.Currency = &v
	return s
}

// SetDiscountPrice sets the DiscountPrice field's value.
func (s *DescribeDBInstancePriceDifferenceOutput) SetDiscountPrice(v float64) *DescribeDBInstancePriceDifferenceOutput {
	s.DiscountPrice = &v
	return s
}

// SetOriginalPrice sets the OriginalPrice field's value.
func (s *DescribeDBInstancePriceDifferenceOutput) SetOriginalPrice(v float64) *DescribeDBInstancePriceDifferenceOutput {
	s.OriginalPrice = &v
	return s
}

// SetPayablePrice sets the PayablePrice field's value.
func (s *DescribeDBInstancePriceDifferenceOutput) SetPayablePrice(v float64) *DescribeDBInstancePriceDifferenceOutput {
	s.PayablePrice = &v
	return s
}

type NodeInfoForDescribeDBInstancePriceDifferenceInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	NodeId *string `type:"string" json:",omitempty"`

	NodeOperateType *string `type:"string" json:",omitempty"`

	NodeSpec *string `type:"string" json:",omitempty"`

	NodeType *string `type:"string" json:",omitempty"`

	ZoneId *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s NodeInfoForDescribeDBInstancePriceDifferenceInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s NodeInfoForDescribeDBInstancePriceDifferenceInput) GoString() string {
	return s.String()
}

// SetNodeId sets the NodeId field's value.
func (s *NodeInfoForDescribeDBInstancePriceDifferenceInput) SetNodeId(v string) *NodeInfoForDescribeDBInstancePriceDifferenceInput {
	s.NodeId = &v
	return s
}

// SetNodeOperateType sets the NodeOperateType field's value.
func (s *NodeInfoForDescribeDBInstancePriceDifferenceInput) SetNodeOperateType(v string) *NodeInfoForDescribeDBInstancePriceDifferenceInput {
	s.NodeOperateType = &v
	return s
}

// SetNodeSpec sets the NodeSpec field's value.
func (s *NodeInfoForDescribeDBInstancePriceDifferenceInput) SetNodeSpec(v string) *NodeInfoForDescribeDBInstancePriceDifferenceInput {
	s.NodeSpec = &v
	return s
}

// SetNodeType sets the NodeType field's value.
func (s *NodeInfoForDescribeDBInstancePriceDifferenceInput) SetNodeType(v string) *NodeInfoForDescribeDBInstancePriceDifferenceInput {
	s.NodeType = &v
	return s
}

// SetZoneId sets the ZoneId field's value.
func (s *NodeInfoForDescribeDBInstancePriceDifferenceInput) SetZoneId(v string) *NodeInfoForDescribeDBInstancePriceDifferenceInput {
	s.ZoneId = &v
	return s
}

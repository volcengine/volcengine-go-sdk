// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package rdspostgresql

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribeDBInstancePriceDetailCommon = "DescribeDBInstancePriceDetail"

// DescribeDBInstancePriceDetailCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeDBInstancePriceDetailCommon operation. The "output" return
// value will be populated with the DescribeDBInstancePriceDetailCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeDBInstancePriceDetailCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeDBInstancePriceDetailCommon Send returns without error.
//
// See DescribeDBInstancePriceDetailCommon for more information on using the DescribeDBInstancePriceDetailCommon
// API call, and error handling.
//
//    // Example sending a request using the DescribeDBInstancePriceDetailCommonRequest method.
//    req, resp := client.DescribeDBInstancePriceDetailCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *RDSPOSTGRESQL) DescribeDBInstancePriceDetailCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeDBInstancePriceDetailCommon,
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

// DescribeDBInstancePriceDetailCommon API operation for RDS_POSTGRESQL.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RDS_POSTGRESQL's
// API operation DescribeDBInstancePriceDetailCommon for usage and error information.
func (c *RDSPOSTGRESQL) DescribeDBInstancePriceDetailCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeDBInstancePriceDetailCommonRequest(input)
	return out, req.Send()
}

// DescribeDBInstancePriceDetailCommonWithContext is the same as DescribeDBInstancePriceDetailCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeDBInstancePriceDetailCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RDSPOSTGRESQL) DescribeDBInstancePriceDetailCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeDBInstancePriceDetailCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeDBInstancePriceDetail = "DescribeDBInstancePriceDetail"

// DescribeDBInstancePriceDetailRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeDBInstancePriceDetail operation. The "output" return
// value will be populated with the DescribeDBInstancePriceDetailCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeDBInstancePriceDetailCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeDBInstancePriceDetailCommon Send returns without error.
//
// See DescribeDBInstancePriceDetail for more information on using the DescribeDBInstancePriceDetail
// API call, and error handling.
//
//    // Example sending a request using the DescribeDBInstancePriceDetailRequest method.
//    req, resp := client.DescribeDBInstancePriceDetailRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *RDSPOSTGRESQL) DescribeDBInstancePriceDetailRequest(input *DescribeDBInstancePriceDetailInput) (req *request.Request, output *DescribeDBInstancePriceDetailOutput) {
	op := &request.Operation{
		Name:       opDescribeDBInstancePriceDetail,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeDBInstancePriceDetailInput{}
	}

	output = &DescribeDBInstancePriceDetailOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// DescribeDBInstancePriceDetail API operation for RDS_POSTGRESQL.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RDS_POSTGRESQL's
// API operation DescribeDBInstancePriceDetail for usage and error information.
func (c *RDSPOSTGRESQL) DescribeDBInstancePriceDetail(input *DescribeDBInstancePriceDetailInput) (*DescribeDBInstancePriceDetailOutput, error) {
	req, out := c.DescribeDBInstancePriceDetailRequest(input)
	return out, req.Send()
}

// DescribeDBInstancePriceDetailWithContext is the same as DescribeDBInstancePriceDetail with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeDBInstancePriceDetail for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RDSPOSTGRESQL) DescribeDBInstancePriceDetailWithContext(ctx volcengine.Context, input *DescribeDBInstancePriceDetailInput, opts ...request.Option) (*DescribeDBInstancePriceDetailOutput, error) {
	req, out := c.DescribeDBInstancePriceDetailRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ChargeInfoForDescribeDBInstancePriceDetailInput struct {
	_ struct{} `type:"structure"`

	AutoRenew *bool `type:"boolean"`

	ChargeType *string `type:"string" enum:"EnumOfChargeTypeForDescribeDBInstancePriceDetailInput"`

	Number *int32 `type:"int32"`

	Period *int32 `type:"int32"`

	PeriodUnit *string `type:"string" enum:"EnumOfPeriodUnitForDescribeDBInstancePriceDetailInput"`
}

// String returns the string representation
func (s ChargeInfoForDescribeDBInstancePriceDetailInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ChargeInfoForDescribeDBInstancePriceDetailInput) GoString() string {
	return s.String()
}

// SetAutoRenew sets the AutoRenew field's value.
func (s *ChargeInfoForDescribeDBInstancePriceDetailInput) SetAutoRenew(v bool) *ChargeInfoForDescribeDBInstancePriceDetailInput {
	s.AutoRenew = &v
	return s
}

// SetChargeType sets the ChargeType field's value.
func (s *ChargeInfoForDescribeDBInstancePriceDetailInput) SetChargeType(v string) *ChargeInfoForDescribeDBInstancePriceDetailInput {
	s.ChargeType = &v
	return s
}

// SetNumber sets the Number field's value.
func (s *ChargeInfoForDescribeDBInstancePriceDetailInput) SetNumber(v int32) *ChargeInfoForDescribeDBInstancePriceDetailInput {
	s.Number = &v
	return s
}

// SetPeriod sets the Period field's value.
func (s *ChargeInfoForDescribeDBInstancePriceDetailInput) SetPeriod(v int32) *ChargeInfoForDescribeDBInstancePriceDetailInput {
	s.Period = &v
	return s
}

// SetPeriodUnit sets the PeriodUnit field's value.
func (s *ChargeInfoForDescribeDBInstancePriceDetailInput) SetPeriodUnit(v string) *ChargeInfoForDescribeDBInstancePriceDetailInput {
	s.PeriodUnit = &v
	return s
}

type ChargeItemPriceForDescribeDBInstancePriceDetailOutput struct {
	_ struct{} `type:"structure"`

	ChargeItemKey *string `type:"string"`

	ChargeItemType *string `type:"string" enum:"EnumOfChargeItemTypeForDescribeDBInstancePriceDetailOutput"`

	ChargeItemValue *int32 `type:"int32"`

	DiscountPrice *float64 `type:"double"`

	NodeNumPerInstance *int32 `type:"int32"`

	OriginalPrice *float64 `type:"double"`

	PayablePrice *float64 `type:"double"`
}

// String returns the string representation
func (s ChargeItemPriceForDescribeDBInstancePriceDetailOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ChargeItemPriceForDescribeDBInstancePriceDetailOutput) GoString() string {
	return s.String()
}

// SetChargeItemKey sets the ChargeItemKey field's value.
func (s *ChargeItemPriceForDescribeDBInstancePriceDetailOutput) SetChargeItemKey(v string) *ChargeItemPriceForDescribeDBInstancePriceDetailOutput {
	s.ChargeItemKey = &v
	return s
}

// SetChargeItemType sets the ChargeItemType field's value.
func (s *ChargeItemPriceForDescribeDBInstancePriceDetailOutput) SetChargeItemType(v string) *ChargeItemPriceForDescribeDBInstancePriceDetailOutput {
	s.ChargeItemType = &v
	return s
}

// SetChargeItemValue sets the ChargeItemValue field's value.
func (s *ChargeItemPriceForDescribeDBInstancePriceDetailOutput) SetChargeItemValue(v int32) *ChargeItemPriceForDescribeDBInstancePriceDetailOutput {
	s.ChargeItemValue = &v
	return s
}

// SetDiscountPrice sets the DiscountPrice field's value.
func (s *ChargeItemPriceForDescribeDBInstancePriceDetailOutput) SetDiscountPrice(v float64) *ChargeItemPriceForDescribeDBInstancePriceDetailOutput {
	s.DiscountPrice = &v
	return s
}

// SetNodeNumPerInstance sets the NodeNumPerInstance field's value.
func (s *ChargeItemPriceForDescribeDBInstancePriceDetailOutput) SetNodeNumPerInstance(v int32) *ChargeItemPriceForDescribeDBInstancePriceDetailOutput {
	s.NodeNumPerInstance = &v
	return s
}

// SetOriginalPrice sets the OriginalPrice field's value.
func (s *ChargeItemPriceForDescribeDBInstancePriceDetailOutput) SetOriginalPrice(v float64) *ChargeItemPriceForDescribeDBInstancePriceDetailOutput {
	s.OriginalPrice = &v
	return s
}

// SetPayablePrice sets the PayablePrice field's value.
func (s *ChargeItemPriceForDescribeDBInstancePriceDetailOutput) SetPayablePrice(v float64) *ChargeItemPriceForDescribeDBInstancePriceDetailOutput {
	s.PayablePrice = &v
	return s
}

type DescribeDBInstancePriceDetailInput struct {
	_ struct{} `type:"structure"`

	ChargeInfo *ChargeInfoForDescribeDBInstancePriceDetailInput `type:"structure"`

	NodeInfo []*NodeInfoForDescribeDBInstancePriceDetailInput `type:"list"`

	// StorageSpace is a required field
	StorageSpace *int32 `min:"20" max:"3000" type:"int32" required:"true"`

	// StorageType is a required field
	StorageType *string `type:"string" required:"true" enum:"EnumOfStorageTypeForDescribeDBInstancePriceDetailInput"`
}

// String returns the string representation
func (s DescribeDBInstancePriceDetailInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeDBInstancePriceDetailInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeDBInstancePriceDetailInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DescribeDBInstancePriceDetailInput"}
	if s.StorageSpace == nil {
		invalidParams.Add(request.NewErrParamRequired("StorageSpace"))
	}
	if s.StorageSpace != nil && *s.StorageSpace < 20 {
		invalidParams.Add(request.NewErrParamMinValue("StorageSpace", 20))
	}
	if s.StorageSpace != nil && *s.StorageSpace > 3000 {
		invalidParams.Add(request.NewErrParamMaxValue("StorageSpace", 3000))
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
func (s *DescribeDBInstancePriceDetailInput) SetChargeInfo(v *ChargeInfoForDescribeDBInstancePriceDetailInput) *DescribeDBInstancePriceDetailInput {
	s.ChargeInfo = v
	return s
}

// SetNodeInfo sets the NodeInfo field's value.
func (s *DescribeDBInstancePriceDetailInput) SetNodeInfo(v []*NodeInfoForDescribeDBInstancePriceDetailInput) *DescribeDBInstancePriceDetailInput {
	s.NodeInfo = v
	return s
}

// SetStorageSpace sets the StorageSpace field's value.
func (s *DescribeDBInstancePriceDetailInput) SetStorageSpace(v int32) *DescribeDBInstancePriceDetailInput {
	s.StorageSpace = &v
	return s
}

// SetStorageType sets the StorageType field's value.
func (s *DescribeDBInstancePriceDetailInput) SetStorageType(v string) *DescribeDBInstancePriceDetailInput {
	s.StorageType = &v
	return s
}

type DescribeDBInstancePriceDetailOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	ChargeItemPrices []*ChargeItemPriceForDescribeDBInstancePriceDetailOutput `type:"list"`

	Currency *string `type:"string"`

	DiscountPrice *float64 `type:"double"`

	InstanceQuantity *int32 `type:"int32"`

	OriginalPrice *float64 `type:"double"`

	PayablePrice *float64 `type:"double"`
}

// String returns the string representation
func (s DescribeDBInstancePriceDetailOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeDBInstancePriceDetailOutput) GoString() string {
	return s.String()
}

// SetChargeItemPrices sets the ChargeItemPrices field's value.
func (s *DescribeDBInstancePriceDetailOutput) SetChargeItemPrices(v []*ChargeItemPriceForDescribeDBInstancePriceDetailOutput) *DescribeDBInstancePriceDetailOutput {
	s.ChargeItemPrices = v
	return s
}

// SetCurrency sets the Currency field's value.
func (s *DescribeDBInstancePriceDetailOutput) SetCurrency(v string) *DescribeDBInstancePriceDetailOutput {
	s.Currency = &v
	return s
}

// SetDiscountPrice sets the DiscountPrice field's value.
func (s *DescribeDBInstancePriceDetailOutput) SetDiscountPrice(v float64) *DescribeDBInstancePriceDetailOutput {
	s.DiscountPrice = &v
	return s
}

// SetInstanceQuantity sets the InstanceQuantity field's value.
func (s *DescribeDBInstancePriceDetailOutput) SetInstanceQuantity(v int32) *DescribeDBInstancePriceDetailOutput {
	s.InstanceQuantity = &v
	return s
}

// SetOriginalPrice sets the OriginalPrice field's value.
func (s *DescribeDBInstancePriceDetailOutput) SetOriginalPrice(v float64) *DescribeDBInstancePriceDetailOutput {
	s.OriginalPrice = &v
	return s
}

// SetPayablePrice sets the PayablePrice field's value.
func (s *DescribeDBInstancePriceDetailOutput) SetPayablePrice(v float64) *DescribeDBInstancePriceDetailOutput {
	s.PayablePrice = &v
	return s
}

type NodeInfoForDescribeDBInstancePriceDetailInput struct {
	_ struct{} `type:"structure"`

	NodeId *string `type:"string"`

	NodeOperateType *string `type:"string" enum:"EnumOfNodeOperateTypeForDescribeDBInstancePriceDetailInput"`

	NodeSpec *string `type:"string"`

	NodeType *string `type:"string" enum:"EnumOfNodeTypeForDescribeDBInstancePriceDetailInput"`

	ZoneId *string `type:"string"`
}

// String returns the string representation
func (s NodeInfoForDescribeDBInstancePriceDetailInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s NodeInfoForDescribeDBInstancePriceDetailInput) GoString() string {
	return s.String()
}

// SetNodeId sets the NodeId field's value.
func (s *NodeInfoForDescribeDBInstancePriceDetailInput) SetNodeId(v string) *NodeInfoForDescribeDBInstancePriceDetailInput {
	s.NodeId = &v
	return s
}

// SetNodeOperateType sets the NodeOperateType field's value.
func (s *NodeInfoForDescribeDBInstancePriceDetailInput) SetNodeOperateType(v string) *NodeInfoForDescribeDBInstancePriceDetailInput {
	s.NodeOperateType = &v
	return s
}

// SetNodeSpec sets the NodeSpec field's value.
func (s *NodeInfoForDescribeDBInstancePriceDetailInput) SetNodeSpec(v string) *NodeInfoForDescribeDBInstancePriceDetailInput {
	s.NodeSpec = &v
	return s
}

// SetNodeType sets the NodeType field's value.
func (s *NodeInfoForDescribeDBInstancePriceDetailInput) SetNodeType(v string) *NodeInfoForDescribeDBInstancePriceDetailInput {
	s.NodeType = &v
	return s
}

// SetZoneId sets the ZoneId field's value.
func (s *NodeInfoForDescribeDBInstancePriceDetailInput) SetZoneId(v string) *NodeInfoForDescribeDBInstancePriceDetailInput {
	s.ZoneId = &v
	return s
}

const (
	// EnumOfChargeItemTypeForDescribeDBInstancePriceDetailOutputPrimary is a EnumOfChargeItemTypeForDescribeDBInstancePriceDetailOutput enum value
	EnumOfChargeItemTypeForDescribeDBInstancePriceDetailOutputPrimary = "Primary"

	// EnumOfChargeItemTypeForDescribeDBInstancePriceDetailOutputSecondary is a EnumOfChargeItemTypeForDescribeDBInstancePriceDetailOutput enum value
	EnumOfChargeItemTypeForDescribeDBInstancePriceDetailOutputSecondary = "Secondary"

	// EnumOfChargeItemTypeForDescribeDBInstancePriceDetailOutputReadOnly is a EnumOfChargeItemTypeForDescribeDBInstancePriceDetailOutput enum value
	EnumOfChargeItemTypeForDescribeDBInstancePriceDetailOutputReadOnly = "ReadOnly"

	// EnumOfChargeItemTypeForDescribeDBInstancePriceDetailOutputStorage is a EnumOfChargeItemTypeForDescribeDBInstancePriceDetailOutput enum value
	EnumOfChargeItemTypeForDescribeDBInstancePriceDetailOutputStorage = "Storage"
)

const (
	// EnumOfChargeTypeForDescribeDBInstancePriceDetailInputPostPaid is a EnumOfChargeTypeForDescribeDBInstancePriceDetailInput enum value
	EnumOfChargeTypeForDescribeDBInstancePriceDetailInputPostPaid = "PostPaid"

	// EnumOfChargeTypeForDescribeDBInstancePriceDetailInputPrePaid is a EnumOfChargeTypeForDescribeDBInstancePriceDetailInput enum value
	EnumOfChargeTypeForDescribeDBInstancePriceDetailInputPrePaid = "PrePaid"
)

const (
	// EnumOfNodeOperateTypeForDescribeDBInstancePriceDetailInputCreate is a EnumOfNodeOperateTypeForDescribeDBInstancePriceDetailInput enum value
	EnumOfNodeOperateTypeForDescribeDBInstancePriceDetailInputCreate = "Create"

	// EnumOfNodeOperateTypeForDescribeDBInstancePriceDetailInputDelete is a EnumOfNodeOperateTypeForDescribeDBInstancePriceDetailInput enum value
	EnumOfNodeOperateTypeForDescribeDBInstancePriceDetailInputDelete = "Delete"

	// EnumOfNodeOperateTypeForDescribeDBInstancePriceDetailInputModify is a EnumOfNodeOperateTypeForDescribeDBInstancePriceDetailInput enum value
	EnumOfNodeOperateTypeForDescribeDBInstancePriceDetailInputModify = "Modify"
)

const (
	// EnumOfNodeTypeForDescribeDBInstancePriceDetailInputPrimary is a EnumOfNodeTypeForDescribeDBInstancePriceDetailInput enum value
	EnumOfNodeTypeForDescribeDBInstancePriceDetailInputPrimary = "Primary"

	// EnumOfNodeTypeForDescribeDBInstancePriceDetailInputSecondary is a EnumOfNodeTypeForDescribeDBInstancePriceDetailInput enum value
	EnumOfNodeTypeForDescribeDBInstancePriceDetailInputSecondary = "Secondary"

	// EnumOfNodeTypeForDescribeDBInstancePriceDetailInputReadOnly is a EnumOfNodeTypeForDescribeDBInstancePriceDetailInput enum value
	EnumOfNodeTypeForDescribeDBInstancePriceDetailInputReadOnly = "ReadOnly"
)

const (
	// EnumOfPeriodUnitForDescribeDBInstancePriceDetailInputMonth is a EnumOfPeriodUnitForDescribeDBInstancePriceDetailInput enum value
	EnumOfPeriodUnitForDescribeDBInstancePriceDetailInputMonth = "Month"

	// EnumOfPeriodUnitForDescribeDBInstancePriceDetailInputYear is a EnumOfPeriodUnitForDescribeDBInstancePriceDetailInput enum value
	EnumOfPeriodUnitForDescribeDBInstancePriceDetailInputYear = "Year"
)

const (
	// EnumOfStorageTypeForDescribeDBInstancePriceDetailInputLocalSsd is a EnumOfStorageTypeForDescribeDBInstancePriceDetailInput enum value
	EnumOfStorageTypeForDescribeDBInstancePriceDetailInputLocalSsd = "LocalSSD"
)
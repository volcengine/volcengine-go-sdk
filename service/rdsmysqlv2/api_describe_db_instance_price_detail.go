// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package rdsmysqlv2

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
func (c *RDSMYSQLV2) DescribeDBInstancePriceDetailCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
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

// DescribeDBInstancePriceDetailCommon API operation for RDS_MYSQL_V2.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RDS_MYSQL_V2's
// API operation DescribeDBInstancePriceDetailCommon for usage and error information.
func (c *RDSMYSQLV2) DescribeDBInstancePriceDetailCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
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
func (c *RDSMYSQLV2) DescribeDBInstancePriceDetailCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
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
func (c *RDSMYSQLV2) DescribeDBInstancePriceDetailRequest(input *DescribeDBInstancePriceDetailInput) (req *request.Request, output *DescribeDBInstancePriceDetailOutput) {
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

// DescribeDBInstancePriceDetail API operation for RDS_MYSQL_V2.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RDS_MYSQL_V2's
// API operation DescribeDBInstancePriceDetail for usage and error information.
func (c *RDSMYSQLV2) DescribeDBInstancePriceDetail(input *DescribeDBInstancePriceDetailInput) (*DescribeDBInstancePriceDetailOutput, error) {
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
func (c *RDSMYSQLV2) DescribeDBInstancePriceDetailWithContext(ctx volcengine.Context, input *DescribeDBInstancePriceDetailInput, opts ...request.Option) (*DescribeDBInstancePriceDetailOutput, error) {
	req, out := c.DescribeDBInstancePriceDetailRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ChargeItemPriceForDescribeDBInstancePriceDetailOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	ChargeItemKey *string `type:"string" json:",omitempty"`

	ChargeItemType *string `type:"string" json:",omitempty"`

	ChargeItemValue *int64 `type:"int64" json:",omitempty"`

	DiscountPrice *string `type:"string" json:",omitempty"`

	OriginalPrice *string `type:"string" json:",omitempty"`

	PayablePrice *string `type:"string" json:",omitempty"`
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
func (s *ChargeItemPriceForDescribeDBInstancePriceDetailOutput) SetChargeItemValue(v int64) *ChargeItemPriceForDescribeDBInstancePriceDetailOutput {
	s.ChargeItemValue = &v
	return s
}

// SetDiscountPrice sets the DiscountPrice field's value.
func (s *ChargeItemPriceForDescribeDBInstancePriceDetailOutput) SetDiscountPrice(v string) *ChargeItemPriceForDescribeDBInstancePriceDetailOutput {
	s.DiscountPrice = &v
	return s
}

// SetOriginalPrice sets the OriginalPrice field's value.
func (s *ChargeItemPriceForDescribeDBInstancePriceDetailOutput) SetOriginalPrice(v string) *ChargeItemPriceForDescribeDBInstancePriceDetailOutput {
	s.OriginalPrice = &v
	return s
}

// SetPayablePrice sets the PayablePrice field's value.
func (s *ChargeItemPriceForDescribeDBInstancePriceDetailOutput) SetPayablePrice(v string) *ChargeItemPriceForDescribeDBInstancePriceDetailOutput {
	s.PayablePrice = &v
	return s
}

type DescribeDBInstancePriceDetailInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	// ChargeType is a required field
	ChargeType *string `type:"string" json:",omitempty" required:"true"`

	NodeInfo []*NodeInfoForDescribeDBInstancePriceDetailInput `type:"list" json:",omitempty"`

	Number *int32 `type:"int32" json:",omitempty"`

	Period *int32 `type:"int32" json:",omitempty"`

	PeriodUnit *string `type:"string" json:",omitempty"`

	ProjectName *string `type:"string" json:",omitempty"`

	ProxyNodeCustom *ProxyNodeCustomForDescribeDBInstancePriceDetailInput `type:"structure" json:",omitempty"`

	// StorageSpace is a required field
	StorageSpace *int32 `type:"int32" json:",omitempty" required:"true"`

	// StorageType is a required field
	StorageType *string `type:"string" json:",omitempty" required:"true"`
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
	if s.ChargeType == nil {
		invalidParams.Add(request.NewErrParamRequired("ChargeType"))
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

// SetChargeType sets the ChargeType field's value.
func (s *DescribeDBInstancePriceDetailInput) SetChargeType(v string) *DescribeDBInstancePriceDetailInput {
	s.ChargeType = &v
	return s
}

// SetNodeInfo sets the NodeInfo field's value.
func (s *DescribeDBInstancePriceDetailInput) SetNodeInfo(v []*NodeInfoForDescribeDBInstancePriceDetailInput) *DescribeDBInstancePriceDetailInput {
	s.NodeInfo = v
	return s
}

// SetNumber sets the Number field's value.
func (s *DescribeDBInstancePriceDetailInput) SetNumber(v int32) *DescribeDBInstancePriceDetailInput {
	s.Number = &v
	return s
}

// SetPeriod sets the Period field's value.
func (s *DescribeDBInstancePriceDetailInput) SetPeriod(v int32) *DescribeDBInstancePriceDetailInput {
	s.Period = &v
	return s
}

// SetPeriodUnit sets the PeriodUnit field's value.
func (s *DescribeDBInstancePriceDetailInput) SetPeriodUnit(v string) *DescribeDBInstancePriceDetailInput {
	s.PeriodUnit = &v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *DescribeDBInstancePriceDetailInput) SetProjectName(v string) *DescribeDBInstancePriceDetailInput {
	s.ProjectName = &v
	return s
}

// SetProxyNodeCustom sets the ProxyNodeCustom field's value.
func (s *DescribeDBInstancePriceDetailInput) SetProxyNodeCustom(v *ProxyNodeCustomForDescribeDBInstancePriceDetailInput) *DescribeDBInstancePriceDetailInput {
	s.ProxyNodeCustom = v
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
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	DescribeDBInstancePriceDetailStr *DescribeDBInstancePriceDetailStrForDescribeDBInstancePriceDetailOutput `type:"structure" json:",omitempty"`
}

// String returns the string representation
func (s DescribeDBInstancePriceDetailOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeDBInstancePriceDetailOutput) GoString() string {
	return s.String()
}

// SetDescribeDBInstancePriceDetailStr sets the DescribeDBInstancePriceDetailStr field's value.
func (s *DescribeDBInstancePriceDetailOutput) SetDescribeDBInstancePriceDetailStr(v *DescribeDBInstancePriceDetailStrForDescribeDBInstancePriceDetailOutput) *DescribeDBInstancePriceDetailOutput {
	s.DescribeDBInstancePriceDetailStr = v
	return s
}

type DescribeDBInstancePriceDetailStrForDescribeDBInstancePriceDetailOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	BillingMethod *string `type:"string" json:",omitempty"`

	ChargeItemPrices []*ChargeItemPriceForDescribeDBInstancePriceDetailOutput `type:"list" json:",omitempty"`

	CouponAmount *string `type:"string" json:",omitempty"`

	Currency *string `type:"string" json:",omitempty"`

	DiscountPrice *string `type:"string" json:",omitempty"`

	OriginalPrice *string `type:"string" json:",omitempty"`

	PayablePrice *string `type:"string" json:",omitempty"`

	Quantity *int32 `type:"int32" json:",omitempty"`

	RefundAmount *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s DescribeDBInstancePriceDetailStrForDescribeDBInstancePriceDetailOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeDBInstancePriceDetailStrForDescribeDBInstancePriceDetailOutput) GoString() string {
	return s.String()
}

// SetBillingMethod sets the BillingMethod field's value.
func (s *DescribeDBInstancePriceDetailStrForDescribeDBInstancePriceDetailOutput) SetBillingMethod(v string) *DescribeDBInstancePriceDetailStrForDescribeDBInstancePriceDetailOutput {
	s.BillingMethod = &v
	return s
}

// SetChargeItemPrices sets the ChargeItemPrices field's value.
func (s *DescribeDBInstancePriceDetailStrForDescribeDBInstancePriceDetailOutput) SetChargeItemPrices(v []*ChargeItemPriceForDescribeDBInstancePriceDetailOutput) *DescribeDBInstancePriceDetailStrForDescribeDBInstancePriceDetailOutput {
	s.ChargeItemPrices = v
	return s
}

// SetCouponAmount sets the CouponAmount field's value.
func (s *DescribeDBInstancePriceDetailStrForDescribeDBInstancePriceDetailOutput) SetCouponAmount(v string) *DescribeDBInstancePriceDetailStrForDescribeDBInstancePriceDetailOutput {
	s.CouponAmount = &v
	return s
}

// SetCurrency sets the Currency field's value.
func (s *DescribeDBInstancePriceDetailStrForDescribeDBInstancePriceDetailOutput) SetCurrency(v string) *DescribeDBInstancePriceDetailStrForDescribeDBInstancePriceDetailOutput {
	s.Currency = &v
	return s
}

// SetDiscountPrice sets the DiscountPrice field's value.
func (s *DescribeDBInstancePriceDetailStrForDescribeDBInstancePriceDetailOutput) SetDiscountPrice(v string) *DescribeDBInstancePriceDetailStrForDescribeDBInstancePriceDetailOutput {
	s.DiscountPrice = &v
	return s
}

// SetOriginalPrice sets the OriginalPrice field's value.
func (s *DescribeDBInstancePriceDetailStrForDescribeDBInstancePriceDetailOutput) SetOriginalPrice(v string) *DescribeDBInstancePriceDetailStrForDescribeDBInstancePriceDetailOutput {
	s.OriginalPrice = &v
	return s
}

// SetPayablePrice sets the PayablePrice field's value.
func (s *DescribeDBInstancePriceDetailStrForDescribeDBInstancePriceDetailOutput) SetPayablePrice(v string) *DescribeDBInstancePriceDetailStrForDescribeDBInstancePriceDetailOutput {
	s.PayablePrice = &v
	return s
}

// SetQuantity sets the Quantity field's value.
func (s *DescribeDBInstancePriceDetailStrForDescribeDBInstancePriceDetailOutput) SetQuantity(v int32) *DescribeDBInstancePriceDetailStrForDescribeDBInstancePriceDetailOutput {
	s.Quantity = &v
	return s
}

// SetRefundAmount sets the RefundAmount field's value.
func (s *DescribeDBInstancePriceDetailStrForDescribeDBInstancePriceDetailOutput) SetRefundAmount(v string) *DescribeDBInstancePriceDetailStrForDescribeDBInstancePriceDetailOutput {
	s.RefundAmount = &v
	return s
}

type NodeInfoForDescribeDBInstancePriceDetailInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	NodeId *string `type:"string" json:",omitempty"`

	NodeOperateType *string `type:"string" json:",omitempty"`

	NodeSpec *string `type:"string" json:",omitempty"`

	NodeType *string `type:"string" json:",omitempty"`

	ZoneId *string `type:"string" json:",omitempty"`
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

type ProxyNodeCustomForDescribeDBInstancePriceDetailInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	CpuNum *int32 `type:"int32" json:",omitempty"`
}

// String returns the string representation
func (s ProxyNodeCustomForDescribeDBInstancePriceDetailInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ProxyNodeCustomForDescribeDBInstancePriceDetailInput) GoString() string {
	return s.String()
}

// SetCpuNum sets the CpuNum field's value.
func (s *ProxyNodeCustomForDescribeDBInstancePriceDetailInput) SetCpuNum(v int32) *ProxyNodeCustomForDescribeDBInstancePriceDetailInput {
	s.CpuNum = &v
	return s
}

// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package billing

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opListSplitBillDetailCommon = "ListSplitBillDetail"

// ListSplitBillDetailCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ListSplitBillDetailCommon operation. The "output" return
// value will be populated with the ListSplitBillDetailCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListSplitBillDetailCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListSplitBillDetailCommon Send returns without error.
//
// See ListSplitBillDetailCommon for more information on using the ListSplitBillDetailCommon
// API call, and error handling.
//
//    // Example sending a request using the ListSplitBillDetailCommonRequest method.
//    req, resp := client.ListSplitBillDetailCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *BILLING) ListSplitBillDetailCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opListSplitBillDetailCommon,
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

// ListSplitBillDetailCommon API operation for BILLING.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for BILLING's
// API operation ListSplitBillDetailCommon for usage and error information.
func (c *BILLING) ListSplitBillDetailCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ListSplitBillDetailCommonRequest(input)
	return out, req.Send()
}

// ListSplitBillDetailCommonWithContext is the same as ListSplitBillDetailCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ListSplitBillDetailCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *BILLING) ListSplitBillDetailCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ListSplitBillDetailCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opListSplitBillDetail = "ListSplitBillDetail"

// ListSplitBillDetailRequest generates a "volcengine/request.Request" representing the
// client's request for the ListSplitBillDetail operation. The "output" return
// value will be populated with the ListSplitBillDetailCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListSplitBillDetailCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListSplitBillDetailCommon Send returns without error.
//
// See ListSplitBillDetail for more information on using the ListSplitBillDetail
// API call, and error handling.
//
//    // Example sending a request using the ListSplitBillDetailRequest method.
//    req, resp := client.ListSplitBillDetailRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *BILLING) ListSplitBillDetailRequest(input *ListSplitBillDetailInput) (req *request.Request, output *ListSplitBillDetailOutput) {
	op := &request.Operation{
		Name:       opListSplitBillDetail,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListSplitBillDetailInput{}
	}

	output = &ListSplitBillDetailOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// ListSplitBillDetail API operation for BILLING.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for BILLING's
// API operation ListSplitBillDetail for usage and error information.
func (c *BILLING) ListSplitBillDetail(input *ListSplitBillDetailInput) (*ListSplitBillDetailOutput, error) {
	req, out := c.ListSplitBillDetailRequest(input)
	return out, req.Send()
}

// ListSplitBillDetailWithContext is the same as ListSplitBillDetail with the addition of
// the ability to pass a context and additional request options.
//
// See ListSplitBillDetail for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *BILLING) ListSplitBillDetailWithContext(ctx volcengine.Context, input *ListSplitBillDetailInput, opts ...request.Option) (*ListSplitBillDetailOutput, error) {
	req, out := c.ListSplitBillDetailRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ListForListSplitBillDetailOutput struct {
	_ struct{} `type:"structure"`

	BillCategory *string `type:"string"`

	BillID *string `type:"string"`

	BillPeriod *string `type:"string"`

	BillingFunction *string `type:"string"`

	BillingMethodCode *string `type:"string"`

	BillingMode *string `type:"string"`

	BusinessMode *string `type:"string"`

	ConfigName *string `type:"string"`

	ConfigurationCode *string `type:"string"`

	CouponDeductionAmount *string `type:"string"`

	CreditCarriedAmount *string `type:"string"`

	Currency *string `type:"string"`

	DeductionCount *string `type:"string"`

	DiscountBillAmount *string `type:"string"`

	DiscountBizBillingFunction *string `type:"string"`

	DiscountBizMeasureInterval *string `type:"string"`

	DiscountBizUnitPrice *string `type:"string"`

	DiscountBizUnitPriceInterval *string `type:"string"`

	EffectiveFactor *string `type:"string"`

	Element *string `type:"string"`

	ElementCode *string `type:"string"`

	ExpandField *string `type:"string"`

	ExpenseTime *string `type:"string"`

	Factor *string `type:"string"`

	FactorCode *string `type:"string"`

	InstanceName *string `type:"string"`

	InstanceNo *string `type:"string"`

	MarketPrice *string `type:"string"`

	MeasureInterval *string `type:"string"`

	OriginalBillAmount *string `type:"string"`

	OwnerUserName *string `type:"string"`

	PaidAmount *string `type:"string"`

	PayableAmount *string `type:"string"`

	PayerUserName *string `type:"string"`

	PreferentialBillAmount *string `type:"string"`

	Price *string `type:"string"`

	PriceInterval *string `type:"string"`

	PriceUnit *string `type:"string"`

	Product *string `type:"string"`

	ProductZh *string `type:"string"`

	Project *string `type:"string"`

	ProjectDisplayName *string `type:"string"`

	Region *string `type:"string"`

	RegionCode *string `type:"string"`

	ReservationInstance *string `type:"string"`

	SellerUserName *string `type:"string"`

	SellingMode *string `type:"string"`

	SettlementType *string `type:"string"`

	SolutionZh *string `type:"string"`

	SplitBillDetailId *string `type:"string"`

	SplitItemAmount *string `type:"string"`

	SplitItemID *string `type:"string"`

	SplitItemName *string `type:"string"`

	SplitItemRatio *string `type:"string"`

	SubjectName *string `type:"string"`

	Tag *string `type:"string"`

	TradeTime *string `type:"string"`

	Unit *string `type:"string"`

	UnpaidAmount *string `type:"string"`

	UseDuration *string `type:"string"`

	UseDurationUnit *string `type:"string"`

	Zone *string `type:"string"`

	ZoneCode *string `type:"string"`
}

// String returns the string representation
func (s ListForListSplitBillDetailOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListForListSplitBillDetailOutput) GoString() string {
	return s.String()
}

// SetBillCategory sets the BillCategory field's value.
func (s *ListForListSplitBillDetailOutput) SetBillCategory(v string) *ListForListSplitBillDetailOutput {
	s.BillCategory = &v
	return s
}

// SetBillID sets the BillID field's value.
func (s *ListForListSplitBillDetailOutput) SetBillID(v string) *ListForListSplitBillDetailOutput {
	s.BillID = &v
	return s
}

// SetBillPeriod sets the BillPeriod field's value.
func (s *ListForListSplitBillDetailOutput) SetBillPeriod(v string) *ListForListSplitBillDetailOutput {
	s.BillPeriod = &v
	return s
}

// SetBillingFunction sets the BillingFunction field's value.
func (s *ListForListSplitBillDetailOutput) SetBillingFunction(v string) *ListForListSplitBillDetailOutput {
	s.BillingFunction = &v
	return s
}

// SetBillingMethodCode sets the BillingMethodCode field's value.
func (s *ListForListSplitBillDetailOutput) SetBillingMethodCode(v string) *ListForListSplitBillDetailOutput {
	s.BillingMethodCode = &v
	return s
}

// SetBillingMode sets the BillingMode field's value.
func (s *ListForListSplitBillDetailOutput) SetBillingMode(v string) *ListForListSplitBillDetailOutput {
	s.BillingMode = &v
	return s
}

// SetBusinessMode sets the BusinessMode field's value.
func (s *ListForListSplitBillDetailOutput) SetBusinessMode(v string) *ListForListSplitBillDetailOutput {
	s.BusinessMode = &v
	return s
}

// SetConfigName sets the ConfigName field's value.
func (s *ListForListSplitBillDetailOutput) SetConfigName(v string) *ListForListSplitBillDetailOutput {
	s.ConfigName = &v
	return s
}

// SetConfigurationCode sets the ConfigurationCode field's value.
func (s *ListForListSplitBillDetailOutput) SetConfigurationCode(v string) *ListForListSplitBillDetailOutput {
	s.ConfigurationCode = &v
	return s
}

// SetCouponDeductionAmount sets the CouponDeductionAmount field's value.
func (s *ListForListSplitBillDetailOutput) SetCouponDeductionAmount(v string) *ListForListSplitBillDetailOutput {
	s.CouponDeductionAmount = &v
	return s
}

// SetCreditCarriedAmount sets the CreditCarriedAmount field's value.
func (s *ListForListSplitBillDetailOutput) SetCreditCarriedAmount(v string) *ListForListSplitBillDetailOutput {
	s.CreditCarriedAmount = &v
	return s
}

// SetCurrency sets the Currency field's value.
func (s *ListForListSplitBillDetailOutput) SetCurrency(v string) *ListForListSplitBillDetailOutput {
	s.Currency = &v
	return s
}

// SetDeductionCount sets the DeductionCount field's value.
func (s *ListForListSplitBillDetailOutput) SetDeductionCount(v string) *ListForListSplitBillDetailOutput {
	s.DeductionCount = &v
	return s
}

// SetDiscountBillAmount sets the DiscountBillAmount field's value.
func (s *ListForListSplitBillDetailOutput) SetDiscountBillAmount(v string) *ListForListSplitBillDetailOutput {
	s.DiscountBillAmount = &v
	return s
}

// SetDiscountBizBillingFunction sets the DiscountBizBillingFunction field's value.
func (s *ListForListSplitBillDetailOutput) SetDiscountBizBillingFunction(v string) *ListForListSplitBillDetailOutput {
	s.DiscountBizBillingFunction = &v
	return s
}

// SetDiscountBizMeasureInterval sets the DiscountBizMeasureInterval field's value.
func (s *ListForListSplitBillDetailOutput) SetDiscountBizMeasureInterval(v string) *ListForListSplitBillDetailOutput {
	s.DiscountBizMeasureInterval = &v
	return s
}

// SetDiscountBizUnitPrice sets the DiscountBizUnitPrice field's value.
func (s *ListForListSplitBillDetailOutput) SetDiscountBizUnitPrice(v string) *ListForListSplitBillDetailOutput {
	s.DiscountBizUnitPrice = &v
	return s
}

// SetDiscountBizUnitPriceInterval sets the DiscountBizUnitPriceInterval field's value.
func (s *ListForListSplitBillDetailOutput) SetDiscountBizUnitPriceInterval(v string) *ListForListSplitBillDetailOutput {
	s.DiscountBizUnitPriceInterval = &v
	return s
}

// SetEffectiveFactor sets the EffectiveFactor field's value.
func (s *ListForListSplitBillDetailOutput) SetEffectiveFactor(v string) *ListForListSplitBillDetailOutput {
	s.EffectiveFactor = &v
	return s
}

// SetElement sets the Element field's value.
func (s *ListForListSplitBillDetailOutput) SetElement(v string) *ListForListSplitBillDetailOutput {
	s.Element = &v
	return s
}

// SetElementCode sets the ElementCode field's value.
func (s *ListForListSplitBillDetailOutput) SetElementCode(v string) *ListForListSplitBillDetailOutput {
	s.ElementCode = &v
	return s
}

// SetExpandField sets the ExpandField field's value.
func (s *ListForListSplitBillDetailOutput) SetExpandField(v string) *ListForListSplitBillDetailOutput {
	s.ExpandField = &v
	return s
}

// SetExpenseTime sets the ExpenseTime field's value.
func (s *ListForListSplitBillDetailOutput) SetExpenseTime(v string) *ListForListSplitBillDetailOutput {
	s.ExpenseTime = &v
	return s
}

// SetFactor sets the Factor field's value.
func (s *ListForListSplitBillDetailOutput) SetFactor(v string) *ListForListSplitBillDetailOutput {
	s.Factor = &v
	return s
}

// SetFactorCode sets the FactorCode field's value.
func (s *ListForListSplitBillDetailOutput) SetFactorCode(v string) *ListForListSplitBillDetailOutput {
	s.FactorCode = &v
	return s
}

// SetInstanceName sets the InstanceName field's value.
func (s *ListForListSplitBillDetailOutput) SetInstanceName(v string) *ListForListSplitBillDetailOutput {
	s.InstanceName = &v
	return s
}

// SetInstanceNo sets the InstanceNo field's value.
func (s *ListForListSplitBillDetailOutput) SetInstanceNo(v string) *ListForListSplitBillDetailOutput {
	s.InstanceNo = &v
	return s
}

// SetMarketPrice sets the MarketPrice field's value.
func (s *ListForListSplitBillDetailOutput) SetMarketPrice(v string) *ListForListSplitBillDetailOutput {
	s.MarketPrice = &v
	return s
}

// SetMeasureInterval sets the MeasureInterval field's value.
func (s *ListForListSplitBillDetailOutput) SetMeasureInterval(v string) *ListForListSplitBillDetailOutput {
	s.MeasureInterval = &v
	return s
}

// SetOriginalBillAmount sets the OriginalBillAmount field's value.
func (s *ListForListSplitBillDetailOutput) SetOriginalBillAmount(v string) *ListForListSplitBillDetailOutput {
	s.OriginalBillAmount = &v
	return s
}

// SetOwnerUserName sets the OwnerUserName field's value.
func (s *ListForListSplitBillDetailOutput) SetOwnerUserName(v string) *ListForListSplitBillDetailOutput {
	s.OwnerUserName = &v
	return s
}

// SetPaidAmount sets the PaidAmount field's value.
func (s *ListForListSplitBillDetailOutput) SetPaidAmount(v string) *ListForListSplitBillDetailOutput {
	s.PaidAmount = &v
	return s
}

// SetPayableAmount sets the PayableAmount field's value.
func (s *ListForListSplitBillDetailOutput) SetPayableAmount(v string) *ListForListSplitBillDetailOutput {
	s.PayableAmount = &v
	return s
}

// SetPayerUserName sets the PayerUserName field's value.
func (s *ListForListSplitBillDetailOutput) SetPayerUserName(v string) *ListForListSplitBillDetailOutput {
	s.PayerUserName = &v
	return s
}

// SetPreferentialBillAmount sets the PreferentialBillAmount field's value.
func (s *ListForListSplitBillDetailOutput) SetPreferentialBillAmount(v string) *ListForListSplitBillDetailOutput {
	s.PreferentialBillAmount = &v
	return s
}

// SetPrice sets the Price field's value.
func (s *ListForListSplitBillDetailOutput) SetPrice(v string) *ListForListSplitBillDetailOutput {
	s.Price = &v
	return s
}

// SetPriceInterval sets the PriceInterval field's value.
func (s *ListForListSplitBillDetailOutput) SetPriceInterval(v string) *ListForListSplitBillDetailOutput {
	s.PriceInterval = &v
	return s
}

// SetPriceUnit sets the PriceUnit field's value.
func (s *ListForListSplitBillDetailOutput) SetPriceUnit(v string) *ListForListSplitBillDetailOutput {
	s.PriceUnit = &v
	return s
}

// SetProduct sets the Product field's value.
func (s *ListForListSplitBillDetailOutput) SetProduct(v string) *ListForListSplitBillDetailOutput {
	s.Product = &v
	return s
}

// SetProductZh sets the ProductZh field's value.
func (s *ListForListSplitBillDetailOutput) SetProductZh(v string) *ListForListSplitBillDetailOutput {
	s.ProductZh = &v
	return s
}

// SetProject sets the Project field's value.
func (s *ListForListSplitBillDetailOutput) SetProject(v string) *ListForListSplitBillDetailOutput {
	s.Project = &v
	return s
}

// SetProjectDisplayName sets the ProjectDisplayName field's value.
func (s *ListForListSplitBillDetailOutput) SetProjectDisplayName(v string) *ListForListSplitBillDetailOutput {
	s.ProjectDisplayName = &v
	return s
}

// SetRegion sets the Region field's value.
func (s *ListForListSplitBillDetailOutput) SetRegion(v string) *ListForListSplitBillDetailOutput {
	s.Region = &v
	return s
}

// SetRegionCode sets the RegionCode field's value.
func (s *ListForListSplitBillDetailOutput) SetRegionCode(v string) *ListForListSplitBillDetailOutput {
	s.RegionCode = &v
	return s
}

// SetReservationInstance sets the ReservationInstance field's value.
func (s *ListForListSplitBillDetailOutput) SetReservationInstance(v string) *ListForListSplitBillDetailOutput {
	s.ReservationInstance = &v
	return s
}

// SetSellerUserName sets the SellerUserName field's value.
func (s *ListForListSplitBillDetailOutput) SetSellerUserName(v string) *ListForListSplitBillDetailOutput {
	s.SellerUserName = &v
	return s
}

// SetSellingMode sets the SellingMode field's value.
func (s *ListForListSplitBillDetailOutput) SetSellingMode(v string) *ListForListSplitBillDetailOutput {
	s.SellingMode = &v
	return s
}

// SetSettlementType sets the SettlementType field's value.
func (s *ListForListSplitBillDetailOutput) SetSettlementType(v string) *ListForListSplitBillDetailOutput {
	s.SettlementType = &v
	return s
}

// SetSolutionZh sets the SolutionZh field's value.
func (s *ListForListSplitBillDetailOutput) SetSolutionZh(v string) *ListForListSplitBillDetailOutput {
	s.SolutionZh = &v
	return s
}

// SetSplitBillDetailId sets the SplitBillDetailId field's value.
func (s *ListForListSplitBillDetailOutput) SetSplitBillDetailId(v string) *ListForListSplitBillDetailOutput {
	s.SplitBillDetailId = &v
	return s
}

// SetSplitItemAmount sets the SplitItemAmount field's value.
func (s *ListForListSplitBillDetailOutput) SetSplitItemAmount(v string) *ListForListSplitBillDetailOutput {
	s.SplitItemAmount = &v
	return s
}

// SetSplitItemID sets the SplitItemID field's value.
func (s *ListForListSplitBillDetailOutput) SetSplitItemID(v string) *ListForListSplitBillDetailOutput {
	s.SplitItemID = &v
	return s
}

// SetSplitItemName sets the SplitItemName field's value.
func (s *ListForListSplitBillDetailOutput) SetSplitItemName(v string) *ListForListSplitBillDetailOutput {
	s.SplitItemName = &v
	return s
}

// SetSplitItemRatio sets the SplitItemRatio field's value.
func (s *ListForListSplitBillDetailOutput) SetSplitItemRatio(v string) *ListForListSplitBillDetailOutput {
	s.SplitItemRatio = &v
	return s
}

// SetSubjectName sets the SubjectName field's value.
func (s *ListForListSplitBillDetailOutput) SetSubjectName(v string) *ListForListSplitBillDetailOutput {
	s.SubjectName = &v
	return s
}

// SetTag sets the Tag field's value.
func (s *ListForListSplitBillDetailOutput) SetTag(v string) *ListForListSplitBillDetailOutput {
	s.Tag = &v
	return s
}

// SetTradeTime sets the TradeTime field's value.
func (s *ListForListSplitBillDetailOutput) SetTradeTime(v string) *ListForListSplitBillDetailOutput {
	s.TradeTime = &v
	return s
}

// SetUnit sets the Unit field's value.
func (s *ListForListSplitBillDetailOutput) SetUnit(v string) *ListForListSplitBillDetailOutput {
	s.Unit = &v
	return s
}

// SetUnpaidAmount sets the UnpaidAmount field's value.
func (s *ListForListSplitBillDetailOutput) SetUnpaidAmount(v string) *ListForListSplitBillDetailOutput {
	s.UnpaidAmount = &v
	return s
}

// SetUseDuration sets the UseDuration field's value.
func (s *ListForListSplitBillDetailOutput) SetUseDuration(v string) *ListForListSplitBillDetailOutput {
	s.UseDuration = &v
	return s
}

// SetUseDurationUnit sets the UseDurationUnit field's value.
func (s *ListForListSplitBillDetailOutput) SetUseDurationUnit(v string) *ListForListSplitBillDetailOutput {
	s.UseDurationUnit = &v
	return s
}

// SetZone sets the Zone field's value.
func (s *ListForListSplitBillDetailOutput) SetZone(v string) *ListForListSplitBillDetailOutput {
	s.Zone = &v
	return s
}

// SetZoneCode sets the ZoneCode field's value.
func (s *ListForListSplitBillDetailOutput) SetZoneCode(v string) *ListForListSplitBillDetailOutput {
	s.ZoneCode = &v
	return s
}

type ListSplitBillDetailInput struct {
	_ struct{} `type:"structure"`

	BillCategory []*string `type:"list"`

	// BillPeriod is a required field
	BillPeriod *string `type:"string" required:"true"`

	BillingMode []*string `type:"list"`

	ExpenseDate *string `type:"string"`

	GroupPeriod *int32 `type:"int32"`

	IgnoreZero *int32 `type:"int32"`

	InstanceNo *string `type:"string"`

	// Limit is a required field
	Limit *int32 `type:"int32" required:"true"`

	NeedRecordNum *int32 `type:"int32"`

	Offset *int32 `type:"int32"`

	OwnerID []*int64 `type:"list"`

	PayerID []*int64 `type:"list"`

	Product []*string `type:"list"`

	SplitItemID *string `type:"string"`
}

// String returns the string representation
func (s ListSplitBillDetailInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListSplitBillDetailInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListSplitBillDetailInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ListSplitBillDetailInput"}
	if s.BillPeriod == nil {
		invalidParams.Add(request.NewErrParamRequired("BillPeriod"))
	}
	if s.Limit == nil {
		invalidParams.Add(request.NewErrParamRequired("Limit"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetBillCategory sets the BillCategory field's value.
func (s *ListSplitBillDetailInput) SetBillCategory(v []*string) *ListSplitBillDetailInput {
	s.BillCategory = v
	return s
}

// SetBillPeriod sets the BillPeriod field's value.
func (s *ListSplitBillDetailInput) SetBillPeriod(v string) *ListSplitBillDetailInput {
	s.BillPeriod = &v
	return s
}

// SetBillingMode sets the BillingMode field's value.
func (s *ListSplitBillDetailInput) SetBillingMode(v []*string) *ListSplitBillDetailInput {
	s.BillingMode = v
	return s
}

// SetExpenseDate sets the ExpenseDate field's value.
func (s *ListSplitBillDetailInput) SetExpenseDate(v string) *ListSplitBillDetailInput {
	s.ExpenseDate = &v
	return s
}

// SetGroupPeriod sets the GroupPeriod field's value.
func (s *ListSplitBillDetailInput) SetGroupPeriod(v int32) *ListSplitBillDetailInput {
	s.GroupPeriod = &v
	return s
}

// SetIgnoreZero sets the IgnoreZero field's value.
func (s *ListSplitBillDetailInput) SetIgnoreZero(v int32) *ListSplitBillDetailInput {
	s.IgnoreZero = &v
	return s
}

// SetInstanceNo sets the InstanceNo field's value.
func (s *ListSplitBillDetailInput) SetInstanceNo(v string) *ListSplitBillDetailInput {
	s.InstanceNo = &v
	return s
}

// SetLimit sets the Limit field's value.
func (s *ListSplitBillDetailInput) SetLimit(v int32) *ListSplitBillDetailInput {
	s.Limit = &v
	return s
}

// SetNeedRecordNum sets the NeedRecordNum field's value.
func (s *ListSplitBillDetailInput) SetNeedRecordNum(v int32) *ListSplitBillDetailInput {
	s.NeedRecordNum = &v
	return s
}

// SetOffset sets the Offset field's value.
func (s *ListSplitBillDetailInput) SetOffset(v int32) *ListSplitBillDetailInput {
	s.Offset = &v
	return s
}

// SetOwnerID sets the OwnerID field's value.
func (s *ListSplitBillDetailInput) SetOwnerID(v []*int64) *ListSplitBillDetailInput {
	s.OwnerID = v
	return s
}

// SetPayerID sets the PayerID field's value.
func (s *ListSplitBillDetailInput) SetPayerID(v []*int64) *ListSplitBillDetailInput {
	s.PayerID = v
	return s
}

// SetProduct sets the Product field's value.
func (s *ListSplitBillDetailInput) SetProduct(v []*string) *ListSplitBillDetailInput {
	s.Product = v
	return s
}

// SetSplitItemID sets the SplitItemID field's value.
func (s *ListSplitBillDetailInput) SetSplitItemID(v string) *ListSplitBillDetailInput {
	s.SplitItemID = &v
	return s
}

type ListSplitBillDetailOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	Limit *int32 `type:"int32"`

	List []*ListForListSplitBillDetailOutput `type:"list"`

	Offset *int32 `type:"int32"`

	Total *int32 `type:"int32"`
}

// String returns the string representation
func (s ListSplitBillDetailOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListSplitBillDetailOutput) GoString() string {
	return s.String()
}

// SetLimit sets the Limit field's value.
func (s *ListSplitBillDetailOutput) SetLimit(v int32) *ListSplitBillDetailOutput {
	s.Limit = &v
	return s
}

// SetList sets the List field's value.
func (s *ListSplitBillDetailOutput) SetList(v []*ListForListSplitBillDetailOutput) *ListSplitBillDetailOutput {
	s.List = v
	return s
}

// SetOffset sets the Offset field's value.
func (s *ListSplitBillDetailOutput) SetOffset(v int32) *ListSplitBillDetailOutput {
	s.Offset = &v
	return s
}

// SetTotal sets the Total field's value.
func (s *ListSplitBillDetailOutput) SetTotal(v int32) *ListSplitBillDetailOutput {
	s.Total = &v
	return s
}

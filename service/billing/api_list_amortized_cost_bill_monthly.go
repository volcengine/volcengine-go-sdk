// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package billing

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opListAmortizedCostBillMonthlyCommon = "ListAmortizedCostBillMonthly"

// ListAmortizedCostBillMonthlyCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ListAmortizedCostBillMonthlyCommon operation. The "output" return
// value will be populated with the ListAmortizedCostBillMonthlyCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListAmortizedCostBillMonthlyCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListAmortizedCostBillMonthlyCommon Send returns without error.
//
// See ListAmortizedCostBillMonthlyCommon for more information on using the ListAmortizedCostBillMonthlyCommon
// API call, and error handling.
//
//    // Example sending a request using the ListAmortizedCostBillMonthlyCommonRequest method.
//    req, resp := client.ListAmortizedCostBillMonthlyCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *BILLING) ListAmortizedCostBillMonthlyCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opListAmortizedCostBillMonthlyCommon,
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

// ListAmortizedCostBillMonthlyCommon API operation for BILLING.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for BILLING's
// API operation ListAmortizedCostBillMonthlyCommon for usage and error information.
func (c *BILLING) ListAmortizedCostBillMonthlyCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ListAmortizedCostBillMonthlyCommonRequest(input)
	return out, req.Send()
}

// ListAmortizedCostBillMonthlyCommonWithContext is the same as ListAmortizedCostBillMonthlyCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ListAmortizedCostBillMonthlyCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *BILLING) ListAmortizedCostBillMonthlyCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ListAmortizedCostBillMonthlyCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opListAmortizedCostBillMonthly = "ListAmortizedCostBillMonthly"

// ListAmortizedCostBillMonthlyRequest generates a "volcengine/request.Request" representing the
// client's request for the ListAmortizedCostBillMonthly operation. The "output" return
// value will be populated with the ListAmortizedCostBillMonthlyCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListAmortizedCostBillMonthlyCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListAmortizedCostBillMonthlyCommon Send returns without error.
//
// See ListAmortizedCostBillMonthly for more information on using the ListAmortizedCostBillMonthly
// API call, and error handling.
//
//    // Example sending a request using the ListAmortizedCostBillMonthlyRequest method.
//    req, resp := client.ListAmortizedCostBillMonthlyRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *BILLING) ListAmortizedCostBillMonthlyRequest(input *ListAmortizedCostBillMonthlyInput) (req *request.Request, output *ListAmortizedCostBillMonthlyOutput) {
	op := &request.Operation{
		Name:       opListAmortizedCostBillMonthly,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListAmortizedCostBillMonthlyInput{}
	}

	output = &ListAmortizedCostBillMonthlyOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// ListAmortizedCostBillMonthly API operation for BILLING.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for BILLING's
// API operation ListAmortizedCostBillMonthly for usage and error information.
func (c *BILLING) ListAmortizedCostBillMonthly(input *ListAmortizedCostBillMonthlyInput) (*ListAmortizedCostBillMonthlyOutput, error) {
	req, out := c.ListAmortizedCostBillMonthlyRequest(input)
	return out, req.Send()
}

// ListAmortizedCostBillMonthlyWithContext is the same as ListAmortizedCostBillMonthly with the addition of
// the ability to pass a context and additional request options.
//
// See ListAmortizedCostBillMonthly for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *BILLING) ListAmortizedCostBillMonthlyWithContext(ctx volcengine.Context, input *ListAmortizedCostBillMonthlyInput, opts ...request.Option) (*ListAmortizedCostBillMonthlyOutput, error) {
	req, out := c.ListAmortizedCostBillMonthlyRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ListAmortizedCostBillMonthlyInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	// AmortizedMonth is a required field
	AmortizedMonth *string `type:"string" json:",omitempty" required:"true"`

	AmortizedType []*string `type:"list" json:",omitempty"`

	BillCategory []*string `type:"list" json:",omitempty"`

	BillPeriod *string `type:"string" json:",omitempty"`

	BillingMode []*string `type:"list" json:",omitempty"`

	IgnoreZero *int32 `type:"int32" json:",omitempty"`

	InstanceNo *string `type:"string" json:",omitempty"`

	// Limit is a required field
	Limit *int32 `type:"int32" json:",omitempty" required:"true"`

	NeedRecordNum *int32 `type:"int32" json:",omitempty"`

	Offset *int32 `type:"int32" json:",omitempty"`

	OwnerID []*int64 `type:"list" json:",omitempty"`

	PayerID []*int64 `type:"list" json:",omitempty"`

	Product []*string `type:"list" json:",omitempty"`
}

// String returns the string representation
func (s ListAmortizedCostBillMonthlyInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListAmortizedCostBillMonthlyInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListAmortizedCostBillMonthlyInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ListAmortizedCostBillMonthlyInput"}
	if s.AmortizedMonth == nil {
		invalidParams.Add(request.NewErrParamRequired("AmortizedMonth"))
	}
	if s.Limit == nil {
		invalidParams.Add(request.NewErrParamRequired("Limit"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetAmortizedMonth sets the AmortizedMonth field's value.
func (s *ListAmortizedCostBillMonthlyInput) SetAmortizedMonth(v string) *ListAmortizedCostBillMonthlyInput {
	s.AmortizedMonth = &v
	return s
}

// SetAmortizedType sets the AmortizedType field's value.
func (s *ListAmortizedCostBillMonthlyInput) SetAmortizedType(v []*string) *ListAmortizedCostBillMonthlyInput {
	s.AmortizedType = v
	return s
}

// SetBillCategory sets the BillCategory field's value.
func (s *ListAmortizedCostBillMonthlyInput) SetBillCategory(v []*string) *ListAmortizedCostBillMonthlyInput {
	s.BillCategory = v
	return s
}

// SetBillPeriod sets the BillPeriod field's value.
func (s *ListAmortizedCostBillMonthlyInput) SetBillPeriod(v string) *ListAmortizedCostBillMonthlyInput {
	s.BillPeriod = &v
	return s
}

// SetBillingMode sets the BillingMode field's value.
func (s *ListAmortizedCostBillMonthlyInput) SetBillingMode(v []*string) *ListAmortizedCostBillMonthlyInput {
	s.BillingMode = v
	return s
}

// SetIgnoreZero sets the IgnoreZero field's value.
func (s *ListAmortizedCostBillMonthlyInput) SetIgnoreZero(v int32) *ListAmortizedCostBillMonthlyInput {
	s.IgnoreZero = &v
	return s
}

// SetInstanceNo sets the InstanceNo field's value.
func (s *ListAmortizedCostBillMonthlyInput) SetInstanceNo(v string) *ListAmortizedCostBillMonthlyInput {
	s.InstanceNo = &v
	return s
}

// SetLimit sets the Limit field's value.
func (s *ListAmortizedCostBillMonthlyInput) SetLimit(v int32) *ListAmortizedCostBillMonthlyInput {
	s.Limit = &v
	return s
}

// SetNeedRecordNum sets the NeedRecordNum field's value.
func (s *ListAmortizedCostBillMonthlyInput) SetNeedRecordNum(v int32) *ListAmortizedCostBillMonthlyInput {
	s.NeedRecordNum = &v
	return s
}

// SetOffset sets the Offset field's value.
func (s *ListAmortizedCostBillMonthlyInput) SetOffset(v int32) *ListAmortizedCostBillMonthlyInput {
	s.Offset = &v
	return s
}

// SetOwnerID sets the OwnerID field's value.
func (s *ListAmortizedCostBillMonthlyInput) SetOwnerID(v []*int64) *ListAmortizedCostBillMonthlyInput {
	s.OwnerID = v
	return s
}

// SetPayerID sets the PayerID field's value.
func (s *ListAmortizedCostBillMonthlyInput) SetPayerID(v []*int64) *ListAmortizedCostBillMonthlyInput {
	s.PayerID = v
	return s
}

// SetProduct sets the Product field's value.
func (s *ListAmortizedCostBillMonthlyInput) SetProduct(v []*string) *ListAmortizedCostBillMonthlyInput {
	s.Product = v
	return s
}

type ListAmortizedCostBillMonthlyOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	Limit *int32 `type:"int32" json:",omitempty"`

	List []*ListForListAmortizedCostBillMonthlyOutput `type:"list" json:",omitempty"`

	Offset *int32 `type:"int32" json:",omitempty"`

	Total *int32 `type:"int32" json:",omitempty"`
}

// String returns the string representation
func (s ListAmortizedCostBillMonthlyOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListAmortizedCostBillMonthlyOutput) GoString() string {
	return s.String()
}

// SetLimit sets the Limit field's value.
func (s *ListAmortizedCostBillMonthlyOutput) SetLimit(v int32) *ListAmortizedCostBillMonthlyOutput {
	s.Limit = &v
	return s
}

// SetList sets the List field's value.
func (s *ListAmortizedCostBillMonthlyOutput) SetList(v []*ListForListAmortizedCostBillMonthlyOutput) *ListAmortizedCostBillMonthlyOutput {
	s.List = v
	return s
}

// SetOffset sets the Offset field's value.
func (s *ListAmortizedCostBillMonthlyOutput) SetOffset(v int32) *ListAmortizedCostBillMonthlyOutput {
	s.Offset = &v
	return s
}

// SetTotal sets the Total field's value.
func (s *ListAmortizedCostBillMonthlyOutput) SetTotal(v int32) *ListAmortizedCostBillMonthlyOutput {
	s.Total = &v
	return s
}

type ListForListAmortizedCostBillMonthlyOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	AmortizedBeginTime *string `type:"string" json:",omitempty"`

	AmortizedDayNum *string `type:"string" json:",omitempty"`

	AmortizedEndTime *string `type:"string" json:",omitempty"`

	AmortizedMonth *string `type:"string" json:",omitempty"`

	AmortizedType *string `type:"string" json:",omitempty"`

	BeforeAmortizedCouponAmount *string `type:"string" json:",omitempty"`

	BeforeAmortizedDiscountBillAmount *string `type:"string" json:",omitempty"`

	BeforeAmortizedOriginalBillAmount *string `type:"string" json:",omitempty"`

	BeforeAmortizedPaidAmount *string `type:"string" json:",omitempty"`

	BeforeAmortizedPayableAmount *string `type:"string" json:",omitempty"`

	BeforeAmortizedPreferentialBillAmount *string `type:"string" json:",omitempty"`

	BeforeAmortizedRoundAmount *string `type:"string" json:",omitempty"`

	BillCategory *string `type:"string" json:",omitempty"`

	BillID *string `type:"string" json:",omitempty"`

	BillPeriod *string `type:"string" json:",omitempty"`

	BillingFunction *string `type:"string" json:",omitempty"`

	BillingMethodCode *string `type:"string" json:",omitempty"`

	BillingMode *string `type:"string" json:",omitempty"`

	BusiPeriod *string `type:"string" json:",omitempty"`

	BusinessMode *string `type:"string" json:",omitempty"`

	ConfigName *string `type:"string" json:",omitempty"`

	Count *string `type:"string" json:",omitempty"`

	CouponAmount *string `type:"string" json:",omitempty"`

	Currency *string `type:"string" json:",omitempty"`

	DailyAmortizedCouponAmount *string `type:"string" json:",omitempty"`

	DailyAmortizedDiscountBillAmount *string `type:"string" json:",omitempty"`

	DailyAmortizedOriginalBillAmount *string `type:"string" json:",omitempty"`

	DailyAmortizedPaidAmount *string `type:"string" json:",omitempty"`

	DailyAmortizedPayableAmount *string `type:"string" json:",omitempty"`

	DailyAmortizedPreferentialBillAmount *string `type:"string" json:",omitempty"`

	DailyAmortizedRoundAmount *string `type:"string" json:",omitempty"`

	DeductionUseDuration *string `type:"string" json:",omitempty"`

	DiscountBillAmount *string `type:"string" json:",omitempty"`

	DiscountBizBillingFunction *string `type:"string" json:",omitempty"`

	DiscountBizMeasureInterval *string `type:"string" json:",omitempty"`

	DiscountBizUnitPrice *string `type:"string" json:",omitempty"`

	DiscountBizUnitPriceInterval *string `type:"string" json:",omitempty"`

	EffectiveFactor *string `type:"string" json:",omitempty"`

	Element *string `type:"string" json:",omitempty"`

	ExpandField *string `type:"string" json:",omitempty"`

	ExpenseTime *string `type:"string" json:",omitempty"`

	Factor *string `type:"string" json:",omitempty"`

	InstanceName *string `type:"string" json:",omitempty"`

	InstanceNo *string `type:"string" json:",omitempty"`

	MarketPrice *string `type:"string" json:",omitempty"`

	MeasureInterval *string `type:"string" json:",omitempty"`

	NowAmortizedCouponAmount *string `type:"string" json:",omitempty"`

	NowAmortizedDiscountBillAmount *string `type:"string" json:",omitempty"`

	NowAmortizedOriginalBillAmount *string `type:"string" json:",omitempty"`

	NowAmortizedPaidAmount *string `type:"string" json:",omitempty"`

	NowAmortizedPayableAmount *string `type:"string" json:",omitempty"`

	NowAmortizedPreferentialBillAmount *string `type:"string" json:",omitempty"`

	NowAmortizedRoundAmount *string `type:"string" json:",omitempty"`

	OriginalBillAmount *string `type:"string" json:",omitempty"`

	OwnerCustomerName *string `type:"string" json:",omitempty"`

	OwnerID *string `type:"string" json:",omitempty"`

	OwnerUserName *string `type:"string" json:",omitempty"`

	PaidAmount *string `type:"string" json:",omitempty"`

	PayableAmount *string `type:"string" json:",omitempty"`

	PayerCustomerName *string `type:"string" json:",omitempty"`

	PayerID *string `type:"string" json:",omitempty"`

	PayerUserName *string `type:"string" json:",omitempty"`

	PreferentialBillAmount *string `type:"string" json:",omitempty"`

	Price *string `type:"string" json:",omitempty"`

	PriceInterval *string `type:"string" json:",omitempty"`

	PriceUnit *string `type:"string" json:",omitempty"`

	Product *string `type:"string" json:",omitempty"`

	ProductZh *string `type:"string" json:",omitempty"`

	Project *string `type:"string" json:",omitempty"`

	ProjectDisplayName *string `type:"string" json:",omitempty"`

	Region *string `type:"string" json:",omitempty"`

	RoundAmount *string `type:"string" json:",omitempty"`

	SellerCustomerName *string `type:"string" json:",omitempty"`

	SellerID *string `type:"string" json:",omitempty"`

	SellerUserName *string `type:"string" json:",omitempty"`

	SplitItemID *string `type:"string" json:",omitempty"`

	SplitItemName *string `type:"string" json:",omitempty"`

	SubjectName *string `type:"string" json:",omitempty"`

	Tag *string `type:"string" json:",omitempty"`

	UnamortizedCouponAmount *string `type:"string" json:",omitempty"`

	UnamortizedDiscountBillAmount *string `type:"string" json:",omitempty"`

	UnamortizedOriginalBillAmount *string `type:"string" json:",omitempty"`

	UnamortizedPaidAmount *string `type:"string" json:",omitempty"`

	UnamortizedPayableAmount *string `type:"string" json:",omitempty"`

	UnamortizedPreferentialBillAmount *string `type:"string" json:",omitempty"`

	UnamortizedRoundAmount *string `type:"string" json:",omitempty"`

	Unit *string `type:"string" json:",omitempty"`

	UseDuration *string `type:"string" json:",omitempty"`

	UseDurationUnit *string `type:"string" json:",omitempty"`

	Zone *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s ListForListAmortizedCostBillMonthlyOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListForListAmortizedCostBillMonthlyOutput) GoString() string {
	return s.String()
}

// SetAmortizedBeginTime sets the AmortizedBeginTime field's value.
func (s *ListForListAmortizedCostBillMonthlyOutput) SetAmortizedBeginTime(v string) *ListForListAmortizedCostBillMonthlyOutput {
	s.AmortizedBeginTime = &v
	return s
}

// SetAmortizedDayNum sets the AmortizedDayNum field's value.
func (s *ListForListAmortizedCostBillMonthlyOutput) SetAmortizedDayNum(v string) *ListForListAmortizedCostBillMonthlyOutput {
	s.AmortizedDayNum = &v
	return s
}

// SetAmortizedEndTime sets the AmortizedEndTime field's value.
func (s *ListForListAmortizedCostBillMonthlyOutput) SetAmortizedEndTime(v string) *ListForListAmortizedCostBillMonthlyOutput {
	s.AmortizedEndTime = &v
	return s
}

// SetAmortizedMonth sets the AmortizedMonth field's value.
func (s *ListForListAmortizedCostBillMonthlyOutput) SetAmortizedMonth(v string) *ListForListAmortizedCostBillMonthlyOutput {
	s.AmortizedMonth = &v
	return s
}

// SetAmortizedType sets the AmortizedType field's value.
func (s *ListForListAmortizedCostBillMonthlyOutput) SetAmortizedType(v string) *ListForListAmortizedCostBillMonthlyOutput {
	s.AmortizedType = &v
	return s
}

// SetBeforeAmortizedCouponAmount sets the BeforeAmortizedCouponAmount field's value.
func (s *ListForListAmortizedCostBillMonthlyOutput) SetBeforeAmortizedCouponAmount(v string) *ListForListAmortizedCostBillMonthlyOutput {
	s.BeforeAmortizedCouponAmount = &v
	return s
}

// SetBeforeAmortizedDiscountBillAmount sets the BeforeAmortizedDiscountBillAmount field's value.
func (s *ListForListAmortizedCostBillMonthlyOutput) SetBeforeAmortizedDiscountBillAmount(v string) *ListForListAmortizedCostBillMonthlyOutput {
	s.BeforeAmortizedDiscountBillAmount = &v
	return s
}

// SetBeforeAmortizedOriginalBillAmount sets the BeforeAmortizedOriginalBillAmount field's value.
func (s *ListForListAmortizedCostBillMonthlyOutput) SetBeforeAmortizedOriginalBillAmount(v string) *ListForListAmortizedCostBillMonthlyOutput {
	s.BeforeAmortizedOriginalBillAmount = &v
	return s
}

// SetBeforeAmortizedPaidAmount sets the BeforeAmortizedPaidAmount field's value.
func (s *ListForListAmortizedCostBillMonthlyOutput) SetBeforeAmortizedPaidAmount(v string) *ListForListAmortizedCostBillMonthlyOutput {
	s.BeforeAmortizedPaidAmount = &v
	return s
}

// SetBeforeAmortizedPayableAmount sets the BeforeAmortizedPayableAmount field's value.
func (s *ListForListAmortizedCostBillMonthlyOutput) SetBeforeAmortizedPayableAmount(v string) *ListForListAmortizedCostBillMonthlyOutput {
	s.BeforeAmortizedPayableAmount = &v
	return s
}

// SetBeforeAmortizedPreferentialBillAmount sets the BeforeAmortizedPreferentialBillAmount field's value.
func (s *ListForListAmortizedCostBillMonthlyOutput) SetBeforeAmortizedPreferentialBillAmount(v string) *ListForListAmortizedCostBillMonthlyOutput {
	s.BeforeAmortizedPreferentialBillAmount = &v
	return s
}

// SetBeforeAmortizedRoundAmount sets the BeforeAmortizedRoundAmount field's value.
func (s *ListForListAmortizedCostBillMonthlyOutput) SetBeforeAmortizedRoundAmount(v string) *ListForListAmortizedCostBillMonthlyOutput {
	s.BeforeAmortizedRoundAmount = &v
	return s
}

// SetBillCategory sets the BillCategory field's value.
func (s *ListForListAmortizedCostBillMonthlyOutput) SetBillCategory(v string) *ListForListAmortizedCostBillMonthlyOutput {
	s.BillCategory = &v
	return s
}

// SetBillID sets the BillID field's value.
func (s *ListForListAmortizedCostBillMonthlyOutput) SetBillID(v string) *ListForListAmortizedCostBillMonthlyOutput {
	s.BillID = &v
	return s
}

// SetBillPeriod sets the BillPeriod field's value.
func (s *ListForListAmortizedCostBillMonthlyOutput) SetBillPeriod(v string) *ListForListAmortizedCostBillMonthlyOutput {
	s.BillPeriod = &v
	return s
}

// SetBillingFunction sets the BillingFunction field's value.
func (s *ListForListAmortizedCostBillMonthlyOutput) SetBillingFunction(v string) *ListForListAmortizedCostBillMonthlyOutput {
	s.BillingFunction = &v
	return s
}

// SetBillingMethodCode sets the BillingMethodCode field's value.
func (s *ListForListAmortizedCostBillMonthlyOutput) SetBillingMethodCode(v string) *ListForListAmortizedCostBillMonthlyOutput {
	s.BillingMethodCode = &v
	return s
}

// SetBillingMode sets the BillingMode field's value.
func (s *ListForListAmortizedCostBillMonthlyOutput) SetBillingMode(v string) *ListForListAmortizedCostBillMonthlyOutput {
	s.BillingMode = &v
	return s
}

// SetBusiPeriod sets the BusiPeriod field's value.
func (s *ListForListAmortizedCostBillMonthlyOutput) SetBusiPeriod(v string) *ListForListAmortizedCostBillMonthlyOutput {
	s.BusiPeriod = &v
	return s
}

// SetBusinessMode sets the BusinessMode field's value.
func (s *ListForListAmortizedCostBillMonthlyOutput) SetBusinessMode(v string) *ListForListAmortizedCostBillMonthlyOutput {
	s.BusinessMode = &v
	return s
}

// SetConfigName sets the ConfigName field's value.
func (s *ListForListAmortizedCostBillMonthlyOutput) SetConfigName(v string) *ListForListAmortizedCostBillMonthlyOutput {
	s.ConfigName = &v
	return s
}

// SetCount sets the Count field's value.
func (s *ListForListAmortizedCostBillMonthlyOutput) SetCount(v string) *ListForListAmortizedCostBillMonthlyOutput {
	s.Count = &v
	return s
}

// SetCouponAmount sets the CouponAmount field's value.
func (s *ListForListAmortizedCostBillMonthlyOutput) SetCouponAmount(v string) *ListForListAmortizedCostBillMonthlyOutput {
	s.CouponAmount = &v
	return s
}

// SetCurrency sets the Currency field's value.
func (s *ListForListAmortizedCostBillMonthlyOutput) SetCurrency(v string) *ListForListAmortizedCostBillMonthlyOutput {
	s.Currency = &v
	return s
}

// SetDailyAmortizedCouponAmount sets the DailyAmortizedCouponAmount field's value.
func (s *ListForListAmortizedCostBillMonthlyOutput) SetDailyAmortizedCouponAmount(v string) *ListForListAmortizedCostBillMonthlyOutput {
	s.DailyAmortizedCouponAmount = &v
	return s
}

// SetDailyAmortizedDiscountBillAmount sets the DailyAmortizedDiscountBillAmount field's value.
func (s *ListForListAmortizedCostBillMonthlyOutput) SetDailyAmortizedDiscountBillAmount(v string) *ListForListAmortizedCostBillMonthlyOutput {
	s.DailyAmortizedDiscountBillAmount = &v
	return s
}

// SetDailyAmortizedOriginalBillAmount sets the DailyAmortizedOriginalBillAmount field's value.
func (s *ListForListAmortizedCostBillMonthlyOutput) SetDailyAmortizedOriginalBillAmount(v string) *ListForListAmortizedCostBillMonthlyOutput {
	s.DailyAmortizedOriginalBillAmount = &v
	return s
}

// SetDailyAmortizedPaidAmount sets the DailyAmortizedPaidAmount field's value.
func (s *ListForListAmortizedCostBillMonthlyOutput) SetDailyAmortizedPaidAmount(v string) *ListForListAmortizedCostBillMonthlyOutput {
	s.DailyAmortizedPaidAmount = &v
	return s
}

// SetDailyAmortizedPayableAmount sets the DailyAmortizedPayableAmount field's value.
func (s *ListForListAmortizedCostBillMonthlyOutput) SetDailyAmortizedPayableAmount(v string) *ListForListAmortizedCostBillMonthlyOutput {
	s.DailyAmortizedPayableAmount = &v
	return s
}

// SetDailyAmortizedPreferentialBillAmount sets the DailyAmortizedPreferentialBillAmount field's value.
func (s *ListForListAmortizedCostBillMonthlyOutput) SetDailyAmortizedPreferentialBillAmount(v string) *ListForListAmortizedCostBillMonthlyOutput {
	s.DailyAmortizedPreferentialBillAmount = &v
	return s
}

// SetDailyAmortizedRoundAmount sets the DailyAmortizedRoundAmount field's value.
func (s *ListForListAmortizedCostBillMonthlyOutput) SetDailyAmortizedRoundAmount(v string) *ListForListAmortizedCostBillMonthlyOutput {
	s.DailyAmortizedRoundAmount = &v
	return s
}

// SetDeductionUseDuration sets the DeductionUseDuration field's value.
func (s *ListForListAmortizedCostBillMonthlyOutput) SetDeductionUseDuration(v string) *ListForListAmortizedCostBillMonthlyOutput {
	s.DeductionUseDuration = &v
	return s
}

// SetDiscountBillAmount sets the DiscountBillAmount field's value.
func (s *ListForListAmortizedCostBillMonthlyOutput) SetDiscountBillAmount(v string) *ListForListAmortizedCostBillMonthlyOutput {
	s.DiscountBillAmount = &v
	return s
}

// SetDiscountBizBillingFunction sets the DiscountBizBillingFunction field's value.
func (s *ListForListAmortizedCostBillMonthlyOutput) SetDiscountBizBillingFunction(v string) *ListForListAmortizedCostBillMonthlyOutput {
	s.DiscountBizBillingFunction = &v
	return s
}

// SetDiscountBizMeasureInterval sets the DiscountBizMeasureInterval field's value.
func (s *ListForListAmortizedCostBillMonthlyOutput) SetDiscountBizMeasureInterval(v string) *ListForListAmortizedCostBillMonthlyOutput {
	s.DiscountBizMeasureInterval = &v
	return s
}

// SetDiscountBizUnitPrice sets the DiscountBizUnitPrice field's value.
func (s *ListForListAmortizedCostBillMonthlyOutput) SetDiscountBizUnitPrice(v string) *ListForListAmortizedCostBillMonthlyOutput {
	s.DiscountBizUnitPrice = &v
	return s
}

// SetDiscountBizUnitPriceInterval sets the DiscountBizUnitPriceInterval field's value.
func (s *ListForListAmortizedCostBillMonthlyOutput) SetDiscountBizUnitPriceInterval(v string) *ListForListAmortizedCostBillMonthlyOutput {
	s.DiscountBizUnitPriceInterval = &v
	return s
}

// SetEffectiveFactor sets the EffectiveFactor field's value.
func (s *ListForListAmortizedCostBillMonthlyOutput) SetEffectiveFactor(v string) *ListForListAmortizedCostBillMonthlyOutput {
	s.EffectiveFactor = &v
	return s
}

// SetElement sets the Element field's value.
func (s *ListForListAmortizedCostBillMonthlyOutput) SetElement(v string) *ListForListAmortizedCostBillMonthlyOutput {
	s.Element = &v
	return s
}

// SetExpandField sets the ExpandField field's value.
func (s *ListForListAmortizedCostBillMonthlyOutput) SetExpandField(v string) *ListForListAmortizedCostBillMonthlyOutput {
	s.ExpandField = &v
	return s
}

// SetExpenseTime sets the ExpenseTime field's value.
func (s *ListForListAmortizedCostBillMonthlyOutput) SetExpenseTime(v string) *ListForListAmortizedCostBillMonthlyOutput {
	s.ExpenseTime = &v
	return s
}

// SetFactor sets the Factor field's value.
func (s *ListForListAmortizedCostBillMonthlyOutput) SetFactor(v string) *ListForListAmortizedCostBillMonthlyOutput {
	s.Factor = &v
	return s
}

// SetInstanceName sets the InstanceName field's value.
func (s *ListForListAmortizedCostBillMonthlyOutput) SetInstanceName(v string) *ListForListAmortizedCostBillMonthlyOutput {
	s.InstanceName = &v
	return s
}

// SetInstanceNo sets the InstanceNo field's value.
func (s *ListForListAmortizedCostBillMonthlyOutput) SetInstanceNo(v string) *ListForListAmortizedCostBillMonthlyOutput {
	s.InstanceNo = &v
	return s
}

// SetMarketPrice sets the MarketPrice field's value.
func (s *ListForListAmortizedCostBillMonthlyOutput) SetMarketPrice(v string) *ListForListAmortizedCostBillMonthlyOutput {
	s.MarketPrice = &v
	return s
}

// SetMeasureInterval sets the MeasureInterval field's value.
func (s *ListForListAmortizedCostBillMonthlyOutput) SetMeasureInterval(v string) *ListForListAmortizedCostBillMonthlyOutput {
	s.MeasureInterval = &v
	return s
}

// SetNowAmortizedCouponAmount sets the NowAmortizedCouponAmount field's value.
func (s *ListForListAmortizedCostBillMonthlyOutput) SetNowAmortizedCouponAmount(v string) *ListForListAmortizedCostBillMonthlyOutput {
	s.NowAmortizedCouponAmount = &v
	return s
}

// SetNowAmortizedDiscountBillAmount sets the NowAmortizedDiscountBillAmount field's value.
func (s *ListForListAmortizedCostBillMonthlyOutput) SetNowAmortizedDiscountBillAmount(v string) *ListForListAmortizedCostBillMonthlyOutput {
	s.NowAmortizedDiscountBillAmount = &v
	return s
}

// SetNowAmortizedOriginalBillAmount sets the NowAmortizedOriginalBillAmount field's value.
func (s *ListForListAmortizedCostBillMonthlyOutput) SetNowAmortizedOriginalBillAmount(v string) *ListForListAmortizedCostBillMonthlyOutput {
	s.NowAmortizedOriginalBillAmount = &v
	return s
}

// SetNowAmortizedPaidAmount sets the NowAmortizedPaidAmount field's value.
func (s *ListForListAmortizedCostBillMonthlyOutput) SetNowAmortizedPaidAmount(v string) *ListForListAmortizedCostBillMonthlyOutput {
	s.NowAmortizedPaidAmount = &v
	return s
}

// SetNowAmortizedPayableAmount sets the NowAmortizedPayableAmount field's value.
func (s *ListForListAmortizedCostBillMonthlyOutput) SetNowAmortizedPayableAmount(v string) *ListForListAmortizedCostBillMonthlyOutput {
	s.NowAmortizedPayableAmount = &v
	return s
}

// SetNowAmortizedPreferentialBillAmount sets the NowAmortizedPreferentialBillAmount field's value.
func (s *ListForListAmortizedCostBillMonthlyOutput) SetNowAmortizedPreferentialBillAmount(v string) *ListForListAmortizedCostBillMonthlyOutput {
	s.NowAmortizedPreferentialBillAmount = &v
	return s
}

// SetNowAmortizedRoundAmount sets the NowAmortizedRoundAmount field's value.
func (s *ListForListAmortizedCostBillMonthlyOutput) SetNowAmortizedRoundAmount(v string) *ListForListAmortizedCostBillMonthlyOutput {
	s.NowAmortizedRoundAmount = &v
	return s
}

// SetOriginalBillAmount sets the OriginalBillAmount field's value.
func (s *ListForListAmortizedCostBillMonthlyOutput) SetOriginalBillAmount(v string) *ListForListAmortizedCostBillMonthlyOutput {
	s.OriginalBillAmount = &v
	return s
}

// SetOwnerCustomerName sets the OwnerCustomerName field's value.
func (s *ListForListAmortizedCostBillMonthlyOutput) SetOwnerCustomerName(v string) *ListForListAmortizedCostBillMonthlyOutput {
	s.OwnerCustomerName = &v
	return s
}

// SetOwnerID sets the OwnerID field's value.
func (s *ListForListAmortizedCostBillMonthlyOutput) SetOwnerID(v string) *ListForListAmortizedCostBillMonthlyOutput {
	s.OwnerID = &v
	return s
}

// SetOwnerUserName sets the OwnerUserName field's value.
func (s *ListForListAmortizedCostBillMonthlyOutput) SetOwnerUserName(v string) *ListForListAmortizedCostBillMonthlyOutput {
	s.OwnerUserName = &v
	return s
}

// SetPaidAmount sets the PaidAmount field's value.
func (s *ListForListAmortizedCostBillMonthlyOutput) SetPaidAmount(v string) *ListForListAmortizedCostBillMonthlyOutput {
	s.PaidAmount = &v
	return s
}

// SetPayableAmount sets the PayableAmount field's value.
func (s *ListForListAmortizedCostBillMonthlyOutput) SetPayableAmount(v string) *ListForListAmortizedCostBillMonthlyOutput {
	s.PayableAmount = &v
	return s
}

// SetPayerCustomerName sets the PayerCustomerName field's value.
func (s *ListForListAmortizedCostBillMonthlyOutput) SetPayerCustomerName(v string) *ListForListAmortizedCostBillMonthlyOutput {
	s.PayerCustomerName = &v
	return s
}

// SetPayerID sets the PayerID field's value.
func (s *ListForListAmortizedCostBillMonthlyOutput) SetPayerID(v string) *ListForListAmortizedCostBillMonthlyOutput {
	s.PayerID = &v
	return s
}

// SetPayerUserName sets the PayerUserName field's value.
func (s *ListForListAmortizedCostBillMonthlyOutput) SetPayerUserName(v string) *ListForListAmortizedCostBillMonthlyOutput {
	s.PayerUserName = &v
	return s
}

// SetPreferentialBillAmount sets the PreferentialBillAmount field's value.
func (s *ListForListAmortizedCostBillMonthlyOutput) SetPreferentialBillAmount(v string) *ListForListAmortizedCostBillMonthlyOutput {
	s.PreferentialBillAmount = &v
	return s
}

// SetPrice sets the Price field's value.
func (s *ListForListAmortizedCostBillMonthlyOutput) SetPrice(v string) *ListForListAmortizedCostBillMonthlyOutput {
	s.Price = &v
	return s
}

// SetPriceInterval sets the PriceInterval field's value.
func (s *ListForListAmortizedCostBillMonthlyOutput) SetPriceInterval(v string) *ListForListAmortizedCostBillMonthlyOutput {
	s.PriceInterval = &v
	return s
}

// SetPriceUnit sets the PriceUnit field's value.
func (s *ListForListAmortizedCostBillMonthlyOutput) SetPriceUnit(v string) *ListForListAmortizedCostBillMonthlyOutput {
	s.PriceUnit = &v
	return s
}

// SetProduct sets the Product field's value.
func (s *ListForListAmortizedCostBillMonthlyOutput) SetProduct(v string) *ListForListAmortizedCostBillMonthlyOutput {
	s.Product = &v
	return s
}

// SetProductZh sets the ProductZh field's value.
func (s *ListForListAmortizedCostBillMonthlyOutput) SetProductZh(v string) *ListForListAmortizedCostBillMonthlyOutput {
	s.ProductZh = &v
	return s
}

// SetProject sets the Project field's value.
func (s *ListForListAmortizedCostBillMonthlyOutput) SetProject(v string) *ListForListAmortizedCostBillMonthlyOutput {
	s.Project = &v
	return s
}

// SetProjectDisplayName sets the ProjectDisplayName field's value.
func (s *ListForListAmortizedCostBillMonthlyOutput) SetProjectDisplayName(v string) *ListForListAmortizedCostBillMonthlyOutput {
	s.ProjectDisplayName = &v
	return s
}

// SetRegion sets the Region field's value.
func (s *ListForListAmortizedCostBillMonthlyOutput) SetRegion(v string) *ListForListAmortizedCostBillMonthlyOutput {
	s.Region = &v
	return s
}

// SetRoundAmount sets the RoundAmount field's value.
func (s *ListForListAmortizedCostBillMonthlyOutput) SetRoundAmount(v string) *ListForListAmortizedCostBillMonthlyOutput {
	s.RoundAmount = &v
	return s
}

// SetSellerCustomerName sets the SellerCustomerName field's value.
func (s *ListForListAmortizedCostBillMonthlyOutput) SetSellerCustomerName(v string) *ListForListAmortizedCostBillMonthlyOutput {
	s.SellerCustomerName = &v
	return s
}

// SetSellerID sets the SellerID field's value.
func (s *ListForListAmortizedCostBillMonthlyOutput) SetSellerID(v string) *ListForListAmortizedCostBillMonthlyOutput {
	s.SellerID = &v
	return s
}

// SetSellerUserName sets the SellerUserName field's value.
func (s *ListForListAmortizedCostBillMonthlyOutput) SetSellerUserName(v string) *ListForListAmortizedCostBillMonthlyOutput {
	s.SellerUserName = &v
	return s
}

// SetSplitItemID sets the SplitItemID field's value.
func (s *ListForListAmortizedCostBillMonthlyOutput) SetSplitItemID(v string) *ListForListAmortizedCostBillMonthlyOutput {
	s.SplitItemID = &v
	return s
}

// SetSplitItemName sets the SplitItemName field's value.
func (s *ListForListAmortizedCostBillMonthlyOutput) SetSplitItemName(v string) *ListForListAmortizedCostBillMonthlyOutput {
	s.SplitItemName = &v
	return s
}

// SetSubjectName sets the SubjectName field's value.
func (s *ListForListAmortizedCostBillMonthlyOutput) SetSubjectName(v string) *ListForListAmortizedCostBillMonthlyOutput {
	s.SubjectName = &v
	return s
}

// SetTag sets the Tag field's value.
func (s *ListForListAmortizedCostBillMonthlyOutput) SetTag(v string) *ListForListAmortizedCostBillMonthlyOutput {
	s.Tag = &v
	return s
}

// SetUnamortizedCouponAmount sets the UnamortizedCouponAmount field's value.
func (s *ListForListAmortizedCostBillMonthlyOutput) SetUnamortizedCouponAmount(v string) *ListForListAmortizedCostBillMonthlyOutput {
	s.UnamortizedCouponAmount = &v
	return s
}

// SetUnamortizedDiscountBillAmount sets the UnamortizedDiscountBillAmount field's value.
func (s *ListForListAmortizedCostBillMonthlyOutput) SetUnamortizedDiscountBillAmount(v string) *ListForListAmortizedCostBillMonthlyOutput {
	s.UnamortizedDiscountBillAmount = &v
	return s
}

// SetUnamortizedOriginalBillAmount sets the UnamortizedOriginalBillAmount field's value.
func (s *ListForListAmortizedCostBillMonthlyOutput) SetUnamortizedOriginalBillAmount(v string) *ListForListAmortizedCostBillMonthlyOutput {
	s.UnamortizedOriginalBillAmount = &v
	return s
}

// SetUnamortizedPaidAmount sets the UnamortizedPaidAmount field's value.
func (s *ListForListAmortizedCostBillMonthlyOutput) SetUnamortizedPaidAmount(v string) *ListForListAmortizedCostBillMonthlyOutput {
	s.UnamortizedPaidAmount = &v
	return s
}

// SetUnamortizedPayableAmount sets the UnamortizedPayableAmount field's value.
func (s *ListForListAmortizedCostBillMonthlyOutput) SetUnamortizedPayableAmount(v string) *ListForListAmortizedCostBillMonthlyOutput {
	s.UnamortizedPayableAmount = &v
	return s
}

// SetUnamortizedPreferentialBillAmount sets the UnamortizedPreferentialBillAmount field's value.
func (s *ListForListAmortizedCostBillMonthlyOutput) SetUnamortizedPreferentialBillAmount(v string) *ListForListAmortizedCostBillMonthlyOutput {
	s.UnamortizedPreferentialBillAmount = &v
	return s
}

// SetUnamortizedRoundAmount sets the UnamortizedRoundAmount field's value.
func (s *ListForListAmortizedCostBillMonthlyOutput) SetUnamortizedRoundAmount(v string) *ListForListAmortizedCostBillMonthlyOutput {
	s.UnamortizedRoundAmount = &v
	return s
}

// SetUnit sets the Unit field's value.
func (s *ListForListAmortizedCostBillMonthlyOutput) SetUnit(v string) *ListForListAmortizedCostBillMonthlyOutput {
	s.Unit = &v
	return s
}

// SetUseDuration sets the UseDuration field's value.
func (s *ListForListAmortizedCostBillMonthlyOutput) SetUseDuration(v string) *ListForListAmortizedCostBillMonthlyOutput {
	s.UseDuration = &v
	return s
}

// SetUseDurationUnit sets the UseDurationUnit field's value.
func (s *ListForListAmortizedCostBillMonthlyOutput) SetUseDurationUnit(v string) *ListForListAmortizedCostBillMonthlyOutput {
	s.UseDurationUnit = &v
	return s
}

// SetZone sets the Zone field's value.
func (s *ListForListAmortizedCostBillMonthlyOutput) SetZone(v string) *ListForListAmortizedCostBillMonthlyOutput {
	s.Zone = &v
	return s
}

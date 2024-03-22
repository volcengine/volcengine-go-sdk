// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package billing

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opListAmortizedCostBillDailyCommon = "ListAmortizedCostBillDaily"

// ListAmortizedCostBillDailyCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ListAmortizedCostBillDailyCommon operation. The "output" return
// value will be populated with the ListAmortizedCostBillDailyCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListAmortizedCostBillDailyCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListAmortizedCostBillDailyCommon Send returns without error.
//
// See ListAmortizedCostBillDailyCommon for more information on using the ListAmortizedCostBillDailyCommon
// API call, and error handling.
//
//    // Example sending a request using the ListAmortizedCostBillDailyCommonRequest method.
//    req, resp := client.ListAmortizedCostBillDailyCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *BILLING) ListAmortizedCostBillDailyCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opListAmortizedCostBillDailyCommon,
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

// ListAmortizedCostBillDailyCommon API operation for BILLING.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for BILLING's
// API operation ListAmortizedCostBillDailyCommon for usage and error information.
func (c *BILLING) ListAmortizedCostBillDailyCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ListAmortizedCostBillDailyCommonRequest(input)
	return out, req.Send()
}

// ListAmortizedCostBillDailyCommonWithContext is the same as ListAmortizedCostBillDailyCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ListAmortizedCostBillDailyCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *BILLING) ListAmortizedCostBillDailyCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ListAmortizedCostBillDailyCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opListAmortizedCostBillDaily = "ListAmortizedCostBillDaily"

// ListAmortizedCostBillDailyRequest generates a "volcengine/request.Request" representing the
// client's request for the ListAmortizedCostBillDaily operation. The "output" return
// value will be populated with the ListAmortizedCostBillDailyCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListAmortizedCostBillDailyCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListAmortizedCostBillDailyCommon Send returns without error.
//
// See ListAmortizedCostBillDaily for more information on using the ListAmortizedCostBillDaily
// API call, and error handling.
//
//    // Example sending a request using the ListAmortizedCostBillDailyRequest method.
//    req, resp := client.ListAmortizedCostBillDailyRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *BILLING) ListAmortizedCostBillDailyRequest(input *ListAmortizedCostBillDailyInput) (req *request.Request, output *ListAmortizedCostBillDailyOutput) {
	op := &request.Operation{
		Name:       opListAmortizedCostBillDaily,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListAmortizedCostBillDailyInput{}
	}

	output = &ListAmortizedCostBillDailyOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// ListAmortizedCostBillDaily API operation for BILLING.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for BILLING's
// API operation ListAmortizedCostBillDaily for usage and error information.
func (c *BILLING) ListAmortizedCostBillDaily(input *ListAmortizedCostBillDailyInput) (*ListAmortizedCostBillDailyOutput, error) {
	req, out := c.ListAmortizedCostBillDailyRequest(input)
	return out, req.Send()
}

// ListAmortizedCostBillDailyWithContext is the same as ListAmortizedCostBillDaily with the addition of
// the ability to pass a context and additional request options.
//
// See ListAmortizedCostBillDaily for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *BILLING) ListAmortizedCostBillDailyWithContext(ctx volcengine.Context, input *ListAmortizedCostBillDailyInput, opts ...request.Option) (*ListAmortizedCostBillDailyOutput, error) {
	req, out := c.ListAmortizedCostBillDailyRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ListAmortizedCostBillDailyInput struct {
	_ struct{} `type:"structure"`

	AmortizedDay *string `type:"string"`

	// AmortizedMonth is a required field
	AmortizedMonth *string `type:"string" required:"true"`

	AmortizedType []*string `type:"list"`

	BillCategory []*string `type:"list"`

	BillPeriod *string `type:"string"`

	BillingMode []*string `type:"list"`

	IgnoreZero *int32 `type:"int32"`

	InstanceNo *string `type:"string"`

	// Limit is a required field
	Limit *int32 `type:"int32" required:"true"`

	NeedRecordNum *int32 `type:"int32"`

	Offset *int32 `type:"int32"`

	Product []*string `type:"list"`
}

// String returns the string representation
func (s ListAmortizedCostBillDailyInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListAmortizedCostBillDailyInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListAmortizedCostBillDailyInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ListAmortizedCostBillDailyInput"}
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

// SetAmortizedDay sets the AmortizedDay field's value.
func (s *ListAmortizedCostBillDailyInput) SetAmortizedDay(v string) *ListAmortizedCostBillDailyInput {
	s.AmortizedDay = &v
	return s
}

// SetAmortizedMonth sets the AmortizedMonth field's value.
func (s *ListAmortizedCostBillDailyInput) SetAmortizedMonth(v string) *ListAmortizedCostBillDailyInput {
	s.AmortizedMonth = &v
	return s
}

// SetAmortizedType sets the AmortizedType field's value.
func (s *ListAmortizedCostBillDailyInput) SetAmortizedType(v []*string) *ListAmortizedCostBillDailyInput {
	s.AmortizedType = v
	return s
}

// SetBillCategory sets the BillCategory field's value.
func (s *ListAmortizedCostBillDailyInput) SetBillCategory(v []*string) *ListAmortizedCostBillDailyInput {
	s.BillCategory = v
	return s
}

// SetBillPeriod sets the BillPeriod field's value.
func (s *ListAmortizedCostBillDailyInput) SetBillPeriod(v string) *ListAmortizedCostBillDailyInput {
	s.BillPeriod = &v
	return s
}

// SetBillingMode sets the BillingMode field's value.
func (s *ListAmortizedCostBillDailyInput) SetBillingMode(v []*string) *ListAmortizedCostBillDailyInput {
	s.BillingMode = v
	return s
}

// SetIgnoreZero sets the IgnoreZero field's value.
func (s *ListAmortizedCostBillDailyInput) SetIgnoreZero(v int32) *ListAmortizedCostBillDailyInput {
	s.IgnoreZero = &v
	return s
}

// SetInstanceNo sets the InstanceNo field's value.
func (s *ListAmortizedCostBillDailyInput) SetInstanceNo(v string) *ListAmortizedCostBillDailyInput {
	s.InstanceNo = &v
	return s
}

// SetLimit sets the Limit field's value.
func (s *ListAmortizedCostBillDailyInput) SetLimit(v int32) *ListAmortizedCostBillDailyInput {
	s.Limit = &v
	return s
}

// SetNeedRecordNum sets the NeedRecordNum field's value.
func (s *ListAmortizedCostBillDailyInput) SetNeedRecordNum(v int32) *ListAmortizedCostBillDailyInput {
	s.NeedRecordNum = &v
	return s
}

// SetOffset sets the Offset field's value.
func (s *ListAmortizedCostBillDailyInput) SetOffset(v int32) *ListAmortizedCostBillDailyInput {
	s.Offset = &v
	return s
}

// SetProduct sets the Product field's value.
func (s *ListAmortizedCostBillDailyInput) SetProduct(v []*string) *ListAmortizedCostBillDailyInput {
	s.Product = v
	return s
}

type ListAmortizedCostBillDailyOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	Limit *int32 `type:"int32"`

	List []*ListForListAmortizedCostBillDailyOutput `type:"list"`

	Offset *int32 `type:"int32"`

	Total *int32 `type:"int32"`
}

// String returns the string representation
func (s ListAmortizedCostBillDailyOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListAmortizedCostBillDailyOutput) GoString() string {
	return s.String()
}

// SetLimit sets the Limit field's value.
func (s *ListAmortizedCostBillDailyOutput) SetLimit(v int32) *ListAmortizedCostBillDailyOutput {
	s.Limit = &v
	return s
}

// SetList sets the List field's value.
func (s *ListAmortizedCostBillDailyOutput) SetList(v []*ListForListAmortizedCostBillDailyOutput) *ListAmortizedCostBillDailyOutput {
	s.List = v
	return s
}

// SetOffset sets the Offset field's value.
func (s *ListAmortizedCostBillDailyOutput) SetOffset(v int32) *ListAmortizedCostBillDailyOutput {
	s.Offset = &v
	return s
}

// SetTotal sets the Total field's value.
func (s *ListAmortizedCostBillDailyOutput) SetTotal(v int32) *ListAmortizedCostBillDailyOutput {
	s.Total = &v
	return s
}

type ListForListAmortizedCostBillDailyOutput struct {
	_ struct{} `type:"structure"`

	AmortizedBeginTime *string `type:"string"`

	AmortizedDay *string `type:"string"`

	AmortizedEndTime *string `type:"string"`

	AmortizedMonth *string `type:"string"`

	AmortizedType *string `type:"string"`

	BillCategory *string `type:"string"`

	BillID *string `type:"string"`

	BillPeriod *string `type:"string"`

	BillingMethodCode *string `type:"string"`

	BillingMode *string `type:"string"`

	BusiPeriod *string `type:"string"`

	BusinessMode *string `type:"string"`

	ConfigName *string `type:"string"`

	ConfigurationCode *string `type:"string"`

	Count *string `type:"string"`

	CouponAmount *string `type:"string"`

	Currency *string `type:"string"`

	DailyAmortizedCouponAmount *string `type:"string"`

	DailyAmortizedDiscountBillAmount *string `type:"string"`

	DailyAmortizedOriginalBillAmount *string `type:"string"`

	DailyAmortizedPaidAmount *string `type:"string"`

	DailyAmortizedPayableAmount *string `type:"string"`

	DailyAmortizedPreferentialBillAmount *string `type:"string"`

	DailyAmortizedRoundAmount *string `type:"string"`

	DiscountBillAmount *string `type:"string"`

	Element *string `type:"string"`

	ElementCode *string `type:"string"`

	ExpandField *string `type:"string"`

	ExpenseTime *string `type:"string"`

	Factor *string `type:"string"`

	FactorCode *string `type:"string"`

	InstanceName *string `type:"string"`

	InstanceNo *string `type:"string"`

	NowAmortizedCouponAmount *string `type:"string"`

	NowAmortizedDiscountBillAmount *string `type:"string"`

	NowAmortizedOriginalBillAmount *string `type:"string"`

	NowAmortizedPaidAmount *string `type:"string"`

	NowAmortizedPayableAmount *string `type:"string"`

	NowAmortizedPreferentialBillAmount *string `type:"string"`

	NowAmortizedRoundAmount *string `type:"string"`

	OriginalBillAmount *string `type:"string"`

	OwnerCustomerName *string `type:"string"`

	OwnerID *string `type:"string"`

	OwnerUserName *string `type:"string"`

	PaidAmount *string `type:"string"`

	PayableAmount *string `type:"string"`

	PayerCustomerName *string `type:"string"`

	PayerID *string `type:"string"`

	PayerUserName *string `type:"string"`

	PreferentialBillAmount *string `type:"string"`

	Price *string `type:"string"`

	PriceUnit *string `type:"string"`

	Product *string `type:"string"`

	ProductZh *string `type:"string"`

	Project *string `type:"string"`

	Region *string `type:"string"`

	RegionCode *string `type:"string"`

	RoundAmount *string `type:"string"`

	SellerCustomerName *string `type:"string"`

	SellerID *string `type:"string"`

	SellerUserName *string `type:"string"`

	SubjectName *string `type:"string"`

	Tag *string `type:"string"`

	Unit *string `type:"string"`

	UseDuration *string `type:"string"`

	UseDurationUnit *string `type:"string"`

	Zone *string `type:"string"`

	ZoneCode *string `type:"string"`
}

// String returns the string representation
func (s ListForListAmortizedCostBillDailyOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListForListAmortizedCostBillDailyOutput) GoString() string {
	return s.String()
}

// SetAmortizedBeginTime sets the AmortizedBeginTime field's value.
func (s *ListForListAmortizedCostBillDailyOutput) SetAmortizedBeginTime(v string) *ListForListAmortizedCostBillDailyOutput {
	s.AmortizedBeginTime = &v
	return s
}

// SetAmortizedDay sets the AmortizedDay field's value.
func (s *ListForListAmortizedCostBillDailyOutput) SetAmortizedDay(v string) *ListForListAmortizedCostBillDailyOutput {
	s.AmortizedDay = &v
	return s
}

// SetAmortizedEndTime sets the AmortizedEndTime field's value.
func (s *ListForListAmortizedCostBillDailyOutput) SetAmortizedEndTime(v string) *ListForListAmortizedCostBillDailyOutput {
	s.AmortizedEndTime = &v
	return s
}

// SetAmortizedMonth sets the AmortizedMonth field's value.
func (s *ListForListAmortizedCostBillDailyOutput) SetAmortizedMonth(v string) *ListForListAmortizedCostBillDailyOutput {
	s.AmortizedMonth = &v
	return s
}

// SetAmortizedType sets the AmortizedType field's value.
func (s *ListForListAmortizedCostBillDailyOutput) SetAmortizedType(v string) *ListForListAmortizedCostBillDailyOutput {
	s.AmortizedType = &v
	return s
}

// SetBillCategory sets the BillCategory field's value.
func (s *ListForListAmortizedCostBillDailyOutput) SetBillCategory(v string) *ListForListAmortizedCostBillDailyOutput {
	s.BillCategory = &v
	return s
}

// SetBillID sets the BillID field's value.
func (s *ListForListAmortizedCostBillDailyOutput) SetBillID(v string) *ListForListAmortizedCostBillDailyOutput {
	s.BillID = &v
	return s
}

// SetBillPeriod sets the BillPeriod field's value.
func (s *ListForListAmortizedCostBillDailyOutput) SetBillPeriod(v string) *ListForListAmortizedCostBillDailyOutput {
	s.BillPeriod = &v
	return s
}

// SetBillingMethodCode sets the BillingMethodCode field's value.
func (s *ListForListAmortizedCostBillDailyOutput) SetBillingMethodCode(v string) *ListForListAmortizedCostBillDailyOutput {
	s.BillingMethodCode = &v
	return s
}

// SetBillingMode sets the BillingMode field's value.
func (s *ListForListAmortizedCostBillDailyOutput) SetBillingMode(v string) *ListForListAmortizedCostBillDailyOutput {
	s.BillingMode = &v
	return s
}

// SetBusiPeriod sets the BusiPeriod field's value.
func (s *ListForListAmortizedCostBillDailyOutput) SetBusiPeriod(v string) *ListForListAmortizedCostBillDailyOutput {
	s.BusiPeriod = &v
	return s
}

// SetBusinessMode sets the BusinessMode field's value.
func (s *ListForListAmortizedCostBillDailyOutput) SetBusinessMode(v string) *ListForListAmortizedCostBillDailyOutput {
	s.BusinessMode = &v
	return s
}

// SetConfigName sets the ConfigName field's value.
func (s *ListForListAmortizedCostBillDailyOutput) SetConfigName(v string) *ListForListAmortizedCostBillDailyOutput {
	s.ConfigName = &v
	return s
}

// SetConfigurationCode sets the ConfigurationCode field's value.
func (s *ListForListAmortizedCostBillDailyOutput) SetConfigurationCode(v string) *ListForListAmortizedCostBillDailyOutput {
	s.ConfigurationCode = &v
	return s
}

// SetCount sets the Count field's value.
func (s *ListForListAmortizedCostBillDailyOutput) SetCount(v string) *ListForListAmortizedCostBillDailyOutput {
	s.Count = &v
	return s
}

// SetCouponAmount sets the CouponAmount field's value.
func (s *ListForListAmortizedCostBillDailyOutput) SetCouponAmount(v string) *ListForListAmortizedCostBillDailyOutput {
	s.CouponAmount = &v
	return s
}

// SetCurrency sets the Currency field's value.
func (s *ListForListAmortizedCostBillDailyOutput) SetCurrency(v string) *ListForListAmortizedCostBillDailyOutput {
	s.Currency = &v
	return s
}

// SetDailyAmortizedCouponAmount sets the DailyAmortizedCouponAmount field's value.
func (s *ListForListAmortizedCostBillDailyOutput) SetDailyAmortizedCouponAmount(v string) *ListForListAmortizedCostBillDailyOutput {
	s.DailyAmortizedCouponAmount = &v
	return s
}

// SetDailyAmortizedDiscountBillAmount sets the DailyAmortizedDiscountBillAmount field's value.
func (s *ListForListAmortizedCostBillDailyOutput) SetDailyAmortizedDiscountBillAmount(v string) *ListForListAmortizedCostBillDailyOutput {
	s.DailyAmortizedDiscountBillAmount = &v
	return s
}

// SetDailyAmortizedOriginalBillAmount sets the DailyAmortizedOriginalBillAmount field's value.
func (s *ListForListAmortizedCostBillDailyOutput) SetDailyAmortizedOriginalBillAmount(v string) *ListForListAmortizedCostBillDailyOutput {
	s.DailyAmortizedOriginalBillAmount = &v
	return s
}

// SetDailyAmortizedPaidAmount sets the DailyAmortizedPaidAmount field's value.
func (s *ListForListAmortizedCostBillDailyOutput) SetDailyAmortizedPaidAmount(v string) *ListForListAmortizedCostBillDailyOutput {
	s.DailyAmortizedPaidAmount = &v
	return s
}

// SetDailyAmortizedPayableAmount sets the DailyAmortizedPayableAmount field's value.
func (s *ListForListAmortizedCostBillDailyOutput) SetDailyAmortizedPayableAmount(v string) *ListForListAmortizedCostBillDailyOutput {
	s.DailyAmortizedPayableAmount = &v
	return s
}

// SetDailyAmortizedPreferentialBillAmount sets the DailyAmortizedPreferentialBillAmount field's value.
func (s *ListForListAmortizedCostBillDailyOutput) SetDailyAmortizedPreferentialBillAmount(v string) *ListForListAmortizedCostBillDailyOutput {
	s.DailyAmortizedPreferentialBillAmount = &v
	return s
}

// SetDailyAmortizedRoundAmount sets the DailyAmortizedRoundAmount field's value.
func (s *ListForListAmortizedCostBillDailyOutput) SetDailyAmortizedRoundAmount(v string) *ListForListAmortizedCostBillDailyOutput {
	s.DailyAmortizedRoundAmount = &v
	return s
}

// SetDiscountBillAmount sets the DiscountBillAmount field's value.
func (s *ListForListAmortizedCostBillDailyOutput) SetDiscountBillAmount(v string) *ListForListAmortizedCostBillDailyOutput {
	s.DiscountBillAmount = &v
	return s
}

// SetElement sets the Element field's value.
func (s *ListForListAmortizedCostBillDailyOutput) SetElement(v string) *ListForListAmortizedCostBillDailyOutput {
	s.Element = &v
	return s
}

// SetElementCode sets the ElementCode field's value.
func (s *ListForListAmortizedCostBillDailyOutput) SetElementCode(v string) *ListForListAmortizedCostBillDailyOutput {
	s.ElementCode = &v
	return s
}

// SetExpandField sets the ExpandField field's value.
func (s *ListForListAmortizedCostBillDailyOutput) SetExpandField(v string) *ListForListAmortizedCostBillDailyOutput {
	s.ExpandField = &v
	return s
}

// SetExpenseTime sets the ExpenseTime field's value.
func (s *ListForListAmortizedCostBillDailyOutput) SetExpenseTime(v string) *ListForListAmortizedCostBillDailyOutput {
	s.ExpenseTime = &v
	return s
}

// SetFactor sets the Factor field's value.
func (s *ListForListAmortizedCostBillDailyOutput) SetFactor(v string) *ListForListAmortizedCostBillDailyOutput {
	s.Factor = &v
	return s
}

// SetFactorCode sets the FactorCode field's value.
func (s *ListForListAmortizedCostBillDailyOutput) SetFactorCode(v string) *ListForListAmortizedCostBillDailyOutput {
	s.FactorCode = &v
	return s
}

// SetInstanceName sets the InstanceName field's value.
func (s *ListForListAmortizedCostBillDailyOutput) SetInstanceName(v string) *ListForListAmortizedCostBillDailyOutput {
	s.InstanceName = &v
	return s
}

// SetInstanceNo sets the InstanceNo field's value.
func (s *ListForListAmortizedCostBillDailyOutput) SetInstanceNo(v string) *ListForListAmortizedCostBillDailyOutput {
	s.InstanceNo = &v
	return s
}

// SetNowAmortizedCouponAmount sets the NowAmortizedCouponAmount field's value.
func (s *ListForListAmortizedCostBillDailyOutput) SetNowAmortizedCouponAmount(v string) *ListForListAmortizedCostBillDailyOutput {
	s.NowAmortizedCouponAmount = &v
	return s
}

// SetNowAmortizedDiscountBillAmount sets the NowAmortizedDiscountBillAmount field's value.
func (s *ListForListAmortizedCostBillDailyOutput) SetNowAmortizedDiscountBillAmount(v string) *ListForListAmortizedCostBillDailyOutput {
	s.NowAmortizedDiscountBillAmount = &v
	return s
}

// SetNowAmortizedOriginalBillAmount sets the NowAmortizedOriginalBillAmount field's value.
func (s *ListForListAmortizedCostBillDailyOutput) SetNowAmortizedOriginalBillAmount(v string) *ListForListAmortizedCostBillDailyOutput {
	s.NowAmortizedOriginalBillAmount = &v
	return s
}

// SetNowAmortizedPaidAmount sets the NowAmortizedPaidAmount field's value.
func (s *ListForListAmortizedCostBillDailyOutput) SetNowAmortizedPaidAmount(v string) *ListForListAmortizedCostBillDailyOutput {
	s.NowAmortizedPaidAmount = &v
	return s
}

// SetNowAmortizedPayableAmount sets the NowAmortizedPayableAmount field's value.
func (s *ListForListAmortizedCostBillDailyOutput) SetNowAmortizedPayableAmount(v string) *ListForListAmortizedCostBillDailyOutput {
	s.NowAmortizedPayableAmount = &v
	return s
}

// SetNowAmortizedPreferentialBillAmount sets the NowAmortizedPreferentialBillAmount field's value.
func (s *ListForListAmortizedCostBillDailyOutput) SetNowAmortizedPreferentialBillAmount(v string) *ListForListAmortizedCostBillDailyOutput {
	s.NowAmortizedPreferentialBillAmount = &v
	return s
}

// SetNowAmortizedRoundAmount sets the NowAmortizedRoundAmount field's value.
func (s *ListForListAmortizedCostBillDailyOutput) SetNowAmortizedRoundAmount(v string) *ListForListAmortizedCostBillDailyOutput {
	s.NowAmortizedRoundAmount = &v
	return s
}

// SetOriginalBillAmount sets the OriginalBillAmount field's value.
func (s *ListForListAmortizedCostBillDailyOutput) SetOriginalBillAmount(v string) *ListForListAmortizedCostBillDailyOutput {
	s.OriginalBillAmount = &v
	return s
}

// SetOwnerCustomerName sets the OwnerCustomerName field's value.
func (s *ListForListAmortizedCostBillDailyOutput) SetOwnerCustomerName(v string) *ListForListAmortizedCostBillDailyOutput {
	s.OwnerCustomerName = &v
	return s
}

// SetOwnerID sets the OwnerID field's value.
func (s *ListForListAmortizedCostBillDailyOutput) SetOwnerID(v string) *ListForListAmortizedCostBillDailyOutput {
	s.OwnerID = &v
	return s
}

// SetOwnerUserName sets the OwnerUserName field's value.
func (s *ListForListAmortizedCostBillDailyOutput) SetOwnerUserName(v string) *ListForListAmortizedCostBillDailyOutput {
	s.OwnerUserName = &v
	return s
}

// SetPaidAmount sets the PaidAmount field's value.
func (s *ListForListAmortizedCostBillDailyOutput) SetPaidAmount(v string) *ListForListAmortizedCostBillDailyOutput {
	s.PaidAmount = &v
	return s
}

// SetPayableAmount sets the PayableAmount field's value.
func (s *ListForListAmortizedCostBillDailyOutput) SetPayableAmount(v string) *ListForListAmortizedCostBillDailyOutput {
	s.PayableAmount = &v
	return s
}

// SetPayerCustomerName sets the PayerCustomerName field's value.
func (s *ListForListAmortizedCostBillDailyOutput) SetPayerCustomerName(v string) *ListForListAmortizedCostBillDailyOutput {
	s.PayerCustomerName = &v
	return s
}

// SetPayerID sets the PayerID field's value.
func (s *ListForListAmortizedCostBillDailyOutput) SetPayerID(v string) *ListForListAmortizedCostBillDailyOutput {
	s.PayerID = &v
	return s
}

// SetPayerUserName sets the PayerUserName field's value.
func (s *ListForListAmortizedCostBillDailyOutput) SetPayerUserName(v string) *ListForListAmortizedCostBillDailyOutput {
	s.PayerUserName = &v
	return s
}

// SetPreferentialBillAmount sets the PreferentialBillAmount field's value.
func (s *ListForListAmortizedCostBillDailyOutput) SetPreferentialBillAmount(v string) *ListForListAmortizedCostBillDailyOutput {
	s.PreferentialBillAmount = &v
	return s
}

// SetPrice sets the Price field's value.
func (s *ListForListAmortizedCostBillDailyOutput) SetPrice(v string) *ListForListAmortizedCostBillDailyOutput {
	s.Price = &v
	return s
}

// SetPriceUnit sets the PriceUnit field's value.
func (s *ListForListAmortizedCostBillDailyOutput) SetPriceUnit(v string) *ListForListAmortizedCostBillDailyOutput {
	s.PriceUnit = &v
	return s
}

// SetProduct sets the Product field's value.
func (s *ListForListAmortizedCostBillDailyOutput) SetProduct(v string) *ListForListAmortizedCostBillDailyOutput {
	s.Product = &v
	return s
}

// SetProductZh sets the ProductZh field's value.
func (s *ListForListAmortizedCostBillDailyOutput) SetProductZh(v string) *ListForListAmortizedCostBillDailyOutput {
	s.ProductZh = &v
	return s
}

// SetProject sets the Project field's value.
func (s *ListForListAmortizedCostBillDailyOutput) SetProject(v string) *ListForListAmortizedCostBillDailyOutput {
	s.Project = &v
	return s
}

// SetRegion sets the Region field's value.
func (s *ListForListAmortizedCostBillDailyOutput) SetRegion(v string) *ListForListAmortizedCostBillDailyOutput {
	s.Region = &v
	return s
}

// SetRegionCode sets the RegionCode field's value.
func (s *ListForListAmortizedCostBillDailyOutput) SetRegionCode(v string) *ListForListAmortizedCostBillDailyOutput {
	s.RegionCode = &v
	return s
}

// SetRoundAmount sets the RoundAmount field's value.
func (s *ListForListAmortizedCostBillDailyOutput) SetRoundAmount(v string) *ListForListAmortizedCostBillDailyOutput {
	s.RoundAmount = &v
	return s
}

// SetSellerCustomerName sets the SellerCustomerName field's value.
func (s *ListForListAmortizedCostBillDailyOutput) SetSellerCustomerName(v string) *ListForListAmortizedCostBillDailyOutput {
	s.SellerCustomerName = &v
	return s
}

// SetSellerID sets the SellerID field's value.
func (s *ListForListAmortizedCostBillDailyOutput) SetSellerID(v string) *ListForListAmortizedCostBillDailyOutput {
	s.SellerID = &v
	return s
}

// SetSellerUserName sets the SellerUserName field's value.
func (s *ListForListAmortizedCostBillDailyOutput) SetSellerUserName(v string) *ListForListAmortizedCostBillDailyOutput {
	s.SellerUserName = &v
	return s
}

// SetSubjectName sets the SubjectName field's value.
func (s *ListForListAmortizedCostBillDailyOutput) SetSubjectName(v string) *ListForListAmortizedCostBillDailyOutput {
	s.SubjectName = &v
	return s
}

// SetTag sets the Tag field's value.
func (s *ListForListAmortizedCostBillDailyOutput) SetTag(v string) *ListForListAmortizedCostBillDailyOutput {
	s.Tag = &v
	return s
}

// SetUnit sets the Unit field's value.
func (s *ListForListAmortizedCostBillDailyOutput) SetUnit(v string) *ListForListAmortizedCostBillDailyOutput {
	s.Unit = &v
	return s
}

// SetUseDuration sets the UseDuration field's value.
func (s *ListForListAmortizedCostBillDailyOutput) SetUseDuration(v string) *ListForListAmortizedCostBillDailyOutput {
	s.UseDuration = &v
	return s
}

// SetUseDurationUnit sets the UseDurationUnit field's value.
func (s *ListForListAmortizedCostBillDailyOutput) SetUseDurationUnit(v string) *ListForListAmortizedCostBillDailyOutput {
	s.UseDurationUnit = &v
	return s
}

// SetZone sets the Zone field's value.
func (s *ListForListAmortizedCostBillDailyOutput) SetZone(v string) *ListForListAmortizedCostBillDailyOutput {
	s.Zone = &v
	return s
}

// SetZoneCode sets the ZoneCode field's value.
func (s *ListForListAmortizedCostBillDailyOutput) SetZoneCode(v string) *ListForListAmortizedCostBillDailyOutput {
	s.ZoneCode = &v
	return s
}
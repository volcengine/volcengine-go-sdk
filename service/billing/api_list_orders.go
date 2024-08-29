// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package billing

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opListOrdersCommon = "ListOrders"

// ListOrdersCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ListOrdersCommon operation. The "output" return
// value will be populated with the ListOrdersCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListOrdersCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListOrdersCommon Send returns without error.
//
// See ListOrdersCommon for more information on using the ListOrdersCommon
// API call, and error handling.
//
//    // Example sending a request using the ListOrdersCommonRequest method.
//    req, resp := client.ListOrdersCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *BILLING) ListOrdersCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opListOrdersCommon,
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

// ListOrdersCommon API operation for BILLING.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for BILLING's
// API operation ListOrdersCommon for usage and error information.
func (c *BILLING) ListOrdersCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ListOrdersCommonRequest(input)
	return out, req.Send()
}

// ListOrdersCommonWithContext is the same as ListOrdersCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ListOrdersCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *BILLING) ListOrdersCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ListOrdersCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opListOrders = "ListOrders"

// ListOrdersRequest generates a "volcengine/request.Request" representing the
// client's request for the ListOrders operation. The "output" return
// value will be populated with the ListOrdersCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListOrdersCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListOrdersCommon Send returns without error.
//
// See ListOrders for more information on using the ListOrders
// API call, and error handling.
//
//    // Example sending a request using the ListOrdersRequest method.
//    req, resp := client.ListOrdersRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *BILLING) ListOrdersRequest(input *ListOrdersInput) (req *request.Request, output *ListOrdersOutput) {
	op := &request.Operation{
		Name:       opListOrders,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListOrdersInput{}
	}

	output = &ListOrdersOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// ListOrders API operation for BILLING.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for BILLING's
// API operation ListOrders for usage and error information.
func (c *BILLING) ListOrders(input *ListOrdersInput) (*ListOrdersOutput, error) {
	req, out := c.ListOrdersRequest(input)
	return out, req.Send()
}

// ListOrdersWithContext is the same as ListOrders with the addition of
// the ability to pass a context and additional request options.
//
// See ListOrders for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *BILLING) ListOrdersWithContext(ctx volcengine.Context, input *ListOrdersInput, opts ...request.Option) (*ListOrdersOutput, error) {
	req, out := c.ListOrdersRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ListOrdersInput struct {
	_ struct{} `type:"structure"`

	// CreateTimeEnd is a required field
	CreateTimeEnd *string `type:"string" required:"true"`

	// CreateTimeStart is a required field
	CreateTimeStart *string `type:"string" required:"true"`

	// MaxResults is a required field
	MaxResults *int32 `type:"int32" required:"true"`

	// NextToken is a required field
	NextToken *string `type:"string" required:"true"`

	// OrderType is a required field
	OrderType *string `type:"string" required:"true"`

	// Status is a required field
	Status *string `type:"string" required:"true"`
}

// String returns the string representation
func (s ListOrdersInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListOrdersInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListOrdersInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ListOrdersInput"}
	if s.CreateTimeEnd == nil {
		invalidParams.Add(request.NewErrParamRequired("CreateTimeEnd"))
	}
	if s.CreateTimeStart == nil {
		invalidParams.Add(request.NewErrParamRequired("CreateTimeStart"))
	}
	if s.MaxResults == nil {
		invalidParams.Add(request.NewErrParamRequired("MaxResults"))
	}
	if s.NextToken == nil {
		invalidParams.Add(request.NewErrParamRequired("NextToken"))
	}
	if s.OrderType == nil {
		invalidParams.Add(request.NewErrParamRequired("OrderType"))
	}
	if s.Status == nil {
		invalidParams.Add(request.NewErrParamRequired("Status"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetCreateTimeEnd sets the CreateTimeEnd field's value.
func (s *ListOrdersInput) SetCreateTimeEnd(v string) *ListOrdersInput {
	s.CreateTimeEnd = &v
	return s
}

// SetCreateTimeStart sets the CreateTimeStart field's value.
func (s *ListOrdersInput) SetCreateTimeStart(v string) *ListOrdersInput {
	s.CreateTimeStart = &v
	return s
}

// SetMaxResults sets the MaxResults field's value.
func (s *ListOrdersInput) SetMaxResults(v int32) *ListOrdersInput {
	s.MaxResults = &v
	return s
}

// SetNextToken sets the NextToken field's value.
func (s *ListOrdersInput) SetNextToken(v string) *ListOrdersInput {
	s.NextToken = &v
	return s
}

// SetOrderType sets the OrderType field's value.
func (s *ListOrdersInput) SetOrderType(v string) *ListOrdersInput {
	s.OrderType = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *ListOrdersInput) SetStatus(v string) *ListOrdersInput {
	s.Status = &v
	return s
}

type ListOrdersOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	MaxResults *int32 `type:"int32"`

	NextToken *string `type:"string"`

	OrderInfos []*OrderInfoForListOrdersOutput `type:"list"`
}

// String returns the string representation
func (s ListOrdersOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListOrdersOutput) GoString() string {
	return s.String()
}

// SetMaxResults sets the MaxResults field's value.
func (s *ListOrdersOutput) SetMaxResults(v int32) *ListOrdersOutput {
	s.MaxResults = &v
	return s
}

// SetNextToken sets the NextToken field's value.
func (s *ListOrdersOutput) SetNextToken(v string) *ListOrdersOutput {
	s.NextToken = &v
	return s
}

// SetOrderInfos sets the OrderInfos field's value.
func (s *ListOrdersOutput) SetOrderInfos(v []*OrderInfoForListOrdersOutput) *ListOrdersOutput {
	s.OrderInfos = v
	return s
}

type OrderInfoForListOrdersOutput struct {
	_ struct{} `type:"structure"`

	BuyerCustomerName *string `type:"string"`

	BuyerID *int32 `type:"int32"`

	CouponAmount *string `type:"string"`

	CreateTime *string `type:"string"`

	DiscountAmount *string `type:"string"`

	OrderID *string `type:"string"`

	OrderType *string `type:"string"`

	OriginalAmount *string `type:"string"`

	PaidAmount *string `type:"string"`

	PayableAmount *string `type:"string"`

	PayerCustomerName *string `type:"string"`

	PayerID *int32 `type:"int32"`

	SellerCustomerName *string `type:"string"`

	SellerID *int32 `type:"int32"`

	Status *string `type:"string"`

	SubjectNo *string `type:"string"`
}

// String returns the string representation
func (s OrderInfoForListOrdersOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s OrderInfoForListOrdersOutput) GoString() string {
	return s.String()
}

// SetBuyerCustomerName sets the BuyerCustomerName field's value.
func (s *OrderInfoForListOrdersOutput) SetBuyerCustomerName(v string) *OrderInfoForListOrdersOutput {
	s.BuyerCustomerName = &v
	return s
}

// SetBuyerID sets the BuyerID field's value.
func (s *OrderInfoForListOrdersOutput) SetBuyerID(v int32) *OrderInfoForListOrdersOutput {
	s.BuyerID = &v
	return s
}

// SetCouponAmount sets the CouponAmount field's value.
func (s *OrderInfoForListOrdersOutput) SetCouponAmount(v string) *OrderInfoForListOrdersOutput {
	s.CouponAmount = &v
	return s
}

// SetCreateTime sets the CreateTime field's value.
func (s *OrderInfoForListOrdersOutput) SetCreateTime(v string) *OrderInfoForListOrdersOutput {
	s.CreateTime = &v
	return s
}

// SetDiscountAmount sets the DiscountAmount field's value.
func (s *OrderInfoForListOrdersOutput) SetDiscountAmount(v string) *OrderInfoForListOrdersOutput {
	s.DiscountAmount = &v
	return s
}

// SetOrderID sets the OrderID field's value.
func (s *OrderInfoForListOrdersOutput) SetOrderID(v string) *OrderInfoForListOrdersOutput {
	s.OrderID = &v
	return s
}

// SetOrderType sets the OrderType field's value.
func (s *OrderInfoForListOrdersOutput) SetOrderType(v string) *OrderInfoForListOrdersOutput {
	s.OrderType = &v
	return s
}

// SetOriginalAmount sets the OriginalAmount field's value.
func (s *OrderInfoForListOrdersOutput) SetOriginalAmount(v string) *OrderInfoForListOrdersOutput {
	s.OriginalAmount = &v
	return s
}

// SetPaidAmount sets the PaidAmount field's value.
func (s *OrderInfoForListOrdersOutput) SetPaidAmount(v string) *OrderInfoForListOrdersOutput {
	s.PaidAmount = &v
	return s
}

// SetPayableAmount sets the PayableAmount field's value.
func (s *OrderInfoForListOrdersOutput) SetPayableAmount(v string) *OrderInfoForListOrdersOutput {
	s.PayableAmount = &v
	return s
}

// SetPayerCustomerName sets the PayerCustomerName field's value.
func (s *OrderInfoForListOrdersOutput) SetPayerCustomerName(v string) *OrderInfoForListOrdersOutput {
	s.PayerCustomerName = &v
	return s
}

// SetPayerID sets the PayerID field's value.
func (s *OrderInfoForListOrdersOutput) SetPayerID(v int32) *OrderInfoForListOrdersOutput {
	s.PayerID = &v
	return s
}

// SetSellerCustomerName sets the SellerCustomerName field's value.
func (s *OrderInfoForListOrdersOutput) SetSellerCustomerName(v string) *OrderInfoForListOrdersOutput {
	s.SellerCustomerName = &v
	return s
}

// SetSellerID sets the SellerID field's value.
func (s *OrderInfoForListOrdersOutput) SetSellerID(v int32) *OrderInfoForListOrdersOutput {
	s.SellerID = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *OrderInfoForListOrdersOutput) SetStatus(v string) *OrderInfoForListOrdersOutput {
	s.Status = &v
	return s
}

// SetSubjectNo sets the SubjectNo field's value.
func (s *OrderInfoForListOrdersOutput) SetSubjectNo(v string) *OrderInfoForListOrdersOutput {
	s.SubjectNo = &v
	return s
}
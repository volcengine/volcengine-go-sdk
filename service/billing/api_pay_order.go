// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package billing

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opPayOrderCommon = "PayOrder"

// PayOrderCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the PayOrderCommon operation. The "output" return
// value will be populated with the PayOrderCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned PayOrderCommon Request to send the API call to the service.
// the "output" return value is not valid until after PayOrderCommon Send returns without error.
//
// See PayOrderCommon for more information on using the PayOrderCommon
// API call, and error handling.
//
//    // Example sending a request using the PayOrderCommonRequest method.
//    req, resp := client.PayOrderCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *BILLING) PayOrderCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opPayOrderCommon,
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

// PayOrderCommon API operation for BILLING.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for BILLING's
// API operation PayOrderCommon for usage and error information.
func (c *BILLING) PayOrderCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.PayOrderCommonRequest(input)
	return out, req.Send()
}

// PayOrderCommonWithContext is the same as PayOrderCommon with the addition of
// the ability to pass a context and additional request options.
//
// See PayOrderCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *BILLING) PayOrderCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.PayOrderCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opPayOrder = "PayOrder"

// PayOrderRequest generates a "volcengine/request.Request" representing the
// client's request for the PayOrder operation. The "output" return
// value will be populated with the PayOrderCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned PayOrderCommon Request to send the API call to the service.
// the "output" return value is not valid until after PayOrderCommon Send returns without error.
//
// See PayOrder for more information on using the PayOrder
// API call, and error handling.
//
//    // Example sending a request using the PayOrderRequest method.
//    req, resp := client.PayOrderRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *BILLING) PayOrderRequest(input *PayOrderInput) (req *request.Request, output *PayOrderOutput) {
	op := &request.Operation{
		Name:       opPayOrder,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &PayOrderInput{}
	}

	output = &PayOrderOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// PayOrder API operation for BILLING.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for BILLING's
// API operation PayOrder for usage and error information.
func (c *BILLING) PayOrder(input *PayOrderInput) (*PayOrderOutput, error) {
	req, out := c.PayOrderRequest(input)
	return out, req.Send()
}

// PayOrderWithContext is the same as PayOrder with the addition of
// the ability to pass a context and additional request options.
//
// See PayOrder for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *BILLING) PayOrderWithContext(ctx volcengine.Context, input *PayOrderInput, opts ...request.Option) (*PayOrderOutput, error) {
	req, out := c.PayOrderRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type PayOrderInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	// OrderID is a required field
	OrderID *string `type:"string" json:",omitempty" required:"true"`
}

// String returns the string representation
func (s PayOrderInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s PayOrderInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PayOrderInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "PayOrderInput"}
	if s.OrderID == nil {
		invalidParams.Add(request.NewErrParamRequired("OrderID"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetOrderID sets the OrderID field's value.
func (s *PayOrderInput) SetOrderID(v string) *PayOrderInput {
	s.OrderID = &v
	return s
}

type PayOrderOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s PayOrderOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s PayOrderOutput) GoString() string {
	return s.String()
}

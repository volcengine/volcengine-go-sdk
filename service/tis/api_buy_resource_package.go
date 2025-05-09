// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package tis

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opBuyResourcePackageCommon = "BuyResourcePackage"

// BuyResourcePackageCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the BuyResourcePackageCommon operation. The "output" return
// value will be populated with the BuyResourcePackageCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned BuyResourcePackageCommon Request to send the API call to the service.
// the "output" return value is not valid until after BuyResourcePackageCommon Send returns without error.
//
// See BuyResourcePackageCommon for more information on using the BuyResourcePackageCommon
// API call, and error handling.
//
//    // Example sending a request using the BuyResourcePackageCommonRequest method.
//    req, resp := client.BuyResourcePackageCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *TIS) BuyResourcePackageCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opBuyResourcePackageCommon,
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

// BuyResourcePackageCommon API operation for TIS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for TIS's
// API operation BuyResourcePackageCommon for usage and error information.
func (c *TIS) BuyResourcePackageCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.BuyResourcePackageCommonRequest(input)
	return out, req.Send()
}

// BuyResourcePackageCommonWithContext is the same as BuyResourcePackageCommon with the addition of
// the ability to pass a context and additional request options.
//
// See BuyResourcePackageCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *TIS) BuyResourcePackageCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.BuyResourcePackageCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opBuyResourcePackage = "BuyResourcePackage"

// BuyResourcePackageRequest generates a "volcengine/request.Request" representing the
// client's request for the BuyResourcePackage operation. The "output" return
// value will be populated with the BuyResourcePackageCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned BuyResourcePackageCommon Request to send the API call to the service.
// the "output" return value is not valid until after BuyResourcePackageCommon Send returns without error.
//
// See BuyResourcePackage for more information on using the BuyResourcePackage
// API call, and error handling.
//
//    // Example sending a request using the BuyResourcePackageRequest method.
//    req, resp := client.BuyResourcePackageRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *TIS) BuyResourcePackageRequest(input *BuyResourcePackageInput) (req *request.Request, output *BuyResourcePackageOutput) {
	op := &request.Operation{
		Name:       opBuyResourcePackage,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &BuyResourcePackageInput{}
	}

	output = &BuyResourcePackageOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// BuyResourcePackage API operation for TIS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for TIS's
// API operation BuyResourcePackage for usage and error information.
func (c *TIS) BuyResourcePackage(input *BuyResourcePackageInput) (*BuyResourcePackageOutput, error) {
	req, out := c.BuyResourcePackageRequest(input)
	return out, req.Send()
}

// BuyResourcePackageWithContext is the same as BuyResourcePackage with the addition of
// the ability to pass a context and additional request options.
//
// See BuyResourcePackage for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *TIS) BuyResourcePackageWithContext(ctx volcengine.Context, input *BuyResourcePackageInput, opts ...request.Option) (*BuyResourcePackageOutput, error) {
	req, out := c.BuyResourcePackageRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type BuyResourcePackageInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	// DeviceName is a required field
	DeviceName *string `type:"string" json:"deviceName,omitempty" required:"true"`

	// Item is a required field
	Item *string `type:"string" json:"item,omitempty" required:"true"`

	// ProductKey is a required field
	ProductKey *string `type:"string" json:"productKey,omitempty" required:"true"`
}

// String returns the string representation
func (s BuyResourcePackageInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s BuyResourcePackageInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *BuyResourcePackageInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "BuyResourcePackageInput"}
	if s.DeviceName == nil {
		invalidParams.Add(request.NewErrParamRequired("DeviceName"))
	}
	if s.Item == nil {
		invalidParams.Add(request.NewErrParamRequired("Item"))
	}
	if s.ProductKey == nil {
		invalidParams.Add(request.NewErrParamRequired("ProductKey"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetDeviceName sets the DeviceName field's value.
func (s *BuyResourcePackageInput) SetDeviceName(v string) *BuyResourcePackageInput {
	s.DeviceName = &v
	return s
}

// SetItem sets the Item field's value.
func (s *BuyResourcePackageInput) SetItem(v string) *BuyResourcePackageInput {
	s.Item = &v
	return s
}

// SetProductKey sets the ProductKey field's value.
func (s *BuyResourcePackageInput) SetProductKey(v string) *BuyResourcePackageInput {
	s.ProductKey = &v
	return s
}

type BuyResourcePackageOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	OrderNo *string `type:"string" json:"orderNo,omitempty"`

	Success *bool `type:"boolean" json:"success,omitempty"`
}

// String returns the string representation
func (s BuyResourcePackageOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s BuyResourcePackageOutput) GoString() string {
	return s.String()
}

// SetOrderNo sets the OrderNo field's value.
func (s *BuyResourcePackageOutput) SetOrderNo(v string) *BuyResourcePackageOutput {
	s.OrderNo = &v
	return s
}

// SetSuccess sets the Success field's value.
func (s *BuyResourcePackageOutput) SetSuccess(v bool) *BuyResourcePackageOutput {
	s.Success = &v
	return s
}

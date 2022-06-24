// Code generated by volcstack with private/model/cli/gen-api/main.go. DO NOT EDIT.

package cen

import (
	"github.com/volcengine/volcstack-go-sdk/volcstack"
	"github.com/volcengine/volcstack-go-sdk/volcstack/request"
	"github.com/volcengine/volcstack-go-sdk/volcstack/response"
	"github.com/volcengine/volcstack-go-sdk/volcstack/volcstackutil"
)

const opRenewCenBandwidthPackageCommon = "RenewCenBandwidthPackage"

// RenewCenBandwidthPackageCommonRequest generates a "volcstack/request.Request" representing the
// client's request for the RenewCenBandwidthPackageCommon operation. The "output" return
// value will be populated with the RenewCenBandwidthPackageCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned RenewCenBandwidthPackageCommon Request to send the API call to the service.
// the "output" return value is not valid until after RenewCenBandwidthPackageCommon Send returns without error.
//
// See RenewCenBandwidthPackageCommon for more information on using the RenewCenBandwidthPackageCommon
// API call, and error handling.
//
//    // Example sending a request using the RenewCenBandwidthPackageCommonRequest method.
//    req, resp := client.RenewCenBandwidthPackageCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CEN) RenewCenBandwidthPackageCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opRenewCenBandwidthPackageCommon,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &map[string]interface{}{}
	}

	output = &map[string]interface{}{}
	req = c.newRequest(op, input, output)

	return
}

// RenewCenBandwidthPackageCommon API operation for CEN.
//
// Returns volcstackerr.Error for service API and SDK errors. Use runtime type assertions
// with volcstackerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCSTACK API reference guide for CEN's
// API operation RenewCenBandwidthPackageCommon for usage and error information.
func (c *CEN) RenewCenBandwidthPackageCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.RenewCenBandwidthPackageCommonRequest(input)
	return out, req.Send()
}

// RenewCenBandwidthPackageCommonWithContext is the same as RenewCenBandwidthPackageCommon with the addition of
// the ability to pass a context and additional request options.
//
// See RenewCenBandwidthPackageCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CEN) RenewCenBandwidthPackageCommonWithContext(ctx volcstack.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.RenewCenBandwidthPackageCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opRenewCenBandwidthPackage = "RenewCenBandwidthPackage"

// RenewCenBandwidthPackageRequest generates a "volcstack/request.Request" representing the
// client's request for the RenewCenBandwidthPackage operation. The "output" return
// value will be populated with the RenewCenBandwidthPackageCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned RenewCenBandwidthPackageCommon Request to send the API call to the service.
// the "output" return value is not valid until after RenewCenBandwidthPackageCommon Send returns without error.
//
// See RenewCenBandwidthPackage for more information on using the RenewCenBandwidthPackage
// API call, and error handling.
//
//    // Example sending a request using the RenewCenBandwidthPackageRequest method.
//    req, resp := client.RenewCenBandwidthPackageRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CEN) RenewCenBandwidthPackageRequest(input *RenewCenBandwidthPackageInput) (req *request.Request, output *RenewCenBandwidthPackageOutput) {
	op := &request.Operation{
		Name:       opRenewCenBandwidthPackage,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &RenewCenBandwidthPackageInput{}
	}

	output = &RenewCenBandwidthPackageOutput{}
	req = c.newRequest(op, input, output)

	return
}

// RenewCenBandwidthPackage API operation for CEN.
//
// Returns volcstackerr.Error for service API and SDK errors. Use runtime type assertions
// with volcstackerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCSTACK API reference guide for CEN's
// API operation RenewCenBandwidthPackage for usage and error information.
func (c *CEN) RenewCenBandwidthPackage(input *RenewCenBandwidthPackageInput) (*RenewCenBandwidthPackageOutput, error) {
	req, out := c.RenewCenBandwidthPackageRequest(input)
	return out, req.Send()
}

// RenewCenBandwidthPackageWithContext is the same as RenewCenBandwidthPackage with the addition of
// the ability to pass a context and additional request options.
//
// See RenewCenBandwidthPackage for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CEN) RenewCenBandwidthPackageWithContext(ctx volcstack.Context, input *RenewCenBandwidthPackageInput, opts ...request.Option) (*RenewCenBandwidthPackageOutput, error) {
	req, out := c.RenewCenBandwidthPackageRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type RenewCenBandwidthPackageInput struct {
	_ struct{} `type:"structure"`

	// CenBandwidthPackageId is a required field
	CenBandwidthPackageId *string `type:"string" required:"true"`

	Period *int64 `type:"integer"`

	PeriodUnit *string `type:"string" enum:"PeriodUnitForRenewCenBandwidthPackageInput"`
}

// String returns the string representation
func (s RenewCenBandwidthPackageInput) String() string {
	return volcstackutil.Prettify(s)
}

// GoString returns the string representation
func (s RenewCenBandwidthPackageInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RenewCenBandwidthPackageInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "RenewCenBandwidthPackageInput"}
	if s.CenBandwidthPackageId == nil {
		invalidParams.Add(request.NewErrParamRequired("CenBandwidthPackageId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetCenBandwidthPackageId sets the CenBandwidthPackageId field's value.
func (s *RenewCenBandwidthPackageInput) SetCenBandwidthPackageId(v string) *RenewCenBandwidthPackageInput {
	s.CenBandwidthPackageId = &v
	return s
}

// SetPeriod sets the Period field's value.
func (s *RenewCenBandwidthPackageInput) SetPeriod(v int64) *RenewCenBandwidthPackageInput {
	s.Period = &v
	return s
}

// SetPeriodUnit sets the PeriodUnit field's value.
func (s *RenewCenBandwidthPackageInput) SetPeriodUnit(v string) *RenewCenBandwidthPackageInput {
	s.PeriodUnit = &v
	return s
}

type RenewCenBandwidthPackageOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s RenewCenBandwidthPackageOutput) String() string {
	return volcstackutil.Prettify(s)
}

// GoString returns the string representation
func (s RenewCenBandwidthPackageOutput) GoString() string {
	return s.String()
}

const (
	// PeriodUnitForRenewCenBandwidthPackageInputMonth is a PeriodUnitForRenewCenBandwidthPackageInput enum value
	PeriodUnitForRenewCenBandwidthPackageInputMonth = "Month"

	// PeriodUnitForRenewCenBandwidthPackageInputYear is a PeriodUnitForRenewCenBandwidthPackageInput enum value
	PeriodUnitForRenewCenBandwidthPackageInputYear = "Year"
)
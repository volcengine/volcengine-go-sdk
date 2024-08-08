// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package transitrouter

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribeTransitRouterBandwidthPackagesBillingCommon = "DescribeTransitRouterBandwidthPackagesBilling"

// DescribeTransitRouterBandwidthPackagesBillingCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeTransitRouterBandwidthPackagesBillingCommon operation. The "output" return
// value will be populated with the DescribeTransitRouterBandwidthPackagesBillingCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeTransitRouterBandwidthPackagesBillingCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeTransitRouterBandwidthPackagesBillingCommon Send returns without error.
//
// See DescribeTransitRouterBandwidthPackagesBillingCommon for more information on using the DescribeTransitRouterBandwidthPackagesBillingCommon
// API call, and error handling.
//
//    // Example sending a request using the DescribeTransitRouterBandwidthPackagesBillingCommonRequest method.
//    req, resp := client.DescribeTransitRouterBandwidthPackagesBillingCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *TRANSITROUTER) DescribeTransitRouterBandwidthPackagesBillingCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeTransitRouterBandwidthPackagesBillingCommon,
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

// DescribeTransitRouterBandwidthPackagesBillingCommon API operation for TRANSITROUTER.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for TRANSITROUTER's
// API operation DescribeTransitRouterBandwidthPackagesBillingCommon for usage and error information.
func (c *TRANSITROUTER) DescribeTransitRouterBandwidthPackagesBillingCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeTransitRouterBandwidthPackagesBillingCommonRequest(input)
	return out, req.Send()
}

// DescribeTransitRouterBandwidthPackagesBillingCommonWithContext is the same as DescribeTransitRouterBandwidthPackagesBillingCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeTransitRouterBandwidthPackagesBillingCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *TRANSITROUTER) DescribeTransitRouterBandwidthPackagesBillingCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeTransitRouterBandwidthPackagesBillingCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeTransitRouterBandwidthPackagesBilling = "DescribeTransitRouterBandwidthPackagesBilling"

// DescribeTransitRouterBandwidthPackagesBillingRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeTransitRouterBandwidthPackagesBilling operation. The "output" return
// value will be populated with the DescribeTransitRouterBandwidthPackagesBillingCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeTransitRouterBandwidthPackagesBillingCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeTransitRouterBandwidthPackagesBillingCommon Send returns without error.
//
// See DescribeTransitRouterBandwidthPackagesBilling for more information on using the DescribeTransitRouterBandwidthPackagesBilling
// API call, and error handling.
//
//    // Example sending a request using the DescribeTransitRouterBandwidthPackagesBillingRequest method.
//    req, resp := client.DescribeTransitRouterBandwidthPackagesBillingRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *TRANSITROUTER) DescribeTransitRouterBandwidthPackagesBillingRequest(input *DescribeTransitRouterBandwidthPackagesBillingInput) (req *request.Request, output *DescribeTransitRouterBandwidthPackagesBillingOutput) {
	op := &request.Operation{
		Name:       opDescribeTransitRouterBandwidthPackagesBilling,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeTransitRouterBandwidthPackagesBillingInput{}
	}

	output = &DescribeTransitRouterBandwidthPackagesBillingOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DescribeTransitRouterBandwidthPackagesBilling API operation for TRANSITROUTER.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for TRANSITROUTER's
// API operation DescribeTransitRouterBandwidthPackagesBilling for usage and error information.
func (c *TRANSITROUTER) DescribeTransitRouterBandwidthPackagesBilling(input *DescribeTransitRouterBandwidthPackagesBillingInput) (*DescribeTransitRouterBandwidthPackagesBillingOutput, error) {
	req, out := c.DescribeTransitRouterBandwidthPackagesBillingRequest(input)
	return out, req.Send()
}

// DescribeTransitRouterBandwidthPackagesBillingWithContext is the same as DescribeTransitRouterBandwidthPackagesBilling with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeTransitRouterBandwidthPackagesBilling for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *TRANSITROUTER) DescribeTransitRouterBandwidthPackagesBillingWithContext(ctx volcengine.Context, input *DescribeTransitRouterBandwidthPackagesBillingInput, opts ...request.Option) (*DescribeTransitRouterBandwidthPackagesBillingOutput, error) {
	req, out := c.DescribeTransitRouterBandwidthPackagesBillingRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DescribeTransitRouterBandwidthPackagesBillingInput struct {
	_ struct{} `type:"structure"`

	PageNumber *int32 `type:"int32"`

	PageSize *int32 `type:"int32"`

	TransitRouterBandwidthPackageIds []*string `type:"list"`
}

// String returns the string representation
func (s DescribeTransitRouterBandwidthPackagesBillingInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeTransitRouterBandwidthPackagesBillingInput) GoString() string {
	return s.String()
}

// SetPageNumber sets the PageNumber field's value.
func (s *DescribeTransitRouterBandwidthPackagesBillingInput) SetPageNumber(v int32) *DescribeTransitRouterBandwidthPackagesBillingInput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *DescribeTransitRouterBandwidthPackagesBillingInput) SetPageSize(v int32) *DescribeTransitRouterBandwidthPackagesBillingInput {
	s.PageSize = &v
	return s
}

// SetTransitRouterBandwidthPackageIds sets the TransitRouterBandwidthPackageIds field's value.
func (s *DescribeTransitRouterBandwidthPackagesBillingInput) SetTransitRouterBandwidthPackageIds(v []*string) *DescribeTransitRouterBandwidthPackagesBillingInput {
	s.TransitRouterBandwidthPackageIds = v
	return s
}

type DescribeTransitRouterBandwidthPackagesBillingOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	PageNumber *int32 `type:"int32"`

	PageSize *int32 `type:"int32"`

	TotalCount *int32 `type:"int32"`

	TransitRouterBandwidthPackages []*TransitRouterBandwidthPackageForDescribeTransitRouterBandwidthPackagesBillingOutput `type:"list"`
}

// String returns the string representation
func (s DescribeTransitRouterBandwidthPackagesBillingOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeTransitRouterBandwidthPackagesBillingOutput) GoString() string {
	return s.String()
}

// SetPageNumber sets the PageNumber field's value.
func (s *DescribeTransitRouterBandwidthPackagesBillingOutput) SetPageNumber(v int32) *DescribeTransitRouterBandwidthPackagesBillingOutput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *DescribeTransitRouterBandwidthPackagesBillingOutput) SetPageSize(v int32) *DescribeTransitRouterBandwidthPackagesBillingOutput {
	s.PageSize = &v
	return s
}

// SetTotalCount sets the TotalCount field's value.
func (s *DescribeTransitRouterBandwidthPackagesBillingOutput) SetTotalCount(v int32) *DescribeTransitRouterBandwidthPackagesBillingOutput {
	s.TotalCount = &v
	return s
}

// SetTransitRouterBandwidthPackages sets the TransitRouterBandwidthPackages field's value.
func (s *DescribeTransitRouterBandwidthPackagesBillingOutput) SetTransitRouterBandwidthPackages(v []*TransitRouterBandwidthPackageForDescribeTransitRouterBandwidthPackagesBillingOutput) *DescribeTransitRouterBandwidthPackagesBillingOutput {
	s.TransitRouterBandwidthPackages = v
	return s
}

type TransitRouterBandwidthPackageForDescribeTransitRouterBandwidthPackagesBillingOutput struct {
	_ struct{} `type:"structure"`

	BillingStatus *int32 `type:"int32"`

	BillingType *int32 `type:"int32"`

	ExpiredTime *string `type:"string"`

	ReclaimTime *string `type:"string"`

	RemainRenewTimes *int32 `type:"int32"`

	RenewType *string `type:"string"`

	TransitRouterBandwidthPackageId *string `type:"string"`
}

// String returns the string representation
func (s TransitRouterBandwidthPackageForDescribeTransitRouterBandwidthPackagesBillingOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s TransitRouterBandwidthPackageForDescribeTransitRouterBandwidthPackagesBillingOutput) GoString() string {
	return s.String()
}

// SetBillingStatus sets the BillingStatus field's value.
func (s *TransitRouterBandwidthPackageForDescribeTransitRouterBandwidthPackagesBillingOutput) SetBillingStatus(v int32) *TransitRouterBandwidthPackageForDescribeTransitRouterBandwidthPackagesBillingOutput {
	s.BillingStatus = &v
	return s
}

// SetBillingType sets the BillingType field's value.
func (s *TransitRouterBandwidthPackageForDescribeTransitRouterBandwidthPackagesBillingOutput) SetBillingType(v int32) *TransitRouterBandwidthPackageForDescribeTransitRouterBandwidthPackagesBillingOutput {
	s.BillingType = &v
	return s
}

// SetExpiredTime sets the ExpiredTime field's value.
func (s *TransitRouterBandwidthPackageForDescribeTransitRouterBandwidthPackagesBillingOutput) SetExpiredTime(v string) *TransitRouterBandwidthPackageForDescribeTransitRouterBandwidthPackagesBillingOutput {
	s.ExpiredTime = &v
	return s
}

// SetReclaimTime sets the ReclaimTime field's value.
func (s *TransitRouterBandwidthPackageForDescribeTransitRouterBandwidthPackagesBillingOutput) SetReclaimTime(v string) *TransitRouterBandwidthPackageForDescribeTransitRouterBandwidthPackagesBillingOutput {
	s.ReclaimTime = &v
	return s
}

// SetRemainRenewTimes sets the RemainRenewTimes field's value.
func (s *TransitRouterBandwidthPackageForDescribeTransitRouterBandwidthPackagesBillingOutput) SetRemainRenewTimes(v int32) *TransitRouterBandwidthPackageForDescribeTransitRouterBandwidthPackagesBillingOutput {
	s.RemainRenewTimes = &v
	return s
}

// SetRenewType sets the RenewType field's value.
func (s *TransitRouterBandwidthPackageForDescribeTransitRouterBandwidthPackagesBillingOutput) SetRenewType(v string) *TransitRouterBandwidthPackageForDescribeTransitRouterBandwidthPackagesBillingOutput {
	s.RenewType = &v
	return s
}

// SetTransitRouterBandwidthPackageId sets the TransitRouterBandwidthPackageId field's value.
func (s *TransitRouterBandwidthPackageForDescribeTransitRouterBandwidthPackagesBillingOutput) SetTransitRouterBandwidthPackageId(v string) *TransitRouterBandwidthPackageForDescribeTransitRouterBandwidthPackagesBillingOutput {
	s.TransitRouterBandwidthPackageId = &v
	return s
}
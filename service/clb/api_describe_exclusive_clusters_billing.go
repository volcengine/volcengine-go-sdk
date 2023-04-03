// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package clb

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribeExclusiveClustersBillingCommon = "DescribeExclusiveClustersBilling"

// DescribeExclusiveClustersBillingCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeExclusiveClustersBillingCommon operation. The "output" return
// value will be populated with the DescribeExclusiveClustersBillingCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeExclusiveClustersBillingCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeExclusiveClustersBillingCommon Send returns without error.
//
// See DescribeExclusiveClustersBillingCommon for more information on using the DescribeExclusiveClustersBillingCommon
// API call, and error handling.
//
//    // Example sending a request using the DescribeExclusiveClustersBillingCommonRequest method.
//    req, resp := client.DescribeExclusiveClustersBillingCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CLB) DescribeExclusiveClustersBillingCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeExclusiveClustersBillingCommon,
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

// DescribeExclusiveClustersBillingCommon API operation for CLB.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CLB's
// API operation DescribeExclusiveClustersBillingCommon for usage and error information.
func (c *CLB) DescribeExclusiveClustersBillingCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeExclusiveClustersBillingCommonRequest(input)
	return out, req.Send()
}

// DescribeExclusiveClustersBillingCommonWithContext is the same as DescribeExclusiveClustersBillingCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeExclusiveClustersBillingCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CLB) DescribeExclusiveClustersBillingCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeExclusiveClustersBillingCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeExclusiveClustersBilling = "DescribeExclusiveClustersBilling"

// DescribeExclusiveClustersBillingRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeExclusiveClustersBilling operation. The "output" return
// value will be populated with the DescribeExclusiveClustersBillingCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeExclusiveClustersBillingCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeExclusiveClustersBillingCommon Send returns without error.
//
// See DescribeExclusiveClustersBilling for more information on using the DescribeExclusiveClustersBilling
// API call, and error handling.
//
//    // Example sending a request using the DescribeExclusiveClustersBillingRequest method.
//    req, resp := client.DescribeExclusiveClustersBillingRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CLB) DescribeExclusiveClustersBillingRequest(input *DescribeExclusiveClustersBillingInput) (req *request.Request, output *DescribeExclusiveClustersBillingOutput) {
	op := &request.Operation{
		Name:       opDescribeExclusiveClustersBilling,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeExclusiveClustersBillingInput{}
	}

	output = &DescribeExclusiveClustersBillingOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DescribeExclusiveClustersBilling API operation for CLB.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CLB's
// API operation DescribeExclusiveClustersBilling for usage and error information.
func (c *CLB) DescribeExclusiveClustersBilling(input *DescribeExclusiveClustersBillingInput) (*DescribeExclusiveClustersBillingOutput, error) {
	req, out := c.DescribeExclusiveClustersBillingRequest(input)
	return out, req.Send()
}

// DescribeExclusiveClustersBillingWithContext is the same as DescribeExclusiveClustersBilling with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeExclusiveClustersBilling for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CLB) DescribeExclusiveClustersBillingWithContext(ctx volcengine.Context, input *DescribeExclusiveClustersBillingInput, opts ...request.Option) (*DescribeExclusiveClustersBillingOutput, error) {
	req, out := c.DescribeExclusiveClustersBillingRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DescribeExclusiveClustersBillingInput struct {
	_ struct{} `type:"structure"`

	ExclusiveClusterIds []*string `type:"list"`

	PageNumber *int64 `type:"integer"`

	PageSize *int64 `type:"integer"`
}

// String returns the string representation
func (s DescribeExclusiveClustersBillingInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeExclusiveClustersBillingInput) GoString() string {
	return s.String()
}

// SetExclusiveClusterIds sets the ExclusiveClusterIds field's value.
func (s *DescribeExclusiveClustersBillingInput) SetExclusiveClusterIds(v []*string) *DescribeExclusiveClustersBillingInput {
	s.ExclusiveClusterIds = v
	return s
}

// SetPageNumber sets the PageNumber field's value.
func (s *DescribeExclusiveClustersBillingInput) SetPageNumber(v int64) *DescribeExclusiveClustersBillingInput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *DescribeExclusiveClustersBillingInput) SetPageSize(v int64) *DescribeExclusiveClustersBillingInput {
	s.PageSize = &v
	return s
}

type DescribeExclusiveClustersBillingOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	ExclusiveClusterBillingConfigs []*ExclusiveClusterBillingConfigForDescribeExclusiveClustersBillingOutput `type:"list"`

	PageNumber *int64 `type:"integer"`

	PageSize *int64 `type:"integer"`

	RequestId *string `type:"string"`

	TotalCount *int64 `type:"integer"`
}

// String returns the string representation
func (s DescribeExclusiveClustersBillingOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeExclusiveClustersBillingOutput) GoString() string {
	return s.String()
}

// SetExclusiveClusterBillingConfigs sets the ExclusiveClusterBillingConfigs field's value.
func (s *DescribeExclusiveClustersBillingOutput) SetExclusiveClusterBillingConfigs(v []*ExclusiveClusterBillingConfigForDescribeExclusiveClustersBillingOutput) *DescribeExclusiveClustersBillingOutput {
	s.ExclusiveClusterBillingConfigs = v
	return s
}

// SetPageNumber sets the PageNumber field's value.
func (s *DescribeExclusiveClustersBillingOutput) SetPageNumber(v int64) *DescribeExclusiveClustersBillingOutput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *DescribeExclusiveClustersBillingOutput) SetPageSize(v int64) *DescribeExclusiveClustersBillingOutput {
	s.PageSize = &v
	return s
}

// SetRequestId sets the RequestId field's value.
func (s *DescribeExclusiveClustersBillingOutput) SetRequestId(v string) *DescribeExclusiveClustersBillingOutput {
	s.RequestId = &v
	return s
}

// SetTotalCount sets the TotalCount field's value.
func (s *DescribeExclusiveClustersBillingOutput) SetTotalCount(v int64) *DescribeExclusiveClustersBillingOutput {
	s.TotalCount = &v
	return s
}

type ExclusiveClusterBillingConfigForDescribeExclusiveClustersBillingOutput struct {
	_ struct{} `type:"structure"`

	BillingType *int64 `type:"integer"`

	ExclusiveClusterId *string `type:"string"`

	ExpiredTime *string `type:"string"`

	InstanceStatus *int64 `type:"integer"`

	ReclaimTime *string `type:"string"`

	RemainRenewTimes *int64 `type:"integer"`

	RenewPeriodTimes *int64 `type:"integer"`

	RenewType *int64 `type:"integer"`
}

// String returns the string representation
func (s ExclusiveClusterBillingConfigForDescribeExclusiveClustersBillingOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ExclusiveClusterBillingConfigForDescribeExclusiveClustersBillingOutput) GoString() string {
	return s.String()
}

// SetBillingType sets the BillingType field's value.
func (s *ExclusiveClusterBillingConfigForDescribeExclusiveClustersBillingOutput) SetBillingType(v int64) *ExclusiveClusterBillingConfigForDescribeExclusiveClustersBillingOutput {
	s.BillingType = &v
	return s
}

// SetExclusiveClusterId sets the ExclusiveClusterId field's value.
func (s *ExclusiveClusterBillingConfigForDescribeExclusiveClustersBillingOutput) SetExclusiveClusterId(v string) *ExclusiveClusterBillingConfigForDescribeExclusiveClustersBillingOutput {
	s.ExclusiveClusterId = &v
	return s
}

// SetExpiredTime sets the ExpiredTime field's value.
func (s *ExclusiveClusterBillingConfigForDescribeExclusiveClustersBillingOutput) SetExpiredTime(v string) *ExclusiveClusterBillingConfigForDescribeExclusiveClustersBillingOutput {
	s.ExpiredTime = &v
	return s
}

// SetInstanceStatus sets the InstanceStatus field's value.
func (s *ExclusiveClusterBillingConfigForDescribeExclusiveClustersBillingOutput) SetInstanceStatus(v int64) *ExclusiveClusterBillingConfigForDescribeExclusiveClustersBillingOutput {
	s.InstanceStatus = &v
	return s
}

// SetReclaimTime sets the ReclaimTime field's value.
func (s *ExclusiveClusterBillingConfigForDescribeExclusiveClustersBillingOutput) SetReclaimTime(v string) *ExclusiveClusterBillingConfigForDescribeExclusiveClustersBillingOutput {
	s.ReclaimTime = &v
	return s
}

// SetRemainRenewTimes sets the RemainRenewTimes field's value.
func (s *ExclusiveClusterBillingConfigForDescribeExclusiveClustersBillingOutput) SetRemainRenewTimes(v int64) *ExclusiveClusterBillingConfigForDescribeExclusiveClustersBillingOutput {
	s.RemainRenewTimes = &v
	return s
}

// SetRenewPeriodTimes sets the RenewPeriodTimes field's value.
func (s *ExclusiveClusterBillingConfigForDescribeExclusiveClustersBillingOutput) SetRenewPeriodTimes(v int64) *ExclusiveClusterBillingConfigForDescribeExclusiveClustersBillingOutput {
	s.RenewPeriodTimes = &v
	return s
}

// SetRenewType sets the RenewType field's value.
func (s *ExclusiveClusterBillingConfigForDescribeExclusiveClustersBillingOutput) SetRenewType(v int64) *ExclusiveClusterBillingConfigForDescribeExclusiveClustersBillingOutput {
	s.RenewType = &v
	return s
}

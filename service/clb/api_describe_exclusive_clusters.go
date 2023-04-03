// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package clb

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribeExclusiveClustersCommon = "DescribeExclusiveClusters"

// DescribeExclusiveClustersCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeExclusiveClustersCommon operation. The "output" return
// value will be populated with the DescribeExclusiveClustersCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeExclusiveClustersCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeExclusiveClustersCommon Send returns without error.
//
// See DescribeExclusiveClustersCommon for more information on using the DescribeExclusiveClustersCommon
// API call, and error handling.
//
//    // Example sending a request using the DescribeExclusiveClustersCommonRequest method.
//    req, resp := client.DescribeExclusiveClustersCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CLB) DescribeExclusiveClustersCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeExclusiveClustersCommon,
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

// DescribeExclusiveClustersCommon API operation for CLB.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CLB's
// API operation DescribeExclusiveClustersCommon for usage and error information.
func (c *CLB) DescribeExclusiveClustersCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeExclusiveClustersCommonRequest(input)
	return out, req.Send()
}

// DescribeExclusiveClustersCommonWithContext is the same as DescribeExclusiveClustersCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeExclusiveClustersCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CLB) DescribeExclusiveClustersCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeExclusiveClustersCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeExclusiveClusters = "DescribeExclusiveClusters"

// DescribeExclusiveClustersRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeExclusiveClusters operation. The "output" return
// value will be populated with the DescribeExclusiveClustersCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeExclusiveClustersCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeExclusiveClustersCommon Send returns without error.
//
// See DescribeExclusiveClusters for more information on using the DescribeExclusiveClusters
// API call, and error handling.
//
//    // Example sending a request using the DescribeExclusiveClustersRequest method.
//    req, resp := client.DescribeExclusiveClustersRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CLB) DescribeExclusiveClustersRequest(input *DescribeExclusiveClustersInput) (req *request.Request, output *DescribeExclusiveClustersOutput) {
	op := &request.Operation{
		Name:       opDescribeExclusiveClusters,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeExclusiveClustersInput{}
	}

	output = &DescribeExclusiveClustersOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DescribeExclusiveClusters API operation for CLB.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CLB's
// API operation DescribeExclusiveClusters for usage and error information.
func (c *CLB) DescribeExclusiveClusters(input *DescribeExclusiveClustersInput) (*DescribeExclusiveClustersOutput, error) {
	req, out := c.DescribeExclusiveClustersRequest(input)
	return out, req.Send()
}

// DescribeExclusiveClustersWithContext is the same as DescribeExclusiveClusters with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeExclusiveClusters for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CLB) DescribeExclusiveClustersWithContext(ctx volcengine.Context, input *DescribeExclusiveClustersInput, opts ...request.Option) (*DescribeExclusiveClustersOutput, error) {
	req, out := c.DescribeExclusiveClustersRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DescribeExclusiveClustersInput struct {
	_ struct{} `type:"structure"`

	ExclusiveClusterIds []*string `type:"list"`

	ExclusiveClusterName *string `type:"string"`

	PageNumber *int64 `type:"integer"`

	PageSize *int64 `type:"integer"`

	ProjectName *string `type:"string"`
}

// String returns the string representation
func (s DescribeExclusiveClustersInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeExclusiveClustersInput) GoString() string {
	return s.String()
}

// SetExclusiveClusterIds sets the ExclusiveClusterIds field's value.
func (s *DescribeExclusiveClustersInput) SetExclusiveClusterIds(v []*string) *DescribeExclusiveClustersInput {
	s.ExclusiveClusterIds = v
	return s
}

// SetExclusiveClusterName sets the ExclusiveClusterName field's value.
func (s *DescribeExclusiveClustersInput) SetExclusiveClusterName(v string) *DescribeExclusiveClustersInput {
	s.ExclusiveClusterName = &v
	return s
}

// SetPageNumber sets the PageNumber field's value.
func (s *DescribeExclusiveClustersInput) SetPageNumber(v int64) *DescribeExclusiveClustersInput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *DescribeExclusiveClustersInput) SetPageSize(v int64) *DescribeExclusiveClustersInput {
	s.PageSize = &v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *DescribeExclusiveClustersInput) SetProjectName(v string) *DescribeExclusiveClustersInput {
	s.ProjectName = &v
	return s
}

type DescribeExclusiveClustersOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	ExclusiveClusters []*ExclusiveClusterForDescribeExclusiveClustersOutput `type:"list"`

	PageNumber *int64 `type:"integer"`

	PageSize *int64 `type:"integer"`

	RequestId *string `type:"string"`

	TotalCount *int64 `type:"integer"`
}

// String returns the string representation
func (s DescribeExclusiveClustersOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeExclusiveClustersOutput) GoString() string {
	return s.String()
}

// SetExclusiveClusters sets the ExclusiveClusters field's value.
func (s *DescribeExclusiveClustersOutput) SetExclusiveClusters(v []*ExclusiveClusterForDescribeExclusiveClustersOutput) *DescribeExclusiveClustersOutput {
	s.ExclusiveClusters = v
	return s
}

// SetPageNumber sets the PageNumber field's value.
func (s *DescribeExclusiveClustersOutput) SetPageNumber(v int64) *DescribeExclusiveClustersOutput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *DescribeExclusiveClustersOutput) SetPageSize(v int64) *DescribeExclusiveClustersOutput {
	s.PageSize = &v
	return s
}

// SetRequestId sets the RequestId field's value.
func (s *DescribeExclusiveClustersOutput) SetRequestId(v string) *DescribeExclusiveClustersOutput {
	s.RequestId = &v
	return s
}

// SetTotalCount sets the TotalCount field's value.
func (s *DescribeExclusiveClustersOutput) SetTotalCount(v int64) *DescribeExclusiveClustersOutput {
	s.TotalCount = &v
	return s
}

type ExclusiveClusterForDescribeExclusiveClustersOutput struct {
	_ struct{} `type:"structure"`

	AvailableZones []*string `type:"list"`

	BillingType *int64 `type:"integer"`

	BusinessStatus *string `type:"string"`

	CreateTime *string `type:"string"`

	DispatchUnitType *string `type:"string"`

	ExclusiveClusterId *string `type:"string"`

	ExclusiveClusterName *string `type:"string"`

	ExpiredTime *string `type:"string"`

	LoadBalancerCount *int64 `type:"integer"`

	LockReason *string `type:"string"`

	MaxPerformanceUnitCount *int64 `type:"integer"`

	PerformanceUnit *string `type:"string"`

	PerformanceUnitCount *int64 `type:"integer"`

	ProjectName *string `type:"string"`

	ReclaimedTime *string `type:"string"`

	Status *string `type:"string"`

	UpdateTime *string `type:"string"`
}

// String returns the string representation
func (s ExclusiveClusterForDescribeExclusiveClustersOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ExclusiveClusterForDescribeExclusiveClustersOutput) GoString() string {
	return s.String()
}

// SetAvailableZones sets the AvailableZones field's value.
func (s *ExclusiveClusterForDescribeExclusiveClustersOutput) SetAvailableZones(v []*string) *ExclusiveClusterForDescribeExclusiveClustersOutput {
	s.AvailableZones = v
	return s
}

// SetBillingType sets the BillingType field's value.
func (s *ExclusiveClusterForDescribeExclusiveClustersOutput) SetBillingType(v int64) *ExclusiveClusterForDescribeExclusiveClustersOutput {
	s.BillingType = &v
	return s
}

// SetBusinessStatus sets the BusinessStatus field's value.
func (s *ExclusiveClusterForDescribeExclusiveClustersOutput) SetBusinessStatus(v string) *ExclusiveClusterForDescribeExclusiveClustersOutput {
	s.BusinessStatus = &v
	return s
}

// SetCreateTime sets the CreateTime field's value.
func (s *ExclusiveClusterForDescribeExclusiveClustersOutput) SetCreateTime(v string) *ExclusiveClusterForDescribeExclusiveClustersOutput {
	s.CreateTime = &v
	return s
}

// SetDispatchUnitType sets the DispatchUnitType field's value.
func (s *ExclusiveClusterForDescribeExclusiveClustersOutput) SetDispatchUnitType(v string) *ExclusiveClusterForDescribeExclusiveClustersOutput {
	s.DispatchUnitType = &v
	return s
}

// SetExclusiveClusterId sets the ExclusiveClusterId field's value.
func (s *ExclusiveClusterForDescribeExclusiveClustersOutput) SetExclusiveClusterId(v string) *ExclusiveClusterForDescribeExclusiveClustersOutput {
	s.ExclusiveClusterId = &v
	return s
}

// SetExclusiveClusterName sets the ExclusiveClusterName field's value.
func (s *ExclusiveClusterForDescribeExclusiveClustersOutput) SetExclusiveClusterName(v string) *ExclusiveClusterForDescribeExclusiveClustersOutput {
	s.ExclusiveClusterName = &v
	return s
}

// SetExpiredTime sets the ExpiredTime field's value.
func (s *ExclusiveClusterForDescribeExclusiveClustersOutput) SetExpiredTime(v string) *ExclusiveClusterForDescribeExclusiveClustersOutput {
	s.ExpiredTime = &v
	return s
}

// SetLoadBalancerCount sets the LoadBalancerCount field's value.
func (s *ExclusiveClusterForDescribeExclusiveClustersOutput) SetLoadBalancerCount(v int64) *ExclusiveClusterForDescribeExclusiveClustersOutput {
	s.LoadBalancerCount = &v
	return s
}

// SetLockReason sets the LockReason field's value.
func (s *ExclusiveClusterForDescribeExclusiveClustersOutput) SetLockReason(v string) *ExclusiveClusterForDescribeExclusiveClustersOutput {
	s.LockReason = &v
	return s
}

// SetMaxPerformanceUnitCount sets the MaxPerformanceUnitCount field's value.
func (s *ExclusiveClusterForDescribeExclusiveClustersOutput) SetMaxPerformanceUnitCount(v int64) *ExclusiveClusterForDescribeExclusiveClustersOutput {
	s.MaxPerformanceUnitCount = &v
	return s
}

// SetPerformanceUnit sets the PerformanceUnit field's value.
func (s *ExclusiveClusterForDescribeExclusiveClustersOutput) SetPerformanceUnit(v string) *ExclusiveClusterForDescribeExclusiveClustersOutput {
	s.PerformanceUnit = &v
	return s
}

// SetPerformanceUnitCount sets the PerformanceUnitCount field's value.
func (s *ExclusiveClusterForDescribeExclusiveClustersOutput) SetPerformanceUnitCount(v int64) *ExclusiveClusterForDescribeExclusiveClustersOutput {
	s.PerformanceUnitCount = &v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *ExclusiveClusterForDescribeExclusiveClustersOutput) SetProjectName(v string) *ExclusiveClusterForDescribeExclusiveClustersOutput {
	s.ProjectName = &v
	return s
}

// SetReclaimedTime sets the ReclaimedTime field's value.
func (s *ExclusiveClusterForDescribeExclusiveClustersOutput) SetReclaimedTime(v string) *ExclusiveClusterForDescribeExclusiveClustersOutput {
	s.ReclaimedTime = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *ExclusiveClusterForDescribeExclusiveClustersOutput) SetStatus(v string) *ExclusiveClusterForDescribeExclusiveClustersOutput {
	s.Status = &v
	return s
}

// SetUpdateTime sets the UpdateTime field's value.
func (s *ExclusiveClusterForDescribeExclusiveClustersOutput) SetUpdateTime(v string) *ExclusiveClusterForDescribeExclusiveClustersOutput {
	s.UpdateTime = &v
	return s
}
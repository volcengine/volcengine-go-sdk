// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package ecs

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribeDeploymentSetsCommon = "DescribeDeploymentSets"

// DescribeDeploymentSetsCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeDeploymentSetsCommon operation. The "output" return
// value will be populated with the DescribeDeploymentSetsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeDeploymentSetsCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeDeploymentSetsCommon Send returns without error.
//
// See DescribeDeploymentSetsCommon for more information on using the DescribeDeploymentSetsCommon
// API call, and error handling.
//
//    // Example sending a request using the DescribeDeploymentSetsCommonRequest method.
//    req, resp := client.DescribeDeploymentSetsCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ECS) DescribeDeploymentSetsCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeDeploymentSetsCommon,
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

// DescribeDeploymentSetsCommon API operation for ECS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ECS's
// API operation DescribeDeploymentSetsCommon for usage and error information.
func (c *ECS) DescribeDeploymentSetsCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeDeploymentSetsCommonRequest(input)
	return out, req.Send()
}

// DescribeDeploymentSetsCommonWithContext is the same as DescribeDeploymentSetsCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeDeploymentSetsCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ECS) DescribeDeploymentSetsCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeDeploymentSetsCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeDeploymentSets = "DescribeDeploymentSets"

// DescribeDeploymentSetsRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeDeploymentSets operation. The "output" return
// value will be populated with the DescribeDeploymentSetsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeDeploymentSetsCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeDeploymentSetsCommon Send returns without error.
//
// See DescribeDeploymentSets for more information on using the DescribeDeploymentSets
// API call, and error handling.
//
//    // Example sending a request using the DescribeDeploymentSetsRequest method.
//    req, resp := client.DescribeDeploymentSetsRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ECS) DescribeDeploymentSetsRequest(input *DescribeDeploymentSetsInput) (req *request.Request, output *DescribeDeploymentSetsOutput) {
	op := &request.Operation{
		Name:       opDescribeDeploymentSets,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeDeploymentSetsInput{}
	}

	output = &DescribeDeploymentSetsOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DescribeDeploymentSets API operation for ECS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ECS's
// API operation DescribeDeploymentSets for usage and error information.
func (c *ECS) DescribeDeploymentSets(input *DescribeDeploymentSetsInput) (*DescribeDeploymentSetsOutput, error) {
	req, out := c.DescribeDeploymentSetsRequest(input)
	return out, req.Send()
}

// DescribeDeploymentSetsWithContext is the same as DescribeDeploymentSets with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeDeploymentSets for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ECS) DescribeDeploymentSetsWithContext(ctx volcengine.Context, input *DescribeDeploymentSetsInput, opts ...request.Option) (*DescribeDeploymentSetsOutput, error) {
	req, out := c.DescribeDeploymentSetsRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type CapacityForDescribeDeploymentSetsOutput struct {
	_ struct{} `type:"structure"`

	AvailableCount *int32 `type:"int32"`

	UsedCount *int32 `type:"int32"`

	ZoneId *string `type:"string"`
}

// String returns the string representation
func (s CapacityForDescribeDeploymentSetsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CapacityForDescribeDeploymentSetsOutput) GoString() string {
	return s.String()
}

// SetAvailableCount sets the AvailableCount field's value.
func (s *CapacityForDescribeDeploymentSetsOutput) SetAvailableCount(v int32) *CapacityForDescribeDeploymentSetsOutput {
	s.AvailableCount = &v
	return s
}

// SetUsedCount sets the UsedCount field's value.
func (s *CapacityForDescribeDeploymentSetsOutput) SetUsedCount(v int32) *CapacityForDescribeDeploymentSetsOutput {
	s.UsedCount = &v
	return s
}

// SetZoneId sets the ZoneId field's value.
func (s *CapacityForDescribeDeploymentSetsOutput) SetZoneId(v string) *CapacityForDescribeDeploymentSetsOutput {
	s.ZoneId = &v
	return s
}

type DeploymentSetForDescribeDeploymentSetsOutput struct {
	_ struct{} `type:"structure"`

	Capacities []*CapacityForDescribeDeploymentSetsOutput `type:"list"`

	CreatedAt *string `type:"string"`

	DeploymentSetDescription *string `type:"string"`

	DeploymentSetId *string `type:"string"`

	DeploymentSetName *string `type:"string"`

	Granularity *string `type:"string"`

	GroupCount *int32 `type:"int32"`

	InstanceAmount *int32 `type:"int32"`

	InstanceIds []*string `type:"list"`

	Quota *int32 `type:"int32"`

	QuotaUsed []*QuotaUsedForDescribeDeploymentSetsOutput `type:"list"`

	Strategy *string `type:"string"`
}

// String returns the string representation
func (s DeploymentSetForDescribeDeploymentSetsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DeploymentSetForDescribeDeploymentSetsOutput) GoString() string {
	return s.String()
}

// SetCapacities sets the Capacities field's value.
func (s *DeploymentSetForDescribeDeploymentSetsOutput) SetCapacities(v []*CapacityForDescribeDeploymentSetsOutput) *DeploymentSetForDescribeDeploymentSetsOutput {
	s.Capacities = v
	return s
}

// SetCreatedAt sets the CreatedAt field's value.
func (s *DeploymentSetForDescribeDeploymentSetsOutput) SetCreatedAt(v string) *DeploymentSetForDescribeDeploymentSetsOutput {
	s.CreatedAt = &v
	return s
}

// SetDeploymentSetDescription sets the DeploymentSetDescription field's value.
func (s *DeploymentSetForDescribeDeploymentSetsOutput) SetDeploymentSetDescription(v string) *DeploymentSetForDescribeDeploymentSetsOutput {
	s.DeploymentSetDescription = &v
	return s
}

// SetDeploymentSetId sets the DeploymentSetId field's value.
func (s *DeploymentSetForDescribeDeploymentSetsOutput) SetDeploymentSetId(v string) *DeploymentSetForDescribeDeploymentSetsOutput {
	s.DeploymentSetId = &v
	return s
}

// SetDeploymentSetName sets the DeploymentSetName field's value.
func (s *DeploymentSetForDescribeDeploymentSetsOutput) SetDeploymentSetName(v string) *DeploymentSetForDescribeDeploymentSetsOutput {
	s.DeploymentSetName = &v
	return s
}

// SetGranularity sets the Granularity field's value.
func (s *DeploymentSetForDescribeDeploymentSetsOutput) SetGranularity(v string) *DeploymentSetForDescribeDeploymentSetsOutput {
	s.Granularity = &v
	return s
}

// SetGroupCount sets the GroupCount field's value.
func (s *DeploymentSetForDescribeDeploymentSetsOutput) SetGroupCount(v int32) *DeploymentSetForDescribeDeploymentSetsOutput {
	s.GroupCount = &v
	return s
}

// SetInstanceAmount sets the InstanceAmount field's value.
func (s *DeploymentSetForDescribeDeploymentSetsOutput) SetInstanceAmount(v int32) *DeploymentSetForDescribeDeploymentSetsOutput {
	s.InstanceAmount = &v
	return s
}

// SetInstanceIds sets the InstanceIds field's value.
func (s *DeploymentSetForDescribeDeploymentSetsOutput) SetInstanceIds(v []*string) *DeploymentSetForDescribeDeploymentSetsOutput {
	s.InstanceIds = v
	return s
}

// SetQuota sets the Quota field's value.
func (s *DeploymentSetForDescribeDeploymentSetsOutput) SetQuota(v int32) *DeploymentSetForDescribeDeploymentSetsOutput {
	s.Quota = &v
	return s
}

// SetQuotaUsed sets the QuotaUsed field's value.
func (s *DeploymentSetForDescribeDeploymentSetsOutput) SetQuotaUsed(v []*QuotaUsedForDescribeDeploymentSetsOutput) *DeploymentSetForDescribeDeploymentSetsOutput {
	s.QuotaUsed = v
	return s
}

// SetStrategy sets the Strategy field's value.
func (s *DeploymentSetForDescribeDeploymentSetsOutput) SetStrategy(v string) *DeploymentSetForDescribeDeploymentSetsOutput {
	s.Strategy = &v
	return s
}

type DescribeDeploymentSetsInput struct {
	_ struct{} `type:"structure"`

	AccountId *string `type:"string"`

	DeploymentSetIds []*string `type:"list"`

	DeploymentSetName *string `type:"string"`

	Granularity *string `type:"string"`

	MaxResults *int32 `type:"int32"`

	NextToken *string `type:"string"`

	PageNumber *int32 `type:"int32"`

	PageSize *int32 `type:"int32"`

	Strategy *string `type:"string"`
}

// String returns the string representation
func (s DescribeDeploymentSetsInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeDeploymentSetsInput) GoString() string {
	return s.String()
}

// SetAccountId sets the AccountId field's value.
func (s *DescribeDeploymentSetsInput) SetAccountId(v string) *DescribeDeploymentSetsInput {
	s.AccountId = &v
	return s
}

// SetDeploymentSetIds sets the DeploymentSetIds field's value.
func (s *DescribeDeploymentSetsInput) SetDeploymentSetIds(v []*string) *DescribeDeploymentSetsInput {
	s.DeploymentSetIds = v
	return s
}

// SetDeploymentSetName sets the DeploymentSetName field's value.
func (s *DescribeDeploymentSetsInput) SetDeploymentSetName(v string) *DescribeDeploymentSetsInput {
	s.DeploymentSetName = &v
	return s
}

// SetGranularity sets the Granularity field's value.
func (s *DescribeDeploymentSetsInput) SetGranularity(v string) *DescribeDeploymentSetsInput {
	s.Granularity = &v
	return s
}

// SetMaxResults sets the MaxResults field's value.
func (s *DescribeDeploymentSetsInput) SetMaxResults(v int32) *DescribeDeploymentSetsInput {
	s.MaxResults = &v
	return s
}

// SetNextToken sets the NextToken field's value.
func (s *DescribeDeploymentSetsInput) SetNextToken(v string) *DescribeDeploymentSetsInput {
	s.NextToken = &v
	return s
}

// SetPageNumber sets the PageNumber field's value.
func (s *DescribeDeploymentSetsInput) SetPageNumber(v int32) *DescribeDeploymentSetsInput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *DescribeDeploymentSetsInput) SetPageSize(v int32) *DescribeDeploymentSetsInput {
	s.PageSize = &v
	return s
}

// SetStrategy sets the Strategy field's value.
func (s *DescribeDeploymentSetsInput) SetStrategy(v string) *DescribeDeploymentSetsInput {
	s.Strategy = &v
	return s
}

type DescribeDeploymentSetsOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	DeploymentSets []*DeploymentSetForDescribeDeploymentSetsOutput `type:"list"`

	NextToken *string `type:"string"`

	PageNumber *int32 `type:"int32"`

	PageSize *int32 `type:"int32"`

	TotalCount *int32 `type:"int32"`
}

// String returns the string representation
func (s DescribeDeploymentSetsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeDeploymentSetsOutput) GoString() string {
	return s.String()
}

// SetDeploymentSets sets the DeploymentSets field's value.
func (s *DescribeDeploymentSetsOutput) SetDeploymentSets(v []*DeploymentSetForDescribeDeploymentSetsOutput) *DescribeDeploymentSetsOutput {
	s.DeploymentSets = v
	return s
}

// SetNextToken sets the NextToken field's value.
func (s *DescribeDeploymentSetsOutput) SetNextToken(v string) *DescribeDeploymentSetsOutput {
	s.NextToken = &v
	return s
}

// SetPageNumber sets the PageNumber field's value.
func (s *DescribeDeploymentSetsOutput) SetPageNumber(v int32) *DescribeDeploymentSetsOutput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *DescribeDeploymentSetsOutput) SetPageSize(v int32) *DescribeDeploymentSetsOutput {
	s.PageSize = &v
	return s
}

// SetTotalCount sets the TotalCount field's value.
func (s *DescribeDeploymentSetsOutput) SetTotalCount(v int32) *DescribeDeploymentSetsOutput {
	s.TotalCount = &v
	return s
}

type QuotaUsedForDescribeDeploymentSetsOutput struct {
	_ struct{} `type:"structure"`

	Count *int32 `type:"int32"`

	ZoneId *string `type:"string"`
}

// String returns the string representation
func (s QuotaUsedForDescribeDeploymentSetsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s QuotaUsedForDescribeDeploymentSetsOutput) GoString() string {
	return s.String()
}

// SetCount sets the Count field's value.
func (s *QuotaUsedForDescribeDeploymentSetsOutput) SetCount(v int32) *QuotaUsedForDescribeDeploymentSetsOutput {
	s.Count = &v
	return s
}

// SetZoneId sets the ZoneId field's value.
func (s *QuotaUsedForDescribeDeploymentSetsOutput) SetZoneId(v string) *QuotaUsedForDescribeDeploymentSetsOutput {
	s.ZoneId = &v
	return s
}

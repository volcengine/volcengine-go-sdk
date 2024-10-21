// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package ecs

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribeReservedInstancesCommon = "DescribeReservedInstances"

// DescribeReservedInstancesCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeReservedInstancesCommon operation. The "output" return
// value will be populated with the DescribeReservedInstancesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeReservedInstancesCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeReservedInstancesCommon Send returns without error.
//
// See DescribeReservedInstancesCommon for more information on using the DescribeReservedInstancesCommon
// API call, and error handling.
//
//    // Example sending a request using the DescribeReservedInstancesCommonRequest method.
//    req, resp := client.DescribeReservedInstancesCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ECS) DescribeReservedInstancesCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeReservedInstancesCommon,
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

// DescribeReservedInstancesCommon API operation for ECS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ECS's
// API operation DescribeReservedInstancesCommon for usage and error information.
func (c *ECS) DescribeReservedInstancesCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeReservedInstancesCommonRequest(input)
	return out, req.Send()
}

// DescribeReservedInstancesCommonWithContext is the same as DescribeReservedInstancesCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeReservedInstancesCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ECS) DescribeReservedInstancesCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeReservedInstancesCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeReservedInstances = "DescribeReservedInstances"

// DescribeReservedInstancesRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeReservedInstances operation. The "output" return
// value will be populated with the DescribeReservedInstancesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeReservedInstancesCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeReservedInstancesCommon Send returns without error.
//
// See DescribeReservedInstances for more information on using the DescribeReservedInstances
// API call, and error handling.
//
//    // Example sending a request using the DescribeReservedInstancesRequest method.
//    req, resp := client.DescribeReservedInstancesRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ECS) DescribeReservedInstancesRequest(input *DescribeReservedInstancesInput) (req *request.Request, output *DescribeReservedInstancesOutput) {
	op := &request.Operation{
		Name:       opDescribeReservedInstances,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeReservedInstancesInput{}
	}

	output = &DescribeReservedInstancesOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DescribeReservedInstances API operation for ECS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ECS's
// API operation DescribeReservedInstances for usage and error information.
func (c *ECS) DescribeReservedInstances(input *DescribeReservedInstancesInput) (*DescribeReservedInstancesOutput, error) {
	req, out := c.DescribeReservedInstancesRequest(input)
	return out, req.Send()
}

// DescribeReservedInstancesWithContext is the same as DescribeReservedInstances with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeReservedInstances for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ECS) DescribeReservedInstancesWithContext(ctx volcengine.Context, input *DescribeReservedInstancesInput, opts ...request.Option) (*DescribeReservedInstancesOutput, error) {
	req, out := c.DescribeReservedInstancesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DescribeReservedInstancesInput struct {
	_ struct{} `type:"structure"`

	HpcClusterId *string `type:"string"`

	InstanceTypeFamilies []*string `type:"list"`

	InstanceTypeIds []*string `type:"list"`

	MaxResults *int32 `type:"int32"`

	NextToken *string `type:"string"`

	ProjectName *string `type:"string"`

	ReservedInstanceIds []*string `type:"list"`

	ReservedInstanceName *string `type:"string"`

	Scope *string `type:"string"`

	Status *string `type:"string"`

	SupportModify *string `type:"string"`

	TagFilters []*TagFilterForDescribeReservedInstancesInput `type:"list"`

	ZoneId *string `type:"string"`
}

// String returns the string representation
func (s DescribeReservedInstancesInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeReservedInstancesInput) GoString() string {
	return s.String()
}

// SetHpcClusterId sets the HpcClusterId field's value.
func (s *DescribeReservedInstancesInput) SetHpcClusterId(v string) *DescribeReservedInstancesInput {
	s.HpcClusterId = &v
	return s
}

// SetInstanceTypeFamilies sets the InstanceTypeFamilies field's value.
func (s *DescribeReservedInstancesInput) SetInstanceTypeFamilies(v []*string) *DescribeReservedInstancesInput {
	s.InstanceTypeFamilies = v
	return s
}

// SetInstanceTypeIds sets the InstanceTypeIds field's value.
func (s *DescribeReservedInstancesInput) SetInstanceTypeIds(v []*string) *DescribeReservedInstancesInput {
	s.InstanceTypeIds = v
	return s
}

// SetMaxResults sets the MaxResults field's value.
func (s *DescribeReservedInstancesInput) SetMaxResults(v int32) *DescribeReservedInstancesInput {
	s.MaxResults = &v
	return s
}

// SetNextToken sets the NextToken field's value.
func (s *DescribeReservedInstancesInput) SetNextToken(v string) *DescribeReservedInstancesInput {
	s.NextToken = &v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *DescribeReservedInstancesInput) SetProjectName(v string) *DescribeReservedInstancesInput {
	s.ProjectName = &v
	return s
}

// SetReservedInstanceIds sets the ReservedInstanceIds field's value.
func (s *DescribeReservedInstancesInput) SetReservedInstanceIds(v []*string) *DescribeReservedInstancesInput {
	s.ReservedInstanceIds = v
	return s
}

// SetReservedInstanceName sets the ReservedInstanceName field's value.
func (s *DescribeReservedInstancesInput) SetReservedInstanceName(v string) *DescribeReservedInstancesInput {
	s.ReservedInstanceName = &v
	return s
}

// SetScope sets the Scope field's value.
func (s *DescribeReservedInstancesInput) SetScope(v string) *DescribeReservedInstancesInput {
	s.Scope = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *DescribeReservedInstancesInput) SetStatus(v string) *DescribeReservedInstancesInput {
	s.Status = &v
	return s
}

// SetSupportModify sets the SupportModify field's value.
func (s *DescribeReservedInstancesInput) SetSupportModify(v string) *DescribeReservedInstancesInput {
	s.SupportModify = &v
	return s
}

// SetTagFilters sets the TagFilters field's value.
func (s *DescribeReservedInstancesInput) SetTagFilters(v []*TagFilterForDescribeReservedInstancesInput) *DescribeReservedInstancesInput {
	s.TagFilters = v
	return s
}

// SetZoneId sets the ZoneId field's value.
func (s *DescribeReservedInstancesInput) SetZoneId(v string) *DescribeReservedInstancesInput {
	s.ZoneId = &v
	return s
}

type DescribeReservedInstancesOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	NextToken *string `type:"string"`

	ReservedInstances []*ReservedInstanceForDescribeReservedInstancesOutput `type:"list"`
}

// String returns the string representation
func (s DescribeReservedInstancesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeReservedInstancesOutput) GoString() string {
	return s.String()
}

// SetNextToken sets the NextToken field's value.
func (s *DescribeReservedInstancesOutput) SetNextToken(v string) *DescribeReservedInstancesOutput {
	s.NextToken = &v
	return s
}

// SetReservedInstances sets the ReservedInstances field's value.
func (s *DescribeReservedInstancesOutput) SetReservedInstances(v []*ReservedInstanceForDescribeReservedInstancesOutput) *DescribeReservedInstancesOutput {
	s.ReservedInstances = v
	return s
}

type ReservedInstanceForDescribeReservedInstancesOutput struct {
	_ struct{} `type:"structure"`

	CreatedAt *string `type:"string"`

	Description *string `type:"string"`

	ExpiredAt *string `type:"string"`

	HpcClusterId *string `type:"string"`

	InstanceCount *int32 `type:"int32"`

	InstanceTypeId *string `type:"string"`

	OfferingType *string `type:"string"`

	ProjectName *string `type:"string"`

	RegionId *string `type:"string"`

	ReservedInstanceId *string `type:"string"`

	ReservedInstanceName *string `type:"string"`

	Scope *string `type:"string"`

	StartAt *string `type:"string"`

	Status *string `type:"string"`

	SupportModify *string `type:"string"`

	Tags []*TagForDescribeReservedInstancesOutput `type:"list"`

	ZoneId *string `type:"string"`
}

// String returns the string representation
func (s ReservedInstanceForDescribeReservedInstancesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ReservedInstanceForDescribeReservedInstancesOutput) GoString() string {
	return s.String()
}

// SetCreatedAt sets the CreatedAt field's value.
func (s *ReservedInstanceForDescribeReservedInstancesOutput) SetCreatedAt(v string) *ReservedInstanceForDescribeReservedInstancesOutput {
	s.CreatedAt = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *ReservedInstanceForDescribeReservedInstancesOutput) SetDescription(v string) *ReservedInstanceForDescribeReservedInstancesOutput {
	s.Description = &v
	return s
}

// SetExpiredAt sets the ExpiredAt field's value.
func (s *ReservedInstanceForDescribeReservedInstancesOutput) SetExpiredAt(v string) *ReservedInstanceForDescribeReservedInstancesOutput {
	s.ExpiredAt = &v
	return s
}

// SetHpcClusterId sets the HpcClusterId field's value.
func (s *ReservedInstanceForDescribeReservedInstancesOutput) SetHpcClusterId(v string) *ReservedInstanceForDescribeReservedInstancesOutput {
	s.HpcClusterId = &v
	return s
}

// SetInstanceCount sets the InstanceCount field's value.
func (s *ReservedInstanceForDescribeReservedInstancesOutput) SetInstanceCount(v int32) *ReservedInstanceForDescribeReservedInstancesOutput {
	s.InstanceCount = &v
	return s
}

// SetInstanceTypeId sets the InstanceTypeId field's value.
func (s *ReservedInstanceForDescribeReservedInstancesOutput) SetInstanceTypeId(v string) *ReservedInstanceForDescribeReservedInstancesOutput {
	s.InstanceTypeId = &v
	return s
}

// SetOfferingType sets the OfferingType field's value.
func (s *ReservedInstanceForDescribeReservedInstancesOutput) SetOfferingType(v string) *ReservedInstanceForDescribeReservedInstancesOutput {
	s.OfferingType = &v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *ReservedInstanceForDescribeReservedInstancesOutput) SetProjectName(v string) *ReservedInstanceForDescribeReservedInstancesOutput {
	s.ProjectName = &v
	return s
}

// SetRegionId sets the RegionId field's value.
func (s *ReservedInstanceForDescribeReservedInstancesOutput) SetRegionId(v string) *ReservedInstanceForDescribeReservedInstancesOutput {
	s.RegionId = &v
	return s
}

// SetReservedInstanceId sets the ReservedInstanceId field's value.
func (s *ReservedInstanceForDescribeReservedInstancesOutput) SetReservedInstanceId(v string) *ReservedInstanceForDescribeReservedInstancesOutput {
	s.ReservedInstanceId = &v
	return s
}

// SetReservedInstanceName sets the ReservedInstanceName field's value.
func (s *ReservedInstanceForDescribeReservedInstancesOutput) SetReservedInstanceName(v string) *ReservedInstanceForDescribeReservedInstancesOutput {
	s.ReservedInstanceName = &v
	return s
}

// SetScope sets the Scope field's value.
func (s *ReservedInstanceForDescribeReservedInstancesOutput) SetScope(v string) *ReservedInstanceForDescribeReservedInstancesOutput {
	s.Scope = &v
	return s
}

// SetStartAt sets the StartAt field's value.
func (s *ReservedInstanceForDescribeReservedInstancesOutput) SetStartAt(v string) *ReservedInstanceForDescribeReservedInstancesOutput {
	s.StartAt = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *ReservedInstanceForDescribeReservedInstancesOutput) SetStatus(v string) *ReservedInstanceForDescribeReservedInstancesOutput {
	s.Status = &v
	return s
}

// SetSupportModify sets the SupportModify field's value.
func (s *ReservedInstanceForDescribeReservedInstancesOutput) SetSupportModify(v string) *ReservedInstanceForDescribeReservedInstancesOutput {
	s.SupportModify = &v
	return s
}

// SetTags sets the Tags field's value.
func (s *ReservedInstanceForDescribeReservedInstancesOutput) SetTags(v []*TagForDescribeReservedInstancesOutput) *ReservedInstanceForDescribeReservedInstancesOutput {
	s.Tags = v
	return s
}

// SetZoneId sets the ZoneId field's value.
func (s *ReservedInstanceForDescribeReservedInstancesOutput) SetZoneId(v string) *ReservedInstanceForDescribeReservedInstancesOutput {
	s.ZoneId = &v
	return s
}

type TagFilterForDescribeReservedInstancesInput struct {
	_ struct{} `type:"structure"`

	Key *string `type:"string"`

	Values []*string `type:"list"`
}

// String returns the string representation
func (s TagFilterForDescribeReservedInstancesInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s TagFilterForDescribeReservedInstancesInput) GoString() string {
	return s.String()
}

// SetKey sets the Key field's value.
func (s *TagFilterForDescribeReservedInstancesInput) SetKey(v string) *TagFilterForDescribeReservedInstancesInput {
	s.Key = &v
	return s
}

// SetValues sets the Values field's value.
func (s *TagFilterForDescribeReservedInstancesInput) SetValues(v []*string) *TagFilterForDescribeReservedInstancesInput {
	s.Values = v
	return s
}

type TagForDescribeReservedInstancesOutput struct {
	_ struct{} `type:"structure"`

	Key *string `type:"string"`

	Value *string `type:"string"`
}

// String returns the string representation
func (s TagForDescribeReservedInstancesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s TagForDescribeReservedInstancesOutput) GoString() string {
	return s.String()
}

// SetKey sets the Key field's value.
func (s *TagForDescribeReservedInstancesOutput) SetKey(v string) *TagForDescribeReservedInstancesOutput {
	s.Key = &v
	return s
}

// SetValue sets the Value field's value.
func (s *TagForDescribeReservedInstancesOutput) SetValue(v string) *TagForDescribeReservedInstancesOutput {
	s.Value = &v
	return s
}

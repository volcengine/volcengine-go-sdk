// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package autoscaling

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribeScalingGroupsCommon = "DescribeScalingGroups"

// DescribeScalingGroupsCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeScalingGroupsCommon operation. The "output" return
// value will be populated with the DescribeScalingGroupsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeScalingGroupsCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeScalingGroupsCommon Send returns without error.
//
// See DescribeScalingGroupsCommon for more information on using the DescribeScalingGroupsCommon
// API call, and error handling.
//
//	// Example sending a request using the DescribeScalingGroupsCommonRequest method.
//	req, resp := client.DescribeScalingGroupsCommonRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *AUTOSCALING) DescribeScalingGroupsCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeScalingGroupsCommon,
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

// DescribeScalingGroupsCommon API operation for AUTO_SCALING.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for AUTO_SCALING's
// API operation DescribeScalingGroupsCommon for usage and error information.
func (c *AUTOSCALING) DescribeScalingGroupsCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeScalingGroupsCommonRequest(input)
	return out, req.Send()
}

// DescribeScalingGroupsCommonWithContext is the same as DescribeScalingGroupsCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeScalingGroupsCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *AUTOSCALING) DescribeScalingGroupsCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeScalingGroupsCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeScalingGroups = "DescribeScalingGroups"

// DescribeScalingGroupsRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeScalingGroups operation. The "output" return
// value will be populated with the DescribeScalingGroupsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeScalingGroupsCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeScalingGroupsCommon Send returns without error.
//
// See DescribeScalingGroups for more information on using the DescribeScalingGroups
// API call, and error handling.
//
//	// Example sending a request using the DescribeScalingGroupsRequest method.
//	req, resp := client.DescribeScalingGroupsRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *AUTOSCALING) DescribeScalingGroupsRequest(input *DescribeScalingGroupsInput) (req *request.Request, output *DescribeScalingGroupsOutput) {
	op := &request.Operation{
		Name:       opDescribeScalingGroups,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeScalingGroupsInput{}
	}

	output = &DescribeScalingGroupsOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DescribeScalingGroups API operation for AUTO_SCALING.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for AUTO_SCALING's
// API operation DescribeScalingGroups for usage and error information.
func (c *AUTOSCALING) DescribeScalingGroups(input *DescribeScalingGroupsInput) (*DescribeScalingGroupsOutput, error) {
	req, out := c.DescribeScalingGroupsRequest(input)
	return out, req.Send()
}

// DescribeScalingGroupsWithContext is the same as DescribeScalingGroups with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeScalingGroups for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *AUTOSCALING) DescribeScalingGroupsWithContext(ctx volcengine.Context, input *DescribeScalingGroupsInput, opts ...request.Option) (*DescribeScalingGroupsOutput, error) {
	req, out := c.DescribeScalingGroupsRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DescribeScalingGroupsInput struct {
	_ struct{} `type:"structure"`

	PageNumber *int32 `type:"int32"`

	PageSize *int32 `type:"int32"`

	ProjectName *string `type:"string"`

	ScalingGroupIds []*string `type:"list"`

	ScalingGroupNames []*string `type:"list"`

	TagFilters []*TagFilterForDescribeScalingGroupsInput `type:"list"`
}

// String returns the string representation
func (s DescribeScalingGroupsInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeScalingGroupsInput) GoString() string {
	return s.String()
}

// SetPageNumber sets the PageNumber field's value.
func (s *DescribeScalingGroupsInput) SetPageNumber(v int32) *DescribeScalingGroupsInput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *DescribeScalingGroupsInput) SetPageSize(v int32) *DescribeScalingGroupsInput {
	s.PageSize = &v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *DescribeScalingGroupsInput) SetProjectName(v string) *DescribeScalingGroupsInput {
	s.ProjectName = &v
	return s
}

// SetScalingGroupIds sets the ScalingGroupIds field's value.
func (s *DescribeScalingGroupsInput) SetScalingGroupIds(v []*string) *DescribeScalingGroupsInput {
	s.ScalingGroupIds = v
	return s
}

// SetScalingGroupNames sets the ScalingGroupNames field's value.
func (s *DescribeScalingGroupsInput) SetScalingGroupNames(v []*string) *DescribeScalingGroupsInput {
	s.ScalingGroupNames = v
	return s
}

// SetTagFilters sets the TagFilters field's value.
func (s *DescribeScalingGroupsInput) SetTagFilters(v []*TagFilterForDescribeScalingGroupsInput) *DescribeScalingGroupsInput {
	s.TagFilters = v
	return s
}

type DescribeScalingGroupsOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	PageNumber *int32 `type:"int32"`

	PageSize *int32 `type:"int32"`

	ScalingGroups []*ScalingGroupForDescribeScalingGroupsOutput `type:"list"`

	TotalCount *int32 `type:"int32"`
}

// String returns the string representation
func (s DescribeScalingGroupsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeScalingGroupsOutput) GoString() string {
	return s.String()
}

// SetPageNumber sets the PageNumber field's value.
func (s *DescribeScalingGroupsOutput) SetPageNumber(v int32) *DescribeScalingGroupsOutput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *DescribeScalingGroupsOutput) SetPageSize(v int32) *DescribeScalingGroupsOutput {
	s.PageSize = &v
	return s
}

// SetScalingGroups sets the ScalingGroups field's value.
func (s *DescribeScalingGroupsOutput) SetScalingGroups(v []*ScalingGroupForDescribeScalingGroupsOutput) *DescribeScalingGroupsOutput {
	s.ScalingGroups = v
	return s
}

// SetTotalCount sets the TotalCount field's value.
func (s *DescribeScalingGroupsOutput) SetTotalCount(v int32) *DescribeScalingGroupsOutput {
	s.TotalCount = &v
	return s
}

type LaunchTemplateOverrideForDescribeScalingGroupsOutput struct {
	_ struct{} `type:"structure"`

	InstanceType *string `type:"string"`
}

// String returns the string representation
func (s LaunchTemplateOverrideForDescribeScalingGroupsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s LaunchTemplateOverrideForDescribeScalingGroupsOutput) GoString() string {
	return s.String()
}

// SetInstanceType sets the InstanceType field's value.
func (s *LaunchTemplateOverrideForDescribeScalingGroupsOutput) SetInstanceType(v string) *LaunchTemplateOverrideForDescribeScalingGroupsOutput {
	s.InstanceType = &v
	return s
}

type ScalingGroupForDescribeScalingGroupsOutput struct {
	_ struct{} `type:"structure"`

	ActiveScalingConfigurationId *string `type:"string"`

	CreatedAt *string `type:"string"`

	DBInstanceIds []*string `type:"list"`

	DefaultCooldown *int32 `type:"int32"`

	DesireInstanceNumber *int32 `type:"int32"`

	HealthCheckType *string `type:"string"`

	InstanceTerminatePolicy *string `type:"string"`

	LaunchTemplateId *string `type:"string"`

	LaunchTemplateOverrides []*LaunchTemplateOverrideForDescribeScalingGroupsOutput `type:"list"`

	LaunchTemplateVersion *string `type:"string"`

	LifecycleState *string `type:"string"`

	MaxInstanceNumber *int32 `type:"int32"`

	MinInstanceNumber *int32 `type:"int32"`

	MultiAZPolicy *string `type:"string"`

	ProjectName *string `type:"string"`

	ScalingGroupId *string `type:"string"`

	ScalingGroupName *string `type:"string"`

	ScalingMode *string `type:"string"`

	ServerGroupAttributes []*ServerGroupAttributeForDescribeScalingGroupsOutput `type:"list"`

	StoppedInstanceCount *int32 `type:"int32"`

	SubnetIds []*string `type:"list"`

	Tags []*TagForDescribeScalingGroupsOutput `type:"list"`

	TotalInstanceCount *int32 `type:"int32"`

	UpdatedAt *string `type:"string"`

	VpcId *string `type:"string"`
}

// String returns the string representation
func (s ScalingGroupForDescribeScalingGroupsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ScalingGroupForDescribeScalingGroupsOutput) GoString() string {
	return s.String()
}

// SetActiveScalingConfigurationId sets the ActiveScalingConfigurationId field's value.
func (s *ScalingGroupForDescribeScalingGroupsOutput) SetActiveScalingConfigurationId(v string) *ScalingGroupForDescribeScalingGroupsOutput {
	s.ActiveScalingConfigurationId = &v
	return s
}

// SetCreatedAt sets the CreatedAt field's value.
func (s *ScalingGroupForDescribeScalingGroupsOutput) SetCreatedAt(v string) *ScalingGroupForDescribeScalingGroupsOutput {
	s.CreatedAt = &v
	return s
}

// SetDBInstanceIds sets the DBInstanceIds field's value.
func (s *ScalingGroupForDescribeScalingGroupsOutput) SetDBInstanceIds(v []*string) *ScalingGroupForDescribeScalingGroupsOutput {
	s.DBInstanceIds = v
	return s
}

// SetDefaultCooldown sets the DefaultCooldown field's value.
func (s *ScalingGroupForDescribeScalingGroupsOutput) SetDefaultCooldown(v int32) *ScalingGroupForDescribeScalingGroupsOutput {
	s.DefaultCooldown = &v
	return s
}

// SetDesireInstanceNumber sets the DesireInstanceNumber field's value.
func (s *ScalingGroupForDescribeScalingGroupsOutput) SetDesireInstanceNumber(v int32) *ScalingGroupForDescribeScalingGroupsOutput {
	s.DesireInstanceNumber = &v
	return s
}

// SetHealthCheckType sets the HealthCheckType field's value.
func (s *ScalingGroupForDescribeScalingGroupsOutput) SetHealthCheckType(v string) *ScalingGroupForDescribeScalingGroupsOutput {
	s.HealthCheckType = &v
	return s
}

// SetInstanceTerminatePolicy sets the InstanceTerminatePolicy field's value.
func (s *ScalingGroupForDescribeScalingGroupsOutput) SetInstanceTerminatePolicy(v string) *ScalingGroupForDescribeScalingGroupsOutput {
	s.InstanceTerminatePolicy = &v
	return s
}

// SetLaunchTemplateId sets the LaunchTemplateId field's value.
func (s *ScalingGroupForDescribeScalingGroupsOutput) SetLaunchTemplateId(v string) *ScalingGroupForDescribeScalingGroupsOutput {
	s.LaunchTemplateId = &v
	return s
}

// SetLaunchTemplateOverrides sets the LaunchTemplateOverrides field's value.
func (s *ScalingGroupForDescribeScalingGroupsOutput) SetLaunchTemplateOverrides(v []*LaunchTemplateOverrideForDescribeScalingGroupsOutput) *ScalingGroupForDescribeScalingGroupsOutput {
	s.LaunchTemplateOverrides = v
	return s
}

// SetLaunchTemplateVersion sets the LaunchTemplateVersion field's value.
func (s *ScalingGroupForDescribeScalingGroupsOutput) SetLaunchTemplateVersion(v string) *ScalingGroupForDescribeScalingGroupsOutput {
	s.LaunchTemplateVersion = &v
	return s
}

// SetLifecycleState sets the LifecycleState field's value.
func (s *ScalingGroupForDescribeScalingGroupsOutput) SetLifecycleState(v string) *ScalingGroupForDescribeScalingGroupsOutput {
	s.LifecycleState = &v
	return s
}

// SetMaxInstanceNumber sets the MaxInstanceNumber field's value.
func (s *ScalingGroupForDescribeScalingGroupsOutput) SetMaxInstanceNumber(v int32) *ScalingGroupForDescribeScalingGroupsOutput {
	s.MaxInstanceNumber = &v
	return s
}

// SetMinInstanceNumber sets the MinInstanceNumber field's value.
func (s *ScalingGroupForDescribeScalingGroupsOutput) SetMinInstanceNumber(v int32) *ScalingGroupForDescribeScalingGroupsOutput {
	s.MinInstanceNumber = &v
	return s
}

// SetMultiAZPolicy sets the MultiAZPolicy field's value.
func (s *ScalingGroupForDescribeScalingGroupsOutput) SetMultiAZPolicy(v string) *ScalingGroupForDescribeScalingGroupsOutput {
	s.MultiAZPolicy = &v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *ScalingGroupForDescribeScalingGroupsOutput) SetProjectName(v string) *ScalingGroupForDescribeScalingGroupsOutput {
	s.ProjectName = &v
	return s
}

// SetScalingGroupId sets the ScalingGroupId field's value.
func (s *ScalingGroupForDescribeScalingGroupsOutput) SetScalingGroupId(v string) *ScalingGroupForDescribeScalingGroupsOutput {
	s.ScalingGroupId = &v
	return s
}

// SetScalingGroupName sets the ScalingGroupName field's value.
func (s *ScalingGroupForDescribeScalingGroupsOutput) SetScalingGroupName(v string) *ScalingGroupForDescribeScalingGroupsOutput {
	s.ScalingGroupName = &v
	return s
}

// SetScalingMode sets the ScalingMode field's value.
func (s *ScalingGroupForDescribeScalingGroupsOutput) SetScalingMode(v string) *ScalingGroupForDescribeScalingGroupsOutput {
	s.ScalingMode = &v
	return s
}

// SetServerGroupAttributes sets the ServerGroupAttributes field's value.
func (s *ScalingGroupForDescribeScalingGroupsOutput) SetServerGroupAttributes(v []*ServerGroupAttributeForDescribeScalingGroupsOutput) *ScalingGroupForDescribeScalingGroupsOutput {
	s.ServerGroupAttributes = v
	return s
}

// SetStoppedInstanceCount sets the StoppedInstanceCount field's value.
func (s *ScalingGroupForDescribeScalingGroupsOutput) SetStoppedInstanceCount(v int32) *ScalingGroupForDescribeScalingGroupsOutput {
	s.StoppedInstanceCount = &v
	return s
}

// SetSubnetIds sets the SubnetIds field's value.
func (s *ScalingGroupForDescribeScalingGroupsOutput) SetSubnetIds(v []*string) *ScalingGroupForDescribeScalingGroupsOutput {
	s.SubnetIds = v
	return s
}

// SetTags sets the Tags field's value.
func (s *ScalingGroupForDescribeScalingGroupsOutput) SetTags(v []*TagForDescribeScalingGroupsOutput) *ScalingGroupForDescribeScalingGroupsOutput {
	s.Tags = v
	return s
}

// SetTotalInstanceCount sets the TotalInstanceCount field's value.
func (s *ScalingGroupForDescribeScalingGroupsOutput) SetTotalInstanceCount(v int32) *ScalingGroupForDescribeScalingGroupsOutput {
	s.TotalInstanceCount = &v
	return s
}

// SetUpdatedAt sets the UpdatedAt field's value.
func (s *ScalingGroupForDescribeScalingGroupsOutput) SetUpdatedAt(v string) *ScalingGroupForDescribeScalingGroupsOutput {
	s.UpdatedAt = &v
	return s
}

// SetVpcId sets the VpcId field's value.
func (s *ScalingGroupForDescribeScalingGroupsOutput) SetVpcId(v string) *ScalingGroupForDescribeScalingGroupsOutput {
	s.VpcId = &v
	return s
}

type ServerGroupAttributeForDescribeScalingGroupsOutput struct {
	_ struct{} `type:"structure"`

	LoadBalancerId *string `type:"string"`

	Port *int32 `type:"int32"`

	ServerGroupId *string `type:"string"`

	Weight *int32 `type:"int32"`
}

// String returns the string representation
func (s ServerGroupAttributeForDescribeScalingGroupsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ServerGroupAttributeForDescribeScalingGroupsOutput) GoString() string {
	return s.String()
}

// SetLoadBalancerId sets the LoadBalancerId field's value.
func (s *ServerGroupAttributeForDescribeScalingGroupsOutput) SetLoadBalancerId(v string) *ServerGroupAttributeForDescribeScalingGroupsOutput {
	s.LoadBalancerId = &v
	return s
}

// SetPort sets the Port field's value.
func (s *ServerGroupAttributeForDescribeScalingGroupsOutput) SetPort(v int32) *ServerGroupAttributeForDescribeScalingGroupsOutput {
	s.Port = &v
	return s
}

// SetServerGroupId sets the ServerGroupId field's value.
func (s *ServerGroupAttributeForDescribeScalingGroupsOutput) SetServerGroupId(v string) *ServerGroupAttributeForDescribeScalingGroupsOutput {
	s.ServerGroupId = &v
	return s
}

// SetWeight sets the Weight field's value.
func (s *ServerGroupAttributeForDescribeScalingGroupsOutput) SetWeight(v int32) *ServerGroupAttributeForDescribeScalingGroupsOutput {
	s.Weight = &v
	return s
}

type TagFilterForDescribeScalingGroupsInput struct {
	_ struct{} `type:"structure"`

	Key *string `type:"string"`

	Value *string `type:"string"`
}

// String returns the string representation
func (s TagFilterForDescribeScalingGroupsInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s TagFilterForDescribeScalingGroupsInput) GoString() string {
	return s.String()
}

// SetKey sets the Key field's value.
func (s *TagFilterForDescribeScalingGroupsInput) SetKey(v string) *TagFilterForDescribeScalingGroupsInput {
	s.Key = &v
	return s
}

// SetValue sets the Value field's value.
func (s *TagFilterForDescribeScalingGroupsInput) SetValue(v string) *TagFilterForDescribeScalingGroupsInput {
	s.Value = &v
	return s
}

type TagForDescribeScalingGroupsOutput struct {
	_ struct{} `type:"structure"`

	Key *string `type:"string"`

	Value *string `type:"string"`
}

// String returns the string representation
func (s TagForDescribeScalingGroupsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s TagForDescribeScalingGroupsOutput) GoString() string {
	return s.String()
}

// SetKey sets the Key field's value.
func (s *TagForDescribeScalingGroupsOutput) SetKey(v string) *TagForDescribeScalingGroupsOutput {
	s.Key = &v
	return s
}

// SetValue sets the Value field's value.
func (s *TagForDescribeScalingGroupsOutput) SetValue(v string) *TagForDescribeScalingGroupsOutput {
	s.Value = &v
	return s
}

// Code generated by volcstack with private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package autoscalingiface provides an interface to enable mocking the AUTO_SCALING service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package autoscaling

import (
	"github.com/volcengine/volcstack-go-sdk/volcstack"
	"github.com/volcengine/volcstack-go-sdk/volcstack/request"
)

// AUTOSCALINGAPI provides an interface to enable mocking the
// autoscaling.AUTOSCALING service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // AUTO_SCALING.
//    func myFunc(svc AUTOSCALINGAPI) bool {
//        // Make svc.AttachInstances request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := autoscaling.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockAUTOSCALINGClient struct {
//        AUTOSCALINGAPI
//    }
//    func (m *mockAUTOSCALINGClient) AttachInstances(input *AttachInstancesInput) (*AttachInstancesOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockAUTOSCALINGClient{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type AUTOSCALINGAPI interface {
	AttachInstancesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	AttachInstancesCommonWithContext(volcstack.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	AttachInstancesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	AttachInstances(*AttachInstancesInput) (*AttachInstancesOutput, error)
	AttachInstancesWithContext(volcstack.Context, *AttachInstancesInput, ...request.Option) (*AttachInstancesOutput, error)
	AttachInstancesRequest(*AttachInstancesInput) (*request.Request, *AttachInstancesOutput)

	AttachLoadBalancersCommon(*map[string]interface{}) (*map[string]interface{}, error)
	AttachLoadBalancersCommonWithContext(volcstack.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	AttachLoadBalancersCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	AttachLoadBalancers(*AttachLoadBalancersInput) (*AttachLoadBalancersOutput, error)
	AttachLoadBalancersWithContext(volcstack.Context, *AttachLoadBalancersInput, ...request.Option) (*AttachLoadBalancersOutput, error)
	AttachLoadBalancersRequest(*AttachLoadBalancersInput) (*request.Request, *AttachLoadBalancersOutput)

	CreateScalingConfigurationCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CreateScalingConfigurationCommonWithContext(volcstack.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateScalingConfigurationCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateScalingConfiguration(*CreateScalingConfigurationInput) (*CreateScalingConfigurationOutput, error)
	CreateScalingConfigurationWithContext(volcstack.Context, *CreateScalingConfigurationInput, ...request.Option) (*CreateScalingConfigurationOutput, error)
	CreateScalingConfigurationRequest(*CreateScalingConfigurationInput) (*request.Request, *CreateScalingConfigurationOutput)

	CreateScalingGroupCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CreateScalingGroupCommonWithContext(volcstack.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateScalingGroupCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateScalingGroup(*CreateScalingGroupInput) (*CreateScalingGroupOutput, error)
	CreateScalingGroupWithContext(volcstack.Context, *CreateScalingGroupInput, ...request.Option) (*CreateScalingGroupOutput, error)
	CreateScalingGroupRequest(*CreateScalingGroupInput) (*request.Request, *CreateScalingGroupOutput)

	CreateScalingPolicyCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CreateScalingPolicyCommonWithContext(volcstack.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateScalingPolicyCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateScalingPolicy(*CreateScalingPolicyInput) (*CreateScalingPolicyOutput, error)
	CreateScalingPolicyWithContext(volcstack.Context, *CreateScalingPolicyInput, ...request.Option) (*CreateScalingPolicyOutput, error)
	CreateScalingPolicyRequest(*CreateScalingPolicyInput) (*request.Request, *CreateScalingPolicyOutput)

	DeleteScalingConfigurationCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteScalingConfigurationCommonWithContext(volcstack.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteScalingConfigurationCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteScalingConfiguration(*DeleteScalingConfigurationInput) (*DeleteScalingConfigurationOutput, error)
	DeleteScalingConfigurationWithContext(volcstack.Context, *DeleteScalingConfigurationInput, ...request.Option) (*DeleteScalingConfigurationOutput, error)
	DeleteScalingConfigurationRequest(*DeleteScalingConfigurationInput) (*request.Request, *DeleteScalingConfigurationOutput)

	DeleteScalingGroupCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteScalingGroupCommonWithContext(volcstack.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteScalingGroupCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteScalingGroup(*DeleteScalingGroupInput) (*DeleteScalingGroupOutput, error)
	DeleteScalingGroupWithContext(volcstack.Context, *DeleteScalingGroupInput, ...request.Option) (*DeleteScalingGroupOutput, error)
	DeleteScalingGroupRequest(*DeleteScalingGroupInput) (*request.Request, *DeleteScalingGroupOutput)

	DeleteScalingPolicyCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteScalingPolicyCommonWithContext(volcstack.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteScalingPolicyCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteScalingPolicy(*DeleteScalingPolicyInput) (*DeleteScalingPolicyOutput, error)
	DeleteScalingPolicyWithContext(volcstack.Context, *DeleteScalingPolicyInput, ...request.Option) (*DeleteScalingPolicyOutput, error)
	DeleteScalingPolicyRequest(*DeleteScalingPolicyInput) (*request.Request, *DeleteScalingPolicyOutput)

	DescribeScalingActivitiesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeScalingActivitiesCommonWithContext(volcstack.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeScalingActivitiesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeScalingActivities(*DescribeScalingActivitiesInput) (*DescribeScalingActivitiesOutput, error)
	DescribeScalingActivitiesWithContext(volcstack.Context, *DescribeScalingActivitiesInput, ...request.Option) (*DescribeScalingActivitiesOutput, error)
	DescribeScalingActivitiesRequest(*DescribeScalingActivitiesInput) (*request.Request, *DescribeScalingActivitiesOutput)

	DescribeScalingConfigurationsCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeScalingConfigurationsCommonWithContext(volcstack.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeScalingConfigurationsCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeScalingConfigurations(*DescribeScalingConfigurationsInput) (*DescribeScalingConfigurationsOutput, error)
	DescribeScalingConfigurationsWithContext(volcstack.Context, *DescribeScalingConfigurationsInput, ...request.Option) (*DescribeScalingConfigurationsOutput, error)
	DescribeScalingConfigurationsRequest(*DescribeScalingConfigurationsInput) (*request.Request, *DescribeScalingConfigurationsOutput)

	DescribeScalingGroupsCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeScalingGroupsCommonWithContext(volcstack.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeScalingGroupsCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeScalingGroups(*DescribeScalingGroupsInput) (*DescribeScalingGroupsOutput, error)
	DescribeScalingGroupsWithContext(volcstack.Context, *DescribeScalingGroupsInput, ...request.Option) (*DescribeScalingGroupsOutput, error)
	DescribeScalingGroupsRequest(*DescribeScalingGroupsInput) (*request.Request, *DescribeScalingGroupsOutput)

	DescribeScalingInstancesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeScalingInstancesCommonWithContext(volcstack.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeScalingInstancesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeScalingInstances(*DescribeScalingInstancesInput) (*DescribeScalingInstancesOutput, error)
	DescribeScalingInstancesWithContext(volcstack.Context, *DescribeScalingInstancesInput, ...request.Option) (*DescribeScalingInstancesOutput, error)
	DescribeScalingInstancesRequest(*DescribeScalingInstancesInput) (*request.Request, *DescribeScalingInstancesOutput)

	DescribeScalingPoliciesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeScalingPoliciesCommonWithContext(volcstack.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeScalingPoliciesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeScalingPolicies(*DescribeScalingPoliciesInput) (*DescribeScalingPoliciesOutput, error)
	DescribeScalingPoliciesWithContext(volcstack.Context, *DescribeScalingPoliciesInput, ...request.Option) (*DescribeScalingPoliciesOutput, error)
	DescribeScalingPoliciesRequest(*DescribeScalingPoliciesInput) (*request.Request, *DescribeScalingPoliciesOutput)

	DetachInstancesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DetachInstancesCommonWithContext(volcstack.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DetachInstancesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DetachInstances(*DetachInstancesInput) (*DetachInstancesOutput, error)
	DetachInstancesWithContext(volcstack.Context, *DetachInstancesInput, ...request.Option) (*DetachInstancesOutput, error)
	DetachInstancesRequest(*DetachInstancesInput) (*request.Request, *DetachInstancesOutput)

	DetachLoadBalancersCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DetachLoadBalancersCommonWithContext(volcstack.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DetachLoadBalancersCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DetachLoadBalancers(*DetachLoadBalancersInput) (*DetachLoadBalancersOutput, error)
	DetachLoadBalancersWithContext(volcstack.Context, *DetachLoadBalancersInput, ...request.Option) (*DetachLoadBalancersOutput, error)
	DetachLoadBalancersRequest(*DetachLoadBalancersInput) (*request.Request, *DetachLoadBalancersOutput)

	DisableScalingGroupCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DisableScalingGroupCommonWithContext(volcstack.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DisableScalingGroupCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DisableScalingGroup(*DisableScalingGroupInput) (*DisableScalingGroupOutput, error)
	DisableScalingGroupWithContext(volcstack.Context, *DisableScalingGroupInput, ...request.Option) (*DisableScalingGroupOutput, error)
	DisableScalingGroupRequest(*DisableScalingGroupInput) (*request.Request, *DisableScalingGroupOutput)

	EnableScalingConfigurationCommon(*map[string]interface{}) (*map[string]interface{}, error)
	EnableScalingConfigurationCommonWithContext(volcstack.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	EnableScalingConfigurationCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	EnableScalingConfiguration(*EnableScalingConfigurationInput) (*EnableScalingConfigurationOutput, error)
	EnableScalingConfigurationWithContext(volcstack.Context, *EnableScalingConfigurationInput, ...request.Option) (*EnableScalingConfigurationOutput, error)
	EnableScalingConfigurationRequest(*EnableScalingConfigurationInput) (*request.Request, *EnableScalingConfigurationOutput)

	EnableScalingGroupCommon(*map[string]interface{}) (*map[string]interface{}, error)
	EnableScalingGroupCommonWithContext(volcstack.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	EnableScalingGroupCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	EnableScalingGroup(*EnableScalingGroupInput) (*EnableScalingGroupOutput, error)
	EnableScalingGroupWithContext(volcstack.Context, *EnableScalingGroupInput, ...request.Option) (*EnableScalingGroupOutput, error)
	EnableScalingGroupRequest(*EnableScalingGroupInput) (*request.Request, *EnableScalingGroupOutput)

	ModifyInstancesProtectionCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ModifyInstancesProtectionCommonWithContext(volcstack.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifyInstancesProtectionCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifyInstancesProtection(*ModifyInstancesProtectionInput) (*ModifyInstancesProtectionOutput, error)
	ModifyInstancesProtectionWithContext(volcstack.Context, *ModifyInstancesProtectionInput, ...request.Option) (*ModifyInstancesProtectionOutput, error)
	ModifyInstancesProtectionRequest(*ModifyInstancesProtectionInput) (*request.Request, *ModifyInstancesProtectionOutput)

	ModifyScalingConfigurationCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ModifyScalingConfigurationCommonWithContext(volcstack.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifyScalingConfigurationCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifyScalingConfiguration(*ModifyScalingConfigurationInput) (*ModifyScalingConfigurationOutput, error)
	ModifyScalingConfigurationWithContext(volcstack.Context, *ModifyScalingConfigurationInput, ...request.Option) (*ModifyScalingConfigurationOutput, error)
	ModifyScalingConfigurationRequest(*ModifyScalingConfigurationInput) (*request.Request, *ModifyScalingConfigurationOutput)

	ModifyScalingGroupCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ModifyScalingGroupCommonWithContext(volcstack.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifyScalingGroupCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifyScalingGroup(*ModifyScalingGroupInput) (*ModifyScalingGroupOutput, error)
	ModifyScalingGroupWithContext(volcstack.Context, *ModifyScalingGroupInput, ...request.Option) (*ModifyScalingGroupOutput, error)
	ModifyScalingGroupRequest(*ModifyScalingGroupInput) (*request.Request, *ModifyScalingGroupOutput)

	ModifyScalingPolicyCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ModifyScalingPolicyCommonWithContext(volcstack.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifyScalingPolicyCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifyScalingPolicy(*ModifyScalingPolicyInput) (*ModifyScalingPolicyOutput, error)
	ModifyScalingPolicyWithContext(volcstack.Context, *ModifyScalingPolicyInput, ...request.Option) (*ModifyScalingPolicyOutput, error)
	ModifyScalingPolicyRequest(*ModifyScalingPolicyInput) (*request.Request, *ModifyScalingPolicyOutput)

	RemoveInstancesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	RemoveInstancesCommonWithContext(volcstack.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	RemoveInstancesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	RemoveInstances(*RemoveInstancesInput) (*RemoveInstancesOutput, error)
	RemoveInstancesWithContext(volcstack.Context, *RemoveInstancesInput, ...request.Option) (*RemoveInstancesOutput, error)
	RemoveInstancesRequest(*RemoveInstancesInput) (*request.Request, *RemoveInstancesOutput)

	StartScalingPolicyCommon(*map[string]interface{}) (*map[string]interface{}, error)
	StartScalingPolicyCommonWithContext(volcstack.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	StartScalingPolicyCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	StartScalingPolicy(*StartScalingPolicyInput) (*StartScalingPolicyOutput, error)
	StartScalingPolicyWithContext(volcstack.Context, *StartScalingPolicyInput, ...request.Option) (*StartScalingPolicyOutput, error)
	StartScalingPolicyRequest(*StartScalingPolicyInput) (*request.Request, *StartScalingPolicyOutput)

	StopScalingPolicyCommon(*map[string]interface{}) (*map[string]interface{}, error)
	StopScalingPolicyCommonWithContext(volcstack.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	StopScalingPolicyCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	StopScalingPolicy(*StopScalingPolicyInput) (*StopScalingPolicyOutput, error)
	StopScalingPolicyWithContext(volcstack.Context, *StopScalingPolicyInput, ...request.Option) (*StopScalingPolicyOutput, error)
	StopScalingPolicyRequest(*StopScalingPolicyInput) (*request.Request, *StopScalingPolicyOutput)
}

var _ AUTOSCALINGAPI = (*AUTOSCALING)(nil)

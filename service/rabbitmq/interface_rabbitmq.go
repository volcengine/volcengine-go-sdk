// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package rabbitmqiface provides an interface to enable mocking the RABBITMQ service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package rabbitmq

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
)

// RABBITMQAPI provides an interface to enable mocking the
// rabbitmq.RABBITMQ service client's API operation,
//
//    // volcengine sdk func uses an SDK service client to make a request to
//    // RABBITMQ.
//    func myFunc(svc RABBITMQAPI) bool {
//        // Make svc.AddTagsToResource request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := rabbitmq.New(sess)
//
//        myFunc(svc)
//    }
//
type RABBITMQAPI interface {
	AddTagsToResourceCommon(*map[string]interface{}) (*map[string]interface{}, error)
	AddTagsToResourceCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	AddTagsToResourceCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	AddTagsToResource(*AddTagsToResourceInput) (*AddTagsToResourceOutput, error)
	AddTagsToResourceWithContext(volcengine.Context, *AddTagsToResourceInput, ...request.Option) (*AddTagsToResourceOutput, error)
	AddTagsToResourceRequest(*AddTagsToResourceInput) (*request.Request, *AddTagsToResourceOutput)

	CreateInstanceCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CreateInstanceCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateInstanceCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateInstance(*CreateInstanceInput) (*CreateInstanceOutput, error)
	CreateInstanceWithContext(volcengine.Context, *CreateInstanceInput, ...request.Option) (*CreateInstanceOutput, error)
	CreateInstanceRequest(*CreateInstanceInput) (*request.Request, *CreateInstanceOutput)

	CreatePublicAddressCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CreatePublicAddressCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreatePublicAddressCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreatePublicAddress(*CreatePublicAddressInput) (*CreatePublicAddressOutput, error)
	CreatePublicAddressWithContext(volcengine.Context, *CreatePublicAddressInput, ...request.Option) (*CreatePublicAddressOutput, error)
	CreatePublicAddressRequest(*CreatePublicAddressInput) (*request.Request, *CreatePublicAddressOutput)

	DeleteInstanceCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteInstanceCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteInstanceCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteInstance(*DeleteInstanceInput) (*DeleteInstanceOutput, error)
	DeleteInstanceWithContext(volcengine.Context, *DeleteInstanceInput, ...request.Option) (*DeleteInstanceOutput, error)
	DeleteInstanceRequest(*DeleteInstanceInput) (*request.Request, *DeleteInstanceOutput)

	DeletePublicAddressCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DeletePublicAddressCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeletePublicAddressCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeletePublicAddress(*DeletePublicAddressInput) (*DeletePublicAddressOutput, error)
	DeletePublicAddressWithContext(volcengine.Context, *DeletePublicAddressInput, ...request.Option) (*DeletePublicAddressOutput, error)
	DeletePublicAddressRequest(*DeletePublicAddressInput) (*request.Request, *DeletePublicAddressOutput)

	DescribeAvailabilityZonesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeAvailabilityZonesCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeAvailabilityZonesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeAvailabilityZones(*DescribeAvailabilityZonesInput) (*DescribeAvailabilityZonesOutput, error)
	DescribeAvailabilityZonesWithContext(volcengine.Context, *DescribeAvailabilityZonesInput, ...request.Option) (*DescribeAvailabilityZonesOutput, error)
	DescribeAvailabilityZonesRequest(*DescribeAvailabilityZonesInput) (*request.Request, *DescribeAvailabilityZonesOutput)

	DescribeInstanceDetailCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeInstanceDetailCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeInstanceDetailCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeInstanceDetail(*DescribeInstanceDetailInput) (*DescribeInstanceDetailOutput, error)
	DescribeInstanceDetailWithContext(volcengine.Context, *DescribeInstanceDetailInput, ...request.Option) (*DescribeInstanceDetailOutput, error)
	DescribeInstanceDetailRequest(*DescribeInstanceDetailInput) (*request.Request, *DescribeInstanceDetailOutput)

	DescribeInstancesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeInstancesCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeInstancesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeInstances(*DescribeInstancesInput) (*DescribeInstancesOutput, error)
	DescribeInstancesWithContext(volcengine.Context, *DescribeInstancesInput, ...request.Option) (*DescribeInstancesOutput, error)
	DescribeInstancesRequest(*DescribeInstancesInput) (*request.Request, *DescribeInstancesOutput)

	DescribePluginsCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribePluginsCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribePluginsCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribePlugins(*DescribePluginsInput) (*DescribePluginsOutput, error)
	DescribePluginsWithContext(volcengine.Context, *DescribePluginsInput, ...request.Option) (*DescribePluginsOutput, error)
	DescribePluginsRequest(*DescribePluginsInput) (*request.Request, *DescribePluginsOutput)

	DescribeRegionsCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeRegionsCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeRegionsCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeRegions(*DescribeRegionsInput) (*DescribeRegionsOutput, error)
	DescribeRegionsWithContext(volcengine.Context, *DescribeRegionsInput, ...request.Option) (*DescribeRegionsOutput, error)
	DescribeRegionsRequest(*DescribeRegionsInput) (*request.Request, *DescribeRegionsOutput)

	DescribeTagsByResourceCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeTagsByResourceCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeTagsByResourceCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeTagsByResource(*DescribeTagsByResourceInput) (*DescribeTagsByResourceOutput, error)
	DescribeTagsByResourceWithContext(volcengine.Context, *DescribeTagsByResourceInput, ...request.Option) (*DescribeTagsByResourceOutput, error)
	DescribeTagsByResourceRequest(*DescribeTagsByResourceInput) (*request.Request, *DescribeTagsByResourceOutput)

	ModifyInstanceAttributesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ModifyInstanceAttributesCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifyInstanceAttributesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifyInstanceAttributes(*ModifyInstanceAttributesInput) (*ModifyInstanceAttributesOutput, error)
	ModifyInstanceAttributesWithContext(volcengine.Context, *ModifyInstanceAttributesInput, ...request.Option) (*ModifyInstanceAttributesOutput, error)
	ModifyInstanceAttributesRequest(*ModifyInstanceAttributesInput) (*request.Request, *ModifyInstanceAttributesOutput)

	ModifyInstanceChargeTypeCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ModifyInstanceChargeTypeCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifyInstanceChargeTypeCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifyInstanceChargeType(*ModifyInstanceChargeTypeInput) (*ModifyInstanceChargeTypeOutput, error)
	ModifyInstanceChargeTypeWithContext(volcengine.Context, *ModifyInstanceChargeTypeInput, ...request.Option) (*ModifyInstanceChargeTypeOutput, error)
	ModifyInstanceChargeTypeRequest(*ModifyInstanceChargeTypeInput) (*request.Request, *ModifyInstanceChargeTypeOutput)

	ModifyInstanceSpecCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ModifyInstanceSpecCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifyInstanceSpecCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifyInstanceSpec(*ModifyInstanceSpecInput) (*ModifyInstanceSpecOutput, error)
	ModifyInstanceSpecWithContext(volcengine.Context, *ModifyInstanceSpecInput, ...request.Option) (*ModifyInstanceSpecOutput, error)
	ModifyInstanceSpecRequest(*ModifyInstanceSpecInput) (*request.Request, *ModifyInstanceSpecOutput)

	ModifyPluginCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ModifyPluginCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifyPluginCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifyPlugin(*ModifyPluginInput) (*ModifyPluginOutput, error)
	ModifyPluginWithContext(volcengine.Context, *ModifyPluginInput, ...request.Option) (*ModifyPluginOutput, error)
	ModifyPluginRequest(*ModifyPluginInput) (*request.Request, *ModifyPluginOutput)

	ModifyUserPasswordCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ModifyUserPasswordCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifyUserPasswordCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifyUserPassword(*ModifyUserPasswordInput) (*ModifyUserPasswordOutput, error)
	ModifyUserPasswordWithContext(volcengine.Context, *ModifyUserPasswordInput, ...request.Option) (*ModifyUserPasswordOutput, error)
	ModifyUserPasswordRequest(*ModifyUserPasswordInput) (*request.Request, *ModifyUserPasswordOutput)

	RemoveTagsFromResourceCommon(*map[string]interface{}) (*map[string]interface{}, error)
	RemoveTagsFromResourceCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	RemoveTagsFromResourceCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	RemoveTagsFromResource(*RemoveTagsFromResourceInput) (*RemoveTagsFromResourceOutput, error)
	RemoveTagsFromResourceWithContext(volcengine.Context, *RemoveTagsFromResourceInput, ...request.Option) (*RemoveTagsFromResourceOutput, error)
	RemoveTagsFromResourceRequest(*RemoveTagsFromResourceInput) (*request.Request, *RemoveTagsFromResourceOutput)

	RestartInstanceCommon(*map[string]interface{}) (*map[string]interface{}, error)
	RestartInstanceCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	RestartInstanceCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	RestartInstance(*RestartInstanceInput) (*RestartInstanceOutput, error)
	RestartInstanceWithContext(volcengine.Context, *RestartInstanceInput, ...request.Option) (*RestartInstanceOutput, error)
	RestartInstanceRequest(*RestartInstanceInput) (*request.Request, *RestartInstanceOutput)
}

var _ RABBITMQAPI = (*RABBITMQ)(nil)
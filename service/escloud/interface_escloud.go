// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package escloudiface provides an interface to enable mocking the ESCLOUD service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package escloud

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
)

// ESCLOUDAPI provides an interface to enable mocking the
// escloud.ESCLOUD service client's API operation,
//
//    // volcengine sdk func uses an SDK service client to make a request to
//    // ESCLOUD.
//    func myFunc(svc ESCLOUDAPI) bool {
//        // Make svc.CreateInstance request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := escloud.New(sess)
//
//        myFunc(svc)
//    }
//
type ESCLOUDAPI interface {
	CreateInstanceCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CreateInstanceCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateInstanceCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateInstance(*CreateInstanceInput) (*CreateInstanceOutput, error)
	CreateInstanceWithContext(volcengine.Context, *CreateInstanceInput, ...request.Option) (*CreateInstanceOutput, error)
	CreateInstanceRequest(*CreateInstanceInput) (*request.Request, *CreateInstanceOutput)

	CreateInstanceInOneStepCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CreateInstanceInOneStepCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateInstanceInOneStepCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateInstanceInOneStep(*CreateInstanceInOneStepInput) (*CreateInstanceInOneStepOutput, error)
	CreateInstanceInOneStepWithContext(volcengine.Context, *CreateInstanceInOneStepInput, ...request.Option) (*CreateInstanceInOneStepOutput, error)
	CreateInstanceInOneStepRequest(*CreateInstanceInOneStepInput) (*request.Request, *CreateInstanceInOneStepOutput)

	DescribeInstanceCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeInstanceCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeInstanceCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeInstance(*DescribeInstanceInput) (*DescribeInstanceOutput, error)
	DescribeInstanceWithContext(volcengine.Context, *DescribeInstanceInput, ...request.Option) (*DescribeInstanceOutput, error)
	DescribeInstanceRequest(*DescribeInstanceInput) (*request.Request, *DescribeInstanceOutput)

	DescribeInstanceNodesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeInstanceNodesCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeInstanceNodesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeInstanceNodes(*DescribeInstanceNodesInput) (*DescribeInstanceNodesOutput, error)
	DescribeInstanceNodesWithContext(volcengine.Context, *DescribeInstanceNodesInput, ...request.Option) (*DescribeInstanceNodesOutput, error)
	DescribeInstanceNodesRequest(*DescribeInstanceNodesInput) (*request.Request, *DescribeInstanceNodesOutput)

	DescribeInstancePluginsCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeInstancePluginsCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeInstancePluginsCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeInstancePlugins(*DescribeInstancePluginsInput) (*DescribeInstancePluginsOutput, error)
	DescribeInstancePluginsWithContext(volcengine.Context, *DescribeInstancePluginsInput, ...request.Option) (*DescribeInstancePluginsOutput, error)
	DescribeInstancePluginsRequest(*DescribeInstancePluginsInput) (*request.Request, *DescribeInstancePluginsOutput)

	DescribeInstancesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeInstancesCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeInstancesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeInstances(*DescribeInstancesInput) (*DescribeInstancesOutput, error)
	DescribeInstancesWithContext(volcengine.Context, *DescribeInstancesInput, ...request.Option) (*DescribeInstancesOutput, error)
	DescribeInstancesRequest(*DescribeInstancesInput) (*request.Request, *DescribeInstancesOutput)

	DescribeNodeAvailableSpecsCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeNodeAvailableSpecsCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeNodeAvailableSpecsCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeNodeAvailableSpecs(*DescribeNodeAvailableSpecsInput) (*DescribeNodeAvailableSpecsOutput, error)
	DescribeNodeAvailableSpecsWithContext(volcengine.Context, *DescribeNodeAvailableSpecsInput, ...request.Option) (*DescribeNodeAvailableSpecsOutput, error)
	DescribeNodeAvailableSpecsRequest(*DescribeNodeAvailableSpecsInput) (*request.Request, *DescribeNodeAvailableSpecsOutput)

	DescribeZonesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeZonesCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeZonesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeZones(*DescribeZonesInput) (*DescribeZonesOutput, error)
	DescribeZonesWithContext(volcengine.Context, *DescribeZonesInput, ...request.Option) (*DescribeZonesOutput, error)
	DescribeZonesRequest(*DescribeZonesInput) (*request.Request, *DescribeZonesOutput)

	ListTagsForResourcesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ListTagsForResourcesCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ListTagsForResourcesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ListTagsForResources(*ListTagsForResourcesInput) (*ListTagsForResourcesOutput, error)
	ListTagsForResourcesWithContext(volcengine.Context, *ListTagsForResourcesInput, ...request.Option) (*ListTagsForResourcesOutput, error)
	ListTagsForResourcesRequest(*ListTagsForResourcesInput) (*request.Request, *ListTagsForResourcesOutput)

	ModifyChargeCodeCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ModifyChargeCodeCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifyChargeCodeCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifyChargeCode(*ModifyChargeCodeInput) (*ModifyChargeCodeOutput, error)
	ModifyChargeCodeWithContext(volcengine.Context, *ModifyChargeCodeInput, ...request.Option) (*ModifyChargeCodeOutput, error)
	ModifyChargeCodeRequest(*ModifyChargeCodeInput) (*request.Request, *ModifyChargeCodeOutput)

	ModifyDeletionProtectionCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ModifyDeletionProtectionCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifyDeletionProtectionCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifyDeletionProtection(*ModifyDeletionProtectionInput) (*ModifyDeletionProtectionOutput, error)
	ModifyDeletionProtectionWithContext(volcengine.Context, *ModifyDeletionProtectionInput, ...request.Option) (*ModifyDeletionProtectionOutput, error)
	ModifyDeletionProtectionRequest(*ModifyDeletionProtectionInput) (*request.Request, *ModifyDeletionProtectionOutput)

	ModifyIpAllowListCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ModifyIpAllowListCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifyIpAllowListCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifyIpAllowList(*ModifyIpAllowListInput) (*ModifyIpAllowListOutput, error)
	ModifyIpAllowListWithContext(volcengine.Context, *ModifyIpAllowListInput, ...request.Option) (*ModifyIpAllowListOutput, error)
	ModifyIpAllowListRequest(*ModifyIpAllowListInput) (*request.Request, *ModifyIpAllowListOutput)

	ModifyMaintenanceSettingCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ModifyMaintenanceSettingCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifyMaintenanceSettingCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifyMaintenanceSetting(*ModifyMaintenanceSettingInput) (*ModifyMaintenanceSettingOutput, error)
	ModifyMaintenanceSettingWithContext(volcengine.Context, *ModifyMaintenanceSettingInput, ...request.Option) (*ModifyMaintenanceSettingOutput, error)
	ModifyMaintenanceSettingRequest(*ModifyMaintenanceSettingInput) (*request.Request, *ModifyMaintenanceSettingOutput)

	ModifyNodeSpecInOneStepCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ModifyNodeSpecInOneStepCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifyNodeSpecInOneStepCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifyNodeSpecInOneStep(*ModifyNodeSpecInOneStepInput) (*ModifyNodeSpecInOneStepOutput, error)
	ModifyNodeSpecInOneStepWithContext(volcengine.Context, *ModifyNodeSpecInOneStepInput, ...request.Option) (*ModifyNodeSpecInOneStepOutput, error)
	ModifyNodeSpecInOneStepRequest(*ModifyNodeSpecInOneStepInput) (*request.Request, *ModifyNodeSpecInOneStepOutput)

	ReleaseInstanceCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ReleaseInstanceCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ReleaseInstanceCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ReleaseInstance(*ReleaseInstanceInput) (*ReleaseInstanceOutput, error)
	ReleaseInstanceWithContext(volcengine.Context, *ReleaseInstanceInput, ...request.Option) (*ReleaseInstanceOutput, error)
	ReleaseInstanceRequest(*ReleaseInstanceInput) (*request.Request, *ReleaseInstanceOutput)

	RenameInstanceCommon(*map[string]interface{}) (*map[string]interface{}, error)
	RenameInstanceCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	RenameInstanceCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	RenameInstance(*RenameInstanceInput) (*RenameInstanceOutput, error)
	RenameInstanceWithContext(volcengine.Context, *RenameInstanceInput, ...request.Option) (*RenameInstanceOutput, error)
	RenameInstanceRequest(*RenameInstanceInput) (*request.Request, *RenameInstanceOutput)

	ResetAdminPasswordCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ResetAdminPasswordCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ResetAdminPasswordCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ResetAdminPassword(*ResetAdminPasswordInput) (*ResetAdminPasswordOutput, error)
	ResetAdminPasswordWithContext(volcengine.Context, *ResetAdminPasswordInput, ...request.Option) (*ResetAdminPasswordOutput, error)
	ResetAdminPasswordRequest(*ResetAdminPasswordInput) (*request.Request, *ResetAdminPasswordOutput)

	RestartInstanceCommon(*map[string]interface{}) (*map[string]interface{}, error)
	RestartInstanceCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	RestartInstanceCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	RestartInstance(*RestartInstanceInput) (*RestartInstanceOutput, error)
	RestartInstanceWithContext(volcengine.Context, *RestartInstanceInput, ...request.Option) (*RestartInstanceOutput, error)
	RestartInstanceRequest(*RestartInstanceInput) (*request.Request, *RestartInstanceOutput)

	RestartNodeCommon(*map[string]interface{}) (*map[string]interface{}, error)
	RestartNodeCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	RestartNodeCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	RestartNode(*RestartNodeInput) (*RestartNodeOutput, error)
	RestartNodeWithContext(volcengine.Context, *RestartNodeInput, ...request.Option) (*RestartNodeOutput, error)
	RestartNodeRequest(*RestartNodeInput) (*request.Request, *RestartNodeOutput)

	TagResourcesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	TagResourcesCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	TagResourcesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	TagResources(*TagResourcesInput) (*TagResourcesOutput, error)
	TagResourcesWithContext(volcengine.Context, *TagResourcesInput, ...request.Option) (*TagResourcesOutput, error)
	TagResourcesRequest(*TagResourcesInput) (*request.Request, *TagResourcesOutput)

	UntagResourcesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	UntagResourcesCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	UntagResourcesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	UntagResources(*UntagResourcesInput) (*UntagResourcesOutput, error)
	UntagResourcesWithContext(volcengine.Context, *UntagResourcesInput, ...request.Option) (*UntagResourcesOutput, error)
	UntagResourcesRequest(*UntagResourcesInput) (*request.Request, *UntagResourcesOutput)
}

var _ ESCLOUDAPI = (*ESCLOUD)(nil)

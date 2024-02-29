// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package kafkaiface provides an interface to enable mocking the KAFKA service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package kafka

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
)

// KAFKAAPI provides an interface to enable mocking the
// kafka.KAFKA service client's API operation,
//
//    // volcengine sdk func uses an SDK service client to make a request to
//    // KAFKA.
//    func myFunc(svc KAFKAAPI) bool {
//        // Make svc.AddTagsToResource request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := kafka.New(sess)
//
//        myFunc(svc)
//    }
//
type KAFKAAPI interface {
	AddTagsToResourceCommon(*map[string]interface{}) (*map[string]interface{}, error)
	AddTagsToResourceCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	AddTagsToResourceCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	AddTagsToResource(*AddTagsToResourceInput) (*AddTagsToResourceOutput, error)
	AddTagsToResourceWithContext(volcengine.Context, *AddTagsToResourceInput, ...request.Option) (*AddTagsToResourceOutput, error)
	AddTagsToResourceRequest(*AddTagsToResourceInput) (*request.Request, *AddTagsToResourceOutput)

	AssociateAllowListCommon(*map[string]interface{}) (*map[string]interface{}, error)
	AssociateAllowListCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	AssociateAllowListCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	AssociateAllowList(*AssociateAllowListInput) (*AssociateAllowListOutput, error)
	AssociateAllowListWithContext(volcengine.Context, *AssociateAllowListInput, ...request.Option) (*AssociateAllowListOutput, error)
	AssociateAllowListRequest(*AssociateAllowListInput) (*request.Request, *AssociateAllowListOutput)

	CreateAclCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CreateAclCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateAclCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateAcl(*CreateAclInput) (*CreateAclOutput, error)
	CreateAclWithContext(volcengine.Context, *CreateAclInput, ...request.Option) (*CreateAclOutput, error)
	CreateAclRequest(*CreateAclInput) (*request.Request, *CreateAclOutput)

	CreateAllowListCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CreateAllowListCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateAllowListCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateAllowList(*CreateAllowListInput) (*CreateAllowListOutput, error)
	CreateAllowListWithContext(volcengine.Context, *CreateAllowListInput, ...request.Option) (*CreateAllowListOutput, error)
	CreateAllowListRequest(*CreateAllowListInput) (*request.Request, *CreateAllowListOutput)

	CreateGroupCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CreateGroupCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateGroupCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateGroup(*CreateGroupInput) (*CreateGroupOutput, error)
	CreateGroupWithContext(volcengine.Context, *CreateGroupInput, ...request.Option) (*CreateGroupOutput, error)
	CreateGroupRequest(*CreateGroupInput) (*request.Request, *CreateGroupOutput)

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

	CreateTopicCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CreateTopicCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateTopicCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateTopic(*CreateTopicInput) (*CreateTopicOutput, error)
	CreateTopicWithContext(volcengine.Context, *CreateTopicInput, ...request.Option) (*CreateTopicOutput, error)
	CreateTopicRequest(*CreateTopicInput) (*request.Request, *CreateTopicOutput)

	CreateUserCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CreateUserCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateUserCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateUser(*CreateUserInput) (*CreateUserOutput, error)
	CreateUserWithContext(volcengine.Context, *CreateUserInput, ...request.Option) (*CreateUserOutput, error)
	CreateUserRequest(*CreateUserInput) (*request.Request, *CreateUserOutput)

	DeleteAclCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteAclCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteAclCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteAcl(*DeleteAclInput) (*DeleteAclOutput, error)
	DeleteAclWithContext(volcengine.Context, *DeleteAclInput, ...request.Option) (*DeleteAclOutput, error)
	DeleteAclRequest(*DeleteAclInput) (*request.Request, *DeleteAclOutput)

	DeleteAllowListCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteAllowListCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteAllowListCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteAllowList(*DeleteAllowListInput) (*DeleteAllowListOutput, error)
	DeleteAllowListWithContext(volcengine.Context, *DeleteAllowListInput, ...request.Option) (*DeleteAllowListOutput, error)
	DeleteAllowListRequest(*DeleteAllowListInput) (*request.Request, *DeleteAllowListOutput)

	DeleteGroupCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteGroupCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteGroupCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteGroup(*DeleteGroupInput) (*DeleteGroupOutput, error)
	DeleteGroupWithContext(volcengine.Context, *DeleteGroupInput, ...request.Option) (*DeleteGroupOutput, error)
	DeleteGroupRequest(*DeleteGroupInput) (*request.Request, *DeleteGroupOutput)

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

	DeleteTopicCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteTopicCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteTopicCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteTopic(*DeleteTopicInput) (*DeleteTopicOutput, error)
	DeleteTopicWithContext(volcengine.Context, *DeleteTopicInput, ...request.Option) (*DeleteTopicOutput, error)
	DeleteTopicRequest(*DeleteTopicInput) (*request.Request, *DeleteTopicOutput)

	DeleteUserCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteUserCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteUserCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteUser(*DeleteUserInput) (*DeleteUserOutput, error)
	DeleteUserWithContext(volcengine.Context, *DeleteUserInput, ...request.Option) (*DeleteUserOutput, error)
	DeleteUserRequest(*DeleteUserInput) (*request.Request, *DeleteUserOutput)

	DescribeAclsCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeAclsCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeAclsCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeAcls(*DescribeAclsInput) (*DescribeAclsOutput, error)
	DescribeAclsWithContext(volcengine.Context, *DescribeAclsInput, ...request.Option) (*DescribeAclsOutput, error)
	DescribeAclsRequest(*DescribeAclsInput) (*request.Request, *DescribeAclsOutput)

	DescribeAllowListDetailCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeAllowListDetailCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeAllowListDetailCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeAllowListDetail(*DescribeAllowListDetailInput) (*DescribeAllowListDetailOutput, error)
	DescribeAllowListDetailWithContext(volcengine.Context, *DescribeAllowListDetailInput, ...request.Option) (*DescribeAllowListDetailOutput, error)
	DescribeAllowListDetailRequest(*DescribeAllowListDetailInput) (*request.Request, *DescribeAllowListDetailOutput)

	DescribeAllowListsCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeAllowListsCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeAllowListsCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeAllowLists(*DescribeAllowListsInput) (*DescribeAllowListsOutput, error)
	DescribeAllowListsWithContext(volcengine.Context, *DescribeAllowListsInput, ...request.Option) (*DescribeAllowListsOutput, error)
	DescribeAllowListsRequest(*DescribeAllowListsInput) (*request.Request, *DescribeAllowListsOutput)

	DescribeAvailabilityZonesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeAvailabilityZonesCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeAvailabilityZonesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeAvailabilityZones(*DescribeAvailabilityZonesInput) (*DescribeAvailabilityZonesOutput, error)
	DescribeAvailabilityZonesWithContext(volcengine.Context, *DescribeAvailabilityZonesInput, ...request.Option) (*DescribeAvailabilityZonesOutput, error)
	DescribeAvailabilityZonesRequest(*DescribeAvailabilityZonesInput) (*request.Request, *DescribeAvailabilityZonesOutput)

	DescribeConsumedPartitionsCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeConsumedPartitionsCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeConsumedPartitionsCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeConsumedPartitions(*DescribeConsumedPartitionsInput) (*DescribeConsumedPartitionsOutput, error)
	DescribeConsumedPartitionsWithContext(volcengine.Context, *DescribeConsumedPartitionsInput, ...request.Option) (*DescribeConsumedPartitionsOutput, error)
	DescribeConsumedPartitionsRequest(*DescribeConsumedPartitionsInput) (*request.Request, *DescribeConsumedPartitionsOutput)

	DescribeConsumedTopicsCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeConsumedTopicsCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeConsumedTopicsCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeConsumedTopics(*DescribeConsumedTopicsInput) (*DescribeConsumedTopicsOutput, error)
	DescribeConsumedTopicsWithContext(volcengine.Context, *DescribeConsumedTopicsInput, ...request.Option) (*DescribeConsumedTopicsOutput, error)
	DescribeConsumedTopicsRequest(*DescribeConsumedTopicsInput) (*request.Request, *DescribeConsumedTopicsOutput)

	DescribeGroupsCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeGroupsCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeGroupsCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeGroups(*DescribeGroupsInput) (*DescribeGroupsOutput, error)
	DescribeGroupsWithContext(volcengine.Context, *DescribeGroupsInput, ...request.Option) (*DescribeGroupsOutput, error)
	DescribeGroupsRequest(*DescribeGroupsInput) (*request.Request, *DescribeGroupsOutput)

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

	DescribeTopicAccessPoliciesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeTopicAccessPoliciesCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeTopicAccessPoliciesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeTopicAccessPolicies(*DescribeTopicAccessPoliciesInput) (*DescribeTopicAccessPoliciesOutput, error)
	DescribeTopicAccessPoliciesWithContext(volcengine.Context, *DescribeTopicAccessPoliciesInput, ...request.Option) (*DescribeTopicAccessPoliciesOutput, error)
	DescribeTopicAccessPoliciesRequest(*DescribeTopicAccessPoliciesInput) (*request.Request, *DescribeTopicAccessPoliciesOutput)

	DescribeTopicParametersCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeTopicParametersCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeTopicParametersCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeTopicParameters(*DescribeTopicParametersInput) (*DescribeTopicParametersOutput, error)
	DescribeTopicParametersWithContext(volcengine.Context, *DescribeTopicParametersInput, ...request.Option) (*DescribeTopicParametersOutput, error)
	DescribeTopicParametersRequest(*DescribeTopicParametersInput) (*request.Request, *DescribeTopicParametersOutput)

	DescribeTopicPartitionsCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeTopicPartitionsCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeTopicPartitionsCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeTopicPartitions(*DescribeTopicPartitionsInput) (*DescribeTopicPartitionsOutput, error)
	DescribeTopicPartitionsWithContext(volcengine.Context, *DescribeTopicPartitionsInput, ...request.Option) (*DescribeTopicPartitionsOutput, error)
	DescribeTopicPartitionsRequest(*DescribeTopicPartitionsInput) (*request.Request, *DescribeTopicPartitionsOutput)

	DescribeTopicsCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeTopicsCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeTopicsCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeTopics(*DescribeTopicsInput) (*DescribeTopicsOutput, error)
	DescribeTopicsWithContext(volcengine.Context, *DescribeTopicsInput, ...request.Option) (*DescribeTopicsOutput, error)
	DescribeTopicsRequest(*DescribeTopicsInput) (*request.Request, *DescribeTopicsOutput)

	DescribeUsersCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeUsersCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeUsersCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeUsers(*DescribeUsersInput) (*DescribeUsersOutput, error)
	DescribeUsersWithContext(volcengine.Context, *DescribeUsersInput, ...request.Option) (*DescribeUsersOutput, error)
	DescribeUsersRequest(*DescribeUsersInput) (*request.Request, *DescribeUsersOutput)

	DisassociateAllowListCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DisassociateAllowListCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DisassociateAllowListCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DisassociateAllowList(*DisassociateAllowListInput) (*DisassociateAllowListOutput, error)
	DisassociateAllowListWithContext(volcengine.Context, *DisassociateAllowListInput, ...request.Option) (*DisassociateAllowListOutput, error)
	DisassociateAllowListRequest(*DisassociateAllowListInput) (*request.Request, *DisassociateAllowListOutput)

	ModifyAllowListCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ModifyAllowListCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifyAllowListCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifyAllowList(*ModifyAllowListInput) (*ModifyAllowListOutput, error)
	ModifyAllowListWithContext(volcengine.Context, *ModifyAllowListInput, ...request.Option) (*ModifyAllowListOutput, error)
	ModifyAllowListRequest(*ModifyAllowListInput) (*request.Request, *ModifyAllowListOutput)

	ModifyGroupCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ModifyGroupCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifyGroupCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifyGroup(*ModifyGroupInput) (*ModifyGroupOutput, error)
	ModifyGroupWithContext(volcengine.Context, *ModifyGroupInput, ...request.Option) (*ModifyGroupOutput, error)
	ModifyGroupRequest(*ModifyGroupInput) (*request.Request, *ModifyGroupOutput)

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

	ModifyInstanceParametersCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ModifyInstanceParametersCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifyInstanceParametersCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifyInstanceParameters(*ModifyInstanceParametersInput) (*ModifyInstanceParametersOutput, error)
	ModifyInstanceParametersWithContext(volcengine.Context, *ModifyInstanceParametersInput, ...request.Option) (*ModifyInstanceParametersOutput, error)
	ModifyInstanceParametersRequest(*ModifyInstanceParametersInput) (*request.Request, *ModifyInstanceParametersOutput)

	ModifyInstanceSpecCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ModifyInstanceSpecCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifyInstanceSpecCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifyInstanceSpec(*ModifyInstanceSpecInput) (*ModifyInstanceSpecOutput, error)
	ModifyInstanceSpecWithContext(volcengine.Context, *ModifyInstanceSpecInput, ...request.Option) (*ModifyInstanceSpecOutput, error)
	ModifyInstanceSpecRequest(*ModifyInstanceSpecInput) (*request.Request, *ModifyInstanceSpecOutput)

	ModifyTopicAccessPoliciesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ModifyTopicAccessPoliciesCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifyTopicAccessPoliciesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifyTopicAccessPolicies(*ModifyTopicAccessPoliciesInput) (*ModifyTopicAccessPoliciesOutput, error)
	ModifyTopicAccessPoliciesWithContext(volcengine.Context, *ModifyTopicAccessPoliciesInput, ...request.Option) (*ModifyTopicAccessPoliciesOutput, error)
	ModifyTopicAccessPoliciesRequest(*ModifyTopicAccessPoliciesInput) (*request.Request, *ModifyTopicAccessPoliciesOutput)

	ModifyTopicAttributesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ModifyTopicAttributesCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifyTopicAttributesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifyTopicAttributes(*ModifyTopicAttributesInput) (*ModifyTopicAttributesOutput, error)
	ModifyTopicAttributesWithContext(volcengine.Context, *ModifyTopicAttributesInput, ...request.Option) (*ModifyTopicAttributesOutput, error)
	ModifyTopicAttributesRequest(*ModifyTopicAttributesInput) (*request.Request, *ModifyTopicAttributesOutput)

	ModifyTopicParametersCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ModifyTopicParametersCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifyTopicParametersCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifyTopicParameters(*ModifyTopicParametersInput) (*ModifyTopicParametersOutput, error)
	ModifyTopicParametersWithContext(volcengine.Context, *ModifyTopicParametersInput, ...request.Option) (*ModifyTopicParametersOutput, error)
	ModifyTopicParametersRequest(*ModifyTopicParametersInput) (*request.Request, *ModifyTopicParametersOutput)

	ModifyUserAuthorityCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ModifyUserAuthorityCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifyUserAuthorityCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifyUserAuthority(*ModifyUserAuthorityInput) (*ModifyUserAuthorityOutput, error)
	ModifyUserAuthorityWithContext(volcengine.Context, *ModifyUserAuthorityInput, ...request.Option) (*ModifyUserAuthorityOutput, error)
	ModifyUserAuthorityRequest(*ModifyUserAuthorityInput) (*request.Request, *ModifyUserAuthorityOutput)

	RemoveTagsFromResourceCommon(*map[string]interface{}) (*map[string]interface{}, error)
	RemoveTagsFromResourceCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	RemoveTagsFromResourceCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	RemoveTagsFromResource(*RemoveTagsFromResourceInput) (*RemoveTagsFromResourceOutput, error)
	RemoveTagsFromResourceWithContext(volcengine.Context, *RemoveTagsFromResourceInput, ...request.Option) (*RemoveTagsFromResourceOutput, error)
	RemoveTagsFromResourceRequest(*RemoveTagsFromResourceInput) (*request.Request, *RemoveTagsFromResourceOutput)

	ResetConsumedOffsetsCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ResetConsumedOffsetsCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ResetConsumedOffsetsCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ResetConsumedOffsets(*ResetConsumedOffsetsInput) (*ResetConsumedOffsetsOutput, error)
	ResetConsumedOffsetsWithContext(volcengine.Context, *ResetConsumedOffsetsInput, ...request.Option) (*ResetConsumedOffsetsOutput, error)
	ResetConsumedOffsetsRequest(*ResetConsumedOffsetsInput) (*request.Request, *ResetConsumedOffsetsOutput)
}

var _ KAFKAAPI = (*KAFKA)(nil)

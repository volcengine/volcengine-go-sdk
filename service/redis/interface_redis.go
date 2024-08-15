// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package redisiface provides an interface to enable mocking the REDIS service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package redis

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
)

// REDISAPI provides an interface to enable mocking the
// redis.REDIS service client's API operation,
//
//    // volcengine sdk func uses an SDK service client to make a request to
//    // REDIS.
//    func myFunc(svc REDISAPI) bool {
//        // Make svc.AddTagsToResource request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := redis.New(sess)
//
//        myFunc(svc)
//    }
//
type REDISAPI interface {
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

	CreateAllowListCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CreateAllowListCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateAllowListCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateAllowList(*CreateAllowListInput) (*CreateAllowListOutput, error)
	CreateAllowListWithContext(volcengine.Context, *CreateAllowListInput, ...request.Option) (*CreateAllowListOutput, error)
	CreateAllowListRequest(*CreateAllowListInput) (*request.Request, *CreateAllowListOutput)

	CreateBackupCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CreateBackupCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateBackupCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateBackup(*CreateBackupInput) (*CreateBackupOutput, error)
	CreateBackupWithContext(volcengine.Context, *CreateBackupInput, ...request.Option) (*CreateBackupOutput, error)
	CreateBackupRequest(*CreateBackupInput) (*request.Request, *CreateBackupOutput)

	CreateDBAccountCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CreateDBAccountCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateDBAccountCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateDBAccount(*CreateDBAccountInput) (*CreateDBAccountOutput, error)
	CreateDBAccountWithContext(volcengine.Context, *CreateDBAccountInput, ...request.Option) (*CreateDBAccountOutput, error)
	CreateDBAccountRequest(*CreateDBAccountInput) (*request.Request, *CreateDBAccountOutput)

	CreateDBEndpointPublicAddressCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CreateDBEndpointPublicAddressCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateDBEndpointPublicAddressCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateDBEndpointPublicAddress(*CreateDBEndpointPublicAddressInput) (*CreateDBEndpointPublicAddressOutput, error)
	CreateDBEndpointPublicAddressWithContext(volcengine.Context, *CreateDBEndpointPublicAddressInput, ...request.Option) (*CreateDBEndpointPublicAddressOutput, error)
	CreateDBEndpointPublicAddressRequest(*CreateDBEndpointPublicAddressInput) (*request.Request, *CreateDBEndpointPublicAddressOutput)

	CreateDBInstanceCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CreateDBInstanceCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateDBInstanceCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateDBInstance(*CreateDBInstanceInput) (*CreateDBInstanceOutput, error)
	CreateDBInstanceWithContext(volcengine.Context, *CreateDBInstanceInput, ...request.Option) (*CreateDBInstanceOutput, error)
	CreateDBInstanceRequest(*CreateDBInstanceInput) (*request.Request, *CreateDBInstanceOutput)

	CreateEnterpriseDBInstanceCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CreateEnterpriseDBInstanceCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateEnterpriseDBInstanceCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateEnterpriseDBInstance(*CreateEnterpriseDBInstanceInput) (*CreateEnterpriseDBInstanceOutput, error)
	CreateEnterpriseDBInstanceWithContext(volcengine.Context, *CreateEnterpriseDBInstanceInput, ...request.Option) (*CreateEnterpriseDBInstanceOutput, error)
	CreateEnterpriseDBInstanceRequest(*CreateEnterpriseDBInstanceInput) (*request.Request, *CreateEnterpriseDBInstanceOutput)

	DecreaseDBInstanceNodeNumberCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DecreaseDBInstanceNodeNumberCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DecreaseDBInstanceNodeNumberCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DecreaseDBInstanceNodeNumber(*DecreaseDBInstanceNodeNumberInput) (*DecreaseDBInstanceNodeNumberOutput, error)
	DecreaseDBInstanceNodeNumberWithContext(volcengine.Context, *DecreaseDBInstanceNodeNumberInput, ...request.Option) (*DecreaseDBInstanceNodeNumberOutput, error)
	DecreaseDBInstanceNodeNumberRequest(*DecreaseDBInstanceNodeNumberInput) (*request.Request, *DecreaseDBInstanceNodeNumberOutput)

	DeleteAllowListCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteAllowListCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteAllowListCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteAllowList(*DeleteAllowListInput) (*DeleteAllowListOutput, error)
	DeleteAllowListWithContext(volcengine.Context, *DeleteAllowListInput, ...request.Option) (*DeleteAllowListOutput, error)
	DeleteAllowListRequest(*DeleteAllowListInput) (*request.Request, *DeleteAllowListOutput)

	DeleteDBAccountCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteDBAccountCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteDBAccountCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteDBAccount(*DeleteDBAccountInput) (*DeleteDBAccountOutput, error)
	DeleteDBAccountWithContext(volcengine.Context, *DeleteDBAccountInput, ...request.Option) (*DeleteDBAccountOutput, error)
	DeleteDBAccountRequest(*DeleteDBAccountInput) (*request.Request, *DeleteDBAccountOutput)

	DeleteDBEndpointPublicAddressCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteDBEndpointPublicAddressCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteDBEndpointPublicAddressCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteDBEndpointPublicAddress(*DeleteDBEndpointPublicAddressInput) (*DeleteDBEndpointPublicAddressOutput, error)
	DeleteDBEndpointPublicAddressWithContext(volcengine.Context, *DeleteDBEndpointPublicAddressInput, ...request.Option) (*DeleteDBEndpointPublicAddressOutput, error)
	DeleteDBEndpointPublicAddressRequest(*DeleteDBEndpointPublicAddressInput) (*request.Request, *DeleteDBEndpointPublicAddressOutput)

	DeleteDBInstanceCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteDBInstanceCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteDBInstanceCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteDBInstance(*DeleteDBInstanceInput) (*DeleteDBInstanceOutput, error)
	DeleteDBInstanceWithContext(volcengine.Context, *DeleteDBInstanceInput, ...request.Option) (*DeleteDBInstanceOutput, error)
	DeleteDBInstanceRequest(*DeleteDBInstanceInput) (*request.Request, *DeleteDBInstanceOutput)

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

	DescribeBackupPlanCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeBackupPlanCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeBackupPlanCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeBackupPlan(*DescribeBackupPlanInput) (*DescribeBackupPlanOutput, error)
	DescribeBackupPlanWithContext(volcengine.Context, *DescribeBackupPlanInput, ...request.Option) (*DescribeBackupPlanOutput, error)
	DescribeBackupPlanRequest(*DescribeBackupPlanInput) (*request.Request, *DescribeBackupPlanOutput)

	DescribeBackupPointDownloadUrlsCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeBackupPointDownloadUrlsCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeBackupPointDownloadUrlsCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeBackupPointDownloadUrls(*DescribeBackupPointDownloadUrlsInput) (*DescribeBackupPointDownloadUrlsOutput, error)
	DescribeBackupPointDownloadUrlsWithContext(volcengine.Context, *DescribeBackupPointDownloadUrlsInput, ...request.Option) (*DescribeBackupPointDownloadUrlsOutput, error)
	DescribeBackupPointDownloadUrlsRequest(*DescribeBackupPointDownloadUrlsInput) (*request.Request, *DescribeBackupPointDownloadUrlsOutput)

	DescribeBackupsCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeBackupsCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeBackupsCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeBackups(*DescribeBackupsInput) (*DescribeBackupsOutput, error)
	DescribeBackupsWithContext(volcengine.Context, *DescribeBackupsInput, ...request.Option) (*DescribeBackupsOutput, error)
	DescribeBackupsRequest(*DescribeBackupsInput) (*request.Request, *DescribeBackupsOutput)

	DescribeBigKeysCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeBigKeysCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeBigKeysCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeBigKeys(*DescribeBigKeysInput) (*DescribeBigKeysOutput, error)
	DescribeBigKeysWithContext(volcengine.Context, *DescribeBigKeysInput, ...request.Option) (*DescribeBigKeysOutput, error)
	DescribeBigKeysRequest(*DescribeBigKeysInput) (*request.Request, *DescribeBigKeysOutput)

	DescribeDBInstanceDetailCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeDBInstanceDetailCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeDBInstanceDetailCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeDBInstanceDetail(*DescribeDBInstanceDetailInput) (*DescribeDBInstanceDetailOutput, error)
	DescribeDBInstanceDetailWithContext(volcengine.Context, *DescribeDBInstanceDetailInput, ...request.Option) (*DescribeDBInstanceDetailOutput, error)
	DescribeDBInstanceDetailRequest(*DescribeDBInstanceDetailInput) (*request.Request, *DescribeDBInstanceDetailOutput)

	DescribeDBInstanceParamsCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeDBInstanceParamsCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeDBInstanceParamsCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeDBInstanceParams(*DescribeDBInstanceParamsInput) (*DescribeDBInstanceParamsOutput, error)
	DescribeDBInstanceParamsWithContext(volcengine.Context, *DescribeDBInstanceParamsInput, ...request.Option) (*DescribeDBInstanceParamsOutput, error)
	DescribeDBInstanceParamsRequest(*DescribeDBInstanceParamsInput) (*request.Request, *DescribeDBInstanceParamsOutput)

	DescribeDBInstanceShardsCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeDBInstanceShardsCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeDBInstanceShardsCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeDBInstanceShards(*DescribeDBInstanceShardsInput) (*DescribeDBInstanceShardsOutput, error)
	DescribeDBInstanceShardsWithContext(volcengine.Context, *DescribeDBInstanceShardsInput, ...request.Option) (*DescribeDBInstanceShardsOutput, error)
	DescribeDBInstanceShardsRequest(*DescribeDBInstanceShardsInput) (*request.Request, *DescribeDBInstanceShardsOutput)

	DescribeDBInstancesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeDBInstancesCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeDBInstancesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeDBInstances(*DescribeDBInstancesInput) (*DescribeDBInstancesOutput, error)
	DescribeDBInstancesWithContext(volcengine.Context, *DescribeDBInstancesInput, ...request.Option) (*DescribeDBInstancesOutput, error)
	DescribeDBInstancesRequest(*DescribeDBInstancesInput) (*request.Request, *DescribeDBInstancesOutput)

	DescribeEnterpriseDBInstanceDetailCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeEnterpriseDBInstanceDetailCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeEnterpriseDBInstanceDetailCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeEnterpriseDBInstanceDetail(*DescribeEnterpriseDBInstanceDetailInput) (*DescribeEnterpriseDBInstanceDetailOutput, error)
	DescribeEnterpriseDBInstanceDetailWithContext(volcengine.Context, *DescribeEnterpriseDBInstanceDetailInput, ...request.Option) (*DescribeEnterpriseDBInstanceDetailOutput, error)
	DescribeEnterpriseDBInstanceDetailRequest(*DescribeEnterpriseDBInstanceDetailInput) (*request.Request, *DescribeEnterpriseDBInstanceDetailOutput)

	DescribeNodeIdsCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeNodeIdsCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeNodeIdsCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeNodeIds(*DescribeNodeIdsInput) (*DescribeNodeIdsOutput, error)
	DescribeNodeIdsWithContext(volcengine.Context, *DescribeNodeIdsInput, ...request.Option) (*DescribeNodeIdsOutput, error)
	DescribeNodeIdsRequest(*DescribeNodeIdsInput) (*request.Request, *DescribeNodeIdsOutput)

	DescribePitrTimeWindowCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribePitrTimeWindowCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribePitrTimeWindowCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribePitrTimeWindow(*DescribePitrTimeWindowInput) (*DescribePitrTimeWindowOutput, error)
	DescribePitrTimeWindowWithContext(volcengine.Context, *DescribePitrTimeWindowInput, ...request.Option) (*DescribePitrTimeWindowOutput, error)
	DescribePitrTimeWindowRequest(*DescribePitrTimeWindowInput) (*request.Request, *DescribePitrTimeWindowOutput)

	DescribeRegionsCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeRegionsCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeRegionsCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeRegions(*DescribeRegionsInput) (*DescribeRegionsOutput, error)
	DescribeRegionsWithContext(volcengine.Context, *DescribeRegionsInput, ...request.Option) (*DescribeRegionsOutput, error)
	DescribeRegionsRequest(*DescribeRegionsInput) (*request.Request, *DescribeRegionsOutput)

	DescribeSlowLogsCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeSlowLogsCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeSlowLogsCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeSlowLogs(*DescribeSlowLogsInput) (*DescribeSlowLogsOutput, error)
	DescribeSlowLogsWithContext(volcengine.Context, *DescribeSlowLogsInput, ...request.Option) (*DescribeSlowLogsOutput, error)
	DescribeSlowLogsRequest(*DescribeSlowLogsInput) (*request.Request, *DescribeSlowLogsOutput)

	DescribeTagsByResourceCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeTagsByResourceCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeTagsByResourceCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeTagsByResource(*DescribeTagsByResourceInput) (*DescribeTagsByResourceOutput, error)
	DescribeTagsByResourceWithContext(volcengine.Context, *DescribeTagsByResourceInput, ...request.Option) (*DescribeTagsByResourceOutput, error)
	DescribeTagsByResourceRequest(*DescribeTagsByResourceInput) (*request.Request, *DescribeTagsByResourceOutput)

	DescribeZonesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeZonesCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeZonesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeZones(*DescribeZonesInput) (*DescribeZonesOutput, error)
	DescribeZonesWithContext(volcengine.Context, *DescribeZonesInput, ...request.Option) (*DescribeZonesOutput, error)
	DescribeZonesRequest(*DescribeZonesInput) (*request.Request, *DescribeZonesOutput)

	DisassociateAllowListCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DisassociateAllowListCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DisassociateAllowListCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DisassociateAllowList(*DisassociateAllowListInput) (*DisassociateAllowListOutput, error)
	DisassociateAllowListWithContext(volcengine.Context, *DisassociateAllowListInput, ...request.Option) (*DisassociateAllowListOutput, error)
	DisassociateAllowListRequest(*DisassociateAllowListInput) (*request.Request, *DisassociateAllowListOutput)

	EnableShardedClusterCommon(*map[string]interface{}) (*map[string]interface{}, error)
	EnableShardedClusterCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	EnableShardedClusterCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	EnableShardedCluster(*EnableShardedClusterInput) (*EnableShardedClusterOutput, error)
	EnableShardedClusterWithContext(volcengine.Context, *EnableShardedClusterInput, ...request.Option) (*EnableShardedClusterOutput, error)
	EnableShardedClusterRequest(*EnableShardedClusterInput) (*request.Request, *EnableShardedClusterOutput)

	FlushDBInstanceCommon(*map[string]interface{}) (*map[string]interface{}, error)
	FlushDBInstanceCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	FlushDBInstanceCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	FlushDBInstance(*FlushDBInstanceInput) (*FlushDBInstanceOutput, error)
	FlushDBInstanceWithContext(volcengine.Context, *FlushDBInstanceInput, ...request.Option) (*FlushDBInstanceOutput, error)
	FlushDBInstanceRequest(*FlushDBInstanceInput) (*request.Request, *FlushDBInstanceOutput)

	IncreaseDBInstanceNodeNumberCommon(*map[string]interface{}) (*map[string]interface{}, error)
	IncreaseDBInstanceNodeNumberCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	IncreaseDBInstanceNodeNumberCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	IncreaseDBInstanceNodeNumber(*IncreaseDBInstanceNodeNumberInput) (*IncreaseDBInstanceNodeNumberOutput, error)
	IncreaseDBInstanceNodeNumberWithContext(volcengine.Context, *IncreaseDBInstanceNodeNumberInput, ...request.Option) (*IncreaseDBInstanceNodeNumberOutput, error)
	IncreaseDBInstanceNodeNumberRequest(*IncreaseDBInstanceNodeNumberInput) (*request.Request, *IncreaseDBInstanceNodeNumberOutput)

	ListDBAccountCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ListDBAccountCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ListDBAccountCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ListDBAccount(*ListDBAccountInput) (*ListDBAccountOutput, error)
	ListDBAccountWithContext(volcengine.Context, *ListDBAccountInput, ...request.Option) (*ListDBAccountOutput, error)
	ListDBAccountRequest(*ListDBAccountInput) (*request.Request, *ListDBAccountOutput)

	ModifyAllowListCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ModifyAllowListCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifyAllowListCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifyAllowList(*ModifyAllowListInput) (*ModifyAllowListOutput, error)
	ModifyAllowListWithContext(volcengine.Context, *ModifyAllowListInput, ...request.Option) (*ModifyAllowListOutput, error)
	ModifyAllowListRequest(*ModifyAllowListInput) (*request.Request, *ModifyAllowListOutput)

	ModifyBackupPlanCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ModifyBackupPlanCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifyBackupPlanCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifyBackupPlan(*ModifyBackupPlanInput) (*ModifyBackupPlanOutput, error)
	ModifyBackupPlanWithContext(volcengine.Context, *ModifyBackupPlanInput, ...request.Option) (*ModifyBackupPlanOutput, error)
	ModifyBackupPlanRequest(*ModifyBackupPlanInput) (*request.Request, *ModifyBackupPlanOutput)

	ModifyDBAccountCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ModifyDBAccountCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifyDBAccountCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifyDBAccount(*ModifyDBAccountInput) (*ModifyDBAccountOutput, error)
	ModifyDBAccountWithContext(volcengine.Context, *ModifyDBAccountInput, ...request.Option) (*ModifyDBAccountOutput, error)
	ModifyDBAccountRequest(*ModifyDBAccountInput) (*request.Request, *ModifyDBAccountOutput)

	ModifyDBInstanceAZConfigureCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ModifyDBInstanceAZConfigureCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifyDBInstanceAZConfigureCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifyDBInstanceAZConfigure(*ModifyDBInstanceAZConfigureInput) (*ModifyDBInstanceAZConfigureOutput, error)
	ModifyDBInstanceAZConfigureWithContext(volcengine.Context, *ModifyDBInstanceAZConfigureInput, ...request.Option) (*ModifyDBInstanceAZConfigureOutput, error)
	ModifyDBInstanceAZConfigureRequest(*ModifyDBInstanceAZConfigureInput) (*request.Request, *ModifyDBInstanceAZConfigureOutput)

	ModifyDBInstanceAdditionalBandwidthPerShardCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ModifyDBInstanceAdditionalBandwidthPerShardCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifyDBInstanceAdditionalBandwidthPerShardCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifyDBInstanceAdditionalBandwidthPerShard(*ModifyDBInstanceAdditionalBandwidthPerShardInput) (*ModifyDBInstanceAdditionalBandwidthPerShardOutput, error)
	ModifyDBInstanceAdditionalBandwidthPerShardWithContext(volcengine.Context, *ModifyDBInstanceAdditionalBandwidthPerShardInput, ...request.Option) (*ModifyDBInstanceAdditionalBandwidthPerShardOutput, error)
	ModifyDBInstanceAdditionalBandwidthPerShardRequest(*ModifyDBInstanceAdditionalBandwidthPerShardInput) (*request.Request, *ModifyDBInstanceAdditionalBandwidthPerShardOutput)

	ModifyDBInstanceChargeTypeCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ModifyDBInstanceChargeTypeCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifyDBInstanceChargeTypeCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifyDBInstanceChargeType(*ModifyDBInstanceChargeTypeInput) (*ModifyDBInstanceChargeTypeOutput, error)
	ModifyDBInstanceChargeTypeWithContext(volcengine.Context, *ModifyDBInstanceChargeTypeInput, ...request.Option) (*ModifyDBInstanceChargeTypeOutput, error)
	ModifyDBInstanceChargeTypeRequest(*ModifyDBInstanceChargeTypeInput) (*request.Request, *ModifyDBInstanceChargeTypeOutput)

	ModifyDBInstanceDeletionProtectionPolicyCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ModifyDBInstanceDeletionProtectionPolicyCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifyDBInstanceDeletionProtectionPolicyCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifyDBInstanceDeletionProtectionPolicy(*ModifyDBInstanceDeletionProtectionPolicyInput) (*ModifyDBInstanceDeletionProtectionPolicyOutput, error)
	ModifyDBInstanceDeletionProtectionPolicyWithContext(volcengine.Context, *ModifyDBInstanceDeletionProtectionPolicyInput, ...request.Option) (*ModifyDBInstanceDeletionProtectionPolicyOutput, error)
	ModifyDBInstanceDeletionProtectionPolicyRequest(*ModifyDBInstanceDeletionProtectionPolicyInput) (*request.Request, *ModifyDBInstanceDeletionProtectionPolicyOutput)

	ModifyDBInstanceMaxConnCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ModifyDBInstanceMaxConnCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifyDBInstanceMaxConnCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifyDBInstanceMaxConn(*ModifyDBInstanceMaxConnInput) (*ModifyDBInstanceMaxConnOutput, error)
	ModifyDBInstanceMaxConnWithContext(volcengine.Context, *ModifyDBInstanceMaxConnInput, ...request.Option) (*ModifyDBInstanceMaxConnOutput, error)
	ModifyDBInstanceMaxConnRequest(*ModifyDBInstanceMaxConnInput) (*request.Request, *ModifyDBInstanceMaxConnOutput)

	ModifyDBInstanceNameCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ModifyDBInstanceNameCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifyDBInstanceNameCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifyDBInstanceName(*ModifyDBInstanceNameInput) (*ModifyDBInstanceNameOutput, error)
	ModifyDBInstanceNameWithContext(volcengine.Context, *ModifyDBInstanceNameInput, ...request.Option) (*ModifyDBInstanceNameOutput, error)
	ModifyDBInstanceNameRequest(*ModifyDBInstanceNameInput) (*request.Request, *ModifyDBInstanceNameOutput)

	ModifyDBInstanceParamsCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ModifyDBInstanceParamsCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifyDBInstanceParamsCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifyDBInstanceParams(*ModifyDBInstanceParamsInput) (*ModifyDBInstanceParamsOutput, error)
	ModifyDBInstanceParamsWithContext(volcengine.Context, *ModifyDBInstanceParamsInput, ...request.Option) (*ModifyDBInstanceParamsOutput, error)
	ModifyDBInstanceParamsRequest(*ModifyDBInstanceParamsInput) (*request.Request, *ModifyDBInstanceParamsOutput)

	ModifyDBInstanceShardCapacityCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ModifyDBInstanceShardCapacityCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifyDBInstanceShardCapacityCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifyDBInstanceShardCapacity(*ModifyDBInstanceShardCapacityInput) (*ModifyDBInstanceShardCapacityOutput, error)
	ModifyDBInstanceShardCapacityWithContext(volcengine.Context, *ModifyDBInstanceShardCapacityInput, ...request.Option) (*ModifyDBInstanceShardCapacityOutput, error)
	ModifyDBInstanceShardCapacityRequest(*ModifyDBInstanceShardCapacityInput) (*request.Request, *ModifyDBInstanceShardCapacityOutput)

	ModifyDBInstanceShardNumberCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ModifyDBInstanceShardNumberCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifyDBInstanceShardNumberCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifyDBInstanceShardNumber(*ModifyDBInstanceShardNumberInput) (*ModifyDBInstanceShardNumberOutput, error)
	ModifyDBInstanceShardNumberWithContext(volcengine.Context, *ModifyDBInstanceShardNumberInput, ...request.Option) (*ModifyDBInstanceShardNumberOutput, error)
	ModifyDBInstanceShardNumberRequest(*ModifyDBInstanceShardNumberInput) (*request.Request, *ModifyDBInstanceShardNumberOutput)

	ModifyDBInstanceSubnetCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ModifyDBInstanceSubnetCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifyDBInstanceSubnetCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifyDBInstanceSubnet(*ModifyDBInstanceSubnetInput) (*ModifyDBInstanceSubnetOutput, error)
	ModifyDBInstanceSubnetWithContext(volcengine.Context, *ModifyDBInstanceSubnetInput, ...request.Option) (*ModifyDBInstanceSubnetOutput, error)
	ModifyDBInstanceSubnetRequest(*ModifyDBInstanceSubnetInput) (*request.Request, *ModifyDBInstanceSubnetOutput)

	ModifyDBInstanceVisitAddressCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ModifyDBInstanceVisitAddressCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifyDBInstanceVisitAddressCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifyDBInstanceVisitAddress(*ModifyDBInstanceVisitAddressInput) (*ModifyDBInstanceVisitAddressOutput, error)
	ModifyDBInstanceVisitAddressWithContext(volcengine.Context, *ModifyDBInstanceVisitAddressInput, ...request.Option) (*ModifyDBInstanceVisitAddressOutput, error)
	ModifyDBInstanceVisitAddressRequest(*ModifyDBInstanceVisitAddressInput) (*request.Request, *ModifyDBInstanceVisitAddressOutput)

	ModifyDBInstanceVpcAuthModeCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ModifyDBInstanceVpcAuthModeCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifyDBInstanceVpcAuthModeCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifyDBInstanceVpcAuthMode(*ModifyDBInstanceVpcAuthModeInput) (*ModifyDBInstanceVpcAuthModeOutput, error)
	ModifyDBInstanceVpcAuthModeWithContext(volcengine.Context, *ModifyDBInstanceVpcAuthModeInput, ...request.Option) (*ModifyDBInstanceVpcAuthModeOutput, error)
	ModifyDBInstanceVpcAuthModeRequest(*ModifyDBInstanceVpcAuthModeInput) (*request.Request, *ModifyDBInstanceVpcAuthModeOutput)

	RemoveTagsFromResourceCommon(*map[string]interface{}) (*map[string]interface{}, error)
	RemoveTagsFromResourceCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	RemoveTagsFromResourceCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	RemoveTagsFromResource(*RemoveTagsFromResourceInput) (*RemoveTagsFromResourceOutput, error)
	RemoveTagsFromResourceWithContext(volcengine.Context, *RemoveTagsFromResourceInput, ...request.Option) (*RemoveTagsFromResourceOutput, error)
	RemoveTagsFromResourceRequest(*RemoveTagsFromResourceInput) (*request.Request, *RemoveTagsFromResourceOutput)

	RestartDBInstanceCommon(*map[string]interface{}) (*map[string]interface{}, error)
	RestartDBInstanceCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	RestartDBInstanceCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	RestartDBInstance(*RestartDBInstanceInput) (*RestartDBInstanceOutput, error)
	RestartDBInstanceWithContext(volcengine.Context, *RestartDBInstanceInput, ...request.Option) (*RestartDBInstanceOutput, error)
	RestartDBInstanceRequest(*RestartDBInstanceInput) (*request.Request, *RestartDBInstanceOutput)

	RestartDBInstanceProxyCommon(*map[string]interface{}) (*map[string]interface{}, error)
	RestartDBInstanceProxyCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	RestartDBInstanceProxyCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	RestartDBInstanceProxy(*RestartDBInstanceProxyInput) (*RestartDBInstanceProxyOutput, error)
	RestartDBInstanceProxyWithContext(volcengine.Context, *RestartDBInstanceProxyInput, ...request.Option) (*RestartDBInstanceProxyOutput, error)
	RestartDBInstanceProxyRequest(*RestartDBInstanceProxyInput) (*request.Request, *RestartDBInstanceProxyOutput)

	RestoreDBInstanceCommon(*map[string]interface{}) (*map[string]interface{}, error)
	RestoreDBInstanceCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	RestoreDBInstanceCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	RestoreDBInstance(*RestoreDBInstanceInput) (*RestoreDBInstanceOutput, error)
	RestoreDBInstanceWithContext(volcengine.Context, *RestoreDBInstanceInput, ...request.Option) (*RestoreDBInstanceOutput, error)
	RestoreDBInstanceRequest(*RestoreDBInstanceInput) (*request.Request, *RestoreDBInstanceOutput)

	StartContinuousBackupCommon(*map[string]interface{}) (*map[string]interface{}, error)
	StartContinuousBackupCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	StartContinuousBackupCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	StartContinuousBackup(*StartContinuousBackupInput) (*StartContinuousBackupOutput, error)
	StartContinuousBackupWithContext(volcengine.Context, *StartContinuousBackupInput, ...request.Option) (*StartContinuousBackupOutput, error)
	StartContinuousBackupRequest(*StartContinuousBackupInput) (*request.Request, *StartContinuousBackupOutput)

	StopContinuousBackupCommon(*map[string]interface{}) (*map[string]interface{}, error)
	StopContinuousBackupCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	StopContinuousBackupCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	StopContinuousBackup(*StopContinuousBackupInput) (*StopContinuousBackupOutput, error)
	StopContinuousBackupWithContext(volcengine.Context, *StopContinuousBackupInput, ...request.Option) (*StopContinuousBackupOutput, error)
	StopContinuousBackupRequest(*StopContinuousBackupInput) (*request.Request, *StopContinuousBackupOutput)

	SwitchOverCommon(*map[string]interface{}) (*map[string]interface{}, error)
	SwitchOverCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	SwitchOverCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	SwitchOver(*SwitchOverInput) (*SwitchOverOutput, error)
	SwitchOverWithContext(volcengine.Context, *SwitchOverInput, ...request.Option) (*SwitchOverOutput, error)
	SwitchOverRequest(*SwitchOverInput) (*request.Request, *SwitchOverOutput)

	UpgradeAllowListVersionCommon(*map[string]interface{}) (*map[string]interface{}, error)
	UpgradeAllowListVersionCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	UpgradeAllowListVersionCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	UpgradeAllowListVersion(*UpgradeAllowListVersionInput) (*UpgradeAllowListVersionOutput, error)
	UpgradeAllowListVersionWithContext(volcengine.Context, *UpgradeAllowListVersionInput, ...request.Option) (*UpgradeAllowListVersionOutput, error)
	UpgradeAllowListVersionRequest(*UpgradeAllowListVersionInput) (*request.Request, *UpgradeAllowListVersionOutput)
}

var _ REDISAPI = (*REDIS)(nil)

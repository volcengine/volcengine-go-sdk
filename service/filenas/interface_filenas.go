// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package filenasiface provides an interface to enable mocking the FILENAS service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package filenas

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
)

// FILENASAPI provides an interface to enable mocking the
// filenas.FILENAS service client's API operation,
//
//    // volcengine sdk func uses an SDK service client to make a request to
//    // FILENAS.
//    func myFunc(svc FILENASAPI) bool {
//        // Make svc.CancelDataFlowTask request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := filenas.New(sess)
//
//        myFunc(svc)
//    }
//
type FILENASAPI interface {
	CancelDataFlowTaskCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CancelDataFlowTaskCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CancelDataFlowTaskCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CancelDataFlowTask(*CancelDataFlowTaskInput) (*CancelDataFlowTaskOutput, error)
	CancelDataFlowTaskWithContext(volcengine.Context, *CancelDataFlowTaskInput, ...request.Option) (*CancelDataFlowTaskOutput, error)
	CancelDataFlowTaskRequest(*CancelDataFlowTaskInput) (*request.Request, *CancelDataFlowTaskOutput)

	CancelDirQuotaCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CancelDirQuotaCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CancelDirQuotaCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CancelDirQuota(*CancelDirQuotaInput) (*CancelDirQuotaOutput, error)
	CancelDirQuotaWithContext(volcengine.Context, *CancelDirQuotaInput, ...request.Option) (*CancelDirQuotaOutput, error)
	CancelDirQuotaRequest(*CancelDirQuotaInput) (*request.Request, *CancelDirQuotaOutput)

	CreateDataFlowCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CreateDataFlowCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateDataFlowCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateDataFlow(*CreateDataFlowInput) (*CreateDataFlowOutput, error)
	CreateDataFlowWithContext(volcengine.Context, *CreateDataFlowInput, ...request.Option) (*CreateDataFlowOutput, error)
	CreateDataFlowRequest(*CreateDataFlowInput) (*request.Request, *CreateDataFlowOutput)

	CreateDataFlowTaskCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CreateDataFlowTaskCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateDataFlowTaskCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateDataFlowTask(*CreateDataFlowTaskInput) (*CreateDataFlowTaskOutput, error)
	CreateDataFlowTaskWithContext(volcengine.Context, *CreateDataFlowTaskInput, ...request.Option) (*CreateDataFlowTaskOutput, error)
	CreateDataFlowTaskRequest(*CreateDataFlowTaskInput) (*request.Request, *CreateDataFlowTaskOutput)

	CreateFileSystemCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CreateFileSystemCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateFileSystemCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateFileSystem(*CreateFileSystemInput) (*CreateFileSystemOutput, error)
	CreateFileSystemWithContext(volcengine.Context, *CreateFileSystemInput, ...request.Option) (*CreateFileSystemOutput, error)
	CreateFileSystemRequest(*CreateFileSystemInput) (*request.Request, *CreateFileSystemOutput)

	CreateMountPointCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CreateMountPointCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateMountPointCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateMountPoint(*CreateMountPointInput) (*CreateMountPointOutput, error)
	CreateMountPointWithContext(volcengine.Context, *CreateMountPointInput, ...request.Option) (*CreateMountPointOutput, error)
	CreateMountPointRequest(*CreateMountPointInput) (*request.Request, *CreateMountPointOutput)

	CreatePermissionGroupCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CreatePermissionGroupCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreatePermissionGroupCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreatePermissionGroup(*CreatePermissionGroupInput) (*CreatePermissionGroupOutput, error)
	CreatePermissionGroupWithContext(volcengine.Context, *CreatePermissionGroupInput, ...request.Option) (*CreatePermissionGroupOutput, error)
	CreatePermissionGroupRequest(*CreatePermissionGroupInput) (*request.Request, *CreatePermissionGroupOutput)

	CreatePreSignedUrlCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CreatePreSignedUrlCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreatePreSignedUrlCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreatePreSignedUrl(*CreatePreSignedUrlInput) (*CreatePreSignedUrlOutput, error)
	CreatePreSignedUrlWithContext(volcengine.Context, *CreatePreSignedUrlInput, ...request.Option) (*CreatePreSignedUrlOutput, error)
	CreatePreSignedUrlRequest(*CreatePreSignedUrlInput) (*request.Request, *CreatePreSignedUrlOutput)

	CreateSnapshotCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CreateSnapshotCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateSnapshotCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateSnapshot(*CreateSnapshotInput) (*CreateSnapshotOutput, error)
	CreateSnapshotWithContext(volcengine.Context, *CreateSnapshotInput, ...request.Option) (*CreateSnapshotOutput, error)
	CreateSnapshotRequest(*CreateSnapshotInput) (*request.Request, *CreateSnapshotOutput)

	DeleteDataFlowCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteDataFlowCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteDataFlowCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteDataFlow(*DeleteDataFlowInput) (*DeleteDataFlowOutput, error)
	DeleteDataFlowWithContext(volcengine.Context, *DeleteDataFlowInput, ...request.Option) (*DeleteDataFlowOutput, error)
	DeleteDataFlowRequest(*DeleteDataFlowInput) (*request.Request, *DeleteDataFlowOutput)

	DeleteDataFlowTaskCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteDataFlowTaskCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteDataFlowTaskCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteDataFlowTask(*DeleteDataFlowTaskInput) (*DeleteDataFlowTaskOutput, error)
	DeleteDataFlowTaskWithContext(volcengine.Context, *DeleteDataFlowTaskInput, ...request.Option) (*DeleteDataFlowTaskOutput, error)
	DeleteDataFlowTaskRequest(*DeleteDataFlowTaskInput) (*request.Request, *DeleteDataFlowTaskOutput)

	DeleteFileSystemCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteFileSystemCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteFileSystemCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteFileSystem(*DeleteFileSystemInput) (*DeleteFileSystemOutput, error)
	DeleteFileSystemWithContext(volcengine.Context, *DeleteFileSystemInput, ...request.Option) (*DeleteFileSystemOutput, error)
	DeleteFileSystemRequest(*DeleteFileSystemInput) (*request.Request, *DeleteFileSystemOutput)

	DeleteMountPointCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteMountPointCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteMountPointCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteMountPoint(*DeleteMountPointInput) (*DeleteMountPointOutput, error)
	DeleteMountPointWithContext(volcengine.Context, *DeleteMountPointInput, ...request.Option) (*DeleteMountPointOutput, error)
	DeleteMountPointRequest(*DeleteMountPointInput) (*request.Request, *DeleteMountPointOutput)

	DeletePermissionGroupCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DeletePermissionGroupCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeletePermissionGroupCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeletePermissionGroup(*DeletePermissionGroupInput) (*DeletePermissionGroupOutput, error)
	DeletePermissionGroupWithContext(volcengine.Context, *DeletePermissionGroupInput, ...request.Option) (*DeletePermissionGroupOutput, error)
	DeletePermissionGroupRequest(*DeletePermissionGroupInput) (*request.Request, *DeletePermissionGroupOutput)

	DeleteSnapshotCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteSnapshotCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteSnapshotCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteSnapshot(*DeleteSnapshotInput) (*DeleteSnapshotOutput, error)
	DeleteSnapshotWithContext(volcengine.Context, *DeleteSnapshotInput, ...request.Option) (*DeleteSnapshotOutput, error)
	DeleteSnapshotRequest(*DeleteSnapshotInput) (*request.Request, *DeleteSnapshotOutput)

	DescribeDataFlowTasksCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeDataFlowTasksCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeDataFlowTasksCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeDataFlowTasks(*DescribeDataFlowTasksInput) (*DescribeDataFlowTasksOutput, error)
	DescribeDataFlowTasksWithContext(volcengine.Context, *DescribeDataFlowTasksInput, ...request.Option) (*DescribeDataFlowTasksOutput, error)
	DescribeDataFlowTasksRequest(*DescribeDataFlowTasksInput) (*request.Request, *DescribeDataFlowTasksOutput)

	DescribeDataFlowsCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeDataFlowsCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeDataFlowsCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeDataFlows(*DescribeDataFlowsInput) (*DescribeDataFlowsOutput, error)
	DescribeDataFlowsWithContext(volcengine.Context, *DescribeDataFlowsInput, ...request.Option) (*DescribeDataFlowsOutput, error)
	DescribeDataFlowsRequest(*DescribeDataFlowsInput) (*request.Request, *DescribeDataFlowsOutput)

	DescribeDirQuotasCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeDirQuotasCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeDirQuotasCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeDirQuotas(*DescribeDirQuotasInput) (*DescribeDirQuotasOutput, error)
	DescribeDirQuotasWithContext(volcengine.Context, *DescribeDirQuotasInput, ...request.Option) (*DescribeDirQuotasOutput, error)
	DescribeDirQuotasRequest(*DescribeDirQuotasInput) (*request.Request, *DescribeDirQuotasOutput)

	DescribeFileSystemOverviewCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeFileSystemOverviewCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeFileSystemOverviewCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeFileSystemOverview(*DescribeFileSystemOverviewInput) (*DescribeFileSystemOverviewOutput, error)
	DescribeFileSystemOverviewWithContext(volcengine.Context, *DescribeFileSystemOverviewInput, ...request.Option) (*DescribeFileSystemOverviewOutput, error)
	DescribeFileSystemOverviewRequest(*DescribeFileSystemOverviewInput) (*request.Request, *DescribeFileSystemOverviewOutput)

	DescribeFileSystemStatisticsCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeFileSystemStatisticsCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeFileSystemStatisticsCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeFileSystemStatistics(*DescribeFileSystemStatisticsInput) (*DescribeFileSystemStatisticsOutput, error)
	DescribeFileSystemStatisticsWithContext(volcengine.Context, *DescribeFileSystemStatisticsInput, ...request.Option) (*DescribeFileSystemStatisticsOutput, error)
	DescribeFileSystemStatisticsRequest(*DescribeFileSystemStatisticsInput) (*request.Request, *DescribeFileSystemStatisticsOutput)

	DescribeFileSystemsCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeFileSystemsCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeFileSystemsCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeFileSystems(*DescribeFileSystemsInput) (*DescribeFileSystemsOutput, error)
	DescribeFileSystemsWithContext(volcengine.Context, *DescribeFileSystemsInput, ...request.Option) (*DescribeFileSystemsOutput, error)
	DescribeFileSystemsRequest(*DescribeFileSystemsInput) (*request.Request, *DescribeFileSystemsOutput)

	DescribeMountPointsCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeMountPointsCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeMountPointsCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeMountPoints(*DescribeMountPointsInput) (*DescribeMountPointsOutput, error)
	DescribeMountPointsWithContext(volcengine.Context, *DescribeMountPointsInput, ...request.Option) (*DescribeMountPointsOutput, error)
	DescribeMountPointsRequest(*DescribeMountPointsInput) (*request.Request, *DescribeMountPointsOutput)

	DescribePermissionGroupsCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribePermissionGroupsCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribePermissionGroupsCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribePermissionGroups(*DescribePermissionGroupsInput) (*DescribePermissionGroupsOutput, error)
	DescribePermissionGroupsWithContext(volcengine.Context, *DescribePermissionGroupsInput, ...request.Option) (*DescribePermissionGroupsOutput, error)
	DescribePermissionGroupsRequest(*DescribePermissionGroupsInput) (*request.Request, *DescribePermissionGroupsOutput)

	DescribePermissionRulesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribePermissionRulesCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribePermissionRulesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribePermissionRules(*DescribePermissionRulesInput) (*DescribePermissionRulesOutput, error)
	DescribePermissionRulesWithContext(volcengine.Context, *DescribePermissionRulesInput, ...request.Option) (*DescribePermissionRulesOutput, error)
	DescribePermissionRulesRequest(*DescribePermissionRulesInput) (*request.Request, *DescribePermissionRulesOutput)

	DescribeRegionsCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeRegionsCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeRegionsCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeRegions(*DescribeRegionsInput) (*DescribeRegionsOutput, error)
	DescribeRegionsWithContext(volcengine.Context, *DescribeRegionsInput, ...request.Option) (*DescribeRegionsOutput, error)
	DescribeRegionsRequest(*DescribeRegionsInput) (*request.Request, *DescribeRegionsOutput)

	DescribeSnapshotsCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeSnapshotsCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeSnapshotsCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeSnapshots(*DescribeSnapshotsInput) (*DescribeSnapshotsOutput, error)
	DescribeSnapshotsWithContext(volcengine.Context, *DescribeSnapshotsInput, ...request.Option) (*DescribeSnapshotsOutput, error)
	DescribeSnapshotsRequest(*DescribeSnapshotsInput) (*request.Request, *DescribeSnapshotsOutput)

	DescribeZonesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeZonesCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeZonesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeZones(*DescribeZonesInput) (*DescribeZonesOutput, error)
	DescribeZonesWithContext(volcengine.Context, *DescribeZonesInput, ...request.Option) (*DescribeZonesOutput, error)
	DescribeZonesRequest(*DescribeZonesInput) (*request.Request, *DescribeZonesOutput)

	ExpandFileSystemCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ExpandFileSystemCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ExpandFileSystemCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ExpandFileSystem(*ExpandFileSystemInput) (*ExpandFileSystemOutput, error)
	ExpandFileSystemWithContext(volcengine.Context, *ExpandFileSystemInput, ...request.Option) (*ExpandFileSystemOutput, error)
	ExpandFileSystemRequest(*ExpandFileSystemInput) (*request.Request, *ExpandFileSystemOutput)

	ModifyFileSystemSpecCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ModifyFileSystemSpecCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifyFileSystemSpecCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifyFileSystemSpec(*ModifyFileSystemSpecInput) (*ModifyFileSystemSpecOutput, error)
	ModifyFileSystemSpecWithContext(volcengine.Context, *ModifyFileSystemSpecInput, ...request.Option) (*ModifyFileSystemSpecOutput, error)
	ModifyFileSystemSpecRequest(*ModifyFileSystemSpecInput) (*request.Request, *ModifyFileSystemSpecOutput)

	SetDirQuotaCommon(*map[string]interface{}) (*map[string]interface{}, error)
	SetDirQuotaCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	SetDirQuotaCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	SetDirQuota(*SetDirQuotaInput) (*SetDirQuotaOutput, error)
	SetDirQuotaWithContext(volcengine.Context, *SetDirQuotaInput, ...request.Option) (*SetDirQuotaOutput, error)
	SetDirQuotaRequest(*SetDirQuotaInput) (*request.Request, *SetDirQuotaOutput)

	StartDataFlowCommon(*map[string]interface{}) (*map[string]interface{}, error)
	StartDataFlowCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	StartDataFlowCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	StartDataFlow(*StartDataFlowInput) (*StartDataFlowOutput, error)
	StartDataFlowWithContext(volcengine.Context, *StartDataFlowInput, ...request.Option) (*StartDataFlowOutput, error)
	StartDataFlowRequest(*StartDataFlowInput) (*request.Request, *StartDataFlowOutput)

	StopDataFlowCommon(*map[string]interface{}) (*map[string]interface{}, error)
	StopDataFlowCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	StopDataFlowCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	StopDataFlow(*StopDataFlowInput) (*StopDataFlowOutput, error)
	StopDataFlowWithContext(volcengine.Context, *StopDataFlowInput, ...request.Option) (*StopDataFlowOutput, error)
	StopDataFlowRequest(*StopDataFlowInput) (*request.Request, *StopDataFlowOutput)

	UpdateDataFlowCommon(*map[string]interface{}) (*map[string]interface{}, error)
	UpdateDataFlowCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	UpdateDataFlowCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	UpdateDataFlow(*UpdateDataFlowInput) (*UpdateDataFlowOutput, error)
	UpdateDataFlowWithContext(volcengine.Context, *UpdateDataFlowInput, ...request.Option) (*UpdateDataFlowOutput, error)
	UpdateDataFlowRequest(*UpdateDataFlowInput) (*request.Request, *UpdateDataFlowOutput)

	UpdateFileSystemCommon(*map[string]interface{}) (*map[string]interface{}, error)
	UpdateFileSystemCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	UpdateFileSystemCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	UpdateFileSystem(*UpdateFileSystemInput) (*UpdateFileSystemOutput, error)
	UpdateFileSystemWithContext(volcengine.Context, *UpdateFileSystemInput, ...request.Option) (*UpdateFileSystemOutput, error)
	UpdateFileSystemRequest(*UpdateFileSystemInput) (*request.Request, *UpdateFileSystemOutput)

	UpdateMountPointCommon(*map[string]interface{}) (*map[string]interface{}, error)
	UpdateMountPointCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	UpdateMountPointCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	UpdateMountPoint(*UpdateMountPointInput) (*UpdateMountPointOutput, error)
	UpdateMountPointWithContext(volcengine.Context, *UpdateMountPointInput, ...request.Option) (*UpdateMountPointOutput, error)
	UpdateMountPointRequest(*UpdateMountPointInput) (*request.Request, *UpdateMountPointOutput)

	UpdatePermissionGroupCommon(*map[string]interface{}) (*map[string]interface{}, error)
	UpdatePermissionGroupCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	UpdatePermissionGroupCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	UpdatePermissionGroup(*UpdatePermissionGroupInput) (*UpdatePermissionGroupOutput, error)
	UpdatePermissionGroupWithContext(volcengine.Context, *UpdatePermissionGroupInput, ...request.Option) (*UpdatePermissionGroupOutput, error)
	UpdatePermissionGroupRequest(*UpdatePermissionGroupInput) (*request.Request, *UpdatePermissionGroupOutput)

	UpdatePermissionRuleCommon(*map[string]interface{}) (*map[string]interface{}, error)
	UpdatePermissionRuleCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	UpdatePermissionRuleCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	UpdatePermissionRule(*UpdatePermissionRuleInput) (*UpdatePermissionRuleOutput, error)
	UpdatePermissionRuleWithContext(volcengine.Context, *UpdatePermissionRuleInput, ...request.Option) (*UpdatePermissionRuleOutput, error)
	UpdatePermissionRuleRequest(*UpdatePermissionRuleInput) (*request.Request, *UpdatePermissionRuleOutput)

	UpdateSnapshotCommon(*map[string]interface{}) (*map[string]interface{}, error)
	UpdateSnapshotCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	UpdateSnapshotCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	UpdateSnapshot(*UpdateSnapshotInput) (*UpdateSnapshotOutput, error)
	UpdateSnapshotWithContext(volcengine.Context, *UpdateSnapshotInput, ...request.Option) (*UpdateSnapshotOutput, error)
	UpdateSnapshotRequest(*UpdateSnapshotInput) (*request.Request, *UpdateSnapshotOutput)
}

var _ FILENASAPI = (*FILENAS)(nil)

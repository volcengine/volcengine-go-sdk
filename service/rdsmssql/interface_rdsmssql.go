// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package rdsmssqliface provides an interface to enable mocking the RDS_MSSQL service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package rdsmssql

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
)

// RDSMSSQLAPI provides an interface to enable mocking the
// rdsmssql.RDSMSSQL service client's API operation,
//
//    // volcengine sdk func uses an SDK service client to make a request to
//    // RDS_MSSQL.
//    func myFunc(svc RDSMSSQLAPI) bool {
//        // Make svc.CreateBackup request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := rdsmssql.New(sess)
//
//        myFunc(svc)
//    }
//
type RDSMSSQLAPI interface {
	CreateBackupCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CreateBackupCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateBackupCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateBackup(*CreateBackupInput) (*CreateBackupOutput, error)
	CreateBackupWithContext(volcengine.Context, *CreateBackupInput, ...request.Option) (*CreateBackupOutput, error)
	CreateBackupRequest(*CreateBackupInput) (*request.Request, *CreateBackupOutput)

	CreateDBInstanceCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CreateDBInstanceCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateDBInstanceCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateDBInstance(*CreateDBInstanceInput) (*CreateDBInstanceOutput, error)
	CreateDBInstanceWithContext(volcengine.Context, *CreateDBInstanceInput, ...request.Option) (*CreateDBInstanceOutput, error)
	CreateDBInstanceRequest(*CreateDBInstanceInput) (*request.Request, *CreateDBInstanceOutput)

	CreateTosRestoreCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CreateTosRestoreCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateTosRestoreCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateTosRestore(*CreateTosRestoreInput) (*CreateTosRestoreOutput, error)
	CreateTosRestoreWithContext(volcengine.Context, *CreateTosRestoreInput, ...request.Option) (*CreateTosRestoreOutput, error)
	CreateTosRestoreRequest(*CreateTosRestoreInput) (*request.Request, *CreateTosRestoreOutput)

	DeleteBackupCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteBackupCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteBackupCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteBackup(*DeleteBackupInput) (*DeleteBackupOutput, error)
	DeleteBackupWithContext(volcengine.Context, *DeleteBackupInput, ...request.Option) (*DeleteBackupOutput, error)
	DeleteBackupRequest(*DeleteBackupInput) (*request.Request, *DeleteBackupOutput)

	DescribeAvailabilityZonesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeAvailabilityZonesCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeAvailabilityZonesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeAvailabilityZones(*DescribeAvailabilityZonesInput) (*DescribeAvailabilityZonesOutput, error)
	DescribeAvailabilityZonesWithContext(volcengine.Context, *DescribeAvailabilityZonesInput, ...request.Option) (*DescribeAvailabilityZonesOutput, error)
	DescribeAvailabilityZonesRequest(*DescribeAvailabilityZonesInput) (*request.Request, *DescribeAvailabilityZonesOutput)

	DescribeBackupsCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeBackupsCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeBackupsCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeBackups(*DescribeBackupsInput) (*DescribeBackupsOutput, error)
	DescribeBackupsWithContext(volcengine.Context, *DescribeBackupsInput, ...request.Option) (*DescribeBackupsOutput, error)
	DescribeBackupsRequest(*DescribeBackupsInput) (*request.Request, *DescribeBackupsOutput)

	DescribeDBInstanceDetailCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeDBInstanceDetailCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeDBInstanceDetailCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeDBInstanceDetail(*DescribeDBInstanceDetailInput) (*DescribeDBInstanceDetailOutput, error)
	DescribeDBInstanceDetailWithContext(volcengine.Context, *DescribeDBInstanceDetailInput, ...request.Option) (*DescribeDBInstanceDetailOutput, error)
	DescribeDBInstanceDetailRequest(*DescribeDBInstanceDetailInput) (*request.Request, *DescribeDBInstanceDetailOutput)

	DescribeDBInstanceParametersCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeDBInstanceParametersCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeDBInstanceParametersCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeDBInstanceParameters(*DescribeDBInstanceParametersInput) (*DescribeDBInstanceParametersOutput, error)
	DescribeDBInstanceParametersWithContext(volcengine.Context, *DescribeDBInstanceParametersInput, ...request.Option) (*DescribeDBInstanceParametersOutput, error)
	DescribeDBInstanceParametersRequest(*DescribeDBInstanceParametersInput) (*request.Request, *DescribeDBInstanceParametersOutput)

	DescribeDBInstancesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeDBInstancesCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeDBInstancesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeDBInstances(*DescribeDBInstancesInput) (*DescribeDBInstancesOutput, error)
	DescribeDBInstancesWithContext(volcengine.Context, *DescribeDBInstancesInput, ...request.Option) (*DescribeDBInstancesOutput, error)
	DescribeDBInstancesRequest(*DescribeDBInstancesInput) (*request.Request, *DescribeDBInstancesOutput)

	DescribeRegionsCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeRegionsCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeRegionsCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeRegions(*DescribeRegionsInput) (*DescribeRegionsOutput, error)
	DescribeRegionsWithContext(volcengine.Context, *DescribeRegionsInput, ...request.Option) (*DescribeRegionsOutput, error)
	DescribeRegionsRequest(*DescribeRegionsInput) (*request.Request, *DescribeRegionsOutput)

	DescribeTosRestoreTaskDetailCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeTosRestoreTaskDetailCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeTosRestoreTaskDetailCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeTosRestoreTaskDetail(*DescribeTosRestoreTaskDetailInput) (*DescribeTosRestoreTaskDetailOutput, error)
	DescribeTosRestoreTaskDetailWithContext(volcengine.Context, *DescribeTosRestoreTaskDetailInput, ...request.Option) (*DescribeTosRestoreTaskDetailOutput, error)
	DescribeTosRestoreTaskDetailRequest(*DescribeTosRestoreTaskDetailInput) (*request.Request, *DescribeTosRestoreTaskDetailOutput)

	DescribeTosRestoreTasksCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeTosRestoreTasksCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeTosRestoreTasksCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeTosRestoreTasks(*DescribeTosRestoreTasksInput) (*DescribeTosRestoreTasksOutput, error)
	DescribeTosRestoreTasksWithContext(volcengine.Context, *DescribeTosRestoreTasksInput, ...request.Option) (*DescribeTosRestoreTasksOutput, error)
	DescribeTosRestoreTasksRequest(*DescribeTosRestoreTasksInput) (*request.Request, *DescribeTosRestoreTasksOutput)

	DownloadBackupCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DownloadBackupCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DownloadBackupCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DownloadBackup(*DownloadBackupInput) (*DownloadBackupOutput, error)
	DownloadBackupWithContext(volcengine.Context, *DownloadBackupInput, ...request.Option) (*DownloadBackupOutput, error)
	DownloadBackupRequest(*DownloadBackupInput) (*request.Request, *DownloadBackupOutput)

	ModifyBackupPolicyCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ModifyBackupPolicyCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifyBackupPolicyCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifyBackupPolicy(*ModifyBackupPolicyInput) (*ModifyBackupPolicyOutput, error)
	ModifyBackupPolicyWithContext(volcengine.Context, *ModifyBackupPolicyInput, ...request.Option) (*ModifyBackupPolicyOutput, error)
	ModifyBackupPolicyRequest(*ModifyBackupPolicyInput) (*request.Request, *ModifyBackupPolicyOutput)

	RestoreToExistedInstanceCommon(*map[string]interface{}) (*map[string]interface{}, error)
	RestoreToExistedInstanceCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	RestoreToExistedInstanceCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	RestoreToExistedInstance(*RestoreToExistedInstanceInput) (*RestoreToExistedInstanceOutput, error)
	RestoreToExistedInstanceWithContext(volcengine.Context, *RestoreToExistedInstanceInput, ...request.Option) (*RestoreToExistedInstanceOutput, error)
	RestoreToExistedInstanceRequest(*RestoreToExistedInstanceInput) (*request.Request, *RestoreToExistedInstanceOutput)
}

var _ RDSMSSQLAPI = (*RDSMSSQL)(nil)

// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package acepiface provides an interface to enable mocking the ACEP service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package acep

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
)

// ACEPAPI provides an interface to enable mocking the
// acep.ACEP service client's API operation,
//
//    // volcengine sdk func uses an SDK service client to make a request to
//    // ACEP.
//    func myFunc(svc ACEPAPI) bool {
//        // Make svc.AddCustomRoute request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := acep.New(sess)
//
//        myFunc(svc)
//    }
//
type ACEPAPI interface {
	AddCustomRouteCommon(*map[string]interface{}) (*map[string]interface{}, error)
	AddCustomRouteCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	AddCustomRouteCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	AddCustomRoute(*AddCustomRouteInput) (*AddCustomRouteOutput, error)
	AddCustomRouteWithContext(volcengine.Context, *AddCustomRouteInput, ...request.Option) (*AddCustomRouteOutput, error)
	AddCustomRouteRequest(*AddCustomRouteInput) (*request.Request, *AddCustomRouteOutput)

	BackupPodCommon(*map[string]interface{}) (*map[string]interface{}, error)
	BackupPodCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	BackupPodCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	BackupPod(*BackupPodInput) (*BackupPodOutput, error)
	BackupPodWithContext(volcengine.Context, *BackupPodInput, ...request.Option) (*BackupPodOutput, error)
	BackupPodRequest(*BackupPodInput) (*request.Request, *BackupPodOutput)

	CancelBackupPodCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CancelBackupPodCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CancelBackupPodCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CancelBackupPod(*CancelBackupPodInput) (*CancelBackupPodOutput, error)
	CancelBackupPodWithContext(volcengine.Context, *CancelBackupPodInput, ...request.Option) (*CancelBackupPodOutput, error)
	CancelBackupPodRequest(*CancelBackupPodInput) (*request.Request, *CancelBackupPodOutput)

	CancelRestorePodCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CancelRestorePodCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CancelRestorePodCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CancelRestorePod(*CancelRestorePodInput) (*CancelRestorePodOutput, error)
	CancelRestorePodWithContext(volcengine.Context, *CancelRestorePodInput, ...request.Option) (*CancelRestorePodOutput, error)
	CancelRestorePodRequest(*CancelRestorePodInput) (*request.Request, *CancelRestorePodOutput)

	DeleteCustomRouteCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteCustomRouteCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteCustomRouteCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteCustomRoute(*DeleteCustomRouteInput) (*DeleteCustomRouteOutput, error)
	DeleteCustomRouteWithContext(volcengine.Context, *DeleteCustomRouteInput, ...request.Option) (*DeleteCustomRouteOutput, error)
	DeleteCustomRouteRequest(*DeleteCustomRouteInput) (*request.Request, *DeleteCustomRouteOutput)

	ListCustomRouteCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ListCustomRouteCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ListCustomRouteCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ListCustomRoute(*ListCustomRouteInput) (*ListCustomRouteOutput, error)
	ListCustomRouteWithContext(volcengine.Context, *ListCustomRouteInput, ...request.Option) (*ListCustomRouteOutput, error)
	ListCustomRouteRequest(*ListCustomRouteInput) (*request.Request, *ListCustomRouteOutput)

	ListPodCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ListPodCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ListPodCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ListPod(*ListPodInput) (*ListPodOutput, error)
	ListPodWithContext(volcengine.Context, *ListPodInput, ...request.Option) (*ListPodOutput, error)
	ListPodRequest(*ListPodInput) (*request.Request, *ListPodOutput)

	MigratePodCommon(*map[string]interface{}) (*map[string]interface{}, error)
	MigratePodCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	MigratePodCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	MigratePod(*MigratePodInput) (*MigratePodOutput, error)
	MigratePodWithContext(volcengine.Context, *MigratePodInput, ...request.Option) (*MigratePodOutput, error)
	MigratePodRequest(*MigratePodInput) (*request.Request, *MigratePodOutput)

	RestorePodCommon(*map[string]interface{}) (*map[string]interface{}, error)
	RestorePodCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	RestorePodCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	RestorePod(*RestorePodInput) (*RestorePodOutput, error)
	RestorePodWithContext(volcengine.Context, *RestorePodInput, ...request.Option) (*RestorePodOutput, error)
	RestorePodRequest(*RestorePodInput) (*request.Request, *RestorePodOutput)

	UpdateCustomRouteCommon(*map[string]interface{}) (*map[string]interface{}, error)
	UpdateCustomRouteCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	UpdateCustomRouteCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	UpdateCustomRoute(*UpdateCustomRouteInput) (*UpdateCustomRouteOutput, error)
	UpdateCustomRouteWithContext(volcengine.Context, *UpdateCustomRouteInput, ...request.Option) (*UpdateCustomRouteOutput, error)
	UpdateCustomRouteRequest(*UpdateCustomRouteInput) (*request.Request, *UpdateCustomRouteOutput)
}

var _ ACEPAPI = (*ACEP)(nil)

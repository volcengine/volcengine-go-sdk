// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package storageebsiface provides an interface to enable mocking the STORAGE_EBS service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package storageebs

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
)

// STORAGEEBSAPI provides an interface to enable mocking the
// storageebs.STORAGEEBS service client's API operation,
//
//    // volcengine sdk func uses an SDK service client to make a request to
//    // STORAGE_EBS.
//    func myFunc(svc STORAGEEBSAPI) bool {
//        // Make svc.AttachVolume request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := storageebs.New(sess)
//
//        myFunc(svc)
//    }
//
type STORAGEEBSAPI interface {
	AttachVolumeCommon(*map[string]interface{}) (*map[string]interface{}, error)
	AttachVolumeCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	AttachVolumeCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	AttachVolume(*AttachVolumeInput) (*AttachVolumeOutput, error)
	AttachVolumeWithContext(volcengine.Context, *AttachVolumeInput, ...request.Option) (*AttachVolumeOutput, error)
	AttachVolumeRequest(*AttachVolumeInput) (*request.Request, *AttachVolumeOutput)

	CreateVolumeCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CreateVolumeCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateVolumeCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateVolume(*CreateVolumeInput) (*CreateVolumeOutput, error)
	CreateVolumeWithContext(volcengine.Context, *CreateVolumeInput, ...request.Option) (*CreateVolumeOutput, error)
	CreateVolumeRequest(*CreateVolumeInput) (*request.Request, *CreateVolumeOutput)

	DeleteVolumeCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteVolumeCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteVolumeCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteVolume(*DeleteVolumeInput) (*DeleteVolumeOutput, error)
	DeleteVolumeWithContext(volcengine.Context, *DeleteVolumeInput, ...request.Option) (*DeleteVolumeOutput, error)
	DeleteVolumeRequest(*DeleteVolumeInput) (*request.Request, *DeleteVolumeOutput)

	DescribeVolumesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeVolumesCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeVolumesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeVolumes(*DescribeVolumesInput) (*DescribeVolumesOutput, error)
	DescribeVolumesWithContext(volcengine.Context, *DescribeVolumesInput, ...request.Option) (*DescribeVolumesOutput, error)
	DescribeVolumesRequest(*DescribeVolumesInput) (*request.Request, *DescribeVolumesOutput)

	DetachVolumeCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DetachVolumeCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DetachVolumeCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DetachVolume(*DetachVolumeInput) (*DetachVolumeOutput, error)
	DetachVolumeWithContext(volcengine.Context, *DetachVolumeInput, ...request.Option) (*DetachVolumeOutput, error)
	DetachVolumeRequest(*DetachVolumeInput) (*request.Request, *DetachVolumeOutput)

	ExtendVolumeCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ExtendVolumeCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ExtendVolumeCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ExtendVolume(*ExtendVolumeInput) (*ExtendVolumeOutput, error)
	ExtendVolumeWithContext(volcengine.Context, *ExtendVolumeInput, ...request.Option) (*ExtendVolumeOutput, error)
	ExtendVolumeRequest(*ExtendVolumeInput) (*request.Request, *ExtendVolumeOutput)

	ModifyVolumeAttributeCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ModifyVolumeAttributeCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifyVolumeAttributeCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifyVolumeAttribute(*ModifyVolumeAttributeInput) (*ModifyVolumeAttributeOutput, error)
	ModifyVolumeAttributeWithContext(volcengine.Context, *ModifyVolumeAttributeInput, ...request.Option) (*ModifyVolumeAttributeOutput, error)
	ModifyVolumeAttributeRequest(*ModifyVolumeAttributeInput) (*request.Request, *ModifyVolumeAttributeOutput)

	ModifyVolumeChargeTypeCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ModifyVolumeChargeTypeCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifyVolumeChargeTypeCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifyVolumeChargeType(*ModifyVolumeChargeTypeInput) (*ModifyVolumeChargeTypeOutput, error)
	ModifyVolumeChargeTypeWithContext(volcengine.Context, *ModifyVolumeChargeTypeInput, ...request.Option) (*ModifyVolumeChargeTypeOutput, error)
	ModifyVolumeChargeTypeRequest(*ModifyVolumeChargeTypeInput) (*request.Request, *ModifyVolumeChargeTypeOutput)

	ModifyVolumeKindCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ModifyVolumeKindCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifyVolumeKindCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifyVolumeKind(*ModifyVolumeKindInput) (*ModifyVolumeKindOutput, error)
	ModifyVolumeKindWithContext(volcengine.Context, *ModifyVolumeKindInput, ...request.Option) (*ModifyVolumeKindOutput, error)
	ModifyVolumeKindRequest(*ModifyVolumeKindInput) (*request.Request, *ModifyVolumeKindOutput)

	TerminateVolumesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	TerminateVolumesCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	TerminateVolumesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	TerminateVolumes(*TerminateVolumesInput) (*TerminateVolumesOutput, error)
	TerminateVolumesWithContext(volcengine.Context, *TerminateVolumesInput, ...request.Option) (*TerminateVolumesOutput, error)
	TerminateVolumesRequest(*TerminateVolumesInput) (*request.Request, *TerminateVolumesOutput)
}

var _ STORAGEEBSAPI = (*STORAGEEBS)(nil)

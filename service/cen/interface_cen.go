// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package ceniface provides an interface to enable mocking the CEN service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package cen

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
)

// CENAPI provides an interface to enable mocking the
// cen.CEN service client's API operation,
//
//    // volcengine sdk func uses an SDK service client to make a request to
//    // CEN.
//    func myFunc(svc CENAPI) bool {
//        // Make svc.AssociateCenBandwidthPackage request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := cen.New(sess)
//
//        myFunc(svc)
//    }
//
type CENAPI interface {
	AssociateCenBandwidthPackageCommon(*map[string]interface{}) (*map[string]interface{}, error)
	AssociateCenBandwidthPackageCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	AssociateCenBandwidthPackageCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	AssociateCenBandwidthPackage(*AssociateCenBandwidthPackageInput) (*AssociateCenBandwidthPackageOutput, error)
	AssociateCenBandwidthPackageWithContext(volcengine.Context, *AssociateCenBandwidthPackageInput, ...request.Option) (*AssociateCenBandwidthPackageOutput, error)
	AssociateCenBandwidthPackageRequest(*AssociateCenBandwidthPackageInput) (*request.Request, *AssociateCenBandwidthPackageOutput)

	AttachInstanceToCenCommon(*map[string]interface{}) (*map[string]interface{}, error)
	AttachInstanceToCenCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	AttachInstanceToCenCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	AttachInstanceToCen(*AttachInstanceToCenInput) (*AttachInstanceToCenOutput, error)
	AttachInstanceToCenWithContext(volcengine.Context, *AttachInstanceToCenInput, ...request.Option) (*AttachInstanceToCenOutput, error)
	AttachInstanceToCenRequest(*AttachInstanceToCenInput) (*request.Request, *AttachInstanceToCenOutput)

	CreateCenCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CreateCenCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateCenCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateCen(*CreateCenInput) (*CreateCenOutput, error)
	CreateCenWithContext(volcengine.Context, *CreateCenInput, ...request.Option) (*CreateCenOutput, error)
	CreateCenRequest(*CreateCenInput) (*request.Request, *CreateCenOutput)

	CreateCenBandwidthPackageCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CreateCenBandwidthPackageCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateCenBandwidthPackageCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateCenBandwidthPackage(*CreateCenBandwidthPackageInput) (*CreateCenBandwidthPackageOutput, error)
	CreateCenBandwidthPackageWithContext(volcengine.Context, *CreateCenBandwidthPackageInput, ...request.Option) (*CreateCenBandwidthPackageOutput, error)
	CreateCenBandwidthPackageRequest(*CreateCenBandwidthPackageInput) (*request.Request, *CreateCenBandwidthPackageOutput)

	CreateCenInterRegionBandwidthCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CreateCenInterRegionBandwidthCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateCenInterRegionBandwidthCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateCenInterRegionBandwidth(*CreateCenInterRegionBandwidthInput) (*CreateCenInterRegionBandwidthOutput, error)
	CreateCenInterRegionBandwidthWithContext(volcengine.Context, *CreateCenInterRegionBandwidthInput, ...request.Option) (*CreateCenInterRegionBandwidthOutput, error)
	CreateCenInterRegionBandwidthRequest(*CreateCenInterRegionBandwidthInput) (*request.Request, *CreateCenInterRegionBandwidthOutput)

	DeleteCenCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteCenCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteCenCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteCen(*DeleteCenInput) (*DeleteCenOutput, error)
	DeleteCenWithContext(volcengine.Context, *DeleteCenInput, ...request.Option) (*DeleteCenOutput, error)
	DeleteCenRequest(*DeleteCenInput) (*request.Request, *DeleteCenOutput)

	DeleteCenBandwidthPackageCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteCenBandwidthPackageCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteCenBandwidthPackageCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteCenBandwidthPackage(*DeleteCenBandwidthPackageInput) (*DeleteCenBandwidthPackageOutput, error)
	DeleteCenBandwidthPackageWithContext(volcengine.Context, *DeleteCenBandwidthPackageInput, ...request.Option) (*DeleteCenBandwidthPackageOutput, error)
	DeleteCenBandwidthPackageRequest(*DeleteCenBandwidthPackageInput) (*request.Request, *DeleteCenBandwidthPackageOutput)

	DeleteCenInterRegionBandwidthCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteCenInterRegionBandwidthCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteCenInterRegionBandwidthCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteCenInterRegionBandwidth(*DeleteCenInterRegionBandwidthInput) (*DeleteCenInterRegionBandwidthOutput, error)
	DeleteCenInterRegionBandwidthWithContext(volcengine.Context, *DeleteCenInterRegionBandwidthInput, ...request.Option) (*DeleteCenInterRegionBandwidthOutput, error)
	DeleteCenInterRegionBandwidthRequest(*DeleteCenInterRegionBandwidthInput) (*request.Request, *DeleteCenInterRegionBandwidthOutput)

	DescribeCenAttachedInstanceAttributesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeCenAttachedInstanceAttributesCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeCenAttachedInstanceAttributesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeCenAttachedInstanceAttributes(*DescribeCenAttachedInstanceAttributesInput) (*DescribeCenAttachedInstanceAttributesOutput, error)
	DescribeCenAttachedInstanceAttributesWithContext(volcengine.Context, *DescribeCenAttachedInstanceAttributesInput, ...request.Option) (*DescribeCenAttachedInstanceAttributesOutput, error)
	DescribeCenAttachedInstanceAttributesRequest(*DescribeCenAttachedInstanceAttributesInput) (*request.Request, *DescribeCenAttachedInstanceAttributesOutput)

	DescribeCenAttachedInstancesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeCenAttachedInstancesCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeCenAttachedInstancesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeCenAttachedInstances(*DescribeCenAttachedInstancesInput) (*DescribeCenAttachedInstancesOutput, error)
	DescribeCenAttachedInstancesWithContext(volcengine.Context, *DescribeCenAttachedInstancesInput, ...request.Option) (*DescribeCenAttachedInstancesOutput, error)
	DescribeCenAttachedInstancesRequest(*DescribeCenAttachedInstancesInput) (*request.Request, *DescribeCenAttachedInstancesOutput)

	DescribeCenAttributesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeCenAttributesCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeCenAttributesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeCenAttributes(*DescribeCenAttributesInput) (*DescribeCenAttributesOutput, error)
	DescribeCenAttributesWithContext(volcengine.Context, *DescribeCenAttributesInput, ...request.Option) (*DescribeCenAttributesOutput, error)
	DescribeCenAttributesRequest(*DescribeCenAttributesInput) (*request.Request, *DescribeCenAttributesOutput)

	DescribeCenBandwidthPackageAttributesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeCenBandwidthPackageAttributesCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeCenBandwidthPackageAttributesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeCenBandwidthPackageAttributes(*DescribeCenBandwidthPackageAttributesInput) (*DescribeCenBandwidthPackageAttributesOutput, error)
	DescribeCenBandwidthPackageAttributesWithContext(volcengine.Context, *DescribeCenBandwidthPackageAttributesInput, ...request.Option) (*DescribeCenBandwidthPackageAttributesOutput, error)
	DescribeCenBandwidthPackageAttributesRequest(*DescribeCenBandwidthPackageAttributesInput) (*request.Request, *DescribeCenBandwidthPackageAttributesOutput)

	DescribeCenBandwidthPackagesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeCenBandwidthPackagesCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeCenBandwidthPackagesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeCenBandwidthPackages(*DescribeCenBandwidthPackagesInput) (*DescribeCenBandwidthPackagesOutput, error)
	DescribeCenBandwidthPackagesWithContext(volcengine.Context, *DescribeCenBandwidthPackagesInput, ...request.Option) (*DescribeCenBandwidthPackagesOutput, error)
	DescribeCenBandwidthPackagesRequest(*DescribeCenBandwidthPackagesInput) (*request.Request, *DescribeCenBandwidthPackagesOutput)

	DescribeCenInterRegionBandwidthAttributesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeCenInterRegionBandwidthAttributesCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeCenInterRegionBandwidthAttributesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeCenInterRegionBandwidthAttributes(*DescribeCenInterRegionBandwidthAttributesInput) (*DescribeCenInterRegionBandwidthAttributesOutput, error)
	DescribeCenInterRegionBandwidthAttributesWithContext(volcengine.Context, *DescribeCenInterRegionBandwidthAttributesInput, ...request.Option) (*DescribeCenInterRegionBandwidthAttributesOutput, error)
	DescribeCenInterRegionBandwidthAttributesRequest(*DescribeCenInterRegionBandwidthAttributesInput) (*request.Request, *DescribeCenInterRegionBandwidthAttributesOutput)

	DescribeCenInterRegionBandwidthsCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeCenInterRegionBandwidthsCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeCenInterRegionBandwidthsCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeCenInterRegionBandwidths(*DescribeCenInterRegionBandwidthsInput) (*DescribeCenInterRegionBandwidthsOutput, error)
	DescribeCenInterRegionBandwidthsWithContext(volcengine.Context, *DescribeCenInterRegionBandwidthsInput, ...request.Option) (*DescribeCenInterRegionBandwidthsOutput, error)
	DescribeCenInterRegionBandwidthsRequest(*DescribeCenInterRegionBandwidthsInput) (*request.Request, *DescribeCenInterRegionBandwidthsOutput)

	DescribeCenRouteEntriesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeCenRouteEntriesCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeCenRouteEntriesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeCenRouteEntries(*DescribeCenRouteEntriesInput) (*DescribeCenRouteEntriesOutput, error)
	DescribeCenRouteEntriesWithContext(volcengine.Context, *DescribeCenRouteEntriesInput, ...request.Option) (*DescribeCenRouteEntriesOutput, error)
	DescribeCenRouteEntriesRequest(*DescribeCenRouteEntriesInput) (*request.Request, *DescribeCenRouteEntriesOutput)

	DescribeCensCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeCensCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeCensCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeCens(*DescribeCensInput) (*DescribeCensOutput, error)
	DescribeCensWithContext(volcengine.Context, *DescribeCensInput, ...request.Option) (*DescribeCensOutput, error)
	DescribeCensRequest(*DescribeCensInput) (*request.Request, *DescribeCensOutput)

	DescribeGrantRulesToCenCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeGrantRulesToCenCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeGrantRulesToCenCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeGrantRulesToCen(*DescribeGrantRulesToCenInput) (*DescribeGrantRulesToCenOutput, error)
	DescribeGrantRulesToCenWithContext(volcengine.Context, *DescribeGrantRulesToCenInput, ...request.Option) (*DescribeGrantRulesToCenOutput, error)
	DescribeGrantRulesToCenRequest(*DescribeGrantRulesToCenInput) (*request.Request, *DescribeGrantRulesToCenOutput)

	DescribeInstanceGrantedRulesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeInstanceGrantedRulesCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeInstanceGrantedRulesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeInstanceGrantedRules(*DescribeInstanceGrantedRulesInput) (*DescribeInstanceGrantedRulesOutput, error)
	DescribeInstanceGrantedRulesWithContext(volcengine.Context, *DescribeInstanceGrantedRulesInput, ...request.Option) (*DescribeInstanceGrantedRulesOutput, error)
	DescribeInstanceGrantedRulesRequest(*DescribeInstanceGrantedRulesInput) (*request.Request, *DescribeInstanceGrantedRulesOutput)

	DetachInstanceFromCenCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DetachInstanceFromCenCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DetachInstanceFromCenCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DetachInstanceFromCen(*DetachInstanceFromCenInput) (*DetachInstanceFromCenOutput, error)
	DetachInstanceFromCenWithContext(volcengine.Context, *DetachInstanceFromCenInput, ...request.Option) (*DetachInstanceFromCenOutput, error)
	DetachInstanceFromCenRequest(*DetachInstanceFromCenInput) (*request.Request, *DetachInstanceFromCenOutput)

	DisassociateCenBandwidthPackageCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DisassociateCenBandwidthPackageCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DisassociateCenBandwidthPackageCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DisassociateCenBandwidthPackage(*DisassociateCenBandwidthPackageInput) (*DisassociateCenBandwidthPackageOutput, error)
	DisassociateCenBandwidthPackageWithContext(volcengine.Context, *DisassociateCenBandwidthPackageInput, ...request.Option) (*DisassociateCenBandwidthPackageOutput, error)
	DisassociateCenBandwidthPackageRequest(*DisassociateCenBandwidthPackageInput) (*request.Request, *DisassociateCenBandwidthPackageOutput)

	GrantInstanceToCenCommon(*map[string]interface{}) (*map[string]interface{}, error)
	GrantInstanceToCenCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	GrantInstanceToCenCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	GrantInstanceToCen(*GrantInstanceToCenInput) (*GrantInstanceToCenOutput, error)
	GrantInstanceToCenWithContext(volcengine.Context, *GrantInstanceToCenInput, ...request.Option) (*GrantInstanceToCenOutput, error)
	GrantInstanceToCenRequest(*GrantInstanceToCenInput) (*request.Request, *GrantInstanceToCenOutput)

	ModifyCenAttributesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ModifyCenAttributesCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifyCenAttributesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifyCenAttributes(*ModifyCenAttributesInput) (*ModifyCenAttributesOutput, error)
	ModifyCenAttributesWithContext(volcengine.Context, *ModifyCenAttributesInput, ...request.Option) (*ModifyCenAttributesOutput, error)
	ModifyCenAttributesRequest(*ModifyCenAttributesInput) (*request.Request, *ModifyCenAttributesOutput)

	ModifyCenBandwidthPackageAttributesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ModifyCenBandwidthPackageAttributesCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifyCenBandwidthPackageAttributesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifyCenBandwidthPackageAttributes(*ModifyCenBandwidthPackageAttributesInput) (*ModifyCenBandwidthPackageAttributesOutput, error)
	ModifyCenBandwidthPackageAttributesWithContext(volcengine.Context, *ModifyCenBandwidthPackageAttributesInput, ...request.Option) (*ModifyCenBandwidthPackageAttributesOutput, error)
	ModifyCenBandwidthPackageAttributesRequest(*ModifyCenBandwidthPackageAttributesInput) (*request.Request, *ModifyCenBandwidthPackageAttributesOutput)

	ModifyCenInterRegionBandwidthAttributesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ModifyCenInterRegionBandwidthAttributesCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifyCenInterRegionBandwidthAttributesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifyCenInterRegionBandwidthAttributes(*ModifyCenInterRegionBandwidthAttributesInput) (*ModifyCenInterRegionBandwidthAttributesOutput, error)
	ModifyCenInterRegionBandwidthAttributesWithContext(volcengine.Context, *ModifyCenInterRegionBandwidthAttributesInput, ...request.Option) (*ModifyCenInterRegionBandwidthAttributesOutput, error)
	ModifyCenInterRegionBandwidthAttributesRequest(*ModifyCenInterRegionBandwidthAttributesInput) (*request.Request, *ModifyCenInterRegionBandwidthAttributesOutput)

	RevokeInstanceFromCenCommon(*map[string]interface{}) (*map[string]interface{}, error)
	RevokeInstanceFromCenCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	RevokeInstanceFromCenCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	RevokeInstanceFromCen(*RevokeInstanceFromCenInput) (*RevokeInstanceFromCenOutput, error)
	RevokeInstanceFromCenWithContext(volcengine.Context, *RevokeInstanceFromCenInput, ...request.Option) (*RevokeInstanceFromCenOutput, error)
	RevokeInstanceFromCenRequest(*RevokeInstanceFromCenInput) (*request.Request, *RevokeInstanceFromCenOutput)
}

var _ CENAPI = (*CEN)(nil)

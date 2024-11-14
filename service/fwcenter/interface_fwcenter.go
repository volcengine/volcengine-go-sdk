// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package fwcenteriface provides an interface to enable mocking the FWCENTER service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package fwcenter

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
)

// FWCENTERAPI provides an interface to enable mocking the
// fwcenter.FWCENTER service client's API operation,
//
//    // volcengine sdk func uses an SDK service client to make a request to
//    // FWCENTER.
//    func myFunc(svc FWCENTERAPI) bool {
//        // Make svc.AddAddressBook request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := fwcenter.New(sess)
//
//        myFunc(svc)
//    }
//
type FWCENTERAPI interface {
	AddAddressBookCommon(*map[string]interface{}) (*map[string]interface{}, error)
	AddAddressBookCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	AddAddressBookCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	AddAddressBook(*AddAddressBookInput) (*AddAddressBookOutput, error)
	AddAddressBookWithContext(volcengine.Context, *AddAddressBookInput, ...request.Option) (*AddAddressBookOutput, error)
	AddAddressBookRequest(*AddAddressBookInput) (*request.Request, *AddAddressBookOutput)

	AddControlPolicyCommon(*map[string]interface{}) (*map[string]interface{}, error)
	AddControlPolicyCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	AddControlPolicyCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	AddControlPolicy(*AddControlPolicyInput) (*AddControlPolicyOutput, error)
	AddControlPolicyWithContext(volcengine.Context, *AddControlPolicyInput, ...request.Option) (*AddControlPolicyOutput, error)
	AddControlPolicyRequest(*AddControlPolicyInput) (*request.Request, *AddControlPolicyOutput)

	AddDnsControlPolicyCommon(*map[string]interface{}) (*map[string]interface{}, error)
	AddDnsControlPolicyCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	AddDnsControlPolicyCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	AddDnsControlPolicy(*AddDnsControlPolicyInput) (*AddDnsControlPolicyOutput, error)
	AddDnsControlPolicyWithContext(volcengine.Context, *AddDnsControlPolicyInput, ...request.Option) (*AddDnsControlPolicyOutput, error)
	AddDnsControlPolicyRequest(*AddDnsControlPolicyInput) (*request.Request, *AddDnsControlPolicyOutput)

	AddNatFirewallControlPolicyCommon(*map[string]interface{}) (*map[string]interface{}, error)
	AddNatFirewallControlPolicyCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	AddNatFirewallControlPolicyCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	AddNatFirewallControlPolicy(*AddNatFirewallControlPolicyInput) (*AddNatFirewallControlPolicyOutput, error)
	AddNatFirewallControlPolicyWithContext(volcengine.Context, *AddNatFirewallControlPolicyInput, ...request.Option) (*AddNatFirewallControlPolicyOutput, error)
	AddNatFirewallControlPolicyRequest(*AddNatFirewallControlPolicyInput) (*request.Request, *AddNatFirewallControlPolicyOutput)

	AddVpcFirewallAclRuleCommon(*map[string]interface{}) (*map[string]interface{}, error)
	AddVpcFirewallAclRuleCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	AddVpcFirewallAclRuleCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	AddVpcFirewallAclRule(*AddVpcFirewallAclRuleInput) (*AddVpcFirewallAclRuleOutput, error)
	AddVpcFirewallAclRuleWithContext(volcengine.Context, *AddVpcFirewallAclRuleInput, ...request.Option) (*AddVpcFirewallAclRuleOutput, error)
	AddVpcFirewallAclRuleRequest(*AddVpcFirewallAclRuleInput) (*request.Request, *AddVpcFirewallAclRuleOutput)

	AssetListCommon(*map[string]interface{}) (*map[string]interface{}, error)
	AssetListCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	AssetListCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	AssetList(*AssetListInput) (*AssetListOutput, error)
	AssetListWithContext(volcengine.Context, *AssetListInput, ...request.Option) (*AssetListOutput, error)
	AssetListRequest(*AssetListInput) (*request.Request, *AssetListOutput)

	DeleteAddressBookCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteAddressBookCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteAddressBookCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteAddressBook(*DeleteAddressBookInput) (*DeleteAddressBookOutput, error)
	DeleteAddressBookWithContext(volcengine.Context, *DeleteAddressBookInput, ...request.Option) (*DeleteAddressBookOutput, error)
	DeleteAddressBookRequest(*DeleteAddressBookInput) (*request.Request, *DeleteAddressBookOutput)

	DeleteControlPolicyCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteControlPolicyCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteControlPolicyCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteControlPolicy(*DeleteControlPolicyInput) (*DeleteControlPolicyOutput, error)
	DeleteControlPolicyWithContext(volcengine.Context, *DeleteControlPolicyInput, ...request.Option) (*DeleteControlPolicyOutput, error)
	DeleteControlPolicyRequest(*DeleteControlPolicyInput) (*request.Request, *DeleteControlPolicyOutput)

	DeleteDnsControlPolicyCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteDnsControlPolicyCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteDnsControlPolicyCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteDnsControlPolicy(*DeleteDnsControlPolicyInput) (*DeleteDnsControlPolicyOutput, error)
	DeleteDnsControlPolicyWithContext(volcengine.Context, *DeleteDnsControlPolicyInput, ...request.Option) (*DeleteDnsControlPolicyOutput, error)
	DeleteDnsControlPolicyRequest(*DeleteDnsControlPolicyInput) (*request.Request, *DeleteDnsControlPolicyOutput)

	DeleteNatFirewallControlPolicyCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteNatFirewallControlPolicyCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteNatFirewallControlPolicyCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteNatFirewallControlPolicy(*DeleteNatFirewallControlPolicyInput) (*DeleteNatFirewallControlPolicyOutput, error)
	DeleteNatFirewallControlPolicyWithContext(volcengine.Context, *DeleteNatFirewallControlPolicyInput, ...request.Option) (*DeleteNatFirewallControlPolicyOutput, error)
	DeleteNatFirewallControlPolicyRequest(*DeleteNatFirewallControlPolicyInput) (*request.Request, *DeleteNatFirewallControlPolicyOutput)

	DeleteVpcFirewallAclRuleCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteVpcFirewallAclRuleCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteVpcFirewallAclRuleCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteVpcFirewallAclRule(*DeleteVpcFirewallAclRuleInput) (*DeleteVpcFirewallAclRuleOutput, error)
	DeleteVpcFirewallAclRuleWithContext(volcengine.Context, *DeleteVpcFirewallAclRuleInput, ...request.Option) (*DeleteVpcFirewallAclRuleOutput, error)
	DeleteVpcFirewallAclRuleRequest(*DeleteVpcFirewallAclRuleInput) (*request.Request, *DeleteVpcFirewallAclRuleOutput)

	DescribeAddressBookCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeAddressBookCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeAddressBookCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeAddressBook(*DescribeAddressBookInput) (*DescribeAddressBookOutput, error)
	DescribeAddressBookWithContext(volcengine.Context, *DescribeAddressBookInput, ...request.Option) (*DescribeAddressBookOutput, error)
	DescribeAddressBookRequest(*DescribeAddressBookInput) (*request.Request, *DescribeAddressBookOutput)

	DescribeControlPolicyCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeControlPolicyCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeControlPolicyCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeControlPolicy(*DescribeControlPolicyInput) (*DescribeControlPolicyOutput, error)
	DescribeControlPolicyWithContext(volcengine.Context, *DescribeControlPolicyInput, ...request.Option) (*DescribeControlPolicyOutput, error)
	DescribeControlPolicyRequest(*DescribeControlPolicyInput) (*request.Request, *DescribeControlPolicyOutput)

	DescribeControlPolicyByRuleIdCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeControlPolicyByRuleIdCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeControlPolicyByRuleIdCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeControlPolicyByRuleId(*DescribeControlPolicyByRuleIdInput) (*DescribeControlPolicyByRuleIdOutput, error)
	DescribeControlPolicyByRuleIdWithContext(volcengine.Context, *DescribeControlPolicyByRuleIdInput, ...request.Option) (*DescribeControlPolicyByRuleIdOutput, error)
	DescribeControlPolicyByRuleIdRequest(*DescribeControlPolicyByRuleIdInput) (*request.Request, *DescribeControlPolicyByRuleIdOutput)

	DescribeControlPolicyPriorUsedCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeControlPolicyPriorUsedCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeControlPolicyPriorUsedCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeControlPolicyPriorUsed(*DescribeControlPolicyPriorUsedInput) (*DescribeControlPolicyPriorUsedOutput, error)
	DescribeControlPolicyPriorUsedWithContext(volcengine.Context, *DescribeControlPolicyPriorUsedInput, ...request.Option) (*DescribeControlPolicyPriorUsedOutput, error)
	DescribeControlPolicyPriorUsedRequest(*DescribeControlPolicyPriorUsedInput) (*request.Request, *DescribeControlPolicyPriorUsedOutput)

	DescribeDnsControlPolicyCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeDnsControlPolicyCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeDnsControlPolicyCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeDnsControlPolicy(*DescribeDnsControlPolicyInput) (*DescribeDnsControlPolicyOutput, error)
	DescribeDnsControlPolicyWithContext(volcengine.Context, *DescribeDnsControlPolicyInput, ...request.Option) (*DescribeDnsControlPolicyOutput, error)
	DescribeDnsControlPolicyRequest(*DescribeDnsControlPolicyInput) (*request.Request, *DescribeDnsControlPolicyOutput)

	DescribeNatFirewallControlPolicyCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeNatFirewallControlPolicyCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeNatFirewallControlPolicyCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeNatFirewallControlPolicy(*DescribeNatFirewallControlPolicyInput) (*DescribeNatFirewallControlPolicyOutput, error)
	DescribeNatFirewallControlPolicyWithContext(volcengine.Context, *DescribeNatFirewallControlPolicyInput, ...request.Option) (*DescribeNatFirewallControlPolicyOutput, error)
	DescribeNatFirewallControlPolicyRequest(*DescribeNatFirewallControlPolicyInput) (*request.Request, *DescribeNatFirewallControlPolicyOutput)

	DescribeNatFirewallControlPolicyPriorityUsedCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeNatFirewallControlPolicyPriorityUsedCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeNatFirewallControlPolicyPriorityUsedCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeNatFirewallControlPolicyPriorityUsed(*DescribeNatFirewallControlPolicyPriorityUsedInput) (*DescribeNatFirewallControlPolicyPriorityUsedOutput, error)
	DescribeNatFirewallControlPolicyPriorityUsedWithContext(volcengine.Context, *DescribeNatFirewallControlPolicyPriorityUsedInput, ...request.Option) (*DescribeNatFirewallControlPolicyPriorityUsedOutput, error)
	DescribeNatFirewallControlPolicyPriorityUsedRequest(*DescribeNatFirewallControlPolicyPriorityUsedInput) (*request.Request, *DescribeNatFirewallControlPolicyPriorityUsedOutput)

	DescribeNatFirewallListCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeNatFirewallListCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeNatFirewallListCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeNatFirewallList(*DescribeNatFirewallListInput) (*DescribeNatFirewallListOutput, error)
	DescribeNatFirewallListWithContext(volcengine.Context, *DescribeNatFirewallListInput, ...request.Option) (*DescribeNatFirewallListOutput, error)
	DescribeNatFirewallListRequest(*DescribeNatFirewallListInput) (*request.Request, *DescribeNatFirewallListOutput)

	DescribeVpcFirewallAclRuleListCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeVpcFirewallAclRuleListCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeVpcFirewallAclRuleListCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeVpcFirewallAclRuleList(*DescribeVpcFirewallAclRuleListInput) (*DescribeVpcFirewallAclRuleListOutput, error)
	DescribeVpcFirewallAclRuleListWithContext(volcengine.Context, *DescribeVpcFirewallAclRuleListInput, ...request.Option) (*DescribeVpcFirewallAclRuleListOutput, error)
	DescribeVpcFirewallAclRuleListRequest(*DescribeVpcFirewallAclRuleListInput) (*request.Request, *DescribeVpcFirewallAclRuleListOutput)

	DescribeVpcFirewallAclRulePriorUsedCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeVpcFirewallAclRulePriorUsedCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeVpcFirewallAclRulePriorUsedCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeVpcFirewallAclRulePriorUsed(*DescribeVpcFirewallAclRulePriorUsedInput) (*DescribeVpcFirewallAclRulePriorUsedOutput, error)
	DescribeVpcFirewallAclRulePriorUsedWithContext(volcengine.Context, *DescribeVpcFirewallAclRulePriorUsedInput, ...request.Option) (*DescribeVpcFirewallAclRulePriorUsedOutput, error)
	DescribeVpcFirewallAclRulePriorUsedRequest(*DescribeVpcFirewallAclRulePriorUsedInput) (*request.Request, *DescribeVpcFirewallAclRulePriorUsedOutput)

	DescribeVpcFirewallListCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeVpcFirewallListCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeVpcFirewallListCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeVpcFirewallList(*DescribeVpcFirewallListInput) (*DescribeVpcFirewallListOutput, error)
	DescribeVpcFirewallListWithContext(volcengine.Context, *DescribeVpcFirewallListInput, ...request.Option) (*DescribeVpcFirewallListOutput, error)
	DescribeVpcFirewallListRequest(*DescribeVpcFirewallListInput) (*request.Request, *DescribeVpcFirewallListOutput)

	DescribeVpcsCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeVpcsCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeVpcsCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeVpcs(*DescribeVpcsInput) (*DescribeVpcsOutput, error)
	DescribeVpcsWithContext(volcengine.Context, *DescribeVpcsInput, ...request.Option) (*DescribeVpcsOutput, error)
	DescribeVpcsRequest(*DescribeVpcsInput) (*request.Request, *DescribeVpcsOutput)

	ModifyAddressBookCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ModifyAddressBookCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifyAddressBookCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifyAddressBook(*ModifyAddressBookInput) (*ModifyAddressBookOutput, error)
	ModifyAddressBookWithContext(volcengine.Context, *ModifyAddressBookInput, ...request.Option) (*ModifyAddressBookOutput, error)
	ModifyAddressBookRequest(*ModifyAddressBookInput) (*request.Request, *ModifyAddressBookOutput)

	ModifyControlPolicyCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ModifyControlPolicyCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifyControlPolicyCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifyControlPolicy(*ModifyControlPolicyInput) (*ModifyControlPolicyOutput, error)
	ModifyControlPolicyWithContext(volcengine.Context, *ModifyControlPolicyInput, ...request.Option) (*ModifyControlPolicyOutput, error)
	ModifyControlPolicyRequest(*ModifyControlPolicyInput) (*request.Request, *ModifyControlPolicyOutput)

	ModifyControlPolicyPositionCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ModifyControlPolicyPositionCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifyControlPolicyPositionCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifyControlPolicyPosition(*ModifyControlPolicyPositionInput) (*ModifyControlPolicyPositionOutput, error)
	ModifyControlPolicyPositionWithContext(volcengine.Context, *ModifyControlPolicyPositionInput, ...request.Option) (*ModifyControlPolicyPositionOutput, error)
	ModifyControlPolicyPositionRequest(*ModifyControlPolicyPositionInput) (*request.Request, *ModifyControlPolicyPositionOutput)

	ModifyDnsControlPolicyCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ModifyDnsControlPolicyCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifyDnsControlPolicyCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifyDnsControlPolicy(*ModifyDnsControlPolicyInput) (*ModifyDnsControlPolicyOutput, error)
	ModifyDnsControlPolicyWithContext(volcengine.Context, *ModifyDnsControlPolicyInput, ...request.Option) (*ModifyDnsControlPolicyOutput, error)
	ModifyDnsControlPolicyRequest(*ModifyDnsControlPolicyInput) (*request.Request, *ModifyDnsControlPolicyOutput)

	ModifyNatFirewallControlPolicyCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ModifyNatFirewallControlPolicyCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifyNatFirewallControlPolicyCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifyNatFirewallControlPolicy(*ModifyNatFirewallControlPolicyInput) (*ModifyNatFirewallControlPolicyOutput, error)
	ModifyNatFirewallControlPolicyWithContext(volcengine.Context, *ModifyNatFirewallControlPolicyInput, ...request.Option) (*ModifyNatFirewallControlPolicyOutput, error)
	ModifyNatFirewallControlPolicyRequest(*ModifyNatFirewallControlPolicyInput) (*request.Request, *ModifyNatFirewallControlPolicyOutput)

	ModifyNatFirewallControlPolicyPositionCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ModifyNatFirewallControlPolicyPositionCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifyNatFirewallControlPolicyPositionCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifyNatFirewallControlPolicyPosition(*ModifyNatFirewallControlPolicyPositionInput) (*ModifyNatFirewallControlPolicyPositionOutput, error)
	ModifyNatFirewallControlPolicyPositionWithContext(volcengine.Context, *ModifyNatFirewallControlPolicyPositionInput, ...request.Option) (*ModifyNatFirewallControlPolicyPositionOutput, error)
	ModifyNatFirewallControlPolicyPositionRequest(*ModifyNatFirewallControlPolicyPositionInput) (*request.Request, *ModifyNatFirewallControlPolicyPositionOutput)

	ModifyVpcFirewallAclRuleCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ModifyVpcFirewallAclRuleCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifyVpcFirewallAclRuleCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifyVpcFirewallAclRule(*ModifyVpcFirewallAclRuleInput) (*ModifyVpcFirewallAclRuleOutput, error)
	ModifyVpcFirewallAclRuleWithContext(volcengine.Context, *ModifyVpcFirewallAclRuleInput, ...request.Option) (*ModifyVpcFirewallAclRuleOutput, error)
	ModifyVpcFirewallAclRuleRequest(*ModifyVpcFirewallAclRuleInput) (*request.Request, *ModifyVpcFirewallAclRuleOutput)

	ModifyVpcFirewallAclRulePositionCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ModifyVpcFirewallAclRulePositionCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifyVpcFirewallAclRulePositionCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifyVpcFirewallAclRulePosition(*ModifyVpcFirewallAclRulePositionInput) (*ModifyVpcFirewallAclRulePositionOutput, error)
	ModifyVpcFirewallAclRulePositionWithContext(volcengine.Context, *ModifyVpcFirewallAclRulePositionInput, ...request.Option) (*ModifyVpcFirewallAclRulePositionOutput, error)
	ModifyVpcFirewallAclRulePositionRequest(*ModifyVpcFirewallAclRulePositionInput) (*request.Request, *ModifyVpcFirewallAclRulePositionOutput)

	UpdateAssetSwitchCommon(*map[string]interface{}) (*map[string]interface{}, error)
	UpdateAssetSwitchCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	UpdateAssetSwitchCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	UpdateAssetSwitch(*UpdateAssetSwitchInput) (*UpdateAssetSwitchOutput, error)
	UpdateAssetSwitchWithContext(volcengine.Context, *UpdateAssetSwitchInput, ...request.Option) (*UpdateAssetSwitchOutput, error)
	UpdateAssetSwitchRequest(*UpdateAssetSwitchInput) (*request.Request, *UpdateAssetSwitchOutput)

	UpdateAssetsCommon(*map[string]interface{}) (*map[string]interface{}, error)
	UpdateAssetsCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	UpdateAssetsCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	UpdateAssets(*UpdateAssetsInput) (*UpdateAssetsOutput, error)
	UpdateAssetsWithContext(volcengine.Context, *UpdateAssetsInput, ...request.Option) (*UpdateAssetsOutput, error)
	UpdateAssetsRequest(*UpdateAssetsInput) (*request.Request, *UpdateAssetsOutput)

	UpdateControlPolicySwitchCommon(*map[string]interface{}) (*map[string]interface{}, error)
	UpdateControlPolicySwitchCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	UpdateControlPolicySwitchCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	UpdateControlPolicySwitch(*UpdateControlPolicySwitchInput) (*UpdateControlPolicySwitchOutput, error)
	UpdateControlPolicySwitchWithContext(volcengine.Context, *UpdateControlPolicySwitchInput, ...request.Option) (*UpdateControlPolicySwitchOutput, error)
	UpdateControlPolicySwitchRequest(*UpdateControlPolicySwitchInput) (*request.Request, *UpdateControlPolicySwitchOutput)

	UpdateDnsControlPolicySwitchCommon(*map[string]interface{}) (*map[string]interface{}, error)
	UpdateDnsControlPolicySwitchCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	UpdateDnsControlPolicySwitchCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	UpdateDnsControlPolicySwitch(*UpdateDnsControlPolicySwitchInput) (*UpdateDnsControlPolicySwitchOutput, error)
	UpdateDnsControlPolicySwitchWithContext(volcengine.Context, *UpdateDnsControlPolicySwitchInput, ...request.Option) (*UpdateDnsControlPolicySwitchOutput, error)
	UpdateDnsControlPolicySwitchRequest(*UpdateDnsControlPolicySwitchInput) (*request.Request, *UpdateDnsControlPolicySwitchOutput)

	UpdateNatFirewallControlPolicySwitchCommon(*map[string]interface{}) (*map[string]interface{}, error)
	UpdateNatFirewallControlPolicySwitchCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	UpdateNatFirewallControlPolicySwitchCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	UpdateNatFirewallControlPolicySwitch(*UpdateNatFirewallControlPolicySwitchInput) (*UpdateNatFirewallControlPolicySwitchOutput, error)
	UpdateNatFirewallControlPolicySwitchWithContext(volcengine.Context, *UpdateNatFirewallControlPolicySwitchInput, ...request.Option) (*UpdateNatFirewallControlPolicySwitchOutput, error)
	UpdateNatFirewallControlPolicySwitchRequest(*UpdateNatFirewallControlPolicySwitchInput) (*request.Request, *UpdateNatFirewallControlPolicySwitchOutput)

	UpdateVpcFirewallAclRuleSwitchCommon(*map[string]interface{}) (*map[string]interface{}, error)
	UpdateVpcFirewallAclRuleSwitchCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	UpdateVpcFirewallAclRuleSwitchCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	UpdateVpcFirewallAclRuleSwitch(*UpdateVpcFirewallAclRuleSwitchInput) (*UpdateVpcFirewallAclRuleSwitchOutput, error)
	UpdateVpcFirewallAclRuleSwitchWithContext(volcengine.Context, *UpdateVpcFirewallAclRuleSwitchInput, ...request.Option) (*UpdateVpcFirewallAclRuleSwitchOutput, error)
	UpdateVpcFirewallAclRuleSwitchRequest(*UpdateVpcFirewallAclRuleSwitchInput) (*request.Request, *UpdateVpcFirewallAclRuleSwitchOutput)
}

var _ FWCENTERAPI = (*FWCENTER)(nil)

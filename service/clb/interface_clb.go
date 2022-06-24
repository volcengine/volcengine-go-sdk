// Code generated by volcstack with private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package clbiface provides an interface to enable mocking the CLB service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package clb

import (
	"github.com/volcengine/volcstack-go-sdk/volcstack"
	"github.com/volcengine/volcstack-go-sdk/volcstack/request"
)

// CLBAPI provides an interface to enable mocking the
// clb.CLB service client's API operation,
//
//    // volcstack sdk func uses an SDK service client to make a request to
//    // CLB.
//    func myFunc(svc CLBAPI) bool {
//        // Make svc.AddAclEntries request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := clb.New(sess)
//
//        myFunc(svc)
//    }
//
type CLBAPI interface {
	AddAclEntriesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	AddAclEntriesCommonWithContext(volcstack.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	AddAclEntriesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	AddAclEntries(*AddAclEntriesInput) (*AddAclEntriesOutput, error)
	AddAclEntriesWithContext(volcstack.Context, *AddAclEntriesInput, ...request.Option) (*AddAclEntriesOutput, error)
	AddAclEntriesRequest(*AddAclEntriesInput) (*request.Request, *AddAclEntriesOutput)

	AddServerGroupBackendServersCommon(*map[string]interface{}) (*map[string]interface{}, error)
	AddServerGroupBackendServersCommonWithContext(volcstack.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	AddServerGroupBackendServersCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	AddServerGroupBackendServers(*AddServerGroupBackendServersInput) (*AddServerGroupBackendServersOutput, error)
	AddServerGroupBackendServersWithContext(volcstack.Context, *AddServerGroupBackendServersInput, ...request.Option) (*AddServerGroupBackendServersOutput, error)
	AddServerGroupBackendServersRequest(*AddServerGroupBackendServersInput) (*request.Request, *AddServerGroupBackendServersOutput)

	ConvertLoadBalancerBillingTypeCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ConvertLoadBalancerBillingTypeCommonWithContext(volcstack.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ConvertLoadBalancerBillingTypeCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ConvertLoadBalancerBillingType(*ConvertLoadBalancerBillingTypeInput) (*ConvertLoadBalancerBillingTypeOutput, error)
	ConvertLoadBalancerBillingTypeWithContext(volcstack.Context, *ConvertLoadBalancerBillingTypeInput, ...request.Option) (*ConvertLoadBalancerBillingTypeOutput, error)
	ConvertLoadBalancerBillingTypeRequest(*ConvertLoadBalancerBillingTypeInput) (*request.Request, *ConvertLoadBalancerBillingTypeOutput)

	CreateAclCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CreateAclCommonWithContext(volcstack.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateAclCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateAcl(*CreateAclInput) (*CreateAclOutput, error)
	CreateAclWithContext(volcstack.Context, *CreateAclInput, ...request.Option) (*CreateAclOutput, error)
	CreateAclRequest(*CreateAclInput) (*request.Request, *CreateAclOutput)

	CreateListenerCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CreateListenerCommonWithContext(volcstack.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateListenerCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateListener(*CreateListenerInput) (*CreateListenerOutput, error)
	CreateListenerWithContext(volcstack.Context, *CreateListenerInput, ...request.Option) (*CreateListenerOutput, error)
	CreateListenerRequest(*CreateListenerInput) (*request.Request, *CreateListenerOutput)

	CreateLoadBalancerCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CreateLoadBalancerCommonWithContext(volcstack.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateLoadBalancerCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateLoadBalancer(*CreateLoadBalancerInput) (*CreateLoadBalancerOutput, error)
	CreateLoadBalancerWithContext(volcstack.Context, *CreateLoadBalancerInput, ...request.Option) (*CreateLoadBalancerOutput, error)
	CreateLoadBalancerRequest(*CreateLoadBalancerInput) (*request.Request, *CreateLoadBalancerOutput)

	CreateRulesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CreateRulesCommonWithContext(volcstack.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateRulesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateRules(*CreateRulesInput) (*CreateRulesOutput, error)
	CreateRulesWithContext(volcstack.Context, *CreateRulesInput, ...request.Option) (*CreateRulesOutput, error)
	CreateRulesRequest(*CreateRulesInput) (*request.Request, *CreateRulesOutput)

	CreateServerGroupCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CreateServerGroupCommonWithContext(volcstack.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateServerGroupCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateServerGroup(*CreateServerGroupInput) (*CreateServerGroupOutput, error)
	CreateServerGroupWithContext(volcstack.Context, *CreateServerGroupInput, ...request.Option) (*CreateServerGroupOutput, error)
	CreateServerGroupRequest(*CreateServerGroupInput) (*request.Request, *CreateServerGroupOutput)

	DeleteAclCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteAclCommonWithContext(volcstack.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteAclCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteAcl(*DeleteAclInput) (*DeleteAclOutput, error)
	DeleteAclWithContext(volcstack.Context, *DeleteAclInput, ...request.Option) (*DeleteAclOutput, error)
	DeleteAclRequest(*DeleteAclInput) (*request.Request, *DeleteAclOutput)

	DeleteCertificateCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteCertificateCommonWithContext(volcstack.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteCertificateCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteCertificate(*DeleteCertificateInput) (*DeleteCertificateOutput, error)
	DeleteCertificateWithContext(volcstack.Context, *DeleteCertificateInput, ...request.Option) (*DeleteCertificateOutput, error)
	DeleteCertificateRequest(*DeleteCertificateInput) (*request.Request, *DeleteCertificateOutput)

	DeleteListenerCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteListenerCommonWithContext(volcstack.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteListenerCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteListener(*DeleteListenerInput) (*DeleteListenerOutput, error)
	DeleteListenerWithContext(volcstack.Context, *DeleteListenerInput, ...request.Option) (*DeleteListenerOutput, error)
	DeleteListenerRequest(*DeleteListenerInput) (*request.Request, *DeleteListenerOutput)

	DeleteLoadBalancerCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteLoadBalancerCommonWithContext(volcstack.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteLoadBalancerCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteLoadBalancer(*DeleteLoadBalancerInput) (*DeleteLoadBalancerOutput, error)
	DeleteLoadBalancerWithContext(volcstack.Context, *DeleteLoadBalancerInput, ...request.Option) (*DeleteLoadBalancerOutput, error)
	DeleteLoadBalancerRequest(*DeleteLoadBalancerInput) (*request.Request, *DeleteLoadBalancerOutput)

	DeleteRulesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteRulesCommonWithContext(volcstack.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteRulesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteRules(*DeleteRulesInput) (*DeleteRulesOutput, error)
	DeleteRulesWithContext(volcstack.Context, *DeleteRulesInput, ...request.Option) (*DeleteRulesOutput, error)
	DeleteRulesRequest(*DeleteRulesInput) (*request.Request, *DeleteRulesOutput)

	DeleteServerGroupCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteServerGroupCommonWithContext(volcstack.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteServerGroupCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteServerGroup(*DeleteServerGroupInput) (*DeleteServerGroupOutput, error)
	DeleteServerGroupWithContext(volcstack.Context, *DeleteServerGroupInput, ...request.Option) (*DeleteServerGroupOutput, error)
	DeleteServerGroupRequest(*DeleteServerGroupInput) (*request.Request, *DeleteServerGroupOutput)

	DescribeAclAttributesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeAclAttributesCommonWithContext(volcstack.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeAclAttributesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeAclAttributes(*DescribeAclAttributesInput) (*DescribeAclAttributesOutput, error)
	DescribeAclAttributesWithContext(volcstack.Context, *DescribeAclAttributesInput, ...request.Option) (*DescribeAclAttributesOutput, error)
	DescribeAclAttributesRequest(*DescribeAclAttributesInput) (*request.Request, *DescribeAclAttributesOutput)

	DescribeAclsCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeAclsCommonWithContext(volcstack.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeAclsCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeAcls(*DescribeAclsInput) (*DescribeAclsOutput, error)
	DescribeAclsWithContext(volcstack.Context, *DescribeAclsInput, ...request.Option) (*DescribeAclsOutput, error)
	DescribeAclsRequest(*DescribeAclsInput) (*request.Request, *DescribeAclsOutput)

	DescribeCertificatesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeCertificatesCommonWithContext(volcstack.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeCertificatesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeCertificates(*DescribeCertificatesInput) (*DescribeCertificatesOutput, error)
	DescribeCertificatesWithContext(volcstack.Context, *DescribeCertificatesInput, ...request.Option) (*DescribeCertificatesOutput, error)
	DescribeCertificatesRequest(*DescribeCertificatesInput) (*request.Request, *DescribeCertificatesOutput)

	DescribeListenerAttributesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeListenerAttributesCommonWithContext(volcstack.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeListenerAttributesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeListenerAttributes(*DescribeListenerAttributesInput) (*DescribeListenerAttributesOutput, error)
	DescribeListenerAttributesWithContext(volcstack.Context, *DescribeListenerAttributesInput, ...request.Option) (*DescribeListenerAttributesOutput, error)
	DescribeListenerAttributesRequest(*DescribeListenerAttributesInput) (*request.Request, *DescribeListenerAttributesOutput)

	DescribeListenerHealthCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeListenerHealthCommonWithContext(volcstack.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeListenerHealthCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeListenerHealth(*DescribeListenerHealthInput) (*DescribeListenerHealthOutput, error)
	DescribeListenerHealthWithContext(volcstack.Context, *DescribeListenerHealthInput, ...request.Option) (*DescribeListenerHealthOutput, error)
	DescribeListenerHealthRequest(*DescribeListenerHealthInput) (*request.Request, *DescribeListenerHealthOutput)

	DescribeListenersCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeListenersCommonWithContext(volcstack.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeListenersCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeListeners(*DescribeListenersInput) (*DescribeListenersOutput, error)
	DescribeListenersWithContext(volcstack.Context, *DescribeListenersInput, ...request.Option) (*DescribeListenersOutput, error)
	DescribeListenersRequest(*DescribeListenersInput) (*request.Request, *DescribeListenersOutput)

	DescribeLoadBalancerAttributesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeLoadBalancerAttributesCommonWithContext(volcstack.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeLoadBalancerAttributesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeLoadBalancerAttributes(*DescribeLoadBalancerAttributesInput) (*DescribeLoadBalancerAttributesOutput, error)
	DescribeLoadBalancerAttributesWithContext(volcstack.Context, *DescribeLoadBalancerAttributesInput, ...request.Option) (*DescribeLoadBalancerAttributesOutput, error)
	DescribeLoadBalancerAttributesRequest(*DescribeLoadBalancerAttributesInput) (*request.Request, *DescribeLoadBalancerAttributesOutput)

	DescribeLoadBalancersCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeLoadBalancersCommonWithContext(volcstack.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeLoadBalancersCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeLoadBalancers(*DescribeLoadBalancersInput) (*DescribeLoadBalancersOutput, error)
	DescribeLoadBalancersWithContext(volcstack.Context, *DescribeLoadBalancersInput, ...request.Option) (*DescribeLoadBalancersOutput, error)
	DescribeLoadBalancersRequest(*DescribeLoadBalancersInput) (*request.Request, *DescribeLoadBalancersOutput)

	DescribeLoadBalancersBillingCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeLoadBalancersBillingCommonWithContext(volcstack.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeLoadBalancersBillingCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeLoadBalancersBilling(*DescribeLoadBalancersBillingInput) (*DescribeLoadBalancersBillingOutput, error)
	DescribeLoadBalancersBillingWithContext(volcstack.Context, *DescribeLoadBalancersBillingInput, ...request.Option) (*DescribeLoadBalancersBillingOutput, error)
	DescribeLoadBalancersBillingRequest(*DescribeLoadBalancersBillingInput) (*request.Request, *DescribeLoadBalancersBillingOutput)

	DescribeRulesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeRulesCommonWithContext(volcstack.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeRulesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeRules(*DescribeRulesInput) (*DescribeRulesOutput, error)
	DescribeRulesWithContext(volcstack.Context, *DescribeRulesInput, ...request.Option) (*DescribeRulesOutput, error)
	DescribeRulesRequest(*DescribeRulesInput) (*request.Request, *DescribeRulesOutput)

	DescribeServerGroupAttributesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeServerGroupAttributesCommonWithContext(volcstack.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeServerGroupAttributesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeServerGroupAttributes(*DescribeServerGroupAttributesInput) (*DescribeServerGroupAttributesOutput, error)
	DescribeServerGroupAttributesWithContext(volcstack.Context, *DescribeServerGroupAttributesInput, ...request.Option) (*DescribeServerGroupAttributesOutput, error)
	DescribeServerGroupAttributesRequest(*DescribeServerGroupAttributesInput) (*request.Request, *DescribeServerGroupAttributesOutput)

	DescribeServerGroupsCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeServerGroupsCommonWithContext(volcstack.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeServerGroupsCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeServerGroups(*DescribeServerGroupsInput) (*DescribeServerGroupsOutput, error)
	DescribeServerGroupsWithContext(volcstack.Context, *DescribeServerGroupsInput, ...request.Option) (*DescribeServerGroupsOutput, error)
	DescribeServerGroupsRequest(*DescribeServerGroupsInput) (*request.Request, *DescribeServerGroupsOutput)

	DisableAccessLogCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DisableAccessLogCommonWithContext(volcstack.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DisableAccessLogCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DisableAccessLog(*DisableAccessLogInput) (*DisableAccessLogOutput, error)
	DisableAccessLogWithContext(volcstack.Context, *DisableAccessLogInput, ...request.Option) (*DisableAccessLogOutput, error)
	DisableAccessLogRequest(*DisableAccessLogInput) (*request.Request, *DisableAccessLogOutput)

	EnableAccessLogCommon(*map[string]interface{}) (*map[string]interface{}, error)
	EnableAccessLogCommonWithContext(volcstack.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	EnableAccessLogCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	EnableAccessLog(*EnableAccessLogInput) (*EnableAccessLogOutput, error)
	EnableAccessLogWithContext(volcstack.Context, *EnableAccessLogInput, ...request.Option) (*EnableAccessLogOutput, error)
	EnableAccessLogRequest(*EnableAccessLogInput) (*request.Request, *EnableAccessLogOutput)

	ModifyAclAttributesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ModifyAclAttributesCommonWithContext(volcstack.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifyAclAttributesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifyAclAttributes(*ModifyAclAttributesInput) (*ModifyAclAttributesOutput, error)
	ModifyAclAttributesWithContext(volcstack.Context, *ModifyAclAttributesInput, ...request.Option) (*ModifyAclAttributesOutput, error)
	ModifyAclAttributesRequest(*ModifyAclAttributesInput) (*request.Request, *ModifyAclAttributesOutput)

	ModifyListenerAttributesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ModifyListenerAttributesCommonWithContext(volcstack.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifyListenerAttributesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifyListenerAttributes(*ModifyListenerAttributesInput) (*ModifyListenerAttributesOutput, error)
	ModifyListenerAttributesWithContext(volcstack.Context, *ModifyListenerAttributesInput, ...request.Option) (*ModifyListenerAttributesOutput, error)
	ModifyListenerAttributesRequest(*ModifyListenerAttributesInput) (*request.Request, *ModifyListenerAttributesOutput)

	ModifyLoadBalancerAttributesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ModifyLoadBalancerAttributesCommonWithContext(volcstack.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifyLoadBalancerAttributesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifyLoadBalancerAttributes(*ModifyLoadBalancerAttributesInput) (*ModifyLoadBalancerAttributesOutput, error)
	ModifyLoadBalancerAttributesWithContext(volcstack.Context, *ModifyLoadBalancerAttributesInput, ...request.Option) (*ModifyLoadBalancerAttributesOutput, error)
	ModifyLoadBalancerAttributesRequest(*ModifyLoadBalancerAttributesInput) (*request.Request, *ModifyLoadBalancerAttributesOutput)

	ModifyRulesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ModifyRulesCommonWithContext(volcstack.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifyRulesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifyRules(*ModifyRulesInput) (*ModifyRulesOutput, error)
	ModifyRulesWithContext(volcstack.Context, *ModifyRulesInput, ...request.Option) (*ModifyRulesOutput, error)
	ModifyRulesRequest(*ModifyRulesInput) (*request.Request, *ModifyRulesOutput)

	ModifyServerGroupAttributesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ModifyServerGroupAttributesCommonWithContext(volcstack.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifyServerGroupAttributesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifyServerGroupAttributes(*ModifyServerGroupAttributesInput) (*ModifyServerGroupAttributesOutput, error)
	ModifyServerGroupAttributesWithContext(volcstack.Context, *ModifyServerGroupAttributesInput, ...request.Option) (*ModifyServerGroupAttributesOutput, error)
	ModifyServerGroupAttributesRequest(*ModifyServerGroupAttributesInput) (*request.Request, *ModifyServerGroupAttributesOutput)

	RemoveAclEntriesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	RemoveAclEntriesCommonWithContext(volcstack.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	RemoveAclEntriesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	RemoveAclEntries(*RemoveAclEntriesInput) (*RemoveAclEntriesOutput, error)
	RemoveAclEntriesWithContext(volcstack.Context, *RemoveAclEntriesInput, ...request.Option) (*RemoveAclEntriesOutput, error)
	RemoveAclEntriesRequest(*RemoveAclEntriesInput) (*request.Request, *RemoveAclEntriesOutput)

	RemoveServerGroupBackendServersCommon(*map[string]interface{}) (*map[string]interface{}, error)
	RemoveServerGroupBackendServersCommonWithContext(volcstack.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	RemoveServerGroupBackendServersCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	RemoveServerGroupBackendServers(*RemoveServerGroupBackendServersInput) (*RemoveServerGroupBackendServersOutput, error)
	RemoveServerGroupBackendServersWithContext(volcstack.Context, *RemoveServerGroupBackendServersInput, ...request.Option) (*RemoveServerGroupBackendServersOutput, error)
	RemoveServerGroupBackendServersRequest(*RemoveServerGroupBackendServersInput) (*request.Request, *RemoveServerGroupBackendServersOutput)

	RenewLoadBalancerCommon(*map[string]interface{}) (*map[string]interface{}, error)
	RenewLoadBalancerCommonWithContext(volcstack.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	RenewLoadBalancerCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	RenewLoadBalancer(*RenewLoadBalancerInput) (*RenewLoadBalancerOutput, error)
	RenewLoadBalancerWithContext(volcstack.Context, *RenewLoadBalancerInput, ...request.Option) (*RenewLoadBalancerOutput, error)
	RenewLoadBalancerRequest(*RenewLoadBalancerInput) (*request.Request, *RenewLoadBalancerOutput)

	SetLoadBalancerRenewalCommon(*map[string]interface{}) (*map[string]interface{}, error)
	SetLoadBalancerRenewalCommonWithContext(volcstack.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	SetLoadBalancerRenewalCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	SetLoadBalancerRenewal(*SetLoadBalancerRenewalInput) (*SetLoadBalancerRenewalOutput, error)
	SetLoadBalancerRenewalWithContext(volcstack.Context, *SetLoadBalancerRenewalInput, ...request.Option) (*SetLoadBalancerRenewalOutput, error)
	SetLoadBalancerRenewalRequest(*SetLoadBalancerRenewalInput) (*request.Request, *SetLoadBalancerRenewalOutput)

	UploadCertificateCommon(*map[string]interface{}) (*map[string]interface{}, error)
	UploadCertificateCommonWithContext(volcstack.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	UploadCertificateCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	UploadCertificate(*UploadCertificateInput) (*UploadCertificateOutput, error)
	UploadCertificateWithContext(volcstack.Context, *UploadCertificateInput, ...request.Option) (*UploadCertificateOutput, error)
	UploadCertificateRequest(*UploadCertificateInput) (*request.Request, *UploadCertificateOutput)
}

var _ CLBAPI = (*CLB)(nil)
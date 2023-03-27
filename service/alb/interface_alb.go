// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package albiface provides an interface to enable mocking the ALB service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package alb

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
)

// ALBAPI provides an interface to enable mocking the
// alb.ALB service client's API operation,
//
//    // volcengine sdk func uses an SDK service client to make a request to
//    // ALB.
//    func myFunc(svc ALBAPI) bool {
//        // Make svc.AddAclEntries request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := alb.New(sess)
//
//        myFunc(svc)
//    }
//
type ALBAPI interface {
	AddAclEntriesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	AddAclEntriesCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	AddAclEntriesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	AddAclEntries(*AddAclEntriesInput) (*AddAclEntriesOutput, error)
	AddAclEntriesWithContext(volcengine.Context, *AddAclEntriesInput, ...request.Option) (*AddAclEntriesOutput, error)
	AddAclEntriesRequest(*AddAclEntriesInput) (*request.Request, *AddAclEntriesOutput)

	AddServerGroupBackendServersCommon(*map[string]interface{}) (*map[string]interface{}, error)
	AddServerGroupBackendServersCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	AddServerGroupBackendServersCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	AddServerGroupBackendServers(*AddServerGroupBackendServersInput) (*AddServerGroupBackendServersOutput, error)
	AddServerGroupBackendServersWithContext(volcengine.Context, *AddServerGroupBackendServersInput, ...request.Option) (*AddServerGroupBackendServersOutput, error)
	AddServerGroupBackendServersRequest(*AddServerGroupBackendServersInput) (*request.Request, *AddServerGroupBackendServersOutput)

	CreateAclCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CreateAclCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateAclCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateAcl(*CreateAclInput) (*CreateAclOutput, error)
	CreateAclWithContext(volcengine.Context, *CreateAclInput, ...request.Option) (*CreateAclOutput, error)
	CreateAclRequest(*CreateAclInput) (*request.Request, *CreateAclOutput)

	CreateCustomizedCfgCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CreateCustomizedCfgCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateCustomizedCfgCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateCustomizedCfg(*CreateCustomizedCfgInput) (*CreateCustomizedCfgOutput, error)
	CreateCustomizedCfgWithContext(volcengine.Context, *CreateCustomizedCfgInput, ...request.Option) (*CreateCustomizedCfgOutput, error)
	CreateCustomizedCfgRequest(*CreateCustomizedCfgInput) (*request.Request, *CreateCustomizedCfgOutput)

	CreateHealthCheckTemplatesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CreateHealthCheckTemplatesCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateHealthCheckTemplatesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateHealthCheckTemplates(*CreateHealthCheckTemplatesInput) (*CreateHealthCheckTemplatesOutput, error)
	CreateHealthCheckTemplatesWithContext(volcengine.Context, *CreateHealthCheckTemplatesInput, ...request.Option) (*CreateHealthCheckTemplatesOutput, error)
	CreateHealthCheckTemplatesRequest(*CreateHealthCheckTemplatesInput) (*request.Request, *CreateHealthCheckTemplatesOutput)

	CreateListenerCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CreateListenerCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateListenerCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateListener(*CreateListenerInput) (*CreateListenerOutput, error)
	CreateListenerWithContext(volcengine.Context, *CreateListenerInput, ...request.Option) (*CreateListenerOutput, error)
	CreateListenerRequest(*CreateListenerInput) (*request.Request, *CreateListenerOutput)

	CreateLoadBalancerCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CreateLoadBalancerCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateLoadBalancerCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateLoadBalancer(*CreateLoadBalancerInput) (*CreateLoadBalancerOutput, error)
	CreateLoadBalancerWithContext(volcengine.Context, *CreateLoadBalancerInput, ...request.Option) (*CreateLoadBalancerOutput, error)
	CreateLoadBalancerRequest(*CreateLoadBalancerInput) (*request.Request, *CreateLoadBalancerOutput)

	CreateRulesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CreateRulesCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateRulesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateRules(*CreateRulesInput) (*CreateRulesOutput, error)
	CreateRulesWithContext(volcengine.Context, *CreateRulesInput, ...request.Option) (*CreateRulesOutput, error)
	CreateRulesRequest(*CreateRulesInput) (*request.Request, *CreateRulesOutput)

	CreateServerGroupCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CreateServerGroupCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateServerGroupCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateServerGroup(*CreateServerGroupInput) (*CreateServerGroupOutput, error)
	CreateServerGroupWithContext(volcengine.Context, *CreateServerGroupInput, ...request.Option) (*CreateServerGroupOutput, error)
	CreateServerGroupRequest(*CreateServerGroupInput) (*request.Request, *CreateServerGroupOutput)

	DeleteAclCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteAclCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteAclCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteAcl(*DeleteAclInput) (*DeleteAclOutput, error)
	DeleteAclWithContext(volcengine.Context, *DeleteAclInput, ...request.Option) (*DeleteAclOutput, error)
	DeleteAclRequest(*DeleteAclInput) (*request.Request, *DeleteAclOutput)

	DeleteCACertificateCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteCACertificateCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteCACertificateCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteCACertificate(*DeleteCACertificateInput) (*DeleteCACertificateOutput, error)
	DeleteCACertificateWithContext(volcengine.Context, *DeleteCACertificateInput, ...request.Option) (*DeleteCACertificateOutput, error)
	DeleteCACertificateRequest(*DeleteCACertificateInput) (*request.Request, *DeleteCACertificateOutput)

	DeleteCertificateCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteCertificateCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteCertificateCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteCertificate(*DeleteCertificateInput) (*DeleteCertificateOutput, error)
	DeleteCertificateWithContext(volcengine.Context, *DeleteCertificateInput, ...request.Option) (*DeleteCertificateOutput, error)
	DeleteCertificateRequest(*DeleteCertificateInput) (*request.Request, *DeleteCertificateOutput)

	DeleteCustomizedCfgCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteCustomizedCfgCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteCustomizedCfgCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteCustomizedCfg(*DeleteCustomizedCfgInput) (*DeleteCustomizedCfgOutput, error)
	DeleteCustomizedCfgWithContext(volcengine.Context, *DeleteCustomizedCfgInput, ...request.Option) (*DeleteCustomizedCfgOutput, error)
	DeleteCustomizedCfgRequest(*DeleteCustomizedCfgInput) (*request.Request, *DeleteCustomizedCfgOutput)

	DeleteHealthCheckTemplatesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteHealthCheckTemplatesCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteHealthCheckTemplatesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteHealthCheckTemplates(*DeleteHealthCheckTemplatesInput) (*DeleteHealthCheckTemplatesOutput, error)
	DeleteHealthCheckTemplatesWithContext(volcengine.Context, *DeleteHealthCheckTemplatesInput, ...request.Option) (*DeleteHealthCheckTemplatesOutput, error)
	DeleteHealthCheckTemplatesRequest(*DeleteHealthCheckTemplatesInput) (*request.Request, *DeleteHealthCheckTemplatesOutput)

	DeleteListenerCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteListenerCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteListenerCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteListener(*DeleteListenerInput) (*DeleteListenerOutput, error)
	DeleteListenerWithContext(volcengine.Context, *DeleteListenerInput, ...request.Option) (*DeleteListenerOutput, error)
	DeleteListenerRequest(*DeleteListenerInput) (*request.Request, *DeleteListenerOutput)

	DeleteLoadBalancerCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteLoadBalancerCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteLoadBalancerCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteLoadBalancer(*DeleteLoadBalancerInput) (*DeleteLoadBalancerOutput, error)
	DeleteLoadBalancerWithContext(volcengine.Context, *DeleteLoadBalancerInput, ...request.Option) (*DeleteLoadBalancerOutput, error)
	DeleteLoadBalancerRequest(*DeleteLoadBalancerInput) (*request.Request, *DeleteLoadBalancerOutput)

	DeleteRulesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteRulesCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteRulesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteRules(*DeleteRulesInput) (*DeleteRulesOutput, error)
	DeleteRulesWithContext(volcengine.Context, *DeleteRulesInput, ...request.Option) (*DeleteRulesOutput, error)
	DeleteRulesRequest(*DeleteRulesInput) (*request.Request, *DeleteRulesOutput)

	DeleteServerGroupCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteServerGroupCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteServerGroupCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteServerGroup(*DeleteServerGroupInput) (*DeleteServerGroupOutput, error)
	DeleteServerGroupWithContext(volcengine.Context, *DeleteServerGroupInput, ...request.Option) (*DeleteServerGroupOutput, error)
	DeleteServerGroupRequest(*DeleteServerGroupInput) (*request.Request, *DeleteServerGroupOutput)

	DescribeAclAttributesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeAclAttributesCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeAclAttributesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeAclAttributes(*DescribeAclAttributesInput) (*DescribeAclAttributesOutput, error)
	DescribeAclAttributesWithContext(volcengine.Context, *DescribeAclAttributesInput, ...request.Option) (*DescribeAclAttributesOutput, error)
	DescribeAclAttributesRequest(*DescribeAclAttributesInput) (*request.Request, *DescribeAclAttributesOutput)

	DescribeAclsCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeAclsCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeAclsCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeAcls(*DescribeAclsInput) (*DescribeAclsOutput, error)
	DescribeAclsWithContext(volcengine.Context, *DescribeAclsInput, ...request.Option) (*DescribeAclsOutput, error)
	DescribeAclsRequest(*DescribeAclsInput) (*request.Request, *DescribeAclsOutput)

	DescribeAllCertificatesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeAllCertificatesCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeAllCertificatesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeAllCertificates(*DescribeAllCertificatesInput) (*DescribeAllCertificatesOutput, error)
	DescribeAllCertificatesWithContext(volcengine.Context, *DescribeAllCertificatesInput, ...request.Option) (*DescribeAllCertificatesOutput, error)
	DescribeAllCertificatesRequest(*DescribeAllCertificatesInput) (*request.Request, *DescribeAllCertificatesOutput)

	DescribeCACertificatesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeCACertificatesCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeCACertificatesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeCACertificates(*DescribeCACertificatesInput) (*DescribeCACertificatesOutput, error)
	DescribeCACertificatesWithContext(volcengine.Context, *DescribeCACertificatesInput, ...request.Option) (*DescribeCACertificatesOutput, error)
	DescribeCACertificatesRequest(*DescribeCACertificatesInput) (*request.Request, *DescribeCACertificatesOutput)

	DescribeCertificatesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeCertificatesCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeCertificatesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeCertificates(*DescribeCertificatesInput) (*DescribeCertificatesOutput, error)
	DescribeCertificatesWithContext(volcengine.Context, *DescribeCertificatesInput, ...request.Option) (*DescribeCertificatesOutput, error)
	DescribeCertificatesRequest(*DescribeCertificatesInput) (*request.Request, *DescribeCertificatesOutput)

	DescribeCustomizedCfgAttributesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeCustomizedCfgAttributesCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeCustomizedCfgAttributesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeCustomizedCfgAttributes(*DescribeCustomizedCfgAttributesInput) (*DescribeCustomizedCfgAttributesOutput, error)
	DescribeCustomizedCfgAttributesWithContext(volcengine.Context, *DescribeCustomizedCfgAttributesInput, ...request.Option) (*DescribeCustomizedCfgAttributesOutput, error)
	DescribeCustomizedCfgAttributesRequest(*DescribeCustomizedCfgAttributesInput) (*request.Request, *DescribeCustomizedCfgAttributesOutput)

	DescribeCustomizedCfgsCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeCustomizedCfgsCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeCustomizedCfgsCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeCustomizedCfgs(*DescribeCustomizedCfgsInput) (*DescribeCustomizedCfgsOutput, error)
	DescribeCustomizedCfgsWithContext(volcengine.Context, *DescribeCustomizedCfgsInput, ...request.Option) (*DescribeCustomizedCfgsOutput, error)
	DescribeCustomizedCfgsRequest(*DescribeCustomizedCfgsInput) (*request.Request, *DescribeCustomizedCfgsOutput)

	DescribeHealthCheckTemplatesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeHealthCheckTemplatesCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeHealthCheckTemplatesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeHealthCheckTemplates(*DescribeHealthCheckTemplatesInput) (*DescribeHealthCheckTemplatesOutput, error)
	DescribeHealthCheckTemplatesWithContext(volcengine.Context, *DescribeHealthCheckTemplatesInput, ...request.Option) (*DescribeHealthCheckTemplatesOutput, error)
	DescribeHealthCheckTemplatesRequest(*DescribeHealthCheckTemplatesInput) (*request.Request, *DescribeHealthCheckTemplatesOutput)

	DescribeListenerAttributesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeListenerAttributesCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeListenerAttributesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeListenerAttributes(*DescribeListenerAttributesInput) (*DescribeListenerAttributesOutput, error)
	DescribeListenerAttributesWithContext(volcengine.Context, *DescribeListenerAttributesInput, ...request.Option) (*DescribeListenerAttributesOutput, error)
	DescribeListenerAttributesRequest(*DescribeListenerAttributesInput) (*request.Request, *DescribeListenerAttributesOutput)

	DescribeListenerHealthCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeListenerHealthCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeListenerHealthCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeListenerHealth(*DescribeListenerHealthInput) (*DescribeListenerHealthOutput, error)
	DescribeListenerHealthWithContext(volcengine.Context, *DescribeListenerHealthInput, ...request.Option) (*DescribeListenerHealthOutput, error)
	DescribeListenerHealthRequest(*DescribeListenerHealthInput) (*request.Request, *DescribeListenerHealthOutput)

	DescribeListenersCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeListenersCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeListenersCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeListeners(*DescribeListenersInput) (*DescribeListenersOutput, error)
	DescribeListenersWithContext(volcengine.Context, *DescribeListenersInput, ...request.Option) (*DescribeListenersOutput, error)
	DescribeListenersRequest(*DescribeListenersInput) (*request.Request, *DescribeListenersOutput)

	DescribeLoadBalancerAttributesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeLoadBalancerAttributesCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeLoadBalancerAttributesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeLoadBalancerAttributes(*DescribeLoadBalancerAttributesInput) (*DescribeLoadBalancerAttributesOutput, error)
	DescribeLoadBalancerAttributesWithContext(volcengine.Context, *DescribeLoadBalancerAttributesInput, ...request.Option) (*DescribeLoadBalancerAttributesOutput, error)
	DescribeLoadBalancerAttributesRequest(*DescribeLoadBalancerAttributesInput) (*request.Request, *DescribeLoadBalancerAttributesOutput)

	DescribeLoadBalancersCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeLoadBalancersCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeLoadBalancersCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeLoadBalancers(*DescribeLoadBalancersInput) (*DescribeLoadBalancersOutput, error)
	DescribeLoadBalancersWithContext(volcengine.Context, *DescribeLoadBalancersInput, ...request.Option) (*DescribeLoadBalancersOutput, error)
	DescribeLoadBalancersRequest(*DescribeLoadBalancersInput) (*request.Request, *DescribeLoadBalancersOutput)

	DescribeRulesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeRulesCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeRulesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeRules(*DescribeRulesInput) (*DescribeRulesOutput, error)
	DescribeRulesWithContext(volcengine.Context, *DescribeRulesInput, ...request.Option) (*DescribeRulesOutput, error)
	DescribeRulesRequest(*DescribeRulesInput) (*request.Request, *DescribeRulesOutput)

	DescribeServerGroupAttributesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeServerGroupAttributesCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeServerGroupAttributesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeServerGroupAttributes(*DescribeServerGroupAttributesInput) (*DescribeServerGroupAttributesOutput, error)
	DescribeServerGroupAttributesWithContext(volcengine.Context, *DescribeServerGroupAttributesInput, ...request.Option) (*DescribeServerGroupAttributesOutput, error)
	DescribeServerGroupAttributesRequest(*DescribeServerGroupAttributesInput) (*request.Request, *DescribeServerGroupAttributesOutput)

	DescribeServerGroupsCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeServerGroupsCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeServerGroupsCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeServerGroups(*DescribeServerGroupsInput) (*DescribeServerGroupsOutput, error)
	DescribeServerGroupsWithContext(volcengine.Context, *DescribeServerGroupsInput, ...request.Option) (*DescribeServerGroupsOutput, error)
	DescribeServerGroupsRequest(*DescribeServerGroupsInput) (*request.Request, *DescribeServerGroupsOutput)

	DescribeZonesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeZonesCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeZonesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeZones(*DescribeZonesInput) (*DescribeZonesOutput, error)
	DescribeZonesWithContext(volcengine.Context, *DescribeZonesInput, ...request.Option) (*DescribeZonesOutput, error)
	DescribeZonesRequest(*DescribeZonesInput) (*request.Request, *DescribeZonesOutput)

	DisableAccessLogCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DisableAccessLogCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DisableAccessLogCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DisableAccessLog(*DisableAccessLogInput) (*DisableAccessLogOutput, error)
	DisableAccessLogWithContext(volcengine.Context, *DisableAccessLogInput, ...request.Option) (*DisableAccessLogOutput, error)
	DisableAccessLogRequest(*DisableAccessLogInput) (*request.Request, *DisableAccessLogOutput)

	DisableHealthLogCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DisableHealthLogCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DisableHealthLogCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DisableHealthLog(*DisableHealthLogInput) (*DisableHealthLogOutput, error)
	DisableHealthLogWithContext(volcengine.Context, *DisableHealthLogInput, ...request.Option) (*DisableHealthLogOutput, error)
	DisableHealthLogRequest(*DisableHealthLogInput) (*request.Request, *DisableHealthLogOutput)

	DisableTLSAccessLogCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DisableTLSAccessLogCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DisableTLSAccessLogCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DisableTLSAccessLog(*DisableTLSAccessLogInput) (*DisableTLSAccessLogOutput, error)
	DisableTLSAccessLogWithContext(volcengine.Context, *DisableTLSAccessLogInput, ...request.Option) (*DisableTLSAccessLogOutput, error)
	DisableTLSAccessLogRequest(*DisableTLSAccessLogInput) (*request.Request, *DisableTLSAccessLogOutput)

	EnableAccessLogCommon(*map[string]interface{}) (*map[string]interface{}, error)
	EnableAccessLogCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	EnableAccessLogCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	EnableAccessLog(*EnableAccessLogInput) (*EnableAccessLogOutput, error)
	EnableAccessLogWithContext(volcengine.Context, *EnableAccessLogInput, ...request.Option) (*EnableAccessLogOutput, error)
	EnableAccessLogRequest(*EnableAccessLogInput) (*request.Request, *EnableAccessLogOutput)

	EnableHealthLogCommon(*map[string]interface{}) (*map[string]interface{}, error)
	EnableHealthLogCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	EnableHealthLogCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	EnableHealthLog(*EnableHealthLogInput) (*EnableHealthLogOutput, error)
	EnableHealthLogWithContext(volcengine.Context, *EnableHealthLogInput, ...request.Option) (*EnableHealthLogOutput, error)
	EnableHealthLogRequest(*EnableHealthLogInput) (*request.Request, *EnableHealthLogOutput)

	EnableTLSAccessLogCommon(*map[string]interface{}) (*map[string]interface{}, error)
	EnableTLSAccessLogCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	EnableTLSAccessLogCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	EnableTLSAccessLog(*EnableTLSAccessLogInput) (*EnableTLSAccessLogOutput, error)
	EnableTLSAccessLogWithContext(volcengine.Context, *EnableTLSAccessLogInput, ...request.Option) (*EnableTLSAccessLogOutput, error)
	EnableTLSAccessLogRequest(*EnableTLSAccessLogInput) (*request.Request, *EnableTLSAccessLogOutput)

	ModifyAclAttributesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ModifyAclAttributesCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifyAclAttributesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifyAclAttributes(*ModifyAclAttributesInput) (*ModifyAclAttributesOutput, error)
	ModifyAclAttributesWithContext(volcengine.Context, *ModifyAclAttributesInput, ...request.Option) (*ModifyAclAttributesOutput, error)
	ModifyAclAttributesRequest(*ModifyAclAttributesInput) (*request.Request, *ModifyAclAttributesOutput)

	ModifyCACertificateAttributesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ModifyCACertificateAttributesCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifyCACertificateAttributesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifyCACertificateAttributes(*ModifyCACertificateAttributesInput) (*ModifyCACertificateAttributesOutput, error)
	ModifyCACertificateAttributesWithContext(volcengine.Context, *ModifyCACertificateAttributesInput, ...request.Option) (*ModifyCACertificateAttributesOutput, error)
	ModifyCACertificateAttributesRequest(*ModifyCACertificateAttributesInput) (*request.Request, *ModifyCACertificateAttributesOutput)

	ModifyCertificateAttributesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ModifyCertificateAttributesCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifyCertificateAttributesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifyCertificateAttributes(*ModifyCertificateAttributesInput) (*ModifyCertificateAttributesOutput, error)
	ModifyCertificateAttributesWithContext(volcengine.Context, *ModifyCertificateAttributesInput, ...request.Option) (*ModifyCertificateAttributesOutput, error)
	ModifyCertificateAttributesRequest(*ModifyCertificateAttributesInput) (*request.Request, *ModifyCertificateAttributesOutput)

	ModifyCustomizedCfgAttributesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ModifyCustomizedCfgAttributesCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifyCustomizedCfgAttributesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifyCustomizedCfgAttributes(*ModifyCustomizedCfgAttributesInput) (*ModifyCustomizedCfgAttributesOutput, error)
	ModifyCustomizedCfgAttributesWithContext(volcengine.Context, *ModifyCustomizedCfgAttributesInput, ...request.Option) (*ModifyCustomizedCfgAttributesOutput, error)
	ModifyCustomizedCfgAttributesRequest(*ModifyCustomizedCfgAttributesInput) (*request.Request, *ModifyCustomizedCfgAttributesOutput)

	ModifyHealthCheckTemplatesAttributesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ModifyHealthCheckTemplatesAttributesCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifyHealthCheckTemplatesAttributesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifyHealthCheckTemplatesAttributes(*ModifyHealthCheckTemplatesAttributesInput) (*ModifyHealthCheckTemplatesAttributesOutput, error)
	ModifyHealthCheckTemplatesAttributesWithContext(volcengine.Context, *ModifyHealthCheckTemplatesAttributesInput, ...request.Option) (*ModifyHealthCheckTemplatesAttributesOutput, error)
	ModifyHealthCheckTemplatesAttributesRequest(*ModifyHealthCheckTemplatesAttributesInput) (*request.Request, *ModifyHealthCheckTemplatesAttributesOutput)

	ModifyListenerAttributesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ModifyListenerAttributesCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifyListenerAttributesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifyListenerAttributes(*ModifyListenerAttributesInput) (*ModifyListenerAttributesOutput, error)
	ModifyListenerAttributesWithContext(volcengine.Context, *ModifyListenerAttributesInput, ...request.Option) (*ModifyListenerAttributesOutput, error)
	ModifyListenerAttributesRequest(*ModifyListenerAttributesInput) (*request.Request, *ModifyListenerAttributesOutput)

	ModifyLoadBalancerAttributesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ModifyLoadBalancerAttributesCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifyLoadBalancerAttributesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifyLoadBalancerAttributes(*ModifyLoadBalancerAttributesInput) (*ModifyLoadBalancerAttributesOutput, error)
	ModifyLoadBalancerAttributesWithContext(volcengine.Context, *ModifyLoadBalancerAttributesInput, ...request.Option) (*ModifyLoadBalancerAttributesOutput, error)
	ModifyLoadBalancerAttributesRequest(*ModifyLoadBalancerAttributesInput) (*request.Request, *ModifyLoadBalancerAttributesOutput)

	ModifyLoadBalancerTypeCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ModifyLoadBalancerTypeCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifyLoadBalancerTypeCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifyLoadBalancerType(*ModifyLoadBalancerTypeInput) (*ModifyLoadBalancerTypeOutput, error)
	ModifyLoadBalancerTypeWithContext(volcengine.Context, *ModifyLoadBalancerTypeInput, ...request.Option) (*ModifyLoadBalancerTypeOutput, error)
	ModifyLoadBalancerTypeRequest(*ModifyLoadBalancerTypeInput) (*request.Request, *ModifyLoadBalancerTypeOutput)

	ModifyRulesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ModifyRulesCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifyRulesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifyRules(*ModifyRulesInput) (*ModifyRulesOutput, error)
	ModifyRulesWithContext(volcengine.Context, *ModifyRulesInput, ...request.Option) (*ModifyRulesOutput, error)
	ModifyRulesRequest(*ModifyRulesInput) (*request.Request, *ModifyRulesOutput)

	ModifyServerGroupAttributesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ModifyServerGroupAttributesCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifyServerGroupAttributesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifyServerGroupAttributes(*ModifyServerGroupAttributesInput) (*ModifyServerGroupAttributesOutput, error)
	ModifyServerGroupAttributesWithContext(volcengine.Context, *ModifyServerGroupAttributesInput, ...request.Option) (*ModifyServerGroupAttributesOutput, error)
	ModifyServerGroupAttributesRequest(*ModifyServerGroupAttributesInput) (*request.Request, *ModifyServerGroupAttributesOutput)

	ModifyServerGroupBackendServersCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ModifyServerGroupBackendServersCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifyServerGroupBackendServersCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifyServerGroupBackendServers(*ModifyServerGroupBackendServersInput) (*ModifyServerGroupBackendServersOutput, error)
	ModifyServerGroupBackendServersWithContext(volcengine.Context, *ModifyServerGroupBackendServersInput, ...request.Option) (*ModifyServerGroupBackendServersOutput, error)
	ModifyServerGroupBackendServersRequest(*ModifyServerGroupBackendServersInput) (*request.Request, *ModifyServerGroupBackendServersOutput)

	RemoveAclEntriesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	RemoveAclEntriesCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	RemoveAclEntriesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	RemoveAclEntries(*RemoveAclEntriesInput) (*RemoveAclEntriesOutput, error)
	RemoveAclEntriesWithContext(volcengine.Context, *RemoveAclEntriesInput, ...request.Option) (*RemoveAclEntriesOutput, error)
	RemoveAclEntriesRequest(*RemoveAclEntriesInput) (*request.Request, *RemoveAclEntriesOutput)

	RemoveServerGroupBackendServersCommon(*map[string]interface{}) (*map[string]interface{}, error)
	RemoveServerGroupBackendServersCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	RemoveServerGroupBackendServersCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	RemoveServerGroupBackendServers(*RemoveServerGroupBackendServersInput) (*RemoveServerGroupBackendServersOutput, error)
	RemoveServerGroupBackendServersWithContext(volcengine.Context, *RemoveServerGroupBackendServersInput, ...request.Option) (*RemoveServerGroupBackendServersOutput, error)
	RemoveServerGroupBackendServersRequest(*RemoveServerGroupBackendServersInput) (*request.Request, *RemoveServerGroupBackendServersOutput)

	ReplaceCACertificateCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ReplaceCACertificateCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ReplaceCACertificateCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ReplaceCACertificate(*ReplaceCACertificateInput) (*ReplaceCACertificateOutput, error)
	ReplaceCACertificateWithContext(volcengine.Context, *ReplaceCACertificateInput, ...request.Option) (*ReplaceCACertificateOutput, error)
	ReplaceCACertificateRequest(*ReplaceCACertificateInput) (*request.Request, *ReplaceCACertificateOutput)

	ReplaceCertificateCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ReplaceCertificateCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ReplaceCertificateCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ReplaceCertificate(*ReplaceCertificateInput) (*ReplaceCertificateOutput, error)
	ReplaceCertificateWithContext(volcengine.Context, *ReplaceCertificateInput, ...request.Option) (*ReplaceCertificateOutput, error)
	ReplaceCertificateRequest(*ReplaceCertificateInput) (*request.Request, *ReplaceCertificateOutput)

	UploadCACertificateCommon(*map[string]interface{}) (*map[string]interface{}, error)
	UploadCACertificateCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	UploadCACertificateCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	UploadCACertificate(*UploadCACertificateInput) (*UploadCACertificateOutput, error)
	UploadCACertificateWithContext(volcengine.Context, *UploadCACertificateInput, ...request.Option) (*UploadCACertificateOutput, error)
	UploadCACertificateRequest(*UploadCACertificateInput) (*request.Request, *UploadCACertificateOutput)

	UploadCertificateCommon(*map[string]interface{}) (*map[string]interface{}, error)
	UploadCertificateCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	UploadCertificateCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	UploadCertificate(*UploadCertificateInput) (*UploadCertificateOutput, error)
	UploadCertificateWithContext(volcengine.Context, *UploadCertificateInput, ...request.Option) (*UploadCertificateOutput, error)
	UploadCertificateRequest(*UploadCertificateInput) (*request.Request, *UploadCertificateOutput)
}

var _ ALBAPI = (*ALB)(nil)

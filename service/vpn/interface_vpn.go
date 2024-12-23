// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package vpniface provides an interface to enable mocking the VPN service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package vpn

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
)

// VPNAPI provides an interface to enable mocking the
// vpn.VPN service client's API operation,
//
//    // volcengine sdk func uses an SDK service client to make a request to
//    // VPN.
//    func myFunc(svc VPNAPI) bool {
//        // Make svc.CreateCustomerGateway request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := vpn.New(sess)
//
//        myFunc(svc)
//    }
//
type VPNAPI interface {
	CreateCustomerGatewayCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CreateCustomerGatewayCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateCustomerGatewayCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateCustomerGateway(*CreateCustomerGatewayInput) (*CreateCustomerGatewayOutput, error)
	CreateCustomerGatewayWithContext(volcengine.Context, *CreateCustomerGatewayInput, ...request.Option) (*CreateCustomerGatewayOutput, error)
	CreateCustomerGatewayRequest(*CreateCustomerGatewayInput) (*request.Request, *CreateCustomerGatewayOutput)

	CreateSslVpnClientCertCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CreateSslVpnClientCertCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateSslVpnClientCertCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateSslVpnClientCert(*CreateSslVpnClientCertInput) (*CreateSslVpnClientCertOutput, error)
	CreateSslVpnClientCertWithContext(volcengine.Context, *CreateSslVpnClientCertInput, ...request.Option) (*CreateSslVpnClientCertOutput, error)
	CreateSslVpnClientCertRequest(*CreateSslVpnClientCertInput) (*request.Request, *CreateSslVpnClientCertOutput)

	CreateSslVpnServerCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CreateSslVpnServerCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateSslVpnServerCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateSslVpnServer(*CreateSslVpnServerInput) (*CreateSslVpnServerOutput, error)
	CreateSslVpnServerWithContext(volcengine.Context, *CreateSslVpnServerInput, ...request.Option) (*CreateSslVpnServerOutput, error)
	CreateSslVpnServerRequest(*CreateSslVpnServerInput) (*request.Request, *CreateSslVpnServerOutput)

	CreateVpnConnectionCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CreateVpnConnectionCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateVpnConnectionCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateVpnConnection(*CreateVpnConnectionInput) (*CreateVpnConnectionOutput, error)
	CreateVpnConnectionWithContext(volcengine.Context, *CreateVpnConnectionInput, ...request.Option) (*CreateVpnConnectionOutput, error)
	CreateVpnConnectionRequest(*CreateVpnConnectionInput) (*request.Request, *CreateVpnConnectionOutput)

	CreateVpnConnectionHealthCheckersCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CreateVpnConnectionHealthCheckersCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateVpnConnectionHealthCheckersCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateVpnConnectionHealthCheckers(*CreateVpnConnectionHealthCheckersInput) (*CreateVpnConnectionHealthCheckersOutput, error)
	CreateVpnConnectionHealthCheckersWithContext(volcengine.Context, *CreateVpnConnectionHealthCheckersInput, ...request.Option) (*CreateVpnConnectionHealthCheckersOutput, error)
	CreateVpnConnectionHealthCheckersRequest(*CreateVpnConnectionHealthCheckersInput) (*request.Request, *CreateVpnConnectionHealthCheckersOutput)

	CreateVpnGatewayCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CreateVpnGatewayCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateVpnGatewayCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateVpnGateway(*CreateVpnGatewayInput) (*CreateVpnGatewayOutput, error)
	CreateVpnGatewayWithContext(volcengine.Context, *CreateVpnGatewayInput, ...request.Option) (*CreateVpnGatewayOutput, error)
	CreateVpnGatewayRequest(*CreateVpnGatewayInput) (*request.Request, *CreateVpnGatewayOutput)

	CreateVpnGatewayRouteCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CreateVpnGatewayRouteCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateVpnGatewayRouteCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateVpnGatewayRoute(*CreateVpnGatewayRouteInput) (*CreateVpnGatewayRouteOutput, error)
	CreateVpnGatewayRouteWithContext(volcengine.Context, *CreateVpnGatewayRouteInput, ...request.Option) (*CreateVpnGatewayRouteOutput, error)
	CreateVpnGatewayRouteRequest(*CreateVpnGatewayRouteInput) (*request.Request, *CreateVpnGatewayRouteOutput)

	DeleteCustomerGatewayCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteCustomerGatewayCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteCustomerGatewayCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteCustomerGateway(*DeleteCustomerGatewayInput) (*DeleteCustomerGatewayOutput, error)
	DeleteCustomerGatewayWithContext(volcengine.Context, *DeleteCustomerGatewayInput, ...request.Option) (*DeleteCustomerGatewayOutput, error)
	DeleteCustomerGatewayRequest(*DeleteCustomerGatewayInput) (*request.Request, *DeleteCustomerGatewayOutput)

	DeleteSslVpnClientCertCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteSslVpnClientCertCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteSslVpnClientCertCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteSslVpnClientCert(*DeleteSslVpnClientCertInput) (*DeleteSslVpnClientCertOutput, error)
	DeleteSslVpnClientCertWithContext(volcengine.Context, *DeleteSslVpnClientCertInput, ...request.Option) (*DeleteSslVpnClientCertOutput, error)
	DeleteSslVpnClientCertRequest(*DeleteSslVpnClientCertInput) (*request.Request, *DeleteSslVpnClientCertOutput)

	DeleteSslVpnServerCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteSslVpnServerCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteSslVpnServerCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteSslVpnServer(*DeleteSslVpnServerInput) (*DeleteSslVpnServerOutput, error)
	DeleteSslVpnServerWithContext(volcengine.Context, *DeleteSslVpnServerInput, ...request.Option) (*DeleteSslVpnServerOutput, error)
	DeleteSslVpnServerRequest(*DeleteSslVpnServerInput) (*request.Request, *DeleteSslVpnServerOutput)

	DeleteVpnConnectionCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteVpnConnectionCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteVpnConnectionCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteVpnConnection(*DeleteVpnConnectionInput) (*DeleteVpnConnectionOutput, error)
	DeleteVpnConnectionWithContext(volcengine.Context, *DeleteVpnConnectionInput, ...request.Option) (*DeleteVpnConnectionOutput, error)
	DeleteVpnConnectionRequest(*DeleteVpnConnectionInput) (*request.Request, *DeleteVpnConnectionOutput)

	DeleteVpnConnectionHealthCheckerCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteVpnConnectionHealthCheckerCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteVpnConnectionHealthCheckerCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteVpnConnectionHealthChecker(*DeleteVpnConnectionHealthCheckerInput) (*DeleteVpnConnectionHealthCheckerOutput, error)
	DeleteVpnConnectionHealthCheckerWithContext(volcengine.Context, *DeleteVpnConnectionHealthCheckerInput, ...request.Option) (*DeleteVpnConnectionHealthCheckerOutput, error)
	DeleteVpnConnectionHealthCheckerRequest(*DeleteVpnConnectionHealthCheckerInput) (*request.Request, *DeleteVpnConnectionHealthCheckerOutput)

	DeleteVpnGatewayCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteVpnGatewayCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteVpnGatewayCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteVpnGateway(*DeleteVpnGatewayInput) (*DeleteVpnGatewayOutput, error)
	DeleteVpnGatewayWithContext(volcengine.Context, *DeleteVpnGatewayInput, ...request.Option) (*DeleteVpnGatewayOutput, error)
	DeleteVpnGatewayRequest(*DeleteVpnGatewayInput) (*request.Request, *DeleteVpnGatewayOutput)

	DeleteVpnGatewayRouteCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteVpnGatewayRouteCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteVpnGatewayRouteCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteVpnGatewayRoute(*DeleteVpnGatewayRouteInput) (*DeleteVpnGatewayRouteOutput, error)
	DeleteVpnGatewayRouteWithContext(volcengine.Context, *DeleteVpnGatewayRouteInput, ...request.Option) (*DeleteVpnGatewayRouteOutput, error)
	DeleteVpnGatewayRouteRequest(*DeleteVpnGatewayRouteInput) (*request.Request, *DeleteVpnGatewayRouteOutput)

	DescribeCustomerGatewayAttributesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeCustomerGatewayAttributesCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeCustomerGatewayAttributesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeCustomerGatewayAttributes(*DescribeCustomerGatewayAttributesInput) (*DescribeCustomerGatewayAttributesOutput, error)
	DescribeCustomerGatewayAttributesWithContext(volcengine.Context, *DescribeCustomerGatewayAttributesInput, ...request.Option) (*DescribeCustomerGatewayAttributesOutput, error)
	DescribeCustomerGatewayAttributesRequest(*DescribeCustomerGatewayAttributesInput) (*request.Request, *DescribeCustomerGatewayAttributesOutput)

	DescribeCustomerGatewaysCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeCustomerGatewaysCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeCustomerGatewaysCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeCustomerGateways(*DescribeCustomerGatewaysInput) (*DescribeCustomerGatewaysOutput, error)
	DescribeCustomerGatewaysWithContext(volcengine.Context, *DescribeCustomerGatewaysInput, ...request.Option) (*DescribeCustomerGatewaysOutput, error)
	DescribeCustomerGatewaysRequest(*DescribeCustomerGatewaysInput) (*request.Request, *DescribeCustomerGatewaysOutput)

	DescribeSslVpnClientCertAttributesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeSslVpnClientCertAttributesCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeSslVpnClientCertAttributesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeSslVpnClientCertAttributes(*DescribeSslVpnClientCertAttributesInput) (*DescribeSslVpnClientCertAttributesOutput, error)
	DescribeSslVpnClientCertAttributesWithContext(volcengine.Context, *DescribeSslVpnClientCertAttributesInput, ...request.Option) (*DescribeSslVpnClientCertAttributesOutput, error)
	DescribeSslVpnClientCertAttributesRequest(*DescribeSslVpnClientCertAttributesInput) (*request.Request, *DescribeSslVpnClientCertAttributesOutput)

	DescribeSslVpnClientCertsCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeSslVpnClientCertsCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeSslVpnClientCertsCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeSslVpnClientCerts(*DescribeSslVpnClientCertsInput) (*DescribeSslVpnClientCertsOutput, error)
	DescribeSslVpnClientCertsWithContext(volcengine.Context, *DescribeSslVpnClientCertsInput, ...request.Option) (*DescribeSslVpnClientCertsOutput, error)
	DescribeSslVpnClientCertsRequest(*DescribeSslVpnClientCertsInput) (*request.Request, *DescribeSslVpnClientCertsOutput)

	DescribeSslVpnServersCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeSslVpnServersCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeSslVpnServersCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeSslVpnServers(*DescribeSslVpnServersInput) (*DescribeSslVpnServersOutput, error)
	DescribeSslVpnServersWithContext(volcengine.Context, *DescribeSslVpnServersInput, ...request.Option) (*DescribeSslVpnServersOutput, error)
	DescribeSslVpnServersRequest(*DescribeSslVpnServersInput) (*request.Request, *DescribeSslVpnServersOutput)

	DescribeVpnConnectionAttributesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeVpnConnectionAttributesCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeVpnConnectionAttributesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeVpnConnectionAttributes(*DescribeVpnConnectionAttributesInput) (*DescribeVpnConnectionAttributesOutput, error)
	DescribeVpnConnectionAttributesWithContext(volcengine.Context, *DescribeVpnConnectionAttributesInput, ...request.Option) (*DescribeVpnConnectionAttributesOutput, error)
	DescribeVpnConnectionAttributesRequest(*DescribeVpnConnectionAttributesInput) (*request.Request, *DescribeVpnConnectionAttributesOutput)

	DescribeVpnConnectionsCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeVpnConnectionsCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeVpnConnectionsCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeVpnConnections(*DescribeVpnConnectionsInput) (*DescribeVpnConnectionsOutput, error)
	DescribeVpnConnectionsWithContext(volcengine.Context, *DescribeVpnConnectionsInput, ...request.Option) (*DescribeVpnConnectionsOutput, error)
	DescribeVpnConnectionsRequest(*DescribeVpnConnectionsInput) (*request.Request, *DescribeVpnConnectionsOutput)

	DescribeVpnGatewayAttributesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeVpnGatewayAttributesCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeVpnGatewayAttributesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeVpnGatewayAttributes(*DescribeVpnGatewayAttributesInput) (*DescribeVpnGatewayAttributesOutput, error)
	DescribeVpnGatewayAttributesWithContext(volcengine.Context, *DescribeVpnGatewayAttributesInput, ...request.Option) (*DescribeVpnGatewayAttributesOutput, error)
	DescribeVpnGatewayAttributesRequest(*DescribeVpnGatewayAttributesInput) (*request.Request, *DescribeVpnGatewayAttributesOutput)

	DescribeVpnGatewayRouteAttributesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeVpnGatewayRouteAttributesCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeVpnGatewayRouteAttributesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeVpnGatewayRouteAttributes(*DescribeVpnGatewayRouteAttributesInput) (*DescribeVpnGatewayRouteAttributesOutput, error)
	DescribeVpnGatewayRouteAttributesWithContext(volcengine.Context, *DescribeVpnGatewayRouteAttributesInput, ...request.Option) (*DescribeVpnGatewayRouteAttributesOutput, error)
	DescribeVpnGatewayRouteAttributesRequest(*DescribeVpnGatewayRouteAttributesInput) (*request.Request, *DescribeVpnGatewayRouteAttributesOutput)

	DescribeVpnGatewayRoutesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeVpnGatewayRoutesCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeVpnGatewayRoutesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeVpnGatewayRoutes(*DescribeVpnGatewayRoutesInput) (*DescribeVpnGatewayRoutesOutput, error)
	DescribeVpnGatewayRoutesWithContext(volcengine.Context, *DescribeVpnGatewayRoutesInput, ...request.Option) (*DescribeVpnGatewayRoutesOutput, error)
	DescribeVpnGatewayRoutesRequest(*DescribeVpnGatewayRoutesInput) (*request.Request, *DescribeVpnGatewayRoutesOutput)

	DescribeVpnGatewaysCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeVpnGatewaysCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeVpnGatewaysCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeVpnGateways(*DescribeVpnGatewaysInput) (*DescribeVpnGatewaysOutput, error)
	DescribeVpnGatewaysWithContext(volcengine.Context, *DescribeVpnGatewaysInput, ...request.Option) (*DescribeVpnGatewaysOutput, error)
	DescribeVpnGatewaysRequest(*DescribeVpnGatewaysInput) (*request.Request, *DescribeVpnGatewaysOutput)

	DescribeVpnGatewaysBillingCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeVpnGatewaysBillingCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeVpnGatewaysBillingCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeVpnGatewaysBilling(*DescribeVpnGatewaysBillingInput) (*DescribeVpnGatewaysBillingOutput, error)
	DescribeVpnGatewaysBillingWithContext(volcengine.Context, *DescribeVpnGatewaysBillingInput, ...request.Option) (*DescribeVpnGatewaysBillingOutput, error)
	DescribeVpnGatewaysBillingRequest(*DescribeVpnGatewaysBillingInput) (*request.Request, *DescribeVpnGatewaysBillingOutput)

	ModifyCustomerGatewayAttributesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ModifyCustomerGatewayAttributesCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifyCustomerGatewayAttributesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifyCustomerGatewayAttributes(*ModifyCustomerGatewayAttributesInput) (*ModifyCustomerGatewayAttributesOutput, error)
	ModifyCustomerGatewayAttributesWithContext(volcengine.Context, *ModifyCustomerGatewayAttributesInput, ...request.Option) (*ModifyCustomerGatewayAttributesOutput, error)
	ModifyCustomerGatewayAttributesRequest(*ModifyCustomerGatewayAttributesInput) (*request.Request, *ModifyCustomerGatewayAttributesOutput)

	ModifySslVpnClientCertCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ModifySslVpnClientCertCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifySslVpnClientCertCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifySslVpnClientCert(*ModifySslVpnClientCertInput) (*ModifySslVpnClientCertOutput, error)
	ModifySslVpnClientCertWithContext(volcengine.Context, *ModifySslVpnClientCertInput, ...request.Option) (*ModifySslVpnClientCertOutput, error)
	ModifySslVpnClientCertRequest(*ModifySslVpnClientCertInput) (*request.Request, *ModifySslVpnClientCertOutput)

	ModifySslVpnServerCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ModifySslVpnServerCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifySslVpnServerCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifySslVpnServer(*ModifySslVpnServerInput) (*ModifySslVpnServerOutput, error)
	ModifySslVpnServerWithContext(volcengine.Context, *ModifySslVpnServerInput, ...request.Option) (*ModifySslVpnServerOutput, error)
	ModifySslVpnServerRequest(*ModifySslVpnServerInput) (*request.Request, *ModifySslVpnServerOutput)

	ModifyVpnConnectionAttributesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ModifyVpnConnectionAttributesCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifyVpnConnectionAttributesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifyVpnConnectionAttributes(*ModifyVpnConnectionAttributesInput) (*ModifyVpnConnectionAttributesOutput, error)
	ModifyVpnConnectionAttributesWithContext(volcengine.Context, *ModifyVpnConnectionAttributesInput, ...request.Option) (*ModifyVpnConnectionAttributesOutput, error)
	ModifyVpnConnectionAttributesRequest(*ModifyVpnConnectionAttributesInput) (*request.Request, *ModifyVpnConnectionAttributesOutput)

	ModifyVpnConnectionHealthCheckerCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ModifyVpnConnectionHealthCheckerCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifyVpnConnectionHealthCheckerCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifyVpnConnectionHealthChecker(*ModifyVpnConnectionHealthCheckerInput) (*ModifyVpnConnectionHealthCheckerOutput, error)
	ModifyVpnConnectionHealthCheckerWithContext(volcengine.Context, *ModifyVpnConnectionHealthCheckerInput, ...request.Option) (*ModifyVpnConnectionHealthCheckerOutput, error)
	ModifyVpnConnectionHealthCheckerRequest(*ModifyVpnConnectionHealthCheckerInput) (*request.Request, *ModifyVpnConnectionHealthCheckerOutput)

	ModifyVpnGatewayAttributesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ModifyVpnGatewayAttributesCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifyVpnGatewayAttributesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifyVpnGatewayAttributes(*ModifyVpnGatewayAttributesInput) (*ModifyVpnGatewayAttributesOutput, error)
	ModifyVpnGatewayAttributesWithContext(volcengine.Context, *ModifyVpnGatewayAttributesInput, ...request.Option) (*ModifyVpnGatewayAttributesOutput, error)
	ModifyVpnGatewayAttributesRequest(*ModifyVpnGatewayAttributesInput) (*request.Request, *ModifyVpnGatewayAttributesOutput)

	RenewVpnGatewayCommon(*map[string]interface{}) (*map[string]interface{}, error)
	RenewVpnGatewayCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	RenewVpnGatewayCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	RenewVpnGateway(*RenewVpnGatewayInput) (*RenewVpnGatewayOutput, error)
	RenewVpnGatewayWithContext(volcengine.Context, *RenewVpnGatewayInput, ...request.Option) (*RenewVpnGatewayOutput, error)
	RenewVpnGatewayRequest(*RenewVpnGatewayInput) (*request.Request, *RenewVpnGatewayOutput)

	ResetVpnConnectionCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ResetVpnConnectionCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ResetVpnConnectionCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ResetVpnConnection(*ResetVpnConnectionInput) (*ResetVpnConnectionOutput, error)
	ResetVpnConnectionWithContext(volcengine.Context, *ResetVpnConnectionInput, ...request.Option) (*ResetVpnConnectionOutput, error)
	ResetVpnConnectionRequest(*ResetVpnConnectionInput) (*request.Request, *ResetVpnConnectionOutput)

	SetVpnGatewayRenewalCommon(*map[string]interface{}) (*map[string]interface{}, error)
	SetVpnGatewayRenewalCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	SetVpnGatewayRenewalCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	SetVpnGatewayRenewal(*SetVpnGatewayRenewalInput) (*SetVpnGatewayRenewalOutput, error)
	SetVpnGatewayRenewalWithContext(volcengine.Context, *SetVpnGatewayRenewalInput, ...request.Option) (*SetVpnGatewayRenewalOutput, error)
	SetVpnGatewayRenewalRequest(*SetVpnGatewayRenewalInput) (*request.Request, *SetVpnGatewayRenewalOutput)
}

var _ VPNAPI = (*VPN)(nil)

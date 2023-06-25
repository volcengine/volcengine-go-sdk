// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package directconnectiface provides an interface to enable mocking the DIRECTCONNECT service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package directconnect

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
)

// DIRECTCONNECTAPI provides an interface to enable mocking the
// directconnect.DIRECTCONNECT service client's API operation,
//
//	// volcengine sdk func uses an SDK service client to make a request to
//	// DIRECTCONNECT.
//	func myFunc(svc DIRECTCONNECTAPI) bool {
//	    // Make svc.CreateBgpPeer request
//	}
//
//	func main() {
//	    sess := session.New()
//	    svc := directconnect.New(sess)
//
//	    myFunc(svc)
//	}
type DIRECTCONNECTAPI interface {
	CreateBgpPeerCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CreateBgpPeerCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateBgpPeerCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateBgpPeer(*CreateBgpPeerInput) (*CreateBgpPeerOutput, error)
	CreateBgpPeerWithContext(volcengine.Context, *CreateBgpPeerInput, ...request.Option) (*CreateBgpPeerOutput, error)
	CreateBgpPeerRequest(*CreateBgpPeerInput) (*request.Request, *CreateBgpPeerOutput)

	CreateDirectConnectConnectionCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CreateDirectConnectConnectionCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateDirectConnectConnectionCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateDirectConnectConnection(*CreateDirectConnectConnectionInput) (*CreateDirectConnectConnectionOutput, error)
	CreateDirectConnectConnectionWithContext(volcengine.Context, *CreateDirectConnectConnectionInput, ...request.Option) (*CreateDirectConnectConnectionOutput, error)
	CreateDirectConnectConnectionRequest(*CreateDirectConnectConnectionInput) (*request.Request, *CreateDirectConnectConnectionOutput)

	CreateDirectConnectConnectionOrderCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CreateDirectConnectConnectionOrderCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateDirectConnectConnectionOrderCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateDirectConnectConnectionOrder(*CreateDirectConnectConnectionOrderInput) (*CreateDirectConnectConnectionOrderOutput, error)
	CreateDirectConnectConnectionOrderWithContext(volcengine.Context, *CreateDirectConnectConnectionOrderInput, ...request.Option) (*CreateDirectConnectConnectionOrderOutput, error)
	CreateDirectConnectConnectionOrderRequest(*CreateDirectConnectConnectionOrderInput) (*request.Request, *CreateDirectConnectConnectionOrderOutput)

	CreateDirectConnectGatewayCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CreateDirectConnectGatewayCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateDirectConnectGatewayCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateDirectConnectGateway(*CreateDirectConnectGatewayInput) (*CreateDirectConnectGatewayOutput, error)
	CreateDirectConnectGatewayWithContext(volcengine.Context, *CreateDirectConnectGatewayInput, ...request.Option) (*CreateDirectConnectGatewayOutput, error)
	CreateDirectConnectGatewayRequest(*CreateDirectConnectGatewayInput) (*request.Request, *CreateDirectConnectGatewayOutput)

	CreateDirectConnectGatewayRouteCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CreateDirectConnectGatewayRouteCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateDirectConnectGatewayRouteCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateDirectConnectGatewayRoute(*CreateDirectConnectGatewayRouteInput) (*CreateDirectConnectGatewayRouteOutput, error)
	CreateDirectConnectGatewayRouteWithContext(volcengine.Context, *CreateDirectConnectGatewayRouteInput, ...request.Option) (*CreateDirectConnectGatewayRouteOutput, error)
	CreateDirectConnectGatewayRouteRequest(*CreateDirectConnectGatewayRouteInput) (*request.Request, *CreateDirectConnectGatewayRouteOutput)

	CreateDirectConnectVirtualInterfaceCommon(*map[string]interface{}) (*map[string]interface{}, error)
	CreateDirectConnectVirtualInterfaceCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateDirectConnectVirtualInterfaceCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateDirectConnectVirtualInterface(*CreateDirectConnectVirtualInterfaceInput) (*CreateDirectConnectVirtualInterfaceOutput, error)
	CreateDirectConnectVirtualInterfaceWithContext(volcengine.Context, *CreateDirectConnectVirtualInterfaceInput, ...request.Option) (*CreateDirectConnectVirtualInterfaceOutput, error)
	CreateDirectConnectVirtualInterfaceRequest(*CreateDirectConnectVirtualInterfaceInput) (*request.Request, *CreateDirectConnectVirtualInterfaceOutput)

	DeleteBgpPeerCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteBgpPeerCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteBgpPeerCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteBgpPeer(*DeleteBgpPeerInput) (*DeleteBgpPeerOutput, error)
	DeleteBgpPeerWithContext(volcengine.Context, *DeleteBgpPeerInput, ...request.Option) (*DeleteBgpPeerOutput, error)
	DeleteBgpPeerRequest(*DeleteBgpPeerInput) (*request.Request, *DeleteBgpPeerOutput)

	DeleteDirectConnectConnectionCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteDirectConnectConnectionCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteDirectConnectConnectionCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteDirectConnectConnection(*DeleteDirectConnectConnectionInput) (*DeleteDirectConnectConnectionOutput, error)
	DeleteDirectConnectConnectionWithContext(volcengine.Context, *DeleteDirectConnectConnectionInput, ...request.Option) (*DeleteDirectConnectConnectionOutput, error)
	DeleteDirectConnectConnectionRequest(*DeleteDirectConnectConnectionInput) (*request.Request, *DeleteDirectConnectConnectionOutput)

	DeleteDirectConnectGatewayCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteDirectConnectGatewayCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteDirectConnectGatewayCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteDirectConnectGateway(*DeleteDirectConnectGatewayInput) (*DeleteDirectConnectGatewayOutput, error)
	DeleteDirectConnectGatewayWithContext(volcengine.Context, *DeleteDirectConnectGatewayInput, ...request.Option) (*DeleteDirectConnectGatewayOutput, error)
	DeleteDirectConnectGatewayRequest(*DeleteDirectConnectGatewayInput) (*request.Request, *DeleteDirectConnectGatewayOutput)

	DeleteDirectConnectGatewayRouteCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteDirectConnectGatewayRouteCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteDirectConnectGatewayRouteCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteDirectConnectGatewayRoute(*DeleteDirectConnectGatewayRouteInput) (*DeleteDirectConnectGatewayRouteOutput, error)
	DeleteDirectConnectGatewayRouteWithContext(volcengine.Context, *DeleteDirectConnectGatewayRouteInput, ...request.Option) (*DeleteDirectConnectGatewayRouteOutput, error)
	DeleteDirectConnectGatewayRouteRequest(*DeleteDirectConnectGatewayRouteInput) (*request.Request, *DeleteDirectConnectGatewayRouteOutput)

	DeleteDirectConnectVirtualInterfaceCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteDirectConnectVirtualInterfaceCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteDirectConnectVirtualInterfaceCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteDirectConnectVirtualInterface(*DeleteDirectConnectVirtualInterfaceInput) (*DeleteDirectConnectVirtualInterfaceOutput, error)
	DeleteDirectConnectVirtualInterfaceWithContext(volcengine.Context, *DeleteDirectConnectVirtualInterfaceInput, ...request.Option) (*DeleteDirectConnectVirtualInterfaceOutput, error)
	DeleteDirectConnectVirtualInterfaceRequest(*DeleteDirectConnectVirtualInterfaceInput) (*request.Request, *DeleteDirectConnectVirtualInterfaceOutput)

	DescribeBgpPeerAttributesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeBgpPeerAttributesCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeBgpPeerAttributesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeBgpPeerAttributes(*DescribeBgpPeerAttributesInput) (*DescribeBgpPeerAttributesOutput, error)
	DescribeBgpPeerAttributesWithContext(volcengine.Context, *DescribeBgpPeerAttributesInput, ...request.Option) (*DescribeBgpPeerAttributesOutput, error)
	DescribeBgpPeerAttributesRequest(*DescribeBgpPeerAttributesInput) (*request.Request, *DescribeBgpPeerAttributesOutput)

	DescribeBgpPeersCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeBgpPeersCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeBgpPeersCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeBgpPeers(*DescribeBgpPeersInput) (*DescribeBgpPeersOutput, error)
	DescribeBgpPeersWithContext(volcengine.Context, *DescribeBgpPeersInput, ...request.Option) (*DescribeBgpPeersOutput, error)
	DescribeBgpPeersRequest(*DescribeBgpPeersInput) (*request.Request, *DescribeBgpPeersOutput)

	DescribeDirectConnectAccessPointsCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeDirectConnectAccessPointsCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeDirectConnectAccessPointsCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeDirectConnectAccessPoints(*DescribeDirectConnectAccessPointsInput) (*DescribeDirectConnectAccessPointsOutput, error)
	DescribeDirectConnectAccessPointsWithContext(volcengine.Context, *DescribeDirectConnectAccessPointsInput, ...request.Option) (*DescribeDirectConnectAccessPointsOutput, error)
	DescribeDirectConnectAccessPointsRequest(*DescribeDirectConnectAccessPointsInput) (*request.Request, *DescribeDirectConnectAccessPointsOutput)

	DescribeDirectConnectConnectionAttributesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeDirectConnectConnectionAttributesCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeDirectConnectConnectionAttributesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeDirectConnectConnectionAttributes(*DescribeDirectConnectConnectionAttributesInput) (*DescribeDirectConnectConnectionAttributesOutput, error)
	DescribeDirectConnectConnectionAttributesWithContext(volcengine.Context, *DescribeDirectConnectConnectionAttributesInput, ...request.Option) (*DescribeDirectConnectConnectionAttributesOutput, error)
	DescribeDirectConnectConnectionAttributesRequest(*DescribeDirectConnectConnectionAttributesInput) (*request.Request, *DescribeDirectConnectConnectionAttributesOutput)

	DescribeDirectConnectConnectionsCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeDirectConnectConnectionsCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeDirectConnectConnectionsCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeDirectConnectConnections(*DescribeDirectConnectConnectionsInput) (*DescribeDirectConnectConnectionsOutput, error)
	DescribeDirectConnectConnectionsWithContext(volcengine.Context, *DescribeDirectConnectConnectionsInput, ...request.Option) (*DescribeDirectConnectConnectionsOutput, error)
	DescribeDirectConnectConnectionsRequest(*DescribeDirectConnectConnectionsInput) (*request.Request, *DescribeDirectConnectConnectionsOutput)

	DescribeDirectConnectGatewayAttributesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeDirectConnectGatewayAttributesCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeDirectConnectGatewayAttributesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeDirectConnectGatewayAttributes(*DescribeDirectConnectGatewayAttributesInput) (*DescribeDirectConnectGatewayAttributesOutput, error)
	DescribeDirectConnectGatewayAttributesWithContext(volcengine.Context, *DescribeDirectConnectGatewayAttributesInput, ...request.Option) (*DescribeDirectConnectGatewayAttributesOutput, error)
	DescribeDirectConnectGatewayAttributesRequest(*DescribeDirectConnectGatewayAttributesInput) (*request.Request, *DescribeDirectConnectGatewayAttributesOutput)

	DescribeDirectConnectGatewayRouteAttributesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeDirectConnectGatewayRouteAttributesCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeDirectConnectGatewayRouteAttributesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeDirectConnectGatewayRouteAttributes(*DescribeDirectConnectGatewayRouteAttributesInput) (*DescribeDirectConnectGatewayRouteAttributesOutput, error)
	DescribeDirectConnectGatewayRouteAttributesWithContext(volcengine.Context, *DescribeDirectConnectGatewayRouteAttributesInput, ...request.Option) (*DescribeDirectConnectGatewayRouteAttributesOutput, error)
	DescribeDirectConnectGatewayRouteAttributesRequest(*DescribeDirectConnectGatewayRouteAttributesInput) (*request.Request, *DescribeDirectConnectGatewayRouteAttributesOutput)

	DescribeDirectConnectGatewayRoutesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeDirectConnectGatewayRoutesCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeDirectConnectGatewayRoutesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeDirectConnectGatewayRoutes(*DescribeDirectConnectGatewayRoutesInput) (*DescribeDirectConnectGatewayRoutesOutput, error)
	DescribeDirectConnectGatewayRoutesWithContext(volcengine.Context, *DescribeDirectConnectGatewayRoutesInput, ...request.Option) (*DescribeDirectConnectGatewayRoutesOutput, error)
	DescribeDirectConnectGatewayRoutesRequest(*DescribeDirectConnectGatewayRoutesInput) (*request.Request, *DescribeDirectConnectGatewayRoutesOutput)

	DescribeDirectConnectGatewaysCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeDirectConnectGatewaysCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeDirectConnectGatewaysCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeDirectConnectGateways(*DescribeDirectConnectGatewaysInput) (*DescribeDirectConnectGatewaysOutput, error)
	DescribeDirectConnectGatewaysWithContext(volcengine.Context, *DescribeDirectConnectGatewaysInput, ...request.Option) (*DescribeDirectConnectGatewaysOutput, error)
	DescribeDirectConnectGatewaysRequest(*DescribeDirectConnectGatewaysInput) (*request.Request, *DescribeDirectConnectGatewaysOutput)

	DescribeDirectConnectVirtualInterfaceAttributesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeDirectConnectVirtualInterfaceAttributesCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeDirectConnectVirtualInterfaceAttributesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeDirectConnectVirtualInterfaceAttributes(*DescribeDirectConnectVirtualInterfaceAttributesInput) (*DescribeDirectConnectVirtualInterfaceAttributesOutput, error)
	DescribeDirectConnectVirtualInterfaceAttributesWithContext(volcengine.Context, *DescribeDirectConnectVirtualInterfaceAttributesInput, ...request.Option) (*DescribeDirectConnectVirtualInterfaceAttributesOutput, error)
	DescribeDirectConnectVirtualInterfaceAttributesRequest(*DescribeDirectConnectVirtualInterfaceAttributesInput) (*request.Request, *DescribeDirectConnectVirtualInterfaceAttributesOutput)

	DescribeDirectConnectVirtualInterfacesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeDirectConnectVirtualInterfacesCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeDirectConnectVirtualInterfacesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeDirectConnectVirtualInterfaces(*DescribeDirectConnectVirtualInterfacesInput) (*DescribeDirectConnectVirtualInterfacesOutput, error)
	DescribeDirectConnectVirtualInterfacesWithContext(volcengine.Context, *DescribeDirectConnectVirtualInterfacesInput, ...request.Option) (*DescribeDirectConnectVirtualInterfacesOutput, error)
	DescribeDirectConnectVirtualInterfacesRequest(*DescribeDirectConnectVirtualInterfacesInput) (*request.Request, *DescribeDirectConnectVirtualInterfacesOutput)

	ModifyBgpPeerAttributesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ModifyBgpPeerAttributesCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifyBgpPeerAttributesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifyBgpPeerAttributes(*ModifyBgpPeerAttributesInput) (*ModifyBgpPeerAttributesOutput, error)
	ModifyBgpPeerAttributesWithContext(volcengine.Context, *ModifyBgpPeerAttributesInput, ...request.Option) (*ModifyBgpPeerAttributesOutput, error)
	ModifyBgpPeerAttributesRequest(*ModifyBgpPeerAttributesInput) (*request.Request, *ModifyBgpPeerAttributesOutput)

	ModifyDirectConnectConnectionAttributesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ModifyDirectConnectConnectionAttributesCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifyDirectConnectConnectionAttributesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifyDirectConnectConnectionAttributes(*ModifyDirectConnectConnectionAttributesInput) (*ModifyDirectConnectConnectionAttributesOutput, error)
	ModifyDirectConnectConnectionAttributesWithContext(volcengine.Context, *ModifyDirectConnectConnectionAttributesInput, ...request.Option) (*ModifyDirectConnectConnectionAttributesOutput, error)
	ModifyDirectConnectConnectionAttributesRequest(*ModifyDirectConnectConnectionAttributesInput) (*request.Request, *ModifyDirectConnectConnectionAttributesOutput)

	ModifyDirectConnectGatewayAttributesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ModifyDirectConnectGatewayAttributesCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifyDirectConnectGatewayAttributesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifyDirectConnectGatewayAttributes(*ModifyDirectConnectGatewayAttributesInput) (*ModifyDirectConnectGatewayAttributesOutput, error)
	ModifyDirectConnectGatewayAttributesWithContext(volcengine.Context, *ModifyDirectConnectGatewayAttributesInput, ...request.Option) (*ModifyDirectConnectGatewayAttributesOutput, error)
	ModifyDirectConnectGatewayAttributesRequest(*ModifyDirectConnectGatewayAttributesInput) (*request.Request, *ModifyDirectConnectGatewayAttributesOutput)

	ModifyDirectConnectVirtualInterfaceAttributesCommon(*map[string]interface{}) (*map[string]interface{}, error)
	ModifyDirectConnectVirtualInterfaceAttributesCommonWithContext(volcengine.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifyDirectConnectVirtualInterfaceAttributesCommonRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifyDirectConnectVirtualInterfaceAttributes(*ModifyDirectConnectVirtualInterfaceAttributesInput) (*ModifyDirectConnectVirtualInterfaceAttributesOutput, error)
	ModifyDirectConnectVirtualInterfaceAttributesWithContext(volcengine.Context, *ModifyDirectConnectVirtualInterfaceAttributesInput, ...request.Option) (*ModifyDirectConnectVirtualInterfaceAttributesOutput, error)
	ModifyDirectConnectVirtualInterfaceAttributesRequest(*ModifyDirectConnectVirtualInterfaceAttributesInput) (*request.Request, *ModifyDirectConnectVirtualInterfaceAttributesOutput)
}

var _ DIRECTCONNECTAPI = (*DIRECTCONNECT)(nil)

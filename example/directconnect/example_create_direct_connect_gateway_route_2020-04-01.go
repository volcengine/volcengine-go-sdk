// Example Code generated by Beijing Volcanoengine Technology.
package directconnectemample

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/service/directconnect"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)

func CreateDirectConnectGatewayRoute() {
	ak, sk, region := "Your AK", "Your SK", "Region"
	config := volcengine.NewConfig().
		WithRegion(region).
		WithCredentials(credentials.NewStaticCredentials(ak, sk, ""))
	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
	svc := directconnect.New(sess)
	createDirectConnectGatewayRouteInput := &directconnect.CreateDirectConnectGatewayRouteInput{
		DestinationCidrBlock:   volcengine.String("172.XX.XX.0/24"),
		DirectConnectGatewayId: volcengine.String("dcg-2fe3zsmkshs59g****"),
		NextHopId:              volcengine.String("dcv-7qthudw0ll6jmc****"),
	}

	resp, err := svc.CreateDirectConnectGatewayRoute(createDirectConnectGatewayRouteInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}

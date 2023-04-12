package vpnexample

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/service/vpn"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)

func DescribeVpnGatewayRoutes() {
	ak, sk, region := "Your AK", "Your SK", "Region"
	config := volcengine.NewConfig().
		WithRegion(region).
		WithCredentials(credentials.NewStaticCredentials(ak, sk, ""))
	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
	svc := vpn.New(sess)
	describeVpnGatewayRoutesInput := &vpn.DescribeVpnGatewayRoutesInput{
		VpnGatewayRouteIds: volcengine.StringSlice([]string{"vgr-3tex2c6c0v844c****"}),
	}

	resp, err := svc.DescribeVpnGatewayRoutes(describeVpnGatewayRoutesInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}

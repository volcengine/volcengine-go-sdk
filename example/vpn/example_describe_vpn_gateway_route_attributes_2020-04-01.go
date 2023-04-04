package vpn_example

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/service/vpn"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
    "github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)
func DescribeVpnGatewayRouteAttributes() {
	ak, sk, region := "Your AK", "Your SK", "Region"
	config := volcengine.NewConfig().
		WithRegion(region).
		WithCredentials(credentials.NewStaticCredentials(ak, sk, "")).
		WithDisableSSL(true)
	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
	svc := vpn.New(sess)
	describeVpnGatewayRouteAttributesInput := &vpn.DescribeVpnGatewayRouteAttributesInput{
		VpnGatewayRouteId: volcengine.String("vgr-3tex2c6c0v844c****"),
	}

	resp, err := svc.DescribeVpnGatewayRouteAttributes(describeVpnGatewayRouteAttributesInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}

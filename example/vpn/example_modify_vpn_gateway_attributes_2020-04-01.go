package vpn_example

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/service/vpn"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)

func ModifyVpnGatewayAttributes() {
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
	modifyVpnGatewayAttributesInput := &vpn.ModifyVpnGatewayAttributesInput{
		Description: volcengine.String("test"),
		VpnGatewayId: volcengine.String("vgw-12bfa2du7fojk17q7y1rk****"),
		VpnGatewayName: volcengine.String("test"),
	}

	resp, err := svc.ModifyVpnGatewayAttributes(modifyVpnGatewayAttributesInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}

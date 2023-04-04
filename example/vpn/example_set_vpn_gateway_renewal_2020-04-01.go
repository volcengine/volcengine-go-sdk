package vpn_example

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/service/vpn"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)

func SetVpnGatewayRenewal() {
	ak, sk, region := "Your AK", "Your SK", "Region"
	config := volcengine.NewConfig().
		WithRegion(region).
		WithCredentials(credentials.NewStaticCredentials(ak, sk, ""))
	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
	svc := vpn.New(sess)
	setVpnGatewayRenewalInput := &vpn.SetVpnGatewayRenewalInput{
		RenewType: volcengine.Int64(3),
		VpnGatewayId: volcengine.String("vgw-2fe7zjsz13ksg5oxruwed****"),
	}

	resp, err := svc.SetVpnGatewayRenewal(setVpnGatewayRenewalInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}

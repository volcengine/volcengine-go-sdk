package main

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/service/vpn"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
    "github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)
func DescribeVpnGateways() {
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
	describeVpnGatewaysInput := &vpn.DescribeVpnGatewaysInput{
		VpnGatewayIds: volcengine.StringSlice([]string{"vgw-12bfa2du7fojk17q7y1rk****"}),
	}

	resp, err := svc.DescribeVpnGateways(describeVpnGatewaysInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}

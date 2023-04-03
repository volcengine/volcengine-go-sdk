package main

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/service/vpn"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
    "github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)
func CreateVpnConnection() {
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
	createVpnConnectionInput := &vpn.CreateVpnConnectionInput{
		CustomerGatewayId: volcengine.String("cgw-274mm8eodvu9s7fap8skw****"),
		Description: volcengine.String("test"),
		VpnConnectionName: volcengine.String("test"),
		VpnGatewayId: volcengine.String("vgw-2752abxsju1vk7fap8sgk****"),
	}

	resp, err := svc.CreateVpnConnection(createVpnConnectionInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}

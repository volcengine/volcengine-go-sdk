package main

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/service/vpn"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
    "github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)
func DeleteCustomerGateway() {
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
	deleteCustomerGatewayInput := &vpn.DeleteCustomerGatewayInput{
		CustomerGatewayId: volcengine.String("cgw-2d670j2o9lc0058ozfddg****"),
	}

	resp, err := svc.DeleteCustomerGateway(deleteCustomerGatewayInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}

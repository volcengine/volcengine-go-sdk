package vpnexample

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/service/vpn"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)

func DescribeCustomerGateways() {
	ak, sk, region := "Your AK", "Your SK", "Region"
	config := volcengine.NewConfig().
		WithRegion(region).
		WithCredentials(credentials.NewStaticCredentials(ak, sk, ""))
	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
	svc := vpn.New(sess)
	describeCustomerGatewaysInput := &vpn.DescribeCustomerGatewaysInput{
		CustomerGatewayIds: volcengine.StringSlice([]string{"cgw-7qthudw0ll6jmc****"}),
		PageNumber:         volcengine.Int64(1),
		PageSize:           volcengine.Int64(20),
	}

	resp, err := svc.DescribeCustomerGateways(describeCustomerGatewaysInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}

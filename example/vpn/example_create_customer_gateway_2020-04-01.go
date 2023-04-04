package vpn_example

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/service/vpn"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)

func CreateCustomerGateway() {
	ak, sk, region := "Your AK", "Your SK", "Region"
	config := volcengine.NewConfig().
		WithRegion(region).
		WithCredentials(credentials.NewStaticCredentials(ak, sk, ""))
	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
	svc := vpn.New(sess)
	createCustomerGatewayInput := &vpn.CreateCustomerGatewayInput{
		CustomerGatewayName: volcengine.String("test"),
		Description:         volcengine.String("test"),
		IpAddress:           volcengine.String("1.1.1.1"),
	}

	resp, err := svc.CreateCustomerGateway(createCustomerGatewayInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}

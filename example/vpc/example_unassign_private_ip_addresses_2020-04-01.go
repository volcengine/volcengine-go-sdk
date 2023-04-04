package vpc_example

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/service/vpc"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)

func UnassignPrivateIpAddresses() {
	ak, sk, region := "Your AK", "Your SK", "Region"
	config := volcengine.NewConfig().
		WithRegion(region).
		WithCredentials(credentials.NewStaticCredentials(ak, sk, ""))
	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
	svc := vpc.New(sess)
	unassignPrivateIpAddressesInput := &vpc.UnassignPrivateIpAddressesInput{
		NetworkInterfaceId: volcengine.String("eni-bp67acfmxazb4ph****"),
		PrivateIpAddress: volcengine.StringSlice([]string{"192.168.XX.10"}),
	}

	resp, err := svc.UnassignPrivateIpAddresses(unassignPrivateIpAddressesInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}

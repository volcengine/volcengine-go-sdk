package vpc_example

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/service/vpc"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
    "github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)
func AssignPrivateIpAddresses() {
	ak, sk, region := "Your AK", "Your SK", "Region"
	config := volcengine.NewConfig().
		WithRegion(region).
		WithCredentials(credentials.NewStaticCredentials(ak, sk, "")).
		WithDisableSSL(true)
	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
	svc := vpc.New(sess)
	assignPrivateIpAddressesInput := &vpc.AssignPrivateIpAddressesInput{
		NetworkInterfaceId: volcengine.String("eni-bp67acfmxazb4p****"),
		PrivateIpAddress: volcengine.StringSlice([]string{"192.168.XX.10","192.168.XX.12"}),
	}

	resp, err := svc.AssignPrivateIpAddresses(assignPrivateIpAddressesInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}

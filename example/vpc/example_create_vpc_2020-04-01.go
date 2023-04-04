package vpc_example

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/service/vpc"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
    "github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)
func CreateVpc() {
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
	createVpcInput := &vpc.CreateVpcInput{
		CidrBlock: volcengine.String("172.XX.XX.0/12"),
		Description: volcengine.String("test"),
		DnsServers: volcengine.StringSlice([]string{"1.XX.XX.1"}),
		VpcName: volcengine.String("test"),
	}

	resp, err := svc.CreateVpc(createVpcInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}

package vpc_example

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/service/vpc"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)
func CreateNetworkInterface() {
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
	createNetworkInterfaceInput := &vpc.CreateNetworkInterfaceInput{
		NetworkInterfaceName: volcengine.String("test"),
		PrimaryIpAddress: volcengine.String("192.XX.XX.10"),
		SecurityGroupIds: volcengine.StringSlice([]string{"sg-123edfgt8hhvj****"}),
		SubnetId: volcengine.String("subnet-xxxxxx"),
	}

	resp, err := svc.CreateNetworkInterface(createNetworkInterfaceInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}

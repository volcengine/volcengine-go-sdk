package vpc_example

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/service/vpc"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)
func DetachNetworkInterface() {
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
	detachNetworkInterfaceInput := &vpc.DetachNetworkInterfaceInput{
		InstanceId: volcengine.String("i-wijfn35c****"),
		NetworkInterfaceId: volcengine.String("eni-bp1fgn8xyz9****"),
	}

	resp, err := svc.DetachNetworkInterface(detachNetworkInterfaceInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}

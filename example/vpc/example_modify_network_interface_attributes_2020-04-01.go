package vpc_example

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/service/vpc"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)

func ModifyNetworkInterfaceAttributes() {
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
	modifyNetworkInterfaceAttributesInput := &vpc.ModifyNetworkInterfaceAttributesInput{
		NetworkInterfaceName: volcengine.String("test"),
	}

	resp, err := svc.ModifyNetworkInterfaceAttributes(modifyNetworkInterfaceAttributesInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}

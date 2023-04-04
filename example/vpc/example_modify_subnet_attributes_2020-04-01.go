package vpc_example

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/service/vpc"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)

func ModifySubnetAttributes() {
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
	modifySubnetAttributesInput := &vpc.ModifySubnetAttributesInput{
		SubnetId: volcengine.String("subnet-bp15qtc7ywecf****"),
	}

	resp, err := svc.ModifySubnetAttributes(modifySubnetAttributesInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}

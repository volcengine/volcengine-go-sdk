package vpc_example

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/service/vpc"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)

func ModifySecurityGroupAttributes() {
	ak, sk, region := "Your AK", "Your SK", "Region"
	config := volcengine.NewConfig().
		WithRegion(region).
		WithCredentials(credentials.NewStaticCredentials(ak, sk, ""))
	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
	svc := vpc.New(sess)
	modifySecurityGroupAttributesInput := &vpc.ModifySecurityGroupAttributesInput{
		SecurityGroupId: volcengine.String("sg-bp67acfmxazb4p****"),
		SecurityGroupName: volcengine.String("test"),
	}

	resp, err := svc.ModifySecurityGroupAttributes(modifySecurityGroupAttributesInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}

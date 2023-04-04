package vpc_example

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/service/vpc"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
    "github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)
func ModifyNetworkAclAttributes() {
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
	modifyNetworkAclAttributesInput := &vpc.ModifyNetworkAclAttributesInput{
		Description: volcengine.String("testDescription"),
		NetworkAclId: volcengine.String("acl-bp1fg655nh68xyz9****"),
		NetworkAclName: volcengine.String("test-acl"),
	}

	resp, err := svc.ModifyNetworkAclAttributes(modifyNetworkAclAttributesInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}

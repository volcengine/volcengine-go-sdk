package vpc_example

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/service/vpc"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
    "github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)
func CreateNetworkAcl() {
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
	createNetworkAclInput := &vpc.CreateNetworkAclInput{
		Description: volcengine.String("testDescription"),
		NetworkAclName: volcengine.String("test-acl"),
		VpcId: volcengine.String("vpc-bp1opxu1zkhn00gzv****"),
	}

	resp, err := svc.CreateNetworkAcl(createNetworkAclInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}

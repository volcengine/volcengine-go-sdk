package main

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/service/vpc"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
    "github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)
func DescribeNetworkAclAttributes() {
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
	describeNetworkAclAttributesInput := &vpc.DescribeNetworkAclAttributesInput{
		NetworkAclId: volcengine.String("nacl-bp1fg655nh68xyz9****"),
	}

	resp, err := svc.DescribeNetworkAclAttributes(describeNetworkAclAttributesInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
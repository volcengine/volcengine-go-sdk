package main

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/service/vpc"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
    "github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)
func DescribeSubnetAttributes() {
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
	describeSubnetAttributesInput := &vpc.DescribeSubnetAttributesInput{
		SubnetId: volcengine.String("subnet-23dscddcffvf3****"),
	}

	resp, err := svc.DescribeSubnetAttributes(describeSubnetAttributesInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
package main

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/service/ecs"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
    "github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)
func DescribeInstanceTypes() {
	ak, sk, region := "Your AK", "Your SK", "Region"
	config := volcengine.NewConfig().
		WithRegion(region).
		WithCredentials(credentials.NewStaticCredentials(ak, sk, "")).
		WithDisableSSL(true)
	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
	svc := ecs.New(sess)
	describeInstanceTypesInput := &ecs.DescribeInstanceTypesInput{
		InstanceTypeIds: volcengine.StringSlice([]string{"ecs.g2i.large"}),
	}

	resp, err := svc.DescribeInstanceTypes(describeInstanceTypesInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
package ecs_example

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/service/ecs"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)

func DescribeInstanceTypeFamilies() {
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
	describeInstanceTypeFamiliesInput := &ecs.DescribeInstanceTypeFamiliesInput{
		ZoneId: volcengine.String("cn-beijing-a"),
	}

	resp, err := svc.DescribeInstanceTypeFamilies(describeInstanceTypeFamiliesInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}

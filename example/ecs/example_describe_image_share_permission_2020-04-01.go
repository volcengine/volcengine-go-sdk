package main

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/service/ecs"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
    "github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)
func DescribeImageSharePermission() {
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
	describeImageSharePermissionInput := &ecs.DescribeImageSharePermissionInput{
		ImageId: volcengine.String("image-ebgy1og99tfct0gw****"),
	}

	resp, err := svc.DescribeImageSharePermission(describeImageSharePermissionInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
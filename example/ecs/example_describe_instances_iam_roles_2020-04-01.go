package main

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/service/ecs"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
    "github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)
func DescribeInstancesIamRoles() {
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
	describeInstancesIamRolesInput := &ecs.DescribeInstancesIamRolesInput{
		InstanceIds: volcengine.StringSlice([]string{"i-cm9tn4z1eohuu4d3****"}),
	}

	resp, err := svc.DescribeInstancesIamRoles(describeInstancesIamRolesInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}

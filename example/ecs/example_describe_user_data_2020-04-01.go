package ecsexample

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/service/ecs"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)

func DescribeUserData() {
	ak, sk, region := "Your AK", "Your SK", "Region"
	config := volcengine.NewConfig().
		WithRegion(region).
		WithCredentials(credentials.NewStaticCredentials(ak, sk, ""))
	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
	svc := ecs.New(sess)
	describeUserDataInput := &ecs.DescribeUserDataInput{
		InstanceId: volcengine.String("i-l8j0j2jynelea7nu****"),
	}

	resp, err := svc.DescribeUserData(describeUserDataInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}

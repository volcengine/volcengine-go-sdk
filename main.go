package main

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/service/autoscaling"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)

func main() {
	ak, sk, region := "Your AK", "Your SK", "cn-beijing"
	config := volcengine.NewConfig().
		WithRegion(region).
		WithCredentials(credentials.NewStaticCredentials(ak, sk, "")).WithEndpointConfigState(true)
	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
	svc := autoscaling.New(sess)
	attachInstancesInput := &autoscaling.AttachInstancesInput{
		InstanceIds:    volcengine.StringSlice([]string{"i-ybmike5l70l8j1ha****"}),
		ScalingGroupId: volcengine.String("scg-ybmssdnnhn5pkgyd****"),
	}

	resp, err := svc.AttachInstances(attachInstancesInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}

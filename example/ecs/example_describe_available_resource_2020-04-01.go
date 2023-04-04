package ecs_example

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/service/ecs"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
    "github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)
func DescribeAvailableResource() {
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
	describeAvailableResourceInput := &ecs.DescribeAvailableResourceInput{
		DestinationResource: volcengine.String("InstanceType"),
		InstanceTypeId: volcengine.String("ecs.g1.large"),
		ZoneId: volcengine.String("cn-*****"),
	}

	resp, err := svc.DescribeAvailableResource(describeAvailableResourceInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}

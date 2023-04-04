package ecs_example

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/service/ecs"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
    "github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)
func DescribeSystemEvents() {
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
	describeSystemEventsInput := &ecs.DescribeSystemEventsInput{
		EventIds: volcengine.StringSlice([]string{"e-3ti9101aju3vj0******"}),
	}

	resp, err := svc.DescribeSystemEvents(describeSystemEventsInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}

package ecs_example

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/service/ecs"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)

func RenewInstance() {
	ak, sk, region := "Your AK", "Your SK", "Region"
	config := volcengine.NewConfig().
		WithRegion(region).
		WithCredentials(credentials.NewStaticCredentials(ak, sk, ""))
	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
	svc := ecs.New(sess)
	renewInstanceInput := &ecs.RenewInstanceInput{
		InstanceId: volcengine.String("i-3thhlu8byl4bwbha****"),
		Period:     volcengine.Int32(1),
		PeriodUnit: volcengine.String("Month"),
	}

	resp, err := svc.RenewInstance(renewInstanceInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}

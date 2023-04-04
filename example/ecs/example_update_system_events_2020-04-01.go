package ecs_example

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/service/ecs"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
    "github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)
func UpdateSystemEvents() {
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
	updateSystemEventsInput := &ecs.UpdateSystemEventsInput{
		EventIds: volcengine.StringSlice([]string{"t-3ti9101aju3vj0******"}),
		Status: volcengine.String("Executing"),
	}

	resp, err := svc.UpdateSystemEvents(updateSystemEventsInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}

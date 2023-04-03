package main

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/service/ecs"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
    "github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)
func GetConsoleOutput() {
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
	getConsoleOutputInput := &ecs.GetConsoleOutputInput{
		InstanceId: volcengine.String("i-3tiefmkskq3vj0******"),
	}

	resp, err := svc.GetConsoleOutput(getConsoleOutputInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
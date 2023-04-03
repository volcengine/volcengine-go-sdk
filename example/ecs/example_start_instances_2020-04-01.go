package main

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/service/ecs"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
    "github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)
func StartInstances() {
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
	startInstancesInput := &ecs.StartInstancesInput{
		InstanceIds: volcengine.StringSlice([]string{"i-ybo349sxoncm9t****","i-ybo349sxolcm9t****"}),
	}

	resp, err := svc.StartInstances(startInstancesInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}

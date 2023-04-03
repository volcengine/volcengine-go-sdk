package main

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/service/ecs"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
    "github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)
func ModifyInstanceSpec() {
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
	modifyInstanceSpecInput := &ecs.ModifyInstanceSpecInput{
		InstanceId: volcengine.String("i-3thhlu8byl4bwbha****"),
		InstanceTypeId: volcengine.String("ecs.g1.large"),
	}

	resp, err := svc.ModifyInstanceSpec(modifyInstanceSpecInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
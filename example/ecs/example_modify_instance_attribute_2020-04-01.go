package ecs_example

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/service/ecs"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
    "github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)
func ModifyInstanceAttribute() {
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
	modifyInstanceAttributeInput := &ecs.ModifyInstanceAttributeInput{
		InstanceId: volcengine.String("i-l8u10sauiu9qj0h*****"),
		Password: volcengine.String("password@123"),
	}

	resp, err := svc.ModifyInstanceAttribute(modifyInstanceAttributeInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}

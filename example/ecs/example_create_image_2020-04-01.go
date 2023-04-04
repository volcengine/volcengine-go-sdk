package ecs_example

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/service/ecs"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)

func CreateImage() {
	ak, sk, region := "Your AK", "Your SK", "Region"
	config := volcengine.NewConfig().
		WithRegion(region).
		WithCredentials(credentials.NewStaticCredentials(ak, sk, ""))
	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
	svc := ecs.New(sess)
	createImageInput := &ecs.CreateImageInput{
		ImageName: volcengine.String("image-1"),
		InstanceId: volcengine.String("i-3teco18f1w5a71******"),
	}

	resp, err := svc.CreateImage(createImageInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}

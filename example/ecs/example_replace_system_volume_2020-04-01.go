package ecs_example

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/service/ecs"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)

func ReplaceSystemVolume() {
	ak, sk, region := "Your AK", "Your SK", "Region"
	config := volcengine.NewConfig().
		WithRegion(region).
		WithCredentials(credentials.NewStaticCredentials(ak, sk, ""))
	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
	svc := ecs.New(sess)
	replaceSystemVolumeInput := &ecs.ReplaceSystemVolumeInput{
		ImageId: volcengine.String("image-38dfk6rfisf6kir6****"),
		InstanceId: volcengine.String("i-a8j6n1i4hojfqpa0****"),
		Password: volcengine.String("Password@123"),
	}

	resp, err := svc.ReplaceSystemVolume(replaceSystemVolumeInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}

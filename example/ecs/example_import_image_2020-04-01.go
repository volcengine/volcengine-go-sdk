package ecs_example

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/service/ecs"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)
func ImportImage() {
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
	importImageInput := &ecs.ImportImageInput{
		Architecture: volcengine.String("amd64"),
		ImageName: volcengine.String("image-1"),
		OsType: volcengine.String("Linux"),
		Platform: volcengine.String("CentOS"),
		PlatformVersion: volcengine.String("7.6"),
		Url: volcengine.String("xxx"),
	}

	resp, err := svc.ImportImage(importImageInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}

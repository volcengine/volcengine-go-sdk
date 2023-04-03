package main

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/service/ecs"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
    "github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)

func main() {
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
	reqTags0 := &ecs.TagForCreateTagsInput{
		Key: volcengine.String("k1"),
		Value: volcengine.String("v1"),
	}
	createTagsInput := &ecs.CreateTagsInput{
		ResourceIds: volcengine.StringSlice([]string{"i-l8u0p77yseabkpak****","i-l8u0p7xyseabkbak****"}),
		ResourceType: volcengine.String("instance"),
		Tags: []*ecs.TagForCreateTagsInput{reqTags0},
	}

	resp, err := svc.CreateTags(createTagsInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
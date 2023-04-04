package ecs_example

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/service/ecs"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)
func DeleteTags() {
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
	deleteTagsInput := &ecs.DeleteTagsInput{
		ResourceIds: volcengine.StringSlice([]string{"i-l8u0p77yseabkpak****","i-l8u0p7xyseabkbak****"}),
		ResourceType: volcengine.String("instance"),
		TagKeys: volcengine.StringSlice([]string{"k1"}),
	}

	resp, err := svc.DeleteTags(deleteTagsInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}

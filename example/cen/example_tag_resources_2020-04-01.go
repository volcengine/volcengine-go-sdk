package cen_example

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/service/cen"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)

func TagResources() {
	ak, sk, region := "Your AK", "Your SK", "Region"
	config := volcengine.NewConfig().
		WithRegion(region).
		WithCredentials(credentials.NewStaticCredentials(ak, sk, ""))
	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
	svc := cen.New(sess)
	reqTags0 := &cen.TagForTagResourcesInput{
		Key: volcengine.String("k1"),
		Value: volcengine.String("v1"),
	}
	reqTags1 := &cen.TagForTagResourcesInput{
		Key: volcengine.String("k2"),
		Value: volcengine.String("v2"),
	}
	tagResourcesInput := &cen.TagResourcesInput{
		ResourceIds: volcengine.StringSlice([]string{"cen-273w3e33y2y9s7fap8u2j****","cen-7qthudw0ll6jmc****"}),
		ResourceType: volcengine.String("cen"),
		Tags: []*cen.TagForTagResourcesInput{reqTags0,reqTags1},
	}

	resp, err := svc.TagResources(tagResourcesInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}

package clb_example

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/service/clb"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)

func TagResources() {
	ak, sk, region := "Your AK", "Your SK", "Region"
	config := volcengine.NewConfig().
		WithRegion(region).
		WithCredentials(credentials.NewStaticCredentials(ak, sk, "")).
		WithDisableSSL(true)
	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
	svc := clb.New(sess)
	reqTags0 := &clb.TagForTagResourcesInput{
		Key: volcengine.String("k1"),
		Value: volcengine.String("v1"),
	}
	reqTags1 := &clb.TagForTagResourcesInput{
		Key: volcengine.String("k2"),
		Value: volcengine.String("v2"),
	}
	tagResourcesInput := &clb.TagResourcesInput{
		ResourceIds: volcengine.StringSlice([]string{"clb-273sdsdsxxxxxp8u2j****","clb-2fe6fszjgeznk5oxruv0u****"}),
		ResourceType: volcengine.String("CLB"),
		Tags: []*clb.TagForTagResourcesInput{reqTags0,reqTags1},
	}

	resp, err := svc.TagResources(tagResourcesInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}

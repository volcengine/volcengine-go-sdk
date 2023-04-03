package main

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/service/vpc"
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
	svc := vpc.New(sess)
	reqTags0 := &vpc.TagForTagResourcesInput{
		Key: volcengine.String("k1"),
		Value: volcengine.String("v1"),
	}
	reqTags1 := &vpc.TagForTagResourcesInput{
		Key: volcengine.String("k2"),
		Value: volcengine.String("v2"),
	}
	tagResourcesInput := &vpc.TagResourcesInput{
		ResourceIds: volcengine.StringSlice([]string{"vpc-273w3e33y2y9s7fap8u2j****","vpc-bp15zckdt37pq72zv****"}),
		ResourceType: volcengine.String("vpc"),
		Tags: []*vpc.TagForTagResourcesInput{reqTags0,reqTags1},
	}

	resp, err := svc.TagResources(tagResourcesInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
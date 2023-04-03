package main

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/service/clb"
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
	svc := clb.New(sess)
	listTagsForResourcesInput := &clb.ListTagsForResourcesInput{
		ResourceIds: volcengine.StringSlice([]string{"clb-273sdsdsxxxxxp8u2j****","clb-2fe6fszjgeznk5oxruv0u****"}),
		ResourceType: volcengine.String("CLB"),
	}

	resp, err := svc.ListTagsForResources(listTagsForResourcesInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
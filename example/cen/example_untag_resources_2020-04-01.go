package cen_example

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/service/cen"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)
func UntagResources() {
	ak, sk, region := "Your AK", "Your SK", "Region"
	config := volcengine.NewConfig().
		WithRegion(region).
		WithCredentials(credentials.NewStaticCredentials(ak, sk, "")).
		WithDisableSSL(true)
	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
	svc := cen.New(sess)
	untagResourcesInput := &cen.UntagResourcesInput{
		ResourceIds: volcengine.StringSlice([]string{"cen-273w3e33y2y9s7fap8u2j****","cen-7qthudw0ll6jmc****"}),
		ResourceType: volcengine.String("cen"),
	}

	resp, err := svc.UntagResources(untagResourcesInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}

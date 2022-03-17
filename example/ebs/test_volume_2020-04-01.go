package main

import (
	"fmt"

	"github.com/volcengine/volcstack-go-sdk/service/storageebs"
	"github.com/volcengine/volcstack-go-sdk/volcstack"
	"github.com/volcengine/volcstack-go-sdk/volcstack/session"
	"github.com/volcengine/volcstack-go-sdk/volcstack/volcstackutil"
)

func main() {
	ak := "ak"
	sk := "sk"
	region := "region"

	config := volcstack.NewConfig().
		WithRegion(region).
		WithAkSk(ak, sk).
		WithDisableSSL(true).WithLogLevel(volcstack.LogDebugWithHTTPBody).WithEndpoint(volcstackutil.NewEndpoint().GetEndpoint())
	sess, _ := session.NewSession(config)

	ebsService := storageebs.New(sess)

	id := "i-uuuuuuuuuuu"

	resp, err := ebsService.DescribeVolumes(&storageebs.DescribeVolumesInput{
		InstanceId: &id,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}

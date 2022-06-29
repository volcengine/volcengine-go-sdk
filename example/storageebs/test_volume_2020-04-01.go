package main

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/service/storageebs"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

func main() {
	ak := "ak"
	sk := "sk"
	region := "region"

	config := volcengine.NewConfig().
		WithRegion(region).
		WithAkSk(ak, sk).
		WithDisableSSL(true).WithLogLevel(volcengine.LogDebugWithHTTPBody).WithEndpoint(volcengineutil.NewEndpoint().GetEndpoint())
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

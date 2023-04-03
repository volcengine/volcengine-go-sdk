package main

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/service/clb"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
    "github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)
func DescribeListenerHealth() {
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
	describeListenerHealthInput := &clb.DescribeListenerHealthInput{
		ListenerId: volcengine.String("lsn-2fek3rgsxhrsw5oxruwec****"),
		OnlyUnHealthy: volcengine.String("true"),
	}

	resp, err := svc.DescribeListenerHealth(describeListenerHealthInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
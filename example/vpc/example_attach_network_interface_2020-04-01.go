package main

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/service/vpc"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
    "github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)
func AttachNetworkInterface() {
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
	attachNetworkInterfaceInput := &vpc.AttachNetworkInterfaceInput{
		InstanceId: volcengine.String("i-wijfn35c****"),
		NetworkInterfaceId: volcengine.String("eni-bp1fg655nh68xyz9****"),
	}

	resp, err := svc.AttachNetworkInterface(attachNetworkInterfaceInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
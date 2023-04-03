package main

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/service/cen"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
    "github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)
func GrantInstanceToCen() {
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
	grantInstanceToCenInput := &cen.GrantInstanceToCenInput{
		CenId: volcengine.String("cen-7qthudw0ll6jmc****"),
		CenOwnerId: volcengine.String("210000****"),
		InstanceId: volcengine.String("vpc-uf6o8d1dj8sjwxi6o****"),
		InstanceRegionId: volcengine.String("cn-beijing"),
		InstanceType: volcengine.String("VPC"),
	}

	resp, err := svc.GrantInstanceToCen(grantInstanceToCenInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
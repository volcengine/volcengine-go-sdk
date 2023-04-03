package main

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/service/vpc"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
    "github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)
func DeleteHaVip() {
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
	deleteHaVipInput := &vpc.DeleteHaVipInput{
		HaVipId: volcengine.String("havip-2zeo05qre24nhrqpy****"),
	}

	resp, err := svc.DeleteHaVip(deleteHaVipInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}

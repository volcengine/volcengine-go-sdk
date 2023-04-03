package main

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/service/ecs"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
    "github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)
func ModifyInstanceChargeType() {
	ak, sk, region := "Your AK", "Your SK", "Region"
	config := volcengine.NewConfig().
		WithRegion(region).
		WithCredentials(credentials.NewStaticCredentials(ak, sk, "")).
		WithDisableSSL(true)
	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
	svc := ecs.New(sess)
	modifyInstanceChargeTypeInput := &ecs.ModifyInstanceChargeTypeInput{
		AutoPay: volcengine.Bool(true),
		IncludeDataVolumes: volcengine.Bool(false),
		InstanceChargeType: volcengine.String("PrePaid"),
		InstanceIds: volcengine.StringSlice([]string{"i-4ay51iinvo8w4nho****"}),
		Period: volcengine.Int32(2),
		PeriodUnit: volcengine.String("Month"),
	}

	resp, err := svc.ModifyInstanceChargeType(modifyInstanceChargeTypeInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}

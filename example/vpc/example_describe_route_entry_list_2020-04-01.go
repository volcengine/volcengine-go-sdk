package main

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/service/vpc"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
    "github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)
func DescribeRouteEntryList() {
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
	describeRouteEntryListInput := &vpc.DescribeRouteEntryListInput{
		RouteTableId: volcengine.String("vtb-2fdzao4h726f45oxruw******"),
	}

	resp, err := svc.DescribeRouteEntryList(describeRouteEntryListInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}

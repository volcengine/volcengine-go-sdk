package main

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/service/vpc"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
    "github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)
func CreateRouteEntry() {
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
	createRouteEntryInput := &vpc.CreateRouteEntryInput{
		DestinationCidrBlock: volcengine.String("8.XX.XX.8/16"),
		NextHopId: volcengine.String("NetworkInterface"),
		NextHopType: volcengine.String("eni-2fdzbqxfwrt345oxru******"),
		RouteTableId: volcengine.String("vtb-2fdzao4h726f45oxruw******"),
	}

	resp, err := svc.CreateRouteEntry(createRouteEntryInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}

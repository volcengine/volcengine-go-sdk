package main

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/service/cen"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
    "github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)
func PublishCenRouteEntry() {
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
	publishCenRouteEntryInput := &cen.PublishCenRouteEntryInput{
		CenId: volcengine.String("cen-2nim00ybaylts7trquyzt****"),
		DestinationCidrBlock: volcengine.String("10.XX.XX.0/24"),
		InstanceId: volcengine.String("vpc-vtbnbb04qw3k2hgi12cv****"),
		InstanceRegionId: volcengine.String("cn-beijing"),
		InstanceType: volcengine.String("VPC"),
	}

	resp, err := svc.PublishCenRouteEntry(publishCenRouteEntryInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}

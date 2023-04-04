package cen_example

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/service/cen"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)

func CreateCenBandwidthPackage() {
	ak, sk, region := "Your AK", "Your SK", "Region"
	config := volcengine.NewConfig().
		WithRegion(region).
		WithCredentials(credentials.NewStaticCredentials(ak, sk, ""))
	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
	svc := cen.New(sess)
	createCenBandwidthPackageInput := &cen.CreateCenBandwidthPackageInput{
		CenBandwidthPackageName:    volcengine.String("cbp-test"),
		CenId:                      volcengine.String("cen-7qthudw0ll6jmc****"),
		Description:                volcengine.String("namedesc"),
		LocalGeographicRegionSetId: volcengine.String("China"),
		PeerGeographicRegionSetId:  volcengine.String("China"),
	}

	resp, err := svc.CreateCenBandwidthPackage(createCenBandwidthPackageInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}

package main

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/service/cen"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
    "github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)
func DeleteCenBandwidthPackage() {
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
	deleteCenBandwidthPackageInput := &cen.DeleteCenBandwidthPackageInput{
		CenBandwidthPackageId: volcengine.String("cbp-4c2zaavbvh5f42****"),
	}

	resp, err := svc.DeleteCenBandwidthPackage(deleteCenBandwidthPackageInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
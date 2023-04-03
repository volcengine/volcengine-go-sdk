package main

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/service/cen"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
    "github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)
func DeleteCenInterRegionBandwidth() {
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
	deleteCenInterRegionBandwidthInput := &cen.DeleteCenInterRegionBandwidthInput{
		InterRegionBandwidthId: volcengine.String("cirb-3tex2x1cwd4c6c0v****"),
	}

	resp, err := svc.DeleteCenInterRegionBandwidth(deleteCenInterRegionBandwidthInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
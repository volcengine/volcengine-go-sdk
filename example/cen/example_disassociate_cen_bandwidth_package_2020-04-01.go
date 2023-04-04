package cen_example

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/service/cen"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)

func DisassociateCenBandwidthPackage() {
	ak, sk, region := "Your AK", "Your SK", "Region"
	config := volcengine.NewConfig().
		WithRegion(region).
		WithCredentials(credentials.NewStaticCredentials(ak, sk, ""))
	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
	svc := cen.New(sess)
	disassociateCenBandwidthPackageInput := &cen.DisassociateCenBandwidthPackageInput{
		CenBandwidthPackageId: volcengine.String("cbp-4c2zaavbvh5fx****"),
		CenId: volcengine.String("cen-7qthudw0ll6jmc****"),
	}

	resp, err := svc.DisassociateCenBandwidthPackage(disassociateCenBandwidthPackageInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}

package cen_example

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/service/cen"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)

func DescribeCenInterRegionBandwidths() {
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
	describeCenInterRegionBandwidthsInput := &cen.DescribeCenInterRegionBandwidthsInput{
		PageSize: volcengine.String("20"),
	}

	resp, err := svc.DescribeCenInterRegionBandwidths(describeCenInterRegionBandwidthsInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}

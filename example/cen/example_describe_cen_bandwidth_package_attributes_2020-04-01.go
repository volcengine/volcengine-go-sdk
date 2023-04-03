package main

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/service/cen"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
    "github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)

func main() {
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
	describeCenBandwidthPackageAttributesInput := &cen.DescribeCenBandwidthPackageAttributesInput{
		CenBandwidthPackageId: volcengine.String("cbp-bp1o94dp5****"),
	}

	resp, err := svc.DescribeCenBandwidthPackageAttributes(describeCenBandwidthPackageAttributesInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
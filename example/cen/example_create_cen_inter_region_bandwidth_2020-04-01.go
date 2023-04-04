package cenexample

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/service/cen"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)

func CreateCenInterRegionBandwidth() {
	ak, sk, region := "Your AK", "Your SK", "Region"
	config := volcengine.NewConfig().
		WithRegion(region).
		WithCredentials(credentials.NewStaticCredentials(ak, sk, ""))
	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
	svc := cen.New(sess)
	createCenInterRegionBandwidthInput := &cen.CreateCenInterRegionBandwidthInput{
		Bandwidth:     volcengine.Int64(1000),
		CenId:         volcengine.String("cen-7qthudw0ll6jmc****"),
		LocalRegionId: volcengine.String("cn-beijing"),
		PeerRegionId:  volcengine.String("cn-nantong"),
	}

	resp, err := svc.CreateCenInterRegionBandwidth(createCenInterRegionBandwidthInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}

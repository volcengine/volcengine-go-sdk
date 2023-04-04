package vpcexample

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/service/vpc"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)

func ReleaseEipAddress() {
	ak, sk, region := "Your AK", "Your SK", "Region"
	config := volcengine.NewConfig().
		WithRegion(region).
		WithCredentials(credentials.NewStaticCredentials(ak, sk, ""))
	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
	svc := vpc.New(sess)
	releaseEipAddressInput := &vpc.ReleaseEipAddressInput{
		AllocationId: volcengine.String("eip-2zewb7ujxscd****"),
	}

	resp, err := svc.ReleaseEipAddress(releaseEipAddressInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}

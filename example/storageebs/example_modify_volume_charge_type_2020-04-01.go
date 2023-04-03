package main

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/service/storageebs"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
    "github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)
func ModifyVolumeChargeType() {
	ak, sk, region := "Your AK", "Your SK", "Region"
	config := volcengine.NewConfig().
		WithRegion(region).
		WithCredentials(credentials.NewStaticCredentials(ak, sk, "")).
		WithDisableSSL(true)
	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
	svc := storageebs.New(sess)
	modifyVolumeChargeTypeInput := &storageebs.ModifyVolumeChargeTypeInput{
		DiskChargeType: volcengine.String("PrePaid"),
		InstanceId: volcengine.String("i-l2soxrgan390t1dk****"),
		VolumeIds: volcengine.StringSlice([]string{"vol-3tiz8tg97u3vj0x0****","vol-3tj20xxpd63vj0wy****"}),
	}

	resp, err := svc.ModifyVolumeChargeType(modifyVolumeChargeTypeInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}

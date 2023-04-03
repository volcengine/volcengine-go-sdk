package main

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/service/storageebs"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
    "github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)
func AttachVolume() {
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
	attachVolumeInput := &storageebs.AttachVolumeInput{
		InstanceId: volcengine.String("i-3tir90q84q3vj0x1****"),
		VolumeId: volcengine.String("vol-3tir92lne23vj0x1****"),
	}

	resp, err := svc.AttachVolume(attachVolumeInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}

package main

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/service/storageebs"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
    "github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)
func ModifyVolumeAttribute() {
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
	modifyVolumeAttributeInput := &storageebs.ModifyVolumeAttributeInput{
		DeleteWithInstance: volcengine.Bool(true),
		VolumeId: volcengine.String("vol-3tir92lne23vj0x1****"),
		VolumeName: volcengine.String("dtest3-volume-1"),
	}

	resp, err := svc.ModifyVolumeAttribute(modifyVolumeAttributeInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}

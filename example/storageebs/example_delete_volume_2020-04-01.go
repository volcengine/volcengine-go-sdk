package storageebsexample

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/service/storageebs"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)

func DeleteVolume() {
	ak, sk, region := "Your AK", "Your SK", "Region"
	config := volcengine.NewConfig().
		WithRegion(region).
		WithCredentials(credentials.NewStaticCredentials(ak, sk, ""))
	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
	svc := storageebs.New(sess)
	deleteVolumeInput := &storageebs.DeleteVolumeInput{
		VolumeId: volcengine.String("vol-3tir92lne23vj0x1****"),
	}

	resp, err := svc.DeleteVolume(deleteVolumeInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}

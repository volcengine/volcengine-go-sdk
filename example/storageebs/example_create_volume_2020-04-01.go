package storageebs_example

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/service/storageebs"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)

func CreateVolume() {
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
	createVolumeInput := &storageebs.CreateVolumeInput{
		Description: volcengine.String("test"),
		Kind: volcengine.String("data"),
		Size: volcengine.JsonNumber("40"),
		VolumeName: volcengine.String("test"),
		VolumeType: volcengine.String("PTSSD"),
		ZoneId: volcengine.String("cn-beijing-a"),
	}

	resp, err := svc.CreateVolume(createVolumeInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}

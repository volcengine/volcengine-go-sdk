package storageebs_example

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/service/storageebs"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
    "github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)
func ExtendVolume() {
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
	extendVolumeInput := &storageebs.ExtendVolumeInput{
		NewSize: volcengine.JsonNumber("80"),
		VolumeId: volcengine.String("vol-3thhlu8byl4bwbha****"),
	}

	resp, err := svc.ExtendVolume(extendVolumeInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}

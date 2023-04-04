package storageebs_example

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/service/storageebs"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
    "github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)
func DescribeVolumes() {
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
	describeVolumesInput := &storageebs.DescribeVolumesInput{
		VolumeIds: volcengine.StringSlice([]string{"vol-3tirbh6lka3vj0x1****"}),
	}

	resp, err := svc.DescribeVolumes(describeVolumesInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}

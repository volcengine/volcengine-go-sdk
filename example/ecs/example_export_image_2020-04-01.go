package main

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/service/ecs"
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
	svc := ecs.New(sess)
	exportImageInput := &ecs.ExportImageInput{
		ImageId: volcengine.String("image-4431h3l7hl31a0******"),
		TOSBucket: volcengine.String("bucket-1"),
		TOSPrefix: volcengine.String("test123"),
	}

	resp, err := svc.ExportImage(exportImageInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
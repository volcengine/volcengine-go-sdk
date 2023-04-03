package main

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/service/cen"
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
	svc := cen.New(sess)
	modifyCenAttributesInput := &cen.ModifyCenAttributesInput{
		CenId: volcengine.String("cen-7qthudw0ll6jmc****"),
		CenName: volcengine.String("cen-test"),
		Description: volcengine.String("test"),
	}

	resp, err := svc.ModifyCenAttributes(modifyCenAttributesInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
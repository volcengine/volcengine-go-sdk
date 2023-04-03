package main

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/service/clb"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
    "github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)
func ModifyListenerAttributes() {
	ak, sk, region := "Your AK", "Your SK", "Region"
	config := volcengine.NewConfig().
		WithRegion(region).
		WithCredentials(credentials.NewStaticCredentials(ak, sk, "")).
		WithDisableSSL(true)
	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
	svc := clb.New(sess)
	modifyListenerAttributesInput := &clb.ModifyListenerAttributesInput{
		ListenerId: volcengine.String("lsn-2fek3rgsxhrsw5oxruwec****"),
		ListenerName: volcengine.String("test"),
		Scheduler: volcengine.String("wlc"),
	}

	resp, err := svc.ModifyListenerAttributes(modifyListenerAttributesInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
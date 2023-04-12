package clbexample

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/service/clb"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)

func DeleteListener() {
	ak, sk, region := "Your AK", "Your SK", "Region"
	config := volcengine.NewConfig().
		WithRegion(region).
		WithCredentials(credentials.NewStaticCredentials(ak, sk, ""))
	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
	svc := clb.New(sess)
	deleteListenerInput := &clb.DeleteListenerInput{
		ListenerId: volcengine.String("lsn-2fek3rgsxhrsw5oxruwec****"),
	}

	resp, err := svc.DeleteListener(deleteListenerInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}

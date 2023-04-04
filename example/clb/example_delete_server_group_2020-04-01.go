package clbexample

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/service/clb"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)

func DeleteServerGroup() {
	ak, sk, region := "Your AK", "Your SK", "Region"
	config := volcengine.NewConfig().
		WithRegion(region).
		WithCredentials(credentials.NewStaticCredentials(ak, sk, ""))
	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
	svc := clb.New(sess)
	deleteServerGroupInput := &clb.DeleteServerGroupInput{
		ServerGroupId: volcengine.String("rsp-bp1o94dp5i6ea****"),
	}

	resp, err := svc.DeleteServerGroup(deleteServerGroupInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}

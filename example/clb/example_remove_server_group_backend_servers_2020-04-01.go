package clb_example

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/service/clb"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)

func RemoveServerGroupBackendServers() {
	ak, sk, region := "Your AK", "Your SK", "Region"
	config := volcengine.NewConfig().
		WithRegion(region).
		WithCredentials(credentials.NewStaticCredentials(ak, sk, ""))
	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
	svc := clb.New(sess)
	removeServerGroupBackendServersInput := &clb.RemoveServerGroupBackendServersInput{
		ServerGroupId: volcengine.String("rsp-bp1o94dp5i6ea****"),
		ServerIds: volcengine.StringSlice([]string{"rs-3cig8e5o0kxs06c6rrsqn****"}),
	}

	resp, err := svc.RemoveServerGroupBackendServers(removeServerGroupBackendServersInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}

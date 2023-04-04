package clb_example

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/service/clb"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)

func AddServerGroupBackendServers() {
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
	reqServers0 := &clb.ServerForAddServerGroupBackendServersInput{
		Description: volcengine.String("ecs1"),
		InstanceId: volcengine.String("i-3tkuehz8oa3vj0wz****"),
		Ip: volcengine.String("192.XX.XX.2"),
		Type: volcengine.String("ecs"),
	}
	addServerGroupBackendServersInput := &clb.AddServerGroupBackendServersInput{
		ServerGroupId: volcengine.String("rsp-bp1o94dp5i6ea****"),
		Servers: []*clb.ServerForAddServerGroupBackendServersInput{reqServers0},
	}

	resp, err := svc.AddServerGroupBackendServers(addServerGroupBackendServersInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}

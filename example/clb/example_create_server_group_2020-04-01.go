package clb_example

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/service/clb"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)

func CreateServerGroup() {
	ak, sk, region := "Your AK", "Your SK", "Region"
	config := volcengine.NewConfig().
		WithRegion(region).
		WithCredentials(credentials.NewStaticCredentials(ak, sk, ""))
	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
	svc := clb.New(sess)
	reqServers0 := &clb.ServerForCreateServerGroupInput{
		InstanceId: volcengine.String("i-3tkuehz8oa3vj0wz****"),
		Ip:         volcengine.String("192.XX.XX.2"),
		Port:       volcengine.Int64(88),
		Type:       volcengine.String("ecs"),
		Weight:     volcengine.Int64(100),
	}
	createServerGroupInput := &clb.CreateServerGroupInput{
		LoadBalancerId:  volcengine.String("clb-bp1b6c719dfa08ex****"),
		ServerGroupName: volcengine.String("myservergroup"),
		Servers:         []*clb.ServerForCreateServerGroupInput{reqServers0},
	}

	resp, err := svc.CreateServerGroup(createServerGroupInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}

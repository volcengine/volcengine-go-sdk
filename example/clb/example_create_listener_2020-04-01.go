package clb_example

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/service/clb"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)

func CreateListener() {
	ak, sk, region := "Your AK", "Your SK", "Region"
	config := volcengine.NewConfig().
		WithRegion(region).
		WithCredentials(credentials.NewStaticCredentials(ak, sk, ""))
	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
	svc := clb.New(sess)
	createListenerInput := &clb.CreateListenerInput{
		AclIds:             volcengine.StringSlice([]string{"acl-3cj44nv0jhhxc6c6rrtet****"}),
		AclStatus:          volcengine.String("on"),
		AclType:            volcengine.String("black"),
		Enabled:            volcengine.String("on"),
		EstablishedTimeout: volcengine.Int64(122),
		ListenerName:       volcengine.String("mylistener"),
		LoadBalancerId:     volcengine.String("clb-bp1o94dp5i6ea****"),
		Port:               volcengine.Int64(12),
		Protocol:           volcengine.String("TCP"),
		Scheduler:          volcengine.String("wrr"),
		ServerGroupId:      volcengine.String("rsp-bp1o94dp5i6ea****"),
	}

	resp, err := svc.CreateListener(createListenerInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}

package clbexample

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/service/clb"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)

func DescribeListeners() {
	ak, sk, region := "Your AK", "Your SK", "Region"
	config := volcengine.NewConfig().
		WithRegion(region).
		WithCredentials(credentials.NewStaticCredentials(ak, sk, ""))
	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
	svc := clb.New(sess)
	describeListenersInput := &clb.DescribeListenersInput{
		ListenerIds:    volcengine.StringSlice([]string{"lsn-2fek3rgsxhrsw5oxr****"}),
		ListenerName:   volcengine.String("test"),
		LoadBalancerId: volcengine.String("clb-bp1o94dp5i6ea****"),
		PageNumber:     volcengine.Int64(1),
		PageSize:       volcengine.Int64(20),
	}

	resp, err := svc.DescribeListeners(describeListenersInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}

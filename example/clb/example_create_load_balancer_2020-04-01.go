package clb_example

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/service/clb"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)

func CreateLoadBalancer() {
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
	createLoadBalancerInput := &clb.CreateLoadBalancerInput{
		LoadBalancerName: volcengine.String("clb-test"),
		LoadBalancerSpec: volcengine.String("small_1"),
		Type: volcengine.String("private"),
		VpcId: volcengine.String("vpc-bp1aevy8sofi8mh1****"),
	}

	resp, err := svc.CreateLoadBalancer(createLoadBalancerInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}

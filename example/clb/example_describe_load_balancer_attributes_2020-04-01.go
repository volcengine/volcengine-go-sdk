package main

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/service/clb"
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
	svc := clb.New(sess)
	describeLoadBalancerAttributesInput := &clb.DescribeLoadBalancerAttributesInput{
		LoadBalancerId: volcengine.String("clb-bp1b6c719dfa08ex****"),
	}

	resp, err := svc.DescribeLoadBalancerAttributes(describeLoadBalancerAttributesInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
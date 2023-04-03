package main

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/service/clb"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
    "github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)
func DescribeLoadBalancersBilling() {
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
	describeLoadBalancersBillingInput := &clb.DescribeLoadBalancersBillingInput{
		LoadBalancerIds: volcengine.StringSlice([]string{"clb-bp1b6c719dfa08ex****"}),
	}

	resp, err := svc.DescribeLoadBalancersBilling(describeLoadBalancersBillingInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}

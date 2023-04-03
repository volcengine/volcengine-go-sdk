package main

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/service/clb"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
    "github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)
func DescribeRules() {
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
	describeRulesInput := &clb.DescribeRulesInput{
		ListenerId: volcengine.String("lsn-2fea4ayvu2g3k5oxruuz****"),
	}

	resp, err := svc.DescribeRules(describeRulesInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}

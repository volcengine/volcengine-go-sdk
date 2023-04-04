package cen_example

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/service/cen"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
    "github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)
func DescribeInstanceGrantedRules() {
	ak, sk, region := "Your AK", "Your SK", "Region"
	config := volcengine.NewConfig().
		WithRegion(region).
		WithCredentials(credentials.NewStaticCredentials(ak, sk, "")).
		WithDisableSSL(true)
	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
	svc := cen.New(sess)
	describeInstanceGrantedRulesInput := &cen.DescribeInstanceGrantedRulesInput{
		InstanceId: volcengine.String("vpc-3rlkeggyn6tc010e****"),
		InstanceType: volcengine.String("VPC"),
	}

	resp, err := svc.DescribeInstanceGrantedRules(describeInstanceGrantedRulesInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}

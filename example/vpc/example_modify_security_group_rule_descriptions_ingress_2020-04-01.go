package main

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/service/vpc"
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
	svc := vpc.New(sess)
	modifySecurityGroupRuleDescriptionsIngressInput := &vpc.ModifySecurityGroupRuleDescriptionsIngressInput{
		CidrIp: volcengine.String("10.XX.XX.0/8"),
		Protocol: volcengine.String("tcp"),
	}

	resp, err := svc.ModifySecurityGroupRuleDescriptionsIngress(modifySecurityGroupRuleDescriptionsIngressInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
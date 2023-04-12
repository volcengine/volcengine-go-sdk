package vpcexample

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/service/vpc"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)

func AuthorizeSecurityGroupEgress() {
	ak, sk, region := "Your AK", "Your SK", "Region"
	config := volcengine.NewConfig().
		WithRegion(region).
		WithCredentials(credentials.NewStaticCredentials(ak, sk, ""))
	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
	svc := vpc.New(sess)
	authorizeSecurityGroupEgressInput := &vpc.AuthorizeSecurityGroupEgressInput{
		CidrIp:          volcengine.String("10.XX.XX.0/8"),
		Policy:          volcengine.String("accept"),
		PortEnd:         volcengine.Int64(22),
		PortStart:       volcengine.Int64(22),
		Priority:        volcengine.Int64(1),
		Protocol:        volcengine.String("tcp"),
		SecurityGroupId: volcengine.String("sg-bp67acfmxazb4p****"),
	}

	resp, err := svc.AuthorizeSecurityGroupEgress(authorizeSecurityGroupEgressInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}

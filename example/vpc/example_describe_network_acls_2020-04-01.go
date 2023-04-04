package vpcexample

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/service/vpc"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)

func DescribeNetworkAcls() {
	ak, sk, region := "Your AK", "Your SK", "Region"
	config := volcengine.NewConfig().
		WithRegion(region).
		WithCredentials(credentials.NewStaticCredentials(ak, sk, ""))
	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
	svc := vpc.New(sess)
	describeNetworkAclsInput := &vpc.DescribeNetworkAclsInput{
		NetworkAclIds:  volcengine.StringSlice([]string{"acl-bp1fg655nh68xyz9****"}),
		NetworkAclName: volcengine.String("test-acl"),
		PageNumber:     volcengine.Int64(1),
		PageSize:       volcengine.Int64(20),
		SubnetId:       volcengine.String("subnet-087k1y0owv0x57ku****"),
		VpcId:          volcengine.String("vpc-bp1opxu1zkhn00gzv****"),
	}

	resp, err := svc.DescribeNetworkAcls(describeNetworkAclsInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}

package main

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/service/vpc"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
    "github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)
func DisassociateVpcCidrBlock() {
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
	disassociateVpcCidrBlockInput := &vpc.DisassociateVpcCidrBlockInput{
		VpcId: volcengine.String("vpc-257gqcdfvx6n****"),
	}

	resp, err := svc.DisassociateVpcCidrBlock(disassociateVpcCidrBlockInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}

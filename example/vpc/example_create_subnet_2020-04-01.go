package vpcexample

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/service/vpc"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)

func CreateSubnet() {
	ak, sk, region := "Your AK", "Your SK", "Region"
	config := volcengine.NewConfig().
		WithRegion(region).
		WithCredentials(credentials.NewStaticCredentials(ak, sk, ""))
	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
	svc := vpc.New(sess)
	createSubnetInput := &vpc.CreateSubnetInput{
		CidrBlock: volcengine.String("172.XX.XX.0/24"),
		VpcId:     volcengine.String("vpc-257gqcdfvx6n****"),
		ZoneId:    volcengine.String("cn-beijing-a"),
	}

	resp, err := svc.CreateSubnet(createSubnetInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}

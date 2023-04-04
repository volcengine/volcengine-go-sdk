package vpcexample

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/service/vpc"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)

func AssociateRouteTable() {
	ak, sk, region := "Your AK", "Your SK", "Region"
	config := volcengine.NewConfig().
		WithRegion(region).
		WithCredentials(credentials.NewStaticCredentials(ak, sk, ""))
	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
	svc := vpc.New(sess)
	associateRouteTableInput := &vpc.AssociateRouteTableInput{
		RouteTableId: volcengine.String("vtb-2fdzao4h726f45****"),
		SubnetId:     volcengine.String("subnet-2fdzaou4liw3k5oxruv****"),
	}

	resp, err := svc.AssociateRouteTable(associateRouteTableInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}

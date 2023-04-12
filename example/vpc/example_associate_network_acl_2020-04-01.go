package vpcexample

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/service/vpc"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)

func AssociateNetworkAcl() {
	ak, sk, region := "Your AK", "Your SK", "Region"
	config := volcengine.NewConfig().
		WithRegion(region).
		WithCredentials(credentials.NewStaticCredentials(ak, sk, ""))
	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
	svc := vpc.New(sess)
	reqResource0 := &vpc.ResourceForAssociateNetworkAclInput{
		ResourceId: volcengine.String("subnet-67acfmxazb4p****"),
	}
	reqResource1 := &vpc.ResourceForAssociateNetworkAclInput{
		ResourceId: volcengine.String("subnet-g655nh68xyz9****"),
	}
	associateNetworkAclInput := &vpc.AssociateNetworkAclInput{
		NetworkAclId: volcengine.String("nacl-bp1fg655nh68xyz9****"),
		Resource:     []*vpc.ResourceForAssociateNetworkAclInput{reqResource0, reqResource1},
	}

	resp, err := svc.AssociateNetworkAcl(associateNetworkAclInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}

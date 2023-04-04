package vpc_example

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/service/vpc"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)

func DisassociateNetworkAcl() {
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
	reqResource0 := &vpc.ResourceForDisassociateNetworkAclInput{
		ResourceId: volcengine.String("subnet-67acfmxazb4p****"),
	}
	reqResource1 := &vpc.ResourceForDisassociateNetworkAclInput{
		ResourceId: volcengine.String("subnet-g655nh68xyz9****"),
	}
	disassociateNetworkAclInput := &vpc.DisassociateNetworkAclInput{
		NetworkAclId: volcengine.String("nacl-bp1fg655nh68xyz9****"),
		Resource: []*vpc.ResourceForDisassociateNetworkAclInput{reqResource0,reqResource1},
	}

	resp, err := svc.DisassociateNetworkAcl(disassociateNetworkAclInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}

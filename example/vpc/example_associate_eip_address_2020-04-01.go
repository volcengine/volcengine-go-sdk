package vpc_example

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/service/vpc"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)

func AssociateEipAddress() {
	ak, sk, region := "Your AK", "Your SK", "Region"
	config := volcengine.NewConfig().
		WithRegion(region).
		WithCredentials(credentials.NewStaticCredentials(ak, sk, ""))
	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
	svc := vpc.New(sess)
	associateEipAddressInput := &vpc.AssociateEipAddressInput{
		AllocationId: volcengine.String("eip-2zeewb7ujxscd****"),
		InstanceId:   volcengine.String("i-2zebbhyczzaweeval****"),
		InstanceType: volcengine.String("EcsInstance"),
	}

	resp, err := svc.AssociateEipAddress(associateEipAddressInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}

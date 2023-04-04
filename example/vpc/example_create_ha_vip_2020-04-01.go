package vpc_example

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/service/vpc"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)

func CreateHaVip() {
	ak, sk, region := "Your AK", "Your SK", "Region"
	config := volcengine.NewConfig().
		WithRegion(region).
		WithCredentials(credentials.NewStaticCredentials(ak, sk, ""))
	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
	svc := vpc.New(sess)
	createHaVipInput := &vpc.CreateHaVipInput{
		HaVipName: volcengine.String("HaVip-test"),
		IpAddress: volcengine.String("192.XX.XX.10"),
		SubnetId:  volcengine.String("subnet-ina9r9xnfpc08gbs****"),
	}

	resp, err := svc.CreateHaVip(createHaVipInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}

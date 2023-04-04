package vpc_example

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/service/vpc"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)

func ModifyVpcAttributes() {
	ak, sk, region := "Your AK", "Your SK", "Region"
	config := volcengine.NewConfig().
		WithRegion(region).
		WithCredentials(credentials.NewStaticCredentials(ak, sk, ""))
	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
	svc := vpc.New(sess)
	modifyVpcAttributesInput := &vpc.ModifyVpcAttributesInput{
		Description: volcengine.String("test"),
		DnsServers: volcengine.StringSlice([]string{"1.XX.XX.1"}),
		VpcId: volcengine.String("vpc-bp15zct37pq72zv****"),
		VpcName: volcengine.String("test"),
	}

	resp, err := svc.ModifyVpcAttributes(modifyVpcAttributesInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}

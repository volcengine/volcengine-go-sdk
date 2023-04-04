package natgateway_example

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/service/natgateway"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)
func CreateSnatEntry() {
	ak, sk, region := "Your AK", "Your SK", "Region"
	config := volcengine.NewConfig().
		WithRegion(region).
		WithCredentials(credentials.NewStaticCredentials(ak, sk, "")).
		WithDisableSSL(true)
	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
	svc := natgateway.New(sess)
	createSnatEntryInput := &natgateway.CreateSnatEntryInput{
		EipId: volcengine.String("eip-2feaac9wtccn459gp67****"),
		NatGatewayId: volcengine.String("ngw-2fedgzyvtzaio59gp675****"),
		SnatEntryName: volcengine.String("snat-01"),
		SubnetId: volcengine.String("subnet-2fe1vklp295a859gp6766****"),
	}

	resp, err := svc.CreateSnatEntry(createSnatEntryInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}

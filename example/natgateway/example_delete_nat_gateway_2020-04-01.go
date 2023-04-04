package natgatewayexample

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/service/natgateway"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)

func DeleteNatGateway() {
	ak, sk, region := "Your AK", "Your SK", "Region"
	config := volcengine.NewConfig().
		WithRegion(region).
		WithCredentials(credentials.NewStaticCredentials(ak, sk, ""))
	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
	svc := natgateway.New(sess)
	deleteNatGatewayInput := &natgateway.DeleteNatGatewayInput{
		NatGatewayId: volcengine.String("ngw-vv3t043k05sm****"),
	}

	resp, err := svc.DeleteNatGateway(deleteNatGatewayInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}

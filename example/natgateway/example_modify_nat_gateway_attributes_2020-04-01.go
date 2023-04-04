package natgateway_example

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/service/natgateway"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)

func ModifyNatGatewayAttributes() {
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
	modifyNatGatewayAttributesInput := &natgateway.ModifyNatGatewayAttributesInput{
		NatGatewayId: volcengine.String("ngw-2fedgzyvtzaio59gp675l****"),
		Spec: volcengine.String("Small"),
	}

	resp, err := svc.ModifyNatGatewayAttributes(modifyNatGatewayAttributesInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}

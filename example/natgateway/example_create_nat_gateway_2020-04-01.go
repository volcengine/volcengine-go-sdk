package natgateway_example

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/service/natgateway"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
    "github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)
func CreateNatGateway() {
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
	createNatGatewayInput := &natgateway.CreateNatGatewayInput{
		Spec: volcengine.String("Small"),
		SubnetId: volcengine.String("subnet-2feypga30rgg059gp67ag****"),
		VpcId: volcengine.String("vpc-2feypfmenesqo59gp67yz****"),
	}

	resp, err := svc.CreateNatGateway(createNatGatewayInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}

package natgatewayexample

import (
	"fmt"

	"github.com/volcengine/volcengine-go-sdk/service/natgateway"
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/credentials"
	"github.com/volcengine/volcengine-go-sdk/volcengine/session"
)

func DescribeNatGatewayAttributes() {
	ak, sk, region := "Your AK", "Your SK", "Region"
	config := volcengine.NewConfig().
		WithRegion(region).
		WithCredentials(credentials.NewStaticCredentials(ak, sk, ""))
	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
	svc := natgateway.New(sess)
	describeNatGatewayAttributesInput := &natgateway.DescribeNatGatewayAttributesInput{
		NatGatewayId: volcengine.String("ngw-2fesmko5zhdz459gp67sc****"),
	}

	resp, err := svc.DescribeNatGatewayAttributes(describeNatGatewayAttributesInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}